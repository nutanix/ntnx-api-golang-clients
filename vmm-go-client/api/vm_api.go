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

type VmApi struct {
	ApiClient *client.ApiClient
}

func NewVmApi(apiClient *client.ApiClient) *VmApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  IP address assignment to a VM NIC.
  IP address assignment to a VM NIC.

  parameters:-
  -> body (vmm.v4.ahv.config.AssignIp) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.AssignIpResponse, error)
*/
func (api *VmApi) AssignIp(body *import2.AssignIp, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.AssignIpResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/assign-ip"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.AssignIpResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Attach a category for the VM.
  Attach a category for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.CategoryReference) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.AttachVmCategoryResponse, error)
*/
func (api *VmApi) AttachVmCategory(body *import2.CategoryReference, extId *string, args ...map[string]interface{}) (*import2.AttachVmCategoryResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/categories/$actions/attach"

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
	unmarshalledResp := new(import2.AttachVmCategoryResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Clone a VM
  Clone a VM

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> body (vmm.v4.ahv.config.CloneOverride) (optional)
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CloneVmResponse, error)
*/
func (api *VmApi) CloneVm(extId *string, body *import2.CloneOverride, args ...map[string]interface{}) (*import2.CloneVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/clone"

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
	unmarshalledResp := new(import2.CloneVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a CD-ROM for the VM.
  Create a CD-ROM for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Cdrom) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateCdromResponse, error)
*/
func (api *VmApi) CreateCdrom(body *import2.Cdrom, vmExtId *string, args ...map[string]interface{}) (*import2.CreateCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.CreateCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a Disk for the VM.
  Create a Disk for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Disk) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateDiskResponse, error)
*/
func (api *VmApi) CreateDisk(body *import2.Disk, vmExtId *string, args ...map[string]interface{}) (*import2.CreateDiskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.CreateDiskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a GPU for the VM.
  Create a GPU for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Gpu) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateGpuResponse, error)
*/
func (api *VmApi) CreateGpu(body *import2.Gpu, vmExtId *string, args ...map[string]interface{}) (*import2.CreateGpuResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.CreateGpuResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a NIC for the VM.
  Create a NIC for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Nic) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateNicResponse, error)
*/
func (api *VmApi) CreateNic(body *import2.Nic, vmExtId *string, args ...map[string]interface{}) (*import2.CreateNicResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.CreateNicResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a SerialPort for the VM.
  Create a SerialPort for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.SerialPort) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateSerialPortResponse, error)
*/
func (api *VmApi) CreateSerialPort(body *import2.SerialPort, vmExtId *string, args ...map[string]interface{}) (*import2.CreateSerialPortResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.CreateSerialPortResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create a VM.
  Create a VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Vm) (required)
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CreateVmResponse, error)
*/
func (api *VmApi) CreateVm(body *import2.Vm, args ...map[string]interface{}) (*import2.CreateVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms"

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
	unmarshalledResp := new(import2.CreateVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Customize a guest VM.
  Customize a guest VM.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestCustomization) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.CustomizeGuestVmResponse, error)
*/
func (api *VmApi) CustomizeGuestVm(body *import2.GuestCustomization, extId *string, args ...map[string]interface{}) (*import2.CustomizeGuestVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/customize-guest"

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
	unmarshalledResp := new(import2.CustomizeGuestVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete CD-ROM for the VM.
  Delete CD-ROM for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a CD-ROM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteCdromResponse, error)
*/
func (api *VmApi) DeleteCdrom(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.DeleteCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete Disk for the VM.
  Delete Disk for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM disk. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteDiskResponse, error)
*/
func (api *VmApi) DeleteDisk(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.DeleteDiskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteDiskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete GPU for the VM.
  Delete GPU for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM GPU. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteGpuResponse, error)
*/
func (api *VmApi) DeleteGpu(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.DeleteGpuResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteGpuResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete NIC for the VM.
  Delete NIC for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteNicResponse, error)
*/
func (api *VmApi) DeleteNic(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.DeleteNicResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteNicResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete SerialPort for the VM.
  Delete SerialPort for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a serial port. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteSerialPortResponse, error)
*/
func (api *VmApi) DeleteSerialPort(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.DeleteSerialPortResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteSerialPortResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete a VM.
  Delete a VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DeleteVmResponse, error)
*/
func (api *VmApi) DeleteVm(extId *string, args ...map[string]interface{}) (*import2.DeleteVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.DeleteVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Detach a category from the VM.
  Detach a category from the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.CategoryReference) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.DetachVmCategoryResponse, error)
