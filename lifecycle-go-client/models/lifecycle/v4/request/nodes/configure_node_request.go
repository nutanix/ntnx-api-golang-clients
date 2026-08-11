package nodes

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
)

// This file holds the request struct for the ConfigureNode operation.

type ConfigureNodeRequest struct {
	// (required)
	Body *import3.ConfigureNodeSpec
}
