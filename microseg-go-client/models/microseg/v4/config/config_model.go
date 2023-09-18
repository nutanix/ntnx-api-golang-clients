/*
 * Generated file models/microseg/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Microseg Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure security policies, address groups and service groups
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/prism/v4/config"
	"time"
)

type AddressGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A user defined annotation for a address group.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of IP range containing start and end IP.
	*/
	IpRanges []IPv4Range `json:"ipRanges,omitempty"`
	/*
	  List of CIDR blocks in the address group.
	*/
	Ipv4Addresses []import1.IPv4Address `json:"ipv4Addresses,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A short identifier for a address group.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AddressGroup) MarshalJSON() ([]byte, error) {
	type AddressGroupProxy AddressGroup
	return json.Marshal(struct {
		*AddressGroupProxy
		Name *string `json:"name,omitempty"`
	}{
		AddressGroupProxy: (*AddressGroupProxy)(p),
		Name:              p.Name,
	})
}

func NewAddressGroup() *AddressGroup {
	p := new(AddressGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/address-groups/{extId} Get operation
*/
type AddressGroupGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddressGroupGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddressGroupGetResponse() *AddressGroupGetResponse {
	p := new(AddressGroupGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddressGroupGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddressGroupGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddressGroupGetResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/address-groups Get operation
*/
type AddressGroupListResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddressGroupListResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddressGroupListResponse() *AddressGroupListResponse {
	p := new(AddressGroupListResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupListResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupListResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddressGroupListResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddressGroupListResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddressGroupListResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/address-groups/$actions/build-policy-association Post operation
*/
type AddressGroupPolicyAssociationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddressGroupPolicyAssociationResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddressGroupPolicyAssociationResponse() *AddressGroupPolicyAssociationResponse {
	p := new(AddressGroupPolicyAssociationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupPolicyAssociationResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupPolicyAssociationResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddressGroupPolicyAssociationResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddressGroupPolicyAssociationResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddressGroupPolicyAssociationResponseData()
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
Reference to policy associated with address group.
*/
type AddressGroupPolicyReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressGroupUuid *string `json:"addressGroupUuid,omitempty"`

	PolicyUuids []string `json:"policyUuids,omitempty"`
}

func NewAddressGroupPolicyReference() *AddressGroupPolicyReference {
	p := new(AddressGroupPolicyReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupPolicyReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupPolicyReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of address group UUID
*/
type AddressGroupReferenceSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressGroupUuidList []string `json:"addressGroupUuidList"`
}

func (p *AddressGroupReferenceSpec) MarshalJSON() ([]byte, error) {
	type AddressGroupReferenceSpecProxy AddressGroupReferenceSpec
	return json.Marshal(struct {
		*AddressGroupReferenceSpecProxy
		AddressGroupUuidList []string `json:"addressGroupUuidList,omitempty"`
	}{
		AddressGroupReferenceSpecProxy: (*AddressGroupReferenceSpecProxy)(p),
		AddressGroupUuidList:           p.AddressGroupUuidList,
	})
}

func NewAddressGroupReferenceSpec() *AddressGroupReferenceSpec {
	p := new(AddressGroupReferenceSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupReferenceSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupReferenceSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/address-groups/{extId} Delete operation
*/
type AddressGroupTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddressGroupTaskResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddressGroupTaskResponse() *AddressGroupTaskResponse {
	p := new(AddressGroupTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AddressGroupTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.AddressGroupTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddressGroupTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddressGroupTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddressGroupTaskResponseData()
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

type AllowType int

const (
	ALLOWTYPE_UNKNOWN  AllowType = 0
	ALLOWTYPE_REDACTED AllowType = 1
	ALLOWTYPE_ALL      AllowType = 2
	ALLOWTYPE_NONE     AllowType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AllowType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AllowType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AllowType) index(name string) AllowType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return AllowType(idx)
		}
	}
	return ALLOWTYPE_UNKNOWN
}

func (e *AllowType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AllowType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AllowType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AllowType) Ref() *AllowType {
	return &e
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/cadmus/banners Get operation
*/
type BannerGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfBannerGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewBannerGetResponse() *BannerGetResponse {
	p := new(BannerGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.BannerGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.BannerGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *BannerGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *BannerGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfBannerGetResponseData()
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

type BannerResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BannerType *BannerType `json:"bannerType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Memory associated with the current mode of operation for cadmus.
	*/
	TierMemoryInGB *int `json:"tierMemoryInGB,omitempty"`
}

func NewBannerResponse() *BannerResponse {
	p := new(BannerResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.BannerResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.BannerResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Shows whether the banner is for a warning(shown in yellow) or for an error(shown in red).
*/
type BannerType int

const (
	BANNERTYPE_UNKNOWN  BannerType = 0
	BANNERTYPE_REDACTED BannerType = 1
	BANNERTYPE_WARNING  BannerType = 2
	BANNERTYPE_ERROR    BannerType = 3
	BANNERTYPE_INFO     BannerType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BannerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BannerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BannerType) index(name string) BannerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
		"INFO",
	}
	for idx := range names {
		if names[idx] == name {
			return BannerType(idx)
		}
	}
	return BANNERTYPE_UNKNOWN
}

func (e *BannerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BannerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BannerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BannerType) Ref() *BannerType {
	return &e
}

type CategoryMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of category mapping.
	*/
	Name *string `json:"name"`

	Resources *CategoryMappingInfo `json:"resources"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CategoryMapping) MarshalJSON() ([]byte, error) {
	type CategoryMappingProxy CategoryMapping
	return json.Marshal(struct {
		*CategoryMappingProxy
		Name      *string              `json:"name,omitempty"`
		Resources *CategoryMappingInfo `json:"resources,omitempty"`
	}{
		CategoryMappingProxy: (*CategoryMappingProxy)(p),
		Name:                 p.Name,
		Resources:            p.Resources,
	})
}

func NewCategoryMapping() *CategoryMapping {
	p := new(CategoryMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CategoryMapping"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.CategoryMapping"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A mapping to an object in Active Directory.
*/
type CategoryMappingAdInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The extId of the directory service that will be used for mapping.
	*/
	DirectoryServiceExtId *string `json:"directoryServiceExtId,omitempty"`
	/*
	  The name of the directory service that will be used for mapping.
	*/
	DirectoryServiceName *string `json:"directoryServiceName,omitempty"`
	/*
	  The objectGUID for the object in AD.
	*/
	ObjectIdentifier *string `json:"objectIdentifier"`
	/*
	  The path for the mapped object in AD.
	*/
	ObjectPath *string `json:"objectPath,omitempty"`

	Status *CategoryMappingAdStatus `json:"status,omitempty"`
}

func (p *CategoryMappingAdInfo) MarshalJSON() ([]byte, error) {
	type CategoryMappingAdInfoProxy CategoryMappingAdInfo
	return json.Marshal(struct {
		*CategoryMappingAdInfoProxy
		ObjectIdentifier *string `json:"objectIdentifier,omitempty"`
	}{
		CategoryMappingAdInfoProxy: (*CategoryMappingAdInfoProxy)(p),
		ObjectIdentifier:           p.ObjectIdentifier,
	})
}

func NewCategoryMappingAdInfo() *CategoryMappingAdInfo {
	p := new(CategoryMappingAdInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CategoryMappingAdInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.CategoryMappingAdInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The mapping status of AD Mapping.
*/
type CategoryMappingAdStatus int

const (
	CATEGORYMAPPINGADSTATUS_UNKNOWN                  CategoryMappingAdStatus = 0
	CATEGORYMAPPINGADSTATUS_REDACTED                 CategoryMappingAdStatus = 1
	CATEGORYMAPPINGADSTATUS_USABLE                   CategoryMappingAdStatus = 2
	CATEGORYMAPPINGADSTATUS_DELETED                  CategoryMappingAdStatus = 3
	CATEGORYMAPPINGADSTATUS_DIRECTORY_NOT_CONFIGURED CategoryMappingAdStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CategoryMappingAdStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USABLE",
		"DELETED",
		"DIRECTORY_NOT_CONFIGURED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CategoryMappingAdStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USABLE",
		"DELETED",
		"DIRECTORY_NOT_CONFIGURED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CategoryMappingAdStatus) index(name string) CategoryMappingAdStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USABLE",
		"DELETED",
		"DIRECTORY_NOT_CONFIGURED",
	}
	for idx := range names {
		if names[idx] == name {
			return CategoryMappingAdStatus(idx)
		}
	}
	return CATEGORYMAPPINGADSTATUS_UNKNOWN
}

func (e *CategoryMappingAdStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CategoryMappingAdStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CategoryMappingAdStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CategoryMappingAdStatus) Ref() *CategoryMappingAdStatus {
	return &e
}

/*
A mapping from an object to a category value.
*/
type CategoryMappingInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AdMapping *CategoryMappingAdInfo `json:"adMapping"`
	/*
	  The name for the category that this mapping is for.
	*/
	CategoryName *string `json:"categoryName"`
	/*
	  The value for the category that this mapping is for.
	*/
	CategoryValue *string `json:"categoryValue"`
}

func (p *CategoryMappingInfo) MarshalJSON() ([]byte, error) {
	type CategoryMappingInfoProxy CategoryMappingInfo
	return json.Marshal(struct {
		*CategoryMappingInfoProxy
		AdMapping     *CategoryMappingAdInfo `json:"adMapping,omitempty"`
		CategoryName  *string                `json:"categoryName,omitempty"`
		CategoryValue *string                `json:"categoryValue,omitempty"`
	}{
		CategoryMappingInfoProxy: (*CategoryMappingInfoProxy)(p),
		AdMapping:                p.AdMapping,
		CategoryName:             p.CategoryName,
		CategoryValue:            p.CategoryValue,
	})
}

func NewCategoryMappingInfo() *CategoryMappingInfo {
	p := new(CategoryMappingInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CategoryMappingInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.CategoryMappingInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Policy-wide options for a security policy.
*/
type ConfigMigrationPolicyOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IsHitlogEnabled *bool `json:"isHitlogEnabled,omitempty"`

	IsIpv6TrafficAllowed *bool `json:"isIpv6TrafficAllowed,omitempty"`
}

func NewConfigMigrationPolicyOptions() *ConfigMigrationPolicyOptions {
	p := new(ConfigMigrationPolicyOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ConfigMigrationPolicyOptions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ConfigMigrationPolicyOptions"}
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
type ConfigMigrationPolicyType int

const (
	CONFIGMIGRATIONPOLICYTYPE_UNKNOWN     ConfigMigrationPolicyType = 0
	CONFIGMIGRATIONPOLICYTYPE_REDACTED    ConfigMigrationPolicyType = 1
	CONFIGMIGRATIONPOLICYTYPE_QUARANTINE  ConfigMigrationPolicyType = 2
	CONFIGMIGRATIONPOLICYTYPE_ISOLATION   ConfigMigrationPolicyType = 3
	CONFIGMIGRATIONPOLICYTYPE_APPLICATION ConfigMigrationPolicyType = 4
	CONFIGMIGRATIONPOLICYTYPE_AD          ConfigMigrationPolicyType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConfigMigrationPolicyType) name(index int) string {
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
func (e ConfigMigrationPolicyType) GetName() string {
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
func (e *ConfigMigrationPolicyType) index(name string) ConfigMigrationPolicyType {
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
			return ConfigMigrationPolicyType(idx)
		}
	}
	return CONFIGMIGRATIONPOLICYTYPE_UNKNOWN
}

func (e *ConfigMigrationPolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConfigMigrationPolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConfigMigrationPolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConfigMigrationPolicyType) Ref() *ConfigMigrationPolicyType {
	return &e
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/flow-migrator/previews/{extId} Get operation
*/
type ConfigMigrationPreviewGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConfigMigrationPreviewGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConfigMigrationPreviewGetResponse() *ConfigMigrationPreviewGetResponse {
	p := new(ConfigMigrationPreviewGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ConfigMigrationPreviewGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ConfigMigrationPreviewGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigMigrationPreviewGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConfigMigrationPreviewGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConfigMigrationPreviewGetResponseData()
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

type ConfigMigrationPreviewResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	NewPolicySpec *NetworkSecurityPolicy `json:"newPolicySpec,omitempty"`
	/*
	  This field corresponds to a list of policy spec which captures the system generated FNS 2.0 policies
	as a result of a single FNS 1.0 policy migration.
	*/
	SystemGeneratedPolicies []NetworkSecurityPolicy `json:"systemGeneratedPolicies,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewConfigMigrationPreviewResponse() *ConfigMigrationPreviewResponse {
	p := new(ConfigMigrationPreviewResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ConfigMigrationPreviewResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ConfigMigrationPreviewResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConfigMigrationSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CountSummary *NetworkSecurityPolicyMigrationCountSummary `json:"countSummary,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	PolicySummaries []NetworkSecurityPolicyMigrationSummary `json:"policySummaries,omitempty"`

	SubnetSummaries []NetworkSecurityPolicyMigrationSubnetSummary `json:"subnetSummaries,omitempty"`

	SummaryFailures []NetworkSecurityPolicyMigrationSummaryFailures `json:"summaryFailures,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewConfigMigrationSummary() *ConfigMigrationSummary {
	p := new(ConfigMigrationSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ConfigMigrationSummary"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ConfigMigrationSummary"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/flow-migrator/summaries/{extId} Get operation
*/
type ConfigMigrationSummaryGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConfigMigrationSummaryGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConfigMigrationSummaryGetResponse() *ConfigMigrationSummaryGetResponse {
	p := new(ConfigMigrationSummaryGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ConfigMigrationSummaryGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ConfigMigrationSummaryGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConfigMigrationSummaryGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConfigMigrationSummaryGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConfigMigrationSummaryGetResponseData()
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

type DirectoryServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The extId of the directory service that will be used for mapping.
	*/
	DirectoryServiceExtId *string `json:"directoryServiceExtId,omitempty"`
	/*
	  The name of the directory service that will be used for mapping.
	*/
	DirectoryServiceName *string `json:"directoryServiceName,omitempty"`
	/*
	  List of domain controllers to be used for event scraping.
	*/
	DomainControllers []DomainController `json:"domainControllers,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Enablement status of the default category.
	*/
	IsDefaultCategoryEnabled *bool `json:"isDefaultCategoryEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The matching criteria used to determine whether an entity will be categorized by identity categorization. If match type is ALL, all the entities will be categorized.
	*/
	MatchingCriterias []MatchingCriteria `json:"matchingCriterias,omitempty"`
	/*
	  Retain default category on user login.
	*/
	ShouldKeepDefaultCategoryOnLogin *bool `json:"shouldKeepDefaultCategoryOnLogin,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDirectoryServer() *DirectoryServer {
	p := new(DirectoryServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServer"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServer"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefaultCategoryEnabled = new(bool)
	*p.IsDefaultCategoryEnabled = false
	p.ShouldKeepDefaultCategoryOnLogin = new(bool)
	*p.ShouldKeepDefaultCategoryOnLogin = false

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/directory-servers Post operation
*/
type DirectoryServerCreateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDirectoryServerCreateResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectoryServerCreateResponse() *DirectoryServerCreateResponse {
	p := new(DirectoryServerCreateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerCreateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServerCreateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DirectoryServerCreateResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DirectoryServerCreateResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDirectoryServerCreateResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/directory-servers/{extId} Delete operation
*/
type DirectoryServerDeleteResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDirectoryServerDeleteResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectoryServerDeleteResponse() *DirectoryServerDeleteResponse {
	p := new(DirectoryServerDeleteResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerDeleteResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServerDeleteResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DirectoryServerDeleteResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DirectoryServerDeleteResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDirectoryServerDeleteResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/directory-servers/{extId} Get operation
*/
type DirectoryServerGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDirectoryServerGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectoryServerGetResponse() *DirectoryServerGetResponse {
	p := new(DirectoryServerGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServerGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DirectoryServerGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DirectoryServerGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDirectoryServerGetResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/directory-servers Get operation
*/
type DirectoryServerListResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDirectoryServerListResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectoryServerListResponse() *DirectoryServerListResponse {
	p := new(DirectoryServerListResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerListResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServerListResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DirectoryServerListResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DirectoryServerListResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDirectoryServerListResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/directory-servers/{extId} Put operation
*/
type DirectoryServerUpdateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDirectoryServerUpdateResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDirectoryServerUpdateResponse() *DirectoryServerUpdateResponse {
	p := new(DirectoryServerUpdateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerUpdateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DirectoryServerUpdateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DirectoryServerUpdateResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DirectoryServerUpdateResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDirectoryServerUpdateResponseData()
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
Configuration for an AD domain controller.
*/
type DomainController struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  IPv4 Address or hostname for the domain controller.
	*/
	Host *string `json:"host"`
}

func (p *DomainController) MarshalJSON() ([]byte, error) {
	type DomainControllerProxy DomainController
	return json.Marshal(struct {
		*DomainControllerProxy
		Host *string `json:"host,omitempty"`
	}{
		DomainControllerProxy: (*DomainControllerProxy)(p),
		Host:                  p.Host,
	})
}

func NewDomainController() *DomainController {
	p := new(DomainController)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DomainController"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DomainController"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/category-mappings Post operation
*/
type DsCategoryMappingCreateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDsCategoryMappingCreateResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDsCategoryMappingCreateResponse() *DsCategoryMappingCreateResponse {
	p := new(DsCategoryMappingCreateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DsCategoryMappingCreateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DsCategoryMappingCreateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DsCategoryMappingCreateResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DsCategoryMappingCreateResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDsCategoryMappingCreateResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/category-mappings/{extId} Delete operation
*/
type DsCategoryMappingDeleteResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDsCategoryMappingDeleteResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDsCategoryMappingDeleteResponse() *DsCategoryMappingDeleteResponse {
	p := new(DsCategoryMappingDeleteResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DsCategoryMappingDeleteResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DsCategoryMappingDeleteResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DsCategoryMappingDeleteResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DsCategoryMappingDeleteResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDsCategoryMappingDeleteResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/category-mappings/{extId} Get operation
*/
type DsCategoryMappingGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDsCategoryMappingGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDsCategoryMappingGetResponse() *DsCategoryMappingGetResponse {
	p := new(DsCategoryMappingGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DsCategoryMappingGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DsCategoryMappingGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DsCategoryMappingGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DsCategoryMappingGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDsCategoryMappingGetResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/category-mappings/{extId} Put operation
*/
type DsCategoryMappingUpdateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDsCategoryMappingUpdateResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDsCategoryMappingUpdateResponse() *DsCategoryMappingUpdateResponse {
	p := new(DsCategoryMappingUpdateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DsCategoryMappingUpdateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DsCategoryMappingUpdateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DsCategoryMappingUpdateResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DsCategoryMappingUpdateResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDsCategoryMappingUpdateResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/category-mappings Get operation
*/
type DsCategoryMappingsGetListResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDsCategoryMappingsGetListResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDsCategoryMappingsGetListResponse() *DsCategoryMappingsGetListResponse {
	p := new(DsCategoryMappingsGetListResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DsCategoryMappingsGetListResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.DsCategoryMappingsGetListResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DsCategoryMappingsGetListResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DsCategoryMappingsGetListResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDsCategoryMappingsGetListResponseData()
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

type FileWrapper struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	File *string `json:"file,omitempty"`
}

func NewFileWrapper() *FileWrapper {
	p := new(FileWrapper)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.FileWrapper"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.FileWrapper"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP range containing start and end IP.
*/
type IPv4Range struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *string `json:"endIp"`

	StartIp *string `json:"startIp"`
}

func (p *IPv4Range) MarshalJSON() ([]byte, error) {
	type IPv4RangeProxy IPv4Range
	return json.Marshal(struct {
		*IPv4RangeProxy
		EndIp   *string `json:"endIp,omitempty"`
		StartIp *string `json:"startIp,omitempty"`
	}{
		IPv4RangeProxy: (*IPv4RangeProxy)(p),
		EndIp:          p.EndIp,
		StartIp:        p.StartIp,
	})
}

func NewIPv4Range() *IPv4Range {
	p := new(IPv4Range)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.IPv4Range"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.IPv4Range"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Icmp Type Code Object
*/
type IcmpTypeCodeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Ignore this field if Code has to be ANY.
	*/
	Code *int `json:"code,omitempty"`
	/*
	  Set this field to true if both Type and Code is ANY.
	*/
	IsAllCodeTypeAllowed *bool `json:"isAllCodeTypeAllowed,omitempty"`
	/*
	  Ignore this field if Type has to be ANY.
	*/
	Type *int `json:"type,omitempty"`
}

func NewIcmpTypeCodeSpec() *IcmpTypeCodeSpec {
	p := new(IcmpTypeCodeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.IcmpTypeCodeSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.IcmpTypeCodeSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAllCodeTypeAllowed = new(bool)
	*p.IsAllCodeTypeAllowed = false

	return p
}

type IntraEntityGroupRuleAction int

const (
	INTRAENTITYGROUPRULEACTION_UNKNOWN  IntraEntityGroupRuleAction = 0
	INTRAENTITYGROUPRULEACTION_REDACTED IntraEntityGroupRuleAction = 1
	INTRAENTITYGROUPRULEACTION_ALLOW    IntraEntityGroupRuleAction = 2
	INTRAENTITYGROUPRULEACTION_DENY     IntraEntityGroupRuleAction = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IntraEntityGroupRuleAction) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e IntraEntityGroupRuleAction) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *IntraEntityGroupRuleAction) index(name string) IntraEntityGroupRuleAction {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DENY",
	}
	for idx := range names {
		if names[idx] == name {
			return IntraEntityGroupRuleAction(idx)
		}
	}
	return INTRAENTITYGROUPRULEACTION_UNKNOWN
}

func (e *IntraEntityGroupRuleAction) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IntraEntityGroupRuleAction:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IntraEntityGroupRuleAction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IntraEntityGroupRuleAction) Ref() *IntraEntityGroupRuleAction {
	return &e
}

/*
The entity to perform matching on. Currently, only the target VM that a logon occurred on is supported.
*/
type MatchEntity int

const (
	MATCHENTITY_UNKNOWN  MatchEntity = 0
	MATCHENTITY_REDACTED MatchEntity = 1
	MATCHENTITY_VM       MatchEntity = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MatchEntity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MatchEntity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MatchEntity) index(name string) MatchEntity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
	}
	for idx := range names {
		if names[idx] == name {
			return MatchEntity(idx)
		}
	}
	return MATCHENTITY_UNKNOWN
}

func (e *MatchEntity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MatchEntity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MatchEntity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MatchEntity) Ref() *MatchEntity {
	return &e
}

/*
The field to match on. Today only NAME is supported, which matches on an entity's name.
*/
type MatchField int

const (
	MATCHFIELD_UNKNOWN  MatchField = 0
	MATCHFIELD_REDACTED MatchField = 1
	MATCHFIELD_NAME     MatchField = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MatchField) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MatchField) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MatchField) index(name string) MatchField {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
	}
	for idx := range names {
		if names[idx] == name {
			return MatchField(idx)
		}
	}
	return MATCHFIELD_UNKNOWN
}

func (e *MatchField) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MatchField:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MatchField) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MatchField) Ref() *MatchField {
	return &e
}

/*
The type of match. Today only CONTAINS and ALL are supported. CONTAINS performs a substring match on the given entity and field for the criteria value whereas ALL allows all string to match on the given entity.
*/
type MatchType int

const (
	MATCHTYPE_UNKNOWN  MatchType = 0
	MATCHTYPE_REDACTED MatchType = 1
	MATCHTYPE_CONTAINS MatchType = 2
	MATCHTYPE_ALL      MatchType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MatchType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONTAINS",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MatchType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONTAINS",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MatchType) index(name string) MatchType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONTAINS",
		"ALL",
	}
	for idx := range names {
		if names[idx] == name {
			return MatchType(idx)
		}
	}
	return MATCHTYPE_UNKNOWN
}

func (e *MatchType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MatchType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MatchType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MatchType) Ref() *MatchType {
	return &e
}

/*
The matching criteria used to determine whether an entity will be categorized by identity categorization. If match type is ALL, all the entities will be categorized.
*/
type MatchingCriteria struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The criteria to use for matching entities to be categorized. Note that depending on the match type, the usage of this value may differ.
	*/
	Criteria *string `json:"criteria,omitempty"`

	MatchEntity *MatchEntity `json:"matchEntity"`

	MatchField *MatchField `json:"matchField"`

	MatchType *MatchType `json:"matchType"`
}

func (p *MatchingCriteria) MarshalJSON() ([]byte, error) {
	type MatchingCriteriaProxy MatchingCriteria
	return json.Marshal(struct {
		*MatchingCriteriaProxy
		MatchEntity *MatchEntity `json:"matchEntity,omitempty"`
		MatchField  *MatchField  `json:"matchField,omitempty"`
		MatchType   *MatchType   `json:"matchType,omitempty"`
	}{
		MatchingCriteriaProxy: (*MatchingCriteriaProxy)(p),
		MatchEntity:           p.MatchEntity,
		MatchField:            p.MatchField,
		MatchType:             p.MatchType,
	})
}

func NewMatchingCriteria() *MatchingCriteria {
	p := new(MatchingCriteria)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.MatchingCriteria"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.MatchingCriteria"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
__Request body for Actual Config Migration__

It contains two required fields:
- forceMonitor
- migrateSecuredSubnets
*/
type MigrationConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  It can only have Boolean values `true` or `false`.<br>
	If true, force all the policies to be in monitor mode after migration.<br>
	The policies would remain in same the same state they were pre migration, otherwise.
	*/
	IsForceMonitor *bool `json:"isForceMonitor"`
	/*
	  It can only have Boolean values `true` or `false`.<br>
	If true, migrate subnets belonging to VMs secured by NSPs.<br>
	ALL subnets are migrated, otherwise.
	*/
	ShouldMigrateSecuredSubnetsOnly *bool `json:"shouldMigrateSecuredSubnetsOnly,omitempty"`
}

func (p *MigrationConfig) MarshalJSON() ([]byte, error) {
	type MigrationConfigProxy MigrationConfig
	return json.Marshal(struct {
		*MigrationConfigProxy
		IsForceMonitor *bool `json:"isForceMonitor,omitempty"`
	}{
		MigrationConfigProxy: (*MigrationConfigProxy)(p),
		IsForceMonitor:       p.IsForceMonitor,
	})
}

func NewMigrationConfig() *MigrationConfig {
	p := new(MigrationConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.MigrationConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.MigrationConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
__Request body for Config Migration__

It contains two fields:
- isDryRun
- spec

The spec can be one of two types based on isDryRun boolean value.

A sample request body with the flag set would look like this:
```
{
"isDryRun": true,
"spec": {
  "policyNames" : [],
  "shouldIncludeSecureSubnetsInfo" : false
  }
}
```

A sample request body with the flag unset would look like this:
```
{
"isDryRun": false,
"spec": {
  "isForceMonitor" : true,
  "shouldMigrateSecuredSubnetsOnly" : false
  }
}
```
*/
type MigrationConfigSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  It can only have Boolean values `true` or `false`.<br>
	If true, it will trigger a dry run of Flow Config Migration. The policies and subnets will not be actually migrated.<br>
	If false, the config migration will be triggered.
	*/
	IsDryRun *bool `json:"isDryRun,omitempty"`
	/*

	 */
	SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`

	Spec *OneOfMigrationConfigSpecSpec `json:"spec"`
}

func (p *MigrationConfigSpec) MarshalJSON() ([]byte, error) {
	type MigrationConfigSpecProxy MigrationConfigSpec
	return json.Marshal(struct {
		*MigrationConfigSpecProxy
		Spec *OneOfMigrationConfigSpecSpec `json:"spec,omitempty"`
	}{
		MigrationConfigSpecProxy: (*MigrationConfigSpecProxy)(p),
		Spec:                     p.Spec,
	})
}

func NewMigrationConfigSpec() *MigrationConfigSpec {
	p := new(MigrationConfigSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.MigrationConfigSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.MigrationConfigSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDryRun = new(bool)
	*p.IsDryRun = false

	return p
}

func (p *MigrationConfigSpec) GetSpec() interface{} {
	if nil == p.Spec {
		return nil
	}
	return p.Spec.GetValue()
}

func (p *MigrationConfigSpec) SetSpec(v interface{}) error {
	if nil == p.Spec {
		p.Spec = NewOneOfMigrationConfigSpecSpec()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/flow-migrator/$actions/migrate Post operation
*/
type MigrationConfigTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrationConfigTaskResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMigrationConfigTaskResponse() *MigrationConfigTaskResponse {
	p := new(MigrationConfigTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.MigrationConfigTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.MigrationConfigTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrationConfigTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrationConfigTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrationConfigTaskResponseData()
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
_Body for requesting Config Migration Summary_

It contains two fields:
  - policyNames
  - isIncludeSecureSubnetsInfo
*/
type MigrationSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of FNS 1.0 policies selected for Migration Summary.
	*/
	PolicyNames []string `json:"policyNames,omitempty"`
	/*
	  It can only have Boolean values `true` or `false`.<br>
	If true, shows migration info ONLY for subnets belonging to VMs secured by NSPs.<br>
	Shows migration info for ALL subnets, otherwise.
	*/
	ShouldIncludeSecureSubnetsInfo *bool `json:"shouldIncludeSecureSubnetsInfo,omitempty"`
}

func NewMigrationSummary() *MigrationSummary {
	p := new(MigrationSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.MigrationSummary"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.MigrationSummary"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A rule for specifying allowed traffic for an application.
*/
type NSPApplicationRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A reference to an address group.
	*/
	DestAddressGroup *string `json:"destAddressGroup,omitempty"`

	DestAllowSpec *AllowType `json:"destAllowSpec,omitempty"`

	DestCategories []string `json:"destCategories,omitempty"`

	DestSubnet *import1.IPv4Address `json:"destSubnet,omitempty"`
	/*
	  Icmp Type Code List.
	*/
	IcmpServices []IcmpTypeCodeSpec `json:"icmpServices,omitempty"`

	IsAllProtocolAllowed *bool `json:"isAllProtocolAllowed,omitempty"`

	NetworkFunctionChainExtId *string `json:"networkFunctionChainExtId,omitempty"`

	SecuredGroup []string `json:"securedGroup"`

	ServiceGroups []string `json:"serviceGroups,omitempty"`
	/*
	  A reference to an address group.
	*/
	SourceAddressGroup *string `json:"sourceAddressGroup,omitempty"`

	SourceAllowSpec *AllowType `json:"sourceAllowSpec,omitempty"`

	SourceCategories []string `json:"sourceCategories,omitempty"`

	SourceSubnet *import1.IPv4Address `json:"sourceSubnet,omitempty"`

	TcpServices []TcpPortRangeSpec `json:"tcpServices,omitempty"`

	UdpServices []UdpPortRangeSpec `json:"udpServices,omitempty"`
}

func (p *NSPApplicationRuleSpec) MarshalJSON() ([]byte, error) {
	type NSPApplicationRuleSpecProxy NSPApplicationRuleSpec
	return json.Marshal(struct {
		*NSPApplicationRuleSpecProxy
		SecuredGroup []string `json:"securedGroup,omitempty"`
	}{
		NSPApplicationRuleSpecProxy: (*NSPApplicationRuleSpecProxy)(p),
		SecuredGroup:                p.SecuredGroup,
	})
}

func NewNSPApplicationRuleSpec() *NSPApplicationRuleSpec {
	p := new(NSPApplicationRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NSPApplicationRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NSPApplicationRuleSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A rule for specifying allowed traffic inside of a secured entity group.
*/
type NSPIntraEntityGroupRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SecuredGroup []string `json:"securedGroup"`

	SecuredGroupAction *IntraEntityGroupRuleAction `json:"securedGroupAction"`
}

func (p *NSPIntraEntityGroupRuleSpec) MarshalJSON() ([]byte, error) {
	type NSPIntraEntityGroupRuleSpecProxy NSPIntraEntityGroupRuleSpec
	return json.Marshal(struct {
		*NSPIntraEntityGroupRuleSpecProxy
		SecuredGroup       []string                    `json:"securedGroup,omitempty"`
		SecuredGroupAction *IntraEntityGroupRuleAction `json:"securedGroupAction,omitempty"`
	}{
		NSPIntraEntityGroupRuleSpecProxy: (*NSPIntraEntityGroupRuleSpecProxy)(p),
		SecuredGroup:                     p.SecuredGroup,
		SecuredGroupAction:               p.SecuredGroupAction,
	})
}

func NewNSPIntraEntityGroupRuleSpec() *NSPIntraEntityGroupRuleSpec {
	p := new(NSPIntraEntityGroupRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NSPIntraEntityGroupRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NSPIntraEntityGroupRuleSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A rule for specifying that two environments should be isolated from each other. Specify both 'firstIsolationGroup' and 'secondIsolationGroup'.
*/
type NSPTwoEnvIsolationRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FirstIsolationGroup []string `json:"firstIsolationGroup"`

	SecondIsolationGroup []string `json:"secondIsolationGroup"`
}

func (p *NSPTwoEnvIsolationRuleSpec) MarshalJSON() ([]byte, error) {
	type NSPTwoEnvIsolationRuleSpecProxy NSPTwoEnvIsolationRuleSpec
	return json.Marshal(struct {
		*NSPTwoEnvIsolationRuleSpecProxy
		FirstIsolationGroup  []string `json:"firstIsolationGroup,omitempty"`
		SecondIsolationGroup []string `json:"secondIsolationGroup,omitempty"`
	}{
		NSPTwoEnvIsolationRuleSpecProxy: (*NSPTwoEnvIsolationRuleSpecProxy)(p),
		FirstIsolationGroup:             p.FirstIsolationGroup,
		SecondIsolationGroup:            p.SecondIsolationGroup,
	})
}

func NewNSPTwoEnvIsolationRuleSpec() *NSPTwoEnvIsolationRuleSpec {
	p := new(NSPTwoEnvIsolationRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NSPTwoEnvIsolationRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NSPTwoEnvIsolationRuleSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NetworkSecurityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  A user defined annotation for a policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  If Hitlog is enabled
	*/
	IsHitlogEnabled *bool `json:"isHitlogEnabled,omitempty"`
	/*
	  If Ipv6 Traffic is allowed
	*/
	IsIpv6TrafficAllowed *bool `json:"isIpv6TrafficAllowed,omitempty"`

	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`

	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Flow Network Security Policy.<br>
	It can have a maximum length of 63 characters.
	*/
	Name *string `json:"name"`
	/*
	  A list of rules that form a policy. For isolation policies, use isolation rules; for application or quarantine policies, use application rules.
	*/
	Rules []NetworkSecurityPolicyRule `json:"rules,omitempty"`

	SecuredGroups []string `json:"securedGroups,omitempty"`

	State *PolicyState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *PolicyType `json:"type"`

	VpcExtId *string `json:"vpcExtId,omitempty"`
}

func (p *NetworkSecurityPolicy) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyProxy NetworkSecurityPolicy
	return json.Marshal(struct {
		*NetworkSecurityPolicyProxy
		Name *string     `json:"name,omitempty"`
		Type *PolicyType `json:"type,omitempty"`
	}{
		NetworkSecurityPolicyProxy: (*NetworkSecurityPolicyProxy)(p),
		Name:                       p.Name,
		Type:                       p.Type,
	})
}

func NewNetworkSecurityPolicy() *NetworkSecurityPolicy {
	p := new(NetworkSecurityPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/policies Get operation
*/
type NetworkSecurityPolicyGetListResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNetworkSecurityPolicyGetListResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyGetListResponse() *NetworkSecurityPolicyGetListResponse {
	p := new(NetworkSecurityPolicyGetListResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyGetListResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyGetListResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyGetListResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyGetListResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNetworkSecurityPolicyGetListResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/policies/{extId} Get operation
*/
type NetworkSecurityPolicyGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNetworkSecurityPolicyGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyGetResponse() *NetworkSecurityPolicyGetResponse {
	p := new(NetworkSecurityPolicyGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNetworkSecurityPolicyGetResponseData()
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
Network security rule import response data.
*/
type NetworkSecurityPolicyImportEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the entity.
	*/
	EntityName *string `json:"entityName"`

	EntityType *NetworkSecurityPolicyImportEntityType `json:"entityType"`

	EntityUpdateType *NetworkSecurityPolicyImportEntityUpdateType `json:"entityUpdateType"`
}

func (p *NetworkSecurityPolicyImportEntity) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyImportEntityProxy NetworkSecurityPolicyImportEntity
	return json.Marshal(struct {
		*NetworkSecurityPolicyImportEntityProxy
		EntityName       *string                                      `json:"entityName,omitempty"`
		EntityType       *NetworkSecurityPolicyImportEntityType       `json:"entityType,omitempty"`
		EntityUpdateType *NetworkSecurityPolicyImportEntityUpdateType `json:"entityUpdateType,omitempty"`
	}{
		NetworkSecurityPolicyImportEntityProxy: (*NetworkSecurityPolicyImportEntityProxy)(p),
		EntityName:                             p.EntityName,
		EntityType:                             p.EntityType,
		EntityUpdateType:                       p.EntityUpdateType,
	})
}

func NewNetworkSecurityPolicyImportEntity() *NetworkSecurityPolicyImportEntity {
	p := new(NetworkSecurityPolicyImportEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyImportEntity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyImportEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the entity.
*/
type NetworkSecurityPolicyImportEntityType int

const (
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_UNKNOWN                NetworkSecurityPolicyImportEntityType = 0
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_REDACTED               NetworkSecurityPolicyImportEntityType = 1
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_CATEGORY               NetworkSecurityPolicyImportEntityType = 2
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_NETWORK_FUNCTION_CHAIN NetworkSecurityPolicyImportEntityType = 3
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_POLICY                 NetworkSecurityPolicyImportEntityType = 4
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_ADDRESS_GROUP          NetworkSecurityPolicyImportEntityType = 5
	NETWORKSECURITYPOLICYIMPORTENTITYTYPE_SERVICE_GROUP          NetworkSecurityPolicyImportEntityType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkSecurityPolicyImportEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"NETWORK_FUNCTION_CHAIN",
		"POLICY",
		"ADDRESS_GROUP",
		"SERVICE_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkSecurityPolicyImportEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"NETWORK_FUNCTION_CHAIN",
		"POLICY",
		"ADDRESS_GROUP",
		"SERVICE_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkSecurityPolicyImportEntityType) index(name string) NetworkSecurityPolicyImportEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORY",
		"NETWORK_FUNCTION_CHAIN",
		"POLICY",
		"ADDRESS_GROUP",
		"SERVICE_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkSecurityPolicyImportEntityType(idx)
		}
	}
	return NETWORKSECURITYPOLICYIMPORTENTITYTYPE_UNKNOWN
}

func (e *NetworkSecurityPolicyImportEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkSecurityPolicyImportEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkSecurityPolicyImportEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkSecurityPolicyImportEntityType) Ref() *NetworkSecurityPolicyImportEntityType {
	return &e
}

/*
Type of update of the entity.
*/
type NetworkSecurityPolicyImportEntityUpdateType int

const (
	NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_UNKNOWN  NetworkSecurityPolicyImportEntityUpdateType = 0
	NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_REDACTED NetworkSecurityPolicyImportEntityUpdateType = 1
	NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_ADDED    NetworkSecurityPolicyImportEntityUpdateType = 2
	NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_DELETED  NetworkSecurityPolicyImportEntityUpdateType = 3
	NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_MODIFIED NetworkSecurityPolicyImportEntityUpdateType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkSecurityPolicyImportEntityUpdateType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDED",
		"DELETED",
		"MODIFIED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkSecurityPolicyImportEntityUpdateType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDED",
		"DELETED",
		"MODIFIED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkSecurityPolicyImportEntityUpdateType) index(name string) NetworkSecurityPolicyImportEntityUpdateType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDED",
		"DELETED",
		"MODIFIED",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkSecurityPolicyImportEntityUpdateType(idx)
		}
	}
	return NETWORKSECURITYPOLICYIMPORTENTITYUPDATETYPE_UNKNOWN
}

func (e *NetworkSecurityPolicyImportEntityUpdateType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkSecurityPolicyImportEntityUpdateType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkSecurityPolicyImportEntityUpdateType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkSecurityPolicyImportEntityUpdateType) Ref() *NetworkSecurityPolicyImportEntityUpdateType {
	return &e
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/policies/$actions/import Post operation
*/
type NetworkSecurityPolicyImportResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNetworkSecurityPolicyImportResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyImportResponse() *NetworkSecurityPolicyImportResponse {
	p := new(NetworkSecurityPolicyImportResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyImportResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyImportResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyImportResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyImportResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNetworkSecurityPolicyImportResponseData()
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
Cumulative and type based policy counts pre and post Flow migration to FNS 2.0.<br>
Contains the summary total policy counts and the policy counts grouped by policy type.
*/
type NetworkSecurityPolicyMigrationCountSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PolicyTypeCountsSummary []NetworkSecurityPolicyMigrationTypeCountInfo `json:"policyTypeCountsSummary,omitempty"`

	TotalCountsSummary *NetworkSecurityPolicyMigrationTotalCountInfo `json:"totalCountsSummary"`
}

func (p *NetworkSecurityPolicyMigrationCountSummary) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyMigrationCountSummaryProxy NetworkSecurityPolicyMigrationCountSummary
	return json.Marshal(struct {
		*NetworkSecurityPolicyMigrationCountSummaryProxy
		TotalCountsSummary *NetworkSecurityPolicyMigrationTotalCountInfo `json:"totalCountsSummary,omitempty"`
	}{
		NetworkSecurityPolicyMigrationCountSummaryProxy: (*NetworkSecurityPolicyMigrationCountSummaryProxy)(p),
		TotalCountsSummary: p.TotalCountsSummary,
	})
}

func NewNetworkSecurityPolicyMigrationCountSummary() *NetworkSecurityPolicyMigrationCountSummary {
	p := new(NetworkSecurityPolicyMigrationCountSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationCountSummary"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationCountSummary"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The meta information about a Flow Network Security Policy.<br>
The info captures the secured Categories, state of policy, policy type, last modified time,
previewId and options.<br>
The PreviewId can be used to fetch the complete FNS 2.0 policy preview using Preview GET API.<br>
A sample call would look like
```
/microseg/v4.0.a1/flow-migrator/preview/8f094d6c-d7b2-32c2-8223-462c6c5e06db
```
*/
type NetworkSecurityPolicyMigrationMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A user defined annotation for a policy metadata during migration.
	*/
	Description *string `json:"description,omitempty"`

	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`

	Name *string `json:"name,omitempty"`

	Options *ConfigMigrationPolicyOptions `json:"options,omitempty"`
	/*
	  This is a system generated identifier which can be used to preview the complete FNS 2.0
	policy corresponding with a FNS 1.0 policy using Preview GET API.
	A sample call would look like
	```
	/microseg/v4.0.a1/flow-migrator/preview/8f094d6c-d7b2-32c2-8223-462c6c5e06db
	```
	*/
	PreviewId *string `json:"previewId,omitempty"`
	/*
	  List of categories external IDs being secured by the Flow Network Security Policy.
	*/
	SecuredGroupCategoryUuids []string `json:"securedGroupCategoryUuids"`

	State *PolicyState `json:"state"`

	Type *ConfigMigrationPolicyType `json:"type,omitempty"`
}

func (p *NetworkSecurityPolicyMigrationMetadata) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyMigrationMetadataProxy NetworkSecurityPolicyMigrationMetadata
	return json.Marshal(struct {
		*NetworkSecurityPolicyMigrationMetadataProxy
		SecuredGroupCategoryUuids []string     `json:"securedGroupCategoryUuids,omitempty"`
		State                     *PolicyState `json:"state,omitempty"`
	}{
		NetworkSecurityPolicyMigrationMetadataProxy: (*NetworkSecurityPolicyMigrationMetadataProxy)(p),
		SecuredGroupCategoryUuids:                   p.SecuredGroupCategoryUuids,
		State:                                       p.State,
	})
}

func NewNetworkSecurityPolicyMigrationMetadata() *NetworkSecurityPolicyMigrationMetadata {
	p := new(NetworkSecurityPolicyMigrationMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationMetadata"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationMetadata"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Subnet information to be communicated for subnet migration as part of Flow migration to FNS 2.0.<br>
The info includes vlanID, vlanName and subnetUuid.
*/
type NetworkSecurityPolicyMigrationSubnetSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SubnetUuid *string `json:"subnetUuid"`

	VlanID *int `json:"vlanID"`

	VlanName *string `json:"vlanName,omitempty"`
}

func (p *NetworkSecurityPolicyMigrationSubnetSummary) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyMigrationSubnetSummaryProxy NetworkSecurityPolicyMigrationSubnetSummary
	return json.Marshal(struct {
		*NetworkSecurityPolicyMigrationSubnetSummaryProxy
		SubnetUuid *string `json:"subnetUuid,omitempty"`
		VlanID     *int    `json:"vlanID,omitempty"`
	}{
		NetworkSecurityPolicyMigrationSubnetSummaryProxy: (*NetworkSecurityPolicyMigrationSubnetSummaryProxy)(p),
		SubnetUuid: p.SubnetUuid,
		VlanID:     p.VlanID,
	})
}

func NewNetworkSecurityPolicyMigrationSubnetSummary() *NetworkSecurityPolicyMigrationSubnetSummary {
	p := new(NetworkSecurityPolicyMigrationSubnetSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationSubnetSummary"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationSubnetSummary"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specifies the policy name along with the policy metadata.<br>
For cases where a single FNS 1.0 policy conversion results in multiple FNS 2.0 policies,
it also contains a list of policy metadata corresponding to those system generated policies.
*/
type NetworkSecurityPolicyMigrationSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PolicyMetadata *NetworkSecurityPolicyMigrationMetadata `json:"policyMetadata"`
	/*
	  Name of the Flow Network Security Policy.<br>
	It can have a maximum length of 63 characters.
	*/
	PolicyName *string `json:"policyName"`
	/*
	  A list of Policy metadata of system generated isolation policies that would get created as part of
	a successful Flow migration.
	*/
	SystemGenPolicyMetadata []NetworkSecurityPolicyMigrationMetadata `json:"systemGenPolicyMetadata,omitempty"`
}

func (p *NetworkSecurityPolicyMigrationSummary) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyMigrationSummaryProxy NetworkSecurityPolicyMigrationSummary
	return json.Marshal(struct {
		*NetworkSecurityPolicyMigrationSummaryProxy
		PolicyMetadata *NetworkSecurityPolicyMigrationMetadata `json:"policyMetadata,omitempty"`
		PolicyName     *string                                 `json:"policyName,omitempty"`
	}{
		NetworkSecurityPolicyMigrationSummaryProxy: (*NetworkSecurityPolicyMigrationSummaryProxy)(p),
		PolicyMetadata: p.PolicyMetadata,
		PolicyName:     p.PolicyName,
	})
}

func NewNetworkSecurityPolicyMigrationSummary() *NetworkSecurityPolicyMigrationSummary {
	p := new(NetworkSecurityPolicyMigrationSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationSummary"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationSummary"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Outlines the error code and description of failure during summarising the migration.
*/
type NetworkSecurityPolicyMigrationSummaryFailures struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	ErrorDescription *string `json:"errorDescription,omitempty"`
}

func NewNetworkSecurityPolicyMigrationSummaryFailures() *NetworkSecurityPolicyMigrationSummaryFailures {
	p := new(NetworkSecurityPolicyMigrationSummaryFailures)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationSummaryFailures"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationSummaryFailures"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cumulative Policy counts pre and post Flow migration to FNS 2.0.
Contains total count of FNS 2.0 policies that'll be created post successful migration,
total count of FNS 1.0 policies before migration and the total count of FNS 2.0 policies
which were system generated as a result of FNS 1.0 to FNS 2.0 conversion.<br>
*/
type NetworkSecurityPolicyMigrationTotalCountInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NewPoliciesCount *int `json:"newPoliciesCount"`

	OldPoliciesCount *int `json:"oldPoliciesCount"`
	/*
	  Count of isolation policies that would be genrated post migration to FNS 2.0.
	*/
	SystemDefinedPoliciesCount *int `json:"systemDefinedPoliciesCount"`
}

func (p *NetworkSecurityPolicyMigrationTotalCountInfo) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyMigrationTotalCountInfoProxy NetworkSecurityPolicyMigrationTotalCountInfo
	return json.Marshal(struct {
		*NetworkSecurityPolicyMigrationTotalCountInfoProxy
		NewPoliciesCount           *int `json:"newPoliciesCount,omitempty"`
		OldPoliciesCount           *int `json:"oldPoliciesCount,omitempty"`
		SystemDefinedPoliciesCount *int `json:"systemDefinedPoliciesCount,omitempty"`
	}{
		NetworkSecurityPolicyMigrationTotalCountInfoProxy: (*NetworkSecurityPolicyMigrationTotalCountInfoProxy)(p),
		NewPoliciesCount:           p.NewPoliciesCount,
		OldPoliciesCount:           p.OldPoliciesCount,
		SystemDefinedPoliciesCount: p.SystemDefinedPoliciesCount,
	})
}

func NewNetworkSecurityPolicyMigrationTotalCountInfo() *NetworkSecurityPolicyMigrationTotalCountInfo {
	p := new(NetworkSecurityPolicyMigrationTotalCountInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationTotalCountInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationTotalCountInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Policy type based counts pre and post Flow migration to FNS 2.0.<br>
This will be grouped by policy types which can be one of following
four types:
  - QUARANTINE POLICY
  - ISOLATION POLICY
  - APPLICATION POLICY
  - AD POLICY
*/
type NetworkSecurityPolicyMigrationTypeCountInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CountSummary *NetworkSecurityPolicyMigrationTotalCountInfo `json:"countSummary,omitempty"`

	Type *ConfigMigrationPolicyType `json:"type,omitempty"`
}

func NewNetworkSecurityPolicyMigrationTypeCountInfo() *NetworkSecurityPolicyMigrationTypeCountInfo {
	p := new(NetworkSecurityPolicyMigrationTypeCountInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyMigrationTypeCountInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyMigrationTypeCountInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NetworkSecurityPolicyRule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A user defined annotation for a rule.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Rule UUID; should be used to identify an individual rule uniquely.
	*/
	Id *string `json:"id,omitempty"`

	SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`

	Spec *OneOfNetworkSecurityPolicyRuleSpec `json:"spec"`

	Type *RuleType `json:"type"`
}

func (p *NetworkSecurityPolicyRule) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyRuleProxy NetworkSecurityPolicyRule
	return json.Marshal(struct {
		*NetworkSecurityPolicyRuleProxy
		Spec *OneOfNetworkSecurityPolicyRuleSpec `json:"spec,omitempty"`
		Type *RuleType                           `json:"type,omitempty"`
	}{
		NetworkSecurityPolicyRuleProxy: (*NetworkSecurityPolicyRuleProxy)(p),
		Spec:                           p.Spec,
		Type:                           p.Type,
	})
}

func NewNetworkSecurityPolicyRule() *NetworkSecurityPolicyRule {
	p := new(NetworkSecurityPolicyRule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyRule"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyRule"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyRule) GetSpec() interface{} {
	if nil == p.Spec {
		return nil
	}
	return p.Spec.GetValue()
}

func (p *NetworkSecurityPolicyRule) SetSpec(v interface{}) error {
	if nil == p.Spec {
		p.Spec = NewOneOfNetworkSecurityPolicyRuleSpec()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/policies/{extId}/rules Get operation
*/
type NetworkSecurityPolicyRulesGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNetworkSecurityPolicyRulesGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyRulesGetResponse() *NetworkSecurityPolicyRulesGetResponse {
	p := new(NetworkSecurityPolicyRulesGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyRulesGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyRulesGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyRulesGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyRulesGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNetworkSecurityPolicyRulesGetResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/policies/$actions/set-allow-ipv6-traffic Post operation
*/
type NetworkSecurityPolicyTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfNetworkSecurityPolicyTaskResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewNetworkSecurityPolicyTaskResponse() *NetworkSecurityPolicyTaskResponse {
	p := new(NetworkSecurityPolicyTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.NetworkSecurityPolicyTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.NetworkSecurityPolicyTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NetworkSecurityPolicyTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *NetworkSecurityPolicyTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfNetworkSecurityPolicyTaskResponseData()
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
List of available tiers
*/
type OptionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of the tier
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Memory in GB of the tier
	*/
	MemoryInGB *int `json:"memoryInGB,omitempty"`

	Name *TierName `json:"name,omitempty"`
}

func NewOptionSpec() *OptionSpec {
	p := new(OptionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.OptionSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.OptionSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PolicyReferenceActionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of external ids for a set of network security policies.
	*/
	PolicyExtIds []string `json:"policyExtIds"`

	Value *bool `json:"value"`
}

func (p *PolicyReferenceActionSpec) MarshalJSON() ([]byte, error) {
	type PolicyReferenceActionSpecProxy PolicyReferenceActionSpec
	return json.Marshal(struct {
		*PolicyReferenceActionSpecProxy
		PolicyExtIds []string `json:"policyExtIds,omitempty"`
		Value        *bool    `json:"value,omitempty"`
	}{
		PolicyReferenceActionSpecProxy: (*PolicyReferenceActionSpecProxy)(p),
		PolicyExtIds:                   p.PolicyExtIds,
		Value:                          p.Value,
	})
}

func NewPolicyReferenceActionSpec() *PolicyReferenceActionSpec {
	p := new(PolicyReferenceActionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.PolicyReferenceActionSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.PolicyReferenceActionSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A list of external ids for a set of network security policies.
*/
type PolicyReferenceSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of external ids for a set of network security policies.
	*/
	PolicyExtIds []string `json:"policyExtIds"`
}

func (p *PolicyReferenceSpec) MarshalJSON() ([]byte, error) {
	type PolicyReferenceSpecProxy PolicyReferenceSpec
	return json.Marshal(struct {
		*PolicyReferenceSpecProxy
		PolicyExtIds []string `json:"policyExtIds,omitempty"`
	}{
		PolicyReferenceSpecProxy: (*PolicyReferenceSpecProxy)(p),
		PolicyExtIds:             p.PolicyExtIds,
	})
}

func NewPolicyReferenceSpec() *PolicyReferenceSpec {
	p := new(PolicyReferenceSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.PolicyReferenceSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.PolicyReferenceSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PolicyReferenceStateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of external ids for a set of network security policies.
	*/
	PolicyExtIds []string `json:"policyExtIds"`

	State *PolicyState `json:"state"`
}

func (p *PolicyReferenceStateSpec) MarshalJSON() ([]byte, error) {
	type PolicyReferenceStateSpecProxy PolicyReferenceStateSpec
	return json.Marshal(struct {
		*PolicyReferenceStateSpecProxy
		PolicyExtIds []string     `json:"policyExtIds,omitempty"`
		State        *PolicyState `json:"state,omitempty"`
	}{
		PolicyReferenceStateSpecProxy: (*PolicyReferenceStateSpecProxy)(p),
		PolicyExtIds:                  p.PolicyExtIds,
		State:                         p.State,
	})
}

func NewPolicyReferenceStateSpec() *PolicyReferenceStateSpec {
	p := new(PolicyReferenceStateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.PolicyReferenceStateSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.PolicyReferenceStateSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Whether the policy is applied or monitored; can be omitted or set null to save the policy without applying or monitoring it.
*/
type PolicyState int

const (
	POLICYSTATE_UNKNOWN  PolicyState = 0
	POLICYSTATE_REDACTED PolicyState = 1
	POLICYSTATE_SAVE     PolicyState = 2
	POLICYSTATE_MONITOR  PolicyState = 3
	POLICYSTATE_ENFORCE  PolicyState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PolicyState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAVE",
		"MONITOR",
		"ENFORCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PolicyState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAVE",
		"MONITOR",
		"ENFORCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PolicyState) index(name string) PolicyState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAVE",
		"MONITOR",
		"ENFORCE",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicyState(idx)
		}
	}
	return POLICYSTATE_UNKNOWN
}

func (e *PolicyState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicyState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicyState) Ref() *PolicyState {
	return &e
}

/*
Defines the type of rules that can be used in a policy.
*/
type PolicyType int

const (
	POLICYTYPE_UNKNOWN     PolicyType = 0
	POLICYTYPE_REDACTED    PolicyType = 1
	POLICYTYPE_QUARANTINE  PolicyType = 2
	POLICYTYPE_ISOLATION   PolicyType = 3
	POLICYTYPE_APPLICATION PolicyType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PolicyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PolicyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PolicyType) index(name string) PolicyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"ISOLATION",
		"APPLICATION",
	}
	for idx := range names {
		if names[idx] == name {
			return PolicyType(idx)
		}
	}
	return POLICYTYPE_UNKNOWN
}

func (e *PolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PolicyType) Ref() *PolicyType {
	return &e
}

/*
The type for a rule - the value chosen here restricts which specification can be chosen.
*/
type RuleType int

const (
	RULETYPE_UNKNOWN           RuleType = 0
	RULETYPE_REDACTED          RuleType = 1
	RULETYPE_QUARANTINE        RuleType = 2
	RULETYPE_TWO_ENV_ISOLATION RuleType = 3
	RULETYPE_APPLICATION       RuleType = 4
	RULETYPE_INTRA_GROUP       RuleType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RuleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"TWO_ENV_ISOLATION",
		"APPLICATION",
		"INTRA_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RuleType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"TWO_ENV_ISOLATION",
		"APPLICATION",
		"INTRA_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RuleType) index(name string) RuleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUARANTINE",
		"TWO_ENV_ISOLATION",
		"APPLICATION",
		"INTRA_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return RuleType(idx)
		}
	}
	return RULETYPE_UNKNOWN
}

func (e *RuleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RuleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RuleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RuleType) Ref() *RuleType {
	return &e
}

type ServiceGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A user defined annotation for a service group.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Icmp Type Code List.
	*/
	IcmpServices []IcmpTypeCodeSpec `json:"icmpServices,omitempty"`
	/*
	  Service Group is system defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A short identifier for a service group.
	*/
	Name *string `json:"name"`
	/*
	  List of TCP ports in the service.
	*/
	TcpServices []TcpPortRangeSpec `json:"tcpServices,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of UDP ports in the service.
	*/
	UdpServices []UdpPortRangeSpec `json:"udpServices,omitempty"`
}

func (p *ServiceGroup) MarshalJSON() ([]byte, error) {
	type ServiceGroupProxy ServiceGroup
	return json.Marshal(struct {
		*ServiceGroupProxy
		Name *string `json:"name,omitempty"`
	}{
		ServiceGroupProxy: (*ServiceGroupProxy)(p),
		Name:              p.Name,
	})
}

func NewServiceGroup() *ServiceGroup {
	p := new(ServiceGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/service-groups/{extId} Get operation
*/
type ServiceGroupGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfServiceGroupGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewServiceGroupGetResponse() *ServiceGroupGetResponse {
	p := new(ServiceGroupGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ServiceGroupGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ServiceGroupGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfServiceGroupGetResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/service-groups Get operation
*/
type ServiceGroupListResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfServiceGroupListResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewServiceGroupListResponse() *ServiceGroupListResponse {
	p := new(ServiceGroupListResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupListResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupListResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ServiceGroupListResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ServiceGroupListResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfServiceGroupListResponseData()
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
REST response for all response codes in API path /microseg/v4.0.a1/config/service-groups/$actions/build-policy-association Post operation
*/
type ServiceGroupPolicyAssociationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfServiceGroupPolicyAssociationResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewServiceGroupPolicyAssociationResponse() *ServiceGroupPolicyAssociationResponse {
	p := new(ServiceGroupPolicyAssociationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupPolicyAssociationResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupPolicyAssociationResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ServiceGroupPolicyAssociationResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ServiceGroupPolicyAssociationResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfServiceGroupPolicyAssociationResponseData()
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
Reference to policy associated with service group.
*/
type ServiceGroupPolicyReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PolicyUuids []string `json:"policyUuids,omitempty"`

	ServiceGroupUuid *string `json:"serviceGroupUuid,omitempty"`
}

func NewServiceGroupPolicyReference() *ServiceGroupPolicyReference {
	p := new(ServiceGroupPolicyReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupPolicyReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupPolicyReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of service group UUID
*/
type ServiceGroupReferenceSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ServiceGroupUuidList []string `json:"serviceGroupUuidList"`
}

func (p *ServiceGroupReferenceSpec) MarshalJSON() ([]byte, error) {
	type ServiceGroupReferenceSpecProxy ServiceGroupReferenceSpec
	return json.Marshal(struct {
		*ServiceGroupReferenceSpecProxy
		ServiceGroupUuidList []string `json:"serviceGroupUuidList,omitempty"`
	}{
		ServiceGroupReferenceSpecProxy: (*ServiceGroupReferenceSpecProxy)(p),
		ServiceGroupUuidList:           p.ServiceGroupUuidList,
	})
}

func NewServiceGroupReferenceSpec() *ServiceGroupReferenceSpec {
	p := new(ServiceGroupReferenceSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupReferenceSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupReferenceSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/service-groups/{extId} Delete operation
*/
type ServiceGroupTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfServiceGroupTaskResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewServiceGroupTaskResponse() *ServiceGroupTaskResponse {
	p := new(ServiceGroupTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ServiceGroupTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.ServiceGroupTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ServiceGroupTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ServiceGroupTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfServiceGroupTaskResponseData()
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
Range of TCP/UDP ports.
*/
type TcpPortRangeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndPort *int `json:"endPort"`

	StartPort *int `json:"startPort"`
}

func (p *TcpPortRangeSpec) MarshalJSON() ([]byte, error) {
	type TcpPortRangeSpecProxy TcpPortRangeSpec
	return json.Marshal(struct {
		*TcpPortRangeSpecProxy
		EndPort   *int `json:"endPort,omitempty"`
		StartPort *int `json:"startPort,omitempty"`
	}{
		TcpPortRangeSpecProxy: (*TcpPortRangeSpecProxy)(p),
		EndPort:               p.EndPort,
		StartPort:             p.StartPort,
	})
}

func NewTcpPortRangeSpec() *TcpPortRangeSpec {
	p := new(TcpPortRangeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.TcpPortRangeSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.TcpPortRangeSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Tier struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActiveTier *TierName `json:"activeTier"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Tier) MarshalJSON() ([]byte, error) {
	type TierProxy Tier
	return json.Marshal(struct {
		*TierProxy
		ActiveTier *TierName `json:"activeTier,omitempty"`
	}{
		TierProxy:  (*TierProxy)(p),
		ActiveTier: p.ActiveTier,
	})
}

func NewTier() *Tier {
	p := new(Tier)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.Tier"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.Tier"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/cadmus/tiers Get operation
*/
type TierGetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTierGetResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTierGetResponse() *TierGetResponse {
	p := new(TierGetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.TierGetResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.TierGetResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TierGetResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TierGetResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTierGetResponseData()
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
Name of the tier
*/
type TierName int

const (
	TIERNAME_UNKNOWN  TierName = 0
	TIERNAME_REDACTED TierName = 1
	TIERNAME_SMALL    TierName = 2
	TIERNAME_MEDIUM   TierName = 3
	TIERNAME_LARGE    TierName = 4
	TIERNAME_XLARGE   TierName = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TierName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
		"XLARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TierName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
		"XLARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TierName) index(name string) TierName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SMALL",
		"MEDIUM",
		"LARGE",
		"XLARGE",
	}
	for idx := range names {
		if names[idx] == name {
			return TierName(idx)
		}
	}
	return TIERNAME_UNKNOWN
}

func (e *TierName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TierName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TierName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TierName) Ref() *TierName {
	return &e
}

type TierResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActiveTier *TierName `json:"activeTier"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of available tiers
	*/
	Options []OptionSpec `json:"options,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *TierResponse) MarshalJSON() ([]byte, error) {
	type TierResponseProxy TierResponse
	return json.Marshal(struct {
		*TierResponseProxy
		ActiveTier *TierName `json:"activeTier,omitempty"`
	}{
		TierResponseProxy: (*TierResponseProxy)(p),
		ActiveTier:        p.ActiveTier,
	})
}

func NewTierResponse() *TierResponse {
	p := new(TierResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.TierResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.TierResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.a1/config/cadmus/tiers Put operation
*/
type TierUpdateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTierUpdateResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTierUpdateResponse() *TierUpdateResponse {
	p := new(TierUpdateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.TierUpdateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.TierUpdateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TierUpdateResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TierUpdateResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTierUpdateResponseData()
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
Range of TCP/UDP ports.
*/
type UdpPortRangeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndPort *int `json:"endPort"`

	StartPort *int `json:"startPort"`
}

func (p *UdpPortRangeSpec) MarshalJSON() ([]byte, error) {
	type UdpPortRangeSpecProxy UdpPortRangeSpec
	return json.Marshal(struct {
		*UdpPortRangeSpecProxy
		EndPort   *int `json:"endPort,omitempty"`
		StartPort *int `json:"startPort,omitempty"`
	}{
		UdpPortRangeSpecProxy: (*UdpPortRangeSpecProxy)(p),
		EndPort:               p.EndPort,
		StartPort:             p.StartPort,
	})
}

func NewUdpPortRangeSpec() *UdpPortRangeSpec {
	p := new(UdpPortRangeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UdpPortRangeSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "microseg.v4.r0.a1.config.UdpPortRangeSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfNetworkSecurityPolicyRuleSpec struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType1    *NSPApplicationRuleSpec      `json:"-"`
	oneOfType2    *NSPIntraEntityGroupRuleSpec `json:"-"`
	oneOfType0    *NSPTwoEnvIsolationRuleSpec  `json:"-"`
}

func NewOneOfNetworkSecurityPolicyRuleSpec() *OneOfNetworkSecurityPolicyRuleSpec {
	p := new(OneOfNetworkSecurityPolicyRuleSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyRuleSpec is nil"))
	}
	switch v.(type) {
	case NSPApplicationRuleSpec:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(NSPApplicationRuleSpec)
		}
		*p.oneOfType1 = v.(NSPApplicationRuleSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case NSPIntraEntityGroupRuleSpec:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(NSPIntraEntityGroupRuleSpec)
		}
		*p.oneOfType2 = v.(NSPIntraEntityGroupRuleSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case NSPTwoEnvIsolationRuleSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(NSPTwoEnvIsolationRuleSpec)
		}
		*p.oneOfType0 = v.(NSPTwoEnvIsolationRuleSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(NSPApplicationRuleSpec)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "microseg.v4.config.NSPApplicationRuleSpec" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(NSPApplicationRuleSpec)
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
	vOneOfType2 := new(NSPIntraEntityGroupRuleSpec)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "microseg.v4.config.NSPIntraEntityGroupRuleSpec" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(NSPIntraEntityGroupRuleSpec)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(NSPTwoEnvIsolationRuleSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.NSPTwoEnvIsolationRuleSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(NSPTwoEnvIsolationRuleSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyRuleSpec"))
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyRuleSpec")
}

type OneOfDirectoryServerCreateResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDirectoryServerCreateResponseData() *OneOfDirectoryServerCreateResponseData {
	p := new(OneOfDirectoryServerCreateResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDirectoryServerCreateResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDirectoryServerCreateResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDirectoryServerCreateResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDirectoryServerCreateResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectoryServerCreateResponseData"))
}

func (p *OneOfDirectoryServerCreateResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDirectoryServerCreateResponseData")
}

type OneOfConfigMigrationSummaryGetResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    *ConfigMigrationSummary `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
}

func NewOneOfConfigMigrationSummaryGetResponseData() *OneOfConfigMigrationSummaryGetResponseData {
	p := new(OneOfConfigMigrationSummaryGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigMigrationSummaryGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigMigrationSummaryGetResponseData is nil"))
	}
	switch v.(type) {
	case ConfigMigrationSummary:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ConfigMigrationSummary)
		}
		*p.oneOfType0 = v.(ConfigMigrationSummary)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigMigrationSummaryGetResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfConfigMigrationSummaryGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ConfigMigrationSummary)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.ConfigMigrationSummary" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ConfigMigrationSummary)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigMigrationSummaryGetResponseData"))
}

func (p *OneOfConfigMigrationSummaryGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfConfigMigrationSummaryGetResponseData")
}

type OneOfDsCategoryMappingCreateResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDsCategoryMappingCreateResponseData() *OneOfDsCategoryMappingCreateResponseData {
	p := new(OneOfDsCategoryMappingCreateResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDsCategoryMappingCreateResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDsCategoryMappingCreateResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDsCategoryMappingCreateResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDsCategoryMappingCreateResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDsCategoryMappingCreateResponseData"))
}

func (p *OneOfDsCategoryMappingCreateResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDsCategoryMappingCreateResponseData")
}

type OneOfTierUpdateResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfTierUpdateResponseData() *OneOfTierUpdateResponseData {
	p := new(OneOfTierUpdateResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTierUpdateResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTierUpdateResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTierUpdateResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTierUpdateResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierUpdateResponseData"))
}

func (p *OneOfTierUpdateResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTierUpdateResponseData")
}

type OneOfNetworkSecurityPolicyTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfNetworkSecurityPolicyTaskResponseData() *OneOfNetworkSecurityPolicyTaskResponseData {
	p := new(OneOfNetworkSecurityPolicyTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyTaskResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyTaskResponseData"))
}

func (p *OneOfNetworkSecurityPolicyTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyTaskResponseData")
}

type OneOfConfigMigrationPreviewGetResponseData struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType400  *import3.ErrorResponse          `json:"-"`
	oneOfType0    *ConfigMigrationPreviewResponse `json:"-"`
}

func NewOneOfConfigMigrationPreviewGetResponseData() *OneOfConfigMigrationPreviewGetResponseData {
	p := new(OneOfConfigMigrationPreviewGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConfigMigrationPreviewGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConfigMigrationPreviewGetResponseData is nil"))
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
	case ConfigMigrationPreviewResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ConfigMigrationPreviewResponse)
		}
		*p.oneOfType0 = v.(ConfigMigrationPreviewResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConfigMigrationPreviewGetResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfConfigMigrationPreviewGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ConfigMigrationPreviewResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.ConfigMigrationPreviewResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ConfigMigrationPreviewResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConfigMigrationPreviewGetResponseData"))
}

func (p *OneOfConfigMigrationPreviewGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfConfigMigrationPreviewGetResponseData")
}

type OneOfAddressGroupPolicyAssociationResponseData struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType0    []AddressGroupPolicyReference `json:"-"`
	oneOfType400  *import3.ErrorResponse        `json:"-"`
}

func NewOneOfAddressGroupPolicyAssociationResponseData() *OneOfAddressGroupPolicyAssociationResponseData {
	p := new(OneOfAddressGroupPolicyAssociationResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddressGroupPolicyAssociationResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddressGroupPolicyAssociationResponseData is nil"))
	}
	switch v.(type) {
	case []AddressGroupPolicyReference:
		p.oneOfType0 = v.([]AddressGroupPolicyReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.AddressGroupPolicyReference>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.AddressGroupPolicyReference>"
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddressGroupPolicyAssociationResponseData) GetValue() interface{} {
	if "List<microseg.v4.config.AddressGroupPolicyReference>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddressGroupPolicyAssociationResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]AddressGroupPolicyReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.AddressGroupPolicyReference" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.AddressGroupPolicyReference>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.AddressGroupPolicyReference>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddressGroupPolicyAssociationResponseData"))
}

func (p *OneOfAddressGroupPolicyAssociationResponseData) MarshalJSON() ([]byte, error) {
	if "List<microseg.v4.config.AddressGroupPolicyReference>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddressGroupPolicyAssociationResponseData")
}

type OneOfServiceGroupListResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ServiceGroup         `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfServiceGroupListResponseData() *OneOfServiceGroupListResponseData {
	p := new(OneOfServiceGroupListResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfServiceGroupListResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfServiceGroupListResponseData is nil"))
	}
	switch v.(type) {
	case []ServiceGroup:
		p.oneOfType0 = v.([]ServiceGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.ServiceGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.ServiceGroup>"
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfServiceGroupListResponseData) GetValue() interface{} {
	if "List<microseg.v4.config.ServiceGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfServiceGroupListResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ServiceGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.ServiceGroup" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.ServiceGroup>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.ServiceGroup>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceGroupListResponseData"))
}

func (p *OneOfServiceGroupListResponseData) MarshalJSON() ([]byte, error) {
	if "List<microseg.v4.config.ServiceGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfServiceGroupListResponseData")
}

type OneOfNetworkSecurityPolicyGetListResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    []NetworkSecurityPolicy `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
	oneOfType1    *FileWrapper            `json:"-"`
}

func NewOneOfNetworkSecurityPolicyGetListResponseData() *OneOfNetworkSecurityPolicyGetListResponseData {
	p := new(OneOfNetworkSecurityPolicyGetListResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyGetListResponseData is nil"))
	}
	switch v.(type) {
	case []NetworkSecurityPolicy:
		p.oneOfType0 = v.([]NetworkSecurityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicy>"
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
	case FileWrapper:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(FileWrapper)
		}
		*p.oneOfType1 = v.(FileWrapper)
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

func (p *OneOfNetworkSecurityPolicyGetListResponseData) GetValue() interface{} {
	if "List<microseg.v4.config.NetworkSecurityPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]NetworkSecurityPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.NetworkSecurityPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicy>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(FileWrapper)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "microseg.v4.config.FileWrapper" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(FileWrapper)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyGetListResponseData"))
}

func (p *OneOfNetworkSecurityPolicyGetListResponseData) MarshalJSON() ([]byte, error) {
	if "List<microseg.v4.config.NetworkSecurityPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyGetListResponseData")
}

type OneOfDsCategoryMappingUpdateResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDsCategoryMappingUpdateResponseData() *OneOfDsCategoryMappingUpdateResponseData {
	p := new(OneOfDsCategoryMappingUpdateResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDsCategoryMappingUpdateResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDsCategoryMappingUpdateResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDsCategoryMappingUpdateResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDsCategoryMappingUpdateResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDsCategoryMappingUpdateResponseData"))
}

func (p *OneOfDsCategoryMappingUpdateResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDsCategoryMappingUpdateResponseData")
}

type OneOfDsCategoryMappingsGetListResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []CategoryMapping      `json:"-"`
}

func NewOneOfDsCategoryMappingsGetListResponseData() *OneOfDsCategoryMappingsGetListResponseData {
	p := new(OneOfDsCategoryMappingsGetListResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDsCategoryMappingsGetListResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDsCategoryMappingsGetListResponseData is nil"))
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
	case []CategoryMapping:
		p.oneOfType0 = v.([]CategoryMapping)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.CategoryMapping>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.CategoryMapping>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDsCategoryMappingsGetListResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.CategoryMapping>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfDsCategoryMappingsGetListResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]CategoryMapping)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.CategoryMapping" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.CategoryMapping>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.CategoryMapping>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDsCategoryMappingsGetListResponseData"))
}

func (p *OneOfDsCategoryMappingsGetListResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.CategoryMapping>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDsCategoryMappingsGetListResponseData")
}

type OneOfMigrationConfigTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfMigrationConfigTaskResponseData() *OneOfMigrationConfigTaskResponseData {
	p := new(OneOfMigrationConfigTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrationConfigTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrationConfigTaskResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfMigrationConfigTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrationConfigTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrationConfigTaskResponseData"))
}

func (p *OneOfMigrationConfigTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrationConfigTaskResponseData")
}

type OneOfMigrationConfigSpecSpec struct {
	Discriminator *string           `json:"-"`
	ObjectType_   *string           `json:"-"`
	oneOfType1    *MigrationSummary `json:"-"`
	oneOfType0    *MigrationConfig  `json:"-"`
}

func NewOneOfMigrationConfigSpecSpec() *OneOfMigrationConfigSpecSpec {
	p := new(OneOfMigrationConfigSpecSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrationConfigSpecSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrationConfigSpecSpec is nil"))
	}
	switch v.(type) {
	case MigrationSummary:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(MigrationSummary)
		}
		*p.oneOfType1 = v.(MigrationSummary)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case MigrationConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(MigrationConfig)
		}
		*p.oneOfType0 = v.(MigrationConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfMigrationConfigSpecSpec) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfMigrationConfigSpecSpec) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(MigrationSummary)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "microseg.v4.config.MigrationSummary" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(MigrationSummary)
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
	vOneOfType0 := new(MigrationConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.MigrationConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(MigrationConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrationConfigSpecSpec"))
}

func (p *OneOfMigrationConfigSpecSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfMigrationConfigSpecSpec")
}

type OneOfNetworkSecurityPolicyImportResponseData struct {
	Discriminator *string                             `json:"-"`
	ObjectType_   *string                             `json:"-"`
	oneOfType0    *import4.TaskReference              `json:"-"`
	oneOfType400  *import3.ErrorResponse              `json:"-"`
	oneOfType1    []NetworkSecurityPolicyImportEntity `json:"-"`
}

func NewOneOfNetworkSecurityPolicyImportResponseData() *OneOfNetworkSecurityPolicyImportResponseData {
	p := new(OneOfNetworkSecurityPolicyImportResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyImportResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	case []NetworkSecurityPolicyImportEntity:
		p.oneOfType1 = v.([]NetworkSecurityPolicyImportEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>" == *p.Discriminator {
		return p.oneOfType1
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new([]NetworkSecurityPolicyImportEntity)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if len(*vOneOfType1) == 0 || "microseg.v4.config.NetworkSecurityPolicyImportEntity" == *((*vOneOfType1)[0].ObjectType_) {
			p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyImportResponseData"))
}

func (p *OneOfNetworkSecurityPolicyImportResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.NetworkSecurityPolicyImportEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyImportResponseData")
}

type OneOfNetworkSecurityPolicyRulesGetResponseData struct {
	Discriminator *string                     `json:"-"`
	ObjectType_   *string                     `json:"-"`
	oneOfType0    []NetworkSecurityPolicyRule `json:"-"`
	oneOfType400  *import3.ErrorResponse      `json:"-"`
}

func NewOneOfNetworkSecurityPolicyRulesGetResponseData() *OneOfNetworkSecurityPolicyRulesGetResponseData {
	p := new(OneOfNetworkSecurityPolicyRulesGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyRulesGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyRulesGetResponseData is nil"))
	}
	switch v.(type) {
	case []NetworkSecurityPolicyRule:
		p.oneOfType0 = v.([]NetworkSecurityPolicyRule)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicyRule>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicyRule>"
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyRulesGetResponseData) GetValue() interface{} {
	if "List<microseg.v4.config.NetworkSecurityPolicyRule>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyRulesGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]NetworkSecurityPolicyRule)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.NetworkSecurityPolicyRule" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.NetworkSecurityPolicyRule>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.NetworkSecurityPolicyRule>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyRulesGetResponseData"))
}

func (p *OneOfNetworkSecurityPolicyRulesGetResponseData) MarshalJSON() ([]byte, error) {
	if "List<microseg.v4.config.NetworkSecurityPolicyRule>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyRulesGetResponseData")
}

type OneOfDirectoryServerUpdateResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDirectoryServerUpdateResponseData() *OneOfDirectoryServerUpdateResponseData {
	p := new(OneOfDirectoryServerUpdateResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDirectoryServerUpdateResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDirectoryServerUpdateResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDirectoryServerUpdateResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDirectoryServerUpdateResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectoryServerUpdateResponseData"))
}

func (p *OneOfDirectoryServerUpdateResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDirectoryServerUpdateResponseData")
}

type OneOfAddressGroupGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *AddressGroup          `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAddressGroupGetResponseData() *OneOfAddressGroupGetResponseData {
	p := new(OneOfAddressGroupGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddressGroupGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddressGroupGetResponseData is nil"))
	}
	switch v.(type) {
	case AddressGroup:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AddressGroup)
		}
		*p.oneOfType0 = v.(AddressGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddressGroupGetResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddressGroupGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AddressGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.AddressGroup" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AddressGroup)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddressGroupGetResponseData"))
}

func (p *OneOfAddressGroupGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddressGroupGetResponseData")
}

type OneOfAddressGroupTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfAddressGroupTaskResponseData() *OneOfAddressGroupTaskResponseData {
	p := new(OneOfAddressGroupTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddressGroupTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddressGroupTaskResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddressGroupTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddressGroupTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddressGroupTaskResponseData"))
}

func (p *OneOfAddressGroupTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddressGroupTaskResponseData")
}

type OneOfServiceGroupGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *ServiceGroup          `json:"-"`
}

func NewOneOfServiceGroupGetResponseData() *OneOfServiceGroupGetResponseData {
	p := new(OneOfServiceGroupGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfServiceGroupGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfServiceGroupGetResponseData is nil"))
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
	case ServiceGroup:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ServiceGroup)
		}
		*p.oneOfType0 = v.(ServiceGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfServiceGroupGetResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfServiceGroupGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ServiceGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.ServiceGroup" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ServiceGroup)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceGroupGetResponseData"))
}

func (p *OneOfServiceGroupGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfServiceGroupGetResponseData")
}

type OneOfDirectoryServerGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *DirectoryServer       `json:"-"`
}

func NewOneOfDirectoryServerGetResponseData() *OneOfDirectoryServerGetResponseData {
	p := new(OneOfDirectoryServerGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDirectoryServerGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDirectoryServerGetResponseData is nil"))
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
	case DirectoryServer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryServer)
		}
		*p.oneOfType0 = v.(DirectoryServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDirectoryServerGetResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDirectoryServerGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(DirectoryServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.DirectoryServer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryServer)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectoryServerGetResponseData"))
}

func (p *OneOfDirectoryServerGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDirectoryServerGetResponseData")
}

type OneOfDirectoryServerListResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []DirectoryServer      `json:"-"`
}

func NewOneOfDirectoryServerListResponseData() *OneOfDirectoryServerListResponseData {
	p := new(OneOfDirectoryServerListResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDirectoryServerListResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDirectoryServerListResponseData is nil"))
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
	case []DirectoryServer:
		p.oneOfType0 = v.([]DirectoryServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.DirectoryServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.DirectoryServer>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDirectoryServerListResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.DirectoryServer>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfDirectoryServerListResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]DirectoryServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.DirectoryServer" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.DirectoryServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.DirectoryServer>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectoryServerListResponseData"))
}

func (p *OneOfDirectoryServerListResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.DirectoryServer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDirectoryServerListResponseData")
}

type OneOfDsCategoryMappingDeleteResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDsCategoryMappingDeleteResponseData() *OneOfDsCategoryMappingDeleteResponseData {
	p := new(OneOfDsCategoryMappingDeleteResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDsCategoryMappingDeleteResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDsCategoryMappingDeleteResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDsCategoryMappingDeleteResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDsCategoryMappingDeleteResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDsCategoryMappingDeleteResponseData"))
}

func (p *OneOfDsCategoryMappingDeleteResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDsCategoryMappingDeleteResponseData")
}

type OneOfNetworkSecurityPolicyGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *NetworkSecurityPolicy `json:"-"`
}

func NewOneOfNetworkSecurityPolicyGetResponseData() *OneOfNetworkSecurityPolicyGetResponseData {
	p := new(OneOfNetworkSecurityPolicyGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNetworkSecurityPolicyGetResponseData is nil"))
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
	case NetworkSecurityPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(NetworkSecurityPolicy)
		}
		*p.oneOfType0 = v.(NetworkSecurityPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(NetworkSecurityPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.NetworkSecurityPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(NetworkSecurityPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyGetResponseData"))
}

func (p *OneOfNetworkSecurityPolicyGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyGetResponseData")
}

type OneOfTierGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *TierResponse          `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfTierGetResponseData() *OneOfTierGetResponseData {
	p := new(OneOfTierGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTierGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTierGetResponseData is nil"))
	}
	switch v.(type) {
	case TierResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TierResponse)
		}
		*p.oneOfType0 = v.(TierResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTierGetResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTierGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(TierResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.TierResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TierResponse)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTierGetResponseData"))
}

func (p *OneOfTierGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTierGetResponseData")
}

type OneOfAddressGroupListResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []AddressGroup         `json:"-"`
}

func NewOneOfAddressGroupListResponseData() *OneOfAddressGroupListResponseData {
	p := new(OneOfAddressGroupListResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddressGroupListResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddressGroupListResponseData is nil"))
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
	case []AddressGroup:
		p.oneOfType0 = v.([]AddressGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.AddressGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.AddressGroup>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddressGroupListResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.AddressGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfAddressGroupListResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]AddressGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.AddressGroup" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.AddressGroup>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.AddressGroup>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddressGroupListResponseData"))
}

func (p *OneOfAddressGroupListResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.AddressGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAddressGroupListResponseData")
}

type OneOfBannerGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *BannerResponse        `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfBannerGetResponseData() *OneOfBannerGetResponseData {
	p := new(OneOfBannerGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBannerGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBannerGetResponseData is nil"))
	}
	switch v.(type) {
	case BannerResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(BannerResponse)
		}
		*p.oneOfType0 = v.(BannerResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBannerGetResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfBannerGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(BannerResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.BannerResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(BannerResponse)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBannerGetResponseData"))
}

func (p *OneOfBannerGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfBannerGetResponseData")
}

type OneOfServiceGroupPolicyAssociationResponseData struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType400  *import3.ErrorResponse        `json:"-"`
	oneOfType0    []ServiceGroupPolicyReference `json:"-"`
}

func NewOneOfServiceGroupPolicyAssociationResponseData() *OneOfServiceGroupPolicyAssociationResponseData {
	p := new(OneOfServiceGroupPolicyAssociationResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfServiceGroupPolicyAssociationResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfServiceGroupPolicyAssociationResponseData is nil"))
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
	case []ServiceGroupPolicyReference:
		p.oneOfType0 = v.([]ServiceGroupPolicyReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.ServiceGroupPolicyReference>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.ServiceGroupPolicyReference>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfServiceGroupPolicyAssociationResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.ServiceGroupPolicyReference>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfServiceGroupPolicyAssociationResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ServiceGroupPolicyReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "microseg.v4.config.ServiceGroupPolicyReference" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.ServiceGroupPolicyReference>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.ServiceGroupPolicyReference>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceGroupPolicyAssociationResponseData"))
}

func (p *OneOfServiceGroupPolicyAssociationResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.ServiceGroupPolicyReference>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfServiceGroupPolicyAssociationResponseData")
}

type OneOfDirectoryServerDeleteResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDirectoryServerDeleteResponseData() *OneOfDirectoryServerDeleteResponseData {
	p := new(OneOfDirectoryServerDeleteResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDirectoryServerDeleteResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDirectoryServerDeleteResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDirectoryServerDeleteResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDirectoryServerDeleteResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDirectoryServerDeleteResponseData"))
}

func (p *OneOfDirectoryServerDeleteResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDirectoryServerDeleteResponseData")
}

type OneOfDsCategoryMappingGetResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *CategoryMapping       `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDsCategoryMappingGetResponseData() *OneOfDsCategoryMappingGetResponseData {
	p := new(OneOfDsCategoryMappingGetResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDsCategoryMappingGetResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDsCategoryMappingGetResponseData is nil"))
	}
	switch v.(type) {
	case CategoryMapping:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CategoryMapping)
		}
		*p.oneOfType0 = v.(CategoryMapping)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDsCategoryMappingGetResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDsCategoryMappingGetResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(CategoryMapping)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.CategoryMapping" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CategoryMapping)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDsCategoryMappingGetResponseData"))
}

func (p *OneOfDsCategoryMappingGetResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDsCategoryMappingGetResponseData")
}

type OneOfServiceGroupTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfServiceGroupTaskResponseData() *OneOfServiceGroupTaskResponseData {
	p := new(OneOfServiceGroupTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfServiceGroupTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfServiceGroupTaskResponseData is nil"))
	}
	switch v.(type) {
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfServiceGroupTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfServiceGroupTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfServiceGroupTaskResponseData"))
}

func (p *OneOfServiceGroupTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfServiceGroupTaskResponseData")
}
