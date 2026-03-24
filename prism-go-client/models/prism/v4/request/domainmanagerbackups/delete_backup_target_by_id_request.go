package domainmanagerbackups

// This file holds the request struct for the DeleteBackupTargetById operation.

type DeleteBackupTargetByIdRequest struct {
	// (required) A unique identifier for the domain manager.
	DomainManagerExtId *string

	// (required) A globally unique identifier of an instance that is suitable for external consumption.
	ExtId *string
}
