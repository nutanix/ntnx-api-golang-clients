package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
	import6 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/domainmanager"
	"net/http"
	"net/url"
	"strings"
)

type DomainManagerApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DomainManagerServiceApi
}

type DomainManagerServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDomainManagerApi(apiClient *client.ApiClient) *DomainManagerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DomainManagerApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDomainManagerServiceApi(a.ApiClient)

	return a
}

func NewDomainManagerServiceApi(apiClient *client.ApiClient) *DomainManagerServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DomainManagerServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Deploys a Prism Central using the provided details. Prism Central Size, Network Config are mandatory fields to deploy Prism Central. The response from this endpoint contains the URL in the task object location header that can be used to track the request status.
func (api *DomainManagerApi) CreateDomainManager(body *import3.DomainManager, args ...map[string]interface{}) (*import3.CreateDomainManagerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDomainManager(context.Background(), &import6.CreateDomainManagerRequest{
		Body: body,
	}, args...)
}

// Deploys a Prism Central using the provided details. Prism Central Size, Network Config are mandatory fields to deploy Prism Central. The response from this endpoint contains the URL in the task object location header that can be used to track the request status.
func (api *DomainManagerServiceApi) CreateDomainManager(ctx context.Context, request *import6.CreateDomainManagerRequest, args ...map[string]interface{}) (*import3.CreateDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/domain-managers"

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

	unmarshalledResp := new(import3.CreateDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the attributes associated with the domain manager (Prism Central) resource based on the provided external identifier. It includes attributes like config, network, node and other information such as size, environment and resource specifications.
func (api *DomainManagerApi) GetDomainManagerById(extId *string, args ...map[string]interface{}) (*import3.GetDomainManagerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDomainManagerById(context.Background(), &import6.GetDomainManagerByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the attributes associated with the domain manager (Prism Central) resource based on the provided external identifier. It includes attributes like config, network, node and other information such as size, environment and resource specifications.
func (api *DomainManagerServiceApi) GetDomainManagerById(ctx context.Context, request *import6.GetDomainManagerByIdRequest, args ...map[string]interface{}) (*import3.GetDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/domain-managers/{extId}"

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

	unmarshalledResp := new(import3.GetDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the product details along with it's current enablement and resize status.
func (api *DomainManagerApi) GetProductById(domainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import5.GetProductApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetProductById(context.Background(), &import6.GetProductByIdRequest{
		DomainManagerExtId: domainManagerExtId,
		ExtId:              extId,
	}, args...)
}

// Retrieves the product details along with it's current enablement and resize status.
func (api *DomainManagerServiceApi) GetProductById(ctx context.Context, request *import6.GetProductByIdRequest, args ...map[string]interface{}) (*import5.GetProductApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{domainManagerExtId}/products/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == request.DomainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import5.GetProductApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of elements representing the domain manager (Prism Central) instance.
func (api *DomainManagerApi) ListDomainManagers(select_ *string, args ...map[string]interface{}) (*import3.ListDomainManagerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDomainManagers(context.Background(), &import6.ListDomainManagersRequest{
		Select_: select_,
	}, args...)
}

// Returns a list of elements representing the domain manager (Prism Central) instance.
func (api *DomainManagerServiceApi) ListDomainManagers(ctx context.Context, request *import6.ListDomainManagersRequest, args ...map[string]interface{}) (*import3.ListDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/domain-managers"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
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

	unmarshalledResp := new(import3.ListDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves a list of all products along with their current enablement and resize status.
func (api *DomainManagerApi) ListProducts(domainManagerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import5.ListProductsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListProducts(context.Background(), &import6.ListProductsRequest{
		DomainManagerExtId: domainManagerExtId,
		Page_:              page_,
		Limit_:             limit_,
		Filter_:            filter_,
		Orderby_:           orderby_,
		Select_:            select_,
	}, args...)
}

// Retrieves a list of all products along with their current enablement and resize status.
func (api *DomainManagerServiceApi) ListProducts(ctx context.Context, request *import6.ListProductsRequest, args ...map[string]interface{}) (*import5.ListProductsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{domainManagerExtId}/products"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == request.DomainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import5.ListProductsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Registers a domain manager (Prism Central) instance to other entities like PE and PC. This process is asynchronous, creating a registration task and returning its UUID.
func (api *DomainManagerApi) Register(extId *string, body *import5.ClusterRegistrationSpec, dryrun_ *bool, args ...map[string]interface{}) (*import5.RegisterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.Register(context.Background(), &import6.RegisterRequest{
		ExtId:   extId,
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Registers a domain manager (Prism Central) instance to other entities like PE and PC. This process is asynchronous, creating a registration task and returning its UUID.
func (api *DomainManagerServiceApi) Register(ctx context.Context, request *import6.RegisterRequest, args ...map[string]interface{}) (*import5.RegisterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{extId}/$actions/register"

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

	unmarshalledResp := new(import5.RegisterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Unregister a registered remote cluster from the local cluster. This process is asynchronous, creating an unregisteration task and returning its UUID.
func (api *DomainManagerApi) Unregister(extId *string, body *import5.ClusterReference, dryrun_ *bool, args ...map[string]interface{}) (*import5.UnregisterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.Unregister(context.Background(), &import6.UnregisterRequest{
		ExtId:   extId,
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Unregister a registered remote cluster from the local cluster. This process is asynchronous, creating an unregisteration task and returning its UUID.
func (api *DomainManagerServiceApi) Unregister(ctx context.Context, request *import6.UnregisterRequest, args ...map[string]interface{}) (*import5.UnregisterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{extId}/$actions/unregister"

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

	unmarshalledResp := new(import5.UnregisterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the status of a given product with the current support focused on updating the enablement state.
func (api *DomainManagerApi) UpdateProductById(domainManagerExtId *string, extId *string, body *import5.Product, dryrun_ *bool, args ...map[string]interface{}) (*import5.UpdateProductApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDomainManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateProductById(context.Background(), &import6.UpdateProductByIdRequest{
		DomainManagerExtId: domainManagerExtId,
		ExtId:              extId,
		Body:               body,
		Dryrun_:            dryrun_,
	}, args...)
}

// Updates the status of a given product with the current support focused on updating the enablement state.
func (api *DomainManagerServiceApi) UpdateProductById(ctx context.Context, request *import6.UpdateProductByIdRequest, args ...map[string]interface{}) (*import5.UpdateProductApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{domainManagerExtId}/products/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == request.DomainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainManagerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.UpdateProductApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
