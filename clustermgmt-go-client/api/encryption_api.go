package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import13 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/encryption"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
	"net/http"
	"net/url"
	"strings"
)

type EncryptionApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EncryptionServiceApi
}

type EncryptionServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEncryptionApi(apiClient *client.ApiClient) *EncryptionApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EncryptionApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEncryptionServiceApi(a.ApiClient)

	return a
}

func NewEncryptionServiceApi(apiClient *client.ApiClient) *EncryptionServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EncryptionServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Download encryption key backup for the cluster(s). This backup can be used to restore the encryption keys.
func (api *EncryptionApi) DownloadEncryptionKeysBackup(extId *string, args ...map[string]interface{}) (*import1.DownloadEncryptionKeysBackupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEncryptionServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DownloadEncryptionKeysBackup(context.Background(), &import13.DownloadEncryptionKeysBackupRequest{
		ExtId: extId,
	}, args...)
}

// Download encryption key backup for the cluster(s). This backup can be used to restore the encryption keys.
func (api *EncryptionServiceApi) DownloadEncryptionKeysBackup(ctx context.Context, request *import13.DownloadEncryptionKeysBackupRequest, args ...map[string]interface{}) (*import1.DownloadEncryptionKeysBackupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/encryption-keys-backup/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import1.NewDownloadEncryptionKeysBackupApiResponse()
			fileDetail := import1.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import3.Flag
			flags = append(flags, import3.Flag{Name: &flagName, Value: &flagValue})
			metadata := import4.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.DownloadEncryptionKeysBackupApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List the fields in the encryption-config sub-entity of the cluster entity.
func (api *EncryptionApi) GetEncryptionConfig(clusterExtId *string, args ...map[string]interface{}) (*import1.GetEncryptionConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEncryptionServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEncryptionConfig(context.Background(), &import13.GetEncryptionConfigRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// List the fields in the encryption-config sub-entity of the cluster entity.
func (api *EncryptionServiceApi) GetEncryptionConfig(ctx context.Context, request *import13.GetEncryptionConfigRequest, args ...map[string]interface{}) (*import1.GetEncryptionConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/encryption-config"

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
	unmarshalledResp := new(import1.GetEncryptionConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Prepare encryption key backup for the specified cluster(s). This backup can be fetched and used to restore the encryption keys.
func (api *EncryptionApi) PrepareEncryptionKeysBackup(body *import1.EncryptionBackupSpec, args ...map[string]interface{}) (*import1.PrepareEncryptionKeysBackupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEncryptionServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PrepareEncryptionKeysBackup(context.Background(), &import13.PrepareEncryptionKeysBackupRequest{
		Body: body,
	}, args...)
}

// Prepare encryption key backup for the specified cluster(s). This backup can be fetched and used to restore the encryption keys.
func (api *EncryptionServiceApi) PrepareEncryptionKeysBackup(ctx context.Context, request *import13.PrepareEncryptionKeysBackupRequest, args ...map[string]interface{}) (*import1.PrepareEncryptionKeysBackupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/$actions/prepare-encryption-keys-backup"

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
	unmarshalledResp := new(import1.PrepareEncryptionKeysBackupApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Rotate the encryption key on a cluster. This can be performed for SOFTWARE encryption type.
func (api *EncryptionApi) RotateEncryptionKeys(clusterExtId *string, body *import1.EncryptionKeyRotationSpec, args ...map[string]interface{}) (*import1.RotateEncryptionKeysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEncryptionServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RotateEncryptionKeys(context.Background(), &import13.RotateEncryptionKeysRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Rotate the encryption key on a cluster. This can be performed for SOFTWARE encryption type.
func (api *EncryptionServiceApi) RotateEncryptionKeys(ctx context.Context, request *import13.RotateEncryptionKeysRequest, args ...map[string]interface{}) (*import1.RotateEncryptionKeysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/$actions/rotate-encryption-keys"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.RotateEncryptionKeysApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Configure encryption on a cluster.
func (api *EncryptionApi) UpdateEncryptionConfig(clusterExtId *string, body *import1.EncryptionConfig, args ...map[string]interface{}) (*import1.UpdateEncryptionConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEncryptionServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateEncryptionConfig(context.Background(), &import13.UpdateEncryptionConfigRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Configure encryption on a cluster.
func (api *EncryptionServiceApi) UpdateEncryptionConfig(ctx context.Context, request *import13.UpdateEncryptionConfigRequest, args ...map[string]interface{}) (*import1.UpdateEncryptionConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/encryption-config"

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
	unmarshalledResp := new(import1.UpdateEncryptionConfigApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
