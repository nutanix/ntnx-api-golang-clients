package templateplacementpolicies

import (
	import16 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/config"
)

// This file holds the request struct for the UpdateTemplatePlacementPolicyById operation.

type UpdateTemplatePlacementPolicyByIdRequest struct {
	// (required) The external identifier of the template placement policy.
	ExtId *string

	// (required) The following parameters are required to update a template placement policy.
	Body *import16.TemplatePlacementPolicy
}
