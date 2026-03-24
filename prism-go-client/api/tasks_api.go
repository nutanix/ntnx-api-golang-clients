package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/config"
	import10 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/request/tasks"
	"net/http"
	"net/url"
	"strings"
)

type TasksApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *TasksServiceApi
}

type TasksServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewTasksApi(apiClient *client.ApiClient) *TasksApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TasksApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewTasksServiceApi(a.ApiClient)

	return a
}

func NewTasksServiceApi(apiClient *client.ApiClient) *TasksServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TasksServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Cancel an ongoing task for the provided taskExtId. This action is supported only if 'isCancelable' is set as True for the task. The task may continue to run for some time after the request has been made, until the backend server deems it is safe to cancel the task. Cancellation requests are idempotent and multiple such requests for the same ongoing task are supported.
func (api *TasksApi) CancelTask(taskExtId *string, args ...map[string]interface{}) (*import3.CancelTaskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CancelTask(context.Background(), &import10.CancelTaskRequest{
		TaskExtId: taskExtId,
	}, args...)
}

// Cancel an ongoing task for the provided taskExtId. This action is supported only if 'isCancelable' is set as True for the task. The task may continue to run for some time after the request has been made, until the backend server deems it is safe to cancel the task. Cancellation requests are idempotent and multiple such requests for the same ongoing task are supported.
func (api *TasksServiceApi) CancelTask(ctx context.Context, request *import10.CancelTaskRequest, args ...map[string]interface{}) (*import3.CancelTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks/{taskExtId}/$actions/cancel"

	// verify the required parameter 'taskExtId' is set
	if nil == request.TaskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*request.TaskExtId, "")), -1)
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

	unmarshalledResp := new(import3.CancelTaskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches an asynchronous operation called a task for the provided external identifier.
func (api *TasksApi) GetTaskById(extId *string, select_ *string, args ...map[string]interface{}) (*import3.GetTaskApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTaskById(context.Background(), &import10.GetTaskByIdRequest{
		ExtId:   extId,
		Select_: select_,
	}, args...)
}

// Fetches an asynchronous operation called a task for the provided external identifier.
func (api *TasksServiceApi) GetTaskById(ctx context.Context, request *import10.GetTaskByIdRequest, args ...map[string]interface{}) (*import3.GetTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks/{extId}"

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

	// Query Params
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

	unmarshalledResp := new(import3.GetTaskApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Fetches details of an individual job associated with a task. When a task is a batch task, it includes multiple actions, each represented as a separate job.
func (api *TasksApi) GetTaskJobById(taskExtId *string, extId *string, select_ *string, args ...map[string]interface{}) (*import3.GetTaskJobApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetTaskJobById(context.Background(), &import10.GetTaskJobByIdRequest{
		TaskExtId: taskExtId,
		ExtId:     extId,
		Select_:   select_,
	}, args...)
}

// Fetches details of an individual job associated with a task. When a task is a batch task, it includes multiple actions, each represented as a separate job.
func (api *TasksServiceApi) GetTaskJobById(ctx context.Context, request *import10.GetTaskJobByIdRequest, args ...map[string]interface{}) (*import3.GetTaskJobApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks/{taskExtId}/jobs/{extId}"

	// verify the required parameter 'taskExtId' is set
	if nil == request.TaskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == request.ExtId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*request.TaskExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*request.ExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
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

	unmarshalledResp := new(import3.GetTaskJobApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists entities associated with a task.
func (api *TasksApi) ListTaskEntities(taskExtId *string, page_ *int, limit_ *int, filter_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListTaskEntitiesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTaskEntities(context.Background(), &import10.ListTaskEntitiesRequest{
		TaskExtId: taskExtId,
		Page_:     page_,
		Limit_:    limit_,
		Filter_:   filter_,
		Select_:   select_,
	}, args...)
}

// Lists entities associated with a task.
func (api *TasksServiceApi) ListTaskEntities(ctx context.Context, request *import10.ListTaskEntitiesRequest, args ...map[string]interface{}) (*import3.ListTaskEntitiesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks/{taskExtId}/affected-entities"

	// verify the required parameter 'taskExtId' is set
	if nil == request.TaskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*request.TaskExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListTaskEntitiesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists jobs associated with a task. When a task is a batch task, it includes multiple actions, each represented as a separate job.
func (api *TasksApi) ListTaskJobs(taskExtId *string, page_ *int, limit_ *int, select_ *string, args ...map[string]interface{}) (*import3.ListTaskJobsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTaskJobs(context.Background(), &import10.ListTaskJobsRequest{
		TaskExtId: taskExtId,
		Page_:     page_,
		Limit_:    limit_,
		Select_:   select_,
	}, args...)
}

// Lists jobs associated with a task. When a task is a batch task, it includes multiple actions, each represented as a separate job.
func (api *TasksServiceApi) ListTaskJobs(ctx context.Context, request *import10.ListTaskJobsRequest, args ...map[string]interface{}) (*import3.ListTaskJobsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks/{taskExtId}/jobs"

	// verify the required parameter 'taskExtId' is set
	if nil == request.TaskExtId {
		return nil, client.ReportError("taskExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(*request.TaskExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListTaskJobsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// List of tasks in the system. The response can be further filtered/sorted using the filter and sort options provided. By default, the response would be sorted by 'createdTime' in descending order.
func (api *TasksApi) ListTasks(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListTasksApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewTasksServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListTasks(context.Background(), &import10.ListTasksRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// List of tasks in the system. The response can be further filtered/sorted using the filter and sort options provided. By default, the response would be sorted by 'createdTime' in descending order.
func (api *TasksServiceApi) ListTasks(ctx context.Context, request *import10.ListTasksRequest, args ...map[string]interface{}) (*import3.ListTasksApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.3/config/tasks"

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

	unmarshalledResp := new(import3.ListTasksApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
