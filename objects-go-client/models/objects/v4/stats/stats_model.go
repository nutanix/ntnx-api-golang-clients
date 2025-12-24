/*
 * Generated file models/objects/v4/stats/stats_model.go.
 *
 * Product version: 4.0.2
 *
 * Part of the Nutanix Objects Storage Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module objects.v4.stats of Nutanix Objects Storage Management APIs
*/
package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/common/v1/stats"
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/error"
	"time"
)

/*
REST response for all response codes in API path /objects/v4.0/stats/object-stores/{extId} Get operation
*/
type GetObjectstoreStatsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetObjectstoreStatsApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetObjectstoreStatsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetObjectstoreStatsApiResponse

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

func (p *GetObjectstoreStatsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetObjectstoreStatsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetObjectstoreStatsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetObjectstoreStatsApiResponse() *GetObjectstoreStatsApiResponse {
	p := new(GetObjectstoreStatsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.stats.GetObjectstoreStatsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetObjectstoreStatsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetObjectstoreStatsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetObjectstoreStatsApiResponseData()
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

type ObjectstoreStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Number of Buckets in the Object store. This stat is updated every 10 minutes.
	*/
	BucketCount []import3.TimeIntValuePair `json:"bucketCount,omitempty"`
	/*
	  Number of HTTP DELETE requests per second on the Objects and Buckets of the Object store.
	*/
	DeleteRequestsPerSecond []TimeFloatValuePair `json:"deleteRequestsPerSecond,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Number of GetBucket operations per second.
	*/
	GetBucketOperationsPerSecond []TimeFloatValuePair `json:"getBucketOperationsPerSecond,omitempty"`
	/*
	  Time to the first byte in milliseconds for Get Object requests.
	*/
	GetObjectTtfbMsecs []TimeFloatValuePair `json:"getObjectTtfbMsecs,omitempty"`
	/*
	  Throughput of HTTP GET requests in bytes per second.
	*/
	GetRequestThroughputBytesPerSecond []TimeFloatValuePair `json:"getRequestThroughputBytesPerSecond,omitempty"`
	/*
	  Number of HTTP GET requests per second on the Objects and Buckets of the Object store.
	*/
	GetRequestsPerSecond []TimeFloatValuePair `json:"getRequestsPerSecond,omitempty"`
	/*
	  Number of HTTP HEAD requests per second on the Objects and Buckets of the Object store.
	*/
	HeadRequestsPerSecond []TimeFloatValuePair `json:"headRequestsPerSecond,omitempty"`
	/*
	  Number of bytes written per second to the Object store from operations on the Objects and Buckets.
	*/
	InboundBytesPerSecond []TimeFloatValuePair `json:"inboundBytesPerSecond,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import2.ApiLink `json:"links,omitempty"`
	/*
	  The number of ListMultipartUploads operations per second.
	*/
	ListMultipartUploadsOperationsPerSecond []TimeFloatValuePair `json:"listMultipartUploadsOperationsPerSecond,omitempty"`
	/*
	  The number of CreateMultipartUpload operations per second.
	*/
	MultipartUploadStartOperationsPerSecond []TimeFloatValuePair `json:"multipartUploadStartOperationsPerSecond,omitempty"`
	/*
	  Number of Network File System (NFS) read requests per second on the objects of the Object store.
	*/
	NfsReadRequestsPerSecond []TimeFloatValuePair `json:"nfsReadRequestsPerSecond,omitempty"`
	/*
	  Throughput of Network File System (NFS) read operations in bytes per second on the objects of the Object store.
	*/
	NfsReadThroughputBytesPerSecond []TimeFloatValuePair `json:"nfsReadThroughputBytesPerSecond,omitempty"`
	/*
	  Number of Network File System (NFS) write requests per second on the objects of the Object store.
	*/
	NfsWriteRequestsPerSecond []TimeFloatValuePair `json:"nfsWriteRequestsPerSecond,omitempty"`
	/*
	  Throughput of Network File System (NFS) write operations in bytes per second on the objects of the Object store.
	*/
	NfsWriteThroughputBytesPerSecond []TimeFloatValuePair `json:"nfsWriteThroughputBytesPerSecond,omitempty"`
	/*
	  Number of Objects in the Object store. This stat is updated every 10 minutes. It includes the count of the incomplete multipart uploads.
	*/
	ObjectCount []import3.TimeIntValuePair `json:"objectCount,omitempty"`
	/*
	  Number of operations per second on the objects of the Object store.
	*/
	ObjectOperationsPerSecond []TimeFloatValuePair `json:"objectOperationsPerSecond,omitempty"`
	/*
	  Number of bytes read per second from the Object store from operations on the Objects and Buckets.
	*/
	OutboundBytesPerSecond []TimeFloatValuePair `json:"outboundBytesPerSecond,omitempty"`
	/*
	  Number of HTTP POST requests per second on the Objects and Buckets of the Object store.
	*/
	PostRequestsPerSecond []TimeFloatValuePair `json:"postRequestsPerSecond,omitempty"`
	/*
	  Throughput of HTTP PUT requests in bytes per second.
	*/
	PutRequestThroughputBytesPerSecond []TimeFloatValuePair `json:"putRequestThroughputBytesPerSecond,omitempty"`
	/*
	  Number of HTTP PUT requests per second on the Objects and Buckets of the Object store.
	*/
	PutRequestsPerSecond []TimeFloatValuePair `json:"putRequestsPerSecond,omitempty"`
	/*
	  The number of SelectObjectContent operations per second.
	*/
	SelectObjectContentOperationsPerSecond []TimeFloatValuePair `json:"selectObjectContentOperationsPerSecond,omitempty"`
	/*
	  Total storage in bytes used by the Objects and Buckets. This stat is updated every 10 minutes. It includes the storage used by the incomplete multipart uploads.
	*/
	StorageUsageBytes []import3.TimeIntValuePair `json:"storageUsageBytes,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func (p *ObjectstoreStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ObjectstoreStats

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

func (p *ObjectstoreStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectstoreStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectstoreStats()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BucketCount != nil {
		p.BucketCount = known.BucketCount
	}
	if known.DeleteRequestsPerSecond != nil {
		p.DeleteRequestsPerSecond = known.DeleteRequestsPerSecond
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.GetBucketOperationsPerSecond != nil {
		p.GetBucketOperationsPerSecond = known.GetBucketOperationsPerSecond
	}
	if known.GetObjectTtfbMsecs != nil {
		p.GetObjectTtfbMsecs = known.GetObjectTtfbMsecs
	}
	if known.GetRequestThroughputBytesPerSecond != nil {
		p.GetRequestThroughputBytesPerSecond = known.GetRequestThroughputBytesPerSecond
	}
	if known.GetRequestsPerSecond != nil {
		p.GetRequestsPerSecond = known.GetRequestsPerSecond
	}
	if known.HeadRequestsPerSecond != nil {
		p.HeadRequestsPerSecond = known.HeadRequestsPerSecond
	}
	if known.InboundBytesPerSecond != nil {
		p.InboundBytesPerSecond = known.InboundBytesPerSecond
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.ListMultipartUploadsOperationsPerSecond != nil {
		p.ListMultipartUploadsOperationsPerSecond = known.ListMultipartUploadsOperationsPerSecond
	}
	if known.MultipartUploadStartOperationsPerSecond != nil {
		p.MultipartUploadStartOperationsPerSecond = known.MultipartUploadStartOperationsPerSecond
	}
	if known.NfsReadRequestsPerSecond != nil {
		p.NfsReadRequestsPerSecond = known.NfsReadRequestsPerSecond
	}
	if known.NfsReadThroughputBytesPerSecond != nil {
		p.NfsReadThroughputBytesPerSecond = known.NfsReadThroughputBytesPerSecond
	}
	if known.NfsWriteRequestsPerSecond != nil {
		p.NfsWriteRequestsPerSecond = known.NfsWriteRequestsPerSecond
	}
	if known.NfsWriteThroughputBytesPerSecond != nil {
		p.NfsWriteThroughputBytesPerSecond = known.NfsWriteThroughputBytesPerSecond
	}
	if known.ObjectCount != nil {
		p.ObjectCount = known.ObjectCount
	}
	if known.ObjectOperationsPerSecond != nil {
		p.ObjectOperationsPerSecond = known.ObjectOperationsPerSecond
	}
	if known.OutboundBytesPerSecond != nil {
		p.OutboundBytesPerSecond = known.OutboundBytesPerSecond
	}
	if known.PostRequestsPerSecond != nil {
		p.PostRequestsPerSecond = known.PostRequestsPerSecond
	}
	if known.PutRequestThroughputBytesPerSecond != nil {
		p.PutRequestThroughputBytesPerSecond = known.PutRequestThroughputBytesPerSecond
	}
	if known.PutRequestsPerSecond != nil {
		p.PutRequestsPerSecond = known.PutRequestsPerSecond
	}
	if known.SelectObjectContentOperationsPerSecond != nil {
		p.SelectObjectContentOperationsPerSecond = known.SelectObjectContentOperationsPerSecond
	}
	if known.StorageUsageBytes != nil {
		p.StorageUsageBytes = known.StorageUsageBytes
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "bucketCount")
	delete(allFields, "deleteRequestsPerSecond")
	delete(allFields, "extId")
	delete(allFields, "getBucketOperationsPerSecond")
	delete(allFields, "getObjectTtfbMsecs")
	delete(allFields, "getRequestThroughputBytesPerSecond")
	delete(allFields, "getRequestsPerSecond")
	delete(allFields, "headRequestsPerSecond")
	delete(allFields, "inboundBytesPerSecond")
	delete(allFields, "links")
	delete(allFields, "listMultipartUploadsOperationsPerSecond")
	delete(allFields, "multipartUploadStartOperationsPerSecond")
	delete(allFields, "nfsReadRequestsPerSecond")
	delete(allFields, "nfsReadThroughputBytesPerSecond")
	delete(allFields, "nfsWriteRequestsPerSecond")
	delete(allFields, "nfsWriteThroughputBytesPerSecond")
	delete(allFields, "objectCount")
	delete(allFields, "objectOperationsPerSecond")
	delete(allFields, "outboundBytesPerSecond")
	delete(allFields, "postRequestsPerSecond")
	delete(allFields, "putRequestThroughputBytesPerSecond")
	delete(allFields, "putRequestsPerSecond")
	delete(allFields, "selectObjectContentOperationsPerSecond")
	delete(allFields, "storageUsageBytes")
	delete(allFields, "tenantId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectstoreStats() *ObjectstoreStats {
	p := new(ObjectstoreStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.stats.ObjectstoreStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
A time value pair representing a stat associated with a given entity at a given point of date and time represented in extended ISO-8601 format.
*/
type TimeFloatValuePair struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The date and time at which the stat was recorded.The value should be in extended ISO-8601 format. For example, start time of 2022-04-23T01:23:45.678+09:00 would consider all stats starting at 1:23:45.678 on the 23rd of April 2022. Details around ISO-8601 format can be found at https://www.iso.org/standard/70907.html
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/*
	  Value of the stat at the recorded date and time in extended ISO-8601 format.
	*/
	Value *float32 `json:"value,omitempty"`
}

func (p *TimeFloatValuePair) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias TimeFloatValuePair

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

func (p *TimeFloatValuePair) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias TimeFloatValuePair
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTimeFloatValuePair()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Timestamp != nil {
		p.Timestamp = known.Timestamp
	}
	if known.Value != nil {
		p.Value = known.Value
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "timestamp")
	delete(allFields, "value")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTimeFloatValuePair() *TimeFloatValuePair {
	p := new(TimeFloatValuePair)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "objects.v4.stats.TimeFloatValuePair"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetObjectstoreStatsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *ObjectstoreStats      `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
}

func NewOneOfGetObjectstoreStatsApiResponseData() *OneOfGetObjectstoreStatsApiResponseData {
	p := new(OneOfGetObjectstoreStatsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetObjectstoreStatsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetObjectstoreStatsApiResponseData is nil"))
	}
	switch v.(type) {
	case ObjectstoreStats:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ObjectstoreStats)
		}
		*p.oneOfType0 = v.(ObjectstoreStats)
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

func (p *OneOfGetObjectstoreStatsApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetObjectstoreStatsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(ObjectstoreStats)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "objects.v4.stats.ObjectstoreStats" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ObjectstoreStats)
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
		if "objects.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetObjectstoreStatsApiResponseData"))
}

func (p *OneOfGetObjectstoreStatsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetObjectstoreStatsApiResponseData")
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
