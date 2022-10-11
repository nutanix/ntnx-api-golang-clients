//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type FileServerApi struct {
  ApiClient *client.ApiClient
}

func NewFileServerApi(apiClient *client.ApiClient) *FileServerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &FileServerApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get file server
    Get file server details.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.FileServerGetApiResponse, error)
*/
func (api *FileServerApi) GetFileServers(args ...map[string]interface{}) (*import1.FileServerGetApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server"


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
    unmarshalledResp := new(import1.FileServerGetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a file server
    Update details of a file server with the given external identifier.  The user has to specify a request body for performing the update. User can specify updated values for fields like `nvmsCount`, `sizeInGib`, `ntpServers`, `vms` etc. Network updates are not allowed as part of update file-server.  A sample request body would look like this:  ``` {   \"containerExtId\": \"18f78959-14a6-4c47-b5db-920460c4b668\",   \"nvmsCount\": 1,   \"sizeInGib\": 1024,   \"ntpServers\": [     {       \"ipv4\": {         \"value\": \"10.1.61.96\"       }     },     {       \"ipv4\": {         \"value\": \"10.40.64.15\"       }     },     {       \"ipv4\": {         \"value\": \"10.40.64.16\"       }     }   ],   \"memoryGib\": 12,   \"vms\": [     {       \"externalNetworks\": [         {           \"ipAddresses\": [             {               \"ipv4\": {                 \"value\": \"10.53.89.31\"               }             }           ]         }       ],       \"vcpus\": 4,       \"memoryGib\": 12,       \"name\": \"NTNX-mar10-fs-1\",       \"extId\": \"1a6d0f9b-be7a-409d-9c05-79974efb5c4a\"     }   ],   \"externalPrimaryNetworkExtId\": \"8064cae7-13c6-4cc9-be63-021dec2448ac\",   \"fileBlockingExtensions\": [     \".doc\"   ],   \"cvmIpAddresses\": [{     \"value\": \"10.51.48.119\"   }],   \"name\": \"mar10-fs\",   \"rebalanceNeeded\": false,   \"vcpus\": 4,   \"version\": \"4.0.0\",   \"dnsDomainName\": \"child4.afs.minerva.com\",   \"dnsServers\": [     {       \"value\": \"10.40.64.16\"     },     {       \"value\": \"10.40.64.16\"     },     {       \"value\": \"10.40.64.16\"     }   ],   \"compressionEnabled\": true } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.FileServer) (required) : A model that represents file server resources.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateFileServerApiResponse, error)
*/
func (api *FileServerApi) UpdateFileServer(body *import1.FileServer, args ...map[string]interface{}) (*import1.UpdateFileServerApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server"

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
    unmarshalledResp := new(import1.UpdateFileServerApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

