package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import7 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/directoryservices"
	"net/http"
	"net/url"
	"strings"
)

type DirectoryServicesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DirectoryServicesServiceApi
}

type DirectoryServicesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDirectoryServicesApi(apiClient *client.ApiClient) *DirectoryServicesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DirectoryServicesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDirectoryServicesServiceApi(a.ApiClient)

	return a
}

func NewDirectoryServicesServiceApi(apiClient *client.ApiClient) *DirectoryServicesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DirectoryServicesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Checks the connection to the directory service.
func (api *DirectoryServicesApi) ConnectionStatusDirectoryService(extId *string, body *import4.DirectoryServiceConnectionRequest, args ...map[string]interface{}) (*import4.ConnectionDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ConnectionStatusDirectoryService(context.Background(), &import7.ConnectionStatusDirectoryServiceRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Checks the connection to the directory service.
func (api *DirectoryServicesServiceApi) ConnectionStatusDirectoryService(ctx context.Context, request *import7.ConnectionStatusDirectoryServiceRequest, args ...map[string]interface{}) (*import4.ConnectionDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services/{extId}/$actions/verify-connection-status"

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

	unmarshalledResp := new(import4.ConnectionDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a directory service.
func (api *DirectoryServicesApi) CreateDirectoryService(body *import4.DirectoryService, args ...map[string]interface{}) (*import4.CreateDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDirectoryService(context.Background(), &import7.CreateDirectoryServiceRequest{
		Body: body,
	}, args...)
}

// Creates a directory service.
func (api *DirectoryServicesServiceApi) CreateDirectoryService(ctx context.Context, request *import7.CreateDirectoryServiceRequest, args ...map[string]interface{}) (*import4.CreateDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services"

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

	unmarshalledResp := new(import4.CreateDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a directory service.
func (api *DirectoryServicesApi) DeleteDirectoryServiceById(extId *string, args ...map[string]interface{}) (*import4.DeleteDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDirectoryServiceById(context.Background(), &import7.DeleteDirectoryServiceByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a directory service.
func (api *DirectoryServicesServiceApi) DeleteDirectoryServiceById(ctx context.Context, request *import7.DeleteDirectoryServiceByIdRequest, args ...map[string]interface{}) (*import4.DeleteDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services/{extId}"

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

	unmarshalledResp := new(import4.DeleteDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a directory service.
func (api *DirectoryServicesApi) GetDirectoryServiceById(extId *string, args ...map[string]interface{}) (*import4.GetDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDirectoryServiceById(context.Background(), &import7.GetDirectoryServiceByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a directory service.
func (api *DirectoryServicesServiceApi) GetDirectoryServiceById(ctx context.Context, request *import7.GetDirectoryServiceByIdRequest, args ...map[string]interface{}) (*import4.GetDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services/{extId}"

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

	unmarshalledResp := new(import4.GetDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all directory service(s).
func (api *DirectoryServicesApi) ListDirectoryServices(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListDirectoryServicesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDirectoryServices(context.Background(), &import7.ListDirectoryServicesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists all directory service(s).
func (api *DirectoryServicesServiceApi) ListDirectoryServices(ctx context.Context, request *import7.ListDirectoryServicesRequest, args ...map[string]interface{}) (*import4.ListDirectoryServicesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services"

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

	unmarshalledResp := new(import4.ListDirectoryServicesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Searches a user or group in the directory service through its external identifier.
func (api *DirectoryServicesApi) SearchDirectoryService(extId *string, body *import4.DirectoryServiceSearchQuery, args ...map[string]interface{}) (*import4.SearchDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SearchDirectoryService(context.Background(), &import7.SearchDirectoryServiceRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Searches a user or group in the directory service through its external identifier.
func (api *DirectoryServicesServiceApi) SearchDirectoryService(ctx context.Context, request *import7.SearchDirectoryServiceRequest, args ...map[string]interface{}) (*import4.SearchDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services/{extId}/$actions/search"

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

	unmarshalledResp := new(import4.SearchDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a directory service.
func (api *DirectoryServicesApi) UpdateDirectoryServiceById(extId *string, body *import4.DirectoryService, args ...map[string]interface{}) (*import4.UpdateDirectoryServiceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDirectoryServicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateDirectoryServiceById(context.Background(), &import7.UpdateDirectoryServiceByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a directory service.
func (api *DirectoryServicesServiceApi) UpdateDirectoryServiceById(ctx context.Context, request *import7.UpdateDirectoryServiceByIdRequest, args ...map[string]interface{}) (*import4.UpdateDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/directory-services/{extId}"

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

	unmarshalledResp := new(import4.UpdateDirectoryServiceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
