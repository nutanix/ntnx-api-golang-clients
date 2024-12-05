/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.config of Nutanix Prism APIs
*/
package config

import (
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

func NewBuildInfo() *BuildInfo {
	p := new(BuildInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BuildInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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

func NewClusterConfig() *ClusterConfig {
	p := new(ClusterConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
	  List of name servers on a cluster. This is part of payload for both cluster create & update operations. For create operation, only ipv4 address / fqdn values are supported currently.
	*/
	NameServers []import1.IPAddressOrFQDN `json:"nameServers,omitempty"`
	/*
	  List of NTP servers on a cluster. This is part of payload for both cluster create & update operations. For create operation, only ipv4 address / fqdn values are supported currently.
	*/
	NtpServers []import1.IPAddressOrFQDN `json:"ntpServers,omitempty"`
}

func NewClusterNetwork() *ClusterNetwork {
	p := new(ClusterNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterNetwork"
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
