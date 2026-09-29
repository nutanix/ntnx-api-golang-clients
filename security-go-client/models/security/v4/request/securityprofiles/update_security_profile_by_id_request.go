package securityprofiles

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/security-go-client/v4/models/security/v4/config"
)

// This file holds the request struct for the UpdateSecurityProfileById operation.

type UpdateSecurityProfileByIdRequest struct {
	// (required) External identifier for cluster security profiles status.
	ExtId *string

	// (required) Security Profiles configure security settings on AOS and AHV. Based on the selected profile, corresponding security
	// configurations are applied to the cluster. Available profiles are Standard, Elevated, and Strict.
	Body *import3.SecurityProfile
}
