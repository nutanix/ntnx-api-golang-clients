/*
 * Generated file models/aiops/v4/stats/stats_model.go.
 *
 * Product version: 4.0.2
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module aiops.v4.stats of Nutanix AIOps APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/error"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
	"time"
)

type BoolList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is a list of booleans.
	*/
	BoolList []bool `json:"boolList"`
}

func (p *BoolList) MarshalJSON() ([]byte, error) {
	type BoolListProxy BoolList

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BoolListProxy
		BoolList []bool `json:"boolList,omitempty"`
	}{
		BoolListProxy: (*BoolListProxy)(p),
		BoolList:      p.BoolList,
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

func (p *BoolList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BoolList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BoolList(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "boolList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBoolList() *BoolList {
	p := new(BoolList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.BoolList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BoolVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is boolean.
	*/
	BoolValue *bool `json:"boolValue"`
}

func (p *BoolVal) MarshalJSON() ([]byte, error) {
	type BoolValProxy BoolVal

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BoolValProxy
		BoolValue *bool `json:"boolValue,omitempty"`
	}{
		BoolValProxy: (*BoolValProxy)(p),
		BoolValue:    p.BoolValue,
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

func (p *BoolVal) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BoolVal
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BoolVal(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "boolValue")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBoolVal() *BoolVal {
	p := new(BoolVal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.BoolVal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DoubleList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is a list of doubles.
	*/
	DoubleList []float64 `json:"doubleList"`
}

func (p *DoubleList) MarshalJSON() ([]byte, error) {
	type DoubleListProxy DoubleList

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DoubleListProxy
		DoubleList []float64 `json:"doubleList,omitempty"`
	}{
		DoubleListProxy: (*DoubleListProxy)(p),
		DoubleList:      p.DoubleList,
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

func (p *DoubleList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DoubleList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DoubleList(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "doubleList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDoubleList() *DoubleList {
	p := new(DoubleList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.DoubleList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DoubleVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is double.
	*/
	DoubleValue *float64 `json:"doubleValue"`
}

func (p *DoubleVal) MarshalJSON() ([]byte, error) {
	type DoubleValProxy DoubleVal

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DoubleValProxy
		DoubleValue *float64 `json:"doubleValue,omitempty"`
	}{
		DoubleValProxy: (*DoubleValProxy)(p),
		DoubleValue:    p.DoubleValue,
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

func (p *DoubleVal) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DoubleVal
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DoubleVal(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "doubleValue")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDoubleVal() *DoubleVal {
	p := new(DoubleVal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.DoubleVal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Entity struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity type of the data supported for a given source. For example VM, cluster etc.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Metrics data for the entity.
	*/
	Metrics []Metric `json:"metrics,omitempty"`
	/*
	  Parent entity types for the given entity.
	*/
	Parents []Entity `json:"parents,omitempty"`
	/*
	  Source name for the vendors. For example 'Nutanix'.
	*/
	Source *string `json:"source,omitempty"`
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
	delete(allFields, "entityType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "metrics")
	delete(allFields, "parents")
	delete(allFields, "source")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Entity"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0/stats/sources/{sourceExtId}/entities/{extId} Get operation
*/
type EntityListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEntityListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *EntityListApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityListApiResponse

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

func (p *EntityListApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityListApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EntityListApiResponse(*known)

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

func NewEntityListApiResponse() *EntityListApiResponse {
	p := new(EntityListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.EntityListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EntityListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EntityListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEntityListApiResponseData()
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

type IntList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is a list of integers.
	*/
	IntList []int64 `json:"intList"`
}

func (p *IntList) MarshalJSON() ([]byte, error) {
	type IntListProxy IntList

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IntListProxy
		IntList []int64 `json:"intList,omitempty"`
	}{
		IntListProxy: (*IntListProxy)(p),
		IntList:      p.IntList,
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

func (p *IntList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IntList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IntList(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "intList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIntList() *IntList {
	p := new(IntList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.IntList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type IntVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is integer.
	*/
	IntValue *int64 `json:"intValue"`
}

func (p *IntVal) MarshalJSON() ([]byte, error) {
	type IntValProxy IntVal

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IntValProxy
		IntValue *int64 `json:"intValue,omitempty"`
	}{
		IntValProxy: (*IntValProxy)(p),
		IntValue:    p.IntValue,
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

func (p *IntVal) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IntVal
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IntVal(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "intValue")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIntVal() *IntVal {
	p := new(IntVal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.IntVal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Metric struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the metric.
	*/
	Name *string `json:"name,omitempty"`

	TimeSeries *TimeSeries `json:"timeSeries,omitempty"`
}

func (p *Metric) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Metric

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

func (p *Metric) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Metric
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Metric(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "timeSeries")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewMetric() *Metric {
	p := new(Metric)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Metric"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Point struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The timestamp of the metric value.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`

	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`
	/*
	  The data point value at the specified timestamp, this value can either be an integer, boolean, string, double, or a list of these types.
	*/
	Value *OneOfPointValue `json:"value,omitempty"`
}

func (p *Point) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Point

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

func (p *Point) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Point
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Point(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPoint() *Point {
	p := new(Point)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Point"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Point) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *Point) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfPointValue()
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
A collection of WhatIf Scenario stats.
*/
type ScenarioStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID of the WhatIf Scenario associated with the Scneario Stats.
	*/
	ScenarioExtId *string `json:"scenarioExtId,omitempty"`
	/*
	  WhatIf Scenario entity statistic time-series.
	*/
	Stats []ScenarioStatsTuple `json:"stats,omitempty"`
}

func (p *ScenarioStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ScenarioStats

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

func (p *ScenarioStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ScenarioStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ScenarioStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "scenarioExtId")
	delete(allFields, "stats")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewScenarioStats() *ScenarioStats {
	p := new(ScenarioStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.ScenarioStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ScenarioStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID of the WhatIf Scenario associated with the Scneario Stats.
	*/
	ScenarioExtId *string `json:"scenarioExtId,omitempty"`
	/*
	  WhatIf Scenario entity statistic time-series.
	*/
	Stats []ScenarioStatsTuple `json:"stats,omitempty"`
}

func (p *ScenarioStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ScenarioStatsProjection

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

func (p *ScenarioStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ScenarioStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ScenarioStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "scenarioExtId")
	delete(allFields, "stats")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewScenarioStatsProjection() *ScenarioStatsProjection {
	p := new(ScenarioStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.ScenarioStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Tuple value of timestamp and CPU/Memory/Storage resource usage, capacity, and effective capacity.
*/
type ScenarioStatsTuple struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  CPU capacity in Ghz.
	*/
	CpuCapacityGhz *float64 `json:"cpuCapacityGhz,omitempty"`
	/*
	  CPU effective capacity in Ghz.
	*/
	CpuEffectiveCapacityGhz *float64 `json:"cpuEffectiveCapacityGhz,omitempty"`
	/*
	  Usage of CPU resources in Ghz.
	*/
	CpuUsageGhz *float64 `json:"cpuUsageGhz,omitempty"`
	/*
	  Memory capacity in Gb.
	*/
	MemoryCapacityGb *float64 `json:"memoryCapacityGb,omitempty"`
	/*
	  Memory effective capacity in Gb.
	*/
	MemoryEffectiveCapacityGb *float64 `json:"memoryEffectiveCapacityGb,omitempty"`
	/*
	  Usage of Memory resources in Gb.
	*/
	MemoryUsageGb *float64 `json:"memoryUsageGb,omitempty"`
	/*
	  Storage capacity in Gb.
	*/
	StorageCapacityGb *float64 `json:"storageCapacityGb,omitempty"`
	/*
	  Storage effective capacity in Gb.
	*/
	StorageEffectiveCapacityGb *float64 `json:"storageEffectiveCapacityGb,omitempty"`
	/*
	  Usage of storage resources in Gb.
	*/
	StorageUsageGb *float64 `json:"storageUsageGb,omitempty"`
	/*
	  Timestamp of a WhatIf Scenario statistic attribute data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *ScenarioStatsTuple) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ScenarioStatsTuple

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

func (p *ScenarioStatsTuple) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ScenarioStatsTuple
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ScenarioStatsTuple(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpuCapacityGhz")
	delete(allFields, "cpuEffectiveCapacityGhz")
	delete(allFields, "cpuUsageGhz")
	delete(allFields, "memoryCapacityGb")
	delete(allFields, "memoryEffectiveCapacityGb")
	delete(allFields, "memoryUsageGb")
	delete(allFields, "storageCapacityGb")
	delete(allFields, "storageEffectiveCapacityGb")
	delete(allFields, "storageUsageGb")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewScenarioStatsTuple() *ScenarioStatsTuple {
	p := new(ScenarioStatsTuple)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.ScenarioStatsTuple"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StrList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is a list of strings.
	*/
	StrList []string `json:"strList"`
}

func (p *StrList) MarshalJSON() ([]byte, error) {
	type StrListProxy StrList

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StrListProxy
		StrList []string `json:"strList,omitempty"`
	}{
		StrListProxy: (*StrListProxy)(p),
		StrList:      p.StrList,
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

func (p *StrList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StrList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StrList(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "strList")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStrList() *StrList {
	p := new(StrList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.StrList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StrVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Value is string.
	*/
	StrValue *string `json:"strValue"`
}

func (p *StrVal) MarshalJSON() ([]byte, error) {
	type StrValProxy StrVal

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StrValProxy
		StrValue *string `json:"strValue,omitempty"`
	}{
		StrValProxy: (*StrValProxy)(p),
		StrValue:    p.StrValue,
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

func (p *StrVal) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StrVal
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StrVal(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "strValue")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStrVal() *StrVal {
	p := new(StrVal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.StrVal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of points for the time series data.
*/
type TimeSeries struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The interval value is used to resample the queried time series data by using down_sampling_operator operator. The default is 86400 seconds.
	*/
	SamplingIntervalSecs *int `json:"samplingIntervalSecs,omitempty"`
	/*
	  The data point value at the specified timestamp, this value can either be an integer, boolean, string, double, or a list of these types.
	*/
	Values []Point `json:"values,omitempty"`
}

func (p *TimeSeries) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TimeSeries

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

func (p *TimeSeries) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TimeSeries
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TimeSeries(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "samplingIntervalSecs")
	delete(allFields, "values")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTimeSeries() *TimeSeries {
	p := new(TimeSeries)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.TimeSeries"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfEntityListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []Entity               `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfEntityListApiResponseData() *OneOfEntityListApiResponseData {
	p := new(OneOfEntityListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEntityListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEntityListApiResponseData is nil"))
	}
	switch v.(type) {
	case []Entity:
		p.oneOfType0 = v.([]Entity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.stats.Entity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.stats.Entity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfEntityListApiResponseData) GetValue() interface{} {
	if "List<aiops.v4.stats.Entity>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfEntityListApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]Entity)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.stats.Entity" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.stats.Entity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.stats.Entity>"
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityListApiResponseData"))
}

func (p *OneOfEntityListApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<aiops.v4.stats.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfEntityListApiResponseData")
}

type OneOfPointValue struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType5    *BoolList   `json:"-"`
	oneOfType6    *IntList    `json:"-"`
	oneOfType3    *DoubleVal  `json:"-"`
	oneOfType1    *BoolVal    `json:"-"`
	oneOfType4    *StrList    `json:"-"`
	oneOfType0    *StrVal     `json:"-"`
	oneOfType7    *DoubleList `json:"-"`
	oneOfType2    *IntVal     `json:"-"`
}

func NewOneOfPointValue() *OneOfPointValue {
	p := new(OneOfPointValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPointValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPointValue is nil"))
	}
	switch v.(type) {
	case BoolList:
		if nil == p.oneOfType5 {
			p.oneOfType5 = new(BoolList)
		}
		*p.oneOfType5 = v.(BoolList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType5.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType5.ObjectType_
	case IntList:
		if nil == p.oneOfType6 {
			p.oneOfType6 = new(IntList)
		}
		*p.oneOfType6 = v.(IntList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType6.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType6.ObjectType_
	case DoubleVal:
		if nil == p.oneOfType3 {
			p.oneOfType3 = new(DoubleVal)
		}
		*p.oneOfType3 = v.(DoubleVal)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType3.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType3.ObjectType_
	case BoolVal:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(BoolVal)
		}
		*p.oneOfType1 = v.(BoolVal)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case StrList:
		if nil == p.oneOfType4 {
			p.oneOfType4 = new(StrList)
		}
		*p.oneOfType4 = v.(StrList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType4.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType4.ObjectType_
	case StrVal:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(StrVal)
		}
		*p.oneOfType0 = v.(StrVal)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case DoubleList:
		if nil == p.oneOfType7 {
			p.oneOfType7 = new(DoubleList)
		}
		*p.oneOfType7 = v.(DoubleList)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType7.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType7.ObjectType_
	case IntVal:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(IntVal)
		}
		*p.oneOfType2 = v.(IntVal)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPointValue) GetValue() interface{} {
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5
	}
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return *p.oneOfType6
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return *p.oneOfType7
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	return nil
}

func (p *OneOfPointValue) UnmarshalJSON(b []byte) error {
	vOneOfType5 := new(BoolList)
	if err := json.Unmarshal(b, vOneOfType5); err == nil {
		if "aiops.v4.stats.BoolList" == *vOneOfType5.ObjectType_ {
			if nil == p.oneOfType5 {
				p.oneOfType5 = new(BoolList)
			}
			*p.oneOfType5 = *vOneOfType5
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType5.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType5.ObjectType_
			return nil
		}
	}
	vOneOfType6 := new(IntList)
	if err := json.Unmarshal(b, vOneOfType6); err == nil {
		if "aiops.v4.stats.IntList" == *vOneOfType6.ObjectType_ {
			if nil == p.oneOfType6 {
				p.oneOfType6 = new(IntList)
			}
			*p.oneOfType6 = *vOneOfType6
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType6.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType6.ObjectType_
			return nil
		}
	}
	vOneOfType3 := new(DoubleVal)
	if err := json.Unmarshal(b, vOneOfType3); err == nil {
		if "aiops.v4.stats.DoubleVal" == *vOneOfType3.ObjectType_ {
			if nil == p.oneOfType3 {
				p.oneOfType3 = new(DoubleVal)
			}
			*p.oneOfType3 = *vOneOfType3
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType3.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType3.ObjectType_
			return nil
		}
	}
	vOneOfType1 := new(BoolVal)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "aiops.v4.stats.BoolVal" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(BoolVal)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	vOneOfType4 := new(StrList)
	if err := json.Unmarshal(b, vOneOfType4); err == nil {
		if "aiops.v4.stats.StrList" == *vOneOfType4.ObjectType_ {
			if nil == p.oneOfType4 {
				p.oneOfType4 = new(StrList)
			}
			*p.oneOfType4 = *vOneOfType4
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType4.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType4.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(StrVal)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "aiops.v4.stats.StrVal" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(StrVal)
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
	vOneOfType7 := new(DoubleList)
	if err := json.Unmarshal(b, vOneOfType7); err == nil {
		if "aiops.v4.stats.DoubleList" == *vOneOfType7.ObjectType_ {
			if nil == p.oneOfType7 {
				p.oneOfType7 = new(DoubleList)
			}
			*p.oneOfType7 = *vOneOfType7
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType7.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType7.ObjectType_
			return nil
		}
	}
	vOneOfType2 := new(IntVal)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "aiops.v4.stats.IntVal" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(IntVal)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPointValue"))
}

func (p *OneOfPointValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5)
	}
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType6)
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType7)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	return nil, errors.New("No value to marshal for OneOfPointValue")
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
