package patchedimages

// This file holds the request struct for the DeletePatchedImageById operation.

type DeletePatchedImageByIdRequest struct {
	// (required) External ID of a patched hypervisor or host OS image
	ExtId *string
}
