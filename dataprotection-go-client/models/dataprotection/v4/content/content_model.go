/*
 * Generated file models/dataprotection/v4/content/content_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Data Protection APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module dataprotection.v4.content of Nutanix Data Protection APIs
*/
package content

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/error"
)

type BaseRecoveryPointSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
}

func NewBaseRecoveryPointSpec() *BaseRecoveryPointSpec {
	p := new(BaseRecoveryPointSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.BaseRecoveryPointSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Changed region comprising of offset, length, and type of the region.
*/
type ChangedRegion struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The length of the region in bytes.
	*/
	Length *int64 `json:"length"`
	/*
	  The start offset of the region in bytes.
	*/
	Offset *int64 `json:"offset"`

	RegionType *RegionType `json:"regionType"`
}

func (p *ChangedRegion) MarshalJSON() ([]byte, error) {
	type ChangedRegionProxy ChangedRegion
	return json.Marshal(struct {
		*ChangedRegionProxy
		Length     *int64      `json:"length,omitempty"`
		Offset     *int64      `json:"offset,omitempty"`
		RegionType *RegionType `json:"regionType,omitempty"`
	}{
		ChangedRegionProxy: (*ChangedRegionProxy)(p),
		Length:             p.Length,
		Offset:             p.Offset,
		RegionType:         p.RegionType,
	})
}

func NewChangedRegion() *ChangedRegion {
	p := new(ChangedRegion)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.ChangedRegion"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0/content/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/disk-recovery-points/{extId}/$actions/compute-changed-regions Post operation
*/
type ChangedVmRegionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangedVmRegionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewChangedVmRegionsApiResponse() *ChangedVmRegionsApiResponse {
	p := new(ChangedVmRegionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.ChangedVmRegionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ChangedVmRegionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ChangedVmRegionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfChangedVmRegionsApiResponseData()
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

/*
REST response for all response codes in API path /dataprotection/v4.0/content/recovery-points/{recoveryPointExtId}/volume-group-recovery-points/{volumeGroupRecoveryPointExtId}/disk-recovery-points/{extId}/$actions/compute-changed-regions Post operation
*/
type ChangedVolumeGroupRegionsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfChangedVolumeGroupRegionsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewChangedVolumeGroupRegionsApiResponse() *ChangedVolumeGroupRegionsApiResponse {
	p := new(ChangedVolumeGroupRegionsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.ChangedVolumeGroupRegionsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ChangedVolumeGroupRegionsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ChangedVolumeGroupRegionsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfChangedVolumeGroupRegionsApiResponseData()
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

/*
Represents operation for which discover cluster api is being called.
*/
type ClusterDiscoverOperation int

const (
	CLUSTERDISCOVEROPERATION_UNKNOWN                 ClusterDiscoverOperation = 0
	CLUSTERDISCOVEROPERATION_REDACTED                ClusterDiscoverOperation = 1
	CLUSTERDISCOVEROPERATION_GET_VSS_METADATA        ClusterDiscoverOperation = 2
	CLUSTERDISCOVEROPERATION_COMPUTE_CHANGED_REGIONS ClusterDiscoverOperation = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterDiscoverOperation) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GET_VSS_METADATA",
		"COMPUTE_CHANGED_REGIONS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterDiscoverOperation) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GET_VSS_METADATA",
		"COMPUTE_CHANGED_REGIONS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterDiscoverOperation) index(name string) ClusterDiscoverOperation {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"GET_VSS_METADATA",
		"COMPUTE_CHANGED_REGIONS",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterDiscoverOperation(idx)
		}
	}
	return CLUSTERDISCOVEROPERATION_UNKNOWN
}

func (e *ClusterDiscoverOperation) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterDiscoverOperation:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterDiscoverOperation) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterDiscoverOperation) Ref() *ClusterDiscoverOperation {
	return &e
}

/*
Request body for discover cluster api containing recovery point details.
*/
type ClusterDiscoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Operation *ClusterDiscoverOperation `json:"operation"`
	/*

	 */
	SpecItemDiscriminator_ *string `json:"$specItemDiscriminator,omitempty"`
	/*
	  Specifications corresponding to the requested operation.
	*/
	Spec *OneOfClusterDiscoverSpecSpec `json:"spec"`
}

