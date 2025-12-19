/*
 * Generated file models/prism/v4/protectpc/protectpc_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Prism APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
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
)

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

func (p *ApiError) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiError

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

func (p *ApiError) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiError
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiError()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ErrorMessageList != nil {
		p.ErrorMessageList = known.ErrorMessageList
	}
	if known.ErrorType != nil {
		p.ErrorType = known.ErrorType
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "errorMessageList")
	delete(allFields, "errorType")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiError() *ApiError {
	p := new(ApiError)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiError"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *ApiLink) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiLink

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

func (p *ApiLink) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiLink
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiLink()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Href != nil {
		p.Href = known.Href
	}
	if known.Rel != nil {
		p.Rel = known.Rel
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "href")
	delete(allFields, "rel")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiLink() *ApiLink {
	p := new(ApiLink)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiLink"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *ApiResponseMetadata) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiResponseMetadata

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

func (p *ApiResponseMetadata) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiResponseMetadata
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiResponseMetadata()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Links != nil {
		p.Links = known.Links
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "links")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiResponseMetadata() *ApiResponseMetadata {
	p := new(ApiResponseMetadata)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiResponseMetadata"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *ApiSuccess) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ApiSuccess

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

func (p *ApiSuccess) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ApiSuccess
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewApiSuccess()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Message != nil {
		p.Message = known.Message
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewApiSuccess() *ApiSuccess {
	p := new(ApiSuccess)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ApiSuccess"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *BackupTargets) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BackupTargets

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

func (p *BackupTargets) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BackupTargets
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBackupTargets()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ClusterUuidList != nil {
		p.ClusterUuidList = known.ClusterUuidList
	}
	if known.ObjectStoreEndpointList != nil {
		p.ObjectStoreEndpointList = known.ObjectStoreEndpointList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "clusterUuidList")
	delete(allFields, "objectStoreEndpointList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBackupTargets() *BackupTargets {
	p := new(BackupTargets)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.BackupTargets"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *BackupTargetsInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias BackupTargetsInfo

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

func (p *BackupTargetsInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias BackupTargetsInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewBackupTargetsInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ObjectStoreEndpointInfoList != nil {
		p.ObjectStoreEndpointInfoList = known.ObjectStoreEndpointInfoList
	}
	if known.PeInfoList != nil {
		p.PeInfoList = known.PeInfoList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "objectStoreEndpointInfoList")
	delete(allFields, "peInfoList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewBackupTargetsInfo() *BackupTargetsInfo {
	p := new(BackupTargetsInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.BackupTargetsInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *EligibleCluster) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EligibleCluster

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

func (p *EligibleCluster) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EligibleCluster
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEligibleCluster()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BackedUpEntitiesCountSupported != nil {
		p.BackedUpEntitiesCountSupported = known.BackedUpEntitiesCountSupported
	}
	if known.ClusterName != nil {
		p.ClusterName = known.ClusterName
	}
	if known.ClusterUuid != nil {
		p.ClusterUuid = known.ClusterUuid
	}
	if known.IsHostingPe != nil {
		p.IsHostingPe = known.IsHostingPe
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backedUpEntitiesCountSupported")
	delete(allFields, "clusterName")
	delete(allFields, "clusterUuid")
	delete(allFields, "isHostingPe")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEligibleCluster() *EligibleCluster {
	p := new(EligibleCluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.EligibleCluster"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *EligibleClusterList) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias EligibleClusterList

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

func (p *EligibleClusterList) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias EligibleClusterList
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewEligibleClusterList()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EligibleClusterList != nil {
		p.EligibleClusterList = known.EligibleClusterList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "eligibleClusterList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewEligibleClusterList() *EligibleClusterList {
	p := new(EligibleClusterList)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.EligibleClusterList"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *FailedRecoveryPointDetails) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FailedRecoveryPointDetails

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

func (p *FailedRecoveryPointDetails) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FailedRecoveryPointDetails
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFailedRecoveryPointDetails()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Message != nil {
		p.Message = known.Message
	}
	if known.RecoveryPointTimestamp != nil {
		p.RecoveryPointTimestamp = known.RecoveryPointTimestamp
	}
	if known.RpoSeconds != nil {
		p.RpoSeconds = known.RpoSeconds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "message")
	delete(allFields, "recoveryPointTimestamp")
	delete(allFields, "rpoSeconds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFailedRecoveryPointDetails() *FailedRecoveryPointDetails {
	p := new(FailedRecoveryPointDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.FailedRecoveryPointDetails"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *FailedRecoveryPointsStats) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias FailedRecoveryPointsStats

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

func (p *FailedRecoveryPointsStats) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias FailedRecoveryPointsStats
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewFailedRecoveryPointsStats()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FailedRecoveryPoints != nil {
		p.FailedRecoveryPoints = known.FailedRecoveryPoints
	}
	if known.TotalFailedRecoveryPoints != nil {
		p.TotalFailedRecoveryPoints = known.TotalFailedRecoveryPoints
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "failedRecoveryPoints")
	delete(allFields, "totalFailedRecoveryPoints")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewFailedRecoveryPointsStats() *FailedRecoveryPointsStats {
	p := new(FailedRecoveryPointsStats)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.FailedRecoveryPointsStats"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *ObjectStoreEndpointInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ObjectStoreEndpointInfo

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

func (p *ObjectStoreEndpointInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ObjectStoreEndpointInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewObjectStoreEndpointInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EntityId != nil {
		p.EntityId = known.EntityId
	}
	if known.IsBackupPaused != nil {
		p.IsBackupPaused = known.IsBackupPaused
	}
	if known.LastSyncTimestamp != nil {
		p.LastSyncTimestamp = known.LastSyncTimestamp
	}
	if known.ObjectStoreEndpoint != nil {
		p.ObjectStoreEndpoint = known.ObjectStoreEndpoint
	}
	if known.PauseBackupMessage != nil {
		p.PauseBackupMessage = known.PauseBackupMessage
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "entityId")
	delete(allFields, "isBackupPaused")
	delete(allFields, "lastSyncTimestamp")
	delete(allFields, "objectStoreEndpoint")
	delete(allFields, "pauseBackupMessage")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewObjectStoreEndpointInfo() *ObjectStoreEndpointInfo {
	p := new(ObjectStoreEndpointInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ObjectStoreEndpointInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *PEInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PEInfo

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

func (p *PEInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PEInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPEInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.IsBackupPaused != nil {
		p.IsBackupPaused = known.IsBackupPaused
	}
	if known.LastSyncTimestamp != nil {
		p.LastSyncTimestamp = known.LastSyncTimestamp
	}
	if known.PauseBackupMessage != nil {
		p.PauseBackupMessage = known.PauseBackupMessage
	}
	if known.PeClusterId != nil {
		p.PeClusterId = known.PeClusterId
	}
	if known.PeName != nil {
		p.PeName = known.PeName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "isBackupPaused")
	delete(allFields, "lastSyncTimestamp")
	delete(allFields, "pauseBackupMessage")
	delete(allFields, "peClusterId")
	delete(allFields, "peName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPEInfo() *PEInfo {
	p := new(PEInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PEInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
This object consists of accessKey, secretAccessKey, certificate for a given entityId.
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
	  Certificate which will be used while validating the object store identity
	*/
	Certificate *string `json:"certificate,omitempty"`
	/*
	  SecretAccessKey for the endpoint flavor.
	*/
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
}

