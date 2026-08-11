package tenants

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
)

// This file holds the request struct for the UpdateTenantDomainAllocationById operation.

type UpdateTenantDomainAllocationByIdRequest struct {
	// (required) External identifier of the tenant.
	TenantExtId *string

	// (required) External identifier of the tenant domain allocation.
	ExtId *string

	// (required)
	Body *import1.DomainAllocation
}
