package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import15 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/ahv/config"
	import16 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/gpuprofiles"
	"net/http"
	"net/url"
	"strings"
)

type GpuProfilesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *GpuProfilesServiceApi
}

type GpuProfilesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewGpuProfilesApi(apiClient *client.ApiClient) *GpuProfilesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GpuProfilesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewGpuProfilesServiceApi(a.ApiClient)

	return a
}

func NewGpuProfilesServiceApi(apiClient *client.ApiClient) *GpuProfilesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GpuProfilesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get a list of all physical GPU profiles available across the managed infrastructure. Physical GPU profiles represent GPU hardware that can be attached to virtual machines in passthrough mode.
func (api *GpuProfilesApi) ListAhvPhysicalGpuProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import15.ListAhvPhysicalGpuProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGpuProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAhvPhysicalGpuProfiles(context.Background(), &import16.ListAhvPhysicalGpuProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Get a list of all physical GPU profiles available across the managed infrastructure. Physical GPU profiles represent GPU hardware that can be attached to virtual machines in passthrough mode.
func (api *GpuProfilesServiceApi) ListAhvPhysicalGpuProfiles(ctx context.Context, request *import16.ListAhvPhysicalGpuProfilesRequest, args ...map[string]interface{}) (*import15.ListAhvPhysicalGpuProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/ahv/config/physical-gpu-profiles"

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
	unmarshalledResp := new(import15.ListAhvPhysicalGpuProfilesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get a list of all virtual GPU profiles available across the managed infrastructure. Virtual GPU profiles define resource allocation templates for virtual GPUs that can be assigned to virtual machines.
func (api *GpuProfilesApi) ListAhvVirtualGpuProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import15.ListAhvVirtualGpuProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGpuProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAhvVirtualGpuProfiles(context.Background(), &import16.ListAhvVirtualGpuProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Get a list of all virtual GPU profiles available across the managed infrastructure. Virtual GPU profiles define resource allocation templates for virtual GPUs that can be assigned to virtual machines.
func (api *GpuProfilesServiceApi) ListAhvVirtualGpuProfiles(ctx context.Context, request *import16.ListAhvVirtualGpuProfilesRequest, args ...map[string]interface{}) (*import15.ListAhvVirtualGpuProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/ahv/config/virtual-gpu-profiles"

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
	unmarshalledResp := new(import15.ListAhvVirtualGpuProfilesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
