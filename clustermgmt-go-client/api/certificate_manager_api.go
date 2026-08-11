package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/certificatemanager"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
	"net/http"
	"net/url"
	"strings"
)

type CertificateManagerApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *CertificateManagerServiceApi
}

type CertificateManagerServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCertificateManagerApi(apiClient *client.ApiClient) *CertificateManagerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CertificateManagerApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewCertificateManagerServiceApi(a.ApiClient)

	return a
}

func NewCertificateManagerServiceApi(apiClient *client.ApiClient) *CertificateManagerServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CertificateManagerServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Add a Certificate Authority (CA) to the specified cluster.
func (api *CertificateManagerApi) CreateCertificateAuthorityByClusterId(clusterExtId *string, body *import1.CertificateAuthority, args ...map[string]interface{}) (*import1.CreateCertificateAuthorityApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCertificateAuthorityByClusterId(context.Background(), &import5.CreateCertificateAuthorityByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Add a Certificate Authority (CA) to the specified cluster.
func (api *CertificateManagerServiceApi) CreateCertificateAuthorityByClusterId(ctx context.Context, request *import5.CreateCertificateAuthorityByClusterIdRequest, args ...map[string]interface{}) (*import1.CreateCertificateAuthorityApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificate-authorities"

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
	unmarshalledResp := new(import1.CreateCertificateAuthorityApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Add a Certificate to the specified cluster.
func (api *CertificateManagerApi) CreateCertificateByClusterId(clusterExtId *string, body *import1.Certificate, dryrun_ *bool, args ...map[string]interface{}) (*import1.CreateCertificateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCertificateByClusterId(context.Background(), &import5.CreateCertificateByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
		Dryrun_:      dryrun_,
	}, args...)
}

// Add a Certificate to the specified cluster.
func (api *CertificateManagerServiceApi) CreateCertificateByClusterId(ctx context.Context, request *import5.CreateCertificateByClusterIdRequest, args ...map[string]interface{}) (*import1.CreateCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificates"

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

	// Query Params
	if request.Dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*request.Dryrun_, ""))
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
	unmarshalledResp := new(import1.CreateCertificateApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Generate a new Certificate Signing Request (CSR) for the specified cluster.
func (api *CertificateManagerApi) CreateCsrByClusterId(clusterExtId *string, body *import1.Csr, args ...map[string]interface{}) (*import1.CreateCsrApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCsrByClusterId(context.Background(), &import5.CreateCsrByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Generate a new Certificate Signing Request (CSR) for the specified cluster.
func (api *CertificateManagerServiceApi) CreateCsrByClusterId(ctx context.Context, request *import5.CreateCsrByClusterIdRequest, args ...map[string]interface{}) (*import1.CreateCsrApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/csrs"

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
	unmarshalledResp := new(import1.CreateCsrApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete a specific Certificate Authority (CA) from the specified cluster.
func (api *CertificateManagerApi) DeleteCertificateAuthorityById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteCertificateAuthorityApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCertificateAuthorityById(context.Background(), &import5.DeleteCertificateAuthorityByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Delete a specific Certificate Authority (CA) from the specified cluster.
func (api *CertificateManagerServiceApi) DeleteCertificateAuthorityById(ctx context.Context, request *import5.DeleteCertificateAuthorityByIdRequest, args ...map[string]interface{}) (*import1.DeleteCertificateAuthorityApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificate-authorities/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.DeleteCertificateAuthorityApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete a specific Certificate Signing Request (CSR) from the specified cluster.
func (api *CertificateManagerApi) DeleteCsrById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteCsrApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCsrById(context.Background(), &import5.DeleteCsrByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Delete a specific Certificate Signing Request (CSR) from the specified cluster.
func (api *CertificateManagerServiceApi) DeleteCsrById(ctx context.Context, request *import5.DeleteCsrByIdRequest, args ...map[string]interface{}) (*import1.DeleteCsrApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/csrs/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.DeleteCsrApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Download a specific Certificate Signing Request (CSR) from the specified cluster.
func (api *CertificateManagerApi) DownloadCsrById(clusterExtId *string, csrExtId *string, args ...map[string]interface{}) (*import1.DownloadCsrApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DownloadCsrById(context.Background(), &import5.DownloadCsrByIdRequest{
		ClusterExtId: clusterExtId,
		CsrExtId:     csrExtId,
	}, args...)
}

// Download a specific Certificate Signing Request (CSR) from the specified cluster.
func (api *CertificateManagerServiceApi) DownloadCsrById(ctx context.Context, request *import5.DownloadCsrByIdRequest, args ...map[string]interface{}) (*import1.DownloadCsrApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/csrs/{csrExtId}/file"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'csrExtId' is set
	if nil == request.CsrExtId {
		return nil, client.ReportError("csrExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"csrExtId"+"}", url.PathEscape(client.ParameterToString(*request.CsrExtId, "")), -1)
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

			response := import1.NewDownloadCsrApiResponse()
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
	unmarshalledResp := new(import1.DownloadCsrApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get details of a specific Certificate Authority (CA) on a particular cluster.
func (api *CertificateManagerApi) GetCertificateAuthorityById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetCertificateAuthorityApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCertificateAuthorityById(context.Background(), &import5.GetCertificateAuthorityByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get details of a specific Certificate Authority (CA) on a particular cluster.
func (api *CertificateManagerServiceApi) GetCertificateAuthorityById(ctx context.Context, request *import5.GetCertificateAuthorityByIdRequest, args ...map[string]interface{}) (*import1.GetCertificateAuthorityApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificate-authorities/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetCertificateAuthorityApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get details of a specific Certificate on a particular cluster.
func (api *CertificateManagerApi) GetCertificateById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetCertificateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCertificateById(context.Background(), &import5.GetCertificateByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get details of a specific Certificate on a particular cluster.
func (api *CertificateManagerServiceApi) GetCertificateById(ctx context.Context, request *import5.GetCertificateByIdRequest, args ...map[string]interface{}) (*import1.GetCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificates/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetCertificateApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get details of a specific Certificate Signing Request (CSR) on a particular cluster.
func (api *CertificateManagerApi) GetCsrById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetCsrApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCsrById(context.Background(), &import5.GetCsrByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get details of a specific Certificate Signing Request (CSR) on a particular cluster.
func (api *CertificateManagerServiceApi) GetCsrById(ctx context.Context, request *import5.GetCsrByIdRequest, args ...map[string]interface{}) (*import1.GetCsrApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/csrs/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetCsrApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List of active Certificate Authorities on a particular cluster.
func (api *CertificateManagerApi) ListCertificateAuthoritiesByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCertificateAuthoritiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCertificateAuthoritiesByClusterId(context.Background(), &import5.ListCertificateAuthoritiesByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// List of active Certificate Authorities on a particular cluster.
func (api *CertificateManagerServiceApi) ListCertificateAuthoritiesByClusterId(ctx context.Context, request *import5.ListCertificateAuthoritiesByClusterIdRequest, args ...map[string]interface{}) (*import1.ListCertificateAuthoritiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificate-authorities"

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
	unmarshalledResp := new(import1.ListCertificateAuthoritiesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List of active Certificates on a particular cluster.
func (api *CertificateManagerApi) ListCertificatesByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCertificatesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCertificatesByClusterId(context.Background(), &import5.ListCertificatesByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// List of active Certificates on a particular cluster.
func (api *CertificateManagerServiceApi) ListCertificatesByClusterId(ctx context.Context, request *import5.ListCertificatesByClusterIdRequest, args ...map[string]interface{}) (*import1.ListCertificatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/certificates"

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
	unmarshalledResp := new(import1.ListCertificatesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List of active Certificate Signing Requests on a particular cluster.
func (api *CertificateManagerApi) ListCsrsByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCsrsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCsrsByClusterId(context.Background(), &import5.ListCsrsByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// List of active Certificate Signing Requests on a particular cluster.
func (api *CertificateManagerServiceApi) ListCsrsByClusterId(ctx context.Context, request *import5.ListCsrsByClusterIdRequest, args ...map[string]interface{}) (*import1.ListCsrsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/csrs"

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
	unmarshalledResp := new(import1.ListCsrsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
