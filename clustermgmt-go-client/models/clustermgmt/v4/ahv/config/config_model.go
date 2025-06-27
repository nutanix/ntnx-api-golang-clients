/*
 * Generated file models/clustermgmt/v4/ahv/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Cluster Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
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
REST response for all response codes in API path /clustermgmt/v4.1/ahv/config/pcie-devices Get operation
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
	*p = ListPcieDevicesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListPcieDevicesApiResponse() *ListPcieDevicesApiResponse {
	p := new(ListPcieDevicesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.ListPcieDevicesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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
	*p = PcieDevice(*known)

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
	delete(allFields, "ownerVmExtId")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPcieDevice() *PcieDevice {
	p := new(PcieDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PcieDevice"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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
	*p = PcieDeviceConfiguration(*known)

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
	p.UnknownFields_ = allFields

	return nil
}

func NewPcieDeviceConfiguration() *PcieDeviceConfiguration {
	p := new(PcieDeviceConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.ahv.config.PcieDeviceConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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

type OneOfListPcieDevicesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []PcieDevice           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	if "List<clustermgmt.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListPcieDevicesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]PcieDevice)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "clustermgmt.v4.ahv.config.PcieDevice" == *((*vOneOfType2001)[0].ObjectType_) {
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
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListPcieDevicesApiResponseData"))
}

func (p *OneOfListPcieDevicesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<clustermgmt.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListPcieDevicesApiResponseData")
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
