package reportartifacts

// This file holds the request struct for the UploadArtifactFile operation.

type UploadArtifactFileRequest struct {
	// (required) Report artifact ID.
	ReportArtifactExtId *string

	// (required) Upload a report artifact by providing the file as part of the request body.
	Path *string
}
