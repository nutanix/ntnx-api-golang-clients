package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/operations"
	import5 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/request/recoveryplanactions"
	"net/http"
	"net/url"
	"strings"
)

type RecoveryPlanActionsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RecoveryPlanActionsServiceApi
}

type RecoveryPlanActionsServiceApi struct {
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

	a.ServiceClient = NewRecoveryPlanActionsServiceApi(a.ApiClient)

	return a
}

func NewRecoveryPlanActionsServiceApi(apiClient *client.ApiClient) *RecoveryPlanActionsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPlanActionsServiceApi{
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
func (api *RecoveryPlanActionsApi) CleanupRecoveryPlanResources(recoveryPlanExtId *string, args ...map[string]interface{}) (*import4.CleanupRecoveryPlanResourcesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlanActionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CleanupRecoveryPlanResources(context.Background(), &import5.CleanupRecoveryPlanResourcesRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
	}, args...)
}

// Delete entities recovered in the last test failover action on recovery plan.<br> #### Task Completion Details <br> The external identifier for the cleaned recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsServiceApi) CleanupRecoveryPlanResources(ctx context.Context, request *import5.CleanupRecoveryPlanResourcesRequest, args ...map[string]interface{}) (*import4.CleanupRecoveryPlanResourcesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/operations/recovery-plans/{recoveryPlanExtId}/$actions/clean-up-resources"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.CleanupRecoveryPlanResourcesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform a planned failover on the recovery plan extID. VMs are powered off on the source before migrating it to the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) PlannedFailoverRecoveryPlan(recoveryPlanExtId *string, body *import4.PlannedFailoverSpec, args ...map[string]interface{}) (*import4.PlannedFailoverRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlanActionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PlannedFailoverRecoveryPlan(context.Background(), &import5.PlannedFailoverRecoveryPlanRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Perform a planned failover on the recovery plan extID. VMs are powered off on the source before migrating it to the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsServiceApi) PlannedFailoverRecoveryPlan(ctx context.Context, request *import5.PlannedFailoverRecoveryPlanRequest, args ...map[string]interface{}) (*import4.PlannedFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/operations/recovery-plans/{recoveryPlanExtId}/$actions/planned-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import4.PlannedFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform a test failover on the recovery plan extId.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) TestFailoverRecoveryPlan(recoveryPlanExtId *string, body *import4.TestFailoverSpec, args ...map[string]interface{}) (*import4.TestFailoverRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlanActionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.TestFailoverRecoveryPlan(context.Background(), &import5.TestFailoverRecoveryPlanRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Perform a test failover on the recovery plan extId.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsServiceApi) TestFailoverRecoveryPlan(ctx context.Context, request *import5.TestFailoverRecoveryPlanRequest, args ...map[string]interface{}) (*import4.TestFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/operations/recovery-plans/{recoveryPlanExtId}/$actions/test-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import4.TestFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Perform an unplanned failover on the recovery plan extId. Restore the entity from the recovery points on the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) UnplannedFailoverRecoveryPlan(recoveryPlanExtId *string, body *import4.UnplannedFailoverSpec, args ...map[string]interface{}) (*import4.UnplannedFailoverRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlanActionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnplannedFailoverRecoveryPlan(context.Background(), &import5.UnplannedFailoverRecoveryPlanRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Perform an unplanned failover on the recovery plan extId. Restore the entity from the recovery points on the target domain manager.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsServiceApi) UnplannedFailoverRecoveryPlan(ctx context.Context, request *import5.UnplannedFailoverRecoveryPlanRequest, args ...map[string]interface{}) (*import4.UnplannedFailoverRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/operations/recovery-plans/{recoveryPlanExtId}/$actions/unplanned-failover"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import4.UnplannedFailoverRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Performs the validation of the recovery plan. The validation includes checks for the presence of entities, networks, categories, etc. referenced in the recovery plan.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsApi) ValidateRecoveryPlan(recoveryPlanExtId *string, body *import4.BaseRecoveryPlanActionSpec, args ...map[string]interface{}) (*import4.ValidateRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlanActionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ValidateRecoveryPlan(context.Background(), &import5.ValidateRecoveryPlanRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Performs the validation of the recovery plan. The validation includes checks for the presence of entities, networks, categories, etc. referenced in the recovery plan.<br> #### Task Completion Details <br> The external identifier for the created recovery plan job is available in the task completion details under the key `recoveryPlanJobExtId`, as well as in the `entitiesAffected` section.
func (api *RecoveryPlanActionsServiceApi) ValidateRecoveryPlan(ctx context.Context, request *import5.ValidateRecoveryPlanRequest, args ...map[string]interface{}) (*import4.ValidateRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/operations/recovery-plans/{recoveryPlanExtId}/$actions/validate"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import4.ValidateRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
