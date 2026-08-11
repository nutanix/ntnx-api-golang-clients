package patchedimages

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
)

// This file holds the request struct for the CreatePatchedImage operation.

type CreatePatchedImageRequest struct {
	// (required)
	Body *import3.PatchedImage
}
