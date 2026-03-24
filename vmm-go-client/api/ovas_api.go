package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import10 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	import13 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/ovas"
	"net/http"
	"net/url"
	"strings"
)

type OvasApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *OvasServiceApi
}

type OvasServiceApi struct {
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

	a.ServiceClient = NewOvasServiceApi(a.ApiClient)

	return a
}

func NewOvasServiceApi(apiClient *client.ApiClient) *OvasServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &OvasServiceApi{
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
func (api *OvasApi) CreateOva(body *import9.Ova, args ...map[string]interface{}) (*import9.CreateOvaApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateOva(context.Background(), &import13.CreateOvaRequest{
		Body: body,
	}, args...)
}

// Creates an OVA using the provided request body. The name and source are mandatory fields to create an OVA.
func (api *OvasServiceApi) CreateOva(ctx context.Context, request *import13.CreateOvaRequest, args ...map[string]interface{}) (*import9.CreateOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas"

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

	unmarshalledResp := new(import9.CreateOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the OVA based on the given external identifier.
func (api *OvasApi) DeleteOvaById(extId *string, args ...map[string]interface{}) (*import9.DeleteOvaApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteOvaById(context.Background(), &import13.DeleteOvaByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the OVA based on the given external identifier.
func (api *OvasServiceApi) DeleteOvaById(ctx context.Context, request *import13.DeleteOvaByIdRequest, args ...map[string]interface{}) (*import9.DeleteOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas/{extId}"

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

	unmarshalledResp := new(import9.DeleteOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deploys a VM from an OVA, allowing you to override the VM configuration if needed.
func (api *OvasApi) DeployOva(extId *string, body *import9.OvaDeploymentSpec, args ...map[string]interface{}) (*import9.DeployOvaApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeployOva(context.Background(), &import13.DeployOvaRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Deploys a VM from an OVA, allowing you to override the VM configuration if needed.
func (api *OvasServiceApi) DeployOva(ctx context.Context, request *import13.DeployOvaRequest, args ...map[string]interface{}) (*import9.DeployOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas/{extId}/$actions/deploy"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import9.DeployOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Downloads an OVA based on the given external identifier. This is a stream download of the OVA file.
func (api *OvasApi) GetFileByOvaId(ovaExtId *string, args ...map[string]interface{}) (*import9.GetOvaFileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetFileByOvaId(context.Background(), &import13.GetFileByOvaIdRequest{
		OvaExtId: ovaExtId,
	}, args...)
}

// Downloads an OVA based on the given external identifier. This is a stream download of the OVA file.
func (api *OvasServiceApi) GetFileByOvaId(ctx context.Context, request *import13.GetFileByOvaIdRequest, args ...map[string]interface{}) (*import9.GetOvaFileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas/{ovaExtId}/file"

	// verify the required parameter 'ovaExtId' is set
	if nil == request.OvaExtId {
		return nil, client.ReportError("ovaExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"ovaExtId"+"}", url.PathEscape(client.ParameterToString(*request.OvaExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
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

			response := import9.NewGetOvaFileApiResponse()
			fileDetail := import9.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import10.Flag
			flags = append(flags, import10.Flag{Name: &flagName, Value: &flagValue})
			metadata := import11.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import9.GetOvaFileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the OVA details for the provided external identifier.
func (api *OvasApi) GetOvaById(extId *string, args ...map[string]interface{}) (*import9.GetOvaApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetOvaById(context.Background(), &import13.GetOvaByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the OVA details for the provided external identifier.
func (api *OvasServiceApi) GetOvaById(ctx context.Context, request *import13.GetOvaByIdRequest, args ...map[string]interface{}) (*import9.GetOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas/{extId}"

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

	unmarshalledResp := new(import9.GetOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This lists all accessible OVAs using the default pagination, which can be customized.
func (api *OvasApi) ListOvas(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import9.ListOvasApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListOvas(context.Background(), &import13.ListOvasRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// This lists all accessible OVAs using the default pagination, which can be customized.
func (api *OvasServiceApi) ListOvas(ctx context.Context, request *import13.ListOvasRequest, args ...map[string]interface{}) (*import9.ListOvasApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
	}
	if request.Filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*request.Filter_, ""))
	}
	if request.Orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*request.Orderby_, ""))
	}
	if request.Select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*request.Select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import9.ListOvasApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates an OVA with the given external identifier. It is always recommended to do a GET operation on a resource before a PUT operation to ensure the correct ETag is used.
func (api *OvasApi) UpdateOvaById(extId *string, body *import9.Ova, args ...map[string]interface{}) (*import9.UpdateOvaApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewOvasServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateOvaById(context.Background(), &import13.UpdateOvaByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates an OVA with the given external identifier. It is always recommended to do a GET operation on a resource before a PUT operation to ensure the correct ETag is used.
func (api *OvasServiceApi) UpdateOvaById(ctx context.Context, request *import13.UpdateOvaByIdRequest, args ...map[string]interface{}) (*import9.UpdateOvaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/ovas/{extId}"

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

	unmarshalledResp := new(import9.UpdateOvaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
