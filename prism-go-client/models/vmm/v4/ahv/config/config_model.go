/*
 * Generated file models/vmm/v4/ahv/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.ahv.config of Nutanix Prism APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
)

/*
If this field is set, the guest will be customized using cloud-init. Either user_data or custom_key_values should be provided. If custom_key_ves are provided then the user data will be generated using these key-value pairs.
*/
type CloudInit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	CloudInitScriptItemDiscriminator_ *string `json:"$cloudInitScriptItemDiscriminator,omitempty"`
	/*
	  The script to use for cloud-init.
	*/
	CloudInitScript *OneOfCloudInitCloudInitScript `json:"cloudInitScript,omitempty"`

	DatasourceType *CloudInitDataSourceType `json:"datasourceType,omitempty"`
	/*
	  The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.
	*/
	Metadata *string `json:"metadata,omitempty"`
}

func (p *CloudInit) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CloudInit

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

func (p *CloudInit) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CloudInit
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CloudInit(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$cloudInitScriptItemDiscriminator")
	delete(allFields, "cloudInitScript")
	delete(allFields, "datasourceType")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCloudInit() *CloudInit {
	p := new(CloudInit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CloudInit"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloudInit) GetCloudInitScript() interface{} {
	if nil == p.CloudInitScript {
		return nil
	}
	return p.CloudInitScript.GetValue()
}

func (p *CloudInit) SetCloudInitScript(v interface{}) error {
	if nil == p.CloudInitScript {
		p.CloudInitScript = NewOneOfCloudInitCloudInitScript()
	}
	e := p.CloudInitScript.SetValue(v)
	if nil == e {
		if nil == p.CloudInitScriptItemDiscriminator_ {
			p.CloudInitScriptItemDiscriminator_ = new(string)
		}
		*p.CloudInitScriptItemDiscriminator_ = *p.CloudInitScript.Discriminator
	}
	return e
}

/*
Type of datasource. Default: CONFIG_DRIVE_V2
*/
type CloudInitDataSourceType int

const (
	CLOUDINITDATASOURCETYPE_UNKNOWN         CloudInitDataSourceType = 0
	CLOUDINITDATASOURCETYPE_REDACTED        CloudInitDataSourceType = 1
	CLOUDINITDATASOURCETYPE_CONFIG_DRIVE_V2 CloudInitDataSourceType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *CloudInitDataSourceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e CloudInitDataSourceType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *CloudInitDataSourceType) index(name string) CloudInitDataSourceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CONFIG_DRIVE_V2",
	}
	for idx := range names {
		if names[idx] == name {
			return CloudInitDataSourceType(idx)
		}
	}
	return CLOUDINITDATASOURCETYPE_UNKNOWN
}

func (e *CloudInitDataSourceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for CloudInitDataSourceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *CloudInitDataSourceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e CloudInitDataSourceType) Ref() *CloudInitDataSourceType {
	return &e
}

/*
A collection of key/value pairs.
*/
type CustomKeyValues struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the individual KeyValuePair elements.
	*/
	KeyValuePairs []import1.KVPair `json:"keyValuePairs,omitempty"`
}

func (p *CustomKeyValues) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CustomKeyValues

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

func (p *CustomKeyValues) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomKeyValues
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CustomKeyValues(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "keyValuePairs")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCustomKeyValues() *CustomKeyValues {
	p := new(CustomKeyValues)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.CustomKeyValues"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The contents of the user_data configuration for cloud-init. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.
*/
type Userdata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The value for the cloud-init user_data.
	*/
	Value *string `json:"value,omitempty"`
}

func (p *Userdata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Userdata

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

func (p *Userdata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Userdata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Userdata(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewUserdata() *Userdata {
	p := new(Userdata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.config.Userdata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfCloudInitCloudInitScript struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType2001 *Userdata        `json:"-"`
	oneOfType2002 *CustomKeyValues `json:"-"`
}

func NewOneOfCloudInitCloudInitScript() *OneOfCloudInitCloudInitScript {
	p := new(OneOfCloudInitCloudInitScript)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloudInitCloudInitScript) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloudInitCloudInitScript is nil"))
	}
	switch v.(type) {
	case Userdata:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Userdata)
		}
		*p.oneOfType2001 = v.(Userdata)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case CustomKeyValues:
		if nil == p.oneOfType2002 {
			p.oneOfType2002 = new(CustomKeyValues)
		}
		*p.oneOfType2002 = v.(CustomKeyValues)
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

func (p *OneOfCloudInitCloudInitScript) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2002
	}
	return nil
}

func (p *OneOfCloudInitCloudInitScript) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(Userdata)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.config.Userdata" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Userdata)
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
	vOneOfType2002 := new(CustomKeyValues)
	if err := json.Unmarshal(b, vOneOfType2002); err == nil {
		if "vmm.v4.ahv.config.CustomKeyValues" == *vOneOfType2002.ObjectType_ {
			if nil == p.oneOfType2002 {
				p.oneOfType2002 = new(CustomKeyValues)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloudInitCloudInitScript"))
}

func (p *OneOfCloudInitCloudInitScript) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType2002 != nil && *p.oneOfType2002.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2002)
	}
	return nil, errors.New("No value to marshal for OneOfCloudInitCloudInitScript")
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
