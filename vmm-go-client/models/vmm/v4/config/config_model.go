/*
 * Generated file models/vmm/v4/config/config_model.go.
 *
 * Product version: 4.0.2-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Configure Tasks
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
)

/**
Indicates the order of device types in which VM should try to boot from. If boot device order is not provided the system will decide appropriate boot device order.
*/
type BootDeviceOrder int

const (
	BOOTDEVICEORDER_UNKNOWN  BootDeviceOrder = 0
	BOOTDEVICEORDER_REDACTED BootDeviceOrder = 1
	BOOTDEVICEORDER_CDROM    BootDeviceOrder = 2
	BOOTDEVICEORDER_DISK     BootDeviceOrder = 3
	BOOTDEVICEORDER_NETWORK  BootDeviceOrder = 4
)

// returns the name of the enum given an ordinal number
func (e *BootDeviceOrder) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *BootDeviceOrder) index(name string) BootDeviceOrder {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CDROM",
		"DISK",
		"NETWORK",
	}
	for idx := range names {
		if names[idx] == name {
			return BootDeviceOrder(idx)
		}
	}
	return BOOTDEVICEORDER_UNKNOWN
}

func (e *BootDeviceOrder) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BootDeviceOrder:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BootDeviceOrder) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BootDeviceOrder) Ref() *BootDeviceOrder {
	return &e
}

/**
Indicates whether the VM should use UEFI boot or Legacy boot. If UEFI boot is enabled then other legacy boot options (like bootDevice and bootDeviceOrderList) are ignored.
*/
type BootType int

const (
	BOOTTYPE_UNKNOWN  BootType = 0
	BOOTTYPE_REDACTED BootType = 1
	BOOTTYPE_UEFI     BootType = 2
	BOOTTYPE_LEGACY   BootType = 3
)

// returns the name of the enum given an ordinal number
func (e *BootType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UEFI",
		"LEGACY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *BootType) index(name string) BootType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"UEFI",
		"LEGACY",
	}
	for idx := range names {
		if names[idx] == name {
			return BootType(idx)
		}
	}
	return BOOTTYPE_UNKNOWN
}

func (e *BootType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for BootType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *BootType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e BootType) Ref() *BootType {
	return &e
}

