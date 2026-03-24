package clusters

// This file holds the request struct for the DeleteClusterById operation.

type DeleteClusterByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ExtId *string

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
