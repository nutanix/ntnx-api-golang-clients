/*
 * Generated file models/vmm/v4/ahv/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Acropolis hypervisor based Virtual Machines, Disks, Nics etc
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import5 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/dataprotection/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/*
The disk migration plan defines how disks are migrated.
*/
type ADSFDiskMigrationPlan struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	StorageContainer *VmDiskContainerReference `json:"storageContainer,omitempty"`
	/*
	  The disks included in the migration plan.
	*/
	VmDisks []MigrateDiskReference `json:"vmDisks,omitempty"`
}

func (p *ADSFDiskMigrationPlan) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ADSFDiskMigrationPlan

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

func (p *ADSFDiskMigrationPlan) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ADSFDiskMigrationPlan
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ADSFDiskMigrationPlan(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "storageContainer")
	delete(allFields, "vmDisks")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewADSFDiskMigrationPlan() *ADSFDiskMigrationPlan {
	p := new(ADSFDiskMigrationPlan)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFDiskMigrationPlan"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Storage configuration for VM.
*/
type ADSFVmStorageConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the virtual disk is pinned to the hot tier or not.
	*/
	IsFlashModeEnabled *bool `json:"isFlashModeEnabled,omitempty"`

	QosConfig *QosConfig `json:"qosConfig,omitempty"`
}

