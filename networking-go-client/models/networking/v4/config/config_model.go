/*
 * Generated file models/networking/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Networking APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Subnets , Virtual Switch and  VPCs amongst others
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
)

/*
Address Type like "EXTERNAL" or "ANY".
*/
type AddressType int

const (
	ADDRESSTYPE_UNKNOWN  AddressType = 0
	ADDRESSTYPE_REDACTED AddressType = 1
	ADDRESSTYPE_ANY      AddressType = 2
	ADDRESSTYPE_EXTERNAL AddressType = 3
	ADDRESSTYPE_SUBNET   AddressType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AddressType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"EXTERNAL",
		"SUBNET",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AddressType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"EXTERNAL",
		"SUBNET",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AddressType) index(name string) AddressType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"EXTERNAL",
		"SUBNET",
	}
	for idx := range names {
		if names[idx] == name {
			return AddressType(idx)
		}
	}
	return ADDRESSTYPE_UNKNOWN
}

func (e *AddressType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AddressType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AddressType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AddressType) Ref() *AddressType {
	return &e
}

/*
Address Type like "EXTERNAL" or "ANY".
*/
type AddressTypeObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressType *AddressType `json:"addressType"`

	SubnetPrefix *IPSubnet `json:"subnetPrefix,omitempty"`
}

func (p *AddressTypeObject) MarshalJSON() ([]byte, error) {
	type AddressTypeObjectProxy AddressTypeObject
	return json.Marshal(struct {
		*AddressTypeObjectProxy
		AddressType *AddressType `json:"addressType,omitempty"`
	}{
		AddressTypeObjectProxy: (*AddressTypeObjectProxy)(p),
		AddressType:            p.AddressType,
	})
}

func NewAddressTypeObject() *AddressTypeObject {
	p := new(AddressTypeObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AddressTypeObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Load balancing algorithm configured for the session.
*/
type Algorithm int

const (
	ALGORITHM_UNKNOWN         Algorithm = 0
	ALGORITHM_REDACTED        Algorithm = 1
	ALGORITHM_FIVE_TUPLE_HASH Algorithm = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Algorithm) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FIVE_TUPLE_HASH",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Algorithm) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FIVE_TUPLE_HASH",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Algorithm) index(name string) Algorithm {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FIVE_TUPLE_HASH",
	}
	for idx := range names {
		if names[idx] == name {
			return Algorithm(idx)
		}
	}
	return ALGORITHM_UNKNOWN
}

func (e *Algorithm) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Algorithm:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Algorithm) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Algorithm) Ref() *Algorithm {
	return &e
}

type Anc struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Resolve name server for the ANC services.
	*/
	AncDomainNameServerList []import1.IPAddress `json:"ancDomainNameServerList,omitempty"`
	/*
	  URL of Atlas Network Controller (ANC).
	*/
	AncUrl *string `json:"ancUrl,omitempty"`
	/*
	  Configuration version for the current ANC. It is the logical timestamp for the current V4 release and will be updated by the actual configuration version in future releases.
	*/
	ConfigVersion *int64 `json:"configVersion,omitempty"`
	/*
	  UUIDs of the clusters excluded from advanced networking.
	*/
	ExcludedClusters []string `json:"excludedClusters,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True if atlas networking is enabled and false otherwise.
	*/
	IsAtlasNetworkingEnabled *bool `json:"isAtlasNetworkingEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  OVN remote address of the Atlas Network Controller (ANC).
	*/
	OvnRemoteAddress *string `json:"ovnRemoteAddress,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAnc() *Anc {
	p := new(Anc)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Anc"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Assignment method for load balancer Virtual IP.
*/
type AssignmentType int

const (
	ASSIGNMENTTYPE_UNKNOWN  AssignmentType = 0
	ASSIGNMENTTYPE_REDACTED AssignmentType = 1
	ASSIGNMENTTYPE_DYNAMIC  AssignmentType = 2
	ASSIGNMENTTYPE_STATIC   AssignmentType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AssignmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"STATIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AssignmentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"STATIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AssignmentType) index(name string) AssignmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"STATIC",
	}
	for idx := range names {
		if names[idx] == name {
			return AssignmentType(idx)
		}
	}
	return ASSIGNMENTTYPE_UNKNOWN
}

func (e *AssignmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AssignmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AssignmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AssignmentType) Ref() *AssignmentType {
	return &e
}

/*
Authentication algorithm.
*/
type AuthenticationAlgorithm int

const (
	AUTHENTICATIONALGORITHM_UNKNOWN  AuthenticationAlgorithm = 0
	AUTHENTICATIONALGORITHM_REDACTED AuthenticationAlgorithm = 1
	AUTHENTICATIONALGORITHM_MD5      AuthenticationAlgorithm = 2
	AUTHENTICATIONALGORITHM_SHA1     AuthenticationAlgorithm = 3
	AUTHENTICATIONALGORITHM_SHA256   AuthenticationAlgorithm = 4
	AUTHENTICATIONALGORITHM_SHA384   AuthenticationAlgorithm = 5
	AUTHENTICATIONALGORITHM_SHA512   AuthenticationAlgorithm = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthenticationAlgorithm) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthenticationAlgorithm) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthenticationAlgorithm) index(name string) AuthenticationAlgorithm {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA1",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthenticationAlgorithm(idx)
		}
	}
	return AUTHENTICATIONALGORITHM_UNKNOWN
}

func (e *AuthenticationAlgorithm) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationAlgorithm:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthenticationAlgorithm) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthenticationAlgorithm) Ref() *AuthenticationAlgorithm {
	return &e
}

/*
Authentication type.
*/
type AuthenticationType int

const (
	AUTHENTICATIONTYPE_UNKNOWN    AuthenticationType = 0
	AUTHENTICATIONTYPE_REDACTED   AuthenticationType = 1
	AUTHENTICATIONTYPE_PLAIN_TEXT AuthenticationType = 2
	AUTHENTICATIONTYPE_MD5        AuthenticationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthenticationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLAIN_TEXT",
		"MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthenticationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLAIN_TEXT",
		"MD5",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthenticationType) index(name string) AuthenticationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLAIN_TEXT",
		"MD5",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthenticationType(idx)
		}
	}
	return AUTHENTICATIONTYPE_UNKNOWN
}

func (e *AuthenticationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthenticationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthenticationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthenticationType) Ref() *AuthenticationType {
	return &e
}

/*
Authorization data.
*/
type AuthorizationData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client Id.
	*/
	ClientId *string `json:"clientId,omitempty"`
	/*
	  Expiry time of the authorization token.
	*/
	ExpiryTimeSeconds *int64 `json:"expiryTimeSeconds,omitempty"`
	/*
	  Subscription ID.
	*/
	SubscriptionId *string `json:"subscriptionId"`
	/*
	  Tenant Id.
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Authorization token.
	*/
	Token *string `json:"token"`
}

func (p *AuthorizationData) MarshalJSON() ([]byte, error) {
	type AuthorizationDataProxy AuthorizationData
	return json.Marshal(struct {
		*AuthorizationDataProxy
		SubscriptionId *string `json:"subscriptionId,omitempty"`
		Token          *string `json:"token,omitempty"`
	}{
		AuthorizationDataProxy: (*AuthorizationDataProxy)(p),
		SubscriptionId:         p.SubscriptionId,
		Token:                  p.Token,
	})
}

func NewAuthorizationData() *AuthorizationData {
	p := new(AuthorizationData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AuthorizationData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
AWS configuration.
*/
type AwsConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external subnet configuration list for the AWS cloud.
	*/
	ExternalSubnetConfigList []AwsExternalSubnetConfig `json:"externalSubnetConfigList"`
}

func (p *AwsConfig) MarshalJSON() ([]byte, error) {
	type AwsConfigProxy AwsConfig
	return json.Marshal(struct {
		*AwsConfigProxy
		ExternalSubnetConfigList []AwsExternalSubnetConfig `json:"externalSubnetConfigList,omitempty"`
	}{
		AwsConfigProxy:           (*AwsConfigProxy)(p),
		ExternalSubnetConfigList: p.ExternalSubnetConfigList,
	})
}

func NewAwsConfig() *AwsConfig {
	p := new(AwsConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AwsConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The external subnet configuration for the AWS cloud.
*/
type AwsExternalSubnetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AwsSubnetConfig *AwsSubnetConfig `json:"awsSubnetConfig"`
	/*
	  Cluster
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
}

func (p *AwsExternalSubnetConfig) MarshalJSON() ([]byte, error) {
	type AwsExternalSubnetConfigProxy AwsExternalSubnetConfig
	return json.Marshal(struct {
		*AwsExternalSubnetConfigProxy
		AwsSubnetConfig *AwsSubnetConfig `json:"awsSubnetConfig,omitempty"`
	}{
		AwsExternalSubnetConfigProxy: (*AwsExternalSubnetConfigProxy)(p),
		AwsSubnetConfig:              p.AwsSubnetConfig,
	})
}

func NewAwsExternalSubnetConfig() *AwsExternalSubnetConfig {
	p := new(AwsExternalSubnetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AwsExternalSubnetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The AWS subnet configuration.
*/
type AwsSubnetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  AWS subnet CIDR.
	*/
	Cidr *string `json:"cidr"`

	GatewayIpAddress *import1.IPAddress `json:"gatewayIpAddress"`
	/*
	  The security group ID list.
	*/
	SecurityGroupReferenceList []string `json:"securityGroupReferenceList"`
	/*
	  AWS subnet Id.
	*/
	SubnetReference *string `json:"subnetReference"`
	/*
	  VPC
	*/
	VpcReference *string `json:"vpcReference"`
}

func (p *AwsSubnetConfig) MarshalJSON() ([]byte, error) {
	type AwsSubnetConfigProxy AwsSubnetConfig
	return json.Marshal(struct {
		*AwsSubnetConfigProxy
		Cidr                       *string            `json:"cidr,omitempty"`
		GatewayIpAddress           *import1.IPAddress `json:"gatewayIpAddress,omitempty"`
		SecurityGroupReferenceList []string           `json:"securityGroupReferenceList,omitempty"`
		SubnetReference            *string            `json:"subnetReference,omitempty"`
		VpcReference               *string            `json:"vpcReference,omitempty"`
	}{
		AwsSubnetConfigProxy:       (*AwsSubnetConfigProxy)(p),
		Cidr:                       p.Cidr,
		GatewayIpAddress:           p.GatewayIpAddress,
		SecurityGroupReferenceList: p.SecurityGroupReferenceList,
		SubnetReference:            p.SubnetReference,
		VpcReference:               p.VpcReference,
	})
}

func NewAwsSubnetConfig() *AwsSubnetConfig {
	p := new(AwsSubnetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AwsSubnetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Azure configuration.
*/
type AzureConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuthorizationData *AuthorizationData `json:"authorizationData,omitempty"`

	BgpInfo *BgpInfo `json:"bgpInfo,omitempty"`
	/*
	  External routing domain UUID list.
	*/
	ExternalRoutingDomainUuidList []string `json:"externalRoutingDomainUuidList,omitempty"`
	/*
	  The external subnet configuration list for the Azure cloud.
	*/
	ExternalSubnetConfigList []AzureExternalSubnetConfig `json:"externalSubnetConfigList"`
	/*
	  True if flow gateway scale out is supported.
	*/
	IsScaleOut *bool `json:"isScaleOut,omitempty"`
}

func (p *AzureConfig) MarshalJSON() ([]byte, error) {
	type AzureConfigProxy AzureConfig
	return json.Marshal(struct {
		*AzureConfigProxy
		ExternalSubnetConfigList []AzureExternalSubnetConfig `json:"externalSubnetConfigList,omitempty"`
	}{
		AzureConfigProxy:         (*AzureConfigProxy)(p),
		ExternalSubnetConfigList: p.ExternalSubnetConfigList,
	})
}

func NewAzureConfig() *AzureConfig {
	p := new(AzureConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AzureConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsScaleOut = new(bool)
	*p.IsScaleOut = false

	return p
}

/*
Binding of Atlas Flow Gateway external subnet with Azure external subnet.
*/
type AzureExternalSubnetBinding struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AzureExternalNetworkPrefix *IPSubnet `json:"azureExternalNetworkPrefix"`
	/*
	  External network interface Id.
	*/
	EniId *string `json:"eniId,omitempty"`

	EniPrimaryIp *import1.IPAddress `json:"eniPrimaryIp,omitempty"`

	PeerToPeerNetworkPrefix *IPSubnet `json:"peerToPeerNetworkPrefix,omitempty"`
	/*
	  Resource group Id.
	*/
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`
	/*
	  MAC address of the router port of Atlas Flow Gateway.
	*/
	RouterPortMac *string `json:"routerPortMac"`
}

func (p *AzureExternalSubnetBinding) MarshalJSON() ([]byte, error) {
	type AzureExternalSubnetBindingProxy AzureExternalSubnetBinding
	return json.Marshal(struct {
		*AzureExternalSubnetBindingProxy
		AzureExternalNetworkPrefix *IPSubnet `json:"azureExternalNetworkPrefix,omitempty"`
		RouterPortMac              *string   `json:"routerPortMac,omitempty"`
	}{
		AzureExternalSubnetBindingProxy: (*AzureExternalSubnetBindingProxy)(p),
		AzureExternalNetworkPrefix:      p.AzureExternalNetworkPrefix,
		RouterPortMac:                   p.RouterPortMac,
	})
}

func NewAzureExternalSubnetBinding() *AzureExternalSubnetBinding {
	p := new(AzureExternalSubnetBinding)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AzureExternalSubnetBinding"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The external subnet configuration for the Azure cloud.
*/
type AzureExternalSubnetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Gateway MAC address of the external subnet.
	*/
	GatewayMacAddress *string `json:"gatewayMacAddress,omitempty"`

	IpConfig *IPConfig `json:"ipConfig"`
	/*
	  List of mappings of private IP (SNAT or floating IP) to public IP.
	*/
	PublicIpMappingList []PublicIpMapping `json:"publicIpMappingList,omitempty"`
}

func (p *AzureExternalSubnetConfig) MarshalJSON() ([]byte, error) {
	type AzureExternalSubnetConfigProxy AzureExternalSubnetConfig
	return json.Marshal(struct {
		*AzureExternalSubnetConfigProxy
		IpConfig *IPConfig `json:"ipConfig,omitempty"`
	}{
		AzureExternalSubnetConfigProxy: (*AzureExternalSubnetConfigProxy)(p),
		IpConfig:                       p.IpConfig,
	})
}

