package status

// This file holds the request struct for the GetStatus operation.

type GetStatusRequest struct {
	// The cluster UUID on which the resource is present or the operation is being performed.
	XClusterId *string
}
