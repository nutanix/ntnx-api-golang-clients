/*
 * Generated file models/prism/v4/management/management_model.go.
 *
 * Product version: 4.4.1-beta-1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.management of Nutanix Multidomain Versioned APIs
*/
package management

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
)

/*
Type of cluster to be connected:
- DOMAIN_MANAGER : Domain manager (Prism Central) instance
- AOS : Prism Element cluster instance
- WITNESS_SERVICE: Witness Service instance
*/
type ClusterType int

const (
	CLUSTERTYPE_UNKNOWN         ClusterType = 0
	CLUSTERTYPE_REDACTED        ClusterType = 1
	CLUSTERTYPE_DOMAIN_MANAGER  ClusterType = 2
	CLUSTERTYPE_AOS             ClusterType = 3
	CLUSTERTYPE_WITNESS_SERVICE ClusterType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DOMAIN_MANAGER",
		"AOS",
		"WITNESS_SERVICE",
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
		"DOMAIN_MANAGER",
		"AOS",
		"WITNESS_SERVICE",
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
		"DOMAIN_MANAGER",
		"AOS",
		"WITNESS_SERVICE",
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
This includes the attributes of a remote cluster, such as the cluster name, cluster type, and address details. The address details comprise the external address (either a virtual IP or FQDN), the port.
*/
type RemoteCluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ClusterType *ClusterType `json:"clusterType,omitempty"`
	/*
	  Version of the cluster. This could be a version of AOS cluster, Domain Manager, or WitnessVM.
	*/
	ClusterVersion *string `json:"clusterVersion,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	ExternalAddress *import1.IPAddressOrFQDN `json:"externalAddress"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  Cluster name of a remote cluster.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  Node IP addresses of a registered cluster.
	*/
	NodeIpAddresses []import1.IPAddressOrFQDN `json:"nodeIpAddresses,omitempty"`
	/*
	  Port of the cluster exposed externally. This would be 9440 by default.
	*/
	Port *int64 `json:"port,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RemoteCluster) MarshalJSON() ([]byte, error) {
	type RemoteClusterProxy RemoteCluster

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RemoteClusterProxy
		ExternalAddress *import1.IPAddressOrFQDN `json:"externalAddress,omitempty"`
	}{
		RemoteClusterProxy: (*RemoteClusterProxy)(p),
		ExternalAddress:    p.ExternalAddress,
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

func (p *RemoteCluster) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RemoteCluster
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRemoteCluster()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterType != nil {
		p.ClusterType = known.ClusterType
	}
	if known.ClusterVersion != nil {
		p.ClusterVersion = known.ClusterVersion
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.ExternalAddress != nil {
		p.ExternalAddress = known.ExternalAddress
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.NodeIpAddresses != nil {
		p.NodeIpAddresses = known.NodeIpAddresses
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterType")
	delete(allFields, "clusterVersion")
	delete(allFields, "extId")
	delete(allFields, "externalAddress")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "nodeIpAddresses")
	delete(allFields, "port")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRemoteCluster() *RemoteCluster {
	p := new(RemoteCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RemoteCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4"}
	p.UnknownFields_ = map[string]interface{}{}

	p.Port = new(int64)
	*p.Port = 9440

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
