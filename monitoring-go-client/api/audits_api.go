package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/audits"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type AuditsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *AuditsServiceApi
}

type AuditsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAuditsApi(apiClient *client.ApiClient) *AuditsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AuditsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewAuditsServiceApi(a.ApiClient)

	return a
}

func NewAuditsServiceApi(apiClient *client.ApiClient) *AuditsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AuditsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches the details of an audit identified by external identifier.
func (api *AuditsApi) GetAuditById(extId *string, args ...map[string]interface{}) (*import1.GetAuditApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuditsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetAuditById(context.Background(), &import4.GetAuditByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the details of an audit identified by external identifier.
func (api *AuditsServiceApi) GetAuditById(ctx context.Context, request *import4.GetAuditByIdRequest, args ...map[string]interface{}) (*import1.GetAuditApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/audits/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GetAuditApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of audits.
func (api *AuditsApi) ListAudits(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListAuditsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAuditsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAudits(context.Background(), &import4.ListAuditsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of audits.
func (api *AuditsServiceApi) ListAudits(ctx context.Context, request *import4.ListAuditsRequest, args ...map[string]interface{}) (*import1.ListAuditsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/audits"

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

	unmarshalledResp := new(import1.ListAuditsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
