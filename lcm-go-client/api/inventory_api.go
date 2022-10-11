//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/operations"
	"encoding/json"
	"net/http"
    "net/url"
)

type InventoryApi struct {
  ApiClient *client.ApiClient
}

func NewInventoryApi(apiClient *client.ApiClient) *InventoryApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &InventoryApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Perform an inventory operation to identify entities that can be updated through LCM
    Perform an LCM inventory operation

    parameters:-
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.operations.InventoryApiResponse, error)
*/
func (api *InventoryApi) Inventory(args ...map[string]interface{}) (*import2.InventoryApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/operations/$actions/performInventory"


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
    unmarshalledResp := new(import2.InventoryApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

