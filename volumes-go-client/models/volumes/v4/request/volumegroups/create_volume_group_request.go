package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the CreateVolumeGroup operation.

type CreateVolumeGroupRequest struct {
	// (required) A model that represents a Volume Group resource.
	Body *import1.VolumeGroup
}
