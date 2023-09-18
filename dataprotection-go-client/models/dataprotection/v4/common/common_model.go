/*
 * Generated file models/dataprotection/v4/common/common_model.go.
 *
 * Product version: 4.0.1-alpha-4
 *
 * Part of the Nutanix Dataprotection Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.common of Nutanix Dataprotection Versioned APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	"time"
)

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
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`
}

func NewBaseRecoveryPoint() *BaseRecoveryPoint {
	p := new(BaseRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.BaseRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a2.common.BaseRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The status of the Recovery point, which indicates whether this Recovery point is fit to be consumed.
*/
type RecoveryPointStatus int

const (
	RECOVERYPOINTSTATUS_UNKNOWN    RecoveryPointStatus = 0
	RECOVERYPOINTSTATUS_REDACTED   RecoveryPointStatus = 1
	RECOVERYPOINTSTATUS_INCOMPLETE RecoveryPointStatus = 2
	RECOVERYPOINTSTATUS_COMPLETE   RecoveryPointStatus = 3
	RECOVERYPOINTSTATUS_EXPIRED    RecoveryPointStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPointStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INCOMPLETE",
		"COMPLETE",
		"EXPIRED",
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
		"INCOMPLETE",
		"COMPLETE",
		"EXPIRED",
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
		"INCOMPLETE",
		"COMPLETE",
		"EXPIRED",
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
Additional metadata provided by the client at the time of Recovery point creation.
*/
type VendorSpecificProperty struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The unique identifier of the vendor. It can be a magic number, UUID or vendor name.
	*/
	VendorId *string `json:"vendorId"`
	/*
	  This is an opaque data interpreted only by the respective vendor.
	*/
	VendorMetadata *string `json:"vendorMetadata"`
}

func (p *VendorSpecificProperty) MarshalJSON() ([]byte, error) {
	type VendorSpecificPropertyProxy VendorSpecificProperty
	return json.Marshal(struct {
		*VendorSpecificPropertyProxy
		VendorId       *string `json:"vendorId,omitempty"`
		VendorMetadata *string `json:"vendorMetadata,omitempty"`
	}{
		VendorSpecificPropertyProxy: (*VendorSpecificPropertyProxy)(p),
		VendorId:                    p.VendorId,
		VendorMetadata:              p.VendorMetadata,
	})
}

func NewVendorSpecificProperty() *VendorSpecificProperty {
	p := new(VendorSpecificProperty)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.common.VendorSpecificProperty"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a2.common.VendorSpecificProperty"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}
