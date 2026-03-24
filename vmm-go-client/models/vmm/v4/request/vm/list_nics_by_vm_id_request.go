package vm

// This file holds the request struct for the ListNicsByVmId operation.

type ListNicsByVmIdRequest struct {
	// (required) A globally unique identifier of a VM of type UUID.
	VmExtId *string

	// A URL query parameter that specifies the page number of the result set. It must be a positive integer between 0 and the
	// maximum number of pages that are available for that resource. Any number out of this range might lead to no results.
	Page_ *int

	// A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer
	// between 1 and 100. Any number out of this range will lead to a validation error. If the limit is not provided, a default
	// value of 50 records will be returned in the result set.
	Limit_ *int

	// A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is
	// evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the
	// response. Expression specified with the $filter must conform to the [OData
	// V4.01](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part1-protocol.html) URL conventions.
	// For example, filter **$filter=name eq 'karbon-ntnx-1.0'** would filter the result on cluster name 'karbon-ntnx1.0',
	// filter **$filter=startswith(name, 'C')** would filter on cluster name starting with 'C'.
	Filter_ *string
}
