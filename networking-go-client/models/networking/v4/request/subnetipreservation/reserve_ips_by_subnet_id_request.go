package subnetipreservation

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the ReserveIpsBySubnetId operation.

type ReserveIpsBySubnetIdRequest struct {
	// (required) UUID of the subnet.
	ExtId *string

	// (required) Request schema to reserve IP addresses on a subnet.
	Body *import4.IpReserveSpec
}
