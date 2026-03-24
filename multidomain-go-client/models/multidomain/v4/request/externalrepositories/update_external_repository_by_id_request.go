package externalrepositories

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the UpdateExternalRepositoryById operation.

type UpdateExternalRepositoryByIdRequest struct {
	// (required) External repository identifier.
	ExtId *string

	// (required) Update an external repository request body.
	Body *import1.ExternalRepository
}
