package patchedimages

// This file holds the request struct for the GetPatchedImageById operation.

type GetPatchedImageByIdRequest struct {
	// (required) External ID of a patched hypervisor or host OS image
	ExtId *string
}
