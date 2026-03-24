package servicegroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateServiceGroupById operation.

type UpdateServiceGroupByIdRequest struct {
	// (required) Service Group UUID.
	ExtId *string

	// (required) Updates the Service Group with the provided ExtID.
	Body *import1.ServiceGroup
}
