package vm

// This file holds the request struct for the GetVmById operation.

type GetVmByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
