//Api classes for lcm's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lcm-go-client/v4/models/lcm/v4/resources"
	"strings"
	"encoding/json"
	"net/http"
    "net/url"
)

type TaskApi struct {
  ApiClient *client.ApiClient
}

func NewTaskApi(apiClient *client.ApiClient) *TaskApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TaskApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Get details of an LCM task identified by the UUID
    Get details of an LCM task identified by the UUID

    parameters:-
    -> uuid (string) (required) : UUID of LCM task
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*lcm.v4.resources.GetTaskByUuidApiResponse, error)
*/
func (api *TaskApi) GetTaskByUuid(uuid *string, args ...map[string]interface{}) (*import1.GetTaskByUuidApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/lcm/v4.0.a1/resources/tasks/{uuid}"

    // verify the required parameter 'uuid' is set
	if nil == uuid {
		return nil, client.ReportError("uuid is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"uuid"+"}", url.PathEscape(client.ParameterToString(*uuid, "")), -1)
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
    unmarshalledResp := new(import1.GetTaskByUuidApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

