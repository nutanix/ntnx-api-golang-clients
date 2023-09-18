/*
 * Generated file models/networking/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Networking Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Fetch stats for VPN connections, VPCs and traffic mirrors
*/
package stats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
)

/*
Layer2 stretch statistics description
*/
type Layer2StretchStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Layer2Stretch string array of round-trip-time.
	*/
	Rtt []string `json:"rtt,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPN connection string array of RX BPS values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/*
	  VPN connection string array of TX BPS values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewLayer2StretchStats() *Layer2StretchStats {
	p := new(Layer2StretchStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.Layer2StretchStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.Layer2StretchStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0.b1/stats/layer2-stretches/{extId} Get operation
*/
type Layer2StretchStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfLayer2StretchStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLayer2StretchStatsApiResponse() *Layer2StretchStatsApiResponse {
	p := new(Layer2StretchStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.Layer2StretchStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.Layer2StretchStatsApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Layer2StretchStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *Layer2StretchStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfLayer2StretchStatsApiResponseData()
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
VPC UUID to reset all routing policy counters to zero.
*/
type RoutingPolicyClearCountersSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RoutingPolicyExtId *string `json:"routingPolicyExtId,omitempty"`

	VpcExtId *string `json:"vpcExtId"`
}

func (p *RoutingPolicyClearCountersSpec) MarshalJSON() ([]byte, error) {
	type RoutingPolicyClearCountersSpecProxy RoutingPolicyClearCountersSpec
	return json.Marshal(struct {
		*RoutingPolicyClearCountersSpecProxy
		VpcExtId *string `json:"vpcExtId,omitempty"`
	}{
		RoutingPolicyClearCountersSpecProxy: (*RoutingPolicyClearCountersSpecProxy)(p),
		VpcExtId:                            p.VpcExtId,
	})
}

func NewRoutingPolicyClearCountersSpec() *RoutingPolicyClearCountersSpec {
	p := new(RoutingPolicyClearCountersSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.RoutingPolicyClearCountersSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.RoutingPolicyClearCountersSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Response of statistics query.
*/
type StatsQueryResponseBase struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewStatsQueryResponseBase() *StatsQueryResponseBase {
	p := new(StatsQueryResponseBase)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.StatsQueryResponseBase"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.StatsQueryResponseBase"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0.b1/stats/routing-policies/$actions/clear Post operation
*/
type TaskReferenceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTaskReferenceApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskReferenceApiResponse() *TaskReferenceApiResponse {
	p := new(TaskReferenceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TaskReferenceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.TaskReferenceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TaskReferenceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TaskReferenceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTaskReferenceApiResponseData()
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
Traffic mirror state value.
*/
type TrafficMirrorState int

const (
	TRAFFICMIRRORSTATE_UNKNOWN  TrafficMirrorState = 0
	TRAFFICMIRRORSTATE_REDACTED TrafficMirrorState = 1
	TRAFFICMIRRORSTATE_ACTIVE   TrafficMirrorState = 2
	TRAFFICMIRRORSTATE_ERROR    TrafficMirrorState = 3
	TRAFFICMIRRORSTATE_DISABLED TrafficMirrorState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TrafficMirrorState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TrafficMirrorState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TrafficMirrorState) index(name string) TrafficMirrorState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACTIVE",
		"ERROR",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return TrafficMirrorState(idx)
		}
	}
	return TRAFFICMIRRORSTATE_UNKNOWN
}

func (e *TrafficMirrorState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TrafficMirrorState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TrafficMirrorState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TrafficMirrorState) Ref() *TrafficMirrorState {
	return &e
}

/*
Traffic mirror stats description.
*/
type TrafficMirrorStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Name of the session.
	*/
	Name *string `json:"name,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`

	State *TrafficMirrorState `json:"state,omitempty"`
	/*
	  Traffic mirror stats state message.
	*/
	StateMessage *string `json:"stateMessage,omitempty"`

	StatsData *TrafficMirrorStatsData `json:"statsData,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewTrafficMirrorStats() *TrafficMirrorStats {
	p := new(TrafficMirrorStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TrafficMirrorStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.TrafficMirrorStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0.b1/stats/traffic-mirrors/{extId} Get operation
*/
type TrafficMirrorStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTrafficMirrorStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTrafficMirrorStatsApiResponse() *TrafficMirrorStatsApiResponse {
	p := new(TrafficMirrorStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TrafficMirrorStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.TrafficMirrorStatsApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TrafficMirrorStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TrafficMirrorStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTrafficMirrorStatsApiResponseData()
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
Traffic mirror stats data values.
*/
type TrafficMirrorStatsData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Traffic mirror bytes transmitted.
	*/
	TransmitByteCount *int64 `json:"transmitByteCount"`
	/*
	  Traffic mirror number of packets transmitted.
	*/
	TransmitPacketCount *int64 `json:"transmitPacketCount"`
}

func (p *TrafficMirrorStatsData) MarshalJSON() ([]byte, error) {
	type TrafficMirrorStatsDataProxy TrafficMirrorStatsData
	return json.Marshal(struct {
		*TrafficMirrorStatsDataProxy
		TransmitByteCount   *int64 `json:"transmitByteCount,omitempty"`
		TransmitPacketCount *int64 `json:"transmitPacketCount,omitempty"`
	}{
		TrafficMirrorStatsDataProxy: (*TrafficMirrorStatsDataProxy)(p),
		TransmitByteCount:           p.TransmitByteCount,
		TransmitPacketCount:         p.TransmitPacketCount,
	})
}

func NewTrafficMirrorStatsData() *TrafficMirrorStatsData {
	p := new(TrafficMirrorStatsData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TrafficMirrorStatsData"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.TrafficMirrorStatsData"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPC North-South statistics description
*/
type VpcNsStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  VPC North-South string array of egress absolute bytes values
	*/
	NorthSouthEgressBytesAbs []string `json:"northSouthEgressBytesAbs,omitempty"`
	/*
	  VPC North-South string array of egress BPS values
	*/
	NorthSouthEgressBytesPerSec []string `json:"northSouthEgressBytesPerSec,omitempty"`
	/*
	  VPC North-South string array of egress absolute packets values
	*/
	NorthSouthEgressPacketsAbs []string `json:"northSouthEgressPacketsAbs,omitempty"`
	/*
	  VPC North-South string array of egress PPS values
	*/
	NorthSouthEgressPacketsPerSec []string `json:"northSouthEgressPacketsPerSec,omitempty"`
	/*
	  VPC North-South string array of ingress absolute bytes values
	*/
	NorthSouthIngressBytesAbs []string `json:"northSouthIngressBytesAbs,omitempty"`
	/*
	  VPC North-South string array of ingress BPS values
	*/
	NorthSouthIngressBytesPerSec []string `json:"northSouthIngressBytesPerSec,omitempty"`
	/*
	  VPC North-South string array of ingress absolute packets values
	*/
	NorthSouthIngressPacketsAbs []string `json:"northSouthIngressPacketsAbs,omitempty"`
	/*
	  VPC North-South string array of ingress PPS values
	*/
	NorthSouthIngressPacketsPerSec []string `json:"northSouthIngressPacketsPerSec,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVpcNsStats() *VpcNsStats {
	p := new(VpcNsStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpcNsStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.VpcNsStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0.b1/stats/vpc/{vpcExtId}/external-subnets/{extId} Get operation
*/
type VpcNsStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVpcNsStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpcNsStatsApiResponse() *VpcNsStatsApiResponse {
	p := new(VpcNsStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpcNsStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.VpcNsStatsApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VpcNsStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VpcNsStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVpcNsStatsApiResponseData()
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
VPN connection statistics description
*/
type VpnConnectionStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPN connection string array of RX BPS values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/*
	  VPN connection string array of TX BPS values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewVpnConnectionStats() *VpnConnectionStats {
	p := new(VpnConnectionStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpnConnectionStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.VpnConnectionStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0.b1/stats/vpn-connections/{extId} Get operation
*/
type VpnConnectionStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVpnConnectionStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpnConnectionStatsApiResponse() *VpnConnectionStatsApiResponse {
	p := new(VpnConnectionStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpnConnectionStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.b1.stats.VpnConnectionStatsApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VpnConnectionStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *VpnConnectionStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfVpnConnectionStatsApiResponseData()
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

type OneOfVpcNsStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *VpcNsStats            `json:"-"`
}

func NewOneOfVpcNsStatsApiResponseData() *OneOfVpcNsStatsApiResponseData {
	p := new(OneOfVpcNsStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVpcNsStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVpcNsStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case VpcNsStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VpcNsStats)
		}
		*p.oneOfType0 = v.(VpcNsStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVpcNsStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfVpcNsStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(VpcNsStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.VpcNsStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VpcNsStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpcNsStatsApiResponseData"))
}

func (p *OneOfVpcNsStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfVpcNsStatsApiResponseData")
}

type OneOfLayer2StretchStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *Layer2StretchStats    `json:"-"`
}

func NewOneOfLayer2StretchStatsApiResponseData() *OneOfLayer2StretchStatsApiResponseData {
	p := new(OneOfLayer2StretchStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLayer2StretchStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLayer2StretchStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Layer2StretchStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Layer2StretchStats)
		}
		*p.oneOfType0 = v.(Layer2StretchStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfLayer2StretchStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfLayer2StretchStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(Layer2StretchStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.Layer2StretchStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Layer2StretchStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLayer2StretchStatsApiResponseData"))
}

func (p *OneOfLayer2StretchStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfLayer2StretchStatsApiResponseData")
}

type OneOfTrafficMirrorStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *TrafficMirrorStats    `json:"-"`
}

func NewOneOfTrafficMirrorStatsApiResponseData() *OneOfTrafficMirrorStatsApiResponseData {
	p := new(OneOfTrafficMirrorStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTrafficMirrorStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTrafficMirrorStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case TrafficMirrorStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TrafficMirrorStats)
		}
		*p.oneOfType0 = v.(TrafficMirrorStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTrafficMirrorStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTrafficMirrorStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(TrafficMirrorStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.TrafficMirrorStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TrafficMirrorStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTrafficMirrorStatsApiResponseData"))
}

func (p *OneOfTrafficMirrorStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTrafficMirrorStatsApiResponseData")
}

type OneOfVpnConnectionStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *VpnConnectionStats    `json:"-"`
}

func NewOneOfVpnConnectionStatsApiResponseData() *OneOfVpnConnectionStatsApiResponseData {
	p := new(OneOfVpnConnectionStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVpnConnectionStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVpnConnectionStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case VpnConnectionStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VpnConnectionStats)
		}
		*p.oneOfType0 = v.(VpnConnectionStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVpnConnectionStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfVpnConnectionStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(VpnConnectionStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.VpnConnectionStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VpnConnectionStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVpnConnectionStatsApiResponseData"))
}

func (p *OneOfVpnConnectionStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfVpnConnectionStatsApiResponseData")
}

type OneOfTaskReferenceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
}

func NewOneOfTaskReferenceApiResponseData() *OneOfTaskReferenceApiResponseData {
	p := new(OneOfTaskReferenceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTaskReferenceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTaskReferenceApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTaskReferenceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTaskReferenceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTaskReferenceApiResponseData"))
}

func (p *OneOfTaskReferenceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTaskReferenceApiResponseData")
}
