//Api classes for lifecycle's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	"net/http"
	"net/url"
	"strings"
)

type PrechecksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewPrechecksApi(apiClient *client.ApiClient) *PrechecksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PrechecksApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Perform LCM prechecks for the intended update operation.
func (api *PrechecksApi) PerformPrechecks(body *import3.PrechecksSpec, xClusterId *string, args ...map[string]interface{}) (*import2.PrechecksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.0.b1/operations/$actions/prechecks"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	if xClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import2.PrechecksApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
