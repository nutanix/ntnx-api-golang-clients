package ipfixexporters

// This file holds the request struct for the GetIpfixExporterById operation.

type GetIpfixExporterByIdRequest struct {
	// (required) UUID of IPFIX exporter.
	ExtId *string
}
