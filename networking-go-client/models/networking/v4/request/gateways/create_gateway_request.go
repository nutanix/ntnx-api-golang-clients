package gateways

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateGateway operation.

type CreateGatewayRequest struct {
	// (required) Create gateway request body.
	Body *import4.Gateway
}
