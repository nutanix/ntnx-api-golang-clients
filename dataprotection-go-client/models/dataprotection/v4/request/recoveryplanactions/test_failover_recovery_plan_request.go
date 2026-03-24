package recoveryplanactions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/operations"
)

// This file holds the request struct for the TestFailoverRecoveryPlan operation.

type TestFailoverRecoveryPlanRequest struct {
	// (required) The external identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) The test failover recovery plan action is used to perform a test failover on the recovery plan.
	Body *import4.TestFailoverSpec
}
