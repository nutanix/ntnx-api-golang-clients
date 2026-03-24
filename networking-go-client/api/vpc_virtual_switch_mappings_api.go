package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import36 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/vpcvirtualswitchmappings"
	"net/http"
	"net/url"
	"strings"
)

type VpcVirtualSwitchMappingsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VpcVirtualSwitchMappingsServiceApi
}

type VpcVirtualSwitchMappingsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVpcVirtualSwitchMappingsApi(apiClient *client.ApiClient) *VpcVirtualSwitchMappingsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcVirtualSwitchMappingsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVpcVirtualSwitchMappingsServiceApi(a.ApiClient)

	return a
}

func NewVpcVirtualSwitchMappingsServiceApi(apiClient *client.ApiClient) *VpcVirtualSwitchMappingsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcVirtualSwitchMappingsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Set VPC for virtual switch mappings traffic config.
func (api *VpcVirtualSwitchMappingsApi) CreateVpcVirtualSwitchMapping(body *[]import4.VpcVirtualSwitchMapping, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpcVirtualSwitchMappingsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVpcVirtualSwitchMapping(context.Background(), &import36.CreateVpcVirtualSwitchMappingRequest{
		Body: body,
	}, args...)
}

// Set VPC for virtual switch mappings traffic config.
func (api *VpcVirtualSwitchMappingsServiceApi) CreateVpcVirtualSwitchMapping(ctx context.Context, request *import36.CreateVpcVirtualSwitchMappingRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpc-virtual-switch-mappings"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the VPC for virtual switch mappings config.
func (api *VpcVirtualSwitchMappingsApi) ListVpcVirtualSwitchMappings(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListVpcVirtualSwitchMappingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpcVirtualSwitchMappingsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVpcVirtualSwitchMappings(context.Background(), &import36.ListVpcVirtualSwitchMappingsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// Get the VPC for virtual switch mappings config.
func (api *VpcVirtualSwitchMappingsServiceApi) ListVpcVirtualSwitchMappings(ctx context.Context, request *import36.ListVpcVirtualSwitchMappingsRequest, args ...map[string]interface{}) (*import4.ListVpcVirtualSwitchMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpc-virtual-switch-mappings"

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

	unmarshalledResp := new(import4.ListVpcVirtualSwitchMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
