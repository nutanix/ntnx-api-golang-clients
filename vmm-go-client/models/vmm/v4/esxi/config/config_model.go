/*
 * Generated file models/vmm/v4/esxi/config/config_model.go.
 *
 * Product version: 4.0.3-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.esxi.config of Nutanix Vmm Versioned APIs
*/
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import3 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/*
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/$actions/assign-owner Post operation
*/
type AssignVmOwnerResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssignVmOwnerResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssignVmOwnerResponse() *AssignVmOwnerResponse {
	p := new(AssignVmOwnerResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.AssignVmOwnerResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.AssignVmOwnerResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssignVmOwnerResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssignVmOwnerResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssignVmOwnerResponseData()
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
List of categories to be associated to the VM.
*/
type AssociateVmCategoriesParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Categories []CategoryReference `json:"categories,omitempty"`
}

func NewAssociateVmCategoriesParams() *AssociateVmCategoriesParams {
	p := new(AssociateVmCategoriesParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.AssociateVmCategoriesParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.AssociateVmCategoriesParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/$actions/associate-categories Post operation
*/
type AssociateVmCategoriesResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfAssociateVmCategoriesResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewAssociateVmCategoriesResponse() *AssociateVmCategoriesResponse {
	p := new(AssociateVmCategoriesResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.AssociateVmCategoriesResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.AssociateVmCategoriesResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *AssociateVmCategoriesResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *AssociateVmCategoriesResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfAssociateVmCategoriesResponseData()
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
Reference to a category.
*/
type CategoryReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier of an instance. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewCategoryReference() *CategoryReference {
	p := new(CategoryReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.CategoryReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.CategoryReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Reference to a cluster.
*/
type ClusterReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier of an instance. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewClusterReference() *ClusterReference {
	p := new(ClusterReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.ClusterReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.ClusterReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
List of categories to be disassociated from the VM.
*/
type DisassociateVmCategoriesParams struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Categories []CategoryReference `json:"categories,omitempty"`
}

func NewDisassociateVmCategoriesParams() *DisassociateVmCategoriesParams {
	p := new(DisassociateVmCategoriesParams)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.DisassociateVmCategoriesParams"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.DisassociateVmCategoriesParams"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/$actions/disassociate-categories Post operation
*/
type DisassociateVmCategoriesResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDisassociateVmCategoriesResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDisassociateVmCategoriesResponse() *DisassociateVmCategoriesResponse {
	p := new(DisassociateVmCategoriesResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.DisassociateVmCategoriesResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.DisassociateVmCategoriesResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DisassociateVmCategoriesResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DisassociateVmCategoriesResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDisassociateVmCategoriesResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools Get operation
*/
type GetNutanixGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetNutanixGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetNutanixGuestToolsResponse() *GetNutanixGuestToolsResponse {
	p := new(GetNutanixGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.GetNutanixGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.GetNutanixGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetNutanixGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetNutanixGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetNutanixGuestToolsResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId} Get operation
*/
type GetVmResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetVmResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetVmResponse() *GetVmResponse {
	p := new(GetVmResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.GetVmResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.GetVmResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetVmResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetVmResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetVmResponseData()
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
Reference to the host, the VM is running on.
*/
type HostReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Globally unique identifier of a host. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewHostReference() *HostReference {
	p := new(HostReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.HostReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.HostReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/insert-iso Post operation
*/
type InsertNutanixVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInsertNutanixVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewInsertNutanixVmGuestToolsResponse() *InsertNutanixVmGuestToolsResponse {
	p := new(InsertNutanixVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.InsertNutanixVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.InsertNutanixVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InsertNutanixVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InsertNutanixVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInsertNutanixVmGuestToolsResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/install Post operation
*/
type InstallNutanixVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfInstallNutanixVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewInstallNutanixVmGuestToolsResponse() *InstallNutanixVmGuestToolsResponse {
	p := new(InstallNutanixVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.InstallNutanixVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.InstallNutanixVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *InstallNutanixVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *InstallNutanixVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfInstallNutanixVmGuestToolsResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms Get operation
*/
type ListVmsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListVmsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewListVmsResponse() *ListVmsResponse {
	p := new(ListVmsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.ListVmsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.ListVmsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListVmsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListVmsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListVmsResponseData()
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
Sign in credentials for the server.
*/
type NutanixCredential struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Password *string `json:"password,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewNutanixCredential() *NutanixCredential {
	p := new(NutanixCredential)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixCredential"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixCredential"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The details about Nutanix Guest Tools for a VM.
*/
type NutanixGuestTools struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Version of Nutanix Guest Tools available on the cluster.
	*/
	AvailableVersion *string `json:"availableVersion,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NutanixGuestToolsCapability `json:"capabilities,omitempty"`
	/*
	  Version of the operating system on the VM.
	*/
	GuestOsVersion *string `json:"guestOsVersion,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools is enabled or not.
	*/
	IsEnabled *bool `json:"isEnabled,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools is installed on the VM or not.
	*/
	IsInstalled *bool `json:"isInstalled,omitempty"`
	/*
	  Indicates whether Nutanix Guest Tools ISO is inserted or not.
	*/
	IsIsoInserted *bool `json:"isIsoInserted,omitempty"`
	/*
	  Indicates whether the communication from VM to CVM is active or not.
	*/
	IsReachable *bool `json:"isReachable,omitempty"`
	/*
	  Indicates whether the VM mobility drivers are installed on the VM or not.
	*/
	IsVmMobilityDriversInstalled *bool `json:"isVmMobilityDriversInstalled,omitempty"`
	/*
	  Indicates whether the VM is configured to take VSS snapshots through NGT or not.
	*/
	IsVssSnapshotCapable *bool `json:"isVssSnapshotCapable,omitempty"`
	/*
	  Version of Nutanix Guest Tools installed on the VM.
	*/
	Version *string `json:"version,omitempty"`
}

func NewNutanixGuestTools() *NutanixGuestTools {
	p := new(NutanixGuestTools)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixGuestTools"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixGuestTools"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The capabilities of the Nutanix Guest Tools in the VM.
*/
type NutanixGuestToolsCapability int

const (
	NUTANIXGUESTTOOLSCAPABILITY_UNKNOWN              NutanixGuestToolsCapability = 0
	NUTANIXGUESTTOOLSCAPABILITY_REDACTED             NutanixGuestToolsCapability = 1
	NUTANIXGUESTTOOLSCAPABILITY_SELF_SERVICE_RESTORE NutanixGuestToolsCapability = 2
	NUTANIXGUESTTOOLSCAPABILITY_VSS_SNAPSHOT         NutanixGuestToolsCapability = 3
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NutanixGuestToolsCapability) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NutanixGuestToolsCapability) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NutanixGuestToolsCapability) index(name string) NutanixGuestToolsCapability {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SELF_SERVICE_RESTORE",
		"VSS_SNAPSHOT",
	}
	for idx := range names {
		if names[idx] == name {
			return NutanixGuestToolsCapability(idx)
		}
	}
	return NUTANIXGUESTTOOLSCAPABILITY_UNKNOWN
}

func (e *NutanixGuestToolsCapability) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NutanixGuestToolsCapability:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NutanixGuestToolsCapability) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NutanixGuestToolsCapability) Ref() *NutanixGuestToolsCapability {
	return &e
}

/*
Argument for inserting a Nutanix Guest Tools ISO into an available slot.
*/
type NutanixGuestToolsInsertConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NutanixGuestToolsCapability `json:"capabilities,omitempty"`
}

func NewNutanixGuestToolsInsertConfig() *NutanixGuestToolsInsertConfig {
	p := new(NutanixGuestToolsInsertConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixGuestToolsInsertConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixGuestToolsInsertConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Arguments for installing Nutanix Guest Tools.
*/
type NutanixGuestToolsInstallConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The list of the application names that are enabled on the guest VM.
	*/
	Capabilities []NutanixGuestToolsCapability `json:"capabilities,omitempty"`

	Credential *NutanixCredential `json:"credential,omitempty"`

	RebootPreference *NutanixRebootPreference `json:"rebootPreference,omitempty"`
}

func NewNutanixGuestToolsInstallConfig() *NutanixGuestToolsInstallConfig {
	p := new(NutanixGuestToolsInstallConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixGuestToolsInstallConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixGuestToolsInstallConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Arguments for upgrading Nutanix Guest Tools.
*/
type NutanixGuestToolsUpgradeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	RebootPreference *NutanixRebootPreference `json:"rebootPreference,omitempty"`
}

func NewNutanixGuestToolsUpgradeConfig() *NutanixGuestToolsUpgradeConfig {
	p := new(NutanixGuestToolsUpgradeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixGuestToolsUpgradeConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixGuestToolsUpgradeConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The restart schedule after installing or upgrading Nutanix Guest Tools.
*/
type NutanixRebootPreference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Schedule *NutanixRebootPreferenceSchedule `json:"schedule,omitempty"`

	ScheduleType *NutanixScheduleType `json:"scheduleType,omitempty"`
}

func NewNutanixRebootPreference() *NutanixRebootPreference {
	p := new(NutanixRebootPreference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixRebootPreference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixRebootPreference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Restart schedule.
*/
type NutanixRebootPreferenceSchedule struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The start time for a scheduled restart.
	*/
	StartTime *time.Time `json:"startTime,omitempty"`
}

func NewNutanixRebootPreferenceSchedule() *NutanixRebootPreferenceSchedule {
	p := new(NutanixRebootPreferenceSchedule)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.NutanixRebootPreferenceSchedule"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.NutanixRebootPreferenceSchedule"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Schedule type for restart.
*/
type NutanixScheduleType int

const (
	NUTANIXSCHEDULETYPE_UNKNOWN   NutanixScheduleType = 0
	NUTANIXSCHEDULETYPE_REDACTED  NutanixScheduleType = 1
	NUTANIXSCHEDULETYPE_SKIP      NutanixScheduleType = 2
	NUTANIXSCHEDULETYPE_IMMEDIATE NutanixScheduleType = 3
	NUTANIXSCHEDULETYPE_LATER     NutanixScheduleType = 4
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *NutanixScheduleType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e NutanixScheduleType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *NutanixScheduleType) index(name string) NutanixScheduleType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SKIP",
		"IMMEDIATE",
		"LATER",
	}
	for idx := range names {
		if names[idx] == name {
			return NutanixScheduleType(idx)
		}
	}
	return NUTANIXSCHEDULETYPE_UNKNOWN
}

func (e *NutanixScheduleType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NutanixScheduleType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NutanixScheduleType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NutanixScheduleType) Ref() *NutanixScheduleType {
	return &e
}

/*
Owner reference.
*/
type OwnerReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EntityType *OwnerReferenceEntityType `json:"entityType,omitempty"`
	/*
	  Globally unique identifier of an instance. It should be of type UUID.
	*/
	ExtId *string `json:"extId,omitempty"`
}

func NewOwnerReference() *OwnerReference {
	p := new(OwnerReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.OwnerReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.OwnerReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OwnerReferenceEntityType int

const (
	OWNERREFERENCEENTITYTYPE_UNKNOWN  OwnerReferenceEntityType = 0
	OWNERREFERENCEENTITYTYPE_REDACTED OwnerReferenceEntityType = 1
	OWNERREFERENCEENTITYTYPE_USER     OwnerReferenceEntityType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *OwnerReferenceEntityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e OwnerReferenceEntityType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *OwnerReferenceEntityType) index(name string) OwnerReferenceEntityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"USER",
	}
	for idx := range names {
		if names[idx] == name {
			return OwnerReferenceEntityType(idx)
		}
	}
	return OWNERREFERENCEENTITYTYPE_UNKNOWN
}

func (e *OwnerReferenceEntityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for OwnerReferenceEntityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *OwnerReferenceEntityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e OwnerReferenceEntityType) Ref() *OwnerReferenceEntityType {
	return &e
}

/*
Ownership information for the VM.
*/
type OwnershipInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Owner *OwnerReference `json:"owner,omitempty"`
}

func NewOwnershipInfo() *OwnershipInfo {
	p := new(OwnershipInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.OwnershipInfo"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.OwnershipInfo"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/uninstall Post operation
*/
type UninstallNutanixVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUninstallNutanixVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUninstallNutanixVmGuestToolsResponse() *UninstallNutanixVmGuestToolsResponse {
	p := new(UninstallNutanixVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.UninstallNutanixVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.UninstallNutanixVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UninstallNutanixVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UninstallNutanixVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUninstallNutanixVmGuestToolsResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools Put operation
*/
type UpdateNutanixGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateNutanixGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpdateNutanixGuestToolsResponse() *UpdateNutanixGuestToolsResponse {
	p := new(UpdateNutanixGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.UpdateNutanixGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.UpdateNutanixGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateNutanixGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateNutanixGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateNutanixGuestToolsResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/esxi/config/vms/{extId}/nutanix-guest-tools/$actions/upgrade Post operation
*/
type UpgradeNutanixVmGuestToolsResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpgradeNutanixVmGuestToolsResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewUpgradeNutanixVmGuestToolsResponse() *UpgradeNutanixVmGuestToolsResponse {
	p := new(UpgradeNutanixVmGuestToolsResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.UpgradeNutanixVmGuestToolsResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.UpgradeNutanixVmGuestToolsResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpgradeNutanixVmGuestToolsResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpgradeNutanixVmGuestToolsResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpgradeNutanixVmGuestToolsResponseData()
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
VM configuration.
*/
type Vm struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Categories for the VM.
	*/
	Categories []CategoryReference `json:"categories,omitempty"`

	Cluster *ClusterReference `json:"cluster,omitempty"`
	/*
	  VM description.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	Host *HostReference `json:"host,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import3.ApiLink `json:"links,omitempty"`
	/*
	  VM name.
	*/
	Name *string `json:"name,omitempty"`

	NutanixGuestTools *NutanixGuestTools `json:"nutanixGuestTools,omitempty"`

	OwnershipInfo *OwnershipInfo `json:"ownershipInfo,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewVm() *Vm {
	p := new(Vm)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.esxi.config.Vm"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.esxi.config.Vm"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfDisassociateVmCategoriesResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfDisassociateVmCategoriesResponseData() *OneOfDisassociateVmCategoriesResponseData {
	p := new(OneOfDisassociateVmCategoriesResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDisassociateVmCategoriesResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDisassociateVmCategoriesResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfDisassociateVmCategoriesResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDisassociateVmCategoriesResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDisassociateVmCategoriesResponseData"))
}

func (p *OneOfDisassociateVmCategoriesResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDisassociateVmCategoriesResponseData")
}

type OneOfAssignVmOwnerResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfAssignVmOwnerResponseData() *OneOfAssignVmOwnerResponseData {
	p := new(OneOfAssignVmOwnerResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssignVmOwnerResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssignVmOwnerResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfAssignVmOwnerResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfAssignVmOwnerResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssignVmOwnerResponseData"))
}

func (p *OneOfAssignVmOwnerResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfAssignVmOwnerResponseData")
}

type OneOfInsertNutanixVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfInsertNutanixVmGuestToolsResponseData() *OneOfInsertNutanixVmGuestToolsResponseData {
	p := new(OneOfInsertNutanixVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInsertNutanixVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInsertNutanixVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfInsertNutanixVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfInsertNutanixVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInsertNutanixVmGuestToolsResponseData"))
}

func (p *OneOfInsertNutanixVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfInsertNutanixVmGuestToolsResponseData")
}

type OneOfUninstallNutanixVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfUninstallNutanixVmGuestToolsResponseData() *OneOfUninstallNutanixVmGuestToolsResponseData {
	p := new(OneOfUninstallNutanixVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUninstallNutanixVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUninstallNutanixVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfUninstallNutanixVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUninstallNutanixVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUninstallNutanixVmGuestToolsResponseData"))
}

func (p *OneOfUninstallNutanixVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUninstallNutanixVmGuestToolsResponseData")
}

type OneOfGetVmResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *Vm                    `json:"-"`
}

func NewOneOfGetVmResponseData() *OneOfGetVmResponseData {
	p := new(OneOfGetVmResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetVmResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetVmResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Vm:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(Vm)
		}
		*p.oneOfType2001 = v.(Vm)
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

func (p *OneOfGetVmResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetVmResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.esxi.config.Vm" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(Vm)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetVmResponseData"))
}

func (p *OneOfGetVmResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetVmResponseData")
}

type OneOfAssociateVmCategoriesResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfAssociateVmCategoriesResponseData() *OneOfAssociateVmCategoriesResponseData {
	p := new(OneOfAssociateVmCategoriesResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfAssociateVmCategoriesResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfAssociateVmCategoriesResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfAssociateVmCategoriesResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfAssociateVmCategoriesResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfAssociateVmCategoriesResponseData"))
}

func (p *OneOfAssociateVmCategoriesResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfAssociateVmCategoriesResponseData")
}

type OneOfUpdateNutanixGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfUpdateNutanixGuestToolsResponseData() *OneOfUpdateNutanixGuestToolsResponseData {
	p := new(OneOfUpdateNutanixGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateNutanixGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateNutanixGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfUpdateNutanixGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateNutanixGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateNutanixGuestToolsResponseData"))
}

func (p *OneOfUpdateNutanixGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateNutanixGuestToolsResponseData")
}

type OneOfInstallNutanixVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfInstallNutanixVmGuestToolsResponseData() *OneOfInstallNutanixVmGuestToolsResponseData {
	p := new(OneOfInstallNutanixVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfInstallNutanixVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfInstallNutanixVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfInstallNutanixVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfInstallNutanixVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfInstallNutanixVmGuestToolsResponseData"))
}

func (p *OneOfInstallNutanixVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfInstallNutanixVmGuestToolsResponseData")
}

type OneOfListVmsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 []Vm                   `json:"-"`
}

func NewOneOfListVmsResponseData() *OneOfListVmsResponseData {
	p := new(OneOfListVmsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListVmsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListVmsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []Vm:
		p.oneOfType2001 = v.([]Vm)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.esxi.config.Vm>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.esxi.config.Vm>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListVmsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.esxi.config.Vm>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListVmsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new([]Vm)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || "vmm.v4.esxi.config.Vm" == *((*vOneOfType2001)[0].ObjectType_) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.esxi.config.Vm>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.esxi.config.Vm>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListVmsResponseData"))
}

func (p *OneOfListVmsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.esxi.config.Vm>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListVmsResponseData")
}

type OneOfUpgradeNutanixVmGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType2001 *import1.TaskReference `json:"-"`
}

func NewOneOfUpgradeNutanixVmGuestToolsResponseData() *OneOfUpgradeNutanixVmGuestToolsResponseData {
	p := new(OneOfUpgradeNutanixVmGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpgradeNutanixVmGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpgradeNutanixVmGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import1.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import1.TaskReference)
		}
		*p.oneOfType2001 = v.(import1.TaskReference)
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

func (p *OneOfUpgradeNutanixVmGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpgradeNutanixVmGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	vOneOfType2001 := new(import1.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import1.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpgradeNutanixVmGuestToolsResponseData"))
}

func (p *OneOfUpgradeNutanixVmGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpgradeNutanixVmGuestToolsResponseData")
}

type OneOfGetNutanixGuestToolsResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType2001 *NutanixGuestTools     `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfGetNutanixGuestToolsResponseData() *OneOfGetNutanixGuestToolsResponseData {
	p := new(OneOfGetNutanixGuestToolsResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetNutanixGuestToolsResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetNutanixGuestToolsResponseData is nil"))
	}
	switch v.(type) {
	case NutanixGuestTools:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(NutanixGuestTools)
		}
		*p.oneOfType2001 = v.(NutanixGuestTools)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
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

func (p *OneOfGetNutanixGuestToolsResponseData) GetValue() interface{} {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfGetNutanixGuestToolsResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType2001 := new(NutanixGuestTools)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if "vmm.v4.esxi.config.NutanixGuestTools" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(NutanixGuestTools)
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
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "vmm.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetNutanixGuestToolsResponseData"))
}

func (p *OneOfGetNutanixGuestToolsResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfGetNutanixGuestToolsResponseData")
}
