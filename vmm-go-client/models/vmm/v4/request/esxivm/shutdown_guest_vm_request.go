package esxivm

// This file holds the request struct for the ShutdownGuestVm operation.

type ShutdownGuestVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
