package tenants

// This file holds the request struct for the DeleteTenantById operation.

type DeleteTenantByIdRequest struct {
	// (required) External identifier of the tenant.
	ExtId *string
}
