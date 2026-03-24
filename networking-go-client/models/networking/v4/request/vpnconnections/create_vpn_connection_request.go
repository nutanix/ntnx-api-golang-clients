package vpnconnections

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateVpnConnection operation.

type CreateVpnConnectionRequest struct {
	// (required) Create VPN connection request body
	Body *import4.VpnConnection
}
