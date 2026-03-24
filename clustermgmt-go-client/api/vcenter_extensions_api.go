package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import15 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/vcenterextensions"
	"net/http"
	"net/url"
	"strings"
)

type VcenterExtensionsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VcenterExtensionsServiceApi
}

type VcenterExtensionsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVcenterExtensionsApi(apiClient *client.ApiClient) *VcenterExtensionsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VcenterExtensionsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVcenterExtensionsServiceApi(a.ApiClient)

	return a
}

func NewVcenterExtensionsServiceApi(apiClient *client.ApiClient) *VcenterExtensionsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VcenterExtensionsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches vCenter server extension information for the given VcenterExtensionExtId.
func (api *VcenterExtensionsApi) GetVcenterExtensionById(extId *string, args ...map[string]interface{}) (*import1.GetVcenterExtensionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVcenterExtensionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVcenterExtensionById(context.Background(), &import15.GetVcenterExtensionByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches vCenter server extension information for the given VcenterExtensionExtId.
func (api *VcenterExtensionsServiceApi) GetVcenterExtensionById(ctx context.Context, request *import15.GetVcenterExtensionByIdRequest, args ...map[string]interface{}) (*import1.GetVcenterExtensionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/vcenter-extensions/{extId}"

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

	unmarshalledResp := new(import1.GetVcenterExtensionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List of vCenter server extensions for the clusters registered under a Prism Central server.
func (api *VcenterExtensionsApi) ListVcenterExtensions(page_ *int, limit_ *int, filter_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListVcenterExtensionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVcenterExtensionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVcenterExtensions(context.Background(), &import15.ListVcenterExtensionsRequest{
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
		Select_: select_,
	}, args...)
}

// List of vCenter server extensions for the clusters registered under a Prism Central server.
func (api *VcenterExtensionsServiceApi) ListVcenterExtensions(ctx context.Context, request *import15.ListVcenterExtensionsRequest, args ...map[string]interface{}) (*import1.ListVcenterExtensionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/vcenter-extensions"

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

	unmarshalledResp := new(import1.ListVcenterExtensionsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Registers Nutanix cluster as a vCenter server extension.
func (api *VcenterExtensionsApi) RegisterVcenterExtension(extId *string, body *import1.VcenterCredentials, args ...map[string]interface{}) (*import1.RegisterVcenterExtensionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVcenterExtensionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RegisterVcenterExtension(context.Background(), &import15.RegisterVcenterExtensionRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Registers Nutanix cluster as a vCenter server extension.
func (api *VcenterExtensionsServiceApi) RegisterVcenterExtension(ctx context.Context, request *import15.RegisterVcenterExtensionRequest, args ...map[string]interface{}) (*import1.RegisterVcenterExtensionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/vcenter-extensions/{extId}/$actions/register"

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

	unmarshalledResp := new(import1.RegisterVcenterExtensionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Unregisters Nutanix cluster as a vCenter server extension.
func (api *VcenterExtensionsApi) UnregisterVcenterExtension(extId *string, body *import1.VcenterCredentials, args ...map[string]interface{}) (*import1.UnregisterVcenterExtensionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVcenterExtensionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnregisterVcenterExtension(context.Background(), &import15.UnregisterVcenterExtensionRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Unregisters Nutanix cluster as a vCenter server extension.
func (api *VcenterExtensionsServiceApi) UnregisterVcenterExtension(ctx context.Context, request *import15.UnregisterVcenterExtensionRequest, args ...map[string]interface{}) (*import1.UnregisterVcenterExtensionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/vcenter-extensions/{extId}/$actions/unregister"

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

	unmarshalledResp := new(import1.UnregisterVcenterExtensionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
