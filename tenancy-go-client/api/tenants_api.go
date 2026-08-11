package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/request/tenants"
	"net/http"
	"net/url"
	"strings"
)

type TenantsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *TenantsServiceApi
}

type TenantsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewTenantsApi(apiClient *client.ApiClient) *TenantsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TenantsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewTenantsServiceApi(a.ApiClient)

	return a
}

func NewTenantsServiceApi(apiClient *client.ApiClient) *TenantsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TenantsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Add a new domain to the list of allowed domains.
func (api *TenantsApi) AddDomain(body *import1.Domain, args ...map[string]interface{}) (*import1.AddDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddDomain(context.Background(), &import3.AddDomainRequest{
		Body: body,
	}, args...)
}

// Add a new domain to the list of allowed domains.
func (api *TenantsServiceApi) AddDomain(ctx context.Context, request *import3.AddDomainRequest, args ...map[string]interface{}) (*import1.AddDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domains"

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
	unmarshalledResp := new(import1.AddDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Create a new tenant.
func (api *TenantsApi) CreateTenant(body *import1.Tenant, args ...map[string]interface{}) (*import1.CreateTenantApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateTenant(context.Background(), &import3.CreateTenantRequest{
		Body: body,
	}, args...)
}

// Create a new tenant.
func (api *TenantsServiceApi) CreateTenant(ctx context.Context, request *import3.CreateTenantRequest, args ...map[string]interface{}) (*import1.CreateTenantApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants"

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
	unmarshalledResp := new(import1.CreateTenantApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Create a new domain allocation for a tenant.
func (api *TenantsApi) CreateTenantDomainAllocation(tenantExtId *string, body *import1.DomainAllocation, args ...map[string]interface{}) (*import1.CreateTenantDomainAllocationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateTenantDomainAllocation(context.Background(), &import3.CreateTenantDomainAllocationRequest{
		TenantExtId: tenantExtId,
		Body:        body,
	}, args...)
}

// Create a new domain allocation for a tenant.
func (api *TenantsServiceApi) CreateTenantDomainAllocation(ctx context.Context, request *import3.CreateTenantDomainAllocationRequest, args ...map[string]interface{}) (*import1.CreateTenantDomainAllocationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
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
	unmarshalledResp := new(import1.CreateTenantDomainAllocationApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete a tenant by its external identifier.
func (api *TenantsApi) DeleteTenantById(extId *string, args ...map[string]interface{}) (*import1.DeleteTenantApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteTenantById(context.Background(), &import3.DeleteTenantByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete a tenant by its external identifier.
func (api *TenantsServiceApi) DeleteTenantById(ctx context.Context, request *import3.DeleteTenantByIdRequest, args ...map[string]interface{}) (*import1.DeleteTenantApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{extId}"

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
	unmarshalledResp := new(import1.DeleteTenantApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete a domain allocation for a tenant by its external identifier.
func (api *TenantsApi) DeleteTenantDomainAllocationById(tenantExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteTenantDomainAllocationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteTenantDomainAllocationById(context.Background(), &import3.DeleteTenantDomainAllocationByIdRequest{
		TenantExtId: tenantExtId,
		ExtId:       extId,
	}, args...)
}

// Delete a domain allocation for a tenant by its external identifier.
func (api *TenantsServiceApi) DeleteTenantDomainAllocationById(ctx context.Context, request *import3.DeleteTenantDomainAllocationByIdRequest, args ...map[string]interface{}) (*import1.DeleteTenantDomainAllocationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId}"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
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
	unmarshalledResp := new(import1.DeleteTenantDomainAllocationApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a specific tenant by its external identifier.
func (api *TenantsApi) GetTenantById(extId *string, args ...map[string]interface{}) (*import1.GetTenantApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTenantById(context.Background(), &import3.GetTenantByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve a specific tenant by its external identifier.
func (api *TenantsServiceApi) GetTenantById(ctx context.Context, request *import3.GetTenantByIdRequest, args ...map[string]interface{}) (*import1.GetTenantApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{extId}"

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
	unmarshalledResp := new(import1.GetTenantApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a specific domain allocation for a tenant by its external identifier.
func (api *TenantsApi) GetTenantDomainAllocationById(tenantExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetTenantDomainAllocationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTenantDomainAllocationById(context.Background(), &import3.GetTenantDomainAllocationByIdRequest{
		TenantExtId: tenantExtId,
		ExtId:       extId,
	}, args...)
}

// Retrieve a specific domain allocation for a tenant by its external identifier.
func (api *TenantsServiceApi) GetTenantDomainAllocationById(ctx context.Context, request *import3.GetTenantDomainAllocationByIdRequest, args ...map[string]interface{}) (*import1.GetTenantDomainAllocationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId}"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetTenantDomainAllocationApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve external network connection settings for a tenant domain allocation.
func (api *TenantsApi) GetTenantExternalNetworkConnection(tenantExtId *string, domainAllocationExtId *string, args ...map[string]interface{}) (*import1.GetTenantExternalNetworkConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTenantExternalNetworkConnection(context.Background(), &import3.GetTenantExternalNetworkConnectionRequest{
		TenantExtId:           tenantExtId,
		DomainAllocationExtId: domainAllocationExtId,
	}, args...)
}

// Retrieve external network connection settings for a tenant domain allocation.
func (api *TenantsServiceApi) GetTenantExternalNetworkConnection(ctx context.Context, request *import3.GetTenantExternalNetworkConnectionRequest, args ...map[string]interface{}) (*import1.GetTenantExternalNetworkConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{domainAllocationExtId}/external-network-connection"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}
	// verify the required parameter 'domainAllocationExtId' is set
	if nil == request.DomainAllocationExtId {
		return nil, client.ReportError("domainAllocationExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"domainAllocationExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainAllocationExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetTenantExternalNetworkConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all allowed domains.
func (api *TenantsApi) ListDomains(args ...map[string]interface{}) (*import1.ListDomainsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDomains(context.Background(), &import3.ListDomainsRequest{}, args...)
}

// List all allowed domains.
func (api *TenantsServiceApi) ListDomains(ctx context.Context, request *import3.ListDomainsRequest, args ...map[string]interface{}) (*import1.ListDomainsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domains"

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
	unmarshalledResp := new(import1.ListDomainsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all domain allocations for a specific tenant.
func (api *TenantsApi) ListTenantDomainAllocations(tenantExtId *string, expand_ *string, args ...map[string]interface{}) (*import1.ListTenantsDomainAllocationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTenantDomainAllocations(context.Background(), &import3.ListTenantDomainAllocationsRequest{
		TenantExtId: tenantExtId,
		Expand_:     expand_,
	}, args...)
}

// List all domain allocations for a specific tenant.
func (api *TenantsServiceApi) ListTenantDomainAllocations(ctx context.Context, request *import3.ListTenantDomainAllocationsRequest, args ...map[string]interface{}) (*import1.ListTenantsDomainAllocationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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
	unmarshalledResp := new(import1.ListTenantsDomainAllocationsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all tenants with pagination support.
func (api *TenantsApi) ListTenants(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.ListTenantsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTenants(context.Background(), &import3.ListTenantsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List all tenants with pagination support.
func (api *TenantsServiceApi) ListTenants(ctx context.Context, request *import3.ListTenantsRequest, args ...map[string]interface{}) (*import1.ListTenantsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants"

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
	unmarshalledResp := new(import1.ListTenantsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update an existing domain allocation for a tenant.
func (api *TenantsApi) UpdateTenantDomainAllocationById(tenantExtId *string, extId *string, body *import1.DomainAllocation, args ...map[string]interface{}) (*import1.UpdateTenantDomainAllocationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateTenantDomainAllocationById(context.Background(), &import3.UpdateTenantDomainAllocationByIdRequest{
		TenantExtId: tenantExtId,
		ExtId:       extId,
		Body:        body,
	}, args...)
}

// Update an existing domain allocation for a tenant.
func (api *TenantsServiceApi) UpdateTenantDomainAllocationById(ctx context.Context, request *import3.UpdateTenantDomainAllocationByIdRequest, args ...map[string]interface{}) (*import1.UpdateTenantDomainAllocationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{extId}"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
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
	unmarshalledResp := new(import1.UpdateTenantDomainAllocationApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update external network connection settings for a tenant domain allocation.
func (api *TenantsApi) UpdateTenantExternalNetworkConnection(tenantExtId *string, domainAllocationExtId *string, body *import1.ExternalNetworkConnConfig, args ...map[string]interface{}) (*import1.UpdateTenantExternalNetworkConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTenantsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateTenantExternalNetworkConnection(context.Background(), &import3.UpdateTenantExternalNetworkConnectionRequest{
		TenantExtId:           tenantExtId,
		DomainAllocationExtId: domainAllocationExtId,
		Body:                  body,
	}, args...)
}

// Update external network connection settings for a tenant domain allocation.
func (api *TenantsServiceApi) UpdateTenantExternalNetworkConnection(ctx context.Context, request *import3.UpdateTenantExternalNetworkConnectionRequest, args ...map[string]interface{}) (*import1.UpdateTenantExternalNetworkConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/tenants/{tenantExtId}/domain-allocations/{domainAllocationExtId}/external-network-connection"

	// verify the required parameter 'tenantExtId' is set
	if nil == request.TenantExtId {
		return nil, client.ReportError("tenantExtId is required and must be specified")
	}
	// verify the required parameter 'domainAllocationExtId' is set
	if nil == request.DomainAllocationExtId {
		return nil, client.ReportError("domainAllocationExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"tenantExtId"+"}", url.PathEscape(client.ParameterToString(*request.TenantExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"domainAllocationExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainAllocationExtId, "")), -1)
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
	unmarshalledResp := new(import1.UpdateTenantExternalNetworkConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
