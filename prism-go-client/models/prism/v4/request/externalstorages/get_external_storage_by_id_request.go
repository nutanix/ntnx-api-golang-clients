package externalstorages

// This file holds the request struct for the GetExternalStorageById operation.

type GetExternalStorageByIdRequest struct {
	// (required) The unique identifier (UUID) of the external storage resource.
	ExtId *string
}
