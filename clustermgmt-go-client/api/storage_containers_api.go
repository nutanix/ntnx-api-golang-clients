package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import14 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/storagecontainers"
	import5 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/stats"
	import6 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StorageContainersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *StorageContainersServiceApi
}

type StorageContainersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewStorageContainersApi(apiClient *client.ApiClient) *StorageContainersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageContainersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewStorageContainersServiceApi(a.ApiClient)

	return a
}

func NewStorageContainersServiceApi(apiClient *client.ApiClient) *StorageContainersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StorageContainersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Clears the thick provisioned space of the provided Storage Container. The location header received in the API response contains the URL of the task object, which can be further used to track the status of the request.
func (api *StorageContainersApi) ClearThickProvisionedSpace(extId *string, xClusterId *string, args ...map[string]interface{}) (*import1.ClearThickProvisionedSpaceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ClearThickProvisionedSpace(context.Background(), &import14.ClearThickProvisionedSpaceRequest{
		ExtId:      extId,
		XClusterId: xClusterId,
	}, args...)
}

// Clears the thick provisioned space of the provided Storage Container. The location header received in the API response contains the URL of the task object, which can be further used to track the status of the request.
func (api *StorageContainersServiceApi) ClearThickProvisionedSpace(ctx context.Context, request *import14.ClearThickProvisionedSpaceRequest, args ...map[string]interface{}) (*import1.ClearThickProvisionedSpaceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}/$actions/clear-thick-provisioned-space"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.ClearThickProvisionedSpaceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a new Storage Container with the specified configuration on the cluster identified by cluster’s external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersApi) CreateStorageContainer(body *import1.StorageContainer, xClusterId *string, args ...map[string]interface{}) (*import1.CreateStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateStorageContainer(context.Background(), &import14.CreateStorageContainerRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Creates a new Storage Container with the specified configuration on the cluster identified by cluster’s external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersServiceApi) CreateStorageContainer(ctx context.Context, request *import14.CreateStorageContainerRequest, args ...map[string]interface{}) (*import1.CreateStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'xClusterId' is set
	if nil == request.XClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
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

	unmarshalledResp := new(import1.CreateStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes an existing Storage Container identified by external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersApi) DeleteStorageContainerById(extId *string, ignoreSmallFiles *bool, args ...map[string]interface{}) (*import1.DeleteStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteStorageContainerById(context.Background(), &import14.DeleteStorageContainerByIdRequest{
		ExtId:            extId,
		IgnoreSmallFiles: ignoreSmallFiles,
	}, args...)
}

