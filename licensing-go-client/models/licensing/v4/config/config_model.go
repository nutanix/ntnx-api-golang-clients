/*
 * Generated file models/licensing/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Licensing APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure licenses.Report allowances and violations
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/prism/v4/config"
	"time"
)

/*
REST response for all response codes in API path /licensing/v4.0/config/license-keys Post operation
*/
type AddLicenseKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddLicenseKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddLicenseKeyApiResponse() *AddLicenseKeyApiResponse {
	p := new(AddLicenseKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AddLicenseKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddLicenseKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddLicenseKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddLicenseKeyApiResponseData()
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

type AddLicenseKeyDryRunApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Status *DryRunStatus `json:"status,omitempty"`

	Summary *LicenseKey `json:"summary,omitempty"`

	ValidationResult *DryRunValidationResult `json:"validationResult,omitempty"`
}

func NewAddLicenseKeyDryRunApiResponse() *AddLicenseKeyDryRunApiResponse {
	p := new(AddLicenseKeyDryRunApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AddLicenseKeyDryRunApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model representing allowance details.
*/
type Allowance struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing a collection of allowance details
	*/
	Details []AllowanceDetail `json:"details,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Indicates feature name like Application Monitoring.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewAllowance() *Allowance {
	p := new(Allowance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Allowance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model representing allowance details.
*/
type AllowanceDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates feature ID like MAX_NODES, APP_MONITORING.
	*/
	FeatureId *string `json:"featureId,omitempty"`

	Scope *Scope `json:"scope,omitempty"`
	/*
	  Attribute for feature value. For boolean, it will be true/false.For integer it will be any numeric value. True value indicates the feature is available to use.
	*/
	Value *string `json:"value,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewAllowanceDetail() *AllowanceDetail {
	p := new(AllowanceDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AllowanceDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AllowanceDetailProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates feature ID like MAX_NODES, APP_MONITORING.
	*/
	FeatureId *string `json:"featureId,omitempty"`

	Scope *Scope `json:"scope,omitempty"`
	/*
	  Attribute for feature value. For boolean, it will be true/false.For integer it will be any numeric value. True value indicates the feature is available to use.
	*/
	Value *string `json:"value,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewAllowanceDetailProjection() *AllowanceDetailProjection {
	p := new(AllowanceDetailProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AllowanceDetailProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AllowanceProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AllowanceDetailProjection []AllowanceDetailProjection `json:"allowanceDetailProjection,omitempty"`
	/*
	  Attribute for capturing a collection of allowance details
	*/
	Details []AllowanceDetail `json:"details,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Indicates feature name like Application Monitoring.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewAllowanceProjection() *AllowanceProjection {
	p := new(AllowanceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AllowanceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /licensing/v4.0/config/$actions/assign-license-keys Post operation
*/
type AssignLicenseKeysApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssignLicenseKeysApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssignLicenseKeysApiResponse() *AssignLicenseKeysApiResponse {
	p := new(AssignLicenseKeysApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AssignLicenseKeysApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssignLicenseKeysApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssignLicenseKeysApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssignLicenseKeysApiResponseData()
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
Model containing cluster identifier and a boolean attribute representing whether it is a PC or PE.
*/
type BaseClusterInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewBaseClusterInfo() *BaseClusterInfo {
	p := new(BaseClusterInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.BaseClusterInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing base licensing attributes like name, type, category and others.
*/
type BaseLicenseInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Indicates the expiration date of the license.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func (p *BaseLicenseInfo) MarshalJSON() ([]byte, error) {
	type BaseLicenseInfoProxy BaseLicenseInfo
	return json.Marshal(struct {
		*BaseLicenseInfoProxy
	}{
		BaseLicenseInfoProxy: (*BaseLicenseInfoProxy)(p),
	})
}

func (p *BaseLicenseInfo) UnmarshalJSON(b []byte) error {
	type CustomBaseLicenseInfo struct {
		ObjectType_    *string                `json:"$objectType,omitempty"`
		Reserved_      map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
		Category       *LicenseCategory       `json:"category,omitempty"`
		ExpiryDate     string                 `json:"expiryDate,omitempty"`
		ExtId          *string                `json:"extId,omitempty"`
		Links          []import2.ApiLink      `json:"links,omitempty"`
		Name           *string                `json:"name,omitempty"`
		SubCategory    *SubCategory           `json:"subCategory,omitempty"`
		TenantId       *string                `json:"tenantId,omitempty"`
		Type           *LicenseType           `json:"type,omitempty"`
	}

	var customBaseLicenseInfo CustomBaseLicenseInfo
	err := json.Unmarshal(b, &customBaseLicenseInfo)
	if err != nil {
		return err
	}

	p.ObjectType_ = customBaseLicenseInfo.ObjectType_
	p.Reserved_ = customBaseLicenseInfo.Reserved_
	p.UnknownFields_ = customBaseLicenseInfo.UnknownFields_
	p.Category = customBaseLicenseInfo.Category
	// Custom date parsing logic for Date field
	if customBaseLicenseInfo.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customBaseLicenseInfo.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct BaseLicenseInfo: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customBaseLicenseInfo.ExtId
	p.Links = customBaseLicenseInfo.Links
	p.Name = customBaseLicenseInfo.Name
	p.SubCategory = customBaseLicenseInfo.SubCategory
	p.TenantId = customBaseLicenseInfo.TenantId
	p.Type = customBaseLicenseInfo.Type

	return nil
}

func NewBaseLicenseInfo() *BaseLicenseInfo {
	p := new(BaseLicenseInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.BaseLicenseInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are PAY_AS_YOU_GO and others.
*/
type BillingPlan int

const (
	BILLINGPLAN_UNKNOWN       BillingPlan = 0
	BILLINGPLAN_REDACTED      BillingPlan = 1
	BILLINGPLAN_FREE_TRIAL    BillingPlan = 2
	BILLINGPLAN_PAY_AS_YOU_GO BillingPlan = 3
	BILLINGPLAN_CLOUD_COMMIT  BillingPlan = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BillingPlan) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FREE_TRIAL",
		"PAY_AS_YOU_GO",
		"CLOUD_COMMIT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BillingPlan) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FREE_TRIAL",
		"PAY_AS_YOU_GO",
		"CLOUD_COMMIT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BillingPlan) index(name string) BillingPlan {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FREE_TRIAL",
		"PAY_AS_YOU_GO",
		"CLOUD_COMMIT",
	}
	for idx := range names {
		if names[idx] == name {
			return BillingPlan(idx)
		}
	}
	return BILLINGPLAN_UNKNOWN
}

func (e *BillingPlan) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BillingPlan:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BillingPlan) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BillingPlan) Ref() *BillingPlan {
	return &e
}

/*
Model for capturing capacity violation. Capacity violations are thrown when a user has not applied licenses to extended capacity.
*/
type CapacityViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates the insufficient quantity of a license. License of given quantity should be applied for a cluster to function properly
	*/
	Shortfall *float64 `json:"shortfall,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func NewCapacityViolation() *CapacityViolation {
	p := new(CapacityViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.CapacityViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing clusters to which license key has been applied.
*/
type ClusterLicenseKeyMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for cluster UUID.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing the license key.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Attribute for capturing quantity used.
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
}

func NewClusterLicenseKeyMapping() *ClusterLicenseKeyMapping {
	p := new(ClusterLicenseKeyMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterLicenseKeyMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterLicenseKeyMappingProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for cluster UUID.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing the license key.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Attribute for capturing quantity used.
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
}

func NewClusterLicenseKeyMappingProjection() *ClusterLicenseKeyMappingProjection {
	p := new(ClusterLicenseKeyMappingProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterLicenseKeyMappingProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are NUTANIX, NON_NUTANIX and others.
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN           ClusterType = 0
	CLUSTERTYPE_REDACTED          ClusterType = 1
	CLUSTERTYPE_NUTANIX           ClusterType = 2
	CLUSTERTYPE_VMWARE            ClusterType = 3
	CLUSTERTYPE_OTHER_NON_NUTANIX ClusterType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"VMWARE",
		"OTHER_NON_NUTANIX",
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
		"NUTANIX",
		"VMWARE",
		"OTHER_NON_NUTANIX",
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
		"NUTANIX",
		"VMWARE",
		"OTHER_NON_NUTANIX",
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
Response object containing cluster and associated compliance details.
*/
type Compliance struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing cluster uuid.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing a collection of service details.
	*/
	Services []Service `json:"services,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewCompliance() *Compliance {
	p := new(Compliance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Compliance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ComplianceProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing cluster uuid.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	ServiceProjection []ServiceProjection `json:"serviceProjection,omitempty"`
	/*
	  Attribute for capturing a collection of service details.
	*/
	Services []Service `json:"services,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewComplianceProjection() *ComplianceProjection {
	p := new(ComplianceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ComplianceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing license consumption details like cluster id and quantity used in that cluster.
*/
type Consumption struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID which is consuming the license.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Indicates quantity of licenses consumed by the cluster.
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
}

func NewConsumption() *Consumption {
	p := new(Consumption)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Consumption"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConsumptionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID which is consuming the license.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Indicates quantity of licenses consumed by the cluster.
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
}

func NewConsumptionProjection() *ConsumptionProjection {
	p := new(ConsumptionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ConsumptionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are HYBRID and others
*/
type ConsumptionType int

const (
	CONSUMPTIONTYPE_UNKNOWN  ConsumptionType = 0
	CONSUMPTIONTYPE_REDACTED ConsumptionType = 1
	CONSUMPTIONTYPE_HYBRID   ConsumptionType = 2
	CONSUMPTIONTYPE_UTILITY  ConsumptionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConsumptionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYBRID",
		"UTILITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConsumptionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYBRID",
		"UTILITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConsumptionType) index(name string) ConsumptionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYBRID",
		"UTILITY",
	}
	for idx := range names {
		if names[idx] == name {
			return ConsumptionType(idx)
		}
	}
	return CONSUMPTIONTYPE_UNKNOWN
}

func (e *ConsumptionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConsumptionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConsumptionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConsumptionType) Ref() *ConsumptionType {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0/config/license-keys/{extId} Delete operation
*/
type DeleteLicenseKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteLicenseKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteLicenseKeyApiResponse() *DeleteLicenseKeyApiResponse {
	p := new(DeleteLicenseKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.DeleteLicenseKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteLicenseKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteLicenseKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteLicenseKeyApiResponseData()
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
Validation Check Object for AddLicenseKeyDryRunApiResponse
*/
type DryRunCheck struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Name *string `json:"name,omitempty"`

	Result *import1.AppMessage `json:"result,omitempty"`
}

func NewDryRunCheck() *DryRunCheck {
	p := new(DryRunCheck)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.DryRunCheck"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The overall status of dry run.
*/
type DryRunStatus int

const (
	DRYRUNSTATUS_UNKNOWN  DryRunStatus = 0
	DRYRUNSTATUS_REDACTED DryRunStatus = 1
	DRYRUNSTATUS_SUCCESS  DryRunStatus = 2
	DRYRUNSTATUS_FAILURE  DryRunStatus = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DryRunStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DryRunStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DryRunStatus) index(name string) DryRunStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCESS",
		"FAILURE",
	}
	for idx := range names {
		if names[idx] == name {
			return DryRunStatus(idx)
		}
	}
	return DRYRUNSTATUS_UNKNOWN
}

func (e *DryRunStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DryRunStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DryRunStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DryRunStatus) Ref() *DryRunStatus {
	return &e
}

/*
Validation Result Object for AddLicenseKeyDryRunApiResponse
*/
type DryRunValidationResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Checks []DryRunCheck `json:"checks,omitempty"`
}

func NewDryRunValidationResult() *DryRunValidationResult {
	p := new(DryRunValidationResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.DryRunValidationResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are NO_NODE_ADDITION, NO_SUPPORT and others.
*/
type EnforcementActions int

const (
	ENFORCEMENTACTIONS_UNKNOWN               EnforcementActions = 0
	ENFORCEMENTACTIONS_REDACTED              EnforcementActions = 1
	ENFORCEMENTACTIONS_NO_NODE_ADDITION      EnforcementActions = 2
	ENFORCEMENTACTIONS_NO_ADVERTISE_CAPACITY EnforcementActions = 3
	ENFORCEMENTACTIONS_NO_SUPPORT            EnforcementActions = 4
	ENFORCEMENTACTIONS_NO_CONTAINER_UPDATE   EnforcementActions = 5
	ENFORCEMENTACTIONS_NO_LOGIN              EnforcementActions = 6
	ENFORCEMENTACTIONS_NO_CLUSTER_PAGE       EnforcementActions = 7
	ENFORCEMENTACTIONS_NO_UPGRADES           EnforcementActions = 8
	ENFORCEMENTACTIONS_NO_SECURITY_PATCH     EnforcementActions = 9
	ENFORCEMENTACTIONS_SHOW_NAGWARE          EnforcementActions = 10
	ENFORCEMENTACTIONS_NO_APP_LAUNCH         EnforcementActions = 11
	ENFORCEMENTACTIONS_NO_RUNBOOK            EnforcementActions = 12
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementActions) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NO_NODE_ADDITION",
		"NO_ADVERTISE_CAPACITY",
		"NO_SUPPORT",
		"NO_CONTAINER_UPDATE",
		"NO_LOGIN",
		"NO_CLUSTER_PAGE",
		"NO_UPGRADES",
		"NO_SECURITY_PATCH",
		"SHOW_NAGWARE",
		"NO_APP_LAUNCH",
		"NO_RUNBOOK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EnforcementActions) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NO_NODE_ADDITION",
		"NO_ADVERTISE_CAPACITY",
		"NO_SUPPORT",
		"NO_CONTAINER_UPDATE",
		"NO_LOGIN",
		"NO_CLUSTER_PAGE",
		"NO_UPGRADES",
		"NO_SECURITY_PATCH",
		"SHOW_NAGWARE",
		"NO_APP_LAUNCH",
		"NO_RUNBOOK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EnforcementActions) index(name string) EnforcementActions {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NO_NODE_ADDITION",
		"NO_ADVERTISE_CAPACITY",
		"NO_SUPPORT",
		"NO_CONTAINER_UPDATE",
		"NO_LOGIN",
		"NO_CLUSTER_PAGE",
		"NO_UPGRADES",
		"NO_SECURITY_PATCH",
		"SHOW_NAGWARE",
		"NO_APP_LAUNCH",
		"NO_RUNBOOK",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementActions(idx)
		}
	}
	return ENFORCEMENTACTIONS_UNKNOWN
}

func (e *EnforcementActions) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementActions:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementActions) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementActions) Ref() *EnforcementActions {
	return &e
}

/*
Possible values are LEVEL_10 and LEVEL_20.
*/
type EnforcementLevel int

const (
	ENFORCEMENTLEVEL_UNKNOWN  EnforcementLevel = 0
	ENFORCEMENTLEVEL_REDACTED EnforcementLevel = 1
	ENFORCEMENTLEVEL_LEVEL_10 EnforcementLevel = 2
	ENFORCEMENTLEVEL_LEVEL_20 EnforcementLevel = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEVEL_10",
		"LEVEL_20",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EnforcementLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEVEL_10",
		"LEVEL_20",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EnforcementLevel) index(name string) EnforcementLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEVEL_10",
		"LEVEL_20",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementLevel(idx)
		}
	}
	return ENFORCEMENTLEVEL_UNKNOWN
}

func (e *EnforcementLevel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementLevel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementLevel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementLevel) Ref() *EnforcementLevel {
	return &e
}

/*
Possible values are All or None.
*/
type EnforcementPolicy int

const (
	ENFORCEMENTPOLICY_UNKNOWN  EnforcementPolicy = 0
	ENFORCEMENTPOLICY_REDACTED EnforcementPolicy = 1
	ENFORCEMENTPOLICY_ALL      EnforcementPolicy = 2
	ENFORCEMENTPOLICY_NONE     EnforcementPolicy = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementPolicy) name(index int) string {
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
func (e EnforcementPolicy) GetName() string {
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
func (e *EnforcementPolicy) index(name string) EnforcementPolicy {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementPolicy(idx)
		}
	}
	return ENFORCEMENTPOLICY_UNKNOWN
}

func (e *EnforcementPolicy) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementPolicy:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementPolicy) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementPolicy) Ref() *EnforcementPolicy {
	return &e
}

/*
Model capturing entitlement details.
*/
type Entitlement struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing cluster UUID for entitlement.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing a collection of entitlement details.
	*/
	Details []EntitlementDetail `json:"details,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  Indicates whether cluster is registered with PC or not.
	*/
	IsRegistered *bool `json:"isRegistered,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Indicates the name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewEntitlement() *Entitlement {
	p := new(Entitlement)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Entitlement"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model capturing entitlement details.
*/
type EntitlementDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Attribute for capturing earliest expiry date across entitlements.
	*/
	EarliestExpiryDate *time.Time `json:"earliestExpiryDate,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates the name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Attribute for capturing quantity for entitlement.
	*/
	Quantity *float64 `json:"quantity,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func (p *EntitlementDetail) MarshalJSON() ([]byte, error) {
	type EntitlementDetailProxy EntitlementDetail
	return json.Marshal(struct {
		*EntitlementDetailProxy
	}{
		EntitlementDetailProxy: (*EntitlementDetailProxy)(p),
	})
}

func (p *EntitlementDetail) UnmarshalJSON(b []byte) error {
	type CustomEntitlementDetail struct {
		ObjectType_        *string                `json:"$objectType,omitempty"`
		Reserved_          map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_     map[string]interface{} `json:"$unknownFields,omitempty"`
		Category           *LicenseCategory       `json:"category,omitempty"`
		EarliestExpiryDate string                 `json:"earliestExpiryDate,omitempty"`
		Meter              *Meter                 `json:"meter,omitempty"`
		Name               *string                `json:"name,omitempty"`
		Quantity           *float64               `json:"quantity,omitempty"`
		Scope              *Scope                 `json:"scope,omitempty"`
		SubCategory        *SubCategory           `json:"subCategory,omitempty"`
		Type               *LicenseType           `json:"type,omitempty"`
	}

	var customEntitlementDetail CustomEntitlementDetail
	err := json.Unmarshal(b, &customEntitlementDetail)
	if err != nil {
		return err
	}

	p.ObjectType_ = customEntitlementDetail.ObjectType_
	p.Reserved_ = customEntitlementDetail.Reserved_
	p.UnknownFields_ = customEntitlementDetail.UnknownFields_
	p.Category = customEntitlementDetail.Category
	// Custom date parsing logic for Date field
	if customEntitlementDetail.EarliestExpiryDate != "" {
		parsedEarliestExpiryDate, err := time.Parse("2006-01-02", customEntitlementDetail.EarliestExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field EarliestExpiryDate in struct EntitlementDetail: %s", err))
		}
		p.EarliestExpiryDate = &parsedEarliestExpiryDate
	}
	p.Meter = customEntitlementDetail.Meter
	p.Name = customEntitlementDetail.Name
	p.Quantity = customEntitlementDetail.Quantity
	p.Scope = customEntitlementDetail.Scope
	p.SubCategory = customEntitlementDetail.SubCategory
	p.Type = customEntitlementDetail.Type

	return nil
}

func NewEntitlementDetail() *EntitlementDetail {
	p := new(EntitlementDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.EntitlementDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EntitlementDetailProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Attribute for capturing earliest expiry date across entitlements.
	*/
	EarliestExpiryDate *time.Time `json:"earliestExpiryDate,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates the name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Attribute for capturing quantity for entitlement.
	*/
	Quantity *float64 `json:"quantity,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func (p *EntitlementDetailProjection) MarshalJSON() ([]byte, error) {
	type EntitlementDetailProjectionProxy EntitlementDetailProjection
	return json.Marshal(struct {
		*EntitlementDetailProjectionProxy
	}{
		EntitlementDetailProjectionProxy: (*EntitlementDetailProjectionProxy)(p),
	})
}

func (p *EntitlementDetailProjection) UnmarshalJSON(b []byte) error {
	type CustomEntitlementDetailProjection struct {
		ObjectType_        *string                `json:"$objectType,omitempty"`
		Reserved_          map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_     map[string]interface{} `json:"$unknownFields,omitempty"`
		Category           *LicenseCategory       `json:"category,omitempty"`
		EarliestExpiryDate string                 `json:"earliestExpiryDate,omitempty"`
		Meter              *Meter                 `json:"meter,omitempty"`
		Name               *string                `json:"name,omitempty"`
		Quantity           *float64               `json:"quantity,omitempty"`
		Scope              *Scope                 `json:"scope,omitempty"`
		SubCategory        *SubCategory           `json:"subCategory,omitempty"`
		Type               *LicenseType           `json:"type,omitempty"`
	}

	var customEntitlementDetailProjection CustomEntitlementDetailProjection
	err := json.Unmarshal(b, &customEntitlementDetailProjection)
	if err != nil {
		return err
	}

	p.ObjectType_ = customEntitlementDetailProjection.ObjectType_
	p.Reserved_ = customEntitlementDetailProjection.Reserved_
	p.UnknownFields_ = customEntitlementDetailProjection.UnknownFields_
	p.Category = customEntitlementDetailProjection.Category
	// Custom date parsing logic for Date field
	if customEntitlementDetailProjection.EarliestExpiryDate != "" {
		parsedEarliestExpiryDate, err := time.Parse("2006-01-02", customEntitlementDetailProjection.EarliestExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field EarliestExpiryDate in struct EntitlementDetailProjection: %s", err))
		}
		p.EarliestExpiryDate = &parsedEarliestExpiryDate
	}
	p.Meter = customEntitlementDetailProjection.Meter
	p.Name = customEntitlementDetailProjection.Name
	p.Quantity = customEntitlementDetailProjection.Quantity
	p.Scope = customEntitlementDetailProjection.Scope
	p.SubCategory = customEntitlementDetailProjection.SubCategory
	p.Type = customEntitlementDetailProjection.Type

	return nil
}

func NewEntitlementDetailProjection() *EntitlementDetailProjection {
	p := new(EntitlementDetailProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.EntitlementDetailProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EntitlementProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing cluster UUID for entitlement.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing a collection of entitlement details.
	*/
	Details []EntitlementDetail `json:"details,omitempty"`

	EntitlementDetailProjection []EntitlementDetailProjection `json:"entitlementDetailProjection,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  Indicates whether cluster is registered with PC or not.
	*/
	IsRegistered *bool `json:"isRegistered,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Indicates the name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterType `json:"type,omitempty"`
}

func NewEntitlementProjection() *EntitlementProjection {
	p := new(EntitlementProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.EntitlementProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
License details of expired license.
*/
type ExpiredLicense struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Indicates the expiration date of the license.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  License ID of expired license.
	*/
	LicenseId *string `json:"licenseId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
	/*
	  Indicates the expired quantity of the license.
	*/
	UsedQuantity *float64 `json:"usedQuantity,omitempty"`
}

func (p *ExpiredLicense) MarshalJSON() ([]byte, error) {
	type ExpiredLicenseProxy ExpiredLicense
	return json.Marshal(struct {
		*ExpiredLicenseProxy
	}{
		ExpiredLicenseProxy: (*ExpiredLicenseProxy)(p),
	})
}

func (p *ExpiredLicense) UnmarshalJSON(b []byte) error {
	type CustomExpiredLicense struct {
		ObjectType_    *string                `json:"$objectType,omitempty"`
		Reserved_      map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
		Category       *LicenseCategory       `json:"category,omitempty"`
		ExpiryDate     string                 `json:"expiryDate,omitempty"`
		ExtId          *string                `json:"extId,omitempty"`
		LicenseId      *string                `json:"licenseId,omitempty"`
		Links          []import2.ApiLink      `json:"links,omitempty"`
		Meter          *Meter                 `json:"meter,omitempty"`
		Name           *string                `json:"name,omitempty"`
		SubCategory    *SubCategory           `json:"subCategory,omitempty"`
		TenantId       *string                `json:"tenantId,omitempty"`
		Type           *LicenseType           `json:"type,omitempty"`
		UsedQuantity   *float64               `json:"usedQuantity,omitempty"`
	}

	var customExpiredLicense CustomExpiredLicense
	err := json.Unmarshal(b, &customExpiredLicense)
	if err != nil {
		return err
	}

	p.ObjectType_ = customExpiredLicense.ObjectType_
	p.Reserved_ = customExpiredLicense.Reserved_
	p.UnknownFields_ = customExpiredLicense.UnknownFields_
	p.Category = customExpiredLicense.Category
	// Custom date parsing logic for Date field
	if customExpiredLicense.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customExpiredLicense.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct ExpiredLicense: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customExpiredLicense.ExtId
	p.LicenseId = customExpiredLicense.LicenseId
	p.Links = customExpiredLicense.Links
	p.Meter = customExpiredLicense.Meter
	p.Name = customExpiredLicense.Name
	p.SubCategory = customExpiredLicense.SubCategory
	p.TenantId = customExpiredLicense.TenantId
	p.Type = customExpiredLicense.Type
	p.UsedQuantity = customExpiredLicense.UsedQuantity

	return nil
}

func NewExpiredLicense() *ExpiredLicense {
	p := new(ExpiredLicense)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ExpiredLicense"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response object containing feature details like name, description, license type and category mapping and others.
*/
type Feature struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	LicenseCategory *LicenseCategory `json:"licenseCategory,omitempty"`

	LicenseSubCategory *SubCategory `json:"licenseSubCategory,omitempty"`

	LicenseType *LicenseType `json:"licenseType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of feature like dp_recovery, dp_backup_tiering.
	*/
	Name *string `json:"name,omitempty"`

	Scope *Scope `json:"scope,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*

	 */
	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  Value of feature, it could be true, false or integer.
	*/
	Value *OneOfFeatureValue `json:"value,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewFeature() *Feature {
	p := new(Feature)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Feature"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Feature) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *Feature) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfFeatureValue()
	}
	e := p.Value.SetValue(v)
	if nil == e {
		if nil == p.ValueItemDiscriminator_ {
			p.ValueItemDiscriminator_ = new(string)
		}
		*p.ValueItemDiscriminator_ = *p.Value.Discriminator
	}
	return e
}

/*
Model for capturing feature details.
*/
type FeatureDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for feature description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Attribute for capturing feature id like VULCAN, APP_MONITORING.
	*/
	FeatureId *string `json:"featureId,omitempty"`
	/*
	  Attribute for capturing feature name like Application Monitoring.
	*/
	Name *string `json:"name,omitempty"`
}

func NewFeatureDetail() *FeatureDetail {
	p := new(FeatureDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type FeatureProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	LicenseCategory *LicenseCategory `json:"licenseCategory,omitempty"`

	LicenseSubCategory *SubCategory `json:"licenseSubCategory,omitempty"`

	LicenseType *LicenseType `json:"licenseType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of feature like dp_recovery, dp_backup_tiering.
	*/
	Name *string `json:"name,omitempty"`

	Scope *Scope `json:"scope,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  Value of feature, it could be true, false or integer.
	*/
	Value *OneOfFeatureProjectionValue `json:"value,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewFeatureProjection() *FeatureProjection {
	p := new(FeatureProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model capturing feature info and associated clusters.
*/
type FeatureViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of entity for which feature violation is thrown.
	Examples: 1) Vm with uuid 000604af-3aa4-9cfe-1c6a-ac1f6b357fb7.
	          2) Storage Container bucket-data-prod
	          3) Protection Domain pd_prod with Application consistency group prod.
	*/
	AffectedEntity *string `json:"affectedEntity,omitempty"`
	/*
	  Attribute for feature description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Attribute for capturing feature id like VULCAN, APP_MONITORING.
	*/
	FeatureId *string `json:"featureId,omitempty"`
	/*
	  Attribute for capturing feature name like Application Monitoring.
	*/
	Name *string `json:"name,omitempty"`
}

func NewFeatureViolation() *FeatureViolation {
	p := new(FeatureViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /licensing/v4.0/config/license-keys/{extId} Get operation
*/
type GetLicenseKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLicenseKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLicenseKeyApiResponse() *GetLicenseKeyApiResponse {
	p := new(GetLicenseKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetLicenseKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLicenseKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLicenseKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLicenseKeyApiResponseData()
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
Model containing license details like id, type, category, sub-category, scope and others.
*/
type License struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Array containing information about the clusters where these licenses are used.
	*/
	ConsumptionDetails []Consumption `json:"consumptionDetails,omitempty"`
	/*
	  Indicates the expiration date of the license.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Total quantity of license consumed.
	*/
	Quantity *float64 `json:"quantity,omitempty"`
	/*
	  Attribute for capturing salesforce generated license identifier e.g. LIC-XXX.
	*/
	SalesforceLicenseId *string `json:"salesforceLicenseId,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func (p *License) MarshalJSON() ([]byte, error) {
	type LicenseProxy License
	return json.Marshal(struct {
		*LicenseProxy
	}{
		LicenseProxy: (*LicenseProxy)(p),
	})
}

func (p *License) UnmarshalJSON(b []byte) error {
	type CustomLicense struct {
		ObjectType_         *string                `json:"$objectType,omitempty"`
		Reserved_           map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_      map[string]interface{} `json:"$unknownFields,omitempty"`
		Category            *LicenseCategory       `json:"category,omitempty"`
		ConsumptionDetails  []Consumption          `json:"consumptionDetails,omitempty"`
		ExpiryDate          string                 `json:"expiryDate,omitempty"`
		ExtId               *string                `json:"extId,omitempty"`
		Links               []import2.ApiLink      `json:"links,omitempty"`
		Meter               *Meter                 `json:"meter,omitempty"`
		Name                *string                `json:"name,omitempty"`
		Quantity            *float64               `json:"quantity,omitempty"`
		SalesforceLicenseId *string                `json:"salesforceLicenseId,omitempty"`
		Scope               *Scope                 `json:"scope,omitempty"`
		SubCategory         *SubCategory           `json:"subCategory,omitempty"`
		TenantId            *string                `json:"tenantId,omitempty"`
		Type                *LicenseType           `json:"type,omitempty"`
	}

	var customLicense CustomLicense
	err := json.Unmarshal(b, &customLicense)
	if err != nil {
		return err
	}

	p.ObjectType_ = customLicense.ObjectType_
	p.Reserved_ = customLicense.Reserved_
	p.UnknownFields_ = customLicense.UnknownFields_
	p.Category = customLicense.Category
	p.ConsumptionDetails = customLicense.ConsumptionDetails
	// Custom date parsing logic for Date field
	if customLicense.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customLicense.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct License: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customLicense.ExtId
	p.Links = customLicense.Links
	p.Meter = customLicense.Meter
	p.Name = customLicense.Name
	p.Quantity = customLicense.Quantity
	p.SalesforceLicenseId = customLicense.SalesforceLicenseId
	p.Scope = customLicense.Scope
	p.SubCategory = customLicense.SubCategory
	p.TenantId = customLicense.TenantId
	p.Type = customLicense.Type

	return nil
}

func NewLicense() *License {
	p := new(License)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.License"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the different categories of license with possible values of Starter, Pro, Ultimate, and others.
*/
type LicenseCategory int

const (
	LICENSECATEGORY_UNKNOWN              LicenseCategory = 0
	LICENSECATEGORY_REDACTED             LicenseCategory = 1
	LICENSECATEGORY_STARTER              LicenseCategory = 2
	LICENSECATEGORY_PRO                  LicenseCategory = 3
	LICENSECATEGORY_ULTIMATE             LicenseCategory = 4
	LICENSECATEGORY_CALM                 LicenseCategory = 5
	LICENSECATEGORY_STANDARD             LicenseCategory = 6
	LICENSECATEGORY_AOS_MINE             LicenseCategory = 7
	LICENSECATEGORY_SOFTWARE_ENCRYPTION  LicenseCategory = 8
	LICENSECATEGORY_ADV_REPLICATION      LicenseCategory = 9
	LICENSECATEGORY_OBJECT               LicenseCategory = 10
	LICENSECATEGORY_ULTIMATE_TRIAL       LicenseCategory = 11
	LICENSECATEGORY_PRISM_STARTER        LicenseCategory = 12
	LICENSECATEGORY_PRO_SPECIAL          LicenseCategory = 13
	LICENSECATEGORY_ADR                  LicenseCategory = 14
	LICENSECATEGORY_SECURITY             LicenseCategory = 15
	LICENSECATEGORY_NKS                  LicenseCategory = 16
	LICENSECATEGORY_APPAUTOMATION        LicenseCategory = 17
	LICENSECATEGORY_NDA                  LicenseCategory = 18
	LICENSECATEGORY_UST                  LicenseCategory = 19
	LICENSECATEGORY_ANALYTICS            LicenseCategory = 20
	LICENSECATEGORY_STANDALONE           LicenseCategory = 21
	LICENSECATEGORY_DRASS                LicenseCategory = 22
	LICENSECATEGORY_CLOUD_NATIVE         LicenseCategory = 23
	LICENSECATEGORY_DATA_ENCRYPTION      LicenseCategory = 24
	LICENSECATEGORY_NDS                  LicenseCategory = 25
	LICENSECATEGORY_NDB                  LicenseCategory = 26
	LICENSECATEGORY_NO_LICENSE           LicenseCategory = 27
	LICENSECATEGORY_NUS_ENCRYPTION       LicenseCategory = 28
	LICENSECATEGORY_NUS_REPLICATION      LicenseCategory = 29
	LICENSECATEGORY_CLOUD_PRO            LicenseCategory = 30
	LICENSECATEGORY_CLOUD_ULTIMATE       LicenseCategory = 31
	LICENSECATEGORY_CLOUD                LicenseCategory = 32
	LICENSECATEGORY_PUBLIC_CLOUD         LicenseCategory = 33
	LICENSECATEGORY_ADVANCED_REPLICATION LicenseCategory = 34
	LICENSECATEGORY_NDK                  LicenseCategory = 35
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LicenseCategory) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"PRO",
		"ULTIMATE",
		"CALM",
		"STANDARD",
		"AOS_MINE",
		"SOFTWARE_ENCRYPTION",
		"ADV_REPLICATION",
		"OBJECT",
		"ULTIMATE_TRIAL",
		"PRISM_STARTER",
		"PRO_SPECIAL",
		"ADR",
		"SECURITY",
		"NKS",
		"APPAUTOMATION",
		"NDA",
		"UST",
		"ANALYTICS",
		"STANDALONE",
		"DRASS",
		"CLOUD_NATIVE",
		"DATA_ENCRYPTION",
		"NDS",
		"NDB",
		"NO_LICENSE",
		"NUS_ENCRYPTION",
		"NUS_REPLICATION",
		"CLOUD_PRO",
		"CLOUD_ULTIMATE",
		"CLOUD",
		"PUBLIC_CLOUD",
		"ADVANCED_REPLICATION",
		"NDK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LicenseCategory) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"PRO",
		"ULTIMATE",
		"CALM",
		"STANDARD",
		"AOS_MINE",
		"SOFTWARE_ENCRYPTION",
		"ADV_REPLICATION",
		"OBJECT",
		"ULTIMATE_TRIAL",
		"PRISM_STARTER",
		"PRO_SPECIAL",
		"ADR",
		"SECURITY",
		"NKS",
		"APPAUTOMATION",
		"NDA",
		"UST",
		"ANALYTICS",
		"STANDALONE",
		"DRASS",
		"CLOUD_NATIVE",
		"DATA_ENCRYPTION",
		"NDS",
		"NDB",
		"NO_LICENSE",
		"NUS_ENCRYPTION",
		"NUS_REPLICATION",
		"CLOUD_PRO",
		"CLOUD_ULTIMATE",
		"CLOUD",
		"PUBLIC_CLOUD",
		"ADVANCED_REPLICATION",
		"NDK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LicenseCategory) index(name string) LicenseCategory {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"PRO",
		"ULTIMATE",
		"CALM",
		"STANDARD",
		"AOS_MINE",
		"SOFTWARE_ENCRYPTION",
		"ADV_REPLICATION",
		"OBJECT",
		"ULTIMATE_TRIAL",
		"PRISM_STARTER",
		"PRO_SPECIAL",
		"ADR",
		"SECURITY",
		"NKS",
		"APPAUTOMATION",
		"NDA",
		"UST",
		"ANALYTICS",
		"STANDALONE",
		"DRASS",
		"CLOUD_NATIVE",
		"DATA_ENCRYPTION",
		"NDS",
		"NDB",
		"NO_LICENSE",
		"NUS_ENCRYPTION",
		"NUS_REPLICATION",
		"CLOUD_PRO",
		"CLOUD_ULTIMATE",
		"CLOUD",
		"PUBLIC_CLOUD",
		"ADVANCED_REPLICATION",
		"NDK",
	}
	for idx := range names {
		if names[idx] == name {
			return LicenseCategory(idx)
		}
	}
	return LICENSECATEGORY_UNKNOWN
}

func (e *LicenseCategory) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LicenseCategory:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LicenseCategory) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LicenseCategory) Ref() *LicenseCategory {
	return &e
}

/*
Possible values are APPLIANCE, SOFTWARE_ONLY and others.
*/
type LicenseClass int

const (
	LICENSECLASS_UNKNOWN       LicenseClass = 0
	LICENSECLASS_REDACTED      LicenseClass = 1
	LICENSECLASS_PRISM_CENTRAL LicenseClass = 2
	LICENSECLASS_APPLIANCE     LicenseClass = 3
	LICENSECLASS_SOFTWARE_ONLY LicenseClass = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LicenseClass) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"APPLIANCE",
		"SOFTWARE_ONLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LicenseClass) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"APPLIANCE",
		"SOFTWARE_ONLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LicenseClass) index(name string) LicenseClass {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_CENTRAL",
		"APPLIANCE",
		"SOFTWARE_ONLY",
	}
	for idx := range names {
		if names[idx] == name {
			return LicenseClass(idx)
		}
	}
	return LICENSECLASS_UNKNOWN
}

func (e *LicenseClass) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LicenseClass:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LicenseClass) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LicenseClass) Ref() *LicenseClass {
	return &e
}

/*
Model for capturing license key detail attributes.
*/
type LicenseKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing add-on categories associated with this license key.
	*/
	Addons []LicenseCategory `json:"addons,omitempty"`
	/*
	  Expansion attribute used for fetching the license key assignment details.
	*/
	AssignmentDetails []ClusterLicenseKeyMapping `json:"assignmentDetails,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Attribute for capturing license key expiration date.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Attribute for capturing license key.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Attribute for capturing license key quantity.
	*/
	Quantity *float64 `json:"quantity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
	/*
	  Attribute for capturing validation data.
	*/
	ValidationDetail *string `json:"validationDetail,omitempty"`
}

func (p *LicenseKey) MarshalJSON() ([]byte, error) {
	type LicenseKeyProxy LicenseKey
	return json.Marshal(struct {
		*LicenseKeyProxy
	}{
		LicenseKeyProxy: (*LicenseKeyProxy)(p),
	})
}

func (p *LicenseKey) UnmarshalJSON(b []byte) error {
	type CustomLicenseKey struct {
		ObjectType_       *string                    `json:"$objectType,omitempty"`
		Reserved_         map[string]interface{}     `json:"$reserved,omitempty"`
		UnknownFields_    map[string]interface{}     `json:"$unknownFields,omitempty"`
		Addons            []LicenseCategory          `json:"addons,omitempty"`
		AssignmentDetails []ClusterLicenseKeyMapping `json:"assignmentDetails,omitempty"`
		Category          *LicenseCategory           `json:"category,omitempty"`
		ExpiryDate        string                     `json:"expiryDate,omitempty"`
		ExtId             *string                    `json:"extId,omitempty"`
		Key               *string                    `json:"key,omitempty"`
		Links             []import2.ApiLink          `json:"links,omitempty"`
		Meter             *Meter                     `json:"meter,omitempty"`
		Quantity          *float64                   `json:"quantity,omitempty"`
		TenantId          *string                    `json:"tenantId,omitempty"`
		Type              *LicenseType               `json:"type,omitempty"`
		ValidationDetail  *string                    `json:"validationDetail,omitempty"`
	}

	var customLicenseKey CustomLicenseKey
	err := json.Unmarshal(b, &customLicenseKey)
	if err != nil {
		return err
	}

	p.ObjectType_ = customLicenseKey.ObjectType_
	p.Reserved_ = customLicenseKey.Reserved_
	p.UnknownFields_ = customLicenseKey.UnknownFields_
	p.Addons = customLicenseKey.Addons
	p.AssignmentDetails = customLicenseKey.AssignmentDetails
	p.Category = customLicenseKey.Category
	// Custom date parsing logic for Date field
	if customLicenseKey.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customLicenseKey.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct LicenseKey: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customLicenseKey.ExtId
	p.Key = customLicenseKey.Key
	p.Links = customLicenseKey.Links
	p.Meter = customLicenseKey.Meter
	p.Quantity = customLicenseKey.Quantity
	p.TenantId = customLicenseKey.TenantId
	p.Type = customLicenseKey.Type
	p.ValidationDetail = customLicenseKey.ValidationDetail

	return nil
}

func NewLicenseKey() *LicenseKey {
	p := new(LicenseKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing license key cluster mapping.
*/
type LicenseKeyAssignSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing the cluster external identifier.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing array of license key mappings.
	*/
	LicenseKeyMappings []LicenseKeyMapping `json:"licenseKeyMappings,omitempty"`
}

func NewLicenseKeyAssignSpec() *LicenseKeyAssignSpec {
	p := new(LicenseKeyAssignSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseKeyAssignSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing license key and quantity mapping.
*/
type LicenseKeyMapping struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing the license key.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Attribute for capturing quantity used.
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
}

func NewLicenseKeyMapping() *LicenseKeyMapping {
	p := new(LicenseKeyMapping)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseKeyMapping"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type LicenseKeyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing add-on categories associated with this license key.
	*/
	Addons []LicenseCategory `json:"addons,omitempty"`
	/*
	  Expansion attribute used for fetching the license key assignment details.
	*/
	AssignmentDetails []ClusterLicenseKeyMapping `json:"assignmentDetails,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`

	ClusterLicenseKeyMappingProjection []ClusterLicenseKeyMappingProjection `json:"clusterLicenseKeyMappingProjection,omitempty"`
	/*
	  Attribute for capturing license key expiration date.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Attribute for capturing license key.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Attribute for capturing license key quantity.
	*/
	Quantity *float64 `json:"quantity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
	/*
	  Attribute for capturing validation data.
	*/
	ValidationDetail *string `json:"validationDetail,omitempty"`
}

func (p *LicenseKeyProjection) MarshalJSON() ([]byte, error) {
	type LicenseKeyProjectionProxy LicenseKeyProjection
	return json.Marshal(struct {
		*LicenseKeyProjectionProxy
	}{
		LicenseKeyProjectionProxy: (*LicenseKeyProjectionProxy)(p),
	})
}

func (p *LicenseKeyProjection) UnmarshalJSON(b []byte) error {
	type CustomLicenseKeyProjection struct {
		ObjectType_                        *string                              `json:"$objectType,omitempty"`
		Reserved_                          map[string]interface{}               `json:"$reserved,omitempty"`
		UnknownFields_                     map[string]interface{}               `json:"$unknownFields,omitempty"`
		Addons                             []LicenseCategory                    `json:"addons,omitempty"`
		AssignmentDetails                  []ClusterLicenseKeyMapping           `json:"assignmentDetails,omitempty"`
		Category                           *LicenseCategory                     `json:"category,omitempty"`
		ClusterLicenseKeyMappingProjection []ClusterLicenseKeyMappingProjection `json:"clusterLicenseKeyMappingProjection,omitempty"`
		ExpiryDate                         string                               `json:"expiryDate,omitempty"`
		ExtId                              *string                              `json:"extId,omitempty"`
		Key                                *string                              `json:"key,omitempty"`
		Links                              []import2.ApiLink                    `json:"links,omitempty"`
		Meter                              *Meter                               `json:"meter,omitempty"`
		Quantity                           *float64                             `json:"quantity,omitempty"`
		TenantId                           *string                              `json:"tenantId,omitempty"`
		Type                               *LicenseType                         `json:"type,omitempty"`
		ValidationDetail                   *string                              `json:"validationDetail,omitempty"`
	}

	var customLicenseKeyProjection CustomLicenseKeyProjection
	err := json.Unmarshal(b, &customLicenseKeyProjection)
	if err != nil {
		return err
	}

	p.ObjectType_ = customLicenseKeyProjection.ObjectType_
	p.Reserved_ = customLicenseKeyProjection.Reserved_
	p.UnknownFields_ = customLicenseKeyProjection.UnknownFields_
	p.Addons = customLicenseKeyProjection.Addons
	p.AssignmentDetails = customLicenseKeyProjection.AssignmentDetails
	p.Category = customLicenseKeyProjection.Category
	p.ClusterLicenseKeyMappingProjection = customLicenseKeyProjection.ClusterLicenseKeyMappingProjection
	// Custom date parsing logic for Date field
	if customLicenseKeyProjection.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customLicenseKeyProjection.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct LicenseKeyProjection: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customLicenseKeyProjection.ExtId
	p.Key = customLicenseKeyProjection.Key
	p.Links = customLicenseKeyProjection.Links
	p.Meter = customLicenseKeyProjection.Meter
	p.Quantity = customLicenseKeyProjection.Quantity
	p.TenantId = customLicenseKeyProjection.TenantId
	p.Type = customLicenseKeyProjection.Type
	p.ValidationDetail = customLicenseKeyProjection.ValidationDetail

	return nil
}

func NewLicenseKeyProjection() *LicenseKeyProjection {
	p := new(LicenseKeyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseKeyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type LicenseProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategory `json:"category,omitempty"`
	/*
	  Array containing information about the clusters where these licenses are used.
	*/
	ConsumptionDetails []Consumption `json:"consumptionDetails,omitempty"`

	ConsumptionProjection []ConsumptionProjection `json:"consumptionProjection,omitempty"`
	/*
	  Indicates the expiration date of the license.
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Meter *Meter `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Total quantity of license consumed.
	*/
	Quantity *float64 `json:"quantity,omitempty"`
	/*
	  Attribute for capturing salesforce generated license identifier e.g. LIC-XXX.
	*/
	SalesforceLicenseId *string `json:"salesforceLicenseId,omitempty"`

	Scope *Scope `json:"scope,omitempty"`

	SubCategory *SubCategory `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseType `json:"type,omitempty"`
}

func (p *LicenseProjection) MarshalJSON() ([]byte, error) {
	type LicenseProjectionProxy LicenseProjection
	return json.Marshal(struct {
		*LicenseProjectionProxy
	}{
		LicenseProjectionProxy: (*LicenseProjectionProxy)(p),
	})
}

func (p *LicenseProjection) UnmarshalJSON(b []byte) error {
	type CustomLicenseProjection struct {
		ObjectType_           *string                 `json:"$objectType,omitempty"`
		Reserved_             map[string]interface{}  `json:"$reserved,omitempty"`
		UnknownFields_        map[string]interface{}  `json:"$unknownFields,omitempty"`
		Category              *LicenseCategory        `json:"category,omitempty"`
		ConsumptionDetails    []Consumption           `json:"consumptionDetails,omitempty"`
		ConsumptionProjection []ConsumptionProjection `json:"consumptionProjection,omitempty"`
		ExpiryDate            string                  `json:"expiryDate,omitempty"`
		ExtId                 *string                 `json:"extId,omitempty"`
		Links                 []import2.ApiLink       `json:"links,omitempty"`
		Meter                 *Meter                  `json:"meter,omitempty"`
		Name                  *string                 `json:"name,omitempty"`
		Quantity              *float64                `json:"quantity,omitempty"`
		SalesforceLicenseId   *string                 `json:"salesforceLicenseId,omitempty"`
		Scope                 *Scope                  `json:"scope,omitempty"`
		SubCategory           *SubCategory            `json:"subCategory,omitempty"`
		TenantId              *string                 `json:"tenantId,omitempty"`
		Type                  *LicenseType            `json:"type,omitempty"`
	}

	var customLicenseProjection CustomLicenseProjection
	err := json.Unmarshal(b, &customLicenseProjection)
	if err != nil {
		return err
	}

	p.ObjectType_ = customLicenseProjection.ObjectType_
	p.Reserved_ = customLicenseProjection.Reserved_
	p.UnknownFields_ = customLicenseProjection.UnknownFields_
	p.Category = customLicenseProjection.Category
	p.ConsumptionDetails = customLicenseProjection.ConsumptionDetails
	p.ConsumptionProjection = customLicenseProjection.ConsumptionProjection
	// Custom date parsing logic for Date field
	if customLicenseProjection.ExpiryDate != "" {
		parsedExpiryDate, err := time.Parse("2006-01-02", customLicenseProjection.ExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ExpiryDate in struct LicenseProjection: %s", err))
		}
		p.ExpiryDate = &parsedExpiryDate
	}
	p.ExtId = customLicenseProjection.ExtId
	p.Links = customLicenseProjection.Links
	p.Meter = customLicenseProjection.Meter
	p.Name = customLicenseProjection.Name
	p.Quantity = customLicenseProjection.Quantity
	p.SalesforceLicenseId = customLicenseProjection.SalesforceLicenseId
	p.Scope = customLicenseProjection.Scope
	p.SubCategory = customLicenseProjection.SubCategory
	p.TenantId = customLicenseProjection.TenantId
	p.Type = customLicenseProjection.Type

	return nil
}

func NewLicenseProjection() *LicenseProjection {
	p := new(LicenseProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing seamless licensing action attributes.
*/
type LicenseStateSyncSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing affected cluster IDs.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  Attribute for capturing the entitlement names like NCM_PRO, AOS_PRO and others.
	*/
	EntitlementNames []string `json:"entitlementNames,omitempty"`

	Operation *SyncOperationType `json:"operation,omitempty"`
}

func NewLicenseStateSyncSpec() *LicenseStateSyncSpec {
	p := new(LicenseStateSyncSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseStateSyncSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the different license types with possible values of Prism, AOS, Calm, and others.
*/
type LicenseType int

const (
	LICENSETYPE_UNKNOWN         LicenseType = 0
	LICENSETYPE_REDACTED        LicenseType = 1
	LICENSETYPE_PRISM           LicenseType = 2
	LICENSETYPE_CALM            LicenseType = 3
	LICENSETYPE_FLOW            LicenseType = 4
	LICENSETYPE_OBJECT          LicenseType = 5
	LICENSETYPE_AOS             LicenseType = 6
	LICENSETYPE_FILE            LicenseType = 7
	LICENSETYPE_VDI             LicenseType = 8
	LICENSETYPE_ROBO            LicenseType = 9
	LICENSETYPE_MINE            LicenseType = 10
	LICENSETYPE_NCI             LicenseType = 11
	LICENSETYPE_NCM             LicenseType = 12
	LICENSETYPE_NCI_D           LicenseType = 13
	LICENSETYPE_NDA_PLATFORM    LicenseType = 14
	LICENSETYPE_UNIFIED_STORAGE LicenseType = 15
	LICENSETYPE_EUC             LicenseType = 16
	LICENSETYPE_OBJECTS         LicenseType = 17
	LICENSETYPE_ERA             LicenseType = 18
	LICENSETYPE_DRS             LicenseType = 19
	LICENSETYPE_NDS             LicenseType = 20
	LICENSETYPE_NDA             LicenseType = 21
	LICENSETYPE_NDS_PLATFORM    LicenseType = 22
	LICENSETYPE_NDB_PLATFORM    LicenseType = 23
	LICENSETYPE_NUS             LicenseType = 24
	LICENSETYPE_NDB             LicenseType = 25
	LICENSETYPE_NCM_CLOUD       LicenseType = 26
	LICENSETYPE_EDGE            LicenseType = 27
	LICENSETYPE_NO_LICENSE      LicenseType = 28
	LICENSETYPE_NCM_EDGE        LicenseType = 29
	LICENSETYPE_NCM_EUC         LicenseType = 30
	LICENSETYPE_NCI_C           LicenseType = 31
	LICENSETYPE_NKP             LicenseType = 32
	LICENSETYPE_NKPFS           LicenseType = 33
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LicenseType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VDI",
		"ROBO",
		"MINE",
		"NCI",
		"NCM",
		"NCI_D",
		"NDA_PLATFORM",
		"UNIFIED_STORAGE",
		"EUC",
		"OBJECTS",
		"ERA",
		"DRS",
		"NDS",
		"NDA",
		"NDS_PLATFORM",
		"NDB_PLATFORM",
		"NUS",
		"NDB",
		"NCM_CLOUD",
		"EDGE",
		"NO_LICENSE",
		"NCM_EDGE",
		"NCM_EUC",
		"NCI_C",
		"NKP",
		"NKPFS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LicenseType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VDI",
		"ROBO",
		"MINE",
		"NCI",
		"NCM",
		"NCI_D",
		"NDA_PLATFORM",
		"UNIFIED_STORAGE",
		"EUC",
		"OBJECTS",
		"ERA",
		"DRS",
		"NDS",
		"NDA",
		"NDS_PLATFORM",
		"NDB_PLATFORM",
		"NUS",
		"NDB",
		"NCM_CLOUD",
		"EDGE",
		"NO_LICENSE",
		"NCM_EDGE",
		"NCM_EUC",
		"NCI_C",
		"NKP",
		"NKPFS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LicenseType) index(name string) LicenseType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VDI",
		"ROBO",
		"MINE",
		"NCI",
		"NCM",
		"NCI_D",
		"NDA_PLATFORM",
		"UNIFIED_STORAGE",
		"EUC",
		"OBJECTS",
		"ERA",
		"DRS",
		"NDS",
		"NDA",
		"NDS_PLATFORM",
		"NDB_PLATFORM",
		"NUS",
		"NDB",
		"NCM_CLOUD",
		"EDGE",
		"NO_LICENSE",
		"NCM_EDGE",
		"NCM_EUC",
		"NCI_C",
		"NKP",
		"NKPFS",
	}
	for idx := range names {
		if names[idx] == name {
			return LicenseType(idx)
		}
	}
	return LICENSETYPE_UNKNOWN
}

func (e *LicenseType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LicenseType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LicenseType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LicenseType) Ref() *LicenseType {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0/config/allowances Get operation
*/
type ListAllowancesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAllowancesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAllowancesApiResponse() *ListAllowancesApiResponse {
	p := new(ListAllowancesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListAllowancesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAllowancesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAllowancesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAllowancesApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/compliances Get operation
*/
type ListCompliancesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCompliancesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCompliancesApiResponse() *ListCompliancesApiResponse {
	p := new(ListCompliancesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListCompliancesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCompliancesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCompliancesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCompliancesApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/entitlements Get operation
*/
type ListEntitlementsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEntitlementsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListEntitlementsApiResponse() *ListEntitlementsApiResponse {
	p := new(ListEntitlementsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListEntitlementsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEntitlementsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEntitlementsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEntitlementsApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/features Get operation
*/
type ListFeaturesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListFeaturesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListFeaturesApiResponse() *ListFeaturesApiResponse {
	p := new(ListFeaturesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListFeaturesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListFeaturesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListFeaturesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListFeaturesApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/license-keys Get operation
*/
type ListLicenseKeysApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLicenseKeysApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLicenseKeysApiResponse() *ListLicenseKeysApiResponse {
	p := new(ListLicenseKeysApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListLicenseKeysApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLicenseKeysApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLicenseKeysApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLicenseKeysApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/licenses Get operation
*/
type ListLicensesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLicensesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLicensesApiResponse() *ListLicensesApiResponse {
	p := new(ListLicensesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListLicensesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLicensesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLicensesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLicensesApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/recommendations Get operation
*/
type ListRecommendationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecommendationsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRecommendationsApiResponse() *ListRecommendationsApiResponse {
	p := new(ListRecommendationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListRecommendationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecommendationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecommendationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecommendationsApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/settings Get operation
*/
type ListSettingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSettingsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSettingsApiResponse() *ListSettingsApiResponse {
	p := new(ListSettingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListSettingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSettingsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSettingsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSettingsApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0/config/violations Get operation
*/
type ListViolationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListViolationsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListViolationsApiResponse() *ListViolationsApiResponse {
	p := new(ListViolationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListViolationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListViolationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListViolationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListViolationsApiResponseData()
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
Model representing license and cluster logical versions.
*/
type LogicalVersion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates cluster logical version. This is used to keep cluster and license portal in sync.
	*/
	ClusterVersion *int `json:"clusterVersion,omitempty"`
	/*
	  Indicates license logical version. This is also used to keep cluster and license portal in sync.
	*/
	LicenseVersion *int `json:"licenseVersion,omitempty"`
}

func NewLogicalVersion() *LogicalVersion {
	p := new(LogicalVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LogicalVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values of meter are CORES, NODE, TIB, FLASH, USERS, VI, VM_PACKS and VM.
*/
type Meter int

const (
	METER_UNKNOWN  Meter = 0
	METER_REDACTED Meter = 1
	METER_CORES    Meter = 2
	METER_NODE     Meter = 3
	METER_TIB      Meter = 4
	METER_FLASH    Meter = 5
	METER_USERS    Meter = 6
	METER_VI       Meter = 7
	METER_VM_PACKS Meter = 8
	METER_VM       Meter = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Meter) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CORES",
		"NODE",
		"TIB",
		"FLASH",
		"USERS",
		"VI",
		"VM_PACKS",
		"VM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Meter) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CORES",
		"NODE",
		"TIB",
		"FLASH",
		"USERS",
		"VI",
		"VM_PACKS",
		"VM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Meter) index(name string) Meter {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CORES",
		"NODE",
		"TIB",
		"FLASH",
		"USERS",
		"VI",
		"VM_PACKS",
		"VM",
	}
	for idx := range names {
		if names[idx] == name {
			return Meter(idx)
		}
	}
	return METER_UNKNOWN
}

func (e *Meter) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Meter:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Meter) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Meter) Ref() *Meter {
	return &e
}

/*
Possible values are UPGRADE, RENEW, REBALANCE and RECLAIM.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN   OperationType = 0
	OPERATIONTYPE_REDACTED  OperationType = 1
	OPERATIONTYPE_UPGRADE   OperationType = 2
	OPERATIONTYPE_REBALANCE OperationType = 3
	OPERATIONTYPE_RECLAIM   OperationType = 4
	OPERATIONTYPE_RENEW     OperationType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UPGRADE",
		"REBALANCE",
		"RECLAIM",
		"RENEW",
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
		"UPGRADE",
		"REBALANCE",
		"RECLAIM",
		"RENEW",
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
		"UPGRADE",
		"REBALANCE",
		"RECLAIM",
		"RENEW",
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
Response object containing portal setting attributes.
*/
type PortalSetting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Authentication token attribute that is used to fetch keyId and apiKey.
	*/
	AuthToken *string `json:"authToken,omitempty"`
	/*
	  Attribute for capturing the endpoint URL. This URL is used for creation of the API Key.
	*/
	EndpointUrl *string `json:"endpointUrl,omitempty"`
	/*
	  Attribute stating whether portal setting is active or not.
	*/
	IsEnabled *bool `json:"isEnabled"`
}

func (p *PortalSetting) MarshalJSON() ([]byte, error) {
	type PortalSettingProxy PortalSetting
	return json.Marshal(struct {
		*PortalSettingProxy
		IsEnabled *bool `json:"isEnabled,omitempty"`
	}{
		PortalSettingProxy: (*PortalSettingProxy)(p),
		IsEnabled:          p.IsEnabled,
	})
}

func NewPortalSetting() *PortalSetting {
	p := new(PortalSetting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.PortalSetting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are MSP and others.
*/
type PostPaidCategory int

const (
	POSTPAIDCATEGORY_UNKNOWN  PostPaidCategory = 0
	POSTPAIDCATEGORY_REDACTED PostPaidCategory = 1
	POSTPAIDCATEGORY_MSP      PostPaidCategory = 2
	POSTPAIDCATEGORY_NC2      PostPaidCategory = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PostPaidCategory) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NC2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PostPaidCategory) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NC2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PostPaidCategory) index(name string) PostPaidCategory {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NC2",
	}
	for idx := range names {
		if names[idx] == name {
			return PostPaidCategory(idx)
		}
	}
	return POSTPAIDCATEGORY_UNKNOWN
}

func (e *PostPaidCategory) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PostPaidCategory:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PostPaidCategory) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PostPaidCategory) Ref() *PostPaidCategory {
	return &e
}

/*
Model for capturing postpaid configuration.
*/
type PostPaidConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BillingPlan *BillingPlan `json:"billingPlan,omitempty"`

	Category *PostPaidCategory `json:"category,omitempty"`

	ConsumptionType *ConsumptionType `json:"consumptionType,omitempty"`
	/*
	  Attribute for postpaid configuration identifier.
	*/
	Id *string `json:"id,omitempty"`
	/*
	  Indicates whether pulse data collection is required.
	*/
	IsPulseRequired *bool `json:"isPulseRequired,omitempty"`
}

func NewPostPaidConfig() *PostPaidConfig {
	p := new(PostPaidConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.PostPaidConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are PRISM, CALM, AOS and others.
*/
type ProductName int

const (
	PRODUCTNAME_UNKNOWN      ProductName = 0
	PRODUCTNAME_REDACTED     ProductName = 1
	PRODUCTNAME_PRISM        ProductName = 2
	PRODUCTNAME_CALM         ProductName = 3
	PRODUCTNAME_FLOW         ProductName = 4
	PRODUCTNAME_OBJECT       ProductName = 5
	PRODUCTNAME_AOS          ProductName = 6
	PRODUCTNAME_FILE         ProductName = 7
	PRODUCTNAME_VOLUME_GROUP ProductName = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProductName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProductName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProductName) index(name string) ProductName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM",
		"CALM",
		"FLOW",
		"OBJECT",
		"AOS",
		"FILE",
		"VOLUME_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return ProductName(idx)
		}
	}
	return PRODUCTNAME_UNKNOWN
}

func (e *ProductName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProductName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProductName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProductName) Ref() *ProductName {
	return &e
}

/*
Model for capturing recommendation response.
*/
type Recommendation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Array containing recommendations for the cluster and its associated licenses.
	*/
	Details []RecommendationDetail `json:"details,omitempty"`
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

func NewRecommendation() *Recommendation {
	p := new(Recommendation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Recommendation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing recommendation details for the cluster and its corresponding licenses.
*/
type RecommendationDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing comments regarding the recommendation like cluster over licensed and others.
	*/
	Comment *string `json:"comment,omitempty"`
	/*
	  Attribute for capturing the license expiration date.
	*/
	LicenseExpiryDate *time.Time `json:"licenseExpiryDate,omitempty"`
	/*
	  Attribute for capturing the license ID.
	*/
	LicenseId *string `json:"licenseId,omitempty"`

	Operation *OperationType `json:"operation,omitempty"`

	Scope *RecommendationScope `json:"scope,omitempty"`
}

func (p *RecommendationDetail) MarshalJSON() ([]byte, error) {
	type RecommendationDetailProxy RecommendationDetail
	return json.Marshal(struct {
		*RecommendationDetailProxy
	}{
		RecommendationDetailProxy: (*RecommendationDetailProxy)(p),
	})
}

func (p *RecommendationDetail) UnmarshalJSON(b []byte) error {
	type CustomRecommendationDetail struct {
		ObjectType_       *string                `json:"$objectType,omitempty"`
		Reserved_         map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_    map[string]interface{} `json:"$unknownFields,omitempty"`
		Comment           *string                `json:"comment,omitempty"`
		LicenseExpiryDate string                 `json:"licenseExpiryDate,omitempty"`
		LicenseId         *string                `json:"licenseId,omitempty"`
		Operation         *OperationType         `json:"operation,omitempty"`
		Scope             *RecommendationScope   `json:"scope,omitempty"`
	}

	var customRecommendationDetail CustomRecommendationDetail
	err := json.Unmarshal(b, &customRecommendationDetail)
	if err != nil {
		return err
	}

	p.ObjectType_ = customRecommendationDetail.ObjectType_
	p.Reserved_ = customRecommendationDetail.Reserved_
	p.UnknownFields_ = customRecommendationDetail.UnknownFields_
	p.Comment = customRecommendationDetail.Comment
	// Custom date parsing logic for Date field
	if customRecommendationDetail.LicenseExpiryDate != "" {
		parsedLicenseExpiryDate, err := time.Parse("2006-01-02", customRecommendationDetail.LicenseExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field LicenseExpiryDate in struct RecommendationDetail: %s", err))
		}
		p.LicenseExpiryDate = &parsedLicenseExpiryDate
	}
	p.LicenseId = customRecommendationDetail.LicenseId
	p.Operation = customRecommendationDetail.Operation
	p.Scope = customRecommendationDetail.Scope

	return nil
}

func NewRecommendationDetail() *RecommendationDetail {
	p := new(RecommendationDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.RecommendationDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are LICENSE and CLUSTER.
*/
type RecommendationScope int

const (
	RECOMMENDATIONSCOPE_UNKNOWN  RecommendationScope = 0
	RECOMMENDATIONSCOPE_REDACTED RecommendationScope = 1
	RECOMMENDATIONSCOPE_LICENSE  RecommendationScope = 2
	RECOMMENDATIONSCOPE_CLUSTER  RecommendationScope = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecommendationScope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecommendationScope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecommendationScope) index(name string) RecommendationScope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE",
		"CLUSTER",
	}
	for idx := range names {
		if names[idx] == name {
			return RecommendationScope(idx)
		}
	}
	return RECOMMENDATIONSCOPE_UNKNOWN
}

func (e *RecommendationScope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecommendationScope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecommendationScope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecommendationScope) Ref() *RecommendationScope {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0/config/$actions/reset-license-state Post operation
*/
type ResetLicenseStateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResetLicenseStateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewResetLicenseStateApiResponse() *ResetLicenseStateApiResponse {
	p := new(ResetLicenseStateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ResetLicenseStateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ResetLicenseStateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ResetLicenseStateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfResetLicenseStateApiResponseData()
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
Possible values are PC and ALL.
*/
type ResetScope int

const (
	RESETSCOPE_UNKNOWN  ResetScope = 0
	RESETSCOPE_REDACTED ResetScope = 1
	RESETSCOPE_PC       ResetScope = 2
	RESETSCOPE_ALL      ResetScope = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResetScope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ResetScope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ResetScope) index(name string) ResetScope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"ALL",
	}
	for idx := range names {
		if names[idx] == name {
			return ResetScope(idx)
		}
	}
	return RESETSCOPE_UNKNOWN
}

func (e *ResetScope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ResetScope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ResetScope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ResetScope) Ref() *ResetScope {
	return &e
}

/*
Possible values of scope are PC, PE & PC_PE.
*/
type Scope int

const (
	SCOPE_UNKNOWN  Scope = 0
	SCOPE_REDACTED Scope = 1
	SCOPE_PC       Scope = 2
	SCOPE_PE       Scope = 3
	SCOPE_PC_PE    Scope = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Scope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Scope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Scope) index(name string) Scope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_PE",
	}
	for idx := range names {
		if names[idx] == name {
			return Scope(idx)
		}
	}
	return SCOPE_UNKNOWN
}

func (e *Scope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Scope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Scope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Scope) Ref() *Scope {
	return &e
}

/*
Model for capturing Service details.
*/
type Service struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing the list of enforcement actions corresponding to this service.
	*/
	EnforcementActions []EnforcementActions `json:"enforcementActions,omitempty"`

	EnforcementLevel *EnforcementLevel `json:"enforcementLevel,omitempty"`
	/*
	  Attribute for capturing the whether the service is compliant.
	*/
	IsCompliant *bool `json:"isCompliant,omitempty"`

	LicenseType *LicenseType `json:"licenseType,omitempty"`

	Name *ProductName `json:"name,omitempty"`
	/*
	  Attribute for capturing the list of violations corresponding to this service.
	*/
	Violations []ServiceViolation `json:"violations,omitempty"`
}

func NewService() *Service {
	p := new(Service)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Service"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ServiceProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing the list of enforcement actions corresponding to this service.
	*/
	EnforcementActions []EnforcementActions `json:"enforcementActions,omitempty"`

	EnforcementLevel *EnforcementLevel `json:"enforcementLevel,omitempty"`
	/*
	  Attribute for capturing the whether the service is compliant.
	*/
	IsCompliant *bool `json:"isCompliant,omitempty"`

	LicenseType *LicenseType `json:"licenseType,omitempty"`

	Name *ProductName `json:"name,omitempty"`
	/*
	  Attribute for capturing the list of violations corresponding to this service.
	*/
	Violations []ServiceViolation `json:"violations,omitempty"`
}

func NewServiceProjection() *ServiceProjection {
	p := new(ServiceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ServiceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing Violation details for compliance.
*/
type ServiceViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing the start date of violation.
	*/
	StartDate *time.Time `json:"startDate,omitempty"`

	Type *ServiceViolationType `json:"type,omitempty"`
}

func (p *ServiceViolation) MarshalJSON() ([]byte, error) {
	type ServiceViolationProxy ServiceViolation
	return json.Marshal(struct {
		*ServiceViolationProxy
	}{
		ServiceViolationProxy: (*ServiceViolationProxy)(p),
	})
}

func (p *ServiceViolation) UnmarshalJSON(b []byte) error {
	type CustomServiceViolation struct {
		ObjectType_    *string                `json:"$objectType,omitempty"`
		Reserved_      map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
		StartDate      string                 `json:"startDate,omitempty"`
		Type           *ServiceViolationType  `json:"type,omitempty"`
	}

	var customServiceViolation CustomServiceViolation
	err := json.Unmarshal(b, &customServiceViolation)
	if err != nil {
		return err
	}

	p.ObjectType_ = customServiceViolation.ObjectType_
	p.Reserved_ = customServiceViolation.Reserved_
	p.UnknownFields_ = customServiceViolation.UnknownFields_
	// Custom date parsing logic for Date field
	if customServiceViolation.StartDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", customServiceViolation.StartDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field StartDate in struct ServiceViolation: %s", err))
		}
		p.StartDate = &parsedStartDate
	}
	p.Type = customServiceViolation.Type

	return nil
}

func NewServiceViolation() *ServiceViolation {
	p := new(ServiceViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ServiceViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are LICENSE-EXPIRED, FEATURE-VIOLATION and CAPACITY-VIOLATION.
*/
type ServiceViolationType int

const (
	SERVICEVIOLATIONTYPE_UNKNOWN            ServiceViolationType = 0
	SERVICEVIOLATIONTYPE_REDACTED           ServiceViolationType = 1
	SERVICEVIOLATIONTYPE_LICENSE_EXPIRED    ServiceViolationType = 2
	SERVICEVIOLATIONTYPE_FEATURE_VIOLATION  ServiceViolationType = 3
	SERVICEVIOLATIONTYPE_CAPACITY_VIOLATION ServiceViolationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ServiceViolationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE_EXPIRED",
		"FEATURE_VIOLATION",
		"CAPACITY_VIOLATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ServiceViolationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE_EXPIRED",
		"FEATURE_VIOLATION",
		"CAPACITY_VIOLATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ServiceViolationType) index(name string) ServiceViolationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE_EXPIRED",
		"FEATURE_VIOLATION",
		"CAPACITY_VIOLATION",
	}
	for idx := range names {
		if names[idx] == name {
			return ServiceViolationType(idx)
		}
	}
	return SERVICEVIOLATIONTYPE_UNKNOWN
}

func (e *ServiceViolationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ServiceViolationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ServiceViolationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ServiceViolationType) Ref() *ServiceViolationType {
	return &e
}

/*
Response object containing setting details.
*/
type Setting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EnforcementPolicy *EnforcementPolicy `json:"enforcementPolicy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates that user is allowed to use non compliant features.
	*/
	HasNonCompliantFeatures *bool `json:"hasNonCompliantFeatures,omitempty"`
	/*
	  Indicates whether ultimate trial is active or not.
	*/
	HasUltimateTrialEnded *bool `json:"hasUltimateTrialEnded,omitempty"`
	/*
	  Used to enable and disable NCC license checks.
	*/
	IsLicenseCheckDisabled *bool `json:"isLicenseCheckDisabled,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  True value indicates that cluster summary file is downloaded but license summary file is not uplaoded.
	*/
	IsStandby *bool `json:"isStandby,omitempty"`

	LicenseClass *LicenseClass `json:"licenseClass,omitempty"`
	/*
	  License key applied to dark site clusters.
	*/
	LicenseKey *string `json:"licenseKey,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LogicalVersion *LogicalVersion `json:"logicalVersion,omitempty"`

	PostPaidConfig *PostPaidConfig `json:"postPaidConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSetting() *Setting {
	p := new(Setting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Setting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasNonCompliantFeatures = new(bool)
	*p.HasNonCompliantFeatures = false
	p.IsLicenseCheckDisabled = new(bool)
	*p.IsLicenseCheckDisabled = false

	return p
}

/*
Possible values are REGISTERED and UNREGISTERED.
*/
type Status int

const (
	STATUS_UNKNOWN      Status = 0
	STATUS_REDACTED     Status = 1
	STATUS_REGISTERED   Status = 2
	STATUS_UNREGISTERED Status = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Status) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGISTERED",
		"UNREGISTERED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Status) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGISTERED",
		"UNREGISTERED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Status) index(name string) Status {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGISTERED",
		"UNREGISTERED",
	}
	for idx := range names {
		if names[idx] == name {
			return Status(idx)
		}
	}
	return STATUS_UNKNOWN
}

func (e *Status) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Status:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Status) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Status) Ref() *Status {
	return &e
}

/*
Possible values are ADDON, UNLIMITED_CAPACITY.
*/
type SubCategory int

const (
	SUBCATEGORY_UNKNOWN            SubCategory = 0
	SUBCATEGORY_REDACTED           SubCategory = 1
	SUBCATEGORY_ADDON              SubCategory = 2
	SUBCATEGORY_UNLIMITED_CAPACITY SubCategory = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SubCategory) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDON",
		"UNLIMITED_CAPACITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SubCategory) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDON",
		"UNLIMITED_CAPACITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SubCategory) index(name string) SubCategory {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDON",
		"UNLIMITED_CAPACITY",
	}
	for idx := range names {
		if names[idx] == name {
			return SubCategory(idx)
		}
	}
	return SUBCATEGORY_UNKNOWN
}

func (e *SubCategory) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SubCategory:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SubCategory) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SubCategory) Ref() *SubCategory {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0/config/$actions/sync-license-state Post operation
*/
type SyncLicenseStateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSyncLicenseStateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSyncLicenseStateApiResponse() *SyncLicenseStateApiResponse {
	p := new(SyncLicenseStateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.SyncLicenseStateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SyncLicenseStateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SyncLicenseStateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSyncLicenseStateApiResponseData()
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
Possible values are UPLOAD_CLUSTER_SUMMARY, APPLY_LICENSE_SUMMARY, REBALANCE, RECLAIM and others.
*/
type SyncOperationType int

const (
	SYNCOPERATIONTYPE_UNKNOWN                             SyncOperationType = 0
	SYNCOPERATIONTYPE_REDACTED                            SyncOperationType = 1
	SYNCOPERATIONTYPE_UPLOAD_CLUSTER_SUMMARY              SyncOperationType = 2
	SYNCOPERATIONTYPE_APPLY_LICENSE_SUMMARY               SyncOperationType = 3
	SYNCOPERATIONTYPE_RECLAIM                             SyncOperationType = 4
	SYNCOPERATIONTYPE_RENEW                               SyncOperationType = 5
	SYNCOPERATIONTYPE_REBALANCE                           SyncOperationType = 6
	SYNCOPERATIONTYPE_UPDATE_METADATA                     SyncOperationType = 7
	SYNCOPERATIONTYPE_RECLAIM_FOR_DESTROYED_CLUSTER       SyncOperationType = 8
	SYNCOPERATIONTYPE_APPLY_LICENSE_SUMMARY_WITH_METADATA SyncOperationType = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SyncOperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UPLOAD_CLUSTER_SUMMARY",
		"APPLY_LICENSE_SUMMARY",
		"RECLAIM",
		"RENEW",
		"REBALANCE",
		"UPDATE_METADATA",
		"RECLAIM_FOR_DESTROYED_CLUSTER",
		"APPLY_LICENSE_SUMMARY_WITH_METADATA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SyncOperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UPLOAD_CLUSTER_SUMMARY",
		"APPLY_LICENSE_SUMMARY",
		"RECLAIM",
		"RENEW",
		"REBALANCE",
		"UPDATE_METADATA",
		"RECLAIM_FOR_DESTROYED_CLUSTER",
		"APPLY_LICENSE_SUMMARY_WITH_METADATA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SyncOperationType) index(name string) SyncOperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UPLOAD_CLUSTER_SUMMARY",
		"APPLY_LICENSE_SUMMARY",
		"RECLAIM",
		"RENEW",
		"REBALANCE",
		"UPDATE_METADATA",
		"RECLAIM_FOR_DESTROYED_CLUSTER",
		"APPLY_LICENSE_SUMMARY_WITH_METADATA",
	}
	for idx := range names {
		if names[idx] == name {
			return SyncOperationType(idx)
		}
	}
	return SYNCOPERATIONTYPE_UNKNOWN
}

func (e *SyncOperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SyncOperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SyncOperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SyncOperationType) Ref() *SyncOperationType {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0/config/portal-setting Put operation
*/
type UpdatePortalSettingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdatePortalSettingApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdatePortalSettingApiResponse() *UpdatePortalSettingApiResponse {
	p := new(UpdatePortalSettingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.UpdatePortalSettingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdatePortalSettingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdatePortalSettingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdatePortalSettingApiResponseData()
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
Possible values of type are BOOLEAN & INTEGER.
*/
type ValueType int

const (
	VALUETYPE_UNKNOWN  ValueType = 0
	VALUETYPE_REDACTED ValueType = 1
	VALUETYPE_BOOLEAN  ValueType = 2
	VALUETYPE_INTEGER  ValueType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ValueType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOLEAN",
		"INTEGER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ValueType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOLEAN",
		"INTEGER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ValueType) index(name string) ValueType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOLEAN",
		"INTEGER",
	}
	for idx := range names {
		if names[idx] == name {
			return ValueType(idx)
		}
	}
	return VALUETYPE_UNKNOWN
}

func (e *ValueType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ValueType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ValueType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ValueType) Ref() *ValueType {
	return &e
}

/*
Response object containing cluster and associated features violations.
*/
type Violation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of capacity violations. Capacity violations are thrown when user has not applied licenses to extended capacity. Example: Cluster is expanded to have new nodes which are not licensed.
	*/
	CapacityViolations []CapacityViolation `json:"capacityViolations,omitempty"`
	/*
	  List of expired licenses for a cluster. Such licenses need to be replaced in order for the cluster to function properly.
	*/
	ExpiredLicenses []ExpiredLicense `json:"expiredLicenses,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of features which are being used by cluster without applying necessary licenses. Feature violations can be resolved by applying necessary licenses.
	*/
	FeatureViolations []FeatureViolation `json:"featureViolations,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element.
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewViolation() *Violation {
	p := new(Violation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Violation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfListCompliancesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21015 []Compliance           `json:"-"`
	oneOfType401   []ComplianceProjection `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfListCompliancesApiResponseData() *OneOfListCompliancesApiResponseData {
	p := new(OneOfListCompliancesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCompliancesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCompliancesApiResponseData is nil"))
	}
	switch v.(type) {
	case []Compliance:
		p.oneOfType21015 = v.([]Compliance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Compliance>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Compliance>"
	case []ComplianceProjection:
		p.oneOfType401 = v.([]ComplianceProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ComplianceProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ComplianceProjection>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListCompliancesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.Compliance>" == *p.Discriminator {
		return p.oneOfType21015
	}
	if "List<licensing.v4.config.ComplianceProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListCompliancesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21015 := new([]Compliance)
	if err := json.Unmarshal(b, vOneOfType21015); err == nil {
		if len(*vOneOfType21015) == 0 || "licensing.v4.config.Compliance" == *((*vOneOfType21015)[0].ObjectType_) {
			p.oneOfType21015 = *vOneOfType21015
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Compliance>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Compliance>"
			return nil
		}
	}
	vOneOfType401 := new([]ComplianceProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.ComplianceProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ComplianceProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ComplianceProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCompliancesApiResponseData"))
}

func (p *OneOfListCompliancesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.Compliance>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21015)
	}
	if "List<licensing.v4.config.ComplianceProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListCompliancesApiResponseData")
}

type OneOfListRecommendationsApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21003 []Recommendation       `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfListRecommendationsApiResponseData() *OneOfListRecommendationsApiResponseData {
	p := new(OneOfListRecommendationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecommendationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecommendationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Recommendation:
		p.oneOfType21003 = v.([]Recommendation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Recommendation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Recommendation>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListRecommendationsApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.Recommendation>" == *p.Discriminator {
		return p.oneOfType21003
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListRecommendationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21003 := new([]Recommendation)
	if err := json.Unmarshal(b, vOneOfType21003); err == nil {
		if len(*vOneOfType21003) == 0 || "licensing.v4.config.Recommendation" == *((*vOneOfType21003)[0].ObjectType_) {
			p.oneOfType21003 = *vOneOfType21003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Recommendation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Recommendation>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecommendationsApiResponseData"))
}

func (p *OneOfListRecommendationsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.Recommendation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21003)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListRecommendationsApiResponseData")
}

type OneOfListSettingsApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21005 []Setting              `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfListSettingsApiResponseData() *OneOfListSettingsApiResponseData {
	p := new(OneOfListSettingsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSettingsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSettingsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Setting:
		p.oneOfType21005 = v.([]Setting)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Setting>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Setting>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListSettingsApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.Setting>" == *p.Discriminator {
		return p.oneOfType21005
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSettingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21005 := new([]Setting)
	if err := json.Unmarshal(b, vOneOfType21005); err == nil {
		if len(*vOneOfType21005) == 0 || "licensing.v4.config.Setting" == *((*vOneOfType21005)[0].ObjectType_) {
			p.oneOfType21005 = *vOneOfType21005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Setting>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Setting>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSettingsApiResponseData"))
}

func (p *OneOfListSettingsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.Setting>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21005)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSettingsApiResponseData")
}

type OneOfFeatureProjectionValue struct {
	Discriminator  *string `json:"-"`
	ObjectType_    *string `json:"-"`
	oneOfType10032 *int    `json:"-"`
	oneOfType10022 *bool   `json:"-"`
}

func NewOneOfFeatureProjectionValue() *OneOfFeatureProjectionValue {
	p := new(OneOfFeatureProjectionValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFeatureProjectionValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFeatureProjectionValue is nil"))
	}
	switch v.(type) {
	case int:
		if nil == p.oneOfType10032 {
			p.oneOfType10032 = new(int)
		}
		*p.oneOfType10032 = v.(int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
	case bool:
		if nil == p.oneOfType10022 {
			p.oneOfType10022 = new(bool)
		}
		*p.oneOfType10022 = v.(bool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfFeatureProjectionValue) GetValue() interface{} {
	if "Integer" == *p.Discriminator {
		return *p.oneOfType10032
	}
	if "Boolean" == *p.Discriminator {
		return *p.oneOfType10022
	}
	return nil
}

func (p *OneOfFeatureProjectionValue) UnmarshalJSON(b []byte) error {
	vOneOfType10032 := new(int)
	if err := json.Unmarshal(b, vOneOfType10032); err == nil {
		if nil == p.oneOfType10032 {
			p.oneOfType10032 = new(int)
		}
		*p.oneOfType10032 = *vOneOfType10032
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
		return nil
	}
	vOneOfType10022 := new(bool)
	if err := json.Unmarshal(b, vOneOfType10022); err == nil {
		if nil == p.oneOfType10022 {
			p.oneOfType10022 = new(bool)
		}
		*p.oneOfType10022 = *vOneOfType10022
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFeatureProjectionValue"))
}

func (p *OneOfFeatureProjectionValue) MarshalJSON() ([]byte, error) {
	if "Integer" == *p.Discriminator {
		return json.Marshal(p.oneOfType10032)
	}
	if "Boolean" == *p.Discriminator {
		return json.Marshal(p.oneOfType10022)
	}
	return nil, errors.New("No value to marshal for OneOfFeatureProjectionValue")
}

type OneOfUpdatePortalSettingApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType12003 []import1.AppMessage   `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdatePortalSettingApiResponseData() *OneOfUpdatePortalSettingApiResponseData {
	p := new(OneOfUpdatePortalSettingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdatePortalSettingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdatePortalSettingApiResponseData is nil"))
	}
	switch v.(type) {
	case []import1.AppMessage:
		p.oneOfType12003 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfUpdatePortalSettingApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType12003
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdatePortalSettingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType12003 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType12003); err == nil {
		if len(*vOneOfType12003) == 0 || "licensing.v4.error.AppMessage" == *((*vOneOfType12003)[0].ObjectType_) {
			p.oneOfType12003 = *vOneOfType12003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdatePortalSettingApiResponseData"))
}

func (p *OneOfUpdatePortalSettingApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType12003)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdatePortalSettingApiResponseData")
}

type OneOfListFeaturesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType21003 []Feature              `json:"-"`
	oneOfType401   []FeatureProjection    `json:"-"`
}

func NewOneOfListFeaturesApiResponseData() *OneOfListFeaturesApiResponseData {
	p := new(OneOfListFeaturesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListFeaturesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListFeaturesApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Feature:
		p.oneOfType21003 = v.([]Feature)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Feature>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Feature>"
	case []FeatureProjection:
		p.oneOfType401 = v.([]FeatureProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.FeatureProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.FeatureProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListFeaturesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Feature>" == *p.Discriminator {
		return p.oneOfType21003
	}
	if "List<licensing.v4.config.FeatureProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListFeaturesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType21003 := new([]Feature)
	if err := json.Unmarshal(b, vOneOfType21003); err == nil {
		if len(*vOneOfType21003) == 0 || "licensing.v4.config.Feature" == *((*vOneOfType21003)[0].ObjectType_) {
			p.oneOfType21003 = *vOneOfType21003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Feature>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Feature>"
			return nil
		}
	}
	vOneOfType401 := new([]FeatureProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.FeatureProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.FeatureProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.FeatureProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListFeaturesApiResponseData"))
}

func (p *OneOfListFeaturesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Feature>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21003)
	}
	if "List<licensing.v4.config.FeatureProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListFeaturesApiResponseData")
}

type OneOfSyncLicenseStateApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21017 *import3.TaskReference `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfSyncLicenseStateApiResponseData() *OneOfSyncLicenseStateApiResponseData {
	p := new(OneOfSyncLicenseStateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSyncLicenseStateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSyncLicenseStateApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType21017 {
			p.oneOfType21017 = new(import3.TaskReference)
		}
		*p.oneOfType21017 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21017.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21017.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfSyncLicenseStateApiResponseData) GetValue() interface{} {
	if p.oneOfType21017 != nil && *p.oneOfType21017.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21017
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfSyncLicenseStateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21017 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType21017); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType21017.ObjectType_ {
			if nil == p.oneOfType21017 {
				p.oneOfType21017 = new(import3.TaskReference)
			}
			*p.oneOfType21017 = *vOneOfType21017
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21017.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21017.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSyncLicenseStateApiResponseData"))
}

func (p *OneOfSyncLicenseStateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType21017 != nil && *p.oneOfType21017.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21017)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfSyncLicenseStateApiResponseData")
}

type OneOfAssignLicenseKeysApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType31009 []import1.AppMessage   `json:"-"`
}

func NewOneOfAssignLicenseKeysApiResponseData() *OneOfAssignLicenseKeysApiResponseData {
	p := new(OneOfAssignLicenseKeysApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssignLicenseKeysApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssignLicenseKeysApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import1.AppMessage:
		p.oneOfType31009 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAssignLicenseKeysApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType31009
	}
	return nil
}

func (p *OneOfAssignLicenseKeysApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType31009 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType31009); err == nil {
		if len(*vOneOfType31009) == 0 || "licensing.v4.error.AppMessage" == *((*vOneOfType31009)[0].ObjectType_) {
			p.oneOfType31009 = *vOneOfType31009
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssignLicenseKeysApiResponseData"))
}

func (p *OneOfAssignLicenseKeysApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType31009)
	}
	return nil, errors.New("No value to marshal for OneOfAssignLicenseKeysApiResponseData")
}

type OneOfDeleteLicenseKeyApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType31011 *interface{}           `json:"-"`
}

func NewOneOfDeleteLicenseKeyApiResponseData() *OneOfDeleteLicenseKeyApiResponseData {
	p := new(OneOfDeleteLicenseKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteLicenseKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteLicenseKeyApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType31011 {
			p.oneOfType31011 = new(interface{})
		}
		*p.oneOfType31011 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfDeleteLicenseKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType31011
	}
	return nil
}

func (p *OneOfDeleteLicenseKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType31011 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType31011); err == nil {
		if nil == *vOneOfType31011 {
			if nil == p.oneOfType31011 {
				p.oneOfType31011 = new(interface{})
			}
			*p.oneOfType31011 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteLicenseKeyApiResponseData"))
}

func (p *OneOfDeleteLicenseKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType31011)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteLicenseKeyApiResponseData")
}

type OneOfGetLicenseKeyApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType31013 *LicenseKey            `json:"-"`
}

func NewOneOfGetLicenseKeyApiResponseData() *OneOfGetLicenseKeyApiResponseData {
	p := new(OneOfGetLicenseKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLicenseKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLicenseKeyApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case LicenseKey:
		if nil == p.oneOfType31013 {
			p.oneOfType31013 = new(LicenseKey)
		}
		*p.oneOfType31013 = v.(LicenseKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType31013.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType31013.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetLicenseKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType31013 != nil && *p.oneOfType31013.ObjectType_ == *p.Discriminator {
		return *p.oneOfType31013
	}
	return nil
}

func (p *OneOfGetLicenseKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType31013 := new(LicenseKey)
	if err := json.Unmarshal(b, vOneOfType31013); err == nil {
		if "licensing.v4.config.LicenseKey" == *vOneOfType31013.ObjectType_ {
			if nil == p.oneOfType31013 {
				p.oneOfType31013 = new(LicenseKey)
			}
			*p.oneOfType31013 = *vOneOfType31013
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType31013.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType31013.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLicenseKeyApiResponseData"))
}

func (p *OneOfGetLicenseKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType31013 != nil && *p.oneOfType31013.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType31013)
	}
	return nil, errors.New("No value to marshal for OneOfGetLicenseKeyApiResponseData")
}

type OneOfListEntitlementsApiResponseData struct {
	Discriminator  *string                 `json:"-"`
	ObjectType_    *string                 `json:"-"`
	oneOfType401   []EntitlementProjection `json:"-"`
	oneOfType400   *import1.ErrorResponse  `json:"-"`
	oneOfType21010 []Entitlement           `json:"-"`
}

func NewOneOfListEntitlementsApiResponseData() *OneOfListEntitlementsApiResponseData {
	p := new(OneOfListEntitlementsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEntitlementsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEntitlementsApiResponseData is nil"))
	}
	switch v.(type) {
	case []EntitlementProjection:
		p.oneOfType401 = v.([]EntitlementProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.EntitlementProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.EntitlementProjection>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Entitlement:
		p.oneOfType21010 = v.([]Entitlement)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Entitlement>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Entitlement>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListEntitlementsApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.EntitlementProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Entitlement>" == *p.Discriminator {
		return p.oneOfType21010
	}
	return nil
}

func (p *OneOfListEntitlementsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]EntitlementProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.EntitlementProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.EntitlementProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.EntitlementProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType21010 := new([]Entitlement)
	if err := json.Unmarshal(b, vOneOfType21010); err == nil {
		if len(*vOneOfType21010) == 0 || "licensing.v4.config.Entitlement" == *((*vOneOfType21010)[0].ObjectType_) {
			p.oneOfType21010 = *vOneOfType21010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Entitlement>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Entitlement>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEntitlementsApiResponseData"))
}

func (p *OneOfListEntitlementsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.EntitlementProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Entitlement>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21010)
	}
	return nil, errors.New("No value to marshal for OneOfListEntitlementsApiResponseData")
}

type OneOfResetLicenseStateApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType21017 []import1.AppMessage   `json:"-"`
}

func NewOneOfResetLicenseStateApiResponseData() *OneOfResetLicenseStateApiResponseData {
	p := new(OneOfResetLicenseStateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfResetLicenseStateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfResetLicenseStateApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []import1.AppMessage:
		p.oneOfType21017 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfResetLicenseStateApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType21017
	}
	return nil
}

func (p *OneOfResetLicenseStateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType21017 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType21017); err == nil {
		if len(*vOneOfType21017) == 0 || "licensing.v4.error.AppMessage" == *((*vOneOfType21017)[0].ObjectType_) {
			p.oneOfType21017 = *vOneOfType21017
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.error.AppMessage>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetLicenseStateApiResponseData"))
}

func (p *OneOfResetLicenseStateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21017)
	}
	return nil, errors.New("No value to marshal for OneOfResetLicenseStateApiResponseData")
}

type OneOfListAllowancesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21012 []Allowance            `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType401   []AllowanceProjection  `json:"-"`
}

func NewOneOfListAllowancesApiResponseData() *OneOfListAllowancesApiResponseData {
	p := new(OneOfListAllowancesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAllowancesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAllowancesApiResponseData is nil"))
	}
	switch v.(type) {
	case []Allowance:
		p.oneOfType21012 = v.([]Allowance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Allowance>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Allowance>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []AllowanceProjection:
		p.oneOfType401 = v.([]AllowanceProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.AllowanceProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.AllowanceProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAllowancesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.Allowance>" == *p.Discriminator {
		return p.oneOfType21012
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.AllowanceProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListAllowancesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21012 := new([]Allowance)
	if err := json.Unmarshal(b, vOneOfType21012); err == nil {
		if len(*vOneOfType21012) == 0 || "licensing.v4.config.Allowance" == *((*vOneOfType21012)[0].ObjectType_) {
			p.oneOfType21012 = *vOneOfType21012
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Allowance>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Allowance>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType401 := new([]AllowanceProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.AllowanceProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.AllowanceProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.AllowanceProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAllowancesApiResponseData"))
}

func (p *OneOfListAllowancesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.Allowance>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21012)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.AllowanceProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListAllowancesApiResponseData")
}

type OneOfListLicensesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType401   []LicenseProjection    `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType21003 []License              `json:"-"`
}

func NewOneOfListLicensesApiResponseData() *OneOfListLicensesApiResponseData {
	p := new(OneOfListLicensesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLicensesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLicensesApiResponseData is nil"))
	}
	switch v.(type) {
	case []LicenseProjection:
		p.oneOfType401 = v.([]LicenseProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.LicenseProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.LicenseProjection>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []License:
		p.oneOfType21003 = v.([]License)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.License>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.License>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLicensesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.LicenseProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.License>" == *p.Discriminator {
		return p.oneOfType21003
	}
	return nil
}

func (p *OneOfListLicensesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]LicenseProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.LicenseProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.LicenseProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.LicenseProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType21003 := new([]License)
	if err := json.Unmarshal(b, vOneOfType21003); err == nil {
		if len(*vOneOfType21003) == 0 || "licensing.v4.config.License" == *((*vOneOfType21003)[0].ObjectType_) {
			p.oneOfType21003 = *vOneOfType21003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.License>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.License>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLicensesApiResponseData"))
}

func (p *OneOfListLicensesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.LicenseProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.License>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21003)
	}
	return nil, errors.New("No value to marshal for OneOfListLicensesApiResponseData")
}

type OneOfListViolationsApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
	oneOfType21008 []Violation            `json:"-"`
}

func NewOneOfListViolationsApiResponseData() *OneOfListViolationsApiResponseData {
	p := new(OneOfListViolationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListViolationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListViolationsApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Violation:
		p.oneOfType21008 = v.([]Violation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.Violation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.Violation>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListViolationsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Violation>" == *p.Discriminator {
		return p.oneOfType21008
	}
	return nil
}

func (p *OneOfListViolationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	vOneOfType21008 := new([]Violation)
	if err := json.Unmarshal(b, vOneOfType21008); err == nil {
		if len(*vOneOfType21008) == 0 || "licensing.v4.config.Violation" == *((*vOneOfType21008)[0].ObjectType_) {
			p.oneOfType21008 = *vOneOfType21008
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.Violation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.Violation>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListViolationsApiResponseData"))
}

func (p *OneOfListViolationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Violation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21008)
	}
	return nil, errors.New("No value to marshal for OneOfListViolationsApiResponseData")
}

type OneOfFeatureValue struct {
	Discriminator  *string `json:"-"`
	ObjectType_    *string `json:"-"`
	oneOfType10032 *int    `json:"-"`
	oneOfType10022 *bool   `json:"-"`
}

func NewOneOfFeatureValue() *OneOfFeatureValue {
	p := new(OneOfFeatureValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFeatureValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFeatureValue is nil"))
	}
	switch v.(type) {
	case int:
		if nil == p.oneOfType10032 {
			p.oneOfType10032 = new(int)
		}
		*p.oneOfType10032 = v.(int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
	case bool:
		if nil == p.oneOfType10022 {
			p.oneOfType10022 = new(bool)
		}
		*p.oneOfType10022 = v.(bool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfFeatureValue) GetValue() interface{} {
	if "Integer" == *p.Discriminator {
		return *p.oneOfType10032
	}
	if "Boolean" == *p.Discriminator {
		return *p.oneOfType10022
	}
	return nil
}

func (p *OneOfFeatureValue) UnmarshalJSON(b []byte) error {
	vOneOfType10032 := new(int)
	if err := json.Unmarshal(b, vOneOfType10032); err == nil {
		if nil == p.oneOfType10032 {
			p.oneOfType10032 = new(int)
		}
		*p.oneOfType10032 = *vOneOfType10032
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
		return nil
	}
	vOneOfType10022 := new(bool)
	if err := json.Unmarshal(b, vOneOfType10022); err == nil {
		if nil == p.oneOfType10022 {
			p.oneOfType10022 = new(bool)
		}
		*p.oneOfType10022 = *vOneOfType10022
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFeatureValue"))
}

func (p *OneOfFeatureValue) MarshalJSON() ([]byte, error) {
	if "Integer" == *p.Discriminator {
		return json.Marshal(p.oneOfType10032)
	}
	if "Boolean" == *p.Discriminator {
		return json.Marshal(p.oneOfType10022)
	}
	return nil, errors.New("No value to marshal for OneOfFeatureValue")
}

type OneOfAddLicenseKeyApiResponseData struct {
	Discriminator  *string                         `json:"-"`
	ObjectType_    *string                         `json:"-"`
	oneOfType401   *AddLicenseKeyDryRunApiResponse `json:"-"`
	oneOfType31003 *LicenseKey                     `json:"-"`
	oneOfType400   *import1.ErrorResponse          `json:"-"`
}

func NewOneOfAddLicenseKeyApiResponseData() *OneOfAddLicenseKeyApiResponseData {
	p := new(OneOfAddLicenseKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddLicenseKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddLicenseKeyApiResponseData is nil"))
	}
	switch v.(type) {
	case AddLicenseKeyDryRunApiResponse:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(AddLicenseKeyDryRunApiResponse)
		}
		*p.oneOfType401 = v.(AddLicenseKeyDryRunApiResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	case LicenseKey:
		if nil == p.oneOfType31003 {
			p.oneOfType31003 = new(LicenseKey)
		}
		*p.oneOfType31003 = v.(LicenseKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType31003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType31003.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfAddLicenseKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType31003 != nil && *p.oneOfType31003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType31003
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddLicenseKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new(AddLicenseKeyDryRunApiResponse)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if "licensing.v4.config.AddLicenseKeyDryRunApiResponse" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(AddLicenseKeyDryRunApiResponse)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	vOneOfType31003 := new(LicenseKey)
	if err := json.Unmarshal(b, vOneOfType31003); err == nil {
		if "licensing.v4.config.LicenseKey" == *vOneOfType31003.ObjectType_ {
			if nil == p.oneOfType31003 {
				p.oneOfType31003 = new(LicenseKey)
			}
			*p.oneOfType31003 = *vOneOfType31003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType31003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType31003.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddLicenseKeyApiResponseData"))
}

func (p *OneOfAddLicenseKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType31003 != nil && *p.oneOfType31003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType31003)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddLicenseKeyApiResponseData")
}

type OneOfListLicenseKeysApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType31007 []LicenseKey           `json:"-"`
	oneOfType401   []LicenseKeyProjection `json:"-"`
	oneOfType400   *import1.ErrorResponse `json:"-"`
}

func NewOneOfListLicenseKeysApiResponseData() *OneOfListLicenseKeysApiResponseData {
	p := new(OneOfListLicenseKeysApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLicenseKeysApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLicenseKeysApiResponseData is nil"))
	}
	switch v.(type) {
	case []LicenseKey:
		p.oneOfType31007 = v.([]LicenseKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.LicenseKey>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.LicenseKey>"
	case []LicenseKeyProjection:
		p.oneOfType401 = v.([]LicenseKeyProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.LicenseKeyProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.LicenseKeyProjection>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfListLicenseKeysApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.LicenseKey>" == *p.Discriminator {
		return p.oneOfType31007
	}
	if "List<licensing.v4.config.LicenseKeyProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListLicenseKeysApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType31007 := new([]LicenseKey)
	if err := json.Unmarshal(b, vOneOfType31007); err == nil {
		if len(*vOneOfType31007) == 0 || "licensing.v4.config.LicenseKey" == *((*vOneOfType31007)[0].ObjectType_) {
			p.oneOfType31007 = *vOneOfType31007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.LicenseKey>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.LicenseKey>"
			return nil
		}
	}
	vOneOfType401 := new([]LicenseKeyProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "licensing.v4.config.LicenseKeyProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.LicenseKeyProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.LicenseKeyProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLicenseKeysApiResponseData"))
}

func (p *OneOfListLicenseKeysApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.LicenseKey>" == *p.Discriminator {
		return json.Marshal(p.oneOfType31007)
	}
	if "List<licensing.v4.config.LicenseKeyProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListLicenseKeysApiResponseData")
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
