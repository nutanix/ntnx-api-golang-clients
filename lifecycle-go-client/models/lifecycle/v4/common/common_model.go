/*
 * Generated file models/lifecycle/v4/common/common_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Common Lifecycle resources.
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/security/v4/config"
)

/*
Available version types.
*/
type AvailableVersionStatus int

const (
	AVAILABLEVERSIONSTATUS_UNKNOWN     AvailableVersionStatus = 0
	AVAILABLEVERSIONSTATUS_REDACTED    AvailableVersionStatus = 1
	AVAILABLEVERSIONSTATUS_RECOMMENDED AvailableVersionStatus = 2
	AVAILABLEVERSIONSTATUS_CRITICAL    AvailableVersionStatus = 3
	AVAILABLEVERSIONSTATUS_LATEST      AvailableVersionStatus = 4
	AVAILABLEVERSIONSTATUS_DEPRECATED  AvailableVersionStatus = 5
	AVAILABLEVERSIONSTATUS_EMERGENCY   AvailableVersionStatus = 6
	AVAILABLEVERSIONSTATUS_AVAILABLE   AvailableVersionStatus = 7
	AVAILABLEVERSIONSTATUS_LTS         AvailableVersionStatus = 8
	AVAILABLEVERSIONSTATUS_STS         AvailableVersionStatus = 9
	AVAILABLEVERSIONSTATUS_ESTS        AvailableVersionStatus = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AvailableVersionStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AvailableVersionStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AvailableVersionStatus) index(name string) AvailableVersionStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RECOMMENDED",
		"CRITICAL",
		"LATEST",
		"DEPRECATED",
		"EMERGENCY",
		"AVAILABLE",
		"LTS",
		"STS",
		"ESTS",
	}
	for idx := range names {
		if names[idx] == name {
			return AvailableVersionStatus(idx)
		}
	}
	return AVAILABLEVERSIONSTATUS_UNKNOWN
}

func (e *AvailableVersionStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AvailableVersionStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AvailableVersionStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AvailableVersionStatus) Ref() *AvailableVersionStatus {
	return &e
}

/*
Checksum type for a third party image.
*/
type CheckSumType int

const (
	CHECKSUMTYPE_UNKNOWN  CheckSumType = 0
	CHECKSUMTYPE_REDACTED CheckSumType = 1
	CHECKSUMTYPE_SHASUM   CheckSumType = 2
	CHECKSUMTYPE_HEX_MD5  CheckSumType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CheckSumType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CheckSumType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CheckSumType) index(name string) CheckSumType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHASUM",
		"HEX_MD5",
	}
	for idx := range names {
		if names[idx] == name {
			return CheckSumType(idx)
		}
	}
	return CHECKSUMTYPE_UNKNOWN
}

func (e *CheckSumType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CheckSumType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CheckSumType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CheckSumType) Ref() *CheckSumType {
	return &e
}

/*
This field indicates LCM is running on a Prism Element or Prism Central cluster.
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN       ClusterType = 0
	CLUSTERTYPE_REDACTED      ClusterType = 1
	CLUSTERTYPE_PRISM_CENTRAL ClusterType = 2
	CLUSTERTYPE_AOS           ClusterType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"AOS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"AOS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterType) index(name string) ClusterType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"AOS",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterType(idx)
		}
	}
	return CLUSTERTYPE_UNKNOWN
}

func (e *ClusterType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterType) Ref() *ClusterType {
	return &e
}

/*
Details of credential used for performing an LCM operations.
*/
type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	CredentialDetailItemDiscriminator_ *string `json:"$credentialDetailItemDiscriminator,omitempty"`
	/*
	  Reference of pre-created credential in credential-store or raw details of credential
	*/
	CredentialDetail *OneOfCredentialCredentialDetail `json:"credentialDetail"`
}

