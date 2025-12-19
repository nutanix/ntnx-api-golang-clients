/*
 * Generated file models/iam/v4/tenant/tenant_model.go.
 *
 * Product version: 4.1.1-beta-2
 *
 * Part of the Nutanix Identity and Access Management APIs
 *
 * (c) 2025 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module iam.v4.tenant of Nutanix Identity and Access Management APIs
*/
package tenant

import (
	"encoding/json"
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/common/v1/response"
)

/*
Tenant details for tenant to be onboarded to IAM.
*/
type Tenant struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  Description of tenant for which configuration is being set up.
	*/
	TenantDescription *string `json:"tenantDescription,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of tenant for which configuration is being set up.
	*/
	TenantName *string `json:"tenantName,omitempty"`
}

func (p *Tenant) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias Tenant

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

func (p *Tenant) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias Tenant
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewTenant()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.TenantDescription != nil {
		p.TenantDescription = known.TenantDescription
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TenantName != nil {
		p.TenantName = known.TenantName
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "extId")
	delete(allFields, "links")
	delete(allFields, "tenantDescription")
	delete(allFields, "tenantId")
	delete(allFields, "tenantName")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.tenant.Tenant"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r1.b2"}
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
