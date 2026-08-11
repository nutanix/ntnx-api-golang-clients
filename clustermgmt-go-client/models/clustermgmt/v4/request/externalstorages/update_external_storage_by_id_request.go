package externalstorages

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateExternalStorageById operation.

type UpdateExternalStorageByIdRequest struct {
	// (required) The unique identifier (UUID) of the external storage.
	ExtId *string

	// (required) Updated configuration for the external storage.
	Body *import1.ExternalStorage
}
