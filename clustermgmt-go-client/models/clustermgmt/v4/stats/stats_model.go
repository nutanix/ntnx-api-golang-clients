/*
 * Generated file models/clustermgmt/v4/stats/stats_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Cluster Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.stats of Nutanix Cluster Management APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/stats"
	"time"
)

/*
Cluster entity statistic attributes.
*/
type ClusterStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number.
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS.
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS.
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS.
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS.
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz.
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  CPU usage (Hz)
	*/
	CpuUsageHz []TimeValuePair `json:"cpuUsageHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes).
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  NCC check score indicating the health of the entity. The value to health mapping is as follows: Good: 100, Info: 98, Warning: 74, Critical: 24, Error: 13, Unknown: -1.
	*/
	HealthCheckScore []TimeValuePair `json:"healthCheckScore,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes).
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Size of memory(in bytes).
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(bytes).
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`
	/*
	  Overall savings (bytes)
	*/
	OverallSavingsBytes []TimeValuePair `json:"overallSavingsBytes,omitempty"`
	/*
	  Overall saving ratio
	*/
	OverallSavingsRatio []TimeValuePair `json:"overallSavingsRatio,omitempty"`
	/*
	  Power instant consumption (watt)
	*/
	PowerConsumptionInstantWatt []TimeValuePair `json:"powerConsumptionInstantWatt,omitempty"`
	/*
	  Recycle bin usage (bytes)
	*/
	RecycleBinUsageBytes []TimeValuePair `json:"recycleBinUsageBytes,omitempty"`
	/*
	  Snapshot capacity (bytes)
	*/
	SnapshotCapacityBytes []TimeValuePair `json:"snapshotCapacityBytes,omitempty"`
	/*
	  Storage capacity(bytes).
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes).
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterStats

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

func (p *ClusterStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateHypervisorMemoryUsagePpm")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmLowerBuf")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmUpperBuf")
	delete(allFields, "controllerAvgIoLatencyUsecs")
	delete(allFields, "controllerAvgIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecs")
	delete(allFields, "controllerAvgReadIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecs")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumIopsLowerBuf")
	delete(allFields, "controllerNumIopsUpperBuf")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumReadIopsLowerBuf")
	delete(allFields, "controllerNumReadIopsUpperBuf")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerNumWriteIopsLowerBuf")
	delete(allFields, "controllerNumWriteIopsUpperBuf")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerReadIoBandwidthKbpsUpperBuf")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerWriteIoBandwidthKbpsUpperBuf")
	delete(allFields, "cpuCapacityHz")
	delete(allFields, "cpuUsageHz")
	delete(allFields, "extId")
	delete(allFields, "freePhysicalStorageBytes")
	delete(allFields, "healthCheckScore")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorCpuUsagePpmLowerBuf")
	delete(allFields, "hypervisorCpuUsagePpmUpperBuf")
	delete(allFields, "ioBandwidthKbps")
	delete(allFields, "ioBandwidthKbpsLowerBuf")
	delete(allFields, "ioBandwidthKbpsUpperBuf")
	delete(allFields, "links")
	delete(allFields, "logicalStorageUsageBytes")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "overallMemoryUsageBytes")
	delete(allFields, "overallSavingsBytes")
	delete(allFields, "overallSavingsRatio")
	delete(allFields, "powerConsumptionInstantWatt")
	delete(allFields, "recycleBinUsageBytes")
	delete(allFields, "snapshotCapacityBytes")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterStats() *ClusterStats {
	p := new(ClusterStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /clustermgmt/v4.1/stats/clusters/{extId} Get operation
*/
type ClusterStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfClusterStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ClusterStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterStatsApiResponse

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

func (p *ClusterStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterStatsApiResponse(*known)

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

func NewClusterStatsApiResponse() *ClusterStatsApiResponse {
	p := new(ClusterStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ClusterStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfClusterStatsApiResponseData()
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
	  Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number.
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS.
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS.
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS.
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS.
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz.
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  CPU usage (Hz)
	*/
	CpuUsageHz []TimeValuePair `json:"cpuUsageHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes).
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  NCC check score indicating the health of the entity. The value to health mapping is as follows: Good: 100, Info: 98, Warning: 74, Critical: 24, Error: 13, Unknown: -1.
	*/
	HealthCheckScore []TimeValuePair `json:"healthCheckScore,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes).
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Size of memory(in bytes).
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(bytes).
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`
	/*
	  Overall savings (bytes)
	*/
	OverallSavingsBytes []TimeValuePair `json:"overallSavingsBytes,omitempty"`
	/*
	  Overall saving ratio
	*/
	OverallSavingsRatio []TimeValuePair `json:"overallSavingsRatio,omitempty"`
	/*
	  Power instant consumption (watt)
	*/
	PowerConsumptionInstantWatt []TimeValuePair `json:"powerConsumptionInstantWatt,omitempty"`
	/*
	  Recycle bin usage (bytes)
	*/
	RecycleBinUsageBytes []TimeValuePair `json:"recycleBinUsageBytes,omitempty"`
	/*
	  Snapshot capacity (bytes)
	*/
	SnapshotCapacityBytes []TimeValuePair `json:"snapshotCapacityBytes,omitempty"`
	/*
	  Storage capacity(bytes).
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes).
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterStatsProjection

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

func (p *ClusterStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateHypervisorMemoryUsagePpm")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmLowerBuf")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmUpperBuf")
	delete(allFields, "controllerAvgIoLatencyUsecs")
	delete(allFields, "controllerAvgIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecs")
	delete(allFields, "controllerAvgReadIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecs")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumIopsLowerBuf")
	delete(allFields, "controllerNumIopsUpperBuf")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumReadIopsLowerBuf")
	delete(allFields, "controllerNumReadIopsUpperBuf")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerNumWriteIopsLowerBuf")
	delete(allFields, "controllerNumWriteIopsUpperBuf")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerReadIoBandwidthKbpsUpperBuf")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerWriteIoBandwidthKbpsUpperBuf")
	delete(allFields, "cpuCapacityHz")
	delete(allFields, "cpuUsageHz")
	delete(allFields, "extId")
	delete(allFields, "freePhysicalStorageBytes")
	delete(allFields, "healthCheckScore")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorCpuUsagePpmLowerBuf")
	delete(allFields, "hypervisorCpuUsagePpmUpperBuf")
	delete(allFields, "ioBandwidthKbps")
	delete(allFields, "ioBandwidthKbpsLowerBuf")
	delete(allFields, "ioBandwidthKbpsUpperBuf")
	delete(allFields, "links")
	delete(allFields, "logicalStorageUsageBytes")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "overallMemoryUsageBytes")
	delete(allFields, "overallSavingsBytes")
	delete(allFields, "overallSavingsRatio")
	delete(allFields, "powerConsumptionInstantWatt")
	delete(allFields, "recycleBinUsageBytes")
	delete(allFields, "snapshotCapacityBytes")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterStatsProjection() *ClusterStatsProjection {
	p := new(ClusterStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.ClusterStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DiskStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Average I/O latency.
	*/
	DiskAvgIoLatencyMicrosec []import3.TimeIntValuePair `json:"diskAvgIoLatencyMicrosec,omitempty"`
	/*
	  Lower limit of data transfer that a Disk can handle per second.
	*/
	DiskBaseIoBandwidthkbps []import3.TimeIntValuePair `json:"diskBaseIoBandwidthkbps,omitempty"`
	/*
	  Lower limit of the latency of I/O operations that the Disk can handle without exceeding its standard latency level.
	*/
	DiskBaseIoLatencyMicrosec []import3.TimeIntValuePair `json:"diskBaseIoLatencyMicrosec,omitempty"`
	/*
	  Lower limit of I/O operations that a Disk can perform per second.
	*/
	DiskBaseNumIops []import3.TimeIntValuePair `json:"diskBaseNumIops,omitempty"`
	/*
	  Lower buffer capacity average read I/O latency, measured in microseconds (μs).
	*/
	DiskBaseReadIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskBaseReadIoAvgLatencyMicrosec,omitempty"`
	/*
	  Lower buffer capacity for the amount of I/O bandwidth that a Disk can handle read operations.
	*/
	DiskBaseReadIoBandwidthkbps []import3.TimeIntValuePair `json:"diskBaseReadIoBandwidthkbps,omitempty"`
	/*
	  Lower buffer capacity for the number of read IOPS that a Disk can handle.
	*/
	DiskBaseReadIops []import3.TimeIntValuePair `json:"diskBaseReadIops,omitempty"`
	/*
	  Lower buffer capacity average write I/O latency, measured in microseconds (μs).
	*/
	DiskBaseWriteIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskBaseWriteIoAvgLatencyMicrosec,omitempty"`
	/*
	  Lower buffer capacity for the amount of I/O bandwidth that a Disk can handle write operations.
	*/
	DiskBaseWriteIoBandwidthkbps []import3.TimeIntValuePair `json:"diskBaseWriteIoBandwidthkbps,omitempty"`
	/*
	  Lower buffer capacity of a number of write I/O per second.
	*/
	DiskBaseWriteIops []import3.TimeIntValuePair `json:"diskBaseWriteIops,omitempty"`
	/*
	  Total storage capacity of a device in bytes.
	*/
	DiskCapacityBytes []import3.TimeIntValuePair `json:"diskCapacityBytes,omitempty"`
	/*
	  Free storage space available on the Disk, measured in bytes.
	*/
	DiskFreeBytes []import3.TimeIntValuePair `json:"diskFreeBytes,omitempty"`
	/*
	  I/O bandwidth in KB per second.
	*/
	DiskIoBandwidthkbps []import3.TimeIntValuePair `json:"diskIoBandwidthkbps,omitempty"`
	/*
	  Number of I/O operations that a Disk performs per second.
	*/
	DiskNumIops []import3.TimeIntValuePair `json:"diskNumIops,omitempty"`
	/*
	  Upper limit of data transfer that a Disk can handle per second.
	*/
	DiskPeakIoBandwidthkbps []import3.TimeIntValuePair `json:"diskPeakIoBandwidthkbps,omitempty"`
	/*
	  Upper limit of the latency of I/O operations that the Disk can handle without exceeding its standard latency level.
	*/
	DiskPeakIoLatencyMicrosec []import3.TimeIntValuePair `json:"diskPeakIoLatencyMicrosec,omitempty"`
	/*
	  Upper limit of I/O operations that a Disk performs per second.
	*/
	DiskPeakNumIops []import3.TimeIntValuePair `json:"diskPeakNumIops,omitempty"`
	/*
	  Upper buffer capacity average read I/O latency, measured in microseconds (μs).
	*/
	DiskPeakReadIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskPeakReadIoAvgLatencyMicrosec,omitempty"`
	/*
	  Upper buffer capacity for the amount of I/O bandwidth that a Disk can handle read operations.
	*/
	DiskPeakReadIoBandwidthkbps []import3.TimeIntValuePair `json:"diskPeakReadIoBandwidthkbps,omitempty"`
	/*
	  Upper buffer capacity for the number of read IOPS that a Disk can handle.
	*/
	DiskPeakReadIops []import3.TimeIntValuePair `json:"diskPeakReadIops,omitempty"`
	/*
	  Upper buffer capacity average write I/O latency, measured in microseconds (μs).
	*/
	DiskPeakWriteIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskPeakWriteIoAvgLatencyMicrosec,omitempty"`
	/*
	  Upper buffer capacity for the amount of I/O bandwidth that a Disk can handle write operations.
	*/
	DiskPeakWriteIoBandwidthkbps []import3.TimeIntValuePair `json:"diskPeakWriteIoBandwidthkbps,omitempty"`
	/*
	  Upper buffer capacity of a number of write I/O per second.
	*/
	DiskPeakWriteIops []import3.TimeIntValuePair `json:"diskPeakWriteIops,omitempty"`
	/*
	  Average read I/O latency, measured in microseconds (μs).
	*/
	DiskReadIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskReadIoAvgLatencyMicrosec,omitempty"`
	/*
	  Number of Disk read I/O per second as reported by Stargate.
	*/
	DiskReadIoBandwidthkbps []import3.TimeIntValuePair `json:"diskReadIoBandwidthkbps,omitempty"`
	/*
	  Disk read I/O, expressed in parts per million.
	*/
	DiskReadIoPpm []import3.TimeIntValuePair `json:"diskReadIoPpm,omitempty"`
	/*
	  Number of read I/O per second.
	*/
	DiskReadIops []import3.TimeIntValuePair `json:"diskReadIops,omitempty"`
	/*
	  Amount of storage currently being used, measured in bytes.
	*/
	DiskUsageBytes []import3.TimeIntValuePair `json:"diskUsageBytes,omitempty"`
	/*
	  Disk space used on a storage device, expressed in parts per million (ppm).
	*/
	DiskUsagePpm []import3.TimeIntValuePair `json:"diskUsagePpm,omitempty"`
	/*
	  Average write I/O latency, measured in microseconds (μs).
	*/
	DiskWriteIoAvgLatencyMicrosec []import3.TimeIntValuePair `json:"diskWriteIoAvgLatencyMicrosec,omitempty"`
	/*
	  Number of Disk write I/O per second reported by Stargate.
	*/
	DiskWriteIoBandwidthkbps []import3.TimeIntValuePair `json:"diskWriteIoBandwidthkbps,omitempty"`
	/*
	  Disk write I/O, expressed in parts per million.
	*/
	DiskWriteIoPpm []import3.TimeIntValuePair `json:"diskWriteIoPpm,omitempty"`
	/*
	  Number of write I/O per second.
	*/
	DiskWriteIops []import3.TimeIntValuePair `json:"diskWriteIops,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DiskStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DiskStats

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

func (p *DiskStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DiskStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = DiskStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "diskAvgIoLatencyMicrosec")
	delete(allFields, "diskBaseIoBandwidthkbps")
	delete(allFields, "diskBaseIoLatencyMicrosec")
	delete(allFields, "diskBaseNumIops")
	delete(allFields, "diskBaseReadIoAvgLatencyMicrosec")
	delete(allFields, "diskBaseReadIoBandwidthkbps")
	delete(allFields, "diskBaseReadIops")
	delete(allFields, "diskBaseWriteIoAvgLatencyMicrosec")
	delete(allFields, "diskBaseWriteIoBandwidthkbps")
	delete(allFields, "diskBaseWriteIops")
	delete(allFields, "diskCapacityBytes")
	delete(allFields, "diskFreeBytes")
	delete(allFields, "diskIoBandwidthkbps")
	delete(allFields, "diskNumIops")
	delete(allFields, "diskPeakIoBandwidthkbps")
	delete(allFields, "diskPeakIoLatencyMicrosec")
	delete(allFields, "diskPeakNumIops")
	delete(allFields, "diskPeakReadIoAvgLatencyMicrosec")
	delete(allFields, "diskPeakReadIoBandwidthkbps")
	delete(allFields, "diskPeakReadIops")
	delete(allFields, "diskPeakWriteIoAvgLatencyMicrosec")
	delete(allFields, "diskPeakWriteIoBandwidthkbps")
	delete(allFields, "diskPeakWriteIops")
	delete(allFields, "diskReadIoAvgLatencyMicrosec")
	delete(allFields, "diskReadIoBandwidthkbps")
	delete(allFields, "diskReadIoPpm")
	delete(allFields, "diskReadIops")
	delete(allFields, "diskUsageBytes")
	delete(allFields, "diskUsagePpm")
	delete(allFields, "diskWriteIoAvgLatencyMicrosec")
	delete(allFields, "diskWriteIoBandwidthkbps")
	delete(allFields, "diskWriteIoPpm")
	delete(allFields, "diskWriteIops")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewDiskStats() *DiskStats {
	p := new(DiskStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.DiskStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /clustermgmt/v4.1/stats/disks/{extId} Get operation
*/
type GetDiskStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDiskStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDiskStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDiskStatsApiResponse

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

func (p *GetDiskStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDiskStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetDiskStatsApiResponse(*known)

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

func NewGetDiskStatsApiResponse() *GetDiskStatsApiResponse {
	p := new(GetDiskStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.GetDiskStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDiskStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDiskStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDiskStatsApiResponseData()
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
REST response for all response codes in API path /clustermgmt/v4.1/stats/storage-containers/{extId} Get operation
*/
type GetStorageContainerStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetStorageContainerStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetStorageContainerStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetStorageContainerStatsApiResponse

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

func (p *GetStorageContainerStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetStorageContainerStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetStorageContainerStatsApiResponse(*known)

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

func NewGetStorageContainerStatsApiResponse() *GetStorageContainerStatsApiResponse {
	p := new(GetStorageContainerStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.GetStorageContainerStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetStorageContainerStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetStorageContainerStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetStorageContainerStatsApiResponseData()
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
Host entity statistic attributes.
*/
type HostStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number.
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS.
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS.
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS.
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS.
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz.
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  CPU usage (Hz)
	*/
	CpuUsageHz []TimeValuePair `json:"cpuUsageHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes).
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  NCC check score indicating the health of the entity. The value to health mapping is as follows: Good: 100, Info: 98, Warning: 74, Critical: 24, Error: 13, Unknown: -1.
	*/
	HealthCheckScore []TimeValuePair `json:"healthCheckScore,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes).
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Size of memory(in bytes).
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(bytes).
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`
	/*
	  Overall memory usage(ppm).
	*/
	OverallMemoryUsagePpm []TimeValuePair `json:"overallMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of overall memory usage(ppm).
	*/
	OverallMemoryUsagePpmLowerBuf []TimeValuePair `json:"overallMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of overall memory usage(ppm).
	*/
	OverallMemoryUsagePpmUpperBuf []TimeValuePair `json:"overallMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Power instant consumption (watt)
	*/
	PowerConsumptionInstantWatt []TimeValuePair `json:"powerConsumptionInstantWatt,omitempty"`
	/*
	  Storage capacity(bytes).
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes).
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *HostStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostStats

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

func (p *HostStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HostStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateHypervisorMemoryUsagePpm")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmLowerBuf")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmUpperBuf")
	delete(allFields, "controllerAvgIoLatencyUsecs")
	delete(allFields, "controllerAvgIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecs")
	delete(allFields, "controllerAvgReadIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecs")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumIopsLowerBuf")
	delete(allFields, "controllerNumIopsUpperBuf")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumReadIopsLowerBuf")
	delete(allFields, "controllerNumReadIopsUpperBuf")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerNumWriteIopsLowerBuf")
	delete(allFields, "controllerNumWriteIopsUpperBuf")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerReadIoBandwidthKbpsUpperBuf")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerWriteIoBandwidthKbpsUpperBuf")
	delete(allFields, "cpuCapacityHz")
	delete(allFields, "cpuUsageHz")
	delete(allFields, "extId")
	delete(allFields, "freePhysicalStorageBytes")
	delete(allFields, "healthCheckScore")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorCpuUsagePpmLowerBuf")
	delete(allFields, "hypervisorCpuUsagePpmUpperBuf")
	delete(allFields, "ioBandwidthKbps")
	delete(allFields, "ioBandwidthKbpsLowerBuf")
	delete(allFields, "ioBandwidthKbpsUpperBuf")
	delete(allFields, "links")
	delete(allFields, "logicalStorageUsageBytes")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "overallMemoryUsageBytes")
	delete(allFields, "overallMemoryUsagePpm")
	delete(allFields, "overallMemoryUsagePpmLowerBuf")
	delete(allFields, "overallMemoryUsagePpmUpperBuf")
	delete(allFields, "powerConsumptionInstantWatt")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHostStats() *HostStats {
	p := new(HostStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /clustermgmt/v4.1/stats/clusters/{clusterExtId}/hosts/{extId} Get operation
*/
type HostStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfHostStatsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *HostStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostStatsApiResponse

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

func (p *HostStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HostStatsApiResponse(*known)

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

func NewHostStatsApiResponse() *HostStatsApiResponse {
	p := new(HostStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HostStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *HostStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfHostStatsApiResponseData()
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
	  Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpm []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmLowerBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Aggregate Hypervisor Memory Usage(ppm).
	*/
	AggregateHypervisorMemoryUsagePpmUpperBuf []TimeValuePair `json:"aggregateHypervisorMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecs []TimeValuePair `json:"controllerAvgIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average IO Latency(usecs).
	*/
	ControllerAvgIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecs []TimeValuePair `json:"controllerAvgReadIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Read IO Latency(usecs).
	*/
	ControllerAvgReadIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgReadIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecs,omitempty"`
	/*
	  Lower Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsLowerBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Average Write IO Latency(usecs).
	*/
	ControllerAvgWriteIoLatencyUsecsUpperBuf []TimeValuePair `json:"controllerAvgWriteIoLatencyUsecsUpperBuf,omitempty"`
	/*
	  Controller IOPS Number.
	*/
	ControllerNumIops []TimeValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Lower Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsLowerBuf []TimeValuePair `json:"controllerNumIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IOPS Number.
	*/
	ControllerNumIopsUpperBuf []TimeValuePair `json:"controllerNumIopsUpperBuf,omitempty"`
	/*
	  Number of controller read IOPS.
	*/
	ControllerNumReadIops []TimeValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Lower Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsLowerBuf []TimeValuePair `json:"controllerNumReadIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller read IOPS.
	*/
	ControllerNumReadIopsUpperBuf []TimeValuePair `json:"controllerNumReadIopsUpperBuf,omitempty"`
	/*
	  Number of controller write IOPS.
	*/
	ControllerNumWriteIops []TimeValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Lower Buf value of Number of controller write IoPS.
	*/
	ControllerNumWriteIopsLowerBuf []TimeValuePair `json:"controllerNumWriteIopsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Number of controller write IOPS.
	*/
	ControllerNumWriteIopsUpperBuf []TimeValuePair `json:"controllerNumWriteIopsUpperBuf,omitempty"`
	/*
	  Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbps []TimeValuePair `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Read IO Bandwidth(kBps).
	*/
	ControllerReadIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerReadIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbps []TimeValuePair `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsLowerBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller Write IO Bandwidth(kBps).
	*/
	ControllerWriteIoBandwidthKbpsUpperBuf []TimeValuePair `json:"controllerWriteIoBandwidthKbpsUpperBuf,omitempty"`
	/*
	  CPU capacity in Hz.
	*/
	CpuCapacityHz []TimeValuePair `json:"cpuCapacityHz,omitempty"`
	/*
	  CPU usage (Hz)
	*/
	CpuUsageHz []TimeValuePair `json:"cpuUsageHz,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Free physical space(bytes).
	*/
	FreePhysicalStorageBytes []TimeValuePair `json:"freePhysicalStorageBytes,omitempty"`
	/*
	  NCC check score indicating the health of the entity. The value to health mapping is as follows: Good: 100, Info: 98, Warning: 74, Critical: 24, Error: 13, Unknown: -1.
	*/
	HealthCheckScore []TimeValuePair `json:"healthCheckScore,omitempty"`
	/*
	  Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpm []TimeValuePair `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  Lower Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmLowerBuf []TimeValuePair `json:"hypervisorCpuUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of Hypervisor CPU Usage(ppm).
	*/
	HypervisorCpuUsagePpmUpperBuf []TimeValuePair `json:"hypervisorCpuUsagePpmUpperBuf,omitempty"`
	/*
	  Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbps []TimeValuePair `json:"ioBandwidthKbps,omitempty"`
	/*
	  Lower Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsLowerBuf []TimeValuePair `json:"ioBandwidthKbpsLowerBuf,omitempty"`
	/*
	  Upper Buf value of Controller IO Bandwidth(kBps).
	*/
	IoBandwidthKbpsUpperBuf []TimeValuePair `json:"ioBandwidthKbpsUpperBuf,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Logical storage usage(bytes).
	*/
	LogicalStorageUsageBytes []TimeValuePair `json:"logicalStorageUsageBytes,omitempty"`
	/*
	  Size of memory(in bytes).
	*/
	MemoryCapacityBytes []TimeValuePair `json:"memoryCapacityBytes,omitempty"`
	/*
	  Overall memory usage(bytes).
	*/
	OverallMemoryUsageBytes []TimeValuePair `json:"overallMemoryUsageBytes,omitempty"`
	/*
	  Overall memory usage(ppm).
	*/
	OverallMemoryUsagePpm []TimeValuePair `json:"overallMemoryUsagePpm,omitempty"`
	/*
	  Lower Buf value of overall memory usage(ppm).
	*/
	OverallMemoryUsagePpmLowerBuf []TimeValuePair `json:"overallMemoryUsagePpmLowerBuf,omitempty"`
	/*
	  Upper Buf value of overall memory usage(ppm).
	*/
	OverallMemoryUsagePpmUpperBuf []TimeValuePair `json:"overallMemoryUsagePpmUpperBuf,omitempty"`
	/*
	  Power instant consumption (watt)
	*/
	PowerConsumptionInstantWatt []TimeValuePair `json:"powerConsumptionInstantWatt,omitempty"`
	/*
	  Storage capacity(bytes).
	*/
	StorageCapacityBytes []TimeValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Storage usage(bytes).
	*/
	StorageUsageBytes []TimeValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *HostStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HostStatsProjection

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

func (p *HostStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HostStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HostStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "aggregateHypervisorMemoryUsagePpm")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmLowerBuf")
	delete(allFields, "aggregateHypervisorMemoryUsagePpmUpperBuf")
	delete(allFields, "controllerAvgIoLatencyUsecs")
	delete(allFields, "controllerAvgIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecs")
	delete(allFields, "controllerAvgReadIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgReadIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecs")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsLowerBuf")
	delete(allFields, "controllerAvgWriteIoLatencyUsecsUpperBuf")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumIopsLowerBuf")
	delete(allFields, "controllerNumIopsUpperBuf")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumReadIopsLowerBuf")
	delete(allFields, "controllerNumReadIopsUpperBuf")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerNumWriteIopsLowerBuf")
	delete(allFields, "controllerNumWriteIopsUpperBuf")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerReadIoBandwidthKbpsUpperBuf")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoBandwidthKbpsLowerBuf")
	delete(allFields, "controllerWriteIoBandwidthKbpsUpperBuf")
	delete(allFields, "cpuCapacityHz")
	delete(allFields, "cpuUsageHz")
	delete(allFields, "extId")
	delete(allFields, "freePhysicalStorageBytes")
	delete(allFields, "healthCheckScore")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorCpuUsagePpmLowerBuf")
	delete(allFields, "hypervisorCpuUsagePpmUpperBuf")
	delete(allFields, "ioBandwidthKbps")
	delete(allFields, "ioBandwidthKbpsLowerBuf")
	delete(allFields, "ioBandwidthKbpsUpperBuf")
	delete(allFields, "links")
	delete(allFields, "logicalStorageUsageBytes")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "overallMemoryUsageBytes")
	delete(allFields, "overallMemoryUsagePpm")
	delete(allFields, "overallMemoryUsagePpmLowerBuf")
	delete(allFields, "overallMemoryUsagePpmUpperBuf")
	delete(allFields, "powerConsumptionInstantWatt")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHostStatsProjection() *HostStatsProjection {
	p := new(HostStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.HostStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StorageContainerStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the Storage Container.
	*/
	ContainerExtId *string `json:"containerExtId,omitempty"`
	/*
	  Average I/O latency in micro secs.
	*/
	ControllerAvgIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgIoLatencyuSecs,omitempty"`
	/*
	  Average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgReadIoLatencyuSecs,omitempty"`
	/*
	  Average read I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgWriteIoLatencyuSecs,omitempty"`
	/*
	  Total I/O bandwidth - kB per second.
	*/
	ControllerIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerIoBandwidthkBps,omitempty"`
	/*
	  Number of I/O per second.
	*/
	ControllerNumIops []import3.TimeIntValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Number of read I/O per second.
	*/
	ControllerNumReadIops []import3.TimeIntValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Number of write I/O per second.
	*/
	ControllerNumWriteIops []import3.TimeIntValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Read I/O bandwidth kB per second.
	*/
	ControllerReadIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerReadIoBandwidthkBps,omitempty"`
	/*
	  Ratio of read I/O to total I/O in PPM.
	*/
	ControllerReadIoRatioPpm []import3.TimeIntValuePair `json:"controllerReadIoRatioPpm,omitempty"`
	/*
	  Write I/O bandwidth kB per second.
	*/
	ControllerWriteIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerWriteIoBandwidthkBps,omitempty"`
	/*
	  Ratio of read I/O to total I/O in PPM.
	*/
	ControllerWriteIoRatioPpm []import3.TimeIntValuePair `json:"controllerWriteIoRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of the cloning technique.
	*/
	DataReductionCloneSavedBytes []import3.TimeIntValuePair `json:"dataReductionCloneSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the cloning technique.
	*/
	DataReductionCloneSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionCloneSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of compression technique.
	*/
	DataReductionCompressionSavedBytes []import3.TimeIntValuePair `json:"dataReductionCompressionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the compression technique.
	*/
	DataReductionCompressionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionCompressionSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication technique.
	*/
	DataReductionDedupSavedBytes []import3.TimeIntValuePair `json:"dataReductionDedupSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the deduplication technique.
	*/
	DataReductionDedupSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionDedupSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of erasure coding technique.
	*/
	DataReductionErasureCodingSavedBytes []import3.TimeIntValuePair `json:"dataReductionErasureCodingSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the erasure coding technique.
	*/
	DataReductionErasureCodingSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionErasureCodingSavingRatioPpm,omitempty"`
	/*
	  Usage in bytes after reduction of ceduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionOverallPostReductionBytes []import3.TimeIntValuePair `json:"dataReductionOverallPostReductionBytes,omitempty"`
	/*
	  Usage in bytes before reduction of deduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionOverallPreReductionBytes []import3.TimeIntValuePair `json:"dataReductionOverallPreReductionBytes,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication, compression, erasure coding, cloning and thin provisioning technique.
	*/
	DataReductionOverallSavedBytes []import3.TimeIntValuePair `json:"dataReductionOverallSavedBytes,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication, compression, erasure coding technique.
	*/
	DataReductionSavedBytes []import3.TimeIntValuePair `json:"dataReductionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of deduplication, compression and erasure coding.
	*/
	DataReductionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of the snapshot technique.
	*/
	DataReductionSnapshotSavedBytes []import3.TimeIntValuePair `json:"dataReductionSnapshotSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of snapshot technique.
	*/
	DataReductionSnapshotSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionSnapshotSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of thin Provisioning technique.
	*/
	DataReductionThinProvisionSavedBytes []import3.TimeIntValuePair `json:"dataReductionThinProvisionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the thin provisioning technique.
	*/
	DataReductionThinProvisionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionThinProvisionSavingRatioPpm,omitempty"`
	/*
	  Saving ratio in PPM consisting of deduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionTotalSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionTotalSavingRatioPpm,omitempty"`
	/*
	  Total amount of savings in bytes as a result of zero writes.
	*/
	DataReductionZeroWriteSavingsBytes []import3.TimeIntValuePair `json:"dataReductionZeroWriteSavingsBytes,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Health of the Storage Container is represented by an integer value in the range 0-100. A higher value indicates better health.
	*/
	Health []import3.TimeIntValuePair `json:"health,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Actual physical disk usage of the Storage Container without considering for the reservation.
	*/
	StorageActualPhysicalUsageBytes []import3.TimeIntValuePair `json:"storageActualPhysicalUsageBytes,omitempty"`
	/*
	  Storage capacity in bytes.
	*/
	StorageCapacityBytes []import3.TimeIntValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Free storage in bytes.
	*/
	StorageFreeBytes []import3.TimeIntValuePair `json:"storageFreeBytes,omitempty"`
	/*
	  The total explicit reserved physical capacity of other Storage Containers in the same Storage Pool.
	*/
	StorageOtherContainersReservedCapacity []import3.TimeIntValuePair `json:"storageOtherContainersReservedCapacity,omitempty"`
	/*
	  The physical usage outside of the explicitly reserved capacity of other Storage Containers in the same Storage Pool.
	*/
	StorageOtherContainersUnreservedCapacity []import3.TimeIntValuePair `json:"storageOtherContainersUnreservedCapacity,omitempty"`
	/*
	  Replication factor of Storage Container.
	*/
	StorageReplicationFactor []import3.TimeIntValuePair `json:"storageReplicationFactor,omitempty"`
	/*
	  Implicit physical reserved capacity (aggregated at the vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedCapacityBytes []import3.TimeIntValuePair `json:"storageReservedCapacityBytes,omitempty"`
	/*
	  The remaining unused space of the implicit reserved capacity(aggregated on vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedFreeBytes []import3.TimeIntValuePair `json:"storageReservedFreeBytes,omitempty"`
	/*
	  The physical usage of the implicit reserved capacity(aggregated on vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedUsageBytes []import3.TimeIntValuePair `json:"storageReservedUsageBytes,omitempty"`
	/*
	  The space that will be reclaimed if all the snapshots in the cluster is deleted. This is the physical snapshot usage with replication factor and data reduction savings taken into account.
	*/
	StorageSnapshotReclaimable []import3.TimeIntValuePair `json:"storageSnapshotReclaimable,omitempty"`
	/*
	  Total usage on HDD tier for the Storage Container in bytes.
	*/
	StorageTierDasSataUsageBytes []import3.TimeIntValuePair `json:"storageTierDasSataUsageBytes,omitempty"`
	/*
	  Total usage on SDD tier for the Storage Container in bytes.
	*/
	StorageTierSsdUsageBytes []import3.TimeIntValuePair `json:"storageTierSsdUsageBytes,omitempty"`
	/*
	  The physical usage from unreserved vDisks(aggregated on thin provisioning vDisks) in bytes.
	*/
	StorageUnreservedUsageBytes []import3.TimeIntValuePair `json:"storageUnreservedUsageBytes,omitempty"`
	/*
	  Used storage in bytes.
	*/
	StorageUsageBytes []import3.TimeIntValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *StorageContainerStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StorageContainerStats

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

