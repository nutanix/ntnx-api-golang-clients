package vm

// This file holds the request struct for the DeletePcieDeviceById operation.

type DeletePcieDeviceByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) Globally unique identifier of a PCIe device of type UUID.
	ExtId *string
}
