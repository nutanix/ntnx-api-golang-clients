/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/clustermgmt/v4/config"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/response"
	import2 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/vmm/v4/ahv/config"
	"time"
)

/*
This model would abstract away the common attributes as part of internal and external networks.
*/
type BaseNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGateway *import1.IPAddressOrFQDN `json:"defaultGateway"`
	/*
	  List of Node (VM) IP addresses used for Prism Central network setup.
	*/
	IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`
	/*
	  Range of IPs used for Prism Central network setup.
	*/
	IpRanges []import1.IpRange `json:"ipRanges,omitempty"`

	SubnetMask *import1.IPAddressOrFQDN `json:"subnetMask"`
}

func (p *BaseNetwork) MarshalJSON() ([]byte, error) {
	type BaseNetworkProxy BaseNetwork

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BaseNetworkProxy
		DefaultGateway *import1.IPAddressOrFQDN `json:"defaultGateway,omitempty"`
		SubnetMask     *import1.IPAddressOrFQDN `json:"subnetMask,omitempty"`
	}{
		BaseNetworkProxy: (*BaseNetworkProxy)(p),
		DefaultGateway:   p.DefaultGateway,
		SubnetMask:       p.SubnetMask,
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

func (p *BaseNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DefaultGateway != nil {
		p.DefaultGateway = known.DefaultGateway
	}
	if known.IpAddresses != nil {
		p.IpAddresses = known.IpAddresses
	}
	if known.IpRanges != nil {
		p.IpRanges = known.IpRanges
	}
	if known.SubnetMask != nil {
		p.SubnetMask = known.SubnetMask
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "defaultGateway")
	delete(allFields, "ipAddresses")
	delete(allFields, "ipRanges")
	delete(allFields, "subnetMask")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseNetwork() *BaseNetwork {
	p := new(BaseNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.BaseNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Bootstrap configuration details for the domain manager (Prism Central).
*/
type BootstrapConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of cloud-init commands required to bootstrap the domain manager (Prism Central) cluster on a startup.
	*/
	CloudInitConfig []import2.CloudInit `json:"cloudInitConfig,omitempty"`

	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty"`
}

