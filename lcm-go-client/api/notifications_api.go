//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/resources"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/common"
	"encoding/json"
	"net/http"
    "net/url"
)

type NotificationsApi struct {
  ApiClient *client.ApiClient
}

func NewNotificationsApi(apiClient *client.ApiClient) *NotificationsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NotificationsApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Generate LCM upgrade notifications
    Generate LCM upgrade notifications

    parameters:-
    -> body (lcm.v4.common.EntityUpdateSpecs) (optional)
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GenUpgradeNotifApiResponse, error)
*/
func (api *NotificationsApi) GenUpgradeNotifications(body *import3.EntityUpdateSpecs, args ...map[string]interface{}) (*import1.GenUpgradeNotifApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/notifications"


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
    unmarshalledResp := new(import1.GenUpgradeNotifApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

