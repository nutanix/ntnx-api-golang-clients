/*
 * Generated file models/files/v4/stats/stats_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Files APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module files.v4.stats of Nutanix Files APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/stats"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
)

/*
REST response for all response codes in API path /files/v4.0/stats/file-servers/{fileServerExtId}/anti-virus-servers/{extId} Get operation
*/
type AntivirusServerStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAntivirusServerStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAntivirusServerStatsApiResponse() *AntivirusServerStatsApiResponse {
	p := new(AntivirusServerStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.AntivirusServerStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AntivirusServerStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AntivirusServerStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAntivirusServerStatsApiResponseData()
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
A model that represents antivirus statistics.
*/
type AntivirusStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The number of files cleaned by the antivirus.
	*/
	CleanedFileCount []import3.TimeIntValuePair `json:"cleanedFileCount,omitempty"`
	/*
	  The number of times the antivirus is disconnected.
	*/
	DisconnectCount []import3.TimeIntValuePair `json:"disconnectCount,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The antivirus average latency in milliseconds.
	*/
	LatencyMs []import3.TimeIntValuePair `json:"latencyMs,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The number of files quarantined by the antivirus.
	*/
	QuarantinedFileCount []import3.TimeIntValuePair `json:"quarantinedFileCount,omitempty"`
	/*
	  The number of files scanned by the antivirus.
	*/
	ScannedFileCount []import3.TimeIntValuePair `json:"scannedFileCount,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The number of threats found by the antivirus.
	*/
	ThreatCount []import3.TimeIntValuePair `json:"threatCount,omitempty"`
	/*
	  The antivirus throughput in bytes per second.
	*/
	ThroughputBps []import3.TimeIntValuePair `json:"throughputBps,omitempty"`
}

func NewAntivirusStats() *AntivirusStats {
	p := new(AntivirusStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.AntivirusStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File server statistics model
*/
type FileServerStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The number of files cleaned by the antivirus.
	*/
	AvCleanedFileCount []import3.TimeIntValuePair `json:"avCleanedFileCount,omitempty"`
	/*
	  The antivirus average latency in milliseconds.
	*/
	AvLatencyMs []import3.TimeIntValuePair `json:"avLatencyMs,omitempty"`
	/*
	  The number of files quarantined by the antivirus.
	*/
	AvQuarantinedFileCount []import3.TimeIntValuePair `json:"avQuarantinedFileCount,omitempty"`
	/*
	  The number of files scanned by the antivirus.
	*/
	AvScannedFileCount []import3.TimeIntValuePair `json:"avScannedFileCount,omitempty"`
	/*
	  The number of threats found by the antivirus.
	*/
	AvThreatCount []import3.TimeIntValuePair `json:"avThreatCount,omitempty"`
	/*
	  The antivirus throughput in bytes per second.
	*/
	AvThroughputBps []import3.TimeIntValuePair `json:"avThroughputBps,omitempty"`
	/*
	  The average Input/Output rate metric is measured in IOPS.
	*/
	AverageIops []import3.TimeIntValuePair `json:"averageIops,omitempty"`
	/*
	  The average access latency metric in microseconds.
	*/
	AverageLatencyUs []import3.TimeIntValuePair `json:"averageLatencyUs,omitempty"`
	/*
	  The average throughput metric in bytes per second.
	*/
	AverageThroughputBps []import3.TimeIntValuePair `json:"averageThroughputBps,omitempty"`
	/*
	  Disk space in bytes used by the file server or mount target dataset.
	*/
	DatasetSpaceUsedBytes []import3.TimeIntValuePair `json:"datasetSpaceUsedBytes,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The length of an antivirus scan queue.
	*/
	IcapDaemonQueueLength []import3.TimeIntValuePair `json:"icapDaemonQueueLength,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The metadata Input/Output rate metric is measured in IOPS.
	*/
	MetadataIops []import3.TimeIntValuePair `json:"metadataIops,omitempty"`
	/*
	  The average metadata latency in microseconds.
	*/
	MetadataLatencyUs []import3.TimeIntValuePair `json:"metadataLatencyUs,omitempty"`
	/*
	  The number of active connections metric.
	*/
	NumberOfConnections []import3.TimeIntValuePair `json:"numberOfConnections,omitempty"`
	/*
	  The number of files metric.
	*/
	NumberOfFiles []import3.TimeIntValuePair `json:"numberOfFiles,omitempty"`
	/*
	  The file server read Input/Output rate metric is measured in IOPS.
	*/
	ReadIops []import3.TimeIntValuePair `json:"readIops,omitempty"`
	/*
	  The average read latency metric in microseconds.
	*/
	ReadLatencyUs []import3.TimeIntValuePair `json:"readLatencyUs,omitempty"`
	/*
	  The read throughput metric in bytes per second.
	*/
	ReadThroughputBps []import3.TimeIntValuePair `json:"readThroughputBps,omitempty"`
	/*
	  Disk space in bytes used by the snapshots.
	*/
	SnapshotUsedBytes []import3.TimeIntValuePair `json:"snapshotUsedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Total tiered data in bytes to the object store.
	*/
	TotalTieredBytes []import3.TimeIntValuePair `json:"totalTieredBytes,omitempty"`
	/*
	  The write Input/Output rate metric is measured in IOPS.
	*/
	WriteIops []import3.TimeIntValuePair `json:"writeIops,omitempty"`
	/*
	  The average write latency metric in microseconds.
	*/
	WriteLatencyUs []import3.TimeIntValuePair `json:"writeLatencyUs,omitempty"`
	/*
	  The write throughput metric in bytes per second.
	*/
	WriteThroughputBps []import3.TimeIntValuePair `json:"writeThroughputBps,omitempty"`
}

func NewFileServerStats() *FileServerStats {
	p := new(FileServerStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.FileServerStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/stats/file-servers/{extId} Get operation
*/
type FileServerStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfFileServerStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFileServerStatsApiResponse() *FileServerStatsApiResponse {
	p := new(FileServerStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.FileServerStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *FileServerStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *FileServerStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfFileServerStatsApiResponseData()
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
A model that represents mount target statistics.
*/
type MountTargetStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The average Input/Output rate metric is measured in IOPS.
	*/
	AverageIops []import3.TimeIntValuePair `json:"averageIops,omitempty"`
	/*
	  The average access latency metric in microseconds.
	*/
	AverageLatencyUs []import3.TimeIntValuePair `json:"averageLatencyUs,omitempty"`
	/*
	  The average throughput metric in bytes per second.
	*/
	AverageThroughputBps []import3.TimeIntValuePair `json:"averageThroughputBps,omitempty"`
	/*
	  Disk space in bytes used by the file server or mount target dataset.
	*/
	DatasetSpaceUsedBytes []import3.TimeIntValuePair `json:"datasetSpaceUsedBytes,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The metadata Input/Output rate metric is measured in IOPS.
	*/
	MetadataIops []import3.TimeIntValuePair `json:"metadataIops,omitempty"`
	/*
	  The average metadata latency in microseconds.
	*/
	MetadataLatencyUs []import3.TimeIntValuePair `json:"metadataLatencyUs,omitempty"`
	/*
	  The number of active connections metric.
	*/
	NumberOfConnections []import3.TimeIntValuePair `json:"numberOfConnections,omitempty"`
	/*
	  The number of files metric.
	*/
	NumberOfFiles []import3.TimeIntValuePair `json:"numberOfFiles,omitempty"`
	/*
	  The file server read Input/Output rate metric is measured in IOPS.
	*/
	ReadIops []import3.TimeIntValuePair `json:"readIops,omitempty"`
	/*
	  The average read latency metric in microseconds.
	*/
	ReadLatencyUs []import3.TimeIntValuePair `json:"readLatencyUs,omitempty"`
	/*
	  The read throughput metric in bytes per second.
	*/
	ReadThroughputBps []import3.TimeIntValuePair `json:"readThroughputBps,omitempty"`
	/*
	  Disk space in bytes used by the snapshots.
	*/
	SnapshotUsedBytes []import3.TimeIntValuePair `json:"snapshotUsedBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The average tiering access latency metric for put operation in milliseconds.
	*/
	TieringLatencyMs []import3.TimeIntValuePair `json:"tieringLatencyMs,omitempty"`
	/*
	  Disk space in bytes used by the mount target.
	*/
	UsedBytes []import3.TimeIntValuePair `json:"usedBytes,omitempty"`
	/*
	  The write Input/Output rate metric is measured in IOPS.
	*/
	WriteIops []import3.TimeIntValuePair `json:"writeIops,omitempty"`
	/*
	  The average write latency metric in microseconds.
	*/
	WriteLatencyUs []import3.TimeIntValuePair `json:"writeLatencyUs,omitempty"`
	/*
	  The write throughput metric in bytes per second.
	*/
	WriteThroughputBps []import3.TimeIntValuePair `json:"writeThroughputBps,omitempty"`
}

func NewMountTargetStats() *MountTargetStats {
	p := new(MountTargetStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.MountTargetStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/stats/file-servers/{fileServerExtId}/mount-targets/{extId} Get operation
*/
type MountTargetStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfMountTargetStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewMountTargetStatsApiResponse() *MountTargetStatsApiResponse {
	p := new(MountTargetStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.stats.MountTargetStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *MountTargetStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *MountTargetStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfMountTargetStatsApiResponseData()
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

type OneOfAntivirusServerStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *AntivirusStats        `json:"-"`
}

func NewOneOfAntivirusServerStatsApiResponseData() *OneOfAntivirusServerStatsApiResponseData {
	p := new(OneOfAntivirusServerStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAntivirusServerStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAntivirusServerStatsApiResponseData is nil"))
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
	case AntivirusStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(AntivirusStats)
		}
		*p.oneOfType0 = v.(AntivirusStats)
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

func (p *OneOfAntivirusServerStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAntivirusServerStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(AntivirusStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.stats.AntivirusStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(AntivirusStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAntivirusServerStatsApiResponseData"))
}

func (p *OneOfAntivirusServerStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAntivirusServerStatsApiResponseData")
}

type OneOfFileServerStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *FileServerStats       `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfFileServerStatsApiResponseData() *OneOfFileServerStatsApiResponseData {
	p := new(OneOfFileServerStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFileServerStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFileServerStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case FileServerStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(FileServerStats)
		}
		*p.oneOfType0 = v.(FileServerStats)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
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

func (p *OneOfFileServerStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfFileServerStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(FileServerStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.stats.FileServerStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(FileServerStats)
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
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFileServerStatsApiResponseData"))
}

func (p *OneOfFileServerStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfFileServerStatsApiResponseData")
}

type OneOfMountTargetStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *MountTargetStats      `json:"-"`
}

func NewOneOfMountTargetStatsApiResponseData() *OneOfMountTargetStatsApiResponseData {
	p := new(OneOfMountTargetStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfMountTargetStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfMountTargetStatsApiResponseData is nil"))
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
	case MountTargetStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(MountTargetStats)
		}
		*p.oneOfType0 = v.(MountTargetStats)
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

func (p *OneOfMountTargetStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfMountTargetStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(MountTargetStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.stats.MountTargetStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(MountTargetStats)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfMountTargetStatsApiResponseData"))
}

func (p *OneOfMountTargetStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfMountTargetStatsApiResponseData")
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
