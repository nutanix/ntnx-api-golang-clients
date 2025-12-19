package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	"net/http"
	"net/url"
	"strings"
)

type VmStartupPoliciesApi struct {
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

	return a
}

// Creates a VM startup policy.
func (api *VmStartupPoliciesApi) CreateVmStartupPolicy(body *import11.VmStartupPolicy, args ...map[string]interface{}) (*import11.CreateVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies"

	// verify the required parameter 'body' is set
	if nil == body {
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import11.CreateVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the requested VM startup policy.
func (api *VmStartupPoliciesApi) DeleteVmStartupPolicyById(extId *string, args ...map[string]interface{}) (*import11.DeleteVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

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

	unmarshalledResp := new(import11.DeleteVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the VM startup policy of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyById(extId *string, args ...map[string]interface{}) (*import11.GetVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

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

	unmarshalledResp := new(import11.GetVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyDependencyConflictById(vmStartupPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import11.GetVmStartupPolicyDependencyConflictApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{extId}"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import11.GetVmStartupPolicyDependencyConflictApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get Start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) GetVmStartupPolicyStartConditionConflictById(vmStartupPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import11.GetVmStartupPolicyStartConditionConflictApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{extId}"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import11.GetVmStartupPolicyStartConditionConflictApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List VM startup policies.
func (api *VmStartupPoliciesApi) ListVmStartupPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import11.ListVmStartupPoliciesApiResponse, error) {
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

	unmarshalledResp := new(import11.ListVmStartupPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependee VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflictDependeeVms(vmStartupPolicyExtId *string, dependencyConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyDependencyConflictDependeeVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{dependencyConflictExtId}/dependee-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'dependencyConflictExtId' is set
	if nil == dependencyConflictExtId {
		return nil, client.ReportError("dependencyConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"dependencyConflictExtId"+"}", url.PathEscape(client.ParameterToString(*dependencyConflictExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyDependencyConflictDependeeVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependent VMs of a Dependency conflict for the provided Dependency conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflictDependentVms(vmStartupPolicyExtId *string, dependencyConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyDependencyConflictDependentVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts/{dependencyConflictExtId}/dependent-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'dependencyConflictExtId' is set
	if nil == dependencyConflictExtId {
		return nil, client.ReportError("dependencyConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"dependencyConflictExtId"+"}", url.PathEscape(client.ParameterToString(*dependencyConflictExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyDependencyConflictDependentVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Dependency Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyDependencyConflicts(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyDependencyConflictsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/dependency-conflicts"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyDependencyConflictsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependee VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflictDependeeVms(vmStartupPolicyExtId *string, startConditionConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyStartConditionConflictDependeeVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{startConditionConflictExtId}/dependee-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'startConditionConflictExtId' is set
	if nil == startConditionConflictExtId {
		return nil, client.ReportError("startConditionConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"startConditionConflictExtId"+"}", url.PathEscape(client.ParameterToString(*startConditionConflictExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyStartConditionConflictDependeeVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List dependent VMs of a start condition conflict for the provided start condition conflict external identifier and VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflictDependentVms(vmStartupPolicyExtId *string, startConditionConflictExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyStartConditionConflictDependentVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts/{startConditionConflictExtId}/dependent-vms"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'startConditionConflictExtId' is set
	if nil == startConditionConflictExtId {
		return nil, client.ReportError("startConditionConflictExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"startConditionConflictExtId"+"}", url.PathEscape(client.ParameterToString(*startConditionConflictExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyStartConditionConflictDependentVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Start condition Conflicts of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyStartConditionConflicts(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyStartConditionConflictsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/start-condition-conflicts"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyStartConditionConflictsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List VM compliances of the provided VM startup policy external identifier.
func (api *VmStartupPoliciesApi) ListVmStartupPolicyVmComplianceStates(vmStartupPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import11.ListVmStartupPolicyVmComplianceStatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{vmStartupPolicyExtId}/vm-compliance-states"

	// verify the required parameter 'vmStartupPolicyExtId' is set
	if nil == vmStartupPolicyExtId {
		return nil, client.ReportError("vmStartupPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmStartupPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*vmStartupPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import11.ListVmStartupPolicyVmComplianceStatesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the requested VM startup policy.
func (api *VmStartupPoliciesApi) UpdateVmStartupPolicyById(extId *string, body *import11.VmStartupPolicy, args ...map[string]interface{}) (*import11.UpdateVmStartupPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-startup-policies/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import11.UpdateVmStartupPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
