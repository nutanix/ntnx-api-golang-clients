package users

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the CreateUser operation.

type CreateUserRequest struct {
	// (required) Information for create user request. It includes details like username, password, userType, emailId, etc.
	Body *import4.User
}
