/*
 * Generated file models/lifecycle/v4/resources/resources_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Lifecycle resources such as reccomendations and notifications.
*/
package resources

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import5 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/prism/v4/config"
	"time"
)

/*
LCM framework version present in the LCM URL.
*/
type AvailableVersion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Available version UUID.
	*/
	AvailableVersionUuid *string `json:"availableVersionUuid,omitempty"`
	/*
	  Component information for the payload based entity.
	*/
	ChildEntities []string `json:"childEntities,omitempty"`
	/*
	  Update custom messages other than release notes to the user about this available version.
	*/
	CustomMessage *string `json:"customMessage,omitempty"`
	/*
	  List of dependencies for the available version.
	*/
	Dependencies []DependentEntity `json:"dependencies,omitempty"`
	/*
	  Reason for disabling the available version.
	*/
	DisablementReason *string `json:"disablementReason,omitempty"`
	/*
	  UUID of the group that this LCM entity is part of.
	*/
	GroupUuid *string `json:"groupUuid,omitempty"`
	/*
	  Indicates if the available update is enabled.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Order of this available version (1 being the lowest and 6 being the highest) when multiple versions are present with different status.
	*/
	Order *int64 `json:"order,omitempty"`
	/*
	  Release date for the entities that need this information.
	*/
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
	/*
	  Release notes corresponding to the update.
	*/
	ReleaseNotes *string `json:"releaseNotes,omitempty"`

	Status *import1.AvailableVersionStatus `json:"status,omitempty"`
	/*
	  LCM framework version present in the LCM URL.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *AvailableVersion) MarshalJSON() ([]byte, error) {
	type AvailableVersionProxy AvailableVersion

	// Step 1: Marshal known fields via proxy to accommodate fields with date format
	baseStruct := struct {
		*AvailableVersionProxy
		ReleaseDate string `json:"releaseDate,omitempty"`
	}{
		AvailableVersionProxy: (*AvailableVersionProxy)(p),
		ReleaseDate: func() string {
			if p.ReleaseDate != nil {
				return p.ReleaseDate.Format("2006-01-02")
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

func (p *AvailableVersion) UnmarshalJSON(b []byte) error {

	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields and custom parsing
	type CustomAvailableVersion struct {
		ObjectType_          *string                         `json:"$objectType,omitempty"`
		Reserved_            map[string]interface{}          `json:"$reserved,omitempty"`
		UnknownFields_       map[string]interface{}          `json:"$unknownFields,omitempty"`
		AvailableVersionUuid *string                         `json:"availableVersionUuid,omitempty"`
		ChildEntities        []string                        `json:"childEntities,omitempty"`
		CustomMessage        *string                         `json:"customMessage,omitempty"`
		Dependencies         []DependentEntity               `json:"dependencies,omitempty"`
		DisablementReason    *string                         `json:"disablementReason,omitempty"`
		GroupUuid            *string                         `json:"groupUuid,omitempty"`
		IsEnabled            *bool                           `json:"isEnabled,omitempty"`
		Order                *int64                          `json:"order,omitempty"`
		ReleaseDate          string                          `json:"releaseDate,omitempty"`
		ReleaseNotes         *string                         `json:"releaseNotes,omitempty"`
		Status               *import1.AvailableVersionStatus `json:"status,omitempty"`
		Version              *string                         `json:"version,omitempty"`
	}

	var knownFields CustomAvailableVersion
	if err := json.Unmarshal(b, &knownFields); err != nil {
		return err
	}

	// Step 3: Assign known fields
	// Handle custom date parsing
	p.ObjectType_ = knownFields.ObjectType_
	// Handle custom date parsing
	p.Reserved_ = knownFields.Reserved_
	// Handle custom date parsing
	p.UnknownFields_ = knownFields.UnknownFields_
	// Handle custom date parsing
	p.AvailableVersionUuid = knownFields.AvailableVersionUuid
	// Handle custom date parsing
	p.ChildEntities = knownFields.ChildEntities
	// Handle custom date parsing
	p.CustomMessage = knownFields.CustomMessage
	// Handle custom date parsing
	p.Dependencies = knownFields.Dependencies
	// Handle custom date parsing
	p.DisablementReason = knownFields.DisablementReason
	// Handle custom date parsing
	p.GroupUuid = knownFields.GroupUuid
	// Handle custom date parsing
	p.IsEnabled = knownFields.IsEnabled
	// Handle custom date parsing
	p.Order = knownFields.Order
	// Handle custom date parsing
	// Custom date parsing logic for Date field
	if knownFields.ReleaseDate != "" {
		parsedReleaseDate, err := time.Parse("2006-01-02", knownFields.ReleaseDate)
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to unmarshal field ReleaseDate in struct AvailableVersion: %s", err))
		}
		p.ReleaseDate = &parsedReleaseDate
	}
	// Handle custom date parsing
	p.ReleaseNotes = knownFields.ReleaseNotes
	// Handle custom date parsing
	p.Status = knownFields.Status
	// Handle custom date parsing
	p.Version = knownFields.Version

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableVersionUuid")
	delete(allFields, "childEntities")
	delete(allFields, "customMessage")
	delete(allFields, "dependencies")
	delete(allFields, "disablementReason")
	delete(allFields, "groupUuid")
	delete(allFields, "isEnabled")
	delete(allFields, "order")
	delete(allFields, "releaseDate")
	delete(allFields, "releaseNotes")
	delete(allFields, "status")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAvailableVersion() *AvailableVersion {
	p := new(AvailableVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.AvailableVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsEnabled = new(bool)
	*p.IsEnabled = true

	return p
}

/*
Details of the LCM bundle
*/
type Bundle struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ChecksumItemDiscriminator_ *string `json:"$checksumItemDiscriminator,omitempty"`
	/*
	  SHA256 sum.
	*/
	Checksum *OneOfBundleChecksum `json:"checksum,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of LCM images
	*/
	Images []Image `json:"images,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the LCM bundle
	*/
	Name *string `json:"name"`
	/*
	  Size of the LCM bundle
	*/
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *BundleType `json:"type,omitempty"`

	Vendor *BundleVendor `json:"vendor"`
}

func (p *Bundle) MarshalJSON() ([]byte, error) {
	type BundleProxy Bundle

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BundleProxy
		Name   *string       `json:"name,omitempty"`
		Vendor *BundleVendor `json:"vendor,omitempty"`
	}{
		BundleProxy: (*BundleProxy)(p),
		Name:        p.Name,
		Vendor:      p.Vendor,
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

func (p *Bundle) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Bundle
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Bundle(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$checksumItemDiscriminator")
	delete(allFields, "checksum")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "images")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "sizeBytes")
	delete(allFields, "tenantId")
	delete(allFields, "type")
	delete(allFields, "vendor")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBundle() *Bundle {
	p := new(Bundle)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Bundle"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Bundle) GetChecksum() interface{} {
	if nil == p.Checksum {
		return nil
	}
	return p.Checksum.GetValue()
}

