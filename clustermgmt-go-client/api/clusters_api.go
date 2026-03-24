package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/operations"
	import7 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/clusters"
	import5 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/stats"
	import6 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/stats"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ClustersApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ClustersServiceApi
}

type ClustersServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewClustersApi(apiClient *client.ApiClient) *ClustersApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClustersApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewClustersServiceApi(a.ApiClient)

	return a
}

func NewClustersServiceApi(apiClient *client.ApiClient) *ClustersServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClustersServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Adds transport ports and protocol details to the SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) AddSnmpTransport(clusterExtId *string, body *import1.SnmpTransport, args ...map[string]interface{}) (*import1.AddSnmpTransportsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AddSnmpTransport(context.Background(), &import7.AddSnmpTransportRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Adds transport ports and protocol details to the SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) AddSnmpTransport(ctx context.Context, request *import7.AddSnmpTransportRequest, args ...map[string]interface{}) (*import1.AddSnmpTransportsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/$actions/add-transports"

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

	unmarshalledResp := new(import1.AddSnmpTransportsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Attach categories to the cluster identified by {clusterExtId}.
func (api *ClustersApi) AssociateCategoriesToCluster(clusterExtId *string, body *import1.CategoryEntityReferences, args ...map[string]interface{}) (*import1.AssociateCategoriesToClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.AssociateCategoriesToCluster(context.Background(), &import7.AssociateCategoriesToClusterRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Attach categories to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) AssociateCategoriesToCluster(ctx context.Context, request *import7.AssociateCategoriesToClusterRequest, args ...map[string]interface{}) (*import1.AssociateCategoriesToClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/associate-categories"

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

	unmarshalledResp := new(import1.AssociateCategoriesToClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Provides information on whether hypervisor ISO upload is required or not. This API is not supported for XEN hypervisor type.
func (api *ClustersApi) CheckHypervisorRequirements(clusterExtId *string, body *import1.HypervisorUploadParam, args ...map[string]interface{}) (*import1.CheckHypervisorRequirementsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CheckHypervisorRequirements(context.Background(), &import7.CheckHypervisorRequirementsRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Provides information on whether hypervisor ISO upload is required or not. This API is not supported for XEN hypervisor type.
func (api *ClustersServiceApi) CheckHypervisorRequirements(ctx context.Context, request *import7.CheckHypervisorRequirementsRequest, args ...map[string]interface{}) (*import1.CheckHypervisorRequirementsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/check-hypervisor-requirements"

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

	unmarshalledResp := new(import1.CheckHypervisorRequirementsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Cluster create operation.
func (api *ClustersApi) CreateCluster(body *import1.Cluster, dryrun_ *bool, args ...map[string]interface{}) (*import1.CreateClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateCluster(context.Background(), &import7.CreateClusterRequest{
		Body:    body,
		Dryrun_: dryrun_,
	}, args...)
}

// Cluster create operation.
func (api *ClustersServiceApi) CreateCluster(ctx context.Context, request *import7.CreateClusterRequest, args ...map[string]interface{}) (*import1.CreateClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters"

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

	unmarshalledResp := new(import1.CreateClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Adds RSYSLOG server configuration to the cluster identified by {clusterExtId}.
func (api *ClustersApi) CreateRsyslogServer(clusterExtId *string, body *import1.RsyslogServer, args ...map[string]interface{}) (*import1.CreateRsyslogServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateRsyslogServer(context.Background(), &import7.CreateRsyslogServerRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Adds RSYSLOG server configuration to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) CreateRsyslogServer(ctx context.Context, request *import7.CreateRsyslogServerRequest, args ...map[string]interface{}) (*import1.CreateRsyslogServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rsyslog-servers"

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

	unmarshalledResp := new(import1.CreateRsyslogServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Add SNMP trap configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) CreateSnmpTrap(clusterExtId *string, body *import1.SnmpTrap, args ...map[string]interface{}) (*import1.CreateSnmpTrapApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateSnmpTrap(context.Background(), &import7.CreateSnmpTrapRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Add SNMP trap configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) CreateSnmpTrap(ctx context.Context, request *import7.CreateSnmpTrapRequest, args ...map[string]interface{}) (*import1.CreateSnmpTrapApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/traps"

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

	unmarshalledResp := new(import1.CreateSnmpTrapApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Adds SNMP user configuration to the cluster identified by {clusterExtId}.
func (api *ClustersApi) CreateSnmpUser(clusterExtId *string, body *import1.SnmpUser, args ...map[string]interface{}) (*import1.CreateSnmpUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateSnmpUser(context.Background(), &import7.CreateSnmpUserRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Adds SNMP user configuration to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) CreateSnmpUser(ctx context.Context, request *import7.CreateSnmpUserRequest, args ...map[string]interface{}) (*import1.CreateSnmpUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/users"

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

	unmarshalledResp := new(import1.CreateSnmpUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the cluster identified by {extId}.
func (api *ClustersApi) DeleteClusterById(extId *string, dryrun_ *bool, args ...map[string]interface{}) (*import1.DeleteClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteClusterById(context.Background(), &import7.DeleteClusterByIdRequest{
		ExtId:   extId,
		Dryrun_: dryrun_,
	}, args...)
}

// Deletes the cluster identified by {extId}.
func (api *ClustersServiceApi) DeleteClusterById(ctx context.Context, request *import7.DeleteClusterByIdRequest, args ...map[string]interface{}) (*import1.DeleteClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes RSYSLOG server configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) DeleteRsyslogServerById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteRsyslogServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteRsyslogServerById(context.Background(), &import7.DeleteRsyslogServerByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Deletes RSYSLOG server configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) DeleteRsyslogServerById(ctx context.Context, request *import7.DeleteRsyslogServerByIdRequest, args ...map[string]interface{}) (*import1.DeleteRsyslogServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rsyslog-servers/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteRsyslogServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes SNMP trap configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) DeleteSnmpTrapById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteSnmpTrapApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteSnmpTrapById(context.Background(), &import7.DeleteSnmpTrapByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Deletes SNMP trap configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) DeleteSnmpTrapById(ctx context.Context, request *import7.DeleteSnmpTrapByIdRequest, args ...map[string]interface{}) (*import1.DeleteSnmpTrapApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/traps/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteSnmpTrapApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes SNMP user configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) DeleteSnmpUserById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.DeleteSnmpUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteSnmpUserById(context.Background(), &import7.DeleteSnmpUserByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Deletes SNMP user configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) DeleteSnmpUserById(ctx context.Context, request *import7.DeleteSnmpUserByIdRequest, args ...map[string]interface{}) (*import1.DeleteSnmpUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/users/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteSnmpUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Detach categories from the cluster identified by {clusterExtId}.
func (api *ClustersApi) DisassociateCategoriesFromCluster(clusterExtId *string, body *import1.CategoryEntityReferences, args ...map[string]interface{}) (*import1.DisassociateCategoriesFromClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DisassociateCategoriesFromCluster(context.Background(), &import7.DisassociateCategoriesFromClusterRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Detach categories from the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) DisassociateCategoriesFromCluster(ctx context.Context, request *import7.DisassociateCategoriesFromClusterRequest, args ...map[string]interface{}) (*import1.DisassociateCategoriesFromClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/disassociate-categories"

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

	unmarshalledResp := new(import1.DisassociateCategoriesFromClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the unconfigured node details such as node UUID, node position, node IP, foundation version and more.
func (api *ClustersApi) DiscoverUnconfiguredNodes(clusterExtId *string, body *import1.NodeDiscoveryParams, args ...map[string]interface{}) (*import1.DiscoverUnconfiguredNodesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DiscoverUnconfiguredNodes(context.Background(), &import7.DiscoverUnconfiguredNodesRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Get the unconfigured node details such as node UUID, node position, node IP, foundation version and more.
func (api *ClustersServiceApi) DiscoverUnconfiguredNodes(ctx context.Context, request *import7.DiscoverUnconfiguredNodesRequest, args ...map[string]interface{}) (*import1.DiscoverUnconfiguredNodesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/discover-unconfigured-nodes"

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

	unmarshalledResp := new(import1.DiscoverUnconfiguredNodesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Enter host identified by {extId} into maintenance mode belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) EnterHostMaintenance(clusterExtId *string, extId *string, body *import4.EnterHostMaintenanceSpec, args ...map[string]interface{}) (*import4.EnterHostMaintenanceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.EnterHostMaintenance(context.Background(), &import7.EnterHostMaintenanceRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Enter host identified by {extId} into maintenance mode belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) EnterHostMaintenance(ctx context.Context, request *import7.EnterHostMaintenanceRequest, args ...map[string]interface{}) (*import4.EnterHostMaintenanceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/operations/clusters/{clusterExtId}/hosts/{extId}/$actions/enter-host-maintenance"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.EnterHostMaintenanceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Exit host identified by {extId} from maintenance mode belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) ExitHostMaintenance(clusterExtId *string, extId *string, body *import4.HostMaintenanceCommonSpec, args ...map[string]interface{}) (*import4.ExitHostMaintenanceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExitHostMaintenance(context.Background(), &import7.ExitHostMaintenanceRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Exit host identified by {extId} from maintenance mode belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ExitHostMaintenance(ctx context.Context, request *import7.ExitHostMaintenanceRequest, args ...map[string]interface{}) (*import4.ExitHostMaintenanceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/operations/clusters/{clusterExtId}/hosts/{extId}/$actions/exit-host-maintenance"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import4.ExitHostMaintenanceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Add node on a cluster. This API is not supported for XEN hypervisor type.
func (api *ClustersApi) ExpandCluster(clusterExtId *string, body *import1.ExpandClusterParams, args ...map[string]interface{}) (*import1.ExpandClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ExpandCluster(context.Background(), &import7.ExpandClusterRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Add node on a cluster. This API is not supported for XEN hypervisor type.
func (api *ClustersServiceApi) ExpandCluster(ctx context.Context, request *import7.ExpandClusterRequest, args ...map[string]interface{}) (*import1.ExpandClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/expand-cluster"

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

	unmarshalledResp := new(import1.ExpandClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a dictionary of cluster networks and available uplinks on the given nodes. This API is not supported for XEN hypervisor type.
func (api *ClustersApi) FetchNodeNetworkingDetails(clusterExtId *string, body *import1.NodeDetails, args ...map[string]interface{}) (*import1.FetchNodeNetworkingDetailsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.FetchNodeNetworkingDetails(context.Background(), &import7.FetchNodeNetworkingDetailsRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Get a dictionary of cluster networks and available uplinks on the given nodes. This API is not supported for XEN hypervisor type.
func (api *ClustersServiceApi) FetchNodeNetworkingDetails(ctx context.Context, request *import7.FetchNodeNetworkingDetailsRequest, args ...map[string]interface{}) (*import1.FetchNodeNetworkingDetailsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/fetch-node-networking-details"

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

	unmarshalledResp := new(import1.FetchNodeNetworkingDetailsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get task response based on the type of request.
func (api *ClustersApi) FetchTaskResponse(extId *string, taskResponseType *import1.TaskResponseType, args ...map[string]interface{}) (*import1.FetchTaskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.FetchTaskResponse(context.Background(), &import7.FetchTaskResponseRequest{
		ExtId:            extId,
		TaskResponseType: taskResponseType,
	}, args...)
}

// Get task response based on the type of request.
func (api *ClustersServiceApi) FetchTaskResponse(ctx context.Context, request *import7.FetchTaskResponseRequest, args ...map[string]interface{}) (*import1.FetchTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/task-response/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'taskResponseType' is set
	if nil == request.TaskResponseType {
		return nil, client.ReportError("taskResponseType is required and must be specified")
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

	// Query Params
	taskResponseTypeQueryParamEnumVal := request.TaskResponseType.GetName()
	queryParams.Add("taskResponseType", client.ParameterToString(taskResponseTypeQueryParamEnumVal, ""))
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

	unmarshalledResp := new(import1.FetchTaskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the cluster entity details identified by {extId}.
func (api *ClustersApi) GetClusterById(extId *string, expand_ *string, args ...map[string]interface{}) (*import1.GetClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetClusterById(context.Background(), &import7.GetClusterByIdRequest{
		ExtId:   extId,
		Expand_: expand_,
	}, args...)
}

// Fetches the cluster entity details identified by {extId}.
func (api *ClustersServiceApi) GetClusterById(ctx context.Context, request *import7.GetClusterByIdRequest, args ...map[string]interface{}) (*import1.GetClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{extId}"

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

	// Query Params
	if request.Expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*request.Expand_, ""))
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

	unmarshalledResp := new(import1.GetClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the statistics data of the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetClusterStats(extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import6.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import5.ClusterStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetClusterStats(context.Background(), &import7.GetClusterStatsRequest{
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Get the statistics data of the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetClusterStats(ctx context.Context, request *import7.GetClusterStatsRequest, args ...map[string]interface{}) (*import5.ClusterStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/stats/clusters/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
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

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.ClusterStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the details of the host identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetHostById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetHostApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetHostById(context.Background(), &import7.GetHostByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Fetches the details of the host identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetHostById(ctx context.Context, request *import7.GetHostByIdRequest, args ...map[string]interface{}) (*import1.GetHostApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{extId}"

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

	unmarshalledResp := new(import1.GetHostApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the host NIC entity of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetHostNicById(clusterExtId *string, hostExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetHostNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetHostNicById(context.Background(), &import7.GetHostNicByIdRequest{
		ClusterExtId: clusterExtId,
		HostExtId:    hostExtId,
		ExtId:        extId,
	}, args...)
}

// Get the host NIC entity of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetHostNicById(ctx context.Context, request *import7.GetHostNicByIdRequest, args ...map[string]interface{}) (*import1.GetHostNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-nics/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == request.HostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*request.HostExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetHostNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the statistics data of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetHostStats(clusterExtId *string, extId *string, startTime_ *time.Time, endTime_ *time.Time, samplingInterval_ *int, statType_ *import6.DownSamplingOperator, select_ *string, args ...map[string]interface{}) (*import5.HostStatsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetHostStats(context.Background(), &import7.GetHostStatsRequest{
		ClusterExtId:      clusterExtId,
		ExtId:             extId,
		StartTime_:        startTime_,
		EndTime_:          endTime_,
		SamplingInterval_: samplingInterval_,
		StatType_:         statType_,
		Select_:           select_,
	}, args...)
}

// Get the statistics data of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetHostStats(ctx context.Context, request *import7.GetHostStatsRequest, args ...map[string]interface{}) (*import5.HostStatsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/stats/clusters/{clusterExtId}/hosts/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'startTime_' is set
	if nil == request.StartTime_ {
		return nil, client.ReportError("startTime_ is required and must be specified")
	}
	// verify the required parameter 'endTime_' is set
	if nil == request.EndTime_ {
		return nil, client.ReportError("endTime_ is required and must be specified")
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

	// Query Params
	queryParams.Add("$startTime", client.ParameterToString(*request.StartTime_, ""))
	queryParams.Add("$endTime", client.ParameterToString(*request.EndTime_, ""))
	if request.SamplingInterval_ != nil {
		queryParams.Add("$samplingInterval", client.ParameterToString(*request.SamplingInterval_, ""))
	}
	if request.StatType_ != nil {
		statType_QueryParamEnumVal := request.StatType_.GetName()
		queryParams.Add("$statType", client.ParameterToString(statType_QueryParamEnumVal, ""))
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

	unmarshalledResp := new(import5.HostStatsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the rackable unit entity details identified by {extId}  of the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetRackableUnitById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetRackableUnitApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRackableUnitById(context.Background(), &import7.GetRackableUnitByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Fetches the rackable unit entity details identified by {extId}  of the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetRackableUnitById(ctx context.Context, request *import7.GetRackableUnitByIdRequest, args ...map[string]interface{}) (*import1.GetRackableUnitApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rackable-units/{extId}"

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

	unmarshalledResp := new(import1.GetRackableUnitApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the RSYSLOG server configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetRsyslogServerById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetRsyslogServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRsyslogServerById(context.Background(), &import7.GetRsyslogServerByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Fetches the RSYSLOG server configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetRsyslogServerById(ctx context.Context, request *import7.GetRsyslogServerByIdRequest, args ...map[string]interface{}) (*import1.GetRsyslogServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rsyslog-servers/{extId}"

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

	unmarshalledResp := new(import1.GetRsyslogServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches SNMP config details of the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetSnmpConfigByClusterId(clusterExtId *string, args ...map[string]interface{}) (*import1.GetSnmpConfigByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSnmpConfigByClusterId(context.Background(), &import7.GetSnmpConfigByClusterIdRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// Fetches SNMP config details of the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetSnmpConfigByClusterId(ctx context.Context, request *import7.GetSnmpConfigByClusterIdRequest, args ...map[string]interface{}) (*import1.GetSnmpConfigByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp"

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

	unmarshalledResp := new(import1.GetSnmpConfigByClusterIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

//  Fetches SNMP trap configuration details identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetSnmpTrapById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetSnmpTrapApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSnmpTrapById(context.Background(), &import7.GetSnmpTrapByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

//  Fetches SNMP trap configuration details identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetSnmpTrapById(ctx context.Context, request *import7.GetSnmpTrapByIdRequest, args ...map[string]interface{}) (*import1.GetSnmpTrapApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/traps/{extId}"

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

	unmarshalledResp := new(import1.GetSnmpTrapApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches SNMP user configuration details identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetSnmpUserById(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetSnmpUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSnmpUserById(context.Background(), &import7.GetSnmpUserByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Fetches SNMP user configuration details identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetSnmpUserById(ctx context.Context, request *import7.GetSnmpUserByIdRequest, args ...map[string]interface{}) (*import1.GetSnmpUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/users/{extId}"

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

	unmarshalledResp := new(import1.GetSnmpUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the virtual NIC entity of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) GetVirtualNicById(clusterExtId *string, hostExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetVirtualNicApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetVirtualNicById(context.Background(), &import7.GetVirtualNicByIdRequest{
		ClusterExtId: clusterExtId,
		HostExtId:    hostExtId,
		ExtId:        extId,
	}, args...)
}

// Get the virtual NIC entity of the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) GetVirtualNicById(ctx context.Context, request *import7.GetVirtualNicByIdRequest, args ...map[string]interface{}) (*import1.GetVirtualNicApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{hostExtId}/virtual-nics/{extId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == request.HostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*request.HostExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetVirtualNicApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all cluster entities registered to Prism Central.
func (api *ClustersApi) ListClusters(page_ *int, limit_ *int, filter_ *string, orderby_ *string, apply_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListClustersApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListClusters(context.Background(), &import7.ListClustersRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Apply_:   apply_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Lists all cluster entities registered to Prism Central.
func (api *ClustersServiceApi) ListClusters(ctx context.Context, request *import7.ListClustersRequest, args ...map[string]interface{}) (*import1.ListClustersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters"

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
	if request.Apply_ != nil {
		queryParams.Add("$apply", client.ParameterToString(*request.Apply_, ""))
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

	unmarshalledResp := new(import1.ListClustersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of all host NICs for all the clusters attached to the PC.
func (api *ClustersApi) ListHostNics(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListHostNicsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListHostNics(context.Background(), &import7.ListHostNicsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of all host NICs for all the clusters attached to the PC.
func (api *ClustersServiceApi) ListHostNics(ctx context.Context, request *import7.ListHostNicsRequest, args ...map[string]interface{}) (*import1.ListHostNicsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/host-nics"

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

	unmarshalledResp := new(import1.ListHostNicsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all host NICs for the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) ListHostNicsByHostId(clusterExtId *string, hostExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListHostNicsByHostIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListHostNicsByHostId(context.Background(), &import7.ListHostNicsByHostIdRequest{
		ClusterExtId: clusterExtId,
		HostExtId:    hostExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// Lists all host NICs for the host identified by {hostExtId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ListHostNicsByHostId(ctx context.Context, request *import7.ListHostNicsByHostIdRequest, args ...map[string]interface{}) (*import1.ListHostNicsByHostIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-nics"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == request.HostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*request.HostExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListHostNicsByHostIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all host entities across clusters registered to Prism Central.
func (api *ClustersApi) ListHosts(page_ *int, limit_ *int, filter_ *string, orderby_ *string, apply_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListHostsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListHosts(context.Background(), &import7.ListHostsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Apply_:   apply_,
		Select_:  select_,
	}, args...)
}

// Lists all host entities across clusters registered to Prism Central.
func (api *ClustersServiceApi) ListHosts(ctx context.Context, request *import7.ListHostsRequest, args ...map[string]interface{}) (*import1.ListHostsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/hosts"

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
	if request.Apply_ != nil {
		queryParams.Add("$apply", client.ParameterToString(*request.Apply_, ""))
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

	unmarshalledResp := new(import1.ListHostsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the hosts associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) ListHostsByClusterId(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, apply_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListHostsByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListHostsByClusterId(context.Background(), &import7.ListHostsByClusterIdRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Apply_:       apply_,
		Select_:      select_,
	}, args...)
}

// Lists all the hosts associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ListHostsByClusterId(ctx context.Context, request *import7.ListHostsByClusterIdRequest, args ...map[string]interface{}) (*import1.ListHostsByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts"

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
	if request.Apply_ != nil {
		queryParams.Add("$apply", client.ParameterToString(*request.Apply_, ""))
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

	unmarshalledResp := new(import1.ListHostsByClusterIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Physical GPU profiles.
func (api *ClustersApi) ListPhysicalGpuProfiles(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.ListPhysicalGpuProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListPhysicalGpuProfiles(context.Background(), &import7.ListPhysicalGpuProfilesRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
	}, args...)
}

// List Physical GPU profiles.
func (api *ClustersServiceApi) ListPhysicalGpuProfiles(ctx context.Context, request *import7.ListPhysicalGpuProfilesRequest, args ...map[string]interface{}) (*import1.ListPhysicalGpuProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/physical-gpu-profiles"

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

	unmarshalledResp := new(import1.ListPhysicalGpuProfilesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the rackable units of the cluster identified by {clusterExtId}.
func (api *ClustersApi) ListRackableUnitsByClusterId(clusterExtId *string, args ...map[string]interface{}) (*import1.ListRackableUnitsByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRackableUnitsByClusterId(context.Background(), &import7.ListRackableUnitsByClusterIdRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// Lists the rackable units of the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ListRackableUnitsByClusterId(ctx context.Context, request *import7.ListRackableUnitsByClusterIdRequest, args ...map[string]interface{}) (*import1.ListRackableUnitsByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rackable-units"

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

	unmarshalledResp := new(import1.ListRackableUnitsByClusterIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists the RSYSLOG server configurations associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) ListRsyslogServersByClusterId(clusterExtId *string, args ...map[string]interface{}) (*import1.ListRsyslogServersByClusterIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRsyslogServersByClusterId(context.Background(), &import7.ListRsyslogServersByClusterIdRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// Lists the RSYSLOG server configurations associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ListRsyslogServersByClusterId(ctx context.Context, request *import7.ListRsyslogServersByClusterIdRequest, args ...map[string]interface{}) (*import1.ListRsyslogServersByClusterIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rsyslog-servers"

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

	unmarshalledResp := new(import1.ListRsyslogServersByClusterIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List Virtual GPU profiles.
func (api *ClustersApi) ListVirtualGpuProfiles(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.ListVirtualGpuProfilesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVirtualGpuProfiles(context.Background(), &import7.ListVirtualGpuProfilesRequest{
		ClusterExtId: clusterExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
	}, args...)
}

// List Virtual GPU profiles.
func (api *ClustersServiceApi) ListVirtualGpuProfiles(ctx context.Context, request *import7.ListVirtualGpuProfilesRequest, args ...map[string]interface{}) (*import1.ListVirtualGpuProfilesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/virtual-gpu-profiles"

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

	unmarshalledResp := new(import1.ListVirtualGpuProfilesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all virtual NICs for the host identified by {extId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersApi) ListVirtualNicsByHostId(clusterExtId *string, hostExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListVirtualNicsByHostIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListVirtualNicsByHostId(context.Background(), &import7.ListVirtualNicsByHostIdRequest{
		ClusterExtId: clusterExtId,
		HostExtId:    hostExtId,
		Page_:        page_,
		Limit_:       limit_,
		Filter_:      filter_,
		Orderby_:     orderby_,
		Select_:      select_,
	}, args...)
}

// Lists all virtual NICs for the host identified by {extId} belonging to the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) ListVirtualNicsByHostId(ctx context.Context, request *import7.ListVirtualNicsByHostIdRequest, args ...map[string]interface{}) (*import1.ListVirtualNicsByHostIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{hostExtId}/virtual-nics"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == request.HostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*request.HostExtId, "")), -1)
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

	unmarshalledResp := new(import1.ListVirtualNicsByHostIdApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes nodes from cluster identified by {extId}.
func (api *ClustersApi) RemoveNode(clusterExtId *string, body *import1.NodeRemovalParams, args ...map[string]interface{}) (*import1.RemoveNodeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RemoveNode(context.Background(), &import7.RemoveNodeRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Removes nodes from cluster identified by {extId}.
func (api *ClustersServiceApi) RemoveNode(ctx context.Context, request *import7.RemoveNodeRequest, args ...map[string]interface{}) (*import1.RemoveNodeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/remove-node"

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

	unmarshalledResp := new(import1.RemoveNodeApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes transport ports and protocol detail from the SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) RemoveSnmpTransport(clusterExtId *string, body *import1.SnmpTransport, args ...map[string]interface{}) (*import1.RemoveSnmpTransportsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.RemoveSnmpTransport(context.Background(), &import7.RemoveSnmpTransportRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Removes transport ports and protocol detail from the SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) RemoveSnmpTransport(ctx context.Context, request *import7.RemoveSnmpTransportRequest, args ...map[string]interface{}) (*import1.RemoveSnmpTransportsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/$actions/remove-transports"

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

	unmarshalledResp := new(import1.RemoveSnmpTransportsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update cluster operation.
func (api *ClustersApi) UpdateClusterById(extId *string, body *import1.Cluster, args ...map[string]interface{}) (*import1.UpdateClusterApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateClusterById(context.Background(), &import7.UpdateClusterByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Update cluster operation.
func (api *ClustersServiceApi) UpdateClusterById(ctx context.Context, request *import7.UpdateClusterByIdRequest, args ...map[string]interface{}) (*import1.UpdateClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{extId}"

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

	unmarshalledResp := new(import1.UpdateClusterApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update RSYSLOG server configuration except RSYSLOG server name as it is a primary key of the entity.
func (api *ClustersApi) UpdateRsyslogServerById(clusterExtId *string, extId *string, body *import1.RsyslogServer, args ...map[string]interface{}) (*import1.UpdateRsyslogServerApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateRsyslogServerById(context.Background(), &import7.UpdateRsyslogServerByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Update RSYSLOG server configuration except RSYSLOG server name as it is a primary key of the entity.
func (api *ClustersServiceApi) UpdateRsyslogServerById(ctx context.Context, request *import7.UpdateRsyslogServerByIdRequest, args ...map[string]interface{}) (*import1.UpdateRsyslogServerApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/rsyslog-servers/{extId}"

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

	unmarshalledResp := new(import1.UpdateRsyslogServerApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the status of SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) UpdateSnmpStatus(clusterExtId *string, body *import1.SnmpStatusParam, args ...map[string]interface{}) (*import1.UpdateSnmpStatusApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSnmpStatus(context.Background(), &import7.UpdateSnmpStatusRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Updates the status of SNMP configuration associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) UpdateSnmpStatus(ctx context.Context, request *import7.UpdateSnmpStatusRequest, args ...map[string]interface{}) (*import1.UpdateSnmpStatusApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/$actions/update-status"

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

	unmarshalledResp := new(import1.UpdateSnmpStatusApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update SNMP trap configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) UpdateSnmpTrapById(clusterExtId *string, extId *string, body *import1.SnmpTrap, args ...map[string]interface{}) (*import1.UpdateSnmpTrapApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSnmpTrapById(context.Background(), &import7.UpdateSnmpTrapByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Update SNMP trap configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) UpdateSnmpTrapById(ctx context.Context, request *import7.UpdateSnmpTrapByIdRequest, args ...map[string]interface{}) (*import1.UpdateSnmpTrapApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/traps/{extId}"

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

	unmarshalledResp := new(import1.UpdateSnmpTrapApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates SNMP user configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersApi) UpdateSnmpUserById(clusterExtId *string, extId *string, body *import1.SnmpUser, args ...map[string]interface{}) (*import1.UpdateSnmpUserApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSnmpUserById(context.Background(), &import7.UpdateSnmpUserByIdRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Updates SNMP user configuration identified by {extId} associated with the cluster identified by {clusterExtId}.
func (api *ClustersServiceApi) UpdateSnmpUserById(ctx context.Context, request *import7.UpdateSnmpUserByIdRequest, args ...map[string]interface{}) (*import1.UpdateSnmpUserApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/snmp/users/{extId}"

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

	unmarshalledResp := new(import1.UpdateSnmpUserApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Validates hypervisor bundle and node uplinks of the node. This API is not supported for XEN hypervisor type.
func (api *ClustersApi) ValidateNode(clusterExtId *string, body *import1.ValidateNodeParam, args ...map[string]interface{}) (*import1.ValidateNodeApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewClustersServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ValidateNode(context.Background(), &import7.ValidateNodeRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// Validates hypervisor bundle and node uplinks of the node. This API is not supported for XEN hypervisor type.
func (api *ClustersServiceApi) ValidateNode(ctx context.Context, request *import7.ValidateNodeRequest, args ...map[string]interface{}) (*import1.ValidateNodeApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/$actions/validate-node"

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

	unmarshalledResp := new(import1.ValidateNodeApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
