package directoryserverconfigs

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/microseg-go-client/v4/models/microseg/v4/config"
)

// This file holds the request struct for the UpdateDsCategoryMappingById operation.

type UpdateDsCategoryMappingByIdRequest struct {
	// (required) Category Mapping UUID.
	ExtId *string

	// (required) Updates the category to directory configuration mapping information with the provided ExtID.
	Body *import1.CategoryMapping
}
