package reportconfig

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the UpdateReportConfigById operation.

type UpdateReportConfigByIdRequest struct {
	// (required) Report configuration UUID.
	ExtId *string

	// (required) The request body for report configuration update consists of name, sections, and other details about the configuration.
	Body *import1.ReportConfig
}
