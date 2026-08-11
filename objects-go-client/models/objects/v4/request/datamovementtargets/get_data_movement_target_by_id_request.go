package datamovementtargets

// This file holds the request struct for the GetDataMovementTargetById operation.

type GetDataMovementTargetByIdRequest struct {
	// (required) The UUID of the Object store.
	ObjectStoreExtId *string

	// (required) The UUID of the data movement target in the Object store.
	ExtId *string
}
