/*
 * Generated file models/dataprotection/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-4
 *
 * Part of the Nutanix Dataprotection Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Hosts, Clusters and other Infrastructure
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import6 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import5 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/prism/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/storage/v4/config"
	"time"
)

/*
Recovery point restore that overrides the AHV VM configuration. The specified properties will be overridden for the restored VM.
*/
type AhvVmOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the restored VM
	*/
	Name *string `json:"name,omitempty"`
}

func NewAhvVmOverrideSpec() *AhvVmOverrideSpec {
	p := new(AhvVmOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AhvVmOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.AhvVmOverrideSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specifies category in key and value object format.
*/
type Category struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the category key.
	*/
	Key *string `json:"key"`
	/*
	  Name of the category value.
	*/
	Value *string `json:"value"`
}

func (p *Category) MarshalJSON() ([]byte, error) {
	type CategoryProxy Category
	return json.Marshal(struct {
		*CategoryProxy
		Key   *string `json:"key,omitempty"`
		Value *string `json:"value,omitempty"`
	}{
		CategoryProxy: (*CategoryProxy)(p),
		Key:           p.Key,
		Value:         p.Value,
	})
}

func NewCategory() *Category {
	p := new(Category)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Category"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Category"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Changed Regions tracking
*/
type ChangedRegions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The offset from where the client must continue the request. This field will not be set when there are no more changed regions to be returned. Note that the nextOffset can be outside the endOffset specified by the client in the request. This helps clients reach the next changed offset faster.
	*/
	NextOffset *int64 `json:"nextOffset,omitempty"`
	/*
	  Changed Regions list comprising of offset and regions length pertaining to the Disk Recovery Point.
	*/
	Regions []Regions `json:"regions,omitempty"`
}

func NewChangedRegions() *ChangedRegions {
	p := new(ChangedRegions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ChangedRegions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ChangedRegions"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/disk-recovery-points/{diskRecoveryPointExtId}/changed-regions Get operation
*/
type ChangedRegionsListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangedRegionsListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewChangedRegionsListApiResponse() *ChangedRegionsListApiResponse {
	p := new(ChangedRegionsListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ChangedRegionsListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ChangedRegionsListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ChangedRegionsListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ChangedRegionsListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfChangedRegionsListApiResponseData()
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
The UUID corresponding to the cluster where an entity resides or is synced.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ExtId *string `json:"extId"`
	/*
	  Name of the Prism Element cluster, this is a read-only field
	*/
	Name *string `json:"name,omitempty"`
}

func (p *ClusterReference) MarshalJSON() ([]byte, error) {
	type ClusterReferenceProxy ClusterReference
	return json.Marshal(struct {
		*ClusterReferenceProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		ClusterReferenceProxy: (*ClusterReferenceProxy)(p),
		ExtId:                 p.ExtId,
	})
}

func NewClusterReference() *ClusterReference {
	p := new(ClusterReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ClusterReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
There are many scenarios where the state of an application must be captured as an aggregate of the internal state of a set of related entities at some point in time. The Consistency group is a collection of all the entities whose snapshot represents the application state at that point in time
*/
type ConsistencyGroup struct {
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

	Members []ConsistencyGroupMember `json:"members,omitempty"`
	/*
	  Name of the Consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this Consistency group. This is a read-only field that is inserted into the Consistency group entity at the time of Consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroup) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProxy ConsistencyGroup
	return json.Marshal(struct {
		*ConsistencyGroupProxy
		Name *string `json:"name,omitempty"`
	}{
		ConsistencyGroupProxy: (*ConsistencyGroupProxy)(p),
		Name:                  p.Name,
	})
}

func NewConsistencyGroup() *ConsistencyGroup {
	p := new(ConsistencyGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/consistency-groups/{extId} Put operation
*/
type ConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConsistencyGroupApiResponse() *ConsistencyGroupApiResponse {
	p := new(ConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConsistencyGroupApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/consistency-groups Get operation
*/
type ConsistencyGroupListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConsistencyGroupListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConsistencyGroupListApiResponse() *ConsistencyGroupListApiResponse {
	p := new(ConsistencyGroupListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConsistencyGroupListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConsistencyGroupListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConsistencyGroupListApiResponseData()
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
External identifier of the members (VM or volume group) of the Consistency group object.
*/
type ConsistencyGroupMember struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the members (VM or volume group) of the Consistency group object.
	*/
	EntityId *string `json:"entityId,omitempty"`

	EntityType *ConsistencyGroupMemberType `json:"entityType"`
	/*
	  External identifier of the members (VM or volume group) of the Consistency group object.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *ConsistencyGroupMember) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupMemberProxy ConsistencyGroupMember
	return json.Marshal(struct {
		*ConsistencyGroupMemberProxy
		EntityType *ConsistencyGroupMemberType `json:"entityType,omitempty"`
	}{
		ConsistencyGroupMemberProxy: (*ConsistencyGroupMemberProxy)(p),
		EntityType:                  p.EntityType,
	})
}

func NewConsistencyGroupMember() *ConsistencyGroupMember {
	p := new(ConsistencyGroupMember)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMember"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupMember"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Consistency group member type
*/
type ConsistencyGroupMemberType int

const (
	CONSISTENCYGROUPMEMBERTYPE_UNKNOWN      ConsistencyGroupMemberType = 0
	CONSISTENCYGROUPMEMBERTYPE_REDACTED     ConsistencyGroupMemberType = 1
	CONSISTENCYGROUPMEMBERTYPE_VM           ConsistencyGroupMemberType = 2
	CONSISTENCYGROUPMEMBERTYPE_VOLUME_GROUP ConsistencyGroupMemberType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConsistencyGroupMemberType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConsistencyGroupMemberType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConsistencyGroupMemberType) index(name string) ConsistencyGroupMemberType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return ConsistencyGroupMemberType(idx)
		}
	}
	return CONSISTENCYGROUPMEMBERTYPE_UNKNOWN
}

func (e *ConsistencyGroupMemberType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConsistencyGroupMemberType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConsistencyGroupMemberType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConsistencyGroupMemberType) Ref() *ConsistencyGroupMemberType {
	return &e
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/consistency-groups/{extId}/$actions/migrate Post operation
*/
type ConsistencyGroupMigrateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConsistencyGroupMigrateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConsistencyGroupMigrateApiResponse() *ConsistencyGroupMigrateApiResponse {
	p := new(ConsistencyGroupMigrateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMigrateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupMigrateApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConsistencyGroupMigrateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConsistencyGroupMigrateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConsistencyGroupMigrateApiResponseData()
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
Specification for the migrate action on the Consistency group.
*/
type ConsistencyGroupMigrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference to the target Availability Zone where the entities need to be migrated.
	*/
	TargetAvailabilityZoneId *string `json:"targetAvailabilityZoneId"`
	/*
	  Reference to the cluster in the target Availability Zone where the entities need to be migrated.
	*/
	TargetClusterId *string `json:"targetClusterId,omitempty"`
}

func (p *ConsistencyGroupMigrationSpec) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupMigrationSpecProxy ConsistencyGroupMigrationSpec
	return json.Marshal(struct {
		*ConsistencyGroupMigrationSpecProxy
		TargetAvailabilityZoneId *string `json:"targetAvailabilityZoneId,omitempty"`
	}{
		ConsistencyGroupMigrationSpecProxy: (*ConsistencyGroupMigrationSpecProxy)(p),
		TargetAvailabilityZoneId:           p.TargetAvailabilityZoneId,
	})
}

func NewConsistencyGroupMigrationSpec() *ConsistencyGroupMigrationSpec {
	p := new(ConsistencyGroupMigrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMigrationSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupMigrationSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConsistencyGroupProjection struct {
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

	Members []ConsistencyGroupMember `json:"members,omitempty"`
	/*
	  Name of the Consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this Consistency group. This is a read-only field that is inserted into the Consistency group entity at the time of Consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroupProjection) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProjectionProxy ConsistencyGroupProjection
	return json.Marshal(struct {
		*ConsistencyGroupProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		ConsistencyGroupProjectionProxy: (*ConsistencyGroupProjectionProxy)(p),
		Name:                            p.Name,
	})
}

func NewConsistencyGroupProjection() *ConsistencyGroupProjection {
	p := new(ConsistencyGroupProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for Create Consistency group Recovery point
*/
type ConsistencyGroupRecoveryPointSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the Consistency group Recovery point.
	*/
	Name *string `json:"name"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`
}

func (p *ConsistencyGroupRecoveryPointSpec) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupRecoveryPointSpecProxy ConsistencyGroupRecoveryPointSpec
	return json.Marshal(struct {
		*ConsistencyGroupRecoveryPointSpecProxy
		Name *string `json:"name,omitempty"`
	}{
		ConsistencyGroupRecoveryPointSpecProxy: (*ConsistencyGroupRecoveryPointSpecProxy)(p),
		Name:                                   p.Name,
	})
}

func NewConsistencyGroupRecoveryPointSpec() *ConsistencyGroupRecoveryPointSpec {
	p := new(ConsistencyGroupRecoveryPointSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupRecoveryPointSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ConsistencyGroupRecoveryPointSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/consistency-groups/{extId}/$actions/create-recovery-point Post operation
*/
type CreateConsistencyGroupRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateConsistencyGroupRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateConsistencyGroupRecoveryPointApiResponse() *CreateConsistencyGroupRecoveryPointApiResponse {
	p := new(CreateConsistencyGroupRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreateConsistencyGroupRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.CreateConsistencyGroupRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateConsistencyGroupRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateConsistencyGroupRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateConsistencyGroupRecoveryPointApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points Post operation
*/
type CreateRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateRecoveryPointApiResponse() *CreateRecoveryPointApiResponse {
	p := new(CreateRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreateRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.CreateRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateRecoveryPointApiResponseData()
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
The data protection details for the protected resource that are relevant to the specified cluster, like the time ranges available for recovery on the given cluster.
*/
type DataProtectionInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LocationReference *DataProtectionSiteReference `json:"locationReference,omitempty"`

	RecoveryInfo *RecoveryInfo `json:"recoveryInfo,omitempty"`
}

func NewDataProtectionInfo() *DataProtectionInfo {
	p := new(DataProtectionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DataProtectionInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DataProtectionSiteReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the Prism Central.
	*/
	PcExtId *string `json:"pcExtId,omitempty"`
}

func NewDataProtectionSiteReference() *DataProtectionSiteReference {
	p := new(DataProtectionSiteReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionSiteReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DataProtectionSiteReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/consistency-groups/{extId} Delete operation
*/
type DeleteConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteConsistencyGroupApiResponse() *DeleteConsistencyGroupApiResponse {
	p := new(DeleteConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DeleteConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DeleteConsistencyGroupApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteConsistencyGroupApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{extId} Delete operation
*/
type DeleteRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRecoveryPointApiResponse() *DeleteRecoveryPointApiResponse {
	p := new(DeleteRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DeleteRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DeleteRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRecoveryPointApiResponseData()
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

type DisasterRecoveryLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Prism Element clusters references. It is required in the scenario when both primary and secondary Prism Elements are managed by the same Prism Central.
	*/
	ClusterReferences []ClusterReference `json:"clusterReferences,omitempty"`
	/*
	  External identifier of the Prism Central.
	*/
	PcExtId *string `json:"pcExtId"`
}

func (p *DisasterRecoveryLocation) MarshalJSON() ([]byte, error) {
	type DisasterRecoveryLocationProxy DisasterRecoveryLocation
	return json.Marshal(struct {
		*DisasterRecoveryLocationProxy
		PcExtId *string `json:"pcExtId,omitempty"`
	}{
		DisasterRecoveryLocationProxy: (*DisasterRecoveryLocationProxy)(p),
		PcExtId:                       p.PcExtId,
	})
}

func NewDisasterRecoveryLocation() *DisasterRecoveryLocation {
	p := new(DisasterRecoveryLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DisasterRecoveryLocation"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DisasterRecoveryLocation"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The configuration of the Disk captured by the Recovery point.
*/
type Disk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Capacity of the virtual disk in bytes.
	*/
	CapacityBytes *int64 `json:"capacityBytes,omitempty"`
	/*
	  Disk external identifier which is captured as part of this Recovery point.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
}

func NewDisk() *Disk {
	p := new(Disk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Disk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Disk"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents the disk Recovery point properties.
*/
type DiskRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	Disk *Disk `json:"disk,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []import5.VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`
}

func NewDiskRecoveryPoint() *DiskRecoveryPoint {
	p := new(DiskRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DiskRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DiskRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/disk-recovery-points/{diskRecoveryPointExtId} Get operation
*/
type DiskRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDiskRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDiskRecoveryPointApiResponse() *DiskRecoveryPointApiResponse {
	p := new(DiskRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DiskRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DiskRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DiskRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DiskRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDiskRecoveryPointApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/disk-recovery-points Get operation
*/
type DiskRecoveryPointListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDiskRecoveryPointListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDiskRecoveryPointListApiResponse() *DiskRecoveryPointListApiResponse {
	p := new(DiskRecoveryPointListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DiskRecoveryPointListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.DiskRecoveryPointListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DiskRecoveryPointListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DiskRecoveryPointListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDiskRecoveryPointListApiResponseData()
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
Recovery point restore that overrides the ESXi VM configuration. The specified properties will be overridden for the restored VM.
*/
type EsxiVmOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the restored VM
	*/
	Name *string `json:"name,omitempty"`
}

func NewEsxiVmOverrideSpec() *EsxiVmOverrideSpec {
	p := new(EsxiVmOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.EsxiVmOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.EsxiVmOverrideSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Architecture of the hardware where the Recovery point is generated.
*/
type HardwareArchitecture int

const (
	HARDWAREARCHITECTURE_UNKNOWN  HardwareArchitecture = 0
	HARDWAREARCHITECTURE_REDACTED HardwareArchitecture = 1
	HARDWAREARCHITECTURE_X86_64   HardwareArchitecture = 2
	HARDWAREARCHITECTURE_PPC      HardwareArchitecture = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HardwareArchitecture) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"X86_64",
		"PPC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HardwareArchitecture) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"X86_64",
		"PPC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HardwareArchitecture) index(name string) HardwareArchitecture {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"X86_64",
		"PPC",
	}
	for idx := range names {
		if names[idx] == name {
			return HardwareArchitecture(idx)
		}
	}
	return HARDWAREARCHITECTURE_UNKNOWN
}

func (e *HardwareArchitecture) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HardwareArchitecture:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HardwareArchitecture) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HardwareArchitecture) Ref() *HardwareArchitecture {
	return &e
}

/*
Information about the Witness host site.
*/
type HostReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HostType *HostType `json:"hostType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness host site.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.HostReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the Witness host site.
*/
type HostType int

const (
	HOSTTYPE_UNKNOWN       HostType = 0
	HOSTTYPE_REDACTED      HostType = 1
	HOSTTYPE_PRISM_ELEMENT HostType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HostType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_ELEMENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HostType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_ELEMENT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HostType) index(name string) HostType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRISM_ELEMENT",
	}
	for idx := range names {
		if names[idx] == name {
			return HostType(idx)
		}
	}
	return HOSTTYPE_UNKNOWN
}

func (e *HostType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HostType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HostType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HostType) Ref() *HostType {
	return &e
}

/*
Hypervisor type where the Recovery point is generated.
*/
type HypervisorType int

const (
	HYPERVISORTYPE_UNKNOWN  HypervisorType = 0
	HYPERVISORTYPE_REDACTED HypervisorType = 1
	HYPERVISORTYPE_AHV      HypervisorType = 2
	HYPERVISORTYPE_HYPERV   HypervisorType = 3
	HYPERVISORTYPE_VMWARE   HypervisorType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HypervisorType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"HYPERV",
		"VMWARE",
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
		"AHV",
		"HYPERV",
		"VMWARE",
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
		"AHV",
		"HYPERV",
		"VMWARE",
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
IP address.
*/
type IpAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  IP address value in IPv4 format.
	*/
	Ipv4 *string `json:"ipv4"`
}

func (p *IpAddress) MarshalJSON() ([]byte, error) {
	type IpAddressProxy IpAddress
	return json.Marshal(struct {
		*IpAddressProxy
		Ipv4 *string `json:"ipv4,omitempty"`
	}{
		IpAddressProxy: (*IpAddressProxy)(p),
		Ipv4:           p.Ipv4,
	})
}

func NewIpAddress() *IpAddress {
	p := new(IpAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.IpAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.IpAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Location reference where the specified Recovery point is present.
*/
type LocationReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster where the Recovery point is present.
	*/
	LocationExtId *string `json:"locationExtId,omitempty"`
}

func NewLocationReference() *LocationReference {
	p := new(LocationReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.LocationReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.LocationReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Power state of the VM when the Recovery point is generated.
*/
type PowerState int

const (
	POWERSTATE_UNKNOWN  PowerState = 0
	POWERSTATE_REDACTED PowerState = 1
	POWERSTATE_OFF      PowerState = 2
	POWERSTATE_ON       PowerState = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PowerState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OFF",
		"ON",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PowerState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OFF",
		"ON",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PowerState) index(name string) PowerState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OFF",
		"ON",
	}
	for idx := range names {
		if names[idx] == name {
			return PowerState(idx)
		}
	}
	return POWERSTATE_UNKNOWN
}

func (e *PowerState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PowerState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PowerState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PowerState) Ref() *PowerState {
	return &e
}

/*
Once a VM or volume group is associated with some protection policy, the schedule(s) in the protection policy kick in to achieve the specified recovery point objective. A protected resource represents the data protection view of such a VM or volume group. It contains information such as the restorable time ranges on the local Prism Central and the state of replication to the targets specified in all the applied protection policies.
*/
type ProtectedResource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The data protection details for the protected resource that are relevant to any of the clusters in the local Prism Central, like the time ranges available for recovery.
	*/
	DataProtectionInfo []DataProtectionInfo `json:"dataProtectionInfo,omitempty"`
	/*
	  The external identifier of the VM or the volume group associated with the protected resource.
	*/
	EntityExtId *string `json:"entityExtId,omitempty"`

	EntityType *ProtectedResourceEntityType `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	ReplicationStates []ReplicationState `json:"replicationStates,omitempty"`

	SourceSiteReference *DataProtectionSiteReference `json:"sourceSiteReference,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewProtectedResource() *ProtectedResource {
	p := new(ProtectedResource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResource"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ProtectedResource"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/protected-resources/{extId} Get operation
*/
type ProtectedResourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfProtectedResourceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewProtectedResourceApiResponse() *ProtectedResourceApiResponse {
	p := new(ProtectedResourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ProtectedResourceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ProtectedResourceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ProtectedResourceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfProtectedResourceApiResponseData()
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
Protected resource entity type.
*/
type ProtectedResourceEntityType int

const (
	PROTECTEDRESOURCEENTITYTYPE_UNKNOWN      ProtectedResourceEntityType = 0
	PROTECTEDRESOURCEENTITYTYPE_REDACTED     ProtectedResourceEntityType = 1
	PROTECTEDRESOURCEENTITYTYPE_VM           ProtectedResourceEntityType = 2
	PROTECTEDRESOURCEENTITYTYPE_VOLUME_GROUP ProtectedResourceEntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtectedResourceEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProtectedResourceEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProtectedResourceEntityType) index(name string) ProtectedResourceEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectedResourceEntityType(idx)
		}
	}
	return PROTECTEDRESOURCEENTITYTYPE_UNKNOWN
}

func (e *ProtectedResourceEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectedResourceEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectedResourceEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectedResourceEntityType) Ref() *ProtectedResourceEntityType {
	return &e
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/protected-resources/{extId}/$actions/promote Post operation
*/
type ProtectedResourcePromoteApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfProtectedResourcePromoteApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewProtectedResourcePromoteApiResponse() *ProtectedResourcePromoteApiResponse {
	p := new(ProtectedResourcePromoteApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourcePromoteApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ProtectedResourcePromoteApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ProtectedResourcePromoteApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ProtectedResourcePromoteApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfProtectedResourcePromoteApiResponseData()
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
Status of replication to a specified target site:
 ACTIVE - Replications are in-progress and the specified recovery point objective is met on the target site.
 SYNCING - The system is trying to meet the specified recovery point objective for the target site.
 DEGRADED - The replication schedule has been degraded to a higher recovery point objective.
 DISABLED - The replication schedule is disabled and there are no ongoing replications.
 DECOUPLED - The connection between recovery and source site is broken, the protected resource at the recovery
 site is out of sync and has been promoted to a live entity.
*/
type ProtectedResourceReplicationStatus int

const (
	PROTECTEDRESOURCEREPLICATIONSTATUS_UNKNOWN   ProtectedResourceReplicationStatus = 0
	PROTECTEDRESOURCEREPLICATIONSTATUS_REDACTED  ProtectedResourceReplicationStatus = 1
	PROTECTEDRESOURCEREPLICATIONSTATUS_ACTIVE    ProtectedResourceReplicationStatus = 2
	PROTECTEDRESOURCEREPLICATIONSTATUS_SYNCING   ProtectedResourceReplicationStatus = 3
	PROTECTEDRESOURCEREPLICATIONSTATUS_DEGRADED  ProtectedResourceReplicationStatus = 4
	PROTECTEDRESOURCEREPLICATIONSTATUS_DISABLED  ProtectedResourceReplicationStatus = 5
	PROTECTEDRESOURCEREPLICATIONSTATUS_DECOUPLED ProtectedResourceReplicationStatus = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtectedResourceReplicationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"SYNCING",
		"DEGRADED",
		"DISABLED",
		"DECOUPLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProtectedResourceReplicationStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"SYNCING",
		"DEGRADED",
		"DISABLED",
		"DECOUPLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProtectedResourceReplicationStatus) index(name string) ProtectedResourceReplicationStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"SYNCING",
		"DEGRADED",
		"DISABLED",
		"DECOUPLED",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectedResourceReplicationStatus(idx)
		}
	}
	return PROTECTEDRESOURCEREPLICATIONSTATUS_UNKNOWN
}

func (e *ProtectedResourceReplicationStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectedResourceReplicationStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectedResourceReplicationStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectedResourceReplicationStatus) Ref() *ProtectedResourceReplicationStatus {
	return &e
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/protected-resources/{extId}/$actions/restore Post operation
*/
type ProtectedResourceRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfProtectedResourceRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewProtectedResourceRestoreApiResponse() *ProtectedResourceRestoreApiResponse {
	p := new(ProtectedResourceRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourceRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ProtectedResourceRestoreApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ProtectedResourceRestoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ProtectedResourceRestoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfProtectedResourceRestoreApiResponseData()
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
Restore action specifications for a minutely scheduled protected resource.
*/
type ProtectedResourceRestoreSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the cluster on which the entity has valid restorable time ranges. The restored entity will be created on the same cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  UTC date and time in ISO 8601 format representing the time from when the state of the entity should be restored. This needs to be a valid time within the restorable time range(s) for the protected resource.
	*/
	RestoreTime *time.Time `json:"restoreTime,omitempty"`
}

func (p *ProtectedResourceRestoreSpec) MarshalJSON() ([]byte, error) {
	type ProtectedResourceRestoreSpecProxy ProtectedResourceRestoreSpec
	return json.Marshal(struct {
		*ProtectedResourceRestoreSpecProxy
		ClusterExtId *string `json:"clusterExtId,omitempty"`
	}{
		ProtectedResourceRestoreSpecProxy: (*ProtectedResourceRestoreSpecProxy)(p),
		ClusterExtId:                      p.ClusterExtId,
	})
}

func NewProtectedResourceRestoreSpec() *ProtectedResourceRestoreSpec {
	p := new(ProtectedResourceRestoreSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourceRestoreSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ProtectedResourceRestoreSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The restorable time range details that can be used to recover the protected resource.
*/
type RecoveryInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of restorable time ranges.
	*/
	RestorableTimeRanges []RestorableTimeRange `json:"restorableTimeRanges,omitempty"`
}

func NewRecoveryInfo() *RecoveryInfo {
	p := new(RecoveryInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the Recovery point along with all the captured VM and volume group Recovery point(s).
*/
type RecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  List of location references where the VM or volume group Recovery points are a part of the specified Recovery point.
	*/
	LocationReferences []LocationReference `json:"locationReferences,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A read only field inserted into Recovery point at the time of Recovery point creation, indicating the external identifier of the user who created this Recovery point.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []import5.VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`
	/*
	  List of VM Recovery points that are a part of the specified top-level Recovery point.
	*/
	VmRecoveryPoints []VmSubRecoveryPoint `json:"vmRecoveryPoints,omitempty"`
	/*
	  List of VM external identifiers that are intended to be captured as part of the Recovery point.
	*/
	VmReferenceList []string `json:"vmReferenceList,omitempty"`
	/*
	  List of volume group Recovery points that are a part of the specified top-level Recovery point.
	*/
	VolumeGroupRecoveryPoints []VolumeGroupSubRecoveryPoint `json:"volumeGroupRecoveryPoints,omitempty"`
	/*
	  List of volume group external identifiers that are intended to be captured as part of the Recovery point.
	*/
	VolumeGroupReferenceList []string `json:"volumeGroupReferenceList,omitempty"`
}

func NewRecoveryPoint() *RecoveryPoint {
	p := new(RecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{extId} Get operation
*/
type RecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecoveryPointApiResponse() *RecoveryPointApiResponse {
	p := new(RecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecoveryPointApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points Get operation
*/
type RecoveryPointListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecoveryPointListApiResponse() *RecoveryPointListApiResponse {
	p := new(RecoveryPointListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoveryPointListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecoveryPointListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecoveryPointListApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{extId}/$actions/replicate Post operation
*/
type RecoveryPointReplicateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointReplicateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecoveryPointReplicateApiResponse() *RecoveryPointReplicateApiResponse {
	p := new(RecoveryPointReplicateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointReplicateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointReplicateApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoveryPointReplicateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecoveryPointReplicateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecoveryPointReplicateApiResponseData()
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
External identifier of the cluster and the Prism Central where the Recovery point should be replicated.
*/
type RecoveryPointReplicationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the Prism Central.
	*/
	PcExtId *string `json:"pcExtId,omitempty"`
	/*
	  Target availability zone reference for the entities to be replicated.
	*/
	TargetAvailabilityZoneId *string `json:"targetAvailabilityZoneId,omitempty"`
	/*
	  Details of the Cluster reference in the target availability zone for the entities to be replicated.
	*/
	TargetClusterId *string `json:"targetClusterId,omitempty"`
}

func NewRecoveryPointReplicationSpec() *RecoveryPointReplicationSpec {
	p := new(RecoveryPointReplicationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointReplicationSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointReplicationSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for the restore action on the top-level Recovery point. For a Recovery point that contains multiple VM or volume group Recovery points, users can selectively trigger restore for any specific set of VM or volume group Recovery point(s). In case no particular selection is made, all VM and volume group Recovery points that are a part of the top-level Recovery point will be restored.
*/
type RecoveryPointRestorationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Recovery points are restored at the associated location reference by default. However, there is no particular location reference associated with Recovery points located on the cloud. In such a case, the client must specify the external identifier of the cluster on which the entity should be restored.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  List of specifications to restore a specific VM Recovery point(s) that are a part of the top-level recovery point. A specific VM Recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VmRecoveryPointRestoreOverrides []VmRecoveryPointRestoreOverride `json:"vmRecoveryPointRestoreOverrides,omitempty"`
	/*
	  List of specifications to restore a specific volume group Recovery point(s) that are a part of the top-level Recovery point. A specific volume group Recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VolumeGroupRecoveryPointRestoreOverrides []VolumeGroupRecoveryPointRestoreOverride `json:"volumeGroupRecoveryPointRestoreOverrides,omitempty"`
}

func NewRecoveryPointRestorationSpec() *RecoveryPointRestorationSpec {
	p := new(RecoveryPointRestorationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRestorationSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointRestorationSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{extId}/$actions/restore Post operation
*/
type RecoveryPointRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecoveryPointRestoreApiResponse() *RecoveryPointRestoreApiResponse {
	p := new(RecoveryPointRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RecoveryPointRestoreApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoveryPointRestoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecoveryPointRestoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecoveryPointRestoreApiResponseData()
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
Changed Regions list comprising of offset and regions length pertaining to the Disk Recovery Point.
*/
type Regions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The length of the regions in bytes.
	*/
	Length *int64 `json:"length,omitempty"`
	/*
	  The byte offset indicating the start of the regions.
	*/
	Offset *int64 `json:"offset,omitempty"`
}

func NewRegions() *Regions {
	p := new(Regions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Regions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Regions"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Replication related information about the protected resource.
*/
type ReplicationState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the protection policy associated with the protected resource.
	*/
	ProtectionPolicyReference *string `json:"protectionPolicyReference,omitempty"`
	/*
	  The recovery point objective of the schedule in seconds.
	*/
	RecoveryPointObjective *int64 `json:"recoveryPointObjective,omitempty"`

	ReplicationStatus *ProtectedResourceReplicationStatus `json:"replicationStatus,omitempty"`

	TargetSiteReference *DataProtectionSiteReference `json:"targetSiteReference,omitempty"`
}

func NewReplicationState() *ReplicationState {
	p := new(ReplicationState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ReplicationState"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ReplicationState"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The start and end time details, both inclusively represent the range of time during which the entity is restorable.
*/
type RestorableTimeRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UTC date and time in ISO 8601 format representing the time when the restorable time range for the entity ends.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  UTC date and time in ISO 8601 format representing the time when the restorable time range for the entity starts.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func NewRestorableTimeRange() *RestorableTimeRange {
	p := new(RestorableTimeRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RestorableTimeRange"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.RestorableTimeRange"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification to set the expiration time of the Recovery point.
*/
type SetExpirationTimeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected. If not specified, the Recovery point will never expire.
	*/
	ExpirationTime *time.Time `json:"expirationTime"`
}

func (p *SetExpirationTimeSpec) MarshalJSON() ([]byte, error) {
	type SetExpirationTimeSpecProxy SetExpirationTimeSpec
	return json.Marshal(struct {
		*SetExpirationTimeSpecProxy
		ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	}{
		SetExpirationTimeSpecProxy: (*SetExpirationTimeSpecProxy)(p),
		ExpirationTime:             p.ExpirationTime,
	})
}

func NewSetExpirationTimeSpec() *SetExpirationTimeSpec {
	p := new(SetExpirationTimeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SetExpirationTimeSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.SetExpirationTimeSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification corresponding to the cluster reference information along with the related AZ information.
*/
type SiteReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID corresponding to the Availability Zone where an entity resides or is synced.
	*/
	AvailabilityZoneReference *string `json:"availabilityZoneReference,omitempty"`
	/*
	  The UUID corresponding to the cluster where an entity resides or is synced.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
}

func NewSiteReference() *SiteReference {
	p := new(SiteReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SiteReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.SiteReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
CIDR of the network gateway.
*/
type Subnet struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4 *import6.IPv4Address `json:"ipv4,omitempty"`
}

func NewSubnet() *Subnet {
	p := new(Subnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Subnet"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Subnet"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/synced-volume-groups/{extId} Get operation
*/
type SyncedVolumeGroupByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSyncedVolumeGroupByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSyncedVolumeGroupByIdApiResponse() *SyncedVolumeGroupByIdApiResponse {
	p := new(SyncedVolumeGroupByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SyncedVolumeGroupByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.SyncedVolumeGroupByIdApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SyncedVolumeGroupByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SyncedVolumeGroupByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSyncedVolumeGroupByIdApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/synced-volume-groups/{extId}/$actions/promote Post operation
*/
type SynchronousReplicaPromoteApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSynchronousReplicaPromoteApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSynchronousReplicaPromoteApiResponse() *SynchronousReplicaPromoteApiResponse {
	p := new(SynchronousReplicaPromoteApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SynchronousReplicaPromoteApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.SynchronousReplicaPromoteApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SynchronousReplicaPromoteApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SynchronousReplicaPromoteApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSynchronousReplicaPromoteApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{extId}/$actions/set-expiration-time Post operation
*/
type UpdateRecoveryPointExpirationTimeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRecoveryPointExpirationTimeApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRecoveryPointExpirationTimeApiResponse() *UpdateRecoveryPointExpirationTimeApiResponse {
	p := new(UpdateRecoveryPointExpirationTimeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.UpdateRecoveryPointExpirationTimeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.UpdateRecoveryPointExpirationTimeApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateRecoveryPointExpirationTimeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateRecoveryPointExpirationTimeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateRecoveryPointExpirationTimeApiResponseData()
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
A model that represents VM Recovery point properties.
*/
type VMRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []import5.VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`

	Vm *Vm `json:"vm,omitempty"`
}

func NewVMRecoveryPoint() *VMRecoveryPoint {
	p := new(VMRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VMRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VMRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId} Get operation
*/
type VMRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVMRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVMRecoveryPointApiResponse() *VMRecoveryPointApiResponse {
	p := new(VMRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VMRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VMRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VMRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VMRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVMRecoveryPointApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/vm-recovery-points Get operation
*/
type VMRecoveryPointListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVMRecoveryPointListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVMRecoveryPointListApiResponse() *VMRecoveryPointListApiResponse {
	p := new(VMRecoveryPointListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VMRecoveryPointListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VMRecoveryPointListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VMRecoveryPointListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VMRecoveryPointListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVMRecoveryPointListApiResponseData()
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
Recovery point validation failure details, if any.
*/
type ValidateRecoveryPointResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Validation failure or warning details.
	*/
	Detail *string `json:"detail,omitempty"`
	/*
	  Possible resolution for the validation failure or warning.
	*/
	Resolution *string `json:"resolution,omitempty"`

	Severity *ValidationSeverity `json:"severity"`
}

func (p *ValidateRecoveryPointResult) MarshalJSON() ([]byte, error) {
	type ValidateRecoveryPointResultProxy ValidateRecoveryPointResult
	return json.Marshal(struct {
		*ValidateRecoveryPointResultProxy
		Severity *ValidationSeverity `json:"severity,omitempty"`
	}{
		ValidateRecoveryPointResultProxy: (*ValidateRecoveryPointResultProxy)(p),
		Severity:                         p.Severity,
	})
}

func NewValidateRecoveryPointResult() *ValidateRecoveryPointResult {
	p := new(ValidateRecoveryPointResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ValidateRecoveryPointResult"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ValidateRecoveryPointResult"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/$actions/validate-restore Post operation
*/
type ValidateRestoreVmRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfValidateRestoreVmRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewValidateRestoreVmRecoveryPointApiResponse() *ValidateRestoreVmRecoveryPointApiResponse {
	p := new(ValidateRestoreVmRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ValidateRestoreVmRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.ValidateRestoreVmRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ValidateRestoreVmRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ValidateRestoreVmRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfValidateRestoreVmRecoveryPointApiResponseData()
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
Validation failure severity. Applicable values are WARNING or ERROR.
*/
type ValidationSeverity int

const (
	VALIDATIONSEVERITY_UNKNOWN  ValidationSeverity = 0
	VALIDATIONSEVERITY_REDACTED ValidationSeverity = 1
	VALIDATIONSEVERITY_WARNING  ValidationSeverity = 2
	VALIDATIONSEVERITY_ERROR    ValidationSeverity = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ValidationSeverity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ValidationSeverity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ValidationSeverity) index(name string) ValidationSeverity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return ValidationSeverity(idx)
		}
	}
	return VALIDATIONSEVERITY_UNKNOWN
}

func (e *ValidationSeverity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ValidationSeverity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ValidationSeverity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ValidationSeverity) Ref() *ValidationSeverity {
	return &e
}

/*
The configuration of the VM captured by the Recovery point.
*/
type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of cores per vCPU.
	*/
	CoresPerVcpu *int `json:"coresPerVcpu,omitempty"`

	HardwareArchitecture *HardwareArchitecture `json:"hardwareArchitecture,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/*
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  The name of the VM that is captured as part of this Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Number of vCPUs.
	*/
	NumVcpus *int `json:"numVcpus,omitempty"`

	PowerState *PowerState `json:"powerState,omitempty"`
	/*
	  VM external identifier which is captured as part of this Recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
	/*
	  Docref(vmNumThreadsPerCoreDesc)
	*/
	VmNumThreadsPerCore *int `json:"vmNumThreadsPerCore,omitempty"`
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Vm"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Vm"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for the restore action on the VM Recovery point.
*/
type VmRecoveryPointRestorationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvVmOverrideSpec *AhvVmOverrideSpec `json:"ahvVmOverrideSpec,omitempty"`
	/*
	  Recovery points are restored at the associated location reference by default. However, there is no particular location reference associated with Recovery points located on the cloud. In such a case, the client must specify the external identifier of the cluster on which the entity should be restored.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	EsxiVmOverrideSpec *EsxiVmOverrideSpec `json:"esxiVmOverrideSpec,omitempty"`
}

func NewVmRecoveryPointRestorationSpec() *VmRecoveryPointRestorationSpec {
	p := new(VmRecoveryPointRestorationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointRestorationSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VmRecoveryPointRestorationSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/$actions/restore Post operation
*/
type VmRecoveryPointRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVmRecoveryPointRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVmRecoveryPointRestoreApiResponse() *VmRecoveryPointRestoreApiResponse {
	p := new(VmRecoveryPointRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VmRecoveryPointRestoreApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VmRecoveryPointRestoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VmRecoveryPointRestoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVmRecoveryPointRestoreApiResponseData()
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

type VmRecoveryPointRestoreOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvVmOverrideSpec *AhvVmOverrideSpec `json:"ahvVmOverrideSpec,omitempty"`

	EsxiVmOverrideSpec *EsxiVmOverrideSpec `json:"esxiVmOverrideSpec,omitempty"`
	/*
	  External identifier of a VM Recovery point, that is a part of the top-level Recovery point.
	*/
	VmRecoveryPointId *string `json:"vmRecoveryPointId"`
}

func (p *VmRecoveryPointRestoreOverride) MarshalJSON() ([]byte, error) {
	type VmRecoveryPointRestoreOverrideProxy VmRecoveryPointRestoreOverride
	return json.Marshal(struct {
		*VmRecoveryPointRestoreOverrideProxy
		VmRecoveryPointId *string `json:"vmRecoveryPointId,omitempty"`
	}{
		VmRecoveryPointRestoreOverrideProxy: (*VmRecoveryPointRestoreOverrideProxy)(p),
		VmRecoveryPointId:                   p.VmRecoveryPointId,
	})
}

func NewVmRecoveryPointRestoreOverride() *VmRecoveryPointRestoreOverride {
	p := new(VmRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VmRecoveryPointRestoreOverride"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents VM Recovery point properties.
*/
type VmSubRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point. This identifier is used to identify the same instances of a Recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of Recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as part of this Recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the VM Recovery point using its URL.
	*/
	VmRecoveryPointId *string `json:"vmRecoveryPointId,omitempty"`
}

func NewVmSubRecoveryPoint() *VmSubRecoveryPoint {
	p := new(VmSubRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmSubRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VmSubRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmSubRecoveryPointProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point. This identifier is used to identify the same instances of a Recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of Recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as part of this Recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the VM Recovery point using its URL.
	*/
	VmRecoveryPointId *string `json:"vmRecoveryPointId,omitempty"`
}

func NewVmSubRecoveryPointProjection() *VmSubRecoveryPointProjection {
	p := new(VmSubRecoveryPointProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmSubRecoveryPointProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VmSubRecoveryPointProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The configuration of the volume group captured by the Recovery point.
*/
type VolumeGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the volume group that is captured as part of this Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Volume group external identifier which is captured as part of this Recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func NewVolumeGroup() *VolumeGroup {
	p := new(VolumeGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Recovery point restore that overrides the volume group configuration. The specified properties will be overridden for the restored volume group.
*/
type VolumeGroupOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the restored volume group.
	*/
	Name *string `json:"name,omitempty"`
}

func NewVolumeGroupOverrideSpec() *VolumeGroupOverrideSpec {
	p := new(VolumeGroupOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupOverrideSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents volume group Recovery point properties.
*/
type VolumeGroupRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import5.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import5.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of additional metadata provided by the client at the time of Recovery point creation.
	*/
	VendorSpecificProperties []import5.VendorSpecificProperty `json:"vendorSpecificProperties,omitempty"`

	VolumeGroup *VolumeGroup `json:"volumeGroup,omitempty"`
}

func NewVolumeGroupRecoveryPoint() *VolumeGroupRecoveryPoint {
	p := new(VolumeGroupRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId} Get operation
*/
type VolumeGroupRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVolumeGroupRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVolumeGroupRecoveryPointApiResponse() *VolumeGroupRecoveryPointApiResponse {
	p := new(VolumeGroupRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPointApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VolumeGroupRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VolumeGroupRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVolumeGroupRecoveryPointApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points Get operation
*/
type VolumeGroupRecoveryPointListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVolumeGroupRecoveryPointListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVolumeGroupRecoveryPointListApiResponse() *VolumeGroupRecoveryPointListApiResponse {
	p := new(VolumeGroupRecoveryPointListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPointListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VolumeGroupRecoveryPointListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VolumeGroupRecoveryPointListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVolumeGroupRecoveryPointListApiResponseData()
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
Specification for the restore action on the volume group Recovery point.
*/
type VolumeGroupRecoveryPointRestorationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Recovery points are restored at the associated location reference by default. However, there is no particular location reference associated with Recovery points located on the cloud. In such a case, the client must specify the external identifier of the cluster on which the entity should be restored.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  The name of the restored volume group.
	*/
	Name *string `json:"name,omitempty"`
}

func NewVolumeGroupRecoveryPointRestorationSpec() *VolumeGroupRecoveryPointRestorationSpec {
	p := new(VolumeGroupRecoveryPointRestorationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointRestorationSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPointRestorationSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/$actions/restore Post operation
*/
type VolumeGroupRecoveryPointRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVolumeGroupRecoveryPointRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVolumeGroupRecoveryPointRestoreApiResponse() *VolumeGroupRecoveryPointRestoreApiResponse {
	p := new(VolumeGroupRecoveryPointRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPointRestoreApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VolumeGroupRecoveryPointRestoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VolumeGroupRecoveryPointRestoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVolumeGroupRecoveryPointRestoreApiResponseData()
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

type VolumeGroupRecoveryPointRestoreOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	VolumeGroupOverrideSpec *VolumeGroupOverrideSpec `json:"volumeGroupOverrideSpec,omitempty"`
	/*
	  External identifier of a volume group Recovery point, that is a part of the top-level Recovery point.
	*/
	VolumeGroupRecoveryPointId *string `json:"volumeGroupRecoveryPointId"`
}

func (p *VolumeGroupRecoveryPointRestoreOverride) MarshalJSON() ([]byte, error) {
	type VolumeGroupRecoveryPointRestoreOverrideProxy VolumeGroupRecoveryPointRestoreOverride
	return json.Marshal(struct {
		*VolumeGroupRecoveryPointRestoreOverrideProxy
		VolumeGroupRecoveryPointId *string `json:"volumeGroupRecoveryPointId,omitempty"`
	}{
		VolumeGroupRecoveryPointRestoreOverrideProxy: (*VolumeGroupRecoveryPointRestoreOverrideProxy)(p),
		VolumeGroupRecoveryPointId:                   p.VolumeGroupRecoveryPointId,
	})
}

func NewVolumeGroupRecoveryPointRestoreOverride() *VolumeGroupRecoveryPointRestoreOverride {
	p := new(VolumeGroupRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupRecoveryPointRestoreOverride"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents volume group Recovery point properties.
*/
type VolumeGroupSubRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point. This identifier is used to identify the same instances of a Recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  Category key-value pairs associated with the volume group at the time of Recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VolumeGroupCategories []string `json:"volumeGroupCategories,omitempty"`
	/*
	  Volume group external identifier which is captured as part of this Recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group Recovery point using its URL.
	*/
	VolumeGroupRecoveryPointId *string `json:"volumeGroupRecoveryPointId,omitempty"`
}

func NewVolumeGroupSubRecoveryPoint() *VolumeGroupSubRecoveryPoint {
	p := new(VolumeGroupSubRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupSubRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupSubRecoveryPoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VolumeGroupSubRecoveryPointProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the Consistency group which the entity was part of at the time of Recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point. This identifier is used to identify the same instances of a Recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  Category key-value pairs associated with the volume group at the time of Recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VolumeGroupCategories []string `json:"volumeGroupCategories,omitempty"`
	/*
	  Volume group external identifier which is captured as part of this Recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group Recovery point using its URL.
	*/
	VolumeGroupRecoveryPointId *string `json:"volumeGroupRecoveryPointId,omitempty"`
}

func NewVolumeGroupSubRecoveryPointProjection() *VolumeGroupSubRecoveryPointProjection {
	p := new(VolumeGroupSubRecoveryPointProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupSubRecoveryPointProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupSubRecoveryPointProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification corresponding to the Volume Group synced state. The primary site reference corresponds to the cluster reference information along with the related AZ information where the entity resides. The secondary site reference corresponds to the cluster site reference information along with the related AZ information where the Volume Group is synced to.
*/
type VolumeGroupSyncContext struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	PrimarySite *SiteReference `json:"primarySite,omitempty"`

	SecondarySite *SiteReference `json:"secondarySite,omitempty"`
}

func NewVolumeGroupSyncContext() *VolumeGroupSyncContext {
	p := new(VolumeGroupSyncContext)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupSyncContext"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VolumeGroupSyncContext"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPC reference to which the network belongs.
*/
type VpcReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the VPC.
	*/
	ExtId *string `json:"extId"`
	/*
	  Name of the VPC, this is a read-only field.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *VpcReference) MarshalJSON() ([]byte, error) {
	type VpcReferenceProxy VpcReference
	return json.Marshal(struct {
		*VpcReferenceProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		VpcReferenceProxy: (*VpcReferenceProxy)(p),
		ExtId:             p.ExtId,
	})
}

func NewVpcReference() *VpcReference {
	p := new(VpcReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VpcReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.VpcReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents the details of a Witness site.
*/
type Witness struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of host references of the Witness.
	*/
	HostReferences []HostReference `json:"hostReferences,omitempty"`
	/*
	  List of IP addresses to be assigned to the Virtual Machine.
	*/
	IpAddresses []string `json:"ipAddresses,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness site.
	*/
	Name *string `json:"name,omitempty"`

	Status *WitnessAvailabilityStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWitness() *Witness {
	p := new(Witness)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Witness"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.Witness"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.a4/config/witness/{extId} Get operation
*/
type WitnessApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfWitnessApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewWitnessApiResponse() *WitnessApiResponse {
	p := new(WitnessApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.WitnessApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "dataprotection.v4.r0.a4.config.WitnessApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *WitnessApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *WitnessApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfWitnessApiResponseData()
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
Availability status of the Witness.
*/
type WitnessAvailabilityStatus int

const (
	WITNESSAVAILABILITYSTATUS_UNKNOWN     WitnessAvailabilityStatus = 0
	WITNESSAVAILABILITYSTATUS_REDACTED    WitnessAvailabilityStatus = 1
	WITNESSAVAILABILITYSTATUS_AVAILABLE   WitnessAvailabilityStatus = 2
	WITNESSAVAILABILITYSTATUS_UNAVAILABLE WitnessAvailabilityStatus = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *WitnessAvailabilityStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABLE",
		"UNAVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e WitnessAvailabilityStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABLE",
		"UNAVAILABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *WitnessAvailabilityStatus) index(name string) WitnessAvailabilityStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVAILABLE",
		"UNAVAILABLE",
	}
	for idx := range names {
		if names[idx] == name {
			return WitnessAvailabilityStatus(idx)
		}
	}
	return WITNESSAVAILABILITYSTATUS_UNKNOWN
}

func (e *WitnessAvailabilityStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for WitnessAvailabilityStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *WitnessAvailabilityStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e WitnessAvailabilityStatus) Ref() *WitnessAvailabilityStatus {
	return &e
}

type OneOfDeleteConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteConsistencyGroupApiResponseData() *OneOfDeleteConsistencyGroupApiResponseData {
	p := new(OneOfDeleteConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteConsistencyGroupApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(interface{})
		}
		*p.oneOfType1 = nil
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

func (p *OneOfDeleteConsistencyGroupApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == *vOneOfType1 {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(interface{})
			}
			*p.oneOfType1 = nil
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
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteConsistencyGroupApiResponseData"))
}

func (p *OneOfDeleteConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteConsistencyGroupApiResponseData")
}

type OneOfDiskRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *DiskRecoveryPoint     `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDiskRecoveryPointApiResponseData() *OneOfDiskRecoveryPointApiResponseData {
	p := new(OneOfDiskRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiskRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiskRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case DiskRecoveryPoint:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DiskRecoveryPoint)
		}
		*p.oneOfType0 = v.(DiskRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfDiskRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDiskRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DiskRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.DiskRecoveryPoint" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DiskRecoveryPoint)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiskRecoveryPointApiResponseData"))
}

func (p *OneOfDiskRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDiskRecoveryPointApiResponseData")
}

type OneOfConsistencyGroupMigrateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfConsistencyGroupMigrateApiResponseData() *OneOfConsistencyGroupMigrateApiResponseData {
	p := new(OneOfConsistencyGroupMigrateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConsistencyGroupMigrateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConsistencyGroupMigrateApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfConsistencyGroupMigrateApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfConsistencyGroupMigrateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConsistencyGroupMigrateApiResponseData"))
}

func (p *OneOfConsistencyGroupMigrateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfConsistencyGroupMigrateApiResponseData")
}

type OneOfSyncedVolumeGroupByIdApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import1.ErrorResponse  `json:"-"`
	oneOfType0    *VolumeGroupSyncContext `json:"-"`
}

func NewOneOfSyncedVolumeGroupByIdApiResponseData() *OneOfSyncedVolumeGroupByIdApiResponseData {
	p := new(OneOfSyncedVolumeGroupByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSyncedVolumeGroupByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSyncedVolumeGroupByIdApiResponseData is nil"))
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
	case VolumeGroupSyncContext:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VolumeGroupSyncContext)
		}
		*p.oneOfType0 = v.(VolumeGroupSyncContext)
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

func (p *OneOfSyncedVolumeGroupByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfSyncedVolumeGroupByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VolumeGroupSyncContext)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.VolumeGroupSyncContext" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VolumeGroupSyncContext)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSyncedVolumeGroupByIdApiResponseData"))
}

func (p *OneOfSyncedVolumeGroupByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfSyncedVolumeGroupByIdApiResponseData")
}

type OneOfVolumeGroupRecoveryPointApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    *VolumeGroupRecoveryPoint `json:"-"`
	oneOfType400  *import1.ErrorResponse    `json:"-"`
}

func NewOneOfVolumeGroupRecoveryPointApiResponseData() *OneOfVolumeGroupRecoveryPointApiResponseData {
	p := new(OneOfVolumeGroupRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVolumeGroupRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVolumeGroupRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case VolumeGroupRecoveryPoint:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VolumeGroupRecoveryPoint)
		}
		*p.oneOfType0 = v.(VolumeGroupRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfVolumeGroupRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVolumeGroupRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VolumeGroupRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.VolumeGroupRecoveryPoint" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VolumeGroupRecoveryPoint)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVolumeGroupRecoveryPointApiResponseData"))
}

func (p *OneOfVolumeGroupRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVolumeGroupRecoveryPointApiResponseData")
}

type OneOfRecoveryPointReplicateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfRecoveryPointReplicateApiResponseData() *OneOfRecoveryPointReplicateApiResponseData {
	p := new(OneOfRecoveryPointReplicateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointReplicateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointReplicateApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfRecoveryPointReplicateApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfRecoveryPointReplicateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointReplicateApiResponseData"))
}

func (p *OneOfRecoveryPointReplicateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointReplicateApiResponseData")
}

type OneOfRecoveryPointRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfRecoveryPointRestoreApiResponseData() *OneOfRecoveryPointRestoreApiResponseData {
	p := new(OneOfRecoveryPointRestoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointRestoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointRestoreApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfRecoveryPointRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfRecoveryPointRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRestoreApiResponseData"))
}

func (p *OneOfRecoveryPointRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRestoreApiResponseData")
}

type OneOfDiskRecoveryPointListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []DiskRecoveryPoint    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDiskRecoveryPointListApiResponseData() *OneOfDiskRecoveryPointListApiResponseData {
	p := new(OneOfDiskRecoveryPointListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiskRecoveryPointListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiskRecoveryPointListApiResponseData is nil"))
	}
	switch v.(type) {
	case []DiskRecoveryPoint:
		p.oneOfType0 = v.([]DiskRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.DiskRecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.DiskRecoveryPoint>"
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

func (p *OneOfDiskRecoveryPointListApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.DiskRecoveryPoint>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDiskRecoveryPointListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]DiskRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.DiskRecoveryPoint" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.DiskRecoveryPoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.DiskRecoveryPoint>"
			return nil

		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiskRecoveryPointListApiResponseData"))
}

func (p *OneOfDiskRecoveryPointListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.DiskRecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDiskRecoveryPointListApiResponseData")
}

type OneOfUpdateRecoveryPointExpirationTimeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateRecoveryPointExpirationTimeApiResponseData() *OneOfUpdateRecoveryPointExpirationTimeApiResponseData {
	p := new(OneOfUpdateRecoveryPointExpirationTimeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRecoveryPointExpirationTimeApiResponseData is nil"))
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

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRecoveryPointExpirationTimeApiResponseData"))
}

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRecoveryPointExpirationTimeApiResponseData")
}

type OneOfProtectedResourceRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfProtectedResourceRestoreApiResponseData() *OneOfProtectedResourceRestoreApiResponseData {
	p := new(OneOfProtectedResourceRestoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfProtectedResourceRestoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfProtectedResourceRestoreApiResponseData is nil"))
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

func (p *OneOfProtectedResourceRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfProtectedResourceRestoreApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourceRestoreApiResponseData"))
}

func (p *OneOfProtectedResourceRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourceRestoreApiResponseData")
}

type OneOfVMRecoveryPointListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []VMRecoveryPoint      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfVMRecoveryPointListApiResponseData() *OneOfVMRecoveryPointListApiResponseData {
	p := new(OneOfVMRecoveryPointListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVMRecoveryPointListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVMRecoveryPointListApiResponseData is nil"))
	}
	switch v.(type) {
	case []VMRecoveryPoint:
		p.oneOfType0 = v.([]VMRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.VMRecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.VMRecoveryPoint>"
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

func (p *OneOfVMRecoveryPointListApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.VMRecoveryPoint>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVMRecoveryPointListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]VMRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.VMRecoveryPoint" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.VMRecoveryPoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.VMRecoveryPoint>"
			return nil

		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVMRecoveryPointListApiResponseData"))
}

func (p *OneOfVMRecoveryPointListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.VMRecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVMRecoveryPointListApiResponseData")
}

type OneOfVolumeGroupRecoveryPointRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfVolumeGroupRecoveryPointRestoreApiResponseData() *OneOfVolumeGroupRecoveryPointRestoreApiResponseData {
	p := new(OneOfVolumeGroupRecoveryPointRestoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVolumeGroupRecoveryPointRestoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVolumeGroupRecoveryPointRestoreApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfVolumeGroupRecoveryPointRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfVolumeGroupRecoveryPointRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVolumeGroupRecoveryPointRestoreApiResponseData"))
}

func (p *OneOfVolumeGroupRecoveryPointRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfVolumeGroupRecoveryPointRestoreApiResponseData")
}

type OneOfVolumeGroupRecoveryPointListApiResponseData struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType0    []VolumeGroupRecoveryPoint `json:"-"`
	oneOfType400  *import1.ErrorResponse     `json:"-"`
}

func NewOneOfVolumeGroupRecoveryPointListApiResponseData() *OneOfVolumeGroupRecoveryPointListApiResponseData {
	p := new(OneOfVolumeGroupRecoveryPointListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVolumeGroupRecoveryPointListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVolumeGroupRecoveryPointListApiResponseData is nil"))
	}
	switch v.(type) {
	case []VolumeGroupRecoveryPoint:
		p.oneOfType0 = v.([]VolumeGroupRecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>"
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

func (p *OneOfVolumeGroupRecoveryPointListApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfVolumeGroupRecoveryPointListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]VolumeGroupRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.VolumeGroupRecoveryPoint" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>"
			return nil

		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVolumeGroupRecoveryPointListApiResponseData"))
}

func (p *OneOfVolumeGroupRecoveryPointListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.VolumeGroupRecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfVolumeGroupRecoveryPointListApiResponseData")
}

type OneOfCreateRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *RecoveryPoint         `json:"-"`
}

func NewOneOfCreateRecoveryPointApiResponseData() *OneOfCreateRecoveryPointApiResponseData {
	p := new(OneOfCreateRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateRecoveryPointApiResponseData is nil"))
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
	case RecoveryPoint:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(RecoveryPoint)
		}
		*p.oneOfType1 = v.(RecoveryPoint)
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

func (p *OneOfCreateRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfCreateRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.RecoveryPoint" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(RecoveryPoint)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRecoveryPointApiResponseData"))
}

func (p *OneOfCreateRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRecoveryPointApiResponseData")
}

type OneOfChangedRegionsListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ChangedRegions        `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfChangedRegionsListApiResponseData() *OneOfChangedRegionsListApiResponseData {
	p := new(OneOfChangedRegionsListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfChangedRegionsListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfChangedRegionsListApiResponseData is nil"))
	}
	switch v.(type) {
	case ChangedRegions:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ChangedRegions)
		}
		*p.oneOfType0 = v.(ChangedRegions)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfChangedRegionsListApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangedRegionsListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ChangedRegions)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.ChangedRegions" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ChangedRegions)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangedRegionsListApiResponseData"))
}

func (p *OneOfChangedRegionsListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangedRegionsListApiResponseData")
}

type OneOfVMRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *VMRecoveryPoint       `json:"-"`
}

func NewOneOfVMRecoveryPointApiResponseData() *OneOfVMRecoveryPointApiResponseData {
	p := new(OneOfVMRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVMRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVMRecoveryPointApiResponseData is nil"))
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
	case VMRecoveryPoint:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VMRecoveryPoint)
		}
		*p.oneOfType0 = v.(VMRecoveryPoint)
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

func (p *OneOfVMRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfVMRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VMRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.VMRecoveryPoint" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VMRecoveryPoint)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVMRecoveryPointApiResponseData"))
}

func (p *OneOfVMRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfVMRecoveryPointApiResponseData")
}

type OneOfRecoveryPointListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []RecoveryPoint        `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfRecoveryPointListApiResponseData() *OneOfRecoveryPointListApiResponseData {
	p := new(OneOfRecoveryPointListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointListApiResponseData is nil"))
	}
	switch v.(type) {
	case []RecoveryPoint:
		p.oneOfType0 = v.([]RecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPoint>"
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

func (p *OneOfRecoveryPointListApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecoveryPointListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.RecoveryPoint" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPoint>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPoint>"
			return nil

		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointListApiResponseData"))
}

func (p *OneOfRecoveryPointListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointListApiResponseData")
}

type OneOfSynchronousReplicaPromoteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfSynchronousReplicaPromoteApiResponseData() *OneOfSynchronousReplicaPromoteApiResponseData {
	p := new(OneOfSynchronousReplicaPromoteApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSynchronousReplicaPromoteApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSynchronousReplicaPromoteApiResponseData is nil"))
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

func (p *OneOfSynchronousReplicaPromoteApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfSynchronousReplicaPromoteApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSynchronousReplicaPromoteApiResponseData"))
}

func (p *OneOfSynchronousReplicaPromoteApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfSynchronousReplicaPromoteApiResponseData")
}

type OneOfConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ConsistencyGroup      `json:"-"`
}

func NewOneOfConsistencyGroupApiResponseData() *OneOfConsistencyGroupApiResponseData {
	p := new(OneOfConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConsistencyGroupApiResponseData is nil"))
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
	case ConsistencyGroup:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ConsistencyGroup)
		}
		*p.oneOfType0 = v.(ConsistencyGroup)
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

func (p *OneOfConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ConsistencyGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.ConsistencyGroup" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ConsistencyGroup)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConsistencyGroupApiResponseData"))
}

func (p *OneOfConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfConsistencyGroupApiResponseData")
}

type OneOfRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RecoveryPoint         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfRecoveryPointApiResponseData() *OneOfRecoveryPointApiResponseData {
	p := new(OneOfRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointApiResponseData is nil"))
	}
	switch v.(type) {
	case RecoveryPoint:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RecoveryPoint)
		}
		*p.oneOfType0 = v.(RecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.RecoveryPoint" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RecoveryPoint)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointApiResponseData"))
}

func (p *OneOfRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointApiResponseData")
}

type OneOfVmRecoveryPointRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfVmRecoveryPointRestoreApiResponseData() *OneOfVmRecoveryPointRestoreApiResponseData {
	p := new(OneOfVmRecoveryPointRestoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRecoveryPointRestoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRecoveryPointRestoreApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfVmRecoveryPointRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfVmRecoveryPointRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRecoveryPointRestoreApiResponseData"))
}

func (p *OneOfVmRecoveryPointRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfVmRecoveryPointRestoreApiResponseData")
}

type OneOfProtectedResourcePromoteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfProtectedResourcePromoteApiResponseData() *OneOfProtectedResourcePromoteApiResponseData {
	p := new(OneOfProtectedResourcePromoteApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfProtectedResourcePromoteApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfProtectedResourcePromoteApiResponseData is nil"))
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

func (p *OneOfProtectedResourcePromoteApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfProtectedResourcePromoteApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourcePromoteApiResponseData"))
}

func (p *OneOfProtectedResourcePromoteApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourcePromoteApiResponseData")
}

type OneOfCreateConsistencyGroupRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *RecoveryPoint         `json:"-"`
}

func NewOneOfCreateConsistencyGroupRecoveryPointApiResponseData() *OneOfCreateConsistencyGroupRecoveryPointApiResponseData {
	p := new(OneOfCreateConsistencyGroupRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateConsistencyGroupRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateConsistencyGroupRecoveryPointApiResponseData is nil"))
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
	case RecoveryPoint:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(RecoveryPoint)
		}
		*p.oneOfType1 = v.(RecoveryPoint)
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

func (p *OneOfCreateConsistencyGroupRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfCreateConsistencyGroupRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.RecoveryPoint" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(RecoveryPoint)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateConsistencyGroupRecoveryPointApiResponseData"))
}

func (p *OneOfCreateConsistencyGroupRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfCreateConsistencyGroupRecoveryPointApiResponseData")
}

type OneOfWitnessApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Witness               `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfWitnessApiResponseData() *OneOfWitnessApiResponseData {
	p := new(OneOfWitnessApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfWitnessApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfWitnessApiResponseData is nil"))
	}
	switch v.(type) {
	case Witness:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Witness)
		}
		*p.oneOfType0 = v.(Witness)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfWitnessApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfWitnessApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Witness)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.Witness" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Witness)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfWitnessApiResponseData"))
}

func (p *OneOfWitnessApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfWitnessApiResponseData")
}

type OneOfConsistencyGroupListApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType0    []ConsistencyGroup           `json:"-"`
	oneOfType400  *import1.ErrorResponse       `json:"-"`
	oneOfType401  []ConsistencyGroupProjection `json:"-"`
}

func NewOneOfConsistencyGroupListApiResponseData() *OneOfConsistencyGroupListApiResponseData {
	p := new(OneOfConsistencyGroupListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConsistencyGroupListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConsistencyGroupListApiResponseData is nil"))
	}
	switch v.(type) {
	case []ConsistencyGroup:
		p.oneOfType0 = v.([]ConsistencyGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.ConsistencyGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.ConsistencyGroup>"
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
	case []ConsistencyGroupProjection:
		p.oneOfType401 = v.([]ConsistencyGroupProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.ConsistencyGroupProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.ConsistencyGroupProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfConsistencyGroupListApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.ConsistencyGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.ConsistencyGroupProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfConsistencyGroupListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ConsistencyGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.ConsistencyGroup" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.ConsistencyGroup>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.ConsistencyGroup>"
			return nil

		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]ConsistencyGroupProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.ConsistencyGroupProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.ConsistencyGroupProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.ConsistencyGroupProjection>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConsistencyGroupListApiResponseData"))
}

func (p *OneOfConsistencyGroupListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.ConsistencyGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.ConsistencyGroupProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfConsistencyGroupListApiResponseData")
}

type OneOfValidateRestoreVmRecoveryPointApiResponseData struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType400  *import1.ErrorResponse        `json:"-"`
	oneOfType0    []ValidateRecoveryPointResult `json:"-"`
}

func NewOneOfValidateRestoreVmRecoveryPointApiResponseData() *OneOfValidateRestoreVmRecoveryPointApiResponseData {
	p := new(OneOfValidateRestoreVmRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfValidateRestoreVmRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfValidateRestoreVmRecoveryPointApiResponseData is nil"))
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
	case []ValidateRecoveryPointResult:
		p.oneOfType0 = v.([]ValidateRecoveryPointResult)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.ValidateRecoveryPointResult>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.ValidateRecoveryPointResult>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfValidateRestoreVmRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.ValidateRecoveryPointResult>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfValidateRestoreVmRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]ValidateRecoveryPointResult)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.config.ValidateRecoveryPointResult" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.ValidateRecoveryPointResult>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.ValidateRecoveryPointResult>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfValidateRestoreVmRecoveryPointApiResponseData"))
}

func (p *OneOfValidateRestoreVmRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.ValidateRecoveryPointResult>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfValidateRestoreVmRecoveryPointApiResponseData")
}

type OneOfDeleteRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType1    *import4.TaskReference `json:"-"`
	oneOfType0    *import3.Task          `json:"-"`
}

func NewOneOfDeleteRecoveryPointApiResponseData() *OneOfDeleteRecoveryPointApiResponseData {
	p := new(OneOfDeleteRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRecoveryPointApiResponseData is nil"))
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
	case import4.TaskReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import4.TaskReference)
		}
		*p.oneOfType1 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case import3.Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Task)
		}
		*p.oneOfType0 = v.(import3.Task)
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

func (p *OneOfDeleteRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import4.TaskReference)
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
	vOneOfType0 := new(import3.Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "storage.v4.config.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Task)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRecoveryPointApiResponseData"))
}

func (p *OneOfDeleteRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRecoveryPointApiResponseData")
}

type OneOfProtectedResourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ProtectedResource     `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfProtectedResourceApiResponseData() *OneOfProtectedResourceApiResponseData {
	p := new(OneOfProtectedResourceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfProtectedResourceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfProtectedResourceApiResponseData is nil"))
	}
	switch v.(type) {
	case ProtectedResource:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ProtectedResource)
		}
		*p.oneOfType0 = v.(ProtectedResource)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfProtectedResourceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfProtectedResourceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ProtectedResource)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.ProtectedResource" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ProtectedResource)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourceApiResponseData"))
}

func (p *OneOfProtectedResourceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourceApiResponseData")
}