func (p *PcEndpointCredentials) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcEndpointCredentials

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

func (p *PcEndpointCredentials) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcEndpointCredentials
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcEndpointCredentials()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.AccessKey != nil {
		p.AccessKey = known.AccessKey
	}
	if known.Certificate != nil {
		p.Certificate = known.Certificate
	}
	if known.SecretAccessKey != nil {
		p.SecretAccessKey = known.SecretAccessKey
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "accessKey")
	delete(allFields, "certificate")
	delete(allFields, "secretAccessKey")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcEndpointCredentials() *PcEndpointCredentials {
	p := new(PcEndpointCredentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcEndpointCredentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type PcEndpointFlavour int

const (
	PCENDPOINTFLAVOUR_KS3        PcEndpointFlavour = 0
	PCENDPOINTFLAVOUR_KOBJECTS   PcEndpointFlavour = 1
	PCENDPOINTFLAVOUR_KGENERICS3 PcEndpointFlavour = 2
	PCENDPOINTFLAVOUR_UNKNOWN    PcEndpointFlavour = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *PcEndpointFlavour) name(index int) string {
	names := [...]string{
		"kS3",
		"kOBJECTS",
		"kGENERICS3",
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
		"kOBJECTS",
		"kGENERICS3",
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
		"kOBJECTS",
		"kGENERICS3",
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
	  Whether custom certificate provided or not.
	*/
	HasCustomCertificate *bool `json:"hasCustomCertificate,omitempty"`
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
	  Skip the verification of certificate during communication with object
	store endpoint
	*/
	SkipCertificateValidation *bool `json:"skipCertificateValidation,omitempty"`
	/*
	  Skip the verification of TLS certificates during communication with object
	store endpoint.
	*/
	SkipTLS *bool `json:"skipTLS,omitempty"`
}

func (p *PcObjectStoreEndpoint) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcObjectStoreEndpoint

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

func (p *PcObjectStoreEndpoint) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcObjectStoreEndpoint
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcObjectStoreEndpoint()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BackupRetentionDays != nil {
		p.BackupRetentionDays = known.BackupRetentionDays
	}
	if known.Bucket != nil {
		p.Bucket = known.Bucket
	}
	if known.EndpointAddress != nil {
		p.EndpointAddress = known.EndpointAddress
	}
	if known.EndpointCredentials != nil {
		p.EndpointCredentials = known.EndpointCredentials
	}
	if known.EndpointFlavour != nil {
		p.EndpointFlavour = known.EndpointFlavour
	}
	if known.HasCustomCertificate != nil {
		p.HasCustomCertificate = known.HasCustomCertificate
	}
	if known.IpAddressOrDomain != nil {
		p.IpAddressOrDomain = known.IpAddressOrDomain
	}
	if known.Port != nil {
		p.Port = known.Port
	}
	if known.Region != nil {
		p.Region = known.Region
	}
	if known.RpoSeconds != nil {
		p.RpoSeconds = known.RpoSeconds
	}
	if known.SkipCertificateValidation != nil {
		p.SkipCertificateValidation = known.SkipCertificateValidation
	}
	if known.SkipTLS != nil {
		p.SkipTLS = known.SkipTLS
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backupRetentionDays")
	delete(allFields, "bucket")
	delete(allFields, "endpointAddress")
	delete(allFields, "endpointCredentials")
	delete(allFields, "endpointFlavour")
	delete(allFields, "hasCustomCertificate")
	delete(allFields, "ipAddressOrDomain")
	delete(allFields, "port")
	delete(allFields, "region")
	delete(allFields, "rpoSeconds")
	delete(allFields, "skipCertificateValidation")
	delete(allFields, "skipTLS")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcObjectStoreEndpoint() *PcObjectStoreEndpoint {
	p := new(PcObjectStoreEndpoint)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcObjectStoreEndpoint"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	p.BackupRetentionDays = new(int)
	*p.BackupRetentionDays = 31
	p.Port = new(string)
	*p.Port = "443"
	p.Region = new(string)
	*p.Region = "us-east-1"
	p.SkipCertificateValidation = new(bool)
	*p.SkipCertificateValidation = false
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

func (p *PcRestoreRootTask) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcRestoreRootTask

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

func (p *PcRestoreRootTask) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcRestoreRootTask
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcRestoreRootTask()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.TaskUuid != nil {
		p.TaskUuid = known.TaskUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "taskUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcRestoreRootTask() *PcRestoreRootTask {
	p := new(PcRestoreRootTask)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcRestoreRootTask"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *PcvmRestoreFile) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcvmRestoreFile

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

func (p *PcvmRestoreFile) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcvmRestoreFile
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcvmRestoreFile()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.EncryptionVersion != nil {
		p.EncryptionVersion = known.EncryptionVersion
	}
	if known.FileContent != nil {
		p.FileContent = known.FileContent
	}
	if known.FilePath != nil {
		p.FilePath = known.FilePath
	}
	if known.IsEncrypted != nil {
		p.IsEncrypted = known.IsEncrypted
	}
	if known.KeyId != nil {
		p.KeyId = known.KeyId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "encryptionVersion")
	delete(allFields, "fileContent")
	delete(allFields, "filePath")
	delete(allFields, "isEncrypted")
	delete(allFields, "keyId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcvmRestoreFile() *PcvmRestoreFile {
	p := new(PcvmRestoreFile)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcvmRestoreFile"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
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

func (p *PcvmRestoreFiles) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias PcvmRestoreFiles

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

func (p *PcvmRestoreFiles) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias PcvmRestoreFiles
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewPcvmRestoreFiles()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.FileList != nil {
		p.FileList = known.FileList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "fileList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewPcvmRestoreFiles() *PcvmRestoreFiles {
	p := new(PcvmRestoreFiles)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.PcvmRestoreFiles"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

func (p *RecoveryStatus) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RecoveryStatus

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

func (p *RecoveryStatus) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RecoveryStatus
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRecoveryStatus()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.OverallCompletionPercentage != nil {
		p.OverallCompletionPercentage = known.OverallCompletionPercentage
	}
	if known.RecoveryState != nil {
		p.RecoveryState = known.RecoveryState
	}
	if known.RecoveryStateTitle != nil {
		p.RecoveryStateTitle = known.RecoveryStateTitle
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "overallCompletionPercentage")
	delete(allFields, "recoveryState")
	delete(allFields, "recoveryStateTitle")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRecoveryStatus() *RecoveryStatus {
	p := new(RecoveryStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RecoveryStatus"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

func (p *ReplicaInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ReplicaInfo

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

func (p *ReplicaInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ReplicaInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewReplicaInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.BackupUuid != nil {
		p.BackupUuid = known.BackupUuid
	}
	if known.ObjectStoreEndpoint != nil {
		p.ObjectStoreEndpoint = known.ObjectStoreEndpoint
	}
	if known.PeClusterId != nil {
		p.PeClusterId = known.PeClusterId
	}
	if known.PeClusterIpList != nil {
		p.PeClusterIpList = known.PeClusterIpList
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "backupUuid")
	delete(allFields, "objectStoreEndpoint")
	delete(allFields, "peClusterId")
	delete(allFields, "peClusterIpList")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewReplicaInfo() *ReplicaInfo {
	p := new(ReplicaInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.ReplicaInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
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

func (p *RpoConfig) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias RpoConfig

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

func (p *RpoConfig) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RpoConfig
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRpoConfig()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.RpoSeconds != nil {
		p.RpoSeconds = known.RpoSeconds
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "rpoSeconds")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRpoConfig() *RpoConfig {
	p := new(RpoConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.protectpc.RpoConfig"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

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
