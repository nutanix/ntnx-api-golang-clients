package recoverypoints

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
)

// This file holds the request struct for the CreateRecoveryPoint operation.

type CreateRecoveryPointRequest struct {
	// (required) Details of the request body to create a recovery point.
	Body *import1.RecoveryPoint
}
