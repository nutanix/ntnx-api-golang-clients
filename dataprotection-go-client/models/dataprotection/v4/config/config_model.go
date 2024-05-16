/*
 * Generated file models/dataprotection/v4/config/config_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Dataprotection Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	import5 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/common"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/prism/v4/config"
	"time"
)

/*
Protected resource/recovery point restore that overrides the AHV VM configuration. The specified properties will be overridden for the restored VM.
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the Amazon S3 bucket.
*/
type AmazonS3Bucket struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the Amazon S3 bucket.
	*/
	BucketName *string `json:"bucketName,omitempty"`
	/*
	  Name of the AWS region.
	*/
	RegionName *string `json:"regionName,omitempty"`
}

func NewAmazonS3Bucket() *AmazonS3Bucket {
	p := new(AmazonS3Bucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AmazonS3Bucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the Microsoft Azure Blob Storage container.
*/
type AzureBlobStorageContainer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the Microsoft Azure Blob Storage container.
	*/
	ContainerName *string `json:"containerName,omitempty"`
	/*
	  Name of the Microsoft Azure Blob Storage account.
	*/
	StorageAccountName *string `json:"storageAccountName,omitempty"`
}

func NewAzureBlobStorageContainer() *AzureBlobStorageContainer {
	p := new(AzureBlobStorageContainer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AzureBlobStorageContainer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response parameters for the compute changed regions. It is a chunk-encoded stream for the redirect request. It can contain a redirect IP, a certificate, or the actual changed regions. If the response contains a redirect IP, the client must re-send the request to the specified IP. The certificate needs to be set as NTNX_IGW_SESSION cookie in the header.
*/
type ComputeChangedRegionsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ResponseItemDiscriminator_ *string `json:"$responseItemDiscriminator,omitempty"`

	Response *OneOfComputeChangedRegionsResponseResponse `json:"response,omitempty"`
}

func NewComputeChangedRegionsResponse() *ComputeChangedRegionsResponse {
	p := new(ComputeChangedRegionsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ComputeChangedRegionsResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ComputeChangedRegionsResponse) GetResponse() interface{} {
	if nil == p.Response {
		return nil
	}
	return p.Response.GetValue()
}

func (p *ComputeChangedRegionsResponse) SetResponse(v interface{}) error {
	if nil == p.Response {
		p.Response = NewOneOfComputeChangedRegionsResponseResponse()
	}
	e := p.Response.SetValue(v)
	if nil == e {
		if nil == p.ResponseItemDiscriminator_ {
			p.ResponseItemDiscriminator_ = new(string)
		}
		*p.ResponseItemDiscriminator_ = *p.Response.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/disk-recovery-points/{extId}/$actions/compute-changed-regions Post operation
*/
type ComputeChangedRegionsResponseModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfComputeChangedRegionsResponseModelData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewComputeChangedRegionsResponseModel() *ComputeChangedRegionsResponseModel {
	p := new(ComputeChangedRegionsResponseModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ComputeChangedRegionsResponseModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ComputeChangedRegionsResponseModel) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ComputeChangedRegionsResponseModel) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfComputeChangedRegionsResponseModelData()
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
There are many scenarios where the state of an application must be captured as an aggregate of the internal state of a set of related entities at some point in time. The Consistency group is a collection of all the entities whose snapshot represents the application state at that point in time
*/
type ConsistencyGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the cluster to which the entities from the Consistency group are associated.
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

	Members []ConsistencyGroupMember `json:"members"`
	/*
	  Name of the Consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this Consistency group. This is a read-only field that is inserted into the Consistency group entity at the time of Consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  The external identifier of the Protection policy that protects the Consistency group.
	*/
	ProtectionPolicyExtId *string `json:"protectionPolicyExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroup) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProxy ConsistencyGroup
	return json.Marshal(struct {
		*ConsistencyGroupProxy
		Members []ConsistencyGroupMember `json:"members,omitempty"`
		Name    *string                  `json:"name,omitempty"`
	}{
		ConsistencyGroupProxy: (*ConsistencyGroupProxy)(p),
		Members:               p.Members,
		Name:                  p.Name,
	})
}

func NewConsistencyGroup() *ConsistencyGroup {
	p := new(ConsistencyGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	EntityExtId *string `json:"entityExtId,omitempty"`

	EntityType *ConsistencyGroupMemberType `json:"entityType"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
Specification for the migrate action on the Consistency group.
*/
type ConsistencyGroupMigrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference to the cluster in the target Availability Zone where the entities need to be migrated.
	*/
	TargetClusterExtId *string `json:"targetClusterExtId,omitempty"`
	/*
	  Reference to the target Availability Zone where the entities need to be migrated.
	*/
	TargetPcExtId *string `json:"targetPcExtId,omitempty"`
}

func NewConsistencyGroupMigrationSpec() *ConsistencyGroupMigrationSpec {
	p := new(ConsistencyGroupMigrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMigrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConsistencyGroupProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the cluster to which the entities from the Consistency group are associated.
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

	Members []ConsistencyGroupMember `json:"members"`
	/*
	  Name of the Consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this Consistency group. This is a read-only field that is inserted into the Consistency group entity at the time of Consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  The external identifier of the Protection policy that protects the Consistency group.
	*/
	ProtectionPolicyExtId *string `json:"protectionPolicyExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroupProjection) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProjectionProxy ConsistencyGroupProjection
	return json.Marshal(struct {
		*ConsistencyGroupProjectionProxy
		Members []ConsistencyGroupMember `json:"members,omitempty"`
		Name    *string                  `json:"name,omitempty"`
	}{
		ConsistencyGroupProjectionProxy: (*ConsistencyGroupProjectionProxy)(p),
		Members:                         p.Members,
		Name:                            p.Name,
	})
}

func NewConsistencyGroupProjection() *ConsistencyGroupProjection {
	p := new(ConsistencyGroupProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/consistency-groups Post operation
*/
type CreateConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateConsistencyGroupApiResponse() *CreateConsistencyGroupApiResponse {
	p := new(CreateConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreateConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateConsistencyGroupApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	MgmtClusterExtId *string `json:"mgmtClusterExtId,omitempty"`
}

func NewDataProtectionSiteReference() *DataProtectionSiteReference {
	p := new(DataProtectionSiteReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionSiteReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/consistency-groups/{extId} Delete operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{extId} Delete operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents the disk recovery point properties.
*/
type DiskRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk external identifier which is captured as part of this recovery point.
	*/
	DiskExtId *string `json:"diskExtId,omitempty"`
	/*
	  Disk recovery point Identifier.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
}

func NewDiskRecoveryPoint() *DiskRecoveryPoint {
	p := new(DiskRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DiskRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Protected resource/recovery point restore that overrides the ESXi VM configuration. The specified properties will be overridden for the restored VM.
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification to set the expiration time of the recovery point.
*/
type ExpirationTimeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current recovery point expires . If not specified, the recovery point will never expire.
	*/
	ExpirationTime *time.Time `json:"expirationTime"`
}

func (p *ExpirationTimeSpec) MarshalJSON() ([]byte, error) {
	type ExpirationTimeSpecProxy ExpirationTimeSpec
	return json.Marshal(struct {
		*ExpirationTimeSpecProxy
		ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	}{
		ExpirationTimeSpecProxy: (*ExpirationTimeSpecProxy)(p),
		ExpirationTime:          p.ExpirationTime,
	})
}

func NewExpirationTimeSpec() *ExpirationTimeSpec {
	p := new(ExpirationTimeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ExpirationTimeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/protected-resources/{extId} Get operation
*/
type GetProtectedResourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetProtectedResourceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetProtectedResourceApiResponse() *GetProtectedResourceApiResponse {
	p := new(GetProtectedResourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetProtectedResourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetProtectedResourceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetProtectedResourceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetProtectedResourceApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{extId} Get operation
*/
type GetRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRecoveryPointApiResponse() *GetRecoveryPointApiResponse {
	p := new(GetRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRecoveryPointApiResponseData()
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness host site.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points Get operation
*/
type ListRecoveryPointsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecoveryPointsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRecoveryPointsApiResponse() *ListRecoveryPointsApiResponse {
	p := new(ListRecoveryPointsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListRecoveryPointsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecoveryPointsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecoveryPointsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecoveryPointsApiResponseData()
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
Location reference where the specified recovery point is present.
*/
type LocationReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster where the recovery point is present.
	*/
	LocationExtId *string `json:"locationExtId,omitempty"`
}

func NewLocationReference() *LocationReference {
	p := new(LocationReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.LocationReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/consistency-groups/{extId}/$actions/migrate Post operation
*/
type MigrateConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMigrateConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMigrateConsistencyGroupApiResponse() *MigrateConsistencyGroupApiResponse {
	p := new(MigrateConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.MigrateConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MigrateConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MigrateConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMigrateConsistencyGroupApiResponseData()
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
Reference to the Nutanix Objects bucket.
*/
type NutanixObjectsBucket struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  URI of the Nutanix Objects bucket.
	*/
	EndPoint *string `json:"endPoint,omitempty"`
}

func NewNutanixObjectsBucket() *NutanixObjectsBucket {
	p := new(NutanixObjectsBucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.NutanixObjectsBucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the S3-compatible object storage.
*/
type ObjectStorageType int

const (
	OBJECTSTORAGETYPE_UNKNOWN            ObjectStorageType = 0
	OBJECTSTORAGETYPE_REDACTED           ObjectStorageType = 1
	OBJECTSTORAGETYPE_AZURE_BLOB_STORAGE ObjectStorageType = 2
	OBJECTSTORAGETYPE_AMAZON_S3          ObjectStorageType = 3
	OBJECTSTORAGETYPE_NUTANIX_OBJECTS    ObjectStorageType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ObjectStorageType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE_BLOB_STORAGE",
		"AMAZON_S3",
		"NUTANIX_OBJECTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ObjectStorageType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE_BLOB_STORAGE",
		"AMAZON_S3",
		"NUTANIX_OBJECTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ObjectStorageType) index(name string) ObjectStorageType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE_BLOB_STORAGE",
		"AMAZON_S3",
		"NUTANIX_OBJECTS",
	}
	for idx := range names {
		if names[idx] == name {
			return ObjectStorageType(idx)
		}
	}
	return OBJECTSTORAGETYPE_UNKNOWN
}

func (e *ObjectStorageType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ObjectStorageType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ObjectStorageType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ObjectStorageType) Ref() *ObjectStorageType {
	return &e
}

/*
Protected resource entity type.
*/
type ProtectedEntityType int

const (
	PROTECTEDENTITYTYPE_UNKNOWN      ProtectedEntityType = 0
	PROTECTEDENTITYTYPE_REDACTED     ProtectedEntityType = 1
	PROTECTEDENTITYTYPE_VM           ProtectedEntityType = 2
	PROTECTEDENTITYTYPE_VOLUME_GROUP ProtectedEntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtectedEntityType) name(index int) string {
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
func (e ProtectedEntityType) GetName() string {
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
func (e *ProtectedEntityType) index(name string) ProtectedEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return ProtectedEntityType(idx)
		}
	}
	return PROTECTEDENTITYTYPE_UNKNOWN
}

func (e *ProtectedEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProtectedEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProtectedEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProtectedEntityType) Ref() *ProtectedEntityType {
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
	  Category key-value pairs associated with the protected resource at the time of protection. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	CategoryFqNames []string `json:"categoryFqNames,omitempty"`
	/*
	  The external identifier of the VM or the volume group associated with the protected resource.
	*/
	EntityExtId *string `json:"entityExtId,omitempty"`

	EntityType *ProtectedEntityType `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	ReplicationStates []ReplicationState `json:"replicationStates,omitempty"`
	/*
	  The data protection details for the protected resource that are relevant to any of the sites in the local Prism Central, like the time ranges available for recovery.
	*/
	SiteProtectionInfo []SiteProtectionInfo `json:"siteProtectionInfo,omitempty"`

	SourceSiteReference *DataProtectionSiteReference `json:"sourceSiteReference,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewProtectedResource() *ProtectedResource {
	p := new(ProtectedResource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResource"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/protected-resources/{extId}/$actions/promote Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
Status of replication to a specified target site.
*/
type ProtectedResourceReplicationStatus int

const (
	PROTECTEDRESOURCEREPLICATIONSTATUS_UNKNOWN     ProtectedResourceReplicationStatus = 0
	PROTECTEDRESOURCEREPLICATIONSTATUS_REDACTED    ProtectedResourceReplicationStatus = 1
	PROTECTEDRESOURCEREPLICATIONSTATUS_IN_SYNC     ProtectedResourceReplicationStatus = 2
	PROTECTEDRESOURCEREPLICATIONSTATUS_SYNCING     ProtectedResourceReplicationStatus = 3
	PROTECTEDRESOURCEREPLICATIONSTATUS_OUT_OF_SYNC ProtectedResourceReplicationStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProtectedResourceReplicationStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_SYNC",
		"SYNCING",
		"OUT_OF_SYNC",
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
		"IN_SYNC",
		"SYNCING",
		"OUT_OF_SYNC",
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
		"IN_SYNC",
		"SYNCING",
		"OUT_OF_SYNC",
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/protected-resources/{extId}/$actions/restore Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

	 */
	OverridesItemDiscriminator_ *string `json:"$overridesItemDiscriminator,omitempty"`
	/*
	  The restore action specifications on the protected resource. This is an optional field and the specified properties will be overridden for the restored VM/VG.
	*/
	Overrides *OneOfProtectedResourceRestoreSpecOverrides `json:"overrides,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ProtectedResourceRestoreSpec) GetOverrides() interface{} {
	if nil == p.Overrides {
		return nil
	}
	return p.Overrides.GetValue()
}

func (p *ProtectedResourceRestoreSpec) SetOverrides(v interface{}) error {
	if nil == p.Overrides {
		p.Overrides = NewOneOfProtectedResourceRestoreSpecOverrides()
	}
	e := p.Overrides.SetValue(v)
	if nil == e {
		if nil == p.OverridesItemDiscriminator_ {
			p.OverridesItemDiscriminator_ = new(string)
		}
		*p.OverridesItemDiscriminator_ = *p.Overrides.Discriminator
	}
	return e
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the recovery point along with all the captured VM and volume group recovery point(s).
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  List of location references where the VM or volume group recovery point are a part of the specified recovery point.
	*/
	LocationReferences []LocationReference `json:"locationReferences,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A read only field inserted into recovery point at the time of recovery point creation, indicating the external identifier of the user who created this recovery point.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`

	RecoveryPointType *import4.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import4.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of VM recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 30 entities. These entities can be a combination of VM(s) and volume group(s).
	*/
	VmRecoveryPoints []VmRecoveryPoint `json:"vmRecoveryPoints,omitempty"`
	/*
	  List of volume group recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 30 entities. These entities can be a combination of VM(s) and volume group(s).
	*/
	VolumeGroupRecoveryPoints []VolumeGroupRecoveryPoint `json:"volumeGroupRecoveryPoints,omitempty"`
}

func NewRecoveryPoint() *RecoveryPoint {
	p := new(RecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RecoveryPointProjection struct {
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  List of location references where the VM or volume group recovery point are a part of the specified recovery point.
	*/
	LocationReferences []LocationReference `json:"locationReferences,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A read only field inserted into recovery point at the time of recovery point creation, indicating the external identifier of the user who created this recovery point.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`

	RecoveryPointType *import4.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import4.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  List of VM recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 30 entities. These entities can be a combination of VM(s) and volume group(s).
	*/
	VmRecoveryPoints []VmRecoveryPoint `json:"vmRecoveryPoints,omitempty"`
	/*
	  List of volume group recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 30 entities. These entities can be a combination of VM(s) and volume group(s).
	*/
	VolumeGroupRecoveryPoints []VolumeGroupRecoveryPoint `json:"volumeGroupRecoveryPoints,omitempty"`
}

func NewRecoveryPointProjection() *RecoveryPointProjection {
	p := new(RecoveryPointProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{extId}/$actions/replicate Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
External identifier of the cluster and the Prism Central where the recovery point should be replicated.
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
}

func NewRecoveryPointReplicationSpec() *RecoveryPointReplicationSpec {
	p := new(RecoveryPointReplicationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointReplicationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A recovery point repository is a reference to a data store which can be used to store recovery points using Multicloud snapshot technology (MST). It can be backed by any S3-compatible object storage solutions, which include options like Nutanix Objects, as well as hyperscaler offerings such as Amazon AWS S3 and Microsoft Azure Blobs.
*/
type RecoveryPointRepository struct {
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
	/*
	  Name of the recovery point repository.
	*/
	Name *string `json:"name,omitempty"`
	/*

	 */
	ObjectStorageReferenceItemDiscriminator_ *string `json:"$objectStorageReferenceItemDiscriminator,omitempty"`

	ObjectStorageReference *OneOfRecoveryPointRepositoryObjectStorageReference `json:"objectStorageReference,omitempty"`

	ObjectStorageType *ObjectStorageType `json:"objectStorageType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRecoveryPointRepository() *RecoveryPointRepository {
	p := new(RecoveryPointRepository)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRepository"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoveryPointRepository) GetObjectStorageReference() interface{} {
	if nil == p.ObjectStorageReference {
		return nil
	}
	return p.ObjectStorageReference.GetValue()
}

func (p *RecoveryPointRepository) SetObjectStorageReference(v interface{}) error {
	if nil == p.ObjectStorageReference {
		p.ObjectStorageReference = NewOneOfRecoveryPointRepositoryObjectStorageReference()
	}
	e := p.ObjectStorageReference.SetValue(v)
	if nil == e {
		if nil == p.ObjectStorageReferenceItemDiscriminator_ {
			p.ObjectStorageReferenceItemDiscriminator_ = new(string)
		}
		*p.ObjectStorageReferenceItemDiscriminator_ = *p.ObjectStorageReference.Discriminator
	}
	return e
}

type RecoveryPointRepositoryProjection struct {
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
	/*
	  Name of the recovery point repository.
	*/
	Name *string `json:"name,omitempty"`

	ObjectStorageReferenceItemDiscriminator_ *string `json:"$objectStorageReferenceItemDiscriminator,omitempty"`

	ObjectStorageReference *OneOfRecoveryPointRepositoryProjectionObjectStorageReference `json:"objectStorageReference,omitempty"`

	ObjectStorageType *ObjectStorageType `json:"objectStorageType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRecoveryPointRepositoryProjection() *RecoveryPointRepositoryProjection {
	p := new(RecoveryPointRepositoryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRepositoryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for the restore action on the top-level recovery point. For a recovery point that contains multiple VM or volume group recovery points, users can selectively trigger restore for any specific set of VM or volume group recovery point(s). In case no particular selection is made, all VM and volume group recovery points that are a part of the top-level recovery point will be restored.
*/
type RecoveryPointRestorationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Recovery points are restored at the associated location reference by default. However, there is no particular location reference associated with recovery points located on the cloud. In such a case, the client must specify the external identifier of the cluster on which the entity should be restored.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  List of specifications to restore a specific VM recovery point(s) that are a part of the top-level recovery point. A specific VM recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VmRecoveryPointRestoreOverrides []VmRecoveryPointRestoreOverride `json:"vmRecoveryPointRestoreOverrides,omitempty"`
	/*
	  List of specifications to restore a specific volume group recovery point(s) that are a part of the top-level recovery point. A specific volume group recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VolumeGroupRecoveryPointRestoreOverrides []VolumeGroupRecoveryPointRestoreOverride `json:"volumeGroupRecoveryPointRestoreOverrides,omitempty"`
}

func NewRecoveryPointRestorationSpec() *RecoveryPointRestorationSpec {
	p := new(RecoveryPointRestorationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRestorationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{extId}/$actions/restore Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

type RedirectionResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A special token in string format. This token must be set as a session cookie in the redirect call.
	*/
	Certificate *string `json:"certificate,omitempty"`

	RedirectIp *IpAddress `json:"redirectIp,omitempty"`
}

func NewRedirectionResponse() *RedirectionResponse {
	p := new(RedirectionResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RedirectionResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Changed region comprising of offset, length, and type of the region.
*/
type Region struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The length of the region in bytes.
	*/
	Length *int64 `json:"length"`
	/*
	  The start offset of the region in bytes.
	*/
	Offset *int64 `json:"offset"`

	RegionType *RegionType `json:"regionType"`
}

func (p *Region) MarshalJSON() ([]byte, error) {
	type RegionProxy Region
	return json.Marshal(struct {
		*RegionProxy
		Length     *int64      `json:"length,omitempty"`
		Offset     *int64      `json:"offset,omitempty"`
		RegionType *RegionType `json:"regionType,omitempty"`
	}{
		RegionProxy: (*RegionProxy)(p),
		Length:      p.Length,
		Offset:      p.Offset,
		RegionType:  p.RegionType,
	})
}

func NewRegion() *Region {
	p := new(Region)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Region"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RegionResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The offset from where the client must continue the request. This field is not set when there are no more changed regions to be returned. Note: the nextOffset can be outside the endOffset specified by the client in the request. This helps clients reach the next changed offset faster.
	*/
	NextOffset *int64 `json:"nextOffset,omitempty"`

	Regions []Region `json:"regions,omitempty"`
}

func NewRegionResponse() *RegionResponse {
	p := new(RegionResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RegionResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of the region.
*/
type RegionType int

const (
	REGIONTYPE_UNKNOWN  RegionType = 0
	REGIONTYPE_REDACTED RegionType = 1
	REGIONTYPE_ZEROED   RegionType = 2
	REGIONTYPE_REGULAR  RegionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RegionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RegionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RegionType) index(name string) RegionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	for idx := range names {
		if names[idx] == name {
			return RegionType(idx)
		}
	}
	return REGIONTYPE_UNKNOWN
}

func (e *RegionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RegionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RegionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RegionType) Ref() *RegionType {
	return &e
}

/*
Replication related information about the protected resource.
*/
type ReplicationState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the Protection policy associated with the protected resource.
	*/
	ProtectionPolicyExtId *string `json:"protectionPolicyExtId,omitempty"`
	/*
	  The recovery point objective of the schedule in seconds.
	*/
	RecoveryPointObjectiveSeconds *int64 `json:"recoveryPointObjectiveSeconds,omitempty"`

	ReplicationStatus *ProtectedResourceReplicationStatus `json:"replicationStatus,omitempty"`

	TargetSiteReference *DataProtectionSiteReference `json:"targetSiteReference,omitempty"`
}

func NewReplicationState() *ReplicationState {
	p := new(ReplicationState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ReplicationState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The data protection details for the Protected resource that are relevant to the specified local or remote site, like the time ranges available for recovery on the given site.
*/
type SiteProtectionInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LocationReference *DataProtectionSiteReference `json:"locationReference,omitempty"`

	RecoveryInfo *RecoveryInfo `json:"recoveryInfo,omitempty"`
}

func NewSiteProtectionInfo() *SiteProtectionInfo {
	p := new(SiteProtectionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SiteProtectionInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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

	Ipv4 *import5.IPv4Address `json:"ipv4,omitempty"`
}

func NewSubnet() *Subnet {
	p := new(Subnet)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Subnet"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0.b1/config/synced-volume-groups/{extId}/$actions/promote Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/consistency-groups/{extId} Put operation
*/
type UpdateConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateConsistencyGroupApiResponse() *UpdateConsistencyGroupApiResponse {
	p := new(UpdateConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.UpdateConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateConsistencyGroupApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.0.b1/config/recovery-points/{extId}/$actions/set-expiration-time Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
A model that represents VM recovery point properties.
*/
type VmRecoveryPoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DiskRecoveryPoints []DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the recovery point. This identifier is used to identify the same instances of a recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId"`
}

func (p *VmRecoveryPoint) MarshalJSON() ([]byte, error) {
	type VmRecoveryPointProxy VmRecoveryPoint
	return json.Marshal(struct {
		*VmRecoveryPointProxy
		VmExtId *string `json:"vmExtId,omitempty"`
	}{
		VmRecoveryPointProxy: (*VmRecoveryPointProxy)(p),
		VmExtId:              p.VmExtId,
	})
}

func NewVmRecoveryPoint() *VmRecoveryPoint {
	p := new(VmRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compute changed region parameters. In these parameters, you can specify a start offset, an end offset, and a reference disk recovery point. All parameters are optional. However, if a reference disk recovery point needs to be set, all three parameters - recovery point Id, VM recovery point Id, and disk recovery point Id must be specified.
*/
type VmRecoveryPointComputeChangedRegionsRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The length to compute the changed region. If the value is not provided, the difference is performed from the start offset to the end of the disk. Note: the end offset might automatically align to a system-defined block boundary.
	*/
	Length *int64 `json:"length,omitempty"`
	/*
	  The start offset value to compute the changed region. If the value is not provided, the difference is executed from the offset of 0. Note: the start offset might automatically align to a system-defined block boundary.
	*/
	Offset *int64 `json:"offset,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the VM recovery point using its URL (Note: This attribute will be removed in future releases, use VM recovery point external identifier instead).
	*/
	ReferenceVmRecoveryPointExtId *string `json:"referenceVmRecoveryPointExtId,omitempty"`
}

func NewVmRecoveryPointComputeChangedRegionsRequest() *VmRecoveryPointComputeChangedRegionsRequest {
	p := new(VmRecoveryPointComputeChangedRegionsRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointComputeChangedRegionsRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmRecoveryPointRestoreOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of a VM recovery point, that is a part of the top-level recovery point (Note: This attribute will be removed in future releases, use VM recovery point external identifier instead).
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
}

func NewVmRecoveryPointRestoreOverride() *VmRecoveryPointRestoreOverride {
	p := new(VmRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Protected resource/recovery point restore that overrides the volume group configuration. The specified properties will be overridden for the restored volume group.
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents volume group recovery point properties.
*/
type VolumeGroupRecoveryPoint struct {
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
	/*
	  Location agnostic identifier of the recovery point. This identifier is used to identify the same instances of a recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Category key-value pairs associated with the volume group at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' will be represented as 'dept/hr'.
	*/
	VolumeGroupCategories []string `json:"volumeGroupCategories,omitempty"`
	/*
	  Volume Group external identifier which is captured as part of this recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId"`
}

func (p *VolumeGroupRecoveryPoint) MarshalJSON() ([]byte, error) {
	type VolumeGroupRecoveryPointProxy VolumeGroupRecoveryPoint
	return json.Marshal(struct {
		*VolumeGroupRecoveryPointProxy
		VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	}{
		VolumeGroupRecoveryPointProxy: (*VolumeGroupRecoveryPointProxy)(p),
		VolumeGroupExtId:              p.VolumeGroupExtId,
	})
}

func NewVolumeGroupRecoveryPoint() *VolumeGroupRecoveryPoint {
	p := new(VolumeGroupRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compute changed region parameters. In these parameters, you can specify a start offset, an end offset, and a reference disk recovery point. All parameters are optional. However, if a reference disk recovery point needs to be set, all three parameters - recovery point Id, volume group recovery point Id, and disk recovery point Id must be specified.
*/
type VolumeGroupRecoveryPointComputeChangedRegionsRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The length to compute the changed region. If the value is not provided, the difference is performed from the start offset to the end of the disk. Note: the end offset might automatically align to a system-defined block boundary.
	*/
	Length *int64 `json:"length,omitempty"`
	/*
	  The start offset value to compute the changed region. If the value is not provided, the difference is executed from the offset of 0. Note: the start offset might automatically align to a system-defined block boundary.
	*/
	Offset *int64 `json:"offset,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group recovery point using its URL (Note: This attribute will be removed in future releases, use volume group recovery point external identifier instead).
	*/
	ReferenceVolumeGroupRecoveryPointExtId *string `json:"referenceVolumeGroupRecoveryPointExtId,omitempty"`
}

func NewVolumeGroupRecoveryPointComputeChangedRegionsRequest() *VolumeGroupRecoveryPointComputeChangedRegionsRequest {
	p := new(VolumeGroupRecoveryPointComputeChangedRegionsRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointComputeChangedRegionsRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VolumeGroupRecoveryPointRestoreOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	VolumeGroupOverrideSpec *VolumeGroupOverrideSpec `json:"volumeGroupOverrideSpec,omitempty"`
	/*
	  External identifier of a volume group recovery point, that is a part of the top-level recovery point.
	*/
	VolumeGroupRecoveryPointExtId *string `json:"volumeGroupRecoveryPointExtId,omitempty"`
}

func NewVolumeGroupRecoveryPointRestoreOverride() *VolumeGroupRecoveryPointRestoreOverride {
	p := new(VolumeGroupRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification corresponding to the volume group synced state. The primary site reference corresponds to the cluster reference information along with the related AZ information where the entity resides. The secondary site reference corresponds to the cluster site reference information along with the related AZ information where the volume group is synced to.
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPC reference.
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
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
	  List of IP addresses for the Witness site.
	*/
	IpAddresses []string `json:"ipAddresses,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness site.
	*/
	Name *string `json:"name,omitempty"`

	Status *WitnessAvailabilityStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWitness() *Witness {
	p := new(Witness)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Witness"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

type OneOfDeleteRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfDeleteRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRecoveryPointApiResponseData"))
}

func (p *OneOfDeleteRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRecoveryPointApiResponseData")
}

type OneOfRecoveryPointReplicateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfRecoveryPointReplicateApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecoveryPointReplicateApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointReplicateApiResponseData"))
}

func (p *OneOfRecoveryPointReplicateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointReplicateApiResponseData")
}

type OneOfRecoveryPointRepositoryProjectionObjectStorageReference struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType1    *AmazonS3Bucket            `json:"-"`
	oneOfType2    *NutanixObjectsBucket      `json:"-"`
	oneOfType0    *AzureBlobStorageContainer `json:"-"`
}

func NewOneOfRecoveryPointRepositoryProjectionObjectStorageReference() *OneOfRecoveryPointRepositoryProjectionObjectStorageReference {
	p := new(OneOfRecoveryPointRepositoryProjectionObjectStorageReference)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointRepositoryProjectionObjectStorageReference is nil"))
	}
	switch v.(type) {
	case AmazonS3Bucket:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(AmazonS3Bucket)
		}
		*p.oneOfType1 = v.(AmazonS3Bucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case NutanixObjectsBucket:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(NutanixObjectsBucket)
		}
		*p.oneOfType2 = v.(NutanixObjectsBucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case AzureBlobStorageContainer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AzureBlobStorageContainer)
		}
		*p.oneOfType0 = v.(AzureBlobStorageContainer)
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

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) GetValue() interface{} {
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

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(AmazonS3Bucket)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.AmazonS3Bucket" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(AmazonS3Bucket)
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
	vOneOfType2 := new(NutanixObjectsBucket)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "dataprotection.v4.config.NutanixObjectsBucket" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(NutanixObjectsBucket)
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
	vOneOfType0 := new(AzureBlobStorageContainer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.AzureBlobStorageContainer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AzureBlobStorageContainer)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRepositoryProjectionObjectStorageReference"))
}

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRepositoryProjectionObjectStorageReference")
}

type OneOfGetProtectedResourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ProtectedResource     `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetProtectedResourceApiResponseData() *OneOfGetProtectedResourceApiResponseData {
	p := new(OneOfGetProtectedResourceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetProtectedResourceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetProtectedResourceApiResponseData is nil"))
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

func (p *OneOfGetProtectedResourceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetProtectedResourceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetProtectedResourceApiResponseData"))
}

func (p *OneOfGetProtectedResourceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetProtectedResourceApiResponseData")
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

type OneOfUpdateRecoveryPointExpirationTimeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
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

type OneOfProtectedResourceRestoreSpecOverrides struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType2    *VolumeGroupOverrideSpec `json:"-"`
	oneOfType0    *AhvVmOverrideSpec       `json:"-"`
	oneOfType1    *EsxiVmOverrideSpec      `json:"-"`
}

func NewOneOfProtectedResourceRestoreSpecOverrides() *OneOfProtectedResourceRestoreSpecOverrides {
	p := new(OneOfProtectedResourceRestoreSpecOverrides)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfProtectedResourceRestoreSpecOverrides) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfProtectedResourceRestoreSpecOverrides is nil"))
	}
	switch v.(type) {
	case VolumeGroupOverrideSpec:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(VolumeGroupOverrideSpec)
		}
		*p.oneOfType2 = v.(VolumeGroupOverrideSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case AhvVmOverrideSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AhvVmOverrideSpec)
		}
		*p.oneOfType0 = v.(AhvVmOverrideSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case EsxiVmOverrideSpec:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(EsxiVmOverrideSpec)
		}
		*p.oneOfType1 = v.(EsxiVmOverrideSpec)
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

func (p *OneOfProtectedResourceRestoreSpecOverrides) GetValue() interface{} {
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

func (p *OneOfProtectedResourceRestoreSpecOverrides) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(VolumeGroupOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "dataprotection.v4.config.VolumeGroupOverrideSpec" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(VolumeGroupOverrideSpec)
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
	vOneOfType0 := new(AhvVmOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.AhvVmOverrideSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AhvVmOverrideSpec)
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
	vOneOfType1 := new(EsxiVmOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.EsxiVmOverrideSpec" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(EsxiVmOverrideSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourceRestoreSpecOverrides"))
}

func (p *OneOfProtectedResourceRestoreSpecOverrides) MarshalJSON() ([]byte, error) {
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourceRestoreSpecOverrides")
}

type OneOfRecoveryPointRepositoryObjectStorageReference struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType1    *AmazonS3Bucket            `json:"-"`
	oneOfType2    *NutanixObjectsBucket      `json:"-"`
	oneOfType0    *AzureBlobStorageContainer `json:"-"`
}

func NewOneOfRecoveryPointRepositoryObjectStorageReference() *OneOfRecoveryPointRepositoryObjectStorageReference {
	p := new(OneOfRecoveryPointRepositoryObjectStorageReference)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoveryPointRepositoryObjectStorageReference is nil"))
	}
	switch v.(type) {
	case AmazonS3Bucket:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(AmazonS3Bucket)
		}
		*p.oneOfType1 = v.(AmazonS3Bucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case NutanixObjectsBucket:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(NutanixObjectsBucket)
		}
		*p.oneOfType2 = v.(NutanixObjectsBucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	case AzureBlobStorageContainer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AzureBlobStorageContainer)
		}
		*p.oneOfType0 = v.(AzureBlobStorageContainer)
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

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) GetValue() interface{} {
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

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(AmazonS3Bucket)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.AmazonS3Bucket" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(AmazonS3Bucket)
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
	vOneOfType2 := new(NutanixObjectsBucket)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "dataprotection.v4.config.NutanixObjectsBucket" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(NutanixObjectsBucket)
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
	vOneOfType0 := new(AzureBlobStorageContainer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.AzureBlobStorageContainer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AzureBlobStorageContainer)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRepositoryObjectStorageReference"))
}

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRepositoryObjectStorageReference")
}

type OneOfProtectedResourceRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
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

type OneOfCreateRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfCreateRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRecoveryPointApiResponseData"))
}

func (p *OneOfCreateRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRecoveryPointApiResponseData")
}

type OneOfSynchronousReplicaPromoteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
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

type OneOfComputeChangedRegionsResponseResponse struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType0    *RegionResponse      `json:"-"`
	oneOfType1    *RedirectionResponse `json:"-"`
}

func NewOneOfComputeChangedRegionsResponseResponse() *OneOfComputeChangedRegionsResponseResponse {
	p := new(OneOfComputeChangedRegionsResponseResponse)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfComputeChangedRegionsResponseResponse) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfComputeChangedRegionsResponseResponse is nil"))
	}
	switch v.(type) {
	case RegionResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RegionResponse)
		}
		*p.oneOfType0 = v.(RegionResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case RedirectionResponse:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(RedirectionResponse)
		}
		*p.oneOfType1 = v.(RedirectionResponse)
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

func (p *OneOfComputeChangedRegionsResponseResponse) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfComputeChangedRegionsResponseResponse) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RegionResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.RegionResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RegionResponse)
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
	vOneOfType1 := new(RedirectionResponse)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.config.RedirectionResponse" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(RedirectionResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfComputeChangedRegionsResponseResponse"))
}

func (p *OneOfComputeChangedRegionsResponseResponse) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfComputeChangedRegionsResponseResponse")
}

type OneOfListRecoveryPointsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    []RecoveryPoint           `json:"-"`
	oneOfType400  *import1.ErrorResponse    `json:"-"`
	oneOfType401  []RecoveryPointProjection `json:"-"`
}

func NewOneOfListRecoveryPointsApiResponseData() *OneOfListRecoveryPointsApiResponseData {
	p := new(OneOfListRecoveryPointsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecoveryPointsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecoveryPointsApiResponseData is nil"))
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
	case []RecoveryPointProjection:
		p.oneOfType401 = v.([]RecoveryPointProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPointProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPointProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRecoveryPointsApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.RecoveryPointProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRecoveryPointsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType401 := new([]RecoveryPointProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {

		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.RecoveryPointProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPointProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPointProjection>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecoveryPointsApiResponseData"))
}

func (p *OneOfListRecoveryPointsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.RecoveryPointProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRecoveryPointsApiResponseData")
}

type OneOfMigrateConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfMigrateConsistencyGroupApiResponseData() *OneOfMigrateConsistencyGroupApiResponseData {
	p := new(OneOfMigrateConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMigrateConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMigrateConsistencyGroupApiResponseData is nil"))
	}
	switch v.(type) {
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

func (p *OneOfMigrateConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfMigrateConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMigrateConsistencyGroupApiResponseData"))
}

func (p *OneOfMigrateConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfMigrateConsistencyGroupApiResponseData")
}

type OneOfCreateConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ConsistencyGroup      `json:"-"`
}

func NewOneOfCreateConsistencyGroupApiResponseData() *OneOfCreateConsistencyGroupApiResponseData {
	p := new(OneOfCreateConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateConsistencyGroupApiResponseData is nil"))
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

func (p *OneOfCreateConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateConsistencyGroupApiResponseData"))
}

func (p *OneOfCreateConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateConsistencyGroupApiResponseData")
}

type OneOfProtectedResourcePromoteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
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

type OneOfGetRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RecoveryPoint         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetRecoveryPointApiResponseData() *OneOfGetRecoveryPointApiResponseData {
	p := new(OneOfGetRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRecoveryPointApiResponseData is nil"))
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

func (p *OneOfGetRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecoveryPointApiResponseData"))
}

func (p *OneOfGetRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecoveryPointApiResponseData")
}

type OneOfRecoveryPointRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfRecoveryPointRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRecoveryPointRestoreApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRestoreApiResponseData"))
}

func (p *OneOfRecoveryPointRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRestoreApiResponseData")
}

type OneOfComputeChangedRegionsResponseModelData struct {
	Discriminator *string                        `json:"-"`
	ObjectType_   *string                        `json:"-"`
	oneOfType0    *ComputeChangedRegionsResponse `json:"-"`
	oneOfType400  *import1.ErrorResponse         `json:"-"`
}

func NewOneOfComputeChangedRegionsResponseModelData() *OneOfComputeChangedRegionsResponseModelData {
	p := new(OneOfComputeChangedRegionsResponseModelData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfComputeChangedRegionsResponseModelData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfComputeChangedRegionsResponseModelData is nil"))
	}
	switch v.(type) {
	case ComputeChangedRegionsResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ComputeChangedRegionsResponse)
		}
		*p.oneOfType0 = v.(ComputeChangedRegionsResponse)
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

func (p *OneOfComputeChangedRegionsResponseModelData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfComputeChangedRegionsResponseModelData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ComputeChangedRegionsResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.config.ComputeChangedRegionsResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ComputeChangedRegionsResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfComputeChangedRegionsResponseModelData"))
}

func (p *OneOfComputeChangedRegionsResponseModelData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfComputeChangedRegionsResponseModelData")
}

type OneOfUpdateConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []import1.AppMessage   `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateConsistencyGroupApiResponseData() *OneOfUpdateConsistencyGroupApiResponseData {
	p := new(OneOfUpdateConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateConsistencyGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case []import1.AppMessage:
		p.oneOfType0 = v.([]import1.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.error.AppMessage>"
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

func (p *OneOfUpdateConsistencyGroupApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "dataprotection.v4.error.AppMessage" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.error.AppMessage>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.error.AppMessage>"
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateConsistencyGroupApiResponseData"))
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateConsistencyGroupApiResponseData")
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
