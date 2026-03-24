package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/stats"
	import14 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/stats"
	import15 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/stats"
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

// Fetches the stats for the specified VM disk. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsApi) GetDiskStatsById(vmExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import14.GetDiskStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDiskStatsById(context.Background(), &import15.GetDiskStatsByIdRequest{
		VmExtId:           vmExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Fetches the stats for the specified VM disk. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsServiceApi) GetDiskStatsById(ctx context.Context, request *import15.GetDiskStatsByIdRequest, args ...map[string]interface{}) (*import14.GetDiskStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/stats/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import14.GetDiskStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the stats for the specified VM NIC. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsApi) GetNicStatsById(vmExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import14.GetNicStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNicStatsById(context.Background(), &import15.GetNicStatsByIdRequest{
		VmExtId:           vmExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Fetches the stats for the specified VM NIC. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsServiceApi) GetNicStatsById(ctx context.Context, request *import15.GetNicStatsByIdRequest, args ...map[string]interface{}) (*import14.GetNicStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/stats/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import14.GetNicStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get VM stats for a given VM. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsApi) GetVmStatsById(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import14.GetVmStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmStatsById(context.Background(), &import15.GetVmStatsByIdRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Get VM stats for a given VM. Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/checkScore'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data
func (api *StatsServiceApi) GetVmStatsById(ctx context.Context, request *import15.GetVmStatsByIdRequest, args ...map[string]interface{}) (*import14.GetVmStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/stats/vms/{extId}"

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

	unmarshalledResp := new(import14.GetVmStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List VM stats for all VMs.  Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/controllerNumIo,stats/hypervisorNumIo'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data; 6) '$orderby'; 7) '$page'; 8) '$limit'; and 9) '$filter': the OData filter to use, e.g. 'stats/hypervisorCpuUsagePpm gt 100000 and stats/guestMemoryUsagePpm lt 2000000.'
func (api *StatsApi) ListVmStats(startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import14.ListVmStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStats(context.Background(), &import15.ListVmStatsRequest{
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Page_:             page_,
		Limit_:            limit_,
		Filter_:           filter_,
		Orderby_:          orderby_,
		Select_:           select_,
	}, args...)
}

// List VM stats for all VMs.  Users can fetch the stats by specifying the following params in the request query: 1) '$select': comma-separated attributes with the prefix 'stats/', e.g. 'stats/controllerNumIo,stats/hypervisorNumIo'. 2) '$startTime': the start time for which stats should be reported, e.g. '2023-01-01T12:00:00.000-08:00'; 3) '$endTime': the end time for which stats should be reported; 4) '$samplingInterval': the sampling interval in seconds at which statistical data should be collected; 5) '$statType': the down-sampling operator to use while performing down-sampling on stats data; 6) '$orderby'; 7) '$page'; 8) '$limit'; and 9) '$filter': the OData filter to use, e.g. 'stats/hypervisorCpuUsagePpm gt 100000 and stats/guestMemoryUsagePpm lt 2000000.'
func (api *StatsServiceApi) ListVmStats(ctx context.Context, request *import15.ListVmStatsRequest, args ...map[string]interface{}) (*import14.ListVmStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/stats/vms"

	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

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

	unmarshalledResp := new(import14.ListVmStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
