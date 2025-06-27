/*
 * Generated file models/catalogCommon/v1/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module catalogCommon.v1.config of Nutanix Virtual Machine Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	"time"
)

/*
Category-based entity filter.
*/
type CategoriesFilter struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Filter to match entities based on the provided categories.
	*/
	CategoryExtIds []string `json:"categoryExtIds"`

	Type *CategoriesMatchType `json:"type"`
}

func (p *CategoriesFilter) MarshalJSON() ([]byte, error) {
	type CategoriesFilterProxy CategoriesFilter

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CategoriesFilterProxy
		CategoryExtIds []string             `json:"categoryExtIds,omitempty"`
		Type           *CategoriesMatchType `json:"type,omitempty"`
	}{
		CategoriesFilterProxy: (*CategoriesFilterProxy)(p),
		CategoryExtIds:        p.CategoryExtIds,
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

func (p *CategoriesFilter) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CategoriesFilter
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CategoriesFilter(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "categoryExtIds")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCategoriesFilter() *CategoriesFilter {
	p := new(CategoriesFilter)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "catalogCommon.v1.config.CategoriesFilter"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Filter matching type.
*/
type CategoriesMatchType int

const (
	CATEGORIESMATCHTYPE_UNKNOWN              CategoriesMatchType = 0
	CATEGORIESMATCHTYPE_REDACTED             CategoriesMatchType = 1
	CATEGORIESMATCHTYPE_CATEGORIES_MATCH_ALL CategoriesMatchType = 2
	CATEGORIESMATCHTYPE_CATEGORIES_MATCH_ANY CategoriesMatchType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CategoriesMatchType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORIES_MATCH_ALL",
		"CATEGORIES_MATCH_ANY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CategoriesMatchType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORIES_MATCH_ALL",
		"CATEGORIES_MATCH_ANY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CategoriesMatchType) index(name string) CategoriesMatchType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CATEGORIES_MATCH_ALL",
		"CATEGORIES_MATCH_ANY",
	}
	for idx := range names {
		if names[idx] == name {
			return CategoriesMatchType(idx)
		}
	}
	return CATEGORIESMATCHTYPE_UNKNOWN
}

func (e *CategoriesMatchType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CategoriesMatchType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CategoriesMatchType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CategoriesMatchType) Ref() *CategoriesMatchType {
	return &e
}

/*
Placement policy model.
*/
type CommonPlacementPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterFilter *CategoriesFilter `json:"clusterFilter,omitempty"`

	ContentFilter *CategoriesFilter `json:"contentFilter,omitempty"`
	/*
	  The time when the placement policy was created.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  External identifier of the user who created the placement policy.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Description of the placement policy.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the placement policy.
	*/
	Name *string `json:"name,omitempty"`

	PlacementType *ContentPlacementType `json:"placementType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The time when the placement policy was last updated.
	*/
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	/*
	  External identifier of the user who updated the placement policy.
	*/
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

func (p *CommonPlacementPolicy) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CommonPlacementPolicy

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

func (p *CommonPlacementPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CommonPlacementPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CommonPlacementPolicy(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterFilter")
	delete(allFields, "contentFilter")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "placementType")
	delete(allFields, "tenantId")
	delete(allFields, "updateTime")
	delete(allFields, "updatedBy")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCommonPlacementPolicy() *CommonPlacementPolicy {
	p := new(CommonPlacementPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "catalogCommon.v1.config.CommonPlacementPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of placement policy.
*/
type ContentPlacementType int

const (
	CONTENTPLACEMENTTYPE_UNKNOWN  ContentPlacementType = 0
	CONTENTPLACEMENTTYPE_REDACTED ContentPlacementType = 1
	CONTENTPLACEMENTTYPE_SOFT     ContentPlacementType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ContentPlacementType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ContentPlacementType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ContentPlacementType) index(name string) ContentPlacementType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFT",
	}
	for idx := range names {
		if names[idx] == name {
			return ContentPlacementType(idx)
		}
	}
	return CONTENTPLACEMENTTYPE_UNKNOWN
}

func (e *ContentPlacementType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ContentPlacementType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ContentPlacementType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ContentPlacementType) Ref() *ContentPlacementType {
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
