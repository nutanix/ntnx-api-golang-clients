package storagecontainers

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateStorageContainerById operation.

type UpdateStorageContainerByIdRequest struct {
	// (required) The external identifier of the Storage Container.
	ExtId *string

	// (required)
	Body *import1.StorageContainer
}
