/*
 * Generated file models/aiops/v4/clusterMetrics/clusterMetrics_model.go.
 *
 * Product version: 4.2.1-beta-1
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  View metrics across clusters
*/
package clusterMetrics

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
)

type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapacityBytes *string `json:"capacityBytes,omitempty"`

	ExternalIpAddress *string `json:"externalIpAddress,omitempty"`

	FreeBytes *string `json:"freeBytes,omitempty"`

	HypervisorCpuUsagePpm *string `json:"hypervisorCpuUsagePpm,omitempty"`

	HypervisorMemoryUsagePpm *string `json:"hypervisorMemoryUsagePpm,omitempty"`

	MemoryCapacityBytes *string `json:"memoryCapacityBytes,omitempty"`

	Name *string `json:"name,omitempty"`

	NumCpus *string `json:"numCpus,omitempty"`

	Rf *string `json:"rf,omitempty"`

	SavedBytes *string `json:"savedBytes,omitempty"`

	UsageBytes *string `json:"usageBytes,omitempty"`

	Uuid *string `json:"uuid,omitempty"`
}

func (p *Cluster) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Cluster

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

func (p *Cluster) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Cluster
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCluster()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapacityBytes != nil {
		p.CapacityBytes = known.CapacityBytes
	}
	if known.ExternalIpAddress != nil {
		p.ExternalIpAddress = known.ExternalIpAddress
	}
	if known.FreeBytes != nil {
		p.FreeBytes = known.FreeBytes
	}
	if known.HypervisorCpuUsagePpm != nil {
		p.HypervisorCpuUsagePpm = known.HypervisorCpuUsagePpm
	}
	if known.HypervisorMemoryUsagePpm != nil {
		p.HypervisorMemoryUsagePpm = known.HypervisorMemoryUsagePpm
	}
	if known.MemoryCapacityBytes != nil {
		p.MemoryCapacityBytes = known.MemoryCapacityBytes
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NumCpus != nil {
		p.NumCpus = known.NumCpus
	}
	if known.Rf != nil {
		p.Rf = known.Rf
	}
	if known.SavedBytes != nil {
		p.SavedBytes = known.SavedBytes
	}
	if known.UsageBytes != nil {
		p.UsageBytes = known.UsageBytes
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityBytes")
	delete(allFields, "externalIpAddress")
	delete(allFields, "freeBytes")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorMemoryUsagePpm")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "name")
	delete(allFields, "numCpus")
	delete(allFields, "rf")
	delete(allFields, "savedBytes")
	delete(allFields, "usageBytes")
	delete(allFields, "uuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterMetrics struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Total amount of available memory in bytes for deploying new VMs
	*/
	AvailableMemorySizeBytes *int64 `json:"availableMemorySizeBytes,omitempty"`
	/*
	  Total number of available VMs for deploying new VMs
	*/
	AvailableVCpuCount *int64 `json:"availableVCpuCount,omitempty"`
	/*
	  Name of the cluster
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  Total memory size in bytes currently being used by the active VMs
	*/
	CurrentMemoryUsageSizeBytes *int64 `json:"currentMemoryUsageSizeBytes,omitempty"`
	/*
	  Total amount of CPU in hertz currently being used by the active VMs
	*/
	CurrentUsedCpuHz *int64 `json:"currentUsedCpuHz,omitempty"`
	/*
	  Total number of vCPUs currently being used by the active VMs
	*/
	CurrentUsedVCpuCount *int64 `json:"currentUsedVCpuCount,omitempty"`
	/*
	  CVM IPs of the nodes in the cluster
	*/
	CvmIPs []import1.IPv4Address `json:"cvmIPs,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalIP *import1.IPv4Address `json:"externalIP,omitempty"`
	/*
	  Powered OFF VM with the largest memory size
	*/
	LargestPoweredOffVMMemory *int64 `json:"largestPoweredOffVMMemory,omitempty"`
	/*
	  Largest powered OFF VM in terms of vCPU count
	*/
	LargestPoweredOffVMVCpu *int64 `json:"largestPoweredOffVMVCpu,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Largest possible memory size in bytes for the next VM that can be added
	*/
	MaxVmMemorySizeBytes *int64 `json:"maxVmMemorySizeBytes,omitempty"`
	/*
	  Largest possible VM size in vCPUs for the next VM that can be added
	*/
	MaxVmVCpuSize *int64 `json:"maxVmVCpuSize,omitempty"`
	/*
	  Total amount of CPU in hertz provisioned for all the VMs
	*/
	ProvisionedCpuHz *int64 `json:"provisionedCpuHz,omitempty"`
	/*
	  Total memory size in bytes provisioned for all the VMs
	*/
	ProvisionedMemoryUsageBytes *int64 `json:"provisionedMemoryUsageBytes,omitempty"`
	/*
	  Total number of vCPUs provisioned for all the VMs
	*/
	ProvisionedVCpuCount *int64 `json:"provisionedVCpuCount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Total amount of CPU in hertz in the cluster
	*/
	TotalCpuHz *int64 `json:"totalCpuHz,omitempty"`
	/*
	  Total memory size in bytes in the cluster
	*/
	TotalMemorySizeBytes *int64 `json:"totalMemorySizeBytes,omitempty"`
	/*
	  Total raw storage bytes still available in the cluster
	*/
	TotalStorageAvailableBytes *int64 `json:"totalStorageAvailableBytes,omitempty"`
	/*
	  Total logical storage bytes still available in the cluster
	*/
	TotalStorageAvailableLogicalBytes *int64 `json:"totalStorageAvailableLogicalBytes,omitempty"`
	/*
	  Total logical storage in bytes capacity after considering RF
	*/
	TotalStorageCapacityLogicalBytes *int64 `json:"totalStorageCapacityLogicalBytes,omitempty"`
	/*
	  Total logical storage capacity in bytes after accounting for savings from dedup, compressions,and erasure coding
	*/
	TotalStorageCapacityLogicalWithSavingsBytes *int64 `json:"totalStorageCapacityLogicalWithSavingsBytes,omitempty"`
	/*
	  Total storage in bytes capacity in the cluster
	*/
	TotalStorageCapacityRawBytes *int64 `json:"totalStorageCapacityRawBytes,omitempty"`
	/*
	  Total raw storage usage by the cluster
	*/
	TotalStorageUsageBytes *int64 `json:"totalStorageUsageBytes,omitempty"`
	/*
	  Total number of vCPUs in the cluster
	*/
	TotalVCpuCount *int64 `json:"totalVCpuCount,omitempty"`
}

func (p *ClusterMetrics) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterMetrics

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

func (p *ClusterMetrics) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterMetrics
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterMetrics()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AvailableMemorySizeBytes != nil {
		p.AvailableMemorySizeBytes = known.AvailableMemorySizeBytes
	}
	if known.AvailableVCpuCount != nil {
		p.AvailableVCpuCount = known.AvailableVCpuCount
	}
	if known.ClusterName != nil {
		p.ClusterName = known.ClusterName
	}
	if known.CurrentMemoryUsageSizeBytes != nil {
		p.CurrentMemoryUsageSizeBytes = known.CurrentMemoryUsageSizeBytes
	}
	if known.CurrentUsedCpuHz != nil {
		p.CurrentUsedCpuHz = known.CurrentUsedCpuHz
	}
	if known.CurrentUsedVCpuCount != nil {
		p.CurrentUsedVCpuCount = known.CurrentUsedVCpuCount
	}
	if known.CvmIPs != nil {
		p.CvmIPs = known.CvmIPs
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ExternalIP != nil {
		p.ExternalIP = known.ExternalIP
	}
	if known.LargestPoweredOffVMMemory != nil {
		p.LargestPoweredOffVMMemory = known.LargestPoweredOffVMMemory
	}
	if known.LargestPoweredOffVMVCpu != nil {
		p.LargestPoweredOffVMVCpu = known.LargestPoweredOffVMVCpu
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.MaxVmMemorySizeBytes != nil {
		p.MaxVmMemorySizeBytes = known.MaxVmMemorySizeBytes
	}
	if known.MaxVmVCpuSize != nil {
		p.MaxVmVCpuSize = known.MaxVmVCpuSize
	}
	if known.ProvisionedCpuHz != nil {
		p.ProvisionedCpuHz = known.ProvisionedCpuHz
	}
	if known.ProvisionedMemoryUsageBytes != nil {
		p.ProvisionedMemoryUsageBytes = known.ProvisionedMemoryUsageBytes
	}
	if known.ProvisionedVCpuCount != nil {
		p.ProvisionedVCpuCount = known.ProvisionedVCpuCount
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalCpuHz != nil {
		p.TotalCpuHz = known.TotalCpuHz
	}
	if known.TotalMemorySizeBytes != nil {
		p.TotalMemorySizeBytes = known.TotalMemorySizeBytes
	}
	if known.TotalStorageAvailableBytes != nil {
		p.TotalStorageAvailableBytes = known.TotalStorageAvailableBytes
	}
	if known.TotalStorageAvailableLogicalBytes != nil {
		p.TotalStorageAvailableLogicalBytes = known.TotalStorageAvailableLogicalBytes
	}
	if known.TotalStorageCapacityLogicalBytes != nil {
		p.TotalStorageCapacityLogicalBytes = known.TotalStorageCapacityLogicalBytes
	}
	if known.TotalStorageCapacityLogicalWithSavingsBytes != nil {
		p.TotalStorageCapacityLogicalWithSavingsBytes = known.TotalStorageCapacityLogicalWithSavingsBytes
	}
	if known.TotalStorageCapacityRawBytes != nil {
		p.TotalStorageCapacityRawBytes = known.TotalStorageCapacityRawBytes
	}
	if known.TotalStorageUsageBytes != nil {
		p.TotalStorageUsageBytes = known.TotalStorageUsageBytes
	}
	if known.TotalVCpuCount != nil {
		p.TotalVCpuCount = known.TotalVCpuCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "availableMemorySizeBytes")
	delete(allFields, "availableVCpuCount")
	delete(allFields, "clusterName")
	delete(allFields, "currentMemoryUsageSizeBytes")
	delete(allFields, "currentUsedCpuHz")
	delete(allFields, "currentUsedVCpuCount")
	delete(allFields, "cvmIPs")
	delete(allFields, "extId")
	delete(allFields, "externalIP")
	delete(allFields, "largestPoweredOffVMMemory")
	delete(allFields, "largestPoweredOffVMVCpu")
	delete(allFields, "links")
	delete(allFields, "maxVmMemorySizeBytes")
	delete(allFields, "maxVmVCpuSize")
	delete(allFields, "provisionedCpuHz")
	delete(allFields, "provisionedMemoryUsageBytes")
	delete(allFields, "provisionedVCpuCount")
	delete(allFields, "tenantId")
	delete(allFields, "totalCpuHz")
	delete(allFields, "totalMemorySizeBytes")
	delete(allFields, "totalStorageAvailableBytes")
	delete(allFields, "totalStorageAvailableLogicalBytes")
	delete(allFields, "totalStorageCapacityLogicalBytes")
	delete(allFields, "totalStorageCapacityLogicalWithSavingsBytes")
	delete(allFields, "totalStorageCapacityRawBytes")
	delete(allFields, "totalStorageUsageBytes")
	delete(allFields, "totalVCpuCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterMetrics() *ClusterMetrics {
	p := new(ClusterMetrics)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterMetrics"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapacityBytes *string `json:"capacityBytes,omitempty"`

	ExternalIpAddress *string `json:"externalIpAddress,omitempty"`

	FreeBytes *string `json:"freeBytes,omitempty"`

	HypervisorCpuUsagePpm *string `json:"hypervisorCpuUsagePpm,omitempty"`

	HypervisorMemoryUsagePpm *string `json:"hypervisorMemoryUsagePpm,omitempty"`

	MemoryCapacityBytes *string `json:"memoryCapacityBytes,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeProjection []NodeProjection `json:"nodeProjection,omitempty"`

	NumCpus *string `json:"numCpus,omitempty"`

	Rf *string `json:"rf,omitempty"`

	SavedBytes *string `json:"savedBytes,omitempty"`

	UsageBytes *string `json:"usageBytes,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	VmProjection []VmProjection `json:"vmProjection,omitempty"`
}

func (p *ClusterProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterProjection

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

func (p *ClusterProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapacityBytes != nil {
		p.CapacityBytes = known.CapacityBytes
	}
	if known.ExternalIpAddress != nil {
		p.ExternalIpAddress = known.ExternalIpAddress
	}
	if known.FreeBytes != nil {
		p.FreeBytes = known.FreeBytes
	}
	if known.HypervisorCpuUsagePpm != nil {
		p.HypervisorCpuUsagePpm = known.HypervisorCpuUsagePpm
	}
	if known.HypervisorMemoryUsagePpm != nil {
		p.HypervisorMemoryUsagePpm = known.HypervisorMemoryUsagePpm
	}
	if known.MemoryCapacityBytes != nil {
		p.MemoryCapacityBytes = known.MemoryCapacityBytes
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NodeProjection != nil {
		p.NodeProjection = known.NodeProjection
	}
	if known.NumCpus != nil {
		p.NumCpus = known.NumCpus
	}
	if known.Rf != nil {
		p.Rf = known.Rf
	}
	if known.SavedBytes != nil {
		p.SavedBytes = known.SavedBytes
	}
	if known.UsageBytes != nil {
		p.UsageBytes = known.UsageBytes
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.VmProjection != nil {
		p.VmProjection = known.VmProjection
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityBytes")
	delete(allFields, "externalIpAddress")
	delete(allFields, "freeBytes")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorMemoryUsagePpm")
	delete(allFields, "memoryCapacityBytes")
	delete(allFields, "name")
	delete(allFields, "nodeProjection")
	delete(allFields, "numCpus")
	delete(allFields, "rf")
	delete(allFields, "savedBytes")
	delete(allFields, "usageBytes")
	delete(allFields, "uuid")
	delete(allFields, "vmProjection")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterProjection() *ClusterProjection {
	p := new(ClusterProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Node struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapacityHz *string `json:"capacityHz,omitempty"`

	HaMemoryReservedBytes *string `json:"haMemoryReservedBytes,omitempty"`

	MemorySizeBytes *string `json:"memorySizeBytes,omitempty"`

	NodeUuid *string `json:"nodeUuid,omitempty"`

	NumCpuThreads *string `json:"numCpuThreads,omitempty"`
}

func (p *Node) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Node

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

func (p *Node) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Node
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNode()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapacityHz != nil {
		p.CapacityHz = known.CapacityHz
	}
	if known.HaMemoryReservedBytes != nil {
		p.HaMemoryReservedBytes = known.HaMemoryReservedBytes
	}
	if known.MemorySizeBytes != nil {
		p.MemorySizeBytes = known.MemorySizeBytes
	}
	if known.NodeUuid != nil {
		p.NodeUuid = known.NodeUuid
	}
	if known.NumCpuThreads != nil {
		p.NumCpuThreads = known.NumCpuThreads
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityHz")
	delete(allFields, "haMemoryReservedBytes")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "nodeUuid")
	delete(allFields, "numCpuThreads")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNode() *Node {
	p := new(Node)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Node"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type NodeProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CapacityHz *string `json:"capacityHz,omitempty"`

	HaMemoryReservedBytes *string `json:"haMemoryReservedBytes,omitempty"`

	MemorySizeBytes *string `json:"memorySizeBytes,omitempty"`

	NodeUuid *string `json:"nodeUuid,omitempty"`

	NumCpuThreads *string `json:"numCpuThreads,omitempty"`
}

func (p *NodeProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NodeProjection

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

func (p *NodeProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CapacityHz != nil {
		p.CapacityHz = known.CapacityHz
	}
	if known.HaMemoryReservedBytes != nil {
		p.HaMemoryReservedBytes = known.HaMemoryReservedBytes
	}
	if known.MemorySizeBytes != nil {
		p.MemorySizeBytes = known.MemorySizeBytes
	}
	if known.NodeUuid != nil {
		p.NodeUuid = known.NodeUuid
	}
	if known.NumCpuThreads != nil {
		p.NumCpuThreads = known.NumCpuThreads
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "capacityHz")
	delete(allFields, "haMemoryReservedBytes")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "nodeUuid")
	delete(allFields, "numCpuThreads")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeProjection() *NodeProjection {
	p := new(NodeProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.NodeProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Node *string `json:"node,omitempty"`

	NumVcpus *string `json:"numVcpus,omitempty"`

	PowerState *string `json:"powerState,omitempty"`

	VmMemory *string `json:"vmMemory,omitempty"`

	VmUuid *string `json:"vmUuid,omitempty"`
}

func (p *Vm) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Vm

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

func (p *Vm) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Vm
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVm()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Node != nil {
		p.Node = known.Node
	}
	if known.NumVcpus != nil {
		p.NumVcpus = known.NumVcpus
	}
	if known.PowerState != nil {
		p.PowerState = known.PowerState
	}
	if known.VmMemory != nil {
		p.VmMemory = known.VmMemory
	}
	if known.VmUuid != nil {
		p.VmUuid = known.VmUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "node")
	delete(allFields, "numVcpus")
	delete(allFields, "powerState")
	delete(allFields, "vmMemory")
	delete(allFields, "vmUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Vm"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Node *string `json:"node,omitempty"`

	NumVcpus *string `json:"numVcpus,omitempty"`

	PowerState *string `json:"powerState,omitempty"`

	VmMemory *string `json:"vmMemory,omitempty"`

	VmUuid *string `json:"vmUuid,omitempty"`
}

func (p *VmProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmProjection

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

func (p *VmProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewVmProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Node != nil {
		p.Node = known.Node
	}
	if known.NumVcpus != nil {
		p.NumVcpus = known.NumVcpus
	}
	if known.PowerState != nil {
		p.PowerState = known.PowerState
	}
	if known.VmMemory != nil {
		p.VmMemory = known.VmMemory
	}
	if known.VmUuid != nil {
		p.VmUuid = known.VmUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "node")
	delete(allFields, "numVcpus")
	delete(allFields, "powerState")
	delete(allFields, "vmMemory")
	delete(allFields, "vmUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVmProjection() *VmProjection {
	p := new(VmProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.VmProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
