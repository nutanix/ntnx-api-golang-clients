package domainmanagerbackups

// This file holds the request struct for the GetRestorePointById operation.

type GetRestorePointByIdRequest struct {
	// (required) A unique identifier obtained from the restore source API that corresponds to the details provided for the
	// restore source.
	RestoreSourceExtId *string

	// (required) A unique identifier for the domain manager.
	RestorableDomainManagerExtId *string

	// (required) Restore point ID for the backup created in cluster/object store.
	ExtId *string
}
