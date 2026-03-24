package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/request/recoveryplans"
	"net/http"
	"net/url"
	"strings"
)

type RecoveryPlansApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RecoveryPlansServiceApi
}

type RecoveryPlansServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRecoveryPlansApi(apiClient *client.ApiClient) *RecoveryPlansApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPlansApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRecoveryPlansServiceApi(a.ApiClient)

	return a
}

func NewRecoveryPlansServiceApi(apiClient *client.ApiClient) *RecoveryPlansServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPlansServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a data services IP mapping.
func (api *RecoveryPlansApi) CreateDataServicesIpMapping(recoveryPlanExtId *string, body *import1.DataServicesIpMapping, args ...map[string]interface{}) (*import1.CreateDataServicesIpMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDataServicesIpMapping(context.Background(), &import4.CreateDataServicesIpMappingRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Create a data services IP mapping.
func (api *RecoveryPlansServiceApi) CreateDataServicesIpMapping(ctx context.Context, request *import4.CreateDataServicesIpMappingRequest, args ...map[string]interface{}) (*import1.CreateDataServicesIpMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/data-services-ip-mappings"

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

	unmarshalledResp := new(import1.CreateDataServicesIpMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create a network mapping.
func (api *RecoveryPlansApi) CreateNetworkMapping(recoveryPlanExtId *string, body *import1.NetworkMapping, args ...map[string]interface{}) (*import1.CreateNetworkMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateNetworkMapping(context.Background(), &import4.CreateNetworkMappingRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Create a network mapping.
func (api *RecoveryPlansServiceApi) CreateNetworkMapping(ctx context.Context, request *import4.CreateNetworkMappingRequest, args ...map[string]interface{}) (*import1.CreateNetworkMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/network-mappings"

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

	unmarshalledResp := new(import1.CreateNetworkMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create a recovery plan.
func (api *RecoveryPlansApi) CreateRecoveryPlan(body *import1.RecoveryPlan, args ...map[string]interface{}) (*import1.CreateRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRecoveryPlan(context.Background(), &import4.CreateRecoveryPlanRequest{
		Body: body,
	}, args...)
}

// Create a recovery plan.
func (api *RecoveryPlansServiceApi) CreateRecoveryPlan(ctx context.Context, request *import4.CreateRecoveryPlanRequest, args ...map[string]interface{}) (*import1.CreateRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

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

	unmarshalledResp := new(import1.CreateRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create a custom recovery setting for VM or Volume Groups.
func (api *RecoveryPlansApi) CreateRecoverySetting(recoveryPlanExtId *string, body *import1.RecoverySetting, args ...map[string]interface{}) (*import1.CreateRecoverySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRecoverySetting(context.Background(), &import4.CreateRecoverySettingRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Create a custom recovery setting for VM or Volume Groups.
func (api *RecoveryPlansServiceApi) CreateRecoverySetting(ctx context.Context, request *import4.CreateRecoverySettingRequest, args ...map[string]interface{}) (*import1.CreateRecoverySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/recovery-settings"

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

	unmarshalledResp := new(import1.CreateRecoverySettingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create a Recovery stage.
func (api *RecoveryPlansApi) CreateRecoveryStage(recoveryPlanExtId *string, body *import1.RecoveryStage, args ...map[string]interface{}) (*import1.CreateRecoveryStageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRecoveryStage(context.Background(), &import4.CreateRecoveryStageRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Body:              body,
	}, args...)
}

// Create a Recovery stage.
func (api *RecoveryPlansServiceApi) CreateRecoveryStage(ctx context.Context, request *import4.CreateRecoveryStageRequest, args ...map[string]interface{}) (*import1.CreateRecoveryStageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/stages"

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

	unmarshalledResp := new(import1.CreateRecoveryStageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the data services IP mapping identified by {extId}.
func (api *RecoveryPlansApi) DeleteDataServicesIpMappingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteDataServicesIpMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDataServicesIpMappingById(context.Background(), &import4.DeleteDataServicesIpMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Delete the data services IP mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) DeleteDataServicesIpMappingById(ctx context.Context, request *import4.DeleteDataServicesIpMappingByIdRequest, args ...map[string]interface{}) (*import1.DeleteDataServicesIpMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/data-services-ip-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteDataServicesIpMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the network mapping identified by {extId}.
func (api *RecoveryPlansApi) DeleteNetworkMappingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteNetworkMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteNetworkMappingById(context.Background(), &import4.DeleteNetworkMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Delete the network mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) DeleteNetworkMappingById(ctx context.Context, request *import4.DeleteNetworkMappingByIdRequest, args ...map[string]interface{}) (*import1.DeleteNetworkMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/network-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteNetworkMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the recovery plan identified by {extId}.
func (api *RecoveryPlansApi) DeleteRecoveryPlanById(extId *string, args ...map[string]interface{}) (*import1.DeleteRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRecoveryPlanById(context.Background(), &import4.DeleteRecoveryPlanByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the recovery plan identified by {extId}.
func (api *RecoveryPlansServiceApi) DeleteRecoveryPlanById(ctx context.Context, request *import4.DeleteRecoveryPlanByIdRequest, args ...map[string]interface{}) (*import1.DeleteRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{extId}"

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

	unmarshalledResp := new(import1.DeleteRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the recovery setting identified by {extId}.
func (api *RecoveryPlansApi) DeleteRecoverySettingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteRecoverySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRecoverySettingById(context.Background(), &import4.DeleteRecoverySettingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Delete the recovery setting identified by {extId}.
func (api *RecoveryPlansServiceApi) DeleteRecoverySettingById(ctx context.Context, request *import4.DeleteRecoverySettingByIdRequest, args ...map[string]interface{}) (*import1.DeleteRecoverySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/recovery-settings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteRecoverySettingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the Recovery stage identified by {extId}.
func (api *RecoveryPlansApi) DeleteRecoveryStageById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteRecoveryStageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRecoveryStageById(context.Background(), &import4.DeleteRecoveryStageByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Delete the Recovery stage identified by {extId}.
func (api *RecoveryPlansServiceApi) DeleteRecoveryStageById(ctx context.Context, request *import4.DeleteRecoveryStageByIdRequest, args ...map[string]interface{}) (*import1.DeleteRecoveryStageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/stages/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteRecoveryStageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the Data services IP mapping identified by {extId}.
func (api *RecoveryPlansApi) GetDataServicesIpMappingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetDataServicesIpMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDataServicesIpMappingById(context.Background(), &import4.GetDataServicesIpMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Get the Data services IP mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) GetDataServicesIpMappingById(ctx context.Context, request *import4.GetDataServicesIpMappingByIdRequest, args ...map[string]interface{}) (*import1.GetDataServicesIpMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/data-services-ip-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetDataServicesIpMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the network mapping identified by {extId}.
func (api *RecoveryPlansApi) GetNetworkMappingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetNetworkMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNetworkMappingById(context.Background(), &import4.GetNetworkMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Get the network mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) GetNetworkMappingById(ctx context.Context, request *import4.GetNetworkMappingByIdRequest, args ...map[string]interface{}) (*import1.GetNetworkMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/network-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetNetworkMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the recovery plan identified by {extId}.
func (api *RecoveryPlansApi) GetRecoveryPlanById(extId *string, args ...map[string]interface{}) (*import1.GetRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecoveryPlanById(context.Background(), &import4.GetRecoveryPlanByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get the recovery plan identified by {extId}.
func (api *RecoveryPlansServiceApi) GetRecoveryPlanById(ctx context.Context, request *import4.GetRecoveryPlanByIdRequest, args ...map[string]interface{}) (*import1.GetRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{extId}"

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

	unmarshalledResp := new(import1.GetRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the custom recovery setting identified by {extId}.
func (api *RecoveryPlansApi) GetRecoverySettingById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetRecoverySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecoverySettingById(context.Background(), &import4.GetRecoverySettingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Get the custom recovery setting identified by {extId}.
func (api *RecoveryPlansServiceApi) GetRecoverySettingById(ctx context.Context, request *import4.GetRecoverySettingByIdRequest, args ...map[string]interface{}) (*import1.GetRecoverySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/recovery-settings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetRecoverySettingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the Recovery stage identified by {extId}.
func (api *RecoveryPlansApi) GetRecoveryStageById(recoveryPlanExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetRecoveryStageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecoveryStageById(context.Background(), &import4.GetRecoveryStageByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
	}, args...)
}

// Get the Recovery stage identified by {extId}.
func (api *RecoveryPlansServiceApi) GetRecoveryStageById(ctx context.Context, request *import4.GetRecoveryStageByIdRequest, args ...map[string]interface{}) (*import1.GetRecoveryStageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/stages/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetRecoveryStageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List data services IP mappings.
func (api *RecoveryPlansApi) ListDataServicesIpMappings(recoveryPlanExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDataServicesIpMappingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDataServicesIpMappings(context.Background(), &import4.ListDataServicesIpMappingsRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Page_:             page_,
		Limit_:            limit_,
		Filter_:           filter_,
		Orderby_:          orderby_,
		Select_:           select_,
	}, args...)
}

// List data services IP mappings.
func (api *RecoveryPlansServiceApi) ListDataServicesIpMappings(ctx context.Context, request *import4.ListDataServicesIpMappingsRequest, args ...map[string]interface{}) (*import1.ListDataServicesIpMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/data-services-ip-mappings"

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

	unmarshalledResp := new(import1.ListDataServicesIpMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List network mappings.
func (api *RecoveryPlansApi) ListNetworkMappings(recoveryPlanExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListNetworkMappingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNetworkMappings(context.Background(), &import4.ListNetworkMappingsRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Page_:             page_,
		Limit_:            limit_,
		Filter_:           filter_,
		Orderby_:          orderby_,
		Select_:           select_,
	}, args...)
}

// List network mappings.
func (api *RecoveryPlansServiceApi) ListNetworkMappings(ctx context.Context, request *import4.ListNetworkMappingsRequest, args ...map[string]interface{}) (*import1.ListNetworkMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/network-mappings"

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

	unmarshalledResp := new(import1.ListNetworkMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List recovery plans.
func (api *RecoveryPlansApi) ListRecoveryPlans(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRecoveryPlansApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecoveryPlans(context.Background(), &import4.ListRecoveryPlansRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List recovery plans.
func (api *RecoveryPlansServiceApi) ListRecoveryPlans(ctx context.Context, request *import4.ListRecoveryPlansRequest, args ...map[string]interface{}) (*import1.ListRecoveryPlansApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans"

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

	unmarshalledResp := new(import1.ListRecoveryPlansApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List custom recovery settings.
func (api *RecoveryPlansApi) ListRecoverySettings(recoveryPlanExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRecoverySettingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecoverySettings(context.Background(), &import4.ListRecoverySettingsRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Page_:             page_,
		Limit_:            limit_,
		Filter_:           filter_,
		Orderby_:          orderby_,
		Select_:           select_,
	}, args...)
}

// List custom recovery settings.
func (api *RecoveryPlansServiceApi) ListRecoverySettings(ctx context.Context, request *import4.ListRecoverySettingsRequest, args ...map[string]interface{}) (*import1.ListRecoverySettingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/recovery-settings"

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

	unmarshalledResp := new(import1.ListRecoverySettingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Recovery stages.
func (api *RecoveryPlansApi) ListRecoveryStages(recoveryPlanExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRecoveryStagesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecoveryStages(context.Background(), &import4.ListRecoveryStagesRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		Page_:             page_,
		Limit_:            limit_,
		Filter_:           filter_,
		Orderby_:          orderby_,
		Select_:           select_,
	}, args...)
}

// List Recovery stages.
func (api *RecoveryPlansServiceApi) ListRecoveryStages(ctx context.Context, request *import4.ListRecoveryStagesRequest, args ...map[string]interface{}) (*import1.ListRecoveryStagesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/stages"

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

	unmarshalledResp := new(import1.ListRecoveryStagesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the data services IP mapping identified by {extId}.
func (api *RecoveryPlansApi) UpdateDataServicesIpMappingById(recoveryPlanExtId *string, extId *string, body *import1.DataServicesIpMapping, args ...map[string]interface{}) (*import1.UpdateDataServicesIpMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDataServicesIpMappingById(context.Background(), &import4.UpdateDataServicesIpMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
		Body:              body,
	}, args...)
}

// Update the data services IP mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) UpdateDataServicesIpMappingById(ctx context.Context, request *import4.UpdateDataServicesIpMappingByIdRequest, args ...map[string]interface{}) (*import1.UpdateDataServicesIpMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/data-services-ip-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateDataServicesIpMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the network mapping identified by {extId}.
func (api *RecoveryPlansApi) UpdateNetworkMappingById(recoveryPlanExtId *string, extId *string, body *import1.NetworkMapping, args ...map[string]interface{}) (*import1.UpdateNetworkMappingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateNetworkMappingById(context.Background(), &import4.UpdateNetworkMappingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
		Body:              body,
	}, args...)
}

// Update the network mapping identified by {extId}.
func (api *RecoveryPlansServiceApi) UpdateNetworkMappingById(ctx context.Context, request *import4.UpdateNetworkMappingByIdRequest, args ...map[string]interface{}) (*import1.UpdateNetworkMappingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/network-mappings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateNetworkMappingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the recovery plan identified by {extId}.
func (api *RecoveryPlansApi) UpdateRecoveryPlanById(extId *string, body *import1.RecoveryPlan, args ...map[string]interface{}) (*import1.UpdateRecoveryPlanApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRecoveryPlanById(context.Background(), &import4.UpdateRecoveryPlanByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the recovery plan identified by {extId}.
func (api *RecoveryPlansServiceApi) UpdateRecoveryPlanById(ctx context.Context, request *import4.UpdateRecoveryPlanByIdRequest, args ...map[string]interface{}) (*import1.UpdateRecoveryPlanApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateRecoveryPlanApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the custom recovery setting identified by {extId}.
func (api *RecoveryPlansApi) UpdateRecoverySettingById(recoveryPlanExtId *string, extId *string, body *import1.RecoverySetting, args ...map[string]interface{}) (*import1.UpdateRecoverySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRecoverySettingById(context.Background(), &import4.UpdateRecoverySettingByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
		Body:              body,
	}, args...)
}

// Update the custom recovery setting identified by {extId}.
func (api *RecoveryPlansServiceApi) UpdateRecoverySettingById(ctx context.Context, request *import4.UpdateRecoverySettingByIdRequest, args ...map[string]interface{}) (*import1.UpdateRecoverySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/recovery-settings/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateRecoverySettingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the Recovery stage identified by {extId}.
func (api *RecoveryPlansApi) UpdateRecoveryStageById(recoveryPlanExtId *string, extId *string, body *import1.RecoveryStage, args ...map[string]interface{}) (*import1.UpdateRecoveryStageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPlansServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRecoveryStageById(context.Background(), &import4.UpdateRecoveryStageByIdRequest{
		RecoveryPlanExtId: recoveryPlanExtId,
		ExtId:             extId,
		Body:              body,
	}, args...)
}

// Update the Recovery stage identified by {extId}.
func (api *RecoveryPlansServiceApi) UpdateRecoveryStageById(ctx context.Context, request *import4.UpdateRecoveryStageByIdRequest, args ...map[string]interface{}) (*import1.UpdateRecoveryStageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/recovery-plans/{recoveryPlanExtId}/stages/{extId}"

	// verify the required parameter 'recoveryPlanExtId' is set
	if nil == request.RecoveryPlanExtId {
		return nil, client.ReportError("recoveryPlanExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPlanExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPlanExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateRecoveryStageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
