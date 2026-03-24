package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import22 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/remoteentities"
	"net/http"
	"net/url"
	"strings"
)

type RemoteEntitiesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RemoteEntitiesServiceApi
}

type RemoteEntitiesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRemoteEntitiesApi(apiClient *client.ApiClient) *RemoteEntitiesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RemoteEntitiesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRemoteEntitiesServiceApi(a.ApiClient)

	return a
}

func NewRemoteEntitiesServiceApi(apiClient *client.ApiClient) *RemoteEntitiesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RemoteEntitiesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Information about a subnet from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) GetRemoteSubnetForClusterById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import4.GetRemoteSubnetApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRemoteSubnetForClusterById(context.Background(), &import22.GetRemoteSubnetForClusterByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Information about a subnet from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) GetRemoteSubnetForClusterById(ctx context.Context, request *import22.GetRemoteSubnetForClusterByIdRequest, args ...map[string]interface{}) (*import4.GetRemoteSubnetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-subnets/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.GetRemoteSubnetApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a VPN connection from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) GetRemoteVpnConnectionForClusterById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import4.RemoteVpnConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRemoteVpnConnectionForClusterById(context.Background(), &import22.GetRemoteVpnConnectionForClusterByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get a VPN connection from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) GetRemoteVpnConnectionForClusterById(ctx context.Context, request *import22.GetRemoteVpnConnectionForClusterByIdRequest, args ...map[string]interface{}) (*import4.RemoteVpnConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-vpn-connections/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.RemoteVpnConnectionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a VTEP gateway from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) GetRemoteVtepGatewayForClusterById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import4.RemoteVtepGatewayApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRemoteVtepGatewayForClusterById(context.Background(), &import22.GetRemoteVtepGatewayForClusterByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get a VTEP gateway from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) GetRemoteVtepGatewayForClusterById(ctx context.Context, request *import22.GetRemoteVtepGatewayForClusterByIdRequest, args ...map[string]interface{}) (*import4.RemoteVtepGatewayApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-vtep-gateways/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.RemoteVtepGatewayApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists subnets from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) ListRemoteSubnetsByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListRemoteSubnetsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRemoteSubnetsByClusterId(context.Background(), &import22.ListRemoteSubnetsByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
	}, args...)
}

// Lists subnets from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) ListRemoteSubnetsByClusterId(ctx context.Context, request *import22.ListRemoteSubnetsByClusterIdRequest, args ...map[string]interface{}) (*import4.ListRemoteSubnetsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-subnets"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListRemoteSubnetsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists VPN connections from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) ListRemoteVpnConnectionsByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.RemoteVpnConnectionListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRemoteVpnConnectionsByClusterId(context.Background(), &import22.ListRemoteVpnConnectionsByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
	}, args...)
}

// Lists VPN connections from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) ListRemoteVpnConnectionsByClusterId(ctx context.Context, request *import22.ListRemoteVpnConnectionsByClusterIdRequest, args ...map[string]interface{}) (*import4.RemoteVpnConnectionListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-vpn-connections"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.RemoteVpnConnectionListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists VTEP gateways from the specified Prism Central cluster.
func (api *RemoteEntitiesApi) ListRemoteVtepGatewaysByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.RemoteVtepGatewayListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRemoteEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRemoteVtepGatewaysByClusterId(context.Background(), &import22.ListRemoteVtepGatewaysByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
	}, args...)
}

// Lists VTEP gateways from the specified Prism Central cluster.
func (api *RemoteEntitiesServiceApi) ListRemoteVtepGatewaysByClusterId(ctx context.Context, request *import22.ListRemoteVtepGatewaysByClusterIdRequest, args ...map[string]interface{}) (*import4.RemoteVtepGatewayListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/clusters/{clusterExtId}/remote-vtep-gateways"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import4.RemoteVtepGatewayListApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
