/*
 * Generated file models/networking/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Networking APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Fetch stats for VPN connections, VPCs and traffic mirrors
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/prism/v4/config"
)

/*
REST response for all response codes in API path /networking/v4.0/stats/layer2-stretches/{extId} Get operation
*/
type GetLayer2StretchStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetLayer2StretchStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetLayer2StretchStatsApiResponse() *GetLayer2StretchStatsApiResponse {
	p := new(GetLayer2StretchStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.GetLayer2StretchStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetLayer2StretchStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetLayer2StretchStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetLayer2StretchStatsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/stats/traffic-mirrors/{extId} Get operation
*/
type GetTrafficMirrorStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTrafficMirrorStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetTrafficMirrorStatsApiResponse() *GetTrafficMirrorStatsApiResponse {
	p := new(GetTrafficMirrorStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.GetTrafficMirrorStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTrafficMirrorStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTrafficMirrorStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTrafficMirrorStatsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/stats/vpc/{vpcExtId}/external-subnets/{extId} Get operation
*/
type GetVpcNsStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVpcNsStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVpcNsStatsApiResponse() *GetVpcNsStatsApiResponse {
	p := new(GetVpcNsStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.GetVpcNsStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVpcNsStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVpcNsStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVpcNsStatsApiResponseData()
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
REST response for all response codes in API path /networking/v4.0/stats/vpn-connections/{extId} Get operation
*/
type GetVpnConnectionStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVpnConnectionStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVpnConnectionStatsApiResponse() *GetVpnConnectionStatsApiResponse {
	p := new(GetVpnConnectionStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.GetVpnConnectionStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVpnConnectionStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVpnConnectionStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVpnConnectionStatsApiResponseData()
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
Layer2 stretch statistics description
*/
type Layer2StretchStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Layer2Stretch string array of round-trip-time.
	*/
	Rtt []string `json:"rtt,omitempty"`

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPN connection string array of ingress kilobits per second values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/*
	  VPN connection string array of egress kilobits per second values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewLayer2StretchStats() *Layer2StretchStats {
	p := new(Layer2StretchStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.Layer2StretchStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ListenerStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  String array of number of bytes received.
	*/
	NumBytes []string `json:"numBytes,omitempty"`
	/*
	  String array of number of requests received.
	*/
	NumPackets []string `json:"numPackets,omitempty"`
	/*
	  String array of number of packets received.
	*/
	NumRequests []string `json:"numRequests,omitempty"`
	/*
	  Load balancer session listener virtual IP address.
	*/
	VirtualIpAddress *string `json:"virtualIpAddress,omitempty"`
}

func NewListenerStats() *ListenerStats {
	p := new(ListenerStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.ListenerStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Load balancer session statistics description.
*/
type LoadBalancerSessionStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of listener stats associated with the load balancer session.
	*/
	ListenerStats []ListenerStats `json:"listenerStats,omitempty"`

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  List of target stats associated with the load balancer session.
	*/
	TargetStats []TargetStats `json:"targetStats,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewLoadBalancerSessionStats() *LoadBalancerSessionStats {
	p := new(LoadBalancerSessionStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.LoadBalancerSessionStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0/stats/load-balancer-sessions/{extId} Get operation
*/
type LoadBalancerSessionStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfLoadBalancerSessionStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLoadBalancerSessionStatsApiResponse() *LoadBalancerSessionStatsApiResponse {
	p := new(LoadBalancerSessionStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.LoadBalancerSessionStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *LoadBalancerSessionStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *LoadBalancerSessionStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfLoadBalancerSessionStatsApiResponseData()
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
	/*
	  UUID of the policy for which the counters need to be cleared.
	*/
	RoutingPolicyExtId *string `json:"routingPolicyExtId,omitempty"`
	/*
	  ExtId of the VPC extId to which the routing policy belongs.
	*/
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
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewStatsQueryResponseBase() *StatsQueryResponseBase {
	p := new(StatsQueryResponseBase)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.StatsQueryResponseBase"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TargetStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  String array of number of bytes received.
	*/
	NumBytes []string `json:"numBytes,omitempty"`
	/*
	  String array of number of requests received.
	*/
	NumPackets []string `json:"numPackets,omitempty"`
	/*
	  String array of number of packets received.
	*/
	NumRequests []string `json:"numRequests,omitempty"`
	/*
	  Load balancer session Target vNIC reference.
	*/
	VirtualNicReference *string `json:"virtualNicReference,omitempty"`
}

func NewTargetStats() *TargetStats {
	p := new(TargetStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TargetStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /networking/v4.0/stats/routing-policies/$actions/clear Post operation
*/
type TaskReferenceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTaskReferenceApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTaskReferenceApiResponse() *TaskReferenceApiResponse {
	p := new(TaskReferenceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TaskReferenceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
Traffic mirror stats description.
*/
type TrafficMirrorStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Traffic mirror bytes transmitted.
	*/
	TransmitByteCount []int64 `json:"transmitByteCount,omitempty"`
	/*
	  Traffic mirror number of packets transmitted.
	*/
	TransmitPacketCount []int64 `json:"transmitPacketCount,omitempty"`
}

func NewTrafficMirrorStats() *TrafficMirrorStats {
	p := new(TrafficMirrorStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.TrafficMirrorStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
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

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVpcNsStats() *VpcNsStats {
	p := new(VpcNsStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpcNsStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
VPN connection statistics description
*/
type VpnConnectionStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	StatType *import3.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  VPN connection string array of ingress kilobits per second values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/*
	  VPN connection string array of egress kilobits per second values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewVpnConnectionStats() *VpnConnectionStats {
	p := new(VpnConnectionStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpnConnectionStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetVpnConnectionStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *VpnConnectionStats    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVpnConnectionStatsApiResponseData() *OneOfGetVpnConnectionStatsApiResponseData {
	p := new(OneOfGetVpnConnectionStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVpnConnectionStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVpnConnectionStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetVpnConnectionStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVpnConnectionStatsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVpnConnectionStatsApiResponseData"))
}

func (p *OneOfGetVpnConnectionStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVpnConnectionStatsApiResponseData")
}

type OneOfTaskReferenceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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

type OneOfGetLayer2StretchStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Layer2StretchStats    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetLayer2StretchStatsApiResponseData() *OneOfGetLayer2StretchStatsApiResponseData {
	p := new(OneOfGetLayer2StretchStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetLayer2StretchStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetLayer2StretchStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetLayer2StretchStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetLayer2StretchStatsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetLayer2StretchStatsApiResponseData"))
}

func (p *OneOfGetLayer2StretchStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetLayer2StretchStatsApiResponseData")
}

type OneOfGetTrafficMirrorStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *TrafficMirrorStats    `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetTrafficMirrorStatsApiResponseData() *OneOfGetTrafficMirrorStatsApiResponseData {
	p := new(OneOfGetTrafficMirrorStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTrafficMirrorStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTrafficMirrorStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetTrafficMirrorStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetTrafficMirrorStatsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTrafficMirrorStatsApiResponseData"))
}

func (p *OneOfGetTrafficMirrorStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetTrafficMirrorStatsApiResponseData")
}

type OneOfLoadBalancerSessionStatsApiResponseData struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    *LoadBalancerSessionStats `json:"-"`
	oneOfType400  *import1.ErrorResponse    `json:"-"`
}

func NewOneOfLoadBalancerSessionStatsApiResponseData() *OneOfLoadBalancerSessionStatsApiResponseData {
	p := new(OneOfLoadBalancerSessionStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLoadBalancerSessionStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLoadBalancerSessionStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case LoadBalancerSessionStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(LoadBalancerSessionStats)
		}
		*p.oneOfType0 = v.(LoadBalancerSessionStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfLoadBalancerSessionStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfLoadBalancerSessionStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(LoadBalancerSessionStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.LoadBalancerSessionStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(LoadBalancerSessionStats)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLoadBalancerSessionStatsApiResponseData"))
}

func (p *OneOfLoadBalancerSessionStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfLoadBalancerSessionStatsApiResponseData")
}

type OneOfGetVpcNsStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *VpcNsStats            `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVpcNsStatsApiResponseData() *OneOfGetVpcNsStatsApiResponseData {
	p := new(OneOfGetVpcNsStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVpcNsStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVpcNsStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetVpcNsStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVpcNsStatsApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "networking.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVpcNsStatsApiResponseData"))
}

func (p *OneOfGetVpcNsStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVpcNsStatsApiResponseData")
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
