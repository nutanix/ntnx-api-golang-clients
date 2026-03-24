package vm

// This file holds the request struct for the GetGpuById operation.

type GetGpuByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) A globally unique identifier of a VM GPU of type UUID.
	ExtId *string
}
