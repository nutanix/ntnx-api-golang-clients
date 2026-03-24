package licensekeys

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
)

// This file holds the request struct for the AssociateLicenseKeys operation.

type AssociateLicenseKeysRequest struct {
	// (required) Path parameter for the license key ID.
	ExtId *string

	// (required) Payload for associating the license key.
	Body *import3.AssociateLicenseKeySpec
}
