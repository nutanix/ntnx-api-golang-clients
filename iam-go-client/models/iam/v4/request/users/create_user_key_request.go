package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the CreateUserKey operation.

type CreateUserKeyRequest struct {
	// (required) External identifier of a user of type service account.
	UserExtId *string

	// (required) Create key of a requested type for a user.
	Body *import4.Key
}