func (p *BootstrapConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BootstrapConfig

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

func (p *BootstrapConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BootstrapConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBootstrapConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CloudInitConfig != nil {
		p.CloudInitConfig = known.CloudInitConfig
	}
	if known.EnvironmentInfo != nil {
		p.EnvironmentInfo = known.EnvironmentInfo
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cloudInitConfig")
	delete(allFields, "environmentInfo")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBootstrapConfig() *BootstrapConfig {
	p := new(BootstrapConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.BootstrapConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain manager (Prism Central) details.
*/
type DomainManager struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Config *DomainManagerClusterConfig `json:"config"`
	/*
	  The timestamp when the domain manager (Prism Central) was created.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The external identifier of the cluster hosting the domain manager (Prism Central) instance.
	*/
	HostingClusterExtId *string `json:"hostingClusterExtId,omitempty"`
	/*
	  Boolean value indicating if the domain manager (Prism Central) is registered with the hosting cluster, that is, Prism Element.
	*/
	IsRegisteredWithHostingCluster *bool `json:"isRegisteredWithHostingCluster,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`

	Network *DomainManagerNetwork `json:"network"`
	/*
	  Domain manager (Prism Central) nodes external identifier.
	*/
	NodeExtIds []string `json:"nodeExtIds,omitempty"`
	/*
	  This configuration enables Prism Central to be deployed in scale-out mode.
	*/
	ShouldEnableHighAvailability *bool `json:"shouldEnableHighAvailability,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *DomainManager) MarshalJSON() ([]byte, error) {
	type DomainManagerProxy DomainManager

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*DomainManagerProxy
		Config  *DomainManagerClusterConfig `json:"config,omitempty"`
		Network *DomainManagerNetwork       `json:"network,omitempty"`
	}{
		DomainManagerProxy: (*DomainManagerProxy)(p),
		Config:             p.Config,
		Network:            p.Network,
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

func (p *DomainManager) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainManager
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainManager()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HostingClusterExtId != nil {
		p.HostingClusterExtId = known.HostingClusterExtId
	}
	if known.IsRegisteredWithHostingCluster != nil {
		p.IsRegisteredWithHostingCluster = known.IsRegisteredWithHostingCluster
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.NodeExtIds != nil {
		p.NodeExtIds = known.NodeExtIds
	}
	if known.ShouldEnableHighAvailability != nil {
		p.ShouldEnableHighAvailability = known.ShouldEnableHighAvailability
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "config")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "hostingClusterExtId")
	delete(allFields, "isRegisteredWithHostingCluster")
	delete(allFields, "links")
	delete(allFields, "network")
	delete(allFields, "nodeExtIds")
	delete(allFields, "shouldEnableHighAvailability")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainManager() *DomainManager {
	p := new(DomainManager)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManager"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ShouldEnableHighAvailability = new(bool)
	*p.ShouldEnableHighAvailability = false

	return p
}

/*
Domain manager (Prism Central) cluster configuration details.
*/
type DomainManagerClusterConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BootstrapConfig *BootstrapConfig `json:"bootstrapConfig,omitempty"`

	BuildInfo *import4.BuildInfo `json:"buildInfo,omitempty"`
	/*
	  The credentials consist of a username and password for a particular user like admin. Users can pass the credentials of admin users currently which will be configured in the create domain manager operation.
	*/
	Credentials []import1.BasicAuth `json:"credentials,omitempty"`
	/*
	  Name of the domain manager (Prism Central).
	*/
	Name *string `json:"name,omitempty"`

	ResourceConfig *DomainManagerResourceConfig `json:"resourceConfig,omitempty"`
	/*
	  A boolean value indicating whether to enable lockdown mode for a cluster.
	*/
	ShouldEnableLockdownMode *bool `json:"shouldEnableLockdownMode,omitempty"`

	Size *Size `json:"size,omitempty"`
}

func (p *DomainManagerClusterConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainManagerClusterConfig

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

func (p *DomainManagerClusterConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainManagerClusterConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainManagerClusterConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BootstrapConfig != nil {
		p.BootstrapConfig = known.BootstrapConfig
	}
	if known.BuildInfo != nil {
		p.BuildInfo = known.BuildInfo
	}
	if known.Credentials != nil {
		p.Credentials = known.Credentials
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ResourceConfig != nil {
		p.ResourceConfig = known.ResourceConfig
	}
	if known.ShouldEnableLockdownMode != nil {
		p.ShouldEnableLockdownMode = known.ShouldEnableLockdownMode
	}
	if known.Size != nil {
		p.Size = known.Size
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bootstrapConfig")
	delete(allFields, "buildInfo")
	delete(allFields, "credentials")
	delete(allFields, "name")
	delete(allFields, "resourceConfig")
	delete(allFields, "shouldEnableLockdownMode")
	delete(allFields, "size")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainManagerClusterConfig() *DomainManagerClusterConfig {
	p := new(DomainManagerClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Domain manager (Prism Central) network configuration details.
*/
type DomainManagerNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Capability *NetworkCapability `json:"capability,omitempty"`

	ExternalAddress *import1.IPAddress `json:"externalAddress,omitempty"`
	/*
	  This configuration is used to manage Prism Central.
	*/
	ExternalNetworks []ExternalNetwork `json:"externalNetworks,omitempty"`
	/*
	  Cluster fully qualified domain name. This is part of payload for cluster update operation only.
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/*
	  List of HTTP Proxy server configuration needed to access a cluster which is hosted behind a HTTP Proxy to not reveal its identity.
	*/
	HttpProxyConfig []import4.HttpProxyConfig `json:"httpProxyConfig,omitempty"`
	/*
	  Targets HTTP traffic to which is exempted from going through the configured HTTP Proxy.
	*/
	HttpProxyWhiteListConfig []import4.HttpProxyWhiteListConfig `json:"httpProxyWhiteListConfig,omitempty"`
	/*
	  This configuration is used to internally manage Prism Central network.
	*/
	InternalNetworks []BaseNetwork `json:"internalNetworks,omitempty"`
	/*
	  List of name servers on a cluster. This is a part of payload for both clusters create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NameServers []import1.IPAddressOrFQDN `json:"nameServers,omitempty"`
	/*
	  List of NTP servers on a cluster. This is a part of payload for both cluster create and update operations. Currently, only IPv4 address and FQDN (fully qualified domain name) values are supported for the create operation.
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func (p *DomainManagerNetwork) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainManagerNetwork

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

func (p *DomainManagerNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainManagerNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainManagerNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Capability != nil {
		p.Capability = known.Capability
	}
	if known.ExternalAddress != nil {
		p.ExternalAddress = known.ExternalAddress
	}
	if known.ExternalNetworks != nil {
		p.ExternalNetworks = known.ExternalNetworks
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
	if known.InternalNetworks != nil {
		p.InternalNetworks = known.InternalNetworks
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
	delete(allFields, "capability")
	delete(allFields, "externalAddress")
	delete(allFields, "externalNetworks")
	delete(allFields, "fqdn")
	delete(allFields, "httpProxyConfig")
	delete(allFields, "httpProxyWhiteListConfig")
	delete(allFields, "internalNetworks")
	delete(allFields, "nameServers")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainManagerNetwork() *DomainManagerNetwork {
	p := new(DomainManagerNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This configuration is used to provide the resource-related details like container external identifiers, number of VCPUs, memory size, data disk size of the domain manager (Prism Central). In the case of a multi-node setup, the sum of resources like number of VCPUs, memory size and data disk size are provided.
*/
type DomainManagerResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the container that will be used to create the domain manager (Prism Central) cluster.
	*/
	ContainerExtIds []string `json:"containerExtIds,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central data disk size allocation at a cluster level.
	*/
	DataDiskSizeBytes *int64 `json:"dataDiskSizeBytes,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central RAM allocation at the cluster level.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  This property is used for readOnly purposes to display Prism Central number of VCPUs allocation.
	*/
	NumVcpus *int `json:"numVcpus,omitempty"`
}

func (p *DomainManagerResourceConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DomainManagerResourceConfig

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

func (p *DomainManagerResourceConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DomainManagerResourceConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDomainManagerResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ContainerExtIds != nil {
		p.ContainerExtIds = known.ContainerExtIds
	}
	if known.DataDiskSizeBytes != nil {
		p.DataDiskSizeBytes = known.DataDiskSizeBytes
	}
	if known.MemorySizeBytes != nil {
		p.MemorySizeBytes = known.MemorySizeBytes
	}
	if known.NumVcpus != nil {
		p.NumVcpus = known.NumVcpus
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "containerExtIds")
	delete(allFields, "dataDiskSizeBytes")
	delete(allFields, "memorySizeBytes")
	delete(allFields, "numVcpus")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDomainManagerResourceConfig() *DomainManagerResourceConfig {
	p := new(DomainManagerResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.DomainManagerResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An object denoting the environment information of the PC. It contains the following fields:<br>
- type: Enums denoting the environment type of the PC.<br>
- providerType: Enums denoting the provider of the cloud PC.<br>
- instanceObj: Enums denoting the instance type of the cloud PC.<br>
*/
type EnvironmentInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ProviderType *ProviderType `json:"providerType,omitempty"`

	ProvisioningType *ProvisioningType `json:"provisioningType,omitempty"`

	Type *EnvironmentType `json:"type,omitempty"`
}

func (p *EnvironmentInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EnvironmentInfo

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

func (p *EnvironmentInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EnvironmentInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEnvironmentInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProviderType != nil {
		p.ProviderType = known.ProviderType
	}
	if known.ProvisioningType != nil {
		p.ProvisioningType = known.ProvisioningType
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "providerType")
	delete(allFields, "provisioningType")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEnvironmentInfo() *EnvironmentInfo {
	p := new(EnvironmentInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.EnvironmentInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Enums denoting the environment type of the PC, that is, on-prem PC or cloud PC.<br>
Following are the supported entity types:
- ONPREM
- NTNX_CLOUD
*/
type EnvironmentType int

const (
	ENVIRONMENTTYPE_UNKNOWN    EnvironmentType = 0
	ENVIRONMENTTYPE_REDACTED   EnvironmentType = 1
	ENVIRONMENTTYPE_ONPREM     EnvironmentType = 2
	ENVIRONMENTTYPE_NTNX_CLOUD EnvironmentType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *EnvironmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e EnvironmentType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *EnvironmentType) index(name string) EnvironmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ONPREM",
		"NTNX_CLOUD",
	}
	for idx := range names {
		if names[idx] == name {
			return EnvironmentType(idx)
		}
	}
	return ENVIRONMENTTYPE_UNKNOWN
}

func (e *EnvironmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for EnvironmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *EnvironmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e EnvironmentType) Ref() *EnvironmentType {
	return &e
}

/*
This configuration is used to manage Prism Central.
*/
type ExternalNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DefaultGateway *import1.IPAddressOrFQDN `json:"defaultGateway"`
	/*
	  List of Node (VM) IP addresses used for Prism Central network setup.
	*/
	IpAddresses []import1.IPAddress `json:"ipAddresses,omitempty"`
	/*
	  Range of IPs used for Prism Central network setup.
	*/
	IpRanges []import1.IpRange `json:"ipRanges,omitempty"`
	/*
	  The network external identifier to which Domain Manager (Prism Central) is to be deployed or is already configured.
	*/
	NetworkExtId *string `json:"networkExtId"`

	SubnetMask *import1.IPAddressOrFQDN `json:"subnetMask"`
}

func (p *ExternalNetwork) MarshalJSON() ([]byte, error) {
	type ExternalNetworkProxy ExternalNetwork

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ExternalNetworkProxy
		DefaultGateway *import1.IPAddressOrFQDN `json:"defaultGateway,omitempty"`
		NetworkExtId   *string                  `json:"networkExtId,omitempty"`
		SubnetMask     *import1.IPAddressOrFQDN `json:"subnetMask,omitempty"`
	}{
		ExternalNetworkProxy: (*ExternalNetworkProxy)(p),
		DefaultGateway:       p.DefaultGateway,
		NetworkExtId:         p.NetworkExtId,
		SubnetMask:           p.SubnetMask,
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

func (p *ExternalNetwork) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ExternalNetwork
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewExternalNetwork()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DefaultGateway != nil {
		p.DefaultGateway = known.DefaultGateway
	}
	if known.IpAddresses != nil {
		p.IpAddresses = known.IpAddresses
	}
	if known.IpRanges != nil {
		p.IpRanges = known.IpRanges
	}
	if known.NetworkExtId != nil {
		p.NetworkExtId = known.NetworkExtId
	}
	if known.SubnetMask != nil {
		p.SubnetMask = known.SubnetMask
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "defaultGateway")
	delete(allFields, "ipAddresses")
	delete(allFields, "ipRanges")
	delete(allFields, "networkExtId")
	delete(allFields, "subnetMask")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewExternalNetwork() *ExternalNetwork {
	p := new(ExternalNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.ExternalNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This property represents network capability of a domain manager which consists of - Ipv4 only, dual stack and ipv6 only networks.
*/
type NetworkCapability int

const (
	NETWORKCAPABILITY_UNKNOWN    NetworkCapability = 0
	NETWORKCAPABILITY_REDACTED   NetworkCapability = 1
	NETWORKCAPABILITY_IPV4       NetworkCapability = 2
	NETWORKCAPABILITY_DUAL_STACK NetworkCapability = 3
	NETWORKCAPABILITY_IPV6       NetworkCapability = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NetworkCapability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4",
		"DUAL_STACK",
		"IPV6",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NetworkCapability) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4",
		"DUAL_STACK",
		"IPV6",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NetworkCapability) index(name string) NetworkCapability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4",
		"DUAL_STACK",
		"IPV6",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkCapability(idx)
		}
	}
	return NETWORKCAPABILITY_UNKNOWN
}

func (e *NetworkCapability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkCapability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkCapability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkCapability) Ref() *NetworkCapability {
	return &e
}

/*
Enums denoting the provider of the cloud, in case of environment type a cloud PC.<br>
The service currently supports the following providers:
- NTNX
- AZURE
- AWS
- GCP
- VSPHERE
*/
type ProviderType int

const (
	PROVIDERTYPE_UNKNOWN  ProviderType = 0
	PROVIDERTYPE_REDACTED ProviderType = 1
	PROVIDERTYPE_NTNX     ProviderType = 2
	PROVIDERTYPE_AZURE    ProviderType = 3
	PROVIDERTYPE_AWS      ProviderType = 4
	PROVIDERTYPE_GCP      ProviderType = 5
	PROVIDERTYPE_VSPHERE  ProviderType = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProviderType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProviderType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProviderType) index(name string) ProviderType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"AZURE",
		"AWS",
		"GCP",
		"VSPHERE",
	}
	for idx := range names {
		if names[idx] == name {
			return ProviderType(idx)
		}
	}
	return PROVIDERTYPE_UNKNOWN
}

func (e *ProviderType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProviderType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProviderType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProviderType) Ref() *ProviderType {
	return &e
}

/*
Enums denoting the instance type of the cloud PC. It indicates whether the PC is created on bare-metal
or on a cloud-provisioned VM. Hence, it supports two possible values:
- NTNX
- NATIVE
*/
type ProvisioningType int

const (
	PROVISIONINGTYPE_UNKNOWN  ProvisioningType = 0
	PROVISIONINGTYPE_REDACTED ProvisioningType = 1
	PROVISIONINGTYPE_NTNX     ProvisioningType = 2
	PROVISIONINGTYPE_NATIVE   ProvisioningType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ProvisioningType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ProvisioningType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ProvisioningType) index(name string) ProvisioningType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NTNX",
		"NATIVE",
	}
	for idx := range names {
		if names[idx] == name {
			return ProvisioningType(idx)
		}
	}
	return PROVISIONINGTYPE_UNKNOWN
}

func (e *ProvisioningType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ProvisioningType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ProvisioningType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ProvisioningType) Ref() *ProvisioningType {
	return &e
}

/*
Domain manager (Prism Central) size is an enumeration of starter, small, large, or extra large starter values.
*/
type Size int

const (
	SIZE_UNKNOWN    Size = 0
	SIZE_REDACTED   Size = 1
	SIZE_STARTER    Size = 2
	SIZE_SMALL      Size = 3
	SIZE_LARGE      Size = 4
	SIZE_EXTRALARGE Size = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Size) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e Size) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *Size) index(name string) Size {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STARTER",
		"SMALL",
		"LARGE",
		"EXTRALARGE",
	}
	for idx := range names {
		if names[idx] == name {
			return Size(idx)
		}
	}
	return SIZE_UNKNOWN
}

func (e *Size) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Size:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Size) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Size) Ref() *Size {
	return &e
}

/*
A reference to a task tracking an asynchronous operation. The status of the task can be queried by making a GET request to the task URI provided in the metadata section of the API response.
*/
type TaskReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier for a task.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func (p *TaskReference) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TaskReference

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

func (p *TaskReference) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TaskReference
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTaskReference()

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

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTaskReference() *TaskReference {
	p := new(TaskReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.config.TaskReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
