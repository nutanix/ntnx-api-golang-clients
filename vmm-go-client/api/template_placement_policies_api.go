package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import16 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/config"
	import17 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/templateplacementpolicies"
	"net/http"
	"net/url"
	"strings"
)

type TemplatePlacementPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *TemplatePlacementPoliciesServiceApi
}

type TemplatePlacementPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewTemplatePlacementPoliciesApi(apiClient *client.ApiClient) *TemplatePlacementPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TemplatePlacementPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewTemplatePlacementPoliciesServiceApi(a.ApiClient)

	return a
}

func NewTemplatePlacementPoliciesServiceApi(apiClient *client.ApiClient) *TemplatePlacementPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TemplatePlacementPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a template placement policy based on the provided request body.
func (api *TemplatePlacementPoliciesApi) CreateTemplatePlacementPolicy(body *import16.TemplatePlacementPolicy, args ...map[string]interface{}) (*import16.CreateTemplatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateTemplatePlacementPolicy(context.Background(), &import17.CreateTemplatePlacementPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a template placement policy based on the provided request body.
func (api *TemplatePlacementPoliciesServiceApi) CreateTemplatePlacementPolicy(ctx context.Context, request *import17.CreateTemplatePlacementPolicyRequest, args ...map[string]interface{}) (*import16.CreateTemplatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/config/template-placement-policies"

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

	unmarshalledResp := new(import16.CreateTemplatePlacementPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the template placement policy associated with the given external identifier.
func (api *TemplatePlacementPoliciesApi) DeleteTemplatePlacementPolicyById(extId *string, args ...map[string]interface{}) (*import16.DeleteTemplatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteTemplatePlacementPolicyById(context.Background(), &import17.DeleteTemplatePlacementPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the template placement policy associated with the given external identifier.
func (api *TemplatePlacementPoliciesServiceApi) DeleteTemplatePlacementPolicyById(ctx context.Context, request *import17.DeleteTemplatePlacementPolicyByIdRequest, args ...map[string]interface{}) (*import16.DeleteTemplatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/config/template-placement-policies/{extId}"

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

	unmarshalledResp := new(import16.DeleteTemplatePlacementPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the details of the template placement policy for the provided external identifier.
func (api *TemplatePlacementPoliciesApi) GetTemplatePlacementPolicyById(extId *string, args ...map[string]interface{}) (*import16.GetTemplatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTemplatePlacementPolicyById(context.Background(), &import17.GetTemplatePlacementPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the details of the template placement policy for the provided external identifier.
func (api *TemplatePlacementPoliciesServiceApi) GetTemplatePlacementPolicyById(ctx context.Context, request *import17.GetTemplatePlacementPolicyByIdRequest, args ...map[string]interface{}) (*import16.GetTemplatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/config/template-placement-policies/{extId}"

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

	unmarshalledResp := new(import16.GetTemplatePlacementPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the template placement policies created on Prism Central, including details such as name, description and more. The API supports operations like filtering, sorting, selection and pagination.
func (api *TemplatePlacementPoliciesApi) ListTemplatePlacementPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import16.ListTemplatePlacementPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTemplatePlacementPolicies(context.Background(), &import17.ListTemplatePlacementPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists the template placement policies created on Prism Central, including details such as name, description and more. The API supports operations like filtering, sorting, selection and pagination.
func (api *TemplatePlacementPoliciesServiceApi) ListTemplatePlacementPolicies(ctx context.Context, request *import17.ListTemplatePlacementPoliciesRequest, args ...map[string]interface{}) (*import16.ListTemplatePlacementPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/config/template-placement-policies"

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

	unmarshalledResp := new(import16.ListTemplatePlacementPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Revise the template placement policy using the specified external identifier and the provided request body. It is highly recommended to perform a GET operation on a resource prior to executing the PUT operation to ensure the correct ETag is utilized.
func (api *TemplatePlacementPoliciesApi) UpdateTemplatePlacementPolicyById(extId *string, body *import16.TemplatePlacementPolicy, args ...map[string]interface{}) (*import16.UpdateTemplatePlacementPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTemplatePlacementPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateTemplatePlacementPolicyById(context.Background(), &import17.UpdateTemplatePlacementPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Revise the template placement policy using the specified external identifier and the provided request body. It is highly recommended to perform a GET operation on a resource prior to executing the PUT operation to ensure the correct ETag is utilized.
func (api *TemplatePlacementPoliciesServiceApi) UpdateTemplatePlacementPolicyById(ctx context.Context, request *import17.UpdateTemplatePlacementPolicyByIdRequest, args ...map[string]interface{}) (*import16.UpdateTemplatePlacementPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/config/template-placement-policies/{extId}"

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

	unmarshalledResp := new(import16.UpdateTemplatePlacementPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
