package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/request/recoverypointstores"
	"net/http"
	"net/url"
	"strings"
)

type RecoveryPointStoresApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RecoveryPointStoresServiceApi
}

type RecoveryPointStoresServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRecoveryPointStoresApi(apiClient *client.ApiClient) *RecoveryPointStoresApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPointStoresApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRecoveryPointStoresServiceApi(a.ApiClient)

	return a
}

func NewRecoveryPointStoresServiceApi(apiClient *client.ApiClient) *RecoveryPointStoresServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPointStoresServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieve a recovery point store by its external identifier.
func (api *RecoveryPointStoresApi) GetRecoveryPointStoreById(extId *string, args ...map[string]interface{}) (*import1.GetRecoveryPointStoreApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointStoresServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecoveryPointStoreById(context.Background(), &import7.GetRecoveryPointStoreByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve a recovery point store by its external identifier.
func (api *RecoveryPointStoresServiceApi) GetRecoveryPointStoreById(ctx context.Context, request *import7.GetRecoveryPointStoreByIdRequest, args ...map[string]interface{}) (*import1.GetRecoveryPointStoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.4/config/recovery-point-stores/{extId}"

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
	unmarshalledResp := new(import1.GetRecoveryPointStoreApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all recovery point stores with support for pagination, filtering, sorting, and field selection.
func (api *RecoveryPointStoresApi) ListRecoveryPointStores(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRecoveryPointStoresApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointStoresServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecoveryPointStores(context.Background(), &import7.ListRecoveryPointStoresRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List all recovery point stores with support for pagination, filtering, sorting, and field selection.
func (api *RecoveryPointStoresServiceApi) ListRecoveryPointStores(ctx context.Context, request *import7.ListRecoveryPointStoresRequest, args ...map[string]interface{}) (*import1.ListRecoveryPointStoresApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.4/config/recovery-point-stores"

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
	unmarshalledResp := new(import1.ListRecoveryPointStoresApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
