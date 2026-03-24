package recoveryplans

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the UpdateRecoverySettingById operation.

type UpdateRecoverySettingByIdRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) recovery setting external identifier.
	ExtId *string

	// (required) Details of the request body to update the custom recovery setting identified by {extId}.
	Body *import1.RecoverySetting
}
