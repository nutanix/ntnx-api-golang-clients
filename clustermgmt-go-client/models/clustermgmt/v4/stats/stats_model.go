/*
 * Generated file models/clustermgmt/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1-beta-1
 *
 * Part of the Nutanix Clustermgmt Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.stats of Nutanix Clustermgmt Versioned APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/stats"
	"time"
)

type ClusterStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes)
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes)
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Overall memory usage(bytes)
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  Storage capacity(bytes)
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes)
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewClusterStats() *ClusterStats {
	p := new(ClusterStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.ClusterStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /clustermgmt/v4.0.b1/stats/clusters/{clusterExtId} Get operation
*/
type ClusterStatsInfoApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfClusterStatsInfoApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewClusterStatsInfoApiResponse() *ClusterStatsInfoApiResponse {
	p := new(ClusterStatsInfoApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStatsInfoApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.ClusterStatsInfoApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterStatsInfoApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ClusterStatsInfoApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfClusterStatsInfoApiResponseData()
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

type ClusterStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes)
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes)
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Overall memory usage(bytes)
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  Storage capacity(bytes)
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes)
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewClusterStatsProjection() *ClusterStatsProjection {
	p := new(ClusterStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.ClusterStatsProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type HostStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes)
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Size of memory(in bytes)
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(ppm)
	*/
	OverallMemoryUsagePpm []TimeValuePair `json:"overallMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of overall memory usage(ppm)
	*/
	OverallMemoryUsagePpmLowerBuf []TimeValuePair `json:"overallMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of overall memory usage(ppm)
	*/
	OverallMemoryUsagePpmUpperBuf []TimeValuePair `json:"overallMemoryUsagePpmUpperBuf,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  Storage capacity(bytes)
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes)
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHostStats() *HostStats {
	p := new(HostStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStats"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.HostStats"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /clustermgmt/v4.0.b1/stats/clusters/{clusterExtId}/hosts/{hostExtId} Get operation
*/
type HostStatsInfoApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfHostStatsInfoApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewHostStatsInfoApiResponse() *HostStatsInfoApiResponse {
	p := new(HostStatsInfoApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStatsInfoApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.HostStatsInfoApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HostStatsInfoApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *HostStatsInfoApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfHostStatsInfoApiResponseData()
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

type HostStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm)
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs)
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs)
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs)
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps)
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps)
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes)
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm)
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps)
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Size of memory(in bytes)
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(ppm)
	*/
	OverallMemoryUsagePpm []TimeValuePair `json:"overallMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of overall memory usage(ppm)
	*/
	OverallMemoryUsagePpmLowerBuf []TimeValuePair `json:"overallMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of overall memory usage(ppm)
	*/
	OverallMemoryUsagePpmUpperBuf []TimeValuePair `json:"overallMemoryUsagePpmUpperBuf,omitempty"`

	StatType *import2.DownSamplingOperator `json:"statType,omitempty"`
	/*
	  Storage capacity(bytes)
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes)
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewHostStatsProjection() *HostStatsProjection {
	p := new(HostStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.HostStatsProjection"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Time - Value pair for time-series stat attributes
*/
type TimeValuePair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Timestamp for given stat attribute(in ISO-8601 format)
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  Value of stat at given timestamp
	*/
	Value *int64 `json:"value,omitempty"`
}

func NewTimeValuePair() *TimeValuePair {
	p := new(TimeValuePair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.TimeValuePair"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.b1.stats.TimeValuePair"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfClusterStatsInfoApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ClusterStats          `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfClusterStatsInfoApiResponseData() *OneOfClusterStatsInfoApiResponseData {
	p := new(OneOfClusterStatsInfoApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterStatsInfoApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterStatsInfoApiResponseData is nil"))
	}
	switch v.(type) {
	case ClusterStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterStats)
		}
		*p.oneOfType0 = v.(ClusterStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfClusterStatsInfoApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfClusterStatsInfoApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ClusterStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.stats.ClusterStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ClusterStats)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterStatsInfoApiResponseData"))
}

func (p *OneOfClusterStatsInfoApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfClusterStatsInfoApiResponseData")
}

type OneOfHostStatsInfoApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *HostStats             `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
}

func NewOneOfHostStatsInfoApiResponseData() *OneOfHostStatsInfoApiResponseData {
	p := new(OneOfHostStatsInfoApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHostStatsInfoApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHostStatsInfoApiResponseData is nil"))
	}
	switch v.(type) {
	case HostStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(HostStats)
		}
		*p.oneOfType0 = v.(HostStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHostStatsInfoApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfHostStatsInfoApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(HostStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.stats.HostStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(HostStats)
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
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHostStatsInfoApiResponseData"))
}

func (p *OneOfHostStatsInfoApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfHostStatsInfoApiResponseData")
}
