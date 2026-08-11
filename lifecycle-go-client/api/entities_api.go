package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import8 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/entities"
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

// Export the entity inventory information of connected clusters as a downloadable file. The exported file contains the full list of LCM entities (software and firmware) discovered across one or more clusters, including their versions, locations, and available updates. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the exported file is available in the completion_details field of the task. Submit an ExportInventorySpec with the desired file format (CSV). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *EntitiesApi) ExportInventory(body *import1.ExportInventorySpec, args ...map[string]interface{}) (*import1.ExportInventoryApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportInventory(context.Background(), &import8.ExportInventoryRequest{
		Body: body,
	}, args...)
}

// Export the entity inventory information of connected clusters as a downloadable file. The exported file contains the full list of LCM entities (software and firmware) discovered across one or more clusters, including their versions, locations, and available updates. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the exported file is available in the completion_details field of the task. Submit an ExportInventorySpec with the desired file format (CSV). Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *EntitiesServiceApi) ExportInventory(ctx context.Context, request *import8.ExportInventoryRequest, args ...map[string]interface{}) (*import1.ExportInventoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/entities/$actions/export"

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
	unmarshalledResp := new(import1.ExportInventoryApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the full details of a single LCM entity by its external identifier. The response includes the entity class, model, type (SOFTWARE or FIRMWARE), current installed version, target upgrade version (if set), available versions with dependency and release-note information, location (node or cluster scope), device identifier, hardware vendor, and sub-entity details. Entities are populated by the inventory operation (POST /$actions/inventory); if no inventory has been run, this endpoint returns a 404. To plan an upgrade for this entity, use POST /$actions/compute-recommendations.
func (api *EntitiesApi) GetEntityById(extId *string, args ...map[string]interface{}) (*import1.GetEntityByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntityById(context.Background(), &import8.GetEntityByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the full details of a single LCM entity by its external identifier. The response includes the entity class, model, type (SOFTWARE or FIRMWARE), current installed version, target upgrade version (if set), available versions with dependency and release-note information, location (node or cluster scope), device identifier, hardware vendor, and sub-entity details. Entities are populated by the inventory operation (POST /$actions/inventory); if no inventory has been run, this endpoint returns a 404. To plan an upgrade for this entity, use POST /$actions/compute-recommendations.
func (api *EntitiesServiceApi) GetEntityById(ctx context.Context, request *import8.GetEntityByIdRequest, args ...map[string]interface{}) (*import1.GetEntityByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/entities/{extId}"

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
	unmarshalledResp := new(import1.GetEntityByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a paginated list of all LCM entities discovered on the cluster. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Each entity includes its class, model, type, current version, available upgrade versions, and location. Entities are populated by the inventory operation (POST /$actions/inventory); if no inventory has been run, the list will be empty. Use this endpoint to browse available updates before computing recommendations (POST /$actions/compute-recommendations) or performing an upgrade (POST /$actions/upgrade).
func (api *EntitiesApi) ListEntities(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListEntitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEntities(context.Background(), &import8.ListEntitiesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieve a paginated list of all LCM entities discovered on the cluster. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Each entity includes its class, model, type, current version, available upgrade versions, and location. Entities are populated by the inventory operation (POST /$actions/inventory); if no inventory has been run, the list will be empty. Use this endpoint to browse available updates before computing recommendations (POST /$actions/compute-recommendations) or performing an upgrade (POST /$actions/upgrade).
func (api *EntitiesServiceApi) ListEntities(ctx context.Context, request *import8.ListEntitiesRequest, args ...map[string]interface{}) (*import1.ListEntitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/entities"

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
	unmarshalledResp := new(import1.ListEntitiesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Pre-stage upgrade images on the cluster before performing an upgrade. Preloading downloads the required images from the LCM repository (or catalog) to the cluster nodes so that they are immediately available when the upgrade is executed, reducing upgrade downtime. Submit a PreloadSpec containing entity update specifications (entity UUID and target version pairs) identifying which entities' images to preload. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; poll the task to track download progress. After preloading, the isImagePresent field on the entity's available versions (GET /entities) will reflect that images are ready. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *EntitiesApi) PreloadArtifacts(body *import6.PreloadSpec, xClusterId *string, args ...map[string]interface{}) (*import7.PreloadArtifactsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PreloadArtifacts(context.Background(), &import8.PreloadArtifactsRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Pre-stage upgrade images on the cluster before performing an upgrade. Preloading downloads the required images from the LCM repository (or catalog) to the cluster nodes so that they are immediately available when the upgrade is executed, reducing upgrade downtime. Submit a PreloadSpec containing entity update specifications (entity UUID and target version pairs) identifying which entities' images to preload. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference; poll the task to track download progress. After preloading, the isImagePresent field on the entity's available versions (GET /entities) will reflect that images are ready. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *EntitiesServiceApi) PreloadArtifacts(ctx context.Context, request *import8.PreloadArtifactsRequest, args ...map[string]interface{}) (*import7.PreloadArtifactsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/preload-artifacts"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import7.PreloadArtifactsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
