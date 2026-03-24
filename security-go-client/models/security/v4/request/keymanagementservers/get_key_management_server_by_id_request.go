package keymanagementservers

// This file holds the request struct for the GetKeyManagementServerById operation.

type GetKeyManagementServerByIdRequest struct {
	// (required) Unique identifier for the key management server of type UUID.
	ExtId *string
}
