package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import14 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/externalstorages"
	"net/http"
	"net/url"
	"strings"
)

type ExternalStoragesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ExternalStoragesServiceApi
}

type ExternalStoragesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewExternalStoragesApi(apiClient *client.ApiClient) *ExternalStoragesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalStoragesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewExternalStoragesServiceApi(a.ApiClient)

	return a
}

func NewExternalStoragesServiceApi(apiClient *client.ApiClient) *ExternalStoragesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ExternalStoragesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new external storage.
func (api *ExternalStoragesApi) CreateExternalStorage(body *import1.ExternalStorage, args ...map[string]interface{}) (*import1.CreateExternalStorageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalStoragesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateExternalStorage(context.Background(), &import14.CreateExternalStorageRequest{
		Body: body,
	}, args...)
}

// Creates a new external storage.
func (api *ExternalStoragesServiceApi) CreateExternalStorage(ctx context.Context, request *import14.CreateExternalStorageRequest, args ...map[string]interface{}) (*import1.CreateExternalStorageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/external-storages"

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
	unmarshalledResp := new(import1.CreateExternalStorageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves details of a specific external storage by its unique identifier. For external storages on PE clusters older than 7.6, `health` may be `$UNKNOWN`, `ETag` and `ownerExtId` will not be returned.
func (api *ExternalStoragesApi) GetExternalStorageById(extId *string, args ...map[string]interface{}) (*import1.GetExternalStorageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalStoragesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetExternalStorageById(context.Background(), &import14.GetExternalStorageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves details of a specific external storage by its unique identifier. For external storages on PE clusters older than 7.6, `health` may be `$UNKNOWN`, `ETag` and `ownerExtId` will not be returned.
func (api *ExternalStoragesServiceApi) GetExternalStorageById(ctx context.Context, request *import14.GetExternalStorageByIdRequest, args ...map[string]interface{}) (*import1.GetExternalStorageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/external-storages/{extId}"

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
	unmarshalledResp := new(import1.GetExternalStorageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves a list of external storages. For external storages on PE clusters older than 7.6, `health` may be `$UNKNOWN`, `ETag` and `ownerExtId` will not be returned.
func (api *ExternalStoragesApi) ListExternalStorages(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, args ...map[string]interface{}) (*import1.GetExternalStorageListApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalStoragesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListExternalStorages(context.Background(), &import14.ListExternalStoragesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
	}, args...)
}

// Retrieves a list of external storages. For external storages on PE clusters older than 7.6, `health` may be `$UNKNOWN`, `ETag` and `ownerExtId` will not be returned.
func (api *ExternalStoragesServiceApi) ListExternalStorages(ctx context.Context, request *import14.ListExternalStoragesRequest, args ...map[string]interface{}) (*import1.GetExternalStorageListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/external-storages"

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
	unmarshalledResp := new(import1.GetExternalStorageListApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates details of a specific external storage by its unique identifier.
func (api *ExternalStoragesApi) UpdateExternalStorageById(extId *string, body *import1.ExternalStorage, args ...map[string]interface{}) (*import1.UpdateExternalStorageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewExternalStoragesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateExternalStorageById(context.Background(), &import14.UpdateExternalStorageByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates details of a specific external storage by its unique identifier.
func (api *ExternalStoragesServiceApi) UpdateExternalStorageById(ctx context.Context, request *import14.UpdateExternalStorageByIdRequest, args ...map[string]interface{}) (*import1.UpdateExternalStorageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/external-storages/{extId}"

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
	unmarshalledResp := new(import1.UpdateExternalStorageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
