//Api classes for networking's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type ClusterCapabilityApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClusterCapabilityApi(apiClient *client.ApiClient) *ClusterCapabilityApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterCapabilityApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// List the capabilities for one or more cluster UUIDs. Requires Prism Central >= pc.2023.3.
func (api *ClusterCapabilityApi) GetClusterCapabilities(page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.ClusterCapabilityApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.0.b1/config/capabilities"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {

		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {

		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {

		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.ClusterCapabilityApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
