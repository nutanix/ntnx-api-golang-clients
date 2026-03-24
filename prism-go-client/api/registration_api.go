package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
	import9 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/registration"
	"net/http"
	"net/url"
	"strings"
)

type RegistrationApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RegistrationServiceApi
}

type RegistrationServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRegistrationApi(apiClient *client.ApiClient) *RegistrationApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RegistrationApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRegistrationServiceApi(a.ApiClient)

	return a
}

func NewRegistrationServiceApi(apiClient *client.ApiClient) *RegistrationServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RegistrationServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Returns the specifications for the specified registration.
func (api *RegistrationApi) GetRegistrationById(domainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import5.GetRegistrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegistrationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRegistrationById(context.Background(), &import9.GetRegistrationByIdRequest{
		DomainManagerExtId: domainManagerExtId,
		ExtId:              extId,
	}, args...)
}

// Returns the specifications for the specified registration.
func (api *RegistrationServiceApi) GetRegistrationById(ctx context.Context, request *import9.GetRegistrationByIdRequest, args ...map[string]interface{}) (*import5.GetRegistrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{domainManagerExtId}/registrations/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == request.DomainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import5.GetRegistrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// This endpoint returns the list of registered clusters to the domain manager. It could be the other domain managers, PEs, WitnessVMs, etc.
func (api *RegistrationApi) ListRegistrations(domainManagerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import5.ListRegistrationApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRegistrationServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListRegistrations(context.Background(), &import9.ListRegistrationsRequest{
		DomainManagerExtId: domainManagerExtId,
		Page_:              page_,
		Limit_:             limit_,
		Filter_:            filter_,
		Orderby_:           orderby_,
		Select_:            select_,
	}, args...)
}

// This endpoint returns the list of registered clusters to the domain manager. It could be the other domain managers, PEs, WitnessVMs, etc.
func (api *RegistrationServiceApi) ListRegistrations(ctx context.Context, request *import9.ListRegistrationsRequest, args ...map[string]interface{}) (*import5.ListRegistrationApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/management/domain-managers/{domainManagerExtId}/registrations"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == request.DomainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*request.DomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import5.ListRegistrationApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
