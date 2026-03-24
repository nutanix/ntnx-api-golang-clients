package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import14 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/usergroups"
	"net/http"
	"net/url"
	"strings"
)

type UserGroupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UserGroupsServiceApi
}

type UserGroupsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUserGroupsApi(apiClient *client.ApiClient) *UserGroupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserGroupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUserGroupsServiceApi(a.ApiClient)

	return a
}

func NewUserGroupsServiceApi(apiClient *client.ApiClient) *UserGroupsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserGroupsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a user group.
func (api *UserGroupsApi) CreateUserGroup(body *import4.UserGroup, args ...map[string]interface{}) (*import4.CreateUserGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateUserGroup(context.Background(), &import14.CreateUserGroupRequest{
		Body: body,
	}, args...)
}

// Creates a user group.
func (api *UserGroupsServiceApi) CreateUserGroup(ctx context.Context, request *import14.CreateUserGroupRequest, args ...map[string]interface{}) (*import4.CreateUserGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/user-groups"

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

	unmarshalledResp := new(import4.CreateUserGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a user group based on the provided external identifier.
func (api *UserGroupsApi) DeleteUserGroupById(extId *string, args ...map[string]interface{}) (*import4.DeleteUserGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteUserGroupById(context.Background(), &import14.DeleteUserGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a user group based on the provided external identifier.
func (api *UserGroupsServiceApi) DeleteUserGroupById(ctx context.Context, request *import14.DeleteUserGroupByIdRequest, args ...map[string]interface{}) (*import4.DeleteUserGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/user-groups/{extId}"

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

	unmarshalledResp := new(import4.DeleteUserGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Displays a user group based on the provided external identifier.
func (api *UserGroupsApi) GetUserGroupById(extId *string, args ...map[string]interface{}) (*import4.GetUserGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUserGroupById(context.Background(), &import14.GetUserGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Displays a user group based on the provided external identifier.
func (api *UserGroupsServiceApi) GetUserGroupById(ctx context.Context, request *import14.GetUserGroupByIdRequest, args ...map[string]interface{}) (*import4.GetUserGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/user-groups/{extId}"

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

	unmarshalledResp := new(import4.GetUserGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all user group(s).
func (api *UserGroupsApi) ListUserGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListUserGroupsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUserGroups(context.Background(), &import14.ListUserGroupsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists all user group(s).
func (api *UserGroupsServiceApi) ListUserGroups(ctx context.Context, request *import14.ListUserGroupsRequest, args ...map[string]interface{}) (*import4.ListUserGroupsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/user-groups"

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

	unmarshalledResp := new(import4.ListUserGroupsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