func NewAzureExternalSubnetConfig() *AzureExternalSubnetConfig {
	p := new(AzureExternalSubnetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.AzureExternalSubnetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Base model of route.
*/
type BaseRoute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  BGP session description.
	*/
	Description *string `json:"description,omitempty"`

	Destination *IPSubnet `json:"destination,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Route name.
	*/
	Name *string `json:"name,omitempty"`

	Nexthop *Nexthop `json:"nexthop,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewBaseRoute() *BaseRoute {
	p := new(BaseRoute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BaseRoute"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
BGP configuration.
*/
type BgpConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Autonomous system number. 0 and 4294967295 are reserved.
	*/
	Asn *int64 `json:"asn,omitempty"`
	/*
	  BPG password.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Redistribute routes over eBGP. Applicable only to network gateways deployed on VLAN subnets with eBGP over VPN.
	*/
	ShouldRedistributeRoutes *bool `json:"shouldRedistributeRoutes,omitempty"`
}

func NewBgpConfig() *BgpConfig {
	p := new(BgpConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BgpConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldRedistributeRoutes = new(bool)
	*p.ShouldRedistributeRoutes = false

	return p
}

/*
BGP info needed for flow gateway scale out model.
*/
type BgpInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LocalBgpGatewayList []LocalBgpGateway `json:"localBgpGatewayList"`

	RemoteBgpGatewayList []RemoteBgpGateway `json:"remoteBgpGatewayList"`
}

func (p *BgpInfo) MarshalJSON() ([]byte, error) {
	type BgpInfoProxy BgpInfo
	return json.Marshal(struct {
		*BgpInfoProxy
		LocalBgpGatewayList  []LocalBgpGateway  `json:"localBgpGatewayList,omitempty"`
		RemoteBgpGatewayList []RemoteBgpGateway `json:"remoteBgpGatewayList,omitempty"`
	}{
		BgpInfoProxy:         (*BgpInfoProxy)(p),
		LocalBgpGatewayList:  p.LocalBgpGatewayList,
		RemoteBgpGatewayList: p.RemoteBgpGatewayList,
	})
}

func NewBgpInfo() *BgpInfo {
	p := new(BgpInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BgpInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Route advertised or received over a BGP session.
*/
type BgpRoute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BgpRouteType *BgpRouteType `json:"bgpRouteType"`
	/*
	  BGP session over which the route is advertised or received.
	*/
	BgpSessionReference *string `json:"bgpSessionReference"`
	/*
	  BGP session description.
	*/
	Description *string `json:"description,omitempty"`

	Destination *IPSubnet `json:"destination,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Route name.
	*/
	Name *string `json:"name,omitempty"`

	Nexthop *Nexthop `json:"nexthop,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BgpRoute) MarshalJSON() ([]byte, error) {
	type BgpRouteProxy BgpRoute
	return json.Marshal(struct {
		*BgpRouteProxy
		BgpRouteType        *BgpRouteType `json:"bgpRouteType,omitempty"`
		BgpSessionReference *string       `json:"bgpSessionReference,omitempty"`
	}{
		BgpRouteProxy:       (*BgpRouteProxy)(p),
		BgpRouteType:        p.BgpRouteType,
		BgpSessionReference: p.BgpSessionReference,
	})
}

func NewBgpRoute() *BgpRoute {
	p := new(BgpRoute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BgpRoute"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the BGP route.
*/
type BgpRouteType int

const (
	BGPROUTETYPE_UNKNOWN              BgpRouteType = 0
	BGPROUTETYPE_REDACTED             BgpRouteType = 1
	BGPROUTETYPE_ADVERTISED           BgpRouteType = 2
	BGPROUTETYPE_RECEIVED             BgpRouteType = 3
	BGPROUTETYPE_RECEIVED_AND_IGNORED BgpRouteType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BgpRouteType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVERTISED",
		"RECEIVED",
		"RECEIVED_AND_IGNORED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BgpRouteType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVERTISED",
		"RECEIVED",
		"RECEIVED_AND_IGNORED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BgpRouteType) index(name string) BgpRouteType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVERTISED",
		"RECEIVED",
		"RECEIVED_AND_IGNORED",
	}
	for idx := range names {
		if names[idx] == name {
			return BgpRouteType(idx)
		}
	}
	return BGPROUTETYPE_UNKNOWN
}

func (e *BgpRouteType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BgpRouteType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BgpRouteType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BgpRouteType) Ref() *BgpRouteType {
	return &e
}

/*
BGP session.
*/
type BgpSession struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  BGP session description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Priority assigned to routes received over this BGP session.
	*/
	DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocalGateway *Gateway `json:"localGateway,omitempty"`

	LocalGatewayInterfaceIpAddress *import1.IPAddress `json:"localGatewayInterfaceIpAddress,omitempty"`
	/*
	  Local BGP gateway reference.
	*/
	LocalGatewayReference *string `json:"localGatewayReference"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  BGP session name.
	*/
	Name *string `json:"name"`
	/*
	  BGP password
	*/
	Password *string `json:"password,omitempty"`

	RemoteGateway *Gateway `json:"remoteGateway,omitempty"`
	/*
	  Remote BGP gateway reference.
	*/
	RemoteGatewayReference *string `json:"remoteGatewayReference"`

	Status *Status `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BgpSession) MarshalJSON() ([]byte, error) {
	type BgpSessionProxy BgpSession
	return json.Marshal(struct {
		*BgpSessionProxy
		LocalGatewayReference  *string `json:"localGatewayReference,omitempty"`
		Name                   *string `json:"name,omitempty"`
		RemoteGatewayReference *string `json:"remoteGatewayReference,omitempty"`
	}{
		BgpSessionProxy:        (*BgpSessionProxy)(p),
		LocalGatewayReference:  p.LocalGatewayReference,
		Name:                   p.Name,
		RemoteGatewayReference: p.RemoteGatewayReference,
	})
}

func NewBgpSession() *BgpSession {
	p := new(BgpSession)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BgpSession"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BgpSessionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  BGP session description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Priority assigned to routes received over this BGP session.
	*/
	DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GatewayProjection *GatewayProjection `json:"gatewayProjection,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocalGateway *Gateway `json:"localGateway,omitempty"`

	LocalGatewayInterfaceIpAddress *import1.IPAddress `json:"localGatewayInterfaceIpAddress,omitempty"`
	/*
	  Local BGP gateway reference.
	*/
	LocalGatewayReference *string `json:"localGatewayReference"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  BGP session name.
	*/
	Name *string `json:"name"`
	/*
	  BGP password
	*/
	Password *string `json:"password,omitempty"`

	RemoteGateway *Gateway `json:"remoteGateway,omitempty"`
	/*
	  Remote BGP gateway reference.
	*/
	RemoteGatewayReference *string `json:"remoteGatewayReference"`

	Status *Status `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BgpSessionProjection) MarshalJSON() ([]byte, error) {
	type BgpSessionProjectionProxy BgpSessionProjection
	return json.Marshal(struct {
		*BgpSessionProjectionProxy
		LocalGatewayReference  *string `json:"localGatewayReference,omitempty"`
		Name                   *string `json:"name,omitempty"`
		RemoteGatewayReference *string `json:"remoteGatewayReference,omitempty"`
	}{
		BgpSessionProjectionProxy: (*BgpSessionProjectionProxy)(p),
		LocalGatewayReference:     p.LocalGatewayReference,
		Name:                      p.Name,
		RemoteGatewayReference:    p.RemoteGatewayReference,
	})
}

func NewBgpSessionProjection() *BgpSessionProjection {
	p := new(BgpSessionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BgpSessionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The types of bond modes
*/
type BondModeType int

const (
	BONDMODETYPE_UNKNOWN       BondModeType = 0
	BONDMODETYPE_REDACTED      BondModeType = 1
	BONDMODETYPE_ACTIVE_BACKUP BondModeType = 2
	BONDMODETYPE_BALANCE_SLB   BondModeType = 3
	BONDMODETYPE_BALANCE_TCP   BondModeType = 4
	BONDMODETYPE_NONE          BondModeType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BondModeType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BondModeType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BondModeType) index(name string) BondModeType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return BondModeType(idx)
		}
	}
	return BONDMODETYPE_UNKNOWN
}

func (e *BondModeType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BondModeType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BondModeType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BondModeType) Ref() *BondModeType {
	return &e
}

/*
Schema of bridge to migrate to a Virtual Switch
*/
type Bridge struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Prism Element cluster reference. This header can be optionally supplied for Virtual Switch list requests, but is deprecated for all other Virtual Switch fetch/create/update/delete requests.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  Input body to migrate to a Virtual Switch
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of bridge to convert from
	*/
	ExistingBridgeName *string `json:"existingBridgeName"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Virtual Switch name to migrate to
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Bridge) MarshalJSON() ([]byte, error) {
	type BridgeProxy Bridge
	return json.Marshal(struct {
		*BridgeProxy
		ExistingBridgeName *string `json:"existingBridgeName,omitempty"`
		Name               *string `json:"name,omitempty"`
	}{
		BridgeProxy:        (*BridgeProxy)(p),
		ExistingBridgeName: p.ExistingBridgeName,
		Name:               p.Name,
	})
}

func NewBridge() *Bridge {
	p := new(Bridge)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Bridge"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BridgeProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Prism Element cluster reference. This header can be optionally supplied for Virtual Switch list requests, but is deprecated for all other Virtual Switch fetch/create/update/delete requests.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  Input body to migrate to a Virtual Switch
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of bridge to convert from
	*/
	ExistingBridgeName *string `json:"existingBridgeName"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Virtual Switch name to migrate to
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BridgeProjection) MarshalJSON() ([]byte, error) {
	type BridgeProjectionProxy BridgeProjection
	return json.Marshal(struct {
		*BridgeProjectionProxy
		ExistingBridgeName *string `json:"existingBridgeName,omitempty"`
		Name               *string `json:"name,omitempty"`
	}{
		BridgeProjectionProxy: (*BridgeProjectionProxy)(p),
		ExistingBridgeName:    p.ExistingBridgeName,
		Name:                  p.Name,
	})
}

func NewBridgeProjection() *BridgeProjection {
	p := new(BridgeProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.BridgeProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Capability dictionary entry with capability name and boolean value indicating support
*/
type Capability struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the capability e.g. "SUPPORTS_PC_SPAN"
	*/
	CapabilityName *string `json:"capabilityName"`
	/*
	  Boolean - true if the capability is supported, false otherwise
	*/
	IsSupported *bool `json:"isSupported"`
}

func (p *Capability) MarshalJSON() ([]byte, error) {
	type CapabilityProxy Capability
	return json.Marshal(struct {
		*CapabilityProxy
		CapabilityName *string `json:"capabilityName,omitempty"`
		IsSupported    *bool   `json:"isSupported,omitempty"`
	}{
		CapabilityProxy: (*CapabilityProxy)(p),
		CapabilityName:  p.CapabilityName,
		IsSupported:     p.IsSupported,
	})
}

func NewCapability() *Capability {
	p := new(Capability)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Capability"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cloud substrate of the network controller, for e.g. Azure.
*/
type CloudSubstrate int

const (
	CLOUDSUBSTRATE_UNKNOWN  CloudSubstrate = 0
	CLOUDSUBSTRATE_REDACTED CloudSubstrate = 1
	CLOUDSUBSTRATE_AZURE    CloudSubstrate = 2
	CLOUDSUBSTRATE_AWS      CloudSubstrate = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CloudSubstrate) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CloudSubstrate) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CloudSubstrate) index(name string) CloudSubstrate {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudSubstrate(idx)
		}
	}
	return CLOUDSUBSTRATE_UNKNOWN
}

func (e *CloudSubstrate) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudSubstrate:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudSubstrate) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudSubstrate) Ref() *CloudSubstrate {
	return &e
}

/*
Input body to configure cluster
*/
type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference ExtId for the cluster. This is a required parameter on Prism Element ; and is optional on Prism Central
	*/
	ExtId *string `json:"extId"`

	GatewayIpAddress *import1.IPv4Address `json:"gatewayIpAddress,omitempty"`
	/*
	  Host configuration array
	*/
	Hosts []Host `json:"hosts"`
}

func (p *Cluster) MarshalJSON() ([]byte, error) {
	type ClusterProxy Cluster
	return json.Marshal(struct {
		*ClusterProxy
		ExtId *string `json:"extId,omitempty"`
		Hosts []Host  `json:"hosts,omitempty"`
	}{
		ClusterProxy: (*ClusterProxy)(p),
		ExtId:        p.ExtId,
		Hosts:        p.Hosts,
	})
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterCapability struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A dictionary with a key as a capability name and a value as a boolean, true if supported by the cluster, false if not.  For example, dictionary will look like { "SUPPORTS_PC_SPAN": true, "SUPPORTS_VGPU_CONSOLE": true }
	*/
	Capabilities []Capability `json:"capabilities"`
	/*
	  Cluster UUID whose capabilities are retrieved.
	*/
	ClusterId *string `json:"clusterId"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterCapability) MarshalJSON() ([]byte, error) {
	type ClusterCapabilityProxy ClusterCapability
	return json.Marshal(struct {
		*ClusterCapabilityProxy
		Capabilities []Capability `json:"capabilities,omitempty"`
		ClusterId    *string      `json:"clusterId,omitempty"`
	}{
		ClusterCapabilityProxy: (*ClusterCapabilityProxy)(p),
		Capabilities:           p.Capabilities,
		ClusterId:              p.ClusterId,
	})
}

func NewClusterCapability() *ClusterCapability {
	p := new(ClusterCapability)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ClusterCapability"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Get the Flow Networking usage of each registered Prism Element cluster.
*/
type ClusterFlowStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flow Networking usage status for every cluster.
	*/
	ClusterStatusList []ClusterStatus `json:"clusterStatusList,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewClusterFlowStatus() *ClusterFlowStatus {
	p := new(ClusterFlowStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ClusterFlowStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Flow Networking usage status for a Prism Element cluster.
*/
type ClusterStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of respective cluster.
	*/
	ClusterReference *string `json:"clusterReference"`
	/*
	  Indicates the flow status on the cluster. It is set to True if the cluster has at least one vNIC that is part of an Atlas subnet
	*/
	HasFlowStatus *bool `json:"hasFlowStatus"`
}

func (p *ClusterStatus) MarshalJSON() ([]byte, error) {
	type ClusterStatusProxy ClusterStatus
	return json.Marshal(struct {
		*ClusterStatusProxy
		ClusterReference *string `json:"clusterReference,omitempty"`
		HasFlowStatus    *bool   `json:"hasFlowStatus,omitempty"`
	}{
		ClusterStatusProxy: (*ClusterStatusProxy)(p),
		ClusterReference:   p.ClusterReference,
		HasFlowStatus:      p.HasFlowStatus,
	})
}

func NewClusterStatus() *ClusterStatus {
	p := new(ClusterStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ClusterStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Current status of the network controller.
*/
type ControllerStatus int

const (
	CONTROLLERSTATUS_UNKNOWN  ControllerStatus = 0
	CONTROLLERSTATUS_REDACTED ControllerStatus = 1
	CONTROLLERSTATUS_UP       ControllerStatus = 2
	CONTROLLERSTATUS_DEGRADED ControllerStatus = 3
	CONTROLLERSTATUS_DOWN     ControllerStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ControllerStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DEGRADED",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ControllerStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DEGRADED",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ControllerStatus) index(name string) ControllerStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DEGRADED",
		"DOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return ControllerStatus(idx)
		}
	}
	return CONTROLLERSTATUS_UNKNOWN
}

func (e *ControllerStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ControllerStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ControllerStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ControllerStatus) Ref() *ControllerStatus {
	return &e
}

/*
Default VLAN stack(Legacy or Advanced) to instatiate VLAN-backed subnets on if advanced networking is enabled
*/
type DefaultVlanStack int

const (
	DEFAULTVLANSTACK_UNKNOWN  DefaultVlanStack = 0
	DEFAULTVLANSTACK_REDACTED DefaultVlanStack = 1
	DEFAULTVLANSTACK_ADVANCED DefaultVlanStack = 2
	DEFAULTVLANSTACK_LEGACY   DefaultVlanStack = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DefaultVlanStack) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVANCED",
		"LEGACY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DefaultVlanStack) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVANCED",
		"LEGACY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DefaultVlanStack) index(name string) DefaultVlanStack {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADVANCED",
		"LEGACY",
	}
	for idx := range names {
		if names[idx] == name {
			return DefaultVlanStack(idx)
		}
	}
	return DEFAULTVLANSTACK_UNKNOWN
}

func (e *DefaultVlanStack) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DefaultVlanStack:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DefaultVlanStack) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DefaultVlanStack) Ref() *DefaultVlanStack {
	return &e
}

/*
List of DHCP options to be configured.
*/
type DhcpOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Boot file name (option 67).
	*/
	BootFileName *string `json:"bootFileName,omitempty"`
	/*
	  The DNS domain name of the client (option 15).
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  List of Domain Name Server addresses (option 6).
	*/
	DomainNameServers []import1.IPAddress `json:"domainNameServers,omitempty"`
	/*
	  List of NTP server addresses (option 42).
	*/
	NtpServers []import1.IPAddress `json:"ntpServers,omitempty"`
	/*
	  The DNS domain search list (option 119).
	*/
	SearchDomains []string `json:"searchDomains,omitempty"`
	/*
	  TFTP server name (option 66).
	*/
	TftpServerName *string `json:"tftpServerName,omitempty"`
}

func NewDhcpOptions() *DhcpOptions {
	p := new(DhcpOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.DhcpOptions"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Dead peer detection configuration for the VPN connection
*/
type DpdConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The amount of time the peer waits for traffic before sending a DPD request
	*/
	IntervalSecs *int64 `json:"intervalSecs,omitempty"`

	Operation *DpdOperation `json:"operation,omitempty"`
	/*
	  The maximum amount of time to wait for a DPD response before marking the peer as dead
	*/
	TimeoutSecs *int64 `json:"timeoutSecs,omitempty"`
}

func NewDpdConfig() *DpdConfig {
	p := new(DpdConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.DpdConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation to be performed on detecting a dead peer. The default is HOLD.
*/
type DpdOperation int

const (
	DPDOPERATION_UNKNOWN  DpdOperation = 0
	DPDOPERATION_REDACTED DpdOperation = 1
	DPDOPERATION_RESTART  DpdOperation = 2
	DPDOPERATION_CLEAR    DpdOperation = 3
	DPDOPERATION_HOLD     DpdOperation = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DpdOperation) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESTART",
		"CLEAR",
		"HOLD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DpdOperation) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESTART",
		"CLEAR",
		"HOLD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DpdOperation) index(name string) DpdOperation {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RESTART",
		"CLEAR",
		"HOLD",
	}
	for idx := range names {
		if names[idx] == name {
			return DpdOperation(idx)
		}
	}
	return DPDOPERATION_UNKNOWN
}

func (e *DpdOperation) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DpdOperation:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DpdOperation) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DpdOperation) Ref() *DpdOperation {
	return &e
}

/*
Encryption algorithm.
*/
type EncryptionAlgorithm int

const (
	ENCRYPTIONALGORITHM_UNKNOWN      EncryptionAlgorithm = 0
	ENCRYPTIONALGORITHM_REDACTED     EncryptionAlgorithm = 1
	ENCRYPTIONALGORITHM_AES128       EncryptionAlgorithm = 2
	ENCRYPTIONALGORITHM_AES256       EncryptionAlgorithm = 3
	ENCRYPTIONALGORITHM_TRIPLE_DES   EncryptionAlgorithm = 4
	ENCRYPTIONALGORITHM_AES256GCM128 EncryptionAlgorithm = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EncryptionAlgorithm) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AES128",
		"AES256",
		"TRIPLE_DES",
		"AES256GCM128",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EncryptionAlgorithm) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AES128",
		"AES256",
		"TRIPLE_DES",
		"AES256GCM128",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EncryptionAlgorithm) index(name string) EncryptionAlgorithm {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AES128",
		"AES256",
		"TRIPLE_DES",
		"AES256GCM128",
	}
	for idx := range names {
		if names[idx] == name {
			return EncryptionAlgorithm(idx)
		}
	}
	return ENCRYPTIONALGORITHM_UNKNOWN
}

func (e *EncryptionAlgorithm) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionAlgorithm:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EncryptionAlgorithm) Ref() *EncryptionAlgorithm {
	return &e
}

/*
Contains the UUID and scope type information for a particular export scope.
*/
type ExportScope struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ScopeType *ScopeType `json:"scopeType,omitempty"`
	/*
	  UUID of the PE or PC cluster.
	*/
	Uuid *string `json:"uuid,omitempty"`
}

func NewExportScope() *ExportScope {
	p := new(ExportScope)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ExportScope"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IPFIX exporter protocol:the permissible values are TCP or UDP.
*/
type ExporterProtocol int

const (
	EXPORTERPROTOCOL_UNKNOWN  ExporterProtocol = 0
	EXPORTERPROTOCOL_REDACTED ExporterProtocol = 1
	EXPORTERPROTOCOL_TCP      ExporterProtocol = 2
	EXPORTERPROTOCOL_UDP      ExporterProtocol = 3
	EXPORTERPROTOCOL_TLS_TCP  ExporterProtocol = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ExporterProtocol) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
		"TLS_TCP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ExporterProtocol) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
		"TLS_TCP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ExporterProtocol) index(name string) ExporterProtocol {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
		"TLS_TCP",
	}
	for idx := range names {
		if names[idx] == name {
			return ExporterProtocol(idx)
		}
	}
	return EXPORTERPROTOCOL_UNKNOWN
}

func (e *ExporterProtocol) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ExporterProtocol:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ExporterProtocol) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ExporterProtocol) Ref() *ExporterProtocol {
	return &e
}

/*
Information about the external subnet, SNAT IPs and the gateway nodes.
*/
type ExternalSubnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Maximum number of active gateway nodes for the VPC external subnet association.
	*/
	ActiveGatewayCount *int `json:"activeGatewayCount,omitempty"`
	/*
	  Currently active gateway nodes that are used for external connectivity.
	*/
	ActiveGatewayNodes []GatewayNodeReference `json:"activeGatewayNodes,omitempty"`
	/*
	  List of IP Addresses used for SNAT, if NAT is enabled on the external subnet. If NAT is not enabled, this specifies the IP address of the VPC port connected to the external gateway.
	*/
	ExternalIps []import1.IPAddress `json:"externalIps,omitempty"`
	/*
	  List of gateway nodes that can be used for external connectivity.
	*/
	GatewayNodes []string `json:"gatewayNodes,omitempty"`
	/*
	  External subnet reference.
	*/
	SubnetReference *string `json:"subnetReference"`
}

func (p *ExternalSubnet) MarshalJSON() ([]byte, error) {
	type ExternalSubnetProxy ExternalSubnet
	return json.Marshal(struct {
		*ExternalSubnetProxy
		SubnetReference *string `json:"subnetReference,omitempty"`
	}{
		ExternalSubnetProxy: (*ExternalSubnetProxy)(p),
		SubnetReference:     p.SubnetReference,
	})
}

func NewExternalSubnet() *ExternalSubnet {
	p := new(ExternalSubnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ExternalSubnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Floating IP address.
*/
type FloatingIPAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4 *FloatingIPv4Address `json:"ipv4,omitempty"`

	Ipv6 *FloatingIPv6Address `json:"ipv6,omitempty"`
}

func NewFloatingIPAddress() *FloatingIPAddress {
	p := new(FloatingIPAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FloatingIPAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (i *FloatingIPAddress) HasIpv4() bool {
	return i.Ipv4 != nil
}
func (i *FloatingIPAddress) HasIpv6() bool {
	return i.Ipv6 != nil
}

func (i *FloatingIPAddress) IsValid() bool {
	return i.HasIpv4() || i.HasIpv6()
}

/*
IPv4 address.
*/
type FloatingIPv4Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Prefix length of the network to which this host IPv4 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  Prefix of the network to which this host IPv4 address belongs.
	*/
	Value *string `json:"value,omitempty"`
}

func NewFloatingIPv4Address() *FloatingIPv4Address {
	p := new(FloatingIPv4Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FloatingIPv4Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 32

	return p
}

/*
IPv6 address.
*/
type FloatingIPv6Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Prefix length of the network to which this host IPv6 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  Prefix of the network to which this host IPv6 address belongs.
	*/
	Value *string `json:"value,omitempty"`
}

func NewFloatingIPv6Address() *FloatingIPv6Address {
	p := new(FloatingIPv6Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FloatingIPv6Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 128

	return p
}

/*
Configure a floating IP.
*/
type FloatingIp struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AssociationItemDiscriminator_ *string `json:"$associationItemDiscriminator,omitempty"`
	/*
	  Association of the Floating IP with either NIC or Private IP
	*/
	Association *OneOfFloatingIpAssociation `json:"association,omitempty"`
	/*
	  Association status of floating IP.
	*/
	AssociationStatus *string `json:"associationStatus,omitempty"`
	/*
	  Description for the Floating IP.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalSubnet *Subnet `json:"externalSubnet,omitempty"`
	/*
	  External subnet reference for the Floating IP to be allocated in on-prem only.
	*/
	ExternalSubnetReference *string `json:"externalSubnetReference,omitempty"`

	FloatingIp *FloatingIPAddress `json:"floatingIp,omitempty"`
	/*
	  Floating IP value in string
	*/
	FloatingIpValue *string `json:"floatingIpValue,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Load Balancer Session reference UUID.
	*/
	LoadBalancerSessionReference *string `json:"loadBalancerSessionReference,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the floating IP.
	*/
	Name *string `json:"name"`
	/*
	  Private IP value in string
	*/
	PrivateIp *string `json:"privateIp,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VmNic *VmNic `json:"vmNic,omitempty"`
	/*
	  VM NIC reference.
	*/
	VmNicReference *string `json:"vmNicReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`
	/*
	  VPC reference UUID
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *FloatingIp) MarshalJSON() ([]byte, error) {
	type FloatingIpProxy FloatingIp
	return json.Marshal(struct {
		*FloatingIpProxy
		Name *string `json:"name,omitempty"`
	}{
		FloatingIpProxy: (*FloatingIpProxy)(p),
		Name:            p.Name,
	})
}

func NewFloatingIp() *FloatingIp {
	p := new(FloatingIp)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FloatingIp"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *FloatingIp) GetAssociation() interface{} {
	if nil == p.Association {
		return nil
	}
	return p.Association.GetValue()
}

func (p *FloatingIp) SetAssociation(v interface{}) error {
	if nil == p.Association {
		p.Association = NewOneOfFloatingIpAssociation()
	}
	e := p.Association.SetValue(v)
	if nil == e {
		if nil == p.AssociationItemDiscriminator_ {
			p.AssociationItemDiscriminator_ = new(string)
		}
		*p.AssociationItemDiscriminator_ = *p.Association.Discriminator
	}
	return e
}

type FloatingIpProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AssociationItemDiscriminator_ *string `json:"$associationItemDiscriminator,omitempty"`
	/*
	  Association of the Floating IP with either NIC or Private IP
	*/
	Association *OneOfFloatingIpProjectionAssociation `json:"association,omitempty"`
	/*
	  Association status of floating IP.
	*/
	AssociationStatus *string `json:"associationStatus,omitempty"`
	/*
	  Description for the Floating IP.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalSubnet *Subnet `json:"externalSubnet,omitempty"`
	/*
	  External subnet reference for the Floating IP to be allocated in on-prem only.
	*/
	ExternalSubnetReference *string `json:"externalSubnetReference,omitempty"`

	FloatingIp *FloatingIPAddress `json:"floatingIp,omitempty"`
	/*
	  Floating IP value in string
	*/
	FloatingIpValue *string `json:"floatingIpValue,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Load Balancer Session reference UUID.
	*/
	LoadBalancerSessionReference *string `json:"loadBalancerSessionReference,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the floating IP.
	*/
	Name *string `json:"name"`
	/*
	  Private IP value in string
	*/
	PrivateIp *string `json:"privateIp,omitempty"`

	SubnetProjection *SubnetProjection `json:"subnetProjection,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VmNic *VmNic `json:"vmNic,omitempty"`

	VmNicProjection *VmNicProjection `json:"vmNicProjection,omitempty"`
	/*
	  VM NIC reference.
	*/
	VmNicReference *string `json:"vmNicReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`

	VpcProjection *VpcProjection `json:"vpcProjection,omitempty"`
	/*
	  VPC reference UUID
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *FloatingIpProjection) MarshalJSON() ([]byte, error) {
	type FloatingIpProjectionProxy FloatingIpProjection
	return json.Marshal(struct {
		*FloatingIpProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		FloatingIpProjectionProxy: (*FloatingIpProjectionProxy)(p),
		Name:                      p.Name,
	})
}

func NewFloatingIpProjection() *FloatingIpProjection {
	p := new(FloatingIpProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FloatingIpProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type FlowGateway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of bindings of the Atlas Flow Gateway external subnet with the Azure external subnet.
	*/
	AzureExternalSubnetBindingList []AzureExternalSubnetBinding `json:"azureExternalSubnetBindingList"`
	/*
	  Chassis UUID of the Atlas Flow Gateway.
	*/
	ChassisUuid *string `json:"chassisUuid"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GatewayStatus *FlowGatewayStatus `json:"gatewayStatus,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Version of the OVN controller
	*/
	OvnControllerVersion *string `json:"ovnControllerVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *FlowGateway) MarshalJSON() ([]byte, error) {
	type FlowGatewayProxy FlowGateway
	return json.Marshal(struct {
		*FlowGatewayProxy
		AzureExternalSubnetBindingList []AzureExternalSubnetBinding `json:"azureExternalSubnetBindingList,omitempty"`
		ChassisUuid                    *string                      `json:"chassisUuid,omitempty"`
	}{
		FlowGatewayProxy:               (*FlowGatewayProxy)(p),
		AzureExternalSubnetBindingList: p.AzureExternalSubnetBindingList,
		ChassisUuid:                    p.ChassisUuid,
	})
}

func NewFlowGateway() *FlowGateway {
	p := new(FlowGateway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FlowGateway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response model for Flow Gateway Keepalive.
*/
type FlowGatewayKeepAlive struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the keep alive request was successful or not.
	*/
	KeepAliveResponse *string `json:"keepAliveResponse,omitempty"`
	/*
	  Version of the network controller on the Prism Central.
	*/
	NetworkControllerVersion *string `json:"networkControllerVersion,omitempty"`
}

func NewFlowGatewayKeepAlive() *FlowGatewayKeepAlive {
	p := new(FlowGatewayKeepAlive)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FlowGatewayKeepAlive"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Request model for Flow Gateway Keepalive.
*/
type FlowGatewayKeepAliveRequestSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Chassis UUID of the Atlas Flow Gateway.
	*/
	ChassisUuid *string `json:"chassisUuid"`
	/*
	  List of external network interface's primary IP addresses corresponding to the external subnet bindings in the flow gateway.
	*/
	EniPrimaryIpList []import1.IPAddress `json:"eniPrimaryIpList,omitempty"`

	GatewayStatus *FlowGatewayStatus `json:"gatewayStatus,omitempty"`
	/*
	  Version of the OVN controller
	*/
	OvnControllerVersion *string `json:"ovnControllerVersion,omitempty"`
}

func (p *FlowGatewayKeepAliveRequestSpec) MarshalJSON() ([]byte, error) {
	type FlowGatewayKeepAliveRequestSpecProxy FlowGatewayKeepAliveRequestSpec
	return json.Marshal(struct {
		*FlowGatewayKeepAliveRequestSpecProxy
		ChassisUuid *string `json:"chassisUuid,omitempty"`
	}{
		FlowGatewayKeepAliveRequestSpecProxy: (*FlowGatewayKeepAliveRequestSpecProxy)(p),
		ChassisUuid:                          p.ChassisUuid,
	})
}

func NewFlowGatewayKeepAliveRequestSpec() *FlowGatewayKeepAliveRequestSpec {
	p := new(FlowGatewayKeepAliveRequestSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FlowGatewayKeepAliveRequestSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
State of the Atlas Flow Gateway.
*/
type FlowGatewayState int

const (
	FLOWGATEWAYSTATE_UNKNOWN      FlowGatewayState = 0
	FLOWGATEWAYSTATE_REDACTED     FlowGatewayState = 1
	FLOWGATEWAYSTATE_HEALTHY      FlowGatewayState = 2
	FLOWGATEWAYSTATE_DOWN         FlowGatewayState = 3
	FLOWGATEWAYSTATE_PROVISIONING FlowGatewayState = 4
	FLOWGATEWAYSTATE_MAINTENANCE  FlowGatewayState = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FlowGatewayState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"DOWN",
		"PROVISIONING",
		"MAINTENANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FlowGatewayState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"DOWN",
		"PROVISIONING",
		"MAINTENANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FlowGatewayState) index(name string) FlowGatewayState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"DOWN",
		"PROVISIONING",
		"MAINTENANCE",
	}
	for idx := range names {
		if names[idx] == name {
			return FlowGatewayState(idx)
		}
	}
	return FLOWGATEWAYSTATE_UNKNOWN
}

func (e *FlowGatewayState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FlowGatewayState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FlowGatewayState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FlowGatewayState) Ref() *FlowGatewayState {
	return &e
}

/*
Status of the Atlas Flow Gateway.
*/
type FlowGatewayStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Status detail of the Atlas Flow Gateway.
	*/
	Detail *string `json:"detail,omitempty"`

	State *FlowGatewayState `json:"state,omitempty"`
}

func NewFlowGatewayStatus() *FlowGatewayStatus {
	p := new(FlowGatewayStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.FlowGatewayStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network gateway.
*/
type Gateway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cloud network on which network gateway is deployed.
	*/
	CloudNetworkReference *string `json:"cloudNetworkReference,omitempty"`

	Deployment *GatewayDeployment `json:"deployment,omitempty"`
	/*
	  Description of the gateway.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Third-party gateway vendor.
	*/
	GatewayDeviceVendor *string `json:"gatewayDeviceVendor,omitempty"`
	/*
	  Software version installed on the gateway appliance.
	*/
	InstalledSoftwareVersion *string `json:"installedSoftwareVersion,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the gateway.
	*/
	Name *string `json:"name,omitempty"`
	/*

	 */
	ServicesItemDiscriminator_ *string `json:"$servicesItemDiscriminator,omitempty"`
	/*
	  Local or remote gateway service type.
	*/
	Services *OneOfGatewayServices `json:"services,omitempty"`

	Status *Status `json:"status,omitempty"`
	/*
	  Supported gateway appliance version.
	*/
	SupportedSoftwareVersion *string `json:"supportedSoftwareVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vm *Vm `json:"vm,omitempty"`
	/*
	  Reference to a dedicated VM on which a local gateway is deployed.
	*/
	VmReference *string `json:"vmReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`
	/*
	  VPC
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewGateway() *Gateway {
	p := new(Gateway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Gateway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Gateway) GetServices() interface{} {
	if nil == p.Services {
		return nil
	}
	return p.Services.GetValue()
}

func (p *Gateway) SetServices(v interface{}) error {
	if nil == p.Services {
		p.Services = NewOneOfGatewayServices()
	}
	e := p.Services.SetValue(v)
	if nil == e {
		if nil == p.ServicesItemDiscriminator_ {
			p.ServicesItemDiscriminator_ = new(string)
		}
		*p.ServicesItemDiscriminator_ = *p.Services.Discriminator
	}
	return e
}

/*
Network gateway deployment configuration.
*/
type GatewayDeployment struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster reference required to identify which on-prem cluster to deploy the gateway VM on.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  List of network interfaces for this gateway.
	*/
	Interfaces []GatewayInterface `json:"interfaces,omitempty"`

	ManagementInterface *GatewayManagementInterface `json:"managementInterface,omitempty"`
	/*
	  vCenter datastore to which the gateway disks and images will be uploaded during deployment.
	*/
	VcenterDatastoreName *string `json:"vcenterDatastoreName,omitempty"`
}

func NewGatewayDeployment() *GatewayDeployment {
	p := new(GatewayDeployment)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayDeployment"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network interface used to deliver network services. If the client  supplies a VPC reference, but does not specify a VPC subnet, the gateway will be deployed on an auto-configured dedicated subnet  within the VPC. If the client specifies a VPC subnet, the gateway  will be deployed on that subnet. If the gateway is deployed on a VLAN subnet, with or without IPAM, then statically-allocated IP addresses on that subnet must be provided for each interface.
*/
type GatewayInterface struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGatewayAddress *import1.IPAddress `json:"defaultGatewayAddress,omitempty"`

	IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
	/*
	  MAC address of this gateway interface.
	*/
	MacAddress *string `json:"macAddress,omitempty"`
	/*
	  MTU of this gateway interface.
	*/
	Mtu *int `json:"mtu,omitempty"`
	/*
	  The VLAN subnet to deploy this network gateway VM on.
	*/
	SubnetReference *string `json:"subnetReference,omitempty"`
}

func NewGatewayInterface() *GatewayInterface {
	p := new(GatewayInterface)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayInterface"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network interface used to deliver network services and for managing the gateway. If a VPC reference is supplied then the gateway will be deployed on a dedicated subnet within the VPC. If a VPC reference is not supplied, then this interface defines the subnet on which the gateway will be deployed, and the address it will be assigned. When a VPC reference is not present, either a subnet reference or a VLAN id must be supplied, along with the address and default gateway of the subnet. A VLAN network without IPAM may be used.
*/
type GatewayManagementInterface struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddress `json:"address,omitempty"`

	DefaultGateway *import1.IPAddress `json:"defaultGateway,omitempty"`
	/*
	  MTU of management interface.
	*/
	Mtu *int `json:"mtu,omitempty"`
	/*
	  The on-prem vlan subnet to deploy the network gateway VM on.
	*/
	SubnetReference *string `json:"subnetReference,omitempty"`
	/*
	  The on-prem VLAN to deploy the gateway on.
	*/
	VlanId *int `json:"vlanId,omitempty"`
}

func NewGatewayManagementInterface() *GatewayManagementInterface {
	p := new(GatewayManagementInterface)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayManagementInterface"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the VNICs attached to the local BGP gateways.
*/
type GatewayNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Index of the NIC.
	*/
	Index *int `json:"index"`

	IpAddress *import1.IPAddress `json:"ipAddress"`
	/*
	  MAC address of the NIC.
	*/
	MacAddress *string `json:"macAddress"`
}

func (p *GatewayNic) MarshalJSON() ([]byte, error) {
	type GatewayNicProxy GatewayNic
	return json.Marshal(struct {
		*GatewayNicProxy
		Index      *int               `json:"index,omitempty"`
		IpAddress  *import1.IPAddress `json:"ipAddress,omitempty"`
		MacAddress *string            `json:"macAddress,omitempty"`
	}{
		GatewayNicProxy: (*GatewayNicProxy)(p),
		Index:           p.Index,
		IpAddress:       p.IpAddress,
		MacAddress:      p.MacAddress,
	})
}

func NewGatewayNic() *GatewayNic {
	p := new(GatewayNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference of gateway nodes.
*/
type GatewayNodeReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of gateway nodes
	*/
	NodeId *string `json:"nodeId,omitempty"`

	NodeIpAddress *import1.IPAddress `json:"nodeIpAddress,omitempty"`
}

func NewGatewayNodeReference() *GatewayNodeReference {
	p := new(GatewayNodeReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayNodeReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type GatewayProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cloud network on which network gateway is deployed.
	*/
	CloudNetworkReference *string `json:"cloudNetworkReference,omitempty"`

	Deployment *GatewayDeployment `json:"deployment,omitempty"`
	/*
	  Description of the gateway.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Third-party gateway vendor.
	*/
	GatewayDeviceVendor *string `json:"gatewayDeviceVendor,omitempty"`
	/*
	  Software version installed on the gateway appliance.
	*/
	InstalledSoftwareVersion *string `json:"installedSoftwareVersion,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the gateway.
	*/
	Name *string `json:"name,omitempty"`

	ServicesItemDiscriminator_ *string `json:"$servicesItemDiscriminator,omitempty"`
	/*
	  Local or remote gateway service type.
	*/
	Services *OneOfGatewayProjectionServices `json:"services,omitempty"`

	Status *Status `json:"status,omitempty"`
	/*
	  Supported gateway appliance version.
	*/
	SupportedSoftwareVersion *string `json:"supportedSoftwareVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vm *Vm `json:"vm,omitempty"`

	VmProjection *VmProjection `json:"vmProjection,omitempty"`
	/*
	  Reference to a dedicated VM on which a local gateway is deployed.
	*/
	VmReference *string `json:"vmReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`

	VpcProjection *VpcProjection `json:"vpcProjection,omitempty"`
	/*
	  VPC
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewGatewayProjection() *GatewayProjection {
	p := new(GatewayProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GatewayProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Local gateway role (acceptor or initiator) in the connection
*/
type GatewayRole int

const (
	GATEWAYROLE_UNKNOWN   GatewayRole = 0
	GATEWAYROLE_REDACTED  GatewayRole = 1
	GATEWAYROLE_INITIATOR GatewayRole = 2
	GATEWAYROLE_ACCEPTOR  GatewayRole = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GatewayRole) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INITIATOR",
		"ACCEPTOR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GatewayRole) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INITIATOR",
		"ACCEPTOR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GatewayRole) index(name string) GatewayRole {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INITIATOR",
		"ACCEPTOR",
	}
	for idx := range names {
		if names[idx] == name {
			return GatewayRole(idx)
		}
	}
	return GATEWAYROLE_UNKNOWN
}

func (e *GatewayRole) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GatewayRole:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GatewayRole) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GatewayRole) Ref() *GatewayRole {
	return &e
}

/*
REST response for all response codes in API path /networking/v4.0/config/bgp-sessions/{bgpSessionExtId}/bgp-routes/{extId} Get operation
*/
type GetBgpRouteApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetBgpRouteApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetBgpRouteApiResponse() *GetBgpRouteApiResponse {
	p := new(GetBgpRouteApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetBgpRouteApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetBgpRouteApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetBgpRouteApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetBgpRouteApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/bgp-sessions/{extId} Get operation
*/
type GetBgpSessionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetBgpSessionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetBgpSessionApiResponse() *GetBgpSessionApiResponse {
	p := new(GetBgpSessionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetBgpSessionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetBgpSessionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetBgpSessionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetBgpSessionApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/floating-ips/{extId} Get operation
*/
type GetFloatingIpApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetFloatingIpApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetFloatingIpApiResponse() *GetFloatingIpApiResponse {
	p := new(GetFloatingIpApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetFloatingIpApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetFloatingIpApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetFloatingIpApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetFloatingIpApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/gateways/{extId} Get operation
*/
type GetGatewayApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGatewayApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetGatewayApiResponse() *GetGatewayApiResponse {
	p := new(GetGatewayApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetGatewayApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetGatewayApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetGatewayApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetGatewayApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/ipfix-exporters/{extId} Get operation
*/
type GetIPFIXExporterApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetIPFIXExporterApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetIPFIXExporterApiResponse() *GetIPFIXExporterApiResponse {
	p := new(GetIPFIXExporterApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetIPFIXExporterApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetIPFIXExporterApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetIPFIXExporterApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetIPFIXExporterApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/layer2-stretches/{extId} Get operation
*/
type GetLayer2StretchApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLayer2StretchApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLayer2StretchApiResponse() *GetLayer2StretchApiResponse {
	p := new(GetLayer2StretchApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetLayer2StretchApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLayer2StretchApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLayer2StretchApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLayer2StretchApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/layer2-stretches/{layer2StretchExtId}/learned-mac-addresses/{extId} Get operation
*/
type GetLearnedMacAddressApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLearnedMacAddressApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLearnedMacAddressApiResponse() *GetLearnedMacAddressApiResponse {
	p := new(GetLearnedMacAddressApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetLearnedMacAddressApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLearnedMacAddressApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLearnedMacAddressApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLearnedMacAddressApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/load-balancer-sessions/{extId} Get operation
*/
type GetLoadBalancerSessionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLoadBalancerSessionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLoadBalancerSessionApiResponse() *GetLoadBalancerSessionApiResponse {
	p := new(GetLoadBalancerSessionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetLoadBalancerSessionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLoadBalancerSessionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLoadBalancerSessionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLoadBalancerSessionApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/controllers/{extId} Get operation
*/
type GetNetworkControllerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNetworkControllerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNetworkControllerApiResponse() *GetNetworkControllerApiResponse {
	p := new(GetNetworkControllerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetNetworkControllerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNetworkControllerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNetworkControllerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNetworkControllerApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-subnets/{extId} Get operation
*/
type GetRemoteSubnetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRemoteSubnetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRemoteSubnetApiResponse() *GetRemoteSubnetApiResponse {
	p := new(GetRemoteSubnetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetRemoteSubnetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRemoteSubnetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRemoteSubnetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRemoteSubnetApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/route-tables/{routeTableExtId}/routes/{extId} Get operation
*/
type GetRouteApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRouteApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRouteApiResponse() *GetRouteApiResponse {
	p := new(GetRouteApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetRouteApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRouteApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRouteApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRouteApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/route-tables/{extId} Get operation
*/
type GetRouteTableApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRouteTableApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRouteTableApiResponse() *GetRouteTableApiResponse {
	p := new(GetRouteTableApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetRouteTableApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRouteTableApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRouteTableApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRouteTableApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/routing-policies/{extId} Get operation
*/
type GetRoutingPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRoutingPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRoutingPolicyApiResponse() *GetRoutingPolicyApiResponse {
	p := new(GetRoutingPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetRoutingPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRoutingPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRoutingPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRoutingPolicyApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/subnets/{extId} Get operation
*/
type GetSubnetApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSubnetApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSubnetApiResponse() *GetSubnetApiResponse {
	p := new(GetSubnetApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetSubnetApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSubnetApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSubnetApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSubnetApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/traffic-mirrors/{extId} Get operation
*/
type GetTrafficMirrorApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTrafficMirrorApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetTrafficMirrorApiResponse() *GetTrafficMirrorApiResponse {
	p := new(GetTrafficMirrorApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetTrafficMirrorApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTrafficMirrorApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTrafficMirrorApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTrafficMirrorApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/uplink-bonds/{extId} Get operation
*/
type GetUplinkBondApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUplinkBondApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUplinkBondApiResponse() *GetUplinkBondApiResponse {
	p := new(GetUplinkBondApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetUplinkBondApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUplinkBondApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUplinkBondApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUplinkBondApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/virtual-switches/{extId} Get operation
*/
type GetVirtualSwitchApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVirtualSwitchApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVirtualSwitchApiResponse() *GetVirtualSwitchApiResponse {
	p := new(GetVirtualSwitchApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetVirtualSwitchApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVirtualSwitchApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVirtualSwitchApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVirtualSwitchApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpcs/{extId} Get operation
*/
type GetVpcApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVpcApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVpcApiResponse() *GetVpcApiResponse {
	p := new(GetVpcApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetVpcApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVpcApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVpcApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVpcApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpn-connections/{extId} Get operation
*/
type GetVpnConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVpnConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVpnConnectionApiResponse() *GetVpnConnectionApiResponse {
	p := new(GetVpnConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.GetVpnConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVpnConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVpnConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVpnConnectionApiResponseData()
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
Health check configuration for the load balancer session.
*/
type HealthCheck struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The number of failure checks after which the target is considered unhealthy.
	*/
	FailureThreshold *int `json:"failureThreshold,omitempty"`
	/*
	  The interval, in seconds, between health checks.
	*/
	IntervalSecs *int `json:"intervalSecs,omitempty"`
	/*
	  The number of successful checks after which the target is considered healthy.
	*/
	SuccessThreshold *int `json:"successThreshold,omitempty"`
	/*
	  The time, in seconds, after which a health check times out.
	*/
	TimeoutSecs *int `json:"timeoutSecs,omitempty"`
}

func NewHealthCheck() *HealthCheck {
	p := new(HealthCheck)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.HealthCheck"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.FailureThreshold = new(int)
	*p.FailureThreshold = 3
	p.IntervalSecs = new(int)
	*p.IntervalSecs = 5
	p.SuccessThreshold = new(int)
	*p.SuccessThreshold = 3
	p.TimeoutSecs = new(int)
	*p.TimeoutSecs = 2

	return p
}

/*
Input body to configure hosts
*/
type Host struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference to the host
	*/
	ExtId *string `json:"extId"`
	/*
	  Host NIC array
	*/
	HostNics []string `json:"hostNics,omitempty"`
	/*
	  Internal bridge name as br0
	*/
	InternalBridgeName *string `json:"internalBridgeName,omitempty"`

	IpAddress *IPv4Subnet `json:"ipAddress,omitempty"`
	/*
	  Internal route table number for the routing rules associated with the IP address on this host
	*/
	RouteTable *int `json:"routeTable,omitempty"`
}

func (p *Host) MarshalJSON() ([]byte, error) {
	type HostProxy Host
	return json.Marshal(struct {
		*HostProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		HostProxy: (*HostProxy)(p),
		ExtId:     p.ExtId,
	})
}

func NewHost() *Host {
	p := new(Host)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Host"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
ICMP parameters to be matched in routing policy.
*/
type ICMPObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ICMP code for the ICMP param to be matched in routing policy.
	*/
	IcmpCode *int `json:"icmpCode,omitempty"`
	/*
	  ICMP type for the ICMP param to be matched in routing policy.
	*/
	IcmpType *int `json:"icmpType,omitempty"`
}

func NewICMPObject() *ICMPObject {
	p := new(ICMPObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ICMPObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP configuration.
*/
type IPConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4 *IPv4Config `json:"ipv4,omitempty"`

	Ipv6 *IPv6Config `json:"ipv6,omitempty"`
}

func NewIPConfig() *IPConfig {
	p := new(IPConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (i *IPConfig) HasIpv4() bool {
	return i.Ipv4 != nil
}
func (i *IPConfig) HasIpv6() bool {
	return i.Ipv6 != nil
}

func (i *IPConfig) IsValid() bool {
	return i.HasIpv4() || i.HasIpv6()
}

type IPFIXExporter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The IP address of the IPFIX collector.
	*/
	CollectorIp *string `json:"collectorIp"`
	/*
	  The port number of the IPFIX collector.
	*/
	CollectorPort *int64 `json:"collectorPort"`
	/*
	  IPFIX exporter description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The maximum export rate in bits per second(bps) at which the exporter should try to export data.
	*/
	ExportRateLimitPerNodeBps *int64 `json:"exportRateLimitPerNodeBps,omitempty"`
	/*
	  List of IPFIX exporter scopes.
	*/
	ExportScopes []ExportScope `json:"exportScopes"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the IPFIX Exporter.
	*/
	Name *string `json:"name"`

	Protocol *ExporterProtocol `json:"protocol"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *IPFIXExporter) MarshalJSON() ([]byte, error) {
	type IPFIXExporterProxy IPFIXExporter
	return json.Marshal(struct {
		*IPFIXExporterProxy
		CollectorIp   *string           `json:"collectorIp,omitempty"`
		CollectorPort *int64            `json:"collectorPort,omitempty"`
		ExportScopes  []ExportScope     `json:"exportScopes,omitempty"`
		Name          *string           `json:"name,omitempty"`
		Protocol      *ExporterProtocol `json:"protocol,omitempty"`
	}{
		IPFIXExporterProxy: (*IPFIXExporterProxy)(p),
		CollectorIp:        p.CollectorIp,
		CollectorPort:      p.CollectorPort,
		ExportScopes:       p.ExportScopes,
		Name:               p.Name,
		Protocol:           p.Protocol,
	})
}

func NewIPFIXExporter() *IPFIXExporter {
	p := new(IPFIXExporter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPFIXExporter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP pool Usage.
*/
type IPPoolUsage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of free IPs.
	*/
	NumFreeIPs *int64 `json:"numFreeIPs,omitempty"`
	/*
	  Total number of IPs in this pool.
	*/
	NumTotalIPs *int64 `json:"numTotalIPs,omitempty"`

	Range *IPv4Pool `json:"range,omitempty"`
}

func NewIPPoolUsage() *IPPoolUsage {
	p := new(IPPoolUsage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPPoolUsage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP subnet.
*/
type IPSubnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4 *IPv4Subnet `json:"ipv4,omitempty"`

	Ipv6 *IPv6Subnet `json:"ipv6,omitempty"`
}

func NewIPSubnet() *IPSubnet {
	p := new(IPSubnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPSubnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (i *IPSubnet) HasIpv4() bool {
	return i.Ipv4 != nil
}
func (i *IPSubnet) HasIpv6() bool {
	return i.Ipv6 != nil
}

func (i *IPSubnet) IsValid() bool {
	return i.HasIpv4() || i.HasIpv6()
}

/*
IP usage statistics.This field is only returned in subnet GET response.
*/
type IPUsage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpPoolUsages []IPPoolUsage `json:"ipPoolUsages,omitempty"`
	/*
	  Number of assigned IPs.
	*/
	NumAssignedIPs *int64 `json:"numAssignedIPs,omitempty"`
	/*
	  Number of free IPs.
	*/
	NumFreeIPs *int64 `json:"numFreeIPs,omitempty"`
	/*
	  Number of MAC addresses.
	*/
	NumMacs *int64 `json:"numMacs,omitempty"`
}

func NewIPUsage() *IPUsage {
	p := new(IPUsage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPUsage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP V4 configuration.
*/
type IPv4Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGatewayIp *import1.IPv4Address `json:"defaultGatewayIp,omitempty"`

	DhcpServerAddress *import1.IPv4Address `json:"dhcpServerAddress,omitempty"`

	IpSubnet *IPv4Subnet `json:"ipSubnet"`
	/*
	  Pool of IP addresses from where IPs are allocated.
	*/
	PoolList []IPv4Pool `json:"poolList,omitempty"`
}

func (p *IPv4Config) MarshalJSON() ([]byte, error) {
	type IPv4ConfigProxy IPv4Config
	return json.Marshal(struct {
		*IPv4ConfigProxy
		IpSubnet *IPv4Subnet `json:"ipSubnet,omitempty"`
	}{
		IPv4ConfigProxy: (*IPv4ConfigProxy)(p),
		IpSubnet:        p.IpSubnet,
	})
}

func NewIPv4Config() *IPv4Config {
	p := new(IPv4Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv4Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Start/end IP address range.
*/
type IPv4Pool struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *import1.IPv4Address `json:"endIp"`

	StartIp *import1.IPv4Address `json:"startIp"`
}

func (p *IPv4Pool) MarshalJSON() ([]byte, error) {
	type IPv4PoolProxy IPv4Pool
	return json.Marshal(struct {
		*IPv4PoolProxy
		EndIp   *import1.IPv4Address `json:"endIp,omitempty"`
		StartIp *import1.IPv4Address `json:"startIp,omitempty"`
	}{
		IPv4PoolProxy: (*IPv4PoolProxy)(p),
		EndIp:         p.EndIp,
		StartIp:       p.StartIp,
	})
}

func NewIPv4Pool() *IPv4Pool {
	p := new(IPv4Pool)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv4Pool"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IPv4 subnet.
*/
type IPv4Subnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ip *import1.IPv4Address `json:"ip"`
	/*
	  Prefix length of the IPv4 subnet.
	*/
	PrefixLength *int `json:"prefixLength"`
}

func (p *IPv4Subnet) MarshalJSON() ([]byte, error) {
	type IPv4SubnetProxy IPv4Subnet
	return json.Marshal(struct {
		*IPv4SubnetProxy
		Ip           *import1.IPv4Address `json:"ip,omitempty"`
		PrefixLength *int                 `json:"prefixLength,omitempty"`
	}{
		IPv4SubnetProxy: (*IPv4SubnetProxy)(p),
		Ip:              p.Ip,
		PrefixLength:    p.PrefixLength,
	})
}

func NewIPv4Subnet() *IPv4Subnet {
	p := new(IPv4Subnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv4Subnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IP V6 configuration.
*/
type IPv6Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGatewayIp *import1.IPv6Address `json:"defaultGatewayIp,omitempty"`

	DhcpServerAddress *import1.IPv6Address `json:"dhcpServerAddress,omitempty"`

	IpSubnet *IPv6Subnet `json:"ipSubnet"`
	/*
	  Pool of IP addresses from where IPs are allocated.
	*/
	PoolList []IPv6Pool `json:"poolList,omitempty"`
}

func (p *IPv6Config) MarshalJSON() ([]byte, error) {
	type IPv6ConfigProxy IPv6Config
	return json.Marshal(struct {
		*IPv6ConfigProxy
		IpSubnet *IPv6Subnet `json:"ipSubnet,omitempty"`
	}{
		IPv6ConfigProxy: (*IPv6ConfigProxy)(p),
		IpSubnet:        p.IpSubnet,
	})
}

func NewIPv6Config() *IPv6Config {
	p := new(IPv6Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv6Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Start/end IP address range.
*/
type IPv6Pool struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *import1.IPv6Address `json:"endIp"`

	StartIp *import1.IPv6Address `json:"startIp"`
}

func (p *IPv6Pool) MarshalJSON() ([]byte, error) {
	type IPv6PoolProxy IPv6Pool
	return json.Marshal(struct {
		*IPv6PoolProxy
		EndIp   *import1.IPv6Address `json:"endIp,omitempty"`
		StartIp *import1.IPv6Address `json:"startIp,omitempty"`
	}{
		IPv6PoolProxy: (*IPv6PoolProxy)(p),
		EndIp:         p.EndIp,
		StartIp:       p.StartIp,
	})
}

func NewIPv6Pool() *IPv6Pool {
	p := new(IPv6Pool)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv6Pool"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IPv6 subnet.
*/
type IPv6Subnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ip *import1.IPv6Address `json:"ip"`
	/*
	  Prefix length of the IPv6 subnet.
	*/
	PrefixLength *int `json:"prefixLength"`
}

func (p *IPv6Subnet) MarshalJSON() ([]byte, error) {
	type IPv6SubnetProxy IPv6Subnet
	return json.Marshal(struct {
		*IPv6SubnetProxy
		Ip           *import1.IPv6Address `json:"ip,omitempty"`
		PrefixLength *int                 `json:"prefixLength,omitempty"`
	}{
		IPv6SubnetProxy: (*IPv6SubnetProxy)(p),
		Ip:              p.Ip,
		PrefixLength:    p.PrefixLength,
	})
}

func NewIPv6Subnet() *IPv6Subnet {
	p := new(IPv6Subnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IPv6Subnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Internal BGP configuration.
*/
type IbgpConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Autonomous system number. 0 and 4294967295 are reserved.
	*/
	Asn *int64 `json:"asn,omitempty"`
	/*
	  BPG password.
	*/
	Password *string `json:"password,omitempty"`

	PeerIp *import1.IPAddress `json:"peerIp,omitempty"`
	/*
	  Redistribute routes over eBGP. Applicable only to network gateways deployed on VLAN subnets with eBGP over VPN.
	*/
	ShouldRedistributeRoutes *bool `json:"shouldRedistributeRoutes,omitempty"`
}

func NewIbgpConfig() *IbgpConfig {
	p := new(IbgpConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IbgpConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldRedistributeRoutes = new(bool)
	*p.ShouldRedistributeRoutes = false

	return p
}

/*
Schema to configure IGMP Snooping on Virtual Switch
*/
type IgmpSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Enable IGMP snooping on this Virtual Switch
	*/
	IsSnoopingEnabled *bool `json:"isSnoopingEnabled,omitempty"`

	QuerierSpec *QuerierSpec `json:"querierSpec,omitempty"`
	/*
	  IGMP Snooping timeout value in seconds
	*/
	SnoopingTimeout *int64 `json:"snoopingTimeout,omitempty"`
}

func NewIgmpSpec() *IgmpSpec {
	p := new(IgmpSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IgmpSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSnoopingEnabled = new(bool)
	*p.IsSnoopingEnabled = false
	p.SnoopingTimeout = new(int64)
	*p.SnoopingTimeout = 300

	return p
}

/*
Describes the routing protocol configuration spec needed by this gateway to peer and learn routes from internal routers using either iBGP or OSPF.
*/
type InternalRoutingConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  iBGP configuration.
	*/
	IbgpConfigList []IbgpConfig `json:"ibgpConfigList,omitempty"`
	/*
	  List of local prefixes to be advertised over eBGP.
	*/
	LocalPrefixList []IPSubnet `json:"localPrefixList,omitempty"`

	OspfConfig *OspfConfig `json:"ospfConfig,omitempty"`
}

func NewInternalRoutingConfig() *InternalRoutingConfig {
	p := new(InternalRoutingConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.InternalRoutingConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input required to reserve IP addresses on a subnet.
*/
type IpReserveSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Optional context to associate with a reserved IP address.
	*/
	ClientContext *string `json:"clientContext,omitempty"`
	/*
	  Number of IP addresses reserved.
	*/
	Count *int64 `json:"count,omitempty"`
	/*
	  List of IP addresses to reserve.
	*/
	IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`

	ReserveType *ReserveType `json:"reserveType"`

	StartIpAddress *import1.IPAddress `json:"startIpAddress,omitempty"`
}

func (p *IpReserveSpec) MarshalJSON() ([]byte, error) {
	type IpReserveSpecProxy IpReserveSpec
	return json.Marshal(struct {
		*IpReserveSpecProxy
		ReserveType *ReserveType `json:"reserveType,omitempty"`
	}{
		IpReserveSpecProxy: (*IpReserveSpecProxy)(p),
		ReserveType:        p.ReserveType,
	})
}

func NewIpReserveSpec() *IpReserveSpec {
	p := new(IpReserveSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IpReserveSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input required to unreserve IP addresses on a subnet.
*/
type IpUnreserveSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Optional context to associate with a reserved IP address.
	*/
	ClientContext *string `json:"clientContext,omitempty"`
	/*
	  Number of IP addresses unreserved.
	*/
	Count *int64 `json:"count,omitempty"`
	/*
	  List of IP addresses to unreserve.
	*/
	IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`

	StartIpAddress *import1.IPAddress `json:"startIpAddress,omitempty"`

	UnreserveType *UnreserveType `json:"unreserveType"`
}

func (p *IpUnreserveSpec) MarshalJSON() ([]byte, error) {
	type IpUnreserveSpecProxy IpUnreserveSpec
	return json.Marshal(struct {
		*IpUnreserveSpecProxy
		UnreserveType *UnreserveType `json:"unreserveType,omitempty"`
	}{
		IpUnreserveSpecProxy: (*IpUnreserveSpecProxy)(p),
		UnreserveType:        p.UnreserveType,
	})
}

func NewIpUnreserveSpec() *IpUnreserveSpec {
	p := new(IpUnreserveSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IpUnreserveSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IPSec configuration
*/
type IpsecConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Diffie-Hellman group value of 14, 19 or 20 to be used for Perfect Forward Secrecy (PFS)
	*/
	EspPfsDhGroupNumber *int `json:"espPfsDhGroupNumber,omitempty"`

	IkeAuthenticationAlgorithm *AuthenticationAlgorithm `json:"ikeAuthenticationAlgorithm,omitempty"`

	IkeEncryptionAlgorithm *EncryptionAlgorithm `json:"ikeEncryptionAlgorithm,omitempty"`
	/*
	  IKE lifetime (seconds)
	*/
	IkeLifetimeSecs *int64 `json:"ikeLifetimeSecs,omitempty"`

	IpsecAuthenticationAlgorithm *AuthenticationAlgorithm `json:"ipsecAuthenticationAlgorithm,omitempty"`

	IpsecEncryptionAlgorithm *EncryptionAlgorithm `json:"ipsecEncryptionAlgorithm,omitempty"`
	/*
	  IPSec lifetime (seconds)
	*/
	IpsecLifetimeSecs *int64 `json:"ipsecLifetimeSecs,omitempty"`
	/*
	  Local IKE authentication Id used for this connection
	*/
	LocalAuthenticationId *string `json:"localAuthenticationId,omitempty"`

	LocalVtiIp *import1.IPAddress `json:"localVtiIp,omitempty"`
	/*
	  Shared secret for authentication between gateway peers
	*/
	PreSharedKey *string `json:"preSharedKey"`
	/*
	  IKE authentication Id of the remote peer
	*/
	RemoteAuthenticationId *string `json:"remoteAuthenticationId,omitempty"`

	RemoteVtiIp *import1.IPAddress `json:"remoteVtiIp,omitempty"`
}

func (p *IpsecConfig) MarshalJSON() ([]byte, error) {
	type IpsecConfigProxy IpsecConfig
	return json.Marshal(struct {
		*IpsecConfigProxy
		PreSharedKey *string `json:"preSharedKey,omitempty"`
	}{
		IpsecConfigProxy: (*IpsecConfigProxy)(p),
		PreSharedKey:     p.PreSharedKey,
	})
}

func NewIpsecConfig() *IpsecConfig {
	p := new(IpsecConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.IpsecConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Layer2Stretch struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ConnectionType *StretchConnectionType `json:"connectionType,omitempty"`
	/*
	  Layer2 stretch configuration details between subnets on two sites.
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

	LocalSiteParams *SiteParams `json:"localSiteParams,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  The MTU size setting for the VXLAN session.
	*/
	Mtu *int `json:"mtu,omitempty"`
	/*
	  Layer2 stretch configuration name.
	*/
	Name *string `json:"name"`

	RemoteSiteParams *SiteParams `json:"remoteSiteParams,omitempty"`

	RemoteStretchStatus []RemoteVtepStretchStatus `json:"remoteStretchStatus,omitempty"`

	StretchStatus *StretchStatus `json:"stretchStatus,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The VXLAN network identifier used to uniquely identify the VXLAN tunnel.
	*/
	Vni *int `json:"vni,omitempty"`
}

func (p *Layer2Stretch) MarshalJSON() ([]byte, error) {
	type Layer2StretchProxy Layer2Stretch
	return json.Marshal(struct {
		*Layer2StretchProxy
		Name *string `json:"name,omitempty"`
	}{
		Layer2StretchProxy: (*Layer2StretchProxy)(p),
		Name:               p.Name,
	})
}

func NewLayer2Stretch() *Layer2Stretch {
	p := new(Layer2Stretch)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Layer2Stretch"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
L4 TCP/UDP protocol parameters to be matched in routing policy.
*/
type LayerFourProtocolObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DestinationPortRanges []PortRange `json:"destinationPortRanges,omitempty"`

	SourcePortRanges []PortRange `json:"sourcePortRanges,omitempty"`
}

func NewLayerFourProtocolObject() *LayerFourProtocolObject {
	p := new(LayerFourProtocolObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LayerFourProtocolObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Layer2 Stretch MAC address properties.
*/
type LearnedMacAddress struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	MacType *MacType `json:"macType,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A reference to the associated remote VTEP Gateway.
	*/
	RemoteGatewayReference *string `json:"remoteGatewayReference,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Hexadecimal representation of this MAC address.
	*/
	Value *string `json:"value,omitempty"`

	VtepIPAddress *import1.IPAddress `json:"vtepIPAddress,omitempty"`
}

func NewLearnedMacAddress() *LearnedMacAddress {
	p := new(LearnedMacAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LearnedMacAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0/config/bgp-sessions/{bgpSessionExtId}/bgp-routes Get operation
*/
type ListBgpRoutesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListBgpRoutesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListBgpRoutesApiResponse() *ListBgpRoutesApiResponse {
	p := new(ListBgpRoutesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListBgpRoutesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListBgpRoutesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListBgpRoutesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListBgpRoutesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/bgp-sessions Get operation
*/
type ListBgpSessionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListBgpSessionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListBgpSessionsApiResponse() *ListBgpSessionsApiResponse {
	p := new(ListBgpSessionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListBgpSessionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListBgpSessionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListBgpSessionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListBgpSessionsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/capabilities Get operation
*/
type ListClusterCapabilitiesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClusterCapabilitiesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListClusterCapabilitiesApiResponse() *ListClusterCapabilitiesApiResponse {
	p := new(ListClusterCapabilitiesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListClusterCapabilitiesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClusterCapabilitiesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClusterCapabilitiesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClusterCapabilitiesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/floating-ips Get operation
*/
type ListFloatingIpsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListFloatingIpsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListFloatingIpsApiResponse() *ListFloatingIpsApiResponse {
	p := new(ListFloatingIpsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListFloatingIpsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListFloatingIpsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListFloatingIpsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListFloatingIpsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/gateways Get operation
*/
type ListGatewaysApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListGatewaysApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListGatewaysApiResponse() *ListGatewaysApiResponse {
	p := new(ListGatewaysApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListGatewaysApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListGatewaysApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListGatewaysApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListGatewaysApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/ipfix-exporters Get operation
*/
type ListIPFIXExportersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListIPFIXExportersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListIPFIXExportersApiResponse() *ListIPFIXExportersApiResponse {
	p := new(ListIPFIXExportersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListIPFIXExportersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListIPFIXExportersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListIPFIXExportersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListIPFIXExportersApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/layer2-stretches Get operation
*/
type ListLayer2StretchesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLayer2StretchesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLayer2StretchesApiResponse() *ListLayer2StretchesApiResponse {
	p := new(ListLayer2StretchesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListLayer2StretchesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLayer2StretchesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLayer2StretchesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLayer2StretchesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/layer2-stretches/{layer2StretchExtId}/learned-mac-addresses Get operation
*/
type ListLearnedMacAddressesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLearnedMacAddressesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLearnedMacAddressesApiResponse() *ListLearnedMacAddressesApiResponse {
	p := new(ListLearnedMacAddressesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListLearnedMacAddressesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLearnedMacAddressesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLearnedMacAddressesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLearnedMacAddressesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/load-balancer-sessions Get operation
*/
type ListLoadBalancerSessionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLoadBalancerSessionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLoadBalancerSessionsApiResponse() *ListLoadBalancerSessionsApiResponse {
	p := new(ListLoadBalancerSessionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListLoadBalancerSessionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLoadBalancerSessionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLoadBalancerSessionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLoadBalancerSessionsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/controllers Get operation
*/
type ListNetworkControllersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNetworkControllersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNetworkControllersApiResponse() *ListNetworkControllersApiResponse {
	p := new(ListNetworkControllersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListNetworkControllersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNetworkControllersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNetworkControllersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNetworkControllersApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/node-schedulable-statuses Get operation
*/
type ListNodeSchedulableStatusesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListNodeSchedulableStatusesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListNodeSchedulableStatusesApiResponse() *ListNodeSchedulableStatusesApiResponse {
	p := new(ListNodeSchedulableStatusesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListNodeSchedulableStatusesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListNodeSchedulableStatusesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListNodeSchedulableStatusesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListNodeSchedulableStatusesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-subnets Get operation
*/
type ListRemoteSubnetsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRemoteSubnetsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRemoteSubnetsApiResponse() *ListRemoteSubnetsApiResponse {
	p := new(ListRemoteSubnetsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListRemoteSubnetsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRemoteSubnetsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRemoteSubnetsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRemoteSubnetsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/route-tables Get operation
*/
type ListRouteTablesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRouteTablesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRouteTablesApiResponse() *ListRouteTablesApiResponse {
	p := new(ListRouteTablesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListRouteTablesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRouteTablesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRouteTablesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRouteTablesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/route-tables/{routeTableExtId}/routes Get operation
*/
type ListRoutesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRoutesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRoutesApiResponse() *ListRoutesApiResponse {
	p := new(ListRoutesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListRoutesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRoutesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRoutesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRoutesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/routing-policies Get operation
*/
type ListRoutingPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRoutingPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRoutingPoliciesApiResponse() *ListRoutingPoliciesApiResponse {
	p := new(ListRoutingPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListRoutingPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRoutingPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRoutingPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRoutingPoliciesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/subnets/{subnetExtId}/reserved-ips Get operation
*/
type ListSubnetReservedIpsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSubnetReservedIpsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSubnetReservedIpsApiResponse() *ListSubnetReservedIpsApiResponse {
	p := new(ListSubnetReservedIpsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListSubnetReservedIpsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSubnetReservedIpsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSubnetReservedIpsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSubnetReservedIpsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/subnets/{subnetExtId}/vnics Get operation
*/
type ListSubnetVnicsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSubnetVnicsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSubnetVnicsApiResponse() *ListSubnetVnicsApiResponse {
	p := new(ListSubnetVnicsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListSubnetVnicsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSubnetVnicsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSubnetVnicsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSubnetVnicsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/subnets Get operation
*/
type ListSubnetsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSubnetsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSubnetsApiResponse() *ListSubnetsApiResponse {
	p := new(ListSubnetsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListSubnetsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSubnetsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSubnetsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSubnetsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/traffic-mirrors Get operation
*/
type ListTrafficMirrorsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListTrafficMirrorsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListTrafficMirrorsApiResponse() *ListTrafficMirrorsApiResponse {
	p := new(ListTrafficMirrorsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListTrafficMirrorsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListTrafficMirrorsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListTrafficMirrorsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListTrafficMirrorsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/uplink-bonds Get operation
*/
type ListUplinkBondsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUplinkBondsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUplinkBondsApiResponse() *ListUplinkBondsApiResponse {
	p := new(ListUplinkBondsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListUplinkBondsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUplinkBondsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUplinkBondsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUplinkBondsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/virtual-switches Get operation
*/
type ListVirtualSwitchesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVirtualSwitchesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVirtualSwitchesApiResponse() *ListVirtualSwitchesApiResponse {
	p := new(ListVirtualSwitchesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListVirtualSwitchesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVirtualSwitchesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVirtualSwitchesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVirtualSwitchesApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpc-virtual-switch-mappings Get operation
*/
type ListVpcVirtualSwitchMappingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVpcVirtualSwitchMappingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVpcVirtualSwitchMappingsApiResponse() *ListVpcVirtualSwitchMappingsApiResponse {
	p := new(ListVpcVirtualSwitchMappingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListVpcVirtualSwitchMappingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVpcVirtualSwitchMappingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVpcVirtualSwitchMappingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVpcVirtualSwitchMappingsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpcs Get operation
*/
type ListVpcsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVpcsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVpcsApiResponse() *ListVpcsApiResponse {
	p := new(ListVpcsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListVpcsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVpcsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVpcsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVpcsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpn-connections Get operation
*/
type ListVpnConnectionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVpnConnectionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVpnConnectionsApiResponse() *ListVpnConnectionsApiResponse {
	p := new(ListVpnConnectionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListVpnConnectionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVpnConnectionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVpnConnectionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVpnConnectionsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/vpn-connections/{vpnConnectionExtId}/vpn-vendor-configs Get operation
*/
type ListVpnVendorConfigsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVpnVendorConfigsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVpnVendorConfigsApiResponse() *ListVpnVendorConfigsApiResponse {
	p := new(ListVpnVendorConfigsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ListVpnVendorConfigsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVpnVendorConfigsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVpnVendorConfigsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVpnVendorConfigsApiResponseData()
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
Listener configuration for the load balancer session.
*/
type Listener struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Doc(LoadBalancerSessionPortRange)
	*/
	PortRanges []PortRange `json:"portRanges"`

	Protocol *Protocol `json:"protocol"`

	VirtualIP *VirtualIP `json:"virtualIP"`
}

func (p *Listener) MarshalJSON() ([]byte, error) {
	type ListenerProxy Listener
	return json.Marshal(struct {
		*ListenerProxy
		PortRanges []PortRange `json:"portRanges,omitempty"`
		Protocol   *Protocol   `json:"protocol,omitempty"`
		VirtualIP  *VirtualIP  `json:"virtualIP,omitempty"`
	}{
		ListenerProxy: (*ListenerProxy)(p),
		PortRanges:    p.PortRanges,
		Protocol:      p.Protocol,
		VirtualIP:     p.VirtualIP,
	})
}

func NewListener() *Listener {
	p := new(Listener)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Listener"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type LoadBalancerSession struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Algorithm *Algorithm `json:"algorithm,omitempty"`
	/*
	  Description of the load balancer session.
	*/
	Description *string `json:"description"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HealthCheckConfig *HealthCheck `json:"healthCheckConfig"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Listener *Listener `json:"listener"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the load balancer session.
	*/
	Name *string `json:"name"`

	TargetsConfig *Target `json:"targetsConfig"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LoadBalancerSessionType `json:"type,omitempty"`
	/*
	  UUID of the Virtual Private Cloud this load balancer session belongs to.
	*/
	VpcReference *string `json:"vpcReference"`
}

func (p *LoadBalancerSession) MarshalJSON() ([]byte, error) {
	type LoadBalancerSessionProxy LoadBalancerSession
	return json.Marshal(struct {
		*LoadBalancerSessionProxy
		Description       *string      `json:"description,omitempty"`
		HealthCheckConfig *HealthCheck `json:"healthCheckConfig,omitempty"`
		Listener          *Listener    `json:"listener,omitempty"`
		Name              *string      `json:"name,omitempty"`
		TargetsConfig     *Target      `json:"targetsConfig,omitempty"`
		VpcReference      *string      `json:"vpcReference,omitempty"`
	}{
		LoadBalancerSessionProxy: (*LoadBalancerSessionProxy)(p),
		Description:              p.Description,
		HealthCheckConfig:        p.HealthCheckConfig,
		Listener:                 p.Listener,
		Name:                     p.Name,
		TargetsConfig:            p.TargetsConfig,
		VpcReference:             p.VpcReference,
	})
}

func NewLoadBalancerSession() *LoadBalancerSession {
	p := new(LoadBalancerSession)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LoadBalancerSession"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Load Balancer Session to which the floating IP is associated.
*/
type LoadBalancerSessionAssociation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Load Balancer Session reference UUID.
	*/
	LoadBalancerSessionReference *string `json:"loadBalancerSessionReference"`
}

func (p *LoadBalancerSessionAssociation) MarshalJSON() ([]byte, error) {
	type LoadBalancerSessionAssociationProxy LoadBalancerSessionAssociation
	return json.Marshal(struct {
		*LoadBalancerSessionAssociationProxy
		LoadBalancerSessionReference *string `json:"loadBalancerSessionReference,omitempty"`
	}{
		LoadBalancerSessionAssociationProxy: (*LoadBalancerSessionAssociationProxy)(p),
		LoadBalancerSessionReference:        p.LoadBalancerSessionReference,
	})
}

func NewLoadBalancerSessionAssociation() *LoadBalancerSessionAssociation {
	p := new(LoadBalancerSessionAssociation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LoadBalancerSessionAssociation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type LoadBalancerSessionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Algorithm *Algorithm `json:"algorithm,omitempty"`
	/*
	  Description of the load balancer session.
	*/
	Description *string `json:"description"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HealthCheckConfig *HealthCheck `json:"healthCheckConfig"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Listener *Listener `json:"listener"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the load balancer session.
	*/
	Name *string `json:"name"`

	TargetsConfig *Target `json:"targetsConfig"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LoadBalancerSessionType `json:"type,omitempty"`
	/*
	  UUID of the Virtual Private Cloud this load balancer session belongs to.
	*/
	VpcReference *string `json:"vpcReference"`
}

func (p *LoadBalancerSessionProjection) MarshalJSON() ([]byte, error) {
	type LoadBalancerSessionProjectionProxy LoadBalancerSessionProjection
	return json.Marshal(struct {
		*LoadBalancerSessionProjectionProxy
		Description       *string      `json:"description,omitempty"`
		HealthCheckConfig *HealthCheck `json:"healthCheckConfig,omitempty"`
		Listener          *Listener    `json:"listener,omitempty"`
		Name              *string      `json:"name,omitempty"`
		TargetsConfig     *Target      `json:"targetsConfig,omitempty"`
		VpcReference      *string      `json:"vpcReference,omitempty"`
	}{
		LoadBalancerSessionProjectionProxy: (*LoadBalancerSessionProjectionProxy)(p),
		Description:                        p.Description,
		HealthCheckConfig:                  p.HealthCheckConfig,
		Listener:                           p.Listener,
		Name:                               p.Name,
		TargetsConfig:                      p.TargetsConfig,
		VpcReference:                       p.VpcReference,
	})
}

func NewLoadBalancerSessionProjection() *LoadBalancerSessionProjection {
	p := new(LoadBalancerSessionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LoadBalancerSessionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of load balancer session.
*/
type LoadBalancerSessionType int

const (
	LOADBALANCERSESSIONTYPE_UNKNOWN               LoadBalancerSessionType = 0
	LOADBALANCERSESSIONTYPE_REDACTED              LoadBalancerSessionType = 1
	LOADBALANCERSESSIONTYPE_NETWORK_LOAD_BALANCER LoadBalancerSessionType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LoadBalancerSessionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NETWORK_LOAD_BALANCER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LoadBalancerSessionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NETWORK_LOAD_BALANCER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LoadBalancerSessionType) index(name string) LoadBalancerSessionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NETWORK_LOAD_BALANCER",
	}
	for idx := range names {
		if names[idx] == name {
			return LoadBalancerSessionType(idx)
		}
	}
	return LOADBALANCERSESSIONTYPE_UNKNOWN
}

func (e *LoadBalancerSessionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoadBalancerSessionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoadBalancerSessionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoadBalancerSessionType) Ref() *LoadBalancerSessionType {
	return &e
}

/*
Local BGP gateway info needed for flow gateway scale out model.
*/
type LocalBgpGateway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ASN number of the BGP gateway.
	*/
	Asn *int `json:"asn"`
	/*
	  Prefix length of the VNIC IP addresses of the local BGP gateways.
	*/
	VnicIpPrefixLength *int `json:"vnicIpPrefixLength"`

	VnicList []GatewayNic `json:"vnicList"`
}

func (p *LocalBgpGateway) MarshalJSON() ([]byte, error) {
	type LocalBgpGatewayProxy LocalBgpGateway
	return json.Marshal(struct {
		*LocalBgpGatewayProxy
		Asn                *int         `json:"asn,omitempty"`
		VnicIpPrefixLength *int         `json:"vnicIpPrefixLength,omitempty"`
		VnicList           []GatewayNic `json:"vnicList,omitempty"`
	}{
		LocalBgpGatewayProxy: (*LocalBgpGatewayProxy)(p),
		Asn:                  p.Asn,
		VnicIpPrefixLength:   p.VnicIpPrefixLength,
		VnicList:             p.VnicList,
	})
}

func NewLocalBgpGateway() *LocalBgpGateway {
	p := new(LocalBgpGateway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LocalBgpGateway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
BGP service hosted on this local gateway.
*/
type LocalBgpService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Autonomous system number. 0 and 4294967295 are reserved.
	*/
	Asn *int64 `json:"asn"`
	/*
	  Reference to the VPC that this network gateway serves as its BGP speaker. Note that this BGP gateway will not service a regular VPC sitting behind a Transit VPC.
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *LocalBgpService) MarshalJSON() ([]byte, error) {
	type LocalBgpServiceProxy LocalBgpService
	return json.Marshal(struct {
		*LocalBgpServiceProxy
		Asn *int64 `json:"asn,omitempty"`
	}{
		LocalBgpServiceProxy: (*LocalBgpServiceProxy)(p),
		Asn:                  p.Asn,
	})
}

func NewLocalBgpService() *LocalBgpService {
	p := new(LocalBgpService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LocalBgpService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Services of this local gateway.
*/
type LocalNetworkServices struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LocalBgpService *LocalBgpService `json:"localBgpService,omitempty"`

	LocalVpnService *LocalVpnService `json:"localVpnService,omitempty"`

	LocalVtepService *LocalVtepService `json:"localVtepService,omitempty"`

	ServiceAddress *import1.IPAddress `json:"serviceAddress,omitempty"`
	/*
	  List of floating IP addresses associated with this local gateway.
	*/
	ServiceAddresses []import1.IPAddress `json:"serviceAddresses,omitempty"`
}

func NewLocalNetworkServices() *LocalNetworkServices {
	p := new(LocalNetworkServices)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LocalNetworkServices"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPN service hosted on this local gateway.
*/
type LocalVpnService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EbgpConfig *BgpConfig `json:"ebgpConfig,omitempty"`

	PeerIgpConfig *InternalRoutingConfig `json:"peerIgpConfig,omitempty"`
}

func NewLocalVpnService() *LocalVpnService {
	p := new(LocalVpnService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LocalVpnService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VTEP service hosted on this local gateway.
*/
type LocalVtepService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VXLAN port.
	*/
	VxlanPort *int `json:"vxlanPort"`
}

func (p *LocalVtepService) MarshalJSON() ([]byte, error) {
	type LocalVtepServiceProxy LocalVtepService
	return json.Marshal(struct {
		*LocalVtepServiceProxy
		VxlanPort *int `json:"vxlanPort,omitempty"`
	}{
		LocalVtepServiceProxy: (*LocalVtepServiceProxy)(p),
		VxlanPort:             p.VxlanPort,
	})
}

func NewLocalVtepService() *LocalVtepService {
	p := new(LocalVtepService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.LocalVtepService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
MAC address.
*/
type MacAddress struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	MacType *MacType `json:"macType,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Hexadecimal representation of this MAC address.
	*/
	Value *string `json:"value,omitempty"`
}

func NewMacAddress() *MacAddress {
	p := new(MacAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.MacAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
MAC address type.
*/
type MacType int

const (
	MACTYPE_UNKNOWN       MacType = 0
	MACTYPE_REDACTED      MacType = 1
	MACTYPE_LEARNED       MacType = 2
	MACTYPE_AUTO_ASSIGNED MacType = 3
	MACTYPE_USER_ASSIGNED MacType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MacType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEARNED",
		"AUTO_ASSIGNED",
		"USER_ASSIGNED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MacType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEARNED",
		"AUTO_ASSIGNED",
		"USER_ASSIGNED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MacType) index(name string) MacType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEARNED",
		"AUTO_ASSIGNED",
		"USER_ASSIGNED",
	}
	for idx := range names {
		if names[idx] == name {
			return MacType(idx)
		}
	}
	return MACTYPE_UNKNOWN
}

func (e *MacType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MacType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MacType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MacType) Ref() *MacType {
	return &e
}

/*
Migration state of the subnet. This field is read-only.
*/
type MigrationState int

const (
	MIGRATIONSTATE_UNKNOWN     MigrationState = 0
	MIGRATIONSTATE_REDACTED    MigrationState = 1
	MIGRATIONSTATE_IN_PROGRESS MigrationState = 2
	MIGRATIONSTATE_FAILED      MigrationState = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MigrationState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MigrationState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MigrationState) index(name string) MigrationState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return MigrationState(idx)
		}
	}
	return MIGRATIONSTATE_UNKNOWN
}

func (e *MigrationState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MigrationState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MigrationState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MigrationState) Ref() *MigrationState {
	return &e
}

type NetworkCloudConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AwsConfig *AwsConfig `json:"awsConfig,omitempty"`

	AzureConfig *AzureConfig `json:"azureConfig,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkCloudConfig() *NetworkCloudConfig {
	p := new(NetworkCloudConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.NetworkCloudConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NetworkController struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CloudSubstrate *CloudSubstrate `json:"cloudSubstrate,omitempty"`

	ControllerStatus *ControllerStatus `json:"controllerStatus,omitempty"`
	/*
	  Network controller software version.
	*/
	ControllerVersion *string `json:"controllerVersion,omitempty"`

	DefaultVlanStack *DefaultVlanStack `json:"defaultVlanStack,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Minimum AHV version that this network controller supports.
	*/
	MinimumAHVVersion *string `json:"minimumAHVVersion,omitempty"`
	/*
	  Minimum NOS version that this network controller supports.
	*/
	MinimumNOSVersion *string `json:"minimumNOSVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkController() *NetworkController {
	p := new(NetworkController)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.NetworkController"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Networking common base object
*/
type NetworkingBaseModel struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNetworkingBaseModel() *NetworkingBaseModel {
	p := new(NetworkingBaseModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.NetworkingBaseModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Route nexthop.
*/
type Nexthop struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NexthopIpAddress *import1.IPAddress `json:"nexthopIpAddress,omitempty"`
	/*
	  Name of the nexthop, where the nexthop is either an IP address, a VPN connection, or a subnet.
	*/
	NexthopName *string `json:"nexthopName,omitempty"`
	/*
	  The reference to a link, such as a VPN connection or a subnet.
	*/
	NexthopReference *string `json:"nexthopReference,omitempty"`

	NexthopType *NexthopType `json:"nexthopType"`
}

func (p *Nexthop) MarshalJSON() ([]byte, error) {
	type NexthopProxy Nexthop
	return json.Marshal(struct {
		*NexthopProxy
		NexthopType *NexthopType `json:"nexthopType,omitempty"`
	}{
		NexthopProxy: (*NexthopProxy)(p),
		NexthopType:  p.NexthopType,
	})
}

func NewNexthop() *Nexthop {
	p := new(Nexthop)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Nexthop"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Nexthop type.
*/
type NexthopType int

const (
	NEXTHOPTYPE_UNKNOWN            NexthopType = 0
	NEXTHOPTYPE_REDACTED           NexthopType = 1
	NEXTHOPTYPE_IP_ADDRESS         NexthopType = 2
	NEXTHOPTYPE_DIRECT_CONNECT_VIF NexthopType = 3
	NEXTHOPTYPE_LOCAL_SUBNET       NexthopType = 4
	NEXTHOPTYPE_EXTERNAL_SUBNET    NexthopType = 5
	NEXTHOPTYPE_VPN_CONNECTION     NexthopType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NexthopType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS",
		"DIRECT_CONNECT_VIF",
		"LOCAL_SUBNET",
		"EXTERNAL_SUBNET",
		"VPN_CONNECTION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NexthopType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS",
		"DIRECT_CONNECT_VIF",
		"LOCAL_SUBNET",
		"EXTERNAL_SUBNET",
		"VPN_CONNECTION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NexthopType) index(name string) NexthopType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS",
		"DIRECT_CONNECT_VIF",
		"LOCAL_SUBNET",
		"EXTERNAL_SUBNET",
		"VPN_CONNECTION",
	}
	for idx := range names {
		if names[idx] == name {
			return NexthopType(idx)
		}
	}
	return NEXTHOPTYPE_UNKNOWN
}

func (e *NexthopType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NexthopType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NexthopType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NexthopType) Ref() *NexthopType {
	return &e
}

/*
Targets that are associated with the load balancer using virtual NIC.
*/
type NicTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Health *TargetHealth `json:"health,omitempty"`
	/*
	  The port value of the load balancer session target.
	*/
	Port *int `json:"port,omitempty"`
	/*
	  Reference to the virtual NIC of the load balancer session target.
	*/
	VirtualNicReference *string `json:"virtualNicReference,omitempty"`
	/*
	  Reference to the VM of the load balancer session target.
	*/
	VmReference *string `json:"vmReference,omitempty"`
}

func NewNicTarget() *NicTarget {
	p := new(NicTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.NicTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Array of node UUIDs and boolean pairs, indicating whether the nodes are storage-only.
*/
type NodeSchedulableStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The boolean value to indicate whether or not node is a storage only node
	*/
	IsNeverSchedulable *bool `json:"isNeverSchedulable,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNodeSchedulableStatus() *NodeSchedulableStatus {
	p := new(NodeSchedulableStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.NodeSchedulableStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
OSPF configuration for route peering with internal routers.
*/
type OspfConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  OSPF area id of this gateway.
	*/
	AreaId *string `json:"areaId,omitempty"`

	AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`
	/*
	  Password for authentication.
	*/
	Password *string `json:"password,omitempty"`
}

func NewOspfConfig() *OspfConfig {
	p := new(OspfConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.OspfConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Owner of DVS
*/
type OwnerType int

const (
	OWNERTYPE_UNKNOWN  OwnerType = 0
	OWNERTYPE_REDACTED OwnerType = 1
	OWNERTYPE_PE       OwnerType = 2
	OWNERTYPE_PC       OwnerType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OwnerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PE",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OwnerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PE",
		"PC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OwnerType) index(name string) OwnerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PE",
		"PC",
	}
	for idx := range names {
		if names[idx] == name {
			return OwnerType(idx)
		}
	}
	return OWNERTYPE_UNKNOWN
}

func (e *OwnerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OwnerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OwnerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OwnerType) Ref() *OwnerType {
	return &e
}

/*
Range of TCP/UDP ports.
*/
type PortRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  End port of TCP/UDP port range.
	*/
	EndPort *int `json:"endPort"`
	/*
	  Start port of TCP/UDP port range.
	*/
	StartPort *int `json:"startPort"`
}

func (p *PortRange) MarshalJSON() ([]byte, error) {
	type PortRangeProxy PortRange
	return json.Marshal(struct {
		*PortRangeProxy
		EndPort   *int `json:"endPort,omitempty"`
		StartPort *int `json:"startPort,omitempty"`
	}{
		PortRangeProxy: (*PortRangeProxy)(p),
		EndPort:        p.EndPort,
		StartPort:      p.StartPort,
	})
}

func NewPortRange() *PortRange {
	p := new(PortRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.PortRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Private IP and VPC to which the floating IP is associated.
*/
type PrivateIpAssociation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PrivateIp *import1.IPAddress `json:"privateIp"`
	/*
	  VPC in which the private IP exists.
	*/
	VpcReference *string `json:"vpcReference"`
}

func (p *PrivateIpAssociation) MarshalJSON() ([]byte, error) {
	type PrivateIpAssociationProxy PrivateIpAssociation
	return json.Marshal(struct {
		*PrivateIpAssociationProxy
		PrivateIp    *import1.IPAddress `json:"privateIp,omitempty"`
		VpcReference *string            `json:"vpcReference,omitempty"`
	}{
		PrivateIpAssociationProxy: (*PrivateIpAssociationProxy)(p),
		PrivateIp:                 p.PrivateIp,
		VpcReference:              p.VpcReference,
	})
}

func NewPrivateIpAssociation() *PrivateIpAssociation {
	p := new(PrivateIpAssociation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.PrivateIpAssociation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
L3/L4 protocol.
*/
type Protocol int

const (
	PROTOCOL_UNKNOWN  Protocol = 0
	PROTOCOL_REDACTED Protocol = 1
	PROTOCOL_TCP      Protocol = 2
	PROTOCOL_UDP      Protocol = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Protocol) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Protocol) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Protocol) index(name string) Protocol {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"TCP",
		"UDP",
	}
	for idx := range names {
		if names[idx] == name {
			return Protocol(idx)
		}
	}
	return PROTOCOL_UNKNOWN
}

func (e *Protocol) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Protocol:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Protocol) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Protocol) Ref() *Protocol {
	return &e
}

type ProtocolNumberObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Routing Policy IP protocol parameter number.
	*/
	ProtocolNumber *int `json:"protocolNumber"`
}

func (p *ProtocolNumberObject) MarshalJSON() ([]byte, error) {
	type ProtocolNumberObjectProxy ProtocolNumberObject
	return json.Marshal(struct {
		*ProtocolNumberObjectProxy
		ProtocolNumber *int `json:"protocolNumber,omitempty"`
	}{
		ProtocolNumberObjectProxy: (*ProtocolNumberObjectProxy)(p),
		ProtocolNumber:            p.ProtocolNumber,
	})
}

func NewProtocolNumberObject() *ProtocolNumberObject {
	p := new(ProtocolNumberObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ProtocolNumberObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Routing Policy IP protocol type.
*/
type ProtocolType int

const (
	PROTOCOLTYPE_UNKNOWN         ProtocolType = 0
	PROTOCOLTYPE_REDACTED        ProtocolType = 1
	PROTOCOLTYPE_ANY             ProtocolType = 2
	PROTOCOLTYPE_ICMP            ProtocolType = 3
	PROTOCOLTYPE_TCP             ProtocolType = 4
	PROTOCOLTYPE_UDP             ProtocolType = 5
	PROTOCOLTYPE_PROTOCOL_NUMBER ProtocolType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtocolType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"ICMP",
		"TCP",
		"UDP",
		"PROTOCOL_NUMBER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProtocolType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"ICMP",
		"TCP",
		"UDP",
		"PROTOCOL_NUMBER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProtocolType) index(name string) ProtocolType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ANY",
		"ICMP",
		"TCP",
		"UDP",
		"PROTOCOL_NUMBER",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtocolType(idx)
		}
	}
	return PROTOCOLTYPE_UNKNOWN
}

func (e *ProtocolType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtocolType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtocolType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtocolType) Ref() *ProtocolType {
	return &e
}

type PublicIpMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PrivateIp *import1.IPAddress `json:"privateIp"`

	PublicIp *import1.IPAddress `json:"publicIp"`
}

func (p *PublicIpMapping) MarshalJSON() ([]byte, error) {
	type PublicIpMappingProxy PublicIpMapping
	return json.Marshal(struct {
		*PublicIpMappingProxy
		PrivateIp *import1.IPAddress `json:"privateIp,omitempty"`
		PublicIp  *import1.IPAddress `json:"publicIp,omitempty"`
	}{
		PublicIpMappingProxy: (*PublicIpMappingProxy)(p),
		PrivateIp:            p.PrivateIp,
		PublicIp:             p.PublicIp,
	})
}

func NewPublicIpMapping() *PublicIpMapping {
	p := new(PublicIpMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.PublicIpMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Quality of Service configuration for the VPN IPSec tunnel
*/
type QosConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Egress traffic limit (Mbps)
	*/
	EgressLimitMbps *int64 `json:"egressLimitMbps,omitempty"`
	/*
	  Ingress traffic limit (Mbps)
	*/
	IngressLimitMbps *int64 `json:"ingressLimitMbps,omitempty"`
}

func NewQosConfig() *QosConfig {
	p := new(QosConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.QosConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schema to configure Querier for Virtual Switch
*/
type QuerierSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Enable IGMP querier on this Virtual Switch
	*/
	IsQuerierEnabled *bool `json:"isQuerierEnabled,omitempty"`
	/*
	  VLAN Id list on which IGMP queries must be sent
	*/
	VlanIdList []int `json:"vlanIdList,omitempty"`
}

func NewQuerierSpec() *QuerierSpec {
	p := new(QuerierSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.QuerierSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsQuerierEnabled = new(bool)
	*p.IsQuerierEnabled = false

	return p
}

/*
Remote BGP gateway info needed for flow gateway scale out model.
*/
type RemoteBgpGateway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  ASN number of the BGP gateway.
	*/
	Asn *int `json:"asn"`

	IpAddress *import1.IPAddress `json:"ipAddress"`
}

func (p *RemoteBgpGateway) MarshalJSON() ([]byte, error) {
	type RemoteBgpGatewayProxy RemoteBgpGateway
	return json.Marshal(struct {
		*RemoteBgpGatewayProxy
		Asn       *int               `json:"asn,omitempty"`
		IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
	}{
		RemoteBgpGatewayProxy: (*RemoteBgpGatewayProxy)(p),
		Asn:                   p.Asn,
		IpAddress:             p.IpAddress,
	})
}

func NewRemoteBgpGateway() *RemoteBgpGateway {
	p := new(RemoteBgpGateway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteBgpGateway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
BGP service hosted on this remote gateway.
*/
type RemoteBgpService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddress `json:"address"`
	/*
	  Autonomous system number. 0 and 4294967295 are reserved.
	*/
	Asn *int64 `json:"asn"`
}

func (p *RemoteBgpService) MarshalJSON() ([]byte, error) {
	type RemoteBgpServiceProxy RemoteBgpService
	return json.Marshal(struct {
		*RemoteBgpServiceProxy
		Address *import1.IPAddress `json:"address,omitempty"`
		Asn     *int64             `json:"asn,omitempty"`
	}{
		RemoteBgpServiceProxy: (*RemoteBgpServiceProxy)(p),
		Address:               p.Address,
		Asn:                   p.Asn,
	})
}

func NewRemoteBgpService() *RemoteBgpService {
	p := new(RemoteBgpService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteBgpService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Services of this remote gateway.
*/
type RemoteNetworkServices struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RemoteBgpService *RemoteBgpService `json:"remoteBgpService,omitempty"`

	RemoteVpnService *RemoteVpnService `json:"remoteVpnService,omitempty"`

	RemoteVtepService *RemoteVtepService `json:"remoteVtepService,omitempty"`
}

func NewRemoteNetworkServices() *RemoteNetworkServices {
	p := new(RemoteNetworkServices)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteNetworkServices"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information about a subnet from the specified Prism Central cluster.
*/
type RemoteSubnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the bridge on the host for the subnet.
	*/
	BridgeName *string `json:"bridgeName,omitempty"`
	/*
	  Cluster Name
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  UUID of the cluster this subnet belongs to.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  Description of the subnet.
	*/
	Description *string `json:"description,omitempty"`

	DhcpOptions *DhcpOptions `json:"dhcpOptions,omitempty"`
	/*
	  List of IPs, which are a subset from the reserved IP address list, that must be advertised to the SDN gateway.
	*/
	DynamicIpAddresses []import1.IPAddress `json:"dynamicIpAddresses,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Hypervisor Type
	*/
	HypervisorType *string `json:"hypervisorType,omitempty"`
	/*
	  IP configuration for the subnet.
	*/
	IpConfig []IPConfig `json:"ipConfig,omitempty"`
	/*
	  IP Prefix in CIDR format.
	*/
	IpPrefix *string `json:"ipPrefix,omitempty"`

	IpUsage *IPUsage `json:"ipUsage,omitempty"`
	/*
	  Indicates whether the subnet is used for advanced networking.
	*/
	IsAdvancedNetworking *bool `json:"isAdvancedNetworking,omitempty"`
	/*
	  Indicates whether the subnet is used for external connectivity.
	*/
	IsExternal *bool `json:"isExternal,omitempty"`
	/*
	  Indicates whether NAT must be enabled for VPCs attached to the subnet. This is supported only for external subnets. NAT is enabled by default on external subnets.
	*/
	IsNatEnabled *bool `json:"isNatEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`

	MigrationState *MigrationState `json:"migrationState,omitempty"`
	/*
	  Name of the subnet.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  UUID of the Network function chain entity that this subnet belongs to (type VLAN only).
	*/
	NetworkFunctionChainReference *string `json:"networkFunctionChainReference,omitempty"`
	/*
	  For VLAN subnet, this field represents VLAN Id, valid range is from 0 to 4095; For overlay subnet, this field represents 24-bit VNI, this field is read-only.
	*/
	NetworkId *int `json:"networkId,omitempty"`
	/*
	  List of IPs that are excluded while allocating IP addresses to VM ports.
	*/
	ReservedIpAddresses []import1.IPAddress `json:"reservedIpAddresses,omitempty"`

	SubnetType *SubnetType `json:"subnetType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VirtualSwitch *VirtualSwitch `json:"virtualSwitch,omitempty"`
	/*
	  UUID of the virtual switch this subnet belongs to (type VLAN only).
	*/
	VirtualSwitchReference *string `json:"virtualSwitchReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`
	/*
	  Name of the VPC associated with the subnet.
	*/
	VpcName *string `json:"vpcName,omitempty"`
	/*
	  UUID of Virtual Private Cloud this subnet belongs to (type Overlay only).
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewRemoteSubnet() *RemoteSubnet {
	p := new(RemoteSubnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteSubnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information about a VPN connection from the specified Prism Central cluster.
*/
type RemoteVpnConnection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  IP prefixes advertised to the remote gateway over BGP.
	*/
	AdvertisedPrefixes []IPSubnet `json:"advertisedPrefixes,omitempty"`
	/*
	  Name of the remote Prism Element cluster where the VPN connection's gateway is located.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  Reference to the remote Prism Element cluster where the VPN connection's gateway is located.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  VPN connection description
	*/
	Description *string `json:"description,omitempty"`

	DpdConfig *DpdConfig `json:"dpdConfig,omitempty"`
	/*
	  Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred
	*/
	DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`

	EbgpStatus *Status `json:"ebgpStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IpsecConfig *IpsecConfig `json:"ipsecConfig"`

	IpsecTunnelStatus *Status `json:"ipsecTunnelStatus,omitempty"`
	/*
	  IP prefixes learned from the remote gateway over BGP.
	*/
	LearnedPrefixes []IPSubnet `json:"learnedPrefixes,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The local VPN gateway reference
	*/
	LocalGatewayReference *string `json:"localGatewayReference"`

	LocalGatewayRole *GatewayRole `json:"localGatewayRole"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  VPN connection name
	*/
	Name *string `json:"name"`

	QosConfig *QosConfig `json:"qosConfig,omitempty"`
	/*
	  The remote VPN gateway reference
	*/
	RemoteGatewayReference *string `json:"remoteGatewayReference"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of the VPC associated with the VPN connection.
	*/
	VpcName *string `json:"vpcName,omitempty"`
	/*
	  Reference to the VPC associated with the VPN connection.
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *RemoteVpnConnection) MarshalJSON() ([]byte, error) {
	type RemoteVpnConnectionProxy RemoteVpnConnection
	return json.Marshal(struct {
		*RemoteVpnConnectionProxy
		IpsecConfig            *IpsecConfig `json:"ipsecConfig,omitempty"`
		LocalGatewayReference  *string      `json:"localGatewayReference,omitempty"`
		LocalGatewayRole       *GatewayRole `json:"localGatewayRole,omitempty"`
		Name                   *string      `json:"name,omitempty"`
		RemoteGatewayReference *string      `json:"remoteGatewayReference,omitempty"`
	}{
		RemoteVpnConnectionProxy: (*RemoteVpnConnectionProxy)(p),
		IpsecConfig:              p.IpsecConfig,
		LocalGatewayReference:    p.LocalGatewayReference,
		LocalGatewayRole:         p.LocalGatewayRole,
		Name:                     p.Name,
		RemoteGatewayReference:   p.RemoteGatewayReference,
	})
}

func NewRemoteVpnConnection() *RemoteVpnConnection {
	p := new(RemoteVpnConnection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVpnConnection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-vpn-connections/{extId} Get operation
*/
type RemoteVpnConnectionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoteVpnConnectionApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoteVpnConnectionApiResponse() *RemoteVpnConnectionApiResponse {
	p := new(RemoteVpnConnectionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVpnConnectionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoteVpnConnectionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoteVpnConnectionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoteVpnConnectionApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-vpn-connections Get operation
*/
type RemoteVpnConnectionListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoteVpnConnectionListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoteVpnConnectionListApiResponse() *RemoteVpnConnectionListApiResponse {
	p := new(RemoteVpnConnectionListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVpnConnectionListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoteVpnConnectionListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoteVpnConnectionListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoteVpnConnectionListApiResponseData()
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
VPN service hosted on this remote gateway.
*/
type RemoteVpnService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EbgpConfig *BgpConfig `json:"ebgpConfig,omitempty"`

	PeerIgpConfig *InternalRoutingConfig `json:"peerIgpConfig,omitempty"`

	ServiceAddress *import1.IPAddress `json:"serviceAddress,omitempty"`
	/*
	  Boolean flag indicating user opt-in for installing Xi LB route in on-prem Prism Central and Prism Element CVMs provided on-prem Prism Central, Prism Element and VPN VM are in the same subnet.
	*/
	ShouldInstallXiRoute *bool `json:"shouldInstallXiRoute,omitempty"`
}

func NewRemoteVpnService() *RemoteVpnService {
	p := new(RemoteVpnService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVpnService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldInstallXiRoute = new(bool)
	*p.ShouldInstallXiRoute = false

	return p
}

type RemoteVtepGateway struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the remote Prism Element cluster that owns the VTEP gateway.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  Reference to the remote Prism Element cluster that owns the VTEP gateway.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  If set to true, VTEP gateway is local. When set to false, VTEP gateway is remote.
	*/
	IsLocal *bool `json:"isLocal,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  VTEP gateway name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of the VPC associated with the VTEP gateway.
	*/
	VpcName *string `json:"vpcName,omitempty"`
	/*
	  Reference to the VPC associated with the VTEP gateway.
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
	/*
	  VXLAN port.
	*/
	VxlanPort *int `json:"vxlanPort,omitempty"`
}

func NewRemoteVtepGateway() *RemoteVtepGateway {
	p := new(RemoteVtepGateway)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVtepGateway"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-vtep-gateways/{extId} Get operation
*/
type RemoteVtepGatewayApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoteVtepGatewayApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoteVtepGatewayApiResponse() *RemoteVtepGatewayApiResponse {
	p := new(RemoteVtepGatewayApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVtepGatewayApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoteVtepGatewayApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoteVtepGatewayApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoteVtepGatewayApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/config/clusters/{clusterExtId}/remote-vtep-gateways Get operation
*/
type RemoteVtepGatewayListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoteVtepGatewayListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoteVtepGatewayListApiResponse() *RemoteVtepGatewayListApiResponse {
	p := new(RemoteVtepGatewayListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVtepGatewayListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoteVtepGatewayListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoteVtepGatewayListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoteVtepGatewayListApiResponseData()
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
VTEP service hosted on this remote gateway.
*/
type RemoteVtepService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Remote VXLAN Tunnel Endpoints configuration.
	*/
	Vteps []Vtep `json:"vteps,omitempty"`
	/*
	  VXLAN port.
	*/
	VxlanPort *int `json:"vxlanPort,omitempty"`
}

func NewRemoteVtepService() *RemoteVtepService {
	p := new(RemoteVtepService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVtepService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.VxlanPort = new(int)
	*p.VxlanPort = 4789

	return p
}

/*
Status of each VTEP. Applicable only when connectionType is VXLAN.
*/
type RemoteVtepStretchStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddress `json:"address,omitempty"`

	Status *StretchStatus `json:"status,omitempty"`
}

func NewRemoteVtepStretchStatus() *RemoteVtepStretchStatus {
	p := new(RemoteVtepStretchStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RemoteVtepStretchStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of fallback action in reroute case when service VM is down.
*/
type RerouteFallbackAction int

const (
	REROUTEFALLBACKACTION_UNKNOWN     RerouteFallbackAction = 0
	REROUTEFALLBACKACTION_REDACTED    RerouteFallbackAction = 1
	REROUTEFALLBACKACTION_ALLOW       RerouteFallbackAction = 2
	REROUTEFALLBACKACTION_DROP        RerouteFallbackAction = 3
	REROUTEFALLBACKACTION_PASSTHROUGH RerouteFallbackAction = 4
	REROUTEFALLBACKACTION_NO_ACTION   RerouteFallbackAction = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RerouteFallbackAction) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DROP",
		"PASSTHROUGH",
		"NO_ACTION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RerouteFallbackAction) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DROP",
		"PASSTHROUGH",
		"NO_ACTION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RerouteFallbackAction) index(name string) RerouteFallbackAction {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOW",
		"DROP",
		"PASSTHROUGH",
		"NO_ACTION",
	}
	for idx := range names {
		if names[idx] == name {
			return RerouteFallbackAction(idx)
		}
	}
	return REROUTEFALLBACKACTION_UNKNOWN
}

func (e *RerouteFallbackAction) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RerouteFallbackAction:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RerouteFallbackAction) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RerouteFallbackAction) Ref() *RerouteFallbackAction {
	return &e
}

/*
Parameters for the reroute action which includes the reroute service IP and the fallback action when the service IP is down.
*/
type RerouteParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EgressServiceIp *import1.IPAddress `json:"egressServiceIp,omitempty"`

	IngressServiceIp *import1.IPAddress `json:"ingressServiceIp,omitempty"`

	RerouteFallbackAction *RerouteFallbackAction `json:"rerouteFallbackAction,omitempty"`

	ServiceIp *import1.IPAddress `json:"serviceIp,omitempty"`
}

func NewRerouteParam() *RerouteParam {
	p := new(RerouteParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RerouteParam"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of IP Address reservation.
*/
type ReserveType int

const (
	RESERVETYPE_UNKNOWN          ReserveType = 0
	RESERVETYPE_REDACTED         ReserveType = 1
	RESERVETYPE_IP_ADDRESS_COUNT ReserveType = 2
	RESERVETYPE_IP_ADDRESS_RANGE ReserveType = 3
	RESERVETYPE_IP_ADDRESS_LIST  ReserveType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReserveType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_COUNT",
		"IP_ADDRESS_RANGE",
		"IP_ADDRESS_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReserveType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_COUNT",
		"IP_ADDRESS_RANGE",
		"IP_ADDRESS_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReserveType) index(name string) ReserveType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_COUNT",
		"IP_ADDRESS_RANGE",
		"IP_ADDRESS_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return ReserveType(idx)
		}
	}
	return RESERVETYPE_UNKNOWN
}

func (e *ReserveType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReserveType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReserveType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReserveType) Ref() *ReserveType {
	return &e
}

/*
Information pertaining to a reserved IP address on a subnet.
*/
type ReservedIp struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Optional context to associate with a reserved IP address.
	*/
	ClientContext *string `json:"clientContext,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Reserved IPv4 address.
	*/
	Ipv4Address *string `json:"ipv4Address,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewReservedIp() *ReservedIp {
	p := new(ReservedIp)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.ReservedIp"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Route of the VPC route table.
*/
type Route struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  BGP session description.
	*/
	Description *string `json:"description,omitempty"`

	Destination *IPSubnet `json:"destination,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  External routing domain to which this route belongs.
	*/
	ExternalRoutingDomainReference *string `json:"externalRoutingDomainReference,omitempty"`
	/*
	  Indicates whether the route is active in the forwarding plane.
	*/
	IsActive *bool `json:"isActive,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Route name.
	*/
	Name *string `json:"name,omitempty"`

	Nexthop *Nexthop `json:"nexthop,omitempty"`
	/*
	  Route priority. A higher value implies greater preference is assigned to the route.
	*/
	Priority *int `json:"priority,omitempty"`
	/*
	  Route table to which this route belongs.
	*/
	RouteTableReference *string `json:"routeTableReference,omitempty"`

	RouteType *RouteType `json:"routeType"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPC to which this route belongs.
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *Route) MarshalJSON() ([]byte, error) {
	type RouteProxy Route
	return json.Marshal(struct {
		*RouteProxy
		RouteType *RouteType `json:"routeType,omitempty"`
	}{
		RouteProxy: (*RouteProxy)(p),
		RouteType:  p.RouteType,
	})
}

func NewRoute() *Route {
	p := new(Route)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Route"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Route table
*/
type RouteTable struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  External routing domain associated with this route table
	*/
	ExternalRoutingDomainReference *string `json:"externalRoutingDomainReference,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPC
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewRouteTable() *RouteTable {
	p := new(RouteTable)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RouteTable"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the VPC route.
*/
type RouteType int

const (
	ROUTETYPE_UNKNOWN  RouteType = 0
	ROUTETYPE_REDACTED RouteType = 1
	ROUTETYPE_DYNAMIC  RouteType = 2
	ROUTETYPE_LOCAL    RouteType = 3
	ROUTETYPE_STATIC   RouteType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RouteType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"LOCAL",
		"STATIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RouteType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"LOCAL",
		"STATIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RouteType) index(name string) RouteType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DYNAMIC",
		"LOCAL",
		"STATIC",
	}
	for idx := range names {
		if names[idx] == name {
			return RouteType(idx)
		}
	}
	return ROUTETYPE_UNKNOWN
}

func (e *RouteType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RouteType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RouteType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RouteType) Ref() *RouteType {
	return &e
}

/*
Schema to configure a routing policy.
*/
type RoutingPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A description of the routing policy.
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

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the routing policy.
	*/
	Name *string `json:"name"`

	Policies []RoutingPolicyRule `json:"policies"`
	/*
	  Priority of the routing policy.
	*/
	Priority *int `json:"priority"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vpc *VpcName `json:"vpc,omitempty"`
	/*
	  ExtId of the VPC extId to which the routing policy belongs.
	*/
	VpcExtId *string `json:"vpcExtId"`
}

func (p *RoutingPolicy) MarshalJSON() ([]byte, error) {
	type RoutingPolicyProxy RoutingPolicy
	return json.Marshal(struct {
		*RoutingPolicyProxy
		Name     *string             `json:"name,omitempty"`
		Policies []RoutingPolicyRule `json:"policies,omitempty"`
		Priority *int                `json:"priority,omitempty"`
		VpcExtId *string             `json:"vpcExtId,omitempty"`
	}{
		RoutingPolicyProxy: (*RoutingPolicyProxy)(p),
		Name:               p.Name,
		Policies:           p.Policies,
		Priority:           p.Priority,
		VpcExtId:           p.VpcExtId,
	})
}

func NewRoutingPolicy() *RoutingPolicy {
	p := new(RoutingPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RoutingPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The action to be taken on the traffic matching the routing policy.
*/
type RoutingPolicyAction struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *RoutingPolicyActionType `json:"actionType"`

	NexthopIpAddress *import1.IPAddress `json:"nexthopIpAddress,omitempty"`

	RerouteParams []RerouteParam `json:"rerouteParams,omitempty"`
}

func (p *RoutingPolicyAction) MarshalJSON() ([]byte, error) {
	type RoutingPolicyActionProxy RoutingPolicyAction
	return json.Marshal(struct {
		*RoutingPolicyActionProxy
		ActionType *RoutingPolicyActionType `json:"actionType,omitempty"`
	}{
		RoutingPolicyActionProxy: (*RoutingPolicyActionProxy)(p),
		ActionType:               p.ActionType,
	})
}

func NewRoutingPolicyAction() *RoutingPolicyAction {
	p := new(RoutingPolicyAction)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RoutingPolicyAction"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Routing policy action type.
*/
type RoutingPolicyActionType int

const (
	ROUTINGPOLICYACTIONTYPE_UNKNOWN  RoutingPolicyActionType = 0
	ROUTINGPOLICYACTIONTYPE_REDACTED RoutingPolicyActionType = 1
	ROUTINGPOLICYACTIONTYPE_PERMIT   RoutingPolicyActionType = 2
	ROUTINGPOLICYACTIONTYPE_DENY     RoutingPolicyActionType = 3
	ROUTINGPOLICYACTIONTYPE_REROUTE  RoutingPolicyActionType = 4
	ROUTINGPOLICYACTIONTYPE_FORWARD  RoutingPolicyActionType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RoutingPolicyActionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PERMIT",
		"DENY",
		"REROUTE",
		"FORWARD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RoutingPolicyActionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PERMIT",
		"DENY",
		"REROUTE",
		"FORWARD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RoutingPolicyActionType) index(name string) RoutingPolicyActionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PERMIT",
		"DENY",
		"REROUTE",
		"FORWARD",
	}
	for idx := range names {
		if names[idx] == name {
			return RoutingPolicyActionType(idx)
		}
	}
	return ROUTINGPOLICYACTIONTYPE_UNKNOWN
}

func (e *RoutingPolicyActionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RoutingPolicyActionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RoutingPolicyActionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RoutingPolicyActionType) Ref() *RoutingPolicyActionType {
	return &e
}

/*
Match condition for the traffic that is entering the VPC.
*/
type RoutingPolicyMatchCondition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Destination *AddressTypeObject `json:"destination"`
	/*

	 */
	ProtocolParametersItemDiscriminator_ *string `json:"$protocolParametersItemDiscriminator,omitempty"`
	/*
	  Protocol parameters of the traffic that is exiting/leaving the VPC.
	*/
	ProtocolParameters *OneOfRoutingPolicyMatchConditionProtocolParameters `json:"protocolParameters,omitempty"`

	ProtocolType *ProtocolType `json:"protocolType"`

	Source *AddressTypeObject `json:"source"`
}

func (p *RoutingPolicyMatchCondition) MarshalJSON() ([]byte, error) {
	type RoutingPolicyMatchConditionProxy RoutingPolicyMatchCondition
	return json.Marshal(struct {
		*RoutingPolicyMatchConditionProxy
		Destination  *AddressTypeObject `json:"destination,omitempty"`
		ProtocolType *ProtocolType      `json:"protocolType,omitempty"`
		Source       *AddressTypeObject `json:"source,omitempty"`
	}{
		RoutingPolicyMatchConditionProxy: (*RoutingPolicyMatchConditionProxy)(p),
		Destination:                      p.Destination,
		ProtocolType:                     p.ProtocolType,
		Source:                           p.Source,
	})
}

func NewRoutingPolicyMatchCondition() *RoutingPolicyMatchCondition {
	p := new(RoutingPolicyMatchCondition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RoutingPolicyMatchCondition"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RoutingPolicyMatchCondition) GetProtocolParameters() interface{} {
	if nil == p.ProtocolParameters {
		return nil
	}
	return p.ProtocolParameters.GetValue()
}

func (p *RoutingPolicyMatchCondition) SetProtocolParameters(v interface{}) error {
	if nil == p.ProtocolParameters {
		p.ProtocolParameters = NewOneOfRoutingPolicyMatchConditionProtocolParameters()
	}
	e := p.ProtocolParameters.SetValue(v)
	if nil == e {
		if nil == p.ProtocolParametersItemDiscriminator_ {
			p.ProtocolParametersItemDiscriminator_ = new(string)
		}
		*p.ProtocolParametersItemDiscriminator_ = *p.ProtocolParameters.Discriminator
	}
	return e
}

type RoutingPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A description of the routing policy.
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

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the routing policy.
	*/
	Name *string `json:"name"`

	Policies []RoutingPolicyRule `json:"policies"`
	/*
	  Priority of the routing policy.
	*/
	Priority *int `json:"priority"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Vpc *VpcName `json:"vpc,omitempty"`
	/*
	  ExtId of the VPC extId to which the routing policy belongs.
	*/
	VpcExtId *string `json:"vpcExtId"`

	VpcNameProjection *VpcNameProjection `json:"vpcNameProjection,omitempty"`
}

func (p *RoutingPolicyProjection) MarshalJSON() ([]byte, error) {
	type RoutingPolicyProjectionProxy RoutingPolicyProjection
	return json.Marshal(struct {
		*RoutingPolicyProjectionProxy
		Name     *string             `json:"name,omitempty"`
		Policies []RoutingPolicyRule `json:"policies,omitempty"`
		Priority *int                `json:"priority,omitempty"`
		VpcExtId *string             `json:"vpcExtId,omitempty"`
	}{
		RoutingPolicyProjectionProxy: (*RoutingPolicyProjectionProxy)(p),
		Name:                         p.Name,
		Policies:                     p.Policies,
		Priority:                     p.Priority,
		VpcExtId:                     p.VpcExtId,
	})
}

func NewRoutingPolicyProjection() *RoutingPolicyProjection {
	p := new(RoutingPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RoutingPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Policy indicating the match rule and the action.
*/
type RoutingPolicyRule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  If True, policies in the reverse direction will be installed with the same action but source and destination will be swapped.
	*/
	IsBidirectional *bool `json:"isBidirectional,omitempty"`

	PolicyAction *RoutingPolicyAction `json:"policyAction"`

	PolicyMatch *RoutingPolicyMatchCondition `json:"policyMatch"`
}

func (p *RoutingPolicyRule) MarshalJSON() ([]byte, error) {
	type RoutingPolicyRuleProxy RoutingPolicyRule
	return json.Marshal(struct {
		*RoutingPolicyRuleProxy
		PolicyAction *RoutingPolicyAction         `json:"policyAction,omitempty"`
		PolicyMatch  *RoutingPolicyMatchCondition `json:"policyMatch,omitempty"`
	}{
		RoutingPolicyRuleProxy: (*RoutingPolicyRuleProxy)(p),
		PolicyAction:           p.PolicyAction,
		PolicyMatch:            p.PolicyMatch,
	})
}

func NewRoutingPolicyRule() *RoutingPolicyRule {
	p := new(RoutingPolicyRule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.RoutingPolicyRule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsBidirectional = new(bool)
	*p.IsBidirectional = false

	return p
}

/*
Scope type:the permissible values are PC or PE.
*/
type ScopeType int

const (
	SCOPETYPE_UNKNOWN  ScopeType = 0
	SCOPETYPE_REDACTED ScopeType = 1
	SCOPETYPE_PC       ScopeType = 2
	SCOPETYPE_PE       ScopeType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScopeType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScopeType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScopeType) index(name string) ScopeType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
	}
	for idx := range names {
		if names[idx] == name {
			return ScopeType(idx)
		}
	}
	return SCOPETYPE_UNKNOWN
}

func (e *ScopeType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScopeType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScopeType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScopeType) Ref() *ScopeType {
	return &e
}

/*
Site-specific stretch configuration parameters.
*/
type SiteParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VPN connection or network gateway with VTEP service used for this Layer2 stretch.
	*/
	ConnectionReference *string `json:"connectionReference,omitempty"`

	DefaultGatewayIPAddress *import1.IPAddress `json:"defaultGatewayIPAddress,omitempty"`
	/*
	  Prism Central cluster reference.
	*/
	PcClusterReference *string `json:"pcClusterReference,omitempty"`

	StretchInterfaceIpAddress *import1.IPAddress `json:"stretchInterfaceIpAddress,omitempty"`
	/*
	  Subnet reference.
	*/
	StretchSubnetReference *string `json:"stretchSubnetReference,omitempty"`

	VpnInterfaceIPAddress *import1.IPAddress `json:"vpnInterfaceIPAddress,omitempty"`
}

func NewSiteParams() *SiteParams {
	p := new(SiteParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.SiteParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Up/Down status of component.
*/
type State int

const (
	STATE_UNKNOWN  State = 0
	STATE_REDACTED State = 1
	STATE_UP       State = 2
	STATE_DOWN     State = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *State) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e State) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *State) index(name string) State {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UP",
		"DOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return State(idx)
		}
	}
	return STATE_UNKNOWN
}

func (e *State) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for State:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *State) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e State) Ref() *State {
	return &e
}

/*
Up/Down status of component and message.
*/
type Status struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Detailed message of component status.
	*/
	Message *string `json:"message,omitempty"`

	State *State `json:"state,omitempty"`
}

func NewStatus() *Status {
	p := new(Status)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Status"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the connection used for stretching the subnet. The default is VPN.
*/
type StretchConnectionType int

const (
	STRETCHCONNECTIONTYPE_UNKNOWN  StretchConnectionType = 0
	STRETCHCONNECTIONTYPE_REDACTED StretchConnectionType = 1
	STRETCHCONNECTIONTYPE_VPN      StretchConnectionType = 2
	STRETCHCONNECTIONTYPE_VXLAN    StretchConnectionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *StretchConnectionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VPN",
		"VXLAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e StretchConnectionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VPN",
		"VXLAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *StretchConnectionType) index(name string) StretchConnectionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VPN",
		"VXLAN",
	}
	for idx := range names {
		if names[idx] == name {
			return StretchConnectionType(idx)
		}
	}
	return STRETCHCONNECTIONTYPE_UNKNOWN
}

func (e *StretchConnectionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StretchConnectionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StretchConnectionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StretchConnectionType) Ref() *StretchConnectionType {
	return &e
}

/*
Current status of the Layer2 extension among subnets.
*/
type StretchStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Detailed text describing the runtime status of this stretch configuration.
	*/
	Detail *string `json:"detail,omitempty"`

	InterfaceState *State `json:"interfaceState,omitempty"`
	/*
	  The round-trip time, in milliseconds, between subnets in this stretch configuration.
	*/
	RoundTripTimeMillis *float32 `json:"roundTripTimeMillis,omitempty"`

	TunnelState *State `json:"tunnelState,omitempty"`
}

func NewStretchStatus() *StretchStatus {
	p := new(StretchStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.StretchStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
API schema for subnet.
*/
type Subnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the bridge on the host for the subnet.
	*/
	BridgeName *string `json:"bridgeName,omitempty"`
	/*
	  Cluster Name
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  UUID of the cluster this subnet belongs to.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  Description of the subnet.
	*/
	Description *string `json:"description,omitempty"`

	DhcpOptions *DhcpOptions `json:"dhcpOptions,omitempty"`
	/*
	  List of IPs, which are a subset from the reserved IP address list, that must be advertised to the SDN gateway.
	*/
	DynamicIpAddresses []import1.IPAddress `json:"dynamicIpAddresses,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Hypervisor Type
	*/
	HypervisorType *string `json:"hypervisorType,omitempty"`
	/*
	  IP configuration for the subnet.
	*/
	IpConfig []IPConfig `json:"ipConfig,omitempty"`
	/*
	  IP Prefix in CIDR format.
	*/
	IpPrefix *string `json:"ipPrefix,omitempty"`

	IpUsage *IPUsage `json:"ipUsage,omitempty"`
	/*
	  Indicates whether the subnet is used for advanced networking.
	*/
	IsAdvancedNetworking *bool `json:"isAdvancedNetworking,omitempty"`
	/*
	  Indicates whether the subnet is used for external connectivity.
	*/
	IsExternal *bool `json:"isExternal,omitempty"`
	/*
	  Indicates whether NAT must be enabled for VPCs attached to the subnet. This is supported only for external subnets. NAT is enabled by default on external subnets.
	*/
	IsNatEnabled *bool `json:"isNatEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`

	MigrationState *MigrationState `json:"migrationState,omitempty"`
	/*
	  Name of the subnet.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  UUID of the Network function chain entity that this subnet belongs to (type VLAN only).
	*/
	NetworkFunctionChainReference *string `json:"networkFunctionChainReference,omitempty"`
	/*
	  For VLAN subnet, this field represents VLAN Id, valid range is from 0 to 4095; For overlay subnet, this field represents 24-bit VNI, this field is read-only.
	*/
	NetworkId *int `json:"networkId,omitempty"`
	/*
	  List of IPs that are excluded while allocating IP addresses to VM ports.
	*/
	ReservedIpAddresses []import1.IPAddress `json:"reservedIpAddresses,omitempty"`

	SubnetType *SubnetType `json:"subnetType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VirtualSwitch *VirtualSwitch `json:"virtualSwitch,omitempty"`
	/*
	  UUID of the virtual switch this subnet belongs to (type VLAN only).
	*/
	VirtualSwitchReference *string `json:"virtualSwitchReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`
	/*
	  UUID of Virtual Private Cloud this subnet belongs to (type Overlay only).
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewSubnet() *Subnet {
	p := new(Subnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Subnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SubnetInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the subnet to be migrated.
	*/
	SubnetUuid *string `json:"subnetUuid"`
}

func (p *SubnetInfo) MarshalJSON() ([]byte, error) {
	type SubnetInfoProxy SubnetInfo
	return json.Marshal(struct {
		*SubnetInfoProxy
		SubnetUuid *string `json:"subnetUuid,omitempty"`
	}{
		SubnetInfoProxy: (*SubnetInfoProxy)(p),
		SubnetUuid:      p.SubnetUuid,
	})
}

func NewSubnetInfo() *SubnetInfo {
	p := new(SubnetInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.SubnetInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SubnetProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the bridge on the host for the subnet.
	*/
	BridgeName *string `json:"bridgeName,omitempty"`
	/*
	  Cluster Name
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  UUID of the cluster this subnet belongs to.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  Description of the subnet.
	*/
	Description *string `json:"description,omitempty"`

	DhcpOptions *DhcpOptions `json:"dhcpOptions,omitempty"`
	/*
	  List of IPs, which are a subset from the reserved IP address list, that must be advertised to the SDN gateway.
	*/
	DynamicIpAddresses []import1.IPAddress `json:"dynamicIpAddresses,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Hypervisor Type
	*/
	HypervisorType *string `json:"hypervisorType,omitempty"`
	/*
	  IP configuration for the subnet.
	*/
	IpConfig []IPConfig `json:"ipConfig,omitempty"`
	/*
	  IP Prefix in CIDR format.
	*/
	IpPrefix *string `json:"ipPrefix,omitempty"`

	IpUsage *IPUsage `json:"ipUsage,omitempty"`
	/*
	  Indicates whether the subnet is used for advanced networking.
	*/
	IsAdvancedNetworking *bool `json:"isAdvancedNetworking,omitempty"`
	/*
	  Indicates whether the subnet is used for external connectivity.
	*/
	IsExternal *bool `json:"isExternal,omitempty"`
	/*
	  Indicates whether NAT must be enabled for VPCs attached to the subnet. This is supported only for external subnets. NAT is enabled by default on external subnets.
	*/
	IsNatEnabled *bool `json:"isNatEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`

	MigrationState *MigrationState `json:"migrationState,omitempty"`
	/*
	  Name of the subnet.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  UUID of the Network function chain entity that this subnet belongs to (type VLAN only).
	*/
	NetworkFunctionChainReference *string `json:"networkFunctionChainReference,omitempty"`
	/*
	  For VLAN subnet, this field represents VLAN Id, valid range is from 0 to 4095; For overlay subnet, this field represents 24-bit VNI, this field is read-only.
	*/
	NetworkId *int `json:"networkId,omitempty"`
	/*
	  List of IPs that are excluded while allocating IP addresses to VM ports.
	*/
	ReservedIpAddresses []import1.IPAddress `json:"reservedIpAddresses,omitempty"`

	SubnetType *SubnetType `json:"subnetType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VirtualSwitch *VirtualSwitch `json:"virtualSwitch,omitempty"`

	VirtualSwitchProjection *VirtualSwitchProjection `json:"virtualSwitchProjection,omitempty"`
	/*
	  UUID of the virtual switch this subnet belongs to (type VLAN only).
	*/
	VirtualSwitchReference *string `json:"virtualSwitchReference,omitempty"`

	Vpc *Vpc `json:"vpc,omitempty"`

	VpcProjection *VpcProjection `json:"vpcProjection,omitempty"`
	/*
	  UUID of Virtual Private Cloud this subnet belongs to (type Overlay only).
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func NewSubnetProjection() *SubnetProjection {
	p := new(SubnetProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.SubnetProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of subnet.
*/
type SubnetType int

const (
	SUBNETTYPE_UNKNOWN  SubnetType = 0
	SUBNETTYPE_REDACTED SubnetType = 1
	SUBNETTYPE_OVERLAY  SubnetType = 2
	SUBNETTYPE_VLAN     SubnetType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SubnetType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OVERLAY",
		"VLAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SubnetType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OVERLAY",
		"VLAN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SubnetType) index(name string) SubnetType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OVERLAY",
		"VLAN",
	}
	for idx := range names {
		if names[idx] == name {
			return SubnetType(idx)
		}
	}
	return SUBNETTYPE_UNKNOWN
}

func (e *SubnetType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SubnetType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SubnetType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SubnetType) Ref() *SubnetType {
	return &e
}

/*
Target configuration for the load balancer session.
*/
type Target struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NicTargets []NicTarget `json:"nicTargets,omitempty"`
}

func NewTarget() *Target {
	p := new(Target)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Target"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Health status of the load balancer session target.
*/
type TargetHealth int

const (
	TARGETHEALTH_UNKNOWN   TargetHealth = 0
	TARGETHEALTH_REDACTED  TargetHealth = 1
	TARGETHEALTH_HEALTHY   TargetHealth = 2
	TARGETHEALTH_UNHEALTHY TargetHealth = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TargetHealth) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"UNHEALTHY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TargetHealth) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"UNHEALTHY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TargetHealth) index(name string) TargetHealth {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HEALTHY",
		"UNHEALTHY",
	}
	for idx := range names {
		if names[idx] == name {
			return TargetHealth(idx)
		}
	}
	return TARGETHEALTH_UNKNOWN
}

func (e *TargetHealth) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TargetHealth:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TargetHealth) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TargetHealth) Ref() *TargetHealth {
	return &e
}

/*
REST response for all response codes in API path /networking/v4.0/config/vpn-connections/{extId} Put operation
*/
type TaskReferenceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTaskReferenceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskReferenceApiResponse() *TaskReferenceApiResponse {
	p := new(TaskReferenceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.TaskReferenceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TaskReferenceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TaskReferenceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTaskReferenceApiResponseData()
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
Mirror network traffic from a set of source ports, to a set of destination ports.
*/
type TrafficMirror struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of cluster UUIDs that are configured for this session. Currently, only 1 cluster is allowed to participate in a session.
	*/
	ClusterReferenceList []string `json:"clusterReferenceList,omitempty"`
	/*
	  Description of the session.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  List of destination ports of the session. Maximum of 2 destination ports are allowed per session. Each session should have at least 1 destination port.
	*/
	DestinationList []TrafficMirrorPort `json:"destinationList,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of host UUIDs that are configured for this session. Currently, only 1 host is allowed to participate in a session.
	*/
	HostReferenceList []string `json:"hostReferenceList,omitempty"`
	/*
	  Indicates whether the port mirroring session is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the session.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  List of source ports of the session. Maximum of 4 source ports are allowed per session. Each session should have at least 1 source port.
	*/
	SourceList []TrafficMirrorSourcePort `json:"sourceList,omitempty"`

	State *TrafficMirrorState `json:"state,omitempty"`
	/*
	  Traffic mirror state message.
	*/
	StateMessage *string `json:"stateMessage,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Traffic mirror virtual switch reference to use for Remote SPAN.
	*/
	VirtualSwitchReference *string `json:"virtualSwitchReference,omitempty"`
}

func NewTrafficMirror() *TrafficMirror {
	p := new(TrafficMirror)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.TrafficMirror"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

/*
Traffic mirror session port description.
*/
type TrafficMirrorPort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the port is up.
	*/
	IsUp *bool `json:"isUp,omitempty"`

	NicType *TrafficMirrorPortNicType `json:"nicType"`
	/*
	  UUID of the NIC that this port belongs to.
	*/
	NicUuid *string `json:"nicUuid"`
}

func (p *TrafficMirrorPort) MarshalJSON() ([]byte, error) {
	type TrafficMirrorPortProxy TrafficMirrorPort
	return json.Marshal(struct {
		*TrafficMirrorPortProxy
		NicType *TrafficMirrorPortNicType `json:"nicType,omitempty"`
		NicUuid *string                   `json:"nicUuid,omitempty"`
	}{
		TrafficMirrorPortProxy: (*TrafficMirrorPortProxy)(p),
		NicType:                p.NicType,
		NicUuid:                p.NicUuid,
	})
}

func NewTrafficMirrorPort() *TrafficMirrorPort {
	p := new(TrafficMirrorPort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.TrafficMirrorPort"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Port NIC type for the Traffic mirror session. Allowed values are HOST_NIC and VIRTUAL_NIC.
*/
type TrafficMirrorPortNicType int

const (
	TRAFFICMIRRORPORTNICTYPE_UNKNOWN     TrafficMirrorPortNicType = 0
	TRAFFICMIRRORPORTNICTYPE_REDACTED    TrafficMirrorPortNicType = 1
	TRAFFICMIRRORPORTNICTYPE_HOST_NIC    TrafficMirrorPortNicType = 2
	TRAFFICMIRRORPORTNICTYPE_VIRTUAL_NIC TrafficMirrorPortNicType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TrafficMirrorPortNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOST_NIC",
		"VIRTUAL_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TrafficMirrorPortNicType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOST_NIC",
		"VIRTUAL_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TrafficMirrorPortNicType) index(name string) TrafficMirrorPortNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOST_NIC",
		"VIRTUAL_NIC",
	}
	for idx := range names {
		if names[idx] == name {
			return TrafficMirrorPortNicType(idx)
		}
	}
	return TRAFFICMIRRORPORTNICTYPE_UNKNOWN
}

func (e *TrafficMirrorPortNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TrafficMirrorPortNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TrafficMirrorPortNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TrafficMirrorPortNicType) Ref() *TrafficMirrorPortNicType {
	return &e
}

/*
Traffic mirror session source port to mirror traffic from.
*/
type TrafficMirrorSourcePort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Direction *TrafficMirrorSourcePortDirection `json:"direction"`
	/*
	  Indicates whether the port is up.
	*/
	IsUp *bool `json:"isUp,omitempty"`

	NicType *TrafficMirrorPortNicType `json:"nicType"`
	/*
	  UUID of the NIC that this port belongs to.
	*/
	NicUuid *string `json:"nicUuid"`
}

func (p *TrafficMirrorSourcePort) MarshalJSON() ([]byte, error) {
	type TrafficMirrorSourcePortProxy TrafficMirrorSourcePort
	return json.Marshal(struct {
		*TrafficMirrorSourcePortProxy
		Direction *TrafficMirrorSourcePortDirection `json:"direction,omitempty"`
		NicType   *TrafficMirrorPortNicType         `json:"nicType,omitempty"`
		NicUuid   *string                           `json:"nicUuid,omitempty"`
	}{
		TrafficMirrorSourcePortProxy: (*TrafficMirrorSourcePortProxy)(p),
		Direction:                    p.Direction,
		NicType:                      p.NicType,
		NicUuid:                      p.NicUuid,
	})
}

func NewTrafficMirrorSourcePort() *TrafficMirrorSourcePort {
	p := new(TrafficMirrorSourcePort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.TrafficMirrorSourcePort"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the direction of traffic that the session will mirror. Allowed values are INGRESS, EGRESS and BIDIRECTIONAL.
*/
type TrafficMirrorSourcePortDirection int

const (
	TRAFFICMIRRORSOURCEPORTDIRECTION_UNKNOWN       TrafficMirrorSourcePortDirection = 0
	TRAFFICMIRRORSOURCEPORTDIRECTION_REDACTED      TrafficMirrorSourcePortDirection = 1
	TRAFFICMIRRORSOURCEPORTDIRECTION_INGRESS       TrafficMirrorSourcePortDirection = 2
	TRAFFICMIRRORSOURCEPORTDIRECTION_EGRESS        TrafficMirrorSourcePortDirection = 3
	TRAFFICMIRRORSOURCEPORTDIRECTION_BIDIRECTIONAL TrafficMirrorSourcePortDirection = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TrafficMirrorSourcePortDirection) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"BIDIRECTIONAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TrafficMirrorSourcePortDirection) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"BIDIRECTIONAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TrafficMirrorSourcePortDirection) index(name string) TrafficMirrorSourcePortDirection {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"BIDIRECTIONAL",
	}
	for idx := range names {
		if names[idx] == name {
			return TrafficMirrorSourcePortDirection(idx)
		}
	}
	return TRAFFICMIRRORSOURCEPORTDIRECTION_UNKNOWN
}

func (e *TrafficMirrorSourcePortDirection) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TrafficMirrorSourcePortDirection:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TrafficMirrorSourcePortDirection) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TrafficMirrorSourcePortDirection) Ref() *TrafficMirrorSourcePortDirection {
	return &e
}

/*
Traffic mirror state value.
*/
type TrafficMirrorState int

const (
	TRAFFICMIRRORSTATE_UNKNOWN  TrafficMirrorState = 0
	TRAFFICMIRRORSTATE_REDACTED TrafficMirrorState = 1
	TRAFFICMIRRORSTATE_ACTIVE   TrafficMirrorState = 2
	TRAFFICMIRRORSTATE_ERROR    TrafficMirrorState = 3
	TRAFFICMIRRORSTATE_DISABLED TrafficMirrorState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TrafficMirrorState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TrafficMirrorState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TrafficMirrorState) index(name string) TrafficMirrorState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return TrafficMirrorState(idx)
		}
	}
	return TRAFFICMIRRORSTATE_UNKNOWN
}

func (e *TrafficMirrorState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TrafficMirrorState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TrafficMirrorState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TrafficMirrorState) Ref() *TrafficMirrorState {
	return &e
}

/*
Type of IP Address unreservation.
*/
type UnreserveType int

const (
	UNRESERVETYPE_UNKNOWN          UnreserveType = 0
	UNRESERVETYPE_REDACTED         UnreserveType = 1
	UNRESERVETYPE_IP_ADDRESS_LIST  UnreserveType = 2
	UNRESERVETYPE_IP_ADDRESS_RANGE UnreserveType = 3
	UNRESERVETYPE_CONTEXT          UnreserveType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UnreserveType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_LIST",
		"IP_ADDRESS_RANGE",
		"CONTEXT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UnreserveType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_LIST",
		"IP_ADDRESS_RANGE",
		"CONTEXT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UnreserveType) index(name string) UnreserveType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IP_ADDRESS_LIST",
		"IP_ADDRESS_RANGE",
		"CONTEXT",
	}
	for idx := range names {
		if names[idx] == name {
			return UnreserveType(idx)
		}
	}
	return UNRESERVETYPE_UNKNOWN
}

func (e *UnreserveType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UnreserveType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UnreserveType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UnreserveType) Ref() *UnreserveType {
	return &e
}

/*
Group host-NICs to function as a singular entity
*/
type UplinkBond struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of Prism Element cluster that the host belongs to
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of host-NIC UUIDs in this bond
	*/
	HostNicReferences []string `json:"hostNicReferences,omitempty"`
	/*
	  Host UUID for the bond
	*/
	HostReference *string `json:"hostReference,omitempty"`

	LacpStatus *UplinkBondLacpStatus `json:"lacpStatus,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the bond
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *UplinkBondType `json:"type,omitempty"`

	VirtualSwitchInfo *UplinkBondVirtualSwitchInfo `json:"virtualSwitchInfo,omitempty"`
}

func NewUplinkBond() *UplinkBond {
	p := new(UplinkBond)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.UplinkBond"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LACP status of the bond (configured, negotiated, or off)
*/
type UplinkBondLacpStatus int

const (
	UPLINKBONDLACPSTATUS_UNKNOWN    UplinkBondLacpStatus = 0
	UPLINKBONDLACPSTATUS_REDACTED   UplinkBondLacpStatus = 1
	UPLINKBONDLACPSTATUS_CONFIGURED UplinkBondLacpStatus = 2
	UPLINKBONDLACPSTATUS_NEGOTIATED UplinkBondLacpStatus = 3
	UPLINKBONDLACPSTATUS_NIL        UplinkBondLacpStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UplinkBondLacpStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIGURED",
		"NEGOTIATED",
		"NIL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UplinkBondLacpStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIGURED",
		"NEGOTIATED",
		"NIL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UplinkBondLacpStatus) index(name string) UplinkBondLacpStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIGURED",
		"NEGOTIATED",
		"NIL",
	}
	for idx := range names {
		if names[idx] == name {
			return UplinkBondLacpStatus(idx)
		}
	}
	return UPLINKBONDLACPSTATUS_UNKNOWN
}

func (e *UplinkBondLacpStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UplinkBondLacpStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UplinkBondLacpStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UplinkBondLacpStatus) Ref() *UplinkBondLacpStatus {
	return &e
}

/*
Type of the bond (active-backup, balance-slb, or balance-tcp)
*/
type UplinkBondType int

const (
	UPLINKBONDTYPE_UNKNOWN       UplinkBondType = 0
	UPLINKBONDTYPE_REDACTED      UplinkBondType = 1
	UPLINKBONDTYPE_ACTIVE_BACKUP UplinkBondType = 2
	UPLINKBONDTYPE_BALANCE_SLB   UplinkBondType = 3
	UPLINKBONDTYPE_BALANCE_TCP   UplinkBondType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UplinkBondType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UplinkBondType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UplinkBondType) index(name string) UplinkBondType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_BACKUP",
		"BALANCE_SLB",
		"BALANCE_TCP",
	}
	for idx := range names {
		if names[idx] == name {
			return UplinkBondType(idx)
		}
	}
	return UPLINKBONDTYPE_UNKNOWN
}

func (e *UplinkBondType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UplinkBondType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UplinkBondType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UplinkBondType) Ref() *UplinkBondType {
	return &e
}

/*
Virtual Switch info associated with bond (if any)
*/
type UplinkBondVirtualSwitchInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the Virtual Switch that this bond is a part of
	*/
	Name *string `json:"name,omitempty"`
	/*
	  UUID of the Virtual Switch this bond is a part of
	*/
	Reference *string `json:"reference,omitempty"`
}

func NewUplinkBondVirtualSwitchInfo() *UplinkBondVirtualSwitchInfo {
	p := new(UplinkBondVirtualSwitchInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.UplinkBondVirtualSwitchInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual IP configuration for the load balancer session listener.
*/
type VirtualIP struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AssignmentType *AssignmentType `json:"assignmentType"`

	IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
	/*
	  Reference to the subnet from which the virtual IP address is allocated.
	*/
	SubnetReference *string `json:"subnetReference"`
}

func (p *VirtualIP) MarshalJSON() ([]byte, error) {
	type VirtualIPProxy VirtualIP
	return json.Marshal(struct {
		*VirtualIPProxy
		AssignmentType  *AssignmentType `json:"assignmentType,omitempty"`
		SubnetReference *string         `json:"subnetReference,omitempty"`
	}{
		VirtualIPProxy:  (*VirtualIPProxy)(p),
		AssignmentType:  p.AssignmentType,
		SubnetReference: p.SubnetReference,
	})
}

func NewVirtualIP() *VirtualIP {
	p := new(VirtualIP)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VirtualIP"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schema to configure a virtual switch
*/
type VirtualSwitch struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BondMode *BondModeType `json:"bondMode"`
	/*
	  Cluster configuration list
	*/
	Clusters []Cluster `json:"clusters"`
	/*
	  Input body to configure a Virtual Switch
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the virtual switch's delete is being processed
	*/
	HasDeleteInProgress *bool `json:"hasDeleteInProgress,omitempty"`
	/*
	  When true, virtual switch configuration is not deployed on every node.
	*/
	HasDeploymentError *bool `json:"hasDeploymentError,omitempty"`
	/*
	  Indicates whether the virtual switch's update is being processed
	*/
	HasUpdateInProgress *bool `json:"hasUpdateInProgress,omitempty"`

	IgmpSpec *IgmpSpec `json:"igmpSpec,omitempty"`
	/*
	  Indicates whether it is a default Virtual Switch which cannot be deleted
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  When true, the node is not put in maintenance mode during the create/update operation.
	*/
	IsQuickMode *bool `json:"isQuickMode,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  MTU
	*/
	Mtu *int64 `json:"mtu,omitempty"`
	/*
	  User-visible Virtual Switch name
	*/
	Name *string `json:"name"`

	OwnerType *OwnerType `json:"ownerType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VirtualSwitch) MarshalJSON() ([]byte, error) {
	type VirtualSwitchProxy VirtualSwitch
	return json.Marshal(struct {
		*VirtualSwitchProxy
		BondMode *BondModeType `json:"bondMode,omitempty"`
		Clusters []Cluster     `json:"clusters,omitempty"`
		Name     *string       `json:"name,omitempty"`
	}{
		VirtualSwitchProxy: (*VirtualSwitchProxy)(p),
		BondMode:           p.BondMode,
		Clusters:           p.Clusters,
		Name:               p.Name,
	})
}

func NewVirtualSwitch() *VirtualSwitch {
	p := new(VirtualSwitch)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VirtualSwitch"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefault = new(bool)
	*p.IsDefault = false
	p.IsQuickMode = new(bool)
	*p.IsQuickMode = false
	p.Mtu = new(int64)
	*p.Mtu = 1500

	return p
}

type VirtualSwitchProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BondMode *BondModeType `json:"bondMode"`
	/*
	  Cluster configuration list
	*/
	Clusters []Cluster `json:"clusters"`
	/*
	  Input body to configure a Virtual Switch
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the virtual switch's delete is being processed
	*/
	HasDeleteInProgress *bool `json:"hasDeleteInProgress,omitempty"`
	/*
	  When true, virtual switch configuration is not deployed on every node.
	*/
	HasDeploymentError *bool `json:"hasDeploymentError,omitempty"`
	/*
	  Indicates whether the virtual switch's update is being processed
	*/
	HasUpdateInProgress *bool `json:"hasUpdateInProgress,omitempty"`

	IgmpSpec *IgmpSpec `json:"igmpSpec,omitempty"`
	/*
	  Indicates whether it is a default Virtual Switch which cannot be deleted
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  When true, the node is not put in maintenance mode during the create/update operation.
	*/
	IsQuickMode *bool `json:"isQuickMode,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  MTU
	*/
	Mtu *int64 `json:"mtu,omitempty"`
	/*
	  User-visible Virtual Switch name
	*/
	Name *string `json:"name"`

	OwnerType *OwnerType `json:"ownerType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VirtualSwitchProjection) MarshalJSON() ([]byte, error) {
	type VirtualSwitchProjectionProxy VirtualSwitchProjection
	return json.Marshal(struct {
		*VirtualSwitchProjectionProxy
		BondMode *BondModeType `json:"bondMode,omitempty"`
		Clusters []Cluster     `json:"clusters,omitempty"`
		Name     *string       `json:"name,omitempty"`
	}{
		VirtualSwitchProjectionProxy: (*VirtualSwitchProjectionProxy)(p),
		BondMode:                     p.BondMode,
		Clusters:                     p.Clusters,
		Name:                         p.Name,
	})
}

func NewVirtualSwitchProjection() *VirtualSwitchProjection {
	p := new(VirtualSwitchProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VirtualSwitchProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsDefault = new(bool)
	*p.IsDefault = false
	p.IsQuickMode = new(bool)
	*p.IsQuickMode = false
	p.Mtu = new(int64)
	*p.Mtu = 1500

	return p
}

/*
Schema for VLAN subnet migration.
*/
type VlanSubnetMigrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VLAN subnet details that need to be migrated.
	*/
	Subnets []SubnetInfo `json:"subnets"`
}

func (p *VlanSubnetMigrationSpec) MarshalJSON() ([]byte, error) {
	type VlanSubnetMigrationSpecProxy VlanSubnetMigrationSpec
	return json.Marshal(struct {
		*VlanSubnetMigrationSpecProxy
		Subnets []SubnetInfo `json:"subnets,omitempty"`
	}{
		VlanSubnetMigrationSpecProxy: (*VlanSubnetMigrationSpecProxy)(p),
		Subnets:                      p.Subnets,
	})
}

func NewVlanSubnetMigrationSpec() *VlanSubnetMigrationSpec {
	p := new(VlanSubnetMigrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VlanSubnetMigrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual Machine properties
*/
type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Virtual Machine name
	*/
	Name *string `json:"name,omitempty"`
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Vm"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtal NIC for projections
*/
type VmNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Private IP value in string
	*/
	PrivateIp *string `json:"privateIp,omitempty"`
}

func NewVmNic() *VmNic {
	p := new(VmNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VmNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VM NIC and VPC to which the VM NIC subnet belongs.
*/
type VmNicAssociation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VM NIC reference.
	*/
	VmNicReference *string `json:"vmNicReference"`
	/*
	  VPC reference to which the VM NIC subnet belongs.
	*/
	VpcReference *string `json:"vpcReference,omitempty"`
}

func (p *VmNicAssociation) MarshalJSON() ([]byte, error) {
	type VmNicAssociationProxy VmNicAssociation
	return json.Marshal(struct {
		*VmNicAssociationProxy
		VmNicReference *string `json:"vmNicReference,omitempty"`
	}{
		VmNicAssociationProxy: (*VmNicAssociationProxy)(p),
		VmNicReference:        p.VmNicReference,
	})
}

func NewVmNicAssociation() *VmNicAssociation {
	p := new(VmNicAssociation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VmNicAssociation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmNicProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Private IP value in string
	*/
	PrivateIp *string `json:"privateIp,omitempty"`
}

func NewVmNicProjection() *VmNicProjection {
	p := new(VmNicProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VmNicProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Virtual Machine name
	*/
	Name *string `json:"name,omitempty"`
}

func NewVmProjection() *VmProjection {
	p := new(VmProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VmProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Virtual NIC object.
*/
type Vnic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The assigned IPv4 addresses on the virtual NIC.
	*/
	AssignedIpv4Addresses []import1.IPAddress `json:"assignedIpv4Addresses,omitempty"`
	/*
	  The assigned secondary IPv4 addresses on the virtual NIC.
	*/
	AssignedSecondaryIpv4Addresses []import1.IPAddress `json:"assignedSecondaryIpv4Addresses,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The learned IPv4 addresses on the virtual NIC.
	*/
	LearnedIpv4Addresses []import1.IPAddress `json:"learnedIpv4Addresses,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The MAC address assigned to the virtual NIC.
	*/
	MacAddress *string `json:"macAddress,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The VM UUID of the virtual NIC.
	*/
	VmReference *string `json:"vmReference,omitempty"`
}

func NewVnic() *Vnic {
	p := new(Vnic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Vnic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VnicMigrationItemSpec struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  NIC network UUID for the destination subnet.
	*/
	NetworkUuid *string `json:"networkUuid,omitempty"`

	RequestedIpAddresses []import1.IPAddress `json:"requestedIpAddresses,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVnicMigrationItemSpec() *VnicMigrationItemSpec {
	p := new(VnicMigrationItemSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VnicMigrationItemSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schema for Virtual Private Cloud (VPC).
*/
type Vpc struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CommonDhcpOptions *VpcDhcpOptions `json:"commonDhcpOptions,omitempty"`
	/*
	  Description of the VPC.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  External routing domain associated with this route table
	*/
	ExternalRoutingDomainReference *string `json:"externalRoutingDomainReference,omitempty"`
	/*
	  List of external subnets that the VPC is attached to.
	*/
	ExternalSubnets []ExternalSubnet `json:"externalSubnets,omitempty"`
	/*
	  CIDR blocks from the VPC which can talk externally without performing NAT. This is applicable when connecting to external subnets which have disabled NAT.
	*/
	ExternallyRoutablePrefixes []IPSubnet `json:"externallyRoutablePrefixes,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the VPC.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  List of IP Addresses used for SNAT.
	*/
	SnatIps []import1.IPAddress `json:"snatIps,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VpcType *VpcType `json:"vpcType,omitempty"`
}

func NewVpc() *Vpc {
	p := new(Vpc)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Vpc"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of DHCP options to be configured.
*/
type VpcDhcpOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Domain Name Server addresses (option 6).
	*/
	DomainNameServers []import1.IPAddress `json:"domainNameServers,omitempty"`
}

func NewVpcDhcpOptions() *VpcDhcpOptions {
	p := new(VpcDhcpOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpcDhcpOptions"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPC name for projections
*/
type VpcName struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VPC name
	*/
	Name *string `json:"name,omitempty"`
}

func NewVpcName() *VpcName {
	p := new(VpcName)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpcName"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VpcNameProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VPC name
	*/
	Name *string `json:"name,omitempty"`
}

func NewVpcNameProjection() *VpcNameProjection {
	p := new(VpcNameProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpcNameProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VpcProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CommonDhcpOptions *VpcDhcpOptions `json:"commonDhcpOptions,omitempty"`
	/*
	  Description of the VPC.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  External routing domain associated with this route table
	*/
	ExternalRoutingDomainReference *string `json:"externalRoutingDomainReference,omitempty"`
	/*
	  List of external subnets that the VPC is attached to.
	*/
	ExternalSubnets []ExternalSubnet `json:"externalSubnets,omitempty"`
	/*
	  CIDR blocks from the VPC which can talk externally without performing NAT. This is applicable when connecting to external subnets which have disabled NAT.
	*/
	ExternallyRoutablePrefixes []IPSubnet `json:"externallyRoutablePrefixes,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  Name of the VPC.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  List of IP Addresses used for SNAT.
	*/
	SnatIps []import1.IPAddress `json:"snatIps,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VpcType *VpcType `json:"vpcType,omitempty"`
}

func NewVpcProjection() *VpcProjection {
	p := new(VpcProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpcProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of VPC.
*/
type VpcType int

const (
	VPCTYPE_UNKNOWN  VpcType = 0
	VPCTYPE_REDACTED VpcType = 1
	VPCTYPE_REGULAR  VpcType = 2
	VPCTYPE_TRANSIT  VpcType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *VpcType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGULAR",
		"TRANSIT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e VpcType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGULAR",
		"TRANSIT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *VpcType) index(name string) VpcType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGULAR",
		"TRANSIT",
	}
	for idx := range names {
		if names[idx] == name {
			return VpcType(idx)
		}
	}
	return VPCTYPE_UNKNOWN
}

func (e *VpcType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VpcType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VpcType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VpcType) Ref() *VpcType {
	return &e
}

type VpcVirtualSwitchMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Whether to permit all traffic through virtual switch or only the ICMP and statistics collection requests.
	*/
	IsAllTrafficPermitted *bool `json:"isAllTrafficPermitted,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  UUID of the virtual switch.
	*/
	VirtualSwitchUuid *string `json:"virtualSwitchUuid"`
}

func (p *VpcVirtualSwitchMapping) MarshalJSON() ([]byte, error) {
	type VpcVirtualSwitchMappingProxy VpcVirtualSwitchMapping
	return json.Marshal(struct {
		*VpcVirtualSwitchMappingProxy
		VirtualSwitchUuid *string `json:"virtualSwitchUuid,omitempty"`
	}{
		VpcVirtualSwitchMappingProxy: (*VpcVirtualSwitchMappingProxy)(p),
		VirtualSwitchUuid:            p.VirtualSwitchUuid,
	})
}

func NewVpcVirtualSwitchMapping() *VpcVirtualSwitchMapping {
	p := new(VpcVirtualSwitchMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpcVirtualSwitchMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Third-party VPN appliance.
*/
type VpnAppliance struct {
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
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  VPN appliance name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPN appliance version.
	*/
	Version *string `json:"version,omitempty"`
}

func NewVpnAppliance() *VpnAppliance {
	p := new(VpnAppliance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpnAppliance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPN connection
*/
type VpnConnection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  IP prefixes advertised to the remote gateway over BGP.
	*/
	AdvertisedPrefixes []IPSubnet `json:"advertisedPrefixes,omitempty"`
	/*
	  VPN connection description
	*/
	Description *string `json:"description,omitempty"`

	DpdConfig *DpdConfig `json:"dpdConfig,omitempty"`
	/*
	  Priority assigned to routes received on this connection over eBGP. A higher priority value indicates that the routes are more preferred
	*/
	DynamicRoutePriority *int `json:"dynamicRoutePriority,omitempty"`

	EbgpStatus *Status `json:"ebgpStatus,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IpsecConfig *IpsecConfig `json:"ipsecConfig"`

	IpsecTunnelStatus *Status `json:"ipsecTunnelStatus,omitempty"`
	/*
	  IP prefixes learned from the remote gateway over BGP.
	*/
	LearnedPrefixes []IPSubnet `json:"learnedPrefixes,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The local VPN gateway reference
	*/
	LocalGatewayReference *string `json:"localGatewayReference"`

	LocalGatewayRole *GatewayRole `json:"localGatewayRole"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  VPN connection name
	*/
	Name *string `json:"name"`

	QosConfig *QosConfig `json:"qosConfig,omitempty"`
	/*
	  The remote VPN gateway reference
	*/
	RemoteGatewayReference *string `json:"remoteGatewayReference"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VpnConnection) MarshalJSON() ([]byte, error) {
	type VpnConnectionProxy VpnConnection
	return json.Marshal(struct {
		*VpnConnectionProxy
		IpsecConfig            *IpsecConfig `json:"ipsecConfig,omitempty"`
		LocalGatewayReference  *string      `json:"localGatewayReference,omitempty"`
		LocalGatewayRole       *GatewayRole `json:"localGatewayRole,omitempty"`
		Name                   *string      `json:"name,omitempty"`
		RemoteGatewayReference *string      `json:"remoteGatewayReference,omitempty"`
	}{
		VpnConnectionProxy:     (*VpnConnectionProxy)(p),
		IpsecConfig:            p.IpsecConfig,
		LocalGatewayReference:  p.LocalGatewayReference,
		LocalGatewayRole:       p.LocalGatewayRole,
		Name:                   p.Name,
		RemoteGatewayReference: p.RemoteGatewayReference,
	})
}

func NewVpnConnection() *VpnConnection {
	p := new(VpnConnection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.VpnConnection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VXLAN Tunnel Endpoint.
*/
type Vtep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddress `json:"address,omitempty"`
}

func NewVtep() *Vtep {
	p := new(Vtep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.config.Vtep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetRouteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Route                 `json:"-"`
}

func NewOneOfGetRouteApiResponseData() *OneOfGetRouteApiResponseData {
	p := new(OneOfGetRouteApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRouteApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRouteApiResponseData is nil"))
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
	case Route:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Route)
		}
		*p.oneOfType0 = v.(Route)
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

func (p *OneOfGetRouteApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRouteApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Route)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.Route" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Route)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRouteApiResponseData"))
}

func (p *OneOfGetRouteApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRouteApiResponseData")
}

type OneOfRemoteVpnConnectionListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []RemoteVpnConnection  `json:"-"`
}

func NewOneOfRemoteVpnConnectionListApiResponseData() *OneOfRemoteVpnConnectionListApiResponseData {
	p := new(OneOfRemoteVpnConnectionListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoteVpnConnectionListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoteVpnConnectionListApiResponseData is nil"))
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
	case []RemoteVpnConnection:
		p.oneOfType0 = v.([]RemoteVpnConnection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RemoteVpnConnection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RemoteVpnConnection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRemoteVpnConnectionListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.RemoteVpnConnection>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfRemoteVpnConnectionListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]RemoteVpnConnection)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.RemoteVpnConnection" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RemoteVpnConnection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RemoteVpnConnection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoteVpnConnectionListApiResponseData"))
}

func (p *OneOfRemoteVpnConnectionListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.RemoteVpnConnection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRemoteVpnConnectionListApiResponseData")
}

type OneOfListLearnedMacAddressesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []LearnedMacAddress    `json:"-"`
}

func NewOneOfListLearnedMacAddressesApiResponseData() *OneOfListLearnedMacAddressesApiResponseData {
	p := new(OneOfListLearnedMacAddressesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLearnedMacAddressesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLearnedMacAddressesApiResponseData is nil"))
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
	case []LearnedMacAddress:
		p.oneOfType0 = v.([]LearnedMacAddress)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.LearnedMacAddress>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.LearnedMacAddress>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLearnedMacAddressesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.LearnedMacAddress>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListLearnedMacAddressesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]LearnedMacAddress)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.LearnedMacAddress" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.LearnedMacAddress>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.LearnedMacAddress>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLearnedMacAddressesApiResponseData"))
}

func (p *OneOfListLearnedMacAddressesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.LearnedMacAddress>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListLearnedMacAddressesApiResponseData")
}

type OneOfGetLoadBalancerSessionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *LoadBalancerSession   `json:"-"`
}

func NewOneOfGetLoadBalancerSessionApiResponseData() *OneOfGetLoadBalancerSessionApiResponseData {
	p := new(OneOfGetLoadBalancerSessionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLoadBalancerSessionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLoadBalancerSessionApiResponseData is nil"))
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
	case LoadBalancerSession:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LoadBalancerSession)
		}
		*p.oneOfType0 = v.(LoadBalancerSession)
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

func (p *OneOfGetLoadBalancerSessionApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetLoadBalancerSessionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(LoadBalancerSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.LoadBalancerSession" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LoadBalancerSession)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLoadBalancerSessionApiResponseData"))
}

func (p *OneOfGetLoadBalancerSessionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetLoadBalancerSessionApiResponseData")
}

type OneOfListRouteTablesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []RouteTable           `json:"-"`
}

func NewOneOfListRouteTablesApiResponseData() *OneOfListRouteTablesApiResponseData {
	p := new(OneOfListRouteTablesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRouteTablesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRouteTablesApiResponseData is nil"))
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
	case []RouteTable:
		p.oneOfType0 = v.([]RouteTable)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RouteTable>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RouteTable>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRouteTablesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.RouteTable>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListRouteTablesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]RouteTable)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.RouteTable" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RouteTable>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RouteTable>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRouteTablesApiResponseData"))
}

func (p *OneOfListRouteTablesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.RouteTable>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListRouteTablesApiResponseData")
}

type OneOfListRemoteSubnetsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []RemoteSubnet         `json:"-"`
}

func NewOneOfListRemoteSubnetsApiResponseData() *OneOfListRemoteSubnetsApiResponseData {
	p := new(OneOfListRemoteSubnetsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRemoteSubnetsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRemoteSubnetsApiResponseData is nil"))
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
	case []RemoteSubnet:
		p.oneOfType0 = v.([]RemoteSubnet)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RemoteSubnet>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RemoteSubnet>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRemoteSubnetsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.RemoteSubnet>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListRemoteSubnetsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]RemoteSubnet)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.RemoteSubnet" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RemoteSubnet>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RemoteSubnet>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRemoteSubnetsApiResponseData"))
}

func (p *OneOfListRemoteSubnetsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.RemoteSubnet>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListRemoteSubnetsApiResponseData")
}

type OneOfRoutingPolicyMatchConditionProtocolParameters struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType1    *ICMPObject              `json:"-"`
	oneOfType2    *ProtocolNumberObject    `json:"-"`
	oneOfType0    *LayerFourProtocolObject `json:"-"`
}

func NewOneOfRoutingPolicyMatchConditionProtocolParameters() *OneOfRoutingPolicyMatchConditionProtocolParameters {
	p := new(OneOfRoutingPolicyMatchConditionProtocolParameters)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRoutingPolicyMatchConditionProtocolParameters is nil"))
	}
	switch v.(type) {
	case ICMPObject:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ICMPObject)
		}
		*p.oneOfType1 = v.(ICMPObject)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case ProtocolNumberObject:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(ProtocolNumberObject)
		}
		*p.oneOfType2 = v.(ProtocolNumberObject)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case LayerFourProtocolObject:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LayerFourProtocolObject)
		}
		*p.oneOfType0 = v.(LayerFourProtocolObject)
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

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) GetValue() interface{} {
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

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(ICMPObject)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "networking.v4.config.ICMPObject" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ICMPObject)
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
	vOneOfType2 := new(ProtocolNumberObject)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "networking.v4.config.ProtocolNumberObject" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(ProtocolNumberObject)
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
	vOneOfType0 := new(LayerFourProtocolObject)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.LayerFourProtocolObject" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LayerFourProtocolObject)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyMatchConditionProtocolParameters"))
}

func (p *OneOfRoutingPolicyMatchConditionProtocolParameters) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRoutingPolicyMatchConditionProtocolParameters")
}

type OneOfGetVirtualSwitchApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *VirtualSwitch         `json:"-"`
}

func NewOneOfGetVirtualSwitchApiResponseData() *OneOfGetVirtualSwitchApiResponseData {
	p := new(OneOfGetVirtualSwitchApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVirtualSwitchApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVirtualSwitchApiResponseData is nil"))
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
	case VirtualSwitch:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VirtualSwitch)
		}
		*p.oneOfType0 = v.(VirtualSwitch)
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

func (p *OneOfGetVirtualSwitchApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetVirtualSwitchApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VirtualSwitch)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.VirtualSwitch" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VirtualSwitch)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVirtualSwitchApiResponseData"))
}

func (p *OneOfGetVirtualSwitchApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetVirtualSwitchApiResponseData")
}

type OneOfListFloatingIpsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []FloatingIp           `json:"-"`
	oneOfType401  []FloatingIpProjection `json:"-"`
}

func NewOneOfListFloatingIpsApiResponseData() *OneOfListFloatingIpsApiResponseData {
	p := new(OneOfListFloatingIpsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListFloatingIpsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListFloatingIpsApiResponseData is nil"))
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
	case []FloatingIp:
		p.oneOfType0 = v.([]FloatingIp)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.FloatingIp>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.FloatingIp>"
	case []FloatingIpProjection:
		p.oneOfType401 = v.([]FloatingIpProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.FloatingIpProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.FloatingIpProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListFloatingIpsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.FloatingIp>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<networking.v4.config.FloatingIpProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListFloatingIpsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]FloatingIp)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.FloatingIp" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.FloatingIp>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.FloatingIp>"
			return nil
		}
	}
	vOneOfType401 := new([]FloatingIpProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.FloatingIpProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.FloatingIpProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.FloatingIpProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListFloatingIpsApiResponseData"))
}

func (p *OneOfListFloatingIpsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.FloatingIp>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<networking.v4.config.FloatingIpProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListFloatingIpsApiResponseData")
}

type OneOfListSubnetsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType401  []SubnetProjection     `json:"-"`
	oneOfType0    []Subnet               `json:"-"`
}

func NewOneOfListSubnetsApiResponseData() *OneOfListSubnetsApiResponseData {
	p := new(OneOfListSubnetsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSubnetsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSubnetsApiResponseData is nil"))
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
	case []SubnetProjection:
		p.oneOfType401 = v.([]SubnetProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.SubnetProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.SubnetProjection>"
	case []Subnet:
		p.oneOfType0 = v.([]Subnet)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Subnet>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Subnet>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSubnetsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.SubnetProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<networking.v4.config.Subnet>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSubnetsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]SubnetProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.SubnetProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.SubnetProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.SubnetProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Subnet)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Subnet" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Subnet>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Subnet>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSubnetsApiResponseData"))
}

func (p *OneOfListSubnetsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.SubnetProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<networking.v4.config.Subnet>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSubnetsApiResponseData")
}

type OneOfGetFloatingIpApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *FloatingIp            `json:"-"`
}

func NewOneOfGetFloatingIpApiResponseData() *OneOfGetFloatingIpApiResponseData {
	p := new(OneOfGetFloatingIpApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetFloatingIpApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetFloatingIpApiResponseData is nil"))
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
	case FloatingIp:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FloatingIp)
		}
		*p.oneOfType0 = v.(FloatingIp)
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

func (p *OneOfGetFloatingIpApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetFloatingIpApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(FloatingIp)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.FloatingIp" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(FloatingIp)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetFloatingIpApiResponseData"))
}

func (p *OneOfGetFloatingIpApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetFloatingIpApiResponseData")
}

type OneOfTaskReferenceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
}

func NewOneOfTaskReferenceApiResponseData() *OneOfTaskReferenceApiResponseData {
	p := new(OneOfTaskReferenceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTaskReferenceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTaskReferenceApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTaskReferenceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTaskReferenceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskReferenceApiResponseData"))
}

func (p *OneOfTaskReferenceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTaskReferenceApiResponseData")
}

type OneOfFloatingIpProjectionAssociation struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType2    *LoadBalancerSessionAssociation `json:"-"`
	oneOfType0    *VmNicAssociation               `json:"-"`
	oneOfType1    *PrivateIpAssociation           `json:"-"`
}

func NewOneOfFloatingIpProjectionAssociation() *OneOfFloatingIpProjectionAssociation {
	p := new(OneOfFloatingIpProjectionAssociation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFloatingIpProjectionAssociation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFloatingIpProjectionAssociation is nil"))
	}
	switch v.(type) {
	case LoadBalancerSessionAssociation:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(LoadBalancerSessionAssociation)
		}
		*p.oneOfType2 = v.(LoadBalancerSessionAssociation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case VmNicAssociation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VmNicAssociation)
		}
		*p.oneOfType0 = v.(VmNicAssociation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case PrivateIpAssociation:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(PrivateIpAssociation)
		}
		*p.oneOfType1 = v.(PrivateIpAssociation)
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

func (p *OneOfFloatingIpProjectionAssociation) GetValue() interface{} {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfFloatingIpProjectionAssociation) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(LoadBalancerSessionAssociation)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "networking.v4.config.LoadBalancerSessionAssociation" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(LoadBalancerSessionAssociation)
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
	vOneOfType0 := new(VmNicAssociation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.VmNicAssociation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VmNicAssociation)
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
	vOneOfType1 := new(PrivateIpAssociation)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "networking.v4.config.PrivateIpAssociation" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(PrivateIpAssociation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFloatingIpProjectionAssociation"))
}

func (p *OneOfFloatingIpProjectionAssociation) MarshalJSON() ([]byte, error) {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfFloatingIpProjectionAssociation")
}

type OneOfListSubnetVnicsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Vnic                 `json:"-"`
}

func NewOneOfListSubnetVnicsApiResponseData() *OneOfListSubnetVnicsApiResponseData {
	p := new(OneOfListSubnetVnicsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSubnetVnicsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSubnetVnicsApiResponseData is nil"))
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
	case []Vnic:
		p.oneOfType0 = v.([]Vnic)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Vnic>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Vnic>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSubnetVnicsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.Vnic>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSubnetVnicsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Vnic)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Vnic" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Vnic>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Vnic>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSubnetVnicsApiResponseData"))
}

func (p *OneOfListSubnetVnicsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.Vnic>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSubnetVnicsApiResponseData")
}

type OneOfGetUplinkBondApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *UplinkBond            `json:"-"`
}

func NewOneOfGetUplinkBondApiResponseData() *OneOfGetUplinkBondApiResponseData {
	p := new(OneOfGetUplinkBondApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUplinkBondApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUplinkBondApiResponseData is nil"))
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
	case UplinkBond:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UplinkBond)
		}
		*p.oneOfType0 = v.(UplinkBond)
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

func (p *OneOfGetUplinkBondApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetUplinkBondApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(UplinkBond)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.UplinkBond" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UplinkBond)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUplinkBondApiResponseData"))
}

func (p *OneOfGetUplinkBondApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetUplinkBondApiResponseData")
}

type OneOfListRoutesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Route                `json:"-"`
}

func NewOneOfListRoutesApiResponseData() *OneOfListRoutesApiResponseData {
	p := new(OneOfListRoutesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRoutesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRoutesApiResponseData is nil"))
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
	case []Route:
		p.oneOfType0 = v.([]Route)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Route>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Route>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRoutesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.Route>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListRoutesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Route)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Route" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Route>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Route>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRoutesApiResponseData"))
}

func (p *OneOfListRoutesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.Route>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListRoutesApiResponseData")
}

type OneOfGatewayProjectionServices struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *RemoteNetworkServices `json:"-"`
	oneOfType0    *LocalNetworkServices  `json:"-"`
}

func NewOneOfGatewayProjectionServices() *OneOfGatewayProjectionServices {
	p := new(OneOfGatewayProjectionServices)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGatewayProjectionServices) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGatewayProjectionServices is nil"))
	}
	switch v.(type) {
	case RemoteNetworkServices:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(RemoteNetworkServices)
		}
		*p.oneOfType1 = v.(RemoteNetworkServices)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case LocalNetworkServices:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LocalNetworkServices)
		}
		*p.oneOfType0 = v.(LocalNetworkServices)
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

func (p *OneOfGatewayProjectionServices) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGatewayProjectionServices) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(RemoteNetworkServices)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "networking.v4.config.RemoteNetworkServices" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(RemoteNetworkServices)
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
	vOneOfType0 := new(LocalNetworkServices)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.LocalNetworkServices" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LocalNetworkServices)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGatewayProjectionServices"))
}

func (p *OneOfGatewayProjectionServices) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGatewayProjectionServices")
}

type OneOfGetRouteTableApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *RouteTable            `json:"-"`
}

func NewOneOfGetRouteTableApiResponseData() *OneOfGetRouteTableApiResponseData {
	p := new(OneOfGetRouteTableApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRouteTableApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRouteTableApiResponseData is nil"))
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
	case RouteTable:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RouteTable)
		}
		*p.oneOfType0 = v.(RouteTable)
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

func (p *OneOfGetRouteTableApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRouteTableApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(RouteTable)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.RouteTable" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RouteTable)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRouteTableApiResponseData"))
}

func (p *OneOfGetRouteTableApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRouteTableApiResponseData")
}

type OneOfListIPFIXExportersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []IPFIXExporter        `json:"-"`
}

func NewOneOfListIPFIXExportersApiResponseData() *OneOfListIPFIXExportersApiResponseData {
	p := new(OneOfListIPFIXExportersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListIPFIXExportersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListIPFIXExportersApiResponseData is nil"))
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
	case []IPFIXExporter:
		p.oneOfType0 = v.([]IPFIXExporter)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.IPFIXExporter>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.IPFIXExporter>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListIPFIXExportersApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.IPFIXExporter>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListIPFIXExportersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]IPFIXExporter)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.IPFIXExporter" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.IPFIXExporter>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.IPFIXExporter>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListIPFIXExportersApiResponseData"))
}

func (p *OneOfListIPFIXExportersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.IPFIXExporter>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListIPFIXExportersApiResponseData")
}

type OneOfListGatewaysApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType401  []GatewayProjection    `json:"-"`
	oneOfType0    []Gateway              `json:"-"`
}

func NewOneOfListGatewaysApiResponseData() *OneOfListGatewaysApiResponseData {
	p := new(OneOfListGatewaysApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListGatewaysApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListGatewaysApiResponseData is nil"))
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
	case []GatewayProjection:
		p.oneOfType401 = v.([]GatewayProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.GatewayProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.GatewayProjection>"
	case []Gateway:
		p.oneOfType0 = v.([]Gateway)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Gateway>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Gateway>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListGatewaysApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.GatewayProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<networking.v4.config.Gateway>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListGatewaysApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]GatewayProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.GatewayProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.GatewayProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.GatewayProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Gateway)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Gateway" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Gateway>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Gateway>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListGatewaysApiResponseData"))
}

func (p *OneOfListGatewaysApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.GatewayProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<networking.v4.config.Gateway>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListGatewaysApiResponseData")
}

type OneOfListBgpRoutesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []BgpRoute             `json:"-"`
}

func NewOneOfListBgpRoutesApiResponseData() *OneOfListBgpRoutesApiResponseData {
	p := new(OneOfListBgpRoutesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListBgpRoutesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListBgpRoutesApiResponseData is nil"))
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
	case []BgpRoute:
		p.oneOfType0 = v.([]BgpRoute)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.BgpRoute>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.BgpRoute>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListBgpRoutesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.BgpRoute>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListBgpRoutesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]BgpRoute)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.BgpRoute" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.BgpRoute>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.BgpRoute>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListBgpRoutesApiResponseData"))
}

func (p *OneOfListBgpRoutesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.BgpRoute>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListBgpRoutesApiResponseData")
}

type OneOfListLayer2StretchesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Layer2Stretch        `json:"-"`
}

func NewOneOfListLayer2StretchesApiResponseData() *OneOfListLayer2StretchesApiResponseData {
	p := new(OneOfListLayer2StretchesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLayer2StretchesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLayer2StretchesApiResponseData is nil"))
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
	case []Layer2Stretch:
		p.oneOfType0 = v.([]Layer2Stretch)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Layer2Stretch>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Layer2Stretch>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLayer2StretchesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.Layer2Stretch>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListLayer2StretchesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Layer2Stretch)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Layer2Stretch" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Layer2Stretch>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Layer2Stretch>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLayer2StretchesApiResponseData"))
}

func (p *OneOfListLayer2StretchesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.Layer2Stretch>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListLayer2StretchesApiResponseData")
}

type OneOfGetSubnetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Subnet                `json:"-"`
}

func NewOneOfGetSubnetApiResponseData() *OneOfGetSubnetApiResponseData {
	p := new(OneOfGetSubnetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSubnetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSubnetApiResponseData is nil"))
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
	case Subnet:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Subnet)
		}
		*p.oneOfType0 = v.(Subnet)
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

func (p *OneOfGetSubnetApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSubnetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Subnet)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.Subnet" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Subnet)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSubnetApiResponseData"))
}

func (p *OneOfGetSubnetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSubnetApiResponseData")
}

type OneOfGetTrafficMirrorApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *TrafficMirror         `json:"-"`
}

func NewOneOfGetTrafficMirrorApiResponseData() *OneOfGetTrafficMirrorApiResponseData {
	p := new(OneOfGetTrafficMirrorApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTrafficMirrorApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTrafficMirrorApiResponseData is nil"))
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
	case TrafficMirror:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TrafficMirror)
		}
		*p.oneOfType0 = v.(TrafficMirror)
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

func (p *OneOfGetTrafficMirrorApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetTrafficMirrorApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(TrafficMirror)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.TrafficMirror" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TrafficMirror)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTrafficMirrorApiResponseData"))
}

func (p *OneOfGetTrafficMirrorApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetTrafficMirrorApiResponseData")
}

type OneOfGetVpnConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *VpnConnection         `json:"-"`
}

func NewOneOfGetVpnConnectionApiResponseData() *OneOfGetVpnConnectionApiResponseData {
	p := new(OneOfGetVpnConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVpnConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVpnConnectionApiResponseData is nil"))
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
	case VpnConnection:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VpnConnection)
		}
		*p.oneOfType0 = v.(VpnConnection)
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

func (p *OneOfGetVpnConnectionApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetVpnConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VpnConnection)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.VpnConnection" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VpnConnection)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVpnConnectionApiResponseData"))
}

func (p *OneOfGetVpnConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetVpnConnectionApiResponseData")
}

type OneOfGatewayServices struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *RemoteNetworkServices `json:"-"`
	oneOfType0    *LocalNetworkServices  `json:"-"`
}

func NewOneOfGatewayServices() *OneOfGatewayServices {
	p := new(OneOfGatewayServices)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGatewayServices) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGatewayServices is nil"))
	}
	switch v.(type) {
	case RemoteNetworkServices:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(RemoteNetworkServices)
		}
		*p.oneOfType1 = v.(RemoteNetworkServices)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case LocalNetworkServices:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LocalNetworkServices)
		}
		*p.oneOfType0 = v.(LocalNetworkServices)
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

func (p *OneOfGatewayServices) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGatewayServices) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(RemoteNetworkServices)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "networking.v4.config.RemoteNetworkServices" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(RemoteNetworkServices)
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
	vOneOfType0 := new(LocalNetworkServices)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.LocalNetworkServices" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LocalNetworkServices)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGatewayServices"))
}

func (p *OneOfGatewayServices) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGatewayServices")
}

type OneOfFloatingIpAssociation struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType2    *LoadBalancerSessionAssociation `json:"-"`
	oneOfType0    *VmNicAssociation               `json:"-"`
	oneOfType1    *PrivateIpAssociation           `json:"-"`
}

func NewOneOfFloatingIpAssociation() *OneOfFloatingIpAssociation {
	p := new(OneOfFloatingIpAssociation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFloatingIpAssociation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFloatingIpAssociation is nil"))
	}
	switch v.(type) {
	case LoadBalancerSessionAssociation:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(LoadBalancerSessionAssociation)
		}
		*p.oneOfType2 = v.(LoadBalancerSessionAssociation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case VmNicAssociation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VmNicAssociation)
		}
		*p.oneOfType0 = v.(VmNicAssociation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case PrivateIpAssociation:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(PrivateIpAssociation)
		}
		*p.oneOfType1 = v.(PrivateIpAssociation)
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

func (p *OneOfFloatingIpAssociation) GetValue() interface{} {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfFloatingIpAssociation) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(LoadBalancerSessionAssociation)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "networking.v4.config.LoadBalancerSessionAssociation" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(LoadBalancerSessionAssociation)
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
	vOneOfType0 := new(VmNicAssociation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.VmNicAssociation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VmNicAssociation)
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
	vOneOfType1 := new(PrivateIpAssociation)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "networking.v4.config.PrivateIpAssociation" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(PrivateIpAssociation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFloatingIpAssociation"))
}

func (p *OneOfFloatingIpAssociation) MarshalJSON() ([]byte, error) {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfFloatingIpAssociation")
}

type OneOfGetRemoteSubnetApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *RemoteSubnet          `json:"-"`
}

func NewOneOfGetRemoteSubnetApiResponseData() *OneOfGetRemoteSubnetApiResponseData {
	p := new(OneOfGetRemoteSubnetApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRemoteSubnetApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRemoteSubnetApiResponseData is nil"))
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
	case RemoteSubnet:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RemoteSubnet)
		}
		*p.oneOfType0 = v.(RemoteSubnet)
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

func (p *OneOfGetRemoteSubnetApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRemoteSubnetApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(RemoteSubnet)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.RemoteSubnet" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RemoteSubnet)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRemoteSubnetApiResponseData"))
}

func (p *OneOfGetRemoteSubnetApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRemoteSubnetApiResponseData")
}

type OneOfListUplinkBondsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []UplinkBond           `json:"-"`
}

func NewOneOfListUplinkBondsApiResponseData() *OneOfListUplinkBondsApiResponseData {
	p := new(OneOfListUplinkBondsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUplinkBondsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUplinkBondsApiResponseData is nil"))
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
	case []UplinkBond:
		p.oneOfType0 = v.([]UplinkBond)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.UplinkBond>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.UplinkBond>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUplinkBondsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.UplinkBond>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUplinkBondsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]UplinkBond)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.UplinkBond" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.UplinkBond>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.UplinkBond>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUplinkBondsApiResponseData"))
}

func (p *OneOfListUplinkBondsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.UplinkBond>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUplinkBondsApiResponseData")
}

type OneOfGetBgpRouteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *BgpRoute              `json:"-"`
}

func NewOneOfGetBgpRouteApiResponseData() *OneOfGetBgpRouteApiResponseData {
	p := new(OneOfGetBgpRouteApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetBgpRouteApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetBgpRouteApiResponseData is nil"))
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
	case BgpRoute:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(BgpRoute)
		}
		*p.oneOfType0 = v.(BgpRoute)
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

func (p *OneOfGetBgpRouteApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetBgpRouteApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(BgpRoute)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.BgpRoute" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(BgpRoute)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBgpRouteApiResponseData"))
}

func (p *OneOfGetBgpRouteApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetBgpRouteApiResponseData")
}

type OneOfGetLayer2StretchApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Layer2Stretch         `json:"-"`
}

func NewOneOfGetLayer2StretchApiResponseData() *OneOfGetLayer2StretchApiResponseData {
	p := new(OneOfGetLayer2StretchApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLayer2StretchApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLayer2StretchApiResponseData is nil"))
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
	case Layer2Stretch:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Layer2Stretch)
		}
		*p.oneOfType0 = v.(Layer2Stretch)
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

func (p *OneOfGetLayer2StretchApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetLayer2StretchApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Layer2Stretch)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.Layer2Stretch" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Layer2Stretch)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLayer2StretchApiResponseData"))
}

func (p *OneOfGetLayer2StretchApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetLayer2StretchApiResponseData")
}

type OneOfGetVpcApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Vpc                   `json:"-"`
}

func NewOneOfGetVpcApiResponseData() *OneOfGetVpcApiResponseData {
	p := new(OneOfGetVpcApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVpcApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVpcApiResponseData is nil"))
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
	case Vpc:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Vpc)
		}
		*p.oneOfType0 = v.(Vpc)
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

func (p *OneOfGetVpcApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetVpcApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Vpc)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.Vpc" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Vpc)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVpcApiResponseData"))
}

func (p *OneOfGetVpcApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetVpcApiResponseData")
}

type OneOfListTrafficMirrorsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []TrafficMirror        `json:"-"`
}

func NewOneOfListTrafficMirrorsApiResponseData() *OneOfListTrafficMirrorsApiResponseData {
	p := new(OneOfListTrafficMirrorsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListTrafficMirrorsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListTrafficMirrorsApiResponseData is nil"))
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
	case []TrafficMirror:
		p.oneOfType0 = v.([]TrafficMirror)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.TrafficMirror>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.TrafficMirror>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListTrafficMirrorsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.TrafficMirror>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListTrafficMirrorsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]TrafficMirror)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.TrafficMirror" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.TrafficMirror>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.TrafficMirror>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListTrafficMirrorsApiResponseData"))
}

func (p *OneOfListTrafficMirrorsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.TrafficMirror>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListTrafficMirrorsApiResponseData")
}

type OneOfListSubnetReservedIpsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []ReservedIp           `json:"-"`
}

func NewOneOfListSubnetReservedIpsApiResponseData() *OneOfListSubnetReservedIpsApiResponseData {
	p := new(OneOfListSubnetReservedIpsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSubnetReservedIpsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSubnetReservedIpsApiResponseData is nil"))
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
	case []ReservedIp:
		p.oneOfType0 = v.([]ReservedIp)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.ReservedIp>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.ReservedIp>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSubnetReservedIpsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.ReservedIp>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSubnetReservedIpsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ReservedIp)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.ReservedIp" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.ReservedIp>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.ReservedIp>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSubnetReservedIpsApiResponseData"))
}

func (p *OneOfListSubnetReservedIpsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.ReservedIp>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSubnetReservedIpsApiResponseData")
}

type OneOfListVpnVendorConfigsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []VpnAppliance         `json:"-"`
}

func NewOneOfListVpnVendorConfigsApiResponseData() *OneOfListVpnVendorConfigsApiResponseData {
	p := new(OneOfListVpnVendorConfigsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVpnVendorConfigsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVpnVendorConfigsApiResponseData is nil"))
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
	case []VpnAppliance:
		p.oneOfType0 = v.([]VpnAppliance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VpnAppliance>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VpnAppliance>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVpnVendorConfigsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.VpnAppliance>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListVpnVendorConfigsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]VpnAppliance)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.VpnAppliance" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VpnAppliance>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VpnAppliance>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVpnVendorConfigsApiResponseData"))
}

func (p *OneOfListVpnVendorConfigsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.VpnAppliance>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListVpnVendorConfigsApiResponseData")
}

type OneOfListVpcsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType401  []VpcProjection        `json:"-"`
	oneOfType0    []Vpc                  `json:"-"`
}

func NewOneOfListVpcsApiResponseData() *OneOfListVpcsApiResponseData {
	p := new(OneOfListVpcsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVpcsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVpcsApiResponseData is nil"))
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
	case []VpcProjection:
		p.oneOfType401 = v.([]VpcProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VpcProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VpcProjection>"
	case []Vpc:
		p.oneOfType0 = v.([]Vpc)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.Vpc>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.Vpc>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVpcsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.VpcProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<networking.v4.config.Vpc>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListVpcsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]VpcProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.VpcProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VpcProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VpcProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Vpc)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.Vpc" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.Vpc>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.Vpc>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVpcsApiResponseData"))
}

func (p *OneOfListVpcsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.VpcProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<networking.v4.config.Vpc>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListVpcsApiResponseData")
}

type OneOfRemoteVtepGatewayListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []RemoteVtepGateway    `json:"-"`
}

func NewOneOfRemoteVtepGatewayListApiResponseData() *OneOfRemoteVtepGatewayListApiResponseData {
	p := new(OneOfRemoteVtepGatewayListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoteVtepGatewayListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoteVtepGatewayListApiResponseData is nil"))
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
	case []RemoteVtepGateway:
		p.oneOfType0 = v.([]RemoteVtepGateway)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RemoteVtepGateway>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RemoteVtepGateway>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRemoteVtepGatewayListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.RemoteVtepGateway>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfRemoteVtepGatewayListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]RemoteVtepGateway)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.RemoteVtepGateway" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RemoteVtepGateway>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RemoteVtepGateway>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoteVtepGatewayListApiResponseData"))
}

func (p *OneOfRemoteVtepGatewayListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.RemoteVtepGateway>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRemoteVtepGatewayListApiResponseData")
}

type OneOfGetGatewayApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Gateway               `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetGatewayApiResponseData() *OneOfGetGatewayApiResponseData {
	p := new(OneOfGetGatewayApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGatewayApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGatewayApiResponseData is nil"))
	}
	switch v.(type) {
	case Gateway:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Gateway)
		}
		*p.oneOfType0 = v.(Gateway)
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

func (p *OneOfGetGatewayApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetGatewayApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Gateway)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.Gateway" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Gateway)
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
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGatewayApiResponseData"))
}

func (p *OneOfGetGatewayApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetGatewayApiResponseData")
}

type OneOfGetBgpSessionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *BgpSession            `json:"-"`
}

func NewOneOfGetBgpSessionApiResponseData() *OneOfGetBgpSessionApiResponseData {
	p := new(OneOfGetBgpSessionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetBgpSessionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetBgpSessionApiResponseData is nil"))
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
	case BgpSession:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(BgpSession)
		}
		*p.oneOfType0 = v.(BgpSession)
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

func (p *OneOfGetBgpSessionApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetBgpSessionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(BgpSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.BgpSession" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(BgpSession)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBgpSessionApiResponseData"))
}

func (p *OneOfGetBgpSessionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetBgpSessionApiResponseData")
}

type OneOfGetNetworkControllerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *NetworkController     `json:"-"`
}

func NewOneOfGetNetworkControllerApiResponseData() *OneOfGetNetworkControllerApiResponseData {
	p := new(OneOfGetNetworkControllerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNetworkControllerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNetworkControllerApiResponseData is nil"))
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
	case NetworkController:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(NetworkController)
		}
		*p.oneOfType0 = v.(NetworkController)
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

func (p *OneOfGetNetworkControllerApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetNetworkControllerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(NetworkController)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.NetworkController" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(NetworkController)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNetworkControllerApiResponseData"))
}

func (p *OneOfGetNetworkControllerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetNetworkControllerApiResponseData")
}

type OneOfRemoteVpnConnectionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RemoteVpnConnection   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfRemoteVpnConnectionApiResponseData() *OneOfRemoteVpnConnectionApiResponseData {
	p := new(OneOfRemoteVpnConnectionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoteVpnConnectionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoteVpnConnectionApiResponseData is nil"))
	}
	switch v.(type) {
	case RemoteVpnConnection:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RemoteVpnConnection)
		}
		*p.oneOfType0 = v.(RemoteVpnConnection)
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

func (p *OneOfRemoteVpnConnectionApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRemoteVpnConnectionApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RemoteVpnConnection)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.RemoteVpnConnection" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RemoteVpnConnection)
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
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoteVpnConnectionApiResponseData"))
}

func (p *OneOfRemoteVpnConnectionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRemoteVpnConnectionApiResponseData")
}

type OneOfListVirtualSwitchesApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    []VirtualSwitch           `json:"-"`
	oneOfType400  *import3.ErrorResponse    `json:"-"`
	oneOfType401  []VirtualSwitchProjection `json:"-"`
}

func NewOneOfListVirtualSwitchesApiResponseData() *OneOfListVirtualSwitchesApiResponseData {
	p := new(OneOfListVirtualSwitchesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVirtualSwitchesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVirtualSwitchesApiResponseData is nil"))
	}
	switch v.(type) {
	case []VirtualSwitch:
		p.oneOfType0 = v.([]VirtualSwitch)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VirtualSwitch>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VirtualSwitch>"
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
	case []VirtualSwitchProjection:
		p.oneOfType401 = v.([]VirtualSwitchProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VirtualSwitchProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VirtualSwitchProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVirtualSwitchesApiResponseData) GetValue() interface{} {
	if "List<networking.v4.config.VirtualSwitch>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.VirtualSwitchProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListVirtualSwitchesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]VirtualSwitch)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.VirtualSwitch" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VirtualSwitch>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VirtualSwitch>"
			return nil
		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]VirtualSwitchProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.VirtualSwitchProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VirtualSwitchProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VirtualSwitchProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVirtualSwitchesApiResponseData"))
}

func (p *OneOfListVirtualSwitchesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<networking.v4.config.VirtualSwitch>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.VirtualSwitchProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListVirtualSwitchesApiResponseData")
}

type OneOfListBgpSessionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType401  []BgpSessionProjection `json:"-"`
	oneOfType0    []BgpSession           `json:"-"`
}

func NewOneOfListBgpSessionsApiResponseData() *OneOfListBgpSessionsApiResponseData {
	p := new(OneOfListBgpSessionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListBgpSessionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListBgpSessionsApiResponseData is nil"))
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
	case []BgpSessionProjection:
		p.oneOfType401 = v.([]BgpSessionProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.BgpSessionProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.BgpSessionProjection>"
	case []BgpSession:
		p.oneOfType0 = v.([]BgpSession)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.BgpSession>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.BgpSession>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListBgpSessionsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.BgpSessionProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<networking.v4.config.BgpSession>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListBgpSessionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]BgpSessionProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.BgpSessionProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.BgpSessionProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.BgpSessionProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]BgpSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.BgpSession" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.BgpSession>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.BgpSession>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListBgpSessionsApiResponseData"))
}

func (p *OneOfListBgpSessionsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.BgpSessionProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<networking.v4.config.BgpSession>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListBgpSessionsApiResponseData")
}

type OneOfListLoadBalancerSessionsApiResponseData struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType400  *import3.ErrorResponse          `json:"-"`
	oneOfType0    []LoadBalancerSession           `json:"-"`
	oneOfType401  []LoadBalancerSessionProjection `json:"-"`
}

func NewOneOfListLoadBalancerSessionsApiResponseData() *OneOfListLoadBalancerSessionsApiResponseData {
	p := new(OneOfListLoadBalancerSessionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLoadBalancerSessionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLoadBalancerSessionsApiResponseData is nil"))
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
	case []LoadBalancerSession:
		p.oneOfType0 = v.([]LoadBalancerSession)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.LoadBalancerSession>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.LoadBalancerSession>"
	case []LoadBalancerSessionProjection:
		p.oneOfType401 = v.([]LoadBalancerSessionProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.LoadBalancerSessionProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.LoadBalancerSessionProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLoadBalancerSessionsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.LoadBalancerSession>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<networking.v4.config.LoadBalancerSessionProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListLoadBalancerSessionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]LoadBalancerSession)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.LoadBalancerSession" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.LoadBalancerSession>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.LoadBalancerSession>"
			return nil
		}
	}
	vOneOfType401 := new([]LoadBalancerSessionProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.LoadBalancerSessionProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.LoadBalancerSessionProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.LoadBalancerSessionProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLoadBalancerSessionsApiResponseData"))
}

func (p *OneOfListLoadBalancerSessionsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.LoadBalancerSession>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<networking.v4.config.LoadBalancerSessionProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListLoadBalancerSessionsApiResponseData")
}

type OneOfListNetworkControllersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []NetworkController    `json:"-"`
}

func NewOneOfListNetworkControllersApiResponseData() *OneOfListNetworkControllersApiResponseData {
	p := new(OneOfListNetworkControllersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNetworkControllersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNetworkControllersApiResponseData is nil"))
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
	case []NetworkController:
		p.oneOfType0 = v.([]NetworkController)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.NetworkController>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.NetworkController>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNetworkControllersApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.NetworkController>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListNetworkControllersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]NetworkController)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.NetworkController" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.NetworkController>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.NetworkController>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNetworkControllersApiResponseData"))
}

func (p *OneOfListNetworkControllersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.NetworkController>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListNetworkControllersApiResponseData")
}

type OneOfListClusterCapabilitiesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []ClusterCapability    `json:"-"`
}

func NewOneOfListClusterCapabilitiesApiResponseData() *OneOfListClusterCapabilitiesApiResponseData {
	p := new(OneOfListClusterCapabilitiesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClusterCapabilitiesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClusterCapabilitiesApiResponseData is nil"))
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
	case []ClusterCapability:
		p.oneOfType0 = v.([]ClusterCapability)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.ClusterCapability>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.ClusterCapability>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListClusterCapabilitiesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.ClusterCapability>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListClusterCapabilitiesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ClusterCapability)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.ClusterCapability" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.ClusterCapability>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.ClusterCapability>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClusterCapabilitiesApiResponseData"))
}

func (p *OneOfListClusterCapabilitiesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.ClusterCapability>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListClusterCapabilitiesApiResponseData")
}

type OneOfListNodeSchedulableStatusesApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
	oneOfType0    []NodeSchedulableStatus `json:"-"`
}

func NewOneOfListNodeSchedulableStatusesApiResponseData() *OneOfListNodeSchedulableStatusesApiResponseData {
	p := new(OneOfListNodeSchedulableStatusesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListNodeSchedulableStatusesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListNodeSchedulableStatusesApiResponseData is nil"))
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
	case []NodeSchedulableStatus:
		p.oneOfType0 = v.([]NodeSchedulableStatus)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.NodeSchedulableStatus>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.NodeSchedulableStatus>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListNodeSchedulableStatusesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.NodeSchedulableStatus>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListNodeSchedulableStatusesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]NodeSchedulableStatus)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.NodeSchedulableStatus" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.NodeSchedulableStatus>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.NodeSchedulableStatus>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListNodeSchedulableStatusesApiResponseData"))
}

func (p *OneOfListNodeSchedulableStatusesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.NodeSchedulableStatus>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListNodeSchedulableStatusesApiResponseData")
}

type OneOfGetIPFIXExporterApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *IPFIXExporter         `json:"-"`
}

func NewOneOfGetIPFIXExporterApiResponseData() *OneOfGetIPFIXExporterApiResponseData {
	p := new(OneOfGetIPFIXExporterApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetIPFIXExporterApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetIPFIXExporterApiResponseData is nil"))
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
	case IPFIXExporter:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(IPFIXExporter)
		}
		*p.oneOfType0 = v.(IPFIXExporter)
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

func (p *OneOfGetIPFIXExporterApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetIPFIXExporterApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(IPFIXExporter)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.IPFIXExporter" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(IPFIXExporter)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetIPFIXExporterApiResponseData"))
}

func (p *OneOfGetIPFIXExporterApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetIPFIXExporterApiResponseData")
}

type OneOfListVpnConnectionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []VpnConnection        `json:"-"`
}

func NewOneOfListVpnConnectionsApiResponseData() *OneOfListVpnConnectionsApiResponseData {
	p := new(OneOfListVpnConnectionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVpnConnectionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVpnConnectionsApiResponseData is nil"))
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
	case []VpnConnection:
		p.oneOfType0 = v.([]VpnConnection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VpnConnection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VpnConnection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVpnConnectionsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.VpnConnection>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListVpnConnectionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]VpnConnection)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.VpnConnection" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VpnConnection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VpnConnection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVpnConnectionsApiResponseData"))
}

func (p *OneOfListVpnConnectionsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.VpnConnection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListVpnConnectionsApiResponseData")
}

type OneOfGetLearnedMacAddressApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *LearnedMacAddress     `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetLearnedMacAddressApiResponseData() *OneOfGetLearnedMacAddressApiResponseData {
	p := new(OneOfGetLearnedMacAddressApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLearnedMacAddressApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLearnedMacAddressApiResponseData is nil"))
	}
	switch v.(type) {
	case LearnedMacAddress:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LearnedMacAddress)
		}
		*p.oneOfType0 = v.(LearnedMacAddress)
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

func (p *OneOfGetLearnedMacAddressApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetLearnedMacAddressApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(LearnedMacAddress)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.LearnedMacAddress" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LearnedMacAddress)
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
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLearnedMacAddressApiResponseData"))
}

func (p *OneOfGetLearnedMacAddressApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetLearnedMacAddressApiResponseData")
}

type OneOfListRoutingPoliciesApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    []RoutingPolicy           `json:"-"`
	oneOfType400  *import3.ErrorResponse    `json:"-"`
	oneOfType401  []RoutingPolicyProjection `json:"-"`
}

func NewOneOfListRoutingPoliciesApiResponseData() *OneOfListRoutingPoliciesApiResponseData {
	p := new(OneOfListRoutingPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRoutingPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRoutingPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []RoutingPolicy:
		p.oneOfType0 = v.([]RoutingPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RoutingPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RoutingPolicy>"
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
	case []RoutingPolicyProjection:
		p.oneOfType401 = v.([]RoutingPolicyProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.RoutingPolicyProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.RoutingPolicyProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRoutingPoliciesApiResponseData) GetValue() interface{} {
	if "List<networking.v4.config.RoutingPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.RoutingPolicyProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRoutingPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]RoutingPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.RoutingPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RoutingPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RoutingPolicy>"
			return nil
		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]RoutingPolicyProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "networking.v4.config.RoutingPolicyProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.RoutingPolicyProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.RoutingPolicyProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRoutingPoliciesApiResponseData"))
}

func (p *OneOfListRoutingPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<networking.v4.config.RoutingPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.RoutingPolicyProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRoutingPoliciesApiResponseData")
}

type OneOfRemoteVtepGatewayApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RemoteVtepGateway     `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfRemoteVtepGatewayApiResponseData() *OneOfRemoteVtepGatewayApiResponseData {
	p := new(OneOfRemoteVtepGatewayApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoteVtepGatewayApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoteVtepGatewayApiResponseData is nil"))
	}
	switch v.(type) {
	case RemoteVtepGateway:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RemoteVtepGateway)
		}
		*p.oneOfType0 = v.(RemoteVtepGateway)
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

func (p *OneOfRemoteVtepGatewayApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRemoteVtepGatewayApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RemoteVtepGateway)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.RemoteVtepGateway" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RemoteVtepGateway)
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
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoteVtepGatewayApiResponseData"))
}

func (p *OneOfRemoteVtepGatewayApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRemoteVtepGatewayApiResponseData")
}

type OneOfGetRoutingPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *RoutingPolicy         `json:"-"`
}

func NewOneOfGetRoutingPolicyApiResponseData() *OneOfGetRoutingPolicyApiResponseData {
	p := new(OneOfGetRoutingPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRoutingPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRoutingPolicyApiResponseData is nil"))
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
	case RoutingPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RoutingPolicy)
		}
		*p.oneOfType0 = v.(RoutingPolicy)
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

func (p *OneOfGetRoutingPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRoutingPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(RoutingPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.config.RoutingPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RoutingPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRoutingPolicyApiResponseData"))
}

func (p *OneOfGetRoutingPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRoutingPolicyApiResponseData")
}

type OneOfListVpcVirtualSwitchMappingsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType400  *import3.ErrorResponse    `json:"-"`
	oneOfType0    []VpcVirtualSwitchMapping `json:"-"`
}

func NewOneOfListVpcVirtualSwitchMappingsApiResponseData() *OneOfListVpcVirtualSwitchMappingsApiResponseData {
	p := new(OneOfListVpcVirtualSwitchMappingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVpcVirtualSwitchMappingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVpcVirtualSwitchMappingsApiResponseData is nil"))
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
	case []VpcVirtualSwitchMapping:
		p.oneOfType0 = v.([]VpcVirtualSwitchMapping)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<networking.v4.config.VpcVirtualSwitchMapping>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<networking.v4.config.VpcVirtualSwitchMapping>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVpcVirtualSwitchMappingsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<networking.v4.config.VpcVirtualSwitchMapping>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListVpcVirtualSwitchMappingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]VpcVirtualSwitchMapping)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "networking.v4.config.VpcVirtualSwitchMapping" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<networking.v4.config.VpcVirtualSwitchMapping>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<networking.v4.config.VpcVirtualSwitchMapping>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVpcVirtualSwitchMappingsApiResponseData"))
}

func (p *OneOfListVpcVirtualSwitchMappingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<networking.v4.config.VpcVirtualSwitchMapping>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListVpcVirtualSwitchMappingsApiResponseData")
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
