package tenants

// This file holds the request struct for the GetTenantExternalNetworkConnection operation.

type GetTenantExternalNetworkConnectionRequest struct {
	// (required) External identifier of the tenant.
	TenantExtId *string

	// (required) External identifier of the domain allocation.
	DomainAllocationExtId *string
}
