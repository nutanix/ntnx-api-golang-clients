//Api classes for prism's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	"strings"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
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
    TaskCancel
    API to cancel an ongoing task.

    parameters:-
    -> taskExtId (string) (required) : Globally unique identifier of a task.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (prism.v4.config.TaskCancelResponse, error)
*/
func (api *TaskApi) TaskCancel(taskExtId *string, args ...map[string]interface{}) (*import2.TaskCancelResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/prism/v4.0.a1/config/tasks/{taskExtId}/$actions/cancel"

    // verify the required parameter 'taskExtId' is set
	if nil == taskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*taskExtId, "")), -1)
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
    unmarshalledResp := new(import2.TaskCancelResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    TaskGet
    API to fetch an asyn operation called task.

    parameters:-
    -> taskExtId (string) (required) : Globally unique identifier of a task.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (prism.v4.config.TaskGetResponse, error)
*/
func (api *TaskApi) TaskGet(taskExtId *string, args ...map[string]interface{}) (*import2.TaskGetResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/prism/v4.0.a1/config/tasks/{taskExtId}"

    // verify the required parameter 'taskExtId' is set
	if nil == taskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*taskExtId, "")), -1)
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
    unmarshalledResp := new(import2.TaskGetResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

