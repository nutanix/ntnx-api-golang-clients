/*
 * Generated file models/aiops/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
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
	return json.Marshal(struct {
		*BoolListProxy
		BoolList []bool `json:"boolList,omitempty"`
	}{
		BoolListProxy: (*BoolListProxy)(p),
		BoolList:      p.BoolList,
	})
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
	return json.Marshal(struct {
		*BoolValProxy
		BoolValue *bool `json:"boolValue,omitempty"`
	}{
		BoolValProxy: (*BoolValProxy)(p),
		BoolValue:    p.BoolValue,
	})
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
	return json.Marshal(struct {
		*DoubleListProxy
		DoubleList []float64 `json:"doubleList,omitempty"`
	}{
		DoubleListProxy: (*DoubleListProxy)(p),
		DoubleList:      p.DoubleList,
	})
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
	return json.Marshal(struct {
		*DoubleValProxy
		DoubleValue *float64 `json:"doubleValue,omitempty"`
	}{
		DoubleValProxy: (*DoubleValProxy)(p),
		DoubleValue:    p.DoubleValue,
	})
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
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
	return json.Marshal(struct {
		*IntListProxy
		IntList []int64 `json:"intList,omitempty"`
	}{
		IntListProxy: (*IntListProxy)(p),
		IntList:      p.IntList,
	})
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
	return json.Marshal(struct {
		*IntValProxy
		IntValue *int64 `json:"intValue,omitempty"`
	}{
		IntValProxy: (*IntValProxy)(p),
		IntValue:    p.IntValue,
	})
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
	return json.Marshal(struct {
		*StrListProxy
		StrList []string `json:"strList,omitempty"`
	}{
		StrListProxy: (*StrListProxy)(p),
		StrList:      p.StrList,
	})
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
	return json.Marshal(struct {
		*StrValProxy
		StrValue *string `json:"strValue,omitempty"`
	}{
		StrValProxy: (*StrValProxy)(p),
		StrValue:    p.StrValue,
	})
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
