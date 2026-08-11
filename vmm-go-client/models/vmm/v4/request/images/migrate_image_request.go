package images

import (
	import11 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the MigrateImage operation.

type MigrateImageRequest struct {
	// (required) The external identifier of an image.
	ExtId *string

	// (required) Reference to the Content Repository to migrate the image to.
	Body *import11.ImageMigrateConfig
}
