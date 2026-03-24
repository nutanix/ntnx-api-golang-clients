package recoveryplans

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the UpdateRecoveryPlanById operation.

type UpdateRecoveryPlanByIdRequest struct {
	// (required) External identifier of the recovery plan.
	ExtId *string

	// (required) Updated recovery plan identified by {extId}. Recovery locations cannot be updated.
	Body *import1.RecoveryPlan
}
