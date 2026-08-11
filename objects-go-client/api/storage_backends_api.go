package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
	import10 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/request/storagebackends"
	"net/http"
	"net/url"
	"strings"
)

type StorageBackendsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StorageBackendsServiceApi
}

type StorageBackendsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStorageBackendsApi(apiClient *client.ApiClient) *StorageBackendsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageBackendsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStorageBackendsServiceApi(a.ApiClient)

	return a
}

func NewStorageBackendsServiceApi(apiClient *client.ApiClient) *StorageBackendsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageBackendsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the details of the specified storage backend.
func (api *StorageBackendsApi) GetStorageBackendById(objectStoreExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetStorageBackendApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageBackendsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStorageBackendById(context.Background(), &import10.GetStorageBackendByIdRequest{
		ObjectStoreExtId: objectStoreExtId,
		ExtId:            extId,
	}, args...)
}

// Get the details of the specified storage backend.
func (api *StorageBackendsServiceApi) GetStorageBackendById(ctx context.Context, request *import10.GetStorageBackendByIdRequest, args ...map[string]interface{}) (*import1.GetStorageBackendApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/storage-backends/{extId}"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetStorageBackendApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get a list of the storage backends that are configured in an Object store. Maximum of 5 storage backends are supported per Object store.
func (api *StorageBackendsApi) ListStorageBackends(objectStoreExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListStorageBackendsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageBackendsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListStorageBackends(context.Background(), &import10.ListStorageBackendsRequest{
		ObjectStoreExtId: objectStoreExtId,
		Page_:            page_,
		Limit_:           limit_,
		Filter_:          filter_,
		Orderby_:         orderby_,
		Select_:          select_,
	}, args...)
}

// Get a list of the storage backends that are configured in an Object store. Maximum of 5 storage backends are supported per Object store.
func (api *StorageBackendsServiceApi) ListStorageBackends(ctx context.Context, request *import10.ListStorageBackendsRequest, args ...map[string]interface{}) (*import1.ListStorageBackendsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/storage-backends"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
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
	unmarshalledResp := new(import1.ListStorageBackendsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
