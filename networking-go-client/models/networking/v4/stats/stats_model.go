/*
 * Generated file models/networking/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1-alpha-1
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
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/common/v1/stats"
	import3 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/error"
)

/**
Layer2 stretch statistics description
*/
type Layer2StretchStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  Layer2Stretch string array of round-trip-time.
	*/
	Rtt []string `json:"rtt,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  VPN connection string array of RX BPS values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/**
	  VPN connection string array of TX BPS values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewLayer2StretchStats() *Layer2StretchStats {
	p := new(Layer2StretchStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.Layer2StretchStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.Layer2StretchStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /networking/v4.0.a1/stats/layer2-stretches/{extId} Get operation
*/
type Layer2StretchStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfLayer2StretchStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewLayer2StretchStatsApiResponse() *Layer2StretchStatsApiResponse {
	p := new(Layer2StretchStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.Layer2StretchStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.Layer2StretchStatsApiResponse"}
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

/**
Status of the clear routing policy counters request.
*/
type RoutingPolicyClearCounterResult struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether the clearing counters operation was successful or not.
	*/
	ClearCountersResponse *string `json:"clearCountersResponse,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewRoutingPolicyClearCounterResult() *RoutingPolicyClearCounterResult {
	p := new(RoutingPolicyClearCounterResult)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.RoutingPolicyClearCounterResult"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.RoutingPolicyClearCounterResult"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /networking/v4.0.a1/stats/routing-policies/{extId}/$actions/clear Post operation
*/
type RoutingPolicyClearCountersApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRoutingPolicyClearCountersApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRoutingPolicyClearCountersApiResponse() *RoutingPolicyClearCountersApiResponse {
	p := new(RoutingPolicyClearCountersApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.RoutingPolicyClearCountersApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.RoutingPolicyClearCountersApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RoutingPolicyClearCountersApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RoutingPolicyClearCountersApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRoutingPolicyClearCountersApiResponseData()
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

/**
VPC UUID to reset all routing policy counters to zero.
*/
type RoutingPolicyClearCountersBody struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VpcExtId *string `json:"vpcExtId"`
}

func (p *RoutingPolicyClearCountersBody) MarshalJSON() ([]byte, error) {
	type RoutingPolicyClearCountersBodyProxy RoutingPolicyClearCountersBody
	return json.Marshal(struct {
		*RoutingPolicyClearCountersBodyProxy
		VpcExtId *string `json:"vpcExtId,omitempty"`
	}{
		RoutingPolicyClearCountersBodyProxy: (*RoutingPolicyClearCountersBodyProxy)(p),
		VpcExtId:                            p.VpcExtId,
	})
}

func NewRoutingPolicyClearCountersBody() *RoutingPolicyClearCountersBody {
	p := new(RoutingPolicyClearCountersBody)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.RoutingPolicyClearCountersBody"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.RoutingPolicyClearCountersBody"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Response of statistics query.
*/
type StatsQueryResponseBase struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewStatsQueryResponseBase() *StatsQueryResponseBase {
	p := new(StatsQueryResponseBase)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.StatsQueryResponseBase"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.StatsQueryResponseBase"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
VPN North-South statistics description
*/
type VpcNsStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/**
	  VPC North-South string array of egress absolute bytes values
	*/
	NorthSouthEgressBytesAbs []string `json:"northSouthEgressBytesAbs,omitempty"`
	/**
	  VPC North-South string array of egress BPS values
	*/
	NorthSouthEgressBytesPerSec []string `json:"northSouthEgressBytesPerSec,omitempty"`
	/**
	  VPC North-South string array of egress absolute packets values
	*/
	NorthSouthEgressPacketsAbs []string `json:"northSouthEgressPacketsAbs,omitempty"`
	/**
	  VPC North-South string array of egress PPS values
	*/
	NorthSouthEgressPacketsPerSec []string `json:"northSouthEgressPacketsPerSec,omitempty"`
	/**
	  VPC North-South string array of ingress absolute bytes values
	*/
	NorthSouthIngressBytesAbs []string `json:"northSouthIngressBytesAbs,omitempty"`
	/**
	  VPC North-South string array of ingress BPS values
	*/
	NorthSouthIngressBytesPerSec []string `json:"northSouthIngressBytesPerSec,omitempty"`
	/**
	  VPC North-South string array of ingress absolute packets values
	*/
	NorthSouthIngressPacketsAbs []string `json:"northSouthIngressPacketsAbs,omitempty"`
	/**
	  VPC North-South string array of ingress PPS values
	*/
	NorthSouthIngressPacketsPerSec []string `json:"northSouthIngressPacketsPerSec,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVpcNsStats() *VpcNsStats {
	p := new(VpcNsStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpcNsStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.VpcNsStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /networking/v4.0.a1/stats/vpc/{vpcUuid}/external-subnets/{extSubnetUuid} Get operation
*/
type VpcNsStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVpcNsStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpcNsStatsApiResponse() *VpcNsStatsApiResponse {
	p := new(VpcNsStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpcNsStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.VpcNsStatsApiResponse"}
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

/**
VPN connection statistics description
*/
type VpnConnectionStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  VPN connection string array of RX BPS values
	*/
	ThroughputRxKbps []string `json:"throughputRxKbps,omitempty"`
	/**
	  VPN connection string array of TX BPS values
	*/
	ThroughputTxKbps []string `json:"throughputTxKbps,omitempty"`
}

func NewVpnConnectionStats() *VpnConnectionStats {
	p := new(VpnConnectionStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpnConnectionStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.VpnConnectionStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /networking/v4.0.a1/stats/vpn-connections/{extId} Get operation
*/
type VpnConnectionStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfVpnConnectionStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewVpnConnectionStatsApiResponse() *VpnConnectionStatsApiResponse {
	p := new(VpnConnectionStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "networking.v4.stats.VpnConnectionStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "networking.v4.r0.a1.stats.VpnConnectionStatsApiResponse"}
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

type OneOfRoutingPolicyClearCountersApiResponseData struct {
	Discriminator *string                          `json:"-"`
	ObjectType_   *string                          `json:"-"`
	oneOfType400  *import3.ErrorResponse           `json:"-"`
	oneOfType0    *RoutingPolicyClearCounterResult `json:"-"`
}

func NewOneOfRoutingPolicyClearCountersApiResponseData() *OneOfRoutingPolicyClearCountersApiResponseData {
	p := new(OneOfRoutingPolicyClearCountersApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRoutingPolicyClearCountersApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRoutingPolicyClearCountersApiResponseData is nil"))
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
	case RoutingPolicyClearCounterResult:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RoutingPolicyClearCounterResult)
		}
		*p.oneOfType0 = v.(RoutingPolicyClearCounterResult)
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

func (p *OneOfRoutingPolicyClearCountersApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfRoutingPolicyClearCountersApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(RoutingPolicyClearCounterResult)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "networking.v4.stats.RoutingPolicyClearCounterResult" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RoutingPolicyClearCounterResult)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRoutingPolicyClearCountersApiResponseData"))
}

func (p *OneOfRoutingPolicyClearCountersApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfRoutingPolicyClearCountersApiResponseData")
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