/**
If this field is set, the guest will be customized using CloudInit. Either userData or KVPair should be provided. If KVPair are provided then the user data will be generated using these key-value pairs.
*/
type CloudInit struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`

	Config *OneOfCloudInitConfig `json:"config,omitempty"`
}

func NewCloudInit() *CloudInit {
	p := new(CloudInit)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.CloudInit"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.CloudInit"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CloudInit) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *CloudInit) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfCloudInitConfig()
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

type CloudInitDataSource struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The contents of the metaData configuration for CloudInit. This can be formatted as YAML or JSON. The value must be base64 encoded.
	*/
	MetaData *string `json:"metaData,omitempty"`
	/**
	  The contents of the userData configuration for CloudInit. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.
	*/
	UserData *string `json:"userData,omitempty"`
}

func NewCloudInitDataSource() *CloudInitDataSource {
	p := new(CloudInitDataSource)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.CloudInitDataSource"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.CloudInitDataSource"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Credentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Password *string `json:"password,omitempty"`

	Username *string `json:"username"`
}

func (p *Credentials) MarshalJSON() ([]byte, error) {
	type CredentialsProxy Credentials
	return json.Marshal(struct {
		*CredentialsProxy
		Username *string `json:"username,omitempty"`
	}{
		CredentialsProxy: (*CredentialsProxy)(p),
		Username:         p.Username,
	})
}

func NewCredentials() *Credentials {
	p := new(Credentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.Credentials"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.Credentials"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DataSourceReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataSourceExtId *string `json:"dataSourceExtId,omitempty"`
	/**
	  This is to indicate if attaching the referenced disk directly. Important - this should only be used by internal services. Direct attaching a disk that is used by another VM will result in data loss.
	*/
	IsDirectAttach *bool `json:"isDirectAttach,omitempty"`

	Kind *string `json:"kind,omitempty"`

	Name *string `json:"name,omitempty"`

	ParentReference *Reference `json:"parentReference,omitempty"`
}

func NewDataSourceReference() *DataSourceReference {
	p := new(DataSourceReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.DataSourceReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.DataSourceReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DiskAdapterType int

const (
	DISKADAPTERTYPE_UNKNOWN  DiskAdapterType = 0
	DISKADAPTERTYPE_REDACTED DiskAdapterType = 1
	DISKADAPTERTYPE_SCSI     DiskAdapterType = 2
	DISKADAPTERTYPE_IDE      DiskAdapterType = 3
	DISKADAPTERTYPE_PCI      DiskAdapterType = 4
	DISKADAPTERTYPE_SATA     DiskAdapterType = 5
	DISKADAPTERTYPE_SPAPR    DiskAdapterType = 6
)

// returns the name of the enum given an ordinal number
func (e *DiskAdapterType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *DiskAdapterType) index(name string) DiskAdapterType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SCSI",
		"IDE",
		"PCI",
		"SATA",
		"SPAPR",
	}
	for idx := range names {
		if names[idx] == name {
			return DiskAdapterType(idx)
		}
	}
	return DISKADAPTERTYPE_UNKNOWN
}

func (e *DiskAdapterType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DiskAdapterType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DiskAdapterType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DiskAdapterType) Ref() *DiskAdapterType {
	return &e
}

type DiskDeviceProperties struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DeviceType *DiskDeviceType `json:"deviceType,omitempty"`

	DiskAddress *VMTemplateDiskAddress `json:"diskAddress,omitempty"`
}

func NewDiskDeviceProperties() *DiskDeviceProperties {
	p := new(DiskDeviceProperties)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.DiskDeviceProperties"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.DiskDeviceProperties"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type DiskDeviceType int

const (
	DISKDEVICETYPE_UNKNOWN      DiskDeviceType = 0
	DISKDEVICETYPE_REDACTED     DiskDeviceType = 1
	DISKDEVICETYPE_DISK         DiskDeviceType = 2
	DISKDEVICETYPE_CDROM        DiskDeviceType = 3
	DISKDEVICETYPE_VOLUME_GROUP DiskDeviceType = 4
)

// returns the name of the enum given an ordinal number
func (e *DiskDeviceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISK",
		"CDROM",
		"VOLUME_GROUP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *DiskDeviceType) index(name string) DiskDeviceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DISK",
		"CDROM",
		"VOLUME_GROUP",
	}
	for idx := range names {
		if names[idx] == name {
			return DiskDeviceType(idx)
		}
	}
	return DISKDEVICETYPE_UNKNOWN
}

func (e *DiskDeviceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for DiskDeviceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *DiskDeviceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e DiskDeviceType) Ref() *DiskDeviceType {
	return &e
}

/**
The mode of this GPU.
*/
type GpuMode int

const (
	GPUMODE_UNKNOWN              GpuMode = 0
	GPUMODE_REDACTED             GpuMode = 1
	GPUMODE_PASSTHROUGH_GRAPHICS GpuMode = 2
	GPUMODE_PASSTHROUGH_COMPUTE  GpuMode = 3
	GPUMODE_VIRTUAL              GpuMode = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
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
		"PASSTHROUGH_GRAPHICS",
		"PASSTHROUGH_COMPUTE",
		"VIRTUAL",
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
The vendor of the GPU.
*/
type GpuVendor int

const (
	GPUVENDOR_UNKNOWN  GpuVendor = 0
	GPUVENDOR_REDACTED GpuVendor = 1
	GPUVENDOR_NVIDIA   GpuVendor = 2
	GPUVENDOR_INTEL    GpuVendor = 3
	GPUVENDOR_AMD      GpuVendor = 4
)

// returns the name of the enum given an ordinal number
func (e *GpuVendor) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *GpuVendor) index(name string) GpuVendor {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NVIDIA",
		"INTEL",
		"AMD",
	}
	for idx := range names {
		if names[idx] == name {
			return GpuVendor(idx)
		}
	}
	return GPUVENDOR_UNKNOWN
}

func (e *GpuVendor) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for GpuVendor:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *GpuVendor) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e GpuVendor) Ref() *GpuVendor {
	return &e
}

/**
Indicates whether IP address is DHCP or Static.
*/
type IPAssignmentType int

const (
	IPASSIGNMENTTYPE_UNKNOWN  IPAssignmentType = 0
	IPASSIGNMENTTYPE_REDACTED IPAssignmentType = 1
	IPASSIGNMENTTYPE_DHCP     IPAssignmentType = 2
	IPASSIGNMENTTYPE_STATIC   IPAssignmentType = 3
)

// returns the name of the enum given an ordinal number
func (e *IPAssignmentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DHCP",
		"STATIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *IPAssignmentType) index(name string) IPAssignmentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DHCP",
		"STATIC",
	}
	for idx := range names {
		if names[idx] == name {
			return IPAssignmentType(idx)
		}
	}
	return IPASSIGNMENTTYPE_UNKNOWN
}

func (e *IPAssignmentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IPAssignmentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IPAssignmentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IPAssignmentType) Ref() *IPAssignmentType {
	return &e
}

/**
Address type. It can only be 'ASSIGNED' in the spec. If no type is specified in the spec, the default type is set to 'ASSIGNED'.
*/
type IPDiscoveredType int

const (
	IPDISCOVEREDTYPE_UNKNOWN  IPDiscoveredType = 0
	IPDISCOVEREDTYPE_REDACTED IPDiscoveredType = 1
	IPDISCOVEREDTYPE_ASSIGNED IPDiscoveredType = 2
	IPDISCOVEREDTYPE_LEARNED  IPDiscoveredType = 3
)

// returns the name of the enum given an ordinal number
func (e *IPDiscoveredType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGNED",
		"LEARNED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *IPDiscoveredType) index(name string) IPDiscoveredType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ASSIGNED",
		"LEARNED",
	}
	for idx := range names {
		if names[idx] == name {
			return IPDiscoveredType(idx)
		}
	}
	return IPDISCOVEREDTYPE_UNKNOWN
}

func (e *IPDiscoveredType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IPDiscoveredType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IPDiscoveredType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IPDiscoveredType) Ref() *IPDiscoveredType {
	return &e
}

/**
Desired mount state of Nutanix Guest Tools ISO.
*/
type IsoMountState int

const (
	ISOMOUNTSTATE_UNKNOWN   IsoMountState = 0
	ISOMOUNTSTATE_REDACTED  IsoMountState = 1
	ISOMOUNTSTATE_MOUNTED   IsoMountState = 2
	ISOMOUNTSTATE_UNMOUNTED IsoMountState = 3
)

// returns the name of the enum given an ordinal number
func (e *IsoMountState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MOUNTED",
		"UNMOUNTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *IsoMountState) index(name string) IsoMountState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"MOUNTED",
		"UNMOUNTED",
	}
	for idx := range names {
		if names[idx] == name {
			return IsoMountState(idx)
		}
	}
	return ISOMOUNTSTATE_UNKNOWN
}

func (e *IsoMountState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for IsoMountState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *IsoMountState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e IsoMountState) Ref() *IsoMountState {
	return &e
}

/**
Machine Type.
*/
type MachineType int

const (
	MACHINETYPE_UNKNOWN  MachineType = 0
	MACHINETYPE_REDACTED MachineType = 1
	MACHINETYPE_PC       MachineType = 2
	MACHINETYPE_PSERIES  MachineType = 3
	MACHINETYPE_Q35      MachineType = 4
)

// returns the name of the enum given an ordinal number
func (e *MachineType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *MachineType) index(name string) MachineType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PC",
		"PSERIES",
		"Q35",
	}
	for idx := range names {
		if names[idx] == name {
			return MachineType(idx)
		}
	}
	return MACHINETYPE_UNKNOWN
}

func (e *MachineType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for MachineType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *MachineType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e MachineType) Ref() *MachineType {
	return &e
}

/**
The type of this Network function NIC. Defaults to INGRESS.
*/
type NetworkFunctionNicType int

const (
	NETWORKFUNCTIONNICTYPE_UNKNOWN  NetworkFunctionNicType = 0
	NETWORKFUNCTIONNICTYPE_REDACTED NetworkFunctionNicType = 1
	NETWORKFUNCTIONNICTYPE_INGRESS  NetworkFunctionNicType = 2
	NETWORKFUNCTIONNICTYPE_EGRESS   NetworkFunctionNicType = 3
	NETWORKFUNCTIONNICTYPE_TAP      NetworkFunctionNicType = 4
)

// returns the name of the enum given an ordinal number
func (e *NetworkFunctionNicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NetworkFunctionNicType) index(name string) NetworkFunctionNicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INGRESS",
		"EGRESS",
		"TAP",
	}
	for idx := range names {
		if names[idx] == name {
			return NetworkFunctionNicType(idx)
		}
	}
	return NETWORKFUNCTIONNICTYPE_UNKNOWN
}

func (e *NetworkFunctionNicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NetworkFunctionNicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NetworkFunctionNicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NetworkFunctionNicType) Ref() *NetworkFunctionNicType {
	return &e
}

/**
Nutanix guest tools is installed or not.
*/
type NgtAvailableState int

const (
	NGTAVAILABLESTATE_UNKNOWN     NgtAvailableState = 0
	NGTAVAILABLESTATE_REDACTED    NgtAvailableState = 1
	NGTAVAILABLESTATE_INSTALLED   NgtAvailableState = 2
	NGTAVAILABLESTATE_UNINSTALLED NgtAvailableState = 3
)

// returns the name of the enum given an ordinal number
func (e *NgtAvailableState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INSTALLED",
		"UNINSTALLED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NgtAvailableState) index(name string) NgtAvailableState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"INSTALLED",
		"UNINSTALLED",
	}
	for idx := range names {
		if names[idx] == name {
			return NgtAvailableState(idx)
		}
	}
	return NGTAVAILABLESTATE_UNKNOWN
}

func (e *NgtAvailableState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NgtAvailableState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NgtAvailableState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NgtAvailableState) Ref() *NgtAvailableState {
	return &e
}

/**
Application names that are enabled.
*/
type NgtEnabledCapability int

const (
	NGTENABLEDCAPABILITY_UNKNOWN              NgtEnabledCapability = 0
	NGTENABLEDCAPABILITY_REDACTED             NgtEnabledCapability = 1
	NGTENABLEDCAPABILITY_SELF_SERVICE_RESTORE NgtEnabledCapability = 2
	NGTENABLEDCAPABILITY_VSS_SNAPSHOT         NgtEnabledCapability = 3
)

// returns the name of the enum given an ordinal number
func (e *NgtEnabledCapability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NgtEnabledCapability) index(name string) NgtEnabledCapability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	for idx := range names {
		if names[idx] == name {
			return NgtEnabledCapability(idx)
		}
	}
	return NGTENABLEDCAPABILITY_UNKNOWN
}

func (e *NgtEnabledCapability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NgtEnabledCapability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NgtEnabledCapability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NgtEnabledCapability) Ref() *NgtEnabledCapability {
	return &e
}

/**
Nutanix Guest Tools is enabled or not.
*/
type NgtState int

const (
	NGTSTATE_UNKNOWN  NgtState = 0
	NGTSTATE_REDACTED NgtState = 1
	NGTSTATE_ENABLED  NgtState = 2
	NGTSTATE_DISABLED NgtState = 3
)

// returns the name of the enum given an ordinal number
func (e *NgtState) name(index int) string {
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

// returns the enum type given a string value
func (e *NgtState) index(name string) NgtState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return NgtState(idx)
		}
	}
	return NGTSTATE_UNKNOWN
}

func (e *NgtState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NgtState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NgtState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NgtState) Ref() *NgtState {
	return &e
}

/**
MAC address of nic to boot from.
*/
type NicMacAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NicMacAddress *string `json:"nicMacAddress,omitempty"`
}

func NewNicMacAddress() *NicMacAddress {
	p := new(NicMacAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.NicMacAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.NicMacAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The model of this NIC.
*/
type NicModel int

const (
	NICMODEL_UNKNOWN  NicModel = 0
	NICMODEL_REDACTED NicModel = 1
	NICMODEL_VIRTIO   NicModel = 2
	NICMODEL_E1000    NicModel = 3
)

// returns the name of the enum given an ordinal number
func (e *NicModel) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NicModel) index(name string) NicModel {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VIRTIO",
		"E1000",
	}
	for idx := range names {
		if names[idx] == name {
			return NicModel(idx)
		}
	}
	return NICMODEL_UNKNOWN
}

func (e *NicModel) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicModel:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicModel) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicModel) Ref() *NicModel {
	return &e
}

/**
The NIC's UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.
*/
type NicType int

const (
	NICTYPE_UNKNOWN              NicType = 0
	NICTYPE_REDACTED             NicType = 1
	NICTYPE_NORMAL_NIC           NicType = 2
	NICTYPE_DIRECT_NIC           NicType = 3
	NICTYPE_NETWORK_FUNCTION_NIC NicType = 4
)

// returns the name of the enum given an ordinal number
func (e *NicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NicType) index(name string) NicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
		"NETWORK_FUNCTION_NIC",
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

/**
The current or desired power state of the VM.
*/
type PowerState int

const (
	POWERSTATE_UNKNOWN  PowerState = 0
	POWERSTATE_REDACTED PowerState = 1
	POWERSTATE_TRUE     PowerState = 2
	POWERSTATE_FALSE    PowerState = 3
)

// returns the name of the enum given an ordinal number
func (e *PowerState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"true",
		"false",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *PowerState) index(name string) PowerState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"true",
		"false",
	}
	for idx := range names {
		if names[idx] == name {
			return PowerState(idx)
		}
	}
	return POWERSTATE_UNKNOWN
}

func (e *PowerState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PowerState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PowerState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PowerState) Ref() *PowerState {
	return &e
}

/**
Power state mechanism (ACPI, GUEST, HARD).
*/
type PowerStateMechanism int

const (
	POWERSTATEMECHANISM_UNKNOWN  PowerStateMechanism = 0
	POWERSTATEMECHANISM_REDACTED PowerStateMechanism = 1
	POWERSTATEMECHANISM_ACPI     PowerStateMechanism = 2
	POWERSTATEMECHANISM_GUEST    PowerStateMechanism = 3
	POWERSTATEMECHANISM_HARD     PowerStateMechanism = 4
)

// returns the name of the enum given an ordinal number
func (e *PowerStateMechanism) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACPI",
		"GUEST",
		"HARD",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *PowerStateMechanism) index(name string) PowerStateMechanism {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACPI",
		"GUEST",
		"HARD",
	}
	for idx := range names {
		if names[idx] == name {
			return PowerStateMechanism(idx)
		}
	}
	return POWERSTATEMECHANISM_UNKNOWN
}

func (e *PowerStateMechanism) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PowerStateMechanism:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PowerStateMechanism) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PowerStateMechanism) Ref() *PowerStateMechanism {
	return &e
}

/**
Reference to an object.
*/
type Reference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ExtId *string `json:"extId,omitempty"`

	Kind *ReferenceType `json:"kind,omitempty"`

	Name *string `json:"name,omitempty"`

	ObjectType *string `json:"objectType,omitempty"`
}

func NewReference() *Reference {
	p := new(Reference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.Reference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.Reference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The kind of the Reference.
*/
type ReferenceType int

const (
	REFERENCETYPE_UNKNOWN                ReferenceType = 0
	REFERENCETYPE_REDACTED               ReferenceType = 1
	REFERENCETYPE_VM                     ReferenceType = 2
	REFERENCETYPE_VOLUME_GROUP           ReferenceType = 3
	REFERENCETYPE_STORAGE_CONTAINER      ReferenceType = 4
	REFERENCETYPE_SUBNET                 ReferenceType = 5
	REFERENCETYPE_NETWORK_FUNCTION_CHAIN ReferenceType = 6
)

// returns the name of the enum given an ordinal number
func (e *ReferenceType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
		"STORAGE_CONTAINER",
		"SUBNET",
		"NETWORK_FUNCTION_CHAIN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ReferenceType) index(name string) ReferenceType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"VOLUME_GROUP",
		"STORAGE_CONTAINER",
		"SUBNET",
		"NETWORK_FUNCTION_CHAIN",
	}
	for idx := range names {
		if names[idx] == name {
			return ReferenceType(idx)
		}
	}
	return REFERENCETYPE_UNKNOWN
}

func (e *ReferenceType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReferenceType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReferenceType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReferenceType) Ref() *ReferenceType {
	return &e
}

/**
If this field is set, the guest will be customized using Sysprep. Either SysprepUnattendXml or KVPair should be provided. If KVPair are provided then the unattended answer file will be generated using these key-value pairs.
*/
type Sysprep struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`

	Config *OneOfSysprepConfig `json:"config,omitempty"`

	ConfigItemDiscriminator *string `json:"configItemDiscriminator,omitempty"`

	InstallType *SysprepInstallType `json:"installType,omitempty"`
}

