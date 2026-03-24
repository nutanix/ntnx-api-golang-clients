package scenarios

// This file holds the request struct for the GetScenarioReport operation.

type GetScenarioReportRequest struct {
	// (required) UUID of the capacity planning scenario to get the report.
	ScenarioExtId *string
}
