package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	import24 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmhostaffinitypolicies"
	"net/http"
	"net/url"
	"strings"
)

type VmHostAffinityPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmHostAffinityPoliciesServiceApi
}

type VmHostAffinityPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmHostAffinityPoliciesApi(apiClient *client.ApiClient) *VmHostAffinityPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmHostAffinityPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmHostAffinityPoliciesServiceApi(a.ApiClient)

	return a
}

func NewVmHostAffinityPoliciesServiceApi(apiClient *client.ApiClient) *VmHostAffinityPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmHostAffinityPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new VM-host affinity policy with the provided configuration.
func (api *VmHostAffinityPoliciesApi) CreateVmHostAffinityPolicy(body *import21.VmHostAffinityPolicy, args ...map[string]interface{}) (*import21.CreateVmHostAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVmHostAffinityPolicy(context.Background(), &import24.CreateVmHostAffinityPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a new VM-host affinity policy with the provided configuration.
func (api *VmHostAffinityPoliciesServiceApi) CreateVmHostAffinityPolicy(ctx context.Context, request *import24.CreateVmHostAffinityPolicyRequest, args ...map[string]interface{}) (*import21.CreateVmHostAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies"

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

	unmarshalledResp := new(import21.CreateVmHostAffinityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the VM-host affinity policy with the given external identifier.
func (api *VmHostAffinityPoliciesApi) DeleteVmHostAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import21.DeleteVmHostAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmHostAffinityPolicyById(context.Background(), &import24.DeleteVmHostAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the VM-host affinity policy with the given external identifier.
func (api *VmHostAffinityPoliciesServiceApi) DeleteVmHostAffinityPolicyById(ctx context.Context, request *import24.DeleteVmHostAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import21.DeleteVmHostAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies/{extId}"

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

	unmarshalledResp := new(import21.DeleteVmHostAffinityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the VM-host affinity policy configuration for the provided VM-host affinity policy external identifier.
func (api *VmHostAffinityPoliciesApi) GetVmHostAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import21.GetVmHostAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmHostAffinityPolicyById(context.Background(), &import24.GetVmHostAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the VM-host affinity policy configuration for the provided VM-host affinity policy external identifier.
func (api *VmHostAffinityPoliciesServiceApi) GetVmHostAffinityPolicyById(ctx context.Context, request *import24.GetVmHostAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import21.GetVmHostAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies/{extId}"

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

	unmarshalledResp := new(import21.GetVmHostAffinityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the VM-Host Affinity policies.
func (api *VmHostAffinityPoliciesApi) ListVmHostAffinityPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import21.ListVmHostAffinityPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmHostAffinityPolicies(context.Background(), &import24.ListVmHostAffinityPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// Lists all the VM-Host Affinity policies.
func (api *VmHostAffinityPoliciesServiceApi) ListVmHostAffinityPolicies(ctx context.Context, request *import24.ListVmHostAffinityPoliciesRequest, args ...map[string]interface{}) (*import21.ListVmHostAffinityPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies"

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

	unmarshalledResp := new(import21.ListVmHostAffinityPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the compliance state of all the VMs associated with the VM-host affinity policy.
func (api *VmHostAffinityPoliciesApi) ListVmHostAffinityPolicyVmComplianceStates(vmHostAffinityPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import21.ListVmHostAffinityPolicyVmComplianceStatesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmHostAffinityPolicyVmComplianceStates(context.Background(), &import24.ListVmHostAffinityPolicyVmComplianceStatesRequest{
		VmHostAffinityPolicyExtId: vmHostAffinityPolicyExtId,
		Page_:                     page_,
		Limit_:                    limit_,
	}, args...)
}

// Lists the compliance state of all the VMs associated with the VM-host affinity policy.
func (api *VmHostAffinityPoliciesServiceApi) ListVmHostAffinityPolicyVmComplianceStates(ctx context.Context, request *import24.ListVmHostAffinityPolicyVmComplianceStatesRequest, args ...map[string]interface{}) (*import21.ListVmHostAffinityPolicyVmComplianceStatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies/{vmHostAffinityPolicyExtId}/vm-compliance-states"

	// verify the required parameter 'vmHostAffinityPolicyExtId' is set
	if nil == request.VmHostAffinityPolicyExtId {
		return nil, client.ReportError("vmHostAffinityPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmHostAffinityPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmHostAffinityPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import21.ListVmHostAffinityPolicyVmComplianceStatesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Re-enforces the VM-host affinity policy on all the VMs associated with the policy.
func (api *VmHostAffinityPoliciesApi) ReEnforceVmHostAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import21.ReEnforceVmHostAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReEnforceVmHostAffinityPolicyById(context.Background(), &import24.ReEnforceVmHostAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Re-enforces the VM-host affinity policy on all the VMs associated with the policy.
func (api *VmHostAffinityPoliciesServiceApi) ReEnforceVmHostAffinityPolicyById(ctx context.Context, request *import24.ReEnforceVmHostAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import21.ReEnforceVmHostAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies/{extId}/$actions/re-enforce"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import21.ReEnforceVmHostAffinityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration of the VM-host affinity policy with the given external identifier.
func (api *VmHostAffinityPoliciesApi) UpdateVmHostAffinityPolicyById(extId *string, body *import21.VmHostAffinityPolicy, args ...map[string]interface{}) (*import21.UpdateVmHostAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmHostAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmHostAffinityPolicyById(context.Background(), &import24.UpdateVmHostAffinityPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the configuration of the VM-host affinity policy with the given external identifier.
func (api *VmHostAffinityPoliciesServiceApi) UpdateVmHostAffinityPolicyById(ctx context.Context, request *import24.UpdateVmHostAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import21.UpdateVmHostAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/policies/vm-host-affinity-policies/{extId}"

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

	unmarshalledResp := new(import21.UpdateVmHostAffinityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
