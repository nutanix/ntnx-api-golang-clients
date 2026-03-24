package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type ConfigApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ConfigServiceApi
}

type ConfigServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewConfigApi(apiClient *client.ApiClient) *ConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ConfigApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewConfigServiceApi(a.ApiClient)

	return a
}

func NewConfigServiceApi(apiClient *client.ApiClient) *ConfigServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ConfigServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get LCM configuration.
func (api *ConfigApi) GetConfig(xClusterId *string, args ...map[string]interface{}) (*import1.GetConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetConfig(context.Background(), &import3.GetConfigRequest{
		XClusterId: xClusterId,
	}, args...)
}

// Get LCM configuration.
func (api *ConfigServiceApi) GetConfig(ctx context.Context, request *import3.GetConfigRequest, args ...map[string]interface{}) (*import1.GetConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/config"

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

	unmarshalledResp := new(import1.GetConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update LCM configuration.
func (api *ConfigApi) UpdateConfig(body *import1.Config, xClusterId *string, args ...map[string]interface{}) (*import1.UpdateConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateConfig(context.Background(), &import3.UpdateConfigRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Update LCM configuration.
func (api *ConfigServiceApi) UpdateConfig(ctx context.Context, request *import3.UpdateConfigRequest, args ...map[string]interface{}) (*import1.UpdateConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/config"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
