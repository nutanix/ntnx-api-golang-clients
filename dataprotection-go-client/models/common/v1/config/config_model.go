/*
 * Generated file models/common/v1/config/config_model.go.
 *
 * Product version: 4.4.1
 *
 * Part of the Nutanix Data Protection APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Nutanix Standard Configuration
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Reference to the Amazon S3 bucket.
*/
type AmazonS3Bucket struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AuthCredentialsItemDiscriminator_ *string `json:"$authCredentialsItemDiscriminator,omitempty"`
	/*
	  Authentication credentials for accessing the object store.
	*/
	AuthCredentials *OneOfAmazonS3BucketAuthCredentials `json:"authCredentials"`
	/*
	  Name of the Amazon S3 bucket.
	*/
	BucketName *string `json:"bucketName"`
	/*
	  AWS region.
	*/
	Region *string `json:"region"`
}

func (p *AmazonS3Bucket) MarshalJSON() ([]byte, error) {
	type AmazonS3BucketProxy AmazonS3Bucket

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AmazonS3BucketProxy
		AuthCredentials *OneOfAmazonS3BucketAuthCredentials `json:"authCredentials,omitempty"`
		BucketName      *string                             `json:"bucketName,omitempty"`
		Region          *string                             `json:"region,omitempty"`
	}{
		AmazonS3BucketProxy: (*AmazonS3BucketProxy)(p),
		AuthCredentials:     p.AuthCredentials,
		BucketName:          p.BucketName,
		Region:              p.Region,
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
	if known.AuthCredentialsItemDiscriminator_ != nil {
		p.AuthCredentialsItemDiscriminator_ = known.AuthCredentialsItemDiscriminator_
	}
	if known.AuthCredentials != nil {
		p.AuthCredentials = known.AuthCredentials
	}
	if known.BucketName != nil {
		p.BucketName = known.BucketName
	}
	if known.Region != nil {
		p.Region = known.Region
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$authCredentialsItemDiscriminator")
	delete(allFields, "authCredentials")
	delete(allFields, "bucketName")
	delete(allFields, "region")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAmazonS3Bucket() *AmazonS3Bucket {
	p := new(AmazonS3Bucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.AmazonS3Bucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AmazonS3Bucket) GetAuthCredentials() interface{} {
	if nil == p.AuthCredentials {
		return nil
	}
	return p.AuthCredentials.GetValue()
}

func (p *AmazonS3Bucket) SetAuthCredentials(v interface{}) error {
	if nil == p.AuthCredentials {
		p.AuthCredentials = NewOneOfAmazonS3BucketAuthCredentials()
	}
	e := p.AuthCredentials.SetValue(v)
	if nil == e {
		if nil == p.AuthCredentialsItemDiscriminator_ {
			p.AuthCredentialsItemDiscriminator_ = new(string)
		}
		*p.AuthCredentialsItemDiscriminator_ = *p.AuthCredentials.Discriminator
	}
	return e
}

/*
Reference to the Microsoft Azure Blob Storage container.
*/
type AzureBlobStorageContainer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AuthCredentialsItemDiscriminator_ *string `json:"$authCredentialsItemDiscriminator,omitempty"`
	/*
	  Authentication credentials for accessing the object store.
	*/
	AuthCredentials *OneOfAzureBlobStorageContainerAuthCredentials `json:"authCredentials"`
	/*
	  Name of the Microsoft Azure Blob Storage container.
	*/
	ContainerName *string `json:"containerName"`
	/*
	  Name of the Microsoft Azure Blob Storage account.
	*/
	StorageAccountName *string `json:"storageAccountName"`
}

func (p *AzureBlobStorageContainer) MarshalJSON() ([]byte, error) {
	type AzureBlobStorageContainerProxy AzureBlobStorageContainer

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AzureBlobStorageContainerProxy
		AuthCredentials    *OneOfAzureBlobStorageContainerAuthCredentials `json:"authCredentials,omitempty"`
		ContainerName      *string                                        `json:"containerName,omitempty"`
		StorageAccountName *string                                        `json:"storageAccountName,omitempty"`
	}{
		AzureBlobStorageContainerProxy: (*AzureBlobStorageContainerProxy)(p),
		AuthCredentials:                p.AuthCredentials,
		ContainerName:                  p.ContainerName,
		StorageAccountName:             p.StorageAccountName,
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
	if known.AuthCredentialsItemDiscriminator_ != nil {
		p.AuthCredentialsItemDiscriminator_ = known.AuthCredentialsItemDiscriminator_
	}
	if known.AuthCredentials != nil {
		p.AuthCredentials = known.AuthCredentials
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
	delete(allFields, "$authCredentialsItemDiscriminator")
	delete(allFields, "authCredentials")
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
	*p.ObjectType_ = "common.v1.config.AzureBlobStorageContainer"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AzureBlobStorageContainer) GetAuthCredentials() interface{} {
	if nil == p.AuthCredentials {
		return nil
	}
	return p.AuthCredentials.GetValue()
}

func (p *AzureBlobStorageContainer) SetAuthCredentials(v interface{}) error {
	if nil == p.AuthCredentials {
		p.AuthCredentials = NewOneOfAzureBlobStorageContainerAuthCredentials()
	}
	e := p.AuthCredentials.SetValue(v)
	if nil == e {
		if nil == p.AuthCredentialsItemDiscriminator_ {
			p.AuthCredentialsItemDiscriminator_ = new(string)
		}
		*p.AuthCredentialsItemDiscriminator_ = *p.AuthCredentials.Discriminator
	}
	return e
}

/*
Azure shared key credentials for storage account authentication.
*/
type AzureSharedKeyCredentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Azure storage account key for authentication.
	*/
	AccountKey *string `json:"accountKey"`
}

func (p *AzureSharedKeyCredentials) MarshalJSON() ([]byte, error) {
	type AzureSharedKeyCredentialsProxy AzureSharedKeyCredentials

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AzureSharedKeyCredentialsProxy
		AccountKey *string `json:"accountKey,omitempty"`
	}{
		AzureSharedKeyCredentialsProxy: (*AzureSharedKeyCredentialsProxy)(p),
		AccountKey:                     p.AccountKey,
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

func (p *AzureSharedKeyCredentials) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AzureSharedKeyCredentials
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAzureSharedKeyCredentials()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccountKey != nil {
		p.AccountKey = known.AccountKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accountKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAzureSharedKeyCredentials() *AzureSharedKeyCredentials {
	p := new(AzureSharedKeyCredentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.AzureSharedKeyCredentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Represents another entity that has been referenced by this entity.
*/
type EntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *EntityType `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Name of the entity represented by this reference.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  URI of entity represented by this reference.
	*/
	Uris []string `json:"uris,omitempty"`
}

func (p *EntityReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityReference

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
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Uris != nil {
		p.Uris = known.Uris
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "name")
	delete(allFields, "uris")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of entity represented by this reference, e.g., VPC.
*/
type EntityType int

const (
	ENTITYTYPE_UNKNOWN             EntityType = 0
	ENTITYTYPE_REDACTED            EntityType = 1
	ENTITYTYPE_CLUSTER             EntityType = 2
	ENTITYTYPE_VM                  EntityType = 3
	ENTITYTYPE_STORAGE_CONTAINER   EntityType = 4
	ENTITYTYPE_VOLUME_GROUP        EntityType = 5
	ENTITYTYPE_TASK                EntityType = 6
	ENTITYTYPE_IMAGE               EntityType = 7
	ENTITYTYPE_CATEGORY            EntityType = 8
	ENTITYTYPE_NODE                EntityType = 9
	ENTITYTYPE_VPC                 EntityType = 10
	ENTITYTYPE_SUBNET              EntityType = 11
	ENTITYTYPE_ROUTING_POLICY      EntityType = 12
	ENTITYTYPE_FLOATING_IP         EntityType = 13
	ENTITYTYPE_VPN_GATEWAY         EntityType = 14
	ENTITYTYPE_VPN_CONNECTION      EntityType = 15
	ENTITYTYPE_DIRECT_CONNECT      EntityType = 16
	ENTITYTYPE_DIRECT_CONNECT_VIF  EntityType = 17
	ENTITYTYPE_VIRTUAL_NIC         EntityType = 18
	ENTITYTYPE_VIRTUAL_SWITCH      EntityType = 19
	ENTITYTYPE_VM_DISK             EntityType = 20
	ENTITYTYPE_VOLUME_DISK         EntityType = 21
	ENTITYTYPE_DISK_RECOVERY_POINT EntityType = 22
	ENTITYTYPE_VTEP_GATEWAY        EntityType = 23
	ENTITYTYPE_RECOVERY_PLAN       EntityType = 24
	ENTITYTYPE_RECOVERY_PLAN_JOB   EntityType = 25
	ENTITYTYPE_AVAILABILITY_ZONE   EntityType = 26
	ENTITYTYPE_VIRTUAL_NETWORK     EntityType = 27
	ENTITYTYPE_CONSISTENCY_GROUP   EntityType = 28
	ENTITYTYPE_SUBNET_EXTENSION    EntityType = 29
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"VM",
		"STORAGE_CONTAINER",
		"VOLUME_GROUP",
		"TASK",
		"IMAGE",
		"CATEGORY",
		"NODE",
		"VPC",
		"SUBNET",
		"ROUTING_POLICY",
		"FLOATING_IP",
		"VPN_GATEWAY",
		"VPN_CONNECTION",
		"DIRECT_CONNECT",
		"DIRECT_CONNECT_VIF",
		"VIRTUAL_NIC",
		"VIRTUAL_SWITCH",
		"VM_DISK",
		"VOLUME_DISK",
		"DISK_RECOVERY_POINT",
		"VTEP_GATEWAY",
		"RECOVERY_PLAN",
		"RECOVERY_PLAN_JOB",
		"AVAILABILITY_ZONE",
		"VIRTUAL_NETWORK",
		"CONSISTENCY_GROUP",
		"SUBNET_EXTENSION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"VM",
		"STORAGE_CONTAINER",
		"VOLUME_GROUP",
		"TASK",
		"IMAGE",
		"CATEGORY",
		"NODE",
		"VPC",
		"SUBNET",
		"ROUTING_POLICY",
		"FLOATING_IP",
		"VPN_GATEWAY",
		"VPN_CONNECTION",
		"DIRECT_CONNECT",
		"DIRECT_CONNECT_VIF",
		"VIRTUAL_NIC",
		"VIRTUAL_SWITCH",
		"VM_DISK",
		"VOLUME_DISK",
		"DISK_RECOVERY_POINT",
		"VTEP_GATEWAY",
		"RECOVERY_PLAN",
		"RECOVERY_PLAN_JOB",
		"AVAILABILITY_ZONE",
		"VIRTUAL_NETWORK",
		"CONSISTENCY_GROUP",
		"SUBNET_EXTENSION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EntityType) index(name string) EntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"VM",
		"STORAGE_CONTAINER",
		"VOLUME_GROUP",
		"TASK",
		"IMAGE",
		"CATEGORY",
		"NODE",
		"VPC",
		"SUBNET",
		"ROUTING_POLICY",
		"FLOATING_IP",
		"VPN_GATEWAY",
		"VPN_CONNECTION",
		"DIRECT_CONNECT",
		"DIRECT_CONNECT_VIF",
		"VIRTUAL_NIC",
		"VIRTUAL_SWITCH",
		"VM_DISK",
		"VOLUME_DISK",
		"DISK_RECOVERY_POINT",
		"VTEP_GATEWAY",
		"RECOVERY_PLAN",
		"RECOVERY_PLAN_JOB",
		"AVAILABILITY_ZONE",
		"VIRTUAL_NETWORK",
		"CONSISTENCY_GROUP",
		"SUBNET_EXTENSION",
	}
	for idx := range names {
		if names[idx] == name {
			return EntityType(idx)
		}
	}
	return ENTITYTYPE_UNKNOWN
}

func (e *EntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EntityType) Ref() *EntityType {
	return &e
}

/*
Many entities in the Nutanix APIs carry flags.  This object captures all the flags associated with that entity through this object.  The field that hosts this type of object must have an attribute called x-bounded-map-keys that tells which flags are actually present for that entity.
*/
type Flag struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the flag.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Value of the flag.
	*/
	Value *bool `json:"value,omitempty"`
}

func (p *Flag) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Flag

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

func (p *Flag) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Flag
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFlag()

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
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFlag() *Flag {
	p := new(Flag)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.Flag"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Value = new(bool)
	*p.Value = false

	return p
}

/*
An unique address that identifies a device on the internet or a local network in IPv4 or IPv6 format.
*/
type IPAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ipv4 *IPv4Address `json:"ipv4,omitempty"`

	Ipv6 *IPv6Address `json:"ipv6,omitempty"`
}

func (p *IPAddress) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IPAddress

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

func (p *IPAddress) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPAddress
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIPAddress()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Ipv4 != nil {
		p.Ipv4 = known.Ipv4
	}
	if known.Ipv6 != nil {
		p.Ipv6 = known.Ipv6
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipv4")
	delete(allFields, "ipv6")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIPAddress() *IPAddress {
	p := new(IPAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.IPAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (i *IPAddress) HasIpv4() bool {
	return i.Ipv4 != nil
}
func (i *IPAddress) HasIpv6() bool {
	return i.Ipv6 != nil
}

func (i *IPAddress) IsValid() bool {
	return i.HasIpv4() || i.HasIpv6()
}

/*
An unique address that identifies a device on the internet or a local network in IPv4 format.
*/
type IPv4Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The prefix length of the network to which this host IPv4 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  The IPv4 address of the host.
	*/
	Value *string `json:"value"`
}

func (p *IPv4Address) MarshalJSON() ([]byte, error) {
	type IPv4AddressProxy IPv4Address

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IPv4AddressProxy
		Value *string `json:"value,omitempty"`
	}{
		IPv4AddressProxy: (*IPv4AddressProxy)(p),
		Value:            p.Value,
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

func (p *IPv4Address) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPv4Address
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIPv4Address()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PrefixLength != nil {
		p.PrefixLength = known.PrefixLength
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "prefixLength")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIPv4Address() *IPv4Address {
	p := new(IPv4Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.IPv4Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 32

	return p
}

/*
An unique address that identifies a device on the internet or a local network in IPv6 format.
*/
type IPv6Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The prefix length of the network to which this host IPv6 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  The IPv6 address of the host.
	*/
	Value *string `json:"value"`
}

func (p *IPv6Address) MarshalJSON() ([]byte, error) {
	type IPv6AddressProxy IPv6Address

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IPv6AddressProxy
		Value *string `json:"value,omitempty"`
	}{
		IPv6AddressProxy: (*IPv6AddressProxy)(p),
		Value:            p.Value,
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

func (p *IPv6Address) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPv6Address
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIPv6Address()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.PrefixLength != nil {
		p.PrefixLength = known.PrefixLength
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "prefixLength")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIPv6Address() *IPv6Address {
	p := new(IPv6Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.IPv6Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 128

	return p
}

/*
A map describing a set of keys and their corresponding values.
*/
type KVPair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The key of the key-value pair.
	*/
	Name *string `json:"name,omitempty"`
	/*

	 */
	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  The value associated with the key for this key-value pair
	*/
	Value *OneOfKVPairValue `json:"value,omitempty"`
}

func (p *KVPair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias KVPair

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

func (p *KVPair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KVPair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKVPair()

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
	if known.ValueItemDiscriminator_ != nil {
		p.ValueItemDiscriminator_ = known.ValueItemDiscriminator_
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewKVPair() *KVPair {
	p := new(KVPair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.KVPair"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *KVPair) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *KVPair) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfKVPairValue()
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
A wrapper schema containing a map with string keys and values.
*/
type MapOfStringWrapper struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A map with string keys and values.
	*/
	Map map[string]string `json:"map,omitempty"`
}

func (p *MapOfStringWrapper) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MapOfStringWrapper

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

func (p *MapOfStringWrapper) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MapOfStringWrapper
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMapOfStringWrapper()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Map != nil {
		p.Map = known.Map
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "map")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMapOfStringWrapper() *MapOfStringWrapper {
	p := new(MapOfStringWrapper)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.MapOfStringWrapper"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Message struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A code that uniquely identifies a message.
	*/
	Code *string `json:"code,omitempty"`
	/*
	  The locale for the message description.
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  The description of the message.
	*/
	Message *string `json:"message,omitempty"`

	Severity *MessageSeverity `json:"severity,omitempty"`
}

func (p *Message) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Message

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

func (p *Message) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Message
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMessage()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Code != nil {
		p.Code = known.Code
	}
	if known.Locale != nil {
		p.Locale = known.Locale
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.Severity != nil {
		p.Severity = known.Severity
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "code")
	delete(allFields, "locale")
	delete(allFields, "message")
	delete(allFields, "severity")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMessage() *Message {
	p := new(Message)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.Message"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Locale = new(string)
	*p.Locale = "en_US"

	return p
}

/*
The message severity.
*/
type MessageSeverity int

const (
	MESSAGESEVERITY_UNKNOWN  MessageSeverity = 0
	MESSAGESEVERITY_REDACTED MessageSeverity = 1
	MESSAGESEVERITY_INFO     MessageSeverity = 2
	MESSAGESEVERITY_WARNING  MessageSeverity = 3
	MESSAGESEVERITY_ERROR    MessageSeverity = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MessageSeverity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MessageSeverity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MessageSeverity) index(name string) MessageSeverity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return MessageSeverity(idx)
		}
	}
	return MESSAGESEVERITY_UNKNOWN
}

func (e *MessageSeverity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MessageSeverity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MessageSeverity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MessageSeverity) Ref() *MessageSeverity {
	return &e
}

/*
Reference to the Nutanix Objects bucket.
*/
type NutanixObjectsBucket struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AuthCredentialsItemDiscriminator_ *string `json:"$authCredentialsItemDiscriminator,omitempty"`
	/*
	  Authentication credentials for accessing the object store.
	*/
	AuthCredentials *OneOfNutanixObjectsBucketAuthCredentials `json:"authCredentials"`
	/*
	  Nutanix Objects bucket name.
	*/
	BucketName *string `json:"bucketName"`
	/*
	  Complete endpoint URL of the Nutanix Objects account including protocol (http/https) and hostname or IP address. Example: https://10.1.2.100
	*/
	EndPoint *string `json:"endPoint"`
}

func (p *NutanixObjectsBucket) MarshalJSON() ([]byte, error) {
	type NutanixObjectsBucketProxy NutanixObjectsBucket

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NutanixObjectsBucketProxy
		AuthCredentials *OneOfNutanixObjectsBucketAuthCredentials `json:"authCredentials,omitempty"`
		BucketName      *string                                   `json:"bucketName,omitempty"`
		EndPoint        *string                                   `json:"endPoint,omitempty"`
	}{
		NutanixObjectsBucketProxy: (*NutanixObjectsBucketProxy)(p),
		AuthCredentials:           p.AuthCredentials,
		BucketName:                p.BucketName,
		EndPoint:                  p.EndPoint,
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
	if known.AuthCredentialsItemDiscriminator_ != nil {
		p.AuthCredentialsItemDiscriminator_ = known.AuthCredentialsItemDiscriminator_
	}
	if known.AuthCredentials != nil {
		p.AuthCredentials = known.AuthCredentials
	}
	if known.BucketName != nil {
		p.BucketName = known.BucketName
	}
	if known.EndPoint != nil {
		p.EndPoint = known.EndPoint
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$authCredentialsItemDiscriminator")
	delete(allFields, "authCredentials")
	delete(allFields, "bucketName")
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
	*p.ObjectType_ = "common.v1.config.NutanixObjectsBucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *NutanixObjectsBucket) GetAuthCredentials() interface{} {
	if nil == p.AuthCredentials {
		return nil
	}
	return p.AuthCredentials.GetValue()
}

func (p *NutanixObjectsBucket) SetAuthCredentials(v interface{}) error {
	if nil == p.AuthCredentials {
		p.AuthCredentials = NewOneOfNutanixObjectsBucketAuthCredentials()
	}
	e := p.AuthCredentials.SetValue(v)
	if nil == e {
		if nil == p.AuthCredentialsItemDiscriminator_ {
			p.AuthCredentialsItemDiscriminator_ = new(string)
		}
		*p.AuthCredentialsItemDiscriminator_ = *p.AuthCredentials.Discriminator
	}
	return e
}

/*
Object store bucket.
*/
type ObjectStoreBucket struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the object store bucket. In case of Azure blob store, this refers to the container name.
	*/
	BucketName *string `json:"bucketName,omitempty"`

	Provider *ObjectStoreProvider `json:"provider,omitempty"`
	/*

	 */
	ProviderConfigItemDiscriminator_ *string `json:"$providerConfigItemDiscriminator,omitempty"`
	/*
	  Configuration details specific to the selected object store provider.
	*/
	ProviderConfig *OneOfObjectStoreBucketProviderConfig `json:"providerConfig"`
}

func (p *ObjectStoreBucket) MarshalJSON() ([]byte, error) {
	type ObjectStoreBucketProxy ObjectStoreBucket

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ObjectStoreBucketProxy
		ProviderConfig *OneOfObjectStoreBucketProviderConfig `json:"providerConfig,omitempty"`
	}{
		ObjectStoreBucketProxy: (*ObjectStoreBucketProxy)(p),
		ProviderConfig:         p.ProviderConfig,
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

func (p *ObjectStoreBucket) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectStoreBucket
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectStoreBucket()

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
	if known.Provider != nil {
		p.Provider = known.Provider
	}
	if known.ProviderConfigItemDiscriminator_ != nil {
		p.ProviderConfigItemDiscriminator_ = known.ProviderConfigItemDiscriminator_
	}
	if known.ProviderConfig != nil {
		p.ProviderConfig = known.ProviderConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bucketName")
	delete(allFields, "provider")
	delete(allFields, "$providerConfigItemDiscriminator")
	delete(allFields, "providerConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectStoreBucket() *ObjectStoreBucket {
	p := new(ObjectStoreBucket)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.ObjectStoreBucket"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ObjectStoreBucket) GetProviderConfig() interface{} {
	if nil == p.ProviderConfig {
		return nil
	}
	return p.ProviderConfig.GetValue()
}

func (p *ObjectStoreBucket) SetProviderConfig(v interface{}) error {
	if nil == p.ProviderConfig {
		p.ProviderConfig = NewOneOfObjectStoreBucketProviderConfig()
	}
	e := p.ProviderConfig.SetValue(v)
	if nil == e {
		if nil == p.ProviderConfigItemDiscriminator_ {
			p.ProviderConfigItemDiscriminator_ = new(string)
		}
		*p.ProviderConfigItemDiscriminator_ = *p.ProviderConfig.Discriminator
	}
	return e
}

/*
Object store provider like Azure, AWS, Nutanix etc.
*/
type ObjectStoreProvider int

const (
	OBJECTSTOREPROVIDER_UNKNOWN  ObjectStoreProvider = 0
	OBJECTSTOREPROVIDER_REDACTED ObjectStoreProvider = 1
	OBJECTSTOREPROVIDER_AZURE    ObjectStoreProvider = 2
	OBJECTSTOREPROVIDER_AWS      ObjectStoreProvider = 3
	OBJECTSTOREPROVIDER_NUTANIX  ObjectStoreProvider = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ObjectStoreProvider) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
		"NUTANIX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ObjectStoreProvider) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
		"NUTANIX",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ObjectStoreProvider) index(name string) ObjectStoreProvider {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AZURE",
		"AWS",
		"NUTANIX",
	}
	for idx := range names {
		if names[idx] == name {
			return ObjectStoreProvider(idx)
		}
	}
	return OBJECTSTOREPROVIDER_UNKNOWN
}

func (e *ObjectStoreProvider) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ObjectStoreProvider:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ObjectStoreProvider) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ObjectStoreProvider) Ref() *ObjectStoreProvider {
	return &e
}

/*
AWS Signature Version 4 credentials for authentication.
*/
type SigV4Credentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access key identifier for authentication.
	*/
	AccessKeyId *string `json:"accessKeyId"`
	/*
	  Secret access key for authentication.
	*/
	SecretAccessKey *string `json:"secretAccessKey"`
}

func (p *SigV4Credentials) MarshalJSON() ([]byte, error) {
	type SigV4CredentialsProxy SigV4Credentials

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SigV4CredentialsProxy
		AccessKeyId     *string `json:"accessKeyId,omitempty"`
		SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	}{
		SigV4CredentialsProxy: (*SigV4CredentialsProxy)(p),
		AccessKeyId:           p.AccessKeyId,
		SecretAccessKey:       p.SecretAccessKey,
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

func (p *SigV4Credentials) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SigV4Credentials
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSigV4Credentials()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessKeyId != nil {
		p.AccessKeyId = known.AccessKeyId
	}
	if known.SecretAccessKey != nil {
		p.SecretAccessKey = known.SecretAccessKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessKeyId")
	delete(allFields, "secretAccessKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSigV4Credentials() *SigV4Credentials {
	p := new(SigV4Credentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.SigV4Credentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model base class whose instances are bound to a specific tenant.  This model adds a tenantId to the base model class that it extends and is automatically set by the server.
*/
type TenantAwareModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *TenantAwareModel) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TenantAwareModel

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

func (p *TenantAwareModel) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TenantAwareModel
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenantAwareModel()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenantAwareModel() *TenantAwareModel {
	p := new(TenantAwareModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.TenantAwareModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfAmazonS3BucketAuthCredentials struct {
	Discriminator *string           `json:"-"`
	ObjectType_   *string           `json:"-"`
	oneOfType2101 *SigV4Credentials `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfAmazonS3BucketAuthCredentials() *OneOfAmazonS3BucketAuthCredentials {
	p := new(OneOfAmazonS3BucketAuthCredentials)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAmazonS3BucketAuthCredentials) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAmazonS3BucketAuthCredentials is nil"))
	}
	switch v.(type) {
	case SigV4Credentials:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(SigV4Credentials)
		}
		*p.oneOfType2101 = v.(SigV4Credentials)
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

func (p *OneOfAmazonS3BucketAuthCredentials) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfAmazonS3BucketAuthCredentials) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2101 := new(SigV4Credentials)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2101)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2101.ObjectType_ != nil && "common.v1.config.SigV4Credentials" == *vOneOfType2101.ObjectType_ {
							if nil == p.oneOfType2101 {
								p.oneOfType2101 = new(SigV4Credentials)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2101 := new(SigV4Credentials)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if vOneOfType2101.ObjectType_ != nil && "common.v1.config.SigV4Credentials" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(SigV4Credentials)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAmazonS3BucketAuthCredentials"))
}

func (p *OneOfAmazonS3BucketAuthCredentials) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfAmazonS3BucketAuthCredentials")
}

type OneOfAzureBlobStorageContainerAuthCredentials struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2101 *AzureSharedKeyCredentials `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfAzureBlobStorageContainerAuthCredentials() *OneOfAzureBlobStorageContainerAuthCredentials {
	p := new(OneOfAzureBlobStorageContainerAuthCredentials)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAzureBlobStorageContainerAuthCredentials) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAzureBlobStorageContainerAuthCredentials is nil"))
	}
	switch v.(type) {
	case AzureSharedKeyCredentials:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(AzureSharedKeyCredentials)
		}
		*p.oneOfType2101 = v.(AzureSharedKeyCredentials)
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

func (p *OneOfAzureBlobStorageContainerAuthCredentials) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfAzureBlobStorageContainerAuthCredentials) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2101 := new(AzureSharedKeyCredentials)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2101)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2101.ObjectType_ != nil && "common.v1.config.AzureSharedKeyCredentials" == *vOneOfType2101.ObjectType_ {
							if nil == p.oneOfType2101 {
								p.oneOfType2101 = new(AzureSharedKeyCredentials)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2101 := new(AzureSharedKeyCredentials)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if vOneOfType2101.ObjectType_ != nil && "common.v1.config.AzureSharedKeyCredentials" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(AzureSharedKeyCredentials)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAzureBlobStorageContainerAuthCredentials"))
}

func (p *OneOfAzureBlobStorageContainerAuthCredentials) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfAzureBlobStorageContainerAuthCredentials")
}

type OneOfNutanixObjectsBucketAuthCredentials struct {
	Discriminator *string           `json:"-"`
	ObjectType_   *string           `json:"-"`
	oneOfType2101 *SigV4Credentials `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfNutanixObjectsBucketAuthCredentials() *OneOfNutanixObjectsBucketAuthCredentials {
	p := new(OneOfNutanixObjectsBucketAuthCredentials)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfNutanixObjectsBucketAuthCredentials) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfNutanixObjectsBucketAuthCredentials is nil"))
	}
	switch v.(type) {
	case SigV4Credentials:
		if nil == p.oneOfType2101 {
			p.oneOfType2101 = new(SigV4Credentials)
		}
		*p.oneOfType2101 = v.(SigV4Credentials)
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

func (p *OneOfNutanixObjectsBucketAuthCredentials) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	return nil
}

func (p *OneOfNutanixObjectsBucketAuthCredentials) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2101 := new(SigV4Credentials)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2101)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2101.ObjectType_ != nil && "common.v1.config.SigV4Credentials" == *vOneOfType2101.ObjectType_ {
							if nil == p.oneOfType2101 {
								p.oneOfType2101 = new(SigV4Credentials)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2101 := new(SigV4Credentials)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if vOneOfType2101.ObjectType_ != nil && "common.v1.config.SigV4Credentials" == *vOneOfType2101.ObjectType_ {
			if nil == p.oneOfType2101 {
				p.oneOfType2101 = new(SigV4Credentials)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfNutanixObjectsBucketAuthCredentials"))
}

func (p *OneOfNutanixObjectsBucketAuthCredentials) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	return nil, errors.New("No value to marshal for OneOfNutanixObjectsBucketAuthCredentials")
}

type OneOfKVPairValue struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType1006 map[string]string    `json:"-"`
	oneOfType1004 *bool                `json:"-"`
	oneOfType1005 []string             `json:"-"`
	oneOfType1003 *int                 `json:"-"`
	oneOfType1008 []int                `json:"-"`
	oneOfType1002 *string              `json:"-"`
	oneOfType1007 []MapOfStringWrapper `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfKVPairValue() *OneOfKVPairValue {
	p := new(OneOfKVPairValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfKVPairValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfKVPairValue is nil"))
	}
	switch v.(type) {
	case map[string]string:
		p.oneOfType1006 = v.(map[string]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Map<String, String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Map<String, String>"
	case bool:
		if nil == p.oneOfType1004 {
			p.oneOfType1004 = new(bool)
		}
		*p.oneOfType1004 = v.(bool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
	case []string:
		p.oneOfType1005 = v.([]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
	case int:
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(int)
		}
		*p.oneOfType1003 = v.(int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
	case []int:
		p.oneOfType1008 = v.([]int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<Integer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<Integer>"
	case string:
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(string)
		}
		*p.oneOfType1002 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	case []MapOfStringWrapper:
		p.oneOfType1007 = v.([]MapOfStringWrapper)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<common.v1.config.MapOfStringWrapper>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<common.v1.config.MapOfStringWrapper>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfKVPairValue) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "Map<String, String>" == *p.Discriminator {
		return p.oneOfType1006
	}
	if "Boolean" == *p.Discriminator {
		return *p.oneOfType1004
	}
	if "List<String>" == *p.Discriminator {
		return p.oneOfType1005
	}
	if "Integer" == *p.Discriminator {
		return *p.oneOfType1003
	}
	if "List<Integer>" == *p.Discriminator {
		return p.oneOfType1008
	}
	if "String" == *p.Discriminator {
		return *p.oneOfType1002
	}
	if "List<common.v1.config.MapOfStringWrapper>" == *p.Discriminator {
		return p.oneOfType1007
	}
	return nil
}

func (p *OneOfKVPairValue) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["Map<String, String>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1006 := new(map[string]string)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1006)
					if unmarshalErr == nil {
						p.oneOfType1006 = *vOneOfType1006
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "Map<String, String>"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "Map<String, String>"
						return nil
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["Boolean"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1004 := new(bool)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1004)
					if unmarshalErr == nil {
						if nil == p.oneOfType1004 {
							p.oneOfType1004 = new(bool)
						}
						*p.oneOfType1004 = *vOneOfType1004
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<String>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new([]string)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						p.oneOfType1005 = *vOneOfType1005
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "List<String>"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "List<String>"
						return nil
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["Integer"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1003 := new(int)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1003)
					if unmarshalErr == nil {
						if nil == p.oneOfType1003 {
							p.oneOfType1003 = new(int)
						}
						*p.oneOfType1003 = *vOneOfType1003
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<Integer>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1008 := new([]int)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1008)
					if unmarshalErr == nil {
						p.oneOfType1008 = *vOneOfType1008
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "List<Integer>"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "List<Integer>"
						return nil
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["String"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1002 := new(string)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1002)
					if unmarshalErr == nil {
						if nil == p.oneOfType1002 {
							p.oneOfType1002 = new(string)
						}
						*p.oneOfType1002 = *vOneOfType1002
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "String"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "String"
						return nil
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<common.v1.config.MapOfStringWrapper>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1007 := new([]MapOfStringWrapper)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1007)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1007 == nil || len(*vOneOfType1007) == 0 || ((*vOneOfType1007)[0].ObjectType_ != nil && "common.v1.config.MapOfStringWrapper" == *((*vOneOfType1007)[0].ObjectType_)) {
							p.oneOfType1007 = *vOneOfType1007
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<common.v1.config.MapOfStringWrapper>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<common.v1.config.MapOfStringWrapper>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1006 := new(map[string]string)
	if err := json.Unmarshal(b, vOneOfType1006); err == nil {
		p.oneOfType1006 = *vOneOfType1006
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Map<String, String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Map<String, String>"
		return nil
	}
	vOneOfType1004 := new(bool)
	if err := json.Unmarshal(b, vOneOfType1004); err == nil {
		if nil == p.oneOfType1004 {
			p.oneOfType1004 = new(bool)
		}
		*p.oneOfType1004 = *vOneOfType1004
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
	vOneOfType1005 := new([]string)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		p.oneOfType1005 = *vOneOfType1005
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
		return nil
	}
	vOneOfType1003 := new(int)
	if err := json.Unmarshal(b, vOneOfType1003); err == nil {
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(int)
		}
		*p.oneOfType1003 = *vOneOfType1003
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
	vOneOfType1008 := new([]int)
	if err := json.Unmarshal(b, vOneOfType1008); err == nil {
		p.oneOfType1008 = *vOneOfType1008
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<Integer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<Integer>"
		return nil
	}
	vOneOfType1002 := new(string)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(string)
		}
		*p.oneOfType1002 = *vOneOfType1002
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
		return nil
	}
	vOneOfType1007 := new([]MapOfStringWrapper)
	if err := json.Unmarshal(b, vOneOfType1007); err == nil {
		if len(*vOneOfType1007) == 0 || (vOneOfType1007 != nil && (*vOneOfType1007)[0].ObjectType_ != nil && "common.v1.config.MapOfStringWrapper" == *((*vOneOfType1007)[0].ObjectType_)) {
			p.oneOfType1007 = *vOneOfType1007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<common.v1.config.MapOfStringWrapper>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<common.v1.config.MapOfStringWrapper>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfKVPairValue"))
}

func (p *OneOfKVPairValue) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "Map<String, String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1006)
	}
	if "Boolean" == *p.Discriminator {
		return json.Marshal(p.oneOfType1004)
	}
	if "List<String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	if "Integer" == *p.Discriminator {
		return json.Marshal(p.oneOfType1003)
	}
	if "List<Integer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1008)
	}
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if "List<common.v1.config.MapOfStringWrapper>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1007)
	}
	return nil, errors.New("No value to marshal for OneOfKVPairValue")
}

type OneOfObjectStoreBucketProviderConfig struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2101 *AzureBlobStorageContainer `json:"-"`
	oneOfType2103 *NutanixObjectsBucket      `json:"-"`
	oneOfType2102 *AmazonS3Bucket            `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfObjectStoreBucketProviderConfig() *OneOfObjectStoreBucketProviderConfig {
	p := new(OneOfObjectStoreBucketProviderConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfObjectStoreBucketProviderConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfObjectStoreBucketProviderConfig is nil"))
	}
	switch v.(type) {
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfObjectStoreBucketProviderConfig) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2101
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2103
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2102
	}
	return nil
}

func (p *OneOfObjectStoreBucketProviderConfig) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2101 := new(AzureBlobStorageContainer)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2101)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2101.ObjectType_ != nil && "common.v1.config.AzureBlobStorageContainer" == *vOneOfType2101.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2103 := new(NutanixObjectsBucket)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2103)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2103.ObjectType_ != nil && "common.v1.config.NutanixObjectsBucket" == *vOneOfType2103.ObjectType_ {
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2102 := new(AmazonS3Bucket)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2102)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2102.ObjectType_ != nil && "common.v1.config.AmazonS3Bucket" == *vOneOfType2102.ObjectType_ {
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2101 := new(AzureBlobStorageContainer)
	if err := json.Unmarshal(b, vOneOfType2101); err == nil {
		if vOneOfType2101.ObjectType_ != nil && "common.v1.config.AzureBlobStorageContainer" == *vOneOfType2101.ObjectType_ {
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
	vOneOfType2103 := new(NutanixObjectsBucket)
	if err := json.Unmarshal(b, vOneOfType2103); err == nil {
		if vOneOfType2103.ObjectType_ != nil && "common.v1.config.NutanixObjectsBucket" == *vOneOfType2103.ObjectType_ {
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
		if vOneOfType2102.ObjectType_ != nil && "common.v1.config.AmazonS3Bucket" == *vOneOfType2102.ObjectType_ {
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfObjectStoreBucketProviderConfig"))
}

func (p *OneOfObjectStoreBucketProviderConfig) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2101 != nil && *p.oneOfType2101.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2101)
	}
	if p.oneOfType2103 != nil && *p.oneOfType2103.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2103)
	}
	if p.oneOfType2102 != nil && *p.oneOfType2102.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2102)
	}
	return nil, errors.New("No value to marshal for OneOfObjectStoreBucketProviderConfig")
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
