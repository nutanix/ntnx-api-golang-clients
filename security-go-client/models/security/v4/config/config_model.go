/*
 * Generated file models/security/v4/config/config_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Security APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module security.v4.config of Nutanix Security APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/prism/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/common"
	import2 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/error"
	"time"
)

/*
REST response for all response codes in API path /security/v4.0.b1/config/public-keys Post operation
*/
type AddPublicKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddPublicKeyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddPublicKeyApiResponse() *AddPublicKeyApiResponse {
	p := new(AddPublicKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.AddPublicKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddPublicKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddPublicKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddPublicKeyApiResponseData()
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
Contains all required specifications for password change operation.
*/
type ChangePasswordSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Existing password of a user account.
	*/
	CurrentPassword *string `json:"currentPassword,omitempty"`
	/*
	  New password for a user account
	*/
	NewPassword *string `json:"newPassword"`
}

func (p *ChangePasswordSpec) MarshalJSON() ([]byte, error) {
	type ChangePasswordSpecProxy ChangePasswordSpec
	return json.Marshal(struct {
		*ChangePasswordSpecProxy
		NewPassword *string `json:"newPassword,omitempty"`
	}{
		ChangePasswordSpecProxy: (*ChangePasswordSpecProxy)(p),
		NewPassword:             p.NewPassword,
	})
}

func NewChangePasswordSpec() *ChangePasswordSpec {
	p := new(ChangePasswordSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ChangePasswordSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.0.b1/config/system-user-passwords/{extId}/$action/change-password Post operation
*/
type ChangeSystemUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangeSystemUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewChangeSystemUserPasswordApiResponse() *ChangeSystemUserPasswordApiResponse {
	p := new(ChangeSystemUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ChangeSystemUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ChangeSystemUserPasswordApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ChangeSystemUserPasswordApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfChangeSystemUserPasswordApiResponseData()
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
REST response for all response codes in API path /security/v4.0.b1/config/public-keys/{extId} Delete operation
*/
type DeletePublicKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeletePublicKeyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeletePublicKeyApiResponse() *DeletePublicKeyApiResponse {
	p := new(DeletePublicKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.DeletePublicKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeletePublicKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeletePublicKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeletePublicKeyApiResponseData()
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
Contains status of all the hardening settings for a cluster.
*/
type HardeningStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvScmaSchedule *Schedule `json:"ahvScmaSchedule,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Represents whether AHV defense knowledge consent banner is enabled on the hypervisor.
	*/
	IsAhvDefenseConsentBannerEnabled *bool `json:"isAhvDefenseConsentBannerEnabled,omitempty"`
	/*
	  Represents whether aide service is enabled on a cluster.
	*/
	IsAideEnabled *bool `json:"isAideEnabled,omitempty"`
	/*
	  Represents whether cluster lockdown mode is enabled on a cluster.
	*/
	IsClusterLockdownEnabled *bool `json:"isClusterLockdownEnabled,omitempty"`
	/*
	  Represents whether Nutanix CVM defense knowledge consent banner is enabled.
	*/
	IsCvmDefenseConsentBannerEnabled *bool `json:"isCvmDefenseConsentBannerEnabled,omitempty"`
	/*
	  Represents whether high strength password is enabled on a cluster.
	*/
	IsHighStrengthPasswordEnabled *bool `json:"isHighStrengthPasswordEnabled,omitempty"`
	/*
	  Represents whether log forwarding is enabled on a cluster.
	*/
	IsLogForwardingEnabled *bool `json:"isLogForwardingEnabled,omitempty"`
	/*
	  Represents the network segmentation settings on a cluster.
	*/
	IsNetworkSegmentationEnabled *bool `json:"isNetworkSegmentationEnabled,omitempty"`
	/*
	  Represents whether host secure boot is enabled on a cluster.
	*/
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	ScmaSchedule *Schedule `json:"scmaSchedule,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHardeningStatus() *HardeningStatus {
	p := new(HardeningStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.HardeningStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type HardeningStatusProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvScmaSchedule *Schedule `json:"ahvScmaSchedule,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterProjection *import4.ClusterProjection `json:"clusterProjection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Represents whether AHV defense knowledge consent banner is enabled on the hypervisor.
	*/
	IsAhvDefenseConsentBannerEnabled *bool `json:"isAhvDefenseConsentBannerEnabled,omitempty"`
	/*
	  Represents whether aide service is enabled on a cluster.
	*/
	IsAideEnabled *bool `json:"isAideEnabled,omitempty"`
	/*
	  Represents whether cluster lockdown mode is enabled on a cluster.
	*/
	IsClusterLockdownEnabled *bool `json:"isClusterLockdownEnabled,omitempty"`
	/*
	  Represents whether Nutanix CVM defense knowledge consent banner is enabled.
	*/
	IsCvmDefenseConsentBannerEnabled *bool `json:"isCvmDefenseConsentBannerEnabled,omitempty"`
	/*
	  Represents whether high strength password is enabled on a cluster.
	*/
	IsHighStrengthPasswordEnabled *bool `json:"isHighStrengthPasswordEnabled,omitempty"`
	/*
	  Represents whether log forwarding is enabled on a cluster.
	*/
	IsLogForwardingEnabled *bool `json:"isLogForwardingEnabled,omitempty"`
	/*
	  Represents the network segmentation settings on a cluster.
	*/
	IsNetworkSegmentationEnabled *bool `json:"isNetworkSegmentationEnabled,omitempty"`
	/*
	  Represents whether host secure boot is enabled on a cluster.
	*/
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	ScmaSchedule *Schedule `json:"scmaSchedule,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHardeningStatusProjection() *HardeningStatusProjection {
	p := new(HardeningStatusProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.HardeningStatusProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains the configuration for PC UI visibility status of all hardening settings.
*/
type HardeningVisibilitySetting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Represents the visibility of aide service state.
	*/
	IsAideVisible *bool `json:"isAideVisible,omitempty"`
	/*
	  Represents the visibility of cluster lockdown state.
	*/
	IsClusterLockdownVisible *bool `json:"isClusterLockdownVisible,omitempty"`
	/*
	  Represents the visibility of banner state.
	*/
	IsDefenseConsentBannerVisible *bool `json:"isDefenseConsentBannerVisible,omitempty"`
	/*
	  Represents the visibility of high strength password state.
	*/
	IsHighStrengthPasswordVisible *bool `json:"isHighStrengthPasswordVisible,omitempty"`
	/*
	  Represents the visibility of log forwarding state.
	*/
	IsLogForwardingVisible *bool `json:"isLogForwardingVisible,omitempty"`
	/*
	  Represents the visibility of network segmentation state.
	*/
	IsNetworkSegmentationVisible *bool `json:"isNetworkSegmentationVisible,omitempty"`
	/*
	  Represents the visibility of security configuration management automation state.
	*/
	IsScmaVisible *bool `json:"isScmaVisible,omitempty"`
	/*
	  Represents the visibility of host secure boot state.
	*/
	IsSecureBootVisible *bool `json:"isSecureBootVisible,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHardeningVisibilitySetting() *HardeningVisibilitySetting {
	p := new(HardeningVisibilitySetting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.HardeningVisibilitySetting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains possible values of password status.
*/
type PasswordStatus int

const (
	PASSWORDSTATUS_UNKNOWN         PasswordStatus = 0
	PASSWORDSTATUS_REDACTED        PasswordStatus = 1
	PASSWORDSTATUS_DEFAULT         PasswordStatus = 2
	PASSWORDSTATUS_SECURE          PasswordStatus = 3
	PASSWORDSTATUS_NOPASSWD        PasswordStatus = 4
	PASSWORDSTATUS_MULTIPLE_ISSUES PasswordStatus = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PasswordStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"SECURE",
		"NOPASSWD",
		"MULTIPLE_ISSUES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PasswordStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"SECURE",
		"NOPASSWD",
		"MULTIPLE_ISSUES",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PasswordStatus) index(name string) PasswordStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"SECURE",
		"NOPASSWD",
		"MULTIPLE_ISSUES",
	}
	for idx := range names {
		if names[idx] == name {
			return PasswordStatus(idx)
		}
	}
	return PASSWORDSTATUS_UNKNOWN
}

func (e *PasswordStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PasswordStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PasswordStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PasswordStatus) Ref() *PasswordStatus {
	return &e
}

/*
Verification status of the verify password operation.
*/
type PasswordVerificationStatus int

const (
	PASSWORDVERIFICATIONSTATUS_UNKNOWN        PasswordVerificationStatus = 0
	PASSWORDVERIFICATIONSTATUS_REDACTED       PasswordVerificationStatus = 1
	PASSWORDVERIFICATIONSTATUS_VERIFIED       PasswordVerificationStatus = 2
	PASSWORDVERIFICATIONSTATUS_NOT_VERIFIED   PasswordVerificationStatus = 3
	PASSWORDVERIFICATIONSTATUS_ACCOUNT_LOCKED PasswordVerificationStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PasswordVerificationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VERIFIED",
		"NOT_VERIFIED",
		"ACCOUNT_LOCKED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PasswordVerificationStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VERIFIED",
		"NOT_VERIFIED",
		"ACCOUNT_LOCKED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PasswordVerificationStatus) index(name string) PasswordVerificationStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VERIFIED",
		"NOT_VERIFIED",
		"ACCOUNT_LOCKED",
	}
	for idx := range names {
		if names[idx] == name {
			return PasswordVerificationStatus(idx)
		}
	}
	return PASSWORDVERIFICATIONSTATUS_UNKNOWN
}

func (e *PasswordVerificationStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PasswordVerificationStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PasswordVerificationStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PasswordVerificationStatus) Ref() *PasswordVerificationStatus {
	return &e
}

/*
Display model to encapsulate cluster public key information.
*/
type PublicKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The ssh key string information.
	*/
	Key *string `json:"key"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  The name of public key.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *PublicKey) MarshalJSON() ([]byte, error) {
	type PublicKeyProxy PublicKey
	return json.Marshal(struct {
		*PublicKeyProxy
		ClusterExtId *string `json:"clusterExtId,omitempty"`
		Key          *string `json:"key,omitempty"`
		Name         *string `json:"name,omitempty"`
	}{
		PublicKeyProxy: (*PublicKeyProxy)(p),
		ClusterExtId:   p.ClusterExtId,
		Key:            p.Key,
		Name:           p.Name,
	})
}

func NewPublicKey() *PublicKey {
	p := new(PublicKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.PublicKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PublicKeyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`

	ClusterProjection *import4.ClusterProjection `json:"clusterProjection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The ssh key string information.
	*/
	Key *string `json:"key"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  The name of public key.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *PublicKeyProjection) MarshalJSON() ([]byte, error) {
	type PublicKeyProjectionProxy PublicKeyProjection
	return json.Marshal(struct {
		*PublicKeyProjectionProxy
		ClusterExtId *string `json:"clusterExtId,omitempty"`
		Key          *string `json:"key,omitempty"`
		Name         *string `json:"name,omitempty"`
	}{
		PublicKeyProjectionProxy: (*PublicKeyProjectionProxy)(p),
		ClusterExtId:             p.ClusterExtId,
		Key:                      p.Key,
		Name:                     p.Name,
	})
}

func NewPublicKeyProjection() *PublicKeyProjection {
	p := new(PublicKeyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.PublicKeyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains possible values for a scheduling a task.
*/
type Schedule int

const (
	SCHEDULE_UNKNOWN  Schedule = 0
	SCHEDULE_REDACTED Schedule = 1
	SCHEDULE_HOURLY   Schedule = 2
	SCHEDULE_DAILY    Schedule = 3
	SCHEDULE_WEEKLY   Schedule = 4
	SCHEDULE_MONTHLY  Schedule = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Schedule) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Schedule) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Schedule) index(name string) Schedule {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	for idx := range names {
		if names[idx] == name {
			return Schedule(idx)
		}
	}
	return SCHEDULE_UNKNOWN
}

func (e *Schedule) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Schedule:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Schedule) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Schedule) Ref() *Schedule {
	return &e
}

/*
Contains supported variants of system products.
*/
type SystemType int

const (
	SYSTEMTYPE_UNKNOWN  SystemType = 0
	SYSTEMTYPE_REDACTED SystemType = 1
	SYSTEMTYPE_PC       SystemType = 2
	SYSTEMTYPE_AOS      SystemType = 3
	SYSTEMTYPE_AHV      SystemType = 4
	SYSTEMTYPE_IPMI     SystemType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SystemType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
		"IPMI",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SystemType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
		"IPMI",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SystemType) index(name string) SystemType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"AOS",
		"AHV",
		"IPMI",
	}
	for idx := range names {
		if names[idx] == name {
			return SystemType(idx)
		}
	}
	return SYSTEMTYPE_UNKNOWN
}

func (e *SystemType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SystemType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SystemType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SystemType) Ref() *SystemType {
	return &e
}

/*
Contains information about system user password.
*/
type SystemUserPassword struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Expiry of a new password.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Shows if high strength password is in use or not.
	*/
	HasHspInUse *bool `json:"hasHspInUse,omitempty"`
	/*
	  Timestamp of last password change.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	PasswordStatus *PasswordStatus `json:"passwordStatus,omitempty"`

	SystemType *SystemType `json:"systemType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  System user name.
	*/
	Username *string `json:"username,omitempty"`
}

