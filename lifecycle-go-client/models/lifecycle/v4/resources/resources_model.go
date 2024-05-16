/*
 * Generated file models/lifecycle/v4/resources/resources_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Lifecycle Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Lifecycle resources such as reccomendations and notifications.
*/
package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/error"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/prism/v4/config"
	"time"
)

/*
LCM framework version present in the LCM URL.
*/
type AvailableVersion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Available version UUID.
	*/
	AvailableVersionUuid *string `json:"availableVersionUuid,omitempty"`
	/*
	  Component information for the payload based entity.
	*/
	ChildEntities []string `json:"childEntities,omitempty"`
	/*
	  Update custom messages other than release notes to the user about this available version.
	*/
	CustomMessage *string `json:"customMessage,omitempty"`
	/*
	  List of dependencies for the available version.
	*/
	Dependencies []DependentEntity `json:"dependencies,omitempty"`
	/*
	  Reason for disabling the available version.
	*/
	DisablementReason *string `json:"disablementReason,omitempty"`
	/*
	  UUID of the group that this LCM entity is part of.
	*/
	GroupUuid *string `json:"groupUuid,omitempty"`
	/*
	  Indicates if the available update is enabled.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Order of this available version (1 being the lowest and 6 being the highest) when multiple versions are present with different status.
	*/
	Order *int64 `json:"order,omitempty"`
	/*
	  Release date for the entities that need this information.
	*/
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
	/*
	  Release notes corresponding to the update.
	*/
	ReleaseNotes *string `json:"releaseNotes,omitempty"`

	Status *import1.AvailableVersionStatus `json:"status,omitempty"`
	/*
	  LCM framework version present in the LCM URL.
	*/
	Version *string `json:"version,omitempty"`
}

