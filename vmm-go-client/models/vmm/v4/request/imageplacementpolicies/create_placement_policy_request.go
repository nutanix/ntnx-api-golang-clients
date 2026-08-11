package imageplacementpolicies

import (
	import8 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
)

// This file holds the request struct for the CreatePlacementPolicy operation.

type CreatePlacementPolicyRequest struct {
	// (required) Request to create an image placement policy.
	Body *import8.PlacementPolicy
}
