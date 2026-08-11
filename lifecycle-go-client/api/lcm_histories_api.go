package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import15 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/lcmhistories"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type LcmHistoriesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *LcmHistoriesServiceApi
}

type LcmHistoriesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewLcmHistoriesApi(apiClient *client.ApiClient) *LcmHistoriesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LcmHistoriesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewLcmHistoriesServiceApi(a.ApiClient)

	return a
}

func NewLcmHistoriesServiceApi(apiClient *client.ApiClient) *LcmHistoriesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LcmHistoriesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Export the LCM operation history of connected clusters as a downloadable file. The exported file contains records of past inventory, upgrade, and upload operations including their statuses, timestamps, and component details. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the exported file is available in the completion_details field of the task. Submit an ExportHistorySpec with the desired file format (CSV). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *LcmHistoriesApi) ExportHistories(body *import1.ExportHistorySpec, args ...map[string]interface{}) (*import1.ExportHistoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportHistories(context.Background(), &import15.ExportHistoriesRequest{
		Body: body,
	}, args...)
}

// Export the LCM operation history of connected clusters as a downloadable file. The exported file contains records of past inventory, upgrade, and upload operations including their statuses, timestamps, and component details. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the exported file is available in the completion_details field of the task. Submit an ExportHistorySpec with the desired file format (CSV). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *LcmHistoriesServiceApi) ExportHistories(ctx context.Context, request *import15.ExportHistoriesRequest, args ...map[string]interface{}) (*import1.ExportHistoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/lcm-histories/$actions/export"

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
	unmarshalledResp := new(import1.ExportHistoriesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the full details of a single LCM operation history entry by its identifier. The response includes the operation type (inventory, upgrade, or upload), start and end timestamps, overall status, user information, LCM framework version, and detailed component-level results. Use this endpoint to inspect the outcome of a specific past operation.
func (api *LcmHistoriesApi) GetLcmHistoryById(extId *string, args ...map[string]interface{}) (*import1.GetLcmHistoryByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLcmHistoryById(context.Background(), &import15.GetLcmHistoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the full details of a single LCM operation history entry by its identifier. The response includes the operation type (inventory, upgrade, or upload), start and end timestamps, overall status, user information, LCM framework version, and detailed component-level results. Use this endpoint to inspect the outcome of a specific past operation.
func (api *LcmHistoriesServiceApi) GetLcmHistoryById(ctx context.Context, request *import15.GetLcmHistoryByIdRequest, args ...map[string]interface{}) (*import1.GetLcmHistoryByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/lcm-histories/{extId}"

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
	unmarshalledResp := new(import1.GetLcmHistoryByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a paginated list of LCM operation history entries across all clusters in a Prism Central environment. Supports standard query parameters for pagination ($page, $limit), filtering ($filter by operation type, status, cluster, time range), sorting ($orderby), and field projection ($select). Each history entry summarizes a past inventory, upgrade, or upload operation. Use POST /lcm-histories/$actions/export to download the full history as a file.
func (api *LcmHistoriesApi) ListLcmHistories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListLcmHistoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLcmHistories(context.Background(), &import15.ListLcmHistoriesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieve a paginated list of LCM operation history entries across all clusters in a Prism Central environment. Supports standard query parameters for pagination ($page, $limit), filtering ($filter by operation type, status, cluster, time range), sorting ($orderby), and field projection ($select). Each history entry summarizes a past inventory, upgrade, or upload operation. Use POST /lcm-histories/$actions/export to download the full history as a file.
func (api *LcmHistoriesServiceApi) ListLcmHistories(ctx context.Context, request *import15.ListLcmHistoriesRequest, args ...map[string]interface{}) (*import1.ListLcmHistoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/lcm-histories"

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
	unmarshalledResp := new(import1.ListLcmHistoriesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
