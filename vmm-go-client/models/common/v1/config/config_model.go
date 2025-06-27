/*
 * Generated file models/common/v1/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Nutanix Standard Configuration
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

/*
Many entities in the Nutanix APIs carry flags.  This object captures all the flags associated with that entity through this object.  The field that hosts this type of object must have an attribute called x-bounded-map-keys that tells which flags are actually present for that entity.
*/
type Flag struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the flag.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Value of the flag.
	*/
	Value *bool `json:"value,omitempty"`
}

func (p *Flag) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Flag

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

func (p *Flag) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Flag
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Flag(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewFlag() *Flag {
	p := new(Flag)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.Flag"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Value = new(bool)
	*p.Value = false

	return p
}

/*
An unique address that identifies a device on the internet or a local network in IPv4 format.
*/
type IPv4Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The prefix length of the network to which this host IPv4 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  The IPv4 address of the host.
	*/
	Value *string `json:"value"`
}

func (p *IPv4Address) MarshalJSON() ([]byte, error) {
	type IPv4AddressProxy IPv4Address

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IPv4AddressProxy
		Value *string `json:"value,omitempty"`
	}{
		IPv4AddressProxy: (*IPv4AddressProxy)(p),
		Value:            p.Value,
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

func (p *IPv4Address) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPv4Address
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IPv4Address(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "prefixLength")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIPv4Address() *IPv4Address {
	p := new(IPv4Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.IPv4Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 32

	return p
}

/*
An unique address that identifies a device on the internet or a local network in IPv6 format.
*/
type IPv6Address struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The prefix length of the network to which this host IPv6 address belongs.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`
	/*
	  The IPv6 address of the host.
	*/
	Value *string `json:"value"`
}

func (p *IPv6Address) MarshalJSON() ([]byte, error) {
	type IPv6AddressProxy IPv6Address

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IPv6AddressProxy
		Value *string `json:"value,omitempty"`
	}{
		IPv6AddressProxy: (*IPv6AddressProxy)(p),
		Value:            p.Value,
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

func (p *IPv6Address) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IPv6Address
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IPv6Address(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "prefixLength")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIPv6Address() *IPv6Address {
	p := new(IPv6Address)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.IPv6Address"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.PrefixLength = new(int)
	*p.PrefixLength = 128

	return p
}

/*
A map describing a set of keys and their corresponding values.
*/
type KVPair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The key of the key-value pair.
	*/
	Name *string `json:"name,omitempty"`
	/*

	 */
	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  The value associated with the key for this key-value pair
	*/
	Value *OneOfKVPairValue `json:"value,omitempty"`
}

func (p *KVPair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias KVPair

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

func (p *KVPair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias KVPair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = KVPair(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewKVPair() *KVPair {
	p := new(KVPair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.KVPair"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *KVPair) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *KVPair) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfKVPairValue()
	}
	e := p.Value.SetValue(v)
	if nil == e {
		if nil == p.ValueItemDiscriminator_ {
			p.ValueItemDiscriminator_ = new(string)
		}
		*p.ValueItemDiscriminator_ = *p.Value.Discriminator
	}
	return e
}

/*
A wrapper schema containing a map with string keys and values.
*/
type MapOfStringWrapper struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A map with string keys and values.
	*/
	Map map[string]string `json:"map,omitempty"`
}

func (p *MapOfStringWrapper) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias MapOfStringWrapper

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

func (p *MapOfStringWrapper) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias MapOfStringWrapper
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = MapOfStringWrapper(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "map")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMapOfStringWrapper() *MapOfStringWrapper {
	p := new(MapOfStringWrapper)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.MapOfStringWrapper"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Message struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A code that uniquely identifies a message.
	*/
	Code *string `json:"code,omitempty"`
	/*
	  The locale for the message description.
	*/
	Locale *string `json:"locale,omitempty"`
	/*
	  The description of the message.
	*/
	Message *string `json:"message,omitempty"`

	Severity *MessageSeverity `json:"severity,omitempty"`
}

func (p *Message) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Message

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

func (p *Message) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Message
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Message(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "code")
	delete(allFields, "locale")
	delete(allFields, "message")
	delete(allFields, "severity")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMessage() *Message {
	p := new(Message)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.Message"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Locale = new(string)
	*p.Locale = "en_US"

	return p
}

/*
The message severity.
*/
type MessageSeverity int

const (
	MESSAGESEVERITY_UNKNOWN  MessageSeverity = 0
	MESSAGESEVERITY_REDACTED MessageSeverity = 1
	MESSAGESEVERITY_INFO     MessageSeverity = 2
	MESSAGESEVERITY_WARNING  MessageSeverity = 3
	MESSAGESEVERITY_ERROR    MessageSeverity = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *MessageSeverity) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e MessageSeverity) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *MessageSeverity) index(name string) MessageSeverity {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INFO",
		"WARNING",
		"ERROR",
	}
	for idx := range names {
		if names[idx] == name {
			return MessageSeverity(idx)
		}
	}
	return MESSAGESEVERITY_UNKNOWN
}

func (e *MessageSeverity) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MessageSeverity:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MessageSeverity) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MessageSeverity) Ref() *MessageSeverity {
	return &e
}

/*
A model base class whose instances are bound to a specific tenant.  This model adds a tenantId to the base model class that it extends and is automatically set by the server.
*/
type TenantAwareModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *TenantAwareModel) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TenantAwareModel

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

func (p *TenantAwareModel) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TenantAwareModel
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TenantAwareModel(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTenantAwareModel() *TenantAwareModel {
	p := new(TenantAwareModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "common.v1.config.TenantAwareModel"
	p.Reserved_ = map[string]interface{}{"$fv": "v1.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfKVPairValue struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType1006 map[string]string    `json:"-"`
	oneOfType1004 *bool                `json:"-"`
	oneOfType1005 []string             `json:"-"`
	oneOfType1003 *int                 `json:"-"`
	oneOfType1008 []int                `json:"-"`
	oneOfType1002 *string              `json:"-"`
	oneOfType1007 []MapOfStringWrapper `json:"-"`
}

func NewOneOfKVPairValue() *OneOfKVPairValue {
	p := new(OneOfKVPairValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfKVPairValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfKVPairValue is nil"))
	}
	switch v.(type) {
	case map[string]string:
		p.oneOfType1006 = v.(map[string]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Map<String, String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Map<String, String>"
	case bool:
		if nil == p.oneOfType1004 {
			p.oneOfType1004 = new(bool)
		}
		*p.oneOfType1004 = v.(bool)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
	case []string:
		p.oneOfType1005 = v.([]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
	case int:
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(int)
		}
		*p.oneOfType1003 = v.(int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
	case []int:
		p.oneOfType1008 = v.([]int)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<Integer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<Integer>"
	case string:
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(string)
		}
		*p.oneOfType1002 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	case []MapOfStringWrapper:
		p.oneOfType1007 = v.([]MapOfStringWrapper)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<common.v1.config.MapOfStringWrapper>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<common.v1.config.MapOfStringWrapper>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfKVPairValue) GetValue() interface{} {
	if "Map<String, String>" == *p.Discriminator {
		return p.oneOfType1006
	}
	if "Boolean" == *p.Discriminator {
		return *p.oneOfType1004
	}
	if "List<String>" == *p.Discriminator {
		return p.oneOfType1005
	}
	if "Integer" == *p.Discriminator {
		return *p.oneOfType1003
	}
	if "List<Integer>" == *p.Discriminator {
		return p.oneOfType1008
	}
	if "String" == *p.Discriminator {
		return *p.oneOfType1002
	}
	if "List<common.v1.config.MapOfStringWrapper>" == *p.Discriminator {
		return p.oneOfType1007
	}
	return nil
}

func (p *OneOfKVPairValue) UnmarshalJSON(b []byte) error {
	vOneOfType1006 := new(map[string]string)
	if err := json.Unmarshal(b, vOneOfType1006); err == nil {
		p.oneOfType1006 = *vOneOfType1006
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Map<String, String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Map<String, String>"
		return nil
	}
	vOneOfType1004 := new(bool)
	if err := json.Unmarshal(b, vOneOfType1004); err == nil {
		if nil == p.oneOfType1004 {
			p.oneOfType1004 = new(bool)
		}
		*p.oneOfType1004 = *vOneOfType1004
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Boolean"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Boolean"
		return nil
	}
	vOneOfType1005 := new([]string)
	if err := json.Unmarshal(b, vOneOfType1005); err == nil {
		p.oneOfType1005 = *vOneOfType1005
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
		return nil
	}
	vOneOfType1003 := new(int)
	if err := json.Unmarshal(b, vOneOfType1003); err == nil {
		if nil == p.oneOfType1003 {
			p.oneOfType1003 = new(int)
		}
		*p.oneOfType1003 = *vOneOfType1003
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "Integer"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "Integer"
		return nil
	}
	vOneOfType1008 := new([]int)
	if err := json.Unmarshal(b, vOneOfType1008); err == nil {
		p.oneOfType1008 = *vOneOfType1008
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<Integer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<Integer>"
		return nil
	}
	vOneOfType1002 := new(string)
	if err := json.Unmarshal(b, vOneOfType1002); err == nil {
		if nil == p.oneOfType1002 {
			p.oneOfType1002 = new(string)
		}
		*p.oneOfType1002 = *vOneOfType1002
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
		return nil
	}
	vOneOfType1007 := new([]MapOfStringWrapper)
	if err := json.Unmarshal(b, vOneOfType1007); err == nil {
		if len(*vOneOfType1007) == 0 || "common.v1.config.MapOfStringWrapper" == *((*vOneOfType1007)[0].ObjectType_) {
			p.oneOfType1007 = *vOneOfType1007
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<common.v1.config.MapOfStringWrapper>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<common.v1.config.MapOfStringWrapper>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfKVPairValue"))
}

func (p *OneOfKVPairValue) MarshalJSON() ([]byte, error) {
	if "Map<String, String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1006)
	}
	if "Boolean" == *p.Discriminator {
		return json.Marshal(p.oneOfType1004)
	}
	if "List<String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1005)
	}
	if "Integer" == *p.Discriminator {
		return json.Marshal(p.oneOfType1003)
	}
	if "List<Integer>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1008)
	}
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType1002)
	}
	if "List<common.v1.config.MapOfStringWrapper>" == *p.Discriminator {
		return json.Marshal(p.oneOfType1007)
	}
	return nil, errors.New("No value to marshal for OneOfKVPairValue")
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
