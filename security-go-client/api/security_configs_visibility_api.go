package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/securityconfigsvisibility"
	"net/http"
	"net/url"
	"strings"
)

type SecurityConfigsVisibilityApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SecurityConfigsVisibilityServiceApi
}

type SecurityConfigsVisibilityServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSecurityConfigsVisibilityApi(apiClient *client.ApiClient) *SecurityConfigsVisibilityApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecurityConfigsVisibilityApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSecurityConfigsVisibilityServiceApi(a.ApiClient)

	return a
}

func NewSecurityConfigsVisibilityServiceApi(apiClient *client.ApiClient) *SecurityConfigsVisibilityServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecurityConfigsVisibilityServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetch the list security configurations settings being displayed on the PC dashboard.
func (api *SecurityConfigsVisibilityApi) GetSecurityConfigsVisibilitySetting(select_ *string, args ...map[string]interface{}) (*import3.GetSecurityConfigsVisibilitySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityConfigsVisibilityServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSecurityConfigsVisibilitySetting(context.Background(), &import8.GetSecurityConfigsVisibilitySettingRequest{
		Select_: select_,
	}, args...)
}

// Fetch the list security configurations settings being displayed on the PC dashboard.
func (api *SecurityConfigsVisibilityServiceApi) GetSecurityConfigsVisibilitySetting(ctx context.Context, request *import8.GetSecurityConfigsVisibilitySettingRequest, args ...map[string]interface{}) (*import3.GetSecurityConfigsVisibilitySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-configs-visibility-setting"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*request.Select_, ""))
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
	unmarshalledResp := new(import3.GetSecurityConfigsVisibilitySettingApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update the list security configurations settings being displayed on the PC dashboard.
func (api *SecurityConfigsVisibilityApi) UpdateSecurityConfigsVisibilitySetting(body *import3.SecurityConfigVisibilitySetting, args ...map[string]interface{}) (*import3.UpdateSecurityConfigsVisibilitySettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecurityConfigsVisibilityServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSecurityConfigsVisibilitySetting(context.Background(), &import8.UpdateSecurityConfigsVisibilitySettingRequest{
		Body: body,
	}, args...)
}

// Update the list security configurations settings being displayed on the PC dashboard.
func (api *SecurityConfigsVisibilityServiceApi) UpdateSecurityConfigsVisibilitySetting(ctx context.Context, request *import8.UpdateSecurityConfigsVisibilitySettingRequest, args ...map[string]interface{}) (*import3.UpdateSecurityConfigsVisibilitySettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.2/config/security-configs-visibility-setting"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.UpdateSecurityConfigsVisibilitySettingApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
