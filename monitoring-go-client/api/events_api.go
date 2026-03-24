package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/client"
	import6 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/request/events"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
	"net/http"
	"net/url"
	"strings"
)

type EventsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *EventsServiceApi
}

type EventsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewEventsApi(apiClient *client.ApiClient) *EventsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EventsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewEventsServiceApi(a.ApiClient)

	return a
}

func NewEventsServiceApi(apiClient *client.ApiClient) *EventsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &EventsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Fetches the details of an event identified by external identifier.
func (api *EventsApi) GetEventById(extId *string, args ...map[string]interface{}) (*import1.GetEventApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEventsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetEventById(context.Background(), &import6.GetEventByIdRequest{
		ExtId: extId,
	}, args...)
}

// Fetches the details of an event identified by external identifier.
func (api *EventsServiceApi) GetEventById(ctx context.Context, request *import6.GetEventByIdRequest, args ...map[string]interface{}) (*import1.GetEventApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/events/{extId}"

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

	unmarshalledResp := new(import1.GetEventApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches a list of events.
func (api *EventsApi) ListEvents(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListEventsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewEventsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListEvents(context.Background(), &import6.ListEventsRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Fetches a list of events.
func (api *EventsServiceApi) ListEvents(ctx context.Context, request *import6.ListEventsRequest, args ...map[string]interface{}) (*import1.ListEventsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/monitoring/v4.2/serviceability/events"

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

	unmarshalledResp := new(import1.ListEventsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
