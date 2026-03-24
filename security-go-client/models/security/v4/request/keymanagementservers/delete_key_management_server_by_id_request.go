package keymanagementservers

// This file holds the request struct for the DeleteKeyManagementServerById operation.

type DeleteKeyManagementServerByIdRequest struct {
	// (required) Unique identifier for the key management server of type UUID.
	ExtId *string
}
