package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import25 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/routingpolicies"
	"net/http"
	"net/url"
	"strings"
)

type RoutingPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RoutingPoliciesServiceApi
}

type RoutingPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRoutingPoliciesApi(apiClient *client.ApiClient) *RoutingPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutingPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRoutingPoliciesServiceApi(a.ApiClient)

	return a
}

func NewRoutingPoliciesServiceApi(apiClient *client.ApiClient) *RoutingPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutingPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a routing policy.
func (api *RoutingPoliciesApi) CreateRoutingPolicy(body *import4.RoutingPolicy, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRoutingPolicy(context.Background(), &import25.CreateRoutingPolicyRequest{
		Body: body,
	}, args...)
}

// Create a routing policy.
func (api *RoutingPoliciesServiceApi) CreateRoutingPolicy(ctx context.Context, request *import25.CreateRoutingPolicyRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/routing-policies"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the routing policy corresponding to the extId.
func (api *RoutingPoliciesApi) DeleteRoutingPolicyById(extId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRoutingPolicyById(context.Background(), &import25.DeleteRoutingPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the routing policy corresponding to the extId.
func (api *RoutingPoliciesServiceApi) DeleteRoutingPolicyById(ctx context.Context, request *import25.DeleteRoutingPolicyByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/routing-policies/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a single routing policy corresponding to the extId.
func (api *RoutingPoliciesApi) GetRoutingPolicyById(extId *string, args ...map[string]interface{}) (*import4.GetRoutingPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRoutingPolicyById(context.Background(), &import25.GetRoutingPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get a single routing policy corresponding to the extId.
func (api *RoutingPoliciesServiceApi) GetRoutingPolicyById(ctx context.Context, request *import25.GetRoutingPolicyByIdRequest, args ...map[string]interface{}) (*import4.GetRoutingPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/routing-policies/{extId}"

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

	unmarshalledResp := new(import4.GetRoutingPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a list of routing policies.
func (api *RoutingPoliciesApi) ListRoutingPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListRoutingPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRoutingPolicies(context.Background(), &import25.ListRoutingPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Get a list of routing policies.
func (api *RoutingPoliciesServiceApi) ListRoutingPolicies(ctx context.Context, request *import25.ListRoutingPoliciesRequest, args ...map[string]interface{}) (*import4.ListRoutingPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/routing-policies"

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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import4.ListRoutingPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the Routing Policy corresponding to the extId.
func (api *RoutingPoliciesApi) UpdateRoutingPolicyById(extId *string, body *import4.RoutingPolicy, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRoutingPolicyById(context.Background(), &import25.UpdateRoutingPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the Routing Policy corresponding to the extId.
func (api *RoutingPoliciesServiceApi) UpdateRoutingPolicyById(ctx context.Context, request *import25.UpdateRoutingPolicyByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/routing-policies/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
