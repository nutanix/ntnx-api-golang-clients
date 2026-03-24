package awsvpcs

// This file holds the request struct for the ListAwsVpcs operation.

type ListAwsVpcsRequest struct {
	// (required) Target cluster UUID.
	XClusterId *string
}
