//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	"net/http"
	"net/url"
)

type SubnetMigrationApi struct {
	ApiClient *client.ApiClient
}

func NewSubnetMigrationApi(apiClient *client.ApiClient) *SubnetMigrationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SubnetMigrationApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Migrate vNICs from Acropolis to Atlas or vice-versa.
  Migrate vNICs from Acropolis to Atlas or vice-versa.

  parameters:-
  -> body (networking.v4.config.SubnetMigration) (required) : Request body for the vNIC migrate operation.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetMigrationApi) MigrateVnics(body *import1.SubnetMigration, args ...map[string]interface{}) (*import1.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.a1/config/subnets/$actions/migrate-vnics"

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
	unmarshalledResp := new(import1.TaskReferenceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
