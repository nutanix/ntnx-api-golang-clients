package config

// This file holds the request struct for the GetConfig operation.

type GetConfigRequest struct {
	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string
}
