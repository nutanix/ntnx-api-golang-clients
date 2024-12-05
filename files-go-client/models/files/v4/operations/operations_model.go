/*
 * Generated file models/files/v4/operations/operations_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Files APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module files.v4.operations of Nutanix Files APIs
*/
package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import4 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/config"
	import3 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/files/v4/error"
	import2 "github.com/nutanix/ntnx-api-golang-clients/files-go-client/v4/models/prism/v4/config"
	"time"
)

/*
Configuration to access the file server active directory server.
*/
type ADServerSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Credential *import1.Credential `json:"credential"`
	/*
	  Use a specific domain controller for the join-domain operation in a multi-DC active directory setup. By default, AFS discovers a site local domain controller for join-domain operation. The preferred domain controller cannot be an IP address. It has to be FQDN (example: dc_name.dom.companyname.com)
	*/
	PreferredDomainController *string `json:"preferredDomainController,omitempty"`
}

func (p *ADServerSpec) MarshalJSON() ([]byte, error) {
	type ADServerSpecProxy ADServerSpec
	return json.Marshal(struct {
		*ADServerSpecProxy
		Credential *import1.Credential `json:"credential,omitempty"`
	}{
		ADServerSpecProxy: (*ADServerSpecProxy)(p),
		Credential:        p.Credential,
	})
}

