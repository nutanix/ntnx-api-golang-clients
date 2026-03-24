package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	import26 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmstartuppolicies"
	"net/http"
	"net/url"
	"strings"
)

type VmStartupPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmStartupPoliciesServiceApi
}

type VmStartupPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmStartupPoliciesApi(apiClient *client.ApiClient) *VmStartupPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmStartupPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmStartupPoliciesServiceApi(a.ApiClient)

	return a
}

func NewVmStartupPoliciesServiceApi(apiClient *client.ApiClient) *VmStartupPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmStartupPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a VM startup policy.
func (api *VmStartupPoliciesApi) CreateVmStartupPolicy(body *import21.VmStartupPolicy, args ...map[string]interface{}) (*import21.CreateVmStartupPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVmStartupPolicy(context.Background(), &import26.CreateVmStartupPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a VM startup policy.
func (api *VmStartupPoliciesServiceApi) CreateVmStartupPolicy(ctx context.Context, request *import26.CreateVmStartupPolicyRequest, args ...map[string]interface{}) (*import21.CreateVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies"

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

	unmarshalledResp := new(import21.CreateVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the requested VM startup policy.
func (api *VmStartupPoliciesApi) DeleteVmStartupPolicyById(extId *string, args ...map[string]interface{}) (*import21.DeleteVmStartupPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmStartupPolicyById(context.Background(), &import26.DeleteVmStartupPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the requested VM startup policy.
func (api *VmStartupPoliciesServiceApi) DeleteVmStartupPolicyById(ctx context.Context, request *import26.DeleteVmStartupPolicyByIdRequest, args ...map[string]interface{}) (*import21.DeleteVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

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

	unmarshalledResp := new(import21.DeleteVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the VM startup policy of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyById(extId *string, args ...map[string]interface{}) (*import21.GetVmStartupPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmStartupPolicyById(context.Background(), &import26.GetVmStartupPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the VM startup policy of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) GetVmStartupPolicyById(ctx context.Context, request *import26.GetVmStartupPolicyByIdRequest, args ...map[string]interface{}) (*import21.GetVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

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

	unmarshalledResp := new(import21.GetVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyDependencyConflictById(vmStartupPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import21.GetVmStartupPolicyDependencyConflictApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmStartupPolicyDependencyConflictById(context.Background(), &import26.GetVmStartupPolicyDependencyConflictByIdRequest{
		VmStartupPolicyExtId: vmStartupPolicyExtId,
		ExtId:                extId,
	}, args...)
}

// Get Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) GetVmStartupPolicyDependencyConflictById(ctx context.Context, request *import26.GetVmStartupPolicyDependencyConflictByIdRequest, args ...map[string]interface{}) (*import21.GetVmStartupPolicyDependencyConflictApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{extId}"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.GetVmStartupPolicyDependencyConflictApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get Start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyStartConditionConflictById(vmStartupPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import21.GetVmStartupPolicyStartConditionConflictApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmStartupPolicyStartConditionConflictById(context.Background(), &import26.GetVmStartupPolicyStartConditionConflictByIdRequest{
		VmStartupPolicyExtId: vmStartupPolicyExtId,
		ExtId:                extId,
	}, args...)
}

// Get Start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) GetVmStartupPolicyStartConditionConflictById(ctx context.Context, request *import26.GetVmStartupPolicyStartConditionConflictByIdRequest, args ...map[string]interface{}) (*import21.GetVmStartupPolicyStartConditionConflictApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{extId}"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.GetVmStartupPolicyStartConditionConflictApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List VM startup policies.
func (api *VmStartupPoliciesApi) ListVmStartupPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import21.ListVmStartupPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicies(context.Background(), &import26.ListVmStartupPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List VM startup policies.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicies(ctx context.Context, request *import26.ListVmStartupPoliciesRequest, args ...map[string]interface{}) (*import21.ListVmStartupPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies"

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

	unmarshalledResp := new(import21.ListVmStartupPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependee VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflictDependeeVms(vmStartupPolicyExtId *string, dependencyConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictDependeeVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyDependencyConflictDependeeVms(context.Background(), &import26.ListVmStartupPolicyDependencyConflictDependeeVmsRequest{
		VmStartupPolicyExtId:    vmStartupPolicyExtId,
		DependencyConflictExtId: dependencyConflictExtId,
		Page_:                   page_,
		Limit_:                  limit_,
	}, args...)
}

// List dependee VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyDependencyConflictDependeeVms(ctx context.Context, request *import26.ListVmStartupPolicyDependencyConflictDependeeVmsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictDependeeVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{dependencyConflictExtId}/dependee-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'dependencyConflictExtId' is set
	if nil == request.DependencyConflictExtId {
		return nil, client.ReportError("dependencyConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"dependencyConflictExtId"+"}", url.PathEscape(client.ParameterToString(*request.DependencyConflictExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyDependencyConflictDependeeVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependent VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflictDependentVms(vmStartupPolicyExtId *string, dependencyConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictDependentVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyDependencyConflictDependentVms(context.Background(), &import26.ListVmStartupPolicyDependencyConflictDependentVmsRequest{
		VmStartupPolicyExtId:    vmStartupPolicyExtId,
		DependencyConflictExtId: dependencyConflictExtId,
		Page_:                   page_,
		Limit_:                  limit_,
	}, args...)
}

// List dependent VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyDependencyConflictDependentVms(ctx context.Context, request *import26.ListVmStartupPolicyDependencyConflictDependentVmsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictDependentVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{dependencyConflictExtId}/dependent-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'dependencyConflictExtId' is set
	if nil == request.DependencyConflictExtId {
		return nil, client.ReportError("dependencyConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"dependencyConflictExtId"+"}", url.PathEscape(client.ParameterToString(*request.DependencyConflictExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyDependencyConflictDependentVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Dependency Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflicts(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyDependencyConflicts(context.Background(), &import26.ListVmStartupPolicyDependencyConflictsRequest{
		VmStartupPolicyExtId: vmStartupPolicyExtId,
		Page_:                page_,
		Limit_:               limit_,
	}, args...)
}

// List Dependency Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyDependencyConflicts(ctx context.Context, request *import26.ListVmStartupPolicyDependencyConflictsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyDependencyConflictsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyDependencyConflictsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependee VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflictDependeeVms(vmStartupPolicyExtId *string, startConditionConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictDependeeVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyStartConditionConflictDependeeVms(context.Background(), &import26.ListVmStartupPolicyStartConditionConflictDependeeVmsRequest{
		VmStartupPolicyExtId:        vmStartupPolicyExtId,
		StartConditionConflictExtId: startConditionConflictExtId,
		Page_:                       page_,
		Limit_:                      limit_,
	}, args...)
}

// List dependee VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyStartConditionConflictDependeeVms(ctx context.Context, request *import26.ListVmStartupPolicyStartConditionConflictDependeeVmsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictDependeeVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{startConditionConflictExtId}/dependee-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'startConditionConflictExtId' is set
	if nil == request.StartConditionConflictExtId {
		return nil, client.ReportError("startConditionConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"startConditionConflictExtId"+"}", url.PathEscape(client.ParameterToString(*request.StartConditionConflictExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyStartConditionConflictDependeeVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependent VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflictDependentVms(vmStartupPolicyExtId *string, startConditionConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictDependentVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyStartConditionConflictDependentVms(context.Background(), &import26.ListVmStartupPolicyStartConditionConflictDependentVmsRequest{
		VmStartupPolicyExtId:        vmStartupPolicyExtId,
		StartConditionConflictExtId: startConditionConflictExtId,
		Page_:                       page_,
		Limit_:                      limit_,
	}, args...)
}

// List dependent VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyStartConditionConflictDependentVms(ctx context.Context, request *import26.ListVmStartupPolicyStartConditionConflictDependentVmsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictDependentVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{startConditionConflictExtId}/dependent-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'startConditionConflictExtId' is set
	if nil == request.StartConditionConflictExtId {
		return nil, client.ReportError("startConditionConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"startConditionConflictExtId"+"}", url.PathEscape(client.ParameterToString(*request.StartConditionConflictExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyStartConditionConflictDependentVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Start condition Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflicts(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyStartConditionConflicts(context.Background(), &import26.ListVmStartupPolicyStartConditionConflictsRequest{
		VmStartupPolicyExtId: vmStartupPolicyExtId,
		Page_:                page_,
		Limit_:               limit_,
	}, args...)
}

