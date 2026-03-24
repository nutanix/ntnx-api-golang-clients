package bgpsessions

import (
	import4 "github.com/nutanix/ntnx-api-golang-clients/networking-go-client/v4/models/networking/v4/config"
)

// This file holds the request struct for the UpdateBgpSessionById operation.

type UpdateBgpSessionByIdRequest struct {
	// (required) BGP session UUID.
	ExtId *string

	// (required) Update BGP session request body.
	Body *import4.BgpSession
}
