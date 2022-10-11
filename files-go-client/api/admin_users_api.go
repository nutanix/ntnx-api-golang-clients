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

type AdminUsersApi struct {
  ApiClient *client.ApiClient
}

func NewAdminUsersApi(apiClient *client.ApiClient) *AdminUsersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AdminUsersApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create admin user
    Create an admin user using the provided request body.  The user needs to specify the `name` and `role` for the new admin user to be created.  A sample request body would look like this: ``` {   \"role\": \"ADMIN\",   \"name\": \"administrator\" } ``` 

    parameters:-
    -> body (files.v4.config.AdminUser) (required) : The model associated with create admin user.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateAdminUserApiResponse, error)
*/
func (api *AdminUsersApi) CreateAdminUser(body *import1.AdminUser, args ...map[string]interface{}) (*import1.CreateAdminUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/admin-users"

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
    unmarshalledResp := new(import1.CreateAdminUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete admin user
    Delete the admin user with the given external identifier.  The user has to specify - a valid external identifier (`adminUserExtId`) of the admin user to be deleted.  How to pass Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/admin-users/38f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/admin-users/38f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> adminUserExtId (string) (required) : Admin user UUID. Example:28f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteAdminUserApiResponse, error)
*/
func (api *AdminUsersApi) DeleteAdminUser(adminUserExtId *string, args ...map[string]interface{}) (*import1.DeleteAdminUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/admin-users/{adminUserExtId}"

    // verify the required parameter 'adminUserExtId' is set
	if nil == adminUserExtId {
		return nil, client.ReportError("adminUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"adminUserExtId"+"}", url.PathEscape(client.ParameterToString(*adminUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteAdminUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get admin user by extId
    Get admin user with the given external identifier.  The user has to specify - a valid external identifier (`adminUserExtId`) of the admin user to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> adminUserExtId (string) (required) : Admin user UUID. Example:28f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.AdminUserApiResponse, error)
*/
func (api *AdminUsersApi) GetAdminUserByExtId(adminUserExtId *string, args ...map[string]interface{}) (*import1.AdminUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/admin-users/{adminUserExtId}"

    // verify the required parameter 'adminUserExtId' is set
	if nil == adminUserExtId {
		return nil, client.ReportError("adminUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"adminUserExtId"+"}", url.PathEscape(client.ParameterToString(*adminUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.AdminUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List admin users
    Get a list of admin users.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported:   - name   - role  A sample request URL would look like this:  ``` /files/v4.0.a2/config/file-server/admin-users?$filter=role eq Schema.Enums.AdminRole'BACKUP_OPERATOR' ```  Example of supported query parameters for admin users LIST API: ```   - ?$page=0&$limit=1   - ?$select=name, role   - ?$orderby=name desc   - ?$filter=role eq Schema.Enums.AdminRole'BACKUP_OPERATOR'   - ?$filter=startswith(name,'test')   - ?$filter=endswith(name,'test')   - ?$filter=contains(name,'admin')   - ?$select=name, role&$filter=startswith(name,'test')   - ?$limit=5&$select=name, role&$orderby=name desc ```  The `$orderby` query parameter allows specifying attributes on which to sort the returned list of admin users.  The following parameters support sorting in admin user list API:   - name   - role  A sample request URL would look like this:  ``` /files/v4.0.a2/config/file-server/admin-users?$orderby=name desc ```  The $select query parameter allows specifying attributes which the user wants to fetch in the returned list of admin users, other attributes will be returned as  a null value.  The following attributes can be selected: ```   - name   - role ```  Some more examples are given below:  1. Order by name in ascending order  ``` /files/v4.0.a2/config/file-server/admin-users?$orderby=name asc ```  2. Select by name  ``` /files/v4.0.a2/config/file-server/admin-users?$select=name ```  3. Paginate the returned list  ``` /files/v4.0.a2/config/file-server/admin-users?$page=0&$limit=1 ```  4. Combination of queries  ``` /files/v4.0.a2/config/file-server/admin-users?$limit=5&$select=name, role&$orderby=name desc ``` 

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - name - role 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - name - role 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.AdminUserListApiResponse, error)
*/
func (api *AdminUsersApi) GetAdminUsers(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.AdminUserListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/admin-users"


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
    unmarshalledResp := new(import1.AdminUserListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update admin user
    Updates an admin user with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`adminUserExtId`) of the admin user to be updated. They also need to provide a request body for performing the update.  A sample request body would look like this: ``` {   \"role\": \"ADMIN\",   \"name\": \"administrator\" } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/admin-users/38f78959-14a6-4c47-b5db-920460c4b668 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/admin-users/38f78959-14a6-4c47-b5db-920460c4b668 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.AdminUser) (required) : The model associated with create admin user.
    -> adminUserExtId (string) (required) : Admin user UUID. Example:28f78959-14a6-4c47-b5db-920460c4b668.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateAdminUserApiResponse, error)
*/
func (api *AdminUsersApi) UpdateAdminUser(body *import1.AdminUser, adminUserExtId *string, args ...map[string]interface{}) (*import1.UpdateAdminUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/admin-users/{adminUserExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'adminUserExtId' is set
	if nil == adminUserExtId {
		return nil, client.ReportError("adminUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"adminUserExtId"+"}", url.PathEscape(client.ParameterToString(*adminUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateAdminUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

