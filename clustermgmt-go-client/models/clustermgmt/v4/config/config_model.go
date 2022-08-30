/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.0.1-alpha-1
 *
 * Part of the Nutanix Clustermgmt Versioned APIs
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Hosts, Clusters and other Infrastructure
*/
package config
import (
  "bytes"
  import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
  import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
  "encoding/json"
  "errors"
  "fmt"
)

/**
Status of Acropolis connection to hypervisor
*/
type AcropolisConnectionState int

const(
  ACROPOLISCONNECTIONSTATE_UNKNOWN AcropolisConnectionState = 0
  ACROPOLISCONNECTIONSTATE_REDACTED AcropolisConnectionState = 1
  ACROPOLISCONNECTIONSTATE_CONNECTED AcropolisConnectionState = 2
  ACROPOLISCONNECTIONSTATE_DISCONNECTED AcropolisConnectionState = 3
)

// returns the name of the enum given an ordinal number
func (e *AcropolisConnectionState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CONNECTED",
    "DISCONNECTED",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *AcropolisConnectionState) index(name string) AcropolisConnectionState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CONNECTED",
    "DISCONNECTED",
  }
  for idx := range names {
    if names[idx] == name {
      return AcropolisConnectionState(idx)
    }
  }
  return ACROPOLISCONNECTIONSTATE_UNKNOWN
}

func (e *AcropolisConnectionState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for AcropolisConnectionState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *AcropolisConnectionState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e AcropolisConnectionState) Ref() *AcropolisConnectionState {
  return &e
}


/**
Attribute item information
*/
type AttributeItem struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Tolerance message attribute key
  */
  Attribute *string `json:"attribute,omitempty"`
  /**
  Tolerance message attribute value
  */
  Value *string `json:"value,omitempty"`
}


func NewAttributeItem() *AttributeItem {
  p := new(AttributeItem)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.AttributeItem"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.AttributeItem"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Build information details
*/
type BuildReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Software build type
  */
  BuildType *string `json:"buildType,omitempty"`
  /**
  Commit Id used for version
  */
  CommitId *string `json:"commitId,omitempty"`
  /**
  Full name of software version
  */
  FullVersion *string `json:"fullVersion,omitempty"`
  /**
  Short commit Id used for version
  */
  ShortCommitId *string `json:"shortCommitId,omitempty"`
  /**
  Software version
  */
  Version *string `json:"version,omitempty"`
}


func NewBuildReference() *BuildReference {
  p := new(BuildReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.BuildReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.BuildReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster arch
*/
type ClusterArchReference int

const(
  CLUSTERARCHREFERENCE_UNKNOWN ClusterArchReference = 0
  CLUSTERARCHREFERENCE_REDACTED ClusterArchReference = 1
  CLUSTERARCHREFERENCE_X86_64 ClusterArchReference = 2
  CLUSTERARCHREFERENCE_PPC64LE ClusterArchReference = 3
)

// returns the name of the enum given an ordinal number
func (e *ClusterArchReference) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "X86_64",
    "PPC64LE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ClusterArchReference) index(name string) ClusterArchReference {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "X86_64",
    "PPC64LE",
  }
  for idx := range names {
    if names[idx] == name {
      return ClusterArchReference(idx)
    }
  }
  return CLUSTERARCHREFERENCE_UNKNOWN
}

func (e *ClusterArchReference) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterArchReference:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ClusterArchReference) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ClusterArchReference) Ref() *ClusterArchReference {
  return &e
}


/**
Cluster configuration details
*/
type ClusterConfigReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Public ssh key details
  */
  AuthorizedPublicKeyList []PublicKey `json:"authorizedPublicKeyList,omitempty"`
  
  BuildInfo *BuildReference `json:"buildInfo,omitempty"`
  
  ClusterArch *ClusterArchReference `json:"clusterArch,omitempty"`
  /**
  Cluster function
  */
  ClusterFunction []ClusterFunctionRef `json:"clusterFunction,omitempty"`
  /**
  Cluster software version details
  */
  ClusterSoftwareMap []SoftwareMapReference `json:"clusterSoftwareMap,omitempty"`
  
  FaultToleranceState *FaultToleranceState `json:"faultToleranceState,omitempty"`
  /**
  Hypervisor type
  */
  HypervisorTypes []HypervisorType `json:"hypervisorTypes,omitempty"`
  /**
  Cluster incarnation Id
  */
  IncarnationId *int64 `json:"incarnationId,omitempty"`
  /**
  Indicates whether the release is categorized as Long-term or not
  */
  IsLts *bool `json:"isLts,omitempty"`
  
  OperationMode *OperationMode `json:"operationMode,omitempty"`
  /**
  Password ssh into cluster is enabled or not
  */
  PasswordRemoteLoginEnabled *bool `json:"passwordRemoteLoginEnabled,omitempty"`
  /**
  Redundancy factor of a cluster
  */
  RedundancyFactor *int64 `json:"redundancyFactor,omitempty"`
  /**
  Remote support status
  */
  RemoteSupport *bool `json:"remoteSupport,omitempty"`
  /**
  Time zone on a cluster
  */
  Timezone *string `json:"timezone,omitempty"`
}


func NewClusterConfigReference() *ClusterConfigReference {
  p := new(ClusterConfigReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ClusterConfigReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ClusterConfigReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}





type ClusterEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  Config *ClusterConfigReference `json:"config,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Cluster name
  */
  Name *string `json:"name,omitempty"`
  
  Network *ClusterNetworkReference `json:"network,omitempty"`
  
  Nodes *NodeReference `json:"nodes,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewClusterEntity() *ClusterEntity {
  p := new(ClusterEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ClusterEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ClusterEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster function
*/
type ClusterFunctionRef int

const(
  CLUSTERFUNCTIONREF_UNKNOWN ClusterFunctionRef = 0
  CLUSTERFUNCTIONREF_REDACTED ClusterFunctionRef = 1
  CLUSTERFUNCTIONREF_AOS ClusterFunctionRef = 2
  CLUSTERFUNCTIONREF_PRISM_CENTRAL ClusterFunctionRef = 3
  CLUSTERFUNCTIONREF_CLOUD_DATA_GATEWAY ClusterFunctionRef = 4
  CLUSTERFUNCTIONREF_AFS ClusterFunctionRef = 5
  CLUSTERFUNCTIONREF_WITNESS ClusterFunctionRef = 6
  CLUSTERFUNCTIONREF_XI_PORTAL ClusterFunctionRef = 7
  CLUSTERFUNCTIONREF_ONE_NODE ClusterFunctionRef = 8
  CLUSTERFUNCTIONREF_TWO_NODE ClusterFunctionRef = 9
)

// returns the name of the enum given an ordinal number
func (e *ClusterFunctionRef) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AOS",
    "PRISM_CENTRAL",
    "CLOUD_DATA_GATEWAY",
    "AFS",
    "WITNESS",
    "XI_PORTAL",
    "ONE_NODE",
    "TWO_NODE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ClusterFunctionRef) index(name string) ClusterFunctionRef {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AOS",
    "PRISM_CENTRAL",
    "CLOUD_DATA_GATEWAY",
    "AFS",
    "WITNESS",
    "XI_PORTAL",
    "ONE_NODE",
    "TWO_NODE",
  }
  for idx := range names {
    if names[idx] == name {
      return ClusterFunctionRef(idx)
    }
  }
  return CLUSTERFUNCTIONREF_UNKNOWN
}

func (e *ClusterFunctionRef) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterFunctionRef:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ClusterFunctionRef) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ClusterFunctionRef) Ref() *ClusterFunctionRef {
  return &e
}


/**
Network details of a cluster
*/
type ClusterNetworkReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  ExternalAddress *IpAddresses `json:"externalAddress,omitempty"`
  /**
  Cluster external data service IP
  */
  ExternalDataServiceIp *string `json:"externalDataServiceIp,omitempty"`
  /**
  Cluster external subnet address
  */
  ExternalSubnet *string `json:"externalSubnet,omitempty"`
  /**
  Cluster internal subnet address
  */
  InternalSubnet *string `json:"internalSubnet,omitempty"`
  
  ManagementServer *ManagementServerRef `json:"managementServer,omitempty"`
  /**
  Cluster NAT or proxy IP that maps to a local IP
  */
  MasqueradingIp *string `json:"masqueradingIp,omitempty"`
  /**
  The port to connect to the cluster when using masquerading IP
  */
  MasqueradingPort *int `json:"masqueradingPort,omitempty"`
  /**
  Name servers on a cluster
  */
  NameServerIpList []string `json:"nameServerIpList,omitempty"`
  /**
  NFS subnet whitelist addresses
  */
  NfsSubnetWhitelist []string `json:"nfsSubnetWhitelist,omitempty"`
  /**
  NTP servers on a cluster
  */
  NtpServerIpList []string `json:"ntpServerIpList,omitempty"`
  
  SmtpServer *SmtpServerRef `json:"smtpServer,omitempty"`
}


