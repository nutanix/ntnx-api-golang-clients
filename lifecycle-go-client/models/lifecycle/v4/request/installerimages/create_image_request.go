package installerimages

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/config"
)

// This file holds the request struct for the CreateImage operation.

type CreateImageRequest struct {
	// (required)
	Body *import3.Image
}
