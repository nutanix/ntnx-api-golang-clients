/*
 * Generated file models/multidomain/v4/config/config_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Monitoring APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module multidomain.v4.config of Nutanix Monitoring APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/common/v1/response"
)

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
	  ID of the project
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
	  Name of the project
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3.b1"}
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
	  ID of the project
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
	  Name of the project
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3.b1"}
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
