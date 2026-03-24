package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/report"
	import9 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/vulnerabilities"
	"net/http"
	"net/url"
	"strings"
)

type VulnerabilitiesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *VulnerabilitiesServiceApi
}

type VulnerabilitiesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVulnerabilitiesApi(apiClient *client.ApiClient) *VulnerabilitiesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VulnerabilitiesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewVulnerabilitiesServiceApi(a.ApiClient)

	return a
}

func NewVulnerabilitiesServiceApi(apiClient *client.ApiClient) *VulnerabilitiesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VulnerabilitiesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Gets the vulnerabilities from the Nutanix Vulnerabilities Database (NXVD).
func (api *VulnerabilitiesApi) ListVulnerabilities(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import6.ListVulnerabilitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewVulnerabilitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVulnerabilities(context.Background(), &import9.ListVulnerabilitiesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Gets the vulnerabilities from the Nutanix Vulnerabilities Database (NXVD).
func (api *VulnerabilitiesServiceApi) ListVulnerabilities(ctx context.Context, request *import9.ListVulnerabilitiesRequest, args ...map[string]interface{}) (*import6.ListVulnerabilitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/report/vulnerabilities"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if request.Page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*request.Page_, ""))
	}
	if request.Limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*request.Limit_, ""))
	}
	if request.Filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*request.Filter_, ""))
	}
	if request.Orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*request.Orderby_, ""))
	}
	if request.Select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*request.Select_, ""))
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import6.ListVulnerabilitiesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
