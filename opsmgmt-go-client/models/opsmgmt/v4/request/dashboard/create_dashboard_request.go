package dashboard

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the CreateDashboard operation.

type CreateDashboardRequest struct {
	// (required) Request body for creating a new dashboard.
	Body *import1.Dashboard
}
