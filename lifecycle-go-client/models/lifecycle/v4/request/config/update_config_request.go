package config

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
)

// This file holds the request struct for the UpdateConfig operation.

type UpdateConfigRequest struct {
	// (required)
	Body *import1.Config

	// The cluster UUID on which the resource is present or the operation is being performed.
	XClusterId *string
}
