package imageratelimitpolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
)

// This file holds the request struct for the UpdateRateLimitPolicyById operation.

type UpdateRateLimitPolicyByIdRequest struct {
	// (required) The external identifier of image rate limit policy.
	ExtId *string

	// (required) Updated image rate limit policy request.
	Body *import6.RateLimitPolicy
}
