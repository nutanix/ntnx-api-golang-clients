package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import16 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/upgrades"
	"net/http"
	"net/url"
	"strings"
)

type UpgradesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UpgradesServiceApi
}

type UpgradesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUpgradesApi(apiClient *client.ApiClient) *UpgradesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUpgradesServiceApi(a.ApiClient)

	return a
}

func NewUpgradesServiceApi(apiClient *client.ApiClient) *UpgradesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UpgradesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Perform upgrade operation to a specific target version for discovered LCM entity/entities.
func (api *UpgradesApi) PerformUpgrade(body *import4.UpgradeSpec, xClusterId *string, dryrun_ *bool, args ...map[string]interface{}) (*import5.UpgradeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUpgradesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformUpgrade(context.Background(), &import16.PerformUpgradeRequest{
		Body:       body,
		XClusterId: xClusterId,
		Dryrun_:    dryrun_,
	}, args...)
}

// Perform upgrade operation to a specific target version for discovered LCM entity/entities.
func (api *UpgradesServiceApi) PerformUpgrade(ctx context.Context, request *import16.PerformUpgradeRequest, args ...map[string]interface{}) (*import5.UpgradeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/operations/$actions/upgrade"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Dryrun_ != nil {
		queryParams.Add("$dryrun", client.ParameterToString(*request.Dryrun_, ""))
	}
	if request.XClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
	}
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import5.UpgradeApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
