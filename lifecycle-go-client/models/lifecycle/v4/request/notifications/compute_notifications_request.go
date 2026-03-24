package notifications

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
)

// This file holds the request struct for the ComputeNotifications operation.

type ComputeNotificationsRequest struct {
	// (required)
	Body *import1.NotificationsSpec

	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string
}
