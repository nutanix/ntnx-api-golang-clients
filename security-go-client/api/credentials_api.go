package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/credentials"
	"net/http"
	"net/url"
	"strings"
)

type CredentialsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *CredentialsServiceApi
}

type CredentialsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCredentialsApi(apiClient *client.ApiClient) *CredentialsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CredentialsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewCredentialsServiceApi(a.ApiClient)

	return a
}

func NewCredentialsServiceApi(apiClient *client.ApiClient) *CredentialsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CredentialsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a credential with the provided fields.
func (api *CredentialsApi) CreateCredential(body *import3.Credential, args ...map[string]interface{}) (*import3.CreateCredentialApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCredentialsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCredential(context.Background(), &import4.CreateCredentialRequest{
		Body: body,
	}, args...)
}

// Create a credential with the provided fields.
func (api *CredentialsServiceApi) CreateCredential(ctx context.Context, request *import4.CreateCredentialRequest, args ...map[string]interface{}) (*import3.CreateCredentialApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/credentials"

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

	unmarshalledResp := new(import3.CreateCredentialApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a Credential with the given ExtId.
func (api *CredentialsApi) DeleteCredentialById(extId *string, args ...map[string]interface{}) (*import3.DeleteCredentialApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCredentialsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCredentialById(context.Background(), &import4.DeleteCredentialByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a Credential with the given ExtId.
func (api *CredentialsServiceApi) DeleteCredentialById(ctx context.Context, request *import4.DeleteCredentialByIdRequest, args ...map[string]interface{}) (*import3.DeleteCredentialApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/credentials/{extId}"

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

	unmarshalledResp := new(import3.DeleteCredentialApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves a credential with the given ExtId.
func (api *CredentialsApi) GetCredentialById(extId *string, args ...map[string]interface{}) (*import3.GetCredentialApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCredentialsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCredentialById(context.Background(), &import4.GetCredentialByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves a credential with the given ExtId.
func (api *CredentialsServiceApi) GetCredentialById(ctx context.Context, request *import4.GetCredentialByIdRequest, args ...map[string]interface{}) (*import3.GetCredentialApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/credentials/{extId}"

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

	unmarshalledResp := new(import3.GetCredentialApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Credentials that match the filter provided.
func (api *CredentialsApi) ListCredentials(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListCredentialsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCredentialsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCredentials(context.Background(), &import4.ListCredentialsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List Credentials that match the filter provided.
func (api *CredentialsServiceApi) ListCredentials(ctx context.Context, request *import4.ListCredentialsRequest, args ...map[string]interface{}) (*import3.ListCredentialsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/credentials"

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

	unmarshalledResp := new(import3.ListCredentialsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a credential with the given ExtId.
func (api *CredentialsApi) UpdateCredentialById(extId *string, body *import3.Credential, args ...map[string]interface{}) (*import3.UpdateCredentialApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCredentialsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateCredentialById(context.Background(), &import4.UpdateCredentialByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a credential with the given ExtId.
func (api *CredentialsServiceApi) UpdateCredentialById(ctx context.Context, request *import4.UpdateCredentialByIdRequest, args ...map[string]interface{}) (*import3.UpdateCredentialApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/config/credentials/{extId}"

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

	unmarshalledResp := new(import3.UpdateCredentialApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
