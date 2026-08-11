package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import14 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/inventory"
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

// Trigger an inventory operation to discover all upgradable software and firmware entities on a cluster. Inventory scans cluster nodes, identifies installed components and their current versions, and checks the configured LCM repository for available updates. The operation is asynchronous and returns a TaskReference; poll the task to track progress. Once completed, discovered entities are available via GET /entities and the cluster summary is updated in GET /lcm-summaries. An optional InventorySpec body allows specifying the inventory type and vendor management credentials. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. Only one inventory operation can run on a cluster at a time; check GET /status before invoking. In dark-site mode, inventory discovers available updates only from bundles that have been uploaded to the cluster (POST /bundles); it does not contact the Nutanix portal.
func (api *InventoryApi) PerformInventory(xClusterId *string, body *import7.InventorySpec, dryrun_ *bool, args ...map[string]interface{}) (*import7.InventoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewInventoryServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformInventory(context.Background(), &import14.PerformInventoryRequest{
		XClusterId: xClusterId,
		Body:       body,
		Dryrun_:    dryrun_,
	}, args...)
}

// Trigger an inventory operation to discover all upgradable software and firmware entities on a cluster. Inventory scans cluster nodes, identifies installed components and their current versions, and checks the configured LCM repository for available updates. The operation is asynchronous and returns a TaskReference; poll the task to track progress. Once completed, discovered entities are available via GET /entities and the cluster summary is updated in GET /lcm-summaries. An optional InventorySpec body allows specifying the inventory type and vendor management credentials. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. Only one inventory operation can run on a cluster at a time; check GET /status before invoking. In dark-site mode, inventory discovers available updates only from bundles that have been uploaded to the cluster (POST /bundles); it does not contact the Nutanix portal.
func (api *InventoryServiceApi) PerformInventory(ctx context.Context, request *import14.PerformInventoryRequest, args ...map[string]interface{}) (*import7.InventoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/inventory"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import7.InventoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
