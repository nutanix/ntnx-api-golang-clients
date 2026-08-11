/*
 * Generated file models/prism/v4/operations/operations_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Volumes APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.operations of Nutanix Volumes APIs
*/
package operations

import (
	"encoding/json"
)

/*
The metadata section on the input specification for performing the batch operation.
*/
type BatchSpecPayloadMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Headers []BatchSpecPayloadMetadataHeader `json:"headers,omitempty"`

	Path []BatchSpecPayloadMetadataPath `json:"path,omitempty"`
}

func (p *BatchSpecPayloadMetadata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BatchSpecPayloadMetadata

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

func (p *BatchSpecPayloadMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BatchSpecPayloadMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBatchSpecPayloadMetadata()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Headers != nil {
		p.Headers = known.Headers
	}
	if known.Path != nil {
		p.Path = known.Path
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "headers")
	delete(allFields, "path")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBatchSpecPayloadMetadata() *BatchSpecPayloadMetadata {
	p := new(BatchSpecPayloadMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.operations.BatchSpecPayloadMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The metadata section on the input specification for performing the batch operation.
*/
type BatchSpecPayloadMetadataHeader struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the header parameter.
	*/
	Name *string `json:"name"`
	/*
	  The value of the header parameter.
	*/
	Value *string `json:"value"`
}

func (p *BatchSpecPayloadMetadataHeader) MarshalJSON() ([]byte, error) {
	type BatchSpecPayloadMetadataHeaderProxy BatchSpecPayloadMetadataHeader

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BatchSpecPayloadMetadataHeaderProxy
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	}{
		BatchSpecPayloadMetadataHeaderProxy: (*BatchSpecPayloadMetadataHeaderProxy)(p),
		Name:                                p.Name,
		Value:                               p.Value,
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

func (p *BatchSpecPayloadMetadataHeader) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BatchSpecPayloadMetadataHeader
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBatchSpecPayloadMetadataHeader()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBatchSpecPayloadMetadataHeader() *BatchSpecPayloadMetadataHeader {
	p := new(BatchSpecPayloadMetadataHeader)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.operations.BatchSpecPayloadMetadataHeader"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The metadata section on the input specification for performing the batch operation.
*/
type BatchSpecPayloadMetadataPath struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The name of the  path parameter.
	*/
	Name *string `json:"name"`
	/*
	  The value of the  path parameter.
	*/
	Value *string `json:"value"`
}

func (p *BatchSpecPayloadMetadataPath) MarshalJSON() ([]byte, error) {
	type BatchSpecPayloadMetadataPathProxy BatchSpecPayloadMetadataPath

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BatchSpecPayloadMetadataPathProxy
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	}{
		BatchSpecPayloadMetadataPathProxy: (*BatchSpecPayloadMetadataPathProxy)(p),
		Name:                              p.Name,
		Value:                             p.Value,
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

func (p *BatchSpecPayloadMetadataPath) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BatchSpecPayloadMetadataPath
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBatchSpecPayloadMetadataPath()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBatchSpecPayloadMetadataPath() *BatchSpecPayloadMetadataPath {
	p := new(BatchSpecPayloadMetadataPath)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.operations.BatchSpecPayloadMetadataPath"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
