package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type SnapshotSchedulesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSnapshotSchedulesApi(apiClient *client.ApiClient) *SnapshotSchedulesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SnapshotSchedulesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new snapshot schedule by using the provided request body. Self service restore (SSR) lets you restore a previous version of a file or share. With SSR, Files takes snapshots of share data in intervals defined by a snapshot schedule. You can enable SSR during or after share creation as the feature is disabled by default.  The user needs to specify the `maxRetentionCount` and `type` of the schedule to be created. User can also provide `daysOfWeek`, `daysOfMonth`, `frequency` and the objectType for these shecdules is `files.v4.config.WeeklySchedule`, `files.v4.config.MonthlySchedule`, `files.v4.config.DailySchedule` respectively.
func (api *SnapshotSchedulesApi) CreateSnapshotSchedule(fileServerExtId *string, body *import3.SnapshotSchedule, args ...map[string]interface{}) (*import3.CreateSnapshotScheduleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.CreateSnapshotScheduleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a snapshot schedule identified by the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and the schedule (`extId`).
func (api *SnapshotSchedulesApi) DeleteSnapshotScheduleById(fileServerExtId *string, extId *string, args ...map[string]interface{}) (*import3.DeleteSnapshotScheduleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.DeleteSnapshotScheduleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a snapshot schedule identified by the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and the schedule (`extId`).
func (api *SnapshotSchedulesApi) GetSnapshotScheduleById(fileServerExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetSnapshotScheduleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
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

	unmarshalledResp := new(import3.GetSnapshotScheduleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a paginated list of snapshot schedules.
func (api *SnapshotSchedulesApi) ListSnapshotSchedules(fileServerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListSnapshotSchedulesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
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

	unmarshalledResp := new(import3.ListSnapshotSchedulesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update a snapshot schedule identified by the given external identifier by using the provided request body.  Specify a valid identifier of the file server (`fileServerExtId`) and the schedule(`extId`).  It is always recommended to perform a GET request on a resource before requesting an UPDATE call.
func (api *SnapshotSchedulesApi) UpdateSnapshotScheduleById(fileServerExtId *string, extId *string, body *import3.SnapshotSchedule, args ...map[string]interface{}) (*import3.UpdateSnapshotScheduleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/snapshot-schedules/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UpdateSnapshotScheduleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
