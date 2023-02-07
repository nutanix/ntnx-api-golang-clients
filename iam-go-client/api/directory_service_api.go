//Api classes for iam's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
	"net/http"
	"net/url"
	"strings"
)

type DirectoryServiceApi struct {
	ApiClient *client.ApiClient
}

func NewDirectoryServiceApi(apiClient *client.ApiClient) *DirectoryServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DirectoryServiceApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Check directory service connection
  Check connection to the directory service

  parameters:-
  -> body (iam.v4.authn.DirectoryServiceConnectionRequest) (required) : Check connection to the directory service
  -> extId (string) (required) : External identifier of the directory service
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.ConnectionDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) ConnectionStatusDirectoryService(body *import3.DirectoryServiceConnectionRequest, extId *string, args ...map[string]interface{}) (*import3.ConnectionDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services/{extId}/connection-status"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
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
	contentTypes := []string{"application/json"}

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.ConnectionDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Create directory service
  Create a directory service

  parameters:-
  -> body (iam.v4.authn.DirectoryService) (required) : Create a directory service
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.CreateDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) CreateDirectoryService(body *import3.DirectoryService, args ...map[string]interface{}) (*import3.CreateDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services"

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

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.CreateDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  List directory service(s)
  List all directory service(s)

  parameters:-
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - createdBy - domainName - extId - name
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - createdTime - domainName - lastUpdatedTime - name
  -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - createdBy - createdTime - directoryType - domainName - extId - groupSearchType - lastUpdatedTime - links - name - openLdapConfiguration - serviceAccount - tenantId - url - whiteListedGroups
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.ListDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) GetDirectoryServiceList(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services"

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

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.ListDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get directory service
  View a directory service

  parameters:-
  -> extId (string) (required) : External identifier of the directory service
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.GetDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) GetDirectoryServicebyId(extId *string, args ...map[string]interface{}) (*import3.GetDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services/{extId}"

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
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.GetDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Search user/group in directory service
  Search user or group in the directory service

  parameters:-
  -> body (iam.v4.authn.DirectoryServiceSearchQuery) (required) : Search user or group in the directory service
  -> extId (string) (required) : External identifier of the directory service
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.SearchDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) SearchDirectoryService(body *import3.DirectoryServiceSearchQuery, extId *string, args ...map[string]interface{}) (*import3.SearchDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services/{extId}/search"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
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
	contentTypes := []string{"application/json"}

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.SearchDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update directory service
  Update a directory service

  parameters:-
  -> body (iam.v4.authn.DirectoryService) (required) : Update a directory service
  -> extId (string) (required) : External identifier of the directory service
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*iam.v4.authn.UpdateDirectoryServiceApiResponse, error)
*/
func (api *DirectoryServiceApi) UpdateDirectoryService(body *import3.DirectoryService, extId *string, args ...map[string]interface{}) (*import3.UpdateDirectoryServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/iam/v4.0.a1/authn/directory-services/{extId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
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
	contentTypes := []string{"application/json"}

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.UpdateDirectoryServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
