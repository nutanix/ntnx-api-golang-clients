/*
 * Generated file models/vmm/v4/ahv/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Data Protection APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.ahv.config of Nutanix Data Protection APIs
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/config"
)

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
The IPv4 address configurations containing fields that can be overridden when restoring the VM.
*/
type VmRestoreIpv4ConfigOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPv4Address `json:"ipAddress,omitempty"`
	/*
	  Secondary IP addresses for the NIC.
	*/
	SecondaryIpAddressList []import1.IPv4Address `json:"secondaryIpAddressList,omitempty"`
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
