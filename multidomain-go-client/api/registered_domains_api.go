package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
	import9 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/request/registereddomains"
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
func (api *RegisteredDomainsApi) CreateRegisteredDomain(body *import3.RegisteredDomain, args ...map[string]interface{}) (*import3.CreateRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRegisteredDomain(context.Background(), &import9.CreateRegisteredDomainRequest{
		Body: body,
	}, args...)
}

// Creates a new registered domain and saves it to database.
func (api *RegisteredDomainsServiceApi) CreateRegisteredDomain(ctx context.Context, request *import9.CreateRegisteredDomainRequest, args ...map[string]interface{}) (*import3.CreateRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains"

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
	unmarshalledResp := new(import3.CreateRegisteredDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the registered domain with the requested external identifier.
func (api *RegisteredDomainsApi) DeleteRegisteredDomainById(extId *string, args ...map[string]interface{}) (*import3.DeleteRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRegisteredDomainById(context.Background(), &import9.DeleteRegisteredDomainByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the registered domain with the requested external identifier.
func (api *RegisteredDomainsServiceApi) DeleteRegisteredDomainById(ctx context.Context, request *import9.DeleteRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import3.DeleteRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains/{extId}"

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
	unmarshalledResp := new(import3.DeleteRegisteredDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves the registered domain by its external identifier.
func (api *RegisteredDomainsApi) GetRegisteredDomainById(extId *string, select_ *string, args ...map[string]interface{}) (*import3.GetRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRegisteredDomainById(context.Background(), &import9.GetRegisteredDomainByIdRequest{
		ExtId:   extId,
		Select_: select_,
	}, args...)
}

// Retrieves the registered domain by its external identifier.
func (api *RegisteredDomainsServiceApi) GetRegisteredDomainById(ctx context.Context, request *import9.GetRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import3.GetRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.GetRegisteredDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all registered domains.
func (api *RegisteredDomainsApi) ListRegisteredDomains(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListRegisteredDomainsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRegisteredDomains(context.Background(), &import9.ListRegisteredDomainsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// List all registered domains.
func (api *RegisteredDomainsServiceApi) ListRegisteredDomains(ctx context.Context, request *import9.ListRegisteredDomainsRequest, args ...map[string]interface{}) (*import3.ListRegisteredDomainsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.ListRegisteredDomainsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Refreshes the API credentials for the registered domain.
func (api *RegisteredDomainsApi) RefreshApiCredentials(extId *string, args ...map[string]interface{}) (*import3.RefreshApiCredentialsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RefreshApiCredentials(context.Background(), &import9.RefreshApiCredentialsRequest{
		ExtId: extId,
	}, args...)
}

// Refreshes the API credentials for the registered domain.
func (api *RegisteredDomainsServiceApi) RefreshApiCredentials(ctx context.Context, request *import9.RefreshApiCredentialsRequest, args ...map[string]interface{}) (*import3.RefreshApiCredentialsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains/{extId}/$actions/refresh-api-credentials"

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
	unmarshalledResp := new(import3.RefreshApiCredentialsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates a registered domain by its external identifier.
func (api *RegisteredDomainsApi) UpdateRegisteredDomainById(extId *string, body *import3.RegisteredDomain, args ...map[string]interface{}) (*import3.UpdateRegisteredDomainApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegisteredDomainsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRegisteredDomainById(context.Background(), &import9.UpdateRegisteredDomainByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a registered domain by its external identifier.
func (api *RegisteredDomainsServiceApi) UpdateRegisteredDomainById(ctx context.Context, request *import9.UpdateRegisteredDomainByIdRequest, args ...map[string]interface{}) (*import3.UpdateRegisteredDomainApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/multidomain/v4.4.b1/config/registered-domains/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import3.UpdateRegisteredDomainApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
