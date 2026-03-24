package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/stats"
	import7 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/request/stats"
	import6 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/stats"
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

// Get the time series stats of an Object store. If $samplingInterval is not set, a value of 120 seconds will be used. If $statType is not set, the operator \"SUM\" will be used. The stats can be fetched for an Object store when it is in the state \"OBJECT_STORE_AVAILABLE\".
func (api *StatsApi) GetObjectstoreStatsById(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import5.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import6.GetObjectstoreStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetObjectstoreStatsById(context.Background(), &import7.GetObjectstoreStatsByIdRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Get the time series stats of an Object store. If $samplingInterval is not set, a value of 120 seconds will be used. If $statType is not set, the operator \"SUM\" will be used. The stats can be fetched for an Object store when it is in the state \"OBJECT_STORE_AVAILABLE\".
func (api *StatsServiceApi) GetObjectstoreStatsById(ctx context.Context, request *import7.GetObjectstoreStatsByIdRequest, args ...map[string]interface{}) (*import6.GetObjectstoreStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/stats/object-stores/{extId}"

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
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import6.GetObjectstoreStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
