package reports

// This file holds the request struct for the DownloadReport operation.

type DownloadReportRequest struct {
	// (required) UUID of the report.
	ReportExtId *string
}
