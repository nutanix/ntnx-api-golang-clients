/*
 * Generated file models/lifecycle/v4/svcmgr/svcmgr_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module lifecycle.v4.svcmgr of Nutanix Lifecycle Management APIs
*/
package svcmgr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/error"
	"time"
)

/*
Ahv client configuration.
*/
type AhvClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  client config CA (Certificate Authority) chain.
	*/
	CaChain *string `json:"caChain,omitempty"`
	/*
	  X.509 certificate to be used for accessing the cluster.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/*
	  X.509 key to be used for accessing the cluster.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Password to be used for accessing the cluster.
	*/
	Password *string `json:"password,omitempty"`

	PrismIp *import1.IPv4Address `json:"prismIp,omitempty"`
	/*
	  Prism Element Port.
	*/
	PrismPort *int64 `json:"prismPort,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`
	/*
	  UserName to be used for accessing the cluster.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *AhvClientConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AhvClientConfig

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

func (p *AhvClientConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AhvClientConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AhvClientConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "caChain")
	delete(allFields, "certificate")
	delete(allFields, "key")
	delete(allFields, "password")
	delete(allFields, "prismIp")
	delete(allFields, "prismPort")
	delete(allFields, "serviceName")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAhvClientConfig() *AhvClientConfig {
	p := new(AhvClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network resources for Ahv.
*/
type AhvNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`
	/*
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`

	NicType *NicType `json:"nicType,omitempty"`
}

