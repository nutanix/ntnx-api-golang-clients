/*
 * Generated file models/dataprotection/v4/common/common_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Volumes APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.common of Nutanix Volumes APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/common/v1/response"
	"time"
)

/*
A model that represents common properties of a recovery point resource.
*/
type BaseRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current recovery point expires and will be removed.
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
	  Location agnostic identifier of the recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier for the project associated with the resource. This field is required in create requests for authorization and must match the project identifier of all associated entities.
	*/
	ProjectExtId *string `json:"projectExtId,omitempty"`

	RecoveryPointType *RecoveryPointType `json:"recoveryPointType,omitempty"`

	SourceLocation *DisasterRecoveryLocation `json:"sourceLocation,omitempty"`

	Status *RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of the recovery point, which is the total space that could be reclaimed after deleting the recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
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
	*p = *NewBaseRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.SourceLocation != nil {
		p.SourceLocation = known.SourceLocation
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}

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
	delete(allFields, "projectExtId")
	delete(allFields, "recoveryPointType")
	delete(allFields, "sourceLocation")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

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
A model that represents common properties of a volume group recovery point resource.
*/
type BaseVolumeGroupRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the consistency group that the volume group was part of at the time of recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DiskRecoveryPoints []DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current recovery point expires and will be removed.
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
	  Location agnostic identifier of the recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier for the project associated with the resource. This field is required in create requests for authorization and must match the project identifier of all associated entities.
	*/
	ProjectExtId *string `json:"projectExtId,omitempty"`

	RecoveryPointType *RecoveryPointType `json:"recoveryPointType,omitempty"`

	SourceLocation *DisasterRecoveryLocation `json:"sourceLocation,omitempty"`

	Status *RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of the recovery point, which is the total space that could be reclaimed after deleting the recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
	/*
	  Category key-value pairs associated with the volume group at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is displayed as 'dept/hr'.
	*/
	VolumeGroupCategories []string `json:"volumeGroupCategories,omitempty"`
	/*
	  Volume group external identifier which is captured as a part of this recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func (p *BaseVolumeGroupRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseVolumeGroupRecoveryPoint

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

func (p *BaseVolumeGroupRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseVolumeGroupRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseVolumeGroupRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConsistencyGroupExtId != nil {
		p.ConsistencyGroupExtId = known.ConsistencyGroupExtId
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.DiskRecoveryPoints != nil {
		p.DiskRecoveryPoints = known.DiskRecoveryPoints
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.SourceLocation != nil {
		p.SourceLocation = known.SourceLocation
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}
	if known.VolumeGroupCategories != nil {
		p.VolumeGroupCategories = known.VolumeGroupCategories
	}
	if known.VolumeGroupExtId != nil {
		p.VolumeGroupExtId = known.VolumeGroupExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "consistencyGroupExtId")
	delete(allFields, "creationTime")
	delete(allFields, "diskRecoveryPoints")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "name")
	delete(allFields, "projectExtId")
	delete(allFields, "recoveryPointType")
	delete(allFields, "sourceLocation")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")
	delete(allFields, "volumeGroupCategories")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseVolumeGroupRecoveryPoint() *BaseVolumeGroupRecoveryPoint {
	p := new(BaseVolumeGroupRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.BaseVolumeGroupRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
disasterRecoveryLocationDesc
*/
type DisasterRecoveryLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster(s) associated with the recovery point at source Prism Central.
	*/
	ClusterExtIds []string `json:"clusterExtIds"`
	/*
	  External identifier of the Prism Central (also known as the Domain Manager) where the recovery point was first created.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *DisasterRecoveryLocation) MarshalJSON() ([]byte, error) {
	type DisasterRecoveryLocationProxy DisasterRecoveryLocation

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DisasterRecoveryLocationProxy
		ClusterExtIds      []string `json:"clusterExtIds,omitempty"`
		DomainManagerExtId *string  `json:"domainManagerExtId,omitempty"`
	}{
		DisasterRecoveryLocationProxy: (*DisasterRecoveryLocationProxy)(p),
		ClusterExtIds:                 p.ClusterExtIds,
		DomainManagerExtId:            p.DomainManagerExtId,
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

func (p *DisasterRecoveryLocation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DisasterRecoveryLocation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDisasterRecoveryLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDisasterRecoveryLocation() *DisasterRecoveryLocation {
	p := new(DisasterRecoveryLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.DisasterRecoveryLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents the disk recovery point properties.
*/
type DiskRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk external identifier captured as a part of the recovery point.
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
	*p = *NewDiskRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DiskExtId != nil {
		p.DiskExtId = known.DiskExtId
	}
	if known.DiskRecoveryPointExtId != nil {
		p.DiskRecoveryPointExtId = known.DiskRecoveryPointExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskExtId")
	delete(allFields, "diskRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

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
The status of the recovery point indicating whether the recovery point is fit to be consumed.
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
Type of the recovery point.
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
