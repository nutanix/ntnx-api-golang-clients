package scenarios

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/aiops-go-client/v4/models/aiops/v4/config"
)

// This file holds the request struct for the CreateSimulation operation.

type CreateSimulationRequest struct {
	// (required) Indicates the parameters required for creating the simulation.
	Body *import1.Simulation
}
