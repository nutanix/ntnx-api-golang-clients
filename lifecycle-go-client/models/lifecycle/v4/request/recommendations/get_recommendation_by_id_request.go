package recommendations

// This file holds the request struct for the GetRecommendationById operation.

type GetRecommendationByIdRequest struct {
	// (required) UUID of LCM Upgrade recommendations resource.
	ExtId *string
}