func (p *ClusterDiscoverSpec) MarshalJSON() ([]byte, error) {
	type ClusterDiscoverSpecProxy ClusterDiscoverSpec
	return json.Marshal(struct {
		*ClusterDiscoverSpecProxy
		Operation *ClusterDiscoverOperation     `json:"operation,omitempty"`
		Spec      *OneOfClusterDiscoverSpecSpec `json:"spec,omitempty"`
	}{
		ClusterDiscoverSpecProxy: (*ClusterDiscoverSpecProxy)(p),
		Operation:                p.Operation,
		Spec:                     p.Spec,
	})
}

func NewClusterDiscoverSpec() *ClusterDiscoverSpec {
	p := new(ClusterDiscoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.ClusterDiscoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ClusterDiscoverSpec) GetSpec() interface{} {
	if nil == p.Spec {
		return nil
	}
	return p.Spec.GetValue()
}

func (p *ClusterDiscoverSpec) SetSpec(v interface{}) error {
	if nil == p.Spec {
		p.Spec = NewOneOfClusterDiscoverSpecSpec()
	}
	e := p.Spec.SetValue(v)
	if nil == e {
		if nil == p.SpecItemDiscriminator_ {
			p.SpecItemDiscriminator_ = new(string)
		}
		*p.SpecItemDiscriminator_ = *p.Spec.Discriminator
	}
	return e
}

/*
Specs required to fetch the cluster information for compute change region operation.
*/
type ComputeChangedRegionsClusterDiscoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DiskRecoveryPointItemDiscriminator_ *string `json:"$diskRecoveryPointItemDiscriminator,omitempty"`
	/*
	  Disk recovery point reference for compute changed regions operation.
	*/
	DiskRecoveryPoint *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint `json:"diskRecoveryPoint"`
	/*

	 */
	ReferenceDiskRecoveryPointItemDiscriminator_ *string `json:"$referenceDiskRecoveryPointItemDiscriminator,omitempty"`
	/*
	  The disk recovery point reference to compute the changed regions.
	*/
	ReferenceDiskRecoveryPoint *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint `json:"referenceDiskRecoveryPoint,omitempty"`
}

func (p *ComputeChangedRegionsClusterDiscoverSpec) MarshalJSON() ([]byte, error) {
	type ComputeChangedRegionsClusterDiscoverSpecProxy ComputeChangedRegionsClusterDiscoverSpec
	return json.Marshal(struct {
		*ComputeChangedRegionsClusterDiscoverSpecProxy
		DiskRecoveryPoint *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint `json:"diskRecoveryPoint,omitempty"`
	}{
		ComputeChangedRegionsClusterDiscoverSpecProxy: (*ComputeChangedRegionsClusterDiscoverSpecProxy)(p),
		DiskRecoveryPoint: p.DiskRecoveryPoint,
	})
}

func NewComputeChangedRegionsClusterDiscoverSpec() *ComputeChangedRegionsClusterDiscoverSpec {
	p := new(ComputeChangedRegionsClusterDiscoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.ComputeChangedRegionsClusterDiscoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ComputeChangedRegionsClusterDiscoverSpec) GetReferenceDiskRecoveryPoint() interface{} {
	if nil == p.ReferenceDiskRecoveryPoint {
		return nil
	}
	return p.ReferenceDiskRecoveryPoint.GetValue()
}

func (p *ComputeChangedRegionsClusterDiscoverSpec) SetReferenceDiskRecoveryPoint(v interface{}) error {
	if nil == p.ReferenceDiskRecoveryPoint {
		p.ReferenceDiskRecoveryPoint = NewOneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint()
	}
	e := p.ReferenceDiskRecoveryPoint.SetValue(v)
	if nil == e {
		if nil == p.ReferenceDiskRecoveryPointItemDiscriminator_ {
			p.ReferenceDiskRecoveryPointItemDiscriminator_ = new(string)
		}
		*p.ReferenceDiskRecoveryPointItemDiscriminator_ = *p.ReferenceDiskRecoveryPoint.Discriminator
	}
	return e
}

/*
Specs containing disk recovery point information used to discover the cluster.
*/
type DiskRecoveryPointReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	RecoveryPointExtId *string `json:"recoveryPointExtId,omitempty"`
}

