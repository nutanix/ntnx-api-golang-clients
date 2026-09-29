package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
	import9 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/securityprofiles"
	"net/http"
	"net/url"
	"strings"
)

type SecurityProfilesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SecurityProfilesServiceApi
}

type SecurityProfilesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSecurityProfilesApi(apiClient *client.ApiClient) *SecurityProfilesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecurityProfilesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSecurityProfilesServiceApi(a.ApiClient)

	return a
}

func NewSecurityProfilesServiceApi(apiClient *client.ApiClient) *SecurityProfilesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecurityProfilesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetch Advanced Configs for cluster like coreDump, consentBanner etc.
func (api *SecurityProfilesApi) GetAdvancedConfig(securityProfileExtId *string, args ...map[string]interface{}) (*import3.GetAdvancedConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetAdvancedConfig(context.Background(), &import9.GetAdvancedConfigRequest{
		SecurityProfileExtId: securityProfileExtId,
	}, args...)
}

// Fetch Advanced Configs for cluster like coreDump, consentBanner etc.
func (api *SecurityProfilesServiceApi) GetAdvancedConfig(ctx context.Context, request *import9.GetAdvancedConfigRequest, args ...map[string]interface{}) (*import3.GetAdvancedConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-profiles/{securityProfileExtId}/advanced-config"

	// verify the required parameter 'securityProfileExtId' is set
	if nil == request.SecurityProfileExtId {
		return nil, client.ReportError("securityProfileExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"securityProfileExtId"+"}", url.PathEscape(client.ParameterToString(*request.SecurityProfileExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetAdvancedConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetch the active security profile on AOS & AHV, profile reflects the current security config posture.
func (api *SecurityProfilesApi) GetSecurityProfileById(extId *string, args ...map[string]interface{}) (*import3.GetSecurityProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSecurityProfileById(context.Background(), &import9.GetSecurityProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetch the active security profile on AOS & AHV, profile reflects the current security config posture.
func (api *SecurityProfilesServiceApi) GetSecurityProfileById(ctx context.Context, request *import9.GetSecurityProfileByIdRequest, args ...map[string]interface{}) (*import3.GetSecurityProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-profiles/{extId}"

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
	unmarshalledResp := new(import3.GetSecurityProfileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetch the list of active security profiles on all registered clusters. Use profile external identifier to fetch/update the security profile on individual cluster.
func (api *SecurityProfilesApi) ListSecurityProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListSecurityProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSecurityProfiles(context.Background(), &import9.ListSecurityProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Fetch the list of active security profiles on all registered clusters. Use profile external identifier to fetch/update the security profile on individual cluster.
func (api *SecurityProfilesServiceApi) ListSecurityProfiles(ctx context.Context, request *import9.ListSecurityProfilesRequest, args ...map[string]interface{}) (*import3.ListSecurityProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-profiles"

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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.ListSecurityProfilesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update Core Dump Configs for userCore, kernelCore and Customise Consent Banner.
func (api *SecurityProfilesApi) UpdateAdvancedConfig(securityProfileExtId *string, body *import3.AdvancedConfig, args ...map[string]interface{}) (*import3.UpdateAdvancedConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateAdvancedConfig(context.Background(), &import9.UpdateAdvancedConfigRequest{
		SecurityProfileExtId: securityProfileExtId,
		Body:                 body,
	}, args...)
}

// Update Core Dump Configs for userCore, kernelCore and Customise Consent Banner.
func (api *SecurityProfilesServiceApi) UpdateAdvancedConfig(ctx context.Context, request *import9.UpdateAdvancedConfigRequest, args ...map[string]interface{}) (*import3.UpdateAdvancedConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-profiles/{securityProfileExtId}/advanced-config"

	// verify the required parameter 'securityProfileExtId' is set
	if nil == request.SecurityProfileExtId {
		return nil, client.ReportError("securityProfileExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"securityProfileExtId"+"}", url.PathEscape(client.ParameterToString(*request.SecurityProfileExtId, "")), -1)
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
	unmarshalledResp := new(import3.UpdateAdvancedConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update the security profile for AOS & AHV, changing profile will trigger a rolling reboot on the cluster.
func (api *SecurityProfilesApi) UpdateSecurityProfileById(extId *string, body *import3.SecurityProfile, args ...map[string]interface{}) (*import3.UpdateSecurityProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSecurityProfileById(context.Background(), &import9.UpdateSecurityProfileByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the security profile for AOS & AHV, changing profile will trigger a rolling reboot on the cluster.
func (api *SecurityProfilesServiceApi) UpdateSecurityProfileById(ctx context.Context, request *import9.UpdateSecurityProfileByIdRequest, args ...map[string]interface{}) (*import3.UpdateSecurityProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-profiles/{extId}"

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
	unmarshalledResp := new(import3.UpdateSecurityProfileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
