/*
 * Generated file models/security/v4/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module security.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
)

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
	*p = BasicAuthCredential(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBasicAuthCredential() *BasicAuthCredential {
	p := new(BasicAuthCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.BasicAuthCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p = IntersightCredential(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")
	delete(allFields, "deploymentType")
	delete(allFields, "type")
	delete(allFields, "url")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIntersightCredential() *IntersightCredential {
	p := new(IntersightCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.IntersightCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	*p = KeyBasedAuth(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiKey")
	delete(allFields, "secretKey")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewKeyBasedAuth() *KeyBasedAuth {
	p := new(KeyBasedAuth)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KeyBasedAuth"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	*p = KeyBasedAuthCredential(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credential")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewKeyBasedAuthCredential() *KeyBasedAuthCredential {
	p := new(KeyBasedAuthCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.KeyBasedAuthCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p = VcenterCredential(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "address")
	delete(allFields, "credential")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVcenterCredential() *VcenterCredential {
	p := new(VcenterCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "security.v4.config.VcenterCredential"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
