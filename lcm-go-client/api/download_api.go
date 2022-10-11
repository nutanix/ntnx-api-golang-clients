//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/common"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/operations"
	"encoding/json"
	"net/http"
    "net/url"
)

type DownloadApi struct {
  ApiClient *client.ApiClient
}

func NewDownloadApi(apiClient *client.ApiClient) *DownloadApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DownloadApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Download an image of a product through LCM
    Initiate LCM image download

    parameters:-
    -> body (lcm.v4.common.ImageDownloadSpecs) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.operations.GetDownloadApiResponse, error)
*/
func (api *DownloadApi) Download(body *import3.ImageDownloadSpecs, args ...map[string]interface{}) (*import2.GetDownloadApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/operations/$actions/downloadImages"


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
    unmarshalledResp := new(import2.GetDownloadApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

