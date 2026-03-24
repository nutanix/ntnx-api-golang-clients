package images

// This file holds the request struct for the GetImageById operation.

type GetImageByIdRequest struct {
	// (required) The external identifier of an image.
	ExtId *string
}
