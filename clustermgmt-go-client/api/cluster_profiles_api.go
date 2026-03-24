package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/clusterprofiles"
	"net/http"
	"net/url"
	"strings"
)

type ClusterProfilesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClusterProfilesServiceApi
}

type ClusterProfilesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClusterProfilesApi(apiClient *client.ApiClient) *ClusterProfilesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterProfilesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClusterProfilesServiceApi(a.ApiClient)

	return a
}

func NewClusterProfilesServiceApi(apiClient *client.ApiClient) *ClusterProfilesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterProfilesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Applies or associates clusters to a profile. The clusters will be added for application and monitoring.
func (api *ClusterProfilesApi) ApplyClusterProfile(extId *string, body *import1.ClusterReferenceListSpec, dryrun_ *bool, args ...map[string]interface{}) (*import1.ApplyClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ApplyClusterProfile(context.Background(), &import3.ApplyClusterProfileRequest{
		ExtId:   extId,
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Applies or associates clusters to a profile. The clusters will be added for application and monitoring.
func (api *ClusterProfilesServiceApi) ApplyClusterProfile(ctx context.Context, request *import3.ApplyClusterProfileRequest, args ...map[string]interface{}) (*import1.ApplyClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles/{extId}/$actions/apply"

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

	unmarshalledResp := new(import1.ApplyClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a cluster profile with the settings provided in the request body.
func (api *ClusterProfilesApi) CreateClusterProfile(body *import1.ClusterProfile, args ...map[string]interface{}) (*import1.CreateClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateClusterProfile(context.Background(), &import3.CreateClusterProfileRequest{
		Body: body,
	}, args...)
}

// Creates a cluster profile with the settings provided in the request body.
func (api *ClusterProfilesServiceApi) CreateClusterProfile(ctx context.Context, request *import3.CreateClusterProfileRequest, args ...map[string]interface{}) (*import1.CreateClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles"

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

	unmarshalledResp := new(import1.CreateClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesApi) DeleteClusterProfileById(extId *string, args ...map[string]interface{}) (*import1.DeleteClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteClusterProfileById(context.Background(), &import3.DeleteClusterProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesServiceApi) DeleteClusterProfileById(ctx context.Context, request *import3.DeleteClusterProfileByIdRequest, args ...map[string]interface{}) (*import1.DeleteClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles/{extId}"

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

	unmarshalledResp := new(import1.DeleteClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Disassociate cluster from a cluster profile. This will halt the application and monitoring of the clusters.
func (api *ClusterProfilesApi) DisassociateClusterFromClusterProfile(extId *string, body *import1.ClusterReferenceListSpec, args ...map[string]interface{}) (*import1.DisassociateClusterFromClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisassociateClusterFromClusterProfile(context.Background(), &import3.DisassociateClusterFromClusterProfileRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Disassociate cluster from a cluster profile. This will halt the application and monitoring of the clusters.
func (api *ClusterProfilesServiceApi) DisassociateClusterFromClusterProfile(ctx context.Context, request *import3.DisassociateClusterFromClusterProfileRequest, args ...map[string]interface{}) (*import1.DisassociateClusterFromClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles/{extId}/$actions/disassociate-cluster"

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

	unmarshalledResp := new(import1.DisassociateClusterFromClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesApi) GetClusterProfileById(extId *string, args ...map[string]interface{}) (*import1.GetClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetClusterProfileById(context.Background(), &import3.GetClusterProfileByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesServiceApi) GetClusterProfileById(ctx context.Context, request *import3.GetClusterProfileByIdRequest, args ...map[string]interface{}) (*import1.GetClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles/{extId}"

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

	unmarshalledResp := new(import1.GetClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of cluster profile entities. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesApi) ListClusterProfiles(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListClusterProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListClusterProfiles(context.Background(), &import3.ListClusterProfilesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of cluster profile entities. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesServiceApi) ListClusterProfiles(ctx context.Context, request *import3.ListClusterProfilesRequest, args ...map[string]interface{}) (*import1.ListClusterProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles"

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

	unmarshalledResp := new(import1.ListClusterProfilesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesApi) UpdateClusterProfileById(extId *string, body *import1.ClusterProfile, dryrun_ *bool, args ...map[string]interface{}) (*import1.UpdateClusterProfileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClusterProfilesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateClusterProfileById(context.Background(), &import3.UpdateClusterProfileByIdRequest{
		ExtId:   extId,
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Updates a cluster profile. A profile consists of different cluster settings like Network Time Protocol(NTP), Domain Name System(DNS), and so on.
func (api *ClusterProfilesServiceApi) UpdateClusterProfileById(ctx context.Context, request *import3.UpdateClusterProfileByIdRequest, args ...map[string]interface{}) (*import1.UpdateClusterProfileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/cluster-profiles/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPut, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateClusterProfileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
