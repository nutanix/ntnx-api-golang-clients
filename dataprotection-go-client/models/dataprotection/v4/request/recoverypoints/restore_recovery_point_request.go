package recoverypoints

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
)

// This file holds the request struct for the RestoreRecoveryPoint operation.

type RestoreRecoveryPointRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	ExtId *string

	// Specification for the restore action on the top-level recovery point. For a recovery point that contains multiple VM or
	// volume group recovery points, users can selectively trigger restore for any specific set of VM or volume group recovery
	// points. In case no particular selection is made, all VM and volume group recovery points that are a part of the
	// top-level recovery point are restored.
	Body *import1.RecoveryPointRestorationSpec
}
