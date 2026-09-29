package securityprofiles

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the UpdateAdvancedConfig operation.

type UpdateAdvancedConfigRequest struct {
	// (required) External identifier for cluster security profiles status.
	SecurityProfileExtId *string

	// (required) Contains advanced configurations for a cluster, such as Core Dump Configs, Consent Banner etc.
	Body *import3.AdvancedConfig
}
