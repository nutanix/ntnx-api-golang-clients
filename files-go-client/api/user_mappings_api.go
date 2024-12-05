package api

import (
	"encoding/json"
	"fmt"
	"github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/client"
	import5 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/config"
	import6 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type UserMappingsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewUserMappingsApi(apiClient *client.ApiClient) *UserMappingsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserMappingsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Download user mapping to CSV file.
func (api *UserMappingsApi) DownloadUserMappings(fileServerExtId *string, args ...map[string]interface{}) (*import3.DownloadUserMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/user-mappings"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

			response := import3.NewDownloadUserMappingsApiResponse()
			fileDetail := import3.NewFileDetail()
			fileDetail.Path = filePath

			flagName := "hasError"
			flagValue := false
			var flags []import5.Flag
			flags = append(flags, import5.Flag{Name: &flagName, Value: &flagValue})
			metadata := import6.NewApiResponseMetadata()
			metadata.Flags = flags
			response.Metadata = metadata
			err = response.SetData(*fileDetail)
			if err != nil {
				return nil, err
			}

			return response, err
		}
	}

	unmarshalledResp := new(import3.DownloadUserMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Upload user mappings using CSV file.  Examples of Various types of mappings 1. Explicit mapping - One-to-one mapping. Add mapping between SMB and NFS users or groups   smbName,nfsId,userOrGroup,mappingType   child4\\\\u11,1,User,OneToOne  2. Explicit mapping - Wildcard-based mapping. Add mapping between SMB and NFS users or groups.    Priority of the mapping will be decided based on the order of the entries.   child4\\\\*,4,Group,WildCard  3. Explicit mapping -  Deny Smb users/group   child4\\\\u13,,User,Deny  4. Explicit mapping -  Deny NFS users/groups   ,1000,Group,Deny  5. Rule based mapping. Add this entry to enable template mapping.    ,,,Template  6. Default mapping - Deny SMB users/Groups with no NFS mapping   ,-,User,Default   ,-,Group,Default  7. Default mapping - Deny NFS users/Groups with no SMB mapping   -,,User,Default   -,,Group,Default
func (api *UserMappingsApi) UploadUserMappings(fileServerExtId *string, path *string, args ...map[string]interface{}) (*import3.UploadUserMappingsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/files/v4.0/config/file-servers/{fileServerExtId}/$actions/upload-user-mappings"

	// verify the required parameter 'fileServerExtId' is set
	if nil == fileServerExtId {
		return nil, client.ReportError("fileServerExtId is required and must be specified")
	}
	// verify the required parameter 'path' is set
	if nil == path {
		return nil, client.ReportError("path is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"fileServerExtId"+"}", url.PathEscape(client.ParameterToString(*fileServerExtId, "")), -1)
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

	authNames := []string{"apiKeyAuthScheme", "basicAuthScheme"}

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, file, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.UploadUserMappingsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
