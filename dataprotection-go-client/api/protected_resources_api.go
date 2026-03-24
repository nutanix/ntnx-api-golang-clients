package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/request/protectedresources"
	"net/http"
	"net/url"
	"strings"
)

type ProtectedResourcesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ProtectedResourcesServiceApi
}

type ProtectedResourcesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewProtectedResourcesApi(apiClient *client.ApiClient) *ProtectedResourcesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ProtectedResourcesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewProtectedResourcesServiceApi(a.ApiClient)

	return a
}

func NewProtectedResourcesServiceApi(apiClient *client.ApiClient) *ProtectedResourcesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ProtectedResourcesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieves the details of the specified protection resource, including available restorable time ranges on the local Prism Central and the replication status to targets defined in the applied protection policies. This is applicable only if the entity is protected with a minutely or synchronous schedule. Other protection schedules are not currently supported by this endpoint and are considered unprotected.
func (api *ProtectedResourcesApi) GetProtectedResourceById(extId *string, args ...map[string]interface{}) (*import1.GetProtectedResourceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectedResourcesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetProtectedResourceById(context.Background(), &import3.GetProtectedResourceByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the details of the specified protection resource, including available restorable time ranges on the local Prism Central and the replication status to targets defined in the applied protection policies. This is applicable only if the entity is protected with a minutely or synchronous schedule. Other protection schedules are not currently supported by this endpoint and are considered unprotected.
func (api *ProtectedResourcesServiceApi) GetProtectedResourceById(ctx context.Context, request *import3.GetProtectedResourceByIdRequest, args ...map[string]interface{}) (*import1.GetProtectedResourceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/protected-resources/{extId}"

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

	unmarshalledResp := new(import1.GetProtectedResourceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Promotes a specified synced entity at the target site. This is only relevant if the synced entity is protected in a synchronous schedule.
func (api *ProtectedResourcesApi) PromoteProtectedResource(extId *string, args ...map[string]interface{}) (*import1.ProtectedResourcePromoteApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectedResourcesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PromoteProtectedResource(context.Background(), &import3.PromoteProtectedResourceRequest{
		ExtId: extId,
	}, args...)
}

// Promotes a specified synced entity at the target site. This is only relevant if the synced entity is protected in a synchronous schedule.
func (api *ProtectedResourcesServiceApi) PromoteProtectedResource(ctx context.Context, request *import3.PromoteProtectedResourceRequest, args ...map[string]interface{}) (*import1.ProtectedResourcePromoteApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/protected-resources/{extId}/$actions/promote"

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

	unmarshalledResp := new(import1.ProtectedResourcePromoteApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Restores the specified protected resource from its state at the given timestamp on the given cluster. This is only relevant if the entity is protected in a minutely schedule at the given timestamp.
func (api *ProtectedResourcesApi) RestoreProtectedResource(extId *string, body *import1.ProtectedResourceRestoreSpec, args ...map[string]interface{}) (*import1.ProtectedResourceRestoreApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectedResourcesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RestoreProtectedResource(context.Background(), &import3.RestoreProtectedResourceRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Restores the specified protected resource from its state at the given timestamp on the given cluster. This is only relevant if the entity is protected in a minutely schedule at the given timestamp.
func (api *ProtectedResourcesServiceApi) RestoreProtectedResource(ctx context.Context, request *import3.RestoreProtectedResourceRequest, args ...map[string]interface{}) (*import1.ProtectedResourceRestoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/protected-resources/{extId}/$actions/restore"

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

	unmarshalledResp := new(import1.ProtectedResourceRestoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
