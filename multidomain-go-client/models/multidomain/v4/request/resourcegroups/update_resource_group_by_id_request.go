package resourcegroups

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the UpdateResourceGroupById operation.

type UpdateResourceGroupByIdRequest struct {
	// (required) UUID of the Resource Group.
	ExtId *string

	// (required) The required parameters to update a Resource Group.
	Body *import3.ResourceGroup
}
