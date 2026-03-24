package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import9 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/systemdefinedpolicies"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type SystemDefinedPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SystemDefinedPoliciesServiceApi
}

type SystemDefinedPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSystemDefinedPoliciesApi(apiClient *client.ApiClient) *SystemDefinedPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SystemDefinedPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSystemDefinedPoliciesServiceApi(a.ApiClient)

	return a
}

func NewSystemDefinedPoliciesServiceApi(apiClient *client.ApiClient) *SystemDefinedPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SystemDefinedPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Retrieves the cluster specific configuration associated with a System-Defined Alert Policy identified by external identifier of the System-Defined Alert Policy for a cluster identified by cluster identifier.
func (api *SystemDefinedPoliciesApi) GetClusterConfigById(systemDefinedPolicyExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetClusterConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetClusterConfigById(context.Background(), &import9.GetClusterConfigByIdRequest{
		SystemDefinedPolicyExtId: systemDefinedPolicyExtId,
		ExtId:                    extId,
	}, args...)
}

// Retrieves the cluster specific configuration associated with a System-Defined Alert Policy identified by external identifier of the System-Defined Alert Policy for a cluster identified by cluster identifier.
func (api *SystemDefinedPoliciesServiceApi) GetClusterConfigById(ctx context.Context, request *import9.GetClusterConfigByIdRequest, args ...map[string]interface{}) (*import1.GetClusterConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs/{extId}"

	// verify the required parameter 'systemDefinedPolicyExtId' is set
	if nil == request.SystemDefinedPolicyExtId {
		return nil, client.ReportError("systemDefinedPolicyExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"systemDefinedPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.SystemDefinedPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetClusterConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get details of a System-Defined Alert Policy identified by external identifier.
func (api *SystemDefinedPoliciesApi) GetSdaPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetSdaPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSdaPolicyById(context.Background(), &import9.GetSdaPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get details of a System-Defined Alert Policy identified by external identifier.
func (api *SystemDefinedPoliciesServiceApi) GetSdaPolicyById(ctx context.Context, request *import9.GetSdaPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetSdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/system-defined-policies/{extId}"

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

	unmarshalledResp := new(import1.GetSdaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves cluster specific configurations associated with a System-Defined Alert Policy identified by the external identifier of the System-Defined Alert Policy.
func (api *SystemDefinedPoliciesApi) ListClusterConfigsBySdaId(systemDefinedPolicyExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListClusterConfigsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListClusterConfigsBySdaId(context.Background(), &import9.ListClusterConfigsBySdaIdRequest{
		SystemDefinedPolicyExtId: systemDefinedPolicyExtId,
		Page_:                    page_,
		Limit_:                   limit_,
		Filter_:                  filter_,
		Orderby_:                 orderby_,
		Select_:                  select_,
	}, args...)
}

// Retrieves cluster specific configurations associated with a System-Defined Alert Policy identified by the external identifier of the System-Defined Alert Policy.
func (api *SystemDefinedPoliciesServiceApi) ListClusterConfigsBySdaId(ctx context.Context, request *import9.ListClusterConfigsBySdaIdRequest, args ...map[string]interface{}) (*import1.ListClusterConfigsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs"

	// verify the required parameter 'systemDefinedPolicyExtId' is set
	if nil == request.SystemDefinedPolicyExtId {
		return nil, client.ReportError("systemDefinedPolicyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"systemDefinedPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.SystemDefinedPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListClusterConfigsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves a list of all System-Defined Alert Policies.
func (api *SystemDefinedPoliciesApi) ListSdaPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListSdaPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSdaPolicies(context.Background(), &import9.ListSdaPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Retrieves a list of all System-Defined Alert Policies.
func (api *SystemDefinedPoliciesServiceApi) ListSdaPolicies(ctx context.Context, request *import9.ListSdaPoliciesRequest, args ...map[string]interface{}) (*import1.ListSdaPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/system-defined-policies"

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

	unmarshalledResp := new(import1.ListSdaPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Modifies cluster-specific configuration associated with a System-Defined Alert Policy identified by external identifier of the System-Defined Alert Policy for a cluster identified by cluster identifier.
func (api *SystemDefinedPoliciesApi) UpdateClusterConfigById(systemDefinedPolicyExtId *string, extId *string, body *import1.ClusterConfig, args ...map[string]interface{}) (*import1.UpdateClusterConfigApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSystemDefinedPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateClusterConfigById(context.Background(), &import9.UpdateClusterConfigByIdRequest{
		SystemDefinedPolicyExtId: systemDefinedPolicyExtId,
		ExtId:                    extId,
		Body:                     body,
	}, args...)
}

// Modifies cluster-specific configuration associated with a System-Defined Alert Policy identified by external identifier of the System-Defined Alert Policy for a cluster identified by cluster identifier.
func (api *SystemDefinedPoliciesServiceApi) UpdateClusterConfigById(ctx context.Context, request *import9.UpdateClusterConfigByIdRequest, args ...map[string]interface{}) (*import1.UpdateClusterConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/alerts/system-defined-policies/{systemDefinedPolicyExtId}/cluster-configs/{extId}"

	// verify the required parameter 'systemDefinedPolicyExtId' is set
	if nil == request.SystemDefinedPolicyExtId {
		return nil, client.ReportError("systemDefinedPolicyExtId is required and must be specified")
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
	uri = strings.Replace(uri, "{"+"systemDefinedPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*request.SystemDefinedPolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.UpdateClusterConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
