package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import24 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/routes"
	"net/http"
	"net/url"
	"strings"
)

type RoutesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RoutesServiceApi
}

type RoutesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRoutesApi(apiClient *client.ApiClient) *RoutesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRoutesServiceApi(a.ApiClient)

	return a
}

func NewRoutesServiceApi(apiClient *client.ApiClient) *RoutesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a user-configured route for a specified .
func (api *RoutesApi) CreateRouteForRouteTable(routeTableExtId *string, body *import4.Route, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRouteForRouteTable(context.Background(), &import24.CreateRouteForRouteTableRequest{
		RouteTableExtId: routeTableExtId,
		Body:            body,
	}, args...)
}

// Create a user-configured route for a specified .
func (api *RoutesServiceApi) CreateRouteForRouteTable(ctx context.Context, request *import24.CreateRouteForRouteTableRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/route-tables/{routeTableExtId}/routes"

	// verify the required parameter 'routeTableExtId' is set
	if nil == request.RouteTableExtId {
		return nil, client.ReportError("routeTableExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"routeTableExtId"+"}", url.PathEscape(client.ParameterToString(*request.RouteTableExtId, "")), -1)
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

// Delete route for the specified {extId}.
func (api *RoutesApi) DeleteRouteForRouteTableById(extId *string, routeTableExtId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRouteForRouteTableById(context.Background(), &import24.DeleteRouteForRouteTableByIdRequest{
		ExtId:           extId,
		RouteTableExtId: routeTableExtId,
	}, args...)
}

// Delete route for the specified {extId}.
func (api *RoutesServiceApi) DeleteRouteForRouteTableById(ctx context.Context, request *import24.DeleteRouteForRouteTableByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/route-tables/{routeTableExtId}/routes/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'routeTableExtId' is set
	if nil == request.RouteTableExtId {
		return nil, client.ReportError("routeTableExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"routeTableExtId"+"}", url.PathEscape(client.ParameterToString(*request.RouteTableExtId, "")), -1)
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

// Get the specified user/system-configured route of the specified route table.
func (api *RoutesApi) GetRouteForRouteTableById(extId *string, routeTableExtId *string, args ...map[string]interface{}) (*import4.GetRouteApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRouteForRouteTableById(context.Background(), &import24.GetRouteForRouteTableByIdRequest{
		ExtId:           extId,
		RouteTableExtId: routeTableExtId,
	}, args...)
}

// Get the specified user/system-configured route of the specified route table.
func (api *RoutesServiceApi) GetRouteForRouteTableById(ctx context.Context, request *import24.GetRouteForRouteTableByIdRequest, args ...map[string]interface{}) (*import4.GetRouteApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/route-tables/{routeTableExtId}/routes/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'routeTableExtId' is set
	if nil == request.RouteTableExtId {
		return nil, client.ReportError("routeTableExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"routeTableExtId"+"}", url.PathEscape(client.ParameterToString(*request.RouteTableExtId, "")), -1)
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

	unmarshalledResp := new(import4.GetRouteApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists user/system-configured routes of the specified route table.
func (api *RoutesApi) ListRoutesByRouteTableId(routeTableExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListRoutesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRoutesByRouteTableId(context.Background(), &import24.ListRoutesByRouteTableIdRequest{
		RouteTableExtId: routeTableExtId,
		Page_:           page_,
		Limit_:          limit_,
		Filter_:         filter_,
		Orderby_:        orderby_,
	}, args...)
}

// Lists user/system-configured routes of the specified route table.
func (api *RoutesServiceApi) ListRoutesByRouteTableId(ctx context.Context, request *import24.ListRoutesByRouteTableIdRequest, args ...map[string]interface{}) (*import4.ListRoutesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/route-tables/{routeTableExtId}/routes"

	// verify the required parameter 'routeTableExtId' is set
	if nil == request.RouteTableExtId {
		return nil, client.ReportError("routeTableExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"routeTableExtId"+"}", url.PathEscape(client.ParameterToString(*request.RouteTableExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListRoutesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the specified user-configured route of the specified route table.
func (api *RoutesApi) UpdateRouteForRouteTableById(extId *string, routeTableExtId *string, body *import4.Route, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRouteForRouteTableById(context.Background(), &import24.UpdateRouteForRouteTableByIdRequest{
		ExtId:           extId,
		RouteTableExtId: routeTableExtId,
		Body:            body,
	}, args...)
}

// Update the specified user-configured route of the specified route table.
func (api *RoutesServiceApi) UpdateRouteForRouteTableById(ctx context.Context, request *import24.UpdateRouteForRouteTableByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/route-tables/{routeTableExtId}/routes/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'routeTableExtId' is set
	if nil == request.RouteTableExtId {
		return nil, client.ReportError("routeTableExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"routeTableExtId"+"}", url.PathEscape(client.ParameterToString(*request.RouteTableExtId, "")), -1)
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
