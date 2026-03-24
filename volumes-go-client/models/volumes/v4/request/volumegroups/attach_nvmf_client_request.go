package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the AttachNvmfClient operation.

type AttachNvmfClientRequest struct {
	// (required) The external identifier of a Volume Group.
	ExtId *string

	// (required) A model representing a NVMe-TCP client that can be associated with a Volume Group as an external attachment.
	Body *import1.NvmfClient
}