func NewSysprep() *Sysprep {
	p := new(Sysprep)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.Sysprep"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.Sysprep"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *Sysprep) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *Sysprep) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfSysprepConfig()
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

/**
Whether the guest will be freshly installed using this unattend configuration, or whether this unattend configuration will be applied to a pre-prepared image. Default is 'PREPARED'.
*/
type SysprepInstallType int

const (
	SYSPREPINSTALLTYPE_UNKNOWN  SysprepInstallType = 0
	SYSPREPINSTALLTYPE_REDACTED SysprepInstallType = 1
	SYSPREPINSTALLTYPE_FRESH    SysprepInstallType = 2
	SYSPREPINSTALLTYPE_PREPARED SysprepInstallType = 3
)

// returns the name of the enum given an ordinal number
func (e *SysprepInstallType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *SysprepInstallType) index(name string) SysprepInstallType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"FRESH",
		"PREPARED",
	}
	for idx := range names {
		if names[idx] == name {
			return SysprepInstallType(idx)
		}
	}
	return SYSPREPINSTALLTYPE_UNKNOWN
}

func (e *SysprepInstallType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for SysprepInstallType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *SysprepInstallType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e SysprepInstallType) Ref() *SysprepInstallType {
	return &e
}

type SysprepUnattendXml struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  This field contains a Sysprep unattend xml definition, as a string. The value must be base64 encoded.
	*/
	SysprepunattendXml *string `json:"SysprepunattendXml,omitempty"`
}

