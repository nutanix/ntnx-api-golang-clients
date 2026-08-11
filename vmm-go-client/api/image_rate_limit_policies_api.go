package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import8 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
	import10 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/imageratelimitpolicies"
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
func (api *ImageRateLimitPoliciesApi) CreateRateLimitPolicy(body *import8.RateLimitPolicy, args ...map[string]interface{}) (*import8.CreateRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRateLimitPolicy(context.Background(), &import10.CreateRateLimitPolicyRequest{
		Body: body,
	}, args...)
}

// Creates an image rate limit policy using the provided request body. The name, rate limit Kbps and cluster entity filter are mandatory fields for creating an image rate limit.
func (api *ImageRateLimitPoliciesServiceApi) CreateRateLimitPolicy(ctx context.Context, request *import10.CreateRateLimitPolicyRequest, args ...map[string]interface{}) (*import8.CreateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/rate-limit-policies"

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
	unmarshalledResp := new(import8.CreateRateLimitPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the image rate limit policy with the given external identifier.
func (api *ImageRateLimitPoliciesApi) DeleteRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import8.DeleteRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRateLimitPolicyById(context.Background(), &import10.DeleteRateLimitPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the image rate limit policy with the given external identifier.
func (api *ImageRateLimitPoliciesServiceApi) DeleteRateLimitPolicyById(ctx context.Context, request *import10.DeleteRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import8.DeleteRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/rate-limit-policies/{extId}"

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
	unmarshalledResp := new(import8.DeleteRateLimitPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves an image rate limit policy details for the provided external identifier.
func (api *ImageRateLimitPoliciesApi) GetRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import8.GetRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRateLimitPolicyById(context.Background(), &import10.GetRateLimitPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves an image rate limit policy details for the provided external identifier.
func (api *ImageRateLimitPoliciesServiceApi) GetRateLimitPolicyById(ctx context.Context, request *import10.GetRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import8.GetRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/rate-limit-policies/{extId}"

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
	unmarshalledResp := new(import8.GetRateLimitPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// The effective rate limit for the Prism Elements. If no rate limit applies to a particular cluster, no entry is returned for that cluster. The API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListEffectiveRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import8.ListEffectiveRateLimitPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEffectiveRateLimitPolicies(context.Background(), &import10.ListEffectiveRateLimitPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// The effective rate limit for the Prism Elements. If no rate limit applies to a particular cluster, no entry is returned for that cluster. The API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesServiceApi) ListEffectiveRateLimitPolicies(ctx context.Context, request *import10.ListEffectiveRateLimitPoliciesRequest, args ...map[string]interface{}) (*import8.ListEffectiveRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/effective-rate-limit-policies"

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
	unmarshalledResp := new(import8.ListEffectiveRateLimitPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists image rate limit policies created on Prism Central along with other details such as, name, description and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import8.ListRateLimitPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRateLimitPolicies(context.Background(), &import10.ListRateLimitPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists image rate limit policies created on Prism Central along with other details such as, name, description and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesServiceApi) ListRateLimitPolicies(ctx context.Context, request *import10.ListRateLimitPoliciesRequest, args ...map[string]interface{}) (*import8.ListRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/rate-limit-policies"

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
	unmarshalledResp := new(import8.ListRateLimitPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the image rate limit policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImageRateLimitPoliciesApi) UpdateRateLimitPolicyById(extId *string, body *import8.RateLimitPolicy, args ...map[string]interface{}) (*import8.UpdateRateLimitPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImageRateLimitPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRateLimitPolicyById(context.Background(), &import10.UpdateRateLimitPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the image rate limit policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImageRateLimitPoliciesServiceApi) UpdateRateLimitPolicyById(ctx context.Context, request *import10.UpdateRateLimitPolicyByIdRequest, args ...map[string]interface{}) (*import8.UpdateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/rate-limit-policies/{extId}"

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
	unmarshalledResp := new(import8.UpdateRateLimitPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
