package roles

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// This file holds the request struct for the UpdateRoleById operation.

type UpdateRoleByIdRequest struct {
	// (required) External identifier for the role.
	ExtId *string

	// (required) Information for the update role request.
	Body *import1.Role
}
