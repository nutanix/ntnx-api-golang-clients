package ipfixexporters

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateIpfixExporterById operation.

type UpdateIpfixExporterByIdRequest struct {
	// (required) UUID of IPFIX exporter.
	ExtId *string

	// (required) Request schema to update the specified IPFIX exporter.
	Body *import4.IPFIXExporter
}
