package domainmanagerbackups

import (
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// This file holds the request struct for the UpdateBackupTargetById operation.

type UpdateBackupTargetByIdRequest struct {
	// (required) A unique identifier for the domain manager.
	DomainManagerExtId *string

	// (required) A globally unique identifier of an instance that is suitable for external consumption.
	ExtId *string

	// (required)
	Body *import5.BackupTarget
}
