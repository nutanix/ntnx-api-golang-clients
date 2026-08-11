package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/request/datamovementtargets"
	"net/http"
	"net/url"
	"strings"
)

type DataMovementTargetsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *DataMovementTargetsServiceApi
}

type DataMovementTargetsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDataMovementTargetsApi(apiClient *client.ApiClient) *DataMovementTargetsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DataMovementTargetsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewDataMovementTargetsServiceApi(a.ApiClient)

	return a
}

func NewDataMovementTargetsServiceApi(apiClient *client.ApiClient) *DataMovementTargetsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DataMovementTargetsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a new data movement target resource under the given Object store. Currently only remote Nutanix Objects data movement targets are supported. Example request body: ``` {   \"name\": \"data-movement-target-1\",   \"type\": \"NTNX_OBJECTS\",   \"details\": {     \"objectStoreExtId\": \"00062664-c65f-f854-65af-7cc25581f458\",     \"objectStoreFqdn\": \"objects-0.pc_nutanix.com\",     \"domainManagerExtId\": \"00062664-c65f-f854-65af-7cc25581f459\",     \"objectStorePublicNetworkIps\": [\"10.10.10.10\", \"10.10.10.11\"],     \"$objectType\": \"objects.v4.config.NTNXObjectsDataMovementTarget\"   } } ```
func (api *DataMovementTargetsApi) CreateDataMovementTarget(objectStoreExtId *string, body *import1.DataMovementTarget, args ...map[string]interface{}) (*import1.CreateDataMovementTargetApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDataMovementTargetsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateDataMovementTarget(context.Background(), &import3.CreateDataMovementTargetRequest{
		ObjectStoreExtId: objectStoreExtId,
		Body:             body,
	}, args...)
}

// Creates a new data movement target resource under the given Object store. Currently only remote Nutanix Objects data movement targets are supported. Example request body: ``` {   \"name\": \"data-movement-target-1\",   \"type\": \"NTNX_OBJECTS\",   \"details\": {     \"objectStoreExtId\": \"00062664-c65f-f854-65af-7cc25581f458\",     \"objectStoreFqdn\": \"objects-0.pc_nutanix.com\",     \"domainManagerExtId\": \"00062664-c65f-f854-65af-7cc25581f459\",     \"objectStorePublicNetworkIps\": [\"10.10.10.10\", \"10.10.10.11\"],     \"$objectType\": \"objects.v4.config.NTNXObjectsDataMovementTarget\"   } } ```
func (api *DataMovementTargetsServiceApi) CreateDataMovementTarget(ctx context.Context, request *import3.CreateDataMovementTargetRequest, args ...map[string]interface{}) (*import1.CreateDataMovementTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/data-movement-targets"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.CreateDataMovementTargetApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Returns the details of the specified data movement target of an Object store.
func (api *DataMovementTargetsApi) GetDataMovementTargetById(objectStoreExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetDataMovementTargetApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDataMovementTargetsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetDataMovementTargetById(context.Background(), &import3.GetDataMovementTargetByIdRequest{
		ObjectStoreExtId: objectStoreExtId,
		ExtId:            extId,
	}, args...)
}

// Returns the details of the specified data movement target of an Object store.
func (api *DataMovementTargetsServiceApi) GetDataMovementTargetById(ctx context.Context, request *import3.GetDataMovementTargetByIdRequest, args ...map[string]interface{}) (*import1.GetDataMovementTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/data-movement-targets/{extId}"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == request.ObjectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*request.ObjectStoreExtId, "")), -1)
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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.GetDataMovementTargetApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Get a list of the data movement targets that are configured in an Object store.
func (api *DataMovementTargetsApi) ListDataMovementTargets(objectStoreExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListDataMovementTargetsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewDataMovementTargetsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListDataMovementTargets(context.Background(), &import3.ListDataMovementTargetsRequest{
		ObjectStoreExtId: objectStoreExtId,
		Page_:            page_,
		Limit_:           limit_,
		Filter_:          filter_,
		Orderby_:         orderby_,
		Select_:          select_,
	}, args...)
}

// Get a list of the data movement targets that are configured in an Object store.
func (api *DataMovementTargetsServiceApi) ListDataMovementTargets(ctx context.Context, request *import3.ListDataMovementTargetsRequest, args ...map[string]interface{}) (*import1.ListDataMovementTargetsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.1/config/object-stores/{objectStoreExtId}/data-movement-targets"

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
	unmarshalledResp := new(import1.ListDataMovementTargetsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
