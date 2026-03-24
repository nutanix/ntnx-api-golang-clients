package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/reportconfig"
	"net/http"
	"net/url"
	"strings"
)

type ReportConfigApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ReportConfigServiceApi
}

type ReportConfigServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewReportConfigApi(apiClient *client.ApiClient) *ReportConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportConfigApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewReportConfigServiceApi(a.ApiClient)

	return a
}

func NewReportConfigServiceApi(apiClient *client.ApiClient) *ReportConfigServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportConfigServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a custom report configuration using the request body.
func (api *ReportConfigApi) CreateReportConfig(body *import1.ReportConfig, args ...map[string]interface{}) (*import1.CreateReportConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateReportConfig(context.Background(), &import7.CreateReportConfigRequest{
		Body: body,
	}, args...)
}

// Create a custom report configuration using the request body.
func (api *ReportConfigServiceApi) CreateReportConfig(ctx context.Context, request *import7.CreateReportConfigRequest, args ...map[string]interface{}) (*import1.CreateReportConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/report-configs"

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

	unmarshalledResp := new(import1.CreateReportConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a report configuration with the given UUID.
func (api *ReportConfigApi) DeleteReportConfigById(extId *string, args ...map[string]interface{}) (*import1.DeleteReportConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteReportConfigById(context.Background(), &import7.DeleteReportConfigByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a report configuration with the given UUID.
func (api *ReportConfigServiceApi) DeleteReportConfigById(ctx context.Context, request *import7.DeleteReportConfigByIdRequest, args ...map[string]interface{}) (*import1.DeleteReportConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/report-configs/{extId}"

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

	unmarshalledResp := new(import1.DeleteReportConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns report configuration information for the provided UUID.
func (api *ReportConfigApi) GetReportConfigById(extId *string, args ...map[string]interface{}) (*import1.GetReportConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetReportConfigById(context.Background(), &import7.GetReportConfigByIdRequest{
		ExtId: extId,
	}, args...)
}

// Returns report configuration information for the provided UUID.
func (api *ReportConfigServiceApi) GetReportConfigById(ctx context.Context, request *import7.GetReportConfigByIdRequest, args ...map[string]interface{}) (*import1.GetReportConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/report-configs/{extId}"

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

	unmarshalledResp := new(import1.GetReportConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Endpoint for listing all report configurations accessible to the user that match the filter criteria.
func (api *ReportConfigApi) ListReportConfigs(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListReportConfigsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListReportConfigs(context.Background(), &import7.ListReportConfigsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Endpoint for listing all report configurations accessible to the user that match the filter criteria.
func (api *ReportConfigServiceApi) ListReportConfigs(ctx context.Context, request *import7.ListReportConfigsRequest, args ...map[string]interface{}) (*import1.ListReportConfigsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/report-configs"

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

	unmarshalledResp := new(import1.ListReportConfigsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the report configuration information for the provided UUID with the given specifications.
func (api *ReportConfigApi) UpdateReportConfigById(extId *string, body *import1.ReportConfig, args ...map[string]interface{}) (*import1.UpdateReportConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateReportConfigById(context.Background(), &import7.UpdateReportConfigByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the report configuration information for the provided UUID with the given specifications.
func (api *ReportConfigServiceApi) UpdateReportConfigById(ctx context.Context, request *import7.UpdateReportConfigByIdRequest, args ...map[string]interface{}) (*import1.UpdateReportConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/report-configs/{extId}"

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

	unmarshalledResp := new(import1.UpdateReportConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
