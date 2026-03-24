package virtualswitches

// This file holds the request struct for the GetVirtualSwitchById operation.

type GetVirtualSwitchByIdRequest struct {
	// (required) UUID of Virtual Switch
	ExtId *string

	// Prism Element cluster reference. This header can be optionally supplied for Virtual Switch list requests, but is
	// deprecated for all other Virtual Switch fetch/create/update/delete requests.
	XClusterId *string
}
