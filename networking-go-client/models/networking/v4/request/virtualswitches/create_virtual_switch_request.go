package virtualswitches

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateVirtualSwitch operation.

type CreateVirtualSwitchRequest struct {
	// (required) Schema to configure a virtual switch
	Body *import4.VirtualSwitch

	// Prism Element cluster reference. This header can be optionally supplied for Virtual Switch list requests, but is
	// deprecated for all other Virtual Switch fetch/create/update/delete requests.
	XClusterId *string
}
