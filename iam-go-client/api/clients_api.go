package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
	import6 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/clients"
	"net/http"
	"net/url"
	"strings"
)

type ClientsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClientsServiceApi
}

type ClientsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClientsApi(apiClient *client.ApiClient) *ClientsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClientsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClientsServiceApi(a.ApiClient)

	return a
}

func NewClientsServiceApi(apiClient *client.ApiClient) *ClientsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClientsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches a registered client information.
func (api *ClientsApi) GetRegisteredClientById(extId *string, args ...map[string]interface{}) (*import1.GetClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRegisteredClientById(context.Background(), &import6.GetRegisteredClientByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a registered client information.
func (api *ClientsServiceApi) GetRegisteredClientById(ctx context.Context, request *import6.GetRegisteredClientByIdRequest, args ...map[string]interface{}) (*import1.GetClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/clients/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
