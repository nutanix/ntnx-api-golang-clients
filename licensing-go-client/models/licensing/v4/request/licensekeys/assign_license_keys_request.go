package licensekeys

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/licensing-go-client/v4/models/licensing/v4/config"
)

// This file holds the request struct for the AssignLicenseKeys operation.

type AssignLicenseKeysRequest struct {
	// (required) Request body containing the cluster and license key consumption mapping.
	Body *import3.LicenseKeyAssignmentSpec
}
