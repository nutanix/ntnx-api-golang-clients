package floatingips

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateFloatingIp operation.

type CreateFloatingIpRequest struct {
	// (required) Task Id corresponding to the Create Floating IP operation.
	Body *import4.FloatingIp
}
