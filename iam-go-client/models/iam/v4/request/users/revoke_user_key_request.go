package users

// This file holds the request struct for the RevokeUserKey operation.

type RevokeUserKeyRequest struct {
	// (required) External identifier of a user of type service account.
	UserExtId *string

	// (required) External identifier of the key.
	ExtId *string
}
