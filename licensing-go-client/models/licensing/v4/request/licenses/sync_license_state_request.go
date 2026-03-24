package licenses

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
)

// This file holds the request struct for the SyncLicenseState operation.

type SyncLicenseStateRequest struct {
	// (required) For seamless licensing, the payload contains the action to be performed, cluster IDs, and other details.
	Body *import3.LicenseStateSyncSpec
}
