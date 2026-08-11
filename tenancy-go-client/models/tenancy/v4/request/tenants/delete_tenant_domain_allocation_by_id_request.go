package tenants

// This file holds the request struct for the DeleteTenantDomainAllocationById operation.

type DeleteTenantDomainAllocationByIdRequest struct {
	// (required) External identifier of the tenant.
	TenantExtId *string

	// (required) External identifier of the tenant domain allocation.
	ExtId *string
}
