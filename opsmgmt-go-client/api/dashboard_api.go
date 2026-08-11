package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/dashboard"
	"net/http"
	"net/url"
	"strings"
)

type DashboardApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DashboardServiceApi
}

type DashboardServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDashboardApi(apiClient *client.ApiClient) *DashboardApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDashboardServiceApi(a.ApiClient)

	return a
}

func NewDashboardServiceApi(apiClient *client.ApiClient) *DashboardServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a new dashboard.
func (api *DashboardApi) CreateDashboard(body *import1.Dashboard, args ...map[string]interface{}) (*import1.CreateDashboardApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDashboard(context.Background(), &import2.CreateDashboardRequest{
		Body: body,
	}, args...)
}

// Create a new dashboard.
func (api *DashboardServiceApi) CreateDashboard(ctx context.Context, request *import2.CreateDashboardRequest, args ...map[string]interface{}) (*import1.CreateDashboardApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboards"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.CreateDashboardApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete an existing dashboard.
func (api *DashboardApi) DeleteDashboardById(extId *string, args ...map[string]interface{}) (*import1.DeleteDashboardApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDashboardById(context.Background(), &import2.DeleteDashboardByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete an existing dashboard.
func (api *DashboardServiceApi) DeleteDashboardById(ctx context.Context, request *import2.DeleteDashboardByIdRequest, args ...map[string]interface{}) (*import1.DeleteDashboardApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboards/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.DeleteDashboardApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get a dashboard by ID.
func (api *DashboardApi) GetDashboardById(extId *string, args ...map[string]interface{}) (*import1.GetDashboardApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDashboardById(context.Background(), &import2.GetDashboardByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get a dashboard by ID.
func (api *DashboardServiceApi) GetDashboardById(ctx context.Context, request *import2.GetDashboardByIdRequest, args ...map[string]interface{}) (*import1.GetDashboardApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboards/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.GetDashboardApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all dashboards.
func (api *DashboardApi) ListDashboards(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDashboardApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDashboards(context.Background(), &import2.ListDashboardsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List all dashboards.
func (api *DashboardServiceApi) ListDashboards(ctx context.Context, request *import2.ListDashboardsRequest, args ...map[string]interface{}) (*import1.ListDashboardApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboards"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.ListDashboardApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update an existing dashboard.
func (api *DashboardApi) UpdateDashboardById(extId *string, body *import1.Dashboard, args ...map[string]interface{}) (*import1.UpdateDashboardApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDashboardById(context.Background(), &import2.UpdateDashboardByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update an existing dashboard.
func (api *DashboardServiceApi) UpdateDashboardById(ctx context.Context, request *import2.UpdateDashboardByIdRequest, args ...map[string]interface{}) (*import1.UpdateDashboardApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboards/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.UpdateDashboardApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
