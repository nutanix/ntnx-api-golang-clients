package recoveryplanactions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/operations"
)

// This file holds the request struct for the UnplannedFailoverRecoveryPlan operation.

type UnplannedFailoverRecoveryPlanRequest struct {
	// (required) The external identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) The unplanned failover recovery plan action is used to perform an unplanned failover on the recovery plan.
	Body *import4.UnplannedFailoverSpec
}
