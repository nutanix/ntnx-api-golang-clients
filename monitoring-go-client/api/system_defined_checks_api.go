package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import8 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/systemdefinedchecks"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type SystemDefinedChecksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SystemDefinedChecksServiceApi
}

type SystemDefinedChecksServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSystemDefinedChecksApi(apiClient *client.ApiClient) *SystemDefinedChecksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SystemDefinedChecksApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSystemDefinedChecksServiceApi(a.ApiClient)

	return a
}

func NewSystemDefinedChecksServiceApi(apiClient *client.ApiClient) *SystemDefinedChecksServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SystemDefinedChecksServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Run System-Defined Checks on a cluster.
func (api *SystemDefinedChecksApi) RunSystemDefinedChecks(clusterExtId *string, body *import1.RunSystemDefinedChecksSpec, args ...map[string]interface{}) (*import1.RunSystemDefinedChecksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedChecksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RunSystemDefinedChecks(context.Background(), &import8.RunSystemDefinedChecksRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Run System-Defined Checks on a cluster.
func (api *SystemDefinedChecksServiceApi) RunSystemDefinedChecks(ctx context.Context, request *import8.RunSystemDefinedChecksRequest, args ...map[string]interface{}) (*import1.RunSystemDefinedChecksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/clusters/{clusterExtId}/$actions/run-system-defined-checks"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import1.RunSystemDefinedChecksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
