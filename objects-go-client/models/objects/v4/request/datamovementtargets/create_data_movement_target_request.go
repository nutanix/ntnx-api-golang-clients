package datamovementtargets

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/objects-go-client/v4/models/objects/v4/config"
)

// This file holds the request struct for the CreateDataMovementTarget operation.

type CreateDataMovementTargetRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) The request body for creating a data movement target.
	Body *import1.DataMovementTarget
}
