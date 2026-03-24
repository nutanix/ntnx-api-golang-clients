package storagecontainers

// This file holds the request struct for the DeleteStorageContainerById operation.

type DeleteStorageContainerByIdRequest struct {
	// (required) The external identifier of the Storage Container.
	ExtId *string

	// Indicates whether the small files should be ignored. Storage containers can have small files that are measured in KBs;
	// these files are ignored when the storage container is marked for removal.
	IgnoreSmallFiles *bool
}