*/
func (api *VmApi) DetachVmCategory(body *import2.CategoryReference, extId *string, args ...map[string]interface{}) (*import2.DetachVmCategoryResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/categories/$actions/detach"

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
	unmarshalledResp := new(import2.DetachVmCategoryResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update CD-ROM for the VM.
  Update CD-ROM for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a CD-ROM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.EjectCdromResponse, error)
*/
func (api *VmApi) EjectCdrom(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.EjectCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}/$actions/eject"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.EjectCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get CD-ROM for the VM.
  Get CD-ROM for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a CD-ROM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetCdromResponse, error)
*/
func (api *VmApi) GetCdromByExtId(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.GetCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get Disk for the VM.
  Get Disk for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM disk. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetDiskResponse, error)
*/
func (api *VmApi) GetDiskByExtId(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.GetDiskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetDiskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get GPU for the VM.
  Get GPU for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM GPU. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetGpuResponse, error)
*/
func (api *VmApi) GetGpuByExtId(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.GetGpuResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetGpuResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get Nutanix Guest Tools information for a VM.
  Get Nutanix Guest Tools information for a VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetGuestToolsResponse, error)
*/
func (api *VmApi) GetGuestTools(extId *string, args ...map[string]interface{}) (*import2.GetGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools-config"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get NIC for the VM identified by the VM external Id.
  Get NIC for the VM identified by the VM external Id.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetNicResponse, error)
*/
func (api *VmApi) GetNicByExtId(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.GetNicResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetNicResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get SerialPort for the VM.
  Get SerialPort for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a serial port. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetSerialPortResponse, error)
*/
func (api *VmApi) GetSerialPortByExtId(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.GetSerialPortResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetSerialPortResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get VM of the provided VM external Id.
  Get VM of the provided VM external Id.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetVmResponse, error)
*/
func (api *VmApi) GetVmByExtId(extId *string, args ...map[string]interface{}) (*import2.GetVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get categories for the VM.
  Get categories for the VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetVmCategoriesResponse, error)
*/
func (api *VmApi) GetVmCategories(extId *string, args ...map[string]interface{}) (*import2.GetVmCategoriesResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/categories"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetVmCategoriesResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get ownership information for the VM identified by the VM external Id.
  Get ownership information for the VM identified by the VM external Id.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.GetVmOwnershipInfoResponse, error)
*/
func (api *VmApi) GetVmOwnershipInfo(extId *string, args ...map[string]interface{}) (*import2.GetVmOwnershipInfoResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/ownership-info"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.GetVmOwnershipInfoResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Restart the VM using Nutanix Guest Tools.
  Restart the VM using Nutanix Guest Tools.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestPowerOptions) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.RebootVmResponse, error)
*/
func (api *VmApi) GuestRebootVm(body *import2.GuestPowerOptions, extId *string, args ...map[string]interface{}) (*import2.RebootVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/guest-reboot"

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
	unmarshalledResp := new(import2.RebootVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Shutdown the VM using Nutanix Guest Tools.
  Shutdown the VM using Nutanix Guest Tools.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestPowerOptions) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ShutdownVmResponse, error)
*/
func (api *VmApi) GuestShutdownVm(body *import2.GuestPowerOptions, extId *string, args ...map[string]interface{}) (*import2.ShutdownVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/guest-shutdown"

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
	unmarshalledResp := new(import2.ShutdownVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Insert an ISO image or VM disk to a CD-ROM for the VM.
  Insert an ISO image or VM disk to a CD-ROM for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.CdromInsert) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a CD-ROM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.InsertCdromResponse, error)
*/
func (api *VmApi) InsertCdrom(body *import2.CdromInsert, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.InsertCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}/$actions/insert"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.InsertCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Insert Nutanix Guest Tools ISO into an available slot.
  Insert Nutanix Guest Tools ISO into an available slot.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestToolsInsertConfig) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.InsertVmGuestToolsResponse, error)
