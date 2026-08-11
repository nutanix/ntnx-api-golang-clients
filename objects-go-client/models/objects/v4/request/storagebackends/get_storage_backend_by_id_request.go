package storagebackends

// This file holds the request struct for the GetStorageBackendById operation.

type GetStorageBackendByIdRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) The extId of the Storage backend.
	ExtId *string
}
