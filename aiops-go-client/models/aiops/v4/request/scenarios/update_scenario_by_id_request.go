package scenarios

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
)

// This file holds the request struct for the UpdateScenarioById operation.

type UpdateScenarioByIdRequest struct {
	// (required) UUID of a capacity planning scenario.
	ExtId *string

	// (required) Capacity planning scenario sent for updation.
	Body *import1.Scenario
}
