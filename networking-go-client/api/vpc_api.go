//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type VpcApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVpcApi(apiClient *client.ApiClient) *VpcApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VpcApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a VPC. Requires Prism Central >= pc.2022.9.
func (api *VpcApi) CreateVpc(body *import1.Vpc, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/vpcs"

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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete the specified VPC. Requires Prism Central >= pc.2022.9.
func (api *VpcApi) DeleteVpc(extId *string, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/vpcs/{extId}"

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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the VPC for the specified UUID. Requires Prism Central >= pc.2022.9.
func (api *VpcApi) GetVpc(extId *string, args ...map[string]interface{}) (*import1.VpcApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/vpcs/{extId}"

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
	unmarshalledResp := new(import1.VpcApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the list of existing VPCs. Requires Prism Central >= pc.2022.9.
func (api *VpcApi) ListVpcs(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.VpcListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/vpcs"

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
	unmarshalledResp := new(import1.VpcListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// Update the specified VPC. Requires Prism Central >= pc.2022.9.
func (api *VpcApi) UpdateVpc(extId *string, body *import1.Vpc, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/vpcs/{extId}"

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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
