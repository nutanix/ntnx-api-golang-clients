package reportartifacts

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/opsmgmt-go-client/v4/models/opsmgmt/v4/content"
)

// This file holds the request struct for the CreateReportArtifact operation.

type CreateReportArtifactRequest struct {
	// (required) The request body to create a report artifact entity. The type and filetype are required fields.
	Body *import3.ReportArtifact
}
