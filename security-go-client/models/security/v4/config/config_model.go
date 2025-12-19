/*
 * Generated file models/security/v4/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Security APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module security.v4.config of Nutanix Security APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/common/v1/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/prism/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/common"
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/error"
	"time"
)

/*
Access information for the Azure Key Vault.
*/
type AzureAccessInformation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Client identifier for the Azure Key Vault.
	*/
	ClientId *string `json:"clientId"`
	/*
	  Client secret for the Azure Key Vault.
	*/
	ClientSecret *string `json:"clientSecret"`
	/*
	  When the client secret is going to expire.
	*/
	CredentialExpiryDate *time.Time `json:"credentialExpiryDate"`
	/*
	  Endpoint URL for the Azure Key Vault.
	*/
	EndpointUrl *string `json:"endpointUrl"`
	/*
	  Master key identifier for the Azure Key Vault.
	*/
	KeyId *string `json:"keyId"`
	/*
	  Tetant identifier for the Azure Key Vault.
	*/
	TenantId *string `json:"tenantId"`
	/*
	  Truncated client secret for the Azure Key Vault.
	*/
	TruncatedClientSecret *string `json:"truncatedClientSecret,omitempty"`
}

func (p *AzureAccessInformation) MarshalJSON() ([]byte, error) {
	type AzureAccessInformationProxy AzureAccessInformation

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AzureAccessInformationProxy
		ClientId             *string `json:"clientId,omitempty"`
		ClientSecret         *string `json:"clientSecret,omitempty"`
		CredentialExpiryDate string  `json:"credentialExpiryDate,omitempty"`
		EndpointUrl          *string `json:"endpointUrl,omitempty"`
		KeyId                *string `json:"keyId,omitempty"`
		TenantId             *string `json:"tenantId,omitempty"`
	}{
		AzureAccessInformationProxy: (*AzureAccessInformationProxy)(p),
		ClientId:                    p.ClientId,
		ClientSecret:                p.ClientSecret,
		CredentialExpiryDate: func() string {
			if p.CredentialExpiryDate != nil {
				return p.CredentialExpiryDate.Format("2006-01-02")
			} else {
				return ""
			}
		}(),
		EndpointUrl: p.EndpointUrl,
		KeyId:       p.KeyId,
		TenantId:    p.TenantId,
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

func (p *AzureAccessInformation) UnmarshalJSON(b []byte) error {

	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields and custom parsing
	type CustomAzureAccessInformation struct {
		ObjectType_           *string                `json:"$objectType,omitempty"`
		Reserved_             map[string]interface{} `json:"$reserved,omitempty"`
		UnknownFields_        map[string]interface{} `json:"$unknownFields,omitempty"`
		ClientId              *string                `json:"clientId"`
		ClientSecret          *string                `json:"clientSecret"`
		CredentialExpiryDate  string                 `json:"credentialExpiryDate"`
		EndpointUrl           *string                `json:"endpointUrl"`
		KeyId                 *string                `json:"keyId"`
		TenantId              *string                `json:"tenantId"`
		TruncatedClientSecret *string                `json:"truncatedClientSecret,omitempty"`
	}

	var knownFields CustomAzureAccessInformation
	if err := json.Unmarshal(b, &knownFields); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAzureAccessInformation()

	// Handle custom date parsing

	if knownFields.ObjectType_ != nil {
		p.ObjectType_ = knownFields.ObjectType_
	}
	// Handle custom date parsing

	if knownFields.Reserved_ != nil {
		p.Reserved_ = knownFields.Reserved_
	}
	// Handle custom date parsing

	if knownFields.UnknownFields_ != nil {
		p.UnknownFields_ = knownFields.UnknownFields_
	}
	// Handle custom date parsing

	if knownFields.ClientId != nil {
		p.ClientId = knownFields.ClientId
	}
	// Handle custom date parsing

	if knownFields.ClientSecret != nil {
		p.ClientSecret = knownFields.ClientSecret
	}
	// Handle custom date parsing

	// Custom date parsing logic for Date field
	if knownFields.CredentialExpiryDate != "" {
		parsedCredentialExpiryDate, err := time.Parse("2006-01-02", knownFields.CredentialExpiryDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field CredentialExpiryDate in struct AzureAccessInformation: %s", err))
		}
		p.CredentialExpiryDate = &parsedCredentialExpiryDate
	}
	// Handle custom date parsing

	if knownFields.EndpointUrl != nil {
		p.EndpointUrl = knownFields.EndpointUrl
	}
	// Handle custom date parsing

	if knownFields.KeyId != nil {
		p.KeyId = knownFields.KeyId
	}
	// Handle custom date parsing

	if knownFields.TenantId != nil {
		p.TenantId = knownFields.TenantId
	}
	// Handle custom date parsing

	if knownFields.TruncatedClientSecret != nil {
		p.TruncatedClientSecret = knownFields.TruncatedClientSecret
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientId")
	delete(allFields, "clientSecret")
	delete(allFields, "credentialExpiryDate")
	delete(allFields, "endpointUrl")
	delete(allFields, "keyId")
	delete(allFields, "tenantId")
	delete(allFields, "truncatedClientSecret")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAzureAccessInformation() *AzureAccessInformation {
	p := new(AzureAccessInformation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.AzureAccessInformation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BasicAuthCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *import1.BasicAuth `json:"credential"`
}

func (p *BasicAuthCredential) MarshalJSON() ([]byte, error) {
	type BasicAuthCredentialProxy BasicAuthCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BasicAuthCredentialProxy
		Credential *import1.BasicAuth `json:"credential,omitempty"`
	}{
		BasicAuthCredentialProxy: (*BasicAuthCredentialProxy)(p),
		Credential:               p.Credential,
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

func (p *BasicAuthCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BasicAuthCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBasicAuthCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBasicAuthCredential() *BasicAuthCredential {
	p := new(BasicAuthCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.BasicAuthCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BmcCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *import1.BasicAuth `json:"credential"`
	/*
	  Pre-defined type of credential.
	*/
	Type *string `json:"type,omitempty"`
}

func (p *BmcCredential) MarshalJSON() ([]byte, error) {
	type BmcCredentialProxy BmcCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BmcCredentialProxy
		Credential *import1.BasicAuth `json:"credential,omitempty"`
	}{
		BmcCredentialProxy: (*BmcCredentialProxy)(p),
		Credential:         p.Credential,
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

func (p *BmcCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BmcCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBmcCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBmcCredential() *BmcCredential {
	p := new(BmcCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.BmcCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.1/config/credentials Post operation
*/
type CreateCredentialApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateCredentialApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateCredentialApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateCredentialApiResponse

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

func (p *CreateCredentialApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateCredentialApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateCredentialApiResponse()

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

func NewCreateCredentialApiResponse() *CreateCredentialApiResponse {
	p := new(CreateCredentialApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.CreateCredentialApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateCredentialApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateCredentialApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateCredentialApiResponseData()
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
REST response for all response codes in API path /security/v4.1/config/key-management-servers Post operation
*/
type CreateKeyManagementServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateKeyManagementServerApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateKeyManagementServerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateKeyManagementServerApiResponse

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

func (p *CreateKeyManagementServerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateKeyManagementServerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateKeyManagementServerApiResponse()

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

func NewCreateKeyManagementServerApiResponse() *CreateKeyManagementServerApiResponse {
	p := new(CreateKeyManagementServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.CreateKeyManagementServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateKeyManagementServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateKeyManagementServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateKeyManagementServerApiResponseData()
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

type Credential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CredentialDetailsItemDiscriminator_ *string `json:"$credentialDetailsItemDiscriminator,omitempty"`
	/*
	  Details of the credential authentication parameters.
	*/
	CredentialDetails *OneOfCredentialCredentialDetails `json:"credentialDetails"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates if the credential is valid.
	*/
	IsValid *bool `json:"isValid,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the credential.
	*/
	Name *string `json:"name"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Credential) MarshalJSON() ([]byte, error) {
	type CredentialProxy Credential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CredentialProxy
		CredentialDetails *OneOfCredentialCredentialDetails `json:"credentialDetails,omitempty"`
		Name              *string                           `json:"name,omitempty"`
	}{
		CredentialProxy:   (*CredentialProxy)(p),
		CredentialDetails: p.CredentialDetails,
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

func (p *Credential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Credential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CredentialDetailsItemDiscriminator_ != nil {
		p.CredentialDetailsItemDiscriminator_ = known.CredentialDetailsItemDiscriminator_
	}
	if known.CredentialDetails != nil {
		p.CredentialDetails = known.CredentialDetails
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsValid != nil {
		p.IsValid = known.IsValid
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
	delete(allFields, "$credentialDetailsItemDiscriminator")
	delete(allFields, "credentialDetails")
	delete(allFields, "extId")
	delete(allFields, "isValid")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCredential() *Credential {
	p := new(Credential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.Credential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Credential) GetCredentialDetails() interface{} {
	if nil == p.CredentialDetails {
		return nil
	}
	return p.CredentialDetails.GetValue()
}

func (p *Credential) SetCredentialDetails(v interface{}) error {
	if nil == p.CredentialDetails {
		p.CredentialDetails = NewOneOfCredentialCredentialDetails()
	}
	e := p.CredentialDetails.SetValue(v)
	if nil == e {
		if nil == p.CredentialDetailsItemDiscriminator_ {
			p.CredentialDetailsItemDiscriminator_ = new(string)
		}
		*p.CredentialDetailsItemDiscriminator_ = *p.CredentialDetails.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /security/v4.1/config/credentials/{extId} Delete operation
*/
type DeleteCredentialApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteCredentialApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteCredentialApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteCredentialApiResponse

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

func (p *DeleteCredentialApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteCredentialApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteCredentialApiResponse()

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

func NewDeleteCredentialApiResponse() *DeleteCredentialApiResponse {
	p := new(DeleteCredentialApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.DeleteCredentialApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteCredentialApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteCredentialApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteCredentialApiResponseData()
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
REST response for all response codes in API path /security/v4.1/config/key-management-servers/{extId} Delete operation
*/
type DeleteKeyManagementServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteKeyManagementServerApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteKeyManagementServerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteKeyManagementServerApiResponse

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

func (p *DeleteKeyManagementServerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteKeyManagementServerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteKeyManagementServerApiResponse()

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

func NewDeleteKeyManagementServerApiResponse() *DeleteKeyManagementServerApiResponse {
	p := new(DeleteKeyManagementServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.DeleteKeyManagementServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteKeyManagementServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteKeyManagementServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteKeyManagementServerApiResponseData()
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

type EndpointInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPAddressOrFQDN `json:"ipAddress"`
	/*
	  Port on which the External Key Manager server runs.
	*/
	Port *int `json:"port"`
}

func (p *EndpointInfo) MarshalJSON() ([]byte, error) {
	type EndpointInfoProxy EndpointInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EndpointInfoProxy
		IpAddress *import1.IPAddressOrFQDN `json:"ipAddress,omitempty"`
		Port      *int                     `json:"port,omitempty"`
	}{
		EndpointInfoProxy: (*EndpointInfoProxy)(p),
		IpAddress:         p.IpAddress,
		Port:              p.Port,
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

func (p *EndpointInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EndpointInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEndpointInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Port != nil {
		p.Port = known.Port
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "port")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEndpointInfo() *EndpointInfo {
	p := new(EndpointInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.EndpointInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Ssh access config of a particular clusters.
*/
type ExternalSshAccessConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if Ssh is currently enabled on the cluster.
	*/
	IsSshEnabled *bool `json:"isSshEnabled,omitempty"`
	/*
	  Duration in Hours for which Ssh should be enabled. Default is 1 hour.
	*/
	SshEnableDurationHours *int64 `json:"sshEnableDurationHours,omitempty"`
	/*
	  Time when the Ssh will be disabled if Ssh is already enabled.
	*/
	SshExpiryTime *time.Time `json:"sshExpiryTime,omitempty"`
	/*
	  Timestamp when Ssh was last enabled.
	*/
	SshlastEnabledTime *time.Time `json:"sshlastEnabledTime,omitempty"`
}

func (p *ExternalSshAccessConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalSshAccessConfig

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

func (p *ExternalSshAccessConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalSshAccessConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalSshAccessConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsSshEnabled != nil {
		p.IsSshEnabled = known.IsSshEnabled
	}
	if known.SshEnableDurationHours != nil {
		p.SshEnableDurationHours = known.SshEnableDurationHours
	}
	if known.SshExpiryTime != nil {
		p.SshExpiryTime = known.SshExpiryTime
	}
	if known.SshlastEnabledTime != nil {
		p.SshlastEnabledTime = known.SshlastEnabledTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isSshEnabled")
	delete(allFields, "sshEnableDurationHours")
	delete(allFields, "sshExpiryTime")
	delete(allFields, "sshlastEnabledTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalSshAccessConfig() *ExternalSshAccessConfig {
	p := new(ExternalSshAccessConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ExternalSshAccessConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.SshEnableDurationHours = new(int64)
	*p.SshEnableDurationHours = 1

	return p
}

type ExternalSshAccessConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if Ssh is currently enabled on the cluster.
	*/
	IsSshEnabled *bool `json:"isSshEnabled,omitempty"`
	/*
	  Duration in Hours for which Ssh should be enabled. Default is 1 hour.
	*/
	SshEnableDurationHours *int64 `json:"sshEnableDurationHours,omitempty"`
	/*
	  Time when the Ssh will be disabled if Ssh is already enabled.
	*/
	SshExpiryTime *time.Time `json:"sshExpiryTime,omitempty"`
	/*
	  Timestamp when Ssh was last enabled.
	*/
	SshlastEnabledTime *time.Time `json:"sshlastEnabledTime,omitempty"`
}

func (p *ExternalSshAccessConfigProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalSshAccessConfigProjection

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

func (p *ExternalSshAccessConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalSshAccessConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalSshAccessConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsSshEnabled != nil {
		p.IsSshEnabled = known.IsSshEnabled
	}
	if known.SshEnableDurationHours != nil {
		p.SshEnableDurationHours = known.SshEnableDurationHours
	}
	if known.SshExpiryTime != nil {
		p.SshExpiryTime = known.SshExpiryTime
	}
	if known.SshlastEnabledTime != nil {
		p.SshlastEnabledTime = known.SshlastEnabledTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isSshEnabled")
	delete(allFields, "sshEnableDurationHours")
	delete(allFields, "sshExpiryTime")
	delete(allFields, "sshlastEnabledTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalSshAccessConfigProjection() *ExternalSshAccessConfigProjection {
	p := new(ExternalSshAccessConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ExternalSshAccessConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.SshEnableDurationHours = new(int64)
	*p.SshEnableDurationHours = 1

	return p
}

/*
REST response for all response codes in API path /security/v4.1/config/credentials/{extId} Get operation
*/
type GetCredentialApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetCredentialApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetCredentialApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetCredentialApiResponse

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

func (p *GetCredentialApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetCredentialApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetCredentialApiResponse()

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

func NewGetCredentialApiResponse() *GetCredentialApiResponse {
	p := new(GetCredentialApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.GetCredentialApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetCredentialApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetCredentialApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetCredentialApiResponseData()
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
REST response for all response codes in API path /security/v4.1/config/key-management-servers/{extId} Get operation
*/
type GetKeyManagementServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetKeyManagementServerApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetKeyManagementServerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetKeyManagementServerApiResponse

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

func (p *GetKeyManagementServerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetKeyManagementServerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetKeyManagementServerApiResponse()

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

func NewGetKeyManagementServerApiResponse() *GetKeyManagementServerApiResponse {
	p := new(GetKeyManagementServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.GetKeyManagementServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetKeyManagementServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetKeyManagementServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetKeyManagementServerApiResponseData()
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
Type of Intersight connection.
*/
type IntersightConnectionType int

const (
	INTERSIGHTCONNECTIONTYPE_UNKNOWN                      IntersightConnectionType = 0
	INTERSIGHTCONNECTIONTYPE_REDACTED                     IntersightConnectionType = 1
	INTERSIGHTCONNECTIONTYPE_INTERSIGHT_SAAS              IntersightConnectionType = 2
	INTERSIGHTCONNECTIONTYPE_INTERSIGHT_VIRTUAL_APPLIANCE IntersightConnectionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *IntersightConnectionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INTERSIGHT_SAAS",
		"INTERSIGHT_VIRTUAL_APPLIANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e IntersightConnectionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INTERSIGHT_SAAS",
		"INTERSIGHT_VIRTUAL_APPLIANCE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *IntersightConnectionType) index(name string) IntersightConnectionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INTERSIGHT_SAAS",
		"INTERSIGHT_VIRTUAL_APPLIANCE",
	}
	for idx := range names {
		if names[idx] == name {
			return IntersightConnectionType(idx)
		}
	}
	return INTERSIGHTCONNECTIONTYPE_UNKNOWN
}

func (e *IntersightConnectionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IntersightConnectionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IntersightConnectionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IntersightConnectionType) Ref() *IntersightConnectionType {
	return &e
}

type IntersightCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *KeyBasedAuth `json:"credential"`

	DeploymentType *IntersightConnectionType `json:"deploymentType"`
	/*
	  Pre-defined type of credential.
	*/
	Type *string `json:"type,omitempty"`
	/*
	  Intersight connection url
	*/
	Url *string `json:"url"`
}

func (p *IntersightCredential) MarshalJSON() ([]byte, error) {
	type IntersightCredentialProxy IntersightCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IntersightCredentialProxy
		Credential     *KeyBasedAuth             `json:"credential,omitempty"`
		DeploymentType *IntersightConnectionType `json:"deploymentType,omitempty"`
		Url            *string                   `json:"url,omitempty"`
	}{
		IntersightCredentialProxy: (*IntersightCredentialProxy)(p),
		Credential:                p.Credential,
		DeploymentType:            p.DeploymentType,
		Url:                       p.Url,
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

func (p *IntersightCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IntersightCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewIntersightCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}
	if known.DeploymentType != nil {
		p.DeploymentType = known.DeploymentType
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Url != nil {
		p.Url = known.Url
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")
	delete(allFields, "deploymentType")
	delete(allFields, "type")
	delete(allFields, "url")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIntersightCredential() *IntersightCredential {
	p := new(IntersightCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.IntersightCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Secret fields of a Key based Authencation credential.
*/
type KeyBasedAuth struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Intersight connection API key.
	*/
	ApiKey *string `json:"apiKey"`
	/*
	  Secret key of the Intersight connection.
	*/
	SecretKey *string `json:"secretKey"`
}

func (p *KeyBasedAuth) MarshalJSON() ([]byte, error) {
	type KeyBasedAuthProxy KeyBasedAuth

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*KeyBasedAuthProxy
		ApiKey    *string `json:"apiKey,omitempty"`
		SecretKey *string `json:"secretKey,omitempty"`
	}{
		KeyBasedAuthProxy: (*KeyBasedAuthProxy)(p),
		ApiKey:            p.ApiKey,
		SecretKey:         p.SecretKey,
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

func (p *KeyBasedAuth) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KeyBasedAuth
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKeyBasedAuth()

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
	if known.SecretKey != nil {
		p.SecretKey = known.SecretKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiKey")
	delete(allFields, "secretKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewKeyBasedAuth() *KeyBasedAuth {
	p := new(KeyBasedAuth)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KeyBasedAuth"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type KeyBasedAuthCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *KeyBasedAuth `json:"credential"`
}

func (p *KeyBasedAuthCredential) MarshalJSON() ([]byte, error) {
	type KeyBasedAuthCredentialProxy KeyBasedAuthCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*KeyBasedAuthCredentialProxy
		Credential *KeyBasedAuth `json:"credential,omitempty"`
	}{
		KeyBasedAuthCredentialProxy: (*KeyBasedAuthCredentialProxy)(p),
		Credential:                  p.Credential,
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

func (p *KeyBasedAuthCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KeyBasedAuthCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKeyBasedAuthCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewKeyBasedAuthCredential() *KeyBasedAuthCredential {
	p := new(KeyBasedAuthCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KeyBasedAuthCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Key Management Server.
*/
type KeyManagementServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	AccessInformationItemDiscriminator_ *string `json:"$accessInformationItemDiscriminator,omitempty"`
	/*
	  KMS access information.
	*/
	AccessInformation *OneOfKeyManagementServerAccessInformation `json:"accessInformation,omitempty"`
	/*
	  Timestamp of KMS Creation.
	*/
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Name of the key management server (KMS).
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *KeyManagementServer) MarshalJSON() ([]byte, error) {
	type KeyManagementServerProxy KeyManagementServer

	// Step 1: Marshal known fields via proxy to accommodate fields with date format
	baseStruct := struct {
		*KeyManagementServerProxy
		CreationTimestamp string `json:"creationTimestamp,omitempty"`
	}{
		KeyManagementServerProxy: (*KeyManagementServerProxy)(p),
		CreationTimestamp: func() string {
			if p.CreationTimestamp != nil {
				return p.CreationTimestamp.Format("2006-01-02")
			} else {
				return ""
			}
		}(),
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

func (p *KeyManagementServer) UnmarshalJSON(b []byte) error {

	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields and custom parsing
	type CustomKeyManagementServer struct {
		ObjectType_                         *string                                    `json:"$objectType,omitempty"`
		Reserved_                           map[string]interface{}                     `json:"$reserved,omitempty"`
		UnknownFields_                      map[string]interface{}                     `json:"$unknownFields,omitempty"`
		AccessInformationItemDiscriminator_ *string                                    `json:"$accessInformationItemDiscriminator,omitempty"`
		AccessInformation                   *OneOfKeyManagementServerAccessInformation `json:"accessInformation,omitempty"`
		CreationTimestamp                   string                                     `json:"creationTimestamp,omitempty"`
		ExtId                               *string                                    `json:"extId,omitempty"`
		Links                               []import4.ApiLink                          `json:"links,omitempty"`
		Name                                *string                                    `json:"name,omitempty"`
		TenantId                            *string                                    `json:"tenantId,omitempty"`
	}

	var knownFields CustomKeyManagementServer
	if err := json.Unmarshal(b, &knownFields); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKeyManagementServer()

	// Handle custom date parsing

	if knownFields.ObjectType_ != nil {
		p.ObjectType_ = knownFields.ObjectType_
	}
	// Handle custom date parsing

	if knownFields.Reserved_ != nil {
		p.Reserved_ = knownFields.Reserved_
	}
	// Handle custom date parsing

	if knownFields.UnknownFields_ != nil {
		p.UnknownFields_ = knownFields.UnknownFields_
	}
	// Handle custom date parsing

	if knownFields.AccessInformationItemDiscriminator_ != nil {
		p.AccessInformationItemDiscriminator_ = knownFields.AccessInformationItemDiscriminator_
	}
	// Handle custom date parsing

	if knownFields.AccessInformation != nil {
		p.AccessInformation = knownFields.AccessInformation
	}
	// Handle custom date parsing

	// Custom date parsing logic for Date field
	if knownFields.CreationTimestamp != "" {
		parsedCreationTimestamp, err := time.Parse("2006-01-02", knownFields.CreationTimestamp)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field CreationTimestamp in struct KeyManagementServer: %s", err))
		}
		p.CreationTimestamp = &parsedCreationTimestamp
	}
	// Handle custom date parsing

	if knownFields.ExtId != nil {
		p.ExtId = knownFields.ExtId
	}
	// Handle custom date parsing

	if knownFields.Links != nil {
		p.Links = knownFields.Links
	}
	// Handle custom date parsing

	if knownFields.Name != nil {
		p.Name = knownFields.Name
	}
	// Handle custom date parsing

	if knownFields.TenantId != nil {
		p.TenantId = knownFields.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$accessInformationItemDiscriminator")
	delete(allFields, "accessInformation")
	delete(allFields, "creationTimestamp")
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

func NewKeyManagementServer() *KeyManagementServer {
	p := new(KeyManagementServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KeyManagementServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *KeyManagementServer) GetAccessInformation() interface{} {
	if nil == p.AccessInformation {
		return nil
	}
	return p.AccessInformation.GetValue()
}

func (p *KeyManagementServer) SetAccessInformation(v interface{}) error {
	if nil == p.AccessInformation {
		p.AccessInformation = NewOneOfKeyManagementServerAccessInformation()
	}
	e := p.AccessInformation.SetValue(v)
	if nil == e {
		if nil == p.AccessInformationItemDiscriminator_ {
			p.AccessInformationItemDiscriminator_ = new(string)
		}
		*p.AccessInformationItemDiscriminator_ = *p.AccessInformation.Discriminator
	}
	return e
}

/*
Access info for KMIP based External Key Manager.
*/
type KmipAccessInformation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the Certificate Authority.
	*/
	CaName *string `json:"caName"`
	/*
	  CA PEM file.
	*/
	CaPem *string `json:"caPem"`
	/*
	  Cert PEM File.
	*/
	CertPem *string `json:"certPem"`
	/*
	  List of endpoints of the External Key Manager server.
	*/
	Endpoints []EndpointInfo `json:"endpoints"`
	/*
	  Private Key used by the External Key Manager.
	*/
	PrivateKey *string `json:"privateKey"`
}

func (p *KmipAccessInformation) MarshalJSON() ([]byte, error) {
	type KmipAccessInformationProxy KmipAccessInformation

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*KmipAccessInformationProxy
		CaName     *string        `json:"caName,omitempty"`
		CaPem      *string        `json:"caPem,omitempty"`
		CertPem    *string        `json:"certPem,omitempty"`
		Endpoints  []EndpointInfo `json:"endpoints,omitempty"`
		PrivateKey *string        `json:"privateKey,omitempty"`
	}{
		KmipAccessInformationProxy: (*KmipAccessInformationProxy)(p),
		CaName:                     p.CaName,
		CaPem:                      p.CaPem,
		CertPem:                    p.CertPem,
		Endpoints:                  p.Endpoints,
		PrivateKey:                 p.PrivateKey,
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

func (p *KmipAccessInformation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KmipAccessInformation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewKmipAccessInformation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CaName != nil {
		p.CaName = known.CaName
	}
	if known.CaPem != nil {
		p.CaPem = known.CaPem
	}
	if known.CertPem != nil {
		p.CertPem = known.CertPem
	}
	if known.Endpoints != nil {
		p.Endpoints = known.Endpoints
	}
	if known.PrivateKey != nil {
		p.PrivateKey = known.PrivateKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "caName")
	delete(allFields, "caPem")
	delete(allFields, "certPem")
	delete(allFields, "endpoints")
	delete(allFields, "privateKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewKmipAccessInformation() *KmipAccessInformation {
	p := new(KmipAccessInformation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KmipAccessInformation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.1/config/credentials Get operation
*/
type ListCredentialsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListCredentialsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListCredentialsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListCredentialsApiResponse

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

func (p *ListCredentialsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListCredentialsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListCredentialsApiResponse()

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

func NewListCredentialsApiResponse() *ListCredentialsApiResponse {
	p := new(ListCredentialsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ListCredentialsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListCredentialsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListCredentialsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListCredentialsApiResponseData()
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
REST response for all response codes in API path /security/v4.1/config/key-management-servers Get operation
*/
type ListKeyManagementServersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListKeyManagementServersApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListKeyManagementServersApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListKeyManagementServersApiResponse

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

func (p *ListKeyManagementServersApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListKeyManagementServersApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListKeyManagementServersApiResponse()

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

func NewListKeyManagementServersApiResponse() *ListKeyManagementServersApiResponse {
	p := new(ListKeyManagementServersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.ListKeyManagementServersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListKeyManagementServersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListKeyManagementServersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListKeyManagementServersApiResponseData()
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
Contains possible values for scheduling a task.
*/
type ScheduleType int

const (
	SCHEDULETYPE_UNKNOWN  ScheduleType = 0
	SCHEDULETYPE_REDACTED ScheduleType = 1
	SCHEDULETYPE_HOURLY   ScheduleType = 2
	SCHEDULETYPE_DAILY    ScheduleType = 3
	SCHEDULETYPE_WEEKLY   ScheduleType = 4
	SCHEDULETYPE_MONTHLY  ScheduleType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScheduleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScheduleType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScheduleType) index(name string) ScheduleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HOURLY",
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
	for idx := range names {
		if names[idx] == name {
			return ScheduleType(idx)
		}
	}
	return SCHEDULETYPE_UNKNOWN
}

func (e *ScheduleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScheduleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScheduleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScheduleType) Ref() *ScheduleType {
	return &e
}

/*
Contains the status of all the security configurations settings for a cluster.
*/
type SecurityConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvScmaSchedule *ScheduleType `json:"ahvScmaSchedule,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterScmaSchedule *ScheduleType `json:"clusterScmaSchedule,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the AHV defence knowledge consent banner is enabled on the hypervisor.
	*/
	IsAhvDefenseConsentBannerEnabled *bool `json:"isAhvDefenseConsentBannerEnabled,omitempty"`
	/*
	  Indicates whether the aide service is enabled on a cluster.
	*/
	IsAideEnabled *bool `json:"isAideEnabled,omitempty"`
	/*
	  Indicates whether the Nutanix CVM defence knowledge consent banner is enabled.
	*/
	IsClusterDefenseConsentBannerEnabled *bool `json:"isClusterDefenseConsentBannerEnabled,omitempty"`
	/*
	  Indicates whether the cluster lockdown mode is enabled on a cluster.
	*/
	IsClusterLockdownEnabled *bool `json:"isClusterLockdownEnabled,omitempty"`
	/*
	  Indicates whether the high strength password is enabled on a cluster.
	*/
	IsHighStrengthPasswordEnabled *bool `json:"isHighStrengthPasswordEnabled,omitempty"`
	/*
	  Indicates whether the log forwarding is enabled on a cluster.
	*/
	IsLogForwardingEnabled *bool `json:"isLogForwardingEnabled,omitempty"`
	/*
	  Indicates whether the host secure start is enabled on a cluster.
	*/
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SecurityConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SecurityConfig

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

func (p *SecurityConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SecurityConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSecurityConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AhvScmaSchedule != nil {
		p.AhvScmaSchedule = known.AhvScmaSchedule
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ClusterScmaSchedule != nil {
		p.ClusterScmaSchedule = known.ClusterScmaSchedule
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsAhvDefenseConsentBannerEnabled != nil {
		p.IsAhvDefenseConsentBannerEnabled = known.IsAhvDefenseConsentBannerEnabled
	}
	if known.IsAideEnabled != nil {
		p.IsAideEnabled = known.IsAideEnabled
	}
	if known.IsClusterDefenseConsentBannerEnabled != nil {
		p.IsClusterDefenseConsentBannerEnabled = known.IsClusterDefenseConsentBannerEnabled
	}
	if known.IsClusterLockdownEnabled != nil {
		p.IsClusterLockdownEnabled = known.IsClusterLockdownEnabled
	}
	if known.IsHighStrengthPasswordEnabled != nil {
		p.IsHighStrengthPasswordEnabled = known.IsHighStrengthPasswordEnabled
	}
	if known.IsLogForwardingEnabled != nil {
		p.IsLogForwardingEnabled = known.IsLogForwardingEnabled
	}
	if known.IsSecureBootEnabled != nil {
		p.IsSecureBootEnabled = known.IsSecureBootEnabled
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
	delete(allFields, "ahvScmaSchedule")
	delete(allFields, "clusterExtId")
	delete(allFields, "clusterScmaSchedule")
	delete(allFields, "extId")
	delete(allFields, "isAhvDefenseConsentBannerEnabled")
	delete(allFields, "isAideEnabled")
	delete(allFields, "isClusterDefenseConsentBannerEnabled")
	delete(allFields, "isClusterLockdownEnabled")
	delete(allFields, "isHighStrengthPasswordEnabled")
	delete(allFields, "isLogForwardingEnabled")
	delete(allFields, "isSecureBootEnabled")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSecurityConfig() *SecurityConfig {
	p := new(SecurityConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SecurityConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SecurityConfigProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvScmaSchedule *ScheduleType `json:"ahvScmaSchedule,omitempty"`
	/*
	  UUID of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`

	ClusterProjection *import5.ClusterProjection `json:"clusterProjection,omitempty"`

	ClusterScmaSchedule *ScheduleType `json:"clusterScmaSchedule,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates whether the AHV defence knowledge consent banner is enabled on the hypervisor.
	*/
	IsAhvDefenseConsentBannerEnabled *bool `json:"isAhvDefenseConsentBannerEnabled,omitempty"`
	/*
	  Indicates whether the aide service is enabled on a cluster.
	*/
	IsAideEnabled *bool `json:"isAideEnabled,omitempty"`
	/*
	  Indicates whether the Nutanix CVM defence knowledge consent banner is enabled.
	*/
	IsClusterDefenseConsentBannerEnabled *bool `json:"isClusterDefenseConsentBannerEnabled,omitempty"`
	/*
	  Indicates whether the cluster lockdown mode is enabled on a cluster.
	*/
	IsClusterLockdownEnabled *bool `json:"isClusterLockdownEnabled,omitempty"`
	/*
	  Indicates whether the high strength password is enabled on a cluster.
	*/
	IsHighStrengthPasswordEnabled *bool `json:"isHighStrengthPasswordEnabled,omitempty"`
	/*
	  Indicates whether the log forwarding is enabled on a cluster.
	*/
	IsLogForwardingEnabled *bool `json:"isLogForwardingEnabled,omitempty"`
	/*
	  Indicates whether the host secure start is enabled on a cluster.
	*/
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SecurityConfigProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SecurityConfigProjection

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

func (p *SecurityConfigProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SecurityConfigProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSecurityConfigProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AhvScmaSchedule != nil {
		p.AhvScmaSchedule = known.AhvScmaSchedule
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.ClusterProjection != nil {
		p.ClusterProjection = known.ClusterProjection
	}
	if known.ClusterScmaSchedule != nil {
		p.ClusterScmaSchedule = known.ClusterScmaSchedule
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsAhvDefenseConsentBannerEnabled != nil {
		p.IsAhvDefenseConsentBannerEnabled = known.IsAhvDefenseConsentBannerEnabled
	}
	if known.IsAideEnabled != nil {
		p.IsAideEnabled = known.IsAideEnabled
	}
	if known.IsClusterDefenseConsentBannerEnabled != nil {
		p.IsClusterDefenseConsentBannerEnabled = known.IsClusterDefenseConsentBannerEnabled
	}
	if known.IsClusterLockdownEnabled != nil {
		p.IsClusterLockdownEnabled = known.IsClusterLockdownEnabled
	}
	if known.IsHighStrengthPasswordEnabled != nil {
		p.IsHighStrengthPasswordEnabled = known.IsHighStrengthPasswordEnabled
	}
	if known.IsLogForwardingEnabled != nil {
		p.IsLogForwardingEnabled = known.IsLogForwardingEnabled
	}
	if known.IsSecureBootEnabled != nil {
		p.IsSecureBootEnabled = known.IsSecureBootEnabled
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
	delete(allFields, "ahvScmaSchedule")
	delete(allFields, "clusterExtId")
	delete(allFields, "clusterProjection")
	delete(allFields, "clusterScmaSchedule")
	delete(allFields, "extId")
	delete(allFields, "isAhvDefenseConsentBannerEnabled")
	delete(allFields, "isAideEnabled")
	delete(allFields, "isClusterDefenseConsentBannerEnabled")
	delete(allFields, "isClusterLockdownEnabled")
	delete(allFields, "isHighStrengthPasswordEnabled")
	delete(allFields, "isLogForwardingEnabled")
	delete(allFields, "isSecureBootEnabled")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSecurityConfigProjection() *SecurityConfigProjection {
	p := new(SecurityConfigProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SecurityConfigProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains the configuration for PC UI visibility status of all security configurations settings.
*/
type SecurityConfigVisibilitySetting struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Represents the visibility of aide service state.
	*/
	IsAideVisible *bool `json:"isAideVisible,omitempty"`
	/*
	  Represents the visibility of cluster lockdown state.
	*/
	IsClusterLockdownVisible *bool `json:"isClusterLockdownVisible,omitempty"`
	/*
	  Represents the visibility of banner state.
	*/
	IsDefenseConsentBannerVisible *bool `json:"isDefenseConsentBannerVisible,omitempty"`
	/*
	  Represents the visibility of high strength password state.
	*/
	IsHighStrengthPasswordVisible *bool `json:"isHighStrengthPasswordVisible,omitempty"`
	/*
	  Represents the visibility of log forwarding state.
	*/
	IsLogForwardingVisible *bool `json:"isLogForwardingVisible,omitempty"`
	/*
	  Represents the visibility of network segmentation state.
	*/
	IsNetworkSegmentationVisible *bool `json:"isNetworkSegmentationVisible,omitempty"`
	/*
	  Represents the visibility of security configuration management automation state.
	*/
	IsScmaVisible *bool `json:"isScmaVisible,omitempty"`
	/*
	  Represents the visibility of host secure boot state.
	*/
	IsSecureBootVisible *bool `json:"isSecureBootVisible,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SecurityConfigVisibilitySetting) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SecurityConfigVisibilitySetting

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

func (p *SecurityConfigVisibilitySetting) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SecurityConfigVisibilitySetting
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSecurityConfigVisibilitySetting()

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
	if known.IsAideVisible != nil {
		p.IsAideVisible = known.IsAideVisible
	}
	if known.IsClusterLockdownVisible != nil {
		p.IsClusterLockdownVisible = known.IsClusterLockdownVisible
	}
	if known.IsDefenseConsentBannerVisible != nil {
		p.IsDefenseConsentBannerVisible = known.IsDefenseConsentBannerVisible
	}
	if known.IsHighStrengthPasswordVisible != nil {
		p.IsHighStrengthPasswordVisible = known.IsHighStrengthPasswordVisible
	}
	if known.IsLogForwardingVisible != nil {
		p.IsLogForwardingVisible = known.IsLogForwardingVisible
	}
	if known.IsNetworkSegmentationVisible != nil {
		p.IsNetworkSegmentationVisible = known.IsNetworkSegmentationVisible
	}
	if known.IsScmaVisible != nil {
		p.IsScmaVisible = known.IsScmaVisible
	}
	if known.IsSecureBootVisible != nil {
		p.IsSecureBootVisible = known.IsSecureBootVisible
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
	delete(allFields, "extId")
	delete(allFields, "isAideVisible")
	delete(allFields, "isClusterLockdownVisible")
	delete(allFields, "isDefenseConsentBannerVisible")
	delete(allFields, "isHighStrengthPasswordVisible")
	delete(allFields, "isLogForwardingVisible")
	delete(allFields, "isNetworkSegmentationVisible")
	delete(allFields, "isScmaVisible")
	delete(allFields, "isSecureBootVisible")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSecurityConfigVisibilitySetting() *SecurityConfigVisibilitySetting {
	p := new(SecurityConfigVisibilitySetting)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SecurityConfigVisibilitySetting"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SecurityConfigVisibilitySettingProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Represents the visibility of aide service state.
	*/
	IsAideVisible *bool `json:"isAideVisible,omitempty"`
	/*
	  Represents the visibility of cluster lockdown state.
	*/
	IsClusterLockdownVisible *bool `json:"isClusterLockdownVisible,omitempty"`
	/*
	  Represents the visibility of banner state.
	*/
	IsDefenseConsentBannerVisible *bool `json:"isDefenseConsentBannerVisible,omitempty"`
	/*
	  Represents the visibility of high strength password state.
	*/
	IsHighStrengthPasswordVisible *bool `json:"isHighStrengthPasswordVisible,omitempty"`
	/*
	  Represents the visibility of log forwarding state.
	*/
	IsLogForwardingVisible *bool `json:"isLogForwardingVisible,omitempty"`
	/*
	  Represents the visibility of network segmentation state.
	*/
	IsNetworkSegmentationVisible *bool `json:"isNetworkSegmentationVisible,omitempty"`
	/*
	  Represents the visibility of security configuration management automation state.
	*/
	IsScmaVisible *bool `json:"isScmaVisible,omitempty"`
	/*
	  Represents the visibility of host secure boot state.
	*/
	IsSecureBootVisible *bool `json:"isSecureBootVisible,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *SecurityConfigVisibilitySettingProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SecurityConfigVisibilitySettingProjection

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

func (p *SecurityConfigVisibilitySettingProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SecurityConfigVisibilitySettingProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSecurityConfigVisibilitySettingProjection()

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
	if known.IsAideVisible != nil {
		p.IsAideVisible = known.IsAideVisible
	}
	if known.IsClusterLockdownVisible != nil {
		p.IsClusterLockdownVisible = known.IsClusterLockdownVisible
	}
	if known.IsDefenseConsentBannerVisible != nil {
		p.IsDefenseConsentBannerVisible = known.IsDefenseConsentBannerVisible
	}
	if known.IsHighStrengthPasswordVisible != nil {
		p.IsHighStrengthPasswordVisible = known.IsHighStrengthPasswordVisible
	}
	if known.IsLogForwardingVisible != nil {
		p.IsLogForwardingVisible = known.IsLogForwardingVisible
	}
	if known.IsNetworkSegmentationVisible != nil {
		p.IsNetworkSegmentationVisible = known.IsNetworkSegmentationVisible
	}
	if known.IsScmaVisible != nil {
		p.IsScmaVisible = known.IsScmaVisible
	}
	if known.IsSecureBootVisible != nil {
		p.IsSecureBootVisible = known.IsSecureBootVisible
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
	delete(allFields, "extId")
	delete(allFields, "isAideVisible")
	delete(allFields, "isClusterLockdownVisible")
	delete(allFields, "isDefenseConsentBannerVisible")
	delete(allFields, "isHighStrengthPasswordVisible")
	delete(allFields, "isLogForwardingVisible")
	delete(allFields, "isNetworkSegmentationVisible")
	delete(allFields, "isScmaVisible")
	delete(allFields, "isSecureBootVisible")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSecurityConfigVisibilitySettingProjection() *SecurityConfigVisibilitySettingProjection {
	p := new(SecurityConfigVisibilitySettingProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.SecurityConfigVisibilitySettingProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /security/v4.1/config/credentials/{extId} Put operation
*/
type UpdateCredentialApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCredentialApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateCredentialApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateCredentialApiResponse

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

func (p *UpdateCredentialApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateCredentialApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateCredentialApiResponse()

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

func NewUpdateCredentialApiResponse() *UpdateCredentialApiResponse {
	p := new(UpdateCredentialApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.UpdateCredentialApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateCredentialApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateCredentialApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateCredentialApiResponseData()
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
REST response for all response codes in API path /security/v4.1/config/key-management-servers/{extId} Put operation
*/
type UpdateKeyManagementServerApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateKeyManagementServerApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateKeyManagementServerApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateKeyManagementServerApiResponse

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

func (p *UpdateKeyManagementServerApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateKeyManagementServerApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateKeyManagementServerApiResponse()

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

func NewUpdateKeyManagementServerApiResponse() *UpdateKeyManagementServerApiResponse {
	p := new(UpdateKeyManagementServerApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.UpdateKeyManagementServerApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateKeyManagementServerApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateKeyManagementServerApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateKeyManagementServerApiResponseData()
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

type VcenterCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address"`

	Credential *import1.BasicAuth `json:"credential"`
	/*
	  Pre-defined type of credential.
	*/
	Type *string `json:"type,omitempty"`
}

func (p *VcenterCredential) MarshalJSON() ([]byte, error) {
	type VcenterCredentialProxy VcenterCredential

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VcenterCredentialProxy
		Address    *import1.IPAddressOrFQDN `json:"address,omitempty"`
		Credential *import1.BasicAuth       `json:"credential,omitempty"`
	}{
		VcenterCredentialProxy: (*VcenterCredentialProxy)(p),
		Address:                p.Address,
		Credential:             p.Credential,
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

func (p *VcenterCredential) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VcenterCredential
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVcenterCredential()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Address != nil {
		p.Address = known.Address
	}
	if known.Credential != nil {
		p.Credential = known.Credential
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "address")
	delete(allFields, "credential")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVcenterCredential() *VcenterCredential {
	p := new(VcenterCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.VcenterCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfDeleteCredentialApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1010 *interface{}           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteCredentialApiResponseData() *OneOfDeleteCredentialApiResponseData {
	p := new(OneOfDeleteCredentialApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteCredentialApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteCredentialApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(interface{})
		}
		*p.oneOfType1010 = nil
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

func (p *OneOfDeleteCredentialApiResponseData) GetValue() interface{} {
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType1010
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteCredentialApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1010 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if nil == *vOneOfType1010 {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(interface{})
			}
			*p.oneOfType1010 = nil
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
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteCredentialApiResponseData"))
}

func (p *OneOfDeleteCredentialApiResponseData) MarshalJSON() ([]byte, error) {
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteCredentialApiResponseData")
}

type OneOfDeleteKeyManagementServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfDeleteKeyManagementServerApiResponseData() *OneOfDeleteKeyManagementServerApiResponseData {
	p := new(OneOfDeleteKeyManagementServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteKeyManagementServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteKeyManagementServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfDeleteKeyManagementServerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteKeyManagementServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteKeyManagementServerApiResponseData"))
}

func (p *OneOfDeleteKeyManagementServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteKeyManagementServerApiResponseData")
}

type OneOfListCredentialsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType3001 []Credential           `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfListCredentialsApiResponseData() *OneOfListCredentialsApiResponseData {
	p := new(OneOfListCredentialsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListCredentialsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListCredentialsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Credential:
		p.oneOfType3001 = v.([]Credential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.config.Credential>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.config.Credential>"
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

func (p *OneOfListCredentialsApiResponseData) GetValue() interface{} {
	if "List<security.v4.config.Credential>" == *p.Discriminator {
		return p.oneOfType3001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListCredentialsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType3001 := new([]Credential)
	if err := json.Unmarshal(b, vOneOfType3001); err == nil {
		if len(*vOneOfType3001) == 0 || "security.v4.config.Credential" == *((*vOneOfType3001)[0].ObjectType_) {
			p.oneOfType3001 = *vOneOfType3001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.config.Credential>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.config.Credential>"
			return nil
		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListCredentialsApiResponseData"))
}

func (p *OneOfListCredentialsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<security.v4.config.Credential>" == *p.Discriminator {
		return json.Marshal(p.oneOfType3001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListCredentialsApiResponseData")
}

type OneOfListKeyManagementServersApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []KeyManagementServer  `json:"-"`
}

func NewOneOfListKeyManagementServersApiResponseData() *OneOfListKeyManagementServersApiResponseData {
	p := new(OneOfListKeyManagementServersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListKeyManagementServersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListKeyManagementServersApiResponseData is nil"))
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
	case []KeyManagementServer:
		p.oneOfType2001 = v.([]KeyManagementServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<security.v4.config.KeyManagementServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<security.v4.config.KeyManagementServer>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListKeyManagementServersApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<security.v4.config.KeyManagementServer>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListKeyManagementServersApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]KeyManagementServer)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "security.v4.config.KeyManagementServer" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<security.v4.config.KeyManagementServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<security.v4.config.KeyManagementServer>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListKeyManagementServersApiResponseData"))
}

func (p *OneOfListKeyManagementServersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<security.v4.config.KeyManagementServer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListKeyManagementServersApiResponseData")
}

type OneOfUpdateCredentialApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1010 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateCredentialApiResponseData() *OneOfUpdateCredentialApiResponseData {
	p := new(OneOfUpdateCredentialApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateCredentialApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateCredentialApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(import2.TaskReference)
		}
		*p.oneOfType1010 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1010.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1010.ObjectType_
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

func (p *OneOfUpdateCredentialApiResponseData) GetValue() interface{} {
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1010
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateCredentialApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1010 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(import2.TaskReference)
			}
			*p.oneOfType1010 = *vOneOfType1010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1010.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1010.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCredentialApiResponseData"))
}

func (p *OneOfUpdateCredentialApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCredentialApiResponseData")
}

type OneOfGetKeyManagementServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *KeyManagementServer   `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfGetKeyManagementServerApiResponseData() *OneOfGetKeyManagementServerApiResponseData {
	p := new(OneOfGetKeyManagementServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetKeyManagementServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetKeyManagementServerApiResponseData is nil"))
	}
	switch v.(type) {
	case KeyManagementServer:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(KeyManagementServer)
		}
		*p.oneOfType2001 = v.(KeyManagementServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfGetKeyManagementServerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetKeyManagementServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(KeyManagementServer)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "security.v4.config.KeyManagementServer" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(KeyManagementServer)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetKeyManagementServerApiResponseData"))
}

func (p *OneOfGetKeyManagementServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetKeyManagementServerApiResponseData")
}

type OneOfCreateCredentialApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType1010 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateCredentialApiResponseData() *OneOfCreateCredentialApiResponseData {
	p := new(OneOfCreateCredentialApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateCredentialApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateCredentialApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(import2.TaskReference)
		}
		*p.oneOfType1010 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1010.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1010.ObjectType_
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

func (p *OneOfCreateCredentialApiResponseData) GetValue() interface{} {
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1010
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateCredentialApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType1010 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType1010.ObjectType_ {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(import2.TaskReference)
			}
			*p.oneOfType1010 = *vOneOfType1010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1010.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1010.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateCredentialApiResponseData"))
}

func (p *OneOfCreateCredentialApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateCredentialApiResponseData")
}

type OneOfGetCredentialApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType1010 *Credential            `json:"-"`
}

func NewOneOfGetCredentialApiResponseData() *OneOfGetCredentialApiResponseData {
	p := new(OneOfGetCredentialApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetCredentialApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetCredentialApiResponseData is nil"))
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
	case Credential:
		if nil == p.oneOfType1010 {
			p.oneOfType1010 = new(Credential)
		}
		*p.oneOfType1010 = v.(Credential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1010.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1010.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetCredentialApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1010
	}
	return nil
}

func (p *OneOfGetCredentialApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType1010 := new(Credential)
	if err := json.Unmarshal(b, vOneOfType1010); err == nil {
		if "security.v4.config.Credential" == *vOneOfType1010.ObjectType_ {
			if nil == p.oneOfType1010 {
				p.oneOfType1010 = new(Credential)
			}
			*p.oneOfType1010 = *vOneOfType1010
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1010.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1010.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetCredentialApiResponseData"))
}

func (p *OneOfGetCredentialApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType1010 != nil && *p.oneOfType1010.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1010)
	}
	return nil, errors.New("No value to marshal for OneOfGetCredentialApiResponseData")
}

type OneOfCreateKeyManagementServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfCreateKeyManagementServerApiResponseData() *OneOfCreateKeyManagementServerApiResponseData {
	p := new(OneOfCreateKeyManagementServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateKeyManagementServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateKeyManagementServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfCreateKeyManagementServerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateKeyManagementServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateKeyManagementServerApiResponseData"))
}

func (p *OneOfCreateKeyManagementServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateKeyManagementServerApiResponseData")
}

type OneOfKeyManagementServerAccessInformation struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType1003 *AzureAccessInformation `json:"-"`
	oneOfType1106 *KmipAccessInformation  `json:"-"`
}

func NewOneOfKeyManagementServerAccessInformation() *OneOfKeyManagementServerAccessInformation {
	p := new(OneOfKeyManagementServerAccessInformation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfKeyManagementServerAccessInformation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfKeyManagementServerAccessInformation is nil"))
	}
	switch v.(type) {
	case AzureAccessInformation:
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(AzureAccessInformation)
		}
		*p.oneOfType1003 = v.(AzureAccessInformation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1003.ObjectType_
	case KmipAccessInformation:
		if nil == p.oneOfType1106 {
			p.oneOfType1106 = new(KmipAccessInformation)
		}
		*p.oneOfType1106 = v.(KmipAccessInformation)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1106.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1106.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfKeyManagementServerAccessInformation) GetValue() interface{} {
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1003
	}
	if p.oneOfType1106 != nil && *p.oneOfType1106.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1106
	}
	return nil
}

func (p *OneOfKeyManagementServerAccessInformation) UnmarshalJSON(b []byte) error {
	vOneOfType1003 := new(AzureAccessInformation)
	if err := json.Unmarshal(b, vOneOfType1003); err == nil {
		if "security.v4.config.AzureAccessInformation" == *vOneOfType1003.ObjectType_ {
			if nil == p.oneOfType1003 {
				p.oneOfType1003 = new(AzureAccessInformation)
			}
			*p.oneOfType1003 = *vOneOfType1003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1003.ObjectType_
			return nil
		}
	}
	vOneOfType1106 := new(KmipAccessInformation)
	if err := json.Unmarshal(b, vOneOfType1106); err == nil {
		if "security.v4.config.KmipAccessInformation" == *vOneOfType1106.ObjectType_ {
			if nil == p.oneOfType1106 {
				p.oneOfType1106 = new(KmipAccessInformation)
			}
			*p.oneOfType1106 = *vOneOfType1106
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1106.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1106.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfKeyManagementServerAccessInformation"))
}

func (p *OneOfKeyManagementServerAccessInformation) MarshalJSON() ([]byte, error) {
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1003)
	}
	if p.oneOfType1106 != nil && *p.oneOfType1106.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1106)
	}
	return nil, errors.New("No value to marshal for OneOfKeyManagementServerAccessInformation")
}

type OneOfCredentialCredentialDetails struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType1003 *IntersightCredential `json:"-"`
	oneOfType1002 *VcenterCredential    `json:"-"`
	oneOfType1001 *BmcCredential        `json:"-"`
}

func NewOneOfCredentialCredentialDetails() *OneOfCredentialCredentialDetails {
	p := new(OneOfCredentialCredentialDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCredentialCredentialDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCredentialCredentialDetails is nil"))
	}
	switch v.(type) {
	case IntersightCredential:
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(IntersightCredential)
		}
		*p.oneOfType1003 = v.(IntersightCredential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1003.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1003.ObjectType_
	case VcenterCredential:
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(VcenterCredential)
		}
		*p.oneOfType1002 = v.(VcenterCredential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1002.ObjectType_
	case BmcCredential:
		if nil == p.oneOfType1001 {
			p.oneOfType1001 = new(BmcCredential)
		}
		*p.oneOfType1001 = v.(BmcCredential)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCredentialCredentialDetails) GetValue() interface{} {
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1003
	}
	if p.oneOfType1002 != nil && *p.oneOfType1002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1002
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1001
	}
	return nil
}

func (p *OneOfCredentialCredentialDetails) UnmarshalJSON(b []byte) error {
	vOneOfType1003 := new(IntersightCredential)
	if err := json.Unmarshal(b, vOneOfType1003); err == nil {
		if "security.v4.config.IntersightCredential" == *vOneOfType1003.ObjectType_ {
			if nil == p.oneOfType1003 {
				p.oneOfType1003 = new(IntersightCredential)
			}
			*p.oneOfType1003 = *vOneOfType1003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1003.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1003.ObjectType_
			return nil
		}
	}
	vOneOfType1002 := new(VcenterCredential)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if "security.v4.config.VcenterCredential" == *vOneOfType1002.ObjectType_ {
			if nil == p.oneOfType1002 {
				p.oneOfType1002 = new(VcenterCredential)
			}
			*p.oneOfType1002 = *vOneOfType1002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1002.ObjectType_
			return nil
		}
	}
	vOneOfType1001 := new(BmcCredential)
	if err := json.Unmarshal(b, vOneOfType1001); err == nil {
		if "security.v4.config.BmcCredential" == *vOneOfType1001.ObjectType_ {
			if nil == p.oneOfType1001 {
				p.oneOfType1001 = new(BmcCredential)
			}
			*p.oneOfType1001 = *vOneOfType1001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCredentialCredentialDetails"))
}

func (p *OneOfCredentialCredentialDetails) MarshalJSON() ([]byte, error) {
	if p.oneOfType1003 != nil && *p.oneOfType1003.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1003)
	}
	if p.oneOfType1002 != nil && *p.oneOfType1002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if p.oneOfType1001 != nil && *p.oneOfType1001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1001)
	}
	return nil, errors.New("No value to marshal for OneOfCredentialCredentialDetails")
}

type OneOfUpdateKeyManagementServerApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfUpdateKeyManagementServerApiResponseData() *OneOfUpdateKeyManagementServerApiResponseData {
	p := new(OneOfUpdateKeyManagementServerApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateKeyManagementServerApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateKeyManagementServerApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfUpdateKeyManagementServerApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateKeyManagementServerApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "security.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateKeyManagementServerApiResponseData"))
}

func (p *OneOfUpdateKeyManagementServerApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateKeyManagementServerApiResponseData")
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
