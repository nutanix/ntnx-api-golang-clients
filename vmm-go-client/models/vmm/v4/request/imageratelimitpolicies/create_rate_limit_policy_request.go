package imageratelimitpolicies

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/images/config"
)

// This file holds the request struct for the CreateRateLimitPolicy operation.

type CreateRateLimitPolicyRequest struct {
	// (required) Request to create an image rate limit policy.
	Body *import6.RateLimitPolicy
}
