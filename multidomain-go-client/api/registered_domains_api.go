package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	import6 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/registereddomains"
	"net/http"
	"net/url"
	"strings"
)

type RegisteredDomainsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RegisteredDomainsServiceApi
}

type RegisteredDomainsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRegisteredDomainsApi(apiClient *client.ApiClient) *RegisteredDomainsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RegisteredDomainsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRegisteredDomainsServiceApi(a.ApiClient)

	return a
}

func NewRegisteredDomainsServiceApi(apiClient *client.ApiClient) *RegisteredDomainsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RegisteredDomainsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new registered domain and saves it to database.
func (api *RegisteredDomainsApi) CreateRegisteredDomain(body *import1.RegisteredDomain, args ...map[string]interface{}) (*import1.CreateRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRegisteredDomain(context.Background(), &import6.CreateRegisteredDomainRequest{
		Body: body,
	}, args...)
}

// Creates a new registered domain and saves it to database.
func (api *RegisteredDomainsServiceApi) CreateRegisteredDomain(ctx context.Context, request *import6.CreateRegisteredDomainRequest, args ...map[string]interface{}) (*import1.CreateRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.3/config/registered-domains"

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

	unmarshalledResp := new(import1.CreateRegisteredDomainApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the registered domain with the requested external identifier.
func (api *RegisteredDomainsApi) DeleteRegisteredDomainById(extId *string, args ...map[string]interface{}) (*import1.DeleteRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRegisteredDomainById(context.Background(), &import6.DeleteRegisteredDomainByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the registered domain with the requested external identifier.
func (api *RegisteredDomainsServiceApi) DeleteRegisteredDomainById(ctx context.Context, request *import6.DeleteRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import1.DeleteRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.3/config/registered-domains/{extId}"

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

	unmarshalledResp := new(import1.DeleteRegisteredDomainApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the registered domain by its external identifier.
func (api *RegisteredDomainsApi) GetRegisteredDomainById(extId *string, select_ *string, args ...map[string]interface{}) (*import1.GetRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRegisteredDomainById(context.Background(), &import6.GetRegisteredDomainByIdRequest{
		ExtId:   extId,
		Select_: select_,
	}, args...)
}

// Retrieves the registered domain by its external identifier.
func (api *RegisteredDomainsServiceApi) GetRegisteredDomainById(ctx context.Context, request *import6.GetRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import1.GetRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.3/config/registered-domains/{extId}"

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

	unmarshalledResp := new(import1.GetRegisteredDomainApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List all registered domains.
func (api *RegisteredDomainsApi) ListRegisteredDomains(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRegisteredDomainsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRegisteredDomains(context.Background(), &import6.ListRegisteredDomainsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// List all registered domains.
func (api *RegisteredDomainsServiceApi) ListRegisteredDomains(ctx context.Context, request *import6.ListRegisteredDomainsRequest, args ...map[string]interface{}) (*import1.ListRegisteredDomainsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.3/config/registered-domains"

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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import1.ListRegisteredDomainsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a registered domain by its external identifier.
func (api *RegisteredDomainsApi) UpdateRegisteredDomainById(extId *string, body *import1.RegisteredDomain, args ...map[string]interface{}) (*import1.UpdateRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRegisteredDomainById(context.Background(), &import6.UpdateRegisteredDomainByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a registered domain by its external identifier.
func (api *RegisteredDomainsServiceApi) UpdateRegisteredDomainById(ctx context.Context, request *import6.UpdateRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import1.UpdateRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.3/config/registered-domains/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateRegisteredDomainApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
