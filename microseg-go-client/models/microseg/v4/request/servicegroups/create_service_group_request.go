package servicegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the CreateServiceGroup operation.

type CreateServiceGroupRequest struct {
	// (required)
	Body *import1.ServiceGroup
}
