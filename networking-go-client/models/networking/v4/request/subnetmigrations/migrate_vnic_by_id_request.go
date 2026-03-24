package subnetmigrations

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the MigrateVnicById operation.

type MigrateVnicByIdRequest struct {
	// (required) UUID of NIC to be migrated.
	ExtId *string

	// (required) Request body for the vNIC migrate operation.
	Body *import4.VnicMigrationItemSpec
}
