package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/globalreportsetting"
	"net/http"
	"net/url"
	"strings"
)

type GlobalReportSettingApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *GlobalReportSettingServiceApi
}

type GlobalReportSettingServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewGlobalReportSettingApi(apiClient *client.ApiClient) *GlobalReportSettingApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GlobalReportSettingApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewGlobalReportSettingServiceApi(a.ApiClient)

	return a
}

func NewGlobalReportSettingServiceApi(apiClient *client.ApiClient) *GlobalReportSettingServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GlobalReportSettingServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// This operation retrieves the global report setting belonging to the user.
func (api *GlobalReportSettingApi) GetGlobalReportSetting(userExtId *string, args ...map[string]interface{}) (*import1.GetGlobalReportSettingApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGlobalReportSettingServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetGlobalReportSetting(context.Background(), &import2.GetGlobalReportSettingRequest{
		UserExtId: userExtId,
	}, args...)
}

// This operation retrieves the global report setting belonging to the user.
func (api *GlobalReportSettingServiceApi) GetGlobalReportSetting(ctx context.Context, request *import2.GetGlobalReportSettingRequest, args ...map[string]interface{}) (*import1.GetGlobalReportSettingApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/user/{userExtId}/global-report-setting"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetGlobalReportSettingApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This operation updates the global report setting belonging to the user.
func (api *GlobalReportSettingApi) UpdateGlobalReportSetting(userExtId *string, body *import1.GlobalReportSetting, args ...map[string]interface{}) (*import1.UpdateGlobalReportSettingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGlobalReportSettingServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateGlobalReportSetting(context.Background(), &import2.UpdateGlobalReportSettingRequest{
		UserExtId: userExtId,
		Body:      body,
	}, args...)
}

// This operation updates the global report setting belonging to the user.
func (api *GlobalReportSettingServiceApi) UpdateGlobalReportSetting(ctx context.Context, request *import2.UpdateGlobalReportSettingRequest, args ...map[string]interface{}) (*import1.UpdateGlobalReportSettingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/user/{userExtId}/global-report-setting"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import1.UpdateGlobalReportSettingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
