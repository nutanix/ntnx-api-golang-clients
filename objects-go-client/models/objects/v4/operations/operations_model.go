/*
 * Generated file models/objects/v4/operations/operations_model.go.
 *
 * Product version: 4.0.1-alpha-2
 *
 * Part of the Nutanix Objects Storage Management APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Perform Objects operations
*/
package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/prism/v4/config"
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
	  CA file to upload.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCertificate() *Certificate {
	p := new(Certificate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.Certificate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates/{extId} Get operation
*/
type CertificateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCertificateApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCertificateApiResponse() *CertificateApiResponse {
	p := new(CertificateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.CertificateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CertificateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CertificateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCertificateApiResponseData()
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
REST response for all response codes in API path /objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates Get operation
*/
type CertificateListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCertificateListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCertificateListApiResponse() *CertificateListApiResponse {
	p := new(CertificateListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.CertificateListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CertificateListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CertificateListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCertificateListApiResponseData()
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

type CertificateProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of alternate FQDNs for accessing the Object store. The FQDNs must consist of at least 2 parts separated by a '.'. Each part can contain upper and lower case letters, digits, hyphens or underscores but must begin and end with a letter. Each part can be up to 63 characters long. For e.g 'objects-0.pc_nutanix.com'.
	*/
	AlternateFqdns []import1.FQDN `json:"alternateFqdns,omitempty"`
	/*
	  CA file to upload.
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCertificateProjection() *CertificateProjection {
	p := new(CertificateProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.CertificateProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Objectstore struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID reference to the Object store certificate.
	*/
	CertificateReference *string `json:"certificateReference,omitempty"`
	/*
	  UUID of the AHV or ESXi cluster.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
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

	State *StateEnum `json:"state,omitempty"`

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

func (p *Objectstore) MarshalJSON() ([]byte, error) {
	type ObjectstoreProxy Objectstore
	return json.Marshal(struct {
		*ObjectstoreProxy
		Name *string `json:"name,omitempty"`
	}{
		ObjectstoreProxy: (*ObjectstoreProxy)(p),
		Name:             p.Name,
	})
}

func NewObjectstore() *Objectstore {
	p := new(Objectstore)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.Objectstore"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /objects/v4.0.a2/operations/object-stores/{extId} Get operation
*/
type ObjectstoreApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfObjectstoreApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewObjectstoreApiResponse() *ObjectstoreApiResponse {
	p := new(ObjectstoreApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.ObjectstoreApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ObjectstoreApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ObjectstoreApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfObjectstoreApiResponseData()
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
REST response for all response codes in API path /objects/v4.0.a2/operations/object-stores Get operation
*/
type ObjectstoreListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfObjectstoreListApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewObjectstoreListApiResponse() *ObjectstoreListApiResponse {
	p := new(ObjectstoreListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.ObjectstoreListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ObjectstoreListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ObjectstoreListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfObjectstoreListApiResponseData()
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

type ObjectstoreProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CertificateProjection *CertificateProjection `json:"certificateProjection,omitempty"`
	/*
	  UUID reference to the Object store certificate.
	*/
	CertificateReference *string `json:"certificateReference,omitempty"`
	/*
	  UUID of the AHV or ESXi cluster.
	*/
	ClusterReference *string `json:"clusterReference,omitempty"`
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

	State *StateEnum `json:"state,omitempty"`

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

func (p *ObjectstoreProjection) MarshalJSON() ([]byte, error) {
	type ObjectstoreProjectionProxy ObjectstoreProjection
	return json.Marshal(struct {
		*ObjectstoreProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		ObjectstoreProjectionProxy: (*ObjectstoreProjectionProxy)(p),
		Name:                       p.Name,
	})
}

func NewObjectstoreProjection() *ObjectstoreProjection {
	p := new(ObjectstoreProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.ObjectstoreProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum for the state of the Object store.
*/
type StateEnum int

const (
	STATEENUM_UNKNOWN                         StateEnum = 0
	STATEENUM_REDACTED                        StateEnum = 1
	STATEENUM_OBJECT_STORE_DEPLOYMENT_ERROR   StateEnum = 2
	STATEENUM_OBJECT_STORE_CREATE_CERT_ERROR  StateEnum = 3
	STATEENUM_OBJECT_STORE_DELETE_ERROR       StateEnum = 4
	STATEENUM_OBJECT_STORE_DELETE_INPUT_ERROR StateEnum = 5
	STATEENUM_OBJECT_STORE_INPUT              StateEnum = 6
	STATEENUM_OBJECT_STORE_DEPLOYING          StateEnum = 7
	STATEENUM_OBJECT_STORE_CREATING_CERT      StateEnum = 8
	STATEENUM_OBJECT_STORE_DELETING           StateEnum = 9
	STATEENUM_OBJECT_STORE_DELETING_INPUT     StateEnum = 10
	STATEENUM_OBJECT_STORE_AVAILABLE          StateEnum = 11
	STATEENUM_OBJECT_STORE_OPERATION_PENDING  StateEnum = 12
	STATEENUM_OBJECT_STORE_OPERATION_ERROR    StateEnum = 13
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *StateEnum) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OBJECT_STORE_DEPLOYMENT_ERROR",
		"OBJECT_STORE_CREATE_CERT_ERROR",
		"OBJECT_STORE_DELETE_ERROR",
		"OBJECT_STORE_DELETE_INPUT_ERROR",
		"OBJECT_STORE_INPUT",
		"OBJECT_STORE_DEPLOYING",
		"OBJECT_STORE_CREATING_CERT",
		"OBJECT_STORE_DELETING",
		"OBJECT_STORE_DELETING_INPUT",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e StateEnum) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OBJECT_STORE_DEPLOYMENT_ERROR",
		"OBJECT_STORE_CREATE_CERT_ERROR",
		"OBJECT_STORE_DELETE_ERROR",
		"OBJECT_STORE_DELETE_INPUT_ERROR",
		"OBJECT_STORE_INPUT",
		"OBJECT_STORE_DEPLOYING",
		"OBJECT_STORE_CREATING_CERT",
		"OBJECT_STORE_DELETING",
		"OBJECT_STORE_DELETING_INPUT",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *StateEnum) index(name string) StateEnum {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OBJECT_STORE_DEPLOYMENT_ERROR",
		"OBJECT_STORE_CREATE_CERT_ERROR",
		"OBJECT_STORE_DELETE_ERROR",
		"OBJECT_STORE_DELETE_INPUT_ERROR",
		"OBJECT_STORE_INPUT",
		"OBJECT_STORE_DEPLOYING",
		"OBJECT_STORE_CREATING_CERT",
		"OBJECT_STORE_DELETING",
		"OBJECT_STORE_DELETING_INPUT",
		"OBJECT_STORE_AVAILABLE",
		"OBJECT_STORE_OPERATION_PENDING",
		"OBJECT_STORE_OPERATION_ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return StateEnum(idx)
		}
	}
	return STATEENUM_UNKNOWN
}

func (e *StateEnum) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StateEnum:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StateEnum) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StateEnum) Ref() *StateEnum {
	return &e
}

/*
REST response for all response codes in API path /objects/v4.0.a2/operations/object-stores/{objectStoreExtId}/certificates Post operation
*/
type TaskReferenceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTaskReferenceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskReferenceApiResponse() *TaskReferenceApiResponse {
	p := new(TaskReferenceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.operations.TaskReferenceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0.a2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TaskReferenceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TaskReferenceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTaskReferenceApiResponseData()
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

type OneOfCertificateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Certificate           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCertificateApiResponseData() *OneOfCertificateApiResponseData {
	p := new(OneOfCertificateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCertificateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCertificateApiResponseData is nil"))
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

func (p *OneOfCertificateApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCertificateApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Certificate)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "objects.v4.operations.Certificate" == *vOneOfType0.ObjectType_ {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCertificateApiResponseData"))
}

func (p *OneOfCertificateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCertificateApiResponseData")
}

type OneOfObjectstoreApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Objectstore           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfObjectstoreApiResponseData() *OneOfObjectstoreApiResponseData {
	p := new(OneOfObjectstoreApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfObjectstoreApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfObjectstoreApiResponseData is nil"))
	}
	switch v.(type) {
	case Objectstore:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Objectstore)
		}
		*p.oneOfType0 = v.(Objectstore)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfObjectstoreApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfObjectstoreApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Objectstore)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "objects.v4.operations.Objectstore" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Objectstore)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfObjectstoreApiResponseData"))
}

func (p *OneOfObjectstoreApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfObjectstoreApiResponseData")
}

type OneOfTaskReferenceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfTaskReferenceApiResponseData() *OneOfTaskReferenceApiResponseData {
	p := new(OneOfTaskReferenceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTaskReferenceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTaskReferenceApiResponseData is nil"))
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

func (p *OneOfTaskReferenceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTaskReferenceApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskReferenceApiResponseData"))
}

func (p *OneOfTaskReferenceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTaskReferenceApiResponseData")
}

type OneOfCertificateListApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType401  []CertificateProjection `json:"-"`
	oneOfType0    []Certificate           `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
}

func NewOneOfCertificateListApiResponseData() *OneOfCertificateListApiResponseData {
	p := new(OneOfCertificateListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCertificateListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCertificateListApiResponseData is nil"))
	}
	switch v.(type) {
	case []CertificateProjection:
		p.oneOfType401 = v.([]CertificateProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.operations.CertificateProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.operations.CertificateProjection>"
	case []Certificate:
		p.oneOfType0 = v.([]Certificate)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.operations.Certificate>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.operations.Certificate>"
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

func (p *OneOfCertificateListApiResponseData) GetValue() interface{} {
	if "List<objects.v4.operations.CertificateProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<objects.v4.operations.Certificate>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCertificateListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]CertificateProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {

		if len(*vOneOfType401) == 0 || "objects.v4.operations.CertificateProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.operations.CertificateProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.operations.CertificateProjection>"
			return nil

		}
	}
	vOneOfType0 := new([]Certificate)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "objects.v4.operations.Certificate" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.operations.Certificate>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.operations.Certificate>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCertificateListApiResponseData"))
}

func (p *OneOfCertificateListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<objects.v4.operations.CertificateProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<objects.v4.operations.Certificate>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCertificateListApiResponseData")
}

type OneOfObjectstoreListApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    []Objectstore           `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
	oneOfType401  []ObjectstoreProjection `json:"-"`
}

func NewOneOfObjectstoreListApiResponseData() *OneOfObjectstoreListApiResponseData {
	p := new(OneOfObjectstoreListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfObjectstoreListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfObjectstoreListApiResponseData is nil"))
	}
	switch v.(type) {
	case []Objectstore:
		p.oneOfType0 = v.([]Objectstore)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.operations.Objectstore>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.operations.Objectstore>"
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
	case []ObjectstoreProjection:
		p.oneOfType401 = v.([]ObjectstoreProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<objects.v4.operations.ObjectstoreProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<objects.v4.operations.ObjectstoreProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfObjectstoreListApiResponseData) GetValue() interface{} {
	if "List<objects.v4.operations.Objectstore>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<objects.v4.operations.ObjectstoreProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfObjectstoreListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Objectstore)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {

		if len(*vOneOfType0) == 0 || "objects.v4.operations.Objectstore" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.operations.Objectstore>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.operations.Objectstore>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]ObjectstoreProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {

		if len(*vOneOfType401) == 0 || "objects.v4.operations.ObjectstoreProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<objects.v4.operations.ObjectstoreProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<objects.v4.operations.ObjectstoreProjection>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfObjectstoreListApiResponseData"))
}

func (p *OneOfObjectstoreListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<objects.v4.operations.Objectstore>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<objects.v4.operations.ObjectstoreProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfObjectstoreListApiResponseData")
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
