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
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVpcNsStatsApi(apiClient *client.ApiClient) *VpcNsStatsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcNsStatsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get VPC North-South statistics. Requires Prism Central >= pc.2022.9.
func (api *VpcNsStatsApi) GetVpcNsStats(vpcExtId *string, extId *string, startTime_ *string, endTime_ *string, samplingInterval_ *int, statType_ *import2.DownSamplingOperator, page_ *int, limit_ *int, select_ *string, args ...map[string]interface{}) (*import3.VpcNsStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/stats/vpc/{vpcExtId}/external-subnets/{extId}"

	// verify the required parameter 'vpcExtId' is set
	if nil == vpcExtId {
		return nil, client.ReportError("vpcExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpcExtId"+"}", url.PathEscape(client.ParameterToString(*vpcExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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
		enumVal := statType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(enumVal, ""))
	}
	if page_ != nil {

		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {

		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if select_ != nil {

		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
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
