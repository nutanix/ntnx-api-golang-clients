package disks

// This file holds the request struct for the GetDiskById operation.

type GetDiskByIdRequest struct {
	// (required) The external identifier of the Disk.
	ExtId *string
}
