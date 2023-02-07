//Api classes for prism's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	"net/http"
	"net/url"
)

type EmailConfigurationApi struct {
	ApiClient *client.ApiClient
}

func NewEmailConfigurationApi(apiClient *client.ApiClient) *EmailConfigurationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EmailConfigurationApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Get alert email configuration.
  Get alert email configuration.

  parameters:-
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.config.EmailConfigurationApiResponse, error)
*/
func (api *EmailConfigurationApi) GetEmailConfiguration(args ...map[string]interface{}) (*import3.EmailConfigurationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/config/email"

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
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.EmailConfigurationApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update alert email configuration.
  Update alert email configuration.

  parameters:-
  -> body (prism.v4.config.EmailConfiguration) (required) : Email configuration sent for the update.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*prism.v4.config.EmailConfigurationApiResponse, error)
*/
func (api *EmailConfigurationApi) UpdateEmailConfiguration(body *import3.EmailConfiguration, args ...map[string]interface{}) (*import3.EmailConfigurationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a1/config/email"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.EmailConfigurationApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
