package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/aws/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/request/awssubnets"
	"net/http"
	"net/url"
	"strings"
)

type AwsSubnetsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *AwsSubnetsServiceApi
}

type AwsSubnetsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewAwsSubnetsApi(apiClient *client.ApiClient) *AwsSubnetsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AwsSubnetsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewAwsSubnetsServiceApi(a.ApiClient)

	return a
}

func NewAwsSubnetsServiceApi(apiClient *client.ApiClient) *AwsSubnetsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &AwsSubnetsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get list of NC2 AWS subnets.
func (api *AwsSubnetsApi) ListAwsSubnets(xClusterId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListAwsSubnetsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewAwsSubnetsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListAwsSubnets(context.Background(), &import2.ListAwsSubnetsRequest{
		XClusterId: xClusterId,
		Page_:      page_,
		Limit_:     limit_,
		Filter_:    filter_,
		Orderby_:   orderby_,
		Select_:    select_,
	}, args...)
}

// Get list of NC2 AWS subnets.
func (api *AwsSubnetsServiceApi) ListAwsSubnets(ctx context.Context, request *import2.ListAwsSubnetsRequest, args ...map[string]interface{}) (*import1.ListAwsSubnetsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/networking/v4.3/aws/config/subnets"

	// verify the required parameter 'xClusterId' is set
	if nil == request.XClusterId {
		return nil, client.ReportError("xClusterId is required and must be specified")
	}

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
	headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
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

	unmarshalledResp := new(import1.ListAwsSubnetsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
