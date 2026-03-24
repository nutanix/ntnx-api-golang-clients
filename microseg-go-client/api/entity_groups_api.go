package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/request/entitygroups"
	"net/http"
	"net/url"
	"strings"
)

type EntityGroupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EntityGroupsServiceApi
}

type EntityGroupsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEntityGroupsApi(apiClient *client.ApiClient) *EntityGroupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntityGroupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEntityGroupsServiceApi(a.ApiClient)

	return a
}

func NewEntityGroupsServiceApi(apiClient *client.ApiClient) *EntityGroupsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntityGroupsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an Entity Group based on the details provided by the user.
func (api *EntityGroupsApi) CreateEntityGroup(body *import1.EntityGroup, args ...map[string]interface{}) (*import1.CreateEntityGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntityGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateEntityGroup(context.Background(), &import4.CreateEntityGroupRequest{
		Body: body,
	}, args...)
}

// Creates an Entity Group based on the details provided by the user.
func (api *EntityGroupsServiceApi) CreateEntityGroup(ctx context.Context, request *import4.CreateEntityGroupRequest, args ...map[string]interface{}) (*import1.CreateEntityGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/entity-groups"

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

	unmarshalledResp := new(import1.CreateEntityGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes an Entity Group based on the provided external identifier.
func (api *EntityGroupsApi) DeleteEntityGroupById(extId *string, args ...map[string]interface{}) (*import1.DeleteEntityGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntityGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteEntityGroupById(context.Background(), &import4.DeleteEntityGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes an Entity Group based on the provided external identifier.
func (api *EntityGroupsServiceApi) DeleteEntityGroupById(ctx context.Context, request *import4.DeleteEntityGroupByIdRequest, args ...map[string]interface{}) (*import1.DeleteEntityGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/entity-groups/{extId}"

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

	unmarshalledResp := new(import1.DeleteEntityGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retreives an Entity Group based on the provided external identifier.
func (api *EntityGroupsApi) GetEntityGroupById(extId *string, args ...map[string]interface{}) (*import1.GetEntityGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntityGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityGroupById(context.Background(), &import4.GetEntityGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retreives an Entity Group based on the provided external identifier.
func (api *EntityGroupsServiceApi) GetEntityGroupById(ctx context.Context, request *import4.GetEntityGroupByIdRequest, args ...map[string]interface{}) (*import1.GetEntityGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/entity-groups/{extId}"

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

	unmarshalledResp := new(import1.GetEntityGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retreives a list of Entity Groups.
func (api *EntityGroupsApi) ListEntityGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListEntityGroupsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntityGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEntityGroups(context.Background(), &import4.ListEntityGroupsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retreives a list of Entity Groups.
func (api *EntityGroupsServiceApi) ListEntityGroups(ctx context.Context, request *import4.ListEntityGroupsRequest, args ...map[string]interface{}) (*import1.ListEntityGroupsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/entity-groups"

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

	unmarshalledResp := new(import1.ListEntityGroupsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates an Entity Group based on the provided external identifier.
func (api *EntityGroupsApi) UpdateEntityGroupById(extId *string, body *import1.EntityGroup, args ...map[string]interface{}) (*import1.UpdateEntityGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntityGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateEntityGroupById(context.Background(), &import4.UpdateEntityGroupByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates an Entity Group based on the provided external identifier.
func (api *EntityGroupsServiceApi) UpdateEntityGroupById(ctx context.Context, request *import4.UpdateEntityGroupByIdRequest, args ...map[string]interface{}) (*import1.UpdateEntityGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/entity-groups/{extId}"

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

	unmarshalledResp := new(import1.UpdateEntityGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
