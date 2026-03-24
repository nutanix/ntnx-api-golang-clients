package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/request/licenses"
	"net/http"
	"net/url"
	"strings"
)

type LicensesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *LicensesServiceApi
}

type LicensesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewLicensesApi(apiClient *client.ApiClient) *LicensesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LicensesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewLicensesServiceApi(a.ApiClient)

	return a
}

func NewLicensesServiceApi(apiClient *client.ApiClient) *LicensesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LicensesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// API to fetch allowances.
func (api *LicensesApi) ListAllowances(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListAllowancesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAllowances(context.Background(), &import5.ListAllowancesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// API to fetch allowances.
func (api *LicensesServiceApi) ListAllowances(ctx context.Context, request *import5.ListAllowancesRequest, args ...map[string]interface{}) (*import3.ListAllowancesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/allowances"

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

	unmarshalledResp := new(import3.ListAllowancesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch list of license compliances across clusters.
func (api *LicensesApi) ListCompliances(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListCompliancesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCompliances(context.Background(), &import5.ListCompliancesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// API to fetch list of license compliances across clusters.
func (api *LicensesServiceApi) ListCompliances(ctx context.Context, request *import5.ListCompliancesRequest, args ...map[string]interface{}) (*import3.ListCompliancesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/compliances"

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

	unmarshalledResp := new(import3.ListCompliancesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch entitlements list.
func (api *LicensesApi) ListEntitlements(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListEntitlementsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEntitlements(context.Background(), &import5.ListEntitlementsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// API to fetch entitlements list.
func (api *LicensesServiceApi) ListEntitlements(ctx context.Context, request *import5.ListEntitlementsRequest, args ...map[string]interface{}) (*import3.ListEntitlementsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/entitlements"

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

	unmarshalledResp := new(import3.ListEntitlementsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch features list.
func (api *LicensesApi) ListFeatures(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListFeaturesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListFeatures(context.Background(), &import5.ListFeaturesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// API to fetch features list.
func (api *LicensesServiceApi) ListFeatures(ctx context.Context, request *import5.ListFeaturesRequest, args ...map[string]interface{}) (*import3.ListFeaturesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/features"

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

	unmarshalledResp := new(import3.ListFeaturesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch a list of licenses along with consumption details using expansion.
func (api *LicensesApi) ListLicenses(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListLicensesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLicenses(context.Background(), &import5.ListLicensesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// API to fetch a list of licenses along with consumption details using expansion.
func (api *LicensesServiceApi) ListLicenses(ctx context.Context, request *import5.ListLicensesRequest, args ...map[string]interface{}) (*import3.ListLicensesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/licenses"

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

	unmarshalledResp := new(import3.ListLicensesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the available licensing recommendations from the license portal.
func (api *LicensesApi) ListRecommendations(page_ *int, limit_ *int, select_ *string, args ...map[string]interface{}) (*import3.ListRecommendationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRecommendations(context.Background(), &import5.ListRecommendationsRequest{
		Page_:   page_,
		Limit_:  limit_,
		Select_: select_,
	}, args...)
}

// Lists all the available licensing recommendations from the license portal.
func (api *LicensesServiceApi) ListRecommendations(ctx context.Context, request *import5.ListRecommendationsRequest, args ...map[string]interface{}) (*import3.ListRecommendationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/recommendations"

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

	unmarshalledResp := new(import3.ListRecommendationsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch license settings list.
func (api *LicensesApi) ListSettings(page_ *int, limit_ *int, args ...map[string]interface{}) (*import3.ListSettingsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSettings(context.Background(), &import5.ListSettingsRequest{
		Page_:  page_,
		Limit_: limit_,
	}, args...)
}

// API to fetch license settings list.
func (api *LicensesServiceApi) ListSettings(ctx context.Context, request *import5.ListSettingsRequest, args ...map[string]interface{}) (*import3.ListSettingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/settings"

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

	unmarshalledResp := new(import3.ListSettingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API to fetch license violations list.
func (api *LicensesApi) ListViolations(page_ *int, limit_ *int, args ...map[string]interface{}) (*import3.ListViolationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListViolations(context.Background(), &import5.ListViolationsRequest{
		Page_:  page_,
		Limit_: limit_,
	}, args...)
}

// API to fetch license violations list.
func (api *LicensesServiceApi) ListViolations(ctx context.Context, request *import5.ListViolationsRequest, args ...map[string]interface{}) (*import3.ListViolationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/violations"

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

	unmarshalledResp := new(import3.ListViolationsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// API for seamless licensing actions like post the cluster summary file, apply license file, and others.
func (api *LicensesApi) SyncLicenseState(body *import3.LicenseStateSyncSpec, args ...map[string]interface{}) (*import3.SyncLicenseStateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicensesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SyncLicenseState(context.Background(), &import5.SyncLicenseStateRequest{
		Body: body,
	}, args...)
}

// API for seamless licensing actions like post the cluster summary file, apply license file, and others.
func (api *LicensesServiceApi) SyncLicenseState(ctx context.Context, request *import5.SyncLicenseStateRequest, args ...map[string]interface{}) (*import3.SyncLicenseStateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/$actions/sync-license-state"

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

	unmarshalledResp := new(import3.SyncLicenseStateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
