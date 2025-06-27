package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import5 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	"net/http"
	"net/url"
	"strings"
)

type OvasApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewOvasApi(apiClient *client.ApiClient) *OvasApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &OvasApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an OVA using the provided request body. The name and source are mandatory fields to create an OVA.
func (api *OvasApi) CreateOva(body *import5.Ova, args ...map[string]interface{}) (*import5.CreateOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas"

	// verify the required parameter 'body' is set
	if nil == body {
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.CreateOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the OVA based on the given external identifier.
func (api *OvasApi) DeleteOvaById(extId *string, args ...map[string]interface{}) (*import5.DeleteOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import5.DeleteOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deploys a VM from an OVA, allowing you to override the VM configuration if needed.
func (api *OvasApi) DeployOva(extId *string, body *import5.OvaDeploymentSpec, args ...map[string]interface{}) (*import5.DeployOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas/{extId}/$actions/deploy"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.DeployOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Downloads an OVA based on the given external identifier. This is a stream download of the OVA file.
func (api *OvasApi) GetFileByOvaId(ovaExtId *string, args ...map[string]interface{}) (*import5.GetOvaFileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas/{ovaExtId}/file"

	// verify the required parameter 'ovaExtId' is set
	if nil == ovaExtId {
		return nil, client.ReportError("ovaExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"ovaExtId"+"}", url.PathEscape(client.ParameterToString(*ovaExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/octet-stream", "application/json"}

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

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import5.NewGetOvaFileApiResponse()
			fileDetail := import5.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import6.Flag
			flags = append(flags, import6.Flag{Name: &flagName, Value: &flagValue})
			metadata := import7.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import5.GetOvaFileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the OVA details for the provided external identifier.
func (api *OvasApi) GetOvaById(extId *string, args ...map[string]interface{}) (*import5.GetOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import5.GetOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This lists all accessible OVAs using the default pagination, which can be customized.
func (api *OvasApi) ListOvas(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import5.ListOvasApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas"

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

	unmarshalledResp := new(import5.ListOvasApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates an OVA with the given external identifier. It is always recommended to do a GET operation on a resource before a PUT operation to ensure the correct ETag is used.
func (api *OvasApi) UpdateOvaById(extId *string, body *import5.Ova, args ...map[string]interface{}) (*import5.UpdateOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/content/ovas/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import5.UpdateOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
