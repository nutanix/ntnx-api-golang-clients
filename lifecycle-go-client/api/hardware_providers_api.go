package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
	import11 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/hardwareproviders"
	"net/http"
	"net/url"
	"strings"
)

type HardwareProvidersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *HardwareProvidersServiceApi
}

type HardwareProvidersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewHardwareProvidersApi(apiClient *client.ApiClient) *HardwareProvidersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &HardwareProvidersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewHardwareProvidersServiceApi(a.ApiClient)

	return a
}

func NewHardwareProvidersServiceApi(apiClient *client.ApiClient) *HardwareProvidersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &HardwareProvidersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a connection to an endpoint for a specific hardware provider
func (api *HardwareProvidersApi) CreateConnectionByHardwareProviderId(hardwareProviderExtId *string, body *import3.Connection, args ...map[string]interface{}) (*import3.CreateConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateConnectionByHardwareProviderId(context.Background(), &import11.CreateConnectionByHardwareProviderIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		Body:                  body,
	}, args...)
}

// Creates a connection to an endpoint for a specific hardware provider
func (api *HardwareProvidersServiceApi) CreateConnectionByHardwareProviderId(ctx context.Context, request *import11.CreateConnectionByHardwareProviderIdRequest, args ...map[string]interface{}) (*import3.CreateConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	unmarshalledResp := new(import3.CreateConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes a hardware provider connection identified by its external ID
func (api *HardwareProvidersApi) DeleteConnectionById(hardwareProviderExtId *string, extId *string, args ...map[string]interface{}) (*import3.DeleteConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteConnectionById(context.Background(), &import11.DeleteConnectionByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ExtId:                 extId,
	}, args...)
}

// Deletes a hardware provider connection identified by its external ID
func (api *HardwareProvidersServiceApi) DeleteConnectionById(ctx context.Context, request *import11.DeleteConnectionByIdRequest, args ...map[string]interface{}) (*import3.DeleteConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	unmarshalledResp := new(import3.DeleteConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a hardware provider connection identified by its external ID
func (api *HardwareProvidersApi) GetConnectionById(hardwareProviderExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetConnectionById(context.Background(), &import11.GetConnectionByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ExtId:                 extId,
	}, args...)
}

// Returns details of a hardware provider connection identified by its external ID
func (api *HardwareProvidersServiceApi) GetConnectionById(ctx context.Context, request *import11.GetConnectionByIdRequest, args ...map[string]interface{}) (*import3.GetConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a node discovered through a specific hardware provider connection
func (api *HardwareProvidersApi) GetConnectionNodeById(hardwareProviderExtId *string, connectionExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetConnectionNodeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetConnectionNodeById(context.Background(), &import11.GetConnectionNodeByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		ExtId:                 extId,
	}, args...)
}

// Returns details of a node discovered through a specific hardware provider connection
func (api *HardwareProvidersServiceApi) GetConnectionNodeById(ctx context.Context, request *import11.GetConnectionNodeByIdRequest, args ...map[string]interface{}) (*import3.GetConnectionNodeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/nodes/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetConnectionNodeApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a hardware provider identified by its external ID
func (api *HardwareProvidersApi) GetHardwareProviderById(extId *string, args ...map[string]interface{}) (*import3.GetHardwareProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetHardwareProviderById(context.Background(), &import11.GetHardwareProviderByIdRequest{
		ExtId: extId,
	}, args...)
}

// Returns details of a hardware provider identified by its external ID
func (api *HardwareProvidersServiceApi) GetHardwareProviderById(ctx context.Context, request *import11.GetHardwareProviderByIdRequest, args ...map[string]interface{}) (*import3.GetHardwareProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{extId}"

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
	unmarshalledResp := new(import3.GetHardwareProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of an IP address pool from a hardware provider connection
func (api *HardwareProvidersApi) GetIpPoolById(hardwareProviderExtId *string, connectionExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetIpPoolApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetIpPoolById(context.Background(), &import11.GetIpPoolByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		ExtId:                 extId,
	}, args...)
}

// Returns details of an IP address pool from a hardware provider connection
func (api *HardwareProvidersServiceApi) GetIpPoolById(ctx context.Context, request *import11.GetIpPoolByIdRequest, args ...map[string]interface{}) (*import3.GetIpPoolApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/ip-pools/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetIpPoolApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a MAC address pool from a hardware provider connection
func (api *HardwareProvidersApi) GetMacPoolById(hardwareProviderExtId *string, connectionExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetMacPoolApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetMacPoolById(context.Background(), &import11.GetMacPoolByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		ExtId:                 extId,
	}, args...)
}

// Returns details of a MAC address pool from a hardware provider connection
func (api *HardwareProvidersServiceApi) GetMacPoolById(ctx context.Context, request *import11.GetMacPoolByIdRequest, args ...map[string]interface{}) (*import3.GetMacPoolApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/mac-pools/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetMacPoolApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a server identity pool from a hardware provider connection
func (api *HardwareProvidersApi) GetServerIdentityPoolById(hardwareProviderExtId *string, connectionExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetServerIdentityPoolApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetServerIdentityPoolById(context.Background(), &import11.GetServerIdentityPoolByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		ExtId:                 extId,
	}, args...)
}

// Returns details of a server identity pool from a hardware provider connection
func (api *HardwareProvidersServiceApi) GetServerIdentityPoolById(ctx context.Context, request *import11.GetServerIdentityPoolByIdRequest, args ...map[string]interface{}) (*import3.GetServerIdentityPoolApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/server-identity-pools/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetServerIdentityPoolApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a paginated list of connections for a specific hardware provider
func (api *HardwareProvidersApi) ListConnectionsByHardwareProviderId(hardwareProviderExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListConnectionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListConnectionsByHardwareProviderId(context.Background(), &import11.ListConnectionsByHardwareProviderIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Returns a paginated list of connections for a specific hardware provider
func (api *HardwareProvidersServiceApi) ListConnectionsByHardwareProviderId(ctx context.Context, request *import11.ListConnectionsByHardwareProviderIdRequest, args ...map[string]interface{}) (*import3.ListConnectionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListConnectionsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a paginated list of all available hardware providers
func (api *HardwareProvidersApi) ListHardwareProviders(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListHardwareProvidersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListHardwareProviders(context.Background(), &import11.ListHardwareProvidersRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Returns a paginated list of all available hardware providers
func (api *HardwareProvidersServiceApi) ListHardwareProviders(ctx context.Context, request *import11.ListHardwareProvidersRequest, args ...map[string]interface{}) (*import3.ListHardwareProvidersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers"

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
	unmarshalledResp := new(import3.ListHardwareProvidersApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a list of IP address pools available in a hardware provider connection
func (api *HardwareProvidersApi) ListIpPoolsByConnectionId(hardwareProviderExtId *string, connectionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListIpPoolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListIpPoolsByConnectionId(context.Background(), &import11.ListIpPoolsByConnectionIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Returns a list of IP address pools available in a hardware provider connection
func (api *HardwareProvidersServiceApi) ListIpPoolsByConnectionId(ctx context.Context, request *import11.ListIpPoolsByConnectionIdRequest, args ...map[string]interface{}) (*import3.ListIpPoolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/ip-pools"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListIpPoolsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a list of MAC address pools available in a hardware provider connection
func (api *HardwareProvidersApi) ListMacPoolsByConnectionId(hardwareProviderExtId *string, connectionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListMacPoolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListMacPoolsByConnectionId(context.Background(), &import11.ListMacPoolsByConnectionIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Returns a list of MAC address pools available in a hardware provider connection
func (api *HardwareProvidersServiceApi) ListMacPoolsByConnectionId(ctx context.Context, request *import11.ListMacPoolsByConnectionIdRequest, args ...map[string]interface{}) (*import3.ListMacPoolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/mac-pools"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListMacPoolsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a list of nodes discovered through a specific hardware provider connection
func (api *HardwareProvidersApi) ListNodesByConnectionId(hardwareProviderExtId *string, connectionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListConnectionNodesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNodesByConnectionId(context.Background(), &import11.ListNodesByConnectionIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Returns a list of nodes discovered through a specific hardware provider connection
func (api *HardwareProvidersServiceApi) ListNodesByConnectionId(ctx context.Context, request *import11.ListNodesByConnectionIdRequest, args ...map[string]interface{}) (*import3.ListConnectionNodesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/nodes"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListConnectionNodesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a list of server identity pools available in a hardware provider connection
func (api *HardwareProvidersApi) ListServerIdentityPoolsByConnectionId(hardwareProviderExtId *string, connectionExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListServerIdentityPoolsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListServerIdentityPoolsByConnectionId(context.Background(), &import11.ListServerIdentityPoolsByConnectionIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ConnectionExtId:       connectionExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Returns a list of server identity pools available in a hardware provider connection
func (api *HardwareProvidersServiceApi) ListServerIdentityPoolsByConnectionId(ctx context.Context, request *import11.ListServerIdentityPoolsByConnectionIdRequest, args ...map[string]interface{}) (*import3.ListServerIdentityPoolsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{connectionExtId}/server-identity-pools"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
	}
	// verify the required parameter 'connectionExtId' is set
	if nil == request.ConnectionExtId {
		return nil, client.ReportError("connectionExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"connectionExtId"+"}", url.PathEscape(client.ParameterToString(*request.ConnectionExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListServerIdentityPoolsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Refresh nodes or resources from a hardware provider connection
func (api *HardwareProvidersApi) RefreshConnection(hardwareProviderExtId *string, extId *string, body *import3.RefreshConnectionSpec, args ...map[string]interface{}) (*import3.RefreshConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RefreshConnection(context.Background(), &import11.RefreshConnectionRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ExtId:                 extId,
		Body:                  body,
	}, args...)
}

// Refresh nodes or resources from a hardware provider connection
func (api *HardwareProvidersServiceApi) RefreshConnection(ctx context.Context, request *import11.RefreshConnectionRequest, args ...map[string]interface{}) (*import3.RefreshConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId}/$actions/refresh"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	unmarshalledResp := new(import3.RefreshConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the details of a hardware provider connection identified by its external ID
func (api *HardwareProvidersApi) UpdateConnectionById(hardwareProviderExtId *string, extId *string, body *import3.Connection, args ...map[string]interface{}) (*import3.UpdateConnectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewHardwareProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateConnectionById(context.Background(), &import11.UpdateConnectionByIdRequest{
		HardwareProviderExtId: hardwareProviderExtId,
		ExtId:                 extId,
		Body:                  body,
	}, args...)
}

// Updates the details of a hardware provider connection identified by its external ID
func (api *HardwareProvidersServiceApi) UpdateConnectionById(ctx context.Context, request *import11.UpdateConnectionByIdRequest, args ...map[string]interface{}) (*import3.UpdateConnectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/hardware-providers/{hardwareProviderExtId}/connections/{extId}"

	// verify the required parameter 'hardwareProviderExtId' is set
	if nil == request.HardwareProviderExtId {
		return nil, client.ReportError("hardwareProviderExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"hardwareProviderExtId"+"}", url.PathEscape(client.ParameterToString(*request.HardwareProviderExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.UpdateConnectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
