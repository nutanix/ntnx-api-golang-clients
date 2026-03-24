package images

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the CreateImage operation.

type CreateImageRequest struct {
	// (required) Request to create an image.
	Body *import9.Image
}