// List Start condition Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyStartConditionConflicts(ctx context.Context, request *import26.ListVmStartupPolicyStartConditionConflictsRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyStartConditionConflictsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyStartConditionConflictsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List VM compliances of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyVmComplianceStates(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmStartupPolicyVmComplianceStatesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmStartupPolicyVmComplianceStates(context.Background(), &import26.ListVmStartupPolicyVmComplianceStatesRequest{
		VmStartupPolicyExtId: vmStartupPolicyExtId,
		Page_:                page_,
		Limit_:               limit_,
	}, args...)
}

// List VM compliances of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesServiceApi) ListVmStartupPolicyVmComplianceStates(ctx context.Context, request *import26.ListVmStartupPolicyVmComplianceStatesRequest, args ...map[string]interface{}) (*import21.ListVmStartupPolicyVmComplianceStatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/vm-compliance-states"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == request.VmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmStartupPolicyVmComplianceStatesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the requested VM startup policy.
func (api *VmStartupPoliciesApi) UpdateVmStartupPolicyById(extId *string, body *import21.VmStartupPolicy, args ...map[string]interface{}) (*import21.UpdateVmStartupPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmStartupPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmStartupPolicyById(context.Background(), &import26.UpdateVmStartupPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the requested VM startup policy.
func (api *VmStartupPoliciesServiceApi) UpdateVmStartupPolicyById(ctx context.Context, request *import26.UpdateVmStartupPolicyByIdRequest, args ...map[string]interface{}) (*import21.UpdateVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

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

	unmarshalledResp := new(import21.UpdateVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
