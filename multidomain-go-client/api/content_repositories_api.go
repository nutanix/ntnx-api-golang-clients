package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/content"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/contentrepositories"
	"net/http"
	"net/url"
	"strings"
)

type ContentRepositoriesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ContentRepositoriesServiceApi
}

type ContentRepositoriesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewContentRepositoriesApi(apiClient *client.ApiClient) *ContentRepositoriesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ContentRepositoriesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewContentRepositoriesServiceApi(a.ApiClient)

	return a
}

func NewContentRepositoriesServiceApi(apiClient *client.ApiClient) *ContentRepositoriesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ContentRepositoriesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a Content Repository to manage content, such as disk images, in the form of Repository items. Name is a mandatory field to create a Content Repository.
func (api *ContentRepositoriesApi) CreateContentRepository(body *import1.ContentRepository, args ...map[string]interface{}) (*import1.CreateContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateContentRepository(context.Background(), &import2.CreateContentRepositoryRequest{
		Body: body,
	}, args...)
}

// Create a Content Repository to manage content, such as disk images, in the form of Repository items. Name is a mandatory field to create a Content Repository.
func (api *ContentRepositoriesServiceApi) CreateContentRepository(ctx context.Context, request *import2.CreateContentRepositoryRequest, args ...map[string]interface{}) (*import1.CreateContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories"

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
	unmarshalledResp := new(import1.CreateContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete the Content Repository with the given external identifier.
func (api *ContentRepositoriesApi) DeleteContentRepositoryById(extId *string, args ...map[string]interface{}) (*import1.DeleteContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteContentRepositoryById(context.Background(), &import2.DeleteContentRepositoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the Content Repository with the given external identifier.
func (api *ContentRepositoriesServiceApi) DeleteContentRepositoryById(ctx context.Context, request *import2.DeleteContentRepositoryByIdRequest, args ...map[string]interface{}) (*import1.DeleteContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}"

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
	unmarshalledResp := new(import1.DeleteContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get the Content Repository details for the provided external identifier.
func (api *ContentRepositoriesApi) GetContentRepositoryById(extId *string, args ...map[string]interface{}) (*import1.GetContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetContentRepositoryById(context.Background(), &import2.GetContentRepositoryByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get the Content Repository details for the provided external identifier.
func (api *ContentRepositoriesServiceApi) GetContentRepositoryById(ctx context.Context, request *import2.GetContentRepositoryByIdRequest, args ...map[string]interface{}) (*import1.GetContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}"

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
	unmarshalledResp := new(import1.GetContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List Content Repositories owned by the user along with Content Repository details like name, description, etc. This operation supports filtering, sorting, selection, and pagination.
func (api *ContentRepositoriesApi) ListContentRepositories(xClusterId *string, xProjectId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListContentRepositoriesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListContentRepositories(context.Background(), &import2.ListContentRepositoriesRequest{
		XClusterId: xClusterId,
		XProjectId: xProjectId,
		Page_:      page_,
		Limit_:     limit_,
		Filter_:    filter_,
		Orderby_:   orderby_,
		Select_:    select_,
	}, args...)
}

// List Content Repositories owned by the user along with Content Repository details like name, description, etc. This operation supports filtering, sorting, selection, and pagination.
func (api *ContentRepositoriesServiceApi) ListContentRepositories(ctx context.Context, request *import2.ListContentRepositoriesRequest, args ...map[string]interface{}) (*import1.ListContentRepositoriesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories"

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
	if request.XClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
	}
	if request.XProjectId != nil {
		headerParams["X-Project-Id"] = client.ParameterToString(*request.XProjectId, "")
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
	unmarshalledResp := new(import1.ListContentRepositoriesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List Content Repository items owned by the user along with details like name, description, etc. This operation supports filtering, sorting, selection and pagination.
func (api *ContentRepositoriesApi) ListContentRepositoryItemsByContentRepositoryId(extId *string, xClusterId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListContentRepositoryItemsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListContentRepositoryItemsByContentRepositoryId(context.Background(), &import2.ListContentRepositoryItemsByContentRepositoryIdRequest{
		ExtId:      extId,
		XClusterId: xClusterId,
		Page_:      page_,
		Limit_:     limit_,
		Filter_:    filter_,
		Orderby_:   orderby_,
		Select_:    select_,
	}, args...)
}

// List Content Repository items owned by the user along with details like name, description, etc. This operation supports filtering, sorting, selection and pagination.
func (api *ContentRepositoriesServiceApi) ListContentRepositoryItemsByContentRepositoryId(ctx context.Context, request *import2.ListContentRepositoryItemsByContentRepositoryIdRequest, args ...map[string]interface{}) (*import1.ListContentRepositoryItemsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}/items"

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
	if request.XClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
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
	unmarshalledResp := new(import1.ListContentRepositoryItemsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Subscribe to a source Content Repository to replicate content.
func (api *ContentRepositoriesApi) SubscribeContentRepository(extId *string, body *import1.ContentRepositorySubscriptionSpec, args ...map[string]interface{}) (*import1.SubscribeContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SubscribeContentRepository(context.Background(), &import2.SubscribeContentRepositoryRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Subscribe to a source Content Repository to replicate content.
func (api *ContentRepositoriesServiceApi) SubscribeContentRepository(ctx context.Context, request *import2.SubscribeContentRepositoryRequest, args ...map[string]interface{}) (*import1.SubscribeContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}/$actions/subscribe"

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
	unmarshalledResp := new(import1.SubscribeContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Sync the items of a subscribed Content Repository by fetching the latest state of items from the publisher.
func (api *ContentRepositoriesApi) SyncContentRepository(extId *string, args ...map[string]interface{}) (*import1.SyncContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SyncContentRepository(context.Background(), &import2.SyncContentRepositoryRequest{
		ExtId: extId,
	}, args...)
}

// Sync the items of a subscribed Content Repository by fetching the latest state of items from the publisher.
func (api *ContentRepositoriesServiceApi) SyncContentRepository(ctx context.Context, request *import2.SyncContentRepositoryRequest, args ...map[string]interface{}) (*import1.SyncContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}/$actions/sync"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.SyncContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update the Content Repository with the given external identifier, using the provided request body. To make sure the correct ETag is used, it is always recommended to do a GET on a resource before doing a PUT.
func (api *ContentRepositoriesApi) UpdateContentRepositoryById(extId *string, body *import1.ContentRepository, args ...map[string]interface{}) (*import1.UpdateContentRepositoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewContentRepositoriesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateContentRepositoryById(context.Background(), &import2.UpdateContentRepositoryByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the Content Repository with the given external identifier, using the provided request body. To make sure the correct ETag is used, it is always recommended to do a GET on a resource before doing a PUT.
func (api *ContentRepositoriesServiceApi) UpdateContentRepositoryById(ctx context.Context, request *import2.UpdateContentRepositoryByIdRequest, args ...map[string]interface{}) (*import1.UpdateContentRepositoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/content/content-repositories/{extId}"

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
	unmarshalledResp := new(import1.UpdateContentRepositoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