// Deletes an existing Storage Container identified by external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersServiceApi) DeleteStorageContainerById(ctx context.Context, request *import14.DeleteStorageContainerByIdRequest, args ...map[string]interface{}) (*import1.DeleteStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}"

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
	if request.IgnoreSmallFiles != nil {
		queryParams.Add("ignoreSmallFiles", client.ParameterToString(*request.IgnoreSmallFiles, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the configuration details of an existing storage container identified by the external identifier. Note: The Storage Containers of PEs with versions prior to AOS 6.5 might have missing attribute data, and the PEs under the self-owned RBAC category might not be visible to non-admin users.
func (api *StorageContainersApi) GetStorageContainerById(extId *string, args ...map[string]interface{}) (*import1.GetStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStorageContainerById(context.Background(), &import14.GetStorageContainerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the configuration details of an existing storage container identified by the external identifier. Note: The Storage Containers of PEs with versions prior to AOS 6.5 might have missing attribute data, and the PEs under the self-owned RBAC category might not be visible to non-admin users.
func (api *StorageContainersServiceApi) GetStorageContainerById(ctx context.Context, request *import14.GetStorageContainerByIdRequest, args ...map[string]interface{}) (*import1.GetStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}"

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

	unmarshalledResp := new(import1.GetStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the statistical information for the Storage Container identified by external identifier..
func (api *StorageContainersApi) GetStorageContainerStats(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import6.DownSamplingOperator, args ...map[string]interface{}) (*import5.GetStorageContainerStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetStorageContainerStats(context.Background(), &import14.GetStorageContainerStatsRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
	}, args...)
}

// Fetches the statistical information for the Storage Container identified by external identifier..
func (api *StorageContainersServiceApi) GetStorageContainerStats(ctx context.Context, request *import14.GetStorageContainerStatsRequest, args ...map[string]interface{}) (*import5.GetStorageContainerStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/stats/storage-containers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
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
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.GetStorageContainerStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the datastores associated with Storage Containers from a cluster.
func (api *StorageContainersApi) ListDataStoresByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.ListDataStoresByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDataStoresByClusterId(context.Background(), &import14.ListDataStoresByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
	}, args...)
}

// Lists all the datastores associated with Storage Containers from a cluster.
func (api *StorageContainersServiceApi) ListDataStoresByClusterId(ctx context.Context, request *import14.ListDataStoresByClusterIdRequest, args ...map[string]interface{}) (*import1.ListDataStoresByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/storage-containers/datastores"

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

	unmarshalledResp := new(import1.ListDataStoresByClusterIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the Storage Containers available in the cluster. Note: The Storage Containers of PEs with versions prior to AOS 6.5 might have missing attribute data, and the PEs under the self-owned RBAC category might not be visible to non-admin users.
func (api *StorageContainersApi) ListStorageContainers(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListStorageContainersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListStorageContainers(context.Background(), &import14.ListStorageContainersRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the Storage Containers available in the cluster. Note: The Storage Containers of PEs with versions prior to AOS 6.5 might have missing attribute data, and the PEs under the self-owned RBAC category might not be visible to non-admin users.
func (api *StorageContainersServiceApi) ListStorageContainers(ctx context.Context, request *import14.ListStorageContainersRequest, args ...map[string]interface{}) (*import1.ListStorageContainersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers"

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

	unmarshalledResp := new(import1.ListStorageContainersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Mounts the Storage Container identified by external identifier on an ESX datastore. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersApi) MountStorageContainer(extId *string, body *import1.DataStoreMount, args ...map[string]interface{}) (*import1.MountStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MountStorageContainer(context.Background(), &import14.MountStorageContainerRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Mounts the Storage Container identified by external identifier on an ESX datastore. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersServiceApi) MountStorageContainer(ctx context.Context, request *import14.MountStorageContainerRequest, args ...map[string]interface{}) (*import1.MountStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}/$actions/mount"

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

	unmarshalledResp := new(import1.MountStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Unmounts the Storage Container identified by external identifier from the ESX datastore. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersApi) UnmountStorageContainer(extId *string, body *import1.DataStoreUnmount, args ...map[string]interface{}) (*import1.UnmountStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnmountStorageContainer(context.Background(), &import14.UnmountStorageContainerRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Unmounts the Storage Container identified by external identifier from the ESX datastore. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersServiceApi) UnmountStorageContainer(ctx context.Context, request *import14.UnmountStorageContainerRequest, args ...map[string]interface{}) (*import1.UnmountStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}/$actions/unmount"

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

	unmarshalledResp := new(import1.UnmountStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the configuration of an existing Storage Container identified by external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersApi) UpdateStorageContainerById(extId *string, body *import1.StorageContainer, args ...map[string]interface{}) (*import1.UpdateStorageContainerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewStorageContainersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateStorageContainerById(context.Background(), &import14.UpdateStorageContainerByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the configuration of an existing Storage Container identified by external identifier. The location header received in the API response contains the URL of the task object, which can be used to further track the status of the request.
func (api *StorageContainersServiceApi) UpdateStorageContainerById(ctx context.Context, request *import14.UpdateStorageContainerByIdRequest, args ...map[string]interface{}) (*import1.UpdateStorageContainerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/storage-containers/{extId}"

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

	unmarshalledResp := new(import1.UpdateStorageContainerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
