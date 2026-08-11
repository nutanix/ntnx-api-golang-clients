package storageconfig

// This file holds the request struct for the GetStorageConfig operation.

type GetStorageConfigRequest struct {
	// (required) The unique identifier (UUID) of the cluster.
	ClusterExtId *string
}
