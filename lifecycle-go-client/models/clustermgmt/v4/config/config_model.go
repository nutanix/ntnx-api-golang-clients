/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.3.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
)

/*
Cluster Fault tolerance. Set desiredClusterFaultTolerance for cluster create and update.
*/
type ClusterFaultToleranceRef int

const (
	CLUSTERFAULTTOLERANCEREF_UNKNOWN       ClusterFaultToleranceRef = 0
	CLUSTERFAULTTOLERANCEREF_REDACTED      ClusterFaultToleranceRef = 1
	CLUSTERFAULTTOLERANCEREF_CFT_0N_AND_0D ClusterFaultToleranceRef = 2
	CLUSTERFAULTTOLERANCEREF_CFT_1N_OR_1D  ClusterFaultToleranceRef = 3
	CLUSTERFAULTTOLERANCEREF_CFT_2N_OR_2D  ClusterFaultToleranceRef = 4
	CLUSTERFAULTTOLERANCEREF_CFT_1N_AND_1D ClusterFaultToleranceRef = 5
	CLUSTERFAULTTOLERANCEREF_CFT_1N_OR_2D  ClusterFaultToleranceRef = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ClusterFaultToleranceRef) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ClusterFaultToleranceRef) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ClusterFaultToleranceRef) index(name string) ClusterFaultToleranceRef {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"CFT_0N_AND_0D",
		"CFT_1N_OR_1D",
		"CFT_2N_OR_2D",
		"CFT_1N_AND_1D",
		"CFT_1N_OR_2D",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterFaultToleranceRef(idx)
		}
	}
	return CLUSTERFAULTTOLERANCEREF_UNKNOWN
}

func (e *ClusterFaultToleranceRef) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterFaultToleranceRef:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterFaultToleranceRef) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterFaultToleranceRef) Ref() *ClusterFaultToleranceRef {
	return &e
}

/*
Non-migratable VM details.
*/
type NonMigratableVmInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HostIp *import1.IPAddress `json:"hostIp,omitempty"`
	/*
	  Reason for a VM to be non-migratable.
	*/
	NonMigratableVmReason *string `json:"nonMigratableVmReason,omitempty"`
	/*
	  Name of the VM.
	*/
	VmName *string `json:"vmName,omitempty"`
	/*
	  UUID of the VM.
	*/
	VmUuid *string `json:"vmUuid,omitempty"`
}

func (p *NonMigratableVmInfo) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias NonMigratableVmInfo

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

func (p *NonMigratableVmInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias NonMigratableVmInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewNonMigratableVmInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.HostIp != nil {
		p.HostIp = known.HostIp
	}
	if known.NonMigratableVmReason != nil {
		p.NonMigratableVmReason = known.NonMigratableVmReason
	}
	if known.VmName != nil {
		p.VmName = known.VmName
	}
	if known.VmUuid != nil {
		p.VmUuid = known.VmUuid
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostIp")
	delete(allFields, "nonMigratableVmReason")
	delete(allFields, "vmName")
	delete(allFields, "vmUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewNonMigratableVmInfo() *NonMigratableVmInfo {
	p := new(NonMigratableVmInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NonMigratableVmInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r3"}
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
