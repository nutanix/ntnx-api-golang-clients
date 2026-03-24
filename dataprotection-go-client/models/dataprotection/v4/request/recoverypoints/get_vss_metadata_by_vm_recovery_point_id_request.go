package recoverypoints

// This file holds the request struct for the GetVssMetadataByVmRecoveryPointId operation.

type GetVssMetadataByVmRecoveryPointIdRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	RecoveryPointExtId *string

	// (required) The external identifier that can be used to identify a VM recovery point.
	VmRecoveryPointExtId *string
}
