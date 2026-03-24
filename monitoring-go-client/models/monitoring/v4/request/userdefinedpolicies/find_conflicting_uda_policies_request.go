package userdefinedpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the FindConflictingUdaPolicies operation.

type FindConflictingUdaPoliciesRequest struct {
	// (required) User-Defined Alert policy sent as a part of the request body to find policies that have a conflicting criteria.
	Body *import1.UserDefinedPolicy
}
