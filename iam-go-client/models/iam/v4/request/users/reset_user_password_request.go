package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the ResetUserPassword operation.

type ResetUserPasswordRequest struct {
	// (required) External identifier of a user.
	ExtId *string

	// (required) Information for reset password request. It requires new password of the user.
	Body *import4.PasswordResetRequest
}
