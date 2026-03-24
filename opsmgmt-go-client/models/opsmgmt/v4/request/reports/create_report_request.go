package reports

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the CreateReport operation.

type CreateReportRequest struct {
	// (required) The request body to generate a new report. The name, config UUID, start and end time are required fields.
	Body *import1.Report
}
