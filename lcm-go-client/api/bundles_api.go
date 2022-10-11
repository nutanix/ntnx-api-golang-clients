//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/resources"
	"strings"
	"encoding/json"
	"net/http"
    "net/url"
)

type BundlesApi struct {
  ApiClient *client.ApiClient
}

func NewBundlesApi(apiClient *client.ApiClient) *BundlesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BundlesApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    DeleteBundle
    Delete bundle for the specified UUID.

    parameters:-
    -> uuid (string) (required) : UUID of the LCM bundle
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.DelBundleApiResponse, error)
*/
func (api *BundlesApi) DeleteBundle(uuid *string, args ...map[string]interface{}) (*import1.DelBundleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/bundles/{uuid}"

    // verify the required parameter 'uuid' is set
	if nil == uuid {
		return nil, client.ReportError("uuid is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"uuid"+"}", url.PathEscape(client.ParameterToString(*uuid, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DelBundleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetBundle
    Get bundle details for bundle UUID.

    parameters:-
    -> uuid (string) (required) : UUID of the LCM bundle
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GetBundleInfoApiResponse, error)
*/
func (api *BundlesApi) GetBundle(uuid *string, args ...map[string]interface{}) (*import1.GetBundleInfoApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/bundles/{uuid}"

    // verify the required parameter 'uuid' is set
	if nil == uuid {
		return nil, client.ReportError("uuid is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"uuid"+"}", url.PathEscape(client.ParameterToString(*uuid, "")), -1)
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
    unmarshalledResp := new(import1.GetBundleInfoApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetBundles
    Query list of bundles

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - type 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GetBundlesApiResponse, error)
*/
func (api *BundlesApi) GetBundles(page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.GetBundlesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/bundles"


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
    unmarshalledResp := new(import1.GetBundlesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    PostBundle
    Create a bundle

    parameters:-
    -> body (lcm.v4.resources.LcmBundlesSpec) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.PostBundlesApiResponse, error)
*/
func (api *BundlesApi) PostBundle(body *import1.LcmBundlesSpec, args ...map[string]interface{}) (*import1.PostBundlesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/bundles"


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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.PostBundlesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    PutBundle
    Update bundle information for a specified UUID

    parameters:-
    -> uuid (string) (required) : UUID of the LCM bundle
    -> body (string) (optional) : The file stream being uploaded into a bundle
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.PutBundleApiResponse, error)
*/
func (api *BundlesApi) PutBundle(uuid *string, body *string, args ...map[string]interface{}) (*import1.PutBundleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/bundles/{uuid}"

    // verify the required parameter 'uuid' is set
	if nil == uuid {
		return nil, client.ReportError("uuid is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"uuid"+"}", url.PathEscape(client.ParameterToString(*uuid, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/octet-stream"}

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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.PutBundleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

