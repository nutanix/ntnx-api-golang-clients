package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import8 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/inventory"
	"net/http"
	"net/url"
	"strings"
)

type InventoryApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *InventoryServiceApi
}

type InventoryServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewInventoryApi(apiClient *client.ApiClient) *InventoryApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &InventoryApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewInventoryServiceApi(a.ApiClient)

	return a
}

func NewInventoryServiceApi(apiClient *client.ApiClient) *InventoryServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &InventoryServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Perform an LCM inventory operation.
func (api *InventoryApi) PerformInventory(xClusterId *string, body *import5.InventorySpec, dryrun_ *bool, args ...map[string]interface{}) (*import5.InventoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewInventoryServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformInventory(context.Background(), &import8.PerformInventoryRequest{
		XClusterId: xClusterId,
		Body:       body,
		Dryrun_:    dryrun_,
	}, args...)
}

// Perform an LCM inventory operation.
func (api *InventoryServiceApi) PerformInventory(ctx context.Context, request *import8.PerformInventoryRequest, args ...map[string]interface{}) (*import5.InventoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/operations/$actions/inventory"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*request.Dryrun_, ""))
	}
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.InventoryApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