func NewSystemUserPassword() *SystemUserPassword {
	p := new(SystemUserPassword)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SystemUserPassword"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SystemUserPasswordProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID for the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterProjection *import4.ClusterProjection `json:"clusterProjection,omitempty"`
	/*
	  Expiry of a new password.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Shows if high strength password is in use or not.
	*/
	HasHspInUse *bool `json:"hasHspInUse,omitempty"`
	/*
	  Timestamp of last password change.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	PasswordStatus *PasswordStatus `json:"passwordStatus,omitempty"`

	SystemType *SystemType `json:"systemType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  System user name.
	*/
	Username *string `json:"username,omitempty"`
}

func NewSystemUserPasswordProjection() *SystemUserPasswordProjection {
	p := new(SystemUserPasswordProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SystemUserPasswordProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.0.b1/config/hardening-statuses/{extId} Put operation
*/
type UpdateHardeningStatusesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateHardeningStatusesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateHardeningStatusesApiResponse() *UpdateHardeningStatusesApiResponse {
	p := new(UpdateHardeningStatusesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.UpdateHardeningStatusesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateHardeningStatusesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateHardeningStatusesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateHardeningStatusesApiResponseData()
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
REST response for all response codes in API path /security/v4.0.b1/config/hardening-visibility-setting Put operation
*/
type UpdateHardeningVisibilitySettingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateHardeningVisibilitySettingApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateHardeningVisibilitySettingApiResponse() *UpdateHardeningVisibilitySettingApiResponse {
	p := new(UpdateHardeningVisibilitySettingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.UpdateHardeningVisibilitySettingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateHardeningVisibilitySettingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateHardeningVisibilitySettingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateHardeningVisibilitySettingApiResponseData()
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
Password verification result for an user account.
*/
type VerifyPassword struct {
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
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VerificationStatus *PasswordVerificationStatus `json:"verificationStatus,omitempty"`
}

func NewVerifyPassword() *VerifyPassword {
	p := new(VerifyPassword)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.VerifyPassword"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains all required information for password verification operation.
*/
type VerifyPasswordSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Existing password of a user account.
	*/
	CurrentPassword *string `json:"currentPassword"`
}

func (p *VerifyPasswordSpec) MarshalJSON() ([]byte, error) {
	type VerifyPasswordSpecProxy VerifyPasswordSpec
	return json.Marshal(struct {
		*VerifyPasswordSpecProxy
		CurrentPassword *string `json:"currentPassword,omitempty"`
	}{
		VerifyPasswordSpecProxy: (*VerifyPasswordSpecProxy)(p),
		CurrentPassword:         p.CurrentPassword,
	})
}

func NewVerifyPasswordSpec() *VerifyPasswordSpec {
	p := new(VerifyPasswordSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.VerifyPasswordSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.0.b1/config/system-user-passwords/{extId}/$action/verify-password Post operation
*/
type VerifySystemUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVerifySystemUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVerifySystemUserPasswordApiResponse() *VerifySystemUserPasswordApiResponse {
	p := new(VerifySystemUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.VerifySystemUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VerifySystemUserPasswordApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VerifySystemUserPasswordApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVerifySystemUserPasswordApiResponseData()
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

type OneOfDeletePublicKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeletePublicKeyApiResponseData() *OneOfDeletePublicKeyApiResponseData {
	p := new(OneOfDeletePublicKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeletePublicKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeletePublicKeyApiResponseData is nil"))
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

func (p *OneOfDeletePublicKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeletePublicKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeletePublicKeyApiResponseData"))
}

func (p *OneOfDeletePublicKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeletePublicKeyApiResponseData")
}

type OneOfChangeSystemUserPasswordApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfChangeSystemUserPasswordApiResponseData() *OneOfChangeSystemUserPasswordApiResponseData {
	p := new(OneOfChangeSystemUserPasswordApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfChangeSystemUserPasswordApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfChangeSystemUserPasswordApiResponseData is nil"))
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

func (p *OneOfChangeSystemUserPasswordApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangeSystemUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangeSystemUserPasswordApiResponseData"))
}

func (p *OneOfChangeSystemUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangeSystemUserPasswordApiResponseData")
}

type OneOfUpdateHardeningVisibilitySettingApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []import2.AppMessage   `json:"-"`
}

func NewOneOfUpdateHardeningVisibilitySettingApiResponseData() *OneOfUpdateHardeningVisibilitySettingApiResponseData {
	p := new(OneOfUpdateHardeningVisibilitySettingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateHardeningVisibilitySettingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateHardeningVisibilitySettingApiResponseData is nil"))
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
	case []import2.AppMessage:
		p.oneOfType2001 = v.([]import2.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateHardeningVisibilitySettingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<security.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateHardeningVisibilitySettingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]import2.AppMessage)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "security.v4.error.AppMessage" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateHardeningVisibilitySettingApiResponseData"))
}

func (p *OneOfUpdateHardeningVisibilitySettingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<security.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateHardeningVisibilitySettingApiResponseData")
}

type OneOfAddPublicKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddPublicKeyApiResponseData() *OneOfAddPublicKeyApiResponseData {
	p := new(OneOfAddPublicKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddPublicKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddPublicKeyApiResponseData is nil"))
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

func (p *OneOfAddPublicKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddPublicKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddPublicKeyApiResponseData"))
}

func (p *OneOfAddPublicKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddPublicKeyApiResponseData")
}

type OneOfUpdateHardeningStatusesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateHardeningStatusesApiResponseData() *OneOfUpdateHardeningStatusesApiResponseData {
	p := new(OneOfUpdateHardeningStatusesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateHardeningStatusesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateHardeningStatusesApiResponseData is nil"))
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

func (p *OneOfUpdateHardeningStatusesApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateHardeningStatusesApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateHardeningStatusesApiResponseData"))
}

func (p *OneOfUpdateHardeningStatusesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateHardeningStatusesApiResponseData")
}

type OneOfVerifySystemUserPasswordApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VerifyPassword        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfVerifySystemUserPasswordApiResponseData() *OneOfVerifySystemUserPasswordApiResponseData {
	p := new(OneOfVerifySystemUserPasswordApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVerifySystemUserPasswordApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVerifySystemUserPasswordApiResponseData is nil"))
	}
	switch v.(type) {
	case VerifyPassword:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VerifyPassword)
		}
		*p.oneOfType2001 = v.(VerifyPassword)
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

func (p *OneOfVerifySystemUserPasswordApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVerifySystemUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VerifyPassword)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "security.v4.config.VerifyPassword" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VerifyPassword)
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVerifySystemUserPasswordApiResponseData"))
}

func (p *OneOfVerifySystemUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVerifySystemUserPasswordApiResponseData")
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
