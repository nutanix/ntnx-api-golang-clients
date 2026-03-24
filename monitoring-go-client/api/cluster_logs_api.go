package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/clusterlogs"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type ClusterLogsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClusterLogsServiceApi
}

type ClusterLogsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClusterLogsApi(apiClient *client.ApiClient) *ClusterLogsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterLogsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClusterLogsServiceApi(a.ApiClient)

	return a
}

func NewClusterLogsServiceApi(apiClient *client.ApiClient) *ClusterLogsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterLogsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Collect logs based on the input parameters.
func (api *ClusterLogsApi) CollectLogs(extId *string, body *import1.LogCollectionSpec, args ...map[string]interface{}) (*import1.CollectLogsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterLogsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CollectLogs(context.Background(), &import5.CollectLogsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Collect logs based on the input parameters.
func (api *ClusterLogsServiceApi) CollectLogs(ctx context.Context, request *import5.CollectLogsRequest, args ...map[string]interface{}) (*import1.CollectLogsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/clusters/{extId}/$actions/collect-logs"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.CollectLogsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of all available tags for a cluster.
func (api *ClusterLogsApi) ListTags(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListTagsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterLogsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTags(context.Background(), &import5.ListTagsRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// Fetches a list of all available tags for a cluster.
func (api *ClusterLogsServiceApi) ListTags(ctx context.Context, request *import5.ListTagsRequest, args ...map[string]interface{}) (*import1.ListTagsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/clusters/{clusterExtId}/tags"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListTagsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
