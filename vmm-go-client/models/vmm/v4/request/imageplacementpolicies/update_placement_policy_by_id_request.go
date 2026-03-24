package imageplacementpolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
)

// This file holds the request struct for the UpdatePlacementPolicyById operation.

type UpdatePlacementPolicyByIdRequest struct {
	// (required) The external identifier of image placement policy.
	ExtId *string

	// (required) Updated image placement policy request.
	Body *import6.PlacementPolicy
}
