package domainmanagerbackups

// This file holds the request struct for the GetRestoreSourceById operation.

type GetRestoreSourceByIdRequest struct {
	// (required) A globally unique identifier of an instance that is suitable for external consumption.
	ExtId *string
}
