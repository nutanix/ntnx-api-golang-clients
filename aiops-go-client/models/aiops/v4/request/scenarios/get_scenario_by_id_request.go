package scenarios

// This file holds the request struct for the GetScenarioById operation.

type GetScenarioByIdRequest struct {
	// (required) UUID of a capacity planning scenario.
	ExtId *string
}
