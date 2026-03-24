package volumegroups

// This file holds the request struct for the ListCategoryAssociationsByVolumeGroupId operation.

//
// Deprecated: The API corresponding to this struct has been deprecated.
type ListCategoryAssociationsByVolumeGroupIdRequest struct {
	// (required) The external identifier of a Volume Group.
	VolumeGroupExtId *string

	// A URL query parameter that specifies the page number of the result set. It must be a positive integer between 0 and the
	// maximum number of pages that are available for that resource. Any number out of this range might lead to no results.
	Page_ *int

	// A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer
	// between 1 and 100. Any number out of this range will lead to a validation error. If the limit is not provided, a default
	// value of 50 records will be returned in the result set.
	Limit_ *int
}
