/*
 * Generated file models/iam/v4/authn/authn_model.go.
 *
 * Product version: 4.1.1-beta-2
 *
 * Part of the Nutanix Identity and Access Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  User , token and identity management
*/
package authn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/error"
	"time"
)

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{extId}/$actions/change-state Post operation
*/
type ActivateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfActivateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ActivateUserApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ActivateUserApiResponse

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

func (p *ActivateUserApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ActivateUserApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewActivateUserApiResponse()

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

func NewActivateUserApiResponse() *ActivateUserApiResponse {
	p := new(ActivateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ActivateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ActivateUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ActivateUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfActivateUserApiResponseData()
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
Algorithm types for OIDC key.
*/
type AlgoType int

const (
	ALGOTYPE_UNKNOWN  AlgoType = 0
	ALGOTYPE_REDACTED AlgoType = 1
	ALGOTYPE_RSA256   AlgoType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AlgoType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AlgoType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AlgoType) index(name string) AlgoType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA256",
	}
	for idx := range names {
		if names[idx] == name {
			return AlgoType(idx)
		}
	}
	return ALGOTYPE_UNKNOWN
}

func (e *AlgoType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AlgoType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AlgoType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AlgoType) Ref() *AlgoType {
	return &e
}

/*
The API key details.
*/
type ApiKeyDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The actual API key value, returned only during creation.
	*/
	ApiKey *string `json:"apiKey,omitempty"`
}

