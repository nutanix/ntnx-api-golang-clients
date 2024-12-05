/*
 * Generated file models/datapolicies/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Data Policies APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Manage disaster recovery and storage policies.
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/dataprotection/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/prism/v4/config"
)

/*
Maintains a rolling window of recovery points for every schedule, starting with the hourly schedule and ending with the schedule created for the specified retention period. For example, if the retention period is 4 weeks, at any given time, you will have 24 of the latest hourly recovery points, 7 of the latest daily recovery points, and 4 of the latest weekly recovery points.
*/
type AutoRollupRetention struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Local *AutoRollupRetentionDetails `json:"local"`

	Remote *AutoRollupRetentionDetails `json:"remote,omitempty"`
}

func (p *AutoRollupRetention) MarshalJSON() ([]byte, error) {
	type AutoRollupRetentionProxy AutoRollupRetention
	return json.Marshal(struct {
		*AutoRollupRetentionProxy
		Local *AutoRollupRetentionDetails `json:"local,omitempty"`
	}{
		AutoRollupRetentionProxy: (*AutoRollupRetentionProxy)(p),
		Local:                    p.Local,
	})
}

func NewAutoRollupRetention() *AutoRollupRetention {
	p := new(AutoRollupRetention)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.AutoRollupRetention"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specifies the auto rollup retention details.
*/
type AutoRollupRetentionDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Multiplier to 'snapshotIntervalType'. For example, if 'snapshotIntervalType' is 'YEARLY' and 'multiple' is 5, then 5 years worth of rollup snapshots will be retained.
	*/
	Frequency *int `json:"frequency"`

	SnapshotIntervalType *SnapshotIntervalType `json:"snapshotIntervalType"`
}

func (p *AutoRollupRetentionDetails) MarshalJSON() ([]byte, error) {
	type AutoRollupRetentionDetailsProxy AutoRollupRetentionDetails
	return json.Marshal(struct {
		*AutoRollupRetentionDetailsProxy
		Frequency            *int                  `json:"frequency,omitempty"`
		SnapshotIntervalType *SnapshotIntervalType `json:"snapshotIntervalType,omitempty"`
	}{
		AutoRollupRetentionDetailsProxy: (*AutoRollupRetentionDetailsProxy)(p),
		Frequency:                       p.Frequency,
		SnapshotIntervalType:            p.SnapshotIntervalType,
	})
}

