package templateplacementpolicies

import (
	import16 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/config"
)

// This file holds the request struct for the CreateTemplatePlacementPolicy operation.

type CreateTemplatePlacementPolicyRequest struct {
	// (required) The following parameters are required to create a template placement policy.
	Body *import16.TemplatePlacementPolicy
}
