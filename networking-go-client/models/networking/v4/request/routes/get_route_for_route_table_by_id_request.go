package routes

// This file holds the request struct for the GetRouteForRouteTableById operation.

type GetRouteForRouteTableByIdRequest struct {
	// (required) Reference for the route.
	ExtId *string

	// (required) Reference for the route table.
	RouteTableExtId *string
}
