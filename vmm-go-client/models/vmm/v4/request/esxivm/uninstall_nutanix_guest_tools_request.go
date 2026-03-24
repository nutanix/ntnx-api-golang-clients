package esxivm

// This file holds the request struct for the UninstallNutanixGuestTools operation.

type UninstallNutanixGuestToolsRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
