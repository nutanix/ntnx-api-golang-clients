package scenarios

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
)

// This file holds the request struct for the UpdateSimulationById operation.

type UpdateSimulationByIdRequest struct {
	// (required) UUID of a simulation.
	ExtId *string

	// (required) Indicates the parameters required for updating the simulation.
	Body *import1.Simulation
}
