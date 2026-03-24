package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/request/storagepolicies"
	"net/http"
	"net/url"
	"strings"
)

type StoragePoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StoragePoliciesServiceApi
}

type StoragePoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStoragePoliciesApi(apiClient *client.ApiClient) *StoragePoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StoragePoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStoragePoliciesServiceApi(a.ApiClient)

	return a
}

func NewStoragePoliciesServiceApi(apiClient *client.ApiClient) *StoragePoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StoragePoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new Storage Policy to simplify storage management at scale.
func (api *StoragePoliciesApi) CreateStoragePolicy(body *import1.StoragePolicy, args ...map[string]interface{}) (*import1.CreateStoragePolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStoragePoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateStoragePolicy(context.Background(), &import5.CreateStoragePolicyRequest{
		Body: body,
	}, args...)
}

// Creates a new Storage Policy to simplify storage management at scale.
func (api *StoragePoliciesServiceApi) CreateStoragePolicy(ctx context.Context, request *import5.CreateStoragePolicyRequest, args ...map[string]interface{}) (*import1.CreateStoragePolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/storage-policies"

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

	unmarshalledResp := new(import1.CreateStoragePolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes an existing Storage Policy identified by its external identifier.
func (api *StoragePoliciesApi) DeleteStoragePolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteStoragePolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStoragePoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteStoragePolicyById(context.Background(), &import5.DeleteStoragePolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes an existing Storage Policy identified by its external identifier.
func (api *StoragePoliciesServiceApi) DeleteStoragePolicyById(ctx context.Context, request *import5.DeleteStoragePolicyByIdRequest, args ...map[string]interface{}) (*import1.DeleteStoragePolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/storage-policies/{extId}"

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

	unmarshalledResp := new(import1.DeleteStoragePolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the Storage Policy identified by its external identifier.
func (api *StoragePoliciesApi) GetStoragePolicyById(extId *string, args ...map[string]interface{}) (*import1.GetStoragePolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStoragePoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStoragePolicyById(context.Background(), &import5.GetStoragePolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the Storage Policy identified by its external identifier.
func (api *StoragePoliciesServiceApi) GetStoragePolicyById(ctx context.Context, request *import5.GetStoragePolicyByIdRequest, args ...map[string]interface{}) (*import1.GetStoragePolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/storage-policies/{extId}"

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

	unmarshalledResp := new(import1.GetStoragePolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of Storage Policies. This operation supports filtering, sorting, selection, and pagination.
func (api *StoragePoliciesApi) ListStoragePolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListStoragePoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStoragePoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListStoragePolicies(context.Background(), &import5.ListStoragePoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of Storage Policies. This operation supports filtering, sorting, selection, and pagination.
func (api *StoragePoliciesServiceApi) ListStoragePolicies(ctx context.Context, request *import5.ListStoragePoliciesRequest, args ...map[string]interface{}) (*import1.ListStoragePoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/storage-policies"

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

	unmarshalledResp := new(import1.ListStoragePoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates an existing Storage Policy identified by its external identifier.
func (api *StoragePoliciesApi) UpdateStoragePolicyById(extId *string, body *import1.StoragePolicy, args ...map[string]interface{}) (*import1.UpdateStoragePolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStoragePoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateStoragePolicyById(context.Background(), &import5.UpdateStoragePolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates an existing Storage Policy identified by its external identifier.
func (api *StoragePoliciesServiceApi) UpdateStoragePolicyById(ctx context.Context, request *import5.UpdateStoragePolicyByIdRequest, args ...map[string]interface{}) (*import1.UpdateStoragePolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/storage-policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateStoragePolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
