package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/request/dataprotectionclustercapabilities"
	"net/http"
	"net/url"
	"strings"
)

type DataProtectionClusterCapabilitiesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DataProtectionClusterCapabilitiesServiceApi
}

type DataProtectionClusterCapabilitiesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDataProtectionClusterCapabilitiesApi(apiClient *client.ApiClient) *DataProtectionClusterCapabilitiesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DataProtectionClusterCapabilitiesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDataProtectionClusterCapabilitiesServiceApi(a.ApiClient)

	return a
}

func NewDataProtectionClusterCapabilitiesServiceApi(apiClient *client.ApiClient) *DataProtectionClusterCapabilitiesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DataProtectionClusterCapabilitiesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// List the data protection capabilities for all the underlying clusters.
func (api *DataProtectionClusterCapabilitiesApi) ListDataProtectionClusterCapabilities(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDPClusterCapabilitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDataProtectionClusterCapabilitiesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDataProtectionClusterCapabilities(context.Background(), &import2.ListDataProtectionClusterCapabilitiesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List the data protection capabilities for all the underlying clusters.
func (api *DataProtectionClusterCapabilitiesServiceApi) ListDataProtectionClusterCapabilities(ctx context.Context, request *import2.ListDataProtectionClusterCapabilitiesRequest, args ...map[string]interface{}) (*import1.ListDPClusterCapabilitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/dataprotection/v4.3/config/cluster-capabilities"

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

	unmarshalledResp := new(import1.ListDPClusterCapabilitiesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
