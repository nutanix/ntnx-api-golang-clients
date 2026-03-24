package subnetmigrations

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the MigrateSubnets operation.

type MigrateSubnetsRequest struct {
	// (required) Request body for subnet migrate operation.
	Body *import4.VlanSubnetMigrationSpec
}
