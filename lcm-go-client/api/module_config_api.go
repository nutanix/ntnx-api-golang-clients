//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/resources"
	"encoding/json"
	"net/http"
    "net/url"
)

type ModuleConfigApi struct {
  ApiClient *client.ApiClient
}

func NewModuleConfigApi(apiClient *client.ApiClient) *ModuleConfigApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ModuleConfigApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get LCM module configuration
    Get LCM module configuration

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GetModuleConfigApiResponse, error)
*/
func (api *ModuleConfigApi) GetModuleConfig(args ...map[string]interface{}) (*import1.GetModuleConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/moduleConfig"


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
    unmarshalledResp := new(import1.GetModuleConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update the LCM module configuration
    Update LCM module configuration

    parameters:-
    -> body (lcm.v4.resources.ModuleConfig) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.PutModuleConfigApiResponse, error)
*/
func (api *ModuleConfigApi) SetModuleConfig(body *import1.ModuleConfig, args ...map[string]interface{}) (*import1.PutModuleConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/moduleConfig"


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
    unmarshalledResp := new(import1.PutModuleConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

