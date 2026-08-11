package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import20 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/prechecks"
	"net/http"
	"net/url"
	"strings"
)

type PrechecksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *PrechecksServiceApi
}

type PrechecksServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewPrechecksApi(apiClient *client.ApiClient) *PrechecksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PrechecksApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewPrechecksServiceApi(a.ApiClient)

	return a
}

func NewPrechecksServiceApi(apiClient *client.ApiClient) *PrechecksServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PrechecksServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Run pre-upgrade validation checks to verify that the cluster and its entities are ready for an upgrade. Prechecks detect conditions that would cause an upgrade to fail, such as insufficient disk space, cluster health issues, incompatible configurations, or pinned VMs that cannot be migrated during host maintenance mode. Submit a PrechecksSpec containing the entityUpdateSpecs (entity UUID and target version pairs) that you intend to use for the upgrade. The entityUpdateSpecs should match the output from POST /$actions/compute-recommendations or be manually constructed from GET /entities data. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; inspect the task result to determine whether all prechecks passed. If prechecks fail, address the reported issues before proceeding to POST /$actions/upgrade. Use the skippedPrecheckFlags field to bypass specific prechecks (e.g. the pinned-VM check). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *PrechecksApi) PerformPrechecks(body *import6.PrechecksSpec, xClusterId *string, dryrun_ *bool, args ...map[string]interface{}) (*import7.PrechecksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewPrechecksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformPrechecks(context.Background(), &import20.PerformPrechecksRequest{
		Body:       body,
		XClusterId: xClusterId,
		Dryrun_:    dryrun_,
	}, args...)
}

// Run pre-upgrade validation checks to verify that the cluster and its entities are ready for an upgrade. Prechecks detect conditions that would cause an upgrade to fail, such as insufficient disk space, cluster health issues, incompatible configurations, or pinned VMs that cannot be migrated during host maintenance mode. Submit a PrechecksSpec containing the entityUpdateSpecs (entity UUID and target version pairs) that you intend to use for the upgrade. The entityUpdateSpecs should match the output from POST /$actions/compute-recommendations or be manually constructed from GET /entities data. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; inspect the task result to determine whether all prechecks passed. If prechecks fail, address the reported issues before proceeding to POST /$actions/upgrade. Use the skippedPrecheckFlags field to bypass specific prechecks (e.g. the pinned-VM check). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *PrechecksServiceApi) PerformPrechecks(ctx context.Context, request *import20.PerformPrechecksRequest, args ...map[string]interface{}) (*import7.PrechecksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/prechecks"

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
	unmarshalledResp := new(import7.PrechecksApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
