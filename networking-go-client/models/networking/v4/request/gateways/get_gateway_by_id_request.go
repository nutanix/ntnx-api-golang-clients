package gateways

// This file holds the request struct for the GetGatewayById operation.

type GetGatewayByIdRequest struct {
	// (required) Gateway UUID.
	ExtId *string
}
