package systemdefinedpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the UpdateClusterConfigById operation.

type UpdateClusterConfigByIdRequest struct {
	// (required) Unique ID of the System-Defined Alert Policy.
	SystemDefinedPolicyExtId *string

	// (required) Cluster UUID.
	ExtId *string

	// (required)
	Body *import1.ClusterConfig
}
