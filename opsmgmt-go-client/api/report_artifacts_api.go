package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/content"
	import6 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/reportartifacts"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ReportArtifactsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ReportArtifactsServiceApi
}

type ReportArtifactsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewReportArtifactsApi(apiClient *client.ApiClient) *ReportArtifactsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportArtifactsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewReportArtifactsServiceApi(a.ApiClient)

	return a
}

func NewReportArtifactsServiceApi(apiClient *client.ApiClient) *ReportArtifactsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportArtifactsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// This operation creates a report artifact using the provided artifact type and file type.
func (api *ReportArtifactsApi) CreateReportArtifact(body *import3.ReportArtifact, args ...map[string]interface{}) (*import3.CreateReportArtifactApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportArtifactsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateReportArtifact(context.Background(), &import6.CreateReportArtifactRequest{
		Body: body,
	}, args...)
}

// This operation creates a report artifact using the provided artifact type and file type.
func (api *ReportArtifactsServiceApi) CreateReportArtifact(ctx context.Context, request *import6.CreateReportArtifactRequest, args ...map[string]interface{}) (*import3.CreateReportArtifactApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/content/report-artifacts"

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

	unmarshalledResp := new(import3.CreateReportArtifactApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This operation downloads the artifact with the given UUID.
func (api *ReportArtifactsApi) DownloadArtifactFile(reportArtifactExtId *string, args ...map[string]interface{}) (*import3.DownloadArtifactileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportArtifactsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DownloadArtifactFile(context.Background(), &import6.DownloadArtifactFileRequest{
		ReportArtifactExtId: reportArtifactExtId,
	}, args...)
}

// This operation downloads the artifact with the given UUID.
func (api *ReportArtifactsServiceApi) DownloadArtifactFile(ctx context.Context, request *import6.DownloadArtifactFileRequest, args ...map[string]interface{}) (*import3.DownloadArtifactileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/content/report-artifacts/{reportArtifactExtId}/file"

	// verify the required parameter 'reportArtifactExtId' is set
	if nil == request.ReportArtifactExtId {
		return nil, client.ReportError("reportArtifactExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"reportArtifactExtId"+"}", url.PathEscape(client.ParameterToString(*request.ReportArtifactExtId, "")), -1)
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

			response := import3.NewDownloadArtifactileApiResponse()
			fileDetail := import3.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import4.Flag
			flags = append(flags, import4.Flag{Name: &flagName, Value: &flagValue})
			metadata := import5.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import3.DownloadArtifactileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Endpoint for listing all report artifacts accessible to the user matching the filter criteria.
func (api *ReportArtifactsApi) ListReportArtifacts(page_ *int, limit_ *int, filter_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListReportArtifactsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportArtifactsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListReportArtifacts(context.Background(), &import6.ListReportArtifactsRequest{
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
		Select_: select_,
	}, args...)
}

// Endpoint for listing all report artifacts accessible to the user matching the filter criteria.
func (api *ReportArtifactsServiceApi) ListReportArtifacts(ctx context.Context, request *import6.ListReportArtifactsRequest, args ...map[string]interface{}) (*import3.ListReportArtifactsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/content/report-artifacts"

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

	unmarshalledResp := new(import3.ListReportArtifactsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This operation uploads an artifact file to the provided UUID.
func (api *ReportArtifactsApi) UploadArtifactFile(reportArtifactExtId *string, path *string, args ...map[string]interface{}) (*import3.UploadArtifactApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportArtifactsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UploadArtifactFile(context.Background(), &import6.UploadArtifactFileRequest{
		ReportArtifactExtId: reportArtifactExtId,
		Path:                path,
	}, args...)
}

// This operation uploads an artifact file to the provided UUID.
func (api *ReportArtifactsServiceApi) UploadArtifactFile(ctx context.Context, request *import6.UploadArtifactFileRequest, args ...map[string]interface{}) (*import3.UploadArtifactApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/content/report-artifacts/{reportArtifactExtId}/$actions/upload"

	// verify the required parameter 'reportArtifactExtId' is set
	if nil == request.ReportArtifactExtId {
		return nil, client.ReportError("reportArtifactExtId is required and must be specified")
	}
	// verify the required parameter 'path' is set
	if nil == request.Path {
		return nil, client.ReportError("path is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"reportArtifactExtId"+"}", url.PathEscape(client.ParameterToString(*request.ReportArtifactExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/octet-stream"}

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

	file, err := os.Open(*request.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	headerParams["Content-Length"] = fmt.Sprintf("%d", fileInfo.Size())
	if headerParams["Content-Disposition"] == "" {
		headerParams["Content-Disposition"] = fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name())
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, file, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UploadArtifactApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
