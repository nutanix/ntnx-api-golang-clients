//Api classes for aiops's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/stats"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
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

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Returns entity and metrics descriptors.
func (api *StatsApi) GetEntityDescriptorsV4(extId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.EntityDescriptorListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.0.a2/stats/sources/{extId}/entity-descriptors"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
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
	unmarshalledResp := new(import1.EntityDescriptorListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns attribute/metric data for the selected entities.
func (api *StatsApi) GetEntityMetricsV4(sourceExtId *string, extId *string, page_ *int, limit_ *int, startTime_ *string, endTime_ *string, samplingInterval_ *int, statType_ *import2.DownSamplingOperator, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.EntityListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.0.a2/stats/sources/{sourceExtId}/entities/{extId}"

	// verify the required parameter 'sourceExtId' is set
	if nil == sourceExtId {
		return nil, client.ReportError("sourceExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
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
	unmarshalledResp := new(import1.EntityListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available entity types for a given source.
func (api *StatsApi) GetEntityTypesV4(extId *string, args ...map[string]interface{}) (*import1.EntityTypeListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.0.a2/stats/sources/{extId}/entity-types"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
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
	unmarshalledResp := new(import1.EntityTypeListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of available sources.
func (api *StatsApi) GetSourcesV4(args ...map[string]interface{}) (*import1.SourceListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.0.a2/stats/sources"

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
	unmarshalledResp := new(import1.SourceListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
