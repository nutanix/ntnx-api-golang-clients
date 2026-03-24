package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import6 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/entities"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type EntitiesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EntitiesServiceApi
}

type EntitiesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEntitiesApi(apiClient *client.ApiClient) *EntitiesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntitiesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEntitiesServiceApi(a.ApiClient)

	return a
}

func NewEntitiesServiceApi(apiClient *client.ApiClient) *EntitiesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntitiesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Download the history and inventory information of connected clusters as a file. This API will return a task, which on completion will provide a URL in the completion_details field to download the file containing the inventory information.
func (api *EntitiesApi) ExportInventory(body *import1.ExportInventorySpec, args ...map[string]interface{}) (*import1.ExportInventoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportInventory(context.Background(), &import6.ExportInventoryRequest{
		Body: body,
	}, args...)
}

// Download the history and inventory information of connected clusters as a file. This API will return a task, which on completion will provide a URL in the completion_details field to download the file containing the inventory information.
func (api *EntitiesServiceApi) ExportInventory(ctx context.Context, request *import6.ExportInventoryRequest, args ...map[string]interface{}) (*import1.ExportInventoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/entities/$actions/export"

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

	unmarshalledResp := new(import1.ExportInventoryApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get details about an LCM entity.
func (api *EntitiesApi) GetEntityById(extId *string, args ...map[string]interface{}) (*import1.GetEntityByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityById(context.Background(), &import6.GetEntityByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get details about an LCM entity.
func (api *EntitiesServiceApi) GetEntityById(ctx context.Context, request *import6.GetEntityByIdRequest, args ...map[string]interface{}) (*import1.GetEntityByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/entities/{extId}"

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

	unmarshalledResp := new(import1.GetEntityByIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query list of LCM entities.
func (api *EntitiesApi) ListEntities(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListEntitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEntities(context.Background(), &import6.ListEntitiesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Query list of LCM entities.
func (api *EntitiesServiceApi) ListEntities(ctx context.Context, request *import6.ListEntitiesRequest, args ...map[string]interface{}) (*import1.ListEntitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/resources/entities"

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

	unmarshalledResp := new(import1.ListEntitiesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Preload artifacts required for upgrade on to the cluster.
func (api *EntitiesApi) PreloadArtifacts(body *import4.PreloadSpec, xClusterId *string, args ...map[string]interface{}) (*import5.PreloadArtifactsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PreloadArtifacts(context.Background(), &import6.PreloadArtifactsRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Preload artifacts required for upgrade on to the cluster.
func (api *EntitiesServiceApi) PreloadArtifacts(ctx context.Context, request *import6.PreloadArtifactsRequest, args ...map[string]interface{}) (*import5.PreloadArtifactsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/operations/$actions/preload-artifacts"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.PreloadArtifactsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
