package api

import (
	"encoding/json"
	"github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/client"
	import3 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
	"net/http"
	"net/url"
	"strings"
)

type DomainManagerBackupsApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewDomainManagerBackupsApi(apiClient *client.ApiClient) *DomainManagerBackupsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &DomainManagerBackupsApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Creates a cluster or object store as the backup target. For a given Prism Central, there can be up to 3 clusters as backup targets  and 1 object store as backup target. If any cluster or object store is not eligible for backup or  lacks appropriate permissions, the API request will fail.  For object store backup targets, specifying backup policy is mandatory along  with the location of the object store.
func (api *DomainManagerBackupsApi) CreateBackupTarget(domainManagerExtId *string, body *import3.BackupTarget, args ...map[string]interface{}) (*import3.CreateBackupTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/backup-targets"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.CreateBackupTargetApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Creates a restore source pointing to a cluster or object store to restore the domain manager. The created  restore source is intended to be deleted after use. If the restore source is not deleted using the deleteRestoreSource API, then it is auto-deleted after sometime. Also note that a restore  source will not contain a backup policy. It is only used to access the backup data at the location from where  the Prism Central may be restored. Credentials used to access the restore source are not validated at the time  of creation of the restore source. They are validated when the restore source is used to fetch data.
func (api *DomainManagerBackupsApi) CreateRestoreSource(body *import3.RestoreSource, args ...map[string]interface{}) (*import3.CreateRestoreSourceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources"

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

	unmarshalledResp := new(import3.CreateRestoreSourceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Removes cluster/object store from the backup targets. This will stop the cluster/object store  from backing up Prism Central data.
func (api *DomainManagerBackupsApi) DeleteBackupTargetById(domainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import3.DeleteBackupTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/backup-targets/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.DeleteBackupTargetApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Deletes a restore source on a cluster/object store.
func (api *DomainManagerBackupsApi) DeleteRestoreSourceById(extId *string, args ...map[string]interface{}) (*import3.DeleteRestoreSourceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{extId}"

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

	unmarshalledResp := new(import3.DeleteRestoreSourceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the backup targets (cluster or object store) from a domain manager and returns the backup configuration and lastSyncTimestamp parameter to the user.
func (api *DomainManagerBackupsApi) GetBackupTargetById(domainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetBackupTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/backup-targets/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.GetBackupTargetApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves detailed information about a specific recovery point and provides essential domain manager information stored in the backup, which is required for the restoration process.
func (api *DomainManagerBackupsApi) GetRestorePointById(restoreSourceExtId *string, restorableDomainManagerExtId *string, extId *string, args ...map[string]interface{}) (*import3.GetRestorePointApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points/{extId}"

	// verify the required parameter 'restoreSourceExtId' is set
	if nil == restoreSourceExtId {
		return nil, client.ReportError("restoreSourceExtId is required and must be specified")
	}
	// verify the required parameter 'restorableDomainManagerExtId' is set
	if nil == restorableDomainManagerExtId {
		return nil, client.ReportError("restorableDomainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"restoreSourceExtId"+"}", url.PathEscape(client.ParameterToString(*restoreSourceExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"restorableDomainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*restorableDomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.GetRestorePointApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Retrieves the restore source from the PE cache store and returns the restore source configuration and external identifier to the user.
func (api *DomainManagerBackupsApi) GetRestoreSourceById(extId *string, args ...map[string]interface{}) (*import3.GetRestoreSourceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{extId}"

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

	unmarshalledResp := new(import3.GetRestoreSourceApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists backup targets (cluster or object store) configured for a given domain manager.
func (api *DomainManagerBackupsApi) ListBackupTargets(domainManagerExtId *string, args ...map[string]interface{}) (*import3.ListBackupTargetsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/backup-targets"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListBackupTargetsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Lists all the domain managers backed up at the object store/cluster.
func (api *DomainManagerBackupsApi) ListRestorableDomainManagers(restoreSourceExtId *string, page_ *int, limit_ *int, filter_ *string, args ...map[string]interface{}) (*import3.ListRestorableDomainManagersApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers"

	// verify the required parameter 'restoreSourceExtId' is set
	if nil == restoreSourceExtId {
		return nil, client.ReportError("restoreSourceExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"restoreSourceExtId"+"}", url.PathEscape(client.ParameterToString(*restoreSourceExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListRestorableDomainManagersApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// The list restore points API allows you to retrieve a list of available restore points,  which are snapshots of the domain manager taken at different times. These restore points can be used to revert the domain manager to a previous state. The list response includes the creation time and identifier ID for the configuration data.<br>  1. For cluster-based backups, only the most recent restore point is available, as backups are continuous.<br> 2. For object store-based backups, multiple restore points may be available, depending on the configured  Recovery Point Objective (RPO) and the retention period set on the bucket.
func (api *DomainManagerBackupsApi) ListRestorePoints(restoreSourceExtId *string, restorableDomainManagerExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListRestorePointsApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points"

	// verify the required parameter 'restoreSourceExtId' is set
	if nil == restoreSourceExtId {
		return nil, client.ReportError("restoreSourceExtId is required and must be specified")
	}
	// verify the required parameter 'restorableDomainManagerExtId' is set
	if nil == restorableDomainManagerExtId {
		return nil, client.ReportError("restorableDomainManagerExtId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"restoreSourceExtId"+"}", url.PathEscape(client.ParameterToString(*restoreSourceExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"restorableDomainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*restorableDomainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.ListRestorePointsApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// The restore domain manager is a task-driven operation to restore a domain manager from a cluster or object store  backup location based on the selected restore point.
func (api *DomainManagerBackupsApi) Restore(restoreSourceExtId *string, restorableDomainManagerExtId *string, extId *string, body *import3.RestoreSpec, args ...map[string]interface{}) (*import3.RestoreApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/restore-sources/{restoreSourceExtId}/restorable-domain-managers/{restorableDomainManagerExtId}/restore-points/{extId}/$actions/restore"

	// verify the required parameter 'restoreSourceExtId' is set
	if nil == restoreSourceExtId {
		return nil, client.ReportError("restoreSourceExtId is required and must be specified")
	}
	// verify the required parameter 'restorableDomainManagerExtId' is set
	if nil == restorableDomainManagerExtId {
		return nil, client.ReportError("restorableDomainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"restoreSourceExtId"+"}", url.PathEscape(client.ParameterToString(*restoreSourceExtId, "")), -1)
	uri = strings.Replace(uri, "{"+"restorableDomainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*restorableDomainManagerExtId, "")), -1)
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

	apiClientResponse, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == apiClientResponse {
		return nil, err
	}

	unmarshalledResp := new(import3.RestoreApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}

// Updates the credentials, RP0, or certificates (for Nutanix Objects only) of the given object store based on the user requirements.  RPO is mandatory to be passed in payload. Credentials and certificates are optional to pass in update backup target api. If credentials or certificates are not passed then these will not be updated for a backup target.
func (api *DomainManagerBackupsApi) UpdateBackupTargetById(domainManagerExtId *string, extId *string, body *import3.BackupTarget, args ...map[string]interface{}) (*import3.UpdateBackupTargetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.1/management/domain-managers/{domainManagerExtId}/backup-targets/{extId}"

	// verify the required parameter 'domainManagerExtId' is set
	if nil == domainManagerExtId {
		return nil, client.ReportError("domainManagerExtId is required and must be specified")
	}
	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"domainManagerExtId"+"}", url.PathEscape(client.ParameterToString(*domainManagerExtId, "")), -1)
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

	unmarshalledResp := new(import3.UpdateBackupTargetApiResponse)
	json.Unmarshal(apiClientResponse.([]byte), &unmarshalledResp)
	return unmarshalledResp, err
}
