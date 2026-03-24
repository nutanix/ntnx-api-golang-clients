package routingpolicystats

import (
	import13 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/stats"
)

// This file holds the request struct for the ClearRoutingPolicyCounters operation.

type ClearRoutingPolicyCountersRequest struct {
	// (required) VPC UUID to reset all routing policy counters to zero.
	Body *import13.RoutingPolicyClearCountersSpec
}
