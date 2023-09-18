/*
 * Generated file models/vmm/v4/templates/templates_model.go.
 *
 * Product version: 4.0.3-alpha-1
 *
 * Part of the Nutanix Vmm Versioned APIs
 *
 * (c) 2023 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module vmm.v4.templates of Nutanix Vmm Versioned APIs
*/
package templates

import (
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/prism/v4/config"
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/config"
	import2 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/error"
	"time"
)

/*
Input to Template Complete Guest OS Update.
*/
type CompleteGuestUpdateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Indicates whether the version has to be a Gold version or not.
	*/
	SetActiveVersion *bool `json:"setActiveVersion,omitempty"`
	/*
	  A description for template version.
	*/
	VersionDescription *string `json:"versionDescription"`
	/*
	  Template version name. This will be the default VM name prefix when a VM is deployed from this template.
	*/
	VersionName *string `json:"versionName"`
}

func (p *CompleteGuestUpdateSpec) MarshalJSON() ([]byte, error) {
	type CompleteGuestUpdateSpecProxy CompleteGuestUpdateSpec
	return json.Marshal(struct {
		*CompleteGuestUpdateSpecProxy
		VersionDescription *string `json:"versionDescription,omitempty"`
		VersionName        *string `json:"versionName,omitempty"`
	}{
		CompleteGuestUpdateSpecProxy: (*CompleteGuestUpdateSpecProxy)(p),
		VersionDescription:           p.VersionDescription,
		VersionName:                  p.VersionName,
	})
}

func NewCompleteGuestUpdateSpec() *CompleteGuestUpdateSpec {
	p := new(CompleteGuestUpdateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.CompleteGuestUpdateSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.CompleteGuestUpdateSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	p.SetActiveVersion = new(bool)
	*p.SetActiveVersion = false

	return p
}

/*
VM reference for Template guest updates.
*/
type GuestUpdateStatus struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  VM reference UUID.
	*/
	DeployedVmReference *string `json:"deployedVmReference,omitempty"`
}

func NewGuestUpdateStatus() *GuestUpdateStatus {
	p := new(GuestUpdateStatus)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.GuestUpdateStatus"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.GuestUpdateStatus"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Input to Template Guest OS update.
*/
type InitiateGuestUpdateSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Template Version number
	*/
	VersionNumber *int `json:"versionNumber,omitempty"`
}

