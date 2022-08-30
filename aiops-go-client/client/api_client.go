//The api client for aiops's golang SDK
package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	basic   = "basic"
	eTag    = "ETag"
	ifMatch = "If-Match"
)

var (
	jsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
	uriCheck        = regexp.MustCompile(`/(?P<namespace>[-\w]+)/v\d+\.\d+(\.[a|b]\d+)?/(?P<suffix>.*)`)
	retryStatusList = []int{408, 503, 504}
	userAgent       = "Nutanix-aiops/v4.0.1-alpha.1"
)

/**
  Generic API client for Swagger client library builds.
  Swagger generic API client. This client handles the client-
  server communication, and is invariant across implementations. Specifics of
  the methods and models for each application are generated from the Swagger
  templates.

  Parameters :-
    Host (string) : Host IPV4, IPV6 or FQDN for all http request made by this client (default : localhost)
    Port (string) : Port for the host to connect to make all http request (default : 9440)
    Username (string) : Username to connect to a cluster
    Password (string) : Password to connect to a cluster
    Debug (bool) : flag to enable debug logging (default : empty)
    VerifySSL (bool) : Verify SSL certificate of cluster (default: true)
    MaxRetryAttempts (int) : Maximum number of retry attempts to be made at a time (default: 5)
    Timeout (time.Duration) : Global timeout for all operations
    RetryInterval (time.Duration) : Interval between successive retry attempts (default: 3 sec)
    LoggerFile (string) : Log file to write activity logs
*/
type ApiClient struct {
	Host             string            `json:"host,omitempty"`
	Port             int               `json:"port,omitempty"`
	Username         string            `json:"username,omitempty"`
	Password         string            `json:"password,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	VerifySSL        bool
	MaxRetryAttempts int               `json:"maxRetryAttempts,omitempty"`
	Timeout          time.Duration     `json:"timeout,omitempty"`
	RetryInterval    time.Duration     `json:"retryInterval,omitempty"`
	LoggerFile       string            `json:"loggerFile,omitempty"`
	defaultHeaders   map[string]string
	retryClient      *retryablehttp.Client
	httpClient       *http.Client
	authentication   map[string]interface{}
	cookie           string
	refreshCookie    bool
	previousAuth     string
	basicAuth        *BasicAuth
}

/**
  Returns a newly generated ApiClient instance populated with default values
*/
func NewApiClient() *ApiClient {

	basicAuth := new(BasicAuth)
	authentication := make(map[string]interface{})
	authentication["basicAuthScheme"] = basicAuth

	a := &ApiClient{
		Host:             "localhost",
		Port:             9440,
		Debug:            false,
		VerifySSL:        true,
		MaxRetryAttempts: 5,
		Timeout:          30 * time.Second,
		RetryInterval:    3 * time.Second,
		defaultHeaders:   make(map[string]string),
		refreshCookie:    true,
		basicAuth:        basicAuth,
		authentication:   authentication,
	}

	a.setupClient()
	return a
}

/*AddDefaultHeader
Adds a default header to current api client instance for all the HTTP calls.
*/
func (a *ApiClient) AddDefaultHeader(headerName string, headerValue string) {
	if headerName == "Authorization" {
		a.cookie = ""
	}

	a.defaultHeaders[headerName] = headerValue
}

// Makes the HTTP request with given options and returns response body as byte array.
func (a *ApiClient) CallApi(uri *string, httpMethod string, body interface{},
	queryParams url.Values, headerParams map[string]string, formParams url.Values,
	accepts []string, contentType []string, authNames []string) ([]byte, error) {
	path := "https://" + a.Host + ":" + strconv.Itoa(a.Port) + *uri

	if headerParams["Authorization"] != "" {
		a.previousAuth = headerParams["Authorization"]
	}

	if a.defaultHeaders["Authorization"] != "" {
		a.previousAuth = a.defaultHeaders["Authorization"]
	}

	// set Content-Type header
	httpContentType := a.selectHeaderContentType(contentType)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

	// set Accept header
	httpHeaderAccept := a.selectHeaderAccept(accepts)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}

	// set NTNX-Request-Id header
	_, requestIdHeaderExists := headerParams["NTNX-Request-Id"]
	_, requestIdDefaultHeaderExists := a.defaultHeaders["NTNX-Request-Id"]
	if !requestIdHeaderExists && !requestIdDefaultHeaderExists {
		headerParams["NTNX-Request-Id"] = uuid.New().String()
	}

	bodyValue := reflect.ValueOf(body)
	if bodyValue.IsValid() && !bodyValue.IsNil() {
		addEtagReferenceToHeader(body, headerParams)
	}

	request, err := a.prepareRequest(path, httpMethod, body, headerParams, queryParams, formParams, authNames)
	if err != nil {
		return nil, err
	}

	if a.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	a.setupClient()
	response, err := a.httpClient.Do(request)

	// Retry one more time without the cookie but with basic auth header
	if response != nil && response.StatusCode == 401 {
		a.refreshCookie = true
		if len(a.previousAuth) > 0 {
			request.Header["Authorization"] = []string{a.previousAuth}
		}
		delete(request.Header, "Cookie")
		response, err = a.httpClient.Do(request)
	}

	if err != nil {
		return nil, err
	}

	if a.Debug {
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	if nil == response {
		return nil, ReportError("response is nil!")
	}

	a.updateCookies(response)

	if response.StatusCode == 204 {
		return nil, nil
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()
	response.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))

	if !(200 <= response.StatusCode && response.StatusCode <= 209) {
		return nil, GenericOpenAPIError {
			Body:   responseBody,
			Status: response.Status,
		}
	} else {
		responseBody := addEtagReferenceToResponse(response.Header, responseBody)
		return responseBody, nil
	}
}

// Get all authentications (key: authentication name, value: authentication)
func (a *ApiClient) GetAuthentications() map[string]interface{} {
	return a.authentication
}

// Get authentication for the given auth name (eg : basic, oauth, bearer, apiKey)
func (a *ApiClient) GetAuthentication(authName string) interface{} {
	return a.authentication[authName]
}

// Helper method to set username for the first HTTP basic authentication.
func (a *ApiClient) SetUserName(username string) {
	a.Username = username
	a.basicAuth.UserName = username
}

// Helper method to set password for the first HTTP basic authentication
func (a *ApiClient) SetPassword(password string) {
	a.Password = password
	a.basicAuth.Password = password
}

// Helper method to set API key value for the first API key authentication
func (a *ApiClient) SetApiKey(key string) error {
	for _, value := range a.authentication {
		if auth, ok := value.(map[string]interface{}); ok {
			if apiKey, ok := auth["apiKey"].(*APIKey); ok {
				apiKey.Key = key
				return nil
			}
		}
	}

	return ReportError("no API key authentication configured!")
}

// Helper method to set API key prefix for the first API key authentication
func (a *ApiClient) SetApiKeyPrefix(apiKeyPrefix string) error {
	for _, value := range a.authentication {
		if auth, ok := value.(map[string]interface{}); ok {
			if apiKey, ok := auth["apiKey"].(*APIKey); ok {
				apiKey.Prefix = apiKeyPrefix
				return nil
			}
		}
	}

	return ReportError("no API key authentication configured!")
}

// Helper method to set access token for the first OAuth2 authentication.
func (a *ApiClient) SetAccessToken(accessToken string) error {
	for _, value := range a.authentication {
		if  auth, ok := value.(*OAuth); ok {
			auth.AccessToken = accessToken
			return nil
		}
	}
	return ReportError("no OAuth2 authentication configured!")
}

/**
  Helper method to set maximum retry attempts.
  After the initial instantiation of ApiClient, maximum retry attempts must be modified only via this method
*/
func (a *ApiClient) SetMaxRetryAttempts(maxRetryAttempts int) {
	a.MaxRetryAttempts = maxRetryAttempts
}

/**
  Helper method to set httpclient timeout.
  After the initial instantiation of ApiClient, timeout must be modified only via this method
*/
func (a *ApiClient) SetTimeout(ms int) {
	dur := time.Duration(ms) * time.Millisecond
	setTimeout(dur, a)
}

func setTimeout(dur time.Duration, apiClient *ApiClient) {
	apiClient.Timeout = dur
	if dur <= 0 {
		log.Printf("Timeout cannot be 0 or less. Setting 30 second as default timeout.")
		apiClient.Timeout = 30 * time.Second
	} else if dur > (30 * time.Minute) {
		log.Printf("Timeout cannot be greater than 30 minutes. Setting 30 minutes as default timeout.")
		apiClient.Timeout = 30 * time.Minute
	}

	apiClient.httpClient.Timeout = apiClient.Timeout
}

/**
  Helper method to enable/disable SSL verification. By default, SSL verification is enabled.
  Please note that disabling SSL verification is not recommended and should only be done for test purposes.
*/
func (a *ApiClient) SetVerifySSL(verifySSL bool) {
	a.VerifySSL = verifySSL
}

/**
  Helper method to set retry back off period.
  After the initial instantiation of ApiClient, back off period must be modified only via this method
*/
func (a *ApiClient) SetRetryIntervalInMilliSeconds(ms int) {
	a.RetryInterval = time.Duration(ms) * time.Millisecond
}

func (a *ApiClient) setupClient() {
	var isRetryClientModified = false
	retryClientValue := reflect.ValueOf(a.retryClient)
	if !retryClientValue.IsValid() || retryClientValue.IsNil() {
		a.retryClient = retryablehttp.NewClient()
		isRetryClientModified = true
	}

	var transport = a.retryClient.HTTPClient.Transport.(*http.Transport)
	if isRetryClientModified || transport.TLSClientConfig == nil || transport.TLSClientConfig.InsecureSkipVerify != !a.VerifySSL {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !a.VerifySSL},
		}
		a.retryClient.HTTPClient.Transport = transport
		isRetryClientModified = true
	}

	if a.retryClient.RetryMax != a.MaxRetryAttempts ||
		a.retryClient.RetryWaitMax != a.RetryInterval ||
		a.retryClient.CheckRetry == nil {
		isRetryClientModified = true
	}

	a.retryClient.RetryMax = a.MaxRetryAttempts
	a.retryClient.RetryWaitMax = a.RetryInterval
	a.retryClient.CheckRetry = retryPolicy

	if a.LoggerFile != "" {
		log := logrus.New()
		leveledLogrus := LeveledLogrus{log}
		leveledLogrus.setLoggerFilePath(a.LoggerFile)
		if a.Debug {
			log.SetLevel(logrus.DebugLevel)
		} else {
			log.SetLevel(logrus.WarnLevel)
		}

		a.retryClient.Logger = retryablehttp.LeveledLogger(&leveledLogrus)
	} else {
		a.retryClient.Logger = nil
	}


	if isRetryClientModified {
		a.httpClient = a.retryClient.StandardClient()
		setTimeout(a.Timeout, a)
	}
}

// SelectHeaderContentType select a content type from the available list.
func (a *ApiClient) selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func (a *ApiClient) selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// prepareRequest build the request
func (a *ApiClient) prepareRequest(
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	authNames []string) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	postBodyValue := reflect.ValueOf(postBody)
	if postBodyValue.IsValid() && !postBodyValue.IsNil() {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}
		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", userAgent)

	// Authentication
	a.SetUserName(a.Username)
	a.SetPassword(a.Password)
	for _, authName := range authNames {
		// Basic HTTP authentication
		if  auth, ok := a.authentication[authName].(*BasicAuth); ok {
			if auth.UserName != "" && auth.Password != "" {
				localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
			}
		// API Key authentication
		} else if auth, ok := a.authentication[authName].(map[string]interface{}); ok {
			var key string
			if apiKey,ok := auth["apiKey"].(*APIKey); ok && apiKey.Prefix != "" {
				key = apiKey.Prefix + " " + apiKey.Key
			} else {
				key = apiKey.Key
			}

			if auth["in"] == "header" {
				localVarRequest.Header.Add(auth["name"].(string), key)
			}
			if auth["in"] == "query" {
				queries := localVarRequest.URL.Query()
				queries.Add(auth["name"].(string), key)
				localVarRequest.URL.RawQuery = queries.Encode()
			}
		// OAuth or Bearer authentication
		} else if auth, ok := a.authentication[authName].(*OAuth); ok {
			localVarRequest.Header.Add("Authorization", "Bearer " + auth.AccessToken)
		} else {
			return nil, ReportError("unknown authentication type: %s", authName)
		}
	}

	for header, value := range a.defaultHeaders {
		localVarRequest.Header.Add(header, value)
	}

	// Add the cookie to the request.
	if len(a.cookie) > 0 {
		localVarRequest.Header.Add("Cookie", a.cookie)
		delete(localVarRequest.Header, "Authorization")
	}

	return localVarRequest, nil
}


// RetryPolicy provides a callback for Client.CheckRetry, specifies retry on
// error codes mentioned in RetryStatusList
func retryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	for _, status := range retryStatusList {
		if resp.StatusCode == status {
			return true, nil
		}
	}
	return false, nil
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// addEtagReferenceToHeader method is used to read ETag and add it to If-Match header
func addEtagReferenceToHeader(body interface{}, headerParams map[string]string) {
	if reflect.ValueOf(body).Elem().Kind() == reflect.Struct {
		if reserved := reflect.ValueOf(body).Elem().FieldByName("Reserved_"); reserved.IsValid() {
			reservedMap := reserved.Interface().(map[string]interface{})
			if etag, etagOk := reservedMap[eTag].(string); etagOk {
				headerParams[ifMatch] = etag
			}
		}
	}
}

/*GetEtag
Get ETag from an object if exists, otherwise returns empty string.
The ETag is usually provided in the response of the GET API calls, which can further be used in other HTTP operations.
*/
func (a *ApiClient) GetEtag(object interface{}) string {
	var reserved reflect.Value
	if reflect.TypeOf(object).Kind() == reflect.Struct {
		reserved = reflect.ValueOf(object).FieldByName("Reserved_")
	} else if reflect.TypeOf(object).Kind() == reflect.Interface || reflect.TypeOf(object).Kind() == reflect.Ptr {
		reserved = reflect.ValueOf(object).Elem().FieldByName("Reserved_")
	} else {
		log.Printf("\nUnrecognized input type %s for %s to retrieve etag!\n", reflect.TypeOf(object).Kind(), object)
		return ""
	}

	if reserved.IsValid() {
		etagKey := strings.ToLower(eTag)
		reservedMap := reserved.Interface().(map[string]interface{})
		for k, v := range reservedMap {
			if strings.ToLower(k) == etagKey {
				return v.(string)
			}
		}
	}

	return ""
}

// addEtagReferenceToResponse method is used to read ETag and add it to response
func addEtagReferenceToResponse(headers http.Header, body []byte) []byte {
	if etag := headers.Get(eTag); etag != "" {
		responseMap := map[string]interface{}{}
		json.Unmarshal(body, &responseMap)
		if r, ok := responseMap["$reserved"].(map[string]interface{}); ok {
			r[eTag] = etag
			if d, ok := responseMap["data"].(map[string]interface{}); ok {
				if r2, ok := d["$reserved"].(map[string]interface{}); ok {
					r2[eTag] = etag
					m, _ := json.Marshal(responseMap)
					return m
				}
			} else if dList, ok := responseMap["data"].([]interface{}); ok {
				for _, d := range dList {
					if d, ok := d.(map[string]interface{}); ok {
						if r3, ok := d["$reserved"].(map[string]interface{}); ok {
							r3[eTag] = etag
						}
					}
				}
				m, _ := json.Marshal(responseMap)
				return m
			}
		}
	}
	return body
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if nil == bodyBuf {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// Set Cookie information to reuse in subsequent requests for a valid response
func (a *ApiClient) updateCookies(response *http.Response) {
	if a.refreshCookie {
		cookiesList := response.Header["Set-Cookie"]
		if len(cookiesList) > 0 {
			cookieFromResponse := ""
			for _, value := range cookiesList {
				finalCookie := strings.SplitN(value, ";", 2)[0]
				if strings.Contains(finalCookie, "=") {
					finalCookie = strings.TrimSpace(finalCookie)
				} else {
					finalCookie = ""
				}

				if finalCookie != "" {
					cookieFromResponse += finalCookie + ";"
				}
			}

			// Remove trailing ";"
			if cookieFromResponse != "" {
				cookieFromResponse = strings.TrimSuffix(cookieFromResponse, ";")
			}

			a.cookie = cookieFromResponse
			a.refreshCookie = false
		}
	}
}

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// OAuth provides OAuth authentication
type OAuth struct {
	AccessToken string
}

// GenericOpenAPIError Provides access to the body (error), status and model on returned errors.
type GenericOpenAPIError struct {
	Body   []byte
	Model  interface{}
	Status string
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return string(e.Body)
}

// Error returns deserialized response body if compatible with GenericOpenAPIError.Model
func (e GenericOpenAPIError) DeserializedModel() interface{} {
	err := json.Unmarshal(e.Body, e.Model)
	if err != nil {
		return nil
	}
	return e.Model
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func ParameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func ParameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// Prevent trying to import "fmt"
func ReportError(format string, b ...interface{}) error {
	return fmt.Errorf(format, b...)
}

type LeveledLogrus struct {
	*logrus.Logger
}

func (l *LeveledLogrus) Error(msg string, keysAndValues ...interface{}) {
	l.WithFields(fields(keysAndValues)).Error(msg)
}

func (l *LeveledLogrus) Info(msg string, keysAndValues ...interface{}) {
	l.WithFields(fields(keysAndValues)).Info(msg)
}
func (l *LeveledLogrus) Debug(msg string, keysAndValues ...interface{}) {
	l.WithFields(fields(keysAndValues)).Debug(msg)
}

func (l *LeveledLogrus) Warn(msg string, keysAndValues ...interface{}) {
	l.WithFields(fields(keysAndValues)).Warn(msg)
}

func fields(keysAndValues []interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	for i := 0; i < len(keysAndValues)-1; i += 2 {
		fields[keysAndValues[i].(string)] = keysAndValues[i+1]
	}

	return fields
}

func (l *LeveledLogrus) setLoggerFilePath(filename string) error {
	logFile, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		l.Error("Error opening log file", "error", err)
		return err
	}

	l.SetOutput(logFile)
	l.SetReportCaller(true)
	return nil
}
