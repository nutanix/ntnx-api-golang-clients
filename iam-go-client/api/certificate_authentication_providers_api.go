//Api classes for iam's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CertificateAuthenticationProvidersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCertificateAuthenticationProvidersApi(apiClient *client.ApiClient) *CertificateAuthenticationProvidersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CertificateAuthenticationProvidersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a Certificate based Authentication provider.
func (api *CertificateAuthenticationProvidersApi) CreateCertAuthProvider(clientCaChain *string, caCertFileName *string, isCertAuthEnabled *bool, name *string, isCacEnabled *bool, dirSvcExtID *string, certRevocationInfo *import1.CertRevocationInfo, createdBy *string, tenantId *string, createdTime *time.Time, links *[]import3.ApiLink, lastUpdatedTime *time.Time, extId *string, args ...map[string]interface{}) (*import1.CreateCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.b2/authn/cert-auth-providers"

	// verify the required parameter 'clientCaChain' is set
	if nil == clientCaChain {
		return nil, client.ReportError("clientCaChain is required and must be specified")
	}
	// verify the required parameter 'caCertFileName' is set
	if nil == caCertFileName {
		return nil, client.ReportError("caCertFileName is required and must be specified")
	}
	// verify the required parameter 'isCertAuthEnabled' is set
	if nil == isCertAuthEnabled {
		return nil, client.ReportError("isCertAuthEnabled is required and must be specified")
	}
	// verify the required parameter 'name' is set
	if nil == name {
		return nil, client.ReportError("name is required and must be specified")
	}
	// verify the required parameter 'isCacEnabled' is set
	if nil == isCacEnabled {
		return nil, client.ReportError("isCacEnabled is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"multipart/form-data"}

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

	// Form Params
	formParams.Add("clientCaChain", client.ParameterToString(*clientCaChain, ""))
	if dirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*dirSvcExtID, ""))
	}
	if certRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*certRevocationInfo, ""))
	}
	formParams.Add("caCertFileName", client.ParameterToString(*caCertFileName, ""))
	formParams.Add("isCertAuthEnabled", client.ParameterToString(*isCertAuthEnabled, ""))
	if createdBy != nil {
		formParams.Add("createdBy", client.ParameterToString(*createdBy, ""))
	}
	if tenantId != nil {
		formParams.Add("tenantId", client.ParameterToString(*tenantId, ""))
	}
	formParams.Add("name", client.ParameterToString(*name, ""))
	if createdTime != nil {
		formParams.Add("createdTime", client.ParameterToString(*createdTime, ""))
	}
	if links != nil {
		formParams.Add("links", client.ParameterToString(*links, "multi"))
	}
	formParams.Add("isCacEnabled", client.ParameterToString(*isCacEnabled, ""))
	if lastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*lastUpdatedTime, ""))
	}
	if extId != nil {
		formParams.Add("extId", client.ParameterToString(*extId, ""))
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateCertAuthProviderApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Delete a Certificate based Authentication provider configuration by its UUID.
func (api *CertificateAuthenticationProvidersApi) DeleteCertAuthProviderById(extId *string, args ...map[string]interface{}) (*import1.DeleteCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.b2/authn/cert-auth-providers/{extId}"

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

	unmarshalledResp := new(import1.DeleteCertAuthProviderApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a Certificate based Authentication provider by its UUID.
func (api *CertificateAuthenticationProvidersApi) GetCertAuthProviderById(extId *string, args ...map[string]interface{}) (*import1.GetCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.b2/authn/cert-auth-providers/{extId}"

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

	unmarshalledResp := new(import1.GetCertAuthProviderApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List all Configured Certificate based Authentication providers.
func (api *CertificateAuthenticationProvidersApi) ListCertAuthProviders(filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCertAuthProvidersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.b2/authn/cert-auth-providers"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
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

	unmarshalledResp := new(import1.ListCertAuthProvidersApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update a Certificate based Authentication provider configuration.
func (api *CertificateAuthenticationProvidersApi) UpdateCertAuthProviderById(extId *string, clientCaChain *string, dirSvcExtID *string, certRevocationInfo *import1.CertRevocationInfo, caCertFileName *string, isCertAuthEnabled *bool, createdBy *string, tenantId *string, name *string, createdTime *time.Time, links *[]import3.ApiLink, isCacEnabled *bool, lastUpdatedTime *time.Time, extId2 *string, args ...map[string]interface{}) (*import1.UpdateCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.b2/authn/cert-auth-providers/{extId}"

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
	contentTypes := []string{"multipart/form-data"}

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

	// Form Params
	if clientCaChain != nil {
		formParams.Add("clientCaChain", client.ParameterToString(*clientCaChain, ""))
	}
	if dirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*dirSvcExtID, ""))
	}
	if certRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*certRevocationInfo, ""))
	}
	if caCertFileName != nil {
		formParams.Add("caCertFileName", client.ParameterToString(*caCertFileName, ""))
	}
	if isCertAuthEnabled != nil {
		formParams.Add("isCertAuthEnabled", client.ParameterToString(*isCertAuthEnabled, ""))
	}
	if createdBy != nil {
		formParams.Add("createdBy", client.ParameterToString(*createdBy, ""))
	}
	if tenantId != nil {
		formParams.Add("tenantId", client.ParameterToString(*tenantId, ""))
	}
	if name != nil {
		formParams.Add("name", client.ParameterToString(*name, ""))
	}
	if createdTime != nil {
		formParams.Add("createdTime", client.ParameterToString(*createdTime, ""))
	}
	if links != nil {
		formParams.Add("links", client.ParameterToString(*links, "multi"))
	}
	if isCacEnabled != nil {
		formParams.Add("isCacEnabled", client.ParameterToString(*isCacEnabled, ""))
	}
	if lastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*lastUpdatedTime, ""))
	}
	if extId2 != nil {
		formParams.Add("extId", client.ParameterToString(*extId2, ""))
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateCertAuthProviderApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