func (p *AhvNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type AhvNetworkResourceConfigProxy AhvNetworkResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AhvNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		AhvNetworkResourceConfigProxy: (*AhvNetworkResourceConfigProxy)(p),
		Network:                       p.Network,
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

func (p *AhvNetworkResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AhvNetworkResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AhvNetworkResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "network")
	delete(allFields, "nicType")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAhvNetworkResourceConfig() *AhvNetworkResourceConfig {
	p := new(AhvNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
AHV platform specific config.
*/
type AhvResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkResource *AhvNetworkResourceConfig `json:"networkResource"`

	VmResource *VmResourceConfig `json:"vmResource,omitempty"`
}

func (p *AhvResourceConfig) MarshalJSON() ([]byte, error) {
	type AhvResourceConfigProxy AhvResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AhvResourceConfigProxy
		NetworkResource *AhvNetworkResourceConfig `json:"networkResource,omitempty"`
	}{
		AhvResourceConfigProxy: (*AhvResourceConfigProxy)(p),
		NetworkResource:        p.NetworkResource,
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

func (p *AhvResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AhvResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = AhvResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkResource")
	delete(allFields, "vmResource")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewAhvResourceConfig() *AhvResourceConfig {
	p := new(AhvResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Application struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	IsInactive *bool `json:"isInactive,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	SubServices []Service `json:"subServices,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *Application) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Application

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

func (p *Application) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Application
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Application(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "isInactive")
	delete(allFields, "lastUpdatedTimestamp")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "subServices")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApplication() *Application {
	p := new(Application)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Application"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ApplicationProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	IsInactive *bool `json:"isInactive,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	SubServices []Service `json:"subServices,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *ApplicationProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApplicationProjection

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

func (p *ApplicationProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApplicationProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ApplicationProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "isInactive")
	delete(allFields, "lastUpdatedTimestamp")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "subServices")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewApplicationProjection() *ApplicationProjection {
	p := new(ApplicationProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ApplicationProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type BaseClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  X.509 certificate to be used for accessing the cluster.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/*
	  X.509 key to be used for accessing the cluster.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Password to be used for accessing the cluster.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  UserName to be used for accessing the cluster.
	*/
	Username *string `json:"username,omitempty"`
}

func (p *BaseClientConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BaseClientConfig

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

func (p *BaseClientConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseClientConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = BaseClientConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "certificate")
	delete(allFields, "key")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBaseClientConfig() *BaseClientConfig {
	p := new(BaseClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.BaseClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Access configuration of the client used to access AHV or ESX clusters.
*/
type ClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`

	Config *OneOfClientConfigConfig `json:"config"`
	/*
	  Name of the client configuration.
	*/
	Name *string `json:"name"`
	/*
	  Type of platform client.
	*/
	Type *string `json:"type"`
}

func (p *ClientConfig) MarshalJSON() ([]byte, error) {
	type ClientConfigProxy ClientConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClientConfigProxy
		Config *OneOfClientConfigConfig `json:"config,omitempty"`
		Name   *string                  `json:"name,omitempty"`
		Type   *string                  `json:"type,omitempty"`
	}{
		ClientConfigProxy: (*ClientConfigProxy)(p),
		Config:            p.Config,
		Name:              p.Name,
		Type:              p.Type,
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

func (p *ClientConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClientConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClientConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configItemDiscriminator")
	delete(allFields, "config")
	delete(allFields, "name")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClientConfig() *ClientConfig {
	p := new(ClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClientConfig) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *ClientConfig) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfClientConfigConfig()
	}
	e := p.Config.SetValue(v)
	if nil == e {
		if nil == p.ConfigItemDiscriminator_ {
			p.ConfigItemDiscriminator_ = new(string)
		}
		*p.ConfigItemDiscriminator_ = *p.Config.Discriminator
	}
	return e
}

/*
A model that represents a Service Cluster for Micro Services Platform(MSP).
*/
type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ControlPlaneConfig *Config `json:"ControlPlaneConfig"`
	/*
	  Key Value pairs that can be used to customize cluster deployment configuration.
	Supported annotations are -
	  - Select storage provisioner service. Accepted values are csi/volume_provisioner.
	    Key: "msp/storage/service"
	    Value: "volume_provisioner/csi"

	  - Disable logging add-ons on the cluster.
	    Key: "msp/logging/disable"
	    Value: "true"

	  - Disable monitoring add-ons on the cluster.
	    Key: "msp/monitoring/disable"
	    Value: "true"
	*/
	Annotations []import1.KVStringPair `json:"annotations,omitempty"`
	/*
	  Client spec needed for VM/Network CRUD operations.
	*/
	ClientConfigs []ClientConfig `json:"clientConfigs,omitempty"`

	ClusterType *ClusterType `json:"clusterType,omitempty"`
	/*
	  The controller version with which the cluster was deployed.
	*/
	ControllerVersion *string `json:"controllerVersion,omitempty"`

	DeploymentType *ClusterDeploymentType `json:"deploymentType,omitempty"`
	/*
	  Brief description about the cluster.
	*/
	Description *string `json:"description,omitempty"`

	Domain *import1.FQDN `json:"domain,omitempty"`
	/*
	  The envoy nodes ips.
	*/
	EnvoyIps []string `json:"envoyIps,omitempty"`
	/*
	  The etcd nodes ips.
	*/
	EtcdIps []string `json:"etcdIps,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The external master ip.
	*/
	ExternalMasterIp *string `json:"externalMasterIp,omitempty"`
	/*
	  Boolean value indicating if cluster is locked or not.
	*/
	IsLockedDown *bool `json:"isLockedDown,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	LoadBalancerConfigs *LoadBalancerConfig `json:"loadBalancerConfigs,omitempty"`
	/*
	  Configuration of the default volume to be attached to a node for logging purpose.
	*/
	LoggingVolumeConfig []LoggingVolumeConfig `json:"loggingVolumeConfig,omitempty"`

	ManagementIps *ClusterManagementIps `json:"managementIps,omitempty"`
	/*
	  The control plane node ips.
	*/
	MasterIps []string `json:"masterIps,omitempty"`
	/*
	  Name of the cluster.
	*/
	Name *string `json:"name"`
	/*
	  Network configuration of the cluster.
	*/
	NetworkConfigs []ClusterNetworkConfig `json:"networkConfigs,omitempty"`
	/*
	  Resource configuration list for the VM resources.
	*/
	ResourceConfigs []ClusterResourceConfig `json:"resourceConfigs,omitempty"`
	/*
	  The status of deployed cluster.
	*/
	Status *string `json:"status,omitempty"`
	/*
	  The storage class specification used for creation of volumes needed for stateful services.
	*/
	StorageClasses []ClusterStorageClass `json:"storageClasses,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The version of the cluster.
	*/
	Version *string `json:"version,omitempty"`

	WorkerConfigs []Config `json:"workerConfigs,omitempty"`
	/*
	  The worker nodes ips.
	*/
	WorkerIps []string `json:"workerIps,omitempty"`
}

func (p *Cluster) MarshalJSON() ([]byte, error) {
	type ClusterProxy Cluster

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterProxy
		ControlPlaneConfig *Config `json:"ControlPlaneConfig,omitempty"`
		Name               *string `json:"name,omitempty"`
	}{
		ClusterProxy:       (*ClusterProxy)(p),
		ControlPlaneConfig: p.ControlPlaneConfig,
		Name:               p.Name,
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
	*p = Cluster(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "ControlPlaneConfig")
	delete(allFields, "annotations")
	delete(allFields, "clientConfigs")
	delete(allFields, "clusterType")
	delete(allFields, "controllerVersion")
	delete(allFields, "deploymentType")
	delete(allFields, "description")
	delete(allFields, "domain")
	delete(allFields, "envoyIps")
	delete(allFields, "etcdIps")
	delete(allFields, "extId")
	delete(allFields, "externalMasterIp")
	delete(allFields, "isLockedDown")
	delete(allFields, "links")
	delete(allFields, "loadBalancerConfigs")
	delete(allFields, "loggingVolumeConfig")
	delete(allFields, "managementIps")
	delete(allFields, "masterIps")
	delete(allFields, "name")
	delete(allFields, "networkConfigs")
	delete(allFields, "resourceConfigs")
	delete(allFields, "status")
	delete(allFields, "storageClasses")
	delete(allFields, "tenantId")
	delete(allFields, "version")
	delete(allFields, "workerConfigs")
	delete(allFields, "workerIps")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enum representing the type of cluster deployment.
*/
type ClusterDeploymentType int

const (
	CLUSTERDEPLOYMENTTYPE_UNKNOWN   ClusterDeploymentType = 0
	CLUSTERDEPLOYMENTTYPE_REDACTED  ClusterDeploymentType = 1
	CLUSTERDEPLOYMENTTYPE_COMBINED  ClusterDeploymentType = 2
	CLUSTERDEPLOYMENTTYPE_SCALE_OUT ClusterDeploymentType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterDeploymentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMBINED",
		"SCALE_OUT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterDeploymentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMBINED",
		"SCALE_OUT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterDeploymentType) index(name string) ClusterDeploymentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMBINED",
		"SCALE_OUT",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterDeploymentType(idx)
		}
	}
	return CLUSTERDEPLOYMENTTYPE_UNKNOWN
}

func (e *ClusterDeploymentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterDeploymentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterDeploymentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterDeploymentType) Ref() *ClusterDeploymentType {
	return &e
}

/*
IPs to manage the cluster.
*/
type ClusterManagementIps struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MasterVip *import1.IPv4Address `json:"masterVip,omitempty"`

	MspDnsIp *import1.IPv4Address `json:"mspDnsIp,omitempty"`
	/*
	  IPAM Name to be used for fetching static IPs that are not provided in the payload.
	*/
	Network *string `json:"network,omitempty"`
}

func (p *ClusterManagementIps) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ClusterManagementIps

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

func (p *ClusterManagementIps) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterManagementIps
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterManagementIps(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "masterVip")
	delete(allFields, "mspDnsIp")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterManagementIps() *ClusterManagementIps {
	p := new(ClusterManagementIps)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterManagementIps"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cluster Network configuration.
*/
type ClusterNetworkConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of DNS servers to be used by this configuration.
	*/
	DnsServers []import1.IPv4Address `json:"dnsServers,omitempty"`

	EsxLocation *VcenterLocation `json:"esxLocation,omitempty"`

	GatewayIp *import1.IPv4Address `json:"gatewayIp,omitempty"`

	IpRanges []IpRange `json:"ipRanges,omitempty"`
	/*
	  Name of the network configuration.
	*/
	Name *string `json:"name"`

	Netmask *import1.IPv4Address `json:"netmask,omitempty"`
	/*
	  Network managed by the IPAM (IP address management) entity.
	*/
	Network *string `json:"network"`
}

func (p *ClusterNetworkConfig) MarshalJSON() ([]byte, error) {
	type ClusterNetworkConfigProxy ClusterNetworkConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterNetworkConfigProxy
		Name    *string `json:"name,omitempty"`
		Network *string `json:"network,omitempty"`
	}{
		ClusterNetworkConfigProxy: (*ClusterNetworkConfigProxy)(p),
		Name:                      p.Name,
		Network:                   p.Network,
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

func (p *ClusterNetworkConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterNetworkConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterNetworkConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "dnsServers")
	delete(allFields, "esxLocation")
	delete(allFields, "gatewayIp")
	delete(allFields, "ipRanges")
	delete(allFields, "name")
	delete(allFields, "netmask")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterNetworkConfig() *ClusterNetworkConfig {
	p := new(ClusterNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of client config to use.
	*/
	Client *string `json:"client,omitempty"`
	/*
	  Labels for this resource config. Its a key value pair of strings
	*/
	Labels []import1.KVStringPair `json:"labels,omitempty"`
	/*
	  Name of resource config.
	*/
	Name *string `json:"name"`

	ResourceConfigItemDiscriminator_ *string `json:"$resourceConfigItemDiscriminator,omitempty"`

	ResourceConfig *OneOfClusterResourceConfigResourceConfig `json:"resourceConfig,omitempty"`
}

func (p *ClusterResourceConfig) MarshalJSON() ([]byte, error) {
	type ClusterResourceConfigProxy ClusterResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterResourceConfigProxy
		Name *string `json:"name,omitempty"`
	}{
		ClusterResourceConfigProxy: (*ClusterResourceConfigProxy)(p),
		Name:                       p.Name,
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

func (p *ClusterResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "client")
	delete(allFields, "labels")
	delete(allFields, "name")
	delete(allFields, "$resourceConfigItemDiscriminator")
	delete(allFields, "resourceConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterResourceConfig() *ClusterResourceConfig {
	p := new(ClusterResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterResourceConfig) GetResourceConfig() interface{} {
	if nil == p.ResourceConfig {
		return nil
	}
	return p.ResourceConfig.GetValue()
}

func (p *ClusterResourceConfig) SetResourceConfig(v interface{}) error {
	if nil == p.ResourceConfig {
		p.ResourceConfig = NewOneOfClusterResourceConfigResourceConfig()
	}
	e := p.ResourceConfig.SetValue(v)
	if nil == e {
		if nil == p.ResourceConfigItemDiscriminator_ {
			p.ResourceConfigItemDiscriminator_ = new(string)
		}
		*p.ResourceConfigItemDiscriminator_ = *p.ResourceConfig.Discriminator
	}
	return e
}

/*
Storage Class to be used by the cluster.
*/
type ClusterStorageClass struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Description *string `json:"description,omitempty"`
	/*
	  Whether this storage class is marked as the default one.
	*/
	IsDefaultStorageClass *bool `json:"isDefaultStorageClass,omitempty"`
	/*
	  MountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"].
	*/
	MountOptions *string `json:"mountOptions,omitempty"`
	/*
	  Unique name of the storage class. The name is used to refer to the storage class.
	*/
	Name *string `json:"name"`

	NutanixStorageClass *NutanixStorageClass `json:"nutanixStorageClass,omitempty"`
	/*
	  Controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class.
	*/
	ReclaimPolicy *string `json:"reclaimPolicy,omitempty"`
	/*
	  Storage type for the storage class. e.g. NutanixVolumes
	*/
	StorageType *string `json:"storageType,omitempty"`
}

func (p *ClusterStorageClass) MarshalJSON() ([]byte, error) {
	type ClusterStorageClassProxy ClusterStorageClass

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterStorageClassProxy
		Name *string `json:"name,omitempty"`
	}{
		ClusterStorageClassProxy: (*ClusterStorageClassProxy)(p),
		Name:                     p.Name,
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

func (p *ClusterStorageClass) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterStorageClass
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterStorageClass(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "description")
	delete(allFields, "isDefaultStorageClass")
	delete(allFields, "mountOptions")
	delete(allFields, "name")
	delete(allFields, "nutanixStorageClass")
	delete(allFields, "reclaimPolicy")
	delete(allFields, "storageType")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterStorageClass() *ClusterStorageClass {
	p := new(ClusterStorageClass)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterStorageClass"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the cluster.
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN  ClusterType = 0
	CLUSTERTYPE_REDACTED ClusterType = 1
	CLUSTERTYPE_CMSP     ClusterType = 2
	CLUSTERTYPE_SMSP     ClusterType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CMSP",
		"SMSP",
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
		"CMSP",
		"SMSP",
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
		"CMSP",
		"SMSP",
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

/*
Cluster upgrade information.
*/
type ClusterUpgradeInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  List of compatible versions for the cluster.
	*/
	CompatibleVersions []string `json:"compatibleVersions"`
	/*
	  Current version of the cluster.
	*/
	CurrentVersion *string `json:"currentVersion"`
}

func (p *ClusterUpgradeInfo) MarshalJSON() ([]byte, error) {
	type ClusterUpgradeInfoProxy ClusterUpgradeInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterUpgradeInfoProxy
		ClusterExtId       *string  `json:"clusterExtId,omitempty"`
		CompatibleVersions []string `json:"compatibleVersions,omitempty"`
		CurrentVersion     *string  `json:"currentVersion,omitempty"`
	}{
		ClusterUpgradeInfoProxy: (*ClusterUpgradeInfoProxy)(p),
		ClusterExtId:            p.ClusterExtId,
		CompatibleVersions:      p.CompatibleVersions,
		CurrentVersion:          p.CurrentVersion,
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

func (p *ClusterUpgradeInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterUpgradeInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ClusterUpgradeInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "compatibleVersions")
	delete(allFields, "currentVersion")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterUpgradeInfo() *ClusterUpgradeInfo {
	p := new(ClusterUpgradeInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterUpgradeInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ComponentDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Error details if any.
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/*
	  Time when the first health failure is seen.
	*/
	FirstFailureTime *time.Time `json:"firstFailureTime,omitempty"`
	/*
	  Messages published by the component.
	*/
	Messages []string `json:"messages,omitempty"`
	/*
	  Number of failures since the last success.
	*/
	NumFailures *int `json:"numFailures,omitempty"`
	/*
	  Time when the health is checked.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *ComponentDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ComponentDetails

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

func (p *ComponentDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ComponentDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ComponentDetails(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "errorMessage")
	delete(allFields, "firstFailureTime")
	delete(allFields, "messages")
	delete(allFields, "numFailures")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewComponentDetails() *ComponentDetails {
	p := new(ComponentDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ComponentDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Config struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NumInstances *int64 `json:"numInstances"`
	/*
	  This describes the name of the resource to be used.
	Actual spec is got from cluster resource configs field using this name.
	*/
	ResourceConfigName *string `json:"resourceConfigName"`
}

func (p *Config) MarshalJSON() ([]byte, error) {
	type ConfigProxy Config

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ConfigProxy
		NumInstances       *int64  `json:"numInstances,omitempty"`
		ResourceConfigName *string `json:"resourceConfigName,omitempty"`
	}{
		ConfigProxy:        (*ConfigProxy)(p),
		NumInstances:       p.NumInstances,
		ResourceConfigName: p.ResourceConfigName,
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

func (p *Config) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Config
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Config(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "numInstances")
	delete(allFields, "resourceConfigName")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewConfig() *Config {
	p := new(Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CustomValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Service *string `json:"service,omitempty"`

	Values []CustomValueItem `json:"values,omitempty"`
}

func (p *CustomValue) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CustomValue

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

func (p *CustomValue) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomValue
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CustomValue(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "service")
	delete(allFields, "values")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCustomValue() *CustomValue {
	p := new(CustomValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.CustomValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CustomValueItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Key *string `json:"key,omitempty"`

	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`

	Value *OneOfCustomValueItemValue `json:"value,omitempty"`
}

func (p *CustomValueItem) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CustomValueItem

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

func (p *CustomValueItem) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CustomValueItem
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = CustomValueItem(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "key")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewCustomValueItem() *CustomValueItem {
	p := new(CustomValueItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.CustomValueItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CustomValueItem) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *CustomValueItem) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfCustomValueItemValue()
	}
	e := p.Value.SetValue(v)
	if nil == e {
		if nil == p.ValueItemDiscriminator_ {
			p.ValueItemDiscriminator_ = new(string)
		}
		*p.ValueItemDiscriminator_ = *p.Value.Discriminator
	}
	return e
}

type Deployment int

const (
	DEPLOYMENT_ONPREM  Deployment = 0
	DEPLOYMENT_AZURE   Deployment = 1
	DEPLOYMENT_AWS     Deployment = 2
	DEPLOYMENT_UNKNOWN Deployment = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Deployment) name(index int) string {
	names := [...]string{
		"ONPREM",
		"AZURE",
		"AWS",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Deployment) GetName() string {
	index := int(e)
	names := [...]string{
		"ONPREM",
		"AZURE",
		"AWS",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Deployment) index(name string) Deployment {
	names := [...]string{
		"ONPREM",
		"AZURE",
		"AWS",
		"$UNKNOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return Deployment(idx)
		}
	}
	return DEPLOYMENT_UNKNOWN
}

func (e *Deployment) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Deployment:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Deployment) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Deployment) Ref() *Deployment {
	return &e
}

/*
Details of the entity.
*/
type EntityReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the entity.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Entity type identified as 'namespace:module[:submodule]:entityType'. For example - vmm:ahv:vm, where vmm is the namepsace, ahv is the module and vm is the entitytype.
	*/
	Rel *string `json:"rel,omitempty"`
}

