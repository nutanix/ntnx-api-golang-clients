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

type VirusScanPolicyApi struct {
  ApiClient *client.ApiClient
}

func NewVirusScanPolicyApi(apiClient *client.ApiClient) *VirusScanPolicyApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VirusScanPolicyApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a mount target virus scan policy
    Create virus scan policy for a mount target identified by mountTargetExtId.  A sample request body would look like this:  ``` {   \"scanOnRead\": true,   \"blockAccessFile\": true,   \"scanTimeoutIntervalInSecs\": 60,   \"mountTargetExtId\": \"f1be4f9f-70b5-4a04-8460-472f032ac9d8\",   \"fileSizeExclusionInBytes\": 0,   \"scanOnWrite\": true,   \"extId\": \"5a008cef-17e4-4658-9978-1bd38f3738bc\" } ``` 

    parameters:-
    -> body (files.v4.config.VirusScanPolicyMountTarget) (required) : virus scan policy Model
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateMountTargetVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) CreateMountTargetVirusScanPolicy(body *import1.VirusScanPolicyMountTarget, mountTargetExtId *string, args ...map[string]interface{}) (*import1.CreateMountTargetVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy"

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
    unmarshalledResp := new(import1.CreateMountTargetVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Create a virus scan policy
    Creates a virus scan policy using the provided request body.  A sample request body would look like this:  ``` {   \"scanOnRead\": true,   \"blockAccessFile\": true,   \"scanTimeoutIntervalInSecs\": 120,   \"fileTypeExclusions\": [     \".txt\",     \".db\"   ],   \"fileSizeExclusionInBytes\": 0,   \"extId\": \"d96df616-c3e4-4ee3-8425-05d6182b27a1\",   \"enableAntiVirus\": false,   \"scanOnWrite\": false } ``` 

    parameters:-
    -> body (files.v4.config.VirusScanPolicy) (required) : virus scan policy Model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) CreateVirusScanPolicy(body *import1.VirusScanPolicy, args ...map[string]interface{}) (*import1.CreateVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/virus-scan-policy"

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
    unmarshalledResp := new(import1.CreateVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete a mount target virus scan policy
    Delete virus scan policy request for a mount target identified by mountTargetExtId. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> virusScanPolicyExtId (string) (required) : The {extId} of the virus scan policy.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteMountTargetVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) DeleteMountTargetVirusScanPolicy(mountTargetExtId *string, virusScanPolicyExtId *string, args ...map[string]interface{}) (*import1.DeleteMountTargetVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'virusScanPolicyExtId' is set
	if nil == virusScanPolicyExtId {
		return nil, client.ReportError("virusScanPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"virusScanPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*virusScanPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteMountTargetVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a mount target virus scan policy
    Get virus scan policy with given external identifier for a mount target with provided mount target  external identifier.  The user has to specify - a valid external identifier `mountTargetExtId` of the mount target and `virusScanPolicyExtId` of virus scan policy which  needs to be fetched. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> virusScanPolicyExtId (string) (required) : The {extId} of the virus scan policy.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MountTargetVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) GetMountTargetVirusScanPolicy(mountTargetExtId *string, virusScanPolicyExtId *string, args ...map[string]interface{}) (*import1.MountTargetVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'virusScanPolicyExtId' is set
	if nil == virusScanPolicyExtId {
		return nil, client.ReportError("virusScanPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"virusScanPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*virusScanPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.MountTargetVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a virus scan policy
    Get a virus scan policy identified by {extId}. 

    parameters:-
    -> virusScanPolicyExtId (string) (required) : The {extId} of the virus scan policy.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.VirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) GetVirusScanPolicy(virusScanPolicyExtId *string, args ...map[string]interface{}) (*import1.VirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/virus-scan-policy/{virusScanPolicyExtId}"

    // verify the required parameter 'virusScanPolicyExtId' is set
	if nil == virusScanPolicyExtId {
		return nil, client.ReportError("virusScanPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"virusScanPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*virusScanPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.VirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List mount target virus scan policies
    List virus scan policy for a mount target with provided mountTargetExtId.  The user has to specify - a valid external identifier (`mountTargetExtId`) of the mount target for which the virus scan policy needs to be fetched.  A sample request URL would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/ff6cafea-a1dc-4d40-776a-cc2ce55be211/virus-scan-policy ``` 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.ListMountTargetVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) ListMountTargetVirusScanPolicy(mountTargetExtId *string, args ...map[string]interface{}) (*import1.ListMountTargetVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy"

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
    unmarshalledResp := new(import1.ListMountTargetVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List virus scan policies
    List virus scan policy for a file server.  A sample request URL would look like this:  ``` /api/files/v4.0.a2/config/file-server/virus-scan-policy ``` 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.ListVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) ListVirusScanPolicy(args ...map[string]interface{}) (*import1.ListVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/virus-scan-policy"


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
    unmarshalledResp := new(import1.ListVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a mount target virus scan policy
    Update virus scan policy request for a mount target identified by mountTargetExtId.  A sample request body would look like this:  ``` {   \"scanOnRead\": true,   \"blockAccessFile\": true,   \"scanTimeoutIntervalInSecs\": 60,   \"fileTypeExclusions\": [     \".pdf\"   ],   \"mountTargetExtId\": \"4e624597-8a37-45aa-b4ac-0f992de0a903\",   \"scanOnWrite\": true,   \"extId\": \"dc9d2193-10db-45eb-8ff3-3bef0fe6fb07\" } ``` 

    parameters:-
    -> body (files.v4.config.VirusScanPolicyMountTarget) (required) : virus scan policy Model
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> virusScanPolicyExtId (string) (required) : The {extId} of the virus scan policy.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateMountTargetVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) UpdateMountTargetVirusScanPolicy(body *import1.VirusScanPolicyMountTarget, mountTargetExtId *string, virusScanPolicyExtId *string, args ...map[string]interface{}) (*import1.UpdateMountTargetVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/virus-scan-policy/{virusScanPolicyExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'virusScanPolicyExtId' is set
	if nil == virusScanPolicyExtId {
		return nil, client.ReportError("virusScanPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"virusScanPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*virusScanPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateMountTargetVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a virus scan policy
    Update virus scan policy for file server identified for the provided extId.  A sample request body would look like this:  ``` {   \"scanOnRead\": true,   \"enableAntiVirus\": true,   \"scanTimeoutIntervalInSecs\": 120,   \"fileTypeExclusions\": [     \".txt\",     \".db\"   ],   \"scanOnWrite\": false,   \"fileSizeExclusionInBytes\": 0,   \"extId\": \"d96df616-c3e4-4ee3-8425-05d6182b27a1\" } ``` 

    parameters:-
    -> body (files.v4.config.VirusScanPolicy) (required) : virus scan policy Model
    -> virusScanPolicyExtId (string) (required) : The {extId} of the virus scan policy.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateVirusScanPolicyApiResponse, error)
*/
func (api *VirusScanPolicyApi) UpdateVirusScanPolicy(body *import1.VirusScanPolicy, virusScanPolicyExtId *string, args ...map[string]interface{}) (*import1.UpdateVirusScanPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/virus-scan-policy/{virusScanPolicyExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'virusScanPolicyExtId' is set
	if nil == virusScanPolicyExtId {
		return nil, client.ReportError("virusScanPolicyExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"virusScanPolicyExtId"+"}", url.PathEscape(client.ParameterToString(*virusScanPolicyExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateVirusScanPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

