package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/imageratelimitpolicies"
	"net/http"
	"net/url"
	"strings"
)

type ImageRateLimitPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ImageRateLimitPoliciesServiceApi
}

type ImageRateLimitPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewImageRateLimitPoliciesApi(apiClient *client.ApiClient) *ImageRateLimitPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImageRateLimitPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewImageRateLimitPoliciesServiceApi(a.ApiClient)

	return a
}

func NewImageRateLimitPoliciesServiceApi(apiClient *client.ApiClient) *ImageRateLimitPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImageRateLimitPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an image rate limit policy using the provided request body. The name, rate limit Kbps and cluster entity filter are mandatory fields for creating an image rate limit.
func (api *ImageRateLimitPoliciesApi) CreateRateLimitPolicy(body *import6.RateLimitPolicy, args ...map[string]interface{}) (*import6.CreateRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRateLimitPolicy(context.Background(), &import8.CreateRateLimitPolicyRequest{
		Body: body,
	}, args...)
}

// Creates an image rate limit policy using the provided request body. The name, rate limit Kbps and cluster entity filter are mandatory fields for creating an image rate limit.
func (api *ImageRateLimitPoliciesServiceApi) CreateRateLimitPolicy(ctx context.Context, request *import8.CreateRateLimitPolicyRequest, args ...map[string]interface{}) (*import6.CreateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/rate-limit-policies"

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

	unmarshalledResp := new(import6.CreateRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the image rate limit policy with the given external identifier.
func (api *ImageRateLimitPoliciesApi) DeleteRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import6.DeleteRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRateLimitPolicyById(context.Background(), &import8.DeleteRateLimitPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the image rate limit policy with the given external identifier.
func (api *ImageRateLimitPoliciesServiceApi) DeleteRateLimitPolicyById(ctx context.Context, request *import8.DeleteRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import6.DeleteRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/rate-limit-policies/{extId}"

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

	unmarshalledResp := new(import6.DeleteRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves an image rate limit policy details for the provided external identifier.
func (api *ImageRateLimitPoliciesApi) GetRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import6.GetRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRateLimitPolicyById(context.Background(), &import8.GetRateLimitPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves an image rate limit policy details for the provided external identifier.
func (api *ImageRateLimitPoliciesServiceApi) GetRateLimitPolicyById(ctx context.Context, request *import8.GetRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import6.GetRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/rate-limit-policies/{extId}"

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

	unmarshalledResp := new(import6.GetRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// The effective rate limit for the Prism Elements. If no rate limit applies to a particular cluster, no entry is returned for that cluster. The API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListEffectiveRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import6.ListEffectiveRateLimitPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEffectiveRateLimitPolicies(context.Background(), &import8.ListEffectiveRateLimitPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// The effective rate limit for the Prism Elements. If no rate limit applies to a particular cluster, no entry is returned for that cluster. The API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesServiceApi) ListEffectiveRateLimitPolicies(ctx context.Context, request *import8.ListEffectiveRateLimitPoliciesRequest, args ...map[string]interface{}) (*import6.ListEffectiveRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/effective-rate-limit-policies"

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

	unmarshalledResp := new(import6.ListEffectiveRateLimitPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists image rate limit policies created on Prism Central along with other details such as, name, description and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import6.ListRateLimitPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRateLimitPolicies(context.Background(), &import8.ListRateLimitPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists image rate limit policies created on Prism Central along with other details such as, name, description and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesServiceApi) ListRateLimitPolicies(ctx context.Context, request *import8.ListRateLimitPoliciesRequest, args ...map[string]interface{}) (*import6.ListRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/rate-limit-policies"

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

	unmarshalledResp := new(import6.ListRateLimitPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the image rate limit policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImageRateLimitPoliciesApi) UpdateRateLimitPolicyById(extId *string, body *import6.RateLimitPolicy, args ...map[string]interface{}) (*import6.UpdateRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRateLimitPolicyById(context.Background(), &import8.UpdateRateLimitPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the image rate limit policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImageRateLimitPoliciesServiceApi) UpdateRateLimitPolicyById(ctx context.Context, request *import8.UpdateRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import6.UpdateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/images/config/rate-limit-policies/{extId}"

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

	unmarshalledResp := new(import6.UpdateRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
