package entities

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
)

// This file holds the request struct for the PreloadArtifacts operation.

type PreloadArtifactsRequest struct {
	// (required)
	Body *import4.PreloadSpec

	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string
}
