package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/request/iscsiclients"
	"net/http"
	"net/url"
	"strings"
)

type IscsiClientsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *IscsiClientsServiceApi
}

type IscsiClientsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewIscsiClientsApi(apiClient *client.ApiClient) *IscsiClientsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &IscsiClientsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewIscsiClientsServiceApi(a.ApiClient)

	return a
}

func NewIscsiClientsServiceApi(apiClient *client.ApiClient) *IscsiClientsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &IscsiClientsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches the iSCSI client details identified by {extId}.
func (api *IscsiClientsApi) GetIscsiClientById(extId *string, args ...map[string]interface{}) (*import1.GetIscsiClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewIscsiClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetIscsiClientById(context.Background(), &import2.GetIscsiClientByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the iSCSI client details identified by {extId}.
func (api *IscsiClientsServiceApi) GetIscsiClientById(ctx context.Context, request *import2.GetIscsiClientByIdRequest, args ...map[string]interface{}) (*import1.GetIscsiClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/iscsi-clients/{extId}"

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

	unmarshalledResp := new(import1.GetIscsiClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the list of iSCSI clients.
func (api *IscsiClientsApi) ListIscsiClients(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListIscsiClientsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewIscsiClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListIscsiClients(context.Background(), &import2.ListIscsiClientsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Fetches the list of iSCSI clients.
func (api *IscsiClientsServiceApi) ListIscsiClients(ctx context.Context, request *import2.ListIscsiClientsRequest, args ...map[string]interface{}) (*import1.ListIscsiClientsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/iscsi-clients"

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

	unmarshalledResp := new(import1.ListIscsiClientsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Modifies the details of an existing iSCSI client configuration identified by {extId}.
func (api *IscsiClientsApi) UpdateIscsiClientById(extId *string, body *import1.IscsiClient, args ...map[string]interface{}) (*import1.UpdateIscsiClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewIscsiClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateIscsiClientById(context.Background(), &import2.UpdateIscsiClientByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Modifies the details of an existing iSCSI client configuration identified by {extId}.
func (api *IscsiClientsServiceApi) UpdateIscsiClientById(ctx context.Context, request *import2.UpdateIscsiClientByIdRequest, args ...map[string]interface{}) (*import1.UpdateIscsiClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/iscsi-clients/{extId}"

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

	unmarshalledResp := new(import1.UpdateIscsiClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
