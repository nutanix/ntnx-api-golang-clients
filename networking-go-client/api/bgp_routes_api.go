package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/bgproutes"
	"net/http"
	"net/url"
	"strings"
)

type BgpRoutesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *BgpRoutesServiceApi
}

type BgpRoutesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewBgpRoutesApi(apiClient *client.ApiClient) *BgpRoutesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BgpRoutesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewBgpRoutesServiceApi(a.ApiClient)

	return a
}

func NewBgpRoutesServiceApi(apiClient *client.ApiClient) *BgpRoutesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BgpRoutesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the specified read-only route of the specified BGP session.
func (api *BgpRoutesApi) GetRouteForBgpSessionById(extId *string, bgpSessionExtId *string, args ...map[string]interface{}) (*import4.GetBgpRouteApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBgpRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRouteForBgpSessionById(context.Background(), &import5.GetRouteForBgpSessionByIdRequest{
		ExtId:           extId,
		BgpSessionExtId: bgpSessionExtId,
	}, args...)
}

// Get the specified read-only route of the specified BGP session.
func (api *BgpRoutesServiceApi) GetRouteForBgpSessionById(ctx context.Context, request *import5.GetRouteForBgpSessionByIdRequest, args ...map[string]interface{}) (*import4.GetBgpRouteApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/bgp-sessions/{bgpSessionExtId}/bgp-routes/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'bgpSessionExtId' is set
	if nil == request.BgpSessionExtId {
		return nil, client.ReportError("bgpSessionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"bgpSessionExtId"+"}", url.PathEscape(client.ParameterToString(*request.BgpSessionExtId, "")), -1)
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

	unmarshalledResp := new(import4.GetBgpRouteApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists read-only routes of the specified BGP session.
func (api *BgpRoutesApi) ListRoutesByBgpSessionId(bgpSessionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListBgpRoutesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBgpRoutesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRoutesByBgpSessionId(context.Background(), &import5.ListRoutesByBgpSessionIdRequest{
		BgpSessionExtId: bgpSessionExtId,
		Page_:           page_,
		Limit_:          limit_,
		Filter_:         filter_,
		Orderby_:        orderby_,
	}, args...)
}

// Lists read-only routes of the specified BGP session.
func (api *BgpRoutesServiceApi) ListRoutesByBgpSessionId(ctx context.Context, request *import5.ListRoutesByBgpSessionIdRequest, args ...map[string]interface{}) (*import4.ListBgpRoutesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/bgp-sessions/{bgpSessionExtId}/bgp-routes"

	// verify the required parameter 'bgpSessionExtId' is set
	if nil == request.BgpSessionExtId {
		return nil, client.ReportError("bgpSessionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"bgpSessionExtId"+"}", url.PathEscape(client.ParameterToString(*request.BgpSessionExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListBgpRoutesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
