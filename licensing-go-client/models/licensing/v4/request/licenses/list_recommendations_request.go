package licenses

// This file holds the request struct for the ListRecommendations operation.

type ListRecommendationsRequest struct {
	// A URL query parameter that specifies the page number of the result set. It must be a positive integer between 0 and the
	// maximum number of pages that are available for that resource. Any number out of this range might lead to no results.
	Page_ *int

	// A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer
	// between 1 and 100. Any number out of this range will lead to a validation error. If the limit is not provided, a default
	// value of 50 records will be returned in the result set.
	Limit_ *int

	// A URL query parameter that allows clients to request a specific set of properties for each entity or complex type.
	// Expression specified with the $select must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions. If a $select
	// expression consists of a single select item that is an asterisk (i.e., *), then all properties on the matching resource
	// will be returned.
	Select_ *string
}
