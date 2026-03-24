package status

// This file holds the request struct for the GetStatus operation.

type GetStatusRequest struct {
	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string
}
