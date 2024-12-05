package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type QuotaPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewQuotaPoliciesApi(apiClient *client.ApiClient) *QuotaPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &QuotaPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a quota policy by using the provided request body.  Specify a valid identifier of the mount target for which quota policy needs to be created. Users can specify `enforcementType`, `principalName`, `principalType`, `sizeInGb` as part of the request body.  Use dry-run workflow with same payload to validate the request. Normal workflow won't validate `principalName` and `principalType` before task creation.
func (api *QuotaPoliciesApi) CreateQuotaPolicy(fileServerExtId *string, mountTargetExtId *string, body *import3.QuotaPolicy, dryrun_ *bool, args ...map[string]interface{}) (*import3.CreateQuotaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*dryrun_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.CreateQuotaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a quota policy with the provided external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and the mount target (`mountTargetExtId`) to which the quota policy belongs and a valid external identifier (`extId`) of the quota policy to be deleted.
func (api *QuotaPoliciesApi) DeleteQuotaPolicyById(fileServerExtId *string, mountTargetExtId *string, extId *string, args ...map[string]interface{}) (*import3.DeleteQuotaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
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

	unmarshalledResp := new(import3.DeleteQuotaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get details of email configuration for the quota policy. User can have single email configuration for file server.
func (api *QuotaPoliciesApi) GetEmailConfig(fileServerExtId *string, args ...map[string]interface{}) (*import3.GetEmailConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/email-config"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

	unmarshalledResp := new(import3.GetEmailConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a quota policy with the given external identifier.  Specify a valid identifier of the file server (`fileServerExtId`) and the mount target (`mountTargetExtId`) to which the quota policy belongs and a valid external identifier (`extId`) of the quota policy to be fetched.
func (api *QuotaPoliciesApi) GetQuotaPolicyById(fileServerExtId *string, mountTargetExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetQuotaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
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

	unmarshalledResp := new(import3.GetQuotaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a paginated list of quota policies for specified mount target.
func (api *QuotaPoliciesApi) ListQuotaPolicies(fileServerExtId *string, mountTargetExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListQuotaPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListQuotaPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the quota policy email configuration.  Specify (`content`) and (`subject`) to update the quota policy email configuration.
func (api *QuotaPoliciesApi) UpdateEmailConfig(fileServerExtId *string, body *import3.EmailConfig, args ...map[string]interface{}) (*import3.UpdateEmailConfigApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/email-config"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

	unmarshalledResp := new(import3.UpdateEmailConfigApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates quota policy with the provided parameters.  Specify a valid identifier of the file server (`fileServerExtId`) and the mount target (`mountTargetExtId`) to which the quota policy belongs and a valid external identifier (`extId`) of the quota policy to be updated.   It is always recommended to perform a GET request on a resource before requesting an UPDATE call.  `principalType` field is non-updatable field.  Use dry-run workflow with same payload to validate the request. Normal workflow won't validate `principalName` and `principalType` before task creation.
func (api *QuotaPoliciesApi) UpdateQuotaPolicyById(fileServerExtId *string, mountTargetExtId *string, extId *string, body *import3.QuotaPolicy, dryrun_ *bool, args ...map[string]interface{}) (*import3.UpdateQuotaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/mount-targets/{mountTargetExtId}/quota-policies/{extId}"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*dryrun_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UpdateQuotaPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
