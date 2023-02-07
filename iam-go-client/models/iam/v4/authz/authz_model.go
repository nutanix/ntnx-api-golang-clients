/*
 * Generated file models/iam/v4/authz/authz_model.go.
 *
 * Product version: 4.0.2-alpha-1
 *
 * Part of the Nutanix Iam Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Access policies, roles and operations
*/
package authz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/error"
	"time"
)

/**
An access policy to define who can perform what operation on which object
*/
type AccessPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessPolicyType *AccessPolicyType `json:"accessPolicyType,omitempty"`
	/**
	  Client that created the access policy
	*/
	ClientName *string `json:"clientName,omitempty"`
	/**
	  User or service name that created the access policy
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  The creation time of the access policy
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  Description of the access policy
	*/
	Description *string `json:"description,omitempty"`
	/**
	  The display name for the access policy
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  The identities for which the access policy is created
	*/
	Identities []Filter `json:"identities,omitempty"`
	/**
	  Flag identifying if the access policy is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/**
	  The time when the access policy was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  The objects being qualified by the access policy
	*/
	Objects []Filter `json:"objects,omitempty"`
	/**
	  The role associated with the access policy
	*/
	Role *string `json:"role,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAccessPolicy() *AccessPolicy {
	p := new(AccessPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AccessPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.AccessPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/**
Type of access policy
*/
type AccessPolicyType int

const (
	ACCESSPOLICYTYPE_UNKNOWN                         AccessPolicyType = 0
	ACCESSPOLICYTYPE_REDACTED                        AccessPolicyType = 1
	ACCESSPOLICYTYPE_USER_DEFINED                    AccessPolicyType = 2
	ACCESSPOLICYTYPE_SERVICE_DEFINED                 AccessPolicyType = 3
	ACCESSPOLICYTYPE_PREDEFINED_READ_ONLY            AccessPolicyType = 4
	ACCESSPOLICYTYPE_PREDEFINED_UPDATE_IDENTITY_ONLY AccessPolicyType = 5
	ACCESSPOLICYTYPE_SERVICE_DEFINED_READ_ONLY       AccessPolicyType = 6
)

// returns the name of the enum given an ordinal number
func (e *AccessPolicyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER_DEFINED",
		"SERVICE_DEFINED",
		"PREDEFINED_READ_ONLY",
		"PREDEFINED_UPDATE_IDENTITY_ONLY",
		"SERVICE_DEFINED_READ_ONLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *AccessPolicyType) index(name string) AccessPolicyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER_DEFINED",
		"SERVICE_DEFINED",
		"PREDEFINED_READ_ONLY",
		"PREDEFINED_UPDATE_IDENTITY_ONLY",
		"SERVICE_DEFINED_READ_ONLY",
	}
	for idx := range names {
		if names[idx] == name {
			return AccessPolicyType(idx)
		}
	}
	return ACCESSPOLICYTYPE_UNKNOWN
}

func (e *AccessPolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AccessPolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AccessPolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AccessPolicyType) Ref() *AccessPolicyType {
	return &e
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/accessPolicies Post operation
*/
type CreateAccessPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateAccessPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateAccessPolicyApiResponse() *CreateAccessPolicyApiResponse {
	p := new(CreateAccessPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateAccessPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.CreateAccessPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateAccessPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateAccessPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateAccessPolicyApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/roles Post operation
*/
type CreateRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateRoleApiResponse() *CreateRoleApiResponse {
	p := new(CreateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.CreateRoleApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateRoleApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/accessPolicies/{extId} Delete operation
*/
type DeleteAccessPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteAccessPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteAccessPolicyApiResponse() *DeleteAccessPolicyApiResponse {
	p := new(DeleteAccessPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteAccessPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.DeleteAccessPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteAccessPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteAccessPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteAccessPolicyApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/roles/{extId} Delete operation
*/
type DeleteRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRoleApiResponse() *DeleteRoleApiResponse {
	p := new(DeleteRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.DeleteRoleApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRoleApiResponseData()
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

type Filter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func NewFilter() *Filter {
	p := new(Filter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Filter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.Filter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/accessPolicies Get operation
*/
type ListAccessPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAccessPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAccessPolicyApiResponse() *ListAccessPolicyApiResponse {
	p := new(ListAccessPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListAccessPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ListAccessPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAccessPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAccessPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAccessPolicyApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/operations Get operation
*/
type ListOperationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListOperationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListOperationApiResponse() *ListOperationApiResponse {
	p := new(ListOperationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListOperationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ListOperationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListOperationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListOperationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListOperationApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/roles Get operation
*/
type ListRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRoleApiResponse() *ListRoleApiResponse {
	p := new(ListRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ListRoleApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRoleApiResponseData()
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

type Operation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Client that created the operation
	*/
	ClientName *string `json:"clientName,omitempty"`
	/**
	  The creation time of the operation
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  Description of the operation
	*/
	Description *string `json:"description,omitempty"`
	/**
	  Name of the operation
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  The time when the operation was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  Type of object associated with this operation
	*/
	ObjectType *string `json:"objectType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewOperation() *Operation {
	p := new(Operation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Operation"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.Operation"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
A role to group the operations and/or sub roles
*/
type Role struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of accessible clients for the role
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/**
	  List of Accessible ObjectTypes for the role
	*/
	AccessibleObjectTypes []string `json:"accessibleObjectTypes,omitempty"`
	/**
	  Client that created the role
	*/
	ClientName *string `json:"clientName,omitempty"`
	/**
	  User or service name that created the role
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  The creation time of the role
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  Description of the role
	*/
	Description *string `json:"description,omitempty"`
	/**
	  The display name for the role
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Flag identifying if the role is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/**
	  The time when the role was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  List of operations for the role
	*/
	Operations []string `json:"operations,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRole() *Role {
	p := new(Role)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Role"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.Role"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/**
Status type for response
*/
type StatusType int

const (
	STATUSTYPE_UNKNOWN        StatusType = 0
	STATUSTYPE_REDACTED       StatusType = 1
	STATUSTYPE_NOERROR        StatusType = 2
	STATUSTYPE_APPERROR       StatusType = 3
	STATUSTYPE_CANCELED       StatusType = 4
	STATUSTYPE_RETRY          StatusType = 5
	STATUSTYPE_TIMEOUT        StatusType = 6
	STATUSTYPE_AUTHORIZED     StatusType = 7
	STATUSTYPE_BADREQUEST     StatusType = 8
	STATUSTYPE_NOTAUTHORIZED  StatusType = 9
	STATUSTYPE_PARTIALSUCCESS StatusType = 10
)

// returns the name of the enum given an ordinal number
func (e *StatusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NoError",
		"AppError",
		"Canceled",
		"Retry",
		"Timeout",
		"Authorized",
		"BadRequest",
		"NotAuthorized",
		"PartialSuccess",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *StatusType) index(name string) StatusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NoError",
		"AppError",
		"Canceled",
		"Retry",
		"Timeout",
		"Authorized",
		"BadRequest",
		"NotAuthorized",
		"PartialSuccess",
	}
	for idx := range names {
		if names[idx] == name {
			return StatusType(idx)
		}
	}
	return STATUSTYPE_UNKNOWN
}

func (e *StatusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StatusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StatusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StatusType) Ref() *StatusType {
	return &e
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/accessPolicies/{extId} Put operation
*/
type UpdateAccessPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateAccessPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateAccessPolicyApiResponse() *UpdateAccessPolicyApiResponse {
	p := new(UpdateAccessPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateAccessPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.UpdateAccessPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateAccessPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateAccessPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateAccessPolicyApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/roles/{extId} Put operation
*/
type UpdateRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRoleApiResponse() *UpdateRoleApiResponse {
	p := new(UpdateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.UpdateRoleApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateRoleApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/accessPolicies/{extId} Get operation
*/
type ViewAccessPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewAccessPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewAccessPolicyApiResponse() *ViewAccessPolicyApiResponse {
	p := new(ViewAccessPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewAccessPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ViewAccessPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewAccessPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewAccessPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewAccessPolicyApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/operations/{extId} Get operation
*/
type ViewOperationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewOperationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewOperationApiResponse() *ViewOperationApiResponse {
	p := new(ViewOperationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewOperationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ViewOperationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewOperationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewOperationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewOperationApiResponseData()
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authz/roles/{extId} Get operation
*/
type ViewRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewRoleApiResponse() *ViewRoleApiResponse {
	p := new(ViewRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authz.ViewRoleApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewRoleApiResponseData()
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

type OneOfDeleteRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteRoleApiResponseData() *OneOfDeleteRoleApiResponseData {
	p := new(OneOfDeleteRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRoleApiResponseData is nil"))
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

func (p *OneOfDeleteRoleApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRoleApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRoleApiResponseData"))
}

func (p *OneOfDeleteRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRoleApiResponseData")
}

type OneOfViewOperationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Operation             `json:"-"`
}

func NewOneOfViewOperationApiResponseData() *OneOfViewOperationApiResponseData {
	p := new(OneOfViewOperationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewOperationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewOperationApiResponseData is nil"))
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
	case Operation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Operation)
		}
		*p.oneOfType0 = v.(Operation)
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

func (p *OneOfViewOperationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfViewOperationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Operation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.Operation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Operation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewOperationApiResponseData"))
}

func (p *OneOfViewOperationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfViewOperationApiResponseData")
}

type OneOfDeleteAccessPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteAccessPolicyApiResponseData() *OneOfDeleteAccessPolicyApiResponseData {
	p := new(OneOfDeleteAccessPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteAccessPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteAccessPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteAccessPolicyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteAccessPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAccessPolicyApiResponseData"))
}

func (p *OneOfDeleteAccessPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteAccessPolicyApiResponseData")
}

type OneOfCreateAccessPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import2.Message       `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateAccessPolicyApiResponseData() *OneOfCreateAccessPolicyApiResponseData {
	p := new(OneOfCreateAccessPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateAccessPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateAccessPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.Message)
		}
		*p.oneOfType0 = v.(import2.Message)
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

func (p *OneOfCreateAccessPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateAccessPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import2.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.Message)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAccessPolicyApiResponseData"))
}

