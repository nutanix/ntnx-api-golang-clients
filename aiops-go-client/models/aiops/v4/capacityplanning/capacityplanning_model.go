/*
 * Generated file models/aiops/v4/capacityplanning/capacityplanning_model.go.
 *
 * Product version: 4.0.2
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module aiops.v4.capacityplanning of Nutanix AIOps APIs
*/
package capacityplanning

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
)

type CapacityDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapacityUnit *CapacityUnit `json:"capacityUnit,omitempty"`
	/*
	  Effective capacity of the resource.
	*/
	EffectiveCapacity *float64 `json:"effectiveCapacity,omitempty"`
	/*
	  Runway details for the given cluster.
	*/
	Runway *int64 `json:"runway,omitempty"`

	RunwayDuration *RunwayDuration `json:"runwayDuration,omitempty"`
	/*
	  Total capacity of the resource.
	*/
	TotalCapacity *float64 `json:"totalCapacity,omitempty"`
	/*
	  Used capacity of the resource.
	*/
	UsedCapacity *float64 `json:"usedCapacity,omitempty"`
}

func (p *CapacityDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CapacityDetails

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

func (p *CapacityDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CapacityDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CapacityDetails(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityUnit")
	delete(allFields, "effectiveCapacity")
	delete(allFields, "runway")
	delete(allFields, "runwayDuration")
	delete(allFields, "totalCapacity")
	delete(allFields, "usedCapacity")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCapacityDetails() *CapacityDetails {
	p := new(CapacityDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.capacityplanning.CapacityDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CapacityUnit int

const (
	CAPACITYUNIT_UNKNOWN  CapacityUnit = 0
	CAPACITYUNIT_REDACTED CapacityUnit = 1
	CAPACITYUNIT_GHZ      CapacityUnit = 2
	CAPACITYUNIT_GB       CapacityUnit = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CapacityUnit) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GHz",
		"GB",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CapacityUnit) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GHz",
		"GB",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CapacityUnit) index(name string) CapacityUnit {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GHz",
		"GB",
	}
	for idx := range names {
		if names[idx] == name {
			return CapacityUnit(idx)
		}
	}
	return CAPACITYUNIT_UNKNOWN
}

func (e *CapacityUnit) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CapacityUnit:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CapacityUnit) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CapacityUnit) Ref() *CapacityUnit {
	return &e
}

type Runway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Cpu *CapacityDetails `json:"cpu,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Memory *CapacityDetails `json:"memory,omitempty"`

	Storage *CapacityDetails `json:"storage,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Runway) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Runway

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

func (p *Runway) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Runway
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Runway(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpu")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "memory")
	delete(allFields, "storage")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRunway() *Runway {
	p := new(Runway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.capacityplanning.Runway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RunwayDuration int

const (
	RUNWAYDURATION_UNKNOWN  RunwayDuration = 0
	RUNWAYDURATION_REDACTED RunwayDuration = 1
	RUNWAYDURATION_DAYS     RunwayDuration = 2
	RUNWAYDURATION_MONTH    RunwayDuration = 3
	RUNWAYDURATION_YEAR     RunwayDuration = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RunwayDuration) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAYS",
		"MONTH",
		"YEAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RunwayDuration) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAYS",
		"MONTH",
		"YEAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RunwayDuration) index(name string) RunwayDuration {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DAYS",
		"MONTH",
		"YEAR",
	}
	for idx := range names {
		if names[idx] == name {
			return RunwayDuration(idx)
		}
	}
	return RUNWAYDURATION_UNKNOWN
}

func (e *RunwayDuration) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RunwayDuration:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RunwayDuration) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RunwayDuration) Ref() *RunwayDuration {
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
