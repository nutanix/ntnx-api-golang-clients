package recoveryplans

// This file holds the request struct for the GetRecoveryPlanById operation.

type GetRecoveryPlanByIdRequest struct {
	// (required) External identifier of the recovery plan.
	ExtId *string
}
