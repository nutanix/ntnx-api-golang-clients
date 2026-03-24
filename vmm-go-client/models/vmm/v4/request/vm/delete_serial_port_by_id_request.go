package vm

// This file holds the request struct for the DeleteSerialPortById operation.

type DeleteSerialPortByIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// (required) The globally unique identifier of a serial port of type UUID.
	ExtId *string
}
