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

type VirtualSwitchApi struct {
	ApiClient *client.ApiClient
}

func NewVirtualSwitchApi(apiClient *client.ApiClient) *VirtualSwitchApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VirtualSwitchApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Create a Virtual Switch
  Create a Virtual Switch

  parameters:-
  -> body (networking.v4.config.VirtualSwitch) (required) : Schema to configure a virtual switch
  -> xClusterId (string) (required) : Prism Element Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *VirtualSwitchApi) CreateVirtualSwitch(body *import1.VirtualSwitch, xClusterId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/virtual-switches"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")

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
  Delete a Virtual Switch
  Delete a Virtual Switch

  parameters:-
  -> xClusterId (string) (required) : Prism Element Cluster UUID
  -> extId (string) (required) : UUID of Virtual Switch
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *VirtualSwitchApi) DeleteVirtualSwitch(xClusterId *string, extId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/virtual-switches/{extId}"

	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}
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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get single Virtual Switch given its UUID
  Get single Virtual Switch given its UUID

  parameters:-
  -> xClusterId (string) (required) : Prism Element Cluster UUID
  -> extId (string) (required) : UUID of Virtual Switch
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.VirtualSwitchApiResponse, error)
*/
func (api *VirtualSwitchApi) GetVirtualSwitch(xClusterId *string, extId *string, args ...map[string]interface{}) (*import1.VirtualSwitchApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/virtual-switches/{extId}"

	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}
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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.VirtualSwitchApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get list of Virtual Switches
  Get list of Virtual Switches

  parameters:-
  -> xClusterId (string) (required) : Prism Element Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.VirtualSwitchListApiResponse, error)
*/
func (api *VirtualSwitchApi) ListVirtualSwitches(xClusterId *string, args ...map[string]interface{}) (*import1.VirtualSwitchListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/virtual-switches"

	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}

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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.VirtualSwitchListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update a Virtual Switch
  Update a Virtual Switch

  parameters:-
  -> body (networking.v4.config.VirtualSwitch) (required) : Schema to configure a virtual switch
  -> xClusterId (string) (required) : Prism Element Cluster UUID
  -> extId (string) (required) : UUID of Virtual Switch
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *VirtualSwitchApi) UpdateVirtualSwitch(body *import1.VirtualSwitch, xClusterId *string, extId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/virtual-switches/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}
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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
