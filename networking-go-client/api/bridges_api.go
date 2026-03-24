package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/bridges"
	"net/http"
	"net/url"
	"strings"
)

type BridgesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *BridgesServiceApi
}

type BridgesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewBridgesApi(apiClient *client.ApiClient) *BridgesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BridgesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewBridgesServiceApi(a.ApiClient)

	return a
}

func NewBridgesServiceApi(apiClient *client.ApiClient) *BridgesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BridgesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a Virtual Switch from an existing bridge.
func (api *BridgesApi) MigrateBridge(body *import4.Bridge, xClusterId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBridgesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MigrateBridge(context.Background(), &import7.MigrateBridgeRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Create a Virtual Switch from an existing bridge.
func (api *BridgesServiceApi) MigrateBridge(ctx context.Context, request *import7.MigrateBridgeRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/virtual-switches/$actions/migrate"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