func NewAutoRollupRetentionDetails() *AutoRollupRetentionDetails {
	p := new(AutoRollupRetentionDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.AutoRollupRetentionDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /datapolicies/v4.0/config/protection-policies Post operation
*/
type CreateProtectionPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateProtectionPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateProtectionPolicyApiResponse() *CreateProtectionPolicyApiResponse {
	p := new(CreateProtectionPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.CreateProtectionPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateProtectionPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateProtectionPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateProtectionPolicyApiResponseData()
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
REST response for all response codes in API path /datapolicies/v4.0/config/protection-policies/{extId} Delete operation
*/
type DeleteProtectionPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteProtectionPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteProtectionPolicyApiResponse() *DeleteProtectionPolicyApiResponse {
	p := new(DeleteProtectionPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.DeleteProtectionPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteProtectionPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteProtectionPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteProtectionPolicyApiResponseData()
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
REST response for all response codes in API path /datapolicies/v4.0/config/protection-policies/{extId} Get operation
*/
type GetProtectionPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetProtectionPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetProtectionPolicyApiResponse() *GetProtectionPolicyApiResponse {
	p := new(GetProtectionPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.GetProtectionPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetProtectionPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetProtectionPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetProtectionPolicyApiResponseData()
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
A simple retention scheme for the policy. If you set the retention number to n, the policy retains the n recent recovery points.
*/
type LinearRetention struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the number of recovery points to retain on the local location.
	*/
	Local *int `json:"local"`
	/*
	  Specifies the number of recovery points to retain on the remote location.
	*/
	Remote *int `json:"remote,omitempty"`
}

func (p *LinearRetention) MarshalJSON() ([]byte, error) {
	type LinearRetentionProxy LinearRetention
	return json.Marshal(struct {
		*LinearRetentionProxy
		Local *int `json:"local,omitempty"`
	}{
		LinearRetentionProxy: (*LinearRetentionProxy)(p),
		Local:                p.Local,
	})
}

func NewLinearRetention() *LinearRetention {
	p := new(LinearRetention)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.LinearRetention"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /datapolicies/v4.0/config/protection-policies Get operation
*/
type ListProtectionPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListProtectionPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListProtectionPoliciesApiResponse() *ListProtectionPoliciesApiResponse {
	p := new(ListProtectionPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.ListProtectionPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListProtectionPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListProtectionPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListProtectionPoliciesApiResponseData()
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
List of Prism Element cluster external identifiers whose associated VMs and volume groups are protected. Only the primary location can have multiple clusters configured, while the other locations can specify only one cluster. Clusters must be specified for replication within the same Prism Central and cannot be specified for an MST type location. All clusters are considered if the cluster external identifier list is empty.
*/
type NutanixCluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Prism Element cluster external identifiers whose associated VMs and volume groups are protected. Only the primary location can have multiple clusters configured, while the other locations can specify only one cluster. Clusters must be specified for replication within the same Prism Central and cannot be specified for an MST type location. All clusters are considered if the cluster external identifier list is empty.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
}

func NewNutanixCluster() *NutanixCluster {
	p := new(NutanixCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.NutanixCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A protection policy automates the process of creating and replicating recovery points. When a protection policy is configured to create local recovery points, the request includes:<br> - The recovery point objective<br> - The retention policy<br> - The entities that need to be protected by specifying the categories in which they are tagged.<br> To automate recovery point replication, you can also specify the replication location(s). Only users with legacy roles, such as admin, can create a Cross-AZ protection policy.
*/
type ProtectionPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the list of external identifiers of categories that must be added to the protection policy. This policy will protect any VM or volume group associated with this category.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  Description of the protection policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Manual deletion of recovery points created by this policy can be driven through a multi-party authorization workflow. Hence, multiple approvers would be required to grant approvals for the delete operation.
	*/
	IsApprovalPolicyNeeded *bool `json:"isApprovalPolicyNeeded,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of the protection policy.
	*/
	Name *string `json:"name"`
	/*
	  External identifier of the owner of the protection policy.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  Specifies the connections between various replication locations and its schedule. Connections from both source-to-target and target-to-source should be specified.
	*/
	ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations"`
	/*
	  Indicates all the locations participating in the protection policy. You can specify up to 3 replication locations.
	*/
	ReplicationLocations []ReplicationLocation `json:"replicationLocations"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ProtectionPolicy) MarshalJSON() ([]byte, error) {
	type ProtectionPolicyProxy ProtectionPolicy
	return json.Marshal(struct {
		*ProtectionPolicyProxy
		Name                      *string                    `json:"name,omitempty"`
		ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations,omitempty"`
		ReplicationLocations      []ReplicationLocation      `json:"replicationLocations,omitempty"`
	}{
		ProtectionPolicyProxy:     (*ProtectionPolicyProxy)(p),
		Name:                      p.Name,
		ReplicationConfigurations: p.ReplicationConfigurations,
		ReplicationLocations:      p.ReplicationLocations,
	})
}

func NewProtectionPolicy() *ProtectionPolicy {
	p := new(ProtectionPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.ProtectionPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ProtectionPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specifies the list of external identifiers of categories that must be added to the protection policy. This policy will protect any VM or volume group associated with this category.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  Description of the protection policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Manual deletion of recovery points created by this policy can be driven through a multi-party authorization workflow. Hence, multiple approvers would be required to grant approvals for the delete operation.
	*/
	IsApprovalPolicyNeeded *bool `json:"isApprovalPolicyNeeded,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  Name of the protection policy.
	*/
	Name *string `json:"name"`
	/*
	  External identifier of the owner of the protection policy.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  Specifies the connections between various replication locations and its schedule. Connections from both source-to-target and target-to-source should be specified.
	*/
	ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations"`
	/*
	  Indicates all the locations participating in the protection policy. You can specify up to 3 replication locations.
	*/
	ReplicationLocations []ReplicationLocation `json:"replicationLocations"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ProtectionPolicyProjection) MarshalJSON() ([]byte, error) {
	type ProtectionPolicyProjectionProxy ProtectionPolicyProjection
	return json.Marshal(struct {
		*ProtectionPolicyProjectionProxy
		Name                      *string                    `json:"name,omitempty"`
		ReplicationConfigurations []ReplicationConfiguration `json:"replicationConfigurations,omitempty"`
		ReplicationLocations      []ReplicationLocation      `json:"replicationLocations,omitempty"`
	}{
		ProtectionPolicyProjectionProxy: (*ProtectionPolicyProjectionProxy)(p),
		Name:                            p.Name,
		ReplicationConfigurations:       p.ReplicationConfigurations,
		ReplicationLocations:            p.ReplicationLocations,
	})
}

func NewProtectionPolicyProjection() *ProtectionPolicyProjection {
	p := new(ProtectionPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.ProtectionPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specifies the connections between various replication locations and its schedule. Connections from both source-to-target and target-to-source should be specified.
*/
type ReplicationConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Label of the source location from the replication locations list, where the entity will be replicated.
	*/
	RemoteLocationLabel *string `json:"remoteLocationLabel,omitempty"`

	Schedule *Schedule `json:"schedule"`
	/*
	  Label of the source location from the replication locations list, where the entity is running. The location of type MST can not be specified as the replication source.
	*/
	SourceLocationLabel *string `json:"sourceLocationLabel"`
}

func (p *ReplicationConfiguration) MarshalJSON() ([]byte, error) {
	type ReplicationConfigurationProxy ReplicationConfiguration
	return json.Marshal(struct {
		*ReplicationConfigurationProxy
		Schedule            *Schedule `json:"schedule,omitempty"`
		SourceLocationLabel *string   `json:"sourceLocationLabel,omitempty"`
	}{
		ReplicationConfigurationProxy: (*ReplicationConfigurationProxy)(p),
		Schedule:                      p.Schedule,
		SourceLocationLabel:           p.SourceLocationLabel,
	})
}

func NewReplicationConfiguration() *ReplicationConfiguration {
	p := new(ReplicationConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.ReplicationConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates all the locations participating in the protection policy. You can specify up to 3 replication locations.
*/
type ReplicationLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the domain manager.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
	/*
	  One of the locations must be specified as the primary location. All the other locations must be connected to the primary location.
	*/
	IsPrimary *bool `json:"isPrimary,omitempty"`
	/*
	  This is a unique user defined label of the replication location. It is used to identify the location in the replication configurations.
	*/
	Label *string `json:"label"`
	/*

	 */
	ReplicationSubLocationItemDiscriminator_ *string `json:"$replicationSubLocationItemDiscriminator,omitempty"`
	/*
	  Specifies the replication sublocations where recovery points can be created or replicated.
	*/
	ReplicationSubLocation *OneOfReplicationLocationReplicationSubLocation `json:"replicationSubLocation,omitempty"`
}

func (p *ReplicationLocation) MarshalJSON() ([]byte, error) {
	type ReplicationLocationProxy ReplicationLocation
	return json.Marshal(struct {
		*ReplicationLocationProxy
		DomainManagerExtId *string `json:"domainManagerExtId,omitempty"`
		Label              *string `json:"label,omitempty"`
	}{
		ReplicationLocationProxy: (*ReplicationLocationProxy)(p),
		DomainManagerExtId:       p.DomainManagerExtId,
		Label:                    p.Label,
	})
}

func NewReplicationLocation() *ReplicationLocation {
	p := new(ReplicationLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.ReplicationLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReplicationLocation) GetReplicationSubLocation() interface{} {
	if nil == p.ReplicationSubLocation {
		return nil
	}
	return p.ReplicationSubLocation.GetValue()
}

func (p *ReplicationLocation) SetReplicationSubLocation(v interface{}) error {
	if nil == p.ReplicationSubLocation {
		p.ReplicationSubLocation = NewOneOfReplicationLocationReplicationSubLocation()
	}
	e := p.ReplicationSubLocation.SetValue(v)
	if nil == e {
		if nil == p.ReplicationSubLocationItemDiscriminator_ {
			p.ReplicationSubLocationItemDiscriminator_ = new(string)
		}
		*p.ReplicationSubLocationItemDiscriminator_ = *p.ReplicationSubLocation.Discriminator
	}
	return e
}

/*
Schedule for protection. The schedule specifies the recovery point objective and the retention policy for the participating locations.
*/
type Schedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The Recovery point objective of the schedule in seconds and specified in multiple of 60 seconds. Only following RPO values can be provided for rollup retention type:<br> Minute(s): 1, 2, 3, 4, 5, 6, 10, 12, 15 <br> Hour(s): 1, 2, 3, 4, 6, 8, 12 <br> Day(s): 1 <br> Week(s): 1, 2
	*/
	RecoveryPointObjectiveTimeSeconds *int `json:"recoveryPointObjectiveTimeSeconds"`

	RecoveryPointType *import4.RecoveryPointType `json:"recoveryPointType,omitempty"`
	/*

	 */
	RetentionItemDiscriminator_ *string `json:"$retentionItemDiscriminator,omitempty"`
	/*
	  Specifies the retention policy for the recovery point schedule.
	*/
	Retention *OneOfScheduleRetention `json:"retention,omitempty"`
	/*
	  Represents the protection start time for the new entities added to the policy after the policy is created in h:m format. The values must be between 00h:00m and 23h:59m and in UTC timezone. It specifies the time when the first snapshot is taken and replicated for any entity added to the policy. If this is not specified, the snapshot is taken immediately and replicated for any new entity added to the policy.
	*/
	StartTime *string `json:"startTime,omitempty"`
	/*
	  Auto suspend timeout if there is a connection failure between locations for synchronous replication. If this value is not set, then the policy will not be suspended.
	*/
	SyncReplicationAutoSuspendTimeoutSeconds *int `json:"syncReplicationAutoSuspendTimeoutSeconds,omitempty"`
}

func (p *Schedule) MarshalJSON() ([]byte, error) {
	type ScheduleProxy Schedule
	return json.Marshal(struct {
		*ScheduleProxy
		RecoveryPointObjectiveTimeSeconds *int `json:"recoveryPointObjectiveTimeSeconds,omitempty"`
	}{
		ScheduleProxy:                     (*ScheduleProxy)(p),
		RecoveryPointObjectiveTimeSeconds: p.RecoveryPointObjectiveTimeSeconds,
	})
}

func NewSchedule() *Schedule {
	p := new(Schedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.Schedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Schedule) GetRetention() interface{} {
	if nil == p.Retention {
		return nil
	}
	return p.Retention.GetValue()
}

func (p *Schedule) SetRetention(v interface{}) error {
	if nil == p.Retention {
		p.Retention = NewOneOfScheduleRetention()
	}
	e := p.Retention.SetValue(v)
	if nil == e {
		if nil == p.RetentionItemDiscriminator_ {
			p.RetentionItemDiscriminator_ = new(string)
		}
		*p.RetentionItemDiscriminator_ = *p.Retention.Discriminator
	}
	return e
}

/*
Snapshot interval period.
*/
type SnapshotIntervalType int

const (
	SNAPSHOTINTERVALTYPE_UNKNOWN  SnapshotIntervalType = 0
	SNAPSHOTINTERVALTYPE_REDACTED SnapshotIntervalType = 1
	SNAPSHOTINTERVALTYPE_HOURLY   SnapshotIntervalType = 2
	SNAPSHOTINTERVALTYPE_DAILY    SnapshotIntervalType = 3
	SNAPSHOTINTERVALTYPE_WEEKLY   SnapshotIntervalType = 4
	SNAPSHOTINTERVALTYPE_MONTHLY  SnapshotIntervalType = 5
	SNAPSHOTINTERVALTYPE_YEARLY   SnapshotIntervalType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnapshotIntervalType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SnapshotIntervalType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SnapshotIntervalType) index(name string) SnapshotIntervalType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
	for idx := range names {
		if names[idx] == name {
			return SnapshotIntervalType(idx)
		}
	}
	return SNAPSHOTINTERVALTYPE_UNKNOWN
}

func (e *SnapshotIntervalType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SnapshotIntervalType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SnapshotIntervalType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SnapshotIntervalType) Ref() *SnapshotIntervalType {
	return &e
}

/*
REST response for all response codes in API path /datapolicies/v4.0/config/protection-policies/{extId} Put operation
*/
type UpdateProtectionPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateProtectionPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateProtectionPolicyApiResponse() *UpdateProtectionPolicyApiResponse {
	p := new(UpdateProtectionPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "datapolicies.v4.config.UpdateProtectionPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateProtectionPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateProtectionPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateProtectionPolicyApiResponseData()
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

type OneOfDeleteProtectionPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteProtectionPolicyApiResponseData() *OneOfDeleteProtectionPolicyApiResponseData {
	p := new(OneOfDeleteProtectionPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteProtectionPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteProtectionPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteProtectionPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "datapolicies.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteProtectionPolicyApiResponseData"))
}

func (p *OneOfDeleteProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteProtectionPolicyApiResponseData")
}

type OneOfGetProtectionPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *ProtectionPolicy      `json:"-"`
}

func NewOneOfGetProtectionPolicyApiResponseData() *OneOfGetProtectionPolicyApiResponseData {
	p := new(OneOfGetProtectionPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetProtectionPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetProtectionPolicyApiResponseData is nil"))
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
	case ProtectionPolicy:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ProtectionPolicy)
		}
		*p.oneOfType2001 = v.(ProtectionPolicy)
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

func (p *OneOfGetProtectionPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "datapolicies.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(ProtectionPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "datapolicies.v4.config.ProtectionPolicy" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ProtectionPolicy)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetProtectionPolicyApiResponseData"))
}

func (p *OneOfGetProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetProtectionPolicyApiResponseData")
}

type OneOfListProtectionPoliciesApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType401  []ProtectionPolicyProjection `json:"-"`
	oneOfType2001 []ProtectionPolicy           `json:"-"`
	oneOfType400  *import2.ErrorResponse       `json:"-"`
}

func NewOneOfListProtectionPoliciesApiResponseData() *OneOfListProtectionPoliciesApiResponseData {
	p := new(OneOfListProtectionPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListProtectionPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListProtectionPoliciesApiResponseData is nil"))
	}
	switch v.(type) {
	case []ProtectionPolicyProjection:
		p.oneOfType401 = v.([]ProtectionPolicyProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<datapolicies.v4.config.ProtectionPolicyProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<datapolicies.v4.config.ProtectionPolicyProjection>"
	case []ProtectionPolicy:
		p.oneOfType2001 = v.([]ProtectionPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<datapolicies.v4.config.ProtectionPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<datapolicies.v4.config.ProtectionPolicy>"
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

func (p *OneOfListProtectionPoliciesApiResponseData) GetValue() interface{} {
	if "List<datapolicies.v4.config.ProtectionPolicyProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<datapolicies.v4.config.ProtectionPolicy>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListProtectionPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]ProtectionPolicyProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "datapolicies.v4.config.ProtectionPolicyProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<datapolicies.v4.config.ProtectionPolicyProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<datapolicies.v4.config.ProtectionPolicyProjection>"
			return nil
		}
	}
	vOneOfType2001 := new([]ProtectionPolicy)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "datapolicies.v4.config.ProtectionPolicy" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<datapolicies.v4.config.ProtectionPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<datapolicies.v4.config.ProtectionPolicy>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "datapolicies.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListProtectionPoliciesApiResponseData"))
}

func (p *OneOfListProtectionPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<datapolicies.v4.config.ProtectionPolicyProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<datapolicies.v4.config.ProtectionPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListProtectionPoliciesApiResponseData")
}

type OneOfReplicationLocationReplicationSubLocation struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType2101 *NutanixCluster `json:"-"`
}

func NewOneOfReplicationLocationReplicationSubLocation() *OneOfReplicationLocationReplicationSubLocation {
	p := new(OneOfReplicationLocationReplicationSubLocation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReplicationLocationReplicationSubLocation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReplicationLocationReplicationSubLocation is nil"))
	}
	switch v.(type) {
	case NutanixCluster:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(NutanixCluster)
		}
		*p.oneOfType2101 = v.(NutanixCluster)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2101.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2101.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfReplicationLocationReplicationSubLocation) GetValue() interface{} {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfReplicationLocationReplicationSubLocation) UnmarshalJSON(b []byte) error {
	vOneOfType2101 := new(NutanixCluster)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "datapolicies.v4.config.NutanixCluster" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(NutanixCluster)
			}
			*p.oneOfType2101 = *vOneOfType2101
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2101.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2101.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReplicationLocationReplicationSubLocation"))
}

func (p *OneOfReplicationLocationReplicationSubLocation) MarshalJSON() ([]byte, error) {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfReplicationLocationReplicationSubLocation")
}

type OneOfScheduleRetention struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType2002 *AutoRollupRetention `json:"-"`
	oneOfType2001 *LinearRetention     `json:"-"`
}

func NewOneOfScheduleRetention() *OneOfScheduleRetention {
	p := new(OneOfScheduleRetention)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfScheduleRetention) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfScheduleRetention is nil"))
	}
	switch v.(type) {
	case AutoRollupRetention:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(AutoRollupRetention)
		}
		*p.oneOfType2002 = v.(AutoRollupRetention)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case LinearRetention:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LinearRetention)
		}
		*p.oneOfType2001 = v.(LinearRetention)
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

func (p *OneOfScheduleRetention) GetValue() interface{} {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfScheduleRetention) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new(AutoRollupRetention)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "datapolicies.v4.config.AutoRollupRetention" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(AutoRollupRetention)
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
	vOneOfType2001 := new(LinearRetention)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "datapolicies.v4.config.LinearRetention" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LinearRetention)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfScheduleRetention"))
}

func (p *OneOfScheduleRetention) MarshalJSON() ([]byte, error) {
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfScheduleRetention")
}

type OneOfUpdateProtectionPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateProtectionPolicyApiResponseData() *OneOfUpdateProtectionPolicyApiResponseData {
	p := new(OneOfUpdateProtectionPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateProtectionPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateProtectionPolicyApiResponseData is nil"))
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

func (p *OneOfUpdateProtectionPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "datapolicies.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateProtectionPolicyApiResponseData"))
}

func (p *OneOfUpdateProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateProtectionPolicyApiResponseData")
}

type OneOfCreateProtectionPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateProtectionPolicyApiResponseData() *OneOfCreateProtectionPolicyApiResponseData {
	p := new(OneOfCreateProtectionPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateProtectionPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateProtectionPolicyApiResponseData is nil"))
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

func (p *OneOfCreateProtectionPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateProtectionPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "datapolicies.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateProtectionPolicyApiResponseData"))
}

func (p *OneOfCreateProtectionPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateProtectionPolicyApiResponseData")
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
