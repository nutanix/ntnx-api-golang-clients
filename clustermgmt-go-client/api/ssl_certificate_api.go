package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
	import13 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/request/sslcertificate"
	"net/http"
	"net/url"
	"strings"
)

type SSLCertificateApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *SSLCertificateServiceApi
}

type SSLCertificateServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewSSLCertificateApi(apiClient *client.ApiClient) *SSLCertificateApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SSLCertificateApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewSSLCertificateServiceApi(a.ApiClient)

	return a
}

func NewSSLCertificateServiceApi(apiClient *client.ApiClient) *SSLCertificateServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SSLCertificateServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Provides detailed information about the SSL certificate in privacy-enhanced mail (.pem) format for the specified cluster.
func (api *SSLCertificateApi) GetSSLCertificate(clusterExtId *string, args ...map[string]interface{}) (*import1.GetSSLCertificateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSSLCertificateServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSSLCertificate(context.Background(), &import13.GetSSLCertificateRequest{
		ClusterExtId: clusterExtId,
	}, args...)
}

// Provides detailed information about the SSL certificate in privacy-enhanced mail (.pem) format for the specified cluster.
func (api *SSLCertificateServiceApi) GetSSLCertificate(ctx context.Context, request *import13.GetSSLCertificateRequest, args ...map[string]interface{}) (*import1.GetSSLCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/ssl-certificate"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import1.GetSSLCertificateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// To update the SSL certificate for a specific cluster, you must provide a valid certificate payload in Base64 format. You can either import a new SSL certificate or replace an existing one by supplying all necessary fields, including the Base64-encoded certificate and private key. Alternatively, you can regenerate a self-signed certificate by specifying the privateKeyAlgorithm, noting that only the RSA_2048 algorithm is supported for SSL certificate regeneration. This process helps maintain the security and integrity of your cluster's communications by allowing you to update or regenerate the SSL certificate as needed.
func (api *SSLCertificateApi) UpdateSSLCertificate(clusterExtId *string, body *import1.SSLCertificate, args ...map[string]interface{}) (*import1.UpdateSSLCertificateApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewSSLCertificateServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSSLCertificate(context.Background(), &import13.UpdateSSLCertificateRequest{
		ClusterExtId: clusterExtId,
		Body:         body,
	}, args...)
}

// To update the SSL certificate for a specific cluster, you must provide a valid certificate payload in Base64 format. You can either import a new SSL certificate or replace an existing one by supplying all necessary fields, including the Base64-encoded certificate and private key. Alternatively, you can regenerate a self-signed certificate by specifying the privateKeyAlgorithm, noting that only the RSA_2048 algorithm is supported for SSL certificate regeneration. This process helps maintain the security and integrity of your cluster's communications by allowing you to update or regenerate the SSL certificate as needed.
func (api *SSLCertificateServiceApi) UpdateSSLCertificate(ctx context.Context, request *import13.UpdateSSLCertificateRequest, args ...map[string]interface{}) (*import1.UpdateSSLCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/clustermgmt/v4.2/config/clusters/{clusterExtId}/ssl-certificate"

	// verify the required parameter 'clusterExtId' is set
	if nil == request.ClusterExtId {
		return nil, client.ReportError("clusterExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"clusterExtId"+"}", url.PathEscape(client.ParameterToString(*request.ClusterExtId, "")), -1)
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

	unmarshalledResp := new(import1.UpdateSSLCertificateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
