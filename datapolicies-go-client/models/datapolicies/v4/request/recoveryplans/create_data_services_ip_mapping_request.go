package recoveryplans

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the CreateDataServicesIpMapping operation.

type CreateDataServicesIpMappingRequest struct {
	// (required) External identifier of the recovery plan.
	RecoveryPlanExtId *string

	// (required) Details of the request body to create a data services IP mapping.
	Body *import1.DataServicesIpMapping
}
