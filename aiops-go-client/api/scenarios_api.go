package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/request/scenarios"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
	"net/http"
	"net/url"
	"strings"
)

type ScenariosApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ScenariosServiceApi
}

type ScenariosServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewScenariosApi(apiClient *client.ApiClient) *ScenariosApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ScenariosApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewScenariosServiceApi(a.ApiClient)

	return a
}

func NewScenariosServiceApi(apiClient *client.ApiClient) *ScenariosServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ScenariosServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a capacity planning scenario which can be used to analyse the future workload.
func (api *ScenariosApi) CreateScenario(body *import1.Scenario, args ...map[string]interface{}) (*import1.CreateScenarioApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateScenario(context.Background(), &import4.CreateScenarioRequest{
		Body: body,
	}, args...)
}

// Creates a capacity planning scenario which can be used to analyse the future workload.
func (api *ScenariosServiceApi) CreateScenario(ctx context.Context, request *import4.CreateScenarioRequest, args ...map[string]interface{}) (*import1.CreateScenarioApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

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

	unmarshalledResp := new(import1.CreateScenarioApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a simulation which can be used in the capacity scenario as a part of VM workload.
func (api *ScenariosApi) CreateSimulation(body *import1.Simulation, args ...map[string]interface{}) (*import1.CreateSimulationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateSimulation(context.Background(), &import4.CreateSimulationRequest{
		Body: body,
	}, args...)
}

// Creates a simulation which can be used in the capacity scenario as a part of VM workload.
func (api *ScenariosServiceApi) CreateSimulation(ctx context.Context, request *import4.CreateSimulationRequest, args ...map[string]interface{}) (*import1.CreateSimulationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/simulations"

	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

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

	unmarshalledResp := new(import1.CreateSimulationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a capacity planning scenario using the provided UUID.
func (api *ScenariosApi) DeleteScenarioById(extId *string, args ...map[string]interface{}) (*import1.DeleteScenarioApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteScenarioById(context.Background(), &import4.DeleteScenarioByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a capacity planning scenario using the provided UUID.
func (api *ScenariosServiceApi) DeleteScenarioById(ctx context.Context, request *import4.DeleteScenarioByIdRequest, args ...map[string]interface{}) (*import1.DeleteScenarioApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteScenarioApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a simulation identified by the provided UUID.
func (api *ScenariosApi) DeleteSimulationById(extId *string, args ...map[string]interface{}) (*import1.DeleteSimulationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteSimulationById(context.Background(), &import4.DeleteSimulationByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes a simulation identified by the provided UUID.
func (api *ScenariosServiceApi) DeleteSimulationById(ctx context.Context, request *import4.DeleteSimulationByIdRequest, args ...map[string]interface{}) (*import1.DeleteSimulationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/simulations/{extId}"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteSimulationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Generates recommendation for a planned capacity scenario using the provided UUID. You can send a polling request to find out the task status. The external identifier provided in the response can be used in GET scenarios API to retrieve the recommendation result upon the task completion.
func (api *ScenariosApi) GenerateRecommendation(extId *string, args ...map[string]interface{}) (*import1.GenerateRecommendationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GenerateRecommendation(context.Background(), &import4.GenerateRecommendationRequest{
		ExtId: extId,
	}, args...)
}

// Generates recommendation for a planned capacity scenario using the provided UUID. You can send a polling request to find out the task status. The external identifier provided in the response can be used in GET scenarios API to retrieve the recommendation result upon the task completion.
func (api *ScenariosServiceApi) GenerateRecommendation(ctx context.Context, request *import4.GenerateRecommendationRequest, args ...map[string]interface{}) (*import1.GenerateRecommendationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}/$actions/generate-recommendation"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GenerateRecommendationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Generates a report for a planned capacity scenario by the provided scenario UUID. You can send a polling request to know the task status. The external identifier from the response can be used in GET report API to fetch the generated report upon the task completion.
func (api *ScenariosApi) GenerateReport(extId *string, args ...map[string]interface{}) (*import1.GenerateReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GenerateReport(context.Background(), &import4.GenerateReportRequest{
		ExtId: extId,
	}, args...)
}

// Generates a report for a planned capacity scenario by the provided scenario UUID. You can send a polling request to know the task status. The external identifier from the response can be used in GET report API to fetch the generated report upon the task completion.
func (api *ScenariosServiceApi) GenerateReport(ctx context.Context, request *import4.GenerateReportRequest, args ...map[string]interface{}) (*import1.GenerateReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}/$actions/generate-report"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GenerateReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Generates runway for planned capacity scenario with the given UUID. You can send a polling request to know about the task status. The external identifier from the response can be used in the GET capacity scenarios API to fetch the runway result upon task completion.
func (api *ScenariosApi) GenerateRunway(extId *string, args ...map[string]interface{}) (*import1.GenerateRunwayApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GenerateRunway(context.Background(), &import4.GenerateRunwayRequest{
		ExtId: extId,
	}, args...)
}

// Generates runway for planned capacity scenario with the given UUID. You can send a polling request to know about the task status. The external identifier from the response can be used in the GET capacity scenarios API to fetch the runway result upon task completion.
func (api *ScenariosServiceApi) GenerateRunway(ctx context.Context, request *import4.GenerateRunwayRequest, args ...map[string]interface{}) (*import1.GenerateRunwayApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}/$actions/generate-runway"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GenerateRunwayApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a single capacity planing scenario using the UUID of the scenario.
func (api *ScenariosApi) GetScenarioById(extId *string, args ...map[string]interface{}) (*import1.GetScenarioApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetScenarioById(context.Background(), &import4.GetScenarioByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a single capacity planing scenario using the UUID of the scenario.
func (api *ScenariosServiceApi) GetScenarioById(ctx context.Context, request *import4.GetScenarioByIdRequest, args ...map[string]interface{}) (*import1.GetScenarioApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}"

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

	unmarshalledResp := new(import1.GetScenarioApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches the generated capacity planning scenario report by the report UUID.
func (api *ScenariosApi) GetScenarioReport(scenarioExtId *string, args ...map[string]interface{}) (*import1.GetScenarioReportApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetScenarioReport(context.Background(), &import4.GetScenarioReportRequest{
		ScenarioExtId: scenarioExtId,
	}, args...)
}

// Fetches the generated capacity planning scenario report by the report UUID.
func (api *ScenariosServiceApi) GetScenarioReport(ctx context.Context, request *import4.GetScenarioReportRequest, args ...map[string]interface{}) (*import1.GetScenarioReportApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{scenarioExtId}/reports"

	// verify the required parameter 'scenarioExtId' is set
	if nil == request.ScenarioExtId {
		return nil, client.ReportError("scenarioExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"scenarioExtId"+"}", url.PathEscape(client.ParameterToString(*request.ScenarioExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/octet-stream", "application/json"}

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

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import1.NewGetScenarioReportApiResponse()
			fileDetail := import1.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import2.Flag
			flags = append(flags, import2.Flag{Name: &flagName, Value: &flagValue})
			metadata := import3.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import1.GetScenarioReportApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a simulation from the provided UUID.
func (api *ScenariosApi) GetSimulationById(extId *string, args ...map[string]interface{}) (*import1.GetSimulationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetSimulationById(context.Background(), &import4.GetSimulationByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches a simulation from the provided UUID.
func (api *ScenariosServiceApi) GetSimulationById(ctx context.Context, request *import4.GetSimulationByIdRequest, args ...map[string]interface{}) (*import1.GetSimulationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/simulations/{extId}"

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

	unmarshalledResp := new(import1.GetSimulationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of capacity planning scenarios in a paginated manner.
func (api *ScenariosApi) ListScenarios(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListScenariosApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListScenarios(context.Background(), &import4.ListScenariosRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of capacity planning scenarios in a paginated manner.
func (api *ScenariosServiceApi) ListScenarios(ctx context.Context, request *import4.ListScenariosRequest, args ...map[string]interface{}) (*import1.ListScenariosApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios"

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

	unmarshalledResp := new(import1.ListScenariosApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of simulations in a paginated manner.
func (api *ScenariosApi) ListSimulations(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListSimulationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListSimulations(context.Background(), &import4.ListSimulationsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of simulations in a paginated manner.
func (api *ScenariosServiceApi) ListSimulations(ctx context.Context, request *import4.ListSimulationsRequest, args ...map[string]interface{}) (*import1.ListSimulationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/simulations"

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

	unmarshalledResp := new(import1.ListSimulationsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a capacity planning scenario using the provided UUID.
func (api *ScenariosApi) UpdateScenarioById(extId *string, body *import1.Scenario, args ...map[string]interface{}) (*import1.UpdateScenarioApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateScenarioById(context.Background(), &import4.UpdateScenarioByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a capacity planning scenario using the provided UUID.
func (api *ScenariosServiceApi) UpdateScenarioById(ctx context.Context, request *import4.UpdateScenarioByIdRequest, args ...map[string]interface{}) (*import1.UpdateScenarioApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/scenarios/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import1.UpdateScenarioApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates a simulation identified by the provided UUID.
func (api *ScenariosApi) UpdateSimulationById(extId *string, body *import1.Simulation, args ...map[string]interface{}) (*import1.UpdateSimulationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewScenariosServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateSimulationById(context.Background(), &import4.UpdateSimulationByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates a simulation identified by the provided UUID.
func (api *ScenariosServiceApi) UpdateSimulationById(ctx context.Context, request *import4.UpdateSimulationByIdRequest, args ...map[string]interface{}) (*import1.UpdateSimulationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/aiops/v4.2.b1/config/simulations/{extId}"

	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == request.Body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
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

	unmarshalledResp := new(import1.UpdateSimulationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
