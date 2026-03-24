package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import19 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/networkcontrollers"
	"net/http"
	"net/url"
	"strings"
)

type NetworkControllersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *NetworkControllersServiceApi
}

type NetworkControllersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewNetworkControllersApi(apiClient *client.ApiClient) *NetworkControllersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NetworkControllersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewNetworkControllersServiceApi(a.ApiClient)

	return a
}

func NewNetworkControllersServiceApi(apiClient *client.ApiClient) *NetworkControllersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NetworkControllersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a network controller. For large Prism Centrals, an additional 3GB of memory and 3 vCPUs are required, for each Prism Central node. For small Prism Centrals, an additional 1GB of memory and 1 vCPU is required, for each Prism Central node.
func (api *NetworkControllersApi) CreateNetworkController(body *import4.NetworkController, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkControllersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateNetworkController(context.Background(), &import19.CreateNetworkControllerRequest{
		Body: body,
	}, args...)
}

// Create a network controller. For large Prism Centrals, an additional 3GB of memory and 3 vCPUs are required, for each Prism Central node. For small Prism Centrals, an additional 1GB of memory and 1 vCPU is required, for each Prism Central node.
func (api *NetworkControllersServiceApi) CreateNetworkController(ctx context.Context, request *import19.CreateNetworkControllerRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/controllers"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete a network controller.
func (api *NetworkControllersApi) DeleteNetworkControllerById(extId *string, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkControllersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteNetworkControllerById(context.Background(), &import19.DeleteNetworkControllerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete a network controller.
func (api *NetworkControllersServiceApi) DeleteNetworkControllerById(ctx context.Context, request *import19.DeleteNetworkControllerByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/controllers/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the network controller with the specified UUID.
func (api *NetworkControllersApi) GetNetworkControllerById(extId *string, args ...map[string]interface{}) (*import4.GetNetworkControllerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkControllersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNetworkControllerById(context.Background(), &import19.GetNetworkControllerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get the network controller with the specified UUID.
func (api *NetworkControllersServiceApi) GetNetworkControllerById(ctx context.Context, request *import19.GetNetworkControllerByIdRequest, args ...map[string]interface{}) (*import4.GetNetworkControllerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/controllers/{extId}"

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

	unmarshalledResp := new(import4.GetNetworkControllerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the list of existing network controllers.
func (api *NetworkControllersApi) ListNetworkControllers(page_ *int, limit_ *int, args ...map[string]interface{}) (*import4.ListNetworkControllersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkControllersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNetworkControllers(context.Background(), &import19.ListNetworkControllersRequest{
		Page_:  page_,
		Limit_: limit_,
	}, args...)
}

// Gets the list of existing network controllers.
func (api *NetworkControllersServiceApi) ListNetworkControllers(ctx context.Context, request *import19.ListNetworkControllersRequest, args ...map[string]interface{}) (*import4.ListNetworkControllersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/controllers"

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

	unmarshalledResp := new(import4.ListNetworkControllersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update a network controller.
func (api *NetworkControllersApi) UpdateNetworkControllerById(extId *string, body *import4.NetworkController, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkControllersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateNetworkControllerById(context.Background(), &import19.UpdateNetworkControllerByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update a network controller.
func (api *NetworkControllersServiceApi) UpdateNetworkControllerById(ctx context.Context, request *import19.UpdateNetworkControllerByIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/controllers/{extId}"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
