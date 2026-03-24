package esxivm

// This file holds the request struct for the SuspendVm operation.

type SuspendVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
