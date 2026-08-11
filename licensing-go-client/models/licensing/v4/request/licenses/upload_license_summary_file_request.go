package licenses

// This file holds the request struct for the UploadLicenseSummaryFile operation.

type UploadLicenseSummaryFileRequest struct {
	// (required) Binary octet-stream payload containing the license summary file.
	Path *string
}
