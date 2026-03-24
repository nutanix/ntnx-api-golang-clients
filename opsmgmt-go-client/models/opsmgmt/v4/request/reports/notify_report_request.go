package reports

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the NotifyReport operation.

type NotifyReportRequest struct {
	// (required) UUID of the report.
	ExtId *string

	// (required) The request body to notify the recipients with the specified report. The recipients and recipient formats are required
	// fields.
	Body *import1.ReportNotificationSpec
}
