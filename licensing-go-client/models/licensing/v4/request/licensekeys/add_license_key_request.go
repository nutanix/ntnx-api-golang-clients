package licensekeys

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
)

// This file holds the request struct for the AddLicenseKey operation.

type AddLicenseKeyRequest struct {
	// (required) Payload for persisting the license key information.
	Body *import3.LicenseKey

	// A URL query parameter that allows long running operations to execute in a dry-run mode providing ability to identify
	// trouble spots and system failures without performing the actual operation. Additionally this mode also offers a summary
	// snapshot of the resultant system in order to better understand how things fit together. The operation runs in dry-run
	// mode only if the provided value is true.
	Dryrun_ *bool
}
