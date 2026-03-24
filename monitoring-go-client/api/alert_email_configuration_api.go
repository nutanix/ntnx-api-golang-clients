package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/alertemailconfiguration"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type AlertEmailConfigurationApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *AlertEmailConfigurationServiceApi
}

type AlertEmailConfigurationServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAlertEmailConfigurationApi(apiClient *client.ApiClient) *AlertEmailConfigurationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AlertEmailConfigurationApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewAlertEmailConfigurationServiceApi(a.ApiClient)

	return a
}

func NewAlertEmailConfigurationServiceApi(apiClient *client.ApiClient) *AlertEmailConfigurationServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AlertEmailConfigurationServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches the configuration that is used to send alert emails.
func (api *AlertEmailConfigurationApi) GetAlertEmailConfiguration(args ...map[string]interface{}) (*import1.GetAlertEmailConfigurationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAlertEmailConfigurationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetAlertEmailConfiguration(context.Background(), &import2.GetAlertEmailConfigurationRequest{}, args...)
}

// Fetches the configuration that is used to send alert emails.
func (api *AlertEmailConfigurationServiceApi) GetAlertEmailConfiguration(ctx context.Context, request *import2.GetAlertEmailConfigurationRequest, args ...map[string]interface{}) (*import1.GetAlertEmailConfigurationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/email-config"

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

	unmarshalledResp := new(import1.GetAlertEmailConfigurationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration that is used to send alert emails.
func (api *AlertEmailConfigurationApi) UpdateAlertEmailConfiguration(body *import1.AlertEmailConfiguration, args ...map[string]interface{}) (*import1.UpdateAlertEmailConfigurationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAlertEmailConfigurationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateAlertEmailConfiguration(context.Background(), &import2.UpdateAlertEmailConfigurationRequest{
		Body: body,
	}, args...)
}

// Updates the configuration that is used to send alert emails.
func (api *AlertEmailConfigurationServiceApi) UpdateAlertEmailConfiguration(ctx context.Context, request *import2.UpdateAlertEmailConfigurationRequest, args ...map[string]interface{}) (*import1.UpdateAlertEmailConfigurationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/email-config"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateAlertEmailConfigurationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