func NewDiskRecoveryPointReference() *DiskRecoveryPointReference {
	p := new(DiskRecoveryPointReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.DiskRecoveryPointReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /dataprotection/v4.0/content/recovery-points/{recoveryPointExtId}/vm-recovery-points/{vmRecoveryPointExtId}/vss-metadata Get operation
*/
type GetVssMetadataApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVssMetadataApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVssMetadataApiResponse() *GetVssMetadataApiResponse {
	p := new(GetVssMetadataApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.GetVssMetadataApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVssMetadataApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVssMetadataApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVssMetadataApiResponseData()
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

/*
Specification of get vss metadata operation.
*/
type GetVssMetadataClusterDiscoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier that can be used to retrieve the VM recovery point using its URL.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId"`
}

func (p *GetVssMetadataClusterDiscoverSpec) MarshalJSON() ([]byte, error) {
	type GetVssMetadataClusterDiscoverSpecProxy GetVssMetadataClusterDiscoverSpec
	return json.Marshal(struct {
		*GetVssMetadataClusterDiscoverSpecProxy
		VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
	}{
		GetVssMetadataClusterDiscoverSpecProxy: (*GetVssMetadataClusterDiscoverSpecProxy)(p),
		VmRecoveryPointExtId:                   p.VmRecoveryPointExtId,
	})
}

func NewGetVssMetadataClusterDiscoverSpec() *GetVssMetadataClusterDiscoverSpec {
	p := new(GetVssMetadataClusterDiscoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.GetVssMetadataClusterDiscoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The type of the region.
*/
type RegionType int

const (
	REGIONTYPE_UNKNOWN  RegionType = 0
	REGIONTYPE_REDACTED RegionType = 1
	REGIONTYPE_ZEROED   RegionType = 2
	REGIONTYPE_REGULAR  RegionType = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *RegionType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e RegionType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *RegionType) index(name string) RegionType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"ZEROED",
		"REGULAR",
	}
	for idx := range names {
		if names[idx] == name {
			return RegionType(idx)
		}
	}
	return REGIONTYPE_UNKNOWN
}

func (e *RegionType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for RegionType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *RegionType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e RegionType) Ref() *RegionType {
	return &e
}

/*
Pass all required IDs to generate a token and discover the cluster. All parameters are optional. However, if a reference disk recovery point needs to be set, the following three parameters must be specified: recoveryPointExtId, vmRecoveryPointExtId, and extId.
*/
type VmDiskRecoveryPointClusterDiscoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to identify a VM recovery point.
	*/
	ReferenceVmRecoveryPointExtId *string `json:"referenceVmRecoveryPointExtId,omitempty"`
}

func NewVmDiskRecoveryPointClusterDiscoverSpec() *VmDiskRecoveryPointClusterDiscoverSpec {
	p := new(VmDiskRecoveryPointClusterDiscoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VmDiskRecoveryPointClusterDiscoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to the disk recovery point that's part of a vm recovery point.
*/
type VmDiskRecoveryPointReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	RecoveryPointExtId *string `json:"recoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to identify a VM recovery point.
	*/
	VmRecoveryPointExtId *string `json:"vmRecoveryPointExtId,omitempty"`
}

func NewVmDiskRecoveryPointReference() *VmDiskRecoveryPointReference {
	p := new(VmDiskRecoveryPointReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VmDiskRecoveryPointReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compute changed region parameters. These parameters allow you to specify a start offset, length, block size, and a reference disk recovery point. All parameters are optional. However, if you need to set a reference disk recovery point, you must specify all three parameters: recovery point ID, VM recovery point ID, and disk recovery point ID.
*/
type VmRecoveryPointChangedRegionsComputeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  When blockSizeByte is set, all returned ranges will start and end at blockSize addresses, and the changed blocks will match the block size. Supported values of blockSizeByte are [32768, 65536, 131072, 262144]. Example: if blockSizeByte is set to 32768 (i.e 32KB), all ranges will start at multiple of 32KB and end at multiple of 32KB value. Default blockSizeByte is set to 32KB
	*/
	BlockSizeByte *int64 `json:"blockSizeByte,omitempty"`
	/*
	  The length to compute the changed region. If the value is not provided, the difference is performed from the start offset to the end of the disk. Note: the end offset might automatically align to a system-defined block boundary.
	*/
	Length *int64 `json:"length,omitempty"`
	/*
	  The start offset value to compute the changed region. If the value is not provided, the difference is executed from the offset of 0. Note: the start offset might automatically align to a system-defined block boundary.
	*/
	Offset *int64 `json:"offset,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to identify a VM recovery point.
	*/
	ReferenceVmRecoveryPointExtId *string `json:"referenceVmRecoveryPointExtId,omitempty"`
}

func NewVmRecoveryPointChangedRegionsComputeSpec() *VmRecoveryPointChangedRegionsComputeSpec {
	p := new(VmRecoveryPointChangedRegionsComputeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VmRecoveryPointChangedRegionsComputeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Pass all required IDs to generate a token and discover the cluster. All parameters are optional. However, if a reference disk recovery point needs to be set, the following three parameters must be specified: recoveryPointExtId,volumeGroupRecoveryPointExtId, and extId.
*/
type VolumeGroupDiskRecoveryPointClusterDiscoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group recovery point using its URL (Note: This attribute will be removed in future releases; therefore use the volume group recovery point external identifier instead).
	*/
	ReferenceVolumeGroupRecoveryPointExtId *string `json:"referenceVolumeGroupRecoveryPointExtId,omitempty"`
}

func NewVolumeGroupDiskRecoveryPointClusterDiscoverSpec() *VolumeGroupDiskRecoveryPointClusterDiscoverSpec {
	p := new(VolumeGroupDiskRecoveryPointClusterDiscoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VolumeGroupDiskRecoveryPointClusterDiscoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Specs containing vm disk recovery point information used to discover the cluster.
*/
type VolumeGroupDiskRecoveryPointReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	DiskRecoveryPointExtId *string `json:"diskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	RecoveryPointExtId *string `json:"recoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group recovery point using its URL (Note: This attribute will be removed in future releases; therefore use the volume group recovery point external identifier instead).
	*/
	VolumeGroupRecoveryPointExtId *string `json:"volumeGroupRecoveryPointExtId,omitempty"`
}

func NewVolumeGroupDiskRecoveryPointReference() *VolumeGroupDiskRecoveryPointReference {
	p := new(VolumeGroupDiskRecoveryPointReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VolumeGroupDiskRecoveryPointReference"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Compute changed region parameters. These parameters allow you to specify a start offset, length, block size, and a reference disk recovery point. All parameters are optional. However, if you need to set a reference disk recovery point, you must specify all three parameters: recovery point ID, volume group recovery point ID, and disk recovery point ID.
*/
type VolumeGroupRecoveryPointChangedRegionsComputeSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  When blockSizeByte is set, all returned ranges will start and end at blockSize addresses, and the changed blocks will match the block size. Supported values of blockSizeByte are [32768, 65536, 131072, 262144]. Example: if blockSizeByte is set to 32768 (i.e 32KB), all ranges will start at multiple of 32KB and end at multiple of 32KB value. Default blockSizeByte is set to 32KB
	*/
	BlockSizeByte *int64 `json:"blockSizeByte,omitempty"`
	/*
	  The length to compute the changed region. If the value is not provided, the difference is performed from the start offset to the end of the disk. Note: the end offset might automatically align to a system-defined block boundary.
	*/
	Length *int64 `json:"length,omitempty"`
	/*
	  The start offset value to compute the changed region. If the value is not provided, the difference is executed from the offset of 0. Note: the start offset might automatically align to a system-defined block boundary.
	*/
	Offset *int64 `json:"offset,omitempty"`
	/*
	  Disk recovery point external identifier.
	*/
	ReferenceDiskRecoveryPointExtId *string `json:"referenceDiskRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the recovery point using its URL.
	*/
	ReferenceRecoveryPointExtId *string `json:"referenceRecoveryPointExtId,omitempty"`
	/*
	  The external identifier that can be used to retrieve the volume group recovery point using its URL (Note: This attribute will be removed in future releases; therefore use the volume group recovery point external identifier instead).
	*/
	ReferenceVolumeGroupRecoveryPointExtId *string `json:"referenceVolumeGroupRecoveryPointExtId,omitempty"`
}

func NewVolumeGroupRecoveryPointChangedRegionsComputeSpec() *VolumeGroupRecoveryPointChangedRegionsComputeSpec {
	p := new(VolumeGroupRecoveryPointChangedRegionsComputeSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "dataprotection.v4.content.VolumeGroupRecoveryPointChangedRegionsComputeSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfChangedVmRegionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ChangedRegion        `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfChangedVmRegionsApiResponseData() *OneOfChangedVmRegionsApiResponseData {
	p := new(OneOfChangedVmRegionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfChangedVmRegionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfChangedVmRegionsApiResponseData is nil"))
	}
	switch v.(type) {
	case []ChangedRegion:
		p.oneOfType0 = v.([]ChangedRegion)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.content.ChangedRegion>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.content.ChangedRegion>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfChangedVmRegionsApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.content.ChangedRegion>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangedVmRegionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ChangedRegion)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.content.ChangedRegion" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.content.ChangedRegion>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.content.ChangedRegion>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangedVmRegionsApiResponseData"))
}

func (p *OneOfChangedVmRegionsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.content.ChangedRegion>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangedVmRegionsApiResponseData")
}

type OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint struct {
	Discriminator *string                                `json:"-"`
	ObjectType_   *string                                `json:"-"`
	oneOfType0    *VmDiskRecoveryPointReference          `json:"-"`
	oneOfType1    *VolumeGroupDiskRecoveryPointReference `json:"-"`
}

func NewOneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint() *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint {
	p := new(OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint is nil"))
	}
	switch v.(type) {
	case VmDiskRecoveryPointReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VmDiskRecoveryPointReference)
		}
		*p.oneOfType0 = v.(VmDiskRecoveryPointReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case VolumeGroupDiskRecoveryPointReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(VolumeGroupDiskRecoveryPointReference)
		}
		*p.oneOfType1 = v.(VolumeGroupDiskRecoveryPointReference)
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

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VmDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.content.VmDiskRecoveryPointReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VmDiskRecoveryPointReference)
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
	vOneOfType1 := new(VolumeGroupDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.content.VolumeGroupDiskRecoveryPointReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(VolumeGroupDiskRecoveryPointReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint"))
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfComputeChangedRegionsClusterDiscoverSpecDiskRecoveryPoint")
}

type OneOfClusterDiscoverSpecSpec struct {
	Discriminator *string                                   `json:"-"`
	ObjectType_   *string                                   `json:"-"`
	oneOfType1    *ComputeChangedRegionsClusterDiscoverSpec `json:"-"`
	oneOfType0    *GetVssMetadataClusterDiscoverSpec        `json:"-"`
}

func NewOneOfClusterDiscoverSpecSpec() *OneOfClusterDiscoverSpecSpec {
	p := new(OneOfClusterDiscoverSpecSpec)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfClusterDiscoverSpecSpec) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfClusterDiscoverSpecSpec is nil"))
	}
	switch v.(type) {
	case ComputeChangedRegionsClusterDiscoverSpec:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ComputeChangedRegionsClusterDiscoverSpec)
		}
		*p.oneOfType1 = v.(ComputeChangedRegionsClusterDiscoverSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case GetVssMetadataClusterDiscoverSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(GetVssMetadataClusterDiscoverSpec)
		}
		*p.oneOfType0 = v.(GetVssMetadataClusterDiscoverSpec)
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

func (p *OneOfClusterDiscoverSpecSpec) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfClusterDiscoverSpecSpec) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(ComputeChangedRegionsClusterDiscoverSpec)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.content.ComputeChangedRegionsClusterDiscoverSpec" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ComputeChangedRegionsClusterDiscoverSpec)
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
	vOneOfType0 := new(GetVssMetadataClusterDiscoverSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.content.GetVssMetadataClusterDiscoverSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(GetVssMetadataClusterDiscoverSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfClusterDiscoverSpecSpec"))
}

func (p *OneOfClusterDiscoverSpecSpec) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfClusterDiscoverSpecSpec")
}

type OneOfChangedVolumeGroupRegionsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    []ChangedRegion        `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfChangedVolumeGroupRegionsApiResponseData() *OneOfChangedVolumeGroupRegionsApiResponseData {
	p := new(OneOfChangedVolumeGroupRegionsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfChangedVolumeGroupRegionsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfChangedVolumeGroupRegionsApiResponseData is nil"))
	}
	switch v.(type) {
	case []ChangedRegion:
		p.oneOfType0 = v.([]ChangedRegion)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<dataprotection.v4.content.ChangedRegion>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<dataprotection.v4.content.ChangedRegion>"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfChangedVolumeGroupRegionsApiResponseData) GetValue() interface{} {
	if "List<dataprotection.v4.content.ChangedRegion>" == *p.Discriminator {
		return p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfChangedVolumeGroupRegionsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new([]ChangedRegion)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "dataprotection.v4.content.ChangedRegion" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<dataprotection.v4.content.ChangedRegion>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<dataprotection.v4.content.ChangedRegion>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfChangedVolumeGroupRegionsApiResponseData"))
}

func (p *OneOfChangedVolumeGroupRegionsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<dataprotection.v4.content.ChangedRegion>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfChangedVolumeGroupRegionsApiResponseData")
}

type OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint struct {
	Discriminator *string                                `json:"-"`
	ObjectType_   *string                                `json:"-"`
	oneOfType1    *VolumeGroupDiskRecoveryPointReference `json:"-"`
	oneOfType0    *VmDiskRecoveryPointReference          `json:"-"`
}

func NewOneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint() *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint {
	p := new(OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint is nil"))
	}
	switch v.(type) {
	case VolumeGroupDiskRecoveryPointReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(VolumeGroupDiskRecoveryPointReference)
		}
		*p.oneOfType1 = v.(VolumeGroupDiskRecoveryPointReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType1.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType1.ObjectType_
	case VmDiskRecoveryPointReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VmDiskRecoveryPointReference)
		}
		*p.oneOfType0 = v.(VmDiskRecoveryPointReference)
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

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint) GetValue() interface{} {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(VolumeGroupDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "dataprotection.v4.content.VolumeGroupDiskRecoveryPointReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(VolumeGroupDiskRecoveryPointReference)
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
	vOneOfType0 := new(VmDiskRecoveryPointReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "dataprotection.v4.content.VmDiskRecoveryPointReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VmDiskRecoveryPointReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint"))
}

func (p *OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint) MarshalJSON() ([]byte, error) {
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfComputeChangedRegionsClusterDiscoverSpecReferenceDiskRecoveryPoint")
}

type OneOfGetVssMetadataApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileDetail            `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVssMetadataApiResponseData() *OneOfGetVssMetadataApiResponseData {
	p := new(OneOfGetVssMetadataApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVssMetadataApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVssMetadataApiResponseData is nil"))
	}
	switch v.(type) {
	case FileDetail:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = v.(FileDetail)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
	case import1.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import1.ErrorResponse)
		}
		*p.oneOfType400 = v.(import1.ErrorResponse)
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

func (p *OneOfGetVssMetadataApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVssMetadataApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(FileDetail)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileDetail)
		}
		*p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "FileDetail"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "FileDetail"
		return nil
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "dataprotection.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import1.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVssMetadataApiResponseData"))
}

func (p *OneOfGetVssMetadataApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && "FileDetail" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVssMetadataApiResponseData")
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
