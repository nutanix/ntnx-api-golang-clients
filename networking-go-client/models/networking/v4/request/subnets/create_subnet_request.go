package subnets

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateSubnet operation.

type CreateSubnetRequest struct {
	// (required) Request schema to create the subnet.
	Body *import4.Subnet
}
