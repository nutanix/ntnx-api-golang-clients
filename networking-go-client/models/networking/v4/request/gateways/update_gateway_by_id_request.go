package gateways

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateGatewayById operation.

type UpdateGatewayByIdRequest struct {
	// (required) Gateway UUID.
	ExtId *string

	// (required) Update gateway request body.
	Body *import4.Gateway
}