func (p *Bundle) SetChecksum(v interface{}) error {
	if nil == p.Checksum {
		p.Checksum = NewOneOfBundleChecksum()
	}
	e := p.Checksum.SetValue(v)
	if nil == e {
		if nil == p.ChecksumItemDiscriminator_ {
			p.ChecksumItemDiscriminator_ = new(string)
		}
		*p.ChecksumItemDiscriminator_ = *p.Checksum.Discriminator
	}
	return e
}

/*
Type of LCM bundle being uploaded
*/
type BundleType int

const (
	BUNDLETYPE_UNKNOWN      BundleType = 0
	BUNDLETYPE_REDACTED     BundleType = 1
	BUNDLETYPE_SOFTWARE     BundleType = 2
	BUNDLETYPE_FIRMWARE     BundleType = 3
	BUNDLETYPE_PRODUCT_META BundleType = 4
	BUNDLETYPE_FRAMEWORK    BundleType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BundleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
		"PRODUCT_META",
		"FRAMEWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BundleType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
		"PRODUCT_META",
		"FRAMEWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BundleType) index(name string) BundleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"FIRMWARE",
		"PRODUCT_META",
		"FRAMEWORK",
	}
	for idx := range names {
		if names[idx] == name {
			return BundleType(idx)
		}
	}
	return BUNDLETYPE_UNKNOWN
}

func (e *BundleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BundleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BundleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BundleType) Ref() *BundleType {
	return &e
}

/*
Owner or vendor of the Bundle
*/
type BundleVendor int

const (
	BUNDLEVENDOR_UNKNOWN     BundleVendor = 0
	BUNDLEVENDOR_REDACTED    BundleVendor = 1
	BUNDLEVENDOR_NUTANIX     BundleVendor = 2
	BUNDLEVENDOR_THIRD_PARTY BundleVendor = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *BundleVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"THIRD_PARTY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e BundleVendor) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"THIRD_PARTY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *BundleVendor) index(name string) BundleVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NUTANIX",
		"THIRD_PARTY",
	}
	for idx := range names {
		if names[idx] == name {
			return BundleVendor(idx)
		}
	}
	return BUNDLEVENDOR_UNKNOWN
}

func (e *BundleVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BundleVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BundleVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BundleVendor) Ref() *BundleVendor {
	return &e
}

/*
Details of a cluster capability.
*/
type Capability int

const (
	CAPABILITY_UNKNOWN       Capability = 0
	CAPABILITY_REDACTED      Capability = 1
	CAPABILITY_MCL_INVENTORY Capability = 2
	CAPABILITY_MCL_UPGRADE   Capability = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Capability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MCL_INVENTORY",
		"MCL_UPGRADE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Capability) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MCL_INVENTORY",
		"MCL_UPGRADE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Capability) index(name string) Capability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MCL_INVENTORY",
		"MCL_UPGRADE",
	}
	for idx := range names {
		if names[idx] == name {
			return Capability(idx)
		}
	}
	return CAPABILITY_UNKNOWN
}

func (e *Capability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Capability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Capability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Capability) Ref() *Capability {
	return &e
}

/*
LCM configuration on the cluster.
*/
type Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The scheduled time in "%H:%M" 24-hour format of the next inventory execution. Used when auto_inventory_enabled is set to True. The default schedule time is 03:00(AM).
	*/
	AutoInventorySchedule *string `json:"autoInventorySchedule,omitempty"`

	ConnectivityType *ConnectivityType `json:"connectivityType,omitempty"`
	/*
	  List of entities for which One-Click upgrades are not available.
	*/
	DeprecatedSoftwareEntities []string `json:"deprecatedSoftwareEntities,omitempty"`
	/*
	  User friendly display version of LCM installed on the cluster.
	*/
	DisplayVersion *string `json:"displayVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Indicates if LCM is enabled to auto-upgrade products. The default value is False.
	*/
	HasModuleAutoUpgradeEnabled *bool `json:"hasModuleAutoUpgradeEnabled,omitempty"`
	/*
	  Indicates if the auto inventory operation is enabled. The default value is set to False.
	*/
	IsAutoInventoryEnabled *bool `json:"isAutoInventoryEnabled,omitempty"`
	/*
	  Indicates if the bundle is uploaded or not.
	*/
	IsFrameworkBundleUploaded *bool `json:"isFrameworkBundleUploaded,omitempty"`
	/*
	  Indicates if the LCM URL has HTTPS enabled. The default value is True.
	*/
	IsHttpsEnabled *bool `json:"isHttpsEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of entities for which One-Click upgrades are supported.
	*/
	SupportedSoftwareEntities []string `json:"supportedSoftwareEntities,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  URL of the LCM repository.
	*/
	Url *string `json:"url,omitempty"`
	/*
	  LCM version installed on the cluster.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *Config) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Config

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

func (p *Config) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Config
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Config(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "autoInventorySchedule")
	delete(allFields, "connectivityType")
	delete(allFields, "deprecatedSoftwareEntities")
	delete(allFields, "displayVersion")
	delete(allFields, "extId")
	delete(allFields, "hasModuleAutoUpgradeEnabled")
	delete(allFields, "isAutoInventoryEnabled")
	delete(allFields, "isFrameworkBundleUploaded")
	delete(allFields, "isHttpsEnabled")
	delete(allFields, "links")
	delete(allFields, "supportedSoftwareEntities")
	delete(allFields, "tenantId")
	delete(allFields, "url")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConfig() *Config {
	p := new(Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.HasModuleAutoUpgradeEnabled = new(bool)
	*p.HasModuleAutoUpgradeEnabled = false
	p.IsAutoInventoryEnabled = new(bool)
	*p.IsAutoInventoryEnabled = false
	p.IsFrameworkBundleUploaded = new(bool)
	*p.IsFrameworkBundleUploaded = false
	p.IsHttpsEnabled = new(bool)
	*p.IsHttpsEnabled = false

	return p
}

/*
This field indicates whether LCM framework on the cluster is running in connected-site mode or darksite mode.
*/
type ConnectivityType int

const (
	CONNECTIVITYTYPE_UNKNOWN                ConnectivityType = 0
	CONNECTIVITYTYPE_REDACTED               ConnectivityType = 1
	CONNECTIVITYTYPE_CONNECTED_SITE         ConnectivityType = 2
	CONNECTIVITYTYPE_DARKSITE_WEB_SERVER    ConnectivityType = 3
	CONNECTIVITYTYPE_DARKSITE_DIRECT_UPLOAD ConnectivityType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConnectivityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConnectivityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConnectivityType) index(name string) ConnectivityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONNECTED_SITE",
		"DARKSITE_WEB_SERVER",
		"DARKSITE_DIRECT_UPLOAD",
	}
	for idx := range names {
		if names[idx] == name {
			return ConnectivityType(idx)
		}
	}
	return CONNECTIVITYTYPE_UNKNOWN
}

func (e *ConnectivityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConnectivityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConnectivityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConnectivityType) Ref() *ConnectivityType {
	return &e
}

/*
REST response for all response codes in API path /lifecycle/v4.1/resources/bundles Post operation
*/
type CreateBundleApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateBundleApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateBundleApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateBundleApiResponse

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

func (p *CreateBundleApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateBundleApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CreateBundleApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCreateBundleApiResponse() *CreateBundleApiResponse {
	p := new(CreateBundleApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.CreateBundleApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateBundleApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateBundleApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateBundleApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/bundles/{extId} Delete operation
*/
type DeleteBundleByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteBundleByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteBundleByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteBundleByIdApiResponse

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

func (p *DeleteBundleByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteBundleByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeleteBundleByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDeleteBundleByIdApiResponse() *DeleteBundleByIdApiResponse {
	p := new(DeleteBundleByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.DeleteBundleByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteBundleByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteBundleByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteBundleByIdApiResponseData()
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
Dependency of an LCM entity available version.
*/
type DependentEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Information of the dependent entity versions for this available entity.
	*/
	DependentVersions []import5.KVPair `json:"dependentVersions,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DependentEntity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DependentEntity

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

