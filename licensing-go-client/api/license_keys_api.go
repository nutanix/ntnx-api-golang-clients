package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/request/licensekeys"
	"net/http"
	"net/url"
	"strings"
)

type LicenseKeysApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *LicenseKeysServiceApi
}

type LicenseKeysServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewLicenseKeysApi(apiClient *client.ApiClient) *LicenseKeysApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LicenseKeysApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewLicenseKeysServiceApi(a.ApiClient)

	return a
}

func NewLicenseKeysServiceApi(apiClient *client.ApiClient) *LicenseKeysServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &LicenseKeysServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Adds the license key in the system.
func (api *LicenseKeysApi) AddLicenseKey(body *import3.LicenseKey, dryrun_ *bool, args ...map[string]interface{}) (*import3.AddLicenseKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddLicenseKey(context.Background(), &import4.AddLicenseKeyRequest{
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Adds the license key in the system.
func (api *LicenseKeysServiceApi) AddLicenseKey(ctx context.Context, request *import4.AddLicenseKeyRequest, args ...map[string]interface{}) (*import3.AddLicenseKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys"

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

	unmarshalledResp := new(import3.AddLicenseKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Assign the license keys to a cluster.
func (api *LicenseKeysApi) AssignLicenseKeys(body *import3.LicenseKeyAssignmentSpec, args ...map[string]interface{}) (*import3.AssignLicenseKeysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssignLicenseKeys(context.Background(), &import4.AssignLicenseKeysRequest{
		Body: body,
	}, args...)
}

// Assign the license keys to a cluster.
func (api *LicenseKeysServiceApi) AssignLicenseKeys(ctx context.Context, request *import4.AssignLicenseKeysRequest, args ...map[string]interface{}) (*import3.AssignLicenseKeysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/$actions/assign-license-keys"

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

	unmarshalledResp := new(import3.AssignLicenseKeysApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Associate parent and child license key.
func (api *LicenseKeysApi) AssociateLicenseKeys(extId *string, body *import3.AssociateLicenseKeySpec, args ...map[string]interface{}) (*import3.AssociateLicenseKeysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociateLicenseKeys(context.Background(), &import4.AssociateLicenseKeysRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Associate parent and child license key.
func (api *LicenseKeysServiceApi) AssociateLicenseKeys(ctx context.Context, request *import4.AssociateLicenseKeysRequest, args ...map[string]interface{}) (*import3.AssociateLicenseKeysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys/{extId}/$actions/associate-license-keys"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import3.AssociateLicenseKeysApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the license key based on the provided external identifier.
func (api *LicenseKeysApi) DeleteLicenseKeyById(extId *string, args ...map[string]interface{}) (*import3.DeleteLicenseKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteLicenseKeyById(context.Background(), &import4.DeleteLicenseKeyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the license key based on the provided external identifier.
func (api *LicenseKeysServiceApi) DeleteLicenseKeyById(ctx context.Context, request *import4.DeleteLicenseKeyByIdRequest, args ...map[string]interface{}) (*import3.DeleteLicenseKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.DeleteLicenseKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the license key details for the provided license key identifier.
func (api *LicenseKeysApi) GetLicenseKeyById(extId *string, args ...map[string]interface{}) (*import3.GetLicenseKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLicenseKeyById(context.Background(), &import4.GetLicenseKeyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the license key details for the provided license key identifier.
func (api *LicenseKeysServiceApi) GetLicenseKeyById(ctx context.Context, request *import4.GetLicenseKeyByIdRequest, args ...map[string]interface{}) (*import3.GetLicenseKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys/{extId}"

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

	unmarshalledResp := new(import3.GetLicenseKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the license keys.
func (api *LicenseKeysApi) ListLicenseKeys(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListLicenseKeysApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLicenseKeys(context.Background(), &import4.ListLicenseKeysRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Fetches the license keys.
func (api *LicenseKeysServiceApi) ListLicenseKeys(ctx context.Context, request *import4.ListLicenseKeysRequest, args ...map[string]interface{}) (*import3.ListLicenseKeysApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys"

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

	unmarshalledResp := new(import3.ListLicenseKeysApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the reclaim license tokens for all reclaimed operations.
func (api *LicenseKeysApi) ListReclaimLicenseTokens(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListReclaimLicenseTokensApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListReclaimLicenseTokens(context.Background(), &import4.ListReclaimLicenseTokensRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches the reclaim license tokens for all reclaimed operations.
func (api *LicenseKeysServiceApi) ListReclaimLicenseTokens(ctx context.Context, request *import4.ListReclaimLicenseTokensRequest, args ...map[string]interface{}) (*import3.ListReclaimLicenseTokensApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/reclaim-license-tokens"

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

	unmarshalledResp := new(import3.ListReclaimLicenseTokensApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Reclaim license key from the cluster.
func (api *LicenseKeysApi) ReclaimLicenseKey(extId *string, body *import3.ReclaimLicenseKeySpec, args ...map[string]interface{}) (*import3.ReclaimLicenseKeyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewLicenseKeysServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReclaimLicenseKey(context.Background(), &import4.ReclaimLicenseKeyRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Reclaim license key from the cluster.
func (api *LicenseKeysServiceApi) ReclaimLicenseKey(ctx context.Context, request *import4.ReclaimLicenseKeyRequest, args ...map[string]interface{}) (*import3.ReclaimLicenseKeyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/licensing/v4.3/config/license-keys/{extId}/$actions/reclaim"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
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

	unmarshalledResp := new(import3.ReclaimLicenseKeyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
