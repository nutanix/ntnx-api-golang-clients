package securityconfigsvisibility

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the UpdateSecurityConfigsVisibilitySetting operation.

type UpdateSecurityConfigsVisibilitySettingRequest struct {
	// (required) Contains the configuration for PC UI visibility status of all security configurations settings.
	Body *import3.SecurityConfigVisibilitySetting
}
