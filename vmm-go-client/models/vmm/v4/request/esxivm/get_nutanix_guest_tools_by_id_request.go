package esxivm

// This file holds the request struct for the GetNutanixGuestToolsById operation.

type GetNutanixGuestToolsByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	ExtId *string
}
