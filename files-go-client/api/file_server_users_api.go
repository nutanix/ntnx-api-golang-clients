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

type FileServerUsersApi struct {
  ApiClient *client.ApiClient
}

func NewFileServerUsersApi(apiClient *client.ApiClient) *FileServerUsersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &FileServerUsersApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a file server user
    Create a file server user using the provided request body.  The user has to specify the request body for creating an user. They need to specify `name`, `password`, `roles` inside request body.  A sample request body would look like this: ``` {   \"password\": \"testuser\",   \"name\": \"testuser\",   \"roles\": [     \"ROLE_CLUSTER_ADMIN\",     \"ROLE_CLUSTER_VIEWER\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.FileServerUser) (required) : File server user model.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateFileServerUserApiResponse, error)
*/
func (api *FileServerUsersApi) CreateFileServerUser(body *import1.FileServerUser, args ...map[string]interface{}) (*import1.CreateFileServerUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/users"

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
    unmarshalledResp := new(import1.CreateFileServerUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete file server user
    Delete file server user with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the file server user to be deleted. 

    parameters:-
    -> fileServerUserExtId (string) (required) : File server user UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteFileServerUserApiResponse, error)
*/
func (api *FileServerUsersApi) DeleteFileServerUser(fileServerUserExtId *string, args ...map[string]interface{}) (*import1.DeleteFileServerUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/users/{fileServerUserExtId}"

    // verify the required parameter 'fileServerUserExtId' is set
	if nil == fileServerUserExtId {
		return nil, client.ReportError("fileServerUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"fileServerUserExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteFileServerUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get file server user by extId
    Get a file server user with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the file server user to be fetched. 

    parameters:-
    -> fileServerUserExtId (string) (required) : File server user UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.FileServerUserApiResponse, error)
*/
func (api *FileServerUsersApi) GetFileServerUserByExtId(fileServerUserExtId *string, args ...map[string]interface{}) (*import1.FileServerUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/users/{fileServerUserExtId}"

    // verify the required parameter 'fileServerUserExtId' is set
	if nil == fileServerUserExtId {
		return nil, client.ReportError("fileServerUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"fileServerUserExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.FileServerUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List all the file server users
    Get a list of file server users.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter).  A sample request URL with limit would look like this:  ``` /api/files/v4.0.a2/config/file-server/users?$limit=1 ``` 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.FileServerUserListApiResponse, error)
*/
func (api *FileServerUsersApi) GetFileServerUsers(args ...map[string]interface{}) (*import1.FileServerUserListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/users"


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
    unmarshalledResp := new(import1.FileServerUserListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update file server user
    Update a file server user with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`extId`) of the file server user to be updated. They also need to provide a request body for performing the update. They need to specify the `name` and `password` of the file server user. They can also specify `roles` in the request body.  A sample request body would look like this:  ``` {   \"password\": \"password\",   \"name\": \"testuser\",   \"roles\": [     \"ROLE_CLUSTER_ADMIN\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.FileServerUser) (required) : File server user model.
    -> fileServerUserExtId (string) (required) : File server user UUID. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateFileServerUserApiResponse, error)
*/
func (api *FileServerUsersApi) UpdateFileServerUser(body *import1.FileServerUser, fileServerUserExtId *string, args ...map[string]interface{}) (*import1.UpdateFileServerUserApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/users/{fileServerUserExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'fileServerUserExtId' is set
	if nil == fileServerUserExtId {
		return nil, client.ReportError("fileServerUserExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"fileServerUserExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerUserExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateFileServerUserApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

