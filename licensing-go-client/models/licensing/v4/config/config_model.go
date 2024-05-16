/*
 * Generated file models/licensing/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Licensing Versioned APIs
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
	import1 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/error"
	"time"
)

/*
Model representing allowance details
*/
type AllowanceDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates feature id like MAX_NODES, APP_MONITORING
	*/
	FeatureId *string `json:"featureId,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	Type *TypeEnum `json:"type,omitempty"`
	/*
	  Attribute for feature value.For boolean, it will be true/false.For integer it will be any numeric value. True value indicates feature is available to use
	*/
	Value *string `json:"value,omitempty"`
}

func NewAllowanceDetail() *AllowanceDetail {
	p := new(AllowanceDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AllowanceDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AllowanceDetailProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates feature id like MAX_NODES, APP_MONITORING
	*/
	FeatureId *string `json:"featureId,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	Type *TypeEnum `json:"type,omitempty"`
	/*
	  Attribute for feature value.For boolean, it will be true/false.For integer it will be any numeric value. True value indicates feature is available to use
	*/
	Value *string `json:"value,omitempty"`
}

func NewAllowanceDetailProjection() *AllowanceDetailProjection {
	p := new(AllowanceDetailProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.AllowanceDetailProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing cluster identifier and a boolean attribute representing whether it is a PC or PE
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
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewBaseClusterInfo() *BaseClusterInfo {
	p := new(BaseClusterInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.BaseClusterInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing base licensing attributes like name, type, category and others
*/
type BaseLicenseInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Indicates expiry date of license
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewBaseLicenseInfo() *BaseLicenseInfo {
	p := new(BaseLicenseInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.BaseLicenseInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing capacity violation. Capacity violations are thrown when user has not applied licenses to extended capacity.
*/
type CapacityViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Indicates the insufficient quantity of a license.License of given quantity should be applied for a cluster to function properly
	*/
	InsufficientQuantity *float64 `json:"insufficientQuantity,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewCapacityViolation() *CapacityViolation {
	p := new(CapacityViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.CapacityViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model capturing cluster details along with associated allowances.
*/
type ClusterAllowance struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing collection of allowance details
	*/
	AllowanceDetails []AllowanceDetail `json:"allowanceDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing cluster name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewClusterAllowance() *ClusterAllowance {
	p := new(ClusterAllowance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterAllowance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterAllowanceProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AllowanceDetailProjection *AllowanceDetailProjection `json:"allowanceDetailProjection,omitempty"`
	/*
	  Attribute for capturing collection of allowance details
	*/
	AllowanceDetails []AllowanceDetail `json:"allowanceDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing cluster name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewClusterAllowanceProjection() *ClusterAllowanceProjection {
	p := new(ClusterAllowanceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterAllowanceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing cluster details along with associated entitlements
*/
type ClusterEntitlement struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates cluster UUID.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Attribute for capturing collection of entitlement details.
	*/
	EntitlementDetails []EntitlementDetail `json:"entitlementDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  Indicates whether cluster is registered with PC or not.
	*/
	IsRegistered *bool `json:"isRegistered,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing cluster name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewClusterEntitlement() *ClusterEntitlement {
	p := new(ClusterEntitlement)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterEntitlement"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterEntitlementProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates cluster UUID.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	EntitlementDetailProjection *EntitlementDetailProjection `json:"entitlementDetailProjection,omitempty"`
	/*
	  Attribute for capturing collection of entitlement details.
	*/
	EntitlementDetails []EntitlementDetail `json:"entitlementDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  Indicates whether cluster is registered with PC or not.
	*/
	IsRegistered *bool `json:"isRegistered,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing cluster name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewClusterEntitlementProjection() *ClusterEntitlementProjection {
	p := new(ClusterEntitlementProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterEntitlementProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are NUTANIX,NON_NUTANIX and others
*/
type ClusterTypeEnum int

const (
	CLUSTERTYPEENUM_UNKNOWN     ClusterTypeEnum = 0
	CLUSTERTYPEENUM_REDACTED    ClusterTypeEnum = 1
	CLUSTERTYPEENUM_NUTANIX     ClusterTypeEnum = 2
	CLUSTERTYPEENUM_VMWARE      ClusterTypeEnum = 3
	CLUSTERTYPEENUM_NON_NUTANIX ClusterTypeEnum = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterTypeEnum) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"VMWARE",
		"NON_NUTANIX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterTypeEnum) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"VMWARE",
		"NON_NUTANIX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterTypeEnum) index(name string) ClusterTypeEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"VMWARE",
		"NON_NUTANIX",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterTypeEnum(idx)
		}
	}
	return CLUSTERTYPEENUM_UNKNOWN
}

func (e *ClusterTypeEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterTypeEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterTypeEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterTypeEnum) Ref() *ClusterTypeEnum {
	return &e
}

/*
Response object containing cluster and associated features violations.
*/
type ClusterViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of capacity violations. Capacity violations are thrown when user has not applied licenses to extended capacity. Example: Cluster is expanded to have new nodes which are not licensed.
	*/
	CapacityViolations []CapacityViolation `json:"capacityViolations,omitempty"`
	/*
	  List of expired licenses for a cluster.Such licenses needs to be replaced in order for cluster to function properly.
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
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewClusterViolation() *ClusterViolation {
	p := new(ClusterViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ClusterViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Attribute for capturing collection of service details.
	*/
	Services []Service `json:"services,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewCompliance() *Compliance {
	p := new(Compliance)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Compliance"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	ServiceProjection *ServiceProjection `json:"serviceProjection,omitempty"`
	/*
	  Attribute for capturing collection of service details.
	*/
	Services []Service `json:"services,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ClusterTypeEnum `json:"type,omitempty"`
}

func NewComplianceProjection() *ComplianceProjection {
	p := new(ComplianceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ComplianceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model containing license consumption details like cluster id and quantity used in that cluster
*/
type Consumption struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID which is consuming the license
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Indicates quantity of licenses consumed by the cluster
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewConsumption() *Consumption {
	p := new(Consumption)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Consumption"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConsumptionProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster UUID which is consuming the license
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Indicates quantity of licenses consumed by the cluster
	*/
	QuantityUsed *float64 `json:"quantityUsed,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewConsumptionProjection() *ConsumptionProjection {
	p := new(ConsumptionProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ConsumptionProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are NO_NODE_ADDITION, NO_SUPPORT and others.
*/
type EnforcementActionsEnum int

const (
	ENFORCEMENTACTIONSENUM_UNKNOWN               EnforcementActionsEnum = 0
	ENFORCEMENTACTIONSENUM_REDACTED              EnforcementActionsEnum = 1
	ENFORCEMENTACTIONSENUM_NO_NODE_ADDITION      EnforcementActionsEnum = 2
	ENFORCEMENTACTIONSENUM_NO_ADVERTISE_CAPACITY EnforcementActionsEnum = 3
	ENFORCEMENTACTIONSENUM_NO_SUPPORT            EnforcementActionsEnum = 4
	ENFORCEMENTACTIONSENUM_NO_CONTAINER_UPDATE   EnforcementActionsEnum = 5
	ENFORCEMENTACTIONSENUM_NO_LOGIN              EnforcementActionsEnum = 6
	ENFORCEMENTACTIONSENUM_NO_CLUSTER_PAGE       EnforcementActionsEnum = 7
	ENFORCEMENTACTIONSENUM_NO_UPGRADES           EnforcementActionsEnum = 8
	ENFORCEMENTACTIONSENUM_NO_SECURITY_PATCH     EnforcementActionsEnum = 9
	ENFORCEMENTACTIONSENUM_SHOW_NAGWARE          EnforcementActionsEnum = 10
	ENFORCEMENTACTIONSENUM_NO_APP_LAUNCH         EnforcementActionsEnum = 11
	ENFORCEMENTACTIONSENUM_NO_RUNBOOK            EnforcementActionsEnum = 12
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementActionsEnum) name(index int) string {
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
func (e EnforcementActionsEnum) GetName() string {
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
func (e *EnforcementActionsEnum) index(name string) EnforcementActionsEnum {
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
			return EnforcementActionsEnum(idx)
		}
	}
	return ENFORCEMENTACTIONSENUM_UNKNOWN
}

func (e *EnforcementActionsEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementActionsEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementActionsEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementActionsEnum) Ref() *EnforcementActionsEnum {
	return &e
}

/*
Possible values are LEVEL_10 and LEVEL_20.
*/
type EnforcementLevelEnum int

const (
	ENFORCEMENTLEVELENUM_UNKNOWN  EnforcementLevelEnum = 0
	ENFORCEMENTLEVELENUM_REDACTED EnforcementLevelEnum = 1
	ENFORCEMENTLEVELENUM_LEVEL_10 EnforcementLevelEnum = 2
	ENFORCEMENTLEVELENUM_LEVEL_20 EnforcementLevelEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementLevelEnum) name(index int) string {
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
func (e EnforcementLevelEnum) GetName() string {
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
func (e *EnforcementLevelEnum) index(name string) EnforcementLevelEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LEVEL_10",
		"LEVEL_20",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementLevelEnum(idx)
		}
	}
	return ENFORCEMENTLEVELENUM_UNKNOWN
}

func (e *EnforcementLevelEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementLevelEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementLevelEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementLevelEnum) Ref() *EnforcementLevelEnum {
	return &e
}

/*
Possible values are All or None
*/
type EnforcementPolicyEnum int

const (
	ENFORCEMENTPOLICYENUM_UNKNOWN  EnforcementPolicyEnum = 0
	ENFORCEMENTPOLICYENUM_REDACTED EnforcementPolicyEnum = 1
	ENFORCEMENTPOLICYENUM_ALL      EnforcementPolicyEnum = 2
	ENFORCEMENTPOLICYENUM_NONE     EnforcementPolicyEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnforcementPolicyEnum) name(index int) string {
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
func (e EnforcementPolicyEnum) GetName() string {
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
func (e *EnforcementPolicyEnum) index(name string) EnforcementPolicyEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALL",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return EnforcementPolicyEnum(idx)
		}
	}
	return ENFORCEMENTPOLICYENUM_UNKNOWN
}

func (e *EnforcementPolicyEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnforcementPolicyEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnforcementPolicyEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnforcementPolicyEnum) Ref() *EnforcementPolicyEnum {
	return &e
}

/*
Model capturing entitlement details
*/
type EntitlementDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Attribute for capturing earliest expiry date across entitlements
	*/
	EarliestExpiryDate *time.Time `json:"earliestExpiryDate,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Attribute for capturing quantity for entitlement
	*/
	Quantity *float64 `json:"quantity,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewEntitlementDetail() *EntitlementDetail {
	p := new(EntitlementDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.EntitlementDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EntitlementDetailProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Attribute for capturing earliest expiry date across entitlements
	*/
	EarliestExpiryDate *time.Time `json:"earliestExpiryDate,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Attribute for capturing quantity for entitlement
	*/
	Quantity *float64 `json:"quantity,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewEntitlementDetailProjection() *EntitlementDetailProjection {
	p := new(EntitlementDetailProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.EntitlementDetailProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
License details of expired license
*/
type ExpiredLicense struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Indicates expiry date of license
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
	Links []import1.ApiLink `json:"links,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
	/*
	  Indicates expired quantity of the license
	*/
	UsedQuantity *float64 `json:"usedQuantity,omitempty"`
}

func NewExpiredLicense() *ExpiredLicense {
	p := new(ExpiredLicense)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ExpiredLicense"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response object containing feature details like name, description, license type and category mapping and others
*/
type Feature struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	LicenseCategory *LicenseCategoryEnum `json:"licenseCategory,omitempty"`

	LicenseSubCategory *SubCategoryEnum `json:"licenseSubCategory,omitempty"`

	LicenseType *LicenseTypeEnum `json:"licenseType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of feature like dp_recovery, dp_backup_tiering
	*/
	Name *string `json:"name,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *TypeEnum `json:"type,omitempty"`
	/*
	  Value of feature, it could be true, false or integer
	*/
	Value *string `json:"value,omitempty"`
}

func NewFeature() *Feature {
	p := new(Feature)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Feature"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model for capturing feature details
*/
type FeatureDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for feature description
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Attribute for capturing feature id like VULCAN, APP_MONITORING
	*/
	FeatureId *string `json:"featureId,omitempty"`
	/*
	  Attribute for capturing feature name like Application Monitoring
	*/
	Name *string `json:"name,omitempty"`
}

func NewFeatureDetail() *FeatureDetail {
	p := new(FeatureDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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

	LicenseCategory *LicenseCategoryEnum `json:"licenseCategory,omitempty"`

	LicenseSubCategory *SubCategoryEnum `json:"licenseSubCategory,omitempty"`

	LicenseType *LicenseTypeEnum `json:"licenseType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of feature like dp_recovery, dp_backup_tiering
	*/
	Name *string `json:"name,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *TypeEnum `json:"type,omitempty"`
	/*
	  Value of feature, it could be true, false or integer
	*/
	Value *string `json:"value,omitempty"`
}

func NewFeatureProjection() *FeatureProjection {
	p := new(FeatureProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model capturing feature info and associated clusters
*/
type FeatureViolation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Description of entity for which feature violation is thrown.
	Examples: 1) Vm with uuid 000604af-3aa4-9cfe-1c6a-ac1f6b357fb7.
	          2) Storage Container bucket-data-prod
	          3) Protection Domain pd_prod with Application consistency group prod
	*/
	AffectedEntity *string `json:"affectedEntity,omitempty"`
	/*
	  Attribute for feature description
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Attribute for capturing feature id like VULCAN, APP_MONITORING
	*/
	FeatureId *string `json:"featureId,omitempty"`
	/*
	  Attribute for capturing feature name like Application Monitoring
	*/
	Name *string `json:"name,omitempty"`
}

func NewFeatureViolation() *FeatureViolation {
	p := new(FeatureViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.FeatureViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /licensing/v4.0.a1/config/clusters/{extId}/allowances Get operation
*/
type GetClusterAllowanceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterAllowanceApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClusterAllowanceApiResponse() *GetClusterAllowanceApiResponse {
	p := new(GetClusterAllowanceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetClusterAllowanceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClusterAllowanceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClusterAllowanceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClusterAllowanceApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/clusters/{extId}/entitlements Get operation
*/
type GetClusterEntitlementApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterEntitlementApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClusterEntitlementApiResponse() *GetClusterEntitlementApiResponse {
	p := new(GetClusterEntitlementApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetClusterEntitlementApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClusterEntitlementApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClusterEntitlementApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClusterEntitlementApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/clusters/{extId}/violations Get operation
*/
type GetClusterViolationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterViolationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClusterViolationApiResponse() *GetClusterViolationApiResponse {
	p := new(GetClusterViolationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetClusterViolationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClusterViolationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClusterViolationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClusterViolationApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/clusters/{extId}/compliance Get operation
*/
type GetComplianceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetComplianceApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetComplianceApiResponse() *GetComplianceApiResponse {
	p := new(GetComplianceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetComplianceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetComplianceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetComplianceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetComplianceApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/licenses/{extId} Get operation
*/
type GetLicenseApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLicenseApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLicenseApiResponse() *GetLicenseApiResponse {
	p := new(GetLicenseApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetLicenseApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLicenseApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLicenseApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLicenseApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/clusters/{extId}/settings Get operation
*/
type GetSettingApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSettingApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSettingApiResponse() *GetSettingApiResponse {
	p := new(GetSettingApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.GetSettingApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSettingApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSettingApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSettingApiResponseData()
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
Model containing license details like id, type, category, sub-category, scope and others
*/
type License struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Array containing information about the clusters where these licenses are used
	*/
	ConsumptionDetails []Consumption `json:"consumptionDetails,omitempty"`
	/*
	  Indicates expiry date of license
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Total quantity of license consumed
	*/
	Quantity *string `json:"quantity,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewLicense() *License {
	p := new(License)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.License"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are Starter,Pro,Ultimate and others
*/
type LicenseCategoryEnum int

const (
	LICENSECATEGORYENUM_UNKNOWN              LicenseCategoryEnum = 0
	LICENSECATEGORYENUM_REDACTED             LicenseCategoryEnum = 1
	LICENSECATEGORYENUM_STARTER              LicenseCategoryEnum = 2
	LICENSECATEGORYENUM_PRO                  LicenseCategoryEnum = 3
	LICENSECATEGORYENUM_ULTIMATE             LicenseCategoryEnum = 4
	LICENSECATEGORYENUM_CALM                 LicenseCategoryEnum = 5
	LICENSECATEGORYENUM_STANDARD             LicenseCategoryEnum = 6
	LICENSECATEGORYENUM_AOS_MINE             LicenseCategoryEnum = 7
	LICENSECATEGORYENUM_SOFTWARE_ENCRYPTION  LicenseCategoryEnum = 8
	LICENSECATEGORYENUM_ADV_REPLICATION      LicenseCategoryEnum = 9
	LICENSECATEGORYENUM_OBJECT               LicenseCategoryEnum = 10
	LICENSECATEGORYENUM_ULTIMATE_TRIAL       LicenseCategoryEnum = 11
	LICENSECATEGORYENUM_PRISM_STARTER        LicenseCategoryEnum = 12
	LICENSECATEGORYENUM_PRO_SPECIAL          LicenseCategoryEnum = 13
	LICENSECATEGORYENUM_ADR                  LicenseCategoryEnum = 14
	LICENSECATEGORYENUM_SECURITY             LicenseCategoryEnum = 15
	LICENSECATEGORYENUM_NKS                  LicenseCategoryEnum = 16
	LICENSECATEGORYENUM_APPAUTOMATION        LicenseCategoryEnum = 17
	LICENSECATEGORYENUM_NDA                  LicenseCategoryEnum = 18
	LICENSECATEGORYENUM_UST                  LicenseCategoryEnum = 19
	LICENSECATEGORYENUM_ANALYTICS            LicenseCategoryEnum = 20
	LICENSECATEGORYENUM_STANDALONE           LicenseCategoryEnum = 21
	LICENSECATEGORYENUM_DRASS                LicenseCategoryEnum = 22
	LICENSECATEGORYENUM_CLOUD_NATIVE         LicenseCategoryEnum = 23
	LICENSECATEGORYENUM_DATA_ENCRYPTION      LicenseCategoryEnum = 24
	LICENSECATEGORYENUM_NDS                  LicenseCategoryEnum = 25
	LICENSECATEGORYENUM_NDB                  LicenseCategoryEnum = 26
	LICENSECATEGORYENUM_NO_LICENSE           LicenseCategoryEnum = 27
	LICENSECATEGORYENUM_NUS_ENCRYPTION       LicenseCategoryEnum = 28
	LICENSECATEGORYENUM_NUS_REPLICATION      LicenseCategoryEnum = 29
	LICENSECATEGORYENUM_CLOUD_PRO            LicenseCategoryEnum = 30
	LICENSECATEGORYENUM_CLOUD_ULTIMATE       LicenseCategoryEnum = 31
	LICENSECATEGORYENUM_CLOUD                LicenseCategoryEnum = 32
	LICENSECATEGORYENUM_PUBLIC_CLOUD         LicenseCategoryEnum = 33
	LICENSECATEGORYENUM_ADVANCED_REPLICATION LicenseCategoryEnum = 34
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LicenseCategoryEnum) name(index int) string {
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
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LicenseCategoryEnum) GetName() string {
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
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LicenseCategoryEnum) index(name string) LicenseCategoryEnum {
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
	}
	for idx := range names {
		if names[idx] == name {
			return LicenseCategoryEnum(idx)
		}
	}
	return LICENSECATEGORYENUM_UNKNOWN
}

func (e *LicenseCategoryEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LicenseCategoryEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LicenseCategoryEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LicenseCategoryEnum) Ref() *LicenseCategoryEnum {
	return &e
}

type LicenseProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Category *LicenseCategoryEnum `json:"category,omitempty"`
	/*
	  Array containing information about the clusters where these licenses are used
	*/
	ConsumptionDetails []Consumption `json:"consumptionDetails,omitempty"`

	ConsumptionProjection *ConsumptionProjection `json:"consumptionProjection,omitempty"`
	/*
	  Indicates expiry date of license
	*/
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Meter *MeterEnum `json:"meter,omitempty"`
	/*
	  Indicates name of the license. Possible values could be NCI Pro, NCM Ultimate
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Total quantity of license consumed
	*/
	Quantity *string `json:"quantity,omitempty"`

	Scope *ScopeEnum `json:"scope,omitempty"`

	SubCategory *SubCategoryEnum `json:"subCategory,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *LicenseTypeEnum `json:"type,omitempty"`
}

func NewLicenseProjection() *LicenseProjection {
	p := new(LicenseProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LicenseProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are Prism,AOS,Calm and others
*/
type LicenseTypeEnum int

const (
	LICENSETYPEENUM_UNKNOWN         LicenseTypeEnum = 0
	LICENSETYPEENUM_REDACTED        LicenseTypeEnum = 1
	LICENSETYPEENUM_PRISM           LicenseTypeEnum = 2
	LICENSETYPEENUM_CALM            LicenseTypeEnum = 3
	LICENSETYPEENUM_FLOW            LicenseTypeEnum = 4
	LICENSETYPEENUM_OBJECT          LicenseTypeEnum = 5
	LICENSETYPEENUM_AOS             LicenseTypeEnum = 6
	LICENSETYPEENUM_FILE            LicenseTypeEnum = 7
	LICENSETYPEENUM_VDI             LicenseTypeEnum = 8
	LICENSETYPEENUM_ROBO            LicenseTypeEnum = 9
	LICENSETYPEENUM_MINE            LicenseTypeEnum = 10
	LICENSETYPEENUM_NCI             LicenseTypeEnum = 11
	LICENSETYPEENUM_NCM             LicenseTypeEnum = 12
	LICENSETYPEENUM_NCI_D           LicenseTypeEnum = 13
	LICENSETYPEENUM_NDA_PLATFORM    LicenseTypeEnum = 14
	LICENSETYPEENUM_UNIFIED_STORAGE LicenseTypeEnum = 15
	LICENSETYPEENUM_EUC             LicenseTypeEnum = 16
	LICENSETYPEENUM_OBJECTS         LicenseTypeEnum = 17
	LICENSETYPEENUM_ERA             LicenseTypeEnum = 18
	LICENSETYPEENUM_DRS             LicenseTypeEnum = 19
	LICENSETYPEENUM_NDS             LicenseTypeEnum = 20
	LICENSETYPEENUM_NDA             LicenseTypeEnum = 21
	LICENSETYPEENUM_NDS_PLATFORM    LicenseTypeEnum = 22
	LICENSETYPEENUM_NDB_PLATFORM    LicenseTypeEnum = 23
	LICENSETYPEENUM_NUS             LicenseTypeEnum = 24
	LICENSETYPEENUM_NDB             LicenseTypeEnum = 25
	LICENSETYPEENUM_NCM_CLOUD       LicenseTypeEnum = 26
	LICENSETYPEENUM_EDGE            LicenseTypeEnum = 27
	LICENSETYPEENUM_NO_LICENSE      LicenseTypeEnum = 28
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LicenseTypeEnum) name(index int) string {
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
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LicenseTypeEnum) GetName() string {
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
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LicenseTypeEnum) index(name string) LicenseTypeEnum {
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
	}
	for idx := range names {
		if names[idx] == name {
			return LicenseTypeEnum(idx)
		}
	}
	return LICENSETYPEENUM_UNKNOWN
}

func (e *LicenseTypeEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LicenseTypeEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LicenseTypeEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LicenseTypeEnum) Ref() *LicenseTypeEnum {
	return &e
}

/*
REST response for all response codes in API path /licensing/v4.0.a1/config/allowances Get operation
*/
type ListClusterAllowancesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClusterAllowancesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListClusterAllowancesApiResponse() *ListClusterAllowancesApiResponse {
	p := new(ListClusterAllowancesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListClusterAllowancesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClusterAllowancesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClusterAllowancesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClusterAllowancesApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/entitlements Get operation
*/
type ListClusterEntitlementsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClusterEntitlementsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListClusterEntitlementsApiResponse() *ListClusterEntitlementsApiResponse {
	p := new(ListClusterEntitlementsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListClusterEntitlementsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClusterEntitlementsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClusterEntitlementsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClusterEntitlementsApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/violations Get operation
*/
type ListClusterViolationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClusterViolationsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListClusterViolationsApiResponse() *ListClusterViolationsApiResponse {
	p := new(ListClusterViolationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListClusterViolationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClusterViolationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClusterViolationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClusterViolationsApiResponseData()
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
REST response for all response codes in API path /licensing/v4.0.a1/config/compliances Get operation
*/
type ListCompliancesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCompliancesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCompliancesApiResponse() *ListCompliancesApiResponse {
	p := new(ListCompliancesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListCompliancesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
REST response for all response codes in API path /licensing/v4.0.a1/config/features Get operation
*/
type ListFeaturesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListFeaturesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListFeaturesApiResponse() *ListFeaturesApiResponse {
	p := new(ListFeaturesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListFeaturesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
REST response for all response codes in API path /licensing/v4.0.a1/config/licenses Get operation
*/
type ListLicensesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLicensesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListLicensesApiResponse() *ListLicensesApiResponse {
	p := new(ListLicensesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListLicensesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
REST response for all response codes in API path /licensing/v4.0.a1/config/settings Get operation
*/
type ListSettingsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSettingsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSettingsApiResponse() *ListSettingsApiResponse {
	p := new(ListSettingsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ListSettingsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
Model representing license and cluster logical versions
*/
type LogicalVersion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates cluster logical version. This is used to keep cluster and license portal in sync
	*/
	Cluster *int `json:"cluster,omitempty"`
	/*
	  Indicates license logical version. This is also used to keep cluster and license portal in sync
	*/
	License *int `json:"license,omitempty"`
}

func NewLogicalVersion() *LogicalVersion {
	p := new(LogicalVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.LogicalVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values of meter are CORES,NODE,TIB,FLASH,USERS,VI,VM_PACKS & VM.
*/
type MeterEnum int

const (
	METERENUM_UNKNOWN  MeterEnum = 0
	METERENUM_REDACTED MeterEnum = 1
	METERENUM_CORES    MeterEnum = 2
	METERENUM_NODE     MeterEnum = 3
	METERENUM_TIB      MeterEnum = 4
	METERENUM_FLASH    MeterEnum = 5
	METERENUM_USERS    MeterEnum = 6
	METERENUM_VI       MeterEnum = 7
	METERENUM_VM_PACKS MeterEnum = 8
	METERENUM_VM       MeterEnum = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MeterEnum) name(index int) string {
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
func (e MeterEnum) GetName() string {
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
func (e *MeterEnum) index(name string) MeterEnum {
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
			return MeterEnum(idx)
		}
	}
	return METERENUM_UNKNOWN
}

func (e *MeterEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MeterEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MeterEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MeterEnum) Ref() *MeterEnum {
	return &e
}

/*
Model for capturing nutanix cluster configuration
*/
type NutanixClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute for capturing billing plan
	*/
	BillingPlan *string `json:"billingPlan,omitempty"`
	/*
	  Attribute for nutanix cluster configuration category like nutanix_cluster
	*/
	Category *string `json:"category,omitempty"`
	/*
	  Indicates consumption model like Subscription
	*/
	Consumption *string `json:"consumption,omitempty"`
	/*
	  Attribute for nutanix cluster configuration identifier
	*/
	Id *string `json:"id,omitempty"`
	/*
	  Indicates whether pulse data collection is required
	*/
	IsPulseRequired *bool `json:"isPulseRequired,omitempty"`
}

func NewNutanixClusterConfig() *NutanixClusterConfig {
	p := new(NutanixClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.NutanixClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are MSP and others
*/
type PostPaidCategoryEnum int

const (
	POSTPAIDCATEGORYENUM_UNKNOWN  PostPaidCategoryEnum = 0
	POSTPAIDCATEGORYENUM_REDACTED PostPaidCategoryEnum = 1
	POSTPAIDCATEGORYENUM_MSP      PostPaidCategoryEnum = 2
	POSTPAIDCATEGORYENUM_NONE     PostPaidCategoryEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PostPaidCategoryEnum) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PostPaidCategoryEnum) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NONE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PostPaidCategoryEnum) index(name string) PostPaidCategoryEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MSP",
		"NONE",
	}
	for idx := range names {
		if names[idx] == name {
			return PostPaidCategoryEnum(idx)
		}
	}
	return POSTPAIDCATEGORYENUM_UNKNOWN
}

func (e *PostPaidCategoryEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PostPaidCategoryEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PostPaidCategoryEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PostPaidCategoryEnum) Ref() *PostPaidCategoryEnum {
	return &e
}

/*
Model for capturing post paid configuration
*/
type PostPaidConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Captures postpaid cluster category like msp
	*/
	Category *string `json:"category,omitempty"`
	/*
	  Attribute for post paid configuration identifier
	*/
	Id *string `json:"id,omitempty"`
	/*
	  Indicates whether pulse data collection is required
	*/
	IsPulseRequired *bool `json:"isPulseRequired,omitempty"`
}

func NewPostPaidConfig() *PostPaidConfig {
	p := new(PostPaidConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.PostPaidConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are PRISM, CALM, AOS and others.
*/
type ProductNameEnum int

const (
	PRODUCTNAMEENUM_UNKNOWN      ProductNameEnum = 0
	PRODUCTNAMEENUM_REDACTED     ProductNameEnum = 1
	PRODUCTNAMEENUM_PRISM        ProductNameEnum = 2
	PRODUCTNAMEENUM_CALM         ProductNameEnum = 3
	PRODUCTNAMEENUM_FLOW         ProductNameEnum = 4
	PRODUCTNAMEENUM_OBJECT       ProductNameEnum = 5
	PRODUCTNAMEENUM_AOS          ProductNameEnum = 6
	PRODUCTNAMEENUM_FILE         ProductNameEnum = 7
	PRODUCTNAMEENUM_VOLUME_GROUP ProductNameEnum = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProductNameEnum) name(index int) string {
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
func (e ProductNameEnum) GetName() string {
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
func (e *ProductNameEnum) index(name string) ProductNameEnum {
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
			return ProductNameEnum(idx)
		}
	}
	return PRODUCTNAMEENUM_UNKNOWN
}

func (e *ProductNameEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProductNameEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProductNameEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProductNameEnum) Ref() *ProductNameEnum {
	return &e
}

/*
Possible values are PC & ALL
*/
type ResetScopeEnum int

const (
	RESETSCOPEENUM_UNKNOWN  ResetScopeEnum = 0
	RESETSCOPEENUM_REDACTED ResetScopeEnum = 1
	RESETSCOPEENUM_PC       ResetScopeEnum = 2
	RESETSCOPEENUM_ALL      ResetScopeEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResetScopeEnum) name(index int) string {
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
func (e ResetScopeEnum) GetName() string {
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
func (e *ResetScopeEnum) index(name string) ResetScopeEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"ALL",
	}
	for idx := range names {
		if names[idx] == name {
			return ResetScopeEnum(idx)
		}
	}
	return RESETSCOPEENUM_UNKNOWN
}

func (e *ResetScopeEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ResetScopeEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ResetScopeEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ResetScopeEnum) Ref() *ResetScopeEnum {
	return &e
}

/*
Possible values of scope are PC,PE,PC_ONLY & PC_PE.
*/
type ScopeEnum int

const (
	SCOPEENUM_UNKNOWN  ScopeEnum = 0
	SCOPEENUM_REDACTED ScopeEnum = 1
	SCOPEENUM_PC       ScopeEnum = 2
	SCOPEENUM_PE       ScopeEnum = 3
	SCOPEENUM_PC_ONLY  ScopeEnum = 4
	SCOPEENUM_PC_PE    ScopeEnum = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScopeEnum) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_ONLY",
		"PC_PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScopeEnum) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_ONLY",
		"PC_PE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScopeEnum) index(name string) ScopeEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PE",
		"PC_ONLY",
		"PC_PE",
	}
	for idx := range names {
		if names[idx] == name {
			return ScopeEnum(idx)
		}
	}
	return SCOPEENUM_UNKNOWN
}

func (e *ScopeEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScopeEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScopeEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScopeEnum) Ref() *ScopeEnum {
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
	EnforcementActions []EnforcementActionsEnum `json:"enforcementActions,omitempty"`

	EnforcementLevel *EnforcementLevelEnum `json:"enforcementLevel,omitempty"`
	/*
	  Attribute for capturing the whether the service is compliant
	*/
	IsCompliant *bool `json:"isCompliant,omitempty"`

	LicenseType *LicenseTypeEnum `json:"licenseType,omitempty"`

	Name *ProductNameEnum `json:"name,omitempty"`
	/*
	  Attribute for capturing the list of violations corresponding to this service.
	*/
	Violations []ServiceViolation `json:"violations,omitempty"`
}

func NewService() *Service {
	p := new(Service)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.Service"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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
	EnforcementActions []EnforcementActionsEnum `json:"enforcementActions,omitempty"`

	EnforcementLevel *EnforcementLevelEnum `json:"enforcementLevel,omitempty"`
	/*
	  Attribute for capturing the whether the service is compliant
	*/
	IsCompliant *bool `json:"isCompliant,omitempty"`

	LicenseType *LicenseTypeEnum `json:"licenseType,omitempty"`

	Name *ProductNameEnum `json:"name,omitempty"`
	/*
	  Attribute for capturing the list of violations corresponding to this service.
	*/
	Violations []ServiceViolation `json:"violations,omitempty"`
}

func NewServiceProjection() *ServiceProjection {
	p := new(ServiceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ServiceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
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

	Type *ServiceViolationEnum `json:"type,omitempty"`
}

func NewServiceViolation() *ServiceViolation {
	p := new(ServiceViolation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "licensing.v4.config.ServiceViolation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are LICENSE-EXPIRED, FEATURE-VIOLATION and CAPACITY-VIOLATION.
*/
type ServiceViolationEnum int

const (
	SERVICEVIOLATIONENUM_UNKNOWN            ServiceViolationEnum = 0
	SERVICEVIOLATIONENUM_REDACTED           ServiceViolationEnum = 1
	SERVICEVIOLATIONENUM_LICENSE_EXPIRED    ServiceViolationEnum = 2
	SERVICEVIOLATIONENUM_FEATURE_VIOLATION  ServiceViolationEnum = 3
	SERVICEVIOLATIONENUM_CAPACITY_VIOLATION ServiceViolationEnum = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ServiceViolationEnum) name(index int) string {
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
func (e ServiceViolationEnum) GetName() string {
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
func (e *ServiceViolationEnum) index(name string) ServiceViolationEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LICENSE_EXPIRED",
		"FEATURE_VIOLATION",
		"CAPACITY_VIOLATION",
	}
	for idx := range names {
		if names[idx] == name {
			return ServiceViolationEnum(idx)
		}
	}
	return SERVICEVIOLATIONENUM_UNKNOWN
}

func (e *ServiceViolationEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ServiceViolationEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ServiceViolationEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ServiceViolationEnum) Ref() *ServiceViolationEnum {
	return &e
}

/*
Response object containing setting details
*/
type Setting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EnforcementPolicy *EnforcementPolicyEnum `json:"enforcementPolicy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  True value indicates that user is allowed to use non compliant features
	*/
	HasNonCompliantFeatures *bool `json:"hasNonCompliantFeatures,omitempty"`
	/*
	  Used to enable and disable NCC license checks
	*/
	IsLicenseCheckDisabled *bool `json:"isLicenseCheckDisabled,omitempty"`
	/*
	  True value indicates cluster UUID represents Prism Central otherwise Prism Element
	*/
	IsMulticluster *bool `json:"isMulticluster,omitempty"`
	/*
	  True value indicates that cluster summary file is downloaded but license summary file is not uplaoded
	*/
	IsStandby *bool `json:"isStandby,omitempty"`
	/*
	  Indicates whether ultimate trial is active or not
	*/
	IsUltimateTrialEnded *bool `json:"isUltimateTrialEnded,omitempty"`
	/*
	  Attribute for capturing license class value
	*/
	LicenseClass *string `json:"licenseClass,omitempty"`
	/*
	  License key applied to dark site clusters
	*/
	LicenseKey *string `json:"licenseKey,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	LogicalVersion *LogicalVersion `json:"logicalVersion,omitempty"`

	NutanixClusterConfig *NutanixClusterConfig `json:"nutanixClusterConfig,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible values are REGISTERED & UNREGISTERED
*/
type StatusEnum int

const (
	STATUSENUM_UNKNOWN      StatusEnum = 0
	STATUSENUM_REDACTED     StatusEnum = 1
	STATUSENUM_REGISTERED   StatusEnum = 2
	STATUSENUM_UNREGISTERED StatusEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *StatusEnum) name(index int) string {
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
func (e StatusEnum) GetName() string {
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
func (e *StatusEnum) index(name string) StatusEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REGISTERED",
		"UNREGISTERED",
	}
	for idx := range names {
		if names[idx] == name {
			return StatusEnum(idx)
		}
	}
	return STATUSENUM_UNKNOWN
}

func (e *StatusEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StatusEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StatusEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StatusEnum) Ref() *StatusEnum {
	return &e
}

/*
Possible values are ADDON, UNLIMITED_CAPACITY
*/
type SubCategoryEnum int

const (
	SUBCATEGORYENUM_UNKNOWN            SubCategoryEnum = 0
	SUBCATEGORYENUM_REDACTED           SubCategoryEnum = 1
	SUBCATEGORYENUM_ADDON              SubCategoryEnum = 2
	SUBCATEGORYENUM_UNLIMITED_CAPACITY SubCategoryEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SubCategoryEnum) name(index int) string {
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
func (e SubCategoryEnum) GetName() string {
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
func (e *SubCategoryEnum) index(name string) SubCategoryEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ADDON",
		"UNLIMITED_CAPACITY",
	}
	for idx := range names {
		if names[idx] == name {
			return SubCategoryEnum(idx)
		}
	}
	return SUBCATEGORYENUM_UNKNOWN
}

func (e *SubCategoryEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SubCategoryEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SubCategoryEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SubCategoryEnum) Ref() *SubCategoryEnum {
	return &e
}

/*
Possible values of type are BOOLEAN & INTEGER
*/
type TypeEnum int

const (
	TYPEENUM_UNKNOWN  TypeEnum = 0
	TYPEENUM_REDACTED TypeEnum = 1
	TYPEENUM_BOOLEAN  TypeEnum = 2
	TYPEENUM_INTEGER  TypeEnum = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TypeEnum) name(index int) string {
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
func (e TypeEnum) GetName() string {
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
func (e *TypeEnum) index(name string) TypeEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOLEAN",
		"INTEGER",
	}
	for idx := range names {
		if names[idx] == name {
			return TypeEnum(idx)
		}
	}
	return TYPEENUM_UNKNOWN
}

func (e *TypeEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TypeEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TypeEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TypeEnum) Ref() *TypeEnum {
	return &e
}

type OneOfListClusterAllowancesApiResponseData struct {
	Discriminator  *string                      `json:"-"`
	ObjectType_    *string                      `json:"-"`
	oneOfType21012 []ClusterAllowance           `json:"-"`
	oneOfType400   *import2.ErrorResponse       `json:"-"`
	oneOfType401   []ClusterAllowanceProjection `json:"-"`
}

func NewOneOfListClusterAllowancesApiResponseData() *OneOfListClusterAllowancesApiResponseData {
	p := new(OneOfListClusterAllowancesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClusterAllowancesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClusterAllowancesApiResponseData is nil"))
	}
	switch v.(type) {
	case []ClusterAllowance:
		p.oneOfType21012 = v.([]ClusterAllowance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ClusterAllowance>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ClusterAllowance>"
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
	case []ClusterAllowanceProjection:
		p.oneOfType401 = v.([]ClusterAllowanceProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ClusterAllowanceProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ClusterAllowanceProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListClusterAllowancesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.ClusterAllowance>" == *p.Discriminator {
		return p.oneOfType21012
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.ClusterAllowanceProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListClusterAllowancesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21012 := new([]ClusterAllowance)
	if err := json.Unmarshal(b, vOneOfType21012); err == nil {

		if len(*vOneOfType21012) == 0 || "licensing.v4.config.ClusterAllowance" == *((*vOneOfType21012)[0].ObjectType_) {
			p.oneOfType21012 = *vOneOfType21012
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ClusterAllowance>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ClusterAllowance>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]ClusterAllowanceProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {

		if len(*vOneOfType401) == 0 || "licensing.v4.config.ClusterAllowanceProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ClusterAllowanceProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ClusterAllowanceProjection>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClusterAllowancesApiResponseData"))
}

func (p *OneOfListClusterAllowancesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.ClusterAllowance>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21012)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.ClusterAllowanceProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListClusterAllowancesApiResponseData")
}

type OneOfGetSettingApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21006 *Setting               `json:"-"`
}

func NewOneOfGetSettingApiResponseData() *OneOfGetSettingApiResponseData {
	p := new(OneOfGetSettingApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSettingApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSettingApiResponseData is nil"))
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
	case Setting:
		if nil == p.oneOfType21006 {
			p.oneOfType21006 = new(Setting)
		}
		*p.oneOfType21006 = v.(Setting)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21006.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21006.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSettingApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType21006 != nil && *p.oneOfType21006.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21006
	}
	return nil
}

func (p *OneOfGetSettingApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType21006 := new(Setting)
	if err := json.Unmarshal(b, vOneOfType21006); err == nil {
		if "licensing.v4.config.Setting" == *vOneOfType21006.ObjectType_ {
			if nil == p.oneOfType21006 {
				p.oneOfType21006 = new(Setting)
			}
			*p.oneOfType21006 = *vOneOfType21006
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21006.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21006.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSettingApiResponseData"))
}

func (p *OneOfGetSettingApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType21006 != nil && *p.oneOfType21006.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21006)
	}
	return nil, errors.New("No value to marshal for OneOfGetSettingApiResponseData")
}

type OneOfGetComplianceApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21016 *Compliance            `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetComplianceApiResponseData() *OneOfGetComplianceApiResponseData {
	p := new(OneOfGetComplianceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetComplianceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetComplianceApiResponseData is nil"))
	}
	switch v.(type) {
	case Compliance:
		if nil == p.oneOfType21016 {
			p.oneOfType21016 = new(Compliance)
		}
		*p.oneOfType21016 = v.(Compliance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21016.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21016.ObjectType_
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

func (p *OneOfGetComplianceApiResponseData) GetValue() interface{} {
	if p.oneOfType21016 != nil && *p.oneOfType21016.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21016
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetComplianceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21016 := new(Compliance)
	if err := json.Unmarshal(b, vOneOfType21016); err == nil {
		if "licensing.v4.config.Compliance" == *vOneOfType21016.ObjectType_ {
			if nil == p.oneOfType21016 {
				p.oneOfType21016 = new(Compliance)
			}
			*p.oneOfType21016 = *vOneOfType21016
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21016.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21016.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetComplianceApiResponseData"))
}

func (p *OneOfGetComplianceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType21016 != nil && *p.oneOfType21016.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21016)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetComplianceApiResponseData")
}

type OneOfGetLicenseApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType21004 *License               `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetLicenseApiResponseData() *OneOfGetLicenseApiResponseData {
	p := new(OneOfGetLicenseApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLicenseApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLicenseApiResponseData is nil"))
	}
	switch v.(type) {
	case License:
		if nil == p.oneOfType21004 {
			p.oneOfType21004 = new(License)
		}
		*p.oneOfType21004 = v.(License)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21004.ObjectType_
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

func (p *OneOfGetLicenseApiResponseData) GetValue() interface{} {
	if p.oneOfType21004 != nil && *p.oneOfType21004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21004
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetLicenseApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType21004 := new(License)
	if err := json.Unmarshal(b, vOneOfType21004); err == nil {
		if "licensing.v4.config.License" == *vOneOfType21004.ObjectType_ {
			if nil == p.oneOfType21004 {
				p.oneOfType21004 = new(License)
			}
			*p.oneOfType21004 = *vOneOfType21004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21004.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21004.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLicenseApiResponseData"))
}

func (p *OneOfGetLicenseApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType21004 != nil && *p.oneOfType21004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21004)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetLicenseApiResponseData")
}

type OneOfListLicensesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType401   []LicenseProjection    `json:"-"`
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.LicenseProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<licensing.v4.config.License>" == *p.Discriminator {
		return p.oneOfType21003
	}
	return nil
}

func (p *OneOfListLicensesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.LicenseProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<licensing.v4.config.License>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21003)
	}
	return nil, errors.New("No value to marshal for OneOfListLicensesApiResponseData")
}

type OneOfGetClusterViolationApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21009 *ClusterViolation      `json:"-"`
}

func NewOneOfGetClusterViolationApiResponseData() *OneOfGetClusterViolationApiResponseData {
	p := new(OneOfGetClusterViolationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterViolationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterViolationApiResponseData is nil"))
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
	case ClusterViolation:
		if nil == p.oneOfType21009 {
			p.oneOfType21009 = new(ClusterViolation)
		}
		*p.oneOfType21009 = v.(ClusterViolation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21009.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21009.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetClusterViolationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType21009 != nil && *p.oneOfType21009.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21009
	}
	return nil
}

func (p *OneOfGetClusterViolationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType21009 := new(ClusterViolation)
	if err := json.Unmarshal(b, vOneOfType21009); err == nil {
		if "licensing.v4.config.ClusterViolation" == *vOneOfType21009.ObjectType_ {
			if nil == p.oneOfType21009 {
				p.oneOfType21009 = new(ClusterViolation)
			}
			*p.oneOfType21009 = *vOneOfType21009
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21009.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21009.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterViolationApiResponseData"))
}

func (p *OneOfGetClusterViolationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType21009 != nil && *p.oneOfType21009.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21009)
	}
	return nil, errors.New("No value to marshal for OneOfGetClusterViolationApiResponseData")
}

type OneOfListClusterEntitlementsApiResponseData struct {
	Discriminator  *string                        `json:"-"`
	ObjectType_    *string                        `json:"-"`
	oneOfType400   *import2.ErrorResponse         `json:"-"`
	oneOfType401   []ClusterEntitlementProjection `json:"-"`
	oneOfType21010 []ClusterEntitlement           `json:"-"`
}

func NewOneOfListClusterEntitlementsApiResponseData() *OneOfListClusterEntitlementsApiResponseData {
	p := new(OneOfListClusterEntitlementsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClusterEntitlementsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClusterEntitlementsApiResponseData is nil"))
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
	case []ClusterEntitlementProjection:
		p.oneOfType401 = v.([]ClusterEntitlementProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ClusterEntitlementProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ClusterEntitlementProjection>"
	case []ClusterEntitlement:
		p.oneOfType21010 = v.([]ClusterEntitlement)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ClusterEntitlement>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ClusterEntitlement>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListClusterEntitlementsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.ClusterEntitlementProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<licensing.v4.config.ClusterEntitlement>" == *p.Discriminator {
		return p.oneOfType21010
	}
	return nil
}

func (p *OneOfListClusterEntitlementsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]ClusterEntitlementProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {

		if len(*vOneOfType401) == 0 || "licensing.v4.config.ClusterEntitlementProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ClusterEntitlementProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ClusterEntitlementProjection>"
			return nil

		}
	}
	vOneOfType21010 := new([]ClusterEntitlement)
	if err := json.Unmarshal(b, vOneOfType21010); err == nil {

		if len(*vOneOfType21010) == 0 || "licensing.v4.config.ClusterEntitlement" == *((*vOneOfType21010)[0].ObjectType_) {
			p.oneOfType21010 = *vOneOfType21010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ClusterEntitlement>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ClusterEntitlement>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClusterEntitlementsApiResponseData"))
}

func (p *OneOfListClusterEntitlementsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.ClusterEntitlementProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<licensing.v4.config.ClusterEntitlement>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21010)
	}
	return nil, errors.New("No value to marshal for OneOfListClusterEntitlementsApiResponseData")
}

type OneOfGetClusterEntitlementApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21011 *ClusterEntitlement    `json:"-"`
}

func NewOneOfGetClusterEntitlementApiResponseData() *OneOfGetClusterEntitlementApiResponseData {
	p := new(OneOfGetClusterEntitlementApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterEntitlementApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterEntitlementApiResponseData is nil"))
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
	case ClusterEntitlement:
		if nil == p.oneOfType21011 {
			p.oneOfType21011 = new(ClusterEntitlement)
		}
		*p.oneOfType21011 = v.(ClusterEntitlement)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21011.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21011.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetClusterEntitlementApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType21011 != nil && *p.oneOfType21011.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21011
	}
	return nil
}

func (p *OneOfGetClusterEntitlementApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType21011 := new(ClusterEntitlement)
	if err := json.Unmarshal(b, vOneOfType21011); err == nil {
		if "licensing.v4.config.ClusterEntitlement" == *vOneOfType21011.ObjectType_ {
			if nil == p.oneOfType21011 {
				p.oneOfType21011 = new(ClusterEntitlement)
			}
			*p.oneOfType21011 = *vOneOfType21011
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21011.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21011.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterEntitlementApiResponseData"))
}

func (p *OneOfGetClusterEntitlementApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType21011 != nil && *p.oneOfType21011.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21011)
	}
	return nil, errors.New("No value to marshal for OneOfGetClusterEntitlementApiResponseData")
}

type OneOfListClusterViolationsApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21008 []ClusterViolation     `json:"-"`
}

func NewOneOfListClusterViolationsApiResponseData() *OneOfListClusterViolationsApiResponseData {
	p := new(OneOfListClusterViolationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClusterViolationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClusterViolationsApiResponseData is nil"))
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
	case []ClusterViolation:
		p.oneOfType21008 = v.([]ClusterViolation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<licensing.v4.config.ClusterViolation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<licensing.v4.config.ClusterViolation>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListClusterViolationsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.ClusterViolation>" == *p.Discriminator {
		return p.oneOfType21008
	}
	return nil
}

func (p *OneOfListClusterViolationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType21008 := new([]ClusterViolation)
	if err := json.Unmarshal(b, vOneOfType21008); err == nil {

		if len(*vOneOfType21008) == 0 || "licensing.v4.config.ClusterViolation" == *((*vOneOfType21008)[0].ObjectType_) {
			p.oneOfType21008 = *vOneOfType21008
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<licensing.v4.config.ClusterViolation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<licensing.v4.config.ClusterViolation>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClusterViolationsApiResponseData"))
}

func (p *OneOfListClusterViolationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.ClusterViolation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21008)
	}
	return nil, errors.New("No value to marshal for OneOfListClusterViolationsApiResponseData")
}

type OneOfListCompliancesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType401   []ComplianceProjection `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21015 []Compliance           `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCompliancesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.ComplianceProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Compliance>" == *p.Discriminator {
		return p.oneOfType21015
	}
	return nil
}

func (p *OneOfListCompliancesApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCompliancesApiResponseData"))
}

func (p *OneOfListCompliancesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.ComplianceProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Compliance>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21015)
	}
	return nil, errors.New("No value to marshal for OneOfListCompliancesApiResponseData")
}

type OneOfListFeaturesApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType401   []FeatureProjection    `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21003 []Feature              `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListFeaturesApiResponseData) GetValue() interface{} {
	if "List<licensing.v4.config.FeatureProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Feature>" == *p.Discriminator {
		return p.oneOfType21003
	}
	return nil
}

