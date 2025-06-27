/*
 * Generated file models/dataprotection/v4/common/common_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.common of Nutanix Virtual Machine Management APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	"time"
)

/*
The backup type specifies the criteria for identifying the files to be backed up. This property should be specified to the application-consistent recovery points for Windows VMs/agents. The following backup types are supported for the application-consistent recovery points:
1. FULL_BACKUP - All files are backed up irrespective of their last backup date/time and state. Also, this backup type updates the backup history of every file that was involved in the recovery point.
2. COPY_BACKUP - All files are backed up regardless of their last backup date/time and state. However, this backup type does not update the backup history of individual files involved in the recovery point.
*/
type BackupType int

const (
	BACKUPTYPE_UNKNOWN     BackupType = 0
	BACKUPTYPE_REDACTED    BackupType = 1
	BACKUPTYPE_FULL_BACKUP BackupType = 2
	BACKUPTYPE_COPY_BACKUP BackupType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BackupType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_BACKUP",
		"COPY_BACKUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BackupType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_BACKUP",
		"COPY_BACKUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BackupType) index(name string) BackupType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FULL_BACKUP",
		"COPY_BACKUP",
	}
	for idx := range names {
		if names[idx] == name {
			return BackupType(idx)
		}
	}
	return BACKUPTYPE_UNKNOWN
}

func (e *BackupType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BackupType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BackupType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BackupType) Ref() *BackupType {
	return &e
}

/*
A model that represents common properties of a Recovery point resources
*/
type BaseRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BaseRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseRecoveryPoint

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *BaseRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BaseRecoveryPoint(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "creationTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "name")
	delete(allFields, "recoveryPointType")
	delete(allFields, "status")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBaseRecoveryPoint() *BaseRecoveryPoint {
	p := new(BaseRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.BaseRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents common properties of a VM Recovery point resources
*/
type BaseVmRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ApplicationConsistentPropertiesItemDiscriminator_ *string `json:"$applicationConsistentPropertiesItemDiscriminator,omitempty"`
	/*
	  User-defined application-consistent properties for the recovery point.
	*/
	ApplicationConsistentProperties *OneOfBaseVmRecoveryPointApplicationConsistentProperties `json:"applicationConsistentProperties,omitempty"`
	/*
	  External identifier of the Consistency group which the VM was part of at the time of recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DiskRecoveryPoints []DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is displayed as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as a part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
}

func (p *BaseVmRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseVmRecoveryPoint

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *BaseVmRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseVmRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BaseVmRecoveryPoint(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$applicationConsistentPropertiesItemDiscriminator")
	delete(allFields, "applicationConsistentProperties")
	delete(allFields, "consistencyGroupExtId")
	delete(allFields, "creationTime")
	delete(allFields, "diskRecoveryPoints")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "name")
	delete(allFields, "recoveryPointType")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "vmCategories")
	delete(allFields, "vmExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBaseVmRecoveryPoint() *BaseVmRecoveryPoint {
	p := new(BaseVmRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.BaseVmRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *BaseVmRecoveryPoint) GetApplicationConsistentProperties() interface{} {
	if nil == p.ApplicationConsistentProperties {
		return nil
	}
	return p.ApplicationConsistentProperties.GetValue()
}

func (p *BaseVmRecoveryPoint) SetApplicationConsistentProperties(v interface{}) error {
	if nil == p.ApplicationConsistentProperties {
		p.ApplicationConsistentProperties = NewOneOfBaseVmRecoveryPointApplicationConsistentProperties()
	}
	e := p.ApplicationConsistentProperties.SetValue(v)
	if nil == e {
		if nil == p.ApplicationConsistentPropertiesItemDiscriminator_ {
			p.ApplicationConsistentPropertiesItemDiscriminator_ = new(string)
		}
		*p.ApplicationConsistentPropertiesItemDiscriminator_ = *p.ApplicationConsistentProperties.Discriminator
	}
	return e
}

/*
A model that represents the disk recovery point properties.
*/
type DiskRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk external identifier which is captured as a part of this recovery point.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
	/*
	  Disk recovery point identifier.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
}

func (p *DiskRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DiskRecoveryPoint

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DiskRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiskRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DiskRecoveryPoint(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskExtId")
	delete(allFields, "diskRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDiskRecoveryPoint() *DiskRecoveryPoint {
	p := new(DiskRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.DiskRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The status of the Recovery point, which indicates whether this Recovery point is fit to be consumed.
*/
type RecoveryPointStatus int

const (
	RECOVERYPOINTSTATUS_UNKNOWN  RecoveryPointStatus = 0
	RECOVERYPOINTSTATUS_REDACTED RecoveryPointStatus = 1
	RECOVERYPOINTSTATUS_COMPLETE RecoveryPointStatus = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPointStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMPLETE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPointStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMPLETE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPointStatus) index(name string) RecoveryPointStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMPLETE",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPointStatus(idx)
		}
	}
	return RECOVERYPOINTSTATUS_UNKNOWN
}

