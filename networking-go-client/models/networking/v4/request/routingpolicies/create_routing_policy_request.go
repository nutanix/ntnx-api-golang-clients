package routingpolicies

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateRoutingPolicy operation.

type CreateRoutingPolicyRequest struct {
	// (required) Schema to configure a routing policy.
	Body *import4.RoutingPolicy
}
