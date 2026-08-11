package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/management"
	import6 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/localdomain"
	"net/http"
	"net/url"
	"strings"
)

type LocalDomainApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *LocalDomainServiceApi
}

type LocalDomainServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewLocalDomainApi(apiClient *client.ApiClient) *LocalDomainApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LocalDomainApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewLocalDomainServiceApi(a.ApiClient)

	return a
}

func NewLocalDomainServiceApi(apiClient *client.ApiClient) *LocalDomainServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LocalDomainServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Returns the domain entity
func (api *LocalDomainApi) GetLocalDomain(args ...map[string]interface{}) (*import5.GetLocalDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLocalDomainServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLocalDomain(context.Background(), &import6.GetLocalDomainRequest{}, args...)
}

// Returns the domain entity
func (api *LocalDomainServiceApi) GetLocalDomain(ctx context.Context, request *import6.GetLocalDomainRequest, args ...map[string]interface{}) (*import5.GetLocalDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/management/local-domain"

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
	unmarshalledResp := new(import5.GetLocalDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Registers the domain with Nutanix Central, establishing a connection between Prism Central and their corresponding Nutanix Central components.
func (api *LocalDomainApi) RegisterLocalDomain(body *import5.LocalDomainRegistrationSpec, args ...map[string]interface{}) (*import5.RegisterLocalDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLocalDomainServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RegisterLocalDomain(context.Background(), &import6.RegisterLocalDomainRequest{
		Body: body,
	}, args...)
}

// Registers the domain with Nutanix Central, establishing a connection between Prism Central and their corresponding Nutanix Central components.
func (api *LocalDomainServiceApi) RegisterLocalDomain(ctx context.Context, request *import6.RegisterLocalDomainRequest, args ...map[string]interface{}) (*import5.RegisterLocalDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/management/local-domain/$actions/register"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import5.RegisterLocalDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Unregister a domain from Nutanix Central.
func (api *LocalDomainApi) UnregisterLocalDomain(args ...map[string]interface{}) (*import5.UnregisterLocalDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLocalDomainServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnregisterLocalDomain(context.Background(), &import6.UnregisterLocalDomainRequest{}, args...)
}

// Unregister a domain from Nutanix Central.
func (api *LocalDomainServiceApi) UnregisterLocalDomain(ctx context.Context, request *import6.UnregisterLocalDomainRequest, args ...map[string]interface{}) (*import5.UnregisterLocalDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/management/local-domain/$actions/unregister"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import5.UnregisterLocalDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
