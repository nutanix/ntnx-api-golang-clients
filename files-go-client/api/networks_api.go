//Api classes for files's golang SDK
package api

import (
    "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	"strings"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type NetworksApi struct {
  ApiClient *client.ApiClient
}

func NewNetworksApi(apiClient *client.ApiClient) *NetworksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NetworksApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create a network
    Creates a new network using the provided request body.  The users need to specify only the `networkExtId` of the nvm network, for the managed network. For the unmanaged networks user can provide `defaultGateway`, `ipAddresses`, `virtualIpAddress`, `subnetMask`, `vlanId`.  A sample request body would look like this: ``` {   \"defaultGateway\": {     \"ipv4\": {       \"value\": \"10.53.80.1\"     }   },   \"ipAddresses\": {     \"ipv4\": {       \"value\": \"10.53.73.42\"     }   },   \"virtualIpAddress\": {     \"ipv4\": {       \"value\": \"\"     }   },   \"subnetMask\": {     \"ipv4\": {       \"value\": \"255.255.240.0\"     }   },   \"networkExtId\": \"8064cae7-13c6-4cc9-be63-021dec2448ac\",   \"vlanId\": 0,   \"extId\": \"bb64cae7-13c6-4cc9-be63-021dec2448bb\" } ``` 

    parameters:-
    -> body (files.v4.config.Network) (required) : A model that represents the file server network resources.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateNetworkApiResponse, error)
*/
func (api *NetworksApi) CreateNetwork(body *import1.Network, args ...map[string]interface{}) (*import1.CreateNetworkApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/networks"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CreateNetworkApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete a network
    Delete the file server network with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the network to be deleted. 

    parameters:-
    -> extId (string) (required) : The extId of the file server network.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteNetworkApiResponse, error)
*/
func (api *NetworksApi) DeleteNetwork(extId *string, args ...map[string]interface{}) (*import1.DeleteNetworkApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/networks/{extId}"

    // verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DeleteNetworkApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get network by extId
    Get the file server network with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the network to be fetched. 

    parameters:-
    -> extId (string) (required) : The extId of the file server network.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.NetworkGetByExtIdApiResponse, error)
*/
func (api *NetworksApi) GetNetworkByExtId(extId *string, args ...map[string]interface{}) (*import1.NetworkGetByExtIdApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/networks/{extId}"

    // verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkGetByExtIdApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get networks
    Get a paginated list of file server external networks.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported: - networkExtId - vlanId - virtualIpAddress - virtualNetworkName - subnetMask - defaultGateway - ipv6PrefixLength - isManaged  A sample call would look like this: ``` /api/files/v4.0.a2/config/file-server/networks?$filter=vlanId gt 10 ``` Example of supported query parameters for the file server network list API: ```  - ?$page=0&$limit=1  - ?$select=vlanId  - ?$filter=vlanId gt 10  - ?$orderby=virtualNetworkName desc  - ?$limit=5&$select=vlanId  - ?$limit=5&$select=virtualNetworkName &$orderby=virtualNetworkName desc ``` The `$orderby` query parameter allows specifying attributes on which to sort the returned list of networks  The following parameters support sorting in the file server network list:   - virtualNetworkName   - isManaged   - vlanId  A sample call would look like this: ``` /api/files/v4.0.a2/config/file-server/network?$orderby=virtualNetworkName desc ```  The `$select` query parameter allows specifying attributes which the user wants to fetch in the returned list of networks, other attributes will be returned as a null value.  The following attributes can be selected: - virtualNetworkName - isManaged - vlanId  Some more examples are given below: 1. Filter by virtualNetworkName: ``` /api/files/v4.0.a2/config/file-server/network?$filter=contains(virtualNetworkName, 'something')   OR /api/files/v4.0.a2/config/file-server/network?$filter=virtualNetworkName eq 'something' ```  2. Order by virtualNetworkName in ascending order ``` /api/files/v4.0.a2/config/file-server/network?$orderby=virtualNetworkName asc ```  3. Order by virtualNetworkName in descending order ``` /api/files/v4.0.a2/config/file-server/network?$orderby=virtualNetworkName desc ``` 4. Select by virtualNetworkName ``` /api/files/v4.0.a2/config/file-server/network?$select=virtualNetworkName ``` 5. Paginate the returned network list ``` /api/files/v4.0.a2/config/file-server/network?$page=0&$limit=1 ``` 7. Combination of queries ```   /api/files/v4.0.a2/config/file-server/network?$limit=5&$select=virtualNetworkName, vlanId &$orderby=virtualNetworkName desc ```  If the user doesn't specify any search query parameters, a list of all external networks will be returned. 

    parameters:-
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - isManaged - virtualNetworkName - vlanId 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - isManaged - virtualNetworkName - vlanId 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - defaultGateway/ipv4/prefixLength - defaultGateway/ipv4/value - defaultGateway/ipv6/prefixLength - defaultGateway/ipv6/value - extId - ipAddresses - ipv6PrefixLength - isManaged - links - networkExtId - subnetMask/ipv4/prefixLength - subnetMask/ipv4/value - subnetMask/ipv6/prefixLength - subnetMask/ipv6/value - tenantId - virtualIpAddress/ipv4/prefixLength - virtualIpAddress/ipv4/value - virtualIpAddress/ipv6/prefixLength - virtualIpAddress/ipv6/value - virtualNetworkName - vlanId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.NetworkListApiResponse, error)
*/
func (api *NetworksApi) GetNetworks(filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.NetworkListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/networks"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

