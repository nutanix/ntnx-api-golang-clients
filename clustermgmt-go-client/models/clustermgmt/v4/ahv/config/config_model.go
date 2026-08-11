/*
 * Generated file models/clustermgmt/v4/ahv/config/config_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Cluster Management APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.ahv.config of Nutanix Cluster Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
	import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
)

/*
REST response for all response codes in API path /clustermgmt/v4.3/ahv/config/physical-gpu-profiles Get operation
*/
type ListAhvPhysicalGpuProfilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAhvPhysicalGpuProfilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListAhvPhysicalGpuProfilesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListAhvPhysicalGpuProfilesApiResponse

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

func (p *ListAhvPhysicalGpuProfilesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListAhvPhysicalGpuProfilesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListAhvPhysicalGpuProfilesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListAhvPhysicalGpuProfilesApiResponse() *ListAhvPhysicalGpuProfilesApiResponse {
	p := new(ListAhvPhysicalGpuProfilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.ListAhvPhysicalGpuProfilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAhvPhysicalGpuProfilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAhvPhysicalGpuProfilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAhvPhysicalGpuProfilesApiResponseData()
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

/*
REST response for all response codes in API path /clustermgmt/v4.3/ahv/config/virtual-gpu-profiles Get operation
*/
type ListAhvVirtualGpuProfilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAhvVirtualGpuProfilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListAhvVirtualGpuProfilesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListAhvVirtualGpuProfilesApiResponse

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

func (p *ListAhvVirtualGpuProfilesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListAhvVirtualGpuProfilesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListAhvVirtualGpuProfilesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListAhvVirtualGpuProfilesApiResponse() *ListAhvVirtualGpuProfilesApiResponse {
	p := new(ListAhvVirtualGpuProfilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.ListAhvVirtualGpuProfilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAhvVirtualGpuProfilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAhvVirtualGpuProfilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAhvVirtualGpuProfilesApiResponseData()
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

/*
REST response for all response codes in API path /clustermgmt/v4.3/ahv/config/pcie-devices Get operation
*/
type ListPcieDevicesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListPcieDevicesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListPcieDevicesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListPcieDevicesApiResponse

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

func (p *ListPcieDevicesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListPcieDevicesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListPcieDevicesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListPcieDevicesApiResponse() *ListPcieDevicesApiResponse {
	p := new(ListPcieDevicesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.ListPcieDevicesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListPcieDevicesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListPcieDevicesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListPcieDevicesApiResponseData()
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

/*
PCI Device entity description
*/
type PcieDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	Configuration *PcieDeviceConfiguration `json:"configuration,omitempty"`
	/*
	  Human readable device description
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  UUID of the host connected to the device
	*/
	HostExtId *string `json:"hostExtId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Maximum number of partitions supported by the device. Only populated when the device is in HOST_PARTITIONED state
	*/
	MaxPartitions *int64 `json:"maxPartitions,omitempty"`
	/*
	  UUID of the VM attached to that device when state is UVM_ASSIGNED or UVM_RESERVED
	*/
	OwnerVmExtId *string `json:"ownerVmExtId,omitempty"`

	State *PcieDeviceState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *PcieDeviceType `json:"type,omitempty"`
}

func (p *PcieDevice) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcieDevice

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

func (p *PcieDevice) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcieDevice
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcieDevice()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostExtId != nil {
		p.HostExtId = known.HostExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MaxPartitions != nil {
		p.MaxPartitions = known.MaxPartitions
	}
	if known.OwnerVmExtId != nil {
		p.OwnerVmExtId = known.OwnerVmExtId
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "configuration")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "hostExtId")
	delete(allFields, "links")
	delete(allFields, "maxPartitions")
	delete(allFields, "ownerVmExtId")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcieDevice() *PcieDevice {
	p := new(PcieDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PcieDevice"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
PCIe device configuration
*/
type PcieDeviceConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Class code
	*/
	ClassId *int64 `json:"classId,omitempty"`
	/*
	  Device ID
	*/
	DeviceId *int64 `json:"deviceId,omitempty"`
	/*
	  Programming interface ID
	*/
	ProgIFace *int64 `json:"progIFace,omitempty"`
	/*
	  Subsystem class ID
	*/
	SubClassId *int64 `json:"subClassId,omitempty"`
	/*
	  Subsystem (device) ID
	*/
	SubSystemId *int64 `json:"subSystemId,omitempty"`
	/*
	  Subsystem vendor ID
	*/
	SubSystemVendorId *int64 `json:"subSystemVendorId,omitempty"`
	/*
	  Vendor ID
	*/
	VendorId *int64 `json:"vendorId,omitempty"`
}

func (p *PcieDeviceConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcieDeviceConfiguration

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

func (p *PcieDeviceConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcieDeviceConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcieDeviceConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClassId != nil {
		p.ClassId = known.ClassId
	}
	if known.DeviceId != nil {
		p.DeviceId = known.DeviceId
	}
	if known.ProgIFace != nil {
		p.ProgIFace = known.ProgIFace
	}
	if known.SubClassId != nil {
		p.SubClassId = known.SubClassId
	}
	if known.SubSystemId != nil {
		p.SubSystemId = known.SubSystemId
	}
	if known.SubSystemVendorId != nil {
		p.SubSystemVendorId = known.SubSystemVendorId
	}
	if known.VendorId != nil {
		p.VendorId = known.VendorId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "classId")
	delete(allFields, "deviceId")
	delete(allFields, "progIFace")
	delete(allFields, "subClassId")
	delete(allFields, "subSystemId")
	delete(allFields, "subSystemVendorId")
	delete(allFields, "vendorId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcieDeviceConfiguration() *PcieDeviceConfiguration {
	p := new(PcieDeviceConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PcieDeviceConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Device state
*/
type PcieDeviceState int

const (
	PCIEDEVICESTATE_UNKNOWN          PcieDeviceState = 0
	PCIEDEVICESTATE_REDACTED         PcieDeviceState = 1
	PCIEDEVICESTATE_UVM_AVAILABLE    PcieDeviceState = 2
	PCIEDEVICESTATE_UVM_RESERVED     PcieDeviceState = 3
	PCIEDEVICESTATE_UVM_ASSIGNED     PcieDeviceState = 4
	PCIEDEVICESTATE_HOST_BROKEN      PcieDeviceState = 5
	PCIEDEVICESTATE_HOST_UNUSED      PcieDeviceState = 6
	PCIEDEVICESTATE_HOST_USED        PcieDeviceState = 7
	PCIEDEVICESTATE_HOST_PARTITIONED PcieDeviceState = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PcieDeviceState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UVM_AVAILABLE",
		"UVM_RESERVED",
		"UVM_ASSIGNED",
		"HOST_BROKEN",
		"HOST_UNUSED",
		"HOST_USED",
		"HOST_PARTITIONED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PcieDeviceState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UVM_AVAILABLE",
		"UVM_RESERVED",
		"UVM_ASSIGNED",
		"HOST_BROKEN",
		"HOST_UNUSED",
		"HOST_USED",
		"HOST_PARTITIONED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PcieDeviceState) index(name string) PcieDeviceState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UVM_AVAILABLE",
		"UVM_RESERVED",
		"UVM_ASSIGNED",
		"HOST_BROKEN",
		"HOST_UNUSED",
		"HOST_USED",
		"HOST_PARTITIONED",
	}
	for idx := range names {
		if names[idx] == name {
			return PcieDeviceState(idx)
		}
	}
	return PCIEDEVICESTATE_UNKNOWN
}

func (e *PcieDeviceState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PcieDeviceState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PcieDeviceState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PcieDeviceState) Ref() *PcieDeviceState {
	return &e
}

/*
Device type according to the PCI code and ID assignment specification
*/
type PcieDeviceType int

const (
	PCIEDEVICETYPE_UNKNOWN                                           PcieDeviceType = 0
	PCIEDEVICETYPE_REDACTED                                          PcieDeviceType = 1
	PCIEDEVICETYPE_MASS_STORAGE_CONTROLLER                           PcieDeviceType = 2
	PCIEDEVICETYPE_NETWORK_CONTROLLER                                PcieDeviceType = 3
	PCIEDEVICETYPE_DISPLAY_CONTROLLER                                PcieDeviceType = 4
	PCIEDEVICETYPE_MULTIMEDIA_DEVICE                                 PcieDeviceType = 5
	PCIEDEVICETYPE_MEMORY_CONTROLLER                                 PcieDeviceType = 6
	PCIEDEVICETYPE_BRIDGE_DEVICE                                     PcieDeviceType = 7
	PCIEDEVICETYPE_SIMPLE_COMMUNICATION_CONTROLLER                   PcieDeviceType = 8
	PCIEDEVICETYPE_BASE_SYSTEM_PERIPHERAL                            PcieDeviceType = 9
	PCIEDEVICETYPE_INPUT_DEVICE                                      PcieDeviceType = 10
	PCIEDEVICETYPE_DOCKING_STATION                                   PcieDeviceType = 11
	PCIEDEVICETYPE_PROCESSOR                                         PcieDeviceType = 12
	PCIEDEVICETYPE_SERIAL_BUS_CONTROLLER                             PcieDeviceType = 13
	PCIEDEVICETYPE_WIRELESS_CONTROLLER                               PcieDeviceType = 14
	PCIEDEVICETYPE_INTELLIGENT_IO_CONTROLLER                         PcieDeviceType = 15
	PCIEDEVICETYPE_SATELLITE_COMMUNICATION_CONTROLLER                PcieDeviceType = 16
	PCIEDEVICETYPE_ENCRYPTION_DECRYPTION_CONTROLLER                  PcieDeviceType = 17
	PCIEDEVICETYPE_DATA_ACQUISITION_AND_SIGNAL_PROCESSING_CONTROLLER PcieDeviceType = 18
	PCIEDEVICETYPE_PROCESSING_ACCELERATOR                            PcieDeviceType = 19
	PCIEDEVICETYPE_NON_ESSENTIAL_INSTRUMENTATION                     PcieDeviceType = 20
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PcieDeviceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MASS_STORAGE_CONTROLLER",
		"NETWORK_CONTROLLER",
		"DISPLAY_CONTROLLER",
		"MULTIMEDIA_DEVICE",
		"MEMORY_CONTROLLER",
		"BRIDGE_DEVICE",
		"SIMPLE_COMMUNICATION_CONTROLLER",
		"BASE_SYSTEM_PERIPHERAL",
		"INPUT_DEVICE",
		"DOCKING_STATION",
		"PROCESSOR",
		"SERIAL_BUS_CONTROLLER",
		"WIRELESS_CONTROLLER",
		"INTELLIGENT_IO_CONTROLLER",
		"SATELLITE_COMMUNICATION_CONTROLLER",
		"ENCRYPTION_DECRYPTION_CONTROLLER",
		"DATA_ACQUISITION_AND_SIGNAL_PROCESSING_CONTROLLER",
		"PROCESSING_ACCELERATOR",
		"NON_ESSENTIAL_INSTRUMENTATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PcieDeviceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MASS_STORAGE_CONTROLLER",
		"NETWORK_CONTROLLER",
		"DISPLAY_CONTROLLER",
		"MULTIMEDIA_DEVICE",
		"MEMORY_CONTROLLER",
		"BRIDGE_DEVICE",
		"SIMPLE_COMMUNICATION_CONTROLLER",
		"BASE_SYSTEM_PERIPHERAL",
		"INPUT_DEVICE",
		"DOCKING_STATION",
		"PROCESSOR",
		"SERIAL_BUS_CONTROLLER",
		"WIRELESS_CONTROLLER",
		"INTELLIGENT_IO_CONTROLLER",
		"SATELLITE_COMMUNICATION_CONTROLLER",
		"ENCRYPTION_DECRYPTION_CONTROLLER",
		"DATA_ACQUISITION_AND_SIGNAL_PROCESSING_CONTROLLER",
		"PROCESSING_ACCELERATOR",
		"NON_ESSENTIAL_INSTRUMENTATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PcieDeviceType) index(name string) PcieDeviceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MASS_STORAGE_CONTROLLER",
		"NETWORK_CONTROLLER",
		"DISPLAY_CONTROLLER",
		"MULTIMEDIA_DEVICE",
		"MEMORY_CONTROLLER",
		"BRIDGE_DEVICE",
		"SIMPLE_COMMUNICATION_CONTROLLER",
		"BASE_SYSTEM_PERIPHERAL",
		"INPUT_DEVICE",
		"DOCKING_STATION",
		"PROCESSOR",
		"SERIAL_BUS_CONTROLLER",
		"WIRELESS_CONTROLLER",
		"INTELLIGENT_IO_CONTROLLER",
		"SATELLITE_COMMUNICATION_CONTROLLER",
		"ENCRYPTION_DECRYPTION_CONTROLLER",
		"DATA_ACQUISITION_AND_SIGNAL_PROCESSING_CONTROLLER",
		"PROCESSING_ACCELERATOR",
		"NON_ESSENTIAL_INSTRUMENTATION",
	}
	for idx := range names {
		if names[idx] == name {
			return PcieDeviceType(idx)
		}
	}
	return PCIEDEVICETYPE_UNKNOWN
}

func (e *PcieDeviceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PcieDeviceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PcieDeviceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PcieDeviceType) Ref() *PcieDeviceType {
	return &e
}

/*
A physical GPU profile representing GPU hardware that can be attached to virtual machines in passthrough mode. These profiles are automatically discovered from installed GPU hardware across registered clusters.
*/
type PhysicalGpuProfile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of clusters where the profile is available.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`

	Configuration *PhysicalGpuProfileConfiguration `json:"configuration,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *PhysicalGpuProfile) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhysicalGpuProfile

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

func (p *PhysicalGpuProfile) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalGpuProfile
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalGpuProfile()

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
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "configuration")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalGpuProfile() *PhysicalGpuProfile {
	p := new(PhysicalGpuProfile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PhysicalGpuProfile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details for a physical GPU profile, including the GPU hardware model identifier.
*/
type PhysicalGpuProfileConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The identifier for the GPU profile. For virtual GPU profiles, this indicates the profile type and resource allocation. For physical GPU profiles, this is the GPU hardware model.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *PhysicalGpuProfileConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhysicalGpuProfileConfiguration

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

func (p *PhysicalGpuProfileConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhysicalGpuProfileConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhysicalGpuProfileConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhysicalGpuProfileConfiguration() *PhysicalGpuProfileConfiguration {
	p := new(PhysicalGpuProfileConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PhysicalGpuProfileConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A virtual GPU profile that defines resource allocation for virtual GPUs. Profiles specify GPU resources such as frame buffer, display heads, and resolution available to virtual machines across registered clusters.
*/
type VirtualGpuProfile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of clusters where the profile is available.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`

	Configuration *VirtualGpuProfileConfiguration `json:"configuration,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VirtualGpuProfile) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VirtualGpuProfile

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

func (p *VirtualGpuProfile) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VirtualGpuProfile
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVirtualGpuProfile()

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
	if known.Configuration != nil {
		p.Configuration = known.Configuration
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "configuration")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVirtualGpuProfile() *VirtualGpuProfile {
	p := new(VirtualGpuProfile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.VirtualGpuProfile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details for a virtual GPU profile, including frame buffer allocation, display capabilities, resolution limits, and licensing requirements.
*/
type VirtualGpuProfileConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The amount of GPU frame buffer memory allocated to the virtual GPU, measured in bytes. Larger frame buffer sizes support higher resolutions and more complex workloads.
	*/
	FrameBufferSizeBytes *int64 `json:"frameBufferSizeBytes,omitempty"`
	/*
	  A list of software license types required or compatible with this virtual GPU profile. The license type determines which features and capabilities are available.
	*/
	Licenses []string `json:"licenses,omitempty"`
	/*
	  The maximum number of virtual GPU instances of this profile that can be created per physical GPU.
	*/
	MaxInstances *int64 `json:"maxInstances,omitempty"`
	/*
	  The maximum number of virtual GPU instances of this profile that can be assigned to a single virtual machine.
	*/
	MaxInstancesPerVm *int64 `json:"maxInstancesPerVm,omitempty"`
	/*
	  The maximum display resolution supported by the profile, formatted as 'width x height' (e.g., '4096x2160').
	*/
	MaxResolution *string `json:"maxResolution,omitempty"`
	/*
	  The identifier for the GPU profile. For virtual GPU profiles, this indicates the profile type and resource allocation. For physical GPU profiles, this is the GPU hardware model.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The number of virtual display heads (monitors) supported by this profile. Each display head can drive an independent display output.
	*/
	NumberOfVirtualDisplayHeads *int64 `json:"numberOfVirtualDisplayHeads,omitempty"`
}

func (p *VirtualGpuProfileConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VirtualGpuProfileConfiguration

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

func (p *VirtualGpuProfileConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VirtualGpuProfileConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVirtualGpuProfileConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FrameBufferSizeBytes != nil {
		p.FrameBufferSizeBytes = known.FrameBufferSizeBytes
	}
	if known.Licenses != nil {
		p.Licenses = known.Licenses
	}
	if known.MaxInstances != nil {
		p.MaxInstances = known.MaxInstances
	}
	if known.MaxInstancesPerVm != nil {
		p.MaxInstancesPerVm = known.MaxInstancesPerVm
	}
	if known.MaxResolution != nil {
		p.MaxResolution = known.MaxResolution
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NumberOfVirtualDisplayHeads != nil {
		p.NumberOfVirtualDisplayHeads = known.NumberOfVirtualDisplayHeads
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "frameBufferSizeBytes")
	delete(allFields, "licenses")
	delete(allFields, "maxInstances")
	delete(allFields, "maxInstancesPerVm")
	delete(allFields, "maxResolution")
	delete(allFields, "name")
	delete(allFields, "numberOfVirtualDisplayHeads")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVirtualGpuProfileConfiguration() *VirtualGpuProfileConfiguration {
	p := new(VirtualGpuProfileConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.VirtualGpuProfileConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListAhvPhysicalGpuProfilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []PhysicalGpuProfile   `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListAhvPhysicalGpuProfilesApiResponseData() *OneOfListAhvPhysicalGpuProfilesApiResponseData {
	p := new(OneOfListAhvPhysicalGpuProfilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAhvPhysicalGpuProfilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAhvPhysicalGpuProfilesApiResponseData is nil"))
	}
	switch v.(type) {
	case []PhysicalGpuProfile:
		p.oneOfType2001 = v.([]PhysicalGpuProfile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAhvPhysicalGpuProfilesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAhvPhysicalGpuProfilesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]PhysicalGpuProfile)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.PhysicalGpuProfile" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new([]PhysicalGpuProfile)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.PhysicalGpuProfile" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAhvPhysicalGpuProfilesApiResponseData"))
}

func (p *OneOfListAhvPhysicalGpuProfilesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<clustermgmt.v4.ahv.config.PhysicalGpuProfile>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAhvPhysicalGpuProfilesApiResponseData")
}

type OneOfListPcieDevicesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []PcieDevice           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListPcieDevicesApiResponseData() *OneOfListPcieDevicesApiResponseData {
	p := new(OneOfListPcieDevicesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListPcieDevicesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListPcieDevicesApiResponseData is nil"))
	}
	switch v.(type) {
	case []PcieDevice:
		p.oneOfType2001 = v.([]PcieDevice)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.ahv.config.PcieDevice>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PcieDevice>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListPcieDevicesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<clustermgmt.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListPcieDevicesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<clustermgmt.v4.ahv.config.PcieDevice>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]PcieDevice)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.PcieDevice" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<clustermgmt.v4.ahv.config.PcieDevice>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PcieDevice>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new([]PcieDevice)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.PcieDevice" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.ahv.config.PcieDevice>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.PcieDevice>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListPcieDevicesApiResponseData"))
}

func (p *OneOfListPcieDevicesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<clustermgmt.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListPcieDevicesApiResponseData")
}

type OneOfListAhvVirtualGpuProfilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []VirtualGpuProfile    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListAhvVirtualGpuProfilesApiResponseData() *OneOfListAhvVirtualGpuProfilesApiResponseData {
	p := new(OneOfListAhvVirtualGpuProfilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAhvVirtualGpuProfilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAhvVirtualGpuProfilesApiResponseData is nil"))
	}
	switch v.(type) {
	case []VirtualGpuProfile:
		p.oneOfType2001 = v.([]VirtualGpuProfile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAhvVirtualGpuProfilesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAhvVirtualGpuProfilesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]VirtualGpuProfile)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.VirtualGpuProfile" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import1.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import1.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new([]VirtualGpuProfile)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "clustermgmt.v4.ahv.config.VirtualGpuProfile" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAhvVirtualGpuProfilesApiResponseData"))
}

func (p *OneOfListAhvVirtualGpuProfilesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<clustermgmt.v4.ahv.config.VirtualGpuProfile>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAhvVirtualGpuProfilesApiResponseData")
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
