package domainmanagerbackups

// This file holds the request struct for the ListBackupTargets operation.

type ListBackupTargetsRequest struct {
	// (required) A unique identifier for the domain manager.
	DomainManagerExtId *string
}
