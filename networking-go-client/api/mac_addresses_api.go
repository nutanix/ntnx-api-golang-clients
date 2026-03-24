package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
	import18 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/macaddresses"
	"net/http"
	"net/url"
	"strings"
)

type MacAddressesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *MacAddressesServiceApi
}

type MacAddressesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewMacAddressesApi(apiClient *client.ApiClient) *MacAddressesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &MacAddressesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewMacAddressesServiceApi(a.ApiClient)

	return a
}

func NewMacAddressesServiceApi(apiClient *client.ApiClient) *MacAddressesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &MacAddressesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get a specified MAC address that the specified Layer2Stretch has learned (i.e. ARP broadcasts, manual configuration).
func (api *MacAddressesApi) GetLearnedMacAddressForLayer2StretchById(layer2StretchExtId *string, extId *string, args ...map[string]interface{}) (*import4.GetLearnedMacAddressApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMacAddressesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetLearnedMacAddressForLayer2StretchById(context.Background(), &import18.GetLearnedMacAddressForLayer2StretchByIdRequest{
		Layer2StretchExtId: layer2StretchExtId,
		ExtId:              extId,
	}, args...)
}

// Get a specified MAC address that the specified Layer2Stretch has learned (i.e. ARP broadcasts, manual configuration).
func (api *MacAddressesServiceApi) GetLearnedMacAddressForLayer2StretchById(ctx context.Context, request *import18.GetLearnedMacAddressForLayer2StretchByIdRequest, args ...map[string]interface{}) (*import4.GetLearnedMacAddressApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/layer2-stretches/{layer2StretchExtId}/learned-mac-addresses/{extId}"

	// verify the required parameter 'layer2StretchExtId' is set
	if nil == request.Layer2StretchExtId {
		return nil, client.ReportError("layer2StretchExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"layer2StretchExtId"+"}", url.PathEscape(client.ParameterToString(*request.Layer2StretchExtId, "")), -1)
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

	unmarshalledResp := new(import4.GetLearnedMacAddressApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the MAC addresses that the specified Layer2Stretch has learned (i.e. ARP broadcasts, manual configuration).
func (api *MacAddressesApi) ListLearnedMacAddressesByLayer2StretchId(layer2StretchExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, args ...map[string]interface{}) (*import4.ListLearnedMacAddressesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewMacAddressesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListLearnedMacAddressesByLayer2StretchId(context.Background(), &import18.ListLearnedMacAddressesByLayer2StretchIdRequest{
		Layer2StretchExtId: layer2StretchExtId,
		Page_:              page_,
		Limit_:             limit_,
		Filter_:            filter_,
		Orderby_:           orderby_,
	}, args...)
}

// Get the MAC addresses that the specified Layer2Stretch has learned (i.e. ARP broadcasts, manual configuration).
func (api *MacAddressesServiceApi) ListLearnedMacAddressesByLayer2StretchId(ctx context.Context, request *import18.ListLearnedMacAddressesByLayer2StretchIdRequest, args ...map[string]interface{}) (*import4.ListLearnedMacAddressesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/config/layer2-stretches/{layer2StretchExtId}/learned-mac-addresses"

	// verify the required parameter 'layer2StretchExtId' is set
	if nil == request.Layer2StretchExtId {
		return nil, client.ReportError("layer2StretchExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"layer2StretchExtId"+"}", url.PathEscape(client.ParameterToString(*request.Layer2StretchExtId, "")), -1)
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

	unmarshalledResp := new(import4.ListLearnedMacAddressesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
