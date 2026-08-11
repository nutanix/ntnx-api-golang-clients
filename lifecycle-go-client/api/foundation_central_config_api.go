package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import9 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/mgmt"
	import10 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/foundationcentralconfig"
	"net/http"
	"net/url"
	"strings"
)

type FoundationCentralConfigApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *FoundationCentralConfigServiceApi
}

type FoundationCentralConfigServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewFoundationCentralConfigApi(apiClient *client.ApiClient) *FoundationCentralConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &FoundationCentralConfigApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewFoundationCentralConfigServiceApi(a.ApiClient)

	return a
}

func NewFoundationCentralConfigServiceApi(apiClient *client.ApiClient) *FoundationCentralConfigServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &FoundationCentralConfigServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Returns the current configuration settings of the Foundation Central service
func (api *FoundationCentralConfigApi) GetFoundationCentralConfig(args ...map[string]interface{}) (*import9.GetFoundationCentralConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewFoundationCentralConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetFoundationCentralConfig(context.Background(), &import10.GetFoundationCentralConfigRequest{}, args...)
}

// Returns the current configuration settings of the Foundation Central service
func (api *FoundationCentralConfigServiceApi) GetFoundationCentralConfig(ctx context.Context, request *import10.GetFoundationCentralConfigRequest, args ...map[string]interface{}) (*import9.GetFoundationCentralConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/mgmt/foundation-central-config"

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
	unmarshalledResp := new(import9.GetFoundationCentralConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the configuration settings of the Foundation Central service
func (api *FoundationCentralConfigApi) UpdateFoundationCentralConfig(body *import9.FoundationCentralConfig, args ...map[string]interface{}) (*import9.UpdateFoundationCentralConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewFoundationCentralConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateFoundationCentralConfig(context.Background(), &import10.UpdateFoundationCentralConfigRequest{
		Body: body,
	}, args...)
}

// Updates the configuration settings of the Foundation Central service
func (api *FoundationCentralConfigServiceApi) UpdateFoundationCentralConfig(ctx context.Context, request *import10.UpdateFoundationCentralConfigRequest, args ...map[string]interface{}) (*import9.UpdateFoundationCentralConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/mgmt/foundation-central-config"

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
	unmarshalledResp := new(import9.UpdateFoundationCentralConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