func (p *Credential) MarshalJSON() ([]byte, error) {
	type CredentialProxy Credential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CredentialProxy
		CredentialDetail *OneOfCredentialCredentialDetail `json:"credentialDetail,omitempty"`
	}{
		CredentialProxy:  (*CredentialProxy)(p),
		CredentialDetail: p.CredentialDetail,
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
	*p = *NewCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CredentialDetailItemDiscriminator_ != nil {
		p.CredentialDetailItemDiscriminator_ = known.CredentialDetailItemDiscriminator_
	}
	if known.CredentialDetail != nil {
		p.CredentialDetail = known.CredentialDetail
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$credentialDetailItemDiscriminator")
	delete(allFields, "credentialDetail")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.Credential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Credential) GetCredentialDetail() interface{} {
	if nil == p.CredentialDetail {
		return nil
	}
	return p.CredentialDetail.GetValue()
}

func (p *Credential) SetCredentialDetail(v interface{}) error {
	if nil == p.CredentialDetail {
		p.CredentialDetail = NewOneOfCredentialCredentialDetail()
	}
	e := p.CredentialDetail.SetValue(v)
	if nil == e {
		if nil == p.CredentialDetailItemDiscriminator_ {
			p.CredentialDetailItemDiscriminator_ = new(string)
		}
		*p.CredentialDetailItemDiscriminator_ = *p.CredentialDetail.Discriminator
	}
	return e
}

/*
Credential Reference from the Credential Store.
*/
type CredentialReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the credential.
	*/
	CredentialExtId *string `json:"credentialExtId"`
}

