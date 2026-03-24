package esxivm

// This file holds the request struct for the ResetVm operation.

type ResetVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
