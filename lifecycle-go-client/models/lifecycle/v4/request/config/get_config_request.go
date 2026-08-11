package config

// This file holds the request struct for the GetConfig operation.

type GetConfigRequest struct {
	// The cluster UUID on which the resource is present or the operation is being performed.
	XClusterId *string
}
