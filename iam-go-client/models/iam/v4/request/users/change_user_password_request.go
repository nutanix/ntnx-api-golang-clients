package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the ChangeUserPassword operation.

type ChangeUserPasswordRequest struct {
	// (required) Information for change password request. It requires the username, old password and new password of the user.
	Body *import4.PasswordChangeRequest
}
