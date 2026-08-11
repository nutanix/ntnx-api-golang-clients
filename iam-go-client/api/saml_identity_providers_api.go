package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import13 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	import14 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/request/samlidentityproviders"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

type SAMLIdentityProvidersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SAMLIdentityProvidersServiceApi
}

type SAMLIdentityProvidersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSAMLIdentityProvidersApi(apiClient *client.ApiClient) *SAMLIdentityProvidersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SAMLIdentityProvidersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSAMLIdentityProvidersServiceApi(a.ApiClient)

	return a
}

func NewSAMLIdentityProvidersServiceApi(apiClient *client.ApiClient) *SAMLIdentityProvidersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SAMLIdentityProvidersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a SAML identity provider from the given request payload.
func (api *SAMLIdentityProvidersApi) CreateSamlIdentityProvider(body *import4.SamlIdentityProvider, args ...map[string]interface{}) (*import4.CreateSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateSamlIdentityProvider(context.Background(), &import14.CreateSamlIdentityProviderRequest{
		Body: body,
	}, args...)
}

// Create a SAML identity provider from the given request payload.
func (api *SAMLIdentityProvidersServiceApi) CreateSamlIdentityProvider(ctx context.Context, request *import14.CreateSamlIdentityProviderRequest, args ...map[string]interface{}) (*import4.CreateSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.CreateSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersApi) DeleteSamlIdentityProviderById(extId *string, args ...map[string]interface{}) (*import4.DeleteSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteSamlIdentityProviderById(context.Background(), &import14.DeleteSamlIdentityProviderByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersServiceApi) DeleteSamlIdentityProviderById(ctx context.Context, request *import14.DeleteSamlIdentityProviderByIdRequest, args ...map[string]interface{}) (*import4.DeleteSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.DeleteSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Fetches a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersApi) GetSamlIdentityProviderById(extId *string, args ...map[string]interface{}) (*import4.GetSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSamlIdentityProviderById(context.Background(), &import14.GetSamlIdentityProviderByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersServiceApi) GetSamlIdentityProviderById(ctx context.Context, request *import14.GetSamlIdentityProviderByIdRequest, args ...map[string]interface{}) (*import4.GetSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.GetSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Downloads SP-Metadata for SAML identity provider.
func (api *SAMLIdentityProvidersApi) GetSamlIdpSpMetadataById(extId *string, args ...map[string]interface{}) (*import4.GetSamlIdpSpMetadataApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSamlIdpSpMetadataById(context.Background(), &import14.GetSamlIdpSpMetadataByIdRequest{
		ExtId: extId,
	}, args...)
}

// Downloads SP-Metadata for SAML identity provider.
func (api *SAMLIdentityProvidersServiceApi) GetSamlIdpSpMetadataById(ctx context.Context, request *import14.GetSamlIdpSpMetadataByIdRequest, args ...map[string]interface{}) (*import4.GetSamlIdpSpMetadataApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers/{extId}/sp-metadata"

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
	accepts := []string{"text/xml", "application/json"}

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	textMediaTypes := []string{"text/event-stream", "text/html", "text/xml", "text/csv", "text/javascript", "text/markdown", "text/vcard", "text/plain"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		responseContentType, _, _ := mime.ParseMediaType(httpResponse.Header.Get("Content-Type"))
		if httpResponse.Header != nil && api.ApiClient.Contains(textMediaTypes, responseContentType) {
			response := import4.NewGetSamlIdpSpMetadataApiResponse()

			flagName := "hasError"
			flagValue := false
			var flags []import13.Flag
			flags = append(flags, import13.Flag{Name: &flagName, Value: &flagValue})
			metadata := import3.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			responseBody, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				return nil, err
			}
			httpResponse.Body.Close()
			err = response.SetData(string(responseBody))
			if err != nil {
				return nil, err
			}
			return response, nil
		}
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.GetSamlIdpSpMetadataApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Downloads SP-Metadata for SAML identity provider.
func (api *SAMLIdentityProvidersApi) GetSamlSpMetadata(args ...map[string]interface{}) (*import4.GetSamlSpMetadataApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSamlSpMetadata(context.Background(), &import14.GetSamlSpMetadataRequest{}, args...)
}

// Downloads SP-Metadata for SAML identity provider.
func (api *SAMLIdentityProvidersServiceApi) GetSamlSpMetadata(ctx context.Context, request *import14.GetSamlSpMetadataRequest, args ...map[string]interface{}) (*import4.GetSamlSpMetadataApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-sp-metadata"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"text/xml", "application/json"}

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	textMediaTypes := []string{"text/event-stream", "text/html", "text/xml", "text/csv", "text/javascript", "text/markdown", "text/vcard", "text/plain"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		responseContentType, _, _ := mime.ParseMediaType(httpResponse.Header.Get("Content-Type"))
		if httpResponse.Header != nil && api.ApiClient.Contains(textMediaTypes, responseContentType) {
			response := import4.NewGetSamlSpMetadataApiResponse()

			responseBody, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				return nil, err
			}
			httpResponse.Body.Close()
			err = response.SetData(string(responseBody))
			if err != nil {
				return nil, err
			}
			return response, nil
		}
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.GetSamlSpMetadataApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists all SAML identity providers configured in the system.
func (api *SAMLIdentityProvidersApi) ListSamlIdentityProviders(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListSamlIdentityProvidersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSamlIdentityProviders(context.Background(), &import14.ListSamlIdentityProvidersRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists all SAML identity providers configured in the system.
func (api *SAMLIdentityProvidersServiceApi) ListSamlIdentityProviders(ctx context.Context, request *import14.ListSamlIdentityProvidersRequest, args ...map[string]interface{}) (*import4.ListSamlIdentityProvidersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.ListSamlIdentityProvidersApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Shares a SAML identity provider with all projects.
func (api *SAMLIdentityProvidersApi) ShareAllSamlIdentityProvider(extId *string, args ...map[string]interface{}) (*import4.ShareAllSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShareAllSamlIdentityProvider(context.Background(), &import14.ShareAllSamlIdentityProviderRequest{
		ExtId: extId,
	}, args...)
}

// Shares a SAML identity provider with all projects.
func (api *SAMLIdentityProvidersServiceApi) ShareAllSamlIdentityProvider(ctx context.Context, request *import14.ShareAllSamlIdentityProviderRequest, args ...map[string]interface{}) (*import4.ShareAllSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/config/saml-identity-provider/{extId}/$actions/share-all"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.ShareAllSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Shares a SAML identity provider with specified project.
func (api *SAMLIdentityProvidersApi) ShareSamlIdentityProvider(extId *string, body *import4.SamlIdentityProviderShareRequest, args ...map[string]interface{}) (*import4.ShareSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ShareSamlIdentityProvider(context.Background(), &import14.ShareSamlIdentityProviderRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Shares a SAML identity provider with specified project.
func (api *SAMLIdentityProvidersServiceApi) ShareSamlIdentityProvider(ctx context.Context, request *import14.ShareSamlIdentityProviderRequest, args ...map[string]interface{}) (*import4.ShareSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/config/saml-identity-provider/{extId}/$actions/share"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.ShareSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Unshare SAML identity provider from all projects, making it unavailable for authentication across all projects.
func (api *SAMLIdentityProvidersApi) UnshareAllSamlIdentityProvider(extId *string, args ...map[string]interface{}) (*import4.UnshareAllSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnshareAllSamlIdentityProvider(context.Background(), &import14.UnshareAllSamlIdentityProviderRequest{
		ExtId: extId,
	}, args...)
}

// Unshare SAML identity provider from all projects, making it unavailable for authentication across all projects.
func (api *SAMLIdentityProvidersServiceApi) UnshareAllSamlIdentityProvider(ctx context.Context, request *import14.UnshareAllSamlIdentityProviderRequest, args ...map[string]interface{}) (*import4.UnshareAllSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/config/saml-identity-provider/{extId}/$actions/unshare-all"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.UnshareAllSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Unshares a SAML identity provider from a specific project.
func (api *SAMLIdentityProvidersApi) UnshareSamlIdentityProvider(extId *string, body *import4.SamlIdentityProviderUnshareRequest, args ...map[string]interface{}) (*import4.UnshareSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnshareSamlIdentityProvider(context.Background(), &import14.UnshareSamlIdentityProviderRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Unshares a SAML identity provider from a specific project.
func (api *SAMLIdentityProvidersServiceApi) UnshareSamlIdentityProvider(ctx context.Context, request *import14.UnshareSamlIdentityProviderRequest, args ...map[string]interface{}) (*import4.UnshareSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/config/saml-identity-provider/{extId}/$actions/unshare"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.UnshareSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersApi) UpdateSamlIdentityProviderById(extId *string, body *import4.SamlIdentityProvider, args ...map[string]interface{}) (*import4.UpdateSamlIdentityProviderApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSAMLIdentityProvidersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSamlIdentityProviderById(context.Background(), &import14.UpdateSamlIdentityProviderByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a SAML identity provider based on its external identifier.
func (api *SAMLIdentityProvidersServiceApi) UpdateSamlIdentityProviderById(ctx context.Context, request *import14.UpdateSamlIdentityProviderByIdRequest, args ...map[string]interface{}) (*import4.UpdateSamlIdentityProviderApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.1.b3/authn/saml-identity-providers/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import4.UpdateSamlIdentityProviderApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