func (p *OneOfCreateAccessPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateAccessPolicyApiResponseData")
}

type OneOfListRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Role                 `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListRoleApiResponseData() *OneOfListRoleApiResponseData {
	p := new(OneOfListRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRoleApiResponseData is nil"))
	}
	switch v.(type) {
	case []Role:
		p.oneOfType0 = v.([]Role)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.Role>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.Role>"
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

func (p *OneOfListRoleApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.Role>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListRoleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Role)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.Role" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.Role>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.Role>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRoleApiResponseData"))
}

func (p *OneOfListRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.Role>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListRoleApiResponseData")
}

type OneOfViewRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2    *interface{}           `json:"-"`
	oneOfType0    *Role                  `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfViewRoleApiResponseData() *OneOfViewRoleApiResponseData {
	p := new(OneOfViewRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewRoleApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(interface{})
		}
		*p.oneOfType2 = nil
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
	case Role:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Role)
		}
		*p.oneOfType0 = v.(Role)
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

func (p *OneOfViewRoleApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewRoleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if nil == *vOneOfType2 {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(interface{})
			}
			*p.oneOfType2 = nil
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
	vOneOfType0 := new(Role)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.Role" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Role)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewRoleApiResponseData"))
}

func (p *OneOfViewRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewRoleApiResponseData")
}

type OneOfCreateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import2.Message       `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateRoleApiResponseData() *OneOfCreateRoleApiResponseData {
	p := new(OneOfCreateRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateRoleApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.Message)
		}
		*p.oneOfType0 = v.(import2.Message)
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

func (p *OneOfCreateRoleApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateRoleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import2.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.Message)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRoleApiResponseData"))
}

func (p *OneOfCreateRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRoleApiResponseData")
}

type OneOfUpdateAccessPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import2.Message       `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateAccessPolicyApiResponseData() *OneOfUpdateAccessPolicyApiResponseData {
	p := new(OneOfUpdateAccessPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateAccessPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateAccessPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.Message)
		}
		*p.oneOfType0 = v.(import2.Message)
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

func (p *OneOfUpdateAccessPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateAccessPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import2.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.Message)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAccessPolicyApiResponseData"))
}

func (p *OneOfUpdateAccessPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateAccessPolicyApiResponseData")
}

type OneOfUpdateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import2.Message       `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateRoleApiResponseData() *OneOfUpdateRoleApiResponseData {
	p := new(OneOfUpdateRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRoleApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.Message)
		}
		*p.oneOfType0 = v.(import2.Message)
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

func (p *OneOfUpdateRoleApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateRoleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import2.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.Message)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRoleApiResponseData"))
}

func (p *OneOfUpdateRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRoleApiResponseData")
}

type OneOfListAccessPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []AccessPolicy         `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListAccessPolicyApiResponseData() *OneOfListAccessPolicyApiResponseData {
	p := new(OneOfListAccessPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAccessPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAccessPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case []AccessPolicy:
		p.oneOfType0 = v.([]AccessPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.AccessPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.AccessPolicy>"
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

func (p *OneOfListAccessPolicyApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.AccessPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListAccessPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]AccessPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.AccessPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.AccessPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.AccessPolicy>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAccessPolicyApiResponseData"))
}

func (p *OneOfListAccessPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.AccessPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListAccessPolicyApiResponseData")
}

type OneOfListOperationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Operation            `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListOperationApiResponseData() *OneOfListOperationApiResponseData {
	p := new(OneOfListOperationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListOperationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListOperationApiResponseData is nil"))
	}
	switch v.(type) {
	case []Operation:
		p.oneOfType0 = v.([]Operation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.Operation>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.Operation>"
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

func (p *OneOfListOperationApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.Operation>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListOperationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Operation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.Operation" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.Operation>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.Operation>"
			return nil

		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListOperationApiResponseData"))
}

func (p *OneOfListOperationApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.Operation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListOperationApiResponseData")
}

type OneOfViewAccessPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *AccessPolicy          `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfViewAccessPolicyApiResponseData() *OneOfViewAccessPolicyApiResponseData {
	p := new(OneOfViewAccessPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewAccessPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewAccessPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case AccessPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AccessPolicy)
		}
		*p.oneOfType0 = v.(AccessPolicy)
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

func (p *OneOfViewAccessPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewAccessPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AccessPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.AccessPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AccessPolicy)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewAccessPolicyApiResponseData"))
}

func (p *OneOfViewAccessPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewAccessPolicyApiResponseData")
}
