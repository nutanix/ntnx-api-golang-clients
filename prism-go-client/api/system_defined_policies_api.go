//Api classes for prism's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/policies"
	"net/http"
	"net/url"
	"strings"
)

type SystemDefinedPoliciesApi struct {
	ApiClient *client.ApiClient
}

func NewSystemDefinedPoliciesApi(apiClient *client.ApiClient) *SystemDefinedPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SystemDefinedPoliciesApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Get a list of all system defined alert policies.
  Get a list of all system defined alert policies.

  parameters:-
  -> globalConfig (bool) (optional) : Indicates whether the system defined alert policies is a default one or not.
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - alertTypeId - title
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - alertTypeId - entityUid - modifiedTimestamp - title
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.policies.ListSdaPoliciesApiResponse, error)
*/
func (api *SystemDefinedPoliciesApi) GetSdaPolicies(globalConfig *bool, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import5.ListSdaPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/policies/system-defined"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if globalConfig != nil {
		queryParams.Add("globalConfig", client.ParameterToString(*globalConfig, ""))
	}
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
	unmarshalledResp := new(import5.ListSdaPoliciesApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get a system defined alert policy based on policy ID.
  Get a system defined alert policy based on policy ID.

  parameters:-
  -> entityUid (string) (required) : Entity UID.
  -> globalConfig (bool) (optional) : Indicates whether the system defined alert policies is a default one or not.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.policies.GetSdaPolicyApiResponse, error)
*/
func (api *SystemDefinedPoliciesApi) GetSdaPolicyById(entityUid *string, globalConfig *bool, args ...map[string]interface{}) (*import5.GetSdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/policies/system-defined/{entityUid}"

	// verify the required parameter 'entityUid' is set
	if nil == entityUid {
		return nil, client.ReportError("entityUid is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"entityUid"+"}", url.PathEscape(client.ParameterToString(*entityUid, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if globalConfig != nil {
		queryParams.Add("globalConfig", client.ParameterToString(*globalConfig, ""))
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
	unmarshalledResp := new(import5.GetSdaPolicyApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update a system defined alert policy.
  Update a system defined alert policy.

  parameters:-
  -> body (prism.v4.policies.SystemDefined) (required) : System defined alert policy sent for the update.
  -> entityUid (string) (required) : Entity UID.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.policies.UpdateSdaPolicyApiResponse, error)
*/
func (api *SystemDefinedPoliciesApi) UpdateSdaPolicy(body *import5.SystemDefined, entityUid *string, args ...map[string]interface{}) (*import5.UpdateSdaPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/policies/system-defined/{entityUid}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'entityUid' is set
	if nil == entityUid {
		return nil, client.ReportError("entityUid is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"entityUid"+"}", url.PathEscape(client.ParameterToString(*entityUid, "")), -1)
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
	unmarshalledResp := new(import5.UpdateSdaPolicyApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
