package userdefinedpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the CreateUdaPolicy operation.

type CreateUdaPolicyRequest struct {
	// (required) User-Defined Alert policy sent as a part of the request body to create a new User-Defined Alert policy.
	Body *import1.UserDefinedPolicy
}