func NewSysprepUnattendXml() *SysprepUnattendXml {
	p := new(SysprepUnattendXml)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.SysprepUnattendXml"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.SysprepUnattendXml"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Object encapsulating Task ID Return Value
*/
type Task struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The external identifier that can be used to retrieve the task using its URL.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewTask() *Task {
	p := new(Task)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.Task"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.Task"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
VM Resources Definition.
*/
type VMTemplate struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BootConfig *VMTemplateBootConfig `json:"bootConfig,omitempty"`
	/**
	  The description of the VM.
	*/
	Description *string `json:"description,omitempty"`
	/**
	  The list of Disks attached to the VM.
	*/
	Disks []VMTemplateDisk `json:"disks,omitempty"`
	/**
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/**
	  The list of GPUs attached to the VM.
	*/
	Gpus []VMTemplateGpu `json:"gpus,omitempty"`

	GuestCustomization *VMTemplateGuestCustomization `json:"guestCustomization,omitempty"`
	/**
	  String that identifies the OS running inside of the guest.
	*/
	GuestOsId *string `json:"guestOsId,omitempty"`

	GuestTools *VMTemplateGuestTools `json:"guestTools,omitempty"`
	/**
	  VM's hardware clock timezone in IANA TZDB format (America/Los_Angeles).
	*/
	HardwareClockTimezone *string `json:"hardwareClockTimezone,omitempty"`
	/**
	  Indicates whether the VM is an agent VM. When their host enters maintenance mode, after normal VMs are evacuated, agent VMs are powered off. When the host is restored, agent VMs are powered on before VMs are restored to normalcy. In other words, agent VMs cannot be HA-protected or live migrated.
	*/
	IsAgentVm *bool `json:"isAgentVm,omitempty"`
	/**
	  Indicates whether to remove AHV branding from VM firmware tables.
	*/
	IsDisableBranding *bool `json:"isDisableBranding,omitempty"`
	/**
	  Indicates whether to passthrough the host's CPU features to the guest. Enabling this will disable live migration of the VM.
	*/
	IsEnableCpuPassthrough *bool `json:"isEnableCpuPassthrough,omitempty"`
	/**
	  Indicates whether the vCPUs should be hard pinned to specific pCPUs.
	*/
	IsVcpuHardPinned *bool `json:"isVcpuHardPinned,omitempty"`
	/**
	  Indicates whether VGA console should be enabled or not.
	*/
	IsVgaConsoleEnabled *bool `json:"isVgaConsoleEnabled,omitempty"`
	/**
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`

	MachineType *MachineType `json:"machineType,omitempty"`
	/**
	  Memory size in MiB.
	*/
	MemorySizeMib *int `json:"memorySizeMib,omitempty"`
	/**
	  The name of the VM.
	*/
	Name *string `json:"name,omitempty"`
	/**
	  The list of NICs attached to the VM.
	*/
	Nics []VMTemplateNic `json:"nics,omitempty"`
	/**
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/**
	  Number of logical threads per core.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
	/**
	  Number of vCPUs per socket.
	*/
	NumVcpusPerSocket *int `json:"numVcpusPerSocket,omitempty"`

	ParentReference *Reference `json:"parentReference,omitempty"`

	PowerState *PowerState `json:"powerState,omitempty"`

	PowerStateMechanism *VMTemplatePowerStateMechanism `json:"powerStateMechanism,omitempty"`
	/**
	  Serial ports configured on the VM.
	*/
	SerialPorts []VMTemplateSerialPort `json:"serialPorts,omitempty"`

	StorageConfig *VMTemplateStorage `json:"storageConfig,omitempty"`
	/**
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	VnumaConfig *VMTemplateVNuma `json:"vnumaConfig,omitempty"`
}

func NewVMTemplate() *VMTemplate {
	p := new(VMTemplate)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplate"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplate"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Indicates which device a VM should boot from.
*/
type VMTemplateBootConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	BootDevice *VMTemplateBootDevice `json:"bootDevice,omitempty"`

	BootDeviceOrder []BootDeviceOrder `json:"bootDeviceOrder,omitempty"`

	BootType *BootType `json:"bootType,omitempty"`
}

func NewVMTemplateBootConfig() *VMTemplateBootConfig {
	p := new(VMTemplateBootConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateBootConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateBootConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Indicates which device a VM should boot from. bootDevice takes precedence over bootDeviceOrder. If both are given then specified bootDevice will be primary boot device and remaining devices will be assigned boot order according to bootDeviceOrder field.
*/
type VMTemplateBootDevice struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`

	Config *OneOfVMTemplateBootDeviceConfig `json:"config,omitempty"`
}

func NewVMTemplateBootDevice() *VMTemplateBootDevice {
	p := new(VMTemplateBootDevice)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateBootDevice"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateBootDevice"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *VMTemplateBootDevice) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *VMTemplateBootDevice) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfVMTemplateBootDeviceConfig()
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

/**
Disk attached to the VM.
*/
type VMTemplateDisk struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	DataSourceReference *DataSourceReference `json:"dataSourceReference,omitempty"`

	DeviceProperties *DiskDeviceProperties `json:"deviceProperties,omitempty"`
	/**
	  The device ID which is used to uniquely identify this particular disk.
	*/
	DiskId *string `json:"diskId,omitempty"`
	/**
	  Size of the disk in Bytes.
	*/
	DiskSizeBytes *int64 `json:"diskSizeBytes,omitempty"`
	/**
	  Size of the disk in MiB. Must match the size specified in diskSizeBytes rounded up to the nearest MiB - when that field is present.
	*/
	DiskSizeMib *int `json:"diskSizeMib,omitempty"`

	StorageConfig *VMTemplateDiskStorage `json:"storageConfig,omitempty"`

	VolumeGroupReference *Reference `json:"volumeGroupReference,omitempty"`
}

func NewVMTemplateDisk() *VMTemplateDisk {
	p := new(VMTemplateDisk)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateDisk"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateDisk"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Disk Address
*/
type VMTemplateDiskAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AdapterType *DiskAdapterType `json:"adapterType,omitempty"`

	DeviceIndex *int `json:"deviceIndex,omitempty"`
}

func NewVMTemplateDiskAddress() *VMTemplateDiskAddress {
	p := new(VMTemplateDiskAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateDiskAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateDiskAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
This preference specifies the storage configuration parameters for VM disks.
*/
type VMTemplateDiskStorage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FlashMode *VMTemplateFlashMode `json:"flashMode,omitempty"`

	StorageContainerReference *Reference `json:"storageContainerReference,omitempty"`
}

func NewVMTemplateDiskStorage() *VMTemplateDiskStorage {
	p := new(VMTemplateDiskStorage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateDiskStorage"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateDiskStorage"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
State of the storage policy to pin this virtual disk to the hot tier. It overrides the VM preference when Flash Mode is enabled on the VM of this virtual disk.
*/
type VMTemplateFlashMode int

const (
	VMTEMPLATEFLASHMODE_UNKNOWN  VMTemplateFlashMode = 0
	VMTEMPLATEFLASHMODE_REDACTED VMTemplateFlashMode = 1
	VMTEMPLATEFLASHMODE_ENABLED  VMTemplateFlashMode = 2
	VMTEMPLATEFLASHMODE_DISABLED VMTemplateFlashMode = 3
)

// returns the name of the enum given an ordinal number
func (e *VMTemplateFlashMode) name(index int) string {
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

// returns the enum type given a string value
func (e *VMTemplateFlashMode) index(name string) VMTemplateFlashMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ENABLED",
		"DISABLED",
	}
	for idx := range names {
		if names[idx] == name {
			return VMTemplateFlashMode(idx)
		}
	}
	return VMTEMPLATEFLASHMODE_UNKNOWN
}

func (e *VMTemplateFlashMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VMTemplateFlashMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VMTemplateFlashMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VMTemplateFlashMode) Ref() *VMTemplateFlashMode {
	return &e
}

/**
Graphics resource information for the Virtual Machine.
*/
type VMTemplateGpu struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  The device ID of the GPU.
	*/
	DeviceId *int `json:"deviceId,omitempty"`

	Mode *GpuMode `json:"mode,omitempty"`

	Vendor *GpuVendor `json:"vendor,omitempty"`
}

func NewVMTemplateGpu() *VMTemplateGpu {
	p := new(VMTemplateGpu)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateGpu"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateGpu"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
VM guests may be customized at boot time using one of several different methods. Currently, cloud-init w/ ConfigDriveV2 (for Linux VMs) and Sysprep (for Windows VMs) are supported. Only ONE OF Sysprep or CloudInit should be provided. Note that guest customization can currently only be set during VM creation. Attempting to change it after creation will result in an error. Additional properties can be specified. For example - in the context of VM template creation if 'isOverridable' is set to 'True' then the deployer can upload their own custom script.
*/
type VMTemplateGuestCustomization struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	ConfigItemDiscriminator_ *string `json:"$configItemDiscriminator,omitempty"`

	Config *OneOfVMTemplateGuestCustomizationConfig `json:"config,omitempty"`

	ConfigItemDiscriminator *string `json:"configItemDiscriminator,omitempty"`
	/**
	  Flag to allow override of customization by deployer.
	*/
	IsOverridable *bool `json:"isOverridable,omitempty"`
}

func NewVMTemplateGuestCustomization() *VMTemplateGuestCustomization {
	p := new(VMTemplateGuestCustomization)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateGuestCustomization"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateGuestCustomization"}
	p.UnknownFields_ = map[string]interface{}{}

	p.IsOverridable = new(bool)
	*p.IsOverridable = false

	return p
}

func (p *VMTemplateGuestCustomization) GetConfig() interface{} {
	if nil == p.Config {
		return nil
	}
	return p.Config.GetValue()
}

func (p *VMTemplateGuestCustomization) SetConfig(v interface{}) error {
	if nil == p.Config {
		p.Config = NewOneOfVMTemplateGuestCustomizationConfig()
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

/**
Extra configs related to power state transition.
*/
type VMTemplateGuestPowerStateTransition struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Indicates whether to execute set script before ngt shutdown/reboot.
	*/
	IsEnableScriptExec *bool `json:"isEnableScriptExec,omitempty"`
	/**
	  Indicates whether to abort ngt shutdown/reboot if script fails.
	*/
	IsShouldFailOnScriptFailure *bool `json:"isShouldFailOnScriptFailure,omitempty"`
}

func NewVMTemplateGuestPowerStateTransition() *VMTemplateGuestPowerStateTransition {
	p := new(VMTemplateGuestPowerStateTransition)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateGuestPowerStateTransition"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateGuestPowerStateTransition"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information regarding guest tools.
*/
type VMTemplateGuestTools struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NutanixGuestTools *VMTemplateNutanixGuestTools `json:"nutanixGuestTools,omitempty"`
}

func NewVMTemplateGuestTools() *VMTemplateGuestTools {
	p := new(VMTemplateGuestTools)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateGuestTools"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateGuestTools"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
An IP address.
*/
type VMTemplateIPAddress struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Gateway IP addresses matching the subnet.
	*/
	GatewayAddresses []import1.IPAddress `json:"gatewayAddresses,omitempty"`

	Ip *import1.IPAddress `json:"ip,omitempty"`

	IpType *IPAssignmentType `json:"ipType,omitempty"`
	/**
	  Prefix length for the IP address.
	*/
	PrefixLength *int `json:"prefixLength,omitempty"`

	Type *IPDiscoveredType `json:"type,omitempty"`
}

func NewVMTemplateIPAddress() *VMTemplateIPAddress {
	p := new(VMTemplateIPAddress)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateIPAddress"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateIPAddress"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Virtual Machine NIC.
*/
type VMTemplateNic struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  IP endpoints for the adapter. Currently, IPv4 addresses are supported.
	*/
	IpEndpoints []VMTemplateIPAddress `json:"ipEndpoints,omitempty"`
	/**
	  Whether or not the NIC is connected. True by default.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`

	MacAddress *NicMacAddress `json:"macAddress,omitempty"`

	Model *NicModel `json:"model,omitempty"`

	NetworkFunctionChainReference *Reference `json:"networkFunctionChainReference,omitempty"`

	NetworkFunctionNicType *NetworkFunctionNicType `json:"networkFunctionNicType,omitempty"`
	/**
	  The NIC's UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.
	*/
	NicId *string `json:"nicId,omitempty"`

	NicType *NicType `json:"nicType,omitempty"`

	SubnetReference *Reference `json:"subnetReference,omitempty"`
	/**
	  List of networks to trunk if vlanMode is TRUNKED. If empty and VLAN mode is TRUNKED, all VLANs are trunked.
	*/
	TrunkedVlans []int `json:"trunkedVlans,omitempty"`

	VlanMode *VlanMode `json:"vlanMode,omitempty"`
}

func NewVMTemplateNic() *VMTemplateNic {
	p := new(VMTemplateNic)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateNic"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateNic"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Information regarding Nutanix Guest Tools.
*/
type VMTemplateNutanixGuestTools struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credentials *Credentials `json:"credentials,omitempty"`
	/**
	  Application names that are enabled.
	*/
	EnabledCapabilities []NgtEnabledCapability `json:"enabledCapabilities,omitempty"`

	IsoMountState *IsoMountState `json:"isoMountState,omitempty"`

	NgtState *NgtAvailableState `json:"ngtState,omitempty"`

	State *NgtState `json:"state,omitempty"`
	/**
	  Desired Version of Nutanix Guest Tools installed on the VM.
	*/
	Version *string `json:"version,omitempty"`
}

func NewVMTemplateNutanixGuestTools() *VMTemplateNutanixGuestTools {
	p := new(VMTemplateNutanixGuestTools)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateNutanixGuestTools"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateNutanixGuestTools"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Indicates the mechanism guiding the VM power state transition. Currently used for the transition to OFF state.
*/
type VMTemplatePowerStateMechanism struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	GuestTransitionConfig *VMTemplateGuestPowerStateTransition `json:"guestTransitionConfig,omitempty"`

	Mechanism *PowerStateMechanism `json:"mechanism,omitempty"`
}

func NewVMTemplatePowerStateMechanism() *VMTemplatePowerStateMechanism {
	p := new(VMTemplatePowerStateMechanism)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplatePowerStateMechanism"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplatePowerStateMechanism"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
Indicates the configuration of serial ports of the VM.
*/
type VMTemplateSerialPort struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Index of the serial port.
	*/
	Index *int `json:"index,omitempty"`
	/**
	  Indicates whether the serial port connection is connected or not.
	*/
	IsConnected *bool `json:"isConnected,omitempty"`
}

func NewVMTemplateSerialPort() *VMTemplateSerialPort {
	p := new(VMTemplateSerialPort)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateSerialPort"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateSerialPort"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
User inputs of storage configuration parameters for VMs.
*/
type VMTemplateStorage struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FlashMode *VMTemplateFlashMode `json:"flashMode,omitempty"`

	QosPolicy *VMTemplateStorageQosPolicy `json:"qosPolicy,omitempty"`
}

func NewVMTemplateStorage() *VMTemplateStorage {
	p := new(VMTemplateStorage)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateStorage"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateStorage"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
The policy rule to be enforced depending on the type of policy. For ex - provisioned/throttled iops would be set for a storage qos policy.
*/
type VMTemplateStorageQosPolicy struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Throttled iops for the entities being governed by the policy. The block size for the IO is 32kB.
	*/
	ThrottledIops *int64 `json:"throttledIops"`
}

func (p *VMTemplateStorageQosPolicy) MarshalJSON() ([]byte, error) {
	type VMTemplateStorageQosPolicyProxy VMTemplateStorageQosPolicy
	return json.Marshal(struct {
		*VMTemplateStorageQosPolicyProxy
		ThrottledIops *int64 `json:"throttledIops,omitempty"`
	}{
		VMTemplateStorageQosPolicyProxy: (*VMTemplateStorageQosPolicyProxy)(p),
		ThrottledIops:                   p.ThrottledIops,
	})
}

func NewVMTemplateStorageQosPolicy() *VMTemplateStorageQosPolicy {
	p := new(VMTemplateStorageQosPolicy)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateStorageQosPolicy"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateStorageQosPolicy"}
	p.UnknownFields_ = map[string]interface{}{}

	p.ThrottledIops = new(int64)
	*p.ThrottledIops = -1

	return p
}

/**
Indicates how VM vNUMA should be configured
*/
type VMTemplateVNuma struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Number of vNUMA nodes. 0 means vNUMA is disabled.
	*/
	NumVNumaNodes *int64 `json:"numVNumaNodes,omitempty"`
}

