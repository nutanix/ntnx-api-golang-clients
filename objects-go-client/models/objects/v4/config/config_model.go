/*
 * Generated file models/objects/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Objects Storage Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module objects.v4.config of Nutanix Objects Storage Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/prism/v4/config"
	"time"
)

type Certificate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of alternate FQDNs for accessing the Object store. The FQDNs must consist of at least 2 parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens or underscores but must begin and end with a letter. Each part can be up to 63 characters long. For e.g 'objects-0.pc_nutanix.com'.
	*/
	AlternateFqdns []import1.FQDN `json:"alternateFqdns,omitempty"`
	/*
	  A list of the IPs included as Subject Alternative Names (SANs) in the certificate. The IPs must be among the public IPs of the Object store (publicNetworkIps).
	*/
	AlternateIps []import1.IPAddress `json:"alternateIps,omitempty"`
	/*
	  The CA certificate or chain to upload.
	*/
	Ca *string `json:"ca,omitempty"`
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
	  The private key to upload.
	*/
	PrivateKey *string `json:"privateKey,omitempty"`
	/*
	  The public certificate to upload.
	*/
	PublicCert *string `json:"publicCert,omitempty"`
	/*
	  If true, the certificate is generated with the provided alternate FQDNs and IPs.
	*/
	ShouldGenerate *bool `json:"shouldGenerate,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCertificate() *Certificate {
	p := new(Certificate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.Certificate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldGenerate = new(bool)
	*p.ShouldGenerate = false

	return p
}

type CertificateProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of alternate FQDNs for accessing the Object store. The FQDNs must consist of at least 2 parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens or underscores but must begin and end with a letter. Each part can be up to 63 characters long. For e.g 'objects-0.pc_nutanix.com'.
	*/
	AlternateFqdns []import1.FQDN `json:"alternateFqdns,omitempty"`
	/*
	  A list of the IPs included as Subject Alternative Names (SANs) in the certificate. The IPs must be among the public IPs of the Object store (publicNetworkIps).
	*/
	AlternateIps []import1.IPAddress `json:"alternateIps,omitempty"`
	/*
	  The CA certificate or chain to upload.
	*/
	Ca *string `json:"ca,omitempty"`
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
	  The private key to upload.
	*/
	PrivateKey *string `json:"privateKey,omitempty"`
	/*
	  The public certificate to upload.
	*/
	PublicCert *string `json:"publicCert,omitempty"`
	/*
	  If true, the certificate is generated with the provided alternate FQDNs and IPs.
	*/
	ShouldGenerate *bool `json:"shouldGenerate,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCertificateProjection() *CertificateProjection {
	p := new(CertificateProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.CertificateProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldGenerate = new(bool)
	*p.ShouldGenerate = false

	return p
}

