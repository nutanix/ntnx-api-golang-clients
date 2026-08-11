package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/bundles"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type BundlesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *BundlesServiceApi
}

type BundlesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewBundlesApi(apiClient *client.ApiClient) *BundlesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BundlesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewBundlesServiceApi(a.ApiClient)

	return a
}

func NewBundlesServiceApi(apiClient *client.ApiClient) *BundlesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BundlesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a new LCM bundle on the cluster. Bundle creation is a 2-step process. Step 1: Stream the bundle bits from the user's local setup to the objects-lite service using the AWS S3 PUT API. Use base64-encoded username and password of a Prism Admin or Super Admin user as credentials. Alternatively, if the bundle is hosted somewhere accessible (e.g., external server or Nutanix Portal), provide a URL source to avoid Step 1; LCM will download the bundle from the URL. Step 2: Finalize the bundle by calling this CreateBundle v4 API with the object key from Step 1 in the name field (or omit name when using URL source). The request body must include the bundle vendor and may include a checksum (for verification). The operation is asynchronous and returns a TaskReference; poll the task to track progress. After a bundle is successfully created, run an inventory (POST /$actions/inventory) to discover the newly available updates from the bundle's images. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *BundlesApi) CreateBundle(body *import1.Bundle, args ...map[string]interface{}) (*import1.CreateBundleApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBundlesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateBundle(context.Background(), &import2.CreateBundleRequest{
		Body: body,
	}, args...)
}

// Create a new LCM bundle on the cluster. Bundle creation is a 2-step process. Step 1: Stream the bundle bits from the user's local setup to the objects-lite service using the AWS S3 PUT API. Use base64-encoded username and password of a Prism Admin or Super Admin user as credentials. Alternatively, if the bundle is hosted somewhere accessible (e.g., external server or Nutanix Portal), provide a URL source to avoid Step 1; LCM will download the bundle from the URL. Step 2: Finalize the bundle by calling this CreateBundle v4 API with the object key from Step 1 in the name field (or omit name when using URL source). The request body must include the bundle vendor and may include a checksum (for verification). The operation is asynchronous and returns a TaskReference; poll the task to track progress. After a bundle is successfully created, run an inventory (POST /$actions/inventory) to discover the newly available updates from the bundle's images. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *BundlesServiceApi) CreateBundle(ctx context.Context, request *import2.CreateBundleRequest, args ...map[string]interface{}) (*import1.CreateBundleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/bundles"

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
	unmarshalledResp := new(import1.CreateBundleApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete an LCM bundle and its associated images from the cluster. This operation is asynchronous and returns a TaskReference; poll the task to confirm deletion completed. Deleting a bundle removes the upgrade images it contains, which may affect the ability to perform upgrades for the corresponding entities until new images are provided. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *BundlesApi) DeleteBundleById(extId *string, args ...map[string]interface{}) (*import1.DeleteBundleByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBundlesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteBundleById(context.Background(), &import2.DeleteBundleByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete an LCM bundle and its associated images from the cluster. This operation is asynchronous and returns a TaskReference; poll the task to confirm deletion completed. Deleting a bundle removes the upgrade images it contains, which may affect the ability to perform upgrades for the corresponding entities until new images are provided. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *BundlesServiceApi) DeleteBundleById(ctx context.Context, request *import2.DeleteBundleByIdRequest, args ...map[string]interface{}) (*import1.DeleteBundleByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/bundles/{extId}"

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
	unmarshalledResp := new(import1.DeleteBundleByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the full details of a specific LCM bundle by its external identifier. The response includes the bundle name, size, type (SOFTWARE, FIRMWARE, PRODUCT_META, FRAMEWORK, or IMAGE_BUNDLE), vendor, creation timestamp, contained images, checksum, and upload mode. Bundles are used primarily in dark-site deployments to deliver upgrade images to a cluster that lacks internet connectivity.
func (api *BundlesApi) GetBundleById(extId *string, args ...map[string]interface{}) (*import1.GetBundleByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBundlesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetBundleById(context.Background(), &import2.GetBundleByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the full details of a specific LCM bundle by its external identifier. The response includes the bundle name, size, type (SOFTWARE, FIRMWARE, PRODUCT_META, FRAMEWORK, or IMAGE_BUNDLE), vendor, creation timestamp, contained images, checksum, and upload mode. Bundles are used primarily in dark-site deployments to deliver upgrade images to a cluster that lacks internet connectivity.
func (api *BundlesServiceApi) GetBundleById(ctx context.Context, request *import2.GetBundleByIdRequest, args ...map[string]interface{}) (*import1.GetBundleByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/bundles/{extId}"

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
	unmarshalledResp := new(import1.GetBundleByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a paginated list of all LCM bundles present on the cluster. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Each bundle entry includes metadata such as name, type, vendor, size, and creation time. For dark-site deployments, bundles are the primary mechanism for delivering upgrade images to the cluster.
func (api *BundlesApi) ListBundles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListBundlesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBundlesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListBundles(context.Background(), &import2.ListBundlesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieve a paginated list of all LCM bundles present on the cluster. Supports standard query parameters for pagination ($page, $limit), filtering ($filter), sorting ($orderby), and field projection ($select). Each bundle entry includes metadata such as name, type, vendor, size, and creation time. For dark-site deployments, bundles are the primary mechanism for delivering upgrade images to the cluster.
func (api *BundlesServiceApi) ListBundles(ctx context.Context, request *import2.ListBundlesRequest, args ...map[string]interface{}) (*import1.ListBundlesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/bundles"

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
	unmarshalledResp := new(import1.ListBundlesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
