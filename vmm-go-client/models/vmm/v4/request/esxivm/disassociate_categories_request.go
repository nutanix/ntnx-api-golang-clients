package esxivm

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/esxi/config"
)

// This file holds the request struct for the DisassociateCategories operation.

type DisassociateCategoriesRequest struct {
	// (required) The globally unique identifier of an instance of type UUID.
	ExtId *string

	// (required)
	Body *import4.DisassociateVmCategoriesParams
}
