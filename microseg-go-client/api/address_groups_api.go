package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/request/addressgroups"
	"net/http"
	"net/url"
	"strings"
)

type AddressGroupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *AddressGroupsServiceApi
}

type AddressGroupsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAddressGroupsApi(apiClient *client.ApiClient) *AddressGroupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AddressGroupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewAddressGroupsServiceApi(a.ApiClient)

	return a
}

func NewAddressGroupsServiceApi(apiClient *client.ApiClient) *AddressGroupsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AddressGroupsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an Address Group.
func (api *AddressGroupsApi) CreateAddressGroup(body *import1.AddressGroup, args ...map[string]interface{}) (*import1.CreateAddressGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAddressGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateAddressGroup(context.Background(), &import2.CreateAddressGroupRequest{
		Body: body,
	}, args...)
}

// Creates an Address Group.
func (api *AddressGroupsServiceApi) CreateAddressGroup(ctx context.Context, request *import2.CreateAddressGroupRequest, args ...map[string]interface{}) (*import1.CreateAddressGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/address-groups"

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

	unmarshalledResp := new(import1.CreateAddressGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the Address Group with the provided ExtID.
func (api *AddressGroupsApi) DeleteAddressGroupById(extId *string, args ...map[string]interface{}) (*import1.DeleteAddressGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAddressGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteAddressGroupById(context.Background(), &import2.DeleteAddressGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the Address Group with the provided ExtID.
func (api *AddressGroupsServiceApi) DeleteAddressGroupById(ctx context.Context, request *import2.DeleteAddressGroupByIdRequest, args ...map[string]interface{}) (*import1.DeleteAddressGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/address-groups/{extId}"

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

	unmarshalledResp := new(import1.DeleteAddressGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets the Address Group with the provided ExtID.
func (api *AddressGroupsApi) GetAddressGroupById(extId *string, args ...map[string]interface{}) (*import1.GetAddressGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAddressGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetAddressGroupById(context.Background(), &import2.GetAddressGroupByIdRequest{
		ExtId: extId,
	}, args...)
}

// Gets the Address Group with the provided ExtID.
func (api *AddressGroupsServiceApi) GetAddressGroupById(ctx context.Context, request *import2.GetAddressGroupByIdRequest, args ...map[string]interface{}) (*import1.GetAddressGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/address-groups/{extId}"

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

	unmarshalledResp := new(import1.GetAddressGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Gets a list of Address Groups.
func (api *AddressGroupsApi) ListAddressGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListAddressGroupsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAddressGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAddressGroups(context.Background(), &import2.ListAddressGroupsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Gets a list of Address Groups.
func (api *AddressGroupsServiceApi) ListAddressGroups(ctx context.Context, request *import2.ListAddressGroupsRequest, args ...map[string]interface{}) (*import1.ListAddressGroupsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/address-groups"

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

	unmarshalledResp := new(import1.ListAddressGroupsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the Address Group with the provided ExtID.
func (api *AddressGroupsApi) UpdateAddressGroupById(extId *string, body *import1.AddressGroup, args ...map[string]interface{}) (*import1.UpdateAddressGroupApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAddressGroupsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateAddressGroupById(context.Background(), &import2.UpdateAddressGroupByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the Address Group with the provided ExtID.
func (api *AddressGroupsServiceApi) UpdateAddressGroupById(ctx context.Context, request *import2.UpdateAddressGroupByIdRequest, args ...map[string]interface{}) (*import1.UpdateAddressGroupApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/microseg/v4.2/config/address-groups/{extId}"

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

	unmarshalledResp := new(import1.UpdateAddressGroupApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
