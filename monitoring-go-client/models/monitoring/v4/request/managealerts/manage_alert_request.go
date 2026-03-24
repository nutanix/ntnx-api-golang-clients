package managealerts

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/monitoring-go-client/v4/models/monitoring/v4/serviceability"
)

// This file holds the request struct for the ManageAlert operation.

type ManageAlertRequest struct {
	// (required) Unique identifier of an alert that can be resolved or acknowledged.
	ExtId *string

	// (required) An alert can be resolved or acknowledged using the following parameters.
	Body *import1.AlertActionSpec
}
