/*
 * Generated file models/microseg/v4/config/config_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Microseg Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	import4 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/prism/v4/config"
	"time"
)

/*
A mapping to an object in Active Directory.
*/
type AdInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The ExtID of the directory service that will be used for mapping.
	*/
	DirectoryServiceReference *string `json:"directoryServiceReference"`
	/*
	  The objectGUID for the object in AD.
	*/
	ObjectIdentifier *string `json:"objectIdentifier"`
	/*
	  The path for the mapped object in AD.
	*/
	ObjectPath *string `json:"objectPath,omitempty"`

	Status *AdStatus `json:"status,omitempty"`
}

func (p *AdInfo) MarshalJSON() ([]byte, error) {
	type AdInfoProxy AdInfo
	return json.Marshal(struct {
		*AdInfoProxy
		DirectoryServiceReference *string `json:"directoryServiceReference,omitempty"`
		ObjectIdentifier          *string `json:"objectIdentifier,omitempty"`
	}{
		AdInfoProxy:               (*AdInfoProxy)(p),
		DirectoryServiceReference: p.DirectoryServiceReference,
		ObjectIdentifier:          p.ObjectIdentifier,
	})
}

func NewAdInfo() *AdInfo {
	p := new(AdInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.AdInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The mapping status of AD Mapping.
*/
type AdStatus int

const (
	ADSTATUS_UNKNOWN                  AdStatus = 0
	ADSTATUS_REDACTED                 AdStatus = 1
	ADSTATUS_USABLE                   AdStatus = 2
	ADSTATUS_DELETED                  AdStatus = 3
	ADSTATUS_DIRECTORY_NOT_CONFIGURED AdStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AdStatus) name(index int) string {
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
func (e AdStatus) GetName() string {
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
func (e *AdStatus) index(name string) AdStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USABLE",
		"DELETED",
		"DIRECTORY_NOT_CONFIGURED",
	}
	for idx := range names {
		if names[idx] == name {
			return AdStatus(idx)
		}
	}
	return ADSTATUS_UNKNOWN
}

func (e *AdStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AdStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AdStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AdStatus) Ref() *AdStatus {
	return &e
}

type AddressGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A user defined annotation for an Address Group.
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
	  List of CIDR blocks in the Address Group.
	*/
	Ipv4Addresses []import1.IPv4Address `json:"ipv4Addresses,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A short identifier for an Address Group.
	*/
	Name *string `json:"name"`
	/*
	  Reference to policy associated with Address Group.
	*/
	PolicyReferences []string `json:"policyReferences,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A specification to how allow mode traffic should be applied, either ALL or NONE.
*/
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
A rule for specifying allowed traffic for an application.
*/
type ApplicationRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of address group references.
	*/
	DestAddressGroupReferences []string `json:"destAddressGroupReferences,omitempty"`

	DestAllowSpec *AllowType `json:"destAllowSpec,omitempty"`
	/*
	  List of categories that define a set of network endpoints as outbound.
	*/
	DestCategoryReferences []string `json:"destCategoryReferences,omitempty"`

	DestSubnet *import1.IPv4Address `json:"destSubnet,omitempty"`
	/*
	  Icmp Type Code List.
	*/
	IcmpServices []IcmpTypeCodeSpec `json:"icmpServices,omitempty"`
	/*
	  Denotes if rule allows traffic for all protocol.
	*/
	IsAllProtocolAllowed *bool `json:"isAllProtocolAllowed,omitempty"`
	/*
	  A reference to the network function chain in the rule.
	*/
	NetworkFunctionChainReference *string `json:"networkFunctionChainReference,omitempty"`
	/*
	  A set of network endpoints which is protected by a Network Security Policy and defined as a list of categories.
	*/
	SecuredGroupCategoryReferences []string `json:"securedGroupCategoryReferences"`

	ServiceGroupReferences []string `json:"serviceGroupReferences,omitempty"`
	/*
	  A list of address group references.
	*/
	SrcAddressGroupReferences []string `json:"srcAddressGroupReferences,omitempty"`

	SrcAllowSpec *AllowType `json:"srcAllowSpec,omitempty"`
	/*
	  List of categories that define a set of network endpoints as inbound.
	*/
	SrcCategoryReferences []string `json:"srcCategoryReferences,omitempty"`

	SrcSubnet *import1.IPv4Address `json:"srcSubnet,omitempty"`

	TcpServices []TcpPortRangeSpec `json:"tcpServices,omitempty"`

	UdpServices []UdpPortRangeSpec `json:"udpServices,omitempty"`
}

func (p *ApplicationRuleSpec) MarshalJSON() ([]byte, error) {
	type ApplicationRuleSpecProxy ApplicationRuleSpec
	return json.Marshal(struct {
		*ApplicationRuleSpecProxy
		SecuredGroupCategoryReferences []string `json:"securedGroupCategoryReferences,omitempty"`
	}{
		ApplicationRuleSpecProxy:       (*ApplicationRuleSpecProxy)(p),
		SecuredGroupCategoryReferences: p.SecuredGroupCategoryReferences,
	})
}

func NewApplicationRuleSpec() *ApplicationRuleSpec {
	p := new(ApplicationRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ApplicationRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CategoryMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AdInfo *AdInfo `json:"adInfo"`
	/*
	  The name for the category that this mapping is for.
	*/
	CategoryName *string `json:"categoryName"`
	/*
	  The value for the category that this mapping is for.
	*/
	CategoryValue *string `json:"categoryValue"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of Category Mapping.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CategoryMapping) MarshalJSON() ([]byte, error) {
	type CategoryMappingProxy CategoryMapping
	return json.Marshal(struct {
		*CategoryMappingProxy
		AdInfo        *AdInfo `json:"adInfo,omitempty"`
		CategoryName  *string `json:"categoryName,omitempty"`
		CategoryValue *string `json:"categoryValue,omitempty"`
		Name          *string `json:"name,omitempty"`
	}{
		CategoryMappingProxy: (*CategoryMappingProxy)(p),
		AdInfo:               p.AdInfo,
		CategoryName:         p.CategoryName,
		CategoryValue:        p.CategoryValue,
		Name:                 p.Name,
	})
}

func NewCategoryMapping() *CategoryMapping {
	p := new(CategoryMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CategoryMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.b1/config/address-groups Post operation
*/
type CreateAddressGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateAddressGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateAddressGroupApiResponse() *CreateAddressGroupApiResponse {
	p := new(CreateAddressGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateAddressGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateAddressGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateAddressGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateAddressGroupApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/directory-server-configs Post operation
*/
type CreateDirectoryServerConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDirectoryServerConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateDirectoryServerConfigApiResponse() *CreateDirectoryServerConfigApiResponse {
	p := new(CreateDirectoryServerConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateDirectoryServerConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDirectoryServerConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDirectoryServerConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDirectoryServerConfigApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/category-mappings Post operation
*/
type CreateDsCategoryMappingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDsCategoryMappingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateDsCategoryMappingApiResponse() *CreateDsCategoryMappingApiResponse {
	p := new(CreateDsCategoryMappingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateDsCategoryMappingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDsCategoryMappingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDsCategoryMappingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDsCategoryMappingApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies Post operation
*/
type CreateNetworkSecurityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNetworkSecurityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateNetworkSecurityPolicyApiResponse() *CreateNetworkSecurityPolicyApiResponse {
	p := new(CreateNetworkSecurityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateNetworkSecurityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNetworkSecurityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNetworkSecurityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNetworkSecurityPolicyApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/$actions/prepare-export Post operation
*/
type CreateNetworkSecurityPolicyExportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNetworkSecurityPolicyExportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateNetworkSecurityPolicyExportApiResponse() *CreateNetworkSecurityPolicyExportApiResponse {
	p := new(CreateNetworkSecurityPolicyExportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateNetworkSecurityPolicyExportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNetworkSecurityPolicyExportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNetworkSecurityPolicyExportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNetworkSecurityPolicyExportApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/$actions/import Post operation
*/
type CreateNetworkSecurityPolicyImportApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateNetworkSecurityPolicyImportApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateNetworkSecurityPolicyImportApiResponse() *CreateNetworkSecurityPolicyImportApiResponse {
	p := new(CreateNetworkSecurityPolicyImportApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateNetworkSecurityPolicyImportApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateNetworkSecurityPolicyImportApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateNetworkSecurityPolicyImportApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateNetworkSecurityPolicyImportApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/service-groups Post operation
*/
type CreateServiceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateServiceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateServiceGroupApiResponse() *CreateServiceGroupApiResponse {
	p := new(CreateServiceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.CreateServiceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateServiceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateServiceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateServiceGroupApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/address-groups/{extId} Delete operation
*/
type DeleteAddressGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteAddressGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteAddressGroupApiResponse() *DeleteAddressGroupApiResponse {
	p := new(DeleteAddressGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DeleteAddressGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteAddressGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteAddressGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteAddressGroupApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/directory-server-configs/{extId} Delete operation
*/
type DeleteDirectoryServerConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDirectoryServerConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteDirectoryServerConfigApiResponse() *DeleteDirectoryServerConfigApiResponse {
	p := new(DeleteDirectoryServerConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DeleteDirectoryServerConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDirectoryServerConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDirectoryServerConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDirectoryServerConfigApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/category-mappings/{extId} Delete operation
*/
type DeleteDsCategoryMappingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDsCategoryMappingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteDsCategoryMappingApiResponse() *DeleteDsCategoryMappingApiResponse {
	p := new(DeleteDsCategoryMappingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DeleteDsCategoryMappingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDsCategoryMappingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDsCategoryMappingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDsCategoryMappingApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/{extId} Delete operation
*/
type DeleteNetworkSecurityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteNetworkSecurityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteNetworkSecurityPolicyApiResponse() *DeleteNetworkSecurityPolicyApiResponse {
	p := new(DeleteNetworkSecurityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DeleteNetworkSecurityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteNetworkSecurityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteNetworkSecurityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteNetworkSecurityPolicyApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/service-groups/{extId} Delete operation
*/
type DeleteServiceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteServiceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteServiceGroupApiResponse() *DeleteServiceGroupApiResponse {
	p := new(DeleteServiceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DeleteServiceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteServiceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteServiceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteServiceGroupApiResponseData()
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

type DirectoryServerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The ExtID of the directory service that will be used for mapping.
	*/
	DirectoryServiceReference *string `json:"directoryServiceReference,omitempty"`
	/*
	  List of domain controllers to be used for event scraping.
	*/
	DomainControllers []import1.IPAddressOrFQDN `json:"domainControllers,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Enablement status of the default category.
	*/
	IsDefaultCategoryEnabled *bool `json:"isDefaultCategoryEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDirectoryServerConfig() *DirectoryServerConfig {
	p := new(DirectoryServerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.DirectoryServerConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.b1/config/address-groups/{extId} Get operation
*/
type GetAddressGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAddressGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAddressGroupApiResponse() *GetAddressGroupApiResponse {
	p := new(GetAddressGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.GetAddressGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAddressGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAddressGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAddressGroupApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/directory-server-configs/{extId} Get operation
*/
type GetDirectoryServerConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDirectoryServerConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDirectoryServerConfigApiResponse() *GetDirectoryServerConfigApiResponse {
	p := new(GetDirectoryServerConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.GetDirectoryServerConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDirectoryServerConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDirectoryServerConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDirectoryServerConfigApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/category-mappings/{extId} Get operation
*/
type GetDsCategoryMappingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDsCategoryMappingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDsCategoryMappingApiResponse() *GetDsCategoryMappingApiResponse {
	p := new(GetDsCategoryMappingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.GetDsCategoryMappingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDsCategoryMappingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDsCategoryMappingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDsCategoryMappingApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/{extId} Get operation
*/
type GetNetworkSecurityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNetworkSecurityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNetworkSecurityPolicyApiResponse() *GetNetworkSecurityPolicyApiResponse {
	p := new(GetNetworkSecurityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.GetNetworkSecurityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNetworkSecurityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNetworkSecurityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNetworkSecurityPolicyApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/service-groups/{extId} Get operation
*/
type GetServiceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetServiceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetServiceGroupApiResponse() *GetServiceGroupApiResponse {
	p := new(GetServiceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.GetServiceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetServiceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetServiceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetServiceGroupApiResponseData()
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Icmp Type Code Object. This object contains values for Icmp type and Icmp code.
*/
type IcmpTypeCodeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Icmp service Code. Ignore this field if Code has to be ANY.
	*/
	Code *int `json:"code,omitempty"`
	/*
	  Set this field to true if both Type and Code is ANY.
	*/
	IsAllAllowed *bool `json:"isAllAllowed,omitempty"`
	/*
	  Icmp service Type. Ignore this field if Type has to be ANY.
	*/
	Type *int `json:"type,omitempty"`
}

func NewIcmpTypeCodeSpec() *IcmpTypeCodeSpec {
	p := new(IcmpTypeCodeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.IcmpTypeCodeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsAllAllowed = new(bool)
	*p.IsAllAllowed = false

	return p
}

/*
Network security rule import response data.
*/
type ImportEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the entity.
	*/
	Name *string `json:"name"`

	Type *ImportEntityType `json:"type"`

	UpdateType *ImportEntityUpdateType `json:"updateType"`
}

func (p *ImportEntity) MarshalJSON() ([]byte, error) {
	type ImportEntityProxy ImportEntity
	return json.Marshal(struct {
		*ImportEntityProxy
		Name       *string                 `json:"name,omitempty"`
		Type       *ImportEntityType       `json:"type,omitempty"`
		UpdateType *ImportEntityUpdateType `json:"updateType,omitempty"`
	}{
		ImportEntityProxy: (*ImportEntityProxy)(p),
		Name:              p.Name,
		Type:              p.Type,
		UpdateType:        p.UpdateType,
	})
}

func NewImportEntity() *ImportEntity {
	p := new(ImportEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ImportEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the entity.
*/
type ImportEntityType int

const (
	IMPORTENTITYTYPE_UNKNOWN                ImportEntityType = 0
	IMPORTENTITYTYPE_REDACTED               ImportEntityType = 1
	IMPORTENTITYTYPE_CATEGORY               ImportEntityType = 2
	IMPORTENTITYTYPE_NETWORK_FUNCTION_CHAIN ImportEntityType = 3
	IMPORTENTITYTYPE_POLICY                 ImportEntityType = 4
	IMPORTENTITYTYPE_ADDRESS_GROUP          ImportEntityType = 5
	IMPORTENTITYTYPE_SERVICE_GROUP          ImportEntityType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ImportEntityType) name(index int) string {
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
func (e ImportEntityType) GetName() string {
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
func (e *ImportEntityType) index(name string) ImportEntityType {
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
			return ImportEntityType(idx)
		}
	}
	return IMPORTENTITYTYPE_UNKNOWN
}

func (e *ImportEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ImportEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ImportEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ImportEntityType) Ref() *ImportEntityType {
	return &e
}

/*
Type of update of the entity.
*/
type ImportEntityUpdateType int

const (
	IMPORTENTITYUPDATETYPE_UNKNOWN  ImportEntityUpdateType = 0
	IMPORTENTITYUPDATETYPE_REDACTED ImportEntityUpdateType = 1
	IMPORTENTITYUPDATETYPE_ADD      ImportEntityUpdateType = 2
	IMPORTENTITYUPDATETYPE_DELETE   ImportEntityUpdateType = 3
	IMPORTENTITYUPDATETYPE_MODIFY   ImportEntityUpdateType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ImportEntityUpdateType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"DELETE",
		"MODIFY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ImportEntityUpdateType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"DELETE",
		"MODIFY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ImportEntityUpdateType) index(name string) ImportEntityUpdateType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADD",
		"DELETE",
		"MODIFY",
	}
	for idx := range names {
		if names[idx] == name {
			return ImportEntityUpdateType(idx)
		}
	}
	return IMPORTENTITYUPDATETYPE_UNKNOWN
}

func (e *ImportEntityUpdateType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ImportEntityUpdateType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ImportEntityUpdateType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ImportEntityUpdateType) Ref() *ImportEntityUpdateType {
	return &e
}

/*
A specification to whether traffic between intra secured group entities should be allowed or denied.
*/
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
A rule for specifying allowed traffic inside of a secured entity group.
*/
type IntraEntityGroupRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SecuredGroupAction *IntraEntityGroupRuleAction `json:"securedGroupAction"`

	SecuredGroupCategoryReferences []string `json:"securedGroupCategoryReferences"`
}

func (p *IntraEntityGroupRuleSpec) MarshalJSON() ([]byte, error) {
	type IntraEntityGroupRuleSpecProxy IntraEntityGroupRuleSpec
	return json.Marshal(struct {
		*IntraEntityGroupRuleSpecProxy
		SecuredGroupAction             *IntraEntityGroupRuleAction `json:"securedGroupAction,omitempty"`
		SecuredGroupCategoryReferences []string                    `json:"securedGroupCategoryReferences,omitempty"`
	}{
		IntraEntityGroupRuleSpecProxy:  (*IntraEntityGroupRuleSpecProxy)(p),
		SecuredGroupAction:             p.SecuredGroupAction,
		SecuredGroupCategoryReferences: p.SecuredGroupCategoryReferences,
	})
}

func NewIntraEntityGroupRuleSpec() *IntraEntityGroupRuleSpec {
	p := new(IntraEntityGroupRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.IntraEntityGroupRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.b1/config/address-groups Get operation
*/
type ListAddressGroupsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAddressGroupsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAddressGroupsApiResponse() *ListAddressGroupsApiResponse {
	p := new(ListAddressGroupsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListAddressGroupsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAddressGroupsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAddressGroupsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAddressGroupsApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/directory-server-configs Get operation
*/
type ListDirectoryServerConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDirectoryServerConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDirectoryServerConfigsApiResponse() *ListDirectoryServerConfigsApiResponse {
	p := new(ListDirectoryServerConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListDirectoryServerConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDirectoryServerConfigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDirectoryServerConfigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDirectoryServerConfigsApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/category-mappings Get operation
*/
type ListDsCategoryMappingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDsCategoryMappingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDsCategoryMappingsApiResponse() *ListDsCategoryMappingsApiResponse {
	p := new(ListDsCategoryMappingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListDsCategoryMappingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDsCategoryMappingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDsCategoryMappingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDsCategoryMappingsApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies Get operation
*/
type ListNetworkSecurityPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNetworkSecurityPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNetworkSecurityPoliciesApiResponse() *ListNetworkSecurityPoliciesApiResponse {
	p := new(ListNetworkSecurityPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListNetworkSecurityPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNetworkSecurityPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNetworkSecurityPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNetworkSecurityPoliciesApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/{extId}/rules Get operation
*/
type ListNetworkSecurityPolicyRulesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNetworkSecurityPolicyRulesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNetworkSecurityPolicyRulesApiResponse() *ListNetworkSecurityPolicyRulesApiResponse {
	p := new(ListNetworkSecurityPolicyRulesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListNetworkSecurityPolicyRulesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNetworkSecurityPolicyRulesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNetworkSecurityPolicyRulesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNetworkSecurityPolicyRulesApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/service-groups Get operation
*/
type ListServiceGroupsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListServiceGroupsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListServiceGroupsApiResponse() *ListServiceGroupsApiResponse {
	p := new(ListServiceGroupsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.ListServiceGroupsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListServiceGroupsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListServiceGroupsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListServiceGroupsApiResponseData()
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NetworkSecurityPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty"`

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
	  If Hitlog is enabled.
	*/
	IsHitlogEnabled *bool `json:"isHitlogEnabled,omitempty"`
	/*
	  If Ipv6 Traffic is allowed.
	*/
	IsIpv6TrafficAllowed *bool `json:"isIpv6TrafficAllowed,omitempty"`

	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`

	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
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

	Scope *SecurityPolicyScope `json:"scope,omitempty"`
	/*
	  Uuids of the secured groups in the NSP.
	*/
	SecuredGroups []string `json:"securedGroups,omitempty"`

	State *SecurityPolicyState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *SecurityPolicyType `json:"type"`
	/*
	  A list of external ids for VPCs, used only when the scope of policy is a list of VPCs.
	*/
	VpcReferences []string `json:"vpcReferences,omitempty"`
}

func (p *NetworkSecurityPolicy) MarshalJSON() ([]byte, error) {
	type NetworkSecurityPolicyProxy NetworkSecurityPolicy
	return json.Marshal(struct {
		*NetworkSecurityPolicyProxy
		Name *string             `json:"name,omitempty"`
		Type *SecurityPolicyType `json:"type,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`

	Spec *OneOfNetworkSecurityPolicyRuleSpec `json:"spec"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

/*
Defines the scope of the policy. Currently, only ALL_VLAN and VPC_LIST are supported. If scope is not provided, the default is set based on whether vpcReferences field is provided or not.
*/
type SecurityPolicyScope int

const (
	SECURITYPOLICYSCOPE_UNKNOWN  SecurityPolicyScope = 0
	SECURITYPOLICYSCOPE_REDACTED SecurityPolicyScope = 1
	SECURITYPOLICYSCOPE_ALL_VLAN SecurityPolicyScope = 2
	SECURITYPOLICYSCOPE_ALL_VPC  SecurityPolicyScope = 3
	SECURITYPOLICYSCOPE_VPC_LIST SecurityPolicyScope = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SecurityPolicyScope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL_VLAN",
		"ALL_VPC",
		"VPC_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SecurityPolicyScope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL_VLAN",
		"ALL_VPC",
		"VPC_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SecurityPolicyScope) index(name string) SecurityPolicyScope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL_VLAN",
		"ALL_VPC",
		"VPC_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return SecurityPolicyScope(idx)
		}
	}
	return SECURITYPOLICYSCOPE_UNKNOWN
}

func (e *SecurityPolicyScope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SecurityPolicyScope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SecurityPolicyScope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SecurityPolicyScope) Ref() *SecurityPolicyScope {
	return &e
}

/*
Whether the policy is applied or monitored; can be omitted or set null to save the policy without applying or monitoring it.
*/
type SecurityPolicyState int

const (
	SECURITYPOLICYSTATE_UNKNOWN  SecurityPolicyState = 0
	SECURITYPOLICYSTATE_REDACTED SecurityPolicyState = 1
	SECURITYPOLICYSTATE_SAVE     SecurityPolicyState = 2
	SECURITYPOLICYSTATE_MONITOR  SecurityPolicyState = 3
	SECURITYPOLICYSTATE_ENFORCE  SecurityPolicyState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SecurityPolicyState) name(index int) string {
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
func (e SecurityPolicyState) GetName() string {
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
func (e *SecurityPolicyState) index(name string) SecurityPolicyState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAVE",
		"MONITOR",
		"ENFORCE",
	}
	for idx := range names {
		if names[idx] == name {
			return SecurityPolicyState(idx)
		}
	}
	return SECURITYPOLICYSTATE_UNKNOWN
}

func (e *SecurityPolicyState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SecurityPolicyState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SecurityPolicyState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SecurityPolicyState) Ref() *SecurityPolicyState {
	return &e
}

/*
Defines the type of rules that can be used in a policy.
*/
type SecurityPolicyType int

const (
	SECURITYPOLICYTYPE_UNKNOWN     SecurityPolicyType = 0
	SECURITYPOLICYTYPE_REDACTED    SecurityPolicyType = 1
	SECURITYPOLICYTYPE_QUARANTINE  SecurityPolicyType = 2
	SECURITYPOLICYTYPE_ISOLATION   SecurityPolicyType = 3
	SECURITYPOLICYTYPE_APPLICATION SecurityPolicyType = 4
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

type ServiceGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A user defined annotation for a Service Group.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A short identifier for a Service Group.
	*/
	Name *string `json:"name"`
	/*
	  Reference to policy associated with Service Group.
	*/
	PolicyReferences []string `json:"policyReferences,omitempty"`
	/*
	  List of TCP ports in the service.
	*/
	TcpServices []TcpPortRangeSpec `json:"tcpServices,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A rule for specifying that two environments should be isolated from each other. Specify both 'firstIsolationGroup' and 'secondIsolationGroup'.
*/
type TwoEnvIsolationRuleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Denotes the first group of category uuids that will be used in an isolation policy.
	*/
	FirstIsolationGroup []string `json:"firstIsolationGroup"`
	/*
	  Denotes the second group of category uuids that will be used in an isolation policy.
	*/
	SecondIsolationGroup []string `json:"secondIsolationGroup"`
}

func (p *TwoEnvIsolationRuleSpec) MarshalJSON() ([]byte, error) {
	type TwoEnvIsolationRuleSpecProxy TwoEnvIsolationRuleSpec
	return json.Marshal(struct {
		*TwoEnvIsolationRuleSpecProxy
		FirstIsolationGroup  []string `json:"firstIsolationGroup,omitempty"`
		SecondIsolationGroup []string `json:"secondIsolationGroup,omitempty"`
	}{
		TwoEnvIsolationRuleSpecProxy: (*TwoEnvIsolationRuleSpecProxy)(p),
		FirstIsolationGroup:          p.FirstIsolationGroup,
		SecondIsolationGroup:         p.SecondIsolationGroup,
	})
}

func NewTwoEnvIsolationRuleSpec() *TwoEnvIsolationRuleSpec {
	p := new(TwoEnvIsolationRuleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.TwoEnvIsolationRuleSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /microseg/v4.0.b1/config/address-groups/{extId} Put operation
*/
type UpdateAddressGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateAddressGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateAddressGroupApiResponse() *UpdateAddressGroupApiResponse {
	p := new(UpdateAddressGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UpdateAddressGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateAddressGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateAddressGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateAddressGroupApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/directory-server-configs/{extId} Put operation
*/
type UpdateDirectoryServerConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDirectoryServerConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateDirectoryServerConfigApiResponse() *UpdateDirectoryServerConfigApiResponse {
	p := new(UpdateDirectoryServerConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UpdateDirectoryServerConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDirectoryServerConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDirectoryServerConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDirectoryServerConfigApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/category-mappings/{extId} Put operation
*/
type UpdateDsCategoryMappingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDsCategoryMappingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateDsCategoryMappingApiResponse() *UpdateDsCategoryMappingApiResponse {
	p := new(UpdateDsCategoryMappingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UpdateDsCategoryMappingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDsCategoryMappingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDsCategoryMappingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDsCategoryMappingApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/policies/{extId} Put operation
*/
type UpdateNetworkSecurityPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNetworkSecurityPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateNetworkSecurityPolicyApiResponse() *UpdateNetworkSecurityPolicyApiResponse {
	p := new(UpdateNetworkSecurityPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UpdateNetworkSecurityPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNetworkSecurityPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNetworkSecurityPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNetworkSecurityPolicyApiResponseData()
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
REST response for all response codes in API path /microseg/v4.0.b1/config/service-groups/{extId} Put operation
*/
type UpdateServiceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateServiceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateServiceGroupApiResponse() *UpdateServiceGroupApiResponse {
	p := new(UpdateServiceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "microseg.v4.config.UpdateServiceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateServiceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateServiceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateServiceGroupApiResponseData()
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

type OneOfCreateServiceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateServiceGroupApiResponseData() *OneOfCreateServiceGroupApiResponseData {
	p := new(OneOfCreateServiceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateServiceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateServiceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateServiceGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateServiceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateServiceGroupApiResponseData"))
}

func (p *OneOfCreateServiceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateServiceGroupApiResponseData")
}

type OneOfGetNetworkSecurityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *NetworkSecurityPolicy `json:"-"`
}

func NewOneOfGetNetworkSecurityPolicyApiResponseData() *OneOfGetNetworkSecurityPolicyApiResponseData {
	p := new(OneOfGetNetworkSecurityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNetworkSecurityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNetworkSecurityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetNetworkSecurityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetNetworkSecurityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNetworkSecurityPolicyApiResponseData"))
}

func (p *OneOfGetNetworkSecurityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetNetworkSecurityPolicyApiResponseData")
}

type OneOfUpdateDirectoryServerConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateDirectoryServerConfigApiResponseData() *OneOfUpdateDirectoryServerConfigApiResponseData {
	p := new(OneOfUpdateDirectoryServerConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDirectoryServerConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDirectoryServerConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfUpdateDirectoryServerConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateDirectoryServerConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDirectoryServerConfigApiResponseData"))
}

func (p *OneOfUpdateDirectoryServerConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDirectoryServerConfigApiResponseData")
}

type OneOfDeleteDsCategoryMappingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfDeleteDsCategoryMappingApiResponseData() *OneOfDeleteDsCategoryMappingApiResponseData {
	p := new(OneOfDeleteDsCategoryMappingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDsCategoryMappingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDsCategoryMappingApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfDeleteDsCategoryMappingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteDsCategoryMappingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDsCategoryMappingApiResponseData"))
}

func (p *OneOfDeleteDsCategoryMappingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDsCategoryMappingApiResponseData")
}

type OneOfGetDsCategoryMappingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *CategoryMapping       `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetDsCategoryMappingApiResponseData() *OneOfGetDsCategoryMappingApiResponseData {
	p := new(OneOfGetDsCategoryMappingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDsCategoryMappingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDsCategoryMappingApiResponseData is nil"))
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
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetDsCategoryMappingApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetDsCategoryMappingApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDsCategoryMappingApiResponseData"))
}

func (p *OneOfGetDsCategoryMappingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetDsCategoryMappingApiResponseData")
}

type OneOfCreateNetworkSecurityPolicyExportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateNetworkSecurityPolicyExportApiResponseData() *OneOfCreateNetworkSecurityPolicyExportApiResponseData {
	p := new(OneOfCreateNetworkSecurityPolicyExportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNetworkSecurityPolicyExportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNetworkSecurityPolicyExportApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateNetworkSecurityPolicyExportApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateNetworkSecurityPolicyExportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNetworkSecurityPolicyExportApiResponseData"))
}

func (p *OneOfCreateNetworkSecurityPolicyExportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNetworkSecurityPolicyExportApiResponseData")
}

type OneOfUpdateServiceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateServiceGroupApiResponseData() *OneOfUpdateServiceGroupApiResponseData {
	p := new(OneOfUpdateServiceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateServiceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateServiceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfUpdateServiceGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateServiceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateServiceGroupApiResponseData"))
}

func (p *OneOfUpdateServiceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateServiceGroupApiResponseData")
}

type OneOfUpdateDsCategoryMappingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateDsCategoryMappingApiResponseData() *OneOfUpdateDsCategoryMappingApiResponseData {
	p := new(OneOfUpdateDsCategoryMappingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDsCategoryMappingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDsCategoryMappingApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfUpdateDsCategoryMappingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateDsCategoryMappingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDsCategoryMappingApiResponseData"))
}

func (p *OneOfUpdateDsCategoryMappingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDsCategoryMappingApiResponseData")
}

type OneOfDeleteDirectoryServerConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfDeleteDirectoryServerConfigApiResponseData() *OneOfDeleteDirectoryServerConfigApiResponseData {
	p := new(OneOfDeleteDirectoryServerConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDirectoryServerConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDirectoryServerConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfDeleteDirectoryServerConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteDirectoryServerConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDirectoryServerConfigApiResponseData"))
}

func (p *OneOfDeleteDirectoryServerConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDirectoryServerConfigApiResponseData")
}

type OneOfListDirectoryServerConfigsApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import4.ErrorResponse  `json:"-"`
	oneOfType0    []DirectoryServerConfig `json:"-"`
}

func NewOneOfListDirectoryServerConfigsApiResponseData() *OneOfListDirectoryServerConfigsApiResponseData {
	p := new(OneOfListDirectoryServerConfigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDirectoryServerConfigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDirectoryServerConfigsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []DirectoryServerConfig:
		p.oneOfType0 = v.([]DirectoryServerConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.DirectoryServerConfig>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.DirectoryServerConfig>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListDirectoryServerConfigsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.DirectoryServerConfig>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListDirectoryServerConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new([]DirectoryServerConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "microseg.v4.config.DirectoryServerConfig" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.DirectoryServerConfig>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.DirectoryServerConfig>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDirectoryServerConfigsApiResponseData"))
}

func (p *OneOfListDirectoryServerConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.DirectoryServerConfig>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListDirectoryServerConfigsApiResponseData")
}

type OneOfGetDirectoryServerConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *DirectoryServerConfig `json:"-"`
}

func NewOneOfGetDirectoryServerConfigApiResponseData() *OneOfGetDirectoryServerConfigApiResponseData {
	p := new(OneOfGetDirectoryServerConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDirectoryServerConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDirectoryServerConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case DirectoryServerConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryServerConfig)
		}
		*p.oneOfType0 = v.(DirectoryServerConfig)
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

func (p *OneOfGetDirectoryServerConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetDirectoryServerConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(DirectoryServerConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.DirectoryServerConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryServerConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDirectoryServerConfigApiResponseData"))
}

func (p *OneOfGetDirectoryServerConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetDirectoryServerConfigApiResponseData")
}

type OneOfDeleteNetworkSecurityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfDeleteNetworkSecurityPolicyApiResponseData() *OneOfDeleteNetworkSecurityPolicyApiResponseData {
	p := new(OneOfDeleteNetworkSecurityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteNetworkSecurityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteNetworkSecurityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfDeleteNetworkSecurityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteNetworkSecurityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteNetworkSecurityPolicyApiResponseData"))
}

func (p *OneOfDeleteNetworkSecurityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteNetworkSecurityPolicyApiResponseData")
}

type OneOfNetworkSecurityPolicyRuleSpec struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    *TwoEnvIsolationRuleSpec  `json:"-"`
	oneOfType2    *IntraEntityGroupRuleSpec `json:"-"`
	oneOfType1    *ApplicationRuleSpec      `json:"-"`
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
	case TwoEnvIsolationRuleSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TwoEnvIsolationRuleSpec)
		}
		*p.oneOfType0 = v.(TwoEnvIsolationRuleSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case IntraEntityGroupRuleSpec:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(IntraEntityGroupRuleSpec)
		}
		*p.oneOfType2 = v.(IntraEntityGroupRuleSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case ApplicationRuleSpec:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApplicationRuleSpec)
		}
		*p.oneOfType1 = v.(ApplicationRuleSpec)
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

func (p *OneOfNetworkSecurityPolicyRuleSpec) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(TwoEnvIsolationRuleSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "microseg.v4.config.TwoEnvIsolationRuleSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TwoEnvIsolationRuleSpec)
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
	vOneOfType2 := new(IntraEntityGroupRuleSpec)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "microseg.v4.config.IntraEntityGroupRuleSpec" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(IntraEntityGroupRuleSpec)
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
	vOneOfType1 := new(ApplicationRuleSpec)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "microseg.v4.config.ApplicationRuleSpec" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApplicationRuleSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNetworkSecurityPolicyRuleSpec"))
}

func (p *OneOfNetworkSecurityPolicyRuleSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfNetworkSecurityPolicyRuleSpec")
}

type OneOfCreateNetworkSecurityPolicyImportApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType1    []ImportEntity         `json:"-"`
}

func NewOneOfCreateNetworkSecurityPolicyImportApiResponseData() *OneOfCreateNetworkSecurityPolicyImportApiResponseData {
	p := new(OneOfCreateNetworkSecurityPolicyImportApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNetworkSecurityPolicyImportApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNetworkSecurityPolicyImportApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case []ImportEntity:
		p.oneOfType1 = v.([]ImportEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<microseg.v4.config.ImportEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<microseg.v4.config.ImportEntity>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateNetworkSecurityPolicyImportApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if "List<microseg.v4.config.ImportEntity>" == *p.Discriminator {
		return p.oneOfType1
	}
	return nil
}

func (p *OneOfCreateNetworkSecurityPolicyImportApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType1 := new([]ImportEntity)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {

		if len(*vOneOfType1) == 0 || "microseg.v4.config.ImportEntity" == *((*vOneOfType1)[0].ObjectType_) {
			p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<microseg.v4.config.ImportEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<microseg.v4.config.ImportEntity>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNetworkSecurityPolicyImportApiResponseData"))
}

func (p *OneOfCreateNetworkSecurityPolicyImportApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<microseg.v4.config.ImportEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNetworkSecurityPolicyImportApiResponseData")
}

type OneOfCreateDsCategoryMappingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateDsCategoryMappingApiResponseData() *OneOfCreateDsCategoryMappingApiResponseData {
	p := new(OneOfCreateDsCategoryMappingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDsCategoryMappingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDsCategoryMappingApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateDsCategoryMappingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateDsCategoryMappingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDsCategoryMappingApiResponseData"))
}

func (p *OneOfCreateDsCategoryMappingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDsCategoryMappingApiResponseData")
}

type OneOfCreateAddressGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateAddressGroupApiResponseData() *OneOfCreateAddressGroupApiResponseData {
	p := new(OneOfCreateAddressGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateAddressGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateAddressGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateAddressGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateAddressGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAddressGroupApiResponseData"))
}

func (p *OneOfCreateAddressGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateAddressGroupApiResponseData")
}

type OneOfGetAddressGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *AddressGroup          `json:"-"`
}

func NewOneOfGetAddressGroupApiResponseData() *OneOfGetAddressGroupApiResponseData {
	p := new(OneOfGetAddressGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAddressGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAddressGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAddressGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAddressGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAddressGroupApiResponseData"))
}

func (p *OneOfGetAddressGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAddressGroupApiResponseData")
}

type OneOfDeleteAddressGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfDeleteAddressGroupApiResponseData() *OneOfDeleteAddressGroupApiResponseData {
	p := new(OneOfDeleteAddressGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteAddressGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteAddressGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfDeleteAddressGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteAddressGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAddressGroupApiResponseData"))
}

func (p *OneOfDeleteAddressGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteAddressGroupApiResponseData")
}

type OneOfListNetworkSecurityPoliciesApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import4.ErrorResponse  `json:"-"`
	oneOfType0    []NetworkSecurityPolicy `json:"-"`
	oneOfType1    *FileDetail             `json:"-"`
}

func NewOneOfListNetworkSecurityPoliciesApiResponseData() *OneOfListNetworkSecurityPoliciesApiResponseData {
	p := new(OneOfListNetworkSecurityPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNetworkSecurityPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNetworkSecurityPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	case FileDetail:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(FileDetail)
		}
		*p.oneOfType1 = v.(FileDetail)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNetworkSecurityPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.NetworkSecurityPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType1 != nil && "FileDetail" == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfListNetworkSecurityPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType1 := new(FileDetail)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(FileDetail)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNetworkSecurityPoliciesApiResponseData"))
}

func (p *OneOfListNetworkSecurityPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.NetworkSecurityPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && "FileDetail" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfListNetworkSecurityPoliciesApiResponseData")
}

type OneOfDeleteServiceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfDeleteServiceGroupApiResponseData() *OneOfDeleteServiceGroupApiResponseData {
	p := new(OneOfDeleteServiceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteServiceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteServiceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfDeleteServiceGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteServiceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteServiceGroupApiResponseData"))
}

func (p *OneOfDeleteServiceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteServiceGroupApiResponseData")
}

type OneOfUpdateAddressGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateAddressGroupApiResponseData() *OneOfUpdateAddressGroupApiResponseData {
	p := new(OneOfUpdateAddressGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateAddressGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateAddressGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfUpdateAddressGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateAddressGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAddressGroupApiResponseData"))
}

func (p *OneOfUpdateAddressGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateAddressGroupApiResponseData")
}

type OneOfUpdateNetworkSecurityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfUpdateNetworkSecurityPolicyApiResponseData() *OneOfUpdateNetworkSecurityPolicyApiResponseData {
	p := new(OneOfUpdateNetworkSecurityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNetworkSecurityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNetworkSecurityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfUpdateNetworkSecurityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateNetworkSecurityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNetworkSecurityPolicyApiResponseData"))
}

func (p *OneOfUpdateNetworkSecurityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNetworkSecurityPolicyApiResponseData")
}

type OneOfListNetworkSecurityPolicyRulesApiResponseData struct {
	Discriminator *string                     `json:"-"`
	ObjectType_   *string                     `json:"-"`
	oneOfType400  *import4.ErrorResponse      `json:"-"`
	oneOfType0    []NetworkSecurityPolicyRule `json:"-"`
}

func NewOneOfListNetworkSecurityPolicyRulesApiResponseData() *OneOfListNetworkSecurityPolicyRulesApiResponseData {
	p := new(OneOfListNetworkSecurityPolicyRulesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNetworkSecurityPolicyRulesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNetworkSecurityPolicyRulesApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNetworkSecurityPolicyRulesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.NetworkSecurityPolicyRule>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListNetworkSecurityPolicyRulesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNetworkSecurityPolicyRulesApiResponseData"))
}

func (p *OneOfListNetworkSecurityPolicyRulesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.NetworkSecurityPolicyRule>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListNetworkSecurityPolicyRulesApiResponseData")
}

type OneOfGetServiceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *ServiceGroup          `json:"-"`
}

func NewOneOfGetServiceGroupApiResponseData() *OneOfGetServiceGroupApiResponseData {
	p := new(OneOfGetServiceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetServiceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetServiceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetServiceGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetServiceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetServiceGroupApiResponseData"))
}

func (p *OneOfGetServiceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetServiceGroupApiResponseData")
}

type OneOfListAddressGroupsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []AddressGroup         `json:"-"`
}

func NewOneOfListAddressGroupsApiResponseData() *OneOfListAddressGroupsApiResponseData {
	p := new(OneOfListAddressGroupsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAddressGroupsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAddressGroupsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListAddressGroupsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.AddressGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAddressGroupsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAddressGroupsApiResponseData"))
}

func (p *OneOfListAddressGroupsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.AddressGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAddressGroupsApiResponseData")
}

type OneOfCreateNetworkSecurityPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateNetworkSecurityPolicyApiResponseData() *OneOfCreateNetworkSecurityPolicyApiResponseData {
	p := new(OneOfCreateNetworkSecurityPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateNetworkSecurityPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateNetworkSecurityPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateNetworkSecurityPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateNetworkSecurityPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateNetworkSecurityPolicyApiResponseData"))
}

func (p *OneOfCreateNetworkSecurityPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateNetworkSecurityPolicyApiResponseData")
}

type OneOfListServiceGroupsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []ServiceGroup         `json:"-"`
}

func NewOneOfListServiceGroupsApiResponseData() *OneOfListServiceGroupsApiResponseData {
	p := new(OneOfListServiceGroupsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListServiceGroupsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListServiceGroupsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListServiceGroupsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.ServiceGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListServiceGroupsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListServiceGroupsApiResponseData"))
}

func (p *OneOfListServiceGroupsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.ServiceGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListServiceGroupsApiResponseData")
}

type OneOfCreateDirectoryServerConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfCreateDirectoryServerConfigApiResponseData() *OneOfCreateDirectoryServerConfigApiResponseData {
	p := new(OneOfCreateDirectoryServerConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDirectoryServerConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDirectoryServerConfigApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfCreateDirectoryServerConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateDirectoryServerConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDirectoryServerConfigApiResponseData"))
}

func (p *OneOfCreateDirectoryServerConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDirectoryServerConfigApiResponseData")
}

type OneOfListDsCategoryMappingsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    []CategoryMapping      `json:"-"`
}

func NewOneOfListDsCategoryMappingsApiResponseData() *OneOfListDsCategoryMappingsApiResponseData {
	p := new(OneOfListDsCategoryMappingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDsCategoryMappingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDsCategoryMappingsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfListDsCategoryMappingsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<microseg.v4.config.CategoryMapping>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListDsCategoryMappingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "microseg.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDsCategoryMappingsApiResponseData"))
}

func (p *OneOfListDsCategoryMappingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<microseg.v4.config.CategoryMapping>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListDsCategoryMappingsApiResponseData")
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
