package recoverypoints

// This file holds the request struct for the GetRecoveryPointById operation.

type GetRecoveryPointByIdRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	ExtId *string
}
