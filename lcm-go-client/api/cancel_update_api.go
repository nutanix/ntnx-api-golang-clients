//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/operations"
	"encoding/json"
	"net/http"
    "net/url"
)

type CancelUpdateApi struct {
  ApiClient *client.ApiClient
}

func NewCancelUpdateApi(apiClient *client.ApiClient) *CancelUpdateApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CancelUpdateApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Set intent to cancel ongoing LCM update
    Cancel LCM Update

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.operations.CancelUpdateApiResponse, error)
*/
func (api *CancelUpdateApi) PostCancelUpdate(args ...map[string]interface{}) (*import2.CancelUpdateApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/operations/$actions/cancelupdate"


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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.CancelUpdateApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

