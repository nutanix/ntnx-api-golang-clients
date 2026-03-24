package esxivm

// This file holds the request struct for the RebootGuestVm operation.

type RebootGuestVmRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
