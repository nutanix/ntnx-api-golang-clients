package esxivm

// This file holds the request struct for the GetVmById operation.

type GetVmByIdRequest struct {
	// (required) The globally unique identifier of an instance of type UUID.
	ExtId *string
}