func (p *EntityReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EntityReference

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

func (p *EntityReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EntityReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EntityReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Environment int

const (
	ENVIRONMENT_DEV     Environment = 0
	ENVIRONMENT_STAGING Environment = 1
	ENVIRONMENT_RELEASE Environment = 2
	ENVIRONMENT_UNKNOWN Environment = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Environment) name(index int) string {
	names := [...]string{
		"DEV",
		"STAGING",
		"RELEASE",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Environment) GetName() string {
	index := int(e)
	names := [...]string{
		"DEV",
		"STAGING",
		"RELEASE",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Environment) index(name string) Environment {
	names := [...]string{
		"DEV",
		"STAGING",
		"RELEASE",
		"$UNKNOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return Environment(idx)
		}
	}
	return ENVIRONMENT_UNKNOWN
}

func (e *Environment) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Environment:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Environment) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Environment) Ref() *Environment {
	return &e
}

/*
Esx client configuration.
*/
type EsxClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  X.509 certificate to be used for accessing the cluster.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/*
	  VCenter extension Key.
	*/
	ExtensionKey *string `json:"extensionKey,omitempty"`
	/*
	  For insecure tls connection.
	*/
	IsInsecure *bool `json:"isInsecure,omitempty"`
	/*
	  X.509 key to be used for accessing the cluster.
	*/
	Key *string `json:"key,omitempty"`
	/*
	  Password to be used for accessing the cluster.
	*/
	Password *string `json:"password,omitempty"`
	/*
	  TLS certificate thumbprint.
	*/
	ThumbPrint *string `json:"thumbPrint,omitempty"`
	/*
	  UserName to be used for accessing the cluster.
	*/
	Username *string `json:"username,omitempty"`

	VcenterIp *import1.IPv4Address `json:"vcenterIp,omitempty"`
	/*
	  Vcenter Port.
	*/
	VcenterPort *int `json:"vcenterPort,omitempty"`
}

