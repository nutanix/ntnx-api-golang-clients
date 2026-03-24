package recoverypoints

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/config"
)

// This file holds the request struct for the SetRecoveryPointExpirationTime operation.

type SetRecoveryPointExpirationTimeRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	ExtId *string

	// (required) Request body details to set the expiration time for the recovery point.
	Body *import1.ExpirationTimeSpec
}
