package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import10 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/passwordmanager"
	"net/http"
	"net/url"
	"strings"
)

type PasswordManagerApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *PasswordManagerServiceApi
}

type PasswordManagerServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewPasswordManagerApi(apiClient *client.ApiClient) *PasswordManagerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PasswordManagerApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewPasswordManagerServiceApi(a.ApiClient)

	return a
}

func NewPasswordManagerServiceApi(apiClient *client.ApiClient) *PasswordManagerServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PasswordManagerServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Initiate change password request for a system user on a supported product.
func (api *PasswordManagerApi) ChangeSystemUserPasswordById(extId *string, body *import1.ChangePasswordSpec, args ...map[string]interface{}) (*import1.ChangeSystemUserPasswordApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewPasswordManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ChangeSystemUserPasswordById(context.Background(), &import10.ChangeSystemUserPasswordByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Initiate change password request for a system user on a supported product.
func (api *PasswordManagerServiceApi) ChangeSystemUserPasswordById(ctx context.Context, request *import10.ChangeSystemUserPasswordByIdRequest, args ...map[string]interface{}) (*import1.ChangeSystemUserPasswordApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/system-user-passwords/{extId}/$actions/change-password"

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

	unmarshalledResp := new(import1.ChangeSystemUserPasswordApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists password status of system user accounts on supported products.
func (api *PasswordManagerApi) ListSystemUserPasswords(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListSystemUserPasswordsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewPasswordManagerServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSystemUserPasswords(context.Background(), &import10.ListSystemUserPasswordsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Lists password status of system user accounts on supported products.
func (api *PasswordManagerServiceApi) ListSystemUserPasswords(ctx context.Context, request *import10.ListSystemUserPasswordsRequest, args ...map[string]interface{}) (*import1.ListSystemUserPasswordsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/system-user-passwords"

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

	unmarshalledResp := new(import1.ListSystemUserPasswordsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
