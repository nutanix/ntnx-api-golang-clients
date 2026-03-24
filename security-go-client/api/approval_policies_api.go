package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/management"
	import2 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/approvalpolicies"
	"net/http"
	"net/url"
	"strings"
)

type ApprovalPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ApprovalPoliciesServiceApi
}

type ApprovalPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewApprovalPoliciesApi(apiClient *client.ApiClient) *ApprovalPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ApprovalPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewApprovalPoliciesServiceApi(a.ApiClient)

	return a
}

func NewApprovalPoliciesServiceApi(apiClient *client.ApiClient) *ApprovalPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ApprovalPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Update the list of secured policies linked to an approval policy with particular external identifier.
func (api *ApprovalPoliciesApi) AssociatePolicies(extId *string, body *import1.AssociatePoliciesSpec, args ...map[string]interface{}) (*import1.AssociatePoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewApprovalPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociatePolicies(context.Background(), &import2.AssociatePoliciesRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update the list of secured policies linked to an approval policy with particular external identifier.
func (api *ApprovalPoliciesServiceApi) AssociatePolicies(ctx context.Context, request *import2.AssociatePoliciesRequest, args ...map[string]interface{}) (*import1.AssociatePoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/management/approval-policies/{extId}/$actions/associate-policies"

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

	unmarshalledResp := new(import1.AssociatePoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Create an approval policy for secure snapshots.
func (api *ApprovalPoliciesApi) CreateApprovalPolicy(body *import1.ApprovalPolicy, args ...map[string]interface{}) (*import1.CreateApprovalPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewApprovalPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateApprovalPolicy(context.Background(), &import2.CreateApprovalPolicyRequest{
		Body: body,
	}, args...)
}

// Create an approval policy for secure snapshots.
func (api *ApprovalPoliciesServiceApi) CreateApprovalPolicy(ctx context.Context, request *import2.CreateApprovalPolicyRequest, args ...map[string]interface{}) (*import1.CreateApprovalPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/management/approval-policies"

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

	unmarshalledResp := new(import1.CreateApprovalPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the details of a specific approval policy by its external identifier.
func (api *ApprovalPoliciesApi) GetApprovalPolicyByExtId(extId *string, select_ *string, args ...map[string]interface{}) (*import1.GetApprovalPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewApprovalPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetApprovalPolicyByExtId(context.Background(), &import2.GetApprovalPolicyByExtIdRequest{
		ExtId:   extId,
		Select_: select_,
	}, args...)
}

// Fetches the details of a specific approval policy by its external identifier.
func (api *ApprovalPoliciesServiceApi) GetApprovalPolicyByExtId(ctx context.Context, request *import2.GetApprovalPolicyByExtIdRequest, args ...map[string]interface{}) (*import1.GetApprovalPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/management/approval-policies/{extId}"

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

	// Query Params
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

	unmarshalledResp := new(import1.GetApprovalPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of approval policies for secure snapshots.
func (api *ApprovalPoliciesApi) ListApprovalPolicies(filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListApprovalPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewApprovalPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListApprovalPolicies(context.Background(), &import2.ListApprovalPoliciesRequest{
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of approval policies for secure snapshots.
func (api *ApprovalPoliciesServiceApi) ListApprovalPolicies(ctx context.Context, request *import2.ListApprovalPoliciesRequest, args ...map[string]interface{}) (*import1.ListApprovalPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/management/approval-policies"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
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

	unmarshalledResp := new(import1.ListApprovalPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the details of a specific approval policy by its external identifier.
func (api *ApprovalPoliciesApi) UpdateApprovalPolicyByExtId(extId *string, body *import1.ApprovalPolicy, args ...map[string]interface{}) (*import1.UpdateApprovalPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewApprovalPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateApprovalPolicyByExtId(context.Background(), &import2.UpdateApprovalPolicyByExtIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the details of a specific approval policy by its external identifier.
func (api *ApprovalPoliciesServiceApi) UpdateApprovalPolicyByExtId(ctx context.Context, request *import2.UpdateApprovalPolicyByExtIdRequest, args ...map[string]interface{}) (*import1.UpdateApprovalPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/management/approval-policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateApprovalPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
