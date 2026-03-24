package subnets

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateSubnetById operation.

type UpdateSubnetByIdRequest struct {
	// (required) UUID of the subnet.
	ExtId *string

	// (required) Request schema to update the specified subnet.
	Body *import4.Subnet
}
