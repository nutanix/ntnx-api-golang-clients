/*
 * Generated file models/dataprotection/v4/operations/operations_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Data Protection APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.operations of Nutanix Data Protection APIs
*/
package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/error"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/prism/v4/config"
	"time"
)

/*
Recovery plan action is a specification describing the recovery plan to be executed, failover directions and type of failover.
*/
type BaseRecoveryPlanActionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []import1.FailoverDirection `json:"failoverDirections"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name"`
}

func (p *BaseRecoveryPlanActionSpec) MarshalJSON() ([]byte, error) {
	type BaseRecoveryPlanActionSpecProxy BaseRecoveryPlanActionSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BaseRecoveryPlanActionSpecProxy
		FailoverDirections []import1.FailoverDirection `json:"failoverDirections,omitempty"`
		Name               *string                     `json:"name,omitempty"`
	}{
		BaseRecoveryPlanActionSpecProxy: (*BaseRecoveryPlanActionSpecProxy)(p),
		FailoverDirections:              p.FailoverDirections,
		Name:                            p.Name,
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

func (p *BaseRecoveryPlanActionSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseRecoveryPlanActionSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseRecoveryPlanActionSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverDirections")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseRecoveryPlanActionSpec() *BaseRecoveryPlanActionSpec {
	p := new(BaseRecoveryPlanActionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.BaseRecoveryPlanActionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/clean-up-resources Post operation
*/
type CleanupRecoveryPlanResourcesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCleanupRecoveryPlanResourcesApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CleanupRecoveryPlanResourcesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CleanupRecoveryPlanResourcesApiResponse

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

func (p *CleanupRecoveryPlanResourcesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CleanupRecoveryPlanResourcesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCleanupRecoveryPlanResourcesApiResponse()

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

func NewCleanupRecoveryPlanResourcesApiResponse() *CleanupRecoveryPlanResourcesApiResponse {
	p := new(CleanupRecoveryPlanResourcesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.CleanupRecoveryPlanResourcesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CleanupRecoveryPlanResourcesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CleanupRecoveryPlanResourcesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCleanupRecoveryPlanResourcesApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/planned-failover Post operation
*/
type PlannedFailoverRecoveryPlanApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPlannedFailoverRecoveryPlanApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *PlannedFailoverRecoveryPlanApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PlannedFailoverRecoveryPlanApiResponse

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

func (p *PlannedFailoverRecoveryPlanApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PlannedFailoverRecoveryPlanApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPlannedFailoverRecoveryPlanApiResponse()

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

func NewPlannedFailoverRecoveryPlanApiResponse() *PlannedFailoverRecoveryPlanApiResponse {
	p := new(PlannedFailoverRecoveryPlanApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.PlannedFailoverRecoveryPlanApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PlannedFailoverRecoveryPlanApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PlannedFailoverRecoveryPlanApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPlannedFailoverRecoveryPlanApiResponseData()
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
The planned failover recovery plan action is used to perform a planned failover on the recovery plan.
*/
type PlannedFailoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []import1.FailoverDirection `json:"failoverDirections"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name"`
	/*
	  Indicates whether to continue the recovery plan action despite validation warnings. For example, if the IP of some virtual machines (VMs) cannot be preserved after recovery, the Recovery Plan action can still be initiated by passing `shouldIgnoreWarnings` as `true`. This value allows the recovery operations to continue despite the validation warnings. Alternatively, the user can choose to resolve the validation warnings and trigger another failover action on recovery plan.
	*/
	ShouldIgnoreWarnings *bool `json:"shouldIgnoreWarnings,omitempty"`
	/*
	  Indicates whether to perform a live migration of virtual machines (VMs) during a `PLANNED_FAILOVER` operation. When not specified or specified as false, the migration is performed for all virtual machines (VMs) in offline mode. When specified as true, the migration is performed for all virtual machines (VMs) in a running state.
	*/
	ShouldLiveMigrateVMs *bool `json:"shouldLiveMigrateVMs,omitempty"`
}

func (p *PlannedFailoverSpec) MarshalJSON() ([]byte, error) {
	type PlannedFailoverSpecProxy PlannedFailoverSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PlannedFailoverSpecProxy
		FailoverDirections []import1.FailoverDirection `json:"failoverDirections,omitempty"`
		Name               *string                     `json:"name,omitempty"`
	}{
		PlannedFailoverSpecProxy: (*PlannedFailoverSpecProxy)(p),
		FailoverDirections:       p.FailoverDirections,
		Name:                     p.Name,
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

func (p *PlannedFailoverSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PlannedFailoverSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPlannedFailoverSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ShouldIgnoreWarnings != nil {
		p.ShouldIgnoreWarnings = known.ShouldIgnoreWarnings
	}
	if known.ShouldLiveMigrateVMs != nil {
		p.ShouldLiveMigrateVMs = known.ShouldLiveMigrateVMs
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverDirections")
	delete(allFields, "name")
	delete(allFields, "shouldIgnoreWarnings")
	delete(allFields, "shouldLiveMigrateVMs")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPlannedFailoverSpec() *PlannedFailoverSpec {
	p := new(PlannedFailoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.PlannedFailoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldIgnoreWarnings = new(bool)
	*p.ShouldIgnoreWarnings = false
	p.ShouldLiveMigrateVMs = new(bool)
	*p.ShouldLiveMigrateVMs = false

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/test-failover Post operation
*/
type TestFailoverRecoveryPlanApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTestFailoverRecoveryPlanApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *TestFailoverRecoveryPlanApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TestFailoverRecoveryPlanApiResponse

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

func (p *TestFailoverRecoveryPlanApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TestFailoverRecoveryPlanApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTestFailoverRecoveryPlanApiResponse()

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

func NewTestFailoverRecoveryPlanApiResponse() *TestFailoverRecoveryPlanApiResponse {
	p := new(TestFailoverRecoveryPlanApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.TestFailoverRecoveryPlanApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TestFailoverRecoveryPlanApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TestFailoverRecoveryPlanApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTestFailoverRecoveryPlanApiResponseData()
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
The test failover recovery plan action is used to perform a test failover on the recovery plan.
*/
type TestFailoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []import1.FailoverDirection `json:"failoverDirections"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name"`
	/*
	  Indicates whether to continue the recovery plan action despite validation warnings. For example, if the IP of some virtual machines (VMs) cannot be preserved after recovery, the Recovery Plan action can still be initiated by passing `shouldIgnoreWarnings` as `true`. This value allows the recovery operations to continue despite the validation warnings. Alternatively, the user can choose to resolve the validation warnings and trigger another failover action on recovery plan.
	*/
	ShouldIgnoreWarnings *bool `json:"shouldIgnoreWarnings,omitempty"`
}

func (p *TestFailoverSpec) MarshalJSON() ([]byte, error) {
	type TestFailoverSpecProxy TestFailoverSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*TestFailoverSpecProxy
		FailoverDirections []import1.FailoverDirection `json:"failoverDirections,omitempty"`
		Name               *string                     `json:"name,omitempty"`
	}{
		TestFailoverSpecProxy: (*TestFailoverSpecProxy)(p),
		FailoverDirections:    p.FailoverDirections,
		Name:                  p.Name,
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

func (p *TestFailoverSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TestFailoverSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTestFailoverSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ShouldIgnoreWarnings != nil {
		p.ShouldIgnoreWarnings = known.ShouldIgnoreWarnings
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverDirections")
	delete(allFields, "name")
	delete(allFields, "shouldIgnoreWarnings")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTestFailoverSpec() *TestFailoverSpec {
	p := new(TestFailoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.TestFailoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldIgnoreWarnings = new(bool)
	*p.ShouldIgnoreWarnings = false

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/unplanned-failover Post operation
*/
type UnplannedFailoverRecoveryPlanApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUnplannedFailoverRecoveryPlanApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UnplannedFailoverRecoveryPlanApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UnplannedFailoverRecoveryPlanApiResponse

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

func (p *UnplannedFailoverRecoveryPlanApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UnplannedFailoverRecoveryPlanApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUnplannedFailoverRecoveryPlanApiResponse()

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

func NewUnplannedFailoverRecoveryPlanApiResponse() *UnplannedFailoverRecoveryPlanApiResponse {
	p := new(UnplannedFailoverRecoveryPlanApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.UnplannedFailoverRecoveryPlanApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UnplannedFailoverRecoveryPlanApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UnplannedFailoverRecoveryPlanApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUnplannedFailoverRecoveryPlanApiResponseData()
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
The unplanned failover recovery plan action is used to perform an unplanned failover on the recovery plan.
*/
type UnplannedFailoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []import1.FailoverDirection `json:"failoverDirections"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name"`
	/*
	  Point in time from which to restore the entities during an `UNPLANNED_FAILOVER` operation. Only ISO-8601 formatted timestamps are supported. For example, `2023-01-02T03:04:05Z`.<br> When specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the specified timestamp. If not specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the start of the Recovery plan job.
	*/
	RecoveryReferenceTime *time.Time `json:"recoveryReferenceTime,omitempty"`
	/*
	  Indicates whether to continue the recovery plan action despite validation warnings. For example, if the IP of some virtual machines (VMs) cannot be preserved after recovery, the Recovery Plan action can still be initiated by passing `shouldIgnoreWarnings` as `true`. This value allows the recovery operations to continue despite the validation warnings. Alternatively, the user can choose to resolve the validation warnings and trigger another failover action on recovery plan.
	*/
	ShouldIgnoreWarnings *bool `json:"shouldIgnoreWarnings,omitempty"`
}

func (p *UnplannedFailoverSpec) MarshalJSON() ([]byte, error) {
	type UnplannedFailoverSpecProxy UnplannedFailoverSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UnplannedFailoverSpecProxy
		FailoverDirections []import1.FailoverDirection `json:"failoverDirections,omitempty"`
		Name               *string                     `json:"name,omitempty"`
	}{
		UnplannedFailoverSpecProxy: (*UnplannedFailoverSpecProxy)(p),
		FailoverDirections:         p.FailoverDirections,
		Name:                       p.Name,
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

func (p *UnplannedFailoverSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UnplannedFailoverSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUnplannedFailoverSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.RecoveryReferenceTime != nil {
		p.RecoveryReferenceTime = known.RecoveryReferenceTime
	}
	if known.ShouldIgnoreWarnings != nil {
		p.ShouldIgnoreWarnings = known.ShouldIgnoreWarnings
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverDirections")
	delete(allFields, "name")
	delete(allFields, "recoveryReferenceTime")
	delete(allFields, "shouldIgnoreWarnings")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUnplannedFailoverSpec() *UnplannedFailoverSpec {
	p := new(UnplannedFailoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.UnplannedFailoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldIgnoreWarnings = new(bool)
	*p.ShouldIgnoreWarnings = false

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/operations/recovery-plans/{recoveryPlanExtId}/$actions/validate Post operation
*/
type ValidateRecoveryPlanApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfValidateRecoveryPlanApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ValidateRecoveryPlanApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ValidateRecoveryPlanApiResponse

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

func (p *ValidateRecoveryPlanApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ValidateRecoveryPlanApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewValidateRecoveryPlanApiResponse()

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

func NewValidateRecoveryPlanApiResponse() *ValidateRecoveryPlanApiResponse {
	p := new(ValidateRecoveryPlanApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.ValidateRecoveryPlanApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ValidateRecoveryPlanApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ValidateRecoveryPlanApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfValidateRecoveryPlanApiResponseData()
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

type ValidateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []import1.FailoverDirection `json:"failoverDirections"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name"`
}

func (p *ValidateSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ValidateSpec

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

func (p *ValidateSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ValidateSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewValidateSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverDirections")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewValidateSpec() *ValidateSpec {
	p := new(ValidateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.operations.ValidateSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfCleanupRecoveryPlanResourcesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
}

func NewOneOfCleanupRecoveryPlanResourcesApiResponseData() *OneOfCleanupRecoveryPlanResourcesApiResponseData {
	p := new(OneOfCleanupRecoveryPlanResourcesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCleanupRecoveryPlanResourcesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCleanupRecoveryPlanResourcesApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
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

func (p *OneOfCleanupRecoveryPlanResourcesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCleanupRecoveryPlanResourcesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCleanupRecoveryPlanResourcesApiResponseData"))
}

func (p *OneOfCleanupRecoveryPlanResourcesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCleanupRecoveryPlanResourcesApiResponseData")
}

type OneOfPlannedFailoverRecoveryPlanApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
}

func NewOneOfPlannedFailoverRecoveryPlanApiResponseData() *OneOfPlannedFailoverRecoveryPlanApiResponseData {
	p := new(OneOfPlannedFailoverRecoveryPlanApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPlannedFailoverRecoveryPlanApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPlannedFailoverRecoveryPlanApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
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

func (p *OneOfPlannedFailoverRecoveryPlanApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfPlannedFailoverRecoveryPlanApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPlannedFailoverRecoveryPlanApiResponseData"))
}

func (p *OneOfPlannedFailoverRecoveryPlanApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfPlannedFailoverRecoveryPlanApiResponseData")
}

type OneOfUnplannedFailoverRecoveryPlanApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
}

func NewOneOfUnplannedFailoverRecoveryPlanApiResponseData() *OneOfUnplannedFailoverRecoveryPlanApiResponseData {
	p := new(OneOfUnplannedFailoverRecoveryPlanApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUnplannedFailoverRecoveryPlanApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUnplannedFailoverRecoveryPlanApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
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

func (p *OneOfUnplannedFailoverRecoveryPlanApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUnplannedFailoverRecoveryPlanApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUnplannedFailoverRecoveryPlanApiResponseData"))
}

func (p *OneOfUnplannedFailoverRecoveryPlanApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUnplannedFailoverRecoveryPlanApiResponseData")
}

type OneOfTestFailoverRecoveryPlanApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
}

func NewOneOfTestFailoverRecoveryPlanApiResponseData() *OneOfTestFailoverRecoveryPlanApiResponseData {
	p := new(OneOfTestFailoverRecoveryPlanApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTestFailoverRecoveryPlanApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTestFailoverRecoveryPlanApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
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

func (p *OneOfTestFailoverRecoveryPlanApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfTestFailoverRecoveryPlanApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTestFailoverRecoveryPlanApiResponseData"))
}

func (p *OneOfTestFailoverRecoveryPlanApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfTestFailoverRecoveryPlanApiResponseData")
}

type OneOfValidateRecoveryPlanApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
}

func NewOneOfValidateRecoveryPlanApiResponseData() *OneOfValidateRecoveryPlanApiResponseData {
	p := new(OneOfValidateRecoveryPlanApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfValidateRecoveryPlanApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfValidateRecoveryPlanApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
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

func (p *OneOfValidateRecoveryPlanApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfValidateRecoveryPlanApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfValidateRecoveryPlanApiResponseData"))
}

func (p *OneOfValidateRecoveryPlanApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfValidateRecoveryPlanApiResponseData")
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
