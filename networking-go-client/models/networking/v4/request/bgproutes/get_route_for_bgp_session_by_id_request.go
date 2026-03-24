package bgproutes

// This file holds the request struct for the GetRouteForBgpSessionById operation.

type GetRouteForBgpSessionByIdRequest struct {
	// (required) Reference for the route.
	ExtId *string

	// (required) Reference for the BGP session.
	BgpSessionExtId *string
}
