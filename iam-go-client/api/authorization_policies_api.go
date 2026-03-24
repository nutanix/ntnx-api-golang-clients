package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
	import2 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/authorizationpolicies"
	"net/http"
	"net/url"
	"strings"
)

type AuthorizationPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *AuthorizationPoliciesServiceApi
}

type AuthorizationPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAuthorizationPoliciesApi(apiClient *client.ApiClient) *AuthorizationPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AuthorizationPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewAuthorizationPoliciesServiceApi(a.ApiClient)

	return a
}

func NewAuthorizationPoliciesServiceApi(apiClient *client.ApiClient) *AuthorizationPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AuthorizationPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an authorization policy.
func (api *AuthorizationPoliciesApi) CreateAuthorizationPolicy(body *import1.AuthorizationPolicy, args ...map[string]interface{}) (*import1.CreateAuthorizationPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuthorizationPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateAuthorizationPolicy(context.Background(), &import2.CreateAuthorizationPolicyRequest{
		Body: body,
	}, args...)
}

// Creates an authorization policy.
func (api *AuthorizationPoliciesServiceApi) CreateAuthorizationPolicy(ctx context.Context, request *import2.CreateAuthorizationPolicyRequest, args ...map[string]interface{}) (*import1.CreateAuthorizationPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/authorization-policies"

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

	unmarshalledResp := new(import1.CreateAuthorizationPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesApi) DeleteAuthorizationPolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteAuthorizationPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuthorizationPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteAuthorizationPolicyById(context.Background(), &import2.DeleteAuthorizationPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesServiceApi) DeleteAuthorizationPolicyById(ctx context.Context, request *import2.DeleteAuthorizationPolicyByIdRequest, args ...map[string]interface{}) (*import1.DeleteAuthorizationPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/authorization-policies/{extId}"

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

	unmarshalledResp := new(import1.DeleteAuthorizationPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Displays an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesApi) GetAuthorizationPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetAuthorizationPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuthorizationPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetAuthorizationPolicyById(context.Background(), &import2.GetAuthorizationPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Displays an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesServiceApi) GetAuthorizationPolicyById(ctx context.Context, request *import2.GetAuthorizationPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetAuthorizationPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/authorization-policies/{extId}"

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

	unmarshalledResp := new(import1.GetAuthorizationPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the authorization policies.
func (api *AuthorizationPoliciesApi) ListAuthorizationPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListAuthorizationPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuthorizationPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAuthorizationPolicies(context.Background(), &import2.ListAuthorizationPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Lists all the authorization policies.
func (api *AuthorizationPoliciesServiceApi) ListAuthorizationPolicies(ctx context.Context, request *import2.ListAuthorizationPoliciesRequest, args ...map[string]interface{}) (*import1.ListAuthorizationPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/authorization-policies"

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

	unmarshalledResp := new(import1.ListAuthorizationPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesApi) UpdateAuthorizationPolicyById(extId *string, body *import1.AuthorizationPolicy, args ...map[string]interface{}) (*import1.UpdateAuthorizationPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuthorizationPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateAuthorizationPolicyById(context.Background(), &import2.UpdateAuthorizationPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates an authorization policy based on the provided external identifier.
func (api *AuthorizationPoliciesServiceApi) UpdateAuthorizationPolicyById(ctx context.Context, request *import2.UpdateAuthorizationPolicyByIdRequest, args ...map[string]interface{}) (*import1.UpdateAuthorizationPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authz/authorization-policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateAuthorizationPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
