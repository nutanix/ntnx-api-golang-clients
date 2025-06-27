package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
	"net/http"
	"net/url"
	"strings"
)

type DomainManagerApi struct {
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

	return a
}

// Deploys a Prism Central using the provided details. Prism Central Size, Network Config are mandatory fields to deploy Prism Central. The response from this endpoint contains the URL in the task object location header that can be used to track the request status.
func (api *DomainManagerApi) CreateDomainManager(body *import2.DomainManager, args ...map[string]interface{}) (*import2.CreateDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/config/domain-managers"

	// verify the required parameter 'body' is set
	if nil == body {
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.CreateDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the attributes associated with the domain manager (Prism Central) resource based on the provided external identifier. It includes attributes like config, network, node and other information such as size, environment and resource specifications.
func (api *DomainManagerApi) GetDomainManagerById(extId *string, args ...map[string]interface{}) (*import2.GetDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/config/domain-managers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.GetDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the product details along with it's current enablement and resize status.
func (api *DomainManagerApi) GetProductById(domainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetProductApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/products/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.GetProductApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of elements representing the domain manager (Prism Central) instance.
func (api *DomainManagerApi) ListDomainManagers(select_ *string, args ...map[string]interface{}) (*import2.ListDomainManagerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/config/domain-managers"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import2.ListDomainManagerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves a list of all products along with their current enablement and resize status.
func (api *DomainManagerApi) ListProducts(domainManagerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListProductsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/products"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.ListProductsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Registers a domain manager (Prism Central) instance to other entities like PE and PC. This process is asynchronous, creating a registration task and returning its UUID.
func (api *DomainManagerApi) Register(extId *string, body *import3.ClusterRegistrationSpec, dryrun_ *bool, args ...map[string]interface{}) (*import3.RegisterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{extId}/$actions/register"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*dryrun_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.RegisterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Unregister a registered remote cluster from the local cluster. This process is asynchronous, creating an unregisteration task and returning its UUID.
func (api *DomainManagerApi) Unregister(extId *string, body *import3.ClusterReference, dryrun_ *bool, args ...map[string]interface{}) (*import3.UnregisterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{extId}/$actions/unregister"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*dryrun_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UnregisterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the status of a given product with the current support focused on updating the enablement state.
func (api *DomainManagerApi) UpdateProductById(domainManagerExtId *string, extId *string, body *import3.Product, args ...map[string]interface{}) (*import3.UpdateProductApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/products/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UpdateProductApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
