package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	import18 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/templates"
	"net/http"
	"net/url"
	"strings"
)

type TemplatesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *TemplatesServiceApi
}

type TemplatesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewTemplatesApi(apiClient *client.ApiClient) *TemplatesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TemplatesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewTemplatesServiceApi(a.ApiClient)

	return a
}

func NewTemplatesServiceApi(apiClient *client.ApiClient) *TemplatesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TemplatesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// This operation cancels the update initiated by the \"Initiate guest OS update\" API for the given template. The temporary VM created during the update process is deleted, and the pending update status is cleared. Please note that any modifications made to the temporary VM will be lost upon cancelling the update operation.
func (api *TemplatesApi) CancelGuestUpdate(extId *string, args ...map[string]interface{}) (*import9.CancelGuestUpdateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CancelGuestUpdate(context.Background(), &import18.CancelGuestUpdateRequest{
		ExtId: extId,
	}, args...)
}

// This operation cancels the update initiated by the \"Initiate guest OS update\" API for the given template. The temporary VM created during the update process is deleted, and the pending update status is cleared. Please note that any modifications made to the temporary VM will be lost upon cancelling the update operation.
func (api *TemplatesServiceApi) CancelGuestUpdate(ctx context.Context, request *import18.CancelGuestUpdateRequest, args ...map[string]interface{}) (*import9.CancelGuestUpdateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}/$actions/cancel-guest-update"

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

	unmarshalledResp := new(import9.CancelGuestUpdateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This operation finalizes the update initiated by the \"Initiate guest OS update\" API for the template. A new version is added to the template, the temporary VM created during the update process is deleted, and the pending update status is cleared.
func (api *TemplatesApi) CompleteGuestUpdate(extId *string, body *import9.CompleteGuestUpdateSpec, args ...map[string]interface{}) (*import9.CompleteGuestUpdateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CompleteGuestUpdate(context.Background(), &import18.CompleteGuestUpdateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// This operation finalizes the update initiated by the \"Initiate guest OS update\" API for the template. A new version is added to the template, the temporary VM created during the update process is deleted, and the pending update status is cleared.
func (api *TemplatesServiceApi) CompleteGuestUpdate(ctx context.Context, request *import18.CompleteGuestUpdateRequest, args ...map[string]interface{}) (*import9.CompleteGuestUpdateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}/$actions/complete-guest-update"

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

	unmarshalledResp := new(import9.CompleteGuestUpdateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a template from the given VM identifier. A template stores the VM configuration and disks from the source VM.
func (api *TemplatesApi) CreateTemplate(body *import9.Template, args ...map[string]interface{}) (*import9.CreateTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateTemplate(context.Background(), &import18.CreateTemplateRequest{
		Body: body,
	}, args...)
}

// Creates a template from the given VM identifier. A template stores the VM configuration and disks from the source VM.
func (api *TemplatesServiceApi) CreateTemplate(ctx context.Context, request *import18.CreateTemplateRequest, args ...map[string]interface{}) (*import9.CreateTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates"

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

	unmarshalledResp := new(import9.CreateTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the template and all of its versions for a given template identifier.
func (api *TemplatesApi) DeleteTemplateById(extId *string, args ...map[string]interface{}) (*import9.DeleteTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteTemplateById(context.Background(), &import18.DeleteTemplateByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the template and all of its versions for a given template identifier.
func (api *TemplatesServiceApi) DeleteTemplateById(ctx context.Context, request *import18.DeleteTemplateByIdRequest, args ...map[string]interface{}) (*import9.DeleteTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}"

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

	unmarshalledResp := new(import9.DeleteTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a version for the given template version identifier.
func (api *TemplatesApi) DeleteTemplateVersionById(templateExtId *string, extId *string, args ...map[string]interface{}) (*import9.DeleteTemplateVersionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteTemplateVersionById(context.Background(), &import18.DeleteTemplateVersionByIdRequest{
		TemplateExtId: templateExtId,
		ExtId:         extId,
	}, args...)
}

// Deletes a version for the given template version identifier.
func (api *TemplatesServiceApi) DeleteTemplateVersionById(ctx context.Context, request *import18.DeleteTemplateVersionByIdRequest, args ...map[string]interface{}) (*import9.DeleteTemplateVersionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{templateExtId}/versions/{extId}"

	// verify the required parameter 'templateExtId' is set
	if nil == request.TemplateExtId {
		return nil, client.ReportError("templateExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(*request.TemplateExtId, "")), -1)
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

	unmarshalledResp := new(import9.DeleteTemplateVersionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deploys one or more VMs from a template. You can specify the number of VMs to deploy and their corresponding VM configuration overrides.
func (api *TemplatesApi) DeployTemplate(extId *string, body *import9.TemplateDeployment, args ...map[string]interface{}) (*import9.DeployTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeployTemplate(context.Background(), &import18.DeployTemplateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Deploys one or more VMs from a template. You can specify the number of VMs to deploy and their corresponding VM configuration overrides.
func (api *TemplatesServiceApi) DeployTemplate(ctx context.Context, request *import18.DeployTemplateRequest, args ...map[string]interface{}) (*import9.DeployTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}/$actions/deploy"

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

	unmarshalledResp := new(import9.DeployTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the template details for a given template identifier.
func (api *TemplatesApi) GetTemplateById(extId *string, args ...map[string]interface{}) (*import9.GetTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTemplateById(context.Background(), &import18.GetTemplateByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the template details for a given template identifier.
func (api *TemplatesServiceApi) GetTemplateById(ctx context.Context, request *import18.GetTemplateByIdRequest, args ...map[string]interface{}) (*import9.GetTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}"

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

	unmarshalledResp := new(import9.GetTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the template version details of the given template and version identifier.
func (api *TemplatesApi) GetTemplateVersionById(templateExtId *string, extId *string, args ...map[string]interface{}) (*import9.GetTemplateVersionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTemplateVersionById(context.Background(), &import18.GetTemplateVersionByIdRequest{
		TemplateExtId: templateExtId,
		ExtId:         extId,
	}, args...)
}

// Retrieves the template version details of the given template and version identifier.
func (api *TemplatesServiceApi) GetTemplateVersionById(ctx context.Context, request *import18.GetTemplateVersionByIdRequest, args ...map[string]interface{}) (*import9.GetTemplateVersionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{templateExtId}/versions/{extId}"

	// verify the required parameter 'templateExtId' is set
	if nil == request.TemplateExtId {
		return nil, client.ReportError("templateExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(*request.TemplateExtId, "")), -1)
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

	unmarshalledResp := new(import9.GetTemplateVersionApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Initiates the process of updating the Guest OS from a given template version identifier. Only one guest OS update can be initiated for a template at a time. A temporary VM is created where the guest OS updates will be applied. The user must make the necessary modifications to this temporary VM. After completing the modifications, the user should issue the \"Complete Guest OS Update\" command to finalize the update. The \"Cancel Guest OS Update\" command can be issued at any time to abort an ongoing update.
func (api *TemplatesApi) InitiateGuestUpdate(extId *string, body *import9.InitiateGuestUpdateSpec, args ...map[string]interface{}) (*import9.InitiateGuestUpdateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.InitiateGuestUpdate(context.Background(), &import18.InitiateGuestUpdateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Initiates the process of updating the Guest OS from a given template version identifier. Only one guest OS update can be initiated for a template at a time. A temporary VM is created where the guest OS updates will be applied. The user must make the necessary modifications to this temporary VM. After completing the modifications, the user should issue the \"Complete Guest OS Update\" command to finalize the update. The \"Cancel Guest OS Update\" command can be issued at any time to abort an ongoing update.
func (api *TemplatesServiceApi) InitiateGuestUpdate(ctx context.Context, request *import18.InitiateGuestUpdateRequest, args ...map[string]interface{}) (*import9.InitiateGuestUpdateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}/$actions/initiate-guest-update"

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

	unmarshalledResp := new(import9.InitiateGuestUpdateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the versions of a template with details such as name, description, VM configuration and so on. This operation supports filtering, sorting and pagination.
func (api *TemplatesApi) ListTemplateVersions(templateExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import9.ListTemplateVersionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTemplateVersions(context.Background(), &import18.ListTemplateVersionsRequest{
		TemplateExtId: templateExtId,
		Page_:         page_,
		Limit_:        limit_,
		Filter_:       filter_,
		Orderby_:      orderby_,
		Select_:       select_,
	}, args...)
}

// Lists all the versions of a template with details such as name, description, VM configuration and so on. This operation supports filtering, sorting and pagination.
func (api *TemplatesServiceApi) ListTemplateVersions(ctx context.Context, request *import18.ListTemplateVersionsRequest, args ...map[string]interface{}) (*import9.ListTemplateVersionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{templateExtId}/versions"

	// verify the required parameter 'templateExtId' is set
	if nil == request.TemplateExtId {
		return nil, client.ReportError("templateExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(*request.TemplateExtId, "")), -1)
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

	unmarshalledResp := new(import9.ListTemplateVersionsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists templates with details such as name, description, VM configuration and so on. This operation supports filtering, sorting and pagination.
func (api *TemplatesApi) ListTemplates(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import9.ListTemplatesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTemplates(context.Background(), &import18.ListTemplatesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists templates with details such as name, description, VM configuration and so on. This operation supports filtering, sorting and pagination.
func (api *TemplatesServiceApi) ListTemplates(ctx context.Context, request *import18.ListTemplatesRequest, args ...map[string]interface{}) (*import9.ListTemplatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates"

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

	unmarshalledResp := new(import9.ListTemplatesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Designate the template version identified by the given identifier as the active version. An active version is a default version for creating VMs from the template and starting the Guest OS Update.
func (api *TemplatesApi) PublishTemplate(extId *string, body *import9.TemplatePublishSpec, args ...map[string]interface{}) (*import9.PublishTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PublishTemplate(context.Background(), &import18.PublishTemplateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Designate the template version identified by the given identifier as the active version. An active version is a default version for creating VMs from the template and starting the Guest OS Update.
func (api *TemplatesServiceApi) PublishTemplate(ctx context.Context, request *import18.PublishTemplateRequest, args ...map[string]interface{}) (*import9.PublishTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}/$actions/publish"

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

	unmarshalledResp := new(import9.PublishTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a template with the given template identifier. This operation updates the template configuration and/or adds a new version to the template. Unless otherwise specified, the newly created version is set as the active version.
func (api *TemplatesApi) UpdateTemplateById(extId *string, body *import9.Template, args ...map[string]interface{}) (*import9.UpdateTemplateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateTemplateById(context.Background(), &import18.UpdateTemplateByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a template with the given template identifier. This operation updates the template configuration and/or adds a new version to the template. Unless otherwise specified, the newly created version is set as the active version.
func (api *TemplatesServiceApi) UpdateTemplateById(ctx context.Context, request *import18.UpdateTemplateByIdRequest, args ...map[string]interface{}) (*import9.UpdateTemplateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/templates/{extId}"

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

	unmarshalledResp := new(import9.UpdateTemplateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
