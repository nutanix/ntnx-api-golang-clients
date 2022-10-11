//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/resources"
	"encoding/json"
	"net/http"
    "net/url"
)

type StatusApi struct {
  ApiClient *client.ApiClient
}

func NewStatusApi(apiClient *client.ApiClient) *StatusApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &StatusApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get the LCM framework status
    Get the LCM framework status

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GetLcmStatusApiResponse, error)
*/
func (api *StatusApi) GetStatus(args ...map[string]interface{}) (*import1.GetLcmStatusApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/status"


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
    unmarshalledResp := new(import1.GetLcmStatusApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

