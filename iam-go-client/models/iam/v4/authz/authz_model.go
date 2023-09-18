/*
 * Generated file models/iam/v4/authz/authz_model.go.
 *
 * Product version: 4.0.1-beta-1
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
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/error"
	"time"
)

/*
Version of the API for the provided associated endpoint
*/
type ApiVersion int

const (
	APIVERSION_UNKNOWN  ApiVersion = 0
	APIVERSION_REDACTED ApiVersion = 1
	APIVERSION_V3       ApiVersion = 2
	APIVERSION_V4       ApiVersion = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApiVersion) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"V3",
		"V4",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ApiVersion) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"V3",
		"V4",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ApiVersion) index(name string) ApiVersion {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"V3",
		"V4",
	}
	for idx := range names {
		if names[idx] == name {
			return ApiVersion(idx)
		}
	}
	return APIVERSION_UNKNOWN
}

func (e *ApiVersion) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApiVersion:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApiVersion) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApiVersion) Ref() *ApiVersion {
	return &e
}

/*
The object which defines the API endpoint, API version and HTTP method for the API endpoint
*/
type AssociatedEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApiVersion *ApiVersion `json:"apiVersion,omitempty"`
	/*
	  Endpoint URL
	*/
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	HttpMethod *HttpMethod `json:"httpMethod,omitempty"`
}

