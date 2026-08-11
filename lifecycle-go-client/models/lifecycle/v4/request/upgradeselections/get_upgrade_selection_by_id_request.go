package upgradeselections

// This file holds the request struct for the GetUpgradeSelectionById operation.

type GetUpgradeSelectionByIdRequest struct {
	// (required) The external identifier (UUID) of the upgrade selection.
	ExtId *string
}
