package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/clustercapabilities"
	"net/http"
	"net/url"
	"strings"
)

type ClusterCapabilitiesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClusterCapabilitiesServiceApi
}

type ClusterCapabilitiesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClusterCapabilitiesApi(apiClient *client.ApiClient) *ClusterCapabilitiesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterCapabilitiesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClusterCapabilitiesServiceApi(a.ApiClient)

	return a
}

func NewClusterCapabilitiesServiceApi(apiClient *client.ApiClient) *ClusterCapabilitiesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterCapabilitiesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// List the capabilities for one or more cluster UUIDs.
func (api *ClusterCapabilitiesApi) ListClusterCapabilities(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListClusterCapabilitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterCapabilitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListClusterCapabilities(context.Background(), &import8.ListClusterCapabilitiesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List the capabilities for one or more cluster UUIDs.
func (api *ClusterCapabilitiesServiceApi) ListClusterCapabilities(ctx context.Context, request *import8.ListClusterCapabilitiesRequest, args ...map[string]interface{}) (*import4.ListClusterCapabilitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/capabilities"

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

	unmarshalledResp := new(import4.ListClusterCapabilitiesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
