package storagecontainers

// This file holds the request struct for the GetStorageContainerById operation.

type GetStorageContainerByIdRequest struct {
	// (required) The external identifier of the Storage Container.
	ExtId *string
}
