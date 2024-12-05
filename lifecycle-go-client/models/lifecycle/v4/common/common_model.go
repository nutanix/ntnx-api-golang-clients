/*
 * Generated file models/lifecycle/v4/common/common_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Common Lifecycle resources.
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
)

/*
Available version types.
*/
type AvailableVersionStatus int

const (
	AVAILABLEVERSIONSTATUS_UNKNOWN     AvailableVersionStatus = 0
	AVAILABLEVERSIONSTATUS_REDACTED    AvailableVersionStatus = 1
	AVAILABLEVERSIONSTATUS_RECOMMENDED AvailableVersionStatus = 2
	AVAILABLEVERSIONSTATUS_CRITICAL    AvailableVersionStatus = 3
	AVAILABLEVERSIONSTATUS_LATEST      AvailableVersionStatus = 4
	AVAILABLEVERSIONSTATUS_DEPRECATED  AvailableVersionStatus = 5
	AVAILABLEVERSIONSTATUS_EMERGENCY   AvailableVersionStatus = 6
	AVAILABLEVERSIONSTATUS_AVAILABLE   AvailableVersionStatus = 7
	AVAILABLEVERSIONSTATUS_LTS         AvailableVersionStatus = 8
	AVAILABLEVERSIONSTATUS_STS         AvailableVersionStatus = 9
	AVAILABLEVERSIONSTATUS_ESTS        AvailableVersionStatus = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AvailableVersionStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AvailableVersionStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AvailableVersionStatus) index(name string) AvailableVersionStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	for idx := range names {
		if names[idx] == name {
			return AvailableVersionStatus(idx)
		}
	}
	return AVAILABLEVERSIONSTATUS_UNKNOWN
}

func (e *AvailableVersionStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AvailableVersionStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AvailableVersionStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AvailableVersionStatus) Ref() *AvailableVersionStatus {
	return &e
}

/*
Checksum type for a third party image.
*/
type CheckSumType int

const (
	CHECKSUMTYPE_UNKNOWN  CheckSumType = 0
	CHECKSUMTYPE_REDACTED CheckSumType = 1
	CHECKSUMTYPE_SHASUM   CheckSumType = 2
	CHECKSUMTYPE_HEX_MD5  CheckSumType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CheckSumType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CheckSumType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CheckSumType) index(name string) CheckSumType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	for idx := range names {
		if names[idx] == name {
			return CheckSumType(idx)
		}
	}
	return CHECKSUMTYPE_UNKNOWN
}

func (e *CheckSumType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CheckSumType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CheckSumType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CheckSumType) Ref() *CheckSumType {
	return &e
}

/*
Specification for deployment of entities.
*/
type DeploySpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of deploy specifications for use in the recommendations API.
	*/
	EntityDeploySpecs []EntityDeploySpec `json:"entityDeploySpecs"`
}

func (p *DeploySpec) MarshalJSON() ([]byte, error) {
	type DeploySpecProxy DeploySpec
	return json.Marshal(struct {
		*DeploySpecProxy
		EntityDeploySpecs []EntityDeploySpec `json:"entityDeploySpecs,omitempty"`
	}{
		DeploySpecProxy:   (*DeploySpecProxy)(p),
		EntityDeploySpecs: p.EntityDeploySpecs,
	})
}