func (p *EsxClientConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EsxClientConfig

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

func (p *EsxClientConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EsxClientConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EsxClientConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "certificate")
	delete(allFields, "extensionKey")
	delete(allFields, "isInsecure")
	delete(allFields, "key")
	delete(allFields, "password")
	delete(allFields, "thumbPrint")
	delete(allFields, "username")
	delete(allFields, "vcenterIp")
	delete(allFields, "vcenterPort")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEsxClientConfig() *EsxClientConfig {
	p := new(EsxClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network resources for Esx.
*/
type EsxNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`

	Location *VcenterLocation `json:"location,omitempty"`
	/*
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`
}

func (p *EsxNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type EsxNetworkResourceConfigProxy EsxNetworkResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EsxNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		EsxNetworkResourceConfigProxy: (*EsxNetworkResourceConfigProxy)(p),
		Network:                       p.Network,
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

func (p *EsxNetworkResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EsxNetworkResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EsxNetworkResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "location")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEsxNetworkResourceConfig() *EsxNetworkResourceConfig {
	p := new(EsxNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
ESX platform specific config.
*/
type EsxResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkResource *EsxNetworkResourceConfig `json:"networkResource"`

	VmResource *VmResourceConfig `json:"vmResource,omitempty"`
}

func (p *EsxResourceConfig) MarshalJSON() ([]byte, error) {
	type EsxResourceConfigProxy EsxResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*EsxResourceConfigProxy
		NetworkResource *EsxNetworkResourceConfig `json:"networkResource,omitempty"`
	}{
		EsxResourceConfigProxy: (*EsxResourceConfigProxy)(p),
		NetworkResource:        p.NetworkResource,
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

func (p *EsxResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EsxResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = EsxResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkResource")
	delete(allFields, "vmResource")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewEsxResourceConfig() *EsxResourceConfig {
	p := new(EsxResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Health struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HealthComponents []HealthComponent `json:"healthComponents,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (p *Health) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Health

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

func (p *Health) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Health
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Health(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "healthComponents")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHealth() *Health {
	p := new(Health)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Health"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Health of the Component.
*/
type HealthComponent struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Details *ComponentDetails `json:"details,omitempty"`
	/*
	  Name of the component.
	*/
	Name *string `json:"name,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (p *HealthComponent) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HealthComponent

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

func (p *HealthComponent) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HealthComponent
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HealthComponent(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "details")
	delete(allFields, "name")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHealthComponent() *HealthComponent {
	p := new(HealthComponent)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HealthComponent"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type History struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	Message *string `json:"message,omitempty"`

	Name *string `json:"name,omitempty"`

	ParentAppUuid *string `json:"parentAppUuid,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *History) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias History

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

func (p *History) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias History
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = History(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "message")
	delete(allFields, "name")
	delete(allFields, "parentAppUuid")
	delete(allFields, "status")
	delete(allFields, "taskUuid")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHistory() *History {
	p := new(History)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.History"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type HistoryProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	Message *string `json:"message,omitempty"`

	Name *string `json:"name,omitempty"`

	ParentAppUuid *string `json:"parentAppUuid,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *HistoryProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias HistoryProjection

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

func (p *HistoryProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HistoryProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = HistoryProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "message")
	delete(allFields, "name")
	delete(allFields, "parentAppUuid")
	delete(allFields, "status")
	delete(allFields, "taskUuid")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewHistoryProjection() *HistoryProjection {
	p := new(HistoryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HistoryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Info struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AirGapServerUri *string `json:"airGapServerUri,omitempty"`

	ChartRepositoryUri *string `json:"chartRepositoryUri,omitempty"`

	Deployment *Deployment `json:"deployment,omitempty"`

	Environment *Environment `json:"environment,omitempty"`

	IsAirGapEnabled *bool `json:"isAirGapEnabled,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *Info) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Info

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

func (p *Info) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Info
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Info(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "airGapServerUri")
	delete(allFields, "chartRepositoryUri")
	delete(allFields, "deployment")
	delete(allFields, "environment")
	delete(allFields, "isAirGapEnabled")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewInfo() *Info {
	p := new(Info)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Info"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Install struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name(or UUID) of the cluster. Optional. If not specified, the default cluster(CMSP) will be used.
	*/
	ClusterID *string `json:"clusterID,omitempty"`

	CustomValues []CustomValue `json:"customValues,omitempty"`

	Name *string `json:"name"`

	Version *string `json:"version"`
}

func (p *Install) MarshalJSON() ([]byte, error) {
	type InstallProxy Install

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*InstallProxy
		Name    *string `json:"name,omitempty"`
		Version *string `json:"version,omitempty"`
	}{
		InstallProxy: (*InstallProxy)(p),
		Name:         p.Name,
		Version:      p.Version,
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

func (p *Install) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Install
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Install(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterID")
	delete(allFields, "customValues")
	delete(allFields, "name")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewInstall() *Install {
	p := new(Install)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Install"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Range of Ips that can be used by the configuration.
*/
type IpRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *import1.IPv4Address `json:"endIp"`

	StartIp *import1.IPv4Address `json:"startIp"`
}

func (p *IpRange) MarshalJSON() ([]byte, error) {
	type IpRangeProxy IpRange

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*IpRangeProxy
		EndIp   *import1.IPv4Address `json:"endIp,omitempty"`
		StartIp *import1.IPv4Address `json:"startIp,omitempty"`
	}{
		IpRangeProxy: (*IpRangeProxy)(p),
		EndIp:        p.EndIp,
		StartIp:      p.StartIp,
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

func (p *IpRange) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias IpRange
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = IpRange(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endIp")
	delete(allFields, "startIp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewIpRange() *IpRange {
	p := new(IpRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.IpRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Load balancer representing id and network config for it.
*/
type LoadBalancer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Identifier for each load balancer. This is generated during cluster creation.
	*/
	Id *string `json:"id,omitempty"`

	NetworkConfig *LoadBalancerNetworkConfig `json:"networkConfig,omitempty"`
}

func (p *LoadBalancer) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LoadBalancer

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

func (p *LoadBalancer) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancer
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancer(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "id")
	delete(allFields, "networkConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancer() *LoadBalancer {
	p := new(LoadBalancer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An access interface is a user's notion of a named load balancer network. It abstracts
out the network details from developers and service deployers (dev-ops).
The name should be something that defines the positioning of the load balancer
in the data center. e.g. "corp", "wan", "dc", "it", "it-hr". Currently in a given
subdomain - the interface should be unique - which greatly simplifies what
the service deployer has to specify to expose the service.
A single load balancer defines one or more interfaces at load balancer creation time.
*/
type LoadBalancerAccessInterface struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  CIDR (Classless Inter-Domain Routing) address of the load balancer.
	*/
	Cidr *string `json:"cidr,omitempty"`
	/*
	  Access Interface IPs of the load balancer.
	*/
	Ips []import1.IPv4Address `json:"ips,omitempty"`
	/*
	  Network identifier of the load balancer.
	*/
	Network *string `json:"network"`
	/*
	  Subnet configuration of the load balancer.
	*/
	Subnet *string `json:"subnet"`
}

func (p *LoadBalancerAccessInterface) MarshalJSON() ([]byte, error) {
	type LoadBalancerAccessInterfaceProxy LoadBalancerAccessInterface

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LoadBalancerAccessInterfaceProxy
		Network *string `json:"network,omitempty"`
		Subnet  *string `json:"subnet,omitempty"`
	}{
		LoadBalancerAccessInterfaceProxy: (*LoadBalancerAccessInterfaceProxy)(p),
		Network:                          p.Network,
		Subnet:                           p.Subnet,
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

func (p *LoadBalancerAccessInterface) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancerAccessInterface
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancerAccessInterface(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cidr")
	delete(allFields, "ips")
	delete(allFields, "network")
	delete(allFields, "subnet")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancerAccessInterface() *LoadBalancerAccessInterface {
	p := new(LoadBalancerAccessInterface)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerAccessInterface"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
null
*/
type LoadBalancerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LoadBalancerConfigItems []LoadBalancerConfigObject `json:"loadBalancerConfigItems"`

	Version *string `json:"version,omitempty"`
}

func (p *LoadBalancerConfig) MarshalJSON() ([]byte, error) {
	type LoadBalancerConfigProxy LoadBalancerConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LoadBalancerConfigProxy
		LoadBalancerConfigItems []LoadBalancerConfigObject `json:"loadBalancerConfigItems,omitempty"`
	}{
		LoadBalancerConfigProxy: (*LoadBalancerConfigProxy)(p),
		LoadBalancerConfigItems: p.LoadBalancerConfigItems,
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

func (p *LoadBalancerConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancerConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancerConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "loadBalancerConfigItems")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancerConfig() *LoadBalancerConfig {
	p := new(LoadBalancerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Load balancer configuration of the cluster.
*/
type LoadBalancerConfigObject struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Action to be taken on load balancer.
	*/
	Action *string `json:"Action,omitempty"`
	/*
	  Kind of LB such as envoyproxy etc.
	*/
	Kind *string `json:"Kind,omitempty"`

	ConfigType *LoadBalancerConfigType `json:"configType,omitempty"`

	Instances []LoadBalancer `json:"instances"`
	/*
	  UUID of the MSP wanting to use this LB.
	*/
	MspUuid *string `json:"mspUuid,omitempty"`
	/*
	  Name of the Load Balancer configuration. The name should be unique for every configuration.
	*/
	Name *string `json:"name"`

	ResourceConfig *LoadBalancerResourceConfig `json:"resourceConfig,omitempty"`

	Type *LoadBalancerType `json:"type,omitempty"`
}

func (p *LoadBalancerConfigObject) MarshalJSON() ([]byte, error) {
	type LoadBalancerConfigObjectProxy LoadBalancerConfigObject

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LoadBalancerConfigObjectProxy
		Instances []LoadBalancer `json:"instances,omitempty"`
		Name      *string        `json:"name,omitempty"`
	}{
		LoadBalancerConfigObjectProxy: (*LoadBalancerConfigObjectProxy)(p),
		Instances:                     p.Instances,
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

func (p *LoadBalancerConfigObject) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancerConfigObject
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancerConfigObject(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "Action")
	delete(allFields, "Kind")
	delete(allFields, "configType")
	delete(allFields, "instances")
	delete(allFields, "mspUuid")
	delete(allFields, "name")
	delete(allFields, "resourceConfig")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancerConfigObject() *LoadBalancerConfigObject {
	p := new(LoadBalancerConfigObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerConfigObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines the type of configuration, static, dynamic etc.
*/
type LoadBalancerConfigType int

const (
	LOADBALANCERCONFIGTYPE_UNKNOWN  LoadBalancerConfigType = 0
	LOADBALANCERCONFIGTYPE_REDACTED LoadBalancerConfigType = 1
	LOADBALANCERCONFIGTYPE_STATIC   LoadBalancerConfigType = 2
	LOADBALANCERCONFIGTYPE_DYNAMIC  LoadBalancerConfigType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LoadBalancerConfigType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC",
		"DYNAMIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LoadBalancerConfigType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC",
		"DYNAMIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LoadBalancerConfigType) index(name string) LoadBalancerConfigType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC",
		"DYNAMIC",
	}
	for idx := range names {
		if names[idx] == name {
			return LoadBalancerConfigType(idx)
		}
	}
	return LOADBALANCERCONFIGTYPE_UNKNOWN
}

func (e *LoadBalancerConfigType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoadBalancerConfigType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoadBalancerConfigType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoadBalancerConfigType) Ref() *LoadBalancerConfigType {
	return &e
}

/*
Network configuration of the load balancer instance.
*/
type LoadBalancerNetworkConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessInterfaces []LoadBalancerAccessInterface `json:"accessInterfaces"`
	/*
	  Name of the client configuration for the load balancer.
	*/
	Client *string `json:"client,omitempty"`
	/*

	 */
	VmNetworkConfigItemDiscriminator_ *string `json:"$vmNetworkConfigItemDiscriminator,omitempty"`

	VmNetworkConfig *OneOfLoadBalancerNetworkConfigVmNetworkConfig `json:"vmNetworkConfig,omitempty"`
}

func (p *LoadBalancerNetworkConfig) MarshalJSON() ([]byte, error) {
	type LoadBalancerNetworkConfigProxy LoadBalancerNetworkConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LoadBalancerNetworkConfigProxy
		AccessInterfaces []LoadBalancerAccessInterface `json:"accessInterfaces,omitempty"`
	}{
		LoadBalancerNetworkConfigProxy: (*LoadBalancerNetworkConfigProxy)(p),
		AccessInterfaces:               p.AccessInterfaces,
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

func (p *LoadBalancerNetworkConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancerNetworkConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancerNetworkConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessInterfaces")
	delete(allFields, "client")
	delete(allFields, "$vmNetworkConfigItemDiscriminator")
	delete(allFields, "vmNetworkConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancerNetworkConfig() *LoadBalancerNetworkConfig {
	p := new(LoadBalancerNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *LoadBalancerNetworkConfig) GetVmNetworkConfig() interface{} {
	if nil == p.VmNetworkConfig {
		return nil
	}
	return p.VmNetworkConfig.GetValue()
}

func (p *LoadBalancerNetworkConfig) SetVmNetworkConfig(v interface{}) error {
	if nil == p.VmNetworkConfig {
		p.VmNetworkConfig = NewOneOfLoadBalancerNetworkConfigVmNetworkConfig()
	}
	e := p.VmNetworkConfig.SetValue(v)
	if nil == e {
		if nil == p.VmNetworkConfigItemDiscriminator_ {
			p.VmNetworkConfigItemDiscriminator_ = new(string)
		}
		*p.VmNetworkConfigItemDiscriminator_ = *p.VmNetworkConfig.Discriminator
	}
	return e
}

/*
load balancer vm resource configuration.
*/
type LoadBalancerResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of VCPUs to be allocated to the vm.
	*/
	Cpu *int64 `json:"cpu,omitempty"`
	/*
	  Memory requirements of VM.
	*/
	DiskMib *int64 `json:"diskMib,omitempty"`
	/*
	  Gold image name/UUID.
	*/
	GoldImageRef *string `json:"goldImageRef,omitempty"`
	/*
	  OS version of the gold image.
	*/
	GoldImageVersion *string `json:"goldImageVersion,omitempty"`
	/*
	  Labels for this resource configuration. Its a key value pair of strings.
	*/
	Labels []import1.KVStringPair `json:"labels,omitempty"`
	/*
	  Memory requirements of VM.
	*/
	MemoryMib *int64 `json:"memoryMib,omitempty"`
}

func (p *LoadBalancerResourceConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LoadBalancerResourceConfig

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

func (p *LoadBalancerResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoadBalancerResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoadBalancerResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpu")
	delete(allFields, "diskMib")
	delete(allFields, "goldImageRef")
	delete(allFields, "goldImageVersion")
	delete(allFields, "labels")
	delete(allFields, "memoryMib")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoadBalancerResourceConfig() *LoadBalancerResourceConfig {
	p := new(LoadBalancerResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of load balancer such as VM, hardware load balancer, container, etc.
*/
type LoadBalancerType int

const (
	LOADBALANCERTYPE_UNKNOWN           LoadBalancerType = 0
	LOADBALANCERTYPE_REDACTED          LoadBalancerType = 1
	LOADBALANCERTYPE_VM                LoadBalancerType = 2
	LOADBALANCERTYPE_CONTAINER_MASTERS LoadBalancerType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LoadBalancerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CONTAINER_MASTERS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LoadBalancerType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CONTAINER_MASTERS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LoadBalancerType) index(name string) LoadBalancerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CONTAINER_MASTERS",
	}
	for idx := range names {
		if names[idx] == name {
			return LoadBalancerType(idx)
		}
	}
	return LOADBALANCERTYPE_UNKNOWN
}

func (e *LoadBalancerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoadBalancerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoadBalancerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoadBalancerType) Ref() *LoadBalancerType {
	return &e
}

/*
Configuration of the default volume to be attached to a node for logging purpose.
*/
type LoggingVolumeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Mount path of the log volume.
	*/
	MountPath *string `json:"mountPath"`
	/*
	  Size of the log volume in MB.
	*/
	SizeMb *int `json:"sizeMb"`
	/*
	  Name of the storage config to be used for configuring the log volume.
	*/
	StorageConfigName *string `json:"storageConfigName"`
}

func (p *LoggingVolumeConfig) MarshalJSON() ([]byte, error) {
	type LoggingVolumeConfigProxy LoggingVolumeConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*LoggingVolumeConfigProxy
		MountPath         *string `json:"mountPath,omitempty"`
		SizeMb            *int    `json:"sizeMb,omitempty"`
		StorageConfigName *string `json:"storageConfigName,omitempty"`
	}{
		LoggingVolumeConfigProxy: (*LoggingVolumeConfigProxy)(p),
		MountPath:                p.MountPath,
		SizeMb:                   p.SizeMb,
		StorageConfigName:        p.StorageConfigName,
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

func (p *LoggingVolumeConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LoggingVolumeConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = LoggingVolumeConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "mountPath")
	delete(allFields, "sizeMb")
	delete(allFields, "storageConfigName")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewLoggingVolumeConfig() *LoggingVolumeConfig {
	p := new(LoggingVolumeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoggingVolumeConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Defines the type of NIC used for the network configuration.
*/
type NicType int

const (
	NICTYPE_UNKNOWN  NicType = 0
	NICTYPE_REDACTED NicType = 1
	NICTYPE_NORMAL   NicType = 2
	NICTYPE_DIRECT   NicType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"DIRECT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NicType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"DIRECT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NicType) index(name string) NicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL",
		"DIRECT",
	}
	for idx := range names {
		if names[idx] == name {
			return NicType(idx)
		}
	}
	return NICTYPE_UNKNOWN
}

func (e *NicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicType) Ref() *NicType {
	return &e
}

type NutanixStorageClass struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Nutanix Cluster Reference.
	*/
	ClusterRef *string `json:"clusterRef"`
	/*
	  File system of the storage volume.
	*/
	FileSystem *string `json:"fileSystem,omitempty"`

	IsFlashMode *bool `json:"isFlashMode,omitempty"`
	/*
	  Name of the storage container to be used.
	*/
	StorageContainer *string `json:"storageContainer,omitempty"`
}

func (p *NutanixStorageClass) MarshalJSON() ([]byte, error) {
	type NutanixStorageClassProxy NutanixStorageClass

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*NutanixStorageClassProxy
		ClusterRef *string `json:"clusterRef,omitempty"`
	}{
		NutanixStorageClassProxy: (*NutanixStorageClassProxy)(p),
		ClusterRef:               p.ClusterRef,
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

func (p *NutanixStorageClass) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NutanixStorageClass
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = NutanixStorageClass(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterRef")
	delete(allFields, "fileSystem")
	delete(allFields, "isFlashMode")
	delete(allFields, "storageContainer")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNutanixStorageClass() *NutanixStorageClass {
	p := new(NutanixStorageClass)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.NutanixStorageClass"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the owner of the task.
*/
type OwnerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the task owner.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Username of the task owner.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *OwnerReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias OwnerReference

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

func (p *OwnerReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias OwnerReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = OwnerReference(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PostAppResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AppUUID *string `json:"appUUID,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`
}

func (p *PostAppResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PostAppResponse

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

func (p *PostAppResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PostAppResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PostAppResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "appUUID")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPostAppResponse() *PostAppResponse {
	p := new(PostAppResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.PostAppResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PostServiceResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ServiceUUID *string `json:"serviceUUID,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`
}

func (p *PostServiceResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PostServiceResponse

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

func (p *PostServiceResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PostServiceResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = PostServiceResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "serviceUUID")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewPostServiceResponse() *PostServiceResponse {
	p := new(PostServiceResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.PostServiceResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type RegistryBundle struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  UUID of a cluster, optional. If not specified, the default controller_msp will be used
	*/
	ClusterUuid *string `json:"clusterUuid,omitempty"`

	Filepath *string `json:"filepath"`
}

func (p *RegistryBundle) MarshalJSON() ([]byte, error) {
	type RegistryBundleProxy RegistryBundle

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RegistryBundleProxy
		Filepath *string `json:"filepath,omitempty"`
	}{
		RegistryBundleProxy: (*RegistryBundleProxy)(p),
		Filepath:            p.Filepath,
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

func (p *RegistryBundle) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RegistryBundle
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = RegistryBundle(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterUuid")
	delete(allFields, "filepath")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewRegistryBundle() *RegistryBundle {
	p := new(RegistryBundle)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.RegistryBundle"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Service struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	IsInactive *bool `json:"isInactive,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *Service) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Service

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

func (p *Service) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Service
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = Service(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "isInactive")
	delete(allFields, "lastUpdatedTimestamp")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewService() *Service {
	p := new(Service)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Service"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ServiceProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ClusterUuid *string `json:"clusterUuid,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	IsInactive *bool `json:"isInactive,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (p *ServiceProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ServiceProjection

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

func (p *ServiceProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ServiceProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ServiceProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "action")
	delete(allFields, "apptype")
	delete(allFields, "clusterUuid")
	delete(allFields, "createdTimestamp")
	delete(allFields, "isInactive")
	delete(allFields, "lastUpdatedTimestamp")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "uuid")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewServiceProjection() *ServiceProjection {
	p := new(ServiceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ServiceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the service.
*/
type Status int

const (
	STATUS_UP      Status = 0
	STATUS_DOWN    Status = 1
	STATUS_UNKNOWN Status = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Status) name(index int) string {
	names := [...]string{
		"UP",
		"DOWN",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Status) GetName() string {
	index := int(e)
	names := [...]string{
		"UP",
		"DOWN",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Status) index(name string) Status {
	names := [...]string{
		"UP",
		"DOWN",
		"$UNKNOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return Status(idx)
		}
	}
	return STATUS_UNKNOWN
}

func (e *Status) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Status:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Status) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Status) Ref() *Status {
	return &e
}

/*
Reference to the parent task associated with the current task.
*/
type TaskReferenceInternal struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of the task.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The URL at which the entity described by the link can be accessed.
	*/
	Href *string `json:"href,omitempty"`
	/*
	  A name that identifies the relationship of the link to the object that is returned by the URL.  The unique value of "self" identifies the URL for the object.
	*/
	Rel *string `json:"rel,omitempty"`
}

func (p *TaskReferenceInternal) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TaskReferenceInternal

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

func (p *TaskReferenceInternal) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TaskReferenceInternal
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TaskReferenceInternal(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "href")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTaskReferenceInternal() *TaskReferenceInternal {
	p := new(TaskReferenceInternal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskReferenceInternal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the task.
*/
type TaskStatus int

const (
	TASKSTATUS_UNKNOWN   TaskStatus = 0
	TASKSTATUS_REDACTED  TaskStatus = 1
	TASKSTATUS_QUEUED    TaskStatus = 2
	TASKSTATUS_RUNNING   TaskStatus = 3
	TASKSTATUS_CANCELING TaskStatus = 4
	TASKSTATUS_SUCCEEDED TaskStatus = 5
	TASKSTATUS_FAILED    TaskStatus = 6
	TASKSTATUS_CANCELED  TaskStatus = 7
	TASKSTATUS_SKIPPED   TaskStatus = 8
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *TaskStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SKIPPED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e TaskStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SKIPPED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *TaskStatus) index(name string) TaskStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"QUEUED",
		"RUNNING",
		"CANCELING",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
		"SKIPPED",
	}
	for idx := range names {
		if names[idx] == name {
			return TaskStatus(idx)
		}
	}
	return TASKSTATUS_UNKNOWN
}

func (e *TaskStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for TaskStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *TaskStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e TaskStatus) Ref() *TaskStatus {
	return &e
}

/*
A single step in the task.
*/
type TaskStep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Message describing the completed steps for the task.
	*/
	Name *string `json:"name,omitempty"`
}

func (p *TaskStep) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TaskStep

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

func (p *TaskStep) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TaskStep
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TaskStep(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTaskStep() *TaskStep {
	p := new(TaskStep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskStep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TaskV2 struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of globally unique identifiers for clusters associated with the task or any of its subtasks.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was completed.
	*/
	CompletedTime *string `json:"completedTime,omitempty"`
	/*
	  Additional details on the task to aid the user with further actions post completion of the task.
	*/
	CompletionDetails []import1.KVPair `json:"completionDetails,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was created.
	*/
	CreatedTime *string `json:"createdTime,omitempty"`
	/*
	  Reference to entities associated with the task.
	*/
	EntitiesAffected []EntityReference `json:"entitiesAffected,omitempty"`
	/*
	  Error details explaining a task failure. These would be populated only in the case of task failures.
	*/
	ErrorMessages []import3.AppMessage `json:"errorMessages,omitempty"`
	/*
	  A globally unique identifier of a task.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Signifies if the task can be cancelled.
	*/
	IsCancelable *bool `json:"isCancelable,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was last updated.
	*/
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`
	/*
	  Provides an error message in the absence of a well-defined error message for the tasks created through legacy APIs.
	*/
	LegacyErrorMessage *string `json:"legacyErrorMessage,omitempty"`
	/*
	  A globally unique identifier of MSP task.
	*/
	MspUuid *string `json:"mspUuid,omitempty"`

	Name *string `json:"name,omitempty"`
	/*
	  The operation name being tracked by the task.
	*/
	Operation *string `json:"operation,omitempty"`
	/*
	  Description of the operation being tracked by the task.
	*/
	OperationDescription *string `json:"operationDescription,omitempty"`

	OwnedBy *OwnerReference `json:"ownedBy,omitempty"`

	ParentTask *TaskReferenceInternal `json:"parentTask,omitempty"`
	/*
	  Task progress expressed as a percentage.
	*/
	ProgressPercentage *int64 `json:"progressPercentage,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was started.
	*/
	StartedTime *string `json:"startedTime,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`
	/*
	  List of steps completed as part of the task.
	*/
	SubSteps []TaskStep `json:"subSteps,omitempty"`
	/*
	  Reference to tasks spawned as children of the current task.
	*/
	SubTasks []TaskReferenceInternal `json:"subTasks,omitempty"`

	Type *string `json:"type,omitempty"`
	/*
	  Warning messages to alert the user of issues which did not directly cause task failure. These can be populated for any task.
	*/
	Warnings []import3.AppMessage `json:"warnings,omitempty"`
}

func (p *TaskV2) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TaskV2

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

func (p *TaskV2) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TaskV2
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TaskV2(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "completedTime")
	delete(allFields, "completionDetails")
	delete(allFields, "createdTime")
	delete(allFields, "entitiesAffected")
	delete(allFields, "errorMessages")
	delete(allFields, "extId")
	delete(allFields, "isCancelable")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyErrorMessage")
	delete(allFields, "mspUuid")
	delete(allFields, "name")
	delete(allFields, "operation")
	delete(allFields, "operationDescription")
	delete(allFields, "ownedBy")
	delete(allFields, "parentTask")
	delete(allFields, "progressPercentage")
	delete(allFields, "startedTime")
	delete(allFields, "status")
	delete(allFields, "subSteps")
	delete(allFields, "subTasks")
	delete(allFields, "type")
	delete(allFields, "warnings")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTaskV2() *TaskV2 {
	p := new(TaskV2)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskV2"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TaskV2Projection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of globally unique identifiers for clusters associated with the task or any of its subtasks.
	*/
	ClusterExtIds []string `json:"clusterExtIds,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was completed.
	*/
	CompletedTime *string `json:"completedTime,omitempty"`
	/*
	  Additional details on the task to aid the user with further actions post completion of the task.
	*/
	CompletionDetails []import1.KVPair `json:"completionDetails,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was created.
	*/
	CreatedTime *string `json:"createdTime,omitempty"`
	/*
	  Reference to entities associated with the task.
	*/
	EntitiesAffected []EntityReference `json:"entitiesAffected,omitempty"`
	/*
	  Error details explaining a task failure. These would be populated only in the case of task failures.
	*/
	ErrorMessages []import3.AppMessage `json:"errorMessages,omitempty"`
	/*
	  A globally unique identifier of a task.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Signifies if the task can be cancelled.
	*/
	IsCancelable *bool `json:"isCancelable,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was last updated.
	*/
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`
	/*
	  Provides an error message in the absence of a well-defined error message for the tasks created through legacy APIs.
	*/
	LegacyErrorMessage *string `json:"legacyErrorMessage,omitempty"`
	/*
	  A globally unique identifier of MSP task.
	*/
	MspUuid *string `json:"mspUuid,omitempty"`

	Name *string `json:"name,omitempty"`
	/*
	  The operation name being tracked by the task.
	*/
	Operation *string `json:"operation,omitempty"`
	/*
	  Description of the operation being tracked by the task.
	*/
	OperationDescription *string `json:"operationDescription,omitempty"`

	OwnedBy *OwnerReference `json:"ownedBy,omitempty"`

	ParentTask *TaskReferenceInternal `json:"parentTask,omitempty"`
	/*
	  Task progress expressed as a percentage.
	*/
	ProgressPercentage *int64 `json:"progressPercentage,omitempty"`
	/*
	  UTC date and time in RFC-3339 format when the task was started.
	*/
	StartedTime *string `json:"startedTime,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`
	/*
	  List of steps completed as part of the task.
	*/
	SubSteps []TaskStep `json:"subSteps,omitempty"`
	/*
	  Reference to tasks spawned as children of the current task.
	*/
	SubTasks []TaskReferenceInternal `json:"subTasks,omitempty"`

	Type *string `json:"type,omitempty"`
	/*
	  Warning messages to alert the user of issues which did not directly cause task failure. These can be populated for any task.
	*/
	Warnings []import3.AppMessage `json:"warnings,omitempty"`
}

func (p *TaskV2Projection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TaskV2Projection

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

func (p *TaskV2Projection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TaskV2Projection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TaskV2Projection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtIds")
	delete(allFields, "completedTime")
	delete(allFields, "completionDetails")
	delete(allFields, "createdTime")
	delete(allFields, "entitiesAffected")
	delete(allFields, "errorMessages")
	delete(allFields, "extId")
	delete(allFields, "isCancelable")
	delete(allFields, "lastUpdatedTime")
	delete(allFields, "legacyErrorMessage")
	delete(allFields, "mspUuid")
	delete(allFields, "name")
	delete(allFields, "operation")
	delete(allFields, "operationDescription")
	delete(allFields, "ownedBy")
	delete(allFields, "parentTask")
	delete(allFields, "progressPercentage")
	delete(allFields, "startedTime")
	delete(allFields, "status")
	delete(allFields, "subSteps")
	delete(allFields, "subTasks")
	delete(allFields, "type")
	delete(allFields, "warnings")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTaskV2Projection() *TaskV2Projection {
	p := new(TaskV2Projection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskV2Projection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Location configuration of the VCenter server.
*/
type VcenterLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster name in Vcenter.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  Data center name in Vcenter.
	*/
	Datacenter *string `json:"datacenter,omitempty"`

	VcenterIp *import1.IPv4Address `json:"vcenterIp,omitempty"`
	/*
	  Vcenter Port.
	*/
	VcenterPort *int `json:"vcenterPort,omitempty"`
}

func (p *VcenterLocation) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VcenterLocation

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

func (p *VcenterLocation) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VcenterLocation
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VcenterLocation(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterName")
	delete(allFields, "datacenter")
	delete(allFields, "vcenterIp")
	delete(allFields, "vcenterPort")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVcenterLocation() *VcenterLocation {
	p := new(VcenterLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VcenterLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Network resources for Vm.
*/
type VmNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`
	/*
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`
}

func (p *VmNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type VmNetworkResourceConfigProxy VmNetworkResourceConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*VmNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		VmNetworkResourceConfigProxy: (*VmNetworkResourceConfigProxy)(p),
		Network:                      p.Network,
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

func (p *VmNetworkResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmNetworkResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmNetworkResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmNetworkResourceConfig() *VmNetworkResourceConfig {
	p := new(VmNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VmNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Resource configuration for the vm.
*/
type VmResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of VCPUs to be allocated to the vm.
	*/
	Cpu *int64 `json:"cpu,omitempty"`
	/*
	  Memory requirements of VM.
	*/
	DiskMib *int64 `json:"diskMib,omitempty"`
	/*
	  Memory requirements of VM.
	*/
	MemoryMib *int64 `json:"memoryMib,omitempty"`
}

func (p *VmResourceConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmResourceConfig

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

func (p *VmResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmResourceConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpu")
	delete(allFields, "diskMib")
	delete(allFields, "memoryMib")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmResourceConfig() *VmResourceConfig {
	p := new(VmResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VmResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfClusterResourceConfigResourceConfig struct {
	Discriminator *string            `json:"-"`
	ObjectType_   *string            `json:"-"`
	oneOfType1    *EsxResourceConfig `json:"-"`
	oneOfType0    *AhvResourceConfig `json:"-"`
}

func NewOneOfClusterResourceConfigResourceConfig() *OneOfClusterResourceConfigResourceConfig {
	p := new(OneOfClusterResourceConfigResourceConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterResourceConfigResourceConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterResourceConfigResourceConfig is nil"))
	}
	switch v.(type) {
	case EsxResourceConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(EsxResourceConfig)
		}
		*p.oneOfType1 = v.(EsxResourceConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case AhvResourceConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AhvResourceConfig)
		}
		*p.oneOfType0 = v.(AhvResourceConfig)
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

func (p *OneOfClusterResourceConfigResourceConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfClusterResourceConfigResourceConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(EsxResourceConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "lifecycle.v4.svcmgr.EsxResourceConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(EsxResourceConfig)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(AhvResourceConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "lifecycle.v4.svcmgr.AhvResourceConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AhvResourceConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterResourceConfigResourceConfig"))
}

func (p *OneOfClusterResourceConfigResourceConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfClusterResourceConfigResourceConfig")
}

type OneOfClientConfigConfig struct {
	Discriminator *string          `json:"-"`
	ObjectType_   *string          `json:"-"`
	oneOfType1    *EsxClientConfig `json:"-"`
	oneOfType0    *AhvClientConfig `json:"-"`
}

func NewOneOfClientConfigConfig() *OneOfClientConfigConfig {
	p := new(OneOfClientConfigConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClientConfigConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClientConfigConfig is nil"))
	}
	switch v.(type) {
	case EsxClientConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(EsxClientConfig)
		}
		*p.oneOfType1 = v.(EsxClientConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case AhvClientConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AhvClientConfig)
		}
		*p.oneOfType0 = v.(AhvClientConfig)
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

func (p *OneOfClientConfigConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfClientConfigConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(EsxClientConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "lifecycle.v4.svcmgr.EsxClientConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(EsxClientConfig)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(AhvClientConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "lifecycle.v4.svcmgr.AhvClientConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AhvClientConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClientConfigConfig"))
}

func (p *OneOfClientConfigConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfClientConfigConfig")
}

type OneOfLoadBalancerNetworkConfigVmNetworkConfig struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType1    *EsxNetworkResourceConfig `json:"-"`
	oneOfType0    *AhvNetworkResourceConfig `json:"-"`
}

func NewOneOfLoadBalancerNetworkConfigVmNetworkConfig() *OneOfLoadBalancerNetworkConfigVmNetworkConfig {
	p := new(OneOfLoadBalancerNetworkConfigVmNetworkConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfLoadBalancerNetworkConfigVmNetworkConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfLoadBalancerNetworkConfigVmNetworkConfig is nil"))
	}
	switch v.(type) {
	case EsxNetworkResourceConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(EsxNetworkResourceConfig)
		}
		*p.oneOfType1 = v.(EsxNetworkResourceConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case AhvNetworkResourceConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AhvNetworkResourceConfig)
		}
		*p.oneOfType0 = v.(AhvNetworkResourceConfig)
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

func (p *OneOfLoadBalancerNetworkConfigVmNetworkConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfLoadBalancerNetworkConfigVmNetworkConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(EsxNetworkResourceConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "lifecycle.v4.svcmgr.EsxNetworkResourceConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(EsxNetworkResourceConfig)
			}
			*p.oneOfType1 = *vOneOfType1
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType1.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType1.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(AhvNetworkResourceConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "lifecycle.v4.svcmgr.AhvNetworkResourceConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AhvNetworkResourceConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfLoadBalancerNetworkConfigVmNetworkConfig"))
}

func (p *OneOfLoadBalancerNetworkConfigVmNetworkConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfLoadBalancerNetworkConfigVmNetworkConfig")
}

type OneOfCustomValueItemValue struct {
	Discriminator *string  `json:"-"`
	ObjectType_   *string  `json:"-"`
	oneOfType0    []string `json:"-"`
	oneOfType1    *string  `json:"-"`
}

func NewOneOfCustomValueItemValue() *OneOfCustomValueItemValue {
	p := new(OneOfCustomValueItemValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCustomValueItemValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCustomValueItemValue is nil"))
	}
	switch v.(type) {
	case []string:
		p.oneOfType0 = v.([]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
	case string:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(string)
		}
		*p.oneOfType1 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCustomValueItemValue) GetValue() interface{} {
	if "List<String>" == *p.Discriminator {
		return p.oneOfType0
	}
	if "String" == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfCustomValueItemValue) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]string)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
		return nil
	}
	vOneOfType1 := new(string)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(string)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCustomValueItemValue"))
}

func (p *OneOfCustomValueItemValue) MarshalJSON() ([]byte, error) {
	if "List<String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfCustomValueItemValue")
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
