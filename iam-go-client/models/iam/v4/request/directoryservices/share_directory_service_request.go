package directoryservices

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the ShareDirectoryService operation.

type ShareDirectoryServiceRequest struct {
	// (required) External identifier of the directory service.
	ExtId *string

	// (required) Shares a directory service with specified project.
	Body *import4.DirectoryServiceShareRequest
}