func (p *DependentEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DependentEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DependentEntity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dependentVersions")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDependentEntity() *DependentEntity {
	p := new(DependentEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.DependentEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of deployable versions based on entity types, versions and its dependencies.
*/
type DeployableVersion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Entity target version.
	*/
	TargetVersion *string `json:"targetVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UpdateDependencies []import1.EntityUpdateSpec `json:"updateDependencies,omitempty"`
}

func (p *DeployableVersion) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeployableVersion

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

func (p *DeployableVersion) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeployableVersion
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DeployableVersion(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "links")
	delete(allFields, "targetVersion")
	delete(allFields, "tenantId")
	delete(allFields, "updateDependencies")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDeployableVersion() *DeployableVersion {
	p := new(DeployableVersion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.DeployableVersion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of an LCM entity.
*/
type Entity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of available versions for an LCM entity to update.
	*/
	AvailableVersions []AvailableVersion `json:"availableVersions,omitempty"`
	/*
	  Component information for the payload based entity.
	*/
	ChildEntities []string `json:"childEntities,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Unique identifier of an LCM entity e.g. "HDD serial number".
	*/
	DeviceId *string `json:"deviceId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  Description of an LCM entity.
	*/
	EntityDescription *string `json:"entityDescription,omitempty"`
	/*
	  Detailed information for the LCM entity. For example, firmware entities contain additional information about NIC and so on.
	*/
	EntityDetails []import5.KVPair `json:"entityDetails,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  UUID of the group that this LCM entity is part of.
	*/
	GroupUuid *string `json:"groupUuid,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  Hardware vendor information.
	*/
	HardwareVendor *string `json:"hardwareVendor,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was last updated.
	*/
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A list of sub-entities applicable to the entity.
	*/
	SubEntities []import1.EntityBaseModel `json:"subEntities,omitempty"`
	/*
	  The requested update version of an LCM entity.
	*/
	TargetVersion *string `json:"targetVersion,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Entity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Entity

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

func (p *Entity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Entity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Entity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableVersions")
	delete(allFields, "childEntities")
	delete(allFields, "clusterExtId")
	delete(allFields, "deviceId")
	delete(allFields, "entityClass")
	delete(allFields, "entityDescription")
	delete(allFields, "entityDetails")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "groupUuid")
	delete(allFields, "hardwareFamily")
	delete(allFields, "hardwareVendor")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "links")
	delete(allFields, "locationInfo")
	delete(allFields, "subEntities")
	delete(allFields, "targetVersion")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Entity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM framework version information.
*/
type FrameworkVersionInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM framework version present in the LCM URL.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/*
	  Current LCM Version.
	*/
	CurrentVersion *string `json:"currentVersion,omitempty"`
	/*
	  Boolean that indicates if LCM framework update is needed.
	*/
	IsUpdateNeeded *bool `json:"isUpdateNeeded,omitempty"`
}

func (p *FrameworkVersionInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FrameworkVersionInfo

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

func (p *FrameworkVersionInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FrameworkVersionInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = FrameworkVersionInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableVersion")
	delete(allFields, "currentVersion")
	delete(allFields, "isUpdateNeeded")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewFrameworkVersionInfo() *FrameworkVersionInfo {
	p := new(FrameworkVersionInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.FrameworkVersionInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsUpdateNeeded = new(bool)
	*p.IsUpdateNeeded = false

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.1/resources/bundles/{extId} Get operation
*/
type GetBundleByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetBundleByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetBundleByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetBundleByIdApiResponse

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

func (p *GetBundleByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetBundleByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetBundleByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetBundleByIdApiResponse() *GetBundleByIdApiResponse {
	p := new(GetBundleByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetBundleByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetBundleByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetBundleByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetBundleByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/config Get operation
*/
type GetConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetConfigApiResponse

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

func (p *GetConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetConfigApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetConfigApiResponse() *GetConfigApiResponse {
	p := new(GetConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetConfigApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/entities/{extId} Get operation
*/
type GetEntityByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetEntityByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetEntityByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetEntityByIdApiResponse

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

func (p *GetEntityByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetEntityByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetEntityByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetEntityByIdApiResponse() *GetEntityByIdApiResponse {
	p := new(GetEntityByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetEntityByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetEntityByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetEntityByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetEntityByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/lcm-summaries/{extId} Get operation
*/
type GetLcmSummaryByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLcmSummaryByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetLcmSummaryByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetLcmSummaryByIdApiResponse

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

func (p *GetLcmSummaryByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetLcmSummaryByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetLcmSummaryByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetLcmSummaryByIdApiResponse() *GetLcmSummaryByIdApiResponse {
	p := new(GetLcmSummaryByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetLcmSummaryByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLcmSummaryByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLcmSummaryByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLcmSummaryByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/notifications/{extId} Get operation
*/
type GetNotificationsByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNotificationsByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetNotificationsByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetNotificationsByIdApiResponse

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

func (p *GetNotificationsByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetNotificationsByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetNotificationsByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetNotificationsByIdApiResponse() *GetNotificationsByIdApiResponse {
	p := new(GetNotificationsByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetNotificationsByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNotificationsByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNotificationsByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNotificationsByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/recommendations/{extId} Get operation
*/
type GetRecommendationByIdApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRecommendationByIdApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetRecommendationByIdApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetRecommendationByIdApiResponse

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

func (p *GetRecommendationByIdApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetRecommendationByIdApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetRecommendationByIdApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetRecommendationByIdApiResponse() *GetRecommendationByIdApiResponse {
	p := new(GetRecommendationByIdApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetRecommendationByIdApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRecommendationByIdApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRecommendationByIdApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRecommendationByIdApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/status Get operation
*/
type GetStatusApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetStatusApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetStatusApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetStatusApiResponse

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

func (p *GetStatusApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetStatusApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetStatusApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetStatusApiResponse() *GetStatusApiResponse {
	p := new(GetStatusApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.GetStatusApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetStatusApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetStatusApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetStatusApiResponseData()
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
LCM image (List of LCM image files).
*/
type Image struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  List of files in the image.
	*/
	Files []ImageFile `json:"files,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  Denotes if the thirdparty version is qualified.
	*/
	IsQualified *bool `json:"isQualified,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Release notes for the LCM image.
	*/
	ReleaseNotes *string `json:"releaseNotes,omitempty"`
	/*
	  Version specification for image metadata JSON.
	*/
	SpecVersion *string `json:"specVersion"`

	Status *import1.AvailableVersionStatus `json:"status,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Image) MarshalJSON() ([]byte, error) {
	type ImageProxy Image

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageProxy
		SpecVersion *string `json:"specVersion,omitempty"`
	}{
		ImageProxy:  (*ImageProxy)(p),
		SpecVersion: p.SpecVersion,
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

func (p *Image) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Image
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Image(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "files")
	delete(allFields, "hardwareFamily")
	delete(allFields, "isQualified")
	delete(allFields, "links")
	delete(allFields, "releaseNotes")
	delete(allFields, "specVersion")
	delete(allFields, "status")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewImage() *Image {
	p := new(Image)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Image"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsQualified = new(bool)
	*p.IsQualified = false

	return p
}

/*
Description of LCM image file.
*/
type ImageFile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM image checksum.
	*/
	Checksum *string `json:"checksum"`

	ChecksumType *import1.CheckSumType `json:"checksumType"`
	/*
	  Image file global catalog item UUID.
	*/
	FileLocationId *string `json:"fileLocationId,omitempty"`
	/*
	  File path for the LCM image.
	*/
	FilePath *string `json:"filePath,omitempty"`
	/*
	  LCM image file name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  LCM image file size.
	*/
	SizeBytes *int64 `json:"sizeBytes"`
}

func (p *ImageFile) MarshalJSON() ([]byte, error) {
	type ImageFileProxy ImageFile

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ImageFileProxy
		Checksum     *string               `json:"checksum,omitempty"`
		ChecksumType *import1.CheckSumType `json:"checksumType,omitempty"`
		SizeBytes    *int64                `json:"sizeBytes,omitempty"`
	}{
		ImageFileProxy: (*ImageFileProxy)(p),
		Checksum:       p.Checksum,
		ChecksumType:   p.ChecksumType,
		SizeBytes:      p.SizeBytes,
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

func (p *ImageFile) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ImageFile
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ImageFile(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "checksum")
	delete(allFields, "checksumType")
	delete(allFields, "fileLocationId")
	delete(allFields, "filePath")
	delete(allFields, "name")
	delete(allFields, "sizeBytes")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewImageFile() *ImageFile {
	p := new(ImageFile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ImageFile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operation type and UUID of an ongoing operation in LCM.
*/
type InProgressOpInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Root task UUID of the operation, if it is in running state.
	*/
	OperationId *string `json:"operationId,omitempty"`

	OperationType *import1.OperationType `json:"operationType,omitempty"`
}

func (p *InProgressOpInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias InProgressOpInfo

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

func (p *InProgressOpInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias InProgressOpInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = InProgressOpInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "operationId")
	delete(allFields, "operationType")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewInProgressOpInfo() *InProgressOpInfo {
	p := new(InProgressOpInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.InProgressOpInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Summary of LCM state on a cluster.
*/
type LcmSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM framework version present in the LCM URL.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/*
	  List of capabilities of cluster. These capabilities are used to determine the features supported by LCM in the cluster.
	*/
	Capabilities []Capability `json:"capabilities,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Current LCM Version.
	*/
	CurrentVersion *string `json:"currentVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Hardware vendor information.
	*/
	HardwareVendor *string `json:"hardwareVendor,omitempty"`

	InProgressOperation *import1.InProgressOpDetails `json:"inProgressOperation,omitempty"`
	/*
	  Boolean that indicates if LCM url is accessible or not
	*/
	IsUrlAccessible *bool `json:"isUrlAccessible,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *LcmSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LcmSummary

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

func (p *LcmSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LcmSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LcmSummary(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableVersion")
	delete(allFields, "capabilities")
	delete(allFields, "clusterExtId")
	delete(allFields, "currentVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareVendor")
	delete(allFields, "inProgressOperation")
	delete(allFields, "isUrlAccessible")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLcmSummary() *LcmSummary {
	p := new(LcmSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.LcmSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsUrlAccessible = new(bool)
	*p.IsUrlAccessible = true

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.1/resources/bundles Get operation
*/
type ListBundlesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListBundlesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListBundlesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListBundlesApiResponse

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

func (p *ListBundlesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListBundlesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListBundlesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListBundlesApiResponse() *ListBundlesApiResponse {
	p := new(ListBundlesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListBundlesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListBundlesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListBundlesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListBundlesApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/entities Get operation
*/
type ListEntitiesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListEntitiesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListEntitiesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListEntitiesApiResponse

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

func (p *ListEntitiesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListEntitiesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListEntitiesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListEntitiesApiResponse() *ListEntitiesApiResponse {
	p := new(ListEntitiesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListEntitiesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListEntitiesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListEntitiesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListEntitiesApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/images Get operation
*/
type ListImagesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListImagesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListImagesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListImagesApiResponse

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

func (p *ListImagesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListImagesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListImagesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListImagesApiResponse() *ListImagesApiResponse {
	p := new(ListImagesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListImagesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListImagesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListImagesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListImagesApiResponseData()
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
REST response for all response codes in API path /lifecycle/v4.1/resources/lcm-summaries Get operation
*/
type ListLcmSummariesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListLcmSummariesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListLcmSummariesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListLcmSummariesApiResponse

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

func (p *ListLcmSummariesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListLcmSummariesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListLcmSummariesApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewListLcmSummariesApiResponse() *ListLcmSummariesApiResponse {
	p := new(ListLcmSummariesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.ListLcmSummariesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListLcmSummariesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListLcmSummariesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListLcmSummariesApiResponseData()
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
The computed LCM upgrade notifications for the given input.
*/
type Notification struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	Notifications []NotificationItem `json:"notifications,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *Notification) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Notification

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

func (p *Notification) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Notification
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Notification(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "notifications")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotification() *Notification {
	p := new(Notification)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.Notification"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Detailed LCM upgrade notification information for this entity.
*/
type NotificationDetail struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A message with notification details. The description of the most disruptive action that will occur on the node or the cluster. INFO, WARNING or NOTICE based message for an entity.
	*/
	Message *string `json:"message,omitempty"`

	SeverityLevel *SeverityLevel `json:"severityLevel,omitempty"`
}

func (p *NotificationDetail) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotificationDetail

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

func (p *NotificationDetail) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotificationDetail
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotificationDetail(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")
	delete(allFields, "severityLevel")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotificationDetail() *NotificationDetail {
	p := new(NotificationDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.NotificationDetail"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM upgrade notification generated for a node or cluster based on specified entity/entities and target version(s).
*/
type NotificationItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of upgrade notifications for this entity.
	*/
	Details []NotificationDetail `json:"details,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`

	HypervisorType *import1.HypervisorType `json:"hypervisorType,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`

	NotificationType *import1.NotificationType `json:"notificationType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Version to upgrade to.
	*/
	ToVersion *string `json:"toVersion,omitempty"`
}

func (p *NotificationItem) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NotificationItem

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

func (p *NotificationItem) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotificationItem
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotificationItem(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "details")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "hypervisorType")
	delete(allFields, "links")
	delete(allFields, "locationInfo")
	delete(allFields, "notificationType")
	delete(allFields, "tenantId")
	delete(allFields, "toVersion")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotificationItem() *NotificationItem {
	p := new(NotificationItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.NotificationItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM upgrade notification generated for a node or cluster based on specified entity/entities and target version(s).
*/
type NotificationsSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Specification of credentials to be provided by the user to perform LCM operations.
	*/
	Credentials []import1.Credential `json:"credentials,omitempty"`
	/*
	  The computed LCM upgrade notifications for the given input.
	*/
	NotificationsSpec []import1.EntityUpdateSpec `json:"notificationsSpec"`
}

func (p *NotificationsSpec) MarshalJSON() ([]byte, error) {
	type NotificationsSpecProxy NotificationsSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NotificationsSpecProxy
		NotificationsSpec []import1.EntityUpdateSpec `json:"notificationsSpec,omitempty"`
	}{
		NotificationsSpecProxy: (*NotificationsSpecProxy)(p),
		NotificationsSpec:      p.NotificationsSpec,
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

func (p *NotificationsSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NotificationsSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NotificationsSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "credentials")
	delete(allFields, "notificationsSpec")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNotificationsSpec() *NotificationsSpec {
	p := new(NotificationsSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.NotificationsSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Upgrade recommendations for LCM entity/entities.
*/
type RecommendationResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of added LCM entities to the input recommendations specification.
	*/
	AddableEntities []UpdatedTargetEntityResult `json:"addableEntities,omitempty"`
	/*
	  Cluster uuid on which the resource is present or operation is being performed.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  List of deployable entities and their dependencies.
	*/
	DeployableVersions []DeployableVersion `json:"deployableVersions,omitempty"`
	/*
	  List of entity update objects for getting recommendations.
	*/
	EntityUpdateSpecs []import1.EntityUpdateSpec `json:"entityUpdateSpecs,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of modified LCM entities from the input recommendations specification.
	*/
	ModifiableEntities []UpdatedTargetEntityResult `json:"modifiableEntities,omitempty"`
	/*
	  List of skipped LCM entities from the input recommendations specification.
	*/
	SkippedEntities []UpdatedTargetEntityResult `json:"skippedEntities,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RecommendationResult) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecommendationResult

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

func (p *RecommendationResult) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecommendationResult
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RecommendationResult(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "addableEntities")
	delete(allFields, "clusterExtId")
	delete(allFields, "deployableVersions")
	delete(allFields, "entityUpdateSpecs")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "modifiableEntities")
	delete(allFields, "skippedEntities")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRecommendationResult() *RecommendationResult {
	p := new(RecommendationResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.RecommendationResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specification to get upgrade recommendations for specific UUID and target version via LCM Recommendation
*/
type RecommendationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	RecommendationSpecItemDiscriminator_ *string `json:"$recommendationSpecItemDiscriminator,omitempty"`
	/*
	  Specification to get upgrade recommendations for specific UUID and target version via LCM Recommendation
	*/
	RecommendationSpec *OneOfRecommendationSpecRecommendationSpec `json:"recommendationSpec"`
}

func (p *RecommendationSpec) MarshalJSON() ([]byte, error) {
	type RecommendationSpecProxy RecommendationSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RecommendationSpecProxy
		RecommendationSpec *OneOfRecommendationSpecRecommendationSpec `json:"recommendationSpec,omitempty"`
	}{
		RecommendationSpecProxy: (*RecommendationSpecProxy)(p),
		RecommendationSpec:      p.RecommendationSpec,
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

func (p *RecommendationSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecommendationSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RecommendationSpec(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$recommendationSpecItemDiscriminator")
	delete(allFields, "recommendationSpec")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRecommendationSpec() *RecommendationSpec {
	p := new(RecommendationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.RecommendationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecommendationSpec) GetRecommendationSpec() interface{} {
	if nil == p.RecommendationSpec {
		return nil
	}
	return p.RecommendationSpec.GetValue()
}

func (p *RecommendationSpec) SetRecommendationSpec(v interface{}) error {
	if nil == p.RecommendationSpec {
		p.RecommendationSpec = NewOneOfRecommendationSpecRecommendationSpec()
	}
	e := p.RecommendationSpec.SetValue(v)
	if nil == e {
		if nil == p.RecommendationSpecItemDiscriminator_ {
			p.RecommendationSpecItemDiscriminator_ = new(string)
		}
		*p.RecommendationSpecItemDiscriminator_ = *p.RecommendationSpec.Discriminator
	}
	return e
}

/*
Severity level of LCM notification.
*/
type SeverityLevel int

const (
	SEVERITYLEVEL_UNKNOWN  SeverityLevel = 0
	SEVERITYLEVEL_REDACTED SeverityLevel = 1
	SEVERITYLEVEL_WARNING  SeverityLevel = 2
	SEVERITYLEVEL_NOTICE   SeverityLevel = 3
	SEVERITYLEVEL_INFO     SeverityLevel = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SeverityLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SeverityLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SeverityLevel) index(name string) SeverityLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"WARNING",
		"NOTICE",
		"INFO",
	}
	for idx := range names {
		if names[idx] == name {
			return SeverityLevel(idx)
		}
	}
	return SEVERITYLEVEL_UNKNOWN
}

func (e *SeverityLevel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SeverityLevel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SeverityLevel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SeverityLevel) Ref() *SeverityLevel {
	return &e
}

/*
Detailed information about the current LCM framework status.
*/
type StatusInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FrameworkVersion *FrameworkVersionInfo `json:"frameworkVersion,omitempty"`

	InProgressOperation *InProgressOpInfo `json:"inProgressOperation,omitempty"`
	/*
	  Boolean that indicates if cancel intent for LCM update is set or not.
	*/
	IsCancelIntentSet *bool `json:"isCancelIntentSet,omitempty"`
	/*
	  Boolean that indicates if LCM url is accessible or not
	*/
	IsUrlAccessible *bool `json:"isUrlAccessible,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Upload task UUID.
	*/
	UploadTaskUuid *string `json:"uploadTaskUuid,omitempty"`
}

func (p *StatusInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StatusInfo

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

func (p *StatusInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StatusInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StatusInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "frameworkVersion")
	delete(allFields, "inProgressOperation")
	delete(allFields, "isCancelIntentSet")
	delete(allFields, "isUrlAccessible")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "uploadTaskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStatusInfo() *StatusInfo {
	p := new(StatusInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.StatusInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsCancelIntentSet = new(bool)
	*p.IsCancelIntentSet = false
	p.IsUrlAccessible = new(bool)
	*p.IsUrlAccessible = true

	return p
}

/*
LCM target entity for which recommendations are requested.
*/
type TargetEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Unique identifier of an LCM entity e.g. "HDD serial number".
	*/
	DeviceId *string `json:"deviceId,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The requested update version of an LCM entity.
	*/
	Version *string `json:"version"`
}

func (p *TargetEntity) MarshalJSON() ([]byte, error) {
	type TargetEntityProxy TargetEntity

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*TargetEntityProxy
		Version *string `json:"version,omitempty"`
	}{
		TargetEntityProxy: (*TargetEntityProxy)(p),
		Version:           p.Version,
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

func (p *TargetEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TargetEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TargetEntity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "deviceId")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "links")
	delete(allFields, "locationInfo")
	delete(allFields, "tenantId")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTargetEntity() *TargetEntity {
	p := new(TargetEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.TargetEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /lifecycle/v4.1/resources/config Put operation
*/
type UpdateConfigApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateConfigApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateConfigApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateConfigApiResponse

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

func (p *UpdateConfigApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateConfigApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdateConfigApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUpdateConfigApiResponse() *UpdateConfigApiResponse {
	p := new(UpdateConfigApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdateConfigApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateConfigApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateConfigApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateConfigApiResponseData()
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
Updated LCM target entity in recommendation result.
*/
type UpdatedTargetEntity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  LCM entity class.
	*/
	EntityClass *string `json:"entityClass,omitempty"`
	/*
	  LCM entity model.
	*/
	EntityModel *string `json:"entityModel,omitempty"`

	EntityType *import1.EntityType `json:"entityType,omitempty"`
	/*
	  Current version of an LCM entity.
	*/
	EntityVersion *string `json:"entityVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A hardware family for a LCM entity.
	*/
	HardwareFamily *string `json:"hardwareFamily,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LocationInfo *import1.LocationInfo `json:"locationInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *UpdatedTargetEntity) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdatedTargetEntity

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

func (p *UpdatedTargetEntity) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdatedTargetEntity
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdatedTargetEntity(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityClass")
	delete(allFields, "entityModel")
	delete(allFields, "entityType")
	delete(allFields, "entityVersion")
	delete(allFields, "extId")
	delete(allFields, "hardwareFamily")
	delete(allFields, "links")
	delete(allFields, "locationInfo")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUpdatedTargetEntity() *UpdatedTargetEntity {
	p := new(UpdatedTargetEntity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdatedTargetEntity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
LCM error for target entity
*/
type UpdatedTargetEntityResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Error message for the target entity that failed in the input recommendations specification.
	*/
	Message *string `json:"message"`

	TargetEntity *UpdatedTargetEntity `json:"targetEntity,omitempty"`
}

func (p *UpdatedTargetEntityResult) MarshalJSON() ([]byte, error) {
	type UpdatedTargetEntityResultProxy UpdatedTargetEntityResult

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UpdatedTargetEntityResultProxy
		Message *string `json:"message,omitempty"`
	}{
		UpdatedTargetEntityResultProxy: (*UpdatedTargetEntityResultProxy)(p),
		Message:                        p.Message,
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

func (p *UpdatedTargetEntityResult) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdatedTargetEntityResult
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = UpdatedTargetEntityResult(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")
	delete(allFields, "targetEntity")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUpdatedTargetEntityResult() *UpdatedTargetEntityResult {
	p := new(UpdatedTargetEntityResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.resources.UpdatedTargetEntityResult"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetEntityByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Entity                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetEntityByIdApiResponseData() *OneOfGetEntityByIdApiResponseData {
	p := new(OneOfGetEntityByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetEntityByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetEntityByIdApiResponseData is nil"))
	}
	switch v.(type) {
	case Entity:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Entity)
		}
		*p.oneOfType2001 = v.(Entity)
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

func (p *OneOfGetEntityByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetEntityByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Entity)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Entity" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Entity)
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
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetEntityByIdApiResponseData"))
}

func (p *OneOfGetEntityByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetEntityByIdApiResponseData")
}

type OneOfRecommendationSpecRecommendationSpec struct {
	Discriminator *string                    `json:"-"`
	ObjectType_   *string                    `json:"-"`
	oneOfType2002 []TargetEntity             `json:"-"`
	oneOfType2003 []import1.EntityUpdateSpec `json:"-"`
	oneOfType2001 []import1.EntityType       `json:"-"`
	oneOfType2004 []import1.EntityDeploySpec `json:"-"`
}

func NewOneOfRecommendationSpecRecommendationSpec() *OneOfRecommendationSpecRecommendationSpec {
	p := new(OneOfRecommendationSpecRecommendationSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecommendationSpecRecommendationSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecommendationSpecRecommendationSpec is nil"))
	}
	switch v.(type) {
	case []TargetEntity:
		p.oneOfType2002 = v.([]TargetEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.TargetEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.TargetEntity>"
	case []import1.EntityUpdateSpec:
		p.oneOfType2003 = v.([]import1.EntityUpdateSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityUpdateSpec>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityUpdateSpec>"
	case []import1.EntityType:
		p.oneOfType2001 = v.([]import1.EntityType)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityType>"
	case []import1.EntityDeploySpec:
		p.oneOfType2004 = v.([]import1.EntityDeploySpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityDeploySpec>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityDeploySpec>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRecommendationSpecRecommendationSpec) GetValue() interface{} {
	if "List<lifecycle.v4.resources.TargetEntity>" == *p.Discriminator {
		return p.oneOfType2002
	}
	if "List<lifecycle.v4.common.EntityUpdateSpec>" == *p.Discriminator {
		return p.oneOfType2003
	}
	if "List<lifecycle.v4.common.EntityType>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if "List<lifecycle.v4.common.EntityDeploySpec>" == *p.Discriminator {
		return p.oneOfType2004
	}
	return nil
}

func (p *OneOfRecommendationSpecRecommendationSpec) UnmarshalJSON(b []byte) error {
	vOneOfType2002 := new([]TargetEntity)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if len(*vOneOfType2002) == 0 || "lifecycle.v4.resources.TargetEntity" == *((*vOneOfType2002)[0].ObjectType_) {
			p.oneOfType2002 = *vOneOfType2002
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.TargetEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.TargetEntity>"
			return nil
		}
	}
	vOneOfType2003 := new([]import1.EntityUpdateSpec)
	if err := json.Unmarshal(b, vOneOfType2003); err == nil {
		if len(*vOneOfType2003) == 0 || "lifecycle.v4.common.EntityUpdateSpec" == *((*vOneOfType2003)[0].ObjectType_) {
			p.oneOfType2003 = *vOneOfType2003
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.common.EntityUpdateSpec>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.common.EntityUpdateSpec>"
			return nil
		}
	}
	vOneOfType2001 := new([]import1.EntityType)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		p.oneOfType2001 = *vOneOfType2001
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.common.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.common.EntityType>"
		return nil
	}
	vOneOfType2004 := new([]import1.EntityDeploySpec)
	if err := json.Unmarshal(b, vOneOfType2004); err == nil {
		if len(*vOneOfType2004) == 0 || "lifecycle.v4.common.EntityDeploySpec" == *((*vOneOfType2004)[0].ObjectType_) {
			p.oneOfType2004 = *vOneOfType2004
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.common.EntityDeploySpec>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.common.EntityDeploySpec>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecommendationSpecRecommendationSpec"))
}

func (p *OneOfRecommendationSpecRecommendationSpec) MarshalJSON() ([]byte, error) {
	if "List<lifecycle.v4.resources.TargetEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	if "List<lifecycle.v4.common.EntityUpdateSpec>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2003)
	}
	if "List<lifecycle.v4.common.EntityType>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if "List<lifecycle.v4.common.EntityDeploySpec>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2004)
	}
	return nil, errors.New("No value to marshal for OneOfRecommendationSpecRecommendationSpec")
}

type OneOfUpdateConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfUpdateConfigApiResponseData() *OneOfUpdateConfigApiResponseData {
	p := new(OneOfUpdateConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateConfigApiResponseData is nil"))
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

func (p *OneOfUpdateConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
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
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateConfigApiResponseData"))
}

func (p *OneOfUpdateConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateConfigApiResponseData")
}

type OneOfListEntitiesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []Entity               `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListEntitiesApiResponseData() *OneOfListEntitiesApiResponseData {
	p := new(OneOfListEntitiesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListEntitiesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListEntitiesApiResponseData is nil"))
	}
	switch v.(type) {
	case []Entity:
		p.oneOfType2001 = v.([]Entity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.Entity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.Entity>"
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

func (p *OneOfListEntitiesApiResponseData) GetValue() interface{} {
	if "List<lifecycle.v4.resources.Entity>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListEntitiesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]Entity)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.Entity" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.Entity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.Entity>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListEntitiesApiResponseData"))
}

func (p *OneOfListEntitiesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<lifecycle.v4.resources.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListEntitiesApiResponseData")
}

type OneOfListBundlesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Bundle               `json:"-"`
}

func NewOneOfListBundlesApiResponseData() *OneOfListBundlesApiResponseData {
	p := new(OneOfListBundlesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListBundlesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListBundlesApiResponseData is nil"))
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
	case []Bundle:
		p.oneOfType2001 = v.([]Bundle)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.Bundle>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.Bundle>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListBundlesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.resources.Bundle>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListBundlesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]Bundle)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.Bundle" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.Bundle>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.Bundle>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListBundlesApiResponseData"))
}

func (p *OneOfListBundlesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.resources.Bundle>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListBundlesApiResponseData")
}

type OneOfCreateBundleApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfCreateBundleApiResponseData() *OneOfCreateBundleApiResponseData {
	p := new(OneOfCreateBundleApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateBundleApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateBundleApiResponseData is nil"))
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

func (p *OneOfCreateBundleApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfCreateBundleApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
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
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateBundleApiResponseData"))
}

func (p *OneOfCreateBundleApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfCreateBundleApiResponseData")
}

type OneOfListLcmSummariesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []LcmSummary           `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfListLcmSummariesApiResponseData() *OneOfListLcmSummariesApiResponseData {
	p := new(OneOfListLcmSummariesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListLcmSummariesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListLcmSummariesApiResponseData is nil"))
	}
	switch v.(type) {
	case []LcmSummary:
		p.oneOfType2001 = v.([]LcmSummary)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.LcmSummary>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.LcmSummary>"
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

func (p *OneOfListLcmSummariesApiResponseData) GetValue() interface{} {
	if "List<lifecycle.v4.resources.LcmSummary>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListLcmSummariesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]LcmSummary)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.LcmSummary" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.LcmSummary>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.LcmSummary>"
			return nil
		}
	}
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListLcmSummariesApiResponseData"))
}

func (p *OneOfListLcmSummariesApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<lifecycle.v4.resources.LcmSummary>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListLcmSummariesApiResponseData")
}

type OneOfGetStatusApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *StatusInfo            `json:"-"`
}

func NewOneOfGetStatusApiResponseData() *OneOfGetStatusApiResponseData {
	p := new(OneOfGetStatusApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetStatusApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetStatusApiResponseData is nil"))
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
	case StatusInfo:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(StatusInfo)
		}
		*p.oneOfType2001 = v.(StatusInfo)
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

func (p *OneOfGetStatusApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetStatusApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(StatusInfo)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.StatusInfo" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(StatusInfo)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetStatusApiResponseData"))
}

func (p *OneOfGetStatusApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetStatusApiResponseData")
}

type OneOfBundleChecksum struct {
	Discriminator *string               `json:"-"`
	ObjectType_   *string               `json:"-"`
	oneOfType2401 *import1.LcmSha256Sum `json:"-"`
	oneOfType2402 *import1.LcmMd5Sum    `json:"-"`
}

func NewOneOfBundleChecksum() *OneOfBundleChecksum {
	p := new(OneOfBundleChecksum)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfBundleChecksum) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfBundleChecksum is nil"))
	}
	switch v.(type) {
	case import1.LcmSha256Sum:
		if nil == p.oneOfType2401 {
			p.oneOfType2401 = new(import1.LcmSha256Sum)
		}
		*p.oneOfType2401 = v.(import1.LcmSha256Sum)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2401.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2401.ObjectType_
	case import1.LcmMd5Sum:
		if nil == p.oneOfType2402 {
			p.oneOfType2402 = new(import1.LcmMd5Sum)
		}
		*p.oneOfType2402 = v.(import1.LcmMd5Sum)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2402.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2402.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfBundleChecksum) GetValue() interface{} {
	if p.oneOfType2401 != nil && *p.oneOfType2401.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2401
	}
	if p.oneOfType2402 != nil && *p.oneOfType2402.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2402
	}
	return nil
}

func (p *OneOfBundleChecksum) UnmarshalJSON(b []byte) error {
	vOneOfType2401 := new(import1.LcmSha256Sum)
	if err := json.Unmarshal(b, vOneOfType2401); err == nil {
		if "lifecycle.v4.common.LcmSha256Sum" == *vOneOfType2401.ObjectType_ {
			if nil == p.oneOfType2401 {
				p.oneOfType2401 = new(import1.LcmSha256Sum)
			}
			*p.oneOfType2401 = *vOneOfType2401
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2401.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2401.ObjectType_
			return nil
		}
	}
	vOneOfType2402 := new(import1.LcmMd5Sum)
	if err := json.Unmarshal(b, vOneOfType2402); err == nil {
		if "lifecycle.v4.common.LcmMd5Sum" == *vOneOfType2402.ObjectType_ {
			if nil == p.oneOfType2402 {
				p.oneOfType2402 = new(import1.LcmMd5Sum)
			}
			*p.oneOfType2402 = *vOneOfType2402
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2402.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2402.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfBundleChecksum"))
}

func (p *OneOfBundleChecksum) MarshalJSON() ([]byte, error) {
	if p.oneOfType2401 != nil && *p.oneOfType2401.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2401)
	}
	if p.oneOfType2402 != nil && *p.oneOfType2402.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2402)
	}
	return nil, errors.New("No value to marshal for OneOfBundleChecksum")
}

type OneOfGetRecommendationByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *RecommendationResult  `json:"-"`
}

func NewOneOfGetRecommendationByIdApiResponseData() *OneOfGetRecommendationByIdApiResponseData {
	p := new(OneOfGetRecommendationByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRecommendationByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRecommendationByIdApiResponseData is nil"))
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
	case RecommendationResult:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(RecommendationResult)
		}
		*p.oneOfType2001 = v.(RecommendationResult)
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

func (p *OneOfGetRecommendationByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetRecommendationByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(RecommendationResult)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.RecommendationResult" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(RecommendationResult)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRecommendationByIdApiResponseData"))
}

func (p *OneOfGetRecommendationByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetRecommendationByIdApiResponseData")
}

type OneOfGetLcmSummaryByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *LcmSummary            `json:"-"`
}

func NewOneOfGetLcmSummaryByIdApiResponseData() *OneOfGetLcmSummaryByIdApiResponseData {
	p := new(OneOfGetLcmSummaryByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLcmSummaryByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLcmSummaryByIdApiResponseData is nil"))
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
	case LcmSummary:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(LcmSummary)
		}
		*p.oneOfType2001 = v.(LcmSummary)
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

func (p *OneOfGetLcmSummaryByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetLcmSummaryByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(LcmSummary)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.LcmSummary" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(LcmSummary)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLcmSummaryByIdApiResponseData"))
}

func (p *OneOfGetLcmSummaryByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetLcmSummaryByIdApiResponseData")
}

type OneOfGetNotificationsByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *Notification          `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfGetNotificationsByIdApiResponseData() *OneOfGetNotificationsByIdApiResponseData {
	p := new(OneOfGetNotificationsByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNotificationsByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNotificationsByIdApiResponseData is nil"))
	}
	switch v.(type) {
	case Notification:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Notification)
		}
		*p.oneOfType2001 = v.(Notification)
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

func (p *OneOfGetNotificationsByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNotificationsByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Notification)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Notification" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Notification)
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
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNotificationsByIdApiResponseData"))
}

func (p *OneOfGetNotificationsByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNotificationsByIdApiResponseData")
}

type OneOfGetBundleByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *Bundle                `json:"-"`
}

func NewOneOfGetBundleByIdApiResponseData() *OneOfGetBundleByIdApiResponseData {
	p := new(OneOfGetBundleByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetBundleByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetBundleByIdApiResponseData is nil"))
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
	case Bundle:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Bundle)
		}
		*p.oneOfType2001 = v.(Bundle)
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

func (p *OneOfGetBundleByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetBundleByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(Bundle)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Bundle" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Bundle)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetBundleByIdApiResponseData"))
}

func (p *OneOfGetBundleByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetBundleByIdApiResponseData")
}

type OneOfGetConfigApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 *Config                `json:"-"`
}

func NewOneOfGetConfigApiResponseData() *OneOfGetConfigApiResponseData {
	p := new(OneOfGetConfigApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetConfigApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetConfigApiResponseData is nil"))
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
	case Config:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Config)
		}
		*p.oneOfType2001 = v.(Config)
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

func (p *OneOfGetConfigApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetConfigApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(Config)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "lifecycle.v4.resources.Config" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Config)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetConfigApiResponseData"))
}

func (p *OneOfGetConfigApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetConfigApiResponseData")
}

type OneOfDeleteBundleByIdApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *import3.TaskReference `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
}

func NewOneOfDeleteBundleByIdApiResponseData() *OneOfDeleteBundleByIdApiResponseData {
	p := new(OneOfDeleteBundleByIdApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteBundleByIdApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteBundleByIdApiResponseData is nil"))
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

func (p *OneOfDeleteBundleByIdApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteBundleByIdApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
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
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteBundleByIdApiResponseData"))
}

func (p *OneOfDeleteBundleByIdApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteBundleByIdApiResponseData")
}

type OneOfListImagesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import4.ErrorResponse `json:"-"`
	oneOfType2001 []Image                `json:"-"`
}

func NewOneOfListImagesApiResponseData() *OneOfListImagesApiResponseData {
	p := new(OneOfListImagesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListImagesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListImagesApiResponseData is nil"))
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
	case []Image:
		p.oneOfType2001 = v.([]Image)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<lifecycle.v4.resources.Image>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<lifecycle.v4.resources.Image>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListImagesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<lifecycle.v4.resources.Image>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListImagesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import4.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "lifecycle.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new([]Image)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "lifecycle.v4.resources.Image" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<lifecycle.v4.resources.Image>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<lifecycle.v4.resources.Image>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListImagesApiResponseData"))
}

func (p *OneOfListImagesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<lifecycle.v4.resources.Image>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListImagesApiResponseData")
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