func NewClusterNetworkReference() *ClusterNetworkReference {
  p := new(ClusterNetworkReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ClusterNetworkReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ClusterNetworkReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster reference of a node
*/
type ClusterReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster name
  */
  Name *string `json:"name,omitempty"`
  /**
  Cluster UUID
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewClusterReference() *ClusterReference {
  p := new(ClusterReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ClusterReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ClusterReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Fault tolerance information of a component
*/
type ComponentFaultTolerance struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  DetailMessage *ToleranceMessage `json:"detailMessage,omitempty"`
  /**
  Time of last update
  */
  LastUpdatesSecs *int64 `json:"lastUpdatesSecs,omitempty"`
  /**
  Maximum fault tolerance
  */
  MaxFaultsTolerated *int `json:"maxFaultsTolerated,omitempty"`
  
  Type *ComponentType `json:"type,omitempty"`
  /**
  Indicates whether the tolerance computation is in progress or not
  */
  UnderComputation *bool `json:"underComputation,omitempty"`
}


func NewComponentFaultTolerance() *ComponentFaultTolerance {
  p := new(ComponentFaultTolerance)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ComponentFaultTolerance"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ComponentFaultTolerance"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of component
*/
type ComponentType int

const(
  COMPONENTTYPE_UNKNOWN ComponentType = 0
  COMPONENTTYPE_REDACTED ComponentType = 1
  COMPONENTTYPE_EXTENT_GROUP_REPLICAS ComponentType = 2
  COMPONENTTYPE_OPLOG_EPISODES ComponentType = 3
  COMPONENTTYPE_CASSANDRA_RING ComponentType = 4
  COMPONENTTYPE_ZOOKEPER_INSTANCES ComponentType = 5
  COMPONENTTYPE_FREE_SPACE ComponentType = 6
  COMPONENTTYPE_STATIC_CONFIG ComponentType = 7
  COMPONENTTYPE_ERASURE_CODE_STRIP_SIZE ComponentType = 8
  COMPONENTTYPE_STARGATE_HEALTH ComponentType = 9
)

// returns the name of the enum given an ordinal number
func (e *ComponentType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "EXTENT_GROUP_REPLICAS",
    "OPLOG_EPISODES",
    "CASSANDRA_RING",
    "ZOOKEPER_INSTANCES",
    "FREE_SPACE",
    "STATIC_CONFIG",
    "ERASURE_CODE_STRIP_SIZE",
    "STARGATE_HEALTH",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ComponentType) index(name string) ComponentType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "EXTENT_GROUP_REPLICAS",
    "OPLOG_EPISODES",
    "CASSANDRA_RING",
    "ZOOKEPER_INSTANCES",
    "FREE_SPACE",
    "STATIC_CONFIG",
    "ERASURE_CODE_STRIP_SIZE",
    "STARGATE_HEALTH",
  }
  for idx := range names {
    if names[idx] == name {
      return ComponentType(idx)
    }
  }
  return COMPONENTTYPE_UNKNOWN
}

func (e *ComponentType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ComponentType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ComponentType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ComponentType) Ref() *ComponentType {
  return &e
}


/**
Host entity with its attributes
*/
type ControllerVmReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Backplane address
  */
  BackplaneAddress *string `json:"backplaneAddress,omitempty"`
  
  ExternalAddress *IpAddresses `json:"externalAddress,omitempty"`
  /**
  Controller VM Id
  */
  Id *int64 `json:"id,omitempty"`
  
  Ipmi *IpmiReference `json:"ipmi,omitempty"`
  /**
  Maintenance mode status
  */
  MaintenanceMode *bool `json:"maintenanceMode,omitempty"`
  /**
  NAT IP address
  */
  NatIp *string `json:"natIp,omitempty"`
  /**
  NAT port
  */
  NatPort *int `json:"natPort,omitempty"`
  /**
  Rackable unit UUID
  */
  RackableUnitUuid *string `json:"rackableUnitUuid,omitempty"`
  /**
  RDMA backplane address
  */
  RdmaBackplaneAddress []string `json:"rdmaBackplaneAddress,omitempty"`
}


func NewControllerVmReference() *ControllerVmReference {
  p := new(ControllerVmReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ControllerVmReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ControllerVmReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Disk details attached to a host
*/
type DiskReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Disk Id
  */
  Id *int64 `json:"id,omitempty"`
  /**
  Disk mount path
  */
  MountPath *string `json:"mountPath,omitempty"`
  /**
  Disk serial Id
  */
  SerialId *string `json:"serialId,omitempty"`
  /**
  Disk size
  */
  Size *int64 `json:"size,omitempty"`
  
  StorageTier *StorageTierReference `json:"storageTier,omitempty"`
  /**
  Disk UUID
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewDiskReference() *DiskReference {
  p := new(DiskReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.DiskReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.DiskReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Domain awareness level
*/
type DomainAwarenessLevel int

const(
  DOMAINAWARENESSLEVEL_UNKNOWN DomainAwarenessLevel = 0
  DOMAINAWARENESSLEVEL_REDACTED DomainAwarenessLevel = 1
  DOMAINAWARENESSLEVEL_NODE DomainAwarenessLevel = 2
  DOMAINAWARENESSLEVEL_BLOCK DomainAwarenessLevel = 3
  DOMAINAWARENESSLEVEL_RACK DomainAwarenessLevel = 4
)

// returns the name of the enum given an ordinal number
func (e *DomainAwarenessLevel) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NODE",
    "BLOCK",
    "RACK",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DomainAwarenessLevel) index(name string) DomainAwarenessLevel {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NODE",
    "BLOCK",
    "RACK",
  }
  for idx := range names {
    if names[idx] == name {
      return DomainAwarenessLevel(idx)
    }
  }
  return DOMAINAWARENESSLEVEL_UNKNOWN
}

func (e *DomainAwarenessLevel) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DomainAwarenessLevel:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DomainAwarenessLevel) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DomainAwarenessLevel) Ref() *DomainAwarenessLevel {
  return &e
}


/**
Domain Fault tolerance configuration
*/
type DomainFaultTolerance struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of components in a domain
  */
  ComponentStatus []ComponentFaultTolerance `json:"componentStatus,omitempty"`
  
  Type *DomainType `json:"type,omitempty"`
}


func NewDomainFaultTolerance() *DomainFaultTolerance {
  p := new(DomainFaultTolerance)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.DomainFaultTolerance"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.DomainFaultTolerance"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of domain
*/
type DomainType int

const(
  DOMAINTYPE_UNKNOWN DomainType = 0
  DOMAINTYPE_REDACTED DomainType = 1
  DOMAINTYPE_DEPRECATED_RACKABLE_UNIT DomainType = 2
  DOMAINTYPE_DEPRECATED_NODE DomainType = 3
  DOMAINTYPE_DEPRECATED_DISK DomainType = 4
  DOMAINTYPE_CUSTOM DomainType = 5
  DOMAINTYPE_DISK DomainType = 6
  DOMAINTYPE_NODE DomainType = 7
  DOMAINTYPE_RACKABLE_UNIT DomainType = 8
  DOMAINTYPE_RACK DomainType = 9
  DOMAINTYPE_CLUSTER DomainType = 10
)

// returns the name of the enum given an ordinal number
func (e *DomainType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DEPRECATED_RACKABLE_UNIT",
    "DEPRECATED_NODE",
    "DEPRECATED_DISK",
    "CUSTOM",
    "DISK",
    "NODE",
    "RACKABLE_UNIT",
    "RACK",
    "CLUSTER",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *DomainType) index(name string) DomainType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DEPRECATED_RACKABLE_UNIT",
    "DEPRECATED_NODE",
    "DEPRECATED_DISK",
    "CUSTOM",
    "DISK",
    "NODE",
    "RACKABLE_UNIT",
    "RACK",
    "CLUSTER",
  }
  for idx := range names {
    if names[idx] == name {
      return DomainType(idx)
    }
  }
  return DOMAINTYPE_UNKNOWN
}

func (e *DomainType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for DomainType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *DomainType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e DomainType) Ref() *DomainType {
  return &e
}


/**
Fault tolerant state of cluster
*/
type FaultToleranceState struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Maximum fault tolerance that is supported currently
  */
  CurrentMaxFaultTolerance *int `json:"currentMaxFaultTolerance,omitempty"`
  /**
  Desired Maximum fault tolerance
  */
  DesiredMaxFaultTolerance *int `json:"desiredMaxFaultTolerance,omitempty"`
  
  DomainAwarenessLevel *DomainAwarenessLevel `json:"domainAwarenessLevel,omitempty"`
}