func NewInitiateGuestUpdateSpec() *InitiateGuestUpdateSpec {
	p := new(InitiateGuestUpdateSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.InitiateGuestUpdateSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.InitiateGuestUpdateSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Template struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The current gold version.
	*/
	ActiveVersionNumber *int `json:"activeVersionNumber,omitempty"`
	/*
	  Created Time.
	*/
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	/*
	  The user who created the Template.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`

	GuestUpdateStatus *GuestUpdateStatus `json:"guestUpdateStatus,omitempty"`
	/*
	  Last Update Time.
	*/
	LastUpdatedAt *time.Time `json:"lastUpdatedAt,omitempty"`
	/*
	  The user who updated the Version in Template.
	*/
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Indicates whether the version has to be a Gold version or not.
	*/
	SetActiveVersion *bool `json:"setActiveVersion,omitempty"`
	/*
	  A description for the Template.
	*/
	TemplateDescription *string `json:"templateDescription,omitempty"`
	/*
	  Template Name.
	*/
	TemplateName *string `json:"templateName,omitempty"`

	TemplateVersionSpec *TemplateVersionSpec `json:"templateVersionSpec,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
}

func NewTemplate() *Template {
	p := new(Template)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.Template"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.Template"}
	p.UnknownFields_ = map[string]interface{}{}

	p.SetActiveVersion = new(bool)
	*p.SetActiveVersion = true

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/templates/{extId} Get operation
*/
type TemplateApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTemplateApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateApiResponse() *TemplateApiResponse {
	p := new(TemplateApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TemplateApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TemplateApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTemplateApiResponseData()
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
Deploy VM Config from the Template.
*/
type TemplateDeployment struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The UUID of the Cluster where the VM has to be deployed. This is mandatory to be specified for creating the VM.
	*/
	ClusterReference *string `json:"clusterReference"`
	/*
	  Number / Count of VMs to be deployed.
	*/
	NumberOfVms *int `json:"numberOfVms"`
	/*
	  The map containing keys as VM index and corresponding VM override config as values which is used at the time of deployment.
	*/
	OverrideVmConfigMap map[string]VmConfigOverride `json:"overrideVmConfigMap,omitempty"`
	/*
	  Template Version number
	*/
	VersionNumber *int `json:"versionNumber,omitempty"`
}

func (p *TemplateDeployment) MarshalJSON() ([]byte, error) {
	type TemplateDeploymentProxy TemplateDeployment
	return json.Marshal(struct {
		*TemplateDeploymentProxy
		ClusterReference *string `json:"clusterReference,omitempty"`
		NumberOfVms      *int    `json:"numberOfVms,omitempty"`
	}{
		TemplateDeploymentProxy: (*TemplateDeploymentProxy)(p),
		ClusterReference:        p.ClusterReference,
		NumberOfVms:             p.NumberOfVms,
	})
}

func NewTemplateDeployment() *TemplateDeployment {
	p := new(TemplateDeployment)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateDeployment"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateDeployment"}
	p.UnknownFields_ = map[string]interface{}{}

	p.NumberOfVms = new(int)
	*p.NumberOfVms = 1

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/templates Get operation
*/
type TemplateListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTemplateListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateListApiResponse() *TemplateListApiResponse {
	p := new(TemplateListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TemplateListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TemplateListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTemplateListApiResponseData()
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

type TemplatePublishSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Template's active version which will be used as default for a VM deployment or a guest update
	*/
	ActiveVersionNumber *int `json:"activeVersionNumber"`
}

func (p *TemplatePublishSpec) MarshalJSON() ([]byte, error) {
	type TemplatePublishSpecProxy TemplatePublishSpec
	return json.Marshal(struct {
		*TemplatePublishSpecProxy
		ActiveVersionNumber *int `json:"activeVersionNumber,omitempty"`
	}{
		TemplatePublishSpecProxy: (*TemplatePublishSpecProxy)(p),
		ActiveVersionNumber:      p.ActiveVersionNumber,
	})
}

func NewTemplatePublishSpec() *TemplatePublishSpec {
	p := new(TemplatePublishSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplatePublishSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplatePublishSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/templates/{extId}/versions/{versionNumber} Get operation
*/
type TemplateVersionApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTemplateVersionApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateVersionApiResponse() *TemplateVersionApiResponse {
	p := new(TemplateVersionApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateVersionApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TemplateVersionApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TemplateVersionApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTemplateVersionApiResponseData()
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
REST response for all response codes in API path /vmm/v4.0.a1/templates/{extId}/versions Get operation
*/
type TemplateVersionListApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTemplateVersionListApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplateVersionListApiResponse() *TemplateVersionListApiResponse {
	p := new(TemplateVersionListApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateVersionListApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionListApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TemplateVersionListApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TemplateVersionListApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTemplateVersionListApiResponseData()
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
Template new version with Version Number and Override Config.
*/
type TemplateVersionReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	OverrideVmConfig *VmConfigOverride `json:"overrideVmConfig"`
	/*
	  Template Version number
	*/
	Version *int `json:"version"`
}

func (p *TemplateVersionReference) MarshalJSON() ([]byte, error) {
	type TemplateVersionReferenceProxy TemplateVersionReference
	return json.Marshal(struct {
		*TemplateVersionReferenceProxy
		OverrideVmConfig *VmConfigOverride `json:"overrideVmConfig,omitempty"`
		Version          *int              `json:"version,omitempty"`
	}{
		TemplateVersionReferenceProxy: (*TemplateVersionReferenceProxy)(p),
		OverrideVmConfig:              p.OverrideVmConfig,
		Version:                       p.Version,
	})
}

func NewTemplateVersionReference() *TemplateVersionReference {
	p := new(TemplateVersionReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateVersionReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TemplateVersionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Created Time.
	*/
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	/*
	  The user who created the Version in Template.
	*/
	CreatedBy *string `json:"createdBy,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Allow or Disallow override of the Guest Customization during Template Deployment using this version.
	*/
	IsGcOverrideEnabled *bool `json:"isGcOverrideEnabled,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity.  It is automatically assigned by the system and is immutable from an API consumer perspective (some use cases may cause this Id to change - for instance a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  A description for template version.
	*/
	VersionDescription *string `json:"versionDescription,omitempty"`
	/*
	  Template version name. This will be the default VM name prefix when a VM is deployed from this template.
	*/
	VersionName *string `json:"versionName,omitempty"`
	/*
	  Template Version number
	*/
	VersionNumber *int `json:"versionNumber,omitempty"`

	VersionSourceItemDiscriminator_ *string `json:"$versionSourceItemDiscriminator,omitempty"`
	/*
	  Source from which a version is created. It can be either a VM or an existing version of the template.
	*/
	VersionSource *OneOfTemplateVersionSpecVersionSource `json:"versionSource,omitempty"`

	VersionSourceDiscriminator *string `json:"versionSourceDiscriminator,omitempty"`
	/*
	  VM configuration spec
	*/
	VmSpec *string `json:"vmSpec,omitempty"`
}

func NewTemplateVersionSpec() *TemplateVersionSpec {
	p := new(TemplateVersionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplateVersionSpec"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplateVersionSpec"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /vmm/v4.0.a1/templates/{extId}/$actions/publish Post operation
*/
type TemplatesTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfTemplatesTaskApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewTemplatesTaskApiResponse() *TemplatesTaskApiResponse {
	p := new(TemplatesTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.TemplatesTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.TemplatesTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *TemplatesTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *TemplatesTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfTemplatesTaskApiResponseData()
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
Overrides specification for VM create from a Template.
*/
type VmConfigOverride struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	GuestCustomization *import4.GuestCustomizationParams `json:"guestCustomization,omitempty"`
	/*
	  Memory size in bytes.
	*/
	MemorySizeBytes *int64 `json:"memorySizeBytes,omitempty"`
	/*
	  VM name.
	*/
	Name *string `json:"name,omitempty"`
	/*
	  NICs attached to the VM.
	*/
	Nics []import4.Nic `json:"nics,omitempty"`
	/*
	  Number of cores per socket.
	*/
	NumCoresPerSocket *int `json:"numCoresPerSocket,omitempty"`
	/*
	  Number of vCPU sockets.
	*/
	NumSockets *int `json:"numSockets,omitempty"`
	/*
	  Number of threads per core.
	*/
	NumThreadsPerCore *int `json:"numThreadsPerCore,omitempty"`
}

func NewVmConfigOverride() *VmConfigOverride {
	p := new(VmConfigOverride)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.VmConfigOverride"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.VmConfigOverride"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
The source VM to create a new template
*/
type VmReference struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The source VM to create a new template
	*/
	ExtId *string `json:"extId"`
	/*
	  VM guests may be customized at boot time using one of several different methods. Currently, cloud-init w/ ConfigDriveV2 (for Linux VMs) and Sysprep (for Windows VMs) are supported. Only ONE OF Sysprep or CloudInit should be provided. Note that guest customization can currently only be set during VM creation. Attempting to change it after creation will result in an error. Additional properties can be specified. For example - in the context of VM template creation if 'isOverridable' is set to 'True' then the deployer can upload their own custom script.
	*/
	GuestCustomization *string `json:"guestCustomization,omitempty"`
}

func (p *VmReference) MarshalJSON() ([]byte, error) {
	type VmReferenceProxy VmReference
	return json.Marshal(struct {
		*VmReferenceProxy
		ExtId *string `json:"extId,omitempty"`
	}{
		VmReferenceProxy: (*VmReferenceProxy)(p),
		ExtId:            p.ExtId,
	})
}

func NewVmReference() *VmReference {
	p := new(VmReference)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "vmm.v4.templates.VmReference"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "vmm.v4.r0.a1.templates.VmReference"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfTemplateVersionListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []TemplateVersionSpec  `json:"-"`
}

func NewOneOfTemplateVersionListApiResponseData() *OneOfTemplateVersionListApiResponseData {
	p := new(OneOfTemplateVersionListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplateVersionListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplateVersionListApiResponseData is nil"))
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
	case []TemplateVersionSpec:
		p.oneOfType0 = v.([]TemplateVersionSpec)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.templates.TemplateVersionSpec>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.templates.TemplateVersionSpec>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTemplateVersionListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.templates.TemplateVersionSpec>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfTemplateVersionListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]TemplateVersionSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "vmm.v4.templates.TemplateVersionSpec" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.templates.TemplateVersionSpec>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.templates.TemplateVersionSpec>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateVersionListApiResponseData"))
}

func (p *OneOfTemplateVersionListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.templates.TemplateVersionSpec>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTemplateVersionListApiResponseData")
}

type OneOfTemplateVersionApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *TemplateVersionSpec   `json:"-"`
}

func NewOneOfTemplateVersionApiResponseData() *OneOfTemplateVersionApiResponseData {
	p := new(OneOfTemplateVersionApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplateVersionApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplateVersionApiResponseData is nil"))
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
	case TemplateVersionSpec:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TemplateVersionSpec)
		}
		*p.oneOfType0 = v.(TemplateVersionSpec)
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

func (p *OneOfTemplateVersionApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTemplateVersionApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(TemplateVersionSpec)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.templates.TemplateVersionSpec" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TemplateVersionSpec)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateVersionApiResponseData"))
}

func (p *OneOfTemplateVersionApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTemplateVersionApiResponseData")
}

type OneOfTemplateVersionSpecVersionSource struct {
	Discriminator *string                   `json:"-"`
	ObjectType_   *string                   `json:"-"`
	oneOfType0    *VmReference              `json:"-"`
	oneOfType1    *TemplateVersionReference `json:"-"`
}

func NewOneOfTemplateVersionSpecVersionSource() *OneOfTemplateVersionSpecVersionSource {
	p := new(OneOfTemplateVersionSpecVersionSource)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplateVersionSpecVersionSource) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplateVersionSpecVersionSource is nil"))
	}
	switch v.(type) {
	case VmReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(VmReference)
		}
		*p.oneOfType0 = v.(VmReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case TemplateVersionReference:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(TemplateVersionReference)
		}
		*p.oneOfType1 = v.(TemplateVersionReference)
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

func (p *OneOfTemplateVersionSpecVersionSource) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return *p.oneOfType1
	}
	return nil
}

func (p *OneOfTemplateVersionSpecVersionSource) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(VmReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.templates.VmReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(VmReference)
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
	vOneOfType1 := new(TemplateVersionReference)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if "vmm.v4.templates.TemplateVersionReference" == *vOneOfType1.ObjectType_ {
			if nil == p.oneOfType1 {
				p.oneOfType1 = new(TemplateVersionReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateVersionSpecVersionSource"))
}

func (p *OneOfTemplateVersionSpecVersionSource) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType1 != nil && *p.oneOfType1.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	return nil, errors.New("No value to marshal for OneOfTemplateVersionSpecVersionSource")
}

type OneOfTemplateListApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    []Template             `json:"-"`
}

func NewOneOfTemplateListApiResponseData() *OneOfTemplateListApiResponseData {
	p := new(OneOfTemplateListApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplateListApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplateListApiResponseData is nil"))
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
	case []Template:
		p.oneOfType0 = v.([]Template)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<vmm.v4.templates.Template>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<vmm.v4.templates.Template>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfTemplateListApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<vmm.v4.templates.Template>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfTemplateListApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new([]Template)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if len(*vOneOfType0) == 0 || "vmm.v4.templates.Template" == *((*vOneOfType0)[0].ObjectType_) {
			p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<vmm.v4.templates.Template>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<vmm.v4.templates.Template>"
			return nil

		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateListApiResponseData"))
}

func (p *OneOfTemplateListApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<vmm.v4.templates.Template>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTemplateListApiResponseData")
}

type OneOfTemplatesTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import3.TaskReference `json:"-"`
}

func NewOneOfTemplatesTaskApiResponseData() *OneOfTemplatesTaskApiResponseData {
	p := new(OneOfTemplatesTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplatesTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplatesTaskApiResponseData is nil"))
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
	case import3.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import3.TaskReference)
		}
		*p.oneOfType0 = v.(import3.TaskReference)
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

func (p *OneOfTemplatesTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTemplatesTaskApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(import3.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import3.TaskReference)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplatesTaskApiResponseData"))
}

func (p *OneOfTemplatesTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTemplatesTaskApiResponseData")
}

type OneOfTemplateApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Template              `json:"-"`
}

func NewOneOfTemplateApiResponseData() *OneOfTemplateApiResponseData {
	p := new(OneOfTemplateApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfTemplateApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfTemplateApiResponseData is nil"))
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
	case Template:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Template)
		}
		*p.oneOfType0 = v.(Template)
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

func (p *OneOfTemplateApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfTemplateApiResponseData) UnmarshalJSON(b []byte) error {
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
	vOneOfType0 := new(Template)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "vmm.v4.templates.Template" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Template)
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
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfTemplateApiResponseData"))
}

func (p *OneOfTemplateApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfTemplateApiResponseData")
}
