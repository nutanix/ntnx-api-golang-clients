/*
 * Generated file models/multidomain/v4/config/config_model.go.
 *
 * Product version: 4.4.1-beta-1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Multi domain configuration.
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/clustermgmt/v4/config"
	import8 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
	import7 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/common"
	import4 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/error"
	import6 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/resources"
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/prism/v4/config"
	import5 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/prism/v4/management"
	"time"
)

/*
API credentials status.
*/
type ApiCredentialStatus int

const (
	APICREDENTIALSTATUS_UNKNOWN     ApiCredentialStatus = 0
	APICREDENTIALSTATUS_REDACTED    ApiCredentialStatus = 1
	APICREDENTIALSTATUS_VALID       ApiCredentialStatus = 2
	APICREDENTIALSTATUS_NEAR_EXPIRY ApiCredentialStatus = 3
	APICREDENTIALSTATUS_EXPIRED     ApiCredentialStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApiCredentialStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"NEAR_EXPIRY",
		"EXPIRED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ApiCredentialStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"NEAR_EXPIRY",
		"EXPIRED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ApiCredentialStatus) index(name string) ApiCredentialStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VALID",
		"NEAR_EXPIRY",
		"EXPIRED",
	}
	for idx := range names {
		if names[idx] == name {
			return ApiCredentialStatus(idx)
		}
	}
	return APICREDENTIALSTATUS_UNKNOWN
}

func (e *ApiCredentialStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApiCredentialStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApiCredentialStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApiCredentialStatus) Ref() *ApiCredentialStatus {
	return &e
}

/*
Application entity
*/
type Application struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Display name for the application
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Fully Qualified Domain Name for the application
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name for the application
	*/
	Name *string `json:"name,omitempty"`

	Reachability *ApplicationReachability `json:"reachability,omitempty"`
	/*
	  Source Relative URL for the application
	*/
	SourceRelativeURL *string `json:"sourceRelativeURL,omitempty"`

	State *ApplicationState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Application) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Application

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

func (p *Application) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Application
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApplication()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Reachability != nil {
		p.Reachability = known.Reachability
	}
	if known.SourceRelativeURL != nil {
		p.SourceRelativeURL = known.SourceRelativeURL
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
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "fqdn")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "reachability")
	delete(allFields, "sourceRelativeURL")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApplication() *Application {
	p := new(Application)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.Application"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ApplicationProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Display name for the application
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Fully Qualified Domain Name for the application
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name for the application
	*/
	Name *string `json:"name,omitempty"`

	Reachability *ApplicationReachability `json:"reachability,omitempty"`
	/*
	  Source Relative URL for the application
	*/
	SourceRelativeURL *string `json:"sourceRelativeURL,omitempty"`

	State *ApplicationState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ApplicationProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApplicationProjection

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

func (p *ApplicationProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApplicationProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApplicationProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Reachability != nil {
		p.Reachability = known.Reachability
	}
	if known.SourceRelativeURL != nil {
		p.SourceRelativeURL = known.SourceRelativeURL
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
	delete(allFields, "displayName")
	delete(allFields, "extId")
	delete(allFields, "fqdn")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "reachability")
	delete(allFields, "sourceRelativeURL")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApplicationProjection() *ApplicationProjection {
	p := new(ApplicationProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ApplicationProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reachability of the application
*/
type ApplicationReachability int

const (
	APPLICATIONREACHABILITY_UNKNOWN     ApplicationReachability = 0
	APPLICATIONREACHABILITY_REDACTED    ApplicationReachability = 1
	APPLICATIONREACHABILITY_REACHABLE   ApplicationReachability = 2
	APPLICATIONREACHABILITY_UNREACHABLE ApplicationReachability = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApplicationReachability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REACHABLE",
		"UNREACHABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ApplicationReachability) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REACHABLE",
		"UNREACHABLE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ApplicationReachability) index(name string) ApplicationReachability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"REACHABLE",
		"UNREACHABLE",
	}
	for idx := range names {
		if names[idx] == name {
			return ApplicationReachability(idx)
		}
	}
	return APPLICATIONREACHABILITY_UNKNOWN
}

func (e *ApplicationReachability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApplicationReachability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApplicationReachability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApplicationReachability) Ref() *ApplicationReachability {
	return &e
}

/*
State of the application
*/
type ApplicationState int

const (
	APPLICATIONSTATE_UNKNOWN  ApplicationState = 0
	APPLICATIONSTATE_REDACTED ApplicationState = 1
	APPLICATIONSTATE_ENABLED  ApplicationState = 2
	APPLICATIONSTATE_DISABLED ApplicationState = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApplicationState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ApplicationState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ApplicationState) index(name string) ApplicationState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return ApplicationState(idx)
		}
	}
	return APPLICATIONSTATE_UNKNOWN
}

func (e *ApplicationState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApplicationState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApplicationState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApplicationState) Ref() *ApplicationState {
	return &e
}

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
	*p.ObjectType_ = "multidomain.v4.config.Bom"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
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
	*p.ObjectType_ = "multidomain.v4.config.BomProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This represents the name for cluster entities group attribute.
*/
type ClusterAttributeName int

const (
	CLUSTERATTRIBUTENAME_UNKNOWN                  ClusterAttributeName = 0
	CLUSTERATTRIBUTENAME_REDACTED                 ClusterAttributeName = 1
	CLUSTERATTRIBUTENAME_NAME                     ClusterAttributeName = 2
	CLUSTERATTRIBUTENAME_UUID                     ClusterAttributeName = 3
	CLUSTERATTRIBUTENAME_DNS_NAME                 ClusterAttributeName = 4
	CLUSTERATTRIBUTENAME_DOMAIN_VERSION           ClusterAttributeName = 5
	CLUSTERATTRIBUTENAME_CRITICAL_ALERT_COUNT     ClusterAttributeName = 6
	CLUSTERATTRIBUTENAME_VM_COUNT                 ClusterAttributeName = 7
	CLUSTERATTRIBUTENAME_HOST_COUNT               ClusterAttributeName = 8
	CLUSTERATTRIBUTENAME_VLAN_COUNT               ClusterAttributeName = 9
	CLUSTERATTRIBUTENAME_STORAGE_CAPACITY         ClusterAttributeName = 10
	CLUSTERATTRIBUTENAME_STORAGE_USAGE            ClusterAttributeName = 11
	CLUSTERATTRIBUTENAME_MEMORY_CAPACITY          ClusterAttributeName = 12
	CLUSTERATTRIBUTENAME_MEMORY_USAGE             ClusterAttributeName = 13
	CLUSTERATTRIBUTENAME_CPU_COUNT                ClusterAttributeName = 14
	CLUSTERATTRIBUTENAME_CPU_USAGE                ClusterAttributeName = 15
	CLUSTERATTRIBUTENAME_DOMAIN_NAME              ClusterAttributeName = 16
	CLUSTERATTRIBUTENAME_IP_ADDRESS               ClusterAttributeName = 17
	CLUSTERATTRIBUTENAME_VCPU_COUNT               ClusterAttributeName = 18
	CLUSTERATTRIBUTENAME_PULSE_STATUS             ClusterAttributeName = 19
	CLUSTERATTRIBUTENAME_IOPS                     ClusterAttributeName = 20
	CLUSTERATTRIBUTENAME_IO_BANDWIDTH             ClusterAttributeName = 21
	CLUSTERATTRIBUTENAME_IO_LATENCY               ClusterAttributeName = 22
	CLUSTERATTRIBUTENAME_AOS_VERSION              ClusterAttributeName = 23
	CLUSTERATTRIBUTENAME_HYPERVISOR               ClusterAttributeName = 24
	CLUSTERATTRIBUTENAME_OPEN_SECURITY_ISSUES     ClusterAttributeName = 25
	CLUSTERATTRIBUTENAME_TOTAL_VULNERABILITIES    ClusterAttributeName = 26
	CLUSTERATTRIBUTENAME_CRITICAL_VULNERABILITIES ClusterAttributeName = 27
	CLUSTERATTRIBUTENAME_STIG_POLICY_VIOLATIONS   ClusterAttributeName = 28
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterAttributeName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"DOMAIN_VERSION",
		"CRITICAL_ALERT_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"DOMAIN_NAME",
		"IP_ADDRESS",
		"VCPU_COUNT",
		"PULSE_STATUS",
		"IOPS",
		"IO_BANDWIDTH",
		"IO_LATENCY",
		"AOS_VERSION",
		"HYPERVISOR",
		"OPEN_SECURITY_ISSUES",
		"TOTAL_VULNERABILITIES",
		"CRITICAL_VULNERABILITIES",
		"STIG_POLICY_VIOLATIONS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterAttributeName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"DOMAIN_VERSION",
		"CRITICAL_ALERT_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"DOMAIN_NAME",
		"IP_ADDRESS",
		"VCPU_COUNT",
		"PULSE_STATUS",
		"IOPS",
		"IO_BANDWIDTH",
		"IO_LATENCY",
		"AOS_VERSION",
		"HYPERVISOR",
		"OPEN_SECURITY_ISSUES",
		"TOTAL_VULNERABILITIES",
		"CRITICAL_VULNERABILITIES",
		"STIG_POLICY_VIOLATIONS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterAttributeName) index(name string) ClusterAttributeName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"DOMAIN_VERSION",
		"CRITICAL_ALERT_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"DOMAIN_NAME",
		"IP_ADDRESS",
		"VCPU_COUNT",
		"PULSE_STATUS",
		"IOPS",
		"IO_BANDWIDTH",
		"IO_LATENCY",
		"AOS_VERSION",
		"HYPERVISOR",
		"OPEN_SECURITY_ISSUES",
		"TOTAL_VULNERABILITIES",
		"CRITICAL_VULNERABILITIES",
		"STIG_POLICY_VIOLATIONS",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterAttributeName(idx)
		}
	}
	return CLUSTERATTRIBUTENAME_UNKNOWN
}

func (e *ClusterAttributeName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterAttributeName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterAttributeName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterAttributeName) Ref() *ClusterAttributeName {
	return &e
}

/*
Cluster object details of cluster to be registered to App Domain PC
*/
type CreateClusterSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of clusters to be created.
	*/
	Clusters []import2.Cluster `json:"clusters"`
	/*
	  List of posthook operations to be run during app domain creation workflow
	*/
	PostHookOps []HookDetails `json:"postHookOps,omitempty"`
	/*
	  List of prehook operations to be run during app domain creation workflow
	*/
	PreHookOps []HookDetails `json:"preHookOps,omitempty"`
}

