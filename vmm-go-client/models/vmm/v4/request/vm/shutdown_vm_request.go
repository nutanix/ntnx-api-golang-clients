package vm

// This file holds the request struct for the ShutdownVm operation.

type ShutdownVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
