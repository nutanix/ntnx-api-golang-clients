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

type VpnVendorConfigApi struct {
	ApiClient *client.ApiClient
}

func NewVpnVendorConfigApi(apiClient *client.ApiClient) *VpnVendorConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpnVendorConfigApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Get VPN vendor device configuration steps.
  Get VPN vendor device configuration steps. If device version is not specified the configuration steps of the latest available device version are returned.

  parameters:-
  -> vpnConnectionId (string) (required) : VPN connection UUID.
  -> vendorName (string) (required) : VPN device vendor name.
  -> deviceVersion (string) (optional) : VPN device version.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*string, error)
*/
func (api *VpnVendorConfigApi) GetVpnVendorConfig(vpnConnectionId *string, vendorName *string, deviceVersion *string, args ...map[string]interface{}) (*string, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/vpn-vendor-configs/{vpnConnectionId}"

	// verify the required parameter 'vpnConnectionId' is set
	if nil == vpnConnectionId {
		return nil, client.ReportError("vpnConnectionId is required and must be specified")
	}
	// verify the required parameter 'vendorName' is set
	if nil == vendorName {
		return nil, client.ReportError("vendorName is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpnConnectionId"+"}", url.PathEscape(client.ParameterToString(*vpnConnectionId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"text/plain", "application/json"}

	// Query Params
	queryParams.Add("vendorName", client.ParameterToString(*vendorName, ""))
	if deviceVersion != nil {
		queryParams.Add("deviceVersion", client.ParameterToString(*deviceVersion, ""))
	}

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
	unmarshalledResp := new(string)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List third-party VPN vendors and devices for which configuration steps are available to download.
  List third-party VPN vendors and devices for which configuration steps are available to download.

  parameters:-
  -> vpnConnectionId (string) (required) : VPN connection UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.VpnVendorListApiResponse, error)
*/
func (api *VpnVendorConfigApi) ListAvailableVpnVendorConfigs(vpnConnectionId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.VpnVendorListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/vpn-vendor-configs/{vpnConnectionId}/available"

	// verify the required parameter 'vpnConnectionId' is set
	if nil == vpnConnectionId {
		return nil, client.ReportError("vpnConnectionId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vpnConnectionId"+"}", url.PathEscape(client.ParameterToString(*vpnConnectionId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}

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
	unmarshalledResp := new(import1.VpnVendorListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
