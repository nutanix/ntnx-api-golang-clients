/*
 * Generated file models/files/v4/infra/infra_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Files Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module for infra APIs for file server
*/
package infra
import (
  "bytes"
  import5 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/config"
  import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
  import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
  "fmt"
  import6 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/prism/v4/config"
  import4 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/storage/v4/config"
)

/**
The type of the volume groups attachment. The acceptable values are "HYPERVISOR", "ISCSI".
*/
type AttachmentType int

const(
  ATTACHMENTTYPE_UNKNOWN AttachmentType = 0
  ATTACHMENTTYPE_REDACTED AttachmentType = 1
  ATTACHMENTTYPE_HYPERVISOR AttachmentType = 2
  ATTACHMENTTYPE_ISCSI AttachmentType = 3
)

// returns the name of the enum given an ordinal number
func (e *AttachmentType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HYPERVISOR",
    "ISCSI",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AttachmentType) index(name string) AttachmentType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HYPERVISOR",
    "ISCSI",
  }
  for idx := range names {
    if names[idx] == name {
      return AttachmentType(idx)
    }
  }
  return ATTACHMENTTYPE_UNKNOWN
}

func (e *AttachmentType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AttachmentType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AttachmentType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AttachmentType) Ref() *AttachmentType {
  return &e
}


/**
Root CA or ICA certificate.
*/
type CaCert struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Root CA or ICA certificate.
  */
  CaCertificate *string `json:"caCertificate,omitempty"`
  /**
  Cluster UUID that owns the certificates.
  */
  OwnerClusterExtId *string `json:"ownerClusterExtId,omitempty"`
  
  OwnerType *OwnerType `json:"ownerType,omitempty"`
}


func NewCaCert() *CaCert {
  p := new(CaCert)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.CaCert"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.CaCert"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Certificate chain to verify TLS certificates.
*/
type CaChain struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Additional CA certificates to be trusted.
  */
  AdditionalCaCertList []CaCert `json:"additionalCaCertList,omitempty"`
  /**
  List of TLS certificates for root CA and ICA.
  */
  CaCertificateList []string `json:"caCertificateList,omitempty"`
  /**
  Cluster UUID that owns the certificates.
  */
  OwnerClusterExtId *string `json:"ownerClusterExtId,omitempty"`
}


func NewCaChain() *CaChain {
  p := new(CaChain)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.CaChain"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.CaChain"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the file server information.
*/
type FileServerInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FileServer *import1.FileServer `json:"fileServer,omitempty"`
  /**
  Genesis status output from NVM.
  */
  GenesisStatus *string `json:"genesisStatus,omitempty"`
  /**
  Get request flag for file server specification.
  */
  IsFileServerInfo *bool `json:"isFileServerInfo,omitempty"`
  /**
  Get request flag for genesis status.
  */
  IsGenesisStatus *bool `json:"isGenesisStatus,omitempty"`
  /**
  Get request flag for HA state.
  */
  IsHaState *bool `json:"isHaState,omitempty"`
  /**
  Flag that represents the file server is in HA state.
  */
  IsInHaState *bool `json:"isInHaState,omitempty"`
}


