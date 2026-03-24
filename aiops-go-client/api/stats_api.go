package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/request/stats"
	import5 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/stats"
	import6 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StatsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StatsServiceApi
}

type StatsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStatsApi(apiClient *client.ApiClient) *StatsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStatsServiceApi(a.ApiClient)

	return a
}

func NewStatsServiceApi(apiClient *client.ApiClient) *StatsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Returns metadata information available for entity and metrics for given entities.
func (api *StatsApi) GetEntityDescriptorsV4(sourceExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.EntityDescriptorListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityDescriptorsV4(context.Background(), &import7.GetEntityDescriptorsV4Request{
		SourceExtId: sourceExtId,
		Page_:       page_,
		Limit_:      limit_,
		Filter_:     filter_,
	}, args...)
}

// Returns metadata information available for entity and metrics for given entities.
func (api *StatsServiceApi) GetEntityDescriptorsV4(ctx context.Context, request *import7.GetEntityDescriptorsV4Request, args ...map[string]interface{}) (*import1.EntityDescriptorListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/sources/{sourceExtId}/entity-descriptors"

	// verify the required parameter 'sourceExtId' is set
	if nil == request.SourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*request.SourceExtId, "")), -1)
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

	unmarshalledResp := new(import1.EntityDescriptorListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of attributes and metrics (time series data) that are available for a given entity type.
func (api *StatsApi) GetEntityMetricsV4(sourceExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, page_ *int, limit_ *int, samplingInterval_ *int, statType_ *import6.DownSamplingOperator, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import5.EntityListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityMetricsV4(context.Background(), &import7.GetEntityMetricsV4Request{
		SourceExtId:       sourceExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		Page_:             page_,
		Limit_:            limit_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Filter_:           filter_,
		Orderby_:          orderby_,
	}, args...)
}

// Returns a list of attributes and metrics (time series data) that are available for a given entity type.
func (api *StatsServiceApi) GetEntityMetricsV4(ctx context.Context, request *import7.GetEntityMetricsV4Request, args ...map[string]interface{}) (*import5.EntityListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/stats/sources/{sourceExtId}/entities/{extId}"

	// verify the required parameter 'sourceExtId' is set
	if nil == request.SourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*request.SourceExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.EntityListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available entity types and their UUIDs for a given source. These UUIDs can further be used for other APIs to retrieve entity metrics and descriptors.
func (api *StatsApi) GetEntityTypesV4(sourceExtId *string, args ...map[string]interface{}) (*import1.EntityTypeListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityTypesV4(context.Background(), &import7.GetEntityTypesV4Request{
		SourceExtId: sourceExtId,
	}, args...)
}

// Returns a list of available entity types and their UUIDs for a given source. These UUIDs can further be used for other APIs to retrieve entity metrics and descriptors.
func (api *StatsServiceApi) GetEntityTypesV4(ctx context.Context, request *import7.GetEntityTypesV4Request, args ...map[string]interface{}) (*import1.EntityTypeListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/sources/{sourceExtId}/entity-types"

	// verify the required parameter 'sourceExtId' is set
	if nil == request.SourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*request.SourceExtId, "")), -1)
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

	unmarshalledResp := new(import1.EntityTypeListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available sources and their UUIDs. These UUIDs can further be used for other APIs to retrieve entity types and their metrics.
func (api *StatsApi) GetSourcesV4(args ...map[string]interface{}) (*import1.SourceListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSourcesV4(context.Background(), &import7.GetSourcesV4Request{}, args...)
}

// Returns a list of available sources and their UUIDs. These UUIDs can further be used for other APIs to retrieve entity types and their metrics.
func (api *StatsServiceApi) GetSourcesV4(ctx context.Context, request *import7.GetSourcesV4Request, args ...map[string]interface{}) (*import1.SourceListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/sources"

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

	unmarshalledResp := new(import1.SourceListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
