package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	import26 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmprofile"
	"net/http"
	"net/url"
	"strings"
)

type VmProfileApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmProfileServiceApi
}

type VmProfileServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmProfileApi(apiClient *client.ApiClient) *VmProfileApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmProfileApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmProfileServiceApi(a.ApiClient)

	return a
}

func NewVmProfileServiceApi(apiClient *client.ApiClient) *VmProfileServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmProfileServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Deploys a VM from a VM Profile.
func (api *VmProfileApi) DeployVmFromVmProfile(extId *string, body *import21.DeployVmFromVmProfileParams, args ...map[string]interface{}) (*import21.DeployVmFromVmProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeployVmFromVmProfile(context.Background(), &import26.DeployVmFromVmProfileRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Deploys a VM from a VM Profile.
func (api *VmProfileServiceApi) DeployVmFromVmProfile(ctx context.Context, request *import26.DeployVmFromVmProfileRequest, args ...map[string]interface{}) (*import21.DeployVmFromVmProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles/{extId}/$actions/deploy"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.DeployVmFromVmProfileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves configuration details for a VM Profile.
func (api *VmProfileApi) GetVmProfileById(extId *string, args ...map[string]interface{}) (*import21.GetVmProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmProfileById(context.Background(), &import26.GetVmProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves configuration details for a VM Profile.
func (api *VmProfileServiceApi) GetVmProfileById(ctx context.Context, request *import26.GetVmProfileByIdRequest, args ...map[string]interface{}) (*import21.GetVmProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles/{extId}"

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
	unmarshalledResp := new(import21.GetVmProfileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists the VM Profiles defined on the system.
func (api *VmProfileApi) ListVmProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import21.ListVmProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmProfiles(context.Background(), &import26.ListVmProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the VM Profiles defined on the system.
func (api *VmProfileServiceApi) ListVmProfiles(ctx context.Context, request *import26.ListVmProfilesRequest, args ...map[string]interface{}) (*import21.ListVmProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.ListVmProfilesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Shares a VM Profile with a project.
func (api *VmProfileApi) ShareVmProfileWithProject(extId *string, body *import21.ShareVmProfileWithProjectParams, args ...map[string]interface{}) (*import21.ShareVmProfileWithProjectApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShareVmProfileWithProject(context.Background(), &import26.ShareVmProfileWithProjectRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Shares a VM Profile with a project.
func (api *VmProfileServiceApi) ShareVmProfileWithProject(ctx context.Context, request *import26.ShareVmProfileWithProjectRequest, args ...map[string]interface{}) (*import21.ShareVmProfileWithProjectApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles/{extId}/$actions/share"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.ShareVmProfileWithProjectApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Unshares a VM Profile from a project.
func (api *VmProfileApi) UnshareVmProfileFromProject(extId *string, body *import21.UnshareVmProfileFromProjectParams, args ...map[string]interface{}) (*import21.UnshareVmProfileFromProjectApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnshareVmProfileFromProject(context.Background(), &import26.UnshareVmProfileFromProjectRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Unshares a VM Profile from a project.
func (api *VmProfileServiceApi) UnshareVmProfileFromProject(ctx context.Context, request *import26.UnshareVmProfileFromProjectRequest, args ...map[string]interface{}) (*import21.UnshareVmProfileFromProjectApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles/{extId}/$actions/unshare"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.UnshareVmProfileFromProjectApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates configuration details for a VM Profile.
func (api *VmProfileApi) UpdateVmProfileById(extId *string, body *import21.VmProfile, args ...map[string]interface{}) (*import21.UpdateVmProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmProfileServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmProfileById(context.Background(), &import26.UpdateVmProfileByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates configuration details for a VM Profile.
func (api *VmProfileServiceApi) UpdateVmProfileById(ctx context.Context, request *import26.UpdateVmProfileByIdRequest, args ...map[string]interface{}) (*import21.UpdateVmProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-profiles/{extId}"

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
	unmarshalledResp := new(import21.UpdateVmProfileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
