package vmstartuppolicies

import (
	import21 "github.com/nutanix/ntnx-api-golang-clients/vmm-go-client/v4/models/vmm/v4/ahv/policies"
)

// This file holds the request struct for the CreateVmStartupPolicy operation.

type CreateVmStartupPolicyRequest struct {
	// (required)
	Body *import21.VmStartupPolicy
}
