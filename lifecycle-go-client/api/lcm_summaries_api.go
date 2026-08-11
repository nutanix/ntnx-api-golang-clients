package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import16 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/lcmsummaries"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type LcmSummariesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *LcmSummariesServiceApi
}

type LcmSummariesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewLcmSummariesApi(apiClient *client.ApiClient) *LcmSummariesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LcmSummariesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewLcmSummariesServiceApi(a.ApiClient)

	return a
}

func NewLcmSummariesServiceApi(apiClient *client.ApiClient) *LcmSummariesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LcmSummariesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieve the LCM summary for a specific cluster by its external identifier. The summary includes the installed framework version, available updates, hardware vendor, connectivity type, cluster capabilities, and indicators for URL accessibility and upgrade availability. The extId corresponds to the cluster UUID.
func (api *LcmSummariesApi) GetLcmSummaryById(extId *string, args ...map[string]interface{}) (*import1.GetLcmSummaryByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmSummariesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLcmSummaryById(context.Background(), &import16.GetLcmSummaryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the LCM summary for a specific cluster by its external identifier. The summary includes the installed framework version, available updates, hardware vendor, connectivity type, cluster capabilities, and indicators for URL accessibility and upgrade availability. The extId corresponds to the cluster UUID.
func (api *LcmSummariesServiceApi) GetLcmSummaryById(ctx context.Context, request *import16.GetLcmSummaryByIdRequest, args ...map[string]interface{}) (*import1.GetLcmSummaryByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/lcm-summaries/{extId}"

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
	unmarshalledResp := new(import1.GetLcmSummaryByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a paginated list of LCM summaries for all clusters in a Prism Central environment. Each summary provides a high-level overview of the LCM state on a cluster, including the installed framework version, available framework version, hardware vendor, connectivity type, cluster capabilities, URL accessibility, and compatibility bundle version. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Summaries are updated after each inventory or upgrade operation.
func (api *LcmSummariesApi) ListLcmSummaries(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListLcmSummariesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLcmSummariesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLcmSummaries(context.Background(), &import16.ListLcmSummariesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieve a paginated list of LCM summaries for all clusters in a Prism Central environment. Each summary provides a high-level overview of the LCM state on a cluster, including the installed framework version, available framework version, hardware vendor, connectivity type, cluster capabilities, URL accessibility, and compatibility bundle version. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Summaries are updated after each inventory or upgrade operation.
func (api *LcmSummariesServiceApi) ListLcmSummaries(ctx context.Context, request *import16.ListLcmSummariesRequest, args ...map[string]interface{}) (*import1.ListLcmSummariesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/lcm-summaries"

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
	unmarshalledResp := new(import1.ListLcmSummariesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
