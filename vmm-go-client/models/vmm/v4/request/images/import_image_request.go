package images

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the ImportImage operation.

type ImportImageRequest struct {
	// (required) Reference to the Prism Element cluster and its images to be imported.
	Body *import9.ImageImportConfig
}
