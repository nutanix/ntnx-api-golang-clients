package lcmsummaries

// This file holds the request struct for the GetLcmSummaryById operation.

type GetLcmSummaryByIdRequest struct {
	// (required) The external identifier (UUID) of the cluster whose LCM summary to retrieve.
	ExtId *string
}
