package recoveryplans

// This file holds the request struct for the DeleteRecoverySettingById operation.

type DeleteRecoverySettingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) recovery setting external identifier.
	ExtId *string
}
