package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/request/dashboardgeoconfiguration"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type DashboardGeoconfigurationApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DashboardGeoconfigurationServiceApi
}

type DashboardGeoconfigurationServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDashboardGeoconfigurationApi(apiClient *client.ApiClient) *DashboardGeoconfigurationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardGeoconfigurationApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDashboardGeoconfigurationServiceApi(a.ApiClient)

	return a
}

func NewDashboardGeoconfigurationServiceApi(apiClient *client.ApiClient) *DashboardGeoconfigurationServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DashboardGeoconfigurationServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Upload a dashboard geoconfiguration by providing the file as part of the request body.
func (api *DashboardGeoconfigurationApi) UploadDashboardGeoconfiguration(path *string, args ...map[string]interface{}) (*import1.UploadDashboardGeoconfigurationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDashboardGeoconfigurationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UploadDashboardGeoconfiguration(context.Background(), &import3.UploadDashboardGeoconfigurationRequest{
		Path: path,
	}, args...)
}

// Upload a dashboard geoconfiguration by providing the file as part of the request body.
func (api *DashboardGeoconfigurationServiceApi) UploadDashboardGeoconfiguration(ctx context.Context, request *import3.UploadDashboardGeoconfigurationRequest, args ...map[string]interface{}) (*import1.UploadDashboardGeoconfigurationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/opsmgmt/v4.1.b1/config/dashboard-settings/$actions/upload-geoconfigurations"

	// verify the required parameter 'path' is set
	if nil == request.Path {
		return nil, client.ReportError("path is required and must be specified")
	}

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.UploadDashboardGeoconfigurationApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
