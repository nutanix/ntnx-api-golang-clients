/*
 * Generated file models/multidomain/v4/resources/resources_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module multidomain.v4.resources of Nutanix Multidomain Versioned APIs
*/
package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
	"time"
)

/*
Information for domain
*/
type Domain struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DomainAuditTrail *DomainAuditTrail `json:"domainAuditTrail,omitempty"`

	DomainDetails *DomainDetails `json:"domainDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Whether domain is reachable from Nutanix central or not.
	*/
	IsReachable *bool `json:"isReachable,omitempty"`
	/*
	  API keyId Reference of domain
	*/
	KeyIdReference *string `json:"keyIdReference,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Location *Location `json:"location,omitempty"`
	/*

	 */
	PlatformDataItemDiscriminator_ *string `json:"$platformDataItemDiscriminator,omitempty"`
	/*
	  Platform specific Data.
	*/
	PlatformData *OneOfDomainPlatformData `json:"platformData,omitempty"`

	RegistrationState *DomainRegistrationState `json:"registrationState,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Communication between Nutanix Central and Prism central will be done through the use of a tunnel provider.
	*/
	TunnelProvider []string `json:"tunnelProvider,omitempty"`
}

func (p *Domain) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Domain

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

func (p *Domain) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Domain
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomain()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DomainAuditTrail != nil {
		p.DomainAuditTrail = known.DomainAuditTrail
	}
	if known.DomainDetails != nil {
		p.DomainDetails = known.DomainDetails
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsReachable != nil {
		p.IsReachable = known.IsReachable
	}
	if known.KeyIdReference != nil {
		p.KeyIdReference = known.KeyIdReference
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.PlatformDataItemDiscriminator_ != nil {
		p.PlatformDataItemDiscriminator_ = known.PlatformDataItemDiscriminator_
	}
	if known.PlatformData != nil {
		p.PlatformData = known.PlatformData
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TunnelProvider != nil {
		p.TunnelProvider = known.TunnelProvider
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domainAuditTrail")
	delete(allFields, "domainDetails")
	delete(allFields, "extId")
	delete(allFields, "isReachable")
	delete(allFields, "keyIdReference")
	delete(allFields, "links")
	delete(allFields, "location")
	delete(allFields, "$platformDataItemDiscriminator")
	delete(allFields, "platformData")
	delete(allFields, "registrationState")
	delete(allFields, "tenantId")
	delete(allFields, "tunnelProvider")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomain() *Domain {
	p := new(Domain)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.Domain"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsReachable = new(bool)
	*p.IsReachable = false

	return p
}

func (p *Domain) GetPlatformData() interface{} {
	if nil == p.PlatformData {
		return nil
	}
	return p.PlatformData.GetValue()
}

func (p *Domain) SetPlatformData(v interface{}) error {
	if nil == p.PlatformData {
		p.PlatformData = NewOneOfDomainPlatformData()
	}
	e := p.PlatformData.SetValue(v)
	if nil == e {
		if nil == p.PlatformDataItemDiscriminator_ {
			p.PlatformDataItemDiscriminator_ = new(string)
		}
		*p.PlatformDataItemDiscriminator_ = *p.PlatformData.Discriminator
	}
	return e
}

/*
Audit trail information of domain such as created at and updated at etc.
*/
type DomainAuditTrail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User/Service which created this domain
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Domain registration date
	*/
	RegisteredAt *time.Time `json:"registeredAt,omitempty"`
	/*
	  User/Service which registered this domain to Nutanix central.
	*/
	RegisteredBy *string `json:"registeredBy,omitempty"`
}

func (p *DomainAuditTrail) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainAuditTrail

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

func (p *DomainAuditTrail) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainAuditTrail
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainAuditTrail()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.RegisteredAt != nil {
		p.RegisteredAt = known.RegisteredAt
	}
	if known.RegisteredBy != nil {
		p.RegisteredBy = known.RegisteredBy
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "registeredAt")
	delete(allFields, "registeredBy")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainAuditTrail() *DomainAuditTrail {
	p := new(DomainAuditTrail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainAuditTrail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Request payload to update state of the domain.
*/
type DomainChangeStateRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RegistrationState *DomainRegistrationState `json:"registrationState"`
}

func (p *DomainChangeStateRequest) MarshalJSON() ([]byte, error) {
	type DomainChangeStateRequestProxy DomainChangeStateRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainChangeStateRequestProxy
		RegistrationState *DomainRegistrationState `json:"registrationState,omitempty"`
	}{
		DomainChangeStateRequestProxy: (*DomainChangeStateRequestProxy)(p),
		RegistrationState:             p.RegistrationState,
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

func (p *DomainChangeStateRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainChangeStateRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainChangeStateRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "registrationState")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainChangeStateRequest() *DomainChangeStateRequest {
	p := new(DomainChangeStateRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainChangeStateRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain state change operation response
*/
type DomainChangeStateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *DomainChangeStateResponse) MarshalJSON() ([]byte, error) {
	type DomainChangeStateResponseProxy DomainChangeStateResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainChangeStateResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		DomainChangeStateResponseProxy: (*DomainChangeStateResponseProxy)(p),
		Message:                        p.Message,
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

func (p *DomainChangeStateResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainChangeStateResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainChangeStateResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainChangeStateResponse() *DomainChangeStateResponse {
	p := new(DomainChangeStateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainChangeStateResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain delete response data
*/
type DomainDeleteResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Delete error list
	*/
	ErrorList []string `json:"errorList,omitempty"`
	/*
	  Domain's status
	*/
	Status *string `json:"status,omitempty"`
	/*
	  Task UUID of async register task
	*/
	TaskUuid *string `json:"taskUuid,omitempty"`
}

func (p *DomainDeleteResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainDeleteResponse

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

func (p *DomainDeleteResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainDeleteResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainDeleteResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ErrorList != nil {
		p.ErrorList = known.ErrorList
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "errorList")
	delete(allFields, "status")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainDeleteResponse() *DomainDeleteResponse {
	p := new(DomainDeleteResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainDeleteResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This contains the basic information about the domain
*/
type DomainDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  DNS name of domain
	*/
	DnsName *string `json:"dnsName,omitempty"`
	/*
	  Fully qualified domain name of domain
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  Name of the domain
	*/
	Name *string `json:"name,omitempty"`
}

func (p *DomainDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainDetails

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

func (p *DomainDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DnsName != nil {
		p.DnsName = known.DnsName
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dnsName")
	delete(allFields, "fqdn")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainDetails() *DomainDetails {
	p := new(DomainDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain export response
*/
type DomainExportResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  File containing details about API credentials and tenant information for the domain.
	*/
	File *string `json:"file,omitempty"`
}

func (p *DomainExportResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainExportResponse

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

func (p *DomainExportResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainExportResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainExportResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.File != nil {
		p.File = known.File
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "file")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainExportResponse() *DomainExportResponse {
	p := new(DomainExportResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainExportResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DomainProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DomainAuditTrail *DomainAuditTrail `json:"domainAuditTrail,omitempty"`

	DomainDetails *DomainDetails `json:"domainDetails,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Whether domain is reachable from Nutanix central or not.
	*/
	IsReachable *bool `json:"isReachable,omitempty"`
	/*
	  API keyId Reference of domain
	*/
	KeyIdReference *string `json:"keyIdReference,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Location *Location `json:"location,omitempty"`

	PlatformDataItemDiscriminator_ *string `json:"$platformDataItemDiscriminator,omitempty"`
	/*
	  Platform specific Data.
	*/
	PlatformData *OneOfDomainProjectionPlatformData `json:"platformData,omitempty"`

	RegistrationState *DomainRegistrationState `json:"registrationState,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Communication between Nutanix Central and Prism central will be done through the use of a tunnel provider.
	*/
	TunnelProvider []string `json:"tunnelProvider,omitempty"`
}

