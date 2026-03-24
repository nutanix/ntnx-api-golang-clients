package vm

// This file holds the request struct for the DeleteGpuById operation.

type DeleteGpuByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) A globally unique identifier of a VM GPU of type UUID.
	ExtId *string
}
