package storageconfig

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateStorageConfig operation.

type UpdateStorageConfigRequest struct {
	// (required) The unique identifier (UUID) of the cluster.
	ClusterExtId *string

	// (required)
	Body *import1.StorageConfig
}
