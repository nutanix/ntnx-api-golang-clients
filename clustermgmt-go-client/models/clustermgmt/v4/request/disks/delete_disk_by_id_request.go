package disks

// This file holds the request struct for the DeleteDiskById operation.

type DeleteDiskByIdRequest struct {
	// (required) The external identifier of the Disk.
	ExtId *string
}