func (p *ADSFVmStorageConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ADSFVmStorageConfig

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

func (p *ADSFVmStorageConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ADSFVmStorageConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ADSFVmStorageConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isFlashModeEnabled")
	delete(allFields, "qosConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewADSFVmStorageConfig() *ADSFVmStorageConfig {
	p := new(ADSFVmStorageConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFVmStorageConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an ADSF Volume Group.
*/
type ADSFVolumeGroupReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of an ADSF volume group. It should be of type UUID.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId"`
}

func (p *ADSFVolumeGroupReference) MarshalJSON() ([]byte, error) {
	type ADSFVolumeGroupReferenceProxy ADSFVolumeGroupReference

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ADSFVolumeGroupReferenceProxy
		VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	}{
		ADSFVolumeGroupReferenceProxy: (*ADSFVolumeGroupReferenceProxy)(p),
		VolumeGroupExtId:              p.VolumeGroupExtId,
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

func (p *ADSFVolumeGroupReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ADSFVolumeGroupReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ADSFVolumeGroupReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewADSFVolumeGroupReference() *ADSFVolumeGroupReference {
	p := new(ADSFVolumeGroupReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ADSFVolumeGroupReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines how all the disks of a VM are migrated to the same target storage container.
*/
type AllDisksMigrationPlan struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	StorageContainer *VmDiskContainerReference `json:"storageContainer,omitempty"`
}

func (p *AllDisksMigrationPlan) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AllDisksMigrationPlan

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

func (p *AllDisksMigrationPlan) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AllDisksMigrationPlan
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AllDisksMigrationPlan(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "storageContainer")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAllDisksMigrationPlan() *AllDisksMigrationPlan {
	p := new(AllDisksMigrationPlan)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AllDisksMigrationPlan"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Advanced Processor Compatibility configuration for the VM. Enabling this retains the CPU model for the VM across power cycles and migrations.
*/
type ApcConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CpuModel *CpuModelReference `json:"cpuModel,omitempty"`
	/*
	  If enabled, the selected CPU model will be retained across live and cold migrations of the VM.
	*/
	IsApcEnabled *bool `json:"isApcEnabled,omitempty"`
}

func (p *ApcConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApcConfig

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

func (p *ApcConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApcConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ApcConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpuModel")
	delete(allFields, "isApcEnabled")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApcConfig() *ApcConfig {
	p := new(ApcConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ApcConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/assign-ip Post operation
*/
type AssignIpApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssignIpApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *AssignIpApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssignIpApiResponse

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

func (p *AssignIpApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssignIpApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AssignIpApiResponse(*known)

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

func NewAssignIpApiResponse() *AssignIpApiResponse {
	p := new(AssignIpApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssignIpApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssignIpApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssignIpApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssignIpApiResponseData()
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
Assigns a specific IP address to a network device attached to a managed network.
*/
type AssignIpParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import4.IPv4Address `json:"ipAddress,omitempty"`
}

func (p *AssignIpParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssignIpParams

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

func (p *AssignIpParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssignIpParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AssignIpParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAssignIpParams() *AssignIpParams {
	p := new(AssignIpParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssignIpParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/assign-owner Post operation
*/
type AssignVmOwnerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssignVmOwnerApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *AssignVmOwnerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssignVmOwnerApiResponse

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

func (p *AssignVmOwnerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssignVmOwnerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AssignVmOwnerApiResponse(*known)

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

func NewAssignVmOwnerApiResponse() *AssignVmOwnerApiResponse {
	p := new(AssignVmOwnerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssignVmOwnerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssignVmOwnerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssignVmOwnerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssignVmOwnerApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/associate-categories Post operation
*/
type AssociateCategoriesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssociateCategoriesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *AssociateCategoriesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssociateCategoriesApiResponse

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

func (p *AssociateCategoriesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssociateCategoriesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AssociateCategoriesApiResponse(*known)

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

func NewAssociateCategoriesApiResponse() *AssociateCategoriesApiResponse {
	p := new(AssociateCategoriesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssociateCategoriesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssociateCategoriesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssociateCategoriesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssociateCategoriesApiResponseData()
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
List of categories to be associated to the VM.
*/
type AssociateVmCategoriesParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Categories []CategoryReference `json:"categories,omitempty"`
}

func (p *AssociateVmCategoriesParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssociateVmCategoriesParams

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

func (p *AssociateVmCategoriesParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssociateVmCategoriesParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AssociateVmCategoriesParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAssociateVmCategoriesParams() *AssociateVmCategoriesParams {
	p := new(AssociateVmCategoriesParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AssociateVmCategoriesParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an availability zone.
*/
type AvailabilityZoneReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of an availability zone type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *AvailabilityZoneReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AvailabilityZoneReference

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

func (p *AvailabilityZoneReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AvailabilityZoneReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AvailabilityZoneReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAvailabilityZoneReference() *AvailabilityZoneReference {
	p := new(AvailabilityZoneReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.AvailabilityZoneReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The disk that will be used to boot the VM.
*/
type BootDeviceDisk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
}

func (p *BootDeviceDisk) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BootDeviceDisk

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

func (p *BootDeviceDisk) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BootDeviceDisk
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BootDeviceDisk(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBootDeviceDisk() *BootDeviceDisk {
	p := new(BootDeviceDisk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.BootDeviceDisk"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The NIC that will be used to boot the VM.
*/
type BootDeviceNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Mac address of the NIC that will be used to boot the VM.
	*/
	MacAddress *string `json:"macAddress,omitempty"`
}

func (p *BootDeviceNic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BootDeviceNic

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

func (p *BootDeviceNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BootDeviceNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BootDeviceNic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "macAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBootDeviceNic() *BootDeviceNic {
	p := new(BootDeviceNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.BootDeviceNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of boot device. Valid values are CDROM, NIC, and DISK.
*/
type BootDeviceType int

const (
	BOOTDEVICETYPE_UNKNOWN  BootDeviceType = 0
	BOOTDEVICETYPE_REDACTED BootDeviceType = 1
	BOOTDEVICETYPE_CDROM    BootDeviceType = 2
	BOOTDEVICETYPE_DISK     BootDeviceType = 3
	BOOTDEVICETYPE_NETWORK  BootDeviceType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BootDeviceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BootDeviceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BootDeviceType) index(name string) BootDeviceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	for idx := range names {
		if names[idx] == name {
			return BootDeviceType(idx)
		}
	}
	return BOOTDEVICETYPE_UNKNOWN
}

func (e *BootDeviceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BootDeviceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BootDeviceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BootDeviceType) Ref() *BootDeviceType {
	return &e
}

/*
Reference to a category.
*/
type CategoryReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM category of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *CategoryReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CategoryReference

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

func (p *CategoryReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CategoryReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CategoryReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCategoryReference() *CategoryReference {
	p := new(CategoryReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CategoryReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual Machine CD-ROM.
*/
type CdRom struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *VmDisk `json:"backingInfo,omitempty"`

	DiskAddress *CdRomAddress `json:"diskAddress,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IsoType *IsoType `json:"isoType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CdRom) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CdRom

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

func (p *CdRom) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CdRom
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CdRom(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backingInfo")
	delete(allFields, "diskAddress")
	delete(allFields, "extId")
	delete(allFields, "isoType")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCdRom() *CdRom {
	p := new(CdRom)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CdRom"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual Machine disk (VM disk).
*/
type CdRomAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BusType *CdRomBusType `json:"busType,omitempty"`
	/*
	  Device index on the bus. This field is ignored unless the bus details are specified.
	*/
	Index *int `json:"index,omitempty"`
}

func (p *CdRomAddress) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CdRomAddress

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

func (p *CdRomAddress) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CdRomAddress
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CdRomAddress(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "busType")
	delete(allFields, "index")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCdRomAddress() *CdRomAddress {
	p := new(CdRomAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CdRomAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Bus type for the device. The acceptable values are: IDE or SATA.
*/
type CdRomBusType int

const (
	CDROMBUSTYPE_UNKNOWN  CdRomBusType = 0
	CDROMBUSTYPE_REDACTED CdRomBusType = 1
	CDROMBUSTYPE_IDE      CdRomBusType = 2
	CDROMBUSTYPE_SATA     CdRomBusType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CdRomBusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IDE",
		"SATA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CdRomBusType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IDE",
		"SATA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CdRomBusType) index(name string) CdRomBusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IDE",
		"SATA",
	}
	for idx := range names {
		if names[idx] == name {
			return CdRomBusType(idx)
		}
	}
	return CDROMBUSTYPE_UNKNOWN
}

func (e *CdRomBusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CdRomBusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CdRomBusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CdRomBusType) Ref() *CdRomBusType {
	return &e
}

/*
Virtual Machine CD-ROM.
*/
type CdRomInsertParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *VmDisk `json:"backingInfo,omitempty"`
}

func (p *CdRomInsertParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CdRomInsertParams

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

func (p *CdRomInsertParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CdRomInsertParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CdRomInsertParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backingInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCdRomInsertParams() *CdRomInsertParams {
	p := new(CdRomInsertParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CdRomInsertParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM clone override specification.
*/
type CloneOverrideParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BootConfigItemDiscriminator_ *string `json:"$bootConfigItemDiscriminator,omitempty"`
	/*
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootConfig *OneOfCloneOverrideParamsBootConfig `json:"bootConfig,omitempty"`

	GuestCustomization *GuestCustomizationParams `json:"guestCustomization,omitempty"`
	/*
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  NICs attached to the VM.
	*/
	Nics []Nic `json:"nics,omitempty"`
	/*
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/*
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/*
	  Number of cores per socket.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
}

func (p *CloneOverrideParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CloneOverrideParams

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

func (p *CloneOverrideParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CloneOverrideParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CloneOverrideParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$bootConfigItemDiscriminator")
	delete(allFields, "bootConfig")
	delete(allFields, "guestCustomization")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "name")
	delete(allFields, "nics")
	delete(allFields, "numCoresPerSocket")
	delete(allFields, "numSockets")
	delete(allFields, "numThreadsPerCore")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCloneOverrideParams() *CloneOverrideParams {
	p := new(CloneOverrideParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloneOverrideParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloneOverrideParams) GetBootConfig() interface{} {
	if nil == p.BootConfig {
		return nil
	}
	return p.BootConfig.GetValue()
}

func (p *CloneOverrideParams) SetBootConfig(v interface{}) error {
	if nil == p.BootConfig {
		p.BootConfig = NewOneOfCloneOverrideParamsBootConfig()
	}
	e := p.BootConfig.SetValue(v)
	if nil == e {
		if nil == p.BootConfigItemDiscriminator_ {
			p.BootConfigItemDiscriminator_ = new(string)
		}
		*p.BootConfigItemDiscriminator_ = *p.BootConfig.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/clone Post operation
*/
type CloneVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCloneVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CloneVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CloneVmApiResponse

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

func (p *CloneVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CloneVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CloneVmApiResponse(*known)

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

func NewCloneVmApiResponse() *CloneVmApiResponse {
	p := new(CloneVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloneVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloneVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CloneVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCloneVmApiResponseData()
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
If this field is set, the guest will be customized using cloud-init. Either user_data or custom_key_values should be provided. If custom_key_ves are provided then the user data will be generated using these key-value pairs.
*/
type CloudInit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	CloudInitScriptItemDiscriminator_ *string `json:"$cloudInitScriptItemDiscriminator,omitempty"`
	/*
	  The script to use for cloud-init.
	*/
	CloudInitScript *OneOfCloudInitCloudInitScript `json:"cloudInitScript,omitempty"`

	DatasourceType *CloudInitDataSourceType `json:"datasourceType,omitempty"`
	/*
	  The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.
	*/
	Metadata *string `json:"metadata,omitempty"`
}

func (p *CloudInit) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CloudInit

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

func (p *CloudInit) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CloudInit
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CloudInit(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$cloudInitScriptItemDiscriminator")
	delete(allFields, "cloudInitScript")
	delete(allFields, "datasourceType")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCloudInit() *CloudInit {
	p := new(CloudInit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloudInit"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloudInit) GetCloudInitScript() interface{} {
	if nil == p.CloudInitScript {
		return nil
	}
	return p.CloudInitScript.GetValue()
}

func (p *CloudInit) SetCloudInitScript(v interface{}) error {
	if nil == p.CloudInitScript {
		p.CloudInitScript = NewOneOfCloudInitCloudInitScript()
	}
	e := p.CloudInitScript.SetValue(v)
	if nil == e {
		if nil == p.CloudInitScriptItemDiscriminator_ {
			p.CloudInitScriptItemDiscriminator_ = new(string)
		}
		*p.CloudInitScriptItemDiscriminator_ = *p.CloudInitScript.Discriminator
	}
	return e
}

/*
Type of datasource. Default: CONFIG_DRIVE_V2
*/
type CloudInitDataSourceType int

const (
	CLOUDINITDATASOURCETYPE_UNKNOWN         CloudInitDataSourceType = 0
	CLOUDINITDATASOURCETYPE_REDACTED        CloudInitDataSourceType = 1
	CLOUDINITDATASOURCETYPE_CONFIG_DRIVE_V2 CloudInitDataSourceType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CloudInitDataSourceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CloudInitDataSourceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CloudInitDataSourceType) index(name string) CloudInitDataSourceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudInitDataSourceType(idx)
		}
	}
	return CLOUDINITDATASOURCETYPE_UNKNOWN
}

func (e *CloudInitDataSourceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudInitDataSourceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudInitDataSourceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudInitDataSourceType) Ref() *CloudInitDataSourceType {
	return &e
}

/*
Reference to a cluster.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of a cluster type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *ClusterReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterReference

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

func (p *ClusterReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterReference() *ClusterReference {
	p := new(ClusterReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The list of additional CPU features to be enabled.
*/
type CpuFeature int

const (
	CPUFEATURE_UNKNOWN                 CpuFeature = 0
	CPUFEATURE_REDACTED                CpuFeature = 1
	CPUFEATURE_HARDWARE_VIRTUALIZATION CpuFeature = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CpuFeature) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HARDWARE_VIRTUALIZATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CpuFeature) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HARDWARE_VIRTUALIZATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CpuFeature) index(name string) CpuFeature {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HARDWARE_VIRTUALIZATION",
	}
	for idx := range names {
		if names[idx] == name {
			return CpuFeature(idx)
		}
	}
	return CPUFEATURE_UNKNOWN
}

func (e *CpuFeature) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CpuFeature:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CpuFeature) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CpuFeature) Ref() *CpuFeature {
	return &e
}

/*
CPU model associated with the VM if Advanced Processor Compatibility(APC) is enabled. If APC is enabled and no CPU model is explicitly set, a default baseline CPU model is picked by the system. See the APC documentation for more information
*/
type CpuModelReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of the CPU model associated with the VM.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Name of the CPU model associated with the VM.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *CpuModelReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CpuModelReference

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

func (p *CpuModelReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CpuModelReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CpuModelReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCpuModelReference() *CpuModelReference {
	p := new(CpuModelReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CpuModelReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms Post operation
*/
type CreateCdRomApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCdRomApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateCdRomApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateCdRomApiResponse

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

func (p *CreateCdRomApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateCdRomApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateCdRomApiResponse(*known)

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

func NewCreateCdRomApiResponse() *CreateCdRomApiResponse {
	p := new(CreateCdRomApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateCdRomApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCdRomApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCdRomApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCdRomApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/disks Post operation
*/
type CreateDiskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDiskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateDiskApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateDiskApiResponse

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

func (p *CreateDiskApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateDiskApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateDiskApiResponse(*known)

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

func NewCreateDiskApiResponse() *CreateDiskApiResponse {
	p := new(CreateDiskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateDiskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDiskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDiskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDiskApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/gpus Post operation
*/
type CreateGpuApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateGpuApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateGpuApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateGpuApiResponse

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

func (p *CreateGpuApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateGpuApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateGpuApiResponse(*known)

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

func NewCreateGpuApiResponse() *CreateGpuApiResponse {
	p := new(CreateGpuApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateGpuApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateGpuApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateGpuApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateGpuApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics Post operation
*/
type CreateNicApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNicApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateNicApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateNicApiResponse

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

func (p *CreateNicApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateNicApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateNicApiResponse(*known)

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

func NewCreateNicApiResponse() *CreateNicApiResponse {
	p := new(CreateNicApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateNicApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNicApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNicApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNicApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/pcie-devices Post operation
*/
type CreatePcieDeviceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreatePcieDeviceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreatePcieDeviceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreatePcieDeviceApiResponse

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

func (p *CreatePcieDeviceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreatePcieDeviceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreatePcieDeviceApiResponse(*known)

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

func NewCreatePcieDeviceApiResponse() *CreatePcieDeviceApiResponse {
	p := new(CreatePcieDeviceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreatePcieDeviceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreatePcieDeviceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreatePcieDeviceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreatePcieDeviceApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/serial-ports Post operation
*/
type CreateSerialPortApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSerialPortApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateSerialPortApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateSerialPortApiResponse

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

func (p *CreateSerialPortApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateSerialPortApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateSerialPortApiResponse(*known)

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

func NewCreateSerialPortApiResponse() *CreateSerialPortApiResponse {
	p := new(CreateSerialPortApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateSerialPortApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateSerialPortApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateSerialPortApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateSerialPortApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vm-recovery-points Post operation
*/
type CreateVMRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVMRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateVMRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateVMRecoveryPointApiResponse

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

func (p *CreateVMRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateVMRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateVMRecoveryPointApiResponse(*known)

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

func NewCreateVMRecoveryPointApiResponse() *CreateVMRecoveryPointApiResponse {
	p := new(CreateVMRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateVMRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateVMRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateVMRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateVMRecoveryPointApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms Post operation
*/
type CreateVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateVmApiResponse

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

func (p *CreateVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateVmApiResponse(*known)

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

func NewCreateVmApiResponse() *CreateVmApiResponse {
	p := new(CreateVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CreateVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateVmApiResponseData()
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
Sign in credentials for the server.
*/
type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password for the server.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Username for the server.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *Credential) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Credential

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

func (p *Credential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Credential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Credential(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Credential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/migrate Post operation
*/
type CrossClusterMigrateVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCrossClusterMigrateVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CrossClusterMigrateVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CrossClusterMigrateVmApiResponse

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

func (p *CrossClusterMigrateVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CrossClusterMigrateVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CrossClusterMigrateVmApiResponse(*known)

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

func NewCrossClusterMigrateVmApiResponse() *CrossClusterMigrateVmApiResponse {
	p := new(CrossClusterMigrateVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CrossClusterMigrateVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CrossClusterMigrateVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CrossClusterMigrateVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCrossClusterMigrateVmApiResponseData()
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
A collection of key/value pairs.
*/
type CustomKeyValues struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the individual KeyValuePair elements.
	*/
	KeyValuePairs []import4.KVPair `json:"keyValuePairs,omitempty"`
}

func (p *CustomKeyValues) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CustomKeyValues

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

func (p *CustomKeyValues) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomKeyValues
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CustomKeyValues(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "keyValuePairs")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCustomKeyValues() *CustomKeyValues {
	p := new(CustomKeyValues)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CustomKeyValues"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/customize-guest Post operation
*/
type CustomizeGuestVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCustomizeGuestVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CustomizeGuestVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CustomizeGuestVmApiResponse

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

func (p *CustomizeGuestVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomizeGuestVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CustomizeGuestVmApiResponse(*known)

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

func NewCustomizeGuestVmApiResponse() *CustomizeGuestVmApiResponse {
	p := new(CustomizeGuestVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CustomizeGuestVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CustomizeGuestVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CustomizeGuestVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCustomizeGuestVmApiResponseData()
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
A reference to a disk or image that contains the contents of a disk.
*/
type DataSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ReferenceItemDiscriminator_ *string `json:"$referenceItemDiscriminator,omitempty"`
	/*
	  A reference to a disk or image that contains the contents of a disk.
	*/
	Reference *OneOfDataSourceReference `json:"reference,omitempty"`
}

func (p *DataSource) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DataSource

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

func (p *DataSource) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DataSource
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DataSource(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$referenceItemDiscriminator")
	delete(allFields, "reference")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDataSource() *DataSource {
	p := new(DataSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DataSource"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DataSource) GetReference() interface{} {
	if nil == p.Reference {
		return nil
	}
	return p.Reference.GetValue()
}

func (p *DataSource) SetReference(v interface{}) error {
	if nil == p.Reference {
		p.Reference = NewOneOfDataSourceReference()
	}
	e := p.Reference.SetValue(v)
	if nil == e {
		if nil == p.ReferenceItemDiscriminator_ {
			p.ReferenceItemDiscriminator_ = new(string)
		}
		*p.ReferenceItemDiscriminator_ = *p.Reference.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms/{extId} Delete operation
*/
type DeleteCdRomApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteCdRomApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteCdRomApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteCdRomApiResponse

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

func (p *DeleteCdRomApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteCdRomApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteCdRomApiResponse(*known)

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

func NewDeleteCdRomApiResponse() *DeleteCdRomApiResponse {
	p := new(DeleteCdRomApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteCdRomApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteCdRomApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteCdRomApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteCdRomApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/disks/{extId} Delete operation
*/
type DeleteDiskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDiskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteDiskApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteDiskApiResponse

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

func (p *DeleteDiskApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteDiskApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteDiskApiResponse(*known)

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

func NewDeleteDiskApiResponse() *DeleteDiskApiResponse {
	p := new(DeleteDiskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteDiskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDiskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDiskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDiskApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/gpus/{extId} Delete operation
*/
type DeleteGpuApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteGpuApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteGpuApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteGpuApiResponse

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

func (p *DeleteGpuApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteGpuApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteGpuApiResponse(*known)

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

func NewDeleteGpuApiResponse() *DeleteGpuApiResponse {
	p := new(DeleteGpuApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteGpuApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteGpuApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteGpuApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteGpuApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId} Delete operation
*/
type DeleteNicApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteNicApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteNicApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteNicApiResponse

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

func (p *DeleteNicApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteNicApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteNicApiResponse(*known)

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

func NewDeleteNicApiResponse() *DeleteNicApiResponse {
	p := new(DeleteNicApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteNicApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteNicApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteNicApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteNicApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/pcie-devices/{extId} Delete operation
*/
type DeletePcieDeviceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeletePcieDeviceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeletePcieDeviceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeletePcieDeviceApiResponse

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

func (p *DeletePcieDeviceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeletePcieDeviceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeletePcieDeviceApiResponse(*known)

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

func NewDeletePcieDeviceApiResponse() *DeletePcieDeviceApiResponse {
	p := new(DeletePcieDeviceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeletePcieDeviceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeletePcieDeviceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeletePcieDeviceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeletePcieDeviceApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Delete operation
*/
type DeleteSerialPortApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSerialPortApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteSerialPortApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteSerialPortApiResponse

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

func (p *DeleteSerialPortApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteSerialPortApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteSerialPortApiResponse(*known)

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

func NewDeleteSerialPortApiResponse() *DeleteSerialPortApiResponse {
	p := new(DeleteSerialPortApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteSerialPortApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSerialPortApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSerialPortApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSerialPortApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId} Delete operation
*/
type DeleteVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteVmApiResponse

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

func (p *DeleteVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteVmApiResponse(*known)

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

func NewDeleteVmApiResponse() *DeleteVmApiResponse {
	p := new(DeleteVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vm-recovery-points/{extId} Delete operation
*/
type DeleteVmRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVmRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteVmRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteVmRecoveryPointApiResponse

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

func (p *DeleteVmRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteVmRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteVmRecoveryPointApiResponse(*known)

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

func NewDeleteVmRecoveryPointApiResponse() *DeleteVmRecoveryPointApiResponse {
	p := new(DeleteVmRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DeleteVmRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVmRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVmRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVmRecoveryPointApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/disassociate-categories Post operation
*/
type DisassociateCategoriesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDisassociateCategoriesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DisassociateCategoriesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DisassociateCategoriesApiResponse

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

func (p *DisassociateCategoriesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DisassociateCategoriesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DisassociateCategoriesApiResponse(*known)

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

func NewDisassociateCategoriesApiResponse() *DisassociateCategoriesApiResponse {
	p := new(DisassociateCategoriesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DisassociateCategoriesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DisassociateCategoriesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DisassociateCategoriesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDisassociateCategoriesApiResponseData()
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
List of categories to be disassociated from the VM.
*/
type DisassociateVmCategoriesParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Categories []CategoryReference `json:"categories,omitempty"`
}

func (p *DisassociateVmCategoriesParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DisassociateVmCategoriesParams

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

func (p *DisassociateVmCategoriesParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DisassociateVmCategoriesParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DisassociateVmCategoriesParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDisassociateVmCategoriesParams() *DisassociateVmCategoriesParams {
	p := new(DisassociateVmCategoriesParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DisassociateVmCategoriesParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual Machine disk (VM disk).
*/
type Disk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BackingInfoItemDiscriminator_ *string `json:"$backingInfoItemDiscriminator,omitempty"`
	/*
	  Supporting storage to create virtual disk on.
	*/
	BackingInfo *OneOfDiskBackingInfo `json:"backingInfo,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Disk) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Disk

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

func (p *Disk) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Disk
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Disk(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$backingInfoItemDiscriminator")
	delete(allFields, "backingInfo")
	delete(allFields, "diskAddress")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDisk() *Disk {
	p := new(Disk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Disk"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Disk) GetBackingInfo() interface{} {
	if nil == p.BackingInfo {
		return nil
	}
	return p.BackingInfo.GetValue()
}

func (p *Disk) SetBackingInfo(v interface{}) error {
	if nil == p.BackingInfo {
		p.BackingInfo = NewOneOfDiskBackingInfo()
	}
	e := p.BackingInfo.SetValue(v)
	if nil == e {
		if nil == p.BackingInfoItemDiscriminator_ {
			p.BackingInfoItemDiscriminator_ = new(string)
		}
		*p.BackingInfoItemDiscriminator_ = *p.BackingInfo.Discriminator
	}
	return e
}

/*
Disk address.
*/
type DiskAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BusType *DiskBusType `json:"busType,omitempty"`
	/*
	  Device index on the bus. This field is ignored unless the bus details are specified.
	*/
	Index *int `json:"index,omitempty"`
}

func (p *DiskAddress) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DiskAddress

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

func (p *DiskAddress) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiskAddress
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DiskAddress(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "busType")
	delete(allFields, "index")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDiskAddress() *DiskAddress {
	p := new(DiskAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DiskAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Bus type for the device. The acceptable values are: SCSI, IDE, PCI, SATA, SPAPR (only PPC).
*/
type DiskBusType int

const (
	DISKBUSTYPE_UNKNOWN  DiskBusType = 0
	DISKBUSTYPE_REDACTED DiskBusType = 1
	DISKBUSTYPE_SCSI     DiskBusType = 2
	DISKBUSTYPE_IDE      DiskBusType = 3
	DISKBUSTYPE_PCI      DiskBusType = 4
	DISKBUSTYPE_SATA     DiskBusType = 5
	DISKBUSTYPE_SPAPR    DiskBusType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DiskBusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DiskBusType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DiskBusType) index(name string) DiskBusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	for idx := range names {
		if names[idx] == name {
			return DiskBusType(idx)
		}
	}
	return DISKBUSTYPE_UNKNOWN
}

func (e *DiskBusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DiskBusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DiskBusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DiskBusType) Ref() *DiskBusType {
	return &e
}

/*
Specifies the VmDisks of a VM for migration and the migration plan for them. If all the disks of a VM need to be migrated to the same storage container, only a single migration plan with only the external ID of the destination storage container is needed.  If the disks are being migrated to different containers, one plan per disk needs to be specified.
*/
type DiskMigrationParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	MigrateDisksItemDiscriminator_ *string `json:"$migrateDisksItemDiscriminator,omitempty"`
	/*
	  Specifies the VmDisks of a VM for migration and the migration plan for them. If all the disks of a VM need to be migrated to the same storage container, only a single migration plan with only the external ID of the destination storage container is needed.  If the disks are being migrated to different containers, one plan per disk needs to be specified.
	*/
	MigrateDisks *OneOfDiskMigrationParamsMigrateDisks `json:"migrateDisks,omitempty"`
}

func (p *DiskMigrationParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DiskMigrationParams

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

func (p *DiskMigrationParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiskMigrationParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DiskMigrationParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$migrateDisksItemDiscriminator")
	delete(allFields, "migrateDisks")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDiskMigrationParams() *DiskMigrationParams {
	p := new(DiskMigrationParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DiskMigrationParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DiskMigrationParams) GetMigrateDisks() interface{} {
	if nil == p.MigrateDisks {
		return nil
	}
	return p.MigrateDisks.GetValue()
}

func (p *DiskMigrationParams) SetMigrateDisks(v interface{}) error {
	if nil == p.MigrateDisks {
		p.MigrateDisks = NewOneOfDiskMigrationParamsMigrateDisks()
	}
	e := p.MigrateDisks.SetValue(v)
	if nil == e {
		if nil == p.MigrateDisksItemDiscriminator_ {
			p.MigrateDisksItemDiscriminator_ = new(string)
		}
		*p.MigrateDisksItemDiscriminator_ = *p.MigrateDisks.Discriminator
	}
	return e
}

/*
Defines a DP Offload NIC.
*/
type DpOffloadNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DpOffloadProfileReference *NicProfileReference `json:"dpOffloadProfileReference"`

	HostPcieDeviceReference *HostPcieDeviceReference `json:"hostPcieDeviceReference,omitempty"`
	/*
	  Indicates whether the NIC is connected or not. Default is True.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`
}

func (p *DpOffloadNic) MarshalJSON() ([]byte, error) {
	type DpOffloadNicProxy DpOffloadNic

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DpOffloadNicProxy
		DpOffloadProfileReference *NicProfileReference `json:"dpOffloadProfileReference,omitempty"`
	}{
		DpOffloadNicProxy:         (*DpOffloadNicProxy)(p),
		DpOffloadProfileReference: p.DpOffloadProfileReference,
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

func (p *DpOffloadNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DpOffloadNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DpOffloadNic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dpOffloadProfileReference")
	delete(allFields, "hostPcieDeviceReference")
	delete(allFields, "isConnected")
	delete(allFields, "macAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDpOffloadNic() *DpOffloadNic {
	p := new(DpOffloadNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DpOffloadNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsConnected = new(bool)
	*p.IsConnected = true

	return p
}

/*
Networking information object for a DP Offload NIC.
*/
type DpOffloadNicNetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4Config *Ipv4Config `json:"ipv4Config,omitempty"`

	Ipv4Info *Ipv4Info `json:"ipv4Info,omitempty"`

	Ipv6Info *Ipv6Info `json:"ipv6Info,omitempty"`
	/*
	  Indicates whether an unknown unicast traffic is forwarded to this NIC or not. This is applicable only for the NICs on the overlay subnets.
	*/
	ShouldAllowUnknownMacs *bool `json:"shouldAllowUnknownMacs,omitempty"`

	Subnet *SubnetReference `json:"subnet,omitempty"`
	/*
	  List of networks to trunk if VLAN mode is marked as TRUNKED. If empty and VLAN mode is set to TRUNKED, all the VLANs are trunked.
	*/
	TrunkedVlans []int `json:"trunkedVlans,omitempty"`

	VlanMode *VlanMode `json:"vlanMode,omitempty"`
}

func (p *DpOffloadNicNetworkInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DpOffloadNicNetworkInfo

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

func (p *DpOffloadNicNetworkInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DpOffloadNicNetworkInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DpOffloadNicNetworkInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipv4Config")
	delete(allFields, "ipv4Info")
	delete(allFields, "ipv6Info")
	delete(allFields, "shouldAllowUnknownMacs")
	delete(allFields, "subnet")
	delete(allFields, "trunkedVlans")
	delete(allFields, "vlanMode")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDpOffloadNicNetworkInfo() *DpOffloadNicNetworkInfo {
	p := new(DpOffloadNicNetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.DpOffloadNicNetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/eject Post operation
*/
type EjectCdRomApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEjectCdRomApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *EjectCdRomApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EjectCdRomApiResponse

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

func (p *EjectCdRomApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EjectCdRomApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EjectCdRomApiResponse(*known)

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

func NewEjectCdRomApiResponse() *EjectCdRomApiResponse {
	p := new(EjectCdRomApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.EjectCdRomApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EjectCdRomApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EjectCdRomApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEjectCdRomApiResponseData()
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
Defines a NIC emulated by the hypervisor. This object is now deprecated. If both deprecated and new objects are present, the new object VirtualEthernetNic takes precedence.
*/
type EmulatedNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the NIC is connected or not. Default is True.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`

	Model *EmulatedNicModel `json:"model,omitempty"`
	/*
	  The number of Tx/Rx queue pairs for this NIC.
	*/
	NumQueues *int `json:"numQueues,omitempty"`
}

func (p *EmulatedNic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EmulatedNic

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

func (p *EmulatedNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EmulatedNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EmulatedNic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isConnected")
	delete(allFields, "macAddress")
	delete(allFields, "model")
	delete(allFields, "numQueues")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEmulatedNic() *EmulatedNic {
	p := new(EmulatedNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.EmulatedNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsConnected = new(bool)
	*p.IsConnected = true
	p.NumQueues = new(int)
	*p.NumQueues = 1

	return p
}

/*
Options for the NIC emulation.
*/
type EmulatedNicModel int

const (
	EMULATEDNICMODEL_UNKNOWN  EmulatedNicModel = 0
	EMULATEDNICMODEL_REDACTED EmulatedNicModel = 1
	EMULATEDNICMODEL_VIRTIO   EmulatedNicModel = 2
	EMULATEDNICMODEL_E1000    EmulatedNicModel = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EmulatedNicModel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EmulatedNicModel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EmulatedNicModel) index(name string) EmulatedNicModel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	for idx := range names {
		if names[idx] == name {
			return EmulatedNicModel(idx)
		}
	}
	return EMULATEDNICMODEL_UNKNOWN
}

func (e *EmulatedNicModel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EmulatedNicModel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EmulatedNicModel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EmulatedNicModel) Ref() *EmulatedNicModel {
	return &e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms/{extId} Get operation
*/
type GetCdRomApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCdRomApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetCdRomApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetCdRomApiResponse

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

func (p *GetCdRomApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetCdRomApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetCdRomApiResponse(*known)

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

func NewGetCdRomApiResponse() *GetCdRomApiResponse {
	p := new(GetCdRomApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetCdRomApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCdRomApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCdRomApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCdRomApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/disks/{extId} Get operation
*/
type GetDiskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDiskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDiskApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDiskApiResponse

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

func (p *GetDiskApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDiskApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetDiskApiResponse(*known)

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

func NewGetDiskApiResponse() *GetDiskApiResponse {
	p := new(GetDiskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetDiskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDiskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDiskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDiskApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/gpus/{extId} Get operation
*/
type GetGpuApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGpuApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetGpuApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetGpuApiResponse

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

func (p *GetGpuApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetGpuApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetGpuApiResponse(*known)

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

func NewGetGpuApiResponse() *GetGpuApiResponse {
	p := new(GetGpuApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetGpuApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGpuApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGpuApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGpuApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools Get operation
*/
type GetGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetGuestToolsApiResponse

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

func (p *GetGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetGuestToolsApiResponse(*known)

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

func NewGetGuestToolsApiResponse() *GetGuestToolsApiResponse {
	p := new(GetGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGuestToolsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId} Get operation
*/
type GetNicApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNicApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetNicApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetNicApiResponse

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

func (p *GetNicApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetNicApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetNicApiResponse(*known)

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

func NewGetNicApiResponse() *GetNicApiResponse {
	p := new(GetNicApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetNicApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNicApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNicApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNicApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/pcie-devices/{extId} Get operation
*/
type GetPcieDeviceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetPcieDeviceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetPcieDeviceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetPcieDeviceApiResponse

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

func (p *GetPcieDeviceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetPcieDeviceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetPcieDeviceApiResponse(*known)

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

func NewGetPcieDeviceApiResponse() *GetPcieDeviceApiResponse {
	p := new(GetPcieDeviceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetPcieDeviceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetPcieDeviceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetPcieDeviceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetPcieDeviceApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Get operation
*/
type GetSerialPortApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSerialPortApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetSerialPortApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetSerialPortApiResponse

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

func (p *GetSerialPortApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetSerialPortApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetSerialPortApiResponse(*known)

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

func NewGetSerialPortApiResponse() *GetSerialPortApiResponse {
	p := new(GetSerialPortApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetSerialPortApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSerialPortApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSerialPortApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSerialPortApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId} Get operation
*/
type GetVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmApiResponse

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

func (p *GetVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVmApiResponse(*known)

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

func NewGetVmApiResponse() *GetVmApiResponse {
	p := new(GetVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vm-recovery-points/{extId} Get operation
*/
type GetVmRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmRecoveryPointApiResponse

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

func (p *GetVmRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVmRecoveryPointApiResponse(*known)

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

func NewGetVmRecoveryPointApiResponse() *GetVmRecoveryPointApiResponse {
	p := new(GetVmRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GetVmRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmRecoveryPointApiResponseData()
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
Graphics resource information for the virtual machine.
*/
type Gpu struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The device Id of the GPU.
	*/
	DeviceId *int `json:"deviceId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Fraction of the physical GPU assigned.
	*/
	Fraction *int `json:"fraction,omitempty"`
	/*
	  GPU frame buffer size in bytes.
	*/
	FrameBufferSizeBytes *int64 `json:"frameBufferSizeBytes,omitempty"`
	/*
	  Last determined guest driver version.
	*/
	GuestDriverVersion *string `json:"guestDriverVersion,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	Mode *GpuMode `json:"mode,omitempty"`
	/*
	  Name of the GPU resource.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Number of supported virtual display heads.
	*/
	NumVirtualDisplayHeads *int `json:"numVirtualDisplayHeads,omitempty"`

	PciAddress *SBDF `json:"pciAddress,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vendor *GpuVendor `json:"vendor,omitempty"`
}

func (p *Gpu) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Gpu

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

func (p *Gpu) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Gpu
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Gpu(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "deviceId")
	delete(allFields, "extId")
	delete(allFields, "fraction")
	delete(allFields, "frameBufferSizeBytes")
	delete(allFields, "guestDriverVersion")
	delete(allFields, "links")
	delete(allFields, "mode")
	delete(allFields, "name")
	delete(allFields, "numVirtualDisplayHeads")
	delete(allFields, "pciAddress")
	delete(allFields, "tenantId")
	delete(allFields, "vendor")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGpu() *Gpu {
	p := new(Gpu)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Gpu"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The mode of this GPU.
*/
type GpuMode int

const (
	GPUMODE_UNKNOWN              GpuMode = 0
	GPUMODE_REDACTED             GpuMode = 1
	GPUMODE_PASSTHROUGH_GRAPHICS GpuMode = 2
	GPUMODE_PASSTHROUGH_COMPUTE  GpuMode = 3
	GPUMODE_VIRTUAL              GpuMode = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GpuMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GpuMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GpuMode) index(name string) GpuMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
	}
	for idx := range names {
		if names[idx] == name {
			return GpuMode(idx)
		}
	}
	return GPUMODE_UNKNOWN
}

func (e *GpuMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GpuMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GpuMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GpuMode) Ref() *GpuMode {
	return &e
}

/*
The vendor of the GPU.
*/
type GpuVendor int

const (
	GPUVENDOR_UNKNOWN  GpuVendor = 0
	GPUVENDOR_REDACTED GpuVendor = 1
	GPUVENDOR_NVIDIA   GpuVendor = 2
	GPUVENDOR_INTEL    GpuVendor = 3
	GPUVENDOR_AMD      GpuVendor = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GpuVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GpuVendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GpuVendor) index(name string) GpuVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	for idx := range names {
		if names[idx] == name {
			return GpuVendor(idx)
		}
	}
	return GPUVENDOR_UNKNOWN
}

func (e *GpuVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GpuVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GpuVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GpuVendor) Ref() *GpuVendor {
	return &e
}

/*
Stage a Sysprep or cloud-init configuration file to be used by the guest for the next boot. Note that the Sysprep command must be used to generalize the Windows VMs before triggering this API call.
*/
type GuestCustomizationParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`
	/*
	  The Nutanix Guest Tools customization settings.
	*/
	Config *OneOfGuestCustomizationParamsConfig `json:"config,omitempty"`
}

func (p *GuestCustomizationParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestCustomizationParams

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

func (p *GuestCustomizationParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestCustomizationParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestCustomizationParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configItemDiscriminator")
	delete(allFields, "config")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestCustomizationParams() *GuestCustomizationParams {
	p := new(GuestCustomizationParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestCustomizationParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GuestCustomizationParams) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *GuestCustomizationParams) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfGuestCustomizationParamsConfig()
	}
	e := p.Config.SetValue(v)
	if nil == e {
		if nil == p.ConfigItemDiscriminator_ {
			p.ConfigItemDiscriminator_ = new(string)
		}
		*p.ConfigItemDiscriminator_ = *p.Config.Discriminator
	}
	return e
}

/*
Nutanix Guest Tools shutdown mode options.
*/
type GuestPowerOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	GuestPowerStateTransitionConfig *GuestPowerStateTransitionConfig `json:"guestPowerStateTransitionConfig,omitempty"`
}

func (p *GuestPowerOptions) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestPowerOptions

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

func (p *GuestPowerOptions) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestPowerOptions
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestPowerOptions(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "guestPowerStateTransitionConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestPowerOptions() *GuestPowerOptions {
	p := new(GuestPowerOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestPowerOptions"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Additional configuration for Nutanix Gust Tools power state transition.
*/
type GuestPowerStateTransitionConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether to run the set script before the VM shutdowns/restarts.
	*/
	ShouldEnableScriptExec *bool `json:"shouldEnableScriptExec,omitempty"`
	/*
	  Indicates whether to abort VM shutdown/restart if the script fails.
	*/
	ShouldFailOnScriptFailure *bool `json:"shouldFailOnScriptFailure,omitempty"`
}

func (p *GuestPowerStateTransitionConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestPowerStateTransitionConfig

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

func (p *GuestPowerStateTransitionConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestPowerStateTransitionConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestPowerStateTransitionConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "shouldEnableScriptExec")
	delete(allFields, "shouldFailOnScriptFailure")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestPowerStateTransitionConfig() *GuestPowerStateTransitionConfig {
	p := new(GuestPowerStateTransitionConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestPowerStateTransitionConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The details about Nutanix Guest Tools for a VM.
*/
type GuestTools struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Version of Nutanix Guest Tools available on the cluster.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`
	/*
	  Version of the operating system on the VM.
	*/
	GuestOsVersion *string `json:"guestOsVersion,omitempty"`
	/*
	  Indicates whether the Nutanix Guest Tools communication over Serial Port is active or not.
	*/
	IsCommunicationActiveOverSerialPort *bool `json:"isCommunicationActiveOverSerialPort,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools is installed on the VM or not.
	*/
	IsInstalled *bool `json:"isInstalled,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools ISO is inserted or not.
	*/
	IsIsoInserted *bool `json:"isIsoInserted,omitempty"`
	/*
	  Indicates whether the communication from VM to CVM is active or not.
	*/
	IsReachable *bool `json:"isReachable,omitempty"`
	/*
	  Indicates whether the VM mobility drivers are installed on the VM or not.
	*/
	IsVmMobilityDriversInstalled *bool `json:"isVmMobilityDriversInstalled,omitempty"`
	/*
	  Indicates whether the VM is configured to take VSS snapshots through NGT or not.
	*/
	IsVssSnapshotCapable *bool `json:"isVssSnapshotCapable,omitempty"`
	/*
	  Version of Nutanix Guest Tools installed on the VM.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *GuestTools) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestTools

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

func (p *GuestTools) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestTools
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestTools(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableVersion")
	delete(allFields, "capabilities")
	delete(allFields, "guestOsVersion")
	delete(allFields, "isCommunicationActiveOverSerialPort")
	delete(allFields, "isEnabled")
	delete(allFields, "isInstalled")
	delete(allFields, "isIsoInserted")
	delete(allFields, "isReachable")
	delete(allFields, "isVmMobilityDriversInstalled")
	delete(allFields, "isVssSnapshotCapable")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestTools() *GuestTools {
	p := new(GuestTools)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestTools"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Argument for inserting a Nutanix Guest Tools ISO into an available slot.
*/
type GuestToolsInsertConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`
	/*
	  Indicates that the Nutanix Guest Tools are already installed on the guest VM, and the ISO is being inserted to update the configuration of these tools.
	*/
	IsConfigOnly *bool `json:"isConfigOnly,omitempty"`
}

func (p *GuestToolsInsertConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestToolsInsertConfig

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

func (p *GuestToolsInsertConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestToolsInsertConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestToolsInsertConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "isConfigOnly")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestToolsInsertConfig() *GuestToolsInsertConfig {
	p := new(GuestToolsInsertConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsInsertConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Arguments for installing Nutanix Guest Tools.
*/
type GuestToolsInstallConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NgtCapability `json:"capabilities,omitempty"`

	Credential *Credential `json:"credential,omitempty"`

	RebootPreference *RebootPreference `json:"rebootPreference,omitempty"`
}

func (p *GuestToolsInstallConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestToolsInstallConfig

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

func (p *GuestToolsInstallConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestToolsInstallConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestToolsInstallConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "credential")
	delete(allFields, "rebootPreference")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestToolsInstallConfig() *GuestToolsInstallConfig {
	p := new(GuestToolsInstallConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsInstallConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Arguments for upgrading Nutanix Guest Tools.
*/
type GuestToolsUpgradeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RebootPreference *RebootPreference `json:"rebootPreference,omitempty"`
}

func (p *GuestToolsUpgradeConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GuestToolsUpgradeConfig

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

func (p *GuestToolsUpgradeConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GuestToolsUpgradeConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GuestToolsUpgradeConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "rebootPreference")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGuestToolsUpgradeConfig() *GuestToolsUpgradeConfig {
	p := new(GuestToolsUpgradeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.GuestToolsUpgradeConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the host PCIe device.
*/
type HostPcieDeviceReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier of the host PCIe device. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *HostPcieDeviceReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostPcieDeviceReference

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

func (p *HostPcieDeviceReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostPcieDeviceReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HostPcieDeviceReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHostPcieDeviceReference() *HostPcieDeviceReference {
	p := new(HostPcieDeviceReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.HostPcieDeviceReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the host, the VM is running on.
*/
type HostReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of a host type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *HostReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostReference

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

func (p *HostReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HostReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Clone a disk from an image.
*/
type ImageReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an image of type UUID.
	*/
	ImageExtId *string `json:"imageExtId"`
}

func (p *ImageReference) MarshalJSON() ([]byte, error) {
	type ImageReferenceProxy ImageReference

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageReferenceProxy
		ImageExtId *string `json:"imageExtId,omitempty"`
	}{
		ImageReferenceProxy: (*ImageReferenceProxy)(p),
		ImageExtId:          p.ImageExtId,
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

func (p *ImageReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ImageReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "imageExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewImageReference() *ImageReference {
	p := new(ImageReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ImageReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms/{extId}/$actions/insert Post operation
*/
type InsertCdRomApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInsertCdRomApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *InsertCdRomApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InsertCdRomApiResponse

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

func (p *InsertCdRomApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InsertCdRomApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = InsertCdRomApiResponse(*known)

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

func NewInsertCdRomApiResponse() *InsertCdRomApiResponse {
	p := new(InsertCdRomApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InsertCdRomApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InsertCdRomApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InsertCdRomApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInsertCdRomApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools/$actions/insert-iso Post operation
*/
type InsertVmGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInsertVmGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *InsertVmGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InsertVmGuestToolsApiResponse

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

func (p *InsertVmGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InsertVmGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = InsertVmGuestToolsApiResponse(*known)

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

func NewInsertVmGuestToolsApiResponse() *InsertVmGuestToolsApiResponse {
	p := new(InsertVmGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InsertVmGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InsertVmGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InsertVmGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInsertVmGuestToolsApiResponseData()
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
Indicates whether the guest will be freshly installed using this unattend configuration, or this unattend configuration will be applied to a pre-prepared image. Default is 'PREPARED'.
*/
type InstallType int

const (
	INSTALLTYPE_UNKNOWN  InstallType = 0
	INSTALLTYPE_REDACTED InstallType = 1
	INSTALLTYPE_FRESH    InstallType = 2
	INSTALLTYPE_PREPARED InstallType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *InstallType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e InstallType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *InstallType) index(name string) InstallType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	for idx := range names {
		if names[idx] == name {
			return InstallType(idx)
		}
	}
	return INSTALLTYPE_UNKNOWN
}

func (e *InstallType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for InstallType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *InstallType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e InstallType) Ref() *InstallType {
	return &e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools/$actions/install Post operation
*/
type InstallVmGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInstallVmGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *InstallVmGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InstallVmGuestToolsApiResponse

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

func (p *InstallVmGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InstallVmGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = InstallVmGuestToolsApiResponse(*known)

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

func NewInstallVmGuestToolsApiResponse() *InstallVmGuestToolsApiResponse {
	p := new(InstallVmGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.InstallVmGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InstallVmGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InstallVmGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInstallVmGuestToolsApiResponseData()
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
The IP address configurations.
*/
type Ipv4Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import4.IPv4Address `json:"ipAddress,omitempty"`
	/*
	  Secondary IP addresses for the NIC.
	*/
	SecondaryIpAddressList []import4.IPv4Address `json:"secondaryIpAddressList,omitempty"`
	/*
	  If set to true (default value), an IP address must be assigned to the VM NIC - either the one explicitly specified by the user or allocated automatically by the IPAM service by not specifying the IP address. If false, then no IP assignment is required for this VM NIC.
	*/
	ShouldAssignIp *bool `json:"shouldAssignIp,omitempty"`
}

func (p *Ipv4Config) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Ipv4Config

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

func (p *Ipv4Config) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Ipv4Config
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Ipv4Config(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "secondaryIpAddressList")
	delete(allFields, "shouldAssignIp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIpv4Config() *Ipv4Config {
	p := new(Ipv4Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Ipv4Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The runtime IP address information of the NIC.
*/
type Ipv4Info struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The IP addresses as found on the guest VM for the NIC.
	*/
	LearnedIpAddresses []import4.IPv4Address `json:"learnedIpAddresses,omitempty"`
}

func (p *Ipv4Info) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Ipv4Info

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

func (p *Ipv4Info) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Ipv4Info
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Ipv4Info(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "learnedIpAddresses")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIpv4Info() *Ipv4Info {
	p := new(Ipv4Info)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Ipv4Info"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The runtime IPv6 address information of the NIC.
*/
type Ipv6Info struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The IPv6 addresses as found on the guest VM for the NIC.
	*/
	LearnedIpv6Addresses []import4.IPv6Address `json:"learnedIpv6Addresses,omitempty"`
}

func (p *Ipv6Info) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Ipv6Info

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

func (p *Ipv6Info) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Ipv6Info
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Ipv6Info(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "learnedIpv6Addresses")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIpv6Info() *Ipv6Info {
	p := new(Ipv6Info)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Ipv6Info"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of ISO image inserted in CD-ROM
*/
type IsoType int

const (
	ISOTYPE_UNKNOWN             IsoType = 0
	ISOTYPE_REDACTED            IsoType = 1
	ISOTYPE_OTHER               IsoType = 2
	ISOTYPE_GUEST_TOOLS         IsoType = 3
	ISOTYPE_GUEST_CUSTOMIZATION IsoType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IsoType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OTHER",
		"GUEST_TOOLS",
		"GUEST_CUSTOMIZATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e IsoType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OTHER",
		"GUEST_TOOLS",
		"GUEST_CUSTOMIZATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *IsoType) index(name string) IsoType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OTHER",
		"GUEST_TOOLS",
		"GUEST_CUSTOMIZATION",
	}
	for idx := range names {
		if names[idx] == name {
			return IsoType(idx)
		}
	}
	return ISOTYPE_UNKNOWN
}

func (e *IsoType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IsoType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IsoType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IsoType) Ref() *IsoType {
	return &e
}

/*
Legacy boot mode and its associated configuration.
*/
type LegacyBoot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BootDeviceItemDiscriminator_ *string `json:"$bootDeviceItemDiscriminator,omitempty"`
	/*
	  The boot device used for legacy boot of the VM.
	*/
	BootDevice *OneOfLegacyBootBootDevice `json:"bootDevice,omitempty"`
	/*
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootOrder []BootDeviceType `json:"bootOrder,omitempty"`
}

func (p *LegacyBoot) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LegacyBoot

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

func (p *LegacyBoot) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LegacyBoot
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LegacyBoot(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$bootDeviceItemDiscriminator")
	delete(allFields, "bootDevice")
	delete(allFields, "bootOrder")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLegacyBoot() *LegacyBoot {
	p := new(LegacyBoot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.LegacyBoot"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *LegacyBoot) GetBootDevice() interface{} {
	if nil == p.BootDevice {
		return nil
	}
	return p.BootDevice.GetValue()
}

func (p *LegacyBoot) SetBootDevice(v interface{}) error {
	if nil == p.BootDevice {
		p.BootDevice = NewOneOfLegacyBootBootDevice()
	}
	e := p.BootDevice.SetValue(v)
	if nil == e {
		if nil == p.BootDeviceItemDiscriminator_ {
			p.BootDeviceItemDiscriminator_ = new(string)
		}
		*p.BootDeviceItemDiscriminator_ = *p.BootDevice.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/cd-roms Get operation
*/
type ListCdRomsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCdRomsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListCdRomsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListCdRomsApiResponse

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

func (p *ListCdRomsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListCdRomsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListCdRomsApiResponse(*known)

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

func NewListCdRomsApiResponse() *ListCdRomsApiResponse {
	p := new(ListCdRomsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListCdRomsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCdRomsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCdRomsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCdRomsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/disks Get operation
*/
type ListDisksApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDisksApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDisksApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDisksApiResponse

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

func (p *ListDisksApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDisksApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListDisksApiResponse(*known)

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

func NewListDisksApiResponse() *ListDisksApiResponse {
	p := new(ListDisksApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListDisksApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDisksApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDisksApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDisksApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/gpus Get operation
*/
type ListGpusApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListGpusApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListGpusApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListGpusApiResponse

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

func (p *ListGpusApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListGpusApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListGpusApiResponse(*known)

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

func NewListGpusApiResponse() *ListGpusApiResponse {
	p := new(ListGpusApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListGpusApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListGpusApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListGpusApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListGpusApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics Get operation
*/
type ListNicsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNicsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListNicsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListNicsApiResponse

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

func (p *ListNicsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListNicsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListNicsApiResponse(*known)

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

func NewListNicsApiResponse() *ListNicsApiResponse {
	p := new(ListNicsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListNicsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNicsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNicsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNicsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/pcie-devices Get operation
*/
type ListPcieDevicesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListPcieDevicesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
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
	*p.ObjectType_ = "vmm.v4.ahv.config.ListPcieDevicesApiResponse"
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/serial-ports Get operation
*/
type ListSerialPortsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSerialPortsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListSerialPortsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListSerialPortsApiResponse

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

func (p *ListSerialPortsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListSerialPortsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListSerialPortsApiResponse(*known)

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

func NewListSerialPortsApiResponse() *ListSerialPortsApiResponse {
	p := new(ListSerialPortsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListSerialPortsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSerialPortsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSerialPortsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSerialPortsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vm-recovery-points Get operation
*/
type ListVmRecoveryPointsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmRecoveryPointsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmRecoveryPointsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmRecoveryPointsApiResponse

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

func (p *ListVmRecoveryPointsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmRecoveryPointsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmRecoveryPointsApiResponse(*known)

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

func NewListVmRecoveryPointsApiResponse() *ListVmRecoveryPointsApiResponse {
	p := new(ListVmRecoveryPointsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListVmRecoveryPointsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmRecoveryPointsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmRecoveryPointsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmRecoveryPointsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms Get operation
*/
type ListVmsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmsApiResponse

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

func (p *ListVmsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmsApiResponse(*known)

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

func NewListVmsApiResponse() *ListVmsApiResponse {
	p := new(ListVmsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ListVmsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmsApiResponseData()
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
Machine type for the VM. Machine type Q35 is required for secure boot and does not support IDE disks.
*/
type MachineType int

const (
	MACHINETYPE_UNKNOWN  MachineType = 0
	MACHINETYPE_REDACTED MachineType = 1
	MACHINETYPE_PC       MachineType = 2
	MACHINETYPE_PSERIES  MachineType = 3
	MACHINETYPE_Q35      MachineType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MachineType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MachineType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MachineType) index(name string) MachineType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	for idx := range names {
		if names[idx] == name {
			return MachineType(idx)
		}
	}
	return MACHINETYPE_UNKNOWN
}

func (e *MachineType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MachineType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MachineType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MachineType) Ref() *MachineType {
	return &e
}

/*
Reference to an existing VM disk that needs to be migrated.
*/
type MigrateDiskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM disk of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
}

func (p *MigrateDiskReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MigrateDiskReference

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

func (p *MigrateDiskReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateDiskReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrateDiskReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMigrateDiskReference() *MigrateDiskReference {
	p := new(MigrateDiskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateDiskReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/migrate Post operation
*/
type MigrateNicApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrateNicApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *MigrateNicApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MigrateNicApiResponse

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

func (p *MigrateNicApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateNicApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrateNicApiResponse(*known)

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

func NewMigrateNicApiResponse() *MigrateNicApiResponse {
	p := new(MigrateNicApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateNicApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrateNicApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrateNicApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrateNicApiResponseData()
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
Configurations for migrating a VM NIC.
*/
type MigrateNicConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import4.IPv4Address `json:"ipAddress,omitempty"`

	MigrateType *MigrateNicType `json:"migrateType"`

	Subnet *SubnetReference `json:"subnet"`
}

func (p *MigrateNicConfig) MarshalJSON() ([]byte, error) {
	type MigrateNicConfigProxy MigrateNicConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*MigrateNicConfigProxy
		MigrateType *MigrateNicType  `json:"migrateType,omitempty"`
		Subnet      *SubnetReference `json:"subnet,omitempty"`
	}{
		MigrateNicConfigProxy: (*MigrateNicConfigProxy)(p),
		MigrateType:           p.MigrateType,
		Subnet:                p.Subnet,
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

func (p *MigrateNicConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateNicConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrateNicConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "migrateType")
	delete(allFields, "subnet")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMigrateNicConfig() *MigrateNicConfig {
	p := new(MigrateNicConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateNicConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of IP address management for NIC migration.
*/
type MigrateNicType int

const (
	MIGRATENICTYPE_UNKNOWN    MigrateNicType = 0
	MIGRATENICTYPE_REDACTED   MigrateNicType = 1
	MIGRATENICTYPE_ASSIGN_IP  MigrateNicType = 2
	MIGRATENICTYPE_RELEASE_IP MigrateNicType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MigrateNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGN_IP",
		"RELEASE_IP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MigrateNicType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGN_IP",
		"RELEASE_IP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MigrateNicType) index(name string) MigrateNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGN_IP",
		"RELEASE_IP",
	}
	for idx := range names {
		if names[idx] == name {
			return MigrateNicType(idx)
		}
	}
	return MIGRATENICTYPE_UNKNOWN
}

func (e *MigrateNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MigrateNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MigrateNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MigrateNicType) Ref() *MigrateNicType {
	return &e
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/migrate-vm-disks Post operation
*/
type MigrateVmDisksApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrateVmDisksApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *MigrateVmDisksApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MigrateVmDisksApiResponse

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

func (p *MigrateVmDisksApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateVmDisksApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrateVmDisksApiResponse(*known)

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

func NewMigrateVmDisksApiResponse() *MigrateVmDisksApiResponse {
	p := new(MigrateVmDisksApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateVmDisksApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrateVmDisksApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrateVmDisksApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrateVmDisksApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/migrate-to-host Post operation
*/
type MigrateVmToHostApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrateVmToHostApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *MigrateVmToHostApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MigrateVmToHostApiResponse

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

func (p *MigrateVmToHostApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateVmToHostApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrateVmToHostApiResponse(*known)

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

func NewMigrateVmToHostApiResponse() *MigrateVmToHostApiResponse {
	p := new(MigrateVmToHostApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrateVmToHostApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrateVmToHostApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrateVmToHostApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrateVmToHostApiResponseData()
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
List of migration plans for VM disks.
*/
type MigrationPlans struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Plans []ADSFDiskMigrationPlan `json:"plans,omitempty"`
}

func (p *MigrationPlans) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MigrationPlans

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

func (p *MigrationPlans) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrationPlans
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MigrationPlans(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "plans")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMigrationPlans() *MigrationPlans {
	p := new(MigrationPlans)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.MigrationPlans"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The network function chain associates with the NIC. Only valid if nic_type is NORMAL_NIC.
*/
type NetworkFunctionChainReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of a network function chain. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *NetworkFunctionChainReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NetworkFunctionChainReference

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

func (p *NetworkFunctionChainReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NetworkFunctionChainReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NetworkFunctionChainReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNetworkFunctionChainReference() *NetworkFunctionChainReference {
	p := new(NetworkFunctionChainReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NetworkFunctionChainReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of this Network function NIC. Defaults to INGRESS.
*/
type NetworkFunctionNicType int

const (
	NETWORKFUNCTIONNICTYPE_UNKNOWN  NetworkFunctionNicType = 0
	NETWORKFUNCTIONNICTYPE_REDACTED NetworkFunctionNicType = 1
	NETWORKFUNCTIONNICTYPE_INGRESS  NetworkFunctionNicType = 2
	NETWORKFUNCTIONNICTYPE_EGRESS   NetworkFunctionNicType = 3
	NETWORKFUNCTIONNICTYPE_TAP      NetworkFunctionNicType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkFunctionNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkFunctionNicType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkFunctionNicType) index(name string) NetworkFunctionNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkFunctionNicType(idx)
		}
	}
	return NETWORKFUNCTIONNICTYPE_UNKNOWN
}

func (e *NetworkFunctionNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkFunctionNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkFunctionNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkFunctionNicType) Ref() *NetworkFunctionNicType {
	return &e
}

/*
The capabilities of the Nutanix Guest Tools in the VM.
*/
type NgtCapability int

const (
	NGTCAPABILITY_UNKNOWN              NgtCapability = 0
	NGTCAPABILITY_REDACTED             NgtCapability = 1
	NGTCAPABILITY_SELF_SERVICE_RESTORE NgtCapability = 2
	NGTCAPABILITY_VSS_SNAPSHOT         NgtCapability = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NgtCapability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NgtCapability) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NgtCapability) index(name string) NgtCapability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	for idx := range names {
		if names[idx] == name {
			return NgtCapability(idx)
		}
	}
	return NGTCAPABILITY_UNKNOWN
}

func (e *NgtCapability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NgtCapability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NgtCapability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NgtCapability) Ref() *NgtCapability {
	return &e
}

/*
Virtual Machine NIC.
*/
type Nic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingInfo *EmulatedNic `json:"backingInfo,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	NetworkInfo *NicNetworkInfo `json:"networkInfo,omitempty"`
	/*

	 */
	NicBackingInfoItemDiscriminator_ *string `json:"$nicBackingInfoItemDiscriminator,omitempty"`
	/*
	  Information about how NIC is associated with a VM.
	*/
	NicBackingInfo *OneOfNicNicBackingInfo `json:"nicBackingInfo,omitempty"`
	/*

	 */
	NicNetworkInfoItemDiscriminator_ *string `json:"$nicNetworkInfoItemDiscriminator,omitempty"`
	/*
	  Networking information object for a NIC.
	*/
	NicNetworkInfo *OneOfNicNicNetworkInfo `json:"nicNetworkInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Nic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Nic

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

func (p *Nic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Nic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Nic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backingInfo")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "networkInfo")
	delete(allFields, "$nicBackingInfoItemDiscriminator")
	delete(allFields, "nicBackingInfo")
	delete(allFields, "$nicNetworkInfoItemDiscriminator")
	delete(allFields, "nicNetworkInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNic() *Nic {
	p := new(Nic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Nic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Nic) GetNicNetworkInfo() interface{} {
	if nil == p.NicNetworkInfo {
		return nil
	}
	return p.NicNetworkInfo.GetValue()
}

func (p *Nic) SetNicNetworkInfo(v interface{}) error {
	if nil == p.NicNetworkInfo {
		p.NicNetworkInfo = NewOneOfNicNicNetworkInfo()
	}
	e := p.NicNetworkInfo.SetValue(v)
	if nil == e {
		if nil == p.NicNetworkInfoItemDiscriminator_ {
			p.NicNetworkInfoItemDiscriminator_ = new(string)
		}
		*p.NicNetworkInfoItemDiscriminator_ = *p.NicNetworkInfo.Discriminator
	}
	return e
}

/*
Networking information object for a NIC. This object is now deprecated. If both deprecated and new objects are present, the new object VirtualEthernetNicNetworkInfo takes precedence.
*/
type NicNetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4Config *Ipv4Config `json:"ipv4Config,omitempty"`

	Ipv4Info *Ipv4Info `json:"ipv4Info,omitempty"`

	NetworkFunctionChain *NetworkFunctionChainReference `json:"networkFunctionChain,omitempty"`

	NetworkFunctionNicType *NetworkFunctionNicType `json:"networkFunctionNicType,omitempty"`

	NicType *NicType `json:"nicType,omitempty"`
	/*
	  Indicates whether an unknown unicast traffic is forwarded to this NIC or not. This is applicable only for the NICs on the overlay subnets.
	*/
	ShouldAllowUnknownMacs *bool `json:"shouldAllowUnknownMacs,omitempty"`

	Subnet *SubnetReference `json:"subnet,omitempty"`
	/*
	  List of networks to trunk if VLAN mode is marked as TRUNKED. If empty and VLAN mode is set to TRUNKED, all the VLANs are trunked.
	*/
	TrunkedVlans []int `json:"trunkedVlans,omitempty"`

	VlanMode *VlanMode `json:"vlanMode,omitempty"`
}

func (p *NicNetworkInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NicNetworkInfo

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

func (p *NicNetworkInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NicNetworkInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NicNetworkInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipv4Config")
	delete(allFields, "ipv4Info")
	delete(allFields, "networkFunctionChain")
	delete(allFields, "networkFunctionNicType")
	delete(allFields, "nicType")
	delete(allFields, "shouldAllowUnknownMacs")
	delete(allFields, "subnet")
	delete(allFields, "trunkedVlans")
	delete(allFields, "vlanMode")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNicNetworkInfo() *NicNetworkInfo {
	p := new(NicNetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NicNetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to a NIC profile.
*/
type NicProfileReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier of a NIC profile. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *NicProfileReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NicProfileReference

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

func (p *NicProfileReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NicProfileReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NicProfileReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNicProfileReference() *NicProfileReference {
	p := new(NicProfileReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NicProfileReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NIC type. Defaults to NORMAL_NIC.
*/
type NicType int

const (
	NICTYPE_UNKNOWN              NicType = 0
	NICTYPE_REDACTED             NicType = 1
	NICTYPE_NORMAL_NIC           NicType = 2
	NICTYPE_DIRECT_NIC           NicType = 3
	NICTYPE_NETWORK_FUNCTION_NIC NicType = 4
	NICTYPE_SPAN_DESTINATION_NIC NicType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
		"SPAN_DESTINATION_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NicType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
		"SPAN_DESTINATION_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NicType) index(name string) NicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
		"SPAN_DESTINATION_NIC",
	}
	for idx := range names {
		if names[idx] == name {
			return NicType(idx)
		}
	}
	return NICTYPE_UNKNOWN
}

func (e *NicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicType) Ref() *NicType {
	return &e
}

/*
Configuration for NVRAM to be presented to the VM.
*/
type NvramDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackingStorageInfo *VmDisk `json:"backingStorageInfo,omitempty"`
}

func (p *NvramDevice) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NvramDevice

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

func (p *NvramDevice) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NvramDevice
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NvramDevice(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backingStorageInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNvramDevice() *NvramDevice {
	p := new(NvramDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.NvramDevice"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM config override spec for OVA VM deploy endpoint.
*/
type OvaVmConfigOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Categories for the VM.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`
	/*
	  CD-ROMs attached to the VM.
	*/
	CdRoms []CdRom `json:"cdRoms,omitempty"`
	/*
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  NICs attached to the VM.
	*/
	Nics []Nic `json:"nics"`
	/*
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/*
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/*
	  Number of threads per core.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
}

func (p *OvaVmConfigOverrideSpec) MarshalJSON() ([]byte, error) {
	type OvaVmConfigOverrideSpecProxy OvaVmConfigOverrideSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*OvaVmConfigOverrideSpecProxy
		Nics []Nic `json:"nics,omitempty"`
	}{
		OvaVmConfigOverrideSpecProxy: (*OvaVmConfigOverrideSpecProxy)(p),
		Nics:                         p.Nics,
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

func (p *OvaVmConfigOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OvaVmConfigOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OvaVmConfigOverrideSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")
	delete(allFields, "cdRoms")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "name")
	delete(allFields, "nics")
	delete(allFields, "numCoresPerSocket")
	delete(allFields, "numSockets")
	delete(allFields, "numThreadsPerCore")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOvaVmConfigOverrideSpec() *OvaVmConfigOverrideSpec {
	p := new(OvaVmConfigOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.OvaVmConfigOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Owner reference.
*/
type OwnerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM owner type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *OwnerReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OwnerReference

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

func (p *OwnerReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OwnerReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OwnerReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Ownership information for the VM.
*/
type OwnershipInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Owner *OwnerReference `json:"owner,omitempty"`
}

func (p *OwnershipInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OwnershipInfo

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

func (p *OwnershipInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OwnershipInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OwnershipInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "owner")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOwnershipInfo() *OwnershipInfo {
	p := new(OwnershipInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.OwnershipInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the information about the PCIe device.
*/
type PcieDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AssignedDeviceInfo *PcieDeviceInfo `json:"assignedDeviceInfo,omitempty"`
	/*

	 */
	BackingInfoItemDiscriminator_ *string `json:"$backingInfoItemDiscriminator,omitempty"`
	/*
	  Indicates the way a PCIe device is associated to the VM.
	*/
	BackingInfo *OneOfPcieDeviceBackingInfo `json:"backingInfo,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
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
	delete(allFields, "assignedDeviceInfo")
	delete(allFields, "$backingInfoItemDiscriminator")
	delete(allFields, "backingInfo")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPcieDevice() *PcieDevice {
	p := new(PcieDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PcieDevice"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PcieDevice) GetBackingInfo() interface{} {
	if nil == p.BackingInfo {
		return nil
	}
	return p.BackingInfo.GetValue()
}

func (p *PcieDevice) SetBackingInfo(v interface{}) error {
	if nil == p.BackingInfo {
		p.BackingInfo = NewOneOfPcieDeviceBackingInfo()
	}
	e := p.BackingInfo.SetValue(v)
	if nil == e {
		if nil == p.BackingInfoItemDiscriminator_ {
			p.BackingInfoItemDiscriminator_ = new(string)
		}
		*p.BackingInfoItemDiscriminator_ = *p.BackingInfo.Discriminator
	}
	return e
}

/*
Information about the attached PCIe device to the VM.
*/
type PcieDeviceInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Device *PcieDeviceReference `json:"device,omitempty"`
}

func (p *PcieDeviceInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcieDeviceInfo

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

func (p *PcieDeviceInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcieDeviceInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PcieDeviceInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "device")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPcieDeviceInfo() *PcieDeviceInfo {
	p := new(PcieDeviceInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PcieDeviceInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the PCIe device.
*/
type PcieDeviceReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier denoting PCIe device label. It should be of type UUID.
	*/
	DeviceExtId *string `json:"deviceExtId,omitempty"`
}

func (p *PcieDeviceReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcieDeviceReference

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

func (p *PcieDeviceReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcieDeviceReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PcieDeviceReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "deviceExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPcieDeviceReference() *PcieDeviceReference {
	p := new(PcieDeviceReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PcieDeviceReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the policy object in use.
*/
type PolicyReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of an instance of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *PolicyReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PolicyReference

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

func (p *PolicyReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyReference() *PolicyReference {
	p := new(PolicyReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PolicyReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/power-cycle Post operation
*/
type PowerCycleVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerCycleVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *PowerCycleVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PowerCycleVmApiResponse

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

func (p *PowerCycleVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PowerCycleVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PowerCycleVmApiResponse(*known)

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

func NewPowerCycleVmApiResponse() *PowerCycleVmApiResponse {
	p := new(PowerCycleVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerCycleVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerCycleVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerCycleVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerCycleVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/power-off Post operation
*/
type PowerOffVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerOffVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *PowerOffVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PowerOffVmApiResponse

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

func (p *PowerOffVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PowerOffVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PowerOffVmApiResponse(*known)

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

func NewPowerOffVmApiResponse() *PowerOffVmApiResponse {
	p := new(PowerOffVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerOffVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerOffVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerOffVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerOffVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/power-on Post operation
*/
type PowerOnVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPowerOnVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *PowerOnVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PowerOnVmApiResponse

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

func (p *PowerOnVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PowerOnVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PowerOnVmApiResponse(*known)

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

func NewPowerOnVmApiResponse() *PowerOnVmApiResponse {
	p := new(PowerOnVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.PowerOnVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PowerOnVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PowerOnVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPowerOnVmApiResponseData()
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
The current power state of the VM.
*/
type PowerState int

const (
	POWERSTATE_UNKNOWN      PowerState = 0
	POWERSTATE_REDACTED     PowerState = 1
	POWERSTATE_ON           PowerState = 2
	POWERSTATE_OFF          PowerState = 3
	POWERSTATE_PAUSED       PowerState = 4
	POWERSTATE_UNDETERMINED PowerState = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PowerState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ON",
		"OFF",
		"PAUSED",
		"UNDETERMINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PowerState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ON",
		"OFF",
		"PAUSED",
		"UNDETERMINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PowerState) index(name string) PowerState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ON",
		"OFF",
		"PAUSED",
		"UNDETERMINED",
	}
	for idx := range names {
		if names[idx] == name {
			return PowerState(idx)
		}
	}
	return POWERSTATE_UNKNOWN
}

func (e *PowerState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PowerState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PowerState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PowerState) Ref() *PowerState {
	return &e
}

/*
Reference to a project.
*/
type ProjectReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of an instance of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *ProjectReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProjectReference

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

func (p *ProjectReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProjectReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ProjectReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewProjectReference() *ProjectReference {
	p := new(ProjectReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ProjectReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of protection policy applied to this VM.
*/
type ProtectionPolicyState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Policy *PolicyReference `json:"policy,omitempty"`
}

func (p *ProtectionPolicyState) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProtectionPolicyState

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

func (p *ProtectionPolicyState) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectionPolicyState
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ProtectionPolicyState(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "policy")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewProtectionPolicyState() *ProtectionPolicyState {
	p := new(ProtectionPolicyState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ProtectionPolicyState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of protection applied on a VM.
*/
type ProtectionType int

const (
	PROTECTIONTYPE_UNKNOWN        ProtectionType = 0
	PROTECTIONTYPE_REDACTED       ProtectionType = 1
	PROTECTIONTYPE_UNPROTECTED    ProtectionType = 2
	PROTECTIONTYPE_PD_PROTECTED   ProtectionType = 3
	PROTECTIONTYPE_RULE_PROTECTED ProtectionType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtectionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPROTECTED",
		"PD_PROTECTED",
		"RULE_PROTECTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProtectionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPROTECTED",
		"PD_PROTECTED",
		"RULE_PROTECTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProtectionType) index(name string) ProtectionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPROTECTED",
		"PD_PROTECTED",
		"RULE_PROTECTED",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectionType(idx)
		}
	}
	return PROTECTIONTYPE_UNKNOWN
}

func (e *ProtectionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectionType) Ref() *ProtectionType {
	return &e
}

/*
QoS parameters to be enforced.
*/
type QosConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Throttled IOPS for the governed entities. The block size for the I/O is 32 kB.
	*/
	ThrottledIops *int `json:"throttledIops,omitempty"`
}

func (p *QosConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias QosConfig

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

func (p *QosConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias QosConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = QosConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "throttledIops")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewQosConfig() *QosConfig {
	p := new(QosConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.QosConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The restart schedule after installing or upgrading Nutanix Guest Tools.
*/
type RebootPreference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Schedule *RebootPreferenceSchedule `json:"schedule,omitempty"`

	ScheduleType *ScheduleType `json:"scheduleType,omitempty"`
}

func (p *RebootPreference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RebootPreference

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

func (p *RebootPreference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RebootPreference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RebootPreference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "schedule")
	delete(allFields, "scheduleType")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRebootPreference() *RebootPreference {
	p := new(RebootPreference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootPreference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Restart schedule.
*/
type RebootPreferenceSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The start time for a scheduled restart.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func (p *RebootPreferenceSchedule) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RebootPreferenceSchedule

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

func (p *RebootPreferenceSchedule) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RebootPreferenceSchedule
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RebootPreferenceSchedule(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "startTime")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRebootPreferenceSchedule() *RebootPreferenceSchedule {
	p := new(RebootPreferenceSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootPreferenceSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/guest-reboot Post operation
*/
type RebootVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRebootVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RebootVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RebootVmApiResponse

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

func (p *RebootVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RebootVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RebootVmApiResponse(*known)

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

func NewRebootVmApiResponse() *RebootVmApiResponse {
	p := new(RebootVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RebootVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RebootVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RebootVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRebootVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId}/$actions/release-ip Post operation
*/
type ReleaseIpApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfReleaseIpApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ReleaseIpApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReleaseIpApiResponse

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

func (p *ReleaseIpApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReleaseIpApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ReleaseIpApiResponse(*known)

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

func NewReleaseIpApiResponse() *ReleaseIpApiResponse {
	p := new(ReleaseIpApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ReleaseIpApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReleaseIpApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ReleaseIpApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfReleaseIpApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/reset Post operation
*/
type ResetVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResetVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ResetVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ResetVmApiResponse

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

func (p *ResetVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResetVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ResetVmApiResponse(*known)

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

func NewResetVmApiResponse() *ResetVmApiResponse {
	p := new(ResetVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ResetVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ResetVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ResetVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfResetVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vm-recovery-points/{extId}/$actions/restore Post operation
*/
type RestoreVmRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRestoreVmRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RestoreVmRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RestoreVmRecoveryPointApiResponse

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

func (p *RestoreVmRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RestoreVmRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RestoreVmRecoveryPointApiResponse(*known)

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

func NewRestoreVmRecoveryPointApiResponse() *RestoreVmRecoveryPointApiResponse {
	p := new(RestoreVmRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RestoreVmRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RestoreVmRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RestoreVmRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRestoreVmRecoveryPointApiResponseData()
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
Input parameters for the VM recovery point restore operation.
*/
type RestoreVmRecoveryPointParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  If set to false, any VM configuration that cannot be restored will be dropped or reset to system defaults. This currently applies to the following fields in the VM configuration: subnets, NIC profiles, project, and categories. It also handles dropping unknown features on older PEs. If set to true, the restore operation will run in normal mode, where any failures will be raised as errors, and the VM will not be restored until the issues are resolved.
	*/
	IsStrictMode *bool `json:"isStrictMode,omitempty"`

	VmConfigOverrideSpec *VmConfigOverrideSpecification `json:"vmConfigOverrideSpec,omitempty"`
}

func (p *RestoreVmRecoveryPointParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RestoreVmRecoveryPointParams

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

func (p *RestoreVmRecoveryPointParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RestoreVmRecoveryPointParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RestoreVmRecoveryPointParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isStrictMode")
	delete(allFields, "vmConfigOverrideSpec")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRestoreVmRecoveryPointParams() *RestoreVmRecoveryPointParams {
	p := new(RestoreVmRecoveryPointParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RestoreVmRecoveryPointParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsStrictMode = new(bool)
	*p.IsStrictMode = true

	return p
}

/*
Input for the VM revert operation. Specify the VM Recovery Point ID to which the VM would be reverted.
*/
type RevertParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the VM Recovery Point.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId"`
}

func (p *RevertParams) MarshalJSON() ([]byte, error) {
	type RevertParamsProxy RevertParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RevertParamsProxy
		VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
	}{
		RevertParamsProxy:    (*RevertParamsProxy)(p),
		VmRecoveryPointExtId: p.VmRecoveryPointExtId,
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

func (p *RevertParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RevertParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RevertParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "vmRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRevertParams() *RevertParams {
	p := new(RevertParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RevertParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/revert Post operation
*/
type RevertVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRevertVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RevertVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RevertVmApiResponse

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

func (p *RevertVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RevertVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RevertVmApiResponse(*known)

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

func NewRevertVmApiResponse() *RevertVmApiResponse {
	p := new(RevertVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.RevertVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RevertVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RevertVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRevertVmApiResponseData()
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
The (S)egment:(B)us:(D)evice.(F)unction hardware address. See https://wiki.xen.org/wiki/Bus:Device.Function_(BDF)_Notation for more details.
*/
type SBDF struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  PCI Bus number in hexadecimal, often padded using a leading zeros to two or four digits.
	*/
	Bus *int `json:"bus,omitempty"`
	/*
	  PCI Device number in hexadecimal, often padded using a leading zero to two digits . Sometimes this is also referred to as the slot number.
	*/
	Device *int `json:"device,omitempty"`
	/*
	  PCI Function number in hexadecimal.
	*/
	Func *int `json:"func,omitempty"`
	/*
	  PCI Segment.
	*/
	Segment *int `json:"segment,omitempty"`
}

func (p *SBDF) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SBDF

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

func (p *SBDF) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SBDF
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SBDF(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bus")
	delete(allFields, "device")
	delete(allFields, "func")
	delete(allFields, "segment")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSBDF() *SBDF {
	p := new(SBDF)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SBDF"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schedule type for restart.
*/
type ScheduleType int

const (
	SCHEDULETYPE_UNKNOWN   ScheduleType = 0
	SCHEDULETYPE_REDACTED  ScheduleType = 1
	SCHEDULETYPE_SKIP      ScheduleType = 2
	SCHEDULETYPE_IMMEDIATE ScheduleType = 3
	SCHEDULETYPE_LATER     ScheduleType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScheduleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScheduleType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScheduleType) index(name string) ScheduleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	for idx := range names {
		if names[idx] == name {
			return ScheduleType(idx)
		}
	}
	return SCHEDULETYPE_UNKNOWN
}

func (e *ScheduleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScheduleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScheduleType) Ref() *ScheduleType {
	return &e
}

/*
Indicates the configuration of serial ports of the VM.
*/
type SerialPort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Index of the serial port. Index 1 is reserved for internal usage.
	*/
	Index *int `json:"index,omitempty"`
	/*
	  Indicates whether the serial port is connected or not.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SerialPort) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SerialPort

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

func (p *SerialPort) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SerialPort
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SerialPort(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "index")
	delete(allFields, "isConnected")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSerialPort() *SerialPort {
	p := new(SerialPort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SerialPort"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/$actions/guest-shutdown Post operation
*/
type ShutdownVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfShutdownVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ShutdownVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ShutdownVmApiResponse

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

func (p *ShutdownVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ShutdownVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ShutdownVmApiResponse(*known)

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

func NewShutdownVmApiResponse() *ShutdownVmApiResponse {
	p := new(ShutdownVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.ShutdownVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ShutdownVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ShutdownVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfShutdownVmApiResponseData()
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
Defines an SR-IOV NIC.
*/
type SriovNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HostPcieDeviceReference *HostPcieDeviceReference `json:"hostPcieDeviceReference,omitempty"`
	/*
	  Indicates whether the NIC is connected or not. Default is True.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`

	SriovProfileReference *NicProfileReference `json:"sriovProfileReference"`
}

func (p *SriovNic) MarshalJSON() ([]byte, error) {
	type SriovNicProxy SriovNic

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SriovNicProxy
		SriovProfileReference *NicProfileReference `json:"sriovProfileReference,omitempty"`
	}{
		SriovNicProxy:         (*SriovNicProxy)(p),
		SriovProfileReference: p.SriovProfileReference,
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

func (p *SriovNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SriovNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SriovNic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostPcieDeviceReference")
	delete(allFields, "isConnected")
	delete(allFields, "macAddress")
	delete(allFields, "sriovProfileReference")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSriovNic() *SriovNic {
	p := new(SriovNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SriovNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsConnected = new(bool)
	*p.IsConnected = true

	return p
}

/*
Networking information object for an SR-IOV NIC.
*/
type SriovNicNetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VLAN ID for the SR-IOV NIC.
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func (p *SriovNicNetworkInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SriovNicNetworkInfo

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

func (p *SriovNicNetworkInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SriovNicNetworkInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SriovNicNetworkInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "vlanId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSriovNicNetworkInfo() *SriovNicNetworkInfo {
	p := new(SriovNicNetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SriovNicNetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network identifier for this adapter. Only valid if nic_type is NORMAL_NIC or DIRECT_NIC.
*/
type SubnetReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The globally unique identifier of a subnet of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *SubnetReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SubnetReference

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

func (p *SubnetReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SubnetReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SubnetReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSubnetReference() *SubnetReference {
	p := new(SubnetReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.SubnetReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
If this field is set, the guest will be customized using Sysprep. Either unattend_xml or custom_key_values should be provided. If custom_key_values are provided then the unattended answer file will be generated using these key-value pairs.
*/
type Sysprep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	InstallType *InstallType `json:"installType,omitempty"`
	/*

	 */
	SysprepScriptItemDiscriminator_ *string `json:"$sysprepScriptItemDiscriminator,omitempty"`
	/*
	  The Sysprep script for guest customisation of VM.
	*/
	SysprepScript *OneOfSysprepSysprepScript `json:"sysprepScript,omitempty"`
}

func (p *Sysprep) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Sysprep

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

func (p *Sysprep) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Sysprep
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Sysprep(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "installType")
	delete(allFields, "$sysprepScriptItemDiscriminator")
	delete(allFields, "sysprepScript")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSysprep() *Sysprep {
	p := new(Sysprep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Sysprep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Sysprep) GetSysprepScript() interface{} {
	if nil == p.SysprepScript {
		return nil
	}
	return p.SysprepScript.GetValue()
}

func (p *Sysprep) SetSysprepScript(v interface{}) error {
	if nil == p.SysprepScript {
		p.SysprepScript = NewOneOfSysprepSysprepScript()
	}
	e := p.SysprepScript.SetValue(v)
	if nil == e {
		if nil == p.SysprepScriptItemDiscriminator_ {
			p.SysprepScriptItemDiscriminator_ = new(string)
		}
		*p.SysprepScriptItemDiscriminator_ = *p.SysprepScript.Discriminator
	}
	return e
}

/*
UEFI boot mode and its associated configuration.
*/
type UefiBoot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	BootDeviceItemDiscriminator_ *string `json:"$bootDeviceItemDiscriminator,omitempty"`
	/*
	  The boot device used for UEFI boot of the VM.
	*/
	BootDevice *OneOfUefiBootBootDevice `json:"bootDevice,omitempty"`
	/*
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootOrder []BootDeviceType `json:"bootOrder,omitempty"`
	/*
	  Indicate whether to enable secure boot or not.
	*/
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	NvramDevice *NvramDevice `json:"nvramDevice,omitempty"`
}

func (p *UefiBoot) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UefiBoot

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

func (p *UefiBoot) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UefiBoot
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UefiBoot(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$bootDeviceItemDiscriminator")
	delete(allFields, "bootDevice")
	delete(allFields, "bootOrder")
	delete(allFields, "isSecureBootEnabled")
	delete(allFields, "nvramDevice")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUefiBoot() *UefiBoot {
	p := new(UefiBoot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UefiBoot"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UefiBoot) GetBootDevice() interface{} {
	if nil == p.BootDevice {
		return nil
	}
	return p.BootDevice.GetValue()
}

func (p *UefiBoot) SetBootDevice(v interface{}) error {
	if nil == p.BootDevice {
		p.BootDevice = NewOneOfUefiBootBootDevice()
	}
	e := p.BootDevice.SetValue(v)
	if nil == e {
		if nil == p.BootDeviceItemDiscriminator_ {
			p.BootDeviceItemDiscriminator_ = new(string)
		}
		*p.BootDeviceItemDiscriminator_ = *p.BootDevice.Discriminator
	}
	return e
}

/*
This field contains a Sysprep unattend xml definition, as a string. The value must be base64 encoded.
*/
type Unattendxml struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value of Sysprep unattend xml definition. The value must be base64 encoded.
	*/
	Value *string `json:"value,omitempty"`
}

func (p *Unattendxml) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Unattendxml

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

func (p *Unattendxml) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Unattendxml
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Unattendxml(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUnattendxml() *Unattendxml {
	p := new(Unattendxml)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Unattendxml"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools/$actions/uninstall Post operation
*/
type UninstallVmGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUninstallVmGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UninstallVmGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UninstallVmGuestToolsApiResponse

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

func (p *UninstallVmGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UninstallVmGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UninstallVmGuestToolsApiResponse(*known)

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

func NewUninstallVmGuestToolsApiResponse() *UninstallVmGuestToolsApiResponse {
	p := new(UninstallVmGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UninstallVmGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UninstallVmGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UninstallVmGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUninstallVmGuestToolsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/disks/{extId} Put operation
*/
type UpdateDiskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDiskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateDiskApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateDiskApiResponse

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

func (p *UpdateDiskApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateDiskApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateDiskApiResponse(*known)

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

func NewUpdateDiskApiResponse() *UpdateDiskApiResponse {
	p := new(UpdateDiskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateDiskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDiskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDiskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDiskApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools Put operation
*/
type UpdateGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateGuestToolsApiResponse

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

func (p *UpdateGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateGuestToolsApiResponse(*known)

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

func NewUpdateGuestToolsApiResponse() *UpdateGuestToolsApiResponse {
	p := new(UpdateGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateGuestToolsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/nics/{extId} Put operation
*/
type UpdateNicApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNicApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateNicApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateNicApiResponse

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

func (p *UpdateNicApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateNicApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateNicApiResponse(*known)

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

func NewUpdateNicApiResponse() *UpdateNicApiResponse {
	p := new(UpdateNicApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateNicApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNicApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNicApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNicApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{vmExtId}/serial-ports/{extId} Put operation
*/
type UpdateSerialPortApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSerialPortApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateSerialPortApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateSerialPortApiResponse

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

func (p *UpdateSerialPortApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateSerialPortApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateSerialPortApiResponse(*known)

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

func NewUpdateSerialPortApiResponse() *UpdateSerialPortApiResponse {
	p := new(UpdateSerialPortApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateSerialPortApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSerialPortApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSerialPortApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSerialPortApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId} Put operation
*/
type UpdateVmApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateVmApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateVmApiResponse

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

func (p *UpdateVmApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateVmApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateVmApiResponse(*known)

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

func NewUpdateVmApiResponse() *UpdateVmApiResponse {
	p := new(UpdateVmApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpdateVmApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/config/vms/{extId}/guest-tools/$actions/upgrade Post operation
*/
type UpgradeVmGuestToolsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpgradeVmGuestToolsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpgradeVmGuestToolsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpgradeVmGuestToolsApiResponse

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

func (p *UpgradeVmGuestToolsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpgradeVmGuestToolsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpgradeVmGuestToolsApiResponse(*known)

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

func NewUpgradeVmGuestToolsApiResponse() *UpgradeVmGuestToolsApiResponse {
	p := new(UpgradeVmGuestToolsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.UpgradeVmGuestToolsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpgradeVmGuestToolsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpgradeVmGuestToolsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpgradeVmGuestToolsApiResponseData()
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
The contents of the user_data configuration for cloud-init. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.
*/
type Userdata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The value for the cloud-init user_data.
	*/
	Value *string `json:"value,omitempty"`
}

func (p *Userdata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Userdata

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

func (p *Userdata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Userdata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Userdata(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUserdata() *Userdata {
	p := new(Userdata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Userdata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines a Virtual Ethernet NIC.
*/
type VirtualEthernetNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the NIC is connected or not. Default is True.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`

	Model *VirtualEthernetNicModel `json:"model,omitempty"`
	/*
	  The number of Tx/Rx queue pairs for this NIC.
	*/
	NumQueues *int `json:"numQueues,omitempty"`
}

func (p *VirtualEthernetNic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VirtualEthernetNic

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

func (p *VirtualEthernetNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VirtualEthernetNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VirtualEthernetNic(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isConnected")
	delete(allFields, "macAddress")
	delete(allFields, "model")
	delete(allFields, "numQueues")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVirtualEthernetNic() *VirtualEthernetNic {
	p := new(VirtualEthernetNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VirtualEthernetNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsConnected = new(bool)
	*p.IsConnected = true
	p.NumQueues = new(int)
	*p.NumQueues = 1

	return p
}

/*
Options for the NIC emulation.
*/
type VirtualEthernetNicModel int

const (
	VIRTUALETHERNETNICMODEL_UNKNOWN  VirtualEthernetNicModel = 0
	VIRTUALETHERNETNICMODEL_REDACTED VirtualEthernetNicModel = 1
	VIRTUALETHERNETNICMODEL_VIRTIO   VirtualEthernetNicModel = 2
	VIRTUALETHERNETNICMODEL_E1000    VirtualEthernetNicModel = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VirtualEthernetNicModel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VirtualEthernetNicModel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VirtualEthernetNicModel) index(name string) VirtualEthernetNicModel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	for idx := range names {
		if names[idx] == name {
			return VirtualEthernetNicModel(idx)
		}
	}
	return VIRTUALETHERNETNICMODEL_UNKNOWN
}

func (e *VirtualEthernetNicModel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VirtualEthernetNicModel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VirtualEthernetNicModel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VirtualEthernetNicModel) Ref() *VirtualEthernetNicModel {
	return &e
}

/*
Networking information object for a Virtual Ethernet NIC.
*/
type VirtualEthernetNicNetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4Config *Ipv4Config `json:"ipv4Config,omitempty"`

	Ipv4Info *Ipv4Info `json:"ipv4Info,omitempty"`

	Ipv6Info *Ipv6Info `json:"ipv6Info,omitempty"`

	NetworkFunctionChain *NetworkFunctionChainReference `json:"networkFunctionChain,omitempty"`

	NetworkFunctionNicType *NetworkFunctionNicType `json:"networkFunctionNicType,omitempty"`

	NicType *NicType `json:"nicType,omitempty"`
	/*
	  Indicates whether an unknown unicast traffic is forwarded to this NIC or not. This is applicable only for the NICs on the overlay subnets.
	*/
	ShouldAllowUnknownMacs *bool `json:"shouldAllowUnknownMacs,omitempty"`

	Subnet *SubnetReference `json:"subnet,omitempty"`
	/*
	  List of networks to trunk if VLAN mode is marked as TRUNKED. If empty and VLAN mode is set to TRUNKED, all the VLANs are trunked.
	*/
	TrunkedVlans []int `json:"trunkedVlans,omitempty"`

	VlanMode *VlanMode `json:"vlanMode,omitempty"`
}

func (p *VirtualEthernetNicNetworkInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VirtualEthernetNicNetworkInfo

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

func (p *VirtualEthernetNicNetworkInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VirtualEthernetNicNetworkInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VirtualEthernetNicNetworkInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipv4Config")
	delete(allFields, "ipv4Info")
	delete(allFields, "ipv6Info")
	delete(allFields, "networkFunctionChain")
	delete(allFields, "networkFunctionNicType")
	delete(allFields, "nicType")
	delete(allFields, "shouldAllowUnknownMacs")
	delete(allFields, "subnet")
	delete(allFields, "trunkedVlans")
	delete(allFields, "vlanMode")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVirtualEthernetNicNetworkInfo() *VirtualEthernetNicNetworkInfo {
	p := new(VirtualEthernetNicNetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VirtualEthernetNicNetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
By default, all the virtual NICs are created in ACCESS mode, which permits only one VLAN per virtual network. TRUNKED mode allows multiple VLANs on a single VM NIC for network-aware user VMs.
*/
type VlanMode int

const (
	VLANMODE_UNKNOWN  VlanMode = 0
	VLANMODE_REDACTED VlanMode = 1
	VLANMODE_ACCESS   VlanMode = 2
	VLANMODE_TRUNK    VlanMode = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VlanMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VlanMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VlanMode) index(name string) VlanMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNK",
	}
	for idx := range names {
		if names[idx] == name {
			return VlanMode(idx)
		}
	}
	return VLANMODE_UNKNOWN
}

func (e *VlanMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VlanMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VlanMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VlanMode) Ref() *VlanMode {
	return &e
}

/*
VM configuration.
*/
type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApcConfig *ApcConfig `json:"apcConfig,omitempty"`

	AvailabilityZone *AvailabilityZoneReference `json:"availabilityZone,omitempty"`
	/*
	  BIOS UUID of the VM of type UUID.
	*/
	BiosUuid *string `json:"biosUuid,omitempty"`
	/*

	 */
	BootConfigItemDiscriminator_ *string `json:"$bootConfigItemDiscriminator,omitempty"`
	/*
	  Indicates the order of device types in which the VM should try to boot from. If the boot device order is not provided the system will decide an appropriate boot device order.
	*/
	BootConfig *OneOfVmBootConfig `json:"bootConfig,omitempty"`
	/*
	  Categories for the VM.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`
	/*
	  CD-ROMs attached to the VM.
	*/
	CdRoms []CdRom `json:"cdRoms,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/*
	  VM creation time.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  VM description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Disks attached to the VM.
	*/
	Disks []Disk `json:"disks,omitempty"`
	/*
	  The list of additional CPU features to be enabled.
	*/
	EnabledCpuFeatures []CpuFeature `json:"enabledCpuFeatures,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Generation UUID of the VM of type UUID.
	*/
	GenerationUuid *string `json:"generationUuid,omitempty"`
	/*
	  GPUs attached to the VM.
	*/
	Gpus []Gpu `json:"gpus,omitempty"`

	GuestCustomization *GuestCustomizationParams `json:"guestCustomization,omitempty"`

	GuestTools *GuestTools `json:"guestTools,omitempty"`
	/*
	  VM hardware clock timezone in IANA TZDB format (America/Los_Angeles).
	*/
	HardwareClockTimezone *string `json:"hardwareClockTimezone,omitempty"`

	Host *HostReference `json:"host,omitempty"`
	/*
	  Indicates whether the VM is an agent VM or not. When their host enters maintenance mode, once the normal VMs are evacuated, the agent VMs are powered off. When the host is restored, agent VMs are powered on before the normal VMs are restored. In other words, agent VMs cannot be HA-protected or live migrated.
	*/
	IsAgentVm *bool `json:"isAgentVm,omitempty"`
	/*
	  Indicates whether to remove AHV branding from VM firmware tables or not.
	*/
	IsBrandingEnabled *bool `json:"isBrandingEnabled,omitempty"`
	/*
	  Indicates whether the VM CPU hotplug is enabled.
	*/
	IsCpuHotplugEnabled *bool `json:"isCpuHotplugEnabled,omitempty"`
	/*
	  Indicates whether to passthrough the host CPU features to the guest or not. Enabling this will make VM incapable of live migration.
	*/
	IsCpuPassthroughEnabled *bool `json:"isCpuPassthroughEnabled,omitempty"`
	/*
	  Indicates whether the VM is currently undergoing cross cluster migration.
	*/
	IsCrossClusterMigrationInProgress *bool `json:"isCrossClusterMigrationInProgress,omitempty"`
	/*
	  Indicates whether the vGPU console is enabled or not.
	*/
	IsGpuConsoleEnabled *bool `json:"isGpuConsoleEnabled,omitempty"`
	/*
	  Indicates whether the VM is capable of live migrating to another host or not, based on its configuration.
	*/
	IsLiveMigrateCapable *bool `json:"isLiveMigrateCapable,omitempty"`
	/*
	  Indicates whether the memory overcommit feature should be enabled for the VM or not. If enabled, parts of the VM memory may reside outside of the hypervisor physical memory. Once enabled, it should be expected that the VM may suffer performance degradation.
	*/
	IsMemoryOvercommitEnabled *bool `json:"isMemoryOvercommitEnabled,omitempty"`
	/*
	  Indicates whether the VM SCSI controller is enabled.
	*/
	IsScsiControllerEnabled *bool `json:"isScsiControllerEnabled,omitempty"`
	/*
	  Indicates whether the vCPUs should be hard pinned to specific pCPUs or not.
	*/
	IsVcpuHardPinningEnabled *bool `json:"isVcpuHardPinningEnabled,omitempty"`
	/*
	  Indicates whether the VGA console should be disabled or not.
	*/
	IsVgaConsoleEnabled *bool `json:"isVgaConsoleEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	MachineType *MachineType `json:"machineType,omitempty"`
	/*
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  NICs attached to the VM.
	*/
	Nics []Nic `json:"nics,omitempty"`
	/*
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/*
	  Number of NUMA nodes. 0 means NUMA is disabled.
	*/
	NumNumaNodes *int `json:"numNumaNodes,omitempty"`
	/*
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/*
	  Number of threads per core.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`

	OwnershipInfo *OwnershipInfo `json:"ownershipInfo,omitempty"`
	/*
	  PCIe devices attached to the VM.
	*/
	PcieDevices []PcieDevice `json:"pcieDevices,omitempty"`

	PowerState *PowerState `json:"powerState,omitempty"`

	Project *ProjectReference `json:"project,omitempty"`

	ProtectionPolicyState *ProtectionPolicyState `json:"protectionPolicyState,omitempty"`

	ProtectionType *ProtectionType `json:"protectionType,omitempty"`
	/*
	  Serial ports configured on the VM.
	*/
	SerialPorts []SerialPort `json:"serialPorts,omitempty"`

	Source *VmSourceReference `json:"source,omitempty"`

	StorageConfig *ADSFVmStorageConfig `json:"storageConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VM last updated time.
	*/
	UpdateTime *time.Time `json:"updateTime,omitempty"`

	VtpmConfig *VtpmConfig `json:"vtpmConfig,omitempty"`
}

func (p *Vm) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Vm

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

func (p *Vm) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Vm
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Vm(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apcConfig")
	delete(allFields, "availabilityZone")
	delete(allFields, "biosUuid")
	delete(allFields, "$bootConfigItemDiscriminator")
	delete(allFields, "bootConfig")
	delete(allFields, "categories")
	delete(allFields, "cdRoms")
	delete(allFields, "cluster")
	delete(allFields, "createTime")
	delete(allFields, "description")
	delete(allFields, "disks")
	delete(allFields, "enabledCpuFeatures")
	delete(allFields, "extId")
	delete(allFields, "generationUuid")
	delete(allFields, "gpus")
	delete(allFields, "guestCustomization")
	delete(allFields, "guestTools")
	delete(allFields, "hardwareClockTimezone")
	delete(allFields, "host")
	delete(allFields, "isAgentVm")
	delete(allFields, "isBrandingEnabled")
	delete(allFields, "isCpuHotplugEnabled")
	delete(allFields, "isCpuPassthroughEnabled")
	delete(allFields, "isCrossClusterMigrationInProgress")
	delete(allFields, "isGpuConsoleEnabled")
	delete(allFields, "isLiveMigrateCapable")
	delete(allFields, "isMemoryOvercommitEnabled")
	delete(allFields, "isScsiControllerEnabled")
	delete(allFields, "isVcpuHardPinningEnabled")
	delete(allFields, "isVgaConsoleEnabled")
	delete(allFields, "links")
	delete(allFields, "machineType")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "name")
	delete(allFields, "nics")
	delete(allFields, "numCoresPerSocket")
	delete(allFields, "numNumaNodes")
	delete(allFields, "numSockets")
	delete(allFields, "numThreadsPerCore")
	delete(allFields, "ownershipInfo")
	delete(allFields, "pcieDevices")
	delete(allFields, "powerState")
	delete(allFields, "project")
	delete(allFields, "protectionPolicyState")
	delete(allFields, "protectionType")
	delete(allFields, "serialPorts")
	delete(allFields, "source")
	delete(allFields, "storageConfig")
	delete(allFields, "tenantId")
	delete(allFields, "updateTime")
	delete(allFields, "vtpmConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Vm"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HardwareClockTimezone = new(string)
	*p.HardwareClockTimezone = "UTC"

	return p
}

func (p *Vm) GetBootConfig() interface{} {
	if nil == p.BootConfig {
		return nil
	}
	return p.BootConfig.GetValue()
}

func (p *Vm) SetBootConfig(v interface{}) error {
	if nil == p.BootConfig {
		p.BootConfig = NewOneOfVmBootConfig()
	}
	e := p.BootConfig.SetValue(v)
	if nil == e {
		if nil == p.BootConfigItemDiscriminator_ {
			p.BootConfigItemDiscriminator_ = new(string)
		}
		*p.BootConfigItemDiscriminator_ = *p.BootConfig.Discriminator
	}
	return e
}

/*
To override the VM configuration captured in the snapshot (VM Recovery Point). Whatever values are specified in this object, will override the corresponding VM config entry from the snapshot. For example, if a new list of NICs are specified, the existing NIC details are replaced with the provided list. If the list of NICs is set to an empty list explicitly, the VM will be created with no NICs. For the case where the NICs are not set in the override spec at all, the NICs from the captured VM config will be attempted during restore.
*/
type VmConfigOverrideSpecification struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Categories to be associated with the VM on successful restore. If not specified, the VM is provisioned without any categories.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`
	/*
	  VM description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of the VM to override with. If not specified, a name is chosen by the system and returned to the task entities when complete.
	*/
	Name *string `json:"name,omitempty"`

	NicSpec *VmRestoreNicConfigSpecification `json:"nicSpec,omitempty"`

	OwnershipInfo *OwnershipInfo `json:"ownershipInfo,omitempty"`
}

func (p *VmConfigOverrideSpecification) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmConfigOverrideSpecification

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

func (p *VmConfigOverrideSpecification) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmConfigOverrideSpecification
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmConfigOverrideSpecification(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")
	delete(allFields, "description")
	delete(allFields, "name")
	delete(allFields, "nicSpec")
	delete(allFields, "ownershipInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmConfigOverrideSpecification() *VmConfigOverrideSpecification {
	p := new(VmConfigOverrideSpecification)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmConfigOverrideSpecification"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cross cluster migration overrides the VM configuration.
*/
type VmCrossClusterMigrateOverrides struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The NIC configuration to apply on the target cluster when migrating a VM from the source cluster to the target cluster.
	*/
	OverrideNicList []Nic `json:"overrideNicList,omitempty"`
}

func (p *VmCrossClusterMigrateOverrides) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmCrossClusterMigrateOverrides

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

func (p *VmCrossClusterMigrateOverrides) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmCrossClusterMigrateOverrides
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmCrossClusterMigrateOverrides(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "overrideNicList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmCrossClusterMigrateOverrides() *VmCrossClusterMigrateOverrides {
	p := new(VmCrossClusterMigrateOverrides)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmCrossClusterMigrateOverrides"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input on how to migrate a VM across clusters.
*/
type VmCrossClusterMigrateParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if the migration is performed with a running VM.
	*/
	IsLiveMigration *bool `json:"isLiveMigration"`

	Overrides *VmCrossClusterMigrateOverrides `json:"overrides,omitempty"`

	TargetAvailabilityZone *AvailabilityZoneReference `json:"targetAvailabilityZone"`

	TargetCluster *ClusterReference `json:"targetCluster,omitempty"`
}

func (p *VmCrossClusterMigrateParams) MarshalJSON() ([]byte, error) {
	type VmCrossClusterMigrateParamsProxy VmCrossClusterMigrateParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VmCrossClusterMigrateParamsProxy
		IsLiveMigration        *bool                      `json:"isLiveMigration,omitempty"`
		TargetAvailabilityZone *AvailabilityZoneReference `json:"targetAvailabilityZone,omitempty"`
	}{
		VmCrossClusterMigrateParamsProxy: (*VmCrossClusterMigrateParamsProxy)(p),
		IsLiveMigration:                  p.IsLiveMigration,
		TargetAvailabilityZone:           p.TargetAvailabilityZone,
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

func (p *VmCrossClusterMigrateParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmCrossClusterMigrateParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmCrossClusterMigrateParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isLiveMigration")
	delete(allFields, "overrides")
	delete(allFields, "targetAvailabilityZone")
	delete(allFields, "targetCluster")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmCrossClusterMigrateParams() *VmCrossClusterMigrateParams {
	p := new(VmCrossClusterMigrateParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmCrossClusterMigrateParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Storage provided by Nutanix ADSF.
*/
type VmDisk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataSource *DataSource `json:"dataSource,omitempty"`
	/*
	  A globally unique identifier of a VM disk of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
	/*
	  Size of the disk in bytes.
	*/
	DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`
	/*
	  Indicates if the disk is undergoing migration to another container.
	*/
	IsMigrationInProgress *bool `json:"isMigrationInProgress,omitempty"`

	StorageConfig *VmDiskStorageConfig `json:"storageConfig,omitempty"`

	StorageContainer *VmDiskContainerReference `json:"storageContainer,omitempty"`
}

func (p *VmDisk) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDisk

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

func (p *VmDisk) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDisk
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDisk(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dataSource")
	delete(allFields, "diskExtId")
	delete(allFields, "diskSizeBytes")
	delete(allFields, "isMigrationInProgress")
	delete(allFields, "storageConfig")
	delete(allFields, "storageContainer")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDisk() *VmDisk {
	p := new(VmDisk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDisk"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This reference is for disk level storage container preference. This preference specifies the storage container to which this disk belongs.
*/
type VmDiskContainerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM disk container. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *VmDiskContainerReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskContainerReference

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

func (p *VmDiskContainerReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskContainerReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskContainerReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskContainerReference() *VmDiskContainerReference {
	p := new(VmDiskContainerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskContainerReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to a Disk Recovery Point of a VM disk.
*/
type VmDiskRecoveryPointReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  An external identifier of the Disk Recovery Point for a VM disk. It should be of type UUID.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
	/*
	  A globally unique identifier of a VM recovery point. It should be of type UUID.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
}

func (p *VmDiskRecoveryPointReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskRecoveryPointReference

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

func (p *VmDiskRecoveryPointReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskRecoveryPointReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskRecoveryPointReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskRecoveryPointExtId")
	delete(allFields, "vmRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskRecoveryPointReference() *VmDiskRecoveryPointReference {
	p := new(VmDiskRecoveryPointReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskRecoveryPointReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an existing VM disk.
*/
type VmDiskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DiskAddress *DiskAddress `json:"diskAddress,omitempty"`
	/*
	  A globally unique identifier of a VM disk of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`

	VmReference *VmReference `json:"vmReference,omitempty"`
}

func (p *VmDiskReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskReference

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

func (p *VmDiskReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskAddress")
	delete(allFields, "diskExtId")
	delete(allFields, "vmReference")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskReference() *VmDiskReference {
	p := new(VmDiskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Storage configuration for VM disks.
*/
type VmDiskStorageConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the virtual disk is pinned to the hot tier or not.
	*/
	IsFlashModeEnabled *bool `json:"isFlashModeEnabled,omitempty"`
}

func (p *VmDiskStorageConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskStorageConfig

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

func (p *VmDiskStorageConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskStorageConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskStorageConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isFlashModeEnabled")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskStorageConfig() *VmDiskStorageConfig {
	p := new(VmDiskStorageConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmDiskStorageConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The destination host to which the VM is migrated.
*/
type VmMigrateToHostParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Host *HostReference `json:"host"`
}

func (p *VmMigrateToHostParams) MarshalJSON() ([]byte, error) {
	type VmMigrateToHostParamsProxy VmMigrateToHostParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VmMigrateToHostParamsProxy
		Host *HostReference `json:"host,omitempty"`
	}{
		VmMigrateToHostParamsProxy: (*VmMigrateToHostParamsProxy)(p),
		Host:                       p.Host,
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

func (p *VmMigrateToHostParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmMigrateToHostParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmMigrateToHostParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "host")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmMigrateToHostParams() *VmMigrateToHostParams {
	p := new(VmMigrateToHostParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmMigrateToHostParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents VM Recovery Point properties.
*/
type VmRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ApplicationConsistentPropertiesItemDiscriminator_ *string `json:"$applicationConsistentPropertiesItemDiscriminator,omitempty"`
	/*
	  User-defined application-consistent properties for the recovery point.
	*/
	ApplicationConsistentProperties *OneOfVmRecoveryPointApplicationConsistentProperties `json:"applicationConsistentProperties,omitempty"`
	/*
	  External identifier of the Consistency group which the VM was part of at the time of recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DiskRecoveryPoints []import5.DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vm *Vm `json:"vm,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is displayed as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as a part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
}

func (p *VmRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRecoveryPoint

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

func (p *VmRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRecoveryPoint(*known)

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
	delete(allFields, "vm")
	delete(allFields, "vmCategories")
	delete(allFields, "vmExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRecoveryPoint() *VmRecoveryPoint {
	p := new(VmRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This is a reference to a VM.
*/
type VmReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *VmReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmReference

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

func (p *VmReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmReference() *VmReference {
	p := new(VmReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The IPv4 address configurations containing fields that can be overridden when restoring the VM.
*/
type VmRestoreIpv4ConfigOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import4.IPv4Address `json:"ipAddress,omitempty"`
	/*
	  Secondary IP addresses for the NIC.
	*/
	SecondaryIpAddressList []import4.IPv4Address `json:"secondaryIpAddressList,omitempty"`
}

func (p *VmRestoreIpv4ConfigOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRestoreIpv4ConfigOverrideSpec

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

func (p *VmRestoreIpv4ConfigOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRestoreIpv4ConfigOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRestoreIpv4ConfigOverrideSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "secondaryIpAddressList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRestoreIpv4ConfigOverrideSpec() *VmRestoreIpv4ConfigOverrideSpec {
	p := new(VmRestoreIpv4ConfigOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRestoreIpv4ConfigOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NIC configuration parameters that can be overridden when restoring the VM.
*/
type VmRestoreNicConfigOverrideParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	NicBackingInfoItemDiscriminator_ *string `json:"$nicBackingInfoItemDiscriminator,omitempty"`
	/*
	  Information about how NIC is associated with a VM.
	*/
	NicBackingInfo *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo `json:"nicBackingInfo,omitempty"`
	/*
	  External identifier of the NIC being overridden from the VM Recovery Point. This is not preserved but is only relevant for mapping the NIC being overridden.
	*/
	NicExtId *string `json:"nicExtId"`
	/*

	 */
	NicNetworkInfoItemDiscriminator_ *string `json:"$nicNetworkInfoItemDiscriminator,omitempty"`
	/*
	  Networking information object for a NIC.
	*/
	NicNetworkInfo *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo `json:"nicNetworkInfo,omitempty"`
}

func (p *VmRestoreNicConfigOverrideParams) MarshalJSON() ([]byte, error) {
	type VmRestoreNicConfigOverrideParamsProxy VmRestoreNicConfigOverrideParams

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VmRestoreNicConfigOverrideParamsProxy
		NicExtId *string `json:"nicExtId,omitempty"`
	}{
		VmRestoreNicConfigOverrideParamsProxy: (*VmRestoreNicConfigOverrideParamsProxy)(p),
		NicExtId:                              p.NicExtId,
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

func (p *VmRestoreNicConfigOverrideParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRestoreNicConfigOverrideParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRestoreNicConfigOverrideParams(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$nicBackingInfoItemDiscriminator")
	delete(allFields, "nicBackingInfo")
	delete(allFields, "nicExtId")
	delete(allFields, "$nicNetworkInfoItemDiscriminator")
	delete(allFields, "nicNetworkInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRestoreNicConfigOverrideParams() *VmRestoreNicConfigOverrideParams {
	p := new(VmRestoreNicConfigOverrideParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRestoreNicConfigOverrideParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VmRestoreNicConfigOverrideParams) GetNicNetworkInfo() interface{} {
	if nil == p.NicNetworkInfo {
		return nil
	}
	return p.NicNetworkInfo.GetValue()
}

func (p *VmRestoreNicConfigOverrideParams) SetNicNetworkInfo(v interface{}) error {
	if nil == p.NicNetworkInfo {
		p.NicNetworkInfo = NewOneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo()
	}
	e := p.NicNetworkInfo.SetValue(v)
	if nil == e {
		if nil == p.NicNetworkInfoItemDiscriminator_ {
			p.NicNetworkInfoItemDiscriminator_ = new(string)
		}
		*p.NicNetworkInfoItemDiscriminator_ = *p.NicNetworkInfo.Discriminator
	}
	return e
}

/*
NIC configuration input for the restored VM. If not specified, the NICs in the VM Recovery Point will be restored.
*/
type VmRestoreNicConfigSpecification struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of NIC parameters to override the corresponding NICs captured in the VM Recovery Point. NIC external identifiers are utilized to map the NICs in this list to the NICs from the VM Recovery Point
	*/
	NicOverrideList []VmRestoreNicConfigOverrideParams `json:"nicOverrideList,omitempty"`
	/*
	  List of NIC UUIDs from the VM Recovery Point to be dropped when restoring the VM.
	*/
	NicRemoveList []string `json:"nicRemoveList,omitempty"`
}

func (p *VmRestoreNicConfigSpecification) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRestoreNicConfigSpecification

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

func (p *VmRestoreNicConfigSpecification) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRestoreNicConfigSpecification
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRestoreNicConfigSpecification(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nicOverrideList")
	delete(allFields, "nicRemoveList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRestoreNicConfigSpecification() *VmRestoreNicConfigSpecification {
	p := new(VmRestoreNicConfigSpecification)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRestoreNicConfigSpecification"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network information object for a Virtual Ethernet NIC, contains fields that can be overridden when restoring the VM.
*/
type VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4Config *VmRestoreIpv4ConfigOverrideSpec `json:"ipv4Config,omitempty"`

	Subnet *SubnetReference `json:"subnet,omitempty"`
}

func (p *VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec

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

func (p *VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipv4Config")
	delete(allFields, "subnet")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRestoreVirtualEthernetNicNetworkInfoOverrideSpec() *VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec {
	p := new(VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines a Virtual Ethernet NIC, contains fields that can be overridden when restoring the VM.
*/
type VmRestoreVirtualEthernetNicOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`
}

func (p *VmRestoreVirtualEthernetNicOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRestoreVirtualEthernetNicOverrideSpec

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

func (p *VmRestoreVirtualEthernetNicOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRestoreVirtualEthernetNicOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmRestoreVirtualEthernetNicOverrideSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "macAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmRestoreVirtualEthernetNicOverrideSpec() *VmRestoreVirtualEthernetNicOverrideSpec {
	p := new(VmRestoreVirtualEthernetNicOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmRestoreVirtualEthernetNicOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an entity from which the VM should be cloned or created.
*/
type VmSourceReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *VmSourceReferenceEntityType `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of a VM of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *VmSourceReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmSourceReference

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

func (p *VmSourceReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmSourceReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmSourceReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityType")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmSourceReference() *VmSourceReference {
	p := new(VmSourceReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VmSourceReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an entity from which the VM should be cloned or created.
*/
type VmSourceReferenceEntityType int

const (
	VMSOURCEREFERENCEENTITYTYPE_UNKNOWN           VmSourceReferenceEntityType = 0
	VMSOURCEREFERENCEENTITYTYPE_REDACTED          VmSourceReferenceEntityType = 1
	VMSOURCEREFERENCEENTITYTYPE_VM                VmSourceReferenceEntityType = 2
	VMSOURCEREFERENCEENTITYTYPE_VM_RECOVERY_POINT VmSourceReferenceEntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VmSourceReferenceEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VM_RECOVERY_POINT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VmSourceReferenceEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VM_RECOVERY_POINT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VmSourceReferenceEntityType) index(name string) VmSourceReferenceEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VM_RECOVERY_POINT",
	}
	for idx := range names {
		if names[idx] == name {
			return VmSourceReferenceEntityType(idx)
		}
	}
	return VMSOURCEREFERENCEENTITYTYPE_UNKNOWN
}

func (e *VmSourceReferenceEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VmSourceReferenceEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VmSourceReferenceEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VmSourceReferenceEntityType) Ref() *VmSourceReferenceEntityType {
	return &e
}

/*
Reference to a disk Recovery Point of a Volume disk.
*/
type VolumeDiskRecoveryPointReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  An external identifier of the Volume disk Recovery Point. It should be of type UUID.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier of the Volume Group Recovery Point. It should be of type UUID.
	*/
	VolumeGroupRecoveryPointExtId *string `json:"volumeGroupRecoveryPointExtId,omitempty"`
}

func (p *VolumeDiskRecoveryPointReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeDiskRecoveryPointReference

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

func (p *VolumeDiskRecoveryPointReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeDiskRecoveryPointReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeDiskRecoveryPointReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskRecoveryPointExtId")
	delete(allFields, "volumeGroupRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeDiskRecoveryPointReference() *VolumeDiskRecoveryPointReference {
	p := new(VolumeDiskRecoveryPointReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VolumeDiskRecoveryPointReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to a Volume Disk.
*/
type VolumeDiskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  An external identifier of the Volume disk. It should be of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
	/*
	  The external identifier of the Volume Group. It should be of type UUID.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func (p *VolumeDiskReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeDiskReference

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

func (p *VolumeDiskReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeDiskReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeDiskReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskExtId")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeDiskReference() *VolumeDiskReference {
	p := new(VolumeDiskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VolumeDiskReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates how the vTPM for the VM should be configured.
*/
type VtpmConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the virtual trusted platform module is enabled for the Guest OS or not.
	*/
	IsVtpmEnabled *bool `json:"isVtpmEnabled,omitempty"`
	/*
	  Virtual trusted platform module version.
	*/
	Version *string `json:"version,omitempty"`

	VtpmDevice *VtpmDevice `json:"vtpmDevice,omitempty"`
}

func (p *VtpmConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VtpmConfig

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

func (p *VtpmConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VtpmConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VtpmConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isVtpmEnabled")
	delete(allFields, "version")
	delete(allFields, "vtpmDevice")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVtpmConfig() *VtpmConfig {
	p := new(VtpmConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VtpmConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A disk for the Virtual trusted platform module.
*/
type VtpmDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of a VM disk of type UUID.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
}

func (p *VtpmDevice) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VtpmDevice

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

func (p *VtpmDevice) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VtpmDevice
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VtpmDevice(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVtpmDevice() *VtpmDevice {
	p := new(VtpmDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.VtpmDevice"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfInstallVmGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfInstallVmGuestToolsApiResponseData() *OneOfInstallVmGuestToolsApiResponseData {
	p := new(OneOfInstallVmGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInstallVmGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInstallVmGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfInstallVmGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInstallVmGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInstallVmGuestToolsApiResponseData"))
}

func (p *OneOfInstallVmGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInstallVmGuestToolsApiResponseData")
}

type OneOfCreatePcieDeviceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreatePcieDeviceApiResponseData() *OneOfCreatePcieDeviceApiResponseData {
	p := new(OneOfCreatePcieDeviceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreatePcieDeviceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreatePcieDeviceApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreatePcieDeviceApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreatePcieDeviceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreatePcieDeviceApiResponseData"))
}

func (p *OneOfCreatePcieDeviceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreatePcieDeviceApiResponseData")
}

type OneOfCrossClusterMigrateVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCrossClusterMigrateVmApiResponseData() *OneOfCrossClusterMigrateVmApiResponseData {
	p := new(OneOfCrossClusterMigrateVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCrossClusterMigrateVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCrossClusterMigrateVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCrossClusterMigrateVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCrossClusterMigrateVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCrossClusterMigrateVmApiResponseData"))
}

func (p *OneOfCrossClusterMigrateVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCrossClusterMigrateVmApiResponseData")
}

type OneOfCloneVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCloneVmApiResponseData() *OneOfCloneVmApiResponseData {
	p := new(OneOfCloneVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloneVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloneVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCloneVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCloneVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloneVmApiResponseData"))
}

func (p *OneOfCloneVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCloneVmApiResponseData")
}

type OneOfCreateGpuApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateGpuApiResponseData() *OneOfCreateGpuApiResponseData {
	p := new(OneOfCreateGpuApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateGpuApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateGpuApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateGpuApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateGpuApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateGpuApiResponseData"))
}

func (p *OneOfCreateGpuApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateGpuApiResponseData")
}

type OneOfResetVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfResetVmApiResponseData() *OneOfResetVmApiResponseData {
	p := new(OneOfResetVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfResetVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfResetVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfResetVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfResetVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetVmApiResponseData"))
}

func (p *OneOfResetVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfResetVmApiResponseData")
}

type OneOfListSerialPortsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []SerialPort           `json:"-"`
}

func NewOneOfListSerialPortsApiResponseData() *OneOfListSerialPortsApiResponseData {
	p := new(OneOfListSerialPortsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSerialPortsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSerialPortsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []SerialPort:
		p.oneOfType2001 = v.([]SerialPort)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.SerialPort>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.SerialPort>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSerialPortsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.SerialPort>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListSerialPortsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]SerialPort)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.SerialPort" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.SerialPort>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.SerialPort>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSerialPortsApiResponseData"))
}

func (p *OneOfListSerialPortsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.SerialPort>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListSerialPortsApiResponseData")
}

type OneOfAssociateCategoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAssociateCategoriesApiResponseData() *OneOfAssociateCategoriesApiResponseData {
	p := new(OneOfAssociateCategoriesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssociateCategoriesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssociateCategoriesApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfAssociateCategoriesApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAssociateCategoriesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssociateCategoriesApiResponseData"))
}

func (p *OneOfAssociateCategoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAssociateCategoriesApiResponseData")
}

type OneOfCloneOverrideParamsBootConfig struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType2001 *LegacyBoot `json:"-"`
	oneOfType2002 *UefiBoot   `json:"-"`
}

func NewOneOfCloneOverrideParamsBootConfig() *OneOfCloneOverrideParamsBootConfig {
	p := new(OneOfCloneOverrideParamsBootConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloneOverrideParamsBootConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloneOverrideParamsBootConfig is nil"))
	}
	switch v.(type) {
	case LegacyBoot:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LegacyBoot)
		}
		*p.oneOfType2001 = v.(LegacyBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case UefiBoot:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(UefiBoot)
		}
		*p.oneOfType2002 = v.(UefiBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCloneOverrideParamsBootConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfCloneOverrideParamsBootConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(LegacyBoot)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.LegacyBoot" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LegacyBoot)
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
	vOneOfType2002 := new(UefiBoot)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.UefiBoot" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(UefiBoot)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloneOverrideParamsBootConfig"))
}

func (p *OneOfCloneOverrideParamsBootConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfCloneOverrideParamsBootConfig")
}

type OneOfUpgradeVmGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpgradeVmGuestToolsApiResponseData() *OneOfUpgradeVmGuestToolsApiResponseData {
	p := new(OneOfUpgradeVmGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpgradeVmGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpgradeVmGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpgradeVmGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpgradeVmGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpgradeVmGuestToolsApiResponseData"))
}

func (p *OneOfUpgradeVmGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpgradeVmGuestToolsApiResponseData")
}

type OneOfDeletePcieDeviceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeletePcieDeviceApiResponseData() *OneOfDeletePcieDeviceApiResponseData {
	p := new(OneOfDeletePcieDeviceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeletePcieDeviceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeletePcieDeviceApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeletePcieDeviceApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeletePcieDeviceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeletePcieDeviceApiResponseData"))
}

func (p *OneOfDeletePcieDeviceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeletePcieDeviceApiResponseData")
}

type OneOfDeleteVmRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVmRecoveryPointApiResponseData() *OneOfDeleteVmRecoveryPointApiResponseData {
	p := new(OneOfDeleteVmRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVmRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVmRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteVmRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVmRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVmRecoveryPointApiResponseData"))
}

func (p *OneOfDeleteVmRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVmRecoveryPointApiResponseData")
}

type OneOfReleaseIpApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfReleaseIpApiResponseData() *OneOfReleaseIpApiResponseData {
	p := new(OneOfReleaseIpApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReleaseIpApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReleaseIpApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfReleaseIpApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfReleaseIpApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReleaseIpApiResponseData"))
}

func (p *OneOfReleaseIpApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfReleaseIpApiResponseData")
}

type OneOfMigrateNicApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfMigrateNicApiResponseData() *OneOfMigrateNicApiResponseData {
	p := new(OneOfMigrateNicApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrateNicApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrateNicApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfMigrateNicApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrateNicApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateNicApiResponseData"))
}

func (p *OneOfMigrateNicApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrateNicApiResponseData")
}

type OneOfPowerCycleVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfPowerCycleVmApiResponseData() *OneOfPowerCycleVmApiResponseData {
	p := new(OneOfPowerCycleVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerCycleVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerCycleVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfPowerCycleVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerCycleVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerCycleVmApiResponseData"))
}

func (p *OneOfPowerCycleVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerCycleVmApiResponseData")
}

type OneOfEjectCdRomApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfEjectCdRomApiResponseData() *OneOfEjectCdRomApiResponseData {
	p := new(OneOfEjectCdRomApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEjectCdRomApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEjectCdRomApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfEjectCdRomApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEjectCdRomApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEjectCdRomApiResponseData"))
}

func (p *OneOfEjectCdRomApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEjectCdRomApiResponseData")
}

type OneOfGetSerialPortApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *SerialPort            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSerialPortApiResponseData() *OneOfGetSerialPortApiResponseData {
	p := new(OneOfGetSerialPortApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSerialPortApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSerialPortApiResponseData is nil"))
	}
	switch v.(type) {
	case SerialPort:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(SerialPort)
		}
		*p.oneOfType2001 = v.(SerialPort)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetSerialPortApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSerialPortApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(SerialPort)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.SerialPort" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(SerialPort)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSerialPortApiResponseData"))
}

func (p *OneOfGetSerialPortApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSerialPortApiResponseData")
}

type OneOfListDisksApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Disk                 `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListDisksApiResponseData() *OneOfListDisksApiResponseData {
	p := new(OneOfListDisksApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDisksApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDisksApiResponseData is nil"))
	}
	switch v.(type) {
	case []Disk:
		p.oneOfType2001 = v.([]Disk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Disk>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Disk>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfListDisksApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.Disk>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListDisksApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Disk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Disk" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Disk>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Disk>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDisksApiResponseData"))
}

func (p *OneOfListDisksApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.Disk>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListDisksApiResponseData")
}

type OneOfListVmRecoveryPointsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []VmRecoveryPoint      `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListVmRecoveryPointsApiResponseData() *OneOfListVmRecoveryPointsApiResponseData {
	p := new(OneOfListVmRecoveryPointsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmRecoveryPointsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmRecoveryPointsApiResponseData is nil"))
	}
	switch v.(type) {
	case []VmRecoveryPoint:
		p.oneOfType2001 = v.([]VmRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.VmRecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.VmRecoveryPoint>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfListVmRecoveryPointsApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.VmRecoveryPoint>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmRecoveryPointsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]VmRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.VmRecoveryPoint" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.VmRecoveryPoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.VmRecoveryPoint>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmRecoveryPointsApiResponseData"))
}

func (p *OneOfListVmRecoveryPointsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.VmRecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmRecoveryPointsApiResponseData")
}

type OneOfVmBootConfig struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType2001 *LegacyBoot `json:"-"`
	oneOfType2002 *UefiBoot   `json:"-"`
}

func NewOneOfVmBootConfig() *OneOfVmBootConfig {
	p := new(OneOfVmBootConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmBootConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmBootConfig is nil"))
	}
	switch v.(type) {
	case LegacyBoot:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LegacyBoot)
		}
		*p.oneOfType2001 = v.(LegacyBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case UefiBoot:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(UefiBoot)
		}
		*p.oneOfType2002 = v.(UefiBoot)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVmBootConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfVmBootConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(LegacyBoot)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.LegacyBoot" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LegacyBoot)
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
	vOneOfType2002 := new(UefiBoot)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.UefiBoot" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(UefiBoot)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmBootConfig"))
}

func (p *OneOfVmBootConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfVmBootConfig")
}

type OneOfVmRecoveryPointApplicationConsistentProperties struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import5.VssProperties `json:"-"`
}

func NewOneOfVmRecoveryPointApplicationConsistentProperties() *OneOfVmRecoveryPointApplicationConsistentProperties {
	p := new(OneOfVmRecoveryPointApplicationConsistentProperties)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRecoveryPointApplicationConsistentProperties is nil"))
	}
	switch v.(type) {
	case import5.VssProperties:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import5.VssProperties)
		}
		*p.oneOfType2001 = v.(import5.VssProperties)
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

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import5.VssProperties)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.common.VssProperties" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import5.VssProperties)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRecoveryPointApplicationConsistentProperties"))
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfVmRecoveryPointApplicationConsistentProperties")
}

type OneOfDataSourceReference struct {
	Discriminator *string                           `json:"-"`
	ObjectType_   *string                           `json:"-"`
	oneOfType2002 *VmDiskReference                  `json:"-"`
	oneOfType2004 *VmDiskRecoveryPointReference     `json:"-"`
	oneOfType2005 *VolumeDiskRecoveryPointReference `json:"-"`
	oneOfType2001 *ImageReference                   `json:"-"`
	oneOfType2003 *VolumeDiskReference              `json:"-"`
}

func NewOneOfDataSourceReference() *OneOfDataSourceReference {
	p := new(OneOfDataSourceReference)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDataSourceReference) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDataSourceReference is nil"))
	}
	switch v.(type) {
	case VmDiskReference:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(VmDiskReference)
		}
		*p.oneOfType2002 = v.(VmDiskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case VmDiskRecoveryPointReference:
		if nil == p.oneOfType2004 {
			p.oneOfType2004 = new(VmDiskRecoveryPointReference)
		}
		*p.oneOfType2004 = v.(VmDiskRecoveryPointReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2004.ObjectType_
	case VolumeDiskRecoveryPointReference:
		if nil == p.oneOfType2005 {
			p.oneOfType2005 = new(VolumeDiskRecoveryPointReference)
		}
		*p.oneOfType2005 = v.(VolumeDiskRecoveryPointReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2005.ObjectType_
	case ImageReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ImageReference)
		}
		*p.oneOfType2001 = v.(ImageReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case VolumeDiskReference:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(VolumeDiskReference)
		}
		*p.oneOfType2003 = v.(VolumeDiskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDataSourceReference) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2004
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2005
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	return nil
}

func (p *OneOfDataSourceReference) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(VmDiskReference)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.VmDiskReference" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(VmDiskReference)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2004 := new(VmDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if "vmm.v4.ahv.config.VmDiskRecoveryPointReference" == *vOneOfType2004.ObjectType_ {
			if nil == p.oneOfType2004 {
				p.oneOfType2004 = new(VmDiskRecoveryPointReference)
			}
			*p.oneOfType2004 = *vOneOfType2004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2004.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2004.ObjectType_
			return nil
		}
	}
	vOneOfType2005 := new(VolumeDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType2005); err == nil {
		if "vmm.v4.ahv.config.VolumeDiskRecoveryPointReference" == *vOneOfType2005.ObjectType_ {
			if nil == p.oneOfType2005 {
				p.oneOfType2005 = new(VolumeDiskRecoveryPointReference)
			}
			*p.oneOfType2005 = *vOneOfType2005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2005.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(ImageReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.ImageReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ImageReference)
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
	vOneOfType2003 := new(VolumeDiskReference)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "vmm.v4.ahv.config.VolumeDiskReference" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(VolumeDiskReference)
			}
			*p.oneOfType2003 = *vOneOfType2003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2003.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDataSourceReference"))
}

func (p *OneOfDataSourceReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2005)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	return nil, errors.New("No value to marshal for OneOfDataSourceReference")
}

type OneOfCloudInitCloudInitScript struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2001 *Userdata        `json:"-"`
	oneOfType2002 *CustomKeyValues `json:"-"`
}

func NewOneOfCloudInitCloudInitScript() *OneOfCloudInitCloudInitScript {
	p := new(OneOfCloudInitCloudInitScript)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloudInitCloudInitScript) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloudInitCloudInitScript is nil"))
	}
	switch v.(type) {
	case Userdata:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Userdata)
		}
		*p.oneOfType2001 = v.(Userdata)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case CustomKeyValues:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CustomKeyValues)
		}
		*p.oneOfType2002 = v.(CustomKeyValues)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCloudInitCloudInitScript) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfCloudInitCloudInitScript) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Userdata)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Userdata" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Userdata)
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
	vOneOfType2002 := new(CustomKeyValues)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CustomKeyValues" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CustomKeyValues)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloudInitCloudInitScript"))
}

func (p *OneOfCloudInitCloudInitScript) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfCloudInitCloudInitScript")
}

type OneOfDiskBackingInfo struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType2001 *VmDisk                   `json:"-"`
	oneOfType2002 *ADSFVolumeGroupReference `json:"-"`
}

func NewOneOfDiskBackingInfo() *OneOfDiskBackingInfo {
	p := new(OneOfDiskBackingInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiskBackingInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiskBackingInfo is nil"))
	}
	switch v.(type) {
	case VmDisk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmDisk)
		}
		*p.oneOfType2001 = v.(VmDisk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case ADSFVolumeGroupReference:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(ADSFVolumeGroupReference)
		}
		*p.oneOfType2002 = v.(ADSFVolumeGroupReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDiskBackingInfo) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfDiskBackingInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmDisk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.VmDisk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmDisk)
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
	vOneOfType2002 := new(ADSFVolumeGroupReference)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.ADSFVolumeGroupReference" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(ADSFVolumeGroupReference)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiskBackingInfo"))
}

func (p *OneOfDiskBackingInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfDiskBackingInfo")
}

type OneOfCreateDiskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateDiskApiResponseData() *OneOfCreateDiskApiResponseData {
	p := new(OneOfCreateDiskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDiskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDiskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateDiskApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateDiskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDiskApiResponseData"))
}

func (p *OneOfCreateDiskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDiskApiResponseData")
}

type OneOfAssignVmOwnerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAssignVmOwnerApiResponseData() *OneOfAssignVmOwnerApiResponseData {
	p := new(OneOfAssignVmOwnerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssignVmOwnerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssignVmOwnerApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfAssignVmOwnerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAssignVmOwnerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssignVmOwnerApiResponseData"))
}

func (p *OneOfAssignVmOwnerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAssignVmOwnerApiResponseData")
}

type OneOfDeleteNicApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteNicApiResponseData() *OneOfDeleteNicApiResponseData {
	p := new(OneOfDeleteNicApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteNicApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteNicApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteNicApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteNicApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNicApiResponseData"))
}

func (p *OneOfDeleteNicApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteNicApiResponseData")
}

type OneOfGetCdRomApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *CdRom                 `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetCdRomApiResponseData() *OneOfGetCdRomApiResponseData {
	p := new(OneOfGetCdRomApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCdRomApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCdRomApiResponseData is nil"))
	}
	switch v.(type) {
	case CdRom:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(CdRom)
		}
		*p.oneOfType2001 = v.(CdRom)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetCdRomApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetCdRomApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(CdRom)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.CdRom" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(CdRom)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCdRomApiResponseData"))
}

func (p *OneOfGetCdRomApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetCdRomApiResponseData")
}

type OneOfInsertCdRomApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfInsertCdRomApiResponseData() *OneOfInsertCdRomApiResponseData {
	p := new(OneOfInsertCdRomApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInsertCdRomApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInsertCdRomApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfInsertCdRomApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInsertCdRomApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInsertCdRomApiResponseData"))
}

func (p *OneOfInsertCdRomApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInsertCdRomApiResponseData")
}

type OneOfSysprepSysprepScript struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2001 *Unattendxml     `json:"-"`
	oneOfType2002 *CustomKeyValues `json:"-"`
}

func NewOneOfSysprepSysprepScript() *OneOfSysprepSysprepScript {
	p := new(OneOfSysprepSysprepScript)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSysprepSysprepScript) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSysprepSysprepScript is nil"))
	}
	switch v.(type) {
	case Unattendxml:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Unattendxml)
		}
		*p.oneOfType2001 = v.(Unattendxml)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case CustomKeyValues:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CustomKeyValues)
		}
		*p.oneOfType2002 = v.(CustomKeyValues)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSysprepSysprepScript) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfSysprepSysprepScript) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Unattendxml)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Unattendxml" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Unattendxml)
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
	vOneOfType2002 := new(CustomKeyValues)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CustomKeyValues" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CustomKeyValues)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSysprepSysprepScript"))
}

func (p *OneOfSysprepSysprepScript) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfSysprepSysprepScript")
}

type OneOfUpdateDiskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateDiskApiResponseData() *OneOfUpdateDiskApiResponseData {
	p := new(OneOfUpdateDiskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDiskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDiskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateDiskApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateDiskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDiskApiResponseData"))
}

func (p *OneOfUpdateDiskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDiskApiResponseData")
}

type OneOfUpdateGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateGuestToolsApiResponseData() *OneOfUpdateGuestToolsApiResponseData {
	p := new(OneOfUpdateGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateGuestToolsApiResponseData"))
}

func (p *OneOfUpdateGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateGuestToolsApiResponseData")
}

type OneOfGetGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *GuestTools            `json:"-"`
}

func NewOneOfGetGuestToolsApiResponseData() *OneOfGetGuestToolsApiResponseData {
	p := new(OneOfGetGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case GuestTools:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(GuestTools)
		}
		*p.oneOfType2001 = v.(GuestTools)
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

func (p *OneOfGetGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(GuestTools)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.GuestTools" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(GuestTools)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGuestToolsApiResponseData"))
}

func (p *OneOfGetGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetGuestToolsApiResponseData")
}

type OneOfUninstallVmGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUninstallVmGuestToolsApiResponseData() *OneOfUninstallVmGuestToolsApiResponseData {
	p := new(OneOfUninstallVmGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUninstallVmGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUninstallVmGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUninstallVmGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUninstallVmGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUninstallVmGuestToolsApiResponseData"))
}

func (p *OneOfUninstallVmGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUninstallVmGuestToolsApiResponseData")
}

type OneOfDisassociateCategoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDisassociateCategoriesApiResponseData() *OneOfDisassociateCategoriesApiResponseData {
	p := new(OneOfDisassociateCategoriesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDisassociateCategoriesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDisassociateCategoriesApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDisassociateCategoriesApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDisassociateCategoriesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisassociateCategoriesApiResponseData"))
}

func (p *OneOfDisassociateCategoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDisassociateCategoriesApiResponseData")
}

type OneOfLegacyBootBootDevice struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType2001 *BootDeviceDisk `json:"-"`
	oneOfType2002 *BootDeviceNic  `json:"-"`
}

func NewOneOfLegacyBootBootDevice() *OneOfLegacyBootBootDevice {
	p := new(OneOfLegacyBootBootDevice)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLegacyBootBootDevice) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLegacyBootBootDevice is nil"))
	}
	switch v.(type) {
	case BootDeviceDisk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(BootDeviceDisk)
		}
		*p.oneOfType2001 = v.(BootDeviceDisk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case BootDeviceNic:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(BootDeviceNic)
		}
		*p.oneOfType2002 = v.(BootDeviceNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfLegacyBootBootDevice) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfLegacyBootBootDevice) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(BootDeviceDisk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.BootDeviceDisk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(BootDeviceDisk)
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
	vOneOfType2002 := new(BootDeviceNic)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.BootDeviceNic" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(BootDeviceNic)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLegacyBootBootDevice"))
}

func (p *OneOfLegacyBootBootDevice) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfLegacyBootBootDevice")
}

type OneOfGetPcieDeviceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *PcieDevice            `json:"-"`
}

func NewOneOfGetPcieDeviceApiResponseData() *OneOfGetPcieDeviceApiResponseData {
	p := new(OneOfGetPcieDeviceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetPcieDeviceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetPcieDeviceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case PcieDevice:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(PcieDevice)
		}
		*p.oneOfType2001 = v.(PcieDevice)
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

func (p *OneOfGetPcieDeviceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetPcieDeviceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(PcieDevice)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.PcieDevice" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(PcieDevice)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetPcieDeviceApiResponseData"))
}

func (p *OneOfGetPcieDeviceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetPcieDeviceApiResponseData")
}

type OneOfGetVmRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VmRecoveryPoint       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetVmRecoveryPointApiResponseData() *OneOfGetVmRecoveryPointApiResponseData {
	p := new(OneOfGetVmRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case VmRecoveryPoint:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmRecoveryPoint)
		}
		*p.oneOfType2001 = v.(VmRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetVmRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.VmRecoveryPoint" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmRecoveryPoint)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmRecoveryPointApiResponseData"))
}

func (p *OneOfGetVmRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmRecoveryPointApiResponseData")
}

type OneOfGetDiskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *Disk                  `json:"-"`
}

func NewOneOfGetDiskApiResponseData() *OneOfGetDiskApiResponseData {
	p := new(OneOfGetDiskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDiskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDiskApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Disk:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Disk)
		}
		*p.oneOfType2001 = v.(Disk)
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

func (p *OneOfGetDiskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetDiskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(Disk)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Disk" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Disk)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDiskApiResponseData"))
}

func (p *OneOfGetDiskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetDiskApiResponseData")
}

type OneOfCustomizeGuestVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCustomizeGuestVmApiResponseData() *OneOfCustomizeGuestVmApiResponseData {
	p := new(OneOfCustomizeGuestVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCustomizeGuestVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCustomizeGuestVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCustomizeGuestVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCustomizeGuestVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCustomizeGuestVmApiResponseData"))
}

func (p *OneOfCustomizeGuestVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCustomizeGuestVmApiResponseData")
}

type OneOfPowerOffVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfPowerOffVmApiResponseData() *OneOfPowerOffVmApiResponseData {
	p := new(OneOfPowerOffVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerOffVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerOffVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfPowerOffVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerOffVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerOffVmApiResponseData"))
}

func (p *OneOfPowerOffVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerOffVmApiResponseData")
}

type OneOfInsertVmGuestToolsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfInsertVmGuestToolsApiResponseData() *OneOfInsertVmGuestToolsApiResponseData {
	p := new(OneOfInsertVmGuestToolsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInsertVmGuestToolsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInsertVmGuestToolsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfInsertVmGuestToolsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfInsertVmGuestToolsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInsertVmGuestToolsApiResponseData"))
}

func (p *OneOfInsertVmGuestToolsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfInsertVmGuestToolsApiResponseData")
}

type OneOfNicNicBackingInfo struct {
	Discriminator *string             `json:"-"`
	ObjectType_   *string             `json:"-"`
	oneOfType2051 *VirtualEthernetNic `json:"-"`
	oneOfType2052 *SriovNic           `json:"-"`
	oneOfType2053 *DpOffloadNic       `json:"-"`
}

func NewOneOfNicNicBackingInfo() *OneOfNicNicBackingInfo {
	p := new(OneOfNicNicBackingInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNicNicBackingInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNicNicBackingInfo is nil"))
	}
	switch v.(type) {
	case VirtualEthernetNic:
		if nil == p.oneOfType2051 {
			p.oneOfType2051 = new(VirtualEthernetNic)
		}
		*p.oneOfType2051 = v.(VirtualEthernetNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2051.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2051.ObjectType_
	case SriovNic:
		if nil == p.oneOfType2052 {
			p.oneOfType2052 = new(SriovNic)
		}
		*p.oneOfType2052 = v.(SriovNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2052.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2052.ObjectType_
	case DpOffloadNic:
		if nil == p.oneOfType2053 {
			p.oneOfType2053 = new(DpOffloadNic)
		}
		*p.oneOfType2053 = v.(DpOffloadNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2053.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2053.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNicNicBackingInfo) GetValue() interface{} {
	if p.oneOfType2051 != nil && *p.oneOfType2051.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2051
	}
	if p.oneOfType2052 != nil && *p.oneOfType2052.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2052
	}
	if p.oneOfType2053 != nil && *p.oneOfType2053.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2053
	}
	return nil
}

func (p *OneOfNicNicBackingInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2051 := new(VirtualEthernetNic)
	if err := json.Unmarshal(b, vOneOfType2051); err == nil {
		if "vmm.v4.ahv.config.VirtualEthernetNic" == *vOneOfType2051.ObjectType_ {
			if nil == p.oneOfType2051 {
				p.oneOfType2051 = new(VirtualEthernetNic)
			}
			*p.oneOfType2051 = *vOneOfType2051
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2051.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2051.ObjectType_
			return nil
		}
	}
	vOneOfType2052 := new(SriovNic)
	if err := json.Unmarshal(b, vOneOfType2052); err == nil {
		if "vmm.v4.ahv.config.SriovNic" == *vOneOfType2052.ObjectType_ {
			if nil == p.oneOfType2052 {
				p.oneOfType2052 = new(SriovNic)
			}
			*p.oneOfType2052 = *vOneOfType2052
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2052.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2052.ObjectType_
			return nil
		}
	}
	vOneOfType2053 := new(DpOffloadNic)
	if err := json.Unmarshal(b, vOneOfType2053); err == nil {
		if "vmm.v4.ahv.config.DpOffloadNic" == *vOneOfType2053.ObjectType_ {
			if nil == p.oneOfType2053 {
				p.oneOfType2053 = new(DpOffloadNic)
			}
			*p.oneOfType2053 = *vOneOfType2053
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2053.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2053.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNicNicBackingInfo"))
}

func (p *OneOfNicNicBackingInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2051 != nil && *p.oneOfType2051.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2051)
	}
	if p.oneOfType2052 != nil && *p.oneOfType2052.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2052)
	}
	if p.oneOfType2053 != nil && *p.oneOfType2053.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2053)
	}
	return nil, errors.New("No value to marshal for OneOfNicNicBackingInfo")
}

type OneOfPowerOnVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfPowerOnVmApiResponseData() *OneOfPowerOnVmApiResponseData {
	p := new(OneOfPowerOnVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPowerOnVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPowerOnVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfPowerOnVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfPowerOnVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPowerOnVmApiResponseData"))
}

func (p *OneOfPowerOnVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfPowerOnVmApiResponseData")
}

type OneOfUpdateVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmApiResponseData() *OneOfUpdateVmApiResponseData {
	p := new(OneOfUpdateVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmApiResponseData"))
}

func (p *OneOfUpdateVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmApiResponseData")
}

type OneOfDiskMigrationParamsMigrateDisks struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2002 *AllDisksMigrationPlan `json:"-"`
	oneOfType2001 *MigrationPlans        `json:"-"`
}

func NewOneOfDiskMigrationParamsMigrateDisks() *OneOfDiskMigrationParamsMigrateDisks {
	p := new(OneOfDiskMigrationParamsMigrateDisks)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiskMigrationParamsMigrateDisks) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiskMigrationParamsMigrateDisks is nil"))
	}
	switch v.(type) {
	case AllDisksMigrationPlan:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(AllDisksMigrationPlan)
		}
		*p.oneOfType2002 = v.(AllDisksMigrationPlan)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case MigrationPlans:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(MigrationPlans)
		}
		*p.oneOfType2001 = v.(MigrationPlans)
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

func (p *OneOfDiskMigrationParamsMigrateDisks) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDiskMigrationParamsMigrateDisks) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(AllDisksMigrationPlan)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.AllDisksMigrationPlan" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(AllDisksMigrationPlan)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(MigrationPlans)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.MigrationPlans" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(MigrationPlans)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiskMigrationParamsMigrateDisks"))
}

func (p *OneOfDiskMigrationParamsMigrateDisks) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDiskMigrationParamsMigrateDisks")
}

type OneOfGetGpuApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *Gpu                   `json:"-"`
}

func NewOneOfGetGpuApiResponseData() *OneOfGetGpuApiResponseData {
	p := new(OneOfGetGpuApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGpuApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGpuApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Gpu:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Gpu)
		}
		*p.oneOfType2001 = v.(Gpu)
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

func (p *OneOfGetGpuApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetGpuApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(Gpu)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Gpu" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Gpu)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGpuApiResponseData"))
}

func (p *OneOfGetGpuApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetGpuApiResponseData")
}

type OneOfMigrateVmDisksApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfMigrateVmDisksApiResponseData() *OneOfMigrateVmDisksApiResponseData {
	p := new(OneOfMigrateVmDisksApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrateVmDisksApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrateVmDisksApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfMigrateVmDisksApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrateVmDisksApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateVmDisksApiResponseData"))
}

func (p *OneOfMigrateVmDisksApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrateVmDisksApiResponseData")
}

type OneOfRestoreVmRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRestoreVmRecoveryPointApiResponseData() *OneOfRestoreVmRecoveryPointApiResponseData {
	p := new(OneOfRestoreVmRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRestoreVmRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRestoreVmRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRestoreVmRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRestoreVmRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRestoreVmRecoveryPointApiResponseData"))
}

func (p *OneOfRestoreVmRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRestoreVmRecoveryPointApiResponseData")
}

type OneOfDeleteCdRomApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteCdRomApiResponseData() *OneOfDeleteCdRomApiResponseData {
	p := new(OneOfDeleteCdRomApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteCdRomApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteCdRomApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteCdRomApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteCdRomApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteCdRomApiResponseData"))
}

func (p *OneOfDeleteCdRomApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteCdRomApiResponseData")
}

type OneOfDeleteSerialPortApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSerialPortApiResponseData() *OneOfDeleteSerialPortApiResponseData {
	p := new(OneOfDeleteSerialPortApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSerialPortApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSerialPortApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteSerialPortApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSerialPortApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSerialPortApiResponseData"))
}

func (p *OneOfDeleteSerialPortApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSerialPortApiResponseData")
}

type OneOfCreateVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateVmApiResponseData() *OneOfCreateVmApiResponseData {
	p := new(OneOfCreateVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVmApiResponseData"))
}

func (p *OneOfCreateVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVmApiResponseData")
}

type OneOfShutdownVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfShutdownVmApiResponseData() *OneOfShutdownVmApiResponseData {
	p := new(OneOfShutdownVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfShutdownVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfShutdownVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfShutdownVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfShutdownVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfShutdownVmApiResponseData"))
}

func (p *OneOfShutdownVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfShutdownVmApiResponseData")
}

type OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo struct {
	Discriminator *string                                             `json:"-"`
	ObjectType_   *string                                             `json:"-"`
	oneOfType2101 *VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec `json:"-"`
}

func NewOneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo() *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo {
	p := new(OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo is nil"))
	}
	switch v.(type) {
	case VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec)
		}
		*p.oneOfType2101 = v.(VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2101.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2101.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo) GetValue() interface{} {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2101 := new(VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "vmm.v4.ahv.config.VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(VmRestoreVirtualEthernetNicNetworkInfoOverrideSpec)
			}
			*p.oneOfType2101 = *vOneOfType2101
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2101.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2101.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo"))
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfVmRestoreNicConfigOverrideParamsNicNetworkInfo")
}

type OneOfRevertVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRevertVmApiResponseData() *OneOfRevertVmApiResponseData {
	p := new(OneOfRevertVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRevertVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRevertVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRevertVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRevertVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRevertVmApiResponseData"))
}

func (p *OneOfRevertVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRevertVmApiResponseData")
}

type OneOfDeleteVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVmApiResponseData() *OneOfDeleteVmApiResponseData {
	p := new(OneOfDeleteVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVmApiResponseData"))
}

func (p *OneOfDeleteVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVmApiResponseData")
}

type OneOfListGpusApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []Gpu                  `json:"-"`
}

func NewOneOfListGpusApiResponseData() *OneOfListGpusApiResponseData {
	p := new(OneOfListGpusApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListGpusApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListGpusApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Gpu:
		p.oneOfType2001 = v.([]Gpu)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Gpu>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Gpu>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListGpusApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.Gpu>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListGpusApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]Gpu)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Gpu" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Gpu>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Gpu>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListGpusApiResponseData"))
}

func (p *OneOfListGpusApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.Gpu>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListGpusApiResponseData")
}

type OneOfCreateNicApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateNicApiResponseData() *OneOfCreateNicApiResponseData {
	p := new(OneOfCreateNicApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNicApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNicApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateNicApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateNicApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNicApiResponseData"))
}

func (p *OneOfCreateNicApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNicApiResponseData")
}

type OneOfNicNicNetworkInfo struct {
	Discriminator *string                        `json:"-"`
	ObjectType_   *string                        `json:"-"`
	oneOfType2102 *SriovNicNetworkInfo           `json:"-"`
	oneOfType2101 *VirtualEthernetNicNetworkInfo `json:"-"`
	oneOfType2103 *DpOffloadNicNetworkInfo       `json:"-"`
}

func NewOneOfNicNicNetworkInfo() *OneOfNicNicNetworkInfo {
	p := new(OneOfNicNicNetworkInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNicNicNetworkInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNicNicNetworkInfo is nil"))
	}
	switch v.(type) {
	case SriovNicNetworkInfo:
		if nil == p.oneOfType2102 {
			p.oneOfType2102 = new(SriovNicNetworkInfo)
		}
		*p.oneOfType2102 = v.(SriovNicNetworkInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2102.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2102.ObjectType_
	case VirtualEthernetNicNetworkInfo:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(VirtualEthernetNicNetworkInfo)
		}
		*p.oneOfType2101 = v.(VirtualEthernetNicNetworkInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2101.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2101.ObjectType_
	case DpOffloadNicNetworkInfo:
		if nil == p.oneOfType2103 {
			p.oneOfType2103 = new(DpOffloadNicNetworkInfo)
		}
		*p.oneOfType2103 = v.(DpOffloadNicNetworkInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2103.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2103.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNicNicNetworkInfo) GetValue() interface{} {
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2102
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2103
	}
	return nil
}

func (p *OneOfNicNicNetworkInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2102 := new(SriovNicNetworkInfo)
	if err := json.Unmarshal(b, vOneOfType2102); err == nil {
		if "vmm.v4.ahv.config.SriovNicNetworkInfo" == *vOneOfType2102.ObjectType_ {
			if nil == p.oneOfType2102 {
				p.oneOfType2102 = new(SriovNicNetworkInfo)
			}
			*p.oneOfType2102 = *vOneOfType2102
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2102.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2102.ObjectType_
			return nil
		}
	}
	vOneOfType2101 := new(VirtualEthernetNicNetworkInfo)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "vmm.v4.ahv.config.VirtualEthernetNicNetworkInfo" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(VirtualEthernetNicNetworkInfo)
			}
			*p.oneOfType2101 = *vOneOfType2101
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2101.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2101.ObjectType_
			return nil
		}
	}
	vOneOfType2103 := new(DpOffloadNicNetworkInfo)
	if err := json.Unmarshal(b, vOneOfType2103); err == nil {
		if "vmm.v4.ahv.config.DpOffloadNicNetworkInfo" == *vOneOfType2103.ObjectType_ {
			if nil == p.oneOfType2103 {
				p.oneOfType2103 = new(DpOffloadNicNetworkInfo)
			}
			*p.oneOfType2103 = *vOneOfType2103
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2103.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2103.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNicNicNetworkInfo"))
}

func (p *OneOfNicNicNetworkInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2102)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2103)
	}
	return nil, errors.New("No value to marshal for OneOfNicNicNetworkInfo")
}

type OneOfUpdateSerialPortApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSerialPortApiResponseData() *OneOfUpdateSerialPortApiResponseData {
	p := new(OneOfUpdateSerialPortApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSerialPortApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSerialPortApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateSerialPortApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSerialPortApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSerialPortApiResponseData"))
}

func (p *OneOfUpdateSerialPortApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSerialPortApiResponseData")
}

type OneOfCreateCdRomApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateCdRomApiResponseData() *OneOfCreateCdRomApiResponseData {
	p := new(OneOfCreateCdRomApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCdRomApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCdRomApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateCdRomApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCdRomApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCdRomApiResponseData"))
}

func (p *OneOfCreateCdRomApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCdRomApiResponseData")
}

type OneOfRebootVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRebootVmApiResponseData() *OneOfRebootVmApiResponseData {
	p := new(OneOfRebootVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRebootVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRebootVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfRebootVmApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRebootVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRebootVmApiResponseData"))
}

func (p *OneOfRebootVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRebootVmApiResponseData")
}

type OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo struct {
	Discriminator *string                                  `json:"-"`
	ObjectType_   *string                                  `json:"-"`
	oneOfType2051 *VmRestoreVirtualEthernetNicOverrideSpec `json:"-"`
}

func NewOneOfVmRestoreNicConfigOverrideParamsNicBackingInfo() *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo {
	p := new(OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo is nil"))
	}
	switch v.(type) {
	case VmRestoreVirtualEthernetNicOverrideSpec:
		if nil == p.oneOfType2051 {
			p.oneOfType2051 = new(VmRestoreVirtualEthernetNicOverrideSpec)
		}
		*p.oneOfType2051 = v.(VmRestoreVirtualEthernetNicOverrideSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2051.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2051.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo) GetValue() interface{} {
	if p.oneOfType2051 != nil && *p.oneOfType2051.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2051
	}
	return nil
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2051 := new(VmRestoreVirtualEthernetNicOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType2051); err == nil {
		if "vmm.v4.ahv.config.VmRestoreVirtualEthernetNicOverrideSpec" == *vOneOfType2051.ObjectType_ {
			if nil == p.oneOfType2051 {
				p.oneOfType2051 = new(VmRestoreVirtualEthernetNicOverrideSpec)
			}
			*p.oneOfType2051 = *vOneOfType2051
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2051.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2051.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo"))
}

func (p *OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2051 != nil && *p.oneOfType2051.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2051)
	}
	return nil, errors.New("No value to marshal for OneOfVmRestoreNicConfigOverrideParamsNicBackingInfo")
}

type OneOfGetVmApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *Vm                    `json:"-"`
}

func NewOneOfGetVmApiResponseData() *OneOfGetVmApiResponseData {
	p := new(OneOfGetVmApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Vm:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Vm)
		}
		*p.oneOfType2001 = v.(Vm)
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

func (p *OneOfGetVmApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetVmApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Vm" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Vm)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmApiResponseData"))
}

func (p *OneOfGetVmApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmApiResponseData")
}

type OneOfUpdateNicApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateNicApiResponseData() *OneOfUpdateNicApiResponseData {
	p := new(OneOfUpdateNicApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNicApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNicApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfUpdateNicApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateNicApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNicApiResponseData"))
}

func (p *OneOfUpdateNicApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNicApiResponseData")
}

type OneOfDeleteDiskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteDiskApiResponseData() *OneOfDeleteDiskApiResponseData {
	p := new(OneOfDeleteDiskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDiskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDiskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteDiskApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteDiskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDiskApiResponseData"))
}

func (p *OneOfDeleteDiskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDiskApiResponseData")
}

type OneOfGuestCustomizationParamsConfig struct {
	Discriminator *string    `json:"-"`
	ObjectType_   *string    `json:"-"`
	oneOfType2001 *Sysprep   `json:"-"`
	oneOfType2002 *CloudInit `json:"-"`
}

func NewOneOfGuestCustomizationParamsConfig() *OneOfGuestCustomizationParamsConfig {
	p := new(OneOfGuestCustomizationParamsConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGuestCustomizationParamsConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGuestCustomizationParamsConfig is nil"))
	}
	switch v.(type) {
	case Sysprep:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Sysprep)
		}
		*p.oneOfType2001 = v.(Sysprep)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case CloudInit:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CloudInit)
		}
		*p.oneOfType2002 = v.(CloudInit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGuestCustomizationParamsConfig) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfGuestCustomizationParamsConfig) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Sysprep)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Sysprep" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Sysprep)
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
	vOneOfType2002 := new(CloudInit)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CloudInit" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CloudInit)
			}
			*p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2002.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGuestCustomizationParamsConfig"))
}

func (p *OneOfGuestCustomizationParamsConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfGuestCustomizationParamsConfig")
}

type OneOfDeleteGpuApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteGpuApiResponseData() *OneOfDeleteGpuApiResponseData {
	p := new(OneOfDeleteGpuApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteGpuApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteGpuApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfDeleteGpuApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteGpuApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteGpuApiResponseData"))
}

func (p *OneOfDeleteGpuApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteGpuApiResponseData")
}

type OneOfUefiBootBootDevice struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType2006 *BootDeviceDisk `json:"-"`
	oneOfType2007 *BootDeviceNic  `json:"-"`
}

func NewOneOfUefiBootBootDevice() *OneOfUefiBootBootDevice {
	p := new(OneOfUefiBootBootDevice)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUefiBootBootDevice) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUefiBootBootDevice is nil"))
	}
	switch v.(type) {
	case BootDeviceDisk:
		if nil == p.oneOfType2006 {
			p.oneOfType2006 = new(BootDeviceDisk)
		}
		*p.oneOfType2006 = v.(BootDeviceDisk)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2006.ObjectType_
	case BootDeviceNic:
		if nil == p.oneOfType2007 {
			p.oneOfType2007 = new(BootDeviceNic)
		}
		*p.oneOfType2007 = v.(BootDeviceNic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2007.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUefiBootBootDevice) GetValue() interface{} {
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2006
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2007
	}
	return nil
}

func (p *OneOfUefiBootBootDevice) UnmarshalJSON(b []byte) error {
	vOneOfType2006 := new(BootDeviceDisk)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if "vmm.v4.ahv.config.BootDeviceDisk" == *vOneOfType2006.ObjectType_ {
			if nil == p.oneOfType2006 {
				p.oneOfType2006 = new(BootDeviceDisk)
			}
			*p.oneOfType2006 = *vOneOfType2006
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2006.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2006.ObjectType_
			return nil
		}
	}
	vOneOfType2007 := new(BootDeviceNic)
	if err := json.Unmarshal(b, vOneOfType2007); err == nil {
		if "vmm.v4.ahv.config.BootDeviceNic" == *vOneOfType2007.ObjectType_ {
			if nil == p.oneOfType2007 {
				p.oneOfType2007 = new(BootDeviceNic)
			}
			*p.oneOfType2007 = *vOneOfType2007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2007.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2007.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUefiBootBootDevice"))
}

func (p *OneOfUefiBootBootDevice) MarshalJSON() ([]byte, error) {
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2007)
	}
	return nil, errors.New("No value to marshal for OneOfUefiBootBootDevice")
}

type OneOfListVmsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Vm                   `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListVmsApiResponseData() *OneOfListVmsApiResponseData {
	p := new(OneOfListVmsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Vm:
		p.oneOfType2001 = v.([]Vm)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Vm>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Vm>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfListVmsApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.config.Vm>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Vm" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Vm>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Vm>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmsApiResponseData"))
}

func (p *OneOfListVmsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.config.Vm>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmsApiResponseData")
}

type OneOfAssignIpApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAssignIpApiResponseData() *OneOfAssignIpApiResponseData {
	p := new(OneOfAssignIpApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssignIpApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssignIpApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfAssignIpApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAssignIpApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssignIpApiResponseData"))
}

func (p *OneOfAssignIpApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAssignIpApiResponseData")
}

type OneOfGetNicApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Nic                   `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetNicApiResponseData() *OneOfGetNicApiResponseData {
	p := new(OneOfGetNicApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNicApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNicApiResponseData is nil"))
	}
	switch v.(type) {
	case Nic:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Nic)
		}
		*p.oneOfType2001 = v.(Nic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetNicApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNicApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Nic)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Nic" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Nic)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNicApiResponseData"))
}

func (p *OneOfGetNicApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNicApiResponseData")
}

type OneOfPcieDeviceBackingInfo struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType2001 *PcieDeviceReference `json:"-"`
}

func NewOneOfPcieDeviceBackingInfo() *OneOfPcieDeviceBackingInfo {
	p := new(OneOfPcieDeviceBackingInfo)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPcieDeviceBackingInfo) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPcieDeviceBackingInfo is nil"))
	}
	switch v.(type) {
	case PcieDeviceReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(PcieDeviceReference)
		}
		*p.oneOfType2001 = v.(PcieDeviceReference)
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

func (p *OneOfPcieDeviceBackingInfo) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfPcieDeviceBackingInfo) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(PcieDeviceReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.PcieDeviceReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(PcieDeviceReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPcieDeviceBackingInfo"))
}

func (p *OneOfPcieDeviceBackingInfo) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfPcieDeviceBackingInfo")
}

type OneOfListPcieDevicesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []PcieDevice           `json:"-"`
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
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []PcieDevice:
		p.oneOfType2001 = v.([]PcieDevice)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.PcieDevice>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.PcieDevice>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListPcieDevicesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListPcieDevicesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]PcieDevice)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.PcieDevice" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.PcieDevice>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.PcieDevice>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListPcieDevicesApiResponseData"))
}

func (p *OneOfListPcieDevicesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.PcieDevice>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListPcieDevicesApiResponseData")
}

type OneOfMigrateVmToHostApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfMigrateVmToHostApiResponseData() *OneOfMigrateVmToHostApiResponseData {
	p := new(OneOfMigrateVmToHostApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrateVmToHostApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrateVmToHostApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfMigrateVmToHostApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrateVmToHostApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateVmToHostApiResponseData"))
}

func (p *OneOfMigrateVmToHostApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrateVmToHostApiResponseData")
}

type OneOfCreateVMRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateVMRecoveryPointApiResponseData() *OneOfCreateVMRecoveryPointApiResponseData {
	p := new(OneOfCreateVMRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVMRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVMRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateVMRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVMRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVMRecoveryPointApiResponseData"))
}

func (p *OneOfCreateVMRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVMRecoveryPointApiResponseData")
}

type OneOfListNicsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []Nic                  `json:"-"`
}

func NewOneOfListNicsApiResponseData() *OneOfListNicsApiResponseData {
	p := new(OneOfListNicsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNicsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNicsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Nic:
		p.oneOfType2001 = v.([]Nic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.Nic>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.Nic>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNicsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.Nic>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListNicsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]Nic)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.Nic" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.Nic>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.Nic>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNicsApiResponseData"))
}

func (p *OneOfListNicsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.Nic>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListNicsApiResponseData")
}

type OneOfCreateSerialPortApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateSerialPortApiResponseData() *OneOfCreateSerialPortApiResponseData {
	p := new(OneOfCreateSerialPortApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateSerialPortApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateSerialPortApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfCreateSerialPortApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateSerialPortApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSerialPortApiResponseData"))
}

func (p *OneOfCreateSerialPortApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSerialPortApiResponseData")
}

type OneOfListCdRomsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []CdRom                `json:"-"`
}

func NewOneOfListCdRomsApiResponseData() *OneOfListCdRomsApiResponseData {
	p := new(OneOfListCdRomsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCdRomsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCdRomsApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []CdRom:
		p.oneOfType2001 = v.([]CdRom)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.config.CdRom>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.config.CdRom>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCdRomsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.config.CdRom>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListCdRomsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]CdRom)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.config.CdRom" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.config.CdRom>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.config.CdRom>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCdRomsApiResponseData"))
}

func (p *OneOfListCdRomsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.config.CdRom>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListCdRomsApiResponseData")
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
