package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import10 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/userdefinedpolicies"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type UserDefinedPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UserDefinedPoliciesServiceApi
}

type UserDefinedPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUserDefinedPoliciesApi(apiClient *client.ApiClient) *UserDefinedPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserDefinedPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUserDefinedPoliciesServiceApi(a.ApiClient)

	return a
}

func NewUserDefinedPoliciesServiceApi(apiClient *client.ApiClient) *UserDefinedPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserDefinedPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new User-Defined Alert policy.
func (api *UserDefinedPoliciesApi) CreateUdaPolicy(body *import1.UserDefinedPolicy, args ...map[string]interface{}) (*import1.CreateUdaPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateUdaPolicy(context.Background(), &import10.CreateUdaPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a new User-Defined Alert policy.
func (api *UserDefinedPoliciesServiceApi) CreateUdaPolicy(ctx context.Context, request *import10.CreateUdaPolicyRequest, args ...map[string]interface{}) (*import1.CreateUdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies"

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

	unmarshalledResp := new(import1.CreateUdaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesApi) DeleteUdaPolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteUdaPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteUdaPolicyById(context.Background(), &import10.DeleteUdaPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesServiceApi) DeleteUdaPolicyById(ctx context.Context, request *import10.DeleteUdaPolicyByIdRequest, args ...map[string]interface{}) (*import1.DeleteUdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId}"

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

	unmarshalledResp := new(import1.DeleteUdaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches all the existing policies with conflicting criteria to a User-Defined Alert policy identified by external identifier..
func (api *UserDefinedPoliciesApi) FindConflictingUdaPolicies(body *import1.UserDefinedPolicy, args ...map[string]interface{}) (*import1.FindConflictingUdaPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.FindConflictingUdaPolicies(context.Background(), &import10.FindConflictingUdaPoliciesRequest{
		Body: body,
	}, args...)
}

// Fetches all the existing policies with conflicting criteria to a User-Defined Alert policy identified by external identifier..
func (api *UserDefinedPoliciesServiceApi) FindConflictingUdaPolicies(ctx context.Context, request *import10.FindConflictingUdaPoliciesRequest, args ...map[string]interface{}) (*import1.FindConflictingUdaPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies/$actions/find-conflicts"

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

	unmarshalledResp := new(import1.FindConflictingUdaPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the details of a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesApi) GetUdaPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetUdaPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUdaPolicyById(context.Background(), &import10.GetUdaPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the details of a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesServiceApi) GetUdaPolicyById(ctx context.Context, request *import10.GetUdaPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetUdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId}"

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

	unmarshalledResp := new(import1.GetUdaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of User-Defined Alert policies.
func (api *UserDefinedPoliciesApi) ListUdaPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListUdaPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUdaPolicies(context.Background(), &import10.ListUdaPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of User-Defined Alert policies.
func (api *UserDefinedPoliciesServiceApi) ListUdaPolicies(ctx context.Context, request *import10.ListUdaPoliciesRequest, args ...map[string]interface{}) (*import1.ListUdaPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies"

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

	unmarshalledResp := new(import1.ListUdaPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesApi) UpdateUdaPolicyById(extId *string, body *import1.UserDefinedPolicy, args ...map[string]interface{}) (*import1.UpdateUdaPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUserDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateUdaPolicyById(context.Background(), &import10.UpdateUdaPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a User-Defined Alert policy identified by external identifier.
func (api *UserDefinedPoliciesServiceApi) UpdateUdaPolicyById(ctx context.Context, request *import10.UpdateUdaPolicyByIdRequest, args ...map[string]interface{}) (*import1.UpdateUdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/user-defined-policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateUdaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
