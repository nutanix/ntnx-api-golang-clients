package lcmsummaries

// This file holds the request struct for the GetLcmSummaryById operation.

type GetLcmSummaryByIdRequest struct {
	// (required) UUID of the LCM summary.
	ExtId *string
}