func (p *ApiKeyDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiKeyDetails

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

func (p *ApiKeyDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiKeyDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiKeyDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApiKey != nil {
		p.ApiKey = known.ApiKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiKeyDetails() *ApiKeyDetails {
	p := new(ApiKeyDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ApiKeyDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Supported authentication methods.
*/
type AuthMethodType int

const (
	AUTHMETHODTYPE_UNKNOWN             AuthMethodType = 0
	AUTHMETHODTYPE_REDACTED            AuthMethodType = 1
	AUTHMETHODTYPE_CLIENT_SECRET_BASIC AuthMethodType = 2
	AUTHMETHODTYPE_PRIVATE_KEY_JWT     AuthMethodType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthMethodType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT_SECRET_BASIC",
		"PRIVATE_KEY_JWT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthMethodType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT_SECRET_BASIC",
		"PRIVATE_KEY_JWT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthMethodType) index(name string) AuthMethodType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLIENT_SECRET_BASIC",
		"PRIVATE_KEY_JWT",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthMethodType(idx)
		}
	}
	return AUTHMETHODTYPE_UNKNOWN
}

func (e *AuthMethodType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthMethodType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthMethodType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthMethodType) Ref() *AuthMethodType {
	return &e
}

/*
Authentication response from validation APIs.
*/
type AuthenticationResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Auth Code URI for redirection to mercury.
	*/
	AuthCodeUri *string `json:"authCodeUri,omitempty"`
}

func (p *AuthenticationResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AuthenticationResponse

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

func (p *AuthenticationResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthenticationResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthenticationResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthCodeUri != nil {
		p.AuthCodeUri = known.AuthCodeUri
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authCodeUri")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAuthenticationResponse() *AuthenticationResponse {
	p := new(AuthenticationResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.AuthenticationResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The authN migration API request body.
*/
type AuthnMigrationRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MigrationType *AuthnMigrationType `json:"migrationType"`
}

func (p *AuthnMigrationRequest) MarshalJSON() ([]byte, error) {
	type AuthnMigrationRequestProxy AuthnMigrationRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AuthnMigrationRequestProxy
		MigrationType *AuthnMigrationType `json:"migrationType,omitempty"`
	}{
		AuthnMigrationRequestProxy: (*AuthnMigrationRequestProxy)(p),
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

func (p *AuthnMigrationRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AuthnMigrationRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAuthnMigrationRequest()

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

func NewAuthnMigrationRequest() *AuthnMigrationRequest {
	p := new(AuthnMigrationRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.AuthnMigrationRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of authn migration.
*/
type AuthnMigrationType int

const (
	AUTHNMIGRATIONTYPE_UNKNOWN                 AuthnMigrationType = 0
	AUTHNMIGRATIONTYPE_REDACTED                AuthnMigrationType = 1
	AUTHNMIGRATIONTYPE_HASH_OIDC_CLIENT_SECRET AuthnMigrationType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AuthnMigrationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HASH_OIDC_CLIENT_SECRET",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AuthnMigrationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HASH_OIDC_CLIENT_SECRET",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AuthnMigrationType) index(name string) AuthnMigrationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HASH_OIDC_CLIENT_SECRET",
	}
	for idx := range names {
		if names[idx] == name {
			return AuthnMigrationType(idx)
		}
	}
	return AUTHNMIGRATIONTYPE_UNKNOWN
}

func (e *AuthnMigrationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AuthnMigrationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AuthnMigrationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AuthnMigrationType) Ref() *AuthnMigrationType {
	return &e
}

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
	Links []import2.ApiLink `json:"links,omitempty"`
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
	*p = *NewBucketsAccessKey()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessKeyName != nil {
		p.AccessKeyName = known.AccessKeyName
	}
	if known.AssignedTo != nil {
		p.AssignedTo = known.AssignedTo
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CreationType != nil {
		p.CreationType = known.CreationType
	}
	if known.ExpiryTime != nil {
		p.ExpiryTime = known.ExpiryTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.SecretAccessKey != nil {
		p.SecretAccessKey = known.SecretAccessKey
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UserId != nil {
		p.UserId = known.UserId
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBucketsAccessKey() *BucketsAccessKey {
	p := new(BucketsAccessKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.BucketsAccessKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
A certificate-based authentication provider.
*/
type CertAuthProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the uploaded CA chain file.
	*/
	CaCertFileName *string `json:"caCertFileName,omitempty"`

	CertRevocationInfo *CertRevocationInfo `json:"certRevocationInfo,omitempty"`
	/*
	  CA chain file.
	*/
	ClientCaChain *string `json:"clientCaChain,omitempty"`
	/*
	  User or service who created the certificate authentication provider.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the certificate authentication provider.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  UUID of an existing directory service.
	*/
	DirSvcExtID *string `json:"dirSvcExtID,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag to enable/disable certificate authentication for the current certificate-based authentication provider.
	*/
	IsCacEnabled *bool `json:"isCacEnabled,omitempty"`
	/*
	  Flag to enable/disable CAC for the current certificate-based authentication provider.
	*/
	IsCertAuthEnabled *bool `json:"isCertAuthEnabled,omitempty"`
	/*
	  Last updated time of the certificate authentication provider.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the certificate-based authentication provider.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CertAuthProvider) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CertAuthProvider

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

func (p *CertAuthProvider) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CertAuthProvider
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCertAuthProvider()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CaCertFileName != nil {
		p.CaCertFileName = known.CaCertFileName
	}
	if known.CertRevocationInfo != nil {
		p.CertRevocationInfo = known.CertRevocationInfo
	}
	if known.ClientCaChain != nil {
		p.ClientCaChain = known.ClientCaChain
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DirSvcExtID != nil {
		p.DirSvcExtID = known.DirSvcExtID
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsCacEnabled != nil {
		p.IsCacEnabled = known.IsCacEnabled
	}
	if known.IsCertAuthEnabled != nil {
		p.IsCertAuthEnabled = known.IsCertAuthEnabled
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
	delete(allFields, "caCertFileName")
	delete(allFields, "certRevocationInfo")
	delete(allFields, "clientCaChain")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "dirSvcExtID")
	delete(allFields, "extId")
	delete(allFields, "isCacEnabled")
	delete(allFields, "isCertAuthEnabled")
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

func NewCertAuthProvider() *CertAuthProvider {
	p := new(CertAuthProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertAuthProvider"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration details used for determining if the client certificate is revoked.
*/
type CertRevocationInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of the CRL distribution points which will be used to fetch the CRLs. For the POST/PUT APIs, you must add this attribute directly to the payload, not under the certRevocationInfo object.
	*/
	CrlDps []string `json:"crlDps,omitempty"`
	/*
	  Interval in seconds at which the CRL should be fetched from the CRLDP. The default is 86400 seconds(1 day). For the POST/PUT APIs, you must add this attribute directly to the payload, not under the certRevocationInfo object.
	*/
	GlobalCrlRefreshInterval *int `json:"globalCrlRefreshInterval,omitempty"`
	/*
	  Flag to enable/disable CRL revocation check. For the POST/PUT APIs, you must add this attribute directly to the payload, not under the certRevocationInfo object.
	*/
	IsCrlEnabled *bool `json:"isCrlEnabled,omitempty"`
	/*
	  Flag to enable/disable OCSP revocation check. For the POST/PUT APIs, you must add this attribute directly to the payload, not under the certRevocationInfo object.
	*/
	IsOcspEnabled *bool `json:"isOcspEnabled,omitempty"`
	/*
	  URL of the OCSP responder used to override the URL from AIA extension. For the POST/PUT APIs, you must add this attribute directly to the payload, not under the certRevocationInfo object.
	*/
	OcspResponder *string `json:"ocspResponder,omitempty"`
}

func (p *CertRevocationInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CertRevocationInfo

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

func (p *CertRevocationInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CertRevocationInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCertRevocationInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CrlDps != nil {
		p.CrlDps = known.CrlDps
	}
	if known.GlobalCrlRefreshInterval != nil {
		p.GlobalCrlRefreshInterval = known.GlobalCrlRefreshInterval
	}
	if known.IsCrlEnabled != nil {
		p.IsCrlEnabled = known.IsCrlEnabled
	}
	if known.IsOcspEnabled != nil {
		p.IsOcspEnabled = known.IsOcspEnabled
	}
	if known.OcspResponder != nil {
		p.OcspResponder = known.OcspResponder
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "crlDps")
	delete(allFields, "globalCrlRefreshInterval")
	delete(allFields, "isCrlEnabled")
	delete(allFields, "isOcspEnabled")
	delete(allFields, "ocspResponder")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCertRevocationInfo() *CertRevocationInfo {
	p := new(CertRevocationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertRevocationInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The certificate login API request body.
*/
type CertificateLoginRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The session ID value tracking the current login session.
	*/
	RelayState *string `json:"relayState"`
}

func (p *CertificateLoginRequest) MarshalJSON() ([]byte, error) {
	type CertificateLoginRequestProxy CertificateLoginRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CertificateLoginRequestProxy
		RelayState *string `json:"relayState,omitempty"`
	}{
		CertificateLoginRequestProxy: (*CertificateLoginRequestProxy)(p),
		RelayState:                   p.RelayState,
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

func (p *CertificateLoginRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CertificateLoginRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCertificateLoginRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RelayState != nil {
		p.RelayState = known.RelayState
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "relayState")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCertificateLoginRequest() *CertificateLoginRequest {
	p := new(CertificateLoginRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertificateLoginRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/users/$actions/change-password Post operation
*/
type ChangeUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangeUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ChangeUserPasswordApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ChangeUserPasswordApiResponse

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

func (p *ChangeUserPasswordApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ChangeUserPasswordApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewChangeUserPasswordApiResponse()

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

func NewChangeUserPasswordApiResponse() *ChangeUserPasswordApiResponse {
	p := new(ChangeUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ChangeUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ChangeUserPasswordApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ChangeUserPasswordApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfChangeUserPasswordApiResponseData()
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
Supported claim types.
*/
type ClaimsType int

const (
	CLAIMSTYPE_UNKNOWN      ClaimsType = 0
	CLAIMSTYPE_REDACTED     ClaimsType = 1
	CLAIMSTYPE_UUID         ClaimsType = 2
	CLAIMSTYPE_TENANT_UUID  ClaimsType = 3
	CLAIMSTYPE_SUB          ClaimsType = 4
	CLAIMSTYPE_EMAIL        ClaimsType = 5
	CLAIMSTYPE_DISPLAY_NAME ClaimsType = 6
	CLAIMSTYPE_GROUPS       ClaimsType = 7
	CLAIMSTYPE_ISS          ClaimsType = 8
	CLAIMSTYPE_AUD          ClaimsType = 9
	CLAIMSTYPE_EXP          ClaimsType = 10
	CLAIMSTYPE_IAT          ClaimsType = 11
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClaimsType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UUID",
		"TENANT_UUID",
		"SUB",
		"EMAIL",
		"DISPLAY_NAME",
		"GROUPS",
		"ISS",
		"AUD",
		"EXP",
		"IAT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClaimsType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UUID",
		"TENANT_UUID",
		"SUB",
		"EMAIL",
		"DISPLAY_NAME",
		"GROUPS",
		"ISS",
		"AUD",
		"EXP",
		"IAT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClaimsType) index(name string) ClaimsType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UUID",
		"TENANT_UUID",
		"SUB",
		"EMAIL",
		"DISPLAY_NAME",
		"GROUPS",
		"ISS",
		"AUD",
		"EXP",
		"IAT",
	}
	for idx := range names {
		if names[idx] == name {
			return ClaimsType(idx)
		}
	}
	return CLAIMSTYPE_UNKNOWN
}

func (e *ClaimsType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClaimsType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClaimsType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClaimsType) Ref() *ClaimsType {
	return &e
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services/{extId}/$actions/verify-connection-status Post operation
*/
type ConnectionDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConnectionDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ConnectionDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ConnectionDirectoryServiceApiResponse

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

func (p *ConnectionDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ConnectionDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewConnectionDirectoryServiceApiResponse()

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

func NewConnectionDirectoryServiceApiResponse() *ConnectionDirectoryServiceApiResponse {
	p := new(ConnectionDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ConnectionDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ConnectionDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ConnectionDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfConnectionDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/cert-auth-providers Post operation
*/
type CreateCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateCertAuthProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateCertAuthProviderApiResponse

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

func (p *CreateCertAuthProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateCertAuthProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateCertAuthProviderApiResponse()

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

func NewCreateCertAuthProviderApiResponse() *CreateCertAuthProviderApiResponse {
	p := new(CreateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCertAuthProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCertAuthProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCertAuthProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services Post operation
*/
type CreateDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateDirectoryServiceApiResponse

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

func (p *CreateDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateDirectoryServiceApiResponse()

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

func NewCreateDirectoryServiceApiResponse() *CreateDirectoryServiceApiResponse {
	p := new(CreateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{userExtId}/keys Post operation
*/
type CreateKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateKeyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateKeyApiResponse

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

func (p *CreateKeyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateKeyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateKeyApiResponse()

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

func NewCreateKeyApiResponse() *CreateKeyApiResponse {
	p := new(CreateKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers Post operation
*/
type CreateSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateSamlIdentityProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateSamlIdentityProviderApiResponse

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

func (p *CreateSamlIdentityProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateSamlIdentityProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateSamlIdentityProviderApiResponse()

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

func NewCreateSamlIdentityProviderApiResponse() *CreateSamlIdentityProviderApiResponse {
	p := new(CreateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateSamlIdentityProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateSamlIdentityProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateSamlIdentityProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users Post operation
*/
type CreateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateUserApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateUserApiResponse

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

func (p *CreateUserApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateUserApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateUserApiResponse()

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

func NewCreateUserApiResponse() *CreateUserApiResponse {
	p := new(CreateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateUserApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/user-groups Post operation
*/
type CreateUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateUserGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateUserGroupApiResponse

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

func (p *CreateUserGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateUserGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateUserGroupApiResponse()

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

func NewCreateUserGroupApiResponse() *CreateUserGroupApiResponse {
	p := new(CreateUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateUserGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateUserGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateUserGroupApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/cert-auth-providers/{extId} Delete operation
*/
type DeleteCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteCertAuthProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteCertAuthProviderApiResponse

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

func (p *DeleteCertAuthProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteCertAuthProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteCertAuthProviderApiResponse()

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

func NewDeleteCertAuthProviderApiResponse() *DeleteCertAuthProviderApiResponse {
	p := new(DeleteCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteCertAuthProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteCertAuthProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteCertAuthProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services/{extId} Delete operation
*/
type DeleteDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteDirectoryServiceApiResponse

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

func (p *DeleteDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteDirectoryServiceApiResponse()

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

func NewDeleteDirectoryServiceApiResponse() *DeleteDirectoryServiceApiResponse {
	p := new(DeleteDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers/{extId} Delete operation
*/
type DeleteSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteSamlIdentityProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteSamlIdentityProviderApiResponse

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

func (p *DeleteSamlIdentityProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteSamlIdentityProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteSamlIdentityProviderApiResponse()

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

func NewDeleteSamlIdentityProviderApiResponse() *DeleteSamlIdentityProviderApiResponse {
	p := new(DeleteSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSamlIdentityProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSamlIdentityProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSamlIdentityProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/user-groups/{extId} Delete operation
*/
type DeleteUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteUserGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteUserGroupApiResponse

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

func (p *DeleteUserGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteUserGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteUserGroupApiResponse()

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

func NewDeleteUserGroupApiResponse() *DeleteUserGroupApiResponse {
	p := new(DeleteUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteUserGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteUserGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteUserGroupApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{userExtId}/keys/{extId} Delete operation
*/
type DeleteUserKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUserKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteUserKeyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteUserKeyApiResponse

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

func (p *DeleteUserKeyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteUserKeyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteUserKeyApiResponse()

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

func NewDeleteUserKeyApiResponse() *DeleteUserKeyApiResponse {
	p := new(DeleteUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteUserKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteUserKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteUserKeyApiResponseData()
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
Information of a directory service.
*/
type DirectoryService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the directory service.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the directory service.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DirectoryType *DirectoryType `json:"directoryType,omitempty"`
	/*
	  Domain name for the directory service.
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupSearchType *GroupSearchType `json:"groupSearchType,omitempty"`
	/*
	  Last updated time of the directory service.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name for the directory service.
	*/
	Name *string `json:"name,omitempty"`

	OpenLdapConfiguration *OpenLdapConfig `json:"openLdapConfiguration,omitempty"`
	/*
	  Secondary URL for the directory service.
	*/
	SecondaryUrls []string `json:"secondaryUrls,omitempty"`

	ServiceAccount *DsServiceAccount `json:"serviceAccount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  URL for the directory service.
	*/
	Url *string `json:"url,omitempty"`
	/*
	  List of allowed user groups for the directory service.
	*/
	WhiteListedGroups []string `json:"whiteListedGroups,omitempty"`
}

func (p *DirectoryService) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DirectoryService

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

func (p *DirectoryService) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryService
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryService()

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
	if known.DirectoryType != nil {
		p.DirectoryType = known.DirectoryType
	}
	if known.DomainName != nil {
		p.DomainName = known.DomainName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.GroupSearchType != nil {
		p.GroupSearchType = known.GroupSearchType
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
	if known.OpenLdapConfiguration != nil {
		p.OpenLdapConfiguration = known.OpenLdapConfiguration
	}
	if known.SecondaryUrls != nil {
		p.SecondaryUrls = known.SecondaryUrls
	}
	if known.ServiceAccount != nil {
		p.ServiceAccount = known.ServiceAccount
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Url != nil {
		p.Url = known.Url
	}
	if known.WhiteListedGroups != nil {
		p.WhiteListedGroups = known.WhiteListedGroups
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "directoryType")
	delete(allFields, "domainName")
	delete(allFields, "extId")
	delete(allFields, "groupSearchType")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "openLdapConfiguration")
	delete(allFields, "secondaryUrls")
	delete(allFields, "serviceAccount")
	delete(allFields, "tenantId")
	delete(allFields, "url")
	delete(allFields, "whiteListedGroups")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryService() *DirectoryService {
	p := new(DirectoryService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryService"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for directory service connection request.
*/
type DirectoryServiceConnectionRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password to connect to the directory service.
	*/
	Password *string `json:"password"`
	/*
	  Username to connect to the directory service.
	*/
	Username *string `json:"username"`
}

func (p *DirectoryServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	type DirectoryServiceConnectionRequestProxy DirectoryServiceConnectionRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceConnectionRequestProxy
		Password *string `json:"password,omitempty"`
		Username *string `json:"username,omitempty"`
	}{
		DirectoryServiceConnectionRequestProxy: (*DirectoryServiceConnectionRequestProxy)(p),
		Password:                               p.Password,
		Username:                               p.Username,
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

func (p *DirectoryServiceConnectionRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceConnectionRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceConnectionRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceConnectionRequest() *DirectoryServiceConnectionRequest {
	p := new(DirectoryServiceConnectionRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceConnectionRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Verify directory service connection response object.
*/
type DirectoryServiceConnectionResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *DirectoryServiceConnectionResponse) MarshalJSON() ([]byte, error) {
	type DirectoryServiceConnectionResponseProxy DirectoryServiceConnectionResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceConnectionResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		DirectoryServiceConnectionResponseProxy: (*DirectoryServiceConnectionResponseProxy)(p),
		Message:                                 p.Message,
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

func (p *DirectoryServiceConnectionResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceConnectionResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceConnectionResponse()

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

func NewDirectoryServiceConnectionResponse() *DirectoryServiceConnectionResponse {
	p := new(DirectoryServiceConnectionResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceConnectionResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
AD/LDAP information of the user.
*/
type DirectoryServiceInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of AD/LDAP groups having the user.
	*/
	Groups []DirectoryServiceInfoGroup `json:"groups"`
	/*
	  List of AD/LDAP OUs having the given user.
	*/
	Ous []DirectoryServiceInfoOu `json:"ous"`
	/*
	  External identifier of a user.
	*/
	UserId *string `json:"userId"`
}

func (p *DirectoryServiceInfo) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoProxy DirectoryServiceInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceInfoProxy
		Groups []DirectoryServiceInfoGroup `json:"groups,omitempty"`
		Ous    []DirectoryServiceInfoOu    `json:"ous,omitempty"`
		UserId *string                     `json:"userId,omitempty"`
	}{
		DirectoryServiceInfoProxy: (*DirectoryServiceInfoProxy)(p),
		Groups:                    p.Groups,
		Ous:                       p.Ous,
		UserId:                    p.UserId,
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

func (p *DirectoryServiceInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.Ous != nil {
		p.Ous = known.Ous
	}
	if known.UserId != nil {
		p.UserId = known.UserId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "groups")
	delete(allFields, "ous")
	delete(allFields, "userId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceInfo() *DirectoryServiceInfo {
	p := new(DirectoryServiceInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of AD group having the user.
*/
type DirectoryServiceInfoGroup struct {
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
	  Name of the group.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DirectoryServiceInfoGroup) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoGroupProxy DirectoryServiceInfoGroup

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceInfoGroupProxy
		Name *string `json:"name,omitempty"`
	}{
		DirectoryServiceInfoGroupProxy: (*DirectoryServiceInfoGroupProxy)(p),
		Name:                           p.Name,
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

func (p *DirectoryServiceInfoGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceInfoGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceInfoGroup()

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
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceInfoGroup() *DirectoryServiceInfoGroup {
	p := new(DirectoryServiceInfoGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfoGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of OUs having the user.
*/
type DirectoryServiceInfoOu struct {
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
	  Name of the OU.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DirectoryServiceInfoOu) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoOuProxy DirectoryServiceInfoOu

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceInfoOuProxy
		Name *string `json:"name,omitempty"`
	}{
		DirectoryServiceInfoOuProxy: (*DirectoryServiceInfoOuProxy)(p),
		Name:                        p.Name,
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

func (p *DirectoryServiceInfoOu) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceInfoOu
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceInfoOu()

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
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceInfoOu() *DirectoryServiceInfoOu {
	p := new(DirectoryServiceInfoOu)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfoOu"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of searched attributes.
*/
type DirectoryServiceSearchAttribute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the attribute.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Value of the attribute.
	*/
	Values []string `json:"values,omitempty"`
}

func (p *DirectoryServiceSearchAttribute) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DirectoryServiceSearchAttribute

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

func (p *DirectoryServiceSearchAttribute) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceSearchAttribute
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceSearchAttribute()

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
	if known.Values != nil {
		p.Values = known.Values
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "values")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceSearchAttribute() *DirectoryServiceSearchAttribute {
	p := new(DirectoryServiceSearchAttribute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchAttribute"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of single search entity.
*/
type DirectoryServiceSearchEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Attributes []DirectoryServiceSearchAttribute `json:"attributes,omitempty"`
	/*
	  Type of entity either user or group.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  Name of the entity in canonical format.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *DirectoryServiceSearchEntity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DirectoryServiceSearchEntity

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

func (p *DirectoryServiceSearchEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceSearchEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceSearchEntity()

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
	if known.EntityType != nil {
		p.EntityType = known.EntityType
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "attributes")
	delete(allFields, "entityType")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceSearchEntity() *DirectoryServiceSearchEntity {
	p := new(DirectoryServiceSearchEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for directory search query.
*/
type DirectoryServiceSearchQuery struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag indicating whether the search should be a wildcard search or not.
	*/
	IsWildcardSearch *bool `json:"isWildcardSearch,omitempty"`
	/*
	  Query string for directory service search.
	*/
	Query *string `json:"query"`
	/*
	  Attributes returned by the search operation.
	*/
	ReturnedAttributes []string `json:"returnedAttributes,omitempty"`
	/*
	  Attributes for search operation. By default, the search will be performed with a common name.
	*/
	SearchedAttributes []string `json:"searchedAttributes,omitempty"`
}

func (p *DirectoryServiceSearchQuery) MarshalJSON() ([]byte, error) {
	type DirectoryServiceSearchQueryProxy DirectoryServiceSearchQuery

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DirectoryServiceSearchQueryProxy
		Query *string `json:"query,omitempty"`
	}{
		DirectoryServiceSearchQueryProxy: (*DirectoryServiceSearchQueryProxy)(p),
		Query:                            p.Query,
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

func (p *DirectoryServiceSearchQuery) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceSearchQuery
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceSearchQuery()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsWildcardSearch != nil {
		p.IsWildcardSearch = known.IsWildcardSearch
	}
	if known.Query != nil {
		p.Query = known.Query
	}
	if known.ReturnedAttributes != nil {
		p.ReturnedAttributes = known.ReturnedAttributes
	}
	if known.SearchedAttributes != nil {
		p.SearchedAttributes = known.SearchedAttributes
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isWildcardSearch")
	delete(allFields, "query")
	delete(allFields, "returnedAttributes")
	delete(allFields, "searchedAttributes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceSearchQuery() *DirectoryServiceSearchQuery {
	p := new(DirectoryServiceSearchQuery)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchQuery"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsWildcardSearch = new(bool)
	*p.IsWildcardSearch = true

	return p
}

/*
Information of directory service search result.
*/
type DirectoryServiceSearchResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Domain name for the directory service.
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  Result of directory service search.
	*/
	SearchResults []DirectoryServiceSearchEntity `json:"searchResults,omitempty"`
}

func (p *DirectoryServiceSearchResult) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DirectoryServiceSearchResult

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

func (p *DirectoryServiceSearchResult) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DirectoryServiceSearchResult
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDirectoryServiceSearchResult()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DomainName != nil {
		p.DomainName = known.DomainName
	}
	if known.SearchResults != nil {
		p.SearchResults = known.SearchResults
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "domainName")
	delete(allFields, "searchResults")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDirectoryServiceSearchResult() *DirectoryServiceSearchResult {
	p := new(DirectoryServiceSearchResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of directory service.
*/
type DirectoryType int

const (
	DIRECTORYTYPE_UNKNOWN          DirectoryType = 0
	DIRECTORYTYPE_REDACTED         DirectoryType = 1
	DIRECTORYTYPE_ACTIVE_DIRECTORY DirectoryType = 2
	DIRECTORYTYPE_OPEN_LDAP        DirectoryType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DirectoryType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_DIRECTORY",
		"OPEN_LDAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DirectoryType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_DIRECTORY",
		"OPEN_LDAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DirectoryType) index(name string) DirectoryType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE_DIRECTORY",
		"OPEN_LDAP",
	}
	for idx := range names {
		if names[idx] == name {
			return DirectoryType(idx)
		}
	}
	return DIRECTORYTYPE_UNKNOWN
}

func (e *DirectoryType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DirectoryType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DirectoryType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DirectoryType) Ref() *DirectoryType {
	return &e
}

/*
Supported subject types.
*/
type DiscoverySubjectType int

const (
	DISCOVERYSUBJECTTYPE_UNKNOWN  DiscoverySubjectType = 0
	DISCOVERYSUBJECTTYPE_REDACTED DiscoverySubjectType = 1
	DISCOVERYSUBJECTTYPE_PUBLIC   DiscoverySubjectType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DiscoverySubjectType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PUBLIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DiscoverySubjectType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PUBLIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DiscoverySubjectType) index(name string) DiscoverySubjectType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PUBLIC",
	}
	for idx := range names {
		if names[idx] == name {
			return DiscoverySubjectType(idx)
		}
	}
	return DISCOVERYSUBJECTTYPE_UNKNOWN
}

func (e *DiscoverySubjectType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DiscoverySubjectType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DiscoverySubjectType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DiscoverySubjectType) Ref() *DiscoverySubjectType {
	return &e
}

/*
Information on the service account to connect to the directory service.
*/
type DsServiceAccount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password to connect to the directory service.
	*/
	Password *string `json:"password"`
	/*
	  Username to connect to the directory service.
	*/
	Username *string `json:"username"`
}

func (p *DsServiceAccount) MarshalJSON() ([]byte, error) {
	type DsServiceAccountProxy DsServiceAccount

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DsServiceAccountProxy
		Password *string `json:"password,omitempty"`
		Username *string `json:"username,omitempty"`
	}{
		DsServiceAccountProxy: (*DsServiceAccountProxy)(p),
		Password:              p.Password,
		Username:              p.Username,
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

func (p *DsServiceAccount) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DsServiceAccount
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDsServiceAccount()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDsServiceAccount() *DsServiceAccount {
	p := new(DsServiceAccount)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DsServiceAccount"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Description of the OIDC provider.
*/
type Federation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClaimMap *FederationClaims `json:"claimMap,omitempty"`
	/*
	  API key for authenticating to the OIDC provider.
	*/
	CloudApiKey *string `json:"cloudApiKey,omitempty"`
	/*
	  Provider side tenant that is registering domain.
	*/
	CloudTenant *string `json:"cloudTenant,omitempty"`
	/*
	  User or service who created the OIDC provider.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the OIDC provider.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  IDP attached to the users of the OIDC provider.
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Last updated time of the OIDC provider.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Well known URL of the OIDC provider.
	*/
	WellKnownConfig *string `json:"wellKnownConfig,omitempty"`
}

func (p *Federation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Federation

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

func (p *Federation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Federation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFederation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClaimMap != nil {
		p.ClaimMap = known.ClaimMap
	}
	if known.CloudApiKey != nil {
		p.CloudApiKey = known.CloudApiKey
	}
	if known.CloudTenant != nil {
		p.CloudTenant = known.CloudTenant
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IdpId != nil {
		p.IdpId = known.IdpId
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
	if known.WellKnownConfig != nil {
		p.WellKnownConfig = known.WellKnownConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "claimMap")
	delete(allFields, "cloudApiKey")
	delete(allFields, "cloudTenant")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "idpId")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "wellKnownConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFederation() *Federation {
	p := new(Federation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Federation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping of claims from IAM to the OIDC provider.
*/
type FederationClaims struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Email of the federation claim.
	*/
	Email *string `json:"email,omitempty"`
	/*
	  Name of the federation claim.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *FederationClaims) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FederationClaims

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

func (p *FederationClaims) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FederationClaims
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFederationClaims()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Email != nil {
		p.Email = known.Email
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "email")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFederationClaims() *FederationClaims {
	p := new(FederationClaims)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.FederationClaims"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/cert-auth-providers/{extId} Get operation
*/
type GetCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetCertAuthProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetCertAuthProviderApiResponse

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

func (p *GetCertAuthProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetCertAuthProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetCertAuthProviderApiResponse()

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

func NewGetCertAuthProviderApiResponse() *GetCertAuthProviderApiResponse {
	p := new(GetCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCertAuthProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCertAuthProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCertAuthProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services/{extId} Get operation
*/
type GetDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDirectoryServiceApiResponse

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

func (p *GetDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetDirectoryServiceApiResponse()

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

func NewGetDirectoryServiceApiResponse() *GetDirectoryServiceApiResponse {
	p := new(GetDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers/{extId} Get operation
*/
type GetSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetSamlIdentityProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetSamlIdentityProviderApiResponse

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

func (p *GetSamlIdentityProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetSamlIdentityProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetSamlIdentityProviderApiResponse()

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

func NewGetSamlIdentityProviderApiResponse() *GetSamlIdentityProviderApiResponse {
	p := new(GetSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSamlIdentityProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSamlIdentityProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSamlIdentityProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers/{extId}/sp-metadata Get operation
*/
type GetSamlIdpSpMetadataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSamlIdpSpMetadataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetSamlIdpSpMetadataApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetSamlIdpSpMetadataApiResponse

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

func (p *GetSamlIdpSpMetadataApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetSamlIdpSpMetadataApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetSamlIdpSpMetadataApiResponse()

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

func NewGetSamlIdpSpMetadataApiResponse() *GetSamlIdpSpMetadataApiResponse {
	p := new(GetSamlIdpSpMetadataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlIdpSpMetadataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSamlIdpSpMetadataApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSamlIdpSpMetadataApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSamlIdpSpMetadataApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-sp-metadata Get operation
*/
type GetSamlSpMetadataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSamlSpMetadataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetSamlSpMetadataApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetSamlSpMetadataApiResponse

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

func (p *GetSamlSpMetadataApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetSamlSpMetadataApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetSamlSpMetadataApiResponse()

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

func NewGetSamlSpMetadataApiResponse() *GetSamlSpMetadataApiResponse {
	p := new(GetSamlSpMetadataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlSpMetadataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSamlSpMetadataApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSamlSpMetadataApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSamlSpMetadataApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{extId} Get operation
*/
type GetUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetUserApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetUserApiResponse

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

func (p *GetUserApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetUserApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetUserApiResponse()

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

func NewGetUserApiResponse() *GetUserApiResponse {
	p := new(GetUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUserApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/user-groups/{extId} Get operation
*/
type GetUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetUserGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetUserGroupApiResponse

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

func (p *GetUserGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetUserGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetUserGroupApiResponse()

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

func NewGetUserGroupApiResponse() *GetUserGroupApiResponse {
	p := new(GetUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUserGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUserGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUserGroupApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{userExtId}/keys/{extId} Get operation
*/
type GetUserKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetUserKeyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetUserKeyApiResponse

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

func (p *GetUserKeyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetUserKeyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetUserKeyApiResponse()

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

func NewGetUserKeyApiResponse() *GetUserKeyApiResponse {
	p := new(GetUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUserKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUserKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUserKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/welcome-banner Get operation
*/
type GetWelcomeBannerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetWelcomeBannerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetWelcomeBannerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetWelcomeBannerApiResponse

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

func (p *GetWelcomeBannerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetWelcomeBannerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetWelcomeBannerApiResponse()

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

func NewGetWelcomeBannerApiResponse() *GetWelcomeBannerApiResponse {
	p := new(GetWelcomeBannerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetWelcomeBannerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetWelcomeBannerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetWelcomeBannerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetWelcomeBannerApiResponseData()
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
Type of grant.
*/
type GrantType int

const (
	GRANTTYPE_UNKNOWN                                         GrantType = 0
	GRANTTYPE_REDACTED                                        GrantType = 1
	GRANTTYPE_AUTHORIZATION_CODE                              GrantType = 2
	GRANTTYPE_REFRESH_TOKEN                                   GrantType = 3
	GRANTTYPE_CLIENT_CREDENTIALS                              GrantType = 4
	GRANTTYPE_PASSWORD                                        GrantType = 5
	GRANTTYPE_URN_IETF_PARAMS_OAUTH_GRANT_TYPE_TOKEN_EXCHANGE GrantType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GrantType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AUTHORIZATION_CODE",
		"REFRESH_TOKEN",
		"CLIENT_CREDENTIALS",
		"PASSWORD",
		"URN:IETF:PARAMS:OAUTH:GRANT-TYPE:TOKEN-EXCHANGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GrantType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AUTHORIZATION_CODE",
		"REFRESH_TOKEN",
		"CLIENT_CREDENTIALS",
		"PASSWORD",
		"URN:IETF:PARAMS:OAUTH:GRANT-TYPE:TOKEN-EXCHANGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GrantType) index(name string) GrantType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AUTHORIZATION_CODE",
		"REFRESH_TOKEN",
		"CLIENT_CREDENTIALS",
		"PASSWORD",
		"URN:IETF:PARAMS:OAUTH:GRANT-TYPE:TOKEN-EXCHANGE",
	}
	for idx := range names {
		if names[idx] == name {
			return GrantType(idx)
		}
	}
	return GRANTTYPE_UNKNOWN
}

func (e *GrantType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GrantType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GrantType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GrantType) Ref() *GrantType {
	return &e
}

/*
Information on a group of users.
*/
type Group struct {
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
	  The name of group.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Group) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Group

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

func (p *Group) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Group
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroup()

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
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGroup() *Group {
	p := new(Group)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Group"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Group membership search type for the directory service.
*/
type GroupSearchType int

const (
	GROUPSEARCHTYPE_UNKNOWN       GroupSearchType = 0
	GROUPSEARCHTYPE_REDACTED      GroupSearchType = 1
	GROUPSEARCHTYPE_NON_RECURSIVE GroupSearchType = 2
	GROUPSEARCHTYPE_RECURSIVE     GroupSearchType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GroupSearchType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NON_RECURSIVE",
		"RECURSIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GroupSearchType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NON_RECURSIVE",
		"RECURSIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GroupSearchType) index(name string) GroupSearchType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NON_RECURSIVE",
		"RECURSIVE",
	}
	for idx := range names {
		if names[idx] == name {
			return GroupSearchType(idx)
		}
	}
	return GROUPSEARCHTYPE_UNKNOWN
}

func (e *GroupSearchType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GroupSearchType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GroupSearchType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GroupSearchType) Ref() *GroupSearchType {
	return &e
}

/*
Type of the user group.
*/
type GroupType int

const (
	GROUPTYPE_UNKNOWN  GroupType = 0
	GROUPTYPE_REDACTED GroupType = 1
	GROUPTYPE_SAML     GroupType = 2
	GROUPTYPE_LDAP     GroupType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GroupType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAML",
		"LDAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GroupType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAML",
		"LDAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GroupType) index(name string) GroupType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SAML",
		"LDAP",
	}
	for idx := range names {
		if names[idx] == name {
			return GroupType(idx)
		}
	}
	return GROUPTYPE_UNKNOWN
}

func (e *GroupType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GroupType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GroupType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GroupType) Ref() *GroupType {
	return &e
}

/*
Information of the IDP.
*/
type IdpMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Certificate for verification.
	*/
	Certificate *string `json:"certificate"`
	/*
	  Entity identifier of identity provider.
	*/
	EntityId *string `json:"entityId"`
	/*
	  Error URL of the identity provider.
	*/
	ErrorUrl *string `json:"errorUrl,omitempty"`
	/*
	  Login URL of the identity provider.
	*/
	LoginUrl *string `json:"loginUrl"`
	/*
	  Logout URL of the identity provider.
	*/
	LogoutUrl *string `json:"logoutUrl,omitempty"`

	NameIdPolicyFormat *NameIdPolicyFormat `json:"nameIdPolicyFormat,omitempty"`
}

func (p *IdpMetadata) MarshalJSON() ([]byte, error) {
	type IdpMetadataProxy IdpMetadata

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IdpMetadataProxy
		Certificate *string `json:"certificate,omitempty"`
		EntityId    *string `json:"entityId,omitempty"`
		LoginUrl    *string `json:"loginUrl,omitempty"`
	}{
		IdpMetadataProxy: (*IdpMetadataProxy)(p),
		Certificate:      p.Certificate,
		EntityId:         p.EntityId,
		LoginUrl:         p.LoginUrl,
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

func (p *IdpMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IdpMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIdpMetadata()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Certificate != nil {
		p.Certificate = known.Certificate
	}
	if known.EntityId != nil {
		p.EntityId = known.EntityId
	}
	if known.ErrorUrl != nil {
		p.ErrorUrl = known.ErrorUrl
	}
	if known.LoginUrl != nil {
		p.LoginUrl = known.LoginUrl
	}
	if known.LogoutUrl != nil {
		p.LogoutUrl = known.LogoutUrl
	}
	if known.NameIdPolicyFormat != nil {
		p.NameIdPolicyFormat = known.NameIdPolicyFormat
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "certificate")
	delete(allFields, "entityId")
	delete(allFields, "errorUrl")
	delete(allFields, "loginUrl")
	delete(allFields, "logoutUrl")
	delete(allFields, "nameIdPolicyFormat")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIdpMetadata() *IdpMetadata {
	p := new(IdpMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.IdpMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the login provider configured.
*/
type IdpType int

const (
	IDPTYPE_UNKNOWN         IdpType = 0
	IDPTYPE_REDACTED        IdpType = 1
	IDPTYPE_LOCAL           IdpType = 2
	IDPTYPE_SAML            IdpType = 3
	IDPTYPE_LDAP            IdpType = 4
	IDPTYPE_CERT            IdpType = 5
	IDPTYPE_SERVICE_ACCOUNT IdpType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IdpType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"local",
		"saml",
		"ldap",
		"cert",
		"service_account",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e IdpType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"local",
		"saml",
		"ldap",
		"cert",
		"service_account",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *IdpType) index(name string) IdpType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"local",
		"saml",
		"ldap",
		"cert",
		"service_account",
	}
	for idx := range names {
		if names[idx] == name {
			return IdpType(idx)
		}
	}
	return IDPTYPE_UNKNOWN
}

func (e *IdpType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IdpType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IdpType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IdpType) Ref() *IdpType {
	return &e
}

/*
Credentials in the form of a unique random key for the users.
*/
type Key struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External client to whom the key is allocated.
	*/
	AssignedTo *string `json:"assignedTo,omitempty"`
	/*
	  User or service who created the key.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the key.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	CreationType *CreationType `json:"creationType,omitempty"`
	/*
	  Brief description of the key.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The time when the key will expire.
	*/
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*

	 */
	KeyDetailsItemDiscriminator_ *string `json:"$keyDetailsItemDiscriminator,omitempty"`
	/*
	  Details specific to type of the key.
	*/
	KeyDetails *OneOfKeyKeyDetails `json:"keyDetails,omitempty"`

	KeyType *KeyKind `json:"keyType"`
	/*
	  User who updated the key.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  The time when the key was updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  The time when the key was last used.
	*/
	LastUsedTime *time.Time `json:"lastUsedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier for the key in the form of a name.
	*/
	Name *string `json:"name"`

	Status *KeyStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Key) MarshalJSON() ([]byte, error) {
	type KeyProxy Key

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*KeyProxy
		KeyType *KeyKind `json:"keyType,omitempty"`
		Name    *string  `json:"name,omitempty"`
	}{
		KeyProxy: (*KeyProxy)(p),
		KeyType:  p.KeyType,
		Name:     p.Name,
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

func (p *Key) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Key
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKey()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssignedTo != nil {
		p.AssignedTo = known.AssignedTo
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CreationType != nil {
		p.CreationType = known.CreationType
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExpiryTime != nil {
		p.ExpiryTime = known.ExpiryTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.KeyDetailsItemDiscriminator_ != nil {
		p.KeyDetailsItemDiscriminator_ = known.KeyDetailsItemDiscriminator_
	}
	if known.KeyDetails != nil {
		p.KeyDetails = known.KeyDetails
	}
	if known.KeyType != nil {
		p.KeyType = known.KeyType
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LastUsedTime != nil {
		p.LastUsedTime = known.LastUsedTime
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
	delete(allFields, "assignedTo")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "creationType")
	delete(allFields, "description")
	delete(allFields, "expiryTime")
	delete(allFields, "extId")
	delete(allFields, "$keyDetailsItemDiscriminator")
	delete(allFields, "keyDetails")
	delete(allFields, "keyType")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "lastUsedTime")
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

func NewKey() *Key {
	p := new(Key)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Key"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Key) GetKeyDetails() interface{} {
	if nil == p.KeyDetails {
		return nil
	}
	return p.KeyDetails.GetValue()
}

func (p *Key) SetKeyDetails(v interface{}) error {
	if nil == p.KeyDetails {
		p.KeyDetails = NewOneOfKeyKeyDetails()
	}
	e := p.KeyDetails.SetValue(v)
	if nil == e {
		if nil == p.KeyDetailsItemDiscriminator_ {
			p.KeyDetailsItemDiscriminator_ = new(string)
		}
		*p.KeyDetailsItemDiscriminator_ = *p.KeyDetails.Discriminator
	}
	return e
}

/*
The type of key.
*/
type KeyKind int

const (
	KEYKIND_UNKNOWN    KeyKind = 0
	KEYKIND_REDACTED   KeyKind = 1
	KEYKIND_API_KEY    KeyKind = 2
	KEYKIND_OBJECT_KEY KeyKind = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *KeyKind) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"API_KEY",
		"OBJECT_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e KeyKind) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"API_KEY",
		"OBJECT_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *KeyKind) index(name string) KeyKind {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"API_KEY",
		"OBJECT_KEY",
	}
	for idx := range names {
		if names[idx] == name {
			return KeyKind(idx)
		}
	}
	return KEYKIND_UNKNOWN
}

func (e *KeyKind) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for KeyKind:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *KeyKind) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e KeyKind) Ref() *KeyKind {
	return &e
}

/*
Migrate keys request.
*/
type KeyMigrationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of keys to be migrated in this API call.
	*/
	MaxKeyCount *int `json:"maxKeyCount"`
}

func (p *KeyMigrationSpec) MarshalJSON() ([]byte, error) {
	type KeyMigrationSpecProxy KeyMigrationSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*KeyMigrationSpecProxy
		MaxKeyCount *int `json:"maxKeyCount,omitempty"`
	}{
		KeyMigrationSpecProxy: (*KeyMigrationSpecProxy)(p),
		MaxKeyCount:           p.MaxKeyCount,
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

func (p *KeyMigrationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KeyMigrationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKeyMigrationSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MaxKeyCount != nil {
		p.MaxKeyCount = known.MaxKeyCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "maxKeyCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewKeyMigrationSpec() *KeyMigrationSpec {
	p := new(KeyMigrationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.KeyMigrationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The status of the key.
*/
type KeyStatus int

const (
	KEYSTATUS_UNKNOWN  KeyStatus = 0
	KEYSTATUS_REDACTED KeyStatus = 1
	KEYSTATUS_VALID    KeyStatus = 2
	KEYSTATUS_REVOKED  KeyStatus = 3
	KEYSTATUS_EXPIRED  KeyStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *KeyStatus) name(index int) string {
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
func (e KeyStatus) GetName() string {
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
func (e *KeyStatus) index(name string) KeyStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"REVOKED",
		"EXPIRED",
	}
	for idx := range names {
		if names[idx] == name {
			return KeyStatus(idx)
		}
	}
	return KEYSTATUS_UNKNOWN
}

func (e *KeyStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for KeyStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *KeyStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e KeyStatus) Ref() *KeyStatus {
	return &e
}

/*
Type of OIDC key.
*/
type KeyType int

const (
	KEYTYPE_UNKNOWN  KeyType = 0
	KEYTYPE_REDACTED KeyType = 1
	KEYTYPE_RSA      KeyType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *KeyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e KeyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *KeyType) index(name string) KeyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RSA",
	}
	for idx := range names {
		if names[idx] == name {
			return KeyType(idx)
		}
	}
	return KEYTYPE_UNKNOWN
}

func (e *KeyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for KeyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *KeyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e KeyType) Ref() *KeyType {
	return &e
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/cert-auth-providers Get operation
*/
type ListCertAuthProvidersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCertAuthProvidersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListCertAuthProvidersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListCertAuthProvidersApiResponse

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

func (p *ListCertAuthProvidersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListCertAuthProvidersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListCertAuthProvidersApiResponse()

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

func NewListCertAuthProvidersApiResponse() *ListCertAuthProvidersApiResponse {
	p := new(ListCertAuthProvidersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListCertAuthProvidersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCertAuthProvidersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCertAuthProvidersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCertAuthProvidersApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services Get operation
*/
type ListDirectoryServicesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDirectoryServicesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListDirectoryServicesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListDirectoryServicesApiResponse

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

func (p *ListDirectoryServicesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListDirectoryServicesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListDirectoryServicesApiResponse()

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

func NewListDirectoryServicesApiResponse() *ListDirectoryServicesApiResponse {
	p := new(ListDirectoryServicesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListDirectoryServicesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDirectoryServicesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDirectoryServicesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDirectoryServicesApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/login-providers Get operation
*/
type ListLoginProvidersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLoginProvidersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListLoginProvidersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListLoginProvidersApiResponse

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

func (p *ListLoginProvidersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListLoginProvidersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListLoginProvidersApiResponse()

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

func NewListLoginProvidersApiResponse() *ListLoginProvidersApiResponse {
	p := new(ListLoginProvidersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListLoginProvidersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLoginProvidersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLoginProvidersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLoginProvidersApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers Get operation
*/
type ListSamlIdentityProvidersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSamlIdentityProvidersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListSamlIdentityProvidersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListSamlIdentityProvidersApiResponse

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

func (p *ListSamlIdentityProvidersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListSamlIdentityProvidersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListSamlIdentityProvidersApiResponse()

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

func NewListSamlIdentityProvidersApiResponse() *ListSamlIdentityProvidersApiResponse {
	p := new(ListSamlIdentityProvidersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListSamlIdentityProvidersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSamlIdentityProvidersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSamlIdentityProvidersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSamlIdentityProvidersApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/user-groups Get operation
*/
type ListUserGroupsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserGroupsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListUserGroupsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListUserGroupsApiResponse

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

func (p *ListUserGroupsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListUserGroupsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListUserGroupsApiResponse()

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

func NewListUserGroupsApiResponse() *ListUserGroupsApiResponse {
	p := new(ListUserGroupsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserGroupsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUserGroupsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUserGroupsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUserGroupsApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{userExtId}/keys Get operation
*/
type ListUserKeysApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserKeysApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListUserKeysApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListUserKeysApiResponse

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

func (p *ListUserKeysApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListUserKeysApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListUserKeysApiResponse()

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

func NewListUserKeysApiResponse() *ListUserKeysApiResponse {
	p := new(ListUserKeysApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserKeysApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUserKeysApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUserKeysApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUserKeysApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users Get operation
*/
type ListUsersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUsersApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListUsersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListUsersApiResponse

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

func (p *ListUsersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListUsersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListUsersApiResponse()

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

func NewListUsersApiResponse() *ListUsersApiResponse {
	p := new(ListUsersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUsersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUsersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUsersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUsersApiResponseData()
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
Description of the login provider configured.
*/
type LoginProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Creation time of the login provider.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag for filtering the connectors for the UI to display.
	*/
	IsBrowserLoginSupported *bool `json:"isBrowserLoginSupported,omitempty"`
	/*
	  Last update time of the login provider.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the login provider configured.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *IdpType `json:"type,omitempty"`
}

func (p *LoginProvider) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LoginProvider

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

func (p *LoginProvider) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoginProvider
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLoginProvider()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsBrowserLoginSupported != nil {
		p.IsBrowserLoginSupported = known.IsBrowserLoginSupported
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
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "isBrowserLoginSupported")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoginProvider() *LoginProvider {
	p := new(LoginProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.LoginProvider"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Request body for creating login request.
*/
type LoginRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Approval prompt value for the login request.
	*/
	ApprovalPrompt *string `json:"approvalPrompt,omitempty"`
	/*
	  Audience value for the login request.
	*/
	Audience *string `json:"audience,omitempty"`
	/*
	  Client identifier for the the login request.
	*/
	ClientId *string `json:"clientId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag to determine if service provider is enabled or not.
	*/
	IsServiceProviderEnabled *bool `json:"isServiceProviderEnabled,omitempty"`
	/*
	  Flag to determine if the user is logged in or not.
	*/
	IsUserLoggedIn *bool `json:"isUserLoggedIn,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier for the login provider to be updated in the login request.
	*/
	LoginProviderId *string `json:"loginProviderId,omitempty"`
	/*
	  Nonce value for the login request.
	*/
	Nonce *string `json:"nonce,omitempty"`
	/*
	  Redirect URI for the login request.
	*/
	RedirectUri *string `json:"redirectUri,omitempty"`

	ResponseType *ResponseType `json:"responseType,omitempty"`
	/*
	  Supported scope values for the login request. The supported values are OPENID, PROFILE, EMAIL, OFFLINE_ACCESS and GROUPS. OPENID is a mandatory field.
	*/
	Scope []Scope `json:"scope,omitempty"`
	/*
	  State value for the login request.
	*/
	State *string `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *LoginRequest) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LoginRequest

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

func (p *LoginRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoginRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLoginRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApprovalPrompt != nil {
		p.ApprovalPrompt = known.ApprovalPrompt
	}
	if known.Audience != nil {
		p.Audience = known.Audience
	}
	if known.ClientId != nil {
		p.ClientId = known.ClientId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsServiceProviderEnabled != nil {
		p.IsServiceProviderEnabled = known.IsServiceProviderEnabled
	}
	if known.IsUserLoggedIn != nil {
		p.IsUserLoggedIn = known.IsUserLoggedIn
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LoginProviderId != nil {
		p.LoginProviderId = known.LoginProviderId
	}
	if known.Nonce != nil {
		p.Nonce = known.Nonce
	}
	if known.RedirectUri != nil {
		p.RedirectUri = known.RedirectUri
	}
	if known.ResponseType != nil {
		p.ResponseType = known.ResponseType
	}
	if known.Scope != nil {
		p.Scope = known.Scope
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "approvalPrompt")
	delete(allFields, "audience")
	delete(allFields, "clientId")
	delete(allFields, "extId")
	delete(allFields, "isServiceProviderEnabled")
	delete(allFields, "isUserLoggedIn")
	delete(allFields, "links")
	delete(allFields, "loginProviderId")
	delete(allFields, "nonce")
	delete(allFields, "redirectUri")
	delete(allFields, "responseType")
	delete(allFields, "scope")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoginRequest() *LoginRequest {
	p := new(LoginRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.LoginRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Name ID Policy format.
*/
type NameIdPolicyFormat int

const (
	NAMEIDPOLICYFORMAT_UNKNOWN                    NameIdPolicyFormat = 0
	NAMEIDPOLICYFORMAT_REDACTED                   NameIdPolicyFormat = 1
	NAMEIDPOLICYFORMAT_EMAILADDRESS               NameIdPolicyFormat = 2
	NAMEIDPOLICYFORMAT_UNSPECIFIED                NameIdPolicyFormat = 3
	NAMEIDPOLICYFORMAT_X509SUBJECTNAME            NameIdPolicyFormat = 4
	NAMEIDPOLICYFORMAT_WINDOWSDOMAINQUALIFIEDNAME NameIdPolicyFormat = 5
	NAMEIDPOLICYFORMAT_ENCRYPTED                  NameIdPolicyFormat = 6
	NAMEIDPOLICYFORMAT_ENTITY                     NameIdPolicyFormat = 7
	NAMEIDPOLICYFORMAT_KERBEROS                   NameIdPolicyFormat = 8
	NAMEIDPOLICYFORMAT_PERSISTENT                 NameIdPolicyFormat = 9
	NAMEIDPOLICYFORMAT_TRANSIENT                  NameIdPolicyFormat = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NameIdPolicyFormat) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"emailAddress",
		"unspecified",
		"X509SubjectName",
		"WindowsDomainQualifiedName",
		"encrypted",
		"entity",
		"kerberos",
		"persistent",
		"transient",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NameIdPolicyFormat) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"emailAddress",
		"unspecified",
		"X509SubjectName",
		"WindowsDomainQualifiedName",
		"encrypted",
		"entity",
		"kerberos",
		"persistent",
		"transient",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NameIdPolicyFormat) index(name string) NameIdPolicyFormat {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"emailAddress",
		"unspecified",
		"X509SubjectName",
		"WindowsDomainQualifiedName",
		"encrypted",
		"entity",
		"kerberos",
		"persistent",
		"transient",
	}
	for idx := range names {
		if names[idx] == name {
			return NameIdPolicyFormat(idx)
		}
	}
	return NAMEIDPOLICYFORMAT_UNKNOWN
}

func (e *NameIdPolicyFormat) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NameIdPolicyFormat:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NameIdPolicyFormat) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NameIdPolicyFormat) Ref() *NameIdPolicyFormat {
	return &e
}

/*
The object key details.
*/
type ObjectKeyDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Access key identifier.
	*/
	AccessKey *string `json:"accessKey,omitempty"`
	/*
	  Secret Access Key, which will be returned during creation of the key of type OBJECT_KEY.
	*/
	SecretKey *string `json:"secretKey,omitempty"`
}

func (p *ObjectKeyDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ObjectKeyDetails

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

func (p *ObjectKeyDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectKeyDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectKeyDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessKey != nil {
		p.AccessKey = known.AccessKey
	}
	if known.SecretKey != nil {
		p.SecretKey = known.SecretKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessKey")
	delete(allFields, "secretKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectKeyDetails() *ObjectKeyDetails {
	p := new(ObjectKeyDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ObjectKeyDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The retrieved auth details.
*/
type ObjectsAuthDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External client to whom this key is allocated.
	*/
	AssignedTo *string `json:"assignedTo"`

	CreationType *CreationType `json:"creationType"`
	/*
	  Any additional user attributes obtained from the authentication provider in the form of key/value pairs.
	*/
	CustomClaims map[string]string `json:"customClaims"`
	/*
	  Email ID of the user.
	*/
	EmailId *string `json:"emailId"`
	/*
	  Identifier of the IDP for the user.
	*/
	IdpId *string `json:"idpId"`
	/*
	  This represents secret access key, which will be returned only during access key creation.
	*/
	SecretAccessKey *string `json:"secretAccessKey"`
	/*
	  ID of tenant for which configuration is being set up.
	*/
	TenantId *string `json:"tenantId"`
	/*
	  User identifier who owns the bucket access key.
	*/
	UserId *string `json:"userId"`

	UserType *UserType `json:"userType"`
	/*
	  Identifier of the user.
	*/
	Username *string `json:"username"`
}

func (p *ObjectsAuthDetails) MarshalJSON() ([]byte, error) {
	type ObjectsAuthDetailsProxy ObjectsAuthDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ObjectsAuthDetailsProxy
		AssignedTo      *string           `json:"assignedTo,omitempty"`
		CreationType    *CreationType     `json:"creationType,omitempty"`
		CustomClaims    map[string]string `json:"customClaims,omitempty"`
		EmailId         *string           `json:"emailId,omitempty"`
		IdpId           *string           `json:"idpId,omitempty"`
		SecretAccessKey *string           `json:"secretAccessKey,omitempty"`
		TenantId        *string           `json:"tenantId,omitempty"`
		UserId          *string           `json:"userId,omitempty"`
		UserType        *UserType         `json:"userType,omitempty"`
		Username        *string           `json:"username,omitempty"`
	}{
		ObjectsAuthDetailsProxy: (*ObjectsAuthDetailsProxy)(p),
		AssignedTo:              p.AssignedTo,
		CreationType:            p.CreationType,
		CustomClaims:            p.CustomClaims,
		EmailId:                 p.EmailId,
		IdpId:                   p.IdpId,
		SecretAccessKey:         p.SecretAccessKey,
		TenantId:                p.TenantId,
		UserId:                  p.UserId,
		UserType:                p.UserType,
		Username:                p.Username,
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

func (p *ObjectsAuthDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectsAuthDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectsAuthDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AssignedTo != nil {
		p.AssignedTo = known.AssignedTo
	}
	if known.CreationType != nil {
		p.CreationType = known.CreationType
	}
	if known.CustomClaims != nil {
		p.CustomClaims = known.CustomClaims
	}
	if known.EmailId != nil {
		p.EmailId = known.EmailId
	}
	if known.IdpId != nil {
		p.IdpId = known.IdpId
	}
	if known.SecretAccessKey != nil {
		p.SecretAccessKey = known.SecretAccessKey
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UserId != nil {
		p.UserId = known.UserId
	}
	if known.UserType != nil {
		p.UserType = known.UserType
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "assignedTo")
	delete(allFields, "creationType")
	delete(allFields, "customClaims")
	delete(allFields, "emailId")
	delete(allFields, "idpId")
	delete(allFields, "secretAccessKey")
	delete(allFields, "tenantId")
	delete(allFields, "userId")
	delete(allFields, "userType")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectsAuthDetails() *ObjectsAuthDetails {
	p := new(ObjectsAuthDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ObjectsAuthDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The specification needed to fetch the auth details.
*/
type ObjectsAuthSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The access key which is used to retrieve the auth details.
	*/
	AccessKey *string `json:"accessKey"`
}

func (p *ObjectsAuthSpec) MarshalJSON() ([]byte, error) {
	type ObjectsAuthSpecProxy ObjectsAuthSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ObjectsAuthSpecProxy
		AccessKey *string `json:"accessKey,omitempty"`
	}{
		ObjectsAuthSpecProxy: (*ObjectsAuthSpecProxy)(p),
		AccessKey:            p.AccessKey,
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

func (p *ObjectsAuthSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectsAuthSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectsAuthSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessKey != nil {
		p.AccessKey = known.AccessKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectsAuthSpec() *ObjectsAuthSpec {
	p := new(ObjectsAuthSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ObjectsAuthSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
JSON Web Key that can be used to verify the signature of tokens issued by IAM.
*/
type OidcKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Alg *AlgoType `json:"alg"`
	/*
	  Exponent for a standard PEM.
	*/
	E *string `json:"e"`
	/*
	  Unique identifier of the key.
	*/
	Kid *string `json:"kid"`

	Kty *KeyType `json:"kty"`
	/*
	  Modulus for a standard PEM.
	*/
	N *string `json:"n"`

	Use *UseType `json:"use"`
	/*
	  x509 certificate chain of the key.
	*/
	X5c *string `json:"x5c"`
}

func (p *OidcKey) MarshalJSON() ([]byte, error) {
	type OidcKeyProxy OidcKey

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*OidcKeyProxy
		Alg *AlgoType `json:"alg,omitempty"`
		E   *string   `json:"e,omitempty"`
		Kid *string   `json:"kid,omitempty"`
		Kty *KeyType  `json:"kty,omitempty"`
		N   *string   `json:"n,omitempty"`
		Use *UseType  `json:"use,omitempty"`
		X5c *string   `json:"x5c,omitempty"`
	}{
		OidcKeyProxy: (*OidcKeyProxy)(p),
		Alg:          p.Alg,
		E:            p.E,
		Kid:          p.Kid,
		Kty:          p.Kty,
		N:            p.N,
		Use:          p.Use,
		X5c:          p.X5c,
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

func (p *OidcKey) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OidcKey
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOidcKey()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Alg != nil {
		p.Alg = known.Alg
	}
	if known.E != nil {
		p.E = known.E
	}
	if known.Kid != nil {
		p.Kid = known.Kid
	}
	if known.Kty != nil {
		p.Kty = known.Kty
	}
	if known.N != nil {
		p.N = known.N
	}
	if known.Use != nil {
		p.Use = known.Use
	}
	if known.X5c != nil {
		p.X5c = known.X5c
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "alg")
	delete(allFields, "e")
	delete(allFields, "kid")
	delete(allFields, "kty")
	delete(allFields, "n")
	delete(allFields, "use")
	delete(allFields, "x5c")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOidcKey() *OidcKey {
	p := new(OidcKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the subject
*/
type OidcSubjectType int

const (
	OIDCSUBJECTTYPE_UNKNOWN         OidcSubjectType = 0
	OIDCSUBJECTTYPE_REDACTED        OidcSubjectType = 1
	OIDCSUBJECTTYPE_LOCAL           OidcSubjectType = 2
	OIDCSUBJECTTYPE_LDAP            OidcSubjectType = 3
	OIDCSUBJECTTYPE_SAML            OidcSubjectType = 4
	OIDCSUBJECTTYPE_CERT            OidcSubjectType = 5
	OIDCSUBJECTTYPE_SERVICE_ACCOUNT OidcSubjectType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OidcSubjectType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"LDAP",
		"SAML",
		"CERT",
		"SERVICE_ACCOUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OidcSubjectType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"LDAP",
		"SAML",
		"CERT",
		"SERVICE_ACCOUNT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OidcSubjectType) index(name string) OidcSubjectType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"LDAP",
		"SAML",
		"CERT",
		"SERVICE_ACCOUNT",
	}
	for idx := range names {
		if names[idx] == name {
			return OidcSubjectType(idx)
		}
	}
	return OIDCSUBJECTTYPE_UNKNOWN
}

func (e *OidcSubjectType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OidcSubjectType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OidcSubjectType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OidcSubjectType) Ref() *OidcSubjectType {
	return &e
}

/*
IAM OpenID userinfo response.
*/
type OidcUserinfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID V5 of the authentication connector that was used to verify the credentials of the user.
	*/
	ConnectorId *string `json:"connectorId,omitempty"`
	/*
	  Any additional user attributes obtained from the authentication provider in the form of key/value pairs.
	*/
	CustomClaims map[string]string `json:"customClaims,omitempty"`
	/*
	  Email address of the subject.
	*/
	Email *string `json:"email,omitempty"`
	/*
	  Expiration time for the identity claims in seconds since epoch.
	*/
	Exp *int64 `json:"exp,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  First name of the user.
	*/
	FirstName *string `json:"firstName,omitempty"`
	/*
	  List of groups obtained from the authentication provider for the user.
	*/
	Groups []Group `json:"groups,omitempty"`
	/*
	  Time at which the identity claims were issued in seconds since epoch.
	*/
	Iat *int64 `json:"iat,omitempty"`
	/*
	  Boolean to indicate if the email address of end user has been verified by IAM.
	*/
	IsEmailVerified *bool `json:"isEmailVerified,omitempty"`
	/*
	  Base URL of the identity service which it asserts as its issuer identifier.
	*/
	Iss *string `json:"iss,omitempty"`
	/*
	  Last name of the user.
	*/
	LastName *string `json:"lastName,omitempty"`
	/*
	  List of legacy Prism roles for the user.
	*/
	LegacyRoles []string `json:"legacyRoles,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Default locale of the user.
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  Middle name of the user.
	*/
	MiddleInitial *string `json:"middleInitial,omitempty"`
	/*
	  Display name of the user including the first and last name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Default region of the user.
	*/
	Region *string `json:"region,omitempty"`

	Status *UserStatusType `json:"status,omitempty"`
	/*
	  Subject identifier. A locally unique identifier for the end user.
	*/
	Sub *string `json:"sub,omitempty"`

	SubType *OidcSubjectType `json:"subType,omitempty"`

	Tenant *Tenant `json:"tenant,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  UUID V5 created for the user by IAM.
	*/
	UserId *string `json:"userId,omitempty"`
}

func (p *OidcUserinfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OidcUserinfo

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

func (p *OidcUserinfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OidcUserinfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOidcUserinfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConnectorId != nil {
		p.ConnectorId = known.ConnectorId
	}
	if known.CustomClaims != nil {
		p.CustomClaims = known.CustomClaims
	}
	if known.Email != nil {
		p.Email = known.Email
	}
	if known.Exp != nil {
		p.Exp = known.Exp
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FirstName != nil {
		p.FirstName = known.FirstName
	}
	if known.Groups != nil {
		p.Groups = known.Groups
	}
	if known.Iat != nil {
		p.Iat = known.Iat
	}
	if known.IsEmailVerified != nil {
		p.IsEmailVerified = known.IsEmailVerified
	}
	if known.Iss != nil {
		p.Iss = known.Iss
	}
	if known.LastName != nil {
		p.LastName = known.LastName
	}
	if known.LegacyRoles != nil {
		p.LegacyRoles = known.LegacyRoles
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Locale != nil {
		p.Locale = known.Locale
	}
	if known.MiddleInitial != nil {
		p.MiddleInitial = known.MiddleInitial
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Region != nil {
		p.Region = known.Region
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.Sub != nil {
		p.Sub = known.Sub
	}
	if known.SubType != nil {
		p.SubType = known.SubType
	}
	if known.Tenant != nil {
		p.Tenant = known.Tenant
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UserId != nil {
		p.UserId = known.UserId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "connectorId")
	delete(allFields, "customClaims")
	delete(allFields, "email")
	delete(allFields, "exp")
	delete(allFields, "extId")
	delete(allFields, "firstName")
	delete(allFields, "groups")
	delete(allFields, "iat")
	delete(allFields, "isEmailVerified")
	delete(allFields, "iss")
	delete(allFields, "lastName")
	delete(allFields, "legacyRoles")
	delete(allFields, "links")
	delete(allFields, "locale")
	delete(allFields, "middleInitial")
	delete(allFields, "name")
	delete(allFields, "region")
	delete(allFields, "status")
	delete(allFields, "sub")
	delete(allFields, "subType")
	delete(allFields, "tenant")
	delete(allFields, "tenantId")
	delete(allFields, "userId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOidcUserinfo() *OidcUserinfo {
	p := new(OidcUserinfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcUserinfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for OpenLDAP directory service.
*/
type OpenLdapConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	UserConfiguration *UserConfiguration `json:"userConfiguration"`

	UserGroupConfiguration *UserGroupConfiguration `json:"userGroupConfiguration"`
}

func (p *OpenLdapConfig) MarshalJSON() ([]byte, error) {
	type OpenLdapConfigProxy OpenLdapConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*OpenLdapConfigProxy
		UserConfiguration      *UserConfiguration      `json:"userConfiguration,omitempty"`
		UserGroupConfiguration *UserGroupConfiguration `json:"userGroupConfiguration,omitempty"`
	}{
		OpenLdapConfigProxy:    (*OpenLdapConfigProxy)(p),
		UserConfiguration:      p.UserConfiguration,
		UserGroupConfiguration: p.UserGroupConfiguration,
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

func (p *OpenLdapConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OpenLdapConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewOpenLdapConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.UserConfiguration != nil {
		p.UserConfiguration = known.UserConfiguration
	}
	if known.UserGroupConfiguration != nil {
		p.UserGroupConfiguration = known.UserGroupConfiguration
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "userConfiguration")
	delete(allFields, "userGroupConfiguration")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOpenLdapConfig() *OpenLdapConfig {
	p := new(OpenLdapConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OpenLdapConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for password change of user.
*/
type PasswordChangeRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  New password of the user.
	*/
	NewPassword *string `json:"newPassword"`
	/*
	  Current password of the user.
	*/
	OldPassword *string `json:"oldPassword"`
	/*
	  Identifier of the user.
	*/
	Username *string `json:"username"`
}

func (p *PasswordChangeRequest) MarshalJSON() ([]byte, error) {
	type PasswordChangeRequestProxy PasswordChangeRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PasswordChangeRequestProxy
		NewPassword *string `json:"newPassword,omitempty"`
		OldPassword *string `json:"oldPassword,omitempty"`
		Username    *string `json:"username,omitempty"`
	}{
		PasswordChangeRequestProxy: (*PasswordChangeRequestProxy)(p),
		NewPassword:                p.NewPassword,
		OldPassword:                p.OldPassword,
		Username:                   p.Username,
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

func (p *PasswordChangeRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PasswordChangeRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPasswordChangeRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NewPassword != nil {
		p.NewPassword = known.NewPassword
	}
	if known.OldPassword != nil {
		p.OldPassword = known.OldPassword
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "newPassword")
	delete(allFields, "oldPassword")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPasswordChangeRequest() *PasswordChangeRequest {
	p := new(PasswordChangeRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordChangeRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Change user password response object.
*/
type PasswordChangeResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *PasswordChangeResponse) MarshalJSON() ([]byte, error) {
	type PasswordChangeResponseProxy PasswordChangeResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PasswordChangeResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		PasswordChangeResponseProxy: (*PasswordChangeResponseProxy)(p),
		Message:                     p.Message,
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

func (p *PasswordChangeResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PasswordChangeResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPasswordChangeResponse()

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

func NewPasswordChangeResponse() *PasswordChangeResponse {
	p := new(PasswordChangeResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordChangeResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for password reset of user.
*/
type PasswordResetRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  New password of the user.
	*/
	NewPassword *string `json:"newPassword"`
}

func (p *PasswordResetRequest) MarshalJSON() ([]byte, error) {
	type PasswordResetRequestProxy PasswordResetRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PasswordResetRequestProxy
		NewPassword *string `json:"newPassword,omitempty"`
	}{
		PasswordResetRequestProxy: (*PasswordResetRequestProxy)(p),
		NewPassword:               p.NewPassword,
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

func (p *PasswordResetRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PasswordResetRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPasswordResetRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NewPassword != nil {
		p.NewPassword = known.NewPassword
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "newPassword")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPasswordResetRequest() *PasswordResetRequest {
	p := new(PasswordResetRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordResetRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User password reset response.
*/
type PasswordResetResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *PasswordResetResponse) MarshalJSON() ([]byte, error) {
	type PasswordResetResponseProxy PasswordResetResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PasswordResetResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		PasswordResetResponseProxy: (*PasswordResetResponseProxy)(p),
		Message:                    p.Message,
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

func (p *PasswordResetResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PasswordResetResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPasswordResetResponse()

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

func NewPasswordResetResponse() *PasswordResetResponse {
	p := new(PasswordResetResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordResetResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{extId}/$actions/reset-password Post operation
*/
type ResetUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResetUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ResetUserPasswordApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ResetUserPasswordApiResponse

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

func (p *ResetUserPasswordApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResetUserPasswordApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResetUserPasswordApiResponse()

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

func NewResetUserPasswordApiResponse() *ResetUserPasswordApiResponse {
	p := new(ResetUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ResetUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ResetUserPasswordApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ResetUserPasswordApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfResetUserPasswordApiResponseData()
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
Supported OIDC response types.
*/
type ResponseType int

const (
	RESPONSETYPE_UNKNOWN  ResponseType = 0
	RESPONSETYPE_REDACTED ResponseType = 1
	RESPONSETYPE_ID_TOKEN ResponseType = 2
	RESPONSETYPE_CODE     ResponseType = 3
	RESPONSETYPE_TOKEN    ResponseType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ResponseType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ID_TOKEN",
		"CODE",
		"TOKEN",
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
		"ID_TOKEN",
		"CODE",
		"TOKEN",
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
		"ID_TOKEN",
		"CODE",
		"TOKEN",
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{userExtId}/keys/{extId}/$actions/revoke Post operation
*/
type RevokeUserKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRevokeUserKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RevokeUserKeyApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RevokeUserKeyApiResponse

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

func (p *RevokeUserKeyApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RevokeUserKeyApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRevokeUserKeyApiResponse()

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

func NewRevokeUserKeyApiResponse() *RevokeUserKeyApiResponse {
	p := new(RevokeUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.RevokeUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RevokeUserKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RevokeUserKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRevokeUserKeyApiResponseData()
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
Information of SAML IDP.
*/
type SamlIdentityProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the SAML identity provider.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the SAML identity provider.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  SAML assertions for a list of custom attribute elements.
	*/
	CustomAttributes []string `json:"customAttributes,omitempty"`
	/*
	  SAML assertion email attribute element.
	*/
	EmailAttribute *string `json:"emailAttribute,omitempty"`
	/*
	  It will be used as an issuer in SAML authnRequest.
	*/
	EntityIssuer *string `json:"entityIssuer,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  SAML assertion groups attribute element.
	*/
	GroupsAttribute *string `json:"groupsAttribute,omitempty"`
	/*
	  Delimiter is used to split the value of an attribute into multiple groups.
	*/
	GroupsDelim *string `json:"groupsDelim,omitempty"`

	IdpMetadata *IdpMetadata `json:"idpMetadata,omitempty"`
	/*
	  Metadata URL that provides IDP details.
	*/
	IdpMetadataUrl *string `json:"idpMetadataUrl,omitempty"`
	/*
	  Base64 encoded metadata in XML format with IDP details.
	*/
	IdpMetadataXml *string `json:"idpMetadataXml,omitempty"`
	/*
	  Flag indicating signing of SAML authnRequests.
	*/
	IsSignedAuthnReqEnabled *bool `json:"isSignedAuthnReqEnabled,omitempty"`
	/*
	  Last updated time of the SAML identity provider.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the IDP.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The URL to redirect the user after successful authentication. Use this optional input when the user needs to be redirected to a different URL other than the default, such as when the PC is behind a reverse proxy and the user needs to be redirected to the reverse proxy URL instead of the PC URL.
	*/
	RedirectURL *string `json:"redirectURL,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  SAML assertion username attribute element.
	*/
	UsernameAttribute *string `json:"usernameAttribute,omitempty"`
}

func (p *SamlIdentityProvider) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SamlIdentityProvider

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

func (p *SamlIdentityProvider) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SamlIdentityProvider
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSamlIdentityProvider()

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
	if known.CustomAttributes != nil {
		p.CustomAttributes = known.CustomAttributes
	}
	if known.EmailAttribute != nil {
		p.EmailAttribute = known.EmailAttribute
	}
	if known.EntityIssuer != nil {
		p.EntityIssuer = known.EntityIssuer
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.GroupsAttribute != nil {
		p.GroupsAttribute = known.GroupsAttribute
	}
	if known.GroupsDelim != nil {
		p.GroupsDelim = known.GroupsDelim
	}
	if known.IdpMetadata != nil {
		p.IdpMetadata = known.IdpMetadata
	}
	if known.IdpMetadataUrl != nil {
		p.IdpMetadataUrl = known.IdpMetadataUrl
	}
	if known.IdpMetadataXml != nil {
		p.IdpMetadataXml = known.IdpMetadataXml
	}
	if known.IsSignedAuthnReqEnabled != nil {
		p.IsSignedAuthnReqEnabled = known.IsSignedAuthnReqEnabled
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
	if known.RedirectURL != nil {
		p.RedirectURL = known.RedirectURL
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UsernameAttribute != nil {
		p.UsernameAttribute = known.UsernameAttribute
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "customAttributes")
	delete(allFields, "emailAttribute")
	delete(allFields, "entityIssuer")
	delete(allFields, "extId")
	delete(allFields, "groupsAttribute")
	delete(allFields, "groupsDelim")
	delete(allFields, "idpMetadata")
	delete(allFields, "idpMetadataUrl")
	delete(allFields, "idpMetadataXml")
	delete(allFields, "isSignedAuthnReqEnabled")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "redirectURL")
	delete(allFields, "tenantId")
	delete(allFields, "usernameAttribute")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSamlIdentityProvider() *SamlIdentityProvider {
	p := new(SamlIdentityProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SamlIdentityProvider"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.EmailAttribute = new(string)
	*p.EmailAttribute = "email"
	p.UsernameAttribute = new(string)
	*p.UsernameAttribute = "name"

	return p
}

type Scope int

const (
	SCOPE_UNKNOWN        Scope = 0
	SCOPE_REDACTED       Scope = 1
	SCOPE_OPENID         Scope = 2
	SCOPE_PROFILE        Scope = 3
	SCOPE_EMAIL          Scope = 4
	SCOPE_OFFLINE_ACCESS Scope = 5
	SCOPE_GROUPS         Scope = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Scope) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"OFFLINE_ACCESS",
		"GROUPS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Scope) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"OFFLINE_ACCESS",
		"GROUPS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Scope) index(name string) Scope {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"OFFLINE_ACCESS",
		"GROUPS",
	}
	for idx := range names {
		if names[idx] == name {
			return Scope(idx)
		}
	}
	return SCOPE_UNKNOWN
}

func (e *Scope) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Scope:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Scope) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Scope) Ref() *Scope {
	return &e
}

/*
Supported OIDC scopes.
*/
type ScopesType int

const (
	SCOPESTYPE_UNKNOWN        ScopesType = 0
	SCOPESTYPE_REDACTED       ScopesType = 1
	SCOPESTYPE_OPENID         ScopesType = 2
	SCOPESTYPE_PROFILE        ScopesType = 3
	SCOPESTYPE_EMAIL          ScopesType = 4
	SCOPESTYPE_GROUPS         ScopesType = 5
	SCOPESTYPE_OFFLINE_ACCESS ScopesType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScopesType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"GROUPS",
		"OFFLINE_ACCESS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScopesType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"GROUPS",
		"OFFLINE_ACCESS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScopesType) index(name string) ScopesType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPENID",
		"PROFILE",
		"EMAIL",
		"GROUPS",
		"OFFLINE_ACCESS",
	}
	for idx := range names {
		if names[idx] == name {
			return ScopesType(idx)
		}
	}
	return SCOPESTYPE_UNKNOWN
}

func (e *ScopesType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScopesType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScopesType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScopesType) Ref() *ScopesType {
	return &e
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services/{extId}/$actions/search Post operation
*/
type SearchDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSearchDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *SearchDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SearchDirectoryServiceApiResponse

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

func (p *SearchDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SearchDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSearchDirectoryServiceApiResponse()

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

func NewSearchDirectoryServiceApiResponse() *SearchDirectoryServiceApiResponse {
	p := new(SearchDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SearchDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SearchDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SearchDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSearchDirectoryServiceApiResponseData()
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
Supported JWS signing algorithms.
*/
type SigningAlgoType int

const (
	SIGNINGALGOTYPE_UNKNOWN  SigningAlgoType = 0
	SIGNINGALGOTYPE_REDACTED SigningAlgoType = 1
	SIGNINGALGOTYPE_RS256    SigningAlgoType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SigningAlgoType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RS256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SigningAlgoType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RS256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SigningAlgoType) index(name string) SigningAlgoType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RS256",
	}
	for idx := range names {
		if names[idx] == name {
			return SigningAlgoType(idx)
		}
	}
	return SIGNINGALGOTYPE_UNKNOWN
}

func (e *SigningAlgoType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SigningAlgoType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SigningAlgoType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SigningAlgoType) Ref() *SigningAlgoType {
	return &e
}

/*
Type of subject for token.
*/
type SubjectType int

const (
	SUBJECTTYPE_UNKNOWN                                   SubjectType = 0
	SUBJECTTYPE_REDACTED                                  SubjectType = 1
	SUBJECTTYPE_URN_IETF_PARAMS_OAUTH_TOKEN_TYPE_ID_TOKEN SubjectType = 2
	SUBJECTTYPE_URN_IETF_PARAMS_OAUTH_TOKEN_TYPE_API_KEY  SubjectType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SubjectType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:ID_TOKEN",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:API_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SubjectType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:ID_TOKEN",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:API_KEY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SubjectType) index(name string) SubjectType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:ID_TOKEN",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:API_KEY",
	}
	for idx := range names {
		if names[idx] == name {
			return SubjectType(idx)
		}
	}
	return SUBJECTTYPE_UNKNOWN
}

func (e *SubjectType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SubjectType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SubjectType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SubjectType) Ref() *SubjectType {
	return &e
}

/*
System configuration resource.
*/
type SystemConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Last update time for the system configuration.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  Name Value Pair of the system configuration.
	*/
	NameValuePair []import3.KVPair `json:"nameValuePair"`
}

func (p *SystemConfig) MarshalJSON() ([]byte, error) {
	type SystemConfigProxy SystemConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SystemConfigProxy
		NameValuePair []import3.KVPair `json:"nameValuePair,omitempty"`
	}{
		SystemConfigProxy: (*SystemConfigProxy)(p),
		NameValuePair:     p.NameValuePair,
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

func (p *SystemConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SystemConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSystemConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.NameValuePair != nil {
		p.NameValuePair = known.NameValuePair
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "nameValuePair")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSystemConfig() *SystemConfig {
	p := new(SystemConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SystemConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information regarding the tenant that the user belongs to.
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the tenant that the user belongs to.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
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
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Tenant"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the OIDC Token request.
*/
type TokenRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Audience for the OIDC token request.
	*/
	Audience *string `json:"audience,omitempty"`
	/*
	  Client assertion for the OIDC token request.
	*/
	ClientAssertion *string `json:"client_assertion,omitempty"`
	/*
	  Client assertion type for the OIDC token request.
	*/
	ClientAssertionType *string `json:"client_assertion_type,omitempty"`
	/*
	  Client identifier for the OIDC token request.
	*/
	ClientId *string `json:"client_id"`
	/*
	  Code for the OIDC token.
	*/
	Code *string `json:"code,omitempty"`
	/*
	  Connector identifier for the connector of a user.
	*/
	ConnectorId *string `json:"connector_id,omitempty"`

	GrantType *GrantType `json:"grant_type"`
	/*
	  Non-tenant issuer for the OIDC token request.
	*/
	NonTenantIssuer *string `json:"non_tenant_issuer,omitempty"`
	/*
	  Redirect URI for the OIDC token request.
	*/
	RedirectUri *string `json:"redirect_uri,omitempty"`
	/*
	  Request token to get the OIDC token.
	*/
	RefreshToken *string `json:"refresh_token,omitempty"`
	/*
	  Scopes for the OIDC token.
	*/
	Scope *string `json:"scope"`
	/*
	  Subject token for the OIDC token request.
	*/
	SubjectToken *string `json:"subject_token,omitempty"`

	SubjectTokenType *SubjectType `json:"subject_token_type,omitempty"`
}

func (p *TokenRequest) MarshalJSON() ([]byte, error) {
	type TokenRequestProxy TokenRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*TokenRequestProxy
		ClientId  *string    `json:"client_id,omitempty"`
		GrantType *GrantType `json:"grant_type,omitempty"`
		Scope     *string    `json:"scope,omitempty"`
	}{
		TokenRequestProxy: (*TokenRequestProxy)(p),
		ClientId:          p.ClientId,
		GrantType:         p.GrantType,
		Scope:             p.Scope,
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

func (p *TokenRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TokenRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTokenRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Audience != nil {
		p.Audience = known.Audience
	}
	if known.ClientAssertion != nil {
		p.ClientAssertion = known.ClientAssertion
	}
	if known.ClientAssertionType != nil {
		p.ClientAssertionType = known.ClientAssertionType
	}
	if known.ClientId != nil {
		p.ClientId = known.ClientId
	}
	if known.Code != nil {
		p.Code = known.Code
	}
	if known.ConnectorId != nil {
		p.ConnectorId = known.ConnectorId
	}
	if known.GrantType != nil {
		p.GrantType = known.GrantType
	}
	if known.NonTenantIssuer != nil {
		p.NonTenantIssuer = known.NonTenantIssuer
	}
	if known.RedirectUri != nil {
		p.RedirectUri = known.RedirectUri
	}
	if known.RefreshToken != nil {
		p.RefreshToken = known.RefreshToken
	}
	if known.Scope != nil {
		p.Scope = known.Scope
	}
	if known.SubjectToken != nil {
		p.SubjectToken = known.SubjectToken
	}
	if known.SubjectTokenType != nil {
		p.SubjectTokenType = known.SubjectTokenType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "audience")
	delete(allFields, "client_assertion")
	delete(allFields, "client_assertion_type")
	delete(allFields, "client_id")
	delete(allFields, "code")
	delete(allFields, "connector_id")
	delete(allFields, "grant_type")
	delete(allFields, "non_tenant_issuer")
	delete(allFields, "redirect_uri")
	delete(allFields, "refresh_token")
	delete(allFields, "scope")
	delete(allFields, "subject_token")
	delete(allFields, "subject_token_type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTokenRequest() *TokenRequest {
	p := new(TokenRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.TokenRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.1.b2/authn/cert-auth-providers/{extId} Put operation
*/
type UpdateCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateCertAuthProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateCertAuthProviderApiResponse

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

func (p *UpdateCertAuthProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateCertAuthProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateCertAuthProviderApiResponse()

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

func NewUpdateCertAuthProviderApiResponse() *UpdateCertAuthProviderApiResponse {
	p := new(UpdateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateCertAuthProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateCertAuthProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateCertAuthProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/directory-services/{extId} Put operation
*/
type UpdateDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateDirectoryServiceApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateDirectoryServiceApiResponse

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

func (p *UpdateDirectoryServiceApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateDirectoryServiceApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateDirectoryServiceApiResponse()

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

func NewUpdateDirectoryServiceApiResponse() *UpdateDirectoryServiceApiResponse {
	p := new(UpdateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/saml-identity-providers/{extId} Put operation
*/
type UpdateSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateSamlIdentityProviderApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateSamlIdentityProviderApiResponse

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

func (p *UpdateSamlIdentityProviderApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateSamlIdentityProviderApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateSamlIdentityProviderApiResponse()

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

func NewUpdateSamlIdentityProviderApiResponse() *UpdateSamlIdentityProviderApiResponse {
	p := new(UpdateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSamlIdentityProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSamlIdentityProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSamlIdentityProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/users/{extId} Put operation
*/
type UpdateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateUserApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateUserApiResponse

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

func (p *UpdateUserApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateUserApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateUserApiResponse()

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

func NewUpdateUserApiResponse() *UpdateUserApiResponse {
	p := new(UpdateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateUserApiResponseData()
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
REST response for all response codes in API path /iam/v4.1.b2/authn/welcome-banner Put operation
*/
type UpdateWelcomeBannerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateWelcomeBannerApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateWelcomeBannerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateWelcomeBannerApiResponse

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

func (p *UpdateWelcomeBannerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateWelcomeBannerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateWelcomeBannerApiResponse()

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

func NewUpdateWelcomeBannerApiResponse() *UpdateWelcomeBannerApiResponse {
	p := new(UpdateWelcomeBannerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateWelcomeBannerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateWelcomeBannerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateWelcomeBannerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateWelcomeBannerApiResponseData()
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
Use for the OIDC key.
*/
type UseType int

const (
	USETYPE_UNKNOWN  UseType = 0
	USETYPE_REDACTED UseType = 1
	USETYPE_SIG      UseType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UseType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SIG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UseType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SIG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UseType) index(name string) UseType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SIG",
	}
	for idx := range names {
		if names[idx] == name {
			return UseType(idx)
		}
	}
	return USETYPE_UNKNOWN
}

func (e *UseType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UseType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UseType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UseType) Ref() *UseType {
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
	AdditionalAttributes []import3.KVPair `json:"additionalAttributes,omitempty"`
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
	Links []import2.ApiLink `json:"links,omitempty"`
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
	*p = *NewUser()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AdditionalAttributes != nil {
		p.AdditionalAttributes = known.AdditionalAttributes
	}
	if known.BucketsAccessKeys != nil {
		p.BucketsAccessKeys = known.BucketsAccessKeys
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.CreationType != nil {
		p.CreationType = known.CreationType
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.EmailId != nil {
		p.EmailId = known.EmailId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FirstName != nil {
		p.FirstName = known.FirstName
	}
	if known.IdpId != nil {
		p.IdpId = known.IdpId
	}
	if known.IsForceResetPasswordEnabled != nil {
		p.IsForceResetPasswordEnabled = known.IsForceResetPasswordEnabled
	}
	if known.LastLoginTime != nil {
		p.LastLoginTime = known.LastLoginTime
	}
	if known.LastName != nil {
		p.LastName = known.LastName
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Locale != nil {
		p.Locale = known.Locale
	}
	if known.MiddleInitial != nil {
		p.MiddleInitial = known.MiddleInitial
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Region != nil {
		p.Region = known.Region
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UserType != nil {
		p.UserType = known.UserType
	}
	if known.Username != nil {
		p.Username = known.Username
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUser() *User {
	p := new(User)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.User"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User configuration for OpenLDAP directory service.
*/
type UserConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Object class in the OpenLDAP system that corresponds to the users.
	*/
	UserObjectClass *string `json:"userObjectClass"`
	/*
	  Base DN for user search.
	*/
	UserSearchBase *string `json:"userSearchBase"`
	/*
	  Unique identifier for each user that can be used in authentication.
	*/
	UsernameAttribute *string `json:"usernameAttribute"`
}

func (p *UserConfiguration) MarshalJSON() ([]byte, error) {
	type UserConfigurationProxy UserConfiguration

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UserConfigurationProxy
		UserObjectClass   *string `json:"userObjectClass,omitempty"`
		UserSearchBase    *string `json:"userSearchBase,omitempty"`
		UsernameAttribute *string `json:"usernameAttribute,omitempty"`
	}{
		UserConfigurationProxy: (*UserConfigurationProxy)(p),
		UserObjectClass:        p.UserObjectClass,
		UserSearchBase:         p.UserSearchBase,
		UsernameAttribute:      p.UsernameAttribute,
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

func (p *UserConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.UserObjectClass != nil {
		p.UserObjectClass = known.UserObjectClass
	}
	if known.UserSearchBase != nil {
		p.UserSearchBase = known.UserSearchBase
	}
	if known.UsernameAttribute != nil {
		p.UsernameAttribute = known.UsernameAttribute
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "userObjectClass")
	delete(allFields, "userSearchBase")
	delete(allFields, "usernameAttribute")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserConfiguration() *UserConfiguration {
	p := new(UserConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the user group.
*/
type UserGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or Service who created the user group.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the user group.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Identifier for the user group in the form of a distinguished name.
	*/
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupType *GroupType `json:"groupType,omitempty"`
	/*
	  Identifier of the IDP for the user group.
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Last updated time of the user group.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Common Name of the user group.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UserGroup) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UserGroup

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

func (p *UserGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserGroup()

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
	if known.DistinguishedName != nil {
		p.DistinguishedName = known.DistinguishedName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.GroupType != nil {
		p.GroupType = known.GroupType
	}
	if known.IdpId != nil {
		p.IdpId = known.IdpId
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
	delete(allFields, "distinguishedName")
	delete(allFields, "extId")
	delete(allFields, "groupType")
	delete(allFields, "idpId")
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

func NewUserGroup() *UserGroup {
	p := new(UserGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User Group configuration for OpenLDAP directory service.
*/
type UserGroupConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute in a group that associates users to the group.
	*/
	GroupMemberAttribute *string `json:"groupMemberAttribute"`
	/*
	  User attribute value that will be used in group entity to associate user to the group.
	*/
	GroupMemberAttributeValue *string `json:"groupMemberAttributeValue"`
	/*
	  Object class in the OpenLDAP system that corresponds to groups.
	*/
	GroupObjectClass *string `json:"groupObjectClass"`
	/*
	  Base DN for group search.
	*/
	GroupSearchBase *string `json:"groupSearchBase"`
}

func (p *UserGroupConfiguration) MarshalJSON() ([]byte, error) {
	type UserGroupConfigurationProxy UserGroupConfiguration

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UserGroupConfigurationProxy
		GroupMemberAttribute      *string `json:"groupMemberAttribute,omitempty"`
		GroupMemberAttributeValue *string `json:"groupMemberAttributeValue,omitempty"`
		GroupObjectClass          *string `json:"groupObjectClass,omitempty"`
		GroupSearchBase           *string `json:"groupSearchBase,omitempty"`
	}{
		UserGroupConfigurationProxy: (*UserGroupConfigurationProxy)(p),
		GroupMemberAttribute:        p.GroupMemberAttribute,
		GroupMemberAttributeValue:   p.GroupMemberAttributeValue,
		GroupObjectClass:            p.GroupObjectClass,
		GroupSearchBase:             p.GroupSearchBase,
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

func (p *UserGroupConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserGroupConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserGroupConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.GroupMemberAttribute != nil {
		p.GroupMemberAttribute = known.GroupMemberAttribute
	}
	if known.GroupMemberAttributeValue != nil {
		p.GroupMemberAttributeValue = known.GroupMemberAttributeValue
	}
	if known.GroupObjectClass != nil {
		p.GroupObjectClass = known.GroupObjectClass
	}
	if known.GroupSearchBase != nil {
		p.GroupSearchBase = known.GroupSearchBase
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "groupMemberAttribute")
	delete(allFields, "groupMemberAttributeValue")
	delete(allFields, "groupObjectClass")
	delete(allFields, "groupSearchBase")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserGroupConfiguration() *UserGroupConfiguration {
	p := new(UserGroupConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserGroupConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information to change state of user.
*/
type UserStateUpdate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Status *UserStatusType `json:"status"`
}

func (p *UserStateUpdate) MarshalJSON() ([]byte, error) {
	type UserStateUpdateProxy UserStateUpdate

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UserStateUpdateProxy
		Status *UserStatusType `json:"status,omitempty"`
	}{
		UserStateUpdateProxy: (*UserStateUpdateProxy)(p),
		Status:               p.Status,
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

func (p *UserStateUpdate) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserStateUpdate
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserStateUpdate()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUserStateUpdate() *UserStateUpdate {
	p := new(UserStateUpdate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserStateUpdate"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User state update response object.
*/
type UserStateUpdateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *UserStateUpdateResponse) MarshalJSON() ([]byte, error) {
	type UserStateUpdateResponseProxy UserStateUpdateResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UserStateUpdateResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		UserStateUpdateResponseProxy: (*UserStateUpdateResponseProxy)(p),
		Message:                      p.Message,
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

func (p *UserStateUpdateResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UserStateUpdateResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUserStateUpdateResponse()

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

func NewUserStateUpdateResponse() *UserStateUpdateResponse {
	p := new(UserStateUpdateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserStateUpdateResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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

/*
Request body for validating credentials.
*/
type ValidateCredentialsRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password of the user.
	*/
	Password *string `json:"password"`
	/*
	  Identifier of the user.
	*/
	Username *string `json:"username"`
}

func (p *ValidateCredentialsRequest) MarshalJSON() ([]byte, error) {
	type ValidateCredentialsRequestProxy ValidateCredentialsRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ValidateCredentialsRequestProxy
		Password *string `json:"password,omitempty"`
		Username *string `json:"username,omitempty"`
	}{
		ValidateCredentialsRequestProxy: (*ValidateCredentialsRequestProxy)(p),
		Password:                        p.Password,
		Username:                        p.Username,
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

func (p *ValidateCredentialsRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ValidateCredentialsRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewValidateCredentialsRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewValidateCredentialsRequest() *ValidateCredentialsRequest {
	p := new(ValidateCredentialsRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ValidateCredentialsRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Welcome banner API.
*/
type WelcomeBanner struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Content of the welcome banner.
	*/
	Content *string `json:"content,omitempty"`
	/*
	  Creation time of the welcome banner.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Flag to denote whether the welcome banner is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Last updated time of the welcome banner.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
}

func (p *WelcomeBanner) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias WelcomeBanner

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

func (p *WelcomeBanner) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias WelcomeBanner
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewWelcomeBanner()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Content != nil {
		p.Content = known.Content
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "content")
	delete(allFields, "createdTime")
	delete(allFields, "isEnabled")
	delete(allFields, "lastUpdatedTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewWelcomeBanner() *WelcomeBanner {
	p := new(WelcomeBanner)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.WelcomeBanner"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

type OneOfDeleteCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteCertAuthProviderApiResponseData() *OneOfDeleteCertAuthProviderApiResponseData {
	p := new(OneOfDeleteCertAuthProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteCertAuthProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteCertAuthProviderApiResponseData is nil"))
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

func (p *OneOfDeleteCertAuthProviderApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteCertAuthProviderApiResponseData"))
}

func (p *OneOfDeleteCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteCertAuthProviderApiResponseData")
}

type OneOfCreateCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateCertAuthProviderApiResponseData() *OneOfCreateCertAuthProviderApiResponseData {
	p := new(OneOfCreateCertAuthProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCertAuthProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCertAuthProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case CertAuthProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CertAuthProvider)
		}
		*p.oneOfType0 = v.(CertAuthProvider)
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

func (p *OneOfCreateCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(CertAuthProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.CertAuthProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CertAuthProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCertAuthProviderApiResponseData"))
}

func (p *OneOfCreateCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCertAuthProviderApiResponseData")
}

type OneOfRevokeUserKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2001 *import1.AppMessage    `json:"-"`
}

func NewOneOfRevokeUserKeyApiResponseData() *OneOfRevokeUserKeyApiResponseData {
	p := new(OneOfRevokeUserKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRevokeUserKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRevokeUserKeyApiResponseData is nil"))
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
	case import1.AppMessage:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.AppMessage)
		}
		*p.oneOfType2001 = v.(import1.AppMessage)
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

func (p *OneOfRevokeUserKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfRevokeUserKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(import1.AppMessage)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "iam.v4.error.AppMessage" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.AppMessage)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRevokeUserKeyApiResponseData"))
}

func (p *OneOfRevokeUserKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfRevokeUserKeyApiResponseData")
}

type OneOfListSamlIdentityProvidersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []SamlIdentityProvider `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListSamlIdentityProvidersApiResponseData() *OneOfListSamlIdentityProvidersApiResponseData {
	p := new(OneOfListSamlIdentityProvidersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSamlIdentityProvidersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSamlIdentityProvidersApiResponseData is nil"))
	}
	switch v.(type) {
	case []SamlIdentityProvider:
		p.oneOfType0 = v.([]SamlIdentityProvider)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.SamlIdentityProvider>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.SamlIdentityProvider>"
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

func (p *OneOfListSamlIdentityProvidersApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authn.SamlIdentityProvider>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListSamlIdentityProvidersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]SamlIdentityProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.SamlIdentityProvider" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.SamlIdentityProvider>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.SamlIdentityProvider>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSamlIdentityProvidersApiResponseData"))
}

func (p *OneOfListSamlIdentityProvidersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authn.SamlIdentityProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListSamlIdentityProvidersApiResponseData")
}

type OneOfGetCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetCertAuthProviderApiResponseData() *OneOfGetCertAuthProviderApiResponseData {
	p := new(OneOfGetCertAuthProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCertAuthProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCertAuthProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case CertAuthProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CertAuthProvider)
		}
		*p.oneOfType0 = v.(CertAuthProvider)
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

func (p *OneOfGetCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(CertAuthProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.CertAuthProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CertAuthProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCertAuthProviderApiResponseData"))
}

func (p *OneOfGetCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetCertAuthProviderApiResponseData")
}

type OneOfDeleteUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteUserGroupApiResponseData() *OneOfDeleteUserGroupApiResponseData {
	p := new(OneOfDeleteUserGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteUserGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteUserGroupApiResponseData is nil"))
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

func (p *OneOfDeleteUserGroupApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteUserGroupApiResponseData"))
}

func (p *OneOfDeleteUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteUserGroupApiResponseData")
}

type OneOfGetSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetSamlIdentityProviderApiResponseData() *OneOfGetSamlIdentityProviderApiResponseData {
	p := new(OneOfGetSamlIdentityProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSamlIdentityProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case SamlIdentityProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SamlIdentityProvider)
		}
		*p.oneOfType0 = v.(SamlIdentityProvider)
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

func (p *OneOfGetSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SamlIdentityProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.SamlIdentityProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SamlIdentityProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSamlIdentityProviderApiResponseData"))
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSamlIdentityProviderApiResponseData")
}

type OneOfDeleteDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteDirectoryServiceApiResponseData() *OneOfDeleteDirectoryServiceApiResponseData {
	p := new(OneOfDeleteDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteDirectoryServiceApiResponseData is nil"))
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

func (p *OneOfDeleteDirectoryServiceApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteDirectoryServiceApiResponseData"))
}

func (p *OneOfDeleteDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteDirectoryServiceApiResponseData")
}

type OneOfUpdateSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSamlIdentityProviderApiResponseData() *OneOfUpdateSamlIdentityProviderApiResponseData {
	p := new(OneOfUpdateSamlIdentityProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSamlIdentityProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case SamlIdentityProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SamlIdentityProvider)
		}
		*p.oneOfType0 = v.(SamlIdentityProvider)
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

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SamlIdentityProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.SamlIdentityProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SamlIdentityProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSamlIdentityProviderApiResponseData"))
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSamlIdentityProviderApiResponseData")
}

type OneOfCreateSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateSamlIdentityProviderApiResponseData() *OneOfCreateSamlIdentityProviderApiResponseData {
	p := new(OneOfCreateSamlIdentityProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateSamlIdentityProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case SamlIdentityProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SamlIdentityProvider)
		}
		*p.oneOfType0 = v.(SamlIdentityProvider)
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

func (p *OneOfCreateSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SamlIdentityProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.SamlIdentityProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SamlIdentityProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSamlIdentityProviderApiResponseData"))
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSamlIdentityProviderApiResponseData")
}

type OneOfCreateUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *User                  `json:"-"`
}

func NewOneOfCreateUserApiResponseData() *OneOfCreateUserApiResponseData {
	p := new(OneOfCreateUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateUserApiResponseData is nil"))
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
	case User:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(User)
		}
		*p.oneOfType0 = v.(User)
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

func (p *OneOfCreateUserApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateUserApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(User)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.User" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(User)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUserApiResponseData"))
}

func (p *OneOfCreateUserApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUserApiResponseData")
}

type OneOfGetUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UserGroup             `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetUserGroupApiResponseData() *OneOfGetUserGroupApiResponseData {
	p := new(OneOfGetUserGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUserGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUserGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case UserGroup:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserGroup)
		}
		*p.oneOfType0 = v.(UserGroup)
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

func (p *OneOfGetUserGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.UserGroup" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserGroup)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUserGroupApiResponseData"))
}

func (p *OneOfGetUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetUserGroupApiResponseData")
}

type OneOfGetSamlSpMetadataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *string                `json:"-"`
}

func NewOneOfGetSamlSpMetadataApiResponseData() *OneOfGetSamlSpMetadataApiResponseData {
	p := new(OneOfGetSamlSpMetadataApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSamlSpMetadataApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSamlSpMetadataApiResponseData is nil"))
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
	case string:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(string)
		}
		*p.oneOfType0 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSamlSpMetadataApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "String" == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSamlSpMetadataApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(string)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(string)
		}
		*p.oneOfType0 = *vOneOfType0
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSamlSpMetadataApiResponseData"))
}

func (p *OneOfGetSamlSpMetadataApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSamlSpMetadataApiResponseData")
}

type OneOfGetUserKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Key                   `json:"-"`
}

func NewOneOfGetUserKeyApiResponseData() *OneOfGetUserKeyApiResponseData {
	p := new(OneOfGetUserKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUserKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUserKeyApiResponseData is nil"))
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
	case Key:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Key)
		}
		*p.oneOfType0 = v.(Key)
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

func (p *OneOfGetUserKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetUserKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Key)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.Key" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Key)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUserKeyApiResponseData"))
}

func (p *OneOfGetUserKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetUserKeyApiResponseData")
}

type OneOfResetUserPasswordApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *PasswordResetResponse `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfResetUserPasswordApiResponseData() *OneOfResetUserPasswordApiResponseData {
	p := new(OneOfResetUserPasswordApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfResetUserPasswordApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfResetUserPasswordApiResponseData is nil"))
	}
	switch v.(type) {
	case PasswordResetResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PasswordResetResponse)
		}
		*p.oneOfType0 = v.(PasswordResetResponse)
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

func (p *OneOfResetUserPasswordApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfResetUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(PasswordResetResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.PasswordResetResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PasswordResetResponse)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetUserPasswordApiResponseData"))
}

func (p *OneOfResetUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfResetUserPasswordApiResponseData")
}

type OneOfGetDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetDirectoryServiceApiResponseData() *OneOfGetDirectoryServiceApiResponseData {
	p := new(OneOfGetDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDirectoryServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case DirectoryService:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryService)
		}
		*p.oneOfType0 = v.(DirectoryService)
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

func (p *OneOfGetDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DirectoryService)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.DirectoryService" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryService)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDirectoryServiceApiResponseData"))
}

func (p *OneOfGetDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetDirectoryServiceApiResponseData")
}

type OneOfUpdateDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateDirectoryServiceApiResponseData() *OneOfUpdateDirectoryServiceApiResponseData {
	p := new(OneOfUpdateDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateDirectoryServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case DirectoryService:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryService)
		}
		*p.oneOfType0 = v.(DirectoryService)
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

func (p *OneOfUpdateDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DirectoryService)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.DirectoryService" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryService)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDirectoryServiceApiResponseData"))
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDirectoryServiceApiResponseData")
}

type OneOfListUsersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []User                 `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListUsersApiResponseData() *OneOfListUsersApiResponseData {
	p := new(OneOfListUsersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUsersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUsersApiResponseData is nil"))
	}
	switch v.(type) {
	case []User:
		p.oneOfType0 = v.([]User)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.User>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.User>"
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

func (p *OneOfListUsersApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authn.User>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListUsersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]User)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.User" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.User>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.User>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUsersApiResponseData"))
}

func (p *OneOfListUsersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authn.User>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListUsersApiResponseData")
}

type OneOfDeleteSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSamlIdentityProviderApiResponseData() *OneOfDeleteSamlIdentityProviderApiResponseData {
	p := new(OneOfDeleteSamlIdentityProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSamlIdentityProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSamlIdentityProviderApiResponseData is nil"))
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

func (p *OneOfDeleteSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSamlIdentityProviderApiResponseData"))
}

func (p *OneOfDeleteSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSamlIdentityProviderApiResponseData")
}

type OneOfChangeUserPasswordApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    *PasswordChangeResponse `json:"-"`
	oneOfType400  *import1.ErrorResponse  `json:"-"`
}

func NewOneOfChangeUserPasswordApiResponseData() *OneOfChangeUserPasswordApiResponseData {
	p := new(OneOfChangeUserPasswordApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfChangeUserPasswordApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfChangeUserPasswordApiResponseData is nil"))
	}
	switch v.(type) {
	case PasswordChangeResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PasswordChangeResponse)
		}
		*p.oneOfType0 = v.(PasswordChangeResponse)
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

func (p *OneOfChangeUserPasswordApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangeUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(PasswordChangeResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.PasswordChangeResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PasswordChangeResponse)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangeUserPasswordApiResponseData"))
}

func (p *OneOfChangeUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangeUserPasswordApiResponseData")
}

type OneOfGetUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *User                  `json:"-"`
}

func NewOneOfGetUserApiResponseData() *OneOfGetUserApiResponseData {
	p := new(OneOfGetUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUserApiResponseData is nil"))
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
	case User:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(User)
		}
		*p.oneOfType0 = v.(User)
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

func (p *OneOfGetUserApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetUserApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(User)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.User" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(User)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUserApiResponseData"))
}

func (p *OneOfGetUserApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetUserApiResponseData")
}

type OneOfGetWelcomeBannerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *WelcomeBanner         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetWelcomeBannerApiResponseData() *OneOfGetWelcomeBannerApiResponseData {
	p := new(OneOfGetWelcomeBannerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetWelcomeBannerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetWelcomeBannerApiResponseData is nil"))
	}
	switch v.(type) {
	case WelcomeBanner:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(WelcomeBanner)
		}
		*p.oneOfType0 = v.(WelcomeBanner)
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

func (p *OneOfGetWelcomeBannerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetWelcomeBannerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(WelcomeBanner)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.WelcomeBanner" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(WelcomeBanner)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetWelcomeBannerApiResponseData"))
}

func (p *OneOfGetWelcomeBannerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetWelcomeBannerApiResponseData")
}

type OneOfCreateUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *UserGroup             `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateUserGroupApiResponseData() *OneOfCreateUserGroupApiResponseData {
	p := new(OneOfCreateUserGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateUserGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateUserGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case UserGroup:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserGroup)
		}
		*p.oneOfType0 = v.(UserGroup)
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

func (p *OneOfCreateUserGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.UserGroup" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserGroup)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUserGroupApiResponseData"))
}

func (p *OneOfCreateUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUserGroupApiResponseData")
}

type OneOfUpdateWelcomeBannerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *WelcomeBanner         `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateWelcomeBannerApiResponseData() *OneOfUpdateWelcomeBannerApiResponseData {
	p := new(OneOfUpdateWelcomeBannerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateWelcomeBannerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateWelcomeBannerApiResponseData is nil"))
	}
	switch v.(type) {
	case WelcomeBanner:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(WelcomeBanner)
		}
		*p.oneOfType0 = v.(WelcomeBanner)
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

func (p *OneOfUpdateWelcomeBannerApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateWelcomeBannerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(WelcomeBanner)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.WelcomeBanner" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(WelcomeBanner)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateWelcomeBannerApiResponseData"))
}

func (p *OneOfUpdateWelcomeBannerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateWelcomeBannerApiResponseData")
}

type OneOfCreateKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Key                   `json:"-"`
}

func NewOneOfCreateKeyApiResponseData() *OneOfCreateKeyApiResponseData {
	p := new(OneOfCreateKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateKeyApiResponseData is nil"))
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
	case Key:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Key)
		}
		*p.oneOfType0 = v.(Key)
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

func (p *OneOfCreateKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(Key)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.Key" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Key)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateKeyApiResponseData"))
}

func (p *OneOfCreateKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateKeyApiResponseData")
}

type OneOfListLoginProvidersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []LoginProvider        `json:"-"`
}

func NewOneOfListLoginProvidersApiResponseData() *OneOfListLoginProvidersApiResponseData {
	p := new(OneOfListLoginProvidersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLoginProvidersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLoginProvidersApiResponseData is nil"))
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
	case []LoginProvider:
		p.oneOfType0 = v.([]LoginProvider)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.LoginProvider>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.LoginProvider>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLoginProvidersApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.LoginProvider>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListLoginProvidersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]LoginProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.LoginProvider" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.LoginProvider>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.LoginProvider>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLoginProvidersApiResponseData"))
}

func (p *OneOfListLoginProvidersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.LoginProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListLoginProvidersApiResponseData")
}

type OneOfListUserGroupsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []UserGroup            `json:"-"`
}

func NewOneOfListUserGroupsApiResponseData() *OneOfListUserGroupsApiResponseData {
	p := new(OneOfListUserGroupsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUserGroupsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUserGroupsApiResponseData is nil"))
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
	case []UserGroup:
		p.oneOfType0 = v.([]UserGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.UserGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.UserGroup>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUserGroupsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.UserGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUserGroupsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]UserGroup)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.UserGroup" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.UserGroup>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.UserGroup>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUserGroupsApiResponseData"))
}

func (p *OneOfListUserGroupsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.UserGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUserGroupsApiResponseData")
}

type OneOfListUserKeysApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Key                  `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListUserKeysApiResponseData() *OneOfListUserKeysApiResponseData {
	p := new(OneOfListUserKeysApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUserKeysApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUserKeysApiResponseData is nil"))
	}
	switch v.(type) {
	case []Key:
		p.oneOfType0 = v.([]Key)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.Key>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.Key>"
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

func (p *OneOfListUserKeysApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authn.Key>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListUserKeysApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Key)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.Key" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.Key>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.Key>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUserKeysApiResponseData"))
}

func (p *OneOfListUserKeysApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authn.Key>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListUserKeysApiResponseData")
}

type OneOfActivateUserApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType0    *UserStateUpdateResponse `json:"-"`
	oneOfType400  *import1.ErrorResponse   `json:"-"`
}

func NewOneOfActivateUserApiResponseData() *OneOfActivateUserApiResponseData {
	p := new(OneOfActivateUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfActivateUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfActivateUserApiResponseData is nil"))
	}
	switch v.(type) {
	case UserStateUpdateResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UserStateUpdateResponse)
		}
		*p.oneOfType0 = v.(UserStateUpdateResponse)
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

func (p *OneOfActivateUserApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfActivateUserApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UserStateUpdateResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.UserStateUpdateResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UserStateUpdateResponse)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfActivateUserApiResponseData"))
}

func (p *OneOfActivateUserApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfActivateUserApiResponseData")
}

type OneOfCreateDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateDirectoryServiceApiResponseData() *OneOfCreateDirectoryServiceApiResponseData {
	p := new(OneOfCreateDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateDirectoryServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case DirectoryService:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryService)
		}
		*p.oneOfType0 = v.(DirectoryService)
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

func (p *OneOfCreateDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DirectoryService)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.DirectoryService" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryService)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDirectoryServiceApiResponseData"))
}

func (p *OneOfCreateDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDirectoryServiceApiResponseData")
}

type OneOfListDirectoryServicesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []DirectoryService     `json:"-"`
}

func NewOneOfListDirectoryServicesApiResponseData() *OneOfListDirectoryServicesApiResponseData {
	p := new(OneOfListDirectoryServicesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDirectoryServicesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDirectoryServicesApiResponseData is nil"))
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
	case []DirectoryService:
		p.oneOfType0 = v.([]DirectoryService)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.DirectoryService>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.DirectoryService>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListDirectoryServicesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.DirectoryService>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListDirectoryServicesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]DirectoryService)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.DirectoryService" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.DirectoryService>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.DirectoryService>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDirectoryServicesApiResponseData"))
}

func (p *OneOfListDirectoryServicesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.DirectoryService>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListDirectoryServicesApiResponseData")
}

type OneOfKeyKeyDetails struct {
	Discriminator *string           `json:"-"`
	ObjectType_   *string           `json:"-"`
	oneOfType0    *ApiKeyDetails    `json:"-"`
	oneOfType1    *ObjectKeyDetails `json:"-"`
}

func NewOneOfKeyKeyDetails() *OneOfKeyKeyDetails {
	p := new(OneOfKeyKeyDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfKeyKeyDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfKeyKeyDetails is nil"))
	}
	switch v.(type) {
	case ApiKeyDetails:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiKeyDetails)
		}
		*p.oneOfType0 = v.(ApiKeyDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ObjectKeyDetails:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ObjectKeyDetails)
		}
		*p.oneOfType1 = v.(ObjectKeyDetails)
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

func (p *OneOfKeyKeyDetails) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfKeyKeyDetails) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ApiKeyDetails)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ApiKeyDetails" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiKeyDetails)
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
	vOneOfType1 := new(ObjectKeyDetails)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "iam.v4.authn.ObjectKeyDetails" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ObjectKeyDetails)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfKeyKeyDetails"))
}

func (p *OneOfKeyKeyDetails) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfKeyKeyDetails")
}

type OneOfGetSamlIdpSpMetadataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *string                `json:"-"`
}

func NewOneOfGetSamlIdpSpMetadataApiResponseData() *OneOfGetSamlIdpSpMetadataApiResponseData {
	p := new(OneOfGetSamlIdpSpMetadataApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSamlIdpSpMetadataApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSamlIdpSpMetadataApiResponseData is nil"))
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
	case string:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(string)
		}
		*p.oneOfType0 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSamlIdpSpMetadataApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "String" == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSamlIdpSpMetadataApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(string)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(string)
		}
		*p.oneOfType0 = *vOneOfType0
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSamlIdpSpMetadataApiResponseData"))
}

func (p *OneOfGetSamlIdpSpMetadataApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSamlIdpSpMetadataApiResponseData")
}

type OneOfUpdateCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfUpdateCertAuthProviderApiResponseData() *OneOfUpdateCertAuthProviderApiResponseData {
	p := new(OneOfUpdateCertAuthProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateCertAuthProviderApiResponseData is nil"))
	}
	switch v.(type) {
	case CertAuthProvider:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CertAuthProvider)
		}
		*p.oneOfType0 = v.(CertAuthProvider)
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

func (p *OneOfUpdateCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(CertAuthProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.CertAuthProvider" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CertAuthProvider)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCertAuthProviderApiResponseData"))
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCertAuthProviderApiResponseData")
}

type OneOfUpdateUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *User                  `json:"-"`
}

func NewOneOfUpdateUserApiResponseData() *OneOfUpdateUserApiResponseData {
	p := new(OneOfUpdateUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateUserApiResponseData is nil"))
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
	case User:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(User)
		}
		*p.oneOfType0 = v.(User)
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

func (p *OneOfUpdateUserApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateUserApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(User)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.User" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(User)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateUserApiResponseData"))
}

func (p *OneOfUpdateUserApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateUserApiResponseData")
}

type OneOfDeleteUserKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteUserKeyApiResponseData() *OneOfDeleteUserKeyApiResponseData {
	p := new(OneOfDeleteUserKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteUserKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteUserKeyApiResponseData is nil"))
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

func (p *OneOfDeleteUserKeyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteUserKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteUserKeyApiResponseData"))
}

func (p *OneOfDeleteUserKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteUserKeyApiResponseData")
}

type OneOfListCertAuthProvidersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []CertAuthProvider     `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListCertAuthProvidersApiResponseData() *OneOfListCertAuthProvidersApiResponseData {
	p := new(OneOfListCertAuthProvidersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCertAuthProvidersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCertAuthProvidersApiResponseData is nil"))
	}
	switch v.(type) {
	case []CertAuthProvider:
		p.oneOfType0 = v.([]CertAuthProvider)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.CertAuthProvider>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.CertAuthProvider>"
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

func (p *OneOfListCertAuthProvidersApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authn.CertAuthProvider>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListCertAuthProvidersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]CertAuthProvider)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.CertAuthProvider" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.CertAuthProvider>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.CertAuthProvider>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCertAuthProvidersApiResponseData"))
}

func (p *OneOfListCertAuthProvidersApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authn.CertAuthProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListCertAuthProvidersApiResponseData")
}

type OneOfConnectionDirectoryServiceApiResponseData struct {
	Discriminator *string                             `json:"-"`
	ObjectType_   *string                             `json:"-"`
	oneOfType0    *DirectoryServiceConnectionResponse `json:"-"`
	oneOfType400  *import1.ErrorResponse              `json:"-"`
}

func NewOneOfConnectionDirectoryServiceApiResponseData() *OneOfConnectionDirectoryServiceApiResponseData {
	p := new(OneOfConnectionDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfConnectionDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfConnectionDirectoryServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case DirectoryServiceConnectionResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryServiceConnectionResponse)
		}
		*p.oneOfType0 = v.(DirectoryServiceConnectionResponse)
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

func (p *OneOfConnectionDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfConnectionDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DirectoryServiceConnectionResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.DirectoryServiceConnectionResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryServiceConnectionResponse)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConnectionDirectoryServiceApiResponseData"))
}

func (p *OneOfConnectionDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfConnectionDirectoryServiceApiResponseData")
}

type OneOfSearchDirectoryServiceApiResponseData struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType0    *DirectoryServiceSearchResult `json:"-"`
	oneOfType400  *import1.ErrorResponse        `json:"-"`
}

func NewOneOfSearchDirectoryServiceApiResponseData() *OneOfSearchDirectoryServiceApiResponseData {
	p := new(OneOfSearchDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSearchDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSearchDirectoryServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case DirectoryServiceSearchResult:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(DirectoryServiceSearchResult)
		}
		*p.oneOfType0 = v.(DirectoryServiceSearchResult)
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

func (p *OneOfSearchDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfSearchDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(DirectoryServiceSearchResult)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.DirectoryServiceSearchResult" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(DirectoryServiceSearchResult)
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
		if "iam.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSearchDirectoryServiceApiResponseData"))
}

func (p *OneOfSearchDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfSearchDirectoryServiceApiResponseData")
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