func NewVMTemplateVNuma() *VMTemplateVNuma {
	p := new(VMTemplateVNuma)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.config.VMTemplateVNuma"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.config.VMTemplateVNuma"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/**
VLAN mode.
*/
type VlanMode int

const (
	VLANMODE_UNKNOWN  VlanMode = 0
	VLANMODE_REDACTED VlanMode = 1
	VLANMODE_ACCESS   VlanMode = 2
	VLANMODE_TRUNKED  VlanMode = 3
)

// returns the name of the enum given an ordinal number
func (e *VlanMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNKED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *VlanMode) index(name string) VlanMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ACCESS",
		"TRUNKED",
	}
	for idx := range names {
		if names[idx] == name {
			return VlanMode(idx)
		}
	}
	return VLANMODE_UNKNOWN
}

func (e *VlanMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for VlanMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *VlanMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e VlanMode) Ref() *VlanMode {
	return &e
}

type OneOfVMTemplateGuestCustomizationConfig struct {
	Discriminator *string    `json:"-"`
	ObjectType_   *string    `json:"-"`
	oneOfType0    *Sysprep   `json:"-"`
	oneOfType1    *CloudInit `json:"-"`
}

func NewOneOfVMTemplateGuestCustomizationConfig() *OneOfVMTemplateGuestCustomizationConfig {
	p := new(OneOfVMTemplateGuestCustomizationConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVMTemplateGuestCustomizationConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVMTemplateGuestCustomizationConfig is nil"))
	}
	switch v.(type) {
	case Sysprep:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Sysprep)
		}
		*p.oneOfType0 = v.(Sysprep)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case CloudInit:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(CloudInit)
		}
		*p.oneOfType1 = v.(CloudInit)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVMTemplateGuestCustomizationConfig) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfVMTemplateGuestCustomizationConfig) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Sysprep)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.config.Sysprep" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Sysprep)
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
	vOneOfType1 := new(CloudInit)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "vmm.v4.config.CloudInit" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(CloudInit)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVMTemplateGuestCustomizationConfig"))
}

