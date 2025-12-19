/*
 * Generated file models/lifecycle/v4/config/config_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module lifecycle.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/prism/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/prism/v4/management"
	"time"
)

/*
Bill of materials which describes the various products and their versions compataible with each other
*/
type Bom struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Bill of materials content i.e list of products and their versions
	*/
	Content *string `json:"content,omitempty"`
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
	  Bill of materials version
	*/
	Version *string `json:"version,omitempty"`
}

func (p *Bom) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Bom

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

func (p *Bom) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Bom
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBom()

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
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "content")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBom() *Bom {
	p := new(Bom)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Bom"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BomProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Bill of materials content i.e list of products and their versions
	*/
	Content *string `json:"content,omitempty"`
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
	  Bill of materials version
	*/
	Version *string `json:"version,omitempty"`
}

func (p *BomProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BomProjection

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

func (p *BomProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BomProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBomProjection()

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
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "content")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBomProjection() *BomProjection {
	p := new(BomProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.BomProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Domain struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BomExtId *string `json:"bomExtId"`
	/*
	  List of clusters managed by App Domain PC/ Management PC
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`

	ClusterProfileExtId *string `json:"clusterProfileExtId"`
	/*
	  The user who created the Domain
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Created Time of Domain
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DomainManagerItemDiscriminator_ *string `json:"$domainManagerItemDiscriminator,omitempty"`
	/*
	  Domain Manager can be one of the following: - DomainManager object which is required to take input for creating a new domain manager (Prism Central) for App Domain PC creation. - RemoteClusterSpec  is mainly required for ingestion usecase of existing PC - RemoteCluster is used for output only to show basic details of existing app domain PC
	*/
	DomainManager *OneOfDomainDomainManager `json:"domainManager,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  List of posthook operations to be run during app domain creation workflow
	*/
	PostHookOps []HookDetails `json:"postHookOps,omitempty"`
	/*
	  List of prehook operations to be run during app domain creation workflow
	*/
	PreHookOps []HookDetails `json:"preHookOps,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *DomainType `json:"type"`
}

func (p *Domain) MarshalJSON() ([]byte, error) {
	type DomainProxy Domain

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainProxy
		BomExtId            *string     `json:"bomExtId,omitempty"`
		ClusterProfileExtId *string     `json:"clusterProfileExtId,omitempty"`
		Type                *DomainType `json:"type,omitempty"`
	}{
		DomainProxy:         (*DomainProxy)(p),
		BomExtId:            p.BomExtId,
		ClusterProfileExtId: p.ClusterProfileExtId,
		Type:                p.Type,
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

func (p *Domain) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Domain
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomain()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BomExtId != nil {
		p.BomExtId = known.BomExtId
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.ClusterProfileExtId != nil {
		p.ClusterProfileExtId = known.ClusterProfileExtId
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DomainManagerItemDiscriminator_ != nil {
		p.DomainManagerItemDiscriminator_ = known.DomainManagerItemDiscriminator_
	}
	if known.DomainManager != nil {
		p.DomainManager = known.DomainManager
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.PostHookOps != nil {
		p.PostHookOps = known.PostHookOps
	}
	if known.PreHookOps != nil {
		p.PreHookOps = known.PreHookOps
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
	delete(allFields, "bomExtId")
	delete(allFields, "clusterExtIds")
	delete(allFields, "clusterProfileExtId")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "$domainManagerItemDiscriminator")
	delete(allFields, "domainManager")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "postHookOps")
	delete(allFields, "preHookOps")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomain() *Domain {
	p := new(Domain)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.Domain"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Domain) GetDomainManager() interface{} {
	if nil == p.DomainManager {
		return nil
	}
	return p.DomainManager.GetValue()
}

func (p *Domain) SetDomainManager(v interface{}) error {
	if nil == p.DomainManager {
		p.DomainManager = NewOneOfDomainDomainManager()
	}
	e := p.DomainManager.SetValue(v)
	if nil == e {
		if nil == p.DomainManagerItemDiscriminator_ {
			p.DomainManagerItemDiscriminator_ = new(string)
		}
		*p.DomainManagerItemDiscriminator_ = *p.DomainManager.Discriminator
	}
	return e
}

type DomainProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BomExtId *string `json:"bomExtId"`

	BomProjection *BomProjection `json:"bomProjection,omitempty"`
	/*
	  List of clusters managed by App Domain PC/ Management PC
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`

	ClusterProfileExtId *string `json:"clusterProfileExtId"`
	/*
	  The user who created the Domain
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Created Time of Domain
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`

	DomainManagerItemDiscriminator_ *string `json:"$domainManagerItemDiscriminator,omitempty"`
	/*
	  Domain Manager can be one of the following: - DomainManager object which is required to take input for creating a new domain manager (Prism Central) for App Domain PC creation. - RemoteClusterSpec  is mainly required for ingestion usecase of existing PC - RemoteCluster is used for output only to show basic details of existing app domain PC
	*/
	DomainManager *OneOfDomainProjectionDomainManager `json:"domainManager,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  List of posthook operations to be run during app domain creation workflow
	*/
	PostHookOps []HookDetails `json:"postHookOps,omitempty"`
	/*
	  List of prehook operations to be run during app domain creation workflow
	*/
	PreHookOps []HookDetails `json:"preHookOps,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *DomainType `json:"type"`
}

func (p *DomainProjection) MarshalJSON() ([]byte, error) {
	type DomainProjectionProxy DomainProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainProjectionProxy
		BomExtId            *string     `json:"bomExtId,omitempty"`
		ClusterProfileExtId *string     `json:"clusterProfileExtId,omitempty"`
		Type                *DomainType `json:"type,omitempty"`
	}{
		DomainProjectionProxy: (*DomainProjectionProxy)(p),
		BomExtId:              p.BomExtId,
		ClusterProfileExtId:   p.ClusterProfileExtId,
		Type:                  p.Type,
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

func (p *DomainProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BomExtId != nil {
		p.BomExtId = known.BomExtId
	}
	if known.BomProjection != nil {
		p.BomProjection = known.BomProjection
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.ClusterProfileExtId != nil {
		p.ClusterProfileExtId = known.ClusterProfileExtId
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.DomainManagerItemDiscriminator_ != nil {
		p.DomainManagerItemDiscriminator_ = known.DomainManagerItemDiscriminator_
	}
	if known.DomainManager != nil {
		p.DomainManager = known.DomainManager
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.PostHookOps != nil {
		p.PostHookOps = known.PostHookOps
	}
	if known.PreHookOps != nil {
		p.PreHookOps = known.PreHookOps
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
	delete(allFields, "bomExtId")
	delete(allFields, "bomProjection")
	delete(allFields, "clusterExtIds")
	delete(allFields, "clusterProfileExtId")
	delete(allFields, "createdBy")
	delete(allFields, "createdTime")
	delete(allFields, "$domainManagerItemDiscriminator")
	delete(allFields, "domainManager")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "postHookOps")
	delete(allFields, "preHookOps")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainProjection() *DomainProjection {
	p := new(DomainProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.DomainProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain Type is an enum having values App Domain and Management Domain
*/
type DomainType int

const (
	DOMAINTYPE_UNKNOWN          DomainType = 0
	DOMAINTYPE_REDACTED         DomainType = 1
	DOMAINTYPE_MANAGEMENTDOMAIN DomainType = 2
	DOMAINTYPE_APPDOMAIN        DomainType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENTDOMAIN",
		"APPDOMAIN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENTDOMAIN",
		"APPDOMAIN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainType) index(name string) DomainType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENTDOMAIN",
		"APPDOMAIN",
	}
	for idx := range names {
		if names[idx] == name {
			return DomainType(idx)
		}
	}
	return DOMAINTYPE_UNKNOWN
}

func (e *DomainType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DomainType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DomainType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DomainType) Ref() *DomainType {
	return &e
}

/*
Hook is a wrapper model over prehook/post hook model to avoid duplication
*/
type HookDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	OperationType *OperationType `json:"operationType,omitempty"`
	/*
	  Prehook/Posthook script name
	*/
	Script *string `json:"script,omitempty"`

	ScriptType *ScriptType `json:"scriptType,omitempty"`
}

func (p *HookDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HookDetails

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

func (p *HookDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HookDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHookDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.OperationType != nil {
		p.OperationType = known.OperationType
	}
	if known.Script != nil {
		p.Script = known.Script
	}
	if known.ScriptType != nil {
		p.ScriptType = known.ScriptType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "operationType")
	delete(allFields, "script")
	delete(allFields, "scriptType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHookDetails() *HookDetails {
	p := new(HookDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.config.HookDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation type is an enum which denotes which of the following operations to be run: - Create Domain Manager
*/
type OperationType int

const (
	OPERATIONTYPE_UNKNOWN             OperationType = 0
	OPERATIONTYPE_REDACTED            OperationType = 1
	OPERATIONTYPE_CREATEDOMAINMANAGER OperationType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CREATEDOMAINMANAGER",
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
		"CREATEDOMAINMANAGER",
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
		"CREATEDOMAINMANAGER",
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
Script Type currently supports only Escript type which is used to run scripts on the cluster
*/
type ScriptType int

const (
	SCRIPTTYPE_UNKNOWN  ScriptType = 0
	SCRIPTTYPE_REDACTED ScriptType = 1
	SCRIPTTYPE_ESCRIPT  ScriptType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ScriptType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESCRIPT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ScriptType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESCRIPT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ScriptType) index(name string) ScriptType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ESCRIPT",
	}
	for idx := range names {
		if names[idx] == name {
			return ScriptType(idx)
		}
	}
	return SCRIPTTYPE_UNKNOWN
}

func (e *ScriptType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ScriptType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ScriptType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ScriptType) Ref() *ScriptType {
	return &e
}

type OneOfDomainDomainManager struct {
	Discriminator  *string                    `json:"-"`
	ObjectType_    *string                    `json:"-"`
	oneOfType12002 *import3.RemoteClusterSpec `json:"-"`
	oneOfType12001 *import2.DomainManager     `json:"-"`
}

func NewOneOfDomainDomainManager() *OneOfDomainDomainManager {
	p := new(OneOfDomainDomainManager)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDomainDomainManager) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDomainDomainManager is nil"))
	}
	switch v.(type) {
	case import3.RemoteClusterSpec:
		if nil == p.oneOfType12002 {
			p.oneOfType12002 = new(import3.RemoteClusterSpec)
		}
		*p.oneOfType12002 = v.(import3.RemoteClusterSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12002.ObjectType_
	case import2.DomainManager:
		if nil == p.oneOfType12001 {
			p.oneOfType12001 = new(import2.DomainManager)
		}
		*p.oneOfType12001 = v.(import2.DomainManager)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDomainDomainManager) GetValue() interface{} {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12002
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12001
	}
	return nil
}

func (p *OneOfDomainDomainManager) UnmarshalJSON(b []byte) error {
	vOneOfType12002 := new(import3.RemoteClusterSpec)
	if err := json.Unmarshal(b, vOneOfType12002); err == nil {
		if "prism.v4.management.RemoteClusterSpec" == *vOneOfType12002.ObjectType_ {
			if nil == p.oneOfType12002 {
				p.oneOfType12002 = new(import3.RemoteClusterSpec)
			}
			*p.oneOfType12002 = *vOneOfType12002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12002.ObjectType_
			return nil
		}
	}
	vOneOfType12001 := new(import2.DomainManager)
	if err := json.Unmarshal(b, vOneOfType12001); err == nil {
		if "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
			if nil == p.oneOfType12001 {
				p.oneOfType12001 = new(import2.DomainManager)
			}
			*p.oneOfType12001 = *vOneOfType12001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainDomainManager"))
}

func (p *OneOfDomainDomainManager) MarshalJSON() ([]byte, error) {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12002)
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12001)
	}
	return nil, errors.New("No value to marshal for OneOfDomainDomainManager")
}

type OneOfDomainProjectionDomainManager struct {
	Discriminator  *string                    `json:"-"`
	ObjectType_    *string                    `json:"-"`
	oneOfType12002 *import3.RemoteClusterSpec `json:"-"`
	oneOfType12001 *import2.DomainManager     `json:"-"`
}

func NewOneOfDomainProjectionDomainManager() *OneOfDomainProjectionDomainManager {
	p := new(OneOfDomainProjectionDomainManager)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDomainProjectionDomainManager) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDomainProjectionDomainManager is nil"))
	}
	switch v.(type) {
	case import3.RemoteClusterSpec:
		if nil == p.oneOfType12002 {
			p.oneOfType12002 = new(import3.RemoteClusterSpec)
		}
		*p.oneOfType12002 = v.(import3.RemoteClusterSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12002.ObjectType_
	case import2.DomainManager:
		if nil == p.oneOfType12001 {
			p.oneOfType12001 = new(import2.DomainManager)
		}
		*p.oneOfType12001 = v.(import2.DomainManager)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDomainProjectionDomainManager) GetValue() interface{} {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12002
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12001
	}
	return nil
}

func (p *OneOfDomainProjectionDomainManager) UnmarshalJSON(b []byte) error {
	vOneOfType12002 := new(import3.RemoteClusterSpec)
	if err := json.Unmarshal(b, vOneOfType12002); err == nil {
		if "prism.v4.management.RemoteClusterSpec" == *vOneOfType12002.ObjectType_ {
			if nil == p.oneOfType12002 {
				p.oneOfType12002 = new(import3.RemoteClusterSpec)
			}
			*p.oneOfType12002 = *vOneOfType12002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12002.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12002.ObjectType_
			return nil
		}
	}
	vOneOfType12001 := new(import2.DomainManager)
	if err := json.Unmarshal(b, vOneOfType12001); err == nil {
		if "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
			if nil == p.oneOfType12001 {
				p.oneOfType12001 = new(import2.DomainManager)
			}
			*p.oneOfType12001 = *vOneOfType12001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType12001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType12001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainProjectionDomainManager"))
}

func (p *OneOfDomainProjectionDomainManager) MarshalJSON() ([]byte, error) {
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12002)
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12001)
	}
	return nil, errors.New("No value to marshal for OneOfDomainProjectionDomainManager")
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
