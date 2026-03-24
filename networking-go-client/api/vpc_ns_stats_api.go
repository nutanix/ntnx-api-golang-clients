package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import12 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	import35 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/vpcnsstats"
	import13 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type VpcNsStatsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VpcNsStatsServiceApi
}

type VpcNsStatsServiceApi struct {
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

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVpcNsStatsServiceApi(a.ApiClient)

	return a
}

func NewVpcNsStatsServiceApi(apiClient *client.ApiClient) *VpcNsStatsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcNsStatsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get VPC North-South statistics.
func (api *VpcNsStatsApi) GetVpcNsStats(vpcExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import12.DownSamplingOperator, page_ *int, limit_ *int, select_ *string, args ...map[string]interface{}) (*import13.GetVpcNsStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpcNsStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVpcNsStats(context.Background(), &import35.GetVpcNsStatsRequest{
		VpcExtId:          vpcExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Page_:             page_,
		Limit_:            limit_,
		Select_:           select_,
	}, args...)
}

// Get VPC North-South statistics.
func (api *VpcNsStatsServiceApi) GetVpcNsStats(ctx context.Context, request *import35.GetVpcNsStatsRequest, args ...map[string]interface{}) (*import13.GetVpcNsStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/stats/vpc/{vpcExtId}/external-subnets/{extId}"

	// verify the required parameter 'vpcExtId' is set
	if nil == request.VpcExtId {
		return nil, client.ReportError("vpcExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"vpcExtId"+"}", url.PathEscape(client.ParameterToString(*request.VpcExtId, "")), -1)
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
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
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

	unmarshalledResp := new(import13.GetVpcNsStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
