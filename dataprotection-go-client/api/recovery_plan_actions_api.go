package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/operations"
	"net/http"
	"net/url"
	"strings"
)

type RecoveryPlanActionsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRecoveryPlanActionsApi(apiClient *client.ApiClient) *RecoveryPlanActionsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPlanActionsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Delete entities recovered in the last test failover action on recovery plan.<br> #### Task Completion Details <br> The external identifier for the cleaned recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) CleanupRecoveryPlanResources(recoveryPlanExtId *string, args ...map[string]interface{}) (*import2.CleanupRecoveryPlanResourcesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/clean-up-resources"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == recoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*recoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.CleanupRecoveryPlanResourcesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform a planned failover on the recovery plan extID. VMs are powered off on the source before migrating it to the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) PlannedFailoverRecoveryPlan(recoveryPlanExtId *string, body *import2.PlannedFailoverSpec, args ...map[string]interface{}) (*import2.PlannedFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/planned-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == recoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*recoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import2.PlannedFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform a test failover on the recovery plan extId.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) TestFailoverRecoveryPlan(recoveryPlanExtId *string, body *import2.TestFailoverSpec, args ...map[string]interface{}) (*import2.TestFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/test-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == recoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*recoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import2.TestFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform an unplanned failover on the recovery plan extId. Restore the entity from the recovery points on the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) UnplannedFailoverRecoveryPlan(recoveryPlanExtId *string, body *import2.UnplannedFailoverSpec, args ...map[string]interface{}) (*import2.UnplannedFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/unplanned-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == recoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*recoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import2.UnplannedFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Performs the validation of the recovery plan. The validation includes checks for the presence of entities, networks, categories, etc. referenced in the recovery plan.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) ValidateRecoveryPlan(recoveryPlanExtId *string, body *import2.BaseRecoveryPlanActionSpec, args ...map[string]interface{}) (*import2.ValidateRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/validate"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == recoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*recoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import2.ValidateRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
