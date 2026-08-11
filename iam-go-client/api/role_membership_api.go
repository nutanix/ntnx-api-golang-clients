package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
	import11 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/rolemembership"
	"net/http"
	"net/url"
	"strings"
)

type RoleMembershipApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RoleMembershipServiceApi
}

type RoleMembershipServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRoleMembershipApi(apiClient *client.ApiClient) *RoleMembershipApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoleMembershipApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRoleMembershipServiceApi(a.ApiClient)

	return a
}

func NewRoleMembershipServiceApi(apiClient *client.ApiClient) *RoleMembershipServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RoleMembershipServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a role membership.
func (api *RoleMembershipApi) CreateRoleMembership(body *import1.RoleMembership, args ...map[string]interface{}) (*import1.CreateRoleMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoleMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRoleMembership(context.Background(), &import11.CreateRoleMembershipRequest{
		Body: body,
	}, args...)
}

// Creates a role membership.
func (api *RoleMembershipServiceApi) CreateRoleMembership(ctx context.Context, request *import11.CreateRoleMembershipRequest, args ...map[string]interface{}) (*import1.CreateRoleMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authz/role-memberships"

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
	unmarshalledResp := new(import1.CreateRoleMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes a role membership based on the provided external identifier.
func (api *RoleMembershipApi) DeleteRoleMembershipById(extId *string, args ...map[string]interface{}) (*import1.DeleteRoleMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoleMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRoleMembershipById(context.Background(), &import11.DeleteRoleMembershipByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a role membership based on the provided external identifier.
func (api *RoleMembershipServiceApi) DeleteRoleMembershipById(ctx context.Context, request *import11.DeleteRoleMembershipByIdRequest, args ...map[string]interface{}) (*import1.DeleteRoleMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authz/role-memberships/{extId}"

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
	unmarshalledResp := new(import1.DeleteRoleMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Displays a role membership based on the provided external identifier.
func (api *RoleMembershipApi) GetRoleMembershipById(extId *string, args ...map[string]interface{}) (*import1.GetRoleMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoleMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRoleMembershipById(context.Background(), &import11.GetRoleMembershipByIdRequest{
		ExtId: extId,
	}, args...)
}

// Displays a role membership based on the provided external identifier.
func (api *RoleMembershipServiceApi) GetRoleMembershipById(ctx context.Context, request *import11.GetRoleMembershipByIdRequest, args ...map[string]interface{}) (*import1.GetRoleMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authz/role-memberships/{extId}"

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
	unmarshalledResp := new(import1.GetRoleMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists role membership summaries. Each record represents a project and returns the count of identities (users and groups) and roles for that project. Use the $filter query parameter to filter by extId to get the summary for a specific project.
func (api *RoleMembershipApi) ListRoleMembershipSummary(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRoleMembershipSummariesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoleMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRoleMembershipSummary(context.Background(), &import11.ListRoleMembershipSummaryRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists role membership summaries. Each record represents a project and returns the count of identities (users and groups) and roles for that project. Use the $filter query parameter to filter by extId to get the summary for a specific project.
func (api *RoleMembershipServiceApi) ListRoleMembershipSummary(ctx context.Context, request *import11.ListRoleMembershipSummaryRequest, args ...map[string]interface{}) (*import1.ListRoleMembershipSummariesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authz/role-membership-summaries"

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
	unmarshalledResp := new(import1.ListRoleMembershipSummariesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all the role memberships.
func (api *RoleMembershipApi) ListRoleMemberships(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListRoleMembershipsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRoleMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRoleMemberships(context.Background(), &import11.ListRoleMembershipsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// List all the role memberships.
func (api *RoleMembershipServiceApi) ListRoleMemberships(ctx context.Context, request *import11.ListRoleMembershipsRequest, args ...map[string]interface{}) (*import1.ListRoleMembershipsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authz/role-memberships"

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
	unmarshalledResp := new(import1.ListRoleMembershipsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
