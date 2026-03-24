package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import16 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/welcomebanner"
	"net/http"
	"net/url"
	"strings"
)

type WelcomeBannerApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *WelcomeBannerServiceApi
}

type WelcomeBannerServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewWelcomeBannerApi(apiClient *client.ApiClient) *WelcomeBannerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &WelcomeBannerApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewWelcomeBannerServiceApi(a.ApiClient)

	return a
}

func NewWelcomeBannerServiceApi(apiClient *client.ApiClient) *WelcomeBannerServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &WelcomeBannerServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches the configured welcome banner.
func (api *WelcomeBannerApi) GetWelcomeBanner(args ...map[string]interface{}) (*import4.GetWelcomeBannerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewWelcomeBannerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetWelcomeBanner(context.Background(), &import16.GetWelcomeBannerRequest{}, args...)
}

// Fetches the configured welcome banner.
func (api *WelcomeBannerServiceApi) GetWelcomeBanner(ctx context.Context, request *import16.GetWelcomeBannerRequest, args ...map[string]interface{}) (*import4.GetWelcomeBannerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/welcome-banner"

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

	unmarshalledResp := new(import4.GetWelcomeBannerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the welcome banner.
func (api *WelcomeBannerApi) UpdateWelcomeBanner(body *import4.WelcomeBanner, args ...map[string]interface{}) (*import4.UpdateWelcomeBannerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewWelcomeBannerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateWelcomeBanner(context.Background(), &import16.UpdateWelcomeBannerRequest{
		Body: body,
	}, args...)
}

// Updates the welcome banner.
func (api *WelcomeBannerServiceApi) UpdateWelcomeBanner(ctx context.Context, request *import16.UpdateWelcomeBannerRequest, args ...map[string]interface{}) (*import4.UpdateWelcomeBannerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/welcome-banner"

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

	unmarshalledResp := new(import4.UpdateWelcomeBannerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
