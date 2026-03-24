package images

// This file holds the request struct for the GetFileByImageId operation.

type GetFileByImageIdRequest struct {
	// (required) The external identifier of an image.
	ImageExtId *string
}
