package resourcegroups

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the CreateResourceGroup operation.

type CreateResourceGroupRequest struct {
	// (required) The required parameters to create a Resource Group.
	Body *import3.ResourceGroup
}
