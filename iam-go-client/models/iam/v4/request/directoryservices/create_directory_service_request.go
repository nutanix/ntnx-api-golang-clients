package directoryservices

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the CreateDirectoryService operation.

type CreateDirectoryServiceRequest struct {
	// (required) Creates a directory service.
	Body *import4.DirectoryService
}
