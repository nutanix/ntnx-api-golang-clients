package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	import23 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmguestcustomizationprofiles"
	"net/http"
	"net/url"
	"strings"
)

type VmGuestCustomizationProfilesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmGuestCustomizationProfilesServiceApi
}

type VmGuestCustomizationProfilesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmGuestCustomizationProfilesApi(apiClient *client.ApiClient) *VmGuestCustomizationProfilesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmGuestCustomizationProfilesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(a.ApiClient)

	return a
}

func NewVmGuestCustomizationProfilesServiceApi(apiClient *client.ApiClient) *VmGuestCustomizationProfilesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmGuestCustomizationProfilesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new VM Guest Customization profile with the provided configuration.
func (api *VmGuestCustomizationProfilesApi) CreateVmGuestCustomizationProfile(body *import19.VmGuestCustomizationProfile, args ...map[string]interface{}) (*import19.CreateVmGuestCustomizationProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVmGuestCustomizationProfile(context.Background(), &import23.CreateVmGuestCustomizationProfileRequest{
		Body: body,
	}, args...)
}

// Creates a new VM Guest Customization profile with the provided configuration.
func (api *VmGuestCustomizationProfilesServiceApi) CreateVmGuestCustomizationProfile(ctx context.Context, request *import23.CreateVmGuestCustomizationProfileRequest, args ...map[string]interface{}) (*import19.CreateVmGuestCustomizationProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vm-guest-customization-profiles"

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

	unmarshalledResp := new(import19.CreateVmGuestCustomizationProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the specified VM Guest Customization Profile.
func (api *VmGuestCustomizationProfilesApi) DeleteVmGuestCustomizationProfileById(extId *string, args ...map[string]interface{}) (*import19.DeleteVmGuestCustomizationProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmGuestCustomizationProfileById(context.Background(), &import23.DeleteVmGuestCustomizationProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the specified VM Guest Customization Profile.
func (api *VmGuestCustomizationProfilesServiceApi) DeleteVmGuestCustomizationProfileById(ctx context.Context, request *import23.DeleteVmGuestCustomizationProfileByIdRequest, args ...map[string]interface{}) (*import19.DeleteVmGuestCustomizationProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vm-guest-customization-profiles/{extId}"

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

	unmarshalledResp := new(import19.DeleteVmGuestCustomizationProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the VM Guest Customization Profile configuration of the provided VM Guest Customization Profile external identifier.
func (api *VmGuestCustomizationProfilesApi) GetVmGuestCustomizationProfileById(extId *string, args ...map[string]interface{}) (*import19.GetVmGuestCustomizationProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmGuestCustomizationProfileById(context.Background(), &import23.GetVmGuestCustomizationProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the VM Guest Customization Profile configuration of the provided VM Guest Customization Profile external identifier.
func (api *VmGuestCustomizationProfilesServiceApi) GetVmGuestCustomizationProfileById(ctx context.Context, request *import23.GetVmGuestCustomizationProfileByIdRequest, args ...map[string]interface{}) (*import19.GetVmGuestCustomizationProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vm-guest-customization-profiles/{extId}"

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

	unmarshalledResp := new(import19.GetVmGuestCustomizationProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists VM Guest Customization Profiles.
func (api *VmGuestCustomizationProfilesApi) ListVmGuestCustomizationProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import19.ListVmGuestCustomizationProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmGuestCustomizationProfiles(context.Background(), &import23.ListVmGuestCustomizationProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists VM Guest Customization Profiles.
func (api *VmGuestCustomizationProfilesServiceApi) ListVmGuestCustomizationProfiles(ctx context.Context, request *import23.ListVmGuestCustomizationProfilesRequest, args ...map[string]interface{}) (*import19.ListVmGuestCustomizationProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vm-guest-customization-profiles"

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

	unmarshalledResp := new(import19.ListVmGuestCustomizationProfilesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the specified VM Guest Customization Profile.
func (api *VmGuestCustomizationProfilesApi) UpdateVmGuestCustomizationProfileById(extId *string, body *import19.VmGuestCustomizationProfile, args ...map[string]interface{}) (*import19.UpdateVmGuestCustomizationProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmGuestCustomizationProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmGuestCustomizationProfileById(context.Background(), &import23.UpdateVmGuestCustomizationProfileByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the specified VM Guest Customization Profile.
func (api *VmGuestCustomizationProfilesServiceApi) UpdateVmGuestCustomizationProfileById(ctx context.Context, request *import23.UpdateVmGuestCustomizationProfileByIdRequest, args ...map[string]interface{}) (*import19.UpdateVmGuestCustomizationProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vm-guest-customization-profiles/{extId}"

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

	unmarshalledResp := new(import19.UpdateVmGuestCustomizationProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
