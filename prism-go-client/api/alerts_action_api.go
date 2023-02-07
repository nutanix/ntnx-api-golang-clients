//Api classes for prism's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/alerts"
	"net/http"
	"net/url"
)

type AlertsActionApi struct {
	ApiClient *client.ApiClient
}

func NewAlertsActionApi(apiClient *client.ApiClient) *AlertsActionApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AlertsActionApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Acknowledge or resolve the specified alerts.
  Acknowledge or resolve the specified alerts.

  parameters:-
  -> body (prism.v4.alerts.AlertAction) (required) : A list of alert UUIDs on which the action should be taken. Action can be RESOLVE or ACKNOWLEDGE.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.alerts.AlertsActionTaskApiResponse, error)
*/
func (api *AlertsActionApi) PerformAlertAction(body *import1.AlertAction, args ...map[string]interface{}) (*import1.AlertsActionTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/alerts/$action/performAction"

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
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AlertsActionTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