/*
REST response for all response codes in API path /objects/v4.0/config/object-stores/{objectStoreExtId}/certificates Post operation
*/
type CreateCertificateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCertificateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateCertificateApiResponse() *CreateCertificateApiResponse {
	p := new(CreateCertificateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.CreateCertificateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCertificateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCertificateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCertificateApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores Post operation
*/
type CreateObjectstoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateObjectstoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateObjectstoreApiResponse() *CreateObjectstoreApiResponse {
	p := new(CreateObjectstoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.CreateObjectstoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateObjectstoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateObjectstoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateObjectstoreApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{extId} Delete operation
*/
type DeleteObjectstoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteObjectstoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteObjectstoreApiResponse() *DeleteObjectstoreApiResponse {
	p := new(DeleteObjectstoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.DeleteObjectstoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteObjectstoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteObjectstoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteObjectstoreApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{objectStoreExtId}/certificates/{certificateExtId}/certificate-authority Get operation
*/
type GetCaApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCaApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCaApiResponse() *GetCaApiResponse {
	p := new(GetCaApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.GetCaApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCaApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCaApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCaApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{objectStoreExtId}/certificates/{extId} Get operation
*/
type GetCertificateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCertificateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCertificateApiResponse() *GetCertificateApiResponse {
	p := new(GetCertificateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.GetCertificateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCertificateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCertificateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCertificateApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{extId} Get operation
*/
type GetObjectstoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetObjectstoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetObjectstoreApiResponse() *GetObjectstoreApiResponse {
	p := new(GetObjectstoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.GetObjectstoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetObjectstoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetObjectstoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetObjectstoreApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{objectStoreExtId}/certificates Get operation
*/
type ListCertificatesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCertificatesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCertificatesApiResponse() *ListCertificatesApiResponse {
	p := new(ListCertificatesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.ListCertificatesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCertificatesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCertificatesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCertificatesApiResponseData()
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
REST response for all response codes in API path /objects/v4.0/config/object-stores Get operation
*/
type ListObjectstoresApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListObjectstoresApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListObjectstoresApiResponse() *ListObjectstoresApiResponse {
	p := new(ListObjectstoresApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.ListObjectstoresApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListObjectstoresApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListObjectstoresApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListObjectstoresApiResponseData()
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

type ObjectStore struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of the UUIDs of the certificates of an Object store.
	*/
	CertificateExtIds []string `json:"certificateExtIds,omitempty"`
	/*
	  UUID of the AHV or ESXi cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  The time when the Object store was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  Object store deployment version.
	*/
	DeploymentVersion *string `json:"deploymentVersion,omitempty"`
	/*
	  A brief description of the Object store.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The DNS domain/subdomain the Object store belongs to. All the Object stores under one Prism Central must have the same domain name. The domain name must consist of at least 2 parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens, or underscores. Each part can be up to 63 characters long. The domain must begin and end with an alphanumeric character. For example - 'objects-0.pc_nutanix.com'.
	*/
	Domain *string `json:"domain,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the Object store was last updated.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  The name of the Object store.
	*/
	Name *string `json:"name"`
	/*
	  The number of worker nodes (VMs) to be created for the Object store. Each worker node requires 10 vCPUs and 32 GiB of memory.
	*/
	NumWorkerNodes *int64 `json:"numWorkerNodes,omitempty"`
	/*
	  A list of static IP addresses used as public IPs to access the Object store.
	*/
	PublicNetworkIps []import1.IPAddress `json:"publicNetworkIps,omitempty"`
	/*
	  Public network reference of the Object store. This is the subnet UUID for an AHV cluster or the IPAM name for an ESXi cluster.
	*/
	PublicNetworkReference *string `json:"publicNetworkReference,omitempty"`
	/*
	  The region in which the Object store is deployed.
	*/
	Region *string `json:"region,omitempty"`

	State *State `json:"state,omitempty"`

	StorageNetworkDnsIp *import1.IPAddress `json:"storageNetworkDnsIp,omitempty"`
	/*
	  Reference to the Storage Network of the Object store. This is the subnet UUID for an AHV cluster or the IPAM name for an ESXi cluster.
	*/
	StorageNetworkReference *string `json:"storageNetworkReference,omitempty"`

	StorageNetworkVip *import1.IPAddress `json:"storageNetworkVip,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Size of the Object store in GiB.
	*/
	TotalCapacityGiB *int64 `json:"totalCapacityGiB,omitempty"`
}

func (p *ObjectStore) MarshalJSON() ([]byte, error) {
	type ObjectStoreProxy ObjectStore
	return json.Marshal(struct {
		*ObjectStoreProxy
		Name *string `json:"name,omitempty"`
	}{
		ObjectStoreProxy: (*ObjectStoreProxy)(p),
		Name:             p.Name,
	})
}

func NewObjectStore() *ObjectStore {
	p := new(ObjectStore)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.ObjectStore"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ObjectStoreProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A list of the UUIDs of the certificates of an Object store.
	*/
	CertificateExtIds []string `json:"certificateExtIds,omitempty"`

	CertificateProjection []CertificateProjection `json:"certificateProjection,omitempty"`
	/*
	  UUID of the AHV or ESXi cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  The time when the Object store was created.
	*/
	CreationTime *time.Time `json:"creationTime,omitempty"`
	/*
	  Object store deployment version.
	*/
	DeploymentVersion *string `json:"deploymentVersion,omitempty"`
	/*
	  A brief description of the Object store.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The DNS domain/subdomain the Object store belongs to. All the Object stores under one Prism Central must have the same domain name. The domain name must consist of at least 2 parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens, or underscores. Each part can be up to 63 characters long. The domain must begin and end with an alphanumeric character. For example - 'objects-0.pc_nutanix.com'.
	*/
	Domain *string `json:"domain,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the Object store was last updated.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Metadata *import1.Metadata `json:"metadata,omitempty"`
	/*
	  The name of the Object store.
	*/
	Name *string `json:"name"`
	/*
	  The number of worker nodes (VMs) to be created for the Object store. Each worker node requires 10 vCPUs and 32 GiB of memory.
	*/
	NumWorkerNodes *int64 `json:"numWorkerNodes,omitempty"`
	/*
	  A list of static IP addresses used as public IPs to access the Object store.
	*/
	PublicNetworkIps []import1.IPAddress `json:"publicNetworkIps,omitempty"`
	/*
	  Public network reference of the Object store. This is the subnet UUID for an AHV cluster or the IPAM name for an ESXi cluster.
	*/
	PublicNetworkReference *string `json:"publicNetworkReference,omitempty"`
	/*
	  The region in which the Object store is deployed.
	*/
	Region *string `json:"region,omitempty"`

	State *State `json:"state,omitempty"`

	StorageNetworkDnsIp *import1.IPAddress `json:"storageNetworkDnsIp,omitempty"`
	/*
	  Reference to the Storage Network of the Object store. This is the subnet UUID for an AHV cluster or the IPAM name for an ESXi cluster.
	*/
	StorageNetworkReference *string `json:"storageNetworkReference,omitempty"`

	StorageNetworkVip *import1.IPAddress `json:"storageNetworkVip,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Size of the Object store in GiB.
	*/
	TotalCapacityGiB *int64 `json:"totalCapacityGiB,omitempty"`
}

func (p *ObjectStoreProjection) MarshalJSON() ([]byte, error) {
	type ObjectStoreProjectionProxy ObjectStoreProjection
	return json.Marshal(struct {
		*ObjectStoreProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		ObjectStoreProjectionProxy: (*ObjectStoreProjectionProxy)(p),
		Name:                       p.Name,
	})
}

func NewObjectStoreProjection() *ObjectStoreProjection {
	p := new(ObjectStoreProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.ObjectStoreProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum for the state of the Object store.
*/
type State int

const (
	STATE_UNKNOWN                           State = 0
	STATE_REDACTED                          State = 1
	STATE_OBJECT_STORE_DEPLOYMENT_FAILED    State = 2
	STATE_OBJECT_STORE_CERT_CREATION_FAILED State = 3
	STATE_OBJECT_STORE_DELETION_FAILED      State = 4
	STATE_UNDEPLOYED_OBJECT_STORE           State = 5
	STATE_DEPLOYING_OBJECT_STORE            State = 6
	STATE_CREATING_OBJECT_STORE_CERT        State = 7
	STATE_DELETING_OBJECT_STORE             State = 8
	STATE_OBJECT_STORE_AVAILABLE            State = 9
	STATE_OBJECT_STORE_OPERATION_PENDING    State = 10
	STATE_OBJECT_STORE_OPERATION_FAILED     State = 11
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *State) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OBJECT_STORE_DEPLOYMENT_FAILED",
		"OBJECT_STORE_CERT_CREATION_FAILED",
		"OBJECT_STORE_DELETION_FAILED",
		"UNDEPLOYED_OBJECT_STORE",
		"DEPLOYING_OBJECT_STORE",
		"CREATING_OBJECT_STORE_CERT",
		"DELETING_OBJECT_STORE",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_FAILED",
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
		"OBJECT_STORE_DEPLOYMENT_FAILED",
		"OBJECT_STORE_CERT_CREATION_FAILED",
		"OBJECT_STORE_DELETION_FAILED",
		"UNDEPLOYED_OBJECT_STORE",
		"DEPLOYING_OBJECT_STORE",
		"CREATING_OBJECT_STORE_CERT",
		"DELETING_OBJECT_STORE",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_FAILED",
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
		"OBJECT_STORE_DEPLOYMENT_FAILED",
		"OBJECT_STORE_CERT_CREATION_FAILED",
		"OBJECT_STORE_DELETION_FAILED",
		"UNDEPLOYED_OBJECT_STORE",
		"DEPLOYING_OBJECT_STORE",
		"CREATING_OBJECT_STORE_CERT",
		"DELETING_OBJECT_STORE",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_FAILED",
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
REST response for all response codes in API path /objects/v4.0/config/object-stores/{extId} Put operation
*/
type UpdateObjectstoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateObjectstoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateObjectstoreApiResponse() *UpdateObjectstoreApiResponse {
	p := new(UpdateObjectstoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.config.UpdateObjectstoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateObjectstoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateObjectstoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateObjectstoreApiResponseData()
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

type OneOfGetObjectstoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ObjectStore           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetObjectstoreApiResponseData() *OneOfGetObjectstoreApiResponseData {
	p := new(OneOfGetObjectstoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetObjectstoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetObjectstoreApiResponseData is nil"))
	}
	switch v.(type) {
	case ObjectStore:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ObjectStore)
		}
		*p.oneOfType0 = v.(ObjectStore)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfGetObjectstoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetObjectstoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ObjectStore)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "objects.v4.config.ObjectStore" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ObjectStore)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetObjectstoreApiResponseData"))
}

func (p *OneOfGetObjectstoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetObjectstoreApiResponseData")
}

type OneOfCreateObjectstoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateObjectstoreApiResponseData() *OneOfCreateObjectstoreApiResponseData {
	p := new(OneOfCreateObjectstoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateObjectstoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateObjectstoreApiResponseData is nil"))
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

func (p *OneOfCreateObjectstoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateObjectstoreApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateObjectstoreApiResponseData"))
}

func (p *OneOfCreateObjectstoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateObjectstoreApiResponseData")
}

type OneOfCreateCertificateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateCertificateApiResponseData() *OneOfCreateCertificateApiResponseData {
	p := new(OneOfCreateCertificateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCertificateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCertificateApiResponseData is nil"))
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

func (p *OneOfCreateCertificateApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCertificateApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCertificateApiResponseData"))
}

func (p *OneOfCreateCertificateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCertificateApiResponseData")
}

type OneOfListObjectstoresApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    []ObjectStore           `json:"-"`
	oneOfType400  *import4.ErrorResponse  `json:"-"`
	oneOfType401  []ObjectStoreProjection `json:"-"`
}

func NewOneOfListObjectstoresApiResponseData() *OneOfListObjectstoresApiResponseData {
	p := new(OneOfListObjectstoresApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListObjectstoresApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListObjectstoresApiResponseData is nil"))
	}
	switch v.(type) {
	case []ObjectStore:
		p.oneOfType0 = v.([]ObjectStore)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.config.ObjectStore>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.config.ObjectStore>"
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
	case []ObjectStoreProjection:
		p.oneOfType401 = v.([]ObjectStoreProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.config.ObjectStoreProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.config.ObjectStoreProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListObjectstoresApiResponseData) GetValue() interface{} {
	if "List<objects.v4.config.ObjectStore>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<objects.v4.config.ObjectStoreProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListObjectstoresApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ObjectStore)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "objects.v4.config.ObjectStore" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.config.ObjectStore>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.config.ObjectStore>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]ObjectStoreProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "objects.v4.config.ObjectStoreProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.config.ObjectStoreProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.config.ObjectStoreProjection>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListObjectstoresApiResponseData"))
}

func (p *OneOfListObjectstoresApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<objects.v4.config.ObjectStore>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<objects.v4.config.ObjectStoreProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListObjectstoresApiResponseData")
}

type OneOfGetCaApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType0    *FileDetail            `json:"-"`
}

func NewOneOfGetCaApiResponseData() *OneOfGetCaApiResponseData {
	p := new(OneOfGetCaApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCaApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCaApiResponseData is nil"))
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
	case FileDetail:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = v.(FileDetail)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetCaApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetCaApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(FileDetail)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCaApiResponseData"))
}

func (p *OneOfGetCaApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetCaApiResponseData")
}

type OneOfDeleteObjectstoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteObjectstoreApiResponseData() *OneOfDeleteObjectstoreApiResponseData {
	p := new(OneOfDeleteObjectstoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteObjectstoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteObjectstoreApiResponseData is nil"))
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

func (p *OneOfDeleteObjectstoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteObjectstoreApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteObjectstoreApiResponseData"))
}

func (p *OneOfDeleteObjectstoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteObjectstoreApiResponseData")
}

type OneOfListCertificatesApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import4.ErrorResponse  `json:"-"`
	oneOfType401  []CertificateProjection `json:"-"`
	oneOfType0    []Certificate           `json:"-"`
}

func NewOneOfListCertificatesApiResponseData() *OneOfListCertificatesApiResponseData {
	p := new(OneOfListCertificatesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCertificatesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCertificatesApiResponseData is nil"))
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
	case []CertificateProjection:
		p.oneOfType401 = v.([]CertificateProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.config.CertificateProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.config.CertificateProjection>"
	case []Certificate:
		p.oneOfType0 = v.([]Certificate)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.config.Certificate>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.config.Certificate>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCertificatesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<objects.v4.config.CertificateProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<objects.v4.config.Certificate>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListCertificatesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]CertificateProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "objects.v4.config.CertificateProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.config.CertificateProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.config.CertificateProjection>"
			return nil
		}
	}
	vOneOfType0 := new([]Certificate)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "objects.v4.config.Certificate" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.config.Certificate>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.config.Certificate>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCertificatesApiResponseData"))
}

func (p *OneOfListCertificatesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<objects.v4.config.CertificateProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<objects.v4.config.Certificate>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListCertificatesApiResponseData")
}

type OneOfGetCertificateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Certificate           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetCertificateApiResponseData() *OneOfGetCertificateApiResponseData {
	p := new(OneOfGetCertificateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCertificateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCertificateApiResponseData is nil"))
	}
	switch v.(type) {
	case Certificate:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Certificate)
		}
		*p.oneOfType0 = v.(Certificate)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfGetCertificateApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetCertificateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Certificate)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "objects.v4.config.Certificate" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Certificate)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCertificateApiResponseData"))
}

func (p *OneOfGetCertificateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetCertificateApiResponseData")
}

type OneOfUpdateObjectstoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateObjectstoreApiResponseData() *OneOfUpdateObjectstoreApiResponseData {
	p := new(OneOfUpdateObjectstoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateObjectstoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateObjectstoreApiResponseData is nil"))
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

func (p *OneOfUpdateObjectstoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateObjectstoreApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateObjectstoreApiResponseData"))
}

func (p *OneOfUpdateObjectstoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateObjectstoreApiResponseData")
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
