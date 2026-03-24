package trafficmirrors

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateTrafficMirror operation.

type CreateTrafficMirrorRequest struct {
	// (required) Create Traffic mirror session request body.
	Body *import4.TrafficMirror
}
