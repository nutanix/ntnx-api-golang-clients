package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import15 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/ahv/config"
	import19 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/pciedevices"
	"net/http"
	"net/url"
	"strings"
)

type PcieDevicesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *PcieDevicesServiceApi
}

type PcieDevicesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewPcieDevicesApi(apiClient *client.ApiClient) *PcieDevicesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PcieDevicesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewPcieDevicesServiceApi(a.ApiClient)

	return a
}

func NewPcieDevicesServiceApi(apiClient *client.ApiClient) *PcieDevicesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &PcieDevicesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetch the PCIe devices
func (api *PcieDevicesApi) ListPcieDevices(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import15.ListPcieDevicesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewPcieDevicesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListPcieDevices(context.Background(), &import19.ListPcieDevicesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetch the PCIe devices
func (api *PcieDevicesServiceApi) ListPcieDevices(ctx context.Context, request *import19.ListPcieDevicesRequest, args ...map[string]interface{}) (*import15.ListPcieDevicesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.3/ahv/config/pcie-devices"

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
	unmarshalledResp := new(import15.ListPcieDevicesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
