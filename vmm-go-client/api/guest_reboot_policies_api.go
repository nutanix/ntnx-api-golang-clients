package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
	import7 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/guestrebootpolicies"
	"net/http"
	"net/url"
	"strings"
)

type GuestRebootPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *GuestRebootPoliciesServiceApi
}

type GuestRebootPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewGuestRebootPoliciesApi(apiClient *client.ApiClient) *GuestRebootPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GuestRebootPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewGuestRebootPoliciesServiceApi(a.ApiClient)

	return a
}

func NewGuestRebootPoliciesServiceApi(apiClient *client.ApiClient) *GuestRebootPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &GuestRebootPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new Guest Reboot Policy with the provided configuration.
func (api *GuestRebootPoliciesApi) CreateGuestRebootPolicy(body *import6.GuestRebootPolicy, args ...map[string]interface{}) (*import6.CreateGuestRebootPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGuestRebootPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateGuestRebootPolicy(context.Background(), &import7.CreateGuestRebootPolicyRequest{
		Body: body,
	}, args...)
}

// Creates a new Guest Reboot Policy with the provided configuration.
func (api *GuestRebootPoliciesServiceApi) CreateGuestRebootPolicy(ctx context.Context, request *import7.CreateGuestRebootPolicyRequest, args ...map[string]interface{}) (*import6.CreateGuestRebootPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/guest-reboot-policies"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.CreateGuestRebootPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the specified Guest Reboot Policy.
func (api *GuestRebootPoliciesApi) DeleteGuestRebootPolicyById(extId *string, args ...map[string]interface{}) (*import6.DeleteGuestRebootPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGuestRebootPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteGuestRebootPolicyById(context.Background(), &import7.DeleteGuestRebootPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the specified Guest Reboot Policy.
func (api *GuestRebootPoliciesServiceApi) DeleteGuestRebootPolicyById(ctx context.Context, request *import7.DeleteGuestRebootPolicyByIdRequest, args ...map[string]interface{}) (*import6.DeleteGuestRebootPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/guest-reboot-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.DeleteGuestRebootPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves the Guest Reboot Policy configuration of the provided Guest Reboot Policy external identifier.
func (api *GuestRebootPoliciesApi) GetGuestRebootPolicyById(extId *string, args ...map[string]interface{}) (*import6.GetGuestRebootPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGuestRebootPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetGuestRebootPolicyById(context.Background(), &import7.GetGuestRebootPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the Guest Reboot Policy configuration of the provided Guest Reboot Policy external identifier.
func (api *GuestRebootPoliciesServiceApi) GetGuestRebootPolicyById(ctx context.Context, request *import7.GetGuestRebootPolicyByIdRequest, args ...map[string]interface{}) (*import6.GetGuestRebootPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/guest-reboot-policies/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.GetGuestRebootPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List Guest Reboot Policies.
func (api *GuestRebootPoliciesApi) ListGuestRebootPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import6.ListGuestRebootPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGuestRebootPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListGuestRebootPolicies(context.Background(), &import7.ListGuestRebootPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List Guest Reboot Policies.
func (api *GuestRebootPoliciesServiceApi) ListGuestRebootPolicies(ctx context.Context, request *import7.ListGuestRebootPoliciesRequest, args ...map[string]interface{}) (*import6.ListGuestRebootPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/guest-reboot-policies"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.ListGuestRebootPoliciesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the specified Guest Reboot Policy.
func (api *GuestRebootPoliciesApi) UpdateGuestRebootPolicyById(extId *string, body *import6.GuestRebootPolicy, args ...map[string]interface{}) (*import6.UpdateGuestRebootPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewGuestRebootPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateGuestRebootPolicyById(context.Background(), &import7.UpdateGuestRebootPolicyByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the specified Guest Reboot Policy.
func (api *GuestRebootPoliciesServiceApi) UpdateGuestRebootPolicyById(ctx context.Context, request *import7.UpdateGuestRebootPolicyByIdRequest, args ...map[string]interface{}) (*import6.UpdateGuestRebootPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/ahv/policies/guest-reboot-policies/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import6.UpdateGuestRebootPolicyApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
