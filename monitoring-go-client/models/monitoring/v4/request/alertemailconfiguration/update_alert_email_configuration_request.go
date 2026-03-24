package alertemailconfiguration

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the UpdateAlertEmailConfiguration operation.

type UpdateAlertEmailConfigurationRequest struct {
	// (required) Email configuration sent for the update.
	Body *import1.AlertEmailConfiguration
}
