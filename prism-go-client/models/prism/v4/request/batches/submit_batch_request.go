package batches

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/operations"
)

// This file holds the request struct for the SubmitBatch operation.

type SubmitBatchRequest struct {
	// (required) The request payload for the batch operation.
	Body *import1.BatchSpec
}
