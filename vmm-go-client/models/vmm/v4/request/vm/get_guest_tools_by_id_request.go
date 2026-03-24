package vm

// This file holds the request struct for the GetGuestToolsById operation.

type GetGuestToolsByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
