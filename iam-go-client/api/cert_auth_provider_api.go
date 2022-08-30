//Api classes for iam's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	"strings"
	import2 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	"time"
	"encoding/json"
	"net/http"
    "net/url"
)

type CertAuthProviderApi struct {
  ApiClient *client.ApiClient
}

func NewCertAuthProviderApi(apiClient *client.ApiClient) *CertAuthProviderApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CertAuthProviderApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create certificate based authentication provider
    Create a certificate based authentication provider

    parameters:-
    -> clientCaChain (string) (required)
    -> dirSvcExtID (string) (optional)
    -> certAuthEnabled (bool) (required)
    -> cacEnabled (bool) (required)
    -> certRevocationInfo (iam.v4.authn.CertRevocationInfo) (optional)
    -> caCertFileName (string) (required)
    -> createdBy (string) (optional)
    -> tenantId (string) (optional)
    -> name (string) (required)
    -> createdTime (time.Time) (optional)
    -> links ([]common.v1.response.ApiLink) (optional)
    -> lastUpdatedTime (time.Time) (optional)
    -> extId (string) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.CreateCertAuthProviderApiResponse, error)
*/
func (api *CertAuthProviderApi) HandleCreateCertAuthProvider(clientCaChain *string, dirSvcExtID *string, certAuthEnabled *bool, cacEnabled *bool, certRevocationInfo *import3.CertRevocationInfo, caCertFileName *string, createdBy *string, tenantId *string, name *string, createdTime *time.Time, links *[]import2.ApiLink, lastUpdatedTime *time.Time, extId *string, args ...map[string]interface{}) (*import3.CreateCertAuthProviderApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/cert-auth-providers"

    // verify the required parameter 'clientCaChain' is set
	if nil == clientCaChain {
		return nil, client.ReportError("clientCaChain is required and must be specified")
	}
    // verify the required parameter 'certAuthEnabled' is set
	if nil == certAuthEnabled {
		return nil, client.ReportError("certAuthEnabled is required and must be specified")
	}
    // verify the required parameter 'cacEnabled' is set
	if nil == cacEnabled {
		return nil, client.ReportError("cacEnabled is required and must be specified")
	}
    // verify the required parameter 'caCertFileName' is set
	if nil == caCertFileName {
		return nil, client.ReportError("caCertFileName is required and must be specified")
	}
    // verify the required parameter 'name' is set
	if nil == name {
		return nil, client.ReportError("name is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"multipart/form-data"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    // Form Params
	formParams.Add("clientCaChain", client.ParameterToString(*clientCaChain, ""))
	if dirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*dirSvcExtID, ""))
	}
	formParams.Add("certAuthEnabled", client.ParameterToString(*certAuthEnabled, ""))
	formParams.Add("cacEnabled", client.ParameterToString(*cacEnabled, ""))
	if certRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*certRevocationInfo, ""))
	}
	formParams.Add("caCertFileName", client.ParameterToString(*caCertFileName, ""))
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
	if lastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*lastUpdatedTime, ""))
	}
	if extId != nil {
		formParams.Add("extId", client.ParameterToString(*extId, ""))
	}

    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.CreateCertAuthProviderApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get certificate based authentication provider
    Get a certificate based authentication provider by its UUID

    parameters:-
    -> extId (string) (required) : External identifier of the certificate based authentication provider
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.GetCertAuthProviderApiResponse, error)
*/
func (api *CertAuthProviderApi) HandleGetCertAuthProvider(extId *string, args ...map[string]interface{}) (*import3.GetCertAuthProviderApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/cert-auth-providers/{extId}"

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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.GetCertAuthProviderApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List certificate based authentication provider(s)
    Get a list of all configured certificate based authentication providers

    parameters:-
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - createdBy - extId - name 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - createdTime - lastUpdatedTime - name 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - caCertFileName - cacEnabled - certAuthEnabled - certRevocationInfo - clientCaChain - createdBy - createdTime - dirSvcExtID - extId - lastUpdatedTime - links - name - tenantId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.ListCertAuthProviderApiResponse, error)
*/
func (api *CertAuthProviderApi) HandleListCertAuthProviders(filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListCertAuthProviderApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/cert-auth-providers"


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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.ListCertAuthProviderApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update certificate based authentication provider
    Update a certificate based authentication provider configuration

    parameters:-
    -> extId (string) (required) : External identifier of the certificate based authentication provider
    -> clientCaChain (string) (optional)
    -> dirSvcExtID (string) (optional)
    -> certAuthEnabled (bool) (optional)
    -> cacEnabled (bool) (optional)
    -> certRevocationInfo (iam.v4.authn.CertRevocationInfo) (optional)
    -> caCertFileName (string) (optional)
    -> createdBy (string) (optional)
    -> tenantId (string) (optional)
    -> name (string) (optional)
    -> createdTime (time.Time) (optional)
    -> links ([]common.v1.response.ApiLink) (optional)
    -> lastUpdatedTime (time.Time) (optional)
    -> extId2 (string) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.UpdateCertAuthProviderApiResponse, error)
*/
func (api *CertAuthProviderApi) HandleUpdateCertAuthProvider(extId *string, clientCaChain *string, dirSvcExtID *string, certAuthEnabled *bool, cacEnabled *bool, certRevocationInfo *import3.CertRevocationInfo, caCertFileName *string, createdBy *string, tenantId *string, name *string, createdTime *time.Time, links *[]import2.ApiLink, lastUpdatedTime *time.Time, extId2 *string, args ...map[string]interface{}) (*import3.UpdateCertAuthProviderApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/cert-auth-providers/{extId}"

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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    // Form Params
	if clientCaChain != nil {
		formParams.Add("clientCaChain", client.ParameterToString(*clientCaChain, ""))
	}
	if dirSvcExtID != nil {
		formParams.Add("dirSvcExtID", client.ParameterToString(*dirSvcExtID, ""))
	}
	if certAuthEnabled != nil {
		formParams.Add("certAuthEnabled", client.ParameterToString(*certAuthEnabled, ""))
	}
	if cacEnabled != nil {
		formParams.Add("cacEnabled", client.ParameterToString(*cacEnabled, ""))
	}
	if certRevocationInfo != nil {
		formParams.Add("certRevocationInfo", client.ParameterToString(*certRevocationInfo, ""))
	}
	if caCertFileName != nil {
		formParams.Add("caCertFileName", client.ParameterToString(*caCertFileName, ""))
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
	if lastUpdatedTime != nil {
		formParams.Add("lastUpdatedTime", client.ParameterToString(*lastUpdatedTime, ""))
	}
	if extId2 != nil {
		formParams.Add("extId", client.ParameterToString(*extId2, ""))
	}

    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.UpdateCertAuthProviderApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

