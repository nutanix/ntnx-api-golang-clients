package entities

import (
	import6 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/common"
)

// This file holds the request struct for the PreloadArtifacts operation.

type PreloadArtifactsRequest struct {
	// (required)
	Body *import6.PreloadSpec

	// The cluster UUID on which the resource is present or the operation is being performed.
	XClusterId *string
}
