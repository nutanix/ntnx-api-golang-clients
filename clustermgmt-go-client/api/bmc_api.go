package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/bmc"
	"net/http"
	"net/url"
	"strings"
)

type BmcApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *BmcServiceApi
}

type BmcServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewBmcApi(apiClient *client.ApiClient) *BmcApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BmcApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewBmcServiceApi(a.ApiClient)

	return a
}

func NewBmcServiceApi(apiClient *client.ApiClient) *BmcServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BmcServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Get BMC details of a host.
func (api *BmcApi) GetBmcInfo(clusterExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetBmcInfoResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBmcServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetBmcInfo(context.Background(), &import2.GetBmcInfoRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
	}, args...)
}

// Get BMC details of a host.
func (api *BmcServiceApi) GetBmcInfo(ctx context.Context, request *import2.GetBmcInfoRequest, args ...map[string]interface{}) (*import1.GetBmcInfoResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{extId}/bmc-info"

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

	unmarshalledResp := new(import1.GetBmcInfoResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the BMC summary information based on the provided identifier.
func (api *BmcApi) UpdateBmcInfo(clusterExtId *string, extId *string, body *import1.BmcInfo, args ...map[string]interface{}) (*import1.UpdateBmcInfoResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBmcServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateBmcInfo(context.Background(), &import2.UpdateBmcInfoRequest{
		ClusterExtId: clusterExtId,
		ExtId:        extId,
		Body:         body,
	}, args...)
}

// Updates the BMC summary information based on the provided identifier.
func (api *BmcServiceApi) UpdateBmcInfo(ctx context.Context, request *import2.UpdateBmcInfoRequest, args ...map[string]interface{}) (*import1.UpdateBmcInfoResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/hosts/{extId}/bmc-info"

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

	unmarshalledResp := new(import1.UpdateBmcInfoResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
