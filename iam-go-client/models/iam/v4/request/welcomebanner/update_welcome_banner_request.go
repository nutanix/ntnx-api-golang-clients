package welcomebanner

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/iam-go-client/v4/models/iam/v4/authn"
)

// This file holds the request struct for the UpdateWelcomeBanner operation.

type UpdateWelcomeBannerRequest struct {
	// (required) Information for the update welcome banner request.
	Body *import4.WelcomeBanner
}