func NewAssociatedEndpoint() *AssociatedEndpoint {
	p := new(AssociatedEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AssociatedEndpoint"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AssociatedEndpoint"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of attributes supported for access control on this entity
*/
type AttributeEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of attributes supported for access control on this entity.
	*/
	AttributeValues []string `json:"attributeValues,omitempty"`
	/*
	  UI display name of the entity
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  List of supported operators for this entity
	*/
	SupportedOperators []EntityOperators `json:"supportedOperators,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAttributeEntity() *AttributeEntity {
	p := new(AttributeEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AttributeEntity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AttributeEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An authorization policy to define who can perform what operation on which entity
*/
type AuthorizationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuthorizationPolicyType *AuthorizationPolicyType `json:"authorizationPolicyType,omitempty"`
	/*
	  Client that created the entity
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the authorization policy
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the authorization policy
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the authorization policy
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The identities for which the authorization policy is created
	*/
	Identities []IdentityFilter `json:"identities,omitempty"`
	/*
	  Flag identifying if the authorization policy is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the authorization policy was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAuthorizationPolicy() *AuthorizationPolicy {
	p := new(AuthorizationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizationPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

type AuthorizationPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AuthorizationPolicyType *AuthorizationPolicyType `json:"authorizationPolicyType,omitempty"`
	/*
	  Client that created the entity
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the authorization policy
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the authorization policy
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the authorization policy
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The identities for which the authorization policy is created
	*/
	Identities []IdentityFilter `json:"identities,omitempty"`
	/*
	  Flag identifying if the authorization policy is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the authorization policy was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy
	*/
	Role *string `json:"role,omitempty"`

	RoleProjection *RoleProjection `json:"roleProjection,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewAuthorizationPolicyProjection() *AuthorizationPolicyProjection {
	p := new(AuthorizationPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizationPolicyProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/*
Type of authorization policy
*/
type AuthorizationPolicyType int

const (
	AUTHORIZATIONPOLICYTYPE_UNKNOWN                         AuthorizationPolicyType = 0
	AUTHORIZATIONPOLICYTYPE_REDACTED                        AuthorizationPolicyType = 1
	AUTHORIZATIONPOLICYTYPE_USER_DEFINED                    AuthorizationPolicyType = 2
	AUTHORIZATIONPOLICYTYPE_SERVICE_DEFINED                 AuthorizationPolicyType = 3
	AUTHORIZATIONPOLICYTYPE_PREDEFINED_READ_ONLY            AuthorizationPolicyType = 4
	AUTHORIZATIONPOLICYTYPE_PREDEFINED_UPDATE_IDENTITY_ONLY AuthorizationPolicyType = 5
	AUTHORIZATIONPOLICYTYPE_SERVICE_DEFINED_READ_ONLY       AuthorizationPolicyType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthorizationPolicyType) name(index int) string {
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

// Returns the name of the enum
func (e AuthorizationPolicyType) GetName() string {
	index := int(e)
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

// Returns the enum type given a string value
func (e *AuthorizationPolicyType) index(name string) AuthorizationPolicyType {
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
			return AuthorizationPolicyType(idx)
		}
	}
	return AUTHORIZATIONPOLICYTYPE_UNKNOWN
}

func (e *AuthorizationPolicyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthorizationPolicyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthorizationPolicyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthorizationPolicyType) Ref() *AuthorizationPolicyType {
	return &e
}

/*
Request body for authorization request
*/
type AuthorizationRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Identity *IdentityObject `json:"identity"`
	/*
	  List of requests to be authorized
	*/
	RequestList []RequestObject `json:"requestList"`
}

func (p *AuthorizationRequest) MarshalJSON() ([]byte, error) {
	type AuthorizationRequestProxy AuthorizationRequest
	return json.Marshal(struct {
		*AuthorizationRequestProxy
		Identity    *IdentityObject `json:"identity,omitempty"`
		RequestList []RequestObject `json:"requestList,omitempty"`
	}{
		AuthorizationRequestProxy: (*AuthorizationRequestProxy)(p),
		Identity:                  p.Identity,
		RequestList:               p.RequestList,
	})
}

func NewAuthorizationRequest() *AuthorizationRequest {
	p := new(AuthorizationRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizationRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity on which the action is being performed
*/
type AuthorizationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of responses for the authorization request
	*/
	Response []AuthorizationResponseObject `json:"response,omitempty"`
}

func NewAuthorizationResponse() *AuthorizationResponse {
	p := new(AuthorizationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizationResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Authorization enforcement response
*/
type AuthorizationResponseObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  True if request was authorized, otherwise False
	*/
	Authorized *interface{} `json:"authorized,omitempty"`
	/*
	  Filter expression
	*/
	Filters *interface{} `json:"filters,omitempty"`

	RequestId *string `json:"requestId,omitempty"`

	ResponseType *ResponseType `json:"responseType,omitempty"`
}

func NewAuthorizationResponseObject() *AuthorizationResponseObject {
	p := new(AuthorizationResponseObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationResponseObject"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizationResponseObject"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/$actions/authorize Post operation
*/
type AuthorizeApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAuthorizeApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAuthorizeApiResponse() *AuthorizeApiResponse {
	p := new(AuthorizeApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizeApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizeApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AuthorizeApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AuthorizeApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAuthorizeApiResponseData()
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
Entity on which the action is being performed
*/
type AuthorizeEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Map of attributes for the entity
	*/
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	/*
	  operation being performed on the entity
	*/
	Operation *string `json:"operation"`
	/*
	  Type of the entity
	*/
	Type *string `json:"type"`
}

func (p *AuthorizeEntity) MarshalJSON() ([]byte, error) {
	type AuthorizeEntityProxy AuthorizeEntity
	return json.Marshal(struct {
		*AuthorizeEntityProxy
		Operation *string `json:"operation,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{
		AuthorizeEntityProxy: (*AuthorizeEntityProxy)(p),
		Operation:            p.Operation,
		Type:                 p.Type,
	})
}

func NewAuthorizeEntity() *AuthorizeEntity {
	p := new(AuthorizeEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizeEntity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.AuthorizeEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the client registered
*/
type Client struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or Service Name that created the Client
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the Client
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Deployment List related to the Client
	*/
	DeploymentList []string `json:"deploymentList,omitempty"`
	/*
	  Description of the client
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI name of the Client
	*/
	DisplayName *string `json:"displayName"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the Client was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the client
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Client) MarshalJSON() ([]byte, error) {
	type ClientProxy Client
	return json.Marshal(struct {
		*ClientProxy
		DisplayName *string `json:"displayName,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{
		ClientProxy: (*ClientProxy)(p),
		DisplayName: p.DisplayName,
		Name:        p.Name,
	})
}

func NewClient() *Client {
	p := new(Client)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Client"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.Client"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/authorization-policies Post operation
*/
type CreateAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateAuthorizationPolicyApiResponse() *CreateAuthorizationPolicyApiResponse {
	p := new(CreateAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.CreateAuthorizationPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/client/register Post operation
*/
type CreateClientApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateClientApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateClientApiResponse() *CreateClientApiResponse {
	p := new(CreateClientApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateClientApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.CreateClientApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateClientApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateClientApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateClientApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/roles Post operation
*/
type CreateRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateRoleApiResponse() *CreateRoleApiResponse {
	p := new(CreateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.CreateRoleApiResponse"}
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

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/authorization-policies/{extId} Delete operation
*/
type DeleteAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteAuthorizationPolicyApiResponse() *DeleteAuthorizationPolicyApiResponse {
	p := new(DeleteAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.DeleteAuthorizationPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/roles/{extId} Delete operation
*/
type DeleteRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRoleApiResponse() *DeleteRoleApiResponse {
	p := new(DeleteRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.DeleteRoleApiResponse"}
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

/*
Information of the entity
*/
type Entity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AttributeList []AttributeEntity `json:"attributeList,omitempty"`
	/*
	  Name of client the Entity belongs to
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the Entity
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the Entity
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the entity
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI display name of the entity
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the Entity was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the entity
	*/
	Name *string `json:"name,omitempty"`
	/*
	  URL provided by the client to search the entities
	*/
	SearchURL *string `json:"searchURL,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Unique Id of a client entity
	*/
	UniqueId *string `json:"uniqueId,omitempty"`
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Entity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.Entity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EntityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func NewEntityFilter() *EntityFilter {
	p := new(EntityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.EntityFilter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.EntityFilter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of authorization policy
*/
type EntityOperators int

const (
	ENTITYOPERATORS_UNKNOWN  EntityOperators = 0
	ENTITYOPERATORS_REDACTED EntityOperators = 1
	ENTITYOPERATORS_ALLOF    EntityOperators = 2
	ENTITYOPERATORS_ANYOF    EntityOperators = 3
	ENTITYOPERATORS_EQ       EntityOperators = 4
	ENTITYOPERATORS_NONEOF   EntityOperators = 5
	ENTITYOPERATORS_NOTEQ    EntityOperators = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EntityOperators) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOF",
		"ANYOF",
		"EQ",
		"NONEOF",
		"NOTEQ",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EntityOperators) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOF",
		"ANYOF",
		"EQ",
		"NONEOF",
		"NOTEQ",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EntityOperators) index(name string) EntityOperators {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ALLOF",
		"ANYOF",
		"EQ",
		"NONEOF",
		"NOTEQ",
	}
	for idx := range names {
		if names[idx] == name {
			return EntityOperators(idx)
		}
	}
	return ENTITYOPERATORS_UNKNOWN
}

func (e *EntityOperators) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EntityOperators:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EntityOperators) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EntityOperators) Ref() *EntityOperators {
	return &e
}

/*
HTTP method for the provided associated endpoint
*/
type HttpMethod int

const (
	HTTPMETHOD_UNKNOWN  HttpMethod = 0
	HTTPMETHOD_REDACTED HttpMethod = 1
	HTTPMETHOD_POST     HttpMethod = 2
	HTTPMETHOD_GET      HttpMethod = 3
	HTTPMETHOD_PUT      HttpMethod = 4
	HTTPMETHOD_PATCH    HttpMethod = 5
	HTTPMETHOD_DELETE   HttpMethod = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HttpMethod) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POST",
		"GET",
		"PUT",
		"PATCH",
		"DELETE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HttpMethod) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POST",
		"GET",
		"PUT",
		"PATCH",
		"DELETE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HttpMethod) index(name string) HttpMethod {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"POST",
		"GET",
		"PUT",
		"PATCH",
		"DELETE",
	}
	for idx := range names {
		if names[idx] == name {
			return HttpMethod(idx)
		}
	}
	return HTTPMETHOD_UNKNOWN
}

func (e *HttpMethod) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HttpMethod:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HttpMethod) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HttpMethod) Ref() *HttpMethod {
	return &e
}

type IdentityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
}

func NewIdentityFilter() *IdentityFilter {
	p := new(IdentityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.IdentityFilter"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.IdentityFilter"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Object defining the attributes of Identity (who)
*/
type IdentityObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Map of attributes for the Identity
	*/
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	IdentityType *IdentityType `json:"identityType,omitempty"`
}

func NewIdentityObject() *IdentityObject {
	p := new(IdentityObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.IdentityObject"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.IdentityObject"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of identity object
*/
type IdentityType int

const (
	IDENTITYTYPE_UNKNOWN        IdentityType = 0
	IDENTITYTYPE_REDACTED       IdentityType = 1
	IDENTITYTYPE_USER           IdentityType = 2
	IDENTITYTYPE_SERVICEACCOUNT IdentityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IdentityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"User",
		"ServiceAccount",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e IdentityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"User",
		"ServiceAccount",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *IdentityType) index(name string) IdentityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"User",
		"ServiceAccount",
	}
	for idx := range names {
		if names[idx] == name {
			return IdentityType(idx)
		}
	}
	return IDENTITYTYPE_UNKNOWN
}

func (e *IdentityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IdentityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IdentityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IdentityType) Ref() *IdentityType {
	return &e
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/authorization-policies Get operation
*/
type ListAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListAuthorizationPolicyApiResponse() *ListAuthorizationPolicyApiResponse {
	p := new(ListAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ListAuthorizationPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/clients Get operation
*/
type ListClientApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListClientApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListClientApiResponse() *ListClientApiResponse {
	p := new(ListClientApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListClientApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ListClientApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListClientApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListClientApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListClientApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/entities Get operation
*/
type ListEntityApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEntityApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListEntityApiResponse() *ListEntityApiResponse {
	p := new(ListEntityApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListEntityApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ListEntityApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEntityApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEntityApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEntityApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/operations Get operation
*/
type ListOperationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListOperationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListOperationApiResponse() *ListOperationApiResponse {
	p := new(ListOperationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListOperationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ListOperationApiResponse"}
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

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/roles Get operation
*/
type ListRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListRoleApiResponse() *ListRoleApiResponse {
	p := new(ListRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ListRoleApiResponse"}
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
	/*
	  List of associated endpoint objects for the operation
	*/
	AssociatedEndpointList []AssociatedEndpoint `json:"associatedEndpointList,omitempty"`
	/*
	  Client that created the entity
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The creation time of the operation
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the operation
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of the operation
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Type of entity associated with this operation
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the operation was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  List of related operations. These are the operations which might need to be given access to, along with the current operation, for certain workflows to succeed
	*/
	RelatedOperationList []string `json:"relatedOperationList,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewOperation() *Operation {
	p := new(Operation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Operation"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.Operation"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The operation type. Currently we support INTERNAL, EXTERNAL and SYSTEM_DEFINED_ONLY
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN             OperationType = 0
	OPERATIONTYPE_REDACTED            OperationType = 1
	OPERATIONTYPE_INTERNAL            OperationType = 2
	OPERATIONTYPE_SYSTEM_DEFINED_ONLY OperationType = 3
	OPERATIONTYPE_EXTERNAL            OperationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INTERNAL",
		"SYSTEM_DEFINED_ONLY",
		"EXTERNAL",
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
		"INTERNAL",
		"SYSTEM_DEFINED_ONLY",
		"EXTERNAL",
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
		"INTERNAL",
		"SYSTEM_DEFINED_ONLY",
		"EXTERNAL",
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
The incoming request
*/
type RequestObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Entity *AuthorizeEntity `json:"entity"`
	/*
	  Unique qualifier for the request
	*/
	RequestID *string `json:"requestID"`

	ResponseType *ResponseType `json:"responseType,omitempty"`
}

func (p *RequestObject) MarshalJSON() ([]byte, error) {
	type RequestObjectProxy RequestObject
	return json.Marshal(struct {
		*RequestObjectProxy
		Entity    *AuthorizeEntity `json:"entity,omitempty"`
		RequestID *string          `json:"requestID,omitempty"`
	}{
		RequestObjectProxy: (*RequestObjectProxy)(p),
		Entity:             p.Entity,
		RequestID:          p.RequestID,
	})
}

func NewRequestObject() *RequestObject {
	p := new(RequestObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.RequestObject"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.RequestObject"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of response expected
*/
type ResponseType int

const (
	RESPONSETYPE_UNKNOWN         ResponseType = 0
	RESPONSETYPE_REDACTED        ResponseType = 1
	RESPONSETYPE_BOOLEANRESPONSE ResponseType = 2
	RESPONSETYPE_FILTERRESPONSE  ResponseType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResponseType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BooleanResponse",
		"FilterResponse",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ResponseType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BooleanResponse",
		"FilterResponse",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ResponseType) index(name string) ResponseType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BooleanResponse",
		"FilterResponse",
	}
	for idx := range names {
		if names[idx] == name {
			return ResponseType(idx)
		}
	}
	return RESPONSETYPE_UNKNOWN
}

func (e *ResponseType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ResponseType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ResponseType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ResponseType) Ref() *ResponseType {
	return &e
}

/*
A role to group the operations
*/
type Role struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Accessible Clients for the role
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/*
	  List of Accessible EntityTypes for the role
	*/
	AccessibleEntityTypes []string `json:"accessibleEntityTypes,omitempty"`
	/*
	  Number of user groups assigned to given role
	*/
	AssignedUserGroupsCount *int64 `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to given role
	*/
	AssignedUsersCount *int64 `json:"assignedUsersCount,omitempty"`
	/*
	  Client that created the entity
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the role
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the role
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the role
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the role
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the role is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the role was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  List of operations for the role
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRole() *Role {
	p := new(Role)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Role"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.Role"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

type RoleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Accessible Clients for the role
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/*
	  List of Accessible EntityTypes for the role
	*/
	AccessibleEntityTypes []string `json:"accessibleEntityTypes,omitempty"`
	/*
	  Number of user groups assigned to given role
	*/
	AssignedUserGroupsCount *int64 `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to given role
	*/
	AssignedUsersCount *int64 `json:"assignedUsersCount,omitempty"`
	/*
	  Client that created the entity
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the role
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the role
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the role
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the role
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the role is system defined or not
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the role was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  List of operations for the role
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRoleProjection() *RoleProjection {
	p := new(RoleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.RoleProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.RoleProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/*
Tenant details for tenant to be onboarded to IAM
*/
type Tenant struct {
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
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of tenant for which configuration is being set up
	*/
	TenantName *string `json:"tenantName,omitempty"`
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Tenant"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.Tenant"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/tenant Post operation
*/
type TenantConfigurationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTenantConfigurationResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTenantConfigurationResponse() *TenantConfigurationResponse {
	p := new(TenantConfigurationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.TenantConfigurationResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.TenantConfigurationResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TenantConfigurationResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TenantConfigurationResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTenantConfigurationResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/authorization-policies/{extId} Put operation
*/
type UpdateAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateAuthorizationPolicyApiResponse() *UpdateAuthorizationPolicyApiResponse {
	p := new(UpdateAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.UpdateAuthorizationPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/clients/{extId} Put operation
*/
type UpdateClientApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateClientApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateClientApiResponse() *UpdateClientApiResponse {
	p := new(UpdateClientApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateClientApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.UpdateClientApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateClientApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateClientApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateClientApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/roles/{extId} Put operation
*/
type UpdateRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRoleApiResponse() *UpdateRoleApiResponse {
	p := new(UpdateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.UpdateRoleApiResponse"}
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

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/authorization-policies/{extId} Get operation
*/
type ViewAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewAuthorizationPolicyApiResponse() *ViewAuthorizationPolicyApiResponse {
	p := new(ViewAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ViewAuthorizationPolicyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/clients/{extId} Get operation
*/
type ViewClientApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewClientApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewClientApiResponse() *ViewClientApiResponse {
	p := new(ViewClientApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewClientApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ViewClientApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewClientApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewClientApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewClientApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/entities/{extId} Get operation
*/
type ViewEntityApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewEntityApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewEntityApiResponse() *ViewEntityApiResponse {
	p := new(ViewEntityApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewEntityApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ViewEntityApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ViewEntityApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ViewEntityApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfViewEntityApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authz/operations/{extId} Get operation
*/
type ViewOperationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewOperationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewOperationApiResponse() *ViewOperationApiResponse {
	p := new(ViewOperationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewOperationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ViewOperationApiResponse"}
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

/*
REST response for all response codes in API path /iam/v4.0.b1/authz/roles/{extId} Get operation
*/
type ViewRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfViewRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewViewRoleApiResponse() *ViewRoleApiResponse {
	p := new(ViewRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ViewRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authz.ViewRoleApiResponse"}
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

type OneOfDeleteAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteAuthorizationPolicyApiResponseData() *OneOfDeleteAuthorizationPolicyApiResponseData {
	p := new(OneOfDeleteAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteAuthorizationPolicyApiResponseData is nil"))
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

func (p *OneOfDeleteAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAuthorizationPolicyApiResponseData"))
}

func (p *OneOfDeleteAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteAuthorizationPolicyApiResponseData")
}

type OneOfUpdateClientApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateClientApiResponseData() *OneOfUpdateClientApiResponseData {
	p := new(OneOfUpdateClientApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateClientApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateClientApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Message)
		}
		*p.oneOfType0 = v.(import3.Message)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfUpdateClientApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateClientApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Message)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateClientApiResponseData"))
}

func (p *OneOfUpdateClientApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateClientApiResponseData")
}

type OneOfListEntityApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Entity               `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListEntityApiResponseData() *OneOfListEntityApiResponseData {
	p := new(OneOfListEntityApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEntityApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEntityApiResponseData is nil"))
	}
	switch v.(type) {
	case []Entity:
		p.oneOfType0 = v.([]Entity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.Entity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.Entity>"
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

func (p *OneOfListEntityApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.Entity>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListEntityApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Entity)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.Entity" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.Entity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.Entity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEntityApiResponseData"))
}

func (p *OneOfListEntityApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListEntityApiResponseData")
}

type OneOfListAuthorizationPolicyApiResponseData struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType401  []AuthorizationPolicyProjection `json:"-"`
	oneOfType400  *import2.ErrorResponse          `json:"-"`
	oneOfType0    []AuthorizationPolicy           `json:"-"`
}

func NewOneOfListAuthorizationPolicyApiResponseData() *OneOfListAuthorizationPolicyApiResponseData {
	p := new(OneOfListAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAuthorizationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case []AuthorizationPolicyProjection:
		p.oneOfType401 = v.([]AuthorizationPolicyProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.AuthorizationPolicyProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.AuthorizationPolicyProjection>"
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
	case []AuthorizationPolicy:
		p.oneOfType0 = v.([]AuthorizationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.AuthorizationPolicy>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.AuthorizationPolicy>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.AuthorizationPolicyProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authz.AuthorizationPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType401 := new([]AuthorizationPolicyProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "iam.v4.authz.AuthorizationPolicyProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.AuthorizationPolicyProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.AuthorizationPolicyProjection>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]AuthorizationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.AuthorizationPolicy" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.AuthorizationPolicy>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.AuthorizationPolicy>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAuthorizationPolicyApiResponseData"))
}

func (p *OneOfListAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.AuthorizationPolicyProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authz.AuthorizationPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAuthorizationPolicyApiResponseData")
}

type OneOfCreateClientApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateClientApiResponseData() *OneOfCreateClientApiResponseData {
	p := new(OneOfCreateClientApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateClientApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateClientApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Message)
		}
		*p.oneOfType0 = v.(import3.Message)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfCreateClientApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateClientApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Message)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateClientApiResponseData"))
}

func (p *OneOfCreateClientApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateClientApiResponseData")
}

type OneOfListRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Role                 `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType401  []RoleProjection       `json:"-"`
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
	case []RoleProjection:
		p.oneOfType401 = v.([]RoleProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.RoleProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.RoleProjection>"
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
	if "List<iam.v4.authz.RoleProjection>" == *p.Discriminator {
		return p.oneOfType401
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType401 := new([]RoleProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "iam.v4.authz.RoleProjection" == *((*vOneOfType401)[0].ObjectType_) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.RoleProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.RoleProjection>"
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
	if "List<iam.v4.authz.RoleProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRoleApiResponseData")
}

type OneOfUpdateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	case import3.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Message)
		}
		*p.oneOfType0 = v.(import3.Message)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	vOneOfType0 := new(import3.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Message)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfViewRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2    *interface{}           `json:"-"`
	oneOfType0    *Role                  `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfCreateAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *AuthorizationPolicy   `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfCreateAuthorizationPolicyApiResponseData() *OneOfCreateAuthorizationPolicyApiResponseData {
	p := new(OneOfCreateAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateAuthorizationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case AuthorizationPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AuthorizationPolicy)
		}
		*p.oneOfType0 = v.(AuthorizationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfCreateAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AuthorizationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.AuthorizationPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AuthorizationPolicy)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAuthorizationPolicyApiResponseData"))
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateAuthorizationPolicyApiResponseData")
}

type OneOfListClientApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Client               `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfListClientApiResponseData() *OneOfListClientApiResponseData {
	p := new(OneOfListClientApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListClientApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListClientApiResponseData is nil"))
	}
	switch v.(type) {
	case []Client:
		p.oneOfType0 = v.([]Client)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authz.Client>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authz.Client>"
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

func (p *OneOfListClientApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authz.Client>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListClientApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Client)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authz.Client" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authz.Client>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authz.Client>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListClientApiResponseData"))
}

func (p *OneOfListClientApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authz.Client>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListClientApiResponseData")
}

type OneOfViewAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *AuthorizationPolicy   `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfViewAuthorizationPolicyApiResponseData() *OneOfViewAuthorizationPolicyApiResponseData {
	p := new(OneOfViewAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewAuthorizationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case AuthorizationPolicy:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AuthorizationPolicy)
		}
		*p.oneOfType0 = v.(AuthorizationPolicy)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfViewAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(AuthorizationPolicy)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.AuthorizationPolicy" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AuthorizationPolicy)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewAuthorizationPolicyApiResponseData"))
}

func (p *OneOfViewAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewAuthorizationPolicyApiResponseData")
}

type OneOfUpdateAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateAuthorizationPolicyApiResponseData() *OneOfUpdateAuthorizationPolicyApiResponseData {
	p := new(OneOfUpdateAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateAuthorizationPolicyApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Message)
		}
		*p.oneOfType0 = v.(import3.Message)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Message)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAuthorizationPolicyApiResponseData"))
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateAuthorizationPolicyApiResponseData")
}

type OneOfDeleteRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfListOperationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Operation            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfTenantConfigurationResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfTenantConfigurationResponseData() *OneOfTenantConfigurationResponseData {
	p := new(OneOfTenantConfigurationResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTenantConfigurationResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTenantConfigurationResponseData is nil"))
	}
	switch v.(type) {
	case import3.Message:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.Message)
		}
		*p.oneOfType0 = v.(import3.Message)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfTenantConfigurationResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfTenantConfigurationResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import3.Message)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "common.v1.config.Message" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.Message)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTenantConfigurationResponseData"))
}

func (p *OneOfTenantConfigurationResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfTenantConfigurationResponseData")
}

type OneOfViewClientApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Client                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfViewClientApiResponseData() *OneOfViewClientApiResponseData {
	p := new(OneOfViewClientApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewClientApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewClientApiResponseData is nil"))
	}
	switch v.(type) {
	case Client:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Client)
		}
		*p.oneOfType0 = v.(Client)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfViewClientApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewClientApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Client)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.Client" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Client)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewClientApiResponseData"))
}

func (p *OneOfViewClientApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewClientApiResponseData")
}

type OneOfCreateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Role                  `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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

type OneOfViewEntityApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Entity                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfViewEntityApiResponseData() *OneOfViewEntityApiResponseData {
	p := new(OneOfViewEntityApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfViewEntityApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfViewEntityApiResponseData is nil"))
	}
	switch v.(type) {
	case Entity:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Entity)
		}
		*p.oneOfType0 = v.(Entity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfViewEntityApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfViewEntityApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Entity)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.Entity" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Entity)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfViewEntityApiResponseData"))
}

func (p *OneOfViewEntityApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfViewEntityApiResponseData")
}

type OneOfAuthorizeApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *AuthorizationResponse `json:"-"`
}

func NewOneOfAuthorizeApiResponseData() *OneOfAuthorizeApiResponseData {
	p := new(OneOfAuthorizeApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAuthorizeApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAuthorizeApiResponseData is nil"))
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
	case AuthorizationResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AuthorizationResponse)
		}
		*p.oneOfType0 = v.(AuthorizationResponse)
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

func (p *OneOfAuthorizeApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAuthorizeApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(AuthorizationResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authz.AuthorizationResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AuthorizationResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAuthorizeApiResponseData"))
}

func (p *OneOfAuthorizeApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAuthorizeApiResponseData")
}

type OneOfViewOperationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
