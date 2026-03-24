package imageratelimitpolicies

// This file holds the request struct for the DeleteRateLimitPolicyById operation.

type DeleteRateLimitPolicyByIdRequest struct {
	// (required) The external identifier of image rate limit policy.
	ExtId *string
}
