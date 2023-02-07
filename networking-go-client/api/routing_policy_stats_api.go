//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/stats"
	"net/http"
	"net/url"
	"strings"
)

type RoutingPolicyStatsApi struct {
	ApiClient *client.ApiClient
}

func NewRoutingPolicyStatsApi(apiClient *client.ApiClient) *RoutingPolicyStatsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoutingPolicyStatsApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Clear the value in packet and byte counters of all Routing Policies in the chosen VPC.
  Clear the value in the packet and byte counters for  all routing policies in the chosen VPC.

  parameters:-
  -> body (networking.v4.stats.RoutingPolicyClearCountersBody) (required) : VPC UUID to reset all routing policy counters to zero.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.stats.RoutingPolicyClearCountersApiResponse, error)
*/
func (api *RoutingPolicyStatsApi) ClearAllRoutingPolicyCounter(body *import3.RoutingPolicyClearCountersBody, args ...map[string]interface{}) (*import3.RoutingPolicyClearCountersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/stats/routing-policies/$actions/clear"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.RoutingPolicyClearCountersApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Clear the value in packet and byte counters of chosen Routing Policy.
  Clear the value in packet and byte counters of the chosen routing policy.

  parameters:-
  -> extId (string) (required) : ExtId of the Routing Policy.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.stats.RoutingPolicyClearCountersApiResponse, error)
*/
func (api *RoutingPolicyStatsApi) ClearRoutingPolicyCounter(extId *string, args ...map[string]interface{}) (*import3.RoutingPolicyClearCountersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/stats/routing-policies/{extId}/$actions/clear"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.RoutingPolicyClearCountersApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
