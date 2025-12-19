/*
 * Generated file models/dataprotection/v4/config/config_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Data Protection APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
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
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/common"
	import3 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/error"
	import5 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/prism/v4/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/vmm/v4/ahv/config"
	"time"
)

type AhvVmOverrideSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Categories to be associated with the VM on successful restore. If not specified, the VM is provisioned without any categories.
	*/
	Categories []import1.CategoryReference `json:"categories,omitempty"`
	/*
	  VM description.
	*/
	Description *string `json:"description,omitempty"`

	GuestToolsSpec *import1.VmRestoreGuestToolsSpecification `json:"guestToolsSpec,omitempty"`
	/*
	  Name of the VM to override with. If not specified, a name is chosen by the system and returned to the task entities when complete.
	*/
	Name *string `json:"name,omitempty"`

	NicSpec *import1.VmRestoreNicConfigSpecification `json:"nicSpec,omitempty"`

	OwnershipInfo *import1.OwnershipInfo `json:"ownershipInfo,omitempty"`
}

func (p *AhvVmOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AhvVmOverrideSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *AhvVmOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AhvVmOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAhvVmOverrideSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Categories != nil {
		p.Categories = known.Categories
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.GuestToolsSpec != nil {
		p.GuestToolsSpec = known.GuestToolsSpec
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NicSpec != nil {
		p.NicSpec = known.NicSpec
	}
	if known.OwnershipInfo != nil {
		p.OwnershipInfo = known.OwnershipInfo
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categories")
	delete(allFields, "description")
	delete(allFields, "guestToolsSpec")
	delete(allFields, "name")
	delete(allFields, "nicSpec")
	delete(allFields, "ownershipInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAhvVmOverrideSpec() *AhvVmOverrideSpec {
	p := new(AhvVmOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AhvVmOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *AmazonS3Bucket) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AmazonS3Bucket

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *AmazonS3Bucket) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AmazonS3Bucket
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAmazonS3Bucket()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BucketName != nil {
		p.BucketName = known.BucketName
	}
	if known.RegionName != nil {
		p.RegionName = known.RegionName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bucketName")
	delete(allFields, "regionName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAmazonS3Bucket() *AmazonS3Bucket {
	p := new(AmazonS3Bucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AmazonS3Bucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *AzureBlobStorageContainer) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AzureBlobStorageContainer

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *AzureBlobStorageContainer) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AzureBlobStorageContainer
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAzureBlobStorageContainer()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ContainerName != nil {
		p.ContainerName = known.ContainerName
	}
	if known.StorageAccountName != nil {
		p.StorageAccountName = known.StorageAccountName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "containerName")
	delete(allFields, "storageAccountName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAzureBlobStorageContainer() *AzureBlobStorageContainer {
	p := new(AzureBlobStorageContainer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.AzureBlobStorageContainer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Capability dictionary entry with capability name and boolean value indicating support.
*/
type Capability struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapabilityName *DataProtectionCapabilityName `json:"capabilityName"`
	/*
	  Boolean - true if the capability is supported, false otherwise.
	*/
	IsSupported *bool `json:"isSupported"`
}

func (p *Capability) MarshalJSON() ([]byte, error) {
	type CapabilityProxy Capability

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CapabilityProxy
		CapabilityName *DataProtectionCapabilityName `json:"capabilityName,omitempty"`
		IsSupported    *bool                         `json:"isSupported,omitempty"`
	}{
		CapabilityProxy: (*CapabilityProxy)(p),
		CapabilityName:  p.CapabilityName,
		IsSupported:     p.IsSupported,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Capability) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Capability
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCapability()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapabilityName != nil {
		p.CapabilityName = known.CapabilityName
	}
	if known.IsSupported != nil {
		p.IsSupported = known.IsSupported
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilityName")
	delete(allFields, "isSupported")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCapability() *Capability {
	p := new(Capability)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Capability"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId}/$actions/discover-cluster Post operation
*/
type ClusterInfoApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfClusterInfoApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ClusterInfoApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterInfoApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ClusterInfoApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterInfoApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterInfoApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterInfoApiResponse() *ClusterInfoApiResponse {
	p := new(ClusterInfoApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ClusterInfoApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterInfoApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ClusterInfoApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfClusterInfoApiResponseData()
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
In many scenarios, it is necessary to capture the state of an application as a combined snapshot of the internal states of a related set of entities at a specific point in time. A consistency group is a collection of these entities, where their combined snapshot reflects the state of an application at that particular point in time.
*/
type ConsistencyGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the cluster to which the entities in the consistency group are associated.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	Members []ConsistencyGroupMember `json:"members"`
	/*
	  Name of the consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this consistency group. This is a read-only field that is inserted into the consistency group entity at the time of the consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  The external identifier of the protection policy that protects the consistency group.
	*/
	ProtectionPolicyExtId *string `json:"protectionPolicyExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroup) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProxy ConsistencyGroup

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConsistencyGroupProxy
		Members []ConsistencyGroupMember `json:"members,omitempty"`
		Name    *string                  `json:"name,omitempty"`
	}{
		ConsistencyGroupProxy: (*ConsistencyGroupProxy)(p),
		Members:               p.Members,
		Name:                  p.Name,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ConsistencyGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConsistencyGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConsistencyGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Members != nil {
		p.Members = known.Members
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.ProtectionPolicyExtId != nil {
		p.ProtectionPolicyExtId = known.ProtectionPolicyExtId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "members")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "protectionPolicyExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConsistencyGroup() *ConsistencyGroup {
	p := new(ConsistencyGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
External identifier of the members (VM or volume group) of the consistency group object.
*/
type ConsistencyGroupMember struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the members (VM or volume group) of the consistency group object.
	*/
	EntityExtId *string `json:"entityExtId"`

	EntityType *ConsistencyGroupMemberType `json:"entityType"`
}

func (p *ConsistencyGroupMember) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupMemberProxy ConsistencyGroupMember

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConsistencyGroupMemberProxy
		EntityExtId *string                     `json:"entityExtId,omitempty"`
		EntityType  *ConsistencyGroupMemberType `json:"entityType,omitempty"`
	}{
		ConsistencyGroupMemberProxy: (*ConsistencyGroupMemberProxy)(p),
		EntityExtId:                 p.EntityExtId,
		EntityType:                  p.EntityType,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ConsistencyGroupMember) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConsistencyGroupMember
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConsistencyGroupMember()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityExtId != nil {
		p.EntityExtId = known.EntityExtId
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityExtId")
	delete(allFields, "entityType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConsistencyGroupMember() *ConsistencyGroupMember {
	p := new(ConsistencyGroupMember)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMember"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Consistency group member type.
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
Specification for the migration action on the consistency group.
*/
type ConsistencyGroupMigrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Reference to the cluster in the target availability zone where the entities need to be migrated.
	*/
	TargetClusterExtId *string `json:"targetClusterExtId,omitempty"`
	/*
	  Reference to the target availability zone where the entities need to be migrated.
	*/
	TargetPcExtId *string `json:"targetPcExtId,omitempty"`
}

func (p *ConsistencyGroupMigrationSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConsistencyGroupMigrationSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ConsistencyGroupMigrationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConsistencyGroupMigrationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConsistencyGroupMigrationSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.TargetClusterExtId != nil {
		p.TargetClusterExtId = known.TargetClusterExtId
	}
	if known.TargetPcExtId != nil {
		p.TargetPcExtId = known.TargetPcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "targetClusterExtId")
	delete(allFields, "targetPcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConsistencyGroupMigrationSpec() *ConsistencyGroupMigrationSpec {
	p := new(ConsistencyGroupMigrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupMigrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConsistencyGroupProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the cluster to which the entities in the consistency group are associated.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	Members []ConsistencyGroupMember `json:"members"`
	/*
	  Name of the consistency group.
	*/
	Name *string `json:"name"`
	/*
	  The external identifier of the user who created this consistency group. This is a read-only field that is inserted into the consistency group entity at the time of the consistency group creation.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  The external identifier of the protection policy that protects the consistency group.
	*/
	ProtectionPolicyExtId *string `json:"protectionPolicyExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConsistencyGroupProjection) MarshalJSON() ([]byte, error) {
	type ConsistencyGroupProjectionProxy ConsistencyGroupProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConsistencyGroupProjectionProxy
		Members []ConsistencyGroupMember `json:"members,omitempty"`
		Name    *string                  `json:"name,omitempty"`
	}{
		ConsistencyGroupProjectionProxy: (*ConsistencyGroupProjectionProxy)(p),
		Members:                         p.Members,
		Name:                            p.Name,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ConsistencyGroupProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConsistencyGroupProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConsistencyGroupProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Members != nil {
		p.Members = known.Members
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.ProtectionPolicyExtId != nil {
		p.ProtectionPolicyExtId = known.ProtectionPolicyExtId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "members")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "protectionPolicyExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConsistencyGroupProjection() *ConsistencyGroupProjection {
	p := new(ConsistencyGroupProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ConsistencyGroupProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/consistency-groups Post operation
*/
type CreateConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateConsistencyGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateConsistencyGroupApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *CreateConsistencyGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateConsistencyGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateConsistencyGroupApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateConsistencyGroupApiResponse() *CreateConsistencyGroupApiResponse {
	p := new(CreateConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreateConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points Post operation
*/
type CreateRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateRecoveryPointApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *CreateRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateRecoveryPointApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateRecoveryPointApiResponse() *CreateRecoveryPointApiResponse {
	p := new(CreateRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreateRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Defines a time window using start and end timestamps, which is used to filter recovery points for deletion. The timestamps are in UTC and in ISO 8601 format.
*/
type CreationTimeRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Start of the time range. Recovery points created after this time are considered.
	*/
	FirstCreationTime *time.Time `json:"firstCreationTime,omitempty"`
	/*
	  End of the time range. Recovery points created before this time are considered.
	*/
	LastCreationTime *time.Time `json:"lastCreationTime,omitempty"`
}

func (p *CreationTimeRange) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreationTimeRange

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *CreationTimeRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreationTimeRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreationTimeRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FirstCreationTime != nil {
		p.FirstCreationTime = known.FirstCreationTime
	}
	if known.LastCreationTime != nil {
		p.LastCreationTime = known.LastCreationTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "firstCreationTime")
	delete(allFields, "lastCreationTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreationTimeRange() *CreationTimeRange {
	p := new(CreationTimeRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.CreationTimeRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Name of the capability e.g. "SUPPORTS_NEAR_SYNC".
*/
type DataProtectionCapabilityName int

const (
	DATAPROTECTIONCAPABILITYNAME_UNKNOWN            DataProtectionCapabilityName = 0
	DATAPROTECTIONCAPABILITYNAME_REDACTED           DataProtectionCapabilityName = 1
	DATAPROTECTIONCAPABILITYNAME_SUPPORTS_NEAR_SYNC DataProtectionCapabilityName = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DataProtectionCapabilityName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUPPORTS_NEAR_SYNC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DataProtectionCapabilityName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUPPORTS_NEAR_SYNC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DataProtectionCapabilityName) index(name string) DataProtectionCapabilityName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUPPORTS_NEAR_SYNC",
	}
	for idx := range names {
		if names[idx] == name {
			return DataProtectionCapabilityName(idx)
		}
	}
	return DATAPROTECTIONCAPABILITYNAME_UNKNOWN
}

func (e *DataProtectionCapabilityName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DataProtectionCapabilityName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DataProtectionCapabilityName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DataProtectionCapabilityName) Ref() *DataProtectionCapabilityName {
	return &e
}

/*
Describes data protection cluster capabilities.
*/
type DataProtectionClusterCapability struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A dictionary with a key as a capability name and a value as a boolean, true if supported by the cluster, false if not.  For example, the dictionary will look like { "SUPPORTS_NEAR_SYNC": true }.
	*/
	Capabilities []Capability `json:"capabilities,omitempty"`
	/*
	  Cluster UUID whose capabilities are retrieved.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DataProtectionClusterCapability) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DataProtectionClusterCapability

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DataProtectionClusterCapability) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DataProtectionClusterCapability
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDataProtectionClusterCapability()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDataProtectionClusterCapability() *DataProtectionClusterCapability {
	p := new(DataProtectionClusterCapability)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionClusterCapability"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DataProtectionClusterCapabilityProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A dictionary with a key as a capability name and a value as a boolean, true if supported by the cluster, false if not.  For example, the dictionary will look like { "SUPPORTS_NEAR_SYNC": true }.
	*/
	Capabilities []Capability `json:"capabilities,omitempty"`
	/*
	  Cluster UUID whose capabilities are retrieved.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DataProtectionClusterCapabilityProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DataProtectionClusterCapabilityProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DataProtectionClusterCapabilityProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DataProtectionClusterCapabilityProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDataProtectionClusterCapabilityProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDataProtectionClusterCapabilityProjection() *DataProtectionClusterCapabilityProjection {
	p := new(DataProtectionClusterCapabilityProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionClusterCapabilityProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the data protection site in Prism Central.
*/
type DataProtectionSiteReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the domain manager.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *DataProtectionSiteReference) MarshalJSON() ([]byte, error) {
	type DataProtectionSiteReferenceProxy DataProtectionSiteReference

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DataProtectionSiteReferenceProxy
		DomainManagerExtId *string `json:"domainManagerExtId,omitempty"`
	}{
		DataProtectionSiteReferenceProxy: (*DataProtectionSiteReferenceProxy)(p),
		DomainManagerExtId:               p.DomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DataProtectionSiteReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DataProtectionSiteReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDataProtectionSiteReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDataProtectionSiteReference() *DataProtectionSiteReference {
	p := new(DataProtectionSiteReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DataProtectionSiteReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/consistency-groups/{extId} Delete operation
*/
type DeleteConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteConsistencyGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteConsistencyGroupApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DeleteConsistencyGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteConsistencyGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteConsistencyGroupApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteConsistencyGroupApiResponse() *DeleteConsistencyGroupApiResponse {
	p := new(DeleteConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DeleteConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-plan-jobs/{extId} Delete operation
*/
type DeleteRecoveryPlanJobApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRecoveryPlanJobApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteRecoveryPlanJobApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteRecoveryPlanJobApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DeleteRecoveryPlanJobApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteRecoveryPlanJobApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteRecoveryPlanJobApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteRecoveryPlanJobApiResponse() *DeleteRecoveryPlanJobApiResponse {
	p := new(DeleteRecoveryPlanJobApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DeleteRecoveryPlanJobApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRecoveryPlanJobApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRecoveryPlanJobApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRecoveryPlanJobApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId} Delete operation
*/
type DeleteRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteRecoveryPointApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DeleteRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteRecoveryPointApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteRecoveryPointApiResponse() *DeleteRecoveryPointApiResponse {
	p := new(DeleteRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DeleteRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

/*
disasterRecoveryLocationDesc
*/
type DisasterRecoveryLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the clusters associated with the recovery point at source Prism Central.
	*/
	ClusterExtIds []string `json:"clusterExtIds"`
	/*
	  External identifier of the Prism Central (also known as the Domain Manager) where the recovery point was first created.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *DisasterRecoveryLocation) MarshalJSON() ([]byte, error) {
	type DisasterRecoveryLocationProxy DisasterRecoveryLocation

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DisasterRecoveryLocationProxy
		ClusterExtIds      []string `json:"clusterExtIds,omitempty"`
		DomainManagerExtId *string  `json:"domainManagerExtId,omitempty"`
	}{
		DisasterRecoveryLocationProxy: (*DisasterRecoveryLocationProxy)(p),
		ClusterExtIds:                 p.ClusterExtIds,
		DomainManagerExtId:            p.DomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DisasterRecoveryLocation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DisasterRecoveryLocation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDisasterRecoveryLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDisasterRecoveryLocation() *DisasterRecoveryLocation {
	p := new(DisasterRecoveryLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.DisasterRecoveryLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This contains the external identifier of the source entity and recovered entity information, such as entity name and the recovered entity external identifier.
*/
type EntityRecoveryResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *import6.EntityType `json:"entityType,omitempty"`
	/*
	  Name of the recovered entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Information about the recovered entity.
	*/
	RecoveredEntityExtId *string `json:"recoveredEntityExtId,omitempty"`
	/*
	  Source entity reference.
	*/
	SourceEntityExtId *string `json:"sourceEntityExtId,omitempty"`
}

func (p *EntityRecoveryResult) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityRecoveryResult

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *EntityRecoveryResult) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityRecoveryResult
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityRecoveryResult()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.RecoveredEntityExtId != nil {
		p.RecoveredEntityExtId = known.RecoveredEntityExtId
	}
	if known.SourceEntityExtId != nil {
		p.SourceEntityExtId = known.SourceEntityExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityType")
	delete(allFields, "name")
	delete(allFields, "recoveredEntityExtId")
	delete(allFields, "sourceEntityExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityRecoveryResult() *EntityRecoveryResult {
	p := new(EntityRecoveryResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.EntityRecoveryResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
External reference of an entity.
*/
type EntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the entity.
	*/
	ExtId *string `json:"extId"`
}

func (p *EntityReference) MarshalJSON() ([]byte, error) {
	type EntityReferenceProxy EntityReference

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EntityReferenceProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		EntityReferenceProxy: (*EntityReferenceProxy)(p),
		ExtId:                p.ExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *EntityReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of validation errors or warnings.
*/
type ErrorGroup int

const (
	ERRORGROUP_UNKNOWN         ErrorGroup = 0
	ERRORGROUP_REDACTED        ErrorGroup = 1
	ERRORGROUP_ENTITY          ErrorGroup = 2
	ERRORGROUP_NETWORK         ErrorGroup = 3
	ERRORGROUP_RECOVERY_POINT  ErrorGroup = 4
	ERRORGROUP_CATEGORY        ErrorGroup = 5
	ERRORGROUP_CONNECTIVITY    ErrorGroup = 6
	ERRORGROUP_INFRASTRUCTURE  ErrorGroup = 7
	ERRORGROUP_CLUSTER_LICENSE ErrorGroup = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ErrorGroup) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"NETWORK",
		"RECOVERY_POINT",
		"CATEGORY",
		"CONNECTIVITY",
		"INFRASTRUCTURE",
		"CLUSTER_LICENSE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ErrorGroup) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"NETWORK",
		"RECOVERY_POINT",
		"CATEGORY",
		"CONNECTIVITY",
		"INFRASTRUCTURE",
		"CLUSTER_LICENSE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ErrorGroup) index(name string) ErrorGroup {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENTITY",
		"NETWORK",
		"RECOVERY_POINT",
		"CATEGORY",
		"CONNECTIVITY",
		"INFRASTRUCTURE",
		"CLUSTER_LICENSE",
	}
	for idx := range names {
		if names[idx] == name {
			return ErrorGroup(idx)
		}
	}
	return ERRORGROUP_UNKNOWN
}

func (e *ErrorGroup) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ErrorGroup:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ErrorGroup) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ErrorGroup) Ref() *ErrorGroup {
	return &e
}

/*
Provides the code, status, severity and summary of the error raised by an API call.
*/
type ErrorMessage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Error code for API request failure.
	*/
	ErrorCode *string `json:"errorCode,omitempty"`

	ErrorGroup *ErrorGroup `json:"errorGroup,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Summary of the error.
	*/
	Message *string `json:"message,omitempty"`

	Severity *Severity `json:"severity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ErrorMessage) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ErrorMessage

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ErrorMessage) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ErrorMessage
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewErrorMessage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ErrorCode != nil {
		p.ErrorCode = known.ErrorCode
	}
	if known.ErrorGroup != nil {
		p.ErrorGroup = known.ErrorGroup
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "errorCode")
	delete(allFields, "errorGroup")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "severity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewErrorMessage() *ErrorMessage {
	p := new(ErrorMessage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ErrorMessage"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Protected resource/recovery point restore that overrides the ESXi VM configuration. The specified properties are overridden for the restored VM.
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

func (p *EsxiVmOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EsxiVmOverrideSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *EsxiVmOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EsxiVmOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEsxiVmOverrideSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEsxiVmOverrideSpec() *EsxiVmOverrideSpec {
	p := new(EsxiVmOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.EsxiVmOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity reference model for execution step affected entities.
*/
type ExecutionStepEntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *ExecutionStepEntityType `json:"entityType,omitempty"`
	/*
	  External identifier of the entity.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Name of the entity.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *ExecutionStepEntityReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExecutionStepEntityReference

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ExecutionStepEntityReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExecutionStepEntityReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExecutionStepEntityReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExecutionStepEntityReference() *ExecutionStepEntityReference {
	p := new(ExecutionStepEntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ExecutionStepEntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of entity being recovered.
*/
type ExecutionStepEntityType int

const (
	EXECUTIONSTEPENTITYTYPE_UNKNOWN           ExecutionStepEntityType = 0
	EXECUTIONSTEPENTITYTYPE_REDACTED          ExecutionStepEntityType = 1
	EXECUTIONSTEPENTITYTYPE_VM                ExecutionStepEntityType = 2
	EXECUTIONSTEPENTITYTYPE_SUBNET            ExecutionStepEntityType = 3
	EXECUTIONSTEPENTITYTYPE_AVAILABILITY_ZONE ExecutionStepEntityType = 4
	EXECUTIONSTEPENTITYTYPE_RECOVERY_PLAN     ExecutionStepEntityType = 5
	EXECUTIONSTEPENTITYTYPE_VIRTUAL_NETWORK   ExecutionStepEntityType = 6
	EXECUTIONSTEPENTITYTYPE_VOLUME_GROUP      ExecutionStepEntityType = 7
	EXECUTIONSTEPENTITYTYPE_CONSISTENCY_GROUP ExecutionStepEntityType = 8
	EXECUTIONSTEPENTITYTYPE_CATEGORY          ExecutionStepEntityType = 9
	EXECUTIONSTEPENTITYTYPE_VPN_GATEWAY       ExecutionStepEntityType = 10
	EXECUTIONSTEPENTITYTYPE_SUBNET_EXTENSION  ExecutionStepEntityType = 11
	EXECUTIONSTEPENTITYTYPE_VPC               ExecutionStepEntityType = 12
	EXECUTIONSTEPENTITYTYPE_RECOVERY_PLAN_JOB ExecutionStepEntityType = 13
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ExecutionStepEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"SUBNET",
		"AVAILABILITY_ZONE",
		"RECOVERY_PLAN",
		"VIRTUAL_NETWORK",
		"VOLUME_GROUP",
		"CONSISTENCY_GROUP",
		"CATEGORY",
		"VPN_GATEWAY",
		"SUBNET_EXTENSION",
		"VPC",
		"RECOVERY_PLAN_JOB",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ExecutionStepEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"SUBNET",
		"AVAILABILITY_ZONE",
		"RECOVERY_PLAN",
		"VIRTUAL_NETWORK",
		"VOLUME_GROUP",
		"CONSISTENCY_GROUP",
		"CATEGORY",
		"VPN_GATEWAY",
		"SUBNET_EXTENSION",
		"VPC",
		"RECOVERY_PLAN_JOB",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ExecutionStepEntityType) index(name string) ExecutionStepEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"SUBNET",
		"AVAILABILITY_ZONE",
		"RECOVERY_PLAN",
		"VIRTUAL_NETWORK",
		"VOLUME_GROUP",
		"CONSISTENCY_GROUP",
		"CATEGORY",
		"VPN_GATEWAY",
		"SUBNET_EXTENSION",
		"VPC",
		"RECOVERY_PLAN_JOB",
	}
	for idx := range names {
		if names[idx] == name {
			return ExecutionStepEntityType(idx)
		}
	}
	return EXECUTIONSTEPENTITYTYPE_UNKNOWN
}

func (e *ExecutionStepEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ExecutionStepEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ExecutionStepEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ExecutionStepEntityType) Ref() *ExecutionStepEntityType {
	return &e
}

/*
Resulted output for an execution of a step.
*/
type ExecutionStepResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ResultItemDiscriminator_ *string `json:"$resultItemDiscriminator,omitempty"`
	/*
	  Resulted output for an execution of a step.
	*/
	Result *OneOfExecutionStepResultResult `json:"result"`
}

func (p *ExecutionStepResult) MarshalJSON() ([]byte, error) {
	type ExecutionStepResultProxy ExecutionStepResult

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ExecutionStepResultProxy
		Result *OneOfExecutionStepResultResult `json:"result,omitempty"`
	}{
		ExecutionStepResultProxy: (*ExecutionStepResultProxy)(p),
		Result:                   p.Result,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ExecutionStepResult) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExecutionStepResult
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExecutionStepResult()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ResultItemDiscriminator_ != nil {
		p.ResultItemDiscriminator_ = known.ResultItemDiscriminator_
	}
	if known.Result != nil {
		p.Result = known.Result
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$resultItemDiscriminator")
	delete(allFields, "result")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExecutionStepResult() *ExecutionStepResult {
	p := new(ExecutionStepResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ExecutionStepResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ExecutionStepResult) GetResult() interface{} {
	if nil == p.Result {
		return nil
	}
	return p.Result.GetValue()
}

func (p *ExecutionStepResult) SetResult(v interface{}) error {
	if nil == p.Result {
		p.Result = NewOneOfExecutionStepResultResult()
	}
	e := p.Result.SetValue(v)
	if nil == e {
		if nil == p.ResultItemDiscriminator_ {
			p.ResultItemDiscriminator_ = new(string)
		}
		*p.ResultItemDiscriminator_ = *p.Result.Discriminator
	}
	return e
}

/*
Specification to set the expiration time of the recovery point.
*/
type ExpirationTimeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current recovery point expires . If not specified, the recovery point never expires.
	*/
	ExpirationTime *time.Time `json:"expirationTime"`
}

func (p *ExpirationTimeSpec) MarshalJSON() ([]byte, error) {
	type ExpirationTimeSpecProxy ExpirationTimeSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ExpirationTimeSpecProxy
		ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	}{
		ExpirationTimeSpecProxy: (*ExpirationTimeSpecProxy)(p),
		ExpirationTime:          p.ExpirationTime,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ExpirationTimeSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExpirationTimeSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExpirationTimeSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "expirationTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExpirationTimeSpec() *ExpirationTimeSpec {
	p := new(ExpirationTimeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ExpirationTimeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Failover direction from the source disaster recovery location to the target disaster recovery location.
*/
type FailoverDirection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SourceCluster *EntityReference `json:"sourceCluster,omitempty"`
	/*
	  External identifier of the source domain manager.
	*/
	SourceDomainManagerExtId *string `json:"sourceDomainManagerExtId"`

	TargetCluster *EntityReference `json:"targetCluster,omitempty"`
	/*
	  External identifier of the target domain manager.
	*/
	TargetDomainManagerExtId *string `json:"targetDomainManagerExtId"`
}

func (p *FailoverDirection) MarshalJSON() ([]byte, error) {
	type FailoverDirectionProxy FailoverDirection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*FailoverDirectionProxy
		SourceDomainManagerExtId *string `json:"sourceDomainManagerExtId,omitempty"`
		TargetDomainManagerExtId *string `json:"targetDomainManagerExtId,omitempty"`
	}{
		FailoverDirectionProxy:   (*FailoverDirectionProxy)(p),
		SourceDomainManagerExtId: p.SourceDomainManagerExtId,
		TargetDomainManagerExtId: p.TargetDomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *FailoverDirection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FailoverDirection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFailoverDirection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.SourceCluster != nil {
		p.SourceCluster = known.SourceCluster
	}
	if known.SourceDomainManagerExtId != nil {
		p.SourceDomainManagerExtId = known.SourceDomainManagerExtId
	}
	if known.TargetCluster != nil {
		p.TargetCluster = known.TargetCluster
	}
	if known.TargetDomainManagerExtId != nil {
		p.TargetDomainManagerExtId = known.TargetDomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "sourceCluster")
	delete(allFields, "sourceDomainManagerExtId")
	delete(allFields, "targetCluster")
	delete(allFields, "targetDomainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFailoverDirection() *FailoverDirection {
	p := new(FailoverDirection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.FailoverDirection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for force deleting recovery points across specified clusters within a given time range.
*/
type ForceDeleteAllRecoveryPointsSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterExtIds []string `json:"clusterExtIds,omitempty"`

	TimeRange *CreationTimeRange `json:"timeRange,omitempty"`
}

func (p *ForceDeleteAllRecoveryPointsSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ForceDeleteAllRecoveryPointsSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ForceDeleteAllRecoveryPointsSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ForceDeleteAllRecoveryPointsSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewForceDeleteAllRecoveryPointsSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.TimeRange != nil {
		p.TimeRange = known.TimeRange
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "timeRange")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewForceDeleteAllRecoveryPointsSpec() *ForceDeleteAllRecoveryPointsSpec {
	p := new(ForceDeleteAllRecoveryPointsSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ForceDeleteAllRecoveryPointsSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/consistency-groups/{extId} Get operation
*/
type GetConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetConsistencyGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetConsistencyGroupApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetConsistencyGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetConsistencyGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetConsistencyGroupApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetConsistencyGroupApiResponse() *GetConsistencyGroupApiResponse {
	p := new(GetConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetConsistencyGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetConsistencyGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetConsistencyGroupApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/protected-resources/{extId} Get operation
*/
type GetProtectedResourceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetProtectedResourceApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetProtectedResourceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetProtectedResourceApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetProtectedResourceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetProtectedResourceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetProtectedResourceApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetProtectedResourceApiResponse() *GetProtectedResourceApiResponse {
	p := new(GetProtectedResourceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetProtectedResourceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-plan-jobs/{extId} Get operation
*/
type GetRecoveryPlanJobApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecoveryPlanJobApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetRecoveryPlanJobApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetRecoveryPlanJobApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetRecoveryPlanJobApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetRecoveryPlanJobApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetRecoveryPlanJobApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetRecoveryPlanJobApiResponse() *GetRecoveryPlanJobApiResponse {
	p := new(GetRecoveryPlanJobApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetRecoveryPlanJobApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRecoveryPlanJobApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRecoveryPlanJobApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRecoveryPlanJobApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId} Get operation
*/
type GetRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetRecoveryPointApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetRecoveryPointApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetRecoveryPointApiResponse() *GetRecoveryPointApiResponse {
	p := new(GetRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{recoveryPointExtId}/vm-recovery-points/{extId} Get operation
*/
type GetVmRecoveryPointApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmRecoveryPointApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmRecoveryPointApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmRecoveryPointApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetVmRecoveryPointApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmRecoveryPointApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetVmRecoveryPointApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetVmRecoveryPointApiResponse() *GetVmRecoveryPointApiResponse {
	p := new(GetVmRecoveryPointApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.GetVmRecoveryPointApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmRecoveryPointApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmRecoveryPointApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmRecoveryPointApiResponseData()
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
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness host site.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *HostReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostReference

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *HostReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHostReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostType != nil {
		p.HostType = known.HostType
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "hostType")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Type of Hypervisor present on the cluster.
*/
type HypervisorType int

const (
	HYPERVISORTYPE_UNKNOWN  HypervisorType = 0
	HYPERVISORTYPE_REDACTED HypervisorType = 1
	HYPERVISORTYPE_AHV      HypervisorType = 2
	HYPERVISORTYPE_ESXI     HypervisorType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HypervisorType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESXI",
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
		"ESXI",
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
		"ESXI",
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
REST response for all response codes in API path /dataprotection/v4.2/config/consistency-groups Get operation
*/
type ListConsistencyGroupsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListConsistencyGroupsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListConsistencyGroupsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListConsistencyGroupsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListConsistencyGroupsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListConsistencyGroupsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListConsistencyGroupsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListConsistencyGroupsApiResponse() *ListConsistencyGroupsApiResponse {
	p := new(ListConsistencyGroupsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListConsistencyGroupsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListConsistencyGroupsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListConsistencyGroupsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListConsistencyGroupsApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/cluster-capabilities Get operation
*/
type ListDPClusterCapabilitiesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDPClusterCapabilitiesApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDPClusterCapabilitiesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDPClusterCapabilitiesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListDPClusterCapabilitiesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDPClusterCapabilitiesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListDPClusterCapabilitiesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListDPClusterCapabilitiesApiResponse() *ListDPClusterCapabilitiesApiResponse {
	p := new(ListDPClusterCapabilitiesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListDPClusterCapabilitiesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDPClusterCapabilitiesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDPClusterCapabilitiesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDPClusterCapabilitiesApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-plan-jobs/{recoveryPlanJobExtId}/execution-steps Get operation
*/
type ListRecoveryPlanJobExecutionStepsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecoveryPlanJobExecutionStepsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRecoveryPlanJobExecutionStepsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRecoveryPlanJobExecutionStepsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListRecoveryPlanJobExecutionStepsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRecoveryPlanJobExecutionStepsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRecoveryPlanJobExecutionStepsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListRecoveryPlanJobExecutionStepsApiResponse() *ListRecoveryPlanJobExecutionStepsApiResponse {
	p := new(ListRecoveryPlanJobExecutionStepsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListRecoveryPlanJobExecutionStepsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecoveryPlanJobExecutionStepsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecoveryPlanJobExecutionStepsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecoveryPlanJobExecutionStepsApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-plan-jobs/{recoveryPlanJobExtId}/validation-errors Get operation
*/
type ListRecoveryPlanJobValidationErrorsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecoveryPlanJobValidationErrorsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRecoveryPlanJobValidationErrorsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRecoveryPlanJobValidationErrorsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListRecoveryPlanJobValidationErrorsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRecoveryPlanJobValidationErrorsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRecoveryPlanJobValidationErrorsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListRecoveryPlanJobValidationErrorsApiResponse() *ListRecoveryPlanJobValidationErrorsApiResponse {
	p := new(ListRecoveryPlanJobValidationErrorsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListRecoveryPlanJobValidationErrorsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecoveryPlanJobValidationErrorsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecoveryPlanJobValidationErrorsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecoveryPlanJobValidationErrorsApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-plan-jobs Get operation
*/
type ListRecoveryPlanJobsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecoveryPlanJobsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRecoveryPlanJobsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRecoveryPlanJobsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListRecoveryPlanJobsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRecoveryPlanJobsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRecoveryPlanJobsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListRecoveryPlanJobsApiResponse() *ListRecoveryPlanJobsApiResponse {
	p := new(ListRecoveryPlanJobsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListRecoveryPlanJobsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRecoveryPlanJobsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRecoveryPlanJobsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRecoveryPlanJobsApiResponseData()
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points Get operation
*/
type ListRecoveryPointsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRecoveryPointsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRecoveryPointsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRecoveryPointsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListRecoveryPointsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRecoveryPointsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRecoveryPointsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListRecoveryPointsApiResponse() *ListRecoveryPointsApiResponse {
	p := new(ListRecoveryPointsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ListRecoveryPointsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *LocationReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LocationReference

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *LocationReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocationReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocationReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LocationExtId != nil {
		p.LocationExtId = known.LocationExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "locationExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocationReference() *LocationReference {
	p := new(LocationReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.LocationReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

func (p *NutanixObjectsBucket) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NutanixObjectsBucket

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *NutanixObjectsBucket) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NutanixObjectsBucket
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNutanixObjectsBucket()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndPoint != nil {
		p.EndPoint = known.EndPoint
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endPoint")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNutanixObjectsBucket() *NutanixObjectsBucket {
	p := new(NutanixObjectsBucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.NutanixObjectsBucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Type of operation being performed.
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN                     OperationType = 0
	OPERATIONTYPE_REDACTED                    OperationType = 1
	OPERATIONTYPE_NETWORK_CREATION            OperationType = 2
	OPERATIONTYPE_SUBNET_EXTENSION_CREATION   OperationType = 3
	OPERATIONTYPE_NETWORK_GATEWAY_CREATION    OperationType = 4
	OPERATIONTYPE_VPC_UPDATION                OperationType = 5
	OPERATIONTYPE_SYNC_RECOVERY_POINTS        OperationType = 6
	OPERATIONTYPE_FLOATING_IP_CREATION        OperationType = 7
	OPERATIONTYPE_RESOURCE_RESERVATION        OperationType = 8
	OPERATIONTYPE_NOTIFY_WITNESS              OperationType = 9
	OPERATIONTYPE_VPC_CREATION                OperationType = 10
	OPERATIONTYPE_CONFIGURE_MANUAL_PROTECTION OperationType = 11
	OPERATIONTYPE_DELAY                       OperationType = 12
	OPERATIONTYPE_ENTITY_MIGRATION            OperationType = 13
	OPERATIONTYPE_ENTITY_RESTORATION          OperationType = 14
	OPERATIONTYPE_ENTITY_LIVE_MIGRATION       OperationType = 15
	OPERATIONTYPE_VM_POWER_ON                 OperationType = 16
	OPERATIONTYPE_ASSIGN_FLOATING_IP          OperationType = 17
	OPERATIONTYPE_STATIC_IP_VALIDATION        OperationType = 18
	OPERATIONTYPE_SCRIPT_EXECUTION            OperationType = 19
	OPERATIONTYPE_VM_VOLUME_IQN_ATTACHMENT    OperationType = 20
	OPERATIONTYPE_VM_VOLUME_ATTACHMENT        OperationType = 21
	OPERATIONTYPE_ENTITY_SNAPSHOT_SYNC        OperationType = 22
	OPERATIONTYPE_AUDIT                       OperationType = 23
	OPERATIONTYPE_NETWORK_DELETION            OperationType = 24
	OPERATIONTYPE_CONSISTENCY_GROUP_CREATE    OperationType = 25
	OPERATIONTYPE_SUBNET_EXTENSION_DELETION   OperationType = 26
	OPERATIONTYPE_NETWORK_GATEWAY_DELETION    OperationType = 27
	OPERATIONTYPE_STAGE_RECOVERY              OperationType = 28
	OPERATIONTYPE_ENTITY_RECOVERY             OperationType = 29
	OPERATIONTYPE_IP_CUSTOMIZATION            OperationType = 30
	OPERATIONTYPE_ENTITIES_CLEANUP            OperationType = 31
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NETWORK_CREATION",
		"SUBNET_EXTENSION_CREATION",
		"NETWORK_GATEWAY_CREATION",
		"VPC_UPDATION",
		"SYNC_RECOVERY_POINTS",
		"FLOATING_IP_CREATION",
		"RESOURCE_RESERVATION",
		"NOTIFY_WITNESS",
		"VPC_CREATION",
		"CONFIGURE_MANUAL_PROTECTION",
		"DELAY",
		"ENTITY_MIGRATION",
		"ENTITY_RESTORATION",
		"ENTITY_LIVE_MIGRATION",
		"VM_POWER_ON",
		"ASSIGN_FLOATING_IP",
		"STATIC_IP_VALIDATION",
		"SCRIPT_EXECUTION",
		"VM_VOLUME_IQN_ATTACHMENT",
		"VM_VOLUME_ATTACHMENT",
		"ENTITY_SNAPSHOT_SYNC",
		"AUDIT",
		"NETWORK_DELETION",
		"CONSISTENCY_GROUP_CREATE",
		"SUBNET_EXTENSION_DELETION",
		"NETWORK_GATEWAY_DELETION",
		"STAGE_RECOVERY",
		"ENTITY_RECOVERY",
		"IP_CUSTOMIZATION",
		"ENTITIES_CLEANUP",
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
		"NETWORK_CREATION",
		"SUBNET_EXTENSION_CREATION",
		"NETWORK_GATEWAY_CREATION",
		"VPC_UPDATION",
		"SYNC_RECOVERY_POINTS",
		"FLOATING_IP_CREATION",
		"RESOURCE_RESERVATION",
		"NOTIFY_WITNESS",
		"VPC_CREATION",
		"CONFIGURE_MANUAL_PROTECTION",
		"DELAY",
		"ENTITY_MIGRATION",
		"ENTITY_RESTORATION",
		"ENTITY_LIVE_MIGRATION",
		"VM_POWER_ON",
		"ASSIGN_FLOATING_IP",
		"STATIC_IP_VALIDATION",
		"SCRIPT_EXECUTION",
		"VM_VOLUME_IQN_ATTACHMENT",
		"VM_VOLUME_ATTACHMENT",
		"ENTITY_SNAPSHOT_SYNC",
		"AUDIT",
		"NETWORK_DELETION",
		"CONSISTENCY_GROUP_CREATE",
		"SUBNET_EXTENSION_DELETION",
		"NETWORK_GATEWAY_DELETION",
		"STAGE_RECOVERY",
		"ENTITY_RECOVERY",
		"IP_CUSTOMIZATION",
		"ENTITIES_CLEANUP",
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
		"NETWORK_CREATION",
		"SUBNET_EXTENSION_CREATION",
		"NETWORK_GATEWAY_CREATION",
		"VPC_UPDATION",
		"SYNC_RECOVERY_POINTS",
		"FLOATING_IP_CREATION",
		"RESOURCE_RESERVATION",
		"NOTIFY_WITNESS",
		"VPC_CREATION",
		"CONFIGURE_MANUAL_PROTECTION",
		"DELAY",
		"ENTITY_MIGRATION",
		"ENTITY_RESTORATION",
		"ENTITY_LIVE_MIGRATION",
		"VM_POWER_ON",
		"ASSIGN_FLOATING_IP",
		"STATIC_IP_VALIDATION",
		"SCRIPT_EXECUTION",
		"VM_VOLUME_IQN_ATTACHMENT",
		"VM_VOLUME_ATTACHMENT",
		"ENTITY_SNAPSHOT_SYNC",
		"AUDIT",
		"NETWORK_DELETION",
		"CONSISTENCY_GROUP_CREATE",
		"SUBNET_EXTENSION_DELETION",
		"NETWORK_GATEWAY_DELETION",
		"STAGE_RECOVERY",
		"ENTITY_RECOVERY",
		"IP_CUSTOMIZATION",
		"ENTITIES_CLEANUP",
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
Summary of the phase in the recovery plan job. This is a generic model to represent the summary in the recovery plan job execution.
*/
type PhaseSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The completed percentage of this phase in the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`

	PhaseType *RecoveryPlanJobPhaseType `json:"phaseType,omitempty"`
	/*
	  List of stages in the phase.
	*/
	Stages []StageSummary `json:"stages,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
}

func (p *PhaseSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PhaseSummary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *PhaseSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PhaseSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPhaseSummary()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.PhaseType != nil {
		p.PhaseType = known.PhaseType
	}
	if known.Stages != nil {
		p.Stages = known.Stages
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "percentageComplete")
	delete(allFields, "phaseType")
	delete(allFields, "stages")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPhaseSummary() *PhaseSummary {
	p := new(PhaseSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.PhaseSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
Once a VM or volume group is associated with some protection policy, the schedules in the protection policy kick in to achieve the specified recovery point objective. A protected resource represents the data protection view of such a VM or volume group. It contains information such as the restorable time ranges on the local Prism Central and the state of replication to the targets specified in all the applied protection policies.
*/
type ProtectedResource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Category key-value pairs associated with the protected resource at the time of protection. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is represented as 'dept/hr'.
	*/
	CategoryFqNames []string `json:"categoryFqNames,omitempty"`
	/*
	  External identifier of the Consistency group which the protected resource is part of.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
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
	Links []import4.ApiLink `json:"links,omitempty"`

	ReplicationStates []ReplicationState `json:"replicationStates,omitempty"`
	/*
	  The data protection details of the protected resource that are relevant to any of the sites in the local Prism Central, like the time ranges available for recovery.
	*/
	SiteProtectionInfo []SiteProtectionInfo `json:"siteProtectionInfo,omitempty"`

	SourceSiteReference *DataProtectionSiteReference `json:"sourceSiteReference,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ProtectedResource) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProtectedResource

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ProtectedResource) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectedResource
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectedResource()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryFqNames != nil {
		p.CategoryFqNames = known.CategoryFqNames
	}
	if known.ConsistencyGroupExtId != nil {
		p.ConsistencyGroupExtId = known.ConsistencyGroupExtId
	}
	if known.EntityExtId != nil {
		p.EntityExtId = known.EntityExtId
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ReplicationStates != nil {
		p.ReplicationStates = known.ReplicationStates
	}
	if known.SiteProtectionInfo != nil {
		p.SiteProtectionInfo = known.SiteProtectionInfo
	}
	if known.SourceSiteReference != nil {
		p.SourceSiteReference = known.SourceSiteReference
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryFqNames")
	delete(allFields, "consistencyGroupExtId")
	delete(allFields, "entityExtId")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "replicationStates")
	delete(allFields, "siteProtectionInfo")
	delete(allFields, "sourceSiteReference")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectedResource() *ProtectedResource {
	p := new(ProtectedResource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResource"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/protected-resources/{extId}/$actions/promote Post operation
*/
type ProtectedResourcePromoteApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfProtectedResourcePromoteApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ProtectedResourcePromoteApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProtectedResourcePromoteApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ProtectedResourcePromoteApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectedResourcePromoteApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectedResourcePromoteApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectedResourcePromoteApiResponse() *ProtectedResourcePromoteApiResponse {
	p := new(ProtectedResourcePromoteApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourcePromoteApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/protected-resources/{extId}/$actions/restore Post operation
*/
type ProtectedResourceRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfProtectedResourceRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ProtectedResourceRestoreApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ProtectedResourceRestoreApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ProtectedResourceRestoreApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectedResourceRestoreApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectedResourceRestoreApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectedResourceRestoreApiResponse() *ProtectedResourceRestoreApiResponse {
	p := new(ProtectedResourceRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourceRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  The external identifier of the cluster on which the entity has valid restorable time ranges. The restored entity is created on the same cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  UTC date and time in ISO 8601 format representing the time from when the state of the entity should be restored. This must be a valid time within the restorable time ranges for the protected resource.
	*/
	RestoreTime *time.Time `json:"restoreTime,omitempty"`
}

func (p *ProtectedResourceRestoreSpec) MarshalJSON() ([]byte, error) {
	type ProtectedResourceRestoreSpecProxy ProtectedResourceRestoreSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProtectedResourceRestoreSpecProxy
		ClusterExtId *string `json:"clusterExtId,omitempty"`
	}{
		ProtectedResourceRestoreSpecProxy: (*ProtectedResourceRestoreSpecProxy)(p),
		ClusterExtId:                      p.ClusterExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ProtectedResourceRestoreSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectedResourceRestoreSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectedResourceRestoreSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.RestoreTime != nil {
		p.RestoreTime = known.RestoreTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "restoreTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectedResourceRestoreSpec() *ProtectedResourceRestoreSpec {
	p := new(ProtectedResourceRestoreSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ProtectedResourceRestoreSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The restorable time range details used to recover the protected resource.
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

func (p *RecoveryInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryInfo

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RestorableTimeRanges != nil {
		p.RestorableTimeRanges = known.RestorableTimeRanges
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "restorableTimeRanges")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryInfo() *RecoveryInfo {
	p := new(RecoveryInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of action performed by the recovery plan job.
*/
type RecoveryPlanActionType int

const (
	RECOVERYPLANACTIONTYPE_UNKNOWN            RecoveryPlanActionType = 0
	RECOVERYPLANACTIONTYPE_REDACTED           RecoveryPlanActionType = 1
	RECOVERYPLANACTIONTYPE_VALIDATE           RecoveryPlanActionType = 2
	RECOVERYPLANACTIONTYPE_PLANNED_FAILOVER   RecoveryPlanActionType = 3
	RECOVERYPLANACTIONTYPE_UNPLANNED_FAILOVER RecoveryPlanActionType = 4
	RECOVERYPLANACTIONTYPE_TEST_FAILOVER      RecoveryPlanActionType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPlanActionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALIDATE",
		"PLANNED_FAILOVER",
		"UNPLANNED_FAILOVER",
		"TEST_FAILOVER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPlanActionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALIDATE",
		"PLANNED_FAILOVER",
		"UNPLANNED_FAILOVER",
		"TEST_FAILOVER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPlanActionType) index(name string) RecoveryPlanActionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALIDATE",
		"PLANNED_FAILOVER",
		"UNPLANNED_FAILOVER",
		"TEST_FAILOVER",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPlanActionType(idx)
		}
	}
	return RECOVERYPLANACTIONTYPE_UNKNOWN
}

func (e *RecoveryPlanActionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPlanActionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPlanActionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPlanActionType) Ref() *RecoveryPlanActionType {
	return &e
}

/*
A recovery plan job specification describing the recovery plan to be executed, failover directions and type of failover.
*/
type RecoveryPlanJob struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *RecoveryPlanActionType `json:"actionType,omitempty"`
	/*
	  Recovery plan job completion time in ISO-8601 format.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  Summary of all the phases involved in the recovery plan job execution:<br> - Pre-processing: Actions carried out prior to the recovery of entities, such as creating a network or synchronizing recovery points as part of a `PLANNED_FAILOVER` operation.<br> - Entity recovery: Steps performed for the recovery of the entities for the recovery plan job.<br> - Post-processing: Actions taken after the recovery of entities, such as creating a Consistency group or auditing the recovery plan job information.
	*/
	ExecutionPhases []PhaseSummary `json:"executionPhases,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []FailoverDirection `json:"failoverDirections,omitempty"`
	/*
	  Indicates whether the recovery plan job was initiated by the witness or not.
	*/
	IsInitiatedByWitness *bool `json:"isInitiatedByWitness,omitempty"`
	/*
	  Indicates whether to perform a live migration of virtual machines (VMs) during a `PLANNED_FAILOVER` operation. When not specified or specified as false, the migration is performed for all virtual machines (VMs) in offline mode. When specified as true, the migration is performed for all virtual machines (VMs) in a running state.
	*/
	IsLiveMigrateVMs *bool `json:"isLiveMigrateVMs,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  External identifier of the owner of the recovery plan job.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  Percentage complete for the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`
	/*
	  The external identifier of the recovery plan.
	*/
	RecoveryPlanExtId *string `json:"recoveryPlanExtId,omitempty"`
	/*
	  Point in time from which to restore the entities during an `UNPLANNED_FAILOVER` operation. Only ISO-8601 formatted timestamps are supported. For example, `2023-01-02T03:04:05Z`.<br> When specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the specified timestamp. If not specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the start of the Recovery plan job.
	*/
	RecoveryReferenceTime *time.Time `json:"recoveryReferenceTime,omitempty"`
	/*
	  Recovery plan job start time in ISO-8601 format.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	ValidationStatus *Summary `json:"validationStatus,omitempty"`
}

func (p *RecoveryPlanJob) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanJob

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanJob) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanJob
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanJob()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ActionType != nil {
		p.ActionType = known.ActionType
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.ExecutionPhases != nil {
		p.ExecutionPhases = known.ExecutionPhases
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.IsInitiatedByWitness != nil {
		p.IsInitiatedByWitness = known.IsInitiatedByWitness
	}
	if known.IsLiveMigrateVMs != nil {
		p.IsLiveMigrateVMs = known.IsLiveMigrateVMs
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.RecoveryPlanExtId != nil {
		p.RecoveryPlanExtId = known.RecoveryPlanExtId
	}
	if known.RecoveryReferenceTime != nil {
		p.RecoveryReferenceTime = known.RecoveryReferenceTime
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.ValidationStatus != nil {
		p.ValidationStatus = known.ValidationStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "actionType")
	delete(allFields, "endTime")
	delete(allFields, "executionPhases")
	delete(allFields, "extId")
	delete(allFields, "failoverDirections")
	delete(allFields, "isInitiatedByWitness")
	delete(allFields, "isLiveMigrateVMs")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "percentageComplete")
	delete(allFields, "recoveryPlanExtId")
	delete(allFields, "recoveryReferenceTime")
	delete(allFields, "startTime")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "validationStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanJob() *RecoveryPlanJob {
	p := new(RecoveryPlanJob)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanJob"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the execution status of a single entity (e.g., a VM) within a recovery plan job.
*/
type RecoveryPlanJobExecutionStatus int

const (
	RECOVERYPLANJOBEXECUTIONSTATUS_UNKNOWN                           RecoveryPlanJobExecutionStatus = 0
	RECOVERYPLANJOBEXECUTIONSTATUS_REDACTED                          RecoveryPlanJobExecutionStatus = 1
	RECOVERYPLANJOBEXECUTIONSTATUS_SUCCEEDED                         RecoveryPlanJobExecutionStatus = 2
	RECOVERYPLANJOBEXECUTIONSTATUS_RUNNING                           RecoveryPlanJobExecutionStatus = 3
	RECOVERYPLANJOBEXECUTIONSTATUS_ABORTED                           RecoveryPlanJobExecutionStatus = 4
	RECOVERYPLANJOBEXECUTIONSTATUS_FAILED                            RecoveryPlanJobExecutionStatus = 5
	RECOVERYPLANJOBEXECUTIONSTATUS_QUEUED                            RecoveryPlanJobExecutionStatus = 6
	RECOVERYPLANJOBEXECUTIONSTATUS_SUCCEEDED_WITH_VALIDATION_WARNING RecoveryPlanJobExecutionStatus = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPlanJobExecutionStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"RUNNING",
		"ABORTED",
		"FAILED",
		"QUEUED",
		"SUCCEEDED_WITH_VALIDATION_WARNING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPlanJobExecutionStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"RUNNING",
		"ABORTED",
		"FAILED",
		"QUEUED",
		"SUCCEEDED_WITH_VALIDATION_WARNING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPlanJobExecutionStatus) index(name string) RecoveryPlanJobExecutionStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUCCEEDED",
		"RUNNING",
		"ABORTED",
		"FAILED",
		"QUEUED",
		"SUCCEEDED_WITH_VALIDATION_WARNING",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPlanJobExecutionStatus(idx)
		}
	}
	return RECOVERYPLANJOBEXECUTIONSTATUS_UNKNOWN
}

func (e *RecoveryPlanJobExecutionStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPlanJobExecutionStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPlanJobExecutionStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPlanJobExecutionStatus) Ref() *RecoveryPlanJobExecutionStatus {
	return &e
}

/*
A step executes an action. Each step has its external identifier and the status of the action it is/was performing.
*/
type RecoveryPlanJobExecutionStep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of affected entities.
	*/
	AffectedEntities []ExecutionStepEntityReference `json:"affectedEntities,omitempty"`
	/*
	  Time in ISO-8601 format when the current step ended.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`

	ErrorMessage *import3.AppMessage `json:"errorMessage,omitempty"`
	/*
	  Result for an execution step.
	*/
	ExecutionStepResults []ExecutionStepResult `json:"executionStepResults,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  The completed percentage for a step in the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`

	Phase *RecoveryPlanJobPhaseType `json:"phase,omitempty"`
	/*
	  External identifier of a recovery plan stage.
	*/
	StageExtId *string `json:"stageExtId,omitempty"`
	/*
	  Time in ISO-8601 format when the current step started.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
	/*
	  User readable message for the action being performed for the current step.
	*/
	StepDescription *string `json:"stepDescription,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPlanJobExecutionStep) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanJobExecutionStep

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanJobExecutionStep) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanJobExecutionStep
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanJobExecutionStep()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.ExecutionStepResults != nil {
		p.ExecutionStepResults = known.ExecutionStepResults
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.Phase != nil {
		p.Phase = known.Phase
	}
	if known.StageExtId != nil {
		p.StageExtId = known.StageExtId
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.StepDescription != nil {
		p.StepDescription = known.StepDescription
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "endTime")
	delete(allFields, "errorMessage")
	delete(allFields, "executionStepResults")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "operationType")
	delete(allFields, "percentageComplete")
	delete(allFields, "phase")
	delete(allFields, "stageExtId")
	delete(allFields, "startTime")
	delete(allFields, "status")
	delete(allFields, "stepDescription")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanJobExecutionStep() *RecoveryPlanJobExecutionStep {
	p := new(RecoveryPlanJobExecutionStep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanJobExecutionStep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RecoveryPlanJobExecutionStepProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of affected entities.
	*/
	AffectedEntities []ExecutionStepEntityReference `json:"affectedEntities,omitempty"`
	/*
	  Time in ISO-8601 format when the current step ended.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`

	ErrorMessage *import3.AppMessage `json:"errorMessage,omitempty"`
	/*
	  Result for an execution step.
	*/
	ExecutionStepResults []ExecutionStepResult `json:"executionStepResults,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  The completed percentage for a step in the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`

	Phase *RecoveryPlanJobPhaseType `json:"phase,omitempty"`
	/*
	  External identifier of a recovery plan stage.
	*/
	StageExtId *string `json:"stageExtId,omitempty"`
	/*
	  Time in ISO-8601 format when the current step started.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
	/*
	  User readable message for the action being performed for the current step.
	*/
	StepDescription *string `json:"stepDescription,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPlanJobExecutionStepProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanJobExecutionStepProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanJobExecutionStepProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanJobExecutionStepProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanJobExecutionStepProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.ExecutionStepResults != nil {
		p.ExecutionStepResults = known.ExecutionStepResults
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.Phase != nil {
		p.Phase = known.Phase
	}
	if known.StageExtId != nil {
		p.StageExtId = known.StageExtId
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.StepDescription != nil {
		p.StepDescription = known.StepDescription
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "endTime")
	delete(allFields, "errorMessage")
	delete(allFields, "executionStepResults")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "operationType")
	delete(allFields, "percentageComplete")
	delete(allFields, "phase")
	delete(allFields, "stageExtId")
	delete(allFields, "startTime")
	delete(allFields, "status")
	delete(allFields, "stepDescription")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanJobExecutionStepProjection() *RecoveryPlanJobExecutionStepProjection {
	p := new(RecoveryPlanJobExecutionStepProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The phase type of the execution step in the Recovery Plan Job.
*/
type RecoveryPlanJobPhaseType int

const (
	RECOVERYPLANJOBPHASETYPE_UNKNOWN         RecoveryPlanJobPhaseType = 0
	RECOVERYPLANJOBPHASETYPE_REDACTED        RecoveryPlanJobPhaseType = 1
	RECOVERYPLANJOBPHASETYPE_PRE_PROCESSING  RecoveryPlanJobPhaseType = 2
	RECOVERYPLANJOBPHASETYPE_ENTITY_RECOVERY RecoveryPlanJobPhaseType = 3
	RECOVERYPLANJOBPHASETYPE_POST_PROCESSING RecoveryPlanJobPhaseType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RecoveryPlanJobPhaseType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRE_PROCESSING",
		"ENTITY_RECOVERY",
		"POST_PROCESSING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RecoveryPlanJobPhaseType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRE_PROCESSING",
		"ENTITY_RECOVERY",
		"POST_PROCESSING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RecoveryPlanJobPhaseType) index(name string) RecoveryPlanJobPhaseType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRE_PROCESSING",
		"ENTITY_RECOVERY",
		"POST_PROCESSING",
	}
	for idx := range names {
		if names[idx] == name {
			return RecoveryPlanJobPhaseType(idx)
		}
	}
	return RECOVERYPLANJOBPHASETYPE_UNKNOWN
}

func (e *RecoveryPlanJobPhaseType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RecoveryPlanJobPhaseType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RecoveryPlanJobPhaseType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RecoveryPlanJobPhaseType) Ref() *RecoveryPlanJobPhaseType {
	return &e
}

type RecoveryPlanJobProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActionType *RecoveryPlanActionType `json:"actionType,omitempty"`
	/*
	  Recovery plan job completion time in ISO-8601 format.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  Summary of all the phases involved in the recovery plan job execution:<br> - Pre-processing: Actions carried out prior to the recovery of entities, such as creating a network or synchronizing recovery points as part of a `PLANNED_FAILOVER` operation.<br> - Entity recovery: Steps performed for the recovery of the entities for the recovery plan job.<br> - Post-processing: Actions taken after the recovery of entities, such as creating a Consistency group or auditing the recovery plan job information.
	*/
	ExecutionPhases []PhaseSummary `json:"executionPhases,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of failover directions from source disaster recovery location to target disaster recovery location.<br> For example, when failing over virtual machines (VMs) and volume groups from clusters C1, C2 registered to domain manager PC1 on source location to cluster C3 registered to domain manager PC1 on target location, there must be one direction between source location (PC1, C1) to target location (PC1, C3) and another mapping between source location (PC1, C2) and target location (PC1, C3).<br> Domain manager is a required parameter while describing disaster recovery location in failover directions.<br> When creating a Recovery Plan Job across two domain managers, the source clusters are not required.<br> Failover direction contains clusters only when failing over between primary and secondary clusters registered to the same domain manager.
	*/
	FailoverDirections []FailoverDirection `json:"failoverDirections,omitempty"`
	/*
	  Indicates whether the recovery plan job was initiated by the witness or not.
	*/
	IsInitiatedByWitness *bool `json:"isInitiatedByWitness,omitempty"`
	/*
	  Indicates whether to perform a live migration of virtual machines (VMs) during a `PLANNED_FAILOVER` operation. When not specified or specified as false, the migration is performed for all virtual machines (VMs) in offline mode. When specified as true, the migration is performed for all virtual machines (VMs) in a running state.
	*/
	IsLiveMigrateVMs *bool `json:"isLiveMigrateVMs,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Recovery Plan Job.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  External identifier of the owner of the recovery plan job.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  Percentage complete for the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`
	/*
	  The external identifier of the recovery plan.
	*/
	RecoveryPlanExtId *string `json:"recoveryPlanExtId,omitempty"`
	/*
	  Point in time from which to restore the entities during an `UNPLANNED_FAILOVER` operation. Only ISO-8601 formatted timestamps are supported. For example, `2023-01-02T03:04:05Z`.<br> When specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the specified timestamp. If not specified, virtual machines (VMs) and volume groups are restored from the latest recovery points created on or before the start of the Recovery plan job.
	*/
	RecoveryReferenceTime *time.Time `json:"recoveryReferenceTime,omitempty"`
	/*
	  Recovery plan job start time in ISO-8601 format.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	ValidationStatus *Summary `json:"validationStatus,omitempty"`
}

func (p *RecoveryPlanJobProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanJobProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanJobProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanJobProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanJobProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ActionType != nil {
		p.ActionType = known.ActionType
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.ExecutionPhases != nil {
		p.ExecutionPhases = known.ExecutionPhases
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FailoverDirections != nil {
		p.FailoverDirections = known.FailoverDirections
	}
	if known.IsInitiatedByWitness != nil {
		p.IsInitiatedByWitness = known.IsInitiatedByWitness
	}
	if known.IsLiveMigrateVMs != nil {
		p.IsLiveMigrateVMs = known.IsLiveMigrateVMs
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.RecoveryPlanExtId != nil {
		p.RecoveryPlanExtId = known.RecoveryPlanExtId
	}
	if known.RecoveryReferenceTime != nil {
		p.RecoveryReferenceTime = known.RecoveryReferenceTime
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.ValidationStatus != nil {
		p.ValidationStatus = known.ValidationStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "actionType")
	delete(allFields, "endTime")
	delete(allFields, "executionPhases")
	delete(allFields, "extId")
	delete(allFields, "failoverDirections")
	delete(allFields, "isInitiatedByWitness")
	delete(allFields, "isLiveMigrateVMs")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "percentageComplete")
	delete(allFields, "recoveryPlanExtId")
	delete(allFields, "recoveryReferenceTime")
	delete(allFields, "startTime")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "validationStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanJobProjection() *RecoveryPlanJobProjection {
	p := new(RecoveryPlanJobProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanJobProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Message used to report recovery plan job errors or warnings.
*/
type RecoveryPlanValidationError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of affected entities.
	*/
	AffectedEntities []import6.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  Error code for API request failure.
	*/
	ErrorCode *string `json:"errorCode,omitempty"`

	ErrorGroup *ErrorGroup `json:"errorGroup,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Summary of the error.
	*/
	Message *string `json:"message,omitempty"`

	RootCauseAnalysis *RootCauseAnalysis `json:"rootCauseAnalysis,omitempty"`

	Severity *Severity `json:"severity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPlanValidationError) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanValidationError

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanValidationError) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanValidationError
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanValidationError()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.ErrorCode != nil {
		p.ErrorCode = known.ErrorCode
	}
	if known.ErrorGroup != nil {
		p.ErrorGroup = known.ErrorGroup
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.RootCauseAnalysis != nil {
		p.RootCauseAnalysis = known.RootCauseAnalysis
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "errorCode")
	delete(allFields, "errorGroup")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "rootCauseAnalysis")
	delete(allFields, "severity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanValidationError() *RecoveryPlanValidationError {
	p := new(RecoveryPlanValidationError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanValidationError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RecoveryPlanValidationErrorProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of affected entities.
	*/
	AffectedEntities []import6.EntityReference `json:"affectedEntities,omitempty"`
	/*
	  Error code for API request failure.
	*/
	ErrorCode *string `json:"errorCode,omitempty"`

	ErrorGroup *ErrorGroup `json:"errorGroup,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Summary of the error.
	*/
	Message *string `json:"message,omitempty"`

	RootCauseAnalysis *RootCauseAnalysis `json:"rootCauseAnalysis,omitempty"`

	Severity *Severity `json:"severity,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPlanValidationErrorProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPlanValidationErrorProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPlanValidationErrorProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPlanValidationErrorProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPlanValidationErrorProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AffectedEntities != nil {
		p.AffectedEntities = known.AffectedEntities
	}
	if known.ErrorCode != nil {
		p.ErrorCode = known.ErrorCode
	}
	if known.ErrorGroup != nil {
		p.ErrorGroup = known.ErrorGroup
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.RootCauseAnalysis != nil {
		p.RootCauseAnalysis = known.RootCauseAnalysis
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "affectedEntities")
	delete(allFields, "errorCode")
	delete(allFields, "errorGroup")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "message")
	delete(allFields, "rootCauseAnalysis")
	delete(allFields, "severity")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPlanValidationErrorProjection() *RecoveryPlanValidationErrorProjection {
	p := new(RecoveryPlanValidationErrorProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPlanValidationErrorProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the recovery point along with all the captured VM and volume group recovery points.
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
	  Indicates whether the recovery point is secure. Approvals are necessary for the manual deletion of secure recovery points.
	*/
	IsSecure *bool `json:"isSecure,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  List of location references where the recovery point is presently located.
	*/
	LocationReferences []LocationReference `json:"locationReferences,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A read only field inserted into the recovery point at the time of recovery point creation, indicating the external identifier of the user who created this recovery point.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`

	RecoveryPointType *import2.RecoveryPointType `json:"recoveryPointType,omitempty"`

	SourceLocation *DisasterRecoveryLocation `json:"sourceLocation,omitempty"`

	Status *import2.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of this recovery point, which is the total space that could be reclaimed after deleting this recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
	/*
	  List of VM recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 32 entities. These entities can be a combination of VMs and volume groups.
	*/
	VmRecoveryPoints []VmRecoveryPoint `json:"vmRecoveryPoints,omitempty"`
	/*
	  List of volume group recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 32 entities. These entities can be a combination of VMs and volume groups.
	*/
	VolumeGroupRecoveryPoints []VolumeGroupRecoveryPoint `json:"volumeGroupRecoveryPoints,omitempty"`
}

func (p *RecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPoint

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsSecure != nil {
		p.IsSecure = known.IsSecure
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.LocationReferences != nil {
		p.LocationReferences = known.LocationReferences
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.SourceLocation != nil {
		p.SourceLocation = known.SourceLocation
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}
	if known.VmRecoveryPoints != nil {
		p.VmRecoveryPoints = known.VmRecoveryPoints
	}
	if known.VolumeGroupRecoveryPoints != nil {
		p.VolumeGroupRecoveryPoints = known.VolumeGroupRecoveryPoints
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "creationTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "isSecure")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "locationReferences")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "recoveryPointType")
	delete(allFields, "sourceLocation")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")
	delete(allFields, "vmRecoveryPoints")
	delete(allFields, "volumeGroupRecoveryPoints")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPoint() *RecoveryPoint {
	p := new(RecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  Indicates whether the recovery point is secure. Approvals are necessary for the manual deletion of secure recovery points.
	*/
	IsSecure *bool `json:"isSecure,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  List of location references where the recovery point is presently located.
	*/
	LocationReferences []LocationReference `json:"locationReferences,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A read only field inserted into the recovery point at the time of recovery point creation, indicating the external identifier of the user who created this recovery point.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`

	RecoveryPointType *import2.RecoveryPointType `json:"recoveryPointType,omitempty"`

	SourceLocation *DisasterRecoveryLocation `json:"sourceLocation,omitempty"`

	Status *import2.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of this recovery point, which is the total space that could be reclaimed after deleting this recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
	/*
	  List of VM recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 32 entities. These entities can be a combination of VMs and volume groups.
	*/
	VmRecoveryPoints []VmRecoveryPoint `json:"vmRecoveryPoints,omitempty"`
	/*
	  List of volume group recovery point that are a part of the specified top-level recovery point. Note that a recovery point can contain a maximum number of 32 entities. These entities can be a combination of VMs and volume groups.
	*/
	VolumeGroupRecoveryPoints []VolumeGroupRecoveryPoint `json:"volumeGroupRecoveryPoints,omitempty"`
}

func (p *RecoveryPointProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsSecure != nil {
		p.IsSecure = known.IsSecure
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.LocationReferences != nil {
		p.LocationReferences = known.LocationReferences
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.SourceLocation != nil {
		p.SourceLocation = known.SourceLocation
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}
	if known.VmRecoveryPoints != nil {
		p.VmRecoveryPoints = known.VmRecoveryPoints
	}
	if known.VolumeGroupRecoveryPoints != nil {
		p.VolumeGroupRecoveryPoints = known.VolumeGroupRecoveryPoints
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "creationTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "isSecure")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "locationReferences")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "recoveryPointType")
	delete(allFields, "sourceLocation")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")
	delete(allFields, "vmRecoveryPoints")
	delete(allFields, "volumeGroupRecoveryPoints")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointProjection() *RecoveryPointProjection {
	p := new(RecoveryPointProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId}/$actions/replicate Post operation
*/
type RecoveryPointReplicateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointReplicateApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RecoveryPointReplicateApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointReplicateApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointReplicateApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointReplicateApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointReplicateApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointReplicateApiResponse() *RecoveryPointReplicateApiResponse {
	p := new(RecoveryPointReplicateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointReplicateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  External identifier of the Prism Central (also known as Domain Manager).
	*/
	PcExtId *string `json:"pcExtId,omitempty"`
}

func (p *RecoveryPointReplicationSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointReplicationSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointReplicationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointReplicationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointReplicationSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.PcExtId != nil {
		p.PcExtId = known.PcExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "pcExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointReplicationSpec() *RecoveryPointReplicationSpec {
	p := new(RecoveryPointReplicationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointReplicationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the recovery point repository.
	*/
	Name *string `json:"name,omitempty"`
	/*

	 */
	ObjectStorageReferenceItemDiscriminator_ *string `json:"$objectStorageReferenceItemDiscriminator,omitempty"`
	/*
	  Object storage reference details.
	*/
	ObjectStorageReference *OneOfRecoveryPointRepositoryObjectStorageReference `json:"objectStorageReference,omitempty"`

	ObjectStorageType *ObjectStorageType `json:"objectStorageType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPointRepository) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointRepository

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointRepository) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointRepository
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointRepository()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ObjectStorageReferenceItemDiscriminator_ != nil {
		p.ObjectStorageReferenceItemDiscriminator_ = known.ObjectStorageReferenceItemDiscriminator_
	}
	if known.ObjectStorageReference != nil {
		p.ObjectStorageReference = known.ObjectStorageReference
	}
	if known.ObjectStorageType != nil {
		p.ObjectStorageType = known.ObjectStorageType
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "$objectStorageReferenceItemDiscriminator")
	delete(allFields, "objectStorageReference")
	delete(allFields, "objectStorageType")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointRepository() *RecoveryPointRepository {
	p := new(RecoveryPointRepository)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRepository"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the recovery point repository.
	*/
	Name *string `json:"name,omitempty"`

	ObjectStorageReferenceItemDiscriminator_ *string `json:"$objectStorageReferenceItemDiscriminator,omitempty"`
	/*
	  Object storage reference details.
	*/
	ObjectStorageReference *OneOfRecoveryPointRepositoryProjectionObjectStorageReference `json:"objectStorageReference,omitempty"`

	ObjectStorageType *ObjectStorageType `json:"objectStorageType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecoveryPointRepositoryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointRepositoryProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointRepositoryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointRepositoryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointRepositoryProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ObjectStorageReferenceItemDiscriminator_ != nil {
		p.ObjectStorageReferenceItemDiscriminator_ = known.ObjectStorageReferenceItemDiscriminator_
	}
	if known.ObjectStorageReference != nil {
		p.ObjectStorageReference = known.ObjectStorageReference
	}
	if known.ObjectStorageType != nil {
		p.ObjectStorageType = known.ObjectStorageType
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "$objectStorageReferenceItemDiscriminator")
	delete(allFields, "objectStorageReference")
	delete(allFields, "objectStorageType")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointRepositoryProjection() *RecoveryPointRepositoryProjection {
	p := new(RecoveryPointRepositoryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRepositoryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification for the restore action on the top-level recovery point. For a recovery point that contains multiple VM or volume group recovery points, users can selectively trigger restore for any specific set of VM or volume group recovery points. In case no particular selection is made, all VM and volume group recovery points that are a part of the top-level recovery point are restored.
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
	  List of specifications to restore a specific VM recovery points that are a part of the top-level recovery point. A specific VM recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VmRecoveryPointRestoreOverrides []VmRecoveryPointRestoreOverride `json:"vmRecoveryPointRestoreOverrides,omitempty"`
	/*
	  List of specifications to restore a specific volume group recovery points that are a part of the top-level recovery point. A specific volume group recovery point can be selected for restore by specifying its external identifier along with override specification (if any).
	*/
	VolumeGroupRecoveryPointRestoreOverrides []VolumeGroupRecoveryPointRestoreOverride `json:"volumeGroupRecoveryPointRestoreOverrides,omitempty"`
}

func (p *RecoveryPointRestorationSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointRestorationSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointRestorationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointRestorationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointRestorationSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.VmRecoveryPointRestoreOverrides != nil {
		p.VmRecoveryPointRestoreOverrides = known.VmRecoveryPointRestoreOverrides
	}
	if known.VolumeGroupRecoveryPointRestoreOverrides != nil {
		p.VolumeGroupRecoveryPointRestoreOverrides = known.VolumeGroupRecoveryPointRestoreOverrides
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "vmRecoveryPointRestoreOverrides")
	delete(allFields, "volumeGroupRecoveryPointRestoreOverrides")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointRestorationSpec() *RecoveryPointRestorationSpec {
	p := new(RecoveryPointRestorationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRestorationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId}/$actions/restore Post operation
*/
type RecoveryPointRestoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoveryPointRestoreApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RecoveryPointRestoreApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryPointRestoreApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecoveryPointRestoreApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryPointRestoreApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryPointRestoreApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryPointRestoreApiResponse() *RecoveryPointRestoreApiResponse {
	p := new(RecoveryPointRestoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecoveryPointRestoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
Details of the Recycle Bin entry along with the entity details.
*/
type RecycleBinEntry struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the categories attached to the entity at the time of deletion.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the user who deleted the entity.
	*/
	DeletedByUserExtId *string `json:"deletedByUserExtId,omitempty"`
	/*
	  Entity deletion time in ISO-8601 format.
	*/
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	/*
	  Absolute time in ISO-8601 format until which the entity's (VM/Volume group) retention is ensured by the Recycle Bin.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Space consumed (in bytes) by the deleted entity in the Recycle Bin.
	*/
	SpaceConsumedBytes *int64 `json:"spaceConsumedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecycleBinEntry) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecycleBinEntry

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecycleBinEntry) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecycleBinEntry
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecycleBinEntry()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryIds != nil {
		p.CategoryIds = known.CategoryIds
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DeletedByUserExtId != nil {
		p.DeletedByUserExtId = known.DeletedByUserExtId
	}
	if known.DeletionTime != nil {
		p.DeletionTime = known.DeletionTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SpaceConsumedBytes != nil {
		p.SpaceConsumedBytes = known.SpaceConsumedBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryIds")
	delete(allFields, "clusterExtId")
	delete(allFields, "deletedByUserExtId")
	delete(allFields, "deletionTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "spaceConsumedBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecycleBinEntry() *RecycleBinEntry {
	p := new(RecycleBinEntry)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecycleBinEntry"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the VM in Recycle Bin.
*/
type RecycleBinVm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the categories attached to the entity at the time of deletion.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the user who deleted the entity.
	*/
	DeletedByUserExtId *string `json:"deletedByUserExtId,omitempty"`
	/*
	  Entity deletion time in ISO-8601 format.
	*/
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	/*
	  Absolute time in ISO-8601 format until which the entity's (VM/Volume group) retention is ensured by the Recycle Bin.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Space consumed (in bytes) by the deleted entity in the Recycle Bin.
	*/
	SpaceConsumedBytes *int64 `json:"spaceConsumedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VM external identifier which is captured as part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
	/*
	  Name of the deleted VM.
	*/
	VmName *string `json:"vmName,omitempty"`
	/*
	  External identifier of the user who owned the VM before it was deleted.
	*/
	VmOwnerExtId *string `json:"vmOwnerExtId,omitempty"`
}

func (p *RecycleBinVm) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecycleBinVm

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecycleBinVm) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecycleBinVm
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecycleBinVm()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryIds != nil {
		p.CategoryIds = known.CategoryIds
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DeletedByUserExtId != nil {
		p.DeletedByUserExtId = known.DeletedByUserExtId
	}
	if known.DeletionTime != nil {
		p.DeletionTime = known.DeletionTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SpaceConsumedBytes != nil {
		p.SpaceConsumedBytes = known.SpaceConsumedBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.VmExtId != nil {
		p.VmExtId = known.VmExtId
	}
	if known.VmName != nil {
		p.VmName = known.VmName
	}
	if known.VmOwnerExtId != nil {
		p.VmOwnerExtId = known.VmOwnerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryIds")
	delete(allFields, "clusterExtId")
	delete(allFields, "deletedByUserExtId")
	delete(allFields, "deletionTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "spaceConsumedBytes")
	delete(allFields, "tenantId")
	delete(allFields, "vmExtId")
	delete(allFields, "vmName")
	delete(allFields, "vmOwnerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecycleBinVm() *RecycleBinVm {
	p := new(RecycleBinVm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecycleBinVm"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RecycleBinVmProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the categories attached to the entity at the time of deletion.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the user who deleted the entity.
	*/
	DeletedByUserExtId *string `json:"deletedByUserExtId,omitempty"`
	/*
	  Entity deletion time in ISO-8601 format.
	*/
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	/*
	  Absolute time in ISO-8601 format until which the entity's (VM/Volume group) retention is ensured by the Recycle Bin.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Space consumed (in bytes) by the deleted entity in the Recycle Bin.
	*/
	SpaceConsumedBytes *int64 `json:"spaceConsumedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VM external identifier which is captured as part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
	/*
	  Name of the deleted VM.
	*/
	VmName *string `json:"vmName,omitempty"`
	/*
	  External identifier of the user who owned the VM before it was deleted.
	*/
	VmOwnerExtId *string `json:"vmOwnerExtId,omitempty"`
}

func (p *RecycleBinVmProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecycleBinVmProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecycleBinVmProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecycleBinVmProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecycleBinVmProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryIds != nil {
		p.CategoryIds = known.CategoryIds
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DeletedByUserExtId != nil {
		p.DeletedByUserExtId = known.DeletedByUserExtId
	}
	if known.DeletionTime != nil {
		p.DeletionTime = known.DeletionTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SpaceConsumedBytes != nil {
		p.SpaceConsumedBytes = known.SpaceConsumedBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.VmExtId != nil {
		p.VmExtId = known.VmExtId
	}
	if known.VmName != nil {
		p.VmName = known.VmName
	}
	if known.VmOwnerExtId != nil {
		p.VmOwnerExtId = known.VmOwnerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryIds")
	delete(allFields, "clusterExtId")
	delete(allFields, "deletedByUserExtId")
	delete(allFields, "deletionTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "spaceConsumedBytes")
	delete(allFields, "tenantId")
	delete(allFields, "vmExtId")
	delete(allFields, "vmName")
	delete(allFields, "vmOwnerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecycleBinVmProjection() *RecycleBinVmProjection {
	p := new(RecycleBinVmProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecycleBinVmProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details about the volume group in Recycle Bin.
*/
type RecycleBinVolumeGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the categories attached to the entity at the time of deletion.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the user who deleted the entity.
	*/
	DeletedByUserExtId *string `json:"deletedByUserExtId,omitempty"`
	/*
	  Entity deletion time in ISO-8601 format.
	*/
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	/*
	  Absolute time in ISO-8601 format until which the entity's (VM/Volume group) retention is ensured by the Recycle Bin.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Space consumed (in bytes) by the deleted entity in the Recycle Bin.
	*/
	SpaceConsumedBytes *int64 `json:"spaceConsumedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Volume Group external identifier which is captured as part of this recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	/*
	  Name of the deleted volume group.
	*/
	VolumeGroupName *string `json:"volumeGroupName,omitempty"`
	/*
	  External identifier of the user who owned the volume group before it was deleted.
	*/
	VolumeGroupOwnerExtId *string `json:"volumeGroupOwnerExtId,omitempty"`
}

func (p *RecycleBinVolumeGroup) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecycleBinVolumeGroup

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecycleBinVolumeGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecycleBinVolumeGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecycleBinVolumeGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryIds != nil {
		p.CategoryIds = known.CategoryIds
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DeletedByUserExtId != nil {
		p.DeletedByUserExtId = known.DeletedByUserExtId
	}
	if known.DeletionTime != nil {
		p.DeletionTime = known.DeletionTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SpaceConsumedBytes != nil {
		p.SpaceConsumedBytes = known.SpaceConsumedBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.VolumeGroupExtId != nil {
		p.VolumeGroupExtId = known.VolumeGroupExtId
	}
	if known.VolumeGroupName != nil {
		p.VolumeGroupName = known.VolumeGroupName
	}
	if known.VolumeGroupOwnerExtId != nil {
		p.VolumeGroupOwnerExtId = known.VolumeGroupOwnerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryIds")
	delete(allFields, "clusterExtId")
	delete(allFields, "deletedByUserExtId")
	delete(allFields, "deletionTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "spaceConsumedBytes")
	delete(allFields, "tenantId")
	delete(allFields, "volumeGroupExtId")
	delete(allFields, "volumeGroupName")
	delete(allFields, "volumeGroupOwnerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecycleBinVolumeGroup() *RecycleBinVolumeGroup {
	p := new(RecycleBinVolumeGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecycleBinVolumeGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RecycleBinVolumeGroupProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the categories attached to the entity at the time of deletion.
	*/
	CategoryIds []string `json:"categoryIds,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  External identifier of the user who deleted the entity.
	*/
	DeletedByUserExtId *string `json:"deletedByUserExtId,omitempty"`
	/*
	  Entity deletion time in ISO-8601 format.
	*/
	DeletionTime *time.Time `json:"deletionTime,omitempty"`
	/*
	  Absolute time in ISO-8601 format until which the entity's (VM/Volume group) retention is ensured by the Recycle Bin.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Space consumed (in bytes) by the deleted entity in the Recycle Bin.
	*/
	SpaceConsumedBytes *int64 `json:"spaceConsumedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Volume Group external identifier which is captured as part of this recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	/*
	  Name of the deleted volume group.
	*/
	VolumeGroupName *string `json:"volumeGroupName,omitempty"`
	/*
	  External identifier of the user who owned the volume group before it was deleted.
	*/
	VolumeGroupOwnerExtId *string `json:"volumeGroupOwnerExtId,omitempty"`
}

func (p *RecycleBinVolumeGroupProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecycleBinVolumeGroupProjection

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RecycleBinVolumeGroupProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecycleBinVolumeGroupProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecycleBinVolumeGroupProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CategoryIds != nil {
		p.CategoryIds = known.CategoryIds
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DeletedByUserExtId != nil {
		p.DeletedByUserExtId = known.DeletedByUserExtId
	}
	if known.DeletionTime != nil {
		p.DeletionTime = known.DeletionTime
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SpaceConsumedBytes != nil {
		p.SpaceConsumedBytes = known.SpaceConsumedBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.VolumeGroupExtId != nil {
		p.VolumeGroupExtId = known.VolumeGroupExtId
	}
	if known.VolumeGroupName != nil {
		p.VolumeGroupName = known.VolumeGroupName
	}
	if known.VolumeGroupOwnerExtId != nil {
		p.VolumeGroupOwnerExtId = known.VolumeGroupOwnerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryIds")
	delete(allFields, "clusterExtId")
	delete(allFields, "deletedByUserExtId")
	delete(allFields, "deletionTime")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "spaceConsumedBytes")
	delete(allFields, "tenantId")
	delete(allFields, "volumeGroupExtId")
	delete(allFields, "volumeGroupName")
	delete(allFields, "volumeGroupOwnerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecycleBinVolumeGroupProjection() *RecycleBinVolumeGroupProjection {
	p := new(RecycleBinVolumeGroupProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RecycleBinVolumeGroupProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Replication-related information about the protected resource.
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

func (p *ReplicationState) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReplicationState

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ReplicationState) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReplicationState
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReplicationState()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProtectionPolicyExtId != nil {
		p.ProtectionPolicyExtId = known.ProtectionPolicyExtId
	}
	if known.RecoveryPointObjectiveSeconds != nil {
		p.RecoveryPointObjectiveSeconds = known.RecoveryPointObjectiveSeconds
	}
	if known.ReplicationStatus != nil {
		p.ReplicationStatus = known.ReplicationStatus
	}
	if known.TargetSiteReference != nil {
		p.TargetSiteReference = known.TargetSiteReference
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "protectionPolicyExtId")
	delete(allFields, "recoveryPointObjectiveSeconds")
	delete(allFields, "replicationStatus")
	delete(allFields, "targetSiteReference")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReplicationState() *ReplicationState {
	p := new(ReplicationState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.ReplicationState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The start and end time details, both inclusive, represent the range of time during which the entity is restorable.
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

func (p *RestorableTimeRange) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RestorableTimeRange

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RestorableTimeRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RestorableTimeRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRestorableTimeRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndTime != nil {
		p.EndTime = known.EndTime
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endTime")
	delete(allFields, "startTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRestorableTimeRange() *RestorableTimeRange {
	p := new(RestorableTimeRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RestorableTimeRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Possible causes, resolutions, impact and additional details to troubleshoot the error message.
*/
type RootCauseAnalysis struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cause of the validation errors or warnings.
	*/
	Cause *string `json:"cause,omitempty"`
	/*
	  List of impacts of the errors or warnings raised.
	*/
	Impacts []string `json:"impacts,omitempty"`
	/*
	  Steps to resolve the errors or warnings.
	*/
	Resolutions []string `json:"resolutions,omitempty"`
}

func (p *RootCauseAnalysis) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RootCauseAnalysis

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *RootCauseAnalysis) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RootCauseAnalysis
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRootCauseAnalysis()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cause != nil {
		p.Cause = known.Cause
	}
	if known.Impacts != nil {
		p.Impacts = known.Impacts
	}
	if known.Resolutions != nil {
		p.Resolutions = known.Resolutions
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cause")
	delete(allFields, "impacts")
	delete(allFields, "resolutions")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRootCauseAnalysis() *RootCauseAnalysis {
	p := new(RootCauseAnalysis)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.RootCauseAnalysis"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Severity of an error message.
*/
type Severity int

const (
	SEVERITY_UNKNOWN  Severity = 0
	SEVERITY_REDACTED Severity = 1
	SEVERITY_ERROR    Severity = 2
	SEVERITY_WARNING  Severity = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Severity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ERROR",
		"WARNING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Severity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ERROR",
		"WARNING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Severity) index(name string) Severity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ERROR",
		"WARNING",
	}
	for idx := range names {
		if names[idx] == name {
			return Severity(idx)
		}
	}
	return SEVERITY_UNKNOWN
}

func (e *Severity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Severity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Severity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Severity) Ref() *Severity {
	return &e
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

	SynchronousReplicationRole *SynchronousReplicationRole `json:"synchronousReplicationRole,omitempty"`
}

func (p *SiteProtectionInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SiteProtectionInfo

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *SiteProtectionInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SiteProtectionInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSiteProtectionInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LocationReference != nil {
		p.LocationReference = known.LocationReference
	}
	if known.RecoveryInfo != nil {
		p.RecoveryInfo = known.RecoveryInfo
	}
	if known.SynchronousReplicationRole != nil {
		p.SynchronousReplicationRole = known.SynchronousReplicationRole
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "locationReference")
	delete(allFields, "recoveryInfo")
	delete(allFields, "synchronousReplicationRole")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSiteProtectionInfo() *SiteProtectionInfo {
	p := new(SiteProtectionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SiteProtectionInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *SiteReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SiteReference

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *SiteReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SiteReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSiteReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AvailabilityZoneReference != nil {
		p.AvailabilityZoneReference = known.AvailabilityZoneReference
	}
	if known.ClusterReference != nil {
		p.ClusterReference = known.ClusterReference
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availabilityZoneReference")
	delete(allFields, "clusterReference")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSiteReference() *SiteReference {
	p := new(SiteReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.SiteReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Recovery stage summary in the recovery plan job execution.
*/
type StageSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The completed percentage of this phase in the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`
	/*
	  External identifier of a recovery plan stage.
	*/
	StageExtId *string `json:"stageExtId,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
}

func (p *StageSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StageSummary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *StageSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StageSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStageSummary()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.StageExtId != nil {
		p.StageExtId = known.StageExtId
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "percentageComplete")
	delete(allFields, "stageExtId")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStageSummary() *StageSummary {
	p := new(StageSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.StageSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Recovery plan job summary. This is a generic model to represent the summary in the recovery plan job execution.
*/
type Summary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The completed percentage of this phase in the recovery plan job.
	*/
	PercentageComplete *int `json:"percentageComplete,omitempty"`

	Status *RecoveryPlanJobExecutionStatus `json:"status,omitempty"`
}

func (p *Summary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Summary

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Summary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Summary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSummary()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PercentageComplete != nil {
		p.PercentageComplete = known.PercentageComplete
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "percentageComplete")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSummary() *Summary {
	p := new(Summary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Summary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Synchronous replication role-related information of the protected resource.
*/
type SynchronousReplicationRole int

const (
	SYNCHRONOUSREPLICATIONROLE_UNKNOWN     SynchronousReplicationRole = 0
	SYNCHRONOUSREPLICATIONROLE_REDACTED    SynchronousReplicationRole = 1
	SYNCHRONOUSREPLICATIONROLE_PRIMARY     SynchronousReplicationRole = 2
	SYNCHRONOUSREPLICATIONROLE_SECONDARY   SynchronousReplicationRole = 3
	SYNCHRONOUSREPLICATIONROLE_INDEPENDENT SynchronousReplicationRole = 4
	SYNCHRONOUSREPLICATIONROLE_DECOUPLED   SynchronousReplicationRole = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SynchronousReplicationRole) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
		"INDEPENDENT",
		"DECOUPLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SynchronousReplicationRole) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
		"INDEPENDENT",
		"DECOUPLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SynchronousReplicationRole) index(name string) SynchronousReplicationRole {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY",
		"SECONDARY",
		"INDEPENDENT",
		"DECOUPLED",
	}
	for idx := range names {
		if names[idx] == name {
			return SynchronousReplicationRole(idx)
		}
	}
	return SYNCHRONOUSREPLICATIONROLE_UNKNOWN
}

func (e *SynchronousReplicationRole) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SynchronousReplicationRole:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SynchronousReplicationRole) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SynchronousReplicationRole) Ref() *SynchronousReplicationRole {
	return &e
}

/*
REST response for all response codes in API path /dataprotection/v4.2/config/consistency-groups/{extId} Put operation
*/
type UpdateConsistencyGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateConsistencyGroupApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateConsistencyGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateConsistencyGroupApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *UpdateConsistencyGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateConsistencyGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateConsistencyGroupApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateConsistencyGroupApiResponse() *UpdateConsistencyGroupApiResponse {
	p := new(UpdateConsistencyGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.UpdateConsistencyGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
REST response for all response codes in API path /dataprotection/v4.2/config/recovery-points/{extId}/$actions/set-expiration-time Post operation
*/
type UpdateRecoveryPointExpirationTimeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRecoveryPointExpirationTimeApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateRecoveryPointExpirationTimeApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateRecoveryPointExpirationTimeApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *UpdateRecoveryPointExpirationTimeApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateRecoveryPointExpirationTimeApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateRecoveryPointExpirationTimeApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateRecoveryPointExpirationTimeApiResponse() *UpdateRecoveryPointExpirationTimeApiResponse {
	p := new(UpdateRecoveryPointExpirationTimeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.UpdateRecoveryPointExpirationTimeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	/*

	 */
	ApplicationConsistentPropertiesItemDiscriminator_ *string `json:"$applicationConsistentPropertiesItemDiscriminator,omitempty"`
	/*
	  User-defined application-consistent properties for the recovery point.
	*/
	ApplicationConsistentProperties *OneOfVmRecoveryPointApplicationConsistentProperties `json:"applicationConsistentProperties,omitempty"`
	/*
	  External identifier of the Consistency group which the VM was part of at the time of recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the Recovery point is created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`

	DiskRecoveryPoints []import2.DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
	/*
	  The UTC date and time in ISO-8601 format when the current Recovery point expires and will be garbage collected.
	*/
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the Recovery point.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  The name of the Recovery point.
	*/
	Name *string `json:"name,omitempty"`

	RecoveryPointType *import2.RecoveryPointType `json:"recoveryPointType,omitempty"`

	Status *import2.RecoveryPointStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of this recovery point, which is the total space that could be reclaimed after deleting this recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
	/*
	  Category key-value pairs associated with the VM at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is displayed as 'dept/hr'.
	*/
	VmCategories []string `json:"vmCategories,omitempty"`
	/*
	  VM external identifier which is captured as a part of this recovery point.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
}

func (p *VmRecoveryPoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRecoveryPoint

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VmRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVmRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApplicationConsistentPropertiesItemDiscriminator_ != nil {
		p.ApplicationConsistentPropertiesItemDiscriminator_ = known.ApplicationConsistentPropertiesItemDiscriminator_
	}
	if known.ApplicationConsistentProperties != nil {
		p.ApplicationConsistentProperties = known.ApplicationConsistentProperties
	}
	if known.ConsistencyGroupExtId != nil {
		p.ConsistencyGroupExtId = known.ConsistencyGroupExtId
	}
	if known.CreationTime != nil {
		p.CreationTime = known.CreationTime
	}
	if known.DiskRecoveryPoints != nil {
		p.DiskRecoveryPoints = known.DiskRecoveryPoints
	}
	if known.ExpirationTime != nil {
		p.ExpirationTime = known.ExpirationTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HypervisorType != nil {
		p.HypervisorType = known.HypervisorType
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.RecoveryPointType != nil {
		p.RecoveryPointType = known.RecoveryPointType
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}
	if known.VmCategories != nil {
		p.VmCategories = known.VmCategories
	}
	if known.VmExtId != nil {
		p.VmExtId = known.VmExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$applicationConsistentPropertiesItemDiscriminator")
	delete(allFields, "applicationConsistentProperties")
	delete(allFields, "consistencyGroupExtId")
	delete(allFields, "creationTime")
	delete(allFields, "diskRecoveryPoints")
	delete(allFields, "expirationTime")
	delete(allFields, "extId")
	delete(allFields, "hypervisorType")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "name")
	delete(allFields, "recoveryPointType")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")
	delete(allFields, "vmCategories")
	delete(allFields, "vmExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVmRecoveryPoint() *VmRecoveryPoint {
	p := new(VmRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmRecoveryPointRestoreOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  If set to true, any VM configuration that cannot be restored or overridden, is dropped or reset to the system defaults and the corresponding warnings are returned. If set to False, the restore operation runs in the normal mode where any failures are raised as errors and the VM is not restored until the issues are fixed. Only applicable for AHV VMs.
	*/
	IsStrictMode *bool `json:"isStrictMode,omitempty"`

	VmOverrideSpecItemDiscriminator_ *string `json:"$vmOverrideSpecItemDiscriminator,omitempty"`
	/*
	  AHV/ESXi specific VM restore override specifications.
	*/
	VmOverrideSpec *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec `json:"vmOverrideSpec,omitempty"`
	/*
	  External identifier of a VM recovery point, that is a part of the top-level recovery point.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
}

func (p *VmRecoveryPointRestoreOverride) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmRecoveryPointRestoreOverride

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VmRecoveryPointRestoreOverride) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmRecoveryPointRestoreOverride
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVmRecoveryPointRestoreOverride()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsStrictMode != nil {
		p.IsStrictMode = known.IsStrictMode
	}
	if known.VmOverrideSpecItemDiscriminator_ != nil {
		p.VmOverrideSpecItemDiscriminator_ = known.VmOverrideSpecItemDiscriminator_
	}
	if known.VmOverrideSpec != nil {
		p.VmOverrideSpec = known.VmOverrideSpec
	}
	if known.VmRecoveryPointExtId != nil {
		p.VmRecoveryPointExtId = known.VmRecoveryPointExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isStrictMode")
	delete(allFields, "$vmOverrideSpecItemDiscriminator")
	delete(allFields, "vmOverrideSpec")
	delete(allFields, "vmRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVmRecoveryPointRestoreOverride() *VmRecoveryPointRestoreOverride {
	p := new(VmRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VmRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsStrictMode = new(bool)
	*p.IsStrictMode = false

	return p
}

func (p *VmRecoveryPointRestoreOverride) GetVmOverrideSpec() interface{} {
	if nil == p.VmOverrideSpec {
		return nil
	}
	return p.VmOverrideSpec.GetValue()
}

func (p *VmRecoveryPointRestoreOverride) SetVmOverrideSpec(v interface{}) error {
	if nil == p.VmOverrideSpec {
		p.VmOverrideSpec = NewOneOfVmRecoveryPointRestoreOverrideVmOverrideSpec()
	}
	e := p.VmOverrideSpec.SetValue(v)
	if nil == e {
		if nil == p.VmOverrideSpecItemDiscriminator_ {
			p.VmOverrideSpecItemDiscriminator_ = new(string)
		}
		*p.VmOverrideSpecItemDiscriminator_ = *p.VmOverrideSpec.Discriminator
	}
	return e
}

/*
Protected resource/recovery point restore that overrides the volume group configuration. The specified properties are overridden for the restored volume group.
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

func (p *VolumeGroupOverrideSpec) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeGroupOverrideSpec

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VolumeGroupOverrideSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupOverrideSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVolumeGroupOverrideSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVolumeGroupOverrideSpec() *VolumeGroupOverrideSpec {
	p := new(VolumeGroupOverrideSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupOverrideSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	  External identifier of the Consistency Group, which the entity was part of at the time of recovery point creation.
	*/
	ConsistencyGroupExtId *string `json:"consistencyGroupExtId,omitempty"`

	DiskRecoveryPoints []import2.DiskRecoveryPoint `json:"diskRecoveryPoints,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Location agnostic identifier of the recovery point. This identifier is used to identify the same instances of a recovery point across different sites.
	*/
	LocationAgnosticId *string `json:"locationAgnosticId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Indicates the total exclusive usage of this recovery point, which is the total space that could be reclaimed after deleting this recovery point.
	*/
	TotalExclusiveUsageBytes *int64 `json:"totalExclusiveUsageBytes,omitempty"`
	/*
	  Category key-value pairs associated with the volume group at the time of recovery point creation. The category key and value are separated by '/'. For example, a category with key 'dept' and value 'hr' is represented as 'dept/hr'.
	*/
	VolumeGroupCategories []string `json:"volumeGroupCategories,omitempty"`
	/*
	  Volume Group external identifier which is captured as part of this recovery point.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId"`
}

func (p *VolumeGroupRecoveryPoint) MarshalJSON() ([]byte, error) {
	type VolumeGroupRecoveryPointProxy VolumeGroupRecoveryPoint

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VolumeGroupRecoveryPointProxy
		VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
	}{
		VolumeGroupRecoveryPointProxy: (*VolumeGroupRecoveryPointProxy)(p),
		VolumeGroupExtId:              p.VolumeGroupExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VolumeGroupRecoveryPoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupRecoveryPoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVolumeGroupRecoveryPoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConsistencyGroupExtId != nil {
		p.ConsistencyGroupExtId = known.ConsistencyGroupExtId
	}
	if known.DiskRecoveryPoints != nil {
		p.DiskRecoveryPoints = known.DiskRecoveryPoints
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationAgnosticId != nil {
		p.LocationAgnosticId = known.LocationAgnosticId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalExclusiveUsageBytes != nil {
		p.TotalExclusiveUsageBytes = known.TotalExclusiveUsageBytes
	}
	if known.VolumeGroupCategories != nil {
		p.VolumeGroupCategories = known.VolumeGroupCategories
	}
	if known.VolumeGroupExtId != nil {
		p.VolumeGroupExtId = known.VolumeGroupExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "consistencyGroupExtId")
	delete(allFields, "diskRecoveryPoints")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "locationAgnosticId")
	delete(allFields, "tenantId")
	delete(allFields, "totalExclusiveUsageBytes")
	delete(allFields, "volumeGroupCategories")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVolumeGroupRecoveryPoint() *VolumeGroupRecoveryPoint {
	p := new(VolumeGroupRecoveryPoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *VolumeGroupRecoveryPointRestoreOverride) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeGroupRecoveryPointRestoreOverride

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VolumeGroupRecoveryPointRestoreOverride) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupRecoveryPointRestoreOverride
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVolumeGroupRecoveryPointRestoreOverride()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.VolumeGroupOverrideSpec != nil {
		p.VolumeGroupOverrideSpec = known.VolumeGroupOverrideSpec
	}
	if known.VolumeGroupRecoveryPointExtId != nil {
		p.VolumeGroupRecoveryPointExtId = known.VolumeGroupRecoveryPointExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "volumeGroupOverrideSpec")
	delete(allFields, "volumeGroupRecoveryPointExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVolumeGroupRecoveryPointRestoreOverride() *VolumeGroupRecoveryPointRestoreOverride {
	p := new(VolumeGroupRecoveryPointRestoreOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupRecoveryPointRestoreOverride"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`

	PrimarySite *SiteReference `json:"primarySite,omitempty"`

	SecondarySite *SiteReference `json:"secondarySite,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VolumeGroupSyncContext) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeGroupSyncContext

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *VolumeGroupSyncContext) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupSyncContext
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVolumeGroupSyncContext()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.PrimarySite != nil {
		p.PrimarySite = known.PrimarySite
	}
	if known.SecondarySite != nil {
		p.SecondarySite = known.SecondarySite
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "primarySite")
	delete(allFields, "secondarySite")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVolumeGroupSyncContext() *VolumeGroupSyncContext {
	p := new(VolumeGroupSyncContext)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.VolumeGroupSyncContext"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Witness site.
	*/
	Name *string `json:"name,omitempty"`

	Status *WitnessAvailabilityStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Witness) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Witness

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *Witness) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Witness
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWitness()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostReferences != nil {
		p.HostReferences = known.HostReferences
	}
	if known.IpAddresses != nil {
		p.IpAddresses = known.IpAddresses
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "hostReferences")
	delete(allFields, "ipAddresses")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWitness() *Witness {
	p := new(Witness)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.config.Witness"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

type OneOfUpdateRecoveryPointExpirationTimeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRecoveryPointExpirationTimeApiResponseData"))
}

func (p *OneOfUpdateRecoveryPointExpirationTimeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRecoveryPointExpirationTimeApiResponseData")
}

type OneOfListConsistencyGroupsApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType400  *import3.ErrorResponse       `json:"-"`
	oneOfType2001 []ConsistencyGroup           `json:"-"`
	oneOfType401  []ConsistencyGroupProjection `json:"-"`
}

func NewOneOfListConsistencyGroupsApiResponseData() *OneOfListConsistencyGroupsApiResponseData {
	p := new(OneOfListConsistencyGroupsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListConsistencyGroupsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListConsistencyGroupsApiResponseData is nil"))
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
	case []ConsistencyGroup:
		p.oneOfType2001 = v.([]ConsistencyGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.ConsistencyGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.ConsistencyGroup>"
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

func (p *OneOfListConsistencyGroupsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.ConsistencyGroup>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if "List<dataprotection.v4.config.ConsistencyGroupProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListConsistencyGroupsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]ConsistencyGroup)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.ConsistencyGroup" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListConsistencyGroupsApiResponseData"))
}

func (p *OneOfListConsistencyGroupsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.ConsistencyGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if "List<dataprotection.v4.config.ConsistencyGroupProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListConsistencyGroupsApiResponseData")
}

type OneOfRecoveryPointRepositoryObjectStorageReference struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2103 *NutanixObjectsBucket      `json:"-"`
	oneOfType2102 *AmazonS3Bucket            `json:"-"`
	oneOfType2101 *AzureBlobStorageContainer `json:"-"`
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
	case NutanixObjectsBucket:
		if nil == p.oneOfType2103 {
			p.oneOfType2103 = new(NutanixObjectsBucket)
		}
		*p.oneOfType2103 = v.(NutanixObjectsBucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2103.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2103.ObjectType_
	case AmazonS3Bucket:
		if nil == p.oneOfType2102 {
			p.oneOfType2102 = new(AmazonS3Bucket)
		}
		*p.oneOfType2102 = v.(AmazonS3Bucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2102.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2102.ObjectType_
	case AzureBlobStorageContainer:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(AzureBlobStorageContainer)
		}
		*p.oneOfType2101 = v.(AzureBlobStorageContainer)
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

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) GetValue() interface{} {
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2103
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2102
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) UnmarshalJSON(b []byte) error {
	vOneOfType2103 := new(NutanixObjectsBucket)
	if err := json.Unmarshal(b, vOneOfType2103); err == nil {
		if "dataprotection.v4.config.NutanixObjectsBucket" == *vOneOfType2103.ObjectType_ {
			if nil == p.oneOfType2103 {
				p.oneOfType2103 = new(NutanixObjectsBucket)
			}
			*p.oneOfType2103 = *vOneOfType2103
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2103.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2103.ObjectType_
			return nil
		}
	}
	vOneOfType2102 := new(AmazonS3Bucket)
	if err := json.Unmarshal(b, vOneOfType2102); err == nil {
		if "dataprotection.v4.config.AmazonS3Bucket" == *vOneOfType2102.ObjectType_ {
			if nil == p.oneOfType2102 {
				p.oneOfType2102 = new(AmazonS3Bucket)
			}
			*p.oneOfType2102 = *vOneOfType2102
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2102.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2102.ObjectType_
			return nil
		}
	}
	vOneOfType2101 := new(AzureBlobStorageContainer)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "dataprotection.v4.config.AzureBlobStorageContainer" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(AzureBlobStorageContainer)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRepositoryObjectStorageReference"))
}

func (p *OneOfRecoveryPointRepositoryObjectStorageReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2103)
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2102)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRepositoryObjectStorageReference")
}

type OneOfListRecoveryPlanJobsApiResponseData struct {
	Discriminator *string                     `json:"-"`
	ObjectType_   *string                     `json:"-"`
	oneOfType400  *import3.ErrorResponse      `json:"-"`
	oneOfType2001 []RecoveryPlanJob           `json:"-"`
	oneOfType401  []RecoveryPlanJobProjection `json:"-"`
}

func NewOneOfListRecoveryPlanJobsApiResponseData() *OneOfListRecoveryPlanJobsApiResponseData {
	p := new(OneOfListRecoveryPlanJobsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecoveryPlanJobsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecoveryPlanJobsApiResponseData is nil"))
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
	case []RecoveryPlanJob:
		p.oneOfType2001 = v.([]RecoveryPlanJob)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJob>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJob>"
	case []RecoveryPlanJobProjection:
		p.oneOfType401 = v.([]RecoveryPlanJobProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.RecoveryPlanJob>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]RecoveryPlanJob)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.RecoveryPlanJob" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJob>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJob>"
			return nil
		}
	}
	vOneOfType401 := new([]RecoveryPlanJobProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.RecoveryPlanJobProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecoveryPlanJobsApiResponseData"))
}

func (p *OneOfListRecoveryPlanJobsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.RecoveryPlanJob>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRecoveryPlanJobsApiResponseData")
}

type OneOfGetRecoveryPlanJobApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *RecoveryPlanJob       `json:"-"`
}

func NewOneOfGetRecoveryPlanJobApiResponseData() *OneOfGetRecoveryPlanJobApiResponseData {
	p := new(OneOfGetRecoveryPlanJobApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRecoveryPlanJobApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRecoveryPlanJobApiResponseData is nil"))
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
	case RecoveryPlanJob:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(RecoveryPlanJob)
		}
		*p.oneOfType2001 = v.(RecoveryPlanJob)
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

func (p *OneOfGetRecoveryPlanJobApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetRecoveryPlanJobApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(RecoveryPlanJob)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.RecoveryPlanJob" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(RecoveryPlanJob)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecoveryPlanJobApiResponseData"))
}

func (p *OneOfGetRecoveryPlanJobApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecoveryPlanJobApiResponseData")
}

type OneOfDeleteRecoveryPlanJobApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
}

func NewOneOfDeleteRecoveryPlanJobApiResponseData() *OneOfDeleteRecoveryPlanJobApiResponseData {
	p := new(OneOfDeleteRecoveryPlanJobApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRecoveryPlanJobApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRecoveryPlanJobApiResponseData is nil"))
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

func (p *OneOfDeleteRecoveryPlanJobApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDeleteRecoveryPlanJobApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRecoveryPlanJobApiResponseData"))
}

func (p *OneOfDeleteRecoveryPlanJobApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRecoveryPlanJobApiResponseData")
}

type OneOfListRecoveryPlanJobExecutionStepsApiResponseData struct {
	Discriminator *string                                  `json:"-"`
	ObjectType_   *string                                  `json:"-"`
	oneOfType400  *import3.ErrorResponse                   `json:"-"`
	oneOfType401  []RecoveryPlanJobExecutionStepProjection `json:"-"`
	oneOfType2001 []RecoveryPlanJobExecutionStep           `json:"-"`
}

func NewOneOfListRecoveryPlanJobExecutionStepsApiResponseData() *OneOfListRecoveryPlanJobExecutionStepsApiResponseData {
	p := new(OneOfListRecoveryPlanJobExecutionStepsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecoveryPlanJobExecutionStepsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecoveryPlanJobExecutionStepsApiResponseData is nil"))
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
	case []RecoveryPlanJobExecutionStepProjection:
		p.oneOfType401 = v.([]RecoveryPlanJobExecutionStepProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>"
	case []RecoveryPlanJobExecutionStep:
		p.oneOfType2001 = v.([]RecoveryPlanJobExecutionStep)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobExecutionStepsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobExecutionStepsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]RecoveryPlanJobExecutionStepProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>"
			return nil
		}
	}
	vOneOfType2001 := new([]RecoveryPlanJobExecutionStep)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.RecoveryPlanJobExecutionStep" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecoveryPlanJobExecutionStepsApiResponseData"))
}

func (p *OneOfListRecoveryPlanJobExecutionStepsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobExecutionStepProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<dataprotection.v4.config.RecoveryPlanJobExecutionStep>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListRecoveryPlanJobExecutionStepsApiResponseData")
}

type OneOfProtectedResourceRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfProtectedResourceRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfProtectedResourceRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourceRestoreApiResponseData"))
}

func (p *OneOfProtectedResourceRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourceRestoreApiResponseData")
}

type OneOfRecoveryPointRepositoryProjectionObjectStorageReference struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2103 *NutanixObjectsBucket      `json:"-"`
	oneOfType2102 *AmazonS3Bucket            `json:"-"`
	oneOfType2101 *AzureBlobStorageContainer `json:"-"`
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
	case NutanixObjectsBucket:
		if nil == p.oneOfType2103 {
			p.oneOfType2103 = new(NutanixObjectsBucket)
		}
		*p.oneOfType2103 = v.(NutanixObjectsBucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2103.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2103.ObjectType_
	case AmazonS3Bucket:
		if nil == p.oneOfType2102 {
			p.oneOfType2102 = new(AmazonS3Bucket)
		}
		*p.oneOfType2102 = v.(AmazonS3Bucket)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2102.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2102.ObjectType_
	case AzureBlobStorageContainer:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(AzureBlobStorageContainer)
		}
		*p.oneOfType2101 = v.(AzureBlobStorageContainer)
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

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) GetValue() interface{} {
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2103
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2102
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) UnmarshalJSON(b []byte) error {
	vOneOfType2103 := new(NutanixObjectsBucket)
	if err := json.Unmarshal(b, vOneOfType2103); err == nil {
		if "dataprotection.v4.config.NutanixObjectsBucket" == *vOneOfType2103.ObjectType_ {
			if nil == p.oneOfType2103 {
				p.oneOfType2103 = new(NutanixObjectsBucket)
			}
			*p.oneOfType2103 = *vOneOfType2103
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2103.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2103.ObjectType_
			return nil
		}
	}
	vOneOfType2102 := new(AmazonS3Bucket)
	if err := json.Unmarshal(b, vOneOfType2102); err == nil {
		if "dataprotection.v4.config.AmazonS3Bucket" == *vOneOfType2102.ObjectType_ {
			if nil == p.oneOfType2102 {
				p.oneOfType2102 = new(AmazonS3Bucket)
			}
			*p.oneOfType2102 = *vOneOfType2102
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2102.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2102.ObjectType_
			return nil
		}
	}
	vOneOfType2101 := new(AzureBlobStorageContainer)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "dataprotection.v4.config.AzureBlobStorageContainer" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(AzureBlobStorageContainer)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRepositoryProjectionObjectStorageReference"))
}

func (p *OneOfRecoveryPointRepositoryProjectionObjectStorageReference) MarshalJSON() ([]byte, error) {
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2103)
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2102)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRepositoryProjectionObjectStorageReference")
}

type OneOfExecutionStepResultResult struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType2101 *EntityRecoveryResult `json:"-"`
}

func NewOneOfExecutionStepResultResult() *OneOfExecutionStepResultResult {
	p := new(OneOfExecutionStepResultResult)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfExecutionStepResultResult) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfExecutionStepResultResult is nil"))
	}
	switch v.(type) {
	case EntityRecoveryResult:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(EntityRecoveryResult)
		}
		*p.oneOfType2101 = v.(EntityRecoveryResult)
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

func (p *OneOfExecutionStepResultResult) GetValue() interface{} {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfExecutionStepResultResult) UnmarshalJSON(b []byte) error {
	vOneOfType2101 := new(EntityRecoveryResult)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "dataprotection.v4.config.EntityRecoveryResult" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(EntityRecoveryResult)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfExecutionStepResultResult"))
}

func (p *OneOfExecutionStepResultResult) MarshalJSON() ([]byte, error) {
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfExecutionStepResultResult")
}

type OneOfDeleteConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType1002 *interface{}           `json:"-"`
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
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(interface{})
		}
		*p.oneOfType1002 = nil
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

func (p *OneOfDeleteConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1002
	}
	return nil
}

func (p *OneOfDeleteConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1002 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if nil == *vOneOfType1002 {
			if nil == p.oneOfType1002 {
				p.oneOfType1002 = new(interface{})
			}
			*p.oneOfType1002 = nil
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteConsistencyGroupApiResponseData"))
}

func (p *OneOfDeleteConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteConsistencyGroupApiResponseData")
}

type OneOfListRecoveryPointsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType400  *import3.ErrorResponse    `json:"-"`
	oneOfType401  []RecoveryPointProjection `json:"-"`
	oneOfType2001 []RecoveryPoint           `json:"-"`
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
	case []RecoveryPoint:
		p.oneOfType2001 = v.([]RecoveryPoint)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPoint>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPoint>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRecoveryPointsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.RecoveryPointProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListRecoveryPointsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.RecoveryPoint" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecoveryPointsApiResponseData"))
}

func (p *OneOfListRecoveryPointsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.RecoveryPointProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<dataprotection.v4.config.RecoveryPoint>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListRecoveryPointsApiResponseData")
}

type OneOfVmRecoveryPointApplicationConsistentProperties struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.VssProperties `json:"-"`
}

func NewOneOfVmRecoveryPointApplicationConsistentProperties() *OneOfVmRecoveryPointApplicationConsistentProperties {
	p := new(OneOfVmRecoveryPointApplicationConsistentProperties)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRecoveryPointApplicationConsistentProperties is nil"))
	}
	switch v.(type) {
	case import2.VssProperties:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.VssProperties)
		}
		*p.oneOfType2001 = v.(import2.VssProperties)
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

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.VssProperties)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.common.VssProperties" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.VssProperties)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRecoveryPointApplicationConsistentProperties"))
}

func (p *OneOfVmRecoveryPointApplicationConsistentProperties) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfVmRecoveryPointApplicationConsistentProperties")
}

type OneOfGetConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *ConsistencyGroup      `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetConsistencyGroupApiResponseData() *OneOfGetConsistencyGroupApiResponseData {
	p := new(OneOfGetConsistencyGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetConsistencyGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetConsistencyGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case ConsistencyGroup:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ConsistencyGroup)
		}
		*p.oneOfType2001 = v.(ConsistencyGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfGetConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(ConsistencyGroup)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.ConsistencyGroup" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ConsistencyGroup)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetConsistencyGroupApiResponseData"))
}

func (p *OneOfGetConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetConsistencyGroupApiResponseData")
}

type OneOfCreateRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfCreateRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCreateRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRecoveryPointApiResponseData"))
}

func (p *OneOfCreateRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRecoveryPointApiResponseData")
}

type OneOfRecoveryPointRestoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfRecoveryPointRestoreApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfRecoveryPointRestoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointRestoreApiResponseData"))
}

func (p *OneOfRecoveryPointRestoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointRestoreApiResponseData")
}

type OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec struct {
	Discriminator *string             `json:"-"`
	ObjectType_   *string             `json:"-"`
	oneOfType2102 *EsxiVmOverrideSpec `json:"-"`
	oneOfType2101 *AhvVmOverrideSpec  `json:"-"`
}

func NewOneOfVmRecoveryPointRestoreOverrideVmOverrideSpec() *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec {
	p := new(OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec is nil"))
	}
	switch v.(type) {
	case EsxiVmOverrideSpec:
		if nil == p.oneOfType2102 {
			p.oneOfType2102 = new(EsxiVmOverrideSpec)
		}
		*p.oneOfType2102 = v.(EsxiVmOverrideSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2102.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2102.ObjectType_
	case AhvVmOverrideSpec:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(AhvVmOverrideSpec)
		}
		*p.oneOfType2101 = v.(AhvVmOverrideSpec)
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

func (p *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec) GetValue() interface{} {
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2102
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec) UnmarshalJSON(b []byte) error {
	vOneOfType2102 := new(EsxiVmOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType2102); err == nil {
		if "dataprotection.v4.config.EsxiVmOverrideSpec" == *vOneOfType2102.ObjectType_ {
			if nil == p.oneOfType2102 {
				p.oneOfType2102 = new(EsxiVmOverrideSpec)
			}
			*p.oneOfType2102 = *vOneOfType2102
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2102.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2102.ObjectType_
			return nil
		}
	}
	vOneOfType2101 := new(AhvVmOverrideSpec)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if "dataprotection.v4.config.AhvVmOverrideSpec" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(AhvVmOverrideSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec"))
}

func (p *OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2102)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfVmRecoveryPointRestoreOverrideVmOverrideSpec")
}

type OneOfCreateConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *ConsistencyGroup      `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
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
	case ConsistencyGroup:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ConsistencyGroup)
		}
		*p.oneOfType2001 = v.(ConsistencyGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfCreateConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(ConsistencyGroup)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.ConsistencyGroup" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ConsistencyGroup)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateConsistencyGroupApiResponseData"))
}

func (p *OneOfCreateConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateConsistencyGroupApiResponseData")
}

type OneOfGetRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *RecoveryPoint         `json:"-"`
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
	case RecoveryPoint:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(RecoveryPoint)
		}
		*p.oneOfType2001 = v.(RecoveryPoint)
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

func (p *OneOfGetRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(RecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.RecoveryPoint" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(RecoveryPoint)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecoveryPointApiResponseData"))
}

func (p *OneOfGetRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecoveryPointApiResponseData")
}

type OneOfDeleteRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfDeleteRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDeleteRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRecoveryPointApiResponseData"))
}

func (p *OneOfDeleteRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRecoveryPointApiResponseData")
}

type OneOfGetProtectedResourceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *ProtectedResource     `json:"-"`
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
	case ProtectedResource:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ProtectedResource)
		}
		*p.oneOfType2001 = v.(ProtectedResource)
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

func (p *OneOfGetProtectedResourceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetProtectedResourceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(ProtectedResource)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.ProtectedResource" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ProtectedResource)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetProtectedResourceApiResponseData"))
}

func (p *OneOfGetProtectedResourceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetProtectedResourceApiResponseData")
}

type OneOfListDPClusterCapabilitiesApiResponseData struct {
	Discriminator *string                                     `json:"-"`
	ObjectType_   *string                                     `json:"-"`
	oneOfType400  *import3.ErrorResponse                      `json:"-"`
	oneOfType401  []DataProtectionClusterCapabilityProjection `json:"-"`
	oneOfType2001 []DataProtectionClusterCapability           `json:"-"`
}

func NewOneOfListDPClusterCapabilitiesApiResponseData() *OneOfListDPClusterCapabilitiesApiResponseData {
	p := new(OneOfListDPClusterCapabilitiesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDPClusterCapabilitiesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDPClusterCapabilitiesApiResponseData is nil"))
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
	case []DataProtectionClusterCapabilityProjection:
		p.oneOfType401 = v.([]DataProtectionClusterCapabilityProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>"
	case []DataProtectionClusterCapability:
		p.oneOfType2001 = v.([]DataProtectionClusterCapability)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.DataProtectionClusterCapability>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.DataProtectionClusterCapability>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListDPClusterCapabilitiesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<dataprotection.v4.config.DataProtectionClusterCapability>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListDPClusterCapabilitiesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]DataProtectionClusterCapabilityProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.DataProtectionClusterCapabilityProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>"
			return nil
		}
	}
	vOneOfType2001 := new([]DataProtectionClusterCapability)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.DataProtectionClusterCapability" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.DataProtectionClusterCapability>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.DataProtectionClusterCapability>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDPClusterCapabilitiesApiResponseData"))
}

func (p *OneOfListDPClusterCapabilitiesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.DataProtectionClusterCapabilityProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<dataprotection.v4.config.DataProtectionClusterCapability>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListDPClusterCapabilitiesApiResponseData")
}

type OneOfListRecoveryPlanJobValidationErrorsApiResponseData struct {
	Discriminator *string                                 `json:"-"`
	ObjectType_   *string                                 `json:"-"`
	oneOfType400  *import3.ErrorResponse                  `json:"-"`
	oneOfType2001 []RecoveryPlanValidationError           `json:"-"`
	oneOfType401  []RecoveryPlanValidationErrorProjection `json:"-"`
}

func NewOneOfListRecoveryPlanJobValidationErrorsApiResponseData() *OneOfListRecoveryPlanJobValidationErrorsApiResponseData {
	p := new(OneOfListRecoveryPlanJobValidationErrorsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRecoveryPlanJobValidationErrorsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRecoveryPlanJobValidationErrorsApiResponseData is nil"))
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
	case []RecoveryPlanValidationError:
		p.oneOfType2001 = v.([]RecoveryPlanValidationError)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanValidationError>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanValidationError>"
	case []RecoveryPlanValidationErrorProjection:
		p.oneOfType401 = v.([]RecoveryPlanValidationErrorProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobValidationErrorsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.config.RecoveryPlanValidationError>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRecoveryPlanJobValidationErrorsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]RecoveryPlanValidationError)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.config.RecoveryPlanValidationError" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanValidationError>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanValidationError>"
			return nil
		}
	}
	vOneOfType401 := new([]RecoveryPlanValidationErrorProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "dataprotection.v4.config.RecoveryPlanValidationErrorProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRecoveryPlanJobValidationErrorsApiResponseData"))
}

func (p *OneOfListRecoveryPlanJobValidationErrorsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.config.RecoveryPlanValidationError>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if "List<dataprotection.v4.config.RecoveryPlanValidationErrorProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRecoveryPlanJobValidationErrorsApiResponseData")
}

type OneOfProtectedResourcePromoteApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfProtectedResourcePromoteApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfProtectedResourcePromoteApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfProtectedResourcePromoteApiResponseData"))
}

func (p *OneOfProtectedResourcePromoteApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfProtectedResourcePromoteApiResponseData")
}

type OneOfUpdateConsistencyGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []import3.AppMessage   `json:"-"`
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
	case []import3.AppMessage:
		p.oneOfType2001 = v.([]import3.AppMessage)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.error.AppMessage>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.error.AppMessage>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]import3.AppMessage)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "dataprotection.v4.error.AppMessage" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateConsistencyGroupApiResponseData"))
}

func (p *OneOfUpdateConsistencyGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<dataprotection.v4.error.AppMessage>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateConsistencyGroupApiResponseData")
}

type OneOfClusterInfoApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.ClusterInfo   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfClusterInfoApiResponseData() *OneOfClusterInfoApiResponseData {
	p := new(OneOfClusterInfoApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterInfoApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterInfoApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ClusterInfo:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.ClusterInfo)
		}
		*p.oneOfType2001 = v.(import2.ClusterInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfClusterInfoApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfClusterInfoApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.ClusterInfo)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.common.ClusterInfo" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.ClusterInfo)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterInfoApiResponseData"))
}

func (p *OneOfClusterInfoApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfClusterInfoApiResponseData")
}

type OneOfGetVmRecoveryPointApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *VmRecoveryPoint       `json:"-"`
}

func NewOneOfGetVmRecoveryPointApiResponseData() *OneOfGetVmRecoveryPointApiResponseData {
	p := new(OneOfGetVmRecoveryPointApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmRecoveryPointApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmRecoveryPointApiResponseData is nil"))
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
	case VmRecoveryPoint:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmRecoveryPoint)
		}
		*p.oneOfType2001 = v.(VmRecoveryPoint)
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

func (p *OneOfGetVmRecoveryPointApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetVmRecoveryPointApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(VmRecoveryPoint)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "dataprotection.v4.config.VmRecoveryPoint" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmRecoveryPoint)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmRecoveryPointApiResponseData"))
}

func (p *OneOfGetVmRecoveryPointApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmRecoveryPointApiResponseData")
}

type OneOfRecoveryPointReplicateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import5.TaskReference `json:"-"`
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

func (p *OneOfRecoveryPointReplicateApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfRecoveryPointReplicateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoveryPointReplicateApiResponseData"))
}

func (p *OneOfRecoveryPointReplicateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfRecoveryPointReplicateApiResponseData")
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
