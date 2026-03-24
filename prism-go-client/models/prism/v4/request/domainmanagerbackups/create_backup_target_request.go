package domainmanagerbackups

import (
	import5 "github.com/nutanix/ntnx-api-golang-clients/prism-go-client/v4/models/prism/v4/management"
)

// This file holds the request struct for the CreateBackupTarget operation.

type CreateBackupTargetRequest struct {
	// (required) A unique identifier for the domain manager.
	DomainManagerExtId *string

	// (required)
	Body *import5.BackupTarget
}
