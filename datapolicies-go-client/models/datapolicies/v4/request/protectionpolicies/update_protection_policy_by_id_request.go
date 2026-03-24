package protectionpolicies

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/datapolicies-go-client/v4/models/datapolicies/v4/config"
)

// This file holds the request struct for the UpdateProtectionPolicyById operation.

type UpdateProtectionPolicyByIdRequest struct {
	// (required) The external identifier of the protection policy.
	ExtId *string

	// (required) Updated protection policy identified by an external identifier.
	Body *import1.ProtectionPolicy
}
