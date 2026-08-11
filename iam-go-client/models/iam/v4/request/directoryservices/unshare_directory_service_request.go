package directoryservices

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UnshareDirectoryService operation.

type UnshareDirectoryServiceRequest struct {
	// (required) External identifier of the directory service.
	ExtId *string

	// (required) Unshares a directory service from a specific project.
	Body *import4.DirectoryServiceUnshareRequest
}
