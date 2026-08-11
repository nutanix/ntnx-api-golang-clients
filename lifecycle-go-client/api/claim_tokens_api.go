package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/claimtokens"
	"net/http"
	"net/url"
	"strings"
)

type ClaimTokensApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClaimTokensServiceApi
}

type ClaimTokensServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClaimTokensApi(apiClient *client.ApiClient) *ClaimTokensApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClaimTokensApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClaimTokensServiceApi(a.ApiClient)

	return a
}

func NewClaimTokensServiceApi(apiClient *client.ApiClient) *ClaimTokensServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClaimTokensServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a claim token with a specified name, expiry time, and maximum usage count
func (api *ClaimTokensApi) CreateClaimToken(body *import3.ClaimToken, args ...map[string]interface{}) (*import3.CreateClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateClaimToken(context.Background(), &import4.CreateClaimTokenRequest{
		Body: body,
	}, args...)
}

// Creates a claim token with a specified name, expiry time, and maximum usage count
func (api *ClaimTokensServiceApi) CreateClaimToken(ctx context.Context, request *import4.CreateClaimTokenRequest, args ...map[string]interface{}) (*import3.CreateClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens"

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
	unmarshalledResp := new(import3.CreateClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the claim token identified by the given external ID
func (api *ClaimTokensApi) DeleteClaimTokenById(extId *string, args ...map[string]interface{}) (*import3.DeleteClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteClaimTokenById(context.Background(), &import4.DeleteClaimTokenByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the claim token identified by the given external ID
func (api *ClaimTokensServiceApi) DeleteClaimTokenById(ctx context.Context, request *import4.DeleteClaimTokenByIdRequest, args ...map[string]interface{}) (*import3.DeleteClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{extId}"

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
	unmarshalledResp := new(import3.DeleteClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a claim token identified by its external ID
func (api *ClaimTokensApi) GetClaimTokenById(extId *string, args ...map[string]interface{}) (*import3.GetClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetClaimTokenById(context.Background(), &import4.GetClaimTokenByIdRequest{
		ExtId: extId,
	}, args...)
}

// Returns details of a claim token identified by its external ID
func (api *ClaimTokensServiceApi) GetClaimTokenById(ctx context.Context, request *import4.GetClaimTokenByIdRequest, args ...map[string]interface{}) (*import3.GetClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{extId}"

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
	unmarshalledResp := new(import3.GetClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns details of a node registered with a specific claim token
func (api *ClaimTokensApi) GetNodeByClaimTokenId(claimTokenExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetNodeByClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNodeByClaimTokenId(context.Background(), &import4.GetNodeByClaimTokenIdRequest{
		ClaimTokenExtId: claimTokenExtId,
		ExtId:           extId,
	}, args...)
}

// Returns details of a node registered with a specific claim token
func (api *ClaimTokensServiceApi) GetNodeByClaimTokenId(ctx context.Context, request *import4.GetNodeByClaimTokenIdRequest, args ...map[string]interface{}) (*import3.GetNodeByClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{claimTokenExtId}/nodes/{extId}"

	// verify the required parameter 'claimTokenExtId' is set
	if nil == request.ClaimTokenExtId {
		return nil, client.ReportError("claimTokenExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"claimTokenExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClaimTokenExtId, "")), -1)
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
	unmarshalledResp := new(import3.GetNodeByClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns the secret value of a claim token identified by its external ID
func (api *ClaimTokensApi) GetSecretByClaimTokenId(extId *string, args ...map[string]interface{}) (*import3.GetClaimTokenSecretApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSecretByClaimTokenId(context.Background(), &import4.GetSecretByClaimTokenIdRequest{
		ExtId: extId,
	}, args...)
}

// Returns the secret value of a claim token identified by its external ID
func (api *ClaimTokensServiceApi) GetSecretByClaimTokenId(ctx context.Context, request *import4.GetSecretByClaimTokenIdRequest, args ...map[string]interface{}) (*import3.GetClaimTokenSecretApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{extId}/secret"

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
	unmarshalledResp := new(import3.GetClaimTokenSecretApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a paginated list of all claim tokens
func (api *ClaimTokensApi) ListClaimTokens(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListClaimTokensApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListClaimTokens(context.Background(), &import4.ListClaimTokensRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Returns a paginated list of all claim tokens
func (api *ClaimTokensServiceApi) ListClaimTokens(ctx context.Context, request *import4.ListClaimTokensRequest, args ...map[string]interface{}) (*import3.ListClaimTokensApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens"

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
	unmarshalledResp := new(import3.ListClaimTokensApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns a list of nodes registered with a specific claim token
func (api *ClaimTokensApi) ListNodesByClaimTokenId(claimTokenExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListNodesByClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNodesByClaimTokenId(context.Background(), &import4.ListNodesByClaimTokenIdRequest{
		ClaimTokenExtId: claimTokenExtId,
		Page_:           page_,
		Limit_:          limit_,
		Filter_:         filter_,
		Orderby_:        orderby_,
		Select_:         select_,
	}, args...)
}

// Returns a list of nodes registered with a specific claim token
func (api *ClaimTokensServiceApi) ListNodesByClaimTokenId(ctx context.Context, request *import4.ListNodesByClaimTokenIdRequest, args ...map[string]interface{}) (*import3.ListNodesByClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{claimTokenExtId}/nodes"

	// verify the required parameter 'claimTokenExtId' is set
	if nil == request.ClaimTokenExtId {
		return nil, client.ReportError("claimTokenExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"claimTokenExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClaimTokenExtId, "")), -1)
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
	unmarshalledResp := new(import3.ListNodesByClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the details of a claim token identified by its external ID
func (api *ClaimTokensApi) UpdateClaimTokenById(extId *string, body *import3.ClaimToken, args ...map[string]interface{}) (*import3.UpdateClaimTokenApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClaimTokensServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateClaimTokenById(context.Background(), &import4.UpdateClaimTokenByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the details of a claim token identified by its external ID
func (api *ClaimTokensServiceApi) UpdateClaimTokenById(ctx context.Context, request *import4.UpdateClaimTokenByIdRequest, args ...map[string]interface{}) (*import3.UpdateClaimTokenApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/config/claim-tokens/{extId}"

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
	unmarshalledResp := new(import3.UpdateClaimTokenApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
