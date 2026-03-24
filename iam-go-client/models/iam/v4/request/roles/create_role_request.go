package roles

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// This file holds the request struct for the CreateRole operation.

type CreateRoleRequest struct {
	// (required) Creates a role.
	Body *import1.Role
}
