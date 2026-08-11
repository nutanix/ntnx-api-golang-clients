package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/request/buckets"
	"net/http"
	"net/url"
	"strings"
)

type BucketsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *BucketsServiceApi
}

type BucketsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewBucketsApi(apiClient *client.ApiClient) *BucketsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BucketsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewBucketsServiceApi(a.ApiClient)

	return a
}

func NewBucketsServiceApi(apiClient *client.ApiClient) *BucketsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &BucketsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// \"Get details of a specific bucket in an Object store. If 'x-ntnx-objects-namespace header' is provided with a valid federated namespace value, the bucket in the federated namespace is returned. Otherwise, the bucket in the local namespace is returned.\"
func (api *BucketsApi) GetBucketByName(objectStoreExtId *string, bucketName *string, xNtnxObjectsNamespace *string, args ...map[string]interface{}) (*import1.GetBucketApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBucketsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetBucketByName(context.Background(), &import2.GetBucketByNameRequest{
		ObjectStoreExtId:      objectStoreExtId,
		BucketName:            bucketName,
		XNtnxObjectsNamespace: xNtnxObjectsNamespace,
	}, args...)
}

// \"Get details of a specific bucket in an Object store. If 'x-ntnx-objects-namespace header' is provided with a valid federated namespace value, the bucket in the federated namespace is returned. Otherwise, the bucket in the local namespace is returned.\"
func (api *BucketsServiceApi) GetBucketByName(ctx context.Context, request *import2.GetBucketByNameRequest, args ...map[string]interface{}) (*import1.GetBucketApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/buckets/{bucketName}"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'bucketName' is set
	if nil == request.BucketName {
		return nil, client.ReportError("bucketName is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"bucketName"+"}", url.PathEscape(client.ParameterToString(*request.BucketName, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	if request.XNtnxObjectsNamespace != nil {
		headerParams["x-ntnx-objects-namespace"] = client.ParameterToString(*request.XNtnxObjectsNamespace, "")
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
	unmarshalledResp := new(import1.GetBucketApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// \"Get a list of the buckets in an Object store. If 'x-ntnx-objects-namespace header' is provided with a valid federated namespace value, the buckets in the federated namespace are returned. Otherwise, the buckets in the local namespace are returned.\"
func (api *BucketsApi) ListBucketsByObjectstoreId(objectStoreExtId *string, xNtnxObjectsNamespace *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListBucketsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewBucketsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListBucketsByObjectstoreId(context.Background(), &import2.ListBucketsByObjectstoreIdRequest{
		ObjectStoreExtId:      objectStoreExtId,
		XNtnxObjectsNamespace: xNtnxObjectsNamespace,
		Page_:                 page_,
		Limit_:                limit_,
		Filter_:               filter_,
		Orderby_:              orderby_,
		Select_:               select_,
	}, args...)
}

// \"Get a list of the buckets in an Object store. If 'x-ntnx-objects-namespace header' is provided with a valid federated namespace value, the buckets in the federated namespace are returned. Otherwise, the buckets in the local namespace are returned.\"
func (api *BucketsServiceApi) ListBucketsByObjectstoreId(ctx context.Context, request *import2.ListBucketsByObjectstoreIdRequest, args ...map[string]interface{}) (*import1.ListBucketsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/buckets"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
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
	if request.XNtnxObjectsNamespace != nil {
		headerParams["x-ntnx-objects-namespace"] = client.ParameterToString(*request.XNtnxObjectsNamespace, "")
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
	unmarshalledResp := new(import1.ListBucketsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
