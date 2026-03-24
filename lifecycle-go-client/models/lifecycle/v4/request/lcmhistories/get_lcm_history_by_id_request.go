package lcmhistories

// This file holds the request struct for the GetLcmHistoryById operation.

type GetLcmHistoryByIdRequest struct {
	// (required) UUID of the LCM history.
	ExtId *string
}
