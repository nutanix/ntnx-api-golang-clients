package directoryservices

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the SearchDirectoryService operation.

type SearchDirectoryServiceRequest struct {
	// (required) External identifier of the directory service.
	ExtId *string

	// (required) Searches a user or group in the directory service through its external identifier.
	Body *import4.DirectoryServiceSearchQuery
}