func NewFaultToleranceState() *FaultToleranceState {
  p := new(FaultToleranceState)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.FaultToleranceState"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.FaultToleranceState"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/host-gpus Get operation
*/
type GetClusterHostGpusResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetClusterHostGpusResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetClusterHostGpusResponse() *GetClusterHostGpusResponse {
  p := new(GetClusterHostGpusResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetClusterHostGpusResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetClusterHostGpusResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetClusterHostGpusResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetClusterHostGpusResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetClusterHostGpusResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId} Get operation
*/
type GetClusterResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetClusterResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetClusterResponse() *GetClusterResponse {
  p := new(GetClusterResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetClusterResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetClusterResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetClusterResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetClusterResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetClusterResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/clusters Get operation
*/
type GetClustersResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetClustersResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetClustersResponse() *GetClustersResponse {
  p := new(GetClustersResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetClustersResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetClustersResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetClustersResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetClustersResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetClustersResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/domain-fault-tolerance-status Get operation
*/
type GetDomainFaultToleranceResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetDomainFaultToleranceResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetDomainFaultToleranceResponse() *GetDomainFaultToleranceResponse {
  p := new(GetDomainFaultToleranceResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetDomainFaultToleranceResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetDomainFaultToleranceResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetDomainFaultToleranceResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetDomainFaultToleranceResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetDomainFaultToleranceResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/gpu-profiles Get operation
*/
type GetGpuProfilesResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetGpuProfilesResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetGpuProfilesResponse() *GetGpuProfilesResponse {
  p := new(GetGpuProfilesResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetGpuProfilesResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetGpuProfilesResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetGpuProfilesResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetGpuProfilesResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetGpuProfilesResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/host/{hostExtId}/host-gpu/{hostGpuExtId} Get operation
*/
type GetHostGpuResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHostGpuResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHostGpuResponse() *GetHostGpuResponse {
  p := new(GetHostGpuResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetHostGpuResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetHostGpuResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHostGpuResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHostGpuResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHostGpuResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/host/{hostExtId}/host-gpus Get operation
*/
type GetHostGpusResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHostGpusResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHostGpusResponse() *GetHostGpusResponse {
  p := new(GetHostGpusResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetHostGpusResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetHostGpusResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHostGpusResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHostGpusResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHostGpusResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/host/{hostExtId} Get operation
*/
type GetHostResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHostResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHostResponse() *GetHostResponse {
  p := new(GetHostResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetHostResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetHostResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHostResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHostResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHostResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/hosts Get operation
*/
type GetHostsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetHostsResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetHostsResponse() *GetHostsResponse {
  p := new(GetHostsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetHostsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetHostsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetHostsResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetHostsResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetHostsResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/rackable-unit/{rackableUnitExtId} Get operation
*/
type GetRackableUnitResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetRackableUnitResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetRackableUnitResponse() *GetRackableUnitResponse {
  p := new(GetRackableUnitResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetRackableUnitResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetRackableUnitResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetRackableUnitResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetRackableUnitResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetRackableUnitResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/rackable-units Get operation
*/
type GetRackableUnitsResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetRackableUnitsResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetRackableUnitsResponse() *GetRackableUnitsResponse {
  p := new(GetRackableUnitsResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetRackableUnitsResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetRackableUnitsResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetRackableUnitsResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetRackableUnitsResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetRackableUnitsResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/rsyslog Get operation
*/
type GetRsyslogServersResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetRsyslogServersResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetRsyslogServersResponse() *GetRsyslogServersResponse {
  p := new(GetRsyslogServersResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetRsyslogServersResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetRsyslogServersResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetRsyslogServersResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetRsyslogServersResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetRsyslogServersResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a1/config/cluster/{clusterExtId}/snmp Get operation
*/
type GetSnmpResponse struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  
  */
  DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`
  
  Data *OneOfGetSnmpResponseData `json:"data,omitempty"`
  
  Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}


func NewGetSnmpResponse() *GetSnmpResponse {
  p := new(GetSnmpResponse)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GetSnmpResponse"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GetSnmpResponse"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}

func (p *GetSnmpResponse) GetData() interface{} {
  if nil == p.Data {
    return nil
  }
  return p.Data.GetValue()
}

func (p *GetSnmpResponse) SetData(v interface{}) error {
  if nil == p.Data {
    p.Data = NewOneOfGetSnmpResponseData()
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
GPU configuration details
*/
type GpuConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  GPU assignable
  */
  Assignable *int64 `json:"assignable,omitempty"`
  /**
  Device Id
  */
  DeviceId *string `json:"deviceId,omitempty"`
  /**
  Device name
  */
  DeviceName *string `json:"deviceName,omitempty"`
  /**
  GPU fraction
  */
  Fraction *int64 `json:"fraction,omitempty"`
  /**
  Frame buffer size in bytes
  */
  FrameBufferSizeBytes *int64 `json:"frameBufferSizeBytes,omitempty"`
  /**
  Guest driver version
  */
  GuestDriverVersion *string `json:"guestDriverVersion,omitempty"`
  /**
  GPU in use
  */
  InUse *bool `json:"inUse,omitempty"`
  /**
  GPU license list
  */
  LicenseList []string `json:"licenseList,omitempty"`
  /**
  Maximum resolution
  */
  MaxResolution *string `json:"maxResolution,omitempty"`
  
  Mode *GpuMode `json:"mode,omitempty"`
  /**
  NUMA node
  */
  NumaNode *string `json:"numaNode,omitempty"`
  /**
  Number of virtual display heads
  */
  NumberOfVirtualDisplayHeads *int64 `json:"numberOfVirtualDisplayHeads,omitempty"`
  /**
  SBDF address
  */
  Sbdf *string `json:"sbdf,omitempty"`
  
  Type *GpuType `json:"type,omitempty"`
  /**
  Vendor name
  */
  VendorName *string `json:"vendorName,omitempty"`
}


func NewGpuConfig() *GpuConfig {
  p := new(GpuConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GpuConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GpuConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
GPU mode
*/
type GpuMode int

const(
  GPUMODE_UNKNOWN GpuMode = 0
  GPUMODE_REDACTED GpuMode = 1
  GPUMODE_UNUSED GpuMode = 2
  GPUMODE_USED_FOR_PASSTHROUGH GpuMode = 3
  GPUMODE_USED_FOR_VIRTUAL GpuMode = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuMode) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UNUSED",
    "USED_FOR_PASSTHROUGH",
    "USED_FOR_VIRTUAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *GpuMode) index(name string) GpuMode {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UNUSED",
    "USED_FOR_PASSTHROUGH",
    "USED_FOR_VIRTUAL",
  }
  for idx := range names {
    if names[idx] == name {
      return GpuMode(idx)
    }
  }
  return GPUMODE_UNKNOWN
}

func (e *GpuMode) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for GpuMode:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *GpuMode) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e GpuMode) Ref() *GpuMode {
  return &e
}


/**
GPU Profile
*/
type GpuProfile struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of UUIDs of virtual machines with an allocated GPU belonging to this profile
  */
  AllocatedVmUuids []string `json:"allocatedVmUuids,omitempty"`
  /**
  Device Id
  */
  DeviceId *string `json:"deviceId,omitempty"`
  
  GpuConfig *GpuConfig `json:"gpuConfig,omitempty"`
}


func NewGpuProfile() *GpuProfile {
  p := new(GpuProfile)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.GpuProfile"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.GpuProfile"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
GPU type
*/
type GpuType int

const(
  GPUTYPE_UNKNOWN GpuType = 0
  GPUTYPE_REDACTED GpuType = 1
  GPUTYPE_PASS_THROUGH_GRAPHICS GpuType = 2
  GPUTYPE_PASS_THROUGH_COMPUTE GpuType = 3
  GPUTYPE_VIRTUAL GpuType = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PASS_THROUGH_GRAPHICS",
    "PASS_THROUGH_COMPUTE",
    "VIRTUAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *GpuType) index(name string) GpuType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PASS_THROUGH_GRAPHICS",
    "PASS_THROUGH_COMPUTE",
    "VIRTUAL",
  }
  for idx := range names {
    if names[idx] == name {
      return GpuType(idx)
    }
  }
  return GPUTYPE_UNKNOWN
}

func (e *GpuType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for GpuType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *GpuType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e GpuType) Ref() *GpuType {
  return &e
}


/**
Host entity with its attributes
*/
type HostEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Boot time in secs
  */
  BootTimeUsecs *int64 `json:"bootTimeUsecs,omitempty"`
  
  Cluster *ClusterReference `json:"cluster,omitempty"`
  
  ControllerVm *ControllerVmReference `json:"controllerVm,omitempty"`
  /**
  CPU capacity in Hz
  */
  CpuCapacityHz *int64 `json:"cpuCapacityHz,omitempty"`
  /**
  CPU frequency in Hz
  */
  CpuFrequencyHz *int64 `json:"cpuFrequencyHz,omitempty"`
  /**
  CPU model name
  */
  CpuModel *string `json:"cpuModel,omitempty"`
  /**
  Default VHD container Id
  */
  DefaultVhdContainerId *string `json:"defaultVhdContainerId,omitempty"`
  /**
  Default VHD container UUID
  */
  DefaultVhdContainerUuid *string `json:"defaultVhdContainerUuid,omitempty"`
  /**
  Default VHD location
  */
  DefaultVhdLocation *string `json:"defaultVhdLocation,omitempty"`
  /**
  Default VM container Id
  */
  DefaultVmContainerId *string `json:"defaultVmContainerId,omitempty"`
  /**
  Default VM container UUID
  */
  DefaultVmContainerUuid *string `json:"defaultVmContainerUuid,omitempty"`
  /**
  Default VM location
  */
  DefaultVmLocation *string `json:"defaultVmLocation,omitempty"`
  /**
  Disks attached to host
  */
  Disk []DiskReference `json:"disk,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Failover cluster FQDN
  */
  FailoverClusterFqdn *string `json:"failoverClusterFqdn,omitempty"`
  /**
  Failover cluster node status
  */
  FailoverClusterNodeStatus *string `json:"failoverClusterNodeStatus,omitempty"`
  /**
  GPU driver version
  */
  GpuDriverVersion *string `json:"gpuDriverVersion,omitempty"`
  /**
  GPU attached list
  */
  GpuList []string `json:"gpuList,omitempty"`
  /**
  Certificate signing request status
  */
  HasCsr *bool `json:"hasCsr,omitempty"`
  /**
  Name of the host
  */
  HostName *string `json:"hostName,omitempty"`
  /**
  Host NICs Id
  */
  HostNicsIdList []string `json:"hostNicsIdList,omitempty"`
  
  HostType *HostTypeEnum `json:"hostType,omitempty"`
  
  Hypervisor *HypervisorReference `json:"hypervisor,omitempty"`
  /**
  Node degraded status
  */
  IsDegraded *bool `json:"isDegraded,omitempty"`
  /**
  Indicates whether the hardware is virtualized or not
  */
  IsHardwareVirtualized *bool `json:"isHardwareVirtualized,omitempty"`
  /**
  Secure boot status
  */
  IsSecureBooted *bool `json:"isSecureBooted,omitempty"`
  /**
  Mapping of key management device to certificate status list
  */
  KeyManagementDeviceToCertStatus []KeyManagementDeviceToCertStatus `json:"keyManagementDeviceToCertStatus,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Memory size in bytes
  */
  MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
  /**
  Number of CPU cores
  */
  NumberOfCpuCores *int64 `json:"numberOfCpuCores,omitempty"`
  /**
  Number of CPU sockets
  */
  NumberOfCpuSockets *int64 `json:"numberOfCpuSockets,omitempty"`
  /**
  Number of CPU threads
  */
  NumberOfCpuThreads *int64 `json:"numberOfCpuThreads,omitempty"`
  /**
  Reboot pending status
  */
  RebootPending *bool `json:"rebootPending,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewHostEntity() *HostEntity {
  p := new(HostEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.HostEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.HostEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Host GPU details
*/
type HostGpuEntity struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster name
  */
  ClusterName *string `json:"clusterName,omitempty"`
  /**
  Cluster UUID
  */
  ClusterUuid *string `json:"clusterUuid,omitempty"`
  
  Config *GpuConfig `json:"config,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  /**
  Controller VM Id
  */
  NodeId *string `json:"nodeId,omitempty"`
  /**
  UUID of a node
  */
  NodeUuid *string `json:"nodeUuid,omitempty"`
  /**
  Number of vGPUs allocated
  */
  NumberOfVgpusAllocated *int64 `json:"numberOfVgpusAllocated,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewHostGpuEntity() *HostGpuEntity {
  p := new(HostGpuEntity)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.HostGpuEntity"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.HostGpuEntity"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of the host
*/
type HostTypeEnum int

const(
  HOSTTYPEENUM_UNKNOWN HostTypeEnum = 0
  HOSTTYPEENUM_REDACTED HostTypeEnum = 1
  HOSTTYPEENUM_HYPER_CONVERGED HostTypeEnum = 2
  HOSTTYPEENUM_COMPUTE_ONLY HostTypeEnum = 3
)

// returns the name of the enum given an ordinal number
func (e *HostTypeEnum) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HYPER_CONVERGED",
    "COMPUTE_ONLY",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *HostTypeEnum) index(name string) HostTypeEnum {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "HYPER_CONVERGED",
    "COMPUTE_ONLY",
  }
  for idx := range names {
    if names[idx] == name {
      return HostTypeEnum(idx)
    }
  }
  return HOSTTYPEENUM_UNKNOWN
}

func (e *HostTypeEnum) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for HostTypeEnum:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *HostTypeEnum) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e HostTypeEnum) Ref() *HostTypeEnum {
  return &e
}


/**
Hypervisor reference of a node
*/
type HypervisorReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  AcropolisConnectionState *AcropolisConnectionState `json:"acropolisConnectionState,omitempty"`
  
  ExternalAddress *IpAddresses `json:"externalAddress,omitempty"`
  /**
  Hypervisor full name
  */
  FullName *string `json:"fullName,omitempty"`
  /**
  Number of VMs
  */
  NumberOfVms *int64 `json:"numberOfVms,omitempty"`
  
  State *HypervisorState `json:"state,omitempty"`
  
  Type *HypervisorType `json:"type,omitempty"`
  /**
  Hypervisor user name
  */
  UserName *string `json:"userName,omitempty"`
}


func NewHypervisorReference() *HypervisorReference {
  p := new(HypervisorReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.HypervisorReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.HypervisorReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Hypervisor state
*/
type HypervisorState int

const(
  HYPERVISORSTATE_UNKNOWN HypervisorState = 0
  HYPERVISORSTATE_REDACTED HypervisorState = 1
  HYPERVISORSTATE_ACROPOLIS_NORMAL HypervisorState = 2
  HYPERVISORSTATE_ENTERING_MAINTENANCE_MODE HypervisorState = 3
  HYPERVISORSTATE_ENTERED_MAINTENANCE_MODE HypervisorState = 4
  HYPERVISORSTATE_RESERVED_FOR_HA_FAILOVER HypervisorState = 5
  HYPERVISORSTATE_ENTERING_MAINTENANCE_MODE_FROM_HA_FAILOVER HypervisorState = 6
  HYPERVISORSTATE_RESERVING_FOR_HA_FAILOVER HypervisorState = 7
  HYPERVISORSTATE_HA_FAILOVER_SOURCE HypervisorState = 8
  HYPERVISORSTATE_HA_FAILOVER_TARGET HypervisorState = 9
  HYPERVISORSTATE_HA_HEALING_SOURCE HypervisorState = 10
  HYPERVISORSTATE_HA_HEALING_TARGET HypervisorState = 11
)

// returns the name of the enum given an ordinal number
func (e *HypervisorState) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACROPOLIS-NORMAL",
    "ENTERING-MAINTENANCE-MODE",
    "ENTERED-MAINTENANCE-MODE",
    "RESERVED-FOR-HA-FAILOVER",
    "ENTERING-MAINTENANCE-MODE-FROM-HA-FAILOVER",
    "RESERVING-FOR-HA-FAILOVER",
    "HA-FAILOVER-SOURCE",
    "HA-FAILOVER-TARGET",
    "HA-HEALING-SOURCE",
    "HA-HEALING-TARGET",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *HypervisorState) index(name string) HypervisorState {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "ACROPOLIS-NORMAL",
    "ENTERING-MAINTENANCE-MODE",
    "ENTERED-MAINTENANCE-MODE",
    "RESERVED-FOR-HA-FAILOVER",
    "ENTERING-MAINTENANCE-MODE-FROM-HA-FAILOVER",
    "RESERVING-FOR-HA-FAILOVER",
    "HA-FAILOVER-SOURCE",
    "HA-FAILOVER-TARGET",
    "HA-HEALING-SOURCE",
    "HA-HEALING-TARGET",
  }
  for idx := range names {
    if names[idx] == name {
      return HypervisorState(idx)
    }
  }
  return HYPERVISORSTATE_UNKNOWN
}

func (e *HypervisorState) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for HypervisorState:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *HypervisorState) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e HypervisorState) Ref() *HypervisorState {
  return &e
}


/**
Hypervisor type
*/
type HypervisorType int

const(
  HYPERVISORTYPE_UNKNOWN HypervisorType = 0
  HYPERVISORTYPE_REDACTED HypervisorType = 1
  HYPERVISORTYPE_AHV HypervisorType = 2
  HYPERVISORTYPE_ESX HypervisorType = 3
  HYPERVISORTYPE_HYPERV HypervisorType = 4
  HYPERVISORTYPE_XEN HypervisorType = 5
  HYPERVISORTYPE_NULLTYPE HypervisorType = 6
)

// returns the name of the enum given an ordinal number
func (e *HypervisorType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AHV",
    "ESX",
    "HYPERV",
    "XEN",
    "NULLTYPE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *HypervisorType) index(name string) HypervisorType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "AHV",
    "ESX",
    "HYPERV",
    "XEN",
    "NULLTYPE",
  }
  for idx := range names {
    if names[idx] == name {
      return HypervisorType(idx)
    }
  }
  return HYPERVISORTYPE_UNKNOWN
}

func (e *HypervisorType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for HypervisorType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *HypervisorType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e HypervisorType) Ref() *HypervisorType {
  return &e
}


/**
IPV4 and IPV6 details
*/
type IpAddresses struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IPV4 address
  */
  Ipv4 *string `json:"ipv4,omitempty"`
  /**
  IPV6 address
  */
  Ipv6 *string `json:"ipv6,omitempty"`
}


func NewIpAddresses() *IpAddresses {
  p := new(IpAddresses)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.IpAddresses"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.IpAddresses"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
IPMI reference
*/
type IpmiReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  IPMI IP address
  */
  Ip *string `json:"ip,omitempty"`
  /**
  IPMI username
  */
  Username *string `json:"username,omitempty"`
}


func NewIpmiReference() *IpmiReference {
  p := new(IpmiReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.IpmiReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.IpmiReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Mapping of key management device to certificate status list
*/
type KeyManagementDeviceToCertStatus struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Key management server name
  */
  KeyManagementServerName *string `json:"keyManagementServerName,omitempty"`
  /**
  Certificate status
  */
  Status *bool `json:"status,omitempty"`
}


func NewKeyManagementDeviceToCertStatus() *KeyManagementDeviceToCertStatus {
  p := new(KeyManagementDeviceToCertStatus)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.KeyManagementDeviceToCertStatus"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.KeyManagementDeviceToCertStatus"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Management server information
*/
type ManagementServerRef struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Indicates whether it is DRS enabled or not
  */
  DrsEnabled *bool `json:"drsEnabled,omitempty"`
  /**
  Indicates whether the host is managed by an entity or not
  */
  InUse *bool `json:"inUse,omitempty"`
  /**
  Management server IP address
  */
  Ip *string `json:"ip,omitempty"`
  /**
  Indicates whether it is registered or not
  */
  IsRegistered *bool `json:"isRegistered,omitempty"`
  
  Type *ManagementServerType `json:"type,omitempty"`
}


func NewManagementServerRef() *ManagementServerRef {
  p := new(ManagementServerRef)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ManagementServerRef"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ManagementServerRef"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Management server type
*/
type ManagementServerType int

const(
  MANAGEMENTSERVERTYPE_UNKNOWN ManagementServerType = 0
  MANAGEMENTSERVERTYPE_REDACTED ManagementServerType = 1
  MANAGEMENTSERVERTYPE_VCENTER ManagementServerType = 2
)

// returns the name of the enum given an ordinal number
func (e *ManagementServerType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VCENTER",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *ManagementServerType) index(name string) ManagementServerType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "VCENTER",
  }
  for idx := range names {
    if names[idx] == name {
      return ManagementServerType(idx)
    }
  }
  return MANAGEMENTSERVERTYPE_UNKNOWN
}

func (e *ManagementServerType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for ManagementServerType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *ManagementServerType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e ManagementServerType) Ref() *ManagementServerType {
  return &e
}


/**
List of nodes in a cluster
*/
type NodeListItemReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Controller VM IP
  */
  ControllerVmIp *string `json:"controllerVmIp,omitempty"`
  /**
  Host IP address
  */
  HostIp *string `json:"hostIp,omitempty"`
  /**
  UUID of a node
  */
  NodeUuid *string `json:"nodeUuid,omitempty"`
}


func NewNodeListItemReference() *NodeListItemReference {
  p := new(NodeListItemReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.NodeListItemReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.NodeListItemReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Node reference for a cluster
*/
type NodeReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of nodes in a cluster
  */
  NodeList []NodeListItemReference `json:"nodeList,omitempty"`
  /**
  Number of nodes in a cluster
  */
  NumberOfNodes *int `json:"numberOfNodes,omitempty"`
}


func NewNodeReference() *NodeReference {
  p := new(NodeReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.NodeReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.NodeReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster operation mode
*/
type OperationMode int

const(
  OPERATIONMODE_UNKNOWN OperationMode = 0
  OPERATIONMODE_REDACTED OperationMode = 1
  OPERATIONMODE_NORMAL OperationMode = 2
  OPERATIONMODE_READ_ONLY OperationMode = 3
  OPERATIONMODE_STAND_ALONE OperationMode = 4
  OPERATIONMODE_SWITCH_TO_TWO_NODE OperationMode = 5
  OPERATIONMODE_OVERRIDE OperationMode = 6
)

// returns the name of the enum given an ordinal number
func (e *OperationMode) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NORMAL",
    "READ-ONLY",
    "STAND-ALONE",
    "SWITCH-TO-TWO-NODE",
    "OVERRIDE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *OperationMode) index(name string) OperationMode {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NORMAL",
    "READ-ONLY",
    "STAND-ALONE",
    "SWITCH-TO-TWO-NODE",
    "OVERRIDE",
  }
  for idx := range names {
    if names[idx] == name {
      return OperationMode(idx)
    }
  }
  return OPERATIONMODE_UNKNOWN
}

func (e *OperationMode) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for OperationMode:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *OperationMode) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e OperationMode) Ref() *OperationMode {
  return &e
}


/**
Public ssh key details
*/
type PublicKey struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Ssh key value
  */
  Key *string `json:"key,omitempty"`
  /**
  Ssh key name
  */
  Name *string `json:"name,omitempty"`
}


func NewPublicKey() *PublicKey {
  p := new(PublicKey)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.PublicKey"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.PublicKey"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Rack reference for the block
*/
type RackReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Cluster entity with its attributes
  */
  Id *int64 `json:"id,omitempty"`
  /**
  Cluster entity with its attributes
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewRackReference() *RackReference {
  p := new(RackReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.RackReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.RackReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Rackable Unit configuration
*/
type RackableUnit struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  A globally unique identifier of an instance that is suitable for external consumption.
  */
  ExtId *string `json:"extId,omitempty"`
  /**
  Rackable unit Id
  */
  Id *int64 `json:"id,omitempty"`
  /**
  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
  */
  Links []import1.ApiLink `json:"links,omitempty"`
  
  Model *RackableUnitModel `json:"model,omitempty"`
  /**
  Rackable unit model name
  */
  ModelName *string `json:"modelName,omitempty"`
  /**
  List of node information registered to the block
  */
  Nodes []RackableUnitNode `json:"nodes,omitempty"`
  
  Rack *RackReference `json:"rack,omitempty"`
  /**
  Rackable unit serial name
  */
  Serial *string `json:"serial,omitempty"`
  /**
  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
  */
  TenantId *string `json:"tenantId,omitempty"`
}


func NewRackableUnit() *RackableUnit {
  p := new(RackableUnit)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.RackableUnit"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.RackableUnit"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Rackable unit model type
*/
type RackableUnitModel int

const(
  RACKABLEUNITMODEL_UNKNOWN RackableUnitModel = 0
  RACKABLEUNITMODEL_REDACTED RackableUnitModel = 1
  RACKABLEUNITMODEL_DESKTOP RackableUnitModel = 2
  RACKABLEUNITMODEL_NX2000 RackableUnitModel = 3
  RACKABLEUNITMODEL_NX3000 RackableUnitModel = 4
  RACKABLEUNITMODEL_NX3050 RackableUnitModel = 5
  RACKABLEUNITMODEL_NX6050 RackableUnitModel = 6
  RACKABLEUNITMODEL_NX6070 RackableUnitModel = 7
  RACKABLEUNITMODEL_NX1050 RackableUnitModel = 8
  RACKABLEUNITMODEL_NX3060 RackableUnitModel = 9
  RACKABLEUNITMODEL_NX6060 RackableUnitModel = 10
  RACKABLEUNITMODEL_NX6080 RackableUnitModel = 11
  RACKABLEUNITMODEL_NX6020 RackableUnitModel = 12
  RACKABLEUNITMODEL_NX7110 RackableUnitModel = 13
  RACKABLEUNITMODEL_NX1020 RackableUnitModel = 14
  RACKABLEUNITMODEL_NX9040 RackableUnitModel = 15
  RACKABLEUNITMODEL_USELAYOUT RackableUnitModel = 16
  RACKABLEUNITMODEL_NULLVALUE RackableUnitModel = 17
)

// returns the name of the enum given an ordinal number
func (e *RackableUnitModel) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DESKTOP",
    "NX2000",
    "NX3000",
    "NX3050",
    "NX6050",
    "NX6070",
    "NX1050",
    "NX3060",
    "NX6060",
    "NX6080",
    "NX6020",
    "NX7110",
    "NX1020",
    "NX9040",
    "USELAYOUT",
    "NULLVALUE",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RackableUnitModel) index(name string) RackableUnitModel {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DESKTOP",
    "NX2000",
    "NX3000",
    "NX3050",
    "NX6050",
    "NX6070",
    "NX1050",
    "NX3060",
    "NX6060",
    "NX6080",
    "NX6020",
    "NX7110",
    "NX1020",
    "NX9040",
    "USELAYOUT",
    "NULLVALUE",
  }
  for idx := range names {
    if names[idx] == name {
      return RackableUnitModel(idx)
    }
  }
  return RACKABLEUNITMODEL_UNKNOWN
}

func (e *RackableUnitModel) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RackableUnitModel:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RackableUnitModel) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RackableUnitModel) Ref() *RackableUnitModel {
  return &e
}


/**
Node information registered to this rackable unit
*/
type RackableUnitNode struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Position of a node in a rackable unit
  */
  Position *int `json:"position,omitempty"`
  /**
  Controller VM Id
  */
  SvmId *int64 `json:"svmId,omitempty"`
  /**
  UUID of a node
  */
  Uuid *string `json:"uuid,omitempty"`
}


func NewRackableUnitNode() *RackableUnitNode {
  p := new(RackableUnitNode)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.RackableUnitNode"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.RackableUnitNode"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
RSYSLOG Module information
*/
type RsyslogModuleItem struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  LogSeverityLevel *RsyslogModuleLogSeverityLevel `json:"logSeverityLevel,omitempty"`
  /**
  Option to log monitor files for a module
  */
  Monitor *bool `json:"monitor,omitempty"`
  
  Name *RsyslogModuleName `json:"name,omitempty"`
}


func NewRsyslogModuleItem() *RsyslogModuleItem {
  p := new(RsyslogModuleItem)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.RsyslogModuleItem"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.RsyslogModuleItem"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
RSYSLOG module log severity level
*/
type RsyslogModuleLogSeverityLevel int

const(
  RSYSLOGMODULELOGSEVERITYLEVEL_UNKNOWN RsyslogModuleLogSeverityLevel = 0
  RSYSLOGMODULELOGSEVERITYLEVEL_REDACTED RsyslogModuleLogSeverityLevel = 1
  RSYSLOGMODULELOGSEVERITYLEVEL_EMERGENCY RsyslogModuleLogSeverityLevel = 2
  RSYSLOGMODULELOGSEVERITYLEVEL_ALERT RsyslogModuleLogSeverityLevel = 3
  RSYSLOGMODULELOGSEVERITYLEVEL_CRITICAL RsyslogModuleLogSeverityLevel = 4
  RSYSLOGMODULELOGSEVERITYLEVEL_ERROR RsyslogModuleLogSeverityLevel = 5
  RSYSLOGMODULELOGSEVERITYLEVEL_WARNING RsyslogModuleLogSeverityLevel = 6
  RSYSLOGMODULELOGSEVERITYLEVEL_NOTICE RsyslogModuleLogSeverityLevel = 7
  RSYSLOGMODULELOGSEVERITYLEVEL_INFO RsyslogModuleLogSeverityLevel = 8
  RSYSLOGMODULELOGSEVERITYLEVEL_DEBUG RsyslogModuleLogSeverityLevel = 9
)

// returns the name of the enum given an ordinal number
func (e *RsyslogModuleLogSeverityLevel) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "EMERGENCY",
    "ALERT",
    "CRITICAL",
    "ERROR",
    "WARNING",
    "NOTICE",
    "INFO",
    "DEBUG",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RsyslogModuleLogSeverityLevel) index(name string) RsyslogModuleLogSeverityLevel {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "EMERGENCY",
    "ALERT",
    "CRITICAL",
    "ERROR",
    "WARNING",
    "NOTICE",
    "INFO",
    "DEBUG",
  }
  for idx := range names {
    if names[idx] == name {
      return RsyslogModuleLogSeverityLevel(idx)
    }
  }
  return RSYSLOGMODULELOGSEVERITYLEVEL_UNKNOWN
}

func (e *RsyslogModuleLogSeverityLevel) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RsyslogModuleLogSeverityLevel:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RsyslogModuleLogSeverityLevel) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RsyslogModuleLogSeverityLevel) Ref() *RsyslogModuleLogSeverityLevel {
  return &e
}


/**
RSYSLOG module name
*/
type RsyslogModuleName int

const(
  RSYSLOGMODULENAME_UNKNOWN RsyslogModuleName = 0
  RSYSLOGMODULENAME_REDACTED RsyslogModuleName = 1
  RSYSLOGMODULENAME_CASSANDRA RsyslogModuleName = 2
  RSYSLOGMODULENAME_CEREBRO RsyslogModuleName = 3
  RSYSLOGMODULENAME_CURATOR RsyslogModuleName = 4
  RSYSLOGMODULENAME_GENESIS RsyslogModuleName = 5
  RSYSLOGMODULENAME_PRISM RsyslogModuleName = 6
  RSYSLOGMODULENAME_STARGATE RsyslogModuleName = 7
  RSYSLOGMODULENAME_SYSLOG_MODULE RsyslogModuleName = 8
  RSYSLOGMODULENAME_ZOOKEEPER RsyslogModuleName = 9
  RSYSLOGMODULENAME_UHARA RsyslogModuleName = 10
  RSYSLOGMODULENAME_LAZAN RsyslogModuleName = 11
  RSYSLOGMODULENAME_API_AUDIT RsyslogModuleName = 12
  RSYSLOGMODULENAME_AUDIT RsyslogModuleName = 13
  RSYSLOGMODULENAME_CALM RsyslogModuleName = 14
  RSYSLOGMODULENAME_EPSILON RsyslogModuleName = 15
  RSYSLOGMODULENAME_ACROPOLIS RsyslogModuleName = 16
  RSYSLOGMODULENAME_MINERVA_CVM RsyslogModuleName = 17
  RSYSLOGMODULENAME_FLOW RsyslogModuleName = 18
  RSYSLOGMODULENAME_FLOW_SERVICE_LOGS RsyslogModuleName = 19
  RSYSLOGMODULENAME_LCM RsyslogModuleName = 20
  RSYSLOGMODULENAME_APLOS RsyslogModuleName = 21
)

// returns the name of the enum given an ordinal number
func (e *RsyslogModuleName) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CASSANDRA",
    "CEREBRO",
    "CURATOR",
    "GENESIS",
    "PRISM",
    "STARGATE",
    "SYSLOG_MODULE",
    "ZOOKEEPER",
    "UHARA",
    "LAZAN",
    "API_AUDIT",
    "AUDIT",
    "CALM",
    "EPSILON",
    "ACROPOLIS",
    "MINERVA_CVM",
    "FLOW",
    "FLOW_SERVICE_LOGS",
    "LCM",
    "APLOS",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RsyslogModuleName) index(name string) RsyslogModuleName {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "CASSANDRA",
    "CEREBRO",
    "CURATOR",
    "GENESIS",
    "PRISM",
    "STARGATE",
    "SYSLOG_MODULE",
    "ZOOKEEPER",
    "UHARA",
    "LAZAN",
    "API_AUDIT",
    "AUDIT",
    "CALM",
    "EPSILON",
    "ACROPOLIS",
    "MINERVA_CVM",
    "FLOW",
    "FLOW_SERVICE_LOGS",
    "LCM",
    "APLOS",
  }
  for idx := range names {
    if names[idx] == name {
      return RsyslogModuleName(idx)
    }
  }
  return RSYSLOGMODULENAME_UNKNOWN
}

func (e *RsyslogModuleName) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RsyslogModuleName:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RsyslogModuleName) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RsyslogModuleName) Ref() *RsyslogModuleName {
  return &e
}


/**
RSYSLOG server protocol type
*/
type RsyslogNetworkProtocol int

const(
  RSYSLOGNETWORKPROTOCOL_UNKNOWN RsyslogNetworkProtocol = 0
  RSYSLOGNETWORKPROTOCOL_REDACTED RsyslogNetworkProtocol = 1
  RSYSLOGNETWORKPROTOCOL_UDP RsyslogNetworkProtocol = 2
  RSYSLOGNETWORKPROTOCOL_TCP RsyslogNetworkProtocol = 3
  RSYSLOGNETWORKPROTOCOL_RELP RsyslogNetworkProtocol = 4
)

// returns the name of the enum given an ordinal number
func (e *RsyslogNetworkProtocol) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UDP",
    "TCP",
    "RELP",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *RsyslogNetworkProtocol) index(name string) RsyslogNetworkProtocol {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UDP",
    "TCP",
    "RELP",
  }
  for idx := range names {
    if names[idx] == name {
      return RsyslogNetworkProtocol(idx)
    }
  }
  return RSYSLOGNETWORKPROTOCOL_UNKNOWN
}

func (e *RsyslogNetworkProtocol) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for RsyslogNetworkProtocol:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *RsyslogNetworkProtocol) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e RsyslogNetworkProtocol) Ref() *RsyslogNetworkProtocol {
  return &e
}


