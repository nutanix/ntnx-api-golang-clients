/*
 * Generated file models/aiops/v4/stats/stats_model.go.
 *
 * Product version: 4.0.3-alpha-2
 *
 * Part of the Nutanix Aiops Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module aiops.v4.stats of Nutanix Aiops Versioned APIs
*/
package stats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/error"
	import3 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
	import4 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/stats"
	"time"
)

type BoolList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.BoolList"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BoolVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.BoolVal"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An abstract representation of an object which has the deprecated information.
*/
type DeprecationModel struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Represents the deprecation name of the attribute.
	*/
	DeprecatedBy *string `json:"deprecatedBy,omitempty"`
	/*
	  Deprecated after the date.
	*/
	DeprecatedDate *string `json:"deprecatedDate,omitempty"`
	/*
	  Represents the reason for the deprecation.
	*/
	DeprecatedReason *string `json:"deprecatedReason,omitempty"`
	/*
	  Deprecated after the version.
	*/
	DeprecatedVersion *string `json:"deprecatedVersion,omitempty"`
}

func NewDeprecationModel() *DeprecationModel {
	p := new(DeprecationModel)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.DeprecationModel"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.DeprecationModel"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DoubleList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.DoubleList"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DoubleVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.DoubleVal"}
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
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
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
	  Source name for the vendors. For example 'nutanix'.
	*/
	Source *string `json:"source,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntity() *Entity {
	p := new(Entity)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Entity"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.Entity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EntityDescriptor struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Readable name for the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`
	/*
	  Entity type of the data supported for a given source. For example VM, cluster etc.
	*/
	EntityType *string `json:"entityType,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	Metrics []MetricDescriptor `json:"metrics,omitempty"`
	/*
	  Parent entity types for the given entity.
	*/
	Parents []EntityDescriptor `json:"parents,omitempty"`
	/*
	  Source name for the vendors. For example 'nutanix'.
	*/
	Source *string `json:"source,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntityDescriptor() *EntityDescriptor {
	p := new(EntityDescriptor)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.EntityDescriptor"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.EntityDescriptor"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0.a2/stats/sources/{extId}/entity-descriptors Get operation
