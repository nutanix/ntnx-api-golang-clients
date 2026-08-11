package tenants

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/tenancy-go-client/v4/models/tenancy/v4/config"
)

// This file holds the request struct for the CreateTenant operation.

type CreateTenantRequest struct {
	// (required)
	Body *import1.Tenant
}
