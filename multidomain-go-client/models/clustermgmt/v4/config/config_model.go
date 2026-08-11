/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.4.1-beta-1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.config of Nutanix Multidomain Versioned APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
	"time"
)

/*
Params associated to the backplane network segmentation. This is part of payload for cluster create operation only.
*/
type BackplaneNetworkParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag to indicate if the backplane segmentation needs to be enabled or not.
	*/
	IsSegmentationEnabled *bool `json:"isSegmentationEnabled,omitempty"`

	Netmask *import1.IPv4Address `json:"netmask,omitempty"`

	Subnet *import1.IPv4Address `json:"subnet,omitempty"`
	/*
	  VLAN Id tagged to the backplane network on the cluster. This is part of cluster create payload.
	*/
	VlanTag *int64 `json:"vlanTag,omitempty"`
}

func (p *BackplaneNetworkParams) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BackplaneNetworkParams

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

func (p *BackplaneNetworkParams) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BackplaneNetworkParams
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBackplaneNetworkParams()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsSegmentationEnabled != nil {
		p.IsSegmentationEnabled = known.IsSegmentationEnabled
	}
	if known.Netmask != nil {
		p.Netmask = known.Netmask
	}
	if known.Subnet != nil {
		p.Subnet = known.Subnet
	}
	if known.VlanTag != nil {
		p.VlanTag = known.VlanTag
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isSegmentationEnabled")
	delete(allFields, "netmask")
	delete(allFields, "subnet")
	delete(allFields, "vlanTag")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBackplaneNetworkParams() *BackplaneNetworkParams {
	p := new(BackplaneNetworkParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BackplaneNetworkParams"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Currently representing the build information to be used for the cluster creation.
*/
type BuildInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Software version.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *BuildInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BuildInfo

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

func (p *BuildInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BuildInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBuildInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBuildInfo() *BuildInfo {
	p := new(BuildInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BuildInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Build information details.
*/
type BuildReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Software build type.
	*/
	BuildType *string `json:"buildType,omitempty"`
	/*
	  Commit Id used for version.
	*/
	CommitId *string `json:"commitId,omitempty"`
	/*
	  Full name of software version.
	*/
	FullVersion *string `json:"fullVersion,omitempty"`
	/*
	  Short commit Id used for version.
	*/
	ShortCommitId *string `json:"shortCommitId,omitempty"`
	/*
	  Software version.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *BuildReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BuildReference

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

func (p *BuildReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BuildReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBuildReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BuildType != nil {
		p.BuildType = known.BuildType
	}
	if known.CommitId != nil {
		p.CommitId = known.CommitId
	}
	if known.FullVersion != nil {
		p.FullVersion = known.FullVersion
	}
	if known.ShortCommitId != nil {
		p.ShortCommitId = known.ShortCommitId
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "buildType")
	delete(allFields, "commitId")
	delete(allFields, "fullVersion")
	delete(allFields, "shortCommitId")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBuildReference() *BuildReference {
	p := new(BuildReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BuildReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster entity with its attributes.
*/
type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Score to indicate how much cluster is eligible for storing domain manager backup.
	*/
	BackupEligibilityScore *int64 `json:"backupEligibilityScore,omitempty"`
	/*
	  List of categories associated to the PE cluster.
	*/
	Categories []string `json:"categories,omitempty"`
	/*
	  UUID of the cluster profile.
	*/
	ClusterProfileExtId *string `json:"clusterProfileExtId,omitempty"`

	Config *ClusterConfigReference `json:"config,omitempty"`
	/*
	  The name of the default container created as part of cluster creation. This is part of payload for cluster create operation only.
	*/
	ContainerName *string `json:"containerName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Number of inefficient VMs.
	*/
	InefficientVmCount *int64 `json:"inefficientVmCount,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Cluster name. This is part of payload for both cluster create &amp; update operations.
	*/
	Name *string `json:"name,omitempty"`

	Network *ClusterNetworkReference `json:"network,omitempty"`

	Nodes *NodeReference `json:"nodes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UpgradeStatus *UpgradeStatus `json:"upgradeStatus,omitempty"`
	/*
	  Number of VMs.
	*/
	VmCount *int64 `json:"vmCount,omitempty"`
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
	if known.BackupEligibilityScore != nil {
		p.BackupEligibilityScore = known.BackupEligibilityScore
	}
	if known.Categories != nil {
		p.Categories = known.Categories
	}
	if known.ClusterProfileExtId != nil {
		p.ClusterProfileExtId = known.ClusterProfileExtId
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.ContainerName != nil {
		p.ContainerName = known.ContainerName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.InefficientVmCount != nil {
		p.InefficientVmCount = known.InefficientVmCount
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.Nodes != nil {
		p.Nodes = known.Nodes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpgradeStatus != nil {
		p.UpgradeStatus = known.UpgradeStatus
	}
	if known.VmCount != nil {
		p.VmCount = known.VmCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backupEligibilityScore")
	delete(allFields, "categories")
	delete(allFields, "clusterProfileExtId")
	delete(allFields, "config")
	delete(allFields, "containerName")
	delete(allFields, "extId")
	delete(allFields, "inefficientVmCount")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "network")
	delete(allFields, "nodes")
	delete(allFields, "tenantId")
	delete(allFields, "upgradeStatus")
	delete(allFields, "vmCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster arch.
*/
type ClusterArchReference int

const (
	CLUSTERARCHREFERENCE_UNKNOWN  ClusterArchReference = 0
	CLUSTERARCHREFERENCE_REDACTED ClusterArchReference = 1
	CLUSTERARCHREFERENCE_X86_64   ClusterArchReference = 2
	CLUSTERARCHREFERENCE_PPC64LE  ClusterArchReference = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterArchReference) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e ClusterArchReference) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *ClusterArchReference) index(name string) ClusterArchReference {
	names := [...]string{
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

/*
Cluster Configuration required for a cluster to function properly.
*/
type ClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`
	/*
	  A boolean value indicating whether to enable lockdown mode for a cluster.
	*/
	ShouldEnableLockdownMode *bool `json:"shouldEnableLockdownMode,omitempty"`
}

func (p *ClusterConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterConfig

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

func (p *ClusterConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BuildInfo != nil {
		p.BuildInfo = known.BuildInfo
	}
	if known.ShouldEnableLockdownMode != nil {
		p.ShouldEnableLockdownMode = known.ShouldEnableLockdownMode
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "buildInfo")
	delete(allFields, "shouldEnableLockdownMode")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterConfig() *ClusterConfig {
	p := new(ClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster configuration details.
*/
type ClusterConfigReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Public ssh key details. This is part of payload for cluster update operation only.
	*/
	AuthorizedPublicKeyList []PublicKey `json:"authorizedPublicKeyList,omitempty"`

	BuildInfo *BuildReference `json:"buildInfo,omitempty"`

	ClusterArch *ClusterArchReference `json:"clusterArch,omitempty"`
	/*
	  Cluster function. This is part of payload for cluster create operation only (allowed enum values for creation are AOS, ONE_NODE &amp; TWO_NODE only).
	*/
	ClusterFunction []ClusterFunctionRef `json:"clusterFunction,omitempty"`
	/*
	  Cluster software version details.
	*/
	ClusterSoftwareMap []SoftwareMapReference `json:"clusterSoftwareMap,omitempty"`

	ClusterType *ClusterType `json:"clusterType,omitempty"`

	EncryptionInTransitStatus *EncryptionStatus `json:"encryptionInTransitStatus,omitempty"`
	/*
	  Encryption option.
	*/
	EncryptionOption []EncryptionOptionInfo `json:"encryptionOption,omitempty"`
	/*
	  Encryption scope.
	*/
	EncryptionScope []EncryptionScopeInfo `json:"encryptionScope,omitempty"`

	FaultToleranceState *FaultToleranceState `json:"faultToleranceState,omitempty"`
	/*
	  Hypervisor type.
	*/
	HypervisorTypes []HypervisorType `json:"hypervisorTypes,omitempty"`
	/*
	  Cluster incarnation Id. This is part of payload for cluster update operation only.
	*/
	IncarnationId *int64 `json:"incarnationId,omitempty"`
	/*
	  Indicates if cluster is available to contact or not.
	*/
	IsAvailable *bool `json:"isAvailable,omitempty"`
	/*
	  Whether to disable monitoring of degraded nodes on the cluster.
	*/
	IsDegradedNodeMonitoringDisabled *bool `json:"isDegradedNodeMonitoringDisabled,omitempty"`
	/*
	  Indicates whether the release is categorized as Long-term or not.
	*/
	IsLts *bool `json:"isLts,omitempty"`
	/*
	  Indicates whether the password ssh into the cluster is enabled or not.
	*/
	IsPasswordRemoteLoginEnabled *bool `json:"isPasswordRemoteLoginEnabled,omitempty"`
	/*
	  Remote support status.
	*/
	IsRemoteSupportEnabled *bool `json:"isRemoteSupportEnabled,omitempty"`

	OperationMode *OperationMode `json:"operationMode,omitempty"`

	PulseStatus *PulseStatus `json:"pulseStatus,omitempty"`
	/*
	  Redundancy factor of a cluster. This is part of payload for both cluster create &amp; update operations.
	*/
	RedundancyFactor *int64 `json:"redundancyFactor,omitempty"`
	/*
	  Time zone on a cluster.
	*/
	Timezone *string `json:"timezone,omitempty"`
}

func (p *ClusterConfigReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterConfigReference

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

func (p *ClusterConfigReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterConfigReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterConfigReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthorizedPublicKeyList != nil {
		p.AuthorizedPublicKeyList = known.AuthorizedPublicKeyList
	}
	if known.BuildInfo != nil {
		p.BuildInfo = known.BuildInfo
	}
	if known.ClusterArch != nil {
		p.ClusterArch = known.ClusterArch
	}
	if known.ClusterFunction != nil {
		p.ClusterFunction = known.ClusterFunction
	}
	if known.ClusterSoftwareMap != nil {
		p.ClusterSoftwareMap = known.ClusterSoftwareMap
	}
	if known.ClusterType != nil {
		p.ClusterType = known.ClusterType
	}
	if known.EncryptionInTransitStatus != nil {
		p.EncryptionInTransitStatus = known.EncryptionInTransitStatus
	}
	if known.EncryptionOption != nil {
		p.EncryptionOption = known.EncryptionOption
	}
	if known.EncryptionScope != nil {
		p.EncryptionScope = known.EncryptionScope
	}
	if known.FaultToleranceState != nil {
		p.FaultToleranceState = known.FaultToleranceState
	}
	if known.HypervisorTypes != nil {
		p.HypervisorTypes = known.HypervisorTypes
	}
	if known.IncarnationId != nil {
		p.IncarnationId = known.IncarnationId
	}
	if known.IsAvailable != nil {
		p.IsAvailable = known.IsAvailable
	}
	if known.IsDegradedNodeMonitoringDisabled != nil {
		p.IsDegradedNodeMonitoringDisabled = known.IsDegradedNodeMonitoringDisabled
	}
	if known.IsLts != nil {
		p.IsLts = known.IsLts
	}
	if known.IsPasswordRemoteLoginEnabled != nil {
		p.IsPasswordRemoteLoginEnabled = known.IsPasswordRemoteLoginEnabled
	}
	if known.IsRemoteSupportEnabled != nil {
		p.IsRemoteSupportEnabled = known.IsRemoteSupportEnabled
	}
	if known.OperationMode != nil {
		p.OperationMode = known.OperationMode
	}
	if known.PulseStatus != nil {
		p.PulseStatus = known.PulseStatus
	}
	if known.RedundancyFactor != nil {
		p.RedundancyFactor = known.RedundancyFactor
	}
	if known.Timezone != nil {
		p.Timezone = known.Timezone
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authorizedPublicKeyList")
	delete(allFields, "buildInfo")
	delete(allFields, "clusterArch")
	delete(allFields, "clusterFunction")
	delete(allFields, "clusterSoftwareMap")
	delete(allFields, "clusterType")
	delete(allFields, "encryptionInTransitStatus")
	delete(allFields, "encryptionOption")
	delete(allFields, "encryptionScope")
	delete(allFields, "faultToleranceState")
	delete(allFields, "hypervisorTypes")
	delete(allFields, "incarnationId")
	delete(allFields, "isAvailable")
	delete(allFields, "isDegradedNodeMonitoringDisabled")
	delete(allFields, "isLts")
	delete(allFields, "isPasswordRemoteLoginEnabled")
	delete(allFields, "isRemoteSupportEnabled")
	delete(allFields, "operationMode")
	delete(allFields, "pulseStatus")
	delete(allFields, "redundancyFactor")
	delete(allFields, "timezone")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterConfigReference() *ClusterConfigReference {
	p := new(ClusterConfigReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterConfigReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster Fault tolerance. Set desiredClusterFaultTolerance for cluster create and update.
*/
type ClusterFaultToleranceRef int

const (
	CLUSTERFAULTTOLERANCEREF_UNKNOWN       ClusterFaultToleranceRef = 0
	CLUSTERFAULTTOLERANCEREF_REDACTED      ClusterFaultToleranceRef = 1
	CLUSTERFAULTTOLERANCEREF_CFT_0N_AND_0D ClusterFaultToleranceRef = 2
	CLUSTERFAULTTOLERANCEREF_CFT_1N_OR_1D  ClusterFaultToleranceRef = 3
	CLUSTERFAULTTOLERANCEREF_CFT_2N_OR_2D  ClusterFaultToleranceRef = 4
	CLUSTERFAULTTOLERANCEREF_CFT_1N_AND_1D ClusterFaultToleranceRef = 5
	CLUSTERFAULTTOLERANCEREF_CFT_1N_OR_2D  ClusterFaultToleranceRef = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterFaultToleranceRef) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterFaultToleranceRef) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterFaultToleranceRef) index(name string) ClusterFaultToleranceRef {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterFaultToleranceRef(idx)
		}
	}
	return CLUSTERFAULTTOLERANCEREF_UNKNOWN
}

func (e *ClusterFaultToleranceRef) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterFaultToleranceRef:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterFaultToleranceRef) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterFaultToleranceRef) Ref() *ClusterFaultToleranceRef {
	return &e
}

/*
Cluster function. This is part of payload for cluster create operation only (allowed enum values for creation are AOS, ONE_NODE &amp; TWO_NODE only).
*/
type ClusterFunctionRef int

const (
	CLUSTERFUNCTIONREF_UNKNOWN            ClusterFunctionRef = 0
	CLUSTERFUNCTIONREF_REDACTED           ClusterFunctionRef = 1
	CLUSTERFUNCTIONREF_AOS                ClusterFunctionRef = 2
	CLUSTERFUNCTIONREF_PRISM_CENTRAL      ClusterFunctionRef = 3
	CLUSTERFUNCTIONREF_CLOUD_DATA_GATEWAY ClusterFunctionRef = 4
	CLUSTERFUNCTIONREF_AFS                ClusterFunctionRef = 5
	CLUSTERFUNCTIONREF_ONE_NODE           ClusterFunctionRef = 6
	CLUSTERFUNCTIONREF_TWO_NODE           ClusterFunctionRef = 7
	CLUSTERFUNCTIONREF_ANALYTICS_PLATFORM ClusterFunctionRef = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterFunctionRef) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"PRISM_CENTRAL",
		"CLOUD_DATA_GATEWAY",
		"AFS",
		"ONE_NODE",
		"TWO_NODE",
		"ANALYTICS_PLATFORM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterFunctionRef) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"PRISM_CENTRAL",
		"CLOUD_DATA_GATEWAY",
		"AFS",
		"ONE_NODE",
		"TWO_NODE",
		"ANALYTICS_PLATFORM",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterFunctionRef) index(name string) ClusterFunctionRef {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AOS",
		"PRISM_CENTRAL",
		"CLOUD_DATA_GATEWAY",
		"AFS",
		"ONE_NODE",
		"TWO_NODE",
		"ANALYTICS_PLATFORM",
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

/*
Network details of a cluster.
*/
type ClusterNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExternalAddress *import1.IPAddress `json:"externalAddress,omitempty"`
	/*
	  Cluster fully qualified domain name. This is part of payload for cluster update operation only.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  List of HTTP Proxy server configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	HttpProxyConfig []HttpProxyConfig `json:"httpProxyConfig,omitempty"`
	/*
	  Targets HTTP traffic to which is exempted from going through the configured HTTP Proxy.
	*/
	HttpProxyWhiteListConfig []HttpProxyWhiteListConfig `json:"httpProxyWhiteListConfig,omitempty"`
	/*
	  List of name servers on a cluster. This is a part of payload for both clusters create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NameServers []import1.IPAddressOrFQDN `json:"nameServers,omitempty"`
	/*
	  List of NTP servers on a cluster. This is a part of payload for both cluster create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *ClusterNetwork) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterNetwork

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

func (p *ClusterNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExternalAddress != nil {
		p.ExternalAddress = known.ExternalAddress
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.HttpProxyConfig != nil {
		p.HttpProxyConfig = known.HttpProxyConfig
	}
	if known.HttpProxyWhiteListConfig != nil {
		p.HttpProxyWhiteListConfig = known.HttpProxyWhiteListConfig
	}
	if known.NameServers != nil {
		p.NameServers = known.NameServers
	}
	if known.NtpServers != nil {
		p.NtpServers = known.NtpServers
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "externalAddress")
	delete(allFields, "fqdn")
	delete(allFields, "httpProxyConfig")
	delete(allFields, "httpProxyWhiteListConfig")
	delete(allFields, "nameServers")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterNetwork() *ClusterNetwork {
	p := new(ClusterNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network details of a cluster.
*/
type ClusterNetworkReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Backplane *BackplaneNetworkParams `json:"backplane,omitempty"`

	ExternalAddress *import1.IPAddress `json:"externalAddress,omitempty"`

	ExternalDataServiceIp *import1.IPAddress `json:"externalDataServiceIp,omitempty"`
	/*
	  Cluster external subnet address.
	*/
	ExternalSubnet *string `json:"externalSubnet,omitempty"`
	/*
	  Cluster fully qualified domain name. This is part of payload for cluster update operation only.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  List of HTTP Proxy server configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	HttpProxyList []HttpProxyConfig `json:"httpProxyList,omitempty"`
	/*
	  Targets HTTP traffic to which is exempted from going through the configured HTTP Proxy.
	*/
	HttpProxyWhiteList []HttpProxyWhiteListConfig `json:"httpProxyWhiteList,omitempty"`
	/*
	  Cluster internal subnet address.
	*/
	InternalSubnet *string `json:"internalSubnet,omitempty"`

	KeyManagementServerType *KeyManagementServerType `json:"keyManagementServerType,omitempty"`

	ManagementServer *ManagementServerRef `json:"managementServer,omitempty"`

	MasqueradingIp *import1.IPAddress `json:"masqueradingIp,omitempty"`
	/*
	  The port to connect to the cluster when using masquerading IP.
	*/
	MasqueradingPort *int `json:"masqueradingPort,omitempty"`
	/*
	  List of name servers on a cluster. This is a part of payload for both clusters create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NameServerIpList []import1.IPAddressOrFQDN `json:"nameServerIpList,omitempty"`
	/*
	  NFS subnet allowlist addresses. This is part of the payload for cluster update operation only.
	*/
	NfsSubnetWhitelist []string `json:"nfsSubnetWhitelist,omitempty"`
	/*
	  List of NTP server configurations.
	*/
	NtpServerConfigList []NtpServerConfig `json:"ntpServerConfigList,omitempty"`
	/*
	  List of NTP servers on a cluster. This is a part of payload for both cluster create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NtpServerIpList []import1.IPAddressOrFQDN `json:"ntpServerIpList,omitempty"`

	SmtpServer *SmtpServerRef `json:"smtpServer,omitempty"`
}

func (p *ClusterNetworkReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterNetworkReference

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

func (p *ClusterNetworkReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterNetworkReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterNetworkReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Backplane != nil {
		p.Backplane = known.Backplane
	}
	if known.ExternalAddress != nil {
		p.ExternalAddress = known.ExternalAddress
	}
	if known.ExternalDataServiceIp != nil {
		p.ExternalDataServiceIp = known.ExternalDataServiceIp
	}
	if known.ExternalSubnet != nil {
		p.ExternalSubnet = known.ExternalSubnet
	}
	if known.Fqdn != nil {
		p.Fqdn = known.Fqdn
	}
	if known.HttpProxyList != nil {
		p.HttpProxyList = known.HttpProxyList
	}
	if known.HttpProxyWhiteList != nil {
		p.HttpProxyWhiteList = known.HttpProxyWhiteList
	}
	if known.InternalSubnet != nil {
		p.InternalSubnet = known.InternalSubnet
	}
	if known.KeyManagementServerType != nil {
		p.KeyManagementServerType = known.KeyManagementServerType
	}
	if known.ManagementServer != nil {
		p.ManagementServer = known.ManagementServer
	}
	if known.MasqueradingIp != nil {
		p.MasqueradingIp = known.MasqueradingIp
	}
	if known.MasqueradingPort != nil {
		p.MasqueradingPort = known.MasqueradingPort
	}
	if known.NameServerIpList != nil {
		p.NameServerIpList = known.NameServerIpList
	}
	if known.NfsSubnetWhitelist != nil {
		p.NfsSubnetWhitelist = known.NfsSubnetWhitelist
	}
	if known.NtpServerConfigList != nil {
		p.NtpServerConfigList = known.NtpServerConfigList
	}
	if known.NtpServerIpList != nil {
		p.NtpServerIpList = known.NtpServerIpList
	}
	if known.SmtpServer != nil {
		p.SmtpServer = known.SmtpServer
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backplane")
	delete(allFields, "externalAddress")
	delete(allFields, "externalDataServiceIp")
	delete(allFields, "externalSubnet")
	delete(allFields, "fqdn")
	delete(allFields, "httpProxyList")
	delete(allFields, "httpProxyWhiteList")
	delete(allFields, "internalSubnet")
	delete(allFields, "keyManagementServerType")
	delete(allFields, "managementServer")
	delete(allFields, "masqueradingIp")
	delete(allFields, "masqueradingPort")
	delete(allFields, "nameServerIpList")
	delete(allFields, "nfsSubnetWhitelist")
	delete(allFields, "ntpServerConfigList")
	delete(allFields, "ntpServerIpList")
	delete(allFields, "smtpServer")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterNetworkReference() *ClusterNetworkReference {
	p := new(ClusterNetworkReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterNetworkReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterProfile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if a configuration of attached clusters can be skipped from monitoring.
	*/
	AllowedOverrides []ConfigType `json:"allowedOverrides,omitempty"`
	/*
	  Count of clusters associated to a cluster profile.
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  Managed cluster information.
	*/
	Clusters []ManagedCluster `json:"clusters,omitempty"`
	/*
	  Creation time of a cluster profile.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  Details of the user who created the cluster profile.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Detailed description of a cluster profile.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The count indicates the number of clusters associated with a cluster profile that has experienced drift. Drifted clusters are those in which the configuration differs from the defined profile. For example, the NTP server has different values on a cluster as compared to the profile it is attached.
	*/
	DriftedClusterCount *int `json:"driftedClusterCount,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FaultToleranceConfig *FaultToleranceConfiguration `json:"faultToleranceConfig,omitempty"`

	HttpProxyConfig *HttpProxyConfiguration `json:"httpProxyConfig,omitempty"`
	/*
	  The last updated time of a cluster profile.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  Details of the user who has recently updated the cluster profile.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the cluster profile.
	*/
	Name *string `json:"name"`
	/*
	  List of name servers on a cluster. This is a part of payload for both clusters create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NameServerIpList []import1.IPAddress `json:"nameServerIpList,omitempty"`
	/*
	  NFS subnet allowlist addresses. This is part of the payload for cluster update operation only.
	*/
	NfsSubnetWhitelist []string `json:"nfsSubnetWhitelist,omitempty"`
	/*
	  List of NTP server configurations.
	*/
	NtpServerConfigList []NtpServerConfig `json:"ntpServerConfigList,omitempty"`
	/*
	  List of NTP servers on a cluster. This is a part of payload for both cluster create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NtpServerIpList []import1.IPAddressOrFQDN `json:"ntpServerIpList,omitempty"`

	PulseStatus *PulseStatus `json:"pulseStatus,omitempty"`

	RebuildReservationConfig *RebuildReservationConfig `json:"rebuildReservationConfig,omitempty"`

	ResilientCapacityWarningThresholdConfig *ResilientCapacityWarningThresholdConfig `json:"resilientCapacityWarningThresholdConfig,omitempty"`
	/*
	  RSYSLOG Server.
	*/
	RsyslogServerList []RsyslogServer `json:"rsyslogServerList,omitempty"`

	SmtpServer *SmtpServerRef `json:"smtpServer,omitempty"`

	SnmpConfig *SnmpConfig `json:"snmpConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterProfile) MarshalJSON() ([]byte, error) {
	type ClusterProfileProxy ClusterProfile

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterProfileProxy
		Name *string `json:"name,omitempty"`
	}{
		ClusterProfileProxy: (*ClusterProfileProxy)(p),
		Name:                p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ClusterProfile) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterProfile
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterProfile()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AllowedOverrides != nil {
		p.AllowedOverrides = known.AllowedOverrides
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.Clusters != nil {
		p.Clusters = known.Clusters
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DriftedClusterCount != nil {
		p.DriftedClusterCount = known.DriftedClusterCount
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FaultToleranceConfig != nil {
		p.FaultToleranceConfig = known.FaultToleranceConfig
	}
	if known.HttpProxyConfig != nil {
		p.HttpProxyConfig = known.HttpProxyConfig
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NameServerIpList != nil {
		p.NameServerIpList = known.NameServerIpList
	}
	if known.NfsSubnetWhitelist != nil {
		p.NfsSubnetWhitelist = known.NfsSubnetWhitelist
	}
	if known.NtpServerConfigList != nil {
		p.NtpServerConfigList = known.NtpServerConfigList
	}
	if known.NtpServerIpList != nil {
		p.NtpServerIpList = known.NtpServerIpList
	}
	if known.PulseStatus != nil {
		p.PulseStatus = known.PulseStatus
	}
	if known.RebuildReservationConfig != nil {
		p.RebuildReservationConfig = known.RebuildReservationConfig
	}
	if known.ResilientCapacityWarningThresholdConfig != nil {
		p.ResilientCapacityWarningThresholdConfig = known.ResilientCapacityWarningThresholdConfig
	}
	if known.RsyslogServerList != nil {
		p.RsyslogServerList = known.RsyslogServerList
	}
	if known.SmtpServer != nil {
		p.SmtpServer = known.SmtpServer
	}
	if known.SnmpConfig != nil {
		p.SnmpConfig = known.SnmpConfig
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "allowedOverrides")
	delete(allFields, "clusterCount")
	delete(allFields, "clusters")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "driftedClusterCount")
	delete(allFields, "extId")
	delete(allFields, "faultToleranceConfig")
	delete(allFields, "httpProxyConfig")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "nameServerIpList")
	delete(allFields, "nfsSubnetWhitelist")
	delete(allFields, "ntpServerConfigList")
	delete(allFields, "ntpServerIpList")
	delete(allFields, "pulseStatus")
	delete(allFields, "rebuildReservationConfig")
	delete(allFields, "resilientCapacityWarningThresholdConfig")
	delete(allFields, "rsyslogServerList")
	delete(allFields, "smtpServer")
	delete(allFields, "snmpConfig")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterProfile() *ClusterProfile {
	p := new(ClusterProfile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterProfile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterProfileProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates if a configuration of attached clusters can be skipped from monitoring.
	*/
	AllowedOverrides []ConfigType `json:"allowedOverrides,omitempty"`
	/*
	  Count of clusters associated to a cluster profile.
	*/
	ClusterCount *int `json:"clusterCount,omitempty"`
	/*
	  Managed cluster information.
	*/
	Clusters []ManagedCluster `json:"clusters,omitempty"`
	/*
	  Creation time of a cluster profile.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  Details of the user who created the cluster profile.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  Detailed description of a cluster profile.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  The count indicates the number of clusters associated with a cluster profile that has experienced drift. Drifted clusters are those in which the configuration differs from the defined profile. For example, the NTP server has different values on a cluster as compared to the profile it is attached.
	*/
	DriftedClusterCount *int `json:"driftedClusterCount,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	FaultToleranceConfig *FaultToleranceConfiguration `json:"faultToleranceConfig,omitempty"`

	HttpProxyConfig *HttpProxyConfiguration `json:"httpProxyConfig,omitempty"`
	/*
	  The last updated time of a cluster profile.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  Details of the user who has recently updated the cluster profile.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the cluster profile.
	*/
	Name *string `json:"name"`
	/*
	  List of name servers on a cluster. This is a part of payload for both clusters create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NameServerIpList []import1.IPAddress `json:"nameServerIpList,omitempty"`
	/*
	  NFS subnet allowlist addresses. This is part of the payload for cluster update operation only.
	*/
	NfsSubnetWhitelist []string `json:"nfsSubnetWhitelist,omitempty"`
	/*
	  List of NTP server configurations.
	*/
	NtpServerConfigList []NtpServerConfig `json:"ntpServerConfigList,omitempty"`
	/*
	  List of NTP servers on a cluster. This is a part of payload for both cluster create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NtpServerIpList []import1.IPAddressOrFQDN `json:"ntpServerIpList,omitempty"`

	PulseStatus *PulseStatus `json:"pulseStatus,omitempty"`

	RebuildReservationConfig *RebuildReservationConfig `json:"rebuildReservationConfig,omitempty"`

	ResilientCapacityWarningThresholdConfig *ResilientCapacityWarningThresholdConfig `json:"resilientCapacityWarningThresholdConfig,omitempty"`
	/*
	  RSYSLOG Server.
	*/
	RsyslogServerList []RsyslogServer `json:"rsyslogServerList,omitempty"`

	SmtpServer *SmtpServerRef `json:"smtpServer,omitempty"`

	SnmpConfig *SnmpConfig `json:"snmpConfig,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ClusterProfileProjection) MarshalJSON() ([]byte, error) {
	type ClusterProfileProjectionProxy ClusterProfileProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterProfileProjectionProxy
		Name *string `json:"name,omitempty"`
	}{
		ClusterProfileProjectionProxy: (*ClusterProfileProjectionProxy)(p),
		Name:                          p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ClusterProfileProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterProfileProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterProfileProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AllowedOverrides != nil {
		p.AllowedOverrides = known.AllowedOverrides
	}
	if known.ClusterCount != nil {
		p.ClusterCount = known.ClusterCount
	}
	if known.Clusters != nil {
		p.Clusters = known.Clusters
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.CreatedBy != nil {
		p.CreatedBy = known.CreatedBy
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.DriftedClusterCount != nil {
		p.DriftedClusterCount = known.DriftedClusterCount
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.FaultToleranceConfig != nil {
		p.FaultToleranceConfig = known.FaultToleranceConfig
	}
	if known.HttpProxyConfig != nil {
		p.HttpProxyConfig = known.HttpProxyConfig
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.LastUpdatedBy != nil {
		p.LastUpdatedBy = known.LastUpdatedBy
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NameServerIpList != nil {
		p.NameServerIpList = known.NameServerIpList
	}
	if known.NfsSubnetWhitelist != nil {
		p.NfsSubnetWhitelist = known.NfsSubnetWhitelist
	}
	if known.NtpServerConfigList != nil {
		p.NtpServerConfigList = known.NtpServerConfigList
	}
	if known.NtpServerIpList != nil {
		p.NtpServerIpList = known.NtpServerIpList
	}
	if known.PulseStatus != nil {
		p.PulseStatus = known.PulseStatus
	}
	if known.RebuildReservationConfig != nil {
		p.RebuildReservationConfig = known.RebuildReservationConfig
	}
	if known.ResilientCapacityWarningThresholdConfig != nil {
		p.ResilientCapacityWarningThresholdConfig = known.ResilientCapacityWarningThresholdConfig
	}
	if known.RsyslogServerList != nil {
		p.RsyslogServerList = known.RsyslogServerList
	}
	if known.SmtpServer != nil {
		p.SmtpServer = known.SmtpServer
	}
	if known.SnmpConfig != nil {
		p.SnmpConfig = known.SnmpConfig
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "allowedOverrides")
	delete(allFields, "clusterCount")
	delete(allFields, "clusters")
	delete(allFields, "createTime")
	delete(allFields, "createdBy")
	delete(allFields, "description")
	delete(allFields, "driftedClusterCount")
	delete(allFields, "extId")
	delete(allFields, "faultToleranceConfig")
	delete(allFields, "httpProxyConfig")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "lastUpdatedBy")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "nameServerIpList")
	delete(allFields, "nfsSubnetWhitelist")
	delete(allFields, "ntpServerConfigList")
	delete(allFields, "ntpServerIpList")
	delete(allFields, "pulseStatus")
	delete(allFields, "rebuildReservationConfig")
	delete(allFields, "resilientCapacityWarningThresholdConfig")
	delete(allFields, "rsyslogServerList")
	delete(allFields, "smtpServer")
	delete(allFields, "snmpConfig")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterProfileProjection() *ClusterProfileProjection {
	p := new(ClusterProfileProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterProfileProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Score to indicate how much cluster is eligible for storing domain manager backup.
	*/
	BackupEligibilityScore *int64 `json:"backupEligibilityScore,omitempty"`
	/*
	  List of categories associated to the PE cluster.
	*/
	Categories []string `json:"categories,omitempty"`
	/*
	  UUID of the cluster profile.
	*/
	ClusterProfileExtId *string `json:"clusterProfileExtId,omitempty"`

	ClusterProfileProjection *ClusterProfileProjection `json:"clusterProfileProjection,omitempty"`

	Config *ClusterConfigReference `json:"config,omitempty"`
	/*
	  The name of the default container created as part of cluster creation. This is part of payload for cluster create operation only.
	*/
	ContainerName *string `json:"containerName,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Number of inefficient VMs.
	*/
	InefficientVmCount *int64 `json:"inefficientVmCount,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Cluster name. This is part of payload for both cluster create &amp; update operations.
	*/
	Name *string `json:"name,omitempty"`

	Network *ClusterNetworkReference `json:"network,omitempty"`

	Nodes *NodeReference `json:"nodes,omitempty"`

	StorageConfigsProjection *StorageConfigsProjection `json:"storageConfigsProjection,omitempty"`

	StorageSummaryProjection *StorageSummaryProjection `json:"storageSummaryProjection,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	UpgradeStatus *UpgradeStatus `json:"upgradeStatus,omitempty"`
	/*
	  Number of VMs.
	*/
	VmCount *int64 `json:"vmCount,omitempty"`
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
	if known.BackupEligibilityScore != nil {
		p.BackupEligibilityScore = known.BackupEligibilityScore
	}
	if known.Categories != nil {
		p.Categories = known.Categories
	}
	if known.ClusterProfileExtId != nil {
		p.ClusterProfileExtId = known.ClusterProfileExtId
	}
	if known.ClusterProfileProjection != nil {
		p.ClusterProfileProjection = known.ClusterProfileProjection
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.ContainerName != nil {
		p.ContainerName = known.ContainerName
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.InefficientVmCount != nil {
		p.InefficientVmCount = known.InefficientVmCount
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.Nodes != nil {
		p.Nodes = known.Nodes
	}
	if known.StorageConfigsProjection != nil {
		p.StorageConfigsProjection = known.StorageConfigsProjection
	}
	if known.StorageSummaryProjection != nil {
		p.StorageSummaryProjection = known.StorageSummaryProjection
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpgradeStatus != nil {
		p.UpgradeStatus = known.UpgradeStatus
	}
	if known.VmCount != nil {
		p.VmCount = known.VmCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backupEligibilityScore")
	delete(allFields, "categories")
	delete(allFields, "clusterProfileExtId")
	delete(allFields, "clusterProfileProjection")
	delete(allFields, "config")
	delete(allFields, "containerName")
	delete(allFields, "extId")
	delete(allFields, "inefficientVmCount")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "network")
	delete(allFields, "nodes")
	delete(allFields, "storageConfigsProjection")
	delete(allFields, "storageSummaryProjection")
	delete(allFields, "tenantId")
	delete(allFields, "upgradeStatus")
	delete(allFields, "vmCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterProjection() *ClusterProjection {
	p := new(ClusterProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the cluster
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN         ClusterType = 0
	CLUSTERTYPE_REDACTED        ClusterType = 1
	CLUSTERTYPE_HYPER_CONVERGED ClusterType = 2
	CLUSTERTYPE_COMPUTE         ClusterType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYPER_CONVERGED",
		"COMPUTE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYPER_CONVERGED",
		"COMPUTE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterType) index(name string) ClusterType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HYPER_CONVERGED",
		"COMPUTE",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterType(idx)
		}
	}
	return CLUSTERTYPE_UNKNOWN
}

func (e *ClusterType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterType) Ref() *ClusterType {
	return &e
}

type ConfigType int

const (
	CONFIGTYPE_UNKNOWN                                     ConfigType = 0
	CONFIGTYPE_REDACTED                                    ConfigType = 1
	CONFIGTYPE_NTP_SERVER_CONFIG                           ConfigType = 2
	CONFIGTYPE_NAME_SERVER_CONFIG                          ConfigType = 3
	CONFIGTYPE_SMTP_SERVER_CONFIG                          ConfigType = 4
	CONFIGTYPE_NFS_SUBNET_WHITELIST_CONFIG                 ConfigType = 5
	CONFIGTYPE_SNMP_SERVER_CONFIG                          ConfigType = 6
	CONFIGTYPE_RSYSLOG_SERVER_CONFIG                       ConfigType = 7
	CONFIGTYPE_PULSE_CONFIG                                ConfigType = 8
	CONFIGTYPE_HTTP_PROXY_CONFIG                           ConfigType = 9
	CONFIGTYPE_FAULT_TOLERANCE_CONFIG                      ConfigType = 10
	CONFIGTYPE_REBUILD_RESERVATION_CONFIG                  ConfigType = 11
	CONFIGTYPE_RESILIENT_CAPACITY_WARNING_THRESHOLD_CONFIG ConfigType = 12
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ConfigType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTP_SERVER_CONFIG",
		"NAME_SERVER_CONFIG",
		"SMTP_SERVER_CONFIG",
		"NFS_SUBNET_WHITELIST_CONFIG",
		"SNMP_SERVER_CONFIG",
		"RSYSLOG_SERVER_CONFIG",
		"PULSE_CONFIG",
		"HTTP_PROXY_CONFIG",
		"FAULT_TOLERANCE_CONFIG",
		"REBUILD_RESERVATION_CONFIG",
		"RESILIENT_CAPACITY_WARNING_THRESHOLD_CONFIG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ConfigType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTP_SERVER_CONFIG",
		"NAME_SERVER_CONFIG",
		"SMTP_SERVER_CONFIG",
		"NFS_SUBNET_WHITELIST_CONFIG",
		"SNMP_SERVER_CONFIG",
		"RSYSLOG_SERVER_CONFIG",
		"PULSE_CONFIG",
		"HTTP_PROXY_CONFIG",
		"FAULT_TOLERANCE_CONFIG",
		"REBUILD_RESERVATION_CONFIG",
		"RESILIENT_CAPACITY_WARNING_THRESHOLD_CONFIG",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ConfigType) index(name string) ConfigType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTP_SERVER_CONFIG",
		"NAME_SERVER_CONFIG",
		"SMTP_SERVER_CONFIG",
		"NFS_SUBNET_WHITELIST_CONFIG",
		"SNMP_SERVER_CONFIG",
		"RSYSLOG_SERVER_CONFIG",
		"PULSE_CONFIG",
		"HTTP_PROXY_CONFIG",
		"FAULT_TOLERANCE_CONFIG",
		"REBUILD_RESERVATION_CONFIG",
		"RESILIENT_CAPACITY_WARNING_THRESHOLD_CONFIG",
	}
	for idx := range names {
		if names[idx] == name {
			return ConfigType(idx)
		}
	}
	return CONFIGTYPE_UNKNOWN
}

func (e *ConfigType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ConfigType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ConfigType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ConfigType) Ref() *ConfigType {
	return &e
}

/*
Domain awareness level corresponds to unit of cluster group. This is part of payload for both cluster create &amp; update operations.
*/
type DomainAwarenessLevel int

const (
	DOMAINAWARENESSLEVEL_UNKNOWN  DomainAwarenessLevel = 0
	DOMAINAWARENESSLEVEL_REDACTED DomainAwarenessLevel = 1
	DOMAINAWARENESSLEVEL_NODE     DomainAwarenessLevel = 2
	DOMAINAWARENESSLEVEL_BLOCK    DomainAwarenessLevel = 3
	DOMAINAWARENESSLEVEL_RACK     DomainAwarenessLevel = 4
	DOMAINAWARENESSLEVEL_DISK     DomainAwarenessLevel = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *DomainAwarenessLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"BLOCK",
		"RACK",
		"DISK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e DomainAwarenessLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"BLOCK",
		"RACK",
		"DISK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *DomainAwarenessLevel) index(name string) DomainAwarenessLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NODE",
		"BLOCK",
		"RACK",
		"DISK",
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

/*
Encryption algorithm used for NTP server authentication.
*/
type EncryptionAlgorithm int

const (
	ENCRYPTIONALGORITHM_UNKNOWN  EncryptionAlgorithm = 0
	ENCRYPTIONALGORITHM_REDACTED EncryptionAlgorithm = 1
	ENCRYPTIONALGORITHM_SHA256   EncryptionAlgorithm = 2
	ENCRYPTIONALGORITHM_SHA384   EncryptionAlgorithm = 3
	ENCRYPTIONALGORITHM_SHA512   EncryptionAlgorithm = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EncryptionAlgorithm) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EncryptionAlgorithm) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EncryptionAlgorithm) index(name string) EncryptionAlgorithm {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	for idx := range names {
		if names[idx] == name {
			return EncryptionAlgorithm(idx)
		}
	}
	return ENCRYPTIONALGORITHM_UNKNOWN
}

func (e *EncryptionAlgorithm) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionAlgorithm:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EncryptionAlgorithm) Ref() *EncryptionAlgorithm {
	return &e
}

/*
Encryption option.
*/
type EncryptionOptionInfo int

const (
	ENCRYPTIONOPTIONINFO_UNKNOWN               EncryptionOptionInfo = 0
	ENCRYPTIONOPTIONINFO_REDACTED              EncryptionOptionInfo = 1
	ENCRYPTIONOPTIONINFO_SOFTWARE              EncryptionOptionInfo = 2
	ENCRYPTIONOPTIONINFO_HARDWARE              EncryptionOptionInfo = 3
	ENCRYPTIONOPTIONINFO_SOFTWARE_AND_HARDWARE EncryptionOptionInfo = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EncryptionOptionInfo) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"HARDWARE",
		"SOFTWARE_AND_HARDWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EncryptionOptionInfo) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"HARDWARE",
		"SOFTWARE_AND_HARDWARE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EncryptionOptionInfo) index(name string) EncryptionOptionInfo {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SOFTWARE",
		"HARDWARE",
		"SOFTWARE_AND_HARDWARE",
	}
	for idx := range names {
		if names[idx] == name {
			return EncryptionOptionInfo(idx)
		}
	}
	return ENCRYPTIONOPTIONINFO_UNKNOWN
}

func (e *EncryptionOptionInfo) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionOptionInfo:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EncryptionOptionInfo) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EncryptionOptionInfo) Ref() *EncryptionOptionInfo {
	return &e
}

/*
Encryption scope.
*/
type EncryptionScopeInfo int

const (
	ENCRYPTIONSCOPEINFO_UNKNOWN   EncryptionScopeInfo = 0
	ENCRYPTIONSCOPEINFO_REDACTED  EncryptionScopeInfo = 1
	ENCRYPTIONSCOPEINFO_CLUSTER   EncryptionScopeInfo = 2
	ENCRYPTIONSCOPEINFO_CONTAINER EncryptionScopeInfo = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EncryptionScopeInfo) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"CONTAINER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EncryptionScopeInfo) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"CONTAINER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EncryptionScopeInfo) index(name string) EncryptionScopeInfo {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CLUSTER",
		"CONTAINER",
	}
	for idx := range names {
		if names[idx] == name {
			return EncryptionScopeInfo(idx)
		}
	}
	return ENCRYPTIONSCOPEINFO_UNKNOWN
}

func (e *EncryptionScopeInfo) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionScopeInfo:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EncryptionScopeInfo) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EncryptionScopeInfo) Ref() *EncryptionScopeInfo {
	return &e
}

/*
Encryption in transit Status.
*/
type EncryptionStatus int

const (
	ENCRYPTIONSTATUS_UNKNOWN  EncryptionStatus = 0
	ENCRYPTIONSTATUS_REDACTED EncryptionStatus = 1
	ENCRYPTIONSTATUS_ENABLED  EncryptionStatus = 2
	ENCRYPTIONSTATUS_DISABLED EncryptionStatus = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EncryptionStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EncryptionStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EncryptionStatus) index(name string) EncryptionStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return EncryptionStatus(idx)
		}
	}
	return ENCRYPTIONSTATUS_UNKNOWN
}

func (e *EncryptionStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EncryptionStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EncryptionStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EncryptionStatus) Ref() *EncryptionStatus {
	return &e
}

/*
The type of external storage provider.
*/
type ExternalStorageProvider int

const (
	EXTERNALSTORAGEPROVIDER_UNKNOWN             ExternalStorageProvider = 0
	EXTERNALSTORAGEPROVIDER_REDACTED            ExternalStorageProvider = 1
	EXTERNALSTORAGEPROVIDER_DELL_POWERFLEX      ExternalStorageProvider = 2
	EXTERNALSTORAGEPROVIDER_EVERPURE_FLASHARRAY ExternalStorageProvider = 3
	EXTERNALSTORAGEPROVIDER_DELL_POWERSTORE     ExternalStorageProvider = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ExternalStorageProvider) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELL_POWERFLEX",
		"EVERPURE_FLASHARRAY",
		"DELL_POWERSTORE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ExternalStorageProvider) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELL_POWERFLEX",
		"EVERPURE_FLASHARRAY",
		"DELL_POWERSTORE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ExternalStorageProvider) index(name string) ExternalStorageProvider {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DELL_POWERFLEX",
		"EVERPURE_FLASHARRAY",
		"DELL_POWERSTORE",
	}
	for idx := range names {
		if names[idx] == name {
			return ExternalStorageProvider(idx)
		}
	}
	return EXTERNALSTORAGEPROVIDER_UNKNOWN
}

func (e *ExternalStorageProvider) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ExternalStorageProvider:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ExternalStorageProvider) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ExternalStorageProvider) Ref() *ExternalStorageProvider {
	return &e
}

/*
List of external storage provider including type, installation status, and compatible version.
*/
type ExternalStorageProviderInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Compatible version of the external storage provider.
	*/
	CompatibleVersion *string `json:"compatibleVersion,omitempty"`
	/*
	  External storage provider is installed on the cluster or not.
	*/
	IsInstalled *bool `json:"isInstalled,omitempty"`

	ProviderType *ExternalStorageProvider `json:"providerType,omitempty"`
}

func (p *ExternalStorageProviderInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ExternalStorageProviderInfo

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

func (p *ExternalStorageProviderInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalStorageProviderInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalStorageProviderInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CompatibleVersion != nil {
		p.CompatibleVersion = known.CompatibleVersion
	}
	if known.IsInstalled != nil {
		p.IsInstalled = known.IsInstalled
	}
	if known.ProviderType != nil {
		p.ProviderType = known.ProviderType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "compatibleVersion")
	delete(allFields, "isInstalled")
	delete(allFields, "providerType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalStorageProviderInfo() *ExternalStorageProviderInfo {
	p := new(ExternalStorageProviderInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ExternalStorageProviderInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Fault tolerant state of cluster.
*/
type FaultToleranceConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DesiredClusterFaultTolerance *ClusterFaultToleranceRef `json:"desiredClusterFaultTolerance,omitempty"`
}

func (p *FaultToleranceConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FaultToleranceConfiguration

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

func (p *FaultToleranceConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FaultToleranceConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFaultToleranceConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DesiredClusterFaultTolerance != nil {
		p.DesiredClusterFaultTolerance = known.DesiredClusterFaultTolerance
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "desiredClusterFaultTolerance")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFaultToleranceConfiguration() *FaultToleranceConfiguration {
	p := new(FaultToleranceConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.FaultToleranceConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Fault tolerant state of cluster.
*/
type FaultToleranceState struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CurrentClusterFaultTolerance *ClusterFaultToleranceRef `json:"currentClusterFaultTolerance,omitempty"`
	/*
	  Maximum fault tolerance that is supported currently.
	*/
	CurrentMaxFaultTolerance *int `json:"currentMaxFaultTolerance,omitempty"`

	DesiredClusterFaultTolerance *ClusterFaultToleranceRef `json:"desiredClusterFaultTolerance,omitempty"`
	/*
	  Maximum fault tolerance desired.
	*/
	DesiredMaxFaultTolerance *int `json:"desiredMaxFaultTolerance,omitempty"`

	DomainAwarenessLevel *DomainAwarenessLevel `json:"domainAwarenessLevel,omitempty"`
	/*
	  Flag to determine if strict rack awareness is enabled for the cluster.
	*/
	IsStrictRackAwarenessEnabled *bool `json:"isStrictRackAwarenessEnabled,omitempty"`

	RedundancyStatus *RedundancyStatusDetails `json:"redundancyStatus,omitempty"`
}

func (p *FaultToleranceState) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FaultToleranceState

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

func (p *FaultToleranceState) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FaultToleranceState
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFaultToleranceState()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CurrentClusterFaultTolerance != nil {
		p.CurrentClusterFaultTolerance = known.CurrentClusterFaultTolerance
	}
	if known.CurrentMaxFaultTolerance != nil {
		p.CurrentMaxFaultTolerance = known.CurrentMaxFaultTolerance
	}
	if known.DesiredClusterFaultTolerance != nil {
		p.DesiredClusterFaultTolerance = known.DesiredClusterFaultTolerance
	}
	if known.DesiredMaxFaultTolerance != nil {
		p.DesiredMaxFaultTolerance = known.DesiredMaxFaultTolerance
	}
	if known.DomainAwarenessLevel != nil {
		p.DomainAwarenessLevel = known.DomainAwarenessLevel
	}
	if known.IsStrictRackAwarenessEnabled != nil {
		p.IsStrictRackAwarenessEnabled = known.IsStrictRackAwarenessEnabled
	}
	if known.RedundancyStatus != nil {
		p.RedundancyStatus = known.RedundancyStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "currentClusterFaultTolerance")
	delete(allFields, "currentMaxFaultTolerance")
	delete(allFields, "desiredClusterFaultTolerance")
	delete(allFields, "desiredMaxFaultTolerance")
	delete(allFields, "domainAwarenessLevel")
	delete(allFields, "isStrictRackAwarenessEnabled")
	delete(allFields, "redundancyStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFaultToleranceState() *FaultToleranceState {
	p := new(FaultToleranceState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.FaultToleranceState"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
HTTP Proxy server configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
*/
type HttpProxyConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPAddress `json:"ipAddress,omitempty"`
	/*
	  HTTP Proxy server name configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	Name *string `json:"name"`
	/*
	  HTTP Proxy server password needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  HTTP Proxy server port configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	Port *int `json:"port,omitempty"`
	/*
	  List of HTTP proxy types.
	*/
	ProxyTypes []HttpProxyType `json:"proxyTypes,omitempty"`
	/*
	  HTTP Proxy server username needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *HttpProxyConfig) MarshalJSON() ([]byte, error) {
	type HttpProxyConfigProxy HttpProxyConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HttpProxyConfigProxy
		Name *string `json:"name,omitempty"`
	}{
		HttpProxyConfigProxy: (*HttpProxyConfigProxy)(p),
		Name:                 p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *HttpProxyConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HttpProxyConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHttpProxyConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.ProxyTypes != nil {
		p.ProxyTypes = known.ProxyTypes
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "name")
	delete(allFields, "password")
	delete(allFields, "port")
	delete(allFields, "proxyTypes")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHttpProxyConfig() *HttpProxyConfig {
	p := new(HttpProxyConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HttpProxyConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
HTTP proxy configuration for the cluster.
*/
type HttpProxyConfiguration struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of HTTP Proxy server configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	ProxyList []HttpProxyConfig `json:"proxyList,omitempty"`
	/*
	  Targets HTTP traffic to which is exempted from going through the configured HTTP Proxy.
	*/
	ProxyWhiteList []HttpProxyWhiteListConfig `json:"proxyWhiteList,omitempty"`
}

func (p *HttpProxyConfiguration) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HttpProxyConfiguration

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

func (p *HttpProxyConfiguration) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HttpProxyConfiguration
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHttpProxyConfiguration()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProxyList != nil {
		p.ProxyList = known.ProxyList
	}
	if known.ProxyWhiteList != nil {
		p.ProxyWhiteList = known.ProxyWhiteList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "proxyList")
	delete(allFields, "proxyWhiteList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHttpProxyConfiguration() *HttpProxyConfiguration {
	p := new(HttpProxyConfiguration)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HttpProxyConfiguration"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
HTTP proxy type which is needed to access a cluster hosted behind a HTTP Proxy.
*/
type HttpProxyType int

const (
	HTTPPROXYTYPE_UNKNOWN  HttpProxyType = 0
	HTTPPROXYTYPE_REDACTED HttpProxyType = 1
	HTTPPROXYTYPE_HTTP     HttpProxyType = 2
	HTTPPROXYTYPE_HTTPS    HttpProxyType = 3
	HTTPPROXYTYPE_SOCKS    HttpProxyType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HttpProxyType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HttpProxyType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HttpProxyType) index(name string) HttpProxyType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HTTP",
		"HTTPS",
		"SOCKS",
	}
	for idx := range names {
		if names[idx] == name {
			return HttpProxyType(idx)
		}
	}
	return HTTPPROXYTYPE_UNKNOWN
}

func (e *HttpProxyType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HttpProxyType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HttpProxyType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HttpProxyType) Ref() *HttpProxyType {
	return &e
}

/*
Targets HTTP traffic to which is exempted from going through the configured HTTP Proxy.
*/
type HttpProxyWhiteListConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Target's identifier which is exempted from going through the configured HTTP Proxy.
	*/
	Target *string `json:"target"`

	TargetType *HttpProxyWhiteListTargetType `json:"targetType"`
}

func (p *HttpProxyWhiteListConfig) MarshalJSON() ([]byte, error) {
	type HttpProxyWhiteListConfigProxy HttpProxyWhiteListConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HttpProxyWhiteListConfigProxy
		Target     *string                       `json:"target,omitempty"`
		TargetType *HttpProxyWhiteListTargetType `json:"targetType,omitempty"`
	}{
		HttpProxyWhiteListConfigProxy: (*HttpProxyWhiteListConfigProxy)(p),
		Target:                        p.Target,
		TargetType:                    p.TargetType,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *HttpProxyWhiteListConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HttpProxyWhiteListConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHttpProxyWhiteListConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Target != nil {
		p.Target = known.Target
	}
	if known.TargetType != nil {
		p.TargetType = known.TargetType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "target")
	delete(allFields, "targetType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHttpProxyWhiteListConfig() *HttpProxyWhiteListConfig {
	p := new(HttpProxyWhiteListConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HttpProxyWhiteListConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the target which is exempted from going through the configured HTTP Proxy.
*/
type HttpProxyWhiteListTargetType int

const (
	HTTPPROXYWHITELISTTARGETTYPE_UNKNOWN            HttpProxyWhiteListTargetType = 0
	HTTPPROXYWHITELISTTARGETTYPE_REDACTED           HttpProxyWhiteListTargetType = 1
	HTTPPROXYWHITELISTTARGETTYPE_IPV4_ADDRESS       HttpProxyWhiteListTargetType = 2
	HTTPPROXYWHITELISTTARGETTYPE_IPV6_ADDRESS       HttpProxyWhiteListTargetType = 3
	HTTPPROXYWHITELISTTARGETTYPE_IPV4_NETWORK_MASK  HttpProxyWhiteListTargetType = 4
	HTTPPROXYWHITELISTTARGETTYPE_DOMAIN_NAME_SUFFIX HttpProxyWhiteListTargetType = 5
	HTTPPROXYWHITELISTTARGETTYPE_HOST_NAME          HttpProxyWhiteListTargetType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HttpProxyWhiteListTargetType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4_ADDRESS",
		"IPV6_ADDRESS",
		"IPV4_NETWORK_MASK",
		"DOMAIN_NAME_SUFFIX",
		"HOST_NAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HttpProxyWhiteListTargetType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4_ADDRESS",
		"IPV6_ADDRESS",
		"IPV4_NETWORK_MASK",
		"DOMAIN_NAME_SUFFIX",
		"HOST_NAME",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HttpProxyWhiteListTargetType) index(name string) HttpProxyWhiteListTargetType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4_ADDRESS",
		"IPV6_ADDRESS",
		"IPV4_NETWORK_MASK",
		"DOMAIN_NAME_SUFFIX",
		"HOST_NAME",
	}
	for idx := range names {
		if names[idx] == name {
			return HttpProxyWhiteListTargetType(idx)
		}
	}
	return HTTPPROXYWHITELISTTARGETTYPE_UNKNOWN
}

func (e *HttpProxyWhiteListTargetType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for HttpProxyWhiteListTargetType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *HttpProxyWhiteListTargetType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e HttpProxyWhiteListTargetType) Ref() *HttpProxyWhiteListTargetType {
	return &e
}

/*
Hypervisor type.
*/
type HypervisorType int

const (
	HYPERVISORTYPE_UNKNOWN    HypervisorType = 0
	HYPERVISORTYPE_REDACTED   HypervisorType = 1
	HYPERVISORTYPE_AHV        HypervisorType = 2
	HYPERVISORTYPE_ESX        HypervisorType = 3
	HYPERVISORTYPE_HYPERV     HypervisorType = 4
	HYPERVISORTYPE_XEN        HypervisorType = 5
	HYPERVISORTYPE_NATIVEHOST HypervisorType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *HypervisorType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
		"HYPERV",
		"XEN",
		"NATIVEHOST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e HypervisorType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
		"HYPERV",
		"XEN",
		"NATIVEHOST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *HypervisorType) index(name string) HypervisorType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AHV",
		"ESX",
		"HYPERV",
		"XEN",
		"NATIVEHOST",
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

/*
Management server type.
*/
type KeyManagementServerType int

const (
	KEYMANAGEMENTSERVERTYPE_UNKNOWN       KeyManagementServerType = 0
	KEYMANAGEMENTSERVERTYPE_REDACTED      KeyManagementServerType = 1
	KEYMANAGEMENTSERVERTYPE_LOCAL         KeyManagementServerType = 2
	KEYMANAGEMENTSERVERTYPE_PRISM_CENTRAL KeyManagementServerType = 3
	KEYMANAGEMENTSERVERTYPE_EXTERNAL      KeyManagementServerType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *KeyManagementServerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"PRISM_CENTRAL",
		"EXTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e KeyManagementServerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"PRISM_CENTRAL",
		"EXTERNAL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *KeyManagementServerType) index(name string) KeyManagementServerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"LOCAL",
		"PRISM_CENTRAL",
		"EXTERNAL",
	}
	for idx := range names {
		if names[idx] == name {
			return KeyManagementServerType(idx)
		}
	}
	return KEYMANAGEMENTSERVERTYPE_UNKNOWN
}

func (e *KeyManagementServerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for KeyManagementServerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *KeyManagementServerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e KeyManagementServerType) Ref() *KeyManagementServerType {
	return &e
}

/*
Managed cluster information.
*/
type ManagedCluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Drifted settings information.
	*/
	ConfigDrifts []ConfigType `json:"configDrifts,omitempty"`
	/*
	  Indicates the UUID of a cluster.
	*/
	ExtId *string `json:"extId"`
	/*
	  Indicates if the attached cluster is compliant with the cluster profile or not.
	*/
	IsCompliant *bool `json:"isCompliant,omitempty"`
	/*
	  The most recent date and time when the cluster profile was monitored across all attached clusters.
	*/
	LastSyncedTime *time.Time `json:"lastSyncedTime,omitempty"`
}

func (p *ManagedCluster) MarshalJSON() ([]byte, error) {
	type ManagedClusterProxy ManagedCluster

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ManagedClusterProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		ManagedClusterProxy: (*ManagedClusterProxy)(p),
		ExtId:               p.ExtId,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *ManagedCluster) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ManagedCluster
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewManagedCluster()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConfigDrifts != nil {
		p.ConfigDrifts = known.ConfigDrifts
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsCompliant != nil {
		p.IsCompliant = known.IsCompliant
	}
	if known.LastSyncedTime != nil {
		p.LastSyncedTime = known.LastSyncedTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "configDrifts")
	delete(allFields, "extId")
	delete(allFields, "isCompliant")
	delete(allFields, "lastSyncedTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewManagedCluster() *ManagedCluster {
	p := new(ManagedCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ManagedCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Management server information.
*/
type ManagementServerRef struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ip *import1.IPAddress `json:"ip,omitempty"`
	/*
	  Indicates whether it is DRS enabled or not.
	*/
	IsDrsEnabled *bool `json:"isDrsEnabled,omitempty"`
	/*
	  Indicates whether the host is managed by an entity or not.
	*/
	IsInUse *bool `json:"isInUse,omitempty"`
	/*
	  Indicates whether it is registered or not.
	*/
	IsRegistered *bool `json:"isRegistered,omitempty"`

	Type *ManagementServerType `json:"type,omitempty"`
}

func (p *ManagementServerRef) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ManagementServerRef

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

func (p *ManagementServerRef) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ManagementServerRef
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewManagementServerRef()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Ip != nil {
		p.Ip = known.Ip
	}
	if known.IsDrsEnabled != nil {
		p.IsDrsEnabled = known.IsDrsEnabled
	}
	if known.IsInUse != nil {
		p.IsInUse = known.IsInUse
	}
	if known.IsRegistered != nil {
		p.IsRegistered = known.IsRegistered
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ip")
	delete(allFields, "isDrsEnabled")
	delete(allFields, "isInUse")
	delete(allFields, "isRegistered")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewManagementServerRef() *ManagementServerRef {
	p := new(ManagementServerRef)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ManagementServerRef"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Management server type.
*/
type ManagementServerType int

const (
	MANAGEMENTSERVERTYPE_UNKNOWN  ManagementServerType = 0
	MANAGEMENTSERVERTYPE_REDACTED ManagementServerType = 1
	MANAGEMENTSERVERTYPE_VCENTER  ManagementServerType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ManagementServerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VCENTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ManagementServerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VCENTER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ManagementServerType) index(name string) ManagementServerType {
	names := [...]string{
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

/*
List of nodes in a cluster.
*/
type NodeListItemReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ControllerVmIp *import1.IPAddress `json:"controllerVmIp,omitempty"`

	HostIp *import1.IPAddress `json:"hostIp,omitempty"`
	/*
	  UUID of the host.
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
}

func (p *NodeListItemReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NodeListItemReference

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

func (p *NodeListItemReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeListItemReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeListItemReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ControllerVmIp != nil {
		p.ControllerVmIp = known.ControllerVmIp
	}
	if known.HostIp != nil {
		p.HostIp = known.HostIp
	}
	if known.NodeUuid != nil {
		p.NodeUuid = known.NodeUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerVmIp")
	delete(allFields, "hostIp")
	delete(allFields, "nodeUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeListItemReference() *NodeListItemReference {
	p := new(NodeListItemReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeListItemReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Node reference for a cluster.
*/
type NodeReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of nodes in a cluster.
	*/
	NodeList []NodeListItemReference `json:"nodeList,omitempty"`
	/*
	  Number of nodes in a cluster.
	*/
	NumberOfNodes *int `json:"numberOfNodes,omitempty"`
}

func (p *NodeReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NodeReference

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

func (p *NodeReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NodeReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNodeReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NodeList != nil {
		p.NodeList = known.NodeList
	}
	if known.NumberOfNodes != nil {
		p.NumberOfNodes = known.NumberOfNodes
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "nodeList")
	delete(allFields, "numberOfNodes")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNodeReference() *NodeReference {
	p := new(NodeReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
NTP server configuration details.
*/
type NtpServerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EncryptionAlgorithm *EncryptionAlgorithm `json:"encryptionAlgorithm,omitempty"`
	/*
	  Encryption key in hexadecimal format used for NTP server authentication.
	*/
	EncryptionKey *string `json:"encryptionKey,omitempty"`
	/*
	  Encryption key Id used for NTP server authentication.
	*/
	EncryptionKeyId *int `json:"encryptionKeyId,omitempty"`

	NtpServerAddress *import1.IPAddressOrFQDN `json:"ntpServerAddress"`
}

func (p *NtpServerConfig) MarshalJSON() ([]byte, error) {
	type NtpServerConfigProxy NtpServerConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NtpServerConfigProxy
		NtpServerAddress *import1.IPAddressOrFQDN `json:"ntpServerAddress,omitempty"`
	}{
		NtpServerConfigProxy: (*NtpServerConfigProxy)(p),
		NtpServerAddress:     p.NtpServerAddress,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *NtpServerConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NtpServerConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNtpServerConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EncryptionAlgorithm != nil {
		p.EncryptionAlgorithm = known.EncryptionAlgorithm
	}
	if known.EncryptionKey != nil {
		p.EncryptionKey = known.EncryptionKey
	}
	if known.EncryptionKeyId != nil {
		p.EncryptionKeyId = known.EncryptionKeyId
	}
	if known.NtpServerAddress != nil {
		p.NtpServerAddress = known.NtpServerAddress
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "encryptionAlgorithm")
	delete(allFields, "encryptionKey")
	delete(allFields, "encryptionKeyId")
	delete(allFields, "ntpServerAddress")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNtpServerConfig() *NtpServerConfig {
	p := new(NtpServerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NtpServerConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster operation mode. This is part of payload for cluster update operation only.
*/
type OperationMode int

const (
	OPERATIONMODE_UNKNOWN            OperationMode = 0
	OPERATIONMODE_REDACTED           OperationMode = 1
	OPERATIONMODE_NORMAL             OperationMode = 2
	OPERATIONMODE_READ_ONLY          OperationMode = 3
	OPERATIONMODE_STAND_ALONE        OperationMode = 4
	OPERATIONMODE_SWITCH_TO_TWO_NODE OperationMode = 5
	OPERATIONMODE_OVERRIDE           OperationMode = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OperationMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"READ_ONLY",
		"STAND_ALONE",
		"SWITCH_TO_TWO_NODE",
		"OVERRIDE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OperationMode) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"READ_ONLY",
		"STAND_ALONE",
		"SWITCH_TO_TWO_NODE",
		"OVERRIDE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OperationMode) index(name string) OperationMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"READ_ONLY",
		"STAND_ALONE",
		"SWITCH_TO_TWO_NODE",
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

/*
PII Scrubbing Level for pulse.
*/
type PIIScrubbingLevel int

const (
	PIISCRUBBINGLEVEL_UNKNOWN  PIIScrubbingLevel = 0
	PIISCRUBBINGLEVEL_REDACTED PIIScrubbingLevel = 1
	PIISCRUBBINGLEVEL_DEFAULT  PIIScrubbingLevel = 2
	PIISCRUBBINGLEVEL_ALL      PIIScrubbingLevel = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PIIScrubbingLevel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PIIScrubbingLevel) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"ALL",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PIIScrubbingLevel) index(name string) PIIScrubbingLevel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"ALL",
	}
	for idx := range names {
		if names[idx] == name {
			return PIIScrubbingLevel(idx)
		}
	}
	return PIISCRUBBINGLEVEL_UNKNOWN
}

func (e *PIIScrubbingLevel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PIIScrubbingLevel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PIIScrubbingLevel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PIIScrubbingLevel) Ref() *PIIScrubbingLevel {
	return &e
}

/*
Public ssh key details. This is part of payload for cluster update operation only.
*/
type PublicKey struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Ssh key value.
	*/
	Key *string `json:"key"`
	/*
	  Ssh key name.
	*/
	Name *string `json:"name"`
}

func (p *PublicKey) MarshalJSON() ([]byte, error) {
	type PublicKeyProxy PublicKey

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PublicKeyProxy
		Key  *string `json:"key,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		PublicKeyProxy: (*PublicKeyProxy)(p),
		Key:            p.Key,
		Name:           p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *PublicKey) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PublicKey
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPublicKey()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Key != nil {
		p.Key = known.Key
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "key")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPublicKey() *PublicKey {
	p := new(PublicKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.PublicKey"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Pulse status for a cluster.
*/
type PulseStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Flag determines whether pulse is enabled for a cluster.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`

	PiiScrubbingLevel *PIIScrubbingLevel `json:"piiScrubbingLevel,omitempty"`
}

func (p *PulseStatus) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PulseStatus

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

func (p *PulseStatus) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PulseStatus
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPulseStatus()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.PiiScrubbingLevel != nil {
		p.PiiScrubbingLevel = known.PiiScrubbingLevel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isEnabled")
	delete(allFields, "piiScrubbingLevel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPulseStatus() *PulseStatus {
	p := new(PulseStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.PulseStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Rebuild capacity reservation configuration.
*/
type RebuildReservationConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Enable rebuild capacity reservation to maintain free space for fault tolerance recovery operations.
	*/
	IsRebuildReservationEnabled *bool `json:"isRebuildReservationEnabled,omitempty"`
}

func (p *RebuildReservationConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RebuildReservationConfig

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

func (p *RebuildReservationConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RebuildReservationConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRebuildReservationConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsRebuildReservationEnabled != nil {
		p.IsRebuildReservationEnabled = known.IsRebuildReservationEnabled
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isRebuildReservationEnabled")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRebuildReservationConfig() *RebuildReservationConfig {
	p := new(RebuildReservationConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RebuildReservationConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsRebuildReservationEnabled = new(bool)
	*p.IsRebuildReservationEnabled = false

	return p
}

/*
Redundancy Status of the cluster
*/
type RedundancyStatusDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Boolean flag to indicate if Cassandra ensemble can meet the desired FT.
	*/
	IsCassandraPreparationDone *bool `json:"isCassandraPreparationDone,omitempty"`
	/*
	  Boolean flag to indicate if Zookeeper ensemble can meet the desired FT.
	*/
	IsZookeeperPreparationDone *bool `json:"isZookeeperPreparationDone,omitempty"`
}

func (p *RedundancyStatusDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RedundancyStatusDetails

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

func (p *RedundancyStatusDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RedundancyStatusDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRedundancyStatusDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsCassandraPreparationDone != nil {
		p.IsCassandraPreparationDone = known.IsCassandraPreparationDone
	}
	if known.IsZookeeperPreparationDone != nil {
		p.IsZookeeperPreparationDone = known.IsZookeeperPreparationDone
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isCassandraPreparationDone")
	delete(allFields, "isZookeeperPreparationDone")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRedundancyStatusDetails() *RedundancyStatusDetails {
	p := new(RedundancyStatusDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RedundancyStatusDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Resilient capacity warning threshold configuration.
*/
type ResilientCapacityWarningThresholdConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Warning threshold percentage for resilient storage capacity.
	*/
	ResilientCapacityWarningThresholdPercentage *int `json:"resilientCapacityWarningThresholdPercentage,omitempty"`
}

func (p *ResilientCapacityWarningThresholdConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ResilientCapacityWarningThresholdConfig

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

func (p *ResilientCapacityWarningThresholdConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResilientCapacityWarningThresholdConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResilientCapacityWarningThresholdConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ResilientCapacityWarningThresholdPercentage != nil {
		p.ResilientCapacityWarningThresholdPercentage = known.ResilientCapacityWarningThresholdPercentage
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "resilientCapacityWarningThresholdPercentage")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewResilientCapacityWarningThresholdConfig() *ResilientCapacityWarningThresholdConfig {
	p := new(ResilientCapacityWarningThresholdConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ResilientCapacityWarningThresholdConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ResilientCapacityWarningThresholdPercentage = new(int)
	*p.ResilientCapacityWarningThresholdPercentage = 75

	return p
}

/*
RSYSLOG Module information.
*/
type RsyslogModuleItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LogSeverityLevel *RsyslogModuleLogSeverityLevel `json:"logSeverityLevel"`

	Name *RsyslogModuleName `json:"name"`
	/*
	  Option to log, monitor/output files of a module.
	*/
	ShouldLogMonitorFiles *bool `json:"shouldLogMonitorFiles,omitempty"`
}

func (p *RsyslogModuleItem) MarshalJSON() ([]byte, error) {
	type RsyslogModuleItemProxy RsyslogModuleItem

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RsyslogModuleItemProxy
		LogSeverityLevel *RsyslogModuleLogSeverityLevel `json:"logSeverityLevel,omitempty"`
		Name             *RsyslogModuleName             `json:"name,omitempty"`
	}{
		RsyslogModuleItemProxy: (*RsyslogModuleItemProxy)(p),
		LogSeverityLevel:       p.LogSeverityLevel,
		Name:                   p.Name,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *RsyslogModuleItem) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RsyslogModuleItem
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRsyslogModuleItem()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LogSeverityLevel != nil {
		p.LogSeverityLevel = known.LogSeverityLevel
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ShouldLogMonitorFiles != nil {
		p.ShouldLogMonitorFiles = known.ShouldLogMonitorFiles
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "logSeverityLevel")
	delete(allFields, "name")
	delete(allFields, "shouldLogMonitorFiles")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRsyslogModuleItem() *RsyslogModuleItem {
	p := new(RsyslogModuleItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RsyslogModuleItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldLogMonitorFiles = new(bool)
	*p.ShouldLogMonitorFiles = true

	return p
}

/*
RSYSLOG module log severity level.
*/
type RsyslogModuleLogSeverityLevel int

const (
	RSYSLOGMODULELOGSEVERITYLEVEL_UNKNOWN   RsyslogModuleLogSeverityLevel = 0
	RSYSLOGMODULELOGSEVERITYLEVEL_REDACTED  RsyslogModuleLogSeverityLevel = 1
	RSYSLOGMODULELOGSEVERITYLEVEL_EMERGENCY RsyslogModuleLogSeverityLevel = 2
	RSYSLOGMODULELOGSEVERITYLEVEL_ALERT     RsyslogModuleLogSeverityLevel = 3
	RSYSLOGMODULELOGSEVERITYLEVEL_CRITICAL  RsyslogModuleLogSeverityLevel = 4
	RSYSLOGMODULELOGSEVERITYLEVEL_ERROR     RsyslogModuleLogSeverityLevel = 5
	RSYSLOGMODULELOGSEVERITYLEVEL_WARNING   RsyslogModuleLogSeverityLevel = 6
	RSYSLOGMODULELOGSEVERITYLEVEL_NOTICE    RsyslogModuleLogSeverityLevel = 7
	RSYSLOGMODULELOGSEVERITYLEVEL_INFO      RsyslogModuleLogSeverityLevel = 8
	RSYSLOGMODULELOGSEVERITYLEVEL_DEBUG     RsyslogModuleLogSeverityLevel = 9
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RsyslogModuleLogSeverityLevel) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e RsyslogModuleLogSeverityLevel) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *RsyslogModuleLogSeverityLevel) index(name string) RsyslogModuleLogSeverityLevel {
	names := [...]string{
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

/*
RSYSLOG module name.
*/
type RsyslogModuleName int

const (
	RSYSLOGMODULENAME_UNKNOWN           RsyslogModuleName = 0
	RSYSLOGMODULENAME_REDACTED          RsyslogModuleName = 1
	RSYSLOGMODULENAME_CASSANDRA         RsyslogModuleName = 2
	RSYSLOGMODULENAME_CEREBRO           RsyslogModuleName = 3
	RSYSLOGMODULENAME_CURATOR           RsyslogModuleName = 4
	RSYSLOGMODULENAME_GENESIS           RsyslogModuleName = 5
	RSYSLOGMODULENAME_PRISM             RsyslogModuleName = 6
	RSYSLOGMODULENAME_STARGATE          RsyslogModuleName = 7
	RSYSLOGMODULENAME_SYSLOG_MODULE     RsyslogModuleName = 8
	RSYSLOGMODULENAME_ZOOKEEPER         RsyslogModuleName = 9
	RSYSLOGMODULENAME_UHARA             RsyslogModuleName = 10
	RSYSLOGMODULENAME_LAZAN             RsyslogModuleName = 11
	RSYSLOGMODULENAME_API_AUDIT         RsyslogModuleName = 12
	RSYSLOGMODULENAME_AUDIT             RsyslogModuleName = 13
	RSYSLOGMODULENAME_CALM              RsyslogModuleName = 14
	RSYSLOGMODULENAME_EPSILON           RsyslogModuleName = 15
	RSYSLOGMODULENAME_ACROPOLIS         RsyslogModuleName = 16
	RSYSLOGMODULENAME_MINERVA_CVM       RsyslogModuleName = 17
	RSYSLOGMODULENAME_FLOW              RsyslogModuleName = 18
	RSYSLOGMODULENAME_FLOW_SERVICE_LOGS RsyslogModuleName = 19
	RSYSLOGMODULENAME_LCM               RsyslogModuleName = 20
	RSYSLOGMODULENAME_APLOS             RsyslogModuleName = 21
	RSYSLOGMODULENAME_NCM_AIOPS         RsyslogModuleName = 22
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RsyslogModuleName) name(index int) string {
	names := [...]string{
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
		"NCM_AIOPS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RsyslogModuleName) GetName() string {
	index := int(e)
	names := [...]string{
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
		"NCM_AIOPS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RsyslogModuleName) index(name string) RsyslogModuleName {
	names := [...]string{
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
		"NCM_AIOPS",
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

/*
RSYSLOG server protocol type.
*/
type RsyslogNetworkProtocol int

const (
	RSYSLOGNETWORKPROTOCOL_UNKNOWN  RsyslogNetworkProtocol = 0
	RSYSLOGNETWORKPROTOCOL_REDACTED RsyslogNetworkProtocol = 1
	RSYSLOGNETWORKPROTOCOL_UDP      RsyslogNetworkProtocol = 2
	RSYSLOGNETWORKPROTOCOL_TCP      RsyslogNetworkProtocol = 3
	RSYSLOGNETWORKPROTOCOL_RELP     RsyslogNetworkProtocol = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RsyslogNetworkProtocol) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e RsyslogNetworkProtocol) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *RsyslogNetworkProtocol) index(name string) RsyslogNetworkProtocol {
	names := [...]string{
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

/*
RSYSLOG Server.
*/
type RsyslogServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IpAddress *import1.IPAddress `json:"ipAddress"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  List of modules registered to RSYSLOG server.
	*/
	Modules []RsyslogModuleItem `json:"modules,omitempty"`

	NetworkProtocol *RsyslogNetworkProtocol `json:"networkProtocol"`
	/*
	  RSYSLOG server port.
	*/
	Port *int `json:"port"`
	/*
	  RSYSLOG server name.
	*/
	ServerName *string `json:"serverName"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RsyslogServer) MarshalJSON() ([]byte, error) {
	type RsyslogServerProxy RsyslogServer

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RsyslogServerProxy
		IpAddress       *import1.IPAddress      `json:"ipAddress,omitempty"`
		NetworkProtocol *RsyslogNetworkProtocol `json:"networkProtocol,omitempty"`
		Port            *int                    `json:"port,omitempty"`
		ServerName      *string                 `json:"serverName,omitempty"`
	}{
		RsyslogServerProxy: (*RsyslogServerProxy)(p),
		IpAddress:          p.IpAddress,
		NetworkProtocol:    p.NetworkProtocol,
		Port:               p.Port,
		ServerName:         p.ServerName,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *RsyslogServer) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RsyslogServer
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRsyslogServer()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Modules != nil {
		p.Modules = known.Modules
	}
	if known.NetworkProtocol != nil {
		p.NetworkProtocol = known.NetworkProtocol
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.ServerName != nil {
		p.ServerName = known.ServerName
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "ipAddress")
	delete(allFields, "links")
	delete(allFields, "modules")
	delete(allFields, "networkProtocol")
	delete(allFields, "port")
	delete(allFields, "serverName")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRsyslogServer() *RsyslogServer {
	p := new(RsyslogServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RsyslogServer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SMTP network details.
*/
type SmtpNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import1.IPAddressOrFQDN `json:"ipAddress"`
	/*
	  SMTP server password.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  SMTP port.
	*/
	Port *int `json:"port,omitempty"`
	/*
	  SMTP server user name.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *SmtpNetwork) MarshalJSON() ([]byte, error) {
	type SmtpNetworkProxy SmtpNetwork

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SmtpNetworkProxy
		IpAddress *import1.IPAddressOrFQDN `json:"ipAddress,omitempty"`
	}{
		SmtpNetworkProxy: (*SmtpNetworkProxy)(p),
		IpAddress:        p.IpAddress,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *SmtpNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SmtpNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSmtpNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IpAddress != nil {
		p.IpAddress = known.IpAddress
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ipAddress")
	delete(allFields, "password")
	delete(allFields, "port")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSmtpNetwork() *SmtpNetwork {
	p := new(SmtpNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SmtpNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SMTP servers on a cluster. This is part of payload for cluster update operation only.
*/
type SmtpServerRef struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  SMTP email address.
	*/
	EmailAddress *string `json:"emailAddress"`

	Server *SmtpNetwork `json:"server"`

	Type *SmtpType `json:"type,omitempty"`
}

func (p *SmtpServerRef) MarshalJSON() ([]byte, error) {
	type SmtpServerRefProxy SmtpServerRef

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SmtpServerRefProxy
		EmailAddress *string      `json:"emailAddress,omitempty"`
		Server       *SmtpNetwork `json:"server,omitempty"`
	}{
		SmtpServerRefProxy: (*SmtpServerRefProxy)(p),
		EmailAddress:       p.EmailAddress,
		Server:             p.Server,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *SmtpServerRef) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SmtpServerRef
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSmtpServerRef()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EmailAddress != nil {
		p.EmailAddress = known.EmailAddress
	}
	if known.Server != nil {
		p.Server = known.Server
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "emailAddress")
	delete(allFields, "server")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSmtpServerRef() *SmtpServerRef {
	p := new(SmtpServerRef)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SmtpServerRef"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of SMTP server.
*/
type SmtpType int

const (
	SMTPTYPE_UNKNOWN  SmtpType = 0
	SMTPTYPE_REDACTED SmtpType = 1
	SMTPTYPE_PLAIN    SmtpType = 2
	SMTPTYPE_STARTTLS SmtpType = 3
	SMTPTYPE_SSL      SmtpType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SmtpType) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SmtpType) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SmtpType) index(name string) SmtpType {
	names := [...]string{
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

/*
SNMP user authentication type.
*/
type SnmpAuthType int

const (
	SNMPAUTHTYPE_UNKNOWN  SnmpAuthType = 0
	SNMPAUTHTYPE_REDACTED SnmpAuthType = 1
	SNMPAUTHTYPE_MD5      SnmpAuthType = 2
	SNMPAUTHTYPE_SHA      SnmpAuthType = 3
	SNMPAUTHTYPE_SHA224   SnmpAuthType = 4
	SNMPAUTHTYPE_SHA256   SnmpAuthType = 5
	SNMPAUTHTYPE_SHA384   SnmpAuthType = 6
	SNMPAUTHTYPE_SHA512   SnmpAuthType = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnmpAuthType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA",
		"SHA224",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SnmpAuthType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA",
		"SHA224",
		"SHA256",
		"SHA384",
		"SHA512",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SnmpAuthType) index(name string) SnmpAuthType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MD5",
		"SHA",
		"SHA224",
		"SHA256",
		"SHA384",
		"SHA512",
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

/*
SNMP information.
*/
type SnmpConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  SNMP status.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  SNMP transport details.
	*/
	Transports []SnmpTransport `json:"transports,omitempty"`
	/*
	  SNMP trap details.
	*/
	Traps []SnmpTrap `json:"traps,omitempty"`
	/*
	  SNMP user information.
	*/
	Users []SnmpUser `json:"users,omitempty"`
}

func (p *SnmpConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SnmpConfig

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

func (p *SnmpConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SnmpConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSnmpConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsEnabled != nil {
		p.IsEnabled = known.IsEnabled
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Transports != nil {
		p.Transports = known.Transports
	}
	if known.Traps != nil {
		p.Traps = known.Traps
	}
	if known.Users != nil {
		p.Users = known.Users
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "isEnabled")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "transports")
	delete(allFields, "traps")
	delete(allFields, "users")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSnmpConfig() *SnmpConfig {
	p := new(SnmpConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SNMP user encryption type.
*/
type SnmpPrivType int

const (
	SNMPPRIVTYPE_UNKNOWN  SnmpPrivType = 0
	SNMPPRIVTYPE_REDACTED SnmpPrivType = 1
	SNMPPRIVTYPE_DES      SnmpPrivType = 2
	SNMPPRIVTYPE_AES      SnmpPrivType = 3
	SNMPPRIVTYPE_AES192   SnmpPrivType = 4
	SNMPPRIVTYPE_AES256   SnmpPrivType = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnmpPrivType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DES",
		"AES",
		"AES192",
		"AES256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SnmpPrivType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DES",
		"AES",
		"AES192",
		"AES256",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SnmpPrivType) index(name string) SnmpPrivType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DES",
		"AES",
		"AES192",
		"AES256",
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

/*
SNMP protocol type.
*/
type SnmpProtocol int

const (
	SNMPPROTOCOL_UNKNOWN  SnmpProtocol = 0
	SNMPPROTOCOL_REDACTED SnmpProtocol = 1
	SNMPPROTOCOL_UDP      SnmpProtocol = 2
	SNMPPROTOCOL_UDP6     SnmpProtocol = 3
	SNMPPROTOCOL_TCP      SnmpProtocol = 4
	SNMPPROTOCOL_TCP6     SnmpProtocol = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnmpProtocol) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SnmpProtocol) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SnmpProtocol) index(name string) SnmpProtocol {
	names := [...]string{
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

/*
SNMP transport details.
*/
type SnmpTransport struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  SNMP port.
	*/
	Port *int `json:"port"`

	Protocol *SnmpProtocol `json:"protocol"`
}

func (p *SnmpTransport) MarshalJSON() ([]byte, error) {
	type SnmpTransportProxy SnmpTransport

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SnmpTransportProxy
		Port     *int          `json:"port,omitempty"`
		Protocol *SnmpProtocol `json:"protocol,omitempty"`
	}{
		SnmpTransportProxy: (*SnmpTransportProxy)(p),
		Port:               p.Port,
		Protocol:           p.Protocol,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *SnmpTransport) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SnmpTransport
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSnmpTransport()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Protocol != nil {
		p.Protocol = known.Protocol
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "port")
	delete(allFields, "protocol")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSnmpTransport() *SnmpTransport {
	p := new(SnmpTransport)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpTransport"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SNMP trap details.
*/
type SnmpTrap struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddress `json:"address"`
	/*
	  Community string(plaintext) for SNMP version 2.0.
	*/
	CommunityString *string `json:"communityString,omitempty"`
	/*
	  SNMP engine Id.
	*/
	EngineId *string `json:"engineId,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  SNMP port.
	*/
	Port *int `json:"port,omitempty"`

	Protocol *SnmpProtocol `json:"protocol,omitempty"`
	/*
	  SNMP receiver name.
	*/
	RecieverName *string `json:"recieverName,omitempty"`
	/*
	  SNMP information status.
	*/
	ShouldInform *bool `json:"shouldInform,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  SNMP username. For SNMP trap v3 version, SNMP username is required parameter.
	*/
	Username *string `json:"username,omitempty"`

	Version *SnmpTrapVersion `json:"version"`
}

func (p *SnmpTrap) MarshalJSON() ([]byte, error) {
	type SnmpTrapProxy SnmpTrap

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SnmpTrapProxy
		Address *import1.IPAddress `json:"address,omitempty"`
		Version *SnmpTrapVersion   `json:"version,omitempty"`
	}{
		SnmpTrapProxy: (*SnmpTrapProxy)(p),
		Address:       p.Address,
		Version:       p.Version,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *SnmpTrap) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SnmpTrap
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSnmpTrap()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Address != nil {
		p.Address = known.Address
	}
	if known.CommunityString != nil {
		p.CommunityString = known.CommunityString
	}
	if known.EngineId != nil {
		p.EngineId = known.EngineId
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Protocol != nil {
		p.Protocol = known.Protocol
	}
	if known.RecieverName != nil {
		p.RecieverName = known.RecieverName
	}
	if known.ShouldInform != nil {
		p.ShouldInform = known.ShouldInform
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Username != nil {
		p.Username = known.Username
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "address")
	delete(allFields, "communityString")
	delete(allFields, "engineId")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "port")
	delete(allFields, "protocol")
	delete(allFields, "recieverName")
	delete(allFields, "shouldInform")
	delete(allFields, "tenantId")
	delete(allFields, "username")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSnmpTrap() *SnmpTrap {
	p := new(SnmpTrap)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpTrap"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
SNMP version.
*/
type SnmpTrapVersion int

const (
	SNMPTRAPVERSION_UNKNOWN  SnmpTrapVersion = 0
	SNMPTRAPVERSION_REDACTED SnmpTrapVersion = 1
	SNMPTRAPVERSION_V2       SnmpTrapVersion = 2
	SNMPTRAPVERSION_V3       SnmpTrapVersion = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnmpTrapVersion) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SnmpTrapVersion) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SnmpTrapVersion) index(name string) SnmpTrapVersion {
	names := [...]string{
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

/*
SNMP user information.
*/
type SnmpUser struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  SNMP user authentication key.
	*/
	AuthKey *string `json:"authKey"`

	AuthType *SnmpAuthType `json:"authType"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  SNMP user encryption key.
	*/
	PrivKey *string `json:"privKey,omitempty"`

	PrivType *SnmpPrivType `json:"privType,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  SNMP username. For SNMP trap v3 version, SNMP username is required parameter.
	*/
	Username *string `json:"username"`
}

func (p *SnmpUser) MarshalJSON() ([]byte, error) {
	type SnmpUserProxy SnmpUser

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SnmpUserProxy
		AuthKey  *string       `json:"authKey,omitempty"`
		AuthType *SnmpAuthType `json:"authType,omitempty"`
		Username *string       `json:"username,omitempty"`
	}{
		SnmpUserProxy: (*SnmpUserProxy)(p),
		AuthKey:       p.AuthKey,
		AuthType:      p.AuthType,
		Username:      p.Username,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *SnmpUser) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SnmpUser
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSnmpUser()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AuthKey != nil {
		p.AuthKey = known.AuthKey
	}
	if known.AuthType != nil {
		p.AuthType = known.AuthType
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.PrivKey != nil {
		p.PrivKey = known.PrivKey
	}
	if known.PrivType != nil {
		p.PrivType = known.PrivType
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authKey")
	delete(allFields, "authType")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "privKey")
	delete(allFields, "privType")
	delete(allFields, "tenantId")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSnmpUser() *SnmpUser {
	p := new(SnmpUser)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpUser"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster software version details.
*/
type SoftwareMapReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SoftwareType *SoftwareTypeRef `json:"softwareType,omitempty"`
	/*
	  Software version.
	*/
	Version *string `json:"version,omitempty"`
}

func (p *SoftwareMapReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SoftwareMapReference

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

func (p *SoftwareMapReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SoftwareMapReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSoftwareMapReference()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.SoftwareType != nil {
		p.SoftwareType = known.SoftwareType
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "softwareType")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSoftwareMapReference() *SoftwareMapReference {
	p := new(SoftwareMapReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SoftwareMapReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Software type.
*/
type SoftwareTypeRef int

const (
	SOFTWARETYPEREF_UNKNOWN       SoftwareTypeRef = 0
	SOFTWARETYPEREF_REDACTED      SoftwareTypeRef = 1
	SOFTWARETYPEREF_NOS           SoftwareTypeRef = 2
	SOFTWARETYPEREF_NCC           SoftwareTypeRef = 3
	SOFTWARETYPEREF_PRISM_CENTRAL SoftwareTypeRef = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SoftwareTypeRef) name(index int) string {
	names := [...]string{
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

// Returns the name of the enum
func (e SoftwareTypeRef) GetName() string {
	index := int(e)
	names := [...]string{
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

// Returns the enum type given a string value
func (e *SoftwareTypeRef) index(name string) SoftwareTypeRef {
	names := [...]string{
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

/*
Configuration details for the storage configuration.
*/
type StorageConfigs struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Whether the delayed maintenance resiliency is enabled.
	*/
	IsDelayedMaintenanceResiliencyEnabled *bool `json:"isDelayedMaintenanceResiliencyEnabled"`
	/*
	  Whether the rebuild reservation is enabled.
	*/
	IsRebuildReservationEnabled *bool `json:"isRebuildReservationEnabled"`
	/*
	  Whether the RF1 is enabled.
	*/
	IsRf1Enabled *bool `json:"isRf1Enabled"`
	/*
	  The threshold percentage for the resilient capacity warning.
	*/
	ResilientCapacityWarningThresholdPercentage *int `json:"resilientCapacityWarningThresholdPercentage"`
	/*
	  The threshold for the storage over provisioning caution. Value must be -1 (default) or between 1 and 100.
	*/
	StorageOverProvisioningCautionThreshold *float64 `json:"storageOverProvisioningCautionThreshold"`
}

func (p *StorageConfigs) MarshalJSON() ([]byte, error) {
	type StorageConfigsProxy StorageConfigs

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StorageConfigsProxy
		IsDelayedMaintenanceResiliencyEnabled       *bool    `json:"isDelayedMaintenanceResiliencyEnabled,omitempty"`
		IsRebuildReservationEnabled                 *bool    `json:"isRebuildReservationEnabled,omitempty"`
		IsRf1Enabled                                *bool    `json:"isRf1Enabled,omitempty"`
		ResilientCapacityWarningThresholdPercentage *int     `json:"resilientCapacityWarningThresholdPercentage,omitempty"`
		StorageOverProvisioningCautionThreshold     *float64 `json:"storageOverProvisioningCautionThreshold,omitempty"`
	}{
		StorageConfigsProxy:                         (*StorageConfigsProxy)(p),
		IsDelayedMaintenanceResiliencyEnabled:       p.IsDelayedMaintenanceResiliencyEnabled,
		IsRebuildReservationEnabled:                 p.IsRebuildReservationEnabled,
		IsRf1Enabled:                                p.IsRf1Enabled,
		ResilientCapacityWarningThresholdPercentage: p.ResilientCapacityWarningThresholdPercentage,
		StorageOverProvisioningCautionThreshold:     p.StorageOverProvisioningCautionThreshold,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *StorageConfigs) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageConfigs
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageConfigs()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsDelayedMaintenanceResiliencyEnabled != nil {
		p.IsDelayedMaintenanceResiliencyEnabled = known.IsDelayedMaintenanceResiliencyEnabled
	}
	if known.IsRebuildReservationEnabled != nil {
		p.IsRebuildReservationEnabled = known.IsRebuildReservationEnabled
	}
	if known.IsRf1Enabled != nil {
		p.IsRf1Enabled = known.IsRf1Enabled
	}
	if known.ResilientCapacityWarningThresholdPercentage != nil {
		p.ResilientCapacityWarningThresholdPercentage = known.ResilientCapacityWarningThresholdPercentage
	}
	if known.StorageOverProvisioningCautionThreshold != nil {
		p.StorageOverProvisioningCautionThreshold = known.StorageOverProvisioningCautionThreshold
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isDelayedMaintenanceResiliencyEnabled")
	delete(allFields, "isRebuildReservationEnabled")
	delete(allFields, "isRf1Enabled")
	delete(allFields, "resilientCapacityWarningThresholdPercentage")
	delete(allFields, "storageOverProvisioningCautionThreshold")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageConfigs() *StorageConfigs {
	p := new(StorageConfigs)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StorageConfigs"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StorageConfigsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Whether the delayed maintenance resiliency is enabled.
	*/
	IsDelayedMaintenanceResiliencyEnabled *bool `json:"isDelayedMaintenanceResiliencyEnabled"`
	/*
	  Whether the rebuild reservation is enabled.
	*/
	IsRebuildReservationEnabled *bool `json:"isRebuildReservationEnabled"`
	/*
	  Whether the RF1 is enabled.
	*/
	IsRf1Enabled *bool `json:"isRf1Enabled"`
	/*
	  The threshold percentage for the resilient capacity warning.
	*/
	ResilientCapacityWarningThresholdPercentage *int `json:"resilientCapacityWarningThresholdPercentage"`
	/*
	  The threshold for the storage over provisioning caution. Value must be -1 (default) or between 1 and 100.
	*/
	StorageOverProvisioningCautionThreshold *float64 `json:"storageOverProvisioningCautionThreshold"`
}

func (p *StorageConfigsProjection) MarshalJSON() ([]byte, error) {
	type StorageConfigsProjectionProxy StorageConfigsProjection

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*StorageConfigsProjectionProxy
		IsDelayedMaintenanceResiliencyEnabled       *bool    `json:"isDelayedMaintenanceResiliencyEnabled,omitempty"`
		IsRebuildReservationEnabled                 *bool    `json:"isRebuildReservationEnabled,omitempty"`
		IsRf1Enabled                                *bool    `json:"isRf1Enabled,omitempty"`
		ResilientCapacityWarningThresholdPercentage *int     `json:"resilientCapacityWarningThresholdPercentage,omitempty"`
		StorageOverProvisioningCautionThreshold     *float64 `json:"storageOverProvisioningCautionThreshold,omitempty"`
	}{
		StorageConfigsProjectionProxy:               (*StorageConfigsProjectionProxy)(p),
		IsDelayedMaintenanceResiliencyEnabled:       p.IsDelayedMaintenanceResiliencyEnabled,
		IsRebuildReservationEnabled:                 p.IsRebuildReservationEnabled,
		IsRf1Enabled:                                p.IsRf1Enabled,
		ResilientCapacityWarningThresholdPercentage: p.ResilientCapacityWarningThresholdPercentage,
		StorageOverProvisioningCautionThreshold:     p.StorageOverProvisioningCautionThreshold,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *StorageConfigsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageConfigsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageConfigsProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsDelayedMaintenanceResiliencyEnabled != nil {
		p.IsDelayedMaintenanceResiliencyEnabled = known.IsDelayedMaintenanceResiliencyEnabled
	}
	if known.IsRebuildReservationEnabled != nil {
		p.IsRebuildReservationEnabled = known.IsRebuildReservationEnabled
	}
	if known.IsRf1Enabled != nil {
		p.IsRf1Enabled = known.IsRf1Enabled
	}
	if known.ResilientCapacityWarningThresholdPercentage != nil {
		p.ResilientCapacityWarningThresholdPercentage = known.ResilientCapacityWarningThresholdPercentage
	}
	if known.StorageOverProvisioningCautionThreshold != nil {
		p.StorageOverProvisioningCautionThreshold = known.StorageOverProvisioningCautionThreshold
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isDelayedMaintenanceResiliencyEnabled")
	delete(allFields, "isRebuildReservationEnabled")
	delete(allFields, "isRf1Enabled")
	delete(allFields, "resilientCapacityWarningThresholdPercentage")
	delete(allFields, "storageOverProvisioningCautionThreshold")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageConfigsProjection() *StorageConfigsProjection {
	p := new(StorageConfigsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StorageConfigsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Storage summary entity and its attribute related to cluster fault tolerance.
*/
type StorageSummary struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster fault tolerance defines how many simultaneous failures within a fault domain the cluster can sustain.
	*/
	ClusterFaultTolerantCapacityInBytes *int64 `json:"clusterFaultTolerantCapacityInBytes,omitempty"`
	/*
	  Returns the usage status of the cluster. Possible values are: 'OK', 'WARNING', 'CRITICAL' or empty.
	*/
	ClusterUsageStatus *string `json:"clusterUsageStatus,omitempty"`
}

func (p *StorageSummary) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StorageSummary

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

func (p *StorageSummary) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageSummary
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageSummary()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterFaultTolerantCapacityInBytes != nil {
		p.ClusterFaultTolerantCapacityInBytes = known.ClusterFaultTolerantCapacityInBytes
	}
	if known.ClusterUsageStatus != nil {
		p.ClusterUsageStatus = known.ClusterUsageStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterFaultTolerantCapacityInBytes")
	delete(allFields, "clusterUsageStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageSummary() *StorageSummary {
	p := new(StorageSummary)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StorageSummary"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StorageSummaryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster fault tolerance defines how many simultaneous failures within a fault domain the cluster can sustain.
	*/
	ClusterFaultTolerantCapacityInBytes *int64 `json:"clusterFaultTolerantCapacityInBytes,omitempty"`
	/*
	  Returns the usage status of the cluster. Possible values are: 'OK', 'WARNING', 'CRITICAL' or empty.
	*/
	ClusterUsageStatus *string `json:"clusterUsageStatus,omitempty"`
}

func (p *StorageSummaryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias StorageSummaryProjection

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

func (p *StorageSummaryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias StorageSummaryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewStorageSummaryProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterFaultTolerantCapacityInBytes != nil {
		p.ClusterFaultTolerantCapacityInBytes = known.ClusterFaultTolerantCapacityInBytes
	}
	if known.ClusterUsageStatus != nil {
		p.ClusterUsageStatus = known.ClusterUsageStatus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterFaultTolerantCapacityInBytes")
	delete(allFields, "clusterUsageStatus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewStorageSummaryProjection() *StorageSummaryProjection {
	p := new(StorageSummaryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StorageSummaryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Upgrade status of a cluster.
*/
type UpgradeStatus int

const (
	UPGRADESTATUS_UNKNOWN     UpgradeStatus = 0
	UPGRADESTATUS_REDACTED    UpgradeStatus = 1
	UPGRADESTATUS_PENDING     UpgradeStatus = 2
	UPGRADESTATUS_DOWNLOADING UpgradeStatus = 3
	UPGRADESTATUS_QUEUED      UpgradeStatus = 4
	UPGRADESTATUS_PREUPGRADE  UpgradeStatus = 5
	UPGRADESTATUS_UPGRADING   UpgradeStatus = 6
	UPGRADESTATUS_SUCCEEDED   UpgradeStatus = 7
	UPGRADESTATUS_FAILED      UpgradeStatus = 8
	UPGRADESTATUS_CANCELLED   UpgradeStatus = 9
	UPGRADESTATUS_SCHEDULED   UpgradeStatus = 10
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *UpgradeStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"DOWNLOADING",
		"QUEUED",
		"PREUPGRADE",
		"UPGRADING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"SCHEDULED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e UpgradeStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"DOWNLOADING",
		"QUEUED",
		"PREUPGRADE",
		"UPGRADING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"SCHEDULED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *UpgradeStatus) index(name string) UpgradeStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"DOWNLOADING",
		"QUEUED",
		"PREUPGRADE",
		"UPGRADING",
		"SUCCEEDED",
		"FAILED",
		"CANCELLED",
		"SCHEDULED",
	}
	for idx := range names {
		if names[idx] == name {
			return UpgradeStatus(idx)
		}
	}
	return UPGRADESTATUS_UNKNOWN
}

func (e *UpgradeStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for UpgradeStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *UpgradeStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e UpgradeStatus) Ref() *UpgradeStatus {
	return &e
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
