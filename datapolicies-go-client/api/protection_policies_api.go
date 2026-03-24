package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/request/protectionpolicies"
	"net/http"
	"net/url"
	"strings"
)

type ProtectionPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ProtectionPoliciesServiceApi
}

type ProtectionPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewProtectionPoliciesApi(apiClient *client.ApiClient) *ProtectionPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ProtectionPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewProtectionPoliciesServiceApi(a.ApiClient)

	return a
}

func NewProtectionPoliciesServiceApi(apiClient *client.ApiClient) *ProtectionPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ProtectionPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new consistency rule with the provided specification.
func (api *ProtectionPoliciesApi) CreateConsistencyRule(protectionPolicyExtId *string, body *import1.ConsistencyRule, args ...map[string]interface{}) (*import1.CreateConsistencyRuleApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateConsistencyRule(context.Background(), &import3.CreateConsistencyRuleRequest{
		ProtectionPolicyExtId: protectionPolicyExtId,
		Body:                  body,
	}, args...)
}

// Creates a new consistency rule with the provided specification.
func (api *ProtectionPoliciesServiceApi) CreateConsistencyRule(ctx context.Context, request *import3.CreateConsistencyRuleRequest, args ...map[string]interface{}) (*import1.CreateConsistencyRuleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{protectionPolicyExtId}/consistency-rules"

	// verify the required parameter 'protectionPolicyExtId' is set
	if nil == request.ProtectionPolicyExtId {
		return nil, client.ReportError("protectionPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"protectionPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.ProtectionPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.CreateConsistencyRuleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a protection policy to automate the recovery point creation and replication process.
func (api *ProtectionPoliciesApi) CreateProtectionPolicy(body *import1.ProtectionPolicy, args ...map[string]interface{}) (*import1.CreateProtectionPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateProtectionPolicy(context.Background(), &import3.CreateProtectionPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a protection policy to automate the recovery point creation and replication process.
func (api *ProtectionPoliciesServiceApi) CreateProtectionPolicy(ctx context.Context, request *import3.CreateProtectionPolicyRequest, args ...map[string]interface{}) (*import1.CreateProtectionPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies"

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

	unmarshalledResp := new(import1.CreateProtectionPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesApi) DeleteConsistencyRuleById(protectionPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteConsistencyRuleApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteConsistencyRuleById(context.Background(), &import3.DeleteConsistencyRuleByIdRequest{
		ProtectionPolicyExtId: protectionPolicyExtId,
		ExtId:                 extId,
	}, args...)
}

// Deletes a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesServiceApi) DeleteConsistencyRuleById(ctx context.Context, request *import3.DeleteConsistencyRuleByIdRequest, args ...map[string]interface{}) (*import1.DeleteConsistencyRuleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{protectionPolicyExtId}/consistency-rules/{extId}"

	// verify the required parameter 'protectionPolicyExtId' is set
	if nil == request.ProtectionPolicyExtId {
		return nil, client.ReportError("protectionPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"protectionPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.ProtectionPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.DeleteConsistencyRuleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the protection policy identified by an external identifier.
func (api *ProtectionPoliciesApi) DeleteProtectionPolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteProtectionPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteProtectionPolicyById(context.Background(), &import3.DeleteProtectionPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the protection policy identified by an external identifier.
func (api *ProtectionPoliciesServiceApi) DeleteProtectionPolicyById(ctx context.Context, request *import3.DeleteProtectionPolicyByIdRequest, args ...map[string]interface{}) (*import1.DeleteProtectionPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{extId}"

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

	unmarshalledResp := new(import1.DeleteProtectionPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the details of a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesApi) GetConsistencyRuleById(protectionPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetConsistencyRuleApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetConsistencyRuleById(context.Background(), &import3.GetConsistencyRuleByIdRequest{
		ProtectionPolicyExtId: protectionPolicyExtId,
		ExtId:                 extId,
	}, args...)
}

// Fetches the details of a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesServiceApi) GetConsistencyRuleById(ctx context.Context, request *import3.GetConsistencyRuleByIdRequest, args ...map[string]interface{}) (*import1.GetConsistencyRuleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{protectionPolicyExtId}/consistency-rules/{extId}"

	// verify the required parameter 'protectionPolicyExtId' is set
	if nil == request.ProtectionPolicyExtId {
		return nil, client.ReportError("protectionPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"protectionPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.ProtectionPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetConsistencyRuleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the protection policy identified by an external identifier.
func (api *ProtectionPoliciesApi) GetProtectionPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetProtectionPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetProtectionPolicyById(context.Background(), &import3.GetProtectionPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the protection policy identified by an external identifier.
func (api *ProtectionPoliciesServiceApi) GetProtectionPolicyById(ctx context.Context, request *import3.GetProtectionPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetProtectionPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{extId}"

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

	unmarshalledResp := new(import1.GetProtectionPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the consistency rules defined in the system. The list of consistency rules can be further filtered using various filtering options.
func (api *ProtectionPoliciesApi) ListConsistencyRulesByProtectionPolicyId(protectionPolicyExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListConsistencyRulesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListConsistencyRulesByProtectionPolicyId(context.Background(), &import3.ListConsistencyRulesByProtectionPolicyIdRequest{
		ProtectionPolicyExtId: protectionPolicyExtId,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// Lists the consistency rules defined in the system. The list of consistency rules can be further filtered using various filtering options.
func (api *ProtectionPoliciesServiceApi) ListConsistencyRulesByProtectionPolicyId(ctx context.Context, request *import3.ListConsistencyRulesByProtectionPolicyIdRequest, args ...map[string]interface{}) (*import1.ListConsistencyRulesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{protectionPolicyExtId}/consistency-rules"

	// verify the required parameter 'protectionPolicyExtId' is set
	if nil == request.ProtectionPolicyExtId {
		return nil, client.ReportError("protectionPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"protectionPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.ProtectionPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListConsistencyRulesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List the protection policies defined on the system. This operation supports filtering, sorting, selection and pagination.
func (api *ProtectionPoliciesApi) ListProtectionPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListProtectionPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListProtectionPolicies(context.Background(), &import3.ListProtectionPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List the protection policies defined on the system. This operation supports filtering, sorting, selection and pagination.
func (api *ProtectionPoliciesServiceApi) ListProtectionPolicies(ctx context.Context, request *import3.ListProtectionPoliciesRequest, args ...map[string]interface{}) (*import1.ListProtectionPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies"

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

	unmarshalledResp := new(import1.ListProtectionPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the details of a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesApi) UpdateConsistencyRuleById(protectionPolicyExtId *string, extId *string, body *import1.ConsistencyRule, args ...map[string]interface{}) (*import1.UpdateConsistencyRuleApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateConsistencyRuleById(context.Background(), &import3.UpdateConsistencyRuleByIdRequest{
		ProtectionPolicyExtId: protectionPolicyExtId,
		ExtId:                 extId,
		Body:                  body,
	}, args...)
}

// Updates the details of a consistency rule for the provided external identifier.
func (api *ProtectionPoliciesServiceApi) UpdateConsistencyRuleById(ctx context.Context, request *import3.UpdateConsistencyRuleByIdRequest, args ...map[string]interface{}) (*import1.UpdateConsistencyRuleApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{protectionPolicyExtId}/consistency-rules/{extId}"

	// verify the required parameter 'protectionPolicyExtId' is set
	if nil == request.ProtectionPolicyExtId {
		return nil, client.ReportError("protectionPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"protectionPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.ProtectionPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.UpdateConsistencyRuleApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the protection policy identified by an external identifier.
func (api *ProtectionPoliciesApi) UpdateProtectionPolicyById(extId *string, body *import1.ProtectionPolicy, args ...map[string]interface{}) (*import1.UpdateProtectionPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewProtectionPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateProtectionPolicyById(context.Background(), &import3.UpdateProtectionPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the protection policy identified by an external identifier.
func (api *ProtectionPoliciesServiceApi) UpdateProtectionPolicyById(ctx context.Context, request *import3.UpdateProtectionPolicyByIdRequest, args ...map[string]interface{}) (*import1.UpdateProtectionPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/protection-policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateProtectionPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
