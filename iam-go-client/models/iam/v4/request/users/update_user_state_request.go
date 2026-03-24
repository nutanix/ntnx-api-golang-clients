package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UpdateUserState operation.

type UpdateUserStateRequest struct {
	// (required) External identifier of a user.
	ExtId *string

	// (required) Updates the active state of a user based on the provided external identifier.
	Body *import4.UserStateUpdate
}
