package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import11 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/cvms"
	"net/http"
	"net/url"
	"strings"
)

type CvmsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *CvmsServiceApi
}

type CvmsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCvmsApi(apiClient *client.ApiClient) *CvmsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CvmsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewCvmsServiceApi(a.ApiClient)

	return a
}

func NewCvmsServiceApi(apiClient *client.ApiClient) *CvmsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CvmsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the details of the CVMs identified by {extId} belonging to the cluster identified by {clusterExtId}.
func (api *CvmsApi) GetCvmById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetCvmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCvmsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetCvmById(context.Background(), &import11.GetCvmByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get the details of the CVMs identified by {extId} belonging to the cluster identified by {clusterExtId}.
func (api *CvmsServiceApi) GetCvmById(ctx context.Context, request *import11.GetCvmByIdRequest, args ...map[string]interface{}) (*import1.GetCvmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/cvms/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.GetCvmApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get the list of all the CVMs belonging to the cluster identified by {clusterExtId}.
func (api *CvmsApi) ListCvmsbyClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCvmsByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCvmsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListCvmsbyClusterId(context.Background(), &import11.ListCvmsbyClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// Get the list of all the CVMs belonging to the cluster identified by {clusterExtId}.
func (api *CvmsServiceApi) ListCvmsbyClusterId(ctx context.Context, request *import11.ListCvmsbyClusterIdRequest, args ...map[string]interface{}) (*import1.ListCvmsByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/cvms"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.ListCvmsByClusterIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Reconfigure CVMs within a cluster identified by {clusterExtId}.
func (api *CvmsApi) ReconfigureCvms(clusterExtId *string, body *import1.CvmReconfigurationSpec, args ...map[string]interface{}) (*import1.ReconfigureCvmsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCvmsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReconfigureCvms(context.Background(), &import11.ReconfigureCvmsRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Reconfigure CVMs within a cluster identified by {clusterExtId}.
func (api *CvmsServiceApi) ReconfigureCvms(ctx context.Context, request *import11.ReconfigureCvmsRequest, args ...map[string]interface{}) (*import1.ReconfigureCvmsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/cvms/$actions/reconfigure"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.ReconfigureCvmsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Update the details of the CVM identified by {extId} belonging to the cluster identified by {clusterExtId}
func (api *CvmsApi) UpdateCvmById(clusterExtId *string, extId *string, body *import1.Cvm, args ...map[string]interface{}) (*import1.UpdateCvmApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewCvmsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateCvmById(context.Background(), &import11.UpdateCvmByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Update the details of the CVM identified by {extId} belonging to the cluster identified by {clusterExtId}
func (api *CvmsServiceApi) UpdateCvmById(ctx context.Context, request *import11.UpdateCvmByIdRequest, args ...map[string]interface{}) (*import1.UpdateCvmApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/config/clusters/{clusterExtId}/cvms/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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
	unmarshalledResp := new(import1.UpdateCvmApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
