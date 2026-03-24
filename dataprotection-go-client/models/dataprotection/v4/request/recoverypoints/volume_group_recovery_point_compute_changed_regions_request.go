package recoverypoints

import (
	import7 "github.com/nutanix/ntnx-api-golang-clients/dataprotection-go-client/v4/models/dataprotection/v4/content"
)

// This file holds the request struct for the VolumeGroupRecoveryPointComputeChangedRegions operation.

type VolumeGroupRecoveryPointComputeChangedRegionsRequest struct {
	// (required) The external identifier that can be used to retrieve the recovery point using its URL.
	RecoveryPointExtId *string

	// (required) External identifier of the volume group recovery point.
	VolumeGroupRecoveryPointExtId *string

	// (required) Disk recovery point identifier.
	ExtId *string

	// (required) Compute changed region parameters. These parameters allow you to specify a start offset, length, block size, and a
	// reference disk recovery point. All parameters are optional. However, if you need to set a reference disk recovery point,
	// you must specify all three parameters: recovery point ID, volume group recovery point ID, and disk recovery point ID.
	Body *import7.VolumeGroupRecoveryPointChangedRegionsComputeSpec
}
