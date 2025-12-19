/*
 * Generated file models/iam/v4/authz/authz_model.go.
 *
 * Product version: 4.1.1-beta-2
 *
 * Part of the Nutanix Identity and Access Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
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
Version of the API for the provided associated endpoint.
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
The object which defines the API endpoint, API version, and HTTP method for the API endpoint.
*/
type AssociatedEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApiVersion *ApiVersion `json:"apiVersion,omitempty"`
	/*
	  Endpoint URL.
	*/
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	HttpMethod *HttpMethod `json:"httpMethod,omitempty"`
}

func (p *AssociatedEndpoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AssociatedEndpoint

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

func (p *AssociatedEndpoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AssociatedEndpoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAssociatedEndpoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApiVersion != nil {
		p.ApiVersion = known.ApiVersion
	}
	if known.EndpointUrl != nil {
		p.EndpointUrl = known.EndpointUrl
	}
	if known.HttpMethod != nil {
		p.HttpMethod = known.HttpMethod
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiVersion")
	delete(allFields, "endpointUrl")
	delete(allFields, "httpMethod")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAssociatedEndpoint() *AssociatedEndpoint {
	p := new(AssociatedEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AssociatedEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of attributes supported for access control on this entity.
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
	  Display name of the entity's attribute used in Authorization Policy filters.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the entity's attribute used in Authorization Policy filters.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  List of supported operators for this entity.
	*/
	SupportedOperators []EntityOperators `json:"supportedOperators,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AttributeEntity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AttributeEntity

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

func (p *AttributeEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AttributeEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAttributeEntity()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AttributeValues != nil {
		p.AttributeValues = known.AttributeValues
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
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
	if known.SupportedOperators != nil {
		p.SupportedOperators = known.SupportedOperators
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributeValues")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "supportedOperators")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAttributeEntity() *AttributeEntity {
	p := new(AttributeEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AttributeEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An authorization policy that defines who can perform what operation on which entity.
*/
type AuthorizationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of user groups assigned to the authorization policy.
	*/
	AssignedUserGroupsCount *int `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the authorization policy.
	*/
	AssignedUsersCount *int `json:"assignedUsersCount,omitempty"`

	AuthorizationPolicyType *AuthorizationPolicyType `json:"authorizationPolicyType,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the authorization policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the authorization policy.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The identities for which the authorization policy is created.
	*/
	Identities []IdentityFilter `json:"identities,omitempty"`
	/*
	  Flag identifying if the authorization policy is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the authorization policy was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AuthorizationPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthorizationPolicy

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

func (p *AuthorizationPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.AuthorizationPolicyType != nil {
		p.AuthorizationPolicyType = known.AuthorizationPolicyType
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Identities != nil {
		p.Identities = known.Identities
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "authorizationPolicyType")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "identities")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "role")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationPolicy() *AuthorizationPolicy {
	p := new(AuthorizationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

type AuthorizationPolicyConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of user groups assigned to the authorization policy.
	*/
	AssignedUserGroupsCount *int `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the authorization policy.
	*/
	AssignedUsersCount *int `json:"assignedUsersCount,omitempty"`

	AuthorizationPolicyType *AuthorizationPolicyType `json:"authorizationPolicyType,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the authorization policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the authorization policy.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The identities for which the authorization policy is created.
	*/
	Identities []IdentityFilter `json:"identities,omitempty"`
	/*
	  Flag identifying if the authorization policy is internal or not.
	*/
	IsInternal *bool `json:"isInternal,omitempty"`
	/*
	  Flag identifying if the authorization policy is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the authorization policy was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Opaque data refers to the legacy filters used in the authorization policy.
	*/
	Opaquedata *string `json:"opaquedata,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AuthorizationPolicyConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthorizationPolicyConfig

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

func (p *AuthorizationPolicyConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationPolicyConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationPolicyConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.AuthorizationPolicyType != nil {
		p.AuthorizationPolicyType = known.AuthorizationPolicyType
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Identities != nil {
		p.Identities = known.Identities
	}
	if known.IsInternal != nil {
		p.IsInternal = known.IsInternal
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Opaquedata != nil {
		p.Opaquedata = known.Opaquedata
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "authorizationPolicyType")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "identities")
	delete(allFields, "isInternal")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "opaquedata")
	delete(allFields, "role")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationPolicyConfig() *AuthorizationPolicyConfig {
	p := new(AuthorizationPolicyConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationPolicyConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsInternal = new(bool)
	*p.IsInternal = false
	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

type AuthorizationPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of user groups assigned to the authorization policy.
	*/
	AssignedUserGroupsCount *int `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the authorization policy.
	*/
	AssignedUsersCount *int `json:"assignedUsersCount,omitempty"`

	AuthorizationPolicyType *AuthorizationPolicyType `json:"authorizationPolicyType,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the authorization policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the authorization policy.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The identities for which the authorization policy is created.
	*/
	Identities []IdentityFilter `json:"identities,omitempty"`
	/*
	  Flag identifying if the authorization policy is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the authorization policy was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`

	RoleProjection *RoleProjection `json:"roleProjection,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *AuthorizationPolicyProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthorizationPolicyProjection

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

func (p *AuthorizationPolicyProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationPolicyProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationPolicyProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.AuthorizationPolicyType != nil {
		p.AuthorizationPolicyType = known.AuthorizationPolicyType
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Identities != nil {
		p.Identities = known.Identities
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.RoleProjection != nil {
		p.RoleProjection = known.RoleProjection
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "authorizationPolicyType")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "identities")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "role")
	delete(allFields, "roleProjection")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationPolicyProjection() *AuthorizationPolicyProjection {
	p := new(AuthorizationPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/*
Type of authorization policy.
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
Request body for the authorization request.
*/
type AuthorizationRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Identity *IdentityObject `json:"identity"`
	/*
	  List of requests to be authorized.
	*/
	RequestList []RequestObject `json:"requestList"`
}

func (p *AuthorizationRequest) MarshalJSON() ([]byte, error) {
	type AuthorizationRequestProxy AuthorizationRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AuthorizationRequestProxy
		Identity    *IdentityObject `json:"identity,omitempty"`
		RequestList []RequestObject `json:"requestList,omitempty"`
	}{
		AuthorizationRequestProxy: (*AuthorizationRequestProxy)(p),
		Identity:                  p.Identity,
		RequestList:               p.RequestList,
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

func (p *AuthorizationRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Identity != nil {
		p.Identity = known.Identity
	}
	if known.RequestList != nil {
		p.RequestList = known.RequestList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "identity")
	delete(allFields, "requestList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationRequest() *AuthorizationRequest {
	p := new(AuthorizationRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity on which the action is being performed.
*/
type AuthorizationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of responses for the Authorization request.
	*/
	Response []AuthorizationResponseObject `json:"response,omitempty"`
}

func (p *AuthorizationResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthorizationResponse

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

func (p *AuthorizationResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Response != nil {
		p.Response = known.Response
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "response")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationResponse() *AuthorizationResponse {
	p := new(AuthorizationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Authorization enforcement response.
*/
type AuthorizationResponseObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  True if request was authorized, otherwise False.
	*/
	Authorized *interface{} `json:"authorized,omitempty"`
	/*
	  Filter expression.
	*/
	Filters *interface{} `json:"filters,omitempty"`
	/*
	  Authorization Request Id.
	*/
	RequestId *string `json:"requestId,omitempty"`

	ResponseType *ResponseType `json:"responseType,omitempty"`
}

func (p *AuthorizationResponseObject) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthorizationResponseObject

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

func (p *AuthorizationResponseObject) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizationResponseObject
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizationResponseObject()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Authorized != nil {
		p.Authorized = known.Authorized
	}
	if known.Filters != nil {
		p.Filters = known.Filters
	}
	if known.RequestId != nil {
		p.RequestId = known.RequestId
	}
	if known.ResponseType != nil {
		p.ResponseType = known.ResponseType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authorized")
	delete(allFields, "filters")
	delete(allFields, "requestId")
	delete(allFields, "responseType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizationResponseObject() *AuthorizationResponseObject {
	p := new(AuthorizationResponseObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizationResponseObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity on which the action is being performed.
*/
type AuthorizeEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Map of attributes for the entity.
	*/
	Attributes *interface{} `json:"attributes,omitempty"`
	/*
	  operation being performed on the entity.
	*/
	Operation *string `json:"operation"`
	/*
	  Type of the entity.
	*/
	Type *string `json:"type"`
}

func (p *AuthorizeEntity) MarshalJSON() ([]byte, error) {
	type AuthorizeEntityProxy AuthorizeEntity

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AuthorizeEntityProxy
		Operation *string `json:"operation,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{
		AuthorizeEntityProxy: (*AuthorizeEntityProxy)(p),
		Operation:            p.Operation,
		Type:                 p.Type,
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

func (p *AuthorizeEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthorizeEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthorizeEntity()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Attributes != nil {
		p.Attributes = known.Attributes
	}
	if known.Operation != nil {
		p.Operation = known.Operation
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributes")
	delete(allFields, "operation")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthorizeEntity() *AuthorizeEntity {
	p := new(AuthorizeEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthorizeEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The authZ migration API request body.
*/
type AuthzMigrationRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MigrationType *AuthzMigrationType `json:"migrationType"`
}

func (p *AuthzMigrationRequest) MarshalJSON() ([]byte, error) {
	type AuthzMigrationRequestProxy AuthzMigrationRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AuthzMigrationRequestProxy
		MigrationType *AuthzMigrationType `json:"migrationType,omitempty"`
	}{
		AuthzMigrationRequestProxy: (*AuthzMigrationRequestProxy)(p),
		MigrationType:              p.MigrationType,
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

func (p *AuthzMigrationRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthzMigrationRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthzMigrationRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MigrationType != nil {
		p.MigrationType = known.MigrationType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "migrationType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthzMigrationRequest() *AuthzMigrationRequest {
	p := new(AuthzMigrationRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.AuthzMigrationRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of authz migration.
*/
type AuthzMigrationType int

const (
	AUTHZMIGRATIONTYPE_UNKNOWN                                             AuthzMigrationType = 0
	AUTHZMIGRATIONTYPE_REDACTED                                            AuthzMigrationType = 1
	AUTHZMIGRATIONTYPE_REMOVE_AP_TENANT_ID_FILTERS                         AuthzMigrationType = 2
	AUTHZMIGRATIONTYPE_NC_USER_AP_TO_USER_DEFINED                          AuthzMigrationType = 3
	AUTHZMIGRATIONTYPE_ASSIGN_BASIC_USER_AP_TO_NC_USER_AP_USERS_USERGROUPS AuthzMigrationType = 4
	AUTHZMIGRATIONTYPE_REMOVE_CLIENT_FOUNDATION_CENTRAL                    AuthzMigrationType = 5
	AUTHZMIGRATIONTYPE_MOVE_ROLES_APS_FROM_PRISM_CLOUD_IAM                 AuthzMigrationType = 6
	AUTHZMIGRATIONTYPE_REMOVE_DUPLICATE_APS_AFTER_NAMESPACE_MIGRATION      AuthzMigrationType = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthzMigrationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REMOVE_AP_TENANT_ID_FILTERS",
		"NC_USER_AP_TO_USER_DEFINED",
		"ASSIGN_BASIC_USER_AP_TO_NC_USER_AP_USERS_USERGROUPS",
		"REMOVE_CLIENT_FOUNDATION_CENTRAL",
		"MOVE_ROLES_APS_FROM_PRISM_CLOUD_IAM",
		"REMOVE_DUPLICATE_APS_AFTER_NAMESPACE_MIGRATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthzMigrationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REMOVE_AP_TENANT_ID_FILTERS",
		"NC_USER_AP_TO_USER_DEFINED",
		"ASSIGN_BASIC_USER_AP_TO_NC_USER_AP_USERS_USERGROUPS",
		"REMOVE_CLIENT_FOUNDATION_CENTRAL",
		"MOVE_ROLES_APS_FROM_PRISM_CLOUD_IAM",
		"REMOVE_DUPLICATE_APS_AFTER_NAMESPACE_MIGRATION",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthzMigrationType) index(name string) AuthzMigrationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REMOVE_AP_TENANT_ID_FILTERS",
		"NC_USER_AP_TO_USER_DEFINED",
		"ASSIGN_BASIC_USER_AP_TO_NC_USER_AP_USERS_USERGROUPS",
		"REMOVE_CLIENT_FOUNDATION_CENTRAL",
		"MOVE_ROLES_APS_FROM_PRISM_CLOUD_IAM",
		"REMOVE_DUPLICATE_APS_AFTER_NAMESPACE_MIGRATION",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthzMigrationType(idx)
		}
	}
	return AUTHZMIGRATIONTYPE_UNKNOWN
}

func (e *AuthzMigrationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthzMigrationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthzMigrationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthzMigrationType) Ref() *AuthzMigrationType {
	return &e
}

/*
An authorization policy that defines who can perform what operation on which entity, without including identity details.
*/
type BaseAuthorizationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BaseAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseAuthorizationPolicy

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

func (p *BaseAuthorizationPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseAuthorizationPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseAuthorizationPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientName")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "role")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseAuthorizationPolicy() *BaseAuthorizationPolicy {
	p := new(BaseAuthorizationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.BaseAuthorizationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BaseAuthorizationPolicyProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BaseRoleProjection *BaseRoleProjection `json:"baseRoleProjection,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BaseAuthorizationPolicyProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseAuthorizationPolicyProjection

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

func (p *BaseAuthorizationPolicyProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseAuthorizationPolicyProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseAuthorizationPolicyProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BaseRoleProjection != nil {
		p.BaseRoleProjection = known.BaseRoleProjection
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "baseRoleProjection")
	delete(allFields, "clientName")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "role")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseAuthorizationPolicyProjection() *BaseAuthorizationPolicyProjection {
	p := new(BaseAuthorizationPolicyProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.BaseAuthorizationPolicyProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A base role is a group of operations, but does not include the operation details.
*/
type BaseRole struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BaseRole) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseRole

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

func (p *BaseRole) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseRole
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseRole()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
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
	delete(allFields, "clientName")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseRole() *BaseRole {
	p := new(BaseRole)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.BaseRole"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BaseRoleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *BaseRoleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseRoleProjection

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

func (p *BaseRoleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseRoleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseRoleProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
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
	delete(allFields, "clientName")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseRoleProjection() *BaseRoleProjection {
	p := new(BaseRoleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.BaseRoleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the client registered.
*/
type Client struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or Service Name that created the client.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the client.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Deployment List related to the client.
	*/
	DeploymentList []string `json:"deploymentList,omitempty"`
	/*
	  Description of the client.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI name of the client.
	*/
	DisplayName *string `json:"displayName"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the client was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the client.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Client) MarshalJSON() ([]byte, error) {
	type ClientProxy Client

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClientProxy
		DisplayName *string `json:"displayName,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{
		ClientProxy: (*ClientProxy)(p),
		DisplayName: p.DisplayName,
		Name:        p.Name,
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

func (p *Client) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Client
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClient()

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
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DeploymentList != nil {
		p.DeploymentList = known.DeploymentList
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
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
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "deploymentList")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClient() *Client {
	p := new(Client)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Client"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the client configuration that is registered.
*/
type ClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or Service Name that created the client.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the client.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Deployment List related to the client.
	*/
	DeploymentList []string `json:"deploymentList,omitempty"`
	/*
	  Description of the client.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI name of the client.
	*/
	DisplayName *string `json:"displayName"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the client was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  Identifier for the old/legacy client name.
	*/
	LegacyClientName *string `json:"legacyClientName,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the client.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClientConfig) MarshalJSON() ([]byte, error) {
	type ClientConfigProxy ClientConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClientConfigProxy
		DisplayName *string `json:"displayName,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{
		ClientConfigProxy: (*ClientConfigProxy)(p),
		DisplayName:       p.DisplayName,
		Name:              p.Name,
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

func (p *ClientConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClientConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClientConfig()

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
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DeploymentList != nil {
		p.DeploymentList = known.DeploymentList
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyClientName != nil {
		p.LegacyClientName = known.LegacyClientName
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
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "deploymentList")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyClientName")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClientConfig() *ClientConfig {
	p := new(ClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IAM configuration changset instance describing the change happened over the last deployed version.
*/
type ConfigChangeset struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Changeset *ConfigChangesetData `json:"changeset,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`

	ConfigType *ConfigType `json:"configType,omitempty"`
	/*
	  The creation time of the IAM configuration changeset record.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  The display name of the IAM configuration entity for which the changeset is generated.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the IAM configuration changeset record was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConfigChangeset) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigChangeset

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

func (p *ConfigChangeset) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangeset
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangeset()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Changeset != nil {
		p.Changeset = known.Changeset
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.ConfigType != nil {
		p.ConfigType = known.ConfigType
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
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
	delete(allFields, "changeset")
	delete(allFields, "clientName")
	delete(allFields, "configType")
	delete(allFields, "createdTime")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangeset() *ConfigChangeset {
	p := new(ConfigChangeset)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangeset"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The difference in the IAM configuration entity from the last deployed version.
*/
type ConfigChangesetData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DisplayNameChangeset *ConfigChangesetDisplayName `json:"displayNameChangeset,omitempty"`

	LegacyObjectListChangeset *ConfigChangesetDataObjects `json:"legacyObjectListChangeset,omitempty"`

	LegacyOperationListChangeset *ConfigChangesetDataOperations `json:"legacyOperationListChangeset,omitempty"`

	LegacyParentOperationListChangeset *ConfigChangesetDataOperations `json:"legacyParentOperationListChangeset,omitempty"`
}

func (p *ConfigChangesetData) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigChangesetData

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

func (p *ConfigChangesetData) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangesetData
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangesetData()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayNameChangeset != nil {
		p.DisplayNameChangeset = known.DisplayNameChangeset
	}
	if known.LegacyObjectListChangeset != nil {
		p.LegacyObjectListChangeset = known.LegacyObjectListChangeset
	}
	if known.LegacyOperationListChangeset != nil {
		p.LegacyOperationListChangeset = known.LegacyOperationListChangeset
	}
	if known.LegacyParentOperationListChangeset != nil {
		p.LegacyParentOperationListChangeset = known.LegacyParentOperationListChangeset
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayNameChangeset")
	delete(allFields, "legacyObjectListChangeset")
	delete(allFields, "legacyOperationListChangeset")
	delete(allFields, "legacyParentOperationListChangeset")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangesetData() *ConfigChangesetData {
	p := new(ConfigChangesetData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangesetData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The changeset data corresponding to IAM configuration "entity".
*/
type ConfigChangesetDataObjects struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Configuration changeset data for added legacyObjectList entries.
	*/
	AddedObjects []string `json:"addedObjects,omitempty"`
	/*
	  Configuration changeset data for removed legacyObjectList entries.
	*/
	RemovedObjects []string `json:"removedObjects,omitempty"`
}

func (p *ConfigChangesetDataObjects) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigChangesetDataObjects

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

func (p *ConfigChangesetDataObjects) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangesetDataObjects
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangesetDataObjects()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AddedObjects != nil {
		p.AddedObjects = known.AddedObjects
	}
	if known.RemovedObjects != nil {
		p.RemovedObjects = known.RemovedObjects
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addedObjects")
	delete(allFields, "removedObjects")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangesetDataObjects() *ConfigChangesetDataObjects {
	p := new(ConfigChangesetDataObjects)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangesetDataObjects"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration changeset data associated with the IAM configuration "Operation".
*/
type ConfigChangesetDataOperations struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Configuration changeset data for added IAM configuration "Operation" entries.
	*/
	AddedOperations []string `json:"addedOperations,omitempty"`
	/*
	  Configuration changeset data for removed IAM configuration "Operation" entries.
	*/
	RemovedOperations []string `json:"removedOperations,omitempty"`
}

func (p *ConfigChangesetDataOperations) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigChangesetDataOperations

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

func (p *ConfigChangesetDataOperations) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangesetDataOperations
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangesetDataOperations()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AddedOperations != nil {
		p.AddedOperations = known.AddedOperations
	}
	if known.RemovedOperations != nil {
		p.RemovedOperations = known.RemovedOperations
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addedOperations")
	delete(allFields, "removedOperations")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangesetDataOperations() *ConfigChangesetDataOperations {
	p := new(ConfigChangesetDataOperations)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangesetDataOperations"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The changeset corresponding to the change in the display name of the IAM configuration entity.
*/
type ConfigChangesetDisplayName struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The current display name of the IAM configuration entity.
	*/
	NewDisplayName *string `json:"newDisplayName"`
	/*
	  The old display name of the IAM configuration entity.
	*/
	PreviousDisplayName *string `json:"previousDisplayName"`
}

func (p *ConfigChangesetDisplayName) MarshalJSON() ([]byte, error) {
	type ConfigChangesetDisplayNameProxy ConfigChangesetDisplayName

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConfigChangesetDisplayNameProxy
		NewDisplayName      *string `json:"newDisplayName,omitempty"`
		PreviousDisplayName *string `json:"previousDisplayName,omitempty"`
	}{
		ConfigChangesetDisplayNameProxy: (*ConfigChangesetDisplayNameProxy)(p),
		NewDisplayName:                  p.NewDisplayName,
		PreviousDisplayName:             p.PreviousDisplayName,
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

func (p *ConfigChangesetDisplayName) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangesetDisplayName
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangesetDisplayName()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NewDisplayName != nil {
		p.NewDisplayName = known.NewDisplayName
	}
	if known.PreviousDisplayName != nil {
		p.PreviousDisplayName = known.PreviousDisplayName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "newDisplayName")
	delete(allFields, "previousDisplayName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangesetDisplayName() *ConfigChangesetDisplayName {
	p := new(ConfigChangesetDisplayName)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangesetDisplayName"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ConfigChangesetProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Changeset *ConfigChangesetData `json:"changeset,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`

	ConfigType *ConfigType `json:"configType,omitempty"`
	/*
	  The creation time of the IAM configuration changeset record.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  The display name of the IAM configuration entity for which the changeset is generated.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the IAM configuration changeset record was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ConfigChangesetProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConfigChangesetProjection

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

func (p *ConfigChangesetProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConfigChangesetProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConfigChangesetProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Changeset != nil {
		p.Changeset = known.Changeset
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.ConfigType != nil {
		p.ConfigType = known.ConfigType
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
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
	delete(allFields, "changeset")
	delete(allFields, "clientName")
	delete(allFields, "configType")
	delete(allFields, "createdTime")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfigChangesetProjection() *ConfigChangesetProjection {
	p := new(ConfigChangesetProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ConfigChangesetProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of IAM configuration.
*/
type ConfigType int

const (
	CONFIGTYPE_UNKNOWN   ConfigType = 0
	CONFIGTYPE_REDACTED  ConfigType = 1
	CONFIGTYPE_OPERATION ConfigType = 2
	CONFIGTYPE_ENTITY    ConfigType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConfigType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPERATION",
		"ENTITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConfigType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPERATION",
		"ENTITY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConfigType) index(name string) ConfigType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPERATION",
		"ENTITY",
	}
	for idx := range names {
		if names[idx] == name {
			return ConfigType(idx)
		}
	}
	return CONFIGTYPE_UNKNOWN
}

func (e *ConfigType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConfigType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConfigType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConfigType) Ref() *ConfigType {
	return &e
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authz/authorization-policies Post operation
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

func (p *CreateAuthorizationPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateAuthorizationPolicyApiResponse

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

func (p *CreateAuthorizationPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateAuthorizationPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateAuthorizationPolicyApiResponse()

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

func NewCreateAuthorizationPolicyApiResponse() *CreateAuthorizationPolicyApiResponse {
	p := new(CreateAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
REST response for all response codes in API path /iam/v4.1.b2/authz/roles Post operation
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

func (p *CreateRoleApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateRoleApiResponse

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

func (p *CreateRoleApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateRoleApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateRoleApiResponse()

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

func NewCreateRoleApiResponse() *CreateRoleApiResponse {
	p := new(CreateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.CreateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
REST response for all response codes in API path /iam/v4.1.b2/authz/authorization-policies/{extId} Delete operation
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

func (p *DeleteAuthorizationPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteAuthorizationPolicyApiResponse

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

func (p *DeleteAuthorizationPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteAuthorizationPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteAuthorizationPolicyApiResponse()

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

func NewDeleteAuthorizationPolicyApiResponse() *DeleteAuthorizationPolicyApiResponse {
	p := new(DeleteAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
REST response for all response codes in API path /iam/v4.1.b2/authz/roles/{extId} Delete operation
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

func (p *DeleteRoleApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteRoleApiResponse

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

func (p *DeleteRoleApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteRoleApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteRoleApiResponse()

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

func NewDeleteRoleApiResponse() *DeleteRoleApiResponse {
	p := new(DeleteRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.DeleteRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
Information of the entity.
*/
type Entity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AttributeList []AttributeEntity `json:"attributeList,omitempty"`
	/*
	  Name of client the entity belongs to.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the entity.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the entity.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the entity.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI display name of the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the entity supports scoping using multiple attributes which will result in a logical AND.
	*/
	IsLogicalAndSupportedForAttributes *bool `json:"isLogicalAndSupportedForAttributes,omitempty"`
	/*
	  The time when the entity was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  URL provided by the client to search the entities.
	*/
	SearchURL *string `json:"searchURL,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Entity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Entity

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

func (p *Entity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Entity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntity()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AttributeList != nil {
		p.AttributeList = known.AttributeList
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsLogicalAndSupportedForAttributes != nil {
		p.IsLogicalAndSupportedForAttributes = known.IsLogicalAndSupportedForAttributes
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.SearchURL != nil {
		p.SearchURL = known.SearchURL
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributeList")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "isLogicalAndSupportedForAttributes")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "searchURL")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Entity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the entity configuration.
*/
type EntityConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AttributeList []AttributeEntity `json:"attributeList,omitempty"`
	/*
	  Name of client the entity belongs to.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the entity.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the entity.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the entity.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  UI display name of the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the entity is internal or not.
	*/
	IsInternal *bool `json:"isInternal,omitempty"`
	/*
	  Indicates whether the entity supports scoping using multiple attributes which will result in a logical AND.
	*/
	IsLogicalAndSupportedForAttributes *bool `json:"isLogicalAndSupportedForAttributes,omitempty"`
	/*
	  The time when the entity was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  List of old displayNames of the given entity which have been consolidated or renamed.
	*/
	LegacyObjectList []string `json:"legacyObjectList,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the entity.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  URL provided by the client to search the entities.
	*/
	SearchURL *string `json:"searchURL,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *EntityConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityConfig

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

func (p *EntityConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AttributeList != nil {
		p.AttributeList = known.AttributeList
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsInternal != nil {
		p.IsInternal = known.IsInternal
	}
	if known.IsLogicalAndSupportedForAttributes != nil {
		p.IsLogicalAndSupportedForAttributes = known.IsLogicalAndSupportedForAttributes
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyObjectList != nil {
		p.LegacyObjectList = known.LegacyObjectList
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.SearchURL != nil {
		p.SearchURL = known.SearchURL
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributeList")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "isInternal")
	delete(allFields, "isLogicalAndSupportedForAttributes")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyObjectList")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "searchURL")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityConfig() *EntityConfig {
	p := new(EntityConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.EntityConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Entity filter object definition for the entities.
*/
type EntityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Information of the entity filter present in the EntityFilter object.
	*/
	EntityFilter *interface{} `json:"entityFilter,omitempty"`
}

func (p *EntityFilter) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityFilter

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

func (p *EntityFilter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityFilter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEntityFilter()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityFilter != nil {
		p.EntityFilter = known.EntityFilter
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityFilter")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityFilter() *EntityFilter {
	p := new(EntityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.EntityFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of authorization policy.
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
REST response for all response codes in API path /iam/v4.1.b2/authz/authorization-policies/{extId} Get operation
*/
type GetAuthorizationPolicyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAuthorizationPolicyApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetAuthorizationPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetAuthorizationPolicyApiResponse

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

func (p *GetAuthorizationPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetAuthorizationPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetAuthorizationPolicyApiResponse()

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

func NewGetAuthorizationPolicyApiResponse() *GetAuthorizationPolicyApiResponse {
	p := new(GetAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.GetAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAuthorizationPolicyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAuthorizationPolicyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAuthorizationPolicyApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/clients/{extId} Get operation
*/
type GetClientApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClientApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetClientApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetClientApiResponse

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

func (p *GetClientApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetClientApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetClientApiResponse()

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

func NewGetClientApiResponse() *GetClientApiResponse {
	p := new(GetClientApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.GetClientApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetClientApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetClientApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetClientApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/entities/{extId} Get operation
*/
type GetEntityApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEntityApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetEntityApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetEntityApiResponse

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

func (p *GetEntityApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetEntityApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetEntityApiResponse()

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

func NewGetEntityApiResponse() *GetEntityApiResponse {
	p := new(GetEntityApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.GetEntityApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEntityApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEntityApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEntityApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/operations/{extId} Get operation
*/
type GetOperationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetOperationApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetOperationApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetOperationApiResponse

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

func (p *GetOperationApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetOperationApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetOperationApiResponse()

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

func NewGetOperationApiResponse() *GetOperationApiResponse {
	p := new(GetOperationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.GetOperationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetOperationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetOperationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetOperationApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/roles/{extId} Get operation
*/
type GetRoleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRoleApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetRoleApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetRoleApiResponse

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

func (p *GetRoleApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetRoleApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetRoleApiResponse()

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

func NewGetRoleApiResponse() *GetRoleApiResponse {
	p := new(GetRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.GetRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRoleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRoleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRoleApiResponseData()
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
HTTP method for the provided associated endpoint.
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

/*
Identity filter object definition for the identities.
*/
type IdentityFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Information of the identity filter present in the IdentityFilter object.
	*/
	IdentityFilter *interface{} `json:"identityFilter,omitempty"`
}

func (p *IdentityFilter) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IdentityFilter

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

func (p *IdentityFilter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IdentityFilter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIdentityFilter()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IdentityFilter != nil {
		p.IdentityFilter = known.IdentityFilter
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "identityFilter")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIdentityFilter() *IdentityFilter {
	p := new(IdentityFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.IdentityFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Object defining the attributes of Identity (who).
*/
type IdentityObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Map of attributes for the Identity.
	*/
	Attributes *interface{} `json:"attributes,omitempty"`

	IdentityType *IdentityType `json:"identityType,omitempty"`
}

func (p *IdentityObject) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias IdentityObject

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

func (p *IdentityObject) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IdentityObject
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIdentityObject()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Attributes != nil {
		p.Attributes = known.Attributes
	}
	if known.IdentityType != nil {
		p.IdentityType = known.IdentityType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributes")
	delete(allFields, "identityType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIdentityObject() *IdentityObject {
	p := new(IdentityObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.IdentityObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of identity object.
*/
type IdentityType int

const (
	IDENTITYTYPE_UNKNOWN  IdentityType = 0
	IDENTITYTYPE_REDACTED IdentityType = 1
	IDENTITYTYPE_USER     IdentityType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IdentityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"User",
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
REST response for all response codes in API path /iam/v4.1.b2/authz/authorization-policies Get operation
*/
type ListAuthorizationPoliciesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListAuthorizationPoliciesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListAuthorizationPoliciesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListAuthorizationPoliciesApiResponse

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

func (p *ListAuthorizationPoliciesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListAuthorizationPoliciesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListAuthorizationPoliciesApiResponse()

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

func NewListAuthorizationPoliciesApiResponse() *ListAuthorizationPoliciesApiResponse {
	p := new(ListAuthorizationPoliciesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListAuthorizationPoliciesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListAuthorizationPoliciesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListAuthorizationPoliciesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListAuthorizationPoliciesApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/entities Get operation
*/
type ListEntitiesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEntitiesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListEntitiesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListEntitiesApiResponse

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

func (p *ListEntitiesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListEntitiesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListEntitiesApiResponse()

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

func NewListEntitiesApiResponse() *ListEntitiesApiResponse {
	p := new(ListEntitiesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListEntitiesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEntitiesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEntitiesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEntitiesApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/operations Get operation
*/
type ListOperationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListOperationsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListOperationsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListOperationsApiResponse

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

func (p *ListOperationsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListOperationsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListOperationsApiResponse()

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

func NewListOperationsApiResponse() *ListOperationsApiResponse {
	p := new(ListOperationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListOperationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListOperationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListOperationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListOperationsApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authz/roles Get operation
*/
type ListRolesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRolesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRolesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRolesApiResponse

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

func (p *ListRolesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRolesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRolesApiResponse()

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

func NewListRolesApiResponse() *ListRolesApiResponse {
	p := new(ListRolesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.ListRolesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRolesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRolesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRolesApiResponseData()
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
Migrate role spec details.
*/
type MigrateRoleSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag indicating whether to include related operations in the final list after migration.
	*/
	ShouldIncludeRelatedOperations *bool `json:"shouldIncludeRelatedOperations"`
}

func (p *MigrateRoleSpec) MarshalJSON() ([]byte, error) {
	type MigrateRoleSpecProxy MigrateRoleSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*MigrateRoleSpecProxy
		ShouldIncludeRelatedOperations *bool `json:"shouldIncludeRelatedOperations,omitempty"`
	}{
		MigrateRoleSpecProxy:           (*MigrateRoleSpecProxy)(p),
		ShouldIncludeRelatedOperations: p.ShouldIncludeRelatedOperations,
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

func (p *MigrateRoleSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrateRoleSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMigrateRoleSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ShouldIncludeRelatedOperations != nil {
		p.ShouldIncludeRelatedOperations = known.ShouldIncludeRelatedOperations
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "shouldIncludeRelatedOperations")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMigrateRoleSpec() *MigrateRoleSpec {
	p := new(MigrateRoleSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.MigrateRoleSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation schema change-specific notes for capturing information to surface to the user to take relevant actions.
*/
type MigrationActionNotes struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specific detailed description around operation schema changes for the end user to understand the impact.
	*/
	Notes *string `json:"notes"`
}

func (p *MigrationActionNotes) MarshalJSON() ([]byte, error) {
	type MigrationActionNotesProxy MigrationActionNotes

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*MigrationActionNotesProxy
		Notes *string `json:"notes,omitempty"`
	}{
		MigrationActionNotesProxy: (*MigrationActionNotesProxy)(p),
		Notes:                     p.Notes,
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

func (p *MigrationActionNotes) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MigrationActionNotes
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewMigrationActionNotes()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Notes != nil {
		p.Notes = known.Notes
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "notes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewMigrationActionNotes() *MigrationActionNotes {
	p := new(MigrationActionNotes)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.MigrationActionNotes"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Operation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of associated endpoint objects for the operation.
	*/
	AssociatedEndpointList []AssociatedEndpoint `json:"associatedEndpointList,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The creation time of the operation.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the operation.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of the operation.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Type of entity associated with this operation.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the operation was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	MigrationActionNotes *MigrationActionNotes `json:"migrationActionNotes,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  List of related operations. These are the operations that may need to be accessed along with the current operation in order for certain workflows to succeed.
	*/
	RelatedOperationList []string `json:"relatedOperationList,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Operation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Operation

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

func (p *Operation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Operation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOperation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssociatedEndpointList != nil {
		p.AssociatedEndpointList = known.AssociatedEndpointList
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MigrationActionNotes != nil {
		p.MigrationActionNotes = known.MigrationActionNotes
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.RelatedOperationList != nil {
		p.RelatedOperationList = known.RelatedOperationList
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associatedEndpointList")
	delete(allFields, "clientName")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "migrationActionNotes")
	delete(allFields, "operationType")
	delete(allFields, "relatedOperationList")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOperation() *Operation {
	p := new(Operation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Operation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OperationConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of associated endpoint objects for the operation.
	*/
	AssociatedEndpointList []AssociatedEndpoint `json:"associatedEndpointList,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The creation time of the operation.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the operation.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Name of the operation.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Type of entity associated with this operation.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag to denote if operation is internal or not.
	*/
	IsInternal *bool `json:"isInternal,omitempty"`
	/*
	  The time when the operation was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  Legacy or existing Operation name(s) which have been replaced/consolidated by the new Operation name.
	*/
	LegacyOperationList []string `json:"legacyOperationList,omitempty"`
	/*
	  Existing operation name(s) which have been split into operations including the given new Operation name.
	*/
	LegacyParentOperationList []string `json:"legacyParentOperationList,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	MigrationActionNotes *MigrationActionNotes `json:"migrationActionNotes,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  List of related operations. These are the operations that may need to be accessed along with the current operation in order for certain workflows to succeed.
	*/
	RelatedOperationList []string `json:"relatedOperationList,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *OperationConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OperationConfig

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

func (p *OperationConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OperationConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOperationConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssociatedEndpointList != nil {
		p.AssociatedEndpointList = known.AssociatedEndpointList
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsInternal != nil {
		p.IsInternal = known.IsInternal
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyOperationList != nil {
		p.LegacyOperationList = known.LegacyOperationList
	}
	if known.LegacyParentOperationList != nil {
		p.LegacyParentOperationList = known.LegacyParentOperationList
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MigrationActionNotes != nil {
		p.MigrationActionNotes = known.MigrationActionNotes
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.RelatedOperationList != nil {
		p.RelatedOperationList = known.RelatedOperationList
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "associatedEndpointList")
	delete(allFields, "clientName")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "isInternal")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyOperationList")
	delete(allFields, "legacyParentOperationList")
	delete(allFields, "links")
	delete(allFields, "migrationActionNotes")
	delete(allFields, "operationType")
	delete(allFields, "relatedOperationList")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOperationConfig() *OperationConfig {
	p := new(OperationConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.OperationConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The impact of operation schema changes. Currently we support IMPACTED, NOT_IMPACTED and UNKNOWN.
*/
type OperationSchemaChangeImpact int

const (
	OPERATIONSCHEMACHANGEIMPACT_UNKNOWN      OperationSchemaChangeImpact = 0
	OPERATIONSCHEMACHANGEIMPACT_REDACTED     OperationSchemaChangeImpact = 1
	OPERATIONSCHEMACHANGEIMPACT_IMPACTED     OperationSchemaChangeImpact = 2
	OPERATIONSCHEMACHANGEIMPACT_NOT_IMPACTED OperationSchemaChangeImpact = 3
	OPERATIONSCHEMACHANGEIMPACT_UNDETERMINED OperationSchemaChangeImpact = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationSchemaChangeImpact) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMPACTED",
		"NOT_IMPACTED",
		"UNDETERMINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationSchemaChangeImpact) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMPACTED",
		"NOT_IMPACTED",
		"UNDETERMINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationSchemaChangeImpact) index(name string) OperationSchemaChangeImpact {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMPACTED",
		"NOT_IMPACTED",
		"UNDETERMINED",
	}
	for idx := range names {
		if names[idx] == name {
			return OperationSchemaChangeImpact(idx)
		}
	}
	return OPERATIONSCHEMACHANGEIMPACT_UNKNOWN
}

func (e *OperationSchemaChangeImpact) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OperationSchemaChangeImpact:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OperationSchemaChangeImpact) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OperationSchemaChangeImpact) Ref() *OperationSchemaChangeImpact {
	return &e
}

/*
The operation type. Currently we support INTERNAL, EXTERNAL and SYSTEM_DEFINED_ONLY.
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
The incoming request.
*/
type RequestObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Entity *AuthorizeEntity `json:"entity"`
	/*
	  Unique qualifier of the request.
	*/
	RequestID *string `json:"requestID"`

	ResponseType *ResponseType `json:"responseType,omitempty"`
}

func (p *RequestObject) MarshalJSON() ([]byte, error) {
	type RequestObjectProxy RequestObject

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RequestObjectProxy
		Entity    *AuthorizeEntity `json:"entity,omitempty"`
		RequestID *string          `json:"requestID,omitempty"`
	}{
		RequestObjectProxy: (*RequestObjectProxy)(p),
		Entity:             p.Entity,
		RequestID:          p.RequestID,
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

func (p *RequestObject) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RequestObject
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRequestObject()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Entity != nil {
		p.Entity = known.Entity
	}
	if known.RequestID != nil {
		p.RequestID = known.RequestID
	}
	if known.ResponseType != nil {
		p.ResponseType = known.ResponseType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entity")
	delete(allFields, "requestID")
	delete(allFields, "responseType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRequestObject() *RequestObject {
	p := new(RequestObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.RequestObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of response expected.
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
A role to group the operations.
*/
type Role struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Accessible clients for the role.
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/*
	  Count of Accessible clients for the role.
	*/
	AccessibleClientsCount *int `json:"accessibleClientsCount,omitempty"`
	/*
	  List of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypes []string `json:"accessibleEntityTypes,omitempty"`
	/*
	  Count of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypesCount *int `json:"accessibleEntityTypesCount,omitempty"`
	/*
	  Number of user groups assigned to the given role.
	*/
	AssignedUserGroupsCount *int64 `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the given role.
	*/
	AssignedUsersCount *int64 `json:"assignedUsersCount,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the role.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the role.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the role.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the role is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the role was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	OperationSchemaChangeImpact *OperationSchemaChangeImpact `json:"operationSchemaChangeImpact,omitempty"`
	/*
	  List of operations for the role.
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Role) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Role

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

func (p *Role) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Role
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRole()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessibleClients != nil {
		p.AccessibleClients = known.AccessibleClients
	}
	if known.AccessibleClientsCount != nil {
		p.AccessibleClientsCount = known.AccessibleClientsCount
	}
	if known.AccessibleEntityTypes != nil {
		p.AccessibleEntityTypes = known.AccessibleEntityTypes
	}
	if known.AccessibleEntityTypesCount != nil {
		p.AccessibleEntityTypesCount = known.AccessibleEntityTypesCount
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.OperationSchemaChangeImpact != nil {
		p.OperationSchemaChangeImpact = known.OperationSchemaChangeImpact
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessibleClients")
	delete(allFields, "accessibleClientsCount")
	delete(allFields, "accessibleEntityTypes")
	delete(allFields, "accessibleEntityTypesCount")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "operationSchemaChangeImpact")
	delete(allFields, "operations")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRole() *Role {
	p := new(Role)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Role"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

type RoleConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of Accessible clients for the role.
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/*
	  Count of Accessible clients for the role.
	*/
	AccessibleClientsCount *int `json:"accessibleClientsCount,omitempty"`
	/*
	  List of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypes []string `json:"accessibleEntityTypes,omitempty"`
	/*
	  Count of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypesCount *int `json:"accessibleEntityTypesCount,omitempty"`
	/*
	  Number of user groups assigned to the given role.
	*/
	AssignedUserGroupsCount *int64 `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the given role.
	*/
	AssignedUsersCount *int64 `json:"assignedUsersCount,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the role.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the role.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the role.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the role is internal or not.
	*/
	IsInternal *bool `json:"isInternal,omitempty"`
	/*
	  Flag identifying if the role is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the role was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  Old IUN of the role.
	*/
	LegacyRole *string `json:"legacyRole,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	OperationSchemaChangeImpact *OperationSchemaChangeImpact `json:"operationSchemaChangeImpact,omitempty"`
	/*
	  List of operations for the role.
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RoleConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RoleConfig

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

func (p *RoleConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RoleConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRoleConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessibleClients != nil {
		p.AccessibleClients = known.AccessibleClients
	}
	if known.AccessibleClientsCount != nil {
		p.AccessibleClientsCount = known.AccessibleClientsCount
	}
	if known.AccessibleEntityTypes != nil {
		p.AccessibleEntityTypes = known.AccessibleEntityTypes
	}
	if known.AccessibleEntityTypesCount != nil {
		p.AccessibleEntityTypesCount = known.AccessibleEntityTypesCount
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsInternal != nil {
		p.IsInternal = known.IsInternal
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyRole != nil {
		p.LegacyRole = known.LegacyRole
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.OperationSchemaChangeImpact != nil {
		p.OperationSchemaChangeImpact = known.OperationSchemaChangeImpact
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessibleClients")
	delete(allFields, "accessibleClientsCount")
	delete(allFields, "accessibleEntityTypes")
	delete(allFields, "accessibleEntityTypesCount")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "isInternal")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyRole")
	delete(allFields, "links")
	delete(allFields, "operationSchemaChangeImpact")
	delete(allFields, "operations")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRoleConfig() *RoleConfig {
	p := new(RoleConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.RoleConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
	  List of Accessible clients for the role.
	*/
	AccessibleClients []string `json:"accessibleClients,omitempty"`
	/*
	  Count of Accessible clients for the role.
	*/
	AccessibleClientsCount *int `json:"accessibleClientsCount,omitempty"`
	/*
	  List of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypes []string `json:"accessibleEntityTypes,omitempty"`
	/*
	  Count of Accessible Entity Types for the role.
	*/
	AccessibleEntityTypesCount *int `json:"accessibleEntityTypesCount,omitempty"`
	/*
	  Number of user groups assigned to the given role.
	*/
	AssignedUserGroupsCount *int64 `json:"assignedUserGroupsCount,omitempty"`
	/*
	  Number of users assigned to the given role.
	*/
	AssignedUsersCount *int64 `json:"assignedUsersCount,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  User or Service Name that created the role.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the role.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Description of the role.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag identifying if the role is system-defined or not.
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  The time when the role was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	OperationSchemaChangeImpact *OperationSchemaChangeImpact `json:"operationSchemaChangeImpact,omitempty"`
	/*
	  List of operations for the role.
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RoleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RoleProjection

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

func (p *RoleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RoleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRoleProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessibleClients != nil {
		p.AccessibleClients = known.AccessibleClients
	}
	if known.AccessibleClientsCount != nil {
		p.AccessibleClientsCount = known.AccessibleClientsCount
	}
	if known.AccessibleEntityTypes != nil {
		p.AccessibleEntityTypes = known.AccessibleEntityTypes
	}
	if known.AccessibleEntityTypesCount != nil {
		p.AccessibleEntityTypesCount = known.AccessibleEntityTypesCount
	}
	if known.AssignedUserGroupsCount != nil {
		p.AssignedUserGroupsCount = known.AssignedUserGroupsCount
	}
	if known.AssignedUsersCount != nil {
		p.AssignedUsersCount = known.AssignedUsersCount
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.OperationSchemaChangeImpact != nil {
		p.OperationSchemaChangeImpact = known.OperationSchemaChangeImpact
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessibleClients")
	delete(allFields, "accessibleClientsCount")
	delete(allFields, "accessibleEntityTypes")
	delete(allFields, "accessibleEntityTypesCount")
	delete(allFields, "assignedUserGroupsCount")
	delete(allFields, "assignedUsersCount")
	delete(allFields, "clientName")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "isSystemDefined")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "operationSchemaChangeImpact")
	delete(allFields, "operations")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRoleProjection() *RoleProjection {
	p := new(RoleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.RoleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsSystemDefined = new(bool)
	*p.IsSystemDefined = true

	return p
}

/*
The authorization policies config that will be seeded.
*/
type SeedConfigAuthorizationPolicyList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The authorization policies config that will be seeded.
	*/
	AuthorizationPolicies []AuthorizationPolicyConfig `json:"authorizationPolicies,omitempty"`
}

func (p *SeedConfigAuthorizationPolicyList) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeedConfigAuthorizationPolicyList

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

func (p *SeedConfigAuthorizationPolicyList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeedConfigAuthorizationPolicyList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeedConfigAuthorizationPolicyList()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthorizationPolicies != nil {
		p.AuthorizationPolicies = known.AuthorizationPolicies
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authorizationPolicies")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeedConfigAuthorizationPolicyList() *SeedConfigAuthorizationPolicyList {
	p := new(SeedConfigAuthorizationPolicyList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.SeedConfigAuthorizationPolicyList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The entities config that will be seeded.
*/
type SeedConfigEntityList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The entities config that will be seeded.
	*/
	Entities []EntityConfig `json:"entities,omitempty"`
}

func (p *SeedConfigEntityList) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeedConfigEntityList

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

func (p *SeedConfigEntityList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeedConfigEntityList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeedConfigEntityList()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entities")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeedConfigEntityList() *SeedConfigEntityList {
	p := new(SeedConfigEntityList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.SeedConfigEntityList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The operation objects that will be seeded.
*/
type SeedConfigOperationList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The operation objects that will be seeded.
	*/
	Operations []OperationConfig `json:"operations,omitempty"`
}

func (p *SeedConfigOperationList) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeedConfigOperationList

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

func (p *SeedConfigOperationList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeedConfigOperationList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeedConfigOperationList()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "operations")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeedConfigOperationList() *SeedConfigOperationList {
	p := new(SeedConfigOperationList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.SeedConfigOperationList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of operation to be performed on the entities.
*/
type SeedConfigOperationType int

const (
	SEEDCONFIGOPERATIONTYPE_UNKNOWN  SeedConfigOperationType = 0
	SEEDCONFIGOPERATIONTYPE_REDACTED SeedConfigOperationType = 1
	SEEDCONFIGOPERATIONTYPE_CREATE   SeedConfigOperationType = 2
	SEEDCONFIGOPERATIONTYPE_UPDATE   SeedConfigOperationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SeedConfigOperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SeedConfigOperationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SeedConfigOperationType) index(name string) SeedConfigOperationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATE",
		"UPDATE",
	}
	for idx := range names {
		if names[idx] == name {
			return SeedConfigOperationType(idx)
		}
	}
	return SEEDCONFIGOPERATIONTYPE_UNKNOWN
}

func (e *SeedConfigOperationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SeedConfigOperationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SeedConfigOperationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SeedConfigOperationType) Ref() *SeedConfigOperationType {
	return &e
}

/*
Request body for the configuration seeding request.
*/
type SeedConfigRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the client for which the entities will be seeded.
	*/
	ClientName *string `json:"clientName"`
	/*

	 */
	ConfigListItemDiscriminator_ *string `json:"$configListItemDiscriminator,omitempty"`
	/*
	  List of entities to be seeded.
	*/
	ConfigList *OneOfSeedConfigRequestConfigList `json:"configList"`

	ConfigType *SeedConfigType `json:"configType"`

	OperationType *SeedConfigOperationType `json:"operationType"`
	/*
	  ID of tenant for which configuration is being set up.
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SeedConfigRequest) MarshalJSON() ([]byte, error) {
	type SeedConfigRequestProxy SeedConfigRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SeedConfigRequestProxy
		ClientName    *string                           `json:"clientName,omitempty"`
		ConfigList    *OneOfSeedConfigRequestConfigList `json:"configList,omitempty"`
		ConfigType    *SeedConfigType                   `json:"configType,omitempty"`
		OperationType *SeedConfigOperationType          `json:"operationType,omitempty"`
	}{
		SeedConfigRequestProxy: (*SeedConfigRequestProxy)(p),
		ClientName:             p.ClientName,
		ConfigList:             p.ConfigList,
		ConfigType:             p.ConfigType,
		OperationType:          p.OperationType,
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

func (p *SeedConfigRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeedConfigRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeedConfigRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.ConfigListItemDiscriminator_ != nil {
		p.ConfigListItemDiscriminator_ = known.ConfigListItemDiscriminator_
	}
	if known.ConfigList != nil {
		p.ConfigList = known.ConfigList
	}
	if known.ConfigType != nil {
		p.ConfigType = known.ConfigType
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientName")
	delete(allFields, "$configListItemDiscriminator")
	delete(allFields, "configList")
	delete(allFields, "configType")
	delete(allFields, "operationType")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeedConfigRequest() *SeedConfigRequest {
	p := new(SeedConfigRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.SeedConfigRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SeedConfigRequest) GetConfigList() interface{} {
	if nil == p.ConfigList {
		return nil
	}
	return p.ConfigList.GetValue()
}

func (p *SeedConfigRequest) SetConfigList(v interface{}) error {
	if nil == p.ConfigList {
		p.ConfigList = NewOneOfSeedConfigRequestConfigList()
	}
	e := p.ConfigList.SetValue(v)
	if nil == e {
		if nil == p.ConfigListItemDiscriminator_ {
			p.ConfigListItemDiscriminator_ = new(string)
		}
		*p.ConfigListItemDiscriminator_ = *p.ConfigList.Discriminator
	}
	return e
}

/*
The role objects that will be seeded.
*/
type SeedConfigRoleList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The role objects that will be seeded.
	*/
	Roles []RoleConfig `json:"roles,omitempty"`
}

func (p *SeedConfigRoleList) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SeedConfigRoleList

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

func (p *SeedConfigRoleList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SeedConfigRoleList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSeedConfigRoleList()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Roles != nil {
		p.Roles = known.Roles
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "roles")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSeedConfigRoleList() *SeedConfigRoleList {
	p := new(SeedConfigRoleList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.SeedConfigRoleList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of entities which will be seeded.
*/
type SeedConfigType int

const (
	SEEDCONFIGTYPE_UNKNOWN              SeedConfigType = 0
	SEEDCONFIGTYPE_REDACTED             SeedConfigType = 1
	SEEDCONFIGTYPE_CLIENT               SeedConfigType = 2
	SEEDCONFIGTYPE_ENTITY               SeedConfigType = 3
	SEEDCONFIGTYPE_OPERATION            SeedConfigType = 4
	SEEDCONFIGTYPE_ROLE                 SeedConfigType = 5
	SEEDCONFIGTYPE_AUTHORIZATION_POLICY SeedConfigType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SeedConfigType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT",
		"ENTITY",
		"OPERATION",
		"ROLE",
		"AUTHORIZATION_POLICY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SeedConfigType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT",
		"ENTITY",
		"OPERATION",
		"ROLE",
		"AUTHORIZATION_POLICY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SeedConfigType) index(name string) SeedConfigType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT",
		"ENTITY",
		"OPERATION",
		"ROLE",
		"AUTHORIZATION_POLICY",
	}
	for idx := range names {
		if names[idx] == name {
			return SeedConfigType(idx)
		}
	}
	return SEEDCONFIGTYPE_UNKNOWN
}

func (e *SeedConfigType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SeedConfigType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SeedConfigType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SeedConfigType) Ref() *SeedConfigType {
	return &e
}

/*
Tenant details for tenant to be onboarded to IAM.
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
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of tenant for which configuration is being set up.
	*/
	TenantName *string `json:"tenantName,omitempty"`
}

func (p *Tenant) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Tenant

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

func (p *Tenant) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Tenant
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenant()

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
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TenantName != nil {
		p.TenantName = known.TenantName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "tenantName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.Tenant"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authz/authorization-policies/{extId} Put operation
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

func (p *UpdateAuthorizationPolicyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateAuthorizationPolicyApiResponse

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

func (p *UpdateAuthorizationPolicyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateAuthorizationPolicyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateAuthorizationPolicyApiResponse()

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

func NewUpdateAuthorizationPolicyApiResponse() *UpdateAuthorizationPolicyApiResponse {
	p := new(UpdateAuthorizationPolicyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateAuthorizationPolicyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
REST response for all response codes in API path /iam/v4.1.b2/authz/roles/{extId} Put operation
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

func (p *UpdateRoleApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateRoleApiResponse

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

func (p *UpdateRoleApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateRoleApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateRoleApiResponse()

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

func NewUpdateRoleApiResponse() *UpdateRoleApiResponse {
	p := new(UpdateRoleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UpdateRoleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
The list of authorization policies associated with the user who is making the API call.
*/
type UserAuthorizationPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  Description of the Authorization Policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The display name for the authorization policy.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  The entities being qualified by the authorization policy.
	*/
	Entities []EntityFilter `json:"entities,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The role associated with the authorization policy.
	*/
	Role *string `json:"role,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UserAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserAuthorizationPolicy

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

func (p *UserAuthorizationPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserAuthorizationPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserAuthorizationPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.Entities != nil {
		p.Entities = known.Entities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Role != nil {
		p.Role = known.Role
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientName")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "entities")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "role")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserAuthorizationPolicy() *UserAuthorizationPolicy {
	p := new(UserAuthorizationPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UserAuthorizationPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List the roles associated with the identity/user making the API call
*/
type UserRole struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The list of operations, presented as FQN, that belong to the given role, where FQN is of the format entity_type:client_name:operation_display_name(e.g. role:IAM:View_User_Role).
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UserRole) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserRole

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

func (p *UserRole) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserRole
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserRole()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientName")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "operations")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserRole() *UserRole {
	p := new(UserRole)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UserRole"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type UserRoleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client that created the entity.
	*/
	ClientName *string `json:"clientName,omitempty"`
	/*
	  The display name of the role.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The list of operations, presented as FQN, that belong to the given role, where FQN is of the format entity_type:client_name:operation_display_name(e.g. role:IAM:View_User_Role).
	*/
	Operations []string `json:"operations,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UserRoleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserRoleProjection

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

func (p *UserRoleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserRoleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserRoleProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientName != nil {
		p.ClientName = known.ClientName
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Operations != nil {
		p.Operations = known.Operations
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientName")
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "operations")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserRoleProjection() *UserRoleProjection {
	p := new(UserRoleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authz.UserRoleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2    *interface{}           `json:"-"`
	oneOfType0    *Role                  `json:"-"`
}

func NewOneOfGetRoleApiResponseData() *OneOfGetRoleApiResponseData {
	p := new(OneOfGetRoleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRoleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRoleApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetRoleApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRoleApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRoleApiResponseData"))
}

func (p *OneOfGetRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRoleApiResponseData")
}

type OneOfGetOperationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Operation             `json:"-"`
}

func NewOneOfGetOperationApiResponseData() *OneOfGetOperationApiResponseData {
	p := new(OneOfGetOperationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetOperationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetOperationApiResponseData is nil"))
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

func (p *OneOfGetOperationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetOperationApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetOperationApiResponseData"))
}

func (p *OneOfGetOperationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetOperationApiResponseData")
}

type OneOfGetAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *AuthorizationPolicy   `json:"-"`
}

func NewOneOfGetAuthorizationPolicyApiResponseData() *OneOfGetAuthorizationPolicyApiResponseData {
	p := new(OneOfGetAuthorizationPolicyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAuthorizationPolicyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAuthorizationPolicyApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAuthorizationPolicyApiResponseData"))
}

func (p *OneOfGetAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAuthorizationPolicyApiResponseData")
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

type OneOfSeedConfigRequestConfigList struct {
	Discriminator *string                            `json:"-"`
	ObjectType_   *string                            `json:"-"`
	oneOfType2005 *SeedConfigAuthorizationPolicyList `json:"-"`
	oneOfType2002 *SeedConfigEntityList              `json:"-"`
	oneOfType2001 *ClientConfig                      `json:"-"`
	oneOfType2003 *SeedConfigOperationList           `json:"-"`
	oneOfType2004 *SeedConfigRoleList                `json:"-"`
}

func NewOneOfSeedConfigRequestConfigList() *OneOfSeedConfigRequestConfigList {
	p := new(OneOfSeedConfigRequestConfigList)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSeedConfigRequestConfigList) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSeedConfigRequestConfigList is nil"))
	}
	switch v.(type) {
	case SeedConfigAuthorizationPolicyList:
		if nil == p.oneOfType2005 {
			p.oneOfType2005 = new(SeedConfigAuthorizationPolicyList)
		}
		*p.oneOfType2005 = v.(SeedConfigAuthorizationPolicyList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2005.ObjectType_
	case SeedConfigEntityList:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(SeedConfigEntityList)
		}
		*p.oneOfType2002 = v.(SeedConfigEntityList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	case ClientConfig:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ClientConfig)
		}
		*p.oneOfType2001 = v.(ClientConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case SeedConfigOperationList:
		if nil == p.oneOfType2003 {
			p.oneOfType2003 = new(SeedConfigOperationList)
		}
		*p.oneOfType2003 = v.(SeedConfigOperationList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2003.ObjectType_
	case SeedConfigRoleList:
		if nil == p.oneOfType2004 {
			p.oneOfType2004 = new(SeedConfigRoleList)
		}
		*p.oneOfType2004 = v.(SeedConfigRoleList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2004.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2004.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSeedConfigRequestConfigList) GetValue() interface{} {
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2005
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2003
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2004
	}
	return nil
}

func (p *OneOfSeedConfigRequestConfigList) UnmarshalJSON(b []byte) error {
	vOneOfType2005 := new(SeedConfigAuthorizationPolicyList)
	if err := json.Unmarshal(b, vOneOfType2005); err == nil {
		if "iam.v4.authz.SeedConfigAuthorizationPolicyList" == *vOneOfType2005.ObjectType_ {
			if nil == p.oneOfType2005 {
				p.oneOfType2005 = new(SeedConfigAuthorizationPolicyList)
			}
			*p.oneOfType2005 = *vOneOfType2005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2005.ObjectType_
			return nil
		}
	}
	vOneOfType2002 := new(SeedConfigEntityList)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "iam.v4.authz.SeedConfigEntityList" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(SeedConfigEntityList)
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
	vOneOfType2001 := new(ClientConfig)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "iam.v4.authz.ClientConfig" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ClientConfig)
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
	vOneOfType2003 := new(SeedConfigOperationList)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if "iam.v4.authz.SeedConfigOperationList" == *vOneOfType2003.ObjectType_ {
			if nil == p.oneOfType2003 {
				p.oneOfType2003 = new(SeedConfigOperationList)
			}
			*p.oneOfType2003 = *vOneOfType2003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2003.ObjectType_
			return nil
		}
	}
	vOneOfType2004 := new(SeedConfigRoleList)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if "iam.v4.authz.SeedConfigRoleList" == *vOneOfType2004.ObjectType_ {
			if nil == p.oneOfType2004 {
				p.oneOfType2004 = new(SeedConfigRoleList)
			}
			*p.oneOfType2004 = *vOneOfType2004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2004.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2004.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSeedConfigRequestConfigList"))
}

func (p *OneOfSeedConfigRequestConfigList) MarshalJSON() ([]byte, error) {
	if p.oneOfType2005 != nil && *p.oneOfType2005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2005)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2003 != nil && *p.oneOfType2003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if p.oneOfType2004 != nil && *p.oneOfType2004.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	return nil, errors.New("No value to marshal for OneOfSeedConfigRequestConfigList")
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

type OneOfListOperationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Operation            `json:"-"`
}

func NewOneOfListOperationsApiResponseData() *OneOfListOperationsApiResponseData {
	p := new(OneOfListOperationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListOperationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListOperationsApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListOperationsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authz.Operation>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListOperationsApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListOperationsApiResponseData"))
}

func (p *OneOfListOperationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authz.Operation>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListOperationsApiResponseData")
}

type OneOfCreateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Role                  `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateRoleApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateRoleApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRoleApiResponseData"))
}

func (p *OneOfCreateRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRoleApiResponseData")
}

type OneOfListAuthorizationPoliciesApiResponseData struct {
	Discriminator *string                         `json:"-"`
	ObjectType_   *string                         `json:"-"`
	oneOfType400  *import2.ErrorResponse          `json:"-"`
	oneOfType401  []AuthorizationPolicyProjection `json:"-"`
	oneOfType0    []AuthorizationPolicy           `json:"-"`
}

func NewOneOfListAuthorizationPoliciesApiResponseData() *OneOfListAuthorizationPoliciesApiResponseData {
	p := new(OneOfListAuthorizationPoliciesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListAuthorizationPoliciesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListAuthorizationPoliciesApiResponseData is nil"))
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

func (p *OneOfListAuthorizationPoliciesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authz.AuthorizationPolicyProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<iam.v4.authz.AuthorizationPolicy>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListAuthorizationPoliciesApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListAuthorizationPoliciesApiResponseData"))
}

func (p *OneOfListAuthorizationPoliciesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authz.AuthorizationPolicyProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<iam.v4.authz.AuthorizationPolicy>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListAuthorizationPoliciesApiResponseData")
}

type OneOfListRolesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType401  []RoleProjection       `json:"-"`
	oneOfType0    []Role                 `json:"-"`
}

func NewOneOfListRolesApiResponseData() *OneOfListRolesApiResponseData {
	p := new(OneOfListRolesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRolesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRolesApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRolesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authz.RoleProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<iam.v4.authz.Role>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListRolesApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRolesApiResponseData"))
}

func (p *OneOfListRolesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authz.RoleProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<iam.v4.authz.Role>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListRolesApiResponseData")
}

type OneOfListEntitiesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Entity               `json:"-"`
}

func NewOneOfListEntitiesApiResponseData() *OneOfListEntitiesApiResponseData {
	p := new(OneOfListEntitiesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEntitiesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEntitiesApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListEntitiesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authz.Entity>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListEntitiesApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEntitiesApiResponseData"))
}

func (p *OneOfListEntitiesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authz.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListEntitiesApiResponseData")
}

type OneOfGetEntityApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Entity                `json:"-"`
}

func NewOneOfGetEntityApiResponseData() *OneOfGetEntityApiResponseData {
	p := new(OneOfGetEntityApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEntityApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEntityApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetEntityApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetEntityApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEntityApiResponseData"))
}

func (p *OneOfGetEntityApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetEntityApiResponseData")
}

type OneOfUpdateRoleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateRoleApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateRoleApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRoleApiResponseData"))
}

func (p *OneOfUpdateRoleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRoleApiResponseData")
}

type OneOfCreateAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *AuthorizationPolicy   `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateAuthorizationPolicyApiResponseData"))
}

func (p *OneOfCreateAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateAuthorizationPolicyApiResponseData")
}

type OneOfUpdateAuthorizationPolicyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import3.Message       `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateAuthorizationPolicyApiResponseData"))
}

func (p *OneOfUpdateAuthorizationPolicyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateAuthorizationPolicyApiResponseData")
}

type OneOfGetClientApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Client                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClientApiResponseData() *OneOfGetClientApiResponseData {
	p := new(OneOfGetClientApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClientApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClientApiResponseData is nil"))
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

func (p *OneOfGetClientApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetClientApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClientApiResponseData"))
}

func (p *OneOfGetClientApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetClientApiResponseData")
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