func (p *OneOfVMTemplateGuestCustomizationConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfVMTemplateGuestCustomizationConfig")
}

type OneOfCloudInitConfig struct {
	Discriminator *string              `json:"-"`
	ObjectType_   *string              `json:"-"`
	oneOfType1    *import1.KVPair      `json:"-"`
	oneOfType0    *CloudInitDataSource `json:"-"`
}

func NewOneOfCloudInitConfig() *OneOfCloudInitConfig {
	p := new(OneOfCloudInitConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCloudInitConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCloudInitConfig is nil"))
	}
	switch v.(type) {
	case import1.KVPair:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import1.KVPair)
		}
		*p.oneOfType1 = v.(import1.KVPair)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case CloudInitDataSource:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(CloudInitDataSource)
		}
		*p.oneOfType0 = v.(CloudInitDataSource)
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

func (p *OneOfCloudInitConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfCloudInitConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(import1.KVPair)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "common.v1.config.KVPair" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import1.KVPair)
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
	vOneOfType0 := new(CloudInitDataSource)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.config.CloudInitDataSource" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(CloudInitDataSource)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCloudInitConfig"))
}

func (p *OneOfCloudInitConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCloudInitConfig")
}

type OneOfVMTemplateBootDeviceConfig struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *VMTemplateDiskAddress `json:"-"`
	oneOfType1    *NicMacAddress         `json:"-"`
}

