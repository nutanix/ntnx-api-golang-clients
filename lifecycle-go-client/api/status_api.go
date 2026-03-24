package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import14 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/status"
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

// Get the LCM framework status.
func (api *StatusApi) GetStatus(xClusterId *string, args ...map[string]interface{}) (*import1.GetStatusApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStatusServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStatus(context.Background(), &import14.GetStatusRequest{
		XClusterId: xClusterId,
	}, args...)
}

// Get the LCM framework status.
func (api *StatusServiceApi) GetStatus(ctx context.Context, request *import14.GetStatusRequest, args ...map[string]interface{}) (*import1.GetStatusApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/status"

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

	unmarshalledResp := new(import1.GetStatusApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
