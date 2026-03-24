package ipfixexporters

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateIpfixExporter operation.

type CreateIpfixExporterRequest struct {
	// (required) Request schema to create the IPFIX exporter.
	Body *import4.IPFIXExporter
}
