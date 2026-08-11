package dashboard

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the UpdateDashboardById operation.

type UpdateDashboardByIdRequest struct {
	// (required) Dashboard ID.
	ExtId *string

	// (required) Request body for updating a dashboard.
	Body *import1.Dashboard
}
