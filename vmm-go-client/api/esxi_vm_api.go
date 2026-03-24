package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/esxi/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/esxivm"
	"net/http"
	"net/url"
	"strings"
)

type EsxiVmApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EsxiVmServiceApi
}

type EsxiVmServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEsxiVmApi(apiClient *client.ApiClient) *EsxiVmApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EsxiVmApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEsxiVmServiceApi(a.ApiClient)

	return a
}

func NewEsxiVmServiceApi(apiClient *client.ApiClient) *EsxiVmServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EsxiVmServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Assign the owner of a virtual machine.
func (api *EsxiVmApi) AssignVmOwner(extId *string, body *import4.OwnershipInfo, args ...map[string]interface{}) (*import4.AssignVmOwnerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssignVmOwner(context.Background(), &import5.AssignVmOwnerRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Assign the owner of a virtual machine.
func (api *EsxiVmServiceApi) AssignVmOwner(ctx context.Context, request *import5.AssignVmOwnerRequest, args ...map[string]interface{}) (*import4.AssignVmOwnerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/assign-owner"

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

	unmarshalledResp := new(import4.AssignVmOwnerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Associate categories to a virtual machine.
func (api *EsxiVmApi) AssociateCategories(extId *string, body *import4.AssociateVmCategoriesParams, args ...map[string]interface{}) (*import4.AssociateCategoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociateCategories(context.Background(), &import5.AssociateCategoriesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Associate categories to a virtual machine.
func (api *EsxiVmServiceApi) AssociateCategories(ctx context.Context, request *import5.AssociateCategoriesRequest, args ...map[string]interface{}) (*import4.AssociateCategoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/associate-categories"

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

	unmarshalledResp := new(import4.AssociateCategoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disassociate categories from a virtual machine.
func (api *EsxiVmApi) DisassociateCategories(extId *string, body *import4.DisassociateVmCategoriesParams, args ...map[string]interface{}) (*import4.DisassociateCategoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisassociateCategories(context.Background(), &import5.DisassociateCategoriesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Disassociate categories from a virtual machine.
func (api *EsxiVmServiceApi) DisassociateCategories(ctx context.Context, request *import5.DisassociateCategoriesRequest, args ...map[string]interface{}) (*import4.DisassociateCategoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/disassociate-categories"

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

	unmarshalledResp := new(import4.DisassociateCategoriesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the Nutanix Guest Tools configuration for a virtual machine.
func (api *EsxiVmApi) GetNutanixGuestToolsById(extId *string, args ...map[string]interface{}) (*import4.GetNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNutanixGuestToolsById(context.Background(), &import5.GetNutanixGuestToolsByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the Nutanix Guest Tools configuration for a virtual machine.
func (api *EsxiVmServiceApi) GetNutanixGuestToolsById(ctx context.Context, request *import5.GetNutanixGuestToolsByIdRequest, args ...map[string]interface{}) (*import4.GetNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools"

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

	unmarshalledResp := new(import4.GetNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves configuration details for a virtual machine.
func (api *EsxiVmApi) GetVmById(extId *string, args ...map[string]interface{}) (*import4.GetVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmById(context.Background(), &import5.GetVmByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves configuration details for a virtual machine.
func (api *EsxiVmServiceApi) GetVmById(ctx context.Context, request *import5.GetVmByIdRequest, args ...map[string]interface{}) (*import4.GetVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}"

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

	unmarshalledResp := new(import4.GetVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Inserts the Nutanix Guest Tools installation and configuration ISO into a virtual machine.
func (api *EsxiVmApi) InsertNutanixGuestTools(extId *string, body *import4.NutanixGuestToolsInsertConfig, args ...map[string]interface{}) (*import4.InsertNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InsertNutanixGuestTools(context.Background(), &import5.InsertNutanixGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Inserts the Nutanix Guest Tools installation and configuration ISO into a virtual machine.
func (api *EsxiVmServiceApi) InsertNutanixGuestTools(ctx context.Context, request *import5.InsertNutanixGuestToolsRequest, args ...map[string]interface{}) (*import4.InsertNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/insert-iso"

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

	unmarshalledResp := new(import4.InsertNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Installs Nutanix Guest Tools in a Virtual Machine by using the provided credentials.
func (api *EsxiVmApi) InstallNutanixGuestTools(extId *string, body *import4.NutanixGuestToolsInstallConfig, args ...map[string]interface{}) (*import4.InstallNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InstallNutanixGuestTools(context.Background(), &import5.InstallNutanixGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Installs Nutanix Guest Tools in a Virtual Machine by using the provided credentials.
func (api *EsxiVmServiceApi) InstallNutanixGuestTools(ctx context.Context, request *import5.InstallNutanixGuestToolsRequest, args ...map[string]interface{}) (*import4.InstallNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/install"

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

	unmarshalledResp := new(import4.InstallNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the Virtual Machines defined on the system. List of Virtual Machines can be further filtered out using various filtering options.
func (api *EsxiVmApi) ListVms(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListVmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVms(context.Background(), &import5.ListVmsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the Virtual Machines defined on the system. List of Virtual Machines can be further filtered out using various filtering options.
func (api *EsxiVmServiceApi) ListVms(ctx context.Context, request *import5.ListVmsRequest, args ...map[string]interface{}) (*import4.ListVmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms"

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

	unmarshalledResp := new(import4.ListVmsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Forceably shuts down a virtual machine which is equivalent to removing the power cable. Note: The forced shutdown may result in data loss if any operations are in progress during the shutdown.
func (api *EsxiVmApi) PowerOffVm(extId *string, args ...map[string]interface{}) (*import4.PowerOffVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PowerOffVm(context.Background(), &import5.PowerOffVmRequest{
		ExtId: extId,
	}, args...)
}

// Forceably shuts down a virtual machine which is equivalent to removing the power cable. Note: The forced shutdown may result in data loss if any operations are in progress during the shutdown.
func (api *EsxiVmServiceApi) PowerOffVm(ctx context.Context, request *import5.PowerOffVmRequest, args ...map[string]interface{}) (*import4.PowerOffVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/power-off"

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

	unmarshalledResp := new(import4.PowerOffVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Powers a Virtual Machine on or resumes it from the suspended state.
func (api *EsxiVmApi) PowerOnVm(extId *string, args ...map[string]interface{}) (*import4.PowerOnVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PowerOnVm(context.Background(), &import5.PowerOnVmRequest{
		ExtId: extId,
	}, args...)
}

// Powers a Virtual Machine on or resumes it from the suspended state.
func (api *EsxiVmServiceApi) PowerOnVm(ctx context.Context, request *import5.PowerOnVmRequest, args ...map[string]interface{}) (*import4.PowerOnVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/power-on"

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

	unmarshalledResp := new(import4.PowerOnVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Issues a command to reboot ESXi guest OS. This operation requires ESXi tools installed.
func (api *EsxiVmApi) RebootGuestVm(extId *string, args ...map[string]interface{}) (*import4.RebootGuestOSApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RebootGuestVm(context.Background(), &import5.RebootGuestVmRequest{
		ExtId: extId,
	}, args...)
}

// Issues a command to reboot ESXi guest OS. This operation requires ESXi tools installed.
func (api *EsxiVmServiceApi) RebootGuestVm(ctx context.Context, request *import5.RebootGuestVmRequest, args ...map[string]interface{}) (*import4.RebootGuestOSApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/guest-reboot"

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

	unmarshalledResp := new(import4.RebootGuestOSApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Sequentially performs the power off and power on operations; any operation between these actions will fail.
func (api *EsxiVmApi) ResetVm(extId *string, args ...map[string]interface{}) (*import4.ResetVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ResetVm(context.Background(), &import5.ResetVmRequest{
		ExtId: extId,
	}, args...)
}

// Sequentially performs the power off and power on operations; any operation between these actions will fail.
func (api *EsxiVmServiceApi) ResetVm(ctx context.Context, request *import5.ResetVmRequest, args ...map[string]interface{}) (*import4.ResetVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/reset"

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

	unmarshalledResp := new(import4.ResetVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point.
func (api *EsxiVmApi) RevertVm(extId *string, body *import4.RevertParams, args ...map[string]interface{}) (*import4.RevertVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RevertVm(context.Background(), &import5.RevertVmRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point.
func (api *EsxiVmServiceApi) RevertVm(ctx context.Context, request *import5.RevertVmRequest, args ...map[string]interface{}) (*import4.RevertVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/revert"

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

	unmarshalledResp := new(import4.RevertVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Issues a command to the ESXi guest OS to perform a clean shut down of services running on it. This operation requires ESXi tools to be installed.
func (api *EsxiVmApi) ShutdownGuestVm(extId *string, args ...map[string]interface{}) (*import4.ShutdownVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShutdownGuestVm(context.Background(), &import5.ShutdownGuestVmRequest{
		ExtId: extId,
	}, args...)
}

// Issues a command to the ESXi guest OS to perform a clean shut down of services running on it. This operation requires ESXi tools to be installed.
func (api *EsxiVmServiceApi) ShutdownGuestVm(ctx context.Context, request *import5.ShutdownGuestVmRequest, args ...map[string]interface{}) (*import4.ShutdownVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/guest-shutdown"

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

	unmarshalledResp := new(import4.ShutdownVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Pause/Suspend execution in an ESXi virtual machine.
func (api *EsxiVmApi) SuspendVm(extId *string, args ...map[string]interface{}) (*import4.SuspendVmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SuspendVm(context.Background(), &import5.SuspendVmRequest{
		ExtId: extId,
	}, args...)
}

// Pause/Suspend execution in an ESXi virtual machine.
func (api *EsxiVmServiceApi) SuspendVm(ctx context.Context, request *import5.SuspendVmRequest, args ...map[string]interface{}) (*import4.SuspendVmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/$actions/suspend"

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

	unmarshalledResp := new(import4.SuspendVmApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Trigger an in-guest uninstallation of Nutanix Guest Tools.
func (api *EsxiVmApi) UninstallNutanixGuestTools(extId *string, args ...map[string]interface{}) (*import4.UninstallNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UninstallNutanixGuestTools(context.Background(), &import5.UninstallNutanixGuestToolsRequest{
		ExtId: extId,
	}, args...)
}

// Trigger an in-guest uninstallation of Nutanix Guest Tools.
func (api *EsxiVmServiceApi) UninstallNutanixGuestTools(ctx context.Context, request *import5.UninstallNutanixGuestToolsRequest, args ...map[string]interface{}) (*import4.UninstallNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/uninstall"

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

	unmarshalledResp := new(import4.UninstallNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Nutanix Guest Tools configuration for a virtual machine.
func (api *EsxiVmApi) UpdateNutanixGuestToolsById(extId *string, body *import4.NutanixGuestTools, args ...map[string]interface{}) (*import4.UpdateNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateNutanixGuestToolsById(context.Background(), &import5.UpdateNutanixGuestToolsByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Nutanix Guest Tools configuration for a virtual machine.
func (api *EsxiVmServiceApi) UpdateNutanixGuestToolsById(ctx context.Context, request *import5.UpdateNutanixGuestToolsByIdRequest, args ...map[string]interface{}) (*import4.UpdateNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools"

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

	unmarshalledResp := new(import4.UpdateNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Trigger an in-guest upgrade of Nutanix Guest Tools.
func (api *EsxiVmApi) UpgradeNutanixGuestTools(extId *string, body *import4.NutanixGuestToolsUpgradeConfig, args ...map[string]interface{}) (*import4.UpgradeNutanixGuestToolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEsxiVmServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpgradeNutanixGuestTools(context.Background(), &import5.UpgradeNutanixGuestToolsRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Trigger an in-guest upgrade of Nutanix Guest Tools.
func (api *EsxiVmServiceApi) UpgradeNutanixGuestTools(ctx context.Context, request *import5.UpgradeNutanixGuestToolsRequest, args ...map[string]interface{}) (*import4.UpgradeNutanixGuestToolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/upgrade"

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

	unmarshalledResp := new(import4.UpgradeNutanixGuestToolsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
