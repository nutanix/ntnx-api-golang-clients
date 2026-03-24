package floatingips

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateFloatingIpById operation.

type UpdateFloatingIpByIdRequest struct {
	// (required) ExtId of the floating IP.
	ExtId *string

	// (required) Configure a floating IP.
	Body *import4.FloatingIp
}
