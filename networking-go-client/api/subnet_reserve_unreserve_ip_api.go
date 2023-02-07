//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type SubnetReserveUnreserveIpApi struct {
	ApiClient *client.ApiClient
}

func NewSubnetReserveUnreserveIpApi(apiClient *client.ApiClient) *SubnetReserveUnreserveIpApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SubnetReserveUnreserveIpApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Get the list of assigned and reserved IP addresses on a subnet.
  Get the list of assigned and reserved IP addresses on a subnet.

  parameters:-
  -> subnetExtId (string) (required) : Subnet ExtID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.SubnetAddressAssignmentListApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) FetchSubnetAddressAssignments(subnetExtId *string, args ...map[string]interface{}) (*import1.SubnetAddressAssignmentListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses"

	// verify the required parameter 'subnetExtId' is set
	if nil == subnetExtId {
		return nil, client.ReportError("subnetExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(*subnetExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.SubnetAddressAssignmentListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Reserve IP addresses on a subnet.
  Reserve IP addresses on a subnet.

  parameters:-
  -> body (networking.v4.config.IpReserveInput) (required) : Request schema to reserve IP addresses on a subnet.
  -> subnetExtId (string) (required) : The UUID of the subnet.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) ReserveIps(body *import1.IpReserveInput, subnetExtId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses/$actions/reserve"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'subnetExtId' is set
	if nil == subnetExtId {
		return nil, client.ReportError("subnetExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(*subnetExtId, "")), -1)
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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Unreserve IP addresses on a subnet.
  Unreserve IP addresses on a subnet.

  parameters:-
  -> body (networking.v4.config.IpUnreserveInput) (required) : Request schema to unreserve IP addresses on a subnet.
  -> subnetExtId (string) (required) : The UUID of the subnet.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) UnreserveIps(body *import1.IpUnreserveInput, subnetExtId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses/$actions/unreserve"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'subnetExtId' is set
	if nil == subnetExtId {
		return nil, client.ReportError("subnetExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(*subnetExtId, "")), -1)
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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
