package users

// This file holds the request struct for the GetUserById operation.

type GetUserByIdRequest struct {
	// (required) External identifier of a user.
	ExtId *string
}
