package bridges

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the MigrateBridge operation.

type MigrateBridgeRequest struct {
	// (required) Schema of bridge to migrate to a Virtual Switch
	Body *import4.Bridge

	// Prism Element cluster reference. This header can be optionally supplied for Virtual Switch list requests, but is
	// deprecated for all other Virtual Switch fetch/create/update/delete requests.
	XClusterId *string
}
