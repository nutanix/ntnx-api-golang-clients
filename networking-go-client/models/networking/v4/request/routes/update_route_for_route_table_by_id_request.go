package routes

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateRouteForRouteTableById operation.

type UpdateRouteForRouteTableByIdRequest struct {
	// (required) Reference for the route.
	ExtId *string

	// (required) Reference for the route table.
	RouteTableExtId *string

	// (required) Request body of the route to update.
	Body *import4.Route
}
