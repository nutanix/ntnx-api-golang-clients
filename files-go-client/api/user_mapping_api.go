//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type UserMappingApi struct {
  ApiClient *client.ApiClient
}

func NewUserMappingApi(apiClient *client.ApiClient) *UserMappingApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserMappingApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a user mapping
    Create a user mapping using the provided request body.  User can specify fields like `explicitMapping`, `deniedAccessSmbUsers`, `ruleBasedMapping`.  A sample request body would look like this:  ``` {   \"explicitMapping\": {     \"deniedAccessSmbUsers\": [       {         \"actionType\": \"DENY_ACCESS\",         \"mappingType\": \"USER\",         \"name\": \"child4\\\\test1\"       },       {         \"actionType\": \"DENY_ACCESS\",         \"mappingType\": \"USER\",         \"name\": \"child4\\\\test2\"       }     ],     \"oneToOneMappings\": [       {         \"identityMappingType\": \"USER\",         \"nfsMapping\": {             \"name\": \"1001\"         },         \"smbMapping\": {             \"name\": \"child4\\\\administrator\"         }       }     ]   },   \"ruleBasedMapping\": {     \"nfsWithNoMappings\": [         {             \"actionType\": \"MAP_IDENTITIES\",             \"mappingType\": \"USER\",             \"name\": \"child4\\\\admin\"         },         {             \"actionType\": \"MAP_IDENTITIES\",             \"mappingType\": \"GROUP\",             \"name\": \"child4\\\\administrator\"         }     ],     \"smbWithNoMappings\": [         {           \"actionType\": \"MAP_IDENTITIES\",           \"mappingType\": \"USER\",           \"name\": \"1001\",           \"userId\": 1001         },         {           \"actionType\": \"MAP_IDENTITIES\",           \"groupId\": 1002,           \"mappingType\": \"GROUP\",           \"name\": \"1002\"         }     ],     \"templateMapping\": \"NO_TEMPLATE_MAPPING\"     }   } ``` 

    parameters:-
    -> body (files.v4.config.UserMapping) (required) : User mapping model.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateUserMappingApiResponse, error)
*/
func (api *UserMappingApi) CreateUserMapping(body *import1.UserMapping, args ...map[string]interface{}) (*import1.CreateUserMappingApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/user-mapping"

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
    unmarshalledResp := new(import1.CreateUserMappingApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a user mapping
    Get a list of user mappings.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UserMappingListApiResponse, error)
*/
func (api *UserMappingApi) GetUserMapping(args ...map[string]interface{}) (*import1.UserMappingListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/user-mapping"


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
    unmarshalledResp := new(import1.UserMappingListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Search a user mapping
    Search a user mapping using the provided request body.  A sample request body would look like this:  ``` {   \"mappingType\": \"USER\",   \"nameIds\": [     \"4321\"   ],   \"actionType\": \"MAP_IDENTITIES\",   \"searchMappingType\": \"NFS_TO_SMB\" } ``` 

    parameters:-
    -> body (files.v4.config.SearchUserMappingRequestSpec) (required) : Search a user mapping request model.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.SearchUserMappingPOSTResponses, error)
*/
func (api *UserMappingApi) SearchUserMapping(body *import1.SearchUserMappingRequestSpec, args ...map[string]interface{}) (*import1.SearchUserMappingPOSTResponses, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/search-user-mapping"

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
    unmarshalledResp := new(import1.SearchUserMappingPOSTResponses)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a user mapping
    Update a user mapping using the provided request body.  The user needs to provide a request body for performing the update. They can specify fields like `explicitMapping`, `deniedAccessSmbUsers`, `ruleBasedMapping`.  A sample request body would look like this:  ``` {   \"explicitMapping\": {     \"deniedAccessSmbUsers\": [       {         \"actionType\": \"DENY_ACCESS\",         \"mappingType\": \"USER\",         \"name\": \"child4\\\\test1\"       },       {         \"actionType\": \"DENY_ACCESS\",         \"mappingType\": \"USER\",         \"name\": \"child4\\\\test2\"       }     ],     \"oneToOneMappings\": [       {         \"identityMappingType\": \"USER\",         \"nfsMapping\": {             \"name\": \"1001\"         },         \"smbMapping\": {             \"name\": \"child4\\\\administrator\"         }       }     ]   },   \"ruleBasedMapping\": {     \"nfsWithNoMappings\": [         {             \"actionType\": \"MAP_IDENTITIES\",             \"mappingType\": \"USER\",             \"name\": \"child4\\\\admin\"         },         {             \"actionType\": \"MAP_IDENTITIES\",             \"mappingType\": \"GROUP\",             \"name\": \"child4\\\\administrator\"         }     ],     \"smbWithNoMappings\": [         {           \"actionType\": \"MAP_IDENTITIES\",           \"mappingType\": \"USER\",           \"name\": \"1001\",           \"userId\": 1001         },         {           \"actionType\": \"MAP_IDENTITIES\",           \"groupId\": 1002,           \"mappingType\": \"GROUP\",           \"name\": \"1002\"         }     ],     \"templateMapping\": \"NO_TEMPLATE_MAPPING\"     }   } ```  User mapping update operation should be performed on a complete GET response body. A GET on a resource is required before doing an UPDATE operation. Sending a partial or incomplete GET response body as update request-body might result in the existing configuration being removed.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/user-mapping ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/user-mapping ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.UserMapping) (required) : User mapping model.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateUserMappingApiResponse, error)
*/
func (api *UserMappingApi) UpdateUserMapping(body *import1.UserMapping, args ...map[string]interface{}) (*import1.UpdateUserMappingApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/user-mapping"

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
    unmarshalledResp := new(import1.UpdateUserMappingApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

