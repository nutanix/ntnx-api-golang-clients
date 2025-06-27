package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/aws/config"
	"net/http"
	"net/url"
	"strings"
)

type AwsVpcsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAwsVpcsApi(apiClient *client.ApiClient) *AwsVpcsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AwsVpcsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the list of NC2 AWS VPCs associated with a Cluster.
func (api *AwsVpcsApi) ListAwsVpcs(xClusterId *string, args ...map[string]interface{}) (*import1.ListAwsVpcsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.1/aws/config/vpcs"

	// verify the required parameter 'xClusterId' is set
	if nil == xClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	headerParams["X-Cluster-Id"] = client.ParameterToString(*xClusterId, "")
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(*string); headerValueOk {
					headerParams[headerKey] = *headerValue
				}
			}
		}
	}

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.ListAwsVpcsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