func (e *RecoveryPointStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPointStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPointStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPointStatus) Ref() *RecoveryPointStatus {
	return &e
}

/*
Type of the Recovery point.
*/
type RecoveryPointType int

const (
	RECOVERYPOINTTYPE_UNKNOWN                RecoveryPointType = 0
	RECOVERYPOINTTYPE_REDACTED               RecoveryPointType = 1
	RECOVERYPOINTTYPE_CRASH_CONSISTENT       RecoveryPointType = 2
	RECOVERYPOINTTYPE_APPLICATION_CONSISTENT RecoveryPointType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPointType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPointType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPointType) index(name string) RecoveryPointType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CRASH_CONSISTENT",
		"APPLICATION_CONSISTENT",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPointType(idx)
		}
	}
	return RECOVERYPOINTTYPE_UNKNOWN
}

func (e *RecoveryPointType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPointType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPointType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPointType) Ref() *RecoveryPointType {
	return &e
}

/*
For Windows VMs/agents, application-consistent recovery points are selected using the Windows-specific Volume Shadow Copy Service (VSS). The enclosed properties, also called VSS properties, are used by the Windows platform to decide the type of application-consistent recovery points to consider.
*/
type VssProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackupType *BackupType `json:"backupType"`
	/*
	  Indicates whether the given set of VSS writers' UUIDs should be included or excluded from the application consistent recovery point. By default, the value is set to false, indicating that all listed VSS writers' UUIDs will be excluded.
	*/
	ShouldIncludeWriters *bool `json:"shouldIncludeWriters,omitempty"`
	/*
	  Indicates whether to store the VSS metadata if the user is interested in application-specific backup/restore. The VSS metadata consists of VSS writers and requester metadata details. These are compressed into a cabinet file(.cab file) during a VSS backup operation. This cabinet file must be saved to the backup media during a backup operation, as it is required during the restore operation.
	*/
	ShouldStoreVssMetadata *bool `json:"shouldStoreVssMetadata,omitempty"`
	/*
	  List of VSS writer UUIDs that are used in an application consistent recovery point. The default values are the system and the registry writer UUIDs.
	*/
	Writers []string `json:"writers,omitempty"`
}

func (p *VssProperties) MarshalJSON() ([]byte, error) {
	type VssPropertiesProxy VssProperties

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VssPropertiesProxy
		BackupType *BackupType `json:"backupType,omitempty"`
	}{
		VssPropertiesProxy: (*VssPropertiesProxy)(p),
		BackupType:         p.BackupType,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VssProperties) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VssProperties
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VssProperties(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backupType")
	delete(allFields, "shouldIncludeWriters")
	delete(allFields, "shouldStoreVssMetadata")
	delete(allFields, "writers")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVssProperties() *VssProperties {
	p := new(VssProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.VssProperties"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldIncludeWriters = new(bool)
	*p.ShouldIncludeWriters = false
	p.ShouldStoreVssMetadata = new(bool)
	*p.ShouldStoreVssMetadata = false

	return p
}

type OneOfBaseVmRecoveryPointApplicationConsistentProperties struct {
	Discriminator *string        `json:"-"`
	ObjectType_   *string        `json:"-"`
	oneOfType2001 *VssProperties `json:"-"`
}

func NewOneOfBaseVmRecoveryPointApplicationConsistentProperties() *OneOfBaseVmRecoveryPointApplicationConsistentProperties {
	p := new(OneOfBaseVmRecoveryPointApplicationConsistentProperties)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBaseVmRecoveryPointApplicationConsistentProperties) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBaseVmRecoveryPointApplicationConsistentProperties is nil"))
	}
	switch v.(type) {
	case VssProperties:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VssProperties)
		}
		*p.oneOfType2001 = v.(VssProperties)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBaseVmRecoveryPointApplicationConsistentProperties) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfBaseVmRecoveryPointApplicationConsistentProperties) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VssProperties)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.common.VssProperties" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VssProperties)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBaseVmRecoveryPointApplicationConsistentProperties"))
}

func (p *OneOfBaseVmRecoveryPointApplicationConsistentProperties) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfBaseVmRecoveryPointApplicationConsistentProperties")
}

type FileDetail struct {
	Path        *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}
