package networkfunctions

// This file holds the request struct for the GetNetworkFunctionById operation.

type GetNetworkFunctionByIdRequest struct {
	// (required) The UUID of the network function.
	ExtId *string
}