*/
type EntityDescriptorListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEntityDescriptorListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEntityDescriptorListApiResponse() *EntityDescriptorListApiResponse {
	p := new(EntityDescriptorListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.EntityDescriptorListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.EntityDescriptorListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EntityDescriptorListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EntityDescriptorListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEntityDescriptorListApiResponseData()
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
REST response for all response codes in API path /aiops/v4.0.a2/stats/sources/{sourceExtId}/entities/{extId} Get operation
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.EntityListApiResponse"}
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

type EntityType struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Entity type of the data supported for a given source. For example VM, cluster etc.
	*/
	EntityTypeName *string `json:"entityTypeName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewEntityType() *EntityType {
	p := new(EntityType)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.EntityType"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.EntityType"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0.a2/stats/sources/{extId}/entity-types Get operation
*/
type EntityTypeListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfEntityTypeListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewEntityTypeListApiResponse() *EntityTypeListApiResponse {
	p := new(EntityTypeListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.EntityTypeListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.EntityTypeListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *EntityTypeListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *EntityTypeListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfEntityTypeListApiResponseData()
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.IntList"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type IntVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.IntVal"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Metric struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Name *string `json:"name,omitempty"`

	TimeSeries *TimeSeries `json:"timeSeries,omitempty"`
}

func NewMetric() *Metric {
	p := new(Metric)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Metric"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.Metric"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type MetricDescriptor struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Additional properties for metric descriptor definition. Like anomaly metric, alert metric and search UI metric.
	*/
	AdditionalProperties []import3.KVPair `json:"additionalProperties,omitempty"`
	/*
	  Default value of the metric.
	*/
	DefaultValue *string `json:"defaultValue,omitempty"`

	Deprecation *DeprecationModel `json:"deprecation,omitempty"`
	/*
	  Readable name for the entity.
	*/
	DisplayName *string `json:"displayName,omitempty"`

	DownsamplingOperator *import4.DownSamplingOperator `json:"downsamplingOperator,omitempty"`
	/*
	  Indicates whether it is an attribute or not.
	*/
	IsAttribute *bool `json:"isAttribute,omitempty"`
	/*
	  Indicator to specify whether the attribute should be persisted as time series data or not.
	*/
	IsAttributePersistedAsTimeSeries *bool `json:"isAttributePersistedAsTimeSeries,omitempty"`
	/*
	  Name of the metric.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  The interval value is used to resample the queried time series data by using down_sampling_operator operator. The default is 86400 seconds.
	*/
	SamplingIntervalSecs *int `json:"samplingIntervalSecs,omitempty"`
	/*
	  Unit for the metric.
	*/
	Unit *string `json:"unit,omitempty"`

	ValueRange *ValueRange `json:"valueRange,omitempty"`

	ValueType *ValueType `json:"valueType,omitempty"`
}

func NewMetricDescriptor() *MetricDescriptor {
	p := new(MetricDescriptor)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.MetricDescriptor"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.MetricDescriptor"}
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.Point"}
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

type Source struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Source name for the vendors. For example 'nutanix'.
	*/
	SourceName *string `json:"sourceName,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSource() *Source {
	p := new(Source)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.Source"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.Source"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /aiops/v4.0.a2/stats/sources Get operation
*/
type SourceListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSourceListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewSourceListApiResponse() *SourceListApiResponse {
	p := new(SourceListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.SourceListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.SourceListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SourceListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SourceListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSourceListApiResponseData()
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

type StrList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.StrList"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StrVal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.StrVal"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TimeSeries struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The interval value is used to resample the queried time series data by using down_sampling_operator operator. The default is 86400 seconds.
	*/
	SamplingIntervalSecs *int `json:"samplingIntervalSecs,omitempty"`

	Values []Point `json:"values,omitempty"`
}

func NewTimeSeries() *TimeSeries {
	p := new(TimeSeries)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.TimeSeries"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.TimeSeries"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Range of the values for the metric.
*/
type ValueRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Max *float64 `json:"max,omitempty"`

	Min *float64 `json:"min,omitempty"`
}

func NewValueRange() *ValueRange {
	p := new(ValueRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.stats.ValueRange"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a2.stats.ValueRange"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Data type of the values for the metric.
*/
type ValueType int

const (
	VALUETYPE_UNKNOWN     ValueType = 0
	VALUETYPE_REDACTED    ValueType = 1
	VALUETYPE_BOOL        ValueType = 2
	VALUETYPE_INT         ValueType = 3
	VALUETYPE_DOUBLE      ValueType = 4
	VALUETYPE_STRING      ValueType = 5
	VALUETYPE_BOOL_LIST   ValueType = 6
	VALUETYPE_INT_LIST    ValueType = 7
	VALUETYPE_DOUBLE_LIST ValueType = 8
	VALUETYPE_STRING_LIST ValueType = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ValueType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ValueType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ValueType) index(name string) ValueType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"BOOL",
		"INT",
		"DOUBLE",
		"STRING",
		"BOOL_LIST",
		"INT_LIST",
		"DOUBLE_LIST",
		"STRING_LIST",
	}
	for idx := range names {
		if names[idx] == name {
			return ValueType(idx)
		}
	}
	return VALUETYPE_UNKNOWN
}

func (e *ValueType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ValueType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ValueType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ValueType) Ref() *ValueType {
	return &e
}

type OneOfEntityListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Entity               `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfEntityListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<aiops.v4.stats.Entity>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfEntityListApiResponseData) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityListApiResponseData"))
}

func (p *OneOfEntityListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<aiops.v4.stats.Entity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfEntityListApiResponseData")
}

type OneOfEntityDescriptorListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []EntityDescriptor     `json:"-"`
}

func NewOneOfEntityDescriptorListApiResponseData() *OneOfEntityDescriptorListApiResponseData {
	p := new(OneOfEntityDescriptorListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEntityDescriptorListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEntityDescriptorListApiResponseData is nil"))
	}
	switch v.(type) {
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
	case []EntityDescriptor:
		p.oneOfType0 = v.([]EntityDescriptor)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.stats.EntityDescriptor>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.stats.EntityDescriptor>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfEntityDescriptorListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<aiops.v4.stats.EntityDescriptor>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfEntityDescriptorListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]EntityDescriptor)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.stats.EntityDescriptor" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.stats.EntityDescriptor>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.stats.EntityDescriptor>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityDescriptorListApiResponseData"))
}

func (p *OneOfEntityDescriptorListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<aiops.v4.stats.EntityDescriptor>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfEntityDescriptorListApiResponseData")
}

type OneOfPointValue struct {
	Discriminator *string     `json:"-"`
	ObjectType_   *string     `json:"-"`
	oneOfType6    *IntList    `json:"-"`
	oneOfType1    *BoolVal    `json:"-"`
	oneOfType5    *BoolList   `json:"-"`
	oneOfType7    *DoubleList `json:"-"`
	oneOfType3    *DoubleVal  `json:"-"`
	oneOfType0    *StrVal     `json:"-"`
	oneOfType2    *IntVal     `json:"-"`
	oneOfType4    *StrList    `json:"-"`
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPointValue) GetValue() interface{} {
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return *p.oneOfType6
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return *p.oneOfType5
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return *p.oneOfType7
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return *p.oneOfType3
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return *p.oneOfType4
	}
	return nil
}

func (p *OneOfPointValue) UnmarshalJSON(b []byte) error {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPointValue"))
}

func (p *OneOfPointValue) MarshalJSON() ([]byte, error) {
	if p.oneOfType6 != nil && *p.oneOfType6.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType6)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType5 != nil && *p.oneOfType5.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType5)
	}
	if p.oneOfType7 != nil && *p.oneOfType7.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType7)
	}
	if p.oneOfType3 != nil && *p.oneOfType3.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType3)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	if p.oneOfType4 != nil && *p.oneOfType4.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType4)
	}
	return nil, errors.New("No value to marshal for OneOfPointValue")
}

type OneOfEntityTypeListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []EntityType           `json:"-"`
}

func NewOneOfEntityTypeListApiResponseData() *OneOfEntityTypeListApiResponseData {
	p := new(OneOfEntityTypeListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfEntityTypeListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfEntityTypeListApiResponseData is nil"))
	}
	switch v.(type) {
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
	case []EntityType:
		p.oneOfType0 = v.([]EntityType)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.stats.EntityType>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.stats.EntityType>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfEntityTypeListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<aiops.v4.stats.EntityType>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfEntityTypeListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]EntityType)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.stats.EntityType" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.stats.EntityType>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.stats.EntityType>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfEntityTypeListApiResponseData"))
}

func (p *OneOfEntityTypeListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<aiops.v4.stats.EntityType>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfEntityTypeListApiResponseData")
}

type OneOfSourceListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Source               `json:"-"`
}

func NewOneOfSourceListApiResponseData() *OneOfSourceListApiResponseData {
	p := new(OneOfSourceListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSourceListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSourceListApiResponseData is nil"))
	}
	switch v.(type) {
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
	case []Source:
		p.oneOfType0 = v.([]Source)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<aiops.v4.stats.Source>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<aiops.v4.stats.Source>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSourceListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<aiops.v4.stats.Source>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfSourceListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]Source)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "aiops.v4.stats.Source" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<aiops.v4.stats.Source>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<aiops.v4.stats.Source>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSourceListApiResponseData"))
}

func (p *OneOfSourceListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<aiops.v4.stats.Source>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfSourceListApiResponseData")
}
