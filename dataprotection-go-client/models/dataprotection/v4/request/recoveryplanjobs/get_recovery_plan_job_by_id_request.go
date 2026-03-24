package recoveryplanjobs

// This file holds the request struct for the GetRecoveryPlanJobById operation.

type GetRecoveryPlanJobByIdRequest struct {
	// (required) The external identifier of the recovery plan job.
	ExtId *string
}
