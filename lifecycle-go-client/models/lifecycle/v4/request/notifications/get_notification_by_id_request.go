package notifications

// This file holds the request struct for the GetNotificationById operation.

type GetNotificationByIdRequest struct {
	// (required) UUID of LCM Upgrade Notification resource.
	ExtId *string
}
