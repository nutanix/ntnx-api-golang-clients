package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/externalrepositories"
	"net/http"
	"net/url"
	"strings"
)

type ExternalRepositoriesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ExternalRepositoriesServiceApi
}

type ExternalRepositoriesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewExternalRepositoriesApi(apiClient *client.ApiClient) *ExternalRepositoriesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalRepositoriesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewExternalRepositoriesServiceApi(a.ApiClient)

	return a
}

func NewExternalRepositoriesServiceApi(apiClient *client.ApiClient) *ExternalRepositoriesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalRepositoriesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create an external repository configuration for instant recovery from backup solutions.
func (api *ExternalRepositoriesApi) CreateExternalRepository(body *import3.ExternalRepository, args ...map[string]interface{}) (*import3.CreateExternalRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateExternalRepository(context.Background(), &import4.CreateExternalRepositoryRequest{
		Body: body,
	}, args...)
}

// Create an external repository configuration for instant recovery from backup solutions.
func (api *ExternalRepositoriesServiceApi) CreateExternalRepository(ctx context.Context, request *import4.CreateExternalRepositoryRequest, args ...map[string]interface{}) (*import3.CreateExternalRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/external-repositories"

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
	unmarshalledResp := new(import3.CreateExternalRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes an external repository by ID.
func (api *ExternalRepositoriesApi) DeleteExternalRepositoryById(extId *string, args ...map[string]interface{}) (*import3.DeleteExternalRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteExternalRepositoryById(context.Background(), &import4.DeleteExternalRepositoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes an external repository by ID.
func (api *ExternalRepositoriesServiceApi) DeleteExternalRepositoryById(ctx context.Context, request *import4.DeleteExternalRepositoryByIdRequest, args ...map[string]interface{}) (*import3.DeleteExternalRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/external-repositories/{extId}"

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
	unmarshalledResp := new(import3.DeleteExternalRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Gets an external repository by ID.
func (api *ExternalRepositoriesApi) GetExternalRepositoryById(extId *string, args ...map[string]interface{}) (*import3.GetExternalRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetExternalRepositoryById(context.Background(), &import4.GetExternalRepositoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets an external repository by ID.
func (api *ExternalRepositoriesServiceApi) GetExternalRepositoryById(ctx context.Context, request *import4.GetExternalRepositoryByIdRequest, args ...map[string]interface{}) (*import3.GetExternalRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/external-repositories/{extId}"

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
	unmarshalledResp := new(import3.GetExternalRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists the external repositories configured for instant recovery from backup solutions.
func (api *ExternalRepositoriesApi) ListExternalRepositories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListExternalRepositoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListExternalRepositories(context.Background(), &import4.ListExternalRepositoriesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the external repositories configured for instant recovery from backup solutions.
func (api *ExternalRepositoriesServiceApi) ListExternalRepositories(ctx context.Context, request *import4.ListExternalRepositoriesRequest, args ...map[string]interface{}) (*import3.ListExternalRepositoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/external-repositories"

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
	unmarshalledResp := new(import3.ListExternalRepositoriesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update the details of an external repository for the provided external identifier.
func (api *ExternalRepositoriesApi) UpdateExternalRepositoryById(extId *string, body *import3.ExternalRepository, args ...map[string]interface{}) (*import3.UpdateExternalRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateExternalRepositoryById(context.Background(), &import4.UpdateExternalRepositoryByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the details of an external repository for the provided external identifier.
func (api *ExternalRepositoriesServiceApi) UpdateExternalRepositoryById(ctx context.Context, request *import4.UpdateExternalRepositoryByIdRequest, args ...map[string]interface{}) (*import3.UpdateExternalRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/external-repositories/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.UpdateExternalRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
