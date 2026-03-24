package images

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the UpdateImageById operation.

type UpdateImageByIdRequest struct {
	// (required) The external identifier of an image.
	ExtId *string

	// (required) Updated image request.
	Body *import9.Image
}
