package authorizationpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authz"
)

// This file holds the request struct for the CreateAuthorizationPolicy operation.

type CreateAuthorizationPolicyRequest struct {
	// (required) Information for the create authorization policy request. It requires the role, identity and entities attributes.
	Body *import1.AuthorizationPolicy
}
