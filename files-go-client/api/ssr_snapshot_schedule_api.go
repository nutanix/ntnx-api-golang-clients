//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	"strings"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type SsrSnapshotScheduleApi struct {
  ApiClient *client.ApiClient
}

func NewSsrSnapshotScheduleApi(apiClient *client.ApiClient) *SsrSnapshotScheduleApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SsrSnapshotScheduleApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a SSR snapshot schedule
    Creates a new SSR snapshot schedule using the provided request body.  The user needs to specify the `localMaxSnapshots` and `snapshotScheduleType` of the schedule to be created. User can also provide `daysOfWeek`, `daysOfMonth`, `snapshotScheduleFrequency`.  A sample request body would look like this:  ``` {   \"snapshotScheduleType\": \"HOURLY\",   \"localMaxSnapshots\": 24,   \"snapshotScheduleFrequency\": 4 } ``` 

    parameters:-
    -> body (files.v4.config.SsrSnapshotSchedule) (required) : SSR snapshot schedule model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateSsrSnapshotScheduleApiResponse, error)
*/
func (api *SsrSnapshotScheduleApi) CreateSsrSnapshotSchedule(body *import1.SsrSnapshotSchedule, args ...map[string]interface{}) (*import1.CreateSsrSnapshotScheduleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules"

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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CreateSsrSnapshotScheduleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete SSR snapshot schedule
    Delete a SSR snapshot schedule with the given external identifier.  The user has to specify - a valid external identifier  (`extId`) of the schedule to be deleted.  How to pass Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/48f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/48f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> ssrSnapshotScheduleExtId (string) (required) : SSR snapshot schedule extId. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteSsrSnapshotScheduleApiResponse, error)
*/
func (api *SsrSnapshotScheduleApi) DeleteSsrSnapshotSchedule(ssrSnapshotScheduleExtId *string, args ...map[string]interface{}) (*import1.DeleteSsrSnapshotScheduleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId}"

    // verify the required parameter 'ssrSnapshotScheduleExtId' is set
	if nil == ssrSnapshotScheduleExtId {
		return nil, client.ReportError("ssrSnapshotScheduleExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"ssrSnapshotScheduleExtId"+"}", url.PathEscape(client.ParameterToString(*ssrSnapshotScheduleExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteSsrSnapshotScheduleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get SSR snapshot schedule by extId
    Get a SSR snapshot schedule with the given external identifier.  The user has to specify - a valid external identifier  (`extId`) of the schedule to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> ssrSnapshotScheduleExtId (string) (required) : SSR snapshot schedule extId. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.SsrSnapshotScheduleApiResponse, error)
*/
func (api *SsrSnapshotScheduleApi) GetSsrSnapshotScheduleByExtId(ssrSnapshotScheduleExtId *string, args ...map[string]interface{}) (*import1.SsrSnapshotScheduleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId}"

    // verify the required parameter 'ssrSnapshotScheduleExtId' is set
	if nil == ssrSnapshotScheduleExtId {
		return nil, client.ReportError("ssrSnapshotScheduleExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"ssrSnapshotScheduleExtId"+"}", url.PathEscape(client.ParameterToString(*ssrSnapshotScheduleExtId, "")), -1)
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
    unmarshalledResp := new(import1.SsrSnapshotScheduleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List SSR snapshot schedules
    Get a paginated list of SSR snapshot schedules.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported: - localMaxSnapshots - snapshotScheduleType  A sample request URL would look like this:  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$filter=snapshotScheduleType eq Schema.Enums.SnapshotScheduleType'WEEKLY' ```  Example of supported query parameters for SSR snapshot schedules LIST API: ```   - ?$page=0&$limit=1   - ?$select=snapshotScheduleType, localMaxSnapshots   - ?$orderby=localMaxSnapshots desc   - ?$filter=snapshotScheduleType eq Schema.Enums.SnapshotScheduleType'WEEKLY'   - ?$filter=localMaxSnapshots eq 4   - ?$filter=localMaxSnapshots ne 4   - ?$select=snapshotScheduleType, localMaxSnapshots&$filter=localMaxSnapshots ne 4   - ?$limit=5&$select=snapshotScheduleType, localMaxSnapshots&$orderby=snapshotScheduleType desc ```  The `$orderby` query parameter allows specifying attributes on which to sort the returned list of SSR snapshot schedules.  The following parameters support sorting in SSR snapshot schedule list API: - localMaxSnapshots - snapshotScheduleType  A sample request URL would look like this:  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$orderby=name desc ```  The $select query parameter allows specifying attributes which the user wants to fetch in the returned list of SSR snapshot schedules, other attributes will be returned as  a null value.  The following attributes can be selected: ```   - localMaxSnapshots   - snapshotScheduleType ```  Some more examples are given below:  1. Order by name in ascending order  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$orderby=localMaxSnapshots asc ```  2. Select by snapshotScheduleType  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$select=snapshotScheduleType ```  3. Paginate the returned list  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$page=0&$limit=1 ```  4. Combination of queries  ``` /files/v4.0.a2/config/file-server/ssr-snapshot-schedules?$limit=5&$select=snapshotScheduleType, localMaxSnapshots&$orderby=localMaxSnapshots desc ``` 

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - localMaxSnapshots - snapshotScheduleType 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - localMaxSnapshots - snapshotScheduleType 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - daysOfMonth - daysOfWeek - extId - links - localMaxSnapshots - snapshotScheduleFrequency - snapshotScheduleType - tenantId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.SsrSnapshotScheduleListApiResponse, error)
*/
func (api *SsrSnapshotScheduleApi) GetSsrSnapshotSchedules(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.SsrSnapshotScheduleListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules"


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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.SsrSnapshotScheduleListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update SSR snapshot schedule
    Update a SSR snapshot schedule with the given external identifier using the provided request body.  The user has to specify - a valid external identifier  (`extId`) of the schedule to be updated. They also need to provide a request body for   performing the update. They can provide an updated value of   `snapshotScheduleFrequency`, `snapshotScheduleType` and `localMaxSnapshots`.  A sample request body would look like this:  ``` {   \"snapshotScheduleType\": \"HOURLY\",   \"localMaxSnapshots\": 24,   \"snapshotScheduleFrequency\": 6 } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/48f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/48f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.SsrSnapshotSchedule) (required) : SSR snapshot schedule model
    -> ssrSnapshotScheduleExtId (string) (required) : SSR snapshot schedule extId. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateSsrSnapshotScheduleApiResponse, error)
*/
func (api *SsrSnapshotScheduleApi) UpdateSsrSnapshotSchedule(body *import1.SsrSnapshotSchedule, ssrSnapshotScheduleExtId *string, args ...map[string]interface{}) (*import1.UpdateSsrSnapshotScheduleApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/ssr-snapshot-schedules/{ssrSnapshotScheduleExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'ssrSnapshotScheduleExtId' is set
	if nil == ssrSnapshotScheduleExtId {
		return nil, client.ReportError("ssrSnapshotScheduleExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"ssrSnapshotScheduleExtId"+"}", url.PathEscape(client.ParameterToString(*ssrSnapshotScheduleExtId, "")), -1)
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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateSsrSnapshotScheduleApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

