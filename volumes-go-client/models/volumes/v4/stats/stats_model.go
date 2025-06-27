/*
 * Generated file models/volumes/v4/stats/stats_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Volumes APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module volumes.v4.stats of Nutanix Volumes APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/error"
	"time"
)

/*
REST response for all response codes in API path /volumes/v4.1/stats/volume-groups/{volumeGroupExtId}/disks/{extId} Get operation
*/
type GetVolumeDiskStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVolumeDiskStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVolumeDiskStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVolumeDiskStatsApiResponse

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

func (p *GetVolumeDiskStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVolumeDiskStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVolumeDiskStatsApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetVolumeDiskStatsApiResponse() *GetVolumeDiskStatsApiResponse {
	p := new(GetVolumeDiskStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.GetVolumeDiskStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVolumeDiskStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVolumeDiskStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVolumeDiskStatsApiResponseData()
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
REST response for all response codes in API path /volumes/v4.1/stats/volume-groups/{extId} Get operation
*/
type GetVolumeGroupStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVolumeGroupStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVolumeGroupStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVolumeGroupStatsApiResponse

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

func (p *GetVolumeGroupStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVolumeGroupStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVolumeGroupStatsApiResponse(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewGetVolumeGroupStatsApiResponse() *GetVolumeGroupStatsApiResponse {
	p := new(GetVolumeGroupStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.GetVolumeGroupStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVolumeGroupStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVolumeGroupStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVolumeGroupStatsApiResponseData()
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

type TimeValuePair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Timestamp is returned in Epoch format.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  Value of the stat at the corresponding timestamp value represented in Int64 format.
	*/
	Value *int64 `json:"value,omitempty"`
}

func (p *TimeValuePair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TimeValuePair

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

func (p *TimeValuePair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TimeValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = TimeValuePair(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewTimeValuePair() *TimeValuePair {
	p := new(TimeValuePair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.TimeValuePair"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that stores the Volume disk stats.
*/
type VolumeDiskStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Controller average I/O latency measured in microseconds.
	*/
	ControllerAvgIOLatencyUsecs []TimeValuePair `json:"controllerAvgIOLatencyUsecs,omitempty"`
	/*
	  Controller average read I/O latency measured in microseconds.
	*/
	ControllerAvgReadIOLatencyUsecs []TimeValuePair `json:"controllerAvgReadIOLatencyUsecs,omitempty"`
	/*
	  Controller average write I/O latency measured in microseconds.
	*/
	ControllerAvgWriteIOLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIOLatencyUsecs,omitempty"`
	/*
	  Controller I/O bandwidth measured in Kbps.
	*/
	ControllerIOBandwidthKBps []TimeValuePair `json:"controllerIOBandwidthKBps,omitempty"`
	/*
	  Controller I/O rate measured in iops.
	*/
	ControllerNumIOPS []TimeValuePair `json:"controllerNumIOPS,omitempty"`
	/*
	  Controller read I/O measured in iops.
	*/
	ControllerNumReadIOPS []TimeValuePair `json:"controllerNumReadIOPS,omitempty"`
	/*
	  Controller write I/O measured in iops.
	*/
	ControllerNumWriteIOPS []TimeValuePair `json:"controllerNumWriteIOPS,omitempty"`
	/*
	  Controller read I/O bandwidth measured in Kbps.
	*/
	ControllerReadIOBandwidthKBps []TimeValuePair `json:"controllerReadIOBandwidthKBps,omitempty"`
	/*
	  Controller user bytes.
	*/
	ControllerUserBytes []TimeValuePair `json:"controllerUserBytes,omitempty"`
	/*
	  Controller write I/O bandwidth measured in Kbps.
	*/
	ControllerWriteIOBandwidthKBps []TimeValuePair `json:"controllerWriteIOBandwidthKBps,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Uuid of the Volume Disk.
	*/
	VolumeDiskExtId *string `json:"volumeDiskExtId,omitempty"`
}

func (p *VolumeDiskStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeDiskStats

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

func (p *VolumeDiskStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeDiskStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeDiskStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIOLatencyUsecs")
	delete(allFields, "controllerAvgReadIOLatencyUsecs")
	delete(allFields, "controllerAvgWriteIOLatencyUsecs")
	delete(allFields, "controllerIOBandwidthKBps")
	delete(allFields, "controllerNumIOPS")
	delete(allFields, "controllerNumReadIOPS")
	delete(allFields, "controllerNumWriteIOPS")
	delete(allFields, "controllerReadIOBandwidthKBps")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteIOBandwidthKBps")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "volumeDiskExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeDiskStats() *VolumeDiskStats {
	p := new(VolumeDiskStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.VolumeDiskStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VolumeDiskStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Controller average I/O latency measured in microseconds.
	*/
	ControllerAvgIOLatencyUsecs []TimeValuePair `json:"controllerAvgIOLatencyUsecs,omitempty"`
	/*
	  Controller average read I/O latency measured in microseconds.
	*/
	ControllerAvgReadIOLatencyUsecs []TimeValuePair `json:"controllerAvgReadIOLatencyUsecs,omitempty"`
	/*
	  Controller average write I/O latency measured in microseconds.
	*/
	ControllerAvgWriteIOLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIOLatencyUsecs,omitempty"`
	/*
	  Controller I/O bandwidth measured in Kbps.
	*/
	ControllerIOBandwidthKBps []TimeValuePair `json:"controllerIOBandwidthKBps,omitempty"`
	/*
	  Controller I/O rate measured in iops.
	*/
	ControllerNumIOPS []TimeValuePair `json:"controllerNumIOPS,omitempty"`
	/*
	  Controller read I/O measured in iops.
	*/
	ControllerNumReadIOPS []TimeValuePair `json:"controllerNumReadIOPS,omitempty"`
	/*
	  Controller write I/O measured in iops.
	*/
	ControllerNumWriteIOPS []TimeValuePair `json:"controllerNumWriteIOPS,omitempty"`
	/*
	  Controller read I/O bandwidth measured in Kbps.
	*/
	ControllerReadIOBandwidthKBps []TimeValuePair `json:"controllerReadIOBandwidthKBps,omitempty"`
	/*
	  Controller user bytes.
	*/
	ControllerUserBytes []TimeValuePair `json:"controllerUserBytes,omitempty"`
	/*
	  Controller write I/O bandwidth measured in Kbps.
	*/
	ControllerWriteIOBandwidthKBps []TimeValuePair `json:"controllerWriteIOBandwidthKBps,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Uuid of the Volume Disk.
	*/
	VolumeDiskExtId *string `json:"volumeDiskExtId,omitempty"`
}

func (p *VolumeDiskStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeDiskStatsProjection

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

func (p *VolumeDiskStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeDiskStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeDiskStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIOLatencyUsecs")
	delete(allFields, "controllerAvgReadIOLatencyUsecs")
	delete(allFields, "controllerAvgWriteIOLatencyUsecs")
	delete(allFields, "controllerIOBandwidthKBps")
	delete(allFields, "controllerNumIOPS")
	delete(allFields, "controllerNumReadIOPS")
	delete(allFields, "controllerNumWriteIOPS")
	delete(allFields, "controllerReadIOBandwidthKBps")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteIOBandwidthKBps")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "volumeDiskExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeDiskStatsProjection() *VolumeDiskStatsProjection {
	p := new(VolumeDiskStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.VolumeDiskStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A model that stores the Volume group stats.
*/
type VolumeGroupStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Controller average I/O latency measured in microseconds.
	*/
	ControllerAvgIOLatencyUsecs []TimeValuePair `json:"controllerAvgIOLatencyUsecs,omitempty"`
	/*
	  Controller average read I/O latency measured in microseconds.
	*/
	ControllerAvgReadIOLatencyUsecs []TimeValuePair `json:"controllerAvgReadIOLatencyUsecs,omitempty"`
	/*
	  Controller average write I/O latency measured in microseconds.
	*/
	ControllerAvgWriteIOLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIOLatencyUsecs,omitempty"`
	/*
	  Controller I/O bandwidth measured in Kbps.
	*/
	ControllerIOBandwidthKBps []TimeValuePair `json:"controllerIOBandwidthKBps,omitempty"`
	/*
	  Controller I/O rate measured in iops.
	*/
	ControllerNumIOPS []TimeValuePair `json:"controllerNumIOPS,omitempty"`
	/*
	  Controller read I/O measured in iops.
	*/
	ControllerNumReadIOPS []TimeValuePair `json:"controllerNumReadIOPS,omitempty"`
	/*
	  Controller write I/O measured in iops.
	*/
	ControllerNumWriteIOPS []TimeValuePair `json:"controllerNumWriteIOPS,omitempty"`
	/*
	  Controller read I/O bandwidth measured in Kbps.
	*/
	ControllerReadIOBandwidthKBps []TimeValuePair `json:"controllerReadIOBandwidthKBps,omitempty"`
	/*
	  Controller user bytes.
	*/
	ControllerUserBytes []TimeValuePair `json:"controllerUserBytes,omitempty"`
	/*
	  Controller write I/O bandwidth measured in Kbps.
	*/
	ControllerWriteIOBandwidthKBps []TimeValuePair `json:"controllerWriteIOBandwidthKBps,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Uuid of the Volume Group.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func (p *VolumeGroupStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeGroupStats

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

func (p *VolumeGroupStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeGroupStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIOLatencyUsecs")
	delete(allFields, "controllerAvgReadIOLatencyUsecs")
	delete(allFields, "controllerAvgWriteIOLatencyUsecs")
	delete(allFields, "controllerIOBandwidthKBps")
	delete(allFields, "controllerNumIOPS")
	delete(allFields, "controllerNumReadIOPS")
	delete(allFields, "controllerNumWriteIOPS")
	delete(allFields, "controllerReadIOBandwidthKBps")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteIOBandwidthKBps")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeGroupStats() *VolumeGroupStats {
	p := new(VolumeGroupStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.VolumeGroupStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VolumeGroupStatsProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Controller average I/O latency measured in microseconds.
	*/
	ControllerAvgIOLatencyUsecs []TimeValuePair `json:"controllerAvgIOLatencyUsecs,omitempty"`
	/*
	  Controller average read I/O latency measured in microseconds.
	*/
	ControllerAvgReadIOLatencyUsecs []TimeValuePair `json:"controllerAvgReadIOLatencyUsecs,omitempty"`
	/*
	  Controller average write I/O latency measured in microseconds.
	*/
	ControllerAvgWriteIOLatencyUsecs []TimeValuePair `json:"controllerAvgWriteIOLatencyUsecs,omitempty"`
	/*
	  Controller I/O bandwidth measured in Kbps.
	*/
	ControllerIOBandwidthKBps []TimeValuePair `json:"controllerIOBandwidthKBps,omitempty"`
	/*
	  Controller I/O rate measured in iops.
	*/
	ControllerNumIOPS []TimeValuePair `json:"controllerNumIOPS,omitempty"`
	/*
	  Controller read I/O measured in iops.
	*/
	ControllerNumReadIOPS []TimeValuePair `json:"controllerNumReadIOPS,omitempty"`
	/*
	  Controller write I/O measured in iops.
	*/
	ControllerNumWriteIOPS []TimeValuePair `json:"controllerNumWriteIOPS,omitempty"`
	/*
	  Controller read I/O bandwidth measured in Kbps.
	*/
	ControllerReadIOBandwidthKBps []TimeValuePair `json:"controllerReadIOBandwidthKBps,omitempty"`
	/*
	  Controller user bytes.
	*/
	ControllerUserBytes []TimeValuePair `json:"controllerUserBytes,omitempty"`
	/*
	  Controller write I/O bandwidth measured in Kbps.
	*/
	ControllerWriteIOBandwidthKBps []TimeValuePair `json:"controllerWriteIOBandwidthKBps,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Uuid of the Volume Group.
	*/
	VolumeGroupExtId *string `json:"volumeGroupExtId,omitempty"`
}

func (p *VolumeGroupStatsProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VolumeGroupStatsProjection

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

func (p *VolumeGroupStatsProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VolumeGroupStatsProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VolumeGroupStatsProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIOLatencyUsecs")
	delete(allFields, "controllerAvgReadIOLatencyUsecs")
	delete(allFields, "controllerAvgWriteIOLatencyUsecs")
	delete(allFields, "controllerIOBandwidthKBps")
	delete(allFields, "controllerNumIOPS")
	delete(allFields, "controllerNumReadIOPS")
	delete(allFields, "controllerNumWriteIOPS")
	delete(allFields, "controllerReadIOBandwidthKBps")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteIOBandwidthKBps")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantId")
	delete(allFields, "volumeGroupExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVolumeGroupStatsProjection() *VolumeGroupStatsProjection {
	p := new(VolumeGroupStatsProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "volumes.v4.stats.VolumeGroupStatsProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetVolumeGroupStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *VolumeGroupStats      `json:"-"`
}

func NewOneOfGetVolumeGroupStatsApiResponseData() *OneOfGetVolumeGroupStatsApiResponseData {
	p := new(OneOfGetVolumeGroupStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVolumeGroupStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVolumeGroupStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case VolumeGroupStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VolumeGroupStats)
		}
		*p.oneOfType0 = v.(VolumeGroupStats)
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

func (p *OneOfGetVolumeGroupStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetVolumeGroupStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "volumes.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VolumeGroupStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "volumes.v4.stats.VolumeGroupStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VolumeGroupStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeGroupStatsApiResponseData"))
}

func (p *OneOfGetVolumeGroupStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetVolumeGroupStatsApiResponseData")
}

type OneOfGetVolumeDiskStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *VolumeDiskStats       `json:"-"`
}

func NewOneOfGetVolumeDiskStatsApiResponseData() *OneOfGetVolumeDiskStatsApiResponseData {
	p := new(OneOfGetVolumeDiskStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVolumeDiskStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVolumeDiskStatsApiResponseData is nil"))
	}
	switch v.(type) {
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
	case VolumeDiskStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VolumeDiskStats)
		}
		*p.oneOfType0 = v.(VolumeDiskStats)
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

func (p *OneOfGetVolumeDiskStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetVolumeDiskStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "volumes.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(VolumeDiskStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "volumes.v4.stats.VolumeDiskStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VolumeDiskStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVolumeDiskStatsApiResponseData"))
}

func (p *OneOfGetVolumeDiskStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetVolumeDiskStatsApiResponseData")
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
