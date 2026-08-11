package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import8 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/imageplacementpolicies"
	"net/http"
	"net/url"
	"strings"
)

type ImagePlacementPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ImagePlacementPoliciesServiceApi
}

type ImagePlacementPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewImagePlacementPoliciesApi(apiClient *client.ApiClient) *ImagePlacementPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImagePlacementPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewImagePlacementPoliciesServiceApi(a.ApiClient)

	return a
}

func NewImagePlacementPoliciesServiceApi(apiClient *client.ApiClient) *ImagePlacementPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImagePlacementPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an image placement policy using the provided request body. The name, placement type, cluster entity filter and image entity filter are mandatory fields for creating an image placement policy.
func (api *ImagePlacementPoliciesApi) CreatePlacementPolicy(body *import8.PlacementPolicy, args ...map[string]interface{}) (*import8.CreatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreatePlacementPolicy(context.Background(), &import9.CreatePlacementPolicyRequest{
		Body: body,
	}, args...)
}

// Creates an image placement policy using the provided request body. The name, placement type, cluster entity filter and image entity filter are mandatory fields for creating an image placement policy.
func (api *ImagePlacementPoliciesServiceApi) CreatePlacementPolicy(ctx context.Context, request *import9.CreatePlacementPolicyRequest, args ...map[string]interface{}) (*import8.CreatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies"

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
	unmarshalledResp := new(import8.CreatePlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the image placement policy with the given external identifier.
func (api *ImagePlacementPoliciesApi) DeletePlacementPolicyById(extId *string, args ...map[string]interface{}) (*import8.DeletePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeletePlacementPolicyById(context.Background(), &import9.DeletePlacementPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the image placement policy with the given external identifier.
func (api *ImagePlacementPoliciesServiceApi) DeletePlacementPolicyById(ctx context.Context, request *import9.DeletePlacementPolicyByIdRequest, args ...map[string]interface{}) (*import8.DeletePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies/{extId}"

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
	unmarshalledResp := new(import8.DeletePlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves the image placement policy details for the provided external identifier.
func (api *ImagePlacementPoliciesApi) GetPlacementPolicyById(extId *string, args ...map[string]interface{}) (*import8.GetPlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetPlacementPolicyById(context.Background(), &import9.GetPlacementPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the image placement policy details for the provided external identifier.
func (api *ImagePlacementPoliciesServiceApi) GetPlacementPolicyById(ctx context.Context, request *import9.GetPlacementPolicyByIdRequest, args ...map[string]interface{}) (*import8.GetPlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies/{extId}"

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
	unmarshalledResp := new(import8.GetPlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists image placement policies created in Prism Central, along with details such as name, description, and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImagePlacementPoliciesApi) ListPlacementPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import8.ListPlacementPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListPlacementPolicies(context.Background(), &import9.ListPlacementPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists image placement policies created in Prism Central, along with details such as name, description, and so on. This API supports operations such as filtering, sorting, selection, and pagination.
func (api *ImagePlacementPoliciesServiceApi) ListPlacementPolicies(ctx context.Context, request *import9.ListPlacementPoliciesRequest, args ...map[string]interface{}) (*import8.ListPlacementPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies"

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
	unmarshalledResp := new(import8.ListPlacementPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Resumes a suspended image placement policy. A suspended image placement policy is one that is not being considered for enforcement.
func (api *ImagePlacementPoliciesApi) ResumePlacementPolicy(extId *string, args ...map[string]interface{}) (*import8.ResumePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ResumePlacementPolicy(context.Background(), &import9.ResumePlacementPolicyRequest{
		ExtId: extId,
	}, args...)
}

// Resumes a suspended image placement policy. A suspended image placement policy is one that is not being considered for enforcement.
func (api *ImagePlacementPoliciesServiceApi) ResumePlacementPolicy(ctx context.Context, request *import9.ResumePlacementPolicyRequest, args ...map[string]interface{}) (*import8.ResumePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies/{extId}/$actions/resume"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import8.ResumePlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Suspends an active image placement policy. An active image placement policy is considered for enforcement.
func (api *ImagePlacementPoliciesApi) SuspendPlacementPolicy(extId *string, body *import8.SuspendPlacementPolicyConfig, args ...map[string]interface{}) (*import8.SuspendPlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SuspendPlacementPolicy(context.Background(), &import9.SuspendPlacementPolicyRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Suspends an active image placement policy. An active image placement policy is considered for enforcement.
func (api *ImagePlacementPoliciesServiceApi) SuspendPlacementPolicy(ctx context.Context, request *import9.SuspendPlacementPolicyRequest, args ...map[string]interface{}) (*import8.SuspendPlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies/{extId}/$actions/suspend"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import8.SuspendPlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the image placement policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImagePlacementPoliciesApi) UpdatePlacementPolicyById(extId *string, body *import8.PlacementPolicy, args ...map[string]interface{}) (*import8.UpdatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdatePlacementPolicyById(context.Background(), &import9.UpdatePlacementPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the image placement policy using the provided request body with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure the correct ETag is used.
func (api *ImagePlacementPoliciesServiceApi) UpdatePlacementPolicyById(ctx context.Context, request *import9.UpdatePlacementPolicyByIdRequest, args ...map[string]interface{}) (*import8.UpdatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/images/config/placement-policies/{extId}"

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
	unmarshalledResp := new(import8.UpdatePlacementPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