/**
RSYSLOG Server
*/
type RsyslogServer struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  RSYSLOG server IP address
  */
  IpAddress *string `json:"ipAddress,omitempty"`
  /**
  List of modules registered to RSYSLOG server
  */
  Modules []RsyslogModuleItem `json:"modules,omitempty"`
  
  NetworkProtocol *RsyslogNetworkProtocol `json:"networkProtocol,omitempty"`
  /**
  RSYSLOG server port
  */
  Port *int `json:"port,omitempty"`
  /**
  RSYSLOG server name
  */
  ServerName *string `json:"serverName,omitempty"`
}


func NewRsyslogServer() *RsyslogServer {
  p := new(RsyslogServer)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.RsyslogServer"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.RsyslogServer"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SMTP network details
*/
type SmtpNetwork struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  IpAddress *IpAddresses `json:"ipAddress,omitempty"`
  /**
  SMTP server password
  */
  Password *string `json:"password,omitempty"`
  /**
  SMTP port
  */
  Port *int `json:"port,omitempty"`
  /**
  SMTP server user name
  */
  Username *string `json:"username,omitempty"`
}


func NewSmtpNetwork() *SmtpNetwork {
  p := new(SmtpNetwork)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SmtpNetwork"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SmtpNetwork"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SMTP servers on a cluster
*/
type SmtpServerRef struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SMTP Email address
  */
  EmailAddress *string `json:"emailAddress,omitempty"`
  
  Server *SmtpNetwork `json:"server,omitempty"`
  
  Type *SmtpType `json:"type,omitempty"`
}


