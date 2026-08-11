package tenants

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
)

// This file holds the request struct for the CreateTenantDomainAllocation operation.

type CreateTenantDomainAllocationRequest struct {
	// (required) External identifier of the tenant.
	TenantExtId *string

	// (required)
	Body *import1.DomainAllocation
}