*/
func (api *VmApi) InsertVmGuestTools(body *import2.GuestToolsInsertConfig, extId *string, args ...map[string]interface{}) (*import2.InsertVmGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/insert-iso"

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
	unmarshalledResp := new(import2.InsertVmGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Install Nutanix Guest Tools.
  Install Nutanix Guest Tools.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestToolsInstallConfig) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.InstallVmGuestToolsResponse, error)
*/
func (api *VmApi) InstallVmGuestTools(body *import2.GuestToolsInstallConfig, extId *string, args ...map[string]interface{}) (*import2.InstallVmGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/install"

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
	unmarshalledResp := new(import2.InstallVmGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List CD-ROMs for the VM.
  List CD-ROMs for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListCdromsResponse, error)
*/
func (api *VmApi) ListCdroms(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import2.ListCdromsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ListCdromsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List Disks for the VM.
  List Disks for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListDisksResponse, error)
*/
func (api *VmApi) ListDisks(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import2.ListDisksResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ListDisksResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List GPUs for the VM.
  List GPUs for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListGpusResponse, error)
*/
func (api *VmApi) ListGpus(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import2.ListGpusResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/gpus"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ListGpusResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List NICs for the VM.
  List NICs for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListNicsResponse, error)
*/
func (api *VmApi) ListNics(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import2.ListNicsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ListNicsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List SerialPort for the VM.
  List SerialPort for the VM.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListSerialPortsResponse, error)
*/
func (api *VmApi) ListSerialPorts(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import2.ListSerialPortsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ListSerialPortsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List VMs.
  List VMs.

  parameters:-
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - cluster/extId - cpuPassthroughEnabled - gpuConsoleEnabled - hardwareClockTimezone - host/extId - isAgentVm - liveMigrateCapable - machineType - memoryOvercommitEnabled - memorySizeBytes - name - numCoresPerSocket - numSockets - numThreadsPerCore - numVnumaNodes - powerState - vgaConsoleEnabled
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - memorySizeBytes - name - numCoresPerSocket - numSockets - numThreadsPerCore - numVnumaNodes
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ListVmsResponse, error)
*/
func (api *VmApi) ListVms(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import2.ListVmsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms"

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
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
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
	unmarshalledResp := new(import2.ListVmsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Migrate a VM NIC.
  Migrate a VM NIC.

  parameters:-
  -> body (vmm.v4.ahv.config.MigrateNicConfig) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.MigrateNicResponse, error)
*/
func (api *VmApi) MigrateNic(body *import2.MigrateNicConfig, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.MigrateNicResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/migrate"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.MigrateNicResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Restart the VM.
  Restart the VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.PowerCycleVmResponse, error)
*/
func (api *VmApi) PowerCycleVm(extId *string, args ...map[string]interface{}) (*import2.PowerCycleVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-cycle"

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
	unmarshalledResp := new(import2.PowerCycleVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Turn off the VM.
  Turn off the VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.PowerOffVmResponse, error)
*/
func (api *VmApi) PowerOffVm(extId *string, args ...map[string]interface{}) (*import2.PowerOffVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-off"

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
	unmarshalledResp := new(import2.PowerOffVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Turn on the VM.
  Turn on the VM.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.PowerOnVmResponse, error)
*/
func (api *VmApi) PowerOnVm(extId *string, args ...map[string]interface{}) (*import2.PowerOnVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/power-on"

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
	unmarshalledResp := new(import2.PowerOnVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Restart the VM using ACPI.
  Restart the VM using ACPI.

  parameters:-
  -> body (vmm.v4.ahv.config.ACPIPowerOptions) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.RebootVmResponse, error)
*/
func (api *VmApi) RebootVm(body *import2.ACPIPowerOptions, extId *string, args ...map[string]interface{}) (*import2.RebootVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/reboot"

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
	unmarshalledResp := new(import2.RebootVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Release an IP address from a VM NIC.
  Release an IP address from a VM NIC.

  parameters:-
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ReleaseIpResponse, error)
*/
func (api *VmApi) ReleaseIp(vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.ReleaseIpResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/release-ip"

	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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
	unmarshalledResp := new(import2.ReleaseIpResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Reset the VM immediately without waiting for the guest VM to shutdown itself.
  Reset the VM immediately without waiting for the guest VM to shutdown itself.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ResetVmResponse, error)
*/
func (api *VmApi) ResetVm(extId *string, args ...map[string]interface{}) (*import2.ResetVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/reset"

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
	unmarshalledResp := new(import2.ResetVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Shutdown the VM using ACPI.
  Shutdown the VM using ACPI.

  parameters:-
  -> body (vmm.v4.ahv.config.ACPIPowerOptions) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.ShutdownVmResponse, error)
*/
func (api *VmApi) ShutdownVm(body *import2.ACPIPowerOptions, extId *string, args ...map[string]interface{}) (*import2.ShutdownVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/shutdown"

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
	unmarshalledResp := new(import2.ShutdownVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Uninstall Nutanix Guest Tools.
  Uninstall Nutanix Guest Tools.

  parameters:-
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UninstallVmGuestToolsResponse, error)
*/
func (api *VmApi) UninstallVmGuestTools(extId *string, args ...map[string]interface{}) (*import2.UninstallVmGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/uninstall"

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
	unmarshalledResp := new(import2.UninstallVmGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update CD-ROM for the VM.
  Update CD-ROM for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Cdrom) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a CD-ROM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateCdromResponse, error)
*/
func (api *VmApi) UpdateCdrom(body *import2.Cdrom, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.UpdateCdromResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/cdroms/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateCdromResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update Disk for the VM.
  Update Disk for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Disk) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM disk. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateDiskResponse, error)
*/
func (api *VmApi) UpdateDisk(body *import2.Disk, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.UpdateDiskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateDiskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update Nutanix Guest Tools information for a VM.
  Update Nutanix Guest Tools information for a VM.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestTools) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateGuestToolsResponse, error)
*/
func (api *VmApi) UpdateGuestTools(body *import2.GuestTools, extId *string, args ...map[string]interface{}) (*import2.UpdateGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools-config"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update NIC for the VM.
  Update NIC for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Nic) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a VM NIC. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateNicResponse, error)
*/
func (api *VmApi) UpdateNic(body *import2.Nic, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.UpdateNicResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateNicResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update SerialPort for the VM.
  Update SerialPort for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.SerialPort) (required)
  -> vmExtId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> extId (string) (required) : Globally unique identifier of a serial port. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateSerialPortResponse, error)
*/
func (api *VmApi) UpdateSerialPort(body *import2.SerialPort, vmExtId *string, extId *string, args ...map[string]interface{}) (*import2.UpdateSerialPortResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'vmExtId' is set
	if nil == vmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*vmExtId, "")), -1)
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateSerialPortResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update VM.
  Update VM.

  parameters:-
  -> body (vmm.v4.ahv.config.Vm) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateVmResponse, error)
*/
func (api *VmApi) UpdateVm(body *import2.Vm, extId *string, args ...map[string]interface{}) (*import2.UpdateVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update categories for the VM.
  Update categories for the VM.

  parameters:-
  -> body ([]vmm.v4.ahv.config.CategoryReference) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateVmCategoriesResponse, error)
*/
func (api *VmApi) UpdateVmCategories(body *[]import2.CategoryReference, extId *string, args ...map[string]interface{}) (*import2.UpdateVmCategoriesResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/categories"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateVmCategoriesResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update ownership information for the VM.
  Update ownership information for the VM.

  parameters:-
  -> body (vmm.v4.ahv.config.OwnershipInfo) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpdateVmOwnershipInfoResponse, error)
*/
func (api *VmApi) UpdateVmOwnershipInfo(body *import2.OwnershipInfo, extId *string, args ...map[string]interface{}) (*import2.UpdateVmOwnershipInfoResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/ownership-info"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.UpdateVmOwnershipInfoResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Upgrade Nutanix Guest Tools.
  Upgrade Nutanix Guest Tools.

  parameters:-
  -> body (vmm.v4.ahv.config.GuestToolsUpgradeConfig) (required)
  -> extId (string) (required) : Globally unique identifier of a VM. It should be of type UUID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*vmm.v4.ahv.config.UpgradeVmGuestToolsResponse, error)
*/
func (api *VmApi) UpgradeVmGuestTools(body *import2.GuestToolsUpgradeConfig, extId *string, args ...map[string]interface{}) (*import2.UpgradeVmGuestToolsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/guest-tools/$actions/upgrade"

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
	unmarshalledResp := new(import2.UpgradeVmGuestToolsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
