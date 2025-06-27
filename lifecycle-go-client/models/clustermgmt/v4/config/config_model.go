/*
 * Generated file models/clustermgmt/v4/config/config_model.go.
 *
 * Product version: 4.1.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module clustermgmt.v4.config of Nutanix Lifecycle Management APIs
*/
package config

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
)

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
	*p = NonMigratableVmInfo(*known)

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "hostIp")
	delete(allFields, "nonMigratableVmReason")
	delete(allFields, "vmName")
	delete(allFields, "vmUuid")

	// Step 5: Assign remaining fields to UnknownFields_
	p.UnknownFields_ = allFields

	return nil
}

func NewNonMigratableVmInfo() *NonMigratableVmInfo {
	p := new(NonMigratableVmInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "clustermgmt.v4.config.NonMigratableVmInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1"}
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
