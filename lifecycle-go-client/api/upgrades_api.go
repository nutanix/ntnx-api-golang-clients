package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import24 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/upgrades"
	"net/http"
	"net/url"
	"strings"
)

type UpgradesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UpgradesServiceApi
}

type UpgradesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUpgradesApi(apiClient *client.ApiClient) *UpgradesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUpgradesServiceApi(a.ApiClient)

	return a
}

func NewUpgradesServiceApi(apiClient *client.ApiClient) *UpgradesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Perform an upgrade operation to update one or more LCM entities to their specified target versions. Before invoking an upgrade, ensure you have: (1) run an inventory (POST /$actions/inventory) to discover entities, (2) computed recommendations (POST /$actions/compute-recommendations) to resolve dependencies and obtain validated entityUpdateSpecs, (3) optionally reviewed notifications (POST /$actions/compute-notifications) to understand the upgrade impact, and (4) run prechecks (POST /$actions/prechecks) to validate cluster readiness. Submit an UpgradeSpec containing the entityUpdateSpecs (entity UUID and target version pairs). For ESX or Hyper-V clusters, provide the managementServer credentials. Use autoHandleFlags to allow LCM to automatically power off pinned VMs during maintenance mode and restore them after the upgrade. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; monitor the task via the Prism tasks API. Use GET /status to check whether an upgrade is already in progress. Only one upgrade can run on a cluster at a time. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. In dark-site mode, ensure the required bundles have been uploaded (POST /bundles) before running the upgrade; images are not downloaded from the portal.
func (api *UpgradesApi) PerformUpgrade(body *import6.UpgradeSpec, xClusterId *string, dryrun_ *bool, args ...map[string]interface{}) (*import7.UpgradeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformUpgrade(context.Background(), &import24.PerformUpgradeRequest{
		Body:       body,
		XClusterId: xClusterId,
		Dryrun_:    dryrun_,
	}, args...)
}

// Perform an upgrade operation to update one or more LCM entities to their specified target versions. Before invoking an upgrade, ensure you have: (1) run an inventory (POST /$actions/inventory) to discover entities, (2) computed recommendations (POST /$actions/compute-recommendations) to resolve dependencies and obtain validated entityUpdateSpecs, (3) optionally reviewed notifications (POST /$actions/compute-notifications) to understand the upgrade impact, and (4) run prechecks (POST /$actions/prechecks) to validate cluster readiness. Submit an UpgradeSpec containing the entityUpdateSpecs (entity UUID and target version pairs). For ESX or Hyper-V clusters, provide the managementServer credentials. Use autoHandleFlags to allow LCM to automatically power off pinned VMs during maintenance mode and restore them after the upgrade. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; monitor the task via the Prism tasks API. Use GET /status to check whether an upgrade is already in progress. Only one upgrade can run on a cluster at a time. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. In dark-site mode, ensure the required bundles have been uploaded (POST /bundles) before running the upgrade; images are not downloaded from the portal.
func (api *UpgradesServiceApi) PerformUpgrade(ctx context.Context, request *import24.PerformUpgradeRequest, args ...map[string]interface{}) (*import7.UpgradeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/upgrade"

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
	unmarshalledResp := new(import7.UpgradeApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
