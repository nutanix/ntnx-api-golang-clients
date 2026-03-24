package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/common/v1/stats"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
	import6 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/request/volumegroups"
	import5 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type VolumeGroupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VolumeGroupsServiceApi
}

type VolumeGroupsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVolumeGroupsApi(apiClient *client.ApiClient) *VolumeGroupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VolumeGroupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVolumeGroupsServiceApi(a.ApiClient)

	return a
}

func NewVolumeGroupsServiceApi(apiClient *client.ApiClient) *VolumeGroupsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VolumeGroupsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Associates a category to a Volume Group identified by {extId}.
func (api *VolumeGroupsApi) AssociateCategory(extId *string, body *import1.CategoryEntityReferences, args ...map[string]interface{}) (*import1.AssociateCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociateCategory(context.Background(), &import6.AssociateCategoryRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Associates a category to a Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) AssociateCategory(ctx context.Context, request *import6.AssociateCategoryRequest, args ...map[string]interface{}) (*import1.AssociateCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/associate-category"

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

	unmarshalledResp := new(import1.AssociateCategoryApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Attaches iSCSI initiator to a Volume Group identified by {extId}.
func (api *VolumeGroupsApi) AttachIscsiClient(extId *string, body *import1.IscsiClient, args ...map[string]interface{}) (*import1.AttachIscsiClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AttachIscsiClient(context.Background(), &import6.AttachIscsiClientRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Attaches iSCSI initiator to a Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) AttachIscsiClient(ctx context.Context, request *import6.AttachIscsiClientRequest, args ...map[string]interface{}) (*import1.AttachIscsiClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/attach-iscsi-client"

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

	unmarshalledResp := new(import1.AttachIscsiClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Attaches a NVMe-TCP client/initiator to a Volume Group identified by its external identifier.
func (api *VolumeGroupsApi) AttachNvmfClient(extId *string, body *import1.NvmfClient, args ...map[string]interface{}) (*import1.AttachNvmfClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AttachNvmfClient(context.Background(), &import6.AttachNvmfClientRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Attaches a NVMe-TCP client/initiator to a Volume Group identified by its external identifier.
func (api *VolumeGroupsServiceApi) AttachNvmfClient(ctx context.Context, request *import6.AttachNvmfClientRequest, args ...map[string]interface{}) (*import1.AttachNvmfClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/attach-nvmf-client"

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

	unmarshalledResp := new(import1.AttachNvmfClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Attaches VM to a Volume Group identified by {extId}.
func (api *VolumeGroupsApi) AttachVm(extId *string, body *import1.VmAttachment, args ...map[string]interface{}) (*import1.AttachVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AttachVm(context.Background(), &import6.AttachVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Attaches VM to a Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) AttachVm(ctx context.Context, request *import6.AttachVmRequest, args ...map[string]interface{}) (*import1.AttachVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/attach-vm"

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

	unmarshalledResp := new(import1.AttachVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a new Volume Disk.
func (api *VolumeGroupsApi) CreateVolumeDisk(volumeGroupExtId *string, body *import1.VolumeDisk, args ...map[string]interface{}) (*import1.CreateVolumeDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVolumeDisk(context.Background(), &import6.CreateVolumeDiskRequest{
		VolumeGroupExtId: volumeGroupExtId,
		Body:             body,
	}, args...)
}

// Creates a new Volume Disk.
func (api *VolumeGroupsServiceApi) CreateVolumeDisk(ctx context.Context, request *import6.CreateVolumeDiskRequest, args ...map[string]interface{}) (*import1.CreateVolumeDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/disks"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.CreateVolumeDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a new Volume Group.
func (api *VolumeGroupsApi) CreateVolumeGroup(body *import1.VolumeGroup, args ...map[string]interface{}) (*import1.CreateVolumeGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVolumeGroup(context.Background(), &import6.CreateVolumeGroupRequest{
		Body: body,
	}, args...)
}

// Creates a new Volume Group.
func (api *VolumeGroupsServiceApi) CreateVolumeGroup(ctx context.Context, request *import6.CreateVolumeGroupRequest, args ...map[string]interface{}) (*import1.CreateVolumeGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups"

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

	unmarshalledResp := new(import1.CreateVolumeGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a Volume Disk identified by {extId} in the Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsApi) DeleteVolumeDiskById(volumeGroupExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteVolumeDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVolumeDiskById(context.Background(), &import6.DeleteVolumeDiskByIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		ExtId:            extId,
	}, args...)
}

// Deletes a Volume Disk identified by {extId} in the Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsServiceApi) DeleteVolumeDiskById(ctx context.Context, request *import6.DeleteVolumeDiskByIdRequest, args ...map[string]interface{}) (*import1.DeleteVolumeDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/disks/{extId}"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteVolumeDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the Volume Group identified by {extId}.
func (api *VolumeGroupsApi) DeleteVolumeGroupById(extId *string, args ...map[string]interface{}) (*import1.DeleteVolumeGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVolumeGroupById(context.Background(), &import6.DeleteVolumeGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) DeleteVolumeGroupById(ctx context.Context, request *import6.DeleteVolumeGroupByIdRequest, args ...map[string]interface{}) (*import1.DeleteVolumeGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}"

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

	unmarshalledResp := new(import1.DeleteVolumeGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Detaches iSCSI initiator identified by {extId} from a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsApi) DetachIscsiClient(extId *string, body *import1.IscsiClientAttachment, args ...map[string]interface{}) (*import1.DetachIscsiClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DetachIscsiClient(context.Background(), &import6.DetachIscsiClientRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Detaches iSCSI initiator identified by {extId} from a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsServiceApi) DetachIscsiClient(ctx context.Context, request *import6.DetachIscsiClientRequest, args ...map[string]interface{}) (*import1.DetachIscsiClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/detach-iscsi-client"

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

	unmarshalledResp := new(import1.DetachIscsiClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Detaches a NVMe-TCP client/initiator from a Volume Group identified by its external identifier.
func (api *VolumeGroupsApi) DetachNvmfClient(extId *string, body *import1.NvmfClient, args ...map[string]interface{}) (*import1.DetachNvmfClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DetachNvmfClient(context.Background(), &import6.DetachNvmfClientRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Detaches a NVMe-TCP client/initiator from a Volume Group identified by its external identifier.
func (api *VolumeGroupsServiceApi) DetachNvmfClient(ctx context.Context, request *import6.DetachNvmfClientRequest, args ...map[string]interface{}) (*import1.DetachNvmfClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/detach-nvmf-client"

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

	unmarshalledResp := new(import1.DetachNvmfClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Detaches VM identified by {extId} from a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsApi) DetachVm(extId *string, body *import1.VmAttachment, args ...map[string]interface{}) (*import1.DetachVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DetachVm(context.Background(), &import6.DetachVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Detaches VM identified by {extId} from a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsServiceApi) DetachVm(ctx context.Context, request *import6.DetachVmRequest, args ...map[string]interface{}) (*import1.DetachVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/detach-vm"

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

	unmarshalledResp := new(import1.DetachVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disassociates a category from a Volume Group identified by {extId}.
func (api *VolumeGroupsApi) DisassociateCategory(extId *string, body *import1.CategoryEntityReferences, args ...map[string]interface{}) (*import1.DisassociateCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisassociateCategory(context.Background(), &import6.DisassociateCategoryRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Disassociates a category from a Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) DisassociateCategory(ctx context.Context, request *import6.DisassociateCategoryRequest, args ...map[string]interface{}) (*import1.DisassociateCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/disassociate-category"

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

	unmarshalledResp := new(import1.DisassociateCategoryApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the Volume Disk identified by {extId} in the Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsApi) GetVolumeDiskById(volumeGroupExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetVolumeDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVolumeDiskById(context.Background(), &import6.GetVolumeDiskByIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		ExtId:            extId,
	}, args...)
}

// Query the Volume Disk identified by {extId} in the Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsServiceApi) GetVolumeDiskById(ctx context.Context, request *import6.GetVolumeDiskByIdRequest, args ...map[string]interface{}) (*import1.GetVolumeDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/disks/{extId}"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetVolumeDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the Volume Disk stats identified by {diskExtId}.
func (api *VolumeGroupsApi) GetVolumeDiskStats(volumeGroupExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import4.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import5.GetVolumeDiskStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVolumeDiskStats(context.Background(), &import6.GetVolumeDiskStatsRequest{
		VolumeGroupExtId:  volumeGroupExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Query the Volume Disk stats identified by {diskExtId}.
func (api *VolumeGroupsServiceApi) GetVolumeDiskStats(ctx context.Context, request *import6.GetVolumeDiskStatsRequest, args ...map[string]interface{}) (*import5.GetVolumeDiskStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/stats/volume-groups/{volumeGroupExtId}/disks/{extId}"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.GetVolumeDiskStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the Volume Group identified by {extId}.
func (api *VolumeGroupsApi) GetVolumeGroupById(extId *string, expand_ *string, args ...map[string]interface{}) (*import1.GetVolumeGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVolumeGroupById(context.Background(), &import6.GetVolumeGroupByIdRequest{
		ExtId:   extId,
		Expand_: expand_,
	}, args...)
}

// Query the Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) GetVolumeGroupById(ctx context.Context, request *import6.GetVolumeGroupByIdRequest, args ...map[string]interface{}) (*import1.GetVolumeGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}"

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

	// Query Params
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import1.GetVolumeGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query for metadata information which is associated with the Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsApi) GetVolumeGroupMetadataById(volumeGroupExtId *string, args ...map[string]interface{}) (*import1.GetVolumeGroupMetadataApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVolumeGroupMetadataById(context.Background(), &import6.GetVolumeGroupMetadataByIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
	}, args...)
}

// Query for metadata information which is associated with the Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsServiceApi) GetVolumeGroupMetadataById(ctx context.Context, request *import6.GetVolumeGroupMetadataByIdRequest, args ...map[string]interface{}) (*import1.GetVolumeGroupMetadataApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/metadata"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetVolumeGroupMetadataApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the Volume Group stats identified by {extId}.
func (api *VolumeGroupsApi) GetVolumeGroupStats(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import4.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import5.GetVolumeGroupStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVolumeGroupStats(context.Background(), &import6.GetVolumeGroupStatsRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Query the Volume Group stats identified by {extId}.
func (api *VolumeGroupsServiceApi) GetVolumeGroupStats(ctx context.Context, request *import6.GetVolumeGroupStatsRequest, args ...map[string]interface{}) (*import5.GetVolumeGroupStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/stats/volume-groups/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
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

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.GetVolumeGroupStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the category details that are associated with the Volume Group identified by {volumeGroupExtId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsApi) ListCategoryAssociationsByVolumeGroupId(volumeGroupExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.ListCategoryAssociationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCategoryAssociationsByVolumeGroupId(context.Background(), &import6.ListCategoryAssociationsByVolumeGroupIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		Page_:            page_,
		Limit_:           limit_,
	}, args...)
}

// Query the category details that are associated with the Volume Group identified by {volumeGroupExtId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsServiceApi) ListCategoryAssociationsByVolumeGroupId(ctx context.Context, request *import6.ListCategoryAssociationsByVolumeGroupIdRequest, args ...map[string]interface{}) (*import1.ListCategoryAssociationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/category-associations"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListCategoryAssociationsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the list of external iSCSI attachments for a Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsApi) ListExternalIscsiAttachmentsByVolumeGroupId(volumeGroupExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListExternalIscsiAttachmentsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListExternalIscsiAttachmentsByVolumeGroupId(context.Background(), &import6.ListExternalIscsiAttachmentsByVolumeGroupIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		Page_:            page_,
		Limit_:           limit_,
		Filter_:          filter_,
		Orderby_:         orderby_,
		Expand_:          expand_,
		Select_:          select_,
	}, args...)
}

// Query the list of external iSCSI attachments for a Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsServiceApi) ListExternalIscsiAttachmentsByVolumeGroupId(ctx context.Context, request *import6.ListExternalIscsiAttachmentsByVolumeGroupIdRequest, args ...map[string]interface{}) (*import1.ListExternalIscsiAttachmentsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/external-iscsi-attachments"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import1.ListExternalIscsiAttachmentsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the list of VM attachments for a Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsApi) ListVmAttachmentsByVolumeGroupId(volumeGroupExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.ListVmAttachmentsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmAttachmentsByVolumeGroupId(context.Background(), &import6.ListVmAttachmentsByVolumeGroupIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		Page_:            page_,
		Limit_:           limit_,
		Filter_:          filter_,
		Orderby_:         orderby_,
	}, args...)
}

// Query the list of VM attachments for a Volume Group identified by {extId}.
//
// Deprecated: This API has been deprecated.
func (api *VolumeGroupsServiceApi) ListVmAttachmentsByVolumeGroupId(ctx context.Context, request *import6.ListVmAttachmentsByVolumeGroupIdRequest, args ...map[string]interface{}) (*import1.ListVmAttachmentsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/vm-attachments"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListVmAttachmentsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the list of disks corresponding to a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsApi) ListVolumeDisksByVolumeGroupId(volumeGroupExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListVolumeDisksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVolumeDisksByVolumeGroupId(context.Background(), &import6.ListVolumeDisksByVolumeGroupIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		Page_:            page_,
		Limit_:           limit_,
		Filter_:          filter_,
		Orderby_:         orderby_,
		Select_:          select_,
	}, args...)
}

// Query the list of disks corresponding to a Volume Group identified by {volumeGroupExtId}.
func (api *VolumeGroupsServiceApi) ListVolumeDisksByVolumeGroupId(ctx context.Context, request *import6.ListVolumeDisksByVolumeGroupIdRequest, args ...map[string]interface{}) (*import1.ListVolumeDisksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/disks"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListVolumeDisksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the list of Volume Groups.
func (api *VolumeGroupsApi) ListVolumeGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListVolumeGroupsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVolumeGroups(context.Background(), &import6.ListVolumeGroupsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Query the list of Volume Groups.
func (api *VolumeGroupsServiceApi) ListVolumeGroups(ctx context.Context, request *import6.ListVolumeGroupsRequest, args ...map[string]interface{}) (*import1.ListVolumeGroupsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups"

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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import1.ListVolumeGroupsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Reverts a Volume Group identified by Volume Group external identifier. This API performs an in-place restore from a specified Volume Group recovery point.
func (api *VolumeGroupsApi) RevertVolumeGroup(extId *string, body *import1.RevertSpec, args ...map[string]interface{}) (*import1.RevertVolumeGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RevertVolumeGroup(context.Background(), &import6.RevertVolumeGroupRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Reverts a Volume Group identified by Volume Group external identifier. This API performs an in-place restore from a specified Volume Group recovery point.
func (api *VolumeGroupsServiceApi) RevertVolumeGroup(ctx context.Context, request *import6.RevertVolumeGroupRequest, args ...map[string]interface{}) (*import1.RevertVolumeGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}/$actions/revert"

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

	unmarshalledResp := new(import1.RevertVolumeGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a specific Volume Disk identified by {extId}.
func (api *VolumeGroupsApi) UpdateVolumeDiskById(volumeGroupExtId *string, extId *string, body *import1.VolumeDisk, args ...map[string]interface{}) (*import1.UpdateVolumeDiskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVolumeDiskById(context.Background(), &import6.UpdateVolumeDiskByIdRequest{
		VolumeGroupExtId: volumeGroupExtId,
		ExtId:            extId,
		Body:             body,
	}, args...)
}

// Updates a specific Volume Disk identified by {extId}.
func (api *VolumeGroupsServiceApi) UpdateVolumeDiskById(ctx context.Context, request *import6.UpdateVolumeDiskByIdRequest, args ...map[string]interface{}) (*import1.UpdateVolumeDiskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{volumeGroupExtId}/disks/{extId}"

	// verify the required parameter 'volumeGroupExtId' is set
	if nil == request.VolumeGroupExtId {
		return nil, client.ReportError("volumeGroupExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupExtId, "")), -1)
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

	unmarshalledResp := new(import1.UpdateVolumeDiskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates details of a specific Volume Group identified by {extId}.
func (api *VolumeGroupsApi) UpdateVolumeGroupById(extId *string, body *import1.VolumeGroup, args ...map[string]interface{}) (*import1.UpdateVolumeGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVolumeGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateVolumeGroupById(context.Background(), &import6.UpdateVolumeGroupByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates details of a specific Volume Group identified by {extId}.
func (api *VolumeGroupsServiceApi) UpdateVolumeGroupById(ctx context.Context, request *import6.UpdateVolumeGroupByIdRequest, args ...map[string]interface{}) (*import1.UpdateVolumeGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/volume-groups/{extId}"

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

	unmarshalledResp := new(import1.UpdateVolumeGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