func (p *OneOfListFeaturesApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListFeaturesApiResponseData"))
}

func (p *OneOfListFeaturesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<licensing.v4.config.FeatureProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Feature>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21003)
	}
	return nil, errors.New("No value to marshal for OneOfListFeaturesApiResponseData")
}

type OneOfListSettingsApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21005 []Setting              `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSettingsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<licensing.v4.config.Setting>" == *p.Discriminator {
		return p.oneOfType21005
	}
	return nil
}

func (p *OneOfListSettingsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSettingsApiResponseData"))
}

func (p *OneOfListSettingsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<licensing.v4.config.Setting>" == *p.Discriminator {
		return json.Marshal(p.oneOfType21005)
	}
	return nil, errors.New("No value to marshal for OneOfListSettingsApiResponseData")
}

type OneOfGetClusterAllowanceApiResponseData struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType400   *import2.ErrorResponse `json:"-"`
	oneOfType21013 *ClusterAllowance      `json:"-"`
}

func NewOneOfGetClusterAllowanceApiResponseData() *OneOfGetClusterAllowanceApiResponseData {
	p := new(OneOfGetClusterAllowanceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterAllowanceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterAllowanceApiResponseData is nil"))
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
	case ClusterAllowance:
		if nil == p.oneOfType21013 {
			p.oneOfType21013 = new(ClusterAllowance)
		}
		*p.oneOfType21013 = v.(ClusterAllowance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType21013.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType21013.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetClusterAllowanceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType21013 != nil && *p.oneOfType21013.ObjectType_ == *p.Discriminator {
		return *p.oneOfType21013
	}
	return nil
}

func (p *OneOfGetClusterAllowanceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "licensing.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType21013 := new(ClusterAllowance)
	if err := json.Unmarshal(b, vOneOfType21013); err == nil {
		if "licensing.v4.config.ClusterAllowance" == *vOneOfType21013.ObjectType_ {
			if nil == p.oneOfType21013 {
				p.oneOfType21013 = new(ClusterAllowance)
			}
			*p.oneOfType21013 = *vOneOfType21013
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType21013.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType21013.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterAllowanceApiResponseData"))
}

func (p *OneOfGetClusterAllowanceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType21013 != nil && *p.oneOfType21013.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType21013)
	}
	return nil, errors.New("No value to marshal for OneOfGetClusterAllowanceApiResponseData")
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
