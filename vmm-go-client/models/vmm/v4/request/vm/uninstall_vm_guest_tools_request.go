package vm

// This file holds the request struct for the UninstallVmGuestTools operation.

type UninstallVmGuestToolsRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
