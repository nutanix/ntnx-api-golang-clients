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

type InfectedFilesApi struct {
  ApiClient *client.ApiClient
}

func NewInfectedFilesApi(apiClient *client.ApiClient) *InfectedFilesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &InfectedFilesApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get infected file by extId
    Get infected files with the given external identifier.  The user has to specify - a valid external identifier (`infectedFileExtId`) of the infected file to be fetched. The user can also specify inside a path parameter - a valid external identifier(`mountTargetExtId`) to which the file belongs. 

    parameters:-
    -> infectedFileExtId (string) (required) : The extId of the infected file.
    -> mountTargetExtId (string) (optional) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.InfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) GetInfectedFileByExtId(infectedFileExtId *string, mountTargetExtId *string, args ...map[string]interface{}) (*import1.InfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/infected-files/{infectedFileExtId}"

    // verify the required parameter 'infectedFileExtId' is set
	if nil == infectedFileExtId {
		return nil, client.ReportError("infectedFileExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"infectedFileExtId"+"}", url.PathEscape(client.ParameterToString(*infectedFileExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if mountTargetExtId != nil {
		queryParams.Add("mountTargetExtId", client.ParameterToString(*mountTargetExtId, ""))
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
    unmarshalledResp := new(import1.InfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List of infected files
    Get a paginated list of infected files.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter).   Example of supported query parameters for infected files LIST API: ```  - ?$page=0&$limit=1  - ?$select=shareName, filePath, threatDescription  - ?$limit=5&$select=shareName, filePath ``` The following attributes can be selected: ```   - mountTargetExtId   - shareName   - filePath   - threatDescription   - scanTimestampUsecs   - quarantined   - icapServer ``` \" 

    parameters:-
    -> mountTargetExtId (string) (optional) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - extId - filePath - fileServerExtId - icapServer - links - mountTargetExtId - quarantined - scanTimestampUsecs - shareName - tenantId - threatDescription 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.InfectedFilesListApiResponse, error)
*/
func (api *InfectedFilesApi) ListInfectedFiles(mountTargetExtId *string, page_ *int, limit_ *int, select_ *string, args ...map[string]interface{}) (*import1.InfectedFilesListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/infected-files"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if mountTargetExtId != nil {
		queryParams.Add("mountTargetExtId", client.ParameterToString(*mountTargetExtId, ""))
	}
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
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
    unmarshalledResp := new(import1.InfectedFilesListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Quarantine infected files
    Quarantine infected files with the given infected files UUIDs in the request body.  A sample request body would look like this: ``` {   \"infectedFileUuids\": [     \"392e08f3-5d27-4655-8936-2a558fd29ef2\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.InfectedFileSpec) (required) : Quarantine infected files request.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.QuarantineInfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) QuarantineInfectedFiles(body *import1.InfectedFileSpec, args ...map[string]interface{}) (*import1.QuarantineInfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/quarantine-infected-files"

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
    unmarshalledResp := new(import1.QuarantineInfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Remove infected files
    Remove infected files with the given infected files UUID in request body.  A sample request body would look like this: ``` {   \"infectedFileUuids\": [     \"392e08f3-5d27-4655-8936-2a558fd29ef2\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.InfectedFileSpec) (required) : Remove infected files request.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.RemoveInfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) RemoveInfectedFiles(body *import1.InfectedFileSpec, args ...map[string]interface{}) (*import1.RemoveInfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/remove-infected-files"

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
    unmarshalledResp := new(import1.RemoveInfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Rescan infected files
    Rescan infected files with the given infected files UUID in request body. A sample request body would look like this: ``` {   \"infectedFileUuids\": [     \"392e08f3-5d27-4655-8936-2a558fd29ef2\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.InfectedFileSpec) (required) : Rescan infected files request
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.RescanInfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) RescanInfectedFiles(body *import1.InfectedFileSpec, args ...map[string]interface{}) (*import1.RescanInfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/rescan-infected-files"

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
    unmarshalledResp := new(import1.RescanInfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Reset infected files
    Reset infected files with the given infected files UUID in request body.  A sample request body would look like this: ``` {   \"infectedFileUuids\": [     \"392e08f3-5d27-4655-8936-2a558fd29ef2\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.InfectedFileSpec) (required) : Reset infected files request.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.ResetInfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) ResetInfectedFiles(body *import1.InfectedFileSpec, args ...map[string]interface{}) (*import1.ResetInfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/reset-infected-files"

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
    unmarshalledResp := new(import1.ResetInfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Unquarantine infected files request.
    Unquarantine infected files using the provided request body.  The user has to specify list of `infectedFileUuids` of files which needs to be unquarantined inside the request body.  A sample request body would look like this: ``` {   \"infectedFileUuids\": [     \"392e08f3-5d27-4655-8936-2a558fd29ef2\"   ] } ``` 

    parameters:-
    -> body (files.v4.config.InfectedFileSpec) (required) : Unquarantine infected files request.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UnQuarantineInfectedFilesApiResponse, error)
*/
func (api *InfectedFilesApi) UnQuarantineInfectedFiles(body *import1.InfectedFileSpec, args ...map[string]interface{}) (*import1.UnQuarantineInfectedFilesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/unquarantine-infected-files"

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
    unmarshalledResp := new(import1.UnQuarantineInfectedFilesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

