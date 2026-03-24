package protectionpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the UpdateConsistencyRuleById operation.

type UpdateConsistencyRuleByIdRequest struct {
	// (required) The external identifier of the protection policy.
	ProtectionPolicyExtId *string

	// (required) The external identifier of the consistency rule.
	ExtId *string

	// (required) There are many scenarios in which it is essential to capture an application's state as an aggregate of the internal
	// states of a group of related entities at a specific moment in time. The consistency rule is a collection of all the
	// entities whose snapshot represents the application state at that point in time.
	Body *import1.ConsistencyRule
}
