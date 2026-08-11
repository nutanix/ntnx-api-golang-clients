package imageplacementpolicies

import (
	import8 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
)

// This file holds the request struct for the SuspendPlacementPolicy operation.

type SuspendPlacementPolicyRequest struct {
	// (required) The external identifier of image placement policy.
	ExtId *string

	// (required)
	Body *import8.SuspendPlacementPolicyConfig
}
