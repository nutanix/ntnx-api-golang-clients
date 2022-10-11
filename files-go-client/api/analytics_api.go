//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	"strings"
	import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/stats"
	"encoding/json"
	"net/http"
    "net/url"
)

type AnalyticsApi struct {
  ApiClient *client.ApiClient
}

func NewAnalyticsApi(apiClient *client.ApiClient) *AnalyticsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AnalyticsApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get Antivirus Stats
    Get the antivirus server stats for the given external identifier.  The user has to specify - a valid external identifier (`antivirusServerExtId`) of the antivirus server for which stats need to be fetched and comma separated list of metrics. The user can also specify `startTimeInUsecs` and `endTimeInUsecs` in query parameters  Available list of metrics are scanned_file_count, threat_count, cleaned_file_count, quarantined_file_count, latency_ms, throughput_bytes, disconnect_count 

    parameters:-
    -> antivirusServerExtId (string) (required) : Antivirus server UUID. Example:18f78959-14a6-4c47-b5db-920460c4b668.
    -> metrics (string) (required) : List of metrics. scanned_file_count, threat_count, cleaned_file_count, quarantined_file_count, latency_ms, throughput_bytes, disconnect_count
    -> startTimeInUsecs (int64) (optional) : Start time in microseconds to retrieve all the stats generated after this timestamp. For example: 1622705280584000
    -> endTimeInUsecs (int64) (optional) : End time in microseconds to retrieve all the stats generated before this timestamp. For example: 1622791680585000
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.stats.AntivirusServerStatsApiResponse, error)
*/
func (api *AnalyticsApi) GetAntivirusServerStats(antivirusServerExtId *string, metrics *string, startTimeInUsecs *int64, endTimeInUsecs *int64, args ...map[string]interface{}) (*import2.AntivirusServerStatsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/stats/file-server/anti-virus-servers/{antivirusServerExtId}"

    // verify the required parameter 'antivirusServerExtId' is set
	if nil == antivirusServerExtId {
		return nil, client.ReportError("antivirusServerExtId is required and must be specified")
	}
    // verify the required parameter 'metrics' is set
	if nil == metrics {
		return nil, client.ReportError("metrics is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"antivirusServerExtId"+"}", url.PathEscape(client.ParameterToString(*antivirusServerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	queryParams.Add("metrics", client.ParameterToString(*metrics, ""))
	if startTimeInUsecs != nil {
		queryParams.Add("startTimeInUsecs", client.ParameterToString(*startTimeInUsecs, ""))
	}
	if endTimeInUsecs != nil {
		queryParams.Add("endTimeInUsecs", client.ParameterToString(*endTimeInUsecs, ""))
	}

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.AntivirusServerStatsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get File server stats
    Get the file server stats for the file server.  The user has to specify a comma separated list of `metrics`. User can also specify `startTimeInUsecs` and `endTimeInUsecs`  in path parameters.  Available list of metrics are number_of_files, number_of_connections, used_bytes, snapshot_used_bytes, latency,throughput, iops, write_latency, read_latency, metadata_latency, write_throughput, read_throughput, read_iops, write_iops, metadata_iops. 

    parameters:-
    -> metrics (string) (required) : List of metrics. number_of_files, number_of_connections, used_bytes, snapshot_used_bytes, latency,throughput, iops, write_latency, read_latency, metadata_latency, write_throughput, read_throughput, read_iops, write_iops, metadata_iops
    -> startTimeInUsecs (int64) (optional) : Start time in microseconds to retrieve all the stats generated after this timestamp. For example: 1622705280584000
    -> endTimeInUsecs (int64) (optional) : End time in microseconds to retrieve all the stats generated before this timestamp. For example: 1622791680585000
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.stats.FileServerStatsApiResponse, error)
*/
func (api *AnalyticsApi) GetFileServerStats(metrics *string, startTimeInUsecs *int64, endTimeInUsecs *int64, args ...map[string]interface{}) (*import2.FileServerStatsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/stats/file-server"

    // verify the required parameter 'metrics' is set
	if nil == metrics {
		return nil, client.ReportError("metrics is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	queryParams.Add("metrics", client.ParameterToString(*metrics, ""))
	if startTimeInUsecs != nil {
		queryParams.Add("startTimeInUsecs", client.ParameterToString(*startTimeInUsecs, ""))
	}
	if endTimeInUsecs != nil {
		queryParams.Add("endTimeInUsecs", client.ParameterToString(*endTimeInUsecs, ""))
	}

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.FileServerStatsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get Mount target stats
    Get the mount target stats for the mount target with the given external identifier.  The user has to specify - a valid external identifier (`mountTargetExtId`) of the mount target for which the stats need to be fetched and comma separated list of metrics. The user can also specify `startTimeInUsecs` and `endTimeInUsecs` in query parameters  Available list of metrics are average_latency, average_throughput, read_latency, metadata_latency, write_throughput, read_throughput, average_iops, write_iops, read_iops, metadata_iops, number_of_files, number_of_connections, space_used_bytes, snapshot_used_bytes, s3_write_latency\" 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> metrics (string) (required) : List of metrics. average_latency, average_throughput, read_latency, metadata_latency, write_throughput, read_throughput, average_iops, write_iops, read_iops, metadata_iops, number_of_files, number_of_connections, space_used_bytes, snapshot_used_bytes
    -> startTimeInUsecs (int64) (optional) : Start time in microseconds to retrieve all the stats generated after this timestamp. For example: 1622705280584000
    -> endTimeInUsecs (int64) (optional) : End time in microseconds to retrieve all the stats generated before this timestamp. For example: 1622791680585000
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.stats.MountTargetStatsApiResponse, error)
*/
func (api *AnalyticsApi) GetMountTargetStats(mountTargetExtId *string, metrics *string, startTimeInUsecs *int64, endTimeInUsecs *int64, args ...map[string]interface{}) (*import2.MountTargetStatsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/stats/file-server/mount-targets/{mountTargetExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'metrics' is set
	if nil == metrics {
		return nil, client.ReportError("metrics is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	queryParams.Add("metrics", client.ParameterToString(*metrics, ""))
	if startTimeInUsecs != nil {
		queryParams.Add("startTimeInUsecs", client.ParameterToString(*startTimeInUsecs, ""))
	}
	if endTimeInUsecs != nil {
		queryParams.Add("endTimeInUsecs", client.ParameterToString(*endTimeInUsecs, ""))
	}

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.MountTargetStatsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

