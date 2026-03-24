package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import39 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/vpnconnections"
	"net/http"
	"net/url"
	"strings"
)

type VpnConnectionsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VpnConnectionsServiceApi
}

type VpnConnectionsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVpnConnectionsApi(apiClient *client.ApiClient) *VpnConnectionsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpnConnectionsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVpnConnectionsServiceApi(a.ApiClient)

	return a
}

func NewVpnConnectionsServiceApi(apiClient *client.ApiClient) *VpnConnectionsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpnConnectionsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create VPN connection.
func (api *VpnConnectionsApi) CreateVpnConnection(body *import4.VpnConnection, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVpnConnection(context.Background(), &import39.CreateVpnConnectionRequest{
		Body: body,
	}, args...)
}

// Create VPN connection.
func (api *VpnConnectionsServiceApi) CreateVpnConnection(ctx context.Context, request *import39.CreateVpnConnectionRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections"

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

// Delete VPN connection request body
func (api *VpnConnectionsApi) DeleteVpnConnectionById(extId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVpnConnectionById(context.Background(), &import39.DeleteVpnConnectionByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete VPN connection request body
func (api *VpnConnectionsServiceApi) DeleteVpnConnectionById(ctx context.Context, request *import39.DeleteVpnConnectionByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get third-party VPN appliance configuration.
func (api *VpnConnectionsApi) GetVpnApplianceForVpnConnectionById(vpnConnectionExtId *string, extId *string, args ...map[string]interface{}) (*string, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVpnApplianceForVpnConnectionById(context.Background(), &import39.GetVpnApplianceForVpnConnectionByIdRequest{
		VpnConnectionExtId: vpnConnectionExtId,
		ExtId:              extId,
	}, args...)
}

// Get third-party VPN appliance configuration.
func (api *VpnConnectionsServiceApi) GetVpnApplianceForVpnConnectionById(ctx context.Context, request *import39.GetVpnApplianceForVpnConnectionByIdRequest, args ...map[string]interface{}) (*string, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections/{vpnConnectionExtId}/vpn-vendor-configs/{extId}"

	// verify the required parameter 'vpnConnectionExtId' is set
	if nil == request.VpnConnectionExtId {
		return nil, client.ReportError("vpnConnectionExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpnConnectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.VpnConnectionExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"text/plain", "application/json"}

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

	unmarshalledResp := new(string)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get VPN connection for a specified ExtId.
func (api *VpnConnectionsApi) GetVpnConnectionById(extId *string, args ...map[string]interface{}) (*import4.GetVpnConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVpnConnectionById(context.Background(), &import39.GetVpnConnectionByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get VPN connection for a specified ExtId.
func (api *VpnConnectionsServiceApi) GetVpnConnectionById(ctx context.Context, request *import39.GetVpnConnectionByIdRequest, args ...map[string]interface{}) (*import4.GetVpnConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections/{extId}"

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

	unmarshalledResp := new(import4.GetVpnConnectionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List third-party VPN appliances for which configurations are available to download.
func (api *VpnConnectionsApi) ListVpnAppliancesByVpnConnectionId(vpnConnectionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListVpnVendorConfigsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVpnAppliancesByVpnConnectionId(context.Background(), &import39.ListVpnAppliancesByVpnConnectionIdRequest{
		VpnConnectionExtId: vpnConnectionExtId,
		Page_:              page_,
		Limit_:             limit_,
		Filter_:            filter_,
		Orderby_:           orderby_,
	}, args...)
}

// List third-party VPN appliances for which configurations are available to download.
func (api *VpnConnectionsServiceApi) ListVpnAppliancesByVpnConnectionId(ctx context.Context, request *import39.ListVpnAppliancesByVpnConnectionIdRequest, args ...map[string]interface{}) (*import4.ListVpnVendorConfigsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections/{vpnConnectionExtId}/vpn-vendor-configs"

	// verify the required parameter 'vpnConnectionExtId' is set
	if nil == request.VpnConnectionExtId {
		return nil, client.ReportError("vpnConnectionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpnConnectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.VpnConnectionExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListVpnVendorConfigsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List the VPN connections.
func (api *VpnConnectionsApi) ListVpnConnections(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListVpnConnectionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVpnConnections(context.Background(), &import39.ListVpnConnectionsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List the VPN connections.
func (api *VpnConnectionsServiceApi) ListVpnConnections(ctx context.Context, request *import39.ListVpnConnectionsRequest, args ...map[string]interface{}) (*import4.ListVpnConnectionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections"

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

	unmarshalledResp := new(import4.ListVpnConnectionsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update VPN connection.
func (api *VpnConnectionsApi) UpdateVpnConnectionById(extId *string, body *import4.VpnConnection, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVpnConnectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVpnConnectionById(context.Background(), &import39.UpdateVpnConnectionByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update VPN connection.
func (api *VpnConnectionsServiceApi) UpdateVpnConnectionById(ctx context.Context, request *import39.UpdateVpnConnectionByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/vpn-connections/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
