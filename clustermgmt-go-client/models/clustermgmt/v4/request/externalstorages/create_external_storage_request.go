package externalstorages

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the CreateExternalStorage operation.

type CreateExternalStorageRequest struct {
	// (required) Configuration for the new external storage to be created.
	Body *import1.ExternalStorage
}
