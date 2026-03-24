package upgradeselections

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
)

// This file holds the request struct for the CreateUpgradeSelection operation.

type CreateUpgradeSelectionRequest struct {
	// (required)
	Body *import1.UpgradeSelection
}
