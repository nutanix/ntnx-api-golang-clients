/*
 * Generated file models/lifecycle/v4/svcmgr/svcmgr_model.go.
 *
 * Product version: 4.2.1
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
Details of a REST API call action.
*/
type APICallDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Endpoint path to invoke.
	*/
	Endpoint *string `json:"endpoint"`
	/*
	  Host of the API endpoint.
	*/
	Host *string `json:"host"`
	/*
	  Port of the API endpoint.
	*/
	Port *int64 `json:"port"`
	/*
	  Request body sent with the API call.
	*/
	RequestBody *string `json:"requestBody,omitempty"`
}

func (p *APICallDetails) MarshalJSON() ([]byte, error) {
	type APICallDetailsProxy APICallDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*APICallDetailsProxy
		Endpoint *string `json:"endpoint,omitempty"`
		Host     *string `json:"host,omitempty"`
		Port     *int64  `json:"port,omitempty"`
	}{
		APICallDetailsProxy: (*APICallDetailsProxy)(p),
		Endpoint:            p.Endpoint,
		Host:                p.Host,
		Port:                p.Port,
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

func (p *APICallDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias APICallDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAPICallDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Endpoint != nil {
		p.Endpoint = known.Endpoint
	}
	if known.Host != nil {
		p.Host = known.Host
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.RequestBody != nil {
		p.RequestBody = known.RequestBody
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endpoint")
	delete(allFields, "host")
	delete(allFields, "port")
	delete(allFields, "requestBody")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAPICallDetails() *APICallDetails {
	p := new(APICallDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.APICallDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

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
	*p = *NewAhvClientConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CaChain != nil {
		p.CaChain = known.CaChain
	}
	if known.Certificate != nil {
		p.Certificate = known.Certificate
	}
	if known.Key != nil {
		p.Key = known.Key
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.PrismIp != nil {
		p.PrismIp = known.PrismIp
	}
	if known.PrismPort != nil {
		p.PrismPort = known.PrismPort
	}
	if known.ServiceName != nil {
		p.ServiceName = known.ServiceName
	}
	if known.Username != nil {
		p.Username = known.Username
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAhvClientConfig() *AhvClientConfig {
	p := new(AhvClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	/*
	  Define multiple NICs for AHV Platform.
	*/
	Nics []AhvNic `json:"nics,omitempty"`
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
	*p = *NewAhvNetworkResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterId != nil {
		p.ClusterId = known.ClusterId
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.NicType != nil {
		p.NicType = known.NicType
	}
	if known.Nics != nil {
		p.Nics = known.Nics
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "network")
	delete(allFields, "nicType")
	delete(allFields, "nics")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAhvNetworkResourceConfig() *AhvNetworkResourceConfig {
	p := new(AhvNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AhvNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Network *string `json:"network,omitempty"`

	NicType *NicType `json:"nicType,omitempty"`

	NicUsage []NicUsage `json:"nicUsage,omitempty"`

	NumQueues *int64 `json:"numQueues,omitempty"`
}

func (p *AhvNic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AhvNic

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

func (p *AhvNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AhvNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAhvNic()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.NicType != nil {
		p.NicType = known.NicType
	}
	if known.NicUsage != nil {
		p.NicUsage = known.NicUsage
	}
	if known.NumQueues != nil {
		p.NumQueues = known.NumQueues
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "network")
	delete(allFields, "nicType")
	delete(allFields, "nicUsage")
	delete(allFields, "numQueues")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAhvNic() *AhvNic {
	p := new(AhvNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewAhvResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NetworkResource != nil {
		p.NetworkResource = known.NetworkResource
	}
	if known.VmResource != nil {
		p.VmResource = known.VmResource
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkResource")
	delete(allFields, "vmResource")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAhvResourceConfig() *AhvResourceConfig {
	p := new(AhvResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AhvResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents an Application Disaster Recovery (App DR) app protection policy.
*/
type AppProtectionPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the SMSP cluster to which the app protection policy is bound.
	*/
	ClusterExtId *string `json:"clusterExtId,omitempty"`
	/*
	  Object creation timestamp.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	HelmApplicationInfo *HelmApplicationInfo `json:"helmApplicationInfo,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Name of the app protection policy.
	*/
	Name *string `json:"name"`
	/*
	  Protection targets that extend coverage across multiple namespaces.
	*/
	ProtectionTargets []ProtectionTarget `json:"protectionTargets,omitempty"`
	/*
	  List of replication target configurations.
	*/
	ReplicationTargets []ReplicationTarget `json:"replicationTargets"`

	SourceCluster *ClusterTarget `json:"sourceCluster"`

	State *AppProtectionPolicyState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Last update timestamp.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
}

func (p *AppProtectionPolicy) MarshalJSON() ([]byte, error) {
	type AppProtectionPolicyProxy AppProtectionPolicy

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AppProtectionPolicyProxy
		Name               *string             `json:"name,omitempty"`
		ReplicationTargets []ReplicationTarget `json:"replicationTargets,omitempty"`
		SourceCluster      *ClusterTarget      `json:"sourceCluster,omitempty"`
	}{
		AppProtectionPolicyProxy: (*AppProtectionPolicyProxy)(p),
		Name:                     p.Name,
		ReplicationTargets:       p.ReplicationTargets,
		SourceCluster:            p.SourceCluster,
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

func (p *AppProtectionPolicy) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AppProtectionPolicy
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAppProtectionPolicy()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.HelmApplicationInfo != nil {
		p.HelmApplicationInfo = known.HelmApplicationInfo
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ProtectionTargets != nil {
		p.ProtectionTargets = known.ProtectionTargets
	}
	if known.ReplicationTargets != nil {
		p.ReplicationTargets = known.ReplicationTargets
	}
	if known.SourceCluster != nil {
		p.SourceCluster = known.SourceCluster
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpdatedTime != nil {
		p.UpdatedTime = known.UpdatedTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "helmApplicationInfo")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "protectionTargets")
	delete(allFields, "replicationTargets")
	delete(allFields, "sourceCluster")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "updatedTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAppProtectionPolicy() *AppProtectionPolicy {
	p := new(AppProtectionPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AppProtectionPolicy"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents the start parameters for an App Protection Policy.
*/
type AppProtectionPolicyStart struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  First time at which the scheduler should trigger.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func (p *AppProtectionPolicyStart) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias AppProtectionPolicyStart

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

func (p *AppProtectionPolicyStart) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AppProtectionPolicyStart
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAppProtectionPolicyStart()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "startTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAppProtectionPolicyStart() *AppProtectionPolicyStart {
	p := new(AppProtectionPolicyStart)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AppProtectionPolicyStart"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
State of the app protection policy.
*/
type AppProtectionPolicyState int

const (
	APPPROTECTIONPOLICYSTATE_UNKNOWN           AppProtectionPolicyState = 0
	APPPROTECTIONPOLICYSTATE_REDACTED          AppProtectionPolicyState = 1
	APPPROTECTIONPOLICYSTATE_RUNNING           AppProtectionPolicyState = 2
	APPPROTECTIONPOLICYSTATE_STOPPED           AppProtectionPolicyState = 3
	APPPROTECTIONPOLICYSTATE_LOCALLY_PROTECTED AppProtectionPolicyState = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *AppProtectionPolicyState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"STOPPED",
		"LOCALLY_PROTECTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e AppProtectionPolicyState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"STOPPED",
		"LOCALLY_PROTECTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *AppProtectionPolicyState) index(name string) AppProtectionPolicyState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"STOPPED",
		"LOCALLY_PROTECTED",
	}
	for idx := range names {
		if names[idx] == name {
			return AppProtectionPolicyState(idx)
		}
	}
	return APPPROTECTIONPOLICYSTATE_UNKNOWN
}

func (e *AppProtectionPolicyState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AppProtectionPolicyState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AppProtectionPolicyState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AppProtectionPolicyState) Ref() *AppProtectionPolicyState {
	return &e
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
	*p = *NewApplication()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.IsInactive != nil {
		p.IsInactive = known.IsInactive
	}
	if known.LastUpdatedTimestamp != nil {
		p.LastUpdatedTimestamp = known.LastUpdatedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.SubServices != nil {
		p.SubServices = known.SubServices
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApplication() *Application {
	p := new(Application)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Application"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewApplicationProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.IsInactive != nil {
		p.IsInactive = known.IsInactive
	}
	if known.LastUpdatedTimestamp != nil {
		p.LastUpdatedTimestamp = known.LastUpdatedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.SubServices != nil {
		p.SubServices = known.SubServices
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApplicationProjection() *ApplicationProjection {
	p := new(ApplicationProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ApplicationProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of the protected entity (application or service).
*/
type ApplicationType int

const (
	APPLICATIONTYPE_UNKNOWN     ApplicationType = 0
	APPLICATIONTYPE_REDACTED    ApplicationType = 1
	APPLICATIONTYPE_APPLICATION ApplicationType = 2
	APPLICATIONTYPE_SERVICE     ApplicationType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ApplicationType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICATION",
		"SERVICE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ApplicationType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICATION",
		"SERVICE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ApplicationType) index(name string) ApplicationType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"APPLICATION",
		"SERVICE",
	}
	for idx := range names {
		if names[idx] == name {
			return ApplicationType(idx)
		}
	}
	return APPLICATIONTYPE_UNKNOWN
}

func (e *ApplicationType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ApplicationType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ApplicationType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ApplicationType) Ref() *ApplicationType {
	return &e
}

/*
Configuration for an asynchronous replication target.
*/
type AsyncReplicationTargetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the replication-target cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  External identifier of the Prism Central managing the replication-target cluster.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`

	Schedule *CronSchedule `json:"schedule"`
	/*
	  Number of most-recent snapshots to retain for asynchronous replication.
	*/
	SnapshotRetentionCount *int `json:"snapshotRetentionCount,omitempty"`
}

func (p *AsyncReplicationTargetConfig) MarshalJSON() ([]byte, error) {
	type AsyncReplicationTargetConfigProxy AsyncReplicationTargetConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*AsyncReplicationTargetConfigProxy
		ClusterExtId       *string       `json:"clusterExtId,omitempty"`
		DomainManagerExtId *string       `json:"domainManagerExtId,omitempty"`
		Schedule           *CronSchedule `json:"schedule,omitempty"`
	}{
		AsyncReplicationTargetConfigProxy: (*AsyncReplicationTargetConfigProxy)(p),
		ClusterExtId:                      p.ClusterExtId,
		DomainManagerExtId:                p.DomainManagerExtId,
		Schedule:                          p.Schedule,
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

func (p *AsyncReplicationTargetConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias AsyncReplicationTargetConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewAsyncReplicationTargetConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}
	if known.Schedule != nil {
		p.Schedule = known.Schedule
	}
	if known.SnapshotRetentionCount != nil {
		p.SnapshotRetentionCount = known.SnapshotRetentionCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "domainManagerExtId")
	delete(allFields, "schedule")
	delete(allFields, "snapshotRetentionCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewAsyncReplicationTargetConfig() *AsyncReplicationTargetConfig {
	p := new(AsyncReplicationTargetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.AsyncReplicationTargetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewBaseClientConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Certificate != nil {
		p.Certificate = known.Certificate
	}
	if known.Key != nil {
		p.Key = known.Key
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.Username != nil {
		p.Username = known.Username
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "certificate")
	delete(allFields, "key")
	delete(allFields, "password")
	delete(allFields, "username")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseClientConfig() *BaseClientConfig {
	p := new(BaseClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.BaseClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Base configuration shared by all replication targets.
*/
type BaseReplicationTargetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the replication-target cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  External identifier of the Prism Central managing the replication-target cluster.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *BaseReplicationTargetConfig) MarshalJSON() ([]byte, error) {
	type BaseReplicationTargetConfigProxy BaseReplicationTargetConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*BaseReplicationTargetConfigProxy
		ClusterExtId       *string `json:"clusterExtId,omitempty"`
		DomainManagerExtId *string `json:"domainManagerExtId,omitempty"`
	}{
		BaseReplicationTargetConfigProxy: (*BaseReplicationTargetConfigProxy)(p),
		ClusterExtId:                     p.ClusterExtId,
		DomainManagerExtId:               p.DomainManagerExtId,
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

func (p *BaseReplicationTargetConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BaseReplicationTargetConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBaseReplicationTargetConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBaseReplicationTargetConfig() *BaseReplicationTargetConfig {
	p := new(BaseReplicationTargetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.BaseReplicationTargetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClientConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ConfigItemDiscriminator_ != nil {
		p.ConfigItemDiscriminator_ = known.ConfigItemDiscriminator_
	}
	if known.Config != nil {
		p.Config = known.Config
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$configItemDiscriminator")
	delete(allFields, "config")
	delete(allFields, "name")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClientConfig() *ClientConfig {
	p := new(ClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	if known.ControlPlaneConfig != nil {
		p.ControlPlaneConfig = known.ControlPlaneConfig
	}
	if known.Annotations != nil {
		p.Annotations = known.Annotations
	}
	if known.ClientConfigs != nil {
		p.ClientConfigs = known.ClientConfigs
	}
	if known.ClusterType != nil {
		p.ClusterType = known.ClusterType
	}
	if known.ControllerVersion != nil {
		p.ControllerVersion = known.ControllerVersion
	}
	if known.DeploymentType != nil {
		p.DeploymentType = known.DeploymentType
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.Domain != nil {
		p.Domain = known.Domain
	}
	if known.EnvoyIps != nil {
		p.EnvoyIps = known.EnvoyIps
	}
	if known.EtcdIps != nil {
		p.EtcdIps = known.EtcdIps
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ExternalMasterIp != nil {
		p.ExternalMasterIp = known.ExternalMasterIp
	}
	if known.IsLockedDown != nil {
		p.IsLockedDown = known.IsLockedDown
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.LoadBalancerConfigs != nil {
		p.LoadBalancerConfigs = known.LoadBalancerConfigs
	}
	if known.LoggingVolumeConfig != nil {
		p.LoggingVolumeConfig = known.LoggingVolumeConfig
	}
	if known.ManagementIps != nil {
		p.ManagementIps = known.ManagementIps
	}
	if known.MasterIps != nil {
		p.MasterIps = known.MasterIps
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NetworkConfigs != nil {
		p.NetworkConfigs = known.NetworkConfigs
	}
	if known.ResourceConfigs != nil {
		p.ResourceConfigs = known.ResourceConfigs
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.StorageClasses != nil {
		p.StorageClasses = known.StorageClasses
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Version != nil {
		p.Version = known.Version
	}
	if known.WorkerConfigs != nil {
		p.WorkerConfigs = known.WorkerConfigs
	}
	if known.WorkerIps != nil {
		p.WorkerIps = known.WorkerIps
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Cluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClusterManagementIps()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MasterVip != nil {
		p.MasterVip = known.MasterVip
	}
	if known.MspDnsIp != nil {
		p.MspDnsIp = known.MspDnsIp
	}
	if known.Network != nil {
		p.Network = known.Network
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "masterVip")
	delete(allFields, "mspDnsIp")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterManagementIps() *ClusterManagementIps {
	p := new(ClusterManagementIps)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterManagementIps"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClusterNetworkConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DnsServers != nil {
		p.DnsServers = known.DnsServers
	}
	if known.EsxLocation != nil {
		p.EsxLocation = known.EsxLocation
	}
	if known.GatewayIp != nil {
		p.GatewayIp = known.GatewayIp
	}
	if known.IpRanges != nil {
		p.IpRanges = known.IpRanges
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Netmask != nil {
		p.Netmask = known.Netmask
	}
	if known.Network != nil {
		p.Network = known.Network
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterNetworkConfig() *ClusterNetworkConfig {
	p := new(ClusterNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClusterResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Client != nil {
		p.Client = known.Client
	}
	if known.Labels != nil {
		p.Labels = known.Labels
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ResourceConfigItemDiscriminator_ != nil {
		p.ResourceConfigItemDiscriminator_ = known.ResourceConfigItemDiscriminator_
	}
	if known.ResourceConfig != nil {
		p.ResourceConfig = known.ResourceConfig
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterResourceConfig() *ClusterResourceConfig {
	p := new(ClusterResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClusterStorageClass()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.IsDefaultStorageClass != nil {
		p.IsDefaultStorageClass = known.IsDefaultStorageClass
	}
	if known.MountOptions != nil {
		p.MountOptions = known.MountOptions
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NutanixStorageClass != nil {
		p.NutanixStorageClass = known.NutanixStorageClass
	}
	if known.ReclaimPolicy != nil {
		p.ReclaimPolicy = known.ReclaimPolicy
	}
	if known.StorageType != nil {
		p.StorageType = known.StorageType
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterStorageClass() *ClusterStorageClass {
	p := new(ClusterStorageClass)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterStorageClass"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to an SMSP cluster and the Prism Central that manages it.
*/
type ClusterTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the SMSP cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  External identifier of the Prism Central that owns the SMSP cluster.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *ClusterTarget) MarshalJSON() ([]byte, error) {
	type ClusterTargetProxy ClusterTarget

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ClusterTargetProxy
		ClusterExtId       *string `json:"clusterExtId,omitempty"`
		DomainManagerExtId *string `json:"domainManagerExtId,omitempty"`
	}{
		ClusterTargetProxy: (*ClusterTargetProxy)(p),
		ClusterExtId:       p.ClusterExtId,
		DomainManagerExtId: p.DomainManagerExtId,
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

func (p *ClusterTarget) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ClusterTarget
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewClusterTarget()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterTarget() *ClusterTarget {
	p := new(ClusterTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewClusterUpgradeInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.CompatibleVersions != nil {
		p.CompatibleVersions = known.CompatibleVersions
	}
	if known.CurrentVersion != nil {
		p.CurrentVersion = known.CurrentVersion
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "compatibleVersions")
	delete(allFields, "currentVersion")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewClusterUpgradeInfo() *ClusterUpgradeInfo {
	p := new(ClusterUpgradeInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ClusterUpgradeInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewComponentDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ErrorMessage != nil {
		p.ErrorMessage = known.ErrorMessage
	}
	if known.FirstFailureTime != nil {
		p.FirstFailureTime = known.FirstFailureTime
	}
	if known.Messages != nil {
		p.Messages = known.Messages
	}
	if known.NumFailures != nil {
		p.NumFailures = known.NumFailures
	}
	if known.Timestamp != nil {
		p.Timestamp = known.Timestamp
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewComponentDetails() *ComponentDetails {
	p := new(ComponentDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ComponentDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NumInstances != nil {
		p.NumInstances = known.NumInstances
	}
	if known.ResourceConfigName != nil {
		p.ResourceConfigName = known.ResourceConfigName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "numInstances")
	delete(allFields, "resourceConfigName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewConfig() *Config {
	p := new(Config)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Config"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Cron schedule configuration.
*/
type CronSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Cron expression that defines the snapshot schedule.
	*/
	CronSchedule *string `json:"cronSchedule"`
	/*
	  First time at which the scheduler should trigger.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
	/*
	  Time zone used to evaluate the cron expression.
	*/
	Timezone *string `json:"timezone"`
}

func (p *CronSchedule) MarshalJSON() ([]byte, error) {
	type CronScheduleProxy CronSchedule

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CronScheduleProxy
		CronSchedule *string `json:"cronSchedule,omitempty"`
		Timezone     *string `json:"timezone,omitempty"`
	}{
		CronScheduleProxy: (*CronScheduleProxy)(p),
		CronSchedule:      p.CronSchedule,
		Timezone:          p.Timezone,
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

func (p *CronSchedule) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CronSchedule
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCronSchedule()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CronSchedule != nil {
		p.CronSchedule = known.CronSchedule
	}
	if known.StartTime != nil {
		p.StartTime = known.StartTime
	}
	if known.Timezone != nil {
		p.Timezone = known.Timezone
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cronSchedule")
	delete(allFields, "startTime")
	delete(allFields, "timezone")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCronSchedule() *CronSchedule {
	p := new(CronSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.CronSchedule"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewCustomValue()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Service != nil {
		p.Service = known.Service
	}
	if known.Values != nil {
		p.Values = known.Values
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "service")
	delete(allFields, "values")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCustomValue() *CustomValue {
	p := new(CustomValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.CustomValue"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewCustomValueItem()

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
	if known.ValueItemDiscriminator_ != nil {
		p.ValueItemDiscriminator_ = known.ValueItemDiscriminator_
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "key")
	delete(allFields, "$valueItemDiscriminator")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCustomValueItem() *CustomValueItem {
	p := new(CustomValueItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.CustomValueItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	DEPLOYMENT_GCP     Deployment = 3
	DEPLOYMENT_UNKNOWN Deployment = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *Deployment) name(index int) string {
	names := [...]string{
		"ONPREM",
		"AZURE",
		"AWS",
		"GCP",
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
		"GCP",
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
		"GCP",
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
	*p = *NewEntityReference()

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
	if known.Rel != nil {
		p.Rel = known.Rel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEntityReference() *EntityReference {
	p := new(EntityReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EntityReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewEsxClientConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Certificate != nil {
		p.Certificate = known.Certificate
	}
	if known.ExtensionKey != nil {
		p.ExtensionKey = known.ExtensionKey
	}
	if known.IsInsecure != nil {
		p.IsInsecure = known.IsInsecure
	}
	if known.Key != nil {
		p.Key = known.Key
	}
	if known.Password != nil {
		p.Password = known.Password
	}
	if known.ThumbPrint != nil {
		p.ThumbPrint = known.ThumbPrint
	}
	if known.Username != nil {
		p.Username = known.Username
	}
	if known.VcenterIp != nil {
		p.VcenterIp = known.VcenterIp
	}
	if known.VcenterPort != nil {
		p.VcenterPort = known.VcenterPort
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEsxClientConfig() *EsxClientConfig {
	p := new(EsxClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxClientConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	/*
	  Define multiple NICs for ESX Platform.
	*/
	Nics []EsxNic `json:"nics,omitempty"`
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
	*p = *NewEsxNetworkResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterId != nil {
		p.ClusterId = known.ClusterId
	}
	if known.Location != nil {
		p.Location = known.Location
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.Nics != nil {
		p.Nics = known.Nics
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "location")
	delete(allFields, "network")
	delete(allFields, "nics")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEsxNetworkResourceConfig() *EsxNetworkResourceConfig {
	p := new(EsxNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type EsxNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Network *string `json:"network,omitempty"`

	NicUsage []NicUsage `json:"nicUsage,omitempty"`
}

func (p *EsxNic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EsxNic

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

func (p *EsxNic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EsxNic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEsxNic()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.NicUsage != nil {
		p.NicUsage = known.NicUsage
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "network")
	delete(allFields, "nicUsage")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEsxNic() *EsxNic {
	p := new(EsxNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxNic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewEsxResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.NetworkResource != nil {
		p.NetworkResource = known.NetworkResource
	}
	if known.VmResource != nil {
		p.VmResource = known.VmResource
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkResource")
	delete(allFields, "vmResource")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEsxResourceConfig() *EsxResourceConfig {
	p := new(EsxResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.EsxResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Payload model used to initiate a failover operation on a protection policy.
*/
type FailoverRequest struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FailoverType *FailoverType `json:"failoverType"`
	/*
	  Action to execute after the failover procedure completes.
	*/
	PostFailoverAction []PrePostAction `json:"postFailoverAction,omitempty"`
	/*
	  Action to execute before the failover procedure starts.
	*/
	PreFailoverAction []PrePostAction `json:"preFailoverAction,omitempty"`

	ReplicationTargetCluster *ClusterTarget `json:"replicationTargetCluster"`
	/*
	  Name of the snapshot to fail over to (async replication only).
	*/
	SnapshotName *string `json:"snapshotName,omitempty"`
}

func (p *FailoverRequest) MarshalJSON() ([]byte, error) {
	type FailoverRequestProxy FailoverRequest

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*FailoverRequestProxy
		FailoverType             *FailoverType  `json:"failoverType,omitempty"`
		ReplicationTargetCluster *ClusterTarget `json:"replicationTargetCluster,omitempty"`
	}{
		FailoverRequestProxy:     (*FailoverRequestProxy)(p),
		FailoverType:             p.FailoverType,
		ReplicationTargetCluster: p.ReplicationTargetCluster,
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

func (p *FailoverRequest) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FailoverRequest
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFailoverRequest()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailoverType != nil {
		p.FailoverType = known.FailoverType
	}
	if known.PostFailoverAction != nil {
		p.PostFailoverAction = known.PostFailoverAction
	}
	if known.PreFailoverAction != nil {
		p.PreFailoverAction = known.PreFailoverAction
	}
	if known.ReplicationTargetCluster != nil {
		p.ReplicationTargetCluster = known.ReplicationTargetCluster
	}
	if known.SnapshotName != nil {
		p.SnapshotName = known.SnapshotName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failoverType")
	delete(allFields, "postFailoverAction")
	delete(allFields, "preFailoverAction")
	delete(allFields, "replicationTargetCluster")
	delete(allFields, "snapshotName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFailoverRequest() *FailoverRequest {
	p := new(FailoverRequest)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.FailoverRequest"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Type of failover operation.
*/
type FailoverType int

const (
	FAILOVERTYPE_UNKNOWN            FailoverType = 0
	FAILOVERTYPE_REDACTED           FailoverType = 1
	FAILOVERTYPE_UNPLANNED_FAILOVER FailoverType = 2
	FAILOVERTYPE_PLANNED_FAILOVER   FailoverType = 3
	FAILOVERTYPE_ROLLBACK           FailoverType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FailoverType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPLANNED_FAILOVER",
		"PLANNED_FAILOVER",
		"ROLLBACK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FailoverType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPLANNED_FAILOVER",
		"PLANNED_FAILOVER",
		"ROLLBACK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FailoverType) index(name string) FailoverType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNPLANNED_FAILOVER",
		"PLANNED_FAILOVER",
		"ROLLBACK",
	}
	for idx := range names {
		if names[idx] == name {
			return FailoverType(idx)
		}
	}
	return FAILOVERTYPE_UNKNOWN
}

func (e *FailoverType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FailoverType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FailoverType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FailoverType) Ref() *FailoverType {
	return &e
}

/*
API group and kind pair describing a Kubernetes resource.
*/
type GroupKind struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  API group of the resource.
	*/
	Group *string `json:"group"`
	/*
	  Kind of the resource.
	*/
	Kind *string `json:"kind"`
}

func (p *GroupKind) MarshalJSON() ([]byte, error) {
	type GroupKindProxy GroupKind

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*GroupKindProxy
		Group *string `json:"group,omitempty"`
		Kind  *string `json:"kind,omitempty"`
	}{
		GroupKindProxy: (*GroupKindProxy)(p),
		Group:          p.Group,
		Kind:           p.Kind,
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

func (p *GroupKind) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GroupKind
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGroupKind()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Group != nil {
		p.Group = known.Group
	}
	if known.Kind != nil {
		p.Kind = known.Kind
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "group")
	delete(allFields, "kind")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGroupKind() *GroupKind {
	p := new(GroupKind)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.GroupKind"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewHealth()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HealthComponents != nil {
		p.HealthComponents = known.HealthComponents
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "healthComponents")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHealth() *Health {
	p := new(Health)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Health"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewHealthComponent()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Details != nil {
		p.Details = known.Details
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "details")
	delete(allFields, "name")
	delete(allFields, "status")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHealthComponent() *HealthComponent {
	p := new(HealthComponent)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HealthComponent"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Information about the Helm-managed application being protected.
*/
type HelmApplicationInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Identifier of the application registered in Service Manager.
	*/
	AppId *string `json:"appId"`

	AppType *ApplicationType `json:"appType"`
}

func (p *HelmApplicationInfo) MarshalJSON() ([]byte, error) {
	type HelmApplicationInfoProxy HelmApplicationInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HelmApplicationInfoProxy
		AppId   *string          `json:"appId,omitempty"`
		AppType *ApplicationType `json:"appType,omitempty"`
	}{
		HelmApplicationInfoProxy: (*HelmApplicationInfoProxy)(p),
		AppId:                    p.AppId,
		AppType:                  p.AppType,
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

func (p *HelmApplicationInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HelmApplicationInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHelmApplicationInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AppId != nil {
		p.AppId = known.AppId
	}
	if known.AppType != nil {
		p.AppType = known.AppType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "appId")
	delete(allFields, "appType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHelmApplicationInfo() *HelmApplicationInfo {
	p := new(HelmApplicationInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HelmApplicationInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Details of a Helm job action.
*/
type HelmJobDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Custom values passed to the Helm job.
	*/
	CustomValues []CustomValue `json:"customValues,omitempty"`
	/*
	  Name of the Helm job chart.
	*/
	Name *string `json:"name"`
	/*
	  Version of the Helm job chart.
	*/
	Version *string `json:"version"`
}

func (p *HelmJobDetails) MarshalJSON() ([]byte, error) {
	type HelmJobDetailsProxy HelmJobDetails

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*HelmJobDetailsProxy
		Name    *string `json:"name,omitempty"`
		Version *string `json:"version,omitempty"`
	}{
		HelmJobDetailsProxy: (*HelmJobDetailsProxy)(p),
		Name:                p.Name,
		Version:             p.Version,
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

func (p *HelmJobDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias HelmJobDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewHelmJobDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CustomValues != nil {
		p.CustomValues = known.CustomValues
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "customValues")
	delete(allFields, "name")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHelmJobDetails() *HelmJobDetails {
	p := new(HelmJobDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HelmJobDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewHistory()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ParentAppUuid != nil {
		p.ParentAppUuid = known.ParentAppUuid
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHistory() *History {
	p := new(History)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.History"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewHistoryProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ParentAppUuid != nil {
		p.ParentAppUuid = known.ParentAppUuid
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewHistoryProjection() *HistoryProjection {
	p := new(HistoryProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.HistoryProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AirGapServerUri != nil {
		p.AirGapServerUri = known.AirGapServerUri
	}
	if known.ChartRepositoryUri != nil {
		p.ChartRepositoryUri = known.ChartRepositoryUri
	}
	if known.Deployment != nil {
		p.Deployment = known.Deployment
	}
	if known.Environment != nil {
		p.Environment = known.Environment
	}
	if known.IsAirGapEnabled != nil {
		p.IsAirGapEnabled = known.IsAirGapEnabled
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewInfo() *Info {
	p := new(Info)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Info"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewInstall()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterID != nil {
		p.ClusterID = known.ClusterID
	}
	if known.CustomValues != nil {
		p.CustomValues = known.CustomValues
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterID")
	delete(allFields, "customValues")
	delete(allFields, "name")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewInstall() *Install {
	p := new(Install)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Install"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewIpRange()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EndIp != nil {
		p.EndIp = known.EndIp
	}
	if known.StartIp != nil {
		p.StartIp = known.StartIp
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "endIp")
	delete(allFields, "startIp")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewIpRange() *IpRange {
	p := new(IpRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.IpRange"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of label selector requirements, all of which must be met.
*/
type LabelMatchExpression struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Label key that the selector applies to.
	*/
	Key *string `json:"key,omitempty"`

	Operator *LabelMatchExpressionOperator `json:"operator,omitempty"`
	/*
	  Label selector expression values.
	*/
	Values []string `json:"values,omitempty"`
}

func (p *LabelMatchExpression) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LabelMatchExpression

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

func (p *LabelMatchExpression) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LabelMatchExpression
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLabelMatchExpression()

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
	if known.Operator != nil {
		p.Operator = known.Operator
	}
	if known.Values != nil {
		p.Values = known.Values
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "key")
	delete(allFields, "operator")
	delete(allFields, "values")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLabelMatchExpression() *LabelMatchExpression {
	p := new(LabelMatchExpression)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LabelMatchExpression"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Operator used in the label requirement.
*/
type LabelMatchExpressionOperator int

const (
	LABELMATCHEXPRESSIONOPERATOR_UNKNOWN    LabelMatchExpressionOperator = 0
	LABELMATCHEXPRESSIONOPERATOR_REDACTED   LabelMatchExpressionOperator = 1
	LABELMATCHEXPRESSIONOPERATOR_IN         LabelMatchExpressionOperator = 2
	LABELMATCHEXPRESSIONOPERATOR_NOT_IN     LabelMatchExpressionOperator = 3
	LABELMATCHEXPRESSIONOPERATOR_EXISTS     LabelMatchExpressionOperator = 4
	LABELMATCHEXPRESSIONOPERATOR_NOT_EXISTS LabelMatchExpressionOperator = 5
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *LabelMatchExpressionOperator) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN",
		"NOT_IN",
		"EXISTS",
		"NOT_EXISTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e LabelMatchExpressionOperator) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN",
		"NOT_IN",
		"EXISTS",
		"NOT_EXISTS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *LabelMatchExpressionOperator) index(name string) LabelMatchExpressionOperator {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN",
		"NOT_IN",
		"EXISTS",
		"NOT_EXISTS",
	}
	for idx := range names {
		if names[idx] == name {
			return LabelMatchExpressionOperator(idx)
		}
	}
	return LABELMATCHEXPRESSIONOPERATOR_UNKNOWN
}

func (e *LabelMatchExpressionOperator) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LabelMatchExpressionOperator:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LabelMatchExpressionOperator) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LabelMatchExpressionOperator) Ref() *LabelMatchExpressionOperator {
	return &e
}

/*
A standard Kubernetes label selector used to filter resources based on their labels.
*/
type LabelSelector struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	MatchExpressions []LabelMatchExpression `json:"matchExpressions,omitempty"`
	/*
	  List of key, value pairs representing label selectors. For example, key "app.kubernetes.io/name", value "test-mysql".
	*/
	MatchLabels []import1.KVStringPair `json:"matchLabels,omitempty"`
}

func (p *LabelSelector) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias LabelSelector

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

func (p *LabelSelector) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias LabelSelector
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewLabelSelector()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MatchExpressions != nil {
		p.MatchExpressions = known.MatchExpressions
	}
	if known.MatchLabels != nil {
		p.MatchLabels = known.MatchLabels
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "matchExpressions")
	delete(allFields, "matchLabels")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLabelSelector() *LabelSelector {
	p := new(LabelSelector)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LabelSelector"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancer()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Id != nil {
		p.Id = known.Id
	}
	if known.NetworkConfig != nil {
		p.NetworkConfig = known.NetworkConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "id")
	delete(allFields, "networkConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancer() *LoadBalancer {
	p := new(LoadBalancer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancer"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancerAccessInterface()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cidr != nil {
		p.Cidr = known.Cidr
	}
	if known.Ips != nil {
		p.Ips = known.Ips
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.Subnet != nil {
		p.Subnet = known.Subnet
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cidr")
	delete(allFields, "ips")
	delete(allFields, "network")
	delete(allFields, "subnet")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancerAccessInterface() *LoadBalancerAccessInterface {
	p := new(LoadBalancerAccessInterface)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerAccessInterface"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancerConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LoadBalancerConfigItems != nil {
		p.LoadBalancerConfigItems = known.LoadBalancerConfigItems
	}
	if known.Version != nil {
		p.Version = known.Version
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "loadBalancerConfigItems")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancerConfig() *LoadBalancerConfig {
	p := new(LoadBalancerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancerConfigObject()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Kind != nil {
		p.Kind = known.Kind
	}
	if known.ConfigType != nil {
		p.ConfigType = known.ConfigType
	}
	if known.Instances != nil {
		p.Instances = known.Instances
	}
	if known.MspUuid != nil {
		p.MspUuid = known.MspUuid
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.ResourceConfig != nil {
		p.ResourceConfig = known.ResourceConfig
	}
	if known.Type != nil {
		p.Type = known.Type
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancerConfigObject() *LoadBalancerConfigObject {
	p := new(LoadBalancerConfigObject)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerConfigObject"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancerNetworkConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessInterfaces != nil {
		p.AccessInterfaces = known.AccessInterfaces
	}
	if known.Client != nil {
		p.Client = known.Client
	}
	if known.VmNetworkConfigItemDiscriminator_ != nil {
		p.VmNetworkConfigItemDiscriminator_ = known.VmNetworkConfigItemDiscriminator_
	}
	if known.VmNetworkConfig != nil {
		p.VmNetworkConfig = known.VmNetworkConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessInterfaces")
	delete(allFields, "client")
	delete(allFields, "$vmNetworkConfigItemDiscriminator")
	delete(allFields, "vmNetworkConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancerNetworkConfig() *LoadBalancerNetworkConfig {
	p := new(LoadBalancerNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoadBalancerResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cpu != nil {
		p.Cpu = known.Cpu
	}
	if known.DiskMib != nil {
		p.DiskMib = known.DiskMib
	}
	if known.GoldImageRef != nil {
		p.GoldImageRef = known.GoldImageRef
	}
	if known.GoldImageVersion != nil {
		p.GoldImageVersion = known.GoldImageVersion
	}
	if known.Labels != nil {
		p.Labels = known.Labels
	}
	if known.MemoryMib != nil {
		p.MemoryMib = known.MemoryMib
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoadBalancerResourceConfig() *LoadBalancerResourceConfig {
	p := new(LoadBalancerResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoadBalancerResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewLoggingVolumeConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.MountPath != nil {
		p.MountPath = known.MountPath
	}
	if known.SizeMb != nil {
		p.SizeMb = known.SizeMb
	}
	if known.StorageConfigName != nil {
		p.StorageConfigName = known.StorageConfigName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "mountPath")
	delete(allFields, "sizeMb")
	delete(allFields, "storageConfigName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewLoggingVolumeConfig() *LoggingVolumeConfig {
	p := new(LoggingVolumeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.LoggingVolumeConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Nic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Network *string `json:"network,omitempty"`

	NicUsage []NicUsage `json:"nicUsage,omitempty"`
}

func (p *Nic) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Nic

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

func (p *Nic) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Nic
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNic()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Network != nil {
		p.Network = known.Network
	}
	if known.NicUsage != nil {
		p.NicUsage = known.NicUsage
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "network")
	delete(allFields, "nicUsage")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNic() *Nic {
	p := new(Nic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Nic"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

/*
Nic usage is used to signify the purpose of Nic. It can be one of MANAGEMENT, DATA, ALL, UNMANAGED, INFRASTRUCTURE.
*/
type NicUsage int

const (
	NICUSAGE_UNKNOWN        NicUsage = 0
	NICUSAGE_REDACTED       NicUsage = 1
	NICUSAGE_MANAGEMENT     NicUsage = 2
	NICUSAGE_DATA           NicUsage = 3
	NICUSAGE_ALL            NicUsage = 4
	NICUSAGE_UNMANAGED      NicUsage = 5
	NICUSAGE_INFRASTRUCTURE NicUsage = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NicUsage) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENT",
		"DATA",
		"ALL",
		"UNMANAGED",
		"INFRASTRUCTURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NicUsage) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENT",
		"DATA",
		"ALL",
		"UNMANAGED",
		"INFRASTRUCTURE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NicUsage) index(name string) NicUsage {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MANAGEMENT",
		"DATA",
		"ALL",
		"UNMANAGED",
		"INFRASTRUCTURE",
	}
	for idx := range names {
		if names[idx] == name {
			return NicUsage(idx)
		}
	}
	return NICUSAGE_UNKNOWN
}

func (e *NicUsage) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicUsage:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicUsage) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicUsage) Ref() *NicUsage {
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
	  Indicates if the storage class uses hypervisor-attached mode in CSI 3.x.
	*/
	IsHypervisorAttached *bool `json:"isHypervisorAttached,omitempty"`
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
	*p = *NewNutanixStorageClass()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterRef != nil {
		p.ClusterRef = known.ClusterRef
	}
	if known.FileSystem != nil {
		p.FileSystem = known.FileSystem
	}
	if known.IsFlashMode != nil {
		p.IsFlashMode = known.IsFlashMode
	}
	if known.IsHypervisorAttached != nil {
		p.IsHypervisorAttached = known.IsHypervisorAttached
	}
	if known.StorageContainer != nil {
		p.StorageContainer = known.StorageContainer
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterRef")
	delete(allFields, "fileSystem")
	delete(allFields, "isFlashMode")
	delete(allFields, "isHypervisorAttached")
	delete(allFields, "storageContainer")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNutanixStorageClass() *NutanixStorageClass {
	p := new(NutanixStorageClass)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.NutanixStorageClass"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewOwnerReference()

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
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewPostAppResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AppUUID != nil {
		p.AppUUID = known.AppUUID
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "appUUID")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPostAppResponse() *PostAppResponse {
	p := new(PostAppResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.PostAppResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewPostServiceResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ServiceUUID != nil {
		p.ServiceUUID = known.ServiceUUID
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "serviceUUID")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPostServiceResponse() *PostServiceResponse {
	p := new(PostServiceResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.PostServiceResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Wrapper for an action executed before or after failover.
*/
type PrePostAction struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ActionDetailsItemDiscriminator_ *string `json:"$actionDetailsItemDiscriminator,omitempty"`

	ActionDetails *OneOfPrePostActionActionDetails `json:"actionDetails"`
}

func (p *PrePostAction) MarshalJSON() ([]byte, error) {
	type PrePostActionProxy PrePostAction

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*PrePostActionProxy
		ActionDetails *OneOfPrePostActionActionDetails `json:"actionDetails,omitempty"`
	}{
		PrePostActionProxy: (*PrePostActionProxy)(p),
		ActionDetails:      p.ActionDetails,
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

func (p *PrePostAction) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PrePostAction
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPrePostAction()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ActionDetailsItemDiscriminator_ != nil {
		p.ActionDetailsItemDiscriminator_ = known.ActionDetailsItemDiscriminator_
	}
	if known.ActionDetails != nil {
		p.ActionDetails = known.ActionDetails
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$actionDetailsItemDiscriminator")
	delete(allFields, "actionDetails")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPrePostAction() *PrePostAction {
	p := new(PrePostAction)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.PrePostAction"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PrePostAction) GetActionDetails() interface{} {
	if nil == p.ActionDetails {
		return nil
	}
	return p.ActionDetails.GetValue()
}

func (p *PrePostAction) SetActionDetails(v interface{}) error {
	if nil == p.ActionDetails {
		p.ActionDetails = NewOneOfPrePostActionActionDetails()
	}
	e := p.ActionDetails.SetValue(v)
	if nil == e {
		if nil == p.ActionDetailsItemDiscriminator_ {
			p.ActionDetailsItemDiscriminator_ = new(string)
		}
		*p.ActionDetailsItemDiscriminator_ = *p.ActionDetails.Discriminator
	}
	return e
}

/*
Protection target grouping Kubernetes resources to protect.
*/
type ProtectionTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Name of the protection target. Used when deploying NDK Application custom resource.
	*/
	Name *string `json:"name"`
	/*
	  Namespace where protected resources are located.
	*/
	Namespace *string `json:"namespace"`
	/*
	  Selectors for protected resources. Each RLS selects resources based on specified label selector and then applies excludes and includes rules for those returned resources. Omitting this will protect all the resources in the namespace.
	*/
	Selectors []ResourceLabelSelector `json:"selectors,omitempty"`
}

func (p *ProtectionTarget) MarshalJSON() ([]byte, error) {
	type ProtectionTargetProxy ProtectionTarget

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ProtectionTargetProxy
		Name      *string `json:"name,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
	}{
		ProtectionTargetProxy: (*ProtectionTargetProxy)(p),
		Name:                  p.Name,
		Namespace:             p.Namespace,
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

func (p *ProtectionTarget) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ProtectionTarget
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewProtectionTarget()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Namespace != nil {
		p.Namespace = known.Namespace
	}
	if known.Selectors != nil {
		p.Selectors = known.Selectors
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")
	delete(allFields, "namespace")
	delete(allFields, "selectors")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewProtectionTarget() *ProtectionTarget {
	p := new(ProtectionTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ProtectionTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewRegistryBundle()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.Filepath != nil {
		p.Filepath = known.Filepath
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterUuid")
	delete(allFields, "filepath")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRegistryBundle() *RegistryBundle {
	p := new(RegistryBundle)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.RegistryBundle"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Replication target definition containing either async- or sync-replication settings.
*/
type ReplicationTarget struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	ReplicationTargetConfigItemDiscriminator_ *string `json:"$replicationTargetConfigItemDiscriminator,omitempty"`

	ReplicationTargetConfig *OneOfReplicationTargetReplicationTargetConfig `json:"replicationTargetConfig"`
}

func (p *ReplicationTarget) MarshalJSON() ([]byte, error) {
	type ReplicationTargetProxy ReplicationTarget

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ReplicationTargetProxy
		ReplicationTargetConfig *OneOfReplicationTargetReplicationTargetConfig `json:"replicationTargetConfig,omitempty"`
	}{
		ReplicationTargetProxy:  (*ReplicationTargetProxy)(p),
		ReplicationTargetConfig: p.ReplicationTargetConfig,
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

func (p *ReplicationTarget) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReplicationTarget
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReplicationTarget()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ReplicationTargetConfigItemDiscriminator_ != nil {
		p.ReplicationTargetConfigItemDiscriminator_ = known.ReplicationTargetConfigItemDiscriminator_
	}
	if known.ReplicationTargetConfig != nil {
		p.ReplicationTargetConfig = known.ReplicationTargetConfig
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$replicationTargetConfigItemDiscriminator")
	delete(allFields, "replicationTargetConfig")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReplicationTarget() *ReplicationTarget {
	p := new(ReplicationTarget)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ReplicationTarget"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ReplicationTarget) GetReplicationTargetConfig() interface{} {
	if nil == p.ReplicationTargetConfig {
		return nil
	}
	return p.ReplicationTargetConfig.GetValue()
}

func (p *ReplicationTarget) SetReplicationTargetConfig(v interface{}) error {
	if nil == p.ReplicationTargetConfig {
		p.ReplicationTargetConfig = NewOneOfReplicationTargetReplicationTargetConfig()
	}
	e := p.ReplicationTargetConfig.SetValue(v)
	if nil == e {
		if nil == p.ReplicationTargetConfigItemDiscriminator_ {
			p.ReplicationTargetConfigItemDiscriminator_ = new(string)
		}
		*p.ReplicationTargetConfigItemDiscriminator_ = *p.ReplicationTargetConfig.Discriminator
	}
	return e
}

/*
Selector that filters resources based on labels and include/exclude rules.
*/
type ResourceLabelSelector struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExcludeResources []GroupKind `json:"excludeResources,omitempty"`

	IncludeResources []GroupKind `json:"includeResources,omitempty"`

	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`
}

func (p *ResourceLabelSelector) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ResourceLabelSelector

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

func (p *ResourceLabelSelector) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ResourceLabelSelector
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewResourceLabelSelector()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExcludeResources != nil {
		p.ExcludeResources = known.ExcludeResources
	}
	if known.IncludeResources != nil {
		p.IncludeResources = known.IncludeResources
	}
	if known.LabelSelector != nil {
		p.LabelSelector = known.LabelSelector
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "excludeResources")
	delete(allFields, "includeResources")
	delete(allFields, "labelSelector")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewResourceLabelSelector() *ResourceLabelSelector {
	p := new(ResourceLabelSelector)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ResourceLabelSelector"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewService()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.IsInactive != nil {
		p.IsInactive = known.IsInactive
	}
	if known.LastUpdatedTimestamp != nil {
		p.LastUpdatedTimestamp = known.LastUpdatedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewService() *Service {
	p := new(Service)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Service"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewServiceProjection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Action != nil {
		p.Action = known.Action
	}
	if known.Apptype != nil {
		p.Apptype = known.Apptype
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.CreatedTimestamp != nil {
		p.CreatedTimestamp = known.CreatedTimestamp
	}
	if known.IsInactive != nil {
		p.IsInactive = known.IsInactive
	}
	if known.LastUpdatedTimestamp != nil {
		p.LastUpdatedTimestamp = known.LastUpdatedTimestamp
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.Uuid != nil {
		p.Uuid = known.Uuid
	}
	if known.Version != nil {
		p.Version = known.Version
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewServiceProjection() *ServiceProjection {
	p := new(ServiceProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.ServiceProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that represents an application snapshot captured by App DR.
*/
type Snapshot struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the associated app protection policy.
	*/
	AppProtectionPolicyExtId *string `json:"appProtectionPolicyExtId"`
	/*
	  Time at which the snapshot operation completed.
	*/
	CompletedTime *time.Time `json:"completedTime,omitempty"`
	/*
	  Object creation timestamp.
	*/
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Point-in-time capture of metadata.
	*/
	Metadata []import1.KVStringPair `json:"metadata,omitempty"`
	/*
	  Name of the snapshot.
	*/
	Name *string `json:"name"`

	Status *SnapshotStatus `json:"status"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Last update timestamp.
	*/
	UpdatedTime *time.Time `json:"updatedTime,omitempty"`
}

func (p *Snapshot) MarshalJSON() ([]byte, error) {
	type SnapshotProxy Snapshot

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SnapshotProxy
		AppProtectionPolicyExtId *string         `json:"appProtectionPolicyExtId,omitempty"`
		Name                     *string         `json:"name,omitempty"`
		Status                   *SnapshotStatus `json:"status,omitempty"`
	}{
		SnapshotProxy:            (*SnapshotProxy)(p),
		AppProtectionPolicyExtId: p.AppProtectionPolicyExtId,
		Name:                     p.Name,
		Status:                   p.Status,
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

func (p *Snapshot) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Snapshot
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSnapshot()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AppProtectionPolicyExtId != nil {
		p.AppProtectionPolicyExtId = known.AppProtectionPolicyExtId
	}
	if known.CompletedTime != nil {
		p.CompletedTime = known.CompletedTime
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.UpdatedTime != nil {
		p.UpdatedTime = known.UpdatedTime
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "appProtectionPolicyExtId")
	delete(allFields, "completedTime")
	delete(allFields, "createdTime")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "metadata")
	delete(allFields, "name")
	delete(allFields, "status")
	delete(allFields, "tenantId")
	delete(allFields, "updatedTime")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSnapshot() *Snapshot {
	p := new(Snapshot)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.Snapshot"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Status of the snapshot lifecycle.
*/
type SnapshotStatus int

const (
	SNAPSHOTSTATUS_UNKNOWN     SnapshotStatus = 0
	SNAPSHOTSTATUS_REDACTED    SnapshotStatus = 1
	SNAPSHOTSTATUS_IN_PROGRESS SnapshotStatus = 2
	SNAPSHOTSTATUS_READY       SnapshotStatus = 3
	SNAPSHOTSTATUS_FAILED      SnapshotStatus = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *SnapshotStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"READY",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e SnapshotStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"READY",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *SnapshotStatus) index(name string) SnapshotStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IN_PROGRESS",
		"READY",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return SnapshotStatus(idx)
		}
	}
	return SNAPSHOTSTATUS_UNKNOWN
}

func (e *SnapshotStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SnapshotStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SnapshotStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SnapshotStatus) Ref() *SnapshotStatus {
	return &e
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
Configuration for a synchronous replication target.
*/
type SyncReplicationTargetConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  External identifier of the replication-target cluster.
	*/
	ClusterExtId *string `json:"clusterExtId"`
	/*
	  External identifier of the Prism Central managing the replication-target cluster.
	*/
	DomainManagerExtId *string `json:"domainManagerExtId"`
}

func (p *SyncReplicationTargetConfig) MarshalJSON() ([]byte, error) {
	type SyncReplicationTargetConfigProxy SyncReplicationTargetConfig

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*SyncReplicationTargetConfigProxy
		ClusterExtId       *string `json:"clusterExtId,omitempty"`
		DomainManagerExtId *string `json:"domainManagerExtId,omitempty"`
	}{
		SyncReplicationTargetConfigProxy: (*SyncReplicationTargetConfigProxy)(p),
		ClusterExtId:                     p.ClusterExtId,
		DomainManagerExtId:               p.DomainManagerExtId,
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

func (p *SyncReplicationTargetConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SyncReplicationTargetConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSyncReplicationTargetConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtId != nil {
		p.ClusterExtId = known.ClusterExtId
	}
	if known.DomainManagerExtId != nil {
		p.DomainManagerExtId = known.DomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterExtId")
	delete(allFields, "domainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSyncReplicationTargetConfig() *SyncReplicationTargetConfig {
	p := new(SyncReplicationTargetConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.SyncReplicationTargetConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p = *NewTaskReferenceInternal()

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
	if known.Href != nil {
		p.Href = known.Href
	}
	if known.Rel != nil {
		p.Rel = known.Rel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "href")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTaskReferenceInternal() *TaskReferenceInternal {
	p := new(TaskReferenceInternal)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskReferenceInternal"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewTaskStep()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Name != nil {
		p.Name = known.Name
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "name")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTaskStep() *TaskStep {
	p := new(TaskStep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskStep"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewTaskV2()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.CompletedTime != nil {
		p.CompletedTime = known.CompletedTime
	}
	if known.CompletionDetails != nil {
		p.CompletionDetails = known.CompletionDetails
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.EntitiesAffected != nil {
		p.EntitiesAffected = known.EntitiesAffected
	}
	if known.ErrorMessages != nil {
		p.ErrorMessages = known.ErrorMessages
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsCancelable != nil {
		p.IsCancelable = known.IsCancelable
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyErrorMessage != nil {
		p.LegacyErrorMessage = known.LegacyErrorMessage
	}
	if known.MspUuid != nil {
		p.MspUuid = known.MspUuid
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Operation != nil {
		p.Operation = known.Operation
	}
	if known.OperationDescription != nil {
		p.OperationDescription = known.OperationDescription
	}
	if known.OwnedBy != nil {
		p.OwnedBy = known.OwnedBy
	}
	if known.ParentTask != nil {
		p.ParentTask = known.ParentTask
	}
	if known.ProgressPercentage != nil {
		p.ProgressPercentage = known.ProgressPercentage
	}
	if known.StartedTime != nil {
		p.StartedTime = known.StartedTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.SubSteps != nil {
		p.SubSteps = known.SubSteps
	}
	if known.SubTasks != nil {
		p.SubTasks = known.SubTasks
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Warnings != nil {
		p.Warnings = known.Warnings
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTaskV2() *TaskV2 {
	p := new(TaskV2)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskV2"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewTaskV2Projection()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterExtIds != nil {
		p.ClusterExtIds = known.ClusterExtIds
	}
	if known.CompletedTime != nil {
		p.CompletedTime = known.CompletedTime
	}
	if known.CompletionDetails != nil {
		p.CompletionDetails = known.CompletionDetails
	}
	if known.CreatedTime != nil {
		p.CreatedTime = known.CreatedTime
	}
	if known.EntitiesAffected != nil {
		p.EntitiesAffected = known.EntitiesAffected
	}
	if known.ErrorMessages != nil {
		p.ErrorMessages = known.ErrorMessages
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.IsCancelable != nil {
		p.IsCancelable = known.IsCancelable
	}
	if known.LastUpdatedTime != nil {
		p.LastUpdatedTime = known.LastUpdatedTime
	}
	if known.LegacyErrorMessage != nil {
		p.LegacyErrorMessage = known.LegacyErrorMessage
	}
	if known.MspUuid != nil {
		p.MspUuid = known.MspUuid
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.Operation != nil {
		p.Operation = known.Operation
	}
	if known.OperationDescription != nil {
		p.OperationDescription = known.OperationDescription
	}
	if known.OwnedBy != nil {
		p.OwnedBy = known.OwnedBy
	}
	if known.ParentTask != nil {
		p.ParentTask = known.ParentTask
	}
	if known.ProgressPercentage != nil {
		p.ProgressPercentage = known.ProgressPercentage
	}
	if known.StartedTime != nil {
		p.StartedTime = known.StartedTime
	}
	if known.Status != nil {
		p.Status = known.Status
	}
	if known.SubSteps != nil {
		p.SubSteps = known.SubSteps
	}
	if known.SubTasks != nil {
		p.SubTasks = known.SubTasks
	}
	if known.Type != nil {
		p.Type = known.Type
	}
	if known.Warnings != nil {
		p.Warnings = known.Warnings
	}

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
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTaskV2Projection() *TaskV2Projection {
	p := new(TaskV2Projection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.TaskV2Projection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewVcenterLocation()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterName != nil {
		p.ClusterName = known.ClusterName
	}
	if known.Datacenter != nil {
		p.Datacenter = known.Datacenter
	}
	if known.VcenterIp != nil {
		p.VcenterIp = known.VcenterIp
	}
	if known.VcenterPort != nil {
		p.VcenterPort = known.VcenterPort
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterName")
	delete(allFields, "datacenter")
	delete(allFields, "vcenterIp")
	delete(allFields, "vcenterPort")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVcenterLocation() *VcenterLocation {
	p := new(VcenterLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VcenterLocation"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewVmNetworkResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterId != nil {
		p.ClusterId = known.ClusterId
	}
	if known.Network != nil {
		p.Network = known.Network
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterId")
	delete(allFields, "network")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVmNetworkResourceConfig() *VmNetworkResourceConfig {
	p := new(VmNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VmNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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
	*p = *NewVmResourceConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Cpu != nil {
		p.Cpu = known.Cpu
	}
	if known.DiskMib != nil {
		p.DiskMib = known.DiskMib
	}
	if known.MemoryMib != nil {
		p.MemoryMib = known.MemoryMib
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "cpu")
	delete(allFields, "diskMib")
	delete(allFields, "memoryMib")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewVmResourceConfig() *VmResourceConfig {
	p := new(VmResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "lifecycle.v4.svcmgr.VmResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

type OneOfReplicationTargetReplicationTargetConfig struct {
	Discriminator *string                       `json:"-"`
	ObjectType_   *string                       `json:"-"`
	oneOfType1    *SyncReplicationTargetConfig  `json:"-"`
	oneOfType0    *AsyncReplicationTargetConfig `json:"-"`
}

func NewOneOfReplicationTargetReplicationTargetConfig() *OneOfReplicationTargetReplicationTargetConfig {
	p := new(OneOfReplicationTargetReplicationTargetConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfReplicationTargetReplicationTargetConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfReplicationTargetReplicationTargetConfig is nil"))
	}
	switch v.(type) {
	case SyncReplicationTargetConfig:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(SyncReplicationTargetConfig)
		}
		*p.oneOfType1 = v.(SyncReplicationTargetConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case AsyncReplicationTargetConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AsyncReplicationTargetConfig)
		}
		*p.oneOfType0 = v.(AsyncReplicationTargetConfig)
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

func (p *OneOfReplicationTargetReplicationTargetConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfReplicationTargetReplicationTargetConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(SyncReplicationTargetConfig)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "lifecycle.v4.svcmgr.SyncReplicationTargetConfig" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(SyncReplicationTargetConfig)
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
	vOneOfType0 := new(AsyncReplicationTargetConfig)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "lifecycle.v4.svcmgr.AsyncReplicationTargetConfig" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AsyncReplicationTargetConfig)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfReplicationTargetReplicationTargetConfig"))
}

func (p *OneOfReplicationTargetReplicationTargetConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfReplicationTargetReplicationTargetConfig")
}

type OneOfPrePostActionActionDetails struct {
	Discriminator *string         `json:"-"`
	ObjectType_   *string         `json:"-"`
	oneOfType1    *APICallDetails `json:"-"`
	oneOfType0    *HelmJobDetails `json:"-"`
}

func NewOneOfPrePostActionActionDetails() *OneOfPrePostActionActionDetails {
	p := new(OneOfPrePostActionActionDetails)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPrePostActionActionDetails) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPrePostActionActionDetails is nil"))
	}
	switch v.(type) {
	case APICallDetails:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(APICallDetails)
		}
		*p.oneOfType1 = v.(APICallDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case HelmJobDetails:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(HelmJobDetails)
		}
		*p.oneOfType0 = v.(HelmJobDetails)
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

func (p *OneOfPrePostActionActionDetails) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfPrePostActionActionDetails) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(APICallDetails)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "lifecycle.v4.svcmgr.APICallDetails" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(APICallDetails)
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
	vOneOfType0 := new(HelmJobDetails)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "lifecycle.v4.svcmgr.HelmJobDetails" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(HelmJobDetails)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPrePostActionActionDetails"))
}

func (p *OneOfPrePostActionActionDetails) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfPrePostActionActionDetails")
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
