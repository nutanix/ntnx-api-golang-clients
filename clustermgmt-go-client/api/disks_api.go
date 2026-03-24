package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import9 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/disks"
	import5 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/stats"
	import6 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DisksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DisksServiceApi
}

type DisksServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDisksApi(apiClient *client.ApiClient) *DisksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DisksApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDisksServiceApi(a.ApiClient)

	return a
}

func NewDisksServiceApi(apiClient *client.ApiClient) *DisksServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DisksServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Repartitions and adds the Disk to a cluster, or adds an old Disk again to a cluster that is marked for removal.
func (api *DisksApi) AddDisk(extId *string, body *import1.DiskAdditionSpec, args ...map[string]interface{}) (*import1.AddDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddDisk(context.Background(), &import9.AddDiskRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Repartitions and adds the Disk to a cluster, or adds an old Disk again to a cluster that is marked for removal.
func (api *DisksServiceApi) AddDisk(ctx context.Context, request *import9.AddDiskRequest, args ...map[string]interface{}) (*import1.AddDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{extId}/$actions/add-disk"

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

	unmarshalledResp := new(import1.AddDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Marks Disk identified by external identifier for removal.
func (api *DisksApi) DeleteDiskById(extId *string, args ...map[string]interface{}) (*import1.DeleteDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDiskById(context.Background(), &import9.DeleteDiskByIdRequest{
		ExtId: extId,
	}, args...)
}

// Marks Disk identified by external identifier for removal.
func (api *DisksServiceApi) DeleteDiskById(ctx context.Context, request *import9.DeleteDiskByIdRequest, args ...map[string]interface{}) (*import1.DeleteDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/disks/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetch the details of the Disk identified by external identifier.
func (api *DisksApi) GetDiskById(extId *string, args ...map[string]interface{}) (*import1.GetDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDiskById(context.Background(), &import9.GetDiskByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetch the details of the Disk identified by external identifier.
func (api *DisksServiceApi) GetDiskById(ctx context.Context, request *import9.GetDiskByIdRequest, args ...map[string]interface{}) (*import1.GetDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/disks/{extId}"

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

	unmarshalledResp := new(import1.GetDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetch the stats information of the Disk identified by external identifier.
func (api *DisksApi) GetDiskStats(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import6.DownSamplingOperator, args ...map[string]interface{}) (*import5.GetDiskStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDiskStats(context.Background(), &import9.GetDiskStatsRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
	}, args...)
}

// Fetch the stats information of the Disk identified by external identifier.
func (api *DisksServiceApi) GetDiskStats(ctx context.Context, request *import9.GetDiskStatsRequest, args ...map[string]interface{}) (*import5.GetDiskStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/stats/disks/{extId}"

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

	unmarshalledResp := new(import5.GetDiskStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches Disk details from all clusters registered with Prism Central.
func (api *DisksApi) ListDisks(page_ *int, limit_ *int, filter_ *string, orderby_ *string, apply_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDisksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDisks(context.Background(), &import9.ListDisksRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Apply_:   apply_,
		Select_:  select_,
	}, args...)
}

// Fetches Disk details from all clusters registered with Prism Central.
func (api *DisksServiceApi) ListDisks(ctx context.Context, request *import9.ListDisksRequest, args ...map[string]interface{}) (*import1.ListDisksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/disks"

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
	if request.Apply_ != nil {
		queryParams.Add("$apply", client.ParameterToString(*request.Apply_, ""))
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

	unmarshalledResp := new(import1.ListDisksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the LED state of a Disk to on or off.
func (api *DisksApi) UpdateDiskLEDState(extId *string, body *import1.LEDStateUpdationSpec, args ...map[string]interface{}) (*import1.UpdateDiskLEDStateTaskResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDisksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDiskLEDState(context.Background(), &import9.UpdateDiskLEDStateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the LED state of a Disk to on or off.
func (api *DisksServiceApi) UpdateDiskLEDState(ctx context.Context, request *import9.UpdateDiskLEDStateRequest, args ...map[string]interface{}) (*import1.UpdateDiskLEDStateTaskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/disks/{extId}/$actions/update-led-state"

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

	unmarshalledResp := new(import1.UpdateDiskLEDStateTaskResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
