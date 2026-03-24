package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/request/directoryserverconfigs"
	"net/http"
	"net/url"
	"strings"
)

type DirectoryServerConfigsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DirectoryServerConfigsServiceApi
}

type DirectoryServerConfigsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDirectoryServerConfigsApi(apiClient *client.ApiClient) *DirectoryServerConfigsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DirectoryServerConfigsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDirectoryServerConfigsServiceApi(a.ApiClient)

	return a
}

func NewDirectoryServerConfigsServiceApi(apiClient *client.ApiClient) *DirectoryServerConfigsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DirectoryServerConfigsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates the mapping between a group in Active Directory and the Category.
func (api *DirectoryServerConfigsApi) CreateCategoryMapping(body *import1.CategoryMapping, args ...map[string]interface{}) (*import1.CreateDsCategoryMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCategoryMapping(context.Background(), &import3.CreateCategoryMappingRequest{
		Body: body,
	}, args...)
}

// Creates the mapping between a group in Active Directory and the Category.
func (api *DirectoryServerConfigsServiceApi) CreateCategoryMapping(ctx context.Context, request *import3.CreateCategoryMappingRequest, args ...map[string]interface{}) (*import1.CreateDsCategoryMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/category-mappings"

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

	unmarshalledResp := new(import1.CreateDsCategoryMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Configures various aspects of identity categorization.
func (api *DirectoryServerConfigsApi) CreateDirectoryServerConfig(body *import1.DirectoryServerConfig, args ...map[string]interface{}) (*import1.CreateDirectoryServerConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDirectoryServerConfig(context.Background(), &import3.CreateDirectoryServerConfigRequest{
		Body: body,
	}, args...)
}

// Configures various aspects of identity categorization.
func (api *DirectoryServerConfigsServiceApi) CreateDirectoryServerConfig(ctx context.Context, request *import3.CreateDirectoryServerConfigRequest, args ...map[string]interface{}) (*import1.CreateDirectoryServerConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/directory-server-configs"

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

	unmarshalledResp := new(import1.CreateDirectoryServerConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the Directory Server with the provided ExtID.
func (api *DirectoryServerConfigsApi) DeleteDirectoryServerConfigById(extId *string, args ...map[string]interface{}) (*import1.DeleteDirectoryServerConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDirectoryServerConfigById(context.Background(), &import3.DeleteDirectoryServerConfigByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the Directory Server with the provided ExtID.
func (api *DirectoryServerConfigsServiceApi) DeleteDirectoryServerConfigById(ctx context.Context, request *import3.DeleteDirectoryServerConfigByIdRequest, args ...map[string]interface{}) (*import1.DeleteDirectoryServerConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/directory-server-configs/{extId}"

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

	unmarshalledResp := new(import1.DeleteDirectoryServerConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the directory configuration with the provided ExtID.
func (api *DirectoryServerConfigsApi) DeleteDsCategoryMappingById(extId *string, args ...map[string]interface{}) (*import1.DeleteDsCategoryMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDsCategoryMappingById(context.Background(), &import3.DeleteDsCategoryMappingByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the directory configuration with the provided ExtID.
func (api *DirectoryServerConfigsServiceApi) DeleteDsCategoryMappingById(ctx context.Context, request *import3.DeleteDsCategoryMappingByIdRequest, args ...map[string]interface{}) (*import1.DeleteDsCategoryMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/category-mappings/{extId}"

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

	unmarshalledResp := new(import1.DeleteDsCategoryMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the list of Directory Server configurations.
func (api *DirectoryServerConfigsApi) GetDirectoryServerConfigById(extId *string, args ...map[string]interface{}) (*import1.GetDirectoryServerConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDirectoryServerConfigById(context.Background(), &import3.GetDirectoryServerConfigByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets the list of Directory Server configurations.
func (api *DirectoryServerConfigsServiceApi) GetDirectoryServerConfigById(ctx context.Context, request *import3.GetDirectoryServerConfigByIdRequest, args ...map[string]interface{}) (*import1.GetDirectoryServerConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/directory-server-configs/{extId}"

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

	unmarshalledResp := new(import1.GetDirectoryServerConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the category to directory configuration information with the provided ExtID.
func (api *DirectoryServerConfigsApi) GetDsCategoryMappingById(extId *string, args ...map[string]interface{}) (*import1.GetDsCategoryMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDsCategoryMappingById(context.Background(), &import3.GetDsCategoryMappingByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets the category to directory configuration information with the provided ExtID.
func (api *DirectoryServerConfigsServiceApi) GetDsCategoryMappingById(ctx context.Context, request *import3.GetDsCategoryMappingByIdRequest, args ...map[string]interface{}) (*import1.GetDsCategoryMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/category-mappings/{extId}"

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

	unmarshalledResp := new(import1.GetDsCategoryMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the list of Directory Server Category Mappings.
func (api *DirectoryServerConfigsApi) ListCategoryMappings(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDsCategoryMappingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCategoryMappings(context.Background(), &import3.ListCategoryMappingsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Gets the list of Directory Server Category Mappings.
func (api *DirectoryServerConfigsServiceApi) ListCategoryMappings(ctx context.Context, request *import3.ListCategoryMappingsRequest, args ...map[string]interface{}) (*import1.ListDsCategoryMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/category-mappings"

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

	unmarshalledResp := new(import1.ListDsCategoryMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the list of Directory Servers.
func (api *DirectoryServerConfigsApi) ListDirectoryServerConfigs(select_ *string, args ...map[string]interface{}) (*import1.ListDirectoryServerConfigsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDirectoryServerConfigs(context.Background(), &import3.ListDirectoryServerConfigsRequest{
		Select_: select_,
	}, args...)
}

// Gets the list of Directory Servers.
func (api *DirectoryServerConfigsServiceApi) ListDirectoryServerConfigs(ctx context.Context, request *import3.ListDirectoryServerConfigsRequest, args ...map[string]interface{}) (*import1.ListDirectoryServerConfigsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/directory-server-configs"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
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

	unmarshalledResp := new(import1.ListDirectoryServerConfigsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Directory Server Config with the provided ExtID.
func (api *DirectoryServerConfigsApi) UpdateDirectoryServerConfigById(extId *string, body *import1.DirectoryServerConfig, args ...map[string]interface{}) (*import1.UpdateDirectoryServerConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDirectoryServerConfigById(context.Background(), &import3.UpdateDirectoryServerConfigByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Directory Server Config with the provided ExtID.
func (api *DirectoryServerConfigsServiceApi) UpdateDirectoryServerConfigById(ctx context.Context, request *import3.UpdateDirectoryServerConfigByIdRequest, args ...map[string]interface{}) (*import1.UpdateDirectoryServerConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/directory-server-configs/{extId}"

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

	unmarshalledResp := new(import1.UpdateDirectoryServerConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the category to directory configuration mapping information with the provided ExtID.
func (api *DirectoryServerConfigsApi) UpdateDsCategoryMappingById(extId *string, body *import1.CategoryMapping, args ...map[string]interface{}) (*import1.UpdateDsCategoryMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServerConfigsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDsCategoryMappingById(context.Background(), &import3.UpdateDsCategoryMappingByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the category to directory configuration mapping information with the provided ExtID.
func (api *DirectoryServerConfigsServiceApi) UpdateDsCategoryMappingById(ctx context.Context, request *import3.UpdateDsCategoryMappingByIdRequest, args ...map[string]interface{}) (*import1.UpdateDsCategoryMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/category-mappings/{extId}"

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

	unmarshalledResp := new(import1.UpdateDsCategoryMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
