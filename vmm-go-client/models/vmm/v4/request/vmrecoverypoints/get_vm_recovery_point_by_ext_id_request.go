package vmrecoverypoints

// This file holds the request struct for the GetVmRecoveryPointByExtId operation.

type GetVmRecoveryPointByExtIdRequest struct {
	// (required) A globally unique identifier of a VM recovery point. It should be of type UUID.
	ExtId *string
}
