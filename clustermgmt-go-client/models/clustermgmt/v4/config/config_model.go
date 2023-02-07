/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.0.2-alpha-2
 *
 * Part of the Nutanix Clustermgmt Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Hosts, Clusters and other Infrastructure
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/error"
	import4 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/prism/v4/config"
)

/**
Status of Acropolis connection to hypervisor
*/
type AcropolisConnectionState int

const (
	ACROPOLISCONNECTIONSTATE_UNKNOWN      AcropolisConnectionState = 0
	ACROPOLISCONNECTIONSTATE_REDACTED     AcropolisConnectionState = 1
	ACROPOLISCONNECTIONSTATE_CONNECTED    AcropolisConnectionState = 2
	ACROPOLISCONNECTIONSTATE_DISCONNECTED AcropolisConnectionState = 3
)

// returns the name of the enum given an ordinal number
func (e *AcropolisConnectionState) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/expand-cluster Post operation
*/
type AddNodeTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddNodeTaskResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddNodeTaskResponse() *AddNodeTaskResponse {
	p := new(AddNodeTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.AddNodeTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AddNodeTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddNodeTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddNodeTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddNodeTaskResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers Post operation
*/
type AddRsyslogServerTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddRsyslogServerTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddRsyslogServerTaskApiResponse() *AddRsyslogServerTaskApiResponse {
	p := new(AddRsyslogServerTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.AddRsyslogServerTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AddRsyslogServerTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddRsyslogServerTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddRsyslogServerTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddRsyslogServerTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/add-transports Post operation
*/
type AddSnmpTransportsTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddSnmpTransportsTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddSnmpTransportsTaskApiResponse() *AddSnmpTransportsTaskApiResponse {
	p := new(AddSnmpTransportsTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.AddSnmpTransportsTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AddSnmpTransportsTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddSnmpTransportsTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddSnmpTransportsTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddSnmpTransportsTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps Post operation
*/
type AddSnmpTrapTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddSnmpTrapTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddSnmpTrapTaskApiResponse() *AddSnmpTrapTaskApiResponse {
	p := new(AddSnmpTrapTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.AddSnmpTrapTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AddSnmpTrapTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddSnmpTrapTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddSnmpTrapTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddSnmpTrapTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users Post operation
*/
type AddSnmpUserTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddSnmpUserTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddSnmpUserTaskApiResponse() *AddSnmpUserTaskApiResponse {
	p := new(AddSnmpUserTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.AddSnmpUserTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AddSnmpUserTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddSnmpUserTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddSnmpUserTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddSnmpUserTaskApiResponseData()
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
Indicates whether the address type is IPV4/IPV6
*/
type AddressType int

const (
	ADDRESSTYPE_UNKNOWN  AddressType = 0
	ADDRESSTYPE_REDACTED AddressType = 1
	ADDRESSTYPE_IPV4     AddressType = 2
	ADDRESSTYPE_IPV6     AddressType = 3
)

// returns the name of the enum given an ordinal number
func (e *AddressType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4",
		"IPV6",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *AddressType) index(name string) AddressType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IPV4",
		"IPV6",
	}
	for idx := range names {
		if names[idx] == name {
			return AddressType(idx)
		}
	}
	return ADDRESSTYPE_UNKNOWN
}

func (e *AddressType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for AddressType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *AddressType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e AddressType) Ref() *AddressType {
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.AttributeItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Block item containing block serial and rack name
*/
type BlockItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit serial name
	*/
	BlockId *string `json:"blockId,omitempty"`
	/**
	  Rack name
	*/
	RackName *string `json:"rackName,omitempty"`
}

func NewBlockItem() *BlockItem {
	p := new(BlockItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BlockItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.BlockItem"}
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.BuildReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Hypervisor bundle information
*/
type BundleInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the hypervisor bundle
	*/
	Name *string `json:"name,omitempty"`
}

func NewBundleInfo() *BundleInfo {
	p := new(BundleInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BundleInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.BundleInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
ISO attributes to validate compatibility
*/
type BundleParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BundleInfo *BundleInfo `json:"bundleInfo,omitempty"`
	/**
	  List of node attributes for validating bundle compatibility
	*/
	NodeList []NodeInfo `json:"nodeList,omitempty"`
}

func NewBundleParam() *BundleParam {
	p := new(BundleParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.BundleParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.BundleParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Cluster arch
*/
type ClusterArchReference int

const (
	CLUSTERARCHREFERENCE_UNKNOWN  ClusterArchReference = 0
	CLUSTERARCHREFERENCE_REDACTED ClusterArchReference = 1
	CLUSTERARCHREFERENCE_X86_64   ClusterArchReference = 2
	CLUSTERARCHREFERENCE_PPC64LE  ClusterArchReference = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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
	  Indicates whether the password ssh into the cluster is enabled or not
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ClusterConfigReference"}
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
	Links []import3.ApiLink `json:"links,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ClusterEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Cluster function
*/
type ClusterFunctionRef int

const (
	CLUSTERFUNCTIONREF_UNKNOWN            ClusterFunctionRef = 0
	CLUSTERFUNCTIONREF_REDACTED           ClusterFunctionRef = 1
	CLUSTERFUNCTIONREF_AOS                ClusterFunctionRef = 2
	CLUSTERFUNCTIONREF_PRISM_CENTRAL      ClusterFunctionRef = 3
	CLUSTERFUNCTIONREF_CLOUD_DATA_GATEWAY ClusterFunctionRef = 4
	CLUSTERFUNCTIONREF_AFS                ClusterFunctionRef = 5
	CLUSTERFUNCTIONREF_WITNESS            ClusterFunctionRef = 6
	CLUSTERFUNCTIONREF_XI_PORTAL          ClusterFunctionRef = 7
	CLUSTERFUNCTIONREF_ONE_NODE           ClusterFunctionRef = 8
	CLUSTERFUNCTIONREF_TWO_NODE           ClusterFunctionRef = 9
)

// returns the name of the enum given an ordinal number
func (e *ClusterFunctionRef) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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

	ExternalAddress *import4.IPAddress `json:"externalAddress,omitempty"`

	ExternalDataServiceIp *import4.IPAddress `json:"externalDataServiceIp,omitempty"`
	/**
	  Cluster external subnet address
	*/
	ExternalSubnet *string `json:"externalSubnet,omitempty"`
	/**
	  Cluster fully qualified domain name
	*/
	Fqdn *string `json:"fqdn,omitempty"`
	/**
	  Cluster internal subnet address
	*/
	InternalSubnet *string `json:"internalSubnet,omitempty"`

	ManagementServer *ManagementServerRef `json:"managementServer,omitempty"`

	MasqueradingIp *import4.IPAddress `json:"masqueradingIp,omitempty"`
	/**
	  The port to connect to the cluster when using masquerading IP
	*/
	MasqueradingPort *int `json:"masqueradingPort,omitempty"`
	/**
	  Name servers on a cluster
	*/
	NameServerIpList []import4.IPAddress `json:"nameServerIpList,omitempty"`
	/**
	  NFS subnet whitelist addresses
	*/
	NfsSubnetWhitelist []string `json:"nfsSubnetWhitelist,omitempty"`
	/**
	  NTP servers on a cluster
	*/
	NtpServerIpList []import4.IPAddress `json:"ntpServerIpList,omitempty"`

	SmtpServer *SmtpServerRef `json:"smtpServer,omitempty"`
}

func NewClusterNetworkReference() *ClusterNetworkReference {
	p := new(ClusterNetworkReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterNetworkReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ClusterNetworkReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Cluster reference for an entity
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ClusterReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/stats Get operation
*/
type ClusterStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfClusterStatsApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewClusterStatsApiResponse() *ClusterStatsApiResponse {
	p := new(ClusterStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ClusterStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ClusterStatsApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ClusterStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfClusterStatsApiResponseData()
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
	LastUpdatesSecs *string `json:"lastUpdatesSecs,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ComponentFaultTolerance"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Type of component
*/
type ComponentType int

const (
	COMPONENTTYPE_UNKNOWN                 ComponentType = 0
	COMPONENTTYPE_REDACTED                ComponentType = 1
	COMPONENTTYPE_EXTENT_GROUP_REPLICAS   ComponentType = 2
	COMPONENTTYPE_OPLOG_EPISODES          ComponentType = 3
	COMPONENTTYPE_CASSANDRA_RING          ComponentType = 4
	COMPONENTTYPE_ZOOKEPER_INSTANCES      ComponentType = 5
	COMPONENTTYPE_FREE_SPACE              ComponentType = 6
	COMPONENTTYPE_STATIC_CONFIG           ComponentType = 7
	COMPONENTTYPE_ERASURE_CODE_STRIP_SIZE ComponentType = 8
	COMPONENTTYPE_STARGATE_HEALTH         ComponentType = 9
)

// returns the name of the enum given an ordinal number
func (e *ComponentType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
Compute node details
*/
type ComputeNodeItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit Id in which node resides
	*/
	BlockId *string `json:"blockId,omitempty"`
	/**
	  List of objects containing digital_certificate_base64 and key_management_server_uuid fields for key management server
	*/
	DigitalCertificateMapList []DigitalCertificateMapReference `json:"digitalCertificateMapList,omitempty"`
	/**
	  Name of the host
	*/
	HypervisorHostname *string `json:"hypervisorHostname,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`

	IpmiIp *import4.IPAddress `json:"ipmiIp,omitempty"`
	/**
	  Rackable unit model type
	*/
	Model *string `json:"model,omitempty"`
	/**
	  Position of a node in a rackable unit
	*/
	NodePosition *string `json:"nodePosition,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
}

func NewComputeNodeItem() *ComputeNodeItem {
	p := new(ComputeNodeItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ComputeNodeItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ComputeNodeItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Host entity with its attributes
*/
type ControllerVmReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BackplaneAddress *import4.IPAddress `json:"backplaneAddress,omitempty"`

	ExternalAddress *import4.IPAddress `json:"externalAddress,omitempty"`
	/**
	  Controller VM Id
	*/
	Id *int64 `json:"id,omitempty"`

	Ipmi *IpmiReference `json:"ipmi,omitempty"`
	/**
	  Maintenance mode status
	*/
	MaintenanceMode *bool `json:"maintenanceMode,omitempty"`

	NatIp *import4.IPAddress `json:"natIp,omitempty"`
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
	RdmaBackplaneAddress []import4.IPAddress `json:"rdmaBackplaneAddress,omitempty"`
}

func NewControllerVmReference() *ControllerVmReference {
	p := new(ControllerVmReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ControllerVmReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ControllerVmReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId} Delete operation
*/
type DeleteRsyslogServerTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteRsyslogServerTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteRsyslogServerTaskApiResponse() *DeleteRsyslogServerTaskApiResponse {
	p := new(DeleteRsyslogServerTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.DeleteRsyslogServerTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DeleteRsyslogServerTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteRsyslogServerTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteRsyslogServerTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteRsyslogServerTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId} Delete operation
*/
type DeleteSnmpTrapTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSnmpTrapTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteSnmpTrapTaskApiResponse() *DeleteSnmpTrapTaskApiResponse {
	p := new(DeleteSnmpTrapTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.DeleteSnmpTrapTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DeleteSnmpTrapTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSnmpTrapTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSnmpTrapTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSnmpTrapTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId} Delete operation
*/
type DeleteSnmpUserTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteSnmpUserTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteSnmpUserTaskApiResponse() *DeleteSnmpUserTaskApiResponse {
	p := new(DeleteSnmpUserTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.DeleteSnmpUserTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DeleteSnmpUserTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteSnmpUserTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteSnmpUserTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteSnmpUserTaskApiResponseData()
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
Object containing digital_certificate_base64 and key_management_server_uuid fields for key management server
*/
type DigitalCertificateMapReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Field containing digital_certificate_base64 and key_management_server_uuid for key management server
	*/
	Key *string `json:"key,omitempty"`
	/**
	  Value for the fields digital_certificate_base64 and key_management_server_uuid for key management server
	*/
	Value *string `json:"value,omitempty"`
}

func NewDigitalCertificateMapReference() *DigitalCertificateMapReference {
	p := new(DigitalCertificateMapReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.DigitalCertificateMapReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DigitalCertificateMapReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/discover-unconfigured-nodes Post operation
*/
type DiscoverNodeTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDiscoverNodeTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDiscoverNodeTaskApiResponse() *DiscoverNodeTaskApiResponse {
	p := new(DiscoverNodeTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.DiscoverNodeTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DiscoverNodeTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DiscoverNodeTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DiscoverNodeTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDiscoverNodeTaskApiResponseData()
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DiskReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Domain awareness level
*/
type DomainAwarenessLevel int

const (
	DOMAINAWARENESSLEVEL_UNKNOWN  DomainAwarenessLevel = 0
	DOMAINAWARENESSLEVEL_REDACTED DomainAwarenessLevel = 1
	DOMAINAWARENESSLEVEL_NODE     DomainAwarenessLevel = 2
	DOMAINAWARENESSLEVEL_BLOCK    DomainAwarenessLevel = 3
	DOMAINAWARENESSLEVEL_RACK     DomainAwarenessLevel = 4
)

// returns the name of the enum given an ordinal number
func (e *DomainAwarenessLevel) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
Domain fault tolerance configuration
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.DomainFaultTolerance"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Type of domain
*/
type DomainType int

const (
	DOMAINTYPE_UNKNOWN                  DomainType = 0
	DOMAINTYPE_REDACTED                 DomainType = 1
	DOMAINTYPE_DEPRECATED_RACKABLE_UNIT DomainType = 2
	DOMAINTYPE_DEPRECATED_NODE          DomainType = 3
	DOMAINTYPE_DEPRECATED_DISK          DomainType = 4
	DOMAINTYPE_CUSTOM                   DomainType = 5
	DOMAINTYPE_DISK                     DomainType = 6
	DOMAINTYPE_NODE                     DomainType = 7
	DOMAINTYPE_RACKABLE_UNIT            DomainType = 8
	DOMAINTYPE_RACK                     DomainType = 9
	DOMAINTYPE_CLUSTER                  DomainType = 10
)

// returns the name of the enum given an ordinal number
func (e *DomainType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
Property of the node to be added
*/
type ExpandClusterParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExtraParams *ExtraParam `json:"extraParams,omitempty"`

	NodeParams *NodeParam `json:"nodeParams,omitempty"`
	/**
	  Indicates if node addition can be skipped
	*/
	SkipAddNode *bool `json:"skipAddNode,omitempty"`
	/**
	  Indicates if pre-expand checks can be skipped for node addition
	*/
	SkipPreExpandChecks *bool `json:"skipPreExpandChecks,omitempty"`
}

func NewExpandClusterParams() *ExpandClusterParams {
	p := new(ExpandClusterParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ExpandClusterParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ExpandClusterParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Extra parameters
*/
type ExtraParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Hyperv *HyperVCredAddNode `json:"hyperv,omitempty"`
	/**
	  Indicates whether the node is compute only or not
	*/
	IsComputeOnly *bool `json:"isComputeOnly,omitempty"`
	/**
	  Indicates if node is compatible or not
	*/
	IsNosCompatible *bool `json:"isNosCompatible,omitempty"`
	/**
	  Indicates whether the node is marked to be never schedulable or not.
	*/
	NeverScheduleable *bool `json:"neverScheduleable,omitempty"`
	/**
	  Indicates if node discovery need to be skipped or not
	*/
	SkipDiscovery *bool `json:"skipDiscovery,omitempty"`
	/**
	  Indicates if node imaging needs to be skipped or not
	*/
	SkipImaging *bool `json:"skipImaging,omitempty"`
	/**
	  Target hypervisor
	*/
	TargetHypervisor *string `json:"targetHypervisor,omitempty"`
	/**
	  Indicates if rack awareness needs to be validated or not
	*/
	ValidateRackAwareness *bool `json:"validateRackAwareness,omitempty"`
}

func NewExtraParam() *ExtraParam {
	p := new(ExtraParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ExtraParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ExtraParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	  Maximum fault tolerance desired.
	*/
	DesiredMaxFaultTolerance *int `json:"desiredMaxFaultTolerance,omitempty"`

	DomainAwarenessLevel *DomainAwarenessLevel `json:"domainAwarenessLevel,omitempty"`
}

func NewFaultToleranceState() *FaultToleranceState {
	p := new(FaultToleranceState)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.FaultToleranceState"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.FaultToleranceState"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/host-gpus Get operation
*/
type GetAllHostGpusResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAllHostGpusResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAllHostGpusResponse() *GetAllHostGpusResponse {
	p := new(GetAllHostGpusResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetAllHostGpusResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetAllHostGpusResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAllHostGpusResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAllHostGpusResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAllHostGpusResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/hosts Get operation
*/
type GetAllHostsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAllHostsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAllHostsResponse() *GetAllHostsResponse {
	p := new(GetAllHostsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetAllHostsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetAllHostsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAllHostsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAllHostsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAllHostsResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/host-gpus Get operation
*/
type GetClusterHostGpusResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterHostGpusResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClusterHostGpusResponse() *GetClusterHostGpusResponse {
	p := new(GetClusterHostGpusResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetClusterHostGpusResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetClusterHostGpusResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId} Get operation
*/
type GetClusterResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClusterResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClusterResponse() *GetClusterResponse {
	p := new(GetClusterResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetClusterResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetClusterResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters Get operation
*/
type GetClustersResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetClustersResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetClustersResponse() *GetClustersResponse {
	p := new(GetClustersResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetClustersResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetClustersResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/domain-fault-tolerance-status Get operation
*/
type GetDomainFaultToleranceResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDomainFaultToleranceResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetDomainFaultToleranceResponse() *GetDomainFaultToleranceResponse {
	p := new(GetDomainFaultToleranceResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetDomainFaultToleranceResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetDomainFaultToleranceResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/gpu-profiles Get operation
*/
type GetGpuProfilesResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetGpuProfilesResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetGpuProfilesResponse() *GetGpuProfilesResponse {
	p := new(GetGpuProfilesResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetGpuProfilesResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetGpuProfilesResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-gpus/{hostGpuExtId} Get operation
*/
type GetHostGpuResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHostGpuResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetHostGpuResponse() *GetHostGpuResponse {
	p := new(GetHostGpuResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetHostGpuResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetHostGpuResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/host-gpus Get operation
*/
type GetHostGpusResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHostGpusResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetHostGpusResponse() *GetHostGpusResponse {
	p := new(GetHostGpusResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetHostGpusResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetHostGpusResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId} Get operation
*/
type GetHostResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHostResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetHostResponse() *GetHostResponse {
	p := new(GetHostResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetHostResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetHostResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts Get operation
*/
type GetHostsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHostsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetHostsResponse() *GetHostsResponse {
	p := new(GetHostsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetHostsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetHostsResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/fetch-node-networking-details Post operation
*/
type GetNodeNetworkingTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNodeNetworkingTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNodeNetworkingTaskApiResponse() *GetNodeNetworkingTaskApiResponse {
	p := new(GetNodeNetworkingTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetNodeNetworkingTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetNodeNetworkingTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNodeNetworkingTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNodeNetworkingTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNodeNetworkingTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rackable-units/{rackableUnitExtId} Get operation
*/
type GetRackableUnitResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRackableUnitResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRackableUnitResponse() *GetRackableUnitResponse {
	p := new(GetRackableUnitResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetRackableUnitResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetRackableUnitResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rackable-units Get operation
*/
type GetRackableUnitsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRackableUnitsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRackableUnitsResponse() *GetRackableUnitsResponse {
	p := new(GetRackableUnitsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetRackableUnitsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetRackableUnitsResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId} Get operation
*/
type GetRsyslogServerResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRsyslogServerResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRsyslogServerResponse() *GetRsyslogServerResponse {
	p := new(GetRsyslogServerResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetRsyslogServerResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetRsyslogServerResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetRsyslogServerResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetRsyslogServerResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetRsyslogServerResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers Get operation
*/
type GetRsyslogServersResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetRsyslogServersResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetRsyslogServersResponse() *GetRsyslogServersResponse {
	p := new(GetRsyslogServersResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetRsyslogServersResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetRsyslogServersResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/tasks/{taskExtId}/$actions/fetch-task-response Post operation
*/
type GetSearchResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSearchResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSearchResponse() *GetSearchResponse {
	p := new(GetSearchResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetSearchResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetSearchResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSearchResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSearchResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSearchResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp Get operation
*/
type GetSnmpResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSnmpResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSnmpResponse() *GetSnmpResponse {
	p := new(GetSnmpResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetSnmpResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetSnmpResponse"}
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId} Get operation
*/
type GetSnmpTrapResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSnmpTrapResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSnmpTrapResponse() *GetSnmpTrapResponse {
	p := new(GetSnmpTrapResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetSnmpTrapResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetSnmpTrapResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSnmpTrapResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSnmpTrapResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSnmpTrapResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId} Get operation
*/
type GetSnmpUserResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetSnmpUserResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetSnmpUserResponse() *GetSnmpUserResponse {
	p := new(GetSnmpUserResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.GetSnmpUserResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GetSnmpUserResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetSnmpUserResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetSnmpUserResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetSnmpUserResponseData()
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
	  Maximum resolution per display heads
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GpuConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
GPU mode
*/
type GpuMode int

const (
	GPUMODE_UNKNOWN              GpuMode = 0
	GPUMODE_REDACTED             GpuMode = 1
	GPUMODE_UNUSED               GpuMode = 2
	GPUMODE_USED_FOR_PASSTHROUGH GpuMode = 3
	GPUMODE_USED_FOR_VIRTUAL     GpuMode = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuMode) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.GpuProfile"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
GPU type
*/
type GpuType int

const (
	GPUTYPE_UNKNOWN               GpuType = 0
	GPUTYPE_REDACTED              GpuType = 1
	GPUTYPE_PASS_THROUGH_GRAPHICS GpuType = 2
	GPUTYPE_PASS_THROUGH_COMPUTE  GpuType = 3
	GPUTYPE_VIRTUAL               GpuType = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	Links []import3.ApiLink `json:"links,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HostEntity"}
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

	Cluster *ClusterReference `json:"cluster,omitempty"`

	Config *GpuConfig `json:"config,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HostGpuEntity"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Host rename parameters
*/
type HostNameParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the host
	*/
	Name *string `json:"name"`
}

func (p *HostNameParam) MarshalJSON() ([]byte, error) {
	type HostNameParamProxy HostNameParam
	return json.Marshal(struct {
		*HostNameParamProxy
		Name *string `json:"name,omitempty"`
	}{
		HostNameParamProxy: (*HostNameParamProxy)(p),
		Name:               p.Name,
	})
}

func NewHostNameParam() *HostNameParam {
	p := new(HostNameParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HostNameParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HostNameParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/hosts/{hostExtId}/$actions/rename-host Post operation
*/
type HostRenameResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfHostRenameResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewHostRenameResponse() *HostRenameResponse {
	p := new(HostRenameResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HostRenameResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HostRenameResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HostRenameResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *HostRenameResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfHostRenameResponseData()
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
Type of the host
*/
type HostTypeEnum int

const (
	HOSTTYPEENUM_UNKNOWN         HostTypeEnum = 0
	HOSTTYPEENUM_REDACTED        HostTypeEnum = 1
	HOSTTYPEENUM_HYPER_CONVERGED HostTypeEnum = 2
	HOSTTYPEENUM_COMPUTE_ONLY    HostTypeEnum = 3
)

// returns the name of the enum given an ordinal number
func (e *HostTypeEnum) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
HyperV Credentials
*/
type HyperVCredAddNode struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DomainDetails *UserInfo `json:"domainDetails,omitempty"`

	FailoverClusterDetails *UserInfo `json:"failoverClusterDetails,omitempty"`
}

func NewHyperVCredAddNode() *HyperVCredAddNode {
	p := new(HyperVCredAddNode)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HyperVCredAddNode"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HyperVCredAddNode"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Map containing key as hypervisor type and value as md5sum of ISO
*/
type HypervisorIsoMap struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Md5sum of ISO
	*/
	Md5sum *string `json:"md5sum,omitempty"`

	Type *HypervisorType `json:"type,omitempty"`
}

func NewHypervisorIsoMap() *HypervisorIsoMap {
	p := new(HypervisorIsoMap)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HypervisorIsoMap"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorIsoMap"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Hypervisor details
*/
type HypervisorReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AcropolisConnectionState *AcropolisConnectionState `json:"acropolisConnectionState,omitempty"`

	ExternalAddress *import4.IPAddress `json:"externalAddress,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Hypervisor state
*/
type HypervisorState int

const (
	HYPERVISORSTATE_UNKNOWN                                    HypervisorState = 0
	HYPERVISORSTATE_REDACTED                                   HypervisorState = 1
	HYPERVISORSTATE_ACROPOLIS_NORMAL                           HypervisorState = 2
	HYPERVISORSTATE_ENTERING_MAINTENANCE_MODE                  HypervisorState = 3
	HYPERVISORSTATE_ENTERED_MAINTENANCE_MODE                   HypervisorState = 4
	HYPERVISORSTATE_RESERVED_FOR_HA_FAILOVER                   HypervisorState = 5
	HYPERVISORSTATE_ENTERING_MAINTENANCE_MODE_FROM_HA_FAILOVER HypervisorState = 6
	HYPERVISORSTATE_RESERVING_FOR_HA_FAILOVER                  HypervisorState = 7
	HYPERVISORSTATE_HA_FAILOVER_SOURCE                         HypervisorState = 8
	HYPERVISORSTATE_HA_FAILOVER_TARGET                         HypervisorState = 9
	HYPERVISORSTATE_HA_HEALING_SOURCE                          HypervisorState = 10
	HYPERVISORSTATE_HA_HEALING_TARGET                          HypervisorState = 11
)

// returns the name of the enum given an ordinal number
func (e *HypervisorState) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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

const (
	HYPERVISORTYPE_UNKNOWN  HypervisorType = 0
	HYPERVISORTYPE_REDACTED HypervisorType = 1
	HYPERVISORTYPE_AHV      HypervisorType = 2
	HYPERVISORTYPE_ESX      HypervisorType = 3
	HYPERVISORTYPE_HYPERV   HypervisorType = 4
	HYPERVISORTYPE_XEN      HypervisorType = 5
	HYPERVISORTYPE_NULLTYPE HypervisorType = 6
)

// returns the name of the enum given an ordinal number
func (e *HypervisorType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
Hypervisor upload required information
*/
type HypervisorUploadInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Error message
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/**
	  Node list containing upload information
	*/
	UploadInfoNodeList []UploadInfoNodeItem `json:"uploadInfoNodeList,omitempty"`
}

func NewHypervisorUploadInfo() *HypervisorUploadInfo {
	p := new(HypervisorUploadInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HypervisorUploadInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorUploadInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Individual node item details for checking whether hypervisor ISO upload is required or not
*/
type HypervisorUploadNodeListItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit Id in which node resides
	*/
	BlockId *string `json:"blockId,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Host version of the node
	*/
	HypervisorVersion *string `json:"hypervisorVersion,omitempty"`
	/**
	  Indicates whether the node is light compute or not
	*/
	IsLightCompute *bool `json:"isLightCompute,omitempty"`
	/**
	  Indicates if node is minimum compute or not
	*/
	MinimumComputeNode *bool `json:"minimumComputeNode,omitempty"`
	/**
	  Rackable unit model type
	*/
	Model *string `json:"model,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
	/**
	  NOS software version of a node
	*/
	NosVersion *string `json:"nosVersion,omitempty"`
	/**
	  Indicates whether the hypervisor is robo mixed or not
	*/
	RoboMixedHypervisor *bool `json:"roboMixedHypervisor,omitempty"`
}

func NewHypervisorUploadNodeListItem() *HypervisorUploadNodeListItem {
	p := new(HypervisorUploadNodeListItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HypervisorUploadNodeListItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorUploadNodeListItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Parameters to get information on whether hypervisor ISO upload is required or not
*/
type HypervisorUploadParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of node details for checking whether hypervisor ISO upload is required or not
	*/
	NodeList []HypervisorUploadNodeListItem `json:"nodeList,omitempty"`
}

func NewHypervisorUploadParam() *HypervisorUploadParam {
	p := new(HypervisorUploadParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HypervisorUploadParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorUploadParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/check-hypervisor-requirements Post operation
*/
type HypervisorUplpadRequiredTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfHypervisorUplpadRequiredTaskResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewHypervisorUplpadRequiredTaskResponse() *HypervisorUplpadRequiredTaskResponse {
	p := new(HypervisorUplpadRequiredTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.HypervisorUplpadRequiredTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.HypervisorUplpadRequiredTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HypervisorUplpadRequiredTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *HypervisorUplpadRequiredTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfHypervisorUplpadRequiredTaskResponseData()
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
IPMI reference
*/
type IpmiReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Ip *import4.IPAddress `json:"ip,omitempty"`
	/**
	  IPMI username
	*/
	Username *string `json:"username,omitempty"`
}

func NewIpmiReference() *IpmiReference {
	p := new(IpmiReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.IpmiReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.IpmiReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/validate-bundle Post operation
*/
type IsBundleCompatibleTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfIsBundleCompatibleTaskResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewIsBundleCompatibleTaskResponse() *IsBundleCompatibleTaskResponse {
	p := new(IsBundleCompatibleTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.IsBundleCompatibleTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.IsBundleCompatibleTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *IsBundleCompatibleTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *IsBundleCompatibleTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfIsBundleCompatibleTaskResponseData()
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.KeyManagementDeviceToCertStatus"}
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

	Ip *import4.IPAddress `json:"ip,omitempty"`
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ManagementServerRef"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Management server type
*/
type ManagementServerType int

const (
	MANAGEMENTSERVERTYPE_UNKNOWN  ManagementServerType = 0
	MANAGEMENTSERVERTYPE_REDACTED ManagementServerType = 1
	MANAGEMENTSERVERTYPE_VCENTER  ManagementServerType = 2
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
Interface name and mac address
*/
type NameMacRef struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Mac address
	*/
	Mac *string `json:"mac,omitempty"`
	/**
	  Interface name
	*/
	Name *string `json:"name,omitempty"`
}

func NewNameMacRef() *NameMacRef {
	p := new(NameMacRef)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NameMacRef"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NameMacRef"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Name and network information
*/
type NameNetworkRef struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Interface name
	*/
	Name *string `json:"name,omitempty"`
	/**
	  List of networks for interface
	*/
	Networks []string `json:"networks,omitempty"`
}

func NewNameNetworkRef() *NameNetworkRef {
	p := new(NameNetworkRef)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NameNetworkRef"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NameNetworkRef"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Network information of HCI and SO nodes
*/
type NetworkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Network information of HCI nodes
	*/
	Hci []NameNetworkRef `json:"hci,omitempty"`
	/**
	  Network information of SO nodes
	*/
	So []NameNetworkRef `json:"so,omitempty"`
}

func NewNetworkInfo() *NetworkInfo {
	p := new(NetworkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NetworkInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NetworkInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Request type and networking details for nodes
*/
type NodeDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Node specific details required to fetch node networking information
	*/
	NodeList []NodeListNetworkingDetails `json:"nodeList"`
	/**
	  Type of request, either it can be expand_cluster or npe
	*/
	RequestType *string `json:"requestType,omitempty"`
}

func (p *NodeDetails) MarshalJSON() ([]byte, error) {
	type NodeDetailsProxy NodeDetails
	return json.Marshal(struct {
		*NodeDetailsProxy
		NodeList []NodeListNetworkingDetails `json:"nodeList,omitempty"`
	}{
		NodeDetailsProxy: (*NodeDetailsProxy)(p),
		NodeList:         p.NodeList,
	})
}

func NewNodeDetails() *NodeDetails {
	p := new(NodeDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Discover unconfigured node details
*/
type NodeDiscoveryParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AddressType *AddressType `json:"addressType,omitempty"`
	/**
	  Interface name that is used for packet broadcasting
	*/
	InterfaceFilterList []string `json:"interfaceFilterList,omitempty"`
	/**
	  IP addresses of the unconfigured nodes
	*/
	IpFilterList []import4.IPAddress `json:"ipFilterList,omitempty"`
	/**
	  Indicates if the discovery is manual or not
	*/
	IsManualDiscovery *bool `json:"isManualDiscovery,omitempty"`
	/**
	  Timeout for discovering nodes
	*/
	Timeout *int64 `json:"timeout,omitempty"`
	/**
	  Unconfigured node UUIDs
	*/
	UuidFilterList []string `json:"uuidFilterList,omitempty"`
}

func NewNodeDiscoveryParams() *NodeDiscoveryParams {
	p := new(NodeDiscoveryParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeDiscoveryParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeDiscoveryParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Node item containing attributes of node
*/
type NodeInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit serial name
	*/
	BlockId *string `json:"blockId,omitempty"`
	/**
	  Current network interface of a node
	*/
	CurrentNetworkInterface *string `json:"currentNetworkInterface,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`
	/**
	  List of objects containing digital_certificate_base64 and key_management_server_uuid fields for key management server
	*/
	DigitalCertificateMapList []DigitalCertificateMapReference `json:"digitalCertificateMapList,omitempty"`
	/**
	  Name of the host
	*/
	HypervisorHostname *string `json:"hypervisorHostname,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Host version of the node
	*/
	HypervisorVersion *string `json:"hypervisorVersion,omitempty"`

	IpmiIp *import4.IPAddress `json:"ipmiIp,omitempty"`
	/**
	  Indicates whether the node is light compute or not
	*/
	IsLightCompute *bool `json:"isLightCompute,omitempty"`
	/**
	  Rackable unit model name
	*/
	Model *string `json:"model,omitempty"`
	/**
	  Position of a node in a rackable unit
	*/
	NodePosition *string `json:"nodePosition,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
	/**
	  NOS software version of a node
	*/
	NosVersion *string `json:"nosVersion,omitempty"`
	/**
	  Indicates whether the hypervisor is robo mixed or not
	*/
	RoboMixedHypervisor *bool `json:"roboMixedHypervisor,omitempty"`
}

func NewNodeInfo() *NodeInfo {
	p := new(NodeInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Node item containing attributes of node
*/
type NodeItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit serial name
	*/
	BlockId *string `json:"blockId,omitempty"`
	/**
	  Current network interface of a node
	*/
	CurrentNetworkInterface *string `json:"currentNetworkInterface,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`
	/**
	  List of objects containing digital_certificate_base64 and key_management_server_uuid fields for key management server
	*/
	DigitalCertificateMapList []DigitalCertificateMapReference `json:"digitalCertificateMapList,omitempty"`
	/**
	  Name of the host
	*/
	HypervisorHostname *string `json:"hypervisorHostname,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Host version of the node
	*/
	HypervisorVersion *string `json:"hypervisorVersion,omitempty"`

	IpmiIp *import4.IPAddress `json:"ipmiIp,omitempty"`
	/**
	  Indicates whether the node is light compute or not
	*/
	IsLightCompute *bool `json:"isLightCompute,omitempty"`
	/**
	  Rackable unit model name
	*/
	Model *string `json:"model,omitempty"`
	/**
	  Active and standby uplink information of the target nodes
	*/
	Networks []UplinkNetworkItem `json:"networks,omitempty"`
	/**
	  Position of a node in a rackable unit
	*/
	NodePosition *string `json:"nodePosition,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
	/**
	  NOS software version of a node
	*/
	NosVersion *string `json:"nosVersion,omitempty"`
	/**
	  Indicates whether the hypervisor is robo mixed or not
	*/
	RoboMixedHypervisor *bool `json:"roboMixedHypervisor,omitempty"`
}

func NewNodeItem() *NodeItem {
	p := new(NodeItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
List of nodes in a cluster
*/
type NodeListItemReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ControllerVmIp *import4.IPAddress `json:"controllerVmIp,omitempty"`

	HostIp *import4.IPAddress `json:"hostIp,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
}

func NewNodeListItemReference() *NodeListItemReference {
	p := new(NodeListItemReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeListItemReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeListItemReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Node specific details required to fetch node networking information
*/
type NodeListNetworkingDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Rackable unit Id in which node resides
	*/
	BlockId *string `json:"blockId,omitempty"`
	/**
	  Current network interface of a node
	*/
	CurrentNetworkInterface *string `json:"currentNetworkInterface,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`
	/**
	  List of objects containing digital_certificate_base64 and key_management_server_uuid fields for key management server
	*/
	DigitalCertificateMapList []DigitalCertificateMapReference `json:"digitalCertificateMapList,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Host version of the node
	*/
	HypervisorVersion *string `json:"hypervisorVersion,omitempty"`

	IpmiIp *import4.IPAddress `json:"ipmiIp,omitempty"`
	/**
	  Indicates whether the node is compute only or not
	*/
	IsComputeOnly *bool `json:"isComputeOnly,omitempty"`
	/**
	  Indicates whether the node is light compute or not
	*/
	IsLightCompute *bool `json:"isLightCompute,omitempty"`
	/**
	  Rackable unit model name
	*/
	Model *string `json:"model,omitempty"`
	/**
	  Position of a node in a rackable unit
	*/
	NodePosition *string `json:"nodePosition,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
	/**
	  NOS software version of a node
	*/
	NosVersion *string `json:"nosVersion,omitempty"`
	/**
	  Indicates whether the hypervisor is robo mixed or not
	*/
	RoboMixedHypervisor *bool `json:"roboMixedHypervisor,omitempty"`
}

func NewNodeListNetworkingDetails() *NodeListNetworkingDetails {
	p := new(NodeListNetworkingDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeListNetworkingDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeListNetworkingDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Network details of nodes
*/
type NodeNetworkingDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkInfo *NetworkInfo `json:"networkInfo,omitempty"`
	/**
	  List of uplinks information for each CVM IP
	*/
	Uplinks []UplinkInfo `json:"uplinks,omitempty"`
	/**
	  List of warning messages
	*/
	Warnings []string `json:"warnings,omitempty"`
}

func NewNodeNetworkingDetails() *NodeNetworkingDetails {
	p := new(NodeNetworkingDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeNetworkingDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeNetworkingDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Parameters of the node to be added
*/
type NodeParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Block list of a cluster
	*/
	BlockList []BlockItem `json:"blockList,omitempty"`

	BundleInfo *BundleInfo `json:"bundleInfo,omitempty"`
	/**
	  List of compute only nodes
	*/
	ComputeNodeList []ComputeNodeItem `json:"computeNodeList,omitempty"`
	/**
	  Hyperv SKU
	*/
	HypervSku *string `json:"hypervSku,omitempty"`
	/**
	  Hypervisor type to md5sum map
	*/
	HypervisorIsos []HypervisorIsoMap `json:"hypervisorIsos,omitempty"`
	/**
	  List of nodes in a cluster
	*/
	NodeList []NodeItem `json:"nodeList,omitempty"`
	/**
	  Indicates if the host networking needs to be skipped or not
	*/
	SkipHostNetworking *bool `json:"skipHostNetworking,omitempty"`
}

func NewNodeParam() *NodeParam {
	p := new(NodeParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeParam"}
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Extra parameters
*/
type NodeRemovalExtraParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates if add node check need to be skip or not
	*/
	SkipAddCheck *bool `json:"skipAddCheck,omitempty"`
	/**
	  Indicates if space check needs to be skip or not
	*/
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	/**
	  Indicates if upgrade check needs to be skip or not
	*/
	SkipUpgradeCheck *bool `json:"skipUpgradeCheck,omitempty"`
}

func NewNodeRemovalExtraParam() *NodeRemovalExtraParam {
	p := new(NodeRemovalExtraParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeRemovalExtraParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeRemovalExtraParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Parameters to remove nodes from cluster
*/
type NodeRemovalParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExtraParams *NodeRemovalExtraParam `json:"extraParams,omitempty"`
	/**
	  List of node UUIDs to remove
	*/
	NodeUuids []string `json:"nodeUuids,omitempty"`
	/**
	  Indicates if prechecks can be skipped for node removal
	*/
	SkipPrechecks *bool `json:"skipPrechecks,omitempty"`
	/**
	  Indicates if node removal can be skipped
	*/
	SkipRemove *bool `json:"skipRemove,omitempty"`
}

func NewNodeRemovalParams() *NodeRemovalParams {
	p := new(NodeRemovalParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NodeRemovalParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.NodeRemovalParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Cluster operation mode
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

// returns the name of the enum given an ordinal number
func (e *OperationMode) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	Key *string `json:"key"`
	/**
	  Ssh key name
	*/
	Name *string `json:"name"`
}

func (p *PublicKey) MarshalJSON() ([]byte, error) {
	type PublicKeyProxy PublicKey
	return json.Marshal(struct {
		*PublicKeyProxy
		Key  *string `json:"key,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		PublicKeyProxy: (*PublicKeyProxy)(p),
		Key:            p.Key,
		Name:           p.Name,
	})
}

func NewPublicKey() *PublicKey {
	p := new(PublicKey)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.PublicKey"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.PublicKey"}
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
	  Rack Id
	*/
	Id *int64 `json:"id,omitempty"`
	/**
	  Rack UUID
	*/
	Uuid *string `json:"uuid,omitempty"`
}

func NewRackReference() *RackReference {
	p := new(RackReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RackReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RackReference"}
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
	Links []import3.ApiLink `json:"links,omitempty"`

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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RackableUnit"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Rackable unit model type
*/
type RackableUnitModel int

const (
	RACKABLEUNITMODEL_UNKNOWN   RackableUnitModel = 0
	RACKABLEUNITMODEL_REDACTED  RackableUnitModel = 1
	RACKABLEUNITMODEL_DESKTOP   RackableUnitModel = 2
	RACKABLEUNITMODEL_NX2000    RackableUnitModel = 3
	RACKABLEUNITMODEL_NX3000    RackableUnitModel = 4
	RACKABLEUNITMODEL_NX3050    RackableUnitModel = 5
	RACKABLEUNITMODEL_NX6050    RackableUnitModel = 6
	RACKABLEUNITMODEL_NX6070    RackableUnitModel = 7
	RACKABLEUNITMODEL_NX1050    RackableUnitModel = 8
	RACKABLEUNITMODEL_NX3060    RackableUnitModel = 9
	RACKABLEUNITMODEL_NX6060    RackableUnitModel = 10
	RACKABLEUNITMODEL_NX6080    RackableUnitModel = 11
	RACKABLEUNITMODEL_NX6020    RackableUnitModel = 12
	RACKABLEUNITMODEL_NX7110    RackableUnitModel = 13
	RACKABLEUNITMODEL_NX1020    RackableUnitModel = 14
	RACKABLEUNITMODEL_NX9040    RackableUnitModel = 15
	RACKABLEUNITMODEL_USELAYOUT RackableUnitModel = 16
	RACKABLEUNITMODEL_NULLVALUE RackableUnitModel = 17
)

// returns the name of the enum given an ordinal number
func (e *RackableUnitModel) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RackableUnitNode"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/remove-node Post operation
*/
type RemoveNodeTaskResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoveNodeTaskResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoveNodeTaskResponse() *RemoveNodeTaskResponse {
	p := new(RemoveNodeTaskResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RemoveNodeTaskResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RemoveNodeTaskResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoveNodeTaskResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoveNodeTaskResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoveNodeTaskResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/remove-transports Post operation
*/
type RemoveSnmpTransportsTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoveSnmpTransportsTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoveSnmpTransportsTaskApiResponse() *RemoveSnmpTransportsTaskApiResponse {
	p := new(RemoveSnmpTransportsTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RemoveSnmpTransportsTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RemoveSnmpTransportsTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoveSnmpTransportsTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoveSnmpTransportsTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoveSnmpTransportsTaskApiResponseData()
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
RSYSLOG Module information
*/
type RsyslogModuleItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	LogSeverityLevel *RsyslogModuleLogSeverityLevel `json:"logSeverityLevel"`
	/**
	  Option to log, monitor/output files of a module
	*/
	Monitor *bool `json:"monitor"`

	Name *RsyslogModuleName `json:"name"`
}

func (p *RsyslogModuleItem) MarshalJSON() ([]byte, error) {
	type RsyslogModuleItemProxy RsyslogModuleItem
	return json.Marshal(struct {
		*RsyslogModuleItemProxy
		LogSeverityLevel *RsyslogModuleLogSeverityLevel `json:"logSeverityLevel,omitempty"`
		Monitor          *bool                          `json:"monitor,omitempty"`
		Name             *RsyslogModuleName             `json:"name,omitempty"`
	}{
		RsyslogModuleItemProxy: (*RsyslogModuleItemProxy)(p),
		LogSeverityLevel:       p.LogSeverityLevel,
		Monitor:                p.Monitor,
		Name:                   p.Name,
	})
}

func NewRsyslogModuleItem() *RsyslogModuleItem {
	p := new(RsyslogModuleItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RsyslogModuleItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RsyslogModuleItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
RSYSLOG module log severity level
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

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
RSYSLOG module name
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
)

// returns the name of the enum given an ordinal number
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
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
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

const (
	RSYSLOGNETWORKPROTOCOL_UNKNOWN  RsyslogNetworkProtocol = 0
	RSYSLOGNETWORKPROTOCOL_REDACTED RsyslogNetworkProtocol = 1
	RSYSLOGNETWORKPROTOCOL_UDP      RsyslogNetworkProtocol = 2
	RSYSLOGNETWORKPROTOCOL_TCP      RsyslogNetworkProtocol = 3
	RSYSLOGNETWORKPROTOCOL_RELP     RsyslogNetworkProtocol = 4
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

type RsyslogServer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	IpAddress *import4.IPAddress `json:"ipAddress"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/**
	  List of modules registered to RSYSLOG server
	*/
	Modules []RsyslogModuleItem `json:"modules,omitempty"`

	NetworkProtocol *RsyslogNetworkProtocol `json:"networkProtocol"`
	/**
	  RSYSLOG server port
	*/
	Port *int `json:"port"`
	/**
	  RSYSLOG server name
	*/
	ServerName *string `json:"serverName"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *RsyslogServer) MarshalJSON() ([]byte, error) {
	type RsyslogServerProxy RsyslogServer
	return json.Marshal(struct {
		*RsyslogServerProxy
		IpAddress       *import4.IPAddress      `json:"ipAddress,omitempty"`
		NetworkProtocol *RsyslogNetworkProtocol `json:"networkProtocol,omitempty"`
		Port            *int                    `json:"port,omitempty"`
		ServerName      *string                 `json:"serverName,omitempty"`
	}{
		RsyslogServerProxy: (*RsyslogServerProxy)(p),
		IpAddress:          p.IpAddress,
		NetworkProtocol:    p.NetworkProtocol,
		Port:               p.Port,
		ServerName:         p.ServerName,
	})
}

func NewRsyslogServer() *RsyslogServer {
	p := new(RsyslogServer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.RsyslogServer"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.RsyslogServer"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Search parameters
*/
type SearchParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	SearchType *SearchType `json:"searchType"`
}

func (p *SearchParams) MarshalJSON() ([]byte, error) {
	type SearchParamsProxy SearchParams
	return json.Marshal(struct {
		*SearchParamsProxy
		SearchType *SearchType `json:"searchType,omitempty"`
	}{
		SearchParamsProxy: (*SearchParamsProxy)(p),
		SearchType:        p.SearchType,
	})
}

func NewSearchParams() *SearchParams {
	p := new(SearchParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SearchParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SearchParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Task Response which is one of node-discovery, networking-details, hypervisor-upload information
*/
type SearchResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ResponseItemDiscriminator_ *string `json:"$responseItemDiscriminator,omitempty"`

	Response *OneOfSearchResponseResponse `json:"response,omitempty"`

	SearchType *SearchType `json:"searchType,omitempty"`
}

func NewSearchResponse() *SearchResponse {
	p := new(SearchResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SearchResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SearchResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SearchResponse) GetResponse() interface{} {
	if nil == p.Response {
		return nil
	}
	return p.Response.GetValue()
}

func (p *SearchResponse) SetResponse(v interface{}) error {
	if nil == p.Response {
		p.Response = NewOneOfSearchResponseResponse()
	}
	e := p.Response.SetValue(v)
	if nil == e {
		if nil == p.ResponseItemDiscriminator_ {
			p.ResponseItemDiscriminator_ = new(string)
		}
		*p.ResponseItemDiscriminator_ = *p.Response.Discriminator
	}
	return e
}

type SearchType int

const (
	SEARCHTYPE_UNKNOWN                SearchType = 0
	SEARCHTYPE_REDACTED               SearchType = 1
	SEARCHTYPE_UNCONFIGURED_NODES     SearchType = 2
	SEARCHTYPE_NETWORKING_DETAILS     SearchType = 3
	SEARCHTYPE_HYPERVISOR_UPLOAD_INFO SearchType = 4
)

// returns the name of the enum given an ordinal number
func (e *SearchType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNCONFIGURED_NODES",
		"NETWORKING_DETAILS",
		"HYPERVISOR_UPLOAD_INFO",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *SearchType) index(name string) SearchType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UNCONFIGURED_NODES",
		"NETWORKING_DETAILS",
		"HYPERVISOR_UPLOAD_INFO",
	}
	for idx := range names {
		if names[idx] == name {
			return SearchType(idx)
		}
	}
	return SEARCHTYPE_UNKNOWN
}

func (e *SearchType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SearchType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SearchType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SearchType) Ref() *SearchType {
	return &e
}

/**
SMTP network details
*/
type SmtpNetwork struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	IpAddress *import4.IPAddress `json:"ipAddress"`
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

func (p *SmtpNetwork) MarshalJSON() ([]byte, error) {
	type SmtpNetworkProxy SmtpNetwork
	return json.Marshal(struct {
		*SmtpNetworkProxy
		IpAddress *import4.IPAddress `json:"ipAddress,omitempty"`
	}{
		SmtpNetworkProxy: (*SmtpNetworkProxy)(p),
		IpAddress:        p.IpAddress,
	})
}

func NewSmtpNetwork() *SmtpNetwork {
	p := new(SmtpNetwork)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SmtpNetwork"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SmtpNetwork"}
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
	  SMTP email address
	*/
	EmailAddress *string `json:"emailAddress"`

	Server *SmtpNetwork `json:"server"`

	Type *SmtpType `json:"type,omitempty"`
}

func (p *SmtpServerRef) MarshalJSON() ([]byte, error) {
	type SmtpServerRefProxy SmtpServerRef
	return json.Marshal(struct {
		*SmtpServerRefProxy
		EmailAddress *string      `json:"emailAddress,omitempty"`
		Server       *SmtpNetwork `json:"server,omitempty"`
	}{
		SmtpServerRefProxy: (*SmtpServerRefProxy)(p),
		EmailAddress:       p.EmailAddress,
		Server:             p.Server,
	})
}

func NewSmtpServerRef() *SmtpServerRef {
	p := new(SmtpServerRef)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SmtpServerRef"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SmtpServerRef"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Type of SMTP server
*/
type SmtpType int

const (
	SMTPTYPE_UNKNOWN  SmtpType = 0
	SMTPTYPE_REDACTED SmtpType = 1
	SMTPTYPE_PLAIN    SmtpType = 2
	SMTPTYPE_STARTTLS SmtpType = 3
	SMTPTYPE_SSL      SmtpType = 4
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
SNMP user authentication type
*/
type SnmpAuthType int

const (
	SNMPAUTHTYPE_UNKNOWN  SnmpAuthType = 0
	SNMPAUTHTYPE_REDACTED SnmpAuthType = 1
	SNMPAUTHTYPE_MD5      SnmpAuthType = 2
	SNMPAUTHTYPE_SHA      SnmpAuthType = 3
)

// returns the name of the enum given an ordinal number
func (e *SnmpAuthType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SnmpConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
SNMP user encryption type
*/
type SnmpPrivType int

const (
	SNMPPRIVTYPE_UNKNOWN  SnmpPrivType = 0
	SNMPPRIVTYPE_REDACTED SnmpPrivType = 1
	SNMPPRIVTYPE_DES      SnmpPrivType = 2
	SNMPPRIVTYPE_AES      SnmpPrivType = 3
)

// returns the name of the enum given an ordinal number
func (e *SnmpPrivType) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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

const (
	SNMPPROTOCOL_UNKNOWN  SnmpProtocol = 0
	SNMPPROTOCOL_REDACTED SnmpProtocol = 1
	SNMPPROTOCOL_UDP      SnmpProtocol = 2
	SNMPPROTOCOL_UDP6     SnmpProtocol = 3
	SNMPPROTOCOL_TCP      SnmpProtocol = 4
	SNMPPROTOCOL_TCP6     SnmpProtocol = 5
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
SNMP status
*/
type SnmpStatusParam struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  SNMP user information
	*/
	Status *bool `json:"status"`
}

func (p *SnmpStatusParam) MarshalJSON() ([]byte, error) {
	type SnmpStatusParamProxy SnmpStatusParam
	return json.Marshal(struct {
		*SnmpStatusParamProxy
		Status *bool `json:"status,omitempty"`
	}{
		SnmpStatusParamProxy: (*SnmpStatusParamProxy)(p),
		Status:               p.Status,
	})
}

func NewSnmpStatusParam() *SnmpStatusParam {
	p := new(SnmpStatusParam)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpStatusParam"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SnmpStatusParam"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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
	Port *int `json:"port"`

	Protocol *SnmpProtocol `json:"protocol"`
}

func (p *SnmpTransport) MarshalJSON() ([]byte, error) {
	type SnmpTransportProxy SnmpTransport
	return json.Marshal(struct {
		*SnmpTransportProxy
		Port     *int          `json:"port,omitempty"`
		Protocol *SnmpProtocol `json:"protocol,omitempty"`
	}{
		SnmpTransportProxy: (*SnmpTransportProxy)(p),
		Port:               p.Port,
		Protocol:           p.Protocol,
	})
}

func NewSnmpTransport() *SnmpTransport {
	p := new(SnmpTransport)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpTransport"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SnmpTransport"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type SnmpTrap struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import4.IPAddress `json:"address"`
	/**
	  Community string(plaintext) for SNMP version 2.0
	*/
	CommunityString *string `json:"communityString,omitempty"`
	/**
	  SNMP engine Id
	*/
	EngineId *string `json:"engineId,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  SNMP information status
	*/
	Inform *bool `json:"inform,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
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
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  SNMP user name
	*/
	Username *string `json:"username,omitempty"`

	Version *SnmpTrapVersion `json:"version"`
}

func (p *SnmpTrap) MarshalJSON() ([]byte, error) {
	type SnmpTrapProxy SnmpTrap
	return json.Marshal(struct {
		*SnmpTrapProxy
		Address *import4.IPAddress `json:"address,omitempty"`
		Version *SnmpTrapVersion   `json:"version,omitempty"`
	}{
		SnmpTrapProxy: (*SnmpTrapProxy)(p),
		Address:       p.Address,
		Version:       p.Version,
	})
}

func NewSnmpTrap() *SnmpTrap {
	p := new(SnmpTrap)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpTrap"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SnmpTrap"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
SNMP version
*/
type SnmpTrapVersion int

const (
	SNMPTRAPVERSION_UNKNOWN  SnmpTrapVersion = 0
	SNMPTRAPVERSION_REDACTED SnmpTrapVersion = 1
	SNMPTRAPVERSION_V2       SnmpTrapVersion = 2
	SNMPTRAPVERSION_V3       SnmpTrapVersion = 3
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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
	AuthKey *string `json:"authKey"`

	AuthType *SnmpAuthType `json:"authType"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/**
	  SNMP user encryption key
	*/
	PrivKey *string `json:"privKey,omitempty"`

	PrivType *SnmpPrivType `json:"privType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/**
	  SNMP user name
	*/
	Username *string `json:"username"`
}

func (p *SnmpUser) MarshalJSON() ([]byte, error) {
	type SnmpUserProxy SnmpUser
	return json.Marshal(struct {
		*SnmpUserProxy
		AuthKey  *string       `json:"authKey,omitempty"`
		AuthType *SnmpAuthType `json:"authType,omitempty"`
		Username *string       `json:"username,omitempty"`
	}{
		SnmpUserProxy: (*SnmpUserProxy)(p),
		AuthKey:       p.AuthKey,
		AuthType:      p.AuthType,
		Username:      p.Username,
	})
}

func NewSnmpUser() *SnmpUser {
	p := new(SnmpUser)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.SnmpUser"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SnmpUser"}
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.SoftwareMapReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Software type
*/
type SoftwareTypeRef int

const (
	SOFTWARETYPEREF_UNKNOWN       SoftwareTypeRef = 0
	SOFTWARETYPEREF_REDACTED      SoftwareTypeRef = 1
	SOFTWARETYPEREF_NOS           SoftwareTypeRef = 2
	SOFTWARETYPEREF_NCC           SoftwareTypeRef = 3
	SOFTWARETYPEREF_PRISM_CENTRAL SoftwareTypeRef = 4
)

// returns the name of the enum given an ordinal number
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

// returns the enum type given a string value
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

/**
Returned list of statistics data.
*/
type StatsData struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Stats query attribute name.
	*/
	Attribute *string `json:"attribute"`
	/**
	  Returned list of statistics data.
	*/
	StatsData []string `json:"statsData,omitempty"`
}

func (p *StatsData) MarshalJSON() ([]byte, error) {
	type StatsDataProxy StatsData
	return json.Marshal(struct {
		*StatsDataProxy
		Attribute *string `json:"attribute,omitempty"`
	}{
		StatsDataProxy: (*StatsDataProxy)(p),
		Attribute:      p.Attribute,
	})
}

func NewStatsData() *StatsData {
	p := new(StatsData)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StatsData"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.StatsData"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Parameters of statistics query.
*/
type StatsQueryParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of queried attributes. All attributes are returned if this parameter is absent.
	*/
	Attributes []string `json:"attributes,omitempty"`
	/**
	  End time in milliseconds of queried statistics data.
	*/
	EndTime *int64 `json:"endTime,omitempty"`
	/**
	  Stats query sampling interval in seconds.
	*/
	SamplingInterval *int64 `json:"samplingInterval,omitempty"`
	/**
	  Start time in milliseconds of queried statistics data.
	*/
	StartTime *int64 `json:"startTime,omitempty"`

	StatsType *StatsType `json:"statsType,omitempty"`
}

func NewStatsQueryParams() *StatsQueryParams {
	p := new(StatsQueryParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StatsQueryParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.StatsQueryParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Response of statistics query.
*/
type StatsQueryResp struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  UUID of queried entity.
	*/
	EntityUuid *string `json:"entityUuid,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/**
	  Returned list of statistics data.
	*/
	StatsData []StatsData `json:"statsData,omitempty"`

	StatsType *StatsType `json:"statsType,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewStatsQueryResp() *StatsQueryResp {
	p := new(StatsQueryResp)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.StatsQueryResp"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.StatsQueryResp"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type StatsType int

const (
	STATSTYPE_UNKNOWN  StatsType = 0
	STATSTYPE_REDACTED StatsType = 1
	STATSTYPE_AVG      StatsType = 2
	STATSTYPE_LAST     StatsType = 3
)

// returns the name of the enum given an ordinal number
func (e *StatsType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVG",
		"LAST",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *StatsType) index(name string) StatsType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"AVG",
		"LAST",
	}
	for idx := range names {
		if names[idx] == name {
			return StatsType(idx)
		}
	}
	return STATSTYPE_UNKNOWN
}

func (e *StatsType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for StatsType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *StatsType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e StatsType) Ref() *StatsType {
	return &e
}

/**
Disk storage Tier type
*/
type StorageTierReference int

const (
	STORAGETIERREFERENCE_UNKNOWN  StorageTierReference = 0
	STORAGETIERREFERENCE_REDACTED StorageTierReference = 1
	STORAGETIERREFERENCE_PCIE_SSD StorageTierReference = 2
	STORAGETIERREFERENCE_SATA_SSD StorageTierReference = 3
	STORAGETIERREFERENCE_HDD      StorageTierReference = 4
)

// returns the name of the enum given an ordinal number
func (e *StorageTierReference) name(index int) string {
	names := [...]string{
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
	names := [...]string{
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
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ToleranceMessage"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
List of unconfigured nodes
*/
type UnconfigureNodeDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of unconfigured nodes
	*/
	NodeList []UnconfiguredNodeListItem `json:"nodeList,omitempty"`
}

func NewUnconfigureNodeDetails() *UnconfigureNodeDetails {
	p := new(UnconfigureNodeDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UnconfigureNodeDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UnconfigureNodeDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type UnconfiguredNodeAttributeMap struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates if cvm interface can work with 1 GIG NIC or not
	*/
	CanWorkWith1GNic *bool `json:"canWorkWith1GNic,omitempty"`
	/**
	  Default workload
	*/
	DefaultWorkload *string `json:"defaultWorkload,omitempty"`
	/**
	  Indicates whether the model is supported or not
	*/
	IsModelSupported *bool `json:"isModelSupported,omitempty"`
	/**
	  LCM family name
	*/
	LcmFamily *string `json:"lcmFamily,omitempty"`
	/**
	  Indicates whether the hypervisor is robo mixed or not
	*/
	RoboMixedHypervisor *bool `json:"roboMixedHypervisor,omitempty"`
}

func NewUnconfiguredNodeAttributeMap() *UnconfiguredNodeAttributeMap {
	p := new(UnconfiguredNodeAttributeMap)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UnconfiguredNodeAttributeMap"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UnconfiguredNodeAttributeMap"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Unconfigured node details
*/
type UnconfiguredNodeListItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster arch
	*/
	Arch *string `json:"arch,omitempty"`

	Attributes *UnconfiguredNodeAttributeMap `json:"attributes,omitempty"`
	/**
	  Cluster ID
	*/
	ClusterId *string `json:"clusterId,omitempty"`
	/**
	  CPU type
	*/
	CpuType []string `json:"cpuType,omitempty"`
	/**
	  Current CVM VLAN tag
	*/
	CurrentCvmVlanTag *string `json:"currentCvmVlanTag,omitempty"`
	/**
	  Current network interface of a node
	*/
	CurrentNetworkInterface *string `json:"currentNetworkInterface,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`
	/**
	  Foundation version
	*/
	FoundationVersion *string `json:"foundationVersion,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`

	HypervisorType *HypervisorType `json:"hypervisorType,omitempty"`
	/**
	  Host version of the node
	*/
	HypervisorVersion *string `json:"hypervisorVersion,omitempty"`
	/**
	  Interface IPV6 address
	*/
	InterfaceIpv6 *string `json:"interfaceIpv6,omitempty"`

	IpmiIp *import4.IPAddress `json:"ipmiIp,omitempty"`
	/**
	  Secure boot status
	*/
	IsSecureBooted *bool `json:"isSecureBooted,omitempty"`
	/**
	  Position of a node in a rackable unit
	*/
	NodePosition *string `json:"nodePosition,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`
	/**
	  NOS software version of a node
	*/
	NosVersion *string `json:"nosVersion,omitempty"`
	/**
	  Maximum number of nodes in rackable-unit
	*/
	RackableUnitMaxNodes *int64 `json:"rackableUnitMaxNodes,omitempty"`
	/**
	  Rackable unit model type
	*/
	RackableUnitModel *string `json:"rackableUnitModel,omitempty"`
	/**
	  Rackable unit serial name
	*/
	RackableUnitSerial *string `json:"rackableUnitSerial,omitempty"`
}

func NewUnconfiguredNodeListItem() *UnconfiguredNodeListItem {
	p := new(UnconfiguredNodeListItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UnconfiguredNodeListItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UnconfiguredNodeListItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId} Put operation
*/
type UpdateClusterTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateClusterTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateClusterTaskApiResponse() *UpdateClusterTaskApiResponse {
	p := new(UpdateClusterTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UpdateClusterTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UpdateClusterTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateClusterTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateClusterTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateClusterTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/rsyslog-servers/{rsyslogServerExtId} Put operation
*/
type UpdateRsyslogServerTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRsyslogServerTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRsyslogServerTaskApiResponse() *UpdateRsyslogServerTaskApiResponse {
	p := new(UpdateRsyslogServerTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UpdateRsyslogServerTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UpdateRsyslogServerTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateRsyslogServerTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateRsyslogServerTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateRsyslogServerTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/$actions/update-status Post operation
*/
type UpdateSnmpStatusTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSnmpStatusTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSnmpStatusTaskApiResponse() *UpdateSnmpStatusTaskApiResponse {
	p := new(UpdateSnmpStatusTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UpdateSnmpStatusTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UpdateSnmpStatusTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSnmpStatusTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSnmpStatusTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSnmpStatusTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/traps/{trapExtId} Put operation
*/
type UpdateSnmpTrapTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSnmpTrapTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSnmpTrapTaskApiResponse() *UpdateSnmpTrapTaskApiResponse {
	p := new(UpdateSnmpTrapTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UpdateSnmpTrapTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UpdateSnmpTrapTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSnmpTrapTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSnmpTrapTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSnmpTrapTaskApiResponseData()
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
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/snmp/users/{userExtId} Put operation
*/
type UpdateSnmpUserTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateSnmpUserTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateSnmpUserTaskApiResponse() *UpdateSnmpUserTaskApiResponse {
	p := new(UpdateSnmpUserTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UpdateSnmpUserTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UpdateSnmpUserTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateSnmpUserTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateSnmpUserTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateSnmpUserTaskApiResponseData()
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
Uplink information for controller VM
*/
type UplinkInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`
	/**
	  Uplink details for a controller VM
	*/
	UplinkList []NameMacRef `json:"uplinkList,omitempty"`
}

func NewUplinkInfo() *UplinkInfo {
	p := new(UplinkInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UplinkInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UplinkInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Active and standby uplink information of the target nodes
*/
type UplinkNetworkItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the uplink
	*/
	Name *string `json:"name,omitempty"`
	/**
	  List of network types
	*/
	Networks []string `json:"networks,omitempty"`

	Uplinks *Uplinks `json:"uplinks,omitempty"`
}

func NewUplinkNetworkItem() *UplinkNetworkItem {
	p := new(UplinkNetworkItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UplinkNetworkItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UplinkNetworkItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Uplink information of the target nodes
*/
type UplinkNode struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	CvmIp *import4.IPAddress `json:"cvmIp,omitempty"`

	HypervisorIp *import4.IPAddress `json:"hypervisorIp,omitempty"`
	/**
	  Active and standby uplink information of the target nodes
	*/
	Networks []UplinkNetworkItem `json:"networks,omitempty"`
}

func NewUplinkNode() *UplinkNode {
	p := new(UplinkNode)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UplinkNode"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UplinkNode"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Active and standby uplink information of the target nodes
*/
type Uplinks struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Active uplink information
	*/
	Active []UplinksField `json:"active,omitempty"`
	/**
	  Standby uplink information
	*/
	Standby []UplinksField `json:"standby,omitempty"`
}

func NewUplinks() *Uplinks {
	p := new(Uplinks)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.Uplinks"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.Uplinks"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Properties of active and standby uplink
*/
type UplinksField struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Mac address
	*/
	Mac *string `json:"mac,omitempty"`
	/**
	  Interface name
	*/
	Name *string `json:"name,omitempty"`
	/**
	  Interface value
	*/
	Value *string `json:"value,omitempty"`
}

func NewUplinksField() *UplinksField {
	p := new(UplinksField)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UplinksField"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UplinksField"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Upload information for a node
*/
type UploadInfoNodeItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Error message if any, for available hypervisor ISO
	*/
	AvailableHypervisorIsoError *string `json:"availableHypervisorIsoError,omitempty"`
	/**
	  Provides information on whether hypervisor ISO upload is required or not
	*/
	HypervisorUploadRequired *bool `json:"hypervisorUploadRequired,omitempty"`
	/**
	  Indicates if imaging is required or not
	*/
	IsImagingMandatory *bool `json:"isImagingMandatory,omitempty"`
	/**
	  Indicates if node is compatible or not
	*/
	IsNodeCompatible *bool `json:"isNodeCompatible,omitempty"`
	/**
	  UUID of a node
	*/
	NodeUuid *string `json:"nodeUuid,omitempty"`

	RequiredHypervisorType *HypervisorType `json:"requiredHypervisorType,omitempty"`
}

func NewUploadInfoNodeItem() *UploadInfoNodeItem {
	p := new(UploadInfoNodeItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UploadInfoNodeItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UploadInfoNodeItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
UserName and Password model
*/
type UserInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster name
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/**
	  Password
	*/
	Password *string `json:"password,omitempty"`
	/**
	  Username
	*/
	UserName *string `json:"userName,omitempty"`
}

func NewUserInfo() *UserInfo {
	p := new(UserInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.UserInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.UserInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
REST response for all response codes in api path /clustermgmt/v4.0.a2/config/clusters/{clusterExtId}/$actions/validate-uplinks Post operation
*/
type ValidateUplinksTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfValidateUplinksTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewValidateUplinksTaskApiResponse() *ValidateUplinksTaskApiResponse {
	p := new(ValidateUplinksTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.ValidateUplinksTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "clustermgmt.v4.r0.a2.config.ValidateUplinksTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ValidateUplinksTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ValidateUplinksTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfValidateUplinksTaskApiResponseData()
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

type OneOfGetClusterResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ClusterEntity         `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClusterResponseData() *OneOfGetClusterResponseData {
	p := new(OneOfGetClusterResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterResponseData is nil"))
	}
	switch v.(type) {
	case ClusterEntity:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ClusterEntity)
		}
		*p.oneOfType0 = v.(ClusterEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ClusterEntity)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []DomainFaultTolerance `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetDomainFaultToleranceResponseData() *OneOfGetDomainFaultToleranceResponseData {
	p := new(OneOfGetDomainFaultToleranceResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDomainFaultToleranceResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDomainFaultToleranceResponseData is nil"))
	}
	switch v.(type) {
	case []DomainFaultTolerance:
		p.oneOfType0 = v.([]DomainFaultTolerance)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.DomainFaultTolerance>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.DomainFaultTolerance>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.DomainFaultTolerance>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.DomainFaultTolerance>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []HostGpuEntity        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClusterHostGpusResponseData() *OneOfGetClusterHostGpusResponseData {
	p := new(OneOfGetClusterHostGpusResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClusterHostGpusResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClusterHostGpusResponseData is nil"))
	}
	switch v.(type) {
	case []HostGpuEntity:
		p.oneOfType0 = v.([]HostGpuEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *HostGpuEntity         `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostGpuResponseData() *OneOfGetHostGpuResponseData {
	p := new(OneOfGetHostGpuResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHostGpuResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHostGpuResponseData is nil"))
	}
	switch v.(type) {
	case HostGpuEntity:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(HostGpuEntity)
		}
		*p.oneOfType0 = v.(HostGpuEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(HostGpuEntity)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfUpdateSnmpUserTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSnmpUserTaskApiResponseData() *OneOfUpdateSnmpUserTaskApiResponseData {
	p := new(OneOfUpdateSnmpUserTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSnmpUserTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSnmpUserTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateSnmpUserTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSnmpUserTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSnmpUserTaskApiResponseData"))
}

func (p *OneOfUpdateSnmpUserTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSnmpUserTaskApiResponseData")
}

type OneOfAddNodeTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddNodeTaskResponseData() *OneOfAddNodeTaskResponseData {
	p := new(OneOfAddNodeTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddNodeTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddNodeTaskResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddNodeTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddNodeTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddNodeTaskResponseData"))
}

func (p *OneOfAddNodeTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddNodeTaskResponseData")
}

type OneOfGetHostGpusResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []HostGpuEntity        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostGpusResponseData() *OneOfGetHostGpusResponseData {
	p := new(OneOfGetHostGpusResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHostGpusResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHostGpusResponseData is nil"))
	}
	switch v.(type) {
	case []HostGpuEntity:
		p.oneOfType0 = v.([]HostGpuEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []RackableUnit         `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRackableUnitsResponseData() *OneOfGetRackableUnitsResponseData {
	p := new(OneOfGetRackableUnitsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRackableUnitsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRackableUnitsResponseData is nil"))
	}
	switch v.(type) {
	case []RackableUnit:
		p.oneOfType0 = v.([]RackableUnit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.RackableUnit>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.RackableUnit>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.RackableUnit>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.RackableUnit>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfDiscoverNodeTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDiscoverNodeTaskApiResponseData() *OneOfDiscoverNodeTaskApiResponseData {
	p := new(OneOfDiscoverNodeTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDiscoverNodeTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDiscoverNodeTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDiscoverNodeTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDiscoverNodeTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDiscoverNodeTaskApiResponseData"))
}

func (p *OneOfDiscoverNodeTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDiscoverNodeTaskApiResponseData")
}

type OneOfGetAllHostsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []HostEntity           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetAllHostsResponseData() *OneOfGetAllHostsResponseData {
	p := new(OneOfGetAllHostsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAllHostsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAllHostsResponseData is nil"))
	}
	switch v.(type) {
	case []HostEntity:
		p.oneOfType0 = v.([]HostEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAllHostsResponseData) GetValue() interface{} {
	if "List<clustermgmt.v4.config.HostEntity>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAllHostsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]HostEntity)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.HostEntity" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAllHostsResponseData"))
}

func (p *OneOfGetAllHostsResponseData) MarshalJSON() ([]byte, error) {
	if "List<clustermgmt.v4.config.HostEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAllHostsResponseData")
}

type OneOfGetSearchResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SearchResponse        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSearchResponseData() *OneOfGetSearchResponseData {
	p := new(OneOfGetSearchResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSearchResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSearchResponseData is nil"))
	}
	switch v.(type) {
	case SearchResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SearchResponse)
		}
		*p.oneOfType0 = v.(SearchResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSearchResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSearchResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SearchResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.SearchResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SearchResponse)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSearchResponseData"))
}

func (p *OneOfGetSearchResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSearchResponseData")
}

type OneOfHostRenameResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfHostRenameResponseData() *OneOfHostRenameResponseData {
	p := new(OneOfHostRenameResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHostRenameResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHostRenameResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHostRenameResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfHostRenameResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHostRenameResponseData"))
}

func (p *OneOfHostRenameResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfHostRenameResponseData")
}

type OneOfRemoveNodeTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRemoveNodeTaskResponseData() *OneOfRemoveNodeTaskResponseData {
	p := new(OneOfRemoveNodeTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoveNodeTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoveNodeTaskResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRemoveNodeTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRemoveNodeTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveNodeTaskResponseData"))
}

func (p *OneOfRemoveNodeTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRemoveNodeTaskResponseData")
}

type OneOfDeleteRsyslogServerTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteRsyslogServerTaskApiResponseData() *OneOfDeleteRsyslogServerTaskApiResponseData {
	p := new(OneOfDeleteRsyslogServerTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteRsyslogServerTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteRsyslogServerTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteRsyslogServerTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteRsyslogServerTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteRsyslogServerTaskApiResponseData"))
}

func (p *OneOfDeleteRsyslogServerTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteRsyslogServerTaskApiResponseData")
}

type OneOfAddSnmpTransportsTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddSnmpTransportsTaskApiResponseData() *OneOfAddSnmpTransportsTaskApiResponseData {
	p := new(OneOfAddSnmpTransportsTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddSnmpTransportsTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddSnmpTransportsTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddSnmpTransportsTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddSnmpTransportsTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddSnmpTransportsTaskApiResponseData"))
}

func (p *OneOfAddSnmpTransportsTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddSnmpTransportsTaskApiResponseData")
}

type OneOfGetRsyslogServerResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RsyslogServer         `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRsyslogServerResponseData() *OneOfGetRsyslogServerResponseData {
	p := new(OneOfGetRsyslogServerResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRsyslogServerResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRsyslogServerResponseData is nil"))
	}
	switch v.(type) {
	case RsyslogServer:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RsyslogServer)
		}
		*p.oneOfType0 = v.(RsyslogServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetRsyslogServerResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetRsyslogServerResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(RsyslogServer)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.RsyslogServer" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RsyslogServer)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetRsyslogServerResponseData"))
}

func (p *OneOfGetRsyslogServerResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetRsyslogServerResponseData")
}

type OneOfGetGpuProfilesResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []GpuProfile           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetGpuProfilesResponseData() *OneOfGetGpuProfilesResponseData {
	p := new(OneOfGetGpuProfilesResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetGpuProfilesResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetGpuProfilesResponseData is nil"))
	}
	switch v.(type) {
	case []GpuProfile:
		p.oneOfType0 = v.([]GpuProfile)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.GpuProfile>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.GpuProfile>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.GpuProfile>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.GpuProfile>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfSearchResponseResponse struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType0    *UnconfigureNodeDetails `json:"-"`
	oneOfType1    *NodeNetworkingDetails  `json:"-"`
	oneOfType2    *HypervisorUploadInfo   `json:"-"`
}

func NewOneOfSearchResponseResponse() *OneOfSearchResponseResponse {
	p := new(OneOfSearchResponseResponse)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSearchResponseResponse) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSearchResponseResponse is nil"))
	}
	switch v.(type) {
	case UnconfigureNodeDetails:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(UnconfigureNodeDetails)
		}
		*p.oneOfType0 = v.(UnconfigureNodeDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case NodeNetworkingDetails:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(NodeNetworkingDetails)
		}
		*p.oneOfType1 = v.(NodeNetworkingDetails)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case HypervisorUploadInfo:
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(HypervisorUploadInfo)
		}
		*p.oneOfType2 = v.(HypervisorUploadInfo)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSearchResponseResponse) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2
	}
	return nil
}

func (p *OneOfSearchResponseResponse) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(UnconfigureNodeDetails)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.UnconfigureNodeDetails" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(UnconfigureNodeDetails)
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
	vOneOfType1 := new(NodeNetworkingDetails)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "clustermgmt.v4.config.NodeNetworkingDetails" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(NodeNetworkingDetails)
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
	vOneOfType2 := new(HypervisorUploadInfo)
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if "clustermgmt.v4.config.HypervisorUploadInfo" == *vOneOfType2.ObjectType_ {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(HypervisorUploadInfo)
			}
			*p.oneOfType2 = *vOneOfType2
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSearchResponseResponse"))
}

func (p *OneOfSearchResponseResponse) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType2 != nil && *p.oneOfType2.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	return nil, errors.New("No value to marshal for OneOfSearchResponseResponse")
}

type OneOfAddRsyslogServerTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddRsyslogServerTaskApiResponseData() *OneOfAddRsyslogServerTaskApiResponseData {
	p := new(OneOfAddRsyslogServerTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddRsyslogServerTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddRsyslogServerTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddRsyslogServerTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddRsyslogServerTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddRsyslogServerTaskApiResponseData"))
}

func (p *OneOfAddRsyslogServerTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddRsyslogServerTaskApiResponseData")
}

type OneOfAddSnmpUserTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddSnmpUserTaskApiResponseData() *OneOfAddSnmpUserTaskApiResponseData {
	p := new(OneOfAddSnmpUserTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddSnmpUserTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddSnmpUserTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddSnmpUserTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddSnmpUserTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddSnmpUserTaskApiResponseData"))
}

func (p *OneOfAddSnmpUserTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddSnmpUserTaskApiResponseData")
}

type OneOfUpdateClusterTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateClusterTaskApiResponseData() *OneOfUpdateClusterTaskApiResponseData {
	p := new(OneOfUpdateClusterTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateClusterTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateClusterTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateClusterTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateClusterTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateClusterTaskApiResponseData"))
}

func (p *OneOfUpdateClusterTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateClusterTaskApiResponseData")
}

type OneOfGetAllHostGpusResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []HostGpuEntity        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetAllHostGpusResponseData() *OneOfGetAllHostGpusResponseData {
	p := new(OneOfGetAllHostGpusResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAllHostGpusResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAllHostGpusResponseData is nil"))
	}
	switch v.(type) {
	case []HostGpuEntity:
		p.oneOfType0 = v.([]HostGpuEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAllHostGpusResponseData) GetValue() interface{} {
	if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetAllHostGpusResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]HostGpuEntity)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "clustermgmt.v4.config.HostGpuEntity" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.HostGpuEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.HostGpuEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAllHostGpusResponseData"))
}

func (p *OneOfGetAllHostGpusResponseData) MarshalJSON() ([]byte, error) {
	if "List<clustermgmt.v4.config.HostGpuEntity>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetAllHostGpusResponseData")
}

type OneOfClusterStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *StatsQueryResp        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfClusterStatsApiResponseData() *OneOfClusterStatsApiResponseData {
	p := new(OneOfClusterStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case StatsQueryResp:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(StatsQueryResp)
		}
		*p.oneOfType0 = v.(StatsQueryResp)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfClusterStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfClusterStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(StatsQueryResp)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.StatsQueryResp" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(StatsQueryResp)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterStatsApiResponseData"))
}

func (p *OneOfClusterStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfClusterStatsApiResponseData")
}

type OneOfUpdateSnmpTrapTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSnmpTrapTaskApiResponseData() *OneOfUpdateSnmpTrapTaskApiResponseData {
	p := new(OneOfUpdateSnmpTrapTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSnmpTrapTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSnmpTrapTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateSnmpTrapTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSnmpTrapTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSnmpTrapTaskApiResponseData"))
}

func (p *OneOfUpdateSnmpTrapTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSnmpTrapTaskApiResponseData")
}

type OneOfRemoveSnmpTransportsTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfRemoveSnmpTransportsTaskApiResponseData() *OneOfRemoveSnmpTransportsTaskApiResponseData {
	p := new(OneOfRemoveSnmpTransportsTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoveSnmpTransportsTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoveSnmpTransportsTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfRemoveSnmpTransportsTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfRemoveSnmpTransportsTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveSnmpTransportsTaskApiResponseData"))
}

func (p *OneOfRemoveSnmpTransportsTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfRemoveSnmpTransportsTaskApiResponseData")
}

type OneOfGetHostsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []HostEntity           `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostsResponseData() *OneOfGetHostsResponseData {
	p := new(OneOfGetHostsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHostsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHostsResponseData is nil"))
	}
	switch v.(type) {
	case []HostEntity:
		p.oneOfType0 = v.([]HostEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.HostEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.HostEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfUpdateRsyslogServerTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateRsyslogServerTaskApiResponseData() *OneOfUpdateRsyslogServerTaskApiResponseData {
	p := new(OneOfUpdateRsyslogServerTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRsyslogServerTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRsyslogServerTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateRsyslogServerTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateRsyslogServerTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRsyslogServerTaskApiResponseData"))
}

func (p *OneOfUpdateRsyslogServerTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRsyslogServerTaskApiResponseData")
}

type OneOfDeleteSnmpTrapTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSnmpTrapTaskApiResponseData() *OneOfDeleteSnmpTrapTaskApiResponseData {
	p := new(OneOfDeleteSnmpTrapTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSnmpTrapTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSnmpTrapTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteSnmpTrapTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSnmpTrapTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSnmpTrapTaskApiResponseData"))
}

func (p *OneOfDeleteSnmpTrapTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSnmpTrapTaskApiResponseData")
}

type OneOfValidateUplinksTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfValidateUplinksTaskApiResponseData() *OneOfValidateUplinksTaskApiResponseData {
	p := new(OneOfValidateUplinksTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfValidateUplinksTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfValidateUplinksTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfValidateUplinksTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfValidateUplinksTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfValidateUplinksTaskApiResponseData"))
}

func (p *OneOfValidateUplinksTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfValidateUplinksTaskApiResponseData")
}

type OneOfDeleteSnmpUserTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfDeleteSnmpUserTaskApiResponseData() *OneOfDeleteSnmpUserTaskApiResponseData {
	p := new(OneOfDeleteSnmpUserTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteSnmpUserTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteSnmpUserTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteSnmpUserTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfDeleteSnmpUserTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteSnmpUserTaskApiResponseData"))
}

func (p *OneOfDeleteSnmpUserTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteSnmpUserTaskApiResponseData")
}

type OneOfIsBundleCompatibleTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfIsBundleCompatibleTaskResponseData() *OneOfIsBundleCompatibleTaskResponseData {
	p := new(OneOfIsBundleCompatibleTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfIsBundleCompatibleTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfIsBundleCompatibleTaskResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfIsBundleCompatibleTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfIsBundleCompatibleTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfIsBundleCompatibleTaskResponseData"))
}

func (p *OneOfIsBundleCompatibleTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfIsBundleCompatibleTaskResponseData")
}

type OneOfAddSnmpTrapTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfAddSnmpTrapTaskApiResponseData() *OneOfAddSnmpTrapTaskApiResponseData {
	p := new(OneOfAddSnmpTrapTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddSnmpTrapTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddSnmpTrapTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfAddSnmpTrapTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfAddSnmpTrapTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddSnmpTrapTaskApiResponseData"))
}

func (p *OneOfAddSnmpTrapTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfAddSnmpTrapTaskApiResponseData")
}

type OneOfGetSnmpUserResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SnmpUser              `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSnmpUserResponseData() *OneOfGetSnmpUserResponseData {
	p := new(OneOfGetSnmpUserResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSnmpUserResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSnmpUserResponseData is nil"))
	}
	switch v.(type) {
	case SnmpUser:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnmpUser)
		}
		*p.oneOfType0 = v.(SnmpUser)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSnmpUserResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSnmpUserResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SnmpUser)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.SnmpUser" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnmpUser)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSnmpUserResponseData"))
}

func (p *OneOfGetSnmpUserResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSnmpUserResponseData")
}

type OneOfHypervisorUplpadRequiredTaskResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfHypervisorUplpadRequiredTaskResponseData() *OneOfHypervisorUplpadRequiredTaskResponseData {
	p := new(OneOfHypervisorUplpadRequiredTaskResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHypervisorUplpadRequiredTaskResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHypervisorUplpadRequiredTaskResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHypervisorUplpadRequiredTaskResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfHypervisorUplpadRequiredTaskResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHypervisorUplpadRequiredTaskResponseData"))
}

func (p *OneOfHypervisorUplpadRequiredTaskResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfHypervisorUplpadRequiredTaskResponseData")
}

type OneOfUpdateSnmpStatusTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfUpdateSnmpStatusTaskApiResponseData() *OneOfUpdateSnmpStatusTaskApiResponseData {
	p := new(OneOfUpdateSnmpStatusTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateSnmpStatusTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateSnmpStatusTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateSnmpStatusTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfUpdateSnmpStatusTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateSnmpStatusTaskApiResponseData"))
}

func (p *OneOfUpdateSnmpStatusTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateSnmpStatusTaskApiResponseData")
}

type OneOfGetSnmpTrapResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SnmpTrap              `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSnmpTrapResponseData() *OneOfGetSnmpTrapResponseData {
	p := new(OneOfGetSnmpTrapResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSnmpTrapResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSnmpTrapResponseData is nil"))
	}
	switch v.(type) {
	case SnmpTrap:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnmpTrap)
		}
		*p.oneOfType0 = v.(SnmpTrap)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetSnmpTrapResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetSnmpTrapResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(SnmpTrap)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "clustermgmt.v4.config.SnmpTrap" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnmpTrap)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetSnmpTrapResponseData"))
}

func (p *OneOfGetSnmpTrapResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetSnmpTrapResponseData")
}

type OneOfGetRsyslogServersResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []RsyslogServer        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRsyslogServersResponseData() *OneOfGetRsyslogServersResponseData {
	p := new(OneOfGetRsyslogServersResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRsyslogServersResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRsyslogServersResponseData is nil"))
	}
	switch v.(type) {
	case []RsyslogServer:
		p.oneOfType0 = v.([]RsyslogServer)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.RsyslogServer>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.RsyslogServer>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.RsyslogServer>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.RsyslogServer>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ClusterEntity        `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetClustersResponseData() *OneOfGetClustersResponseData {
	p := new(OneOfGetClustersResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetClustersResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetClustersResponseData is nil"))
	}
	switch v.(type) {
	case []ClusterEntity:
		p.oneOfType0 = v.([]ClusterEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<clustermgmt.v4.config.ClusterEntity>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<clustermgmt.v4.config.ClusterEntity>"
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<clustermgmt.v4.config.ClusterEntity>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<clustermgmt.v4.config.ClusterEntity>"
			return nil

		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfGetHostResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *HostEntity            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetHostResponseData() *OneOfGetHostResponseData {
	p := new(OneOfGetHostResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHostResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHostResponseData is nil"))
	}
	switch v.(type) {
	case HostEntity:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(HostEntity)
		}
		*p.oneOfType0 = v.(HostEntity)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(HostEntity)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *SnmpConfig            `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetSnmpResponseData() *OneOfGetSnmpResponseData {
	p := new(OneOfGetSnmpResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetSnmpResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetSnmpResponseData is nil"))
	}
	switch v.(type) {
	case SnmpConfig:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SnmpConfig)
		}
		*p.oneOfType0 = v.(SnmpConfig)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SnmpConfig)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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

type OneOfGetNodeNetworkingTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *import1.TaskReference `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetNodeNetworkingTaskApiResponseData() *OneOfGetNodeNetworkingTaskApiResponseData {
	p := new(OneOfGetNodeNetworkingTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNodeNetworkingTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNodeNetworkingTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import1.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import1.TaskReference)
		}
		*p.oneOfType0 = v.(import1.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetNodeNetworkingTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNodeNetworkingTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import1.TaskReference)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNodeNetworkingTaskApiResponseData"))
}

func (p *OneOfGetNodeNetworkingTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNodeNetworkingTaskApiResponseData")
}

type OneOfGetRackableUnitResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *RackableUnit          `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetRackableUnitResponseData() *OneOfGetRackableUnitResponseData {
	p := new(OneOfGetRackableUnitResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetRackableUnitResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetRackableUnitResponseData is nil"))
	}
	switch v.(type) {
	case RackableUnit:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(RackableUnit)
		}
		*p.oneOfType0 = v.(RackableUnit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
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
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(RackableUnit)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "clustermgmt.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
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
