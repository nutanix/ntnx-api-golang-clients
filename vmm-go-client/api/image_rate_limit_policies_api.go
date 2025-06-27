package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
	"net/http"
	"net/url"
	"strings"
)

type ImageRateLimitPoliciesApi struct {
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

	return a
}

// Creates an image rate limit policy using the provided request body. The name, rate limit Kbps and cluster entity filter are mandatory fields for creating an image rate limit.
func (api *ImageRateLimitPoliciesApi) CreateRateLimitPolicy(body *import4.RateLimitPolicy, args ...map[string]interface{}) (*import4.CreateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/rate-limit-policies"

	// verify the required parameter 'body' is set
	if nil == body {
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.CreateRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the image rate limit policy with the given external identifier.
func (api *ImageRateLimitPoliciesApi) DeleteRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import4.DeleteRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/rate-limit-policies/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.DeleteRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves an image rate limit policy details for the provided external identifier.
func (api *ImageRateLimitPoliciesApi) GetRateLimitPolicyById(extId *string, args ...map[string]interface{}) (*import4.GetRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/rate-limit-policies/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.GetRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// The effective rate limit for the Prism Elements. If no rate limit applies to a particular cluster, no entry is returned for that cluster. The API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListEffectiveRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListEffectiveRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/effective-rate-limit-policies"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.ListEffectiveRateLimitPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists image rate limit policies created on Prism Central along with other details such as, name, description and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImageRateLimitPoliciesApi) ListRateLimitPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListRateLimitPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/rate-limit-policies"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.ListRateLimitPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the image rate limit policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImageRateLimitPoliciesApi) UpdateRateLimitPolicyById(extId *string, body *import4.RateLimitPolicy, args ...map[string]interface{}) (*import4.UpdateRateLimitPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.1/images/config/rate-limit-policies/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.UpdateRateLimitPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
