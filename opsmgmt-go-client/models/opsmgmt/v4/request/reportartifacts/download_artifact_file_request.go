package reportartifacts

// This file holds the request struct for the DownloadArtifactFile operation.

type DownloadArtifactFileRequest struct {
	// (required) Report artifact ID.
	ReportArtifactExtId *string
}
