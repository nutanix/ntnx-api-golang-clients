package volumegroups

// This file holds the request struct for the DeleteVolumeDiskById operation.

type DeleteVolumeDiskByIdRequest struct {
	// (required) The external identifier of a Volume Group.
	VolumeGroupExtId *string

	// (required) The external identifier of a Volume Disk.
	ExtId *string
}
