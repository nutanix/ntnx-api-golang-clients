package notifications

// This file holds the request struct for the GetNotificationById operation.

type GetNotificationByIdRequest struct {
	// (required) The resource identifier (UUID) of the computed upgrade notification. Obtained from the completion_details of the task
	// returned by POST /$actions/compute-notifications.
	ExtId *string
}
