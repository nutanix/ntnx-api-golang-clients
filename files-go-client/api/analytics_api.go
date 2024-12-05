package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/stats"
	import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AnalyticsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAnalyticsApi(apiClient *client.ApiClient) *AnalyticsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AnalyticsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the antivirus server statistics for the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and identifier of the antivirus server (`extId`) and comma-separated list of metrics in `$select`. The user can also specify `$startTime` and `$endTime` in query parameters.  `$samplingInterval` should be either empty, 300(default) or 600. `$statType` is not supported for this API.  Following are some valid GET antivirus statistics query examples: ``` - ?$select=scannedFileCount,threatCount - ?$select=scannedFileCount,threatCount&$startTime=2023-01-09T07:23:45.678&$endTime=2023-01-09T18:23:45.678 ```
func (api *AnalyticsApi) GetAntivirusServerStats(fileServerExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import2.AntivirusServerStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/stats/file-servers/{fileServerExtId}/anti-virus-servers/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == startTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == endTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*startTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*endTime_, ""))
	if samplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*samplingInterval_, ""))
	}
	if statType_ != nil {
		statType_QueryParamEnumVal := statType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.AntivirusServerStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the file server statistics for the specified file server.  Specify a comma-separated list of metrics in `$select`. Users can also specify `$startTime` and `$endTime` in query parameters.  `$samplingInterval` should be either empty, 300(default) or 600. `$statType` is not supported for this API.  Following are some valid GET file server statistics query examples: ``` - ?$select=numberOfFiles,numberOfConnections - ?$select=numberOfFiles,numberOfConnections&$startTime=2023-01-09T07:23:45.678&$endTime=2023-01-09T18:23:45.678 ```
func (api *AnalyticsApi) GetFileServerStats(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import2.FileServerStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/stats/file-servers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == startTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == endTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*startTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*endTime_, ""))
	if samplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*samplingInterval_, ""))
	}
	if statType_ != nil {
		statType_QueryParamEnumVal := statType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.FileServerStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the mount target statistics for the mount target with the given external identifier.  Specify a valid file server identifier (`fileServerExtId`) and a mount target identifier (`extId`) and comma-separated list of metrics in `$select`. The user can also specify `$startTime` and `$endTime` in query parameters.  `$samplingInterval` should be either empty, 300(default) or 600. `$statType` is not supported for this API.  Following are some valid GET mount target statistics examples: ``` - ?$select=averageLatencyMs,numberOfFiles - ?$select=averageLatencyMs,numberOfFiles&$startTime=2023-01-09T07:23:45.678&$endTime=2023-01-09T18:23:45.678 ```
func (api *AnalyticsApi) GetMountTargetStats(fileServerExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import1.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import2.MountTargetStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/stats/file-servers/{fileServerExtId}/mount-targets/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == startTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == endTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*startTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*endTime_, ""))
	if samplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*samplingInterval_, ""))
	}
	if statType_ != nil {
		statType_QueryParamEnumVal := statType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.MountTargetStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