func NewSmtpServerRef() *SmtpServerRef {
  p := new(SmtpServerRef)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SmtpServerRef"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SmtpServerRef"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Type of SMTP server
*/
type SmtpType int

const(
  SMTPTYPE_UNKNOWN SmtpType = 0
  SMTPTYPE_REDACTED SmtpType = 1
  SMTPTYPE_PLAIN SmtpType = 2
  SMTPTYPE_STARTTLS SmtpType = 3
  SMTPTYPE_SSL SmtpType = 4
)

// returns the name of the enum given an ordinal number
func (e *SmtpType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PLAIN",
    "STARTTLS",
    "SSL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SmtpType) index(name string) SmtpType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PLAIN",
    "STARTTLS",
    "SSL",
  }
  for idx := range names {
    if names[idx] == name {
      return SmtpType(idx)
    }
  }
  return SMTPTYPE_UNKNOWN
}

func (e *SmtpType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SmtpType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SmtpType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SmtpType) Ref() *SmtpType {
  return &e
}


/**
SNMP user authentication type
*/
type SnmpAuthType int

const(
  SNMPAUTHTYPE_UNKNOWN SnmpAuthType = 0
  SNMPAUTHTYPE_REDACTED SnmpAuthType = 1
  SNMPAUTHTYPE_MD5 SnmpAuthType = 2
  SNMPAUTHTYPE_SHA SnmpAuthType = 3
)

