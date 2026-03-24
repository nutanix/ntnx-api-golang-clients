package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/request/servicegroups"
	"net/http"
	"net/url"
	"strings"
)

type ServiceGroupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ServiceGroupsServiceApi
}

type ServiceGroupsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewServiceGroupsApi(apiClient *client.ApiClient) *ServiceGroupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ServiceGroupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewServiceGroupsServiceApi(a.ApiClient)

	return a
}

func NewServiceGroupsServiceApi(apiClient *client.ApiClient) *ServiceGroupsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ServiceGroupsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a Service Group.
func (api *ServiceGroupsApi) CreateServiceGroup(body *import1.ServiceGroup, args ...map[string]interface{}) (*import1.CreateServiceGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewServiceGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateServiceGroup(context.Background(), &import8.CreateServiceGroupRequest{
		Body: body,
	}, args...)
}

// Creates a Service Group.
func (api *ServiceGroupsServiceApi) CreateServiceGroup(ctx context.Context, request *import8.CreateServiceGroupRequest, args ...map[string]interface{}) (*import1.CreateServiceGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/service-groups"

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

	unmarshalledResp := new(import1.CreateServiceGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the Service Group with the provided ExtID.
func (api *ServiceGroupsApi) DeleteServiceGroupById(extId *string, args ...map[string]interface{}) (*import1.DeleteServiceGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewServiceGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteServiceGroupById(context.Background(), &import8.DeleteServiceGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the Service Group with the provided ExtID.
func (api *ServiceGroupsServiceApi) DeleteServiceGroupById(ctx context.Context, request *import8.DeleteServiceGroupByIdRequest, args ...map[string]interface{}) (*import1.DeleteServiceGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/service-groups/{extId}"

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

	unmarshalledResp := new(import1.DeleteServiceGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the Service Group with the provided ExtID.
func (api *ServiceGroupsApi) GetServiceGroupById(extId *string, args ...map[string]interface{}) (*import1.GetServiceGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewServiceGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetServiceGroupById(context.Background(), &import8.GetServiceGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets the Service Group with the provided ExtID.
func (api *ServiceGroupsServiceApi) GetServiceGroupById(ctx context.Context, request *import8.GetServiceGroupByIdRequest, args ...map[string]interface{}) (*import1.GetServiceGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/service-groups/{extId}"

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

	unmarshalledResp := new(import1.GetServiceGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets a list of Service Groups.
func (api *ServiceGroupsApi) ListServiceGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListServiceGroupsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewServiceGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListServiceGroups(context.Background(), &import8.ListServiceGroupsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Gets a list of Service Groups.
func (api *ServiceGroupsServiceApi) ListServiceGroups(ctx context.Context, request *import8.ListServiceGroupsRequest, args ...map[string]interface{}) (*import1.ListServiceGroupsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/service-groups"

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

	unmarshalledResp := new(import1.ListServiceGroupsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Service Group with the provided ExtID.
func (api *ServiceGroupsApi) UpdateServiceGroupById(extId *string, body *import1.ServiceGroup, args ...map[string]interface{}) (*import1.UpdateServiceGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewServiceGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateServiceGroupById(context.Background(), &import8.UpdateServiceGroupByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Service Group with the provided ExtID.
func (api *ServiceGroupsServiceApi) UpdateServiceGroupById(ctx context.Context, request *import8.UpdateServiceGroupByIdRequest, args ...map[string]interface{}) (*import1.UpdateServiceGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/service-groups/{extId}"

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

	unmarshalledResp := new(import1.UpdateServiceGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
