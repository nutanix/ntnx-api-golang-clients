/*
 * Generated file models/vmm/v4/ahv/stats/stats_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Virtual Machine Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.ahv.stats of Nutanix Virtual Machine Management APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/*
REST response for all response codes in API path /vmm/v4.1/ahv/stats/vms/{vmExtId}/disks/{extId} Get operation
*/
type GetDiskStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetDiskStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetDiskStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetDiskStatsApiResponse

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

func (p *GetDiskStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetDiskStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetDiskStatsApiResponse(*known)

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

func NewGetDiskStatsApiResponse() *GetDiskStatsApiResponse {
	p := new(GetDiskStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.GetDiskStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetDiskStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetDiskStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetDiskStatsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/stats/vms/{vmExtId}/nics/{extId} Get operation
*/
type GetNicStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNicStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetNicStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetNicStatsApiResponse

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

func (p *GetNicStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetNicStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetNicStatsApiResponse(*known)

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

func NewGetNicStatsApiResponse() *GetNicStatsApiResponse {
	p := new(GetNicStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.GetNicStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNicStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNicStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNicStatsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/stats/vms/{extId} Get operation
*/
type GetVmStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetVmStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetVmStatsApiResponse

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

func (p *GetVmStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetVmStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = GetVmStatsApiResponse(*known)

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

func NewGetVmStatsApiResponse() *GetVmStatsApiResponse {
	p := new(GetVmStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.GetVmStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmStatsApiResponseData()
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
REST response for all response codes in API path /vmm/v4.1/ahv/stats/vms Get operation
*/
type ListVmStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListVmStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListVmStatsApiResponse

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

func (p *ListVmStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListVmStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = ListVmStatsApiResponse(*known)

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

func NewListVmStatsApiResponse() *ListVmStatsApiResponse {
	p := new(ListVmStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.ListVmStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmStatsApiResponseData()
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
A collection of VM disk stats.
*/
type VmDiskStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Stats []VmDiskStatsTuple `json:"stats,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The VM external ID associated with the VM Disk stats.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
}

func (p *VmDiskStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskStats

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

func (p *VmDiskStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "stats")
	delete(allFields, "tenantId")
	delete(allFields, "vmExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskStats() *VmDiskStats {
	p := new(VmDiskStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmDiskStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A collection of VM disk stats.
*/
type VmDiskStatsTuple struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM disk controller average I/O latency in microseconds.
	*/
	ControllerAvgIoLatencyMicros *int64 `json:"controllerAvgIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyMicros *int64 `json:"controllerAvgReadIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average read I/O size in kilobytes.
	*/
	ControllerAvgReadIoSizeKb *int64 `json:"controllerAvgReadIoSizeKb,omitempty"`
	/*
	  The VM disk controller average write I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyMicros *int64 `json:"controllerAvgWriteIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average write I/O size in kilobytes.
	*/
	ControllerAvgWriteIoSizeKb *int64 `json:"controllerAvgWriteIoSizeKb,omitempty"`
	/*
	  The VM disk controller number of frontend read operations.
	*/
	ControllerFrontendReadOps *int64 `json:"controllerFrontendReadOps,omitempty"`
	/*
	  The VM disk controller number of frontend write operations.
	*/
	ControllerFrontendWriteOps *int64 `json:"controllerFrontendWriteOps,omitempty"`
	/*
	  The VM disk controller number of I/O bandwidth in kilobytes per second.
	*/
	ControllerIoBandwidthKbps *int64 `json:"controllerIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller number of I/O.
	*/
	ControllerNumIo *int64 `json:"controllerNumIo,omitempty"`
	/*
	  The VM disk controller number of I/O operations per second.
	*/
	ControllerNumIops *int64 `json:"controllerNumIops,omitempty"`
	/*
	  The VM disk controller number of read I/O.
	*/
	ControllerNumReadIo *int64 `json:"controllerNumReadIo,omitempty"`
	/*
	  The VM disk controller number of read I/O operations per second.
	*/
	ControllerNumReadIops *int64 `json:"controllerNumReadIops,omitempty"`
	/*
	  The VM disk controller number of write I/O.
	*/
	ControllerNumWriteIo *int64 `json:"controllerNumWriteIo,omitempty"`
	/*
	  The VM disk controller number of random read operations.
	*/
	ControllerRandomReadOps *int64 `json:"controllerRandomReadOps,omitempty"`
	/*
	  The VM disk controller number of random write operations.
	*/
	ControllerRandomWriteOps *int64 `json:"controllerRandomWriteOps,omitempty"`
	/*
	  The VM disk controller number of read I/O bandwidth in kilobytes per second.
	*/
	ControllerReadIoBandwidthKbps *int64 `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller percentage of read I/O in parts per million.
	*/
	ControllerReadIoPpm *int64 `json:"controllerReadIoPpm,omitempty"`
	/*
	  The VM disk controller percentage of sequential I/O in parts per million.
	*/
	ControllerSeqIoPpm *int64 `json:"controllerSeqIoPpm,omitempty"`
	/*
	  The VM disk controller timespan in microseconds.
	*/
	ControllerTimespanMicros *int64 `json:"controllerTimespanMicros,omitempty"`
	/*
	  The VM disk controller total I/O size in kilobytes.
	*/
	ControllerTotalIoSizeKb *int64 `json:"controllerTotalIoSizeKb,omitempty"`
	/*
	  The VM disk controller total I/O time in microseconds.
	*/
	ControllerTotalIoTimeMicros *int64 `json:"controllerTotalIoTimeMicros,omitempty"`
	/*
	  The VM disk controller total read I/O size in kilobytes.
	*/
	ControllerTotalReadIoSizeKb *int64 `json:"controllerTotalReadIoSizeKb,omitempty"`
	/*
	  The VM disk controller total read I/O time in microseconds.
	*/
	ControllerTotalReadIoTimeMicros *int64 `json:"controllerTotalReadIoTimeMicros,omitempty"`
	/*
	  The VM disk controller user bytes.
	*/
	ControllerUserBytes *int64 `json:"controllerUserBytes,omitempty"`
	/*
	  The VM disk controller vDisk CPU time in microseconds.
	*/
	ControllerVdiskCpuTimeMicros *int64 `json:"controllerVdiskCpuTimeMicros,omitempty"`
	/*
	  The VM disk controller write I/O bandwidth in kilobytes per second.
	*/
	ControllerWriteIoBandwidthKbps *int64 `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller percentage of write I/O in parts per million.
	*/
	ControllerWriteIoPpm *int64 `json:"controllerWriteIoPpm,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmDiskStatsTuple) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskStatsTuple

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

func (p *VmDiskStatsTuple) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskStatsTuple
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskStatsTuple(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoSizeKb")
	delete(allFields, "controllerAvgWriteIoLatencyMicros")
	delete(allFields, "controllerAvgWriteIoSizeKb")
	delete(allFields, "controllerFrontendReadOps")
	delete(allFields, "controllerFrontendWriteOps")
	delete(allFields, "controllerIoBandwidthKbps")
	delete(allFields, "controllerNumIo")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumReadIo")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumWriteIo")
	delete(allFields, "controllerRandomReadOps")
	delete(allFields, "controllerRandomWriteOps")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoPpm")
	delete(allFields, "controllerSeqIoPpm")
	delete(allFields, "controllerTimespanMicros")
	delete(allFields, "controllerTotalIoSizeKb")
	delete(allFields, "controllerTotalIoTimeMicros")
	delete(allFields, "controllerTotalReadIoSizeKb")
	delete(allFields, "controllerTotalReadIoTimeMicros")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerVdiskCpuTimeMicros")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoPpm")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskStatsTuple() *VmDiskStatsTuple {
	p := new(VmDiskStatsTuple)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmDiskStatsTuple"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmDiskStatsTupleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM disk controller average I/O latency in microseconds.
	*/
	ControllerAvgIoLatencyMicros *int64 `json:"controllerAvgIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyMicros *int64 `json:"controllerAvgReadIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average read I/O size in kilobytes.
	*/
	ControllerAvgReadIoSizeKb *int64 `json:"controllerAvgReadIoSizeKb,omitempty"`
	/*
	  The VM disk controller average write I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyMicros *int64 `json:"controllerAvgWriteIoLatencyMicros,omitempty"`
	/*
	  The VM disk controller average write I/O size in kilobytes.
	*/
	ControllerAvgWriteIoSizeKb *int64 `json:"controllerAvgWriteIoSizeKb,omitempty"`
	/*
	  The VM disk controller number of frontend read operations.
	*/
	ControllerFrontendReadOps *int64 `json:"controllerFrontendReadOps,omitempty"`
	/*
	  The VM disk controller number of frontend write operations.
	*/
	ControllerFrontendWriteOps *int64 `json:"controllerFrontendWriteOps,omitempty"`
	/*
	  The VM disk controller number of I/O bandwidth in kilobytes per second.
	*/
	ControllerIoBandwidthKbps *int64 `json:"controllerIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller number of I/O.
	*/
	ControllerNumIo *int64 `json:"controllerNumIo,omitempty"`
	/*
	  The VM disk controller number of I/O operations per second.
	*/
	ControllerNumIops *int64 `json:"controllerNumIops,omitempty"`
	/*
	  The VM disk controller number of read I/O.
	*/
	ControllerNumReadIo *int64 `json:"controllerNumReadIo,omitempty"`
	/*
	  The VM disk controller number of read I/O operations per second.
	*/
	ControllerNumReadIops *int64 `json:"controllerNumReadIops,omitempty"`
	/*
	  The VM disk controller number of write I/O.
	*/
	ControllerNumWriteIo *int64 `json:"controllerNumWriteIo,omitempty"`
	/*
	  The VM disk controller number of random read operations.
	*/
	ControllerRandomReadOps *int64 `json:"controllerRandomReadOps,omitempty"`
	/*
	  The VM disk controller number of random write operations.
	*/
	ControllerRandomWriteOps *int64 `json:"controllerRandomWriteOps,omitempty"`
	/*
	  The VM disk controller number of read I/O bandwidth in kilobytes per second.
	*/
	ControllerReadIoBandwidthKbps *int64 `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller percentage of read I/O in parts per million.
	*/
	ControllerReadIoPpm *int64 `json:"controllerReadIoPpm,omitempty"`
	/*
	  The VM disk controller percentage of sequential I/O in parts per million.
	*/
	ControllerSeqIoPpm *int64 `json:"controllerSeqIoPpm,omitempty"`
	/*
	  The VM disk controller timespan in microseconds.
	*/
	ControllerTimespanMicros *int64 `json:"controllerTimespanMicros,omitempty"`
	/*
	  The VM disk controller total I/O size in kilobytes.
	*/
	ControllerTotalIoSizeKb *int64 `json:"controllerTotalIoSizeKb,omitempty"`
	/*
	  The VM disk controller total I/O time in microseconds.
	*/
	ControllerTotalIoTimeMicros *int64 `json:"controllerTotalIoTimeMicros,omitempty"`
	/*
	  The VM disk controller total read I/O size in kilobytes.
	*/
	ControllerTotalReadIoSizeKb *int64 `json:"controllerTotalReadIoSizeKb,omitempty"`
	/*
	  The VM disk controller total read I/O time in microseconds.
	*/
	ControllerTotalReadIoTimeMicros *int64 `json:"controllerTotalReadIoTimeMicros,omitempty"`
	/*
	  The VM disk controller user bytes.
	*/
	ControllerUserBytes *int64 `json:"controllerUserBytes,omitempty"`
	/*
	  The VM disk controller vDisk CPU time in microseconds.
	*/
	ControllerVdiskCpuTimeMicros *int64 `json:"controllerVdiskCpuTimeMicros,omitempty"`
	/*
	  The VM disk controller write I/O bandwidth in kilobytes per second.
	*/
	ControllerWriteIoBandwidthKbps *int64 `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  The VM disk controller percentage of write I/O in parts per million.
	*/
	ControllerWriteIoPpm *int64 `json:"controllerWriteIoPpm,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmDiskStatsTupleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmDiskStatsTupleProjection

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

func (p *VmDiskStatsTupleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmDiskStatsTupleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmDiskStatsTupleProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "controllerAvgIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoSizeKb")
	delete(allFields, "controllerAvgWriteIoLatencyMicros")
	delete(allFields, "controllerAvgWriteIoSizeKb")
	delete(allFields, "controllerFrontendReadOps")
	delete(allFields, "controllerFrontendWriteOps")
	delete(allFields, "controllerIoBandwidthKbps")
	delete(allFields, "controllerNumIo")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumReadIo")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumWriteIo")
	delete(allFields, "controllerRandomReadOps")
	delete(allFields, "controllerRandomWriteOps")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoPpm")
	delete(allFields, "controllerSeqIoPpm")
	delete(allFields, "controllerTimespanMicros")
	delete(allFields, "controllerTotalIoSizeKb")
	delete(allFields, "controllerTotalIoTimeMicros")
	delete(allFields, "controllerTotalReadIoSizeKb")
	delete(allFields, "controllerTotalReadIoTimeMicros")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerVdiskCpuTimeMicros")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoPpm")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmDiskStatsTupleProjection() *VmDiskStatsTupleProjection {
	p := new(VmDiskStatsTupleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmDiskStatsTupleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A collection of VM NIC stats.
*/
type VmNicStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Stats []VmNicStatsTuple `json:"stats,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The VM external ID associated with the VM NIC stats.
	*/
	VmExtId *string `json:"vmExtId,omitempty"`
}

func (p *VmNicStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmNicStats

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

func (p *VmNicStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmNicStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmNicStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "stats")
	delete(allFields, "tenantId")
	delete(allFields, "vmExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmNicStats() *VmNicStats {
	p := new(VmNicStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmNicStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A collection of VM NIC stats.
*/
type VmNicStatsTuple struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM NIC number of dropped received packets.
	*/
	NetworkDroppedReceivedPackets *int64 `json:"networkDroppedReceivedPackets,omitempty"`
	/*
	  The VM NIC number of dropped transmitted packets.
	*/
	NetworkDroppedTransmittedPackets *int64 `json:"networkDroppedTransmittedPackets,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmNicStatsTuple) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmNicStatsTuple

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

func (p *VmNicStatsTuple) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmNicStatsTuple
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmNicStatsTuple(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkDroppedReceivedPackets")
	delete(allFields, "networkDroppedTransmittedPackets")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmNicStatsTuple() *VmNicStatsTuple {
	p := new(VmNicStatsTuple)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmNicStatsTuple"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmNicStatsTupleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM NIC number of dropped received packets.
	*/
	NetworkDroppedReceivedPackets *int64 `json:"networkDroppedReceivedPackets,omitempty"`
	/*
	  The VM NIC number of dropped transmitted packets.
	*/
	NetworkDroppedTransmittedPackets *int64 `json:"networkDroppedTransmittedPackets,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmNicStatsTupleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmNicStatsTupleProjection

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

func (p *VmNicStatsTupleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmNicStatsTupleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmNicStatsTupleProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "networkDroppedReceivedPackets")
	delete(allFields, "networkDroppedTransmittedPackets")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmNicStatsTupleProjection() *VmNicStatsTupleProjection {
	p := new(VmNicStatsTupleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmNicStatsTupleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A collection of VM stats.
*/
type VmStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Stats []VmStatsTuple `json:"stats,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *VmStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmStats

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

func (p *VmStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmStats(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "stats")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmStats() *VmStats {
	p := new(VmStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A collection of VM stats.
*/
type VmStatsTuple struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM NCC health check score.
	*/
	CheckScore *int64 `json:"checkScore,omitempty"`
	/*
	  The UUID of the cluster on which the VM resides.
	*/
	Cluster *string `json:"cluster,omitempty"`
	/*
	  The VM controller average I/O latency in microseconds.
	*/
	ControllerAvgIoLatencyMicros *int64 `json:"controllerAvgIoLatencyMicros,omitempty"`
	/*
	  The VM controller average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyMicros *int64 `json:"controllerAvgReadIoLatencyMicros,omitempty"`
	/*
	  The VM controller average read I/O size in kilobytes.
	*/
	ControllerAvgReadIoSizeKb *int64 `json:"controllerAvgReadIoSizeKb,omitempty"`
	/*
	  The VM controller average write I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyMicros *int64 `json:"controllerAvgWriteIoLatencyMicros,omitempty"`
	/*
	  The VM controller average write I/O size in kilobytes.
	*/
	ControllerAvgWriteIoSizeKb *int64 `json:"controllerAvgWriteIoSizeKb,omitempty"`
	/*
	  The VM controller I/O bandwidth in kilobytes per second.
	*/
	ControllerIoBandwidthKbps *int64 `json:"controllerIoBandwidthKbps,omitempty"`
	/*
	  The VM controller number of I/O requests.
	*/
	ControllerNumIo *int64 `json:"controllerNumIo,omitempty"`
	/*
	  The VM controller number of I/O operations per second.
	*/
	ControllerNumIops *int64 `json:"controllerNumIops,omitempty"`
	/*
	  The VM controller number of random I/O.
	*/
	ControllerNumRandomIo *int64 `json:"controllerNumRandomIo,omitempty"`
	/*
	  The VM controller number of read I/O.
	*/
	ControllerNumReadIo *int64 `json:"controllerNumReadIo,omitempty"`
	/*
	  The VM controller number of read I/O operations per second.
	*/
	ControllerNumReadIops *int64 `json:"controllerNumReadIops,omitempty"`
	/*
	  The VM controller number of sequential I/Os.
	*/
	ControllerNumSeqIo *int64 `json:"controllerNumSeqIo,omitempty"`
	/*
	  The VM controller number of write I/O.
	*/
	ControllerNumWriteIo *int64 `json:"controllerNumWriteIo,omitempty"`
	/*
	  The VM controller number of write I/O operations per second.
	*/
	ControllerNumWriteIops *int64 `json:"controllerNumWriteIops,omitempty"`
	/*
	  Amount of data in bytes drained to HDD tier of extent store by the VM controller.
	*/
	ControllerOplogDrainDestHddBytes *int64 `json:"controllerOplogDrainDestHddBytes,omitempty"`
	/*
	  Amount of data in bytes drained to SSD tier of extent store by the VM controller.
	*/
	ControllerOplogDrainDestSsdBytes *int64 `json:"controllerOplogDrainDestSsdBytes,omitempty"`
	/*
	  The VM controller number of random I/O PPM.
	*/
	ControllerRandomIoPpm *int64 `json:"controllerRandomIoPpm,omitempty"`
	/*
	  The VM controller number of read I/O bandwidth in kilobytes per second.
	*/
	ControllerReadIoBandwidthKbps *int64 `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  The VM controller number of read I/O PPM.
	*/
	ControllerReadIoPpm *int64 `json:"controllerReadIoPpm,omitempty"`
	/*
	  The VM disk controller number of bytes for local read source: HDD estore.
	*/
	ControllerReadSourceEstoreHddLocalBytes *int64 `json:"controllerReadSourceEstoreHddLocalBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for remote read source: HDD estore.
	*/
	ControllerReadSourceEstoreHddRemoteBytes *int64 `json:"controllerReadSourceEstoreHddRemoteBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for local read source: SSD estore.
	*/
	ControllerReadSourceEstoreSsdLocalBytes *int64 `json:"controllerReadSourceEstoreSsdLocalBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for remote read source: SSD estore.
	*/
	ControllerReadSourceEstoreSsdRemoteBytes *int64 `json:"controllerReadSourceEstoreSsdRemoteBytes,omitempty"`
	/*
	  The VM controller read source oplog in bytes.
	*/
	ControllerReadSourceOplogBytes *int64 `json:"controllerReadSourceOplogBytes,omitempty"`
	/*
	  The VM controller number of sequential I/O PPM.
	*/
	ControllerSeqIoPpm *int64 `json:"controllerSeqIoPpm,omitempty"`
	/*
	  The VM controller shared usage in bytes.
	*/
	ControllerSharedUsageBytes *int64 `json:"controllerSharedUsageBytes,omitempty"`
	/*
	  The VM controller snapshot usage in bytes.
	*/
	ControllerSnapshotUsageBytes *int64 `json:"controllerSnapshotUsageBytes,omitempty"`
	/*
	  The VM controller total usage on SSD tier for the VM.
	*/
	ControllerStorageTierSsdUsageBytes *int64 `json:"controllerStorageTierSsdUsageBytes,omitempty"`
	/*
	  The VM controller timespan in microseconds.
	*/
	ControllerTimespanMicros *int64 `json:"controllerTimespanMicros,omitempty"`
	/*
	  The VM controller number of total I/O size in kilobytes.
	*/
	ControllerTotalIoSizeKb *int64 `json:"controllerTotalIoSizeKb,omitempty"`
	/*
	  The VM controller number of total I/O time in microseconds.
	*/
	ControllerTotalIoTimeMicros *int64 `json:"controllerTotalIoTimeMicros,omitempty"`
	/*
	  The VM controller number of total read I/O size in kilobytes.
	*/
	ControllerTotalReadIoSizeKb *int64 `json:"controllerTotalReadIoSizeKb,omitempty"`
	/*
	  The VM controller number of total read I/O time in microseconds.
	*/
	ControllerTotalReadIoTimeMicros *int64 `json:"controllerTotalReadIoTimeMicros,omitempty"`
	/*
	  The VM controller number of total transformed usage in bytes.
	*/
	ControllerTotalTransformedUsageBytes *int64 `json:"controllerTotalTransformedUsageBytes,omitempty"`
	/*
	  The VM controller user bytes.
	*/
	ControllerUserBytes *int64 `json:"controllerUserBytes,omitempty"`
	/*
	  The VM controller write destination HDD estore in bytes.
	*/
	ControllerWriteDestEstoreHddBytes *int64 `json:"controllerWriteDestEstoreHddBytes,omitempty"`
	/*
	  The VM controller write destination SSD estore in bytes.
	*/
	ControllerWriteDestEstoreSsdBytes *int64 `json:"controllerWriteDestEstoreSsdBytes,omitempty"`
	/*
	  The VM controller write I/O bandwidth in kilobytes per second.
	*/
	ControllerWriteIoBandwidthKbps *int64 `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  The VM controller percentage of write I/O in parts per million.
	*/
	ControllerWriteIoPpm *int64 `json:"controllerWriteIoPpm,omitempty"`
	/*
	  The read working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondReadMb *int64 `json:"controllerWss120SecondReadMb,omitempty"`
	/*
	  The combined working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondUnionMb *int64 `json:"controllerWss120SecondUnionMb,omitempty"`
	/*
	  The write working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondWriteMb *int64 `json:"controllerWss120SecondWriteMb,omitempty"`
	/*
	  The read working set size of the VM disk controller in megabytes, 1 hour measurement interval.
	*/
	ControllerWss3600SecondReadMb *int64 `json:"controllerWss3600SecondReadMb,omitempty"`
	/*
	  The combined working set size of the VM disk controller in megabytes, 1-hour measurement interval.
	*/
	ControllerWss3600SecondUnionMb *int64 `json:"controllerWss3600SecondUnionMb,omitempty"`
	/*
	  The write working set size of the VM disk controller in megabytes, 1-hour measurement interval.
	*/
	ControllerWss3600SecondWriteMb *int64 `json:"controllerWss3600SecondWriteMb,omitempty"`
	/*
	  The VM disk capacity in bytes.
	*/
	DiskCapacityBytes *int64 `json:"diskCapacityBytes,omitempty"`
	/*
	  The VM disk usage in PPM.
	*/
	DiskUsagePpm *int64 `json:"diskUsagePpm,omitempty"`
	/*
	  The VM frame buffer usage in PPM.
	*/
	FrameBufferUsagePpm *int64 `json:"frameBufferUsagePpm,omitempty"`
	/*
	  The VM GPU usage in PPM.
	*/
	GpuUsagePpm *int64 `json:"gpuUsagePpm,omitempty"`
	/*
	  The VM memory usage in percentage reported by a guest.
	*/
	GuestMemoryUsagePpm *int64 `json:"guestMemoryUsagePpm,omitempty"`
	/*
	  The average I/O latency of the VM in microseconds
	*/
	HypervisorAvgIoLatencyMicros *int64 `json:"hypervisorAvgIoLatencyMicros,omitempty"`
	/*
	  Percentage of time that the VM was ready, but could not get scheduled to run.
	*/
	HypervisorCpuReadyTimePpm *int64 `json:"hypervisorCpuReadyTimePpm,omitempty"`
	/*
	  The CPU usage of the VM in parts per million.
	*/
	HypervisorCpuUsagePpm *int64 `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  The I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorIoBandwidthKbps *int64 `json:"hypervisorIoBandwidthKbps,omitempty"`
	/*
	  The target memory balloon of the VM to be reclaimed in bytes.
	*/
	HypervisorMemoryBalloonReclaimTargetBytes *int64 `json:"hypervisorMemoryBalloonReclaimTargetBytes,omitempty"`
	/*
	  The memory balloon reclaimed by the VM in bytes.
	*/
	HypervisorMemoryBalloonReclaimedBytes *int64 `json:"hypervisorMemoryBalloonReclaimedBytes,omitempty"`
	/*
	  Consolidated guest memory usage in percentage.
	*/
	HypervisorMemoryUsagePpm *int64 `json:"hypervisorMemoryUsagePpm,omitempty"`
	/*
	  The number of I/O by the VM.
	*/
	HypervisorNumIo *int64 `json:"hypervisorNumIo,omitempty"`
	/*
	  The number of I/O operations by the VM per second.
	*/
	HypervisorNumIops *int64 `json:"hypervisorNumIops,omitempty"`
	/*
	  The number of read I/O operations by the VM.
	*/
	HypervisorNumReadIo *int64 `json:"hypervisorNumReadIo,omitempty"`
	/*
	  The number of read I/O operations by the VM per second.
	*/
	HypervisorNumReadIops *int64 `json:"hypervisorNumReadIops,omitempty"`
	/*
	  The number of packets dropped by the VM during reception.
	*/
	HypervisorNumReceivePacketsDropped *int64 `json:"hypervisorNumReceivePacketsDropped,omitempty"`
	/*
	  The number of bytes received by the VM.
	*/
	HypervisorNumReceivedBytes *int64 `json:"hypervisorNumReceivedBytes,omitempty"`
	/*
	  The number of packets dropped by the VM during transmission.
	*/
	HypervisorNumTransmitPacketsDropped *int64 `json:"hypervisorNumTransmitPacketsDropped,omitempty"`
	/*
	  The number of bytes transmitted by the VM.
	*/
	HypervisorNumTransmittedBytes *int64 `json:"hypervisorNumTransmittedBytes,omitempty"`
	/*
	  The number of write I/O by the VM.
	*/
	HypervisorNumWriteIo *int64 `json:"hypervisorNumWriteIo,omitempty"`
	/*
	  The number of write I/O operations by the VM per second.
	*/
	HypervisorNumWriteIops *int64 `json:"hypervisorNumWriteIops,omitempty"`
	/*
	  The number of read I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorReadIoBandwidthKbps *int64 `json:"hypervisorReadIoBandwidthKbps,omitempty"`
	/*
	  The swap in rate of the VM in kilobytes per second.
	*/
	HypervisorSwapInRateKbps *int64 `json:"hypervisorSwapInRateKbps,omitempty"`
	/*
	  The swap out rate of the VM in kilobytes per second.
	*/
	HypervisorSwapOutRateKbps *int64 `json:"hypervisorSwapOutRateKbps,omitempty"`
	/*
	  The timespan of the VM in microseconds.
	*/
	HypervisorTimespanMicros *int64 `json:"hypervisorTimespanMicros,omitempty"`
	/*
	  The total I/O size of the VM in kilobytes.
	*/
	HypervisorTotalIoSizeKb *int64 `json:"hypervisorTotalIoSizeKb,omitempty"`
	/*
	  The total I/O time of the VM in microseconds.
	*/
	HypervisorTotalIoTimeMicros *int64 `json:"hypervisorTotalIoTimeMicros,omitempty"`
	/*
	  The total read I/O size of the VM in kilobytes.
	*/
	HypervisorTotalReadIoSizeKb *int64 `json:"hypervisorTotalReadIoSizeKb,omitempty"`
	/*
	  Hypervisor type of the VM.
	*/
	HypervisorType *string `json:"hypervisorType,omitempty"`
	/*
	  The running time of the VM in microseconds.
	*/
	HypervisorVmRunningTimeUsecs *int64 `json:"hypervisorVmRunningTimeUsecs,omitempty"`
	/*
	  The write I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorWriteIoBandwidthKbps *int64 `json:"hypervisorWriteIoBandwidthKbps,omitempty"`
	/*
	  The reserved memory of the VM in bytes.
	*/
	MemoryReservedBytes *int64 `json:"memoryReservedBytes,omitempty"`
	/*
	  The VM memory usage bytes.
	*/
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	/*
	  The VM memory usage in PPM.
	*/
	MemoryUsagePpm *int64 `json:"memoryUsagePpm,omitempty"`
	/*
	  The VM number of vCPUs used in PPM.
	*/
	NumVcpusUsedPpm *int64 `json:"numVcpusUsedPpm,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmStatsTuple) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmStatsTuple

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

func (p *VmStatsTuple) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmStatsTuple
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmStatsTuple(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "checkScore")
	delete(allFields, "cluster")
	delete(allFields, "controllerAvgIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoSizeKb")
	delete(allFields, "controllerAvgWriteIoLatencyMicros")
	delete(allFields, "controllerAvgWriteIoSizeKb")
	delete(allFields, "controllerIoBandwidthKbps")
	delete(allFields, "controllerNumIo")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumRandomIo")
	delete(allFields, "controllerNumReadIo")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumSeqIo")
	delete(allFields, "controllerNumWriteIo")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerOplogDrainDestHddBytes")
	delete(allFields, "controllerOplogDrainDestSsdBytes")
	delete(allFields, "controllerRandomIoPpm")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoPpm")
	delete(allFields, "controllerReadSourceEstoreHddLocalBytes")
	delete(allFields, "controllerReadSourceEstoreHddRemoteBytes")
	delete(allFields, "controllerReadSourceEstoreSsdLocalBytes")
	delete(allFields, "controllerReadSourceEstoreSsdRemoteBytes")
	delete(allFields, "controllerReadSourceOplogBytes")
	delete(allFields, "controllerSeqIoPpm")
	delete(allFields, "controllerSharedUsageBytes")
	delete(allFields, "controllerSnapshotUsageBytes")
	delete(allFields, "controllerStorageTierSsdUsageBytes")
	delete(allFields, "controllerTimespanMicros")
	delete(allFields, "controllerTotalIoSizeKb")
	delete(allFields, "controllerTotalIoTimeMicros")
	delete(allFields, "controllerTotalReadIoSizeKb")
	delete(allFields, "controllerTotalReadIoTimeMicros")
	delete(allFields, "controllerTotalTransformedUsageBytes")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteDestEstoreHddBytes")
	delete(allFields, "controllerWriteDestEstoreSsdBytes")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoPpm")
	delete(allFields, "controllerWss120SecondReadMb")
	delete(allFields, "controllerWss120SecondUnionMb")
	delete(allFields, "controllerWss120SecondWriteMb")
	delete(allFields, "controllerWss3600SecondReadMb")
	delete(allFields, "controllerWss3600SecondUnionMb")
	delete(allFields, "controllerWss3600SecondWriteMb")
	delete(allFields, "diskCapacityBytes")
	delete(allFields, "diskUsagePpm")
	delete(allFields, "frameBufferUsagePpm")
	delete(allFields, "gpuUsagePpm")
	delete(allFields, "guestMemoryUsagePpm")
	delete(allFields, "hypervisorAvgIoLatencyMicros")
	delete(allFields, "hypervisorCpuReadyTimePpm")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorIoBandwidthKbps")
	delete(allFields, "hypervisorMemoryBalloonReclaimTargetBytes")
	delete(allFields, "hypervisorMemoryBalloonReclaimedBytes")
	delete(allFields, "hypervisorMemoryUsagePpm")
	delete(allFields, "hypervisorNumIo")
	delete(allFields, "hypervisorNumIops")
	delete(allFields, "hypervisorNumReadIo")
	delete(allFields, "hypervisorNumReadIops")
	delete(allFields, "hypervisorNumReceivePacketsDropped")
	delete(allFields, "hypervisorNumReceivedBytes")
	delete(allFields, "hypervisorNumTransmitPacketsDropped")
	delete(allFields, "hypervisorNumTransmittedBytes")
	delete(allFields, "hypervisorNumWriteIo")
	delete(allFields, "hypervisorNumWriteIops")
	delete(allFields, "hypervisorReadIoBandwidthKbps")
	delete(allFields, "hypervisorSwapInRateKbps")
	delete(allFields, "hypervisorSwapOutRateKbps")
	delete(allFields, "hypervisorTimespanMicros")
	delete(allFields, "hypervisorTotalIoSizeKb")
	delete(allFields, "hypervisorTotalIoTimeMicros")
	delete(allFields, "hypervisorTotalReadIoSizeKb")
	delete(allFields, "hypervisorType")
	delete(allFields, "hypervisorVmRunningTimeUsecs")
	delete(allFields, "hypervisorWriteIoBandwidthKbps")
	delete(allFields, "memoryReservedBytes")
	delete(allFields, "memoryUsageBytes")
	delete(allFields, "memoryUsagePpm")
	delete(allFields, "numVcpusUsedPpm")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmStatsTuple() *VmStatsTuple {
	p := new(VmStatsTuple)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmStatsTuple"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type VmStatsTupleProjection struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The VM NCC health check score.
	*/
	CheckScore *int64 `json:"checkScore,omitempty"`
	/*
	  The UUID of the cluster on which the VM resides.
	*/
	Cluster *string `json:"cluster,omitempty"`
	/*
	  The VM controller average I/O latency in microseconds.
	*/
	ControllerAvgIoLatencyMicros *int64 `json:"controllerAvgIoLatencyMicros,omitempty"`
	/*
	  The VM controller average read I/O latency in microseconds.
	*/
	ControllerAvgReadIoLatencyMicros *int64 `json:"controllerAvgReadIoLatencyMicros,omitempty"`
	/*
	  The VM controller average read I/O size in kilobytes.
	*/
	ControllerAvgReadIoSizeKb *int64 `json:"controllerAvgReadIoSizeKb,omitempty"`
	/*
	  The VM controller average write I/O latency in microseconds.
	*/
	ControllerAvgWriteIoLatencyMicros *int64 `json:"controllerAvgWriteIoLatencyMicros,omitempty"`
	/*
	  The VM controller average write I/O size in kilobytes.
	*/
	ControllerAvgWriteIoSizeKb *int64 `json:"controllerAvgWriteIoSizeKb,omitempty"`
	/*
	  The VM controller I/O bandwidth in kilobytes per second.
	*/
	ControllerIoBandwidthKbps *int64 `json:"controllerIoBandwidthKbps,omitempty"`
	/*
	  The VM controller number of I/O requests.
	*/
	ControllerNumIo *int64 `json:"controllerNumIo,omitempty"`
	/*
	  The VM controller number of I/O operations per second.
	*/
	ControllerNumIops *int64 `json:"controllerNumIops,omitempty"`
	/*
	  The VM controller number of random I/O.
	*/
	ControllerNumRandomIo *int64 `json:"controllerNumRandomIo,omitempty"`
	/*
	  The VM controller number of read I/O.
	*/
	ControllerNumReadIo *int64 `json:"controllerNumReadIo,omitempty"`
	/*
	  The VM controller number of read I/O operations per second.
	*/
	ControllerNumReadIops *int64 `json:"controllerNumReadIops,omitempty"`
	/*
	  The VM controller number of sequential I/Os.
	*/
	ControllerNumSeqIo *int64 `json:"controllerNumSeqIo,omitempty"`
	/*
	  The VM controller number of write I/O.
	*/
	ControllerNumWriteIo *int64 `json:"controllerNumWriteIo,omitempty"`
	/*
	  The VM controller number of write I/O operations per second.
	*/
	ControllerNumWriteIops *int64 `json:"controllerNumWriteIops,omitempty"`
	/*
	  Amount of data in bytes drained to HDD tier of extent store by the VM controller.
	*/
	ControllerOplogDrainDestHddBytes *int64 `json:"controllerOplogDrainDestHddBytes,omitempty"`
	/*
	  Amount of data in bytes drained to SSD tier of extent store by the VM controller.
	*/
	ControllerOplogDrainDestSsdBytes *int64 `json:"controllerOplogDrainDestSsdBytes,omitempty"`
	/*
	  The VM controller number of random I/O PPM.
	*/
	ControllerRandomIoPpm *int64 `json:"controllerRandomIoPpm,omitempty"`
	/*
	  The VM controller number of read I/O bandwidth in kilobytes per second.
	*/
	ControllerReadIoBandwidthKbps *int64 `json:"controllerReadIoBandwidthKbps,omitempty"`
	/*
	  The VM controller number of read I/O PPM.
	*/
	ControllerReadIoPpm *int64 `json:"controllerReadIoPpm,omitempty"`
	/*
	  The VM disk controller number of bytes for local read source: HDD estore.
	*/
	ControllerReadSourceEstoreHddLocalBytes *int64 `json:"controllerReadSourceEstoreHddLocalBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for remote read source: HDD estore.
	*/
	ControllerReadSourceEstoreHddRemoteBytes *int64 `json:"controllerReadSourceEstoreHddRemoteBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for local read source: SSD estore.
	*/
	ControllerReadSourceEstoreSsdLocalBytes *int64 `json:"controllerReadSourceEstoreSsdLocalBytes,omitempty"`
	/*
	  The VM disk controller number of bytes for remote read source: SSD estore.
	*/
	ControllerReadSourceEstoreSsdRemoteBytes *int64 `json:"controllerReadSourceEstoreSsdRemoteBytes,omitempty"`
	/*
	  The VM controller read source oplog in bytes.
	*/
	ControllerReadSourceOplogBytes *int64 `json:"controllerReadSourceOplogBytes,omitempty"`
	/*
	  The VM controller number of sequential I/O PPM.
	*/
	ControllerSeqIoPpm *int64 `json:"controllerSeqIoPpm,omitempty"`
	/*
	  The VM controller shared usage in bytes.
	*/
	ControllerSharedUsageBytes *int64 `json:"controllerSharedUsageBytes,omitempty"`
	/*
	  The VM controller snapshot usage in bytes.
	*/
	ControllerSnapshotUsageBytes *int64 `json:"controllerSnapshotUsageBytes,omitempty"`
	/*
	  The VM controller total usage on SSD tier for the VM.
	*/
	ControllerStorageTierSsdUsageBytes *int64 `json:"controllerStorageTierSsdUsageBytes,omitempty"`
	/*
	  The VM controller timespan in microseconds.
	*/
	ControllerTimespanMicros *int64 `json:"controllerTimespanMicros,omitempty"`
	/*
	  The VM controller number of total I/O size in kilobytes.
	*/
	ControllerTotalIoSizeKb *int64 `json:"controllerTotalIoSizeKb,omitempty"`
	/*
	  The VM controller number of total I/O time in microseconds.
	*/
	ControllerTotalIoTimeMicros *int64 `json:"controllerTotalIoTimeMicros,omitempty"`
	/*
	  The VM controller number of total read I/O size in kilobytes.
	*/
	ControllerTotalReadIoSizeKb *int64 `json:"controllerTotalReadIoSizeKb,omitempty"`
	/*
	  The VM controller number of total read I/O time in microseconds.
	*/
	ControllerTotalReadIoTimeMicros *int64 `json:"controllerTotalReadIoTimeMicros,omitempty"`
	/*
	  The VM controller number of total transformed usage in bytes.
	*/
	ControllerTotalTransformedUsageBytes *int64 `json:"controllerTotalTransformedUsageBytes,omitempty"`
	/*
	  The VM controller user bytes.
	*/
	ControllerUserBytes *int64 `json:"controllerUserBytes,omitempty"`
	/*
	  The VM controller write destination HDD estore in bytes.
	*/
	ControllerWriteDestEstoreHddBytes *int64 `json:"controllerWriteDestEstoreHddBytes,omitempty"`
	/*
	  The VM controller write destination SSD estore in bytes.
	*/
	ControllerWriteDestEstoreSsdBytes *int64 `json:"controllerWriteDestEstoreSsdBytes,omitempty"`
	/*
	  The VM controller write I/O bandwidth in kilobytes per second.
	*/
	ControllerWriteIoBandwidthKbps *int64 `json:"controllerWriteIoBandwidthKbps,omitempty"`
	/*
	  The VM controller percentage of write I/O in parts per million.
	*/
	ControllerWriteIoPpm *int64 `json:"controllerWriteIoPpm,omitempty"`
	/*
	  The read working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondReadMb *int64 `json:"controllerWss120SecondReadMb,omitempty"`
	/*
	  The combined working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondUnionMb *int64 `json:"controllerWss120SecondUnionMb,omitempty"`
	/*
	  The write working set size of the VM disk controller in megabytes, 2 minute measurement interval.
	*/
	ControllerWss120SecondWriteMb *int64 `json:"controllerWss120SecondWriteMb,omitempty"`
	/*
	  The read working set size of the VM disk controller in megabytes, 1 hour measurement interval.
	*/
	ControllerWss3600SecondReadMb *int64 `json:"controllerWss3600SecondReadMb,omitempty"`
	/*
	  The combined working set size of the VM disk controller in megabytes, 1-hour measurement interval.
	*/
	ControllerWss3600SecondUnionMb *int64 `json:"controllerWss3600SecondUnionMb,omitempty"`
	/*
	  The write working set size of the VM disk controller in megabytes, 1-hour measurement interval.
	*/
	ControllerWss3600SecondWriteMb *int64 `json:"controllerWss3600SecondWriteMb,omitempty"`
	/*
	  The VM disk capacity in bytes.
	*/
	DiskCapacityBytes *int64 `json:"diskCapacityBytes,omitempty"`
	/*
	  The VM disk usage in PPM.
	*/
	DiskUsagePpm *int64 `json:"diskUsagePpm,omitempty"`
	/*
	  The VM frame buffer usage in PPM.
	*/
	FrameBufferUsagePpm *int64 `json:"frameBufferUsagePpm,omitempty"`
	/*
	  The VM GPU usage in PPM.
	*/
	GpuUsagePpm *int64 `json:"gpuUsagePpm,omitempty"`
	/*
	  The VM memory usage in percentage reported by a guest.
	*/
	GuestMemoryUsagePpm *int64 `json:"guestMemoryUsagePpm,omitempty"`
	/*
	  The average I/O latency of the VM in microseconds
	*/
	HypervisorAvgIoLatencyMicros *int64 `json:"hypervisorAvgIoLatencyMicros,omitempty"`
	/*
	  Percentage of time that the VM was ready, but could not get scheduled to run.
	*/
	HypervisorCpuReadyTimePpm *int64 `json:"hypervisorCpuReadyTimePpm,omitempty"`
	/*
	  The CPU usage of the VM in parts per million.
	*/
	HypervisorCpuUsagePpm *int64 `json:"hypervisorCpuUsagePpm,omitempty"`
	/*
	  The I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorIoBandwidthKbps *int64 `json:"hypervisorIoBandwidthKbps,omitempty"`
	/*
	  The target memory balloon of the VM to be reclaimed in bytes.
	*/
	HypervisorMemoryBalloonReclaimTargetBytes *int64 `json:"hypervisorMemoryBalloonReclaimTargetBytes,omitempty"`
	/*
	  The memory balloon reclaimed by the VM in bytes.
	*/
	HypervisorMemoryBalloonReclaimedBytes *int64 `json:"hypervisorMemoryBalloonReclaimedBytes,omitempty"`
	/*
	  Consolidated guest memory usage in percentage.
	*/
	HypervisorMemoryUsagePpm *int64 `json:"hypervisorMemoryUsagePpm,omitempty"`
	/*
	  The number of I/O by the VM.
	*/
	HypervisorNumIo *int64 `json:"hypervisorNumIo,omitempty"`
	/*
	  The number of I/O operations by the VM per second.
	*/
	HypervisorNumIops *int64 `json:"hypervisorNumIops,omitempty"`
	/*
	  The number of read I/O operations by the VM.
	*/
	HypervisorNumReadIo *int64 `json:"hypervisorNumReadIo,omitempty"`
	/*
	  The number of read I/O operations by the VM per second.
	*/
	HypervisorNumReadIops *int64 `json:"hypervisorNumReadIops,omitempty"`
	/*
	  The number of packets dropped by the VM during reception.
	*/
	HypervisorNumReceivePacketsDropped *int64 `json:"hypervisorNumReceivePacketsDropped,omitempty"`
	/*
	  The number of bytes received by the VM.
	*/
	HypervisorNumReceivedBytes *int64 `json:"hypervisorNumReceivedBytes,omitempty"`
	/*
	  The number of packets dropped by the VM during transmission.
	*/
	HypervisorNumTransmitPacketsDropped *int64 `json:"hypervisorNumTransmitPacketsDropped,omitempty"`
	/*
	  The number of bytes transmitted by the VM.
	*/
	HypervisorNumTransmittedBytes *int64 `json:"hypervisorNumTransmittedBytes,omitempty"`
	/*
	  The number of write I/O by the VM.
	*/
	HypervisorNumWriteIo *int64 `json:"hypervisorNumWriteIo,omitempty"`
	/*
	  The number of write I/O operations by the VM per second.
	*/
	HypervisorNumWriteIops *int64 `json:"hypervisorNumWriteIops,omitempty"`
	/*
	  The number of read I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorReadIoBandwidthKbps *int64 `json:"hypervisorReadIoBandwidthKbps,omitempty"`
	/*
	  The swap in rate of the VM in kilobytes per second.
	*/
	HypervisorSwapInRateKbps *int64 `json:"hypervisorSwapInRateKbps,omitempty"`
	/*
	  The swap out rate of the VM in kilobytes per second.
	*/
	HypervisorSwapOutRateKbps *int64 `json:"hypervisorSwapOutRateKbps,omitempty"`
	/*
	  The timespan of the VM in microseconds.
	*/
	HypervisorTimespanMicros *int64 `json:"hypervisorTimespanMicros,omitempty"`
	/*
	  The total I/O size of the VM in kilobytes.
	*/
	HypervisorTotalIoSizeKb *int64 `json:"hypervisorTotalIoSizeKb,omitempty"`
	/*
	  The total I/O time of the VM in microseconds.
	*/
	HypervisorTotalIoTimeMicros *int64 `json:"hypervisorTotalIoTimeMicros,omitempty"`
	/*
	  The total read I/O size of the VM in kilobytes.
	*/
	HypervisorTotalReadIoSizeKb *int64 `json:"hypervisorTotalReadIoSizeKb,omitempty"`
	/*
	  Hypervisor type of the VM.
	*/
	HypervisorType *string `json:"hypervisorType,omitempty"`
	/*
	  The running time of the VM in microseconds.
	*/
	HypervisorVmRunningTimeUsecs *int64 `json:"hypervisorVmRunningTimeUsecs,omitempty"`
	/*
	  The write I/O bandwidth of the VM in kilobytes per second.
	*/
	HypervisorWriteIoBandwidthKbps *int64 `json:"hypervisorWriteIoBandwidthKbps,omitempty"`
	/*
	  The reserved memory of the VM in bytes.
	*/
	MemoryReservedBytes *int64 `json:"memoryReservedBytes,omitempty"`
	/*
	  The VM memory usage bytes.
	*/
	MemoryUsageBytes *int64 `json:"memoryUsageBytes,omitempty"`
	/*
	  The VM memory usage in PPM.
	*/
	MemoryUsagePpm *int64 `json:"memoryUsagePpm,omitempty"`
	/*
	  The VM number of vCPUs used in PPM.
	*/
	NumVcpusUsedPpm *int64 `json:"numVcpusUsedPpm,omitempty"`
	/*
	  The timestamp of a specific VM stats response data point.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func (p *VmStatsTupleProjection) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias VmStatsTupleProjection

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

func (p *VmStatsTupleProjection) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias VmStatsTupleProjection
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = VmStatsTupleProjection(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "checkScore")
	delete(allFields, "cluster")
	delete(allFields, "controllerAvgIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoLatencyMicros")
	delete(allFields, "controllerAvgReadIoSizeKb")
	delete(allFields, "controllerAvgWriteIoLatencyMicros")
	delete(allFields, "controllerAvgWriteIoSizeKb")
	delete(allFields, "controllerIoBandwidthKbps")
	delete(allFields, "controllerNumIo")
	delete(allFields, "controllerNumIops")
	delete(allFields, "controllerNumRandomIo")
	delete(allFields, "controllerNumReadIo")
	delete(allFields, "controllerNumReadIops")
	delete(allFields, "controllerNumSeqIo")
	delete(allFields, "controllerNumWriteIo")
	delete(allFields, "controllerNumWriteIops")
	delete(allFields, "controllerOplogDrainDestHddBytes")
	delete(allFields, "controllerOplogDrainDestSsdBytes")
	delete(allFields, "controllerRandomIoPpm")
	delete(allFields, "controllerReadIoBandwidthKbps")
	delete(allFields, "controllerReadIoPpm")
	delete(allFields, "controllerReadSourceEstoreHddLocalBytes")
	delete(allFields, "controllerReadSourceEstoreHddRemoteBytes")
	delete(allFields, "controllerReadSourceEstoreSsdLocalBytes")
	delete(allFields, "controllerReadSourceEstoreSsdRemoteBytes")
	delete(allFields, "controllerReadSourceOplogBytes")
	delete(allFields, "controllerSeqIoPpm")
	delete(allFields, "controllerSharedUsageBytes")
	delete(allFields, "controllerSnapshotUsageBytes")
	delete(allFields, "controllerStorageTierSsdUsageBytes")
	delete(allFields, "controllerTimespanMicros")
	delete(allFields, "controllerTotalIoSizeKb")
	delete(allFields, "controllerTotalIoTimeMicros")
	delete(allFields, "controllerTotalReadIoSizeKb")
	delete(allFields, "controllerTotalReadIoTimeMicros")
	delete(allFields, "controllerTotalTransformedUsageBytes")
	delete(allFields, "controllerUserBytes")
	delete(allFields, "controllerWriteDestEstoreHddBytes")
	delete(allFields, "controllerWriteDestEstoreSsdBytes")
	delete(allFields, "controllerWriteIoBandwidthKbps")
	delete(allFields, "controllerWriteIoPpm")
	delete(allFields, "controllerWss120SecondReadMb")
	delete(allFields, "controllerWss120SecondUnionMb")
	delete(allFields, "controllerWss120SecondWriteMb")
	delete(allFields, "controllerWss3600SecondReadMb")
	delete(allFields, "controllerWss3600SecondUnionMb")
	delete(allFields, "controllerWss3600SecondWriteMb")
	delete(allFields, "diskCapacityBytes")
	delete(allFields, "diskUsagePpm")
	delete(allFields, "frameBufferUsagePpm")
	delete(allFields, "gpuUsagePpm")
	delete(allFields, "guestMemoryUsagePpm")
	delete(allFields, "hypervisorAvgIoLatencyMicros")
	delete(allFields, "hypervisorCpuReadyTimePpm")
	delete(allFields, "hypervisorCpuUsagePpm")
	delete(allFields, "hypervisorIoBandwidthKbps")
	delete(allFields, "hypervisorMemoryBalloonReclaimTargetBytes")
	delete(allFields, "hypervisorMemoryBalloonReclaimedBytes")
	delete(allFields, "hypervisorMemoryUsagePpm")
	delete(allFields, "hypervisorNumIo")
	delete(allFields, "hypervisorNumIops")
	delete(allFields, "hypervisorNumReadIo")
	delete(allFields, "hypervisorNumReadIops")
	delete(allFields, "hypervisorNumReceivePacketsDropped")
	delete(allFields, "hypervisorNumReceivedBytes")
	delete(allFields, "hypervisorNumTransmitPacketsDropped")
	delete(allFields, "hypervisorNumTransmittedBytes")
	delete(allFields, "hypervisorNumWriteIo")
	delete(allFields, "hypervisorNumWriteIops")
	delete(allFields, "hypervisorReadIoBandwidthKbps")
	delete(allFields, "hypervisorSwapInRateKbps")
	delete(allFields, "hypervisorSwapOutRateKbps")
	delete(allFields, "hypervisorTimespanMicros")
	delete(allFields, "hypervisorTotalIoSizeKb")
	delete(allFields, "hypervisorTotalIoTimeMicros")
	delete(allFields, "hypervisorTotalReadIoSizeKb")
	delete(allFields, "hypervisorType")
	delete(allFields, "hypervisorVmRunningTimeUsecs")
	delete(allFields, "hypervisorWriteIoBandwidthKbps")
	delete(allFields, "memoryReservedBytes")
	delete(allFields, "memoryUsageBytes")
	delete(allFields, "memoryUsagePpm")
	delete(allFields, "numVcpusUsedPpm")
	delete(allFields, "timestamp")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewVmStatsTupleProjection() *VmStatsTupleProjection {
	p := new(VmStatsTupleProjection)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.ahv.stats.VmStatsTupleProjection"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetDiskStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VmDiskStats           `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetDiskStatsApiResponseData() *OneOfGetDiskStatsApiResponseData {
	p := new(OneOfGetDiskStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetDiskStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetDiskStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case VmDiskStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmDiskStats)
		}
		*p.oneOfType2001 = v.(VmDiskStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfGetDiskStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetDiskStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmDiskStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.stats.VmDiskStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmDiskStats)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetDiskStatsApiResponseData"))
}

func (p *OneOfGetDiskStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetDiskStatsApiResponseData")
}

type OneOfGetNicStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType2001 *VmNicStats            `json:"-"`
}

func NewOneOfGetNicStatsApiResponseData() *OneOfGetNicStatsApiResponseData {
	p := new(OneOfGetNicStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNicStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNicStatsApiResponseData is nil"))
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
	case VmNicStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmNicStats)
		}
		*p.oneOfType2001 = v.(VmNicStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetNicStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetNicStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType2001 := new(VmNicStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.stats.VmNicStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmNicStats)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNicStatsApiResponseData"))
}

func (p *OneOfGetNicStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetNicStatsApiResponseData")
}

type OneOfListVmStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 []VmStats              `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfListVmStatsApiResponseData() *OneOfListVmStatsApiResponseData {
	p := new(OneOfListVmStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case []VmStats:
		p.oneOfType2001 = v.([]VmStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.ahv.stats.VmStats>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.ahv.stats.VmStats>"
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

func (p *OneOfListVmStatsApiResponseData) GetValue() interface{} {
	if "List<vmm.v4.ahv.stats.VmStats>" == *p.Discriminator {
		return p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfListVmStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new([]VmStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.ahv.stats.VmStats" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.ahv.stats.VmStats>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.ahv.stats.VmStats>"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmStatsApiResponseData"))
}

func (p *OneOfListVmStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if "List<vmm.v4.ahv.stats.VmStats>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfListVmStatsApiResponseData")
}

type OneOfGetVmStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *VmStats               `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetVmStatsApiResponseData() *OneOfGetVmStatsApiResponseData {
	p := new(OneOfGetVmStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case VmStats:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(VmStats)
		}
		*p.oneOfType2001 = v.(VmStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
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

func (p *OneOfGetVmStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetVmStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(VmStats)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.ahv.stats.VmStats" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(VmStats)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmStatsApiResponseData"))
}

func (p *OneOfGetVmStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmStatsApiResponseData")
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
