package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import18 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/notifications"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type NotificationsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *NotificationsServiceApi
}

type NotificationsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewNotificationsApi(apiClient *client.ApiClient) *NotificationsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NotificationsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewNotificationsServiceApi(a.ApiClient)

	return a
}

func NewNotificationsServiceApi(apiClient *client.ApiClient) *NotificationsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &NotificationsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Compute upgrade notifications for a set of LCM entities and their target versions. Notifications describe the impact of the planned upgrade on cluster nodes and entities, including disruptive actions such as host reboots, VM migrations, and maintenance-mode entry. Submit a NotificationsSpec containing entity update specifications (entity UUID and target version pairs) and optional credentials for third-party vendor management. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, retrieve the resource identifier from the task's completion_details field and use GET /notifications/{extId} to fetch the computed notification results. The result resource is valid for 1 hour from creation. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *NotificationsApi) ComputeNotifications(body *import1.NotificationsSpec, xClusterId *string, args ...map[string]interface{}) (*import7.ComputeNotificationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNotificationsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ComputeNotifications(context.Background(), &import18.ComputeNotificationsRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Compute upgrade notifications for a set of LCM entities and their target versions. Notifications describe the impact of the planned upgrade on cluster nodes and entities, including disruptive actions such as host reboots, VM migrations, and maintenance-mode entry. Submit a NotificationsSpec containing entity update specifications (entity UUID and target version pairs) and optional credentials for third-party vendor management. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, retrieve the resource identifier from the task's completion_details field and use GET /notifications/{extId} to fetch the computed notification results. The result resource is valid for 1 hour from creation. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency.
func (api *NotificationsServiceApi) ComputeNotifications(ctx context.Context, request *import18.ComputeNotificationsRequest, args ...map[string]interface{}) (*import7.ComputeNotificationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/compute-notifications"

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

	if request.XClusterId != nil {
		headerParams["X-Cluster-Id"] = client.ParameterToString(*request.XClusterId, "")
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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import7.ComputeNotificationsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the computed upgrade notification result by its resource identifier. The resource identifier is obtained from the completion_details field of the task returned by POST /$actions/compute-notifications. The result contains notification items grouped by entity and location, each with severity levels and descriptive messages about the upgrade impact. The notification result is valid for 1 hour from the time it was computed. Review notifications before proceeding with prechecks (POST /$actions/prechecks) and upgrade (POST /$actions/upgrade) to understand the operational impact.
func (api *NotificationsApi) GetNotificationById(extId *string, args ...map[string]interface{}) (*import1.GetNotificationsByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewNotificationsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetNotificationById(context.Background(), &import18.GetNotificationByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the computed upgrade notification result by its resource identifier. The resource identifier is obtained from the completion_details field of the task returned by POST /$actions/compute-notifications. The result contains notification items grouped by entity and location, each with severity levels and descriptive messages about the upgrade impact. The notification result is valid for 1 hour from the time it was computed. Review notifications before proceeding with prechecks (POST /$actions/prechecks) and upgrade (POST /$actions/upgrade) to understand the operational impact.
func (api *NotificationsServiceApi) GetNotificationById(ctx context.Context, request *import18.GetNotificationByIdRequest, args ...map[string]interface{}) (*import1.GetNotificationsByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/notifications/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import1.GetNotificationsByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
