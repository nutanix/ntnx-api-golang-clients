package batches

// This file holds the request struct for the GetBatchById operation.

type GetBatchByIdRequest struct {
	// (required) The external identifier of the Batch.
	ExtId *string
}
