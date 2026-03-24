package systemdefinedchecks

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the RunSystemDefinedChecks operation.

type RunSystemDefinedChecksRequest struct {
	// (required) Unique identifier for the cluster for which run System-Defined Checks is requested.
	ClusterExtId *string

	// (required) Required parameters for running the System-Defined Checks.
	Body *import1.RunSystemDefinedChecksSpec
}
