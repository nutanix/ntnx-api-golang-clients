package prechecks

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
)

// This file holds the request struct for the PerformPrechecks operation.

type PerformPrechecksRequest struct {
	// (required)
	Body *import4.PrechecksSpec

	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
