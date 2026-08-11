package contentrepositories

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/content"
)

// This file holds the request struct for the UpdateContentRepositoryById operation.

type UpdateContentRepositoryByIdRequest struct {
	// (required) The external identifier of a Content Repository.
	ExtId *string

	// (required) Update Content Repository request.
	Body *import1.ContentRepository
}
