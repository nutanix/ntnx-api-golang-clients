package reports

// This file holds the request struct for the GetReportById operation.

type GetReportByIdRequest struct {
	// (required) UUID of the report.
	ExtId *string
}
