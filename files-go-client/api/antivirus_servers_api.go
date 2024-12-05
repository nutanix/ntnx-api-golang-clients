package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type AntivirusServersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAntivirusServersApi(apiClient *client.ApiClient) *AntivirusServersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AntivirusServersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a new antivirus server. Please refer model section to get more details about fields.
func (api *AntivirusServersApi) CreateAntivirusServer(fileServerExtId *string, body *import3.AntivirusServer, args ...map[string]interface{}) (*import3.CreateAntivirusServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.CreateAntivirusServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the antivirus server with the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and  of the antivirus server (`extId`).
func (api *AntivirusServersApi) DeleteAntivirusServerById(fileServerExtId *string, extId *string, args ...map[string]interface{}) (*import3.DeleteAntivirusServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.DeleteAntivirusServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get antivirus server for the provided external identifier.  Specify a valid identifier of the file server (`fileServerExtId`)  and  of the antivirus server (`extId`).
func (api *AntivirusServersApi) GetAntivirusServerById(fileServerExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetAntivirusServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.GetAntivirusServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a paginated list of antivirus servers.
func (api *AntivirusServersApi) ListAntivirusServers(fileServerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListAntivirusServersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.ListAntivirusServersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Test antivirus server connection with the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and  of the antivirus server(`extId`) for which the connection needs to be tested.
func (api *AntivirusServersApi) TestConnectionAntivirusServer(fileServerExtId *string, extId *string, args ...map[string]interface{}) (*import3.TestConnectionAntivirusServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId}/$actions/test-connection"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.TestConnectionAntivirusServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update antivirus server by using the provided request body.  Specify a valid identifier (`fileServerExtId`) of the file server and (`extId`) of the antivirus server to be updated.  It is always recommended to perform a GET request on a resource before requesting an UPDATE call.
func (api *AntivirusServersApi) UpdateAntivirusServerById(fileServerExtId *string, extId *string, body *import3.AntivirusServer, args ...map[string]interface{}) (*import3.UpdateAntivirusServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/anti-virus-servers/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UpdateAntivirusServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