func (p *DomainProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainProjection

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

func (p *DomainProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DomainAuditTrail != nil {
		p.DomainAuditTrail = known.DomainAuditTrail
	}
	if known.DomainDetails != nil {
		p.DomainDetails = known.DomainDetails
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsReachable != nil {
		p.IsReachable = known.IsReachable
	}
	if known.KeyIdReference != nil {
		p.KeyIdReference = known.KeyIdReference
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.PlatformDataItemDiscriminator_ != nil {
		p.PlatformDataItemDiscriminator_ = known.PlatformDataItemDiscriminator_
	}
	if known.PlatformData != nil {
		p.PlatformData = known.PlatformData
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TunnelProvider != nil {
		p.TunnelProvider = known.TunnelProvider
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domainAuditTrail")
	delete(allFields, "domainDetails")
	delete(allFields, "extId")
	delete(allFields, "isReachable")
	delete(allFields, "keyIdReference")
	delete(allFields, "links")
	delete(allFields, "location")
	delete(allFields, "$platformDataItemDiscriminator")
	delete(allFields, "platformData")
	delete(allFields, "registrationState")
	delete(allFields, "tenantId")
	delete(allFields, "tunnelProvider")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainProjection() *DomainProjection {
	p := new(DomainProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsReachable = new(bool)
	*p.IsReachable = false

	return p
}

/*
Domain register request payload
*/
type DomainRegisterRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  CSR for the domain register request
	*/
	Csr *string `json:"csr"`
}

func (p *DomainRegisterRequest) MarshalJSON() ([]byte, error) {
	type DomainRegisterRequestProxy DomainRegisterRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainRegisterRequestProxy
		Csr *string `json:"csr,omitempty"`
	}{
		DomainRegisterRequestProxy: (*DomainRegisterRequestProxy)(p),
		Csr:                        p.Csr,
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

func (p *DomainRegisterRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainRegisterRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainRegisterRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Csr != nil {
		p.Csr = known.Csr
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "csr")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainRegisterRequest() *DomainRegisterRequest {
	p := new(DomainRegisterRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainRegisterRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain register on cloud response
*/
type DomainRegisterResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client certificate value for domain registration to be used by ikat on pc
	*/
	ClientCertificate *string `json:"clientCertificate"`
	/*
	  Cloud IAM's URL. This is used by domain to establish federation between domain IAM and cloud IAM.
	*/
	CloudIAMURL *string `json:"cloudIAMURL"`
	/*
	  Ikat Hub Certificate chain value for domain registration to be used by ikat on domain
	*/
	HubCertChain *string `json:"hubCertChain"`
	/*
	  Ikat hub endpoint for tunnel connectivity.
	*/
	HubEndpoint *string `json:"hubEndpoint"`

	Location *Location `json:"location"`
	/*
	  The domain has a subdomain assigned to it. This is used by the domain to form the Ikat Hub address.
	*/
	SubDomainName *string `json:"subDomainName"`
}

func (p *DomainRegisterResponse) MarshalJSON() ([]byte, error) {
	type DomainRegisterResponseProxy DomainRegisterResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainRegisterResponseProxy
		ClientCertificate *string   `json:"clientCertificate,omitempty"`
		CloudIAMURL       *string   `json:"cloudIAMURL,omitempty"`
		HubCertChain      *string   `json:"hubCertChain,omitempty"`
		HubEndpoint       *string   `json:"hubEndpoint,omitempty"`
		Location          *Location `json:"location,omitempty"`
		SubDomainName     *string   `json:"subDomainName,omitempty"`
	}{
		DomainRegisterResponseProxy: (*DomainRegisterResponseProxy)(p),
		ClientCertificate:           p.ClientCertificate,
		CloudIAMURL:                 p.CloudIAMURL,
		HubCertChain:                p.HubCertChain,
		HubEndpoint:                 p.HubEndpoint,
		Location:                    p.Location,
		SubDomainName:               p.SubDomainName,
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

func (p *DomainRegisterResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainRegisterResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainRegisterResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientCertificate != nil {
		p.ClientCertificate = known.ClientCertificate
	}
	if known.CloudIAMURL != nil {
		p.CloudIAMURL = known.CloudIAMURL
	}
	if known.HubCertChain != nil {
		p.HubCertChain = known.HubCertChain
	}
	if known.HubEndpoint != nil {
		p.HubEndpoint = known.HubEndpoint
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.SubDomainName != nil {
		p.SubDomainName = known.SubDomainName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientCertificate")
	delete(allFields, "cloudIAMURL")
	delete(allFields, "hubCertChain")
	delete(allFields, "hubEndpoint")
	delete(allFields, "location")
	delete(allFields, "subDomainName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainRegisterResponse() *DomainRegisterResponse {
	p := new(DomainRegisterResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.DomainRegisterResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the state of domain registration.
*/
type DomainRegistrationState int

const (
	DOMAINREGISTRATIONSTATE_UNKNOWN       DomainRegistrationState = 0
	DOMAINREGISTRATIONSTATE_REDACTED      DomainRegistrationState = 1
	DOMAINREGISTRATIONSTATE_UNREGISTERED  DomainRegistrationState = 2
	DOMAINREGISTRATIONSTATE_REGISTERING   DomainRegistrationState = 3
	DOMAINREGISTRATIONSTATE_REGISTERED    DomainRegistrationState = 4
	DOMAINREGISTRATIONSTATE_UNREGISTERING DomainRegistrationState = 5
	DOMAINREGISTRATIONSTATE_DRAFT         DomainRegistrationState = 6
	DOMAINREGISTRATIONSTATE_ERROR         DomainRegistrationState = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainRegistrationState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"DRAFT",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainRegistrationState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"DRAFT",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainRegistrationState) index(name string) DomainRegistrationState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"DRAFT",
		"ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return DomainRegistrationState(idx)
		}
	}
	return DOMAINREGISTRATIONSTATE_UNKNOWN
}

func (e *DomainRegistrationState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DomainRegistrationState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DomainRegistrationState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DomainRegistrationState) Ref() *DomainRegistrationState {
	return &e
}

/*
Location information.
*/
type Location struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The location coordinates, that is latitude and longitude.
	*/
	Coordinates *string `json:"coordinates,omitempty"`
	/*
	  Location Country.
	*/
	Country *string `json:"country"`
	/*
	  Location name
	*/
	Name *string `json:"name"`
	/*
	  Location state.
	*/
	SubCountry *string `json:"subCountry,omitempty"`
}

func (p *Location) MarshalJSON() ([]byte, error) {
	type LocationProxy Location

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LocationProxy
		Country *string `json:"country,omitempty"`
		Name    *string `json:"name,omitempty"`
	}{
		LocationProxy: (*LocationProxy)(p),
		Country:       p.Country,
		Name:          p.Name,
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

func (p *Location) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Location
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Coordinates != nil {
		p.Coordinates = known.Coordinates
	}
	if known.Country != nil {
		p.Country = known.Country
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.SubCountry != nil {
		p.SubCountry = known.SubCountry
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "coordinates")
	delete(allFields, "country")
	delete(allFields, "name")
	delete(allFields, "subCountry")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocation() *Location {
	p := new(Location)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.Location"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Prism Central information such as cluster count, pc version etc.
*/
type PrismCentral struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of clusters in the domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  PC Version of the domain
	*/
	Version *string `json:"version,omitempty"`
}

func (p *PrismCentral) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PrismCentral

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

func (p *PrismCentral) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PrismCentral
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPrismCentral()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterCount")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPrismCentral() *PrismCentral {
	p := new(PrismCentral)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.resources.PrismCentral"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfDomainPlatformData struct {
	Discriminator *string       `json:"-"`
	ObjectType_   *string       `json:"-"`
	oneOfType0    *PrismCentral `json:"-"`
}

func NewOneOfDomainPlatformData() *OneOfDomainPlatformData {
	p := new(OneOfDomainPlatformData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDomainPlatformData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDomainPlatformData is nil"))
	}
	switch v.(type) {
	case PrismCentral:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PrismCentral)
		}
		*p.oneOfType0 = v.(PrismCentral)
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

func (p *OneOfDomainPlatformData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDomainPlatformData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(PrismCentral)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "multidomain.v4.resources.PrismCentral" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PrismCentral)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainPlatformData"))
}

func (p *OneOfDomainPlatformData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDomainPlatformData")
}

type OneOfDomainProjectionPlatformData struct {
	Discriminator *string       `json:"-"`
	ObjectType_   *string       `json:"-"`
	oneOfType0    *PrismCentral `json:"-"`
}

func NewOneOfDomainProjectionPlatformData() *OneOfDomainProjectionPlatformData {
	p := new(OneOfDomainProjectionPlatformData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDomainProjectionPlatformData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDomainProjectionPlatformData is nil"))
	}
	switch v.(type) {
	case PrismCentral:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PrismCentral)
		}
		*p.oneOfType0 = v.(PrismCentral)
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

func (p *OneOfDomainProjectionPlatformData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDomainProjectionPlatformData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(PrismCentral)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "multidomain.v4.resources.PrismCentral" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PrismCentral)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainProjectionPlatformData"))
}

func (p *OneOfDomainProjectionPlatformData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDomainProjectionPlatformData")
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