func (p *CreateClusterSpec) MarshalJSON() ([]byte, error) {
	type CreateClusterSpecProxy CreateClusterSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CreateClusterSpecProxy
		Clusters []import2.Cluster `json:"clusters,omitempty"`
	}{
		CreateClusterSpecProxy: (*CreateClusterSpecProxy)(p),
		Clusters:               p.Clusters,
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

func (p *CreateClusterSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateClusterSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateClusterSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Clusters != nil {
		p.Clusters = known.Clusters
	}
	if known.PostHookOps != nil {
		p.PostHookOps = known.PostHookOps
	}
	if known.PreHookOps != nil {
		p.PreHookOps = known.PreHookOps
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusters")
	delete(allFields, "postHookOps")
	delete(allFields, "preHookOps")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateClusterSpec() *CreateClusterSpec {
	p := new(CreateClusterSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.CreateClusterSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/config/external-repositories Post operation
*/
type CreateExternalRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateExternalRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateExternalRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateExternalRepositoryApiResponse

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

func (p *CreateExternalRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateExternalRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateExternalRepositoryApiResponse()

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

func NewCreateExternalRepositoryApiResponse() *CreateExternalRepositoryApiResponse {
	p := new(CreateExternalRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.CreateExternalRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateExternalRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateExternalRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateExternalRepositoryApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/projects Post operation
*/
type CreateProjectApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateProjectApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateProjectApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateProjectApiResponse

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

func (p *CreateProjectApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateProjectApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateProjectApiResponse()

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

func NewCreateProjectApiResponse() *CreateProjectApiResponse {
	p := new(CreateProjectApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.CreateProjectApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateProjectApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateProjectApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateProjectApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains Post operation
*/
type CreateRegisteredDomainApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateRegisteredDomainApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateRegisteredDomainApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateRegisteredDomainApiResponse

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

func (p *CreateRegisteredDomainApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateRegisteredDomainApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateRegisteredDomainApiResponse()

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

func NewCreateRegisteredDomainApiResponse() *CreateRegisteredDomainApiResponse {
	p := new(CreateRegisteredDomainApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.CreateRegisteredDomainApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateRegisteredDomainApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateRegisteredDomainApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateRegisteredDomainApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/resource-groups Post operation
*/
type CreateResourceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateResourceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateResourceGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateResourceGroupApiResponse

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

func (p *CreateResourceGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateResourceGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateResourceGroupApiResponse()

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

func NewCreateResourceGroupApiResponse() *CreateResourceGroupApiResponse {
	p := new(CreateResourceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.CreateResourceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateResourceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateResourceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateResourceGroupApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/external-repositories/{extId} Delete operation
*/
type DeleteExternalRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteExternalRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteExternalRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteExternalRepositoryApiResponse

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

func (p *DeleteExternalRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteExternalRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteExternalRepositoryApiResponse()

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

func NewDeleteExternalRepositoryApiResponse() *DeleteExternalRepositoryApiResponse {
	p := new(DeleteExternalRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DeleteExternalRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteExternalRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteExternalRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteExternalRepositoryApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/projects/{extId} Delete operation
*/
type DeleteProjectApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteProjectApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteProjectApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteProjectApiResponse

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

func (p *DeleteProjectApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteProjectApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteProjectApiResponse()

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

func NewDeleteProjectApiResponse() *DeleteProjectApiResponse {
	p := new(DeleteProjectApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DeleteProjectApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteProjectApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteProjectApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteProjectApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains/{extId} Delete operation
*/
type DeleteRegisteredDomainApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRegisteredDomainApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteRegisteredDomainApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteRegisteredDomainApiResponse

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

func (p *DeleteRegisteredDomainApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteRegisteredDomainApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteRegisteredDomainApiResponse()

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

func NewDeleteRegisteredDomainApiResponse() *DeleteRegisteredDomainApiResponse {
	p := new(DeleteRegisteredDomainApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DeleteRegisteredDomainApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRegisteredDomainApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRegisteredDomainApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRegisteredDomainApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/resource-groups/{extId} Delete operation
*/
type DeleteResourceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteResourceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteResourceGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteResourceGroupApiResponse

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

func (p *DeleteResourceGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteResourceGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteResourceGroupApiResponse()

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

func NewDeleteResourceGroupApiResponse() *DeleteResourceGroupApiResponse {
	p := new(DeleteResourceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DeleteResourceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteResourceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteResourceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteResourceGroupApiResponseData()
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

type Domain struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Bill of materials ExtID
	*/
	BomExtId *string `json:"bomExtId"`
	/*
	  Total number of clusters registered to this domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  List of clusters managed by App Domain PC/ Management PC
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  Cluster Profile External ID is a unique identifier for a cluster profile used to apply common settings to a Domain.
	*/
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
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
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
	delete(allFields, "clusterCount")
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
	*p.ObjectType_ = "multidomain.v4.config.Domain"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
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

/*
This represents the name for the domain entities group attribute.
*/
type DomainAttributeName int

const (
	DOMAINATTRIBUTENAME_UNKNOWN              DomainAttributeName = 0
	DOMAINATTRIBUTENAME_REDACTED             DomainAttributeName = 1
	DOMAINATTRIBUTENAME_NAME                 DomainAttributeName = 2
	DOMAINATTRIBUTENAME_UUID                 DomainAttributeName = 3
	DOMAINATTRIBUTENAME_DNS_NAME             DomainAttributeName = 4
	DOMAINATTRIBUTENAME_LOCATION             DomainAttributeName = 5
	DOMAINATTRIBUTENAME_VERSION              DomainAttributeName = 6
	DOMAINATTRIBUTENAME_CRITICAL_ALERT_COUNT DomainAttributeName = 7
	DOMAINATTRIBUTENAME_CLUSTER_COUNT        DomainAttributeName = 8
	DOMAINATTRIBUTENAME_VM_COUNT             DomainAttributeName = 9
	DOMAINATTRIBUTENAME_HOST_COUNT           DomainAttributeName = 10
	DOMAINATTRIBUTENAME_VLAN_COUNT           DomainAttributeName = 11
	DOMAINATTRIBUTENAME_SCALE_FACTOR         DomainAttributeName = 12
	DOMAINATTRIBUTENAME_IP_ADDRESS           DomainAttributeName = 13
	DOMAINATTRIBUTENAME_SIZE                 DomainAttributeName = 14
	DOMAINATTRIBUTENAME_VPC_COUNT            DomainAttributeName = 15
	DOMAINATTRIBUTENAME_STORAGE_CAPACITY     DomainAttributeName = 16
	DOMAINATTRIBUTENAME_STORAGE_USAGE        DomainAttributeName = 17
	DOMAINATTRIBUTENAME_MEMORY_CAPACITY      DomainAttributeName = 18
	DOMAINATTRIBUTENAME_MEMORY_USAGE         DomainAttributeName = 19
	DOMAINATTRIBUTENAME_CPU_COUNT            DomainAttributeName = 20
	DOMAINATTRIBUTENAME_CPU_USAGE            DomainAttributeName = 21
	DOMAINATTRIBUTENAME_CONNECTIVITY_STATUS  DomainAttributeName = 22
	DOMAINATTRIBUTENAME_REGISTERED_AT        DomainAttributeName = 23
	DOMAINATTRIBUTENAME_LAST_SYNCED_AT       DomainAttributeName = 24
	DOMAINATTRIBUTENAME_CREATED_BY           DomainAttributeName = 25
	DOMAINATTRIBUTENAME_ENV_TYPE             DomainAttributeName = 26
	DOMAINATTRIBUTENAME_PROVIDER_TYPE        DomainAttributeName = 27
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainAttributeName) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"LOCATION",
		"VERSION",
		"CRITICAL_ALERT_COUNT",
		"CLUSTER_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"SCALE_FACTOR",
		"IP_ADDRESS",
		"SIZE",
		"VPC_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"CONNECTIVITY_STATUS",
		"REGISTERED_AT",
		"LAST_SYNCED_AT",
		"CREATED_BY",
		"ENV_TYPE",
		"PROVIDER_TYPE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainAttributeName) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"LOCATION",
		"VERSION",
		"CRITICAL_ALERT_COUNT",
		"CLUSTER_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"SCALE_FACTOR",
		"IP_ADDRESS",
		"SIZE",
		"VPC_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"CONNECTIVITY_STATUS",
		"REGISTERED_AT",
		"LAST_SYNCED_AT",
		"CREATED_BY",
		"ENV_TYPE",
		"PROVIDER_TYPE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainAttributeName) index(name string) DomainAttributeName {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NAME",
		"UUID",
		"DNS_NAME",
		"LOCATION",
		"VERSION",
		"CRITICAL_ALERT_COUNT",
		"CLUSTER_COUNT",
		"VM_COUNT",
		"HOST_COUNT",
		"VLAN_COUNT",
		"SCALE_FACTOR",
		"IP_ADDRESS",
		"SIZE",
		"VPC_COUNT",
		"STORAGE_CAPACITY",
		"STORAGE_USAGE",
		"MEMORY_CAPACITY",
		"MEMORY_USAGE",
		"CPU_COUNT",
		"CPU_USAGE",
		"CONNECTIVITY_STATUS",
		"REGISTERED_AT",
		"LAST_SYNCED_AT",
		"CREATED_BY",
		"ENV_TYPE",
		"PROVIDER_TYPE",
	}
	for idx := range names {
		if names[idx] == name {
			return DomainAttributeName(idx)
		}
	}
	return DOMAINATTRIBUTENAME_UNKNOWN
}

func (e *DomainAttributeName) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DomainAttributeName:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DomainAttributeName) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DomainAttributeName) Ref() *DomainAttributeName {
	return &e
}

/*
Domain platform data
*/
type DomainPlatformData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of clusters in the domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  Domain connectivity status
	*/
	ConnectivityStatus *string `json:"connectivityStatus,omitempty"`
	/*
	  CPU capacity of the domain
	*/
	CpuCount *int `json:"cpuCount,omitempty"`
	/*
	  CPU usage of the domain
	*/
	CpuUsage *float64 `json:"cpuUsage,omitempty"`
	/*
	  User/Service which created this domain
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Number of alerts in the domain
	*/
	CriticalAlertCount *int `json:"criticalAlertCount,omitempty"`
	/*
	  DNS name of domain
	*/
	DnsName *string `json:"dnsName,omitempty"`
	/*
	  Error messages from domain
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`

	ExtId *string `json:"extId,omitempty"`
	/*
	  Number of hosts in the domain
	*/
	HostCount *int `json:"hostCount,omitempty"`
	/*
	  Domain last synced date
	*/
	LastSyncedAt *time.Time `json:"lastSyncedAt,omitempty"`

	Location *Location `json:"location,omitempty"`
	/*
	  Memory capacity of the domain
	*/
	MemoryCapacity *int64 `json:"memoryCapacity,omitempty"`
	/*
	  Memory usage of the domain
	*/
	MemoryUsage *float64 `json:"memoryUsage,omitempty"`
	/*
	  Name of the registered domain.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  PC Environment Type of the domain
	*/
	PcEnvType *string `json:"pcEnvType,omitempty"`
	/*
	  The IP address of the Prism Central.
	*/
	PcIpAddress *string `json:"pcIpAddress,omitempty"`
	/*
	  PC Provider Type of the domain
	*/
	PcProviderType *string `json:"pcProviderType,omitempty"`

	PcSize *PCSize `json:"pcSize,omitempty"`
	/*
	  PC Version of the domain
	*/
	PcVersion *string `json:"pcVersion,omitempty"`
	/*
	  VM count of Prism Central
	*/
	PcVmCount *int `json:"pcVmCount,omitempty"`
	/*
	  Domain registration date
	*/
	RegisteredAt *time.Time `json:"registeredAt,omitempty"`
	/*
	  Storage capacity of the domain
	*/
	StorageCapacity *int64 `json:"storageCapacity,omitempty"`
	/*
	  Storage usage of the domain
	*/
	StorageUsage *float64 `json:"storageUsage,omitempty"`
	/*
	  Number of VLANs in the domain
	*/
	VlanCount *int `json:"vlanCount,omitempty"`
	/*
	  Number of VMs in the domain
	*/
	VmCount *int `json:"vmCount,omitempty"`
	/*
	  Number of VPCs in the domain
	*/
	VpcCount *int `json:"vpcCount,omitempty"`
}

func (p *DomainPlatformData) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainPlatformData

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

func (p *DomainPlatformData) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainPlatformData
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainPlatformData()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.ConnectivityStatus != nil {
		p.ConnectivityStatus = known.ConnectivityStatus
	}
	if known.CpuCount != nil {
		p.CpuCount = known.CpuCount
	}
	if known.CpuUsage != nil {
		p.CpuUsage = known.CpuUsage
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CriticalAlertCount != nil {
		p.CriticalAlertCount = known.CriticalAlertCount
	}
	if known.DnsName != nil {
		p.DnsName = known.DnsName
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostCount != nil {
		p.HostCount = known.HostCount
	}
	if known.LastSyncedAt != nil {
		p.LastSyncedAt = known.LastSyncedAt
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.MemoryCapacity != nil {
		p.MemoryCapacity = known.MemoryCapacity
	}
	if known.MemoryUsage != nil {
		p.MemoryUsage = known.MemoryUsage
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PcEnvType != nil {
		p.PcEnvType = known.PcEnvType
	}
	if known.PcIpAddress != nil {
		p.PcIpAddress = known.PcIpAddress
	}
	if known.PcProviderType != nil {
		p.PcProviderType = known.PcProviderType
	}
	if known.PcSize != nil {
		p.PcSize = known.PcSize
	}
	if known.PcVersion != nil {
		p.PcVersion = known.PcVersion
	}
	if known.PcVmCount != nil {
		p.PcVmCount = known.PcVmCount
	}
	if known.RegisteredAt != nil {
		p.RegisteredAt = known.RegisteredAt
	}
	if known.StorageCapacity != nil {
		p.StorageCapacity = known.StorageCapacity
	}
	if known.StorageUsage != nil {
		p.StorageUsage = known.StorageUsage
	}
	if known.VlanCount != nil {
		p.VlanCount = known.VlanCount
	}
	if known.VmCount != nil {
		p.VmCount = known.VmCount
	}
	if known.VpcCount != nil {
		p.VpcCount = known.VpcCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterCount")
	delete(allFields, "connectivityStatus")
	delete(allFields, "cpuCount")
	delete(allFields, "cpuUsage")
	delete(allFields, "createdBy")
	delete(allFields, "criticalAlertCount")
	delete(allFields, "dnsName")
	delete(allFields, "errorMessage")
	delete(allFields, "extId")
	delete(allFields, "hostCount")
	delete(allFields, "lastSyncedAt")
	delete(allFields, "location")
	delete(allFields, "memoryCapacity")
	delete(allFields, "memoryUsage")
	delete(allFields, "name")
	delete(allFields, "pcEnvType")
	delete(allFields, "pcIpAddress")
	delete(allFields, "pcProviderType")
	delete(allFields, "pcSize")
	delete(allFields, "pcVersion")
	delete(allFields, "pcVmCount")
	delete(allFields, "registeredAt")
	delete(allFields, "storageCapacity")
	delete(allFields, "storageUsage")
	delete(allFields, "vlanCount")
	delete(allFields, "vmCount")
	delete(allFields, "vpcCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainPlatformData() *DomainPlatformData {
	p := new(DomainPlatformData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DomainPlatformData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DomainPlatformDataProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of clusters in the domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  Domain connectivity status
	*/
	ConnectivityStatus *string `json:"connectivityStatus,omitempty"`
	/*
	  CPU capacity of the domain
	*/
	CpuCount *int `json:"cpuCount,omitempty"`
	/*
	  CPU usage of the domain
	*/
	CpuUsage *float64 `json:"cpuUsage,omitempty"`
	/*
	  User/Service which created this domain
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Number of alerts in the domain
	*/
	CriticalAlertCount *int `json:"criticalAlertCount,omitempty"`
	/*
	  DNS name of domain
	*/
	DnsName *string `json:"dnsName,omitempty"`

	DomainProjection *import6.DomainProjection `json:"domainProjection,omitempty"`
	/*
	  Error messages from domain
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`

	ExtId *string `json:"extId,omitempty"`
	/*
	  Number of hosts in the domain
	*/
	HostCount *int `json:"hostCount,omitempty"`
	/*
	  Domain last synced date
	*/
	LastSyncedAt *time.Time `json:"lastSyncedAt,omitempty"`

	Location *Location `json:"location,omitempty"`
	/*
	  Memory capacity of the domain
	*/
	MemoryCapacity *int64 `json:"memoryCapacity,omitempty"`
	/*
	  Memory usage of the domain
	*/
	MemoryUsage *float64 `json:"memoryUsage,omitempty"`
	/*
	  Name of the registered domain.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  PC Environment Type of the domain
	*/
	PcEnvType *string `json:"pcEnvType,omitempty"`
	/*
	  The IP address of the Prism Central.
	*/
	PcIpAddress *string `json:"pcIpAddress,omitempty"`
	/*
	  PC Provider Type of the domain
	*/
	PcProviderType *string `json:"pcProviderType,omitempty"`

	PcSize *PCSize `json:"pcSize,omitempty"`
	/*
	  PC Version of the domain
	*/
	PcVersion *string `json:"pcVersion,omitempty"`
	/*
	  VM count of Prism Central
	*/
	PcVmCount *int `json:"pcVmCount,omitempty"`
	/*
	  Domain registration date
	*/
	RegisteredAt *time.Time `json:"registeredAt,omitempty"`
	/*
	  Storage capacity of the domain
	*/
	StorageCapacity *int64 `json:"storageCapacity,omitempty"`
	/*
	  Storage usage of the domain
	*/
	StorageUsage *float64 `json:"storageUsage,omitempty"`
	/*
	  Number of VLANs in the domain
	*/
	VlanCount *int `json:"vlanCount,omitempty"`
	/*
	  Number of VMs in the domain
	*/
	VmCount *int `json:"vmCount,omitempty"`
	/*
	  Number of VPCs in the domain
	*/
	VpcCount *int `json:"vpcCount,omitempty"`
}

func (p *DomainPlatformDataProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainPlatformDataProjection

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

func (p *DomainPlatformDataProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainPlatformDataProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainPlatformDataProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.ConnectivityStatus != nil {
		p.ConnectivityStatus = known.ConnectivityStatus
	}
	if known.CpuCount != nil {
		p.CpuCount = known.CpuCount
	}
	if known.CpuUsage != nil {
		p.CpuUsage = known.CpuUsage
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CriticalAlertCount != nil {
		p.CriticalAlertCount = known.CriticalAlertCount
	}
	if known.DnsName != nil {
		p.DnsName = known.DnsName
	}
	if known.DomainProjection != nil {
		p.DomainProjection = known.DomainProjection
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostCount != nil {
		p.HostCount = known.HostCount
	}
	if known.LastSyncedAt != nil {
		p.LastSyncedAt = known.LastSyncedAt
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.MemoryCapacity != nil {
		p.MemoryCapacity = known.MemoryCapacity
	}
	if known.MemoryUsage != nil {
		p.MemoryUsage = known.MemoryUsage
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PcEnvType != nil {
		p.PcEnvType = known.PcEnvType
	}
	if known.PcIpAddress != nil {
		p.PcIpAddress = known.PcIpAddress
	}
	if known.PcProviderType != nil {
		p.PcProviderType = known.PcProviderType
	}
	if known.PcSize != nil {
		p.PcSize = known.PcSize
	}
	if known.PcVersion != nil {
		p.PcVersion = known.PcVersion
	}
	if known.PcVmCount != nil {
		p.PcVmCount = known.PcVmCount
	}
	if known.RegisteredAt != nil {
		p.RegisteredAt = known.RegisteredAt
	}
	if known.StorageCapacity != nil {
		p.StorageCapacity = known.StorageCapacity
	}
	if known.StorageUsage != nil {
		p.StorageUsage = known.StorageUsage
	}
	if known.VlanCount != nil {
		p.VlanCount = known.VlanCount
	}
	if known.VmCount != nil {
		p.VmCount = known.VmCount
	}
	if known.VpcCount != nil {
		p.VpcCount = known.VpcCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterCount")
	delete(allFields, "connectivityStatus")
	delete(allFields, "cpuCount")
	delete(allFields, "cpuUsage")
	delete(allFields, "createdBy")
	delete(allFields, "criticalAlertCount")
	delete(allFields, "dnsName")
	delete(allFields, "domainProjection")
	delete(allFields, "errorMessage")
	delete(allFields, "extId")
	delete(allFields, "hostCount")
	delete(allFields, "lastSyncedAt")
	delete(allFields, "location")
	delete(allFields, "memoryCapacity")
	delete(allFields, "memoryUsage")
	delete(allFields, "name")
	delete(allFields, "pcEnvType")
	delete(allFields, "pcIpAddress")
	delete(allFields, "pcProviderType")
	delete(allFields, "pcSize")
	delete(allFields, "pcVersion")
	delete(allFields, "pcVmCount")
	delete(allFields, "registeredAt")
	delete(allFields, "storageCapacity")
	delete(allFields, "storageUsage")
	delete(allFields, "vlanCount")
	delete(allFields, "vmCount")
	delete(allFields, "vpcCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainPlatformDataProjection() *DomainPlatformDataProjection {
	p := new(DomainPlatformDataProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DomainPlatformDataProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DomainProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Bill of materials ExtID
	*/
	BomExtId *string `json:"bomExtId"`

	BomProjection *BomProjection `json:"bomProjection,omitempty"`
	/*
	  Total number of clusters registered to this domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  List of clusters managed by App Domain PC/ Management PC
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  Cluster Profile External ID is a unique identifier for a cluster profile used to apply common settings to a Domain.
	*/
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
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
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
	delete(allFields, "clusterCount")
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
	*p.ObjectType_ = "multidomain.v4.config.DomainProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains details about the tunnel used for the domain.
*/
type DomainTunnel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Tunnel endpoint for connectivity. Prism Central make use of this endpoint to reach tunnel.
	*/
	Endpoint *string `json:"endpoint,omitempty"`

	Provider *import7.TunnelProvider `json:"provider,omitempty"`
}

func (p *DomainTunnel) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainTunnel

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

func (p *DomainTunnel) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainTunnel
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainTunnel()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Endpoint != nil {
		p.Endpoint = known.Endpoint
	}
	if known.Provider != nil {
		p.Provider = known.Provider
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endpoint")
	delete(allFields, "provider")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainTunnel() *DomainTunnel {
	p := new(DomainTunnel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.DomainTunnel"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
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

type ExternalRepository struct {
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

	LocationItemDiscriminator_ *string `json:"$locationItemDiscriminator,omitempty"`
	/*
	  Location details.
	*/
	Location *OneOfExternalRepositoryLocation `json:"location"`
	/*
	  External repository name.
	*/
	Name *string `json:"name"`
	/*
	  External identifier of the owner of the repository.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ExternalRepository) MarshalJSON() ([]byte, error) {
	type ExternalRepositoryProxy ExternalRepository

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ExternalRepositoryProxy
		Location *OneOfExternalRepositoryLocation `json:"location,omitempty"`
		Name     *string                          `json:"name,omitempty"`
	}{
		ExternalRepositoryProxy: (*ExternalRepositoryProxy)(p),
		Location:                p.Location,
		Name:                    p.Name,
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

func (p *ExternalRepository) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalRepository
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalRepository()

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
	if known.LocationItemDiscriminator_ != nil {
		p.LocationItemDiscriminator_ = known.LocationItemDiscriminator_
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
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
	delete(allFields, "$locationItemDiscriminator")
	delete(allFields, "location")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalRepository() *ExternalRepository {
	p := new(ExternalRepository)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ExternalRepository"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ExternalRepository) GetLocation() interface{} {
	if nil == p.Location {
		return nil
	}
	return p.Location.GetValue()
}

func (p *ExternalRepository) SetLocation(v interface{}) error {
	if nil == p.Location {
		p.Location = NewOneOfExternalRepositoryLocation()
	}
	e := p.Location.SetValue(v)
	if nil == e {
		if nil == p.LocationItemDiscriminator_ {
			p.LocationItemDiscriminator_ = new(string)
		}
		*p.LocationItemDiscriminator_ = *p.Location.Discriminator
	}
	return e
}

/*
The location coordinates, that is latitude and longitude.
*/
type GeographicCoordinates struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The latitude part of the location.
	*/
	Latitude *float64 `json:"latitude,omitempty"`
	/*
	  The longitude part of the location.
	*/
	Longitude *float64 `json:"longitude,omitempty"`
}

func (p *GeographicCoordinates) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GeographicCoordinates

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

func (p *GeographicCoordinates) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GeographicCoordinates
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGeographicCoordinates()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Latitude != nil {
		p.Latitude = known.Latitude
	}
	if known.Longitude != nil {
		p.Longitude = known.Longitude
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "latitude")
	delete(allFields, "longitude")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGeographicCoordinates() *GeographicCoordinates {
	p := new(GeographicCoordinates)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GeographicCoordinates"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/config/external-repositories/{extId} Get operation
*/
type GetExternalRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetExternalRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetExternalRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetExternalRepositoryApiResponse

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

func (p *GetExternalRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetExternalRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetExternalRepositoryApiResponse()

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

func NewGetExternalRepositoryApiResponse() *GetExternalRepositoryApiResponse {
	p := new(GetExternalRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GetExternalRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetExternalRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetExternalRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetExternalRepositoryApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/projects/{extId} Get operation
*/
type GetProjectApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetProjectApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetProjectApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetProjectApiResponse

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

func (p *GetProjectApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetProjectApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetProjectApiResponse()

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

func NewGetProjectApiResponse() *GetProjectApiResponse {
	p := new(GetProjectApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GetProjectApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetProjectApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetProjectApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetProjectApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains/{extId} Get operation
*/
type GetRegisteredDomainApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRegisteredDomainApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetRegisteredDomainApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetRegisteredDomainApiResponse

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

func (p *GetRegisteredDomainApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetRegisteredDomainApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetRegisteredDomainApiResponse()

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

func NewGetRegisteredDomainApiResponse() *GetRegisteredDomainApiResponse {
	p := new(GetRegisteredDomainApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GetRegisteredDomainApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRegisteredDomainApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRegisteredDomainApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRegisteredDomainApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/resource-groups/{extId} Get operation
*/
type GetResourceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetResourceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetResourceGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetResourceGroupApiResponse

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

func (p *GetResourceGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetResourceGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetResourceGroupApiResponse()

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

func NewGetResourceGroupApiResponse() *GetResourceGroupApiResponse {
	p := new(GetResourceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GetResourceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetResourceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetResourceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetResourceGroupApiResponseData()
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
Information about the group view.
*/
type GroupView struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Attribute list for the group view.
	*/
	Attributes []GroupViewAttribute `json:"attributes,omitempty"`

	EntityType *GroupViewEntityType `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name for the group view.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *GroupView) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GroupView

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

func (p *GroupView) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GroupView
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroupView()

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
	delete(allFields, "attributes")
	delete(allFields, "entityType")
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

func NewGroupView() *GroupView {
	p := new(GroupView)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GroupView"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Attribute of a group-view
*/
type GroupViewAttribute struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  This represents the user-facing display name for the view group attribute.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  This represents API model key for the attribute.
	*/
	ModelPropertyName *string `json:"modelPropertyName,omitempty"`
	/*

	 */
	NameItemDiscriminator_ *string `json:"$nameItemDiscriminator,omitempty"`
	/*
	  This represents the name for the group attribute used in the database table.
	*/
	Name *OneOfGroupViewAttributeName `json:"name"`
	/*
	  This represents the parent model name for the attribute.
	*/
	ParentModelName *string `json:"parentModelName,omitempty"`
}

func (p *GroupViewAttribute) MarshalJSON() ([]byte, error) {
	type GroupViewAttributeProxy GroupViewAttribute

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*GroupViewAttributeProxy
		Name *OneOfGroupViewAttributeName `json:"name,omitempty"`
	}{
		GroupViewAttributeProxy: (*GroupViewAttributeProxy)(p),
		Name:                    p.Name,
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

func (p *GroupViewAttribute) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GroupViewAttribute
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroupViewAttribute()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DisplayName != nil {
		p.DisplayName = known.DisplayName
	}
	if known.ModelPropertyName != nil {
		p.ModelPropertyName = known.ModelPropertyName
	}
	if known.NameItemDiscriminator_ != nil {
		p.NameItemDiscriminator_ = known.NameItemDiscriminator_
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ParentModelName != nil {
		p.ParentModelName = known.ParentModelName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "displayName")
	delete(allFields, "modelPropertyName")
	delete(allFields, "$nameItemDiscriminator")
	delete(allFields, "name")
	delete(allFields, "parentModelName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGroupViewAttribute() *GroupViewAttribute {
	p := new(GroupViewAttribute)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.GroupViewAttribute"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GroupViewAttribute) GetName() interface{} {
	if nil == p.Name {
		return nil
	}
	return p.Name.GetValue()
}

func (p *GroupViewAttribute) SetName(v interface{}) error {
	if nil == p.Name {
		p.Name = NewOneOfGroupViewAttributeName()
	}
	e := p.Name.SetValue(v)
	if nil == e {
		if nil == p.NameItemDiscriminator_ {
			p.NameItemDiscriminator_ = new(string)
		}
		*p.NameItemDiscriminator_ = *p.Name.Discriminator
	}
	return e
}

/*
Name for the entity to which group view belongs.
*/
type GroupViewEntityType int

const (
	GROUPVIEWENTITYTYPE_UNKNOWN  GroupViewEntityType = 0
	GROUPVIEWENTITYTYPE_REDACTED GroupViewEntityType = 1
	GROUPVIEWENTITYTYPE_DOMAIN   GroupViewEntityType = 2
	GROUPVIEWENTITYTYPE_CLUSTER  GroupViewEntityType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *GroupViewEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e GroupViewEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN",
		"CLUSTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *GroupViewEntityType) index(name string) GroupViewEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN",
		"CLUSTER",
	}
	for idx := range names {
		if names[idx] == name {
			return GroupViewEntityType(idx)
		}
	}
	return GROUPVIEWENTITYTYPE_UNKNOWN
}

func (e *GroupViewEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GroupViewEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GroupViewEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GroupViewEntityType) Ref() *GroupViewEntityType {
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
	*p.ObjectType_ = "multidomain.v4.config.HookDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/config/external-repositories Get operation
*/
type ListExternalRepositoriesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListExternalRepositoriesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListExternalRepositoriesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListExternalRepositoriesApiResponse

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

func (p *ListExternalRepositoriesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListExternalRepositoriesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListExternalRepositoriesApiResponse()

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

func NewListExternalRepositoriesApiResponse() *ListExternalRepositoriesApiResponse {
	p := new(ListExternalRepositoriesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ListExternalRepositoriesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListExternalRepositoriesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListExternalRepositoriesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListExternalRepositoriesApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/locations Get operation
*/
type ListLocationsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLocationsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListLocationsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListLocationsApiResponse

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

func (p *ListLocationsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListLocationsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListLocationsApiResponse()

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

func NewListLocationsApiResponse() *ListLocationsApiResponse {
	p := new(ListLocationsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ListLocationsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLocationsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLocationsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLocationsApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/projects Get operation
*/
type ListProjectsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListProjectsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListProjectsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListProjectsApiResponse

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

func (p *ListProjectsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListProjectsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListProjectsApiResponse()

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

func NewListProjectsApiResponse() *ListProjectsApiResponse {
	p := new(ListProjectsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ListProjectsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListProjectsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListProjectsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListProjectsApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains Get operation
*/
type ListRegisteredDomainsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListRegisteredDomainsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListRegisteredDomainsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListRegisteredDomainsApiResponse

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

func (p *ListRegisteredDomainsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListRegisteredDomainsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListRegisteredDomainsApiResponse()

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

func NewListRegisteredDomainsApiResponse() *ListRegisteredDomainsApiResponse {
	p := new(ListRegisteredDomainsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ListRegisteredDomainsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListRegisteredDomainsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListRegisteredDomainsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListRegisteredDomainsApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/resource-groups Get operation
*/
type ListResourceGroupsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListResourceGroupsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListResourceGroupsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListResourceGroupsApiResponse

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

func (p *ListResourceGroupsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListResourceGroupsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListResourceGroupsApiResponse()

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

func NewListResourceGroupsApiResponse() *ListResourceGroupsApiResponse {
	p := new(ListResourceGroupsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ListResourceGroupsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListResourceGroupsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListResourceGroupsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListResourceGroupsApiResponseData()
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
Location information.
*/
type Location struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Coordinates *GeographicCoordinates `json:"coordinates,omitempty"`
	/*
	  Location Country.
	*/
	Country *string `json:"country,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Location state.
	*/
	State *string `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Location) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Location

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

func (p *Location) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Location
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Coordinates != nil {
		p.Coordinates = known.Coordinates
	}
	if known.Country != nil {
		p.Country = known.Country
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
	delete(allFields, "coordinates")
	delete(allFields, "country")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocation() *Location {
	p := new(Location)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.Location"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type LocationProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Coordinates *GeographicCoordinates `json:"coordinates,omitempty"`
	/*
	  Location Country.
	*/
	Country *string `json:"country,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location name
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Location state.
	*/
	State *string `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *LocationProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LocationProjection

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

func (p *LocationProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LocationProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLocationProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Coordinates != nil {
		p.Coordinates = known.Coordinates
	}
	if known.Country != nil {
		p.Country = known.Country
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
	delete(allFields, "coordinates")
	delete(allFields, "country")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "state")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLocationProjection() *LocationProjection {
	p := new(LocationProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.LocationProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NFS server address details in FQDN:Port Number format.
*/
type NFSServerAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import8.IPAddressOrFQDN `json:"address"`
	/*
	  Server port number.
	*/
	Port *int `json:"port"`
}

func (p *NFSServerAddress) MarshalJSON() ([]byte, error) {
	type NFSServerAddressProxy NFSServerAddress

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NFSServerAddressProxy
		Address *import8.IPAddressOrFQDN `json:"address,omitempty"`
		Port    *int                     `json:"port,omitempty"`
	}{
		NFSServerAddressProxy: (*NFSServerAddressProxy)(p),
		Address:               p.Address,
		Port:                  p.Port,
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

func (p *NFSServerAddress) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NFSServerAddress
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNFSServerAddress()

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
	if known.Port != nil {
		p.Port = known.Port
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "address")
	delete(allFields, "port")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNFSServerAddress() *NFSServerAddress {
	p := new(NFSServerAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.NFSServerAddress"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NFS configuration details.
*/
type NfsRepository struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of NFS server addresses.
	*/
	Addresses []NFSServerAddress `json:"addresses"`
	/*
	  NFS repository export name details.
	*/
	ExportName *string `json:"exportName"`
}

func (p *NfsRepository) MarshalJSON() ([]byte, error) {
	type NfsRepositoryProxy NfsRepository

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NfsRepositoryProxy
		Addresses  []NFSServerAddress `json:"addresses,omitempty"`
		ExportName *string            `json:"exportName,omitempty"`
	}{
		NfsRepositoryProxy: (*NfsRepositoryProxy)(p),
		Addresses:          p.Addresses,
		ExportName:         p.ExportName,
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

func (p *NfsRepository) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NfsRepository
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNfsRepository()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Addresses != nil {
		p.Addresses = known.Addresses
	}
	if known.ExportName != nil {
		p.ExportName = known.ExportName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addresses")
	delete(allFields, "exportName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNfsRepository() *NfsRepository {
	p := new(NfsRepository)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.NfsRepository"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
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
The size of the Prism Central.
*/
type PCSize int

const (
	PCSIZE_UNKNOWN  PCSize = 0
	PCSIZE_REDACTED PCSize = 1
	PCSIZE_XSMALL   PCSize = 2
	PCSIZE_SMALL    PCSize = 3
	PCSIZE_LARGE    PCSize = 4
	PCSIZE_XLARGE   PCSize = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PCSize) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XSMALL",
		"SMALL",
		"LARGE",
		"XLARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PCSize) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XSMALL",
		"SMALL",
		"LARGE",
		"XLARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PCSize) index(name string) PCSize {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"XSMALL",
		"SMALL",
		"LARGE",
		"XLARGE",
	}
	for idx := range names {
		if names[idx] == name {
			return PCSize(idx)
		}
	}
	return PCSIZE_UNKNOWN
}

func (e *PCSize) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PCSize:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PCSize) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PCSize) Ref() *PCSize {
	return &e
}

/*
Platform data for the domain.
*/
type PlatformData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of clusters in the domain
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  Domain connectivity status
	*/
	ConnectivityStatus *string `json:"connectivityStatus,omitempty"`
	/*
	  PC Environment Type of the domain
	*/
	PcEnvType *string `json:"pcEnvType,omitempty"`
	/*
	  PC Provider Type of the domain
	*/
	PcProviderType *string `json:"pcProviderType,omitempty"`
	/*
	  PC Version of the domain
	*/
	PcVersion *string `json:"pcVersion,omitempty"`
}

func (p *PlatformData) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PlatformData

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

func (p *PlatformData) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PlatformData
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPlatformData()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.ConnectivityStatus != nil {
		p.ConnectivityStatus = known.ConnectivityStatus
	}
	if known.PcEnvType != nil {
		p.PcEnvType = known.PcEnvType
	}
	if known.PcProviderType != nil {
		p.PcProviderType = known.PcProviderType
	}
	if known.PcVersion != nil {
		p.PcVersion = known.PcVersion
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterCount")
	delete(allFields, "connectivityStatus")
	delete(allFields, "pcEnvType")
	delete(allFields, "pcProviderType")
	delete(allFields, "pcVersion")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPlatformData() *PlatformData {
	p := new(PlatformData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.PlatformData"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A logical grouping construct
*/
type Project struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The ID of the user who created the project
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The timestamp when the project was created
	*/
	CreatedTimestamp *int64 `json:"createdTimestamp,omitempty"`
	/*
	  Description of the project
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  ID of the project. This field is immutable.
	*/
	Id *string `json:"id"`
	/*
	  Indicates if the project is the default project
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  Indicates if the project is system defined
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The timestamp when the project was last modified
	*/
	ModifiedTimestamp *int64 `json:"modifiedTimestamp,omitempty"`
	/*
	  Name of the project. This field is immutable.
	*/
	Name *string `json:"name"`

	State *ProjectState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The ID of the user who last updated the project
	*/
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type ProjectProxy Project

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProjectProxy
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		ProjectProxy: (*ProjectProxy)(p),
		Id:           p.Id,
		Name:         p.Name,
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

func (p *Project) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Project
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProject()

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
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Id != nil {
		p.Id = known.Id
	}
	if known.IsDefault != nil {
		p.IsDefault = known.IsDefault
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ModifiedTimestamp != nil {
		p.ModifiedTimestamp = known.ModifiedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpdatedBy != nil {
		p.UpdatedBy = known.UpdatedBy
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "createdTimestamp")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "id")
	delete(allFields, "isDefault")
	delete(allFields, "isSystemDefined")
	delete(allFields, "links")
	delete(allFields, "modifiedTimestamp")
	delete(allFields, "name")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "updatedBy")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProject() *Project {
	p := new(Project)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.Project"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ProjectProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The ID of the user who created the project
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The timestamp when the project was created
	*/
	CreatedTimestamp *int64 `json:"createdTimestamp,omitempty"`
	/*
	  Description of the project
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  ID of the project. This field is immutable.
	*/
	Id *string `json:"id"`
	/*
	  Indicates if the project is the default project
	*/
	IsDefault *bool `json:"isDefault,omitempty"`
	/*
	  Indicates if the project is system defined
	*/
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The timestamp when the project was last modified
	*/
	ModifiedTimestamp *int64 `json:"modifiedTimestamp,omitempty"`
	/*
	  Name of the project. This field is immutable.
	*/
	Name *string `json:"name"`

	State *ProjectState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The ID of the user who last updated the project
	*/
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

func (p *ProjectProjection) MarshalJSON() ([]byte, error) {
	type ProjectProjectionProxy ProjectProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProjectProjectionProxy
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		ProjectProjectionProxy: (*ProjectProjectionProxy)(p),
		Id:                     p.Id,
		Name:                   p.Name,
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

func (p *ProjectProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProjectProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProjectProjection()

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
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Id != nil {
		p.Id = known.Id
	}
	if known.IsDefault != nil {
		p.IsDefault = known.IsDefault
	}
	if known.IsSystemDefined != nil {
		p.IsSystemDefined = known.IsSystemDefined
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ModifiedTimestamp != nil {
		p.ModifiedTimestamp = known.ModifiedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpdatedBy != nil {
		p.UpdatedBy = known.UpdatedBy
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createdBy")
	delete(allFields, "createdTimestamp")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "id")
	delete(allFields, "isDefault")
	delete(allFields, "isSystemDefined")
	delete(allFields, "links")
	delete(allFields, "modifiedTimestamp")
	delete(allFields, "name")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "updatedBy")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProjectProjection() *ProjectProjection {
	p := new(ProjectProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ProjectProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The state of the project. Possible values are: ACTIVE, DELETING
*/
type ProjectState int

const (
	PROJECTSTATE_UNKNOWN  ProjectState = 0
	PROJECTSTATE_REDACTED ProjectState = 1
	PROJECTSTATE_ACTIVE   ProjectState = 2
	PROJECTSTATE_DELETING ProjectState = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProjectState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"DELETING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProjectState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"DELETING",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProjectState) index(name string) ProjectState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"DELETING",
	}
	for idx := range names {
		if names[idx] == name {
			return ProjectState(idx)
		}
	}
	return PROJECTSTATE_UNKNOWN
}

func (e *ProjectState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProjectState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProjectState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProjectState) Ref() *ProjectState {
	return &e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains/{extId}/$actions/refresh-api-credentials Post operation
*/
type RefreshApiCredentialsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRefreshApiCredentialsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *RefreshApiCredentialsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RefreshApiCredentialsApiResponse

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

func (p *RefreshApiCredentialsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RefreshApiCredentialsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRefreshApiCredentialsApiResponse()

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

func NewRefreshApiCredentialsApiResponse() *RefreshApiCredentialsApiResponse {
	p := new(RefreshApiCredentialsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RefreshApiCredentialsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RefreshApiCredentialsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RefreshApiCredentialsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRefreshApiCredentialsApiResponseData()
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
Information of the registered domain.
*/
type RegisteredDomain struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApiCredentialStatus *ApiCredentialStatus `json:"apiCredentialStatus,omitempty"`
	/*
	  The user that created the registered domain on Nutanix Central.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The entity that issued the credentials for the domain to use as authentication for Nutanix central communication.
	*/
	CredentialIssuer *string `json:"credentialIssuer,omitempty"`
	/*
	  DNS name of the registered domain.
	*/
	DnsName *string `json:"dnsName,omitempty"`
	/*
	  The external ID of the domain
	*/
	DomainExtId *string `json:"domainExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Fully qualified domain name  of the registered domain.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location UUID of the domain.
	*/
	LocationExtId *string `json:"locationExtId,omitempty"`
	/*
	  Name of the registered domain.
	*/
	Name *string `json:"name,omitempty"`

	PlatformData *PlatformData `json:"platformData,omitempty"`
	/*
	  Date and time of domain registration.
	*/
	RegisteredTime *time.Time `json:"registeredTime,omitempty"`

	RegistrationConfig *RegistrationConfig `json:"registrationConfig,omitempty"`

	RegistrationState *import7.RegistrationState `json:"registrationState,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Tunnel *DomainTunnel `json:"tunnel,omitempty"`
}

func (p *RegisteredDomain) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RegisteredDomain

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

func (p *RegisteredDomain) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegisteredDomain
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegisteredDomain()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApiCredentialStatus != nil {
		p.ApiCredentialStatus = known.ApiCredentialStatus
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CredentialIssuer != nil {
		p.CredentialIssuer = known.CredentialIssuer
	}
	if known.DnsName != nil {
		p.DnsName = known.DnsName
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationExtId != nil {
		p.LocationExtId = known.LocationExtId
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PlatformData != nil {
		p.PlatformData = known.PlatformData
	}
	if known.RegisteredTime != nil {
		p.RegisteredTime = known.RegisteredTime
	}
	if known.RegistrationConfig != nil {
		p.RegistrationConfig = known.RegistrationConfig
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Tunnel != nil {
		p.Tunnel = known.Tunnel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiCredentialStatus")
	delete(allFields, "createdBy")
	delete(allFields, "credentialIssuer")
	delete(allFields, "dnsName")
	delete(allFields, "domainExtId")
	delete(allFields, "extId")
	delete(allFields, "fqdn")
	delete(allFields, "links")
	delete(allFields, "locationExtId")
	delete(allFields, "name")
	delete(allFields, "platformData")
	delete(allFields, "registeredTime")
	delete(allFields, "registrationConfig")
	delete(allFields, "registrationState")
	delete(allFields, "tenantId")
	delete(allFields, "tunnel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegisteredDomain() *RegisteredDomain {
	p := new(RegisteredDomain)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RegisteredDomain"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Request payload to update state of the registered domain.
*/
type RegisteredDomainChangeStateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RegistrationState *import7.RegistrationState `json:"registrationState"`
}

func (p *RegisteredDomainChangeStateSpec) MarshalJSON() ([]byte, error) {
	type RegisteredDomainChangeStateSpecProxy RegisteredDomainChangeStateSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RegisteredDomainChangeStateSpecProxy
		RegistrationState *import7.RegistrationState `json:"registrationState,omitempty"`
	}{
		RegisteredDomainChangeStateSpecProxy: (*RegisteredDomainChangeStateSpecProxy)(p),
		RegistrationState:                    p.RegistrationState,
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

func (p *RegisteredDomainChangeStateSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegisteredDomainChangeStateSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegisteredDomainChangeStateSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "registrationState")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegisteredDomainChangeStateSpec() *RegisteredDomainChangeStateSpec {
	p := new(RegisteredDomainChangeStateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RegisteredDomainChangeStateSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RegisteredDomainProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApiCredentialStatus *ApiCredentialStatus `json:"apiCredentialStatus,omitempty"`
	/*
	  The user that created the registered domain on Nutanix Central.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  The entity that issued the credentials for the domain to use as authentication for Nutanix central communication.
	*/
	CredentialIssuer *string `json:"credentialIssuer,omitempty"`
	/*
	  DNS name of the registered domain.
	*/
	DnsName *string `json:"dnsName,omitempty"`
	/*
	  The external ID of the domain
	*/
	DomainExtId *string `json:"domainExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Fully qualified domain name  of the registered domain.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Location UUID of the domain.
	*/
	LocationExtId *string `json:"locationExtId,omitempty"`

	LocationProjection *LocationProjection `json:"locationProjection,omitempty"`
	/*
	  Name of the registered domain.
	*/
	Name *string `json:"name,omitempty"`

	PlatformData *PlatformData `json:"platformData,omitempty"`
	/*
	  Date and time of domain registration.
	*/
	RegisteredTime *time.Time `json:"registeredTime,omitempty"`

	RegistrationConfig *RegistrationConfig `json:"registrationConfig,omitempty"`

	RegistrationState *import7.RegistrationState `json:"registrationState,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Tunnel *DomainTunnel `json:"tunnel,omitempty"`
}

func (p *RegisteredDomainProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RegisteredDomainProjection

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

func (p *RegisteredDomainProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegisteredDomainProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegisteredDomainProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ApiCredentialStatus != nil {
		p.ApiCredentialStatus = known.ApiCredentialStatus
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.CredentialIssuer != nil {
		p.CredentialIssuer = known.CredentialIssuer
	}
	if known.DnsName != nil {
		p.DnsName = known.DnsName
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LocationExtId != nil {
		p.LocationExtId = known.LocationExtId
	}
	if known.LocationProjection != nil {
		p.LocationProjection = known.LocationProjection
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PlatformData != nil {
		p.PlatformData = known.PlatformData
	}
	if known.RegisteredTime != nil {
		p.RegisteredTime = known.RegisteredTime
	}
	if known.RegistrationConfig != nil {
		p.RegistrationConfig = known.RegistrationConfig
	}
	if known.RegistrationState != nil {
		p.RegistrationState = known.RegistrationState
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Tunnel != nil {
		p.Tunnel = known.Tunnel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "apiCredentialStatus")
	delete(allFields, "createdBy")
	delete(allFields, "credentialIssuer")
	delete(allFields, "dnsName")
	delete(allFields, "domainExtId")
	delete(allFields, "extId")
	delete(allFields, "fqdn")
	delete(allFields, "links")
	delete(allFields, "locationExtId")
	delete(allFields, "locationProjection")
	delete(allFields, "name")
	delete(allFields, "platformData")
	delete(allFields, "registeredTime")
	delete(allFields, "registrationConfig")
	delete(allFields, "registrationState")
	delete(allFields, "tenantId")
	delete(allFields, "tunnel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegisteredDomainProjection() *RegisteredDomainProjection {
	p := new(RegisteredDomainProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RegisteredDomainProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Registered domain trust request payload
*/
type RegisteredDomainTrustSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The request value signed by the client for domain registration.
	*/
	ClientSigningRequest *string `json:"clientSigningRequest"`
}

func (p *RegisteredDomainTrustSpec) MarshalJSON() ([]byte, error) {
	type RegisteredDomainTrustSpecProxy RegisteredDomainTrustSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RegisteredDomainTrustSpecProxy
		ClientSigningRequest *string `json:"clientSigningRequest,omitempty"`
	}{
		RegisteredDomainTrustSpecProxy: (*RegisteredDomainTrustSpecProxy)(p),
		ClientSigningRequest:           p.ClientSigningRequest,
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

func (p *RegisteredDomainTrustSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegisteredDomainTrustSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegisteredDomainTrustSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClientSigningRequest != nil {
		p.ClientSigningRequest = known.ClientSigningRequest
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clientSigningRequest")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegisteredDomainTrustSpec() *RegisteredDomainTrustSpec {
	p := new(RegisteredDomainTrustSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RegisteredDomainTrustSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Configuration information for the domain registration.
*/
type RegistrationConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credentials *import7.RegistrationCredentials `json:"credentials,omitempty"`
	/*
	  Unique identifier for the Prism Central.
	*/
	DomainExtId *string `json:"domainExtId,omitempty"`
	/*
	  Prism Central should use this URL to communicate with Nutanix Central.
	*/
	TargetUrl *string `json:"targetUrl,omitempty"`
	/*
	  Unique identifier of the tenant.
	*/
	TenantExtId *string `json:"tenantExtId,omitempty"`
}

func (p *RegistrationConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RegistrationConfig

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

func (p *RegistrationConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegistrationConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRegistrationConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Credentials != nil {
		p.Credentials = known.Credentials
	}
	if known.DomainExtId != nil {
		p.DomainExtId = known.DomainExtId
	}
	if known.TargetUrl != nil {
		p.TargetUrl = known.TargetUrl
	}
	if known.TenantExtId != nil {
		p.TenantExtId = known.TenantExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credentials")
	delete(allFields, "domainExtId")
	delete(allFields, "targetUrl")
	delete(allFields, "tenantExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegistrationConfig() *RegistrationConfig {
	p := new(RegistrationConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.RegistrationConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Resource Group entity with its attributes.
*/
type ResourceGroup struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Capabilities and features for this Resource Group.
	*/
	Capabilities []import8.KVPair `json:"capabilities,omitempty"`
	/*
	  The time when the resource group was created in ISO 8601 format.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  UUID of the user who created this Resource Group.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the resource group was last updated in ISO 8601 format.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  UUID of the user who last updated this Resource Group.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Resource Group.
	*/
	Name *string `json:"name"`
	/*
	  List of placement targets that define where resources from this Resource Group can be deployed.
	*/
	PlacementTargets []TargetDetails `json:"placementTargets,omitempty"`
	/*
	  UUID of the project that owns this Resource Group.
	*/
	ProjectExtId *string `json:"projectExtId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ResourceGroup) MarshalJSON() ([]byte, error) {
	type ResourceGroupProxy ResourceGroup

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ResourceGroupProxy
		Name         *string `json:"name,omitempty"`
		ProjectExtId *string `json:"projectExtId,omitempty"`
	}{
		ResourceGroupProxy: (*ResourceGroupProxy)(p),
		Name:               p.Name,
		ProjectExtId:       p.ProjectExtId,
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

func (p *ResourceGroup) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResourceGroup
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResourceGroup()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PlacementTargets != nil {
		p.PlacementTargets = known.PlacementTargets
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "extId")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "placementTargets")
	delete(allFields, "projectExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewResourceGroup() *ResourceGroup {
	p := new(ResourceGroup)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ResourceGroup"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ResourceGroupProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Capabilities and features for this Resource Group.
	*/
	Capabilities []import8.KVPair `json:"capabilities,omitempty"`
	/*
	  The time when the resource group was created in ISO 8601 format.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  UUID of the user who created this Resource Group.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The time when the resource group was last updated in ISO 8601 format.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  UUID of the user who last updated this Resource Group.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the Resource Group.
	*/
	Name *string `json:"name"`
	/*
	  List of placement targets that define where resources from this Resource Group can be deployed.
	*/
	PlacementTargets []TargetDetails `json:"placementTargets,omitempty"`
	/*
	  UUID of the project that owns this Resource Group.
	*/
	ProjectExtId *string `json:"projectExtId"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ResourceGroupProjection) MarshalJSON() ([]byte, error) {
	type ResourceGroupProjectionProxy ResourceGroupProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ResourceGroupProjectionProxy
		Name         *string `json:"name,omitempty"`
		ProjectExtId *string `json:"projectExtId,omitempty"`
	}{
		ResourceGroupProjectionProxy: (*ResourceGroupProjectionProxy)(p),
		Name:                         p.Name,
		ProjectExtId:                 p.ProjectExtId,
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

func (p *ResourceGroupProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResourceGroupProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResourceGroupProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.PlacementTargets != nil {
		p.PlacementTargets = known.PlacementTargets
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "extId")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "placementTargets")
	delete(allFields, "projectExtId")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewResourceGroupProjection() *ResourceGroupProjection {
	p := new(ResourceGroupProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.ResourceGroupProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

/*
Storage container that provides persistent storage resources.
*/
type StorageContainerDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Capabilities and features of the storage container.
	*/
	Capabilities []import8.KVPair `json:"capabilities,omitempty"`
	/*
	  UUID of the storage container.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *StorageContainerDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StorageContainerDetails

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

func (p *StorageContainerDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageContainerDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageContainerDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageContainerDetails() *StorageContainerDetails {
	p := new(StorageContainerDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.StorageContainerDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Target details containing cluster configuration and storage resources.
*/
type TargetDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Capabilities and features available at this Placement Target.
	*/
	Capabilities []import8.KVPair `json:"capabilities,omitempty"`
	/*
	  UUID of the AOS cluster.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  List of storage containers available for this cluster target.
	*/
	StorageContainers []StorageContainerDetails `json:"storageContainers,omitempty"`
}

func (p *TargetDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TargetDetails

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

func (p *TargetDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TargetDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTargetDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capabilities != nil {
		p.Capabilities = known.Capabilities
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.StorageContainers != nil {
		p.StorageContainers = known.StorageContainers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capabilities")
	delete(allFields, "clusterExtId")
	delete(allFields, "storageContainers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTargetDetails() *TargetDetails {
	p := new(TargetDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.TargetDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/config/external-repositories/{extId} Put operation
*/
type UpdateExternalRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateExternalRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateExternalRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateExternalRepositoryApiResponse

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

func (p *UpdateExternalRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateExternalRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateExternalRepositoryApiResponse()

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

func NewUpdateExternalRepositoryApiResponse() *UpdateExternalRepositoryApiResponse {
	p := new(UpdateExternalRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.UpdateExternalRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateExternalRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateExternalRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateExternalRepositoryApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/projects/{extId} Put operation
*/
type UpdateProjectApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateProjectApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateProjectApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateProjectApiResponse

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

func (p *UpdateProjectApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateProjectApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateProjectApiResponse()

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

func NewUpdateProjectApiResponse() *UpdateProjectApiResponse {
	p := new(UpdateProjectApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.UpdateProjectApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateProjectApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateProjectApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateProjectApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/registered-domains/{extId} Put operation
*/
type UpdateRegisteredDomainApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRegisteredDomainApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateRegisteredDomainApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateRegisteredDomainApiResponse

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

func (p *UpdateRegisteredDomainApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateRegisteredDomainApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateRegisteredDomainApiResponse()

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

func NewUpdateRegisteredDomainApiResponse() *UpdateRegisteredDomainApiResponse {
	p := new(UpdateRegisteredDomainApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.UpdateRegisteredDomainApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateRegisteredDomainApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateRegisteredDomainApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateRegisteredDomainApiResponseData()
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
REST response for all response codes in API path /multidomain/v4.4.b1/config/resource-groups/{extId} Put operation
*/
type UpdateResourceGroupApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateResourceGroupApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateResourceGroupApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateResourceGroupApiResponse

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

func (p *UpdateResourceGroupApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateResourceGroupApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateResourceGroupApiResponse()

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

func NewUpdateResourceGroupApiResponse() *UpdateResourceGroupApiResponse {
	p := new(UpdateResourceGroupApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.config.UpdateResourceGroupApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateResourceGroupApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateResourceGroupApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateResourceGroupApiResponseData()
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

type OneOfCreateRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateRegisteredDomainApiResponseData() *OneOfCreateRegisteredDomainApiResponseData {
	p := new(OneOfCreateRegisteredDomainApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateRegisteredDomainApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateRegisteredDomainApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRegisteredDomainApiResponseData"))
}

func (p *OneOfCreateRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRegisteredDomainApiResponseData")
}

type OneOfListExternalRepositoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []ExternalRepository   `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListExternalRepositoriesApiResponseData() *OneOfListExternalRepositoriesApiResponseData {
	p := new(OneOfListExternalRepositoriesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListExternalRepositoriesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListExternalRepositoriesApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ExternalRepository:
		p.oneOfType2001 = v.([]ExternalRepository)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.ExternalRepository>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.ExternalRepository>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListExternalRepositoriesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.ExternalRepository>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListExternalRepositoriesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.ExternalRepository>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]ExternalRepository)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.config.ExternalRepository" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.ExternalRepository>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.ExternalRepository>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType2001 := new([]ExternalRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.config.ExternalRepository" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.ExternalRepository>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.ExternalRepository>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListExternalRepositoriesApiResponseData"))
}

func (p *OneOfListExternalRepositoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.ExternalRepository>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListExternalRepositoriesApiResponseData")
}

type OneOfUpdateRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateRegisteredDomainApiResponseData() *OneOfUpdateRegisteredDomainApiResponseData {
	p := new(OneOfUpdateRegisteredDomainApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRegisteredDomainApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRegisteredDomainApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRegisteredDomainApiResponseData"))
}

func (p *OneOfUpdateRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRegisteredDomainApiResponseData")
}

type OneOfGetResourceGroupApiResponseData struct {
	Discriminator *string                  `json:"-"`
	ObjectType_   *string                  `json:"-"`
	oneOfType400  *import4.ErrorResponse   `json:"-"`
	oneOfType401  *ResourceGroupProjection `json:"-"`
	oneOfType2001 *ResourceGroup           `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetResourceGroupApiResponseData() *OneOfGetResourceGroupApiResponseData {
	p := new(OneOfGetResourceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetResourceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetResourceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ResourceGroupProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(ResourceGroupProjection)
		}
		*p.oneOfType401 = v.(ResourceGroupProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	case ResourceGroup:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ResourceGroup)
		}
		*p.oneOfType2001 = v.(ResourceGroup)
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

func (p *OneOfGetResourceGroupApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetResourceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new(ResourceGroupProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.ResourceGroupProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(ResourceGroupProjection)
							}
							*p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType401.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(ResourceGroup)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.config.ResourceGroup" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(ResourceGroup)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType401 := new(ResourceGroupProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.ResourceGroupProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(ResourceGroupProjection)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(ResourceGroup)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.config.ResourceGroup" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ResourceGroup)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetResourceGroupApiResponseData"))
}

func (p *OneOfGetResourceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetResourceGroupApiResponseData")
}

type OneOfDomainProjectionDomainManager struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType12002 *import5.RemoteCluster `json:"-"`
	oneOfType12001 *import3.DomainManager `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
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
	case import5.RemoteCluster:
		if nil == p.oneOfType12002 {
			p.oneOfType12002 = new(import5.RemoteCluster)
		}
		*p.oneOfType12002 = v.(import5.RemoteCluster)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12002.ObjectType_
	case import3.DomainManager:
		if nil == p.oneOfType12001 {
			p.oneOfType12001 = new(import3.DomainManager)
		}
		*p.oneOfType12001 = v.(import3.DomainManager)
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
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12002
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12001
	}
	return nil
}

func (p *OneOfDomainProjectionDomainManager) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType12002 := new(import5.RemoteCluster)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType12002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType12002.ObjectType_ != nil && "prism.v4.management.RemoteCluster" == *vOneOfType12002.ObjectType_ {
							if nil == p.oneOfType12002 {
								p.oneOfType12002 = new(import5.RemoteCluster)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType12001 := new(import3.DomainManager)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType12001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType12001.ObjectType_ != nil && "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
							if nil == p.oneOfType12001 {
								p.oneOfType12001 = new(import3.DomainManager)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType12002 := new(import5.RemoteCluster)
	if err := json.Unmarshal(b, vOneOfType12002); err == nil {
		if vOneOfType12002.ObjectType_ != nil && "prism.v4.management.RemoteCluster" == *vOneOfType12002.ObjectType_ {
			if nil == p.oneOfType12002 {
				p.oneOfType12002 = new(import5.RemoteCluster)
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
	vOneOfType12001 := new(import3.DomainManager)
	if err := json.Unmarshal(b, vOneOfType12001); err == nil {
		if vOneOfType12001.ObjectType_ != nil && "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
			if nil == p.oneOfType12001 {
				p.oneOfType12001 = new(import3.DomainManager)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainProjectionDomainManager"))
}

func (p *OneOfDomainProjectionDomainManager) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12002)
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12001)
	}
	return nil, errors.New("No value to marshal for OneOfDomainProjectionDomainManager")
}

type OneOfGroupViewAttributeName struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType1    *DomainAttributeName  `json:"-"`
	oneOfType0    *ClusterAttributeName `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGroupViewAttributeName() *OneOfGroupViewAttributeName {
	p := new(OneOfGroupViewAttributeName)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGroupViewAttributeName) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGroupViewAttributeName is nil"))
	}
	switch v.(type) {
	case DomainAttributeName:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(DomainAttributeName)
		}
		*p.oneOfType1 = v.(DomainAttributeName)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "multidomain.v4.config.DomainAttributeName"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "multidomain.v4.config.DomainAttributeName"
	case ClusterAttributeName:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterAttributeName)
		}
		*p.oneOfType0 = v.(ClusterAttributeName)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "multidomain.v4.config.ClusterAttributeName"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "multidomain.v4.config.ClusterAttributeName"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGroupViewAttributeName) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "multidomain.v4.config.DomainAttributeName" == *p.Discriminator {
		return *p.oneOfType1
	}
	if "multidomain.v4.config.ClusterAttributeName" == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGroupViewAttributeName) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["multidomain.v4.config.DomainAttributeName"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1 := new(DomainAttributeName)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1)
					if unmarshalErr == nil {
						if nil == p.oneOfType1 {
							p.oneOfType1 = new(DomainAttributeName)
						}
						*p.oneOfType1 = *vOneOfType1
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "multidomain.v4.config.DomainAttributeName"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "multidomain.v4.config.DomainAttributeName"
						return nil
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["multidomain.v4.config.ClusterAttributeName"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(ClusterAttributeName)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						if nil == p.oneOfType0 {
							p.oneOfType0 = new(ClusterAttributeName)
						}
						*p.oneOfType0 = *vOneOfType0
						if nil == p.Discriminator {
							p.Discriminator = new(string)
						}
						*p.Discriminator = "multidomain.v4.config.ClusterAttributeName"
						if nil == p.ObjectType_ {
							p.ObjectType_ = new(string)
						}
						*p.ObjectType_ = "multidomain.v4.config.ClusterAttributeName"
						return nil
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType1 := new(DomainAttributeName)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(DomainAttributeName)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "multidomain.v4.config.DomainAttributeName"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "multidomain.v4.config.DomainAttributeName"
		return nil
	}
	vOneOfType0 := new(ClusterAttributeName)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterAttributeName)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "multidomain.v4.config.ClusterAttributeName"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "multidomain.v4.config.ClusterAttributeName"
		return nil
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGroupViewAttributeName"))
}

func (p *OneOfGroupViewAttributeName) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "multidomain.v4.config.DomainAttributeName" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if "multidomain.v4.config.ClusterAttributeName" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGroupViewAttributeName")
}

type OneOfDeleteExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteExternalRepositoryApiResponseData() *OneOfDeleteExternalRepositoryApiResponseData {
	p := new(OneOfDeleteExternalRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteExternalRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteExternalRepositoryApiResponseData"))
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteExternalRepositoryApiResponseData")
}

type OneOfUpdateExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateExternalRepositoryApiResponseData() *OneOfUpdateExternalRepositoryApiResponseData {
	p := new(OneOfUpdateExternalRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateExternalRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateExternalRepositoryApiResponseData"))
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateExternalRepositoryApiResponseData")
}

type OneOfCreateExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateExternalRepositoryApiResponseData() *OneOfCreateExternalRepositoryApiResponseData {
	p := new(OneOfCreateExternalRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateExternalRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateExternalRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateExternalRepositoryApiResponseData"))
}

func (p *OneOfCreateExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateExternalRepositoryApiResponseData")
}

type OneOfDeleteProjectApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteProjectApiResponseData() *OneOfDeleteProjectApiResponseData {
	p := new(OneOfDeleteProjectApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteProjectApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteProjectApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteProjectApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteProjectApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteProjectApiResponseData"))
}

func (p *OneOfDeleteProjectApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteProjectApiResponseData")
}

type OneOfDeleteRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteRegisteredDomainApiResponseData() *OneOfDeleteRegisteredDomainApiResponseData {
	p := new(OneOfDeleteRegisteredDomainApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRegisteredDomainApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRegisteredDomainApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRegisteredDomainApiResponseData"))
}

func (p *OneOfDeleteRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRegisteredDomainApiResponseData")
}

type OneOfGetExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2002 *ExternalRepository    `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetExternalRepositoryApiResponseData() *OneOfGetExternalRepositoryApiResponseData {
	p := new(OneOfGetExternalRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetExternalRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetExternalRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ExternalRepository:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(ExternalRepository)
		}
		*p.oneOfType2002 = v.(ExternalRepository)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2002.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfGetExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2002 := new(ExternalRepository)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2002.ObjectType_ != nil && "multidomain.v4.config.ExternalRepository" == *vOneOfType2002.ObjectType_ {
							if nil == p.oneOfType2002 {
								p.oneOfType2002 = new(ExternalRepository)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType2002 := new(ExternalRepository)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if vOneOfType2002.ObjectType_ != nil && "multidomain.v4.config.ExternalRepository" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(ExternalRepository)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetExternalRepositoryApiResponseData"))
}

func (p *OneOfGetExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfGetExternalRepositoryApiResponseData")
}

type OneOfListRegisteredDomainsApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType0    []RegisteredDomain           `json:"-"`
	oneOfType400  *import4.ErrorResponse       `json:"-"`
	oneOfType401  []RegisteredDomainProjection `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListRegisteredDomainsApiResponseData() *OneOfListRegisteredDomainsApiResponseData {
	p := new(OneOfListRegisteredDomainsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListRegisteredDomainsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListRegisteredDomainsApiResponseData is nil"))
	}
	switch v.(type) {
	case []RegisteredDomain:
		p.oneOfType0 = v.([]RegisteredDomain)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.RegisteredDomain>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomain>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []RegisteredDomainProjection:
		p.oneOfType401 = v.([]RegisteredDomainProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.RegisteredDomainProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomainProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListRegisteredDomainsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<multidomain.v4.config.RegisteredDomain>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.RegisteredDomainProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRegisteredDomainsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.RegisteredDomain>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]RegisteredDomain)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "multidomain.v4.config.RegisteredDomain" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.RegisteredDomain>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomain>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.RegisteredDomainProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]RegisteredDomainProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.RegisteredDomainProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.RegisteredDomainProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomainProjection>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new([]RegisteredDomain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "multidomain.v4.config.RegisteredDomain" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.RegisteredDomain>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomain>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType401 := new([]RegisteredDomainProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.RegisteredDomainProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.RegisteredDomainProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.RegisteredDomainProjection>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRegisteredDomainsApiResponseData"))
}

func (p *OneOfListRegisteredDomainsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<multidomain.v4.config.RegisteredDomain>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.RegisteredDomainProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRegisteredDomainsApiResponseData")
}

type OneOfListResourceGroupsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType400  *import4.ErrorResponse    `json:"-"`
	oneOfType401  []ResourceGroupProjection `json:"-"`
	oneOfType2001 []ResourceGroup           `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListResourceGroupsApiResponseData() *OneOfListResourceGroupsApiResponseData {
	p := new(OneOfListResourceGroupsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListResourceGroupsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListResourceGroupsApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ResourceGroupProjection:
		p.oneOfType401 = v.([]ResourceGroupProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.ResourceGroupProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroupProjection>"
	case []ResourceGroup:
		p.oneOfType2001 = v.([]ResourceGroup)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.ResourceGroup>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroup>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListResourceGroupsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.ResourceGroupProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if "List<multidomain.v4.config.ResourceGroup>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListResourceGroupsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.ResourceGroupProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]ResourceGroupProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.ResourceGroupProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.ResourceGroupProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroupProjection>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.ResourceGroup>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]ResourceGroup)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.config.ResourceGroup" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.ResourceGroup>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroup>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType401 := new([]ResourceGroupProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.ResourceGroupProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.ResourceGroupProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroupProjection>"
			return nil
		}
	}
	vOneOfType2001 := new([]ResourceGroup)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.config.ResourceGroup" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.ResourceGroup>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.ResourceGroup>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListResourceGroupsApiResponseData"))
}

func (p *OneOfListResourceGroupsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.ResourceGroupProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if "List<multidomain.v4.config.ResourceGroup>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListResourceGroupsApiResponseData")
}

type OneOfUpdateProjectApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateProjectApiResponseData() *OneOfUpdateProjectApiResponseData {
	p := new(OneOfUpdateProjectApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateProjectApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateProjectApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateProjectApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateProjectApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateProjectApiResponseData"))
}

func (p *OneOfUpdateProjectApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateProjectApiResponseData")
}

type OneOfDomainDomainManager struct {
	Discriminator  *string                `json:"-"`
	ObjectType_    *string                `json:"-"`
	oneOfType12002 *import5.RemoteCluster `json:"-"`
	oneOfType12001 *import3.DomainManager `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
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
	case import5.RemoteCluster:
		if nil == p.oneOfType12002 {
			p.oneOfType12002 = new(import5.RemoteCluster)
		}
		*p.oneOfType12002 = v.(import5.RemoteCluster)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType12002.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType12002.ObjectType_
	case import3.DomainManager:
		if nil == p.oneOfType12001 {
			p.oneOfType12001 = new(import3.DomainManager)
		}
		*p.oneOfType12001 = v.(import3.DomainManager)
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
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12002
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType12001
	}
	return nil
}

func (p *OneOfDomainDomainManager) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType12002 := new(import5.RemoteCluster)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType12002)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType12002.ObjectType_ != nil && "prism.v4.management.RemoteCluster" == *vOneOfType12002.ObjectType_ {
							if nil == p.oneOfType12002 {
								p.oneOfType12002 = new(import5.RemoteCluster)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType12001 := new(import3.DomainManager)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType12001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType12001.ObjectType_ != nil && "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
							if nil == p.oneOfType12001 {
								p.oneOfType12001 = new(import3.DomainManager)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType12002 := new(import5.RemoteCluster)
	if err := json.Unmarshal(b, vOneOfType12002); err == nil {
		if vOneOfType12002.ObjectType_ != nil && "prism.v4.management.RemoteCluster" == *vOneOfType12002.ObjectType_ {
			if nil == p.oneOfType12002 {
				p.oneOfType12002 = new(import5.RemoteCluster)
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
	vOneOfType12001 := new(import3.DomainManager)
	if err := json.Unmarshal(b, vOneOfType12001); err == nil {
		if vOneOfType12001.ObjectType_ != nil && "prism.v4.config.DomainManager" == *vOneOfType12001.ObjectType_ {
			if nil == p.oneOfType12001 {
				p.oneOfType12001 = new(import3.DomainManager)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDomainDomainManager"))
}

func (p *OneOfDomainDomainManager) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType12002 != nil && *p.oneOfType12002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12002)
	}
	if p.oneOfType12001 != nil && *p.oneOfType12001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType12001)
	}
	return nil, errors.New("No value to marshal for OneOfDomainDomainManager")
}

type OneOfListProjectsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  []ProjectProjection    `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType1005 []Project              `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListProjectsApiResponseData() *OneOfListProjectsApiResponseData {
	p := new(OneOfListProjectsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListProjectsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListProjectsApiResponseData is nil"))
	}
	switch v.(type) {
	case []ProjectProjection:
		p.oneOfType401 = v.([]ProjectProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.ProjectProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.ProjectProjection>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Project:
		p.oneOfType1005 = v.([]Project)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.Project>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.Project>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListProjectsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<multidomain.v4.config.ProjectProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.Project>" == *p.Discriminator {
		return p.oneOfType1005
	}
	return nil
}

func (p *OneOfListProjectsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.ProjectProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]ProjectProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.ProjectProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.ProjectProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.ProjectProjection>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.Project>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new([]Project)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType1005 == nil || len(*vOneOfType1005) == 0 || ((*vOneOfType1005)[0].ObjectType_ != nil && "multidomain.v4.config.Project" == *((*vOneOfType1005)[0].ObjectType_)) {
							p.oneOfType1005 = *vOneOfType1005
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.Project>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.Project>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType401 := new([]ProjectProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.ProjectProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.ProjectProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.ProjectProjection>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType1005 := new([]Project)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		if len(*vOneOfType1005) == 0 || (vOneOfType1005 != nil && (*vOneOfType1005)[0].ObjectType_ != nil && "multidomain.v4.config.Project" == *((*vOneOfType1005)[0].ObjectType_)) {
			p.oneOfType1005 = *vOneOfType1005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.Project>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.Project>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListProjectsApiResponseData"))
}

func (p *OneOfListProjectsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<multidomain.v4.config.ProjectProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.Project>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	return nil, errors.New("No value to marshal for OneOfListProjectsApiResponseData")
}

type OneOfExternalRepositoryLocation struct {
	Discriminator *string        `json:"-"`
	ObjectType_   *string        `json:"-"`
	oneOfType2001 *NfsRepository `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfExternalRepositoryLocation() *OneOfExternalRepositoryLocation {
	p := new(OneOfExternalRepositoryLocation)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfExternalRepositoryLocation) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfExternalRepositoryLocation is nil"))
	}
	switch v.(type) {
	case NfsRepository:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(NfsRepository)
		}
		*p.oneOfType2001 = v.(NfsRepository)
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

func (p *OneOfExternalRepositoryLocation) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfExternalRepositoryLocation) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(NfsRepository)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.config.NfsRepository" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(NfsRepository)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(NfsRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.config.NfsRepository" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(NfsRepository)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfExternalRepositoryLocation"))
}

func (p *OneOfExternalRepositoryLocation) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfExternalRepositoryLocation")
}

type OneOfDeleteResourceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteResourceGroupApiResponseData() *OneOfDeleteResourceGroupApiResponseData {
	p := new(OneOfDeleteResourceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteResourceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteResourceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfDeleteResourceGroupApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteResourceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteResourceGroupApiResponseData"))
}

func (p *OneOfDeleteResourceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteResourceGroupApiResponseData")
}

type OneOfUpdateResourceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateResourceGroupApiResponseData() *OneOfUpdateResourceGroupApiResponseData {
	p := new(OneOfUpdateResourceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateResourceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateResourceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfUpdateResourceGroupApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateResourceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateResourceGroupApiResponseData"))
}

func (p *OneOfUpdateResourceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateResourceGroupApiResponseData")
}

type OneOfRefreshApiCredentialsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfRefreshApiCredentialsApiResponseData() *OneOfRefreshApiCredentialsApiResponseData {
	p := new(OneOfRefreshApiCredentialsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRefreshApiCredentialsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRefreshApiCredentialsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfRefreshApiCredentialsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRefreshApiCredentialsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRefreshApiCredentialsApiResponseData"))
}

func (p *OneOfRefreshApiCredentialsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRefreshApiCredentialsApiResponseData")
}

type OneOfGetProjectApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType401  *ProjectProjection     `json:"-"`
	oneOfType1005 *Project               `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetProjectApiResponseData() *OneOfGetProjectApiResponseData {
	p := new(OneOfGetProjectApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetProjectApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetProjectApiResponseData is nil"))
	}
	switch v.(type) {
	case ProjectProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(ProjectProjection)
		}
		*p.oneOfType401 = v.(ProjectProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	case Project:
		if nil == p.oneOfType1005 {
			p.oneOfType1005 = new(Project)
		}
		*p.oneOfType1005 = v.(Project)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1005.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1005.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfGetProjectApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1005
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetProjectApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new(ProjectProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.ProjectProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(ProjectProjection)
							}
							*p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType401.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType1005 := new(Project)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType1005)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType1005.ObjectType_ != nil && "multidomain.v4.config.Project" == *vOneOfType1005.ObjectType_ {
							if nil == p.oneOfType1005 {
								p.oneOfType1005 = new(Project)
							}
							*p.oneOfType1005 = *vOneOfType1005
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType1005.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType1005.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType401 := new(ProjectProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.ProjectProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(ProjectProjection)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	vOneOfType1005 := new(Project)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		if vOneOfType1005.ObjectType_ != nil && "multidomain.v4.config.Project" == *vOneOfType1005.ObjectType_ {
			if nil == p.oneOfType1005 {
				p.oneOfType1005 = new(Project)
			}
			*p.oneOfType1005 = *vOneOfType1005
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1005.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1005.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetProjectApiResponseData"))
}

func (p *OneOfGetProjectApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	if p.oneOfType1005 != nil && *p.oneOfType1005.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetProjectApiResponseData")
}

type OneOfListLocationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Location             `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType401  []LocationProjection   `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListLocationsApiResponseData() *OneOfListLocationsApiResponseData {
	p := new(OneOfListLocationsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLocationsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLocationsApiResponseData is nil"))
	}
	switch v.(type) {
	case []Location:
		p.oneOfType0 = v.([]Location)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.Location>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.Location>"
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []LocationProjection:
		p.oneOfType401 = v.([]LocationProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.config.LocationProjection>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.config.LocationProjection>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListLocationsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if "List<multidomain.v4.config.Location>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.LocationProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListLocationsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.Location>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new([]Location)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType0 == nil || len(*vOneOfType0) == 0 || ((*vOneOfType0)[0].ObjectType_ != nil && "multidomain.v4.config.Location" == *((*vOneOfType0)[0].ObjectType_)) {
							p.oneOfType0 = *vOneOfType0
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.Location>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.Location>"
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.config.LocationProjection>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new([]LocationProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType401 == nil || len(*vOneOfType401) == 0 || ((*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.LocationProjection" == *((*vOneOfType401)[0].ObjectType_)) {
							p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.config.LocationProjection>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.config.LocationProjection>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType0 := new([]Location)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || (vOneOfType0 != nil && (*vOneOfType0)[0].ObjectType_ != nil && "multidomain.v4.config.Location" == *((*vOneOfType0)[0].ObjectType_)) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.Location>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.Location>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType401 := new([]LocationProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || (vOneOfType401 != nil && (*vOneOfType401)[0].ObjectType_ != nil && "multidomain.v4.config.LocationProjection" == *((*vOneOfType401)[0].ObjectType_)) {
			p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.config.LocationProjection>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.config.LocationProjection>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLocationsApiResponseData"))
}

func (p *OneOfListLocationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if "List<multidomain.v4.config.Location>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.LocationProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListLocationsApiResponseData")
}

type OneOfCreateProjectApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateProjectApiResponseData() *OneOfCreateProjectApiResponseData {
	p := new(OneOfCreateProjectApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateProjectApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateProjectApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateProjectApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateProjectApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateProjectApiResponseData"))
}

func (p *OneOfCreateProjectApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateProjectApiResponseData")
}

type OneOfGetRegisteredDomainApiResponseData struct {
	Discriminator *string                     `json:"-"`
	ObjectType_   *string                     `json:"-"`
	oneOfType400  *import4.ErrorResponse      `json:"-"`
	oneOfType0    *RegisteredDomain           `json:"-"`
	oneOfType401  *RegisteredDomainProjection `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetRegisteredDomainApiResponseData() *OneOfGetRegisteredDomainApiResponseData {
	p := new(OneOfGetRegisteredDomainApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRegisteredDomainApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRegisteredDomainApiResponseData is nil"))
	}
	switch v.(type) {
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case RegisteredDomain:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RegisteredDomain)
		}
		*p.oneOfType0 = v.(RegisteredDomain)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case RegisteredDomainProjection:
		if nil == p.oneOfType401 {
			p.oneOfType401 = new(RegisteredDomainProjection)
		}
		*p.oneOfType401 = v.(RegisteredDomainProjection)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType401.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType401
	}
	return nil
}

func (p *OneOfGetRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType0 := new(RegisteredDomain)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType0)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType0.ObjectType_ != nil && "multidomain.v4.config.RegisteredDomain" == *vOneOfType0.ObjectType_ {
							if nil == p.oneOfType0 {
								p.oneOfType0 = new(RegisteredDomain)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType401 := new(RegisteredDomainProjection)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType401)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.RegisteredDomainProjection" == *vOneOfType401.ObjectType_ {
							if nil == p.oneOfType401 {
								p.oneOfType401 = new(RegisteredDomainProjection)
							}
							*p.oneOfType401 = *vOneOfType401
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType401.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType401.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	vOneOfType0 := new(RegisteredDomain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if vOneOfType0.ObjectType_ != nil && "multidomain.v4.config.RegisteredDomain" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RegisteredDomain)
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
	vOneOfType401 := new(RegisteredDomainProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if vOneOfType401.ObjectType_ != nil && "multidomain.v4.config.RegisteredDomainProjection" == *vOneOfType401.ObjectType_ {
			if nil == p.oneOfType401 {
				p.oneOfType401 = new(RegisteredDomainProjection)
			}
			*p.oneOfType401 = *vOneOfType401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType401.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRegisteredDomainApiResponseData"))
}

func (p *OneOfGetRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType401 != nil && *p.oneOfType401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfGetRegisteredDomainApiResponseData")
}

type OneOfCreateResourceGroupApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateResourceGroupApiResponseData() *OneOfCreateResourceGroupApiResponseData {
	p := new(OneOfCreateResourceGroupApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateResourceGroupApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateResourceGroupApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import3.TaskReference)
		}
		*p.oneOfType2001 = v.(import3.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import4.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import4.ErrorResponse)
		}
		*p.oneOfType400 = v.(import4.ErrorResponse)
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

func (p *OneOfCreateResourceGroupApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateResourceGroupApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import3.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import3.TaskReference)
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
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import4.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import4.ErrorResponse)
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
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import3.TaskReference)
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
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import4.ErrorResponse)
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
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateResourceGroupApiResponseData"))
}

func (p *OneOfCreateResourceGroupApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateResourceGroupApiResponseData")
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
