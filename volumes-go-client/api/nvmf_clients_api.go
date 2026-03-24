package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/request/nvmfclients"
	"net/http"
	"net/url"
	"strings"
)

type NvmfClientsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *NvmfClientsServiceApi
}

type NvmfClientsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewNvmfClientsApi(apiClient *client.ApiClient) *NvmfClientsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NvmfClientsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewNvmfClientsServiceApi(a.ApiClient)

	return a
}

func NewNvmfClientsServiceApi(apiClient *client.ApiClient) *NvmfClientsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NvmfClientsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches NVMe-TCP client details identified by its external identifier.
func (api *NvmfClientsApi) GetNvmfClientById(extId *string, args ...map[string]interface{}) (*import1.GetNvmfClientApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNvmfClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNvmfClientById(context.Background(), &import3.GetNvmfClientByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches NVMe-TCP client details identified by its external identifier.
func (api *NvmfClientsServiceApi) GetNvmfClientById(ctx context.Context, request *import3.GetNvmfClientByIdRequest, args ...map[string]interface{}) (*import1.GetNvmfClientApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/nvmf-clients/{extId}"

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

	unmarshalledResp := new(import1.GetNvmfClientApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of all the NVMe-TCP clients.
func (api *NvmfClientsApi) ListNvmfClients(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListNvmfClientsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNvmfClientsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListNvmfClients(context.Background(), &import3.ListNvmfClientsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Expand_:  expand_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of all the NVMe-TCP clients.
func (api *NvmfClientsServiceApi) ListNvmfClients(ctx context.Context, request *import3.ListNvmfClientsRequest, args ...map[string]interface{}) (*import1.ListNvmfClientsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/volumes/v4.2/config/nvmf-clients"

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

	unmarshalledResp := new(import1.ListNvmfClientsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
