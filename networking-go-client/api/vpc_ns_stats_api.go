//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/stats"
	"net/http"
	"net/url"
	"strings"
)

type VpcNsStatsApi struct {
	ApiClient *client.ApiClient
}

func NewVpcNsStatsApi(apiClient *client.ApiClient) *VpcNsStatsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcNsStatsApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Get VPC North-South statistics.
  Get VPC North-South statistics

  parameters:-
  -> vpcUuid (string) (required) : The UUID of the VPC.
  -> extSubnetUuid (string) (required) : The UUID of the VPC.
  -> startTime_ (string) (optional) : The start time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For example, start time of 2022-04-23T01:23:45.678+09:00 would consider all stats starting at 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
  -> endTime_ (string) (optional) : The end time of the period for which stats should be reported. The value should be in extended ISO-8601 format. For example, end time of 2022-04-23T013:23:45.678+09:00 would consider all stats till 13:23:45 .678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
  -> samplingInterval_ (int) (optional) : The sampling interval in seconds at which statistical data should be collected For example, do you want performance statistics every 30 seconds? Every 60 seconds?
  -> statType_ (common.v1.stats.DownSamplingOperator) (optional)
  -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - entityUuid - extId - links - northSouthEgressBytesAbs - northSouthEgressBytesPerSec - northSouthEgressPacketsAbs - northSouthEgressPacketsPerSec - northSouthIngressBytesAbs - northSouthIngressBytesPerSec - northSouthIngressPacketsAbs - northSouthIngressPacketsPerSec - statType - tenantId
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.stats.VpcNsStatsApiResponse, error)
*/
func (api *VpcNsStatsApi) GetVpcNsStats(vpcUuid *string, extSubnetUuid *string, startTime_ *string, endTime_ *string, samplingInterval_ *int, statType_ *import2.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import3.VpcNsStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/stats/vpc/{vpcUuid}/external-subnets/{extSubnetUuid}"

	// verify the required parameter 'vpcUuid' is set
	if nil == vpcUuid {
		return nil, client.ReportError("vpcUuid is required and must be specified")
	}
	// verify the required parameter 'extSubnetUuid' is set
	if nil == extSubnetUuid {
		return nil, client.ReportError("extSubnetUuid is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpcUuid"+"}", url.PathEscape(client.ParameterToString(*vpcUuid, "")), -1)
	uri = strings.Replace(uri, "{"+"extSubnetUuid"+"}", url.PathEscape(client.ParameterToString(*extSubnetUuid, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if startTime_ != nil {
		queryParams.Add("$startTime", client.ParameterToString(*startTime_, ""))
	}
	if endTime_ != nil {
		queryParams.Add("$endTime", client.ParameterToString(*endTime_, ""))
	}
	if samplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*samplingInterval_, ""))
	}
	if statType_ != nil {
		queryParams.Add("$statType", client.ParameterToString(*statType_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.VpcNsStatsApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
