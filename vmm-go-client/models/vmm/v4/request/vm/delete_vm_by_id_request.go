package vm

// This file holds the request struct for the DeleteVmById operation.

type DeleteVmByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
