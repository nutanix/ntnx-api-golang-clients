package clusterlogs

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the CollectLogs operation.

type CollectLogsRequest struct {
	// (required) The UUID of the cluster for which you want to collect logs.
	ExtId *string

	// (required) Input parameters for collecting logs.
	Body *import1.LogCollectionSpec
}
