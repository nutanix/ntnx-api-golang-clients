package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import12 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/prechecks"
	"net/http"
	"net/url"
	"strings"
)

type PrechecksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *PrechecksServiceApi
}

type PrechecksServiceApi struct {
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

	a.ServiceClient = NewPrechecksServiceApi(a.ApiClient)

	return a
}

func NewPrechecksServiceApi(apiClient *client.ApiClient) *PrechecksServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PrechecksServiceApi{
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
func (api *PrechecksApi) PerformPrechecks(body *import4.PrechecksSpec, xClusterId *string, dryrun_ *bool, args ...map[string]interface{}) (*import5.PrechecksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewPrechecksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.PerformPrechecks(context.Background(), &import12.PerformPrechecksRequest{
		Body:       body,
		XClusterId: xClusterId,
		Dryrun_:    dryrun_,
	}, args...)
}

// Perform LCM prechecks for the intended update operation.
func (api *PrechecksServiceApi) PerformPrechecks(ctx context.Context, request *import12.PerformPrechecksRequest, args ...map[string]interface{}) (*import5.PrechecksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.2/operations/$actions/prechecks"

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

	unmarshalledResp := new(import5.PrechecksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
