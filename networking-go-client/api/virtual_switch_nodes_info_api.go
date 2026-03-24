package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import33 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/virtualswitchnodesinfo"
	"net/http"
	"net/url"
	"strings"
)

type VirtualSwitchNodesInfoApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VirtualSwitchNodesInfoServiceApi
}

type VirtualSwitchNodesInfoServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVirtualSwitchNodesInfoApi(apiClient *client.ApiClient) *VirtualSwitchNodesInfoApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VirtualSwitchNodesInfoApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVirtualSwitchNodesInfoServiceApi(a.ApiClient)

	return a
}

func NewVirtualSwitchNodesInfoServiceApi(apiClient *client.ApiClient) *VirtualSwitchNodesInfoServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VirtualSwitchNodesInfoServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Check to see whether a node in a cluster is a storage-only node or not
func (api *VirtualSwitchNodesInfoApi) ListNodeSchedulableStatus(xClusterId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListNodeSchedulableStatusesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVirtualSwitchNodesInfoServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNodeSchedulableStatus(context.Background(), &import33.ListNodeSchedulableStatusRequest{
		XClusterId: xClusterId,
		Page_:      page_,
		Limit_:     limit_,
		Filter_:    filter_,
		Orderby_:   orderby_,
	}, args...)
}

// Check to see whether a node in a cluster is a storage-only node or not
func (api *VirtualSwitchNodesInfoServiceApi) ListNodeSchedulableStatus(ctx context.Context, request *import33.ListNodeSchedulableStatusRequest, args ...map[string]interface{}) (*import4.ListNodeSchedulableStatusesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/node-schedulable-statuses"

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
	if request.XClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
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

	unmarshalledResp := new(import4.ListNodeSchedulableStatusesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
