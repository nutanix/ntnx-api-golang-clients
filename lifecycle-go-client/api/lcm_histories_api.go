package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import9 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/lcmhistories"
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

// Download the history information of connected clusters as a file. This API will return a task, which on completion will provide a URL in the completion_details field to download the file containing the history information.
func (api *LcmHistoriesApi) ExportHistories(body *import1.ExportHistorySpec, args ...map[string]interface{}) (*import1.ExportHistoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportHistories(context.Background(), &import9.ExportHistoriesRequest{
		Body: body,
	}, args...)
}

// Download the history information of connected clusters as a file. This API will return a task, which on completion will provide a URL in the completion_details field to download the file containing the history information.
func (api *LcmHistoriesServiceApi) ExportHistories(ctx context.Context, request *import9.ExportHistoriesRequest, args ...map[string]interface{}) (*import1.ExportHistoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/lcm-histories/$actions/export"

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

	unmarshalledResp := new(import1.ExportHistoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query an LCM History entry by its id.
func (api *LcmHistoriesApi) GetLcmHistoryById(extId *string, args ...map[string]interface{}) (*import1.GetLcmHistoryByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLcmHistoryById(context.Background(), &import9.GetLcmHistoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Query an LCM History entry by its id.
func (api *LcmHistoriesServiceApi) GetLcmHistoryById(ctx context.Context, request *import9.GetLcmHistoryByIdRequest, args ...map[string]interface{}) (*import1.GetLcmHistoryByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/lcm-histories/{extId}"

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

	unmarshalledResp := new(import1.GetLcmHistoryByIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query list of LCM histories.
func (api *LcmHistoriesApi) ListLcmHistories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListLcmHistoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmHistoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLcmHistories(context.Background(), &import9.ListLcmHistoriesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Query list of LCM histories.
func (api *LcmHistoriesServiceApi) ListLcmHistories(ctx context.Context, request *import9.ListLcmHistoriesRequest, args ...map[string]interface{}) (*import1.ListLcmHistoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/lcm-histories"

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

	unmarshalledResp := new(import1.ListLcmHistoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
