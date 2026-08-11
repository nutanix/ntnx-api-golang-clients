/*
 * Generated file models/datapolicies/v4/operations/operations_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Data Policies APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Disaster recovery and storage policies like overriding protection domain.
*/
package operations

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/dataprotection/v4/common"
)

/*
Category override spec.
*/
type CategoryOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the category.
	*/
	CategoryName *string `json:"categoryName"`
	/*
	  Value of the category.
	*/
	CategoryValue *string `json:"categoryValue,omitempty"`
}

func (p *CategoryOverrideSpec) MarshalJSON() ([]byte, error) {
	type CategoryOverrideSpecProxy CategoryOverrideSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CategoryOverrideSpecProxy
		CategoryName *string `json:"categoryName,omitempty"`
	}{
		CategoryOverrideSpecProxy: (*CategoryOverrideSpecProxy)(p),
		CategoryName:              p.CategoryName,
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

func (p *CategoryOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CategoryOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCategoryOverrideSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryName != nil {
		p.CategoryName = known.CategoryName
	}
	if known.CategoryValue != nil {
		p.CategoryValue = known.CategoryValue
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryName")
	delete(allFields, "categoryValue")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCategoryOverrideSpec() *CategoryOverrideSpec {
	p := new(CategoryOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.operations.CategoryOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for converting protection domains to entity centric data protection.
*/
type ConvertProtectionDomainsSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of protection domains to be converted and migrated to entity centric data protection.
	*/
	ProtectionDomainSpecs []ProtectionDomainSpec `json:"protectionDomainSpecs"`
}

func (p *ConvertProtectionDomainsSpec) MarshalJSON() ([]byte, error) {
	type ConvertProtectionDomainsSpecProxy ConvertProtectionDomainsSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConvertProtectionDomainsSpecProxy
		ProtectionDomainSpecs []ProtectionDomainSpec `json:"protectionDomainSpecs,omitempty"`
	}{
		ConvertProtectionDomainsSpecProxy: (*ConvertProtectionDomainsSpecProxy)(p),
		ProtectionDomainSpecs:             p.ProtectionDomainSpecs,
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

func (p *ConvertProtectionDomainsSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConvertProtectionDomainsSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConvertProtectionDomainsSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProtectionDomainSpecs != nil {
		p.ProtectionDomainSpecs = known.ProtectionDomainSpecs
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "protectionDomainSpecs")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConvertProtectionDomainsSpec() *ConvertProtectionDomainsSpec {
	p := new(ConvertProtectionDomainsSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.operations.ConvertProtectionDomainsSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Protection domain override spec.
*/
type ProtectionDomainOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CategorySpec *CategoryOverrideSpec `json:"categorySpec,omitempty"`

	RecoveryPointType *import1.RecoveryPointType `json:"recoveryPointType,omitempty"`
	/*
	  Flag to indicate whether a recovery plan should be created for this protection domain during the conversion workflow.
	*/
	ShouldCreateRecoveryPlan *bool `json:"shouldCreateRecoveryPlan,omitempty"`
	/*
	  Flag to indicate whether the recovery point objective should be set to weekly for this protection domain during the conversion workflow. This is only applicable if the protection domain has custom schedules requiring snapshots on particular days of the week.
	*/
	ShouldDisableDaySpecificSnapshots *bool `json:"shouldDisableDaySpecificSnapshots,omitempty"`
}

func (p *ProtectionDomainOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProtectionDomainOverrideSpec

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

func (p *ProtectionDomainOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectionDomainOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectionDomainOverrideSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategorySpec != nil {
		p.CategorySpec = known.CategorySpec
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.ShouldCreateRecoveryPlan != nil {
		p.ShouldCreateRecoveryPlan = known.ShouldCreateRecoveryPlan
	}
	if known.ShouldDisableDaySpecificSnapshots != nil {
		p.ShouldDisableDaySpecificSnapshots = known.ShouldDisableDaySpecificSnapshots
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categorySpec")
	delete(allFields, "recoveryPointType")
	delete(allFields, "shouldCreateRecoveryPlan")
	delete(allFields, "shouldDisableDaySpecificSnapshots")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectionDomainOverrideSpec() *ProtectionDomainOverrideSpec {
	p := new(ProtectionDomainOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.operations.ProtectionDomainOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldCreateRecoveryPlan = new(bool)
	*p.ShouldCreateRecoveryPlan = true
	p.ShouldDisableDaySpecificSnapshots = new(bool)
	*p.ShouldDisableDaySpecificSnapshots = false

	return p
}

/*
Protection domain spec.
*/
type ProtectionDomainSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the protection domain.
	*/
	ProtectionDomainName *string `json:"protectionDomainName"`

	ProtectionDomainOverrideSpec *ProtectionDomainOverrideSpec `json:"protectionDomainOverrideSpec,omitempty"`
}

func (p *ProtectionDomainSpec) MarshalJSON() ([]byte, error) {
	type ProtectionDomainSpecProxy ProtectionDomainSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProtectionDomainSpecProxy
		ProtectionDomainName *string `json:"protectionDomainName,omitempty"`
	}{
		ProtectionDomainSpecProxy: (*ProtectionDomainSpecProxy)(p),
		ProtectionDomainName:      p.ProtectionDomainName,
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

func (p *ProtectionDomainSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectionDomainSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectionDomainSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProtectionDomainName != nil {
		p.ProtectionDomainName = known.ProtectionDomainName
	}
	if known.ProtectionDomainOverrideSpec != nil {
		p.ProtectionDomainOverrideSpec = known.ProtectionDomainOverrideSpec
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "protectionDomainName")
	delete(allFields, "protectionDomainOverrideSpec")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectionDomainSpec() *ProtectionDomainSpec {
	p := new(ProtectionDomainSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.operations.ProtectionDomainSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
