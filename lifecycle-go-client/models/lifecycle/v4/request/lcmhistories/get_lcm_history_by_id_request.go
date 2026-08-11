package lcmhistories

// This file holds the request struct for the GetLcmHistoryById operation.

type GetLcmHistoryByIdRequest struct {
	// (required) The external identifier (UUID) of the LCM history entry.
	ExtId *string
}
