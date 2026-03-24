package bgpsessions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the CreateBgpSession operation.

type CreateBgpSessionRequest struct {
	// (required) Create BGP session request body.
	Body *import4.BgpSession
}
