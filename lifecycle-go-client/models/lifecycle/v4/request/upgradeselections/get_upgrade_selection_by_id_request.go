package upgradeselections

// This file holds the request struct for the GetUpgradeSelectionById operation.

type GetUpgradeSelectionByIdRequest struct {
	// (required) ExtId of the LCM Upgrade Selection
	ExtId *string
}
