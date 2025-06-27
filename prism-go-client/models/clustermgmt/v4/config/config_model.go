/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.config of Nutanix Prism APIs
*/
package config

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/config"
)

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
	*p = BuildInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "version")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewBuildInfo() *BuildInfo {
	p := new(BuildInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BuildInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p = ClusterConfig(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "buildInfo")
	delete(allFields, "shouldEnableLockdownMode")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterConfig() *ClusterConfig {
	p := new(ClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	*p = ClusterNetwork(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "externalAddress")
	delete(allFields, "fqdn")
	delete(allFields, "nameServers")
	delete(allFields, "ntpServers")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewClusterNetwork() *ClusterNetwork {
	p := new(ClusterNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterNetwork"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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
