package volumegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/volumes-go-client/v4/models/volumes/v4/config"
)

// This file holds the request struct for the DetachIscsiClient operation.

type DetachIscsiClientRequest struct {
	// (required) The external identifier of a Volume Group.
	ExtId *string

	// (required) A model that represents an iSCSI client that can be associated with a Volume Group as an external attachment. It
	// contains the minimal properties required for the attachment.
	Body *import1.IscsiClientAttachment
}
