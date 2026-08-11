package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import23 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/upgradeselections"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type UpgradeSelectionsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UpgradeSelectionsServiceApi
}

type UpgradeSelectionsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUpgradeSelectionsApi(apiClient *client.ApiClient) *UpgradeSelectionsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradeSelectionsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUpgradeSelectionsServiceApi(a.ApiClient)

	return a
}

func NewUpgradeSelectionsServiceApi(apiClient *client.ApiClient) *UpgradeSelectionsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradeSelectionsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a new upgrade selection to save the intended set of entity upgrades for a cluster. This is primarily used in dark-site workflows where the user selects entities to upgrade, saves the selection, and then exports it to generate a download helper for obtaining the required bundles from the Nutanix Portal. Submit an UpgradeSelection object containing the selectedUpgrades (entity UUID and target version pairs) and the clusterExtId. The operation is asynchronous and returns a TaskReference. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. After creating a selection, use POST /upgrade-selections/{extId}/$actions/export to generate the download helper package.
func (api *UpgradeSelectionsApi) CreateUpgradeSelection(body *import1.UpgradeSelection, args ...map[string]interface{}) (*import1.CreateUpgradeSelectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradeSelectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateUpgradeSelection(context.Background(), &import23.CreateUpgradeSelectionRequest{
		Body: body,
	}, args...)
}

// Create a new upgrade selection to save the intended set of entity upgrades for a cluster. This is primarily used in dark-site workflows where the user selects entities to upgrade, saves the selection, and then exports it to generate a download helper for obtaining the required bundles from the Nutanix Portal. Submit an UpgradeSelection object containing the selectedUpgrades (entity UUID and target version pairs) and the clusterExtId. The operation is asynchronous and returns a TaskReference. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. After creating a selection, use POST /upgrade-selections/{extId}/$actions/export to generate the download helper package.
func (api *UpgradeSelectionsServiceApi) CreateUpgradeSelection(ctx context.Context, request *import23.CreateUpgradeSelectionRequest, args ...map[string]interface{}) (*import1.CreateUpgradeSelectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/upgrade-selections"

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
	unmarshalledResp := new(import1.CreateUpgradeSelectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete an upgrade selection that is no longer needed. This operation is asynchronous and returns a TaskReference. Deleting a selection does not affect any bundles or images that have already been uploaded to the cluster. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *UpgradeSelectionsApi) DeleteUpgradeSelectionById(extId *string, args ...map[string]interface{}) (*import1.DeleteUpgradeSelectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradeSelectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteUpgradeSelectionById(context.Background(), &import23.DeleteUpgradeSelectionByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete an upgrade selection that is no longer needed. This operation is asynchronous and returns a TaskReference. Deleting a selection does not affect any bundles or images that have already been uploaded to the cluster. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *UpgradeSelectionsServiceApi) DeleteUpgradeSelectionById(ctx context.Context, request *import23.DeleteUpgradeSelectionByIdRequest, args ...map[string]interface{}) (*import1.DeleteUpgradeSelectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/upgrade-selections/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.DeleteUpgradeSelectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Generate a download_helper.zip file containing scripts and instructions for downloading the LCM dark-site bundles required for the selected upgrades from the Nutanix Portal. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the download_helper.zip file is available in the completion_details field of the task. The download helper includes metadata about which bundles are needed and helper scripts for automated download. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *UpgradeSelectionsApi) ExportUpgradeSelection(extId *string, args ...map[string]interface{}) (*import1.ExportUpgradeSelectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradeSelectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportUpgradeSelection(context.Background(), &import23.ExportUpgradeSelectionRequest{
		ExtId: extId,
	}, args...)
}

// Generate a download_helper.zip file containing scripts and instructions for downloading the LCM dark-site bundles required for the selected upgrades from the Nutanix Portal. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, the URL to download the download_helper.zip file is available in the completion_details field of the task. The download helper includes metadata about which bundles are needed and helper scripts for automated download. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *UpgradeSelectionsServiceApi) ExportUpgradeSelection(ctx context.Context, request *import23.ExportUpgradeSelectionRequest, args ...map[string]interface{}) (*import1.ExportUpgradeSelectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/upgrade-selections/{extId}/$actions/export"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.ExportUpgradeSelectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the full details of a specific upgrade selection by its identifier. The response includes the list of selected entity upgrades (entity UUID and target version pairs), the cluster identifier, and the current selection status (PENDING_UPLOAD, UPGRADE_READY, or STALE_SELECTION). Use this endpoint to check whether all required bundles have been uploaded and the selection is ready for upgrade.
func (api *UpgradeSelectionsApi) GetUpgradeSelectionById(extId *string, args ...map[string]interface{}) (*import1.GetUpgradeSelectionApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradeSelectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUpgradeSelectionById(context.Background(), &import23.GetUpgradeSelectionByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the full details of a specific upgrade selection by its identifier. The response includes the list of selected entity upgrades (entity UUID and target version pairs), the cluster identifier, and the current selection status (PENDING_UPLOAD, UPGRADE_READY, or STALE_SELECTION). Use this endpoint to check whether all required bundles have been uploaded and the selection is ready for upgrade.
func (api *UpgradeSelectionsServiceApi) GetUpgradeSelectionById(ctx context.Context, request *import23.GetUpgradeSelectionByIdRequest, args ...map[string]interface{}) (*import1.GetUpgradeSelectionApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/upgrade-selections/{extId}"

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
	unmarshalledResp := new(import1.GetUpgradeSelectionApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a paginated list of all upgrade selections on the cluster. Supports standard query parameters for pagination ($page, $limit) and filtering ($filter). Each selection shows the saved entity upgrades, cluster identifier, and current status.
func (api *UpgradeSelectionsApi) ListUpgradeSelections(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListUpgradeSelectionsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradeSelectionsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUpgradeSelections(context.Background(), &import23.ListUpgradeSelectionsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieve a paginated list of all upgrade selections on the cluster. Supports standard query parameters for pagination ($page, $limit) and filtering ($filter). Each selection shows the saved entity upgrades, cluster identifier, and current status.
func (api *UpgradeSelectionsServiceApi) ListUpgradeSelections(ctx context.Context, request *import23.ListUpgradeSelectionsRequest, args ...map[string]interface{}) (*import1.ListUpgradeSelectionsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/upgrade-selections"

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
	unmarshalledResp := new(import1.ListUpgradeSelectionsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
