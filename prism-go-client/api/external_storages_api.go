package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/externalstorages"
	"net/http"
	"net/url"
	"strings"
)

type ExternalStoragesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ExternalStoragesServiceApi
}

type ExternalStoragesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewExternalStoragesApi(apiClient *client.ApiClient) *ExternalStoragesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalStoragesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewExternalStoragesServiceApi(a.ApiClient)

	return a
}

func NewExternalStoragesServiceApi(apiClient *client.ApiClient) *ExternalStoragesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalStoragesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieves details of a specific external storage resource by its unique identifier.
func (api *ExternalStoragesApi) GetExternalStorageById(extId *string, args ...map[string]interface{}) (*import3.GetExternalStorageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalStoragesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetExternalStorageById(context.Background(), &import8.GetExternalStorageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves details of a specific external storage resource by its unique identifier.
func (api *ExternalStoragesServiceApi) GetExternalStorageById(ctx context.Context, request *import8.GetExternalStorageByIdRequest, args ...map[string]interface{}) (*import3.GetExternalStorageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/external-storages/{extId}"

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

	unmarshalledResp := new(import3.GetExternalStorageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
