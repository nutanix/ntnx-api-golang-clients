package vm

// This file holds the request struct for the DeleteDiskById operation.

type DeleteDiskByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) A globally unique identifier of a VM disk of type UUID.
	ExtId *string
}
