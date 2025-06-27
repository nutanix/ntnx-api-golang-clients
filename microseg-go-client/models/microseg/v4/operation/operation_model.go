/*
 * Generated file models/microseg/v4/operation/operation_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Flow Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module microseg.v4.operation of Nutanix Flow Management APIs
*/
package operation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
	"time"
)

/*
_Body for requesting Config Upgrade Summary_

It contains two fields:
  - policyNames
  - shouldIncludeSecureSubnetsInfo
*/
type DryRunSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of FNS policies selected for upgrade.
	*/
	PolicyNames []string `json:"policyNames,omitempty"`
	/*
	  If set to true, displays upgrade information only for subnets with VMs secured by NSPs.<br>
	Else, the upgrade information is displayed for all the subnets.
	*/
	ShouldIncludeSecureSubnetsInfo *bool `json:"shouldIncludeSecureSubnetsInfo,omitempty"`
}

func (p *DryRunSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DryRunSpec

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

func (p *DryRunSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DryRunSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DryRunSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "policyNames")
	delete(allFields, "shouldIncludeSecureSubnetsInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDryRunSpec() *DryRunSpec {
	p := new(DryRunSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.DryRunSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DryRunSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CountSummary *TotalPolicyCount `json:"countSummary,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	PolicySummaries []PolicyUpgradeSummary `json:"policySummaries,omitempty"`

	SubnetSummaries []SubnetUpgradeSummary `json:"subnetSummaries,omitempty"`
	/*
	  Summary of failures from the flow upgrade dry-run summary report.
	*/
	SummaryFailures []import2.Message `json:"summaryFailures,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DryRunSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DryRunSummary

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

func (p *DryRunSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DryRunSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DryRunSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "countSummary")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "policySummaries")
	delete(allFields, "subnetSummaries")
	delete(allFields, "summaryFailures")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDryRunSummary() *DryRunSummary {
	p := new(DryRunSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.DryRunSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cumulative Policy counts pre and post Flow upgrade to FNS Next-Gen.
Contains total count of FNS Next-Gen policies that'll be created post successful upgrade,
total count of FNS policies before upgrade and the total count of FNS Next-Gen policies
which were system generated as a result of FNS to FNS Next-Gen conversion.<br>
*/
type PolicyCount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of policies after the upgrade.
	*/
	NewPoliciesCount *int `json:"newPoliciesCount"`
	/*
	  Number of policies before the migration.
	*/
	OldPoliciesCount *int `json:"oldPoliciesCount"`
	/*
	  Count of isolation policies that would be generated post upgrade to FNS Next-Gen.
	*/
	SystemDefinedPoliciesCount *int `json:"systemDefinedPoliciesCount"`
}

func (p *PolicyCount) MarshalJSON() ([]byte, error) {
	type PolicyCountProxy PolicyCount

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PolicyCountProxy
		NewPoliciesCount           *int `json:"newPoliciesCount,omitempty"`
		OldPoliciesCount           *int `json:"oldPoliciesCount,omitempty"`
		SystemDefinedPoliciesCount *int `json:"systemDefinedPoliciesCount,omitempty"`
	}{
		PolicyCountProxy:           (*PolicyCountProxy)(p),
		NewPoliciesCount:           p.NewPoliciesCount,
		OldPoliciesCount:           p.OldPoliciesCount,
		SystemDefinedPoliciesCount: p.SystemDefinedPoliciesCount,
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

func (p *PolicyCount) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyCount
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyCount(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "newPoliciesCount")
	delete(allFields, "oldPoliciesCount")
	delete(allFields, "systemDefinedPoliciesCount")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyCount() *PolicyCount {
	p := new(PolicyCount)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.PolicyCount"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Count statistics for number of policies before and after the upgrade as well as type of policy.
*/
type PolicyCountByType struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CountSummary *PolicyCount `json:"countSummary,omitempty"`

	Type *SecurityPolicyType `json:"type,omitempty"`
}

func (p *PolicyCountByType) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PolicyCountByType

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

func (p *PolicyCountByType) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyCountByType
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyCountByType(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "countSummary")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyCountByType() *PolicyCountByType {
	p := new(PolicyCountByType)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.PolicyCountByType"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The meta information about a Flow Network Security Policy.<br>
The info captures the secured Categories, state of policy, policy type, last modified time,
previewId and options.<br>
The PreviewId can be used to fetch the complete FNS Next-Gen policy preview using Preview GET API.<br>
A sample call would look like
```
/microseg/v4.0.a1/operation/policy-preview/8f094d6c-d7b2-32c2-8223-462c6c5e06db
```
*/
type PolicyMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A user defined annotation for a policy metadata during upgrade.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Last update time for the network security policy.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  Name of the network security policy.
	*/
	Name *string `json:"name,omitempty"`

	Options *SecurityPolicyOptions `json:"options,omitempty"`
	/*
	  This is a system generated identifier which can be used to preview the complete FNS Next-Gen
	policy corresponding with a FNS policy using Preview GET API.
	A sample call would look like
	```
	/microseg/v4.0.a1/operation/policy-preview/8f094d6c-d7b2-32c2-8223-462c6c5e06db
	```
	*/
	PreviewReference *string `json:"previewReference,omitempty"`
	/*
	  List of categories external IDs being secured by the Flow Network Security Policy.
	*/
	SecuredGroupCategoryReferences []string `json:"securedGroupCategoryReferences,omitempty"`

	State *import3.SecurityPolicyState `json:"state"`

	Type *SecurityPolicyType `json:"type,omitempty"`
}

func (p *PolicyMetadata) MarshalJSON() ([]byte, error) {
	type PolicyMetadataProxy PolicyMetadata

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PolicyMetadataProxy
		State *import3.SecurityPolicyState `json:"state,omitempty"`
	}{
		PolicyMetadataProxy: (*PolicyMetadataProxy)(p),
		State:               p.State,
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

func (p *PolicyMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyMetadata(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "description")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "name")
	delete(allFields, "options")
	delete(allFields, "previewReference")
	delete(allFields, "securedGroupCategoryReferences")
	delete(allFields, "state")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyMetadata() *PolicyMetadata {
	p := new(PolicyMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.PolicyMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PolicyPreview struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	NewPolicySpec *import3.NetworkSecurityPolicy `json:"newPolicySpec,omitempty"`
	/*
	  This field corresponds to a list of policy specifications which captures the system generated FNS Next-Gen policies
	as a result of a single FNS policy upgrade.
	*/
	SystemGeneratedPolicies []import3.NetworkSecurityPolicy `json:"systemGeneratedPolicies,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *PolicyPreview) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PolicyPreview

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

func (p *PolicyPreview) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyPreview
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyPreview(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "newPolicySpec")
	delete(allFields, "systemGeneratedPolicies")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyPreview() *PolicyPreview {
	p := new(PolicyPreview)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.PolicyPreview"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specifies the policy name along with the policy metadata.<br>
For cases where a single FNS policy conversion results in multiple FNS Next-Gen policies,
it also contains a list of policy metadata corresponding to those system generated policies.
*/
type PolicyUpgradeSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PolicyMetadata *PolicyMetadata `json:"policyMetadata"`
	/*
	  Name of the Flow Network Security Policy.<br>
	It can have a maximum length of 63 characters.
	*/
	PolicyName *string `json:"policyName"`
	/*
	  A list of Policy metadata of system generated isolation policies that would get created as part of
	a successful Flow upgrade.
	*/
	SystemGenPolicyMetadata []PolicyMetadata `json:"systemGenPolicyMetadata,omitempty"`
}

func (p *PolicyUpgradeSummary) MarshalJSON() ([]byte, error) {
	type PolicyUpgradeSummaryProxy PolicyUpgradeSummary

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PolicyUpgradeSummaryProxy
		PolicyMetadata *PolicyMetadata `json:"policyMetadata,omitempty"`
		PolicyName     *string         `json:"policyName,omitempty"`
	}{
		PolicyUpgradeSummaryProxy: (*PolicyUpgradeSummaryProxy)(p),
		PolicyMetadata:            p.PolicyMetadata,
		PolicyName:                p.PolicyName,
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

func (p *PolicyUpgradeSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PolicyUpgradeSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PolicyUpgradeSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "policyMetadata")
	delete(allFields, "policyName")
	delete(allFields, "systemGenPolicyMetadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPolicyUpgradeSummary() *PolicyUpgradeSummary {
	p := new(PolicyUpgradeSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.PolicyUpgradeSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Policy-wide options for a security policy.
*/
type SecurityPolicyOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  If Hitlog is enabled.
	*/
	IsHitlogEnabled *bool `json:"isHitlogEnabled,omitempty"`
	/*
	  If Ipv6 Traffic is allowed.
	*/
	IsIpv6TrafficAllowed *bool `json:"isIpv6TrafficAllowed,omitempty"`
}

func (p *SecurityPolicyOptions) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SecurityPolicyOptions

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

func (p *SecurityPolicyOptions) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SecurityPolicyOptions
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SecurityPolicyOptions(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isHitlogEnabled")
	delete(allFields, "isIpv6TrafficAllowed")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSecurityPolicyOptions() *SecurityPolicyOptions {
	p := new(SecurityPolicyOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.SecurityPolicyOptions"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsHitlogEnabled = new(bool)
	*p.IsHitlogEnabled = false
	p.IsIpv6TrafficAllowed = new(bool)
	*p.IsIpv6TrafficAllowed = false

	return p
}

/*
Defines the type of rules that can be used in a policy.<br>
It can be one of following types
  - QUARANTINE POLICY
  - ISOLATION POLICY
  - APPLICATION POLICY
  - AD POLICY
*/
type SecurityPolicyType int

const (
	SECURITYPOLICYTYPE_UNKNOWN     SecurityPolicyType = 0
	SECURITYPOLICYTYPE_REDACTED    SecurityPolicyType = 1
	SECURITYPOLICYTYPE_QUARANTINE  SecurityPolicyType = 2
	SECURITYPOLICYTYPE_ISOLATION   SecurityPolicyType = 3
	SECURITYPOLICYTYPE_APPLICATION SecurityPolicyType = 4
	SECURITYPOLICYTYPE_AD          SecurityPolicyType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SecurityPolicyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
		"AD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SecurityPolicyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
		"AD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SecurityPolicyType) index(name string) SecurityPolicyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
		"AD",
	}
	for idx := range names {
		if names[idx] == name {
			return SecurityPolicyType(idx)
		}
	}
	return SECURITYPOLICYTYPE_UNKNOWN
}

func (e *SecurityPolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SecurityPolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SecurityPolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SecurityPolicyType) Ref() *SecurityPolicyType {
	return &e
}

/*
Subnet information to be communicated for subnet upgrade as part of Flow upgrade to FNS Next-Gen.<br>
The info includes vlanID, vlanName and subnetUuid.
*/
type SubnetUpgradeSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Subnet ExtID used in upgrade.
	*/
	SubnetReference *string `json:"subnetReference,omitempty"`
	/*
	  Vlan id used in the upgrade.
	*/
	VlanID *int `json:"vlanID"`
	/*
	  Vlan name in the upgrade.
	*/
	VlanName *string `json:"vlanName,omitempty"`
}

func (p *SubnetUpgradeSummary) MarshalJSON() ([]byte, error) {
	type SubnetUpgradeSummaryProxy SubnetUpgradeSummary

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SubnetUpgradeSummaryProxy
		VlanID *int `json:"vlanID,omitempty"`
	}{
		SubnetUpgradeSummaryProxy: (*SubnetUpgradeSummaryProxy)(p),
		VlanID:                    p.VlanID,
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

func (p *SubnetUpgradeSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SubnetUpgradeSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = SubnetUpgradeSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "subnetReference")
	delete(allFields, "vlanID")
	delete(allFields, "vlanName")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewSubnetUpgradeSummary() *SubnetUpgradeSummary {
	p := new(SubnetUpgradeSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.SubnetUpgradeSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cumulative and type based policy counts pre and post Flow upgrade to FNS Next-Gen.<br>
Contains the summary total policy counts and the policy counts grouped by policy type.
*/
type TotalPolicyCount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PolicyTypeCountsSummary []PolicyCountByType `json:"policyTypeCountsSummary,omitempty"`

	TotalCountsSummary *PolicyCount `json:"totalCountsSummary"`
}

func (p *TotalPolicyCount) MarshalJSON() ([]byte, error) {
	type TotalPolicyCountProxy TotalPolicyCount

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*TotalPolicyCountProxy
		TotalCountsSummary *PolicyCount `json:"totalCountsSummary,omitempty"`
	}{
		TotalPolicyCountProxy: (*TotalPolicyCountProxy)(p),
		TotalCountsSummary:    p.TotalCountsSummary,
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

func (p *TotalPolicyCount) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TotalPolicyCount
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TotalPolicyCount(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "policyTypeCountsSummary")
	delete(allFields, "totalCountsSummary")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTotalPolicyCount() *TotalPolicyCount {
	p := new(TotalPolicyCount)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.TotalPolicyCount"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
__Request body for Config Upgrade__

It contains one field:
- spec

The spec can be one of two types based on $dryrun boolean value passed in query parameter.

A sample request body with the $dryrun set to true in query parameter would look like this:
```
{
"spec": {
  "policyNames" : [],
  "shouldIncludeSecureSubnetsInfo" : false
  }
}
```

A sample request body with the $dryrun not set as part of query parameter would look like this:
```
{
"spec": {
  "shouldSetMonitorMode" : true,
  "shouldUpgradeSecuredSubnetsOnly" : false
  }
}
```
*/
type UpgradeJob struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
	/*
	  Specifies the type of upgradeJobSpec (UpgradeSpec or DryRunSpec).
	*/
	Spec *OneOfUpgradeJobSpec `json:"spec"`
}

func (p *UpgradeJob) MarshalJSON() ([]byte, error) {
	type UpgradeJobProxy UpgradeJob

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UpgradeJobProxy
		Spec *OneOfUpgradeJobSpec `json:"spec,omitempty"`
	}{
		UpgradeJobProxy: (*UpgradeJobProxy)(p),
		Spec:            p.Spec,
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

func (p *UpgradeJob) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpgradeJob
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpgradeJob(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$specItemDiscriminator")
	delete(allFields, "spec")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUpgradeJob() *UpgradeJob {
	p := new(UpgradeJob)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.UpgradeJob"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpgradeJob) GetSpec() interface{} {
	if nil == p.Spec {
		return nil
	}
	return p.Spec.GetValue()
}

func (p *UpgradeJob) SetSpec(v interface{}) error {
	if nil == p.Spec {
		p.Spec = NewOneOfUpgradeJobSpec()
	}
	e := p.Spec.SetValue(v)
	if nil == e {
		if nil == p.SpecItemDiscriminator_ {
			p.SpecItemDiscriminator_ = new(string)
		}
		*p.SpecItemDiscriminator_ = *p.Spec.Discriminator
	}
	return e
}

/*
__Request body for Actual Config Upgrade__

It contains two required fields:
- forceMonitor
- upgradeSecuredSubnets
*/
type UpgradeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  If set to true, it forces all the policies to be in monitor mode after the upgrade completes.
	For the default false case, all the policies would remain in the same pre-upgrade state.
	*/
	ShouldSetMonitorMode *bool `json:"shouldSetMonitorMode"`
	/*
	  If set to true, only the subnets with NSP secured VMs are upgraded. For the value set to false, all the subnets are
	upgraded.
	*/
	ShouldUpgradeSecuredSubnetsOnly *bool `json:"shouldUpgradeSecuredSubnetsOnly,omitempty"`
}

func (p *UpgradeSpec) MarshalJSON() ([]byte, error) {
	type UpgradeSpecProxy UpgradeSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UpgradeSpecProxy
		ShouldSetMonitorMode *bool `json:"shouldSetMonitorMode,omitempty"`
	}{
		UpgradeSpecProxy:     (*UpgradeSpecProxy)(p),
		ShouldSetMonitorMode: p.ShouldSetMonitorMode,
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

func (p *UpgradeSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpgradeSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpgradeSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "shouldSetMonitorMode")
	delete(allFields, "shouldUpgradeSecuredSubnetsOnly")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUpgradeSpec() *UpgradeSpec {
	p := new(UpgradeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.operation.UpgradeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfUpgradeJobSpec struct {
	Discriminator *string      `json:"-"`
	ObjectType_   *string      `json:"-"`
	oneOfType0    *UpgradeSpec `json:"-"`
	oneOfType1    *DryRunSpec  `json:"-"`
}

func NewOneOfUpgradeJobSpec() *OneOfUpgradeJobSpec {
	p := new(OneOfUpgradeJobSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpgradeJobSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpgradeJobSpec is nil"))
	}
	switch v.(type) {
	case UpgradeSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UpgradeSpec)
		}
		*p.oneOfType0 = v.(UpgradeSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case DryRunSpec:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(DryRunSpec)
		}
		*p.oneOfType1 = v.(DryRunSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpgradeJobSpec) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfUpgradeJobSpec) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UpgradeSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.operation.UpgradeSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UpgradeSpec)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType1 := new(DryRunSpec)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "microseg.v4.operation.DryRunSpec" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(DryRunSpec)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpgradeJobSpec"))
}

func (p *OneOfUpgradeJobSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfUpgradeJobSpec")
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
