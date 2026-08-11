/*
 * Generated file models/catalogCommon/v1/config/config_model.go.
 *
 * Product version: 4.4.1-beta-1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module catalogCommon.v1.config of Nutanix Multidomain Versioned APIs
*/
package config

import (
	"encoding/json"
)

/*
The Object details for creating entity.
*/
type ObjectStoreSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Key that identifies the source Object in the bucket.
	*/
	Key *string `json:"key"`
}

func (p *ObjectStoreSource) MarshalJSON() ([]byte, error) {
	type ObjectStoreSourceProxy ObjectStoreSource

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ObjectStoreSourceProxy
		Key *string `json:"key,omitempty"`
	}{
		ObjectStoreSourceProxy: (*ObjectStoreSourceProxy)(p),
		Key:                    p.Key,
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

func (p *ObjectStoreSource) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectStoreSource
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectStoreSource()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Key != nil {
		p.Key = known.Key
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "key")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectStoreSource() *ObjectStoreSource {
	p := new(ObjectStoreSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "catalogCommon.v1.config.ObjectStoreSource"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The source URL details for creating an entity.
*/
type UrlSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Ignore the certificate errors, if the value is true. Default is false.
	*/
	ShouldAllowInsecureUrl *bool `json:"shouldAllowInsecureUrl,omitempty"`
	/*
	  URL to create an entity.
	*/
	Url *string `json:"url"`
}

func (p *UrlSource) MarshalJSON() ([]byte, error) {
	type UrlSourceProxy UrlSource

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*UrlSourceProxy
		Url *string `json:"url,omitempty"`
	}{
		UrlSourceProxy: (*UrlSourceProxy)(p),
		Url:            p.Url,
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

func (p *UrlSource) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UrlSource
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUrlSource()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ShouldAllowInsecureUrl != nil {
		p.ShouldAllowInsecureUrl = known.ShouldAllowInsecureUrl
	}
	if known.Url != nil {
		p.Url = known.Url
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "shouldAllowInsecureUrl")
	delete(allFields, "url")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUrlSource() *UrlSource {
	p := new(UrlSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "catalogCommon.v1.config.UrlSource"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldAllowInsecureUrl = new(bool)
	*p.ShouldAllowInsecureUrl = false

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
