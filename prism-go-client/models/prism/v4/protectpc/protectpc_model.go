/*
 * Generated file models/prism/v4/protectpc/protectpc_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.protectpc of Nutanix Prism APIs
*/
package protectpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import2 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/error"
)

/*
REST response for all response codes in API path /prism/v4.0/protectpc/add-replicas Post operation
*/
type AddReplicasApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAddReplicasApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAddReplicasApiResponse() *AddReplicasApiResponse {
	p := new(AddReplicasApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.AddReplicasApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AddReplicasApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AddReplicasApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAddReplicasApiResponseData()
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
The error response that we want to return to the user.
*/
type ApiError struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The error message field where the error response will be put.
	*/
	ErrorMessageList []string `json:"errorMessageList,omitempty"`
	/*
	  The type of error, like INVALID_INPUT, INTERNAL_SERVER_ERROR, etc.
	*/
	ErrorType *string `json:"errorType,omitempty"`
}

func NewApiError() *ApiError {
	p := new(ApiError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
HATEOAS links for the request.  For paginated requests includes prev/next/first and last links
*/
type ApiLink struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The URL that points to the relationship.
	*/
	Href *string `json:"href,omitempty"`
	/*
	  The name of the relationship.
	*/
	Rel *string `json:"rel,omitempty"`
}

func NewApiLink() *ApiLink {
	p := new(ApiLink)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiLink"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Metadata associated with API responses.
*/
type ApiResponseMetadata struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Links []ApiLink `json:"links,omitempty"`
}

func NewApiResponseMetadata() *ApiResponseMetadata {
	p := new(ApiResponseMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiResponseMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The success response that we want to return to the user.
*/
type ApiSuccess struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The success message field where the success response message will be put.
	*/
	Message *string `json:"message,omitempty"`
}

func NewApiSuccess() *ApiSuccess {
	p := new(ApiSuccess)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiSuccess"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of pe cluster uuid or object store endpoints which needs to be
provided while making a post request, to add a PE or object store as a
backup for the PC data.
*/
type BackupTargets struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  List of PE cluster uuid.
	*/
	ClusterUuidList []string `json:"clusterUuidList,omitempty"`

	ObjectStoreEndpointList []PcObjectStoreEndpoint `json:"objectStoreEndpointList,omitempty"`
}

func NewBackupTargets() *BackupTargets {
	p := new(BackupTargets)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.BackupTargets"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model to contain information regarding PE's which are already added as replicas. It is a list of PEInfo model.
*/
type BackupTargetsInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ObjectStoreEndpointInfoList []ObjectStoreEndpointInfo `json:"objectStoreEndpointInfoList,omitempty"`

	PeInfoList []PEInfo `json:"peInfoList,omitempty"`
}

func NewBackupTargetsInfo() *BackupTargetsInfo {
	p := new(BackupTargetsInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.BackupTargetsInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Eligible cluster object for PC backup. Containing some basic properties like
cluster_uuid, cluster_name, remainingStorage, totalStorage
*/
type EligibleCluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The amount of entities which can be backed up to the PE from PC.
	*/
	BackedUpEntitiesCountSupported *int64 `json:"backedUpEntitiesCountSupported,omitempty"`
	/*
	  Name of the eligible PE(cluster).
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/*
	  Uuid of the eligible PE(cluster).
	*/
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	/*
	  Whether the candidate PE was hosting PC.
	*/
	IsHostingPe *bool `json:"isHostingPe,omitempty"`
}

func NewEligibleCluster() *EligibleCluster {
	p := new(EligibleCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.EligibleCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
An object containing a list of eligible cluster object. These are the
eligible cluster for which PC backup is possible.
*/
type EligibleClusterList struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EligibleClusterList []EligibleCluster `json:"eligibleClusterList,omitempty"`
}

func NewEligibleClusterList() *EligibleClusterList {
	p := new(EligibleClusterList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.EligibleClusterList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Failed recovery stats details i.e timestamp, rpo and failure message.
*/
type FailedRecoveryPointDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Failure reason because of which backup failed for that particular timestamp.
	*/
	Message *string `json:"message,omitempty"`
	/*
	  Timestamp at which backup failed for defined object store.
	*/
	RecoveryPointTimestamp *int64 `json:"recoveryPointTimestamp,omitempty"`
	/*
	  A RPO value in seconds to be configured.
	*/
	RpoSeconds *int `json:"rpoSeconds,omitempty"`
}

func NewFailedRecoveryPointDetails() *FailedRecoveryPointDetails {
	p := new(FailedRecoveryPointDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.FailedRecoveryPointDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Contains the total failed recovery counts and stats details list.
*/
type FailedRecoveryPointsStats struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FailedRecoveryPoints []FailedRecoveryPointDetails `json:"failedRecoveryPoints,omitempty"`
	/*
	  Count of failed recovery points in last 30 days.
	*/
	TotalFailedRecoveryPoints *int64 `json:"totalFailedRecoveryPoints,omitempty"`
}

func NewFailedRecoveryPointsStats() *FailedRecoveryPointsStats {
	p := new(FailedRecoveryPointsStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.FailedRecoveryPointsStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model to contain the information of the replica ObjectStoreEndpoints. It contains information like ObjectStore endpointAddress, endpointFlavour, and lastSyncTimestamp.
*/
type ObjectStoreEndpointInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A unique id corresponding to the object store.
	*/
	EntityId *string `json:"entityId,omitempty"`
	/*
	  Tells whether the backup is paused on the given PE or not.
	*/
	IsBackupPaused *bool `json:"isBackupPaused,omitempty"`
	/*
	  The last sync time signifies the time at which the backup was last synced to PE. This time will be updated every 30min.
	*/
	LastSyncTimestamp *int64 `json:"lastSyncTimestamp,omitempty"`

	ObjectStoreEndpoint *PcObjectStoreEndpoint `json:"objectStoreEndpoint,omitempty"`
	/*
	  Tells the reason why the backup might be paused. Will be empty if isBackupPaused field is false.
	*/
	PauseBackupMessage *string `json:"pauseBackupMessage,omitempty"`
}

func NewObjectStoreEndpointInfo() *ObjectStoreEndpointInfo {
	p := new(ObjectStoreEndpointInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ObjectStoreEndpointInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Model to contain the information of the replica PE. It contains information like PE clusterUuid, PE name, and lastSyncTimestamp.
*/
type PEInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Tells whether the backup is paused on the given PE or not.
	*/
	IsBackupPaused *bool `json:"isBackupPaused,omitempty"`
	/*
	  The last sync time signifies the time at which the backup was last synced to PE. This time will be updated every 30min.
	*/
	LastSyncTimestamp *int64 `json:"lastSyncTimestamp,omitempty"`
	/*
	  Tells the reason why the backup might be paused. Will be empty if isBackupPaused field is false.
	*/
	PauseBackupMessage *string `json:"pauseBackupMessage,omitempty"`
	/*
	  PE cluster uuid. A unique id corresponding to the cluster.
	*/
	PeClusterId *string `json:"peClusterId,omitempty"`
	/*
	  A human redable name of the PE cluster.
	*/
	PeName *string `json:"peName,omitempty"`
}

func NewPEInfo() *PEInfo {
	p := new(PEInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PEInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This object consists of accessKey, secretAccessKey for a given entityId.
*/
type PcEndpointCredentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  AccessKey for the endpoint flavor
	*/
	AccessKey *string `json:"accessKey,omitempty"`
	/*
	  SecretAccessKey for the endpoint flavor.
	*/
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}

func NewPcEndpointCredentials() *PcEndpointCredentials {
	p := new(PcEndpointCredentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcEndpointCredentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PcEndpointFlavour int

const (
	PCENDPOINTFLAVOUR_KS3     PcEndpointFlavour = 0
	PCENDPOINTFLAVOUR_UNKNOWN PcEndpointFlavour = 1
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PcEndpointFlavour) name(index int) string {
	names := [...]string{
		"kS3",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e PcEndpointFlavour) GetName() string {
	index := int(e)
	names := [...]string{
		"kS3",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *PcEndpointFlavour) index(name string) PcEndpointFlavour {
	names := [...]string{
		"kS3",
		"$UNKNOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return PcEndpointFlavour(idx)
		}
	}
	return PCENDPOINTFLAVOUR_UNKNOWN
}

func (e *PcEndpointFlavour) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for PcEndpointFlavour:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *PcEndpointFlavour) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e PcEndpointFlavour) Ref() *PcEndpointFlavour {
	return &e
}

/*
The endpoint of the object store where backup data of Prism Central is
present.
*/
type PcObjectStoreEndpoint struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Retention days configured for backup in Object Store.
	*/
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`
	/*
	  The bucket name of the object store endpoint where backup data of Prism Central
	is stored.
	*/
	Bucket *string `json:"bucket,omitempty"`
	/*
	  The endpoint address of the object store where backup data of Prism Central
	is present.
	*/
	EndpointAddress *string `json:"endpointAddress,omitempty"`

	EndpointCredentials *PcEndpointCredentials `json:"endpointCredentials,omitempty"`

	EndpointFlavour *PcEndpointFlavour `json:"endpointFlavour,omitempty"`
	/*
	  The ip address or domain of the object store endpoint where backup data of Prism Central
	is stored.
	*/
	IpAddressOrDomain *string `json:"ipAddressOrDomain,omitempty"`
	/*
	  The port of the object store endpoint where backup data of Prism Central
	is stored.
	*/
	Port *string `json:"port,omitempty"`
	/*
	  The region name of the object store endpoint where backup data of Prism Central
	is stored.
	*/
	Region *string `json:"region,omitempty"`
	/*
	  A RPO value in seconds to be configured.
	*/
	RpoSeconds *int `json:"rpoSeconds,omitempty"`
	/*
	  Skip the verification of TLS certificates during communication with object
	store endpoint.
	*/
	SkipTLS *bool `json:"skipTLS,omitempty"`
}

func NewPcObjectStoreEndpoint() *PcObjectStoreEndpoint {
	p := new(PcObjectStoreEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcObjectStoreEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	p.BackupRetentionDays = new(int)
	*p.BackupRetentionDays = 31
	p.Region = new(string)
	*p.Region = "us-east-1"
	p.SkipTLS = new(bool)
	*p.SkipTLS = false

	return p
}

/*
Model to return the root task uuid created at PC for the recovery flow. As success this uuid will be returned telling that the task has been initiated and now the PE have to wait until the tasks finishes.
*/
type PcRestoreRootTask struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Task uuid of the root PC restore task.
	*/
	TaskUuid *string `json:"taskUuid,omitempty"`
}

func NewPcRestoreRootTask() *PcRestoreRootTask {
	p := new(PcRestoreRootTask)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcRestoreRootTask"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
File object containing file path and content.
*/
type PcvmRestoreFile struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EncryptionVersion *string `json:"encryptionVersion,omitempty"`
	/*
	  Contents of the file to be restored.
	*/
	FileContent *string `json:"fileContent,omitempty"`
	/*
	  Path of the file to be restored.
	*/
	FilePath *string `json:"filePath,omitempty"`
	/*
	  Whether the file is encrypted.
	*/
	IsEncrypted *bool `json:"isEncrypted,omitempty"`

	KeyId *string `json:"keyId,omitempty"`
}

func NewPcvmRestoreFile() *PcvmRestoreFile {
	p := new(PcvmRestoreFile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcvmRestoreFile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of files to be restored in new PC.
*/
type PcvmRestoreFiles struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	FileList []PcvmRestoreFile `json:"fileList,omitempty"`
}

func NewPcvmRestoreFiles() *PcvmRestoreFiles {
	p := new(PcvmRestoreFiles)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcvmRestoreFiles"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/protectpc/recover Post operation
*/
type RecoverPcApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRecoverPcApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRecoverPcApiResponse() *RecoverPcApiResponse {
	p := new(RecoverPcApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RecoverPcApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RecoverPcApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RecoverPcApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRecoverPcApiResponseData()
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
An object containing Recovery state and overall completion percentage
of recovery task.
*/
type RecoveryStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Overall completion percentage of task.
	*/
	OverallCompletionPercentage *int `json:"overallCompletionPercentage,omitempty"`
	/*
	  Recovery state of recovery task. Possible values could be IDFDataRestore,
	WaitForProcessesToReconcile.
	*/
	RecoveryState *string `json:"recoveryState,omitempty"`
	/*
	  Recovery state title, the message that appears to the user.
	*/
	RecoveryStateTitle *string `json:"recoveryStateTitle,omitempty"`
}

func NewRecoveryStatus() *RecoveryStatus {
	p := new(RecoveryStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RecoveryStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/protectpc/replica/{backupTargetID} Delete operation
*/
type RemoveReplicaResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRemoveReplicaResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRemoveReplicaResponse() *RemoveReplicaResponse {
	p := new(RemoveReplicaResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RemoveReplicaResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RemoveReplicaResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RemoveReplicaResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRemoveReplicaResponseData()
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
Contains all the IPs of the Replica PEs and PE cluster uuid which is
required to make request on the PE3. Recovered PC will try to call all
of them sequentially if it does not work.
*/
type ReplicaInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  BackupUuid for the particular backup for which recovery is triggered.
	*/
	BackupUuid *string `json:"backupUuid,omitempty"`

	ObjectStoreEndpoint *PcObjectStoreEndpoint `json:"objectStoreEndpoint,omitempty"`
	/*
	  PE cluster uuid. A unique id corresponding to the cluster.
	*/
	PeClusterId *string `json:"peClusterId,omitempty"`

	PeClusterIpList []string `json:"peClusterIpList,omitempty"`
}

func NewReplicaInfo() *ReplicaInfo {
	p := new(ReplicaInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ReplicaInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/protectpc/restore-files Post operation
*/
type RestoreFilesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfRestoreFilesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewRestoreFilesApiResponse() *RestoreFilesApiResponse {
	p := new(RestoreFilesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RestoreFilesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *RestoreFilesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *RestoreFilesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfRestoreFilesApiResponseData()
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
Object consisting of rpo value to be configured for the given entityId.
*/
type RpoConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A RPO value in seconds to be configured.
	*/
	RpoSeconds *int `json:"rpoSeconds,omitempty"`
}

func NewRpoConfig() *RpoConfig {
	p := new(RpoConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RpoConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /prism/v4.0/protectpc/stop-services Post operation
*/
type StopServicesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfStopServicesApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewStopServicesApiResponse() *StopServicesApiResponse {
	p := new(StopServicesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.StopServicesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *StopServicesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *StopServicesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfStopServicesApiResponseData()
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
REST response for all response codes in API path /prism/v4.0/protectpc/objectstore/{entityId}/$actions/update-credentials Post operation
*/
type UpdateCredentialApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateCredentialApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateCredentialApiResponse() *UpdateCredentialApiResponse {
	p := new(UpdateCredentialApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.UpdateCredentialApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateCredentialApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateCredentialApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateCredentialApiResponseData()
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
REST response for all response codes in API path /prism/v4.0/protectpc/objectstore/{entityId}/$actions/update-rpo Post operation
*/
type UpdateRpoApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateRpoApiResponseData `json:"data,omitempty"`

	Metadata *import2.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateRpoApiResponse() *UpdateRpoApiResponse {
	p := new(UpdateRpoApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.UpdateRpoApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateRpoApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateRpoApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateRpoApiResponseData()
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

type OneOfRestoreFilesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiError              `json:"-"`
	oneOfType2    *interface{}           `json:"-"`
}

func NewOneOfRestoreFilesApiResponseData() *OneOfRestoreFilesApiResponseData {
	p := new(OneOfRestoreFilesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRestoreFilesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRestoreFilesApiResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(interface{})
		}
		*p.oneOfType2 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
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
	case ApiError:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiError)
		}
		*p.oneOfType0 = v.(ApiError)
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

func (p *OneOfRestoreFilesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType2
	}
	return nil
}

func (p *OneOfRestoreFilesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if nil == *vOneOfType2 {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(interface{})
			}
			*p.oneOfType2 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRestoreFilesApiResponseData"))
}

func (p *OneOfRestoreFilesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	return nil, errors.New("No value to marshal for OneOfRestoreFilesApiResponseData")
}

type OneOfUpdateCredentialApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiSuccess            `json:"-"`
	oneOfType1    *ApiError              `json:"-"`
}

func NewOneOfUpdateCredentialApiResponseData() *OneOfUpdateCredentialApiResponseData {
	p := new(OneOfUpdateCredentialApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateCredentialApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateCredentialApiResponseData is nil"))
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
	case ApiSuccess:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiSuccess)
		}
		*p.oneOfType0 = v.(ApiSuccess)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ApiError:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApiError)
		}
		*p.oneOfType1 = v.(ApiError)
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

func (p *OneOfUpdateCredentialApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfUpdateCredentialApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiSuccess)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiSuccess" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiSuccess)
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
	vOneOfType1 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateCredentialApiResponseData"))
}

func (p *OneOfUpdateCredentialApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateCredentialApiResponseData")
}

type OneOfRecoverPcApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *PcRestoreRootTask     `json:"-"`
	oneOfType1    *ApiError              `json:"-"`
}

func NewOneOfRecoverPcApiResponseData() *OneOfRecoverPcApiResponseData {
	p := new(OneOfRecoverPcApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRecoverPcApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRecoverPcApiResponseData is nil"))
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
	case PcRestoreRootTask:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PcRestoreRootTask)
		}
		*p.oneOfType0 = v.(PcRestoreRootTask)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ApiError:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApiError)
		}
		*p.oneOfType1 = v.(ApiError)
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

func (p *OneOfRecoverPcApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfRecoverPcApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(PcRestoreRootTask)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.PcRestoreRootTask" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PcRestoreRootTask)
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
	vOneOfType1 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRecoverPcApiResponseData"))
}

func (p *OneOfRecoverPcApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfRecoverPcApiResponseData")
}

type OneOfUpdateRpoApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiSuccess            `json:"-"`
	oneOfType1    *ApiError              `json:"-"`
}

func NewOneOfUpdateRpoApiResponseData() *OneOfUpdateRpoApiResponseData {
	p := new(OneOfUpdateRpoApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateRpoApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateRpoApiResponseData is nil"))
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
	case ApiSuccess:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiSuccess)
		}
		*p.oneOfType0 = v.(ApiSuccess)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ApiError:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApiError)
		}
		*p.oneOfType1 = v.(ApiError)
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

func (p *OneOfUpdateRpoApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfUpdateRpoApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiSuccess)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiSuccess" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiSuccess)
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
	vOneOfType1 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateRpoApiResponseData"))
}

func (p *OneOfUpdateRpoApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateRpoApiResponseData")
}

type OneOfAddReplicasApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiSuccess            `json:"-"`
	oneOfType1    *ApiError              `json:"-"`
}

func NewOneOfAddReplicasApiResponseData() *OneOfAddReplicasApiResponseData {
	p := new(OneOfAddReplicasApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAddReplicasApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAddReplicasApiResponseData is nil"))
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
	case ApiSuccess:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiSuccess)
		}
		*p.oneOfType0 = v.(ApiSuccess)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ApiError:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApiError)
		}
		*p.oneOfType1 = v.(ApiError)
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

func (p *OneOfAddReplicasApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfAddReplicasApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiSuccess)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiSuccess" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiSuccess)
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
	vOneOfType1 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAddReplicasApiResponseData"))
}

func (p *OneOfAddReplicasApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfAddReplicasApiResponseData")
}

type OneOfStopServicesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiSuccess            `json:"-"`
	oneOfType1    *ApiError              `json:"-"`
}

func NewOneOfStopServicesApiResponseData() *OneOfStopServicesApiResponseData {
	p := new(OneOfStopServicesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfStopServicesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfStopServicesApiResponseData is nil"))
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
	case ApiSuccess:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiSuccess)
		}
		*p.oneOfType0 = v.(ApiSuccess)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case ApiError:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(ApiError)
		}
		*p.oneOfType1 = v.(ApiError)
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

func (p *OneOfStopServicesApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfStopServicesApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiSuccess)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiSuccess" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiSuccess)
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
	vOneOfType1 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfStopServicesApiResponseData"))
}

func (p *OneOfStopServicesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfStopServicesApiResponseData")
}

type OneOfRemoveReplicaResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import1.ErrorResponse `json:"-"`
	oneOfType0    *ApiError              `json:"-"`
	oneOfType2    *interface{}           `json:"-"`
}

func NewOneOfRemoveReplicaResponseData() *OneOfRemoveReplicaResponseData {
	p := new(OneOfRemoveReplicaResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfRemoveReplicaResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfRemoveReplicaResponseData is nil"))
	}
	if nil == v {
		if nil == p.oneOfType2 {
			p.oneOfType2 = new(interface{})
		}
		*p.oneOfType2 = nil
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "EMPTY"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "EMPTY"
		return nil
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
	case ApiError:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(ApiError)
		}
		*p.oneOfType0 = v.(ApiError)
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

func (p *OneOfRemoveReplicaResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if "EMPTY" == *p.Discriminator {
		return *p.oneOfType2
	}
	return nil
}

func (p *OneOfRemoveReplicaResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2 := new(interface{})
	if err := json.Unmarshal(b, vOneOfType2); err == nil {
		if nil == *vOneOfType2 {
			if nil == p.oneOfType2 {
				p.oneOfType2 = new(interface{})
			}
			*p.oneOfType2 = nil
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "EMPTY"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "EMPTY"
			return nil
		}
	}
	vOneOfType400 := new(import1.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "prism.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
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
	vOneOfType0 := new(ApiError)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.protectpc.ApiError" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(ApiError)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfRemoveReplicaResponseData"))
}

func (p *OneOfRemoveReplicaResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if "EMPTY" == *p.Discriminator {
		return json.Marshal(p.oneOfType2)
	}
	return nil, errors.New("No value to marshal for OneOfRemoveReplicaResponseData")
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
