package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UpdateUserById operation.

type UpdateUserByIdRequest struct {
	// (required) External identifier of a user.
	ExtId *string

	// (required) Information for updating the user request. It can be used to update fields like displayName, firstName, lastName, email,
	// etc. The username and userType fields cannot be updated.
	Body *import4.User
}
