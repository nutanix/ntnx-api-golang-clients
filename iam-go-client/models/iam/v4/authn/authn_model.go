/*
 * Generated file models/iam/v4/authn/authn_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Iam Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId}/$actions/change-state Post operation
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

func NewActivateUserApiResponse() *ActivateUserApiResponse {
	p := new(ActivateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ActivateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ActivateUserApiResponse"}
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
Algorithm types for OIDC key
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
Credentials in the form of a unique random key for the service account
*/
type ApiKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The actual API key value, returned only during creation
	*/
	ApiKey *string `json:"apiKey,omitempty"`
	/*
	  User or service who created the API key
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the API key
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	CustomOptions *ApiKeyCustomOptions `json:"customOptions,omitempty"`
	/*
	  Brief description of the API key
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The expiry time of the API key
	*/
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the API key was last used
	*/
	LastUsedTime *time.Time `json:"lastUsedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier for the API key in the form of a name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  An optional set of audiences for the API key, meant to restrict its access to APIs
	*/
	Scope []string `json:"scope,omitempty"`

	Status *ApiKeyStatusType `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewApiKey() *ApiKey {
	p := new(ApiKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ApiKey"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ApiKey"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Optional customizations for the API key
*/
type ApiKeyCustomOptions struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Custom set of characters for the API key
	*/
	Alphabet *string `json:"alphabet"`
	/*
	  Custom length of the API key
	*/
	Length *int `json:"length"`
}

func (p *ApiKeyCustomOptions) MarshalJSON() ([]byte, error) {
	type ApiKeyCustomOptionsProxy ApiKeyCustomOptions
	return json.Marshal(struct {
		*ApiKeyCustomOptionsProxy
		Alphabet *string `json:"alphabet,omitempty"`
		Length   *int    `json:"length,omitempty"`
	}{
		ApiKeyCustomOptionsProxy: (*ApiKeyCustomOptionsProxy)(p),
		Alphabet:                 p.Alphabet,
		Length:                   p.Length,
	})
}

func NewApiKeyCustomOptions() *ApiKeyCustomOptions {
	p := new(ApiKeyCustomOptions)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ApiKeyCustomOptions"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ApiKeyCustomOptions"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The revocation status of the API key
*/
type ApiKeyStatusType int

const (
	APIKEYSTATUSTYPE_UNKNOWN  ApiKeyStatusType = 0
	APIKEYSTATUSTYPE_REDACTED ApiKeyStatusType = 1
	APIKEYSTATUSTYPE_VALID    ApiKeyStatusType = 2
	APIKEYSTATUSTYPE_REVOKED  ApiKeyStatusType = 3
	APIKEYSTATUSTYPE_EXPIRED  ApiKeyStatusType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApiKeyStatusType) name(index int) string {
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
func (e ApiKeyStatusType) GetName() string {
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
func (e *ApiKeyStatusType) index(name string) ApiKeyStatusType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"REVOKED",
		"EXPIRED",
	}
	for idx := range names {
		if names[idx] == name {
			return ApiKeyStatusType(idx)
		}
	}
	return APIKEYSTATUSTYPE_UNKNOWN
}

func (e *ApiKeyStatusType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApiKeyStatusType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApiKeyStatusType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApiKeyStatusType) Ref() *ApiKeyStatusType {
	return &e
}

/*
Body of the validate request, consists of an audience
*/
type ApiKeyValidateRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Audience *string `json:"audience,omitempty"`
}

func NewApiKeyValidateRequest() *ApiKeyValidateRequest {
	p := new(ApiKeyValidateRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ApiKeyValidateRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ApiKeyValidateRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response of the validate request, consists of the API key details and the service account ID
*/
type ApiKeyValidateResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The actual API key value, returned only during creation
	*/
	ApiKey *string `json:"apiKey,omitempty"`
	/*
	  User or service who created the API key
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the API key
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	CustomOptions *ApiKeyCustomOptions `json:"customOptions,omitempty"`
	/*
	  Brief description of the API key
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The expiry time of the API key
	*/
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the API key was last used
	*/
	LastUsedTime *time.Time `json:"lastUsedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier for the API key in the form of a name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  An optional set of audiences for the API key, meant to restrict its access to APIs
	*/
	Scope []string `json:"scope,omitempty"`

	ServiceAccountId *string `json:"serviceAccountId,omitempty"`

	Status *ApiKeyStatusType `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewApiKeyValidateResponse() *ApiKeyValidateResponse {
	p := new(ApiKeyValidateResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ApiKeyValidateResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ApiKeyValidateResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Supported Auth methods
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
Information of Bucket access key
*/
type BucketsAccessKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the bucket access key
	*/
	AccessKeyName *string `json:"accessKeyName"`
	/*
	  Creation time for the bucket access key
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Secret access key, it will be returned only during bucket access key creation
	*/
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  User Identifier who owns the bucket access key
	*/
	UserId *string `json:"userId"`
}

func (p *BucketsAccessKey) MarshalJSON() ([]byte, error) {
	type BucketsAccessKeyProxy BucketsAccessKey
	return json.Marshal(struct {
		*BucketsAccessKeyProxy
		AccessKeyName *string `json:"accessKeyName,omitempty"`
		UserId        *string `json:"userId,omitempty"`
	}{
		BucketsAccessKeyProxy: (*BucketsAccessKeyProxy)(p),
		AccessKeyName:         p.AccessKeyName,
		UserId:                p.UserId,
	})
}

func NewBucketsAccessKey() *BucketsAccessKey {
	p := new(BucketsAccessKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.BucketsAccessKey"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.BucketsAccessKey"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A certificate based authentication provider
*/
type CertAuthProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the uploaded CA chain file
	*/
	CaCertFileName *string `json:"caCertFileName,omitempty"`

	CertRevocationInfo *CertRevocationInfo `json:"certRevocationInfo,omitempty"`
	/*
	  Ca chain file
	*/
	ClientCaChain *string `json:"clientCaChain,omitempty"`
	/*
	  User or service who created the CertAuthProvider
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the CertAuthProvider
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  UUID of an existing directory service
	*/
	DirSvcExtID *string `json:"dirSvcExtID,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Flag to enable/disable Cert Auth for the current certificate based authentication provider.
	*/
	IsCacEnabled *bool `json:"isCacEnabled,omitempty"`
	/*
	  Flag to enable/disable CAC for the current certificate based authentication provider.
	*/
	IsCertAuthEnabled *bool `json:"isCertAuthEnabled,omitempty"`
	/*
	  Last updated time of the CertAuthProvider
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the certificate based authentication provider.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewCertAuthProvider() *CertAuthProvider {
	p := new(CertAuthProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertAuthProvider"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CertAuthProvider"}
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
	  List of the CRL distribution points which will be used to fetch the CRLs
	*/
	CrlDps []string `json:"crlDps,omitempty"`
	/*
	  Interval in seconds at which the CRL should be fetched from the CRLDP, default = 86400 seconds(1 day)
	*/
	GlobalCrlRefreshInterval *int `json:"globalCrlRefreshInterval,omitempty"`
	/*
	  Flag to enable/disable CRL revocation check
	*/
	IsCrlEnabled *bool `json:"isCrlEnabled,omitempty"`
	/*
	  Flag to enable/disable OCSP revocation check
	*/
	IsOcspEnabled *bool `json:"isOcspEnabled,omitempty"`
	/*
	  URL of the OCSP responder used to override the url from AIA extension
	*/
	OcspResponder *string `json:"ocspResponder,omitempty"`
}

func NewCertRevocationInfo() *CertRevocationInfo {
	p := new(CertRevocationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertRevocationInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CertRevocationInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authn/users/$actions/change-password Post operation
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

func NewChangeUserPasswordApiResponse() *ChangeUserPasswordApiResponse {
	p := new(ChangeUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ChangeUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ChangeUserPasswordApiResponse"}
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
Supported claims types
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services/{extId}/$actions/verify-connection-status Post operation
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

func NewConnectionDirectoryServiceApiResponse() *ConnectionDirectoryServiceApiResponse {
	p := new(ConnectionDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ConnectionDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ConnectionDirectoryServiceApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId}/api-keys Post operation
*/
type CreateApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateApiKeyApiResponse() *CreateApiKeyApiResponse {
	p := new(CreateApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateApiKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/cert-auth-providers Post operation
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

func NewCreateCertAuthProviderApiResponse() *CreateCertAuthProviderApiResponse {
	p := new(CreateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateCertAuthProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services Post operation
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

func NewCreateDirectoryServiceApiResponse() *CreateDirectoryServiceApiResponse {
	p := new(CreateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateDirectoryServiceApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/federation Post operation
*/
type CreateFederationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateFederationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateFederationApiResponse() *CreateFederationApiResponse {
	p := new(CreateFederationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateFederationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateFederationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateFederationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateFederationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateFederationApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-identity-providers Post operation
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

func NewCreateSamlIdentityProviderApiResponse() *CreateSamlIdentityProviderApiResponse {
	p := new(CreateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateSamlIdentityProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts Post operation
*/
type CreateServiceAccountApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateServiceAccountApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateServiceAccountApiResponse() *CreateServiceAccountApiResponse {
	p := new(CreateServiceAccountApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateServiceAccountApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateServiceAccountApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateServiceAccountApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateServiceAccountApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateServiceAccountApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId}/access-keys Post operation
*/
type CreateServiceAccountKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateServiceAccountKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateServiceAccountKeyApiResponse() *CreateServiceAccountKeyApiResponse {
	p := new(CreateServiceAccountKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateServiceAccountKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateServiceAccountKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateServiceAccountKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateServiceAccountKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateServiceAccountKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users Post operation
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

func NewCreateUserApiResponse() *CreateUserApiResponse {
	p := new(CreateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateUserApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/user-groups Post operation
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

func NewCreateUserGroupApiResponse() *CreateUserGroupApiResponse {
	p := new(CreateUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateUserGroupApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId}/buckets-access-keys Post operation
*/
type CreateUserKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUserKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUserKeyApiResponse() *CreateUserKeyApiResponse {
	p := new(CreateUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.CreateUserKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateUserKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateUserKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateUserKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{svcAccExtId}/api-keys/{extId} Delete operation
*/
type DeleteApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteApiKeyApiResponse() *DeleteApiKeyApiResponse {
	p := new(DeleteApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteApiKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/cert-auth-providers/{extId} Delete operation
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

func NewDeleteCertAuthProviderApiResponse() *DeleteCertAuthProviderApiResponse {
	p := new(DeleteCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteCertAuthProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services/{extId} Delete operation
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

func NewDeleteDirectoryServiceApiResponse() *DeleteDirectoryServiceApiResponse {
	p := new(DeleteDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteDirectoryServiceApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/federation Delete operation
*/
type DeleteFederationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteFederationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteFederationApiResponse() *DeleteFederationApiResponse {
	p := new(DeleteFederationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteFederationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteFederationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteFederationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteFederationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteFederationApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-identity-providers/{extId} Delete operation
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

func NewDeleteSamlIdentityProviderApiResponse() *DeleteSamlIdentityProviderApiResponse {
	p := new(DeleteSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteSamlIdentityProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId} Delete operation
*/
type DeleteServiceAccountApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteServiceAccountApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteServiceAccountApiResponse() *DeleteServiceAccountApiResponse {
	p := new(DeleteServiceAccountApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteServiceAccountApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteServiceAccountApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteServiceAccountApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteServiceAccountApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteServiceAccountApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{svcAccExtId}/access-keys/{extId} Delete operation
*/
type DeleteServiceAccountKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteServiceAccountKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteServiceAccountKeyApiResponse() *DeleteServiceAccountKeyApiResponse {
	p := new(DeleteServiceAccountKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteServiceAccountKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteServiceAccountKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteServiceAccountKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteServiceAccountKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteServiceAccountKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId} Delete operation
*/
type DeleteUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteUserApiResponse() *DeleteUserApiResponse {
	p := new(DeleteUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteUserApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteUserApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/user-groups/{extId} Delete operation
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

func NewDeleteUserGroupApiResponse() *DeleteUserGroupApiResponse {
	p := new(DeleteUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteUserGroupApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{userExtId}/buckets-access-keys/{extId} Delete operation
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

func NewDeleteUserKeyApiResponse() *DeleteUserKeyApiResponse {
	p := new(DeleteUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DeleteUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DeleteUserKeyApiResponse"}
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
Information of a directory service
*/
type DirectoryService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the directory service
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the directory service
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DirectoryType *DirectoryType `json:"directoryType,omitempty"`
	/*
	  Domain name for the directory service
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupSearchType *GroupSearchType `json:"groupSearchType,omitempty"`
	/*
	  Last updated time of the directory service
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name for the directory service
	*/
	Name *string `json:"name,omitempty"`

	OpenLdapConfiguration *OpenLdapConfig `json:"openLdapConfiguration,omitempty"`
	/*
	  Secondary Url for the directory service
	*/
	SecondaryUrls []string `json:"secondaryUrls,omitempty"`

	ServiceAccount *DsServiceAccount `json:"serviceAccount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Url for the directory service
	*/
	Url *string `json:"url,omitempty"`
	/*
	  List of allowed user groups for the directory service
	*/
	WhiteListedGroups []string `json:"whiteListedGroups,omitempty"`
}

func NewDirectoryService() *DirectoryService {
	p := new(DirectoryService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryService"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryService"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for directory service connection request
*/
type DirectoryServiceConnectionRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password to connect to the directory service
	*/
	Password *string `json:"password"`
	/*
	  Username to connect to the directory service
	*/
	Username *string `json:"username"`
}

func (p *DirectoryServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	type DirectoryServiceConnectionRequestProxy DirectoryServiceConnectionRequest
	return json.Marshal(struct {
		*DirectoryServiceConnectionRequestProxy
		Password *string `json:"password,omitempty"`
		Username *string `json:"username,omitempty"`
	}{
		DirectoryServiceConnectionRequestProxy: (*DirectoryServiceConnectionRequestProxy)(p),
		Password:                               p.Password,
		Username:                               p.Username,
	})
}

func NewDirectoryServiceConnectionRequest() *DirectoryServiceConnectionRequest {
	p := new(DirectoryServiceConnectionRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceConnectionRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceConnectionRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
AD/LDAP information of the user
*/
type DirectoryServiceInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of AD/LDAP groups having the user
	*/
	Groups []DirectoryServiceInfoGroup `json:"groups"`
	/*
	  List of AD/LDAP OUs having the given user
	*/
	Ous []DirectoryServiceInfoOu `json:"ous"`
	/*
	  External Identifier of the user
	*/
	UserId *string `json:"userId"`
}

func (p *DirectoryServiceInfo) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoProxy DirectoryServiceInfo
	return json.Marshal(struct {
		*DirectoryServiceInfoProxy
		Groups []DirectoryServiceInfoGroup `json:"groups,omitempty"`
		Ous    []DirectoryServiceInfoOu    `json:"ous,omitempty"`
		UserId *string                     `json:"userId,omitempty"`
	}{
		DirectoryServiceInfoProxy: (*DirectoryServiceInfoProxy)(p),
		Groups:                    p.Groups,
		Ous:                       p.Ous,
		UserId:                    p.UserId,
	})
}

func NewDirectoryServiceInfo() *DirectoryServiceInfo {
	p := new(DirectoryServiceInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of AD group having the user
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
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the group
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DirectoryServiceInfoGroup) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoGroupProxy DirectoryServiceInfoGroup
	return json.Marshal(struct {
		*DirectoryServiceInfoGroupProxy
		Name *string `json:"name,omitempty"`
	}{
		DirectoryServiceInfoGroupProxy: (*DirectoryServiceInfoGroupProxy)(p),
		Name:                           p.Name,
	})
}

func NewDirectoryServiceInfoGroup() *DirectoryServiceInfoGroup {
	p := new(DirectoryServiceInfoGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfoGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceInfoGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of OUs having the user
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
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the OU
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DirectoryServiceInfoOu) MarshalJSON() ([]byte, error) {
	type DirectoryServiceInfoOuProxy DirectoryServiceInfoOu
	return json.Marshal(struct {
		*DirectoryServiceInfoOuProxy
		Name *string `json:"name,omitempty"`
	}{
		DirectoryServiceInfoOuProxy: (*DirectoryServiceInfoOuProxy)(p),
		Name:                        p.Name,
	})
}

func NewDirectoryServiceInfoOu() *DirectoryServiceInfoOu {
	p := new(DirectoryServiceInfoOu)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceInfoOu"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceInfoOu"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of searched attributes
*/
type DirectoryServiceSearchAttribute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the attribute
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Value of the attribute
	*/
	Values []string `json:"values,omitempty"`
}

func NewDirectoryServiceSearchAttribute() *DirectoryServiceSearchAttribute {
	p := new(DirectoryServiceSearchAttribute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchAttribute"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceSearchAttribute"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of single search entity
*/
type DirectoryServiceSearchEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Attributes []DirectoryServiceSearchAttribute `json:"attributes,omitempty"`
	/*
	  Type of entity either user or group
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  Name of the entity in canonical format
	*/
	Name *string `json:"name,omitempty"`
}

func NewDirectoryServiceSearchEntity() *DirectoryServiceSearchEntity {
	p := new(DirectoryServiceSearchEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchEntity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceSearchEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for directory search query
*/
type DirectoryServiceSearchQuery struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag indicating whether the search should be a wildcard search or not
	*/
	IsWildcardSearch *bool `json:"isWildcardSearch,omitempty"`
	/*
	  Query string for directory service search
	*/
	Query *string `json:"query"`
	/*
	  Attributes the search operation returns
	*/
	ReturnedAttributes []string `json:"returnedAttributes,omitempty"`
	/*
	  Attributes for search operation. By default search will be performed with common name
	*/
	SearchedAttributes []string `json:"searchedAttributes,omitempty"`
}

func (p *DirectoryServiceSearchQuery) MarshalJSON() ([]byte, error) {
	type DirectoryServiceSearchQueryProxy DirectoryServiceSearchQuery
	return json.Marshal(struct {
		*DirectoryServiceSearchQueryProxy
		Query *string `json:"query,omitempty"`
	}{
		DirectoryServiceSearchQueryProxy: (*DirectoryServiceSearchQueryProxy)(p),
		Query:                            p.Query,
	})
}

func NewDirectoryServiceSearchQuery() *DirectoryServiceSearchQuery {
	p := new(DirectoryServiceSearchQuery)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchQuery"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceSearchQuery"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsWildcardSearch = new(bool)
	*p.IsWildcardSearch = true

	return p
}

/*
Information of directory service search result
*/
type DirectoryServiceSearchResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Domain name for the directory service
	*/
	DomainName *string `json:"domainName,omitempty"`
	/*
	  Result of directory service search
	*/
	SearchResults []DirectoryServiceSearchEntity `json:"searchResults,omitempty"`
}

func NewDirectoryServiceSearchResult() *DirectoryServiceSearchResult {
	p := new(DirectoryServiceSearchResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchResult"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DirectoryServiceSearchResult"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of directory service
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
Supported subject types
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
Information of service account to connect to the directory service
*/
type DsServiceAccount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Password to connect to the directory service
	*/
	Password *string `json:"password"`
	/*
	  Username to connect to the directory service
	*/
	Username *string `json:"username"`
}

func (p *DsServiceAccount) MarshalJSON() ([]byte, error) {
	type DsServiceAccountProxy DsServiceAccount
	return json.Marshal(struct {
		*DsServiceAccountProxy
		Password *string `json:"password,omitempty"`
		Username *string `json:"username,omitempty"`
	}{
		DsServiceAccountProxy: (*DsServiceAccountProxy)(p),
		Password:              p.Password,
		Username:              p.Username,
	})
}

func NewDsServiceAccount() *DsServiceAccount {
	p := new(DsServiceAccount)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DsServiceAccount"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.DsServiceAccount"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Description of the OIDC provider
*/
type Federation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClaimMap *FederationClaims `json:"claimMap,omitempty"`
	/*
	  API key for authenticating to the OIDC provider
	*/
	CloudApiKey *string `json:"cloudApiKey,omitempty"`
	/*
	  Provider side tenant that is registering domain
	*/
	CloudTenant *string `json:"cloudTenant,omitempty"`
	/*
	  User or service who created the OIDC provider
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the OIDC provider
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  IDP attached to users of the OIDC provider
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Last updated time of the OIDC provider
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Well known URL of the OIDC provider
	*/
	WellKnownConfig *string `json:"wellKnownConfig,omitempty"`
}

func NewFederation() *Federation {
	p := new(Federation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Federation"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.Federation"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Mapping of claims from IAM to the OIDC provider
*/
type FederationClaims struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Email *string `json:"email,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewFederationClaims() *FederationClaims {
	p := new(FederationClaims)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.FederationClaims"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.FederationClaims"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{svcAccExtId}/api-keys/{extId} Get operation
*/
type GetApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetApiKeyApiResponse() *GetApiKeyApiResponse {
	p := new(GetApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetApiKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/cert-auth-providers/{extId} Get operation
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

func NewGetCertAuthProviderApiResponse() *GetCertAuthProviderApiResponse {
	p := new(GetCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetCertAuthProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services/{extId} Get operation
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

func NewGetDirectoryServiceApiResponse() *GetDirectoryServiceApiResponse {
	p := new(GetDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetDirectoryServiceApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/federation Get operation
*/
type GetFederationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetFederationApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetFederationApiResponse() *GetFederationApiResponse {
	p := new(GetFederationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetFederationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetFederationApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetFederationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetFederationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetFederationApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-identity-providers/{extId} Get operation
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

func NewGetSamlIdentityProviderApiResponse() *GetSamlIdentityProviderApiResponse {
	p := new(GetSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetSamlIdentityProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-sp-metadata Get operation
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

func NewGetSamlSpMetadataApiResponse() *GetSamlSpMetadataApiResponse {
	p := new(GetSamlSpMetadataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlSpMetadataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetSamlSpMetadataApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/buckets-access-keys/{accessKeyId}/secret-key Get operation
*/
type GetSecretKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSecretKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSecretKeyApiResponse() *GetSecretKeyApiResponse {
	p := new(GetSecretKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSecretKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetSecretKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSecretKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSecretKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSecretKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId} Get operation
*/
type GetServiceAccountApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetServiceAccountApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetServiceAccountApiResponse() *GetServiceAccountApiResponse {
	p := new(GetServiceAccountApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetServiceAccountApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetServiceAccountApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetServiceAccountApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetServiceAccountApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetServiceAccountApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{svcAccExtId}/access-keys/{extId} Get operation
*/
type GetServiceAccountKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetServiceAccountKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetServiceAccountKeyApiResponse() *GetServiceAccountKeyApiResponse {
	p := new(GetServiceAccountKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetServiceAccountKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetServiceAccountKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetServiceAccountKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetServiceAccountKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetServiceAccountKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId} Get operation
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

func NewGetUserApiResponse() *GetUserApiResponse {
	p := new(GetUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetUserApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/user-groups/{extId} Get operation
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

func NewGetUserGroupApiResponse() *GetUserGroupApiResponse {
	p := new(GetUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetUserGroupApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/oidc/userinfo Get operation
*/
type GetUserInfoApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserInfoApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUserInfoApiResponse() *GetUserInfoApiResponse {
	p := new(GetUserInfoApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserInfoApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.GetUserInfoApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetUserInfoApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetUserInfoApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetUserInfoApiResponseData()
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
Type of grant
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
Information of a group of users
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
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The name of group.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewGroup() *Group {
	p := new(Group)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Group"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.Group"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Group membership search type for the directory service
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
Type of the user group
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
Information of the IDP
*/
type IdpMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Certificate for verification
	*/
	Certificate *string `json:"certificate"`
	/*
	  Entity Id of Identity provider
	*/
	EntityId *string `json:"entityId"`
	/*
	  Error URL of the Identity provider
	*/
	ErrorUrl *string `json:"errorUrl,omitempty"`
	/*
	  Login URL of the Identity provider
	*/
	LoginUrl *string `json:"loginUrl"`
	/*
	  Logout URL of the Identity provider
	*/
	LogoutUrl *string `json:"logoutUrl,omitempty"`

	NameIdPolicyFormat *NameIdPolicyFormat `json:"nameIdPolicyFormat,omitempty"`
}

func (p *IdpMetadata) MarshalJSON() ([]byte, error) {
	type IdpMetadataProxy IdpMetadata
	return json.Marshal(struct {
		*IdpMetadataProxy
		Certificate *string `json:"certificate,omitempty"`
		EntityId    *string `json:"entityId,omitempty"`
		LoginUrl    *string `json:"loginUrl,omitempty"`
	}{
		IdpMetadataProxy: (*IdpMetadataProxy)(p),
		Certificate:      p.Certificate,
		EntityId:         p.EntityId,
		LoginUrl:         p.LoginUrl,
	})
}

func NewIdpMetadata() *IdpMetadata {
	p := new(IdpMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.IdpMetadata"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.IdpMetadata"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of OIDC key
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId}/api-keys Get operation
*/
type ListApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListApiKeyApiResponse() *ListApiKeyApiResponse {
	p := new(ListApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListApiKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/cert-auth-providers Get operation
*/
type ListCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCertAuthProviderApiResponse() *ListCertAuthProviderApiResponse {
	p := new(ListCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListCertAuthProviderApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCertAuthProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCertAuthProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCertAuthProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services Get operation
*/
type ListDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDirectoryServiceApiResponse() *ListDirectoryServiceApiResponse {
	p := new(ListDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListDirectoryServiceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListDirectoryServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListDirectoryServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListDirectoryServiceApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-identity-providers Get operation
*/
type ListSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSamlIdentityProviderApiResponse() *ListSamlIdentityProviderApiResponse {
	p := new(ListSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListSamlIdentityProviderApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListSamlIdentityProviderApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListSamlIdentityProviderApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListSamlIdentityProviderApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts Get operation
*/
type ListServiceAccountApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListServiceAccountApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListServiceAccountApiResponse() *ListServiceAccountApiResponse {
	p := new(ListServiceAccountApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListServiceAccountApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListServiceAccountApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListServiceAccountApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListServiceAccountApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListServiceAccountApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{extId}/access-keys Get operation
*/
type ListServiceAccountKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListServiceAccountKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListServiceAccountKeyApiResponse() *ListServiceAccountKeyApiResponse {
	p := new(ListServiceAccountKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListServiceAccountKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListServiceAccountKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListServiceAccountKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListServiceAccountKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListServiceAccountKeyApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users Get operation
*/
type ListUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUserApiResponse() *ListUserApiResponse {
	p := new(ListUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListUserApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUserApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUserApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUserApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/user-groups Get operation
*/
type ListUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUserGroupApiResponse() *ListUserGroupApiResponse {
	p := new(ListUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListUserGroupApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUserGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUserGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUserGroupApiResponseData()
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId}/buckets-access-keys Get operation
*/
type ListUserKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUserKeyApiResponse() *ListUserKeyApiResponse {
	p := new(ListUserKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ListUserKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListUserKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListUserKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListUserKeyApiResponseData()
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
Name ID Policy format
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
Information for OIDC client
*/
type OidcClient struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Application type for OIDC client
	*/
	ApplicationType *string `json:"application_type,omitempty"`
	/*
	  Unique client identifier generated by the server
	*/
	ClientId *string `json:"client_id,omitempty"`
	/*
	  Name of OIDC client
	*/
	ClientName *string `json:"client_name"`
	/*
	  Credential generated by the server for authentication of the OIDC client
	*/
	ClientSecret *string `json:"client_secret,omitempty"`
	/*
	  Creation time of the OIDC client
	*/
	CreatedTime *time.Time `json:"created_time,omitempty"`
	/*
	  Last Update time of the OIDC client
	*/
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	/*
	  Redirect URIs for OIDC client
	*/
	RedirectUris []string `json:"redirect_uris"`
	/*
	  SHA256 of a certificate public key
	*/
	Sha256ThumbprintList []string `json:"sha256_thumbprint_list,omitempty"`
}

func (p *OidcClient) MarshalJSON() ([]byte, error) {
	type OidcClientProxy OidcClient
	return json.Marshal(struct {
		*OidcClientProxy
		ClientName   *string  `json:"client_name,omitempty"`
		RedirectUris []string `json:"redirect_uris,omitempty"`
	}{
		OidcClientProxy: (*OidcClientProxy)(p),
		ClientName:      p.ClientName,
		RedirectUris:    p.RedirectUris,
	})
}

func NewOidcClient() *OidcClient {
	p := new(OidcClient)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcClient"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OidcClient"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
JSON Web Key that can be used to verify the signature of tokens issued by IAM
*/
type OidcKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Alg *AlgoType `json:"alg"`
	/*
	  Exponent for a standard pem
	*/
	E *string `json:"e"`
	/*
	  Unique identifier for the key
	*/
	Kid *string `json:"kid"`

	Kty *KeyType `json:"kty"`
	/*
	  Modulus for a standard pem
	*/
	N *string `json:"n"`

	Use *UseType `json:"use"`
	/*
	  x509 certificate chain for the key
	*/
	X5c *string `json:"x5c"`
}

func (p *OidcKey) MarshalJSON() ([]byte, error) {
	type OidcKeyProxy OidcKey
	return json.Marshal(struct {
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
	})
}

func NewOidcKey() *OidcKey {
	p := new(OidcKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcKey"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OidcKey"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of IAM OpenID keys
*/
type OidcKeys struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of JSON Web Keys that can be used to verify the signature of tokens issued by IAM
	*/
	Keys []OidcKey `json:"keys"`
}

func (p *OidcKeys) MarshalJSON() ([]byte, error) {
	type OidcKeysProxy OidcKeys
	return json.Marshal(struct {
		*OidcKeysProxy
		Keys []OidcKey `json:"keys,omitempty"`
	}{
		OidcKeysProxy: (*OidcKeysProxy)(p),
		Keys:          p.Keys,
	})
}

func NewOidcKeys() *OidcKeys {
	p := new(OidcKeys)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcKeys"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OidcKeys"}
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
IAM OpenID userinfo response
*/
type OidcUserinfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID V5 of the authentication connector that was used to verify the credentials of the user
	*/
	ConnectorId *string `json:"connectorId,omitempty"`
	/*
	  Any additional user attributes obtained from the authentication provider in the form of key/value pairs.
	*/
	CustomClaims map[string]string `json:"customClaims,omitempty"`
	/*
	  Email address of the subject
	*/
	Email *string `json:"email,omitempty"`
	/*
	  Expiration time for the identity claims in seconds since epoch
	*/
	Exp *int64 `json:"exp,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of groups obtained from the authentication provider for the user
	*/
	Groups []Group `json:"groups,omitempty"`
	/*
	  Time at which the identity claims were issued in seconds since epoch
	*/
	Iat *int64 `json:"iat,omitempty"`
	/*
	  Boolean to indicate if the email address of end user has been verified by IAM
	*/
	IsEmailVerified *bool `json:"isEmailVerified,omitempty"`
	/*
	  Base URL of the identity service which it asserts as its issuer identifier
	*/
	Iss *string `json:"iss,omitempty"`
	/*
	  List of legacy Prism roles for the user
	*/
	LegacyRoles []string `json:"legacyRoles,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Display name of the user including the first and last name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Subject Identifier. A locally unique identifier for the end user
	*/
	Sub *string `json:"sub,omitempty"`

	SubType *OidcSubjectType `json:"subType,omitempty"`

	Tenant *Tenant `json:"tenant,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  UUID V5 created for the user by IAM
	*/
	UserId *string `json:"userId,omitempty"`
}

func NewOidcUserinfo() *OidcUserinfo {
	p := new(OidcUserinfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcUserinfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OidcUserinfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of OIDC well-known config
*/
type OidcWellKnownConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  URL of the OIDC authorization endpoint
	*/
	AuthorizationEndpoint *string `json:"authorization_endpoint"`
	/*
	  List of supported claims
	*/
	ClaimsSupported []ClaimsType `json:"claims_supported"`
	/*
	  List of the JWS signing algorithms supported for the ID Token
	*/
	IdTokenSigningAlgValuesSupported []SigningAlgoType `json:"id_token_signing_alg_values_supported"`
	/*
	  Base URL of the identity service which it asserts as its issuer identifier
	*/
	Issuer *string `json:"issuer"`
	/*
	  URL of JSON Web Key Set document of the identity service
	*/
	JwksUri *string `json:"jwks_uri"`
	/*
	  URL of the dynamic client registration endpoint
	*/
	RegistrationEndpoint *string `json:"registration_endpoint"`
	/*
	  List of supported OIDC response type values
	*/
	ResponseTypesSupported []ResponseType `json:"response_types_supported"`
	/*
	  List of supported OIDC scope values
	*/
	ScopesSupported []ScopesType `json:"scopes_supported"`
	/*
	  List of supported subject types
	*/
	SubjectTypesSupported []DiscoverySubjectType `json:"subject_types_supported"`
	/*
	  URL of the OIDC token endpoint
	*/
	TokenEndpoint *string `json:"token_endpoint"`
	/*
	  List of client authentication methods supported by the token endpoint
	*/
	TokenEndpointAuthMethodsSupported []AuthMethodType `json:"token_endpoint_auth_methods_supported"`
}

func (p *OidcWellKnownConfig) MarshalJSON() ([]byte, error) {
	type OidcWellKnownConfigProxy OidcWellKnownConfig
	return json.Marshal(struct {
		*OidcWellKnownConfigProxy
		AuthorizationEndpoint             *string                `json:"authorization_endpoint,omitempty"`
		ClaimsSupported                   []ClaimsType           `json:"claims_supported,omitempty"`
		IdTokenSigningAlgValuesSupported  []SigningAlgoType      `json:"id_token_signing_alg_values_supported,omitempty"`
		Issuer                            *string                `json:"issuer,omitempty"`
		JwksUri                           *string                `json:"jwks_uri,omitempty"`
		RegistrationEndpoint              *string                `json:"registration_endpoint,omitempty"`
		ResponseTypesSupported            []ResponseType         `json:"response_types_supported,omitempty"`
		ScopesSupported                   []ScopesType           `json:"scopes_supported,omitempty"`
		SubjectTypesSupported             []DiscoverySubjectType `json:"subject_types_supported,omitempty"`
		TokenEndpoint                     *string                `json:"token_endpoint,omitempty"`
		TokenEndpointAuthMethodsSupported []AuthMethodType       `json:"token_endpoint_auth_methods_supported,omitempty"`
	}{
		OidcWellKnownConfigProxy:          (*OidcWellKnownConfigProxy)(p),
		AuthorizationEndpoint:             p.AuthorizationEndpoint,
		ClaimsSupported:                   p.ClaimsSupported,
		IdTokenSigningAlgValuesSupported:  p.IdTokenSigningAlgValuesSupported,
		Issuer:                            p.Issuer,
		JwksUri:                           p.JwksUri,
		RegistrationEndpoint:              p.RegistrationEndpoint,
		ResponseTypesSupported:            p.ResponseTypesSupported,
		ScopesSupported:                   p.ScopesSupported,
		SubjectTypesSupported:             p.SubjectTypesSupported,
		TokenEndpoint:                     p.TokenEndpoint,
		TokenEndpointAuthMethodsSupported: p.TokenEndpointAuthMethodsSupported,
	})
}

func NewOidcWellKnownConfig() *OidcWellKnownConfig {
	p := new(OidcWellKnownConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OidcWellKnownConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OidcWellKnownConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration for OpenLDAP directory service
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
	return json.Marshal(struct {
		*OpenLdapConfigProxy
		UserConfiguration      *UserConfiguration      `json:"userConfiguration,omitempty"`
		UserGroupConfiguration *UserGroupConfiguration `json:"userGroupConfiguration,omitempty"`
	}{
		OpenLdapConfigProxy:    (*OpenLdapConfigProxy)(p),
		UserConfiguration:      p.UserConfiguration,
		UserGroupConfiguration: p.UserGroupConfiguration,
	})
}

func NewOpenLdapConfig() *OpenLdapConfig {
	p := new(OpenLdapConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.OpenLdapConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.OpenLdapConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for password change of user
*/
type PasswordChangeRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  New password for the user
	*/
	NewPassword *string `json:"newPassword"`
	/*
	  Current password of the user
	*/
	OldPassword *string `json:"oldPassword"`
	/*
	  Identifier for the user in the form an email address
	*/
	Username *string `json:"username"`
}

func (p *PasswordChangeRequest) MarshalJSON() ([]byte, error) {
	type PasswordChangeRequestProxy PasswordChangeRequest
	return json.Marshal(struct {
		*PasswordChangeRequestProxy
		NewPassword *string `json:"newPassword,omitempty"`
		OldPassword *string `json:"oldPassword,omitempty"`
		Username    *string `json:"username,omitempty"`
	}{
		PasswordChangeRequestProxy: (*PasswordChangeRequestProxy)(p),
		NewPassword:                p.NewPassword,
		OldPassword:                p.OldPassword,
		Username:                   p.Username,
	})
}

func NewPasswordChangeRequest() *PasswordChangeRequest {
	p := new(PasswordChangeRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordChangeRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.PasswordChangeRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information for password reset of user
*/
type PasswordResetRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  New password for the user
	*/
	NewPassword *string `json:"newPassword"`
}

func (p *PasswordResetRequest) MarshalJSON() ([]byte, error) {
	type PasswordResetRequestProxy PasswordResetRequest
	return json.Marshal(struct {
		*PasswordResetRequestProxy
		NewPassword *string `json:"newPassword,omitempty"`
	}{
		PasswordResetRequestProxy: (*PasswordResetRequestProxy)(p),
		NewPassword:               p.NewPassword,
	})
}

func NewPasswordResetRequest() *PasswordResetRequest {
	p := new(PasswordResetRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.PasswordResetRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.PasswordResetRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId}/$actions/reset-password Post operation
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

func NewResetUserPasswordApiResponse() *ResetUserPasswordApiResponse {
	p := new(ResetUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ResetUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ResetUserPasswordApiResponse"}
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
Supported OIDC response types
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
REST response for all response codes in API path /iam/v4.0.b1/authn/service-accounts/{svcAccExtId}/api-keys/{extId}/$actions/revoke Post operation
*/
type RevokeApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRevokeApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRevokeApiKeyApiResponse() *RevokeApiKeyApiResponse {
	p := new(RevokeApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.RevokeApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.RevokeApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RevokeApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RevokeApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRevokeApiKeyApiResponseData()
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
Information of SAML IDP
*/
type SamlIdentityProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the SAML Identity Provider
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the SAML Identity Provider
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  SAML assertions for list of custom attribute elements
	*/
	CustomAttributes []string `json:"customAttributes,omitempty"`
	/*
	  SAML assertion email attribute element
	*/
	EmailAttribute *string `json:"emailAttribute,omitempty"`
	/*
	  It will be used as Issuer in SAML authnRequest
	*/
	EntityIssuer *string `json:"entityIssuer,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  SAML assertion groups attribute element
	*/
	GroupsAttribute *string `json:"groupsAttribute,omitempty"`
	/*
	  Delimiter is used to split the value of attribute into multiple groups
	*/
	GroupsDelim *string `json:"groupsDelim,omitempty"`

	IdpMetadata *IdpMetadata `json:"idpMetadata,omitempty"`
	/*
	  Metadata url that provides IDP details
	*/
	IdpMetadataUrl *string `json:"idpMetadataUrl,omitempty"`
	/*
	  Base64 encoded metadata in xml format with IDP details
	*/
	IdpMetadataXml *string `json:"idpMetadataXml,omitempty"`
	/*
	  Flag indicating signing of SAML authnRequests
	*/
	IsSignedAuthnReqEnabled *bool `json:"isSignedAuthnReqEnabled,omitempty"`
	/*
	  Last updated time of the SAML Identity Provider
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Unique name of the IDP
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  SAML assertion username attribute element
	*/
	UsernameAttribute *string `json:"usernameAttribute,omitempty"`
}

func NewSamlIdentityProvider() *SamlIdentityProvider {
	p := new(SamlIdentityProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SamlIdentityProvider"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.SamlIdentityProvider"}
	p.UnknownFields_ = map[string]interface{}{}

	p.EmailAttribute = new(string)
	*p.EmailAttribute = "email"
	p.UsernameAttribute = new(string)
	*p.UsernameAttribute = "name"

	return p
}

/*
Type of scope
*/
type ScopeType int

const (
	SCOPETYPE_UNKNOWN  ScopeType = 0
	SCOPETYPE_REDACTED ScopeType = 1
	SCOPETYPE_OPEN_ID  ScopeType = 2
	SCOPETYPE_PROFILE  ScopeType = 3
	SCOPETYPE_EMAIL    ScopeType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScopeType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScopeType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScopeType) index(name string) ScopeType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
	}
	for idx := range names {
		if names[idx] == name {
			return ScopeType(idx)
		}
	}
	return SCOPETYPE_UNKNOWN
}

func (e *ScopeType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScopeType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScopeType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScopeType) Ref() *ScopeType {
	return &e
}

/*
Supported OIDC scopes
*/
type ScopesType int

const (
	SCOPESTYPE_UNKNOWN  ScopesType = 0
	SCOPESTYPE_REDACTED ScopesType = 1
	SCOPESTYPE_OPEN_ID  ScopesType = 2
	SCOPESTYPE_PROFILE  ScopesType = 3
	SCOPESTYPE_EMAIL    ScopesType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScopesType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
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
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
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
		"OPEN_ID",
		"PROFILE",
		"EMAIL",
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services/{extId}/$actions/search Post operation
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

func NewSearchDirectoryServiceApiResponse() *SearchDirectoryServiceApiResponse {
	p := new(SearchDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SearchDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.SearchDirectoryServiceApiResponse"}
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
Get secret key and user details for given access key id
*/
type SecretKeyRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External Identifier for Bucket Access Key
	*/
	AccessKeyId *string `json:"accessKeyId"`
}

func (p *SecretKeyRequest) MarshalJSON() ([]byte, error) {
	type SecretKeyRequestProxy SecretKeyRequest
	return json.Marshal(struct {
		*SecretKeyRequestProxy
		AccessKeyId *string `json:"accessKeyId,omitempty"`
	}{
		SecretKeyRequestProxy: (*SecretKeyRequestProxy)(p),
		AccessKeyId:           p.AccessKeyId,
	})
}

func NewSecretKeyRequest() *SecretKeyRequest {
	p := new(SecretKeyRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SecretKeyRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.SecretKeyRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
IAM secret key response
*/
type SecretKeyResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Email Id for the user
	*/
	EmailId *string `json:"emailId"`
	/*
	  Identifier of the IDP for the user
	*/
	IdpId *string `json:"idpId"`
	/*
	  Secret access key, it will be returned only during bucket access key creation
	*/
	SecretAccessKey *string `json:"secretAccessKey"`
	/*
	  ID of tenant for which configuration is being set up
	*/
	TenantId *string `json:"tenantId"`
	/*
	  User Identifier who owns the bucket access key
	*/
	UserId *string `json:"userId"`

	UserType *UserType `json:"userType"`
}

func (p *SecretKeyResponse) MarshalJSON() ([]byte, error) {
	type SecretKeyResponseProxy SecretKeyResponse
	return json.Marshal(struct {
		*SecretKeyResponseProxy
		EmailId         *string   `json:"emailId,omitempty"`
		IdpId           *string   `json:"idpId,omitempty"`
		SecretAccessKey *string   `json:"secretAccessKey,omitempty"`
		TenantId        *string   `json:"tenantId,omitempty"`
		UserId          *string   `json:"userId,omitempty"`
		UserType        *UserType `json:"userType,omitempty"`
	}{
		SecretKeyResponseProxy: (*SecretKeyResponseProxy)(p),
		EmailId:                p.EmailId,
		IdpId:                  p.IdpId,
		SecretAccessKey:        p.SecretAccessKey,
		TenantId:               p.TenantId,
		UserId:                 p.UserId,
		UserType:               p.UserType,
	})
}

func NewSecretKeyResponse() *SecretKeyResponse {
	p := new(SecretKeyResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SecretKeyResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.SecretKeyResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A special type of identity meant for non-browser API clients
*/
type ServiceAccount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the service account
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The creation time of the service account
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Brief description of the service account
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Email ID of the service account
	*/
	EmailId *string `json:"emailId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the service account was last updated
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier for the service account in the form of a name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Nutanix access keys for the service account
	*/
	NutanixAccessKeys []ServiceAccountKey `json:"nutanixAccessKeys,omitempty"`
	/*
	  Nutanix API keys for the service account
	*/
	NutanixApiKeys []ApiKey `json:"nutanixApiKeys,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewServiceAccount() *ServiceAccount {
	p := new(ServiceAccount)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ServiceAccount"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ServiceAccount"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Credentials in the form of a private key and certificate for the service account, for authentication over mTLS
*/
type ServiceAccountKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  X509 certificate signed by IAM in PEM format (includes leaf as well as ICAs), returned only during creation
	*/
	Certificate *string `json:"certificate,omitempty"`
	/*
	  Creation time for the nutanix access key
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  User or service who created the nutanix access key
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Brief description of the service account key
	*/
	Description *string `json:"description,omitempty"`
	/*
	  Expiry time of the certificate
	*/
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Identifier of the service account key in the form of a name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  RSA private key encoded in PEM format, returned only during creation
	*/
	PrivateKey *string `json:"privateKey,omitempty"`
	/*
	  Serial number of the certificate
	*/
	SerialNumber *string `json:"serialNumber,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewServiceAccountKey() *ServiceAccountKey {
	p := new(ServiceAccountKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ServiceAccountKey"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ServiceAccountKey"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Supported JWS signing algorithms
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
Type of subject for token
*/
type SubjectType int

const (
	SUBJECTTYPE_UNKNOWN                                   SubjectType = 0
	SUBJECTTYPE_REDACTED                                  SubjectType = 1
	SUBJECTTYPE_URN_IETF_PARAMS_OAUTH_TOKEN_TYPE_ID_TOKEN SubjectType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SubjectType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"URN:IETF:PARAMS:OAUTH:TOKEN-TYPE:ID_TOKEN",
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
Information regarding the tenant that the user belongs to
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
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the tenant that the user belongs to
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.Tenant"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.Tenant"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the OIDC Token request
*/
type TokenRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Audience for the OIDC token request
	*/
	Audience *string `json:"audience,omitempty"`
	/*
	  Client assertion for the OIDC token request
	*/
	ClientAssertion *string `json:"client_assertion,omitempty"`
	/*
	  Client assertion type for the OIDC token request
	*/
	ClientAssertionType *string `json:"client_assertion_type,omitempty"`
	/*
	  Client Id for the OIDC token request
	*/
	ClientId *string `json:"client_id"`
	/*
	  Code for the OIDC token
	*/
	Code *string `json:"code,omitempty"`
	/*
	  Connector Identifier for connector of user
	*/
	ConnectorId *string `json:"connector_id,omitempty"`

	GrantType *GrantType `json:"grant_type"`
	/*
	  Non Tenant Issuer for the OIDC token request
	*/
	NonTenantIssuer *string `json:"non_tenant_issuer,omitempty"`
	/*
	  Redirect Uri for the OIDC token request
	*/
	RedirectUri *string `json:"redirect_uri,omitempty"`
	/*
	  Request token to get the OIDC token
	*/
	RefreshToken *string `json:"refresh_token,omitempty"`

	Scope *ScopeType `json:"scope"`
	/*
	  Subject Token for the OIDC token request
	*/
	SubjectToken *string `json:"subject_token,omitempty"`

	SubjectTokenType *SubjectType `json:"subject_token_type,omitempty"`
}

func (p *TokenRequest) MarshalJSON() ([]byte, error) {
	type TokenRequestProxy TokenRequest
	return json.Marshal(struct {
		*TokenRequestProxy
		ClientId  *string    `json:"client_id,omitempty"`
		GrantType *GrantType `json:"grant_type,omitempty"`
		Scope     *ScopeType `json:"scope,omitempty"`
	}{
		TokenRequestProxy: (*TokenRequestProxy)(p),
		ClientId:          p.ClientId,
		GrantType:         p.GrantType,
		Scope:             p.Scope,
	})
}

func NewTokenRequest() *TokenRequest {
	p := new(TokenRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.TokenRequest"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.TokenRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the OIDC token response
*/
type TokenResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Generated Access Token to be utilized by the client
	*/
	AccessToken *string `json:"access_token,omitempty"`
	/*
	  Token expiration time in seconds
	*/
	ExpiresIn *int64 `json:"expires_in"`
	/*
	  Generated ID Token to be utilized by the client
	*/
	IdToken *string `json:"id_token"`
	/*
	  Refresh token which is issued in case of auth code, refresh token or token exchange flows
	*/
	RefreshToken *string `json:"refresh_token,omitempty"`
	/*
	  Generated ID Token type
	*/
	TokenType *string `json:"token_type"`
}

func (p *TokenResponse) MarshalJSON() ([]byte, error) {
	type TokenResponseProxy TokenResponse
	return json.Marshal(struct {
		*TokenResponseProxy
		ExpiresIn *int64  `json:"expires_in,omitempty"`
		IdToken   *string `json:"id_token,omitempty"`
		TokenType *string `json:"token_type,omitempty"`
	}{
		TokenResponseProxy: (*TokenResponseProxy)(p),
		ExpiresIn:          p.ExpiresIn,
		IdToken:            p.IdToken,
		TokenType:          p.TokenType,
	})
}

func NewTokenResponse() *TokenResponse {
	p := new(TokenResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.TokenResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.TokenResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /iam/v4.0.b1/authn/cert-auth-providers/{extId} Put operation
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

func NewUpdateCertAuthProviderApiResponse() *UpdateCertAuthProviderApiResponse {
	p := new(UpdateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UpdateCertAuthProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/directory-services/{extId} Put operation
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

func NewUpdateDirectoryServiceApiResponse() *UpdateDirectoryServiceApiResponse {
	p := new(UpdateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UpdateDirectoryServiceApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/saml-identity-providers/{extId} Put operation
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

func NewUpdateSamlIdentityProviderApiResponse() *UpdateSamlIdentityProviderApiResponse {
	p := new(UpdateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UpdateSamlIdentityProviderApiResponse"}
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
REST response for all response codes in API path /iam/v4.0.b1/authn/users/{extId} Put operation
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

func NewUpdateUserApiResponse() *UpdateUserApiResponse {
	p := new(UpdateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UpdateUserApiResponse"}
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
Use for the OIDC key
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
Information of the user
*/
type User struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Any additional attribute for the user
	*/
	AdditionalAttributes []import3.KVPair `json:"additionalAttributes,omitempty"`
	/*
	  Bucket access keys for the user
	*/
	BucketsAccessKeys []BucketsAccessKey `json:"bucketsAccessKeys,omitempty"`
	/*
	  User or service who created the user
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the user
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Display name for the user
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Email Id for the user
	*/
	EmailId *string `json:"emailId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  First name for the user
	*/
	FirstName *string `json:"firstName,omitempty"`
	/*
	  Identifier of the IDP for the user
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Flag to force the user to reset password
	*/
	IsForceResetPasswordEnabled *bool `json:"isForceResetPasswordEnabled,omitempty"`
	/*
	  Last successful logged in time for the user
	*/
	LastLoginTime *time.Time `json:"lastLoginTime,omitempty"`
	/*
	  Last name for the user
	*/
	LastName *string `json:"lastName,omitempty"`
	/*
	  Last updated time of the user
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Default locale for the user
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  Middle name for the user
	*/
	MiddleInitial *string `json:"middleInitial,omitempty"`
	/*
	  Password for the user
	*/
	Password *string `json:"password,omitempty"`
	/*
	  Default Region for the user
	*/
	Region *string `json:"region,omitempty"`

	Status *UserStatusType `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserType *UserType `json:"userType,omitempty"`
	/*
	  Identifier for the user in the form an email address
	*/
	Username *string `json:"username,omitempty"`
}

func NewUser() *User {
	p := new(User)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.User"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.User"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User configuration for OpenLDAP directory service
*/
type UserConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Object class in the OpenLDAP system that corresponds to users
	*/
	UserObjectClass *string `json:"userObjectClass"`
	/*
	  Base DN for user search
	*/
	UserSearchBase *string `json:"userSearchBase"`
	/*
	  Unique identifier for each user which can be used in authentication
	*/
	UsernameAttribute *string `json:"usernameAttribute"`
}

func (p *UserConfiguration) MarshalJSON() ([]byte, error) {
	type UserConfigurationProxy UserConfiguration
	return json.Marshal(struct {
		*UserConfigurationProxy
		UserObjectClass   *string `json:"userObjectClass,omitempty"`
		UserSearchBase    *string `json:"userSearchBase,omitempty"`
		UsernameAttribute *string `json:"usernameAttribute,omitempty"`
	}{
		UserConfigurationProxy: (*UserConfigurationProxy)(p),
		UserObjectClass:        p.UserObjectClass,
		UserSearchBase:         p.UserSearchBase,
		UsernameAttribute:      p.UsernameAttribute,
	})
}

func NewUserConfiguration() *UserConfiguration {
	p := new(UserConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserConfiguration"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UserConfiguration"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information of the user group
*/
type UserGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User or service who created the user group
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Creation time of the user group
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  Identifier for the user group in the form of a distinguished name
	*/
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupType *GroupType `json:"groupType,omitempty"`
	/*
	  Identifier of the IDP for the user group
	*/
	IdpId *string `json:"idpId,omitempty"`
	/*
	  Last updated time of the user group
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Common Name of the user group
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewUserGroup() *UserGroup {
	p := new(UserGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UserGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
User group configuration for OpenLDAP directory service
*/
type UserGroupConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute in a group that associates users to the group
	*/
	GroupMemberAttribute *string `json:"groupMemberAttribute"`
	/*
	  User attribute value that will be used in group entity to associate user to the group
	*/
	GroupMemberAttributeValue *string `json:"groupMemberAttributeValue"`
	/*
	  Object class in the OpenLDAP system that corresponds to groups
	*/
	GroupObjectClass *string `json:"groupObjectClass"`
	/*
	  Base DN for group search
	*/
	GroupSearchBase *string `json:"groupSearchBase"`
}

func (p *UserGroupConfiguration) MarshalJSON() ([]byte, error) {
	type UserGroupConfigurationProxy UserGroupConfiguration
	return json.Marshal(struct {
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
	})
}

func NewUserGroupConfiguration() *UserGroupConfiguration {
	p := new(UserGroupConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserGroupConfiguration"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UserGroupConfiguration"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information to change state of user
*/
type UserStateUpdate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Status *UserStatusType `json:"status"`
}

func (p *UserStateUpdate) MarshalJSON() ([]byte, error) {
	type UserStateUpdateProxy UserStateUpdate
	return json.Marshal(struct {
		*UserStateUpdateProxy
		Status *UserStatusType `json:"status,omitempty"`
	}{
		UserStateUpdateProxy: (*UserStateUpdateProxy)(p),
		Status:               p.Status,
	})
}

func NewUserStateUpdate() *UserStateUpdate {
	p := new(UserStateUpdate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserStateUpdate"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.UserStateUpdate"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the user
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
Type of the user
*/
type UserType int

const (
	USERTYPE_UNKNOWN  UserType = 0
	USERTYPE_REDACTED UserType = 1
	USERTYPE_LOCAL    UserType = 2
	USERTYPE_SAML     UserType = 3
	USERTYPE_LDAP     UserType = 4
	USERTYPE_EXTERNAL UserType = 5
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
REST response for all response codes in API path /iam/v4.0.b1/authn/api-keys/$actions/validate Post operation
*/
type ValidateApiKeyApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfValidateApiKeyApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewValidateApiKeyApiResponse() *ValidateApiKeyApiResponse {
	p := new(ValidateApiKeyApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ValidateApiKeyApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.b1.authn.ValidateApiKeyApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ValidateApiKeyApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ValidateApiKeyApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfValidateApiKeyApiResponseData()
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

type OneOfGetServiceAccountKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ServiceAccountKey     `json:"-"`
}

func NewOneOfGetServiceAccountKeyApiResponseData() *OneOfGetServiceAccountKeyApiResponseData {
	p := new(OneOfGetServiceAccountKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetServiceAccountKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetServiceAccountKeyApiResponseData is nil"))
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
	case ServiceAccountKey:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ServiceAccountKey)
		}
		*p.oneOfType0 = v.(ServiceAccountKey)
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

func (p *OneOfGetServiceAccountKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetServiceAccountKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ServiceAccountKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ServiceAccountKey" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ServiceAccountKey)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetServiceAccountKeyApiResponseData"))
}

func (p *OneOfGetServiceAccountKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetServiceAccountKeyApiResponseData")
}

type OneOfListDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []DirectoryService     `json:"-"`
}

func NewOneOfListDirectoryServiceApiResponseData() *OneOfListDirectoryServiceApiResponseData {
	p := new(OneOfListDirectoryServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListDirectoryServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListDirectoryServiceApiResponseData is nil"))
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

func (p *OneOfListDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.DirectoryService>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListDirectoryServiceApiResponseData"))
}

func (p *OneOfListDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.DirectoryService>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListDirectoryServiceApiResponseData")
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

type OneOfUpdateSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSamlIdentityProviderApiResponseData"))
}

func (p *OneOfUpdateSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSamlIdentityProviderApiResponseData")
}

type OneOfCreateDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateDirectoryServiceApiResponseData"))
}

func (p *OneOfCreateDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateDirectoryServiceApiResponseData")
}

type OneOfGetUserInfoApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *OidcUserinfo          `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetUserInfoApiResponseData() *OneOfGetUserInfoApiResponseData {
	p := new(OneOfGetUserInfoApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetUserInfoApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetUserInfoApiResponseData is nil"))
	}
	switch v.(type) {
	case OidcUserinfo:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(OidcUserinfo)
		}
		*p.oneOfType0 = v.(OidcUserinfo)
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

func (p *OneOfGetUserInfoApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetUserInfoApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(OidcUserinfo)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.OidcUserinfo" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(OidcUserinfo)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUserInfoApiResponseData"))
}

func (p *OneOfGetUserInfoApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetUserInfoApiResponseData")
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

type OneOfDeleteServiceAccountKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteServiceAccountKeyApiResponseData() *OneOfDeleteServiceAccountKeyApiResponseData {
	p := new(OneOfDeleteServiceAccountKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteServiceAccountKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteServiceAccountKeyApiResponseData is nil"))
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

func (p *OneOfDeleteServiceAccountKeyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteServiceAccountKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteServiceAccountKeyApiResponseData"))
}

func (p *OneOfDeleteServiceAccountKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteServiceAccountKeyApiResponseData")
}

type OneOfCreateServiceAccountKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ServiceAccountKey     `json:"-"`
}

func NewOneOfCreateServiceAccountKeyApiResponseData() *OneOfCreateServiceAccountKeyApiResponseData {
	p := new(OneOfCreateServiceAccountKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateServiceAccountKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateServiceAccountKeyApiResponseData is nil"))
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
	case ServiceAccountKey:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ServiceAccountKey)
		}
		*p.oneOfType0 = v.(ServiceAccountKey)
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

func (p *OneOfCreateServiceAccountKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateServiceAccountKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ServiceAccountKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ServiceAccountKey" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ServiceAccountKey)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateServiceAccountKeyApiResponseData"))
}

func (p *OneOfCreateServiceAccountKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateServiceAccountKeyApiResponseData")
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

type OneOfConnectionDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfConnectionDirectoryServiceApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfConnectionDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfConnectionDirectoryServiceApiResponseData"))
}

func (p *OneOfConnectionDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfConnectionDirectoryServiceApiResponseData")
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

type OneOfCreateSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateSamlIdentityProviderApiResponseData"))
}

func (p *OneOfCreateSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateSamlIdentityProviderApiResponseData")
}

type OneOfDeleteServiceAccountApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteServiceAccountApiResponseData() *OneOfDeleteServiceAccountApiResponseData {
	p := new(OneOfDeleteServiceAccountApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteServiceAccountApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteServiceAccountApiResponseData is nil"))
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

func (p *OneOfDeleteServiceAccountApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteServiceAccountApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteServiceAccountApiResponseData"))
}

func (p *OneOfDeleteServiceAccountApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteServiceAccountApiResponseData")
}

type OneOfCreateApiKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiKey                `json:"-"`
}

func NewOneOfCreateApiKeyApiResponseData() *OneOfCreateApiKeyApiResponseData {
	p := new(OneOfCreateApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateApiKeyApiResponseData is nil"))
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
	case ApiKey:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiKey)
		}
		*p.oneOfType0 = v.(ApiKey)
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

func (p *OneOfCreateApiKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ApiKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ApiKey" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiKey)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateApiKeyApiResponseData"))
}

func (p *OneOfCreateApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateApiKeyApiResponseData")
}

type OneOfListUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []UserGroup            `json:"-"`
}

func NewOneOfListUserGroupApiResponseData() *OneOfListUserGroupApiResponseData {
	p := new(OneOfListUserGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUserGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUserGroupApiResponseData is nil"))
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

func (p *OneOfListUserGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.UserGroup>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUserGroupApiResponseData"))
}

func (p *OneOfListUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.UserGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUserGroupApiResponseData")
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

type OneOfGetServiceAccountApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ServiceAccount        `json:"-"`
}

func NewOneOfGetServiceAccountApiResponseData() *OneOfGetServiceAccountApiResponseData {
	p := new(OneOfGetServiceAccountApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetServiceAccountApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetServiceAccountApiResponseData is nil"))
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
	case ServiceAccount:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ServiceAccount)
		}
		*p.oneOfType0 = v.(ServiceAccount)
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

func (p *OneOfGetServiceAccountApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetServiceAccountApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ServiceAccount)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ServiceAccount" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ServiceAccount)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetServiceAccountApiResponseData"))
}

func (p *OneOfGetServiceAccountApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetServiceAccountApiResponseData")
}

type OneOfGetFederationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Federation            `json:"-"`
}

func NewOneOfGetFederationApiResponseData() *OneOfGetFederationApiResponseData {
	p := new(OneOfGetFederationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetFederationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetFederationApiResponseData is nil"))
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
	case Federation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Federation)
		}
		*p.oneOfType0 = v.(Federation)
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

func (p *OneOfGetFederationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetFederationApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(Federation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.Federation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Federation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetFederationApiResponseData"))
}

func (p *OneOfGetFederationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetFederationApiResponseData")
}

type OneOfGetSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *SamlIdentityProvider  `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSamlIdentityProviderApiResponseData"))
}

func (p *OneOfGetSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSamlIdentityProviderApiResponseData")
}

type OneOfCreateServiceAccountApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ServiceAccount        `json:"-"`
}

func NewOneOfCreateServiceAccountApiResponseData() *OneOfCreateServiceAccountApiResponseData {
	p := new(OneOfCreateServiceAccountApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateServiceAccountApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateServiceAccountApiResponseData is nil"))
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
	case ServiceAccount:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ServiceAccount)
		}
		*p.oneOfType0 = v.(ServiceAccount)
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

func (p *OneOfCreateServiceAccountApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateServiceAccountApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ServiceAccount)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ServiceAccount" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ServiceAccount)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateServiceAccountApiResponseData"))
}

func (p *OneOfCreateServiceAccountApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateServiceAccountApiResponseData")
}

type OneOfUpdateCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCertAuthProviderApiResponseData"))
}

func (p *OneOfUpdateCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCertAuthProviderApiResponseData")
}

type OneOfListCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []CertAuthProvider     `json:"-"`
}

func NewOneOfListCertAuthProviderApiResponseData() *OneOfListCertAuthProviderApiResponseData {
	p := new(OneOfListCertAuthProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCertAuthProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCertAuthProviderApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.CertAuthProvider>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCertAuthProviderApiResponseData"))
}

func (p *OneOfListCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.CertAuthProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListCertAuthProviderApiResponseData")
}

type OneOfDeleteFederationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteFederationApiResponseData() *OneOfDeleteFederationApiResponseData {
	p := new(OneOfDeleteFederationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteFederationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteFederationApiResponseData is nil"))
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

func (p *OneOfDeleteFederationApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteFederationApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteFederationApiResponseData"))
}

func (p *OneOfDeleteFederationApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteFederationApiResponseData")
}

type OneOfDeleteApiKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteApiKeyApiResponseData() *OneOfDeleteApiKeyApiResponseData {
	p := new(OneOfDeleteApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteApiKeyApiResponseData is nil"))
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

func (p *OneOfDeleteApiKeyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteApiKeyApiResponseData"))
}

func (p *OneOfDeleteApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteApiKeyApiResponseData")
}

type OneOfListServiceAccountKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ServiceAccountKey    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListServiceAccountKeyApiResponseData() *OneOfListServiceAccountKeyApiResponseData {
	p := new(OneOfListServiceAccountKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListServiceAccountKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListServiceAccountKeyApiResponseData is nil"))
	}
	switch v.(type) {
	case []ServiceAccountKey:
		p.oneOfType0 = v.([]ServiceAccountKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.ServiceAccountKey>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.ServiceAccountKey>"
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

func (p *OneOfListServiceAccountKeyApiResponseData) GetValue() interface{} {
	if "List<iam.v4.authn.ServiceAccountKey>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListServiceAccountKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ServiceAccountKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.ServiceAccountKey" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.ServiceAccountKey>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.ServiceAccountKey>"
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListServiceAccountKeyApiResponseData"))
}

func (p *OneOfListServiceAccountKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<iam.v4.authn.ServiceAccountKey>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListServiceAccountKeyApiResponseData")
}

type OneOfCreateUserKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *BucketsAccessKey      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfCreateUserKeyApiResponseData() *OneOfCreateUserKeyApiResponseData {
	p := new(OneOfCreateUserKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateUserKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateUserKeyApiResponseData is nil"))
	}
	switch v.(type) {
	case BucketsAccessKey:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(BucketsAccessKey)
		}
		*p.oneOfType0 = v.(BucketsAccessKey)
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

func (p *OneOfCreateUserKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateUserKeyApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(BucketsAccessKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.BucketsAccessKey" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(BucketsAccessKey)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUserKeyApiResponseData"))
}

func (p *OneOfCreateUserKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUserKeyApiResponseData")
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

type OneOfCreateCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCertAuthProviderApiResponseData"))
}

func (p *OneOfCreateCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCertAuthProviderApiResponseData")
}

type OneOfGetUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *UserGroup             `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetUserGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetUserGroupApiResponseData"))
}

func (p *OneOfGetUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetUserGroupApiResponseData")
}

type OneOfGetSamlSpMetadataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfGetSamlSpMetadataApiResponseData) GetValue() interface{} {
	if "String" == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSamlSpMetadataApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSamlSpMetadataApiResponseData"))
}

func (p *OneOfGetSamlSpMetadataApiResponseData) MarshalJSON() ([]byte, error) {
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSamlSpMetadataApiResponseData")
}

type OneOfGetCertAuthProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *CertAuthProvider      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetCertAuthProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetCertAuthProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCertAuthProviderApiResponseData"))
}

func (p *OneOfGetCertAuthProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetCertAuthProviderApiResponseData")
}

type OneOfValidateApiKeyApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import1.ErrorResponse  `json:"-"`
	oneOfType0    *ApiKeyValidateResponse `json:"-"`
}

func NewOneOfValidateApiKeyApiResponseData() *OneOfValidateApiKeyApiResponseData {
	p := new(OneOfValidateApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfValidateApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfValidateApiKeyApiResponseData is nil"))
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
	case ApiKeyValidateResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiKeyValidateResponse)
		}
		*p.oneOfType0 = v.(ApiKeyValidateResponse)
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

func (p *OneOfValidateApiKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfValidateApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ApiKeyValidateResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ApiKeyValidateResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiKeyValidateResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfValidateApiKeyApiResponseData"))
}

func (p *OneOfValidateApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfValidateApiKeyApiResponseData")
}

type OneOfListServiceAccountApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []ServiceAccount       `json:"-"`
}

func NewOneOfListServiceAccountApiResponseData() *OneOfListServiceAccountApiResponseData {
	p := new(OneOfListServiceAccountApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListServiceAccountApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListServiceAccountApiResponseData is nil"))
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
	case []ServiceAccount:
		p.oneOfType0 = v.([]ServiceAccount)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.ServiceAccount>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.ServiceAccount>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListServiceAccountApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.ServiceAccount>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListServiceAccountApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]ServiceAccount)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.ServiceAccount" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.ServiceAccount>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.ServiceAccount>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListServiceAccountApiResponseData"))
}

func (p *OneOfListServiceAccountApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.ServiceAccount>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListServiceAccountApiResponseData")
}

type OneOfListSamlIdentityProviderApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []SamlIdentityProvider `json:"-"`
}

func NewOneOfListSamlIdentityProviderApiResponseData() *OneOfListSamlIdentityProviderApiResponseData {
	p := new(OneOfListSamlIdentityProviderApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListSamlIdentityProviderApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListSamlIdentityProviderApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListSamlIdentityProviderApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.SamlIdentityProvider>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListSamlIdentityProviderApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListSamlIdentityProviderApiResponseData"))
}

func (p *OneOfListSamlIdentityProviderApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.SamlIdentityProvider>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListSamlIdentityProviderApiResponseData")
}

type OneOfActivateUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfActivateUserApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfActivateUserApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfActivateUserApiResponseData"))
}

func (p *OneOfActivateUserApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfActivateUserApiResponseData")
}

type OneOfListUserKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []BucketsAccessKey     `json:"-"`
}

func NewOneOfListUserKeyApiResponseData() *OneOfListUserKeyApiResponseData {
	p := new(OneOfListUserKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUserKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUserKeyApiResponseData is nil"))
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
	case []BucketsAccessKey:
		p.oneOfType0 = v.([]BucketsAccessKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.BucketsAccessKey>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.BucketsAccessKey>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUserKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.BucketsAccessKey>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUserKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]BucketsAccessKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.BucketsAccessKey" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.BucketsAccessKey>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.BucketsAccessKey>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUserKeyApiResponseData"))
}

func (p *OneOfListUserKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.BucketsAccessKey>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUserKeyApiResponseData")
}

type OneOfListApiKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []ApiKey               `json:"-"`
}

func NewOneOfListApiKeyApiResponseData() *OneOfListApiKeyApiResponseData {
	p := new(OneOfListApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListApiKeyApiResponseData is nil"))
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
	case []ApiKey:
		p.oneOfType0 = v.([]ApiKey)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<iam.v4.authn.ApiKey>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<iam.v4.authn.ApiKey>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListApiKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.ApiKey>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]ApiKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "iam.v4.authn.ApiKey" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<iam.v4.authn.ApiKey>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<iam.v4.authn.ApiKey>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListApiKeyApiResponseData"))
}

func (p *OneOfListApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.ApiKey>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListApiKeyApiResponseData")
}

type OneOfRevokeApiKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfRevokeApiKeyApiResponseData() *OneOfRevokeApiKeyApiResponseData {
	p := new(OneOfRevokeApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRevokeApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRevokeApiKeyApiResponseData is nil"))
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

func (p *OneOfRevokeApiKeyApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRevokeApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRevokeApiKeyApiResponseData"))
}

func (p *OneOfRevokeApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRevokeApiKeyApiResponseData")
}

type OneOfCreateFederationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *Federation            `json:"-"`
}

func NewOneOfCreateFederationApiResponseData() *OneOfCreateFederationApiResponseData {
	p := new(OneOfCreateFederationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateFederationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateFederationApiResponseData is nil"))
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
	case Federation:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Federation)
		}
		*p.oneOfType0 = v.(Federation)
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

func (p *OneOfCreateFederationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateFederationApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(Federation)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.Federation" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Federation)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateFederationApiResponseData"))
}

func (p *OneOfCreateFederationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateFederationApiResponseData")
}

type OneOfGetApiKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiKey                `json:"-"`
}

func NewOneOfGetApiKeyApiResponseData() *OneOfGetApiKeyApiResponseData {
	p := new(OneOfGetApiKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetApiKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetApiKeyApiResponseData is nil"))
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
	case ApiKey:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiKey)
		}
		*p.oneOfType0 = v.(ApiKey)
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

func (p *OneOfGetApiKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetApiKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(ApiKey)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.ApiKey" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiKey)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetApiKeyApiResponseData"))
}

func (p *OneOfGetApiKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetApiKeyApiResponseData")
}

type OneOfResetUserPasswordApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
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

func (p *OneOfResetUserPasswordApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfResetUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResetUserPasswordApiResponseData"))
}

func (p *OneOfResetUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfResetUserPasswordApiResponseData")
}

type OneOfListUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    []User                 `json:"-"`
}

func NewOneOfListUserApiResponseData() *OneOfListUserApiResponseData {
	p := new(OneOfListUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListUserApiResponseData is nil"))
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListUserApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<iam.v4.authn.User>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListUserApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListUserApiResponseData"))
}

func (p *OneOfListUserApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<iam.v4.authn.User>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListUserApiResponseData")
}

type OneOfChangeUserPasswordApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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

func (p *OneOfChangeUserPasswordApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangeUserPasswordApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangeUserPasswordApiResponseData"))
}

func (p *OneOfChangeUserPasswordApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangeUserPasswordApiResponseData")
}

type OneOfUpdateDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateDirectoryServiceApiResponseData"))
}

func (p *OneOfUpdateDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateDirectoryServiceApiResponseData")
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

type OneOfSearchDirectoryServiceApiResponseData struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType400  *import1.ErrorResponse        `json:"-"`
	oneOfType0    *DirectoryServiceSearchResult `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSearchDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfSearchDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSearchDirectoryServiceApiResponseData"))
}

func (p *OneOfSearchDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfSearchDirectoryServiceApiResponseData")
}

type OneOfDeleteUserApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1    *interface{}           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfDeleteUserApiResponseData() *OneOfDeleteUserApiResponseData {
	p := new(OneOfDeleteUserApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteUserApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteUserApiResponseData is nil"))
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

func (p *OneOfDeleteUserApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteUserApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteUserApiResponseData"))
}

func (p *OneOfDeleteUserApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteUserApiResponseData")
}

type OneOfGetSecretKeyApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *SecretKeyResponse     `json:"-"`
}

func NewOneOfGetSecretKeyApiResponseData() *OneOfGetSecretKeyApiResponseData {
	p := new(OneOfGetSecretKeyApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSecretKeyApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSecretKeyApiResponseData is nil"))
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
	case SecretKeyResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SecretKeyResponse)
		}
		*p.oneOfType0 = v.(SecretKeyResponse)
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

func (p *OneOfGetSecretKeyApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetSecretKeyApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(SecretKeyResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "iam.v4.authn.SecretKeyResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SecretKeyResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSecretKeyApiResponseData"))
}

func (p *OneOfGetSecretKeyApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetSecretKeyApiResponseData")
}

type OneOfCreateUserGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *UserGroup             `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateUserGroupApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateUserGroupApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateUserGroupApiResponseData"))
}

func (p *OneOfCreateUserGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateUserGroupApiResponseData")
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

type OneOfGetDirectoryServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *DirectoryService      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetDirectoryServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetDirectoryServiceApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDirectoryServiceApiResponseData"))
}

func (p *OneOfGetDirectoryServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetDirectoryServiceApiResponseData")
}