func (p *CredentialReference) MarshalJSON() ([]byte, error) {
	type CredentialReferenceProxy CredentialReference

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CredentialReferenceProxy
		CredentialExtId *string `json:"credentialExtId,omitempty"`
	}{
		CredentialReferenceProxy: (*CredentialReferenceProxy)(p),
		CredentialExtId:          p.CredentialExtId,
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

func (p *CredentialReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CredentialReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCredentialReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CredentialExtId != nil {
		p.CredentialExtId = known.CredentialExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credentialExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCredentialReference() *CredentialReference {
	p := new(CredentialReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.CredentialReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for deployment of entities.
*/
type DeploySpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of deploy specifications for use in the recommendations API.
	*/
	EntityDeploySpecs []EntityDeploySpec `json:"entityDeploySpecs"`
}

func (p *DeploySpec) MarshalJSON() ([]byte, error) {
	type DeploySpecProxy DeploySpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DeploySpecProxy
		EntityDeploySpecs []EntityDeploySpec `json:"entityDeploySpecs,omitempty"`
	}{
		DeploySpecProxy:   (*DeploySpecProxy)(p),
		EntityDeploySpecs: p.EntityDeploySpecs,
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

func (p *DeploySpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeploySpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeploySpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityDeploySpecs != nil {
		p.EntityDeploySpecs = known.EntityDeploySpecs
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityDeploySpecs")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeploySpec() *DeploySpec {
	p := new(DeploySpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.DeploySpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The list of properties that can be expanded on the LCM entity.
*/
type EntityBaseModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *EntityBaseModel) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityBaseModel

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

func (p *EntityBaseModel) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityBaseModel
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityBaseModel()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityClass != nil {
		p.EntityClass = known.EntityClass
	}
	if known.EntityModel != nil {
		p.EntityModel = known.EntityModel
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.EntityVersion != nil {
		p.EntityVersion = known.EntityVersion
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HardwareFamily != nil {
		p.HardwareFamily = known.HardwareFamily
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
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityBaseModel() *EntityBaseModel {
	p := new(EntityBaseModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityBaseModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A specification defining the entity being deployed and its version.
*/
type EntityDeploySpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityIdentifier *EntityBaseModel `json:"entityIdentifier"`
}

func (p *EntityDeploySpec) MarshalJSON() ([]byte, error) {
	type EntityDeploySpecProxy EntityDeploySpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EntityDeploySpecProxy
		EntityIdentifier *EntityBaseModel `json:"entityIdentifier,omitempty"`
	}{
		EntityDeploySpecProxy: (*EntityDeploySpecProxy)(p),
		EntityIdentifier:      p.EntityIdentifier,
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

func (p *EntityDeploySpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityDeploySpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityDeploySpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityIdentifier != nil {
		p.EntityIdentifier = known.EntityIdentifier
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityIdentifier")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityDeploySpec() *EntityDeploySpec {
	p := new(EntityDeploySpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityDeploySpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of an LCM entity.
*/
type EntityType int

const (
	ENTITYTYPE_UNKNOWN  EntityType = 0
	ENTITYTYPE_REDACTED EntityType = 1
	ENTITYTYPE_SOFTWARE EntityType = 2
	ENTITYTYPE_FIRMWARE EntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EntityType) index(name string) EntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
	}
	for idx := range names {
		if names[idx] == name {
			return EntityType(idx)
		}
	}
	return ENTITYTYPE_UNKNOWN
}

func (e *EntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EntityType) Ref() *EntityType {
	return &e
}

/*
Specification for running an update operation.
*/
type EntityUpdateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the LCM entity.
	*/
	EntityUuid *string `json:"entityUuid"`
	/*
	  Version to upgrade to.
	*/
	ToVersion *string `json:"toVersion"`
}

func (p *EntityUpdateSpec) MarshalJSON() ([]byte, error) {
	type EntityUpdateSpecProxy EntityUpdateSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EntityUpdateSpecProxy
		EntityUuid *string `json:"entityUuid,omitempty"`
		ToVersion  *string `json:"toVersion,omitempty"`
	}{
		EntityUpdateSpecProxy: (*EntityUpdateSpecProxy)(p),
		EntityUuid:            p.EntityUuid,
		ToVersion:             p.ToVersion,
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

func (p *EntityUpdateSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityUpdateSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityUpdateSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityUuid != nil {
		p.EntityUuid = known.EntityUuid
	}
	if known.ToVersion != nil {
		p.ToVersion = known.ToVersion
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityUuid")
	delete(allFields, "toVersion")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityUpdateSpec() *EntityUpdateSpec {
	p := new(EntityUpdateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.EntityUpdateSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File format of the exported data.
*/
type FileFormat int

const (
	FILEFORMAT_UNKNOWN  FileFormat = 0
	FILEFORMAT_REDACTED FileFormat = 1
	FILEFORMAT_CSV      FileFormat = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FileFormat) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CSV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FileFormat) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CSV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FileFormat) index(name string) FileFormat {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CSV",
	}
	for idx := range names {
		if names[idx] == name {
			return FileFormat(idx)
		}
	}
	return FILEFORMAT_UNKNOWN
}

func (e *FileFormat) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FileFormat:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FileFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FileFormat) Ref() *FileFormat {
	return &e
}

/*
Type of Hypervisor present in the cluster.
*/
type HypervisorType int

const (
	HYPERVISORTYPE_UNKNOWN  HypervisorType = 0
	HYPERVISORTYPE_REDACTED HypervisorType = 1
	HYPERVISORTYPE_ESX      HypervisorType = 2
	HYPERVISORTYPE_AHV      HypervisorType = 3
	HYPERVISORTYPE_HYPERV   HypervisorType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HypervisorType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HypervisorType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HypervisorType) index(name string) HypervisorType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESX",
		"AHV",
		"HYPERV",
	}
	for idx := range names {
		if names[idx] == name {
			return HypervisorType(idx)
		}
	}
	return HYPERVISORTYPE_UNKNOWN
}

func (e *HypervisorType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HypervisorType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HypervisorType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HypervisorType) Ref() *HypervisorType {
	return &e
}

/*
Details of the in progress LCM operation.
*/
type InProgressOpDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  Task ext id of the in progress LCM operation.
	*/
	TaskExtId *string `json:"taskExtId,omitempty"`
}

func (p *InProgressOpDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InProgressOpDetails

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

func (p *InProgressOpDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InProgressOpDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewInProgressOpDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.TaskExtId != nil {
		p.TaskExtId = known.TaskExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "operationType")
	delete(allFields, "taskExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewInProgressOpDetails() *InProgressOpDetails {
	p := new(InProgressOpDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.InProgressOpDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
MD5Sum of the bundle.
*/
type LcmMd5Sum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hex digest of the MD5 sum.
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *LcmMd5Sum) MarshalJSON() ([]byte, error) {
	type LcmMd5SumProxy LcmMd5Sum

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LcmMd5SumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		LcmMd5SumProxy: (*LcmMd5SumProxy)(p),
		HexDigest:      p.HexDigest,
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

func (p *LcmMd5Sum) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LcmMd5Sum
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLcmMd5Sum()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HexDigest != nil {
		p.HexDigest = known.HexDigest
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hexDigest")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLcmMd5Sum() *LcmMd5Sum {
	p := new(LcmMd5Sum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LcmMd5Sum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Sha256sum of the bundle.
*/
type LcmSha256Sum struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Hex digest of the SHA256 sum.
	*/
	HexDigest *string `json:"hexDigest"`
}

func (p *LcmSha256Sum) MarshalJSON() ([]byte, error) {
	type LcmSha256SumProxy LcmSha256Sum

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LcmSha256SumProxy
		HexDigest *string `json:"hexDigest,omitempty"`
	}{
		LcmSha256SumProxy: (*LcmSha256SumProxy)(p),
		HexDigest:         p.HexDigest,
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

func (p *LcmSha256Sum) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LcmSha256Sum
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLcmSha256Sum()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HexDigest != nil {
		p.HexDigest = known.HexDigest
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hexDigest")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLcmSha256Sum() *LcmSha256Sum {
	p := new(LcmSha256Sum)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LcmSha256Sum"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Location info corresponds to a tuple of location type (either node/cluster) and ExtID
*/
type LocationInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the location.
	*/
	LocationName *string `json:"locationName,omitempty"`

	LocationType *LocationType `json:"locationType,omitempty"`
	/*
	  Location UUID of the resource.
	*/
	Uuid *string `json:"uuid,omitempty"`
}

func (p *LocationInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LocationInfo

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

func (p *LocationInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocationInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocationInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LocationName != nil {
		p.LocationName = known.LocationName
	}
	if known.LocationType != nil {
		p.LocationType = known.LocationType
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "locationName")
	delete(allFields, "locationType")
	delete(allFields, "uuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocationInfo() *LocationInfo {
	p := new(LocationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.LocationInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Scope of entity represented in LCM. This could be either Node or cluster type.
*/
type LocationType int

const (
	LOCATIONTYPE_UNKNOWN  LocationType = 0
	LOCATIONTYPE_REDACTED LocationType = 1
	LOCATIONTYPE_NODE     LocationType = 2
	LOCATIONTYPE_CLUSTER  LocationType = 3
	LOCATIONTYPE_PC       LocationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LocationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LocationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LocationType) index(name string) LocationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"CLUSTER",
		"PC",
	}
	for idx := range names {
		if names[idx] == name {
			return LocationType(idx)
		}
	}
	return LOCATIONTYPE_UNKNOWN
}

func (e *LocationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LocationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LocationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LocationType) Ref() *LocationType {
	return &e
}

/*
Cluster management server configuration used while updating clusters with ESX or Hyper-V.
*/
type ManagementServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType"`
	/*
	  IP address of the management server.
	*/
	Ip *string `json:"ip"`
	/*
	  Password to login to the management server.
	*/
	Password *string `json:"password"`
	/*
	  Username to login to the management server.
	*/
	Username *string `json:"username"`
}

func (p *ManagementServer) MarshalJSON() ([]byte, error) {
	type ManagementServerProxy ManagementServer

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ManagementServerProxy
		HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
		Ip             *string         `json:"ip,omitempty"`
		Password       *string         `json:"password,omitempty"`
		Username       *string         `json:"username,omitempty"`
	}{
		ManagementServerProxy: (*ManagementServerProxy)(p),
		HypervisorType:        p.HypervisorType,
		Ip:                    p.Ip,
		Password:              p.Password,
		Username:              p.Username,
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

func (p *ManagementServer) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ManagementServer
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewManagementServer()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HypervisorType != nil {
		p.HypervisorType = known.HypervisorType
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hypervisorType")
	delete(allFields, "ip")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewManagementServer() *ManagementServer {
	p := new(ManagementServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.ManagementServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of LCM upgrade notification generated. It can be any one of the values like Entity, Location, Generic or Workflow. The only types supported today are Entity and Location.
*/
type NotificationType int

const (
	NOTIFICATIONTYPE_UNKNOWN  NotificationType = 0
	NOTIFICATIONTYPE_REDACTED NotificationType = 1
	NOTIFICATIONTYPE_ENTITY   NotificationType = 2
	NOTIFICATIONTYPE_LOCATION NotificationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NotificationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NotificationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NotificationType) index(name string) NotificationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"LOCATION",
	}
	for idx := range names {
		if names[idx] == name {
			return NotificationType(idx)
		}
	}
	return NOTIFICATIONTYPE_UNKNOWN
}

func (e *NotificationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NotificationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NotificationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NotificationType) Ref() *NotificationType {
	return &e
}

/*
Details about credential required for running list of operations on a cluster.
*/
type OperationCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the credential.
	*/
	CredentialExtId *string `json:"credentialExtId,omitempty"`

	VendorManagementName *VendorManagementName `json:"vendorManagementName,omitempty"`
}

func (p *OperationCredential) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OperationCredential

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

func (p *OperationCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OperationCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOperationCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CredentialExtId != nil {
		p.CredentialExtId = known.CredentialExtId
	}
	if known.VendorManagementName != nil {
		p.VendorManagementName = known.VendorManagementName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credentialExtId")
	delete(allFields, "vendorManagementName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOperationCredential() *OperationCredential {
	p := new(OperationCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.OperationCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Current status of the operation.
*/
type OperationStatus int

const (
	OPERATIONSTATUS_UNKNOWN   OperationStatus = 0
	OPERATIONSTATUS_REDACTED  OperationStatus = 1
	OPERATIONSTATUS_SUCCEEDED OperationStatus = 2
	OPERATIONSTATUS_FAILED    OperationStatus = 3
	OPERATIONSTATUS_RUNNING   OperationStatus = 4
	OPERATIONSTATUS_SUSPENDED OperationStatus = 5
	OPERATIONSTATUS_CANCELED  OperationStatus = 6
	OPERATIONSTATUS_CANCELING OperationStatus = 7
	OPERATIONSTATUS_QUEUED    OperationStatus = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"RUNNING",
		"SUSPENDED",
		"CANCELED",
		"CANCELING",
		"QUEUED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"RUNNING",
		"SUSPENDED",
		"CANCELED",
		"CANCELING",
		"QUEUED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationStatus) index(name string) OperationStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"FAILED",
		"RUNNING",
		"SUSPENDED",
		"CANCELED",
		"CANCELING",
		"QUEUED",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationStatus(idx)
		}
	}
	return OPERATIONSTATUS_UNKNOWN
}

func (e *OperationStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationStatus) Ref() *OperationStatus {
	return &e
}

/*
Type of the operation tracked by the task.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN   OperationType = 0
	OPERATIONTYPE_REDACTED  OperationType = 1
	OPERATIONTYPE_INVENTORY OperationType = 2
	OPERATIONTYPE_PRECHECKS OperationType = 3
	OPERATIONTYPE_UPGRADE   OperationType = 4
	OPERATIONTYPE_NONE      OperationType = 5
	OPERATIONTYPE_UPLOAD    OperationType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
		"UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
		"UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationType) index(name string) OperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INVENTORY",
		"PRECHECKS",
		"UPGRADE",
		"NONE",
		"UPLOAD",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationType(idx)
		}
	}
	return OPERATIONTYPE_UNKNOWN
}

func (e *OperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationType) Ref() *OperationType {
	return &e
}

/*
Specification for running a Precheck operation.
*/
type PrechecksSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credentials []Credential `json:"credentials,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`

	ManagementServer *ManagementServer `json:"managementServer,omitempty"`
	/*
	  List of prechecks to skip. The allowed value is 'powerOffUvms' that skips the pinned VM prechecks.
	*/
	SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
}

func (p *PrechecksSpec) MarshalJSON() ([]byte, error) {
	type PrechecksSpecProxy PrechecksSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PrechecksSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		PrechecksSpecProxy: (*PrechecksSpecProxy)(p),
		EntityUpdateSpecs:  p.EntityUpdateSpecs,
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

func (p *PrechecksSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PrechecksSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPrechecksSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credentials != nil {
		p.Credentials = known.Credentials
	}
	if known.EntityUpdateSpecs != nil {
		p.EntityUpdateSpecs = known.EntityUpdateSpecs
	}
	if known.ManagementServer != nil {
		p.ManagementServer = known.ManagementServer
	}
	if known.SkippedPrecheckFlags != nil {
		p.SkippedPrecheckFlags = known.SkippedPrecheckFlags
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credentials")
	delete(allFields, "entityUpdateSpecs")
	delete(allFields, "managementServer")
	delete(allFields, "skippedPrecheckFlags")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPrechecksSpec() *PrechecksSpec {
	p := new(PrechecksSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.PrechecksSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A specification defining the entity being preloaded and its version.
*/
type PreloadSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`
}

func (p *PreloadSpec) MarshalJSON() ([]byte, error) {
	type PreloadSpecProxy PreloadSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PreloadSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		PreloadSpecProxy:  (*PreloadSpecProxy)(p),
		EntityUpdateSpecs: p.EntityUpdateSpecs,
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

func (p *PreloadSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PreloadSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPreloadSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityUpdateSpecs != nil {
		p.EntityUpdateSpecs = known.EntityUpdateSpecs
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityUpdateSpecs")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPreloadSpec() *PreloadSpec {
	p := new(PreloadSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.PreloadSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
System auto-management flag to handle system operation during upgrade.
*/
type SystemAutoMgmtFlag int

const (
	SYSTEMAUTOMGMTFLAG_UNKNOWN                  SystemAutoMgmtFlag = 0
	SYSTEMAUTOMGMTFLAG_REDACTED                 SystemAutoMgmtFlag = 1
	SYSTEMAUTOMGMTFLAG_POWER_OFF_UVMS           SystemAutoMgmtFlag = 2
	SYSTEMAUTOMGMTFLAG_MIGRATE_POWERED_OFF_UVMS SystemAutoMgmtFlag = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SystemAutoMgmtFlag) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
		"MIGRATE_POWERED_OFF_UVMS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SystemAutoMgmtFlag) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
		"MIGRATE_POWERED_OFF_UVMS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SystemAutoMgmtFlag) index(name string) SystemAutoMgmtFlag {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POWER_OFF_UVMS",
		"MIGRATE_POWERED_OFF_UVMS",
	}
	for idx := range names {
		if names[idx] == name {
			return SystemAutoMgmtFlag(idx)
		}
	}
	return SYSTEMAUTOMGMTFLAG_UNKNOWN
}

func (e *SystemAutoMgmtFlag) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SystemAutoMgmtFlag:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SystemAutoMgmtFlag) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SystemAutoMgmtFlag) Ref() *SystemAutoMgmtFlag {
	return &e
}

/*
Specification for an upgrade operation of an entity to a particular target version.
*/
type UpgradeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of automated system operations to perform, to avoid precheck failure and let the system restore state after an update is complete. The allowed flag is: - 'powerOffUvms': This allows the system to automatically power off user VMs which cannot be migrated to other hosts and power them on when the update is done. This option can avoid pinned VM precheck failure on the host which needs to enter maintenance mode during the update and allow the update to go through.
	*/
	AutoHandleFlags []SystemAutoMgmtFlag `json:"autoHandleFlags,omitempty"`

	Credentials []Credential `json:"credentials,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs"`

	ManagementServer *ManagementServer `json:"managementServer,omitempty"`
	/*
	  Number of seconds LCM waits for the VMs to come up after exiting host maintenance mode.
	*/
	MaxWaitTimeInSecs *int `json:"maxWaitTimeInSecs,omitempty"`
	/*
	  List of prechecks to skip. The allowed value is 'powerOffUvms' that skips the pinned VM prechecks.
	*/
	SkippedPrecheckFlags []SystemAutoMgmtFlag `json:"skippedPrecheckFlags,omitempty"`
}

func (p *UpgradeSpec) MarshalJSON() ([]byte, error) {
	type UpgradeSpecProxy UpgradeSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UpgradeSpecProxy
		EntityUpdateSpecs []EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	}{
		UpgradeSpecProxy:  (*UpgradeSpecProxy)(p),
		EntityUpdateSpecs: p.EntityUpdateSpecs,
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
	*p = *NewUpgradeSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AutoHandleFlags != nil {
		p.AutoHandleFlags = known.AutoHandleFlags
	}
	if known.Credentials != nil {
		p.Credentials = known.Credentials
	}
	if known.EntityUpdateSpecs != nil {
		p.EntityUpdateSpecs = known.EntityUpdateSpecs
	}
	if known.ManagementServer != nil {
		p.ManagementServer = known.ManagementServer
	}
	if known.MaxWaitTimeInSecs != nil {
		p.MaxWaitTimeInSecs = known.MaxWaitTimeInSecs
	}
	if known.SkippedPrecheckFlags != nil {
		p.SkippedPrecheckFlags = known.SkippedPrecheckFlags
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "autoHandleFlags")
	delete(allFields, "credentials")
	delete(allFields, "entityUpdateSpecs")
	delete(allFields, "managementServer")
	delete(allFields, "maxWaitTimeInSecs")
	delete(allFields, "skippedPrecheckFlags")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpgradeSpec() *UpgradeSpec {
	p := new(UpgradeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.UpgradeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
vendor management credentials for inventory operation
*/
type VendorManagementCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	CredentialSpecItemDiscriminator_ *string `json:"$credentialSpecItemDiscriminator,omitempty"`
	/*
	  Specification of credentials to be provided by the user to perform LCM operations.
	*/
	CredentialSpec *OneOfVendorManagementCredentialCredentialSpec `json:"credentialSpec"`
}

func (p *VendorManagementCredential) MarshalJSON() ([]byte, error) {
	type VendorManagementCredentialProxy VendorManagementCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VendorManagementCredentialProxy
		CredentialSpec *OneOfVendorManagementCredentialCredentialSpec `json:"credentialSpec,omitempty"`
	}{
		VendorManagementCredentialProxy: (*VendorManagementCredentialProxy)(p),
		CredentialSpec:                  p.CredentialSpec,
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

func (p *VendorManagementCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VendorManagementCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVendorManagementCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CredentialSpecItemDiscriminator_ != nil {
		p.CredentialSpecItemDiscriminator_ = known.CredentialSpecItemDiscriminator_
	}
	if known.CredentialSpec != nil {
		p.CredentialSpec = known.CredentialSpec
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$credentialSpecItemDiscriminator")
	delete(allFields, "credentialSpec")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVendorManagementCredential() *VendorManagementCredential {
	p := new(VendorManagementCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.common.VendorManagementCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VendorManagementCredential) GetCredentialSpec() interface{} {
	if nil == p.CredentialSpec {
		return nil
	}
	return p.CredentialSpec.GetValue()
}

func (p *VendorManagementCredential) SetCredentialSpec(v interface{}) error {
	if nil == p.CredentialSpec {
		p.CredentialSpec = NewOneOfVendorManagementCredentialCredentialSpec()
	}
	e := p.CredentialSpec.SetValue(v)
	if nil == e {
		if nil == p.CredentialSpecItemDiscriminator_ {
			p.CredentialSpecItemDiscriminator_ = new(string)
		}
		*p.CredentialSpecItemDiscriminator_ = *p.CredentialSpec.Discriminator
	}
	return e
}

/*
Name of the vendor management software that manages fleet of servers. This could be one of Intersight or Vcenter or UCS.
*/
type VendorManagementName int

const (
	VENDORMANAGEMENTNAME_UNKNOWN  VendorManagementName = 0
	VENDORMANAGEMENTNAME_REDACTED VendorManagementName = 1
	VENDORMANAGEMENTNAME_UCS      VendorManagementName = 2
	VENDORMANAGEMENTNAME_ISM      VendorManagementName = 3
	VENDORMANAGEMENTNAME_VCENTER  VendorManagementName = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VendorManagementName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UCS",
		"ISM",
		"VCENTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VendorManagementName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UCS",
		"ISM",
		"VCENTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VendorManagementName) index(name string) VendorManagementName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UCS",
		"ISM",
		"VCENTER",
	}
	for idx := range names {
		if names[idx] == name {
			return VendorManagementName(idx)
		}
	}
	return VENDORMANAGEMENTNAME_UNKNOWN
}

func (e *VendorManagementName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VendorManagementName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VendorManagementName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VendorManagementName) Ref() *VendorManagementName {
	return &e
}

type OneOfVendorManagementCredentialCredentialSpec struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType2002 *import2.VcenterCredential    `json:"-"`
	oneOfType2001 *import2.IntersightCredential `json:"-"`
}

func NewOneOfVendorManagementCredentialCredentialSpec() *OneOfVendorManagementCredentialCredentialSpec {
	p := new(OneOfVendorManagementCredentialCredentialSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVendorManagementCredentialCredentialSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVendorManagementCredentialCredentialSpec is nil"))
	}
	switch v.(type) {
	case import2.VcenterCredential:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(import2.VcenterCredential)
		}
		*p.oneOfType2002 = v.(import2.VcenterCredential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case import2.IntersightCredential:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.IntersightCredential)
		}
		*p.oneOfType2001 = v.(import2.IntersightCredential)
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

func (p *OneOfVendorManagementCredentialCredentialSpec) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfVendorManagementCredentialCredentialSpec) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(import2.VcenterCredential)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "security.v4.config.VcenterCredential" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(import2.VcenterCredential)
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
	vOneOfType2001 := new(import2.IntersightCredential)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "security.v4.config.IntersightCredential" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.IntersightCredential)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVendorManagementCredentialCredentialSpec"))
}

func (p *OneOfVendorManagementCredentialCredentialSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfVendorManagementCredentialCredentialSpec")
}

type OneOfCredentialCredentialDetail struct {
	Discriminator *string                     `json:"-"`
	ObjectType_   *string                     `json:"-"`
	oneOfType2001 *CredentialReference        `json:"-"`
	oneOfType2002 *VendorManagementCredential `json:"-"`
}

func NewOneOfCredentialCredentialDetail() *OneOfCredentialCredentialDetail {
	p := new(OneOfCredentialCredentialDetail)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCredentialCredentialDetail) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCredentialCredentialDetail is nil"))
	}
	switch v.(type) {
	case CredentialReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(CredentialReference)
		}
		*p.oneOfType2001 = v.(CredentialReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case VendorManagementCredential:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(VendorManagementCredential)
		}
		*p.oneOfType2002 = v.(VendorManagementCredential)
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

func (p *OneOfCredentialCredentialDetail) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfCredentialCredentialDetail) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(CredentialReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.common.CredentialReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(CredentialReference)
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
	vOneOfType2002 := new(VendorManagementCredential)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "lifecycle.v4.common.VendorManagementCredential" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(VendorManagementCredential)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCredentialCredentialDetail"))
}

func (p *OneOfCredentialCredentialDetail) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfCredentialCredentialDetail")
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
