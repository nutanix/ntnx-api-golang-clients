package vm

// This file holds the request struct for the PowerOffVm operation.

type PowerOffVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
