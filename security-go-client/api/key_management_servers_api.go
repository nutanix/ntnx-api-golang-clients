package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/keymanagementservers"
	"net/http"
	"net/url"
	"strings"
)

type KeyManagementServersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *KeyManagementServersServiceApi
}

type KeyManagementServersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewKeyManagementServersApi(apiClient *client.ApiClient) *KeyManagementServersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &KeyManagementServersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewKeyManagementServersServiceApi(a.ApiClient)

	return a
}

func NewKeyManagementServersServiceApi(apiClient *client.ApiClient) *KeyManagementServersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &KeyManagementServersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new key management server with the specified access credentials.
func (api *KeyManagementServersApi) CreateKeyManagementServer(body *import3.KeyManagementServer, args ...map[string]interface{}) (*import3.CreateKeyManagementServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewKeyManagementServersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateKeyManagementServer(context.Background(), &import5.CreateKeyManagementServerRequest{
		Body: body,
	}, args...)
}

// Creates a new key management server with the specified access credentials.
func (api *KeyManagementServersServiceApi) CreateKeyManagementServer(ctx context.Context, request *import5.CreateKeyManagementServerRequest, args ...map[string]interface{}) (*import3.CreateKeyManagementServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/key-management-servers"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.CreateKeyManagementServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a key management server identified by its unique identifier.
func (api *KeyManagementServersApi) DeleteKeyManagementServerById(extId *string, args ...map[string]interface{}) (*import3.DeleteKeyManagementServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewKeyManagementServersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteKeyManagementServerById(context.Background(), &import5.DeleteKeyManagementServerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a key management server identified by its unique identifier.
func (api *KeyManagementServersServiceApi) DeleteKeyManagementServerById(ctx context.Context, request *import5.DeleteKeyManagementServerByIdRequest, args ...map[string]interface{}) (*import3.DeleteKeyManagementServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/key-management-servers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.DeleteKeyManagementServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the details of a key management server identified by its unique identifier.
func (api *KeyManagementServersApi) GetKeyManagementServerById(extId *string, args ...map[string]interface{}) (*import3.GetKeyManagementServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewKeyManagementServersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetKeyManagementServerById(context.Background(), &import5.GetKeyManagementServerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the details of a key management server identified by its unique identifier.
func (api *KeyManagementServersServiceApi) GetKeyManagementServerById(ctx context.Context, request *import5.GetKeyManagementServerByIdRequest, args ...map[string]interface{}) (*import3.GetKeyManagementServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/key-management-servers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.GetKeyManagementServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Provide a comprehensive list of all key management servers, including their access details and relevant attributes.
func (api *KeyManagementServersApi) ListKeyManagementServers(args ...map[string]interface{}) (*import3.ListKeyManagementServersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewKeyManagementServersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListKeyManagementServers(context.Background(), &import5.ListKeyManagementServersRequest{}, args...)
}

// Provide a comprehensive list of all key management servers, including their access details and relevant attributes.
func (api *KeyManagementServersServiceApi) ListKeyManagementServers(ctx context.Context, request *import5.ListKeyManagementServersRequest, args ...map[string]interface{}) (*import3.ListKeyManagementServersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/key-management-servers"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.ListKeyManagementServersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a key management server with a new set of access credentials.
func (api *KeyManagementServersApi) UpdateKeyManagementServerById(extId *string, body *import3.KeyManagementServer, args ...map[string]interface{}) (*import3.UpdateKeyManagementServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewKeyManagementServersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateKeyManagementServerById(context.Background(), &import5.UpdateKeyManagementServerByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a key management server with a new set of access credentials.
func (api *KeyManagementServersServiceApi) UpdateKeyManagementServerById(ctx context.Context, request *import5.UpdateKeyManagementServerByIdRequest, args ...map[string]interface{}) (*import3.UpdateKeyManagementServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/key-management-servers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UpdateKeyManagementServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
