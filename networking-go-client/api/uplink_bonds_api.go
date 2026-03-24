package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import32 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/uplinkbonds"
	"net/http"
	"net/url"
	"strings"
)

type UplinkBondsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *UplinkBondsServiceApi
}

type UplinkBondsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUplinkBondsApi(apiClient *client.ApiClient) *UplinkBondsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UplinkBondsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewUplinkBondsServiceApi(a.ApiClient)

	return a
}

func NewUplinkBondsServiceApi(apiClient *client.ApiClient) *UplinkBondsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UplinkBondsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get the uplink bond for the given extId.
func (api *UplinkBondsApi) GetUplinkBondById(extId *string, args ...map[string]interface{}) (*import4.GetUplinkBondApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUplinkBondsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetUplinkBondById(context.Background(), &import32.GetUplinkBondByIdRequest{
		ExtId: extId,
	}, args...)
}

// Get the uplink bond for the given extId.
func (api *UplinkBondsServiceApi) GetUplinkBondById(ctx context.Context, request *import32.GetUplinkBondByIdRequest, args ...map[string]interface{}) (*import4.GetUplinkBondApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/uplink-bonds/{extId}"

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

	unmarshalledResp := new(import4.GetUplinkBondApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List uplink bonds.
func (api *UplinkBondsApi) ListUplinkBonds(page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListUplinkBondsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewUplinkBondsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListUplinkBonds(context.Background(), &import32.ListUplinkBondsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
	}, args...)
}

// List uplink bonds.
func (api *UplinkBondsServiceApi) ListUplinkBonds(ctx context.Context, request *import32.ListUplinkBondsRequest, args ...map[string]interface{}) (*import4.ListUplinkBondsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/uplink-bonds"

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

	unmarshalledResp := new(import4.ListUplinkBondsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
