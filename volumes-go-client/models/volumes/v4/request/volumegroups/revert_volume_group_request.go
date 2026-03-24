package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the RevertVolumeGroup operation.

type RevertVolumeGroupRequest struct {
	// (required) The external identifier of a Volume Group.
	ExtId *string

	// (required) Specify the Volume Group recovery point ID to which the Volume Group would be reverted.
	Body *import1.RevertSpec
}
