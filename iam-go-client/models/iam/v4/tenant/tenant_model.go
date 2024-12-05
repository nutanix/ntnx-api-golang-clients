/*
 * Generated file models/iam/v4/tenant/tenant_model.go.
 *
 * Product version: 4.0.1
 *
 * Part of the Nutanix IAM Versioned APIs
 *
 * (c) 2024 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module iam.v4.tenant of Nutanix IAM Versioned APIs
*/
package tenant

import (
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
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this Id to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Name of tenant for which configuration is being set up.
	*/
	TenantName *string `json:"tenantName,omitempty"`
}

func NewTenant() *Tenant {
	p := new(Tenant)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "iam.v4.tenant.Tenant"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r0"}
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