func NewADServerSpec() *ADServerSpec {
	p := new(ADServerSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.ADServerSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/operations/$actions/ad-dns-failover Post operation
*/
type AdDnsFailoverApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAdDnsFailoverApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAdDnsFailoverApiResponse() *AdDnsFailoverApiResponse {
	p := new(AdDnsFailoverApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.AdDnsFailoverApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AdDnsFailoverApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AdDnsFailoverApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAdDnsFailoverApiResponseData()
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

type AdDnsFailoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActiveDirectory *ADServerSpec `json:"activeDirectory,omitempty"`

	Dns *import1.DnsRecordSpec `json:"dns,omitempty"`
	/*
	  The external identifier of the primary file server for the replication.
	*/
	PrimaryFileServerExtId *string `json:"primaryFileServerExtId"`
	/*
	  The external identifier of the secondary file server for the replication.
	*/
	SecondaryFileServerExtId *string `json:"secondaryFileServerExtId"`

	Type *FailoverType `json:"type,omitempty"`
}

func NewAdDnsFailoverSpec() *AdDnsFailoverSpec {
	p := new(AdDnsFailoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.AdDnsFailoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/operations/$actions/failover Post operation
*/
type FailoverApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfFailoverApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewFailoverApiResponse() *FailoverApiResponse {
	p := new(FailoverApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.FailoverApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *FailoverApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *FailoverApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfFailoverApiResponseData()
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
Failover model.
*/
type FailoverSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActiveDirectory *ADServerSpec `json:"activeDirectory,omitempty"`

	Dns *import1.DnsRecordSpec `json:"dns,omitempty"`
	/*
	  The external identifier of the primary file server for the replication.
	*/
	PrimaryFileServerExtId *string `json:"primaryFileServerExtId"`
	/*
	  The external identifier of the secondary file server for the replication.
	*/
	SecondaryFileServerExtId *string `json:"secondaryFileServerExtId"`

	Type *FailoverType `json:"type,omitempty"`
}

func (p *FailoverSpec) MarshalJSON() ([]byte, error) {
	type FailoverSpecProxy FailoverSpec
	return json.Marshal(struct {
		*FailoverSpecProxy
		PrimaryFileServerExtId   *string `json:"primaryFileServerExtId,omitempty"`
		SecondaryFileServerExtId *string `json:"secondaryFileServerExtId,omitempty"`
	}{
		FailoverSpecProxy:        (*FailoverSpecProxy)(p),
		PrimaryFileServerExtId:   p.PrimaryFileServerExtId,
		SecondaryFileServerExtId: p.SecondaryFileServerExtId,
	})
}

func NewFailoverSpec() *FailoverSpec {
	p := new(FailoverSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.FailoverSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Failover type value.
*/
type FailoverType int

const (
	FAILOVERTYPE_UNKNOWN   FailoverType = 0
	FAILOVERTYPE_REDACTED  FailoverType = 1
	FAILOVERTYPE_PLANNED   FailoverType = 2
	FAILOVERTYPE_UNPLANNED FailoverType = 3
	FAILOVERTYPE_FAILBACK  FailoverType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *FailoverType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLANNED",
		"UNPLANNED",
		"FAILBACK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e FailoverType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLANNED",
		"UNPLANNED",
		"FAILBACK",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *FailoverType) index(name string) FailoverType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PLANNED",
		"UNPLANNED",
		"FAILBACK",
	}
	for idx := range names {
		if names[idx] == name {
			return FailoverType(idx)
		}
	}
	return FAILOVERTYPE_UNKNOWN
}

func (e *FailoverType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for FailoverType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *FailoverType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e FailoverType) Ref() *FailoverType {
	return &e
}

/*
REST response for all response codes in API path /files/v4.0/operations/replication-jobs/{extId} Get operation
*/
type GetReplicationJobApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetReplicationJobApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetReplicationJobApiResponse() *GetReplicationJobApiResponse {
	p := new(GetReplicationJobApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.GetReplicationJobApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetReplicationJobApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetReplicationJobApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetReplicationJobApiResponseData()
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
Status of each job.
*/
type JobStatus int

const (
	JOBSTATUS_UNKNOWN   JobStatus = 0
	JOBSTATUS_REDACTED  JobStatus = 1
	JOBSTATUS_RUNNING   JobStatus = 2
	JOBSTATUS_QUEUED    JobStatus = 3
	JOBSTATUS_SUCCEEDED JobStatus = 4
	JOBSTATUS_CANCELLED JobStatus = 5
	JOBSTATUS_PAUSED    JobStatus = 6
	JOBSTATUS_FAILED    JobStatus = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *JobStatus) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"QUEUED",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e JobStatus) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"QUEUED",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *JobStatus) index(name string) JobStatus {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"RUNNING",
		"QUEUED",
		"SUCCEEDED",
		"CANCELLED",
		"PAUSED",
		"FAILED",
	}
	for idx := range names {
		if names[idx] == name {
			return JobStatus(idx)
		}
	}
	return JOBSTATUS_UNKNOWN
}

func (e *JobStatus) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for JobStatus:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *JobStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e JobStatus) Ref() *JobStatus {
	return &e
}

/*
REST response for all response codes in API path /files/v4.0/operations/replication-jobs Get operation
*/
type ListReplicationJobsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListReplicationJobsApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListReplicationJobsApiResponse() *ListReplicationJobsApiResponse {
	p := new(ListReplicationJobsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.ListReplicationJobsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListReplicationJobsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListReplicationJobsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListReplicationJobsApiResponseData()
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
Migration plan job details.
*/
type MigrationJob struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The average data throughput for the migration sub-plan in bytes/sec.
	*/
	AverageDataThroughputBps *float32 `json:"averageDataThroughputBps,omitempty"`
	/*
	  The average entities throughput for migration sub-plan in entities/sec.
	*/
	AverageEntitiesThroughput *int64 `json:"averageEntitiesThroughput,omitempty"`
	/*
	  The average average file scan rate for migration sub-plan in files/sec.
	*/
	AverageFilesScannedPerSecond *int64 `json:"averageFilesScannedPerSecond,omitempty"`
	/*
	  End time for the migration sub-plan.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  The iteration number for the migration sub-plan.
	*/
	IterationNumber *int64 `json:"iterationNumber,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  The number of migration sub-plan deleted directories.
	*/
	NumberOfDirectoriesDeleted *int64 `json:"numberOfDirectoriesDeleted,omitempty"`
	/*
	  The number of the transferred failed directories for the migration sub-plan.
	*/
	NumberOfDirectoriesFailed *int64 `json:"numberOfDirectoriesFailed,omitempty"`
	/*
	  The number of the transferred directories for the migration sub-plan.
	*/
	NumberOfDirectoriesTransferred *int64 `json:"numberOfDirectoriesTransferred,omitempty"`
	/*
	  The number of migration sub-plan deleted files.
	*/
	NumberOfFilesDeleted *int64 `json:"numberOfFilesDeleted,omitempty"`
	/*
	  The number of transferred failed files for the migration sub-plan.
	*/
	NumberOfFilesFailed *int64 `json:"numberOfFilesFailed,omitempty"`
	/*
	  The number of migration sub-plan files scanned.
	*/
	NumberOfFilesScanned *int64 `json:"numberOfFilesScanned,omitempty"`
	/*
	  The number of migration sub-plan files skipped.
	*/
	NumberOfFilesSkipped *int64 `json:"numberOfFilesSkipped,omitempty"`
	/*
	  The number of the transferred files for the migration sub-plan.
	*/
	NumberOfFilesTransferred *int64 `json:"numberOfFilesTransferred,omitempty"`
	/*
	  The number of migration sub-plan deleted streams.
	*/
	NumberOfStreamsDeleted *int64 `json:"numberOfStreamsDeleted,omitempty"`
	/*
	  The number of migration sub-plan failed streams.
	*/
	NumberOfStreamsFailed *int64 `json:"numberOfStreamsFailed,omitempty"`
	/*
	  The number of migration sub-plan streams transferred.
	*/
	NumberOfStreamsTransferred *int64 `json:"numberOfStreamsTransferred,omitempty"`

	OverallStatus *MigrationStatusInfo `json:"overallStatus,omitempty"`
	/*
	  Start time for the migration sub-plan.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
	/*
	  List of sub-job status. When a migration leader's job moves to running state, it creates node jobs on each node of the file server, and the status of the node job can be different than the leader job, for example: one node job can be queued, but another node job can be in a running state, but to maintain consistency, we cannot change the status of the leader job based on node jobs.
	*/
	SubJobsStatus []MigrationStatusInfo `json:"subJobsStatus,omitempty"`
	/*
	  The extId of migration sub-plan.
	*/
	SubPlanExtId *string `json:"subPlanExtId,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  The skipped bytes for the migration sub-plan.
	*/
	TotalDataSkippedBytes *int64 `json:"totalDataSkippedBytes,omitempty"`
	/*
	  The transferred bytes for the migration sub-plan.
	*/
	TotalDataTransferredBytes *int64 `json:"totalDataTransferredBytes,omitempty"`
	/*
	  The number of migration sub-plan bytes metadata transferred.
	*/
	TotalMetadataTransferredBytes *int64 `json:"totalMetadataTransferredBytes,omitempty"`
}

func NewMigrationJob() *MigrationJob {
	p := new(MigrationJob)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.MigrationJob"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.AverageDataThroughputBps = new(float32)
	*p.AverageDataThroughputBps = 0
	p.AverageEntitiesThroughput = new(int64)
	*p.AverageEntitiesThroughput = 0
	p.AverageFilesScannedPerSecond = new(int64)
	*p.AverageFilesScannedPerSecond = 0
	p.NumberOfDirectoriesDeleted = new(int64)
	*p.NumberOfDirectoriesDeleted = 0
	p.NumberOfDirectoriesFailed = new(int64)
	*p.NumberOfDirectoriesFailed = 0
	p.NumberOfDirectoriesTransferred = new(int64)
	*p.NumberOfDirectoriesTransferred = 0
	p.NumberOfFilesDeleted = new(int64)
	*p.NumberOfFilesDeleted = 0
	p.NumberOfFilesFailed = new(int64)
	*p.NumberOfFilesFailed = 0
	p.NumberOfFilesScanned = new(int64)
	*p.NumberOfFilesScanned = 0
	p.NumberOfFilesSkipped = new(int64)
	*p.NumberOfFilesSkipped = 0
	p.NumberOfFilesTransferred = new(int64)
	*p.NumberOfFilesTransferred = 0
	p.NumberOfStreamsDeleted = new(int64)
	*p.NumberOfStreamsDeleted = 0
	p.NumberOfStreamsFailed = new(int64)
	*p.NumberOfStreamsFailed = 0
	p.NumberOfStreamsTransferred = new(int64)
	*p.NumberOfStreamsTransferred = 0
	p.TotalDataSkippedBytes = new(int64)
	*p.TotalDataSkippedBytes = 0
	p.TotalDataTransferredBytes = new(int64)
	*p.TotalDataTransferredBytes = 0
	p.TotalMetadataTransferredBytes = new(int64)
	*p.TotalMetadataTransferredBytes = 0

	return p
}

/*
Status information details
*/
type MigrationStatusInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Status *import1.MigrationStatus `json:"status"`
	/*
	  Indicates the status message of the migration sub-plan. It updates after each migration step, informing the current migration sub-plan stage. This is a read-only field.
	*/
	StatusMsg *string `json:"statusMsg"`
}

func (p *MigrationStatusInfo) MarshalJSON() ([]byte, error) {
	type MigrationStatusInfoProxy MigrationStatusInfo
	return json.Marshal(struct {
		*MigrationStatusInfoProxy
		Status    *import1.MigrationStatus `json:"status,omitempty"`
		StatusMsg *string                  `json:"statusMsg,omitempty"`
	}{
		MigrationStatusInfoProxy: (*MigrationStatusInfoProxy)(p),
		Status:                   p.Status,
		StatusMsg:                p.StatusMsg,
	})
}

func NewMigrationStatusInfo() *MigrationStatusInfo {
	p := new(MigrationStatusInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.MigrationStatusInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Post failover, resume replication direction enum value.
*/
type ReplicationDirection int

const (
	REPLICATIONDIRECTION_UNKNOWN              ReplicationDirection = 0
	REPLICATIONDIRECTION_REDACTED             ReplicationDirection = 1
	REPLICATIONDIRECTION_PRIMARY_TO_SECONDARY ReplicationDirection = 2
	REPLICATIONDIRECTION_SECONDARY_TO_PRIMARY ReplicationDirection = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ReplicationDirection) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY_TO_SECONDARY",
		"SECONDARY_TO_PRIMARY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ReplicationDirection) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY_TO_SECONDARY",
		"SECONDARY_TO_PRIMARY",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ReplicationDirection) index(name string) ReplicationDirection {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PRIMARY_TO_SECONDARY",
		"SECONDARY_TO_PRIMARY",
	}
	for idx := range names {
		if names[idx] == name {
			return ReplicationDirection(idx)
		}
	}
	return REPLICATIONDIRECTION_UNKNOWN
}

func (e *ReplicationDirection) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ReplicationDirection:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ReplicationDirection) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ReplicationDirection) Ref() *ReplicationDirection {
	return &e
}

/*
Replication job details.
*/
type ReplicationJob struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Average throughput for the replication job in Bytes per second.
	*/
	AverageThroughputBps *int64 `json:"averageThroughputBps,omitempty"`
	/*
	  Actual bytes transferred for the replication job.
	*/
	BytesTransferred *int64 `json:"bytesTransferred,omitempty"`
	/*
	  End time for the replication job in ISO format.
	*/
	EndTime *time.Time `json:"endTime,omitempty"`
	/*
	  Estimated bytes for the replication job.
	*/
	EstimatedBytes *int64 `json:"estimatedBytes,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Delete propagation set/unset the replication job. If delete propagation is enabled, the deletes on source will be propagated to target
	*/
	IsDeletePropagationEnabled *bool `json:"isDeletePropagationEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import4.ApiLink `json:"links,omitempty"`
	/*
	  Number of Estimated files for the replication job.
	*/
	NumberOfEstimatedFiles *int64 `json:"numberOfEstimatedFiles,omitempty"`
	/*
	  Number of Files failed for the replication job.
	*/
	NumberOfFilesFailed *int64 `json:"numberOfFilesFailed,omitempty"`
	/*
	  Number of Actual Files transferred for the replication job.
	*/
	NumberOfFilesTransferred *int64 `json:"numberOfFilesTransferred,omitempty"`
	/*
	  External identifier of the policy.
	*/
	PolicyExtId *string `json:"policyExtId,omitempty"`
	/*
	  Progress (%) of the replication job.
	*/
	ProgressPercentage *int64 `json:"progressPercentage,omitempty"`

	ReplicationSummary *import1.ReplicationSummary `json:"replicationSummary,omitempty"`
	/*
	  The external identifier of the primary file server for the replication.
	*/
	SourceFileServerExtId *string `json:"sourceFileServerExtId,omitempty"`
	/*
	  The external identifier of the source mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
	*/
	SourceMountTargetExtId *string `json:"sourceMountTargetExtId,omitempty"`
	/*
	  User provided source path for the replication job.
	*/
	SourceMountTargetPath *string `json:"sourceMountTargetPath,omitempty"`
	/*
	  Start time for the replication job in ISO format.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`

	Status *JobStatus `json:"status,omitempty"`
	/*
	  Status Message for the replication job. If the job failed or completed with errors, this field provides the reason for the failure. If the job succeeded, this field provides the success message - Replication is complete.
	*/
	StatusMessage *string `json:"statusMessage,omitempty"`
	/*
	  The external identifier of the secondary file server for the replication.
	*/
	TargetFileServerExtId *string `json:"targetFileServerExtId,omitempty"`
	/*
	  The external identifier of the target mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
	*/
	TargetMountTargetExtId *string `json:"targetMountTargetExtId,omitempty"`
	/*
	  User provided target path for the replication job.
	*/
	TargetMountTargetPath *string `json:"targetMountTargetPath,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewReplicationJob() *ReplicationJob {
	p := new(ReplicationJob)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.ReplicationJob"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /files/v4.0/operations/$actions/resume-replication Post operation
*/
type ResumeReplicationApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfResumeReplicationApiResponseData `json:"data,omitempty"`

	Metadata *import4.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewResumeReplicationApiResponse() *ResumeReplicationApiResponse {
	p := new(ResumeReplicationApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.ResumeReplicationApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ResumeReplicationApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ResumeReplicationApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfResumeReplicationApiResponseData()
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
This contains the parameters to resume mount target level replication after performing failover.
*/
type ResumeReplicationSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ActiveDirectory *ADServerSpec `json:"activeDirectory,omitempty"`

	Dns *import1.DnsRecordSpec `json:"dns,omitempty"`
	/*
	  The external identifier of the primary file server for the replication.
	*/
	PrimaryFileServerExtId *string `json:"primaryFileServerExtId"`

	ReplicationDirection *ReplicationDirection `json:"replicationDirection,omitempty"`
	/*
	  The external identifier of the secondary file server for the replication.
	*/
	SecondaryFileServerExtId *string `json:"secondaryFileServerExtId"`

	Type *FailoverType `json:"type,omitempty"`
}

func (p *ResumeReplicationSpec) MarshalJSON() ([]byte, error) {
	type ResumeReplicationSpecProxy ResumeReplicationSpec
	return json.Marshal(struct {
		*ResumeReplicationSpecProxy
		PrimaryFileServerExtId   *string `json:"primaryFileServerExtId,omitempty"`
		SecondaryFileServerExtId *string `json:"secondaryFileServerExtId,omitempty"`
	}{
		ResumeReplicationSpecProxy: (*ResumeReplicationSpecProxy)(p),
		PrimaryFileServerExtId:     p.PrimaryFileServerExtId,
		SecondaryFileServerExtId:   p.SecondaryFileServerExtId,
	})
}

func NewResumeReplicationSpec() *ResumeReplicationSpec {
	p := new(ResumeReplicationSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "files.v4.operations.ResumeReplicationSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfAdDnsFailoverApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
}

func NewOneOfAdDnsFailoverApiResponseData() *OneOfAdDnsFailoverApiResponseData {
	p := new(OneOfAdDnsFailoverApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAdDnsFailoverApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAdDnsFailoverApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfAdDnsFailoverApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfAdDnsFailoverApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAdDnsFailoverApiResponseData"))
}

func (p *OneOfAdDnsFailoverApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfAdDnsFailoverApiResponseData")
}

type OneOfFailoverApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
}

func NewOneOfFailoverApiResponseData() *OneOfFailoverApiResponseData {
	p := new(OneOfFailoverApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfFailoverApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfFailoverApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfFailoverApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfFailoverApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfFailoverApiResponseData"))
}

func (p *OneOfFailoverApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfFailoverApiResponseData")
}

type OneOfResumeReplicationApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *import2.TaskReference `json:"-"`
}

func NewOneOfResumeReplicationApiResponseData() *OneOfResumeReplicationApiResponseData {
	p := new(OneOfResumeReplicationApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfResumeReplicationApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfResumeReplicationApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import2.TaskReference)
		}
		*p.oneOfType0 = v.(import2.TaskReference)
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

func (p *OneOfResumeReplicationApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfResumeReplicationApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import2.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfResumeReplicationApiResponseData"))
}

func (p *OneOfResumeReplicationApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfResumeReplicationApiResponseData")
}

type OneOfGetReplicationJobApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    *ReplicationJob        `json:"-"`
}

func NewOneOfGetReplicationJobApiResponseData() *OneOfGetReplicationJobApiResponseData {
	p := new(OneOfGetReplicationJobApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetReplicationJobApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetReplicationJobApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ReplicationJob:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ReplicationJob)
		}
		*p.oneOfType0 = v.(ReplicationJob)
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

func (p *OneOfGetReplicationJobApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetReplicationJobApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new(ReplicationJob)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "files.v4.operations.ReplicationJob" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ReplicationJob)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetReplicationJobApiResponseData"))
}

func (p *OneOfGetReplicationJobApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetReplicationJobApiResponseData")
}

type OneOfListReplicationJobsApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType0    []ReplicationJob       `json:"-"`
}

func NewOneOfListReplicationJobsApiResponseData() *OneOfListReplicationJobsApiResponseData {
	p := new(OneOfListReplicationJobsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListReplicationJobsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListReplicationJobsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ReplicationJob:
		p.oneOfType0 = v.([]ReplicationJob)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<files.v4.operations.ReplicationJob>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<files.v4.operations.ReplicationJob>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListReplicationJobsApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<files.v4.operations.ReplicationJob>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfListReplicationJobsApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "files.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
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
	vOneOfType0 := new([]ReplicationJob)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "files.v4.operations.ReplicationJob" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<files.v4.operations.ReplicationJob>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<files.v4.operations.ReplicationJob>"
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListReplicationJobsApiResponseData"))
}

func (p *OneOfListReplicationJobsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<files.v4.operations.ReplicationJob>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfListReplicationJobsApiResponseData")
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
