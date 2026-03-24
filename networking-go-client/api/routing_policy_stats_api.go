package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import26 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/routingpolicystats"
	import13 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/stats"
	"net/http"
	"net/url"
	"strings"
)

type RoutingPolicyStatsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RoutingPolicyStatsServiceApi
}

type RoutingPolicyStatsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRoutingPolicyStatsApi(apiClient *client.ApiClient) *RoutingPolicyStatsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutingPolicyStatsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRoutingPolicyStatsServiceApi(a.ApiClient)

	return a
}

func NewRoutingPolicyStatsServiceApi(apiClient *client.ApiClient) *RoutingPolicyStatsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutingPolicyStatsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Clear the value in packet and byte counters of all Routing Policies in the chosen VPC or a particular routing policy in the chosen VPC.
func (api *RoutingPolicyStatsApi) ClearRoutingPolicyCounters(body *import13.RoutingPolicyClearCountersSpec, args ...map[string]interface{}) (*import13.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoutingPolicyStatsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ClearRoutingPolicyCounters(context.Background(), &import26.ClearRoutingPolicyCountersRequest{
		Body: body,
	}, args...)
}

// Clear the value in packet and byte counters of all Routing Policies in the chosen VPC or a particular routing policy in the chosen VPC.
func (api *RoutingPolicyStatsServiceApi) ClearRoutingPolicyCounters(ctx context.Context, request *import26.ClearRoutingPolicyCountersRequest, args ...map[string]interface{}) (*import13.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/stats/routing-policies/$actions/clear"

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

	unmarshalledResp := new(import13.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
