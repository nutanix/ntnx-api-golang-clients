/*
 * Generated file models/vmm/v4/ahv/policies/policies_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.ahv.policies of Nutanix Virtual Machine Management APIs
*/
package policies

import (
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/*
Reference to a category.
*/
type CategoryReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the category.
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
	*p.ObjectType_ = "vmm.v4.ahv.policies.CategoryReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The VM-VM anti-affinity policy is not compliant for the VM as the cluster hosting the VM is not running on a supported AOS version.
*/
type ClusterNotSupportedForVmAntiAffinity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *ClusterNotSupportedForVmAntiAffinity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterNotSupportedForVmAntiAffinity

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

func (p *ClusterNotSupportedForVmAntiAffinity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterNotSupportedForVmAntiAffinity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterNotSupportedForVmAntiAffinity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterNotSupportedForVmAntiAffinity() *ClusterNotSupportedForVmAntiAffinity {
	p := new(ClusterNotSupportedForVmAntiAffinity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ClusterNotSupportedForVmAntiAffinity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to a cluster.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the cluster.
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
	*p.ObjectType_ = "vmm.v4.ahv.policies.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The compliance status of the VM is compliant.
*/
type CompliantVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *CompliantVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CompliantVmAntiAffinityPolicy

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

func (p *CompliantVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CompliantVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CompliantVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCompliantVmAntiAffinityPolicy() *CompliantVmAntiAffinityPolicy {
	p := new(CompliantVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.CompliantVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM is compliant with the VM-host affinity policy.
*/
type CompliantVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *CompliantVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CompliantVmHostAffinityPolicy

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

func (p *CompliantVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CompliantVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CompliantVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCompliantVmHostAffinityPolicy() *CompliantVmHostAffinityPolicy {
	p := new(CompliantVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.CompliantVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The current VM-VM anti-affinity policy is not applied on the VM as the VM is already part of a legacy VM Group.
*/
type ConflictingLegacyVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID of the legacy VM Group which the VM is the part of.
	*/
	LegacyVmAntiAffinityPolicyExtId *string `json:"legacyVmAntiAffinityPolicyExtId,omitempty"`
}

func (p *ConflictingLegacyVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConflictingLegacyVmAntiAffinityPolicy

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

func (p *ConflictingLegacyVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConflictingLegacyVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ConflictingLegacyVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "legacyVmAntiAffinityPolicyExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConflictingLegacyVmAntiAffinityPolicy() *ConflictingLegacyVmAntiAffinityPolicy {
	p := new(ConflictingLegacyVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ConflictingLegacyVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The current VM-host affinity policy is not applied to the VM because the VM is already part of a legacy VM-host affinity policy.
*/
type ConflictingLegacyVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *ConflictingLegacyVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConflictingLegacyVmHostAffinityPolicy

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

func (p *ConflictingLegacyVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConflictingLegacyVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ConflictingLegacyVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConflictingLegacyVmHostAffinityPolicy() *ConflictingLegacyVmHostAffinityPolicy {
	p := new(ConflictingLegacyVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ConflictingLegacyVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The current VM-VM anti-affinity policy is not applied on the VM as another conflicting policy has already been applied to the VM.
*/
type ConflictingVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID of the conflicting VM-VM anti-affinity policy which is applied on the VM.
	*/
	ConflictingVmAntiAffinityPolicyExtId *string `json:"conflictingVmAntiAffinityPolicyExtId,omitempty"`
}

func (p *ConflictingVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConflictingVmAntiAffinityPolicy

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

func (p *ConflictingVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConflictingVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ConflictingVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "conflictingVmAntiAffinityPolicyExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConflictingVmAntiAffinityPolicy() *ConflictingVmAntiAffinityPolicy {
	p := new(ConflictingVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ConflictingVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The current VM-host affinity policy is not applied to the VM because another policy with an earlier creation time is applied to the VM.
*/
type ConflictingVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the conflicting VM-host affinity policy which is applied on the VM.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *ConflictingVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConflictingVmHostAffinityPolicy

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

func (p *ConflictingVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConflictingVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ConflictingVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConflictingVmHostAffinityPolicy() *ConflictingVmHostAffinityPolicy {
	p := new(ConflictingVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ConflictingVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies Post operation
*/
type CreateVmAntiAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVmAntiAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateVmAntiAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateVmAntiAffinityPolicyApiResponse

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

func (p *CreateVmAntiAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateVmAntiAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateVmAntiAffinityPolicyApiResponse(*known)

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

func NewCreateVmAntiAffinityPolicyApiResponse() *CreateVmAntiAffinityPolicyApiResponse {
	p := new(CreateVmAntiAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.CreateVmAntiAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateVmAntiAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateVmAntiAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateVmAntiAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies Post operation
*/
type CreateVmHostAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateVmHostAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateVmHostAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateVmHostAffinityPolicyApiResponse

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

func (p *CreateVmHostAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateVmHostAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateVmHostAffinityPolicyApiResponse(*known)

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

func NewCreateVmHostAffinityPolicyApiResponse() *CreateVmHostAffinityPolicyApiResponse {
	p := new(CreateVmHostAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.CreateVmHostAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateVmHostAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateVmHostAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateVmHostAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/legacy-vm-anti-affinity-policies/{extId} Delete operation
*/
type DeleteLegacyVmAntiAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteLegacyVmAntiAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteLegacyVmAntiAffinityPolicyApiResponse

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

func (p *DeleteLegacyVmAntiAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteLegacyVmAntiAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteLegacyVmAntiAffinityPolicyApiResponse(*known)

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

func NewDeleteLegacyVmAntiAffinityPolicyApiResponse() *DeleteLegacyVmAntiAffinityPolicyApiResponse {
	p := new(DeleteLegacyVmAntiAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.DeleteLegacyVmAntiAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteLegacyVmAntiAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteLegacyVmAntiAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies/{extId} Delete operation
*/
type DeleteVmAntiAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVmAntiAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteVmAntiAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteVmAntiAffinityPolicyApiResponse

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

func (p *DeleteVmAntiAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteVmAntiAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteVmAntiAffinityPolicyApiResponse(*known)

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

func NewDeleteVmAntiAffinityPolicyApiResponse() *DeleteVmAntiAffinityPolicyApiResponse {
	p := new(DeleteVmAntiAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.DeleteVmAntiAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVmAntiAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVmAntiAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVmAntiAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies/{extId} Delete operation
*/
type DeleteVmHostAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteVmHostAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteVmHostAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteVmHostAffinityPolicyApiResponse

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

func (p *DeleteVmHostAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteVmHostAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteVmHostAffinityPolicyApiResponse(*known)

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

func NewDeleteVmHostAffinityPolicyApiResponse() *DeleteVmHostAffinityPolicyApiResponse {
	p := new(DeleteVmHostAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.DeleteVmHostAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteVmHostAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteVmHostAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteVmHostAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies/{extId} Get operation
*/
type GetVmAntiAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmAntiAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmAntiAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmAntiAffinityPolicyApiResponse

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

func (p *GetVmAntiAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmAntiAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVmAntiAffinityPolicyApiResponse(*known)

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

func NewGetVmAntiAffinityPolicyApiResponse() *GetVmAntiAffinityPolicyApiResponse {
	p := new(GetVmAntiAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.GetVmAntiAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmAntiAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmAntiAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmAntiAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies/{extId} Get operation
*/
type GetVmHostAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmHostAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmHostAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmHostAffinityPolicyApiResponse

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

func (p *GetVmHostAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmHostAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVmHostAffinityPolicyApiResponse(*known)

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

func NewGetVmHostAffinityPolicyApiResponse() *GetVmHostAffinityPolicyApiResponse {
	p := new(GetVmHostAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.GetVmHostAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmHostAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmHostAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmHostAffinityPolicyApiResponseData()
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
Reference to a host.
*/
type HostReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the host.
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
	*p.ObjectType_ = "vmm.v4.ahv.policies.HostReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Legacy VM-VM Anti-Affinity policy configuration.
*/
type LegacyVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  VM Anti-Affinity policy name corresponding to a legacy VM-VM Anti-Affinity policy.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VMs that are part of the legacy VM-VM Anti-Affinity policy.
	*/
	Vms []VmReference `json:"vms,omitempty"`
}

func (p *LegacyVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LegacyVmAntiAffinityPolicy

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

func (p *LegacyVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LegacyVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LegacyVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cluster")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")
	delete(allFields, "vms")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLegacyVmAntiAffinityPolicy() *LegacyVmAntiAffinityPolicy {
	p := new(LegacyVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/policies/legacy-vm-anti-affinity-policies Get operation
*/
type ListLegacyVmAntiAffinityPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListLegacyVmAntiAffinityPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListLegacyVmAntiAffinityPoliciesApiResponse

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

func (p *ListLegacyVmAntiAffinityPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListLegacyVmAntiAffinityPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListLegacyVmAntiAffinityPoliciesApiResponse(*known)

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

func NewListLegacyVmAntiAffinityPoliciesApiResponse() *ListLegacyVmAntiAffinityPoliciesApiResponse {
	p := new(ListLegacyVmAntiAffinityPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ListLegacyVmAntiAffinityPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLegacyVmAntiAffinityPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLegacyVmAntiAffinityPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLegacyVmAntiAffinityPoliciesApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies Get operation
*/
type ListVmAntiAffinityPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmAntiAffinityPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmAntiAffinityPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmAntiAffinityPoliciesApiResponse

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

func (p *ListVmAntiAffinityPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmAntiAffinityPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmAntiAffinityPoliciesApiResponse(*known)

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

func NewListVmAntiAffinityPoliciesApiResponse() *ListVmAntiAffinityPoliciesApiResponse {
	p := new(ListVmAntiAffinityPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ListVmAntiAffinityPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmAntiAffinityPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmAntiAffinityPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmAntiAffinityPoliciesApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies/{vmAntiAffinityPolicyExtId}/vm-compliance-states Get operation
*/
type ListVmAntiAffinityPolicyVmComplianceStatesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmAntiAffinityPolicyVmComplianceStatesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmAntiAffinityPolicyVmComplianceStatesApiResponse

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

func (p *ListVmAntiAffinityPolicyVmComplianceStatesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmAntiAffinityPolicyVmComplianceStatesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmAntiAffinityPolicyVmComplianceStatesApiResponse(*known)

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

func NewListVmAntiAffinityPolicyVmComplianceStatesApiResponse() *ListVmAntiAffinityPolicyVmComplianceStatesApiResponse {
	p := new(ListVmAntiAffinityPolicyVmComplianceStatesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ListVmAntiAffinityPolicyVmComplianceStatesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmAntiAffinityPolicyVmComplianceStatesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmAntiAffinityPolicyVmComplianceStatesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies Get operation
*/
type ListVmHostAffinityPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmHostAffinityPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmHostAffinityPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmHostAffinityPoliciesApiResponse

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

func (p *ListVmHostAffinityPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmHostAffinityPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmHostAffinityPoliciesApiResponse(*known)

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

func NewListVmHostAffinityPoliciesApiResponse() *ListVmHostAffinityPoliciesApiResponse {
	p := new(ListVmHostAffinityPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ListVmHostAffinityPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmHostAffinityPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmHostAffinityPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmHostAffinityPoliciesApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies/{vmHostAffinityPolicyExtId}/vm-compliance-states Get operation
*/
type ListVmHostAffinityPolicyVmComplianceStatesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmHostAffinityPolicyVmComplianceStatesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmHostAffinityPolicyVmComplianceStatesApiResponse

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

func (p *ListVmHostAffinityPolicyVmComplianceStatesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmHostAffinityPolicyVmComplianceStatesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmHostAffinityPolicyVmComplianceStatesApiResponse(*known)

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

func NewListVmHostAffinityPolicyVmComplianceStatesApiResponse() *ListVmHostAffinityPolicyVmComplianceStatesApiResponse {
	p := new(ListVmHostAffinityPolicyVmComplianceStatesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ListVmHostAffinityPolicyVmComplianceStatesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmHostAffinityPolicyVmComplianceStatesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmHostAffinityPolicyVmComplianceStatesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData()
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
There are no hosts linked with the VM-host affinity policy.
*/
type NoHostsForVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *NoHostsForVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NoHostsForVmHostAffinityPolicy

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

func (p *NoHostsForVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NoHostsForVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NoHostsForVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNoHostsForVmHostAffinityPolicy() *NoHostsForVmHostAffinityPolicy {
	p := new(NoHostsForVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NoHostsForVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The compliance status of the VM is non-compliant.
*/
type NonCompliantVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	NonComplianceReasonItemDiscriminator_ *string `json:"$nonComplianceReasonItemDiscriminator,omitempty"`
	/*
	  The reason for non-compliance of the VM-VM anti-affinity policy for the VM.
	*/
	NonComplianceReason *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason `json:"nonComplianceReason,omitempty"`
}

func (p *NonCompliantVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NonCompliantVmAntiAffinityPolicy

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

func (p *NonCompliantVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NonCompliantVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NonCompliantVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$nonComplianceReasonItemDiscriminator")
	delete(allFields, "nonComplianceReason")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNonCompliantVmAntiAffinityPolicy() *NonCompliantVmAntiAffinityPolicy {
	p := new(NonCompliantVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NonCompliantVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NonCompliantVmAntiAffinityPolicy) GetNonComplianceReason() interface{} {
	if nil == p.NonComplianceReason {
		return nil
	}
	return p.NonComplianceReason.GetValue()
}

func (p *NonCompliantVmAntiAffinityPolicy) SetNonComplianceReason(v interface{}) error {
	if nil == p.NonComplianceReason {
		p.NonComplianceReason = NewOneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason()
	}
	e := p.NonComplianceReason.SetValue(v)
	if nil == e {
		if nil == p.NonComplianceReasonItemDiscriminator_ {
			p.NonComplianceReasonItemDiscriminator_ = new(string)
		}
		*p.NonComplianceReasonItemDiscriminator_ = *p.NonComplianceReason.Discriminator
	}
	return e
}

/*
VM is non-compliant with the VM-host affinity policy.
*/
type NonCompliantVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	NonComplianceReasonItemDiscriminator_ *string `json:"$nonComplianceReasonItemDiscriminator,omitempty"`
	/*
	  The reason why the VM does not comply with the VM host affinity policy.
	*/
	NonComplianceReason *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason `json:"nonComplianceReason,omitempty"`
}

func (p *NonCompliantVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NonCompliantVmHostAffinityPolicy

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

func (p *NonCompliantVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NonCompliantVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NonCompliantVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$nonComplianceReasonItemDiscriminator")
	delete(allFields, "nonComplianceReason")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNonCompliantVmHostAffinityPolicy() *NonCompliantVmHostAffinityPolicy {
	p := new(NonCompliantVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NonCompliantVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NonCompliantVmHostAffinityPolicy) GetNonComplianceReason() interface{} {
	if nil == p.NonComplianceReason {
		return nil
	}
	return p.NonComplianceReason.GetValue()
}

func (p *NonCompliantVmHostAffinityPolicy) SetNonComplianceReason(v interface{}) error {
	if nil == p.NonComplianceReason {
		p.NonComplianceReason = NewOneOfNonCompliantVmHostAffinityPolicyNonComplianceReason()
	}
	e := p.NonComplianceReason.SetValue(v)
	if nil == e {
		if nil == p.NonComplianceReasonItemDiscriminator_ {
			p.NonComplianceReasonItemDiscriminator_ = new(string)
		}
		*p.NonComplianceReasonItemDiscriminator_ = *p.NonComplianceReason.Discriminator
	}
	return e
}

/*
The VM-VM anti-affinity policy cannot be made compliant for the VM as the number of VMs in the policy are more than the number of hosts on the cluster which is hosting the current VM.
*/
type NotEnoughHostsForVmAntiAffinity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *NotEnoughHostsForVmAntiAffinity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotEnoughHostsForVmAntiAffinity

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

func (p *NotEnoughHostsForVmAntiAffinity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotEnoughHostsForVmAntiAffinity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotEnoughHostsForVmAntiAffinity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotEnoughHostsForVmAntiAffinity() *NotEnoughHostsForVmAntiAffinity {
	p := new(NotEnoughHostsForVmAntiAffinity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NotEnoughHostsForVmAntiAffinity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The VM-VM anti-affinity policy cannot be made compliant for the VM as there are not enough resources on the cluster to enforce the policy. This may be related to pinned VMs, VM-Host affinity policies, etc.
*/
type NotEnoughResourcesForVmAntiAffinity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *NotEnoughResourcesForVmAntiAffinity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotEnoughResourcesForVmAntiAffinity

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

func (p *NotEnoughResourcesForVmAntiAffinity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotEnoughResourcesForVmAntiAffinity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotEnoughResourcesForVmAntiAffinity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotEnoughResourcesForVmAntiAffinity() *NotEnoughResourcesForVmAntiAffinity {
	p := new(NotEnoughResourcesForVmAntiAffinity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NotEnoughResourcesForVmAntiAffinity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The VM-host affinity policy cannot be applied on the VM as there are not enough resources on the hosts in the cluster associated with the policy.
*/
type NotEnoughResourcesForVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *NotEnoughResourcesForVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotEnoughResourcesForVmHostAffinityPolicy

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

func (p *NotEnoughResourcesForVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotEnoughResourcesForVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotEnoughResourcesForVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotEnoughResourcesForVmHostAffinityPolicy() *NotEnoughResourcesForVmHostAffinityPolicy {
	p := new(NotEnoughResourcesForVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.NotEnoughResourcesForVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The VM-VM anti-affinity policy cannot be made compliant for the VM due to some other issue.
*/
type OtherVmAntiAffinityPolicyNonComplianceReason struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *OtherVmAntiAffinityPolicyNonComplianceReason) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OtherVmAntiAffinityPolicyNonComplianceReason

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

func (p *OtherVmAntiAffinityPolicyNonComplianceReason) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OtherVmAntiAffinityPolicyNonComplianceReason
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OtherVmAntiAffinityPolicyNonComplianceReason(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOtherVmAntiAffinityPolicyNonComplianceReason() *OtherVmAntiAffinityPolicyNonComplianceReason {
	p := new(OtherVmAntiAffinityPolicyNonComplianceReason)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.OtherVmAntiAffinityPolicyNonComplianceReason"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM is non-compliant with the VM-host affinity policy due to an internal error.
*/
type OtherVmHostAffinityPolicyVmNonComplianceReason struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *OtherVmHostAffinityPolicyVmNonComplianceReason) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OtherVmHostAffinityPolicyVmNonComplianceReason

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

func (p *OtherVmHostAffinityPolicyVmNonComplianceReason) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OtherVmHostAffinityPolicyVmNonComplianceReason
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OtherVmHostAffinityPolicyVmNonComplianceReason(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOtherVmHostAffinityPolicyVmNonComplianceReason() *OtherVmHostAffinityPolicyVmNonComplianceReason {
	p := new(OtherVmHostAffinityPolicyVmNonComplianceReason)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.OtherVmHostAffinityPolicyVmNonComplianceReason"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
PE is of an older version which doesn't support category based VM-Host Affinity policies. Re-enforce Host Affinity policy after upgrading PE to supported version.
*/
type PeNotCapableForVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Minimum AOS version which supports category based VM-Host Affinity policies.
	*/
	MinimumAosVersionRequired *string `json:"minimumAosVersionRequired,omitempty"`
}

func (p *PeNotCapableForVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PeNotCapableForVmHostAffinityPolicy

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

func (p *PeNotCapableForVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PeNotCapableForVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PeNotCapableForVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "minimumAosVersionRequired")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPeNotCapableForVmHostAffinityPolicy() *PeNotCapableForVmHostAffinityPolicy {
	p := new(PeNotCapableForVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.PeNotCapableForVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The compliance status of the VM is PendingVmAntiAffinityPolicy. Policy enforcement is still in progress.
*/
type PendingVmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *PendingVmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PendingVmAntiAffinityPolicy

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

func (p *PendingVmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PendingVmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PendingVmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPendingVmAntiAffinityPolicy() *PendingVmAntiAffinityPolicy {
	p := new(PendingVmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.PendingVmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The compliance status of VM is pending. The policy enforcement is still in-progress.
*/
type PendingVmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func (p *PendingVmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PendingVmHostAffinityPolicy

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

func (p *PendingVmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PendingVmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PendingVmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPendingVmHostAffinityPolicy() *PendingVmHostAffinityPolicy {
	p := new(PendingVmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.PendingVmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies/{extId}/$actions/re-enforce Post operation
*/
type ReEnforceVmHostAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfReEnforceVmHostAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ReEnforceVmHostAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReEnforceVmHostAffinityPolicyApiResponse

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

func (p *ReEnforceVmHostAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReEnforceVmHostAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ReEnforceVmHostAffinityPolicyApiResponse(*known)

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

func NewReEnforceVmHostAffinityPolicyApiResponse() *ReEnforceVmHostAffinityPolicyApiResponse {
	p := new(ReEnforceVmHostAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.ReEnforceVmHostAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReEnforceVmHostAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ReEnforceVmHostAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfReEnforceVmHostAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-anti-affinity-policies/{extId} Put operation
*/
type UpdateVmAntiAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmAntiAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateVmAntiAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateVmAntiAffinityPolicyApiResponse

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

func (p *UpdateVmAntiAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateVmAntiAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateVmAntiAffinityPolicyApiResponse(*known)

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

func NewUpdateVmAntiAffinityPolicyApiResponse() *UpdateVmAntiAffinityPolicyApiResponse {
	p := new(UpdateVmAntiAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.UpdateVmAntiAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmAntiAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmAntiAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmAntiAffinityPolicyApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/policies/vm-host-affinity-policies/{extId} Put operation
*/
type UpdateVmHostAffinityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateVmHostAffinityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateVmHostAffinityPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateVmHostAffinityPolicyApiResponse

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

func (p *UpdateVmHostAffinityPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateVmHostAffinityPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateVmHostAffinityPolicyApiResponse(*known)

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

func NewUpdateVmHostAffinityPolicyApiResponse() *UpdateVmHostAffinityPolicyApiResponse {
	p := new(UpdateVmHostAffinityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.UpdateVmHostAffinityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateVmHostAffinityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateVmHostAffinityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateVmHostAffinityPolicyApiResponseData()
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
Reference to a user.
*/
type UserReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the user.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *UserReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserReference

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

func (p *UserReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UserReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUserReference() *UserReference {
	p := new(UserReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.UserReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM-VM anti-affinity policy configuration.
*/
type VmAntiAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Categories configured for the VM-VM anti-affinity policy.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`
	/*
	  VM-VM anti-affinity policy creation time.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`
	/*
	  VM-VM anti-affinity policy description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  VM-VM anti-affinity policy name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Number of compliant VMs which are part of the VM-VM anti-affinity policy.
	*/
	NumCompliantVms *int64 `json:"numCompliantVms,omitempty"`
	/*
	  Number of non-compliant VMs which are part of the VM-VM anti-affinity policy.
	*/
	NumNonCompliantVms *int64 `json:"numNonCompliantVms,omitempty"`
	/*
	  Number of VMs with compliance state as pending, which are part of the VM-VM anti-affinity policy.
	*/
	NumPendingVms *int64 `json:"numPendingVms,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VM-VM anti-affinity policy last updated time.
	*/
	UpdateTime *time.Time `json:"updateTime,omitempty"`

	UpdatedBy *UserReference `json:"updatedBy,omitempty"`
}

func (p *VmAntiAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmAntiAffinityPolicy

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

func (p *VmAntiAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmAntiAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmAntiAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "numCompliantVms")
	delete(allFields, "numNonCompliantVms")
	delete(allFields, "numPendingVms")
	delete(allFields, "tenantId")
	delete(allFields, "updateTime")
	delete(allFields, "updatedBy")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmAntiAffinityPolicy() *VmAntiAffinityPolicy {
	p := new(VmAntiAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.VmAntiAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compliance information of a VM which is part of the VM-VM anti-affinity policy.
*/
type VmAntiAffinityPolicyVmComplianceState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of categories through which the VM is associated to the VM-VM anti-affinity policy.
	*/
	AssociatedCategories []CategoryReference `json:"associatedCategories,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/*

	 */
	ComplianceStatusItemDiscriminator_ *string `json:"$complianceStatusItemDiscriminator,omitempty"`
	/*
	  The compliance status of the VM.
	*/
	ComplianceStatus *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus `json:"complianceStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	Host *HostReference `json:"host,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VmAntiAffinityPolicyVmComplianceState) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmAntiAffinityPolicyVmComplianceState

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

func (p *VmAntiAffinityPolicyVmComplianceState) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmAntiAffinityPolicyVmComplianceState
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmAntiAffinityPolicyVmComplianceState(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associatedCategories")
	delete(allFields, "cluster")
	delete(allFields, "$complianceStatusItemDiscriminator")
	delete(allFields, "complianceStatus")
	delete(allFields, "extId")
	delete(allFields, "host")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmAntiAffinityPolicyVmComplianceState() *VmAntiAffinityPolicyVmComplianceState {
	p := new(VmAntiAffinityPolicyVmComplianceState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VmAntiAffinityPolicyVmComplianceState) GetComplianceStatus() interface{} {
	if nil == p.ComplianceStatus {
		return nil
	}
	return p.ComplianceStatus.GetValue()
}

func (p *VmAntiAffinityPolicyVmComplianceState) SetComplianceStatus(v interface{}) error {
	if nil == p.ComplianceStatus {
		p.ComplianceStatus = NewOneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus()
	}
	e := p.ComplianceStatus.SetValue(v)
	if nil == e {
		if nil == p.ComplianceStatusItemDiscriminator_ {
			p.ComplianceStatusItemDiscriminator_ = new(string)
		}
		*p.ComplianceStatusItemDiscriminator_ = *p.ComplianceStatus.Discriminator
	}
	return e
}

/*
VM-Host Affinity policy configuration.
*/
type VmHostAffinityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VM-Host Affinity policy creation time.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`
	/*
	  VM-Host Affinity policy description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Categories through which host is associated with the VM-host affinity policy.
	*/
	HostCategories []CategoryReference `json:"hostCategories,omitempty"`

	LastUpdatedBy *UserReference `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  VM-Host Affinity policy name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Number of VMs which are compliant with the VM-host affinity policy.
	*/
	NumCompliantVms *int64 `json:"numCompliantVms,omitempty"`
	/*
	  Number of hosts associated with the VM-host affinity policy.
	*/
	NumHosts *int64 `json:"numHosts,omitempty"`
	/*
	  Number of VMs which are not compliant with the VM-host affinity policy.
	*/
	NumNonCompliantVms *int64 `json:"numNonCompliantVms,omitempty"`
	/*
	  Number of VMs associated with the VM-host affinity policy.
	*/
	NumVms *int64 `json:"numVms,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VM-Host Affinity policy last updated time.
	*/
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	/*
	  Categories through which VM is associated with the VM-host affinity policy.
	*/
	VmCategories []CategoryReference `json:"vmCategories,omitempty"`
}

func (p *VmHostAffinityPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmHostAffinityPolicy

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

func (p *VmHostAffinityPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmHostAffinityPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmHostAffinityPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "hostCategories")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "numCompliantVms")
	delete(allFields, "numHosts")
	delete(allFields, "numNonCompliantVms")
	delete(allFields, "numVms")
	delete(allFields, "tenantId")
	delete(allFields, "updateTime")
	delete(allFields, "vmCategories")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmHostAffinityPolicy() *VmHostAffinityPolicy {
	p := new(VmHostAffinityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.VmHostAffinityPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compliance information of the VM associated with the VM-host affinity policy.
*/
type VmHostAffinityPolicyVmComplianceState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of categories by which VM is associated to the VM-host affinity policy.
	*/
	AssociatedCategories []CategoryReference `json:"associatedCategories,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/*

	 */
	ComplianceStatusItemDiscriminator_ *string `json:"$complianceStatusItemDiscriminator,omitempty"`
	/*
	  The compliance status of the VM.
	*/
	ComplianceStatus *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus `json:"complianceStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	Host *HostReference `json:"host,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VmHostAffinityPolicyVmComplianceState) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmHostAffinityPolicyVmComplianceState

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

func (p *VmHostAffinityPolicyVmComplianceState) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmHostAffinityPolicyVmComplianceState
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmHostAffinityPolicyVmComplianceState(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associatedCategories")
	delete(allFields, "cluster")
	delete(allFields, "$complianceStatusItemDiscriminator")
	delete(allFields, "complianceStatus")
	delete(allFields, "extId")
	delete(allFields, "host")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmHostAffinityPolicyVmComplianceState() *VmHostAffinityPolicyVmComplianceState {
	p := new(VmHostAffinityPolicyVmComplianceState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VmHostAffinityPolicyVmComplianceState) GetComplianceStatus() interface{} {
	if nil == p.ComplianceStatus {
		return nil
	}
	return p.ComplianceStatus.GetValue()
}

func (p *VmHostAffinityPolicyVmComplianceState) SetComplianceStatus(v interface{}) error {
	if nil == p.ComplianceStatus {
		p.ComplianceStatus = NewOneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus()
	}
	e := p.ComplianceStatus.SetValue(v)
	if nil == e {
		if nil == p.ComplianceStatusItemDiscriminator_ {
			p.ComplianceStatusItemDiscriminator_ = new(string)
		}
		*p.ComplianceStatusItemDiscriminator_ = *p.ComplianceStatus.Discriminator
	}
	return e
}

/*
Reference to a VM.
*/
type VmReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external ID (UUID) of the VM.
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
	*p.ObjectType_ = "vmm.v4.ahv.policies.VmReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason struct {
	Discriminator *string                                         `json:"-"`
	ObjectType_   *string                                         `json:"-"`
	oneOfType2003 *NoHostsForVmHostAffinityPolicy                 `json:"-"`
	oneOfType2001 *ConflictingVmHostAffinityPolicy                `json:"-"`
	oneOfType2002 *ConflictingLegacyVmHostAffinityPolicy          `json:"-"`
	oneOfType2005 *PeNotCapableForVmHostAffinityPolicy            `json:"-"`
	oneOfType2004 *NotEnoughResourcesForVmHostAffinityPolicy      `json:"-"`
	oneOfType2006 *OtherVmHostAffinityPolicyVmNonComplianceReason `json:"-"`
}

func NewOneOfNonCompliantVmHostAffinityPolicyNonComplianceReason() *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason {
	p := new(OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason is nil"))
	}
	switch v.(type) {
	case NoHostsForVmHostAffinityPolicy:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(NoHostsForVmHostAffinityPolicy)
		}
		*p.oneOfType2003 = v.(NoHostsForVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	case ConflictingVmHostAffinityPolicy:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ConflictingVmHostAffinityPolicy)
		}
		*p.oneOfType2001 = v.(ConflictingVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case ConflictingLegacyVmHostAffinityPolicy:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(ConflictingLegacyVmHostAffinityPolicy)
		}
		*p.oneOfType2002 = v.(ConflictingLegacyVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case PeNotCapableForVmHostAffinityPolicy:
		if nil == p.oneOfType2005 {
			p.oneOfType2005 = new(PeNotCapableForVmHostAffinityPolicy)
		}
		*p.oneOfType2005 = v.(PeNotCapableForVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2005.ObjectType_
	case NotEnoughResourcesForVmHostAffinityPolicy:
		if nil == p.oneOfType2004 {
			p.oneOfType2004 = new(NotEnoughResourcesForVmHostAffinityPolicy)
		}
		*p.oneOfType2004 = v.(NotEnoughResourcesForVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2004.ObjectType_
	case OtherVmHostAffinityPolicyVmNonComplianceReason:
		if nil == p.oneOfType2006 {
			p.oneOfType2006 = new(OtherVmHostAffinityPolicyVmNonComplianceReason)
		}
		*p.oneOfType2006 = v.(OtherVmHostAffinityPolicyVmNonComplianceReason)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2006.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason) GetValue() interface{} {
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2005
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2004
	}
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2006
	}
	return nil
}

func (p *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason) UnmarshalJSON(b []byte) error {
	vOneOfType2003 := new(NoHostsForVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "vmm.v4.ahv.policies.NoHostsForVmHostAffinityPolicy" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(NoHostsForVmHostAffinityPolicy)
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
	vOneOfType2001 := new(ConflictingVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.policies.ConflictingVmHostAffinityPolicy" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ConflictingVmHostAffinityPolicy)
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
	vOneOfType2002 := new(ConflictingLegacyVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.policies.ConflictingLegacyVmHostAffinityPolicy" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(ConflictingLegacyVmHostAffinityPolicy)
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
	vOneOfType2005 := new(PeNotCapableForVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2005); err == nil {
		if "vmm.v4.ahv.policies.PeNotCapableForVmHostAffinityPolicy" == *vOneOfType2005.ObjectType_ {
			if nil == p.oneOfType2005 {
				p.oneOfType2005 = new(PeNotCapableForVmHostAffinityPolicy)
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
	vOneOfType2004 := new(NotEnoughResourcesForVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if "vmm.v4.ahv.policies.NotEnoughResourcesForVmHostAffinityPolicy" == *vOneOfType2004.ObjectType_ {
			if nil == p.oneOfType2004 {
				p.oneOfType2004 = new(NotEnoughResourcesForVmHostAffinityPolicy)
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
	vOneOfType2006 := new(OtherVmHostAffinityPolicyVmNonComplianceReason)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if "vmm.v4.ahv.policies.OtherVmHostAffinityPolicyVmNonComplianceReason" == *vOneOfType2006.ObjectType_ {
			if nil == p.oneOfType2006 {
				p.oneOfType2006 = new(OtherVmHostAffinityPolicyVmNonComplianceReason)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason"))
}

func (p *OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason) MarshalJSON() ([]byte, error) {
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2005)
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	return nil, errors.New("No value to marshal for OneOfNonCompliantVmHostAffinityPolicyNonComplianceReason")
}

type OneOfGetVmHostAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VmHostAffinityPolicy  `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetVmHostAffinityPolicyApiResponseData() *OneOfGetVmHostAffinityPolicyApiResponseData {
	p := new(OneOfGetVmHostAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmHostAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmHostAffinityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case VmHostAffinityPolicy:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmHostAffinityPolicy)
		}
		*p.oneOfType2001 = v.(VmHostAffinityPolicy)
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

func (p *OneOfGetVmHostAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmHostAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.policies.VmHostAffinityPolicy" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmHostAffinityPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmHostAffinityPolicyApiResponseData"))
}

func (p *OneOfGetVmHostAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmHostAffinityPolicyApiResponseData")
}

type OneOfListVmHostAffinityPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []VmHostAffinityPolicy `json:"-"`
}

func NewOneOfListVmHostAffinityPoliciesApiResponseData() *OneOfListVmHostAffinityPoliciesApiResponseData {
	p := new(OneOfListVmHostAffinityPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmHostAffinityPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmHostAffinityPoliciesApiResponseData is nil"))
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
	case []VmHostAffinityPolicy:
		p.oneOfType2001 = v.([]VmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVmHostAffinityPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListVmHostAffinityPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType2001 := new([]VmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.policies.VmHostAffinityPolicy" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmHostAffinityPoliciesApiResponseData"))
}

func (p *OneOfListVmHostAffinityPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.policies.VmHostAffinityPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListVmHostAffinityPoliciesApiResponseData")
}

type OneOfDeleteVmHostAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVmHostAffinityPolicyApiResponseData() *OneOfDeleteVmHostAffinityPolicyApiResponseData {
	p := new(OneOfDeleteVmHostAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVmHostAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVmHostAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteVmHostAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVmHostAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVmHostAffinityPolicyApiResponseData"))
}

func (p *OneOfDeleteVmHostAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVmHostAffinityPolicyApiResponseData")
}

type OneOfUpdateVmHostAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmHostAffinityPolicyApiResponseData() *OneOfUpdateVmHostAffinityPolicyApiResponseData {
	p := new(OneOfUpdateVmHostAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmHostAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmHostAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfUpdateVmHostAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmHostAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmHostAffinityPolicyApiResponseData"))
}

func (p *OneOfUpdateVmHostAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmHostAffinityPolicyApiResponseData")
}

type OneOfCreateVmAntiAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateVmAntiAffinityPolicyApiResponseData() *OneOfCreateVmAntiAffinityPolicyApiResponseData {
	p := new(OneOfCreateVmAntiAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVmAntiAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVmAntiAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfCreateVmAntiAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVmAntiAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVmAntiAffinityPolicyApiResponseData"))
}

func (p *OneOfCreateVmAntiAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVmAntiAffinityPolicyApiResponseData")
}

type OneOfCreateVmHostAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateVmHostAffinityPolicyApiResponseData() *OneOfCreateVmHostAffinityPolicyApiResponseData {
	p := new(OneOfCreateVmHostAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateVmHostAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateVmHostAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfCreateVmHostAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateVmHostAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateVmHostAffinityPolicyApiResponseData"))
}

func (p *OneOfCreateVmHostAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateVmHostAffinityPolicyApiResponseData")
}

type OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData struct {
	Discriminator *string                                 `json:"-"`
	ObjectType_   *string                                 `json:"-"`
	oneOfType400  *import2.ErrorResponse                  `json:"-"`
	oneOfType2001 []VmHostAffinityPolicyVmComplianceState `json:"-"`
}

func NewOneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData() *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData {
	p := new(OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData is nil"))
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
	case []VmHostAffinityPolicyVmComplianceState:
		p.oneOfType2001 = v.([]VmHostAffinityPolicyVmComplianceState)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType2001 := new([]VmHostAffinityPolicyVmComplianceState)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData"))
}

func (p *OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.policies.VmHostAffinityPolicyVmComplianceState>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListVmHostAffinityPolicyVmComplianceStatesApiResponseData")
}

type OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData() *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData {
	p := new(OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData"))
}

func (p *OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteLegacyVmAntiAffinityPolicyApiResponseData")
}

type OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason struct {
	Discriminator *string                                       `json:"-"`
	ObjectType_   *string                                       `json:"-"`
	oneOfType2002 *ConflictingVmAntiAffinityPolicy              `json:"-"`
	oneOfType2006 *OtherVmAntiAffinityPolicyNonComplianceReason `json:"-"`
	oneOfType2007 *ClusterNotSupportedForVmAntiAffinity         `json:"-"`
	oneOfType2003 *ConflictingLegacyVmAntiAffinityPolicy        `json:"-"`
	oneOfType2005 *NotEnoughResourcesForVmAntiAffinity          `json:"-"`
	oneOfType2004 *NotEnoughHostsForVmAntiAffinity              `json:"-"`
}

func NewOneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason() *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason {
	p := new(OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason is nil"))
	}
	switch v.(type) {
	case ConflictingVmAntiAffinityPolicy:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(ConflictingVmAntiAffinityPolicy)
		}
		*p.oneOfType2002 = v.(ConflictingVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case OtherVmAntiAffinityPolicyNonComplianceReason:
		if nil == p.oneOfType2006 {
			p.oneOfType2006 = new(OtherVmAntiAffinityPolicyNonComplianceReason)
		}
		*p.oneOfType2006 = v.(OtherVmAntiAffinityPolicyNonComplianceReason)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2006.ObjectType_
	case ClusterNotSupportedForVmAntiAffinity:
		if nil == p.oneOfType2007 {
			p.oneOfType2007 = new(ClusterNotSupportedForVmAntiAffinity)
		}
		*p.oneOfType2007 = v.(ClusterNotSupportedForVmAntiAffinity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2007.ObjectType_
	case ConflictingLegacyVmAntiAffinityPolicy:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(ConflictingLegacyVmAntiAffinityPolicy)
		}
		*p.oneOfType2003 = v.(ConflictingLegacyVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	case NotEnoughResourcesForVmAntiAffinity:
		if nil == p.oneOfType2005 {
			p.oneOfType2005 = new(NotEnoughResourcesForVmAntiAffinity)
		}
		*p.oneOfType2005 = v.(NotEnoughResourcesForVmAntiAffinity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2005.ObjectType_
	case NotEnoughHostsForVmAntiAffinity:
		if nil == p.oneOfType2004 {
			p.oneOfType2004 = new(NotEnoughHostsForVmAntiAffinity)
		}
		*p.oneOfType2004 = v.(NotEnoughHostsForVmAntiAffinity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2004.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2006
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2007
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2005
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2004
	}
	return nil
}

func (p *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(ConflictingVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.policies.ConflictingVmAntiAffinityPolicy" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(ConflictingVmAntiAffinityPolicy)
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
	vOneOfType2006 := new(OtherVmAntiAffinityPolicyNonComplianceReason)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if "vmm.v4.ahv.policies.OtherVmAntiAffinityPolicyNonComplianceReason" == *vOneOfType2006.ObjectType_ {
			if nil == p.oneOfType2006 {
				p.oneOfType2006 = new(OtherVmAntiAffinityPolicyNonComplianceReason)
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
	vOneOfType2007 := new(ClusterNotSupportedForVmAntiAffinity)
	if err := json.Unmarshal(b, vOneOfType2007); err == nil {
		if "vmm.v4.ahv.policies.ClusterNotSupportedForVmAntiAffinity" == *vOneOfType2007.ObjectType_ {
			if nil == p.oneOfType2007 {
				p.oneOfType2007 = new(ClusterNotSupportedForVmAntiAffinity)
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
	vOneOfType2003 := new(ConflictingLegacyVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "vmm.v4.ahv.policies.ConflictingLegacyVmAntiAffinityPolicy" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(ConflictingLegacyVmAntiAffinityPolicy)
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
	vOneOfType2005 := new(NotEnoughResourcesForVmAntiAffinity)
	if err := json.Unmarshal(b, vOneOfType2005); err == nil {
		if "vmm.v4.ahv.policies.NotEnoughResourcesForVmAntiAffinity" == *vOneOfType2005.ObjectType_ {
			if nil == p.oneOfType2005 {
				p.oneOfType2005 = new(NotEnoughResourcesForVmAntiAffinity)
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
	vOneOfType2004 := new(NotEnoughHostsForVmAntiAffinity)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if "vmm.v4.ahv.policies.NotEnoughHostsForVmAntiAffinity" == *vOneOfType2004.ObjectType_ {
			if nil == p.oneOfType2004 {
				p.oneOfType2004 = new(NotEnoughHostsForVmAntiAffinity)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason"))
}

func (p *OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2007)
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2005)
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	return nil, errors.New("No value to marshal for OneOfNonCompliantVmAntiAffinityPolicyNonComplianceReason")
}

type OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData struct {
	Discriminator *string                                 `json:"-"`
	ObjectType_   *string                                 `json:"-"`
	oneOfType2001 []VmAntiAffinityPolicyVmComplianceState `json:"-"`
	oneOfType400  *import2.ErrorResponse                  `json:"-"`
}

func NewOneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData() *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData {
	p := new(OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData is nil"))
	}
	switch v.(type) {
	case []VmAntiAffinityPolicyVmComplianceState:
		p.oneOfType2001 = v.([]VmAntiAffinityPolicyVmComplianceState)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>"
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

func (p *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]VmAntiAffinityPolicyVmComplianceState)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>"
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData"))
}

func (p *OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.policies.VmAntiAffinityPolicyVmComplianceState>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmAntiAffinityPolicyVmComplianceStatesApiResponseData")
}

type OneOfDeleteVmAntiAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteVmAntiAffinityPolicyApiResponseData() *OneOfDeleteVmAntiAffinityPolicyApiResponseData {
	p := new(OneOfDeleteVmAntiAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteVmAntiAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteVmAntiAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteVmAntiAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteVmAntiAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteVmAntiAffinityPolicyApiResponseData"))
}

func (p *OneOfDeleteVmAntiAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteVmAntiAffinityPolicyApiResponseData")
}

type OneOfListLegacyVmAntiAffinityPoliciesApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType400  *import2.ErrorResponse       `json:"-"`
	oneOfType2001 []LegacyVmAntiAffinityPolicy `json:"-"`
}

func NewOneOfListLegacyVmAntiAffinityPoliciesApiResponseData() *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData {
	p := new(OneOfListLegacyVmAntiAffinityPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLegacyVmAntiAffinityPoliciesApiResponseData is nil"))
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
	case []LegacyVmAntiAffinityPolicy:
		p.oneOfType2001 = v.([]LegacyVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType2001 := new([]LegacyVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLegacyVmAntiAffinityPoliciesApiResponseData"))
}

func (p *OneOfListLegacyVmAntiAffinityPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.ahv.policies.LegacyVmAntiAffinityPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListLegacyVmAntiAffinityPoliciesApiResponseData")
}

type OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus struct {
	Discriminator *string                           `json:"-"`
	ObjectType_   *string                           `json:"-"`
	oneOfType2006 *CompliantVmAntiAffinityPolicy    `json:"-"`
	oneOfType2007 *NonCompliantVmAntiAffinityPolicy `json:"-"`
	oneOfType2008 *PendingVmAntiAffinityPolicy      `json:"-"`
}

func NewOneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus() *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus {
	p := new(OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus is nil"))
	}
	switch v.(type) {
	case CompliantVmAntiAffinityPolicy:
		if nil == p.oneOfType2006 {
			p.oneOfType2006 = new(CompliantVmAntiAffinityPolicy)
		}
		*p.oneOfType2006 = v.(CompliantVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2006.ObjectType_
	case NonCompliantVmAntiAffinityPolicy:
		if nil == p.oneOfType2007 {
			p.oneOfType2007 = new(NonCompliantVmAntiAffinityPolicy)
		}
		*p.oneOfType2007 = v.(NonCompliantVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2007.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2007.ObjectType_
	case PendingVmAntiAffinityPolicy:
		if nil == p.oneOfType2008 {
			p.oneOfType2008 = new(PendingVmAntiAffinityPolicy)
		}
		*p.oneOfType2008 = v.(PendingVmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2008.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2008.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus) GetValue() interface{} {
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2006
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2007
	}
	if p.oneOfType2008 != nil && *p.oneOfType2008.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2008
	}
	return nil
}

func (p *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus) UnmarshalJSON(b []byte) error {
	vOneOfType2006 := new(CompliantVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2006); err == nil {
		if "vmm.v4.ahv.policies.CompliantVmAntiAffinityPolicy" == *vOneOfType2006.ObjectType_ {
			if nil == p.oneOfType2006 {
				p.oneOfType2006 = new(CompliantVmAntiAffinityPolicy)
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
	vOneOfType2007 := new(NonCompliantVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2007); err == nil {
		if "vmm.v4.ahv.policies.NonCompliantVmAntiAffinityPolicy" == *vOneOfType2007.ObjectType_ {
			if nil == p.oneOfType2007 {
				p.oneOfType2007 = new(NonCompliantVmAntiAffinityPolicy)
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
	vOneOfType2008 := new(PendingVmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2008); err == nil {
		if "vmm.v4.ahv.policies.PendingVmAntiAffinityPolicy" == *vOneOfType2008.ObjectType_ {
			if nil == p.oneOfType2008 {
				p.oneOfType2008 = new(PendingVmAntiAffinityPolicy)
			}
			*p.oneOfType2008 = *vOneOfType2008
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2008.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2008.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus"))
}

func (p *OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus) MarshalJSON() ([]byte, error) {
	if p.oneOfType2006 != nil && *p.oneOfType2006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2006)
	}
	if p.oneOfType2007 != nil && *p.oneOfType2007.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2007)
	}
	if p.oneOfType2008 != nil && *p.oneOfType2008.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2008)
	}
	return nil, errors.New("No value to marshal for OneOfVmAntiAffinityPolicyVmComplianceStateComplianceStatus")
}

type OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus struct {
	Discriminator *string                           `json:"-"`
	ObjectType_   *string                           `json:"-"`
	oneOfType2002 *NonCompliantVmHostAffinityPolicy `json:"-"`
	oneOfType2003 *PendingVmHostAffinityPolicy      `json:"-"`
	oneOfType2001 *CompliantVmHostAffinityPolicy    `json:"-"`
}

func NewOneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus() *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus {
	p := new(OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus is nil"))
	}
	switch v.(type) {
	case NonCompliantVmHostAffinityPolicy:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(NonCompliantVmHostAffinityPolicy)
		}
		*p.oneOfType2002 = v.(NonCompliantVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case PendingVmHostAffinityPolicy:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(PendingVmHostAffinityPolicy)
		}
		*p.oneOfType2003 = v.(PendingVmHostAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	case CompliantVmHostAffinityPolicy:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(CompliantVmHostAffinityPolicy)
		}
		*p.oneOfType2001 = v.(CompliantVmHostAffinityPolicy)
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

func (p *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(NonCompliantVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.policies.NonCompliantVmHostAffinityPolicy" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(NonCompliantVmHostAffinityPolicy)
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
	vOneOfType2003 := new(PendingVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "vmm.v4.ahv.policies.PendingVmHostAffinityPolicy" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(PendingVmHostAffinityPolicy)
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
	vOneOfType2001 := new(CompliantVmHostAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.policies.CompliantVmHostAffinityPolicy" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(CompliantVmHostAffinityPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus"))
}

func (p *OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfVmHostAffinityPolicyVmComplianceStateComplianceStatus")
}

type OneOfGetVmAntiAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VmAntiAffinityPolicy  `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetVmAntiAffinityPolicyApiResponseData() *OneOfGetVmAntiAffinityPolicyApiResponseData {
	p := new(OneOfGetVmAntiAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmAntiAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmAntiAffinityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case VmAntiAffinityPolicy:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmAntiAffinityPolicy)
		}
		*p.oneOfType2001 = v.(VmAntiAffinityPolicy)
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

func (p *OneOfGetVmAntiAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmAntiAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.policies.VmAntiAffinityPolicy" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmAntiAffinityPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmAntiAffinityPolicyApiResponseData"))
}

func (p *OneOfGetVmAntiAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmAntiAffinityPolicyApiResponseData")
}

type OneOfUpdateVmAntiAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateVmAntiAffinityPolicyApiResponseData() *OneOfUpdateVmAntiAffinityPolicyApiResponseData {
	p := new(OneOfUpdateVmAntiAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateVmAntiAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateVmAntiAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfUpdateVmAntiAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateVmAntiAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateVmAntiAffinityPolicyApiResponseData"))
}

func (p *OneOfUpdateVmAntiAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateVmAntiAffinityPolicyApiResponseData")
}

type OneOfListVmAntiAffinityPoliciesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []VmAntiAffinityPolicy `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListVmAntiAffinityPoliciesApiResponseData() *OneOfListVmAntiAffinityPoliciesApiResponseData {
	p := new(OneOfListVmAntiAffinityPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmAntiAffinityPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmAntiAffinityPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []VmAntiAffinityPolicy:
		p.oneOfType2001 = v.([]VmAntiAffinityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>"
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

func (p *OneOfListVmAntiAffinityPoliciesApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmAntiAffinityPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]VmAntiAffinityPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.policies.VmAntiAffinityPolicy" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>"
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmAntiAffinityPoliciesApiResponseData"))
}

func (p *OneOfListVmAntiAffinityPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.policies.VmAntiAffinityPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmAntiAffinityPoliciesApiResponseData")
}

type OneOfReEnforceVmHostAffinityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfReEnforceVmHostAffinityPolicyApiResponseData() *OneOfReEnforceVmHostAffinityPolicyApiResponseData {
	p := new(OneOfReEnforceVmHostAffinityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReEnforceVmHostAffinityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReEnforceVmHostAffinityPolicyApiResponseData is nil"))
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

func (p *OneOfReEnforceVmHostAffinityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfReEnforceVmHostAffinityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReEnforceVmHostAffinityPolicyApiResponseData"))
}

func (p *OneOfReEnforceVmHostAffinityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfReEnforceVmHostAffinityPolicyApiResponseData")
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
