package scenarios

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
)

// This file holds the request struct for the CreateScenario operation.

type CreateScenarioRequest struct {
	// (required) Capacity planning scenario sent for creation.
	Body *import1.Scenario
}
