package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import22 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/status"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type StatusApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StatusServiceApi
}

type StatusServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStatusApi(apiClient *client.ApiClient) *StatusApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatusApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStatusServiceApi(a.ApiClient)

	return a
}

func NewStatusServiceApi(apiClient *client.ApiClient) *StatusServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatusServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieve the current operational status of the LCM framework on a cluster. The response includes the installed framework version, whether a framework update is available, any in-progress operation (inventory, prechecks, upgrade, or upload) with its task identifier, the cancel-intent flag, URL accessibility, connectivity type (connected-site or dark-site), and restricted-mode information. Always check this endpoint before starting a new operation (POST /$actions/inventory, POST /$actions/prechecks, or POST /$actions/upgrade) to confirm no conflicting operation is already running. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself.
func (api *StatusApi) GetStatus(xClusterId *string, args ...map[string]interface{}) (*import1.GetStatusApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatusServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStatus(context.Background(), &import22.GetStatusRequest{
		XClusterId: xClusterId,
	}, args...)
}

// Retrieve the current operational status of the LCM framework on a cluster. The response includes the installed framework version, whether a framework update is available, any in-progress operation (inventory, prechecks, upgrade, or upload) with its task identifier, the cancel-intent flag, URL accessibility, connectivity type (connected-site or dark-site), and restricted-mode information. Always check this endpoint before starting a new operation (POST /$actions/inventory, POST /$actions/prechecks, or POST /$actions/upgrade) to confirm no conflicting operation is already running. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself.
func (api *StatusServiceApi) GetStatus(ctx context.Context, request *import22.GetStatusRequest, args ...map[string]interface{}) (*import1.GetStatusApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/status"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.GetStatusApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
