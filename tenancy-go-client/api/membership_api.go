package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/request/membership"
	"net/http"
	"net/url"
	"strings"
)

type MembershipApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *MembershipServiceApi
}

type MembershipServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewMembershipApi(apiClient *client.ApiClient) *MembershipApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &MembershipApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewMembershipServiceApi(a.ApiClient)

	return a
}

func NewMembershipServiceApi(apiClient *client.ApiClient) *MembershipServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &MembershipServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Create a membership for a user within a single domain. This operation is asynchronous and returns a membership reference.
func (api *MembershipApi) CreateDomainMembership(body *import1.Membership, args ...map[string]interface{}) (*import1.CreateDomainMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDomainMembership(context.Background(), &import2.CreateDomainMembershipRequest{
		Body: body,
	}, args...)
}

// Create a membership for a user within a single domain. This operation is asynchronous and returns a membership reference.
func (api *MembershipServiceApi) CreateDomainMembership(ctx context.Context, request *import2.CreateDomainMembershipRequest, args ...map[string]interface{}) (*import1.CreateDomainMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domain-memberships"

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
	unmarshalledResp := new(import1.CreateDomainMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Delete an existing membership and revoke access for the associated domain. This operation is asynchronous.
func (api *MembershipApi) DeleteDomainMembershipById(extId *string, args ...map[string]interface{}) (*import1.DeleteDomainMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteDomainMembershipById(context.Background(), &import2.DeleteDomainMembershipByIdRequest{
		ExtId: extId,
	}, args...)
}

// Delete an existing membership and revoke access for the associated domain. This operation is asynchronous.
func (api *MembershipServiceApi) DeleteDomainMembershipById(ctx context.Context, request *import2.DeleteDomainMembershipByIdRequest, args ...map[string]interface{}) (*import1.DeleteDomainMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domain-memberships/{extId}"

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
	unmarshalledResp := new(import1.DeleteDomainMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve a specific membership by its identifier.
func (api *MembershipApi) GetDomainMembershipById(extId *string, args ...map[string]interface{}) (*import1.GetDomainMembershipApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDomainMembershipById(context.Background(), &import2.GetDomainMembershipByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve a specific membership by its identifier.
func (api *MembershipServiceApi) GetDomainMembershipById(ctx context.Context, request *import2.GetDomainMembershipByIdRequest, args ...map[string]interface{}) (*import1.GetDomainMembershipApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domain-memberships/{extId}"

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
	unmarshalledResp := new(import1.GetDomainMembershipApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// List all memberships for the tenant in a flat structure across all domains.
func (api *MembershipApi) ListDomainMemberships(page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import1.ListDomainMembershipsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMembershipServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDomainMemberships(context.Background(), &import2.ListDomainMembershipsRequest{
		Page_:   page_,
		Limit_:  limit_,
		Filter_: filter_,
	}, args...)
}

// List all memberships for the tenant in a flat structure across all domains.
func (api *MembershipServiceApi) ListDomainMemberships(ctx context.Context, request *import2.ListDomainMembershipsRequest, args ...map[string]interface{}) (*import1.ListDomainMembershipsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/tenancy/v4.0.a1/config/domain-memberships"

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
	unmarshalledResp := new(import1.ListDomainMembershipsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
