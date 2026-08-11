package upgradeselections

// This file holds the request struct for the ExportUpgradeSelection operation.

type ExportUpgradeSelectionRequest struct {
	// (required) The external identifier (UUID) of the upgrade selection.
	ExtId *string
}
