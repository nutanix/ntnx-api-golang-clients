package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import19 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	import20 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vm"
	"net/http"
	"net/url"
	"strings"
)

type VmApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmServiceApi
}

type VmServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmApi(apiClient *client.ApiClient) *VmApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmServiceApi(a.ApiClient)

	return a
}

func NewVmServiceApi(apiClient *client.ApiClient) *VmServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Adds to the custom attributes of the VM.
func (api *VmApi) AddVmCustomAttributes(extId *string, body *import19.UpdateCustomAttributesParams, args ...map[string]interface{}) (*import19.AddVmCustomAttributesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddVmCustomAttributes(context.Background(), &import20.AddVmCustomAttributesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Adds to the custom attributes of the VM.
func (api *VmServiceApi) AddVmCustomAttributes(ctx context.Context, request *import20.AddVmCustomAttributesRequest, args ...map[string]interface{}) (*import19.AddVmCustomAttributesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/add-custom-attributes"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.AddVmCustomAttributesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Adds to the custom attributes of the VM disk entity.
func (api *VmApi) AddVmDiskCustomAttributes(vmExtId *string, extId *string, body *import19.UpdateCustomAttributesParams, args ...map[string]interface{}) (*import19.AddVmDiskCustomAttributesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddVmDiskCustomAttributes(context.Background(), &import20.AddVmDiskCustomAttributesRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Adds to the custom attributes of the VM disk entity.
func (api *VmServiceApi) AddVmDiskCustomAttributes(ctx context.Context, request *import20.AddVmDiskCustomAttributesRequest, args ...map[string]interface{}) (*import19.AddVmDiskCustomAttributesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}/$actions/add-custom-attributes"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.AddVmDiskCustomAttributesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Assigns a specific IP address to a network device attached to a managed network.
func (api *VmApi) AssignIpById(vmExtId *string, extId *string, body *import19.AssignIpParams, args ...map[string]interface{}) (*import19.AssignIpApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssignIpById(context.Background(), &import20.AssignIpByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Assigns a specific IP address to a network device attached to a managed network.
func (api *VmServiceApi) AssignIpById(ctx context.Context, request *import20.AssignIpByIdRequest, args ...map[string]interface{}) (*import19.AssignIpApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/assign-ip"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.AssignIpApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Assign the owner of a virtual machine.
func (api *VmApi) AssignVmOwner(extId *string, body *import19.OwnershipInfo, args ...map[string]interface{}) (*import19.AssignVmOwnerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssignVmOwner(context.Background(), &import20.AssignVmOwnerRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Assign the owner of a virtual machine.
func (api *VmServiceApi) AssignVmOwner(ctx context.Context, request *import20.AssignVmOwnerRequest, args ...map[string]interface{}) (*import19.AssignVmOwnerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/assign-owner"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.AssignVmOwnerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Associate categories to a virtual machine.
func (api *VmApi) AssociateCategories(extId *string, body *import19.AssociateVmCategoriesParams, args ...map[string]interface{}) (*import19.AssociateCategoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociateCategories(context.Background(), &import20.AssociateCategoriesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Associate categories to a virtual machine.
func (api *VmServiceApi) AssociateCategories(ctx context.Context, request *import20.AssociateCategoriesRequest, args ...map[string]interface{}) (*import19.AssociateCategoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/associate-categories"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.AssociateCategoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Clones an existing virtual machine. This operation also includes updating the existing configuration during the clone operation.
func (api *VmApi) CloneVm(extId *string, body *import19.CloneOverrideParams, args ...map[string]interface{}) (*import19.CloneVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CloneVm(context.Background(), &import20.CloneVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Clones an existing virtual machine. This operation also includes updating the existing configuration during the clone operation.
func (api *VmServiceApi) CloneVm(ctx context.Context, request *import20.CloneVmRequest, args ...map[string]interface{}) (*import19.CloneVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/clone"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.CloneVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates and attaches a CD-ROM device to a virtual machine.
func (api *VmApi) CreateCdRom(vmExtId *string, body *import19.CdRom, args ...map[string]interface{}) (*import19.CreateCdRomApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCdRom(context.Background(), &import20.CreateCdRomRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Creates and attaches a CD-ROM device to a virtual machine.
func (api *VmServiceApi) CreateCdRom(ctx context.Context, request *import20.CreateCdRomRequest, args ...map[string]interface{}) (*import19.CreateCdRomApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreateCdRomApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates and attaches a disk device to a virtual machine.
func (api *VmApi) CreateDisk(vmExtId *string, body *import19.Disk, args ...map[string]interface{}) (*import19.CreateDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDisk(context.Background(), &import20.CreateDiskRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Creates and attaches a disk device to a virtual machine.
func (api *VmServiceApi) CreateDisk(ctx context.Context, request *import20.CreateDiskRequest, args ...map[string]interface{}) (*import19.CreateDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreateDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Attaches a GPU device to a virtual machine
func (api *VmApi) CreateGpu(vmExtId *string, body *import19.Gpu, args ...map[string]interface{}) (*import19.CreateGpuApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateGpu(context.Background(), &import20.CreateGpuRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Attaches a GPU device to a virtual machine
func (api *VmServiceApi) CreateGpu(ctx context.Context, request *import20.CreateGpuRequest, args ...map[string]interface{}) (*import19.CreateGpuApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/gpus"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreateGpuApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates and attaches a network device to a virtual machine.
func (api *VmApi) CreateNic(vmExtId *string, body *import19.Nic, args ...map[string]interface{}) (*import19.CreateNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateNic(context.Background(), &import20.CreateNicRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Creates and attaches a network device to a virtual machine.
func (api *VmServiceApi) CreateNic(ctx context.Context, request *import20.CreateNicRequest, args ...map[string]interface{}) (*import19.CreateNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreateNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates and attaches a PCIe device to a virtual machine.
func (api *VmApi) CreatePcieDevice(vmExtId *string, body *import19.PcieDevice, args ...map[string]interface{}) (*import19.CreatePcieDeviceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreatePcieDevice(context.Background(), &import20.CreatePcieDeviceRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Creates and attaches a PCIe device to a virtual machine.
func (api *VmServiceApi) CreatePcieDevice(ctx context.Context, request *import20.CreatePcieDeviceRequest, args ...map[string]interface{}) (*import19.CreatePcieDeviceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/pcie-devices"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreatePcieDeviceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates and attaches a serial port device to a virtual machine.
func (api *VmApi) CreateSerialPort(vmExtId *string, body *import19.SerialPort, args ...map[string]interface{}) (*import19.CreateSerialPortApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateSerialPort(context.Background(), &import20.CreateSerialPortRequest{
		VmExtId: vmExtId,
		Body:    body,
	}, args...)
}

// Creates and attaches a serial port device to a virtual machine.
func (api *VmServiceApi) CreateSerialPort(ctx context.Context, request *import20.CreateSerialPortRequest, args ...map[string]interface{}) (*import19.CreateSerialPortApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/serial-ports"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
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

	unmarshalledResp := new(import19.CreateSerialPortApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a Virtual Machine with the provided configuration.
func (api *VmApi) CreateVm(body *import19.Vm, args ...map[string]interface{}) (*import19.CreateVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVm(context.Background(), &import20.CreateVmRequest{
		Body: body,
	}, args...)
}

// Creates a Virtual Machine with the provided configuration.
func (api *VmServiceApi) CreateVm(ctx context.Context, request *import20.CreateVmRequest, args ...map[string]interface{}) (*import19.CreateVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms"

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

	unmarshalledResp := new(import19.CreateVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Migrates a VM from the current cluster to a new target cluster. Supports the dry-run option, which if used results only in validation of the cross-cluster migration parameters and pre-requisites.
func (api *VmApi) CrossClusterMigrateVm(extId *string, body *import19.VmCrossClusterMigrateParams, dryrun_ *bool, args ...map[string]interface{}) (*import19.CrossClusterMigrateVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CrossClusterMigrateVm(context.Background(), &import20.CrossClusterMigrateVmRequest{
		ExtId:   extId,
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Migrates a VM from the current cluster to a new target cluster. Supports the dry-run option, which if used results only in validation of the cross-cluster migration parameters and pre-requisites.
func (api *VmServiceApi) CrossClusterMigrateVm(ctx context.Context, request *import20.CrossClusterMigrateVmRequest, args ...map[string]interface{}) (*import19.CrossClusterMigrateVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/migrate"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.CrossClusterMigrateVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Stage a Sysprep or cloud-init configuration file to be used by the guest for the next boot. Note that the Sysprep command must be used to generalize the Windows VMs before triggering this API call.
func (api *VmApi) CustomizeGuestVm(extId *string, body *import19.GuestCustomizationParams, args ...map[string]interface{}) (*import19.CustomizeGuestVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CustomizeGuestVm(context.Background(), &import20.CustomizeGuestVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Stage a Sysprep or cloud-init configuration file to be used by the guest for the next boot. Note that the Sysprep command must be used to generalize the Windows VMs before triggering this API call.
func (api *VmServiceApi) CustomizeGuestVm(ctx context.Context, request *import20.CustomizeGuestVmRequest, args ...map[string]interface{}) (*import19.CustomizeGuestVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/customize-guest"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.CustomizeGuestVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified CD-ROM device from a virtual machine.
func (api *VmApi) DeleteCdRomById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeleteCdRomApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCdRomById(context.Background(), &import20.DeleteCdRomByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified CD-ROM device from a virtual machine.
func (api *VmServiceApi) DeleteCdRomById(ctx context.Context, request *import20.DeleteCdRomByIdRequest, args ...map[string]interface{}) (*import19.DeleteCdRomApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteCdRomApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified disk device from a virtual machine.
func (api *VmApi) DeleteDiskById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeleteDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDiskById(context.Background(), &import20.DeleteDiskByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified disk device from a virtual machine.
func (api *VmServiceApi) DeleteDiskById(ctx context.Context, request *import20.DeleteDiskByIdRequest, args ...map[string]interface{}) (*import19.DeleteDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified GPU device from a virtual machine
func (api *VmApi) DeleteGpuById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeleteGpuApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteGpuById(context.Background(), &import20.DeleteGpuByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified GPU device from a virtual machine
func (api *VmServiceApi) DeleteGpuById(ctx context.Context, request *import20.DeleteGpuByIdRequest, args ...map[string]interface{}) (*import19.DeleteGpuApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/gpus/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteGpuApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified network device from a virtual machine.
func (api *VmApi) DeleteNicById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeleteNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteNicById(context.Background(), &import20.DeleteNicByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified network device from a virtual machine.
func (api *VmServiceApi) DeleteNicById(ctx context.Context, request *import20.DeleteNicByIdRequest, args ...map[string]interface{}) (*import19.DeleteNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified PCIe device from a virtual machine.
func (api *VmApi) DeletePcieDeviceById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeletePcieDeviceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeletePcieDeviceById(context.Background(), &import20.DeletePcieDeviceByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified PCIe device from a virtual machine.
func (api *VmServiceApi) DeletePcieDeviceById(ctx context.Context, request *import20.DeletePcieDeviceByIdRequest, args ...map[string]interface{}) (*import19.DeletePcieDeviceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/pcie-devices/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeletePcieDeviceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes the specified serial port from a virtual machine.
func (api *VmApi) DeleteSerialPortById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DeleteSerialPortApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteSerialPortById(context.Background(), &import20.DeleteSerialPortByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Removes the specified serial port from a virtual machine.
func (api *VmServiceApi) DeleteSerialPortById(ctx context.Context, request *import20.DeleteSerialPortByIdRequest, args ...map[string]interface{}) (*import19.DeleteSerialPortApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteSerialPortApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the specified virtual machine.
func (api *VmApi) DeleteVmById(extId *string, args ...map[string]interface{}) (*import19.DeleteVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmById(context.Background(), &import20.DeleteVmByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the specified virtual machine.
func (api *VmServiceApi) DeleteVmById(ctx context.Context, request *import20.DeleteVmByIdRequest, args ...map[string]interface{}) (*import19.DeleteVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DeleteVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disable hydration for a VM CD-ROM. The hydration process will be stopped and the CD-ROM will not be migrated to the local storage container.
func (api *VmApi) DisableVmCdRomHydration(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DisableVmCdRomHydrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisableVmCdRomHydration(context.Background(), &import20.DisableVmCdRomHydrationRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Disable hydration for a VM CD-ROM. The hydration process will be stopped and the CD-ROM will not be migrated to the local storage container.
func (api *VmServiceApi) DisableVmCdRomHydration(ctx context.Context, request *import20.DisableVmCdRomHydrationRequest, args ...map[string]interface{}) (*import19.DisableVmCdRomHydrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/disable-hydration"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DisableVmCdRomHydrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disable hydration for a VM disk. The hydration process will be stopped and the disk will not be migrated to the local storage container.
func (api *VmApi) DisableVmDiskHydration(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.DisableVmDiskHydrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisableVmDiskHydration(context.Background(), &import20.DisableVmDiskHydrationRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Disable hydration for a VM disk. The hydration process will be stopped and the disk will not be migrated to the local storage container.
func (api *VmServiceApi) DisableVmDiskHydration(ctx context.Context, request *import20.DisableVmDiskHydrationRequest, args ...map[string]interface{}) (*import19.DisableVmDiskHydrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}/$actions/disable-hydration"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.DisableVmDiskHydrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disassociate categories from a virtual machine.
func (api *VmApi) DisassociateCategories(extId *string, body *import19.DisassociateVmCategoriesParams, args ...map[string]interface{}) (*import19.DisassociateCategoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisassociateCategories(context.Background(), &import20.DisassociateCategoriesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Disassociate categories from a virtual machine.
func (api *VmServiceApi) DisassociateCategories(ctx context.Context, request *import20.DisassociateCategoriesRequest, args ...map[string]interface{}) (*import19.DisassociateCategoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/disassociate-categories"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.DisassociateCategoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Ejects the ISO currently inserted into a CD-ROM device on a Virtual Machine.
func (api *VmApi) EjectCdRomById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.EjectCdRomApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.EjectCdRomById(context.Background(), &import20.EjectCdRomByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Ejects the ISO currently inserted into a CD-ROM device on a Virtual Machine.
func (api *VmServiceApi) EjectCdRomById(ctx context.Context, request *import20.EjectCdRomByIdRequest, args ...map[string]interface{}) (*import19.EjectCdRomApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/eject"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.EjectCdRomApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Enable a VM CD-ROM to be hydrated to the local storage container in the background. If the VM CD-ROM hydration has been disabled, this will resume the hydration process.
func (api *VmApi) EnableVmCdRomHydration(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.EnableVmCdRomHydrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.EnableVmCdRomHydration(context.Background(), &import20.EnableVmCdRomHydrationRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Enable a VM CD-ROM to be hydrated to the local storage container in the background. If the VM CD-ROM hydration has been disabled, this will resume the hydration process.
func (api *VmServiceApi) EnableVmCdRomHydration(ctx context.Context, request *import20.EnableVmCdRomHydrationRequest, args ...map[string]interface{}) (*import19.EnableVmCdRomHydrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/enable-hydration"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.EnableVmCdRomHydrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Enable a VM disk to be hydrated to the local storage container in the background. If the VM disk hydration has been disabled, this will resume the hydration process.
func (api *VmApi) EnableVmDiskHydration(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.EnableVmDiskHydrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.EnableVmDiskHydration(context.Background(), &import20.EnableVmDiskHydrationRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Enable a VM disk to be hydrated to the local storage container in the background. If the VM disk hydration has been disabled, this will resume the hydration process.
func (api *VmServiceApi) EnableVmDiskHydration(ctx context.Context, request *import20.EnableVmDiskHydrationRequest, args ...map[string]interface{}) (*import19.EnableVmDiskHydrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}/$actions/enable-hydration"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.EnableVmDiskHydrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Generates a token to launch a VM console.
func (api *VmApi) GenerateConsoleTokenById(extId *string, args ...map[string]interface{}) (*import19.GenerateConsoleTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GenerateConsoleTokenById(context.Background(), &import20.GenerateConsoleTokenByIdRequest{
		ExtId: extId,
	}, args...)
}

// Generates a token to launch a VM console.
func (api *VmServiceApi) GenerateConsoleTokenById(ctx context.Context, request *import20.GenerateConsoleTokenByIdRequest, args ...map[string]interface{}) (*import19.GenerateConsoleTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/generate-console-token"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GenerateConsoleTokenApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided CD-ROM device attached to a virtual machine.
func (api *VmApi) GetCdRomById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetCdRomApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCdRomById(context.Background(), &import20.GetCdRomByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided CD-ROM device attached to a virtual machine.
func (api *VmServiceApi) GetCdRomById(ctx context.Context, request *import20.GetCdRomByIdRequest, args ...map[string]interface{}) (*import19.GetCdRomApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetCdRomApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided disk device attached to a virtual machine.
func (api *VmApi) GetDiskById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDiskById(context.Background(), &import20.GetDiskByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided disk device attached to a virtual machine.
func (api *VmServiceApi) GetDiskById(ctx context.Context, request *import20.GetDiskByIdRequest, args ...map[string]interface{}) (*import19.GetDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided GPU device attached to a virtual machine.
func (api *VmApi) GetGpuById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetGpuApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetGpuById(context.Background(), &import20.GetGpuByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided GPU device attached to a virtual machine.
func (api *VmServiceApi) GetGpuById(ctx context.Context, request *import20.GetGpuByIdRequest, args ...map[string]interface{}) (*import19.GetGpuApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/gpus/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetGpuApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the Nutanix Guest Tools configuration for a virtual machine.
func (api *VmApi) GetGuestToolsById(extId *string, args ...map[string]interface{}) (*import19.GetGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetGuestToolsById(context.Background(), &import20.GetGuestToolsByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the Nutanix Guest Tools configuration for a virtual machine.
func (api *VmServiceApi) GetGuestToolsById(ctx context.Context, request *import20.GetGuestToolsByIdRequest, args ...map[string]interface{}) (*import19.GetGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided network device attached to a virtual machine.
func (api *VmApi) GetNicById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNicById(context.Background(), &import20.GetNicByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided network device attached to a virtual machine.
func (api *VmServiceApi) GetNicById(ctx context.Context, request *import20.GetNicByIdRequest, args ...map[string]interface{}) (*import19.GetNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided PCIe device attached to a virtual machine.
func (api *VmApi) GetPcieDeviceById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetPcieDeviceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetPcieDeviceById(context.Background(), &import20.GetPcieDeviceByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided PCIe device attached to a virtual machine.
func (api *VmServiceApi) GetPcieDeviceById(ctx context.Context, request *import20.GetPcieDeviceByIdRequest, args ...map[string]interface{}) (*import19.GetPcieDeviceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/pcie-devices/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetPcieDeviceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for the provided serial port attached to a virtual machine.
func (api *VmApi) GetSerialPortById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.GetSerialPortApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSerialPortById(context.Background(), &import20.GetSerialPortByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Retrieves configuration details for the provided serial port attached to a virtual machine.
func (api *VmServiceApi) GetSerialPortById(ctx context.Context, request *import20.GetSerialPortByIdRequest, args ...map[string]interface{}) (*import19.GetSerialPortApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetSerialPortApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for a virtual machine.
func (api *VmApi) GetVmById(extId *string, args ...map[string]interface{}) (*import19.GetVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmById(context.Background(), &import20.GetVmByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves configuration details for a virtual machine.
func (api *VmServiceApi) GetVmById(ctx context.Context, request *import20.GetVmByIdRequest, args ...map[string]interface{}) (*import19.GetVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.GetVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Inserts the defined ISO into a CD-ROM device attached to a Virtual Machine.
func (api *VmApi) InsertCdRomById(vmExtId *string, extId *string, body *import19.CdRomInsertParams, args ...map[string]interface{}) (*import19.InsertCdRomApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InsertCdRomById(context.Background(), &import20.InsertCdRomByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Inserts the defined ISO into a CD-ROM device attached to a Virtual Machine.
func (api *VmServiceApi) InsertCdRomById(ctx context.Context, request *import20.InsertCdRomByIdRequest, args ...map[string]interface{}) (*import19.InsertCdRomApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/insert"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.InsertCdRomApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Inserts the Nutanix Guest Tools installation and configuration ISO into a virtual machine.
func (api *VmApi) InsertVmGuestTools(extId *string, body *import19.GuestToolsInsertConfig, args ...map[string]interface{}) (*import19.InsertVmGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InsertVmGuestTools(context.Background(), &import20.InsertVmGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Inserts the Nutanix Guest Tools installation and configuration ISO into a virtual machine.
func (api *VmServiceApi) InsertVmGuestTools(ctx context.Context, request *import20.InsertVmGuestToolsRequest, args ...map[string]interface{}) (*import19.InsertVmGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools/$actions/insert-iso"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.InsertVmGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Installs Nutanix Guest Tools in a Virtual Machine by using the provided credentials.
func (api *VmApi) InstallVmGuestTools(extId *string, body *import19.GuestToolsInstallConfig, args ...map[string]interface{}) (*import19.InstallVmGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InstallVmGuestTools(context.Background(), &import20.InstallVmGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Installs Nutanix Guest Tools in a Virtual Machine by using the provided credentials.
func (api *VmServiceApi) InstallVmGuestTools(ctx context.Context, request *import20.InstallVmGuestToolsRequest, args ...map[string]interface{}) (*import19.InstallVmGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools/$actions/install"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.InstallVmGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the CD-ROM devices attached to a virtual machine.
func (api *VmApi) ListCdRomsByVmId(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import19.ListCdRomsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCdRomsByVmId(context.Background(), &import20.ListCdRomsByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
	}, args...)
}

// Lists the CD-ROM devices attached to a virtual machine.
func (api *VmServiceApi) ListCdRomsByVmId(ctx context.Context, request *import20.ListCdRomsByVmIdRequest, args ...map[string]interface{}) (*import19.ListCdRomsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/cd-roms"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListCdRomsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the disk devices attached to a virtual machine.
func (api *VmApi) ListDisksByVmId(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import19.ListDisksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDisksByVmId(context.Background(), &import20.ListDisksByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
	}, args...)
}

// Lists the disk devices attached to a virtual machine.
func (api *VmServiceApi) ListDisksByVmId(ctx context.Context, request *import20.ListDisksByVmIdRequest, args ...map[string]interface{}) (*import19.ListDisksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListDisksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the GPU devices attached to a virtual machine
func (api *VmApi) ListGpusByVmId(vmExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import19.ListGpusApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListGpusByVmId(context.Background(), &import20.ListGpusByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
	}, args...)
}

// Lists the GPU devices attached to a virtual machine
func (api *VmServiceApi) ListGpusByVmId(ctx context.Context, request *import20.ListGpusByVmIdRequest, args ...map[string]interface{}) (*import19.ListGpusApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/gpus"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
	}
	if request.Filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*request.Filter_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListGpusApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the network devices attached to a virtual machine.
func (api *VmApi) ListNicsByVmId(vmExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import19.ListNicsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNicsByVmId(context.Background(), &import20.ListNicsByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
	}, args...)
}

// Lists the network devices attached to a virtual machine.
func (api *VmServiceApi) ListNicsByVmId(ctx context.Context, request *import20.ListNicsByVmIdRequest, args ...map[string]interface{}) (*import19.ListNicsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
	}
	if request.Filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*request.Filter_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListNicsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the PCIe devices attached to a virtual machine.
func (api *VmApi) ListPcieDevicesByVmId(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import19.ListPcieDevicesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListPcieDevicesByVmId(context.Background(), &import20.ListPcieDevicesByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
	}, args...)
}

// Lists all the PCIe devices attached to a virtual machine.
func (api *VmServiceApi) ListPcieDevicesByVmId(ctx context.Context, request *import20.ListPcieDevicesByVmIdRequest, args ...map[string]interface{}) (*import19.ListPcieDevicesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/pcie-devices"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListPcieDevicesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the serial ports attached to a virtual machine.
func (api *VmApi) ListSerialPortsByVmId(vmExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import19.ListSerialPortsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSerialPortsByVmId(context.Background(), &import20.ListSerialPortsByVmIdRequest{
		VmExtId: vmExtId,
		Page_:   page_,
		Limit_:  limit_,
	}, args...)
}

// Lists the serial ports attached to a virtual machine.
func (api *VmServiceApi) ListSerialPortsByVmId(ctx context.Context, request *import20.ListSerialPortsByVmIdRequest, args ...map[string]interface{}) (*import19.ListSerialPortsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/serial-ports"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListSerialPortsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the Virtual Machines defined on the system. List of Virtual Machines can be further filtered out using various filtering options.
func (api *VmApi) ListVms(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import19.ListVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVms(context.Background(), &import20.ListVmsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the Virtual Machines defined on the system. List of Virtual Machines can be further filtered out using various filtering options.
func (api *VmServiceApi) ListVms(ctx context.Context, request *import20.ListVmsRequest, args ...map[string]interface{}) (*import19.ListVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
	}
	if request.Filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*request.Filter_, ""))
	}
	if request.Orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*request.Orderby_, ""))
	}
	if request.Select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*request.Select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ListVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Migrates a network device attached to a Virtual Machine to another subnet. This can be used to move network devices between VLAN and VPC subnets.
func (api *VmApi) MigrateNicById(vmExtId *string, extId *string, body *import19.MigrateNicConfig, args ...map[string]interface{}) (*import19.MigrateNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MigrateNicById(context.Background(), &import20.MigrateNicByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Migrates a network device attached to a Virtual Machine to another subnet. This can be used to move network devices between VLAN and VPC subnets.
func (api *VmServiceApi) MigrateNicById(ctx context.Context, request *import20.MigrateNicByIdRequest, args ...map[string]interface{}) (*import19.MigrateNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/migrate"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.MigrateNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Specifies the VmDisks of a VM for migration and the migration plan for them. If all the disks of a VM need to be migrated to the same storage container, only a single migration plan with only the external ID of the destination storage container is needed.  If the disks are being migrated to different containers, one plan per disk needs to be specified.
func (api *VmApi) MigrateVmDisks(extId *string, body *import19.DiskMigrationParams, args ...map[string]interface{}) (*import19.MigrateVmDisksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MigrateVmDisks(context.Background(), &import20.MigrateVmDisksRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Specifies the VmDisks of a VM for migration and the migration plan for them. If all the disks of a VM need to be migrated to the same storage container, only a single migration plan with only the external ID of the destination storage container is needed.  If the disks are being migrated to different containers, one plan per disk needs to be specified.
func (api *VmServiceApi) MigrateVmDisks(ctx context.Context, request *import20.MigrateVmDisksRequest, args ...map[string]interface{}) (*import19.MigrateVmDisksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/migrate-vm-disks"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.MigrateVmDisksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Migrate a VM to another host within the same cluster.
func (api *VmApi) MigrateVmToHost(extId *string, body *import19.VmMigrateToHostParams, args ...map[string]interface{}) (*import19.MigrateVmToHostApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MigrateVmToHost(context.Background(), &import20.MigrateVmToHostRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Migrate a VM to another host within the same cluster.
func (api *VmServiceApi) MigrateVmToHost(ctx context.Context, request *import20.MigrateVmToHostRequest, args ...map[string]interface{}) (*import19.MigrateVmToHostApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/migrate-to-host"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.MigrateVmToHostApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Force a power-cycle for a virtual machine. This operation is equivalent to requesting a VM power-off followed by the VM power-on. Power cycling a VM is slower than resetting it, but it will be creating a fresh instance of the virtual machine. When resetting, the same instance is restarted within the context of the running VM instance.
func (api *VmApi) PowerCycleVm(extId *string, args ...map[string]interface{}) (*import19.PowerCycleVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PowerCycleVm(context.Background(), &import20.PowerCycleVmRequest{
		ExtId: extId,
	}, args...)
}

// Force a power-cycle for a virtual machine. This operation is equivalent to requesting a VM power-off followed by the VM power-on. Power cycling a VM is slower than resetting it, but it will be creating a fresh instance of the virtual machine. When resetting, the same instance is restarted within the context of the running VM instance.
func (api *VmServiceApi) PowerCycleVm(ctx context.Context, request *import20.PowerCycleVmRequest, args ...map[string]interface{}) (*import19.PowerCycleVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/power-cycle"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.PowerCycleVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Forceably shuts down a virtual machine which is equivalent to removing the power cable. Note: The forced shutdown may result in data loss if any operations are in progress during the shutdown.
func (api *VmApi) PowerOffVm(extId *string, args ...map[string]interface{}) (*import19.PowerOffVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PowerOffVm(context.Background(), &import20.PowerOffVmRequest{
		ExtId: extId,
	}, args...)
}

// Forceably shuts down a virtual machine which is equivalent to removing the power cable. Note: The forced shutdown may result in data loss if any operations are in progress during the shutdown.
func (api *VmServiceApi) PowerOffVm(ctx context.Context, request *import20.PowerOffVmRequest, args ...map[string]interface{}) (*import19.PowerOffVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/power-off"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.PowerOffVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Turns a Virtual Machine on.
func (api *VmApi) PowerOnVm(extId *string, args ...map[string]interface{}) (*import19.PowerOnVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PowerOnVm(context.Background(), &import20.PowerOnVmRequest{
		ExtId: extId,
	}, args...)
}

// Turns a Virtual Machine on.
func (api *VmServiceApi) PowerOnVm(ctx context.Context, request *import20.PowerOnVmRequest, args ...map[string]interface{}) (*import19.PowerOnVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/power-on"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.PowerOnVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Collaborative reboot of a Virtual Machine, requesting Nutanix Guest Tools to trigger a reboot from within the VM.
func (api *VmApi) RebootGuestVm(extId *string, body *import19.GuestPowerOptions, args ...map[string]interface{}) (*import19.RebootVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RebootGuestVm(context.Background(), &import20.RebootGuestVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Collaborative reboot of a Virtual Machine, requesting Nutanix Guest Tools to trigger a reboot from within the VM.
func (api *VmServiceApi) RebootGuestVm(ctx context.Context, request *import20.RebootGuestVmRequest, args ...map[string]interface{}) (*import19.RebootVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/guest-reboot"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.RebootVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Collaborative reboot of a Virtual Machine through the ACPI support in the operating system.
func (api *VmApi) RebootVm(extId *string, args ...map[string]interface{}) (*import19.RebootVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RebootVm(context.Background(), &import20.RebootVmRequest{
		ExtId: extId,
	}, args...)
}

// Collaborative reboot of a Virtual Machine through the ACPI support in the operating system.
func (api *VmServiceApi) RebootVm(ctx context.Context, request *import20.RebootVmRequest, args ...map[string]interface{}) (*import19.RebootVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/reboot"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.RebootVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Releases the IP address assignment from a network device attached to a managed network.
func (api *VmApi) ReleaseIpById(vmExtId *string, extId *string, args ...map[string]interface{}) (*import19.ReleaseIpApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReleaseIpById(context.Background(), &import20.ReleaseIpByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
	}, args...)
}

// Releases the IP address assignment from a network device attached to a managed network.
func (api *VmServiceApi) ReleaseIpById(ctx context.Context, request *import20.ReleaseIpByIdRequest, args ...map[string]interface{}) (*import19.ReleaseIpApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/release-ip"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ReleaseIpApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes from the custom attributes of the VM.
func (api *VmApi) RemoveVmCustomAttributes(extId *string, body *import19.UpdateCustomAttributesParams, args ...map[string]interface{}) (*import19.RemoveVmCustomAttributesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RemoveVmCustomAttributes(context.Background(), &import20.RemoveVmCustomAttributesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Removes from the custom attributes of the VM.
func (api *VmServiceApi) RemoveVmCustomAttributes(ctx context.Context, request *import20.RemoveVmCustomAttributesRequest, args ...map[string]interface{}) (*import19.RemoveVmCustomAttributesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/remove-custom-attributes"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.RemoveVmCustomAttributesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes from the custom attributes of the VM Disk entity.
func (api *VmApi) RemoveVmDiskCustomAttributes(vmExtId *string, extId *string, body *import19.UpdateCustomAttributesParams, args ...map[string]interface{}) (*import19.RemoveVmDiskCustomAttributesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RemoveVmDiskCustomAttributes(context.Background(), &import20.RemoveVmDiskCustomAttributesRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Removes from the custom attributes of the VM Disk entity.
func (api *VmServiceApi) RemoveVmDiskCustomAttributes(ctx context.Context, request *import20.RemoveVmDiskCustomAttributesRequest, args ...map[string]interface{}) (*import19.RemoveVmDiskCustomAttributesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}/$actions/remove-custom-attributes"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.RemoveVmDiskCustomAttributesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Force reset of a Virtual Machine, without waiting for the guest VM to shutdown itself. Resetting a VM is faster than power-cycle as the reset occurs within the context of the running virtual machine instance rather than creating a fresh instance.
func (api *VmApi) ResetVm(extId *string, args ...map[string]interface{}) (*import19.ResetVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ResetVm(context.Background(), &import20.ResetVmRequest{
		ExtId: extId,
	}, args...)
}

// Force reset of a Virtual Machine, without waiting for the guest VM to shutdown itself. Resetting a VM is faster than power-cycle as the reset occurs within the context of the running virtual machine instance rather than creating a fresh instance.
func (api *VmServiceApi) ResetVm(ctx context.Context, request *import20.ResetVmRequest, args ...map[string]interface{}) (*import19.ResetVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/reset"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ResetVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point. The project, categories and VM owner reference will not be reverted as part of this operation by default.
func (api *VmApi) RevertVm(extId *string, body *import19.RevertParams, args ...map[string]interface{}) (*import19.RevertVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RevertVm(context.Background(), &import20.RevertVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point. The project, categories and VM owner reference will not be reverted as part of this operation by default.
func (api *VmServiceApi) RevertVm(ctx context.Context, request *import20.RevertVmRequest, args ...map[string]interface{}) (*import19.RevertVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/revert"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.RevertVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Collaborative shutdown of a Virtual Machine, requesting Nutanix Guest Tools to trigger a shutdown from within the VM.
func (api *VmApi) ShutdownGuestVm(extId *string, body *import19.GuestPowerOptions, args ...map[string]interface{}) (*import19.ShutdownVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShutdownGuestVm(context.Background(), &import20.ShutdownGuestVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Collaborative shutdown of a Virtual Machine, requesting Nutanix Guest Tools to trigger a shutdown from within the VM.
func (api *VmServiceApi) ShutdownGuestVm(ctx context.Context, request *import20.ShutdownGuestVmRequest, args ...map[string]interface{}) (*import19.ShutdownVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/guest-shutdown"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.ShutdownVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Collaborative shutdown of a Virtual Machine through the ACPI support in the operating system.
func (api *VmApi) ShutdownVm(extId *string, args ...map[string]interface{}) (*import19.ShutdownVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShutdownVm(context.Background(), &import20.ShutdownVmRequest{
		ExtId: extId,
	}, args...)
}

// Collaborative shutdown of a Virtual Machine through the ACPI support in the operating system.
func (api *VmServiceApi) ShutdownVm(ctx context.Context, request *import20.ShutdownVmRequest, args ...map[string]interface{}) (*import19.ShutdownVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/$actions/shutdown"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.ShutdownVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Trigger an in-guest uninstallation of Nutanix Guest Tools.
func (api *VmApi) UninstallVmGuestTools(extId *string, args ...map[string]interface{}) (*import19.UninstallVmGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UninstallVmGuestTools(context.Background(), &import20.UninstallVmGuestToolsRequest{
		ExtId: extId,
	}, args...)
}

// Trigger an in-guest uninstallation of Nutanix Guest Tools.
func (api *VmServiceApi) UninstallVmGuestTools(ctx context.Context, request *import20.UninstallVmGuestToolsRequest, args ...map[string]interface{}) (*import19.UninstallVmGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools/$actions/uninstall"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UninstallVmGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration details for the provided disk device.
func (api *VmApi) UpdateDiskById(vmExtId *string, extId *string, body *import19.Disk, args ...map[string]interface{}) (*import19.UpdateDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDiskById(context.Background(), &import20.UpdateDiskByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Updates the configuration details for the provided disk device.
func (api *VmServiceApi) UpdateDiskById(ctx context.Context, request *import20.UpdateDiskByIdRequest, args ...map[string]interface{}) (*import19.UpdateDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/disks/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UpdateDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Nutanix Guest Tools configuration for a virtual machine.
func (api *VmApi) UpdateGuestToolsById(extId *string, body *import19.GuestTools, args ...map[string]interface{}) (*import19.UpdateGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateGuestToolsById(context.Background(), &import20.UpdateGuestToolsByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Nutanix Guest Tools configuration for a virtual machine.
func (api *VmServiceApi) UpdateGuestToolsById(ctx context.Context, request *import20.UpdateGuestToolsByIdRequest, args ...map[string]interface{}) (*import19.UpdateGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UpdateGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration details for the provided network device.
func (api *VmApi) UpdateNicById(vmExtId *string, extId *string, body *import19.Nic, args ...map[string]interface{}) (*import19.UpdateNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateNicById(context.Background(), &import20.UpdateNicByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Updates the configuration details for the provided network device.
func (api *VmServiceApi) UpdateNicById(ctx context.Context, request *import20.UpdateNicByIdRequest, args ...map[string]interface{}) (*import19.UpdateNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/nics/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UpdateNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration details for the provided serial port.
func (api *VmApi) UpdateSerialPortById(vmExtId *string, extId *string, body *import19.SerialPort, args ...map[string]interface{}) (*import19.UpdateSerialPortApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSerialPortById(context.Background(), &import20.UpdateSerialPortByIdRequest{
		VmExtId: vmExtId,
		ExtId:   extId,
		Body:    body,
	}, args...)
}

// Updates the configuration details for the provided serial port.
func (api *VmServiceApi) UpdateSerialPortById(ctx context.Context, request *import20.UpdateSerialPortByIdRequest, args ...map[string]interface{}) (*import19.UpdateSerialPortApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{vmExtId}/serial-ports/{extId}"

	// verify the required parameter 'vmExtId' is set
	if nil == request.VmExtId {
		return nil, client.ReportError("vmExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UpdateSerialPortApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates configuration details for a virtual machine.
func (api *VmApi) UpdateVmById(extId *string, body *import19.Vm, args ...map[string]interface{}) (*import19.UpdateVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVmById(context.Background(), &import20.UpdateVmByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates configuration details for a virtual machine.
func (api *VmServiceApi) UpdateVmById(ctx context.Context, request *import20.UpdateVmByIdRequest, args ...map[string]interface{}) (*import19.UpdateVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import19.UpdateVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Trigger an in-guest upgrade of Nutanix Guest Tools.
func (api *VmApi) UpgradeVmGuestTools(extId *string, body *import19.GuestToolsUpgradeConfig, args ...map[string]interface{}) (*import19.UpgradeVmGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpgradeVmGuestTools(context.Background(), &import20.UpgradeVmGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Trigger an in-guest upgrade of Nutanix Guest Tools.
func (api *VmServiceApi) UpgradeVmGuestTools(ctx context.Context, request *import20.UpgradeVmGuestToolsRequest, args ...map[string]interface{}) (*import19.UpgradeVmGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/ahv/config/vms/{extId}/guest-tools/$actions/upgrade"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import19.UpgradeVmGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
