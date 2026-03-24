package vm

// This file holds the request struct for the PowerCycleVm operation.

type PowerCycleVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
