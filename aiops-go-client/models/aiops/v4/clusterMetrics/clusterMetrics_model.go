/*
 * Generated file models/aiops/v4/clusterMetrics/clusterMetrics_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix AIOps APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  View metrics across clusters
*/
package clusterMetrics

import (
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

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
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

func NewClusterMetrics() *ClusterMetrics {
	p := new(ClusterMetrics)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterMetrics"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewClusterProjection() *ClusterProjection {
	p := new(ClusterProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewNode() *Node {
	p := new(Node)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Node"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewNodeProjection() *NodeProjection {
	p := new(NodeProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.NodeProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.Vm"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewVmProjection() *VmProjection {
	p := new(VmProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "aiops.v4.clusterMetrics.VmProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
