/*
 * Generated file models/prism/v4/management/management_model.go.
 *
 * Product version: 4.2.1
 *
 * Part of the Nutanix Lifecycle Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.management of Nutanix Lifecycle Management APIs
*/
package management

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/common/v1/config"
)

/*
Credentials to connect to a remote cluster.
*/
type Credentials struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Authentication *import1.BasicAuth `json:"authentication"`
}

func (p *Credentials) MarshalJSON() ([]byte, error) {
	type CredentialsProxy Credentials

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*CredentialsProxy
		Authentication *import1.BasicAuth `json:"authentication,omitempty"`
	}{
		CredentialsProxy: (*CredentialsProxy)(p),
		Authentication:   p.Authentication,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *Credentials) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Credentials
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCredentials()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Authentication != nil {
		p.Authentication = known.Authentication
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "authentication")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCredentials() *Credentials {
	p := new(Credentials)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.Credentials"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r2"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Address configuration of a remote cluster. It requires the address of the remote, that is an IP or domain name along with the basic authentication credentials.
*/
type RemoteClusterSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Address *import1.IPAddressOrFQDN `json:"address"`

	Credentials *Credentials `json:"credentials"`
}

func (p *RemoteClusterSpec) MarshalJSON() ([]byte, error) {
	type RemoteClusterSpecProxy RemoteClusterSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*RemoteClusterSpecProxy
		Address     *import1.IPAddressOrFQDN `json:"address,omitempty"`
		Credentials *Credentials             `json:"credentials,omitempty"`
	}{
		RemoteClusterSpecProxy: (*RemoteClusterSpecProxy)(p),
		Address:                p.Address,
		Credentials:            p.Credentials,
	}

	known, err := json.Marshal(baseStruct)
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

func (p *RemoteClusterSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias RemoteClusterSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewRemoteClusterSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.Address != nil {
		p.Address = known.Address
	}
	if known.Credentials != nil {
		p.Credentials = known.Credentials
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "address")
	delete(allFields, "credentials")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewRemoteClusterSpec() *RemoteClusterSpec {
	p := new(RemoteClusterSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "prism.v4.management.RemoteClusterSpec"
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
