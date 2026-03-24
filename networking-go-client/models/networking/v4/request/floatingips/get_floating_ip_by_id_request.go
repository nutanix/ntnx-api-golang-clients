package floatingips

// This file holds the request struct for the GetFloatingIpById operation.

type GetFloatingIpByIdRequest struct {
	// (required) ExtId of the floating IP.
	ExtId *string
}