func NewFileServerInfo() *FileServerInfo {
  p := new(FileServerInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FileServerInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FileServerInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the file server storage specification - a set of volume groups with volume disks and ISCSI clients.
*/
type FileServerStorageSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AttachmentType *AttachmentType `json:"attachmentType"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  ExtId of the task.
  */
  TaskExtId *string `json:"taskExtId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  List of volume group objects.
  */
  VolumeGroups []FilesVolumeGroup `json:"volumeGroups"`
}

func (p *FileServerStorageSpec) MarshalJSON() ([]byte, error) {
  type FileServerStorageSpecProxy FileServerStorageSpec
  return json.Marshal(struct {
    *FileServerStorageSpecProxy
    AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
    FileServerExtId *string `json:"fileServerExtId,omitempty"`
    VolumeGroups []FilesVolumeGroup `json:"volumeGroups,omitempty"`
  }{
    FileServerStorageSpecProxy : (*FileServerStorageSpecProxy)(p),
    AttachmentType : p.AttachmentType,
    FileServerExtId : p.FileServerExtId,
    VolumeGroups : p.VolumeGroups,
  })
}

func NewFileServerStorageSpec() *FileServerStorageSpec {
  p := new(FileServerStorageSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FileServerStorageSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FileServerStorageSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the files infrastructure lock specification.
*/
type FilesInfraLock struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Endpoint *FilesInfraLockEndpoint `json:"endpoint,omitempty"`
  /**
  Unique Id for the lock.
  */
  LockId *string `json:"lockId,omitempty"`
  /**
  Operation description message.
  */
  Message *string `json:"message,omitempty"`
  /**
  Operation name that acquired the lock.
  */
  OperationName *string `json:"operationName,omitempty"`
  
  Scope *FilesInfraLockScope `json:"scope,omitempty"`
  /**
  Identification for the scope. For example: taskId
  */
  ScopeId *string `json:"scopeId,omitempty"`
  /**
  Expiry time for the lock.
  */
  TimeoutSecs *int64 `json:"timeoutSecs,omitempty"`
  /**
  Time when lock was last acquired.
  */
  Timestamp *int64 `json:"timestamp,omitempty"`
}


func NewFilesInfraLock() *FilesInfraLock {
  p := new(FilesInfraLock)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLock"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLock"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-infra-locks/$actions/acquire Post operation
*/
type FilesInfraLockAcquireApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFilesInfraLockAcquireApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFilesInfraLockAcquireApiResponse() *FilesInfraLockAcquireApiResponse {
  p := new(FilesInfraLockAcquireApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLockAcquireApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLockAcquireApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FilesInfraLockAcquireApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FilesInfraLockAcquireApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFilesInfraLockAcquireApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
A model that represents the response for files infrastructure lock acquire operation.
*/
type FilesInfraLockAcquireResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  LockSpec *FilesInfraLock `json:"lockSpec,omitempty"`
  /**
  Status of the lock - locked/released.
  */
  Status *bool `json:"status,omitempty"`
}


func NewFilesInfraLockAcquireResponse() *FilesInfraLockAcquireResponse {
  p := new(FilesInfraLockAcquireResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLockAcquireResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLockAcquireResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The supported lock endpoints.
*/
type FilesInfraLockEndpoint int

const(
  FILESINFRALOCKENDPOINT_UNKNOWN FilesInfraLockEndpoint = 0
  FILESINFRALOCKENDPOINT_REDACTED FilesInfraLockEndpoint = 1
  FILESINFRALOCKENDPOINT_FILES FilesInfraLockEndpoint = 2
  FILESINFRALOCKENDPOINT_PLATFORM FilesInfraLockEndpoint = 3
)

// returns the name of the enum given an ordinal number
func (e *FilesInfraLockEndpoint) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FILES",
    "PLATFORM",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FilesInfraLockEndpoint) index(name string) FilesInfraLockEndpoint {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "FILES",
    "PLATFORM",
  }
  for idx := range names {
    if names[idx] == name {
      return FilesInfraLockEndpoint(idx)
    }
  }
  return FILESINFRALOCKENDPOINT_UNKNOWN
}

func (e *FilesInfraLockEndpoint) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FilesInfraLockEndpoint:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FilesInfraLockEndpoint) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FilesInfraLockEndpoint) Ref() *FilesInfraLockEndpoint {
  return &e
}


/**
A model that represents the list of files infrastructure lock specification(s).
*/
type FilesInfraLockList struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of files infrastructure lock specification(s).
  */
  LockSpecList []FilesInfraLock `json:"lockSpecList,omitempty"`
}


func NewFilesInfraLockList() *FilesInfraLockList {
  p := new(FilesInfraLockList)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLockList"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLockList"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-infra-locks/$actions/release Post operation
*/
type FilesInfraLockReleaseApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFilesInfraLockReleaseApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFilesInfraLockReleaseApiResponse() *FilesInfraLockReleaseApiResponse {
  p := new(FilesInfraLockReleaseApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLockReleaseApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLockReleaseApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FilesInfraLockReleaseApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FilesInfraLockReleaseApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFilesInfraLockReleaseApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
A model that represents the response for files infrastructure lock release operation.
*/
type FilesInfraLockReleaseResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Status of the lock - locked/released.
  */
  Status *bool `json:"status,omitempty"`
}


func NewFilesInfraLockReleaseResponse() *FilesInfraLockReleaseResponse {
  p := new(FilesInfraLockReleaseResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesInfraLockReleaseResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesInfraLockReleaseResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
The supported lock scopes. The acceptable values are \"MANUAL\" or \"TASK\".
*/
type FilesInfraLockScope int

const(
  FILESINFRALOCKSCOPE_UNKNOWN FilesInfraLockScope = 0
  FILESINFRALOCKSCOPE_REDACTED FilesInfraLockScope = 1
  FILESINFRALOCKSCOPE_MANUAL FilesInfraLockScope = 2
  FILESINFRALOCKSCOPE_TASK FilesInfraLockScope = 3
)

// returns the name of the enum given an ordinal number
func (e *FilesInfraLockScope) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MANUAL",
    "TASK",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *FilesInfraLockScope) index(name string) FilesInfraLockScope {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MANUAL",
    "TASK",
  }
  for idx := range names {
    if names[idx] == name {
      return FilesInfraLockScope(idx)
    }
  }
  return FILESINFRALOCKSCOPE_UNKNOWN
}

func (e *FilesInfraLockScope) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for FilesInfraLockScope:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *FilesInfraLockScope) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e FilesInfraLockScope) Ref() *FilesInfraLockScope {
  return &e
}


/**
A model that represents the files platform service request and response. Files platform service comprises of file server, NVM and storage information. This is requested from platform to files.
*/
type FilesService struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FileServerInfo *FileServerInfo `json:"fileServerInfo,omitempty"`
  
  NvmInfo *FilesServiceNvmInfo `json:"nvmInfo,omitempty"`
  
  StorageInfo *StorageInfo `json:"storageInfo,omitempty"`
}


