package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import10 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/gateways"
	"net/http"
	"net/url"
	"strings"
)

type GatewaysApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *GatewaysServiceApi
}

type GatewaysServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewGatewaysApi(apiClient *client.ApiClient) *GatewaysApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GatewaysApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewGatewaysServiceApi(a.ApiClient)

	return a
}

func NewGatewaysServiceApi(apiClient *client.ApiClient) *GatewaysServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GatewaysServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create gateway.
func (api *GatewaysApi) CreateGateway(body *import4.Gateway, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateGateway(context.Background(), &import10.CreateGatewayRequest{
		Body: body,
	}, args...)
}

// Create gateway.
func (api *GatewaysServiceApi) CreateGateway(ctx context.Context, request *import10.CreateGatewayRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways"

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

// Delete gateway for the specified UUID.
func (api *GatewaysApi) DeleteGatewayById(extId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteGatewayById(context.Background(), &import10.DeleteGatewayByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete gateway for the specified UUID.
func (api *GatewaysServiceApi) DeleteGatewayById(ctx context.Context, request *import10.DeleteGatewayByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways/{extId}"

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

// Get the gateway for the specified extId.
func (api *GatewaysApi) GetGatewayById(extId *string, args ...map[string]interface{}) (*import4.GetGatewayApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetGatewayById(context.Background(), &import10.GetGatewayByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get the gateway for the specified extId.
func (api *GatewaysServiceApi) GetGatewayById(ctx context.Context, request *import10.GetGatewayByIdRequest, args ...map[string]interface{}) (*import4.GetGatewayApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways/{extId}"

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

	unmarshalledResp := new(import4.GetGatewayApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the list of the existing Network gateways.
func (api *GatewaysApi) ListGateways(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListGatewaysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListGateways(context.Background(), &import10.ListGatewaysRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Get the list of the existing Network gateways.
func (api *GatewaysServiceApi) ListGateways(ctx context.Context, request *import10.ListGatewaysRequest, args ...map[string]interface{}) (*import4.ListGatewaysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways"

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

	unmarshalledResp := new(import4.ListGatewaysApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update gateway.
func (api *GatewaysApi) UpdateGatewayById(extId *string, body *import4.Gateway, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateGatewayById(context.Background(), &import10.UpdateGatewayByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update gateway.
func (api *GatewaysServiceApi) UpdateGatewayById(ctx context.Context, request *import10.UpdateGatewayByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways/{extId}"

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

// Upgrade the gateway for the specified UUID.
func (api *GatewaysApi) UpgradeGatewayById(extId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGatewaysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpgradeGatewayById(context.Background(), &import10.UpgradeGatewayByIdRequest{
		ExtId: extId,
	}, args...)
}

// Upgrade the gateway for the specified UUID.
func (api *GatewaysServiceApi) UpgradeGatewayById(ctx context.Context, request *import10.UpgradeGatewayByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/gateways/{extId}/$actions/upgrade"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
