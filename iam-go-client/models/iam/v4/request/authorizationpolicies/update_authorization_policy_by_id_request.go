package authorizationpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// This file holds the request struct for the UpdateAuthorizationPolicyById operation.

type UpdateAuthorizationPolicyByIdRequest struct {
	// (required) External identifier for the role.
	ExtId *string

	// (required) Information for the update authorization policy request.
	Body *import1.AuthorizationPolicy
}
