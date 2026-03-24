package routes

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateRouteForRouteTable operation.

type CreateRouteForRouteTableRequest struct {
	// (required) Reference for the route table.
	RouteTableExtId *string

	// (required) Request body for creating a route for a specified route table.
	Body *import4.Route
}
