package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/content"
	import8 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/reports"
	"net/http"
	"net/url"
	"strings"
)

type ReportsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ReportsServiceApi
}

type ReportsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewReportsApi(apiClient *client.ApiClient) *ReportsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewReportsServiceApi(a.ApiClient)

	return a
}

func NewReportsServiceApi(apiClient *client.ApiClient) *ReportsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ReportsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a report for the specified report configuration UUID.
func (api *ReportsApi) CreateReport(body *import1.Report, args ...map[string]interface{}) (*import1.CreateReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateReport(context.Background(), &import8.CreateReportRequest{
		Body: body,
	}, args...)
}

// Creates a report for the specified report configuration UUID.
func (api *ReportsServiceApi) CreateReport(ctx context.Context, request *import8.CreateReportRequest, args ...map[string]interface{}) (*import1.CreateReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/reports"

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

	unmarshalledResp := new(import1.CreateReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the report associated with the given UUID.
func (api *ReportsApi) DeleteReportById(extId *string, args ...map[string]interface{}) (*import1.DeleteReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteReportById(context.Background(), &import8.DeleteReportByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the report associated with the given UUID.
func (api *ReportsServiceApi) DeleteReportById(ctx context.Context, request *import8.DeleteReportByIdRequest, args ...map[string]interface{}) (*import1.DeleteReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/reports/{extId}"

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

	unmarshalledResp := new(import1.DeleteReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Downloads a generated report by utilizing the given report type and UUID.
func (api *ReportsApi) DownloadReport(reportExtId *string, args ...map[string]interface{}) (*import3.DownloadReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DownloadReport(context.Background(), &import8.DownloadReportRequest{
		ReportExtId: reportExtId,
	}, args...)
}

// Downloads a generated report by utilizing the given report type and UUID.
func (api *ReportsServiceApi) DownloadReport(ctx context.Context, request *import8.DownloadReportRequest, args ...map[string]interface{}) (*import3.DownloadReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/content/reports/{reportExtId}/file"

	// verify the required parameter 'reportExtId' is set
	if nil == request.ReportExtId {
		return nil, client.ReportError("reportExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"reportExtId"+"}", url.PathEscape(client.ParameterToString(*request.ReportExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/pdf", "application/zip", "application/json"}

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

			response := import3.NewDownloadReportApiResponse()
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

	unmarshalledResp := new(import3.DownloadReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the report associated with the given UUID.
func (api *ReportsApi) GetReportById(extId *string, args ...map[string]interface{}) (*import1.GetReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetReportById(context.Background(), &import8.GetReportByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the report associated with the given UUID.
func (api *ReportsServiceApi) GetReportById(ctx context.Context, request *import8.GetReportByIdRequest, args ...map[string]interface{}) (*import1.GetReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/reports/{extId}"

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

	unmarshalledResp := new(import1.GetReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of all the generated reports
func (api *ReportsApi) ListReports(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListReportsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListReports(context.Background(), &import8.ListReportsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of all the generated reports
func (api *ReportsServiceApi) ListReports(ctx context.Context, request *import8.ListReportsRequest, args ...map[string]interface{}) (*import1.ListReportsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/reports"

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

	unmarshalledResp := new(import1.ListReportsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Notifies the specified recipients by emailing the generated reports for the given UUID.
func (api *ReportsApi) NotifyReport(extId *string, body *import1.ReportNotificationSpec, args ...map[string]interface{}) (*import1.NotifyReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewReportsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.NotifyReport(context.Background(), &import8.NotifyReportRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Notifies the specified recipients by emailing the generated reports for the given UUID.
func (api *ReportsServiceApi) NotifyReport(ctx context.Context, request *import8.NotifyReportRequest, args ...map[string]interface{}) (*import1.NotifyReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.0/config/reports/{extId}/$actions/notify-recipients"

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

	unmarshalledResp := new(import1.NotifyReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
