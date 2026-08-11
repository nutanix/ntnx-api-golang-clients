package recoveryplanactions

// This file holds the request struct for the ResumeProtection operation.

type ResumeProtectionRequest struct {
	// (required) The external identifier of the recovery plan.
	RecoveryPlanExtId *string
}