func (p *StorageContainerStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageContainerStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StorageContainerStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "containerExtId")
	delete(allFields, "controllerAvgIoLatencyuSecs")
	delete(allFields, "controllerAvgReadIoLatencyuSecs")
	delete(allFields, "controllerAvgWriteIoLatencyuSecs")
	delete(allFields, "controllerIoBandwidthkBps")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerReadIoBandwidthkBps")
	delete(allFields, "controllerReadIoRatioPpm")
	delete(allFields, "controllerWriteIoBandwidthkBps")
	delete(allFields, "controllerWriteIoRatioPpm")
	delete(allFields, "dataReductionCloneSavedBytes")
	delete(allFields, "dataReductionCloneSavingRatioPpm")
	delete(allFields, "dataReductionCompressionSavedBytes")
	delete(allFields, "dataReductionCompressionSavingRatioPpm")
	delete(allFields, "dataReductionDedupSavedBytes")
	delete(allFields, "dataReductionDedupSavingRatioPpm")
	delete(allFields, "dataReductionErasureCodingSavedBytes")
	delete(allFields, "dataReductionErasureCodingSavingRatioPpm")
	delete(allFields, "dataReductionOverallPostReductionBytes")
	delete(allFields, "dataReductionOverallPreReductionBytes")
	delete(allFields, "dataReductionOverallSavedBytes")
	delete(allFields, "dataReductionSavedBytes")
	delete(allFields, "dataReductionSavingRatioPpm")
	delete(allFields, "dataReductionSnapshotSavedBytes")
	delete(allFields, "dataReductionSnapshotSavingRatioPpm")
	delete(allFields, "dataReductionThinProvisionSavedBytes")
	delete(allFields, "dataReductionThinProvisionSavingRatioPpm")
	delete(allFields, "dataReductionTotalSavingRatioPpm")
	delete(allFields, "dataReductionZeroWriteSavingsBytes")
	delete(allFields, "extId")
	delete(allFields, "health")
	delete(allFields, "links")
	delete(allFields, "storageActualPhysicalUsageBytes")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageFreeBytes")
	delete(allFields, "storageOtherContainersReservedCapacity")
	delete(allFields, "storageOtherContainersUnreservedCapacity")
	delete(allFields, "storageReplicationFactor")
	delete(allFields, "storageReservedCapacityBytes")
	delete(allFields, "storageReservedFreeBytes")
	delete(allFields, "storageReservedUsageBytes")
	delete(allFields, "storageSnapshotReclaimable")
	delete(allFields, "storageTierDasSataUsageBytes")
	delete(allFields, "storageTierSsdUsageBytes")
	delete(allFields, "storageUnreservedUsageBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStorageContainerStats() *StorageContainerStats {
	p := new(StorageContainerStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.StorageContainerStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StorageContainerStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the Storage Container.
	*/
	ContainerExtId *string `json:"containerExtId,omitempty"`
	/*
	  Average I/O latency in micro secs.
	*/
	ControllerAvgIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgIoLatencyuSecs,omitempty"`
	/*
	  Average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgReadIoLatencyuSecs,omitempty"`
	/*
	  Average read I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyuSecs []import3.TimeIntValuePair `json:"controllerAvgWriteIoLatencyuSecs,omitempty"`
	/*
	  Total I/O bandwidth - kB per second.
	*/
	ControllerIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerIoBandwidthkBps,omitempty"`
	/*
	  Number of I/O per second.
	*/
	ControllerNumIops []import3.TimeIntValuePair `json:"controllerNumIops,omitempty"`
	/*
	  Number of read I/O per second.
	*/
	ControllerNumReadIops []import3.TimeIntValuePair `json:"controllerNumReadIops,omitempty"`
	/*
	  Number of write I/O per second.
	*/
	ControllerNumWriteIops []import3.TimeIntValuePair `json:"controllerNumWriteIops,omitempty"`
	/*
	  Read I/O bandwidth kB per second.
	*/
	ControllerReadIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerReadIoBandwidthkBps,omitempty"`
	/*
	  Ratio of read I/O to total I/O in PPM.
	*/
	ControllerReadIoRatioPpm []import3.TimeIntValuePair `json:"controllerReadIoRatioPpm,omitempty"`
	/*
	  Write I/O bandwidth kB per second.
	*/
	ControllerWriteIoBandwidthkBps []import3.TimeIntValuePair `json:"controllerWriteIoBandwidthkBps,omitempty"`
	/*
	  Ratio of read I/O to total I/O in PPM.
	*/
	ControllerWriteIoRatioPpm []import3.TimeIntValuePair `json:"controllerWriteIoRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of the cloning technique.
	*/
	DataReductionCloneSavedBytes []import3.TimeIntValuePair `json:"dataReductionCloneSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the cloning technique.
	*/
	DataReductionCloneSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionCloneSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of compression technique.
	*/
	DataReductionCompressionSavedBytes []import3.TimeIntValuePair `json:"dataReductionCompressionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the compression technique.
	*/
	DataReductionCompressionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionCompressionSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication technique.
	*/
	DataReductionDedupSavedBytes []import3.TimeIntValuePair `json:"dataReductionDedupSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the deduplication technique.
	*/
	DataReductionDedupSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionDedupSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of erasure coding technique.
	*/
	DataReductionErasureCodingSavedBytes []import3.TimeIntValuePair `json:"dataReductionErasureCodingSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the erasure coding technique.
	*/
	DataReductionErasureCodingSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionErasureCodingSavingRatioPpm,omitempty"`
	/*
	  Usage in bytes after reduction of ceduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionOverallPostReductionBytes []import3.TimeIntValuePair `json:"dataReductionOverallPostReductionBytes,omitempty"`
	/*
	  Usage in bytes before reduction of deduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionOverallPreReductionBytes []import3.TimeIntValuePair `json:"dataReductionOverallPreReductionBytes,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication, compression, erasure coding, cloning and thin provisioning technique.
	*/
	DataReductionOverallSavedBytes []import3.TimeIntValuePair `json:"dataReductionOverallSavedBytes,omitempty"`
	/*
	  Storage saving in bytes as a result of deduplication, compression, erasure coding technique.
	*/
	DataReductionSavedBytes []import3.TimeIntValuePair `json:"dataReductionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of deduplication, compression and erasure coding.
	*/
	DataReductionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of the snapshot technique.
	*/
	DataReductionSnapshotSavedBytes []import3.TimeIntValuePair `json:"dataReductionSnapshotSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of snapshot technique.
	*/
	DataReductionSnapshotSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionSnapshotSavingRatioPpm,omitempty"`
	/*
	  Storage saving in bytes as a result of thin Provisioning technique.
	*/
	DataReductionThinProvisionSavedBytes []import3.TimeIntValuePair `json:"dataReductionThinProvisionSavedBytes,omitempty"`
	/*
	  Saving ratio in PPM as a result of the thin provisioning technique.
	*/
	DataReductionThinProvisionSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionThinProvisionSavingRatioPpm,omitempty"`
	/*
	  Saving ratio in PPM consisting of deduplication, compression, erasure coding, cloning, and thin provisioning.
	*/
	DataReductionTotalSavingRatioPpm []import3.TimeIntValuePair `json:"dataReductionTotalSavingRatioPpm,omitempty"`
	/*
	  Total amount of savings in bytes as a result of zero writes.
	*/
	DataReductionZeroWriteSavingsBytes []import3.TimeIntValuePair `json:"dataReductionZeroWriteSavingsBytes,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Health of the Storage Container is represented by an integer value in the range 0-100. A higher value indicates better health.
	*/
	Health []import3.TimeIntValuePair `json:"health,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Actual physical disk usage of the Storage Container without considering for the reservation.
	*/
	StorageActualPhysicalUsageBytes []import3.TimeIntValuePair `json:"storageActualPhysicalUsageBytes,omitempty"`
	/*
	  Storage capacity in bytes.
	*/
	StorageCapacityBytes []import3.TimeIntValuePair `json:"storageCapacityBytes,omitempty"`
	/*
	  Free storage in bytes.
	*/
	StorageFreeBytes []import3.TimeIntValuePair `json:"storageFreeBytes,omitempty"`
	/*
	  The total explicit reserved physical capacity of other Storage Containers in the same Storage Pool.
	*/
	StorageOtherContainersReservedCapacity []import3.TimeIntValuePair `json:"storageOtherContainersReservedCapacity,omitempty"`
	/*
	  The physical usage outside of the explicitly reserved capacity of other Storage Containers in the same Storage Pool.
	*/
	StorageOtherContainersUnreservedCapacity []import3.TimeIntValuePair `json:"storageOtherContainersUnreservedCapacity,omitempty"`
	/*
	  Replication factor of Storage Container.
	*/
	StorageReplicationFactor []import3.TimeIntValuePair `json:"storageReplicationFactor,omitempty"`
	/*
	  Implicit physical reserved capacity (aggregated at the vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedCapacityBytes []import3.TimeIntValuePair `json:"storageReservedCapacityBytes,omitempty"`
	/*
	  The remaining unused space of the implicit reserved capacity(aggregated on vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedFreeBytes []import3.TimeIntValuePair `json:"storageReservedFreeBytes,omitempty"`
	/*
	  The physical usage of the implicit reserved capacity(aggregated on vDisk level due to thick provisioning) in bytes.
	*/
	StorageReservedUsageBytes []import3.TimeIntValuePair `json:"storageReservedUsageBytes,omitempty"`
	/*
	  The space that will be reclaimed if all the snapshots in the cluster is deleted. This is the physical snapshot usage with replication factor and data reduction savings taken into account.
	*/
	StorageSnapshotReclaimable []import3.TimeIntValuePair `json:"storageSnapshotReclaimable,omitempty"`
	/*
	  Total usage on HDD tier for the Storage Container in bytes.
	*/
	StorageTierDasSataUsageBytes []import3.TimeIntValuePair `json:"storageTierDasSataUsageBytes,omitempty"`
	/*
	  Total usage on SDD tier for the Storage Container in bytes.
	*/
	StorageTierSsdUsageBytes []import3.TimeIntValuePair `json:"storageTierSsdUsageBytes,omitempty"`
	/*
	  The physical usage from unreserved vDisks(aggregated on thin provisioning vDisks) in bytes.
	*/
	StorageUnreservedUsageBytes []import3.TimeIntValuePair `json:"storageUnreservedUsageBytes,omitempty"`
	/*
	  Used storage in bytes.
	*/
	StorageUsageBytes []import3.TimeIntValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *StorageContainerStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StorageContainerStatsProjection

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

func (p *StorageContainerStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageContainerStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = StorageContainerStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "containerExtId")
	delete(allFields, "controllerAvgIoLatencyuSecs")
	delete(allFields, "controllerAvgReadIoLatencyuSecs")
	delete(allFields, "controllerAvgWriteIoLatencyuSecs")
	delete(allFields, "controllerIoBandwidthkBps")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerReadIoBandwidthkBps")
	delete(allFields, "controllerReadIoRatioPpm")
	delete(allFields, "controllerWriteIoBandwidthkBps")
	delete(allFields, "controllerWriteIoRatioPpm")
	delete(allFields, "dataReductionCloneSavedBytes")
	delete(allFields, "dataReductionCloneSavingRatioPpm")
	delete(allFields, "dataReductionCompressionSavedBytes")
	delete(allFields, "dataReductionCompressionSavingRatioPpm")
	delete(allFields, "dataReductionDedupSavedBytes")
	delete(allFields, "dataReductionDedupSavingRatioPpm")
	delete(allFields, "dataReductionErasureCodingSavedBytes")
	delete(allFields, "dataReductionErasureCodingSavingRatioPpm")
	delete(allFields, "dataReductionOverallPostReductionBytes")
	delete(allFields, "dataReductionOverallPreReductionBytes")
	delete(allFields, "dataReductionOverallSavedBytes")
	delete(allFields, "dataReductionSavedBytes")
	delete(allFields, "dataReductionSavingRatioPpm")
	delete(allFields, "dataReductionSnapshotSavedBytes")
	delete(allFields, "dataReductionSnapshotSavingRatioPpm")
	delete(allFields, "dataReductionThinProvisionSavedBytes")
	delete(allFields, "dataReductionThinProvisionSavingRatioPpm")
	delete(allFields, "dataReductionTotalSavingRatioPpm")
	delete(allFields, "dataReductionZeroWriteSavingsBytes")
	delete(allFields, "extId")
	delete(allFields, "health")
	delete(allFields, "links")
	delete(allFields, "storageActualPhysicalUsageBytes")
	delete(allFields, "storageCapacityBytes")
	delete(allFields, "storageFreeBytes")
	delete(allFields, "storageOtherContainersReservedCapacity")
	delete(allFields, "storageOtherContainersUnreservedCapacity")
	delete(allFields, "storageReplicationFactor")
	delete(allFields, "storageReservedCapacityBytes")
	delete(allFields, "storageReservedFreeBytes")
	delete(allFields, "storageReservedUsageBytes")
	delete(allFields, "storageSnapshotReclaimable")
	delete(allFields, "storageTierDasSataUsageBytes")
	delete(allFields, "storageTierSsdUsageBytes")
	delete(allFields, "storageUnreservedUsageBytes")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewStorageContainerStatsProjection() *StorageContainerStatsProjection {
	p := new(StorageContainerStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.StorageContainerStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Time - Value pair for time-series stat attributes.
*/
type TimeValuePair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Timestamp for given stat attribute(in ISO-8601 format).
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  Value of stat at given timestamp.
	*/
	Value *int64 `json:"value,omitempty"`
}

func (p *TimeValuePair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TimeValuePair

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

func (p *TimeValuePair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TimeValuePair(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTimeValuePair() *TimeValuePair {
	p := new(TimeValuePair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.stats.TimeValuePair"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfHostStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *HostStats             `json:"-"`
}

func NewOneOfHostStatsApiResponseData() *OneOfHostStatsApiResponseData {
	p := new(OneOfHostStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHostStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHostStatsApiResponseData is nil"))
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
	case HostStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(HostStats)
		}
		*p.oneOfType2001 = v.(HostStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHostStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfHostStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(HostStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "clustermgmt.v4.stats.HostStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(HostStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHostStatsApiResponseData"))
}

func (p *OneOfHostStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfHostStatsApiResponseData")
}

type OneOfClusterStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *ClusterStats          `json:"-"`
}

func NewOneOfClusterStatsApiResponseData() *OneOfClusterStatsApiResponseData {
	p := new(OneOfClusterStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterStatsApiResponseData is nil"))
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
	case ClusterStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ClusterStats)
		}
		*p.oneOfType2001 = v.(ClusterStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfClusterStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfClusterStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(ClusterStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "clustermgmt.v4.stats.ClusterStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ClusterStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterStatsApiResponseData"))
}

func (p *OneOfClusterStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfClusterStatsApiResponseData")
}

type OneOfGetStorageContainerStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *StorageContainerStats `json:"-"`
}

func NewOneOfGetStorageContainerStatsApiResponseData() *OneOfGetStorageContainerStatsApiResponseData {
	p := new(OneOfGetStorageContainerStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetStorageContainerStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetStorageContainerStatsApiResponseData is nil"))
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
	case StorageContainerStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(StorageContainerStats)
		}
		*p.oneOfType2001 = v.(StorageContainerStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetStorageContainerStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetStorageContainerStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(StorageContainerStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "clustermgmt.v4.stats.StorageContainerStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(StorageContainerStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetStorageContainerStatsApiResponseData"))
}

func (p *OneOfGetStorageContainerStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetStorageContainerStatsApiResponseData")
}

type OneOfGetDiskStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *DiskStats             `json:"-"`
}

func NewOneOfGetDiskStatsApiResponseData() *OneOfGetDiskStatsApiResponseData {
	p := new(OneOfGetDiskStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDiskStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDiskStatsApiResponseData is nil"))
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
	case DiskStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(DiskStats)
		}
		*p.oneOfType2001 = v.(DiskStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetDiskStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetDiskStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(DiskStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "clustermgmt.v4.stats.DiskStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(DiskStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDiskStatsApiResponseData"))
}

func (p *OneOfGetDiskStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetDiskStatsApiResponseData")
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
