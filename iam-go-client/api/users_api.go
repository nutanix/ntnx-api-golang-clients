package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import15 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/users"
	"net/http"
	"net/url"
	"strings"
)

type UsersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UsersServiceApi
}

type UsersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUsersApi(apiClient *client.ApiClient) *UsersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UsersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUsersServiceApi(a.ApiClient)

	return a
}

func NewUsersServiceApi(apiClient *client.ApiClient) *UsersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UsersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Change Password for a user with the given username.
func (api *UsersApi) ChangeUserPassword(body *import4.PasswordChangeRequest, args ...map[string]interface{}) (*import4.ChangeUserPasswordApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ChangeUserPassword(context.Background(), &import15.ChangeUserPasswordRequest{
		Body: body,
	}, args...)
}

// Change Password for a user with the given username.
func (api *UsersServiceApi) ChangeUserPassword(ctx context.Context, request *import15.ChangeUserPasswordRequest, args ...map[string]interface{}) (*import4.ChangeUserPasswordApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/$actions/change-password"

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

	unmarshalledResp := new(import4.ChangeUserPasswordApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a user of the type specified in the request.
func (api *UsersApi) CreateUser(body *import4.User, args ...map[string]interface{}) (*import4.CreateUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateUser(context.Background(), &import15.CreateUserRequest{
		Body: body,
	}, args...)
}

// Creates a user of the type specified in the request.
func (api *UsersServiceApi) CreateUser(ctx context.Context, request *import15.CreateUserRequest, args ...map[string]interface{}) (*import4.CreateUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users"

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

	unmarshalledResp := new(import4.CreateUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create key of a requested type for a user.
func (api *UsersApi) CreateUserKey(userExtId *string, body *import4.Key, args ...map[string]interface{}) (*import4.CreateKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateUserKey(context.Background(), &import15.CreateUserKeyRequest{
		UserExtId: userExtId,
		Body:      body,
	}, args...)
}

// Create key of a requested type for a user.
func (api *UsersServiceApi) CreateUserKey(ctx context.Context, request *import15.CreateUserKeyRequest, args ...map[string]interface{}) (*import4.CreateKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{userExtId}/keys"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import4.CreateKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the requested key.
func (api *UsersApi) DeleteUserKeyById(userExtId *string, extId *string, args ...map[string]interface{}) (*import4.DeleteUserKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteUserKeyById(context.Background(), &import15.DeleteUserKeyByIdRequest{
		UserExtId: userExtId,
		ExtId:     extId,
	}, args...)
}

// Delete the requested key.
func (api *UsersServiceApi) DeleteUserKeyById(ctx context.Context, request *import15.DeleteUserKeyByIdRequest, args ...map[string]interface{}) (*import4.DeleteUserKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{userExtId}/keys/{extId}"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import4.DeleteUserKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a user based on its external identifier.
func (api *UsersApi) GetUserById(extId *string, args ...map[string]interface{}) (*import4.GetUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUserById(context.Background(), &import15.GetUserByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a user based on its external identifier.
func (api *UsersServiceApi) GetUserById(ctx context.Context, request *import15.GetUserByIdRequest, args ...map[string]interface{}) (*import4.GetUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{extId}"

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

	unmarshalledResp := new(import4.GetUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the requested key through the provided external identifier for the user and the key.
func (api *UsersApi) GetUserKeyById(userExtId *string, extId *string, args ...map[string]interface{}) (*import4.GetUserKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUserKeyById(context.Background(), &import15.GetUserKeyByIdRequest{
		UserExtId: userExtId,
		ExtId:     extId,
	}, args...)
}

// Fetches the requested key through the provided external identifier for the user and the key.
func (api *UsersServiceApi) GetUserKeyById(ctx context.Context, request *import15.GetUserKeyByIdRequest, args ...map[string]interface{}) (*import4.GetUserKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{userExtId}/keys/{extId}"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import4.GetUserKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List all keys identified by the external identifier of a user.
func (api *UsersApi) ListUserKeys(userExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListUserKeysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUserKeys(context.Background(), &import15.ListUserKeysRequest{
		UserExtId: userExtId,
		Page_:     page_,
		Limit_:    limit_,
		Filter_:   filter_,
		Orderby_:  orderby_,
		Select_:   select_,
	}, args...)
}

// List all keys identified by the external identifier of a user.
func (api *UsersServiceApi) ListUserKeys(ctx context.Context, request *import15.ListUserKeysRequest, args ...map[string]interface{}) (*import4.ListUserKeysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{userExtId}/keys"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListUserKeysApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Returns a list of all non-internal user(s) in the system.
func (api *UsersApi) ListUsers(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListUsersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUsers(context.Background(), &import15.ListUsersRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Returns a list of all non-internal user(s) in the system.
func (api *UsersServiceApi) ListUsers(ctx context.Context, request *import15.ListUsersRequest, args ...map[string]interface{}) (*import4.ListUsersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users"

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

	unmarshalledResp := new(import4.ListUsersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Reset user password based on the provided external identifier.
func (api *UsersApi) ResetUserPassword(extId *string, body *import4.PasswordResetRequest, args ...map[string]interface{}) (*import4.ResetUserPasswordApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ResetUserPassword(context.Background(), &import15.ResetUserPasswordRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Reset user password based on the provided external identifier.
func (api *UsersServiceApi) ResetUserPassword(ctx context.Context, request *import15.ResetUserPasswordRequest, args ...map[string]interface{}) (*import4.ResetUserPasswordApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{extId}/$actions/reset-password"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.ResetUserPasswordApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Revoke the requested key.
func (api *UsersApi) RevokeUserKey(userExtId *string, extId *string, args ...map[string]interface{}) (*import4.RevokeUserKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RevokeUserKey(context.Background(), &import15.RevokeUserKeyRequest{
		UserExtId: userExtId,
		ExtId:     extId,
	}, args...)
}

// Revoke the requested key.
func (api *UsersServiceApi) RevokeUserKey(ctx context.Context, request *import15.RevokeUserKeyRequest, args ...map[string]interface{}) (*import4.RevokeUserKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{userExtId}/keys/{extId}/$actions/revoke"

	// verify the required parameter 'userExtId' is set
	if nil == request.UserExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*request.UserExtId, "")), -1)
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

	unmarshalledResp := new(import4.RevokeUserKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a user based on the provided external identifier.
func (api *UsersApi) UpdateUserById(extId *string, body *import4.User, args ...map[string]interface{}) (*import4.UpdateUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateUserById(context.Background(), &import15.UpdateUserByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a user based on the provided external identifier.
func (api *UsersServiceApi) UpdateUserById(ctx context.Context, request *import15.UpdateUserByIdRequest, args ...map[string]interface{}) (*import4.UpdateUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{extId}"

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

	unmarshalledResp := new(import4.UpdateUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the active state of a user based on the provided external identifier.
func (api *UsersApi) UpdateUserState(extId *string, body *import4.UserStateUpdate, args ...map[string]interface{}) (*import4.ActivateUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUsersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateUserState(context.Background(), &import15.UpdateUserStateRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the active state of a user based on the provided external identifier.
func (api *UsersServiceApi) UpdateUserState(ctx context.Context, request *import15.UpdateUserStateRequest, args ...map[string]interface{}) (*import4.ActivateUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/users/{extId}/$actions/change-state"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.ActivateUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
