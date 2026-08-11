package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	import23 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmantiaffinitypolicies"
	"net/http"
	"net/url"
	"strings"
)

type VmAntiAffinityPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmAntiAffinityPoliciesServiceApi
}

type VmAntiAffinityPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmAntiAffinityPoliciesApi(apiClient *client.ApiClient) *VmAntiAffinityPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmAntiAffinityPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(a.ApiClient)

	return a
}

func NewVmAntiAffinityPoliciesServiceApi(apiClient *client.ApiClient) *VmAntiAffinityPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmAntiAffinityPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesApi) CreateVmAntiAffinityPolicy(body *import6.VmAntiAffinityPolicy, args ...map[string]interface{}) (*import6.CreateVmAntiAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVmAntiAffinityPolicy(context.Background(), &import23.CreateVmAntiAffinityPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesServiceApi) CreateVmAntiAffinityPolicy(ctx context.Context, request *import23.CreateVmAntiAffinityPolicyRequest, args ...map[string]interface{}) (*import6.CreateVmAntiAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.CreateVmAntiAffinityPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the legacy VM-VM anti-affinity policy. The legacy VM-VM anti-affinity policies refer to the anti-affinity policies configured using VM groups through aCLI in Prism Element.
func (api *VmAntiAffinityPoliciesApi) DeleteLegacyVmAntiAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import6.DeleteLegacyVmAntiAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteLegacyVmAntiAffinityPolicyById(context.Background(), &import23.DeleteLegacyVmAntiAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the legacy VM-VM anti-affinity policy. The legacy VM-VM anti-affinity policies refer to the anti-affinity policies configured using VM groups through aCLI in Prism Element.
func (api *VmAntiAffinityPoliciesServiceApi) DeleteLegacyVmAntiAffinityPolicyById(ctx context.Context, request *import23.DeleteLegacyVmAntiAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import6.DeleteLegacyVmAntiAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/legacy-vm-anti-affinity-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.DeleteLegacyVmAntiAffinityPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the requested VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesApi) DeleteVmAntiAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import6.DeleteVmAntiAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmAntiAffinityPolicyById(context.Background(), &import23.DeleteVmAntiAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the requested VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesServiceApi) DeleteVmAntiAffinityPolicyById(ctx context.Context, request *import23.DeleteVmAntiAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import6.DeleteVmAntiAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.DeleteVmAntiAffinityPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetches the VM-VM anti-affinity policy of the provided VM-VM anti-affinity policy external identifier.
func (api *VmAntiAffinityPoliciesApi) GetVmAntiAffinityPolicyById(extId *string, args ...map[string]interface{}) (*import6.GetVmAntiAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmAntiAffinityPolicyById(context.Background(), &import23.GetVmAntiAffinityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the VM-VM anti-affinity policy of the provided VM-VM anti-affinity policy external identifier.
func (api *VmAntiAffinityPoliciesServiceApi) GetVmAntiAffinityPolicyById(ctx context.Context, request *import23.GetVmAntiAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import6.GetVmAntiAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.GetVmAntiAffinityPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List legacy VM-VM anti-affinity policies. The legacy VM-VM anti-affinity policies refer to the anti-affinity policies configured using VM groups through aCLI in Prism Element.
func (api *VmAntiAffinityPoliciesApi) ListLegacyVmAntiAffinityPolicies(page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import6.ListLegacyVmAntiAffinityPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLegacyVmAntiAffinityPolicies(context.Background(), &import23.ListLegacyVmAntiAffinityPoliciesRequest{
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
	}, args...)
}

// List legacy VM-VM anti-affinity policies. The legacy VM-VM anti-affinity policies refer to the anti-affinity policies configured using VM groups through aCLI in Prism Element.
func (api *VmAntiAffinityPoliciesServiceApi) ListLegacyVmAntiAffinityPolicies(ctx context.Context, request *import23.ListLegacyVmAntiAffinityPoliciesRequest, args ...map[string]interface{}) (*import6.ListLegacyVmAntiAffinityPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/legacy-vm-anti-affinity-policies"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.ListLegacyVmAntiAffinityPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List VM-VM anti-affinity policies.
func (api *VmAntiAffinityPoliciesApi) ListVmAntiAffinityPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import6.ListVmAntiAffinityPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmAntiAffinityPolicies(context.Background(), &import23.ListVmAntiAffinityPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List VM-VM anti-affinity policies.
func (api *VmAntiAffinityPoliciesServiceApi) ListVmAntiAffinityPolicies(ctx context.Context, request *import23.ListVmAntiAffinityPoliciesRequest, args ...map[string]interface{}) (*import6.ListVmAntiAffinityPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.ListVmAntiAffinityPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists compliance states of VMs in the VM-VM anti-affinity policy of the provided VM-VM anti-affinity policy external identifier.
func (api *VmAntiAffinityPoliciesApi) ListVmAntiAffinityPolicyVmComplianceStates(vmAntiAffinityPolicyExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import6.ListVmAntiAffinityPolicyVmComplianceStatesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmAntiAffinityPolicyVmComplianceStates(context.Background(), &import23.ListVmAntiAffinityPolicyVmComplianceStatesRequest{
		VmAntiAffinityPolicyExtId: vmAntiAffinityPolicyExtId,
		Page_:                     page_,
		Limit_:                    limit_,
	}, args...)
}

// Lists compliance states of VMs in the VM-VM anti-affinity policy of the provided VM-VM anti-affinity policy external identifier.
func (api *VmAntiAffinityPoliciesServiceApi) ListVmAntiAffinityPolicyVmComplianceStates(ctx context.Context, request *import23.ListVmAntiAffinityPolicyVmComplianceStatesRequest, args ...map[string]interface{}) (*import6.ListVmAntiAffinityPolicyVmComplianceStatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies/{vmAntiAffinityPolicyExtId}/vm-compliance-states"

	// verify the required parameter 'vmAntiAffinityPolicyExtId' is set
	if nil == request.VmAntiAffinityPolicyExtId {
		return nil, client.ReportError("vmAntiAffinityPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmAntiAffinityPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmAntiAffinityPolicyExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.ListVmAntiAffinityPolicyVmComplianceStatesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the requested VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesApi) UpdateVmAntiAffinityPolicyById(extId *string, body *import6.VmAntiAffinityPolicy, args ...map[string]interface{}) (*import6.UpdateVmAntiAffinityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmAntiAffinityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmAntiAffinityPolicyById(context.Background(), &import23.UpdateVmAntiAffinityPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the requested VM-VM anti-affinity policy.
func (api *VmAntiAffinityPoliciesServiceApi) UpdateVmAntiAffinityPolicyById(ctx context.Context, request *import23.UpdateVmAntiAffinityPolicyByIdRequest, args ...map[string]interface{}) (*import6.UpdateVmAntiAffinityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/vm-anti-affinity-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.UpdateVmAntiAffinityPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
