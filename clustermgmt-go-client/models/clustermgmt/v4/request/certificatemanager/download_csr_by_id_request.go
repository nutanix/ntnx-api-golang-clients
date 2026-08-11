package certificatemanager

// This file holds the request struct for the DownloadCsrById operation.

type DownloadCsrByIdRequest struct {
	// (required) The unique external identifier for the cluster.
	ClusterExtId *string

	// (required) External identifier of the Certificate Signing Request (CSR).
	CsrExtId *string
}
