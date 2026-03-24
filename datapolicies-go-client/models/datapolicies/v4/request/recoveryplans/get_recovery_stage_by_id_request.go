package recoveryplans

// This file holds the request struct for the GetRecoveryStageById operation.

type GetRecoveryStageByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) External identifier of the Recovery stage.
	ExtId *string
}
