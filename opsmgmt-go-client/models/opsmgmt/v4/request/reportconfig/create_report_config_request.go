package reportconfig

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/config"
)

// This file holds the request struct for the CreateReportConfig operation.

type CreateReportConfigRequest struct {
	// (required) The request body for report configuration creation consists of name, sections, and other details about the
	// configuration.
	Body *import1.ReportConfig
}
