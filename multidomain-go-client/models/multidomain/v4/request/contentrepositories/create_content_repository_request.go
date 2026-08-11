package contentrepositories

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/content"
)

// This file holds the request struct for the CreateContentRepository operation.

type CreateContentRepositoryRequest struct {
	// (required) Create a Content Repository request.
	Body *import1.ContentRepository
}
