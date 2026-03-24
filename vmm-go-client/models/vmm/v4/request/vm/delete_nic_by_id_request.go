package vm

// This file holds the request struct for the DeleteNicById operation.

type DeleteNicByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) A globally unique identifier of a VM NIC of type UUID.
	ExtId *string
}
