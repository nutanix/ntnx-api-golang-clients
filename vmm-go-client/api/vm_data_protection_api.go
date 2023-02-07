//Api classes for vmm's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	"net/http"
	"net/url"
	"strings"
)

type VmDataProtectionApi struct {
	ApiClient *client.ApiClient
}

func NewVmDataProtectionApi(apiClient *client.ApiClient) *VmDataProtectionApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmDataProtectionApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point.
  Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point.

  parameters:-
  -> body (vmm.v4.ahv.config.VmRevert) (required) : Specify the VM Recovery Point Id to which the VM would be reverted.
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.RevertVmResponse, error)
*/
func (api *VmDataProtectionApi) RevertVm(body *import2.VmRevert, extId *string, args ...map[string]interface{}) (*import2.RevertVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/revert"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
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
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.RevertVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
