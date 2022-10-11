//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type NameServicesApi struct {
  ApiClient *client.ApiClient
}

func NewNameServicesApi(apiClient *client.ApiClient) *NameServicesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NameServicesApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Configure file server name services
    Configure/Unconfigure name services of a file server.  This helps in configuring and un-configuring name services of a file server for different protocols.  A sample request body would look like this:  ``` {   \"adDomain\": {     \"username\": \"ADUserName\",     \"password\": \"password\",     \"addUserAsFsAdmin\": false,     \"domainName\": \"child4.xyz.minerva.com\",     \"protocolType\": \"SMB\",     \"overwriteUserAccount\": true,     \"preferredDomainController\": \"phxengminpd3.child4.xyz.minerva.com\"   },   \"ldapDomain\": {     \"protocolType\": \"NFS\",     \"binddn\": \"\",     \"base\": \"dc=ldaps,dc=afs,dc=minerva,dc=com\",     \"uri\": \"ldap://10.51.38.36\"   },   \"nfsVersion\": \"NFSV3V4\" } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/$actions/configure-name-services ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/$actions/configure-name-services ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.NameServiceSpec) (required) : Configure the name services model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.ConfigureNameServiceApiResponse, error)
*/
func (api *NameServicesApi) ConfigureNameServicesByFileServer(body *import1.NameServiceSpec, args ...map[string]interface{}) (*import1.ConfigureNameServiceApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/configure-name-services"

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
    unmarshalledResp := new(import1.ConfigureNameServiceApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Name services by file server
    Get name services configured for a file server.  The response contains information of various name services configured on a file server like `adDomain`, `ldapDomain`, `localDomain`, `nfsVersion`.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.NameServicesApiResponse, error)
*/
func (api *NameServicesApi) GetNameServicesByFileServer(args ...map[string]interface{}) (*import1.NameServicesApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/name-services"


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
    unmarshalledResp := new(import1.NameServicesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

