package esxivm

// This file holds the request struct for the PowerOnVm operation.

type PowerOnVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
