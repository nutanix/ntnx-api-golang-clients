/*
 * Generated file models/prism/v4/mgmt/mgmt_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.mgmt of Nutanix Prism APIs
*/
package mgmt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
)

type EnvironmentInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  This field signifies billing API hostname which serves billing apis.
	*/
	BillingApiHost *string `json:"billingApiHost,omitempty"`
	/*
	  This field signifies the billing UI hostname.
	*/
	BillingHost *string `json:"billingHost,omitempty"`
	/*
	  This signifies region/az info. There can be different billing units for different region.
	CFS service running in PC and PE (in XI DC) currently pass this info along with metrix and spec
	to metering and telemetry pipeline.
	*/
	CellFqdn *string `json:"cellFqdn,omitempty"`
	/*
	  This field signifies tenant region, to be selected by customers on MCM UI.
	*/
	CloudRegionName *string `json:"cloudRegionName,omitempty"`
	/*
	  This field signifies cloud site name for a given company. A company can have multiple povisioned sites.
	This will be provided by customers on MCM UI.
	*/
	CloudSiteName *string `json:"cloudSiteName,omitempty"`

	EnvironmentType *EnvironmentType `json:"environmentType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	InstanceType *InstanceType `json:"instanceType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  This field signifies local availability zone name derived from tenant region.
	*/
	LocalAzName *string `json:"localAzName,omitempty"`
	/*
	  This field signifies mynutanix URL.
	*/
	MyNutanixUrl *string `json:"myNutanixUrl,omitempty"`

	OlbVirtualAddress *LbAddress `json:"olbVirtualAddress,omitempty"`
	/*
	  This field signifies DNS mapped fully qualified domain name provided by MCM.
	*/
	PcExternalUrl *string `json:"pcExternalUrl,omitempty"`

	ProviderType *ProviderType `json:"providerType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The unique UUID provisioned to the tenant.
	*/
	TenantUuid *string `json:"tenantUuid,omitempty"`

	XlbVirtualAddress *LbAddress `json:"xlbVirtualAddress,omitempty"`
}

func (p *EnvironmentInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EnvironmentInfo

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

func (p *EnvironmentInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EnvironmentInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEnvironmentInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BillingApiHost != nil {
		p.BillingApiHost = known.BillingApiHost
	}
	if known.BillingHost != nil {
		p.BillingHost = known.BillingHost
	}
	if known.CellFqdn != nil {
		p.CellFqdn = known.CellFqdn
	}
	if known.CloudRegionName != nil {
		p.CloudRegionName = known.CloudRegionName
	}
	if known.CloudSiteName != nil {
		p.CloudSiteName = known.CloudSiteName
	}
	if known.EnvironmentType != nil {
		p.EnvironmentType = known.EnvironmentType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.InstanceType != nil {
		p.InstanceType = known.InstanceType
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocalAzName != nil {
		p.LocalAzName = known.LocalAzName
	}
	if known.MyNutanixUrl != nil {
		p.MyNutanixUrl = known.MyNutanixUrl
	}
	if known.OlbVirtualAddress != nil {
		p.OlbVirtualAddress = known.OlbVirtualAddress
	}
	if known.PcExternalUrl != nil {
		p.PcExternalUrl = known.PcExternalUrl
	}
	if known.ProviderType != nil {
		p.ProviderType = known.ProviderType
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TenantUuid != nil {
		p.TenantUuid = known.TenantUuid
	}
	if known.XlbVirtualAddress != nil {
		p.XlbVirtualAddress = known.XlbVirtualAddress
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "billingApiHost")
	delete(allFields, "billingHost")
	delete(allFields, "cellFqdn")
	delete(allFields, "cloudRegionName")
	delete(allFields, "cloudSiteName")
	delete(allFields, "environmentType")
	delete(allFields, "extId")
	delete(allFields, "instanceType")
	delete(allFields, "links")
	delete(allFields, "localAzName")
	delete(allFields, "myNutanixUrl")
	delete(allFields, "olbVirtualAddress")
	delete(allFields, "pcExternalUrl")
	delete(allFields, "providerType")
	delete(allFields, "tenantId")
	delete(allFields, "tenantUuid")
	delete(allFields, "xlbVirtualAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEnvironmentInfo() *EnvironmentInfo {
	p := new(EnvironmentInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.mgmt.EnvironmentInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An enum denoting the environment type of the PC. Whether it is an OnPrem PC or a cloud PC.<br>
Following are the supported entity types:
- ONPREM
- NTNX_CLOUD
*/
type EnvironmentType int

const (
	ENVIRONMENTTYPE_UNKNOWN    EnvironmentType = 0
	ENVIRONMENTTYPE_REDACTED   EnvironmentType = 1
	ENVIRONMENTTYPE_ONPREM     EnvironmentType = 2
	ENVIRONMENTTYPE_NTNX_CLOUD EnvironmentType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnvironmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EnvironmentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EnvironmentType) index(name string) EnvironmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	for idx := range names {
		if names[idx] == name {
			return EnvironmentType(idx)
		}
	}
	return ENVIRONMENTTYPE_UNKNOWN
}

func (e *EnvironmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnvironmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnvironmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnvironmentType) Ref() *EnvironmentType {
	return &e
}

/*
An enum denoting the instance type of the cloud PC. Indicates whether PC is created on baremetal
or on a cloud provisioned VM.. Hence it supports two possible values:
- NTNX_PROVISIONED
- NATIVE_PROVISIONED
*/
type InstanceType int

const (
	INSTANCETYPE_UNKNOWN            InstanceType = 0
	INSTANCETYPE_REDACTED           InstanceType = 1
	INSTANCETYPE_NTNX_PROVISIONED   InstanceType = 2
	INSTANCETYPE_NATIVE_PROVISIONED InstanceType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *InstanceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX_PROVISIONED",
		"NATIVE_PROVISIONED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e InstanceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX_PROVISIONED",
		"NATIVE_PROVISIONED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *InstanceType) index(name string) InstanceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX_PROVISIONED",
		"NATIVE_PROVISIONED",
	}
	for idx := range names {
		if names[idx] == name {
			return InstanceType(idx)
		}
	}
	return INSTANCETYPE_UNKNOWN
}

func (e *InstanceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for InstanceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *InstanceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e InstanceType) Ref() *InstanceType {
	return &e
}

type LbAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	Fqdn *string `json:"fqdn,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Ipv6 *string `json:"ipv6,omitempty"`

	IsBackup *bool `json:"isBackup,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Port *int `json:"port,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *LbAddress) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LbAddress

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

