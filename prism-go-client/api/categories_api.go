package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/categories"
	"net/http"
	"net/url"
	"strings"
)

type CategoriesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *CategoriesServiceApi
}

type CategoriesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCategoriesApi(apiClient *client.ApiClient) *CategoriesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CategoriesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewCategoriesServiceApi(a.ApiClient)

	return a
}

func NewCategoriesServiceApi(apiClient *client.ApiClient) *CategoriesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CategoriesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a category with a given key and value pair. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 201 response with the created category.
func (api *CategoriesApi) CreateCategory(body *import3.Category, args ...map[string]interface{}) (*import3.CreateCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCategory(context.Background(), &import4.CreateCategoryRequest{
		Body: body,
	}, args...)
}

// Creates a category with a given key and value pair. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 201 response with the created category.
func (api *CategoriesServiceApi) CreateCategory(ctx context.Context, request *import4.CreateCategoryRequest, args ...map[string]interface{}) (*import3.CreateCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories"

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
	unmarshalledResp := new(import3.CreateCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes a category with the given external identifier. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 204 response indicating successful deletion.
func (api *CategoriesApi) DeleteCategoryById(extId *string, args ...map[string]interface{}) (*import3.DeleteCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCategoryById(context.Background(), &import4.DeleteCategoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a category with the given external identifier. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 204 response indicating successful deletion.
func (api *CategoriesServiceApi) DeleteCategoryById(ctx context.Context, request *import4.DeleteCategoryByIdRequest, args ...map[string]interface{}) (*import3.DeleteCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories/{extId}"

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
	unmarshalledResp := new(import3.DeleteCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetches the details of a category with the given external identifier. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central).
func (api *CategoriesApi) GetCategoryById(extId *string, expand_ *string, args ...map[string]interface{}) (*import3.GetCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCategoryById(context.Background(), &import4.GetCategoryByIdRequest{
		ExtId:   extId,
		Expand_: expand_,
	}, args...)
}

// Fetches the details of a category with the given external identifier. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central).
func (api *CategoriesServiceApi) GetCategoryById(ctx context.Context, request *import4.GetCategoryByIdRequest, args ...map[string]interface{}) (*import3.GetCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.GetCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetches a list of categories with pagination, filtering, sorting, selection, and optional expansion of associated entity counts. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central).<br> Note: This endpoint only supports $expand=associations. For detailed associations with category extId, use the GET /categories/{extId}?$expand=detailedAssociations endpoint.
func (api *CategoriesApi) ListCategories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, apply_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListCategoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCategories(context.Background(), &import4.ListCategoriesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Apply_:   apply_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of categories with pagination, filtering, sorting, selection, and optional expansion of associated entity counts. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central).<br> Note: This endpoint only supports $expand=associations. For detailed associations with category extId, use the GET /categories/{extId}?$expand=detailedAssociations endpoint.
func (api *CategoriesServiceApi) ListCategories(ctx context.Context, request *import4.ListCategoriesRequest, args ...map[string]interface{}) (*import3.ListCategoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories"

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
	if request.Apply_ != nil {
		queryParams.Add("$apply", client.ParameterToString(*request.Apply_, ""))
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.ListCategoriesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Shares a category with a specified project. This allows the project to access the category.
func (api *CategoriesApi) ShareCategory(categoryExtId *string, body *import3.ShareCategoryRequest, args ...map[string]interface{}) (*import3.ShareCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShareCategory(context.Background(), &import4.ShareCategoryRequest{
		CategoryExtId: categoryExtId,
		Body:          body,
	}, args...)
}

// Shares a category with a specified project. This allows the project to access the category.
func (api *CategoriesServiceApi) ShareCategory(ctx context.Context, request *import4.ShareCategoryRequest, args ...map[string]interface{}) (*import3.ShareCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories/{categoryExtId}/$actions/share"

	// verify the required parameter 'categoryExtId' is set
	if nil == request.CategoryExtId {
		return nil, client.ReportError("categoryExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"categoryExtId"+"}", url.PathEscape(client.ParameterToString(*request.CategoryExtId, "")), -1)
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
	unmarshalledResp := new(import3.ShareCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Unshares a category from a specified project. This removes the project's access to the category.
func (api *CategoriesApi) UnshareCategory(categoryExtId *string, body *import3.UnshareCategoryRequest, args ...map[string]interface{}) (*import3.UnshareCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnshareCategory(context.Background(), &import4.UnshareCategoryRequest{
		CategoryExtId: categoryExtId,
		Body:          body,
	}, args...)
}

// Unshares a category from a specified project. This removes the project's access to the category.
func (api *CategoriesServiceApi) UnshareCategory(ctx context.Context, request *import4.UnshareCategoryRequest, args ...map[string]interface{}) (*import3.UnshareCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories/{categoryExtId}/$actions/unshare"

	// verify the required parameter 'categoryExtId' is set
	if nil == request.CategoryExtId {
		return nil, client.ReportError("categoryExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"categoryExtId"+"}", url.PathEscape(client.ParameterToString(*request.CategoryExtId, "")), -1)
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
	unmarshalledResp := new(import3.UnshareCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the description, value, and owner properties of a category. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 200 response with the update status.
func (api *CategoriesApi) UpdateCategoryById(extId *string, body *import3.Category, args ...map[string]interface{}) (*import3.UpdateCategoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCategoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateCategoryById(context.Background(), &import4.UpdateCategoryByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the description, value, and owner properties of a category. This endpoint supports both multidomain categories on NC (Nutanix Central) and local categories on PC (Prism Central). For multidomain categories on NC, this operation is asynchronous and returns a 202 response with a task reference. For local categories on PC, this operation is synchronous and returns a 200 response with the update status.
func (api *CategoriesServiceApi) UpdateCategoryById(ctx context.Context, request *import4.UpdateCategoryByIdRequest, args ...map[string]interface{}) (*import3.UpdateCategoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.4/config/categories/{extId}"

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
	unmarshalledResp := new(import3.UpdateCategoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
