package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import21 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/storageconfig"
	"net/http"
	"net/url"
	"strings"
)

type StorageConfigApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StorageConfigServiceApi
}

type StorageConfigServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStorageConfigApi(apiClient *client.ApiClient) *StorageConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageConfigApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStorageConfigServiceApi(a.ApiClient)

	return a
}

func NewStorageConfigServiceApi(apiClient *client.ApiClient) *StorageConfigServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageConfigServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieves the storage configuration for the cluster identified by {clusterExtId}.
func (api *StorageConfigApi) GetStorageConfig(clusterExtId *string, args ...map[string]interface{}) (*import1.GetStorageConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStorageConfig(context.Background(), &import21.GetStorageConfigRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// Retrieves the storage configuration for the cluster identified by {clusterExtId}.
func (api *StorageConfigServiceApi) GetStorageConfig(ctx context.Context, request *import21.GetStorageConfigRequest, args ...map[string]interface{}) (*import1.GetStorageConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/storage-config"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetStorageConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the storage configuration for the cluster identified by {clusterExtId}.
func (api *StorageConfigApi) UpdateStorageConfig(clusterExtId *string, body *import1.StorageConfig, args ...map[string]interface{}) (*import1.UpdateStorageConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageConfigServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateStorageConfig(context.Background(), &import21.UpdateStorageConfigRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Updates the storage configuration for the cluster identified by {clusterExtId}.
func (api *StorageConfigServiceApi) UpdateStorageConfig(ctx context.Context, request *import21.UpdateStorageConfigRequest, args ...map[string]interface{}) (*import1.UpdateStorageConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/storage-config"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.UpdateStorageConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
