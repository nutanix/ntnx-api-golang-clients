/*
 * Generated file models/lcm/v4/common/common_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Lcm Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module of common models
*/
package common
import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
)

/**
Collect metrics boolean attribute
*/
type CollectMetrics struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Boolean argument to collect metrics data or not
  */
  CollectMetrics *bool `json:"collectMetrics,omitempty"`
}


func NewCollectMetrics() *CollectMetrics {
  p := new(CollectMetrics)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.CollectMetrics"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.CollectMetrics"}
  p.UnknownFields_ = map[string]interface{}{}

  p.CollectMetrics = new(bool)
  *p.CollectMetrics = false


  return p
}




/**
A specification defining the entity being deployed and its version
*/
type EntityDeploySpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityIdentifierSpec *EntityIdentifierSpec `json:"entityIdentifierSpec,omitempty"`
}


func NewEntityDeploySpec() *EntityDeploySpec {
  p := new(EntityDeploySpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.EntityDeploySpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.EntityDeploySpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Unique identifier of an entity in LCM
*/
type EntityIdentifierSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass,omitempty"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel,omitempty"`
  /**
  A hardware family for a LCM entity
  */
  HardwareFamily *string `json:"hardwareFamily,omitempty"`
}


func NewEntityIdentifierSpec() *EntityIdentifierSpec {
  p := new(EntityIdentifierSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.EntityIdentifierSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.EntityIdentifierSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Specification for running an update operation
*/
type EntityUpdateSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  UUID of the entity to be used
  */
  EntityUuid *string `json:"entityUuid,omitempty"`
  /**
  Entity target version
  */
  Version *string `json:"version,omitempty"`
}


func NewEntityUpdateSpec() *EntityUpdateSpec {
  p := new(EntityUpdateSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.EntityUpdateSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.EntityUpdateSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of entity update objects for getting recommendations
*/
type EntityUpdateSpecs struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
}


func NewEntityUpdateSpecs() *EntityUpdateSpecs {
  p := new(EntityUpdateSpecs)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.EntityUpdateSpecs"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.EntityUpdateSpecs"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of hypervisor present in the cluster
*/
type HypervisorType int

const(
  HYPERVISORTYPE_UNKNOWN HypervisorType = 0
  HYPERVISORTYPE_REDACTED HypervisorType = 1
  HYPERVISORTYPE_ESX HypervisorType = 2
  HYPERVISORTYPE_AHV HypervisorType = 3
  HYPERVISORTYPE_HYPERV HypervisorType = 4
  HYPERVISORTYPE_XENSERVER HypervisorType = 5
  HYPERVISORTYPE_NONE HypervisorType = 6
)

// returns the name of the enum given an ordinal number
func (e *HypervisorType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "esx",
    "ahv",
    "hyperv",
    "xenserver",
    "none",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *HypervisorType) index(name string) HypervisorType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "esx",
    "ahv",
    "hyperv",
    "xenserver",
    "none",
  }
  for idx := range names {
    if names[idx] == name {
      return HypervisorType(idx)
    }
  }
  return HYPERVISORTYPE_UNKNOWN
}

func (e *HypervisorType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for HypervisorType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *HypervisorType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e HypervisorType) Ref() *HypervisorType {
  return &e
}


/**
LCM image download specification
*/
type ImageDownloadSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM entity class
  */
  EntityClass *string `json:"entityClass"`
  /**
  LCM entity model
  */
  EntityModel *string `json:"entityModel"`
  /**
  UUID of the entity to be used
  */
  EntityUuid *string `json:"entityUuid,omitempty"`
  /**
  A hardware family for a LCM entity
  */
  HardwareFamily *string `json:"hardwareFamily,omitempty"`
  /**
  The location of the LCM entity, it can be at the cluster level or at the node level, in the format of "cluster:uuid" or "node:uuid"
  */
  LocationId *string `json:"locationId,omitempty"`
  /**
  The requested update version of an LCM entity.
  */
  Version *string `json:"version"`
}

func (p *ImageDownloadSpec) MarshalJSON() ([]byte, error) {
  type ImageDownloadSpecProxy ImageDownloadSpec
  return json.Marshal(struct {
    *ImageDownloadSpecProxy
    EntityClass *string `json:"entityClass,omitempty"`
    EntityModel *string `json:"entityModel,omitempty"`
    Version *string `json:"version,omitempty"`
  }{
    ImageDownloadSpecProxy : (*ImageDownloadSpecProxy)(p),
    EntityClass : p.EntityClass,
    EntityModel : p.EntityModel,
    Version : p.Version,
  })
}

func NewImageDownloadSpec() *ImageDownloadSpec {
  p := new(ImageDownloadSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.ImageDownloadSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.ImageDownloadSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
List of LCM images to download
*/
type ImageDownloadSpecs struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  LCM image download specification
  */
  ImageDownloadSpecs []ImageDownloadSpec `json:"imageDownloadSpecs,omitempty"`
}


func NewImageDownloadSpecs() *ImageDownloadSpecs {
  p := new(ImageDownloadSpecs)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.ImageDownloadSpecs"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.ImageDownloadSpecs"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Available version types
*/
type LcmAvailableVersionStatus int

const(
  LCMAVAILABLEVERSIONSTATUS_UNKNOWN LcmAvailableVersionStatus = 0
  LCMAVAILABLEVERSIONSTATUS_REDACTED LcmAvailableVersionStatus = 1
  LCMAVAILABLEVERSIONSTATUS_RECOMMENDED LcmAvailableVersionStatus = 2
  LCMAVAILABLEVERSIONSTATUS_CRITICAL LcmAvailableVersionStatus = 3
  LCMAVAILABLEVERSIONSTATUS_LATEST LcmAvailableVersionStatus = 4
  LCMAVAILABLEVERSIONSTATUS_DEPRECATED LcmAvailableVersionStatus = 5
  LCMAVAILABLEVERSIONSTATUS_EMERGENCY LcmAvailableVersionStatus = 6
  LCMAVAILABLEVERSIONSTATUS_AVAILABLE LcmAvailableVersionStatus = 7
  LCMAVAILABLEVERSIONSTATUS_LTS LcmAvailableVersionStatus = 8
  LCMAVAILABLEVERSIONSTATUS_STS LcmAvailableVersionStatus = 9
)

// returns the name of the enum given an ordinal number
func (e *LcmAvailableVersionStatus) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "recommended",
    "critical",
    "latest",
    "deprecated",
    "emergency",
    "available",
    "LTS",
    "STS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *LcmAvailableVersionStatus) index(name string) LcmAvailableVersionStatus {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "recommended",
    "critical",
    "latest",
    "deprecated",
    "emergency",
    "available",
    "LTS",
    "STS",
  }
  for idx := range names {
    if names[idx] == name {
      return LcmAvailableVersionStatus(idx)
    }
  }
  return LCMAVAILABLEVERSIONSTATUS_UNKNOWN
}

func (e *LcmAvailableVersionStatus) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for LcmAvailableVersionStatus:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *LcmAvailableVersionStatus) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e LcmAvailableVersionStatus) Ref() *LcmAvailableVersionStatus {
  return &e
}


/**
Type of an LCM entity
*/
type LcmEntityType int

const(
  LCMENTITYTYPE_UNKNOWN LcmEntityType = 0
  LCMENTITYTYPE_REDACTED LcmEntityType = 1
  LCMENTITYTYPE_SOFTWARE LcmEntityType = 2
  LCMENTITYTYPE_FIRMWARE LcmEntityType = 3
)

// returns the name of the enum given an ordinal number
func (e *LcmEntityType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "software",
    "firmware",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *LcmEntityType) index(name string) LcmEntityType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "software",
    "firmware",
  }
  for idx := range names {
    if names[idx] == name {
      return LcmEntityType(idx)
    }
  }
  return LCMENTITYTYPE_UNKNOWN
}

func (e *LcmEntityType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for LcmEntityType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *LcmEntityType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e LcmEntityType) Ref() *LcmEntityType {
  return &e
}


/**
Cluster management server configuration used while updating clusters with ESX or HyperV
*/
type ManagementServer struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
  /**
  IP address of the management server
  */
  Ip *string `json:"ip,omitempty"`
  /**
  Password to login to the management server
  */
  Password *string `json:"password,omitempty"`
  /**
  User name to login to the management server
  */
  Username *string `json:"username,omitempty"`
}


func NewManagementServer() *ManagementServer {
  p := new(ManagementServer)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.ManagementServer"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.ManagementServer"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Specification for running a Precheck operation
*/
type PrecheckSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
  
  ManagementServer *ManagementServer `json:"managementServer,omitempty"`
  /**
  List of prechecks to skip. The allowed value is 'power_off_vm' that skips the pinned VM prechecks
  */
  SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
}


func NewPrecheckSpec() *PrecheckSpec {
  p := new(PrecheckSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.PrecheckSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.PrecheckSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
System auto-management flag
*/
type SystemAutoMgmtFlag int

const(
  SYSTEMAUTOMGMTFLAG_UNKNOWN SystemAutoMgmtFlag = 0
  SYSTEMAUTOMGMTFLAG_REDACTED SystemAutoMgmtFlag = 1
  SYSTEMAUTOMGMTFLAG_POWEROFFUVMS SystemAutoMgmtFlag = 2
)

// returns the name of the enum given an ordinal number
func (e *SystemAutoMgmtFlag) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "powerOffUvms",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SystemAutoMgmtFlag) index(name string) SystemAutoMgmtFlag {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "powerOffUvms",
  }
  for idx := range names {
    if names[idx] == name {
      return SystemAutoMgmtFlag(idx)
    }
  }
  return SYSTEMAUTOMGMTFLAG_UNKNOWN
}

func (e *SystemAutoMgmtFlag) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SystemAutoMgmtFlag:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SystemAutoMgmtFlag) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SystemAutoMgmtFlag) Ref() *SystemAutoMgmtFlag {
  return &e
}


/**
Specification for running an update operation
*/
type UpdateSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of automated system operations to perform, to avoid precheck failure and let the system restore state after an update is complete. The allowed flag is: - 'power_off_vm': This allows the system to automatically power off user VMs which cannot be migrated to other hosts and power them on when the update is done. This option can avoid pinned VM precheck failure on the host which needs to enter maintenance mode during the update and allow the update to go through.
  */
  AutoHandleFlags []SystemAutoMgmtFlag `json:"autoHandleFlags,omitempty"`
  
  EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
  
  ManagementServer *ManagementServer `json:"managementServer,omitempty"`
  /**
  List of prechecks to skip. The allowed value is 'power_off_vm' that skips the pinned VM prechecks
  */
  SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
  /**
  Number of seconds LCM waits for the VMs to come up after exiting host maintenance mode
  */
  WaitInSecForAppUp *int `json:"waitInSecForAppUp,omitempty"`
}


func NewUpdateSpec() *UpdateSpec {
  p := new(UpdateSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.UpdateSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.UpdateSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Identifies a version of an entity type in LCM
*/
type VersionIdentifierSpec struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  EntityIdentifierSpec *EntityIdentifierSpec `json:"entityIdentifierSpec,omitempty"`
  /**
  Entity target version
  */
  Version *string `json:"version,omitempty"`
}


func NewVersionIdentifierSpec() *VersionIdentifierSpec {
  p := new(VersionIdentifierSpec)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "lcm.v4.common.VersionIdentifierSpec"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "lcm.v4.r0.a1.common.VersionIdentifierSpec"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



