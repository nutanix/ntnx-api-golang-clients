package api

import (
	"context"
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/client"
	import12 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import13 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
	import14 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/request/images"
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
func (api *ImagesApi) CreateImage(body *import11.Image, args ...map[string]interface{}) (*import11.CreateImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.CreateImage(context.Background(), &import14.CreateImageRequest{
		Body: body,
	}, args...)
}

// Creates an image using the provided request body. The name, type and source are mandatory fields to create an image.
func (api *ImagesServiceApi) CreateImage(ctx context.Context, request *import14.CreateImageRequest, args ...map[string]interface{}) (*import11.CreateImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.CreateImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Deletes the image with the given external identifier.
func (api *ImagesApi) DeleteImageById(extId *string, args ...map[string]interface{}) (*import11.DeleteImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.DeleteImageById(context.Background(), &import14.DeleteImageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Deletes the image with the given external identifier.
func (api *ImagesServiceApi) DeleteImageById(ctx context.Context, request *import14.DeleteImageByIdRequest, args ...map[string]interface{}) (*import11.DeleteImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.DeleteImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Downloads the image with the given external identifier.
func (api *ImagesApi) GetFileByImageId(imageExtId *string, args ...map[string]interface{}) (*import11.GetImageFileApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetFileByImageId(context.Background(), &import14.GetFileByImageIdRequest{
		ImageExtId: imageExtId,
	}, args...)
}

// Downloads the image with the given external identifier.
func (api *ImagesServiceApi) GetFileByImageId(ctx context.Context, request *import14.GetFileByImageIdRequest, args ...map[string]interface{}) (*import11.GetImageFileApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/{imageExtId}/file"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	binaryMediaTypes := []string{"application/octet-stream", "application/pdf", "application/zip"}
	if httpResponse, ok := apiClientResponse.(*http.Response); ok {
		if api.ApiClient.Contains(binaryMediaTypes, httpResponse.Header.Get("Content-Type")) {
			// Download file
			filePath, err := api.ApiClient.DownloadFile(httpResponse)
			if err != nil {
				return nil, err
			}

			response := import11.NewGetImageFileApiResponse()
			fileDetail := import11.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import12.Flag
			flags = append(flags, import12.Flag{Name: &flagName, Value: &flagValue})
			metadata := import13.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.GetImageFileApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Retrieves the image details for the provided external identifier.
func (api *ImagesApi) GetImageById(extId *string, args ...map[string]interface{}) (*import11.GetImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.GetImageById(context.Background(), &import14.GetImageByIdRequest{
		ExtId: extId,
	}, args...)
}

// Retrieves the image details for the provided external identifier.
func (api *ImagesServiceApi) GetImageById(ctx context.Context, request *import14.GetImageByIdRequest, args ...map[string]interface{}) (*import11.GetImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/{extId}"

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
	unmarshalledResp := new(import11.GetImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Imports images owned by a registered Prism Element cluster into the current Prism Central.
func (api *ImagesApi) ImportImage(body *import11.ImageImportConfig, args ...map[string]interface{}) (*import11.ImportImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ImportImage(context.Background(), &import14.ImportImageRequest{
		Body: body,
	}, args...)
}

// Imports images owned by a registered Prism Element cluster into the current Prism Central.
func (api *ImagesServiceApi) ImportImage(ctx context.Context, request *import14.ImportImageRequest, args ...map[string]interface{}) (*import11.ImportImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/$actions/import"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.ImportImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Lists images owned by Prism Central, along with image details such as name, description, type, and so on. This API supports operation such as filtering, sorting, selection, and pagination.
func (api *ImagesApi) ListImages(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import11.ListImagesApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.ListImages(context.Background(), &import14.ListImagesRequest{
		Page_:    page_,
		Limit_:   limit_,
		Filter_:  filter_,
		Orderby_: orderby_,
		Select_:  select_,
	}, args...)
}

// Lists images owned by Prism Central, along with image details such as name, description, type, and so on. This API supports operation such as filtering, sorting, selection, and pagination.
func (api *ImagesServiceApi) ListImages(ctx context.Context, request *import14.ListImagesRequest, args ...map[string]interface{}) (*import11.ListImagesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images"

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
	unmarshalledResp := new(import11.ListImagesApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Migrate the image with the given external identifier to the provided Content Repository. This moves an existing image to an existing Content Repository. The association of an image to a Content Repository is immutable.
func (api *ImagesApi) MigrateImage(extId *string, body *import11.ImageMigrateConfig, args ...map[string]interface{}) (*import11.MigrateImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.MigrateImage(context.Background(), &import14.MigrateImageRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Migrate the image with the given external identifier to the provided Content Repository. This moves an existing image to an existing Content Repository. The association of an image to a Content Repository is immutable.
func (api *ImagesServiceApi) MigrateImage(ctx context.Context, request *import14.MigrateImageRequest, args ...map[string]interface{}) (*import11.MigrateImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/{extId}/$actions/migrate"

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

	apiClientResponse, err := api.ApiClient.CallApiWithContext(ctx, &uri, http.MethodPost, request.Body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.MigrateImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}

// Updates the image with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure that the correct ETag is used.
func (api *ImagesApi) UpdateImageById(extId *string, body *import11.Image, args ...map[string]interface{}) (*import11.UpdateImageApiResponse, error) {
	if api.ServiceClient == nil {
		api.ServiceClient = NewImagesServiceApi(api.ApiClient)
	}
	return api.ServiceClient.UpdateImageById(context.Background(), &import14.UpdateImageByIdRequest{
		ExtId: extId,
		Body:  body,
	}, args...)
}

// Updates the image with the given external identifier. It is always recommended to perform a GET operation on a resource before performing a PUT operation to ensure that the correct ETag is used.
func (api *ImagesServiceApi) UpdateImageById(ctx context.Context, request *import14.UpdateImageByIdRequest, args ...map[string]interface{}) (*import11.UpdateImageApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.3/content/images/{extId}"

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
	if _, ok := apiClientResponse.(*client.EmptyResponse); ok {
		return nil, nil
	}

	// Response is already []byte (JSON content)
	unmarshalledResp := new(import11.UpdateImageApiResponse)
	if err = json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp); err != nil {
		return nil, err
	}
	return unmarshalledResp, err
}
