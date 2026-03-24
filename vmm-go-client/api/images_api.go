package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import10 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	import12 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/images"
	"net/http"
	"net/url"
	"strings"
)

type ImagesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
	ServiceClient *ImagesServiceApi
}

type ImagesServiceApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewImagesApi(apiClient *client.ApiClient) *ImagesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImagesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	a.ServiceClient = NewImagesServiceApi(a.ApiClient)

	return a
}

func NewImagesServiceApi(apiClient *client.ApiClient) *ImagesServiceApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ImagesServiceApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates an image using the provided request body. The name, type and source are mandatory fields to create an image.
func (api *ImagesApi) CreateImage(body *import9.Image, args ...map[string]interface{}) (*import9.CreateImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateImage(context.Background(), &import12.CreateImageRequest{
		Body: body,
	}, args...)
}

// Creates an image using the provided request body. The name, type and source are mandatory fields to create an image.
func (api *ImagesServiceApi) CreateImage(ctx context.Context, request *import12.CreateImageRequest, args ...map[string]interface{}) (*import9.CreateImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images"

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

	unmarshalledResp := new(import9.CreateImageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes the image with the given external identifier.
func (api *ImagesApi) DeleteImageById(extId *string, args ...map[string]interface{}) (*import9.DeleteImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteImageById(context.Background(), &import12.DeleteImageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the image with the given external identifier.
func (api *ImagesServiceApi) DeleteImageById(ctx context.Context, request *import12.DeleteImageByIdRequest, args ...map[string]interface{}) (*import9.DeleteImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images/{extId}"

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

	unmarshalledResp := new(import9.DeleteImageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Downloads the image with the given external identifier.
func (api *ImagesApi) GetFileByImageId(imageExtId *string, args ...map[string]interface{}) (*import9.GetImageFileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetFileByImageId(context.Background(), &import12.GetFileByImageIdRequest{
		ImageExtId: imageExtId,
	}, args...)
}

// Downloads the image with the given external identifier.
func (api *ImagesServiceApi) GetFileByImageId(ctx context.Context, request *import12.GetFileByImageIdRequest, args ...map[string]interface{}) (*import9.GetImageFileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images/{imageExtId}/file"

	// verify the required parameter 'imageExtId' is set
	if nil == request.ImageExtId {
		return nil, client.ReportError("imageExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"imageExtId"+"}", url.PathEscape(client.ParameterToString(*request.ImageExtId, "")), -1)
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

			response := import9.NewGetImageFileApiResponse()
			fileDetail := import9.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import10.Flag
			flags = append(flags, import10.Flag{Name: &flagName, Value: &flagValue})
			metadata := import11.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import9.GetImageFileApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the image details for the provided external identifier.
func (api *ImagesApi) GetImageById(extId *string, args ...map[string]interface{}) (*import9.GetImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetImageById(context.Background(), &import12.GetImageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the image details for the provided external identifier.
func (api *ImagesServiceApi) GetImageById(ctx context.Context, request *import12.GetImageByIdRequest, args ...map[string]interface{}) (*import9.GetImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images/{extId}"

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

	unmarshalledResp := new(import9.GetImageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Imports images owned by a registered Prism Element cluster into the current Prism Central.
func (api *ImagesApi) ImportImage(body *import9.ImageImportConfig, args ...map[string]interface{}) (*import9.ImportImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ImportImage(context.Background(), &import12.ImportImageRequest{
		Body: body,
	}, args...)
}

// Imports images owned by a registered Prism Element cluster into the current Prism Central.
func (api *ImagesServiceApi) ImportImage(ctx context.Context, request *import12.ImportImageRequest, args ...map[string]interface{}) (*import9.ImportImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images/$actions/import"

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

	unmarshalledResp := new(import9.ImportImageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists images owned by Prism Central, along with image details such as name, description, type, and so on. This API supports operation such as filtering, sorting, selection, and pagination.
func (api *ImagesApi) ListImages(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import9.ListImagesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListImages(context.Background(), &import12.ListImagesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists images owned by Prism Central, along with image details such as name, description, type, and so on. This API supports operation such as filtering, sorting, selection, and pagination.
func (api *ImagesServiceApi) ListImages(ctx context.Context, request *import12.ListImagesRequest, args ...map[string]interface{}) (*import9.ListImagesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images"

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

	unmarshalledResp := new(import9.ListImagesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the image with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure that the correct ETag is used.
func (api *ImagesApi) UpdateImageById(extId *string, body *import9.Image, args ...map[string]interface{}) (*import9.UpdateImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateImageById(context.Background(), &import12.UpdateImageByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the image with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure that the correct ETag is used.
func (api *ImagesServiceApi) UpdateImageById(ctx context.Context, request *import12.UpdateImageByIdRequest, args ...map[string]interface{}) (*import9.UpdateImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.2/content/images/{extId}"

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

	unmarshalledResp := new(import9.UpdateImageApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
