package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/request/entitysyncpolicies"
	"net/http"
	"net/url"
	"strings"
)

type EntitySyncPoliciesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EntitySyncPoliciesServiceApi
}

type EntitySyncPoliciesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEntitySyncPoliciesApi(apiClient *client.ApiClient) *EntitySyncPoliciesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntitySyncPoliciesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEntitySyncPoliciesServiceApi(a.ApiClient)

	return a
}

func NewEntitySyncPoliciesServiceApi(apiClient *client.ApiClient) *EntitySyncPoliciesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EntitySyncPoliciesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches an entity sync policy identified by its external identifier.
func (api *EntitySyncPoliciesApi) GetEntitySyncPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetEntitySyncPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitySyncPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEntitySyncPolicyById(context.Background(), &import2.GetEntitySyncPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches an entity sync policy identified by its external identifier.
func (api *EntitySyncPoliciesServiceApi) GetEntitySyncPolicyById(ctx context.Context, request *import2.GetEntitySyncPolicyByIdRequest, args ...map[string]interface{}) (*import1.GetEntitySyncPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/entity-sync-policies/{extId}"

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

	unmarshalledResp := new(import1.GetEntitySyncPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists entity sync policies.
func (api *EntitySyncPoliciesApi) ListEntitySyncPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListEntitySyncPoliciesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitySyncPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEntitySyncPolicies(context.Background(), &import2.ListEntitySyncPoliciesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists entity sync policies.
func (api *EntitySyncPoliciesServiceApi) ListEntitySyncPolicies(ctx context.Context, request *import2.ListEntitySyncPoliciesRequest, args ...map[string]interface{}) (*import1.ListEntitySyncPoliciesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/entity-sync-policies"

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

	unmarshalledResp := new(import1.ListEntitySyncPoliciesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Syncs the entity identified by  its external identifier to a remote domain manager. It overrides the entity on the remote domain manager in the event of a conflict.
func (api *EntitySyncPoliciesApi) SyncEntitySyncPolicyById(extId *string, args ...map[string]interface{}) (*import1.SyncEntitySyncPolicyApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEntitySyncPoliciesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.SyncEntitySyncPolicyById(context.Background(), &import2.SyncEntitySyncPolicyByIdRequest{
		ExtId: extId,
	}, args...)
}

// Syncs the entity identified by  its external identifier to a remote domain manager. It overrides the entity on the remote domain manager in the event of a conflict.
func (api *EntitySyncPoliciesServiceApi) SyncEntitySyncPolicyById(ctx context.Context, request *import2.SyncEntitySyncPolicyByIdRequest, args ...map[string]interface{}) (*import1.SyncEntitySyncPolicyApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/datapolicies/v4.2/config/entity-sync-policies/{extId}/$actions/sync-entity"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.SyncEntitySyncPolicyApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
