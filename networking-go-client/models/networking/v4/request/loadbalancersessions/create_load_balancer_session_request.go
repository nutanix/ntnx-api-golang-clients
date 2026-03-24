package loadbalancersessions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateLoadBalancerSession operation.

type CreateLoadBalancerSessionRequest struct {
	// (required) Request schema to create the load balancer session.
	Body *import4.LoadBalancerSession
}
