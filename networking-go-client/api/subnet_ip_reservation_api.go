package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import27 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/subnetipreservation"
	"net/http"
	"net/url"
	"strings"
)

type SubnetIPReservationApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SubnetIPReservationServiceApi
}

type SubnetIPReservationServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSubnetIPReservationApi(apiClient *client.ApiClient) *SubnetIPReservationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SubnetIPReservationApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSubnetIPReservationServiceApi(a.ApiClient)

	return a
}

func NewSubnetIPReservationServiceApi(apiClient *client.ApiClient) *SubnetIPReservationServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SubnetIPReservationServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches a list of reserved IP addresses on a particular managed subnet.
func (api *SubnetIPReservationApi) ListReservedIpsBySubnetId(subnetExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import4.ListSubnetReservedIpsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSubnetIPReservationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListReservedIpsBySubnetId(context.Background(), &import27.ListReservedIpsBySubnetIdRequest{
		SubnetExtId: subnetExtId,
		Page_:       page_,
		Limit_:      limit_,
		Filter_:     filter_,
		Orderby_:    orderby_,
		Select_:     select_,
	}, args...)
}

// Fetches a list of reserved IP addresses on a particular managed subnet.
func (api *SubnetIPReservationServiceApi) ListReservedIpsBySubnetId(ctx context.Context, request *import27.ListReservedIpsBySubnetIdRequest, args ...map[string]interface{}) (*import4.ListSubnetReservedIpsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/subnets/{subnetExtId}/reserved-ips"

	// verify the required parameter 'subnetExtId' is set
	if nil == request.SubnetExtId {
		return nil, client.ReportError("subnetExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(*request.SubnetExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListSubnetReservedIpsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Reserve IP addresses on a subnet.
func (api *SubnetIPReservationApi) ReserveIpsBySubnetId(extId *string, body *import4.IpReserveSpec, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSubnetIPReservationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ReserveIpsBySubnetId(context.Background(), &import27.ReserveIpsBySubnetIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Reserve IP addresses on a subnet.
func (api *SubnetIPReservationServiceApi) ReserveIpsBySubnetId(ctx context.Context, request *import27.ReserveIpsBySubnetIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/subnets/{extId}/addresses/$actions/reserve"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Unreserve IP addresses on a subnet.
func (api *SubnetIPReservationApi) UnreserveIpsBySubnetId(extId *string, body *import4.IpUnreserveSpec, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSubnetIPReservationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UnreserveIpsBySubnetId(context.Background(), &import27.UnreserveIpsBySubnetIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Unreserve IP addresses on a subnet.
func (api *SubnetIPReservationServiceApi) UnreserveIpsBySubnetId(ctx context.Context, request *import27.UnreserveIpsBySubnetIdRequest, args ...map[string]interface{}) (*import4.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/subnets/{extId}/addresses/$actions/unreserve"

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

	unmarshalledResp := new(import4.TaskReferenceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
