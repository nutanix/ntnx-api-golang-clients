package imageratelimitpolicies

// This file holds the request struct for the GetRateLimitPolicyById operation.

type GetRateLimitPolicyByIdRequest struct {
	// (required) The external identifier of image rate limit policy.
	ExtId *string
}
