package api

import (
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ObjectStoresApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewObjectStoresApi(apiClient *client.ApiClient) *ObjectStoresApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ObjectStoresApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// This operation creates a new default certificate and keys. It also creates the alternate FQDNs and alternate IPs for the Object store. The certificate of an Object store can be created when it is in a OBJECT_STORE_AVAILABLE or OBJECT_STORE_CERT_CREATION_FAILED state. If the publicCert, privateKey, and ca values are provided in the request body, these values are used to create the new certificate. If these values are not provided, a new certificate will be generated if 'shouldGenerate' is set to true and if it is set to false, the existing certificate will be used as the new certificate. Optionally, a list of additional alternate FQDNs and alternate IPs can be provided. These alternateFqdns and alternateIps must be included in the CA certificate if it has been provided.
func (api *ObjectStoresApi) CreateCertificate(objectStoreExtId *string, path *string, args ...map[string]interface{}) (*import1.CreateCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{objectStoreExtId}/certificates"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == objectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'path' is set
	if nil == path {
		return nil, client.ReportError("path is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*objectStoreExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/octet-stream"}

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

	file, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	headerParams["Content-Length"] = fmt.Sprintf("%d", fileInfo.Size())
	if headerParams["Content-Disposition"] == "" {
		headerParams["Content-Disposition"] = fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name())
	}

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, file, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateCertificateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Run the prechecks, create and start the deployment of an Object store on Prism Central.
func (api *ObjectStoresApi) CreateObjectstore(body *import1.ObjectStore, args ...map[string]interface{}) (*import1.CreateObjectstoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores"

	// verify the required parameter 'body' is set
	if nil == body {
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.CreateObjectstoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Send a request to delete an Object store for the provided UUID and clean up its data. An Object store can only be deleted when it is in the state UNDEPLOYED_OBJECT_STORE, OBJECT_STORE_AVAILABLE, OBJECT_STORE_DEPLOYMENT_FAILED or OBJECT_STORE_DELETION_FAILED and does not contain any buckets.
func (api *ObjectStoresApi) DeleteObjectstoreById(extId *string, args ...map[string]interface{}) (*import1.DeleteObjectstoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.DeleteObjectstoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Download the certificate authority of an Object store certificate.
func (api *ObjectStoresApi) GetCaByCertificateId(objectStoreExtId *string, certificateExtId *string, args ...map[string]interface{}) (*import1.GetCaApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{objectStoreExtId}/certificates/{certificateExtId}/certificate-authority"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == objectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'certificateExtId' is set
	if nil == certificateExtId {
		return nil, client.ReportError("certificateExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*objectStoreExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"certificateExtId"+"}", url.PathEscape(client.ParameterToString(*certificateExtId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
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

			response := import1.NewGetCaApiResponse()
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

	unmarshalledResp := new(import1.GetCaApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the details of the SSL certificate which can be used to connect to an Object store.
func (api *ObjectStoresApi) GetCertificateById(objectStoreExtId *string, extId *string, args ...map[string]interface{}) (*import1.GetCertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{objectStoreExtId}/certificates/{extId}"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == objectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*objectStoreExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GetCertificateApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get an Object store for the provided UUID. `state` string (State of the Object store) Enum for the state of the Object store. | Enum | Description | | ----------- | ----------- | | `\"DEPLOYING_OBJECT_STORE\"` | The Object store will be in this state during an ongoing deployment of the Object store. The Object store will be unavailable through S3 APIs in this state. The Object store will move to the OBJECT_STORE_AVAILABLE state if the deployment succeeds, and to the OBJECT_STORE_DEPLOYMENT_FAILED state if the deployment fails. | | `\"OBJECT_STORE_OPERATION_FAILED\"` | The Object store is in this state when there is an error while performing an operation on the Object store. The Object store may not be available through S3 APIs in this state. | | `\"OBJECT_STORE_CERT_CREATION_FAILED\"` |  An Object store enters this state if there is an error while creating the Object store certificate. Creating a new certificate can be retried from this state. The Object store may not be available through S3 APIs in this state. | | `\"OBJECT_STORE_OPERATION_PENDING\"` |  The Object store is in this state during an ongoing operation on the Object store. The Object store may not be available through S3 APIs in this state. The Object store will enter the OBJECT_STORE_OPERATION_FAILED state if the operation fails, or the OBJECT_STORE_AVAILABLE state if the operation is successful. | | `\"UNDEPLOYED_OBJECT_STORE\"` |  The Object store is in this state if it has not been deployed. | | `\"CREATING_OBJECT_STORE_CERT\"` |  The Object store is in this state during a certificate creation for the Object store. The Object store will be unavailable through S3 APIs in this state. It will move to the OBJECT_STORE_AVAILABLE state if the certificate was created successfully, or to the OBJECT_STORE_CERT_CREATION_FAILED state if an error occurs while creating the certificate. | | `\"OBJECT_STORE_AVAILABLE\"`  | An Object store is in this state if its deployment was successful, and there are no ongoing operations on the Object store. The Object store will be available through S3 APIs in this state. In this state, the Object store can be deleted or a new certificate can be created for this Object store. | | `\"OBJECT_STORE_DELETION_FAILED\"` |  An Object store enters this state if there is an error deleting the Object store. The Object store will not be available through S3 APIs in this state. Deleting the Object store can be retried from this state. | | `\"OBJECT_STORE_DEPLOYMENT_FAILED\"` |  An Object store enters this state when its deployment fails. The Object store deployment can be retried or the Object store can be deleted from this state. | | `\"DELETING_OBJECT_STORE\"` |  A deployed Object store is in this state when the Object store is being deleted. The Object store will be unavailable through S3 APIs in this state. It can be either deleted or move to the OBJECT_STORE_DELETION_FAILED state if the deletion fails. |
func (api *ObjectStoresApi) GetObjectstoreById(extId *string, args ...map[string]interface{}) (*import1.GetObjectstoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.GetObjectstoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a list of the SSL certificates which can be used to access an Object store.
func (api *ObjectStoresApi) ListCertificatesByObjectstoreId(objectStoreExtId *string, page_ *int, limit_ *int, filter_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListCertificatesApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{objectStoreExtId}/certificates"

	// verify the required parameter 'objectStoreExtId' is set
	if nil == objectStoreExtId {
		return nil, client.ReportError("objectStoreExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"objectStoreExtId"+"}", url.PathEscape(client.ParameterToString(*objectStoreExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.ListCertificatesApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a list of all the Object store details on the registered Prism Central.
func (api *ObjectStoresApi) ListObjectstores(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.ListObjectstoresApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.ListObjectstoresApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update an Object store. The deployment of an Object store can be restarted from the state OBJECT_STORE_DEPLOYMENT_FAILED.
func (api *ObjectStoresApi) UpdateObjectstoreById(extId *string, body *import1.ObjectStore, args ...map[string]interface{}) (*import1.UpdateObjectstoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0/config/object-stores/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
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

	authNames := []string{"basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import1.UpdateObjectstoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
