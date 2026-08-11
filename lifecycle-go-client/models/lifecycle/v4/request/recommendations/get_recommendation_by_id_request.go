package recommendations

// This file holds the request struct for the GetRecommendationById operation.

type GetRecommendationByIdRequest struct {
	// (required) The resource identifier (UUID) of the computed upgrade recommendation. Obtained from the completion_details of the task
	// returned by POST /$actions/compute-recommendations.
	ExtId *string
}
