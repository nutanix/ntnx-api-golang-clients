//Api classes for objects's golang SDK
package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/client"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/operations"
	"net/http"
	"net/url"
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

// This operation creates a new default certificate and keys. It also creates the alternate FQDNs for the Object store. The certificate of an Object store can be created when it is in a OBJECT_STORE_AVAILABLE or OBJECT_STORE_CREATE_CERT_ERROR state. If the publicCert, privateKey, and ca values are provided in the request body, these values are used to create the new certificate. If these values are not provided, they will be generated for the new certificate. Optionally, a list of additional alternate FQDNs can be provided.
func (api *ObjectStoresApi) CreateCertificate(objectStoreExtId *string, publicCert *string, privateKey *string, metadata *import1.Metadata, alternateFqdns *[]import1.FQDN, tenantId *string, links *[]import2.ApiLink, extId *string, ca *string, args ...map[string]interface{}) (*import3.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates"

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
	contentTypes := []string{"multipart/form-data"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	// Form Params
	if publicCert != nil {
		formParams.Add("publicCert", client.ParameterToString(*publicCert, ""))
	}
	if privateKey != nil {
		formParams.Add("privateKey", client.ParameterToString(*privateKey, ""))
	}
	if metadata != nil {
		formParams.Add("metadata", client.ParameterToString(*metadata, ""))
	}
	if alternateFqdns != nil {
		formParams.Add("alternateFqdns", client.ParameterToString(*alternateFqdns, "multi"))
	}
	if tenantId != nil {
		formParams.Add("tenantId", client.ParameterToString(*tenantId, ""))
	}
	if links != nil {
		formParams.Add("links", client.ParameterToString(*links, "multi"))
	}
	if extId != nil {
		formParams.Add("extId", client.ParameterToString(*extId, ""))
	}
	if ca != nil {
		formParams.Add("ca", client.ParameterToString(*ca, ""))
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.TaskReferenceApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Run the prechecks, create and start the deployment of an Object store on Prism Central.
func (api *ObjectStoresApi) CreateObjectstore(body *import3.Objectstore, args ...map[string]interface{}) (*import3.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.TaskReferenceApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Send a request to delete an Object store for the provided UUID and clean up its data. An Object store can only be deleted when it is in the state OBJECT_STORE_INPUT, OBJECT_STORE_DELETE_INPUT_ERROR, OBJECT_STORE_AVAILABLE, OBJECT_STORE_DEPLOYMENT_ERROR or OBJECT_STORE_DELETE_ERROR and does not contain any buckets.
func (api *ObjectStoresApi) DeleteObjectstore(extId *string, args ...map[string]interface{}) (*import3.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.TaskReferenceApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get the details of the SSL certificate which can be used to connect to an Object store.
func (api *ObjectStoresApi) GetCertificate(objectStoreExtId *string, extId *string, args ...map[string]interface{}) (*import3.CertificateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.CertificateApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a list of the SSL certificates which can be used to access an Object store.
func (api *ObjectStoresApi) GetCertificates(objectStoreExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.CertificateListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates"

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
	if orderby_ != nil {

		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if select_ != nil {

		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.CertificateListApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get an Object store for the provided UUID.      `state`   string (State of the Object store)   Enum for the state of the Object store.   | Enum | Description | | ----------- | ----------- | | `\"OBJECT_STORE_DEPLOYING\"` | The Object store will be in this state during an ongoing deployment of the Object store. The Object store will be unavailable through S3 APIs in this state. The Object store will move to the OBJECT_STORE_AVAILABLE state if the deployment succeeds, and to the OBJECT_STORE_DEPLOYMENT_ERROR state if the deployment fails. | | `\"OBJECT_STORE_OPERATION_ERROR\"` | The Object store is in this state when there is an error while performing an operation on the Object store. The Object store may not be available through S3 APIs in this state. | | `\"OBJECT_STORE_CREATE_CERT_ERROR\"` |  An Object store enters this state if there is an error while creating the Object store certificate. Creating a new certificate can be retried from this state. The Object store may not be available through S3 APIs in this state. | | `\"OBJECT_STORE_OPERATION_PENDING\"` |  The Object store is in this state during an ongoing operation on the Object store. The Object store may not be available through S3 APIs in this state. The Object store will enter the OBJECT_STORE_OPERATION_ERROR state if the operation fails, or the OBJECT_STORE_AVAILABLE state if the operation is successful. | | `\"OBJECT_STORE_INPUT\"` |  The Object store is in this state if it has not been deployed. | | `\"OBJECT_STORE_CREATING_CERT\"` |  The Object store is in this state during a certificate creation for the Object store. The Object store will be unavailable through S3 APIs in this state. It will move to the OBJECT_STORE_AVAILABLE state if the certificate was created successfully, or to the OBJECT_STORE_CREATE_CERT_ERROR state if an error occurs while creating the certificate. | | `\"OBJECT_STORE_AVAILABLE\"`  | An Object store is in this state if its deployment was successful, and there are no ongoing operations on the Object store. The Object store will be available through S3 APIs in this state. In this state, the Object store can be deleted or a new certificate can be created for this Object store. | | `\"OBJECT_STORE_DELETE_INPUT_ERROR\"` |  An undeployed Object store enters this state if there is an error deleting it. Deleting the Object store can be retried from this state. | | `\"OBJECT_STORE_DELETING_INPUT\"` |  An undeployed Object store being deleted is in this state. It will be either deleted or move to the OBJECT_STORE_DELETE_INPUT_ERROR state if the deletion fails. | | `\"OBJECT_STORE_DELETE_ERROR\"` |  An Object store enters this state if there is an error deleting the Object store. The Object store will not be available through S3 APIs in this state. Deleting the Object store can be retried from this state. | | `\"OBJECT_STORE_DEPLOYMENT_ERROR\"` |  An Object store enters this state when its deployment fails. The Object store deployment can be retried or the Object store can be deleted from this state. | | `\"OBJECT_STORE_DELETING\"` |  A deployed Object store is in this state when the Object store is being deleted. The Object store will be unavailable through S3 APIs in this state. It can be either deleted or move to the OBJECT_STORE_DELETE_ERROR state if the deletion fails. |
func (api *ObjectStoresApi) GetObjectstore(extId *string, expand_ *string, args ...map[string]interface{}) (*import3.ObjectstoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{extId}"

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

	// Query Params
	if expand_ != nil {

		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.ObjectstoreApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Get a list of all the Object store details on the registered Prism Central.
func (api *ObjectStoresApi) GetObjectstores(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import3.ObjectstoreListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.ObjectstoreListApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Update an Object store. The deployment of an Object store can be restarted from the state OBJECT_STORE_DEPLOYMENT_ERROR.
func (api *ObjectStoresApi) UpdateObjectstore(extId *string, body *import3.Objectstore, args ...map[string]interface{}) (*import3.TaskReferenceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/objects/v4.0.a2/operations/object-stores/{extId}"

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
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}

	unmarshalledResp := new(import3.TaskReferenceApiResponse)
	json.Unmarshal(responseBody.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
