package domainmanagerbackups

import (
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// This file holds the request struct for the Restore operation.

type RestoreRequest struct {
	// (required) A unique identifier obtained from the restore source API that corresponds to the details provided for the
	// restore source.
	RestoreSourceExtId *string

	// (required) A unique identifier for the domain manager.
	RestorableDomainManagerExtId *string

	// (required) Restore point ID for the backup created in cluster/object store.
	ExtId *string

	// (required)
	Body *import5.RestoreSpec
}
