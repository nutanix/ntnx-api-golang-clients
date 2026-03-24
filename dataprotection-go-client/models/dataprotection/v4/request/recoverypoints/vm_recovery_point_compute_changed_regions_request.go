package recoverypoints

import (
	import7 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/content"
)

// This file holds the request struct for the VmRecoveryPointComputeChangedRegions operation.

type VmRecoveryPointComputeChangedRegionsRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	RecoveryPointExtId *string

	// (required) The external identifier that can be used to identify a VM recovery point.
	VmRecoveryPointExtId *string

	// (required) Disk recovery point identifier.
	ExtId *string

	// (required) Compute changed region parameters. These parameters allow you to specify a start offset, length, block size, and a
	// reference disk recovery point. All parameters are optional. However, if you need to set a reference disk recovery point,
	// you must specify all three parameters: recovery point ID, VM recovery point ID, and disk recovery point ID.
	Body *import7.VmRecoveryPointChangedRegionsComputeSpec
}
