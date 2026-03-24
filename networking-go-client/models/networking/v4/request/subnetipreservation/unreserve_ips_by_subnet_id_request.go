package subnetipreservation

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UnreserveIpsBySubnetId operation.

type UnreserveIpsBySubnetIdRequest struct {
	// (required) UUID of the subnet.
	ExtId *string

	// (required) Request schema to unreserve IP addresses on a subnet.
	Body *import4.IpUnreserveSpec
}
