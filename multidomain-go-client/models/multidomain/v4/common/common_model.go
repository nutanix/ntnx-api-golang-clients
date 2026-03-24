/*
 * Generated file models/multidomain/v4/common/common_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module multidomain.v4.common of Nutanix Multidomain Versioned APIs
*/
package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Base model for action API response.
*/
type ActionBaseResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action API successful message.
	*/
	Message *string `json:"message"`
}

func (p *ActionBaseResponse) MarshalJSON() ([]byte, error) {
	type ActionBaseResponseProxy ActionBaseResponse

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ActionBaseResponseProxy
		Message *string `json:"message,omitempty"`
	}{
		ActionBaseResponseProxy: (*ActionBaseResponseProxy)(p),
		Message:                 p.Message,
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

func (p *ActionBaseResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ActionBaseResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewActionBaseResponse()

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

func NewActionBaseResponse() *ActionBaseResponse {
	p := new(ActionBaseResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.common.ActionBaseResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
API response for the domain.
*/
type DomainError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  User friendly error message against error received from domain
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/*
	  Original response received from the domain.
	*/
	RawError *string `json:"rawError,omitempty"`
	/*
	  Resource article URL for domain error handling. for example KB article link
	*/
	ResourceURL *string `json:"resourceURL,omitempty"`
}

func (p *DomainError) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainError

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

func (p *DomainError) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainError
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainError()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.RawError != nil {
		p.RawError = known.RawError
	}
	if known.ResourceURL != nil {
		p.ResourceURL = known.ResourceURL
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "errorMessage")
	delete(allFields, "rawError")
	delete(allFields, "resourceURL")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainError() *DomainError {
	p := new(DomainError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.common.DomainError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This contains the basic information about the domain
*/
type DomainInformation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  This stores the uuid of the domain
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Name of the domain
	*/
	Name *string `json:"name,omitempty"`
}

func (p *DomainInformation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainInformation

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

func (p *DomainInformation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainInformation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainInformation()

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
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainInformation() *DomainInformation {
	p := new(DomainInformation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.common.DomainInformation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains details about credentials to be used by the registered domain (i.e. Prism Central) to talk to Nutanix Central.
*/
type RegistrationCredentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  API key which is used by Prism Central as authentication for Nutanix Central communication.
	*/
	ApiKey *string `json:"apiKey"`
	/*
	  This is used by Prism Central to form authentication token for Nutanix Central communication.
	*/
	KeyId *string `json:"keyId"`
}

func (p *RegistrationCredentials) MarshalJSON() ([]byte, error) {
	type RegistrationCredentialsProxy RegistrationCredentials

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RegistrationCredentialsProxy
		ApiKey *string `json:"apiKey,omitempty"`
		KeyId  *string `json:"keyId,omitempty"`
	}{
		RegistrationCredentialsProxy: (*RegistrationCredentialsProxy)(p),
		ApiKey:                       p.ApiKey,
		KeyId:                        p.KeyId,
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

func (p *RegistrationCredentials) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegistrationCredentials
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegistrationCredentials()

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
	if known.KeyId != nil {
		p.KeyId = known.KeyId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiKey")
	delete(allFields, "keyId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegistrationCredentials() *RegistrationCredentials {
	p := new(RegistrationCredentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.common.RegistrationCredentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Indicates the state of domain registration.
*/
type RegistrationState int

const (
	REGISTRATIONSTATE_UNKNOWN              RegistrationState = 0
	REGISTRATIONSTATE_REDACTED             RegistrationState = 1
	REGISTRATIONSTATE_UNREGISTERED         RegistrationState = 2
	REGISTRATIONSTATE_REGISTERING          RegistrationState = 3
	REGISTRATIONSTATE_REGISTERED           RegistrationState = 4
	REGISTRATIONSTATE_UNREGISTERING        RegistrationState = 5
	REGISTRATIONSTATE_REGISTRATION_ERROR   RegistrationState = 6
	REGISTRATIONSTATE_UNREGISTRATION_ERROR RegistrationState = 7
	REGISTRATIONSTATE_DRAFT                RegistrationState = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RegistrationState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"REGISTRATION_ERROR",
		"UNREGISTRATION_ERROR",
		"DRAFT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RegistrationState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"REGISTRATION_ERROR",
		"UNREGISTRATION_ERROR",
		"DRAFT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RegistrationState) index(name string) RegistrationState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNREGISTERED",
		"REGISTERING",
		"REGISTERED",
		"UNREGISTERING",
		"REGISTRATION_ERROR",
		"UNREGISTRATION_ERROR",
		"DRAFT",
	}
	for idx := range names {
		if names[idx] == name {
			return RegistrationState(idx)
		}
	}
	return REGISTRATIONSTATE_UNKNOWN
}

func (e *RegistrationState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RegistrationState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RegistrationState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RegistrationState) Ref() *RegistrationState {
	return &e
}

/*
Contains information about the tenant under which the domain is added on Nutanix Central.
*/
type TenantInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The tenant's unique identifier.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The tenant's name.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *TenantInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TenantInfo

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

func (p *TenantInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TenantInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenantInfo()

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
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenantInfo() *TenantInfo {
	p := new(TenantInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.common.TenantInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of tunnel providers for the domain.
*/
type TunnelProvider int

const (
	TUNNELPROVIDER_UNKNOWN  TunnelProvider = 0
	TUNNELPROVIDER_REDACTED TunnelProvider = 1
	TUNNELPROVIDER_IKAT     TunnelProvider = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TunnelProvider) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IKAT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TunnelProvider) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IKAT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TunnelProvider) index(name string) TunnelProvider {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IKAT",
	}
	for idx := range names {
		if names[idx] == name {
			return TunnelProvider(idx)
		}
	}
	return TUNNELPROVIDER_UNKNOWN
}

func (e *TunnelProvider) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TunnelProvider:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TunnelProvider) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TunnelProvider) Ref() *TunnelProvider {
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
