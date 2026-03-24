package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import5 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/certificateauthenticationproviders"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CertificateAuthenticationProvidersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *CertificateAuthenticationProvidersServiceApi
}

type CertificateAuthenticationProvidersServiceApi struct {
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

	a.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(a.ApiClient)

	return a
}

func NewCertificateAuthenticationProvidersServiceApi(apiClient *client.ApiClient) *CertificateAuthenticationProvidersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CertificateAuthenticationProvidersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a certificate-based authentication provider.
func (api *CertificateAuthenticationProvidersApi) CreateCertAuthProvider(clientCaChain *string, caCertFileName *string, isCertAuthEnabled *bool, name *string, isCacEnabled *bool, dirSvcExtID *string, certRevocationInfo *import4.CertRevocationInfo, createdBy *string, tenantId *string, createdTime *time.Time, links *[]import3.ApiLink, lastUpdatedTime *time.Time, extId *string, args ...map[string]interface{}) (*import4.CreateCertAuthProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCertAuthProvider(context.Background(), &import5.CreateCertAuthProviderRequest{
		ClientCaChain:      clientCaChain,
		CaCertFileName:     caCertFileName,
		IsCertAuthEnabled:  isCertAuthEnabled,
		Name:               name,
		IsCacEnabled:       isCacEnabled,
		DirSvcExtID:        dirSvcExtID,
		CertRevocationInfo: certRevocationInfo,
		CreatedBy:          createdBy,
		TenantId:           tenantId,
		CreatedTime:        createdTime,
		Links:              links,
		LastUpdatedTime:    lastUpdatedTime,
		ExtId:              extId,
	}, args...)
}

