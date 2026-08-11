package tenants

// This file holds the request struct for the GetTenantById operation.

type GetTenantByIdRequest struct {
	// (required) External identifier of the tenant.
	ExtId *string
}
