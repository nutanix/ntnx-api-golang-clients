package recoveryplanactions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/operations"
)

// This file holds the request struct for the ValidateRecoveryPlan operation.

type ValidateRecoveryPlanRequest struct {
	// (required) The external identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) The validate Recovery Plan action is used to validate the recovery plan.
	Body *import4.BaseRecoveryPlanActionSpec
}
