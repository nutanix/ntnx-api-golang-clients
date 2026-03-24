package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import8 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/config"
	import9 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/content"
	import10 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/request/recoverypoints"
	"net/http"
	"net/url"
	"strings"
)

type RecoveryPointsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RecoveryPointsServiceApi
}

type RecoveryPointsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRecoveryPointsApi(apiClient *client.ApiClient) *RecoveryPointsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPointsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRecoveryPointsServiceApi(a.ApiClient)

	return a
}

func NewRecoveryPointsServiceApi(apiClient *client.ApiClient) *RecoveryPointsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecoveryPointsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a new recovery point.<br> #### Task Completion Details <br> External identifier of the created recovery point can be found in the task completion details under the key `recoveryPointExtId`.
func (api *RecoveryPointsApi) CreateRecoveryPoint(body *import1.RecoveryPoint, args ...map[string]interface{}) (*import1.CreateRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRecoveryPoint(context.Background(), &import10.CreateRecoveryPointRequest{
		Body: body,
	}, args...)
}

// Create a new recovery point.<br> #### Task Completion Details <br> External identifier of the created recovery point can be found in the task completion details under the key `recoveryPointExtId`.
func (api *RecoveryPointsServiceApi) CreateRecoveryPoint(ctx context.Context, request *import10.CreateRecoveryPointRequest, args ...map[string]interface{}) (*import1.CreateRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points"

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

	unmarshalledResp := new(import1.CreateRecoveryPointApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the recovery point identified by {extId}.
func (api *RecoveryPointsApi) DeleteRecoveryPointById(extId *string, args ...map[string]interface{}) (*import1.DeleteRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRecoveryPointById(context.Background(), &import10.DeleteRecoveryPointByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete the recovery point identified by {extId}.
func (api *RecoveryPointsServiceApi) DeleteRecoveryPointById(ctx context.Context, request *import10.DeleteRecoveryPointByIdRequest, args ...map[string]interface{}) (*import1.DeleteRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}"

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

	unmarshalledResp := new(import1.DeleteRecoveryPointApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// It returns cluster details where the given recovery point is located, and a certificate to access the endpoint. The certificate must be set as a NTNX_IGW_SESSION cookie in the header. For example, Cookie: NTNX_IGW_SESSION='certificate'
func (api *RecoveryPointsApi) DiscoverClusterForRecoveryPointId(extId *string, body *import7.ClusterDiscoverSpec, args ...map[string]interface{}) (*import1.ClusterInfoApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DiscoverClusterForRecoveryPointId(context.Background(), &import10.DiscoverClusterForRecoveryPointIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// It returns cluster details where the given recovery point is located, and a certificate to access the endpoint. The certificate must be set as a NTNX_IGW_SESSION cookie in the header. For example, Cookie: NTNX_IGW_SESSION='certificate'
func (api *RecoveryPointsServiceApi) DiscoverClusterForRecoveryPointId(ctx context.Context, request *import10.DiscoverClusterForRecoveryPointIdRequest, args ...map[string]interface{}) (*import1.ClusterInfoApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}/$actions/discover-cluster"

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

	unmarshalledResp := new(import1.ClusterInfoApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the recovery point identified by {extId}.
func (api *RecoveryPointsApi) GetRecoveryPointById(extId *string, args ...map[string]interface{}) (*import1.GetRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecoveryPointById(context.Background(), &import10.GetRecoveryPointByIdRequest{
		ExtId: extId,
	}, args...)
}

// Query the recovery point identified by {extId}.
func (api *RecoveryPointsServiceApi) GetRecoveryPointById(ctx context.Context, request *import10.GetRecoveryPointByIdRequest, args ...map[string]interface{}) (*import1.GetRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}"

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

	unmarshalledResp := new(import1.GetRecoveryPointApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Query the VM recovery point identified by {extId}.
func (api *RecoveryPointsApi) GetVmRecoveryPointById(recoveryPointExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetVmRecoveryPointApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVmRecoveryPointById(context.Background(), &import10.GetVmRecoveryPointByIdRequest{
		RecoveryPointExtId: recoveryPointExtId,
		ExtId:              extId,
	}, args...)
}

// Query the VM recovery point identified by {extId}.
func (api *RecoveryPointsServiceApi) GetVmRecoveryPointById(ctx context.Context, request *import10.GetVmRecoveryPointByIdRequest, args ...map[string]interface{}) (*import1.GetVmRecoveryPointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{extId}"

	// verify the required parameter 'recoveryPointExtId' is set
	if nil == request.RecoveryPointExtId {
		return nil, client.ReportError("recoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPointExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetVmRecoveryPointApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// The metadata documents of Volume Shadow Copy Service (VSS) writers and requesters are called VSS metadata. During a VSS backup operation, the VSS metadata is compressed into a cabinet file, which is in a .cab file format designed to store compressed files. This cabinet file must be saved to the backup media during a backup operation, as it is required during a restore operation. This API returns the VSS metadata (cabinet file) of a VM recovery point under a composite recovery point that is identified by an external identifier. This external identifier was saved during the recovery point creation operation.
func (api *RecoveryPointsApi) GetVssMetadataByVmRecoveryPointId(recoveryPointExtId *string, vmRecoveryPointExtId *string, args ...map[string]interface{}) (*import7.GetVssMetadataApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVssMetadataByVmRecoveryPointId(context.Background(), &import10.GetVssMetadataByVmRecoveryPointIdRequest{
		RecoveryPointExtId:   recoveryPointExtId,
		VmRecoveryPointExtId: vmRecoveryPointExtId,
	}, args...)
}

// The metadata documents of Volume Shadow Copy Service (VSS) writers and requesters are called VSS metadata. During a VSS backup operation, the VSS metadata is compressed into a cabinet file, which is in a .cab file format designed to store compressed files. This cabinet file must be saved to the backup media during a backup operation, as it is required during a restore operation. This API returns the VSS metadata (cabinet file) of a VM recovery point under a composite recovery point that is identified by an external identifier. This external identifier was saved during the recovery point creation operation.
func (api *RecoveryPointsServiceApi) GetVssMetadataByVmRecoveryPointId(ctx context.Context, request *import10.GetVssMetadataByVmRecoveryPointIdRequest, args ...map[string]interface{}) (*import7.GetVssMetadataApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/content/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/vss-metadata"

	// verify the required parameter 'recoveryPointExtId' is set
	if nil == request.RecoveryPointExtId {
		return nil, client.ReportError("recoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'vmRecoveryPointExtId' is set
	if nil == request.VmRecoveryPointExtId {
		return nil, client.ReportError("vmRecoveryPointExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPointExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"vmRecoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmRecoveryPointExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/octet-stream", "application/json"}

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

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import7.NewGetVssMetadataApiResponse()
			fileDetail := import7.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import8.Flag
			flags = append(flags, import8.Flag{Name: &flagName, Value: &flagValue})
			metadata := import9.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import7.GetVssMetadataApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List of recovery points.
func (api *RecoveryPointsApi) ListRecoveryPoints(xClusterId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRecoveryPointsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecoveryPoints(context.Background(), &import10.ListRecoveryPointsRequest{
		XClusterId: xClusterId,
		Page_:      page_,
		Limit_:     limit_,
		Filter_:    filter_,
		Orderby_:   orderby_,
		Select_:    select_,
	}, args...)
}

// List of recovery points.
func (api *RecoveryPointsServiceApi) ListRecoveryPoints(ctx context.Context, request *import10.ListRecoveryPointsRequest, args ...map[string]interface{}) (*import1.ListRecoveryPointsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points"

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

	unmarshalledResp := new(import1.ListRecoveryPointsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Replicate the recovery point identified by {extId}.<br> #### Task Completion Details <br> External identifier of the replicated recovery point can be found in the task completion details under the key `recoveryPointExtId`.
func (api *RecoveryPointsApi) ReplicateRecoveryPoint(extId *string, body *import1.RecoveryPointReplicationSpec, args ...map[string]interface{}) (*import1.RecoveryPointReplicateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReplicateRecoveryPoint(context.Background(), &import10.ReplicateRecoveryPointRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Replicate the recovery point identified by {extId}.<br> #### Task Completion Details <br> External identifier of the replicated recovery point can be found in the task completion details under the key `recoveryPointExtId`.
func (api *RecoveryPointsServiceApi) ReplicateRecoveryPoint(ctx context.Context, request *import10.ReplicateRecoveryPointRequest, args ...map[string]interface{}) (*import1.RecoveryPointReplicateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}/$actions/replicate"

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

	unmarshalledResp := new(import1.RecoveryPointReplicateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Restore a recovery point identified by {extId}.<br> #### Task Completion Details <br> A comma separated list of the created VM and volume group external identifiers can be found in the task completion details under the keys `vmExtIds` and `volumeGroupExtIds` respectively.
func (api *RecoveryPointsApi) RestoreRecoveryPoint(extId *string, body *import1.RecoveryPointRestorationSpec, args ...map[string]interface{}) (*import1.RecoveryPointRestoreApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RestoreRecoveryPoint(context.Background(), &import10.RestoreRecoveryPointRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Restore a recovery point identified by {extId}.<br> #### Task Completion Details <br> A comma separated list of the created VM and volume group external identifiers can be found in the task completion details under the keys `vmExtIds` and `volumeGroupExtIds` respectively.
func (api *RecoveryPointsServiceApi) RestoreRecoveryPoint(ctx context.Context, request *import10.RestoreRecoveryPointRequest, args ...map[string]interface{}) (*import1.RecoveryPointRestoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}/$actions/restore"

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

	unmarshalledResp := new(import1.RecoveryPointRestoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Set the expiration time for the recovery point identified by {extId}.
func (api *RecoveryPointsApi) SetRecoveryPointExpirationTime(extId *string, body *import1.ExpirationTimeSpec, args ...map[string]interface{}) (*import1.UpdateRecoveryPointExpirationTimeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SetRecoveryPointExpirationTime(context.Background(), &import10.SetRecoveryPointExpirationTimeRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Set the expiration time for the recovery point identified by {extId}.
func (api *RecoveryPointsServiceApi) SetRecoveryPointExpirationTime(ctx context.Context, request *import10.SetRecoveryPointExpirationTimeRequest, args ...map[string]interface{}) (*import1.UpdateRecoveryPointExpirationTimeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/recovery-points/{extId}/$actions/set-expiration-time"

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

	unmarshalledResp := new(import1.UpdateRecoveryPointExpirationTimeApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Displays the calculated metadata with the changed region details between any two VM disk recovery points of a file. This API can be used for incremental and differential backups, as well as for a full backup, because it indicates the regions that are zeros, helping to avoid copying zero regions.
func (api *RecoveryPointsApi) VmRecoveryPointComputeChangedRegions(recoveryPointExtId *string, vmRecoveryPointExtId *string, extId *string, body *import7.VmRecoveryPointChangedRegionsComputeSpec, args ...map[string]interface{}) (*import7.ChangedVmRegionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.VmRecoveryPointComputeChangedRegions(context.Background(), &import10.VmRecoveryPointComputeChangedRegionsRequest{
		RecoveryPointExtId:   recoveryPointExtId,
		VmRecoveryPointExtId: vmRecoveryPointExtId,
		ExtId:                extId,
		Body:                 body,
	}, args...)
}

// Displays the calculated metadata with the changed region details between any two VM disk recovery points of a file. This API can be used for incremental and differential backups, as well as for a full backup, because it indicates the regions that are zeros, helping to avoid copying zero regions.
func (api *RecoveryPointsServiceApi) VmRecoveryPointComputeChangedRegions(ctx context.Context, request *import10.VmRecoveryPointComputeChangedRegionsRequest, args ...map[string]interface{}) (*import7.ChangedVmRegionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/content/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/disk-recovery-points/{extId}/$actions/compute-changed-regions"

	// verify the required parameter 'recoveryPointExtId' is set
	if nil == request.RecoveryPointExtId {
		return nil, client.ReportError("recoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'vmRecoveryPointExtId' is set
	if nil == request.VmRecoveryPointExtId {
		return nil, client.ReportError("vmRecoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPointExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"vmRecoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.VmRecoveryPointExtId, "")), -1)
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

	unmarshalledResp := new(import7.ChangedVmRegionsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Displays the calculated metadata with the changed region details between any two volume group disk recovery points of a file. This API can be used for incremental and differential backups, as well as for a full backup, because it indicates the regions that are zeros, helping to avoid copying zero regions.
func (api *RecoveryPointsApi) VolumeGroupRecoveryPointComputeChangedRegions(recoveryPointExtId *string, volumeGroupRecoveryPointExtId *string, extId *string, body *import7.VolumeGroupRecoveryPointChangedRegionsComputeSpec, args ...map[string]interface{}) (*import7.ChangedVolumeGroupRegionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecoveryPointsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.VolumeGroupRecoveryPointComputeChangedRegions(context.Background(), &import10.VolumeGroupRecoveryPointComputeChangedRegionsRequest{
		RecoveryPointExtId:            recoveryPointExtId,
		VolumeGroupRecoveryPointExtId: volumeGroupRecoveryPointExtId,
		ExtId:                         extId,
		Body:                          body,
	}, args...)
}

// Displays the calculated metadata with the changed region details between any two volume group disk recovery points of a file. This API can be used for incremental and differential backups, as well as for a full backup, because it indicates the regions that are zeros, helping to avoid copying zero regions.
func (api *RecoveryPointsServiceApi) VolumeGroupRecoveryPointComputeChangedRegions(ctx context.Context, request *import10.VolumeGroupRecoveryPointComputeChangedRegionsRequest, args ...map[string]interface{}) (*import7.ChangedVolumeGroupRegionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/content/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/disk-recovery-points/{extId}/$actions/compute-changed-regions"

	// verify the required parameter 'recoveryPointExtId' is set
	if nil == request.RecoveryPointExtId {
		return nil, client.ReportError("recoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'volumeGroupRecoveryPointExtId' is set
	if nil == request.VolumeGroupRecoveryPointExtId {
		return nil, client.ReportError("volumeGroupRecoveryPointExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"recoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.RecoveryPointExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"volumeGroupRecoveryPointExtId"+"}", url.PathEscape(client.ParameterToString(*request.VolumeGroupRecoveryPointExtId, "")), -1)
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

	unmarshalledResp := new(import7.ChangedVolumeGroupRegionsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
