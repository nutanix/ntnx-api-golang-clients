package ovas

import (
	import9 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/content"
)

// This file holds the request struct for the DeployOva operation.

type DeployOvaRequest struct {
	// (required) The external identifier for an OVA.
	ExtId *string

	// (required) Request to deploy a VM from an OVA.
	Body *import9.OvaDeploymentSpec
}
