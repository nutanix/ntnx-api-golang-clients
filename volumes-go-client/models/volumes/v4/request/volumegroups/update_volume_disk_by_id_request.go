package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the UpdateVolumeDiskById operation.

type UpdateVolumeDiskByIdRequest struct {
	// (required) The external identifier of a Volume Group.
	VolumeGroupExtId *string

	// (required) The external identifier of a Volume Disk.
	ExtId *string

	// (required) A model that represents a Volume Disk associated with a Volume Group, supported by a backing file on DSF.
	Body *import1.VolumeDisk
}
