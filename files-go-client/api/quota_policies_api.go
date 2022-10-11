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

type QuotaPoliciesApi struct {
  ApiClient *client.ApiClient
}

func NewQuotaPoliciesApi(apiClient *client.ApiClient) *QuotaPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &QuotaPoliciesApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a quota policy
    Create a quota policy using the provided request body.  The user has to specify - a valid external identifier (extId) of the mount target for which quota policy needs to be created. They can specify `enforcementType`, `principalValue`, `principalType`, `enforcementType` as part of request body.  A sample request body would look like this: ``` {   \"enableNotifications\": true,   \"sizeInBytes\": 1073741800,   \"enforcementType\": \"SOFT\",   \"principalValue\": \"administrator\",   \"mountTargetExtId\": \"9c1e537d-6777-4c22-5d41-ddd0c3337aa9\",   \"principalType\": \"USER\",   \"notificationRecipients\": [     \"tempemail@test.com\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.QuotaPolicy) (required) : Quota policy model.
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateQuotaPolicyApiResponse, error)
*/
func (api *QuotaPoliciesApi) CreateQuotaPolicy(body *import1.QuotaPolicy, mountTargetExtId *string, args ...map[string]interface{}) (*import1.CreateQuotaPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
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
    unmarshalledResp := new(import1.CreateQuotaPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete a quota policy
    Delete a quota policy with the provided extId.  The user has to specify - a valid external identifier (`extId`) of the of the mount target to which quota policy belongs to and a valid external identifier (`extId`) of the quota policy to be deleted.  How to pass Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies/48f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies/48f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> quotaPolicyExtId (string) (required) : Quota policy UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteQuotaPolicyApiResponse, error)
*/
func (api *QuotaPoliciesApi) DeleteQuotaPolicy(mountTargetExtId *string, quotaPolicyExtId *string, args ...map[string]interface{}) (*import1.DeleteQuotaPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'quotaPolicyExtId' is set
	if nil == quotaPolicyExtId {
		return nil, client.ReportError("quotaPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"quotaPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*quotaPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteQuotaPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get the quota policy email configuration
    Get the quota policy email configuration. 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetQuotaEmailConfigApiResponse, error)
*/
func (api *QuotaPoliciesApi) GetQuotaEmailConfig(args ...map[string]interface{}) (*import1.GetQuotaEmailConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/quota-email-config"


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
    unmarshalledResp := new(import1.GetQuotaEmailConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List all quota policies
    Get a paginated list of quota policies.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported: - enforcementType - principalValue - sizeInBytes  A sample call would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$filter=enforcementType eq Schema.Enums.EnforcementType'SOFT' ```  Example of supported query parameters for nested Quota policy LIST API: ```  - ?$page=0&$limit=1  - ?$select=principalType, principalValue, sizeInBytes  - ?$orderby=enforcementType desc  - ?$filter=enforcementType eq Schema.Enums.EnforcementType'SOFT'  - ?$limit=1&$select=principalType, sizeInBytes&$orderby=enforcementType desc&$filter=enforcementType eq Schema.Enums.EnforcementType'SOFT' ``` The following attributes can be selected: ``` - mountTargetExtId - principalValue - sizeInBytes - enableNotifications - isEnableNotifications ```  Some more examples are given below: 1. Order by enforcementType in ascending order  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$orderby=enforcementType asc ```  2. Order by enforcementType in descending order  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$orderby=enforcementType desc ```  3. Select by principalType  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$select=principalType ```  4. Select by principalValue  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$select=principalValue ```  5. Paginate the returned policies list  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$page=0&$limit=1 ```  6. Combination of queries  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies?$limit=1&$select=principalType, sizeInBytes&$orderby=enforcementType desc&$filter=enforcementType eq Schema.Enums.EnforcementType'SOFT' ``` 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - enforcementType - principalValue - sizeInBytes 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - enforcementType - principalValue 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - enableNotifications - enforcementType - extId - isEnableNotifications - links - mountTargetExtId - notificationRecipients - principalType - principalValue - sizeInBytes - tenantId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.QuotaPolicyListApiResponse, error)
*/
func (api *QuotaPoliciesApi) GetQuotaPolicies(mountTargetExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.QuotaPolicyListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}

    // Path Params
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
    unmarshalledResp := new(import1.QuotaPolicyListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a quota policy by extId
    Get a quota policy with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the of the mount target to which quota policy belongs to and a valid external identifier (`extId`) of the quota policy to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> quotaPolicyExtId (string) (required) : Quota policy UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.QuotaPolicyApiResponse, error)
*/
func (api *QuotaPoliciesApi) GetQuotaPolicyByExtId(mountTargetExtId *string, quotaPolicyExtId *string, args ...map[string]interface{}) (*import1.QuotaPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'quotaPolicyExtId' is set
	if nil == quotaPolicyExtId {
		return nil, client.ReportError("quotaPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"quotaPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*quotaPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.QuotaPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update the quota policy email configuration
    Update the quota policy email configuration.  The user has to specify (`content`) and (`subject`) to update the quota policy email configuration.  A sample request body would look like this: ``` {   \"content\": \"Updated email content\",   \"subject\": \"Updated email subject\" } ``` 

    parameters:-
    -> body (files.v4.config.QuotaEmailConfig) (required) : Quota policy email configuration model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateQuotaEmailConfigApiResponse, error)
*/
func (api *QuotaPoliciesApi) UpdateQuotaEmailConfig(body *import1.QuotaEmailConfig, args ...map[string]interface{}) (*import1.UpdateQuotaEmailConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/quota-email-config"

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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateQuotaEmailConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a quota policy.
    Update quota policy.  The user has to specify - a valid external identifier (`extId`) of the of the mount target to which quota policy belongs to and a valid external identifier (`extId`) of the quota policy to be updated. They also need to provide a request body for performing the update. They can provide parameters like `enableNotifications`, `sizeInBytes`, `enforcementType`.  A sample request body would look like this: ``` {   \"enableNotifications\": false,   \"sizeInBytes\": 1073741800,   \"enforcementType\": \"SOFT\",   \"principalValue\": \"administrator\",   \"mountTargetExtId\": \"9c1e537d-6777-4c22-5d41-ddd0c3337aa9\",   \"principalType\": \"USER\" } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies/48f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/quota-policies/48f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.QuotaPolicy) (required) : Quota policy model.
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> quotaPolicyExtId (string) (required) : Quota policy UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateQuotaPolicyApiResponse, error)
*/
func (api *QuotaPoliciesApi) UpdateQuotaPolicy(body *import1.QuotaPolicy, mountTargetExtId *string, quotaPolicyExtId *string, args ...map[string]interface{}) (*import1.UpdateQuotaPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/quota-policies/{quotaPolicyExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'quotaPolicyExtId' is set
	if nil == quotaPolicyExtId {
		return nil, client.ReportError("quotaPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"quotaPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*quotaPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateQuotaPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

