package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/dashboardsettings"
	"net/http"
	"net/url"
	"strings"
)

type DashboardSettingsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DashboardSettingsServiceApi
}

type DashboardSettingsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDashboardSettingsApi(apiClient *client.ApiClient) *DashboardSettingsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardSettingsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDashboardSettingsServiceApi(a.ApiClient)

	return a
}

func NewDashboardSettingsServiceApi(apiClient *client.ApiClient) *DashboardSettingsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardSettingsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieve the global dashboard settings.
func (api *DashboardSettingsApi) GetDashboardSettings(args ...map[string]interface{}) (*import1.GetDashboardSettingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardSettingsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDashboardSettings(context.Background(), &import4.GetDashboardSettingsRequest{}, args...)
}

// Retrieve the global dashboard settings.
func (api *DashboardSettingsServiceApi) GetDashboardSettings(ctx context.Context, request *import4.GetDashboardSettingsRequest, args ...map[string]interface{}) (*import1.GetDashboardSettingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboard-settings"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.GetDashboardSettingsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
