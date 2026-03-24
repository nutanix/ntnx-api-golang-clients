package recoveryplans

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the CreateRecoveryPlan operation.

type CreateRecoveryPlanRequest struct {
	// (required) A recovery plan orchestrates disaster recovery of protected VMs and volume groups on primary Nutanix clusters to
	// secondary Nutanix clusters registered to the same or different Domain manager.
	Body *import1.RecoveryPlan
}
