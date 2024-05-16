//Api classes for microseg's golang SDK
package api

import (
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type NetworkSecurityPoliciesApi struct {
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

	return a
}

// Imports all the Network Security Policies specified by the data file. The Policies Import (POST /policies/$actions/import) includes a `dry_run` option, providing a summary detailing the potential impact of importing the data file. Set `NTNX-Dry-Run` to `True` in the Request headers to achieve this.
func (api *NetworkSecurityPoliciesApi) ApplyNetworkSecurityPolicyImport(path *string, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyImportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/$actions/import"

	// verify the required parameter 'path' is set
	if nil == path {
		return nil, client.ReportError("path is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/octet-stream"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	file, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	headerParams["Content-Length"] = fmt.Sprintf("%d", fileInfo.Size())
	if headerParams["Content-Disposition"] == "" {
		headerParams["Content-Disposition"] = fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name())
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, file, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyImportApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates the Network Security Policy.
func (api *NetworkSecurityPoliciesApi) CreateNetworkSecurityPolicy(body *import1.NetworkSecurityPolicy, args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) DeleteNetworkSecurityPolicyById(extId *string, args ...map[string]interface{}) (*import1.DeleteNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteNetworkSecurityPolicyApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Prepares and exports all the Network Security Policies in the system. Export is achieved using two APIs. 1. POST /policies/$actions/prepare-export (ASYNC) - Serializes Network Security Policies and stores in DB. 2. GET /policies, Accept: application/octet-stream (SYNC) - Retrieves from DB and flushes to response.
func (api *NetworkSecurityPoliciesApi) ExportNetworkSecurityPolicy(args ...map[string]interface{}) (*import1.CreateNetworkSecurityPolicyExportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/$actions/prepare-export"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateNetworkSecurityPolicyExportApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.GetNetworkSecurityPolicyApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets a list of Network Security Policies.
func (api *NetworkSecurityPoliciesApi) ListNetworkSecurityPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListNetworkSecurityPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json", "application/octet-stream"}

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	isFileOperation := headerParams["Accept"] == "application/octet-stream" || (len(accepts) == 1 && accepts[0] == "application/octet-stream")
	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	if isFileOperation {
		filePath := responseBody.(*string)
		response := import1.NewListNetworkSecurityPoliciesApiResponse()
		fileDetail := import1.NewFileDetail()
		fileDetail.Path = filePath
		err = response.SetData(*fileDetail)
		if err != nil {
			return nil, err
		}

		return response, err
	} else {
		unmarshalledResp := new(import1.ListNetworkSecurityPoliciesApiResponse)
		json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
		return unmarshalledResp, err
	}
}

// Gets the list of Network Security Policy rules with the provided policy ExtID.
func (api *NetworkSecurityPoliciesApi) ListNetworkSecurityPolicyRules(extId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.ListNetworkSecurityPolicyRulesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/{extId}/rules"

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

	// Query Params
	if page_ != nil {

		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {

		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.ListNetworkSecurityPolicyRulesApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Network Security Policy with the provided ExtID.
func (api *NetworkSecurityPoliciesApi) UpdateNetworkSecurityPolicyById(extId *string, body *import1.NetworkSecurityPolicy, args ...map[string]interface{}) (*import1.UpdateNetworkSecurityPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.0.b1/config/policies/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateNetworkSecurityPolicyApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
