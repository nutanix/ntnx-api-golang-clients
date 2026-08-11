package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/client"
	import7 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/operations"
	import21 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/request/recommendations"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
	"net/http"
	"net/url"
	"strings"
)

type RecommendationsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *RecommendationsServiceApi
}

type RecommendationsServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewRecommendationsApi(apiClient *client.ApiClient) *RecommendationsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecommendationsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewRecommendationsServiceApi(a.ApiClient)

	return a
}

func NewRecommendationsServiceApi(apiClient *client.ApiClient) *RecommendationsServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &RecommendationsServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Compute upgrade recommendations for a set of LCM entities. Recommendations resolve inter-entity dependencies, identify additional entities that must be co-upgraded, and produce a validated list of entityUpdateSpecs ready for use in prechecks (POST /$actions/prechecks) and upgrades (POST /$actions/upgrade). Submit a RecommendationSpec containing one of four input types: entity types (SOFTWARE/FIRMWARE) for broad recommendations, target entities for entity-class/model-based selection, entity update specs for specific UUID-to-version mappings, or deploy specs for pre-deployment planning. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, retrieve the resource identifier from the task's completion_details field and use GET /recommendations/{extId} to fetch the computed recommendation results. The result resource is valid for 1 hour from creation. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. In dark-site mode, recommendations are limited to updates available from uploaded bundles.
func (api *RecommendationsApi) ComputeRecommendations(body *import1.RecommendationSpec, xClusterId *string, args ...map[string]interface{}) (*import7.ComputeRecommendationsApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecommendationsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ComputeRecommendations(context.Background(), &import21.ComputeRecommendationsRequest{
		Body:       body,
		XClusterId: xClusterId,
	}, args...)
}

// Compute upgrade recommendations for a set of LCM entities. Recommendations resolve inter-entity dependencies, identify additional entities that must be co-upgraded, and produce a validated list of entityUpdateSpecs ready for use in prechecks (POST /$actions/prechecks) and upgrades (POST /$actions/upgrade). Submit a RecommendationSpec containing one of four input types: entity types (SOFTWARE/FIRMWARE) for broad recommendations, target entities for entity-class/model-based selection, entity update specs for specific UUID-to-version mappings, or deploy specs for pre-deployment planning. When operating from Prism Central, supply the X-Cluster-Id header to target a specific Prism Element cluster; if omitted, the operation targets the Prism Central cluster itself. The operation is asynchronous and returns a TaskReference. Once the task completes successfully, retrieve the resource identifier from the task's completion_details field and use GET /recommendations/{extId} to fetch the computed recommendation results. The result resource is valid for 1 hour from creation. Requires the NTNX-Request-Id header (a UUID v4 value) for idempotency. In dark-site mode, recommendations are limited to updates available from uploaded bundles.
func (api *RecommendationsServiceApi) ComputeRecommendations(ctx context.Context, request *import21.ComputeRecommendationsRequest, args ...map[string]interface{}) (*import7.ComputeRecommendationsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/operations/$actions/compute-recommendations"

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
	unmarshalledResp := new(import7.ComputeRecommendationsApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieve the computed upgrade recommendation result by its resource identifier. The resource identifier is obtained from the completion_details field of the task returned by POST /$actions/compute-recommendations. The result includes the resolved entityUpdateSpecs (entity UUID and target version pairs), any entities that were skipped, modified, or added during dependency resolution, and deployable version information. The recommendation result is valid for 1 hour from the time it was computed. Use the entityUpdateSpecs from the result as input to POST /$actions/prechecks and POST /$actions/upgrade.
func (api *RecommendationsApi) GetRecommendationById(extId *string, args ...map[string]interface{}) (*import1.GetRecommendationByIdApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewRecommendationsServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetRecommendationById(context.Background(), &import21.GetRecommendationByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieve the computed upgrade recommendation result by its resource identifier. The resource identifier is obtained from the completion_details field of the task returned by POST /$actions/compute-recommendations. The result includes the resolved entityUpdateSpecs (entity UUID and target version pairs), any entities that were skipped, modified, or added during dependency resolution, and deployable version information. The recommendation result is valid for 1 hour from the time it was computed. Use the entityUpdateSpecs from the result as input to POST /$actions/prechecks and POST /$actions/upgrade.
func (api *RecommendationsServiceApi) GetRecommendationById(ctx context.Context, request *import21.GetRecommendationByIdRequest, args ...map[string]interface{}) (*import1.GetRecommendationByIdApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/lifecycle/v4.3/resources/recommendations/{extId}"

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
	unmarshalledResp := new(import1.GetRecommendationByIdApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