// returns the name of the enum given an ordinal number
func (e *SnmpAuthType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MD5",
    "SHA",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnmpAuthType) index(name string) SnmpAuthType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "MD5",
    "SHA",
  }
  for idx := range names {
    if names[idx] == name {
      return SnmpAuthType(idx)
    }
  }
  return SNMPAUTHTYPE_UNKNOWN
}

func (e *SnmpAuthType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnmpAuthType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnmpAuthType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnmpAuthType) Ref() *SnmpAuthType {
  return &e
}


/**
SNMP information
*/
type SnmpConfig struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SNMP status
  */
  Status *bool `json:"status,omitempty"`
  /**
  SNMP transport details
  */
  Transports []SnmpTransport `json:"transports,omitempty"`
  /**
  SNMP trap details
  */
  Traps []SnmpTrap `json:"traps,omitempty"`
  /**
  SNMP user information
  */
  Users []SnmpUser `json:"users,omitempty"`
}


func NewSnmpConfig() *SnmpConfig {
  p := new(SnmpConfig)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SnmpConfig"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SnmpConfig"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SNMP user encryption type
*/
type SnmpPrivType int

const(
  SNMPPRIVTYPE_UNKNOWN SnmpPrivType = 0
  SNMPPRIVTYPE_REDACTED SnmpPrivType = 1
  SNMPPRIVTYPE_DES SnmpPrivType = 2
  SNMPPRIVTYPE_AES SnmpPrivType = 3
)

// returns the name of the enum given an ordinal number
func (e *SnmpPrivType) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DES",
    "AES",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnmpPrivType) index(name string) SnmpPrivType {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "DES",
    "AES",
  }
  for idx := range names {
    if names[idx] == name {
      return SnmpPrivType(idx)
    }
  }
  return SNMPPRIVTYPE_UNKNOWN
}

