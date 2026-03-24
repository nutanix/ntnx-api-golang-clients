package entitygroups

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateEntityGroupById operation.

type UpdateEntityGroupByIdRequest struct {
	// (required) The external identifier of the Entity Group.
	ExtId *string

	// (required) Updates the details of the Entity Group.
	Body *import1.EntityGroup
}
