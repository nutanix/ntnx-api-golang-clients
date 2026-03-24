/*
 * Generated file models/multidomain/v4/config/config_model.go.
 *
 * Product version: 4.3.1
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
	import6 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
	import5 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/common"
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/resources"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/prism/v4/config"
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
REST response for all response codes in API path /multidomain/v4.3/config/external-repositories Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/registered-domains Post operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/external-repositories/{extId} Delete operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/registered-domains/{extId} Delete operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	DomainProjection *import4.DomainProjection `json:"domainProjection,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	Provider *import5.TunnelProvider `json:"provider,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.3/config/external-repositories/{extId} Get operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/registered-domains/{extId} Get operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/external-repositories Get operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/locations Get operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/registered-domains Get operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	Address *import6.IPAddressOrFQDN `json:"address"`
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
		Address *import6.IPAddressOrFQDN `json:"address,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

	RegistrationState *import5.RegistrationState `json:"registrationState,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	RegistrationState *import5.RegistrationState `json:"registrationState"`
}

func (p *RegisteredDomainChangeStateSpec) MarshalJSON() ([]byte, error) {
	type RegisteredDomainChangeStateSpecProxy RegisteredDomainChangeStateSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RegisteredDomainChangeStateSpecProxy
		RegistrationState *import5.RegistrationState `json:"registrationState,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	RegistrationState *import5.RegistrationState `json:"registrationState,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

	Credentials *import5.RegistrationCredentials `json:"credentials,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.3/config/external-repositories/{extId} Put operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
REST response for all response codes in API path /multidomain/v4.3/config/registered-domains/{extId} Put operation
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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

type OneOfDeleteExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteExternalRepositoryApiResponseData"))
}

func (p *OneOfDeleteExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteExternalRepositoryApiResponseData")
}

type OneOfGetRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *RegisteredDomain      `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(RegisteredDomain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "multidomain.v4.config.RegisteredDomain" == *vOneOfType0.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRegisteredDomainApiResponseData"))
}

func (p *OneOfGetRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetRegisteredDomainApiResponseData")
}

type OneOfUpdateExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateExternalRepositoryApiResponseData"))
}

func (p *OneOfUpdateExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateExternalRepositoryApiResponseData")
}

type OneOfUpdateRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
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
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfUpdateRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfUpdateRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRegisteredDomainApiResponseData"))
}

func (p *OneOfUpdateRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRegisteredDomainApiResponseData")
}

type OneOfExternalRepositoryLocation struct {
	Discriminator *string        `json:"-"`
	ObjectType_   *string        `json:"-"`
	oneOfType2001 *NfsRepository `json:"-"`
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
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfExternalRepositoryLocation) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(NfsRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "multidomain.v4.config.NfsRepository" == *vOneOfType2001.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfExternalRepositoryLocation"))
}

func (p *OneOfExternalRepositoryLocation) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfExternalRepositoryLocation")
}

type OneOfCreateExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateExternalRepositoryApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCreateExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateExternalRepositoryApiResponseData"))
}

func (p *OneOfCreateExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCreateExternalRepositoryApiResponseData")
}

type OneOfListLocationsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []Location             `json:"-"`
	oneOfType401  []LocationProjection   `json:"-"`
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.Location>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<multidomain.v4.config.LocationProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListLocationsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]Location)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "multidomain.v4.config.Location" == *((*vOneOfType0)[0].ObjectType_) {
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
	vOneOfType401 := new([]LocationProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "multidomain.v4.config.LocationProjection" == *((*vOneOfType401)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLocationsApiResponseData"))
}

func (p *OneOfListLocationsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.Location>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<multidomain.v4.config.LocationProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListLocationsApiResponseData")
}

type OneOfDeleteRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
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
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfDeleteRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRegisteredDomainApiResponseData"))
}

func (p *OneOfDeleteRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRegisteredDomainApiResponseData")
}

type OneOfListRegisteredDomainsApiResponseData struct {
	Discriminator *string                      `json:"-"`
	ObjectType_   *string                      `json:"-"`
	oneOfType400  *import3.ErrorResponse       `json:"-"`
	oneOfType0    []RegisteredDomain           `json:"-"`
	oneOfType401  []RegisteredDomainProjection `json:"-"`
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.RegisteredDomain>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "List<multidomain.v4.config.RegisteredDomainProjection>" == *p.Discriminator {
		return p.oneOfType401
	}
	return nil
}

func (p *OneOfListRegisteredDomainsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new([]RegisteredDomain)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "multidomain.v4.config.RegisteredDomain" == *((*vOneOfType0)[0].ObjectType_) {
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
	vOneOfType401 := new([]RegisteredDomainProjection)
	if err := json.Unmarshal(b, vOneOfType401); err == nil {
		if len(*vOneOfType401) == 0 || "multidomain.v4.config.RegisteredDomainProjection" == *((*vOneOfType401)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListRegisteredDomainsApiResponseData"))
}

func (p *OneOfListRegisteredDomainsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.RegisteredDomain>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "List<multidomain.v4.config.RegisteredDomainProjection>" == *p.Discriminator {
		return json.Marshal(p.oneOfType401)
	}
	return nil, errors.New("No value to marshal for OneOfListRegisteredDomainsApiResponseData")
}

type OneOfGroupViewAttributeName struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType1    *DomainAttributeName  `json:"-"`
	oneOfType0    *ClusterAttributeName `json:"-"`
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
	if "multidomain.v4.config.DomainAttributeName" == *p.Discriminator {
		return *p.oneOfType1
	}
	if "multidomain.v4.config.ClusterAttributeName" == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGroupViewAttributeName) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGroupViewAttributeName"))
}

func (p *OneOfGroupViewAttributeName) MarshalJSON() ([]byte, error) {
	if "multidomain.v4.config.DomainAttributeName" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if "multidomain.v4.config.ClusterAttributeName" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGroupViewAttributeName")
}

type OneOfGetExternalRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2002 *ExternalRepository    `json:"-"`
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfGetExternalRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2002 := new(ExternalRepository)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "multidomain.v4.config.ExternalRepository" == *vOneOfType2002.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetExternalRepositoryApiResponseData"))
}

func (p *OneOfGetExternalRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfGetExternalRepositoryApiResponseData")
}

type OneOfListExternalRepositoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []ExternalRepository   `json:"-"`
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
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.config.ExternalRepository>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListExternalRepositoriesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]ExternalRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "multidomain.v4.config.ExternalRepository" == *((*vOneOfType2001)[0].ObjectType_) {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListExternalRepositoriesApiResponseData"))
}

func (p *OneOfListExternalRepositoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.config.ExternalRepository>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListExternalRepositoriesApiResponseData")
}

type OneOfCreateRegisteredDomainApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
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
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfCreateRegisteredDomainApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCreateRegisteredDomainApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateRegisteredDomainApiResponseData"))
}

func (p *OneOfCreateRegisteredDomainApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCreateRegisteredDomainApiResponseData")
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
