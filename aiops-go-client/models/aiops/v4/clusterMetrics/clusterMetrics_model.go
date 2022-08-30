/*
 * Generated file models/aiops/v4/clusterMetrics/clusterMetrics_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Aiops Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  View metrics across clusters
*/
package clusterMetrics
import (
  import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/error"
  import3 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/config"
  import2 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
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
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.Cluster"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /aiops/v4.0.a1/cluster-metrics/{extId} Get operation
*/
type ClusterApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfClusterApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewClusterApiResponse() *ClusterApiResponse {
  p := new(ClusterApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.ClusterApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ClusterApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ClusterApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfClusterApiResponseData()
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
REST response for all response codes in api path /aiops/v4.0.a1/cluster-metrics Get operation
*/
type ClusterListApiResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfClusterListApiResponseData `json:"data,omitempty"`
  
  Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewClusterListApiResponse() *ClusterListApiResponse {
  p := new(ClusterListApiResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterListApiResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.ClusterListApiResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *ClusterListApiResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *ClusterListApiResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfClusterListApiResponseData()
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




type ClusterMetrics struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Total amount of available memory in bytes for deploying new VMs
  */
  AvailableMemorySizeBytes *int64 `json:"availableMemorySizeBytes,omitempty"`
  /**
  Total number of available VMs for deploying new VMs
  */
  AvailableVCpuCount *int64 `json:"availableVCpuCount,omitempty"`
  /**
  Name of the cluster
  */
  ClusterName *string `json:"clusterName,omitempty"`
  /**
  Total memory size in bytes currently being used by the active VMs
  */
  CurrentMemoryUsageSizeBytes *int64 `json:"currentMemoryUsageSizeBytes,omitempty"`
  /**
  Total amount of CPU in hertz currently being used by the active VMs
  */
  CurrentUsedCpuHz *int64 `json:"currentUsedCpuHz,omitempty"`
  /**
  Total number of vCPUs currently being used by the active VMs
  */
  CurrentUsedVCpuCount *int64 `json:"currentUsedVCpuCount,omitempty"`
  /**
  CVM IPs of the nodes in the cluster
  */
  CvmIPs []import3.IPv4Address `json:"cvmIPs,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  
  ExternalIP *import3.IPv4Address `json:"externalIP,omitempty"`
  /**
  Powered OFF VM with the largest memory size
  */
  LargestPoweredOffVMMemory *int64 `json:"largestPoweredOffVMMemory,omitempty"`
  /**
  Largest powered OFF VM in terms of vCPU count
  */
  LargestPoweredOffVMVCpu *int64 `json:"largestPoweredOffVMVCpu,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import2.ApiLink `json:"links,omitempty"`
  /**
  Largest possible memory size in bytes for the next VM that can be added
  */
  MaxVmMemorySizeBytes *int64 `json:"maxVmMemorySizeBytes,omitempty"`
  /**
  Largest possible VM size in vCPUs for the next VM that can be added
  */
  MaxVmVCpuSize *int64 `json:"maxVmVCpuSize,omitempty"`
  /**
  Total amount of CPU in hertz provisioned for all the VMs
  */
  ProvisionedCpuHz *int64 `json:"provisionedCpuHz,omitempty"`
  /**
  Total memory size in bytes provisioned for all the VMs
  */
  ProvisionedMemoryUsageBytes *int64 `json:"provisionedMemoryUsageBytes,omitempty"`
  /**
  Total number of vCPUs provisioned for all the VMs
  */
  ProvisionedVCpuCount *int64 `json:"provisionedVCpuCount,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
  /**
  Total amount of CPU in hertz in the cluster
  */
  TotalCpuHz *int64 `json:"totalCpuHz,omitempty"`
  /**
  Total memory size in bytes in the cluster
  */
  TotalMemorySizeBytes *int64 `json:"totalMemorySizeBytes,omitempty"`
  /**
  Total raw storage bytes still available in the cluster
  */
  TotalStorageAvailableBytes *int64 `json:"totalStorageAvailableBytes,omitempty"`
  /**
  Total logical storage bytes still available in the cluster
  */
  TotalStorageAvailableLogicalBytes *int64 `json:"totalStorageAvailableLogicalBytes,omitempty"`
  /**
  Total logical storage in bytes capacity after considering RF
  */
  TotalStorageCapacityLogicalBytes *int64 `json:"totalStorageCapacityLogicalBytes,omitempty"`
  /**
  Total logical storage capacity in bytes after accounting for savings from dedup, compressions,and erasure coding
  */
  TotalStorageCapacityLogicalWithSavingsBytes *int64 `json:"totalStorageCapacityLogicalWithSavingsBytes,omitempty"`
  /**
  Total storage in bytes capacity in the cluster
  */
  TotalStorageCapacityRawBytes *int64 `json:"totalStorageCapacityRawBytes,omitempty"`
  /**
  Total raw storage usage by the cluster
  */
  TotalStorageUsageBytes *int64 `json:"totalStorageUsageBytes,omitempty"`
  /**
  Total number of vCPUs in the cluster
  */
  TotalVCpuCount *int64 `json:"totalVCpuCount,omitempty"`
}


func NewClusterMetrics() *ClusterMetrics {
  p := new(ClusterMetrics)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterMetrics"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.ClusterMetrics"}
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
  
  NodeProjection *NodeProjection `json:"nodeProjection,omitempty"`
  
  NumCpus *string `json:"numCpus,omitempty"`
  
  Rf *string `json:"rf,omitempty"`
  
  SavedBytes *string `json:"savedBytes,omitempty"`
  
  UsageBytes *string `json:"usageBytes,omitempty"`
  
  Uuid *string `json:"uuid,omitempty"`
  
  VmProjection *VmProjection `json:"vmProjection,omitempty"`
}


func NewClusterProjection() *ClusterProjection {
  p := new(ClusterProjection)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "aiops.v4.clusterMetrics.ClusterProjection"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.ClusterProjection"}
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
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.Node"}
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
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.NodeProjection"}
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
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.Vm"}
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
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "aiops.v4.r0.a1.clusterMetrics.VmProjection"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfClusterApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *ClusterMetrics `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfClusterApiResponseData() *OneOfClusterApiResponseData {
  p := new(OneOfClusterApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfClusterApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfClusterApiResponseData is nil"))
  }
  switch v.(type) {
    case ClusterMetrics:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterMetrics)}
      *p.oneOfType0 = v.(ClusterMetrics)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfClusterApiResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfClusterApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(ClusterMetrics)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "aiops.v4.clusterMetrics.ClusterMetrics" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterMetrics)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterApiResponseData"))
}

func (p *OneOfClusterApiResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfClusterApiResponseData")
}

type OneOfClusterListApiResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []ClusterMetrics `json:"-"`
  oneOfType400 *import1.ErrorResponse `json:"-"`
}

func NewOneOfClusterListApiResponseData() *OneOfClusterListApiResponseData {
  p := new(OneOfClusterListApiResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfClusterListApiResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfClusterListApiResponseData is nil"))
  }
  switch v.(type) {
    case []ClusterMetrics:
      p.oneOfType0 = v.([]ClusterMetrics)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<aiops.v4.clusterMetrics.ClusterMetrics>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<aiops.v4.clusterMetrics.ClusterMetrics>"
    case import1.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = v.(import1.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfClusterListApiResponseData) GetValue() interface{} {
  if "List<aiops.v4.clusterMetrics.ClusterMetrics>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfClusterListApiResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]ClusterMetrics)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "aiops.v4.clusterMetrics.ClusterMetrics" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<aiops.v4.clusterMetrics.ClusterMetrics>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<aiops.v4.clusterMetrics.ClusterMetrics>"
      return nil

    }
  }
  vOneOfType400 := new(import1.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "aiops.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import1.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterListApiResponseData"))
}

func (p *OneOfClusterListApiResponseData) MarshalJSON() ([]byte, error) {
  if "List<aiops.v4.clusterMetrics.ClusterMetrics>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfClusterListApiResponseData")
}

