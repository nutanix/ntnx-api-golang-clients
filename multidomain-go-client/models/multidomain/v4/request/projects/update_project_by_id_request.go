package projects

import (
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/config"
)

// This file holds the request struct for the UpdateProjectById operation.

type UpdateProjectByIdRequest struct {
	// (required) ExtId of the project
	ExtId *string

	// (required)
	Body *import3.Project
}
