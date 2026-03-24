package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/report"
	import8 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/request/securitysummaries"
	"net/http"
	"net/url"
	"strings"
)

type SecuritySummariesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SecuritySummariesServiceApi
}

type SecuritySummariesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSecuritySummariesApi(apiClient *client.ApiClient) *SecuritySummariesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecuritySummariesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSecuritySummariesServiceApi(a.ApiClient)

	return a
}

func NewSecuritySummariesServiceApi(apiClient *client.ApiClient) *SecuritySummariesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SecuritySummariesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the current number of 'issues' and their trend for each cluster.
func (api *SecuritySummariesApi) ListSecuritySummaries(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import6.ListSecuritySummariesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecuritySummariesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSecuritySummaries(context.Background(), &import8.ListSecuritySummariesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Get the current number of 'issues' and their trend for each cluster.
func (api *SecuritySummariesServiceApi) ListSecuritySummaries(ctx context.Context, request *import8.ListSecuritySummariesRequest, args ...map[string]interface{}) (*import6.ListSecuritySummariesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/report/security-summaries"

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
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import6.ListSecuritySummariesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Triggers a refresh operation for updating security stats on Prism Central and Prism Element.
func (api *SecuritySummariesApi) RefreshSecuritySummaries(args ...map[string]interface{}) (*import6.RefreshSecuritySummariesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSecuritySummariesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RefreshSecuritySummaries(context.Background(), &import8.RefreshSecuritySummariesRequest{}, args...)
}

// Triggers a refresh operation for updating security stats on Prism Central and Prism Element.
func (api *SecuritySummariesServiceApi) RefreshSecuritySummaries(ctx context.Context, request *import8.RefreshSecuritySummariesRequest, args ...map[string]interface{}) (*import6.RefreshSecuritySummariesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/security/v4.1/report/security-summaries/$actions/refresh"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import6.RefreshSecuritySummariesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
