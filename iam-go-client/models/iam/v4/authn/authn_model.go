/*
 * Generated file models/iam/v4/authn/authn_model.go.
 *
 * Product version: 4.0.2-alpha-1
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users/{extId}/$actions/change-state Post operation
*/
type ActivateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfActivateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewActivateUserApiResponse() *ActivateUserApiResponse {
	p := new(ActivateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ActivateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ActivateUserApiResponse"}
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

/**
A certificate based authentication provider
*/
type CertAuthProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the uploaded CA chain file
	*/
	CaCertFileName *string `json:"caCertFileName"`
	/**
	  Flag to enable/disable Cert Auth for the current certificate based authentication provider.
	*/
	CacEnabled *bool `json:"cacEnabled"`
	/**
	  Flag to enable/disable CAC for the current certificate based authentication provider.
	*/
	CertAuthEnabled *bool `json:"certAuthEnabled"`

	CertRevocationInfo *CertRevocationInfo `json:"certRevocationInfo,omitempty"`
	/**
	  CA chain file
	*/
	ClientCaChain *string `json:"clientCaChain"`
	/**
	  User or service who created the CertAuthProvider
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Creation time of the certificate based authentication provider
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  UUID of an existing directory service
	*/
	DirSvcExtID *string `json:"dirSvcExtID,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  Last updated time of the certificate based authentication provider
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Unique name of the certificate based authentication provider.
	*/
	Name *string `json:"name"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *CertAuthProvider) MarshalJSON() ([]byte, error) {
	type CertAuthProviderProxy CertAuthProvider
	return json.Marshal(struct {
		*CertAuthProviderProxy
		CaCertFileName  *string `json:"caCertFileName,omitempty"`
		CacEnabled      *bool   `json:"cacEnabled,omitempty"`
		CertAuthEnabled *bool   `json:"certAuthEnabled,omitempty"`
		ClientCaChain   *string `json:"clientCaChain,omitempty"`
		Name            *string `json:"name,omitempty"`
	}{
		CertAuthProviderProxy: (*CertAuthProviderProxy)(p),
		CaCertFileName:        p.CaCertFileName,
		CacEnabled:            p.CacEnabled,
		CertAuthEnabled:       p.CertAuthEnabled,
		ClientCaChain:         p.ClientCaChain,
		Name:                  p.Name,
	})
}

func NewCertAuthProvider() *CertAuthProvider {
	p := new(CertAuthProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertAuthProvider"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CertAuthProvider"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Configuration details used for determining if the client certificate is revoked.
*/
type CertRevocationInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of the CRL distribution points which will be used to fetch the CRLs
	*/
	CrlDps []string `json:"crlDps,omitempty"`
	/**
	  Flag to enable/disable CRL revocation check
	*/
	CrlEnabled *bool `json:"crlEnabled,omitempty"`
	/**
	  Interval in seconds at which the CRL should be fetched from the CRLDP, default = 86400 seconds(1 day)
	*/
	GlobalCrlRefreshInterval *int `json:"globalCrlRefreshInterval,omitempty"`
	/**
	  Flag to enable/disable OCSP revocation check
	*/
	OcspEnabled *bool `json:"ocspEnabled,omitempty"`
	/**
	  URL of the OCSP responder used to override the URL from AIA extension
	*/
	OcspResponder *string `json:"ocspResponder,omitempty"`
}

func NewCertRevocationInfo() *CertRevocationInfo {
	p := new(CertRevocationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CertRevocationInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CertRevocationInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users/$actions/change-password Post operation
*/
type ChangeUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangeUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewChangeUserPasswordApiResponse() *ChangeUserPasswordApiResponse {
	p := new(ChangeUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ChangeUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ChangeUserPasswordApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services/{extId}/connection-status Post operation
*/
type ConnectionDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfConnectionDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewConnectionDirectoryServiceApiResponse() *ConnectionDirectoryServiceApiResponse {
	p := new(ConnectionDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ConnectionDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ConnectionDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/cert-auth-providers Post operation
*/
type CreateCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateCertAuthProviderApiResponse() *CreateCertAuthProviderApiResponse {
	p := new(CreateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CreateCertAuthProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services Post operation
*/
type CreateDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateDirectoryServiceApiResponse() *CreateDirectoryServiceApiResponse {
	p := new(CreateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CreateDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/saml-identity-providers Post operation
*/
type CreateSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateSamlIdentityProviderApiResponse() *CreateSamlIdentityProviderApiResponse {
	p := new(CreateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CreateSamlIdentityProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users Post operation
*/
type CreateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUserApiResponse() *CreateUserApiResponse {
	p := new(CreateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CreateUserApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/user-groups Post operation
*/
type CreateUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewCreateUserGroupApiResponse() *CreateUserGroupApiResponse {
	p := new(CreateUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.CreateUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.CreateUserGroupApiResponse"}
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

/**
Information of a directory service
*/
type DirectoryService struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  User or service who created the directory service
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Creation time of the directory service
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DirectoryType *DirectoryType `json:"directoryType"`
	/**
	  Domain name for the directory service
	*/
	DomainName *string `json:"domainName"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupSearchType *GroupSearchType `json:"groupSearchType,omitempty"`
	/**
	  Last updated time of the directory service
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Name for the directory service
	*/
	Name *string `json:"name"`

	OpenLdapConfiguration *OpenLdapConfig `json:"openLdapConfiguration,omitempty"`

	ServiceAccount *DsServiceAccount `json:"serviceAccount"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  URL for the directory service
	*/
	Url *string `json:"url"`
	/**
	  List of allowed user groups for the directory service
	*/
	WhiteListedGroups []string `json:"whiteListedGroups,omitempty"`
}

func (p *DirectoryService) MarshalJSON() ([]byte, error) {
	type DirectoryServiceProxy DirectoryService
	return json.Marshal(struct {
		*DirectoryServiceProxy
		DirectoryType  *DirectoryType    `json:"directoryType,omitempty"`
		DomainName     *string           `json:"domainName,omitempty"`
		Name           *string           `json:"name,omitempty"`
		ServiceAccount *DsServiceAccount `json:"serviceAccount,omitempty"`
		Url            *string           `json:"url,omitempty"`
	}{
		DirectoryServiceProxy: (*DirectoryServiceProxy)(p),
		DirectoryType:         p.DirectoryType,
		DomainName:            p.DomainName,
		Name:                  p.Name,
		ServiceAccount:        p.ServiceAccount,
		Url:                   p.Url,
	})
}

func NewDirectoryService() *DirectoryService {
	p := new(DirectoryService)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryService"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryService"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information for directory service connection request
*/
type DirectoryServiceConnectionRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Password to connect to the directory service
	*/
	Password *string `json:"password"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceConnectionRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
AD/LDAP information of the user
*/
type DirectoryServiceInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of AD/LDAP groups having the user
	*/
	Groups []DirectoryServiceInfoGroup `json:"groups"`
	/**
	  List of AD/LDAP OUs having the given user
	*/
	Ous []DirectoryServiceInfoOu `json:"ous"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information of AD group having the user
*/
type DirectoryServiceInfoGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Name of the group
	*/
	Name *string `json:"name"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceInfoGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information of OUs having the user
*/
type DirectoryServiceInfoOu struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Name of the OU
	*/
	Name *string `json:"name"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceInfoOu"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information of searched attributes
*/
type DirectoryServiceSearchAttribute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the attribute
	*/
	Name *string `json:"name,omitempty"`
	/**
	  Value of the attribute
	*/
	Values []string `json:"values,omitempty"`
}

func NewDirectoryServiceSearchAttribute() *DirectoryServiceSearchAttribute {
	p := new(DirectoryServiceSearchAttribute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchAttribute"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceSearchAttribute"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information of single search entity
*/
type DirectoryServiceSearchEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Attributes []DirectoryServiceSearchAttribute `json:"attributes,omitempty"`
	/**
	  Type of entity either user or group
	*/
	EntityType *string `json:"entityType,omitempty"`
	/**
	  Name of the entity in canonical format
	*/
	Name *string `json:"name,omitempty"`
}

func NewDirectoryServiceSearchEntity() *DirectoryServiceSearchEntity {
	p := new(DirectoryServiceSearchEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchEntity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceSearchEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information for directory search query
*/
type DirectoryServiceSearchQuery struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Flag indicating whether the search should be a wildcard search or not
	*/
	IsWildcardSearch *bool `json:"isWildcardSearch,omitempty"`
	/**
	  Query string for directory service search
	*/
	Query *string `json:"query"`
	/**
	  Attributes the search operation returns
	*/
	ReturnedAttributes []string `json:"returnedAttributes,omitempty"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceSearchQuery"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsWildcardSearch = new(bool)
	*p.IsWildcardSearch = true

	return p
}

/**
Information of directory service search result
*/
type DirectoryServiceSearchResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Domain name for the directory service
	*/
	DomainName *string `json:"domainName,omitempty"`
	/**
	  Result of directory service search
	*/
	SearchResults []DirectoryServiceSearchEntity `json:"searchResults,omitempty"`
}

func NewDirectoryServiceSearchResult() *DirectoryServiceSearchResult {
	p := new(DirectoryServiceSearchResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.DirectoryServiceSearchResult"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DirectoryServiceSearchResult"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Type of directory service
*/
type DirectoryType int

const (
	DIRECTORYTYPE_UNKNOWN          DirectoryType = 0
	DIRECTORYTYPE_REDACTED         DirectoryType = 1
	DIRECTORYTYPE_ACTIVE_DIRECTORY DirectoryType = 2
	DIRECTORYTYPE_OPEN_LDAP        DirectoryType = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
Service account information to connect to the directory service
*/
type DsServiceAccount struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Password to connect to the directory service
	*/
	Password *string `json:"password"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.DsServiceAccount"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/cert-auth-providers/{extId} Get operation
*/
type GetCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetCertAuthProviderApiResponse() *GetCertAuthProviderApiResponse {
	p := new(GetCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetCertAuthProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services/{extId} Get operation
*/
type GetDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDirectoryServiceApiResponse() *GetDirectoryServiceApiResponse {
	p := new(GetDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/saml-identity-providers/{extId} Get operation
*/
type GetSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSamlIdentityProviderApiResponse() *GetSamlIdentityProviderApiResponse {
	p := new(GetSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetSamlIdentityProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/saml-identity-providers/saml20/sp-metadata Get operation
*/
type GetSamlSpMetadataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSamlSpMetadataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSamlSpMetadataApiResponse() *GetSamlSpMetadataApiResponse {
	p := new(GetSamlSpMetadataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetSamlSpMetadataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetSamlSpMetadataApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users/{extId} Get operation
*/
type GetUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUserApiResponse() *GetUserApiResponse {
	p := new(GetUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetUserApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/user-groups/{extId} Get operation
*/
type GetUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetUserGroupApiResponse() *GetUserGroupApiResponse {
	p := new(GetUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.GetUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.GetUserGroupApiResponse"}
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

/**
Group membership search type for the directory service
*/
type GroupSearchType int

const (
	GROUPSEARCHTYPE_UNKNOWN       GroupSearchType = 0
	GROUPSEARCHTYPE_REDACTED      GroupSearchType = 1
	GROUPSEARCHTYPE_NON_RECURSIVE GroupSearchType = 2
	GROUPSEARCHTYPE_RECURSIVE     GroupSearchType = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
Type of the user group
*/
type GroupType int

const (
	GROUPTYPE_UNKNOWN  GroupType = 0
	GROUPTYPE_REDACTED GroupType = 1
	GROUPTYPE_SAML     GroupType = 2
	GROUPTYPE_LDAP     GroupType = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
Information of the IDP
*/
type IdpProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Certificate for verification
	*/
	Certificate *string `json:"certificate"`
	/**
	  Entity Id of Identity provider
	*/
	EntityId *string `json:"entityId"`
	/**
	  Error URL of the Identity provider
	*/
	ErrorUrl *string `json:"errorUrl,omitempty"`
	/**
	  Login URL of the Identity provider
	*/
	LoginUrl *string `json:"loginUrl"`
	/**
	  Logout URL of the Identity provider
	*/
	LogoutUrl *string `json:"logoutUrl,omitempty"`

	NameIdPolicyFormat *NameIdPolicyFormat `json:"nameIdPolicyFormat,omitempty"`
}

func (p *IdpProperties) MarshalJSON() ([]byte, error) {
	type IdpPropertiesProxy IdpProperties
	return json.Marshal(struct {
		*IdpPropertiesProxy
		Certificate *string `json:"certificate,omitempty"`
		EntityId    *string `json:"entityId,omitempty"`
		LoginUrl    *string `json:"loginUrl,omitempty"`
	}{
		IdpPropertiesProxy: (*IdpPropertiesProxy)(p),
		Certificate:        p.Certificate,
		EntityId:           p.EntityId,
		LoginUrl:           p.LoginUrl,
	})
}

func NewIdpProperties() *IdpProperties {
	p := new(IdpProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.IdpProperties"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.IdpProperties"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/cert-auth-providers Get operation
*/
type ListCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListCertAuthProviderApiResponse() *ListCertAuthProviderApiResponse {
	p := new(ListCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ListCertAuthProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services Get operation
*/
type ListDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListDirectoryServiceApiResponse() *ListDirectoryServiceApiResponse {
	p := new(ListDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ListDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/saml-identity-providers Get operation
*/
type ListSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListSamlIdentityProviderApiResponse() *ListSamlIdentityProviderApiResponse {
	p := new(ListSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ListSamlIdentityProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users Get operation
*/
type ListUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUserApiResponse() *ListUserApiResponse {
	p := new(ListUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ListUserApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/user-groups Get operation
*/
type ListUserGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListUserGroupApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListUserGroupApiResponse() *ListUserGroupApiResponse {
	p := new(ListUserGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ListUserGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ListUserGroupApiResponse"}
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

/**
Name Id Policy format
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

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.OpenLdapConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information for password change of user
*/
type PasswordChangeRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  New password for the user
	*/
	NewPassword *string `json:"newPassword"`
	/**
	  Current password of the user
	*/
	OldPassword *string `json:"oldPassword"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.PasswordChangeRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information for password reset of user
*/
type PasswordResetRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.PasswordResetRequest"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users/{extId}/$actions/reset-password Post operation
*/
type ResetUserPasswordApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResetUserPasswordApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewResetUserPasswordApiResponse() *ResetUserPasswordApiResponse {
	p := new(ResetUserPasswordApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.ResetUserPasswordApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.ResetUserPasswordApiResponse"}
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

/**
Information of SAML IDP
*/
type SamlIdentityProvider struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  User or service who created the SAML Identity Provider
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Creation time of the SAML Identity Provider
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  SAML assertions for list of custom attribute elements
	*/
	CustomAttr []string `json:"customAttr,omitempty"`
	/**
	  SAML assertion email attribute element
	*/
	EmailAttr *string `json:"emailAttr"`
	/**
	  Flag indicating signing of SAML authentication requests
	*/
	EnableSignedAuthnReq *bool `json:"enableSignedAuthnReq,omitempty"`
	/**
	  It will be used as issuer in SAML authentication request
	*/
	EntityIssuer *string `json:"entityIssuer,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  SAML assertion groups attribute element
	*/
	GroupsAttr *string `json:"groupsAttr,omitempty"`
	/**
	  Delimiter is used to split the value of attribute into multiple groups
	*/
	GroupsDelim *string `json:"groupsDelim,omitempty"`
	/**
	  Base64 encoded metadata in xml format with IDP details
	*/
	IdpMetadata *string `json:"idpMetadata,omitempty"`
	/**
	  Metadata URL that provides IDP details
	*/
	IdpMetadataUrl *string `json:"idpMetadataUrl,omitempty"`

	IdpProperties *IdpProperties `json:"idpProperties,omitempty"`
	/**
	  Last updated time of the SAML Identity Provider
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Unique name of the IDP
	*/
	Name *string `json:"name"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  SAML assertion username attribute element
	*/
	UsernameAttr *string `json:"usernameAttr"`
}

func (p *SamlIdentityProvider) MarshalJSON() ([]byte, error) {
	type SamlIdentityProviderProxy SamlIdentityProvider
	return json.Marshal(struct {
		*SamlIdentityProviderProxy
		EmailAttr    *string `json:"emailAttr,omitempty"`
		Name         *string `json:"name,omitempty"`
		UsernameAttr *string `json:"usernameAttr,omitempty"`
	}{
		SamlIdentityProviderProxy: (*SamlIdentityProviderProxy)(p),
		EmailAttr:                 p.EmailAttr,
		Name:                      p.Name,
		UsernameAttr:              p.UsernameAttr,
	})
}

func NewSamlIdentityProvider() *SamlIdentityProvider {
	p := new(SamlIdentityProvider)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SamlIdentityProvider"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.SamlIdentityProvider"}
	p.UnknownFields_ = map[string]interface{}{}

	p.EmailAttr = new(string)
	*p.EmailAttr = "email"
	p.UsernameAttr = new(string)
	*p.UsernameAttr = "name"

	return p
}

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services/{extId}/search Post operation
*/
type SearchDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSearchDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSearchDirectoryServiceApiResponse() *SearchDirectoryServiceApiResponse {
	p := new(SearchDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.SearchDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.SearchDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/cert-auth-providers/{extId} Put operation
*/
type UpdateCertAuthProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCertAuthProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateCertAuthProviderApiResponse() *UpdateCertAuthProviderApiResponse {
	p := new(UpdateCertAuthProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateCertAuthProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UpdateCertAuthProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/directory-services/{extId} Put operation
*/
type UpdateDirectoryServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateDirectoryServiceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateDirectoryServiceApiResponse() *UpdateDirectoryServiceApiResponse {
	p := new(UpdateDirectoryServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateDirectoryServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UpdateDirectoryServiceApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/saml-identity-providers/{extId} Put operation
*/
type UpdateSamlIdentityProviderApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSamlIdentityProviderApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSamlIdentityProviderApiResponse() *UpdateSamlIdentityProviderApiResponse {
	p := new(UpdateSamlIdentityProviderApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateSamlIdentityProviderApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UpdateSamlIdentityProviderApiResponse"}
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

/**
REST response for all response codes in api path /iam/v4.0.a1/authn/users/{extId} Put operation
*/
type UpdateUserApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateUserApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateUserApiResponse() *UpdateUserApiResponse {
	p := new(UpdateUserApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UpdateUserApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UpdateUserApiResponse"}
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

/**
Information of the user
*/
type User struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Any additional attribute for the user
	*/
	AdditionalAttributes []import3.KVPair `json:"additionalAttributes,omitempty"`
	/**
	  User or service who created the user
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Creation time of the user
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  Display name for the user
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/**
	  Email Id for the user
	*/
	EmailId *string `json:"emailId,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  First name for the user
	*/
	FirstName *string `json:"firstName,omitempty"`
	/**
	  Flag to force the user to reset password
	*/
	ForceResetPassword *bool `json:"forceResetPassword,omitempty"`
	/**
	  Identifier of the IDP for the user
	*/
	IdpId *string `json:"idpId,omitempty"`
	/**
	  Last successful logged in time for the user
	*/
	LastLoginTime *time.Time `json:"lastLoginTime,omitempty"`
	/**
	  Last name for the user
	*/
	LastName *string `json:"lastName,omitempty"`
	/**
	  Last updated time of the user
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Default locale for the user
	*/
	Locale *string `json:"locale,omitempty"`
	/**
	  Middle name for the user
	*/
	MiddleInitial *string `json:"middleInitial,omitempty"`
	/**
	  Password for the user
	*/
	Password *string `json:"password,omitempty"`
	/**
	  Default region for the user
	*/
	Region *string `json:"region,omitempty"`

	Status *UserStatusType `json:"status,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UserType *UserType `json:"userType"`
	/**
	  Identifier for the user in the form an email address
	*/
	Username *string `json:"username"`
}

func (p *User) MarshalJSON() ([]byte, error) {
	type UserProxy User
	return json.Marshal(struct {
		*UserProxy
		UserType *UserType `json:"userType,omitempty"`
		Username *string   `json:"username,omitempty"`
	}{
		UserProxy: (*UserProxy)(p),
		UserType:  p.UserType,
		Username:  p.Username,
	})
}

func NewUser() *User {
	p := new(User)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.User"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.User"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
User configuration for OpenLDAP directory service
*/
type UserConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Object class in the OpenLDAP system that corresponds to users
	*/
	UserObjectClass *string `json:"userObjectClass"`
	/**
	  Base DN for user search
	*/
	UserSearchBase *string `json:"userSearchBase"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UserConfiguration"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information of the user group
*/
type UserGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  User or service who created the user group
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/**
	  Creation time of the user group
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/**
	  Identifier for the user group in the form of a distinguished name
	*/
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GroupType *GroupType `json:"groupType,omitempty"`
	/**
	  Identifier of the IDP for the user group
	*/
	IdpId *string `json:"idpId,omitempty"`
	/**
	  Last updated time of the user group
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/**
	  Common name of the user group
	*/
	Name *string `json:"name,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewUserGroup() *UserGroup {
	p := new(UserGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.authn.UserGroup"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UserGroup"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
User group configuration for OpenLDAP directory service
*/
type UserGroupConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Attribute in a group that associates users to the group
	*/
	GroupMemberAttribute *string `json:"groupMemberAttribute"`
	/**
	  User attribute value that will be used in group entity to associate user to the group
	*/
	GroupMemberAttributeValue *string `json:"groupMemberAttributeValue"`
	/**
	  Object class in the OpenLDAP system that corresponds to groups
	*/
	GroupObjectClass *string `json:"groupObjectClass"`
	/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UserGroupConfiguration"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "iam.v4.r0.a1.authn.UserStateUpdate"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Status of the user
*/
type UserStatusType int

const (
	USERSTATUSTYPE_UNKNOWN  UserStatusType = 0
	USERSTATUSTYPE_REDACTED UserStatusType = 1
	USERSTATUSTYPE_ACTIVE   UserStatusType = 2
	USERSTATUSTYPE_INACTIVE UserStatusType = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
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

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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
