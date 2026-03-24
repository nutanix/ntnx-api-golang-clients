package clusters

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/clustermgmt-go-client/v4/models/clustermgmt/v4/config"
)

// This file holds the request struct for the UpdateRsyslogServerById operation.

type UpdateRsyslogServerByIdRequest struct {
	// (required) Indicates the UUID of a cluster.
	ClusterExtId *string

	// (required) RSYSLOG server UUID.
	ExtId *string

	// (required) RSYSLOG server to update.
	Body *import1.RsyslogServer
}
