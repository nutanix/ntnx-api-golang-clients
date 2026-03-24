package recoveryplans

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the UpdateRecoveryStageById operation.

type UpdateRecoveryStageByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) External identifier of the Recovery stage.
	ExtId *string

	// (required) Details of the request body to update the Recovery stage identified by {extId}.
	Body *import1.RecoveryStage
}
