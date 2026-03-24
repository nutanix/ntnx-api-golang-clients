package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/config"
	import6 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	import7 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/request/networksecuritypolicies"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type NetworkSecurityPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *NetworkSecurityPoliciesServiceApi
}

type NetworkSecurityPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewNetworkSecurityPoliciesApi(apiClient *client.ApiClient) *NetworkSecurityPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NetworkSecurityPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewNetworkSecurityPoliciesServiceApi(a.ApiClient)

	return a
}

func NewNetworkSecurityPoliciesServiceApi(apiClient *client.ApiClient) *NetworkSecurityPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NetworkSecurityPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Imports all the Network Security Policies specified by the data file.
func (api *NetworkSecurityPoliciesApi) ApplyNetworkSecurityPolicyImport(path *string, nTNXPurgePolicies *bool, dryrun_ *bool, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyImportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ApplyNetworkSecurityPolicyImport(context.Background(), &import7.ApplyNetworkSecurityPolicyImportRequest{
		Path:              path,
		NTNXPurgePolicies: nTNXPurgePolicies,
		Dryrun_:           dryrun_,
	}, args...)
}

// Imports all the Network Security Policies specified by the data file.
func (api *NetworkSecurityPoliciesServiceApi) ApplyNetworkSecurityPolicyImport(ctx context.Context, request *import7.ApplyNetworkSecurityPolicyImportRequest, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyImportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/$actions/import"

	// verify the required parameter 'path' is set
	if nil == request.Path {
		return nil, client.ReportError("path is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/octet-stream"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*request.Dryrun_, ""))
	}
	if request.NTNXPurgePolicies != nil {
		headerParams["NTNX-Purge-Policies"] = client.ParameterToString(*request.NTNXPurgePolicies, "")
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

	file, err := os.Open(*request.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	headerParams["Content-Length"] = fmt.Sprintf("%d", fileInfo.Size())
	if headerParams["Content-Disposition"] == "" {
		headerParams["Content-Disposition"] = fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name())
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, file, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyImportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates the Network Security Policy.
func (api *NetworkSecurityPoliciesApi) CreateNetworkSecurityPolicy(body *import1.NetworkSecurityPolicy, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateNetworkSecurityPolicy(context.Background(), &import7.CreateNetworkSecurityPolicyRequest{
		Body: body,
	}, args...)
}

// Creates the Network Security Policy.
func (api *NetworkSecurityPoliciesServiceApi) CreateNetworkSecurityPolicy(ctx context.Context, request *import7.CreateNetworkSecurityPolicyRequest, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies"

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

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) DeleteNetworkSecurityPolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteNetworkSecurityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteNetworkSecurityPolicyById(context.Background(), &import7.DeleteNetworkSecurityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesServiceApi) DeleteNetworkSecurityPolicyById(ctx context.Context, request *import7.DeleteNetworkSecurityPolicyByIdRequest, args ...map[string]interface{}) (*import1.DeleteNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/{extId}"

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

	unmarshalledResp := new(import1.DeleteNetworkSecurityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Prepares and exports all the Network Security Policies in the system. Export is achieved using two APIs. 1. POST /policies/$actions/prepare-export (ASYNC) - Serializes Network Security Policies and stores in DB. 2. GET /policies, Accept: application/octet-stream (SYNC) - Retrieves from DB and flushes to response.
func (api *NetworkSecurityPoliciesApi) ExportNetworkSecurityPolicy(body *import1.NetworkSecurityPolicyExportSpec, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyExportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExportNetworkSecurityPolicy(context.Background(), &import7.ExportNetworkSecurityPolicyRequest{
		Body: body,
	}, args...)
}

// Prepares and exports all the Network Security Policies in the system. Export is achieved using two APIs. 1. POST /policies/$actions/prepare-export (ASYNC) - Serializes Network Security Policies and stores in DB. 2. GET /policies, Accept: application/octet-stream (SYNC) - Retrieves from DB and flushes to response.
func (api *NetworkSecurityPoliciesServiceApi) ExportNetworkSecurityPolicy(ctx context.Context, request *import7.ExportNetworkSecurityPolicyRequest, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyExportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/$actions/prepare-export"

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

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyExportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetNetworkSecurityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNetworkSecurityPolicyById(context.Background(), &import7.GetNetworkSecurityPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesServiceApi) GetNetworkSecurityPolicyById(ctx context.Context, request *import7.GetNetworkSecurityPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/{extId}"

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

	unmarshalledResp := new(import1.GetNetworkSecurityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets a list of Network Security Policies.
func (api *NetworkSecurityPoliciesApi) ListNetworkSecurityPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListNetworkSecurityPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNetworkSecurityPolicies(context.Background(), &import7.ListNetworkSecurityPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Gets a list of Network Security Policies.
func (api *NetworkSecurityPoliciesServiceApi) ListNetworkSecurityPolicies(ctx context.Context, request *import7.ListNetworkSecurityPoliciesRequest, args ...map[string]interface{}) (*import1.ListNetworkSecurityPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json", "application/octet-stream"}

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

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import1.NewListNetworkSecurityPoliciesApiResponse()
			fileDetail := import1.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import5.Flag
			flags = append(flags, import5.Flag{Name: &flagName, Value: &flagValue})
			metadata := import6.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import1.ListNetworkSecurityPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the list of Network Security Policy rules with the provided policy ExtID.
func (api *NetworkSecurityPoliciesApi) ListNetworkSecurityPolicyRules(policyExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListNetworkSecurityPolicyRulesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNetworkSecurityPolicyRules(context.Background(), &import7.ListNetworkSecurityPolicyRulesRequest{
		PolicyExtId: policyExtId,
		Page_:       page_,
		Limit_:      limit_,
		Filter_:     filter_,
		Orderby_:    orderby_,
		Select_:     select_,
	}, args...)
}

// Gets the list of Network Security Policy rules with the provided policy ExtID.
func (api *NetworkSecurityPoliciesServiceApi) ListNetworkSecurityPolicyRules(ctx context.Context, request *import7.ListNetworkSecurityPolicyRulesRequest, args ...map[string]interface{}) (*import1.ListNetworkSecurityPolicyRulesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/{policyExtId}/rules"

	// verify the required parameter 'policyExtId' is set
	if nil == request.PolicyExtId {
		return nil, client.ReportError("policyExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"policyExtId"+"}", url.PathEscape(client.ParameterToString(*request.PolicyExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListNetworkSecurityPolicyRulesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) UpdateNetworkSecurityPolicyById(extId *string, body *import1.NetworkSecurityPolicy, args ...map[string]interface{}) (*import1.UpdateNetworkSecurityPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNetworkSecurityPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateNetworkSecurityPolicyById(context.Background(), &import7.UpdateNetworkSecurityPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesServiceApi) UpdateNetworkSecurityPolicyById(ctx context.Context, request *import7.UpdateNetworkSecurityPolicyByIdRequest, args ...map[string]interface{}) (*import1.UpdateNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/policies/{extId}"

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

	unmarshalledResp := new(import1.UpdateNetworkSecurityPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
