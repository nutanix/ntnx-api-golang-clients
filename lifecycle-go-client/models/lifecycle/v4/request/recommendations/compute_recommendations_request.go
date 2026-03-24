package recommendations

import (
	import1 "github.com/nutanix/ntnx-api-golang-clients/lifecycle-go-client/v4/models/lifecycle/v4/resources"
)

// This file holds the request struct for the ComputeRecommendations operation.

type ComputeRecommendationsRequest struct {
	// (required)
	Body *import1.RecommendationSpec

	// Cluster uuid on which the resource is present or operation is being performed.
	XClusterId *string
}