func NewFilesService() *FilesService {
  p := new(FilesService)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesService"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesService"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-service Post operation
*/
type FilesServiceInfoGetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfFilesServiceInfoGetApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewFilesServiceInfoGetApiResponse() *FilesServiceInfoGetApiResponse {
  p := new(FilesServiceInfoGetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesServiceInfoGetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesServiceInfoGetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *FilesServiceInfoGetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *FilesServiceInfoGetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfFilesServiceInfoGetApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
A model that represents the files service NVM information.
*/
type FilesServiceNvmInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Get request flag for NVM health.
  */
  IsNvmHealth *bool `json:"isNvmHealth,omitempty"`
  /**
  Get request flag for NVM UUID to VM UUID mapping.
  */
  IsNvmUuidToVmUuidMapping *bool `json:"isNvmUuidToVmUuidMapping,omitempty"`
  /**
  Request field for the count of NVM objects to be deleted.
  */
  NumOfNvmsToBeRemoved *int `json:"numOfNvmsToBeRemoved,omitempty"`
  /**
  List of NVM objects.
  */
  NvmList []Nvm `json:"nvmList,omitempty"`
}


func NewFilesServiceNvmInfo() *FilesServiceNvmInfo {
  p := new(FilesServiceNvmInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesServiceNvmInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesServiceNvmInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
File server storage object representing volume group resource.
*/
type FilesVolumeGroup struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The UUID of cluster that will host the Volume Group. This is mandatory to be specified for creating a volume group on PC.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  /**
  Service/user who created this Volume Group.
  */
  CreatedBy *string `json:"createdBy,omitempty"`
  /**
  The description of the Volume Group.
  */
  Description *string `json:"description,omitempty"`
  /**
  List of volume disk objects.
  */
  Disks []import4.VolumeDisk `json:"disks,omitempty"`
  
  EnabledAuthentications *import4.AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  List of ISCSI client objects.
  */
  IscsiClients []import4.IscsiClient `json:"iscsiClients,omitempty"`
  /**
  iSCSI target full name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetName *string `json:"iscsiTargetName,omitempty"`
  /**
  iSCSI target prefix-name. The spec should not comprise both iscsi target prefix and iscsi target name.
  */
  IscsiTargetPrefix *string `json:"iscsiTargetPrefix,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Whether to enable volume group load balancing for VM attachments. This cannot be enabled if there are iSCSI client attachments already associated with the Volume Group, and vice-versa.
  */
  LoadBalanceVmAttachments *bool `json:"loadBalanceVmAttachments,omitempty"`
  /**
  The name of the Volume Group.
  */
  Name *string `json:"name,omitempty"`
  
  SharingStatus *import4.SharingStatus `json:"sharingStatus,omitempty"`
  
  StorageFeatures *import4.StorageFeatures `json:"storageFeatures,omitempty"`
  /**
  Target Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  TargetSecret *string `json:"targetSecret,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  List of file server VMs associated with the file server. This is a read-only field.
  */
  VmAttachments []import4.VmAttachment `json:"vmAttachments,omitempty"`
}


func NewFilesVolumeGroup() *FilesVolumeGroup {
  p := new(FilesVolumeGroup)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.FilesVolumeGroup"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.FilesVolumeGroup"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Http proxy model.
*/
type HttpProxy struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IP address or FQDN name for the http proxy.
  */
  Address *string `json:"address,omitempty"`
  /**
  Id string for the http proxy.
  */
  IdString *string `json:"idString,omitempty"`
  /**
  Password for prism central API authentication.
  */
  Password *string `json:"password,omitempty"`
  /**
  Port for the http proxy.
  */
  Port *int `json:"port,omitempty"`
  /**
  List of supported types by the proxy.
  */
  ProxyTypeList []ProxyType `json:"proxyTypeList,omitempty"`
  /**
  Username for prism central API authentication.
  */
  Username *string `json:"username,omitempty"`
}


func NewHttpProxy() *HttpProxy {
  p := new(HttpProxy)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.HttpProxy"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.HttpProxy"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-infra-locks Get operation
*/
type ListFilesInfraLockInfoApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfListFilesInfraLockInfoApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewListFilesInfraLockInfoApiResponse() *ListFilesInfraLockInfoApiResponse {
  p := new(ListFilesInfraLockInfoApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.ListFilesInfraLockInfoApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.ListFilesInfraLockInfoApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ListFilesInfraLockInfoApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ListFilesInfraLockInfoApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfListFilesInfraLockInfoApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Platform migrate request
*/
type MigratePlatform struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  PcConfig *PcConfig `json:"pcConfig,omitempty"`
}


func NewMigratePlatform() *MigratePlatform {
  p := new(MigratePlatform)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.MigratePlatform"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.MigratePlatform"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Platform update notification request model.
*/
type NotifyPlatform struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  PcConfig *PcConfig `json:"pcConfig,omitempty"`
  
  PeConfig *PeConfig `json:"peConfig,omitempty"`
  
  PreUpgradeNvm *PreUpgradeNvm `json:"preUpgradeNvm,omitempty"`
  
  PulseConfig *PulseConfig `json:"pulseConfig,omitempty"`
  
  TlsCertificateConfig *TlsCertificateConfig `json:"tlsCertificateConfig,omitempty"`
}


func NewNotifyPlatform() *NotifyPlatform {
  p := new(NotifyPlatform)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NotifyPlatform"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NotifyPlatform"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM resource - NVM UUID, mapped VM UUID etc.
*/
type Nvm struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The NVM UUID.
  */
  Uuid *string `json:"uuid,omitempty"`
  /**
  Acropolis UUID of the file server VM. This is a read-only field.
  */
  VmUuid *string `json:"vmUuid,omitempty"`
}


func NewNvm() *Nvm {
  p := new(Nvm)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.Nvm"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.Nvm"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM add information.
*/
type NvmAddInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  CVM IPV4 address list of the cluster.
  */
  CvmIps []string `json:"cvmIps,omitempty"`
  /**
  DNS domain name for the file server.
  */
  DnsDomainName *string `json:"dnsDomainName,omitempty"`
  /**
  DNS IPV4 address list for file server.
  */
  DnsIpv4AddressList []string `json:"dnsIpv4AddressList,omitempty"`
  /**
  NTP IPV4 address list for file server.
  */
  NtpIpv4AddressList []string `json:"ntpIpv4AddressList,omitempty"`
  /**
  List of VMs to be added.
  */
  NvmAddList []import1.VM `json:"nvmAddList,omitempty"`
  /**
  File server size in Gb.
  */
  SizeInGb *float64 `json:"sizeInGb,omitempty"`
}


func NewNvmAddInfo() *NvmAddInfo {
  p := new(NvmAddInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NvmAddInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NvmAddInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM power off information.
*/
type NvmPowerOffInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of VM UUIDs to be powered off.
  */
  NvmPowerOffList []string `json:"nvmPowerOffList,omitempty"`
}


func NewNvmPowerOffInfo() *NvmPowerOffInfo {
  p := new(NvmPowerOffInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NvmPowerOffInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NvmPowerOffInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM power-on information.
*/
type NvmPowerOnInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of VM UUIDs to be powered on.
  */
  NvmPowerOnList []string `json:"nvmPowerOnList,omitempty"`
}


func NewNvmPowerOnInfo() *NvmPowerOnInfo {
  p := new(NvmPowerOnInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NvmPowerOnInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NvmPowerOnInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM remove information.
*/
type NvmRemoveInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of VM UUIDs to be removed.
  */
  NvmRemoveList []string `json:"nvmRemoveList,omitempty"`
}


func NewNvmRemoveInfo() *NvmRemoveInfo {
  p := new(NvmRemoveInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NvmRemoveInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NvmRemoveInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the NVM update information.
*/
type NvmUpdateInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of VMs to be updated.
  */
  NvmUpdateList []import1.VM `json:"nvmUpdateList,omitempty"`
}


func NewNvmUpdateInfo() *NvmUpdateInfo {
  p := new(NvmUpdateInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.NvmUpdateInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.NvmUpdateInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Platform type.The acceptable values are \"kPE\",\"kPC\", \"kXi\", \"kMinerva\", \"kOther\".
*/
type OwnerType int

const(
  OWNERTYPE_UNKNOWN OwnerType = 0
  OWNERTYPE_REDACTED OwnerType = 1
  OWNERTYPE_KPE OwnerType = 2
  OWNERTYPE_KPC OwnerType = 3
  OWNERTYPE_KXI OwnerType = 4
  OWNERTYPE_KMINERVA OwnerType = 5
  OWNERTYPE_KOTHER OwnerType = 6
)

// returns the name of the enum given an ordinal number
func (e *OwnerType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kPE",
    "kPC",
    "kXi",
    "kMinerva",
    "kOther",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *OwnerType) index(name string) OwnerType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kPE",
    "kPC",
    "kXi",
    "kMinerva",
    "kOther",
  }
  for idx := range names {
    if names[idx] == name {
      return OwnerType(idx)
    }
  }
  return OWNERTYPE_UNKNOWN
}

func (e *OwnerType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for OwnerType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *OwnerType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e OwnerType) Ref() *OwnerType {
  return &e
}


/**
Prism central configuration model.
*/
type PcConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Files manager version.
  */
  FmVersion *string `json:"fmVersion,omitempty"`
  /**
  Prism central node IP address list.
  */
  NodeIps []import5.IPv4Address `json:"nodeIps,omitempty"`
  /**
  Password for Prism central API authentications.
  */
  PcApiPassword *string `json:"pcApiPassword,omitempty"`
  /**
  Username for Prism central API authentications.
  */
  PcApiUser *string `json:"pcApiUser,omitempty"`
  /**
  The Prism central UUID.
  */
  PcClusterExtId *string `json:"pcClusterExtId,omitempty"`
  /**
  Prism central version.
  */
  PcVersion *string `json:"pcVersion,omitempty"`
  /**
  Prism central SSH public keys.
  */
  SshPublicKeys []string `json:"sshPublicKeys,omitempty"`
  
  VirtualIpAddress *import5.IPv4Address `json:"virtualIpAddress,omitempty"`
}


func NewPcConfig() *PcConfig {
  p := new(PcConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PcConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PcConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
PE config
*/
type PeConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Prism central node IP address list.
  */
  CvmIps []import5.IPv4Address `json:"cvmIps,omitempty"`
  
  DataserviceIp *import5.IPv4Address `json:"dataserviceIp,omitempty"`
}


func NewPeConfig() *PeConfig {
  p := new(PeConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PeConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PeConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Level of applying personal identification information scrubbing on pulse data. The acceptable values are \"kAll\", \"kPartial\", \"kNone\", \"kAuto\".
*/
type PiiScrubbingLevel int

const(
  PIISCRUBBINGLEVEL_UNKNOWN PiiScrubbingLevel = 0
  PIISCRUBBINGLEVEL_REDACTED PiiScrubbingLevel = 1
  PIISCRUBBINGLEVEL_KALL PiiScrubbingLevel = 2
  PIISCRUBBINGLEVEL_KPARTIAL PiiScrubbingLevel = 3
  PIISCRUBBINGLEVEL_KNONE PiiScrubbingLevel = 4
  PIISCRUBBINGLEVEL_KAUTO PiiScrubbingLevel = 5
)

// returns the name of the enum given an ordinal number
func (e *PiiScrubbingLevel) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kAll",
    "kPartial",
    "kNone",
    "kAuto",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *PiiScrubbingLevel) index(name string) PiiScrubbingLevel {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kAll",
    "kPartial",
    "kNone",
    "kAuto",
  }
  for idx := range names {
    if names[idx] == name {
      return PiiScrubbingLevel(idx)
    }
  }
  return PIISCRUBBINGLEVEL_UNKNOWN
}

func (e *PiiScrubbingLevel) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for PiiScrubbingLevel:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *PiiScrubbingLevel) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e PiiScrubbingLevel) Ref() *PiiScrubbingLevel {
  return &e
}


/**
A model that represents the platform files information for request and response.
*/
type PlatformFiles struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  /**
  Number of generate UUID requests.
  */
  GenerateNumUuids *int `json:"generateNumUuids,omitempty"`
  /**
  List of generated UUIDs.
  */
  GeneratedUuids []string `json:"generatedUuids,omitempty"`
  
  NvmInfo *PlatformFilesNvmInfo `json:"nvmInfo,omitempty"`
  
  PlatformInfo *PlatformInfo `json:"platformInfo,omitempty"`
}


func NewPlatformFiles() *PlatformFiles {
  p := new(PlatformFiles)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFiles"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFiles"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/platform-files Post operation
*/
type PlatformFilesGetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlatformFilesGetApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlatformFilesGetApiResponse() *PlatformFilesGetApiResponse {
  p := new(PlatformFilesGetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFilesGetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFilesGetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlatformFilesGetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlatformFilesGetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlatformFilesGetApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
A model that represents the platform files NVM information.
*/
type PlatformFilesNvmInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Get request flag for NVM information.
  */
  IsNvmInfo *bool `json:"isNvmInfo,omitempty"`
  /**
  List of NVM objects.
  */
  NvmList []import1.VM `json:"nvmList,omitempty"`
}


func NewPlatformFilesNvmInfo() *PlatformFilesNvmInfo {
  p := new(PlatformFilesNvmInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFilesNvmInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFilesNvmInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the platform files update request.
*/
type PlatformFilesUpdate struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The cluster associated with the file server. This contains AOS cluster UUID. This is a read-only field.
  */
  ClusterUuid *string `json:"clusterUuid,omitempty"`
  /**
  The {extId} of the file server. This is a read-only field.
  */
  FileServerExtId *string `json:"fileServerExtId,omitempty"`
  
  NvmAddInfo *NvmAddInfo `json:"nvmAddInfo,omitempty"`
  
  NvmPowerOffInfo *NvmPowerOffInfo `json:"nvmPowerOffInfo,omitempty"`
  
  NvmPowerOnInfo *NvmPowerOnInfo `json:"nvmPowerOnInfo,omitempty"`
  
  NvmRemoveInfo *NvmRemoveInfo `json:"nvmRemoveInfo,omitempty"`
  
  NvmUpdateInfo *NvmUpdateInfo `json:"nvmUpdateInfo,omitempty"`
  /**
  ExtId of the task.
  */
  TaskUuid *string `json:"taskUuid,omitempty"`
}


func NewPlatformFilesUpdate() *PlatformFilesUpdate {
  p := new(PlatformFilesUpdate)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFilesUpdate"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFilesUpdate"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/platform-files/$actions/update Post operation
*/
type PlatformFilesUpdateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlatformFilesUpdateApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlatformFilesUpdateApiResponse() *PlatformFilesUpdateApiResponse {
  p := new(PlatformFilesUpdateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFilesUpdateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFilesUpdateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlatformFilesUpdateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlatformFilesUpdateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlatformFilesUpdateApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
A model that represents the platform files virtual network information along with available IPV4 addresses.
*/
type PlatformFilesVirtualNetworkInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of available IPV4 addresses to allocate from the virtual network address pool.
  */
  AvailableIPv4AddressList []string `json:"availableIPv4AddressList,omitempty"`
  
  DefaultGateway *import5.IPAddress `json:"defaultGateway,omitempty"`
  /**
  UUID of the the file server network.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Name of the network.
  */
  Name *string `json:"name,omitempty"`
  
  SubnetMask *import5.IPAddress `json:"subnetMask,omitempty"`
}


func NewPlatformFilesVirtualNetworkInfo() *PlatformFilesVirtualNetworkInfo {
  p := new(PlatformFilesVirtualNetworkInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformFilesVirtualNetworkInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformFilesVirtualNetworkInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents the platform and files storage provider information.
*/
type PlatformInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Get request flag for platform network information of the files manager.
  */
  IsNetworkInfo *bool `json:"isNetworkInfo,omitempty"`
  /**
  Get request flag for files storage provider information.
  */
  IsStorageProviderInfo *bool `json:"isStorageProviderInfo,omitempty"`
  /**
  Get request flag for vCenter connection status.
  */
  IsVcenterStatus *bool `json:"isVcenterStatus,omitempty"`
  /**
  Get request flag for files manager and the corresponding platform version.
  */
  IsVersionInfo *bool `json:"isVersionInfo,omitempty"`
  /**
  Get request flag for virtual network information.
  */
  IsVirtualNetworkInfo *bool `json:"isVirtualNetworkInfo,omitempty"`
  /**
  Platform FQDN for the files manager.
  */
  ManagerFqdn *string `json:"managerFqdn,omitempty"`
  /**
  Platform cluster node IPs for the files manager.
  */
  ManagerNodeIps []import5.IPAddress `json:"managerNodeIps,omitempty"`
  /**
  Platform version for the files manager.
  */
  ManagerPlatformVersion *string `json:"managerPlatformVersion,omitempty"`
  /**
  Files manager version.
  */
  ManagerVersion *string `json:"managerVersion,omitempty"`
  
  ManagerVirtualIpAddress *import5.IPAddress `json:"managerVirtualIpAddress,omitempty"`
  
  StorageProviderDsIp *import5.IPAddress `json:"storageProviderDsIp,omitempty"`
  /**
  Files storage provider cluster node IPs.
  */
  StorageProviderNodeIps []import5.IPAddress `json:"storageProviderNodeIps,omitempty"`
  /**
  Flag to represent if the vCenter is connected to the cluster.
  */
  VCenterConnected *bool `json:"vCenterConnected,omitempty"`
  /**
  List of virtual network information.
  */
  VirtualNetworkInformation []PlatformFilesVirtualNetworkInfo `json:"virtualNetworkInformation,omitempty"`
  /**
  List of virtual network UUIDs.
  */
  VirtualNetworkUuids []string `json:"virtualNetworkUuids,omitempty"`
}


func NewPlatformInfo() *PlatformInfo {
  p := new(PlatformInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-service/$actions/migrate-platform Post operation
*/
type PlatformMigrateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlatformMigrateApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlatformMigrateApiResponse() *PlatformMigrateApiResponse {
  p := new(PlatformMigrateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformMigrateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformMigrateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlatformMigrateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlatformMigrateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlatformMigrateApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/files-service/$actions/notify-platform Post operation
*/
type PlatformNotifyApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfPlatformNotifyApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewPlatformNotifyApiResponse() *PlatformNotifyApiResponse {
  p := new(PlatformNotifyApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PlatformNotifyApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PlatformNotifyApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *PlatformNotifyApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *PlatformNotifyApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfPlatformNotifyApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Pre upgrade operation specification with the target software details for the NVM upgrade.
*/
type PreUpgradeNvm struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Full release version string of the software.
  */
  FullReleaseVersion *string `json:"fullReleaseVersion,omitempty"`
  /**
  Run pre-upgrade checks alone.
  */
  PreUpgradeChecksOnly *bool `json:"preUpgradeChecksOnly,omitempty"`
  /**
  Software version.
  */
  Version *string `json:"version,omitempty"`
}


func NewPreUpgradeNvm() *PreUpgradeNvm {
  p := new(PreUpgradeNvm)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PreUpgradeNvm"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PreUpgradeNvm"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Proxy type. The acceptable values are \"kHttp\", \"kHttps\", \"kSocks\".
*/
type ProxyType int

const(
  PROXYTYPE_UNKNOWN ProxyType = 0
  PROXYTYPE_REDACTED ProxyType = 1
  PROXYTYPE_KHTTP ProxyType = 2
  PROXYTYPE_KHTTPS ProxyType = 3
  PROXYTYPE_KSOCKS ProxyType = 4
)

// returns the name of the enum given an ordinal number
func (e *ProxyType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kHttp",
    "kHttps",
    "kSocks",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ProxyType) index(name string) ProxyType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kHttp",
    "kHttps",
    "kSocks",
  }
  for idx := range names {
    if names[idx] == name {
      return ProxyType(idx)
    }
  }
  return PROXYTYPE_UNKNOWN
}

func (e *ProxyType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ProxyType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ProxyType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ProxyType) Ref() *ProxyType {
  return &e
}


/**
Pulse configuration model.
*/
type PulseConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of http proxy servers for pulse data.
  */
  HttpProxyList []HttpProxy `json:"httpProxyList,omitempty"`
  
  PiiScrubbingLevel *PiiScrubbingLevel `json:"piiScrubbingLevel,omitempty"`
  /**
  Enable pulse on file server.
  */
  PulseEnabled *bool `json:"pulseEnabled,omitempty"`
}


func NewPulseConfig() *PulseConfig {
  p := new(PulseConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.PulseConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.PulseConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Source type of certificate. The acceptable values are \"kPC\", \"kPE\"
 \"kGenesis\", \"kManual\".
*/
type SourceType int

const(
  SOURCETYPE_UNKNOWN SourceType = 0
  SOURCETYPE_REDACTED SourceType = 1
  SOURCETYPE_KPC SourceType = 2
  SOURCETYPE_KPE SourceType = 3
  SOURCETYPE_KGENESIS SourceType = 4
  SOURCETYPE_KMANUAL SourceType = 5
)

// returns the name of the enum given an ordinal number
func (e *SourceType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kPC",
    "kPE",
    "kGenesis",
    "kManual",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SourceType) index(name string) SourceType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "kPC",
    "kPE",
    "kGenesis",
    "kManual",
  }
  for idx := range names {
    if names[idx] == name {
      return SourceType(idx)
    }
  }
  return SOURCETYPE_UNKNOWN
}

func (e *SourceType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SourceType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SourceType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SourceType) Ref() *SourceType {
  return &e
}


/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/storage-specs Post operation
*/
type StorageCreateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfStorageCreateApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewStorageCreateApiResponse() *StorageCreateApiResponse {
  p := new(StorageCreateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.StorageCreateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.StorageCreateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *StorageCreateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *StorageCreateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfStorageCreateApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/storage-specs Delete operation
*/
type StorageDeleteApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfStorageDeleteApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewStorageDeleteApiResponse() *StorageDeleteApiResponse {
  p := new(StorageDeleteApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.StorageDeleteApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.StorageDeleteApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *StorageDeleteApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *StorageDeleteApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfStorageDeleteApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/storage-specs Get operation
*/
type StorageGetApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfStorageGetApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewStorageGetApiResponse() *StorageGetApiResponse {
  p := new(StorageGetApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.StorageGetApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.StorageGetApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *StorageGetApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *StorageGetApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfStorageGetApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
Storage information model.
*/
type StorageInfo struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Flag that represents whether the storage attachments are healthy.
  */
  AreAttachmentsHealthy *bool `json:"areAttachmentsHealthy,omitempty"`
  /**
  Get request flag for storage information including volume group list.
  */
  IsStorageInfo *bool `json:"isStorageInfo,omitempty"`
  /**
  Get request flag for storage health verification of ISCSI or hypervisor attachments.
  */
  IsVerifyAttachmentsHealth *bool `json:"isVerifyAttachmentsHealth,omitempty"`
  /**
  List of volume groups set UUIDs for the file server.
  */
  VgSetList []string `json:"vgSetList,omitempty"`
}


func NewStorageInfo() *StorageInfo {
  p := new(StorageInfo)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.StorageInfo"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.StorageInfo"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /files/v4.0.a2/infra/file-server/storage-specs Patch operation
*/
type StorageUpdateApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfStorageUpdateApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewStorageUpdateApiResponse() *StorageUpdateApiResponse {
  p := new(StorageUpdateApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.StorageUpdateApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.StorageUpdateApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *StorageUpdateApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *StorageUpdateApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfStorageUpdateApiResponseData()
  }
  e := p.Data.SetValue(v)
  if nil == e {
    if nil == p.DataItemDiscriminator_ {
      p.DataItemDiscriminator_ = new(string)
    }
    *p.DataItemDiscriminator_ = *p.Data.Discriminator
  }
  return e
}



/**
TLS certificate configuration for mutual TLS communications between services.
*/
type TlsCertificateConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  CaChain *CaChain `json:"caChain,omitempty"`
  /**
  File server ICA certificate signed by platform ICA.
  */
  SignedCertificate *string `json:"signedCertificate,omitempty"`
  
  SourceType *SourceType `json:"sourceType,omitempty"`
}


func NewTlsCertificateConfig() *TlsCertificateConfig {
  p := new(TlsCertificateConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "files.v4.infra.TlsCertificateConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "files.v4.r0.a2.infra.TlsCertificateConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfStorageCreateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType1 []import5.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfStorageCreateApiResponseData() *OneOfStorageCreateApiResponseData {
  p := new(OneOfStorageCreateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfStorageCreateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfStorageCreateApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case []import5.Message:
      p.oneOfType1 = v.([]import5.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfStorageCreateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfStorageCreateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new([]import5.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfStorageCreateApiResponseData"))
}

func (p *OneOfStorageCreateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfStorageCreateApiResponseData")
}

type OneOfStorageGetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType1 []import5.Message `json:"-"`
  oneOfType0 *FileServerStorageSpec `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfStorageGetApiResponseData() *OneOfStorageGetApiResponseData {
  p := new(OneOfStorageGetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfStorageGetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfStorageGetApiResponseData is nil"))
  }
  switch v.(type) {
    case []import5.Message:
      p.oneOfType1 = v.([]import5.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case FileServerStorageSpec:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerStorageSpec)}
      *p.oneOfType0 = v.(FileServerStorageSpec)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfStorageGetApiResponseData) GetValue() interface{} {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfStorageGetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType1 := new([]import5.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType0 := new(FileServerStorageSpec)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.FileServerStorageSpec" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FileServerStorageSpec)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfStorageGetApiResponseData"))
}

func (p *OneOfStorageGetApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfStorageGetApiResponseData")
}

type OneOfListFilesInfraLockInfoApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FilesInfraLockList `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfListFilesInfraLockInfoApiResponseData() *OneOfListFilesInfraLockInfoApiResponseData {
  p := new(OneOfListFilesInfraLockInfoApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfListFilesInfraLockInfoApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfListFilesInfraLockInfoApiResponseData is nil"))
  }
  switch v.(type) {
    case FilesInfraLockList:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockList)}
      *p.oneOfType0 = v.(FilesInfraLockList)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfListFilesInfraLockInfoApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfListFilesInfraLockInfoApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FilesInfraLockList)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.FilesInfraLockList" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockList)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListFilesInfraLockInfoApiResponseData"))
}

func (p *OneOfListFilesInfraLockInfoApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfListFilesInfraLockInfoApiResponseData")
}

type OneOfFilesServiceInfoGetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *FilesService `json:"-"`
}

func NewOneOfFilesServiceInfoGetApiResponseData() *OneOfFilesServiceInfoGetApiResponseData {
  p := new(OneOfFilesServiceInfoGetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFilesServiceInfoGetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFilesServiceInfoGetApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case FilesService:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FilesService)}
      *p.oneOfType0 = v.(FilesService)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFilesServiceInfoGetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfFilesServiceInfoGetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(FilesService)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.FilesService" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FilesService)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFilesServiceInfoGetApiResponseData"))
}

func (p *OneOfFilesServiceInfoGetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfFilesServiceInfoGetApiResponseData")
}

type OneOfPlatformNotifyApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfPlatformNotifyApiResponseData() *OneOfPlatformNotifyApiResponseData {
  p := new(OneOfPlatformNotifyApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlatformNotifyApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlatformNotifyApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlatformNotifyApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfPlatformNotifyApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlatformNotifyApiResponseData"))
}

func (p *OneOfPlatformNotifyApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfPlatformNotifyApiResponseData")
}

type OneOfFilesInfraLockAcquireApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *FilesInfraLockAcquireResponse `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfFilesInfraLockAcquireApiResponseData() *OneOfFilesInfraLockAcquireApiResponseData {
  p := new(OneOfFilesInfraLockAcquireApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFilesInfraLockAcquireApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFilesInfraLockAcquireApiResponseData is nil"))
  }
  switch v.(type) {
    case FilesInfraLockAcquireResponse:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockAcquireResponse)}
      *p.oneOfType0 = v.(FilesInfraLockAcquireResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFilesInfraLockAcquireApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfFilesInfraLockAcquireApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(FilesInfraLockAcquireResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.FilesInfraLockAcquireResponse" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockAcquireResponse)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFilesInfraLockAcquireApiResponseData"))
}

func (p *OneOfFilesInfraLockAcquireApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfFilesInfraLockAcquireApiResponseData")
}

type OneOfPlatformFilesUpdateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfPlatformFilesUpdateApiResponseData() *OneOfPlatformFilesUpdateApiResponseData {
  p := new(OneOfPlatformFilesUpdateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlatformFilesUpdateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlatformFilesUpdateApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlatformFilesUpdateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfPlatformFilesUpdateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlatformFilesUpdateApiResponseData"))
}

func (p *OneOfPlatformFilesUpdateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfPlatformFilesUpdateApiResponseData")
}

type OneOfFilesInfraLockReleaseApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *FilesInfraLockReleaseResponse `json:"-"`
}

func NewOneOfFilesInfraLockReleaseApiResponseData() *OneOfFilesInfraLockReleaseApiResponseData {
  p := new(OneOfFilesInfraLockReleaseApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfFilesInfraLockReleaseApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfFilesInfraLockReleaseApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case FilesInfraLockReleaseResponse:
      if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockReleaseResponse)}
      *p.oneOfType0 = v.(FilesInfraLockReleaseResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfFilesInfraLockReleaseApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfFilesInfraLockReleaseApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(FilesInfraLockReleaseResponse)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.FilesInfraLockReleaseResponse" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(FilesInfraLockReleaseResponse)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFilesInfraLockReleaseApiResponseData"))
}

func (p *OneOfFilesInfraLockReleaseApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfFilesInfraLockReleaseApiResponseData")
}

type OneOfPlatformFilesGetApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
  oneOfType0 *PlatformFiles `json:"-"`
}

func NewOneOfPlatformFilesGetApiResponseData() *OneOfPlatformFilesGetApiResponseData {
  p := new(OneOfPlatformFilesGetApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlatformFilesGetApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlatformFilesGetApiResponseData is nil"))
  }
  switch v.(type) {
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    case PlatformFiles:
      if nil == p.oneOfType0 {p.oneOfType0 = new(PlatformFiles)}
      *p.oneOfType0 = v.(PlatformFiles)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlatformFilesGetApiResponseData) GetValue() interface{} {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  return nil
}

func (p *OneOfPlatformFilesGetApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  vOneOfType0 := new(PlatformFiles)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "files.v4.infra.PlatformFiles" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(PlatformFiles)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlatformFilesGetApiResponseData"))
}

func (p *OneOfPlatformFilesGetApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  return nil, errors.New("No value to marshal for OneOfPlatformFilesGetApiResponseData")
}

type OneOfStorageDeleteApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType1 []import5.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfStorageDeleteApiResponseData() *OneOfStorageDeleteApiResponseData {
  p := new(OneOfStorageDeleteApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfStorageDeleteApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfStorageDeleteApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case []import5.Message:
      p.oneOfType1 = v.([]import5.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfStorageDeleteApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfStorageDeleteApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new([]import5.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfStorageDeleteApiResponseData"))
}

func (p *OneOfStorageDeleteApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfStorageDeleteApiResponseData")
}

type OneOfStorageUpdateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType1 []import5.Message `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfStorageUpdateApiResponseData() *OneOfStorageUpdateApiResponseData {
  p := new(OneOfStorageUpdateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfStorageUpdateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfStorageUpdateApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case []import5.Message:
      p.oneOfType1 = v.([]import5.Message)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfStorageUpdateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return p.oneOfType1
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfStorageUpdateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType1 := new([]import5.Message)
  if err := json.Unmarshal(b, vOneOfType1); err == nil {
    if len(*vOneOfType1) == 0 || "common.v1.config.Message" == *((*vOneOfType1)[0].ObjectType_) {
      p.oneOfType1 = *vOneOfType1
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<common.v1.config.Message>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<common.v1.config.Message>"
      return nil

    }
  }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfStorageUpdateApiResponseData"))
}

func (p *OneOfStorageUpdateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if "List<common.v1.config.Message>" == *p.Discriminator {
    return json.Marshal(p.oneOfType1)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfStorageUpdateApiResponseData")
}

type OneOfPlatformMigrateApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *import6.TaskReference `json:"-"`
  oneOfType400 *import3.ErrorResponse `json:"-"`
}

func NewOneOfPlatformMigrateApiResponseData() *OneOfPlatformMigrateApiResponseData {
  p := new(OneOfPlatformMigrateApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfPlatformMigrateApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfPlatformMigrateApiResponseData is nil"))
  }
  switch v.(type) {
    case import6.TaskReference:
      if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = v.(import6.TaskReference)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import3.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = v.(import3.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfPlatformMigrateApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfPlatformMigrateApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(import6.TaskReference)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(import6.TaskReference)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import3.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import3.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlatformMigrateApiResponseData"))
}

func (p *OneOfPlatformMigrateApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfPlatformMigrateApiResponseData")
}