func (e *SnmpPrivType) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnmpPrivType:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnmpPrivType) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnmpPrivType) Ref() *SnmpPrivType {
  return &e
}


/**
SNMP protocol type
*/
type SnmpProtocol int

const(
  SNMPPROTOCOL_UNKNOWN SnmpProtocol = 0
  SNMPPROTOCOL_REDACTED SnmpProtocol = 1
  SNMPPROTOCOL_UDP SnmpProtocol = 2
  SNMPPROTOCOL_UDP6 SnmpProtocol = 3
  SNMPPROTOCOL_TCP SnmpProtocol = 4
  SNMPPROTOCOL_TCP6 SnmpProtocol = 5
)

// returns the name of the enum given an ordinal number
func (e *SnmpProtocol) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UDP",
    "UDP6",
    "TCP",
    "TCP6",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnmpProtocol) index(name string) SnmpProtocol {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "UDP",
    "UDP6",
    "TCP",
    "TCP6",
  }
  for idx := range names {
    if names[idx] == name {
      return SnmpProtocol(idx)
    }
  }
  return SNMPPROTOCOL_UNKNOWN
}

func (e *SnmpProtocol) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnmpProtocol:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnmpProtocol) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnmpProtocol) Ref() *SnmpProtocol {
  return &e
}


/**
SNMP transport details
*/
type SnmpTransport struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SNMP port
  */
  Port *int `json:"port,omitempty"`
  
  Protocol *SnmpProtocol `json:"protocol,omitempty"`
}


func NewSnmpTransport() *SnmpTransport {
  p := new(SnmpTransport)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SnmpTransport"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SnmpTransport"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SNMP trap details
*/
type SnmpTrap struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SNMP trap address
  */
  Address *string `json:"address,omitempty"`
  /**
  Plaintext community string for SNMP version 2.0
  */
  CommunityString *string `json:"communityString,omitempty"`
  /**
  SNMP engine Id
  */
  EngineId *string `json:"engineId,omitempty"`
  /**
  SNMP information status
  */
  Inform *bool `json:"inform,omitempty"`
  /**
  SNMP port
  */
  Port *int `json:"port,omitempty"`
  
  Protocol *SnmpProtocol `json:"protocol,omitempty"`
  /**
  SNMP receiver name
  */
  RecieverName *string `json:"recieverName,omitempty"`
  /**
  SNMP user name
  */
  Username *string `json:"username,omitempty"`
  
  Version *SnmpTrapVersion `json:"version,omitempty"`
}


func NewSnmpTrap() *SnmpTrap {
  p := new(SnmpTrap)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SnmpTrap"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SnmpTrap"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
SNMP version
*/
type SnmpTrapVersion int

const(
  SNMPTRAPVERSION_UNKNOWN SnmpTrapVersion = 0
  SNMPTRAPVERSION_REDACTED SnmpTrapVersion = 1
  SNMPTRAPVERSION_V2 SnmpTrapVersion = 2
  SNMPTRAPVERSION_V3 SnmpTrapVersion = 3
)

// returns the name of the enum given an ordinal number
func (e *SnmpTrapVersion) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "V2",
    "V3",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SnmpTrapVersion) index(name string) SnmpTrapVersion {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "V2",
    "V3",
  }
  for idx := range names {
    if names[idx] == name {
      return SnmpTrapVersion(idx)
    }
  }
  return SNMPTRAPVERSION_UNKNOWN
}

func (e *SnmpTrapVersion) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SnmpTrapVersion:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SnmpTrapVersion) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SnmpTrapVersion) Ref() *SnmpTrapVersion {
  return &e
}


/**
SNMP user information
*/
type SnmpUser struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  SNMP user authentication key
  */
  AuthKey *string `json:"authKey,omitempty"`
  
  AuthType *SnmpAuthType `json:"authType,omitempty"`
  /**
  SNMP user encryption key
  */
  PrivKey *string `json:"privKey,omitempty"`
  
  PrivType *SnmpPrivType `json:"privType,omitempty"`
  /**
  SNMP user name
  */
  Username *string `json:"username,omitempty"`
}


func NewSnmpUser() *SnmpUser {
  p := new(SnmpUser)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SnmpUser"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SnmpUser"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Cluster software version details
*/
type SoftwareMapReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  
  SoftwareType *SoftwareTypeRef `json:"softwareType,omitempty"`
  /**
  Software version
  */
  Version *string `json:"version,omitempty"`
}


func NewSoftwareMapReference() *SoftwareMapReference {
  p := new(SoftwareMapReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.SoftwareMapReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.SoftwareMapReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}




/**
Software type
*/
type SoftwareTypeRef int

const(
  SOFTWARETYPEREF_UNKNOWN SoftwareTypeRef = 0
  SOFTWARETYPEREF_REDACTED SoftwareTypeRef = 1
  SOFTWARETYPEREF_NOS SoftwareTypeRef = 2
  SOFTWARETYPEREF_NCC SoftwareTypeRef = 3
  SOFTWARETYPEREF_PRISM_CENTRAL SoftwareTypeRef = 4
)

// returns the name of the enum given an ordinal number
func (e *SoftwareTypeRef) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NOS",
    "NCC",
    "PRISM_CENTRAL",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *SoftwareTypeRef) index(name string) SoftwareTypeRef {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "NOS",
    "NCC",
    "PRISM_CENTRAL",
  }
  for idx := range names {
    if names[idx] == name {
      return SoftwareTypeRef(idx)
    }
  }
  return SOFTWARETYPEREF_UNKNOWN
}

func (e *SoftwareTypeRef) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for SoftwareTypeRef:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *SoftwareTypeRef) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e SoftwareTypeRef) Ref() *SoftwareTypeRef {
  return &e
}


/**
Disk storage Tier type
*/
type StorageTierReference int

const(
  STORAGETIERREFERENCE_UNKNOWN StorageTierReference = 0
  STORAGETIERREFERENCE_REDACTED StorageTierReference = 1
  STORAGETIERREFERENCE_PCIE_SSD StorageTierReference = 2
  STORAGETIERREFERENCE_SATA_SSD StorageTierReference = 3
  STORAGETIERREFERENCE_HDD StorageTierReference = 4
)

// returns the name of the enum given an ordinal number
func (e *StorageTierReference) name(index int) string {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PCIE_SSD",
    "SATA_SSD",
    "HDD",
  }
  if index < 0 || index >= len(names) {
    return "$UNKNOWN"
  }
  return names[index]
}
// returns the enum type given a string value
func (e *StorageTierReference) index(name string) StorageTierReference {
  names := [...]string {
    "$UNKNOWN",
    "$REDACTED",
    "PCIE_SSD",
    "SATA_SSD",
    "HDD",
  }
  for idx := range names {
    if names[idx] == name {
      return StorageTierReference(idx)
    }
  }
  return STORAGETIERREFERENCE_UNKNOWN
}

func (e *StorageTierReference) UnmarshalJSON(b []byte) error {
  var enumStr string
  if err := json.Unmarshal(b, &enumStr); err != nil {
    return errors.New(fmt.Sprintf("Unable to unmarshal for StorageTierReference:%s", err))
  }
  *e = e.index(enumStr)
  return nil
}

func (e *StorageTierReference) MarshalJSON() ([]byte, error) {
  b := bytes.NewBufferString(`"`)
  b.WriteString(e.name(int(*e)))
  b.WriteString(`"`)
  return b.Bytes(), nil
}

func (e StorageTierReference) Ref() *StorageTierReference {
  return &e
}


/**
Message contains the component domain fault tolerance text details
*/
type ToleranceMessage struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  List of tolerance message attributes
  */
  AttributeList []AttributeItem `json:"attributeList,omitempty"`
  /**
  Message Id
  */
  Id *string `json:"id,omitempty"`
}


func NewToleranceMessage() *ToleranceMessage {
  p := new(ToleranceMessage)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "clustermgmt.v4.config.ToleranceMessage"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a1.config.ToleranceMessage"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



type OneOfGetHostsResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []HostEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostsResponseData() *OneOfGetHostsResponseData {
  p := new(OneOfGetHostsResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHostsResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHostsResponseData is nil"))
  }
  switch v.(type) {
    case []HostEntity:
      p.oneOfType0 = v.([]HostEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHostsResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.HostEntity>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetHostsResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]HostEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.HostEntity" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHostsResponseData"))
}

func (p *OneOfGetHostsResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.HostEntity>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetHostsResponseData")
}

type OneOfGetClusterResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *ClusterEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClusterResponseData() *OneOfGetClusterResponseData {
  p := new(OneOfGetClusterResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetClusterResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetClusterResponseData is nil"))
  }
  switch v.(type) {
    case ClusterEntity:
      if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterEntity)}
      *p.oneOfType0 = v.(ClusterEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetClusterResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetClusterResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(ClusterEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "clustermgmt.v4.config.ClusterEntity" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(ClusterEntity)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterResponseData"))
}

func (p *OneOfGetClusterResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetClusterResponseData")
}

type OneOfGetDomainFaultToleranceResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []DomainFaultTolerance `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetDomainFaultToleranceResponseData() *OneOfGetDomainFaultToleranceResponseData {
  p := new(OneOfGetDomainFaultToleranceResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetDomainFaultToleranceResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetDomainFaultToleranceResponseData is nil"))
  }
  switch v.(type) {
    case []DomainFaultTolerance:
      p.oneOfType0 = v.([]DomainFaultTolerance)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.DomainFaultTolerance>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.DomainFaultTolerance>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetDomainFaultToleranceResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.DomainFaultTolerance>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetDomainFaultToleranceResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]DomainFaultTolerance)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.DomainFaultTolerance" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.DomainFaultTolerance>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.DomainFaultTolerance>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDomainFaultToleranceResponseData"))
}

func (p *OneOfGetDomainFaultToleranceResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.DomainFaultTolerance>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetDomainFaultToleranceResponseData")
}

type OneOfGetClusterHostGpusResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []HostGpuEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClusterHostGpusResponseData() *OneOfGetClusterHostGpusResponseData {
  p := new(OneOfGetClusterHostGpusResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetClusterHostGpusResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetClusterHostGpusResponseData is nil"))
  }
  switch v.(type) {
    case []HostGpuEntity:
      p.oneOfType0 = v.([]HostGpuEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetClusterHostGpusResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetClusterHostGpusResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]HostGpuEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.HostGpuEntity" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClusterHostGpusResponseData"))
}

func (p *OneOfGetClusterHostGpusResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetClusterHostGpusResponseData")
}

type OneOfGetHostGpuResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *HostGpuEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostGpuResponseData() *OneOfGetHostGpuResponseData {
  p := new(OneOfGetHostGpuResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHostGpuResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHostGpuResponseData is nil"))
  }
  switch v.(type) {
    case HostGpuEntity:
      if nil == p.oneOfType0 {p.oneOfType0 = new(HostGpuEntity)}
      *p.oneOfType0 = v.(HostGpuEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHostGpuResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetHostGpuResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(HostGpuEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "clustermgmt.v4.config.HostGpuEntity" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(HostGpuEntity)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHostGpuResponseData"))
}

func (p *OneOfGetHostGpuResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetHostGpuResponseData")
}

type OneOfGetHostGpusResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []HostGpuEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostGpusResponseData() *OneOfGetHostGpusResponseData {
  p := new(OneOfGetHostGpusResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHostGpusResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHostGpusResponseData is nil"))
  }
  switch v.(type) {
    case []HostGpuEntity:
      p.oneOfType0 = v.([]HostGpuEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHostGpusResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetHostGpusResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]HostGpuEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.HostGpuEntity" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHostGpusResponseData"))
}

func (p *OneOfGetHostGpusResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetHostGpusResponseData")
}

type OneOfGetRackableUnitsResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []RackableUnit `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRackableUnitsResponseData() *OneOfGetRackableUnitsResponseData {
  p := new(OneOfGetRackableUnitsResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetRackableUnitsResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetRackableUnitsResponseData is nil"))
  }
  switch v.(type) {
    case []RackableUnit:
      p.oneOfType0 = v.([]RackableUnit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.RackableUnit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.RackableUnit>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetRackableUnitsResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.RackableUnit>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetRackableUnitsResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]RackableUnit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.RackableUnit" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.RackableUnit>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.RackableUnit>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRackableUnitsResponseData"))
}

func (p *OneOfGetRackableUnitsResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.RackableUnit>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetRackableUnitsResponseData")
}

type OneOfGetRsyslogServersResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []RsyslogServer `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRsyslogServersResponseData() *OneOfGetRsyslogServersResponseData {
  p := new(OneOfGetRsyslogServersResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetRsyslogServersResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetRsyslogServersResponseData is nil"))
  }
  switch v.(type) {
    case []RsyslogServer:
      p.oneOfType0 = v.([]RsyslogServer)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.RsyslogServer>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.RsyslogServer>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetRsyslogServersResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.RsyslogServer>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetRsyslogServersResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]RsyslogServer)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.RsyslogServer" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.RsyslogServer>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.RsyslogServer>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRsyslogServersResponseData"))
}

func (p *OneOfGetRsyslogServersResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.RsyslogServer>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetRsyslogServersResponseData")
}

type OneOfGetClustersResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []ClusterEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClustersResponseData() *OneOfGetClustersResponseData {
  p := new(OneOfGetClustersResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetClustersResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetClustersResponseData is nil"))
  }
  switch v.(type) {
    case []ClusterEntity:
      p.oneOfType0 = v.([]ClusterEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.ClusterEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.ClusterEntity>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetClustersResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.ClusterEntity>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetClustersResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]ClusterEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.ClusterEntity" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.ClusterEntity>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.ClusterEntity>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetClustersResponseData"))
}

func (p *OneOfGetClustersResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.ClusterEntity>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetClustersResponseData")
}

type OneOfGetGpuProfilesResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 []GpuProfile `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetGpuProfilesResponseData() *OneOfGetGpuProfilesResponseData {
  p := new(OneOfGetGpuProfilesResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetGpuProfilesResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetGpuProfilesResponseData is nil"))
  }
  switch v.(type) {
    case []GpuProfile:
      p.oneOfType0 = v.([]GpuProfile)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.GpuProfile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.GpuProfile>"
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetGpuProfilesResponseData) GetValue() interface{} {
  if "List<clustermgmt.v4.config.GpuProfile>" == *p.Discriminator {
    return p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetGpuProfilesResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new([]GpuProfile)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.GpuProfile" == *((*vOneOfType0)[0].ObjectType_) {
      p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = "List<clustermgmt.v4.config.GpuProfile>"
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = "List<clustermgmt.v4.config.GpuProfile>"
      return nil

    }
  }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetGpuProfilesResponseData"))
}

func (p *OneOfGetGpuProfilesResponseData) MarshalJSON() ([]byte, error) {
  if "List<clustermgmt.v4.config.GpuProfile>" == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetGpuProfilesResponseData")
}

type OneOfGetHostResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *HostEntity `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostResponseData() *OneOfGetHostResponseData {
  p := new(OneOfGetHostResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetHostResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetHostResponseData is nil"))
  }
  switch v.(type) {
    case HostEntity:
      if nil == p.oneOfType0 {p.oneOfType0 = new(HostEntity)}
      *p.oneOfType0 = v.(HostEntity)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetHostResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetHostResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(HostEntity)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "clustermgmt.v4.config.HostEntity" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(HostEntity)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHostResponseData"))
}

func (p *OneOfGetHostResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetHostResponseData")
}

type OneOfGetSnmpResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *SnmpConfig `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSnmpResponseData() *OneOfGetSnmpResponseData {
  p := new(OneOfGetSnmpResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetSnmpResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetSnmpResponseData is nil"))
  }
  switch v.(type) {
    case SnmpConfig:
      if nil == p.oneOfType0 {p.oneOfType0 = new(SnmpConfig)}
      *p.oneOfType0 = v.(SnmpConfig)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetSnmpResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetSnmpResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(SnmpConfig)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "clustermgmt.v4.config.SnmpConfig" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(SnmpConfig)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSnmpResponseData"))
}

func (p *OneOfGetSnmpResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetSnmpResponseData")
}

type OneOfGetRackableUnitResponseData struct {
  Discriminator *string `json:"-"`
  ObjectType_ *string `json:"-"`
  oneOfType0 *RackableUnit `json:"-"`
  oneOfType400 *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRackableUnitResponseData() *OneOfGetRackableUnitResponseData {
  p := new(OneOfGetRackableUnitResponseData)
  p.Discriminator = new(string)
  p.ObjectType_ = new(string)
  return p
}

func (p *OneOfGetRackableUnitResponseData) SetValue (v interface {}) error {
  if nil == p {
    return errors.New(fmt.Sprintf("OneOfGetRackableUnitResponseData is nil"))
  }
  switch v.(type) {
    case RackableUnit:
      if nil == p.oneOfType0 {p.oneOfType0 = new(RackableUnit)}
      *p.oneOfType0 = v.(RackableUnit)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
    case import2.ErrorResponse:
      if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = v.(import2.ErrorResponse)
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
    default:
      return errors.New(fmt.Sprintf("%T(%v) is not expected type", v,v))
  }
  return nil
}

func (p *OneOfGetRackableUnitResponseData) GetValue() interface{} {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return *p.oneOfType0
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return *p.oneOfType400
  }
  return nil
}

func (p *OneOfGetRackableUnitResponseData) UnmarshalJSON(b []byte) error {
  vOneOfType0 := new(RackableUnit)
  if err := json.Unmarshal(b, vOneOfType0); err == nil {
    if "clustermgmt.v4.config.RackableUnit" == *vOneOfType0.ObjectType_ {
          if nil == p.oneOfType0 {p.oneOfType0 = new(RackableUnit)}
      *p.oneOfType0 = *vOneOfType0
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType0.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType0.ObjectType_
      return nil
    }
    }
  vOneOfType400 := new(import2.ErrorResponse)
  if err := json.Unmarshal(b, vOneOfType400); err == nil {
    if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
          if nil == p.oneOfType400 {p.oneOfType400 = new(import2.ErrorResponse)}
      *p.oneOfType400 = *vOneOfType400
      if nil == p.Discriminator {p.Discriminator = new(string)}
      *p.Discriminator = *p.oneOfType400.ObjectType_
      if nil == p.ObjectType_ {p.ObjectType_ = new(string)}
      *p.ObjectType_ = *p.oneOfType400.ObjectType_
      return nil
    }
    }
  return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRackableUnitResponseData"))
}

func (p *OneOfGetRackableUnitResponseData) MarshalJSON() ([]byte, error) {
  if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType0)
  }
  if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
    return json.Marshal(p.oneOfType400)
  }
  return nil, errors.New("No value to marshal for OneOfGetRackableUnitResponseData")
}

