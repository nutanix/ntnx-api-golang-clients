//Api classes for clustermgmt's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type ClusterApi struct {
	ApiClient *client.ApiClient
}

func NewClusterApi(apiClient *client.ApiClient) *ClusterApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ClusterApi{
		ApiClient: apiClient,
	}
	return a
}

/**
  Add node on a cluster
  Add node on a cluster

  parameters:-
  -> body (clustermgmt.v4.config.ExpandClusterParams) (required) : Property of the node to be added
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.AddNodeTaskResponse, error)
*/
func (api *ClusterApi) AddNode(body *import1.ExpandClusterParams, clusterExtId *string, args ...map[string]interface{}) (*import1.AddNodeTaskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/expand-cluster"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AddNodeTaskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Add new RSYSLOG server configuration
  Add new RSYSLOG server configuration

  parameters:-
  -> body (clustermgmt.v4.config.RsyslogServer) (required) : RSYSLOG server to add
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.AddRsyslogServerTaskApiResponse, error)
*/
func (api *ClusterApi) AddRsyslogServer(body *import1.RsyslogServer, clusterExtId *string, args ...map[string]interface{}) (*import1.AddRsyslogServerTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AddRsyslogServerTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Add SNMP transports
  Add SNMP transports

  parameters:-
  -> body ([]clustermgmt.v4.config.SnmpTransport) (required) : SNMP transports to add
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.AddSnmpTransportsTaskApiResponse, error)
*/
func (api *ClusterApi) AddSnmpTransport(body *[]import1.SnmpTransport, clusterExtId *string, args ...map[string]interface{}) (*import1.AddSnmpTransportsTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/add-transports"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AddSnmpTransportsTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Add SNMP trap
  Add SNMP trap

  parameters:-
  -> body (clustermgmt.v4.config.SnmpTrap) (required) : SNMP trap to add
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.AddSnmpTrapTaskApiResponse, error)
*/
func (api *ClusterApi) AddSnmpTrap(body *import1.SnmpTrap, clusterExtId *string, args ...map[string]interface{}) (*import1.AddSnmpTrapTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AddSnmpTrapTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Add SNMP user
  Add SNMP user

  parameters:-
  -> body (clustermgmt.v4.config.SnmpUser) (required) : SNMP user to add
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.AddSnmpUserTaskApiResponse, error)
*/
func (api *ClusterApi) AddSnmpUser(body *import1.SnmpUser, clusterExtId *string, args ...map[string]interface{}) (*import1.AddSnmpUserTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.AddSnmpUserTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete RSYSLOG Server
  Delete RSYSLOG Server

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> rsyslogServerExtId (string) (required) : RSYSLOG server UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.DeleteRsyslogServerTaskApiResponse, error)
*/
func (api *ClusterApi) DeleteRsyslogServer(clusterExtId *string, rsyslogServerExtId *string, args ...map[string]interface{}) (*import1.DeleteRsyslogServerTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'rsyslogServerExtId' is set
	if nil == rsyslogServerExtId {
		return nil, client.ReportError("rsyslogServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"rsyslogServerExtId"+"}", url.PathEscape(client.ParameterToString(*rsyslogServerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DeleteRsyslogServerTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete SNMP trap
  Delete SNMP trap

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> trapExtId (string) (required) : SNMP trap UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.DeleteSnmpTrapTaskApiResponse, error)
*/
func (api *ClusterApi) DeleteSnmpTrap(clusterExtId *string, trapExtId *string, args ...map[string]interface{}) (*import1.DeleteSnmpTrapTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'trapExtId' is set
	if nil == trapExtId {
		return nil, client.ReportError("trapExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"trapExtId"+"}", url.PathEscape(client.ParameterToString(*trapExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DeleteSnmpTrapTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Delete SNMP user
  Delete SNMP user

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> userExtId (string) (required) : SNMP user UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.DeleteSnmpUserTaskApiResponse, error)
*/
func (api *ClusterApi) DeleteSnmpUser(clusterExtId *string, userExtId *string, args ...map[string]interface{}) (*import1.DeleteSnmpUserTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'userExtId' is set
	if nil == userExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*userExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DeleteSnmpUserTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Discover unconfigured nodes
  Get the unconfigured node details such as node UUID, node position, node IP, foundation version and more.

  parameters:-
  -> body (clustermgmt.v4.config.NodeDiscoveryParams) (required) : Discover unconfigured node details
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.DiscoverNodeTaskApiResponse, error)
*/
func (api *ClusterApi) DiscoverUnconfiguredNodes(body *import1.NodeDiscoveryParams, clusterExtId *string, args ...map[string]interface{}) (*import1.DiscoverNodeTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/discover-unconfigured-nodes"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DiscoverNodeTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get all host GPUs
  Get all host GPUs

  parameters:-
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - cluster/uuid - numberOfVgpusAllocated
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - numberOfVgpusAllocated
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetAllHostGpusResponse, error)
*/
func (api *ClusterApi) GetAllHostGpus(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetAllHostGpusResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/host-gpus"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAllHostGpusResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get all host entities
  Get all host entities

  parameters:-
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - bootTimeUsecs - cluster/uuid - cpuCapacityHz - cpuFrequencyHz - cpuModel - defaultVhdContainerId - defaultVhdContainerUuid - defaultVhdLocation - defaultVmContainerId - defaultVmContainerUuid - defaultVmLocation - gpuDriverVersion - hostName - memorySizeBytes - numberOfCpuCores - numberOfCpuSockets - numberOfCpuThreads
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - hostName
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetAllHostsResponse, error)
*/
func (api *ClusterApi) GetAllHosts(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetAllHostsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/hosts"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAllHostsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get cluster entity
  Get cluster entity

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetClusterResponse, error)
*/
func (api *ClusterApi) GetCluster(clusterExtId *string, args ...map[string]interface{}) (*import1.GetClusterResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetClusterResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get cluster entities
  Get cluster entities

  parameters:-
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - extId - name
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - name
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetClustersResponse, error)
*/
func (api *ClusterApi) GetClusters(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetClustersResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetClustersResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get domain fault tolerance status
  Get domain fault tolerance status

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetDomainFaultToleranceResponse, error)
*/
func (api *ClusterApi) GetDomainFaultToleranceStatus(clusterExtId *string, args ...map[string]interface{}) (*import1.GetDomainFaultToleranceResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/domain-fault-tolerance-status"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetDomainFaultToleranceResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get GPU profiles
  Get GPU profiles

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetGpuProfilesResponse, error)
*/
func (api *ClusterApi) GetGpuProfiles(clusterExtId *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.GetGpuProfilesResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/gpu-profiles"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetGpuProfilesResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get host entity
  Get host entity

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> hostExtId (string) (required) : Host UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetHostResponse, error)
*/
func (api *ClusterApi) GetHost(clusterExtId *string, hostExtId *string, args ...map[string]interface{}) (*import1.GetHostResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == hostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*hostExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetHostResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get host GPU
  Get host GPU

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> hostExtId (string) (required) : Host UUID
  -> hostGpuExtId (string) (required) : Host GPU UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetHostGpuResponse, error)
*/
func (api *ClusterApi) GetHostGpu(clusterExtId *string, hostExtId *string, hostGpuExtId *string, args ...map[string]interface{}) (*import1.GetHostGpuResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-gpus/{hostGpuExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == hostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}
	// verify the required parameter 'hostGpuExtId' is set
	if nil == hostGpuExtId {
		return nil, client.ReportError("hostGpuExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*hostExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostGpuExtId"+"}", url.PathEscape(client.ParameterToString(*hostGpuExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetHostGpuResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get host GPUs on a particular host
  Get host GPUs on a particular host

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> hostExtId (string) (required) : Host UUID
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - cluster/uuid - numberOfVgpusAllocated
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - numberOfVgpusAllocated
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetHostGpusResponse, error)
*/
func (api *ClusterApi) GetHostGpus(clusterExtId *string, hostExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetHostGpusResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-gpus"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == hostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*hostExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetHostGpusResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get all host GPUs of a cluster
  Get all host GPUs of a cluster

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - cluster/uuid - numberOfVgpusAllocated
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - numberOfVgpusAllocated
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetClusterHostGpusResponse, error)
*/
func (api *ClusterApi) GetHostGpusOfCluster(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetClusterHostGpusResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/host-gpus"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetClusterHostGpusResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get host entities of a cluster
  Get host entities of a cluster

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource. Any number out of this range will lead to no results being returned.
  -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will lead to a validation error. If the limit is not provided a default value of 50 records will be returned in the result set.
  -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - bootTimeUsecs - cluster/uuid - cpuCapacityHz - cpuFrequencyHz - cpuModel - defaultVhdContainerId - defaultVhdContainerUuid - defaultVhdLocation - defaultVmContainerId - defaultVmContainerUuid - defaultVmLocation - gpuDriverVersion - hostName - memorySizeBytes - numberOfCpuCores - numberOfCpuSockets - numberOfCpuThreads
  -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - hostName
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetHostsResponse, error)
*/
func (api *ClusterApi) GetHosts(clusterExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import1.GetHostsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetHostsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get rackable unit entity
  Get rackable unit entity

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> rackableUnitExtId (string) (required) : Rackable unit UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetRackableUnitResponse, error)
*/
func (api *ClusterApi) GetRackableUnit(clusterExtId *string, rackableUnitExtId *string, args ...map[string]interface{}) (*import1.GetRackableUnitResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rackable-units/{rackableUnitExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'rackableUnitExtId' is set
	if nil == rackableUnitExtId {
		return nil, client.ReportError("rackableUnitExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"rackableUnitExtId"+"}", url.PathEscape(client.ParameterToString(*rackableUnitExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetRackableUnitResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get rackable unit entities of a cluster
  Get rackable unit entities of a cluster

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetRackableUnitsResponse, error)
*/
func (api *ClusterApi) GetRackableUnits(clusterExtId *string, args ...map[string]interface{}) (*import1.GetRackableUnitsResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rackable-units"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetRackableUnitsResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get RSYSLOG server configuration
  Get RSYSLOG server configuration

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> rsyslogServerExtId (string) (required) : RSYSLOG server UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetRsyslogServerResponse, error)
*/
func (api *ClusterApi) GetRsyslogServer(clusterExtId *string, rsyslogServerExtId *string, args ...map[string]interface{}) (*import1.GetRsyslogServerResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'rsyslogServerExtId' is set
	if nil == rsyslogServerExtId {
		return nil, client.ReportError("rsyslogServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"rsyslogServerExtId"+"}", url.PathEscape(client.ParameterToString(*rsyslogServerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetRsyslogServerResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get RSYSLOG servers configuration details
  Get RSYSLOG servers configuration details

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetRsyslogServersResponse, error)
*/
func (api *ClusterApi) GetRsyslogServers(clusterExtId *string, args ...map[string]interface{}) (*import1.GetRsyslogServersResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetRsyslogServersResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get SNMP information
  Get SNMP information

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetSnmpResponse, error)
*/
func (api *ClusterApi) GetSnmp(clusterExtId *string, args ...map[string]interface{}) (*import1.GetSnmpResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetSnmpResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get SNMP trap configuration
  Get SNMP trap configuration

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> trapExtId (string) (required) : SNMP trap UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetSnmpTrapResponse, error)
*/
func (api *ClusterApi) GetSnmpTrap(clusterExtId *string, trapExtId *string, args ...map[string]interface{}) (*import1.GetSnmpTrapResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'trapExtId' is set
	if nil == trapExtId {
		return nil, client.ReportError("trapExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"trapExtId"+"}", url.PathEscape(client.ParameterToString(*trapExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetSnmpTrapResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get SNMP user configuration
  Get SNMP user configuration

  parameters:-
  -> clusterExtId (string) (required) : Cluster UUID
  -> userExtId (string) (required) : SNMP user UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetSnmpUserResponse, error)
*/
func (api *ClusterApi) GetSnmpUser(clusterExtId *string, userExtId *string, args ...map[string]interface{}) (*import1.GetSnmpUserResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId}"

	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'userExtId' is set
	if nil == userExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*userExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetSnmpUserResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get hypervisor bundle compatibility information
  Provides details on whether hypervisor bundle is compatible or not

  parameters:-
  -> body (clustermgmt.v4.config.BundleParam) (required) : ISO attributes to validate compatibility
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.IsBundleCompatibleTaskResponse, error)
*/
func (api *ClusterApi) IsBundleCompatible(body *import1.BundleParam, clusterExtId *string, args ...map[string]interface{}) (*import1.IsBundleCompatibleTaskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/validate-bundle"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.IsBundleCompatibleTaskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get hypervisor ISO upload information
  Provides information on whether hypervisor ISO upload is required or not

  parameters:-
  -> body (clustermgmt.v4.config.HypervisorUploadParam) (required) : Parameters to get information on whether hypervisor ISO upload is required or not
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.HypervisorUplpadRequiredTaskResponse, error)
*/
func (api *ClusterApi) IsHypervisorUploadRequired(body *import1.HypervisorUploadParam, clusterExtId *string, args ...map[string]interface{}) (*import1.HypervisorUplpadRequiredTaskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/check-hypervisor-requirements"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.HypervisorUplpadRequiredTaskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get network information of unconfigured nodes
  Get a dictionary of cluster networks and available uplinks on the given nodes

  parameters:-
  -> body (clustermgmt.v4.config.NodeDetails) (required) : Node specific details required to fetch node networking information
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetNodeNetworkingTaskApiResponse, error)
*/
func (api *ClusterApi) NodesNetworkingDetails(body *import1.NodeDetails, clusterExtId *string, args ...map[string]interface{}) (*import1.GetNodeNetworkingTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/fetch-node-networking-details"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetNodeNetworkingTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Removes nodes from cluster
  Removes nodes from cluster

  parameters:-
  -> body (clustermgmt.v4.config.NodeRemovalParams) (required) : Parameters to remove nodes from cluster
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.RemoveNodeTaskResponse, error)
*/
func (api *ClusterApi) RemoveNode(body *import1.NodeRemovalParams, clusterExtId *string, args ...map[string]interface{}) (*import1.RemoveNodeTaskResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/remove-node"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.RemoveNodeTaskResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Remove SNMP transports
  Remove SNMP transports

  parameters:-
  -> body ([]clustermgmt.v4.config.SnmpTransport) (required) : SNMP transports to remove
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.RemoveSnmpTransportsTaskApiResponse, error)
*/
func (api *ClusterApi) RemoveSnmpTransport(body *[]import1.SnmpTransport, clusterExtId *string, args ...map[string]interface{}) (*import1.RemoveSnmpTransportsTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/remove-transports"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.RemoveSnmpTransportsTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Rename AHV host in a cluster.
  Rename AHV host in a cluster. Requests are processed in parallel and if multiple rename requests are submitted they can be executed in random order.

  parameters:-
  -> body (clustermgmt.v4.config.HostNameParam) (required) : Host rename parameters
  -> clusterExtId (string) (required) : Cluster UUID
  -> hostExtId (string) (required) : Host UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.HostRenameResponse, error)
*/
func (api *ClusterApi) RenameHost(body *import1.HostNameParam, clusterExtId *string, hostExtId *string, args ...map[string]interface{}) (*import1.HostRenameResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/$actions/rename-host"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'hostExtId' is set
	if nil == hostExtId {
		return nil, client.ReportError("hostExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"hostExtId"+"}", url.PathEscape(client.ParameterToString(*hostExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.HostRenameResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Get task response based on the type of request
  Get task response based on the type of request

  parameters:-
  -> body (clustermgmt.v4.config.SearchParams) (required) : Search parameters
  -> taskExtId (string) (required) : The external identifier of the task.
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.GetSearchResponse, error)
*/
func (api *ClusterApi) SearchTask(body *import1.SearchParams, taskExtId *string, args ...map[string]interface{}) (*import1.GetSearchResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/tasks/{taskExtId}/$actions/fetch-task-response"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'taskExtId' is set
	if nil == taskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*taskExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetSearchResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update cluster
  Update cluster

  parameters:-
  -> body (clustermgmt.v4.config.ClusterEntity) (required) : Cluster resource to update
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.UpdateClusterTaskApiResponse, error)
*/
func (api *ClusterApi) UpdateCluster(body *import1.ClusterEntity, clusterExtId *string, args ...map[string]interface{}) (*import1.UpdateClusterTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.UpdateClusterTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update RSYSLOG server configuration
  Update RSYSLOG server configuration

  parameters:-
  -> body (clustermgmt.v4.config.RsyslogServer) (required) : RSYSLOG server to update
  -> clusterExtId (string) (required) : Cluster UUID
  -> rsyslogServerExtId (string) (required) : RSYSLOG server UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.UpdateRsyslogServerTaskApiResponse, error)
*/
func (api *ClusterApi) UpdateRsyslogServer(body *import1.RsyslogServer, clusterExtId *string, rsyslogServerExtId *string, args ...map[string]interface{}) (*import1.UpdateRsyslogServerTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'rsyslogServerExtId' is set
	if nil == rsyslogServerExtId {
		return nil, client.ReportError("rsyslogServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"rsyslogServerExtId"+"}", url.PathEscape(client.ParameterToString(*rsyslogServerExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.UpdateRsyslogServerTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update SNMP status
  Update SNMP status

  parameters:-
  -> body (clustermgmt.v4.config.SnmpStatusParam) (required) : SNMP status
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.UpdateSnmpStatusTaskApiResponse, error)
*/
func (api *ClusterApi) UpdateSnmpStatus(body *import1.SnmpStatusParam, clusterExtId *string, args ...map[string]interface{}) (*import1.UpdateSnmpStatusTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/update-status"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.UpdateSnmpStatusTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update SNMP trap
  Update SNMP trap

  parameters:-
  -> body (clustermgmt.v4.config.SnmpTrap) (required) : SNMP trap to update
  -> clusterExtId (string) (required) : Cluster UUID
  -> trapExtId (string) (required) : SNMP trap UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.UpdateSnmpTrapTaskApiResponse, error)
*/
func (api *ClusterApi) UpdateSnmpTrap(body *import1.SnmpTrap, clusterExtId *string, trapExtId *string, args ...map[string]interface{}) (*import1.UpdateSnmpTrapTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'trapExtId' is set
	if nil == trapExtId {
		return nil, client.ReportError("trapExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"trapExtId"+"}", url.PathEscape(client.ParameterToString(*trapExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.UpdateSnmpTrapTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Update SNMP user
  Update SNMP user

  parameters:-
  -> body (clustermgmt.v4.config.SnmpUser) (required) : SNMP user to update
  -> clusterExtId (string) (required) : Cluster UUID
  -> userExtId (string) (required) : SNMP user UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.UpdateSnmpUserTaskApiResponse, error)
*/
func (api *ClusterApi) UpdateSnmpUser(body *import1.SnmpUser, clusterExtId *string, userExtId *string, args ...map[string]interface{}) (*import1.UpdateSnmpUserTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId}"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'userExtId' is set
	if nil == userExtId {
		return nil, client.ReportError("userExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"userExtId"+"}", url.PathEscape(client.ParameterToString(*userExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.UpdateSnmpUserTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
  Validates uplink information for the nodes
  Validates whether the uplink information for the given nodes can be used to add nodes to the cluster

  parameters:-
  -> body ([]clustermgmt.v4.config.UplinkNode) (required) : Parameters for validating uplinks
  -> clusterExtId (string) (required) : Cluster UUID
  -> args (map[string]interface{}) (optional) : Additional Arguments

  returns: (*clustermgmt.v4.config.ValidateUplinksTaskApiResponse, error)
*/
func (api *ClusterApi) ValidateUplinks(body *[]import1.UplinkNode, clusterExtId *string, args ...map[string]interface{}) (*import1.ValidateUplinksTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/validate-uplinks"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
	// verify the required parameter 'clusterExtId' is set
	if nil == clusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*clusterExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.ValidateUplinksTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
