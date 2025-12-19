package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/stats"
	import5 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StatsApi struct {
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

	return a
}

// Returns metadata information available for entity and metrics for given entities.
func (api *StatsApi) GetEntityDescriptorsV4(sourceExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.EntityDescriptorListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/sources/{sourceExtId}/entity-descriptors"

	// verify the required parameter 'sourceExtId' is set
	if nil == sourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*sourceExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
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

	unmarshalledResp := new(import1.EntityDescriptorListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of attributes and metrics (time series data) that are available for a given entity type.
func (api *StatsApi) GetEntityMetricsV4(sourceExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, page_ *int, limit_ *int, samplingInterval_ *int, statType_ *import5.DownSamplingOperator, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.EntityListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/stats/sources/{sourceExtId}/entities/{extId}"

	// verify the required parameter 'sourceExtId' is set
	if nil == sourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*sourceExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	queryParams.Add("$startTime", client.ParameterToString(*startTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*endTime_, ""))
	if samplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*samplingInterval_, ""))
	}
	if statType_ != nil {
		statType_QueryParamEnumVal := statType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
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

	unmarshalledResp := new(import4.EntityListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available entity types and their UUIDs for a given source. These UUIDs can further be used for other APIs to retrieve entity metrics and descriptors.
func (api *StatsApi) GetEntityTypesV4(sourceExtId *string, args ...map[string]interface{}) (*import1.EntityTypeListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/sources/{sourceExtId}/entity-types"

	// verify the required parameter 'sourceExtId' is set
	if nil == sourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"sourceExtId"+"}", url.PathEscape(client.ParameterToString(*sourceExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.EntityTypeListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available sources and their UUIDs. These UUIDs can further be used for other APIs to retrieve entity types and their metrics.
func (api *StatsApi) GetSourcesV4(args ...map[string]interface{}) (*import1.SourceListApiResponse, error) {
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.SourceListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