// Creates a certificate-based authentication provider.
func (api *CertificateAuthenticationProvidersServiceApi) CreateCertAuthProvider(ctx context.Context, request *import5.CreateCertAuthProviderRequest, args ...map[string]interface{}) (*import4.CreateCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/cert-auth-providers"

	// verify the required parameter 'clientCaChain' is set
	if nil == request.ClientCaChain {
		return nil, client.ReportError("clientCaChain is required and must be specified")
	}
	// verify the required parameter 'caCertFileName' is set
	if nil == request.CaCertFileName {
		return nil, client.ReportError("caCertFileName is required and must be specified")
	}
	// verify the required parameter 'isCertAuthEnabled' is set
	if nil == request.IsCertAuthEnabled {
		return nil, client.ReportError("isCertAuthEnabled is required and must be specified")
	}
	// verify the required parameter 'name' is set
	if nil == request.Name {
		return nil, client.ReportError("name is required and must be specified")
	}
	// verify the required parameter 'isCacEnabled' is set
	if nil == request.IsCacEnabled {
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
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	// Form Params
	formParams.Add("clientCaChain", client.ParameterToString(*request.ClientCaChain, ""))
	if request.DirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*request.DirSvcExtID, ""))
	}
	if request.CertRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*request.CertRevocationInfo, ""))
	}
	formParams.Add("caCertFileName", client.ParameterToString(*request.CaCertFileName, ""))
	formParams.Add("isCertAuthEnabled", client.ParameterToString(*request.IsCertAuthEnabled, ""))
	if request.CreatedBy != nil {
		formParams.Add("createdBy", client.ParameterToString(*request.CreatedBy, ""))
	}
	if request.TenantId != nil {
		formParams.Add("tenantId", client.ParameterToString(*request.TenantId, ""))
	}
	formParams.Add("name", client.ParameterToString(*request.Name, ""))
	if request.CreatedTime != nil {
		formParams.Add("createdTime", client.ParameterToString(*request.CreatedTime, ""))
	}
	if request.Links != nil {
		formParams.Add("links", client.ParameterToString(*request.Links, "multi"))
	}
	formParams.Add("isCacEnabled", client.ParameterToString(*request.IsCacEnabled, ""))
	if request.LastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*request.LastUpdatedTime, ""))
	}
	if request.ExtId != nil {
		formParams.Add("extId", client.ParameterToString(*request.ExtId, ""))
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.CreateCertAuthProviderApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a certificate-based authentication provider configuration for the given UUID.
func (api *CertificateAuthenticationProvidersApi) DeleteCertAuthProviderById(extId *string, args ...map[string]interface{}) (*import4.DeleteCertAuthProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteCertAuthProviderById(context.Background(), &import5.DeleteCertAuthProviderByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a certificate-based authentication provider configuration for the given UUID.
func (api *CertificateAuthenticationProvidersServiceApi) DeleteCertAuthProviderById(ctx context.Context, request *import5.DeleteCertAuthProviderByIdRequest, args ...map[string]interface{}) (*import4.DeleteCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/cert-auth-providers/{extId}"

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

	unmarshalledResp := new(import4.DeleteCertAuthProviderApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a certificate-based authentication provider by its UUID.
func (api *CertificateAuthenticationProvidersApi) GetCertAuthProviderById(extId *string, args ...map[string]interface{}) (*import4.GetCertAuthProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCertAuthProviderById(context.Background(), &import5.GetCertAuthProviderByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a certificate-based authentication provider by its UUID.
func (api *CertificateAuthenticationProvidersServiceApi) GetCertAuthProviderById(ctx context.Context, request *import5.GetCertAuthProviderByIdRequest, args ...map[string]interface{}) (*import4.GetCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/cert-auth-providers/{extId}"

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

	unmarshalledResp := new(import4.GetCertAuthProviderApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all configured certificate-based authentication providers.
func (api *CertificateAuthenticationProvidersApi) ListCertAuthProviders(page_ *int, limit_ *int, args ...map[string]interface{}) (*import4.ListCertAuthProvidersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCertAuthProviders(context.Background(), &import5.ListCertAuthProvidersRequest{
		Page_:  page_,
		Limit_: limit_,
	}, args...)
}

// Lists all configured certificate-based authentication providers.
func (api *CertificateAuthenticationProvidersServiceApi) ListCertAuthProviders(ctx context.Context, request *import5.ListCertAuthProvidersRequest, args ...map[string]interface{}) (*import4.ListCertAuthProvidersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/cert-auth-providers"

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

	unmarshalledResp := new(import4.ListCertAuthProvidersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a certificate-based authentication provider configuration.
func (api *CertificateAuthenticationProvidersApi) UpdateCertAuthProviderById(extId *string, clientCaChain *string, caCertFileName *string, isCertAuthEnabled *bool, name *string, isCacEnabled *bool, dirSvcExtID *string, certRevocationInfo *import4.CertRevocationInfo, createdBy *string, tenantId *string, createdTime *time.Time, links *[]import3.ApiLink, lastUpdatedTime *time.Time, extId2 *string, args ...map[string]interface{}) (*import4.UpdateCertAuthProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCertificateAuthenticationProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateCertAuthProviderById(context.Background(), &import5.UpdateCertAuthProviderByIdRequest{
		ExtId:              extId,
		ClientCaChain:      clientCaChain,
		CaCertFileName:     caCertFileName,
		IsCertAuthEnabled:  isCertAuthEnabled,
		Name:               name,
		IsCacEnabled:       isCacEnabled,
		DirSvcExtID:        dirSvcExtID,
		CertRevocationInfo: certRevocationInfo,
		CreatedBy:          createdBy,
		TenantId:           tenantId,
		CreatedTime:        createdTime,
		Links:              links,
		LastUpdatedTime:    lastUpdatedTime,
		ExtId2:             extId2,
	}, args...)
}

// Updates a certificate-based authentication provider configuration.
func (api *CertificateAuthenticationProvidersServiceApi) UpdateCertAuthProviderById(ctx context.Context, request *import5.UpdateCertAuthProviderByIdRequest, args ...map[string]interface{}) (*import4.UpdateCertAuthProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b2/authn/cert-auth-providers/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'clientCaChain' is set
	if nil == request.ClientCaChain {
		return nil, client.ReportError("clientCaChain is required and must be specified")
	}
	// verify the required parameter 'caCertFileName' is set
	if nil == request.CaCertFileName {
		return nil, client.ReportError("caCertFileName is required and must be specified")
	}
	// verify the required parameter 'isCertAuthEnabled' is set
	if nil == request.IsCertAuthEnabled {
		return nil, client.ReportError("isCertAuthEnabled is required and must be specified")
	}
	// verify the required parameter 'name' is set
	if nil == request.Name {
		return nil, client.ReportError("name is required and must be specified")
	}
	// verify the required parameter 'isCacEnabled' is set
	if nil == request.IsCacEnabled {
		return nil, client.ReportError("isCacEnabled is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	// Form Params
	formParams.Add("clientCaChain", client.ParameterToString(*request.ClientCaChain, ""))
	if request.DirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*request.DirSvcExtID, ""))
	}
	if request.CertRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*request.CertRevocationInfo, ""))
	}
	formParams.Add("caCertFileName", client.ParameterToString(*request.CaCertFileName, ""))
	formParams.Add("isCertAuthEnabled", client.ParameterToString(*request.IsCertAuthEnabled, ""))
	if request.CreatedBy != nil {
		formParams.Add("createdBy", client.ParameterToString(*request.CreatedBy, ""))
	}
	if request.TenantId != nil {
		formParams.Add("tenantId", client.ParameterToString(*request.TenantId, ""))
	}
	formParams.Add("name", client.ParameterToString(*request.Name, ""))
	if request.CreatedTime != nil {
		formParams.Add("createdTime", client.ParameterToString(*request.CreatedTime, ""))
	}
	if request.Links != nil {
		formParams.Add("links", client.ParameterToString(*request.Links, "multi"))
	}
	formParams.Add("isCacEnabled", client.ParameterToString(*request.IsCacEnabled, ""))
	if request.LastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*request.LastUpdatedTime, ""))
	}
	if request.ExtId2 != nil {
		formParams.Add("extId", client.ParameterToString(*request.ExtId2, ""))
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.UpdateCertAuthProviderApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