func NewAvailableVersion() *AvailableVersion {
	p := new(AvailableVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.AvailableVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

/*
LCM configuration on the cluster.
*/
type Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The scheduled time in "%H:%M" 24-hour format of the next inventory execution. Used when auto_inventory_enabled is set to True. The default schedule time is 03:00(AM).
	*/
	AutoInventorySchedule *string `json:"autoInventorySchedule,omitempty"`

	ConnectivityType *ConnectivityType `json:"connectivityType,omitempty"`
	/*
	  List of entities for which One-Click upgrades are not available.
	*/
	DeprecatedSoftwareEntities []string `json:"deprecatedSoftwareEntities,omitempty"`
	/*
	  User friendly display version of LCM installed on the cluster.
	*/
	DisplayVersion *string `json:"displayVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates if LCM is enabled to auto-upgrade products. The default value is False.
	*/
	HasModuleAutoUpgradeEnabled *bool `json:"hasModuleAutoUpgradeEnabled,omitempty"`
	/*
	  Indicates if the auto inventory operation is enabled. The default value is set to False.
	*/
	IsAutoInventoryEnabled *bool `json:"isAutoInventoryEnabled,omitempty"`
	/*
	  Indicates if the bundle is uploaded or not.
	*/
	IsFrameworkBundleUploaded *bool `json:"isFrameworkBundleUploaded,omitempty"`
	/*
	  Indicates if the LCM URL has HTTPS enabled. The default value is True.
	*/
	IsHttpsEnabled *bool `json:"isHttpsEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of entities for which One-Click upgrades are supported.
	*/
	SupportedSoftwareEntities []string `json:"supportedSoftwareEntities,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  URL of the LCM repository.
	*/
	Url *string `json:"url,omitempty"`
	/*
	  LCM version installed on the cluster.
	*/
	Version *string `json:"version,omitempty"`
}

func NewConfig() *Config {
	p := new(Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasModuleAutoUpgradeEnabled = new(bool)
	*p.HasModuleAutoUpgradeEnabled = false
	p.IsAutoInventoryEnabled = new(bool)
	*p.IsAutoInventoryEnabled = false
	p.IsFrameworkBundleUploaded = new(bool)
	*p.IsFrameworkBundleUploaded = false
	p.IsHttpsEnabled = new(bool)
	*p.IsHttpsEnabled = false

	return p
}

/*
This field indicates whether LCM framework on the cluster is running in connected-site mode or darksite mode.
*/
type ConnectivityType int

const (
	CONNECTIVITYTYPE_UNKNOWN                ConnectivityType = 0
	CONNECTIVITYTYPE_REDACTED               ConnectivityType = 1
	CONNECTIVITYTYPE_CONNECTED_SITE         ConnectivityType = 2
	CONNECTIVITYTYPE_DARKSITE_WEB_SERVER    ConnectivityType = 3
	CONNECTIVITYTYPE_DARKSITE_DIRECT_UPLOAD ConnectivityType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConnectivityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConnectivityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConnectivityType) index(name string) ConnectivityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	for idx := range names {
		if names[idx] == name {
			return ConnectivityType(idx)
		}
	}
	return CONNECTIVITYTYPE_UNKNOWN
}

func (e *ConnectivityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConnectivityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConnectivityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConnectivityType) Ref() *ConnectivityType {
	return &e
}

/*
Dependency of an LCM entity available version.
*/
type DependentEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Information of the dependent entity versions for this available entity.
	*/
	DependentVersions []import3.KVPair `json:"dependentVersions,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewDependentEntity() *DependentEntity {
	p := new(DependentEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.DependentEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of deployable versions based on entity types, versions and its dependencies.
*/
type DeployableVersion struct {
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

	EntityType *import1.EntityType `json:"entityType,omitempty"`
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Entity target version.
	*/
	TargetVersion *string `json:"targetVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UpdateDependencies []import1.EntityUpdateSpec `json:"updateDependencies,omitempty"`
}

func NewDeployableVersion() *DeployableVersion {
	p := new(DeployableVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.DeployableVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of an LCM entity.
*/
type Entity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of available versions for an LCM entity to update.
	*/
	AvailableVersions []AvailableVersion `json:"availableVersions,omitempty"`
	/*
	  Component information for the payload based entity.
	*/
	ChildEntities []string `json:"childEntities,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Unique identifier of an LCM entity e.g. "HDD serial number".
	*/
	DeviceId *string `json:"deviceId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  Description of an LCM entity.
	*/
	EntityDescription *string `json:"entityDescription,omitempty"`
	/*
	  Detailed information for the LCM entity. For example, firmware entities contain additional information about NIC and so on.
	*/
	EntityDetails []import3.KVPair `json:"entityDetails,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  UUID of the group that this LCM entity is part of.
	*/
	GroupUuid *string `json:"groupUuid,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A list of sub-entities applicable to the entity.
	*/
	SubEntities []import1.EntityBaseModel `json:"subEntities,omitempty"`
	/*
	  The requested update version of an LCM entity.
	*/
	TargetVersion *string `json:"targetVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Entity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM framework version information.
*/
type FrameworkVersionInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM framework version present in the LCM URL.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/*
	  Current LCM Version.
	*/
	CurrentVersion *string `json:"currentVersion,omitempty"`
	/*
	  Boolean that indicates if LCM framework update is needed.
	*/
	IsUpdateNeeded *bool `json:"isUpdateNeeded,omitempty"`
}

func NewFrameworkVersionInfo() *FrameworkVersionInfo {
	p := new(FrameworkVersionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.FrameworkVersionInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsUpdateNeeded = new(bool)
	*p.IsUpdateNeeded = false

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/config Get operation
*/
type GetConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetConfigApiResponse() *GetConfigApiResponse {
	p := new(GetConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetConfigApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/entities/{extId} Get operation
*/
type GetEntityByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEntityByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetEntityByIdApiResponse() *GetEntityByIdApiResponse {
	p := new(GetEntityByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetEntityByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEntityByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEntityByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEntityByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/notifications/{extId} Get operation
*/
type GetNotificationsByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNotificationsByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNotificationsByIdApiResponse() *GetNotificationsByIdApiResponse {
	p := new(GetNotificationsByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetNotificationsByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNotificationsByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNotificationsByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNotificationsByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/recommendations/{extId} Get operation
*/
type GetRecommendationByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecommendationByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRecommendationByIdApiResponse() *GetRecommendationByIdApiResponse {
	p := new(GetRecommendationByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetRecommendationByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRecommendationByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRecommendationByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRecommendationByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/status Get operation
*/
type GetStatusApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetStatusApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetStatusApiResponse() *GetStatusApiResponse {
	p := new(GetStatusApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetStatusApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetStatusApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetStatusApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetStatusApiResponseData()
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
LCM image (List of LCM image files).
*/
type Image struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	Files []ImageFile `json:"files"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  Denotes if the thirdparty version is qualified.
	*/
	IsQualified *bool `json:"isQualified,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Release notes for the LCM image.
	*/
	ReleaseNotes *string `json:"releaseNotes,omitempty"`
	/*
	  Version specification for image metadata JSON.
	*/
	SpecVersion *string `json:"specVersion"`

	Status *import1.AvailableVersionStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Image) MarshalJSON() ([]byte, error) {
	type ImageProxy Image
	return json.Marshal(struct {
		*ImageProxy
		Files       []ImageFile `json:"files,omitempty"`
		SpecVersion *string     `json:"specVersion,omitempty"`
	}{
		ImageProxy:  (*ImageProxy)(p),
		Files:       p.Files,
		SpecVersion: p.SpecVersion,
	})
}

func NewImage() *Image {
	p := new(Image)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Image"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsQualified = new(bool)
	*p.IsQualified = false

	return p
}

/*
Description of LCM image file.
*/
type ImageFile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM image checksum.
	*/
	Checksum *string `json:"checksum"`

	ChecksumType *import1.CheckSumType `json:"checksumType"`
	/*
	  Image file global catalog item UUID.
	*/
	FileLocationId *string `json:"fileLocationId,omitempty"`
	/*
	  File path for the LCM image.
	*/
	FilePath *string `json:"filePath,omitempty"`
	/*
	  LCM image file name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  LCM image file size.
	*/
	SizeBytes *int64 `json:"sizeBytes"`
}

func (p *ImageFile) MarshalJSON() ([]byte, error) {
	type ImageFileProxy ImageFile
	return json.Marshal(struct {
		*ImageFileProxy
		Checksum     *string               `json:"checksum,omitempty"`
		ChecksumType *import1.CheckSumType `json:"checksumType,omitempty"`
		SizeBytes    *int64                `json:"sizeBytes,omitempty"`
	}{
		ImageFileProxy: (*ImageFileProxy)(p),
		Checksum:       p.Checksum,
		ChecksumType:   p.ChecksumType,
		SizeBytes:      p.SizeBytes,
	})
}

func NewImageFile() *ImageFile {
	p := new(ImageFile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ImageFile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation type and UUID of an ongoing operation in LCM.
*/
type InProgressOpInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Root task UUID of the operation, if it is in running state.
	*/
	OperationId *string `json:"operationId,omitempty"`

	OperationType *import1.OperationType `json:"operationType,omitempty"`
}

func NewInProgressOpInfo() *InProgressOpInfo {
	p := new(InProgressOpInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.InProgressOpInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/entities Get operation
*/
type ListEntitiesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEntitiesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListEntitiesApiResponse() *ListEntitiesApiResponse {
	p := new(ListEntitiesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListEntitiesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEntitiesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEntitiesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEntitiesApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/images Get operation
*/
type ListImagesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListImagesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListImagesApiResponse() *ListImagesApiResponse {
	p := new(ListImagesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListImagesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListImagesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListImagesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListImagesApiResponseData()
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
The computed LCM upgrade notifications for the given input.
*/
type Notification struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Notifications []NotificationItem `json:"notifications,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewNotification() *Notification {
	p := new(Notification)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Notification"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Detailed LCM upgrade notification information for this entity.
*/
type NotificationDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A message with notification details. The description of the most disruptive action that will occur on the node or the cluster. INFO, WARNING or NOTICE based message for an entity.
	*/
	Message *string `json:"message,omitempty"`

	SeverityLevel *SeverityLevel `json:"severityLevel,omitempty"`
}

func NewNotificationDetail() *NotificationDetail {
	p := new(NotificationDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.NotificationDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM upgrade notification generated for a node or cluster based on specified entity/entities and target version(s).
*/
type NotificationItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of upgrade notifications for this entity.
	*/
	Details []NotificationDetail `json:"details,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
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

	HypervisorType *import1.HypervisorType `json:"hypervisorType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`

	NotificationType *import1.NotificationType `json:"notificationType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Version to upgrade to.
	*/
	ToVersion *string `json:"toVersion,omitempty"`
}

func NewNotificationItem() *NotificationItem {
	p := new(NotificationItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.NotificationItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Upgrade recommendations for LCM entity/entities.
*/
type RecommendationResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of added LCM entities to the input recommendations specification.
	*/
	AddableEntities []UpdatedTargetEntityResult `json:"addableEntities,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  List of deployable entities and their dependencies.
	*/
	DeployableVersions []DeployableVersion `json:"deployableVersions,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []import1.EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of modified LCM entities from the input recommendations specification.
	*/
	ModifiableEntities []UpdatedTargetEntityResult `json:"modifiableEntities,omitempty"`
	/*
	  List of skipped LCM entities from the input recommendations specification.
	*/
	SkippedEntities []UpdatedTargetEntityResult `json:"skippedEntities,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRecommendationResult() *RecommendationResult {
	p := new(RecommendationResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.RecommendationResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification to get upgrade recommendations for specific UUID and target version via LCM Recommendation
*/
type RecommendationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	RecommendationSpecItemDiscriminator_ *string `json:"$recommendationSpecItemDiscriminator,omitempty"`

	RecommendationSpec *OneOfRecommendationSpecRecommendationSpec `json:"recommendationSpec"`
}

func (p *RecommendationSpec) MarshalJSON() ([]byte, error) {
	type RecommendationSpecProxy RecommendationSpec
	return json.Marshal(struct {
		*RecommendationSpecProxy
		RecommendationSpec *OneOfRecommendationSpecRecommendationSpec `json:"recommendationSpec,omitempty"`
	}{
		RecommendationSpecProxy: (*RecommendationSpecProxy)(p),
		RecommendationSpec:      p.RecommendationSpec,
	})
}

func NewRecommendationSpec() *RecommendationSpec {
	p := new(RecommendationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.RecommendationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecommendationSpec) GetRecommendationSpec() interface{} {
	if nil == p.RecommendationSpec {
		return nil
	}
	return p.RecommendationSpec.GetValue()
}

func (p *RecommendationSpec) SetRecommendationSpec(v interface{}) error {
	if nil == p.RecommendationSpec {
		p.RecommendationSpec = NewOneOfRecommendationSpecRecommendationSpec()
	}
	e := p.RecommendationSpec.SetValue(v)
	if nil == e {
		if nil == p.RecommendationSpecItemDiscriminator_ {
			p.RecommendationSpecItemDiscriminator_ = new(string)
		}
		*p.RecommendationSpecItemDiscriminator_ = *p.RecommendationSpec.Discriminator
	}
	return e
}

/*
Severity level of LCM notification.
*/
type SeverityLevel int

const (
	SEVERITYLEVEL_UNKNOWN  SeverityLevel = 0
	SEVERITYLEVEL_REDACTED SeverityLevel = 1
	SEVERITYLEVEL_WARNING  SeverityLevel = 2
	SEVERITYLEVEL_NOTICE   SeverityLevel = 3
	SEVERITYLEVEL_INFO     SeverityLevel = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SeverityLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SeverityLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SeverityLevel) index(name string) SeverityLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	for idx := range names {
		if names[idx] == name {
			return SeverityLevel(idx)
		}
	}
	return SEVERITYLEVEL_UNKNOWN
}

func (e *SeverityLevel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SeverityLevel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SeverityLevel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SeverityLevel) Ref() *SeverityLevel {
	return &e
}

/*
Detailed information about the current LCM framework status.
*/
type StatusInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FrameworkVersion *FrameworkVersionInfo `json:"frameworkVersion,omitempty"`

	InProgressOperation *InProgressOpInfo `json:"inProgressOperation,omitempty"`
	/*
	  Boolean that indicates if cancel intent for LCM update is set or not.
	*/
	IsCancelIntentSet *bool `json:"isCancelIntentSet,omitempty"`
	/*
	  Upload task UUID.
	*/
	UploadTaskUuid *string `json:"uploadTaskUuid,omitempty"`
}

func NewStatusInfo() *StatusInfo {
	p := new(StatusInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.StatusInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsCancelIntentSet = new(bool)
	*p.IsCancelIntentSet = false

	return p
}

/*
LCM target entity for which recommendations are requested.
*/
type TargetEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique identifier of an LCM entity e.g. "HDD serial number".
	*/
	DeviceId *string `json:"deviceId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
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
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The requested update version of an LCM entity.
	*/
	Version *string `json:"version"`
}

func (p *TargetEntity) MarshalJSON() ([]byte, error) {
	type TargetEntityProxy TargetEntity
	return json.Marshal(struct {
		*TargetEntityProxy
		Version *string `json:"version,omitempty"`
	}{
		TargetEntityProxy: (*TargetEntityProxy)(p),
		Version:           p.Version,
	})
}

func NewTargetEntity() *TargetEntity {
	p := new(TargetEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.TargetEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.0.b1/resources/config Put operation
*/
type UpdateConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateConfigApiResponse() *UpdateConfigApiResponse {
	p := new(UpdateConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdateConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateConfigApiResponseData()
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
Updated LCM target entity in recommendation result.
*/
type UpdatedTargetEntity struct {
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

	EntityType *import1.EntityType `json:"entityType,omitempty"`
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
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewUpdatedTargetEntity() *UpdatedTargetEntity {
	p := new(UpdatedTargetEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdatedTargetEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM error for target entity
*/
type UpdatedTargetEntityResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Error message for the target entity that failed in the input recommendations specification.
	*/
	Message *string `json:"message"`

	TargetEntity *UpdatedTargetEntity `json:"targetEntity,omitempty"`
}

func (p *UpdatedTargetEntityResult) MarshalJSON() ([]byte, error) {
	type UpdatedTargetEntityResultProxy UpdatedTargetEntityResult
	return json.Marshal(struct {
		*UpdatedTargetEntityResultProxy
		Message *string `json:"message,omitempty"`
	}{
		UpdatedTargetEntityResultProxy: (*UpdatedTargetEntityResultProxy)(p),
		Message:                        p.Message,
	})
}

func NewUpdatedTargetEntityResult() *UpdatedTargetEntityResult {
	p := new(UpdatedTargetEntityResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdatedTargetEntityResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetRecommendationByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *RecommendationResult  `json:"-"`
}

func NewOneOfGetRecommendationByIdApiResponseData() *OneOfGetRecommendationByIdApiResponseData {
	p := new(OneOfGetRecommendationByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRecommendationByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRecommendationByIdApiResponseData is nil"))
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
	case RecommendationResult:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(RecommendationResult)
		}
		*p.oneOfType2001 = v.(RecommendationResult)
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

func (p *OneOfGetRecommendationByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetRecommendationByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(RecommendationResult)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.RecommendationResult" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(RecommendationResult)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecommendationByIdApiResponseData"))
}

func (p *OneOfGetRecommendationByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecommendationByIdApiResponseData")
}

type OneOfGetStatusApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *StatusInfo            `json:"-"`
}

func NewOneOfGetStatusApiResponseData() *OneOfGetStatusApiResponseData {
	p := new(OneOfGetStatusApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetStatusApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetStatusApiResponseData is nil"))
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
	case StatusInfo:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(StatusInfo)
		}
		*p.oneOfType2001 = v.(StatusInfo)
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

func (p *OneOfGetStatusApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetStatusApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(StatusInfo)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.StatusInfo" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(StatusInfo)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetStatusApiResponseData"))
}

func (p *OneOfGetStatusApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetStatusApiResponseData")
}

type OneOfListImagesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Image                `json:"-"`
}

func NewOneOfListImagesApiResponseData() *OneOfListImagesApiResponseData {
	p := new(OneOfListImagesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListImagesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListImagesApiResponseData is nil"))
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
	case []Image:
		p.oneOfType2001 = v.([]Image)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.Image>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.Image>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListImagesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.resources.Image>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListImagesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]Image)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {

		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.Image" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.Image>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.Image>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListImagesApiResponseData"))
}

func (p *OneOfListImagesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.resources.Image>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListImagesApiResponseData")
}

type OneOfRecommendationSpecRecommendationSpec struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2002 []TargetEntity             `json:"-"`
	oneOfType2001 []import1.EntityType       `json:"-"`
	oneOfType2004 []import1.EntityDeploySpec `json:"-"`
	oneOfType2003 []import1.EntityUpdateSpec `json:"-"`
}

func NewOneOfRecommendationSpecRecommendationSpec() *OneOfRecommendationSpecRecommendationSpec {
	p := new(OneOfRecommendationSpecRecommendationSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecommendationSpecRecommendationSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecommendationSpecRecommendationSpec is nil"))
	}
	switch v.(type) {
	case []TargetEntity:
		p.oneOfType2002 = v.([]TargetEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.TargetEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.TargetEntity>"
	case []import1.EntityType:
		p.oneOfType2001 = v.([]import1.EntityType)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityType>"
	case []import1.EntityDeploySpec:
		p.oneOfType2004 = v.([]import1.EntityDeploySpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityDeploySpec>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityDeploySpec>"
	case []import1.EntityUpdateSpec:
		p.oneOfType2003 = v.([]import1.EntityUpdateSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityUpdateSpec>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityUpdateSpec>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRecommendationSpecRecommendationSpec) GetValue() interface{} {
	if "List<lifecycle.v4.resources.TargetEntity>" == *p.Discriminator {
		return p.oneOfType2002
	}
	if "List<lifecycle.v4.common.EntityType>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if "List<lifecycle.v4.common.EntityDeploySpec>" == *p.Discriminator {
		return p.oneOfType2004
	}
	if "List<lifecycle.v4.common.EntityUpdateSpec>" == *p.Discriminator {
		return p.oneOfType2003
	}
	return nil
}

func (p *OneOfRecommendationSpecRecommendationSpec) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new([]TargetEntity)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {

		if len(*vOneOfType2002) == 0 || "lifecycle.v4.resources.TargetEntity" == *((*vOneOfType2002)[0].ObjectType_) {
			p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.TargetEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.TargetEntity>"
			return nil

		}
	}
	vOneOfType2001 := new([]import1.EntityType)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		p.oneOfType2001 = *vOneOfType2001
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityType>"
		return nil
	}
	vOneOfType2004 := new([]import1.EntityDeploySpec)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {

		if len(*vOneOfType2004) == 0 || "lifecycle.v4.common.EntityDeploySpec" == *((*vOneOfType2004)[0].ObjectType_) {
			p.oneOfType2004 = *vOneOfType2004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.common.EntityDeploySpec>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.common.EntityDeploySpec>"
			return nil

		}
	}
	vOneOfType2003 := new([]import1.EntityUpdateSpec)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {

		if len(*vOneOfType2003) == 0 || "lifecycle.v4.common.EntityUpdateSpec" == *((*vOneOfType2003)[0].ObjectType_) {
			p.oneOfType2003 = *vOneOfType2003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.common.EntityUpdateSpec>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.common.EntityUpdateSpec>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecommendationSpecRecommendationSpec"))
}

func (p *OneOfRecommendationSpecRecommendationSpec) MarshalJSON() ([]byte, error) {
	if "List<lifecycle.v4.resources.TargetEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if "List<lifecycle.v4.common.EntityType>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if "List<lifecycle.v4.common.EntityDeploySpec>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	if "List<lifecycle.v4.common.EntityUpdateSpec>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	return nil, errors.New("No value to marshal for OneOfRecommendationSpecRecommendationSpec")
}

type OneOfGetEntityByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *Entity                `json:"-"`
}

func NewOneOfGetEntityByIdApiResponseData() *OneOfGetEntityByIdApiResponseData {
	p := new(OneOfGetEntityByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEntityByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEntityByIdApiResponseData is nil"))
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
	case Entity:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Entity)
		}
		*p.oneOfType2001 = v.(Entity)
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

func (p *OneOfGetEntityByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetEntityByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(Entity)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Entity" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Entity)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEntityByIdApiResponseData"))
}

func (p *OneOfGetEntityByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetEntityByIdApiResponseData")
}

type OneOfGetConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *Config                `json:"-"`
}

func NewOneOfGetConfigApiResponseData() *OneOfGetConfigApiResponseData {
	p := new(OneOfGetConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetConfigApiResponseData is nil"))
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
	case Config:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Config)
		}
		*p.oneOfType2001 = v.(Config)
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

func (p *OneOfGetConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(Config)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Config" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Config)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetConfigApiResponseData"))
}

func (p *OneOfGetConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetConfigApiResponseData")
}

type OneOfGetNotificationsByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Notification          `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetNotificationsByIdApiResponseData() *OneOfGetNotificationsByIdApiResponseData {
	p := new(OneOfGetNotificationsByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNotificationsByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNotificationsByIdApiResponseData is nil"))
	}
	switch v.(type) {
	case Notification:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Notification)
		}
		*p.oneOfType2001 = v.(Notification)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfGetNotificationsByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNotificationsByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Notification)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Notification" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Notification)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNotificationsByIdApiResponseData"))
}

func (p *OneOfGetNotificationsByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNotificationsByIdApiResponseData")
}

type OneOfListEntitiesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Entity               `json:"-"`
}

func NewOneOfListEntitiesApiResponseData() *OneOfListEntitiesApiResponseData {
	p := new(OneOfListEntitiesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEntitiesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEntitiesApiResponseData is nil"))
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
	case []Entity:
		p.oneOfType2001 = v.([]Entity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.Entity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.Entity>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListEntitiesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.resources.Entity>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListEntitiesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]Entity)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {

		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.Entity" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.Entity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.Entity>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEntitiesApiResponseData"))
}

func (p *OneOfListEntitiesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.resources.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListEntitiesApiResponseData")
}

type OneOfUpdateConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
}

func NewOneOfUpdateConfigApiResponseData() *OneOfUpdateConfigApiResponseData {
	p := new(OneOfUpdateConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateConfigApiResponseData is nil"))
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
	case import5.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import5.TaskReference)
		}
		*p.oneOfType2001 = v.(import5.TaskReference)
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

func (p *OneOfUpdateConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(import5.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import5.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateConfigApiResponseData"))
}

func (p *OneOfUpdateConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateConfigApiResponseData")
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
