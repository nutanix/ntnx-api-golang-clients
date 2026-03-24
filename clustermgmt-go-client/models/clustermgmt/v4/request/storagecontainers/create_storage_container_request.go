package storagecontainers

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateStorageContainer operation.

type CreateStorageContainerRequest struct {
	// (required)
	Body *import1.StorageContainer

	// (required) The external identifier of the remote cluster to which the request is forwarded.
	XClusterId *string
}
