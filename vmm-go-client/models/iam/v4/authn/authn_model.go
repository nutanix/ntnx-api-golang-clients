/*
 * Generated file models/iam/v4/authn/authn_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module iam.v4.authn of Nutanix Virtual Machine Management APIs
*/
package authn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	"time"
)

/*
Information of bucket access key.
*/
type BucketsAccessKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the bucket access key.
	*/
	AccessKeyName *string `json:"accessKeyName"`
	/*
	  External client to whom this key is allocated.
	*/
	AssignedTo *string `json:"assignedTo,omitempty"`
	/*
	  Service account user who created the buckets access key.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time for the bucket access key.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	CreationType *CreationType `json:"creationType,omitempty"`
	/*
	  The expiry time of the buckets access Key.
	*/
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Entity that updated the buckets access key.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  Creation time for the bucket access key.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  This represents secret access key, which will be returned only during access key creation.
	*/
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	Status *BucketsAccessKeyStatusType `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  User identifier who owns the bucket access key.
	*/
	UserId *string `json:"userId,omitempty"`
}

func (p *BucketsAccessKey) MarshalJSON() ([]byte, error) {
	type BucketsAccessKeyProxy BucketsAccessKey

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BucketsAccessKeyProxy
		AccessKeyName *string `json:"accessKeyName,omitempty"`
	}{
		BucketsAccessKeyProxy: (*BucketsAccessKeyProxy)(p),
		AccessKeyName:         p.AccessKeyName,
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

func (p *BucketsAccessKey) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BucketsAccessKey
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BucketsAccessKey(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessKeyName")
	delete(allFields, "assignedTo")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "creationType")
	delete(allFields, "expiryTime")
	delete(allFields, "extId")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "secretAccessKey")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "userId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBucketsAccessKey() *BucketsAccessKey {
	p := new(BucketsAccessKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.BucketsAccessKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The status of the buckets access key, that is, valid, expired or revoked.
*/
type BucketsAccessKeyStatusType int

const (
	BUCKETSACCESSKEYSTATUSTYPE_UNKNOWN  BucketsAccessKeyStatusType = 0
	BUCKETSACCESSKEYSTATUSTYPE_REDACTED BucketsAccessKeyStatusType = 1
	BUCKETSACCESSKEYSTATUSTYPE_VALID    BucketsAccessKeyStatusType = 2
	BUCKETSACCESSKEYSTATUSTYPE_REVOKED  BucketsAccessKeyStatusType = 3
	BUCKETSACCESSKEYSTATUSTYPE_EXPIRED  BucketsAccessKeyStatusType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BucketsAccessKeyStatusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"REVOKED",
		"EXPIRED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BucketsAccessKeyStatusType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"REVOKED",
		"EXPIRED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BucketsAccessKeyStatusType) index(name string) BucketsAccessKeyStatusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"REVOKED",
		"EXPIRED",
	}
	for idx := range names {
		if names[idx] == name {
			return BucketsAccessKeyStatusType(idx)
		}
	}
	return BUCKETSACCESSKEYSTATUSTYPE_UNKNOWN
}

func (e *BucketsAccessKeyStatusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BucketsAccessKeyStatusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BucketsAccessKeyStatusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BucketsAccessKeyStatusType) Ref() *BucketsAccessKeyStatusType {
	return &e
}

/*
The creation mechanism of this entity.
*/
type CreationType int

const (
	CREATIONTYPE_UNKNOWN        CreationType = 0
	CREATIONTYPE_REDACTED       CreationType = 1
	CREATIONTYPE_PREDEFINED     CreationType = 2
	CREATIONTYPE_USERDEFINED    CreationType = 3
	CREATIONTYPE_SERVICEDEFINED CreationType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CreationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PREDEFINED",
		"USERDEFINED",
		"SERVICEDEFINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CreationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PREDEFINED",
		"USERDEFINED",
		"SERVICEDEFINED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CreationType) index(name string) CreationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PREDEFINED",
		"USERDEFINED",
		"SERVICEDEFINED",
	}
	for idx := range names {
		if names[idx] == name {
			return CreationType(idx)
		}
	}
	return CREATIONTYPE_UNKNOWN
}

func (e *CreationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CreationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CreationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CreationType) Ref() *CreationType {
	return &e
}

/*
Information of the user.
*/
type User struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates additional attributes of the user.
	*/
	AdditionalAttributes []import2.KVPair `json:"additionalAttributes,omitempty"`
	/*
	  Bucket access keys for the user.
	*/
	BucketsAccessKeys []BucketsAccessKey `json:"bucketsAccessKeys,omitempty"`
	/*
	  User or Service who created the user.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the user.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	CreationType *CreationType `json:"creationType,omitempty"`
	/*
	  Description of the user.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Display name of the user. For LDAP and SAML users, this is set from AD config.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Email ID of the user.
	*/
	EmailId *string `json:"emailId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  First name of the user.
	*/
	FirstName *string `json:"firstName,omitempty"`
	/*
	  Identifier of the IDP for the user.
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Flag to force the user to reset password.
	*/
	IsForceResetPasswordEnabled *bool `json:"isForceResetPasswordEnabled,omitempty"`
	/*
	  The last successful login time for the user.
	*/
	LastLoginTime *time.Time `json:"lastLoginTime,omitempty"`
	/*
	  Last name of the user.
	*/
	LastName *string `json:"lastName,omitempty"`
	/*
	  Last updated by this user ID.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  The last updated time for the user.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Default locale of the user.
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  Middle name of the user.
	*/
	MiddleInitial *string `json:"middleInitial,omitempty"`
	/*
	  Password of the user.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Default region of the user.
	*/
	Region *string `json:"region,omitempty"`

	Status *UserStatusType `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserType *UserType `json:"userType,omitempty"`
	/*
	  Identifier of the user.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *User) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias User

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

func (p *User) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias User
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = User(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "additionalAttributes")
	delete(allFields, "bucketsAccessKeys")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "creationType")
	delete(allFields, "description")
	delete(allFields, "displayName")
	delete(allFields, "emailId")
	delete(allFields, "extId")
	delete(allFields, "firstName")
	delete(allFields, "idpId")
	delete(allFields, "isForceResetPasswordEnabled")
	delete(allFields, "lastLoginTime")
	delete(allFields, "lastName")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "locale")
	delete(allFields, "middleInitial")
	delete(allFields, "password")
	delete(allFields, "region")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "userType")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUser() *User {
	p := new(User)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.User"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the user.
*/
type UserStatusType int

const (
	USERSTATUSTYPE_UNKNOWN  UserStatusType = 0
	USERSTATUSTYPE_REDACTED UserStatusType = 1
	USERSTATUSTYPE_ACTIVE   UserStatusType = 2
	USERSTATUSTYPE_INACTIVE UserStatusType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UserStatusType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"INACTIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UserStatusType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"INACTIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UserStatusType) index(name string) UserStatusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"INACTIVE",
	}
	for idx := range names {
		if names[idx] == name {
			return UserStatusType(idx)
		}
	}
	return USERSTATUSTYPE_UNKNOWN
}

func (e *UserStatusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UserStatusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UserStatusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UserStatusType) Ref() *UserStatusType {
	return &e
}

/*
User type like SAML user or local user, and so on.
*/
type UserType int

const (
	USERTYPE_UNKNOWN         UserType = 0
	USERTYPE_REDACTED        UserType = 1
	USERTYPE_LOCAL           UserType = 2
	USERTYPE_SAML            UserType = 3
	USERTYPE_LDAP            UserType = 4
	USERTYPE_EXTERNAL        UserType = 5
	USERTYPE_SERVICE_ACCOUNT UserType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UserType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"SAML",
		"LDAP",
		"EXTERNAL",
		"SERVICE_ACCOUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UserType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"SAML",
		"LDAP",
		"EXTERNAL",
		"SERVICE_ACCOUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UserType) index(name string) UserType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"SAML",
		"LDAP",
		"EXTERNAL",
		"SERVICE_ACCOUNT",
	}
	for idx := range names {
		if names[idx] == name {
			return UserType(idx)
		}
	}
	return USERTYPE_UNKNOWN
}

func (e *UserType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UserType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UserType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UserType) Ref() *UserType {
	return &e
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
