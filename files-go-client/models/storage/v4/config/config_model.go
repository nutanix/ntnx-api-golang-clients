/*
 * Generated file models/storage/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Files Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module storage.v4.config of Nutanix Files Versioned APIs
*/
package config
import (
  "bytes"
  import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/config"
  import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
)

/**
The type of authentication enabled for the Volume Group. If this is set to CHAP, the target/client secret must be provided.
*/
type AuthenticationType int

const(
  AUTHENTICATIONTYPE_UNKNOWN AuthenticationType = 0
  AUTHENTICATIONTYPE_REDACTED AuthenticationType = 1
  AUTHENTICATIONTYPE_CHAP AuthenticationType = 2
  AUTHENTICATIONTYPE_NONE AuthenticationType = 3
)

// returns the name of the enum given an ordinal number
func (e *AuthenticationType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CHAP",
    "NONE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AuthenticationType) index(name string) AuthenticationType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CHAP",
    "NONE",
  }
  for idx := range names {
    if names[idx] == name {
      return AuthenticationType(idx)
    }
  }
  return AUTHENTICATIONTYPE_UNKNOWN
}

func (e *AuthenticationType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AuthenticationType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AuthenticationType) Ref() *AuthenticationType {
  return &e
}


/**
Storage features for the Volume disks.
*/
type DiskStorageFeatures struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FlashMode *FlashMode `json:"flashMode,omitempty"`
}


func NewDiskStorageFeatures() *DiskStorageFeatures {
  p := new(DiskStorageFeatures)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.DiskStorageFeatures"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.DiskStorageFeatures"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Performance setting to avoid down migration of data from the hot tier. When specified, all the virtual disks of the volume group are pinned to the higher tier, unless overrides are specified for some of the virtual disks.
*/
type FlashMode struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IsEnabled *bool `json:"isEnabled,omitempty"`
}


func NewFlashMode() *FlashMode {
  p := new(FlashMode)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.FlashMode"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.FlashMode"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents iSCSI Client that can be associated with a volume group as an external attachment.
*/
type IscsiClient struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  iSCSI initiator Client Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  ClientSecret *string `json:"clientSecret,omitempty"`
  /**
  The UUID of cluster that will host the Volume Group. This is mandatory to be specified for creating a volume group on PC.
  */
  ClusterReference *string `json:"clusterReference,omitempty"`
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  iSCSI Initiator Name.
  */
  IscsiInitiatorName *string `json:"iscsiInitiatorName,omitempty"`
  
  IscsiInitiatorNetworkId *import1.IPAddressOrFQDN `json:"iscsiInitiatorNetworkId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  
  TargetParams []TargetParam `json:"targetParams,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewIscsiClient() *IscsiClient {
  p := new(IscsiClient)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.IscsiClient"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.IscsiClient"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Whether the Volume Group can be shared across multiple iSCSI initiators. The mode cannot be changed from SHARED to NOT_SHARED on a volume group with multiple attachments. Similarly, a volume group cannot be associated with more than one attachment as long as it is in exclusive mode.
*/
type SharingStatus int

const(
  SHARINGSTATUS_UNKNOWN SharingStatus = 0
  SHARINGSTATUS_REDACTED SharingStatus = 1
  SHARINGSTATUS_SHARED SharingStatus = 2
  SHARINGSTATUS_NOT_SHARED SharingStatus = 3
)

// returns the name of the enum given an ordinal number
func (e *SharingStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARED",
    "NOT_SHARED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SharingStatus) index(name string) SharingStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "SHARED",
    "NOT_SHARED",
  }
  for idx := range names {
    if names[idx] == name {
      return SharingStatus(idx)
    }
  }
  return SHARINGSTATUS_UNKNOWN
}

func (e *SharingStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SharingStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SharingStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SharingStatus) Ref() *SharingStatus {
  return &e
}


/**
Storage features for the Volume Groups.
*/
type StorageFeatures struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  FlashMode *FlashMode `json:"flashMode,omitempty"`
}


func NewStorageFeatures() *StorageFeatures {
  p := new(StorageFeatures)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.StorageFeatures"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.StorageFeatures"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of iSCSI targets' parameters that will be visible and accessible to the iSCSI client.
*/
type TargetParam struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Number of virtual targets to generate for the iSCSI target.
  */
  NumVirtualTargets *int `json:"numVirtualTargets,omitempty"`
}


func NewTargetParam() *TargetParam {
  p := new(TargetParam)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.TargetParam"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.TargetParam"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents a VM reference that can be associated with a volume group as a hypervisor attachment.
*/
type VmAttachment struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  The external identifier of the VM.
  */
  ExtId *string `json:"extId"`
  /**
  The index on the scsi bus to attach the VM to the Volume Group.
  */
  Index *int `json:"index,omitempty"`
}

func (p *VmAttachment) MarshalJSON() ([]byte, error) {
  type VmAttachmentProxy VmAttachment
  return json.Marshal(struct {
    *VmAttachmentProxy
    ExtId *string `json:"extId,omitempty"`
  }{
    VmAttachmentProxy : (*VmAttachmentProxy)(p),
    ExtId : p.ExtId,
  })
}

func NewVmAttachment() *VmAttachment {
  p := new(VmAttachment)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VmAttachment"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.VmAttachment"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents volume disk which is associated with a volume group, and is supported by a backing file on DSF.
*/
type VolumeDisk struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Description of the Volume disk.
  */
  Description *string `json:"description,omitempty"`
  
  DiskDataSourceReference *import1.EntityReference `json:"diskDataSourceReference,omitempty"`
  /**
  Size of the disk in bytes.
  */
  DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`
  
  DiskStorageFeatures *DiskStorageFeatures `json:"diskStorageFeatures,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Index of the disk in a Volume Group. This field is immutable.
  */
  Index *int `json:"index,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Storage container on which the disk must be created.
  */
  StorageContainerId *string `json:"storageContainerId,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewVolumeDisk() *VolumeDisk {
  p := new(VolumeDisk)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeDisk"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.VolumeDisk"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
A model that represents Volume Group resources.
*/
type VolumeGroup struct {
  
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
  
  EnabledAuthentications *AuthenticationType `json:"enabledAuthentications,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
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
  
  SharingStatus *SharingStatus `json:"sharingStatus,omitempty"`
  
  StorageFeatures *StorageFeatures `json:"storageFeatures,omitempty"`
  /**
  Target Secret in case of CHAP authentication. This field should not be provided in case the authentication type is not set to CHAP.
  */
  TargetSecret *string `json:"targetSecret,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewVolumeGroup() *VolumeGroup {
  p := new(VolumeGroup)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "storage.v4.config.VolumeGroup"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "storage.v4.r0.a1.config.VolumeGroup"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