func (p *LbAddress) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LbAddress
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLbAddress()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.Ipv6 != nil {
		p.Ipv6 = known.Ipv6
	}
	if known.IsBackup != nil {
		p.IsBackup = known.IsBackup
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "fqdn")
	delete(allFields, "ip")
	delete(allFields, "ipv6")
	delete(allFields, "isBackup")
	delete(allFields, "links")
	delete(allFields, "port")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLbAddress() *LbAddress {
	p := new(LbAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.mgmt.LbAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An enum denoting the provider of the cloud in case the environment type is a cloud PC.<br>
The following providers are currently supported by the service:
- NTNX
- AZURE
- AWS
- GCP
*/
type ProviderType int

const (
	PROVIDERTYPE_UNKNOWN  ProviderType = 0
	PROVIDERTYPE_REDACTED ProviderType = 1
	PROVIDERTYPE_NTNX     ProviderType = 2
	PROVIDERTYPE_AZURE    ProviderType = 3
	PROVIDERTYPE_AWS      ProviderType = 4
	PROVIDERTYPE_GCP      ProviderType = 5
	PROVIDERTYPE_VSPHERE  ProviderType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProviderType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProviderType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProviderType) index(name string) ProviderType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	for idx := range names {
		if names[idx] == name {
			return ProviderType(idx)
		}
	}
	return PROVIDERTYPE_UNKNOWN
}

func (e *ProviderType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProviderType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProviderType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProviderType) Ref() *ProviderType {
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
