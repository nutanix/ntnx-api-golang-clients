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

type MountTargetsApi struct {
  ApiClient *client.ApiClient
}

func NewMountTargetsApi(apiClient *client.ApiClient) *MountTargetsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &MountTargetsApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a mount target
    Create a mount target using the provided request body.  The user has to specify `name` and  `type` of the mount target which needs to be created. User can also specify `maxSizeGib`, `path`, `secondaryProtocol`, `state` etc of the mount target.  A sample request body would look like this: ``` {   \"protocol\": \"NFS\",   \"name\": \"sharetest\",   \"maxSizeGib\": 1,   \"extId\": \"2f2193f0-0c9e-4be4-571a-949694761ff3\",   \"enableCompression\": true,   \"secondaryProtocol\": [     \"NONE\"   ],   \"nfsProperties\": {     \"authenticationType\": \"SYSTEM\",     \"clientAccessType\": {       \"defaultAccessType\": \"READ_WRITE\"     }   },   \"type\": \"GENERAL\",   \"description\": \"share\" } ``` 

    parameters:-
    -> body (files.v4.config.MountTarget) (required) : Mount target model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateMountTargetApiResponse, error)
*/
func (api *MountTargetsApi) CreateMountTargets(body *import1.MountTarget, args ...map[string]interface{}) (*import1.CreateMountTargetApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets"

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
    unmarshalledResp := new(import1.CreateMountTargetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete a mount target
    Delete a mount target with the given external identifier.  The user has to specify - a valid external identifier (`mountTargetExtId`) of the mount target to be deleted. User can also specify `force` flag to delete mount target forcefully.  How to use Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9 ````  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> force (bool) (optional) : Force flag to delete the mount target. If set to true, mount target will be deleted even if there are users accessing data from it right now.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteMountTargetApiResponse, error)
*/
func (api *MountTargetsApi) DeleteMountTargetById(mountTargetExtId *string, force *bool, args ...map[string]interface{}) (*import1.DeleteMountTargetApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}"

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
	if force != nil {
		queryParams.Add("force", client.ParameterToString(*force, ""))
	}

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
    unmarshalledResp := new(import1.DeleteMountTargetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List all mount targets
    Get a paginated list of mount targets.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported:   - name   - type  A sample request URL would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$filter=type eq Schema.Enums.MountTargetType'DISTRIBUTED' ```  Example of supported query parameters for mount targets LIST API: ```   - ?$page=0&$limit=1   - ?$select=name, description   - ?$orderby=name desc   - ?$filter=type eq Schema.Enums.MountTargetType'DISTRIBUTED'   - ?$filter=startswith(name,'test')   - ?$filter=endswith(name,'test')   - ?$filter=contains(name,'nested')   - ?$select=name, description, path&$filter=startswith(name,'test')   - ?$limit=5&$select=name, description, path&$orderby=name desc ```  The `$orderby` query parameter allows specifying attributes on which to sort the returned list of mount targets.  The following parameters support sorting in mount target list API:   - enableCompression   - isEnableCompression   - name   - path   - protocol   - type  A sample request URL would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$orderby=name desc ```  The $select query parameter allows specifying attributes which the user wants to fetch in the returned list of mount targets, other attributes will be returned as  a null value.  The following attributes can be selected: ```   - name   - description   - fileServerExtId   - maxSizeGib   - path   - parentMountTargetExtId   - enableCompression   - isEnableCompression   - enablePreviousVersion   - isEnablePreviousVersion   - longnameEnabled   - isLongnameEnabled ```  Some more examples are given below:  1. Order by name in ascending order  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$orderby=name asc ```  2. Select by name  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$select=name ```  3. Paginate the returned list  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$page=0&$limit=1 ```  4. Combination of queries  ``` /api/files/v4.0.a2/config/file-server/mount-targets?$limit=5&$select=name, description, path&$orderby=name desc ``` 

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - name - type 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - enableCompression - isEnableCompression - name - path - protocol - type 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - blockedClients - description - enableCompression - enablePreviousVersion - extId - fileBlockingExtensions - fileServerExtId - isEnableCompression - isEnablePreviousVersion - isLongnameEnabled - links - longnameEnabled - maxSizeGib - multiProtocolProperties - name - nfsProperties - parentMountTargetExtId - path - protocol - secondaryProtocol - smbProperties/enableAccessBasedEnumeration - smbProperties/enableCa - smbProperties/enableSmb3Encryption - smbProperties/isEnableAccessBasedEnumeration - smbProperties/isEnableCa - smbProperties/isEnableSmb3Encryption - state - statusType - tenantId - type - workloadType - wormSpec 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MountTargetListApiResponse, error)
*/
func (api *MountTargetsApi) GetAllMountTargets(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.MountTargetListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets"


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
    unmarshalledResp := new(import1.MountTargetListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get the mimetype of files for the given inode genId pair.
    Query the mimetypes of files of mount target with given external identifier, Inode and genId.  The user has to specify - a valid external identifier (`mountTargetExtId`) of mount target for which mimetypes of files need to be fetched and a list of Inode and genId pairs (`inodeGenIdList`) in the path parameter. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> inodeGenIdList ([]string) (required) : List of Inode and genId pairs in the format inode:genId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MimeTypeResponseApiResponse, error)
*/
func (api *MountTargetsApi) GetFileByInodeGenId(mountTargetExtId *string, inodeGenIdList *[]string, args ...map[string]interface{}) (*import1.MimeTypeResponseApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/files"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'inodeGenIdList' is set
	if nil == inodeGenIdList {
		return nil, client.ReportError("inodeGenIdList is required and must be specified")
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
	queryParams.Add("inodeGenIdList", client.ParameterToString(*inodeGenIdList, "csv"))

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
    unmarshalledResp := new(import1.MimeTypeResponseApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a mount target by extId
    Get the mount target with the given external identifier.  The user has to specify - a valid external identifier (`mountTargetExtId`) of the mount target to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MountTargetApiResponse, error)
*/
func (api *MountTargetsApi) GetMountTargetByUuid(mountTargetExtId *string, args ...map[string]interface{}) (*import1.MountTargetApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}"

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
    unmarshalledResp := new(import1.MountTargetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a mount target
    Update a mount target with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`mountTargetExtId`) of the mount target to be updated. They also need to provide a request body for performing the update.  A sample request body would look like this:  ``` {   \"protocol\": \"NFS\",   \"name\": \"sharetest\",   \"maxSizeGib\": 1,   \"extId\": \"9c1e537d-6777-4c22-5d41-ddd0c3337aa9\",   \"enableCompression\": true,   \"secondaryProtocol\": [     \"NONE\"   ],   \"nfsProperties\": {     \"authenticationType\": \"SYSTEM\",     \"clientAccessType\": {       \"defaultAccessType\": \"READ_WRITE\"     }   },   \"type\": \"GENERAL\",   \"description\": \"share\" } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to use Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9 ````  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.MountTarget) (required) : Mount target model
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateMountTargetApiResponse, error)
*/
func (api *MountTargetsApi) UpdateMountTargets(body *import1.MountTarget, mountTargetExtId *string, args ...map[string]interface{}) (*import1.UpdateMountTargetApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}"

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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateMountTargetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

