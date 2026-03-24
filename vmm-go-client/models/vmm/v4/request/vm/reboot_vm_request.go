package vm

// This file holds the request struct for the RebootVm operation.

type RebootVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
