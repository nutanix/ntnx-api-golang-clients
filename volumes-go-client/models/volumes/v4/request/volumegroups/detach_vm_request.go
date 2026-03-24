package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the DetachVm operation.

type DetachVmRequest struct {
	// (required) The external identifier of a Volume Group.
	ExtId *string

	// (required) A model that represents a VM reference that can be associated with a Volume Group as an AHV hypervisor attachment.
	Body *import1.VmAttachment
}