func NewDeploySpec() *DeploySpec {
	p := new(DeploySpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.DeploySpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The list of properties that can be expanded on the LCM entity.
*/
type EntityBaseModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntityBaseModel() *EntityBaseModel {
	p := new(EntityBaseModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityBaseModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A specification defining the entity being deployed and its version.
*/
type EntityDeploySpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityIdentifier *EntityBaseModel `json:"entityIdentifier"`
}

func (p *EntityDeploySpec) MarshalJSON() ([]byte, error) {
	type EntityDeploySpecProxy EntityDeploySpec
	return json.Marshal(struct {
		*EntityDeploySpecProxy
		EntityIdentifier *EntityBaseModel `json:"entityIdentifier,omitempty"`
	}{
		EntityDeploySpecProxy: (*EntityDeploySpecProxy)(p),
		EntityIdentifier:      p.EntityIdentifier,
	})
}

func NewEntityDeploySpec() *EntityDeploySpec {
	p := new(EntityDeploySpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityDeploySpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of an LCM entity.
*/
type EntityType int

const (
	ENTITYTYPE_UNKNOWN  EntityType = 0
	ENTITYTYPE_REDACTED EntityType = 1
	ENTITYTYPE_SOFTWARE EntityType = 2
	ENTITYTYPE_FIRMWARE EntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EntityType) index(name string) EntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	for idx := range names {
		if names[idx] == name {
			return EntityType(idx)
		}
	}
	return ENTITYTYPE_UNKNOWN
}

func (e *EntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EntityType) Ref() *EntityType {
	return &e
}

/*
Specification for running an update operation.
*/
type EntityUpdateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the LCM entity.
	*/
	EntityUuid *string `json:"entityUuid"`
	/*
	  Version to upgrade to.
	*/
	ToVersion *string `json:"toVersion"`
}

func (p *EntityUpdateSpec) MarshalJSON() ([]byte, error) {
	type EntityUpdateSpecProxy EntityUpdateSpec
	return json.Marshal(struct {
		*EntityUpdateSpecProxy
		EntityUuid *string `json:"entityUuid,omitempty"`
		ToVersion  *string `json:"toVersion,omitempty"`
	}{
		EntityUpdateSpecProxy: (*EntityUpdateSpecProxy)(p),
		EntityUuid:            p.EntityUuid,
		ToVersion:             p.ToVersion,
	})
}

func NewEntityUpdateSpec() *EntityUpdateSpec {
	p := new(EntityUpdateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityUpdateSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of Hypervisor present in the cluster.
*/
type HypervisorType int

const (
	HYPERVISORTYPE_UNKNOWN  HypervisorType = 0
	HYPERVISORTYPE_REDACTED HypervisorType = 1
	HYPERVISORTYPE_ESX      HypervisorType = 2
	HYPERVISORTYPE_AHV      HypervisorType = 3
	HYPERVISORTYPE_HYPERV   HypervisorType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HypervisorType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HypervisorType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HypervisorType) index(name string) HypervisorType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
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

/*
Details of the in progress LCM operation.
*/
type InProgressOpDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  Task ext id of the in progress LCM operation.
	*/
	TaskExtId *string `json:"taskExtId,omitempty"`
}

func NewInProgressOpDetails() *InProgressOpDetails {
	p := new(InProgressOpDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.InProgressOpDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
MD5Sum of the bundle.
*/
type LcmMd5Sum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hex digest of the MD5 sum.
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *LcmMd5Sum) MarshalJSON() ([]byte, error) {
	type LcmMd5SumProxy LcmMd5Sum
	return json.Marshal(struct {
		*LcmMd5SumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		LcmMd5SumProxy: (*LcmMd5SumProxy)(p),
		HexDigest:      p.HexDigest,
	})
}

func NewLcmMd5Sum() *LcmMd5Sum {
	p := new(LcmMd5Sum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LcmMd5Sum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Sha256sum of the bundle.
*/
type LcmSha256Sum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hex digest of the SHA256 sum.
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *LcmSha256Sum) MarshalJSON() ([]byte, error) {
	type LcmSha256SumProxy LcmSha256Sum
	return json.Marshal(struct {
		*LcmSha256SumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		LcmSha256SumProxy: (*LcmSha256SumProxy)(p),
		HexDigest:         p.HexDigest,
	})
}

func NewLcmSha256Sum() *LcmSha256Sum {
	p := new(LcmSha256Sum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LcmSha256Sum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Location info corresponds to a tuple of location type (either node/cluster) and ExtID
*/
type LocationInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LocationType *LocationType `json:"locationType,omitempty"`
	/*
	  Location UUID of the resource.
	*/
	Uuid *string `json:"uuid,omitempty"`
}

func NewLocationInfo() *LocationInfo {
	p := new(LocationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LocationInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Scope of entity represented in LCM. This could be either Node or cluster type.
*/
type LocationType int

const (
	LOCATIONTYPE_UNKNOWN  LocationType = 0
	LOCATIONTYPE_REDACTED LocationType = 1
	LOCATIONTYPE_NODE     LocationType = 2
	LOCATIONTYPE_CLUSTER  LocationType = 3
	LOCATIONTYPE_PC       LocationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LocationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LocationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LocationType) index(name string) LocationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	for idx := range names {
		if names[idx] == name {
			return LocationType(idx)
		}
	}
	return LOCATIONTYPE_UNKNOWN
}

func (e *LocationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LocationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LocationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LocationType) Ref() *LocationType {
	return &e
}

/*
Cluster management server configuration used while updating clusters with ESX or Hyper-V.
*/
type ManagementServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType"`
	/*
	  IP address of the management server.
	*/
	Ip *string `json:"ip"`
	/*
	  Password to login to the management server.
	*/
	Password *string `json:"password"`
	/*
	  Username to login to the management server.
	*/
	Username *string `json:"username"`
}

func (p *ManagementServer) MarshalJSON() ([]byte, error) {
	type ManagementServerProxy ManagementServer
	return json.Marshal(struct {
		*ManagementServerProxy
		HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
		Ip             *string         `json:"ip,omitempty"`
		Password       *string         `json:"password,omitempty"`
		Username       *string         `json:"username,omitempty"`
	}{
		ManagementServerProxy: (*ManagementServerProxy)(p),
		HypervisorType:        p.HypervisorType,
		Ip:                    p.Ip,
		Password:              p.Password,
		Username:              p.Username,
	})
}

func NewManagementServer() *ManagementServer {
	p := new(ManagementServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.ManagementServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of LCM upgrade notification generated. It can be any one of the values like Entity, Location, Generic or Workflow. The only types supported today are Entity and Location.
*/
type NotificationType int

const (
	NOTIFICATIONTYPE_UNKNOWN  NotificationType = 0
	NOTIFICATIONTYPE_REDACTED NotificationType = 1
	NOTIFICATIONTYPE_ENTITY   NotificationType = 2
	NOTIFICATIONTYPE_LOCATION NotificationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NotificationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NotificationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NotificationType) index(name string) NotificationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	for idx := range names {
		if names[idx] == name {
			return NotificationType(idx)
		}
	}
	return NOTIFICATIONTYPE_UNKNOWN
}

func (e *NotificationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NotificationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NotificationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NotificationType) Ref() *NotificationType {
	return &e
}

/*
Type of the operation tracked by the task.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN   OperationType = 0
	OPERATIONTYPE_REDACTED  OperationType = 1
	OPERATIONTYPE_INVENTORY OperationType = 2
	OPERATIONTYPE_PRECHECKS OperationType = 3
	OPERATIONTYPE_UPGRADE   OperationType = 4
	OPERATIONTYPE_NONE      OperationType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationType(idx)
		}
	}
	return OPERATIONTYPE_UNKNOWN
}

func (e *OperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationType) Ref() *OperationType {
	return &e
}

/*
Specification for running a Precheck operation.
*/
type PrechecksSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`

	ManagementServer *ManagementServer `json:"managementServer,omitempty"`
	/*
	  List of prechecks to skip. The allowed value is 'powerOffUvms' that skips the pinned VM prechecks.
	*/
	SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
}

func (p *PrechecksSpec) MarshalJSON() ([]byte, error) {
	type PrechecksSpecProxy PrechecksSpec
	return json.Marshal(struct {
		*PrechecksSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		PrechecksSpecProxy: (*PrechecksSpecProxy)(p),
		EntityUpdateSpecs:  p.EntityUpdateSpecs,
	})
}

func NewPrechecksSpec() *PrechecksSpec {
	p := new(PrechecksSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.PrechecksSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A specification defining the entity being preloaded and its version.
*/
type PreloadSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`
}

func (p *PreloadSpec) MarshalJSON() ([]byte, error) {
	type PreloadSpecProxy PreloadSpec
	return json.Marshal(struct {
		*PreloadSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		PreloadSpecProxy:  (*PreloadSpecProxy)(p),
		EntityUpdateSpecs: p.EntityUpdateSpecs,
	})
}

func NewPreloadSpec() *PreloadSpec {
	p := new(PreloadSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.PreloadSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
System auto-management flag to handle system operation during upgrade.
*/
type SystemAutoMgmtFlag int

const (
	SYSTEMAUTOMGMTFLAG_UNKNOWN        SystemAutoMgmtFlag = 0
	SYSTEMAUTOMGMTFLAG_REDACTED       SystemAutoMgmtFlag = 1
	SYSTEMAUTOMGMTFLAG_POWER_OFF_UVMS SystemAutoMgmtFlag = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SystemAutoMgmtFlag) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SystemAutoMgmtFlag) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SystemAutoMgmtFlag) index(name string) SystemAutoMgmtFlag {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
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

/*
Specification for an upgrade operation of an entity to a particular target version.
*/
type UpgradeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of automated system operations to perform, to avoid precheck failure and let the system restore state after an update is complete. The allowed flag is: - 'powerOffUvms': This allows the system to automatically power off user VMs which cannot be migrated to other hosts and power them on when the update is done. This option can avoid pinned VM precheck failure on the host which needs to enter maintenance mode during the update and allow the update to go through.
	*/
	AutoHandleFlags []SystemAutoMgmtFlag `json:"autoHandleFlags,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`

	ManagementServer *ManagementServer `json:"managementServer,omitempty"`
	/*
	  Number of seconds LCM waits for the VMs to come up after exiting host maintenance mode.
	*/
	MaxWaitTimeInSecs *int `json:"maxWaitTimeInSecs,omitempty"`
	/*
	  List of prechecks to skip. The allowed value is 'powerOffUvms' that skips the pinned VM prechecks.
	*/
	SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
}

func (p *UpgradeSpec) MarshalJSON() ([]byte, error) {
	type UpgradeSpecProxy UpgradeSpec
	return json.Marshal(struct {
		*UpgradeSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		UpgradeSpecProxy:  (*UpgradeSpecProxy)(p),
		EntityUpdateSpecs: p.EntityUpdateSpecs,
	})
}

func NewUpgradeSpec() *UpgradeSpec {
	p := new(UpgradeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.UpgradeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
