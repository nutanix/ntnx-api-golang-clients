package trafficmirrors

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateTrafficMirrorById operation.

type UpdateTrafficMirrorByIdRequest struct {
	// (required) UUID of the session.
	ExtId *string

	// (required) Update Traffic mirror session request body.
	Body *import4.TrafficMirror
}