func NewOneOfVMTemplateBootDeviceConfig() *OneOfVMTemplateBootDeviceConfig {
	p := new(OneOfVMTemplateBootDeviceConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfVMTemplateBootDeviceConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfVMTemplateBootDeviceConfig is nil"))
	}
	switch v.(type) {
	case VMTemplateDiskAddress:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VMTemplateDiskAddress)
		}
		*p.oneOfType0 = v.(VMTemplateDiskAddress)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case NicMacAddress:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(NicMacAddress)
		}
		*p.oneOfType1 = v.(NicMacAddress)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfVMTemplateBootDeviceConfig) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfVMTemplateBootDeviceConfig) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VMTemplateDiskAddress)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.config.VMTemplateDiskAddress" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VMTemplateDiskAddress)
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
	vOneOfType1 := new(NicMacAddress)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "vmm.v4.config.NicMacAddress" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(NicMacAddress)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfVMTemplateBootDeviceConfig"))
}

func (p *OneOfVMTemplateBootDeviceConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfVMTemplateBootDeviceConfig")
}

type OneOfSysprepConfig struct {
	Discriminator *string             `json:"-"`
	ObjectType_   *string             `json:"-"`
	oneOfType1    *import1.KVPair     `json:"-"`
	oneOfType0    *SysprepUnattendXml `json:"-"`
}

func NewOneOfSysprepConfig() *OneOfSysprepConfig {
	p := new(OneOfSysprepConfig)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSysprepConfig) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSysprepConfig is nil"))
	}
	switch v.(type) {
	case import1.KVPair:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(import1.KVPair)
		}
		*p.oneOfType1 = v.(import1.KVPair)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case SysprepUnattendXml:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(SysprepUnattendXml)
		}
		*p.oneOfType0 = v.(SysprepUnattendXml)
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

func (p *OneOfSysprepConfig) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfSysprepConfig) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(import1.KVPair)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "common.v1.config.KVPair" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(import1.KVPair)
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
	vOneOfType0 := new(SysprepUnattendXml)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.config.SysprepUnattendXml" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(SysprepUnattendXml)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSysprepConfig"))
}

func (p *OneOfSysprepConfig) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfSysprepConfig")
}
