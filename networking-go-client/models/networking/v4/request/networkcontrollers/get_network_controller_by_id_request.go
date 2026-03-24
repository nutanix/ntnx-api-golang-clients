package networkcontrollers

// This file holds the request struct for the GetNetworkControllerById operation.

type GetNetworkControllerByIdRequest struct {
	// (required) The UUID of the network controller
	ExtId *string
}
