package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	import27 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/vmrecoverypoints"
	"net/http"
	"net/url"
	"strings"
)

type VmRecoveryPointsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VmRecoveryPointsServiceApi
}

type VmRecoveryPointsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmRecoveryPointsApi(apiClient *client.ApiClient) *VmRecoveryPointsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmRecoveryPointsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVmRecoveryPointsServiceApi(a.ApiClient)

	return a
}

func NewVmRecoveryPointsServiceApi(apiClient *client.ApiClient) *VmRecoveryPointsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmRecoveryPointsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an AHV VM recovery point based on the provided recovery point details.
func (api *VmRecoveryPointsApi) CreateVmRecoveryPoint(body *import21.VmRecoveryPoint, args ...map[string]interface{}) (*import21.CreateVMRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateVmRecoveryPoint(context.Background(), &import27.CreateVmRecoveryPointRequest{
		Body: body,
	}, args...)
}

// Creates an AHV VM recovery point based on the provided recovery point details.
func (api *VmRecoveryPointsServiceApi) CreateVmRecoveryPoint(ctx context.Context, request *import27.CreateVmRecoveryPointRequest, args ...map[string]interface{}) (*import21.CreateVMRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-recovery-points"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.CreateVMRecoveryPointApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes an AHV VM recovery point based on the provided identifier.
func (api *VmRecoveryPointsApi) DeleteVmRecoveryPointByExtId(extId *string, args ...map[string]interface{}) (*import21.DeleteVmRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteVmRecoveryPointByExtId(context.Background(), &import27.DeleteVmRecoveryPointByExtIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes an AHV VM recovery point based on the provided identifier.
func (api *VmRecoveryPointsServiceApi) DeleteVmRecoveryPointByExtId(ctx context.Context, request *import27.DeleteVmRecoveryPointByExtIdRequest, args ...map[string]interface{}) (*import21.DeleteVmRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-recovery-points/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.DeleteVmRecoveryPointApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves the AHV VM recovery point details, including both VM recovery point and VM configuration.
func (api *VmRecoveryPointsApi) GetVmRecoveryPointByExtId(extId *string, args ...map[string]interface{}) (*import21.GetVmRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmRecoveryPointByExtId(context.Background(), &import27.GetVmRecoveryPointByExtIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the AHV VM recovery point details, including both VM recovery point and VM configuration.
func (api *VmRecoveryPointsServiceApi) GetVmRecoveryPointByExtId(ctx context.Context, request *import27.GetVmRecoveryPointByExtIdRequest, args ...map[string]interface{}) (*import21.GetVmRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-recovery-points/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.GetVmRecoveryPointApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List the AHV VM recovery points defined on the PC. It can be further refined using various filtering.
func (api *VmRecoveryPointsApi) ListVmRecoveryPoints(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import21.ListVmRecoveryPointsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVmRecoveryPoints(context.Background(), &import27.ListVmRecoveryPointsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List the AHV VM recovery points defined on the PC. It can be further refined using various filtering.
func (api *VmRecoveryPointsServiceApi) ListVmRecoveryPoints(ctx context.Context, request *import27.ListVmRecoveryPointsRequest, args ...map[string]interface{}) (*import21.ListVmRecoveryPointsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-recovery-points"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.ListVmRecoveryPointsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Restores a new VM from an AHV VM recovery point using the VM configuration captured in the VM recovery point and the parameters provided in the restore API request body.
func (api *VmRecoveryPointsApi) RestoreVmRecoveryPoint(extId *string, body *import21.RestoreVmRecoveryPointParams, args ...map[string]interface{}) (*import21.RestoreVmRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVmRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RestoreVmRecoveryPoint(context.Background(), &import27.RestoreVmRecoveryPointRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Restores a new VM from an AHV VM recovery point using the VM configuration captured in the VM recovery point and the parameters provided in the restore API request body.
func (api *VmRecoveryPointsServiceApi) RestoreVmRecoveryPoint(ctx context.Context, request *import27.RestoreVmRecoveryPointRequest, args ...map[string]interface{}) (*import21.RestoreVmRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/config/vm-recovery-points/{extId}/$actions/restore"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import21.RestoreVmRecoveryPointApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
